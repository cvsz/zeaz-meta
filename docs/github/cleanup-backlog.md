# GitHub Portfolio Cleanup Backlog

Tracking epic: #16

## Guardrails

- Do not delete, archive, transfer, rename, or change visibility through this backlog.
- Public documents must not reveal private repository names or metadata.
- Apply changes in small reviewed batches with evidence and rollback notes.
- Repositories classified as `archive-review` remain active until explicitly approved otherwise.

## Phase 0 — Establish the control plane

- [x] Select `zeaz-meta` as the governance repository.
- [x] Define lifecycle states and approval boundaries.
- [x] Publish an initial public portfolio map.
- [x] Add a machine-readable public catalog.
- [ ] Add an automated inventory refresh that produces a pull request rather than mutating repositories directly.
- [ ] Add schema validation for `config/github-portfolio.yaml`.

## Phase 1 — Strategic repository baseline

Priority repositories: `zeaz-meta`, `z-platform`, `zdash`, `zai-coder`, `aicoder`, `zcodex`, `zc`, and `zytgen`.

For each repository:

- [ ] confirm lifecycle and owner
- [ ] normalize description and topics
- [ ] confirm default branch and branch protection
- [ ] confirm README status and setup instructions
- [ ] confirm security policy and license
- [ ] confirm CI and release evidence
- [ ] identify duplicate capability and migration dependencies

## Phase 2 — Product-family decisions

### AI coding and agent stack

Repositories: `zai-coder`, `aicoder`, `zcodex`, `zc`, `codex-universal`, `agent-harness-generator`, `zagents-generator`, `git-env-manager`, `GPTxCODEX-CONFIG`, `zsp-aitool`.

- [ ] create a feature and architecture matrix
- [ ] select source-of-truth boundaries
- [ ] document compatibility targets
- [ ] identify packages or services that belong in `z-platform`
- [ ] create migration proposals before any consolidation

### TikTok commerce

Repositories: `tiktok-shop-sdk`, `tiktok-shop-bot`, `tiktokshop-php`.

- [ ] distinguish SDK, bot/application, and legacy PHP client responsibilities
- [ ] document credentials and webhook boundaries
- [ ] select maintained SDK/API contracts

### Messaging automation

Repositories: `zLinebot`, `zLinebot-automos`, `zlttbots`, `zttlbots`.

- [ ] identify duplicates and typo-derived names
- [ ] select one product naming convention
- [ ] document migration paths before rename or archive proposals

### Traccar

Repositories: `traccar`, `traccar-web`, `traccar-web-1`, mobile clients/managers, and `TraccarInstaller`.

- [ ] identify upstream origins and revisions
- [ ] identify unique local patches
- [ ] select reference copies worth retaining
- [ ] prepare explicit archive proposals for duplicates only after patch preservation

### ClipBucket

Repositories: `clipbucket`, `clipbucketv3`.

- [ ] identify source-of-truth version
- [ ] document upstream and local modifications
- [ ] prepare migration or archive-review proposal

## Phase 3 — Reference and fork hygiene

For every `reference` repository:

- [ ] add upstream URL and license
- [ ] record pinned/tracked revision
- [ ] document local patches and divergence
- [ ] disable unreviewed workflows with write permissions or secrets
- [ ] define update cadence
- [ ] clarify whether issues and pull requests are accepted locally

Initial queue:

- `open-webui`
- `vscode`
- `ComfyUI`
- `opencode`
- `one-api`
- `adk-python`
- language documentation samples
- Copilot/Claude/Gemini reference collections

## Phase 4 — Historical and security-sensitive review

- [ ] add a prominent research-only and safety notice where appropriate
- [ ] scan for secrets, personal data, binaries, and unsupported dependencies
- [ ] verify license and redistribution rights
- [ ] identify unique historical value
- [ ] prepare keep/private/archive/delete proposals separately

`Mirai-Source-Code` must receive an explicit security and legitimate-purpose review before any other portfolio action.

## Phase 5 — Issue, pull-request, and branch hygiene

Per active repository:

- [ ] triage open pull requests by ready / blocked / obsolete / superseded
- [ ] resolve or convert stale PRs into tracked issues
- [ ] label issues by type, priority, lifecycle, and product family
- [ ] close duplicates with references
- [ ] identify unmerged branches with unique commits
- [ ] remove branches only after evidence review

## Phase 6 — Approved changes

Each destructive or externally visible change receives a dedicated issue containing:

- repository and current lifecycle
- evidence and dependency analysis
- proposed action
- migration and redirect plan
- rollback or recovery plan
- owner approval record

Only then may a repository be renamed, archived, transferred, made private/public, or deleted.

## Definition of done

- public catalog reflects the current visible portfolio
- private catalog exists in a private control plane without leaking names publicly
- every strategic repository has owner, lifecycle, purpose, CI, security, and release metadata
- reference repositories declare upstream and divergence
- duplicate families have approved source-of-truth decisions
- archive candidates have separate explicit owner approvals
