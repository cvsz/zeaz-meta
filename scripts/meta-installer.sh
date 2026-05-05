#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MODE="${MODE:-validate}"
CLEAN_OS="${CLEAN_OS:-0}"
SOURCE_ROOT="${SOURCE_ROOT:-$ROOT/sources}"
DOMAIN="${DOMAIN:-zeaz.dev}"
CLOUDFLARE_ZONE_ID="${CLOUDFLARE_ZONE_ID:-}"
ARGOCD_SERVER="${ARGOCD_SERVER:-}"

log() { printf '[meta-installer] %s\n' "$*"; }
need() { command -v "$1" >/dev/null 2>&1 || { echo "$1 is required" >&2; exit 1; }; }

usage() {
  cat <<USAGE
Usage: MODE=validate|deploy CLEAN_OS=0|1 DOMAIN=zeaz.dev $0

Idempotent enterprise installer. Default MODE=validate performs no cluster or DNS mutation.
Set MODE=deploy to apply Cloudflare/ArgoCD/Kubernetes actions. Secrets are read only from
environment variables and are never written to disk.
USAGE
}

[[ "${1:-}" == "--help" ]] && { usage; exit 0; }

need git
need go
need python3
if ! command -v gh >/dev/null 2>&1; then
  log "warning: gh not found; scripts/full-spectrum-audit.py will use git fallback for public repos"
fi

if [[ "$CLEAN_OS" == "1" ]]; then
  if [[ "$EUID" -ne 0 ]]; then
    echo "CLEAN_OS=1 requires root" >&2
    exit 1
  fi
  "$ROOT/scripts/clean-os.sh"
fi

log "generating source inventory into docs/generated"
"$ROOT/scripts/full-spectrum-audit.py" --source-root "$SOURCE_ROOT" --output-dir "$ROOT/docs/generated"

log "validating Go platform"
(cd "$ROOT/platform" && go mod tidy && go test ./... && go build ./cmd/...)

if command -v terraform >/dev/null 2>&1; then
  log "validating Terraform syntax"
  terraform -chdir="$ROOT/infra/terraform" fmt -check
else
  log "warning: terraform not found; skipping Terraform syntax validation"
fi

if command -v kubectl >/dev/null 2>&1; then
  log "validating Kubernetes manifests client-side"
  kubectl apply --dry-run=client -f "$ROOT/infra/kubernetes/zeaz-v19.yaml" >/dev/null
else
  log "warning: kubectl not found; skipping Kubernetes manifest validation"
fi

if [[ "$MODE" == "validate" ]]; then
  log "validate mode complete; no DNS, cluster, or host persistence changes were made"
  exit 0
fi

[[ "$MODE" == "deploy" ]] || { echo "invalid MODE=$MODE" >&2; exit 1; }
need kubectl
need argocd
need curl

if [[ -n "$CLOUDFLARE_ZONE_ID" ]]; then
  log "provisioning Cloudflare wildcard and service DNS records for *.$DOMAIN"
  "$ROOT/platform/deploy/cloudflare/provision-zone.sh" "$CLOUDFLARE_ZONE_ID" "$DOMAIN"
else
  log "CLOUDFLARE_ZONE_ID unset; skipping DNS mutation"
fi

log "applying Kubernetes base manifests"
kubectl apply -f "$ROOT/infra/kubernetes/zeaz-v19.yaml"
kubectl apply -k "$ROOT/platform/deploy/kubernetes/base"

if [[ -n "$ARGOCD_SERVER" ]]; then
  log "syncing ArgoCD app-of-apps"
  "$ROOT/scripts/safe-deploy.sh" zeaz-meta platform/deploy/argocd/app-of-apps.yaml
else
  log "ARGOCD_SERVER unset; skipping ArgoCD sync"
fi
