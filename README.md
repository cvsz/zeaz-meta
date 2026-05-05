# ZeaZ Meta Platform

This repository is the enterprise architecture, security, and infrastructure control plane for consolidating the cvsz ZeaZ ecosystem into a unified Go microservices platform.

## Contents

- `docs/repo-audit.md` — repo-by-repo architecture, feature, dependency, runtime, and security audit.
- `docs/source-inventory.md` — clone/intake evidence, endpoint hits, automation side effects, and quarantine status for unavailable repositories.
- `docs/unified-architecture.md` — domain model, service boundaries, API surface, diagrams, rollout plan, TSS/Vault/HSM design, validation, and operations guide.
- `docs/security-risk-report.md` — OWASP-level risk matrix, attack surface map, exploit scenarios, and hardening baseline.
- `docs/production-deployment-guide.md` — production promotion, Terraform, GitOps, TSS, rollback, failure recovery, load, and chaos procedures.
- `platform/api/openapi.yaml` — unified tenant-scoped API contract for gateway routing and service ownership.
- `platform/db/migrations/001_core_schema.sql` — tenant, user, wallet, ledger, payment, and immutable audit database schema.
- `platform/` — production-oriented Go service scaffold, signed webhook verification, replay protection, container builds, Kubernetes, Terraform, Vault, SPIRE, observability, GitOps, Cloudflare, and CI/CD assets.
- `scripts/bootstrap.sh` — clean local bootstrap and validation entrypoint.
- `scripts/clean-os.sh` — removes legacy cron/systemd/watchdog/docker residue before deterministic GitOps bootstrap.
- `scripts/safe-deploy.sh` — ArgoCD sync/wait deployment wrapper for controlled promotion.
- `scripts/audit-repos.sh` — repeatable source repository evidence inventory command.

## Target Runtime

```text
[ Cloudflare + API Gateway ]
        ↓ mTLS / JWT / tenant context
[ Go Microservices ]
        ↓ NATS JetStream events
[ Postgres + Redis + Immutable Audit Store ]
        ↓ policy-controlled egress
[ LINE / TikTok / Stripe / Chain RPC / HSM / Vault ]
```

## Bootstrap

```bash
./scripts/bootstrap.sh
```
