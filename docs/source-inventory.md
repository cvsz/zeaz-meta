# Repository Intake and Evidence Inventory

Audit date: 2026-05-06 UTC.

## Repository-list expansion

The regeneration incorporated the two requested repository-list pages:

- `https://github.com/cvsz?tab=repositories` — public REST catalog returned 59 repositories: `ABTPi18n`, `adk-python`, `adminer`, `aioci-content-summary`, `awesome-copilot`, `cars-for-sale-php-website`, `class.truewallet.php`, `clipbucket`, `clipbucketv3`, `cloud-code-custom-samples-example`, `codex-universal`, `CORA`, `FluxCP`, `golang-samples`, `livescan`, `Mirai-Source-Code`, `mssql-jdbc`, `NEMO`, `nodejs-docs-samples`, `open-webui`, `openapi-builder`, `openkore`, `passport-reader`, `phonegap-plugins`, `python-docs-samples`, `rathena`, `rs-csv-importer`, `SECGenerator`, `THGirlclock`, `tiktok-shop-bot`, `tiktok-shop-sdk`, `tiktokshop-api-client`, `tiktokshop-php`, `traccar`, `traccar-client-android`, `traccar-client-ios`, `traccar-manager-android`, `traccar-manager-ios`, `traccar-web`, `traccar-web-1`, `TraccarInstaller`, `vesta`, `wowza-server`, `zcino`, `zeaz-meta`, `zeaz-platform`, `ZeaZDev-Omega`, `zGaming`, `zgitcp`, `zLinebot`, `zLinebot-automos`, `zlttbots`, `zSafeGuard`, `zspin`, `ztsaff`, `zttlbots`, `zvath`, `zwallet`, `zypto`.
- `https://github.com/orgs/ZeaZDev/repositories` — public REST catalog returned 8 repositories: `vscode`, `ZeaClicker`, `zeazchain`, `ZeaZDev`, `ZeaZDev-Omega`, `zeazdev-repo`, `zeaztools`, `zlms-prod`.

The implementation scope was expanded from the original 18 cvsz targets to 30 ZeaZ/TikTok/runtime repositories: the original list, the newly discovered cvsz ZeaZ repos (`zcino`, `zGaming`, `zSafeGuard`, `zspin`), and every public `ZeaZDev/*` org repository. Non-ZeaZ forks and sample repositories from the cvsz catalog are recorded above as catalog evidence but are not imported into the production meta architecture unless an owner explicitly promotes them into the ZeaZ runtime boundary.

## Intake commands executed

`gh` was installed and the audit scripts attempt GitHub CLI first for each clone and metadata lookup. This non-interactive container has no GitHub authentication token/session, so `gh repo clone` and `gh repo view` returned the standard authentication prompt. The scripts then used anonymous read-only HTTPS `git clone --depth 1` and public GitHub REST metadata fallback for public repositories.

```bash
apt-get update && apt-get install -y gh
python3 scripts/full-spectrum-audit.py --source-root sources --output-dir docs/generated
python3 scripts/full-spectrum-audit.py --source-root sources --output-dir docs/generated --skip-clone
```

## Clone and static metrics

