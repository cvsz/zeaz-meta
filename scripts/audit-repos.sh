#!/usr/bin/env bash
set -euo pipefail

ROOT="${1:-$(pwd)/sources}"
mkdir -p "$ROOT"
repos=(zgitcp zwallet ABTPi18n zypto ZeaZDev-Omega ztsaff zLinebot zlttbots zttlbots zTTato-Platform zvath tiktok-shop-bot tiktokshop-api-client tiktok-shop-sdk tiktokshop-php zLinebot-automos zeaz-platform zeapay)

for repo in "${repos[@]}"; do
  if [[ -d "$ROOT/$repo/.git" ]]; then
    echo "exists $repo"
    continue
  fi
  if command -v gh >/dev/null 2>&1; then
    gh repo clone "cvsz/$repo" "$ROOT/$repo" || git clone --depth 1 "https://github.com/cvsz/$repo.git" "$ROOT/$repo" || echo "blocked $repo"
  else
    git clone --depth 1 "https://github.com/cvsz/$repo.git" "$ROOT/$repo" || echo "blocked $repo"
  fi
done

printf 'repo|files|endpoint_hits|runtime_scripts|risk_hits\n'
for dir in "$ROOT"/*; do
  [[ -d "$dir/.git" ]] || continue
  repo="${dir##*/}"
  files=$(find "$dir" -path '*/.git' -prune -o -type f -printf . | wc -c)
  endpoint_hits=$(rg -n "(@(app|router)\.(get|post|put|delete|patch)|app\.(get|post|put|delete|patch)\(|router\.(get|post|put|delete|patch)\(|@(Get|Post|Put|Delete|Patch)\(|webhook|/health|/api/|/v1/)" "$dir" -g '!**/.git/**' | wc -l || true)
  runtime_scripts=$(find "$dir" -path '*/.git' -prune -o -type f \( -name '*.sh' -o -name '*cron*' -o -name '*systemd*' -o -name '*watchdog*' -o -name 'Dockerfile' -o -name 'docker-compose*.yml' \) -printf . | wc -c)
  risk_hits=$(rg -n "(while true|crontab|systemctl|sudo |curl .*\|.*sh|eval\(|exec\(|child_process|subprocess|os\.system|shell=True|PRIVATE_KEY|SECRET|TOKEN|PASSWORD|\.env|watchdog|selfheal|pm2|nohup)" "$dir" -g '!**/.git/**' | wc -l || true)
  printf '%s|%s|%s|%s|%s\n' "$repo" "$files" "$endpoint_hits" "$runtime_scripts" "$risk_hits"
done | sort
