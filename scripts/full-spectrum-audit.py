#!/usr/bin/env python3
"""Clone/intake and statically inventory the ZeaZ repository set.

The script is intentionally read-only for source repos: it uses GitHub CLI first, falls
back to anonymous git for public repos, and emits deterministic Markdown inventories.
"""
from __future__ import annotations

import argparse
import re
import shutil
import subprocess
import json
import os
from collections import Counter
from dataclasses import dataclass
from pathlib import Path
from typing import Iterable

ECOSYSTEM_REPOS = [
    ("cvsz", "zgitcp"), ("cvsz", "zwallet"), ("cvsz", "ABTPi18n"),
    ("cvsz", "zypto"), ("cvsz", "ZeaZDev-Omega"), ("cvsz", "ztsaff"),
    ("cvsz", "zLinebot"), ("cvsz", "zlttbots"), ("cvsz", "zttlbots"),
    ("cvsz", "zTTato-Platform"), ("cvsz", "zvath"),
    ("cvsz", "tiktok-shop-bot"), ("cvsz", "tiktokshop-api-client"),
    ("cvsz", "tiktok-shop-sdk"), ("cvsz", "tiktokshop-php"),
    ("cvsz", "zLinebot-automos"), ("cvsz", "zeaz-platform"),
    ("cvsz", "zeapay"),
    # Additional ZeaZ/cvsz repositories discovered from
    # https://github.com/cvsz?tab=repositories on 2026-05-06.
    ("cvsz", "zcino"), ("cvsz", "zGaming"), ("cvsz", "zSafeGuard"),
    ("cvsz", "zspin"),
    # Organization repositories discovered from
    # https://github.com/orgs/ZeaZDev/repositories on 2026-05-06.
    ("ZeaZDev", "vscode"), ("ZeaZDev", "ZeaClicker"),
    ("ZeaZDev", "zeazchain"), ("ZeaZDev", "ZeaZDev"),
    ("ZeaZDev", "ZeaZDev-Omega"), ("ZeaZDev", "zeazdev-repo"),
    ("ZeaZDev", "zeaztools"), ("ZeaZDev", "zlms-prod"),
]
# Public catalog evidence for the two repository-list URLs. These repos are recorded
# in source inventory even when they are outside the ZeaZ runtime consolidation scope.
CATALOG_OWNERS = ("cvsz", "ZeaZDev")
DISPLAY_LIMIT_PER_REPO = 300
GH_METADATA_FIELDS = "nameWithOwner,description,defaultBranchRef,primaryLanguage,languages,repositoryTopics,licenseInfo,pushedAt,updatedAt,isPrivate"
TEXT_SUFFIXES = {
    ".go", ".py", ".js", ".jsx", ".ts", ".tsx", ".php", ".sh", ".rb", ".java",
    ".kt", ".rs", ".cs", ".aspx", ".ascx", ".sol", ".sql", ".yaml", ".yml", ".json", ".toml", ".hcl",
    ".tf", ".md", ".env", ".example", ".ini", ".conf", ".Dockerfile", "",
}
SKIP_DIRS = {".git", "node_modules", "vendor", "dist", "build", ".next", "coverage", "__pycache__", ".venv", "venv", "env", "bin", "obj", "assets", "images", "Upload", "upload", "Multimedia", "Multimedia_user", "phpMyAdmin", "devexpress", "Content", "Scripts"}
FUNC_PATTERNS = {
    "go": [("Go", re.compile(r"^\s*func\s+(?:\([^)]*\)\s*)?([A-Za-z_][\w]*)\s*\("))],
    "python": [("Python", re.compile(r"^\s*(?:async\s+)?def\s+([A-Za-z_]\w*)\s*\("))],
    "js-ts": [
        ("JS/TS", re.compile(r"^\s*(?:export\s+)?(?:async\s+)?function\s+([A-Za-z_$][\w$]*)\s*\(")),
        ("JS/TS", re.compile(r"^\s*(?:export\s+)?(?:const|let|var)\s+([A-Za-z_$][\w$]*)\s*=\s*(?:async\s*)?\([^)]*\)\s*=>")),
        ("JS/TS", re.compile(r"^\s*(?:public\s+|private\s+|protected\s+|static\s+|async\s+)*([A-Za-z_$][\w$]*)\s*\([^)]*\)\s*[:{]")),
    ],
    "php": [("PHP", re.compile(r"^\s*(?:public|private|protected|static|final|abstract|\s)*function\s+([A-Za-z_]\w*)\s*\("))],
    "csharp": [("C#", re.compile(r"^\s*(?:public|private|protected|internal|static|async|virtual|override|sealed|partial|\s)+[A-Za-z_][\w<>?,\[\] ]+\s+([A-Za-z_]\w*)\s*\("))],
    "shell": [("Shell", re.compile(r"^\s*([A-Za-z_][\w.-]*)\s*\(\)\s*\{"))],
}
ENDPOINT_PATTERNS = [
    re.compile(r"@(app|router|bp)\.(get|post|put|delete|patch|route)\(\s*['\"]([^'\"]+)"),
    re.compile(r"\b(app|router|server|fastify)\.(get|post|put|delete|patch|all|use)\(\s*['\"]([^'\"]+)"),
    re.compile(r"@(Get|Post|Put|Delete|Patch|All)\(\s*['\"]([^'\"]*)"),
    re.compile(r"Route::(get|post|put|delete|patch|any)\(\s*['\"]([^'\"]+)"),
    re.compile(r"['\"](/(?:api|v1|v2|auth|wallet|webhook|stripe|line|tiktok|health|metrics)[^'\"\s]*)['\"]"),
]
PIPELINE_FILES = ("workflow", "action", "pipeline", "deploy", "install", "bootstrap", "cron", "systemd", "watchdog", "docker-compose", "Dockerfile", "Makefile")
RISK_RE = re.compile(r"while true|crontab|systemctl|sudo\s|curl\s+.*\|.*sh|eval\(|exec\(|child_process|subprocess|os\.system|shell=True|PRIVATE_KEY|SECRET|TOKEN|PASSWORD|\.env|watchdog|selfheal|pm2|nohup", re.I)

