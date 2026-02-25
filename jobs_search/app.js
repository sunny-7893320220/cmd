// ─────────────────────────────────────────────
// State
// ─────────────────────────────────────────────
let items = []; // formerly companies
let jobId = null;
let pollInterval = null;
let allResults = [];

// ─────────────────────────────────────────────
// DOM refs
// ─────────────────────────────────────────────
const dropZone = document.getElementById("dropZone");
const fileInput = document.getElementById("fileInput");
const previewSection = document.getElementById("previewSection");
const previewCount = document.getElementById("previewCount");
const previewTags = document.getElementById("previewTags");
const clearBtn = document.getElementById("clearBtn");
const searchBtn = document.getElementById("searchBtn");
const progressSection = document.getElementById("progressSection");
const progressText = document.getElementById("progressText");
const progressPct = document.getElementById("progressPct");
const progressBar = document.getElementById("progressBar");
const cancelBtn = document.getElementById("cancelBtn");
const resultsSection = document.getElementById("resultsSection");
const resultsCount = document.getElementById("resultsCount");
const resultsBody = document.getElementById("resultsBody");
const exportBtn = document.getElementById("exportBtn");
const filterInput = document.getElementById("filterInput");

// Manual search
const quickSearchBtn = document.getElementById("quickSearchBtn");
const manualCompany = document.getElementById("manualCompany");
const manualPerson = document.getElementById("manualPerson");

// ─────────────────────────────────────────────
// File drag & drop
// ─────────────────────────────────────────────
dropZone.addEventListener("dragover", e => { e.preventDefault(); dropZone.classList.add("drag-over"); });
dropZone.addEventListener("dragleave", () => dropZone.classList.remove("drag-over"));
dropZone.addEventListener("drop", e => {
    e.preventDefault();
    dropZone.classList.remove("drag-over");
    const file = e.dataTransfer.files[0];
    if (file) handleFile(file);
});
dropZone.addEventListener("click", () => fileInput.click());
fileInput.addEventListener("change", () => { if (fileInput.files[0]) handleFile(fileInput.files[0]); });

// ─────────────────────────────────────────────
// File parsing
// ─────────────────────────────────────────────
function handleFile(file) {
    const name = file.name.toLowerCase();
    const reader = new FileReader();
    reader.onload = e => {
        const text = e.target.result;
        if (name.endsWith(".csv")) {
            parseCSV(text);
        } else if (name.endsWith(".md") || name.endsWith(".txt")) {
            parseMarkdown(text);
        } else {
            alert("Please upload a .csv or .md file.");
        }
    };
    reader.readAsText(file);
}

function parseCSV(text) {
    const result = Papa.parse(text.trim(), { header: true, skipEmptyLines: true });
    let parsedItems = [];

    const headers = result.meta.fields || [];
    const companyCol = headers.find(h => /company|org/i.test(h)) || headers[0];
    const personCol = headers.find(h => /name|person|contact/i.test(h) && h !== companyCol);

    if (companyCol && result.data.length) {
        parsedItems = result.data.map(row => ({
            company: (row[companyCol] || "").trim(),
            person: personCol ? (row[personCol] || "").trim() : ""
        })).filter(i => i.company);
    } else {
        // No header — try splitting by comma/tab
        const lines = text.split("\n").map(l => l.trim()).filter(l => l && !l.startsWith("#"));
        parsedItems = lines.map(line => {
            const parts = line.split(/[,\t]/);
            return {
                company: parts[0].trim(),
                person: parts[1] ? parts[1].trim() : ""
            };
        }).filter(i => i.company);
    }
    setItems(parsedItems);
}

