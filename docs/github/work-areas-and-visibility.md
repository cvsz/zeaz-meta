# Work Areas and Repository Visibility

Snapshot date: 2026-08-06

## Objective

This document separates the `cvsz` GitHub portfolio into clear work areas, states what each area is responsible for, and defines whether repositories should normally be public or private.

A visibility recommendation is not authorization to change visibility. Every repository visibility change requires a repository-specific review and explicit owner approval.

## Visibility classes

| Class | Meaning | Decision rule |
|---|---|---|
| `public-by-design` | Intended to be reusable, inspectable, or community-facing | Keep public when license, security, and documentation are complete |
| `public-after-review` | Can be public, but only after secrets, internal topology, licensing, and proprietary assets are cleared | Public is conditional, not assumed |
| `private-by-default` | Commercial product, internal workflow, account automation, or unfinished product | Keep private until a release decision explicitly opens it |
| `private-required` | Contains or is expected to contain production infrastructure, payment/authentication logic, customer/account operations, confidential business rules, or sensitive datasets | Must remain private unless redesigned and independently security-reviewed |
| `archive-or-quarantine-review` | Historical, duplicated, unsupported, unsafe, or unclear-purpose repository | Freeze changes; review for archive, private quarantine, or deletion separately |
| `reference-public` | Upstream fork, mirror, sample, or research dependency | May remain public when upstream, license, revision, and local divergence are declared |

## Workstream 1 — Governance and portfolio control

### Purpose

- maintain repository inventory and lifecycle status
- define naming, ownership, security, CI, and archival policies
- track migration and consolidation decisions
- prevent public leakage of private repository metadata

### Source of truth

- `zeaz-meta`

### Default visibility

`public-by-design`

Governance policies should be public and auditable, but private repository names, production topology, secrets, customer data, and confidential business decisions must remain in a separate private control plane.

### Required outputs

- public portfolio catalog
- private inventory stored privately
- repository lifecycle policy
- visibility review queue
- cleanup and archive proposals
- decision records and evidence

## Workstream 2 — Core platform and shared runtime

### Purpose

- provide common identity, gateway, agent, workflow, observability, audit, deployment, and integration contracts
- prevent duplicated platform services across product repositories
- define stable extension boundaries for applications

### Repositories

- `z-platform`
- `zdash`

### Default visibility

`public-after-review`

The platform may be public when it contains reusable architecture and no production secrets, internal-only endpoints, customer configuration, proprietary datasets, restricted model artifacts, or live deployment topology.

### Required separation

Public repository:

- reusable platform source
- public API contracts
- local development environment
- example configuration
- architecture and security model

Private repository or private configuration store:

- production Terraform state and inventories
- DNS and network topology
- credentials and secret references that expose account structure
- customer/tenant configuration
- incident records and private audit data
- commercial deployment overlays

## Workstream 3 — AI coding, agent, and developer tooling

### Purpose

- CLI and agent runtimes
- provider adapters and protocol compatibility
- repository automation and development environment management
- reusable agent harnesses and configuration tooling

### Repositories

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

### Default visibility

`public-after-review`

These should be public when they are generic developer tools and all provider credentials remain server-side or user-supplied at runtime. Commercial routing policy, private prompts, customer telemetry, internal model allowlists, and production provider configuration belong in private repositories or secret stores.

### Portfolio action

- create a capability matrix
- assign one source of truth for CLI, gateway, agent runtime, provider registry, config generator, and installer
- convert overlapping repositories into packages, adapters, references, or migration sources
- do not archive until unique functionality and history are preserved

## Workstream 4 — Reusable SDKs, libraries, schemas, and examples

### Purpose

- publish reusable integration code
- provide stable API contracts and typed clients
- provide examples that contain no production data

### Public candidates

- `tiktok-shop-sdk`
- `tiktokshop-php`
- `ABTPi18n`
- `openapi-builder`
- public portions of `SECGenerator`

### Default visibility

`public-by-design`

### Public release requirements

- clear license and third-party attribution
- versioning and changelog
- no embedded app credentials, merchant IDs, account IDs, callback secrets, or production URLs
- examples use placeholders and synthetic data
- maintained API version and deprecation policy

## Workstream 5 — Product applications and commercial SaaS

### Purpose

- deliver end-user products
- contain business workflows, monetization, product analytics, UI, and product-specific orchestration

### Current product candidates

- `zytgen`
- `zveo`
- `zAcademy`
- `learnship`
- `my-virtual-office`
- `zsticker`
- `zGaming`
- `OmegaQuant`
- `zkbtrader`
- `zcino`
- `zeye`
- `ZeaZDev-Omega`

### Default visibility

`private-by-default`

A product repository becomes public only when the owner intentionally chooses an open-source business model and completes a licensing, security, dependency, brand, and proprietary-asset review.

### Recommended split

Public repository when appropriate:

- SDKs and extension contracts
- example plugins
- public documentation
- self-hosted community edition

Private repository:

- commercial edition
- billing and entitlement logic
- internal prompts and evaluation datasets
- growth experiments and analytics
- production deployment overlays
- customer-specific integrations