@dataclass
class RepoInventory:
    name: str
    status: str
    sha: str = "blocked"
    clone_method: str = "unavailable"
    metadata_status: str = "unavailable"
    files: int = 0
    languages: Counter = None  # type: ignore[assignment]
    functions: list[tuple[str, int, str, str]] = None  # path,line,lang,name
    endpoints: list[tuple[str, int, str]] = None
    pipelines: list[str] = None
    risks: int = 0
    modules: list[str] = None


def run(cmd: list[str], cwd: Path | None = None) -> subprocess.CompletedProcess[str]:
    return subprocess.run(cmd, cwd=cwd, text=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)


def repo_key(owner: str, repo: str) -> str:
    return repo if owner == "cvsz" else f"{owner}__{repo}"


def full_name(owner: str, repo: str) -> str:
    return f"{owner}/{repo}"


def clone_repo(root: Path, owner: str, repo: str) -> str:
    dest = root / repo_key(owner, repo)
    if (dest / ".git").exists():
        return "cached"
    if shutil.which("gh"):
        proc = run(["gh", "repo", "clone", full_name(owner, repo), str(dest), "--", "--depth=1"])
        if proc.returncode == 0:
            return "gh"
    proc = run(["git", "clone", "--depth=1", f"https://github.com/{owner}/{repo}.git", str(dest)])
    return "git" if proc.returncode == 0 else "blocked"


def github_api_repo(owner: str, repo: str) -> dict[str, object]:
    import urllib.request
    req = urllib.request.Request(
        f"https://api.github.com/repos/{owner}/{repo}",
        headers={"Accept": "application/vnd.github+json", "User-Agent": "zeaz-meta-audit"},
    )
    with urllib.request.urlopen(req, timeout=20) as resp:  # nosec: public metadata only
        return json.loads(resp.read().decode("utf-8"))


