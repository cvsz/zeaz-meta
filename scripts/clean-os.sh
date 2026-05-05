#!/usr/bin/env bash
set -euo pipefail

if [[ "${EUID}" -ne 0 ]]; then
  echo "clean-os.sh must run as root on disposable/bootstrap hosts" >&2
  exit 1
fi

log() { printf '[clean-os] %s\n' "$*"; }

log "stopping legacy process managers and watchdog units if present"
systemctl stop pm2-root.service pm2.service docker-compose.service zeaz-watchdog.service zlinebot-watchdog.service 2>/dev/null || true
systemctl disable pm2-root.service pm2.service docker-compose.service zeaz-watchdog.service zlinebot-watchdog.service 2>/dev/null || true

log "removing cron persistence entries installed by legacy repositories"
for user in root ubuntu ec2-user; do
  if id "$user" >/dev/null 2>&1; then
    crontab -u "$user" -l 2>/dev/null | sed '/zLinebot/d;/zlttbots/d;/ztsaff/d;/zeaz/d;/watchdog/d;/selfheal/d' | crontab -u "$user" - 2>/dev/null || true
  fi
done
rm -f /etc/cron.d/zeaz* /etc/cron.d/zlinebot* /etc/cron.d/zlttbots* /etc/cron.d/ztsaff* 2>/dev/null || true

log "removing legacy systemd service files"
find /etc/systemd/system -maxdepth 1 -type f \( -iname '*zeaz*.service' -o -iname '*zlinebot*.service' -o -iname '*zlttbots*.service' -o -iname '*ztsaff*.service' -o -iname '*watchdog*.service' \) -print -delete 2>/dev/null || true
systemctl daemon-reload 2>/dev/null || true

log "stopping local docker-compose residue without deleting container runtime packages"
if command -v docker >/dev/null 2>&1; then
  docker ps --format '{{.ID}} {{.Names}}' | awk '/zeaz|zlinebot|zlttbots|ztsaff|zgitea|tiktok/{print $1}' | xargs -r docker stop || true
  docker ps -a --format '{{.ID}} {{.Names}}' | awk '/zeaz|zlinebot|zlttbots|ztsaff|zgitea|tiktok/{print $1}' | xargs -r docker rm || true
fi

log "removing generated local secret files from known legacy paths"
find /opt /srv /var/tmp /tmp -maxdepth 4 -type f \( -name '.env' -o -name '.env.*' -o -name '*secret*' -o -name '*token*' \) \
  | awk '/zeaz|zlinebot|zlttbots|ztsaff|zgitea|tiktok|automos/' \
  | xargs -r shred -u 2>/dev/null || true

log "clean OS baseline complete; deploy through Terraform/Packer/ArgoCD only"
