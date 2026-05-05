# Security Risk Report

## Attack surface map

| Surface | Exposed components | Primary controls in target state |
|---|---|---|
| Public HTTP APIs | auth, wallet, payment, swap, bot, affiliate, notification, audit | Cloudflare WAF/rate limiting, Gateway JWT validation, strict tenant context, OPA/RBAC, request body limits |
| Provider webhooks | Stripe, PromptPay/bank, LINE, TikTok Shop | Raw-body HMAC/signature verification, timestamp skew checks, Redis replay cache, idempotency table, dead-letter queue |
| Money movement | wallet transfers, payment settlement, swap orders, signer requests | Serializable double-entry ledger, policy risk checks, TSS quorum, immutable audit, per-tenant limits |
| CI/CD and supply chain | GitHub Actions, images, Terraform, ArgoCD | SAST/dependency/container scans, OIDC deploy identity, signed images, protected branches, Argo sync windows |
| Infrastructure mutation | Terraform/Packer/Kubernetes/Cloudflare | GitOps-only changes, remote state locks, least-privilege tokens, no host installers |
| Secrets and keys | OAuth credentials, provider API keys, signing shares | Vault dynamic/static secrets with rotation, PKCS#11/HSM custody, no `.env` production secrets |
| Background workers | affiliate outreach, bot sends, swaps, Drive/content sync | NATS durable consumers, leases, budgets, kill switches, audit on each external side effect |
| Observability data | logs, traces, metrics, audit exports | PII redaction, tenant labels, immutable object-lock audit archive, SIEM correlation rules |

## Risk matrix

| Severity | Risk | Evidence from intake | Exploit scenario | Mitigation |
|---|---|---|---|---|
| Critical | Unreviewed payment repository (`zeapay`) unavailable | Clone blocked by credentials | Hidden payment code contains unauthenticated callback or unsafe settlement path and is imported blindly | Do not import. Rebuild payment-service from audited contracts; require authenticated audit before any code reuse |
| Critical | Host persistence/installers/watchdogs | High script counts in `ztsaff`, `zlttbots`, `zLinebot`, `zLinebot-automos`, `zypto` | Attacker modifies installer/self-heal loop to re-add backdoor after rollback | Quarantine scripts; deploy only signed Terraform/Packer/K8s/ArgoCD; run `scripts/clean-os.sh` on nodes before bootstrap |
| Critical | Fragmented money-movement logic | Wallet/payment/swap appears across `zwallet`, bot repos, Omega, unavailable `zeapay` | Double settlement or cross-tenant transfer bypass through weaker service | Canonical wallet/payment schemas, idempotency keys, serializable ledger, unified RBAC, TSS signer policy |
| High | Duplicated webhook verification | LINE/TikTok/Stripe handlers in multiple repos and SDKs | Forged webhook accepted by one implementation; replay triggers duplicate payout/message | Central Go verifier in `platform/internal/integration`, raw body validation, timestamp and nonce replay cache, provider event table |
| High | Secret leakage and `.env` reliance | Secret-like risk patterns and `.env.example`/generated secret scripts | Static provider token committed or copied into container image | Vault-backed config, secret scanning in CI, no production `.env`, short-lived dynamic secrets |
| High | Unsafe shell execution | Shell-heavy repos with `sudo`, `systemctl`, `curl`, generated scripts | Command injection in deploy/config variable executes as root | Remove runtime shell from services; restrict bootstrap scripts to deterministic, reviewed commands with shellcheck |
| High | Missing service-to-service identity | Legacy services call each other through ad hoc URLs | Compromised pod invokes payment/wallet internals | SPIFFE/SPIRE mTLS, NetworkPolicy deny-by-default, OPA method authorization |
| Medium | Supply-chain drift | Multiple language ecosystems and generated dependency files | Vulnerable transitive package persists in unused runtime | Go-first consolidation, lockfiles, Trivy/govulncheck/npm/composer scans, SBOMs |
| Medium | Autonomous loop runaway | Schedulers/workers/AI/bot loops in automation repos | API spam, TikTok/LINE ban, unexpected cloud spend | NATS leases, per-tenant budgets, kill switch, rate limiter, alert on anomaly |
| Medium | Observability gaps | Local loggers and partial audit stores | Incident lacks traceable settlement/auth history | Central audit-service, OpenTelemetry, Prometheus, ELK/Filebeat, immutable exports |
| Low | SDK examples and tests with secret-shaped strings | SDK risk-pattern hits | False-positive leak noise hides real secrets | Ignore examples from deploy artifacts but scan them in CI with allowlist review |

## Exploit scenarios and required controls

### Forged payment settlement
1. Attacker discovers legacy Stripe/PromptPay webhook route without canonical signature validation.
2. Attacker posts a crafted `payment_succeeded` payload for another tenant.
3. Legacy handler marks invoice paid and emits wallet credit.
4. **Target prevention:** gateway blocks direct legacy route; `payment-service` validates raw signature, provider timestamp, replay key, tenant-bound payment intent, and idempotent settlement transaction before any ledger event.

### Installer persistence after incident rollback
1. Attacker commits or tampers with a shell installer/self-heal script.
2. Host cron/systemd/watchdog re-runs the script after defenders remove the malicious container.
3. Backdoor returns with root privileges.
4. **Target prevention:** host installers are not production deploy inputs; Packer images are immutable; `scripts/clean-os.sh` removes cron/systemd/docker residue before bootstrap; ArgoCD owns application state.

### Cross-tenant bot data leakage
1. A bot webhook endpoint lacks tenant context and reuses a global Redis or database key.
2. Attacker replays another tenant's event ID or command state.
3. Bot sends or reveals messages under wrong tenant.
4. **Target prevention:** tenant context is mandatory at gateway/service middleware; Redis keys include tenant/provider/event; RBAC scopes and audit events are tenant-bound.

### Signing key compromise through fake MPC
1. Application code pretends to use MPC but locally computes deterministic signatures or stores complete private key material.
2. Compromised app pod signs arbitrary transfers.
3. **Target prevention:** `tss-signer` is isolated, fails closed without quorum transcript, requires policy-approved session IDs, and stores shares only in Vault/HSM-backed participant pods.

## Security hardening baseline

- **Vault:** KV v2 for provider secrets, database dynamic credentials, transit for envelope encryption, rotation policies by provider and tenant.
- **HSM:** PKCS#11 library mounted only into signer participants, CloudHSM partition per environment, key labels bound to signer policy IDs.
- **Zero trust:** SPIFFE identities for all service accounts, mTLS everywhere, NetworkPolicy default-deny, no public pod IPs.
- **RBAC:** tenant roles (`owner`, `admin`, `operator`, `viewer`, `bot`, `service`) mapped to explicit scopes; break-glass role requires approval and logs every read/write.
- **Audit:** all policy decisions, webhooks, ledger mutations, signer sessions, deploy syncs, and break-glass actions append immutable audit events.
