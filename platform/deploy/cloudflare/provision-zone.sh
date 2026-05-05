#!/usr/bin/env bash
set -euo pipefail

: "${CLOUDFLARE_API_TOKEN:?required}"
: "${CLOUDFLARE_ZONE_ID:?required}"
: "${PRIMARY_ORIGIN:?required primary ingress hostname}"
: "${FAILOVER_ORIGIN:?required failover ingress hostname}"

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
  local existing
  existing=$(api GET "/zones/${CLOUDFLARE_ZONE_ID}/dns_records?type=CNAME&name=${name}" | jq -r '.result[0].id // empty')
  local payload
  payload=$(jq -nc --arg name "$name" --arg target "$target" '{type:"CNAME",name:$name,content:$target,proxied:true,ttl:1}')
  if [[ -n "$existing" ]]; then
    api PUT "/zones/${CLOUDFLARE_ZONE_ID}/dns_records/${existing}" "$payload" >/dev/null
  else
    api POST "/zones/${CLOUDFLARE_ZONE_ID}/dns_records" "$payload" >/dev/null
  fi
}

upsert_cname "*.zeaz.dev" "$PRIMARY_ORIGIN"
upsert_cname "api.zeaz.dev" "$PRIMARY_ORIGIN"
upsert_cname "failover.zeaz.dev" "$FAILOVER_ORIGIN"

api PATCH "/zones/${CLOUDFLARE_ZONE_ID}/settings/ssl" '{"value":"strict"}' >/dev/null
api PATCH "/zones/${CLOUDFLARE_ZONE_ID}/settings/always_use_https" '{"value":"on"}' >/dev/null
api PATCH "/zones/${CLOUDFLARE_ZONE_ID}/settings/min_tls_version" '{"value":"1.2"}' >/dev/null
api PATCH "/zones/${CLOUDFLARE_ZONE_ID}/settings/tls_1_3" '{"value":"on"}' >/dev/null

echo "cloudflare zone provisioned for *.zeaz.dev"