## Workstream 6 — Account automation, bots, and social integrations

### Purpose

- operate connected social, messaging, marketplace, and merchant accounts
- execute scheduled or event-driven automation
- process webhooks, sessions, tokens, and account-specific rules

### Repositories requiring private review

- `tiktok-shop-bot`
- `zLinebot`
- `zLinebot-automos`
- `zlttbots`
- `zttlbots`

### Default visibility

`private-required`

Bot frameworks may be extracted into public libraries, but deployed automation logic and account operations should remain private because they commonly expose webhook structure, operational rules, anti-abuse assumptions, tenant IDs, channel IDs, session handling, or commercial workflows.

### Required controls

- least-privilege tokens
- webhook signature validation and replay protection
- account-level audit trail
- environment isolation
- no credentials in repository history
- explicit rate-limit and failure policy

## Workstream 7 — Payments, wallets, identity, and sensitive data

### Purpose

- payment processing
- wallet, ledger, settlement, and financial workflows
- authentication and identity-adjacent functionality
- document or personal-data processing

### Repositories requiring private or quarantine review

- `zwallet`
- `zypto`
- `zvath`
- `class.truewallet.php`
- `passport-reader`

### Default visibility

`private-required`

Reusable cryptographic libraries or open protocols may be public, but product payment logic, merchant integrations, signing policy, fraud controls, settlement rules, identity processing, and production schemas must remain private unless subjected to independent security and compliance review.

### Immediate review priorities

1. secret and credential history scan
2. personal-data and sample-document scan
3. payment provider terms and licensing review
4. threat model and key-management boundary
5. decision to retain, rewrite, quarantine, or archive

## Workstream 8 — Infrastructure, deployment, and production operations

### Purpose

- deploy and operate production systems
- manage cloud, DNS, networking, observability, storage, backup, and recovery
- define environment-specific configuration

### Repositories requiring private review

- `zlms-prod`
- `wowza-server`
- any repository whose primary purpose is production deployment or server configuration

### Default visibility

`private-required`

Public infrastructure modules are acceptable only when they are generic, sanitized, and contain no account identifiers, IP addresses, DNS zones, bucket names, tenant topology, state files, private images, backup paths, or privileged operational procedures.

### Recommended split

Public:

- reusable Terraform/OpenTofu modules
- generic Helm charts
- Dockerfiles and local development compose files
- security hardening guidance

Private:

- environment overlays
- state and inventories
- production values
- incident and recovery data
- vendor account configuration

## Workstream 9 — Upstream references, historical code, and archive review

### Purpose

- preserve upstream research, experiments, legacy products, or migration evidence
- separate maintained products from code retained only for reference

### Reference-public candidates

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

These may remain public if upstream attribution, license, revision, divergence, and update policy are explicit.

### Archive-or-quarantine review candidates

- Traccar family, including duplicate `traccar-web` variants
- ClipBucket family
- Ragnarok/game-server family
- `THGirlclock`
- `phonegap-plugins`
- `mssql-jdbc`
- `adminer`
- `rs-csv-importer`
- `vesta`
- `cars-for-sale-php-website`
- `Mirai-Source-Code`

### Special security decision

`Mirai-Source-Code` should not be treated as a normal public product. It requires an explicit legitimate-research justification, safety notice, license review, secret/binary scan, and owner decision between private quarantine, archival, or deletion.

## Initial visibility recommendation

### Keep public

- `zeaz-meta`
- reusable SDKs and libraries after validation
- reference repositories with complete upstream attribution
- generic developer tooling after secret and licensing review

### Public, but conditional

- `z-platform`
- `zdash`
- AI coding and agent repositories
- defensive security tools

The condition is removal of production topology, credentials, customer data, proprietary assets, and confidential operational policy.

### Move to private review queue

- product applications that are not intentionally open source
- account automation and bot repositories
- wallet, payment, identity, and personal-data repositories
- production infrastructure and deployment repositories
- repositories containing operational merchant or social-account logic

### Archive or quarantine review

- duplicates
- unsupported historical projects
- old imported upstream code with unclear local value
- security-sensitive research code
- repositories with unclear licensing or purpose

## Visibility decision checklist

A repository may be public only when every answer below is acceptable:

- Is the license clear for all source and bundled assets?
- Is the repository free of credentials and sensitive history?
- Are production endpoints, topology, account IDs, and tenant data removed?
- Is customer or personal data absent?
- Are proprietary prompts, datasets, models, and commercial rules intentionally disclosed?
- Are copied workflows safe with public-fork execution?
- Is vulnerability reporting documented?
- Is the repository purpose and support status clear?
- Can the owner maintain public issues and dependency/security updates?

A single unacceptable answer places the repository in `private-by-default`, `private-required`, or `archive-or-quarantine-review` until remediated.

## Execution order

1. protect governance and core platform repositories
2. move production infrastructure and high-risk automation into private review first
3. review payment, wallet, identity, and document-processing repositories
4. separate generic SDKs from deployed product logic
5. define product source-of-truth repositories
6. document upstream/reference repositories
7. review duplicates and historical repositories
8. apply visibility changes only through owner-approved repository-specific changes
