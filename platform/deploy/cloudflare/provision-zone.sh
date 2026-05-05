#!/usr/bin/env bash
set -euo pipefail

ZONE_ID="${1:-${CLOUDFLARE_ZONE_ID:-}}"
DOMAIN="${2:-${DOMAIN:-zeaz.dev}}"
: "${CLOUDFLARE_API_TOKEN:?required}"
: "${ZONE_ID:?required as arg1 or CLOUDFLARE_ZONE_ID}"
: "${PRIMARY_ORIGIN:?required primary ingress hostname}"
: "${FAILOVER_ORIGIN:?required failover ingress hostname}"
command -v jq >/dev/null 2>&1 || { echo "jq is required" >&2; exit 1; }

SERVICES_CSV="${SERVICES_CSV:-api,auth,wallet,payment,swap,bot,affiliate,notification,audit,signer,grafana,prometheus,tracing,argocd,vault}"

api() {
  local method="$1" path="$2" data="${3:-}"
  if [[ -n "$data" ]]; then
    curl -fsS -X "$method" "https://api.cloudflare.com/client/v4${path}" \
      -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
      -H 'Content-Type: application/json' \
      --data "$data"
  else
    curl -fsS -X "$method" "https://api.cloudflare.com/client/v4${path}" \
      -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
      -H 'Content-Type: application/json'
  fi
}

upsert_cname() {
  local name="$1" target="$2"
  local existing payload
  existing=$(api GET "/zones/${ZONE_ID}/dns_records?type=CNAME&name=${name}" | jq -r '.result[0].id // empty')
  payload=$(jq -nc --arg name "$name" --arg target "$target" '{type:"CNAME",name:$name,content:$target,proxied:true,ttl:1}')
  if [[ -n "$existing" ]]; then
    api PUT "/zones/${ZONE_ID}/dns_records/${existing}" "$payload" >/dev/null
  else
    api POST "/zones/${ZONE_ID}/dns_records" "$payload" >/dev/null
  fi
}

upsert_cname "*.${DOMAIN}" "$PRIMARY_ORIGIN"
upsert_cname "failover.${DOMAIN}" "$FAILOVER_ORIGIN"
IFS=',' read -r -a services <<< "$SERVICES_CSV"
for service in "${services[@]}"; do
  service="${service//[[:space:]]/}"
  [[ -n "$service" ]] || continue
  target="$PRIMARY_ORIGIN"
  [[ "$service" == "failover" ]] && target="$FAILOVER_ORIGIN"
  upsert_cname "${service}.${DOMAIN}" "$target"
done

api PATCH "/zones/${ZONE_ID}/settings/ssl" '{"value":"strict"}' >/dev/null
api PATCH "/zones/${ZONE_ID}/settings/always_use_https" '{"value":"on"}' >/dev/null
api PATCH "/zones/${ZONE_ID}/settings/min_tls_version" '{"value":"1.2"}' >/dev/null
api PATCH "/zones/${ZONE_ID}/settings/tls_1_3" '{"value":"on"}' >/dev/null

echo "cloudflare zone provisioned for *.${DOMAIN} and services: ${SERVICES_CSV}"
