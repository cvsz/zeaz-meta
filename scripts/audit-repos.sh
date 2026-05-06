#!/usr/bin/env bash
set -euo pipefail

ROOT="${1:-$(pwd)/sources}"
mkdir -p "$ROOT"
repos=(
  cvsz/zgitcp cvsz/zwallet cvsz/ABTPi18n cvsz/zypto cvsz/ZeaZDev-Omega
  cvsz/ztsaff cvsz/zLinebot cvsz/zlttbots cvsz/zttlbots cvsz/zTTato-Platform
  cvsz/zvath cvsz/tiktok-shop-bot cvsz/tiktokshop-api-client cvsz/tiktok-shop-sdk
  cvsz/tiktokshop-php cvsz/zLinebot-automos cvsz/zeaz-platform cvsz/zeapay
  cvsz/zcino cvsz/zGaming cvsz/zSafeGuard cvsz/zspin
  ZeaZDev/vscode ZeaZDev/ZeaClicker ZeaZDev/zeazchain ZeaZDev/ZeaZDev
  ZeaZDev/ZeaZDev-Omega ZeaZDev/zeazdev-repo ZeaZDev/zeaztools ZeaZDev/zlms-prod
)

repo_dir_name() {
  local full="$1" owner="${full%%/*}" repo="${full##*/}"
  if [[ "$owner" == "cvsz" ]]; then
    printf '%s\n' "$repo"
  else
    printf '%s__%s\n' "$owner" "$repo"
  fi
}

for full in "${repos[@]}"; do
  dir="$(repo_dir_name "$full")"
  if [[ -d "$ROOT/$dir/.git" ]]; then
    echo "exists $full"
    continue
  fi
  if command -v gh >/dev/null 2>&1; then
    gh repo clone "$full" "$ROOT/$dir" -- --depth=1 || git clone --depth 1 "https://github.com/$full.git" "$ROOT/$dir" || echo "blocked $full"
  else
    git clone --depth 1 "https://github.com/$full.git" "$ROOT/$dir" || echo "blocked $full"
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
