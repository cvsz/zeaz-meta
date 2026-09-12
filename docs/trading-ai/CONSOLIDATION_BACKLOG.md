# ZTrader / AI Ecosystem Consolidation Backlog

## P0 — security and ownership
- [ ] Zaiman: remove tracked runtime account data and rotate any exposed credentials
- [ ] Zaiman: add real lint/test/secret-scan/security CI gates
- [ ] Publish canonical ownership map
- [ ] Freeze duplicate trading-platform feature development

## P0 — canonical integrations
- [ ] zworkforce -> zksato advisory-intent contract
- [ ] zworkforce -> zwallet read-only on-chain evidence contract
- [ ] zworkforce -> zaiman model-gateway contract
- [ ] Correlated request/trace IDs across all three integrations

## P1 — evidence/data plane
- [ ] social collector adapter
- [ ] EVM on-chain collector
- [ ] Solana on-chain collector
- [ ] holder/deployer/LP provenance
- [ ] freshness/quality metadata
- [ ] durable historical evidence

## P1 — validation
- [ ] paper execution via zksato
- [ ] fee/slippage-aware signal evaluation
- [ ] walk-forward and out-of-sample tests
- [ ] forward-test dataset
- [ ] calibration report

## P1 — UX and operations
- [ ] zdash canonical trader UI
- [ ] LINE alerts/operator commands
- [ ] provider latency/cost/error/fallback metrics
- [ ] trading risk/PnL/position metrics

## P2 — production
- [ ] z-platform GitOps manifests
- [ ] GHCR image pipeline
- [ ] ArgoCD rollout
- [ ] k3s NetworkPolicy/PDB/HPA where appropriate
- [ ] backup/restore evidence
- [ ] incident/rollback drill

## Legacy parity migrations
- [ ] ABTPi18n parity matrix
- [ ] zkbtrader parity matrix
- [ ] OmegaQuant useful research inventory
- [ ] zypto validated feature inventory
- [ ] zeaz/apps/ztrader parity matrix
- [ ] cmeerp/apps/ztrader parity matrix
- [ ] one-api compatibility-test migration to zaiman

No source repository is deleted by this backlog. Archive/delete is a separate gated action.
