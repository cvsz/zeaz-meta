# ZeaZDev Repository Consolidation Plan

Status: Approved
Owner: ZeaZDev
Scope: cvsz GitHub repositories

## Objective

Clean, merge, combine, refactor, and optimize the ZeaZDev repository portfolio so the GitHub profile presents a professional product ecosystem instead of scattered experiments, forks, and prototypes.

## Public Flagship Repositories

These repositories should remain public and receive the highest level of polish:

| Repository | Role |
|---|---|
| `zeaz-platform` | Main ecosystem platform and company/product foundation |
| `zeaz-meta` | Architecture, standards, roadmap, governance, and persistent AI context |
| `zeaz-ai-tooling` | AI development tooling, Codex automation, Git workflow control, and OpenAPI generation |
| `zeaz-linebot-suite` | Full-stack LINE Bot platform with automation and AI workflows |
| `zeaz-wallet` | Wallet, swap, and payment product |
| `zeaz-lms` | Learning management system |
| `zeaz-safeguard` | Mobile safety and protection platform |

## Rename Targets

| Current | Target |
|---|---|
| `zwallet` | `zeaz-wallet` |
| `zlms-prod` | `zeaz-lms` |
| `zSafeGuard` | `zeaz-safeguard` |

## Merge Targets

### zeaz-ai-tooling

| Current Repo | Target Path |
|---|---|
| `zcodex` | `tools/codex-installer` |
| `zgitcp` | `apps/git-control-panel` |
| `openapi-builder` | `packages/openapi-builder` |

### zeaz-linebot-suite

| Current Repo | Target Path |
|---|---|
| `zLinebot` | repository root |
| `zLinebot-automos` | `apps/automation` |

## Private / Internal Targets

These repositories should be made private until they have clear scope, documentation, compliance review, and production-readiness:

| Repository | Target Group |
|---|---|
| `zlttbots` | `zeaz-social-automation` |
| `zvath` | `zeaz-social-automation` |
| `ztsaff` | `zeaz-social-automation` |
| `zypto` | `zeaz-trading-lab` |
| `ABTPi18n` | `zeaz-trading-lab` |
| `zGaming` | `zeaz-gaming-lab` |
| `livescan` | `zeaz-gaming-lab` |
| `zspin` | review/private |
| `zcino` | review/private |
| `zveo` | review/private |
| `zttlbots` | review/private |
| `ZeaZDev-Omega` | review; merge or archive |

## Archive Targets

Archive forks, samples, duplicates, and reference-only repositories unless they are actively modified and productized:

- `python-docs-samples`
- `nodejs-docs-samples`
- `golang-samples`
- `adk-python`
- `open-webui`
- `awesome-copilot`
- `tiktokshop-api-client`
- `tiktok-shop-bot`

## Repository Rule

Public repositories must be polished products or reusable tools.
Private repositories are for experiments, automation, trading, gaming, internal tools, and compliance-sensitive work.
Archived repositories are forks, samples, duplicates, and outdated prototypes.

## Required Public Repo Standard

Every public flagship repo should include:

- `README.md`
- `.faf`
- `docs/architecture.md`
- `docs/roadmap.md`
- `SECURITY.md`
- `LICENSE` or explicit proprietary notice
- `.github/workflows/ci.yml` where applicable
- clear scope, status, and maintenance policy

## Migration Commit Convention

Use Conventional Commits:

```bash
chore(repo): consolidate ZeaZDev repository portfolio
chore(repo): migrate project into ZeaZDev product ecosystem
chore(repo): deprecate repository after consolidation
```

## Deprecation Banner

Use this in old repositories after migration:

```md
> Status: Archived / Deprecated
> This repository has been consolidated into the ZeaZDev product ecosystem.
```
