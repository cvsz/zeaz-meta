#!/usr/bin/env bash
set -euo pipefail
: "${CLOUDFLARE_API_TOKEN:?required}"
ZONE_ID="${CLOUDFLARE_ZONE_ID:?required}"
DOMAIN="${DOMAIN:-zeaz.dev}"
SUBDOMAIN="${SUBDOMAIN:?required, for example api}"
TARGET="${TARGET:?required CNAME target}"
command -v jq >/dev/null 2>&1 || { echo "jq is required" >&2; exit 1; }

name="${SUBDOMAIN}.${DOMAIN}"
record_id=$(curl -fsS "https://api.cloudflare.com/client/v4/zones/${ZONE_ID}/dns_records?type=CNAME&name=${name}" \
  -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" | jq -r '.result[0].id // empty')
payload=$(jq -nc --arg name "$name" --arg target "$TARGET" '{type:"CNAME",name:$name,content:$target,proxied:true,ttl:1}')
if [[ -n "$record_id" ]]; then
  curl -fsS -X PUT "https://api.cloudflare.com/client/v4/zones/${ZONE_ID}/dns_records/${record_id}" \
    -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" -H 'Content-Type: application/json' --data "$payload"
else
  curl -fsS -X POST "https://api.cloudflare.com/client/v4/zones/${ZONE_ID}/dns_records" \
    -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" -H 'Content-Type: application/json' --data "$payload"
fi
