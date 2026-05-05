#!/usr/bin/env bash
set -euo pipefail
: "${ENVIRONMENT:?dev, staging, or prod}"
: "${ARGOCD_SERVER:?required}"
: "${ARGOCD_AUTH_TOKEN:?required}"

case "$ENVIRONMENT" in dev|staging|prod) ;; *) echo "invalid ENVIRONMENT" >&2; exit 2;; esac

argocd --server "$ARGOCD_SERVER" --auth-token "$ARGOCD_AUTH_TOKEN" app sync "zeaz-${ENVIRONMENT}" --prune --timeout 600
argocd --server "$ARGOCD_SERVER" --auth-token "$ARGOCD_AUTH_TOKEN" app wait "zeaz-${ENVIRONMENT}" --health --operation --timeout 600
