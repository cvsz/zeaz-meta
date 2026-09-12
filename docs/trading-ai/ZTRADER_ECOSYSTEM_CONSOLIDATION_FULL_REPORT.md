# ZeaZ Trading + AI Ecosystem Consolidation — Full Report

Date: 2026-09-12
Owner: ZeaZDev
Status: Approved consolidation baseline
Mode: Non-destructive migration; no repository deletion before parity gates

## Executive decision

The ZeaZ ecosystem currently contains multiple overlapping trading, AI gateway, dashboard, wallet, and agent implementations. The target state is a domain-owned service architecture with one canonical repository per capability. Duplicate implementations become migration sources, then are frozen/deprecated only after parity evidence exists.

Canonical ownership:

| Domain | Canonical repository |
|---|---|
| Trading intelligence / crypto research | `cvsz/zworkforce` |
| Trading execution / paper / risk / reconciliation | `cvsz/zksato` |
| On-chain / wallet / DEX / liquidity | `cvsz/zwallet` |
| AI gateway / provider routing | `cvsz/zaiman` |
| Trading dashboard / observability UI | `cvsz/zdash` |
| LINE operator interface / notifications | `cvsz/zLinebot-automos` |
| Agent templates / MCP | `cvsz/zagents-generator` |
| Production / GitOps destination | `cvsz/z-platform` |
| Ecosystem governance / architecture | `cvsz/zeaz-meta` |
| Commercial agent/skill platform | `cvsz/zomega` |
| Network controller | `cvsz/zos` |

## Canonical runtime architecture

```text
                         zWorkforce
                    Agents / Orchestration
                            |
                            v
                   zTrader Intelligence
                            |
          +-----------------+-----------------+
          |                 |                 |
          v                 v                 v
       zWallet           zksato            Zaiman
    On-chain / DEX    Risk / Execution    AI Gateway
          |                 |                 |
          +--------+--------+                 |
                   v                          |
             PostgreSQL / Redis               |
                   |                          |
                   v                          |
                 zDash <----------------------+
                   |
                   v
           zLinebot-automos

Deployment:
GitHub -> CI -> GHCR -> ArgoCD -> z-platform -> k3s

Governance:
zeaz-meta
```

## Current zTrader implementation

Canonical intelligence implementation:
`cvsz/zworkforce/packages/ztrader/zeaz-intelligence`

Merged baseline includes:
- narrative lifecycle scoring
- low-cap discovery
- whale / smart-money analysis
- rug / contract / liquidity risk
- Opportunity / Risk / Confidence scoring
- trade-plan generation
- portfolio allocation
- deterministic backtesting
- alerts
- watchlists
- SQLite persistence
- SSE events
- Prometheus metrics
- Docker hardening
- Kubernetes manifests
- execution policy that always keeps live execution disabled by default

The intelligence service remains decision-support-first. Live-money execution must not be triggered directly from social/narrative outputs.

## Canonical service boundaries

### zworkforce

Owns:
- zTrader intelligence API
- social/narrative evidence normalization
- discovery and ranking
- whale/smart-money evidence normalization
- opportunity/risk/confidence scoring
- trading-agent orchestration
- consensus and policy outputs

Must not own:
- exchange credentials
- exchange order placement
- private wallet signing
- duplicate production dashboard stack

Next:
1. live data collectors
2. evidence provenance/freshness
3. connector contracts to zksato, zwallet, zaiman
4. agent supervisor integration

### zksato

Owns:
- deterministic execution boundary
- paper trading
- durable order/fill state
- pre-trade risk
- reconciliation
- exchange adapters
- commission/slippage-aware research
- walk-forward/OOS evaluation
- portfolio exposure limits
- kill switches and approvals

zTrader signals must enter zksato as advisory intents. They never bypass deterministic risk.

### zwallet

Owns:
- EVM / Solana / Bitcoin chain interaction
- wallet and transaction data
- DEX/RFQ routing
- liquidity and price-impact evidence
- wallet-flow and smart-money source adapters
- simulation-before-execution

Private keys/MPC shares never move into zTrader.

### zaiman

Owns:
- OpenAI-compatible gateway
- provider/model registry
- provider account routing
- fallback
- quota-aware routing
- SSE/model protocol adaptation
- OpenAI Responses/Codex/Gemini-compatible surfaces
- local-model routing

Production blockers:
- remove tracked credential-like account artifacts
- rotate any real credentials that were ever committed
- real CI/test/lint gates
- fail-closed production authentication
- secret scanning
- image/dependency pinning
- provider-level observability

### zdash

Owns:
- operator dashboard
- intelligence visualization
- paper PnL/risk display
- observability views
- backtest/research visualization

Execution logic already present in zdash should be migrated or converted to API clients; zdash must not remain an independent order execution authority.

### zLinebot-automos

Owns:
- LINE notifications
- operator commands
- alert delivery

Trading logic in this repository must not become a separate execution engine.

## Migration-source repositories

Repositories below remain readable during migration but should not receive new independent trading-platform features:

### cvsz/ABTPi18n
Extract/preserve:
- i18n UI
- strategy metadata
- CCXT patterns
- monitoring patterns
- useful tests
- notification/websocket patterns

Destination:
- execution/research -> zksato
- UI/visualization -> zdash
- intelligence -> zworkforce