| Repository | Status | Commit | Files | Clone | GH metadata | Top languages | Functions | APIs/endpoints | Pipelines | Risk hits |
|---|---:|---|---:|---|---|---|---:|---:|---:|---:|
| `cvsz/zgitcp` | cloned | `77075fe9a660` | 2 | git | captured | Markdown:1, Shell:1 | 38 | 0 | 0 | 60 |
| `cvsz/zwallet` | cloned | `8164cb36fd63` | 460 | git | captured | TypeScript:120, YAML:77, Markdown:66, json:60, Python:44 | 419 | 78 | 32 | 495 |
| `cvsz/ABTPi18n` | cloned | `6749c6f6c414` | 225 | git | captured | Python:84, Markdown:73, TypeScript:15, YAML:15, Shell:11 | 411 | 100 | 10 | 778 |
| `cvsz/zypto` | cloned | `ef0fdd5ce580` | 3757 | git | captured | YAML:1498, Python:1247, Other:337, Markdown:285, toml:102 | 1821 | 275 | 59 | 356 |
| `cvsz/ZeaZDev-Omega` | cloned | `d0e4ba7b48f5` | 57 | git | captured | TypeScript:28, json:11, Markdown:8, Solidity:5, Other:1 | 38 | 18 | 4 | 235 |
| `cvsz/ztsaff` | cloned | `f4e0e25f255d` | 223 | git | captured | Shell:120, YAML:31, Other:16, Python:14, Markdown:10 | 107 | 64 | 58 | 3429 |
| `cvsz/zLinebot` | cloned | `0c0c3ec4d7d5` | 486 | git | captured | TypeScript:243, YAML:56, Shell:44, Markdown:42, Python:25 | 566 | 37 | 48 | 646 |
| `cvsz/zlttbots` | cloned | `139bde44dfa3` | 891 | git | captured | Python:352, Markdown:94, YAML:93, Shell:77, Other:69 | 1448 | 177 | 42 | 1062 |
| `cvsz/zttlbots` | cloned | `40949c6038d1` | 52 | git | captured | TypeScript:28, Shell:8, YAML:5, Markdown:4, Other:3 | 56 | 7 | 6 | 99 |
| `cvsz/zTTato-Platform` | blocked | `blocked` | 0 | unavailable | blocked | n/a | 0 | 0 | 0 | 0 |
| `cvsz/zvath` | cloned | `9553b27b0c23` | 219 | git | captured | Python:162, YAML:22, TypeScript:18, SQL:6, Other:3 | 446 | 12 | 9 | 244 |
| `cvsz/tiktok-shop-bot` | cloned | `0e6128856c86` | 11 | git | captured | Python:6, Markdown:4, txt:1 | 22 | 0 | 0 | 1 |
| `cvsz/tiktokshop-api-client` | cloned | `68dc36ff8524` | 16 | git | captured | PHP:9, Other:2, Markdown:1, example:1, json:1 | 32 | 7 | 0 | 69 |
| `cvsz/tiktok-shop-sdk` | cloned | `0c53da6cbba9` | 374 | git | captured | TypeScript:223, Markdown:85, json:35, Other:7, JavaScript:5 | 1921 | 16 | 6 | 1475 |
| `cvsz/tiktokshop-php` | cloned | `dbbec213d9d1` | 57 | git | captured | PHP:49, Markdown:2, Other:2, YAML:2, json:1 | 389 | 7 | 1 | 78 |
| `cvsz/zLinebot-automos` | cloned | `ba8382907971` | 321 | git | captured | JavaScript:98, YAML:54, Shell:29, TypeScript:27, Python:26 | 491 | 115 | 32 | 565 |
| `cvsz/zeaz-platform` | cloned | `b22f68a3f6b8` | 1 | git | captured | Shell:1 | 0 | 0 | 1 | 17 |
| `cvsz/zeapay` | blocked | `blocked` | 0 | unavailable | blocked | n/a | 0 | 0 | 0 | 0 |
| `cvsz/zcino` | cloned | `f2e95ee2b857` | 72 | git | captured | Go:33, TypeScript:17, Other:7, SQL:4, json:4 | 172 | 3 | 1 | 109 |
| `cvsz/zGaming` | cloned | `10f08a39a3ea` | 257 | git | captured | Shell:65, Python:54, PHP:40, TypeScript:30, Markdown:23 | 314 | 12 | 16 | 329 |
| `cvsz/zSafeGuard` | cloned | `94d388d96cc5` | 39 | git | captured | Python:8, Shell:7, JavaScript:5, YAML:5, Markdown:4 | 37 | 11 | 6 | 13 |
| `cvsz/zspin` | cloned | `09550f99b3cc` | 201 | git | captured | Python:113, YAML:33, TypeScript:15, Markdown:12, json:9 | 334 | 24 | 12 | 101 |
| `ZeaZDev/vscode` | cloned | `ea4ccd5f449b` | 8574 | git | captured | TypeScript:5959, json:969, css:295, Other:173, JavaScript:142 | 64674 | 13 | 235 | 23916 |
| `ZeaZDev/ZeaClicker` | cloned | `6a4be8b804fa` | 29 | git | captured | Kotlin:10, Markdown:5, xml:5, gradle:3, Other:2 | 0 | 0 | 0 | 3 |
| `ZeaZDev/zeazchain` | cloned | `78ae5ed30c75` | 57 | git | captured | Markdown:18, YAML:9, TypeScript:8, json:7, Other:4 | 6 | 0 | 4 | 76 |
| `ZeaZDev/ZeaZDev` | cloned | `742079fa3f27` | 82 | git | captured | Shell:34, Markdown:17, TypeScript:9, json:7, Solidity:3 | 186 | 139 | 5 | 2517 |
| `ZeaZDev/ZeaZDev-Omega` | cloned | `41adac95ad68` | 156 | git | captured | TypeScript:54, Markdown:45, Solidity:15, YAML:14, json:12 | 203 | 156 | 16 | 1189 |
| `ZeaZDev/zeazdev-repo` | cloned | `5b73116fbde4` | 24 | git | captured | Python:11, TypeScript:3, Markdown:2, service:2, Other:1 | 51 | 13 | 1 | 32 |
| `ZeaZDev/zeaztools` | cloned | `7a971921b431` | 1 | git | captured | Other:1 | 0 | 0 | 0 | 0 |
| `ZeaZDev/zlms-prod` | cloned | `f26c75d7083b` | 7355 | git | captured | png:3129, gif:1181, ASP.NET:699, PHP:520, html:468 | 6044 | 0 | 40 | 4067 |

