#!/usr/bin/env bash
set -euo pipefail
: "${CLOUDFLARE_API_TOKEN:?required}"
: "${CLOUDFLARE_ZONE_ID:?required}"
: "${SUBDOMAIN:?required, for example api}"
: "${TARGET:?required CNAME target}"
name="${SUBDOMAIN}.zeaz.dev"
record_id=$(curl -fsS "https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_ID}/dns_records?type=CNAME&name=${name}" \
  -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" | jq -r '.result[0].id // empty')
payload=$(jq -nc --arg name "$name" --arg target "$TARGET" '{type:"CNAME",name:$name,content:$target,proxied:true,ttl:1}')
if [[ -n "$record_id" ]]; then
  curl -fsS -X PUT "https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_ID}/dns_records/${record_id}" \
    -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" -H 'Content-Type: application/json' --data "$payload"
else
  curl -fsS -X POST "https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_ID}/dns_records" \
    -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" -H 'Content-Type: application/json' --data "$payload"
fi
