# Production Deployment Guide

## 1. Promotion gates

1. Merge only signed commits that pass Go tests, container builds, dependency review, CodeQL, secret scanning, and image vulnerability scans.
2. Build one immutable image per service: `wallet-service`, `payment-service`, `swap-engine`, `bot-service`, `affiliate-service`, `notification-service`, `auth-service`, `audit-service`, and `tss-signer`.
3. Sign every image with keyless Sigstore and publish provenance before ArgoCD is allowed to sync production.
4. Apply database migrations through a single migration job that uses Vault-issued credentials with a 15-minute lease.

## 2. Infrastructure rollout

```bash
cd platform/deploy/terraform/envs/prod
terraform init
terraform plan -out=tfplan
terraform apply tfplan
```

The production Terraform stack creates isolated primary and failover EKS clusters, optional GKE hybrid capacity, encrypted data planes, Cloudflare wildcard DNS/WAF/rate limits, private control-plane endpoints, regional node pools, KMS keys, Aurora PostgreSQL, and Redis multi-AZ replication.

## 3. Cluster security bootstrap

```bash
kubectl apply -f platform/deploy/spire/server.conf
kubectl apply -k platform/deploy/kubernetes/overlays/prod
vault policy write zeaz-platform platform/deploy/vault/policies.hcl
```

SPIRE issues workload identities under `spiffe://zeaz.dev`, Vault maps Kubernetes service accounts to short-lived secret leases, and Kubernetes NetworkPolicies deny all traffic except mesh, Vault, NATS, data, and observability egress.

## 4. GitOps release

```bash
kubectl apply -f platform/deploy/argocd/app-of-apps.yaml
argocd app sync zeaz-platform
argocd app wait zeaz-platform --health
```

Rollback is performed by reverting the Git commit that changed image tags or manifests and re-syncing ArgoCD. Emergency rollback can pin `targetRevision` to a known healthy commit.

## 5. TSS signer operating model

`tss-signer` runs five pods across at least three zones on the dedicated signer pool. A signing operation is accepted only when tenant scope, nonce freshness, policy ID, chain allow-list, risk decision, FROST/tss-lib aggregate proof, participant SPIFFE transcripts, and quorum approval are present. The coordinator fails closed without a quorum proof; cryptographic participants load shares from Vault and wrap local material with HSM/PKCS#11 handles.

## 6. Failure and recovery procedures

| Failure | Detection | Recovery |
|---|---|---|
| Primary region outage | Cloudflare health checks and ArgoCD cluster health fail. | Promote failover ingress in Cloudflare, unpause failover consumers, verify Postgres recovery point, run smoke tests. |
| Vault sealed | Service secret lease errors and signer denials increase. | Unseal or fail over Vault, revoke partially issued leases, restart affected pods after leases are healthy. |
| NATS JetStream leader loss | Publish latency and consumer lag alerts. | Allow cluster election, verify stream replicas, replay durable consumers from last acknowledged sequence. |
| Webhook replay attack | SIEM `webhook-replay` alert. | Rotate provider secret in Vault, block source at Cloudflare WAF, replay accepted events from audit log to verify no double spend. |
| Signer participant loss | TSS quorum alert. | Drain failed node, reschedule signer pod on signer pool, verify quorum freshness before resuming high-value signing. |

## 7. Load and chaos validation

- Load test tenant-authenticated API flows at the gateway: token exchange, wallet transfer, payment intent, swap order, LINE webhook burst, TikTok webhook burst, and audit read.
- Chaos test AZ loss, signer pod loss, Redis primary failover, Aurora writer promotion, NATS stream replica loss, and Cloudflare origin failover.
- Recovery is accepted only when idempotency prevents duplicate ledger entries, audit hashes remain continuous, and tenant-scoped SLO dashboards return to green.

## 8. Clean bootstrap and no-persistence rule

Before reusing any host that previously ran legacy repositories, execute `scripts/clean-os.sh` as root, then deploy only through Terraform/Packer/ArgoCD. Runtime services must not install cron entries, systemd units, PM2 processes, self-heal loops, or generated `.env` secret files.