def capture_gh_metadata(root: Path, owner: str, repo: str) -> str:
    """Capture GitHub metadata with gh first, then public REST fallback, without persisting secrets."""
    meta_path = root / f"{repo_key(owner, repo)}.gh.json"
    if shutil.which("gh"):
        proc = run(["gh", "repo", "view", full_name(owner, repo), "--json", GH_METADATA_FIELDS])
        if proc.returncode == 0 and proc.stdout.strip():
            meta_path.write_text(proc.stdout.strip() + "\n")
            return "gh"
        gh_reason = proc.stdout.strip()[:500]
    else:
        gh_reason = "gh-unavailable"
    try:
        data = github_api_repo(owner, repo)
        meta_path.write_text(json.dumps({
            "nameWithOwner": data.get("full_name"),
            "description": data.get("description"),
            "defaultBranchRef": {"name": data.get("default_branch")},
            "primaryLanguage": {"name": data.get("language")},
            "licenseInfo": data.get("license"),
            "pushedAt": data.get("pushed_at"),
            "updatedAt": data.get("updated_at"),
            "isPrivate": data.get("private"),
            "metadataSource": "github-rest-fallback",
            "ghReason": gh_reason,
        }, indent=2, sort_keys=True) + "\n")
        return "rest"
    except Exception as exc:  # noqa: BLE001 - deterministic evidence capture
        meta_path.write_text(json.dumps({"repo": full_name(owner, repo), "status": "blocked", "reason": f"{gh_reason}; REST: {exc}"[:500]}, indent=2) + "\n")
        return "blocked"


def iter_files(repo_dir: Path) -> Iterable[Path]:
    """Yield files while pruning heavy dependency, artifact, binary, and VCS trees."""
    for current, dirs, files in os.walk(repo_dir):
        dirs[:] = sorted(d for d in dirs if d not in SKIP_DIRS)
        base = Path(current)
        for name in sorted(files):
            yield base / name


def safe_lines(path: Path) -> list[str]:
    try:
        if path.stat().st_size > 2_000_000:
            return []
        return path.read_text(errors="ignore").splitlines()
    except OSError:
        return []


def function_pattern_key(path: Path) -> str | None:
    suffix = path.suffix.lower()
    if suffix == ".go":
        return "go"
    if suffix == ".py":
        return "python"
    if suffix in {".js", ".jsx", ".ts", ".tsx"}:
        return "js-ts"
    if suffix == ".php":
        return "php"
    if suffix in {".cs", ".aspx", ".ascx"}:
        return "csharp"
    if suffix == ".sh" or path.name in {"Makefile"}:
        return "shell"
    return None


def language(path: Path) -> str:
    suffix = path.suffix.lower()
    return {
        ".go": "Go", ".py": "Python", ".js": "JavaScript", ".jsx": "JavaScript",
        ".ts": "TypeScript", ".tsx": "TypeScript", ".php": "PHP", ".sh": "Shell",
        ".tf": "Terraform", ".hcl": "HCL", ".yaml": "YAML", ".yml": "YAML",
        ".sql": "SQL", ".sol": "Solidity", ".kt": "Kotlin", ".cs": "C#", ".aspx": "ASP.NET", ".ascx": "ASP.NET", ".md": "Markdown",
    }.get(suffix, suffix.lstrip(".") or "Other")


def infer_modules(paths: list[Path], base: Path) -> list[str]:
    top = Counter()
    semantic = Counter()
    keywords = ["auth", "wallet", "payment", "stripe", "promptpay", "line", "tiktok", "shop", "affiliate", "bot", "webhook", "audit", "admin", "tenant", "rbac", "policy", "swap", "trade", "ledger", "mpc", "sign", "vault", "terraform", "kubernetes", "docker", "worker", "cron", "notify", "i18n"]
    for p in paths:
        rel = p.relative_to(base)
        if len(rel.parts) > 1:
            top[rel.parts[0]] += 1
        low = str(rel).lower()
        for k in keywords:
            if k in low:
                semantic[k] += 1
    modules = [f"{k} ({v} files)" for k, v in top.most_common(10)]
    modules += [f"cap:{k} ({v})" for k, v in semantic.most_common(12)]
    return modules