Risk-pattern hits are conservative review signals, not exploit proof. They count occurrences of host mutation, persistence, dynamic execution, shell/process invocation, token-like identifiers, `.env`, watchdog/self-heal terms, PM2/nohup, and similar operational-risk markers.

## Repository quarantine decisions

| Repository class | Production admission decision |
|---|---|
| SDK-only repos (`tiktok-shop-sdk`, `tiktokshop-api-client`, `tiktokshop-php`) | Reference-only. Port signing/webhook logic into typed Go integration adapters and keep SDK runtimes out of hot paths. |
| Host-mutating repos (`zgitcp`, `zeaz-platform`, `ztsaff`, `zlttbots`, `zLinebot`, `zLinebot-automos`, `zypto`, `ZeaZDev/ZeaZDev`, `ZeaZDev/ZeaZDev-Omega`, `ZeaZDev/zlms-prod`) | Quarantine scripts. Only signed Terraform, Packer, Kubernetes, and ArgoCD artifacts from this meta repository may mutate infrastructure. |
| Money-movement/signing repos (`zwallet`, `zeapay`, `ZeaZDev-Omega`, `zeazchain`, payment paths in bot repos) | Require ledger migration, idempotency, audit append, transaction tests, Vault leases, and isolated threshold signer/HSM boundaries before user funds or keys are handled. |
| Unaudited/blocked repos (`zTTato-Platform`, `zeapay`) | Blocked from production import until authenticated clone, commit SHA, dependency lockfiles, and secret scan results are supplied. |
| Large legacy/application imports (`ZeaZDev/vscode`, `ZeaZDev/zlms-prod`) | Cloned for evidence. Treat as reference/legacy workloads; extract only deliberate domain logic into normalized Go services after package, license, and secret review. |

## High-value endpoint evidence

Representative source endpoints and API expressions continue to map into the normalized service boundaries in `docs/generated/full-spectrum-analysis.md` and `docs/generated/function-api-inventory.md`. The generated function/API inventory is display-capped per repository to keep this meta repository reviewable while retaining exact summary counts and reproducible extraction logic in `scripts/full-spectrum-audit.py`.
