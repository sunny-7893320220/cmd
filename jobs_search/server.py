import os
import re
import time
import uuid
import json
import threading
import logging
from urllib.parse import urlparse, urljoin

import requests
from bs4 import BeautifulSoup
from flask import Flask, request, jsonify, send_from_directory
from flask_cors import CORS
from dotenv import load_dotenv

load_dotenv()

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")
logger = logging.getLogger(__name__)

app = Flask(__name__, static_folder=".", static_url_path="")
CORS(app)

HUNTER_API_KEY = os.getenv("HUNTER_API_KEY", "")

# In-memory job store
jobs: dict[str, dict] = {}

HEADERS = {
    "User-Agent": (
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 "
        "(KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
    ),
    "Accept-Language": "en-US,en;q=0.9",
}

# ─────────────────────────────────────────────
# Helpers
# ─────────────────────────────────────────────

def google_search(query: str, num: int = 5) -> list[str]:
    """Return up to `num` URLs from a DuckDuckGo HTML search (no API key needed)."""
    try:
        url = "https://html.duckduckgo.com/html/"
        params = {"q": query}
        resp = requests.post(url, data=params, headers=HEADERS, timeout=10)
        soup = BeautifulSoup(resp.text, "lxml")
        links = []
        for a in soup.select("a.result__url"):
            href = a.get("href", "")
            if href and href.startswith("http"):
                links.append(href)
            if len(links) >= num:
                break
        # fallback: result__a links
        if not links:
            for a in soup.select("a.result__a"):
                href = a.get("href", "")
                if href and href.startswith("http"):
                    links.append(href)
                if len(links) >= num:
                    break
        return links
    except Exception as e:
        logger.warning(f"google_search error for '{query}': {e}")
        return []


def find_company_website(company: str) -> str:
    """Find the official website of a company."""
    results = google_search(f"{company} official website")
    skip_domains = {"linkedin.com", "facebook.com", "twitter.com", "instagram.com",
                    "youtube.com", "wikipedia.org", "glassdoor.com", "indeed.com",
                    "crunchbase.com", "bloomberg.com", "ambitionbox.com"}
    for url in results:
        try:
            domain = urlparse(url).netloc.lower().replace("www.", "")
            if not any(s in domain for s in skip_domains):
                return url
        except Exception:
            continue
    return ""


def find_linkedin(company: str, person: str = "") -> str:
    """Find the LinkedIn company page or person profile URL."""
    if person:
        query = f"{person} {company} LinkedIn"
    else:
        query = f"{company} LinkedIn company page site:linkedin.com"
        
    results = google_search(query)
    
    pattern = "linkedin.com/in/" if person else "linkedin.com/company/"
    
    for url in results:
        if pattern in url:
            # Clean tracking params
            parsed = urlparse(url)
            return f"{parsed.scheme}://{parsed.netloc}{parsed.path}"
    
    # Fallback for company search
    if not person:
        results2 = google_search(f"{company} LinkedIn company")
        for url in results2:
            if "linkedin.com/company/" in url:
                parsed = urlparse(url)
                return f"{parsed.scheme}://{parsed.netloc}{parsed.path}"
    return ""


def scrape_emails_from_page(url: str) -> list[str]:
    """Scrape email addresses from a given URL."""
    try:
        resp = requests.get(url, headers=HEADERS, timeout=10)
        text = resp.text
        emails = re.findall(r"[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}", text)
        # Filter out image/asset emails and common non-HR patterns
        hr_keywords = ["hr", "careers", "jobs", "recruit", "talent", "people", "hiring", "contact", "info"]
        hr_emails = [e for e in emails if any(k in e.lower() for k in hr_keywords)]
        return list(dict.fromkeys(hr_emails or emails))  # deduplicate, prefer HR emails
    except Exception:
        return []


def scrape_phones_from_page(url: str) -> list[str]:
    """Scrape phone numbers from a given URL."""
    try:
        resp = requests.get(url, headers=HEADERS, timeout=10)
        text = resp.text
        # Remove HTML tags
        soup = BeautifulSoup(text, "lxml")
        plain = soup.get_text(" ")
        # Match common phone patterns
        phones = re.findall(
            r"(?:\+?\d[\d\s\-().]{7,}\d)",
            plain
        )
        cleaned = []
        for p in phones:
            p = p.strip()
            digits = re.sub(r"\D", "", p)
            if 7 <= len(digits) <= 15:
                cleaned.append(p)
        return list(dict.fromkeys(cleaned))[:3]
    except Exception:
        return []


def find_contact_page(base_url: str) -> str:
    """Try to find the contact/about page of a website."""
    contact_paths = ["/contact", "/contact-us", "/about", "/about-us",
                     "/team", "/hr", "/careers", "/jobs"]
    for path in contact_paths:
        try:
            url = urljoin(base_url, path)
            resp = requests.get(url, headers=HEADERS, timeout=8)
            if resp.status_code == 200 and len(resp.text) > 500:
                return url
        except Exception:
            continue
    return base_url