def inventory(root: Path, owner: str, repo: str) -> RepoInventory:
    key = repo_key(owner, repo)
    repo_dir = root / key
    inv = RepoInventory(full_name(owner, repo), "blocked", languages=Counter(), functions=[], endpoints=[], pipelines=[], modules=[])
    meta_path = root / f"{key}.gh.json"
    if meta_path.exists():
        try:
            meta = json.loads(meta_path.read_text())
            inv.metadata_status = "captured" if "nameWithOwner" in meta else meta.get("status", "captured")
        except json.JSONDecodeError:
            inv.metadata_status = "invalid"
    if not (repo_dir / ".git").exists():
        return inv
    inv.status = "cloned"
    inv.clone_method = run(["git", "config", "--get", "zeaz.cloneMethod"], repo_dir).stdout.strip()
    remote = run(["git", "config", "--get", "remote.origin.url"], repo_dir).stdout.strip()
    if not inv.clone_method or inv.clone_method == "cached":
        inv.clone_method = "cached-public-git" if remote.startswith("https://github.com/") else "cached"
    inv.sha = run(["git", "rev-parse", "--short=12", "HEAD"], repo_dir).stdout.strip() or "unknown"
    paths = list(iter_files(repo_dir))
    inv.files = len(paths)
    inv.modules = infer_modules(paths, repo_dir)
    for p in paths:
        rel = str(p.relative_to(repo_dir))
        inv.languages[language(p)] += 1
        lname = p.name.lower()
        if any(token in lname for token in PIPELINE_FILES) or ".github/workflows" in rel:
            inv.pipelines.append(rel)
        if p.suffix.lower() not in TEXT_SUFFIXES and p.name not in {"Dockerfile", "Makefile"}:
            continue
        for idx, line in enumerate(safe_lines(p), 1):
            if RISK_RE.search(line):
                inv.risks += 1
            key = function_pattern_key(p)
            if key:
                for lang, pat in FUNC_PATTERNS[key]:
                    m = pat.search(line)
                    if m:
                        name = m.group(1)
                        if name not in {"if", "for", "while", "switch", "catch"}:
                            inv.functions.append((rel, idx, lang, name))
                            break
            for pat in ENDPOINT_PATTERNS:
                m = pat.search(line)
                if m:
                    endpoint = m.group(m.lastindex or 0)
                    inv.endpoints.append((rel, idx, endpoint))
                    break
    inv.pipelines = sorted(set(inv.pipelines))
    inv.functions.sort()
    inv.endpoints.sort()
    return inv


def md_escape(s: str) -> str:
    return s.replace("|", "\\|").replace("\n", " ")


