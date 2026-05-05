# Unified Production Architecture

## Before-state architecture

```text
External users, LINE, TikTok, Stripe, PromptPay, chain RPCs
  ├─> multiple Node/FastAPI/Nest/PHP/Python apps
  │     ├─ duplicated auth/session/OAuth logic
  │     ├─ duplicated webhook/signature logic
  │     ├─ local Docker/PM2/shell/scheduler orchestration
  │     ├─ ad hoc Redis/Postgres/SQLite/local files
  │     └─ local logs and partial audit stores
  └─> shell installers, watchdogs, self-heal loops, compose stacks, generated YAML
```

This creates unbounded execution paths, uncertain ownership of money movement, inconsistent replay/auth controls, and operational drift.

## After-state architecture

```text
Cloudflare WAF / Bot Management / Global Load Balancer / TLS
  ↓
Gateway API / NGINX Ingress with JWT validation, request IDs, tenant extraction
  ↓ mTLS (SPIFFE/SPIRE), network policies, OPA/RBAC
Go microservices
  ├─ auth-service: tenants, users, sessions, OIDC/OAuth, RBAC
  ├─ wallet-service: accounts, balances, double-entry transfers, signer requests
  ├─ payment-service: Stripe, PromptPay, cards/banks, idempotent webhooks
  ├─ swap-engine: quote/risk/order/backtest/settlement orchestration
  ├─ bot-service: LINE/TikTok webhooks, canonical verifier, replay cache
  ├─ affiliate-service: campaigns, creator discovery, consent/rate budgets
  ├─ notification-service: LINE/email/push/SSE/WebSocket fanout
  ├─ audit-service: append-only audit, SIEM export, break-glass logs
  └─ tss-signer: isolated threshold-signing API, DKG/signing rounds, HSM/Vault custody
  ↓
NATS JetStream for durable events, exactly-once-effect semantics via idempotency table
  ↓
Aurora/Postgres global database, Redis HA cache, immutable object-lock audit archive
  ↓
External integrations through egress allow-lists and per-tenant provider credentials in Vault
```

## Bounded contexts and service contracts

| Service | Synchronous REST/gRPC scope | Async subjects | Storage owner | External dependencies |
|---|---|---|---|---|
| `auth-service` | token exchange, registration, login, OAuth link/callback, roles/scopes | `auth.*`, `security.policy.*` | `tenants`, `users`, `sessions`, `roles`, `permissions` | OIDC providers, Vault |
| `wallet-service` | wallet create/list, transfer request/status, balance read | `wallet.*`, `ledger.*`, `signing.requested` | `wallet_accounts`, `ledger_entries`, `idempotency_keys` | `tss-signer`, Postgres |
| `payment-service` | payment intents, refunds, provider webhooks | `payment.*`, `audit.event.appended` | `payment_intents`, `provider_webhooks`, `idempotency_keys` | Stripe, PromptPay/bank APIs |
| `swap-engine` | quote, order submit/cancel/status, backtest | `swap.*`, `risk.*`, `signing.requested` | `swap_orders`, `risk_limits`, `idempotency_keys` | chain RPC, `tss-signer` |
| `bot-service` | LINE/TikTok webhook ingress, bot command execution | `bot.*`, `notification.*` | `webhook_replay`, command state | LINE, TikTok Shop |
| `affiliate-service` | campaigns, creator lists, message approvals | `affiliate.*`, `notification.*` | `campaigns`, `creator_contacts`, consent ledger | TikTok Shop/creator APIs |
| `notification-service` | templated sends, stream publish | `notification.*`, `content.*` | `notification_jobs`, template assets | LINE, email/SMS/push providers |
| `audit-service` | audit read/export/append | `audit.*`, `security.alert.*` | `audit_events`, object-lock exports | SIEM/ELK |
| `tss-signer` | DKG session, signing request, quorum transcript, aggregate signature verification | `signing.*` | signer metadata only; no raw private keys | Vault transit, PKCS#11/CloudHSM, isolated signers |

## API design

- Public APIs are tenant-scoped REST endpoints in `platform/api/openapi.yaml`.
- Internal service-to-service APIs use gRPC with SPIFFE-authenticated mTLS and explicit per-method authorization.
- Every mutating REST endpoint requires an idempotency key, tenant ID, user/service principal, trace ID, and audit event emission.
- Webhooks validate provider signatures over raw request bodies before JSON parsing; the Go implementation covers LINE base64 HMAC, TikTok canonical-header HMAC, and bounded replay rejection before payload handling.

## Control-flow model

