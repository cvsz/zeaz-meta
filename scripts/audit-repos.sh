#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:-/zeazdev/source-repos}"
for repo in zgitcp zwallet ABTPi18n zypto ZeaZDev-Omega ztsaff zLinebot zlttbots zttlbots zTTato-Platform zvath tiktok-shop-bot tiktokshop-api-client tiktok-shop-sdk tiktokshop-php zLinebot-automos zeaz-platform zeapay; do
  if [[ ! -d "$ROOT/$repo" ]]; then
    printf '%s\tmissing\n' "$repo"
    continue
  fi
  files=$(find "$ROOT/$repo" -type f \( -path '*/.git/*' -o -path '*/node_modules/*' -o -path '*/vendor/*' \) -prune -o -type f -print | wc -l | tr -d ' ')
  manifests=$(find "$ROOT/$repo" -maxdepth 4 -type f \( -name package.json -o -name composer.json -o -name go.mod -o -name requirements.txt -o -name pyproject.toml \) -print | paste -sd ',' -)
  endpoints=$( (rg -n --glob '!node_modules' --glob '!vendor' --glob '!*lock*' '(app\.(get|post|put|delete|patch)|router\.(get|post|put|delete|patch)|@(Get|Post|Put|Delete|Patch)\(|FastAPI\(|APIRouter\(|Route::(get|post|put|delete)|http.HandleFunc|fastify\.(get|post|put|delete|patch))' "$ROOT/$repo" || true) | wc -l | tr -d ' ')
  automation=$( (rg -n --glob '!node_modules' --glob '!vendor' --glob '!*lock*' '(cron|schedule|setInterval|while true|watchdog|pm2|supervisor|nohup|crontab)' "$ROOT/$repo" || true) | wc -l | tr -d ' ')
  printf '%s\tfiles=%s\tmanifests=%s\tendpoint_hits=%s\tautomation_hits=%s\n' "$repo" "$files" "$manifests" "$endpoints" "$automation"
done
