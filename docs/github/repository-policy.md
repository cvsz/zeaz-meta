# Repository Governance Policy

## Purpose

This policy defines how repositories owned by `cvsz` are created, classified, maintained, consolidated, and retired. It applies to public and private repositories. Public catalogs must never expose private repository names or metadata.

## Lifecycle states

### `governance`

Used for portfolio policy, evidence, architecture, and control-plane automation.

Required:

- documented owner and scope
- reviewed changes through pull requests
- auditable catalog updates
- no product runtime secrets

### `core`

Used for strategic platforms and dependencies whose failure affects multiple products.

Required:

- `main` as the preferred default branch for new or migrated repositories
- protected review path
- CI covering format, lint, typecheck, tests, security, and deployment contracts where applicable
- `README.md`, `SECURITY.md`, `LICENSE`, `CHANGELOG.md`, and ownership declaration
- explicit compatibility and migration policy

### `active`

Used for products and tools with planned development.

Required:

- clear product purpose and target user
- roadmap or active milestone
- owner and support status
- release and validation instructions
- issue and pull-request triage

### `incubating`

Used for experiments and early products.

Required:

- hypothesis or problem statement
- next measurable milestone
- relationship to existing repositories
- review date
- graduation, merge, maintenance, or archive-review decision

### `reference`

Used for forks, mirrors, samples, upstream dependencies, and research copies.

Required:

- upstream repository and license
- pinned or tracked upstream revision
- divergence and local patch policy
- update cadence
- clear statement that the repository is not an independently owned product when applicable

### `maintenance`

Used for supported repositories without planned feature development.

Required:

- supported versions
- security response expectations
- dependency update policy
- end-of-support criteria

### `archive-review`

Used only as a proposal state.

A repository may enter archive review when it is duplicated, superseded, abandoned, unsupported, unsafe to maintain, or no longer aligned with the portfolio. Archive review does not authorize archival.

## Naming policy

- Prefer lowercase kebab-case for new repositories.
- Use the `z-` or established `z` product prefix only when the repository is an owned ZeaZ product or platform component.
- Do not use suffixes such as `-1`, `-new`, `-final`, `-copy`, or version numbers as a long-term product identity.
- Forks and mirrors should retain an upstream-recognizable name and declare their reference status.
- Renames require an impact review covering package names, container images, CI, deployment URLs, imports, submodules, documentation, and redirects.

## Repository metadata baseline

Every `core` and `active` public repository should have:

- concise GitHub description
- relevant topics such as `zeaz`, `platform`, `product`, `agent`, `reference`, or `incubating`
- homepage only when an owned and maintained endpoint exists
- `README.md` with status, purpose, architecture, setup, validation, security, and support
- license appropriate to the source and dependencies
- vulnerability reporting instructions
- declared owner or maintainer
- lifecycle status

## Branch and pull-request policy

- Prefer one complete vertical slice per pull request.
- Do not commit partial migrations to the default branch.
- Revalidate the repository before editing.
- Preserve backward compatibility unless the change explicitly contains a migration plan.
- Use squash merge for atomic product slices unless repository history requires another method.
- Do not merge with failing required checks or unresolved security review threads.
- Delete stale branches only after confirming they contain no unique commits or recovery value.

## Consolidation policy

Before merging or replacing repositories:

1. inventory features, APIs, data, deployment assets, secrets boundaries, licenses, and history
2. declare the target source of truth
3. define compatibility and migration paths
4. preserve or explicitly archive relevant history
5. validate the target implementation
6. redirect documentation and consumers
7. propose archival separately after migration evidence is complete

Git submodules are acceptable when independent history and release ownership must be preserved. Monorepo import is preferred only when build, release, and ownership truly become unified.

## Fork and reference policy

Reference repositories must not be presented as original ZeaZ products without substantial independently maintained work. Security updates and upstream divergence must be visible. GitHub Actions copied from upstream are not automatically trusted and must be reviewed before enabling secrets or write permissions.

## Security-sensitive repositories

Repositories containing malware samples, exploit research, payment code, authentication code, wallet code, automation credentials, or device-control logic require:

- explicit legitimate purpose
- prominent safety and scope notice
- no live credentials or personal data
- controlled CI permissions
- dependency and secret scanning
- owner approval before public release or continued distribution

## Destructive-change approval

The following actions always require explicit, repository-specific owner approval:

- archive or unarchive
- delete
- transfer ownership
- change visibility
- rename
- rewrite default-branch history
- remove releases or tags
- revoke public access to an existing integration

A portfolio document, automated classification, or `archive-review` label is never sufficient authorization.