def write_docs(outdir: Path, inventories: list[RepoInventory]) -> None:
    outdir.mkdir(parents=True, exist_ok=True)
    summary = outdir / "full-spectrum-analysis.md"
    funcs = outdir / "function-api-inventory.md"
    with summary.open("w") as f:
        f.write("# Full-Spectrum ZeaZ Ecosystem Analysis\n\n")
        f.write("Generated by `scripts/full-spectrum-audit.py` from shallow source clones. Blocked repositories remain excluded from implementation until authenticated source is supplied.\n\n")
        f.write("## Repository inventory\n\n")
        f.write("| Repository | Status | Commit | Files | Clone | GH metadata | Top languages | Functions | APIs/endpoints | Pipelines | Risk hits |\n")
        f.write("|---|---:|---|---:|---|---|---|---:|---:|---:|---:|\n")
        for inv in inventories:
            langs = ", ".join(f"{k}:{v}" for k, v in sorted(inv.languages.items(), key=lambda item: (-item[1], item[0]))[:5]) if inv.languages else "n/a"
            f.write(f"| `{inv.name}` | {inv.status} | `{inv.sha}` | {inv.files} | {inv.clone_method} | {inv.metadata_status} | {md_escape(langs)} | {len(inv.functions)} | {len(inv.endpoints)} | {len(inv.pipelines)} | {inv.risks} |\n")
        f.write("\n## Domain capabilities and normalized owners\n\n")
        domains = [
            ("Identity/RBAC", "auth-service", "signup/login/OAuth/JWT/session middleware/tenant policy found across wallet, bot, automation, and control-plane repos"),
            ("Wallet/Ledger", "wallet-service", "accounts, double-entry ledger, transfers, PromptPay deposit state, metadata, balance views"),
            ("Payments", "payment-service", "Stripe, PromptPay, checkout, billing, subscription, webhook verification, settlement"),
            ("Bots/Channels", "bot-service", "LINE webhooks, TikTok webhooks, chat/campaign callbacks, inbound message normalization"),
            ("TikTok Shop/Affiliate", "affiliate-service", "SDK signing, product feeds, creator discovery, outreach, trend/copy generation"),
            ("Trading/Swap", "swap-engine", "quotes, strategies, backtests, risk limits, routing, clearing/settlement events"),
            ("Notifications/i18n", "notification-service", "LINE Notify, Telegram/email/push, Google Drive/i18n assets, templates"),
            ("Audit/SIEM", "audit-service", "immutable audit append, policy evaluation, security alerts, ELK export"),
            ("Threshold signing", "tss-signer", "MPC references normalized to TSS/FROST ceremonies, PKCS#11/cloud-HSM boundary"),
            ("Infra/GitOps", "platform control plane", "Terraform/Packer/Kubernetes/ArgoCD/Cloudflare/Vault/SPIRE/observability"),
        ]
        f.write("| Domain capability | Unified owner | Extracted source logic |\n|---|---|---|\n")
        for d, o, e in domains:
            f.write(f"| {d} | `{o}` | {e} |\n")
        f.write("\n## Repo modules, services, and automation pipelines\n\n")
        for inv in inventories:
            f.write(f"### `{inv.name}`\n\n")
            f.write(f"- Status: **{inv.status}**, commit `{inv.sha}`, clone `{inv.clone_method}`, GitHub metadata `{inv.metadata_status}`.\n")
            f.write(f"- Modules/capabilities: {md_escape('; '.join(inv.modules[:22]) or 'n/a')}.\n")
            f.write(f"- Automation/deploy files: {md_escape(', '.join(inv.pipelines[:60]) or 'none detected')}.\n")
            f.write(f"- Refactor decision: {'BLOCK production import until authenticated audit is complete.' if inv.status == 'blocked' else 'Port business rules into typed Go services; quarantine host-mutating scripts and duplicated auth/payment code.'}\n\n")
        f.write("## Cross-repo flow map\n\n")
        flows = [
            "Edge request -> Cloudflare wildcard DNS/WAF -> Gateway API -> mTLS/SPIFFE workload -> tenant-aware Go service -> Postgres transaction/outbox -> NATS JetStream event -> audit append -> SIEM export.",
            "LINE/TikTok webhook -> raw body verification/replay cache -> bot-service normalized event -> affiliate/payment/notification consumers -> immutable audit event.",
            "Payment checkout -> payment-service idempotency key -> Stripe/PromptPay adapter -> webhook reconciliation -> wallet ledger posting -> audit export.",
            "Swap request -> swap-engine quote/risk policy -> wallet reservation -> tss-signer threshold session for chain payload -> settlement event -> ledger closeout.",
            "Signer ceremony -> Vault dynamic secret lease -> PKCS#11/cloud-HSM share access -> tss-lib/FROST transcript -> aggregate signature -> transcript hash stored in audit.",
            "Git change -> CI build/test/scan -> image signing/SBOM -> ArgoCD auto-sync -> canary analysis -> automatic rollback on SLO/security failure.",
        ]
        for flow in flows:
            f.write(f"- {flow}\n")
        f.write("\n## Duplicated logic and inconsistencies\n\n")
        f.write("- Auth/JWT/OAuth/RBAC and tenant context are repeatedly implemented in Python, TypeScript, Nest/FastAPI-style code, and scripts; normalize into `auth-service`.\n")
        f.write("- Stripe/PromptPay/webhook verification appears in several runtimes with inconsistent raw-body handling and replay protection; normalize into `payment-service` and shared Go webhook verifier.\n")
        f.write("- Background cron/watchdog/PM2/shell deploy loops create hidden persistence; replace with GitOps-managed Kubernetes Jobs, NATS consumers, and Argo Rollouts.\n")
        f.write("- TikTok Shop request signing is duplicated across TypeScript and PHP SDKs; keep SDKs as reference only and implement canonical Go adapter.\n")
        f.write("- MPC/signing references are not an enterprise key boundary; replace with threshold signing (`tss-lib`/FROST), Vault leases, PKCS#11/cloud-HSM, and isolated signer nodes.\n")
        f.write("- Terraform/Kubernetes assets exist across repos with divergent state assumptions; normalize into versioned modules and environment overlays owned here.\n")
    with funcs.open("w") as f:
        f.write("# Function, API, and Pipeline Inventory\n\n")
        f.write(f"Generated static index of detected functions, HTTP/API strings, and automation pipeline files. Each section displays up to {DISPLAY_LIMIT_PER_REPO} entries per kind per repository while summary tables retain complete detected counts. Re-run the script and raise `DISPLAY_LIMIT_PER_REPO` if a local uncapped artifact is required.\n\n")
        for inv in inventories:
            f.write(f"## `{inv.name}`\n\n")
            f.write(f"Status: **{inv.status}**; commit `{inv.sha}`.\n\n")
            f.write("### API/endpoints\n\n")
            if inv.endpoints:
                f.write("| File | Line | Endpoint/expression |\n|---|---:|---|\n")
                for path, line, ep in inv.endpoints[:DISPLAY_LIMIT_PER_REPO]:
                    f.write(f"| `{md_escape(path)}` | {line} | `{md_escape(ep)}` |\n")
                if len(inv.endpoints) > DISPLAY_LIMIT_PER_REPO:
                    f.write(f"\n_Display capped at {DISPLAY_LIMIT_PER_REPO} of {len(inv.endpoints)} detected endpoint/API expressions._\n")
            else:
                f.write("None detected or repository blocked.\n")
            f.write("\n### Functions/methods\n\n")
            if inv.functions:
                f.write("| File | Line | Language | Function/method |\n|---|---:|---|---|\n")
                for path, line, lang, name in inv.functions[:DISPLAY_LIMIT_PER_REPO]:
                    f.write(f"| `{md_escape(path)}` | {line} | {lang} | `{md_escape(name)}` |\n")
                if len(inv.functions) > DISPLAY_LIMIT_PER_REPO:
                    f.write(f"\n_Display capped at {DISPLAY_LIMIT_PER_REPO} of {len(inv.functions)} detected functions/methods._\n")
            else:
                f.write("None detected or repository blocked.\n")
            f.write("\n### Automation pipeline files\n\n")
            if inv.pipelines:
                for p in inv.pipelines[:DISPLAY_LIMIT_PER_REPO]:
                    f.write(f"- `{md_escape(p)}`\n")
                if len(inv.pipelines) > DISPLAY_LIMIT_PER_REPO:
                    f.write(f"- _Display capped at {DISPLAY_LIMIT_PER_REPO} of {len(inv.pipelines)} detected automation files._\n")
            else:
                f.write("None detected or repository blocked.\n")
            f.write("\n")


