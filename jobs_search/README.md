# Company Research Tool

A web app that takes a list of companies and automatically finds:
- 🌐 **Official Website**
- 💼 **LinkedIn Company Page**
- 📧 **HR Email Address**
- 📞 **HR Phone Number**

---

## Quick Start

### 1. Install dependencies
```bash
cd /home/saikrishna/cmd/jobs_search
pip install -r requirements.txt
```

### 2. (Optional) Set up Hunter.io API key for better email finding
```bash
cp .env.example .env
# Edit .env and add your Hunter.io API key (free: 25 searches/month)
# https://hunter.io/api-keys
```

### 3. Start the server
```bash
python server.py
```

### 4. Open in browser
Go to: **http://localhost:5000**

---

## How to Use

1. **Upload** your company list as `.csv` or `.md`
2. Click **Start Research** — the tool searches ~10–15 sec per company
3. Watch results populate live in the table
4. Click **Export CSV** to download all results

---

## File Formats

### CSV (`companies.csv`)
```csv
company_name
Google
Microsoft
Infosys
```

### Markdown (`companies.md`)
```markdown
# Companies
- Google
- Microsoft
- Infosys
```

---

## How It Searches

| Data | Method |
|------|--------|
| Website | DuckDuckGo search → first non-social result |
| LinkedIn | DuckDuckGo `site:linkedin.com/company/` search |
| HR Email | Hunter.io API (if key set) → scrape contact page |
| HR Phone | Scrape company contact/about page |

> **Note**: LinkedIn blocks scraping, so only the company page URL is found (not profile details).

---

## Limits
- Max **150 companies per batch**
- ~10–15 seconds per company (polite rate limiting)
- 100 companies ≈ 20–25 minutes

---

## 🐳 Deploy on Minikube

### 1. Point Docker to Minikube's daemon
```bash
eval $(minikube docker-env)
```

### 2. Build the Docker image inside Minikube
```bash
cd /home/saikrishna/cmd/jobs_search
docker build -t jobs-search:latest .
```

### 3. (Optional) Add Hunter.io API key
Edit `k8s/deployment.yaml` and set the `HUNTER_API_KEY` value in the Secret.

### 4. Apply the manifests
```bash
kubectl apply -f k8s/deployment.yaml
```

### 5. Access the app
```bash
minikube service jobs-search-svc -n jobs-search
# OR
minikube ip   # then open http://<ip>:30080
```

### Useful commands
```bash
# Check pod status
kubectl get pods -n jobs-search

# View logs
kubectl logs -n jobs-search deploy/jobs-search -f

# Delete everything
kubectl delete namespace jobs-search
```

---

## Files
```
jobs_search/
├── Dockerfile       # Container image
├── .dockerignore
├── k8s/
│   └── deployment.yaml  # Namespace + Deployment + Service
├── server.py        # Flask backend
├── index.html       # Web UI
├── style.css        # Styling
├── app.js           # Frontend logic
├── requirements.txt # Python deps
├── companies.csv    # Sample CSV
├── companies.md     # Sample Markdown
└── .env.example     # API key config
```