function parseMarkdown(text) {
    const lines = text.split("\n");
    const parsedItems = [];
    for (const line of lines) {
        const stripped = line.trim();
        // Match bullet lists: "- Company, Person" or "* Company, Person"
        const bulletMatch = stripped.match(/^[-*•]\s+(.+)/) || stripped.match(/^\d+\.\s+(.+)/);
        if (bulletMatch) {
            const content = bulletMatch[1].trim();
            const parts = content.split(/[,|]/);
            parsedItems.push({
                company: parts[0].trim(),
                person: parts[1] ? parts[1].trim() : ""
            });
            continue;
        }
        // Match table rows: | Company | Person | ...
        if (stripped.startsWith("|") && !stripped.startsWith("|---") && !stripped.startsWith("| ---")) {
            const cells = stripped.split("|").map(c => c.trim()).filter(Boolean);
            if (cells.length && !/^company|^name/i.test(cells[0])) {
                parsedItems.push({
                    company: cells[0],
                    person: cells[1] || ""
                });
            }
            continue;
        }
        // Plain non-heading lines
        if (stripped && !stripped.startsWith("#") && stripped.length > 1) {
            const parts = stripped.split(/[,|]/);
            parsedItems.push({
                company: parts[0].trim(),
                person: parts[1] ? parts[1].trim() : ""
            });
        }
    }
    setItems(parsedItems);
}

function setItems(newItems) {
    items = newItems.slice(0, 150);
    if (items.length === 0) {
        alert("No company/person names found in the file. Check the format.");
        return;
    }
    renderPreview();
}

function renderPreview() {
    previewCount.textContent = `${items.length} items found`;
    previewTags.innerHTML = items.map(i => {
        const label = i.person ? `${i.company} (${i.person})` : i.company;
        return `<span class="tag">${escHtml(label)}</span>`;
    }).join("");
    previewSection.style.display = "block";
    previewSection.scrollIntoView({ behavior: "smooth", block: "nearest" });
}

clearBtn.addEventListener("click", () => {
    items = [];
    fileInput.value = "";
    previewSection.style.display = "none";
});

// ─────────────────────────────────────────────
// Search
// ─────────────────────────────────────────────
searchBtn.addEventListener("click", startSearch);
quickSearchBtn.addEventListener("click", startQuickSearch);

function startQuickSearch() {
    const company = manualCompany.value.trim();
    const person = manualPerson.value.trim();
    if (!company) {
        alert("Please enter a company name.");
        return;
    }
    items = [{ company, person }];
    startSearch();
}

async function startSearch() {
    if (!items.length) return;
    searchBtn.disabled = true;

    // Reset results
    allResults = [];
    resultsBody.innerHTML = "";
    resultsCount.textContent = "";

    // Show progress
    progressSection.style.display = "block";
    progressSection.scrollIntoView({ behavior: "smooth" });
    updateProgress(0, items.length, 0);

    try {
        const resp = await fetch("/search", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ items }),
        });
        const data = await resp.json();
        if (data.error) { alert(data.error); searchBtn.disabled = false; return; }
        jobId = data.job_id;
        pollInterval = setInterval(pollStatus, 3000);
    } catch (err) {
        alert("Could not connect to the backend. Is server.py running?\n\nRun: python server.py");
        searchBtn.disabled = false;
    }
}

async function pollStatus() {
    if (!jobId) return;
    try {
        const resp = await fetch(`/status/${jobId}`);
        const data = await resp.json();

        allResults = data.results || [];
        updateProgress(data.progress, data.total, data.completed);
        renderResults(allResults);

        if (data.status === "done" || data.status === "cancelled") {
            clearInterval(pollInterval);
            pollInterval = null;
            searchBtn.disabled = false;
            progressSection.style.display = "none";
            resultsSection.style.display = "block";
            resultsSection.scrollIntoView({ behavior: "smooth" });
        }
    } catch (err) {
        console.error("Poll error:", err);
    }
}

cancelBtn.addEventListener("click", async () => {
    if (jobId) {
        await fetch(`/cancel/${jobId}`, { method: "POST" });
        clearInterval(pollInterval);
        pollInterval = null;
        searchBtn.disabled = false;
        progressSection.style.display = "none";
        if (allResults.length) {
            resultsSection.style.display = "block";
        }
    }
});

