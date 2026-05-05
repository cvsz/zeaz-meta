# ZeaZ Meta Platform

This repository is the enterprise architecture, security, and infrastructure control plane for consolidating the cvsz ZeaZ ecosystem into a unified Go microservices platform.

## Contents

- `docs/repo-audit.md` — repo-by-repo architecture, feature, dependency, runtime, and security audit.
- `docs/unified-architecture.md` — domain model, service boundaries, API surface, diagrams, rollout plan, TSS/Vault/HSM design, validation, and operations guide.
- `platform/` — production-oriented Go service scaffold, container builds, Kubernetes, Terraform, Vault, SPIRE, observability, GitOps, Cloudflare, and CI/CD assets.
- `scripts/bootstrap.sh` — clean local bootstrap and validation entrypoint.

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
