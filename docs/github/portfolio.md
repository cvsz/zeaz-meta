# cvsz GitHub Portfolio

Snapshot date: 2026-08-06

## Scope

This document organizes the public repositories visible through the connected GitHub account. Private repository names and metadata are deliberately excluded because this governance repository is public.

The classifications below are operational starting points. They are not archive, deletion, transfer, or rename decisions. Any destructive or externally visible repository change requires explicit owner approval in a separate issue or pull request.

## Portfolio model

| Lifecycle | Meaning | Required action |
|---|---|---|
| `governance` | Portfolio policy, architecture, evidence, and control plane | Keep current; protect review process |
| `core` | Strategic platform or runtime dependency | Highest validation and compatibility requirements |
| `active` | Product or tool receiving planned development | Maintain owner, roadmap, CI, security, and release metadata |
| `incubating` | Early product, prototype, or unclear product-market boundary | Define outcome and graduation/closure date |
| `reference` | Upstream mirror, fork, sample, or research dependency | Record upstream, divergence, and update policy |
| `maintenance` | Supported but not under active feature development | Security and compatibility fixes only |
| `archive-review` | Duplicate, stale, historical, or superseded candidate | Review evidence; do not archive automatically |

## Strategic control plane

### Governance

- `zeaz-meta` — repository governance, architecture, security, source inventory, and portfolio evidence.

### Core platform

- `z-platform` — primary Z Platform monorepo and integration boundary.
- `zdash` — governed product integrated into `z-platform` through a pinned revision at `apps/zdash`.

### AI coding and agent toolchain

- `zai-coder`
- `aicoder`
- `zcodex`
- `zc`
- `codex-universal`
- `agent-harness-generator`
- `zagents-generator`
- `git-env-manager`
- `GPTxCODEX-CONFIG`
- `zsp-aitool`

These repositories require a consolidation review because their capabilities overlap. The goal is a declared product boundary and migration path, not an immediate merge or archive.

### Product and integration families

- Content and media: `zytgen`, `zveo`, `zsticker`, `aioci-content-summary`.
- Commerce and TikTok: `tiktok-shop-sdk`, `tiktok-shop-bot`, `tiktokshop-php`.
- Messaging automation: `zLinebot`, `zLinebot-automos`, `zlttbots`, `zttlbots`.
- Learning and workspace: `zAcademy`, `learnship`, `my-virtual-office`.
- Wallet and identity-adjacent: `zwallet`, `zypto`, `zvath`.
- Operations and security: `zSafeGuard`, `zgitcp`, `livescan`, `SECGenerator`.

## Reference and upstream review queue

The following repositories look like upstream projects, documentation samples, forks, or research references. Each needs an `UPSTREAM.md` or README section declaring the upstream URL, pinned revision, divergence, license, and update policy:

- `open-webui`
- `vscode`
- `ComfyUI`
- `opencode`
- `one-api`
- `adk-python`
- `python-docs-samples`
- `nodejs-docs-samples`
- `golang-samples`
- `cloud-code-custom-samples-example`
- `awesome-copilot`
- `everything-claude-code`
- `everything-gemini-code`
- `free-claude-code`

## Incubation review queue

These public repositories should receive a one-page product decision: owner, user, problem, relationship to `z-platform`, next milestone, and review date.

- `ZeaZDev-Omega`
- `ABTPi18n`
- `openapi-builder`
- `ztsaff`
- `zGaming`
- `zspin`
- `zvath`
- `zcino`
- `zeye`
- `zkbtrader`
- `OmegaQuant`
- `zAcademy`
- `zsticker`

## Duplicate and legacy review queues

### Traccar family

- `traccar`
- `traccar-web`
- `traccar-web-1`
- `traccar-manager-android`
- `traccar-client-android`
- `traccar-client-ios`
- `traccar-manager-ios`
- `TraccarInstaller`

### ClipBucket family

- `clipbucket`
- `clipbucketv3`

### Ragnarok and game-server family

- `rathena`
- `openkore`
- `NEMO`
- `FluxCP`
- `CORA`

### Historical utilities and imported code

- `THGirlclock`
- `phonegap-plugins`
- `passport-reader`
- `mssql-jdbc`
- `adminer`
- `Mirai-Source-Code`
- `rs-csv-importer`
- `class.truewallet.php`
- `vesta`
- `cars-for-sale-php-website`
- `wowza-server`

These are `archive-review`, not approved archives. Security-sensitive imported code such as `Mirai-Source-Code` requires an explicit research-only purpose, prominent warning, and distribution review.

## Portfolio priorities

1. Keep `zeaz-meta`, `z-platform`, and `zdash` governed and green.
2. Resolve product boundaries across the AI coding/agent repositories.
3. Add upstream declarations to reference repositories.
4. Normalize repository descriptions, topics, README status, default branch, security policy, and ownership.
5. Review duplicate families and create explicit keep/migrate/archive proposals.
6. Apply approved changes one repository at a time with rollback notes.

## Privacy boundary

Private repositories follow the same lifecycle policy but are maintained in a private inventory. Their names, visibility, purpose, and metadata must not be copied into this public document.
