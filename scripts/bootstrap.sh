#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
if ! command -v go >/dev/null 2>&1; then
  echo "go is required" >&2
  exit 1
fi
(
  cd platform
  go mod tidy
  go test ./...
)
echo "Bootstrap complete: Go modules resolved and tests passed."