function updateProgress(pct, total, completed) {
    progressBar.style.width = pct + "%";
    progressPct.textContent = pct + "%";
    progressText.textContent = `${completed} / ${total} completed`;
}

// ─────────────────────────────────────────────
// Results rendering
// ─────────────────────────────────────────────
function renderResults(results) {
    resultsSection.style.display = "block";
    const filter = filterInput.value.toLowerCase();
    const filtered = filter
        ? results.filter(r =>
            (r.company && r.company.toLowerCase().includes(filter)) ||
            (r.person && r.person.toLowerCase().includes(filter))
        )
        : results;

    resultsCount.textContent = `(${results.length})`;
    resultsBody.innerHTML = filtered.map((r, i) => rowHtml(i + 1, r)).join("");
}

function rowHtml(idx, r) {
    const website = r.website ? linkHtml(r.website, shorten(r.website, 30)) : emptyCell();
    const linkedin = r.linkedin ? linkHtml(r.linkedin, "LinkedIn") : emptyCell();
    const email = r.hr_email ? `<a class="cell-link" href="mailto:${escHtml(r.hr_email)}">${escHtml(r.hr_email)}</a>` : emptyCell();
    const phone = r.hr_phone ? escHtml(r.hr_phone) : emptyCell();

    const googleQuery = r.person ? `${r.person} ${r.company}` : r.company;
    const googleLink = `<a class="google-link" href="https://www.google.com/search?q=${encodeURIComponent(googleQuery)}" target="_blank" title="Search on Google">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
            <circle cx="11" cy="11" r="8"/><path d="m21 21-4.35-4.35"/>
        </svg>
    </a>`;

    const status = badgeHtml(r.status);

    return `<tr>
    <td>${idx}</td>
    <td class="company-name">${escHtml(r.company)}</td>
    <td>${escHtml(r.person || "—")}</td>
    <td>${website}</td>
    <td>${linkedin}</td>
    <td>${email}</td>
    <td>${phone}</td>
    <td>${googleLink}</td>
    <td>${status}</td>
  </tr>`;
}

function linkHtml(url, label) {
    return `<a class="cell-link" href="${escHtml(url)}" target="_blank" rel="noopener">
    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
      <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"/>
      <polyline points="15 3 21 3 21 9"/><line x1="10" y1="14" x2="21" y2="3"/>
    </svg>
    ${escHtml(label)}
  </a>`;
}

function emptyCell() { return `<span class="cell-empty">—</span>`; }

function badgeHtml(status) {
    if (status === "done") return `<span class="badge badge-done">✓ Done</span>`;
    if (status === "error") return `<span class="badge badge-error">✕ Error</span>`;
    return `<span class="badge badge-searching"><span class="spinner"></span> Searching</span>`;
}

function shorten(url, max) {
    try {
        const u = new URL(url);
        const s = u.hostname.replace("www.", "") + u.pathname;
        return s.length > max ? s.slice(0, max) + "…" : s;
    } catch { return url.slice(0, max); }
}

// ─────────────────────────────────────────────
// Filter
// ─────────────────────────────────────────────
filterInput.addEventListener("input", () => renderResults(allResults));

// ─────────────────────────────────────────────
// Export CSV
// ─────────────────────────────────────────────
exportBtn.addEventListener("click", () => {
    if (!allResults.length) return;
    const headers = ["Company", "Person", "Website", "LinkedIn", "HR Email", "HR Phone", "Status"];
    const rows = allResults.map(r => [
        r.company, r.person, r.website, r.linkedin, r.hr_email, r.hr_phone, r.status
    ]);
    const csv = [headers, ...rows]
        .map(row => row.map(cell => `"${(cell || "").replace(/"/g, '""')}"`).join(","))
        .join("\n");

    const blob = new Blob([csv], { type: "text/csv" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `company_research_${new Date().toISOString().slice(0, 10)}.csv`;
    a.click();
    URL.revokeObjectURL(url);
});

// ─────────────────────────────────────────────
// Utility
// ─────────────────────────────────────────────
function escHtml(str) {
    return String(str || "")
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;");
}