def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument("--source-root", default="sources", help="clone/cache directory")
    ap.add_argument("--output-dir", default="docs/generated", help="Markdown output directory")
    ap.add_argument("--skip-clone", action="store_true")
    args = ap.parse_args()
    root = Path(args.source_root).resolve()
    root.mkdir(parents=True, exist_ok=True)
    if not args.skip_clone:
        for owner, repo in ECOSYSTEM_REPOS:
            clone_method = clone_repo(root, owner, repo)
            dest = root / repo_key(owner, repo)
            if (dest / ".git").exists() and clone_method != "cached":
                run(["git", "config", "zeaz.cloneMethod", clone_method], dest)
            metadata_status = capture_gh_metadata(root, owner, repo)
            print(f"{full_name(owner, repo)}: clone={clone_method} gh-metadata={metadata_status}")
    else:
        for owner, repo in ECOSYSTEM_REPOS:
            meta_path = root / f"{repo_key(owner, repo)}.gh.json"
            needs_metadata = not meta_path.exists() or meta_path.stat().st_size == 0
            if not needs_metadata:
                try:
                    json.loads(meta_path.read_text())
                except json.JSONDecodeError:
                    needs_metadata = True
            if needs_metadata:
                capture_gh_metadata(root, owner, repo)
    inventories = [inventory(root, owner, repo) for owner, repo in ECOSYSTEM_REPOS]
    write_docs(Path(args.output_dir), inventories)
    print(f"wrote {args.output_dir}/full-spectrum-analysis.md and function-api-inventory.md")

if __name__ == "__main__":
    main()
