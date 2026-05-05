#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
export MODE="${MODE:-deploy}"
exec "$ROOT/scripts/meta-installer.sh" "$@"