1. Cloudflare blocks malicious/bot/rate-limit traffic and forwards only TLS requests for `*.zeaz.dev`.
2. Gateway validates JWT issuer/audience, injects tenant/user headers, enforces request size/timeouts, and records access logs.
3. Service validates scope through RBAC and tenant context; policy decisions are audited.
4. Mutating command opens a serializable database transaction, claims an idempotency key, writes domain state, writes outbox events, and commits.
5. Outbox dispatcher publishes to NATS JetStream with tenant and trace metadata.
6. Consumers process with durable subscriptions, max-delivery limits, dead-letter subjects, and budget/kill-switch checks.
7. Audit service receives security and business events, stores append-only rows, and exports object-lock batches to SIEM.

## Data-flow model

```text
Webhook/payment/trade command
  -> gateway auth and tenant context
  -> domain service command handler
  -> Postgres serializable transaction + idempotency claim
  -> NATS event publication
  -> audit append + SIEM pipeline
  -> notification or external provider adapter
```

Money movement uses double-entry ledger rows and never updates balances without a matching debit/credit pair. Provider callbacks can only settle an existing intent with a matching idempotency/provider event key.

## Runtime lifecycle model

- **Build:** GitHub Actions tests, vets, scans, builds per-service images, signs artifacts, and updates GitOps manifests.
- **Provision:** Terraform creates multi-region EKS/GKE-ready networks, private clusters, encrypted data planes, Cloudflare DNS/WAF, and state locking.
- **Bake:** Packer creates immutable golden nodes with auditd, SSM/Ops agents, disabled password login, and no application secrets.
- **Deploy:** ArgoCD syncs environment overlays; Argo Rollouts performs canary analysis with Prometheus metrics.
- **Operate:** HPA scales stateless services; NATS consumers scale independently; failover uses Cloudflare load balancing and database replicas.
- **Recover:** roll back GitOps revision, pause rollouts, replay NATS dead-letter events after fix, and restore point-in-time database snapshots if invariants fail.

## TSS signer security boundary

`tss-signer` is not a wallet service and does not expose arbitrary signing. It accepts only policy-approved signing sessions from `wallet-service` or `swap-engine`, verifies tenant/key/policy/quorum metadata, coordinates isolated signers, and returns an aggregate signature plus transcript hash. Private key shares never leave Vault/HSM-backed signer pods.

Signing boundary controls:

- Signer namespace has dedicated nodes, taints, NetworkPolicy deny-by-default, no public ingress, and egress only to Vault/HSM/NATS/audit.
- DKG requires an approved ceremony record, participant SPIFFE IDs, HSM key labels, and quorum policy.
- Signing requires tenant risk approval, replay-resistant nonce, payload digest, policy ID, participant transcript commitments, aggregate signature verification, and audit append.
- Break-glass requires two-person approval and produces immutable SIEM events.

## Multi-region high availability

| Layer | Active-active design | Failover behavior |
|---|---|---|
| Edge | Cloudflare proxied wildcard DNS and load balancing | Health checks remove degraded region; WAF/rate rules remain global |
| Kubernetes | Primary and failover clusters in separate regions | ArgoCD app-of-apps targets both; workloads maintain minimum regional replicas |
| Data | Aurora/Postgres global cluster pattern, Redis regional cache, object-lock audit archives | Writes route to elected primary database writer; read-only degraded mode if quorum is unavailable |
| Events | NATS JetStream per region with mirror/source streams for critical subjects | Consumers pause on split brain; idempotency prevents duplicate side effects after replay |
| Signer | Independent signer quorums with region-aware policy | Signing requires configured regional quorum; no automatic key export between regions |

## Refactor roadmap

1. **Freeze and quarantine:** block shell installers, watchdogs, PM2 start scripts, self-heal loops, and runtime-generated secrets from deploy paths.
2. **Introduce edge/gateway:** route all public domains through Cloudflare and Gateway API; reject direct pod/service exposure.
3. **Centralize identity:** migrate all auth/OAuth/JWT/RBAC into `auth-service` with tenant-scoped claims.
4. **Move ledger first:** migrate wallet tables into the canonical schema and add double-entry invariant tests before payment/swap traffic.
5. **Canonicalize webhooks:** implement raw-body verification and replay caches for Stripe/PromptPay/LINE/TikTok in Go.
6. **Event outbox:** add idempotency and outbox publication to each mutating endpoint.
7. **Replace background loops:** convert cron/watchdog/autonomous loops into NATS consumers with leases, rate budgets, and dead-letter queues.
8. **Signer isolation:** deploy `tss-signer`, Vault, HSM/PKCS#11 integration, and ceremony workflows; remove fake MPC paths.
9. **Observability/SIEM:** enforce trace IDs, RED metrics, audit export, and alert rules before production traffic.
10. **Progressive migration:** shadow read legacy outputs, dual-write audit-only events, canary tenants, then retire legacy repos.