def find_hr_email_hunter(domain: str, person: str = "") -> str:
    """Use Hunter.io API to find email for a domain and optional person."""
    if not HUNTER_API_KEY or HUNTER_API_KEY == "your_hunter_api_key_here":
        return ""
    try:
        params = {
            "domain": domain,
            "api_key": HUNTER_API_KEY,
            "limit": 5,
        }
        if person:
            # If person is provided, we use the Email Finder API or search for them
            # For simplicity, we just look for them in the domain search results
            pass
        else:
            params["department"] = "human_resources"

        resp = requests.get(
            "https://api.hunter.io/v2/domain-search",
            params=params,
            timeout=10,
        )
        data = resp.json()
        emails = data.get("data", {}).get("emails", [])
        
        if person:
            name_parts = person.lower().split()
            for e in emails:
                val = e.get("value", "").lower()
                if all(p in val for p in name_parts):
                    return e.get("value", "")
        
        if emails:
            return emails[0].get("value", "")
    except Exception as e:
        logger.warning(f"Hunter.io error: {e}")
    return ""


def research_company(company: str, person: str = "") -> dict:
    """Full research pipeline for one company and optional person."""
    result = {
        "company": company,
        "person": person,
        "website": "",
        "linkedin": "",
        "hr_email": "",
        "hr_phone": "",
        "status": "searching",
    }

    # 1. Find website
    website = find_company_website(company)
    result["website"] = website
    time.sleep(1.5)

    # 2. Find LinkedIn
    linkedin = find_linkedin(company, person)
    result["linkedin"] = linkedin
    time.sleep(1.5)

    # 3. Find HR email
    hr_email = ""
    if website:
        domain = urlparse(website).netloc.replace("www.", "")
        # Try Hunter.io first
        hr_email = find_hr_email_hunter(domain, person)
        if not hr_email:
            # Scrape contact page
            contact_url = find_contact_page(website)
            emails = scrape_emails_from_page(contact_url)
            
            if person:
                name_parts = person.lower().split()
                matches = [e for e in emails if all(p in e.lower() for p in name_parts)]
                if matches:
                    hr_email = matches[0]

            if not hr_email and emails:
                hr_email = emails[0]
            if not hr_email and contact_url != website:
                emails2 = scrape_emails_from_page(website)
                if person:
                    name_parts = person.lower().split()
                    matches = [e for e in emails2 if all(p in e.lower() for p in name_parts)]
                    if matches:
                        hr_email = matches[0]
                if not hr_email and emails2:
                    hr_email = emails2[0]
    result["hr_email"] = hr_email
    time.sleep(1)

    # 4. Find HR phone
    hr_phone = ""
    if website:
        contact_url = find_contact_page(website)
        phones = scrape_phones_from_page(contact_url)
        if phones:
            hr_phone = phones[0]
    result["hr_phone"] = hr_phone

    result["status"] = "done"
    return result


# ─────────────────────────────────────────────
# Background worker
# ─────────────────────────────────────────────

def run_job(job_id: str, items: list[dict]):
    jobs[job_id]["status"] = "running"
    total = len(items)
    results = []

    for i, item in enumerate(items):
        company = item.get("company", "").strip()
        person = item.get("person", "").strip()
        
        if not company:
            continue
            
        logger.info(f"[{job_id}] Researching {i+1}/{total}: {company} (Person: {person})")
        try:
            data = research_company(company, person)
        except Exception as e:
            logger.error(f"Error researching {company}: {e}")
            data = {
                "company": company,
                "person": person,
                "website": "",
                "linkedin": "",
                "hr_email": "",
                "hr_phone": "",
                "status": "error",
            }
        results.append(data)
        jobs[job_id]["progress"] = round((i + 1) / total * 100)
        jobs[job_id]["results"] = results
        # Polite delay between companies (avoid rate limiting)
        time.sleep(2)

    jobs[job_id]["status"] = "done"
    logger.info(f"[{job_id}] Job complete. {len(results)} items processed.")


# ─────────────────────────────────────────────
# Routes
# ─────────────────────────────────────────────

@app.route("/")
def index():
    return send_from_directory(".", "index.html")


@app.route("/search", methods=["POST"])
def start_search():
    data = request.get_json()
    items = data.get("items", [])
    if not items:
        # Fallback for old format (list of strings)
        companies = data.get("companies", [])
        if companies:
            items = [{"company": c, "person": ""} for c in companies]
            
    if not items:
        return jsonify({"error": "No companies provided"}), 400
    if len(items) > 150:
        return jsonify({"error": "Max 150 companies per batch"}), 400

    job_id = str(uuid.uuid4())
    jobs[job_id] = {
        "status": "queued",
        "progress": 0,
        "results": [],
        "total": len(items),
    }
    thread = threading.Thread(target=run_job, args=(job_id, items), daemon=True)
    thread.start()
    return jsonify({"job_id": job_id})


@app.route("/status/<job_id>")
def job_status(job_id):
    job = jobs.get(job_id)
    if not job:
        return jsonify({"error": "Job not found"}), 404
    return jsonify({
        "status": job["status"],
        "progress": job["progress"],
        "total": job["total"],
        "completed": len(job["results"]),
        "results": job["results"],
    })


@app.route("/cancel/<job_id>", methods=["POST"])
def cancel_job(job_id):
    if job_id in jobs:
        jobs[job_id]["status"] = "cancelled"
    return jsonify({"ok": True})


if __name__ == "__main__":
    print("=" * 55)
    print("  Company Research Tool — http://localhost:5000")
    print("=" * 55)
    app.run(host="0.0.0.0", port=5000, debug=False)
