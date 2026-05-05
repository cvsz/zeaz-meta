#!/usr/bin/env bash
set -euo pipefail
: "${CLOUDFLARE_API_TOKEN:?required}"
: "${CLOUDFLARE_ZONE_ID:?required}"
: "${SUBDOMAIN:?required, for example api}"
: "${TARGET:?required CNAME target}"
curl -fsS -X POST "https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_ID}/dns_records" \
  -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
  -H 'Content-Type: application/json' \
  --data "{\"type\":\"CNAME\",\"name\":\"${SUBDOMAIN}.zeaz.dev\",\"content\":\"${TARGET}\",\"proxied\":true,\"ttl\":1}"
