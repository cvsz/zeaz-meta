#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SOURCE_ROOT="${SOURCE_ROOT:-$ROOT/sources}"
OUTPUT_DIR="${OUTPUT_DIR:-$ROOT/docs/generated}"
APPLY="${APPLY:-0}"

log() { printf '[prepare-all-repos] %s\n' "$*"; }
have() { command -v "$1" >/dev/null 2>&1; }

install_hint() {
  local tool="$1" pkg="${2:-$1}"
  if have "$tool"; then
    log "$tool already available: $(command -v "$tool")"
    return
  fi
  if [[ "$APPLY" != "1" ]]; then
    log "dry-run: would install $pkg for missing tool $tool (set APPLY=1 to mutate this disposable builder)"
    return
  fi
  if have apt-get; then
    sudo apt-get update
    sudo apt-get install -y "$pkg"
  elif have brew; then
    brew install "$pkg"
  else
    log "warning: no supported package manager found for $tool; install $pkg manually"
  fi
}

log "checking clone/build toolchain without writing credentials"
install_hint git git
install_hint gh gh
install_hint python3 python3
install_hint go golang-go
install_hint jq jq
install_hint docker docker.io
install_hint kubectl kubectl
install_hint terraform terraform
install_hint packer packer

if have gh; then
  if gh auth status >/dev/null 2>&1; then
    log "gh is authenticated; GitHub metadata and private repositories can be inspected"
  else
    log "gh is installed but unauthenticated; audit falls back to anonymous git for public repositories"
  fi
fi

log "running deterministic repository intake into $SOURCE_ROOT"
python3 "$ROOT/scripts/full-spectrum-audit.py" --source-root "$SOURCE_ROOT" --output-dir "$OUTPUT_DIR"

log "validating refactored Go meta-platform"
(cd "$ROOT/platform" && go mod tidy && go test ./... && go build ./cmd/...)

log "preparation complete; no secrets were persisted by this script"
