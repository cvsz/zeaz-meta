# Source Repository Intake and Evidence Inventory

Audit date: 2026-05-05 UTC.

## Intake commands executed

The mandated GitHub CLI command list was executed after installing `gh` in the container. `gh repo clone` could not clone because the environment had no GitHub authentication token/session, so anonymous HTTPS `git clone --depth 1` was used as the deterministic read-only fallback for public repositories.

```bash
apt-get update && apt-get install -y gh
mkdir -p /zeazdev/zeaz-meta/sources
cd /zeazdev/zeaz-meta/sources
for r in zgitcp zwallet ABTPi18n zypto ZeaZDev-Omega ztsaff zLinebot zlttbots zttlbots zTTato-Platform zvath tiktok-shop-bot tiktokshop-api-client tiktok-shop-sdk tiktokshop-php zLinebot-automos zeaz-platform zeapay; do
  gh repo clone cvsz/$r $r || git clone --depth 1 https://github.com/cvsz/$r.git $r
done
```

`zTTato-Platform` and `zeapay` required GitHub credentials over anonymous HTTPS and remain quarantined as unaudited critical inputs. They must not be imported into production until an authenticated source bundle, commit SHA, dependency lockfiles, and secret scan results are supplied.

## Clone and static metrics

| Repository | Status | Files | Dominant file families | Endpoint hits | Runtime scripts/configs | Risk-pattern hits |
|---|---:|---:|---|---:|---:|---:|
| `ABTPi18n` | cloned | 225 | Python, Markdown, TSX, YAML, shell | 302 | 14 | 199 |
| `ZeaZDev-Omega` | cloned | 58 | TypeScript, TSX, Solidity, Prisma | 21 | 2 | 66 |
| `tiktok-shop-bot` | cloned | 11 | Python CLI, templates | 0 | 0 | 1 |
| `tiktok-shop-sdk` | cloned | 376 | TypeScript SDK, docs, tests | 59 | 0 | 785 |
| `tiktokshop-api-client` | cloned | 16 | PHP Composer SDK | 7 | 0 | 11 |
| `tiktokshop-php` | cloned | 57 | PHP SDK | 53 | 0 | 0 |
| `zLinebot` | cloned | 486 | TypeScript, shell, YAML, Python, SQL | 172 | 54 | 369 |
| `zLinebot-automos` | cloned | 321 | JavaScript, YAML, shell, TypeScript, Python | 219 | 38 | 261 |
| `zeaz-platform` | cloned | 1 | shell | 1 | 1 | 17 |
| `zgitcp` | cloned | 2 | shell, Markdown | 0 | 1 | 38 |
| `zlttbots` | cloned | 893 | Python, shell, JS, YAML, TypeScript | 440 | 142 | 608 |
| `ztsaff` | cloned | 223 | shell, YAML, Python, Terraform | 265 | 143 | 1326 |
| `zttlbots` | cloned | 52 | TypeScript, shell, YAML, SQL | 23 | 11 | 50 |
| `zvath` | cloned | 219 | Python, YAML, TypeScript/TSX, SQL | 41 | 3 | 42 |
| `zwallet` | cloned | 460 | TypeScript, Python, YAML, Kotlin, SQL | 161 | 17 | 117 |
| `zypto` | cloned | 3757 | YAML, Python, Markdown, Terraform, TOML | 316 | 356 | 123 |
| `zTTato-Platform` | blocked | n/a | unavailable | n/a | n/a | n/a |
| `zeapay` | blocked | n/a | unavailable | n/a | n/a | n/a |

Risk-pattern hits were generated with a conservative scan for `while true`, `crontab`, `systemctl`, `sudo`, shell piping, `eval`, `exec`, `child_process`, `subprocess`, `os.system`, `shell=True`, secret-like identifiers, `.env`, watchdog/self-heal terms, `pm2`, and `nohup`. These hits are evidence for review priority, not proof that every line is exploitable.

## Repository quarantine decisions

| Repository | Production admission decision |
|---|---|
| SDK-only repos (`tiktok-shop-sdk`, `tiktokshop-api-client`, `tiktokshop-php`) | Reference-only. Do not deploy SDK runtimes; port signing/webhook logic into the Go integration adapter and keep third-party SDKs out of hot paths. |
| Host-mutating repos (`zgitcp`, `zeaz-platform`, `ztsaff`, `zlttbots`, `zLinebot`, `zLinebot-automos`, `zypto`) | Quarantine scripts. Only signed Terraform, Packer, Kubernetes, and ArgoCD artifacts from this meta repository may mutate infrastructure. |
| Money-movement repos (`zwallet`, `zeapay`, `ZeaZDev-Omega`, payment paths in bot repos) | Require ledger migration, idempotency, audit append, transaction tests, and signer isolation before any user funds are handled. |
| Unaudited repos (`zTTato-Platform`, `zeapay`) | Blocked from production import until authenticated clone and full scan are complete. |
