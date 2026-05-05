# Source Repository Intake and Evidence Inventory

Audit date: 2026-05-05 UTC.

## Intake commands executed

The environment did not contain GitHub CLI (`gh`), so the mandatory `gh repo clone` list was attempted through anonymous HTTPS Git as a deterministic fallback:

```bash
mkdir -p /zeazdev/source-repos
cd /zeazdev/source-repos
for r in zgitcp zwallet ABTPi18n zypto ZeaZDev-Omega ztsaff zLinebot zlttbots zttlbots zTTato-Platform zvath tiktok-shop-bot tiktokshop-api-client tiktok-shop-sdk tiktokshop-php zLinebot-automos zeaz-platform zeapay; do
  git clone --depth 1 https://github.com/cvsz/$r.git "$r"
done
```

`zTTato-Platform` and `zeapay` returned GitHub credential prompts over anonymous HTTPS and must remain quarantined until an authenticated clone is supplied.

## Endpoint and automation evidence by repository

| Repository | Runtime evidence | Endpoint / command evidence | Automation and side-effect evidence | Consolidation decision |
|---|---|---|---|---|
| `zgitcp` | Shell/docs bootstrap repository. | README/install scripts only. | Imperative bootstrap flow. | Remove host mutation; retain as signed GitOps bootstrap instructions only. |
| `zwallet` | TypeScript services, Python admin, Kotlin/mobile, SQL, Go security controller. | Gateway routes include `/v1/wallets`, `/v1/wallets/transfer`, `/v1/deposits/promptpay`, `/v1/withdrawals/promptpay`, `/v1/transactions/index`, `/v1/swaps/orchestrate`. | Reconciliation requirements and infra installers. | Split into wallet-service, payment-service, swap-engine, auth-service, audit-service. |
| `ABTPi18n` | FastAPI backend, Next frontend, Python workers/tools. | `/health`, `/auth/login`, `/exchange/keys`, `/strategies`, `/dashboard/pnl`, portfolio endpoints, PromptPay endpoints. | Celery beat references, Drive sync, secret rotation scripts. | Move auth to auth-service; move Drive and strategy notifications to bot/notification services. |
| `zypto` | Large Python control-plane, generated microservices, Terraform/K8s/security assets, Go TSS experiments. | `/auth/signup`, `/auth/me`, `/auth/token`, `/healthz`, `/create-checkout`, sales CRM routes, admin RBAC/audit/SLO routes. | CronJobs, bootstrap cleanup, recovery scripts, cloud clients, generated service fleet. | Keep policy/zero-trust patterns; collapse generated services into typed Go control-plane services. |
| `ZeaZDev-Omega` | NestJS backend, Solidity contracts, frontend miniapp. | Rewards, fintech card/banking, DeFi quote/stake, World ID, wallet connect, game sessions. | Contract deployment workflow. | Map fintech to payment-service, DeFi to swap-engine, identity to auth-service, contracts to audit-controlled release. |
| `ztsaff` | Node/Python/Terraform Gitea and TikTok affiliate platform. | Auth, dashboard, wallet deposit, transactions, rent plans, scheduler/worker controls. | Shell installers, `.env` trees, cron-like scheduler. | Affiliate flows move to affiliate-service with explicit event workers and Vault secrets. |
| `zLinebot` | TypeScript/Python/SQL/Cloudflare Workers. | LINE webhook, TikTok OAuth callback/webhook, Stripe webhook, metrics, automation/log/inference APIs. | Worker deploy scripts, websocket services, compose stacks. | Standardize webhook validation in bot-service and notification-service. |
| `zlttbots` | Python/JavaScript/shell automation mega repo. | Health APIs, automation tools, memory/vector store endpoints. | Installers, hidden loops, Cloudflare automation scripts. | Keep allowed automation as idempotent NATS jobs; remove host installers/watchdogs. |
| `zttlbots` | Fastify TypeScript bot. | `/healthz`, LINE webhook, TikTok webhook, SSE and WebSocket streams. | Direct bot runtime. | Move to bot-service ingress with tenant rate limits and audit stream. |
| `zvath` | FastAPI/TypeScript AI security SaaS. | `/healthz`, tenant `me`, Stripe webhook, generate, OAuth URL/callback/refresh. | Queue/trend workers and scheduler loops. | Auth/OAuth to auth-service; generation/audit side effects to notification/audit. |
| `tiktok-shop-bot` | Python CLI. | Outreach CLI commands; markdown templates. | Dedupe/rate-limit local persistence. | Replace CLI side effects with affiliate-service campaigns and bot-service delivery workers. |
| `tiktokshop-api-client` | PHP Composer SDK. | TikTok Shop signed request helper. | None; synchronous client. | Reference only; Go integration adapter owns signing with Vault. |
| `tiktok-shop-sdk` | TypeScript SDK. | TikTok Shop modules/examples. | Generated docs/examples. | Reference only; do not deploy SDK runtime. |
| `tiktokshop-php` | PHP SDK. | Client/resource/webhook abstractions. | None; synchronous SDK. | Reference only; centralize webhook verifier in Go. |
| `zLinebot-automos` | Python/Node/TS/Terraform/K8s autonomous trading/bot. | Register/chat/checkout, auth, portfolio, backtest, webhook, smart router, orderbook/arb/analytics. | Autonomous loops, generated secrets script, deploy/install scripts. | Trading becomes swap-engine jobs; chat/webhooks become bot-service; billing becomes payment-service. |
| `zeaz-platform` | Shell bootstrap stub. | Minimal script. | Imperative bootstrap. | Replace with `scripts/bootstrap.sh` and GitOps app-of-apps. |
| `zTTato-Platform` | Unavailable. | Unavailable. | Unavailable. | Quarantine as unknown until authenticated audit completes. |
| `zeapay` | Unavailable. | Unavailable. | Unavailable. | Treat as critical payment blind spot; do not import without authenticated source review. |

## Risk classification

| Risk | Repositories | Required mitigation in this meta refactor |
|---|---|---|
| Host-mutating installers and implicit persistence | zgitcp, ABTPi18n, zypto, ztsaff, zlttbots, zLinebot-automos, zeaz-platform | Replace with declarative Terraform, Kubernetes, ArgoCD, and audited bootstrap only. |
| Money movement without unified idempotency | zwallet, ABTPi18n, ZeaZDev-Omega, ztsaff, zLinebot-automos | Enforce tenant ledger schema, idempotency keys, NATS outbox, reconciliation, and append-only audit. |
| Webhook signature fragmentation | zLinebot, zttlbots, tiktokshop-php, zLinebot-automos, ABTPi18n | Centralize LINE/TikTok/Stripe/PromptPay verification in bot/payment services. |
| Secret sprawl in SDKs/scripts/examples | ABTPi18n, zwallet, ztsaff, tiktok SDK repos, zLinebot-automos | Vault dynamic secrets, sealed CI variables, secret scanning, no local `.env` promotion. |
| Autonomous loops and schedulers | zypto, zvath, zLinebot-automos, zlttbots | Convert to explicit event consumers with leases, DLQs, rate limits, and human approval for high-risk actions. |