### cvsz/zkbtrader
Extract/preserve:
- fail-closed risk tests
- paper trading safety
- audit patterns
- safe defaults

Destination:
- zksato

### cvsz/OmegaQuant
Status:
- research/foundation source

Extract:
- quant research ideas
- strategy/portfolio concepts

Do not create a third execution platform.

### cvsz/zypto
Status:
- legacy/research source with broad generated/variant surface

Extract only:
- validated execution patterns
- infra/compliance patterns
- relevant tests

Do not bulk-merge variants.

### cvsz/zeaz/apps/ztrader
Status:
- duplicate zTrader implementation

Extract:
- unique tested features only

Then freeze after parity.

### cvsz/cmeerp/apps/ztrader
Status:
- duplicate zTrader implementation

Extract:
- TradingAgents integration
- unique strategies/adapters/tests not present in canonical stack

Then freeze after parity.

### cvsz/zeaz-platform
Role:
- legacy migration/research source

New production work should target cvsz/z-platform unless a migration requires temporary changes.

### cvsz/one-api
Decision:
- not canonical AI gateway
- current runtime is a minimal mock/prototype surface compared with zaiman

Action:
1. preserve useful API-contract/robustness tests
2. migrate compatibility tests to zaiman
3. optionally retain a thin compatibility facade temporarily
4. deprecate/archive only after clients are moved

## Lifecycle policy

```text
ACTIVE
  -> FEATURE_FREEZE
  -> INVENTORY
  -> MIGRATION
  -> PARITY_TEST
  -> DEPRECATED
  -> ARCHIVE
```

Destructive deletion is prohibited before:
- destination feature exists
- source/destination behavioral mapping exists
- tests are migrated or explicitly retired
- production references/imports are removed
- CI passes
- rollback path exists
- owner approval for final archive/delete step

## Integration contracts

### zTrader -> zksato

Input must contain:
- signal id
- symbol/market
- opportunity score
- risk score
- confidence score
- narrative state
- whale state
- evidence timestamp
- invalidation conditions
- proposed trade structure

zksato alone determines:
- whether an intent is accepted
- notional size
- order type
- execution permission
- kill-switch state
- reconciliation handling

### zTrader -> zwallet

Read-only evidence first:
- holder/flow evidence
- wallet clusters
- LP/liquidity evidence
- transfer flows
- chain freshness
- simulation evidence

No signing authority is delegated to zTrader.

### zTrader -> zaiman

All LLM-dependent intelligence should route through the canonical AI gateway where practical.

Required metadata:
- task class
- preferred model/provider
- fallback class
- latency budget
- token budget
- trace/request id
- data-sensitivity policy

### Services -> zdash

Use APIs/events/metrics, not duplicated domain logic.

Recommended:
- SSE for live intelligence events
- Prometheus for service telemetry
- authenticated REST for operational details

## Security findings and required remediation

### zaiman P0
A tracked account-data file contains account identifiers and credential-like fields. Treat any real credential previously committed as exposed.

Required:
- rotate affected credentials
- remove runtime account data from Git
- add deny patterns to .gitignore
- secret scanning in CI
- review Git history where real secrets existed
- encrypted runtime secret storage
- production API-key requirement
- secure cookies
- default-secret startup rejection

### General
- no exchange withdrawal permission
- no private wallet keys in trading/intelligence repos
- live trading remains disabled by default
- production mutations require explicit policy/approval
- pinned images/dependencies for production
- audited request IDs across service boundaries

## Delivery phases

### Phase A — Governance and freeze
- publish this report
- publish machine-readable ownership
- create migration epics
- stop new duplicate platform features

### Phase B — P0 canonical integrations
1. zworkforce <-> zksato
2. zworkforce <-> zwallet
3. zworkforce <-> zaiman

### Phase C — Evidence plane
- social collectors
- on-chain collectors
- provenance/freshness
- durable historical datasets

### Phase D — Paper + forward validation
- advisory signal -> zksato paper intent
- fills/slippage/fees
- walk-forward/OOS
- forward-test evidence
- signal calibration

### Phase E — UI/alerts
- zdash as canonical dashboard
- LINE alert/operator integration
- provider and trading observability

### Phase F — production platform
- z-platform manifests
- GHCR
- ArgoCD
- k3s
- secrets
- metrics/logs/traces
- backup/restore evidence

### Phase G — legacy retirement
For each migration source:
- parity matrix
- final test evidence
- references removed
- deprecation notice
- archive approval

## Acceptance criteria

The consolidation is complete only when:

1. No production trade can bypass zksato deterministic risk.
2. zworkforce is the only canonical intelligence/agent decision layer.
3. zwallet is the on-chain/wallet boundary.
4. zaiman is the canonical provider gateway.
5. zdash reads canonical APIs/events instead of maintaining a second execution authority.
6. zLinebot is an operator/notification surface only.
7. duplicate zTrader implementations are feature-frozen and parity-migrated.
8. legacy repos have explicit lifecycle status.
9. production deploys from z-platform through GitOps.
10. governance ownership is machine-readable in zeaz-meta.
11. all P0 security findings are remediated.
12. live trading remains gated until backtest + OOS + paper + forward evidence exists.

## Final policy

The target architecture favors consolidation over additional repositories. New functionality must be implemented in the canonical owner for its domain or exposed through a versioned integration contract.
