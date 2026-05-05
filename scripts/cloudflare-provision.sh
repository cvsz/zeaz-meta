#!/usr/bin/env bash
set -euo pipefail
: "${CLOUDFLARE_API_TOKEN:?CLOUDFLARE_API_TOKEN is required}"
: "${CLOUDFLARE_ZONE_ID:?CLOUDFLARE_ZONE_ID is required}"
: "${TARGET_CNAME:=gateway.zeaz.dev}"
record_name="${1:-*}"
payload=$(jq -nc --arg type CNAME --arg name "$record_name" --arg content "$TARGET_CNAME" '{type:$type,name:$name,content:$content,ttl:1,proxied:true}')
curl --fail-with-body -sS -X POST "https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_ID}/dns_records" \
  -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
  -H "Content-Type: application/json" \
  --data "$payload"
