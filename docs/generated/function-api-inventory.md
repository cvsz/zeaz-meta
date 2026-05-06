# Function, API, and Pipeline Inventory

Generated static index of detected functions, HTTP/API strings, and automation pipeline files. Each section displays up to 300 entries per kind per repository while summary tables retain complete detected counts. Re-run the script and raise `DISPLAY_LIMIT_PER_REPO` if a local uncapped artifact is required.

## `cvsz/zgitcp`

Status: **cloned**; commit `77075fe9a660`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `git-control-panel.sh` | 43 | Shell | `_log_level_value` |
| `git-control-panel.sh` | 54 | Shell | `_log` |
| `git-control-panel.sh` | 65 | Shell | `log_debug` |
| `git-control-panel.sh` | 66 | Shell | `log_info` |
| `git-control-panel.sh` | 67 | Shell | `log_warn` |
| `git-control-panel.sh` | 68 | Shell | `log_error` |
| `git-control-panel.sh` | 73 | Shell | `json_emit` |
| `git-control-panel.sh` | 90 | Shell | `cleanup` |
| `git-control-panel.sh` | 93 | Shell | `pause` |
| `git-control-panel.sh` | 101 | Shell | `confirm` |
| `git-control-panel.sh` | 115 | Shell | `command_exists` |
| `git-control-panel.sh` | 120 | Shell | `load_config` |
| `git-control-panel.sh` | 140 | Shell | `save_config` |
| `git-control-panel.sh` | 153 | Shell | `edit_config_interactive` |
| `git-control-panel.sh` | 178 | Shell | `check_git_repo` |
| `git-control-panel.sh` | 191 | Shell | `install_gh` |
| `git-control-panel.sh` | 211 | Shell | `install_tea` |
| `git-control-panel.sh` | 286 | Shell | `install_gitea_server` |
| `git-control-panel.sh` | 331 | Shell | `install_ssh_server` |
| `git-control-panel.sh` | 348 | Shell | `connect_host_remote` |
| `git-control-panel.sh` | 368 | Shell | `full_automated_setup` |
| `git-control-panel.sh` | 398 | Shell | `ensure_version_file` |
| `git-control-panel.sh` | 405 | Shell | `generate_changelog` |
| `git-control-panel.sh` | 425 | Shell | `bump_version` |
| `git-control-panel.sh` | 445 | Shell | `create_artifact` |
| `git-control-panel.sh` | 461 | Shell | `upload_artifact_github` |
| `git-control-panel.sh` | 472 | Shell | `upload_artifact_gitea` |
| `git-control-panel.sh` | 489 | Shell | `gh_auth_token` |
| `git-control-panel.sh` | 502 | Shell | `gitea_auth_token` |
| `git-control-panel.sh` | 518 | Shell | `autoheal` |
| `git-control-panel.sh` | 547 | Shell | `self_check` |
| `git-control-panel.sh` | 585 | Shell | `show_usage` |
| `git-control-panel.sh` | 658 | Shell | `cli_release` |
| `git-control-panel.sh` | 732 | Shell | `cli_package` |
| `git-control-panel.sh` | 762 | Shell | `extract_github_owner` |
| `git-control-panel.sh` | 784 | Shell | `stack_generate_source` |
| `git-control-panel.sh` | 891 | Shell | `interactive_menu` |
| `git-control-panel.sh` | 1016 | Shell | `main` |

### Automation pipeline files

None detected or repository blocked.

## `cvsz/zwallet`

Status: **cloned**; commit `8164cb36fd63`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `admin/panel/enterprise/main.py` | 47 | `/admin/health` |
| `admin/panel/enterprise/main.py` | 52 | `/admin/security/metrics` |
| `admin/panel/enterprise/main.py` | 62 | `/admin/security/unblock` |
| `admin/panel/enterprise/main.py` | 75 | `/admin/security/shadow_unban` |
| `admin/panel/enterprise/main.py` | 88 | `/admin/audit/logs` |
| `admin/panel/main.py` | 19 | `/admin/health` |
| `admin/panel/main.py` | 24 | `/admin/security/metrics` |
| `admin/panel/main.py` | 34 | `/admin/security/unblock` |
| `android-app/app/src/main/java/com/zwallet/data/remote/Api.kt` | 8 | `/api/v1/swap/quote` |
| `api/app/main.py` | 54 | `/v1/auth/register` |
| `api/app/main.py` | 64 | `/v1/auth/login` |
| `api/app/main.py` | 71 | `/v1/wallet/transfer` |
| `api/app/main.py` | 80 | `/v1/ai/inference/transaction-anomaly` |
| `api/app/main.py` | 95 | `/v1/ai/inference/user-behavior` |
| `api/app/main.py` | 109 | `/v1/ai/inference/swap-recommendation` |
| `api/app/main.py` | 130 | `/quote` |
| `api/app/main.py` | 139 | `/execute` |
| `api/src/app.ts` | 10 | `/health` |
| `api/src/app.ts` | 19 | `/wallet/balance` |
| `api/src/app.ts` | 25 | `/wallet/transfer` |
| `api/src/app.ts` | 43 | `/deposit/promptpay` |
| `api/src/app.ts` | 52 | `/withdraw/promptpay` |
| `apps/api/src/server.ts` | 20 | `/swap/quote` |
| `apps/api/src/server.ts` | 42 | `/tx/send` |
| `apps/api/src/walletRoutes.ts` | 8 | `/wallet/transfer/preview` |
| `apps/api/src/worldcoinRepos.ts` | 88 | `/integrations/github/worldcoin/repositories` |
| `backend/services/gateway/src/app.ts` | 55 | `/health` |
| `backend/services/gateway/src/app.ts` | 56 | `/v1/auth/register` |
| `backend/services/gateway/src/app.ts` | 58 | `/v1/auth/device/bind` |
| `backend/services/gateway/src/app.ts` | 59 | `/v1/auth/refresh` |
| `backend/services/gateway/src/app.ts` | 60 | `/v1/auth/login` |
| `backend/services/gateway/src/app.ts` | 61 | `/v1/wallet-metadata` |
| `backend/services/gateway/src/app.ts` | 62 | `/v1/transactions/index` |
| `backend/services/gateway/src/app.ts` | 63 | `/v1/swaps/orchestrate` |
| `backend/services/gateway/src/app.ts` | 64 | `/v1/prices/:symbol` |
| `backend/services/gateway/src/app.ts` | 65 | `/v1/audit-logs` |
| `backend/services/gateway/src/app.ts` | 67 | `/v1/compliance/kyc/start` |
| `backend/services/gateway/src/app.ts` | 74 | `/v1/card/issue` |
| `backend/services/gateway/src/app.ts` | 81 | `/v1/webhooks/issuer/auth` |
| `backend/services/gateway/src/app.ts` | 106 | `/v1/card/:id/freeze` |
| `backend/services/gateway/src/app.ts` | 113 | `/v1/transactions/lifecycle` |
| `backend/services/gateway/src/app.ts` | 194 | `/v1/transactions/lifecycle/:txId` |
| `backend/services/gateway/src/app.ts` | 201 | `/v1/mpc/sign-transaction` |
| `backend/services/gateway/src/app.ts` | 213 | `/v1/aa/user-operations` |
| `backend/services/gateway/src/app.ts` | 222 | `/v1/aa/user-operations/:userOpHash/status` |
| `backend/services/gateway/src/app.ts` | 226 | `/v1/flow/wallet-sign-swap` |
| `backend/services/gateway/src/plugins/security.ts` | 60 | `/health` |
| `backend/services/gateway/test/app.test.ts` | 20 | `/v1/wallet-metadata` |
| `backend/services/gateway/test/e2e.wallet-tx-swap.test.ts` | 19 | `/v1/wallet-metadata` |
| `backend/services/gateway/test/e2e.wallet-tx-swap.test.ts` | 21 | `/v1/transactions/lifecycle` |
| `backend/services/gateway/test/e2e.wallet-tx-swap.test.ts` | 23 | `/v1/flow/wallet-sign-swap` |
| `backend/services/gateway/test/helpers.ts` | 27 | `/v1/auth/register` |
| `backend/services/gateway/test/helpers.ts` | 28 | `/v1/auth/login` |
| `backend/services/gateway/test/integration.api-db.test.ts` | 16 | `/v1/wallet-metadata` |
| `backend/services/gateway/test/rpc-mpc-aa.test.ts` | 36 | `/v1/mpc/sign-transaction` |
| `backend/services/gateway/test/rpc-mpc-aa.test.ts` | 40 | `/v1/aa/user-operations` |
| `backend/services/gateway/test/rpc-mpc-aa.test.ts` | 41 | `/v1/aa/user-operations` |
| `backend/services/gateway/test/security.attack-paths.test.ts` | 11 | `/v1/tx/verify-signature` |
| `backend/services/gateway/test/security.attack-paths.test.ts` | 28 | `/health` |
| `backend/services/gateway/test/security.attack-paths.test.ts` | 35 | `/health` |
| `backend/services/gateway/test/security.attack-paths.test.ts` | 49 | `/v1/transactions/lifecycle` |
| `backend/services/indexer-service/src/index.ts` | 11 | `/health` |
| `backend/services/indexer-service/src/index.ts` | 24 | `/worker/run-once` |
| `backend/services/notify-service/src/index.ts` | 4 | `/health` |
| `backend/services/policy-service/src/app.ts` | 11 | `/health` |
| `backend/services/policy-service/src/app.ts` | 32 | `/v1/policy/pre-sign` |
| `backend/services/portfolio-service/src/index.ts` | 4 | `/health` |
| `backend/services/portfolio-service/src/index.ts` | 6 | `/v1/portfolio/display` |
| `backend/services/swap-service/src/index.ts` | 277 | `/health` |
| `backend/services/swap-service/src/index.ts` | 279 | `/v1/swaps/quote` |
| `backend/services/swap-service/src/index.ts` | 296 | `/v1/swaps/execute` |
| `backend/services/tx-orchestrator/src/index.ts` | 26 | `/health` |
| `backend/services/tx-orchestrator/src/index.ts` | 28 | `/v1/tx/simulate` |
| `backend/services/tx-orchestrator/src/index.ts` | 37 | `/v1/tx/verify-signature` |
| `backend/services/wallet-service/src/index.ts` | 5 | `/health` |
| `backend/services/wallet-service/src/index.ts` | 7 | `/v1/wallets/default` |
| `packages/shared-types/src/contracts.ts` | 49 | `/v1/transfers` |
| `packages/shared-types/src/contracts.ts` | 98 | `/v1/transfers` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `.codex/templates/ledger.ts` | 4 | JS/TS | `validate` |
| `admin/panel/enterprise/audit_store.py` | 13 | Python | `get_redis` |
| `admin/panel/enterprise/audit_store.py` | 19 | Python | `append_audit` |
| `admin/panel/enterprise/main.py` | 20 | Python | `get_redis` |
| `admin/panel/enterprise/main.py` | 27 | Python | `verify_admin` |
| `admin/panel/enterprise/main.py` | 40 | Python | `audit_log` |
| `admin/panel/enterprise/main.py` | 48 | Python | `health` |
| `admin/panel/enterprise/main.py` | 53 | Python | `metrics` |
| `admin/panel/enterprise/main.py` | 63 | Python | `unblock` |
| `admin/panel/enterprise/main.py` | 76 | Python | `shadow_unban` |
| `admin/panel/enterprise/main.py` | 89 | Python | `audit_logs` |
| `admin/panel/enterprise/security_hardening.py` | 8 | Python | `enforce_mtls` |
| `admin/panel/enterprise/zero_trust.py` | 5 | Python | `enforce_zero_trust` |
| `admin/panel/main.py` | 14 | Python | `verify_admin` |
| `admin/panel/main.py` | 20 | Python | `health` |
| `admin/panel/main.py` | 25 | Python | `security_metrics` |
| `admin/panel/main.py` | 35 | Python | `unblock` |
| `api/app/application/ai.py` | 18 | Python | `tx_features` |
| `api/app/application/ai.py` | 28 | Python | `behavior_features` |
| `api/app/application/ai.py` | 42 | Python | `__init__` |
| `api/app/application/ai.py` | 46 | Python | `upsert` |
| `api/app/application/ai.py` | 55 | Python | `similarity_search` |
| `api/app/application/ai.py` | 75 | Python | `_cosine_similarity` |
| `api/app/application/ai.py` | 86 | Python | `_bootstrap_neighbors` |
| `api/app/application/ai.py` | 99 | Python | `__init__` |
| `api/app/application/ai.py` | 102 | Python | `detect_transaction_anomaly` |
| `api/app/application/ai.py` | 116 | Python | `analyze_user_behavior` |
| `api/app/application/ai.py` | 126 | Python | `smart_swap_recommendation` |
| `api/app/application/schemas.py` | 10 | Python | `strong_password` |
| `api/app/application/services.py` | 8 | Python | `__init__` |
| `api/app/application/services.py` | 11 | Python | `register` |
| `api/app/application/services.py` | 18 | Python | `login` |
| `api/app/application/services.py` | 26 | Python | `__init__` |
| `api/app/application/services.py` | 29 | Python | `transfer` |
| `api/app/application/swap.py` | 24 | Python | `fetch_quotes` |
| `api/app/application/swap.py` | 55 | Python | `normalize_routes` |
| `api/app/application/swap.py` | 62 | Python | `simulate_transaction` |
| `api/app/application/swap.py` | 70 | Python | `execute_with_fallback` |
| `api/app/core/security.py` | 9 | Python | `hash_password` |
| `api/app/core/security.py` | 13 | Python | `verify_password` |
| `api/app/core/security.py` | 17 | Python | `create_access_token` |
| `api/app/core/security.py` | 23 | Python | `decode_access_token` |
| `api/app/events/idempotency.py` | 10 | Python | `register_once` |
| `api/app/events/idempotency.py` | 17 | Python | `is_duplicate` |
| `api/app/infrastructure/blockchain.py` | 7 | Python | `__init__` |
| `api/app/infrastructure/blockchain.py` | 10 | Python | `transfer_eth` |
| `api/app/infrastructure/db.py` | 14 | Python | `get_session` |
| `api/app/infrastructure/repositories.py` | 9 | Python | `__init__` |
| `api/app/infrastructure/repositories.py` | 12 | Python | `create` |
| `api/app/infrastructure/repositories.py` | 18 | Python | `find_by_email` |
| `api/app/interfaces/http/deps.py` | 10 | Python | `require_user` |
| `api/app/interfaces/http/deps.py` | 18 | Python | `session_dep` |
| `api/app/main.py` | 36 | Python | `lifespan` |
| `api/app/main.py` | 50 | Python | `rate_limit_handler` |
| `api/app/main.py` | 56 | Python | `register` |
| `api/app/main.py` | 66 | Python | `login` |
| `api/app/main.py` | 73 | Python | `transfer` |
| `api/app/main.py` | 82 | Python | `transaction_anomaly_inference` |
| `api/app/main.py` | 97 | Python | `user_behavior_inference` |
| `api/app/main.py` | 111 | Python | `swap_recommendation_inference` |
| `api/app/main.py` | 132 | Python | `quote` |
| `api/app/main.py` | 141 | Python | `execute` |
| `api/app/middleware/idempotency_global.py` | 8 | Python | `__init__` |
| `api/app/middleware/idempotency_global.py` | 12 | Python | `dispatch` |
| `api/app/middleware/security.py` | 17 | Python | `dispatch` |
| `api/app/middleware/security_advanced.py` | 26 | Python | `contains_active_markup` |
| `api/app/middleware/security_advanced.py` | 30 | Python | `get_redis` |
| `api/app/middleware/security_advanced.py` | 37 | Python | `extract_identity` |
| `api/app/middleware/security_advanced.py` | 50 | Python | `dispatch` |
| `api/app/middleware/security_autonomous.py` | 14 | Python | `dispatch` |
| `api/app/middleware/security_battle_ready.py` | 12 | Python | `dispatch` |
| `api/app/middleware/security_distributed.py` | 24 | Python | `contains_active_markup` |
| `api/app/middleware/security_distributed.py` | 28 | Python | `get_redis` |
| `api/app/middleware/security_distributed.py` | 35 | Python | `dispatch` |
| `api/app/middleware/security_intelligent.py` | 11 | Python | `dispatch` |
| `api/app/middleware/security_unified.py` | 12 | Python | `dispatch` |
| `api/app/middleware/security_unified.py` | 17 | Python | `terminal` |
| `api/app/middleware/security_unified.py` | 20 | Python | `run_intelligent` |
| `api/app/middleware/security_unified.py` | 23 | Python | `run_autonomous` |
| `api/app/middleware/security_unified.py` | 29 | Python | `enforce_runtime_prerequisites` |
| `api/app/observability/alerts.py` | 10 | Python | `check_and_alert` |
| `api/app/observability/structured_logger.py` | 10 | Python | `log_event` |
| `api/app/security/adaptive_rate_limiter.py` | 14 | Python | `get_redis` |
| `api/app/security/adaptive_rate_limiter.py` | 20 | Python | `check_rate` |
| `api/app/security/anomaly_detector.py` | 13 | Python | `get_redis` |
| `api/app/security/anomaly_detector.py` | 19 | Python | `record_request` |
| `api/app/security/anomaly_detector.py` | 31 | Python | `is_anomalous` |
| `api/app/security/bot_detection.py` | 13 | Python | `is_bot` |
| `api/app/security/circuit_breaker.py` | 12 | Python | `get_redis` |
| `api/app/security/circuit_breaker.py` | 18 | Python | `is_blocked` |
| `api/app/security/circuit_breaker.py` | 25 | Python | `block` |
| `api/app/security/fingerprint.py` | 6 | Python | `build_fingerprint` |
| `api/app/security/risk_engine.py` | 12 | Python | `get_redis` |
| `api/app/security/risk_engine.py` | 18 | Python | `update_score` |
| `api/app/security/risk_engine.py` | 29 | Python | `is_high_risk` |
| `api/app/security/shadow_ban.py` | 13 | Python | `get_redis` |
| `api/app/security/shadow_ban.py` | 19 | Python | `shadow_ban` |
| `api/app/security/shadow_ban.py` | 26 | Python | `is_shadow_banned` |
| `api/app/security/shadow_ban.py` | 33 | Python | `apply_shadow_behavior` |
| `api/src/app.ts` | 15 | JS/TS | `readBalance` |
| `api/src/app.ts` | 16 | JS/TS | `writeBalance` |
| `api/src/application/deposit.usecase.ts` | 7 | JS/TS | `createDeposit` |
| `api/src/application/deposit.usecase.ts` | 32 | JS/TS | `confirmDeposit` |
| `api/src/application/ledger.service.ts` | 11 | JS/TS | `postDoubleEntry` |
| `api/src/application/payment.lifecycle.ts` | 6 | JS/TS | `markWithdrawSent` |
| `api/src/application/payment.lifecycle.ts` | 39 | JS/TS | `markWithdrawFailed` |
| `api/src/application/payment.lifecycle.ts` | 79 | JS/TS | `markWithdrawSettled` |
| `api/src/application/reconciliation.job.ts` | 6 | JS/TS | `runReconciliation` |
| `api/src/application/saga/transaction.saga.ts` | 3 | JS/TS | `transactionSaga` |
| `api/src/application/transaction.usecase.ts` | 4 | JS/TS | `processTransaction` |
| `api/src/application/transfer.usecase.ts` | 9 | JS/TS | `transfer` |
| `api/src/application/withdraw.usecase.ts` | 7 | JS/TS | `createWithdraw` |
| `api/src/core/final.enforcement.ts` | 40 | JS/TS | `updateBalance` |
| `api/src/core/final.enforcement.ts` | 59 | JS/TS | `postTransaction` |
| `api/src/core/final.enforcement.ts` | 110 | JS/TS | `runOutboxWorker` |
| `api/src/domain/ledger.service.ts` | 14 | JS/TS | `executeTransaction` |
| `api/src/infrastructure/audit.immutable.ts` | 3 | JS/TS | `immutableAudit` |
| `api/src/infrastructure/audit.ts` | 1 | JS/TS | `audit` |
| `api/src/infrastructure/idempotency.global.ts` | 5 | JS/TS | `enforceIdempotency` |
| `api/src/infrastructure/idempotency.global.ts` | 16 | JS/TS | `checkIdempotency` |
| `api/src/infrastructure/idempotency.global.ts` | 23 | JS/TS | `enforceIdempotency` |
| `api/src/infrastructure/idempotency.persistent.ts` | 1 | JS/TS | `enforceIdempotencyPersistent` |
| `api/src/infrastructure/kafka.producer.ts` | 10 | JS/TS | `sendEvent` |
| `api/src/infrastructure/outbox.repository.ts` | 5 | JS/TS | `saveOutboxEvent` |
| `api/src/infrastructure/outbox.repository.ts` | 13 | JS/TS | `fetchPendingEvents` |
| `api/src/infrastructure/outbox.repository.ts` | 21 | JS/TS | `markEventProcessed` |
| `api/src/infrastructure/outbox.worker.ts` | 5 | JS/TS | `runOutboxWorker` |
| `api/src/infrastructure/reconciliation.job.ts` | 5 | JS/TS | `reconcile` |
| `api/src/infrastructure/shard.router.ts` | 3 | JS/TS | `shard` |
| `api/src/infrastructure/sharding.ts` | 5 | JS/TS | `getShardKey` |
| `api/src/infrastructure/vault.injector.ts` | 3 | JS/TS | `loadSecrets` |
| `api/src/infrastructure/webhook.handler.ts` | 7 | JS/TS | `handleWebhook` |
| `api/src/middleware/idempotency.ts` | 5 | JS/TS | `idempotency` |
| `api/src/middleware/identity.ts` | 3 | JS/TS | `extractIdentity` |
| `api/src/middleware/identity.zerotrust.ts` | 3 | JS/TS | `extractIdentity` |
| `api/src/middleware/rateLimit.ts` | 5 | JS/TS | `rateLimit` |
| `api/src/middleware/trace.ts` | 3 | JS/TS | `traceMiddleware` |
| `api/src/security/behavioral.risk.engine.ts` | 11 | JS/TS | `computeRiskScore` |
| `api/src/security/device.fingerprint.ts` | 3 | JS/TS | `generateDeviceFingerprint` |
| `api/src/security/hmac.middleware.ts` | 6 | JS/TS | `verifyHmac` |
| `api/src/security/hmac.v2.middleware.ts` | 6 | JS/TS | `hashBody` |
| `api/src/security/hmac.v2.middleware.ts` | 11 | JS/TS | `verifyHmacV2` |
| `api/src/security/rate.signature.ts` | 5 | JS/TS | `enforceSignatureRate` |
| `api/src/security/replay.protection.ts` | 5 | JS/TS | `preventReplay` |
| `api/src/security/replay.v2.ts` | 5 | JS/TS | `preventReplayV2` |
| `api/src/security/request.attestation.ts` | 3 | JS/TS | `verifyAttestation` |
| `api/src/services/clearing.engine.ts` | 5 | JS/TS | `runClearing` |
| `api/src/services/fraud.service.ts` | 1 | JS/TS | `fraudScore` |
| `api/src/services/fraud.service.ts` | 11 | JS/TS | `fraudDecision` |
| `api/src/services/liquidity.manager.ts` | 9 | JS/TS | `reserveLiquidity` |
| `api/src/services/liquidity.manager.ts` | 19 | JS/TS | `releaseLiquidity` |
| `api/src/services/settlement.engine.ts` | 5 | JS/TS | `settleTransactions` |
| `api/tests/conftest.py` | 19 | Python | `validate_email` |
| `api/tests/conftest.py` | 31 | Python | `api_client` |
| `api/tests/conftest.py` | 35 | Python | `_fake_user` |
| `api/tests/conftest.py` | 39 | Python | `_no_lifespan` |
| `api/tests/test_integration_api_events_failures.py` | 7 | Python | `test_integration_event_pipeline_transaction_and_behavior` |
| `api/tests/test_integration_api_events_failures.py` | 22 | Python | `test_failure_rpc_down` |
| `api/tests/test_integration_api_events_failures.py` | 26 | Python | `_always_rpc_fail` |
| `api/tests/test_integration_api_events_failures.py` | 38 | Python | `test_failure_insufficient_liquidity_and_invalid_tx` |
| `api/tests/test_swap_orchestrator.py` | 4 | Python | `test_fetch_and_normalize_quotes` |
| `api/tests/test_swap_orchestrator.py` | 14 | Python | `test_fetch_quotes_rejects_invalid_slippage_and_amount` |
| `api/tests/test_swap_orchestrator.py` | 29 | Python | `test_simulation_rejects_slippage` |
| `api/tests/test_unit_crypto_swap.py` | 11 | Python | `sign_transaction` |
| `api/tests/test_unit_crypto_swap.py` | 18 | Python | `get_transaction_count` |
| `api/tests/test_unit_crypto_swap.py` | 22 | Python | `send_raw_transaction` |
| `api/tests/test_unit_crypto_swap.py` | 25 | Python | `hex` |
| `api/tests/test_unit_crypto_swap.py` | 35 | Python | `to_wei` |
| `api/tests/test_unit_crypto_swap.py` | 43 | Python | `test_crypto_transfer_builds_and_submits_transaction` |
| `api/tests/test_unit_crypto_swap.py` | 57 | Python | `test_swap_route_selection_includes_gas_and_sorted_descending` |
| `api/tests/test_unit_crypto_swap.py` | 67 | Python | `test_swap_execute_fails_when_all_routes_invalid_slippage` |
| `apps/api/src/walletRoutes.ts` | 7 | JS/TS | `createWalletRoutes` |
| `apps/api/src/worldcoinRepos.ts` | 35 | JS/TS | `fetchOrgReposPage` |
| `apps/api/src/worldcoinRepos.ts` | 64 | JS/TS | `deepFindWorldcoinRepositories` |
| `apps/api/src/worldcoinRepos.ts` | 87 | JS/TS | `createWorldcoinReposRoute` |
| `apps/world/src/main.js` | 10 | JS/TS | `isAddress` |
| `apps/world/src/main.js` | 11 | JS/TS | `isAmount` |
| `apps/world/src/main.js` | 12 | JS/TS | `escapeHtml` |
| `apps/world/src/main.js` | 14 | JS/TS | `section` |
| `apps/world/src/main.js` | 18 | JS/TS | `renderWallet` |
| `apps/world/src/main.js` | 33 | JS/TS | `renderSwap` |
| `apps/world/src/main.js` | 52 | JS/TS | `renderCard` |
| `apps/world/src/main.js` | 66 | JS/TS | `renderFiat` |
| `apps/world/src/main.js` | 85 | JS/TS | `render` |
| `apps/world/src/main.js` | 90 | JS/TS | `bind` |
| `backend/packages/shared/src/index.ts` | 15 | JS/TS | `nowIso` |
| `backend/packages/shared/src/indexer.ts` | 106 | JS/TS | `set` |
| `backend/packages/shared/src/indexer.ts` | 114 | JS/TS | `ingest` |
| `backend/packages/shared/src/indexer.ts` | 121 | JS/TS | `queryByOwner` |
| `backend/packages/shared/src/indexer.ts` | 133 | JS/TS | `constructor` |
| `backend/packages/shared/src/indexer.ts` | 135 | JS/TS | `discover` |
| `backend/packages/shared/src/wallet-engine.ts` | 73 | JS/TS | `wipe` |
| `backend/packages/shared/src/wallet-engine.ts` | 75 | JS/TS | `hex` |
| `backend/packages/shared/src/wallet-engine.ts` | 77 | JS/TS | `fromHex` |
| `backend/packages/shared/src/wallet-engine.ts` | 79 | JS/TS | `deriveKek` |
| `backend/packages/shared/src/wallet-engine.ts` | 83 | JS/TS | `generateMnemonic` |
| `backend/packages/shared/src/wallet-engine.ts` | 94 | JS/TS | `validateMnemonic` |
| `backend/packages/shared/src/wallet-engine.ts` | 99 | JS/TS | `deriveSeed` |
| `backend/packages/shared/src/wallet-engine.ts` | 106 | JS/TS | `deriveAccount` |
| `backend/packages/shared/src/wallet-engine.ts` | 117 | JS/TS | `encryptPrivateKey` |
| `backend/packages/shared/src/wallet-engine.ts` | 137 | JS/TS | `decryptPrivateKey` |
| `backend/packages/shared/src/wallet-engine.ts` | 147 | JS/TS | `signTransaction` |
| `backend/packages/shared/src/wallet-engine.ts` | 163 | JS/TS | `verifyTransactionSignature` |
| `backend/packages/shared/src/wallet-engine.ts` | 169 | JS/TS | `simulateTransaction` |
| `backend/packages/shared/src/wallet-engine.ts` | 183 | JS/TS | `estimateGas` |
| `backend/packages/shared/src/wallet-engine.ts` | 192 | JS/TS | `nextNonce` |
| `backend/packages/shared/src/wallet-engine.ts` | 196 | JS/TS | `verifyReplayProtection` |
| `backend/packages/shared/src/wallet-engine.ts` | 202 | JS/TS | `constantTimeEqual` |
| `backend/packages/shared/src/wallet-engine.ts` | 215 | JS/TS | `pathFor` |
| `backend/packages/shared/src/wallet-engine.ts` | 226 | JS/TS | `addressFor` |
| `backend/packages/shared/src/wallet-engine.ts` | 236 | JS/TS | `put` |
| `backend/packages/shared/src/wallet-engine.ts` | 240 | JS/TS | `get` |
| `backend/packages/shared/src/wallet-engine.ts` | 244 | JS/TS | `delete` |
| `backend/packages/shared/src/wallet-engine.ts` | 250 | JS/TS | `isAvailable` |
| `backend/packages/shared/src/wallet-engine.ts` | 251 | JS/TS | `authorize` |
| `backend/services/gateway/src/app.ts` | 19 | JS/TS | `buildApp` |
| `backend/services/gateway/src/app.ts` | 50 | JS/TS | `authGuard` |
| `backend/services/gateway/src/lib/rpc-provider-pool.ts` | 16 | JS/TS | `constructor` |
| `backend/services/gateway/src/lib/rpc-provider-pool.ts` | 19 | JS/TS | `classify` |
| `backend/services/gateway/src/lib/rpc-provider-pool.ts` | 22 | JS/TS | `isAvailable` |
| `backend/services/gateway/src/lib/rpc-provider-pool.ts` | 27 | JS/TS | `markResult` |
| `backend/services/gateway/src/lib/rpc-provider-pool.ts` | 33 | JS/TS | `call` |
| `backend/services/gateway/src/services/bundler.ts` | 14 | JS/TS | `simulate` |
| `backend/services/gateway/src/services/bundler.ts` | 15 | JS/TS | `submit` |
| `backend/services/gateway/src/services/bundler.ts` | 16 | JS/TS | `getStatus` |
| `backend/services/gateway/src/services/bundler.ts` | 21 | JS/TS | `constructor` |
| `backend/services/gateway/src/services/bundler.ts` | 22 | JS/TS | `createAndSubmit` |
| `backend/services/gateway/src/services/bundler.ts` | 32 | JS/TS | `status` |
| `backend/services/gateway/src/services/card/card-orchestrator.ts` | 30 | JS/TS | `issueCard` |
| `backend/services/gateway/src/services/card/card-orchestrator.ts` | 36 | JS/TS | `handleAuth` |
| `backend/services/gateway/src/services/card/card-orchestrator.ts` | 46 | JS/TS | `freeze` |
| `backend/services/gateway/src/services/card/card-orchestrator.ts` | 51 | JS/TS | `callIssuer` |
| `backend/services/gateway/src/services/compliance/kyc-service.ts` | 11 | JS/TS | `constructor` |
| `backend/services/gateway/src/services/compliance/kyc-service.ts` | 13 | JS/TS | `start` |
| `backend/services/gateway/src/services/compliance/kyc-service.ts` | 24 | JS/TS | `webhookUpdate` |
| `backend/services/gateway/src/services/liquidity/offramp-service.ts` | 17 | JS/TS | `preFund` |
| `backend/services/gateway/src/services/liquidity/offramp-service.ts` | 35 | JS/TS | `fetchQuote` |
| `backend/services/gateway/src/services/mpc.ts` | 18 | JS/TS | `createWallet` |
| `backend/services/gateway/src/services/mpc.ts` | 19 | JS/TS | `getAddress` |
| `backend/services/gateway/src/services/mpc.ts` | 20 | JS/TS | `signTransaction` |
| `backend/services/gateway/src/services/mpc.ts` | 21 | JS/TS | `signMessage` |
| `backend/services/gateway/src/services/mpc.ts` | 22 | JS/TS | `healthCheck` |
| `backend/services/gateway/src/services/mpc.ts` | 26 | JS/TS | `constructor` |
| `backend/services/gateway/src/services/mpc.ts` | 27 | JS/TS | `signTransaction` |
| `backend/services/gateway/src/services/risk/risk-engine.ts` | 15 | JS/TS | `evaluate` |
| `backend/services/gateway/src/services/state-store.ts` | 25 | JS/TS | `constructor` |
| `backend/services/gateway/src/services/state-store.ts` | 39 | JS/TS | `createUser` |
| `backend/services/gateway/src/services/state-store.ts` | 46 | JS/TS | `getUser` |
| `backend/services/gateway/src/services/state-store.ts` | 52 | JS/TS | `bindDevice` |
| `backend/services/gateway/src/services/state-store.ts` | 53 | JS/TS | `hasDevice` |
| `backend/services/gateway/src/services/state-store.ts` | 54 | JS/TS | `appendAudit` |
| `backend/services/gateway/src/services/state-store.ts` | 55 | JS/TS | `readAudit` |
| `backend/services/gateway/src/services/state-store.ts` | 57 | JS/TS | `writeTxLifecycle` |
| `backend/services/gateway/src/services/state-store.ts` | 61 | JS/TS | `readTxLifecycle` |
| `backend/services/gateway/test/helpers.ts` | 3 | JS/TS | `fakeRedis` |
| `backend/services/gateway/test/helpers.ts` | 14 | JS/TS | `resetStore` |
| `backend/services/gateway/test/helpers.ts` | 26 | JS/TS | `registerAndLogin` |
| `backend/services/gateway/test/rpc-mpc-aa.test.ts` | 17 | JS/TS | `delay` |
| `backend/services/indexer-service/src/chains.ts` | 37 | JS/TS | `parseEvmTransfer` |
| `backend/services/indexer-service/src/chains.ts` | 52 | JS/TS | `parseSolanaTransaction` |
| `backend/services/indexer-service/src/chains.ts` | 78 | JS/TS | `parseBitcoinUtxoEvent` |
| `backend/services/indexer-service/src/state-store.ts` | 4 | JS/TS | `tryStartJob` |
| `backend/services/indexer-service/src/state-store.ts` | 5 | JS/TS | `markEventProcessed` |
| `backend/services/indexer-service/src/state-store.ts` | 6 | JS/TS | `applyBalanceDeltas` |
| `backend/services/indexer-service/src/state-store.ts` | 7 | JS/TS | `applyUtxoDeltas` |
| `backend/services/indexer-service/src/state-store.ts` | 16 | JS/TS | `tryStartJob` |
| `backend/services/indexer-service/src/state-store.ts` | 23 | JS/TS | `markEventProcessed` |
| `backend/services/indexer-service/src/state-store.ts` | 30 | JS/TS | `applyBalanceDeltas` |
| `backend/services/indexer-service/src/state-store.ts` | 38 | JS/TS | `applyUtxoDeltas` |
| `backend/services/indexer-service/src/worker.ts` | 13 | JS/TS | `publish` |
| `backend/services/indexer-service/src/worker.ts` | 14 | JS/TS | `consume` |
| `backend/services/indexer-service/src/worker.ts` | 19 | JS/TS | `publish` |
| `backend/services/indexer-service/src/worker.ts` | 20 | JS/TS | `consume` |
| `backend/services/indexer-service/src/worker.ts` | 24 | JS/TS | `constructor` |
| `backend/services/indexer-service/src/worker.ts` | 26 | JS/TS | `drainOnce` |
| `backend/services/indexer-service/src/worker.ts` | 47 | JS/TS | `toStateUpdates` |
| `backend/services/indexer-service/src/worker.ts` | 53 | JS/TS | `fingerprint` |
| `backend/services/indexer-service/src/worker.ts` | 57 | JS/TS | `isChainMessage` |
| `backend/services/policy-service/src/app.ts` | 6 | JS/TS | `createApp` |
| `backend/services/policy-service/src/contracts.ts` | 25 | JS/TS | `evaluatePreSign` |
| `backend/services/policy-service/src/service.ts` | 6 | JS/TS | `evaluatePreSign` |
| `backend/services/swap-service/src/index.ts` | 86 | JS/TS | `buildRoutes` |
| `backend/services/swap-service/src/index.ts` | 98 | JS/TS | `getQuoteRoutes` |
| `backend/services/swap-service/src/index.ts` | 102 | JS/TS | `routeScore` |
| `backend/services/swap-service/src/index.ts` | 111 | JS/TS | `fetchRoutesFromAggregators` |
| `backend/services/swap-service/src/index.ts` | 127 | JS/TS | `normalizeRoute` |
| `backend/services/swap-service/src/index.ts` | 141 | JS/TS | `calculatePriceImpactPct` |
| `backend/services/swap-service/src/index.ts` | 148 | JS/TS | `simulateRoute` |
| `backend/services/swap-service/src/index.ts` | 162 | JS/TS | `slippageAdjusted` |
| `backend/services/swap-service/src/index.ts` | 166 | JS/TS | `gasUsd` |
| `backend/services/swap-service/src/index.ts` | 173 | JS/TS | `baseRoute` |
| `backend/services/swap-service/src/index.ts` | 190 | JS/TS | `oneInch` |
| `backend/services/swap-service/src/index.ts` | 195 | JS/TS | `uniswapV2` |
| `backend/services/swap-service/src/index.ts` | 200 | JS/TS | `uniswapV3` |
| `backend/services/swap-service/src/index.ts` | 210 | JS/TS | `jupiter` |
| `backend/services/swap-service/src/index.ts` | 215 | JS/TS | `raydium` |
| `backend/services/swap-service/src/index.ts` | 220 | JS/TS | `rfqSolver` |
| `backend/services/swap-service/src/index.ts` | 229 | JS/TS | `executeWithFallback` |
| `backend/services/swap-service/src/index.ts` | 243 | JS/TS | `executeSingleRoute` |
| `infra/installer/install.prod.sh` | 19 | Shell | `log` |

_Display capped at 300 of 419 detected functions/methods._

### Automation pipeline files

- `.codex/workflows.yaml`
- `.github/workflows/chaos.yml`
- `.github/workflows/ci.yml`
- `.github/workflows/cicd.yml`
- `.github/workflows/codeql.yml`
- `.github/workflows/codex-autofix.yml`
- `.github/workflows/codex-pr-gate.yml`
- `.github/workflows/repo-hygiene.yml`
- `.github/workflows/ssot-enforcement.yml`
- `.github/workflows/supply-chain.yml`
- `api/src/application/saga/transaction.saga.ts`
- `api/src/application/transaction.usecase.ts`
- `docker-compose.yml`
- `docs/DEVOPS_PIPELINE.md`
- `docs/production-validation-pipeline.md`
- `infra/chaos/chaos-pipeline.yaml`
- `infra/docker/docker-compose.devops.yml`
- `infra/docker/docker-compose.prod.yml`
- `infra/helm/templates/api-deployment.yaml`
- `infra/installer/install.prod.sh`
- `infra/installer/install.sh`
- `infra/k8s/api-deployment.yaml`
- `infra/k8s/helm/zwallet/templates/deployment.yaml`
- `infra/k8s/operators/deploy-operators.sh`
- `infra/observability/docker-compose.siem.yml`
- `infra/redis/docker-compose.yml`
- `k8s/base/gateway-deployment.yaml`
- `k8s/deployment.yaml`
- `packages/crypto-core/src/txPipeline.ts`
- `scripts/install.sh`
- `scripts/installer.config.json`
- `scripts/run-installer.mjs`

## `cvsz/ABTPi18n`

Status: **cloned**; commit `6749c6f6c414`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `.github/agents/abtpro_omega.agent.md` | 98 | `/` |
| `apps/backend/main.py` | 71 | `/metrics` |
| `apps/backend/main.py` | 107 | `/health` |
| `apps/backend/main.py` | 126 | `/auth/login` |
| `apps/backend/main.py` | 135 | `/exchange/keys` |
| `apps/backend/main.py` | 153 | `/strategies` |
| `apps/backend/main.py` | 158 | `/dashboard/pnl` |
| `apps/backend/main.py` | 173 | `/auth` |
| `apps/backend/main.py` | 188 | `/health` |
| `apps/backend/src/api/audit_endpoints.py` | 32 | `/logs` |
| `apps/backend/src/api/audit_endpoints.py` | 86 | `/logs/{log_id}` |
| `apps/backend/src/api/audit_endpoints.py` | 97 | `/stats` |
| `apps/backend/src/api/audit_endpoints.py` | 129 | `/export` |
| `apps/backend/src/api/auth_endpoints.py` | 33 | `/google/authorize` |
| `apps/backend/src/api/auth_endpoints.py` | 46 | `/google/login` |
| `apps/backend/src/api/auth_endpoints.py` | 74 | `/google/callback` |
| `apps/backend/src/api/auth_endpoints.py` | 116 | `/me` |
| `apps/backend/src/api/auth_endpoints.py` | 130 | `/logout` |
| `apps/backend/src/api/backtest_endpoints.py` | 43 | `/run` |
| `apps/backend/src/api/backtest_endpoints.py` | 73 | `/runs` |
| `apps/backend/src/api/backtest_endpoints.py` | 83 | `/runs/{run_id}` |
| `apps/backend/src/api/backtest_endpoints.py` | 93 | `/runs/{run_id}` |
| `apps/backend/src/api/backtest_endpoints.py` | 106 | `/paper/start` |
| `apps/backend/src/api/backtest_endpoints.py` | 125 | `/paper/stop` |
| `apps/backend/src/api/backtest_endpoints.py` | 137 | `/paper/status/{session_id}` |
| `apps/backend/src/api/backtest_endpoints.py` | 147 | `/paper/sessions` |
| `apps/backend/src/api/bot_endpoints.py` | 29 | `/start` |
| `apps/backend/src/api/bot_endpoints.py` | 49 | `/stop` |
| `apps/backend/src/api/health_endpoints.py` | 28 | `/detailed` |
| `apps/backend/src/api/health_endpoints.py` | 92 | `/database` |
| `apps/backend/src/api/health_endpoints.py` | 128 | `/redis` |
| `apps/backend/src/api/health_endpoints.py` | 175 | `/system` |
| `apps/backend/src/api/ml_endpoints.py` | 68 | `/signal/score` |
| `apps/backend/src/api/ml_endpoints.py` | 94 | `/signal/history` |
| `apps/backend/src/api/ml_endpoints.py` | 106 | `/signal/stats` |
| `apps/backend/src/api/ml_endpoints.py` | 120 | `/volatility/predict` |
| `apps/backend/src/api/ml_endpoints.py` | 155 | `/volatility/history` |
| `apps/backend/src/api/ml_endpoints.py` | 166 | `/volatility/accuracy` |
| `apps/backend/src/api/ml_endpoints.py` | 182 | `/tune/start` |
| `apps/backend/src/api/ml_endpoints.py` | 198 | `/tune/status/{job_id}` |
| `apps/backend/src/api/ml_endpoints.py` | 211 | `/tune/results/{job_id}` |
| `apps/backend/src/api/ml_endpoints.py` | 224 | `/tune/apply` |
| `apps/backend/src/api/ml_endpoints.py` | 238 | `/models/list` |
| `apps/backend/src/api/ml_endpoints.py` | 255 | `/models/train` |
| `apps/backend/src/api/ml_endpoints.py` | 269 | `/models/performance` |
| `apps/backend/src/api/payment_endpoints.py` | 32 | `/promptpay/create` |
| `apps/backend/src/api/payment_endpoints.py` | 49 | `/webhook/promptpay` |
| `apps/backend/src/api/payment_endpoints.py` | 72 | `/wallet` |
| `apps/backend/src/api/payment_endpoints.py` | 82 | `/transactions` |
| `apps/backend/src/api/plugin_endpoints.py` | 29 | `/available` |
| `apps/backend/src/api/plugin_endpoints.py` | 44 | `/installed` |
| `apps/backend/src/api/plugin_endpoints.py` | 54 | `/install` |
| `apps/backend/src/api/plugin_endpoints.py` | 68 | `/uninstall/{user_plugin_id}` |
| `apps/backend/src/api/plugin_endpoints.py` | 81 | `/{user_plugin_id}/toggle` |
| `apps/backend/src/api/plugin_endpoints.py` | 91 | `/discover` |
| `apps/backend/src/api/portfolio_endpoints.py` | 26 | `/summary` |
| `apps/backend/src/api/portfolio_endpoints.py` | 36 | `/accounts` |
| `apps/backend/src/api/portfolio_endpoints.py` | 46 | `/accounts` |
| `apps/backend/src/api/portfolio_endpoints.py` | 63 | `/accounts/{account_id}` |
| `apps/backend/src/api/portfolio_endpoints.py` | 73 | `/accounts/{account_id}/sync` |
| `apps/backend/src/api/portfolio_endpoints.py` | 83 | `/accounts/{account_id}/performance` |
| `apps/backend/src/api/preferences_endpoints.py` | 46 | `/preferences/{user_id}` |
| `apps/backend/src/api/preferences_endpoints.py` | 72 | `/preferences` |
| `apps/backend/src/api/preferences_endpoints.py` | 116 | `/theme` |
| `apps/backend/src/api/preferences_endpoints.py` | 131 | `/language` |
| `apps/backend/src/api/preferences_endpoints.py` | 146 | `/notifications/preferences/{user_id}` |
| `apps/backend/src/api/preferences_endpoints.py` | 170 | `/notifications/preferences` |
| `apps/backend/src/api/rental_endpoints.py` | 29 | `/contract` |
| `apps/backend/src/api/rental_endpoints.py` | 41 | `/contract` |
| `apps/backend/src/api/rental_endpoints.py` | 55 | `/renew` |
| `apps/backend/src/api/rental_endpoints.py` | 69 | `/plans` |
| `apps/backend/src/api/rental_endpoints.py` | 76 | `/plans/{plan_id}` |
| `apps/backend/src/api/rental_endpoints.py` | 85 | `/features` |
| `apps/backend/src/api/rental_endpoints.py` | 97 | `/has-feature/{feature_name}` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 33 | `/rotation/schedule` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 51 | `/rotation/rotate` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 94 | `/rotation/history` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 114 | `/rotation/due` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 139 | `/rotation/policy` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 173 | `/rotation/complete/{secret_type}/{secret_name}` |
| `apps/backend/src/api/telegram_endpoints.py` | 35 | `/link` |
| `apps/backend/src/api/telegram_endpoints.py` | 48 | `/unlink` |
| `apps/backend/src/api/telegram_endpoints.py` | 59 | `/status/{user_id}` |
| `apps/backend/src/api/telegram_endpoints.py` | 66 | `/notify/test` |
| `apps/backend/src/api/tradingview_endpoints.py` | 87 | `/webhook` |
| `apps/backend/src/api/tradingview_endpoints.py` | 174 | `/alerts` |
| `apps/backend/src/api/tradingview_endpoints.py` | 228 | `/config` |
| `apps/backend/src/services/audit_middleware.py` | 60 | `/metrics` |
| `apps/backend/src/services/audit_middleware.py` | 61 | `/health` |
| `docs/enterprise/CONFIGURATION.en.md` | 453 | `/metrics` |
| `docs/enterprise/CONFIGURATION.th.md` | 453 | `/metrics` |
| `docs/phases/phase2/PHASE2_IMPLEMENTATION_SUMMARY.md` | 608 | `/metrics` |
| `docs/phases/phase2/PHASE2_IMPLEMENTATION_SUMMARY.md` | 630 | `/metrics` |
| `docs/phases/phase2/PHASE2_IMPLEMENTATION_SUMMARY.md` | 798 | `/metrics` |
| `docs/phases/phase6/PHASE6_IMPLEMENTATION_SUMMARY.md` | 413 | `/signal/score` |
| `docs/phases/phase6/PHASE6_IMPLEMENTATION_SUMMARY.md` | 444 | `/tune/start` |
| `docs/phases/phase6/PHASE6_IMPLEMENTATION_SUMMARY.md` | 477 | `/tune/status/{job_id}` |
| `docs/phases/phase6/PHASE6_IMPLEMENTATION_SUMMARY.md` | 501 | `/volatility/predict` |
| `docs/phases/phase6/PHASE6_IMPLEMENTATION_SUMMARY.md` | 732 | `/ml/metrics/signal_accuracy` |
| `monitoring/prometheus.yml` | 14 | `/metrics` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `apps/backend/main.py` | 88 | Python | `startup` |
| `apps/backend/main.py` | 98 | Python | `shutdown` |
| `apps/backend/main.py` | 108 | Python | `health` |
| `apps/backend/main.py` | 127 | Python | `login` |
| `apps/backend/main.py` | 136 | Python | `save_exchange_keys` |
| `apps/backend/main.py` | 154 | Python | `list_strategies` |
| `apps/backend/main.py` | 159 | Python | `dashboard_pnl` |
| `apps/backend/src/api/audit_endpoints.py` | 33 | Python | `get_audit_logs` |
| `apps/backend/src/api/audit_endpoints.py` | 87 | Python | `get_audit_log` |
| `apps/backend/src/api/audit_endpoints.py` | 98 | Python | `get_audit_stats` |
| `apps/backend/src/api/audit_endpoints.py` | 130 | Python | `export_audit_logs` |
| `apps/backend/src/api/auth_endpoints.py` | 34 | Python | `google_authorize` |
| `apps/backend/src/api/auth_endpoints.py` | 47 | Python | `google_login` |
| `apps/backend/src/api/auth_endpoints.py` | 75 | Python | `google_callback` |
| `apps/backend/src/api/auth_endpoints.py` | 117 | Python | `get_current_user` |
| `apps/backend/src/api/auth_endpoints.py` | 131 | Python | `logout` |
| `apps/backend/src/api/backtest_endpoints.py` | 44 | Python | `create_backtest` |
| `apps/backend/src/api/backtest_endpoints.py` | 74 | Python | `list_backtests` |
| `apps/backend/src/api/backtest_endpoints.py` | 84 | Python | `get_backtest_results` |
| `apps/backend/src/api/backtest_endpoints.py` | 94 | Python | `delete_backtest` |
| `apps/backend/src/api/backtest_endpoints.py` | 107 | Python | `start_paper_trading` |
| `apps/backend/src/api/backtest_endpoints.py` | 126 | Python | `stop_paper_trading` |
| `apps/backend/src/api/backtest_endpoints.py` | 138 | Python | `get_paper_trading_status` |
| `apps/backend/src/api/backtest_endpoints.py` | 148 | Python | `list_paper_trading_sessions` |
| `apps/backend/src/api/bot_endpoints.py` | 30 | Python | `start_bot` |
| `apps/backend/src/api/bot_endpoints.py` | 50 | Python | `stop_bot` |
| `apps/backend/src/api/health_endpoints.py` | 20 | Python | `basic_health_check` |
| `apps/backend/src/api/health_endpoints.py` | 29 | Python | `detailed_health_check` |
| `apps/backend/src/api/health_endpoints.py` | 93 | Python | `database_health_check` |
| `apps/backend/src/api/health_endpoints.py` | 129 | Python | `redis_health_check` |
| `apps/backend/src/api/health_endpoints.py` | 176 | Python | `system_health_check` |
| `apps/backend/src/api/ml_endpoints.py` | 69 | Python | `score_signal` |
| `apps/backend/src/api/ml_endpoints.py` | 95 | Python | `get_signal_history` |
| `apps/backend/src/api/ml_endpoints.py` | 107 | Python | `get_signal_stats` |
| `apps/backend/src/api/ml_endpoints.py` | 121 | Python | `predict_volatility` |
| `apps/backend/src/api/ml_endpoints.py` | 156 | Python | `get_volatility_history` |
| `apps/backend/src/api/ml_endpoints.py` | 167 | Python | `get_volatility_accuracy` |
| `apps/backend/src/api/ml_endpoints.py` | 183 | Python | `start_tuning` |
| `apps/backend/src/api/ml_endpoints.py` | 199 | Python | `get_tuning_status` |
| `apps/backend/src/api/ml_endpoints.py` | 212 | Python | `get_tuning_results` |
| `apps/backend/src/api/ml_endpoints.py` | 225 | Python | `apply_optimized_params` |
| `apps/backend/src/api/ml_endpoints.py` | 239 | Python | `list_models` |
| `apps/backend/src/api/ml_endpoints.py` | 256 | Python | `train_model` |
| `apps/backend/src/api/ml_endpoints.py` | 270 | Python | `get_model_performance` |
| `apps/backend/src/api/payment_endpoints.py` | 33 | Python | `create_promptpay_payment` |
| `apps/backend/src/api/payment_endpoints.py` | 50 | Python | `promptpay_webhook` |
| `apps/backend/src/api/payment_endpoints.py` | 73 | Python | `get_wallet_balance` |
| `apps/backend/src/api/payment_endpoints.py` | 83 | Python | `get_transaction_history` |
| `apps/backend/src/api/plugin_endpoints.py` | 30 | Python | `list_available_plugins` |
| `apps/backend/src/api/plugin_endpoints.py` | 45 | Python | `list_installed_plugins` |
| `apps/backend/src/api/plugin_endpoints.py` | 55 | Python | `install_plugin` |
| `apps/backend/src/api/plugin_endpoints.py` | 69 | Python | `uninstall_plugin` |
| `apps/backend/src/api/plugin_endpoints.py` | 82 | Python | `toggle_plugin` |
| `apps/backend/src/api/plugin_endpoints.py` | 92 | Python | `discover_plugins` |
| `apps/backend/src/api/portfolio_endpoints.py` | 27 | Python | `get_portfolio_summary` |
| `apps/backend/src/api/portfolio_endpoints.py` | 37 | Python | `list_accounts` |
| `apps/backend/src/api/portfolio_endpoints.py` | 47 | Python | `create_account` |
| `apps/backend/src/api/portfolio_endpoints.py` | 64 | Python | `delete_account` |
| `apps/backend/src/api/portfolio_endpoints.py` | 74 | Python | `sync_account_positions` |
| `apps/backend/src/api/portfolio_endpoints.py` | 84 | Python | `get_account_performance` |
| `apps/backend/src/api/preferences_endpoints.py` | 47 | Python | `get_user_preferences` |
| `apps/backend/src/api/preferences_endpoints.py` | 73 | Python | `update_user_preferences` |
| `apps/backend/src/api/preferences_endpoints.py` | 117 | Python | `update_theme` |
| `apps/backend/src/api/preferences_endpoints.py` | 132 | Python | `update_language` |
| `apps/backend/src/api/preferences_endpoints.py` | 147 | Python | `get_notification_preferences` |
| `apps/backend/src/api/preferences_endpoints.py` | 171 | Python | `update_notification_preferences` |
| `apps/backend/src/api/rental_endpoints.py` | 30 | Python | `get_user_contract` |
| `apps/backend/src/api/rental_endpoints.py` | 42 | Python | `create_contract` |
| `apps/backend/src/api/rental_endpoints.py` | 56 | Python | `renew_contract` |
| `apps/backend/src/api/rental_endpoints.py` | 70 | Python | `list_plans` |
| `apps/backend/src/api/rental_endpoints.py` | 77 | Python | `get_plan_info` |
| `apps/backend/src/api/rental_endpoints.py` | 86 | Python | `get_contract_features` |
| `apps/backend/src/api/rental_endpoints.py` | 98 | Python | `check_feature_access` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 34 | Python | `get_rotation_schedule` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 52 | Python | `rotate_secret` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 95 | Python | `get_rotation_history` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 115 | Python | `get_secrets_due_for_rotation` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 140 | Python | `update_rotation_policy` |
| `apps/backend/src/api/secret_rotation_endpoints.py` | 174 | Python | `mark_rotation_complete` |
| `apps/backend/src/api/telegram_endpoints.py` | 36 | Python | `link_telegram` |
| `apps/backend/src/api/telegram_endpoints.py` | 49 | Python | `unlink_telegram` |
| `apps/backend/src/api/telegram_endpoints.py` | 60 | Python | `get_telegram_status` |
| `apps/backend/src/api/telegram_endpoints.py` | 67 | Python | `send_test_notification` |
| `apps/backend/src/api/tradingview_endpoints.py` | 57 | Python | `verify_webhook_secret` |
| `apps/backend/src/api/tradingview_endpoints.py` | 88 | Python | `tradingview_webhook` |
| `apps/backend/src/api/tradingview_endpoints.py` | 175 | Python | `list_tradingview_alerts` |
| `apps/backend/src/api/tradingview_endpoints.py` | 229 | Python | `get_webhook_config` |
| `apps/backend/src/auth/google_provider.py` | 16 | Python | `__init__` |
| `apps/backend/src/auth/google_provider.py` | 30 | Python | `get_authorization_url` |
| `apps/backend/src/auth/google_provider.py` | 52 | Python | `exchange_code_for_token` |
| `apps/backend/src/auth/oauth_service.py` | 19 | Python | `__init__` |
| `apps/backend/src/auth/oauth_service.py` | 24 | Python | `verify_google_token` |
| `apps/backend/src/auth/oauth_service.py` | 44 | Python | `get_or_create_user` |
| `apps/backend/src/auth/oauth_service.py` | 108 | Python | `generate_session_token` |
| `apps/backend/src/backtesting/backtest_service.py` | 21 | Python | `__init__` |
| `apps/backend/src/backtesting/backtest_service.py` | 24 | Python | `create_backtest` |
| `apps/backend/src/backtesting/backtest_service.py` | 61 | Python | `run_backtest` |
| `apps/backend/src/backtesting/backtest_service.py` | 119 | Python | `_generate_mock_backtest_results` |
| `apps/backend/src/backtesting/backtest_service.py` | 165 | Python | `get_backtest_results` |
| `apps/backend/src/backtesting/backtest_service.py` | 196 | Python | `list_user_backtests` |
| `apps/backend/src/backtesting/backtest_service.py` | 221 | Python | `delete_backtest` |
| `apps/backend/src/backtesting/backtest_service.py` | 242 | Python | `start_paper_trading` |
| `apps/backend/src/backtesting/backtest_service.py` | 278 | Python | `stop_paper_trading` |
| `apps/backend/src/backtesting/backtest_service.py` | 322 | Python | `get_paper_trading_status` |
| `apps/backend/src/backtesting/backtest_service.py` | 357 | Python | `list_paper_trading_sessions` |
| `apps/backend/src/ml/reinforcement/environment.py` | 11 | Python | `__init__` |
| `apps/backend/src/ml/reinforcement/environment.py` | 27 | Python | `reset` |
| `apps/backend/src/ml/reinforcement/environment.py` | 40 | Python | `step` |
| `apps/backend/src/ml/reinforcement/environment.py` | 68 | Python | `get_current_params` |
| `apps/backend/src/ml/reinforcement/environment.py` | 72 | Python | `_get_state` |
| `apps/backend/src/ml/reinforcement/environment.py` | 82 | Python | `_apply_action` |
| `apps/backend/src/ml/reinforcement/environment.py` | 100 | Python | `_execute_trading_step` |
| `apps/backend/src/ml/reinforcement/environment.py` | 133 | Python | `_generate_signal` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 15 | Python | `__init__` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 34 | Python | `optimize` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 97 | Python | `_get_action` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 115 | Python | `_random_action` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 121 | Python | `_best_action` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 135 | Python | `_get_possible_actions` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 143 | Python | `_update_q_value` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 157 | Python | `_evaluate_params` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 220 | Python | `_simulate_strategy` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 228 | Python | `_calculate_sharpe_ratio` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 239 | Python | `_calculate_max_drawdown` |
| `apps/backend/src/ml/reinforcement/tuner.py` | 248 | Python | `_calculate_win_rate` |
| `apps/backend/src/ml/signal_quality/feature_extractor.py` | 13 | Python | `extract_signal_features` |
| `apps/backend/src/ml/signal_quality/feature_extractor.py` | 93 | Python | `extract_market_features` |
| `apps/backend/src/ml/signal_quality/scorer.py` | 18 | Python | `__init__` |
| `apps/backend/src/ml/signal_quality/scorer.py` | 34 | Python | `train` |
| `apps/backend/src/ml/signal_quality/scorer.py` | 70 | Python | `score_signal` |
| `apps/backend/src/ml/signal_quality/scorer.py` | 136 | Python | `save_model` |
| `apps/backend/src/ml/signal_quality/scorer.py` | 149 | Python | `load_model` |
| `apps/backend/src/ml/volatility/features.py` | 13 | Python | `extract` |
| `apps/backend/src/ml/volatility/features.py` | 108 | Python | `_count_price_jumps` |
| `apps/backend/src/ml/volatility/features.py` | 114 | Python | `_calculate_trend_strength` |
| `apps/backend/src/ml/volatility/predictor.py` | 17 | Python | `__init__` |
| `apps/backend/src/ml/volatility/predictor.py` | 33 | Python | `train` |
| `apps/backend/src/ml/volatility/predictor.py` | 69 | Python | `predict_volatility` |
| `apps/backend/src/ml/volatility/predictor.py` | 140 | Python | `save_model` |
| `apps/backend/src/ml/volatility/predictor.py` | 153 | Python | `load_model` |
| `apps/backend/src/plugins/plugin_loader.py` | 29 | Python | `__init__` |
| `apps/backend/src/plugins/plugin_loader.py` | 34 | Python | `initialize` |
| `apps/backend/src/plugins/plugin_loader.py` | 39 | Python | `cleanup` |
| `apps/backend/src/plugins/plugin_loader.py` | 50 | Python | `on_candle` |
| `apps/backend/src/plugins/plugin_loader.py` | 69 | Python | `calculate` |
| `apps/backend/src/plugins/plugin_loader.py` | 85 | Python | `__init__` |
| `apps/backend/src/plugins/plugin_loader.py` | 93 | Python | `discover_plugins` |
| `apps/backend/src/plugins/plugin_loader.py` | 146 | Python | `register_plugin` |
| `apps/backend/src/plugins/plugin_loader.py` | 197 | Python | `install_plugin_for_user` |
| `apps/backend/src/plugins/plugin_loader.py` | 243 | Python | `toggle_plugin` |
| `apps/backend/src/plugins/plugin_loader.py` | 256 | Python | `get_user_plugins` |
| `apps/backend/src/plugins/plugin_loader.py` | 279 | Python | `load_plugin` |
| `apps/backend/src/plugins/plugin_loader.py` | 316 | Python | `unload_plugin` |
| `apps/backend/src/plugins/plugin_loader.py` | 324 | Python | `list_available_plugins` |
| `apps/backend/src/security/crypto_service.py` | 21 | Python | `encrypt_data` |
| `apps/backend/src/security/crypto_service.py` | 28 | Python | `decrypt_data` |
| `apps/backend/src/services/audit_middleware.py` | 19 | Python | `__init__` |
| `apps/backend/src/services/audit_middleware.py` | 23 | Python | `get_user_id` |
| `apps/backend/src/services/audit_middleware.py` | 39 | Python | `get_action_from_method` |
| `apps/backend/src/services/audit_middleware.py` | 56 | Python | `should_log` |
| `apps/backend/src/services/audit_middleware.py` | 73 | Python | `dispatch` |
| `apps/backend/src/services/audit_service.py` | 23 | Python | `__init__` |
| `apps/backend/src/services/audit_service.py` | 27 | Python | `connect` |
| `apps/backend/src/services/audit_service.py` | 32 | Python | `disconnect` |
| `apps/backend/src/services/audit_service.py` | 37 | Python | `sanitize_data` |
| `apps/backend/src/services/audit_service.py` | 69 | Python | `log_api_call` |
| `apps/backend/src/services/audit_service.py` | 126 | Python | `get_logs` |
| `apps/backend/src/services/audit_service.py` | 209 | Python | `get_log_by_id` |
| `apps/backend/src/services/audit_service.py` | 235 | Python | `get_stats` |
| `apps/backend/src/services/audit_service.py` | 296 | Python | `cleanup_old_logs` |
| `apps/backend/src/services/exchange_service.py` | 19 | Python | `for_exchange` |
| `apps/backend/src/services/gdrive_loader.py` | 28 | Python | `__init__` |
| `apps/backend/src/services/gdrive_loader.py` | 38 | Python | `download_file` |
| `apps/backend/src/services/gdrive_loader.py` | 71 | Python | `download_folder` |
| `apps/backend/src/services/gdrive_loader.py` | 111 | Python | `load_strategy_config` |
| `apps/backend/src/services/gdrive_loader.py` | 139 | Python | `load_all_configs` |
| `apps/backend/src/services/gdrive_loader.py` | 171 | Python | `_extract_file_id` |
| `apps/backend/src/services/gdrive_loader.py` | 180 | Python | `_extract_folder_id` |
| `apps/backend/src/services/gdrive_loader.py` | 189 | Python | `clear_cache` |
| `apps/backend/src/services/gdrive_loader.py` | 200 | Python | `load_tradingview_strategies_from_gdrive` |
| `apps/backend/src/services/metrics_service.py` | 105 | Python | `record_trade` |
| `apps/backend/src/services/metrics_service.py` | 116 | Python | `record_strategy_signal` |
| `apps/backend/src/services/metrics_service.py` | 123 | Python | `record_risk_check` |
| `apps/backend/src/services/metrics_service.py` | 129 | Python | `record_circuit_breaker_trip` |
| `apps/backend/src/services/metrics_service.py` | 134 | Python | `update_risk_metrics` |
| `apps/backend/src/services/metrics_service.py` | 152 | Python | `update_bot_status` |
| `apps/backend/src/services/metrics_service.py` | 160 | Python | `time_strategy_execution` |
| `apps/backend/src/services/metrics_service.py` | 165 | Python | `record_websocket_connection` |
| `apps/backend/src/services/metrics_service.py` | 171 | Python | `record_websocket_message` |
| `apps/backend/src/services/metrics_service.py` | 176 | Python | `record_websocket_error` |
| `apps/backend/src/services/metrics_service.py` | 181 | Python | `record_exchange_api_call` |
| `apps/backend/src/services/metrics_service.py` | 194 | Python | `set_system_info` |
| `apps/backend/src/services/notification_service.py` | 17 | Python | `__init__` |
| `apps/backend/src/services/notification_service.py` | 21 | Python | `send_trade_notification` |
| `apps/backend/src/services/notification_service.py` | 65 | Python | `send_risk_alert` |
| `apps/backend/src/services/notification_service.py` | 101 | Python | `send_system_notification` |
| `apps/backend/src/services/notification_service.py` | 131 | Python | `send_daily_summary` |
| `apps/backend/src/services/portfolio_service.py` | 23 | Python | `__init__` |
| `apps/backend/src/services/portfolio_service.py` | 27 | Python | `create_account` |
| `apps/backend/src/services/portfolio_service.py` | 67 | Python | `list_accounts` |
| `apps/backend/src/services/portfolio_service.py` | 90 | Python | `sync_account_positions` |
| `apps/backend/src/services/portfolio_service.py` | 162 | Python | `get_portfolio_summary` |
| `apps/backend/src/services/portfolio_service.py` | 219 | Python | `get_account_performance` |
| `apps/backend/src/services/portfolio_service.py` | 255 | Python | `_get_exchange_instance` |
| `apps/backend/src/services/portfolio_service.py` | 275 | Python | `delete_account` |
| `apps/backend/src/services/promptpay_service.py` | 23 | Python | `__init__` |
| `apps/backend/src/services/promptpay_service.py` | 28 | Python | `generate_qr_code` |
| `apps/backend/src/services/promptpay_service.py` | 63 | Python | `create_payment_intent` |
| `apps/backend/src/services/promptpay_service.py` | 120 | Python | `verify_webhook_signature` |
| `apps/backend/src/services/promptpay_service.py` | 137 | Python | `process_payment_callback` |
| `apps/backend/src/services/promptpay_service.py` | 190 | Python | `_credit_wallet` |
| `apps/backend/src/services/promptpay_service.py` | 206 | Python | `get_wallet_balance` |
| `apps/backend/src/services/promptpay_service.py` | 227 | Python | `get_transaction_history` |
| `apps/backend/src/services/rental_service.py` | 58 | Python | `__init__` |
| `apps/backend/src/services/rental_service.py` | 61 | Python | `get_user_contract` |
| `apps/backend/src/services/rental_service.py` | 94 | Python | `create_contract` |
| `apps/backend/src/services/rental_service.py` | 130 | Python | `renew_contract` |
| `apps/backend/src/services/rental_service.py` | 187 | Python | `check_contract_expiry` |
| `apps/backend/src/services/rental_service.py` | 230 | Python | `expire_contract` |
| `apps/backend/src/services/rental_service.py` | 255 | Python | `has_feature` |
| `apps/backend/src/services/rental_service.py` | 283 | Python | `get_plan_info` |
| `apps/backend/src/services/rental_service.py` | 287 | Python | `list_all_plans` |
| `apps/backend/src/services/secret_rotation_service.py` | 18 | Python | `__init__` |
| `apps/backend/src/services/secret_rotation_service.py` | 25 | Python | `connect` |
| `apps/backend/src/services/secret_rotation_service.py` | 30 | Python | `disconnect` |
| `apps/backend/src/services/secret_rotation_service.py` | 35 | Python | `hash_secret` |
| `apps/backend/src/services/secret_rotation_service.py` | 47 | Python | `create_rotation_record` |
| `apps/backend/src/services/secret_rotation_service.py` | 110 | Python | `get_rotation_schedule` |
| `apps/backend/src/services/secret_rotation_service.py` | 151 | Python | `get_rotation_history` |
| `apps/backend/src/services/secret_rotation_service.py` | 193 | Python | `get_secrets_due_for_rotation` |
| `apps/backend/src/services/secret_rotation_service.py` | 229 | Python | `update_rotation_policy` |
| `apps/backend/src/services/secret_rotation_service.py` | 281 | Python | `mark_rotation_complete` |
| `apps/backend/src/services/telegram_service.py` | 19 | Python | `__init__` |
| `apps/backend/src/services/telegram_service.py` | 24 | Python | `send_message` |
| `apps/backend/src/services/telegram_service.py` | 37 | Python | `verify_chat` |
| `apps/backend/src/services/telegram_service.py` | 48 | Python | `link_telegram_account` |
| `apps/backend/src/services/telegram_service.py` | 98 | Python | `unlink_telegram_account` |
| `apps/backend/src/services/telegram_service.py` | 124 | Python | `get_telegram_link_status` |
| `apps/backend/src/services/telegram_service.py` | 142 | Python | `generate_verification_code` |
| `apps/backend/src/services/websocket_service.py` | 20 | Python | `__init__` |
| `apps/backend/src/services/websocket_service.py` | 33 | Python | `connect` |
| `apps/backend/src/services/websocket_service.py` | 52 | Python | `disconnect` |
| `apps/backend/src/services/websocket_service.py` | 68 | Python | `subscribe_ticker` |
| `apps/backend/src/services/websocket_service.py` | 89 | Python | `unsubscribe_ticker` |
| `apps/backend/src/services/websocket_service.py` | 102 | Python | `_stream_ticker` |
| `apps/backend/src/services/websocket_service.py` | 147 | Python | `subscribe_trades` |
| `apps/backend/src/services/websocket_service.py` | 166 | Python | `_stream_trades` |
| `apps/backend/src/services/websocket_service.py` | 210 | Python | `get_active_subscriptions` |
| `apps/backend/src/services/websocket_service.py` | 222 | Python | `get_websocket_instance` |
| `apps/backend/src/trading/bot_runner.py` | 21 | Python | `__init__` |
| `apps/backend/src/trading/bot_runner.py` | 25 | Python | `assess` |
| `apps/backend/src/trading/bot_runner.py` | 35 | Python | `__init__` |
| `apps/backend/src/trading/bot_runner.py` | 45 | Python | `load_bot` |
| `apps/backend/src/trading/bot_runner.py` | 52 | Python | `fetch_ohlcv` |
| `apps/backend/src/trading/bot_runner.py` | 55 | Python | `run_loop` |
| `apps/backend/src/trading/bot_runner.py` | 120 | Python | `record_trade` |
| `apps/backend/src/trading/risk_manager.py` | 16 | Python | `__init__` |
| `apps/backend/src/trading/risk_manager.py` | 26 | Python | `update_equity` |
| `apps/backend/src/trading/risk_manager.py` | 39 | Python | `is_drawdown_exceeded` |
| `apps/backend/src/trading/risk_manager.py` | 47 | Python | `get_metrics` |
| `apps/backend/src/trading/risk_manager.py` | 67 | Python | `__init__` |
| `apps/backend/src/trading/risk_manager.py` | 87 | Python | `record_trade_outcome` |
| `apps/backend/src/trading/risk_manager.py` | 105 | Python | `trip_breaker` |
| `apps/backend/src/trading/risk_manager.py` | 111 | Python | `is_tripped` |
| `apps/backend/src/trading/risk_manager.py` | 124 | Python | `check_trade_rate_limit` |
| `apps/backend/src/trading/risk_manager.py` | 133 | Python | `get_status` |
| `apps/backend/src/trading/risk_manager.py` | 156 | Python | `__init__` |
| `apps/backend/src/trading/risk_manager.py` | 175 | Python | `assess` |
| `apps/backend/src/trading/risk_manager.py` | 218 | Python | `update_equity_from_trades` |
| `apps/backend/src/trading/risk_manager.py` | 228 | Python | `record_trade_result` |
| `apps/backend/src/trading/risk_manager.py` | 237 | Python | `get_metrics` |
| `apps/backend/src/trading/strategies/breakout_strategy.py` | 17 | Python | `__init__` |
| `apps/backend/src/trading/strategies/breakout_strategy.py` | 29 | Python | `execute` |
| `apps/backend/src/trading/strategies/mean_reversion_strategy.py` | 18 | Python | `__init__` |
| `apps/backend/src/trading/strategies/mean_reversion_strategy.py` | 43 | Python | `execute` |
| `apps/backend/src/trading/strategies/rsi_cross_strategy.py` | 18 | Python | `__init__` |
| `apps/backend/src/trading/strategies/rsi_cross_strategy.py` | 24 | Python | `compute_rsi` |
| `apps/backend/src/trading/strategies/rsi_cross_strategy.py` | 35 | Python | `execute` |
| `apps/backend/src/trading/strategies/tradingview_strategy.py` | 25 | Python | `__init__` |
| `apps/backend/src/trading/strategies/tradingview_strategy.py` | 36 | Python | `execute` |
| `apps/backend/src/trading/strategies/vwap_strategy.py` | 17 | Python | `__init__` |
| `apps/backend/src/trading/strategies/vwap_strategy.py` | 27 | Python | `calculate_vwap` |
| `apps/backend/src/trading/strategies/vwap_strategy.py` | 33 | Python | `execute` |
| `apps/backend/src/trading/strategy_interface.py` | 15 | Python | `execute` |
| `apps/backend/src/trading/strategy_interface.py` | 26 | Python | `register` |
| `apps/backend/src/trading/strategy_interface.py` | 30 | Python | `create` |
| `apps/backend/src/trading/strategy_interface.py` | 36 | Python | `list_names` |
| `apps/backend/src/utils/auth.py` | 12 | Python | `create_session_response` |
| `apps/backend/src/utils/database.py` | 17 | Python | `get_db_connection` |
| `apps/backend/src/utils/database.py` | 35 | Python | `time_database_operation` |
| `apps/backend/src/utils/dependencies.py` | 12 | Python | `get_current_user_id` |
| `apps/backend/src/utils/dependencies.py` | 39 | Python | `get_optional_user_id` |
| `apps/backend/src/utils/exceptions.py` | 11 | Python | `raise_bad_request` |
| `apps/backend/src/utils/exceptions.py` | 16 | Python | `raise_not_found` |
| `apps/backend/src/utils/exceptions.py` | 21 | Python | `raise_internal_error` |
| `apps/backend/src/utils/exceptions.py` | 26 | Python | `handle_service_error` |
| `apps/backend/src/worker/tasks.py` | 24 | Python | `run_bot_loop` |
| `apps/backend/src/worker/tasks.py` | 28 | Python | `run_bot_async` |
| `apps/backend/src/worker/tasks.py` | 37 | Python | `check_contract_expiry` |
| `apps/backend/src/worker/tasks.py` | 42 | Python | `check_contract_expiry_async` |

_Display capped at 300 of 411 detected functions/methods._

### Automation pipeline files

- `.github/workflows/agent-runner.yml`
- `.github/workflows/lint-test.yml`
- `.github/workflows/metaultra.yml`
- `.github/workflows/release.yml`
- `.github/workflows/security-scan.yml`
- `docker-compose.yml`
- `docs/enterprise/DEPLOYMENT.en.md`
- `docs/setup/INSTALLER_PLATFORM_REQUIREMENTS.md`
- `install.sh`
- `scripts/zeaz_meta_installer.sh`

## `cvsz/zypto`

Status: **cloned**; commit `ef0fdd5ce580`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `cert_saas/api/app.py` | 8 | `/attestation` |
| `compliance_marketplace/api/app.py` | 6 | `/health` |
| `decentralized/api/main.py` | 11 | `/node/register` |
| `decentralized/api/main.py` | 18 | `/nodes` |
| `decentralized/api/main.py` | 23 | `/account/create` |
| `decentralized/api/main.py` | 29 | `/account/fund` |
| `decentralized/api/main.py` | 35 | `/job/submit` |
| `enterprise/dashboard/api.py` | 10 | `/dashboard` |
| `enterprise/frontend/pages/EnterpriseDashboard.tsx` | 8 | `/api/dashboard` |
| `frontend/pages/CRM.tsx` | 9 | `/api/pipeline` |
| `frontend/pages/Forecast.tsx` | 9 | `/api/forecast` |
| `pitchdeck/demo/app.py` | 68 | `/api/tenants` |
| `pitchdeck/demo/app.py` | 73 | `/api/tenant/{tenant_id}/snapshot` |
| `pitchdeck/demo/app.py` | 100 | `/` |
| `pitchdeck/demo/app.py` | 142 | `/api/tenants` |
| `saas/backend/main.py` | 24 | `/auth/signup` |
| `saas/backend/main.py` | 30 | `/auth/me` |
| `saas/backend/main.py` | 57 | `/healthz` |
| `saas/backend/stripe_api.py` | 15 | `/create-checkout` |
| `sales_crm/api.py` | 9 | `/pipeline` |
| `sales_crm/api.py` | 14 | `/forecast` |
| `services/tss-signing/cmd/server/main.go` | 21 | `/healthz` |
| `services/tss-signing/cmd/server/main.go` | 25 | `/v1/sign` |
| `testing/locustfile.py` | 9 | `/api/metrics` |
| `tests/contract/openapi_snapshot.json` | 8 | `/metrics/{tenant}` |
| `tests/contract/test_orders_contract.py` | 11 | `/metrics/tenantA` |
| `tests/security/test_tenant_isolation.py` | 9 | `/metrics/tenantA` |
| `tests/security/test_tenant_isolation.py` | 16 | `/metrics/tenantB` |
| `v1/app/api/main.py` | 13 | `/auth` |
| `v1/app/api/routes_admin.py` | 11 | `/strategy/toggle` |
| `v1/app/api/routes_admin.py` | 18 | `/enterprise/rbac/assign` |
| `v1/app/api/routes_admin.py` | 24 | `/enterprise/rbac` |
| `v1/app/api/routes_admin.py` | 30 | `/enterprise/audit` |
| `v1/app/api/routes_admin.py` | 36 | `/enterprise/slo` |
| `v1/app/api/routes_admin.py` | 42 | `/enterprise/slo` |
| `v1/app/api/routes_admin.py` | 48 | `/growth/campaign` |
| `v1/app/api/routes_admin.py` | 54 | `/growth/automation/actions` |
| `v1/app/api/routes_admin.py` | 60 | `/growth/analytics` |
| `v1/app/api/routes_aiops.py` | 11 | `/kill` |
| `v1/app/api/routes_aiops.py` | 18 | `/limit` |
| `v1/app/api/routes_auth.py` | 8 | `/login` |
| `v1/app/api/routes_god.py` | 15 | `/kill-switch` |
| `v1/app/api/routes_god.py` | 22 | `/budget` |
| `v1/app/api/routes_god.py` | 30 | `/circuit` |
| `v1/app/api/routes_user.py` | 15 | `/portfolio` |
| `v1/app/api/routes_user.py` | 25 | `/onboarding/stage` |
| `v1/app/main.py` | 14 | `/health` |
| `v1/infra/scripts/ui.py` | 27 | `/deploy` |
| `v1/infra/scripts/ui.py` | 41 | `/status` |
| `zcyptobot-hyperscale/scripts/bootstrap.py` | 19 | `/health` |
| `zcyptobot-hyperscale/services/service_0001/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0002/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0003/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0004/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0005/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0006/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0007/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0008/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0009/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0010/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0011/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0012/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0013/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0014/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0015/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0016/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0017/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0018/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0019/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0020/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0021/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0022/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0023/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0024/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0025/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0026/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0027/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0028/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0029/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0030/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0031/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0032/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0033/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0034/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0035/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0036/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0037/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0038/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0039/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0040/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0041/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0042/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0043/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0044/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0045/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0046/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0047/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0048/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0049/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0050/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0051/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0052/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0053/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0054/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0055/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0056/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0057/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0058/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0059/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0060/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0061/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0062/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0063/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0064/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0065/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0066/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0067/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0068/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0069/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0070/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0071/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0072/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0073/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0074/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0075/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0076/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0077/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0078/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0079/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0080/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0081/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0082/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0083/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0084/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0085/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0086/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0087/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0088/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0089/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0090/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0091/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0092/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0093/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0094/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0095/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0096/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0097/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0098/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0099/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0100/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0101/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0102/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0103/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0104/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0105/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0106/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0107/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0108/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0109/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0110/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0111/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0112/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0113/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0114/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0115/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0116/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0117/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0118/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0119/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0120/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0121/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0122/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0123/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0124/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0125/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0126/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0127/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0128/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0129/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0130/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0131/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0132/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0133/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0134/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0135/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0136/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0137/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0138/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0139/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0140/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0141/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0142/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0143/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0144/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0145/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0146/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0147/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0148/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0149/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0150/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0151/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0152/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0153/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0154/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0155/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0156/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0157/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0158/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0159/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0160/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0161/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0162/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0163/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0164/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0165/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0166/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0167/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0168/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0169/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0170/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0171/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0172/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0173/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0174/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0175/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0176/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0177/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0178/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0179/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0180/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0181/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0182/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0183/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0184/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0185/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0186/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0187/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0188/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0189/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0190/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0191/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0192/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0193/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0194/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0195/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0196/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0197/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0198/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0199/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0200/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0201/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0202/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0203/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0204/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0205/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0206/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0207/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0208/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0209/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0210/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0211/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0212/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0213/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0214/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0215/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0216/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0217/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0218/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0219/src/main.py` | 5 | `/health` |
| `zcyptobot-hyperscale/services/service_0220/src/main.py` | 5 | `/health` |
| `zeaz-live/backend/main.py` | 38 | `/metrics/{tenant}` |
| `zeaz-live/backend/main.py` | 45 | `/admin/usage/{tenant}` |
| `zeaz-live/backend/main.py` | 78 | `/i18n/messages` |
| `zeaz-prod/backend/main.py` | 23 | `/auth/token` |
| `zeaz_super_gen.sh` | 20 | `/health` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `adoption/incentives.py` | 1 | Python | `incentive` |
| `ai/agents/negotiation.py` | 2 | Python | `propose` |
| `ai/agents/negotiation.py` | 7 | Python | `propose` |
| `ai/agents/negotiation.py` | 12 | Python | `propose` |
| `ai/agents/negotiation.py` | 17 | Python | `decide` |
| `ai/control/engine.py` | 11 | Python | `propose` |
| `ai/control/engine.py` | 22 | Python | `decide` |
| `ai/control/self_evolve.py` | 14 | Python | `evolve` |
| `ai/cost_opt/model.py` | 2 | Python | `score` |
| `ai/cost_opt/planner.py` | 4 | Python | `choose_region` |
| `ai/fraud/detector.py` | 4 | Python | `is_fraud` |
| `ai/fraud/graph_model.py` | 6 | Python | `add_tx` |
| `ai/fraud/graph_model.py` | 10 | Python | `anomaly_score` |
| `ai/org/system.py` | 2 | Python | `__init__` |
| `ai/org/system.py` | 5 | Python | `act` |
| `ai/org/system.py` | 10 | Python | `__init__` |
| `ai/org/system.py` | 18 | Python | `run` |
| `ai/org_v2/engine.py` | 6 | Python | `__init__` |
| `ai/org_v2/engine.py` | 9 | Python | `generate_code` |
| `ai/org_v2/engine.py` | 12 | Python | `deploy` |
| `ai/org_v2/engine.py` | 17 | Python | `__init__` |
| `ai/org_v2/engine.py` | 25 | Python | `execute` |
| `ai/pricing/model.py` | 6 | Python | `__init__` |
| `ai/pricing/model.py` | 14 | Python | `forward` |
| `ai/pricing/online_infer.py` | 4 | Python | `predict` |
| `ai/pricing/trainer.py` | 9 | Python | `train_step` |
| `ai/pricing_v2/bandit.py` | 5 | Python | `__init__` |
| `ai/pricing_v2/bandit.py` | 9 | Python | `select` |
| `ai/pricing_v2/bandit.py` | 13 | Python | `update` |
| `ai/world_model/model.py` | 5 | Python | `__init__` |
| `ai/world_model/model.py` | 11 | Python | `forward` |
| `ai/world_model/planner.py` | 3 | Python | `plan` |
| `apps/api/middleware/policy.py` | 6 | Python | `policy_middleware` |
| `audit_simulation/evidence/controls.py` | 1 | Python | `validate` |
| `backend/audit.py` | 7 | Python | `log_event` |
| `backend/audit.py` | 18 | Python | `query_logs` |
| `backend/middleware_rbac.py` | 4 | Python | `rbac_middleware` |
| `backend/rbac.py` | 8 | Python | `authorize` |
| `backend/slo.py` | 13 | Python | `query` |
| `backend/slo.py` | 23 | Python | `latency_p95` |
| `backend/slo.py` | 27 | Python | `error_rate` |
| `backend/tenant.py` | 1 | Python | `tenant_filter` |
| `bi/ingest/batch_loader.py` | 13 | Python | `_hash_user_id` |
| `bi/ingest/batch_loader.py` | 17 | Python | `load_file` |
| `bi/ingest/kafka_consumer.py` | 29 | Python | `_hash_user_id` |
| `bi/ingest/kafka_consumer.py` | 33 | Python | `_normalize_event` |
| `bi/ingest/kafka_consumer.py` | 47 | Python | `run` |
| `billing/invoice.py` | 4 | Python | `generate_invoice` |
| `billing/tax.py` | 1 | Python | `apply_tax` |
| `cert_saas/api/app.py` | 9 | Python | `attest` |
| `cert_saas/api/auth.py` | 4 | Python | `authorize` |
| `cert_saas/api/tenants.py` | 4 | Python | `create` |
| `cert_saas/billing/metering.py` | 4 | Python | `record` |
| `cert_saas/billing/pricing.py` | 1 | Python | `price` |
| `cert_saas/core/attestation.py` | 1 | Python | `generate` |
| `cert_saas/core/control_engine.py` | 8 | Python | `run` |
| `cert_saas/core/evidence_store.py` | 8 | Python | `put` |
| `cert_saas/core/policy_engine.py` | 1 | Python | `evaluate` |
| `cert_saas/core/scheduler.py` | 4 | Python | `loop` |
| `compliance_marketplace/api/app.py` | 7 | Python | `health` |
| `compliance_marketplace/api/auth.py` | 4 | Python | `authorize` |
| `compliance_marketplace/api/tenants.py` | 4 | Python | `create` |
| `compliance_marketplace/core/catalog.py` | 4 | Python | `add` |
| `compliance_marketplace/core/escrow.py` | 4 | Python | `hold` |
| `compliance_marketplace/core/escrow.py` | 10 | Python | `release` |
| `compliance_marketplace/core/listing.py` | 4 | Python | `search` |
| `compliance_marketplace/core/orders.py` | 4 | Python | `create` |
| `compliance_marketplace/core/ratings.py` | 4 | Python | `rate` |
| `compliance_marketplace/core/ratings.py` | 10 | Python | `avg` |
| `compliance_marketplace/core/settlement.py` | 1 | Python | `split` |
| `compliance_marketplace/governance/disputes.py` | 1 | Python | `resolve` |
| `core/audit/log.py` | 5 | Python | `log` |
| `data_treaty/enforcement/audit.py` | 6 | Python | `log` |
| `data_treaty/enforcement/geo_policy.py` | 1 | Python | `allow` |
| `data_treaty/enforcement/transfer_gate.py` | 4 | Python | `transfer` |
| `data_treaty/interoperability/federation.py` | 1 | Python | `federate` |
| `data_treaty/interoperability/portability.py` | 1 | Python | `export` |
| `decentralized/api/main.py` | 12 | Python | `node_register` |
| `decentralized/api/main.py` | 19 | Python | `nodes` |
| `decentralized/api/main.py` | 24 | Python | `acc_create` |
| `decentralized/api/main.py` | 30 | Python | `acc_fund` |
| `decentralized/api/main.py` | 36 | Python | `job_submit` |
| `decentralized/coordinator/federated_scheduler.py` | 11 | Python | `select` |
| `decentralized/coordinator/federated_scheduler.py` | 41 | Python | `audit_decision` |
| `decentralized/coordinator/scheduler.py` | 11 | Python | `__init__` |
| `decentralized/coordinator/scheduler.py` | 14 | Python | `submit_job` |
| `decentralized/coordinator/scheduler.py` | 19 | Python | `add_provider` |
| `decentralized/coordinator/scheduler.py` | 24 | Python | `run` |
| `decentralized/fairness/drf.py` | 1 | Python | `dominant_share` |
| `decentralized/fairness/drf.py` | 9 | Python | `pick_min_dominant` |
| `decentralized/fairness/quota.py` | 4 | Python | `set_limit` |
| `decentralized/fairness/quota.py` | 8 | Python | `allow` |
| `decentralized/k8s/deploy_job.py` | 4 | Python | `deploy` |
| `decentralized/ledger/accounts.py` | 6 | Python | `create_account` |
| `decentralized/ledger/accounts.py` | 10 | Python | `credit` |
| `decentralized/ledger/accounts.py` | 14 | Python | `debit` |
| `decentralized/ledger/accounts.py` | 20 | Python | `balance` |
| `decentralized/ledger/escrow.py` | 4 | Python | `lock` |
| `decentralized/ledger/escrow.py` | 9 | Python | `release` |
| `decentralized/ledger/escrow.py` | 14 | Python | `refund` |
| `decentralized/ledger/settlement.py` | 5 | Python | `settle_success` |
| `decentralized/ledger/settlement.py` | 9 | Python | `settle_refund` |
| `decentralized/market/matching.py` | 4 | Python | `match` |
| `decentralized/market/orderbook.py` | 19 | Python | `add` |
| `decentralized/market/pricing.py` | 1 | Python | `dynamic_price` |
| `decentralized/p2p/gossip.py` | 4 | Python | `subscribe` |
| `decentralized/p2p/gossip.py` | 8 | Python | `publish` |
| `decentralized/p2p/libp2p_node.py` | 11 | Python | `on_msg` |
| `decentralized/p2p/libp2p_node.py` | 16 | Python | `start` |
| `decentralized/p2p/libp2p_node.py` | 19 | Python | `loop` |
| `decentralized/p2p/peer_store.py` | 6 | Python | `add_peer` |
| `decentralized/p2p/peer_store.py` | 10 | Python | `list_peers` |
| `decentralized/registry/node_registry.py` | 6 | Python | `register` |
| `decentralized/registry/node_registry.py` | 11 | Python | `list_nodes` |
| `decentralized/registry/node_registry.py` | 15 | Python | `get` |
| `decentralized/reputation/score.py` | 4 | Python | `update` |
| `decentralized/reputation/score.py` | 10 | Python | `get` |
| `decentralized/tee/attest_verify.py` | 5 | Python | `verify_quote` |
| `decentralized/trust/byzantine.py` | 4 | Python | `filter_outliers` |
| `decentralized/trust/byzantine.py` | 11 | Python | `quorum` |
| `decentralized/trust/scoring.py` | 4 | Python | `update` |
| `decentralized/trust/scoring.py` | 12 | Python | `get` |
| `decentralized/zk/audit.py` | 4 | Python | `append` |
| `decentralized/zk/audit.py` | 10 | Python | `verify_chain` |
| `decentralized/zk/commit.py` | 6 | Python | `commit` |
| `default_layer/control_plane/jwks_sync.py` | 4 | Python | `update` |
| `default_layer/control_plane/jwks_sync.py` | 8 | Python | `get` |
| `default_layer/control_plane/policy.py` | 1 | Python | `allow` |
| `default_layer/control_plane/revocation.py` | 4 | Python | `revoke` |
| `default_layer/control_plane/revocation.py` | 8 | Python | `is_revoked` |
| `default_layer/control_plane/router.py` | 1 | Python | `choose_issuer` |
| `default_layer/data_plane/audit_log.py` | 6 | Python | `log` |
| `default_layer/data_plane/metrics.py` | 1 | Python | `kpis` |
| `default_layer/sdk/adapters/aws.py` | 1 | Python | `map_to_iam` |
| `default_layer/sdk/adapters/azure.py` | 1 | Python | `map_to_entra` |
| `default_layer/sdk/adapters/gcp.py` | 1 | Python | `map_to_sa` |
| `default_layer/sdk/cache.py` | 4 | Python | `get` |
| `default_layer/sdk/cache.py` | 8 | Python | `set` |
| `default_layer/sdk/middleware/express.ts` | 1 | JS/TS | `gidMiddleware` |
| `default_layer/sdk/middleware/fastapi.py` | 5 | Python | `gid_middleware` |
| `default_layer/sdk/middleware/fastapi.py` | 6 | Python | `middleware` |
| `devops/oidc/fastapi_oidc.py` | 11 | Python | `_jwk_client` |
| `devops/oidc/fastapi_oidc.py` | 15 | Python | `verify_oidc` |
| `dora/reporting/incident_report.py` | 2 | Python | `report` |
| `dora/reporting/kpi.py` | 2 | Python | `kpi` |
| `dora/third_party/register.py` | 5 | Python | `add` |
| `dora/third_party/risk_scoring.py` | 3 | Python | `score` |
| `economy/compliance/aml.py` | 4 | Python | `suspicious` |
| `economy/compliance/limits.py` | 6 | Python | `check` |
| `economy/credits/escrow.py` | 8 | Python | `lock` |
| `economy/credits/escrow.py` | 14 | Python | `release` |
| `economy/credits/ledger.py` | 8 | Python | `mint` |
| `economy/credits/ledger.py` | 13 | Python | `spend` |
| `economy/credits/ledger.py` | 20 | Python | `transfer` |
| `economy/credits/ledger.py` | 26 | Python | `balance` |
| `economy/governance/policy_engine.py` | 6 | Python | `apply` |
| `economy/governance/proposal.py` | 6 | Python | `create` |
| `economy/governance/voting.py` | 6 | Python | `vote` |
| `economy/governance/voting.py` | 10 | Python | `result` |
| `economy/pricing_ai/model.py` | 7 | Python | `__init__` |
| `economy/pricing_ai/model.py` | 10 | Python | `price` |
| `economy/pricing_ai/model.py` | 14 | Python | `update` |
| `economy/pricing_ai/trainer.py` | 4 | Python | `train_step` |
| `economy/qos/controller.py` | 6 | Python | `priority` |
| `economy/qos/controller.py` | 11 | Python | `allocate` |
| `economy/reputation/sybil_guard.py` | 8 | Python | `register` |
| `endgame/competition/payoff.py` | 1 | Python | `payoff` |
| `endgame/competition/simulator.py` | 4 | Python | `step` |
| `endgame/competition/strategies.py` | 1 | Python | `strategy` |
| `endgame/governance/compliance_bridge.py` | 1 | Python | `validate` |
| `endgame/governance/dao.py` | 2 | Python | `__init__` |
| `endgame/governance/dao.py` | 6 | Python | `propose` |
| `endgame/governance/dao.py` | 9 | Python | `vote` |
| `endgame/governance/policy_engine.py` | 1 | Python | `enforce` |
| `endgame/governance/treasury.py` | 2 | Python | `__init__` |
| `endgame/governance/treasury.py` | 5 | Python | `allocate` |
| `endgame/governance/voting.py` | 1 | Python | `result` |
| `enterprise/dashboard/api.py` | 11 | Python | `dashboard` |
| `enterprise/dashboard/calculator.py` | 4 | Python | `compute_metrics` |
| `enterprise/onboarding/flow.py` | 4 | Python | `stage` |
| `enterprise/onboarding/pilot.py` | 4 | Python | `pilot_program` |
| `enterprise/sales/acquisition.py` | 6 | Python | `select_target` |
| `enterprise/sales/deal_closing.py` | 4 | Python | `close` |
| `enterprise/sales/outreach.py` | 4 | Python | `message` |
| `eu_tender/pricing/bid_optimizer.py` | 1 | Python | `optimize_bid` |
| `eu_tender/pricing/strategy.py` | 1 | Python | `tender_price` |
| `eu_tender/scoring/evaluator.py` | 1 | Python | `score` |
| `expansion/contracts.py` | 1 | Python | `expand_contract` |
| `expansion/multi_region.py` | 1 | Python | `expand_regions` |
| `expansion/upsell.py` | 1 | Python | `upsell` |
| `feature_store/offline.py` | 6 | Python | `insert` |
| `feature_store/online.py` | 8 | Python | `put` |
| `feature_store/online.py` | 12 | Python | `get` |
| `finance/export_model.py` | 10 | Python | `build_model` |
| `finance/export_model.py` | 35 | Python | `export_model` |
| `finance/model.py` | 15 | Python | `project` |
| `finance/model.py` | 25 | Python | `dcf` |
| `finance/model.py` | 29 | Python | `unit_economics` |
| `frontend/components/Card.tsx` | 8 | JS/TS | `Card` |
| `geo_conflicts/conflicts.py` | 1 | Python | `conflict` |
| `geo_conflicts/resolver.py` | 4 | Python | `resolve` |
| `gid_ref/conformance/runner.py` | 3 | Python | `run` |
| `gid_ref/sdk/python/issuer.py` | 1 | Python | `issue` |
| `gid_ref/sdk/python/jwks.py` | 4 | Python | `publish` |
| `gid_ref/sdk/python/jwks.py` | 8 | Python | `get` |
| `gid_ref/sdk/python/verifier.py` | 1 | Python | `verify` |
| `gid_ref/sdk/typescript/issuer.ts` | 1 | JS/TS | `issue` |
| `gid_ref/sdk/typescript/verifier.ts` | 1 | JS/TS | `verify` |
| `gid_standard/sdk/python/issuer.py` | 9 | Python | `__init__` |
| `gid_standard/sdk/python/issuer.py` | 16 | Python | `_encode` |
| `gid_standard/sdk/python/issuer.py` | 20 | Python | `mint` |
| `gid_standard/sdk/python/policy.py` | 5 | Python | `__init__` |
| `gid_standard/sdk/python/policy.py` | 8 | Python | `evaluate` |
| `gid_standard/sdk/python/verifier.py` | 8 | Python | `__init__` |
| `gid_standard/sdk/python/verifier.py` | 13 | Python | `_decode` |
| `gid_standard/sdk/python/verifier.py` | 17 | Python | `verify` |
| `gid_standard/sdk/typescript/issuer.ts` | 14 | JS/TS | `constructor` |
| `gid_standard/sdk/typescript/issuer.ts` | 18 | JS/TS | `mint` |
| `gid_standard/sdk/typescript/verifier.ts` | 6 | JS/TS | `verify` |
| `gid_standard/tests/test_policy.py` | 4 | Python | `test_cross_border_pii_denied_by_default` |
| `gid_standard/tests/test_policy.py` | 10 | Python | `test_same_region_pii_allowed` |
| `gid_standard/tests/test_tokens.py` | 10 | Python | `_enc` |
| `gid_standard/tests/test_tokens.py` | 15 | Python | `test_valid_token_passes` |
| `gid_standard/tests/test_tokens.py` | 21 | Python | `test_expired_rejected` |
| `gid_standard/tests/test_tokens.py` | 28 | Python | `test_untrusted_rejected` |
| `gid_standard/tests/test_tokens.py` | 35 | Python | `test_revoked_rejected` |
| `global_identity/core/issuer.py` | 7 | Python | `_b64` |
| `global_identity/core/issuer.py` | 11 | Python | `issue` |
| `global_identity/core/key_rotation.py` | 4 | Python | `rotate` |
| `global_identity/core/revocation.py` | 4 | Python | `revoke` |
| `global_identity/core/revocation.py` | 8 | Python | `is_revoked` |
| `global_identity/core/verifier.py` | 9 | Python | `_b64d` |
| `global_identity/core/verifier.py` | 14 | Python | `verify` |
| `global_identity/policy/access.py` | 1 | Python | `allow` |
| `global_identity/policy/risk.py` | 1 | Python | `score` |
| `global_identity/registry/attestations.py` | 4 | Python | `put` |
| `global_identity/registry/attestations.py` | 8 | Python | `get` |
| `global_identity/registry/issuers.py` | 4 | Python | `add` |
| `global_identity/registry/issuers.py` | 8 | Python | `is_trusted` |
| `global_identity/transport/channel_binding.py` | 1 | Python | `bind` |
| `global_identity/transport/mtls.py` | 1 | Python | `enforce` |
| `global_sovereign/routing/geo_policy.py` | 1 | Python | `route` |
| `global_sovereign/routing/residency.py` | 1 | Python | `allow` |
| `growth/analytics/tracker.py` | 4 | Python | `track` |
| `growth/funnel/optimizer.py` | 1 | Python | `conversion_rate` |
| `growth/funnel/optimizer.py` | 5 | Python | `optimize` |
| `growth/onboarding/engine.py` | 1 | Python | `onboarding_stage` |
| `growth_ai/bandits/thompson.py` | 7 | Python | `__init__` |
| `growth_ai/bandits/thompson.py` | 11 | Python | `select` |
| `growth_ai/bandits/thompson.py` | 15 | Python | `update` |
| `growth_ai/campaigns/manager.py` | 8 | Python | `run_campaign` |
| `growth_ai/llm/client.py` | 6 | Python | `generate` |
| `growth_ai/personalization/engine.py` | 7 | Python | `segment` |
| `growth_ai/personalization/engine.py` | 15 | Python | `message` |
| `hyperscale_1b/metrics.py` | 4 | Python | `metrics` |
| `infra/aws/billing.py` | 13 | Python | `cost_usage` |
| `infra/aws/billing.py` | 27 | Python | `current_month_to_date` |
| `infra/aws_client.py` | 6 | Python | `list_instances` |
| `infra/azure_client.py` | 7 | Python | `list_instances` |
| `infra/cost/controller.py` | 4 | Python | `check_cost` |
| `infra/cost/optimizer.py` | 1 | Python | `score` |
| `infra/cost/optimizer.py` | 9 | Python | `choose` |
| `infra/gcp_client.py` | 6 | Python | `list_instances` |
| `infra/identity/bind.py` | 4 | Python | `issue_spiffe_id` |
| `infra/policy/opa_client.py` | 6 | Python | `check_tenant` |
| `infra/policy/opa_client.py` | 12 | Python | `check_cost` |
| `infra/provisioning/engine.py` | 12 | Python | `create_instance` |
| `infra/recovery/auto.py` | 1 | Python | `recover` |
| `infra/tee/verify.py` | 9 | Python | `verify_quote` |
| `intelligence/causal/uplift.py` | 11 | Python | `__init__` |
| `intelligence/causal/uplift.py` | 15 | Python | `add` |
| `intelligence/causal/uplift.py` | 22 | Python | `uplift` |
| `intelligence/causal/uplift.py` | 28 | Python | `is_rollout_safe` |
| `intelligence/experimentation/ab.py` | 11 | Python | `assign` |
| `intelligence/experimentation/ab.py` | 17 | Python | `record` |
| `intelligence/experimentation/ab.py` | 24 | Python | `compare` |
| `intelligence/experimentation/shadow.py` | 8 | Python | `shadow_decision` |
| `intelligence/planner/budget.py` | 15 | Python | `can_spend` |
| `intelligence/planner/budget.py` | 22 | Python | `allocate` |
| `intelligence/planner/budget.py` | 31 | Python | `remaining` |
| `intelligence/realtime/feature_join.py` | 12 | Python | `_redis_client` |
| `intelligence/realtime/feature_join.py` | 17 | Python | `get_online` |
| `intelligence/realtime/feature_join.py` | 38 | Python | `join_features` |
| `ipo/finance.py` | 3 | Python | `unit_economics` |
| `ipo/finance.py` | 14 | Python | `burn_multiple` |
| `ipo/reporting.py` | 3 | Python | `quarterly` |
| `ipo_simulation/valuation.py` | 4 | Python | `valuation` |
| `ledger/adapters/bank_adapter.py` | 5 | Python | `charge` |
| `ledger/adapters/bank_adapter.py` | 8 | Python | `payout` |
| `ledger/adapters/base.py` | 6 | Python | `charge` |
| `ledger/adapters/base.py` | 10 | Python | `payout` |
| `ledger/adapters/stripe_adapter.py` | 10 | Python | `__init__` |
| `ledger/adapters/stripe_adapter.py` | 15 | Python | `charge` |
| `ledger/adapters/stripe_adapter.py` | 29 | Python | `payout` |
| `ledger/settlement_router.py` | 12 | Python | `charge` |
| `ledger/settlement_router.py` | 16 | Python | `payout` |
| `moat_simulation/metrics.py` | 1 | Python | `kpis` |
| `moat_simulation/model.py` | 1 | Python | `moat` |
| `monitoring/prometheus_client.py` | 4 | Python | `query_prometheus` |
| `monitoring/prometheus_client.py` | 10 | Python | `current_kpi_state` |

_Display capped at 300 of 1821 detected functions/methods._

### Automation pipeline files

- `.github/workflows/codeql.yml`
- `.github/workflows/dast.yml`
- `.github/workflows/deep_scan.yml`
- `.github/workflows/full-security-stack.yml`
- `.github/workflows/osv.yml`
- `.github/workflows/pip-audit.yml`
- `.github/workflows/provenance.yml`
- `.github/workflows/quality-gate.yml`
- `.github/workflows/sast.yml`
- `.github/workflows/sbom.yml`
- `.github/workflows/secrets.yml`
- `.github/workflows/security-dashboard.yml`
- `.github/workflows/security-gate.yml`
- `.github/workflows/sign.yml`
- `.github/workflows/sovereign-pipeline.yml`
- `decentralized/k8s/deploy_job.py`
- `enterprise/frontend/pages/DealsPipeline.tsx`
- `helm/clickhouse/install.sh`
- `helm/kafka/install.sh`
- `helm/redis/install.sh`
- `infra/docker-compose-bi.yml`
- `infra/k8s/deployment.yaml`
- `infra/platform/github-actions/enterprise-pipeline.yml`
- `k8s/base/deployment.yaml`
- `pilot/scripts/deploy.sh`
- `revenue/sales/pipeline.py`
- `saas/docker-compose.yml`
- `saas/infra/docker-compose.yml`
- `sales/pipeline.py`
- `sales_crm/pipeline.py`
- `scripts/zeaz-enterprise-install.sh`
- `scripts/zeaz-os-bootstrap.sh`
- `security/vault/install.sh`
- `tender_ai/pipeline/workflow.py`
- `v1/.github/workflows/build-deploy.yml`
- `v1/app/aiops/enterprise_pipeline.py`
- `v1/deploy/helm/zcyptobot/templates/deployment.yaml`
- `v1/devops/ci/github-actions.yml`
- `v1/docker-compose.yml`
- `v1/infra/aiops/cronjob.yaml`
- `v1/infra/ansible/bootstrap.yml`
- `v1/infra/ansible/install-k8s.yml`
- `v1/infra/ansible/install-monitoring.yml`
- `v1/infra/ansible/install-vault.yml`
- `v1/infra/k8s/zcyptobot/deployment.yaml`
- `v1/infra/scripts/bootstrap.sh`
- `v1/infra/scripts/deploy.sh`
- `v1/src/zcyptobot/pipeline.py`
- `v1/tests/test_pipeline.py`
- `zcyptobot-hyperscale/.github/workflows/cd.yml`
- `zcyptobot-hyperscale/.github/workflows/ci.yml`
- `zcyptobot-hyperscale/ai/training/pipeline.yaml`
- `zcyptobot-hyperscale/scripts/bootstrap.py`
- `zeaz-hyperscale/scripts/bootstrap_k3s.sh`
- `zeaz-hyperscale/scripts/deploy_cluster.sh`
- `zeaz-hyperscale/scripts/helm_install.sh`
- `zeaz-v12/helm/zeaz/templates/aiops-deployment.yaml`
- `zeaz-v12/helm/zeaz/templates/api-deployment.yaml`
- `zeaz/infra/docker-compose.yml`

## `cvsz/ZeaZDev-Omega`

Status: **cloned**; commit `d0e4ba7b48f5`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `apps/backend/src/modules/auth/auth.controller.ts` | 27 | `world-id/verify` |
| `apps/backend/src/modules/auth/auth.controller.ts` | 83 | `wallet/connect` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 14 | `swap/quote` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 29 | `stake` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 40 | `stake/:id/claim` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 45 | `stake/user/:userId` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 18 | `card/issue` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 35 | `card/user/:userId` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 40 | `bank/thai/deposit` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 55 | `bank/thai/withdraw` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 70 | `transactions/user/:userId` |
| `apps/backend/src/modules/game/game.controller.ts` | 14 | `slots/play` |
| `apps/backend/src/modules/game/game.controller.ts` | 31 | `slots/complete` |
| `apps/backend/src/modules/game/game.controller.ts` | 46 | `sessions/user/:userId` |
| `apps/backend/src/modules/game/game.controller.ts` | 51 | `sessions/:sessionId` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 14 | `claim` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 33 | `user/:userId` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 38 | `user/:userId/eligibility` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `apps/backend/src/main.ts` | 11 | JS/TS | `bootstrap` |
| `apps/backend/src/modules/auth/worldcoin.service.ts` | 15 | JS/TS | `constructor` |
| `apps/backend/src/modules/auth/worldcoin.service.ts` | 94 | JS/TS | `parseProof` |
| `apps/backend/src/modules/auth/worldcoin.service.ts` | 109 | JS/TS | `getNullifierHash` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 12 | JS/TS | `constructor` |
| `apps/backend/src/modules/defi/defi.service.ts` | 16 | JS/TS | `constructor` |
| `apps/backend/src/modules/defi/defi.service.ts` | 36 | JS/TS | `getSwapQuote` |
| `apps/backend/src/modules/defi/defi.service.ts` | 55 | JS/TS | `createStake` |
| `apps/backend/src/modules/defi/defi.service.ts` | 76 | JS/TS | `claimStakeRewards` |
| `apps/backend/src/modules/defi/defi.service.ts` | 112 | JS/TS | `getUserStakes` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 16 | JS/TS | `constructor` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 21 | JS/TS | `deposit` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 61 | JS/TS | `withdraw` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 116 | JS/TS | `getUserTransactions` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 139 | JS/TS | `verifyBankAccount` |
| `apps/backend/src/modules/fintech/card.service.ts` | 16 | JS/TS | `constructor` |
| `apps/backend/src/modules/fintech/card.service.ts` | 21 | JS/TS | `issueCard` |
| `apps/backend/src/modules/fintech/card.service.ts` | 69 | JS/TS | `getUserCard` |
| `apps/backend/src/modules/fintech/card.service.ts` | 91 | JS/TS | `generateMockCardNumber` |
| `apps/backend/src/modules/fintech/card.service.ts` | 98 | JS/TS | `maskCardNumber` |
| `apps/backend/src/modules/game/game.controller.ts` | 12 | JS/TS | `constructor` |
| `apps/backend/src/modules/game/game.service.ts` | 12 | JS/TS | `constructor` |
| `apps/backend/src/modules/game/game.service.ts` | 96 | JS/TS | `getUserSessions` |
| `apps/backend/src/modules/game/game.service.ts` | 119 | JS/TS | `getSession` |
| `apps/backend/src/modules/game/game.service.ts` | 144 | JS/TS | `generateSlotResult` |
| `apps/backend/src/modules/game/game.service.ts` | 168 | JS/TS | `generateSymbols` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 12 | JS/TS | `constructor` |
| `apps/backend/src/modules/rewards/rewards.service.ts` | 127 | JS/TS | `getUserRewards` |
| `apps/backend/src/modules/rewards/rewards.service.ts` | 144 | JS/TS | `checkEligibility` |
| `apps/backend/src/prisma.service.ts` | 12 | JS/TS | `onModuleInit` |
| `apps/backend/src/prisma.service.ts` | 16 | JS/TS | `onModuleDestroy` |
| `apps/frontend-miniapp/src/screens/RewardScreen.tsx` | 11 | JS/TS | `handleClaimDaily` |
| `apps/frontend-miniapp/src/screens/WorldIDScreen.tsx` | 17 | JS/TS | `handleVerifyWorldID` |
| `install.sh` | 30 | Shell | `print_success` |
| `install.sh` | 34 | Shell | `print_error` |
| `install.sh` | 38 | Shell | `print_warning` |
| `install.sh` | 42 | Shell | `print_info` |
| `packages/contracts/scripts/deploy.ts` | 9 | JS/TS | `main` |

### Automation pipeline files

- `INSTALLER_OS_REQUIREMENTS.md`
- `docker-compose.yml`
- `install.sh`
- `packages/contracts/scripts/deploy.ts`

## `cvsz/ztsaff`

Status: **cloned**; commit `f4e0e25f255d`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22550 | `/health` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22559 | `/release/info` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22574 | `/health` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22577 | `/wallet` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22582 | `/register` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22613 | `/login` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22683 | `/me` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22688 | `/user/dashboard` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22717 | `/wallet` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22722 | `/wallet/deposit` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22761 | `/wallet/transactions` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22770 | `/rent/plans` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22775 | `/rent/subscribe` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22801 | `/me/rentals` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22813 | `/generate` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22828 | `/my-scripts` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22833 | `/product-feed/import` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22870 | `/products` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22875 | `/video/generate-from-feed` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22912 | `/video-jobs` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22917 | `/showcase/upload` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22938 | `/showcase/uploads` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 22943 | `/admin/dashboard` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 23009 | `/admin/meta-dashboard` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 23013 | `/admin/master-meta-dashboard` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 23017 | `/admin/users` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 23022 | `/admin/rentals` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 23033 | `/admin/users/:id/role` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 24442 | `/health` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 25086 | `/health` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 25087 | `/health` |
| `ROOT_PROJECT_SOURCE_MERGED.md` | 25092 | `/webhook` |
| `tiktok-review-saas/backend/server.js` | 100 | `/health` |
| `tiktok-review-saas/backend/server.js` | 109 | `/release/info` |
| `tiktok-review-saas/backend/server.js` | 124 | `/health` |
| `tiktok-review-saas/backend/server.js` | 127 | `/wallet` |
| `tiktok-review-saas/backend/server.js` | 132 | `/register` |
| `tiktok-review-saas/backend/server.js` | 163 | `/login` |
| `tiktok-review-saas/backend/server.js` | 233 | `/me` |
| `tiktok-review-saas/backend/server.js` | 238 | `/user/dashboard` |
| `tiktok-review-saas/backend/server.js` | 267 | `/wallet` |
| `tiktok-review-saas/backend/server.js` | 272 | `/wallet/deposit` |
| `tiktok-review-saas/backend/server.js` | 311 | `/wallet/transactions` |
| `tiktok-review-saas/backend/server.js` | 320 | `/rent/plans` |
| `tiktok-review-saas/backend/server.js` | 325 | `/rent/subscribe` |
| `tiktok-review-saas/backend/server.js` | 351 | `/me/rentals` |
| `tiktok-review-saas/backend/server.js` | 363 | `/generate` |
| `tiktok-review-saas/backend/server.js` | 378 | `/my-scripts` |
| `tiktok-review-saas/backend/server.js` | 383 | `/product-feed/import` |
| `tiktok-review-saas/backend/server.js` | 420 | `/products` |
| `tiktok-review-saas/backend/server.js` | 425 | `/video/generate-from-feed` |
| `tiktok-review-saas/backend/server.js` | 462 | `/video-jobs` |
| `tiktok-review-saas/backend/server.js` | 467 | `/showcase/upload` |
| `tiktok-review-saas/backend/server.js` | 488 | `/showcase/uploads` |
| `tiktok-review-saas/backend/server.js` | 493 | `/admin/dashboard` |
| `tiktok-review-saas/backend/server.js` | 559 | `/admin/meta-dashboard` |
| `tiktok-review-saas/backend/server.js` | 563 | `/admin/master-meta-dashboard` |
| `tiktok-review-saas/backend/server.js` | 567 | `/admin/users` |
| `tiktok-review-saas/backend/server.js` | 572 | `/admin/rentals` |
| `tiktok-review-saas/backend/server.js` | 583 | `/admin/users/:id/role` |
| `zgitea-v8/edge-router/app.py` | 89 | `/health` |
| `zgitea-v8/webhook/ingress.py` | 34 | `/health` |
| `zgitea-v8/webhook/ingress.py` | 35 | `/health` |
| `zgitea-v8/webhook/ingress.py` | 40 | `/webhook` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `exports/gitea-plathform-clean/zgitea-installer.sh` | 30 | Shell | `install_gvisor` |
| `exports/gitea-plathform-clean/zgitea-installer.sh` | 38 | Shell | `test_gvisor` |
| `exports/gitea-plathform-clean/zgitea-installer.sh` | 52 | Shell | `gen_secret` |
| `exports/gitea-plathform/zgitea-installer__10.sh` | 47 | Shell | `install_gvisor` |
| `exports/gitea-plathform/zgitea-installer__10.sh` | 66 | Shell | `test_gvisor` |
| `exports/gitea-plathform/zgitea-installer__10.sh` | 81 | Shell | `init_secret` |
| `exports/gitea-plathform/zgitea-installer__10.sh` | 298 | Shell | `install_cron` |
| `exports/gitea-plathform/zgitea-installer__11.sh` | 31 | Shell | `gen_secret` |
| `exports/gitea-plathform/zgitea-installer__12.sh` | 29 | Shell | `install_gvisor` |
| `exports/gitea-plathform/zgitea-installer__12.sh` | 37 | Shell | `test_gvisor` |
| `exports/gitea-plathform/zgitea-installer__12.sh` | 51 | Shell | `gen_secret` |
| `exports/gitea-plathform/zgitea-installer__14.sh` | 29 | Shell | `install_gvisor` |
| `exports/gitea-plathform/zgitea-installer__14.sh` | 37 | Shell | `test_gvisor` |
| `exports/gitea-plathform/zgitea-installer__14.sh` | 51 | Shell | `gen_secret` |
| `exports/gitea-plathform/zgitea-installer__15.sh` | 30 | Shell | `install_gvisor` |
| `exports/gitea-plathform/zgitea-installer__15.sh` | 38 | Shell | `test_gvisor` |
| `exports/gitea-plathform/zgitea-installer__15.sh` | 52 | Shell | `gen_secret` |
| `exports/gitea-plathform/zgitea-installer__8.sh` | 47 | Shell | `install_gvisor` |
| `exports/gitea-plathform/zgitea-installer__8.sh` | 68 | Shell | `test_gvisor` |
| `exports/gitea-plathform/zgitea-installer__9.sh` | 47 | Shell | `install_gvisor` |
| `exports/gitea-plathform/zgitea-installer__9.sh` | 66 | Shell | `test_gvisor` |
| `exports/gitea-plathform/zgitea-installer__9.sh` | 238 | Shell | `install_cron` |
| `tiktok-review-saas/backend/server.js` | 40 | JS/TS | `authRateLimit` |
| `tiktok-review-saas/backend/server.js` | 58 | JS/TS | `normalizeEmail` |
| `tiktok-review-saas/backend/server.js` | 62 | JS/TS | `isValidEmail` |
| `tiktok-review-saas/backend/server.js` | 67 | JS/TS | `safeText` |
| `tiktok-review-saas/backend/server.js` | 71 | JS/TS | `parsePositiveInt` |
| `tiktok-review-saas/backend/server.js` | 81 | JS/TS | `asyncHandler` |
| `tiktok-review-saas/backend/server.js` | 85 | JS/TS | `ensureWallet` |
| `tiktok-review-saas/backend/server.js` | 185 | JS/TS | `auth` |
| `tiktok-review-saas/backend/server.js` | 201 | JS/TS | `adminOnly` |
| `tiktok-review-saas/backend/server.js` | 212 | JS/TS | `generateScript` |
| `tiktok-review-saas/backend/server.js` | 217 | JS/TS | `generateVideoPackage` |
| `tiktok-review-saas/backend/server.js` | 509 | JS/TS | `getMetaDashboardPayload` |
| `tiktok-review-saas/install.sh` | 18 | Shell | `run_docker_cmd` |
| `tiktok-review-saas/install.sh` | 26 | Shell | `install_docker` |
| `tiktok-review-saas/install.sh` | 48 | Shell | `install_compose_fallback` |
| `tiktok-review-saas/install.sh` | 64 | Shell | `generate_env_file` |
| `tiktok-review-saas/install.sh` | 92 | Shell | `start_stack` |
| `tiktok-review-saas/install.sh` | 96 | Shell | `main` |
| `tools/deep_scan_merge_to_root.py` | 33 | Python | `parse_args` |
| `tools/deep_scan_merge_to_root.py` | 56 | Python | `should_skip` |
| `tools/deep_scan_merge_to_root.py` | 64 | Python | `collect_source_files` |
| `tools/deep_scan_merge_to_root.py` | 77 | Python | `build_merged_document` |
| `tools/deep_scan_merge_to_root.py` | 108 | Python | `delete_files` |
| `tools/deep_scan_merge_to_root.py` | 117 | Python | `main` |
| `tools/merge_gitea_plathform.py` | 15 | Python | `parse_variant` |
| `tools/merge_gitea_plathform.py` | 32 | Python | `is_export_helper_file` |
| `tools/merge_gitea_plathform.py` | 36 | Python | `main` |
| `zgitea-v10-monorepo/apps/ai-scheduler/main.py` | 11 | Python | `predict_workers` |
| `zgitea-v10-monorepo/apps/control-plane/main.py` | 5 | Python | `elect_leader` |
| `zgitea-v10-monorepo/apps/control-plane/main.py` | 10 | Python | `main` |
| `zgitea-v10-monorepo/apps/edge-router/main.py` | 11 | Python | `do_POST` |
| `zgitea-v10-monorepo/apps/edge-router/main.py` | 34 | Python | `run` |
| `zgitea-v10-monorepo/apps/worker/main.py` | 5 | Python | `run_job` |
| `zgitea-v10-monorepo/apps/worker/main.py` | 12 | Python | `main` |
| `zgitea-v10-monorepo/scripts/sovereign-compliance-audit-v10.sh` | 16 | Shell | `die` |
| `zgitea-v10-monorepo/scripts/sovereign-compliance-audit-v10.sh` | 21 | Shell | `warn` |
| `zgitea-v10-monorepo/scripts/sovereign-compliance-audit-v10.sh` | 25 | Shell | `require_cmd` |
| `zgitea-v10-monorepo/scripts/sovereign-compliance-audit-v10.sh` | 29 | Shell | `require_optional_cmd` |
| `zgitea-v10-monorepo/scripts/sovereign-compliance-audit-v10.sh` | 45 | Shell | `quarantine_target` |
| `zgitea-v10-monorepo/tests/test_scheduler.py` | 6 | Python | `load_module` |
| `zgitea-v10-monorepo/tests/test_scheduler.py` | 14 | Python | `test_predict_workers_increases_for_load` |
| `zgitea-v8/autoscaler/autoscaler.py` | 20 | Python | `queue_depth` |
| `zgitea-v8/autoscaler/autoscaler.py` | 29 | Python | `desired_workers` |
| `zgitea-v8/autoscaler/autoscaler.py` | 33 | Python | `scale_to` |
| `zgitea-v8/autoscaler/autoscaler.py` | 43 | Python | `run` |
| `zgitea-v8/edge-router/app.py` | 32 | Python | `_distance_score` |
| `zgitea-v8/edge-router/app.py` | 36 | Python | `_geoip_hint` |
| `zgitea-v8/edge-router/app.py` | 47 | Python | `_probe` |
| `zgitea-v8/edge-router/app.py` | 70 | Python | `choose_region` |
| `zgitea-v8/edge-router/app.py` | 90 | Python | `health` |
| `zgitea-v8/edge-router/app.py` | 98 | Python | `route` |
| `zgitea-v8/replication-bridge/bridge.py` | 22 | Python | `stream_names` |
| `zgitea-v8/replication-bridge/bridge.py` | 26 | Python | `ensure_groups` |
| `zgitea-v8/replication-bridge/bridge.py` | 37 | Python | `replication_id` |
| `zgitea-v8/replication-bridge/bridge.py` | 42 | Python | `fanout` |
| `zgitea-v8/replication-bridge/bridge.py` | 61 | Python | `poll_region` |
| `zgitea-v8/replication-bridge/bridge.py` | 77 | Python | `reclaim_stale` |
| `zgitea-v8/replication-bridge/bridge.py` | 89 | Python | `main` |
| `zgitea-v8/scheduler/scheduler.py` | 23 | Python | `stream_for` |
| `zgitea-v8/scheduler/scheduler.py` | 29 | Python | `region_leader_lock` |
| `zgitea-v8/scheduler/scheduler.py` | 41 | Python | `ensure_groups` |
| `zgitea-v8/scheduler/scheduler.py` | 51 | Python | `dispatch` |
| `zgitea-v8/scheduler/scheduler.py` | 58 | Python | `reconcile_stale` |
| `zgitea-v8/scheduler/scheduler.py` | 66 | Python | `run_once` |
| `zgitea-v8/scheduler/scheduler.py` | 77 | Python | `main` |
| `zgitea-v8/webhook/ingress.py` | 19 | Python | `read_secret` |
| `zgitea-v8/webhook/ingress.py` | 26 | Python | `verify` |
| `zgitea-v8/webhook/ingress.py` | 36 | Python | `health` |
| `zgitea-v8/webhook/ingress.py` | 41 | Python | `webhook` |
| `zgitea-v8/worker/runner.py` | 24 | Python | `read_secret` |
| `zgitea-v8/worker/runner.py` | 31 | Python | `ensure_groups` |
| `zgitea-v8/worker/runner.py` | 41 | Python | `run_wasm` |
| `zgitea-v8/worker/runner.py` | 49 | Python | `encrypt_blob` |
| `zgitea-v8/worker/runner.py` | 61 | Python | `process_job` |
| `zgitea-v8/worker/runner.py` | 83 | Python | `ack_result` |
| `zgitea-v8/worker/runner.py` | 89 | Python | `run` |
| `zt-gitea-v10/install-v10.sh` | 10 | Shell | `require_cmd` |
| `zt-gitea-v10/install-v10.sh` | 17 | Shell | `log` |
| `zt-gitea-v10/scripts/audit-runtime.sh` | 4 | Shell | `say` |
| `zt-gitea-v10/scripts/clean-os.sh` | 4 | Shell | `log` |
| `zt-gitea-v10/scripts/sovereign-compliance-audit-v10.sh` | 16 | Shell | `die` |
| `zt-gitea-v10/scripts/sovereign-compliance-audit-v10.sh` | 21 | Shell | `warn` |
| `zt-gitea-v10/scripts/sovereign-compliance-audit-v10.sh` | 25 | Shell | `require_cmd` |
| `zt-gitea-v10/scripts/sovereign-compliance-audit-v10.sh` | 29 | Shell | `require_optional_cmd` |
| `zt-gitea-v10/scripts/sovereign-compliance-audit-v10.sh` | 45 | Shell | `quarantine_target` |

### Automation pipeline files

- `exports/gitea-plathform-clean/.github/workflows/ci.yml`
- `exports/gitea-plathform-clean/.github/workflows/deploy.yml`
- `exports/gitea-plathform-clean/docker-compose.bluegreen.yml`
- `exports/gitea-plathform-clean/docker-compose.ephemeral-v5.yml`
- `exports/gitea-plathform-clean/docker-compose.ephemeral-v6.yml`
- `exports/gitea-plathform-clean/docker-compose.ephemeral-v7.yml`
- `exports/gitea-plathform-clean/docker-compose.yml`
- `exports/gitea-plathform-clean/fix-cron-safe.sh`
- `exports/gitea-plathform-clean/install-docker-clean.sh`
- `exports/gitea-plathform-clean/install-v10.sh`
- `exports/gitea-plathform-clean/regions/region-a/docker-compose.yml`
- `exports/gitea-plathform-clean/scripts/deploy-bluegreen.sh`
- `exports/gitea-plathform-clean/scripts/install.sh`
- `exports/gitea-plathform-clean/zgitea-installer.sh`
- `exports/gitea-plathform/.github/workflows/ci.yml`
- `exports/gitea-plathform/.github/workflows/deploy.yml`
- `exports/gitea-plathform/.github/workflows/deploy__2.yml`
- `exports/gitea-plathform/docker-compose.bluegreen.yml`
- `exports/gitea-plathform/docker-compose.ephemeral-v5.yml`
- `exports/gitea-plathform/docker-compose.ephemeral-v6.yml`
- `exports/gitea-plathform/docker-compose.ephemeral-v7.yml`
- `exports/gitea-plathform/docker-compose.yml`
- `exports/gitea-plathform/docker-compose__2.yml`
- `exports/gitea-plathform/docker-compose__3.yml`
- `exports/gitea-plathform/docker-compose__4.yml`
- `exports/gitea-plathform/fix-cron-safe.sh`
- `exports/gitea-plathform/install-docker-clean.sh`
- `exports/gitea-plathform/install-v10.sh`
- `exports/gitea-plathform/regions/region-a/docker-compose.yml`
- `exports/gitea-plathform/scripts/deploy-bluegreen.sh`
- `exports/gitea-plathform/scripts/install.sh`
- `exports/gitea-plathform/zgitea-installer.sh`
- `exports/gitea-plathform/zgitea-installer__10.sh`
- `exports/gitea-plathform/zgitea-installer__11.sh`
- `exports/gitea-plathform/zgitea-installer__12.sh`
- `exports/gitea-plathform/zgitea-installer__13.sh`
- `exports/gitea-plathform/zgitea-installer__14.sh`
- `exports/gitea-plathform/zgitea-installer__15.sh`
- `exports/gitea-plathform/zgitea-installer__2.sh`
- `exports/gitea-plathform/zgitea-installer__3.sh`
- `exports/gitea-plathform/zgitea-installer__4.sh`
- `exports/gitea-plathform/zgitea-installer__5.sh`
- `exports/gitea-plathform/zgitea-installer__6.sh`
- `exports/gitea-plathform/zgitea-installer__7.sh`
- `exports/gitea-plathform/zgitea-installer__8.sh`
- `exports/gitea-plathform/zgitea-installer__9.sh`
- `tiktok-review-saas/docker-compose.yml`
- `tiktok-review-saas/install.sh`
- `zgitea-v10-monorepo/.github/workflows/ci.yml`
- `zgitea-v10-monorepo/deploy/docker-compose.yml`
- `zgitea-v10-monorepo/scripts/bootstrap.sh`
- `zgitea-v10-monorepo/zgitea-installer.sh`
- `zgitea-v8/edge-router/docker-compose.yml`
- `zgitea-v8/regions/region-a/docker-compose.yml`
- `zgitea-v8/regions/region-b/docker-compose.yml`
- `zgitea-v8/regions/region-c/docker-compose.yml`
- `zgitea-v8/scripts/bootstrap.sh`
- `zt-gitea-v10/install-v10.sh`

## `cvsz/zLinebot`

Status: **cloned**; commit `0c0c3ec4d7d5`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `admin/src/pages/Billing.jsx` | 7 | `/api/admin/billing` |
| `admin/src/pages/TikTokShopPanel.jsx` | 31 | `/api/admin/tiktok-shop/overview` |
| `admin/src/pages/TikTokShopPanel.jsx` | 43 | `/api/admin/tiktok-shop/sync-showcase` |
| `admin/src/pages/TikTokShopPanel.jsx` | 56 | `/api/admin/tiktok-shop/intelligence` |
| `admin/src/pages/TikTokShopPanel.jsx` | 68 | `/api/admin/tiktok-shop/export` |
| `admin/src/pages/TikTokShopPanel.jsx` | 86 | `/api/admin/tiktok-shop/auto-video` |
| `app/src/line/webhook.ts` | 43 | `/line` |
| `app/src/main.ts` | 78 | `/webhook` |
| `app/src/main.ts` | 79 | `/webhook` |
| `app/src/main.ts` | 83 | `/health` |
| `app/src/main.ts` | 85 | `/` |
| `app/src/main.ts` | 86 | `/` |
| `app/src/main.ts` | 100 | `/` |
| `app/src/routes/admin.billing.ts` | 8 | `/admin/billing` |
| `app/src/routes/feedback.ts` | 15 | `/api/feedback` |
| `app/src/routes/tiktok.ts` | 33 | `/auth/tiktok/url` |
| `app/src/routes/tiktok.ts` | 54 | `/auth/tiktok/callback` |
| `app/src/routes/tiktok.ts` | 86 | `/webhook/tiktok` |
| `app/src/routes/webhook.promptpay.ts` | 6 | `/promptpay` |
| `app/src/routes/webhook.stripe.ts` | 21 | `/stripe` |
| `app/src/services/rlhf.ts` | 60 | `/api/feedback` |
| `apps/api/src/routes/analytics.ts` | 4 | `/stats` |
| `apps/api/src/routes/auth.ts` | 6 | `/tiktok/callback` |
| `apps/api/src/routes/automation.ts` | 5 | `/` |
| `apps/api/src/routes/health.ts` | 4 | `/health` |
| `apps/api/src/routes/logs.ts` | 4 | `/` |
| `apps/api/src/routes/stripeWebhook.ts` | 11 | `/stripe/webhook` |
| `apps/api/src/routes/webhook.ts` | 337 | `/tiktok` |
| `apps/api/src/routes/webhook.ts` | 399 | `/line` |
| `apps/api/src/server.ts` | 35 | `/metrics` |
| `apps/api/src/server.ts` | 40 | `/auth` |
| `apps/api/src/server.ts` | 45 | `/webhook` |
| `cloud/worker/main.py` | 11 | `/infer` |
| `k8s/api-deployment.yaml` | 18 | `/metrics` |
| `k8s/deploy.yaml` | 31 | `/metrics` |
| `k8s/zlinebot-config.yaml` | 10 | `/webhook` |
| `mobile/api.ts` | 2 | `/api/admin/retrain` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `admin/src/pages/Automations.jsx` | 7 | JS/TS | `addAction` |
| `admin/src/pages/Automations.jsx` | 11 | JS/TS | `saveAutomation` |
| `admin/src/pages/Builder.jsx` | 34 | JS/TS | `saveFlow` |
| `admin/src/pages/TikTokShopPanel.jsx` | 30 | JS/TS | `loadOverview` |
| `admin/src/pages/TikTokShopPanel.jsx` | 40 | JS/TS | `syncShowcase` |
| `admin/src/pages/TikTokShopPanel.jsx` | 53 | JS/TS | `loadIntelligence` |
| `admin/src/pages/TikTokShopPanel.jsx` | 65 | JS/TS | `downloadExcel` |
| `admin/src/pages/TikTokShopPanel.jsx` | 83 | JS/TS | `generateVideos` |
| `admin/src/pages/TikTokShopPanel.jsx` | 101 | JS/TS | `toggleProduct` |
| `app/src/agents/actuators.ts` | 8 | JS/TS | `act` |
| `app/src/agents/auction.ts` | 1 | JS/TS | `scoreBid` |
| `app/src/agents/automl.ts` | 15 | JS/TS | `runAutoML` |
| `app/src/agents/automl.ts` | 33 | JS/TS | `evalModel` |
| `app/src/agents/automl.ts` | 44 | JS/TS | `threshold` |
| `app/src/agents/autonomous.ts` | 16 | JS/TS | `tick` |
| `app/src/agents/economy.ts` | 13 | JS/TS | `step` |
| `app/src/agents/economy.ts` | 30 | JS/TS | `marginAfterGlobal` |
| `app/src/agents/guardrails.ts` | 7 | JS/TS | `guard` |
| `app/src/agents/killswitch.ts` | 6 | JS/TS | `shouldHalt` |
| `app/src/agents/ledger.ts` | 1 | JS/TS | `record` |
| `app/src/agents/market.ts` | 7 | JS/TS | `selectAgent` |
| `app/src/agents/modelManager.ts` | 16 | JS/TS | `evaluateCandidate` |
| `app/src/agents/modelManager.ts` | 37 | JS/TS | `deployIfBetter` |
| `app/src/agents/negotiation.basic.ts` | 8 | JS/TS | `select` |
| `app/src/agents/negotiation.ts` | 20 | JS/TS | `negotiate` |
| `app/src/agents/policy.ts` | 32 | JS/TS | `decide` |
| `app/src/agents/pricing.ts` | 3 | JS/TS | `pricingAgent` |
| `app/src/agents/profit.ts` | 1 | JS/TS | `profit` |
| `app/src/agents/registry.ts` | 3 | JS/TS | `register` |
| `app/src/agents/registry.ts` | 7 | JS/TS | `run` |
| `app/src/agents/revenue.ts` | 1 | JS/TS | `payout` |
| `app/src/agents/risk.ts` | 3 | JS/TS | `riskAgent` |
| `app/src/agents/sales.ts` | 3 | JS/TS | `salesAgent` |
| `app/src/agents/sensors.ts` | 3 | JS/TS | `readState` |
| `app/src/agents/sim.ts` | 5 | JS/TS | `simulate` |
| `app/src/agents/supply.ts` | 3 | JS/TS | `supplyAgent` |
| `app/src/agi/recommender.ts` | 5 | JS/TS | `retrieve` |
| `app/src/agi/recommender.ts` | 9 | JS/TS | `rankBatch` |
| `app/src/agi/recommender.ts` | 15 | JS/TS | `toFeatures` |
| `app/src/agi/recommender.ts` | 19 | JS/TS | `recommend` |
| `app/src/ai/edge.ts` | 1 | JS/TS | `edgeAI` |
| `app/src/automation/ai.ts` | 10 | JS/TS | `generateReply` |
| `app/src/automation/compiler.ts` | 5 | JS/TS | `compileFlow` |
| `app/src/automation/plugins.ts` | 5 | JS/TS | `registerPlugin` |
| `app/src/automation/plugins.ts` | 9 | JS/TS | `runPlugin` |
| `app/src/causal/intervene.ts` | 3 | JS/TS | `doIntervene` |
| `app/src/causal/notears_like.ts` | 1 | JS/TS | `h` |
| `app/src/causal/notears_like.ts` | 13 | JS/TS | `score` |
| `app/src/causal/pc_lite.ts` | 3 | JS/TS | `buildGraph` |
| `app/src/compliance/evidence.ts` | 4 | JS/TS | `collectEvidence` |
| `app/src/compliance/generator.ts` | 7 | JS/TS | `generateReport` |
| `app/src/dr/failover.ts` | 1 | JS/TS | `shouldFailover` |
| `app/src/dr/health.ts` | 1 | JS/TS | `checkPrimary` |
| `app/src/dr/router.ts` | 3 | JS/TS | `routeRequest` |
| `app/src/econ/global.ts` | 3 | JS/TS | `globalObjective` |
| `app/src/econ/global.ts` | 8 | JS/TS | `acceptChange` |
| `app/src/econ/ledger.ts` | 4 | JS/TS | `credit` |
| `app/src/econ/payout.ts` | 1 | JS/TS | `payoutOnChain` |
| `app/src/events/producer.ts` | 7 | JS/TS | `ensureConnected` |
| `app/src/events/producer.ts` | 13 | JS/TS | `emit` |
| `app/src/federation/aggregate.ts` | 1 | JS/TS | `aggregate` |
| `app/src/federation/gossip.ts` | 5 | JS/TS | `getPeers` |
| `app/src/federation/gossip.ts` | 13 | JS/TS | `broadcastUpdate` |
| `app/src/health.ts` | 3 | JS/TS | `health` |
| `app/src/identity/federated.basic.ts` | 3 | JS/TS | `hashId` |
| `app/src/identity/federated.basic.ts` | 7 | JS/TS | `shareLink` |
| `app/src/identity/federated.ts` | 6 | JS/TS | `federatedResolve` |
| `app/src/identity/link.ts` | 3 | JS/TS | `link` |
| `app/src/identity/resolve.ts` | 1 | JS/TS | `resolveIdentity` |
| `app/src/line/flex.ts` | 16 | JS/TS | `productFlex` |
| `app/src/line/flex.ts` | 52 | JS/TS | `buildRecommendationFlex` |
| `app/src/line/flex.ts` | 75 | JS/TS | `buildAgentActionFlex` |
| `app/src/line/flex.ts` | 98 | JS/TS | `recommendationFlex` |
| `app/src/line/flex.ts` | 102 | JS/TS | `agentActionText` |
| `app/src/line/flex.ts` | 114 | JS/TS | `textMessage` |
| `app/src/line/handler.ts` | 13 | JS/TS | `emitCoordinatedRewardEvent` |
| `app/src/line/handler.ts` | 27 | JS/TS | `handleMessage` |
| `app/src/line/handler.ts` | 60 | JS/TS | `handlePostback` |
| `app/src/line/webhook.ts` | 93 | JS/TS | `isDuplicateLineEvent` |
| `app/src/line/webhook.ts` | 99 | JS/TS | `handleEvent` |
| `app/src/line/webhook.ts` | 123 | JS/TS | `sendReply` |
| `app/src/main.ts` | 57 | JS/TS | `origin` |
| `app/src/main.ts` | 148 | JS/TS | `shutdown` |
| `app/src/middleware/errorHandler.ts` | 3 | JS/TS | `errorHandler` |
| `app/src/middleware/rateLimit.ts` | 10 | JS/TS | `buildRateLimit` |
| `app/src/middleware/rateLimit.ts` | 26 | JS/TS | `routeRateLimit` |
| `app/src/middleware/rateLimit.ts` | 30 | JS/TS | `rateLimitByIp` |
| `app/src/middleware/riskGuard.ts` | 5 | JS/TS | `riskGuard` |
| `app/src/middleware/schema.ts` | 5 | JS/TS | `safeTenantSchema` |
| `app/src/middleware/schema.ts` | 11 | JS/TS | `setTenantSchema` |
| `app/src/middleware/tenant.ts` | 8 | JS/TS | `readHeader` |
| `app/src/middleware/tenant.ts` | 12 | JS/TS | `isAuthorizedTenantKey` |
| `app/src/middleware/tenant.ts` | 16 | JS/TS | `resolveTenantId` |
| `app/src/middleware/tenant.ts` | 21 | JS/TS | `tenant` |
| `app/src/middleware/trace.ts` | 6 | JS/TS | `trace` |
| `app/src/middleware/zeroTrust.ts` | 4 | JS/TS | `verify` |
| `app/src/models/lifecycle.ts` | 15 | JS/TS | `evolve` |
| `app/src/org/controlPlane.ts` | 7 | JS/TS | `controlTick` |
| `app/src/org/executor.ts` | 3 | JS/TS | `execute` |
| `app/src/org/guardrails.ts` | 5 | JS/TS | `enforce` |
| `app/src/org/planner.ts` | 7 | JS/TS | `plan` |
| `app/src/org/provision.ts` | 3 | JS/TS | `createCluster` |
| `app/src/org/router.ts` | 1 | JS/TS | `routeOrg` |
| `app/src/privacy/confidential.basic.ts` | 6 | JS/TS | `encrypt` |
| `app/src/privacy/confidential.basic.ts` | 14 | JS/TS | `decrypt` |
| `app/src/privacy/filters.ts` | 8 | JS/TS | `minimize` |
| `app/src/queue/automation.ts` | 19 | JS/TS | `enqueueAutomationJob` |
| `app/src/queue/automation.ts` | 30 | JS/TS | `parseCompiledSteps` |
| `app/src/queue/automation.ts` | 39 | JS/TS | `matchesCondition` |
| `app/src/queue/automation.ts` | 48 | JS/TS | `executeCompiledFlow` |
| `app/src/queue/automation.ts` | 69 | JS/TS | `processAutomation` |
| `app/src/queue/automation.ts` | 90 | JS/TS | `startAutomationWorker` |
| `app/src/queue/producer.ts` | 5 | JS/TS | `pushTask` |
| `app/src/rl/counterfactual.ts` | 1 | JS/TS | `uplift` |
| `app/src/rl/dqn.ts` | 33 | JS/TS | `randomMatrix` |
| `app/src/rl/dqn.ts` | 39 | JS/TS | `dot` |
| `app/src/rl/dqn.ts` | 57 | JS/TS | `maxIndex` |
| `app/src/rl/dqn.ts` | 67 | JS/TS | `samplePrioritizedBatch` |
| `app/src/rl/dqn.ts` | 87 | JS/TS | `encodeState` |
| `app/src/rl/dqn.ts` | 122 | JS/TS | `constructor` |
| `app/src/rl/dqn.ts` | 140 | JS/TS | `getEpsilon` |
| `app/src/rl/dqn.ts` | 144 | JS/TS | `selectAction` |
| `app/src/rl/dqn.ts` | 151 | JS/TS | `storeTransition` |
| `app/src/rl/dqn.ts` | 164 | JS/TS | `trainStep` |
| `app/src/rl/dqn.ts` | 200 | JS/TS | `toAgentAction` |
| `app/src/rl/dqn.ts` | 215 | JS/TS | `softUpdateTarget` |
| `app/src/rl/dqn.ts` | 232 | JS/TS | `configureDQN` |
| `app/src/rl/dqn.ts` | 237 | JS/TS | `getDQN` |
| `app/src/rl/dr.ts` | 1 | JS/TS | `dr` |
| `app/src/rl/env.ts` | 14 | JS/TS | `createStaticEnvironment` |
| `app/src/rl/env.ts` | 16 | JS/TS | `reset` |
| `app/src/rl/env.ts` | 19 | JS/TS | `step` |
| `app/src/rl/ips.ts` | 1 | JS/TS | `ips` |
| `app/src/rl/multi-agent-reward.ts` | 49 | JS/TS | `setMultiAgentRewardConfig` |
| `app/src/rl/multi-agent-reward.ts` | 55 | JS/TS | `computeCoordinatedReward` |
| `app/src/rl/multi-agent-reward.ts` | 114 | JS/TS | `initializeMultiAgentRewardSystem` |
| `app/src/rl/policy.ts` | 19 | JS/TS | `hasDQN` |
| `app/src/rl/policy.ts` | 39 | JS/TS | `dqnActionIndex` |
| `app/src/rl/policy.ts` | 59 | JS/TS | `getRLAction` |
| `app/src/rl/policy.ts` | 79 | JS/TS | `observeTransition` |
| `app/src/rl/qlearning.ts` | 18 | JS/TS | `getStateKey` |
| `app/src/rl/qlearning.ts` | 24 | JS/TS | `ensureState` |
| `app/src/rl/qlearning.ts` | 39 | JS/TS | `bestActionKey` |
| `app/src/rl/qlearning.ts` | 45 | JS/TS | `toAgentAction` |
| `app/src/rl/qlearning.ts` | 54 | JS/TS | `selectAction` |
| `app/src/rl/qlearning.ts` | 68 | JS/TS | `updateQTable` |
| `app/src/rl/qlearning.ts` | 86 | JS/TS | `snapshotQTable` |
| `app/src/rl/reward.ts` | 27 | JS/TS | `setRewardConfig` |
| `app/src/rl/reward.ts` | 33 | JS/TS | `initializeRewardSystem` |
| `app/src/rl/reward.ts` | 41 | JS/TS | `computeReward` |
| `app/src/rl/train.ts` | 6 | JS/TS | `trainEpisode` |
| `app/src/rl/utils.ts` | 1 | JS/TS | `argMax` |
| `app/src/rl/utils.ts` | 18 | JS/TS | `epsilonGreedy` |
| `app/src/rl/utils.ts` | 22 | JS/TS | `clamp` |
| `app/src/routes/webhook.stripe.ts` | 9 | JS/TS | `getStripeClient` |
| `app/src/security/basic.ts` | 3 | JS/TS | `verify` |
| `app/src/security/encryption.ts` | 4 | JS/TS | `loadEncryptionKey` |
| `app/src/security/encryption.ts` | 17 | JS/TS | `encrypt` |
| `app/src/security/secret-validator.ts` | 6 | JS/TS | `hasStrongComposition` |
| `app/src/security/signature.ts` | 3 | JS/TS | `timingSafeEqualStrings` |
| `app/src/security/signature.ts` | 11 | JS/TS | `verifyLineSignature` |
| `app/src/security/signature.ts` | 26 | JS/TS | `verifyTikTokSignature` |
| `app/src/security/tee.ts` | 1 | JS/TS | `secureCompute` |
| `app/src/sensors/kpi.ts` | 12 | JS/TS | `readKPIs` |
| `app/src/services/ab.metrics.ts` | 13 | JS/TS | `trackAbEvent` |
| `app/src/services/ab.ts` | 11 | JS/TS | `getBucket` |
| `app/src/services/ab.ts` | 16 | JS/TS | `assignExperiment` |
| `app/src/services/affiliate.ts` | 4 | JS/TS | `applyAffiliate` |
| `app/src/services/agent.ts` | 5 | JS/TS | `salesAgent` |
| `app/src/services/agi.ts` | 11 | JS/TS | `agiRecommend` |
| `app/src/services/ai.router.ts` | 4 | JS/TS | `localInfer` |
| `app/src/services/ai.router.ts` | 9 | JS/TS | `inferSmart` |
| `app/src/services/ai.ts` | 1 | JS/TS | `generateReply` |
| `app/src/services/alerts.ts` | 6 | JS/TS | `checkAlerts` |
| `app/src/services/analytics.ts` | 17 | JS/TS | `emitEvent` |
| `app/src/services/analytics.ts` | 29 | JS/TS | `startAggregator` |
| `app/src/services/analytics.ts` | 48 | JS/TS | `snapshot` |
| `app/src/services/arbitrage.exec.ts` | 3 | JS/TS | `reserve` |
| `app/src/services/arbitrage.exec.ts` | 7 | JS/TS | `placeSell` |
| `app/src/services/arbitrage.exec.ts` | 11 | JS/TS | `placeBuy` |
| `app/src/services/arbitrage.exec.ts` | 15 | JS/TS | `assertOpportunity` |
| `app/src/services/arbitrage.exec.ts` | 25 | JS/TS | `executeArb` |
| `app/src/services/arbitrage.ts` | 20 | JS/TS | `findArb` |
| `app/src/services/audit.ts` | 15 | JS/TS | `lastHash` |
| `app/src/services/audit.ts` | 23 | JS/TS | `auditLog` |
| `app/src/services/auth.ts` | 7 | JS/TS | `sign` |
| `app/src/services/auth.ts` | 21 | JS/TS | `verifyToken` |
| `app/src/services/auth.ts` | 44 | JS/TS | `authorize` |
| `app/src/services/automation.ts` | 13 | JS/TS | `createAutomationRule` |
| `app/src/services/automation.ts` | 25 | JS/TS | `createFlowAutomationRule` |
| `app/src/services/automation.ts` | 38 | JS/TS | `listAutomationRules` |
| `app/src/services/automation.ts` | 50 | JS/TS | `matchesCondition` |
| `app/src/services/automation.ts` | 65 | JS/TS | `triggerAutomations` |
| `app/src/services/bandit.causal.ts` | 8 | JS/TS | `causalScore` |
| `app/src/services/bandit.causal.ts` | 20 | JS/TS | `meanCausalScore` |
| `app/src/services/bandit.contextual.ts` | 19 | JS/TS | `key` |
| `app/src/services/bandit.contextual.ts` | 21 | JS/TS | `identity` |
| `app/src/services/bandit.contextual.ts` | 27 | JS/TS | `dot` |
| `app/src/services/bandit.contextual.ts` | 35 | JS/TS | `matVec` |
| `app/src/services/bandit.contextual.ts` | 39 | JS/TS | `addVec` |
| `app/src/services/bandit.contextual.ts` | 43 | JS/TS | `scaleVec` |
| `app/src/services/bandit.contextual.ts` | 47 | JS/TS | `normalize` |
| `app/src/services/bandit.contextual.ts` | 52 | JS/TS | `getArm` |
| `app/src/services/bandit.contextual.ts` | 64 | JS/TS | `setArm` |
| `app/src/services/bandit.contextual.ts` | 68 | JS/TS | `shermanMorrison` |
| `app/src/services/bandit.contextual.ts` | 83 | JS/TS | `selectContextual` |
| `app/src/services/bandit.contextual.ts` | 123 | JS/TS | `updateContextual` |
| `app/src/services/bandit.ts` | 12 | JS/TS | `key` |
| `app/src/services/bandit.ts` | 14 | JS/TS | `getArm` |
| `app/src/services/bandit.ts` | 19 | JS/TS | `setArm` |
| `app/src/services/bandit.ts` | 23 | JS/TS | `sampleBeta` |
| `app/src/services/bandit.ts` | 29 | JS/TS | `selectArm` |
| `app/src/services/bandit.ts` | 52 | JS/TS | `updateArm` |
| `app/src/services/billing.ts` | 12 | JS/TS | `createInvoice` |
| `app/src/services/billing.ts` | 24 | JS/TS | `markPaid` |
| `app/src/services/billing.ts` | 28 | JS/TS | `setSubscriptionPlan` |
| `app/src/services/billing.ts` | 38 | JS/TS | `getPlanLimits` |
| `app/src/services/budget.ts` | 3 | JS/TS | `setBudget` |
| `app/src/services/budget.ts` | 7 | JS/TS | `spend` |
| `app/src/services/budget.ts` | 14 | JS/TS | `remainingBudget` |
| `app/src/services/causal.online.ts` | 3 | JS/TS | `uplift` |
| `app/src/services/causal.stream.ts` | 6 | JS/TS | `startCausalStreamConsumer` |
| `app/src/services/crypto.ts` | 9 | JS/TS | `getKey` |
| `app/src/services/crypto.ts` | 23 | JS/TS | `encrypt` |
| `app/src/services/crypto.ts` | 34 | JS/TS | `decrypt` |
| `app/src/services/diffusion.ts` | 5 | JS/TS | `generateEmbedding` |
| `app/src/services/econ.ts` | 8 | JS/TS | `objective` |
| `app/src/services/embed.ts` | 1 | JS/TS | `embed` |
| `app/src/services/eval.dr.ts` | 8 | JS/TS | `dr` |
| `app/src/services/eval.ips.ts` | 7 | JS/TS | `ips` |
| `app/src/services/events.reward.ts` | 3 | JS/TS | `onClick` |
| `app/src/services/events.reward.ts` | 7 | JS/TS | `onView` |
| `app/src/services/explain.ts` | 6 | JS/TS | `explain` |
| `app/src/services/export.ts` | 14 | JS/TS | `escapeCsv` |
| `app/src/services/export.ts` | 20 | JS/TS | `exportLedger` |
| `app/src/services/fallback.ts` | 3 | JS/TS | `safeInfer` |
| `app/src/services/feature.online.ts` | 5 | JS/TS | `getFeatures` |
| `app/src/services/feature.sync.ts` | 20 | JS/TS | `parseFeatureEvent` |
| `app/src/services/feature.sync.ts` | 37 | JS/TS | `startFeatureSyncConsumer` |
| `app/src/services/featureRegistry.ts` | 18 | JS/TS | `registerFeature` |
| `app/src/services/featureRegistry.ts` | 28 | JS/TS | `registerLineage` |
| `app/src/services/featureStore.ts` | 16 | JS/TS | `userKey` |
| `app/src/services/featureStore.ts` | 17 | JS/TS | `itemKey` |
| `app/src/services/featureStore.ts` | 19 | JS/TS | `wrapFeature` |
| `app/src/services/featureStore.ts` | 27 | JS/TS | `unwrapFeature` |
| `app/src/services/featureStore.ts` | 40 | JS/TS | `setUserFeatures` |
| `app/src/services/featureStore.ts` | 49 | JS/TS | `getUserFeatures` |
| `app/src/services/featureStore.ts` | 54 | JS/TS | `setItemFeatures` |
| `app/src/services/featureStore.ts` | 63 | JS/TS | `getItemFeatures` |
| `app/src/services/featureStore.ts` | 68 | JS/TS | `backfillFeaturesFromOffline` |
| `app/src/services/federated.ts` | 1 | JS/TS | `mask` |
| `app/src/services/finalStage.ts` | 39 | JS/TS | `runFinalStage` |
| `app/src/services/foundation.ts` | 7 | JS/TS | `loadFoundationModel` |
| `app/src/services/foundation.ts` | 17 | JS/TS | `isFoundationLoaded` |
| `app/src/services/foundation.ts` | 21 | JS/TS | `rankBatch` |
| `app/src/services/fulfillment.ts` | 7 | JS/TS | `allocate` |
| `app/src/services/identity.ts` | 3 | JS/TS | `resolveIdentity` |
| `app/src/services/infer.route.ts` | 3 | JS/TS | `infer` |
| `app/src/services/infer.ts` | 7 | JS/TS | `loadModel` |
| `app/src/services/infer.ts` | 17 | JS/TS | `inferScore` |
| `app/src/services/infer.ts` | 43 | JS/TS | `isModelLoaded` |
| `app/src/services/kafka.consumer.ts` | 3 | JS/TS | `startKafkaConsumer` |
| `app/src/services/kafka.ts` | 9 | JS/TS | `getProducer` |
| `app/src/services/kafka.ts` | 21 | JS/TS | `publish` |
| `app/src/services/killswitch.ts` | 3 | JS/TS | `enableKill` |
| `app/src/services/killswitch.ts` | 7 | JS/TS | `disableKill` |
| `app/src/services/killswitch.ts` | 11 | JS/TS | `isHalted` |
| `app/src/services/logger.ts` | 3 | JS/TS | `anonymizeUserId` |
| `app/src/services/logger.ts` | 7 | JS/TS | `sanitizeMeta` |
| `app/src/services/logger.ts` | 19 | JS/TS | `log` |
| `app/src/services/loyalty.ts` | 3 | JS/TS | `addPoints` |
| `app/src/services/loyalty.ts` | 13 | JS/TS | `getPoints` |
| `app/src/services/memory.ts` | 5 | JS/TS | `saveMemory` |
| `app/src/services/memory.ts` | 17 | JS/TS | `searchMemory` |
| `app/src/services/metrics.ts` | 3 | JS/TS | `inc` |
| `app/src/services/metrics.ts` | 7 | JS/TS | `metrics` |
| `app/src/services/order.ts` | 5 | JS/TS | `createOrder` |
| `app/src/services/ort.ts` | 8 | JS/TS | `create` |
| `app/src/services/ort.ts` | 14 | JS/TS | `getOrt` |
| `app/src/services/payment.ts` | 3 | JS/TS | `promptpayQR` |
| `app/src/services/policy.update.ts` | 7 | JS/TS | `updatePolicy` |
| `app/src/services/pricing.ts` | 1 | JS/TS | `optimalPrice` |
| `app/src/services/privacy.ts` | 3 | JS/TS | `requireSalt` |
| `app/src/services/privacy.ts` | 11 | JS/TS | `pseudo` |
| `app/src/services/privacy.ts` | 15 | JS/TS | `anonymizeEmail` |
| `app/src/services/push.ts` | 1 | JS/TS | `sendPush` |
| `app/src/services/rank.ts` | 9 | JS/TS | `dot` |
| `app/src/services/rank.ts` | 36 | JS/TS | `normalizeScore` |
| `app/src/services/rank.ts` | 44 | JS/TS | `hybridRankScore` |
| `app/src/services/rank.ts` | 60 | JS/TS | `resolveUserVec` |
| `app/src/services/rank.ts` | 69 | JS/TS | `estimateRlScore` |
| `app/src/services/rank.ts` | 75 | JS/TS | `rankProducts` |
| `app/src/services/ranker.transformer.ts` | 5 | JS/TS | `loadRanker` |
| `app/src/services/ranker.transformer.ts` | 15 | JS/TS | `scoreWithTransformer` |
| `app/src/services/reco.hybrid.ts` | 5 | JS/TS | `hybridRecommend` |
| `app/src/services/reco.index.ts` | 6 | JS/TS | `indexProduct` |
| `app/src/services/reco.search.ts` | 16 | JS/TS | `recommendVector` |
| `app/src/services/recommend.ts` | 11 | JS/TS | `recommendProducts` |
| `app/src/services/reconcile.ts` | 6 | JS/TS | `reconcile` |
| `app/src/services/redis.ts` | 14 | JS/TS | `ensureRedisConnected` |

_Display capped at 300 of 566 detected functions/methods._

### Automation pipeline files

- `.github/workflows/cloudflare-config.yml`
- `.github/workflows/codacy.yml`
- `.github/workflows/codeql.yml`
- `.github/workflows/codex-master-release.yml`
- `.github/workflows/dependabot.yml`
- `.github/workflows/deploy.yml`
- `.github/workflows/lint.yml`
- `.github/workflows/master-meta-full-source-ci-cd.yml`
- `.github/workflows/zlinebot-master-k8s-final-release.yml`
- `bootstrap.sh`
- `cloudflare/deploy.sh`
- `deploy-k8s.sh`
- `deploy.sh`
- `docker-compose.blue.yml`
- `docker-compose.green.yml`
- `docker-compose.yml`
- `docs/INSTALL_FULL.md`
- `docs/install_manual_en.md`
- `docs/install_manual_th.md`
- `infra/deploy.sh`
- `install-monitoring.sh`
- `install-observability.sh`
- `install.sh`
- `install_full.sh`
- `install_ultimate.sh`
- `k8s/api-deployment.yaml`
- `k8s/deploy.yaml`
- `k8s/deployment.yaml`
- `k8s/worker-deployment.yaml`
- `scripts/bootstrap-saas-foundation.sh`
- `scripts/deploy-k8s.sh`
- `scripts/deploy.sh`
- `scripts/install_android_deploy.sh`
- `scripts/install_auto.sh`
- `scripts/install_crdb_nocost.sh`
- `scripts/install_full.sh`
- `scripts/install_ios_deploy.sh`
- `scripts/install_istio.sh`
- `scripts/install_k3s.sh`
- `scripts/install_mesh_nocost.sh`
- `scripts/install_meta_fullstack_android_app_deploy.sh`
- `scripts/install_meta_fullstack_ios_app_deploy.sh`
- `scripts/install_meta_fullstack_mobile_app_deploy.sh`
- `scripts/install_mobile_fullstack_deploy.sh`
- `scripts/install_no_cost.sh`
- `scripts/install_secure.sh`
- `tools/codex_engine/pipeline.py`
- `watchdog.sh`

## `cvsz/zlttbots`

Status: **cloned**; commit `139bde44dfa3`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `AGENTS.md` | 272 | `/health` |
| `apps/affiliate-marketing/src/index.ts` | 43 | `/healthz` |
| `apps/affiliate-marketing/src/index.ts` | 47 | `/v1/trends/analyze` |
| `apps/affiliate-marketing/src/index.ts` | 57 | `/v1/copy/generate` |
| `apps/affiliate-marketing/src/index.ts` | 67 | `/v1/copy/optimize` |
| `apps/affiliate-marketing/src/index.ts` | 77 | `/v1/system/tiktok-shopee` |
| `apps/api-gateway/src/index.ts` | 23 | `/auth` |
| `apps/api-gateway/src/index.ts` | 74 | `/healthz` |
| `apps/api-gateway/src/middleware.ts` | 15 | `/healthz` |
| `apps/auth-service/src/index.ts` | 50 | `/healthz` |
| `apps/auth-service/src/index.ts` | 54 | `/register` |
| `apps/auth-service/src/index.ts` | 84 | `/login` |
| `apps/billing-service/src/index.ts` | 63 | `/webhook` |
| `apps/billing-service/src/index.ts` | 69 | `/healthz` |
| `apps/billing-service/src/index.ts` | 73 | `/checkout` |
| `apps/billing-service/src/index.ts` | 122 | `/webhook` |
| `apps/billing-service/src/index.ts` | 166 | `/plan/:userId` |
| `apps/frontend/src/App.tsx` | 28 | `/auth/register` |
| `apps/frontend/src/App.tsx` | 42 | `/auth/login` |
| `apps/frontend/vite.config.ts` | 14 | `/auth` |
| `apps/log-service/src/index.ts` | 65 | `/log` |
| `apps/log-service/src/index.ts` | 90 | `/healthz` |
| `apps/platform-core/src/index.ts` | 143 | `/healthz` |
| `apps/platform-core/src/index.ts` | 147 | `/github/repos` |
| `apps/platform-core/src/index.ts` | 157 | `/deploy` |
| `apps/platform-core/src/index.ts` | 234 | `/deploy/github` |
| `apps/platform-core/src/index.ts` | 262 | `/projects/public` |
| `apps/platform-core/src/index.ts` | 267 | `/orgs` |
| `apps/platform-core/src/index.ts` | 282 | `/orgs/:orgId/members` |
| `apps/worker-ai/src/index.ts` | 47 | `/healthz` |
| `apps/worker-ai/src/index.ts` | 51 | `/fix` |
| `generate-zlttbots.sh` | 1173 | `/analytics/revenue` |
| `generate-zlttbots.sh` | 1182 | `/analytics/conversion` |
| `generate-zlttbots.sh` | 1191 | `/analytics/campaigns` |
| `generate-zlttbots.sh` | 1200 | `/analytics/products` |
| `generate-zlttbots.sh` | 1412 | `/farm/health` |
| `generate-zlttbots.sh` | 1421 | `/farm/run` |
| `generate-zlttbots.sh` | 2379 | `/predict` |
| `generate-zlttbots.sh` | 2611 | `/farm/job` |
| `generate-zlttbots.sh` | 2619 | `/farm/status` |
| `generate-zlttbots.sh` | 2741 | `/render` |
| `generate-zlttbots.sh` | 3016 | `/crawl` |
| `generate-zlttbots.sh` | 3225 | `/arbitrage` |
| `infrastructure/cloudflare/waf-rules.tf` | 30 | `/auth/token\` |
| `infrastructure/k8s/global/failover.yaml` | 9 | `/healthz` |
| `installer/zlttbots-ai-full-stack-installer.sh` | 170 | `/status` |
| `installer/zlttbots-ai-full-stack-installer.sh` | 174 | `/kill` |
| `installer/zlttbots-ai-full-stack-installer.sh` | 180 | `/start` |
| `installer/zlttbots-ai-full-stack-installer.sh` | 285 | `/health` |
| `scripts/generate-enterprise-v8.sh` | 28 | `/health` |
| `scripts/generate-enterprise-v8.sh` | 32 | `/run-growth-cycle` |
| `scripts/generate-enterprise-v8.sh` | 73 | `/discover` |
| `scripts/generate-enterprise-v8.sh` | 117 | `/optimize` |
| `services/account-farm/src/api/server.js` | 11 | `/farm/health` |
| `services/account-farm/src/api/server.js` | 20 | `/farm/run` |
| `services/affiliate-webhook/src/main.py` | 46 | `/healthz` |
| `services/affiliate-webhook/src/main.py` | 65 | `/conversion` |
| `services/ai-orchestrator/src/main.py` | 10 | `/health` |
| `services/ai-orchestrator/src/main.py` | 14 | `/run-growth-cycle` |
| `services/analytics/src/api/server.js` | 10 | `/healthz` |
| `services/analytics/src/api/server.js` | 14 | `/analytics/summary` |
| `services/analytics/src/api/server.js` | 30 | `/analytics/revenue` |
| `services/analytics/src/api/server.js` | 35 | `/analytics/conversion` |
| `services/analytics/src/api/server.js` | 40 | `/analytics/campaigns` |
| `services/analytics/src/api/server.js` | 45 | `/analytics/products` |
| `services/arbitrage-engine/src/api/server.py` | 126 | `/health/live` |
| `services/arbitrage-engine/src/api/server.py` | 131 | `/health/ready` |
| `services/arbitrage-engine/src/api/server.py` | 140 | `/healthz` |
| `services/arbitrage-engine/src/api/server.py` | 155 | `/metrics` |
| `services/arbitrage-engine/src/api/server.py` | 160 | `/arbitrage` |
| `services/arbitrage-engine/src/api/server.py` | 169 | `/arbitrage/scan` |
| `services/arbitrage-engine/src/api/server.py` | 195 | `/affiliate/payouts/ingest` |
| `services/arbitrage-engine/src/api/server.py` | 215 | `/affiliate/sync` |
| `services/arbitrage-engine/src/api/server.py` | 237 | `/publishing/jobs` |
| `services/arbitrage-engine/src/api/server.py` | 245 | `/publishing/run-daily` |
| `services/arbitrage-engine/src/api/server.py` | 254 | `/publishing/counters/{tenant_id}` |
| `services/arbitrage-engine/src/api/server.py` | 260 | `/videos` |
| `services/arbitrage-engine/src/api/server.py` | 266 | `/performance` |
| `services/arbitrage-engine/src/api/server.py` | 276 | `/reporting/posted-products/{tenant_id}` |
| `services/billing-service/src/main.py` | 18 | `/healthz` |
| `services/billing-service/src/main.py` | 23 | `/charge` |
| `services/budget-allocator/src/main.py` | 44 | `/healthz` |
| `services/budget-allocator/src/main.py` | 49 | `/allocate` |
| `services/campaign-optimizer/src/main.py` | 39 | `/optimize` |
| `services/capital-allocator/src/main.py` | 42 | `/healthz` |
| `services/capital-allocator/src/main.py` | 52 | `/allocate` |
| `services/click-tracker/src/server/server.js` | 34 | `/healthz` |
| `services/exchange/src/main.py` | 39 | `/order` |
| `services/execution-engine/src/main.py` | 80 | `/healthz` |
| `services/execution-engine/src/main.py` | 94 | `/publish` |
| `services/execution-engine/src/main.py` | 115 | `/status/{external_id}` |
| `services/feature-store/src/main.py` | 109 | `/healthz` |
| `services/feature-store/src/main.py` | 119 | `/features/{campaign_id}` |
| `services/feature-store/src/main.py` | 126 | `/features/{campaign_id}` |
| `services/feature-store/src/main.py` | 162 | `/features/{campaign_id}` |
| `services/feature-store/src/main.py` | 171 | `/features` |
| `services/federation/src/main.py` | 152 | `/healthz` |
| `services/federation/src/main.py` | 167 | `/nodes` |
| `services/federation/src/main.py` | 197 | `/register` |
| `services/gpu-renderer/src/api/server.py` | 20 | `/healthz` |
| `services/gpu-renderer/src/api/server.py` | 39 | `/metrics` |
| `services/gpu-renderer/src/api/server.py` | 44 | `/render` |
| `services/jwt-auth/src/main.py` | 129 | `/healthz` |
| `services/jwt-auth/src/main.py` | 134 | `/.well-known/jwks.json` |
| `services/jwt-auth/src/main.py` | 148 | `/token` |
| `services/jwt-auth/src/main.py` | 200 | `/introspect` |
| `services/landing-service/src/main.py` | 9 | `/healthz` |
| `services/landing-service/src/main.py` | 14 | `/landing/{product_id}` |
| `services/market-crawler/src/api/server.py` | 19 | `/healthz` |
| `services/market-crawler/src/api/server.py` | 38 | `/metrics` |
| `services/market-crawler/src/api/server.py` | 43 | `/crawl` |
| `services/market-orchestrator/src/main.py` | 19 | `/healthz` |
| `services/market-orchestrator/src/main.py` | 24 | `/launch` |
| `services/master-orchestrator/src/main.py` | 178 | `/healthz` |
| `services/master-orchestrator/src/main.py` | 183 | `/deployments` |
| `services/master-orchestrator/src/main.py` | 199 | `/deployments/{deployment_id}` |
| `services/master-orchestrator/src/main.py` | 208 | `/deployments/events` |
| `services/master-orchestrator/src/main.py` | 233 | `/run-cycle` |
| `services/master-orchestrator/src/main.py` | 238 | `/economy/run` |
| `services/master-orchestrator/src/main.py` | 243 | `/federation/run` |
| `services/master-orchestrator/src/main.py` | 273 | `/campaign/run` |
| `services/master-orchestrator/src/main.py` | 314 | `/profit-mode/activate` |
| `services/model-registry/src/main.py` | 128 | `/healthz` |
| `services/model-registry/src/main.py` | 137 | `/register` |
| `services/model-registry/src/main.py` | 153 | `/latest/{name}` |
| `services/model-service/src/main.py` | 106 | `/healthz` |
| `services/model-service/src/main.py` | 117 | `/metrics` |
| `services/model-service/src/main.py` | 122 | `/predict` |
| `services/model-service/src/main.py` | 127 | `/predict_async` |
| `services/model-service/src/main.py` | 153 | `/result/{job_id}` |
| `services/model-service/src/main.py` | 161 | `/result/{job_id}/wait` |
| `services/payment/main.py` | 29 | `/healthz` |
| `services/payment/main.py` | 34 | `/checkout` |
| `services/payment/webhook.py` | 79 | `/webhook` |
| `services/pr-bot/src/server.py` | 20 | `/healthz` |
| `services/pr-bot/src/server.py` | 25 | `/webhook` |
| `services/product-discovery/src/main.py` | 161 | `/healthz` |
| `services/product-discovery/src/main.py` | 181 | `/discover` |
| `services/product-generator/src/main.py` | 19 | `/healthz` |
| `services/product-generator/src/main.py` | 24 | `/generate` |
| `services/reward-collector/src/main.py` | 40 | `/healthz` |
| `services/reward-collector/src/main.py` | 45 | `/reward` |
| `services/rl-coordinator/src/main.py` | 48 | `/healthz` |
| `services/rl-coordinator/src/main.py` | 70 | `/decide` |
| `services/rl-coordinator/src/main.py` | 95 | `/update` |
| `services/rl-engine/src/main.py` | 106 | `/healthz` |
| `services/rl-engine/src/main.py` | 122 | `/select` |
| `services/rl-engine/src/main.py` | 153 | `/update` |
| `services/rl-policy/src/main.py` | 25 | `/healthz` |
| `services/rl-policy/src/main.py` | 30 | `/train` |
| `services/rtb-engine/src/main.py` | 74 | `/healthz` |
| `services/rtb-engine/src/main.py` | 79 | `/bid` |
| `services/rtb-engine/src/main.py` | 116 | `/openrtb/bid` |
| `services/scaling-engine/src/main.py` | 23 | `/healthz` |
| `services/scaling-engine/src/main.py` | 28 | `/scale` |
| `services/scheduler/src/main.py` | 38 | `/healthz` |
| `services/scheduler/src/main.py` | 92 | `/assign` |
| `services/security-api/src/app/routes.py` | 11 | `/healthz` |
| `services/security-api/src/app/routes.py` | 16 | `/scan` |
| `services/security-api/src/app/routes.py` | 27 | `/scan/{scan_id}` |
| `services/tenant-service/src/main.py` | 58 | `/healthz` |
| `services/tenant-service/src/main.py` | 72 | `/tenant` |
| `services/tiktok-farm/src/api/server.js` | 8 | `/farm/job` |
| `services/tiktok-farm/src/api/server.js` | 16 | `/farm/status` |
| `services/tracking/server.py` | 59 | `/healthz` |
| `services/tracking/server.py` | 64 | `/t/{campaign_id}` |
| `services/viral-predictor/src/api/server.py` | 23 | `/healthz` |
| `services/viral-predictor/src/api/server.py` | 47 | `/metrics` |
| `services/viral-predictor/src/api/server.py` | 52 | `/predict` |
| `tests/runtime/test_service_health_import_safe.py` | 23 | `/healthz` |
| `tests/runtime/test_service_health_import_safe.py` | 30 | `/healthz` |
| `tests/runtime/test_service_health_import_safe.py` | 37 | `/healthz` |
| `tests/security_platform/test_security_api.py` | 20 | `/healthz` |
| `tests/test_feature_impact_dive.py` | 46 | `/healthz` |
| `tests/test_feature_impact_dive.py` | 52 | `/deploy` |
| `tests/test_feature_impact_dive.py` | 68 | `/healthz` |
| `tests/test_tracking_and_payment_webhook.py` | 78 | `/webhook` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `agents/analyzer.py` | 4 | Python | `analyze` |
| `agents/ci.py` | 9 | Python | `_run_pytest` |
| `agents/ci.py` | 23 | Python | `fix_ci` |
| `agents/guard.py` | 6 | Python | `validate_runtime` |
| `agents/orchestrator.py` | 17 | Python | `main` |
| `agents/pr.py` | 6 | Python | `_run` |
| `agents/pr.py` | 11 | Python | `create_patch` |
| `agents/refactor.py` | 4 | Python | `refactor` |
| `agents/security.py` | 4 | Python | `secure` |
| `apps/affiliate-marketing/src/config.ts` | 12 | JS/TS | `loadConfig` |
| `apps/affiliate-marketing/src/engine.ts` | 8 | JS/TS | `calculateMomentum` |
| `apps/affiliate-marketing/src/engine.ts` | 14 | JS/TS | `rankProducts` |
| `apps/affiliate-marketing/src/engine.ts` | 24 | JS/TS | `generateCaption` |
| `apps/affiliate-marketing/src/engine.ts` | 44 | JS/TS | `optimizeCaption` |
| `apps/affiliate-marketing/src/engine.ts` | 62 | JS/TS | `buildTikTokShopeeSystem` |
| `apps/api-gateway/src/index.ts` | 116 | JS/TS | `shutdown` |
| `apps/api-gateway/src/middleware.ts` | 17 | JS/TS | `normalizePath` |
| `apps/api-gateway/src/middleware.ts` | 25 | JS/TS | `isJwtClaims` |
| `apps/api-gateway/src/middleware.ts` | 33 | JS/TS | `auth` |
| `apps/auth-service/src/index.ts` | 114 | JS/TS | `shutdown` |
| `apps/billing-service/src/index.ts` | 171 | JS/TS | `start` |
| `apps/billing-service/src/index.ts` | 177 | JS/TS | `shutdown` |
| `apps/frontend/src/App.tsx` | 22 | JS/TS | `register` |
| `apps/frontend/src/App.tsx` | 36 | JS/TS | `login` |
| `apps/frontend/src/App.tsx` | 59 | JS/TS | `deploy` |
| `apps/frontend/src/App.tsx` | 79 | JS/TS | `deployRepo` |
| `apps/frontend/src/App.tsx` | 91 | JS/TS | `setStatus` |
| `apps/frontend/src/App.tsx` | 94 | JS/TS | `loadGallery` |
| `apps/frontend/src/Logs.tsx` | 3 | JS/TS | `getDefaultWsUrl` |
| `apps/log-service/src/index.ts` | 57 | JS/TS | `broadcastLog` |
| `apps/log-service/src/index.ts` | 94 | JS/TS | `start` |
| `apps/log-service/src/index.ts` | 103 | JS/TS | `shutdown` |
| `apps/platform-core/src/github.ts` | 3 | JS/TS | `getRepos` |
| `apps/platform-core/src/index.ts` | 119 | JS/TS | `getUserContext` |
| `apps/platform-core/src/index.ts` | 128 | JS/TS | `requireUserContextOrRespond` |
| `apps/platform-core/src/index.ts` | 308 | JS/TS | `initDatabase` |
| `apps/platform-core/src/index.ts` | 325 | JS/TS | `start` |
| `apps/platform-core/src/index.ts` | 334 | JS/TS | `shutdown` |
| `apps/platform-core/src/limits.ts` | 7 | JS/TS | `assertAllowedBillingUrl` |
| `apps/platform-core/src/limits.ts` | 26 | JS/TS | `checkLimits` |
| `apps/platform-core/src/org.ts` | 6 | JS/TS | `createOrg` |
| `apps/platform-core/src/org.ts` | 17 | JS/TS | `getOrg` |
| `apps/platform-core/src/org.ts` | 21 | JS/TS | `addMember` |
| `apps/platform-core/src/org.ts` | 35 | JS/TS | `requireRole` |
| `apps/platform-core/src/public.ts` | 1 | JS/TS | `generatePublicUrl` |
| `apps/worker-ai/src/index.ts` | 115 | JS/TS | `shutdown` |
| `enterprise_maturity/full_upgrade.py` | 36 | Python | `validate` |
| `enterprise_maturity/go_to_market.py` | 15 | Python | `validate` |
| `enterprise_maturity/go_to_market.py` | 29 | Python | `validate` |
| `enterprise_maturity/go_to_market.py` | 43 | Python | `validate` |
| `enterprise_maturity/go_to_market.py` | 60 | Python | `validate` |
| `enterprise_maturity/go_to_market.py` | 73 | Python | `validate` |
| `enterprise_maturity/go_to_market.py` | 87 | Python | `_select_hook` |
| `enterprise_maturity/go_to_market.py` | 93 | Python | `generate_daily_plan` |
| `enterprise_maturity/go_to_market.py` | 144 | Python | `generate_post` |
| `enterprise_maturity/go_to_market.py` | 162 | Python | `generate_pitch_deck_outline` |
| `enterprise_maturity/go_to_market.py` | 178 | Python | `generate_date_stamped_status` |
| `enterprise_maturity/governance.py` | 14 | Python | `validate` |
| `enterprise_maturity/governance.py` | 27 | Python | `assert_backward_compatible` |
| `enterprise_maturity/operations.py` | 22 | Python | `remaining_budget` |
| `enterprise_maturity/operations.py` | 47 | Python | `__init__` |
| `enterprise_maturity/operations.py` | 50 | Python | `route` |
| `enterprise_maturity/operations.py` | 64 | Python | `__init__` |
| `enterprise_maturity/operations.py` | 68 | Python | `run` |
| `enterprise_maturity/performance.py` | 20 | Python | `__init__` |
| `enterprise_maturity/performance.py` | 24 | Python | `recommend` |
| `enterprise_maturity/performance.py` | 35 | Python | `__init__` |
| `enterprise_maturity/performance.py` | 40 | Python | `admit` |
| `enterprise_maturity/performance.py` | 51 | Python | `complete` |
| `enterprise_maturity/resilience.py` | 17 | Python | `run` |
| `enterprise_maturity/resilience.py` | 31 | Python | `__init__` |
| `enterprise_maturity/resilience.py` | 37 | Python | `_is_open` |
| `enterprise_maturity/resilience.py` | 43 | Python | `call` |
| `enterprise_maturity/resilience.py` | 61 | Python | `__init__` |
| `enterprise_maturity/resilience.py` | 64 | Python | `execute` |
| `enterprise_maturity/security.py` | 20 | Python | `is_rotation_due` |
| `enterprise_maturity/security.py` | 28 | Python | `__init__` |
| `enterprise_maturity/security.py` | 32 | Python | `put` |
| `enterprise_maturity/security.py` | 39 | Python | `get` |
| `enterprise_maturity/security.py` | 44 | Python | `due_for_rotation` |
| `enterprise_maturity/security.py` | 60 | Python | `__init__` |
| `enterprise_maturity/security.py` | 63 | Python | `authorize` |
| `enterprise_maturity/security.py` | 80 | Python | `digest` |
| `enterprise_maturity/security.py` | 88 | Python | `__init__` |
| `enterprise_maturity/security.py` | 91 | Python | `emit` |
| `enterprise_maturity/security.py` | 95 | Python | `list_events` |
| `enterprise_maturity/v3_runtime/api_gateway.py` | 17 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/api_gateway.py` | 22 | Python | `add_route` |
| `enterprise_maturity/v3_runtime/api_gateway.py` | 25 | Python | `route` |
| `enterprise_maturity/v3_runtime/api_gateway.py` | 38 | Python | `routes` |
| `enterprise_maturity/v3_runtime/autoscaling.py` | 18 | Python | `reconcile` |
| `enterprise_maturity/v3_runtime/crawler_cluster.py` | 19 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/crawler_cluster.py` | 24 | Python | `submit` |
| `enterprise_maturity/v3_runtime/crawler_cluster.py` | 32 | Python | `reconcile` |
| `enterprise_maturity/v3_runtime/crawler_cluster.py` | 36 | Python | `dispatch_batch` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 30 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 38 | Python | `assess` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 76 | Python | `_validate_metric` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 102 | Python | `optimize` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 131 | Python | `_validate` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 162 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 168 | Python | `_connect` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 171 | Python | `_initialize` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 199 | Python | `record_violation` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 217 | Python | `record_fix` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 241 | Python | `snapshot` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 266 | Python | `_validate_violation` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 271 | Python | `_validate_fix` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 289 | Python | `detect_drift` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 295 | Python | `heal` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 302 | Python | `_validate` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 316 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/deployment_pipeline.py` | 320 | Python | `format_auto_fix_notification` |
| `enterprise_maturity/v3_runtime/gpu_scheduler.py` | 20 | Python | `available_gb` |
| `enterprise_maturity/v3_runtime/gpu_scheduler.py` | 27 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/gpu_scheduler.py` | 32 | Python | `schedule` |
| `enterprise_maturity/v3_runtime/gpu_scheduler.py` | 41 | Python | `release` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 28 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 57 | Python | `bootstrap` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 65 | Python | `enqueue_domain_event` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 68 | Python | `schedule_render` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 73 | Python | `evaluate_multi_region_canary` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 76 | Python | `optimize_pipeline_execution` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 79 | Python | `record_policy_violation` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 82 | Python | `record_policy_fix` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 85 | Python | `detect_manifest_drift` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 94 | Python | `heal_manifest` |
| `enterprise_maturity/v3_runtime/orchestrator.py` | 97 | Python | `prepare_auto_fix_notifications` |
| `enterprise_maturity/v3_runtime/queue_system.py` | 17 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/queue_system.py` | 20 | Python | `publish` |
| `enterprise_maturity/v3_runtime/queue_system.py` | 23 | Python | `consume` |
| `enterprise_maturity/v3_runtime/queue_system.py` | 30 | Python | `depth` |
| `enterprise_maturity/v3_runtime/queue_system.py` | 33 | Python | `topics` |
| `enterprise_maturity/v3_runtime/service_discovery.py` | 16 | Python | `endpoint` |
| `enterprise_maturity/v3_runtime/service_discovery.py` | 23 | Python | `__init__` |
| `enterprise_maturity/v3_runtime/service_discovery.py` | 26 | Python | `register` |
| `enterprise_maturity/v3_runtime/service_discovery.py` | 31 | Python | `deregister` |
| `enterprise_maturity/v3_runtime/service_discovery.py` | 34 | Python | `resolve` |
| `enterprise_maturity/v3_runtime/service_discovery.py` | 37 | Python | `snapshot` |
| `enterprise_maturity/v3_upgrade.py` | 38 | Python | `__init__` |
| `enterprise_maturity/v3_upgrade.py` | 49 | Python | `default_blueprint` |
| `enterprise_maturity/v3_upgrade.py` | 122 | Python | `validate` |
| `enterprise_maturity/v3_upgrade.py` | 142 | Python | `phase_checklist` |
| `enterprise_maturity/v3_upgrade.py` | 149 | Python | `_require_unique` |
| `infrastructure/scripts/check-iac-policy.sh` | 8 | Shell | `check_k8s_file` |
| `infrastructure/scripts/node-dependency-scan.sh` | 10 | Shell | `run_outdated_fallback` |
| `infrastructure/scripts/node-dependency-scan.sh` | 19 | Shell | `run_audit` |
| `infrastructure/start/workers.sh` | 7 | Shell | `run_worker` |
| `installer/full-stack-installer.sh` | 11 | Shell | `log` |
| `installer/full-stack-installer.sh` | 15 | Shell | `fail` |
| `installer/full-stack-installer.sh` | 20 | Shell | `usage` |
| `installer/full-stack-installer.sh` | 41 | Shell | `require_cmd` |
| `installer/full-stack-installer.sh` | 46 | Shell | `file_has_pattern` |
| `installer/full-stack-installer.sh` | 57 | Shell | `compose_cmd` |
| `installer/full-stack-installer.sh` | 67 | Shell | `check_prereqs` |
| `installer/full-stack-installer.sh` | 84 | Shell | `create_env_if_needed` |
| `installer/full-stack-installer.sh` | 95 | Shell | `validate_env_security` |
| `installer/full-stack-installer.sh` | 111 | Shell | `start_stack` |
| `installer/full-stack-installer.sh` | 122 | Shell | `health_probe` |
| `installer/full-stack-installer.sh` | 133 | Shell | `verify_stack` |
| `installer/full-stack-installer.sh` | 142 | Shell | `stop_stack` |
| `installer/full-stack-installer.sh` | 147 | Shell | `uninstall_stack` |
| `installer/full-stack-installer.sh` | 152 | Shell | `run_deploy` |
| `installer/full-stack-installer.sh` | 161 | Shell | `main` |
| `installer/zlttbots-ai-full-stack-installer.sh` | 33 | Shell | `log` |
| `memory/learning.py` | 6 | Python | `learn_from_ci` |
| `memory/learning.py` | 10 | Python | `suggest_fix` |
| `memory/vector_store.py` | 8 | Python | `__init__` |
| `memory/vector_store.py` | 11 | Python | `_embed` |
| `memory/vector_store.py` | 14 | Python | `add` |
| `memory/vector_store.py` | 18 | Python | `search` |
| `packages/shared-runtime/src/circuitBreaker.ts` | 18 | JS/TS | `constructor` |
| `packages/shared-runtime/src/circuitBreaker.ts` | 24 | JS/TS | `getState` |
| `packages/shared-runtime/src/circuitBreaker.ts` | 29 | JS/TS | `syncOpenWindow` |
| `packages/shared-runtime/src/circuitBreaker.ts` | 64 | JS/TS | `onSuccess` |
| `packages/shared-runtime/src/circuitBreaker.ts` | 70 | JS/TS | `onFailure` |
| `packages/shared-runtime/src/idempotency.ts` | 14 | JS/TS | `constructor` |
| `packages/shared-runtime/src/idempotency.ts` | 18 | JS/TS | `get` |
| `packages/shared-runtime/src/idempotency.ts` | 30 | JS/TS | `set` |
| `packages/shared-runtime/src/idempotency.ts` | 40 | JS/TS | `buildIdempotencyKey` |
| `packages/shared-runtime/src/kafka.ts` | 16 | JS/TS | `constructor` |
| `packages/shared-runtime/src/kafka.ts` | 31 | JS/TS | `connect` |
| `packages/shared-runtime/src/kafka.ts` | 35 | JS/TS | `disconnect` |
| `packages/shared-runtime/src/kafka.ts` | 39 | JS/TS | `sendExactlyOnce` |
| `packages/shared-runtime/src/otel.ts` | 7 | JS/TS | `parseOtelHeaders` |
| `packages/shared-runtime/src/otel.ts` | 26 | JS/TS | `initTelemetry` |
| `packages/shared-runtime/src/retry.ts` | 18 | JS/TS | `delay` |
| `scripts/deploy-global.sh` | 7 | Shell | `require_cmd` |
| `scripts/deploy-zlttbots-platform.sh` | 7 | Shell | `log` |
| `scripts/deploy-zlttbots-production.sh` | 125 | Shell | `check_service` |
| `scripts/extract_gemini_source.py` | 40 | Python | `digest` |
| `scripts/extract_gemini_source.py` | 44 | Python | `_iter_code_fences` |
| `scripts/extract_gemini_source.py` | 49 | Python | `extract_artifacts` |
| `scripts/extract_gemini_source.py` | 107 | Python | `_safe_join` |
| `scripts/extract_gemini_source.py` | 115 | Python | `write_artifacts` |
| `scripts/extract_gemini_source.py` | 144 | Python | `parse_args` |
| `scripts/extract_gemini_source.py` | 151 | Python | `main` |
| `scripts/feature_impact_dive.py` | 59 | Python | `total_features` |
| `scripts/feature_impact_dive.py` | 97 | Python | `_read_lines` |
| `scripts/feature_impact_dive.py` | 101 | Python | `_extract_endpoint_signatures` |
| `scripts/feature_impact_dive.py` | 112 | Python | `extract_compose_services` |
| `scripts/feature_impact_dive.py` | 136 | Python | `extract_app_features` |
| `scripts/feature_impact_dive.py` | 158 | Python | `extract_runtime_service_features` |
| `scripts/feature_impact_dive.py` | 192 | Python | `extract_documented_features` |
| `scripts/feature_impact_dive.py` | 209 | Python | `build_impact_report` |
| `scripts/feature_impact_dive.py` | 218 | Python | `build_surface_manifest` |
| `scripts/feature_impact_dive.py` | 279 | Python | `manifest_to_dict` |
| `scripts/feature_impact_dive.py` | 294 | Python | `_deep_sort` |
| `scripts/feature_impact_dive.py` | 303 | Python | `resolve_output_path` |
| `scripts/feature_impact_dive.py` | 310 | Python | `write_manifest` |
| `scripts/feature_impact_dive.py` | 319 | Python | `_validate_manifest_shape` |
| `scripts/feature_impact_dive.py` | 351 | Python | `validate_manifest` |
| `scripts/feature_impact_dive.py` | 408 | Python | `_render_section` |
| `scripts/feature_impact_dive.py` | 413 | Python | `format_markdown` |
| `scripts/feature_impact_dive.py` | 459 | Python | `write_report` |
| `scripts/feature_impact_dive.py` | 464 | Python | `main` |
| `scripts/full_migration.py` | 93 | Python | `is_text_candidate` |
| `scripts/full_migration.py` | 102 | Python | `is_excluded` |
| `scripts/full_migration.py` | 106 | Python | `iter_files` |
| `scripts/full_migration.py` | 119 | Python | `apply_mapping` |
| `scripts/full_migration.py` | 129 | Python | `migrate_file_contents` |
| `scripts/full_migration.py` | 144 | Python | `plan_renames` |
| `scripts/full_migration.py` | 168 | Python | `apply_renames` |
| `scripts/full_migration.py` | 182 | Python | `run` |
| `scripts/full_migration.py` | 202 | Python | `main` |
| `scripts/generate-domain-env.sh` | 8 | Shell | `usage` |
| `scripts/generate-domain-env.sh` | 22 | Shell | `require_non_empty` |
| `scripts/install-zeaz-edge-stack.sh` | 165 | Shell | `create_dns` |
| `scripts/install-zlttbots-platform.sh` | 7 | Shell | `log` |
| `scripts/install-zlttbots-platform.sh` | 8 | Shell | `fail` |
| `scripts/install-zlttbots-platform.sh` | 10 | Shell | `require_cmd` |
| `scripts/install-zlttbots-platform.sh` | 14 | Shell | `ensure_env` |
| `scripts/install-zlttbots-platform.sh` | 48 | Shell | `install_node_dependencies` |
| `scripts/install-zlttbots-platform.sh` | 58 | Shell | `validate_compose` |
| `scripts/migrate-edge-domain.sh` | 70 | Shell | `cf_api` |
| `scripts/migrate-edge-domain.sh` | 81 | Shell | `ensure_dns` |
| `scripts/node-services-lib.sh` | 13 | Shell | `zlttbots_node_discover_services` |
| `scripts/node-services-lib.sh` | 22 | Shell | `zlttbots_node_print_service_list` |
| `scripts/node-services-lib.sh` | 26 | Shell | `zlttbots_node_service_dir` |
| `scripts/node-services-lib.sh` | 30 | Shell | `zlttbots_node_service_package` |
| `scripts/node-services-lib.sh` | 34 | Shell | `zlttbots_node_prepare_runtime_dirs` |
| `scripts/node-services-lib.sh` | 38 | Shell | `zlttbots_node_has_package` |
| `scripts/provision-cloudflare-edge.sh` | 62 | Shell | `cf_api` |
| `scripts/provision-cloudflare-edge.sh` | 73 | Shell | `ensure_dns` |
| `scripts/rebrand_zlttbots_to_zlttbots.py` | 77 | Python | `is_excluded` |
| `scripts/rebrand_zlttbots_to_zlttbots.py` | 81 | Python | `should_scan_file` |
| `scripts/rebrand_zlttbots_to_zlttbots.py` | 89 | Python | `iter_candidate_files` |
| `scripts/rebrand_zlttbots_to_zlttbots.py` | 100 | Python | `apply_replacements` |
| `scripts/rebrand_zlttbots_to_zlttbots.py` | 113 | Python | `audit_and_rebrand` |
| `scripts/rebrand_zlttbots_to_zlttbots.py` | 146 | Python | `write_report` |
| `scripts/rebrand_zlttbots_to_zlttbots.py` | 151 | Python | `main` |
| `scripts/repo_cleanup_audit.py` | 56 | Python | `load_config` |
| `scripts/repo_cleanup_audit.py` | 71 | Python | `is_excluded` |
| `scripts/repo_cleanup_audit.py` | 75 | Python | `list_source_files` |
| `scripts/repo_cleanup_audit.py` | 94 | Python | `candidate_local_module_targets` |
| `scripts/repo_cleanup_audit.py` | 114 | Python | `resolve_python_import` |
| `scripts/repo_cleanup_audit.py` | 142 | Python | `build_dependency_graph` |
| `scripts/repo_cleanup_audit.py` | 166 | Python | `normalize_path_list` |
| `scripts/repo_cleanup_audit.py` | 175 | Python | `traverse_reachable` |
| `scripts/repo_cleanup_audit.py` | 189 | Python | `analyze` |
| `scripts/repo_cleanup_audit.py` | 211 | Python | `write_report` |
| `scripts/repo_cleanup_audit.py` | 228 | Python | `quarantine_files` |
| `scripts/repo_cleanup_audit.py` | 246 | Python | `main` |
| `scripts/security_autofix_engine.py` | 31 | Python | `_iter_findings` |
| `scripts/security_autofix_engine.py` | 42 | Python | `_safe_replace` |
| `scripts/security_autofix_engine.py` | 52 | Python | `apply_repo_fixes` |
| `scripts/security_autofix_engine.py` | 82 | Python | `build_report` |
| `scripts/security_autofix_engine.py` | 98 | Python | `main` |
| `scripts/start-zlttbots-platform.sh` | 7 | Shell | `log` |
| `scripts/start-zlttbots-platform.sh` | 8 | Shell | `fail` |
| `scripts/start-zlttbots-platform.sh` | 23 | Shell | `compose_cmd` |
| `scripts/start-zlttbots-platform.sh` | 33 | Shell | `is_snapshot_cache_error` |
| `scripts/start-zlttbots-platform.sh` | 38 | Shell | `compose_up_with_recovery` |
| `scripts/start-zlttbots-platform.sh` | 65 | Shell | `ensure_env` |
| `scripts/start-zlttbots-platform.sh` | 69 | Shell | `print_usage` |
| `scripts/start-zlttbots-platform.sh` | 84 | Shell | `resolve_mode` |
| `scripts/start-zlttbots-platform.sh` | 108 | Shell | `smoke_check` |
| `scripts/start-zlttbots.sh` | 16 | Shell | `log` |
| `scripts/start-zlttbots.sh` | 90 | Shell | `install_node_service` |
| `scripts/start-zlttbots.sh` | 125 | Shell | `run` |
| `scripts/test-integration.sh` | 8 | Shell | `log` |
| `scripts/test-integration.sh` | 12 | Shell | `assert_running` |
| `scripts/test-integration.sh` | 25 | Shell | `assert_http_2xx` |
| `scripts/upgrade-zlttbots-platform.sh` | 7 | Shell | `log` |
| `scripts/upgrade-zlttbots-platform.sh` | 8 | Shell | `fail` |
| `scripts/upgrade-zlttbots-platform.sh` | 10 | Shell | `compose_cmd` |
| `scripts/validate_docs_snippets.py` | 15 | Python | `iter_markdown_files` |
| `scripts/validate_docs_snippets.py` | 19 | Python | `validate` |
| `scripts/validate_docs_snippets.py` | 39 | Python | `main` |
| `scripts/zlttbots-edge-doctor.sh` | 72 | Shell | `cf_api` |
| `scripts/zlttbots-manager.sh` | 7 | Shell | `fail` |
| `scripts/zlttbots-manager.sh` | 9 | Shell | `compose_cmd` |
| `scripts/zlttbots-manager.sh` | 19 | Shell | `status_cmd` |
| `scripts/zlttbots-manager.sh` | 25 | Shell | `logs_cmd` |
| `scripts/zlttbots-manager.sh` | 33 | Shell | `usage` |
| `scripts/zlttbots-node.sh` | 10 | Shell | `usage` |
| `scripts/zlttbots-supervisor.sh` | 17 | Shell | `check_service` |
| `services/account-farm/src/core/accounts.js` | 3 | JS/TS | `getAccount` |
| `services/account-farm/src/core/jobs.js` | 3 | JS/TS | `nextJob` |
| `services/account-farm/src/core/jobs.js` | 13 | JS/TS | `completeJob` |

_Display capped at 300 of 1448 detected functions/methods._

### Automation pipeline files

- `.github/workflows/ai-bot.yml`
- `.github/workflows/ai-pr-review.yml`
- `.github/workflows/build.yml`
- `.github/workflows/codeql-analysis.yml`
- `.github/workflows/enterprise-platform-delivery.yml`
- `.github/workflows/infra-autotune-train.yml`
- `.github/workflows/nextjs.yml`
- `.github/workflows/security-autofix.yml`
- `.github/workflows/zlttbots-ci.yml`
- `cloudflare-devops/docker/docker-compose.cloudflare.yml`
- `cloudflare-devops/install.sh`
- `deploy/helm/zlttbots/templates/deployment.yaml`
- `docker-compose.enterprise.yml`
- `docker-compose.yml`
- `docs/full-stack-installation.md`
- `docs/installation.md`
- `docs/operations/event-pipeline.md`
- `enterprise_maturity/v3_runtime/deployment_pipeline.py`
- `infra/docker-compose.yml`
- `infrastructure/ci/deploy.sh`
- `infrastructure/monitoring/docker-compose.monitoring.yml`
- `infrastructure/postgres/docker-compose.postgres.yml`
- `infrastructure/scripts/bootstrap-platform.sh`
- `infrastructure/scripts/deploy-kubernetes.sh`
- `installer/full-stack-installer.sh`
- `installer/zlttbots-ai-full-stack-installer.sh`
- `scripts/deploy-cloudflare-edge.sh`
- `scripts/deploy-global.sh`
- `scripts/deploy-zlttbots-platform.sh`
- `scripts/deploy-zlttbots-production.sh`
- `scripts/install-istio.sh`
- `scripts/install-node-services.sh`
- `scripts/install-zeaz-edge-stack.sh`
- `scripts/install-zlttbots-cloudflare.sh`
- `scripts/install-zlttbots-platform.sh`
- `scripts/k8s-bootstrap.sh`
- `scripts/k8s-deploy-enterprise.sh`
- `scripts/node-install.sh`
- `services/data/pipeline.py`
- `services/master-orchestrator/src/deployment_controller.py`
- `services/rl-trainer/src/autonomy/infra_tuner_pipeline.py`
- `tests/test_profit_mode_pipeline.py`

## `cvsz/zttlbots`

Status: **cloned**; commit `40949c6038d1`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `zlinebot-lean/app/src/routes/line.ts` | 8 | `/` |
| `zlinebot-lean/app/src/routes/stream.ts` | 10 | `/sse` |
| `zlinebot-lean/app/src/routes/stream.ts` | 27 | `/ws` |
| `zlinebot-lean/app/src/routes/tiktok.ts` | 7 | `/` |
| `zlinebot-lean/app/src/server.ts` | 22 | `/healthz` |
| `zlinebot-lean/app/src/server.ts` | 23 | `/webhook/line` |
| `zlinebot-lean/app/src/server.ts` | 24 | `/webhook/tiktok` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `zlinebot-lean/app/src/billing/cost.ts` | 1 | JS/TS | `calcCost` |
| `zlinebot-lean/app/src/billing/guard.ts` | 3 | JS/TS | `tenantQuotaKey` |
| `zlinebot-lean/app/src/billing/guard.ts` | 7 | JS/TS | `getTenantQuota` |
| `zlinebot-lean/app/src/billing/guard.ts` | 11 | JS/TS | `enforceQuota` |
| `zlinebot-lean/app/src/billing/ledger.ts` | 14 | JS/TS | `recordUsageEvent` |
| `zlinebot-lean/app/src/billing/meter.ts` | 13 | JS/TS | `recordUsage` |
| `zlinebot-lean/app/src/core/config.ts` | 10 | JS/TS | `required` |
| `zlinebot-lean/app/src/core/region-routing.ts` | 5 | JS/TS | `nearestRegion` |
| `zlinebot-lean/app/src/core/region-routing.ts` | 9 | JS/TS | `routeRegion` |
| `zlinebot-lean/app/src/core/security.ts` | 3 | JS/TS | `timingSafeEqual` |
| `zlinebot-lean/app/src/core/security.ts` | 11 | JS/TS | `lineSignature` |
| `zlinebot-lean/app/src/llm/agent.ts` | 4 | JS/TS | `agent` |
| `zlinebot-lean/app/src/llm/bandit.ts` | 10 | JS/TS | `totalCount` |
| `zlinebot-lean/app/src/llm/bandit.ts` | 14 | JS/TS | `ucb` |
| `zlinebot-lean/app/src/llm/bandit.ts` | 18 | JS/TS | `selectArm` |
| `zlinebot-lean/app/src/llm/bandit.ts` | 22 | JS/TS | `observeReward` |
| `zlinebot-lean/app/src/llm/cache.ts` | 6 | JS/TS | `buildPromptCacheKey` |
| `zlinebot-lean/app/src/llm/cache.ts` | 10 | JS/TS | `getCachedResponse` |
| `zlinebot-lean/app/src/llm/cache.ts` | 14 | JS/TS | `setCachedResponse` |
| `zlinebot-lean/app/src/llm/embedding.ts` | 5 | JS/TS | `getProvider` |
| `zlinebot-lean/app/src/llm/embedding.ts` | 14 | JS/TS | `validateText` |
| `zlinebot-lean/app/src/llm/embedding.ts` | 20 | JS/TS | `parseEmbedding` |
| `zlinebot-lean/app/src/llm/embedding.ts` | 39 | JS/TS | `embed` |
| `zlinebot-lean/app/src/llm/llm.service.ts` | 10 | JS/TS | `completeAndMeter` |
| `zlinebot-lean/app/src/llm/router.ts` | 8 | JS/TS | `callOpenAI` |
| `zlinebot-lean/app/src/llm/router.ts` | 38 | JS/TS | `routeLLM` |
| `zlinebot-lean/app/src/llm/router.ts` | 55 | JS/TS | `recordOutcome` |
| `zlinebot-lean/app/src/llm/safety.ts` | 8 | JS/TS | `sanitizePrompt` |
| `zlinebot-lean/app/src/llm/token-optimizer.ts` | 1 | JS/TS | `optimizePrompt` |
| `zlinebot-lean/app/src/llm/tool-runner.ts` | 3 | JS/TS | `validateArgs` |
| `zlinebot-lean/app/src/llm/tool-runner.ts` | 23 | JS/TS | `runTool` |
| `zlinebot-lean/app/src/llm/tools.ts` | 14 | JS/TS | `getBalance` |
| `zlinebot-lean/app/src/routes/line.ts` | 7 | JS/TS | `lineRoutes` |
| `zlinebot-lean/app/src/routes/stream.ts` | 5 | JS/TS | `getQueryParam` |
| `zlinebot-lean/app/src/routes/stream.ts` | 9 | JS/TS | `streamRoutes` |
| `zlinebot-lean/app/src/routes/tiktok.ts` | 6 | JS/TS | `tiktokRoutes` |
| `zlinebot-lean/app/src/sandbox/runner.ts` | 4 | JS/TS | `runWasm` |
| `zlinebot-lean/app/src/server.ts` | 12 | JS/TS | `bootstrap` |
| `zlinebot-lean/app/src/services/line.service.ts` | 5 | JS/TS | `handleLineMessageEvent` |
| `zlinebot-lean/app/src/services/tiktok.service.ts` | 3 | JS/TS | `handleTikTokEvent` |
| `zlinebot-lean/app/src/utils/validator.ts` | 1 | JS/TS | `asObject` |
| `zlinebot-lean/scripts/meta-master-global.sh` | 6 | Shell | `failover` |
| `zlinebot-lean/scripts/meta-master-v3.sh` | 6 | Shell | `check_cost_spike` |
| `zlinebot-lean/scripts/meta-master-v3.sh` | 10 | Shell | `disable_tenant` |
| `zlinebot-lean/scripts/meta-master-v3.sh` | 14 | Shell | `check_region` |
| `zlinebot-lean/scripts/meta-master-v3.sh` | 19 | Shell | `shift_traffic` |
| `zlinebot-lean/scripts/meta-master-v3.sh` | 24 | Shell | `check_lag` |
| `zlinebot-lean/scripts/meta-master-v3.sh` | 28 | Shell | `scale_worker` |
| `zlinebot-lean/scripts/meta-master-v4.sh` | 15 | Shell | `log` |
| `zlinebot-lean/scripts/meta-master-v4.sh` | 19 | Shell | `check_billing_drift` |
| `zlinebot-lean/scripts/meta-master-v4.sh` | 27 | Shell | `check_cost_spike` |
| `zlinebot-lean/scripts/meta-master-v4.sh` | 34 | Shell | `check_region_health` |
| `zlinebot-lean/scripts/meta-master-v4.sh` | 41 | Shell | `check_kafka_lag` |
| `zlinebot-lean/scripts/meta-master-v4.sh` | 51 | Shell | `scale_workers` |
| `zlinebot-lean/scripts/meta-master-v4.sh` | 56 | Shell | `main` |
| `zlinebot-lean/scripts/meta-master.sh` | 6 | Shell | `heal` |

### Automation pipeline files

- `zlinebot-lean/.github/workflows/deploy.yml`
- `zlinebot-lean/docker/docker-compose.yml`
- `zlinebot-lean/infra/docker/docker-compose.yml`
- `zlinebot-lean/infra/k8s/gpu-deployment.yaml`
- `zlinebot-lean/scripts/bootstrap-infra.sh`
- `zlinebot-lean/scripts/install.sh`

## `cvsz/zTTato-Platform`

Status: **blocked**; commit `blocked`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

None detected or repository blocked.

## `cvsz/zvath`

Status: **cloned**; commit `9553b27b0c23`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `apps/api/main.py` | 22 | `/healthz` |
| `apps/api/main.py` | 27 | `/tenants/{org_id}/me` |
| `apps/api/routes/health.py` | 17 | `/health` |
| `apps/api/routes/oauth.py` | 94 | `/url` |
| `apps/api/routes/oauth.py` | 128 | `/callback` |
| `apps/api/routes/oauth.py` | 188 | `/refresh` |
| `apps/api/routes/stripe_webhook.py` | 9 | `/stripe` |
| `apps/api/routes/stripe_webhook.py` | 21 | `/webhook` |
| `apps/web/app/(auth)/login/page.tsx` | 25 | `/api/auth/login` |
| `apps/web/app/dashboard/page.tsx` | 12 | `/api/dashboard` |
| `tests/test_api_security_hardening.py` | 159 | `/stripe/webhook` |
| `tests/test_api_security_hardening.py` | 172 | `/stripe/webhook` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `apps/ai-brain/agent.py` | 10 | Python | `__init__` |
| `apps/ai-brain/agent.py` | 15 | Python | `run_cycle` |
| `apps/ai-brain/agent.py` | 36 | Python | `store` |
| `apps/ai-brain/inference.py` | 6 | Python | `predict` |
| `apps/ai-brain/loop.py` | 6 | Python | `loop` |
| `apps/ai-brain/pipeline.py` | 6 | Python | `generate_pipeline` |
| `apps/ai-brain/swarm/analyst.py` | 4 | Python | `summarize` |
| `apps/ai-brain/swarm/creator.py` | 7 | Python | `generate_caption` |
| `apps/ai-brain/swarm/creator.py` | 10 | Python | `_validate_plan` |
| `apps/ai-brain/swarm/creator.py` | 31 | Python | `create` |
| `apps/ai-brain/swarm/critic.py` | 6 | Python | `approve` |
| `apps/ai-brain/swarm/optimizer.py` | 4 | Python | `variants` |
| `apps/ai-brain/swarm/orchestrator.py` | 14 | Python | `emit` |
| `apps/ai-brain/swarm/orchestrator.py` | 18 | Python | `start_cycle` |
| `apps/ai-brain/swarm/orchestrator.py` | 22 | Python | `route` |
| `apps/ai-brain/swarm/scheduler.py` | 4 | Python | `schedule` |
| `apps/ai-brain/swarm/strategist.py` | 4 | Python | `plan` |
| `apps/ai-brain/trend_engine.py` | 13 | Python | `__init__` |
| `apps/ai-brain/trend_engine.py` | 19 | Python | `tiktok` |
| `apps/ai-brain/trend_engine.py` | 22 | Python | `youtube` |
| `apps/ai-brain/trend_engine.py` | 26 | Python | `_normalize` |
| `apps/ai-brain/trend_engine.py` | 29 | Python | `aggregate` |
| `apps/anti-ban-agent/stealth.py` | 27 | Python | `jitter` |
| `apps/anti-ban-agent/stealth.py` | 33 | Python | `_proxy_config` |
| `apps/anti-ban-agent/stealth.py` | 40 | Python | `safe_post` |
| `apps/api/deps.py` | 8 | Python | `auth_context` |
| `apps/api/deps.py` | 15 | Python | `require_admin_or_god` |
| `apps/api/main.py` | 23 | Python | `healthz` |
| `apps/api/main.py` | 28 | Python | `my_tenant` |
| `apps/api/routes/generate.py` | 16 | Python | `generate` |
| `apps/api/routes/health.py` | 18 | Python | `health` |
| `apps/api/routes/inference.py` | 8 | Python | `predict` |
| `apps/api/routes/oauth.py` | 24 | Python | `_client_id` |
| `apps/api/routes/oauth.py` | 28 | Python | `_client_secret` |
| `apps/api/routes/oauth.py` | 32 | Python | `_redirect_uri` |
| `apps/api/routes/oauth.py` | 36 | Python | `_redis_client` |
| `apps/api/routes/oauth.py` | 40 | Python | `_oauth_state_key` |
| `apps/api/routes/oauth.py` | 44 | Python | `_build_state` |
| `apps/api/routes/oauth.py` | 48 | Python | `_build_pkce_verifier` |
| `apps/api/routes/oauth.py` | 52 | Python | `_build_pkce_challenge` |
| `apps/api/routes/oauth.py` | 57 | Python | `_is_local_env` |
| `apps/api/routes/oauth.py` | 62 | Python | `_persist_state` |
| `apps/api/routes/oauth.py` | 76 | Python | `_consume_state` |
| `apps/api/routes/oauth.py` | 95 | Python | `get_oauth_url` |
| `apps/api/routes/oauth.py` | 129 | Python | `exchange_code` |
| `apps/api/routes/oauth.py` | 177 | Python | `_normalize_token_response` |
| `apps/api/routes/oauth.py` | 189 | Python | `refresh_access_token` |
| `apps/api/routes/stripe_webhook.py` | 12 | Python | `_webhook_secret` |
| `apps/api/routes/stripe_webhook.py` | 16 | Python | `_is_local_env` |
| `apps/api/routes/stripe_webhook.py` | 22 | Python | `webhook` |
| `apps/mobile/src/hooks/useAutomation.ts` | 8 | JS/TS | `runAutoPost` |
| `apps/stealth-agent/engine.py` | 14 | Python | `post` |
| `apps/web/app/(auth)/login/page.tsx` | 19 | JS/TS | `handleSubmit` |
| `apps/web/app/admin/page.tsx` | 65 | JS/TS | `AdminCard` |
| `apps/web/app/api/auth/login/route.ts` | 10 | JS/TS | `POST` |
| `apps/web/app/api/dashboard/route.ts` | 4 | JS/TS | `GET` |
| `apps/web/app/dashboard/page.tsx` | 11 | JS/TS | `fetchData` |
| `apps/web/app/dashboard/page.tsx` | 114 | JS/TS | `Card` |
| `apps/web/app/page.tsx` | 123 | JS/TS | `Metric` |
| `apps/web/components/ai-table.tsx` | 13 | JS/TS | `AITable` |
| `apps/web/components/pricing.tsx` | 3 | JS/TS | `Pricing` |
| `apps/web/components/pricing.tsx` | 31 | JS/TS | `Card` |
| `apps/web/lib/i18n.ts` | 183 | JS/TS | `resolveLocale` |
| `apps/web/lib/i18n.ts` | 188 | JS/TS | `getMessages` |
| `apps/web/middleware.ts` | 8 | JS/TS | `getRole` |
| `apps/web/middleware.ts` | 14 | JS/TS | `isAuthorized` |
| `apps/web/middleware.ts` | 20 | JS/TS | `withLang` |
| `apps/web/middleware.ts` | 27 | JS/TS | `middleware` |
| `apps/worker/analytics_consumer.py` | 17 | Python | `consume_forever` |
| `apps/worker/competitor_scraper.py` | 11 | Python | `_proxy_config` |
| `apps/worker/competitor_scraper.py` | 18 | Python | `fetch_profile` |
| `apps/worker/competitor_scraper.py` | 31 | Python | `extract_signals` |
| `apps/worker/competitor_scraper.py` | 39 | Python | `scan` |
| `apps/worker/event_worker.py` | 11 | Python | `consume` |
| `apps/worker/event_worker.py` | 22 | Python | `process` |
| `apps/worker/kafka_worker.py` | 16 | Python | `_bootstrap_servers` |
| `apps/worker/kafka_worker.py` | 20 | Python | `_jobs_topic` |
| `apps/worker/kafka_worker.py` | 24 | Python | `_dlq_topic` |
| `apps/worker/kafka_worker.py` | 28 | Python | `_worker_group` |
| `apps/worker/kafka_worker.py` | 32 | Python | `_build_consumer` |
| `apps/worker/kafka_worker.py` | 43 | Python | `_build_producer` |
| `apps/worker/kafka_worker.py` | 50 | Python | `process_generate` |
| `apps/worker/kafka_worker.py` | 54 | Python | `process_post` |
| `apps/worker/kafka_worker.py` | 58 | Python | `process` |
| `apps/worker/kafka_worker.py` | 74 | Python | `handle` |
| `apps/worker/kafka_worker.py` | 78 | Python | `_send_to_dlq` |
| `apps/worker/kafka_worker.py` | 89 | Python | `worker` |
| `apps/worker/queue.py` | 17 | Python | `labels` |
| `apps/worker/queue.py` | 20 | Python | `inc` |
| `apps/worker/queue.py` | 23 | Python | `Counter` |
| `apps/worker/queue.py` | 27 | Python | `_redis_host` |
| `apps/worker/queue.py` | 31 | Python | `_max_retry` |
| `apps/worker/queue.py` | 35 | Python | `_sleep_seconds` |
| `apps/worker/queue.py` | 39 | Python | `_max_queue_length` |
| `apps/worker/queue.py` | 43 | Python | `_idempotency_ttl_seconds` |
| `apps/worker/queue.py` | 56 | Python | `process` |
| `apps/worker/queue.py` | 63 | Python | `classify_error` |
| `apps/worker/queue.py` | 75 | Python | `_idempotency_key` |
| `apps/worker/queue.py` | 83 | Python | `enqueue_job` |
| `apps/worker/queue.py` | 100 | Python | `worker` |
| `apps/worker/trend_ingest.py` | 19 | Python | `_extract_lines` |
| `apps/worker/trend_ingest.py` | 26 | Python | `fetch_tiktok_trends` |
| `apps/worker/trend_ingest.py` | 36 | Python | `loop` |
| `infra/kubeflow/pipeline.py` | 17 | Python | `train_pipeline` |
| `packages/ai/bandit/thompson.py` | 9 | Python | `__init__` |
| `packages/ai/bandit/thompson.py` | 15 | Python | `select` |
| `packages/ai/bandit/thompson.py` | 22 | Python | `update` |
| `packages/ai/deep_model.py` | 5 | Python | `build_viral_net` |
| `packages/ai/deep_model.py` | 7 | Python | `__init__` |
| `packages/ai/deep_model.py` | 16 | Python | `forward` |
| `packages/ai/embedding.py` | 8 | Python | `_model` |
| `packages/ai/embedding.py` | 12 | Python | `embed` |
| `packages/ai/feature_store/offline_loader.py` | 11 | Python | `load_parquet` |
| `packages/ai/feature_store/offline_store.py` | 16 | Python | `write_parquet` |
| `packages/ai/feature_store/repo.py` | 6 | Python | `get_features` |
| `packages/ai/infer.py` | 16 | Python | `_sanitize_text` |
| `packages/ai/infer.py` | 22 | Python | `predict` |
| `packages/ai/mlops/canary.py` | 18 | Python | `__init__` |
| `packages/ai/mlops/canary.py` | 24 | Python | `route` |
| `packages/ai/mlops/causal.py` | 8 | Python | `_mean` |
| `packages/ai/mlops/causal.py` | 15 | Python | `estimate_uplift` |
| `packages/ai/mlops/compliance.py` | 21 | Python | `passed` |
| `packages/ai/mlops/compliance.py` | 25 | Python | `failed` |
| `packages/ai/mlops/compliance.py` | 29 | Python | `status` |
| `packages/ai/mlops/compliance.py` | 46 | Python | `certified` |
| `packages/ai/mlops/compliance.py` | 59 | Python | `assess` |
| `packages/ai/mlops/compliance.py` | 80 | Python | `__init__` |
| `packages/ai/mlops/compliance.py` | 83 | Python | `run` |
| `packages/ai/mlops/drift.py` | 18 | Python | `drifted` |
| `packages/ai/mlops/drift.py` | 25 | Python | `__init__` |
| `packages/ai/mlops/drift.py` | 29 | Python | `_to_list` |
| `packages/ai/mlops/drift.py` | 35 | Python | `score` |
| `packages/ai/mlops/drift.py` | 51 | Python | `detect` |
| `packages/ai/mlops/explainability.py` | 18 | Python | `explain_with_shap` |
| `packages/ai/mlops/explainability.py` | 23 | Python | `explain_with_lime` |
| `packages/ai/mlops/federated.py` | 25 | Python | `__init__` |
| `packages/ai/mlops/federated.py` | 29 | Python | `aggregate` |
| `packages/ai/mlops/federation_network.py` | 36 | Python | `__init__` |
| `packages/ai/mlops/federation_network.py` | 41 | Python | `run_round` |
| `packages/ai/mlops/governance.py` | 43 | Python | `__init__` |
| `packages/ai/mlops/governance.py` | 52 | Python | `evaluate` |
| `packages/ai/mlops/homomorphic.py` | 26 | Python | `__init__` |
| `packages/ai/mlops/homomorphic.py` | 29 | Python | `encrypt` |
| `packages/ai/mlops/homomorphic.py` | 32 | Python | `decrypt` |
| `packages/ai/mlops/homomorphic.py` | 35 | Python | `add` |
| `packages/ai/mlops/homomorphic.py` | 38 | Python | `linear_infer` |
| `packages/ai/mlops/lineage.py` | 20 | Python | `create_event` |
| `packages/ai/mlops/lineage.py` | 28 | Python | `to_openlineage_payload` |
| `packages/ai/mlops/lineage.py` | 36 | Python | `to_marquez_payload` |
| `packages/ai/mlops/monitoring.py` | 16 | Python | `message` |
| `packages/ai/mlops/monitoring.py` | 26 | Python | `__init__` |
| `packages/ai/mlops/monitoring.py` | 30 | Python | `evaluate` |
| `packages/ai/mlops/pipeline.py` | 21 | Python | `as_dict` |
| `packages/ai/mlops/policy.py` | 18 | Python | `__init__` |
| `packages/ai/mlops/policy.py` | 21 | Python | `register_deny_rule` |
| `packages/ai/mlops/policy.py` | 29 | Python | `evaluate` |
| `packages/ai/mlops/privacy.py` | 35 | Python | `__init__` |
| `packages/ai/mlops/privacy.py` | 39 | Python | `privatize_gradient` |
| `packages/ai/mlops/privacy.py` | 57 | Python | `detect` |
| `packages/ai/mlops/privacy.py` | 63 | Python | `anonymize` |
| `packages/ai/mlops/safe_rl.py` | 17 | Python | `__init__` |
| `packages/ai/mlops/safe_rl.py` | 23 | Python | `enforce` |
| `packages/ai/mlops/security.py` | 35 | Python | `score` |
| `packages/ai/mlops/security.py` | 44 | Python | `__init__` |
| `packages/ai/mlops/security.py` | 47 | Python | `evaluate_prompt` |
| `packages/ai/mlops/security.py` | 51 | Python | `run` |
| `packages/ai/mlops/security.py` | 76 | Python | `__init__` |
| `packages/ai/mlops/security.py` | 80 | Python | `defend` |
| `packages/ai/mlops/security.py` | 125 | Python | `__init__` |
| `packages/ai/mlops/security.py` | 129 | Python | `ingest` |
| `packages/ai/mlops/security.py` | 140 | Python | `lookup` |
| `packages/ai/mlops/security.py` | 143 | Python | `top_indicators` |
| `packages/ai/mlops/security.py` | 151 | Python | `__init__` |
| `packages/ai/mlops/security.py` | 156 | Python | `ingest` |
| `packages/ai/mlops/security.py` | 161 | Python | `summarize` |
| `packages/ai/mlops/security.py` | 172 | Python | `__init__` |
| `packages/ai/mlops/security.py` | 177 | Python | `_legacy_signature` |
| `packages/ai/mlops/security.py` | 180 | Python | `_signature` |
| `packages/ai/mlops/security.py` | 184 | Python | `embed` |
| `packages/ai/mlops/security.py` | 190 | Python | `verify` |
| `packages/ai/online/shadow.py` | 6 | Python | `should_promote` |
| `packages/ai/online/train.py` | 9 | Python | `train_from_stream` |
| `packages/ai/recommendations.py` | 11 | Python | `cosine_similarity` |
| `packages/ai/recommendations.py` | 20 | Python | `cluster_vectors` |
| `packages/ai/recommendations.py` | 56 | Python | `recommend_from_clusters` |
| `packages/ai/recommendations.py` | 83 | Python | `iter_cluster_groups` |
| `packages/ai/registry/load_model.py` | 8 | Python | `load_latest` |
| `packages/ai/registry/log_model.py` | 9 | Python | `log` |
| `packages/ai/rl/agent.py` | 11 | Python | `__init__` |
| `packages/ai/rl/agent.py` | 15 | Python | `forward` |
| `packages/ai/rl/agent.py` | 20 | Python | `__init__` |
| `packages/ai/rl/agent.py` | 24 | Python | `act` |
| `packages/ai/rl/agent.py` | 30 | Python | `update` |
| `packages/ai/rl/train_loop.py` | 12 | Python | `train` |
| `packages/ai/search.py` | 6 | Python | `search` |
| `packages/ai/service.py` | 10 | Python | `__init__` |
| `packages/ai/service.py` | 13 | Python | `embed` |
| `packages/ai/service.py` | 16 | Python | `score` |
| `packages/ai/train.py` | 23 | Python | `train` |
| `packages/ai/train_dl.py` | 9 | Python | `train_step` |
| `packages/ai/training/dataset.py` | 6 | Python | `build_dataset` |
| `packages/ai/training/train.py` | 6 | Python | `fine_tune` |
| `packages/ai/vector_store.py` | 7 | Python | `insert_embeddings` |
| `packages/ai/viral_score.py` | 6 | Python | `score` |
| `packages/analytics/ab.py` | 6 | Python | `variant` |
| `packages/analytics/competitor.py` | 4 | Python | `extract_trend_patterns` |
| `packages/analytics/competitor_signals.py` | 4 | Python | `normalize` |
| `packages/analytics/competitor_signals.py` | 21 | Python | `top_patterns` |
| `packages/analytics/engine.py` | 4 | Python | `compute_ctr` |
| `packages/analytics/engine.py` | 14 | Python | `detect_winner` |
| `packages/analytics/producer.py` | 13 | Python | `send_event` |
| `packages/anti_ban/behavior.py` | 7 | Python | `human_delay` |
| `packages/anti_ban/behavior.py` | 11 | Python | `scroll_simulation` |
| `packages/anti_ban/fingerprint.py` | 6 | Python | `generate_fingerprint` |
| `packages/billing/ledger.py` | 15 | Python | `charge` |
| `packages/billing/service.py` | 21 | Python | `__init__` |
| `packages/billing/service.py` | 27 | Python | `create_subscription` |
| `packages/billing/service.py` | 31 | Python | `log_payment` |
| `packages/billing/service.py` | 47 | Python | `charge` |
| `packages/billing/thai_gateway.py` | 10 | Python | `configure` |
| `packages/billing/thai_gateway.py` | 20 | Python | `charge` |
| `packages/core/audit.py` | 8 | Python | `log` |
| `packages/core/audit.py` | 27 | Python | `risk_level` |
| `packages/core/audit.py` | 35 | Python | `build_compliance_summary` |
| `packages/core/backpressure.py` | 12 | Python | `allow_enqueue` |
| `packages/core/backpressure.py` | 17 | Python | `allow` |
| `packages/core/budget.py` | 11 | Python | `__init__` |
| `packages/core/budget.py` | 14 | Python | `get_usage` |
| `packages/core/budget.py` | 21 | Python | `check` |
| `packages/core/circuit.py` | 12 | Python | `__init__` |
| `packages/core/circuit.py` | 18 | Python | `call` |
| `packages/core/confidential.py` | 21 | Python | `__init__` |
| `packages/core/confidential.py` | 34 | Python | `evaluate` |
| `packages/core/crdt.py` | 24 | Python | `_dominates` |
| `packages/core/crdt.py` | 29 | Python | `merge` |
| `packages/core/crdt.py` | 37 | Python | `update` |
| `packages/core/crypto.py` | 15 | Python | `_key_bytes` |
| `packages/core/crypto.py` | 22 | Python | `_keystream` |
| `packages/core/crypto.py` | 32 | Python | `_xor_bytes` |
| `packages/core/crypto.py` | 36 | Python | `encrypt` |
| `packages/core/crypto.py` | 46 | Python | `decrypt` |
| `packages/core/crypto.py` | 72 | Python | `rotate_token` |
| `packages/core/db_ctx.py` | 10 | Python | `tenant_conn` |
| `packages/core/db_ctx.py` | 22 | Python | `tenant_context` |
| `packages/core/db_ctx.py` | 29 | Python | `tenant_session` |
| `packages/core/db_ctx.py` | 35 | Python | `set_org` |
| `packages/core/db_ctx.py` | 41 | Python | `reset_org` |
| `packages/core/geo_router.py` | 10 | Python | `route_request` |
| `packages/core/guard.py` | 9 | Python | `guarded_call` |
| `packages/core/idempotency.py` | 21 | Python | `__init__` |
| `packages/core/idempotency.py` | 25 | Python | `set_if_absent` |
| `packages/core/idempotency.py` | 40 | Python | `__init__` |
| `packages/core/idempotency.py` | 52 | Python | `check_and_set` |
| `packages/core/idempotency.py` | 62 | Python | `key` |
| `packages/core/idempotency.py` | 70 | Python | `acquire` |
| `packages/core/idempotency.py` | 74 | Python | `check_and_set` |
| `packages/core/identity.py` | 16 | Python | `uri` |
| `packages/core/identity.py` | 20 | Python | `parse_spiffe_id` |
| `packages/core/identity.py` | 44 | Python | `__init__` |
| `packages/core/identity.py` | 47 | Python | `verify` |
| `packages/core/quantum_safe.py` | 34 | Python | `derive_hybrid_key` |
| `packages/core/quantum_safe.py` | 64 | Python | `encrypt` |
| `packages/core/quantum_safe.py` | 73 | Python | `decrypt` |
| `packages/core/quantum_safe.py` | 90 | Python | `_keystream` |
| `packages/core/quota.py` | 9 | Python | `__init__` |
| `packages/core/quota.py` | 12 | Python | `enforce` |
| `packages/core/quota_global.py` | 10 | Python | `enforce_global` |
| `packages/core/rate_limit.py` | 10 | Python | `allow` |
| `packages/core/secrets.py` | 10 | Python | `_client` |
| `packages/core/secrets.py` | 37 | Python | `_cache_ttl_seconds` |
| `packages/core/secrets.py` | 41 | Python | `invalidate_secret` |
| `packages/core/secrets.py` | 45 | Python | `get_secret` |
| `packages/core/security.py` | 29 | Python | `_jwt_secret` |
| `packages/core/security.py` | 36 | Python | `_jwks_url` |
| `packages/core/security.py` | 40 | Python | `_get_jwks_key` |
| `packages/core/security.py` | 59 | Python | `create_token` |
| `packages/core/security.py` | 71 | Python | `verify_token` |
| `packages/integrations/publisher.py` | 10 | Python | `post` |
| `packages/integrations/publisher.py` | 20 | Python | `post` |
| `packages/integrations/token_store.py` | 7 | Python | `expired` |
| `packages/integrations/token_store.py` | 11 | Python | `is_expired` |
| `packages/integrations/token_store.py` | 15 | Python | `refresh` |
| `packages/observability/metrics.py` | 14 | Python | `inc` |
| `packages/observability/metrics.py` | 21 | Python | `__enter__` |
| `packages/observability/metrics.py` | 25 | Python | `__exit__` |
| `packages/observability/metrics.py` | 32 | Python | `time` |
| `packages/observability/metrics.py` | 36 | Python | `counter_factory` |
| `packages/observability/metrics.py` | 40 | Python | `histogram_factory` |
| `packages/observability/metrics.py` | 52 | Python | `track` |
| `packages/observability/metrics.py` | 56 | Python | `wrapper` |
| `packages/observability/tracer.py` | 10 | Python | `init_tracer` |
| `packages/observability/tracing.py` | 13 | Python | `trace_span` |
| `packages/observability/tracing.py` | 14 | Python | `decorator` |
| `packages/observability/tracing.py` | 16 | Python | `wrapper` |
| `packages/observability/tracing.py` | 28 | Python | `span` |
| `packages/queue/producer.py` | 17 | Python | `send` |
| `packages/revenue/cac.py` | 4 | Python | `calculate_cac` |
| `packages/revenue/cohort.py` | 6 | Python | `retention_rate` |
| `packages/revenue/engine.py` | 22 | Python | `__init__` |
| `packages/revenue/engine.py` | 27 | Python | `track_usage` |

_Display capped at 300 of 446 detected functions/methods._

### Automation pipeline files

- `.github/workflows/ci.yml`
- `.github/workflows/nodejs-webpack.yml`
- `.github/workflows/pylint.yml`
- `.github/workflows/python-version-consistency.yml`
- `apps/ai-brain/pipeline.py`
- `infra/helm/viral/templates/deployment.yaml`
- `infra/kubeflow/pipeline.py`
- `packages/ai/mlops/pipeline.py`
- `scripts/install.sh`

## `cvsz/tiktok-shop-bot`

Status: **cloned**; commit `0e6128856c86`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `src/cli.py` | 11 | Python | `cmd_validate` |
| `src/cli.py` | 18 | Python | `cmd_render` |
| `src/cli.py` | 35 | Python | `cmd_send` |
| `src/cli.py` | 49 | Python | `cmd_followup` |
| `src/cli.py` | 58 | Python | `cmd_update` |
| `src/cli.py` | 66 | Python | `cmd_report` |
| `src/cli.py` | 80 | Python | `build` |
| `src/cli.py` | 122 | Python | `main` |
| `src/dedupe.py` | 4 | Python | `_norm_id` |
| `src/dedupe.py` | 9 | Python | `dedupe_df` |
| `src/outreach.py` | 7 | Python | `prepare_vars` |
| `src/outreach.py` | 16 | Python | `send_email_batch` |
| `src/outreach.py` | 39 | Python | `manual_dm_checklist` |
| `src/rate_limit.py` | 5 | Python | `_now_hour_bucket` |
| `src/rate_limit.py` | 6 | Python | `_now_day_bucket` |
| `src/rate_limit.py` | 8 | Python | `in_quiet_hours` |
| `src/rate_limit.py` | 14 | Python | `natural_delay` |
| `src/rate_limit.py` | 20 | Python | `__init__` |
| `src/rate_limit.py` | 25 | Python | `_roll` |
| `src/rate_limit.py` | 28 | Python | `allow` |
| `src/rate_limit.py` | 31 | Python | `mark` |
| `src/templating.py` | 5 | Python | `render_template` |

### Automation pipeline files

None detected or repository blocked.

## `cvsz/tiktokshop-api-client`

Status: **cloned**; commit `68dc36ff8524`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `README.MD` | 23 | `/api/v2/token/get` |
| `README.MD` | 46 | `/api/v2/token/refresh` |
| `README.MD` | 70 | `/api/shop/get_authorized_shop` |
| `src/Request/General/GeneralWithOutBody.php` | 27 | `/api/shop/get_authorized_shop` |
| `tests/AuthRequestTest.php` | 39 | `/api/v2/token/get` |
| `tests/AuthRequestTest.php` | 61 | `/api/v2/token/refresh` |
| `tests/GeneralRequestTest.php` | 33 | `/api/shop/get_authorized_shop` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `src/Common/SignGenerator.php` | 7 | PHP | `generateSign` |
| `src/Common/TiktokShopConfig.php` | 29 | PHP | `__construct` |
| `src/Common/TiktokShopConfig.php` | 42 | PHP | `getAppKey` |
| `src/Common/TiktokShopConfig.php` | 50 | PHP | `setAppKey` |
| `src/Common/TiktokShopConfig.php` | 60 | PHP | `getAccessToken` |
| `src/Common/TiktokShopConfig.php` | 68 | PHP | `setAccessToken` |
| `src/Common/TiktokShopConfig.php` | 76 | PHP | `getRefreshToken` |
| `src/Common/TiktokShopConfig.php` | 84 | PHP | `setRefreshToken` |
| `src/Common/TiktokShopConfig.php` | 93 | PHP | `getShopId` |
| `src/Common/TiktokShopConfig.php` | 101 | PHP | `setShopId` |
| `src/Common/TiktokShopConfig.php` | 109 | PHP | `getSecretKey` |
| `src/Common/TiktokShopConfig.php` | 117 | PHP | `setSecretKey` |
| `src/Common/TiktokShopConfig.php` | 125 | PHP | `getTimeOut` |
| `src/Common/TiktokShopConfig.php` | 133 | PHP | `setTimeOut` |
| `src/Request/Auth/AuthWithOutBody.php` | 20 | PHP | `makeGetMethod` |
| `src/Request/General/GeneralWithBody.php` | 19 | PHP | `makeMethod` |
| `src/Request/General/GeneralWithOutBody.php` | 19 | PHP | `makeGetMethod` |
| `src/Resource/Auth/TiktokShopAuthResource.php` | 16 | PHP | `httpCallGet` |
| `src/Resource/Auth/TiktokShopAuthResource.php` | 29 | PHP | `generateAuthUrl` |
| `src/Resource/General/TiktokShopGeneralResource.php` | 16 | PHP | `httpCallGet` |
| `src/Resource/General/TiktokShopGeneralResource.php` | 26 | PHP | `httpCallPost` |
| `src/Resource/General/TiktokShopGeneralResource.php` | 36 | PHP | `httpCallPut` |
| `src/Resource/General/TiktokShopGeneralResource.php` | 47 | PHP | `httpCallPatch` |
| `src/Resource/General/TiktokShopGeneralResource.php` | 58 | PHP | `httpCallDelete` |
| `tests/AuthRequestTest.php` | 12 | PHP | `setUp` |
| `tests/AuthRequestTest.php` | 18 | PHP | `tearDown` |
| `tests/AuthRequestTest.php` | 23 | PHP | `test_generate_auth_url` |
| `tests/AuthRequestTest.php` | 30 | PHP | `test_generate_access_token` |
| `tests/AuthRequestTest.php` | 52 | PHP | `test_refresh_token` |
| `tests/GeneralRequestTest.php` | 12 | PHP | `setUp` |
| `tests/GeneralRequestTest.php` | 18 | PHP | `tearDown` |
| `tests/GeneralRequestTest.php` | 23 | PHP | `test_get_authorized_shop` |

### Automation pipeline files

None detected or repository blocked.

## `cvsz/tiktok-shop-sdk`

Status: **cloned**; commit `0c53da6cbba9`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `packages/sdk/src/client/__tests__/Request.test.ts` | 333 | `/api/test` |
| `packages/sdk/src/client/__tests__/Request.test.ts` | 362 | `/api/test` |
| `packages/sdk/src/client/__tests__/Request.test.ts` | 397 | `/api/test` |
| `packages/sdk/src/client/__tests__/Request.test.ts` | 439 | `/api/private` |
| `packages/sdk/src/client/__tests__/authRequest.test.ts` | 54 | `/api/v2/token/get` |
| `packages/sdk/src/client/__tests__/authRequest.test.ts` | 111 | `/api/v2/token/refresh` |
| `packages/sdk/src/client/__tests__/authRequest.test.ts` | 171 | `/api/v2/token/get` |
| `packages/sdk/src/client/__tests__/authRequest.test.ts` | 229 | `/api/v2/token/get` |
| `packages/sdk/src/modules/AuthModule.ts` | 29 | `/api/v2/token/get` |
| `packages/sdk/src/modules/AuthModule.ts` | 51 | `/api/v2/token/refresh` |
| `packages/sdk/src/modules/ShopModule.ts` | 27 | `/authorization/202309/shops` |
| `packages/sdk/src/modules/ShopModule.ts` | 42 | `/authorization/202405/category_assets` |
| `packages/sdk/src/modules/__tests__/auth.test.ts` | 41 | `/api/v2/token/get` |
| `packages/sdk/src/modules/__tests__/auth.test.ts` | 70 | `/api/v2/token/refresh` |
| `packages/sdk/src/modules/__tests__/shopModule.test.ts` | 25 | `/authorization/202309/shops` |
| `packages/sdk/src/modules/__tests__/shopModule.test.ts` | 44 | `/authorization/202405/category_assets` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2 | JS/TS | `makeMap` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9 | JS/TS | `NOOP` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11 | JS/TS | `NO` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12 | JS/TS | `isOn` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 14 | JS/TS | `isModelListener` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 16 | JS/TS | `remove` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 23 | JS/TS | `hasOwn` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 25 | JS/TS | `isMap` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 26 | JS/TS | `isSet` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 27 | JS/TS | `isDate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 28 | JS/TS | `isRegExp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 29 | JS/TS | `isFunction` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 30 | JS/TS | `isString` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 31 | JS/TS | `isSymbol` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 32 | JS/TS | `isObject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 33 | JS/TS | `isPromise` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 37 | JS/TS | `toTypeString` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 38 | JS/TS | `toRawType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 41 | JS/TS | `isPlainObject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 42 | JS/TS | `isIntegerKey` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 50 | JS/TS | `cacheStringFunction` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 76 | JS/TS | `hasChanged` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 77 | JS/TS | `invokeArrayFns` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 82 | JS/TS | `def` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 90 | JS/TS | `looseToNumber` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 94 | JS/TS | `toNumber` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 99 | JS/TS | `getGlobalThis` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 104 | JS/TS | `normalizeStyle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 124 | JS/TS | `parseStringStyle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 134 | JS/TS | `stringifyStyle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 147 | JS/TS | `normalizeClass` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 167 | JS/TS | `normalizeProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 191 | JS/TS | `includeBooleanAttr` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 203 | JS/TS | `isRenderableAttrValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 211 | JS/TS | `getEscapedCssVarName` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 217 | JS/TS | `looseCompareArrays` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 225 | JS/TS | `looseEqual` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 263 | JS/TS | `looseIndexOf` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 266 | JS/TS | `isRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 269 | JS/TS | `toDisplayString` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 272 | JS/TS | `replacer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 296 | JS/TS | `stringifySymbol` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 304 | JS/TS | `normalizeCssVarValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 323 | JS/TS | `warn` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 328 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 345 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 362 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 378 | JS/TS | `run` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 395 | JS/TS | `on` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 405 | JS/TS | `off` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 411 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 440 | JS/TS | `effectScope` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 443 | JS/TS | `getCurrentScope` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 446 | JS/TS | `onScopeDispose` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 458 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 470 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 473 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 485 | JS/TS | `notify` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 493 | JS/TS | `run` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 518 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 529 | JS/TS | `trigger` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 541 | JS/TS | `runIfDirty` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 553 | JS/TS | `batch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 563 | JS/TS | `startBatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 566 | JS/TS | `endBatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 601 | JS/TS | `prepareDeps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 608 | JS/TS | `cleanupDeps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 628 | JS/TS | `isDirty` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 639 | JS/TS | `refreshComputed` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 675 | JS/TS | `removeSub` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 701 | JS/TS | `removeDep` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 712 | JS/TS | `effect` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 730 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 735 | JS/TS | `pauseTracking` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 739 | JS/TS | `resetTracking` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 743 | JS/TS | `cleanupEffect` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 758 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 767 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 780 | JS/TS | `track` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 824 | JS/TS | `trigger` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 829 | JS/TS | `notify` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 857 | JS/TS | `addSub` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 888 | JS/TS | `track` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 911 | JS/TS | `trigger` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 917 | JS/TS | `run` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 982 | JS/TS | `getDepFromReactive` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 986 | JS/TS | `reactiveReadArray` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 992 | JS/TS | `shallowReadArray` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 996 | JS/TS | `toWrapped` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1007 | JS/TS | `concat` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1012 | JS/TS | `entries` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1018 | JS/TS | `every` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1021 | JS/TS | `filter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1031 | JS/TS | `find` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1041 | JS/TS | `findIndex` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1044 | JS/TS | `findLast` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1054 | JS/TS | `findLastIndex` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1058 | JS/TS | `forEach` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1061 | JS/TS | `includes` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1064 | JS/TS | `indexOf` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1067 | JS/TS | `join` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1071 | JS/TS | `lastIndexOf` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1074 | JS/TS | `map` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1077 | JS/TS | `pop` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1080 | JS/TS | `push` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1083 | JS/TS | `reduce` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1086 | JS/TS | `reduceRight` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1089 | JS/TS | `shift` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1093 | JS/TS | `some` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1096 | JS/TS | `splice` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1099 | JS/TS | `toReversed` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1102 | JS/TS | `toSorted` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1105 | JS/TS | `toSpliced` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1108 | JS/TS | `unshift` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1111 | JS/TS | `values` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1115 | JS/TS | `iterator` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1131 | JS/TS | `apply` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1154 | JS/TS | `reduce` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1170 | JS/TS | `searchProxy` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1180 | JS/TS | `noTracking` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1192 | JS/TS | `hasOwnProperty2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1199 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1203 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1258 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1261 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1301 | JS/TS | `deleteProperty` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1310 | JS/TS | `has` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1317 | JS/TS | `ownKeys` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1327 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1330 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1339 | JS/TS | `deleteProperty` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1353 | JS/TS | `toShallow` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1354 | JS/TS | `getProto` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1355 | JS/TS | `createIterableMethod` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1371 | JS/TS | `next` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1385 | JS/TS | `createReadonlyMethod` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1397 | JS/TS | `createInstrumentations` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1399 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1424 | JS/TS | `has` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1436 | JS/TS | `forEach` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1455 | JS/TS | `add` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1468 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1490 | JS/TS | `delete` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1507 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1536 | JS/TS | `createInstrumentationGetter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1565 | JS/TS | `checkIdentityKeys` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1578 | JS/TS | `targetTypeMap` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1592 | JS/TS | `getTargetType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1595 | JS/TS | `reactive` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1607 | JS/TS | `shallowReactive` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1616 | JS/TS | `readonly` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1625 | JS/TS | `shallowReadonly` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1634 | JS/TS | `createReactiveObject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1663 | JS/TS | `isReactive` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1669 | JS/TS | `isReadonly` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1672 | JS/TS | `isShallow` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1675 | JS/TS | `isProxy` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1678 | JS/TS | `toRaw` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1682 | JS/TS | `markRaw` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1688 | JS/TS | `toReactive` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1689 | JS/TS | `toReadonly` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1690 | JS/TS | `isRef2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1693 | JS/TS | `ref` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1696 | JS/TS | `shallowRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1699 | JS/TS | `createRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1706 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1747 | JS/TS | `triggerRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1761 | JS/TS | `unref` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1764 | JS/TS | `toValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1779 | JS/TS | `proxyRefs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1783 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1798 | JS/TS | `customRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1801 | JS/TS | `toRefs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1812 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1850 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1860 | JS/TS | `toRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1871 | JS/TS | `propertyToRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1875 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1893 | JS/TS | `notify` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1921 | JS/TS | `computed` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1951 | JS/TS | `getCurrentWatcher` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1954 | JS/TS | `onWatcherCleanup` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1965 | JS/TS | `watch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1967 | JS/TS | `warnInvalidSource` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 1974 | JS/TS | `reactiveGetter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2038 | JS/TS | `watchHandle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2052 | JS/TS | `job` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2121 | JS/TS | `traverse` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2156 | JS/TS | `pushWarningContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2159 | JS/TS | `popWarningContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2163 | JS/TS | `warn$1` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2200 | JS/TS | `getComponentTrace` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2221 | JS/TS | `formatTrace` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2229 | JS/TS | `formatTraceEntry` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2240 | JS/TS | `formatProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2251 | JS/TS | `formatProp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2267 | JS/TS | `assertNumber` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2340 | JS/TS | `callWithErrorHandling` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2347 | JS/TS | `callWithAsyncErrorHandling` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2369 | JS/TS | `handleError` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2400 | JS/TS | `logError` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2429 | JS/TS | `nextTick` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2433 | JS/TS | `findInsertionIndex` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2448 | JS/TS | `queueJob` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2462 | JS/TS | `queueFlush` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2467 | JS/TS | `queuePostFlushCb` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2480 | JS/TS | `flushPreFlushCbs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2505 | JS/TS | `flushPostFlushCbs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2534 | JS/TS | `getId` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2535 | JS/TS | `flushJobs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2576 | JS/TS | `checkRecursiveUpdates` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2601 | JS/TS | `registerHMR` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2610 | JS/TS | `unregisterHMR` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2613 | JS/TS | `createRecord` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2623 | JS/TS | `normalizeClassComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2626 | JS/TS | `rerender` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2645 | JS/TS | `reload` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2695 | JS/TS | `updateComponentDef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2703 | JS/TS | `tryWrap` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2718 | JS/TS | `emit$1` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2725 | JS/TS | `setDevtoolsHook$1` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2757 | JS/TS | `devtoolsInitApp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2765 | JS/TS | `devtoolsUnmountApp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2780 | JS/TS | `devtoolsComponentRemoved` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2786 | JS/TS | `createDevtoolsComponentHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2805 | JS/TS | `createDevtoolsPerformanceHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2810 | JS/TS | `devtoolsComponentEmit` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2821 | JS/TS | `setCurrentRenderingInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2827 | JS/TS | `pushScopeId` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2830 | JS/TS | `popScopeId` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2833 | JS/TS | `withScopeId` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2834 | JS/TS | `withCtx` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2839 | JS/TS | `renderFnWithContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2863 | JS/TS | `validateDirectiveName` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2868 | JS/TS | `withDirectives` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2899 | JS/TS | `invokeDirectiveHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2921 | JS/TS | `isTeleport` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2922 | JS/TS | `isTeleportDisabled` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2923 | JS/TS | `isTeleportDeferred` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2924 | JS/TS | `isTargetSVG` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2925 | JS/TS | `isTargetMathML` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2926 | JS/TS | `resolveTarget` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2953 | JS/TS | `process` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2971 | JS/TS | `mount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 2985 | JS/TS | `mountToTarget` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3125 | JS/TS | `remove` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3157 | JS/TS | `moveTeleport` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3182 | JS/TS | `hydrateTeleport` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3185 | JS/TS | `hydrateDisabledTeleport` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3251 | JS/TS | `updateCssVars` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3269 | JS/TS | `prepareAnchor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3281 | JS/TS | `useTransitionState` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3317 | JS/TS | `recursiveGetSubtree` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3324 | JS/TS | `setup` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3405 | JS/TS | `findNonCommentChild` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3426 | JS/TS | `getLeavingNodesForType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3435 | JS/TS | `resolveTransitionHooks` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3455 | JS/TS | `callHook3` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3463 | JS/TS | `callAsyncHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3475 | JS/TS | `beforeEnter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3496 | JS/TS | `enter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3529 | JS/TS | `leave` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3563 | JS/TS | `clone` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3577 | JS/TS | `emptyPlaceholder` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3584 | JS/TS | `getInnerChild$1` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3604 | JS/TS | `setTransitionHooks` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3615 | JS/TS | `getTransitionRawChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3637 | JS/TS | `defineComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3644 | JS/TS | `useId` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3655 | JS/TS | `markAsyncBoundary` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3659 | JS/TS | `useTemplateRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3686 | JS/TS | `setRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3731 | JS/TS | `canSetRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3755 | JS/TS | `doSet` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3793 | JS/TS | `job` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3809 | JS/TS | `invalidatePendingSetRef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3817 | JS/TS | `logMismatchError` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3824 | JS/TS | `isSVGContainer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3825 | JS/TS | `isMathMLContainer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3826 | JS/TS | `getContainerType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3832 | JS/TS | `isComment` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3833 | JS/TS | `createHydrationFunctions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3847 | JS/TS | `hydrate2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3861 | JS/TS | `hydrateNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 3864 | JS/TS | `onMismatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4039 | JS/TS | `hydrateElement` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4165 | JS/TS | `hydrateChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4227 | JS/TS | `hydrateFragment` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4250 | JS/TS | `handleMismatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4298 | JS/TS | `locateClosingAnchor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4315 | JS/TS | `replaceNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4328 | JS/TS | `isTemplateNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4333 | JS/TS | `propHasMismatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4392 | JS/TS | `format` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4406 | JS/TS | `toClassSet` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4409 | JS/TS | `isSetEqual` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4420 | JS/TS | `toStyleMap` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4432 | JS/TS | `isMapEqual` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4443 | JS/TS | `resolveCssVars` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4479 | JS/TS | `isMismatchAllowed` |

_Display capped at 300 of 1921 detected functions/methods._

### Automation pipeline files

- `.github/workflows/ci.yml`
- `.github/workflows/publish.yml`
- `apps/docs/core-api/finance/get-transactions-by-order.md`
- `apps/docs/core-api/finance/get-transactions-by-statement.md`
- `apps/examples/src/finances/getTransactionsByOrder.ts`
- `apps/examples/src/finances/getTransactionsByStatement.ts`

## `cvsz/tiktokshop-php`

Status: **cloned**; commit `dbbec213d9d1`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `src/Auth.php` | 52 | `/api/v2/token/get` |
| `src/Auth.php` | 71 | `/api/v2/token/refresh` |
| `tests/AuthTest.php` | 56 | `/api/v2/token/refresh` |
| `tests/AuthTest.php` | 83 | `/api/v2/token/get` |
| `tests/Resources/EventTest.php` | 26 | `/webhooks` |
| `tests/Resources/EventTest.php` | 32 | `/webhooks` |
| `tests/Resources/EventTest.php` | 38 | `/webhooks` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `src/Auth.php` | 25 | PHP | `__construct` |
| `src/Auth.php` | 33 | PHP | `createAuthRequest` |
| `src/Auth.php` | 50 | PHP | `getToken` |
| `src/Auth.php` | 69 | PHP | `refreshNewToken` |
| `src/Client.php` | 100 | PHP | `__construct` |
| `src/Client.php` | 109 | PHP | `useSandboxMode` |
| `src/Client.php` | 117 | PHP | `useVersion` |
| `src/Client.php` | 122 | PHP | `getAppKey` |
| `src/Client.php` | 127 | PHP | `getAppSecret` |
| `src/Client.php` | 132 | PHP | `setAccessToken` |
| `src/Client.php` | 137 | PHP | `setShopCipher` |
| `src/Client.php` | 142 | PHP | `auth` |
| `src/Client.php` | 147 | PHP | `webhook` |
| `src/Client.php` | 162 | PHP | `modifyRequestBeforeSend` |
| `src/Client.php` | 198 | PHP | `httpClient` |
| `src/Client.php` | 222 | PHP | `prepareSignature` |
| `src/Client.php` | 260 | PHP | `__get` |
| `src/Client.php` | 279 | PHP | `call` |
| `src/Client.php` | 296 | PHP | `get` |
| `src/Client.php` | 301 | PHP | `post` |
| `src/Resource.php` | 36 | PHP | `useVersion` |
| `src/Resource.php` | 53 | PHP | `useHttpClient` |
| `src/Resource.php` | 63 | PHP | `call` |
| `src/Resource.php` | 94 | PHP | `getLastMessage` |
| `src/Resource.php` | 99 | PHP | `getLastRequestId` |
| `src/Resource.php` | 108 | PHP | `handleErrorResponse` |
| `src/Resource.php` | 123 | PHP | `extractParams` |
| `src/Resource.php` | 136 | PHP | `dataTypeCast` |
| `src/Resources/AffiliateCreator.php` | 21 | PHP | `addShowcaseProducts` |
| `src/Resources/AffiliateCreator.php` | 30 | PHP | `getShowcaseProducts` |
| `src/Resources/AffiliateCreator.php` | 42 | PHP | `getCreatorProfile` |
| `src/Resources/AffiliateCreator.php` | 47 | PHP | `searchOpenCollaborationProduct` |
| `src/Resources/AffiliateCreator.php` | 59 | PHP | `searchTargetCollaborations` |
| `src/Resources/AffiliateCreator.php` | 71 | PHP | `searchAffiliateOrders` |
| `src/Resources/AffiliatePartner.php` | 21 | PHP | `createAffiliatePartnerCampaign` |
| `src/Resources/AffiliatePartner.php` | 29 | PHP | `getAffiliatePartnerCampaignProductList` |
| `src/Resources/AffiliatePartner.php` | 36 | PHP | `getAffiliatePartnerCampaignDetail` |
| `src/Resources/AffiliatePartner.php` | 43 | PHP | `reviewAffiliatePartnerCampaignProduct` |
| `src/Resources/AffiliatePartner.php` | 51 | PHP | `editAffiliatePartnerCampaign` |
| `src/Resources/AffiliatePartner.php` | 59 | PHP | `getAffiliatePartnerCampaignList` |
| `src/Resources/AffiliatePartner.php` | 66 | PHP | `publishAffiliatePartnerCampaign` |
| `src/Resources/AffiliatePartner.php` | 73 | PHP | `generateAffiliatePartnerCampaignProductLink` |
| `src/Resources/AffiliateSeller.php` | 21 | PHP | `editOpenCollaborationSettings` |
| `src/Resources/AffiliateSeller.php` | 28 | PHP | `searchOpenCollaborationProduct` |
| `src/Resources/AffiliateSeller.php` | 40 | PHP | `searchSellerAffiliateOrders` |
| `src/Resources/AffiliateSeller.php` | 52 | PHP | `createOpenCollaboration` |
| `src/Resources/AffiliateSeller.php` | 63 | PHP | `createTargetCollaboration` |
| `src/Resources/AffiliateSeller.php` | 70 | PHP | `removeCreatorAffiliateFromCollaboration` |
| `src/Resources/AffiliateSeller.php` | 80 | PHP | `getMarketplaceCreatorPerformance` |
| `src/Resources/AffiliateSeller.php` | 85 | PHP | `searchCreatorOnMarketplace` |
| `src/Resources/AffiliateSeller.php` | 97 | PHP | `generateAffiliateProductPromotionLink` |
| `src/Resources/AffiliateSeller.php` | 102 | PHP | `searchSampleApplicationsFulfillments` |
| `src/Resources/AffiliateSeller.php` | 109 | PHP | `reviewSampleApplications` |
| `src/Resources/AffiliateSeller.php` | 119 | PHP | `getOpenCollaborationSampleRules` |
| `src/Resources/AffiliateSeller.php` | 128 | PHP | `searchSampleApplications` |
| `src/Resources/AffiliateSeller.php` | 140 | PHP | `editOpenCollaborationSampleRule` |
| `src/Resources/AffiliateSeller.php` | 147 | PHP | `removeTargetCollaboration` |
| `src/Resources/AffiliateSeller.php` | 153 | PHP | `queryTargetCollaborationDetail` |
| `src/Resources/AffiliateSeller.php` | 159 | PHP | `searchTargetCollaborations` |
| `src/Resources/AffiliateSeller.php` | 171 | PHP | `updateTargetCollaboration` |
| `src/Resources/AffiliateSeller.php` | 178 | PHP | `searchOpenCollaboration` |
| `src/Resources/AffiliateSeller.php` | 190 | PHP | `getOpenCollaborationSettings` |
| `src/Resources/AffiliateSeller.php` | 195 | PHP | `removeOpenCollaboration` |
| `src/Resources/AffiliateSeller.php` | 200 | PHP | `getOpenCollaborationCreatorContentDetail` |
| `src/Resources/AffiliateSeller.php` | 210 | PHP | `getMessageInConversation` |
| `src/Resources/AffiliateSeller.php` | 221 | PHP | `getConversationList` |
| `src/Resources/AffiliateSeller.php` | 232 | PHP | `sendImMessage` |
| `src/Resources/AffiliateSeller.php` | 242 | PHP | `createConversationWithCreator` |
| `src/Resources/AffiliateSeller.php` | 252 | PHP | `markConversationRead` |
| `src/Resources/AffiliateSeller.php` | 261 | PHP | `getLatestUnreadMessages` |
| `src/Resources/Analytics.php` | 21 | PHP | `getShopPerformance` |
| `src/Resources/Analytics.php` | 28 | PHP | `getShopProductPerformance` |
| `src/Resources/Analytics.php` | 35 | PHP | `getShopProductPerformanceList` |
| `src/Resources/Analytics.php` | 42 | PHP | `getShopSkuPerformance` |
| `src/Resources/Analytics.php` | 49 | PHP | `getShopSkuPerformanceList` |
| `src/Resources/Analytics.php` | 56 | PHP | `getShopVideoPerformanceList` |
| `src/Resources/Analytics.php` | 63 | PHP | `getShopVideoPerformanceOverview` |
| `src/Resources/Analytics.php` | 70 | PHP | `getShopVideoPerformance` |
| `src/Resources/Analytics.php` | 77 | PHP | `getShopVideoProductPerformanceList` |
| `src/Resources/Authorization.php` | 19 | PHP | `getAuthorizedShop` |
| `src/Resources/Authorization.php` | 24 | PHP | `getAuthorizedCategoryAssets` |
| `src/Resources/CustomerService.php` | 20 | PHP | `getConversationMessages` |
| `src/Resources/CustomerService.php` | 31 | PHP | `getConversations` |
| `src/Resources/CustomerService.php` | 42 | PHP | `sendMessage` |
| `src/Resources/CustomerService.php` | 52 | PHP | `getAgentSettings` |
| `src/Resources/CustomerService.php` | 57 | PHP | `updateAgentSettings` |
| `src/Resources/CustomerService.php` | 64 | PHP | `uploadBuyerMessagesImage` |
| `src/Resources/CustomerService.php` | 77 | PHP | `readMessage` |
| `src/Resources/CustomerService.php` | 82 | PHP | `createConversation` |
| `src/Resources/Event.php` | 20 | PHP | `getShopWebhooks` |
| `src/Resources/Event.php` | 25 | PHP | `updateShopWebhook` |
| `src/Resources/Event.php` | 35 | PHP | `deleteShopWebhook` |
| `src/Resources/Finance.php` | 20 | PHP | `getOrderStatementTransactions` |
| `src/Resources/Finance.php` | 25 | PHP | `getStatementTransactions` |
| `src/Resources/Finance.php` | 36 | PHP | `getWithdrawals` |
| `src/Resources/Finance.php` | 47 | PHP | `getStatements` |
| `src/Resources/Finance.php` | 58 | PHP | `getPayments` |
| `src/Resources/Finance.php` | 69 | PHP | `getTransactionsByOrder` |
| `src/Resources/FulfilledByTiktok.php` | 21 | PHP | `getFbtMerchantOnboardedRegions` |
| `src/Resources/FulfilledByTiktok.php` | 26 | PHP | `getFbtWarehouseList` |
| `src/Resources/FulfilledByTiktok.php` | 31 | PHP | `getInboundOrder` |
| `src/Resources/FulfilledByTiktok.php` | 40 | PHP | `searchFbtInventory` |
| `src/Resources/FulfilledByTiktok.php` | 52 | PHP | `searchFbtInventoryRecord` |
| `src/Resources/FulfilledByTiktok.php` | 64 | PHP | `searchGoodsInfo` |
| `src/Resources/Fulfillment.php` | 21 | PHP | `searchCombinablePackages` |
| `src/Resources/Fulfillment.php` | 32 | PHP | `getPackageShippingDocument` |
| `src/Resources/Fulfillment.php` | 42 | PHP | `getPackageHandoverTimeSlots` |
| `src/Resources/Fulfillment.php` | 47 | PHP | `getTracking` |
| `src/Resources/Fulfillment.php` | 52 | PHP | `updatePackageShippingInfo` |
| `src/Resources/Fulfillment.php` | 62 | PHP | `searchPackage` |
| `src/Resources/Fulfillment.php` | 74 | PHP | `shipPackage` |
| `src/Resources/Fulfillment.php` | 85 | PHP | `getPackageDetail` |
| `src/Resources/Fulfillment.php` | 90 | PHP | `fulfillmentUploadDeliveryImage` |
| `src/Resources/Fulfillment.php` | 103 | PHP | `fulfillmentUploadDeliveryFile` |
| `src/Resources/Fulfillment.php` | 124 | PHP | `updatePackageDeliveryStatus` |
| `src/Resources/Fulfillment.php` | 133 | PHP | `batchShipPackages` |
| `src/Resources/Fulfillment.php` | 142 | PHP | `getOrderSplitAttributes` |
| `src/Resources/Fulfillment.php` | 151 | PHP | `splitOrders` |
| `src/Resources/Fulfillment.php` | 161 | PHP | `combinePackage` |
| `src/Resources/Fulfillment.php` | 170 | PHP | `uncombinePackages` |
| `src/Resources/Fulfillment.php` | 179 | PHP | `markPackageAsShipped` |
| `src/Resources/Fulfillment.php` | 190 | PHP | `getEligibleShippingService` |
| `src/Resources/Fulfillment.php` | 197 | PHP | `createPackages` |
| `src/Resources/Fulfillment.php` | 206 | PHP | `updateShippingInfo` |
| `src/Resources/GlobalProduct.php` | 20 | PHP | `createGlobalProduct` |
| `src/Resources/GlobalProduct.php` | 27 | PHP | `getGlobalProduct` |
| `src/Resources/GlobalProduct.php` | 32 | PHP | `publishGlobalProduct` |
| `src/Resources/GlobalProduct.php` | 39 | PHP | `getGlobalCategories` |
| `src/Resources/GlobalProduct.php` | 46 | PHP | `editGlobalProduct` |
| `src/Resources/GlobalProduct.php` | 53 | PHP | `getGlobalAttributes` |
| `src/Resources/GlobalProduct.php` | 60 | PHP | `getGlobalCategoryRules` |
| `src/Resources/GlobalProduct.php` | 65 | PHP | `recommendGlobalCategories` |
| `src/Resources/GlobalProduct.php` | 72 | PHP | `updateGlobalInventory` |
| `src/Resources/GlobalProduct.php` | 79 | PHP | `searchGlobalProducts` |
| `src/Resources/GlobalProduct.php` | 90 | PHP | `deleteGlobalProducts` |
| `src/Resources/Logistic.php` | 19 | PHP | `getWarehouseDeliveryOptions` |
| `src/Resources/Logistic.php` | 24 | PHP | `getShippingProvider` |
| `src/Resources/Logistic.php` | 29 | PHP | `getWarehouseList` |
| `src/Resources/Logistic.php` | 34 | PHP | `getGlobalSellerWarehouse` |
| `src/Resources/Order.php` | 30 | PHP | `getOrderDetail` |
| `src/Resources/Order.php` | 39 | PHP | `getOrderList` |
| `src/Resources/Order.php` | 51 | PHP | `addExternalOrderReferences` |
| `src/Resources/Order.php` | 60 | PHP | `searchOrderByExternalOrderReferences` |
| `src/Resources/Order.php` | 70 | PHP | `getExternalOrderReferences` |
| `src/Resources/Order.php` | 79 | PHP | `getPriceDetail` |
| `src/Resources/Product.php` | 23 | PHP | `uploadProductFile` |
| `src/Resources/Product.php` | 44 | PHP | `uploadProductImage` |
| `src/Resources/Product.php` | 61 | PHP | `createProduct` |
| `src/Resources/Product.php` | 68 | PHP | `deleteProducts` |
| `src/Resources/Product.php` | 77 | PHP | `editProduct` |
| `src/Resources/Product.php` | 84 | PHP | `updateInventory` |
| `src/Resources/Product.php` | 91 | PHP | `getProduct` |
| `src/Resources/Product.php` | 98 | PHP | `deactivateProducts` |
| `src/Resources/Product.php` | 107 | PHP | `activateProducts` |
| `src/Resources/Product.php` | 116 | PHP | `recoverProducts` |
| `src/Resources/Product.php` | 125 | PHP | `updatePrice` |
| `src/Resources/Product.php` | 132 | PHP | `getCategories` |
| `src/Resources/Product.php` | 139 | PHP | `getBrands` |
| `src/Resources/Product.php` | 150 | PHP | `createCustomBrand` |
| `src/Resources/Product.php` | 159 | PHP | `getAttributes` |
| `src/Resources/Product.php` | 166 | PHP | `getCategoryRules` |
| `src/Resources/Product.php` | 173 | PHP | `recommendCategory` |
| `src/Resources/Product.php` | 185 | PHP | `checkListingPrerequisites` |
| `src/Resources/Product.php` | 190 | PHP | `partialEditProduct` |
| `src/Resources/Product.php` | 197 | PHP | `searchProducts` |
| `src/Resources/Product.php` | 209 | PHP | `inventorySearch` |
| `src/Resources/Product.php` | 216 | PHP | `searchResponsiblePersons` |
| `src/Resources/Product.php` | 228 | PHP | `searchManufacturers` |
| `src/Resources/Product.php` | 240 | PHP | `createResponsiblePerson` |
| `src/Resources/Product.php` | 247 | PHP | `createManufacturer` |
| `src/Resources/Product.php` | 254 | PHP | `partialEditManufacturer` |
| `src/Resources/Product.php` | 261 | PHP | `partialEditResponsiblePerson` |
| `src/Resources/Promotion.php` | 20 | PHP | `createActivity` |
| `src/Resources/Promotion.php` | 33 | PHP | `updateActivityProduct` |
| `src/Resources/Promotion.php` | 43 | PHP | `removeActivityProduct` |
| `src/Resources/Promotion.php` | 53 | PHP | `searchActivities` |
| `src/Resources/Promotion.php` | 60 | PHP | `getActivity` |
| `src/Resources/Promotion.php` | 65 | PHP | `updateActivity` |
| `src/Resources/Promotion.php` | 76 | PHP | `deactivateActivity` |
| `src/Resources/Promotion.php` | 81 | PHP | `searchCoupons` |
| `src/Resources/Promotion.php` | 89 | PHP | `getCoupon` |
| `src/Resources/ReturnRefund.php` | 20 | PHP | `searchCancellations` |
| `src/Resources/ReturnRefund.php` | 31 | PHP | `approveCancellation` |
| `src/Resources/ReturnRefund.php` | 40 | PHP | `rejectCancellation` |
| `src/Resources/ReturnRefund.php` | 50 | PHP | `searchReturns` |
| `src/Resources/ReturnRefund.php` | 61 | PHP | `approveReturn` |
| `src/Resources/ReturnRefund.php` | 71 | PHP | `rejectReturn` |
| `src/Resources/ReturnRefund.php` | 81 | PHP | `getAftersaleEligibility` |
| `src/Resources/ReturnRefund.php` | 88 | PHP | `getRejectReasons` |
| `src/Resources/ReturnRefund.php` | 95 | PHP | `calculateRefund` |
| `src/Resources/ReturnRefund.php` | 102 | PHP | `getReturnRecords` |
| `src/Resources/ReturnRefund.php` | 107 | PHP | `cancelOrder` |
| `src/Resources/ReturnRefund.php` | 114 | PHP | `createReturn` |
| `src/Resources/ReturnRefund.php` | 129 | PHP | `generateIdempotencyKey` |
| `src/Resources/Seller.php` | 19 | PHP | `getActiveShopList` |
| `src/Resources/Seller.php` | 24 | PHP | `getSellerPermissions` |
| `src/Resources/Supplychain.php` | 20 | PHP | `confirmPackageShipment` |
| `src/Webhook.php` | 36 | PHP | `__construct` |
| `src/Webhook.php` | 41 | PHP | `getType` |
| `src/Webhook.php` | 46 | PHP | `getShopId` |
| `src/Webhook.php` | 51 | PHP | `getData` |
| `src/Webhook.php` | 56 | PHP | `getTimestamp` |
| `src/Webhook.php` | 61 | PHP | `capture` |
| `src/Webhook.php` | 78 | PHP | `verify` |
| `tests/AuthTest.php` | 30 | PHP | `setUp` |
| `tests/AuthTest.php` | 35 | PHP | `testRefreshNewToken` |
| `tests/AuthTest.php` | 62 | PHP | `testGetToken` |
| `tests/AuthTest.php` | 89 | PHP | `testCreateAuthRequest` |
| `tests/AuthTest.php` | 103 | PHP | `test__construct` |
| `tests/ClientTest.php` | 31 | PHP | `setUp` |
| `tests/ClientTest.php` | 38 | PHP | `testSetAccessToken` |
| `tests/ClientTest.php` | 49 | PHP | `testAuth` |
| `tests/ClientTest.php` | 54 | PHP | `test__get` |
| `tests/ClientTest.php` | 75 | PHP | `test__construct` |
| `tests/ClientTest.php` | 81 | PHP | `testPrepareSignature` |
| `tests/ClientTest.php` | 85 | PHP | `prepareSignature` |
| `tests/ClientTest.php` | 98 | PHP | `testModifyRequestBeforeSend` |
| `tests/ResourceTest.php` | 27 | PHP | `setUp` |
| `tests/ResourceTest.php` | 34 | PHP | `testCall` |
| `tests/ResourceTest.php` | 48 | PHP | `testLastMessageAndRequestId` |
| `tests/ResourceTest.php` | 63 | PHP | `testChangeAPIVersion` |
| `tests/Resources/AffiliateCreatorTest.php` | 22 | PHP | `testAddShowcaseProducts` |
| `tests/Resources/AffiliateCreatorTest.php` | 28 | PHP | `testGetShowcaseProducts` |
| `tests/Resources/AffiliateCreatorTest.php` | 34 | PHP | `testGetCreatorProfile` |
| `tests/Resources/AffiliateCreatorTest.php` | 40 | PHP | `testSearchOpenCollaborationProduct` |
| `tests/Resources/AffiliateCreatorTest.php` | 46 | PHP | `testSearchTargetCollaborations` |
| `tests/Resources/AffiliateCreatorTest.php` | 52 | PHP | `testSearchAffiliateOrders` |
| `tests/Resources/AffiliatePartnerTest.php` | 14 | PHP | `testCreateAffiliatePartnerCampaign` |
| `tests/Resources/AffiliatePartnerTest.php` | 20 | PHP | `testGetAffiliatePartnerCampaignProductList` |
| `tests/Resources/AffiliatePartnerTest.php` | 26 | PHP | `testGetAffiliatePartnerCampaignDetail` |
| `tests/Resources/AffiliatePartnerTest.php` | 32 | PHP | `testReviewAffiliatePartnerCampaignProduct` |
| `tests/Resources/AffiliatePartnerTest.php` | 38 | PHP | `testEditAffiliatePartnerCampaign` |
| `tests/Resources/AffiliatePartnerTest.php` | 44 | PHP | `testGetAffiliatePartnerCampaignList` |
| `tests/Resources/AffiliatePartnerTest.php` | 50 | PHP | `testPublishAffiliatePartnerCampaign` |
| `tests/Resources/AffiliatePartnerTest.php` | 56 | PHP | `testGenerateAffiliatePartnerCampaignProductLink` |
| `tests/Resources/AffiliateSellerTest.php` | 23 | PHP | `testEditOpenCollaborationSettings` |
| `tests/Resources/AffiliateSellerTest.php` | 29 | PHP | `testSearchOpenCollaborationProduct` |
| `tests/Resources/AffiliateSellerTest.php` | 35 | PHP | `testSearchSellerAffiliateOrders` |
| `tests/Resources/AffiliateSellerTest.php` | 41 | PHP | `testCreateOpenCollaboration` |
| `tests/Resources/AffiliateSellerTest.php` | 47 | PHP | `testCreateTargetCollaboration` |
| `tests/Resources/AffiliateSellerTest.php` | 53 | PHP | `testRemoveCreatorAffiliateFromCollaboration` |
| `tests/Resources/AffiliateSellerTest.php` | 59 | PHP | `testGetMarketplaceCreatorPerformance` |
| `tests/Resources/AffiliateSellerTest.php` | 65 | PHP | `testSearchCreatorOnMarketplace` |
| `tests/Resources/AffiliateSellerTest.php` | 71 | PHP | `testGenerateAffiliateProductPromotionLink` |
| `tests/Resources/AffiliateSellerTest.php` | 77 | PHP | `testSearchSampleApplicationsFulfillments` |
| `tests/Resources/AffiliateSellerTest.php` | 83 | PHP | `testReviewSampleApplications` |
| `tests/Resources/AffiliateSellerTest.php` | 89 | PHP | `testGetOpenCollaborationSampleRules` |
| `tests/Resources/AffiliateSellerTest.php` | 95 | PHP | `testSearchSampleApplications` |
| `tests/Resources/AffiliateSellerTest.php` | 101 | PHP | `testEditOpenCollaborationSampleRule` |
| `tests/Resources/AffiliateSellerTest.php` | 107 | PHP | `testRemoveTargetCollaboration` |
| `tests/Resources/AffiliateSellerTest.php` | 113 | PHP | `testQueryTargetCollaborationDetail` |
| `tests/Resources/AffiliateSellerTest.php` | 119 | PHP | `testSearchTargetCollaborations` |
| `tests/Resources/AffiliateSellerTest.php` | 125 | PHP | `testUpdateTargetCollaboration` |
| `tests/Resources/AffiliateSellerTest.php` | 131 | PHP | `testSearchOpenCollaboration` |
| `tests/Resources/AffiliateSellerTest.php` | 137 | PHP | `testGetOpenCollaborationSettings` |
| `tests/Resources/AffiliateSellerTest.php` | 143 | PHP | `testRemoveOpenCollaboration` |
| `tests/Resources/AffiliateSellerTest.php` | 149 | PHP | `testGetOpenCollaborationCreatorContentDetail` |
| `tests/Resources/AffiliateSellerTest.php` | 155 | PHP | `testGetMessageInConversation` |
| `tests/Resources/AffiliateSellerTest.php` | 161 | PHP | `testGetConversationList` |
| `tests/Resources/AffiliateSellerTest.php` | 167 | PHP | `testSendImMessage` |
| `tests/Resources/AffiliateSellerTest.php` | 173 | PHP | `testCreateConversationWithCreator` |
| `tests/Resources/AffiliateSellerTest.php` | 179 | PHP | `testMarkConversationRead` |
| `tests/Resources/AffiliateSellerTest.php` | 185 | PHP | `testGetLatestUnreadMessages` |
| `tests/Resources/AnalyticsTest.php` | 23 | PHP | `testGetShopPerformance` |
| `tests/Resources/AnalyticsTest.php` | 29 | PHP | `testGetShopProductPerformance` |
| `tests/Resources/AnalyticsTest.php` | 35 | PHP | `testGetShopProductPerformanceList` |
| `tests/Resources/AnalyticsTest.php` | 41 | PHP | `testGetShopSkuPerformance` |
| `tests/Resources/AnalyticsTest.php` | 47 | PHP | `testGetShopSkuPerformanceList` |
| `tests/Resources/AnalyticsTest.php` | 53 | PHP | `testGetShopVideoPerformance` |
| `tests/Resources/AnalyticsTest.php` | 60 | PHP | `testGetShopVideoPerformanceList` |
| `tests/Resources/AnalyticsTest.php` | 66 | PHP | `testGetShopVideoPerformanceOverview` |
| `tests/Resources/AnalyticsTest.php` | 72 | PHP | `testGetShopVideoProductPerformanceList` |
| `tests/Resources/AuthorizationTest.php` | 21 | PHP | `testGetAuthorizedShop` |
| `tests/Resources/AuthorizationTest.php` | 27 | PHP | `testGetAuthorizedCategoryAssets` |
| `tests/Resources/CustomerServiceTest.php` | 18 | PHP | `testGetConversationMessages` |
| `tests/Resources/CustomerServiceTest.php` | 24 | PHP | `testGetConversations` |
| `tests/Resources/CustomerServiceTest.php` | 30 | PHP | `testSendMessage` |
| `tests/Resources/CustomerServiceTest.php` | 36 | PHP | `testGetAgentSettings` |
| `tests/Resources/CustomerServiceTest.php` | 42 | PHP | `testUpdateAgentSettings` |
| `tests/Resources/CustomerServiceTest.php` | 48 | PHP | `testUploadBuyerMessagesImage` |
| `tests/Resources/CustomerServiceTest.php` | 54 | PHP | `testReadMessage` |
| `tests/Resources/CustomerServiceTest.php` | 60 | PHP | `testCreateConversation` |
| `tests/Resources/EventTest.php` | 23 | PHP | `testUpdateShopWebhook` |
| `tests/Resources/EventTest.php` | 29 | PHP | `testGetShopWebhooks` |
| `tests/Resources/EventTest.php` | 35 | PHP | `testDeleteShopWebhook` |
| `tests/Resources/FinanceTest.php` | 21 | PHP | `testGetPayments` |
| `tests/Resources/FinanceTest.php` | 27 | PHP | `testGetOrderStatementTransactions` |
| `tests/Resources/FinanceTest.php` | 34 | PHP | `testGetStatementTransactions` |
| `tests/Resources/FinanceTest.php` | 41 | PHP | `testGetStatements` |
| `tests/Resources/FinanceTest.php` | 47 | PHP | `testGetWithdrawals` |
| `tests/Resources/FulfilledByTiktokTest.php` | 11 | PHP | `testGetFbtMerchantOnboardedRegions` |
| `tests/Resources/FulfilledByTiktokTest.php` | 17 | PHP | `getFbtWarehouseList` |
| `tests/Resources/FulfilledByTiktokTest.php` | 23 | PHP | `getInboundOrder` |
| `tests/Resources/FulfilledByTiktokTest.php` | 29 | PHP | `testSearchFbtInventory` |
| `tests/Resources/FulfilledByTiktokTest.php` | 35 | PHP | `testSearchFbtInventoryRecord` |
| `tests/Resources/FulfilledByTiktokTest.php` | 41 | PHP | `testSearchGoodsInfo` |
| `tests/Resources/FulfillmentTest.php` | 21 | PHP | `testGetPackageShippingDocument` |
| `tests/Resources/FulfillmentTest.php` | 28 | PHP | `testGetPackageHandoverTimeSlots` |
| `tests/Resources/FulfillmentTest.php` | 35 | PHP | `testUpdatePackageDeliveryStatus` |
| `tests/Resources/FulfillmentTest.php` | 41 | PHP | `testSearchCombinablePackages` |

_Display capped at 300 of 389 detected functions/methods._

### Automation pipeline files

- `.github/workflows/ci.yml`

## `cvsz/zLinebot-automos`

Status: **cloned**; commit `ba8382907971`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `api-gateway/server.ts` | 9 | `/health` |
| `api/main.py` | 33 | `/health` |
| `api/main.py` | 38 | `/api/register` |
| `api/main.py` | 47 | `/api/chat` |
| `api/main.py` | 61 | `/api/checkout` |
| `api/main.py` | 82 | `/api/create-checkout` |
| `api/main.py` | 84 | `/api/checkout` |
| `autonomos/api/routes.js` | 57 | `/auth/register` |
| `autonomos/api/routes.js` | 58 | `/auth/login` |
| `autonomos/api/routes.js` | 60 | `/portfolio` |
| `autonomos/api/routes.js` | 70 | `/backtest` |
| `autonomos/api/routes.js` | 83 | `/optimize` |
| `autonomos/api/routes.js` | 93 | `/pipeline/tune` |
| `autonomos/api/routes.js` | 99 | `/pipeline/run` |
| `autonomos/api/routes.js` | 120 | `/trading/queue` |
| `autonomos/api/routes.js` | 130 | `/risk` |
| `autonomos/api/routes.js` | 132 | `/market/cache/:symbol` |
| `autonomos/api/routes.js` | 142 | `/market/cache/:symbol` |
| `autonomos/api/routes.js` | 153 | `/kyc` |
| `autonomos/api/routes.js` | 163 | `/copy/follow` |
| `autonomos/api/routes.js` | 174 | `/copy/unfollow` |
| `autonomos/api/routes.js` | 185 | `/copy/leaderboard` |
| `autonomos/api/routes.js` | 195 | `/billing/checkout` |
| `autonomos/api/routes.js` | 204 | `/billing/webhook` |
| `autonomos/api/routes.js` | 206 | `/admin/users` |
| `autonomos/api/routes.js` | 215 | `/admin/subscriptions` |
| `autonomos/api/routes.js` | 224 | `/admin/trades` |
| `autonomos/api/routes.js` | 233 | `/admin/logs` |
| `autonomos/api/server.js` | 11 | `/api/billing/webhook` |
| `autonomos/api/server.js` | 14 | `/webhook` |
| `autonomos/api/server.js` | 25 | `/api` |
| `backend-node/server.js` | 17 | `/health` |
| `backend-node/server.js` | 21 | `/api/create-checkout` |
| `backend-node/server.js` | 42 | `/api/chat` |
| `backend-node/webhook.js` | 11 | `/health` |
| `backend-node/webhook.js` | 15 | `/webhook` |
| `backend/api/main.py` | 626 | `/health` |
| `backend/api/main.py` | 630 | `/api/health` |
| `backend/api/main.py` | 634 | `/api/register` |
| `backend/api/main.py` | 686 | `/api/login` |
| `backend/api/main.py` | 721 | `/api/me` |
| `backend/api/main.py` | 731 | `/api/branding` |
| `backend/api/main.py` | 746 | `/api/branding` |
| `backend/api/main.py` | 777 | `/api/team` |
| `backend/api/main.py` | 788 | `/api/team` |
| `backend/api/main.py` | 823 | `/api/chat` |
| `backend/api/main.py` | 830 | `/webhook/{tenant_id}` |
| `backend/api/main.py` | 921 | `/api/leads` |
| `backend/api/main.py` | 935 | `/api/export/tiktok.csv` |
| `backend/api/main.py` | 966 | `/api/leads/{lead_id}` |
| `backend/api/main.py` | 988 | `/api/stats` |
| `backend/api/main.py` | 1012 | `/api/revenue/daily` |
| `backend/api/main.py` | 1029 | `/api/templates` |
| `backend/api/main.py` | 1049 | `/api/templates` |
| `backend/api/main.py` | 1059 | `/api/broadcast` |
| `backend/api/main.py` | 1111 | `/api/campaigns` |
| `backend/api/main.py` | 1121 | `/api/billing/checkout` |
| `backend/api/main.py` | 1139 | `/stripe/webhook` |
| `backend/api/tests/test_auth.py` | 24 | `/api/register` |
| `backend/api/tests/test_auth.py` | 30 | `/api/login` |
| `backend/api/tests/test_auth.py` | 48 | `/api/register` |
| `backend/api/tests/test_auth.py` | 64 | `/api/register` |
| `backend/api/tests/test_auth.py` | 65 | `/api/register` |
| `backend/api/tests/test_auth.py` | 75 | `/api/register` |
| `backend/api/tests/test_auth.py` | 77 | `/api/login` |
| `backend/api/tests/test_auth.py` | 79 | `/api/login` |
| `backend/api/tests/test_auth.py` | 83 | `/api/login` |
| `backend/api/tests/test_auth.py` | 100 | `/api/register` |
| `backend/api/tests/test_auth.py` | 112 | `/api/register` |
| `backend/api/tests/test_auth.py` | 126 | `/webhook/{session[` |
| `backend/api/tests/test_auth.py` | 139 | `/api/leads/1` |
| `backend/api/tests/test_auth.py` | 140 | `/api/leads?status=invalid` |
| `backend/api/tests/test_auth.py` | 142 | `/api/broadcast` |
| `backend/api/tests/test_auth.py` | 159 | `/api/templates` |
| `backend/api/tests/test_auth.py` | 164 | `/api/broadcast` |
| `backend/api/tests/test_auth.py` | 181 | `/api/team` |
| `backend/api/tests/test_auth.py` | 187 | `/api/login` |
| `backend/api/tests/test_auth.py` | 191 | `/api/stats` |
| `backend/api/tests/test_auth.py` | 192 | `/api/export/tiktok.csv` |
| `backend/api/tests/test_auth.py` | 193 | `/api/leads` |
| `backend/api/tests/test_auth.py` | 256 | `/webhook/tenant-1` |
| `backend/api/tests/test_auth.py` | 276 | `/webhook/tenant-2` |
| `backend/api/tests/test_auth.py` | 291 | `/webhook/{tenant_id}` |
| `backend/api/tests/test_auth.py` | 314 | `/webhook/{tenant_id}` |
| `backend/api/tests/test_auth.py` | 358 | `/api/branding` |
| `backend/api/tests/test_auth.py` | 362 | `/api/branding` |
| `backend/api/tests/test_auth.py` | 376 | `/api/team` |
| `backend/api/tests/test_auth.py` | 381 | `/api/login` |
| `backend/api/tests/test_auth.py` | 386 | `/api/team` |
| `backend/api/tests/test_auth.py` | 390 | `/api/team` |
| `backend/api/tests/test_auth.py` | 407 | `/webhook/{session[` |
| `backend/api/tests/test_auth.py` | 420 | `/api/export/tiktok.csv` |
| `backend/api/tests/test_auth.py` | 458 | `/webhook/{tenant_id}` |
| `frontend/src/lib/api.js` | 87 | `/api/register` |
| `frontend/src/lib/api.js` | 91 | `/api/login` |
| `frontend/src/lib/api.js` | 95 | `/api/chat` |
| `frontend/src/lib/api.js` | 99 | `/api/me` |
| `frontend/src/lib/api.js` | 103 | `/api/stats` |
| `frontend/src/lib/api.js` | 116 | `/api/revenue/daily` |
| `frontend/src/lib/api.js` | 120 | `/api/templates` |
| `frontend/src/lib/api.js` | 124 | `/api/templates` |
| `frontend/src/lib/api.js` | 128 | `/api/campaigns` |
| `frontend/src/lib/api.js` | 132 | `/api/broadcast` |
| `frontend/src/lib/api.js` | 136 | `/api/billing/checkout` |
| `landing/src/App.jsx` | 9 | `/api/create-checkout` |
| `landing/src/components/DemoChat.jsx` | 14 | `/api/chat` |
| `scripts/zlinebot_autonomos.sh` | 180 | `/webhook` |
| `tools/control_panel/app.py` | 80 | `/` |
| `tools/control_panel/app.py` | 90 | `/run` |
| `tools/control_panel/app.py` | 97 | `/logs/<log_name>` |
| `tools/control_panel/app.py` | 110 | `/logs/<log_name>/stream` |
| `zBAutomos/api/src/market/routes.ts` | 4 | `/snapshot` |
| `zBAutomos/api/src/server.ts` | 28 | `/health` |
| `zBAutomos/api/src/server.ts` | 30 | `/auth` |
| `zBAutomos/api/src/ws/routes.ts` | 4 | `/status` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `ai-agent/agent.py` | 6 | Python | `run_agent` |
| `ai/optimizer.ts` | 13 | JS/TS | `scoreResult` |
| `ai/optimizer.ts` | 18 | JS/TS | `optimize` |
| `analytics/exposure.js` | 1 | JS/TS | `exposure` |
| `analytics/exposure.js` | 9 | JS/TS | `netExposure` |
| `analytics/pnl.js` | 1 | JS/TS | `pnl` |
| `analytics/pnl.js` | 9 | JS/TS | `pnlByStrategy` |
| `analytics/var.js` | 1 | JS/TS | `VaR` |
| `analytics/var.js` | 15 | JS/TS | `cVaR` |
| `api/main.py` | 34 | Python | `health` |
| `api/main.py` | 39 | Python | `register` |
| `api/main.py` | 48 | Python | `chat` |
| `api/main.py` | 62 | Python | `checkout` |
| `api/main.py` | 83 | Python | `legacy_checkout_redirect` |
| `arb/latency_engine.js` | 37 | JS/TS | `normalize` |
| `arb/latency_engine.js` | 59 | JS/TS | `updateQuote` |
| `arb/latency_engine.js` | 67 | JS/TS | `stale` |
| `arb/latency_engine.js` | 71 | JS/TS | `pairKey` |
| `arb/latency_engine.js` | 75 | JS/TS | `estimateRoundTripFees` |
| `arb/latency_engine.js` | 79 | JS/TS | `canTrade` |
| `arb/latency_engine.js` | 85 | JS/TS | `placeBuy` |
| `arb/latency_engine.js` | 91 | JS/TS | `placeSell` |
| `arb/latency_engine.js` | 97 | JS/TS | `executeArb` |
| `arb/latency_engine.js` | 121 | JS/TS | `checkPair` |
| `arb/latency_engine.js` | 143 | JS/TS | `checkArb` |
| `arb/latency_engine.js` | 154 | JS/TS | `connect` |
| `autonomos/admin/api.js` | 3 | JS/TS | `getAdminUsers` |
| `autonomos/admin/api.js` | 8 | JS/TS | `getAdminTrades` |
| `autonomos/admin/api.js` | 13 | JS/TS | `getAdminSubscriptions` |
| `autonomos/admin/api.js` | 20 | JS/TS | `getAdminLogs` |
| `autonomos/agents/brain.js` | 3 | JS/TS | `brain` |
| `autonomos/agents/orchestrator.js` | 1 | JS/TS | `runAgents` |
| `autonomos/ai/chatgpt.js` | 5 | JS/TS | `askAI` |
| `autonomos/ai/tuner.js` | 3 | JS/TS | `optimize` |
| `autonomos/ai/tuner.js` | 10 | JS/TS | `candidateStrategy` |
| `autonomos/analytics/exposure.js` | 1 | JS/TS | `exposure` |
| `autonomos/analytics/metrics.js` | 1 | JS/TS | `sharpe` |
| `autonomos/analytics/metrics.js` | 12 | JS/TS | `drawdown` |
| `autonomos/analytics/metrics.js` | 27 | JS/TS | `winRate` |
| `autonomos/analytics/pnl.js` | 1 | JS/TS | `pnl` |
| `autonomos/analytics/var.js` | 1 | JS/TS | `VaR` |
| `autonomos/api/routes.js` | 20 | JS/TS | `requireAdmin` |
| `autonomos/api/routes.js` | 25 | JS/TS | `generateFakePrices` |
| `autonomos/api/routes.js` | 38 | JS/TS | `generateMarketData` |
| `autonomos/auth/jwt.js` | 3 | JS/TS | `generateToken` |
| `autonomos/auth/login.js` | 5 | JS/TS | `login` |
| `autonomos/auth/register.js` | 4 | JS/TS | `register` |
| `autonomos/autonomous.js` | 3 | JS/TS | `loop` |
| `autonomos/backtest/engine.js` | 3 | JS/TS | `safeNumber` |
| `autonomos/backtest/engine.js` | 9 | JS/TS | `constructor` |
| `autonomos/backtest/engine.js` | 27 | JS/TS | `run` |
| `autonomos/backtest/engine.js` | 72 | JS/TS | `shouldStopLoss` |
| `autonomos/backtest/engine.js` | 76 | JS/TS | `shouldTakeProfit` |
| `autonomos/backtest/engine.js` | 80 | JS/TS | `openPosition` |
| `autonomos/backtest/engine.js` | 92 | JS/TS | `closePosition` |
| `autonomos/backtest/engine.js` | 116 | JS/TS | `currentEquity` |
| `autonomos/backtest/engine.js` | 120 | JS/TS | `getState` |
| `autonomos/backtest/engine.js` | 132 | JS/TS | `results` |
| `autonomos/backtest/engine.js` | 151 | JS/TS | `sharpe` |
| `autonomos/backtest/engine.js` | 160 | JS/TS | `maxDrawdown` |
| `autonomos/backtest/engine.js` | 175 | JS/TS | `calculateMetrics` |
| `autonomos/backtest/engine.js` | 191 | JS/TS | `calculateTradeStats` |
| `autonomos/backtest/engine.test.js` | 8 | JS/TS | `strategy` |
| `autonomos/backtest/engine.test.js` | 27 | JS/TS | `strategy` |
| `autonomos/billing/stripe.js` | 11 | JS/TS | `createCheckout` |
| `autonomos/billing/stripe.js` | 32 | JS/TS | `webhook` |
| `autonomos/cache/marketCache.js` | 6 | JS/TS | `setCachedPrice` |
| `autonomos/cache/marketCache.js` | 11 | JS/TS | `getCachedPrice` |
| `autonomos/copy/follow.js` | 3 | JS/TS | `follow` |
| `autonomos/copy/follow.js` | 18 | JS/TS | `unfollow` |
| `autonomos/copy/leaderboard.js` | 3 | JS/TS | `getLeaderboard` |
| `autonomos/copy/propagate.js` | 4 | JS/TS | `propagate` |
| `autonomos/core/trading_loop.js` | 8 | JS/TS | `buildParameterizedStrategy` |
| `autonomos/core/trading_loop.js` | 19 | JS/TS | `tradingLoop` |
| `autonomos/data/tick_collector.js` | 4 | JS/TS | `startTickCollector` |
| `autonomos/data/tick_storage.js` | 4 | JS/TS | `startTickStorage` |
| `autonomos/db.js` | 7 | JS/TS | `getPool` |
| `autonomos/db.js` | 21 | JS/TS | `query` |
| `autonomos/db.js` | 26 | JS/TS | `closeDb` |
| `autonomos/dl/orderbook_model.py` | 11 | Python | `__init__` |
| `autonomos/dl/orderbook_model.py` | 20 | Python | `forward` |
| `autonomos/execution/engine.js` | 6 | JS/TS | `smartExecute` |
| `autonomos/execution/engine.js` | 40 | JS/TS | `twapOrder` |
| `autonomos/execution/engine.js` | 63 | JS/TS | `bestExecution` |
| `autonomos/execution/engine.js` | 80 | JS/TS | `sleep` |
| `autonomos/execution/live.js` | 5 | JS/TS | `getClient` |
| `autonomos/execution/live.js` | 16 | JS/TS | `executeTrade` |
| `autonomos/kyc/service.js` | 3 | JS/TS | `submitKYC` |
| `autonomos/kyc/service.js` | 19 | JS/TS | `logAction` |
| `autonomos/memory/memory.js` | 5 | JS/TS | `saveMessage` |
| `autonomos/memory/memory.js` | 10 | JS/TS | `getRecentMessages` |
| `autonomos/middleware/auth.js` | 3 | JS/TS | `auth` |
| `autonomos/portfolio/manager.js` | 1 | JS/TS | `combine` |
| `autonomos/queue/tradingQueue.js` | 11 | JS/TS | `enqueueUserRun` |
| `autonomos/reporting/report.js` | 4 | JS/TS | `generateReport` |
| `autonomos/risk/manager.js` | 3 | JS/TS | `safeNumber` |
| `autonomos/risk/manager.js` | 8 | JS/TS | `getMaxDailyLoss` |
| `autonomos/risk/manager.js` | 12 | JS/TS | `checkRisk` |
| `autonomos/risk/manager.js` | 32 | JS/TS | `getRiskState` |
| `autonomos/risk/manager.js` | 36 | JS/TS | `resetRisk` |
| `autonomos/saas/auth.js` | 3 | JS/TS | `register` |
| `autonomos/saas/auth.js` | 16 | JS/TS | `login` |
| `autonomos/saas/auth.js` | 20 | JS/TS | `listUsers` |
| `autonomos/saas/billing.js` | 1 | JS/TS | `checkPlan` |
| `autonomos/strategy/advanced.js` | 1 | JS/TS | `strategy` |
| `autonomos/trading/arbitrage.js` | 1 | JS/TS | `findArb` |
| `autonomos/trading/backtest.js` | 7 | JS/TS | `priceSeriesToCandles` |
| `autonomos/trading/backtest.js` | 15 | JS/TS | `backtest` |
| `autonomos/trading/backtest.js` | 28 | JS/TS | `wrappedStrategy` |
| `autonomos/trading/backtest.js` | 54 | JS/TS | `metrics` |
| `autonomos/trading/binance.js` | 8 | JS/TS | `marketBuy` |
| `autonomos/trading/binance.js` | 12 | JS/TS | `marketSell` |
| `autonomos/trading/binance_safe.js` | 5 | JS/TS | `createClient` |
| `autonomos/trading/binance_safe.js` | 13 | JS/TS | `validateTrade` |
| `autonomos/trading/binance_safe.js` | 23 | JS/TS | `safeBuy` |
| `autonomos/trading/binance_safe.js` | 34 | JS/TS | `safeSell` |
| `autonomos/trading/copyTrading.js` | 3 | JS/TS | `follow` |
| `autonomos/trading/copyTrading.js` | 8 | JS/TS | `unfollow` |
| `autonomos/trading/copyTrading.js` | 12 | JS/TS | `getFollowers` |
| `autonomos/trading/copyTrading.js` | 18 | JS/TS | `propagateTrade` |
| `autonomos/trading/copyTrading.js` | 29 | JS/TS | `rankBots` |
| `autonomos/trading/engine.js` | 11 | JS/TS | `resolveAction` |
| `autonomos/trading/engine.js` | 20 | JS/TS | `runTradingEngine` |
| `autonomos/trading/engine.js` | 58 | JS/TS | `run` |
| `autonomos/trading/latency_engine.js` | 10 | JS/TS | `constructor` |
| `autonomos/trading/latency_engine.js` | 16 | JS/TS | `connect` |
| `autonomos/trading/latency_engine.js` | 49 | JS/TS | `normalize` |
| `autonomos/trading/latency_engine.js` | 70 | JS/TS | `checkArb` |
| `autonomos/trading/latency_engine.js` | 101 | JS/TS | `close` |
| `autonomos/trading/latency_engine.js` | 106 | JS/TS | `createDefaultEngine` |
| `autonomos/trading/market_maker.js` | 1 | JS/TS | `makeMarket` |
| `autonomos/trading/multiStrategy.js` | 10 | JS/TS | `combinedStrategy` |
| `autonomos/trading/optimizer.js` | 3 | JS/TS | `optimize` |
| `autonomos/trading/portfolio_optimizer.js` | 1 | JS/TS | `mean` |
| `autonomos/trading/portfolio_optimizer.js` | 6 | JS/TS | `variance` |
| `autonomos/trading/portfolio_optimizer.js` | 12 | JS/TS | `optimizePortfolio` |
| `autonomos/trading/profitGuard.js` | 3 | JS/TS | `guard` |
| `autonomos/trading/profitGuard.js` | 14 | JS/TS | `resetDailyLoss` |
| `autonomos/trading/risk.js` | 1 | JS/TS | `riskControl` |
| `autonomos/trading/rl_agent.js` | 4 | JS/TS | `constructor` |
| `autonomos/trading/rl_agent.js` | 11 | JS/TS | `getState` |
| `autonomos/trading/rl_agent.js` | 19 | JS/TS | `chooseAction` |
| `autonomos/trading/rl_agent.js` | 34 | JS/TS | `update` |
| `autonomos/trading/strategy.js` | 3 | JS/TS | `strategy` |
| `autonomos/trading/strategy_alt.js` | 3 | JS/TS | `strategy` |
| `autonomos/trading/userEngine.js` | 7 | JS/TS | `strategyFromName` |
| `autonomos/trading/userEngine.js` | 22 | JS/TS | `getUser` |
| `autonomos/trading/userEngine.js` | 27 | JS/TS | `runUser` |
| `autonomos/workers/cluster.js` | 5 | JS/TS | `startWorker` |
| `autonomos/ws/server.js` | 5 | JS/TS | `startWebSocketServer` |
| `autonomos/ws/server.js` | 16 | JS/TS | `broadcast` |
| `backend-node/config.js` | 3 | JS/TS | `toInt` |
| `backend-node/config.js` | 8 | JS/TS | `splitCsv` |
| `backend/api/main.py` | 247 | Python | `__init__` |
| `backend/api/main.py` | 250 | Python | `allow` |
| `backend/api/main.py` | 263 | Python | `make_engine` |
| `backend/api/main.py` | 270 | Python | `verify_line_signature` |
| `backend/api/main.py` | 283 | Python | `resolve_cors_origins` |
| `backend/api/main.py` | 288 | Python | `hash_password` |
| `backend/api/main.py` | 298 | Python | `verify_password` |
| `backend/api/main.py` | 317 | Python | `slugify_tenant` |
| `backend/api/main.py` | 323 | Python | `extract_phone` |
| `backend/api/main.py` | 329 | Python | `extract_name` |
| `backend/api/main.py` | 336 | Python | `clean_text` |
| `backend/api/main.py` | 340 | Python | `normalize_role` |
| `backend/api/main.py` | 347 | Python | `normalize_status` |
| `backend/api/main.py` | 359 | Python | `calculate_score` |
| `backend/api/main.py` | 374 | Python | `determine_status` |
| `backend/api/main.py` | 388 | Python | `generate_sales_reply` |
| `backend/api/main.py` | 426 | Python | `build_json_headers` |
| `backend/api/main.py` | 434 | Python | `send_line_reply` |
| `backend/api/main.py` | 446 | Python | `send_line_push` |
| `backend/api/main.py` | 458 | Python | `send_sheet_sync_stub` |
| `backend/api/main.py` | 473 | Python | `serialize_lead` |
| `backend/api/main.py` | 478 | Python | `serialize_campaign` |
| `backend/api/main.py` | 485 | Python | `serialize_template` |
| `backend/api/main.py` | 489 | Python | `serialize_team_member` |
| `backend/api/main.py` | 493 | Python | `queue_broadcast_event` |
| `backend/api/main.py` | 516 | Python | `check_rate_limit` |
| `backend/api/main.py` | 532 | Python | `get_current_user` |
| `backend/api/main.py` | 555 | Python | `require_role` |
| `backend/api/main.py` | 556 | Python | `dependency` |
| `backend/api/main.py` | 565 | Python | `get_tenant_scope` |
| `backend/api/main.py` | 584 | Python | `_health_payload` |
| `backend/api/main.py` | 589 | Python | `create_app` |
| `backend/api/main.py` | 594 | Python | `lifespan` |
| `backend/api/main.py` | 614 | Python | `db_session` |
| `backend/api/main.py` | 622 | Python | `attach_session_factory` |
| `backend/api/main.py` | 627 | Python | `health` |
| `backend/api/main.py` | 631 | Python | `health_api` |
| `backend/api/main.py` | 635 | Python | `register` |
| `backend/api/main.py` | 687 | Python | `login` |
| `backend/api/main.py` | 722 | Python | `me` |
| `backend/api/main.py` | 732 | Python | `get_branding` |
| `backend/api/main.py` | 747 | Python | `update_branding` |
| `backend/api/main.py` | 778 | Python | `list_team_members` |
| `backend/api/main.py` | 789 | Python | `add_team_member` |
| `backend/api/main.py` | 824 | Python | `chat` |
| `backend/api/main.py` | 831 | Python | `webhook` |
| `backend/api/main.py` | 922 | Python | `get_leads` |
| `backend/api/main.py` | 936 | Python | `export_tiktok_csv` |
| `backend/api/main.py` | 967 | Python | `update_lead` |
| `backend/api/main.py` | 989 | Python | `get_stats` |
| `backend/api/main.py` | 1013 | Python | `get_revenue_daily` |
| `backend/api/main.py` | 1030 | Python | `create_template` |
| `backend/api/main.py` | 1050 | Python | `get_templates` |
| `backend/api/main.py` | 1060 | Python | `create_broadcast` |
| `backend/api/main.py` | 1112 | Python | `get_campaigns` |
| `backend/api/main.py` | 1122 | Python | `create_checkout_session` |
| `backend/api/main.py` | 1140 | Python | `stripe_webhook` |
| `backend/api/tests/test_auth.py` | 15 | Python | `make_client` |
| `backend/api/tests/test_auth.py` | 21 | Python | `register_and_login` |
| `backend/api/tests/test_auth.py` | 37 | Python | `auth_headers` |
| `backend/api/tests/test_auth.py` | 44 | Python | `test_register_persists_user_and_returns_profile` |
| `backend/api/tests/test_auth.py` | 60 | Python | `test_register_rejects_duplicate_username` |
| `backend/api/tests/test_auth.py` | 72 | Python | `test_login_requires_persisted_credentials` |
| `backend/api/tests/test_auth.py` | 96 | Python | `test_register_rejects_short_passwords` |
| `backend/api/tests/test_auth.py` | 108 | Python | `test_register_rejects_invalid_username_characters` |
| `backend/api/tests/test_auth.py` | 120 | Python | `test_protected_endpoints_reject_invalid_status_payloads` |
| `backend/api/tests/test_auth.py` | 153 | Python | `test_template_and_broadcast_require_non_empty_messages` |
| `backend/api/tests/test_auth.py` | 175 | Python | `test_staff_cannot_access_admin_only_stats_or_export` |
| `backend/api/tests/test_auth.py` | 200 | Python | `test_import_does_not_touch_database_at_module_import` |
| `backend/api/tests/test_auth.py` | 217 | Python | `test_queue_broadcast_event_uses_redis_when_available` |
| `backend/api/tests/test_auth.py` | 219 | Python | `__init__` |
| `backend/api/tests/test_auth.py` | 222 | Python | `rpush` |
| `backend/api/tests/test_auth.py` | 233 | Python | `__init__` |
| `backend/api/tests/test_auth.py` | 248 | Python | `test_webhook_rejects_invalid_line_signature_when_secret_is_configured` |
| `backend/api/tests/test_auth.py` | 265 | Python | `test_webhook_accepts_valid_line_signature_when_secret_is_configured` |
| `backend/api/tests/test_auth.py` | 285 | Python | `test_webhook_ignores_non_message_events` |
| `backend/api/tests/test_auth.py` | 308 | Python | `test_webhook_ignores_empty_text_messages` |
| `backend/api/tests/test_auth.py` | 332 | Python | `test_import_with_default_sqlite_does_not_create_db_file` |
| `backend/api/tests/test_auth.py` | 352 | Python | `test_branding_can_be_updated_by_admin` |
| `backend/api/tests/test_auth.py` | 370 | Python | `test_team_member_management_enforces_roles` |
| `backend/api/tests/test_auth.py` | 401 | Python | `test_tiktok_csv_export_returns_tenant_scoped_rows` |
| `backend/api/tests/test_auth.py` | 429 | Python | `test_webhook_posts_sheet_sync_when_webhook_url_is_configured` |
| `backend/api/tests/test_auth.py` | 437 | Python | `__init__` |
| `backend/api/tests/test_auth.py` | 440 | Python | `raise_for_status` |
| `backend/api/tests/test_auth.py` | 443 | Python | `json` |
| `backend/api/tests/test_auth.py` | 446 | Python | `fake_post` |
| `backend/worker/tests/test_worker.py` | 9 | Python | `__init__` |
| `backend/worker/tests/test_worker.py` | 12 | Python | `send` |
| `backend/worker/tests/test_worker.py` | 16 | Python | `get` |
| `backend/worker/tests/test_worker.py` | 21 | Python | `flush` |
| `backend/worker/tests/test_worker.py` | 24 | Python | `close` |
| `backend/worker/tests/test_worker.py` | 29 | Python | `__init__` |
| `backend/worker/tests/test_worker.py` | 32 | Python | `poll` |
| `backend/worker/tests/test_worker.py` | 35 | Python | `commit` |
| `backend/worker/tests/test_worker.py` | 39 | Python | `close` |
| `backend/worker/tests/test_worker.py` | 44 | Python | `__init__` |
| `backend/worker/tests/test_worker.py` | 47 | Python | `brpop` |
| `backend/worker/tests/test_worker.py` | 50 | Python | `close` |
| `backend/worker/tests/test_worker.py` | 66 | Python | `make_worker` |
| `backend/worker/tests/test_worker.py` | 109 | Python | `test_process_message_persists_rows_and_increments_replies` |
| `backend/worker/tests/test_worker.py` | 136 | Python | `test_process_broadcast_filters_recipients_and_updates_campaign` |
| `backend/worker/tests/test_worker.py` | 172 | Python | `test_process_broadcast_marks_failures_when_line_push_fails` |
| `backend/worker/tests/test_worker.py` | 200 | Python | `test_publish_dead_letter_includes_context` |
| `backend/worker/worker.py` | 36 | Python | `from_env` |
| `backend/worker/worker.py` | 53 | Python | `__init__` |
| `backend/worker/worker.py` | 74 | Python | `_make_engine` |
| `backend/worker/worker.py` | 80 | Python | `stop` |
| `backend/worker/worker.py` | 84 | Python | `close` |
| `backend/worker/worker.py` | 96 | Python | `run` |
| `backend/worker/worker.py` | 144 | Python | `process_record` |
| `backend/worker/worker.py` | 157 | Python | `process_broadcast` |
| `backend/worker/worker.py` | 262 | Python | `process_message` |
| `backend/worker/worker.py` | 310 | Python | `line_push` |
| `backend/worker/worker.py` | 338 | Python | `publish_dead_letter` |
| `backend/worker/worker.py` | 348 | Python | `_required_text` |
| `backend/worker/worker.py` | 355 | Python | `_optional_text` |
| `backend/worker/worker.py` | 360 | Python | `_first_text` |
| `backend/worker/worker.py` | 368 | Python | `_required_int` |
| `backend/worker/worker.py` | 375 | Python | `_optional_int` |
| `backend/worker/worker.py` | 385 | Python | `configure_logging` |
| `backend/worker/worker.py` | 392 | Python | `main` |
| `billing/stripe.py` | 8 | Python | `create_checkout` |
| `codex.sh` | 17 | Shell | `log` |
| `codex.sh` | 21 | Shell | `require_cmd` |
| `codex.sh` | 28 | Shell | `random_secret` |
| `codex.sh` | 33 | Shell | `ensure_repo` |
| `codex.sh` | 52 | Shell | `resolve_paths` |
| `codex.sh` | 67 | Shell | `install_dependencies` |
| `codex.sh` | 75 | Shell | `setup_postgres_redis` |
| `codex.sh` | 107 | Shell | `apply_schema` |
| `codex.sh` | 123 | Shell | `create_env_file` |
| `codex.sh` | 151 | Shell | `create_run_script` |
| `codex.sh` | 165 | Shell | `install_codex_cli` |
| `codex.sh` | 171 | Shell | `write_codex_config` |
| `codex.sh` | 192 | Shell | `verify_openai_api_key` |
| `codex.sh` | 198 | Shell | `run_codex_auto_setup` |
| `codex.sh` | 216 | Shell | `launch_application` |
| `codex.sh` | 221 | Shell | `package_k8s_bundle` |
| `codex.sh` | 243 | Shell | `usage` |
| `codex.sh` | 263 | Shell | `main` |
| `core/affiliate/service.ts` | 10 | JS/TS | `awardReferralCredit` |
| `core/affiliate/service.ts` | 19 | JS/TS | `createReferral` |
| `dl/orderbook_model.py` | 10 | Python | `__init__` |
| `dl/orderbook_model.py` | 15 | Python | `forward` |
| `dl/orderbook_model.py` | 21 | Python | `build_feature_vector` |
| `dl/orderbook_model.py` | 34 | Python | `softmax_probs` |
| `execution/smartRouter.ts` | 9 | JS/TS | `bestExecution` |

_Display capped at 300 of 491 detected functions/methods._

### Automation pipeline files

- `.github/workflows/ci.yml`
- `.github/workflows/deploy-k8s.yml`
- `.github/workflows/deploy.yml`
- `.github/workflows/self-healing-cicd.yml`
- `deploy.sh`
- `docker-compose.yml`
- `docker/docker-compose.enterprise.yml`
- `infra/docker-compose.yml`
- `infrastructure/scripts/auto-fix-pipeline.sh`
- `install_and_run.sh`
- `install_full_stack.sh`
- `installer/install.sh`
- `k8s/api-deployment.yaml`
- `k8s/backup/cronjob.yaml`
- `k8s/deployment.yaml`
- `k8s/deployments/api-deployment.yaml`
- `k8s/deployments/dashboard-deployment.yaml`
- `k8s/deployments/scheduler-deployment.yaml`
- `k8s/deployments/worker-deployment.yaml`
- `k8s/enterprise-api-deployment.yaml`
- `k8s/ingress/cloudflare-tunnel-deployment.yaml`
- `k8s/worker-deployment.yaml`
- `k8s/zlinebot-deployment.yaml`
- `master_installer.sh`
- `one-click-deploy-config-installer-starter.sh`
- `scripts/deploy-images.sh`
- `stack-workflow-manager.sh`
- `ubuntu_stack_installer.sh`
- `zBAutomos/infra/docker-compose.yml`
- `zBAutomos/install.sh`
- `zBAutomos/install_zba.sh`
- `zeaz_ai_full_stack_installer.sh`

## `cvsz/zeaz-platform`

Status: **cloned**; commit `b22f68a3f6b8`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

- `install-ubuntu-dev-stack.sh`

## `cvsz/zeapay`

Status: **blocked**; commit `blocked`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

None detected or repository blocked.

## `cvsz/zcino`

Status: **cloned**; commit `f2e95ee2b857`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `frontend/components/games/game-card.tsx` | 34 | `/api/track/click` |
| `frontend/components/games/game-card.tsx` | 40 | `/api/track/click` |
| `internal/transport/policy_middleware_test.go` | 37 | `/wallet/balance` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `cmd/compliance/main.go` | 17 | Go | `main` |
| `cmd/compliance/main.go` | 43 | Go | `loadBlockedCountries` |
| `cmd/compliance/main.go` | 57 | Go | `getEnv` |
| `cmd/compliance/main_test.go` | 5 | Go | `TestLoadBlockedCountriesDefaults` |
| `cmd/compliance/main_test.go` | 12 | Go | `TestLoadBlockedCountriesNormalizesInput` |
| `frontend/app/api/games/route.ts` | 45 | JS/TS | `readNumber` |
| `frontend/app/api/games/route.ts` | 52 | JS/TS | `readFilters` |
| `frontend/app/api/games/route.ts` | 62 | JS/TS | `isGameCategory` |
| `frontend/app/api/games/route.ts` | 66 | JS/TS | `createThumbnail` |
| `frontend/app/api/games/route.ts` | 78 | JS/TS | `mapCatalogGame` |
| `frontend/app/api/games/route.ts` | 94 | JS/TS | `catalogQuery` |
| `frontend/app/api/games/route.ts` | 111 | JS/TS | `fetchCatalogProviders` |
| `frontend/app/api/games/route.ts` | 127 | JS/TS | `fetchCatalogGames` |
| `frontend/app/api/games/route.ts` | 170 | JS/TS | `GET` |
| `frontend/app/api/track/click/route.ts` | 5 | JS/TS | `POST` |
| `frontend/app/providers/query-provider.tsx` | 7 | JS/TS | `QueryProvider` |
| `frontend/components/games/filter-sidebar.tsx` | 20 | JS/TS | `FilterSidebar` |
| `frontend/components/games/filter-sidebar.tsx` | 100 | JS/TS | `onChange` |
| `frontend/components/games/filter-sidebar.tsx` | 115 | JS/TS | `onChange` |
| `frontend/components/games/game-card.tsx` | 13 | JS/TS | `GameCard` |
| `frontend/components/games/game-card.tsx` | 14 | JS/TS | `trackClick` |
| `frontend/components/games/game-card.tsx` | 97 | JS/TS | `getQueryParam` |
| `frontend/components/games/game-card.tsx` | 102 | JS/TS | `getSessionID` |
| `frontend/components/games/game-grid.tsx` | 16 | JS/TS | `GameGrid` |
| `frontend/components/games/game-lobby.tsx` | 13 | JS/TS | `GameLobby` |
| `frontend/components/games/search-bar.tsx` | 3 | JS/TS | `SearchBar` |
| `frontend/hooks/use-games.ts` | 8 | JS/TS | `useGames` |
| `frontend/lib/api.ts` | 5 | JS/TS | `buildGamesQuery` |
| `frontend/lib/api.ts` | 20 | JS/TS | `fetchGames` |
| `frontend/lib/games-data.ts` | 40 | JS/TS | `createThumbnail` |
| `frontend/lib/games-data.ts` | 70 | JS/TS | `getFilterOptions` |
| `frontend/lib/games-data.ts` | 81 | JS/TS | `getGames` |
| `internal/auth/jwt.go` | 31 | Go | `NewTokenManager` |
| `internal/auth/jwt.go` | 35 | Go | `Generate` |
| `internal/auth/jwt.go` | 51 | Go | `Validate` |
| `internal/auth/jwt.go` | 68 | Go | `WithClaims` |
| `internal/auth/jwt.go` | 72 | Go | `ClaimsFromContext` |
| `internal/auth/jwt.go` | 77 | Go | `HasRole` |
| `internal/auth/jwt_test.go` | 8 | Go | `TestTokenManagerGenerateValidate` |
| `internal/auth/jwt_test.go` | 27 | Go | `TestTokenManagerRejectsWrongSecret` |
| `internal/cache/game_cache.go` | 32 | Go | `NewRedisGameCache` |
| `internal/cache/game_cache.go` | 36 | Go | `GetGame` |
| `internal/cache/game_cache.go` | 42 | Go | `SetGame` |
| `internal/cache/game_cache.go` | 46 | Go | `GetGamePage` |
| `internal/cache/game_cache.go` | 52 | Go | `SetGamePage` |
| `internal/cache/game_cache.go` | 56 | Go | `GetProviders` |
| `internal/cache/game_cache.go` | 62 | Go | `SetProviders` |
| `internal/cache/game_cache.go` | 66 | Go | `getJSON` |
| `internal/cache/game_cache.go` | 80 | Go | `setJSON` |
| `internal/cache/game_cache.go` | 91 | Go | `gameKey` |
| `internal/cache/game_cache.go` | 95 | Go | `providersKey` |
| `internal/cache/game_cache.go` | 99 | Go | `gamesKey` |
| `internal/config/config.go` | 54 | Go | `Load` |
| `internal/config/config.go` | 144 | Go | `getEnv` |
| `internal/config/config.go` | 151 | Go | `getEnvInt` |
| `internal/config/config.go` | 163 | Go | `getEnvDuration` |
| `internal/config/config.go` | 175 | Go | `getEnvDurationOrDefault` |
| `internal/database/postgres.go` | 12 | Go | `NewPostgresPool` |
| `internal/database/redis.go` | 12 | Go | `NewRedisClient` |
| `internal/domain/errors.go` | 14 | Go | `Error` |
| `internal/domain/errors.go` | 18 | Go | `Unwrap` |
| `internal/domain/game.go` | 20 | Go | `IsValid` |
| `internal/domain/game.go` | 40 | Go | `Validate` |
| `internal/domain/game.go` | 81 | Go | `Offset` |
| `internal/domain/tracking.go` | 19 | Go | `IsValid` |
| `internal/domain/tracking.go` | 47 | Go | `Validate` |
| `internal/handler/auth_handler.go` | 30 | Go | `NewAuthHandler` |
| `internal/handler/auth_handler.go` | 34 | Go | `Token` |
| `internal/handler/auth_handler.go` | 54 | Go | `decodeLoginRequest` |
| `internal/handler/catalog_handler.go` | 30 | Go | `NewCatalogHandler` |
| `internal/handler/catalog_handler.go` | 34 | Go | `ListGames` |
| `internal/handler/catalog_handler.go` | 49 | Go | `GetGame` |
| `internal/handler/catalog_handler.go` | 65 | Go | `ListProviders` |
| `internal/handler/catalog_handler.go` | 74 | Go | `handleServiceError` |
| `internal/handler/catalog_handler.go` | 87 | Go | `parseListGamesQuery` |
| `internal/handler/catalog_handler.go` | 112 | Go | `parseRTPRange` |
| `internal/handler/catalog_handler.go` | 142 | Go | `parsePositiveInt` |
| `internal/handler/catalog_handler.go` | 153 | Go | `writeJSON` |
| `internal/handler/catalog_handler.go` | 159 | Go | `writeError` |
| `internal/handler/tracking_handler.go` | 52 | Go | `NewTrackingHandler` |
| `internal/handler/tracking_handler.go` | 56 | Go | `TrackImpression` |
| `internal/handler/tracking_handler.go` | 60 | Go | `TrackClick` |
| `internal/handler/tracking_handler.go` | 64 | Go | `track` |
| `internal/handler/tracking_handler.go` | 89 | Go | `handleTrackingError` |
| `internal/handler/tracking_handler.go` | 102 | Go | `decodeTrackingRequest` |
| `internal/handler/tracking_handler.go` | 116 | Go | `toEvent` |
| `internal/handler/tracking_handler_test.go` | 22 | Go | `Track` |
| `internal/handler/tracking_handler_test.go` | 27 | Go | `Start` |
| `internal/handler/tracking_handler_test.go` | 28 | Go | `Stop` |
| `internal/handler/tracking_handler_test.go` | 30 | Go | `TestTrackingHandlerAcceptsClick` |
| `internal/handler/tracking_handler_test.go` | 60 | Go | `TestTrackingRequestUsesProvidedOccurredAt` |
| `internal/logger/logger.go` | 5 | Go | `New` |
| `internal/provider/mock_providers.go` | 19 | Go | `NewMockPGSoftProvider` |
| `internal/provider/mock_providers.go` | 26 | Go | `GetGames` |
| `internal/provider/mock_providers.go` | 63 | Go | `NewMockPragmaticPlayProvider` |
| `internal/provider/mock_providers.go` | 70 | Go | `GetGames` |
| `internal/provider/mock_providers.go` | 91 | Go | `varianceToVolatility` |
| `internal/provider/mock_providers.go` | 118 | Go | `NewMockEvolutionProvider` |
| `internal/provider/mock_providers.go` | 125 | Go | `GetGames` |
| `internal/provider/normalizer.go` | 11 | Go | `NewNormalizer` |
| `internal/provider/normalizer.go` | 16 | Go | `GetGames` |
| `internal/provider/normalizer_test.go` | 9 | Go | `TestMockProvidersNormalizeToCanonicalCatalog` |
| `internal/provider/normalizer_test.go` | 39 | Go | `TestNormalizeGameConvertsProviderFields` |
| `internal/provider/provider.go` | 41 | Go | `NormalizeGame` |
| `internal/provider/provider.go` | 67 | Go | `stableGameID` |
| `internal/provider/provider.go` | 72 | Go | `normalizeCategory` |
| `internal/provider/provider.go` | 86 | Go | `normalizeRTP` |
| `internal/provider/provider.go` | 93 | Go | `normalizeVolatility` |
| `internal/provider/provider.go` | 104 | Go | `normalizeThumbnailURL` |
| `internal/repository/game_repository.go` | 26 | Go | `NewPostgresGameRepository` |
| `internal/repository/game_repository.go` | 30 | Go | `ListGames` |
| `internal/repository/game_repository.go` | 67 | Go | `GetGameByID` |
| `internal/repository/game_repository.go` | 87 | Go | `ListProviders` |
| `internal/repository/game_repository.go` | 114 | Go | `buildGameFilters` |
| `internal/repository/game_repository.go` | 149 | Go | `scanGame` |
| `internal/repository/tracking_repository.go` | 22 | Go | `NewPostgresTrackingRepository` |
| `internal/repository/tracking_repository.go` | 26 | Go | `InsertTrackingEvents` |
| `internal/repository/tracking_repository.go` | 67 | Go | `nullString` |
| `internal/repository/tracking_repository.go` | 74 | Go | `nullInt64` |
| `internal/repository/tracking_repository.go` | 81 | Go | `metadataOrEmpty` |
| `internal/service/catalog_service.go` | 31 | Go | `NewCatalogService` |
| `internal/service/catalog_service.go` | 35 | Go | `ListGames` |
| `internal/service/catalog_service.go` | 68 | Go | `GetGame` |
| `internal/service/catalog_service.go` | 88 | Go | `ListProviders` |
| `internal/service/catalog_service.go` | 104 | Go | `validateFilter` |
| `internal/service/catalog_service.go` | 123 | Go | `IsNotFound` |
| `internal/service/catalog_service_test.go` | 17 | Go | `GetGame` |
| `internal/service/catalog_service_test.go` | 20 | Go | `SetGame` |
| `internal/service/catalog_service_test.go` | 21 | Go | `GetGamePage` |
| `internal/service/catalog_service_test.go` | 24 | Go | `SetGamePage` |
| `internal/service/catalog_service_test.go` | 27 | Go | `GetProviders` |
| `internal/service/catalog_service_test.go` | 28 | Go | `SetProviders` |
| `internal/service/catalog_service_test.go` | 34 | Go | `ListGames` |
| `internal/service/catalog_service_test.go` | 38 | Go | `TestCatalogServiceRejectsInvalidRTPRange` |
| `internal/service/tracking_service.go` | 43 | Go | `NewTrackingService` |
| `internal/service/tracking_service.go` | 62 | Go | `Start` |
| `internal/service/tracking_service.go` | 68 | Go | `Track` |
| `internal/service/tracking_service.go` | 87 | Go | `Stop` |
| `internal/service/tracking_service.go` | 103 | Go | `run` |
| `internal/service/tracking_service_test.go` | 21 | Go | `InsertTrackingEvents` |
| `internal/service/tracking_service_test.go` | 29 | Go | `count` |
| `internal/service/tracking_service_test.go` | 35 | Go | `TestTrackingServiceBatchesBySize` |
| `internal/service/tracking_service_test.go` | 52 | Go | `TestTrackingServiceRejectsInvalidEvent` |
| `internal/service/tracking_service_test.go` | 62 | Go | `validTrackingEvent` |
| `internal/service/tracking_service_test.go` | 74 | Go | `assertEventually` |
| `internal/transport/auth_middleware.go` | 10 | Go | `optionalAuth` |
| `internal/transport/auth_middleware.go` | 28 | Go | `requireRole` |
| `internal/transport/auth_middleware.go` | 45 | Go | `authenticateBearer` |
| `internal/transport/auth_middleware.go` | 57 | Go | `writeAuthError` |
| `internal/transport/policy_middleware.go` | 61 | Go | `policyGuard` |
| `internal/transport/policy_middleware.go` | 91 | Go | `inspectPolicy` |
| `internal/transport/policy_middleware.go` | 109 | Go | `shouldInspectBody` |
| `internal/transport/policy_middleware.go` | 118 | Go | `readAndRestoreBody` |
| `internal/transport/policy_middleware.go` | 133 | Go | `evaluatePolicy` |
| `internal/transport/policy_middleware.go` | 153 | Go | `tokenSet` |
| `internal/transport/policy_middleware.go` | 161 | Go | `writePolicyBlock` |
| `internal/transport/policy_middleware_test.go` | 13 | Go | `TestPolicyGuardAllowsCatalogRequests` |
| `internal/transport/policy_middleware_test.go` | 31 | Go | `TestPolicyGuardBlocksProhibitedEndpointPaths` |
| `internal/transport/policy_middleware_test.go` | 58 | Go | `TestPolicyGuardBlocksProhibitedRequestBody` |
| `internal/transport/policy_middleware_test.go` | 71 | Go | `TestPolicyGuardRestoresAllowedRequestBody` |
| `internal/transport/policy_middleware_test.go` | 93 | Go | `assertPolicyBlocked` |
| `internal/transport/rate_limit.go` | 24 | Go | `newRateLimiter` |
| `internal/transport/rate_limit.go` | 39 | Go | `middleware` |
| `internal/transport/rate_limit.go` | 50 | Go | `allow` |
| `internal/transport/rate_limit.go` | 76 | Go | `cleanup` |
| `internal/transport/rate_limit.go` | 84 | Go | `clientIP` |
| `internal/transport/rate_limit_test.go` | 10 | Go | `TestRateLimiterBlocksAfterBurst` |
| `internal/transport/router.go` | 15 | Go | `NewRouter` |
| `internal/transport/router.go` | 42 | Go | `requestLogger` |
| `internal/transport/router.go` | 59 | Go | `recoverer` |
| `internal/transport/router.go` | 76 | Go | `WriteHeader` |
| `main.go` | 27 | Go | `main` |

### Automation pipeline files

- `infra/docker-compose.yml`

## `cvsz/zGaming`

Status: **cloned**; commit `10f08a39a3ea`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `api/gateway/server.ts` | 60 | `/health` |
| `api/gateway/server.ts` | 62 | `/ready` |
| `api/gateway/server.ts` | 76 | `/metrics` |
| `api/gateway/server.ts` | 90 | `/auth/token` |
| `api/gateway/server.ts` | 116 | `/internal/webhook` |
| `generator/phases/107-meta-orchestrator.sh` | 156 | `/health` |
| `generator/phases/107-meta-orchestrator.sh` | 158 | `/auth/token` |
| `generator/phases/107-meta-orchestrator.sh` | 168 | `/secure` |
| `generator/phases/60-frontend.sh` | 58 | `/api/login.php` |
| `generator/phases/60-frontend.sh` | 68 | `/api/launch/pragmatic.php` |
| `generator/phases/60-frontend.sh` | 109 | `/api/admin/users.php` |
| `infra/kubernetes/api-deployment.yaml` | 10 | `/health` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `api/gateway/server.ts` | 9 | JS/TS | `startServer` |
| `api/gateway/webhook-auth.ts` | 13 | JS/TS | `signWebhookPayload` |
| `api/gateway/webhook-auth.ts` | 17 | JS/TS | `verifyWebhookSignature` |
| `backend/api/login.php` | 7 | PHP | `consume_nonce` |
| `backend/api/regulator/common.php` | 11 | PHP | `regulator_guard` |
| `backend/api/regulator/common.php` | 24 | PHP | `pagination` |
| `backend/api/regulator/common.php` | 31 | PHP | `export_data` |
| `backend/audit/AuditTrail.php` | 7 | PHP | `__construct` |
| `backend/audit/AuditTrail.php` | 11 | PHP | `append` |
| `backend/compliance/KycService.php` | 9 | PHP | `__construct` |
| `backend/compliance/KycService.php` | 13 | PHP | `assessUser` |
| `backend/compliance/KycService.php` | 24 | PHP | `isWithdrawalAllowed` |
| `backend/compliance/KycService.php` | 31 | PHP | `shouldFlagHighRisk` |
| `backend/compliance/OnfidoProvider.php` | 9 | PHP | `verifyUser` |
| `backend/compliance/OnfidoProvider.php` | 20 | PHP | `sanctionsCheck` |
| `backend/compliance/SumsubProvider.php` | 9 | PHP | `verifyUser` |
| `backend/compliance/SumsubProvider.php` | 20 | PHP | `sanctionsCheck` |
| `backend/finance/AmlEngine.php` | 7 | PHP | `__construct` |
| `backend/finance/AmlEngine.php` | 11 | PHP | `scanUser` |
| `backend/finance/AmlEngine.php` | 18 | PHP | `flagLargeTransactions` |
| `backend/finance/AmlEngine.php` | 27 | PHP | `flagRapidTransactions` |
| `backend/finance/AmlEngine.php` | 37 | PHP | `flagKycRisk` |
| `backend/finance/AmlEngine.php` | 51 | PHP | `createFlag` |
| `backend/finance/FraudEngine.php` | 7 | PHP | `__construct` |
| `backend/finance/FraudEngine.php` | 11 | PHP | `assess` |
| `backend/finance/ReconciliationEngine.php` | 7 | PHP | `__construct` |
| `backend/finance/ReconciliationEngine.php` | 11 | PHP | `runForDate` |
| `backend/finance/SettlementWorker.php` | 11 | PHP | `__construct` |
| `backend/finance/SettlementWorker.php` | 16 | PHP | `runOnce` |
| `backend/finance/SettlementWorker.php` | 43 | PHP | `processJob` |
| `backend/finance/WalletTierService.php` | 13 | PHP | `__construct` |
| `backend/finance/WalletTierService.php` | 21 | PHP | `creditDeposit` |
| `backend/finance/WalletTierService.php` | 30 | PHP | `requestWithdrawal` |
| `backend/finance/WalletTierService.php` | 84 | PHP | `isReadOnlyMode` |
| `backend/finance/WalletTierService.php` | 89 | PHP | `isWithdrawalsDisabled` |
| `backend/finance/WalletTierService.php` | 94 | PHP | `isUserFrozen` |
| `backend/finance/WalletTierService.php` | 101 | PHP | `isControlEnabled` |
| `backend/finance/WalletTierService.php` | 108 | PHP | `approveColdWithdrawal` |
| `backend/finance/WalletTierService.php` | 114 | PHP | `topUpHotFromCold` |
| `backend/finance/WalletTierService.php` | 125 | PHP | `newTransactionId` |
| `backend/interfaces/KycProviderInterface.php` | 8 | PHP | `verifyUser` |
| `backend/interfaces/KycProviderInterface.php` | 10 | PHP | `sanctionsCheck` |
| `backend/interfaces/SignerInterface.php` | 7 | PHP | `signTransaction` |
| `backend/interfaces/SignerInterface.php` | 9 | PHP | `getAddress` |
| `backend/interfaces/SignerInterface.php` | 11 | PHP | `verifySignature` |
| `backend/lib/Database.php` | 9 | PHP | `conn` |
| `backend/lib/auth.php` | 5 | PHP | `auth_env` |
| `backend/lib/auth.php` | 18 | PHP | `base64url_encode` |
| `backend/lib/auth.php` | 23 | PHP | `jwt_sign` |
| `backend/lib/auth.php` | 34 | PHP | `hash_nonce` |
| `backend/lib/auth.php` | 39 | PHP | `read_json_input` |
| `backend/lib/auth.php` | 48 | PHP | `json_response` |
| `backend/observability/StructuredLogger.php` | 7 | PHP | `__construct` |
| `backend/observability/StructuredLogger.php` | 11 | PHP | `info` |
| `backend/observability/StructuredLogger.php` | 16 | PHP | `warning` |
| `backend/observability/StructuredLogger.php` | 21 | PHP | `error` |
| `backend/observability/StructuredLogger.php` | 26 | PHP | `log` |
| `backend/security/AdminAccess.php` | 15 | PHP | `__construct` |
| `backend/security/AdminAccess.php` | 19 | PHP | `authorize` |
| `backend/security/ServiceAuth.php` | 10 | PHP | `__construct` |
| `backend/security/ServiceAuth.php` | 14 | PHP | `mintToken` |
| `backend/security/ServiceAuth.php` | 34 | PHP | `verifyIncoming` |
| `backend/security/ServiceAuth.php` | 66 | PHP | `parse_service_secrets_env` |
| `backend/security/ServiceAuth.php` | 78 | PHP | `require_service_auth` |
| `backend/signer/AwsKmsSigner.php` | 10 | PHP | `__construct` |
| `backend/signer/AwsKmsSigner.php` | 17 | PHP | `signTransaction` |
| `backend/signer/AwsKmsSigner.php` | 23 | PHP | `getAddress` |
| `backend/signer/AwsKmsSigner.php` | 28 | PHP | `verifySignature` |
| `backend/signer/FireblocksSigner.php` | 10 | PHP | `__construct` |
| `backend/signer/FireblocksSigner.php` | 18 | PHP | `signTransaction` |
| `backend/signer/FireblocksSigner.php` | 24 | PHP | `getAddress` |
| `backend/signer/FireblocksSigner.php` | 29 | PHP | `verifySignature` |
| `backend/signer/HsmSigner.php` | 10 | PHP | `__construct` |
| `backend/signer/HsmSigner.php` | 17 | PHP | `signTransaction` |
| `backend/signer/HsmSigner.php` | 23 | PHP | `getAddress` |
| `backend/signer/HsmSigner.php` | 28 | PHP | `verifySignature` |
| `backend/signer/LocalSigner.php` | 10 | PHP | `__construct` |
| `backend/signer/LocalSigner.php` | 18 | PHP | `signTransaction` |
| `backend/signer/LocalSigner.php` | 24 | PHP | `getAddress` |
| `backend/signer/LocalSigner.php` | 29 | PHP | `verifySignature` |
| `backend/signer/SigningAuditLogger.php` | 9 | PHP | `__construct` |
| `backend/signer/SigningAuditLogger.php` | 14 | PHP | `logRequest` |
| `core/orchestrator/kernel.ts` | 5 | JS/TS | `init` |
| `core/orchestrator/kernel.ts` | 6 | JS/TS | `shutdown` |
| `core/orchestrator/kernel.ts` | 13 | JS/TS | `register` |
| `core/orchestrator/kernel.ts` | 22 | JS/TS | `boot` |
| `core/orchestrator/kernel.ts` | 34 | JS/TS | `shutdown` |
| `core/orchestrator/kernel.ts` | 43 | JS/TS | `getBus` |
| `core/plugin-loader/loader.ts` | 4 | JS/TS | `loadModules` |
| `factory/publisher/resumable-upload.ts` | 22 | JS/TS | `acceptChunk` |
| `factory/publisher/resumable-upload.ts` | 39 | JS/TS | `finalize` |
| `factory/render/stateless-render-farm.ts` | 18 | JS/TS | `constructor` |
| `factory/render/stateless-render-farm.ts` | 24 | JS/TS | `dispatch` |
| `factory/render/stateless-render-farm.ts` | 40 | JS/TS | `hpaTargetMetrics` |
| `factory/scout/autonomous-content-factory.ts` | 9 | JS/TS | `autonomousContentFactoryLoop` |
| `frontend-admin/src/transaction-review.ts` | 3 | JS/TS | `validateIntent` |
| `frontend-player/src/intent-simulation.ts` | 10 | JS/TS | `renderIntent` |
| `generator/lib/assert.sh` | 27 | Shell | `fail` |
| `generator/lib/assert.sh` | 32 | Shell | `ok` |
| `generator/meta-master.sh` | 31 | Shell | `print_header` |
| `generator/meta-master.sh` | 38 | Shell | `usage` |
| `generator/meta-master.sh` | 59 | Shell | `fail` |
| `generator/meta-master.sh` | 106 | Shell | `assert_layout` |
| `generator/meta-master.sh` | 110 | Shell | `validate_phase_catalog` |
| `generator/meta-master.sh` | 133 | Shell | `phase_index` |
| `generator/meta-master.sh` | 147 | Shell | `run_phase` |
| `generator/meta-master.sh` | 159 | Shell | `print_summary` |
| `generator/meta-master.sh` | 178 | Shell | `run_all` |
| `generator/meta-master.sh` | 233 | Shell | `main` |
| `generator/phases/110-runtest-report.sh` | 21 | Shell | `log_ok` |
| `generator/phases/110-runtest-report.sh` | 24 | Shell | `log_fail` |
| `generator/phases/70-nginx.sh` | 11 | Shell | `require_docker_ready` |
| `generator/phases/70-nginx.sh` | 23 | Shell | `warn_linux_desktop_credsstore` |
| `generator/phases/85-backup.sh` | 45 | Shell | `fix_linux_docker_credsstore` |
| `generator/phases/85-backup.sh` | 58 | Shell | `require_docker_ready` |
| `generator/phases/85-backup.sh` | 78 | Shell | `ensure_db_container` |
| `generator/phases/87-dr-test.sh` | 11 | Shell | `cleanup_named_containers` |
| `generator/stages/lib/common.sh` | 23 | Shell | `log` |
| `generator/stages/lib/common.sh` | 31 | Shell | `progress` |
| `generator/stages/lib/common.sh` | 35 | Shell | `check_stage_state` |
| `generator/stages/lib/common.sh` | 44 | Shell | `mark_stage_completed` |
| `generator/stages/lib/common.sh` | 50 | Shell | `generate_file_checksums` |
| `generator/stages/lib/common.sh` | 59 | Shell | `generate_manifest` |
| `generator/stages/lib/common.sh` | 81 | Shell | `rollback_stage` |
| `generator/stages/lib/common.sh` | 85 | Shell | `run_phase` |
| `install_and_run.sh` | 6 | Shell | `log` |
| `install_and_run.sh` | 7 | Shell | `fail` |
| `install_and_run.sh` | 9 | Shell | `usage` |
| `installer/zgaming-ultra-installer.sh` | 22 | Shell | `log_json` |
| `installer/zgaming-ultra-installer.sh` | 28 | Shell | `print_banner` |
| `installer/zgaming-ultra-installer.sh` | 37 | Shell | `usage` |
| `installer/zgaming-ultra-installer.sh` | 52 | Shell | `print_plan` |
| `installer/zgaming-ultra-installer.sh` | 69 | Shell | `rollback` |
| `installer/zgaming-ultra-installer.sh` | 79 | Shell | `require_bins` |
| `installer/zgaming-ultra-installer.sh` | 95 | Shell | `check_runtime` |
| `installer/zgaming-ultra-installer.sh` | 104 | Shell | `extract_repo_metadata` |
| `installer/zgaming-ultra-installer.sh` | 120 | Shell | `generate_sbom_lite` |
| `installer/zgaming-ultra-installer.sh` | 147 | Shell | `emit_infrastructure_inventory` |
| `installer/zgaming-ultra-installer.sh` | 173 | Shell | `compliance_checks` |
| `installer/zgaming-ultra-installer.sh` | 203 | Shell | `write_audit_report` |
| `installer/zgaming-ultra-installer.sh` | 237 | Shell | `build_release_package` |
| `installer/zgaming-ultra-installer.sh` | 266 | Shell | `run_diagnostics` |
| `installer/zgaming-ultra-installer.sh` | 276 | Shell | `run_chaos` |
| `installer/zgaming-ultra-installer.sh` | 284 | Shell | `run_quick` |
| `installer/zgaming-ultra-installer.sh` | 296 | Shell | `run_full` |
| `installer/zgaming-ultra-installer.sh` | 304 | Shell | `run_no_cost` |
| `installer/zgaming-ultra-installer.sh` | 315 | Shell | `run_full_project` |
| `installer/zgaming-ultra-installer.sh` | 324 | Shell | `menu` |
| `installer/zgaming-ultra-installer.sh` | 340 | Shell | `main` |
| `modules/game-engine/cascade.ts` | 13 | JS/TS | `isBaseMatchSymbol` |
| `modules/game-engine/cascade.ts` | 17 | JS/TS | `findClusters` |
| `modules/game-engine/cascade.ts` | 83 | JS/TS | `applyCascadeAndRefill` |
| `modules/game-engine/engine.ts` | 24 | JS/TS | `spin` |
| `modules/game-engine/paytable.ts` | 12 | JS/TS | `payoutForCluster` |
| `modules/game-engine/provably-fair.ts` | 9 | JS/TS | `createServerSeedCommitment` |
| `modules/game-engine/provably-fair.ts` | 16 | JS/TS | `sha256` |
| `modules/game-engine/provably-fair.ts` | 20 | JS/TS | `verifyServerSeedReveal` |
| `modules/game-engine/provably-fair.ts` | 31 | JS/TS | `buildSeedTrace` |
| `modules/game-engine/provably-fair.ts` | 35 | JS/TS | `computeProvablyFairRoll` |
| `modules/game-engine/reels.ts` | 13 | JS/TS | `weightedChoice` |
| `modules/game-engine/reels.ts` | 33 | JS/TS | `generateGrid` |
| `modules/game-engine/rng.ts` | 21 | JS/TS | `nextFloat` |
| `modules/game-engine/rng.ts` | 29 | JS/TS | `nextInt` |
| `modules/game-engine/rng.ts` | 37 | JS/TS | `getCurrentNonce` |
| `modules/game-engine/rtp-controller.ts` | 26 | JS/TS | `snapshot` |
| `modules/game-engine/rtp-controller.ts` | 34 | JS/TS | `update` |
| `modules/game-engine/rtp-controller.ts` | 43 | JS/TS | `adjust` |
| `modules/ledger/ledger.ts` | 28 | JS/TS | `transferWithIdempotency` |
| `modules/ledger/ledger.ts` | 55 | JS/TS | `verifyLedgerIntegrity` |
| `modules/ledger/ledger.ts` | 123 | JS/TS | `getDerivedBalance` |
| `modules/ledger/ledger.ts` | 134 | JS/TS | `computeImmutableHash` |
| `modules/ledger/ledger.ts` | 148 | JS/TS | `computeLedgerChainHash` |
| `modules/ledger/ledger.ts` | 160 | JS/TS | `transferOnce` |
| `modules/ledger/ledger.ts` | 270 | JS/TS | `validateAmount` |
| `modules/ledger/ledger.ts` | 284 | JS/TS | `decimalPlaces` |
| `modules/ledger/ledger.ts` | 289 | JS/TS | `ensureCoreTables` |
| `modules/ledger/ledger.ts` | 347 | JS/TS | `ensureSerializable` |
| `modules/ledger/ledger.ts` | 351 | JS/TS | `safeRollback` |
| `modules/ledger/ledger.ts` | 359 | JS/TS | `isRetryableSerializationError` |
| `modules/wallet/chain-validation.ts` | 3 | JS/TS | `validateChainId` |
| `modules/wallet/eth.ts` | 8 | JS/TS | `hex` |
| `modules/wallet/eth.ts` | 12 | JS/TS | `hashHex` |
| `modules/wallet/eth.ts` | 23 | JS/TS | `transfer` |
| `modules/wallet/index.ts` | 21 | JS/TS | `constructor` |
| `modules/wallet/index.ts` | 29 | JS/TS | `transfer` |
| `modules/wallet/index.ts` | 33 | JS/TS | `simulateTransfer` |
| `modules/wallet/rpc.ts` | 11 | JS/TS | `constructor` |
| `modules/wallet/rpc.ts` | 17 | JS/TS | `resolve` |
| `modules/wallet/rpc.ts` | 27 | JS/TS | `withFallbacks` |
| `modules/wallet/signer.ts` | 9 | JS/TS | `constructor` |
| `modules/wallet/signer.ts` | 11 | JS/TS | `sign` |
| `modules/wallet/signer.ts` | 26 | JS/TS | `createEnvBackedHmacProvider` |
| `modules/wallet/signer.ts` | 41 | JS/TS | `constructor` |
| `modules/wallet/signer.ts` | 43 | JS/TS | `sign` |
| `modules/wallet/simulation.ts` | 7 | JS/TS | `simulateTransfer` |
| `modules/wallet/simulation.ts` | 45 | JS/TS | `failed` |
| `modules/wallet/sol.ts` | 8 | JS/TS | `base58Like` |
| `modules/wallet/sol.ts` | 12 | JS/TS | `hashHex` |
| `modules/wallet/sol.ts` | 23 | JS/TS | `transfer` |
| `modules/wallet/types.ts` | 53 | JS/TS | `sign` |
| `modules/wallet/wagmi-viem.ts` | 12 | JS/TS | `createDeterministicWagmiConfig` |
| `scripts/emergency-control.php` | 50 | PHP | `setControl` |
| `scripts/emergency-control.php` | 58 | PHP | `setUserFreeze` |
| `scripts/fix-all.sh` | 14 | Shell | `log` |
| `scripts/fix-all.sh` | 18 | Shell | `compute_hash` |
| `scripts/full_logic_scan.py` | 40 | Python | `read_text` |
| `scripts/full_logic_scan.py` | 44 | Python | `check_phase_integrity` |
| `scripts/full_logic_scan.py` | 102 | Python | `check_repo_signals` |
| `scripts/full_logic_scan.py` | 161 | Python | `build_upgrade_actions` |
| `scripts/full_logic_scan.py` | 182 | Python | `write_upgrade_markdown` |
| `scripts/full_logic_scan.py` | 201 | Python | `run` |
| `scripts/full_logic_scan.py` | 237 | Python | `parse_args` |
| `scripts/full_logic_scan.py` | 253 | Python | `main` |
| `scripts/meta-master-v2.sh` | 8 | Shell | `log` |
| `scripts/meta-master-v2.sh` | 23 | Shell | `check_docker` |
| `scripts/meta-master-v2.sh` | 27 | Shell | `check_db` |
| `scripts/meta-master-v2.sh` | 31 | Shell | `check_network` |
| `scripts/meta-master-v2.sh` | 39 | Shell | `recover_docker` |
| `scripts/meta-master-v2.sh` | 45 | Shell | `recover_db` |
| `scripts/meta-master-v2.sh` | 72 | Shell | `recover_network` |
| `scripts/meta-master-v2.sh` | 80 | Shell | `ensure_dep` |
| `scripts/meta-master-v2.sh` | 100 | Shell | `run_phase` |
| `scripts/security-baseline-check.sh` | 6 | Shell | `check_no_tracked_paths` |
| `scripts/security-baseline-check.sh` | 21 | Shell | `check_secret_manifests` |
| `services/adversarial/self_play.py` | 9 | Python | `generate` |
| `services/adversarial/self_play.py` | 18 | Python | `score` |
| `services/adversarial/self_play.py` | 27 | Python | `self_play` |
| `services/agents/agent_compliance.py` | 4 | Python | `evaluate_compliance` |
| `services/agents/agent_profit.py` | 4 | Python | `evaluate_profit` |
| `services/agents/agent_risk.py` | 4 | Python | `evaluate_risk` |
| `services/agents/economic_agent.py` | 5 | Python | `__init__` |
| `services/agents/economic_agent.py` | 10 | Python | `act` |
| `services/agents/economic_agent.py` | 16 | Python | `execute` |
| `services/autonomous-agent/agent.py` | 7 | Python | `main` |
| `services/autonomous-controller/controller.py` | 11 | Python | `decide_action` |
| `services/autonomous_controller/controller.py` | 11 | Python | `decide_action` |
| `services/bandit/segment_bandit.py` | 9 | Python | `__init__` |
| `services/bandit/segment_bandit.py` | 13 | Python | `select` |
| `services/bandit/segment_bandit.py` | 18 | Python | `update` |
| `services/causal/uplift_model.py` | 6 | Python | `__init__` |
| `services/causal/uplift_model.py` | 10 | Python | `fit` |
| `services/causal/uplift_model.py` | 14 | Python | `predict_uplift` |
| `services/central-bank/policy_ai.py` | 7 | Python | `__init__` |
| `services/central-bank/policy_ai.py` | 10 | Python | `update` |
| `services/central-bank/policy_ai.py` | 21 | Python | `supply_adjustment` |
| `services/feature-pipeline/flink_job.py` | 9 | Python | `compute_features` |
| `services/feature-pipeline/flink_job.py` | 19 | Python | `build_job` |
| `services/federated/fl_client.py` | 7 | Python | `__init__` |
| `services/federated/fl_client.py` | 10 | Python | `train` |
| `services/federated/fl_server.py` | 7 | Python | `__init__` |
| `services/federated/fl_server.py` | 10 | Python | `aggregate` |
| `services/game-theory/nash_solver.py` | 8 | Python | `best_response` |
| `services/game-theory/nash_solver.py` | 12 | Python | `compute_equilibrium` |
| `services/game_theory/nash_solver.py` | 8 | Python | `best_response` |
| `services/game_theory/nash_solver.py` | 12 | Python | `compute_equilibrium` |
| `services/hierarchy/manager.py` | 9 | Python | `__init__` |
| `services/hierarchy/manager.py` | 16 | Python | `route` |
| `services/hierarchy/manager.py` | 24 | Python | `decide` |
| `services/intents/intent_engine.py` | 5 | Python | `__init__` |
| `services/intents/intent_engine.py` | 12 | Python | `plan` |
| `services/intents/intent_engine.py` | 16 | Python | `optimize` |
| `services/market/mechanism.py` | 5 | Python | `__init__` |
| `services/market/mechanism.py` | 8 | Python | `submit` |
| `services/market/mechanism.py` | 11 | Python | `clear` |
| `services/mechanism/mechanism_design.py` | 5 | Python | `__init__` |
| `services/mechanism/mechanism_design.py` | 9 | Python | `payoff` |
| `services/meta-learning/meta_learner.py` | 9 | Python | `__init__` |
| `services/meta-learning/meta_learner.py` | 12 | Python | `update` |
| `services/meta-learning/meta_learner.py` | 17 | Python | `normalize` |
| `services/meta_learning/meta_learner.py` | 9 | Python | `__init__` |
| `services/meta_learning/meta_learner.py` | 12 | Python | `update` |
| `services/meta_learning/meta_learner.py` | 17 | Python | `normalize` |
| `services/mpc/secure_aggregation.py` | 7 | Python | `__init__` |
| `services/mpc/secure_aggregation.py` | 10 | Python | `split_secret` |
| `services/mpc/secure_aggregation.py` | 15 | Python | `reconstruct` |
| `services/multi-agent/orchestrator.py` | 11 | Python | `__init__` |
| `services/multi-agent/orchestrator.py` | 14 | Python | `decide` |
| `services/multi-agent/orchestrator.py` | 25 | Python | `evaluate` |
| `services/multi_agent/orchestrator.py` | 11 | Python | `__init__` |
| `services/multi_agent/orchestrator.py` | 14 | Python | `decide` |
| `services/multi_agent/orchestrator.py` | 25 | Python | `evaluate` |
| `services/policy/adaptive_policy.py` | 1 | Python | `decide` |
| `services/risk-engine/app.py` | 15 | Python | `infer` |
| `services/risk-engine/app.py` | 32 | Python | `_features` |
| `services/risk-engine/app.py` | 42 | Python | `main` |
| `services/rl-agent/online_rl.py` | 13 | Python | `__init__` |
| `services/rl-agent/online_rl.py` | 16 | Python | `predict` |
| `services/rl-agent/online_rl.py` | 20 | Python | `update` |
| `services/rl-agent/online_rl.py` | 24 | Python | `_softmax` |
| `services/rl-agent/online_rl.py` | 32 | Python | `main` |
| `services/rl-agent/reward_consumer.py` | 12 | Python | `__init__` |
| `services/rl-agent/reward_consumer.py` | 15 | Python | `update` |
| `services/rl-agent/reward_consumer.py` | 22 | Python | `compute_reward` |
| `services/rl-agent/reward_consumer.py` | 30 | Python | `main` |
| `services/training/pipeline.py` | 8 | Python | `train` |
| `services/treasury/treasury_ai.py` | 9 | Python | `__init__` |
| `services/treasury/treasury_ai.py` | 12 | Python | `optimize` |
| `services/treasury/treasury_ai.py` | 25 | Python | `softmax` |
| `services/world-model/model.py` | 10 | Python | `__init__` |
| `services/world-model/model.py` | 13 | Python | `train` |

_Display capped at 300 of 314 detected functions/methods._

### Automation pipeline files

- `.github/workflows/ci-cd.yml`
- `.github/workflows/ci.yml`
- `.github/workflows/deploy.yml`
- `.github/workflows/github-actions-demo.yml`
- `backend/api/regulator/transactions.php`
- `deploy.sh`
- `docker-compose.ci.yml`
- `docker-compose.yml`
- `docs/clean-installer-ultra-meta-2026.md`
- `frontend-admin/src/transaction-review.ts`
- `infra/kubernetes/api-deployment.yaml`
- `install_and_run.sh`
- `installer/artifacts/workflow-plan.txt`
- `installer/zgaming-ultra-installer.sh`
- `scripts/export-transaction-history.php`
- `services/training/pipeline.py`

## `cvsz/zSafeGuard`

Status: **cloned**; commit `94d388d96cc5`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `ai/loadtest/locustfile.py` | 17 | `/metrics` |
| `ai/main.py` | 49 | `/` |
| `ai/main.py` | 54 | `/health` |
| `ai/main.py` | 59 | `/metrics` |
| `ai/main.py` | 73 | `/events` |
| `ai/main.py` | 78 | `/report` |
| `ai/main.py` | 109 | `/analyze` |
| `ai/main.py` | 125 | `/stream` |
| `ai/tests/test_api.py` | 119 | `/metrics` |
| `dashboard/pages/index.js` | 21 | `/api/stats` |
| `dashboard/pages/index.js` | 59 | `/api/simulate?batch=8` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `ai/ensemble.py` | 12 | Python | `_load_models` |
| `ai/ensemble.py` | 18 | Python | `_fallback_predict` |
| `ai/ensemble.py` | 26 | Python | `predict` |
| `ai/loadtest/locustfile.py` | 8 | Python | `analyze` |
| `ai/loadtest/locustfile.py` | 16 | Python | `metrics` |
| `ai/loadtest/locustfile.py` | 20 | Python | `report` |
| `ai/loadtest/locustfile.py` | 24 | Python | `events` |
| `ai/main.py` | 50 | Python | `home` |
| `ai/main.py` | 55 | Python | `health` |
| `ai/main.py` | 60 | Python | `metrics` |
| `ai/main.py` | 74 | Python | `events` |
| `ai/main.py` | 79 | Python | `report` |
| `ai/main.py` | 110 | Python | `analyze` |
| `ai/main.py` | 126 | Python | `stream` |
| `ai/main.py` | 127 | Python | `event_generator` |
| `ai/schemas.py` | 35 | Python | `validate_feature_range` |
| `ai/tests/test_api.py` | 19 | Python | `setup_function` |
| `ai/tests/test_api.py` | 23 | Python | `_run_server` |
| `ai/tests/test_api.py` | 27 | Python | `_wait_for_server` |
| `ai/tests/test_api.py` | 37 | Python | `_request` |
| `ai/tests/test_api.py` | 55 | Python | `live_server` |
| `ai/tests/test_api.py` | 64 | Python | `test_analyze_request_model_validation` |
| `ai/tests/test_api.py` | 72 | Python | `test_analyze_success_and_event_capture` |
| `ai/tests/test_api.py` | 88 | Python | `test_validation_errors` |
| `ai/tests/test_api.py` | 107 | Python | `test_report_and_metrics` |
| `ai/tests/test_ensemble.py` | 5 | Python | `predict_proba` |
| `ai/tests/test_ensemble.py` | 11 | Python | `predict` |
| `ai/tests/test_ensemble.py` | 16 | Python | `test_predict_weighted_score` |
| `codex.sh` | 10 | Shell | `usage` |
| `codex.sh` | 14 | Shell | `require_cmd` |
| `codex.sh` | 23 | Shell | `run_step` |
| `codex.sh` | 34 | Shell | `main` |
| `dashboard/pages/api/simulate.js` | 3 | JS/TS | `randomFeatures` |
| `dashboard/pages/index.js` | 5 | JS/TS | `ScoreBadge` |
| `dashboard/pages/index.js` | 20 | JS/TS | `loadStats` |
| `dashboard/pages/index.js` | 57 | JS/TS | `runSimulation` |
| `deploy.sh` | 32 | Shell | `cleanup` |

### Automation pipeline files

- `.github/workflows/ci-cd.yml`
- `deploy.sh`
- `install.sh`
- `install_full.sh`
- `install_ultimate.sh`
- `k8s/deployment.yaml`

## `cvsz/zspin`

Status: **cloned**; commit `09550f99b3cc`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `examples/betting-platform-mvp/backend/src/modules/admin/admin.controller.ts` | 12 | `users` |
| `examples/betting-platform-mvp/backend/src/modules/betting/betting.controller.ts` | 27 | `:userId/place` |
| `examples/betting-platform-mvp/backend/src/modules/betting/betting.controller.ts` | 35 | `:betId/settle` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.controller.ts` | 15 | `:userId/deposit` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.controller.ts` | 23 | `:userId/withdraw` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.controller.ts` | 31 | `:userId/balance` |
| `examples/zspin-cloud-saas/backend/main.py` | 26 | `/login` |
| `examples/zspin-cloud-saas/backend/main.py` | 32 | `/query` |
| `examples/zspin-cloud-saas/backend/main.py` | 39 | `/analytics` |
| `examples/zspin-cloud-saas/backend/main.py` | 45 | `/billing/checkout` |
| `examples/zspin-cloud-saas/waitlist/server.py` | 12 | `/waitlist` |
| `src/zspin/apiserver.py` | 56 | `/apply` |
| `src/zspin/apiserver.py` | 68 | `/resources` |
| `src/zspin/control_plane/api.py` | 48 | `/deploy` |
| `src/zspin/control_plane/api.py` | 66 | `/tenants` |
| `src/zspin/control_plane/api.py` | 71 | `/clusters` |
| `src/zspin/control_plane/api.py` | 76 | `/billing` |
| `src/zspin/kv/api.py` | 42 | `/v1/kv` |
| `src/zspin/kv/api.py` | 47 | `/v1/kv` |
| `src/zspin/kv/api.py` | 52 | `/v1/kv` |
| `src/zspin/kv/api.py` | 57 | `/v1/kv/prefix` |
| `src/zspin/kv/api.py` | 62 | `/v1/kv/cas` |
| `src/zspin/kv/api.py` | 67 | `/v1/lease` |
| `src/zspin/kv/watch.py` | 28 | `/v1/watch/stream` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `examples/betting-platform-mvp/admin-panel/src/AdminDashboard.tsx` | 10 | JS/TS | `AdminDashboard` |
| `examples/betting-platform-mvp/backend/src/main.ts` | 5 | JS/TS | `bootstrap` |
| `examples/betting-platform-mvp/backend/src/modules/admin/admin.controller.ts` | 13 | JS/TS | `getUsers` |
| `examples/betting-platform-mvp/backend/src/modules/betting/betting.controller.ts` | 25 | JS/TS | `constructor` |
| `examples/betting-platform-mvp/backend/src/modules/betting/betting.service.ts` | 27 | JS/TS | `placeBet` |
| `examples/betting-platform-mvp/backend/src/modules/betting/betting.service.ts` | 45 | JS/TS | `settleBet` |
| `examples/betting-platform-mvp/backend/src/modules/realtime/realtime.gateway.ts` | 14 | JS/TS | `sendBalanceUpdate` |
| `examples/betting-platform-mvp/backend/src/modules/realtime/realtime.gateway.ts` | 18 | JS/TS | `sendBetUpdate` |
| `examples/betting-platform-mvp/backend/src/modules/users/users.service.ts` | 13 | JS/TS | `findAll` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.controller.ts` | 13 | JS/TS | `constructor` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.service.ts` | 14 | JS/TS | `constructor` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.service.ts` | 16 | JS/TS | `deposit` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.service.ts` | 23 | JS/TS | `withdraw` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.service.ts` | 30 | JS/TS | `getBalance` |
| `examples/betting-platform-mvp/backend/src/modules/wallet/wallet.service.ts` | 34 | JS/TS | `ensureWallet` |
| `examples/betting-platform-mvp/frontend/pages/dashboard.tsx` | 17 | JS/TS | `onBalance` |
| `examples/betting-platform-mvp/frontend/pages/dashboard.tsx` | 22 | JS/TS | `onBet` |
| `examples/enterprise-blueprint/mlops/dags/retrain_pipeline.py` | 9 | Python | `run_training` |
| `examples/enterprise-blueprint/mlops/dags/retrain_pipeline.py` | 13 | Python | `run_validation` |
| `examples/enterprise-blueprint/mlops/dags/retrain_pipeline.py` | 17 | Python | `deploy_canary` |
| `examples/enterprise-blueprint/mlops/mlflow_tracking_example.py` | 9 | Python | `train_example_model` |
| `examples/zspin-cloud-saas/backend/analytics.py` | 8 | Python | `get_stats` |
| `examples/zspin-cloud-saas/backend/auth.py` | 15 | Python | `create_token` |
| `examples/zspin-cloud-saas/backend/auth.py` | 25 | Python | `verify_token` |
| `examples/zspin-cloud-saas/backend/billing.py` | 12 | Python | `create_checkout` |
| `examples/zspin-cloud-saas/backend/main.py` | 15 | Python | `get_user` |
| `examples/zspin-cloud-saas/backend/main.py` | 27 | Python | `login` |
| `examples/zspin-cloud-saas/backend/main.py` | 33 | Python | `query` |
| `examples/zspin-cloud-saas/backend/main.py` | 40 | Python | `analytics` |
| `examples/zspin-cloud-saas/backend/main.py` | 46 | Python | `billing_checkout` |
| `examples/zspin-cloud-saas/backend/usage.py` | 10 | Python | `meter` |
| `examples/zspin-cloud-saas/frontend/src/Dashboard.js` | 6 | JS/TS | `runQuery` |
| `examples/zspin-cloud-saas/frontend/src/Login.js` | 6 | JS/TS | `handleLogin` |
| `examples/zspin-cloud-saas/waitlist/server.py` | 13 | Python | `join` |
| `scripts/deep_audit.py` | 23 | Python | `iter_python_files` |
| `scripts/deep_audit.py` | 32 | Python | `find_dangerous_patterns` |
| `scripts/deep_audit.py` | 90 | Python | `gather_inventory` |
| `scripts/deep_audit.py` | 104 | Python | `compliance_matrix` |
| `scripts/deep_audit.py` | 125 | Python | `compute_score` |
| `scripts/deep_audit.py` | 136 | Python | `build_report` |
| `scripts/deep_audit.py` | 175 | Python | `render_markdown` |
| `scripts/deep_audit.py` | 224 | Python | `main` |
| `scripts/validate.py` | 11 | Python | `run` |
| `scripts/validate.py` | 20 | Python | `main` |
| `src/zspin/aiops.py` | 5 | Python | `analyze` |
| `src/zspin/aiops.py` | 12 | Python | `feedback` |
| `src/zspin/apikeys.py` | 8 | Python | `generate_key` |
| `src/zspin/apikeys.py` | 14 | Python | `get_tenant` |
| `src/zspin/apiserver.py` | 14 | Python | `__init__` |
| `src/zspin/apiserver.py` | 17 | Python | `propose` |
| `src/zspin/apiserver.py` | 32 | Python | `init_app` |
| `src/zspin/apiserver.py` | 57 | Python | `apply` |
| `src/zspin/apiserver.py` | 69 | Python | `list_resources` |
| `src/zspin/audit.py` | 14 | Python | `write_audit_report` |
| `src/zspin/auth.py` | 13 | Python | `_b64url_encode` |
| `src/zspin/auth.py` | 17 | Python | `_b64url_decode` |
| `src/zspin/auth.py` | 22 | Python | `create_token` |
| `src/zspin/auth.py` | 39 | Python | `verify_token` |
| `src/zspin/billing.py` | 10 | Python | `calculate` |
| `src/zspin/cache.py` | 9 | Python | `get_cache` |
| `src/zspin/cli.py` | 14 | Python | `build_parser` |
| `src/zspin/cli.py` | 101 | Python | `_run_cost` |
| `src/zspin/cli.py` | 113 | Python | `main` |
| `src/zspin/cli.py` | 202 | Python | `reconcile` |
| `src/zspin/compliance.py` | 17 | Python | `_has_warn` |
| `src/zspin/compliance.py` | 21 | Python | `evaluate_controls` |
| `src/zspin/config.py` | 37 | Python | `_to_bool` |
| `src/zspin/config.py` | 51 | Python | `_normalize_required_binaries` |
| `src/zspin/config.py` | 62 | Python | `load_config` |
| `src/zspin/control_loop.py` | 8 | Python | `__init__` |
| `src/zspin/control_loop.py` | 20 | Python | `run` |
| `src/zspin/control_loop.py` | 29 | Python | `run_operator_loop` |
| `src/zspin/control_plane/api.py` | 18 | Python | `_resolve_identity` |
| `src/zspin/control_plane/api.py` | 39 | Python | `_apply_rate_limit` |
| `src/zspin/control_plane/api.py` | 49 | Python | `deploy` |
| `src/zspin/control_plane/api.py` | 67 | Python | `list_tenants` |
| `src/zspin/control_plane/api.py` | 72 | Python | `list_clusters` |
| `src/zspin/control_plane/api.py` | 77 | Python | `get_billing` |
| `src/zspin/control_plane/manager.py` | 10 | Python | `__init__` |
| `src/zspin/control_plane/manager.py` | 19 | Python | `register_tenant` |
| `src/zspin/control_plane/manager.py` | 22 | Python | `list_tenants` |
| `src/zspin/control_plane/manager.py` | 25 | Python | `list_clusters` |
| `src/zspin/control_plane/manager.py` | 28 | Python | `deploy` |
| `src/zspin/controller.py` | 7 | Python | `__init__` |
| `src/zspin/controller.py` | 11 | Python | `run` |
| `src/zspin/crd.py` | 5 | Python | `__init__` |
| `src/zspin/crd.py` | 12 | Python | `__init__` |
| `src/zspin/crd.py` | 15 | Python | `apply` |
| `src/zspin/crd.py` | 18 | Python | `list` |
| `src/zspin/crd.py` | 21 | Python | `get` |
| `src/zspin/db.py` | 9 | Python | `get_conn` |
| `src/zspin/diagnostics.py` | 21 | Python | `run_diagnostics` |
| `src/zspin/distributed_db/columnar.py` | 5 | Python | `__init__` |
| `src/zspin/distributed_db/columnar.py` | 8 | Python | `insert` |
| `src/zspin/distributed_db/columnar.py` | 12 | Python | `scan` |
| `src/zspin/distributed_db/engine.py` | 22 | Python | `__init__` |
| `src/zspin/distributed_db/engine.py` | 44 | Python | `query` |
| `src/zspin/distributed_db/engine.py` | 58 | Python | `_apply_mvcc_write` |
| `src/zspin/distributed_db/engine.py` | 66 | Python | `apply` |
| `src/zspin/distributed_db/engine.py` | 76 | Python | `snapshot_now` |
| `src/zspin/distributed_db/engine.py` | 79 | Python | `_load_snapshot` |
| `src/zspin/distributed_db/engine.py` | 91 | Python | `_replay_wal` |
| `src/zspin/distributed_db/geo.py` | 5 | Python | `__init__` |
| `src/zspin/distributed_db/geo.py` | 12 | Python | `get_replicas` |
| `src/zspin/distributed_db/index.py` | 5 | Python | `__init__` |
| `src/zspin/distributed_db/index.py` | 8 | Python | `add` |
| `src/zspin/distributed_db/index.py` | 11 | Python | `search` |
| `src/zspin/distributed_db/mvcc.py` | 7 | Python | `__init__` |
| `src/zspin/distributed_db/mvcc.py` | 11 | Python | `write` |
| `src/zspin/distributed_db/mvcc.py` | 16 | Python | `read` |
| `src/zspin/distributed_db/sharding.py` | 5 | Python | `__init__` |
| `src/zspin/distributed_db/sharding.py` | 11 | Python | `get_shard` |
| `src/zspin/distributed_db/sharding.py` | 18 | Python | `rebalance` |
| `src/zspin/distributed_db/snapshot.py` | 9 | Python | `__init__` |
| `src/zspin/distributed_db/snapshot.py` | 12 | Python | `save` |
| `src/zspin/distributed_db/snapshot.py` | 16 | Python | `load` |
| `src/zspin/distributed_db/truetime.py` | 7 | Python | `__init__` |
| `src/zspin/distributed_db/truetime.py` | 10 | Python | `now` |
| `src/zspin/distributed_db/truetime.py` | 14 | Python | `after` |
| `src/zspin/distributed_db/tx2pc.py` | 8 | Python | `__init__` |
| `src/zspin/distributed_db/tx2pc.py` | 12 | Python | `prepare` |
| `src/zspin/distributed_db/tx2pc.py` | 16 | Python | `commit` |
| `src/zspin/distributed_db/txn.py` | 9 | Python | `__init__` |
| `src/zspin/distributed_db/txn.py` | 14 | Python | `read` |
| `src/zspin/distributed_db/txn.py` | 17 | Python | `write` |
| `src/zspin/distributed_db/txn.py` | 20 | Python | `commit` |
| `src/zspin/distributed_db/txn.py` | 26 | Python | `rollback` |
| `src/zspin/distributed_db/vector.py` | 7 | Python | `__init__` |
| `src/zspin/distributed_db/vector.py` | 10 | Python | `add` |
| `src/zspin/distributed_db/vector.py` | 13 | Python | `search` |
| `src/zspin/distributed_db/vector.py` | 14 | Python | `dist` |
| `src/zspin/distributed_db/wal.py` | 10 | Python | `__init__` |
| `src/zspin/distributed_db/wal.py` | 14 | Python | `append` |
| `src/zspin/distributed_db/wal.py` | 20 | Python | `replay` |
| `src/zspin/distributed_loop.py` | 7 | Python | `__init__` |
| `src/zspin/distributed_loop.py` | 12 | Python | `run` |
| `src/zspin/etcd.py` | 7 | Python | `__init__` |
| `src/zspin/etcd.py` | 11 | Python | `put` |
| `src/zspin/etcd.py` | 16 | Python | `get` |
| `src/zspin/etcd.py` | 19 | Python | `history` |
| `src/zspin/events.py` | 12 | Python | `send` |
| `src/zspin/events.py` | 16 | Python | `_build_producer` |
| `src/zspin/events.py` | 29 | Python | `publish_event` |
| `src/zspin/finops.py` | 7 | Python | `analyze` |
| `src/zspin/finops.py` | 20 | Python | `estimate_cost` |
| `src/zspin/go_live.py` | 29 | Python | `_now_utc` |
| `src/zspin/go_live.py` | 33 | Python | `_write_json` |
| `src/zspin/go_live.py` | 39 | Python | `_copy_source_tree` |
| `src/zspin/go_live.py` | 73 | Python | `_write_linux_installer` |
| `src/zspin/go_live.py` | 98 | Python | `_write_windows_installer` |
| `src/zspin/go_live.py` | 119 | Python | `run_go_live_installer_bundle` |
| `src/zspin/installer.py` | 13 | Python | `_attempt_autoheal` |
| `src/zspin/installer.py` | 22 | Python | `run_workflow` |
| `src/zspin/intelligence.py` | 5 | Python | `decide` |
| `src/zspin/k8s/registry.py` | 1 | Python | `pod_key` |
| `src/zspin/k8s/registry.py` | 5 | Python | `service_key` |
| `src/zspin/kv/api.py` | 30 | Python | `init` |
| `src/zspin/kv/api.py` | 36 | Python | `_require_store` |
| `src/zspin/kv/api.py` | 43 | Python | `put` |
| `src/zspin/kv/api.py` | 48 | Python | `get` |
| `src/zspin/kv/api.py` | 53 | Python | `delete` |
| `src/zspin/kv/api.py` | 58 | Python | `prefix` |
| `src/zspin/kv/api.py` | 63 | Python | `cas` |
| `src/zspin/kv/api.py` | 68 | Python | `grant_lease` |
| `src/zspin/kv/lease.py` | 10 | Python | `__init__` |
| `src/zspin/kv/lease.py` | 18 | Python | `grant` |
| `src/zspin/kv/lease.py` | 24 | Python | `revoke` |
| `src/zspin/kv/lease.py` | 28 | Python | `_loop` |
| `src/zspin/kv/revision.py` | 9 | Python | `__init__` |
| `src/zspin/kv/revision.py` | 13 | Python | `next` |
| `src/zspin/kv/revision.py` | 19 | Python | `current` |
| `src/zspin/kv/store.py` | 15 | Python | `__init__` |
| `src/zspin/kv/store.py` | 22 | Python | `put` |
| `src/zspin/kv/store.py` | 37 | Python | `get` |
| `src/zspin/kv/store.py` | 40 | Python | `delete` |
| `src/zspin/kv/store.py` | 43 | Python | `prefix` |
| `src/zspin/kv/store.py` | 50 | Python | `cas` |
| `src/zspin/kv/store.py` | 56 | Python | `watch` |
| `src/zspin/kv/store.py` | 60 | Python | `_notify_watchers` |
| `src/zspin/kv/watch.py` | 17 | Python | `init` |
| `src/zspin/kv/watch.py` | 22 | Python | `_require_store` |
| `src/zspin/kv/watch.py` | 29 | Python | `watch_stream` |
| `src/zspin/kv/watch.py` | 33 | Python | `event_stream` |
| `src/zspin/leader.py` | 8 | Python | `__init__` |
| `src/zspin/leader.py` | 14 | Python | `try_acquire` |
| `src/zspin/leader.py` | 24 | Python | `renew` |
| `src/zspin/logging_utils.py` | 12 | Python | `format` |
| `src/zspin/logging_utils.py` | 25 | Python | `build_logger` |
| `src/zspin/main.py` | 14 | Python | `init_app` |
| `src/zspin/master_meta.py` | 26 | Python | `_deterministic_ai_risk_score` |
| `src/zspin/master_meta.py` | 34 | Python | `_run_release_packaging` |
| `src/zspin/master_meta.py` | 48 | Python | `_write_json` |
| `src/zspin/master_meta.py` | 54 | Python | `run_master_meta_bundle` |
| `src/zspin/metering.py` | 8 | Python | `record_deploy` |
| `src/zspin/metering.py` | 12 | Python | `record_cpu` |
| `src/zspin/metering.py` | 16 | Python | `get_usage` |
| `src/zspin/multicluster.py` | 5 | Python | `select_cluster` |
| `src/zspin/observability/alerts.py` | 10 | Python | `__init__` |
| `src/zspin/observability/alerts.py` | 16 | Python | `handle` |
| `src/zspin/observability/alerts.py` | 38 | Python | `map_alert` |
| `src/zspin/observability/metrics.py` | 8 | Python | `start_metrics_server` |
| `src/zspin/observability/metrics.py` | 12 | Python | `track_request` |
| `src/zspin/operator.py` | 7 | Python | `__init__` |
| `src/zspin/operator.py` | 12 | Python | `reconcile` |
| `src/zspin/platform/deploy.py` | 6 | Python | `deploy` |
| `src/zspin/platform/gitops.py` | 6 | Python | `apply_gitops` |
| `src/zspin/platform/runtime.py` | 9 | Python | `_run` |
| `src/zspin/platform/runtime.py` | 17 | Python | `deploy` |
| `src/zspin/platform/runtime.py` | 21 | Python | `scale` |
| `src/zspin/platform/runtime.py` | 31 | Python | `restart` |
| `src/zspin/platform/services.py` | 5 | Python | `__init__` |
| `src/zspin/platform/services.py` | 9 | Python | `to_dict` |
| `src/zspin/policy.py` | 8 | Python | `evaluate` |
| `src/zspin/raft/cluster.py` | 8 | Python | `handle_append_entries` |
| `src/zspin/raft/cluster.py` | 19 | Python | `install_snapshot` |
| `src/zspin/raft/cluster.py` | 28 | Python | `append_entries` |
| `src/zspin/raft/cluster.py` | 46 | Python | `install_snapshot` |
| `src/zspin/raft/cluster.py` | 53 | Python | `__init__` |
| `src/zspin/raft/cluster.py` | 56 | Python | `register` |
| `src/zspin/raft/cluster.py` | 59 | Python | `clients_for` |
| `src/zspin/raft/log.py` | 14 | Python | `__init__` |
| `src/zspin/raft/log.py` | 19 | Python | `append` |
| `src/zspin/raft/log.py` | 22 | Python | `match` |
| `src/zspin/raft/log.py` | 29 | Python | `delete_conflicts` |
| `src/zspin/raft/log.py` | 36 | Python | `commit_to` |
| `src/zspin/raft/log.py` | 39 | Python | `apply` |
| `src/zspin/raft/node.py` | 17 | Python | `__init__` |
| `src/zspin/raft/node.py` | 40 | Python | `attach_clients` |
| `src/zspin/raft/node.py` | 43 | Python | `become_leader` |
| `src/zspin/raft/node.py` | 52 | Python | `heartbeat_loop` |
| `src/zspin/raft/node.py` | 57 | Python | `become_follower` |
| `src/zspin/raft/node.py` | 65 | Python | `start_election` |
| `src/zspin/raft/node.py` | 101 | Python | `append_entry` |
| `src/zspin/raft/node.py` | 104 | Python | `propose` |
| `src/zspin/raft/node.py` | 115 | Python | `_replicate_over_http` |
| `src/zspin/raft/node.py` | 126 | Python | `send_heartbeat` |
| `src/zspin/raft/node.py` | 153 | Python | `update_commit_index` |
| `src/zspin/raft/node.py` | 166 | Python | `handle_append_entries` |
| `src/zspin/raft/node.py` | 202 | Python | `append_entries` |
| `src/zspin/raft/node.py` | 213 | Python | `request_vote` |
| `src/zspin/raft/node.py` | 248 | Python | `create_snapshot` |
| `src/zspin/raft/node.py` | 259 | Python | `take_snapshot` |
| `src/zspin/raft/node.py` | 262 | Python | `install_snapshot` |
| `src/zspin/raft/node.py` | 281 | Python | `install_snapshot_rpc` |
| `src/zspin/raft/node.py` | 289 | Python | `add_node` |
| `src/zspin/raft/node.py` | 301 | Python | `add_peer` |
| `src/zspin/raft/node.py` | 305 | Python | `remove_peer` |
| `src/zspin/raft/server.py` | 12 | Python | `__init__` |
| `src/zspin/raft/server.py` | 15 | Python | `RequestVote` |
| `src/zspin/raft/server.py` | 22 | Python | `AppendEntries` |
| `src/zspin/raft/server.py` | 28 | Python | `serve` |
| `src/zspin/raft/snapshot.py` | 10 | Python | `__init__` |
| `src/zspin/raft/snapshot.py` | 13 | Python | `save` |
| `src/zspin/raft/snapshot.py` | 21 | Python | `load` |
| `src/zspin/raft/state.py` | 10 | Python | `__init__` |
| `src/zspin/raft/state.py` | 16 | Python | `save` |
| `src/zspin/raft/state.py` | 20 | Python | `load` |
| `src/zspin/raft/state.py` | 28 | Python | `set_term` |
| `src/zspin/raft/state.py` | 33 | Python | `vote_for` |
| `src/zspin/raft/state_machine.py` | 7 | Python | `__init__` |
| `src/zspin/raft/state_machine.py` | 11 | Python | `register_handler` |
| `src/zspin/raft/state_machine.py` | 14 | Python | `apply` |
| `src/zspin/raft/storage.py` | 8 | Python | `__init__` |
| `src/zspin/raft/storage.py` | 11 | Python | `append` |
| `src/zspin/raft/storage.py` | 15 | Python | `read_all` |
| `src/zspin/rbac.py` | 10 | Python | `check_permission` |
| `src/zspin/remediation.py` | 12 | Python | `__init__` |
| `src/zspin/remediation.py` | 22 | Python | `run` |
| `src/zspin/remediation.py` | 30 | Python | `scale_service` |
| `src/zspin/remediation.py` | 43 | Python | `restart_service` |
| `src/zspin/rpc/client.py` | 8 | Python | `send` |
| `src/zspin/rpc/server.py` | 10 | Python | `do_POST` |
| `src/zspin/rpc/server.py` | 40 | Python | `log_message` |
| `src/zspin/rpc/server.py` | 45 | Python | `start_rpc` |
| `src/zspin/sbom.py` | 11 | Python | `generate_sbom` |
| `src/zspin/scaling.py` | 48 | Python | `__init__` |
| `src/zspin/scaling.py` | 52 | Python | `_validate` |
| `src/zspin/scaling.py` | 60 | Python | `_forecast_peak_rps` |
| `src/zspin/scaling.py` | 66 | Python | `_target_replicas` |
| `src/zspin/scaling.py` | 72 | Python | `_hpa_spec` |
| `src/zspin/scaling.py` | 105 | Python | `_canary_plan` |
| `src/zspin/scaling.py` | 122 | Python | `_multi_region_plan` |
| `src/zspin/scaling.py` | 138 | Python | `_diagnostics` |
| `src/zspin/scaling.py` | 151 | Python | `plan` |
| `src/zspin/scaling.py` | 179 | Python | `load_scaling_inputs` |
| `src/zspin/scaling.py` | 206 | Python | `generate_scaling_bundle` |
| `src/zspin/scheduler.py` | 7 | Python | `__init__` |
| `src/zspin/scheduler.py` | 10 | Python | `schedule` |
| `src/zspin/scheduler_scoring.py` | 5 | Python | `__init__` |
| `src/zspin/scheduler_scoring.py` | 11 | Python | `score` |
| `src/zspin/scheduler_scoring.py` | 14 | Python | `select` |
| `src/zspin/slo.py` | 5 | Python | `__init__` |
| `src/zspin/slo.py` | 11 | Python | `evaluate` |
| `src/zspin/sql/distributed_executor.py` | 8 | Python | `__init__` |
| `src/zspin/sql/distributed_executor.py` | 12 | Python | `execute` |
| `src/zspin/sql/executor.py` | 11 | Python | `__init__` |
| `src/zspin/sql/executor.py` | 16 | Python | `execute` |
| `src/zspin/sql/optimizer.py` | 5 | Python | `estimate_cost` |
| `src/zspin/sql/optimizer.py` | 12 | Python | `choose` |
| `src/zspin/sql/parser.py` | 5 | Python | `parse` |

_Display capped at 300 of 334 detected functions/methods._

### Automation pipeline files

- `.github/workflows/deploy.yaml`
- `docs/workflow.md`
- `examples/betting-platform-mvp/docker-compose.yml`
- `examples/enterprise-blueprint/kafka/docker-compose.kafka.yml`
- `examples/enterprise-blueprint/mlops/dags/retrain_pipeline.py`
- `examples/enterprise-blueprint/observability/docker-compose.observability.yml`
- `helm/zspin/templates/deployment.yaml`
- `k8s/deployment.yaml`
- `observability/prometheus-deploy.yaml`
- `src/zspin/installer.py`
- `src/zspin/platform/deploy.py`
- `tests/test_workflow.py`

## `ZeaZDev/vscode`

Status: **cloned**; commit `ea4ccd5f449b`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `scripts/playground-server.ts` | 16 | `/` |
| `scripts/playground-server.ts` | 37 | `/file-changes` |
| `src/vs/base/common/oauth.ts` | 835 | `/authorize` |
| `src/vs/base/test/common/labels.test.ts` | 180 | `/auth/c:/some/folder/file.txt` |
| `src/vs/base/test/common/labels.test.ts` | 181 | `/auth/c:/some/folder/file.txt` |
| `src/vs/base/test/common/uri.test.ts` | 76 | `/api/files/test.me` |
| `src/vs/base/test/common/uri.test.ts` | 108 | `/api/files/test.me` |
| `src/vs/base/test/common/uri.test.ts` | 109 | `/api/files/test.me` |
| `src/vs/base/test/common/uri.test.ts` | 110 | `/api/files/test.me` |
| `src/vs/base/test/common/uri.test.ts` | 111 | `/api/files/test.me` |
| `src/vs/base/test/common/uri.test.ts` | 112 | `/api/files/test.me` |
| `src/vs/base/test/common/uri.test.ts` | 113 | `/api/files/test.me` |
| `src/vs/base/test/common/uri.test.ts` | 138 | `/api/files/test.me` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `.eslint-plugin-local/code-amd-node-module.ts` | 21 | JS/TS | `create` |
| `.eslint-plugin-local/code-amd-node-module.ts` | 39 | JS/TS | `checkImport` |
| `.eslint-plugin-local/code-declare-service-brand.ts` | 16 | JS/TS | `create` |
| `.eslint-plugin-local/code-ensure-no-disposables-leak-in-test.ts` | 20 | JS/TS | `create` |
| `.eslint-plugin-local/code-import-patterns.ts` | 33 | JS/TS | `isLayerAllowRule` |
| `.eslint-plugin-local/code-import-patterns.ts` | 57 | JS/TS | `create` |
| `.eslint-plugin-local/code-import-patterns.ts` | 78 | JS/TS | `_processOptions` |
| `.eslint-plugin-local/code-import-patterns.ts` | 93 | JS/TS | `orSegment` |
| `.eslint-plugin-local/code-import-patterns.ts` | 125 | JS/TS | `findLayer` |
| `.eslint-plugin-local/code-import-patterns.ts` | 134 | JS/TS | `generateConfig` |
| `.eslint-plugin-local/code-import-patterns.ts` | 227 | JS/TS | `_checkImport` |
| `.eslint-plugin-local/code-import-patterns.ts` | 284 | JS/TS | `getRelativeFilename` |
| `.eslint-plugin-local/code-layering.ts` | 37 | JS/TS | `create` |
| `.eslint-plugin-local/code-limited-top-functions.ts` | 31 | JS/TS | `create` |
| `.eslint-plugin-local/code-limited-top-functions.ts` | 57 | JS/TS | `ExportNamedDeclaration` |
| `.eslint-plugin-local/code-must-use-result.ts` | 20 | JS/TS | `create` |
| `.eslint-plugin-local/code-must-use-super-dispose.ts` | 12 | JS/TS | `create` |
| `.eslint-plugin-local/code-must-use-super-dispose.ts` | 13 | JS/TS | `doesCallSuperDispose` |
| `.eslint-plugin-local/code-no-any-casts.ts` | 11 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-dangerous-type-assertions.ts` | 12 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-deep-import-of-internal.ts` | 29 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-deep-import-of-internal.ts` | 64 | JS/TS | `splitParts` |
| `.eslint-plugin-local/code-no-global-document-listener.ts` | 10 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-global-document-listener.ts` | 12 | JS/TS | `CallExpression` |
| `.eslint-plugin-local/code-no-in-operator.ts` | 29 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-in-operator.ts` | 31 | JS/TS | `checkInOperator` |
| `.eslint-plugin-local/code-no-localization-template-literals.ts` | 34 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-localization-template-literals.ts` | 36 | JS/TS | `checkCallExpression` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 27 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 28 | JS/TS | `reportIfLocalized` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 53 | JS/TS | `isModelDescriptionKey` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 63 | JS/TS | `isModelDescriptionAssignmentTarget` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 70 | JS/TS | `containsLocalizedCall` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 73 | JS/TS | `visit` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 104 | JS/TS | `isLocalizeCall` |
| `.eslint-plugin-local/code-no-localized-model-description.ts` | 115 | JS/TS | `isLocalizeCallee` |
| `.eslint-plugin-local/code-no-native-private.ts` | 18 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-nls-in-standalone-editor.ts` | 19 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-observable-get-in-reactive-context.ts` | 20 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-observable-get-in-reactive-context.ts` | 48 | JS/TS | `checkFunctionForObservableGetCalls` |
| `.eslint-plugin-local/code-no-observable-get-in-reactive-context.ts` | 55 | JS/TS | `traverse` |
| `.eslint-plugin-local/code-no-observable-get-in-reactive-context.ts` | 81 | JS/TS | `isObservableGetCall` |
| `.eslint-plugin-local/code-no-observable-get-in-reactive-context.ts` | 111 | JS/TS | `getReaderParameterName` |
| `.eslint-plugin-local/code-no-observable-get-in-reactive-context.ts` | 124 | JS/TS | `isReactiveFunctionWithReader` |
| `.eslint-plugin-local/code-no-observable-get-in-reactive-context.ts` | 131 | JS/TS | `walkChildren` |
| `.eslint-plugin-local/code-no-potentially-unsafe-disposables.ts` | 16 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-potentially-unsafe-disposables.ts` | 17 | JS/TS | `checkVariableDeclaration` |
| `.eslint-plugin-local/code-no-potentially-unsafe-disposables.ts` | 24 | JS/TS | `checkProperty` |
| `.eslint-plugin-local/code-no-reader-after-await.ts` | 11 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-reader-after-await.ts` | 39 | JS/TS | `checkFunctionForAwaitBeforeReader` |
| `.eslint-plugin-local/code-no-reader-after-await.ts` | 47 | JS/TS | `collectPositions` |
| `.eslint-plugin-local/code-no-reader-after-await.ts` | 121 | JS/TS | `getMethodName` |
| `.eslint-plugin-local/code-no-reader-after-await.ts` | 129 | JS/TS | `isReaderMethodCall` |
| `.eslint-plugin-local/code-no-reader-after-await.ts` | 152 | JS/TS | `getReaderParameterName` |
| `.eslint-plugin-local/code-no-reader-after-await.ts` | 163 | JS/TS | `isFunctionWithReader` |
| `.eslint-plugin-local/code-no-runtime-import.ts` | 32 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-standalone-editor.ts` | 22 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-static-self-ref.ts` | 15 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-static-self-ref.ts` | 17 | JS/TS | `checkProperty` |
| `.eslint-plugin-local/code-no-test-async-suite.ts` | 10 | JS/TS | `isCallExpression` |
| `.eslint-plugin-local/code-no-test-async-suite.ts` | 14 | JS/TS | `isFunctionExpression` |
| `.eslint-plugin-local/code-no-test-async-suite.ts` | 20 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-test-async-suite.ts` | 21 | JS/TS | `hasAsyncSuite` |
| `.eslint-plugin-local/code-no-test-only.ts` | 11 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-unexternalized-strings.ts` | 10 | JS/TS | `isStringLiteral` |
| `.eslint-plugin-local/code-no-unexternalized-strings.ts` | 14 | JS/TS | `isDoubleQuoted` |
| `.eslint-plugin-local/code-no-unexternalized-strings.ts` | 42 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-unexternalized-strings.ts` | 47 | JS/TS | `collectDoubleQuotedStrings` |
| `.eslint-plugin-local/code-no-unexternalized-strings.ts` | 53 | JS/TS | `visitLocalizeCall` |
| `.eslint-plugin-local/code-no-unexternalized-strings.ts` | 100 | JS/TS | `visitL10NCall` |
| `.eslint-plugin-local/code-no-unexternalized-strings.ts` | 119 | JS/TS | `reportBadStringsAndBadKeys` |
| `.eslint-plugin-local/code-no-unused-expressions.ts` | 55 | JS/TS | `create` |
| `.eslint-plugin-local/code-no-unused-expressions.ts` | 66 | JS/TS | `looksLikeDirective` |
| `.eslint-plugin-local/code-no-unused-expressions.ts` | 91 | JS/TS | `directives` |
| `.eslint-plugin-local/code-no-unused-expressions.ts` | 101 | JS/TS | `isDirective` |
| `.eslint-plugin-local/code-no-unused-expressions.ts` | 115 | JS/TS | `isValidExpression` |
| `.eslint-plugin-local/code-no-unused-expressions.ts` | 143 | JS/TS | `ExpressionStatement` |
| `.eslint-plugin-local/code-parameter-properties-must-have-explicit-accessibility.ts` | 16 | JS/TS | `create` |
| `.eslint-plugin-local/code-parameter-properties-must-have-explicit-accessibility.ts` | 17 | JS/TS | `check` |
| `.eslint-plugin-local/code-policy-localization-key-match.ts` | 37 | JS/TS | `create` |
| `.eslint-plugin-local/code-policy-localization-key-match.ts` | 39 | JS/TS | `checkLocalizationObject` |
| `.eslint-plugin-local/code-policy-localization-key-match.ts` | 117 | JS/TS | `isInPolicyBlock` |
| `.eslint-plugin-local/code-translation-remind.ts` | 23 | JS/TS | `create` |
| `.eslint-plugin-local/code-translation-remind.ts` | 27 | JS/TS | `_checkImport` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 10 | JS/TS | `testValidUsage` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 30 | JS/TS | `testInvalidUsage` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 65 | JS/TS | `testComplexCases` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 83 | JS/TS | `process` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 108 | JS/TS | `testValidComplexCases` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 125 | JS/TS | `processValues` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 139 | JS/TS | `testEdgeCases` |
| `.eslint-plugin-local/tests/code-no-observable-get-in-reactive-context-test.ts` | 165 | JS/TS | `testQuickFixScenarios` |
| `.eslint-plugin-local/tests/code-no-reader-after-await-test.ts` | 10 | JS/TS | `testValidUsage` |
| `.eslint-plugin-local/tests/code-no-reader-after-await-test.ts` | 24 | JS/TS | `testInvalidUsage` |
| `.eslint-plugin-local/tests/code-no-reader-after-await-test.ts` | 49 | JS/TS | `testComplexCases` |
| `.eslint-plugin-local/tests/code-no-reader-after-await-test.ts` | 76 | JS/TS | `testValidComplexCases` |
| `.eslint-plugin-local/utils.ts` | 10 | JS/TS | `createImportRuleListener` |
| `.eslint-plugin-local/utils.ts` | 12 | JS/TS | `_checkImport` |
| `.eslint-plugin-local/vscode-dts-cancellation.ts` | 18 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-create-func.ts` | 18 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-event-naming.ts` | 27 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-event-naming.ts` | 86 | JS/TS | `getIdent` |
| `.eslint-plugin-local/vscode-dts-interface-naming.ts` | 21 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-literal-or-types.ts` | 17 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-provider-naming.ts` | 20 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-string-type-literals.ts` | 20 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-use-export.ts` | 19 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-use-thenable.ts` | 18 | JS/TS | `create` |
| `.eslint-plugin-local/vscode-dts-vscode-in-comments.ts` | 18 | JS/TS | `create` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 10 | JS/TS | `activate` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 20 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 29 | JS/TS | `dispose` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 37 | JS/TS | `all` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 48 | JS/TS | `_refresh` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 63 | JS/TS | `findNodeAtPosition` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 70 | JS/TS | `findImportAt` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 80 | JS/TS | `provideCompletionItems` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 139 | JS/TS | `provideCodeActions` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 148 | JS/TS | `_provideFixAll` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 192 | JS/TS | `_provideFix` |
| `.vscode/extensions/vscode-selfhost-import-aid/src/extension.ts` | 202 | JS/TS | `_provideFixesForDiag` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 21 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 23 | JS/TS | `add` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 41 | JS/TS | `report` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 62 | JS/TS | `add` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 71 | JS/TS | `report` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 81 | JS/TS | `add` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 157 | JS/TS | `add` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/coverageProvider.ts` | 165 | JS/TS | `testDetails` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/debounce.ts` | 9 | JS/TS | `debounce` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/debounce.ts` | 11 | JS/TS | `debounced` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 27 | JS/TS | `getWorkspaceFolderForTestFile` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 41 | JS/TS | `activate` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 46 | JS/TS | `provideFollowup` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 56 | JS/TS | `resolveHandler` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 249 | JS/TS | `updateNodeForDocument` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 278 | JS/TS | `deactivate` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 282 | JS/TS | `getOrCreateFile` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 305 | JS/TS | `gatherTestItems` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 311 | JS/TS | `startWatchingWorkspace` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/extension.ts` | 341 | JS/TS | `getPendingTestMap` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 32 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 76 | JS/TS | `dispose` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 87 | JS/TS | `formatJsonValue` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 125 | JS/TS | `detectFailingDeepStrictEqualAssertion` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 161 | JS/TS | `fromNode` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 177 | JS/TS | `atPosition` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 191 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 204 | JS/TS | `findNodeAt` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 213 | JS/TS | `getAllTestStatusMessagesAt` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 237 | JS/TS | `getTestResultsWithUri` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failingDeepStrictEqualAssertFixer.ts` | 240 | JS/TS | `walk` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failureTracker.ts` | 36 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failureTracker.ts` | 64 | JS/TS | `getGitState` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failureTracker.ts` | 95 | JS/TS | `append` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failureTracker.ts` | 113 | JS/TS | `captureGitState` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failureTracker.ts` | 133 | JS/TS | `dispose` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/failureTracker.ts` | 137 | JS/TS | `exec` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 40 | JS/TS | `provideRelatedCode` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 73 | JS/TS | `provideRelatedTests` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 115 | JS/TS | `didChange` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 132 | JS/TS | `getNode` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 138 | JS/TS | `discoverOutwards` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 154 | JS/TS | `discoverOutwardsInner` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 172 | JS/TS | `syncNode` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 209 | JS/TS | `uriToImportPath` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/importGraph.ts` | 222 | JS/TS | `importPathToUri` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/metadata.ts` | 16 | JS/TS | `getId` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/metadata.ts` | 22 | JS/TS | `attachTestMessageMetadata` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/metadata.ts` | 46 | JS/TS | `getTestMessageMetadata` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/snapshot.ts` | 11 | JS/TS | `registerSnapshotUpdate` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/sourceUtils.ts` | 18 | JS/TS | `extractTestFromNode` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/stackTraceParser.ts` | 23 | JS/TS | `isStackLike` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/stackTraceParser.ts` | 26 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/streamSplitter.ts` | 21 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 109 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 121 | JS/TS | `dispose` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 152 | JS/TS | `scanTestOutput` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 163 | JS/TS | `enqueueOutput` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 382 | JS/TS | `forceCRLF` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 384 | JS/TS | `sourcemapStack` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 411 | JS/TS | `outputToString` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 414 | JS/TS | `tryMakeMarkdown` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 439 | JS/TS | `getSourceLocationMapper` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 491 | JS/TS | `getSourceLocation` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 519 | JS/TS | `getSourceFile` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 539 | JS/TS | `completeSourceMapUrl` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 551 | JS/TS | `loadSourceMap` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 557 | JS/TS | `promise` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 580 | JS/TS | `replaceAllLocations` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 613 | JS/TS | `deriveStackLocations` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testOutputScanner.ts` | 647 | JS/TS | `deriveSourceLocation` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testTree.ts` | 19 | JS/TS | `clearFileDiagnostics` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testTree.ts` | 24 | JS/TS | `guessWorkspaceFolder` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testTree.ts` | 63 | JS/TS | `getId` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testTree.ts` | 67 | JS/TS | `getLabel` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testTree.ts` | 71 | JS/TS | `updateFromDisk` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/testTree.ts` | 102 | JS/TS | `traverse` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 43 | JS/TS | `initializeBlocks` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 77 | JS/TS | `clone` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 84 | JS/TS | `cover` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 89 | JS/TS | `uncovered` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 104 | JS/TS | `setCovered` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 111 | JS/TS | `push` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 171 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 179 | JS/TS | `getLineLength` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 189 | JS/TS | `getLineOfOffset` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/v8CoverageWrangling.ts` | 207 | JS/TS | `toLineColumn` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 18 | JS/TS | `escapeRe` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 27 | JS/TS | `constructor` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 29 | JS/TS | `run` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 40 | JS/TS | `debug` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 72 | JS/TS | `createDebugAdapterTracker` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 80 | JS/TS | `onDidSendMessage` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 85 | JS/TS | `onWillReceiveMessage` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 126 | JS/TS | `findOpenPort` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 142 | JS/TS | `getEnvironment` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 161 | JS/TS | `addTestFileRunPath` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 172 | JS/TS | `addTestCaseOrSuite` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 218 | JS/TS | `readProductJson` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 230 | JS/TS | `createWaitServer` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 259 | JS/TS | `binaryPath` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 273 | JS/TS | `getDefaultArgs` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 280 | JS/TS | `binaryPath` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 286 | JS/TS | `getDefaultArgs` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 293 | JS/TS | `binaryPath` |
| `.vscode/extensions/vscode-selfhost-test-provider/src/vscodeTestRunner.ts` | 299 | JS/TS | `getDefaultArgs` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 12 | JS/TS | `activate` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 32 | JS/TS | `registerSettingsCompletions` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 34 | JS/TS | `provideCompletionItems` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 40 | JS/TS | `registerVariableCompletions` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 42 | JS/TS | `provideCompletionItems` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 85 | JS/TS | `isCompletingInsidePropertyStringValue` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 97 | JS/TS | `isLocationInsideTopLevelProperty` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 105 | JS/TS | `registerExtensionsCompletions` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 109 | JS/TS | `registerExtensionsCompletionsInExtensionsDocument` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 111 | JS/TS | `provideCompletionItems` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 123 | JS/TS | `registerExtensionsCompletionsInWorkspaceConfigurationDocument` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 125 | JS/TS | `provideCompletionItems` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 137 | JS/TS | `getReplaceRange` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 149 | JS/TS | `provideDocumentSymbols` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 183 | JS/TS | `registerContextKeyCompletions` |
| `extensions/configuration-editing/src/configurationEditingMain.ts` | 202 | JS/TS | `provideCompletionItems` |
| `extensions/configuration-editing/src/extensionsProposals.ts` | 9 | JS/TS | `provideInstalledExtensionProposals` |
| `extensions/configuration-editing/src/importExportProfiles.ts` | 17 | JS/TS | `getOctokit` |
| `extensions/configuration-editing/src/importExportProfiles.ts` | 35 | JS/TS | `saveProfile` |
| `extensions/configuration-editing/src/importExportProfiles.ts` | 53 | JS/TS | `getPublicOctokit` |
| `extensions/configuration-editing/src/importExportProfiles.ts` | 63 | JS/TS | `readProfile` |
| `extensions/configuration-editing/src/importExportProfiles.ts` | 64 | JS/TS | `readProfile` |
| `extensions/configuration-editing/src/importExportProfiles.ts` | 65 | JS/TS | `readProfile` |
| `extensions/configuration-editing/src/node/net.ts` | 17 | JS/TS | `getAgent` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 14 | JS/TS | `constructor` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 16 | JS/TS | `provideCompletionItems` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 72 | JS/TS | `getReplaceRange` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 83 | JS/TS | `isCompletingPropertyValue` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 95 | JS/TS | `provideWindowTitleCompletionItems` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 107 | JS/TS | `getText` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 134 | JS/TS | `provideEditorLabelCompletionItems` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 146 | JS/TS | `getText` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 159 | JS/TS | `provideFilesAssociationsCompletionItems` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 189 | JS/TS | `provideExcludeCompletionItems` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 253 | JS/TS | `provideLanguageCompletionItems` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 265 | JS/TS | `provideLanguageCompletionItemsForLanguageOverrides` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 277 | JS/TS | `provideLanguageOverridesCompletionItems` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 321 | JS/TS | `providePortsAttributesCompletionItem` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 346 | JS/TS | `newSimpleCompletionItem` |
| `extensions/configuration-editing/src/settingsDocumentHelper.ts` | 355 | JS/TS | `newSnippetCompletionItem` |
| `extensions/configuration-editing/src/test/completion.test.ts` | 545 | JS/TS | `testCompletion` |
| `extensions/configuration-editing/src/test/completion.test.ts` | 582 | JS/TS | `setTestContent` |
| `extensions/css-language-features/client/src/browser/cssClientMain.ts` | 15 | JS/TS | `activate` |
| `extensions/css-language-features/client/src/browser/cssClientMain.ts` | 33 | JS/TS | `deactivate` |
| `extensions/css-language-features/client/src/cssClient.ts` | 39 | JS/TS | `startClient` |
| `extensions/css-language-features/client/src/cssClient.ts` | 61 | JS/TS | `provideCompletionItem` |
| `extensions/css-language-features/client/src/cssClient.ts` | 63 | JS/TS | `updateRanges` |
| `extensions/css-language-features/client/src/cssClient.ts` | 70 | JS/TS | `updateLabel` |
| `extensions/css-language-features/client/src/cssClient.ts` | 79 | JS/TS | `updateProposals` |
| `extensions/css-language-features/client/src/cssClient.ts` | 122 | JS/TS | `initCompletionProvider` |
| `extensions/css-language-features/client/src/cssClient.ts` | 126 | JS/TS | `provideCompletionItems` |
| `extensions/css-language-features/client/src/cssClient.ts` | 152 | JS/TS | `applyCodeAction` |
| `extensions/css-language-features/client/src/cssClient.ts` | 170 | JS/TS | `updateFormatterRegistration` |
| `extensions/css-language-features/client/src/cssClient.ts` | 177 | JS/TS | `provideDocumentRangeFormattingEdits` |
| `extensions/css-language-features/client/src/customData.ts` | 9 | JS/TS | `getCustomDataSource` |
| `extensions/css-language-features/client/src/customData.ts` | 40 | JS/TS | `getCustomDataPathsInAllWorkspaces` |
| `extensions/css-language-features/client/src/customData.ts` | 49 | JS/TS | `collect` |
| `extensions/css-language-features/client/src/customData.ts` | 77 | JS/TS | `getCustomDataPathsFromAllExtensions` |
| `extensions/css-language-features/client/src/dropOrPaste/dropOrPasteResource.ts` | 64 | JS/TS | `getUriList` |
| `extensions/css-language-features/client/src/dropOrPaste/dropOrPasteResource.ts` | 82 | JS/TS | `createUriListSnippet` |
| `extensions/css-language-features/client/src/dropOrPaste/dropOrPasteResource.ts` | 107 | JS/TS | `pasteAsCssUrlByDefault` |
| `extensions/css-language-features/client/src/dropOrPaste/dropOrPasteResource.ts` | 118 | JS/TS | `getRelativePath` |
| `extensions/css-language-features/client/src/dropOrPaste/dropOrPasteResource.ts` | 134 | JS/TS | `registerDropOrPasteResourceSupport` |
| `extensions/css-language-features/client/src/dropOrPaste/shared.ts` | 21 | JS/TS | `getDocumentDir` |
| `extensions/css-language-features/client/src/dropOrPaste/shared.ts` | 29 | JS/TS | `getParentDocumentUri` |
| `extensions/css-language-features/client/src/dropOrPaste/uriList.ts` | 8 | JS/TS | `splitUriList` |
| `extensions/css-language-features/client/src/dropOrPaste/uriList.ts` | 12 | JS/TS | `parseUriList` |
| `extensions/css-language-features/client/src/dropOrPaste/uriList.ts` | 20 | JS/TS | `from` |
| `extensions/css-language-features/client/src/node/cssClientMain.ts` | 15 | JS/TS | `activate` |
| `extensions/css-language-features/client/src/node/cssClientMain.ts` | 43 | JS/TS | `deactivate` |
| `extensions/css-language-features/client/src/node/nodeFs.ts` | 10 | JS/TS | `getNodeFSRequestService` |
| `extensions/css-language-features/client/src/node/nodeFs.ts` | 11 | JS/TS | `ensureFileUri` |
| `extensions/css-language-features/client/src/node/nodeFs.ts` | 17 | JS/TS | `getContent` |

_Display capped at 300 of 64674 detected functions/methods._

### Automation pipeline files

- `.config/1espt/PipelineAutobaseliningConfig.yml`
- `.devcontainer/install-vscode.sh`
- `.github/workflows/check-clean-git-state.sh`
- `.github/workflows/copilot-setup-steps.yml`
- `.github/workflows/monaco-editor.yml`
- `.github/workflows/no-package-lock-changes.yml`
- `.github/workflows/no-yarn-lock-changes.yml`
- `.github/workflows/pr-darwin-test.yml`
- `.github/workflows/pr-linux-cli-test.yml`
- `.github/workflows/pr-linux-test.yml`
- `.github/workflows/pr-node-modules.yml`
- `.github/workflows/pr-win32-test.yml`
- `.github/workflows/pr.yml`
- `.github/workflows/telemetry.yml`
- `extensions/emmet/src/abbreviationActions.ts`
- `extensions/emmet/src/test/abbreviationAction.test.ts`
- `extensions/emmet/src/test/cssAbbreviationAction.test.ts`
- `extensions/emmet/src/test/tagActions.test.ts`
- `extensions/git/src/actionButton.ts`
- `extensions/postinstall.mjs`
- `extensions/terminal-suggest/src/completions/upstream/docker-compose.ts`
- `extensions/typescript-language-features/resources/walkthroughs/install-node-js.svg`
- `extensions/typescript-language-features/src/filesystems/autoInstallerFs.ts`
- `extensions/typescript-language-features/src/languageFeatures/util/codeAction.ts`
- `extensions/typescript-language-features/web/src/typingsInstaller/typingsInstaller.ts`
- `src/bootstrap-cli.ts`
- `src/bootstrap-esm.ts`
- `src/bootstrap-fork.ts`
- `src/bootstrap-import.ts`
- `src/bootstrap-meta.ts`
- `src/bootstrap-node.ts`
- `src/bootstrap-server.ts`
- `src/vs/base/browser/ui/actionbar/actionViewItems.ts`
- `src/vs/base/browser/ui/actionbar/actionbar.css`
- `src/vs/base/browser/ui/actionbar/actionbar.ts`
- `src/vs/base/browser/ui/dropdown/dropdownActionViewItem.ts`
- `src/vs/base/common/actions.ts`
- `src/vs/base/common/observableInternal/transaction.ts`
- `src/vs/base/common/worker/webWorkerBootstrap.ts`
- `src/vs/base/test/browser/actionbar.test.ts`
- `src/vs/editor/common/editorAction.ts`
- `src/vs/editor/common/languages/enterAction.ts`
- `src/vs/editor/contrib/codeAction/browser/codeAction.ts`
- `src/vs/editor/contrib/codeAction/browser/codeActionCommands.ts`
- `src/vs/editor/contrib/codeAction/browser/codeActionContributions.ts`
- `src/vs/editor/contrib/codeAction/browser/codeActionController.ts`
- `src/vs/editor/contrib/codeAction/browser/codeActionKeybindingResolver.ts`
- `src/vs/editor/contrib/codeAction/browser/codeActionMenu.ts`
- `src/vs/editor/contrib/codeAction/browser/codeActionModel.ts`
- `src/vs/editor/contrib/codeAction/test/browser/codeAction.test.ts`
- `src/vs/editor/contrib/codeAction/test/browser/codeActionKeybindingResolver.test.ts`
- `src/vs/editor/contrib/codeAction/test/browser/codeActionModel.test.ts`
- `src/vs/editor/contrib/colorPicker/browser/standaloneColorPicker/standaloneColorPickerActions.ts`
- `src/vs/editor/contrib/format/browser/formatActions.ts`
- `src/vs/editor/contrib/gpu/browser/gpuActions.ts`
- `src/vs/editor/contrib/hover/browser/hoverActionIds.ts`
- `src/vs/editor/contrib/hover/browser/hoverActions.ts`
- `src/vs/editor/contrib/stickyScroll/browser/stickyScrollActions.ts`
- `src/vs/platform/accessibilitySignal/browser/media/chatUserActionRequired.mp3`
- `src/vs/platform/accessibilitySignal/browser/media/codeActionApplied.mp3`
- `src/vs/platform/accessibilitySignal/browser/media/codeActionTriggered.mp3`
- `src/vs/platform/action/common/action.ts`
- `src/vs/platform/action/common/actionCommonCategories.ts`
- `src/vs/platform/actionWidget/browser/actionList.ts`
- `src/vs/platform/actionWidget/browser/actionWidget.css`
- `src/vs/platform/actionWidget/browser/actionWidget.ts`
- `src/vs/platform/actionWidget/browser/actionWidgetDropdown.ts`
- `src/vs/platform/actionWidget/common/actionWidget.ts`
- `src/vs/platform/actions/browser/actionViewItemService.ts`
- `src/vs/platform/actions/browser/actionWidgetDropdownActionViewItem.ts`
- `src/vs/platform/actions/browser/dropdownActionViewItemWithKeybinding.ts`
- `src/vs/platform/actions/browser/dropdownWithPrimaryActionViewItem.ts`
- `src/vs/platform/actions/browser/menuEntryActionViewItem.css`
- `src/vs/platform/actions/browser/menuEntryActionViewItem.ts`
- `src/vs/platform/actions/common/actions.contribution.ts`
- `src/vs/platform/actions/common/actions.ts`
- `src/vs/platform/actions/common/menuResetAction.ts`
- `src/vs/platform/quickinput/browser/quickInputActions.ts`
- `src/vs/workbench/api/common/extHostTypes/codeActionKind.ts`
- `src/vs/workbench/browser/actions/developerActions.ts`
- `src/vs/workbench/browser/actions/helpActions.ts`
- `src/vs/workbench/browser/actions/layoutActions.ts`
- `src/vs/workbench/browser/actions/media/actions.css`
- `src/vs/workbench/browser/actions/navigationActions.ts`
- `src/vs/workbench/browser/actions/quickAccessActions.ts`
- `src/vs/workbench/browser/actions/textInputActions.ts`
- `src/vs/workbench/browser/actions/windowActions.ts`
- `src/vs/workbench/browser/actions/workspaceActions.ts`
- `src/vs/workbench/browser/parts/activitybar/media/activityaction.css`
- `src/vs/workbench/browser/parts/auxiliarybar/auxiliaryBarActions.ts`
- `src/vs/workbench/browser/parts/compositeBarActions.ts`
- `src/vs/workbench/browser/parts/editor/editorActions.ts`
- `src/vs/workbench/browser/parts/notifications/media/notificationsActions.css`
- `src/vs/workbench/browser/parts/notifications/notificationsActions.ts`
- `src/vs/workbench/browser/parts/panel/panelActions.ts`
- `src/vs/workbench/browser/parts/sidebar/sidebarActions.ts`
- `src/vs/workbench/browser/parts/statusbar/statusbarActions.ts`
- `src/vs/workbench/browser/parts/titlebar/titlebarActions.ts`
- `src/vs/workbench/browser/parts/views/viewMenuActions.ts`
- `src/vs/workbench/contrib/accessibility/browser/accessibleViewActions.ts`
- `src/vs/workbench/contrib/authentication/browser/actions/manageAccountPreferencesForExtensionAction.ts`
- `src/vs/workbench/contrib/authentication/browser/actions/manageAccountPreferencesForMcpServerAction.ts`
- `src/vs/workbench/contrib/authentication/browser/actions/manageAccountsAction.ts`
- `src/vs/workbench/contrib/authentication/browser/actions/manageDynamicAuthenticationProvidersAction.ts`
- `src/vs/workbench/contrib/authentication/browser/actions/manageTrustedExtensionsForAccountAction.ts`
- `src/vs/workbench/contrib/authentication/browser/actions/manageTrustedMcpServersForAccountAction.ts`
- `src/vs/workbench/contrib/authentication/browser/actions/signOutOfAccountAction.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatAccessibilityActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatCodeblockActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatContextActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatContinueInAction.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatCopyActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatDeveloperActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatElicitationActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatExecuteActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatFileTreeActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatLanguageModelActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatMoveActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatNewActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatPromptNavigationActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatQuickInputActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatSessionActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatTitleActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/chatToolActions.ts`
- `src/vs/workbench/contrib/chat/browser/actions/manageModelsActions.ts`
- `src/vs/workbench/contrib/chat/browser/agentSessions/agentSessionsActions.ts`
- `src/vs/workbench/contrib/chat/browser/agentSessions/media/agentsessionsactions.css`
- `src/vs/workbench/contrib/chat/browser/chatContentParts/chatMcpServersInteractionContentPart.ts`
- `src/vs/workbench/contrib/chat/browser/chatContentParts/media/chatMcpServersInteractionContent.css`
- `src/vs/workbench/contrib/chat/browser/chatContentParts/toolInvocationParts/chatExtensionsInstallToolSubPart.ts`
- `src/vs/workbench/contrib/chat/browser/chatEditing/chatEditingActions.ts`
- `src/vs/workbench/contrib/chat/browser/chatEditing/chatEditingEditorActions.ts`
- `src/vs/workbench/contrib/chat/browser/chatSessions/chatSessionPickerActionItem.ts`
- `src/vs/workbench/contrib/chat/browser/modelPicker/modePickerActionItem.ts`
- `src/vs/workbench/contrib/chat/browser/modelPicker/modelPickerActionItem.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/attachInstructionsAction.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/chatModeActions.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/newPromptFileActions.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/promptCodingAgentActionContribution.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/promptCodingAgentActionOverlay.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/promptFileActions.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/runPromptAction.ts`
- `src/vs/workbench/contrib/chat/browser/promptSyntax/saveAsPromptFileActions.ts`
- `src/vs/workbench/contrib/chat/common/chatActions.ts`
- `src/vs/workbench/contrib/chat/common/promptSyntax/languageProviders/promptCodeActions.ts`
- `src/vs/workbench/contrib/chat/electron-browser/actions/chatDeveloperActions.ts`
- `src/vs/workbench/contrib/chat/electron-browser/actions/media/voiceChatActions.css`
- `src/vs/workbench/contrib/chat/electron-browser/actions/voiceChatActions.ts`
- `src/vs/workbench/contrib/chat/test/electron-browser/voiceChatActions.test.ts`
- `src/vs/workbench/contrib/codeActions/browser/codeActions.contribution.ts`
- `src/vs/workbench/contrib/codeActions/browser/codeActionsContribution.ts`
- `src/vs/workbench/contrib/codeEditor/electron-browser/inputClipboardActions.ts`
- `src/vs/workbench/contrib/comments/browser/commentFormActions.ts`
- `src/vs/workbench/contrib/comments/browser/commentThreadAdditionalActions.ts`
- `src/vs/workbench/contrib/comments/browser/commentsViewActions.ts`
- `src/vs/workbench/contrib/comments/browser/reactionsAction.ts`
- `src/vs/workbench/contrib/debug/browser/debugActionViewItems.ts`
- `src/vs/workbench/contrib/debug/browser/debugEditorActions.ts`
- `src/vs/workbench/contrib/emmet/browser/emmetActions.ts`
- `src/vs/workbench/contrib/emmet/test/browser/emmetAction.test.ts`
- `src/vs/workbench/contrib/extensions/browser/extensionsActions.ts`
- `src/vs/workbench/contrib/extensions/browser/media/extensionActions.css`
- `src/vs/workbench/contrib/extensions/common/installExtensionsTool.ts`
- `src/vs/workbench/contrib/extensions/common/reportExtensionIssueAction.ts`
- `src/vs/workbench/contrib/extensions/electron-browser/debugExtensionHostAction.ts`
- `src/vs/workbench/contrib/extensions/electron-browser/extensionsActions.ts`
- `src/vs/workbench/contrib/extensions/electron-browser/extensionsSlowActions.ts`
- `src/vs/workbench/contrib/extensions/test/electron-browser/extensionsActions.test.ts`
- `src/vs/workbench/contrib/files/browser/fileActions.contribution.ts`
- `src/vs/workbench/contrib/files/browser/fileActions.ts`
- `src/vs/workbench/contrib/files/electron-browser/fileActions.contribution.ts`
- `src/vs/workbench/contrib/files/test/browser/fileActions.test.ts`
- `src/vs/workbench/contrib/format/browser/formatActionsMultiple.ts`
- `src/vs/workbench/contrib/format/browser/formatActionsNone.ts`
- `src/vs/workbench/contrib/inlineChat/browser/inlineChatActions.ts`
- `src/vs/workbench/contrib/inlineChat/electron-browser/inlineChatActions.ts`
- `src/vs/workbench/contrib/list/browser/listResizeColumnAction.ts`
- `src/vs/workbench/contrib/localization/common/localizationsActions.ts`
- `src/vs/workbench/contrib/logs/common/logsActions.ts`
- `src/vs/workbench/contrib/logs/electron-browser/logsActions.ts`
- `src/vs/workbench/contrib/markers/browser/markersViewActions.css`
- `src/vs/workbench/contrib/markers/browser/markersViewActions.ts`
- `src/vs/workbench/contrib/mcp/browser/mcpServerActions.ts`
- `src/vs/workbench/contrib/mcp/browser/media/mcpServerAction.css`
- `src/vs/workbench/contrib/mcp/common/discovery/installedMcpServersDiscovery.ts`
- `src/vs/workbench/contrib/mergeEditor/browser/view/conflictActions.ts`
- `src/vs/workbench/contrib/multiDiffEditor/browser/actions.ts`
- `src/vs/workbench/contrib/notebook/browser/contrib/cellDiagnostics/cellDiagnosticsActions.ts`
- `src/vs/workbench/contrib/notebook/browser/contrib/layout/layoutActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/apiActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/cellOutputActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/chat/cellChatActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/coreActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/editActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/executeActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/insertCellActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/layoutActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/notebookIndentationActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/sectionActions.ts`
- `src/vs/workbench/contrib/notebook/browser/controller/variablesActions.ts`
- `src/vs/workbench/contrib/notebook/browser/diff/notebookDiffActions.ts`
- `src/vs/workbench/contrib/notebook/browser/media/notebookKernelActionViewItem.css`
- `src/vs/workbench/contrib/notebook/browser/view/cellParts/cellActionView.ts`
- `src/vs/workbench/contrib/notebook/test/browser/contrib/layoutActions.test.ts`
- `src/vs/workbench/contrib/outline/browser/outlineActions.ts`
- `src/vs/workbench/contrib/preferences/browser/preferencesActions.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsBase.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsCopy.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsFind.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsNav.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsRemoveReplace.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsSymbol.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsTextQuickAccess.ts`
- `src/vs/workbench/contrib/search/browser/searchActionsTopBar.ts`
- `src/vs/workbench/contrib/search/test/browser/searchActions.test.ts`
- `src/vs/workbench/contrib/searchEditor/browser/searchEditorActions.ts`
- `src/vs/workbench/contrib/snippets/browser/commands/abstractSnippetsActions.ts`
- `src/vs/workbench/contrib/snippets/browser/snippetCodeActionProvider.ts`
- `src/vs/workbench/contrib/terminal/browser/terminalActions.ts`
- `src/vs/workbench/contrib/terminal/test/browser/terminalActions.test.ts`
- `src/vs/workbench/contrib/terminalContrib/chat/browser/terminalChatActions.ts`
- `src/vs/workbench/contrib/terminalContrib/quickFix/browser/terminalQuickFixBuiltinActions.ts`
- `src/vs/workbench/contrib/terminalContrib/voice/browser/terminalVoiceActions.ts`
- `src/vs/workbench/contrib/testing/browser/testExplorerActions.ts`
- `src/vs/workbench/contrib/userDataProfile/browser/userDataProfileActions.ts`
- `src/vs/workbench/contrib/welcomeWalkthrough/browser/walkThroughActions.ts`
- `src/vs/workbench/electron-browser/actions/developerActions.ts`
- `src/vs/workbench/electron-browser/actions/installActions.ts`
- `src/vs/workbench/electron-browser/actions/media/actions.css`
- `src/vs/workbench/electron-browser/actions/windowActions.ts`
- `src/vscode-dts/vscode.proposed.codeActionAI.d.ts`
- `src/vscode-dts/vscode.proposed.codeActionRanges.d.ts`
- `src/vscode-dts/vscode.proposed.contribCommentEditorActionsMenu.d.ts`
- `src/vscode-dts/vscode.proposed.scmActionButton.d.ts`

## `ZeaZDev/ZeaClicker`

Status: **cloned**; commit `6a4be8b804fa`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

None detected or repository blocked.

## `ZeaZDev/zeazchain`

Status: **cloned**; commit `78ae5ed30c75`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `apps/api/src/main.ts` | 14 | JS/TS | `bootstrap` |
| `apps/api/src/modules/auth/world-id.service.ts` | 40 | JS/TS | `verify` |
| `apps/miniapp/components/WorldIDButton.tsx` | 34 | JS/TS | `resolveSecureStore` |
| `apps/miniapp/components/WorldIDButton.tsx` | 46 | JS/TS | `persistProof` |
| `apps/miniapp/lib/app-config.ts` | 21 | JS/TS | `loadConfig` |
| `setup_dev_env.sh` | 16 | Shell | `command_exists` |

### Automation pipeline files

- `.github/workflows/mcp-install.yml`
- `infra/docker-compose.yml`
- `script/install-mcp.sh`
- `script/workflows/mcp-install.yml`

## `ZeaZDev/ZeaZDev`

Status: **cloned**; commit `742079fa3f27`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `AUTH_SYSTEM.md` | 124 | `/api/auth/nonce` |
| `AUTH_SYSTEM.md` | 162 | `/api/auth/verify` |
| `AUTH_SYSTEM.md` | 197 | `/api/auth/nonce` |
| `AUTH_SYSTEM.md` | 216 | `/api/auth/verify` |
| `AUTH_SYSTEM.md` | 344 | `/api/worldid/verify` |
| `AUTH_SYSTEM.md` | 390 | `/api/worldid/verify` |
| `AUTH_SYSTEM.md` | 522 | `/api/auth/refresh` |
| `AUTH_SYSTEM.md` | 706 | `/api/stake` |
| `AUTH_SYSTEM.md` | 714 | `/api/admin/users/:id/ban` |
| `AUTH_SYSTEM.md` | 734 | `/api/auth/2fa/enable` |
| `AUTH_SYSTEM.md` | 759 | `/api/auth/2fa/verify` |
| `AUTH_SYSTEM.md` | 791 | `/api/auth/login/2fa` |
| `AUTH_SYSTEM.md` | 844 | `/api/kyc/init` |
| `AUTH_SYSTEM.md` | 863 | `/api/kyc/webhook` |
| `AUTH_SYSTEM.md` | 930 | `/api/` |
| `AUTH_SYSTEM.md` | 931 | `/api/auth/` |
| `AUTH_SYSTEM.md` | 942 | `/api/auth/csrf-token` |
| `AUTH_SYSTEM.md` | 955 | `/api/developer/keys/create` |
| `DEPLOYMENT.md` | 930 | `/health` |
| `Release/ZeaZDev-Release-v10.0.sh` | 18 | `/api/v1/claim` |
| `Release/ZeaZDev-Release-v10.0.sh` | 575 | `/api/results` |
| `Release/ZeaZDev-Release-v10.0.sh` | 584 | `/api/logtail` |
| `Release/ZeaZDev-Release-v10.0.sh` | 658 | `/api/v1/claim` |
| `Release/ZeaZDev-Release-v10.0.sh` | 738 | `/api/results` |
| `Release/ZeaZDev-Release-v10.0.sh` | 743 | `/api/logtail` |
| `Release/ZeaZDev-Release-v10.1.sh` | 18 | `/api/health` |
| `Release/ZeaZDev-Release-v10.1.sh` | 572 | `/api/health` |
| `Release/ZeaZDev-Release-v10.1.sh` | 577 | `/api/results` |
| `Release/ZeaZDev-Release-v10.1.sh` | 586 | `/api/logtail` |
| `Release/ZeaZDev-Release-v10.1.sh` | 651 | `/api/health` |
| `Release/ZeaZDev-Release-v10.1.sh` | 664 | `/api/v1/claim` |
| `Release/ZeaZDev-Release-v10.1.sh` | 744 | `/api/results` |
| `Release/ZeaZDev-Release-v10.1.sh` | 749 | `/api/logtail` |
| `Release/ZeaZDev-Release-v5.2.sh` | 392 | `/api/results` |
| `Release/ZeaZDev-Release-v5.2.sh` | 405 | `/api/logtail` |
| `Release/ZeaZDev-Release-v5.4.sh` | 450 | `/api/results` |
| `Release/ZeaZDev-Release-v5.4.sh` | 462 | `/api/logtail` |
| `Release/ZeaZDev-Release-v5.6.sh` | 245 | `/api/results` |
| `Release/ZeaZDev-Release-v5.6.sh` | 247 | `/api/logtail` |
| `Release/ZeaZDev-Release-v5.6.sh` | 296 | `/api/results` |
| `Release/ZeaZDev-Release-v5.7.sh` | 538 | `/api/results` |
| `Release/ZeaZDev-Release-v5.7.sh` | 551 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.0.sh` | 383 | `/api/results` |
| `Release/ZeaZDev-Release-v6.0.sh` | 392 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.0.sh` | 433 | `/api/results` |
| `Release/ZeaZDev-Release-v6.0.sh` | 434 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.1.sh` | 407 | `/api/results` |
| `Release/ZeaZDev-Release-v6.1.sh` | 416 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.1.sh` | 457 | `/api/results` |
| `Release/ZeaZDev-Release-v6.1.sh` | 458 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.2.sh` | 413 | `/api/results` |
| `Release/ZeaZDev-Release-v6.2.sh` | 422 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.2.sh` | 463 | `/api/results` |
| `Release/ZeaZDev-Release-v6.2.sh` | 464 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.3.sh` | 415 | `/api/results` |
| `Release/ZeaZDev-Release-v6.3.sh` | 424 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.3.sh` | 465 | `/api/results` |
| `Release/ZeaZDev-Release-v6.3.sh` | 466 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.4.sh` | 419 | `/api/results` |
| `Release/ZeaZDev-Release-v6.4.sh` | 428 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.4.sh` | 469 | `/api/results` |
| `Release/ZeaZDev-Release-v6.4.sh` | 470 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.5.sh` | 423 | `/api/results` |
| `Release/ZeaZDev-Release-v6.5.sh` | 432 | `/api/logtail` |
| `Release/ZeaZDev-Release-v6.5.sh` | 473 | `/api/results` |
| `Release/ZeaZDev-Release-v6.5.sh` | 474 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.0.sh` | 463 | `/api/results` |
| `Release/ZeaZDev-Release-v7.0.sh` | 472 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.0.sh` | 513 | `/api/results` |
| `Release/ZeaZDev-Release-v7.0.sh` | 514 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.1.sh` | 463 | `/api/results` |
| `Release/ZeaZDev-Release-v7.1.sh` | 472 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.1.sh` | 513 | `/api/results` |
| `Release/ZeaZDev-Release-v7.1.sh` | 514 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.2.sh` | 517 | `/api/results` |
| `Release/ZeaZDev-Release-v7.2.sh` | 526 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.2.sh` | 567 | `/api/results` |
| `Release/ZeaZDev-Release-v7.2.sh` | 568 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.3.sh` | 459 | `/api/results` |
| `Release/ZeaZDev-Release-v7.3.sh` | 468 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.3.sh` | 509 | `/api/results` |
| `Release/ZeaZDev-Release-v7.3.sh` | 510 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.4.sh` | 448 | `/api/results` |
| `Release/ZeaZDev-Release-v7.4.sh` | 457 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.4.sh` | 498 | `/api/results` |
| `Release/ZeaZDev-Release-v7.4.sh` | 499 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.5.sh` | 501 | `/api/results` |
| `Release/ZeaZDev-Release-v7.5.sh` | 510 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.5.sh` | 551 | `/api/results` |
| `Release/ZeaZDev-Release-v7.5.sh` | 552 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.6.sh` | 463 | `/api/results` |
| `Release/ZeaZDev-Release-v7.6.sh` | 472 | `/api/logtail` |
| `Release/ZeaZDev-Release-v7.6.sh` | 513 | `/api/results` |
| `Release/ZeaZDev-Release-v7.6.sh` | 514 | `/api/logtail` |
| `Release/ZeaZDev-Release-v8.0.sh` | 466 | `/api/results` |
| `Release/ZeaZDev-Release-v8.0.sh` | 475 | `/api/logtail` |
| `Release/ZeaZDev-Release-v8.0.sh` | 516 | `/api/results` |
| `Release/ZeaZDev-Release-v8.0.sh` | 517 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.0.sh` | 552 | `/api/results` |
| `Release/ZeaZDev-Release-v9.0.sh` | 561 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.0.sh` | 602 | `/api/results` |
| `Release/ZeaZDev-Release-v9.0.sh` | 607 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.1.sh` | 525 | `/api/results` |
| `Release/ZeaZDev-Release-v9.1.sh` | 534 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.1.sh` | 575 | `/api/results` |
| `Release/ZeaZDev-Release-v9.1.sh` | 580 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.2.sh` | 527 | `/api/results` |
| `Release/ZeaZDev-Release-v9.2.sh` | 536 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.2.sh` | 577 | `/api/results` |
| `Release/ZeaZDev-Release-v9.2.sh` | 582 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.3.sh` | 589 | `/api/results` |
| `Release/ZeaZDev-Release-v9.3.sh` | 598 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.3.sh` | 639 | `/api/results` |
| `Release/ZeaZDev-Release-v9.3.sh` | 644 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.4.sh` | 526 | `/api/results` |
| `Release/ZeaZDev-Release-v9.4.sh` | 535 | `/api/logtail` |
| `Release/ZeaZDev-Release-v9.4.sh` | 576 | `/api/results` |
| `Release/ZeaZDev-Release-v9.4.sh` | 581 | `/api/logtail` |
| `Release/ZeaZDev-Release.sh` | 401 | `/api/results` |
| `Release/ZeaZDev-Release.sh` | 414 | `/api/logtail` |
| `Rewards/api.php` | 42 | `/api/checkin` |
| `ZeaZDev-Release-v10.1.sh` | 18 | `/api/health` |
| `ZeaZDev-Release-v10.1.sh` | 572 | `/api/health` |
| `ZeaZDev-Release-v10.1.sh` | 577 | `/api/results` |
| `ZeaZDev-Release-v10.1.sh` | 586 | `/api/logtail` |
| `ZeaZDev-Release-v10.1.sh` | 651 | `/api/health` |
| `ZeaZDev-Release-v10.1.sh` | 664 | `/api/v1/claim` |
| `ZeaZDev-Release-v10.1.sh` | 744 | `/api/results` |
| `ZeaZDev-Release-v10.1.sh` | 749 | `/api/logtail` |
| `ZeaZDev-Repair-Dashboard.sh` | 49 | `/api/health` |
| `ZeaZDev-Repair-Dashboard.sh` | 54 | `/api/vault` |
| `server/verifier.js` | 84 | `/api/verify-worldid` |
| `server/verifier.js` | 140 | `/api/check-verification/:address` |
| `server/verifier.js` | 163 | `/api/checkin-status/:address` |
| `server/verifier.js` | 190 | `/api/claim-airdrop` |
| `server/verifier.js` | 220 | `/api/daily-checkin` |
| `server/verifier.js` | 255 | `/api/swap-quote` |
| `server/verifier.js` | 284 | `/api/execute-swap` |
| `server/verifier.js` | 358 | `/health` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `Release/ZeaZDev-Release-v10.0.sh` | 30 | Shell | `info` |
| `Release/ZeaZDev-Release-v10.0.sh` | 31 | Shell | `warn` |
| `Release/ZeaZDev-Release-v10.0.sh` | 32 | Shell | `err` |
| `Release/ZeaZDev-Release-v10.0.sh` | 435 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v10.1.sh` | 28 | Shell | `info` |
| `Release/ZeaZDev-Release-v10.1.sh` | 29 | Shell | `warn` |
| `Release/ZeaZDev-Release-v10.1.sh` | 30 | Shell | `err` |
| `Release/ZeaZDev-Release-v10.1.sh` | 430 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v5.2.sh` | 11 | Shell | `info` |
| `Release/ZeaZDev-Release-v5.2.sh` | 12 | Shell | `warn` |
| `Release/ZeaZDev-Release-v5.2.sh` | 13 | Shell | `err` |
| `Release/ZeaZDev-Release-v5.2.sh` | 69 | Shell | `normalize_rpc` |
| `Release/ZeaZDev-Release-v5.2.sh` | 101 | Shell | `try_fetch_world_router` |
| `Release/ZeaZDev-Release-v5.2.sh` | 130 | Shell | `write_hardhat_config` |
| `Release/ZeaZDev-Release-v5.2.sh` | 254 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v5.4.sh` | 19 | Shell | `info` |
| `Release/ZeaZDev-Release-v5.4.sh` | 20 | Shell | `warn` |
| `Release/ZeaZDev-Release-v5.4.sh` | 21 | Shell | `err` |
| `Release/ZeaZDev-Release-v5.4.sh` | 89 | Shell | `normalize_rpc` |
| `Release/ZeaZDev-Release-v5.4.sh` | 102 | Shell | `ensure_swap_if_low_memory` |
| `Release/ZeaZDev-Release-v5.4.sh` | 124 | Shell | `try_fetch_world_router` |
| `Release/ZeaZDev-Release-v5.4.sh` | 277 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v5.5.sh` | 17 | Shell | `info` |
| `Release/ZeaZDev-Release-v5.5.sh` | 18 | Shell | `warn` |
| `Release/ZeaZDev-Release-v5.5.sh` | 19 | Shell | `err` |
| `Release/ZeaZDev-Release-v5.5.sh` | 189 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v5.6.sh` | 17 | Shell | `info` |
| `Release/ZeaZDev-Release-v5.6.sh` | 18 | Shell | `warn` |
| `Release/ZeaZDev-Release-v5.6.sh` | 19 | Shell | `err` |
| `Release/ZeaZDev-Release-v5.6.sh` | 98 | Shell | `install_nginx_and_certbot` |
| `Release/ZeaZDev-Release-v5.6.sh` | 127 | Shell | `configure_nginx_proxy` |
| `Release/ZeaZDev-Release-v5.6.sh` | 160 | Shell | `request_ssl` |
| `Release/ZeaZDev-Release-v5.6.sh` | 179 | Shell | `create_systemd_services` |
| `Release/ZeaZDev-Release-v5.6.sh` | 212 | Shell | `bootstrap_project_files` |
| `Release/ZeaZDev-Release-v5.6.sh` | 306 | Shell | `start_services_and_confirm` |
| `Release/ZeaZDev-Release-v5.6.sh` | 325 | Shell | `main_flow` |
| `Release/ZeaZDev-Release-v5.7.sh` | 17 | Shell | `info` |
| `Release/ZeaZDev-Release-v5.7.sh` | 18 | Shell | `warn` |
| `Release/ZeaZDev-Release-v5.7.sh` | 19 | Shell | `err` |
| `Release/ZeaZDev-Release-v5.7.sh` | 36 | Shell | `create_dir` |
| `Release/ZeaZDev-Release-v5.7.sh` | 69 | Shell | `normalize_rpc` |
| `Release/ZeaZDev-Release-v5.7.sh` | 80 | Shell | `check_command` |
| `Release/ZeaZDev-Release-v5.7.sh` | 123 | Shell | `install_system_deps` |
| `Release/ZeaZDev-Release-v5.7.sh` | 142 | Shell | `install_node_npm` |
| `Release/ZeaZDev-Release-v5.7.sh` | 386 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v5.8.sh` | 17 | Shell | `info` |
| `Release/ZeaZDev-Release-v5.8.sh` | 18 | Shell | `warn` |
| `Release/ZeaZDev-Release-v5.8.sh` | 19 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.0.sh` | 21 | Shell | `info` |
| `Release/ZeaZDev-Release-v6.0.sh` | 22 | Shell | `warn` |
| `Release/ZeaZDev-Release-v6.0.sh` | 23 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.0.sh` | 248 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v6.1.sh` | 22 | Shell | `info` |
| `Release/ZeaZDev-Release-v6.1.sh` | 23 | Shell | `warn` |
| `Release/ZeaZDev-Release-v6.1.sh` | 24 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.1.sh` | 258 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v6.2.sh` | 24 | Shell | `info` |
| `Release/ZeaZDev-Release-v6.2.sh` | 25 | Shell | `warn` |
| `Release/ZeaZDev-Release-v6.2.sh` | 26 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.2.sh` | 267 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v6.3.sh` | 26 | Shell | `info` |
| `Release/ZeaZDev-Release-v6.3.sh` | 27 | Shell | `warn` |
| `Release/ZeaZDev-Release-v6.3.sh` | 28 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.3.sh` | 269 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v6.4.sh` | 29 | Shell | `info` |
| `Release/ZeaZDev-Release-v6.4.sh` | 30 | Shell | `warn` |
| `Release/ZeaZDev-Release-v6.4.sh` | 31 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.4.sh` | 272 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v6.5.sh` | 34 | Shell | `info` |
| `Release/ZeaZDev-Release-v6.5.sh` | 35 | Shell | `warn` |
| `Release/ZeaZDev-Release-v6.5.sh` | 36 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.5.sh` | 277 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v6.9.sh` | 30 | Shell | `info` |
| `Release/ZeaZDev-Release-v6.9.sh` | 31 | Shell | `warn` |
| `Release/ZeaZDev-Release-v6.9.sh` | 32 | Shell | `err` |
| `Release/ZeaZDev-Release-v6.9.sh` | 323 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v7.0.sh` | 23 | Shell | `info` |
| `Release/ZeaZDev-Release-v7.0.sh` | 24 | Shell | `warn` |
| `Release/ZeaZDev-Release-v7.0.sh` | 25 | Shell | `err` |
| `Release/ZeaZDev-Release-v7.0.sh` | 316 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v7.1.sh` | 23 | Shell | `info` |
| `Release/ZeaZDev-Release-v7.1.sh` | 24 | Shell | `warn` |
| `Release/ZeaZDev-Release-v7.1.sh` | 25 | Shell | `err` |
| `Release/ZeaZDev-Release-v7.1.sh` | 316 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v7.2.sh` | 29 | Shell | `info` |
| `Release/ZeaZDev-Release-v7.2.sh` | 30 | Shell | `warn` |
| `Release/ZeaZDev-Release-v7.2.sh` | 31 | Shell | `err` |
| `Release/ZeaZDev-Release-v7.2.sh` | 366 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v7.3.sh` | 32 | Shell | `info` |
| `Release/ZeaZDev-Release-v7.3.sh` | 33 | Shell | `warn` |
| `Release/ZeaZDev-Release-v7.3.sh` | 34 | Shell | `err` |
| `Release/ZeaZDev-Release-v7.3.sh` | 317 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v7.4.sh` | 27 | Shell | `info` |
| `Release/ZeaZDev-Release-v7.4.sh` | 28 | Shell | `warn` |
| `Release/ZeaZDev-Release-v7.4.sh` | 29 | Shell | `err` |
| `Release/ZeaZDev-Release-v7.4.sh` | 312 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v7.5.sh` | 28 | Shell | `info` |
| `Release/ZeaZDev-Release-v7.5.sh` | 29 | Shell | `warn` |
| `Release/ZeaZDev-Release-v7.5.sh` | 30 | Shell | `err` |
| `Release/ZeaZDev-Release-v7.5.sh` | 358 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v7.6.sh` | 26 | Shell | `info` |
| `Release/ZeaZDev-Release-v7.6.sh` | 27 | Shell | `warn` |
| `Release/ZeaZDev-Release-v7.6.sh` | 28 | Shell | `err` |
| `Release/ZeaZDev-Release-v7.6.sh` | 328 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v8.0.sh` | 29 | Shell | `info` |
| `Release/ZeaZDev-Release-v8.0.sh` | 30 | Shell | `warn` |
| `Release/ZeaZDev-Release-v8.0.sh` | 31 | Shell | `err` |
| `Release/ZeaZDev-Release-v8.0.sh` | 331 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v9.0.sh` | 29 | Shell | `info` |
| `Release/ZeaZDev-Release-v9.0.sh` | 30 | Shell | `warn` |
| `Release/ZeaZDev-Release-v9.0.sh` | 31 | Shell | `err` |
| `Release/ZeaZDev-Release-v9.0.sh` | 416 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v9.1.sh` | 27 | Shell | `info` |
| `Release/ZeaZDev-Release-v9.1.sh` | 28 | Shell | `warn` |
| `Release/ZeaZDev-Release-v9.1.sh` | 29 | Shell | `err` |
| `Release/ZeaZDev-Release-v9.1.sh` | 389 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v9.2.sh` | 28 | Shell | `info` |
| `Release/ZeaZDev-Release-v9.2.sh` | 29 | Shell | `warn` |
| `Release/ZeaZDev-Release-v9.2.sh` | 30 | Shell | `err` |
| `Release/ZeaZDev-Release-v9.2.sh` | 391 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v9.3.sh` | 28 | Shell | `info` |
| `Release/ZeaZDev-Release-v9.3.sh` | 29 | Shell | `warn` |
| `Release/ZeaZDev-Release-v9.3.sh` | 30 | Shell | `err` |
| `Release/ZeaZDev-Release-v9.3.sh` | 446 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release-v9.4.sh` | 29 | Shell | `info` |
| `Release/ZeaZDev-Release-v9.4.sh` | 30 | Shell | `warn` |
| `Release/ZeaZDev-Release-v9.4.sh` | 31 | Shell | `err` |
| `Release/ZeaZDev-Release-v9.4.sh` | 392 | Shell | `run_network_job` |
| `Release/ZeaZDev-Release.sh` | 20 | Shell | `info` |
| `Release/ZeaZDev-Release.sh` | 21 | Shell | `warn` |
| `Release/ZeaZDev-Release.sh` | 22 | Shell | `err` |
| `Release/ZeaZDev-Release.sh` | 78 | Shell | `normalize_rpc` |
| `Release/ZeaZDev-Release.sh` | 110 | Shell | `try_fetch_world_router` |
| `Release/ZeaZDev-Release.sh` | 139 | Shell | `write_hardhat_config` |
| `Release/ZeaZDev-Release.sh` | 263 | Shell | `run_network_job` |
| `Rewards/api.php` | 23 | PHP | `resp` |
| `Rewards/api.php` | 25 | PHP | `isValidEthereumAddress` |
| `ZeaZDev-Release-v10.1.sh` | 28 | Shell | `info` |
| `ZeaZDev-Release-v10.1.sh` | 29 | Shell | `warn` |
| `ZeaZDev-Release-v10.1.sh` | 30 | Shell | `err` |
| `ZeaZDev-Release-v10.1.sh` | 430 | Shell | `run_network_job` |
| `mini-app/src/screens/AuthGate.tsx` | 39 | JS/TS | `checkExistingAuth` |
| `mini-app/src/screens/AuthGate.tsx` | 56 | JS/TS | `handleVerifySuccess` |
| `mini-app/src/screens/AuthGate.tsx` | 121 | JS/TS | `handleVerifyFailure` |
| `mini-app/src/screens/AuthGate.tsx` | 133 | JS/TS | `startVerification` |
| `mini-app/src/screens/RewardScreen.tsx` | 57 | JS/TS | `loadUserStatus` |
| `mini-app/src/screens/RewardScreen.tsx` | 80 | JS/TS | `handleClaimAirdrop` |
| `mini-app/src/screens/RewardScreen.tsx` | 94 | JS/TS | `executeClaimAirdrop` |
| `mini-app/src/screens/RewardScreen.tsx` | 125 | JS/TS | `handleCheckIn` |
| `mini-app/src/screens/SwapTradeScreen.tsx` | 64 | JS/TS | `loadQuote` |
| `mini-app/src/screens/SwapTradeScreen.tsx` | 90 | JS/TS | `handleSwap` |
| `mini-app/src/screens/SwapTradeScreen.tsx` | 121 | JS/TS | `confirmSwap` |
| `mini-app/src/screens/SwapTradeScreen.tsx` | 135 | JS/TS | `executeSwapTransaction` |
| `mini-app/src/screens/SwapTradeScreen.tsx` | 176 | JS/TS | `swapTokenPositions` |
| `mini-app/src/screens/SwapTradeScreen.tsx` | 187 | JS/TS | `selectToken` |
| `mini-app/src/screens/WalletScreen.tsx` | 53 | JS/TS | `loadBalances` |
| `mini-app/src/screens/WalletScreen.tsx` | 76 | JS/TS | `handleSend` |
| `mini-app/src/screens/WalletScreen.tsx` | 110 | JS/TS | `executeSend` |
| `mini-app/src/screens/WalletScreen.tsx` | 151 | JS/TS | `copyAddress` |
| `mini-app/src/services/rewards.ts` | 21 | JS/TS | `getCheckInStatus` |
| `mini-app/src/services/rewards.ts` | 58 | JS/TS | `claimAirdrop` |
| `mini-app/src/services/rewards.ts` | 97 | JS/TS | `dailyCheckIn` |
| `mini-app/src/services/swap.ts` | 26 | JS/TS | `getSwapQuote` |
| `mini-app/src/services/swap.ts` | 74 | JS/TS | `executeSwap` |
| `mini-app/src/services/wallet.ts` | 15 | JS/TS | `getBalance` |
| `mini-app/src/services/wallet.ts` | 54 | JS/TS | `sendToken` |
| `mini-app/src/services/worldid.ts` | 27 | JS/TS | `verifyProof` |
| `mini-app/src/services/worldid.ts` | 57 | JS/TS | `checkVerificationStatus` |
| `server/verifier.js` | 312 | JS/TS | `verifyWithWorldID` |
| `server/verifier.js` | 350 | JS/TS | `parseProof` |
| `start-all-phases.sh` | 53 | Shell | `print_banner` |
| `start-all-phases.sh` | 63 | Shell | `log_info` |
| `start-all-phases.sh` | 67 | Shell | `log_warn` |
| `start-all-phases.sh` | 71 | Shell | `log_error` |
| `start-all-phases.sh` | 75 | Shell | `log_step` |
| `start-all-phases.sh` | 83 | Shell | `check_prerequisites` |
| `start-all-phases.sh` | 121 | Shell | `check_directory_structure` |
| `start-all-phases.sh` | 154 | Shell | `install_dependencies` |
| `start-all-phases.sh` | 193 | Shell | `check_environment` |
| `start-all-phases.sh` | 257 | Shell | `start_backend` |
| `start-all-phases.sh` | 304 | Shell | `start_frontend` |
| `start-all-phases.sh` | 352 | Shell | `check_contracts` |
| `start-all-phases.sh` | 377 | Shell | `show_status` |
| `start-all-phases.sh` | 435 | Shell | `stop_services` |
| `start-all-phases.sh` | 473 | Shell | `show_help` |
| `start-all-phases.sh` | 514 | Shell | `main` |

### Automation pipeline files

- `.github/workflows/auto-assign.yml`
- `.github/workflows/deploy.yml`
- `.github/workflows/proof-html.yml`
- `DEPLOYMENT.md`
- `scripts/deploy.js`

## `ZeaZDev/ZeaZDev-Omega`

Status: **cloned**; commit `41adac95ad68`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 8 | `metric` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 19 | `metrics/:userId` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 36 | `dashboard/:userId` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 41 | `behavior/:userId` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 46 | `recommendations/:userId` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 51 | `compliance/snapshot` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 56 | `prediction` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 68 | `predictions/:userId` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 73 | `fraud/:userId` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 79 | `fraud/:alertId/resolve` |
| `apps/backend/src/modules/auth/auth.controller.ts` | 33 | `world-id/verify` |
| `apps/backend/src/modules/auth/auth.controller.ts` | 89 | `wallet/connect` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 25 | `quote` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 41 | `initiate` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 63 | `complete` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 77 | `transactions/:userId` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 85 | `transaction/:hash` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 93 | `chains` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 101 | `chains/:chainId` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 109 | `liquidity/:chainId` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 125 | `liquidity/add` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 143 | `liquidity/remove` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 161 | `stats` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 20 | `swap/quote` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 29 | `swap/execute` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 50 | `pool/liquidity` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 58 | `stake` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 69 | `stake/:id/claim` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 74 | `stake/:id/unstake` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 82 | `stake/user/:userId` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 87 | `stake/analytics/:userId` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 23 | `whitelabel` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 47 | `whitelabel/:orgId` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 52 | `whitelabel/:orgId` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 68 | `whitelabel/:orgId/features` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 73 | `whitelabel/:orgId/features` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 84 | `api-usage/:orgId` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 89 | `rate-limits` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 97 | `developer/app` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 104 | `developer/:developerId/apps` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 109 | `developer/app/:appId` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 120 | `plugins` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 128 | `plugins` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 153 | `plugins/install` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 160 | `plugins/installed/:orgId` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 165 | `plugins/uninstall` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 175 | `sdk/list` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 180 | `sdk/example` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 191 | `api/endpoints` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 520 | `/api/wallet/create` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 521 | `/api/wallet/:userId` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 522 | `/api/wallet/:userId/balance` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 528 | `/api/bridge/quote` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 529 | `/api/bridge/initiate` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 530 | `/api/bridge/transaction/:hash` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 536 | `/api/game/types` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 537 | `/api/game/slots/play` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 538 | `/api/game/poker/play` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 539 | `/api/game/stats/:userId` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 545 | `/api/fintech/promptpay/generate` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 546 | `/api/fintech/promptpay/verify/:id` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 547 | `/api/fintech/card/issue` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 553 | `/api/defi/stake` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 554 | `/api/defi/unstake` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 555 | `/api/defi/swap` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 24 | `card/issue` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 41 | `card/user/:userId` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 46 | `bank/thai/deposit` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 61 | `bank/thai/withdraw` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 76 | `transactions/user/:userId` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 84 | `promptpay/generate` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 102 | `promptpay/verify/:transactionId` |
| `apps/backend/src/modules/fintech/fintech.controller.ts` | 110 | `promptpay/webhook` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 21 | `banks` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 30 | `currencies` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 39 | `exchange-rates` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 48 | `kyc/levels` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 57 | `stats` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 70 | `kyc/submit` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 87 | `kyc/:userId` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 100 | `bank-accounts` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 115 | `bank-accounts/verify` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 127 | `bank-accounts/:userId` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 140 | `deposits/quote` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 153 | `deposits` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 172 | `withdrawals/quote` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 185 | `withdrawals` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 204 | `transactions/:userId` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 222 | `transactions/details/:txId` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 235 | `cards/virtual` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 248 | `cards/physical` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 262 | `cards/:userId` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 271 | `cards/fund` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 284 | `cards/:cardId/transactions` |
| `apps/backend/src/modules/game/game.controller.ts` | 23 | `types` |
| `apps/backend/src/modules/game/game.controller.ts` | 31 | `stats/:userId` |
| `apps/backend/src/modules/game/game.controller.ts` | 39 | `slots/play` |
| `apps/backend/src/modules/game/game.controller.ts` | 56 | `slots/complete` |
| `apps/backend/src/modules/game/game.controller.ts` | 74 | `poker/create` |
| `apps/backend/src/modules/game/game.controller.ts` | 91 | `poker/play` |
| `apps/backend/src/modules/game/game.controller.ts` | 107 | `roulette/play` |
| `apps/backend/src/modules/game/game.controller.ts` | 129 | `sports/events` |
| `apps/backend/src/modules/game/game.controller.ts` | 137 | `sports/bet` |
| `apps/backend/src/modules/game/game.controller.ts` | 159 | `sessions/user/:userId` |
| `apps/backend/src/modules/game/game.controller.ts` | 164 | `sessions/:sessionId` |
| `apps/backend/src/modules/game/game.controller.ts` | 169 | `leaderboard/:gameType?` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 24 | `slots/spin` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 40 | `slots/history/:userId` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 53 | `slots/verify` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 67 | `nfts/:userId` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 76 | `tournaments/create` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 94 | `tournaments/join` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 104 | `tournaments/active` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 109 | `tournaments/:tournamentId/leaderboard` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 118 | `leaderboard` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 123 | `stats/:userId` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 20 | `proposal` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 42 | `vote` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 60 | `proposal/:id` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 65 | `proposals` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 76 | `voting-power/:address` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 82 | `proposal/:id/execute` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 87 | `proposal/:id/cancel` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 20 | `claim` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 39 | `user/:userId` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 44 | `user/:userId/eligibility` |
| `apps/backend/src/modules/social/social.controller.ts` | 8 | `profile` |
| `apps/backend/src/modules/social/social.controller.ts` | 13 | `profile/:userId` |
| `apps/backend/src/modules/social/social.controller.ts` | 18 | `profile/:userId` |
| `apps/backend/src/modules/social/social.controller.ts` | 23 | `follow` |
| `apps/backend/src/modules/social/social.controller.ts` | 28 | `unfollow` |
| `apps/backend/src/modules/social/social.controller.ts` | 33 | `followers/:userId` |
| `apps/backend/src/modules/social/social.controller.ts` | 38 | `following/:userId` |
| `apps/backend/src/modules/social/social.controller.ts` | 43 | `post` |
| `apps/backend/src/modules/social/social.controller.ts` | 48 | `feed/:userId` |
| `apps/backend/src/modules/social/social.controller.ts` | 53 | `post/:postId/like` |
| `apps/backend/src/modules/social/social.controller.ts` | 58 | `achievements/:userId` |
| `apps/backend/src/modules/social/social.controller.ts` | 63 | `leaderboard` |
| `docs/API_DOCUMENTATION.md` | 521 | `/api/wallet/create` |
| `docs/API_DOCUMENTATION.md` | 522 | `/api/wallet/:userId` |
| `docs/API_DOCUMENTATION.md` | 523 | `/api/wallet/:userId/balance` |
| `docs/API_DOCUMENTATION.md` | 529 | `/api/bridge/quote` |
| `docs/API_DOCUMENTATION.md` | 530 | `/api/bridge/initiate` |
| `docs/API_DOCUMENTATION.md` | 531 | `/api/bridge/transaction/:hash` |
| `docs/API_DOCUMENTATION.md` | 586 | `/api/wallet/create` |
| `docs/API_DOCUMENTATION.md` | 596 | `/api/bridge/initiate` |
| `docs/ENTERPRISE_FEATURES.md` | 514 | `/api/wallet/create` |
| `docs/PHASE8_COMPLETION.md` | 149 | `/api/wallet/create` |
| `docs/PHASE8_COMPLETION.md` | 150 | `/api/wallet/:userId` |
| `docs/PHASE8_COMPLETION.md` | 151 | `/api/wallet/:userId/balance` |
| `docs/PHASE8_COMPLETION.md` | 157 | `/api/bridge/quote` |
| `docs/PHASE8_COMPLETION.md` | 158 | `/api/bridge/initiate` |
| `docs/PHASE8_COMPLETION.md` | 159 | `/api/bridge/transaction/:hash` |
| `docs/PHASE8_COMPLETION.md` | 200 | `/api/wallet/create` |
| `docs/PHASE9_10_SECURITY_SUMMARY.md` | 177 | `leaderboard` |
| `docs/PHASE9_10_SECURITY_SUMMARY.md` | 182 | `fraud/:userId` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `apps/backend/src/main.ts` | 17 | JS/TS | `bootstrap` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 6 | JS/TS | `constructor` |
| `apps/backend/src/modules/analytics/analytics.controller.ts` | 52 | JS/TS | `getComplianceSnapshot` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 6 | JS/TS | `constructor` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 8 | JS/TS | `trackMetric` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 29 | JS/TS | `getUserMetrics` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 44 | JS/TS | `getDashboard` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 62 | JS/TS | `createPrediction` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 73 | JS/TS | `getUserPredictions` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 84 | JS/TS | `updatePredictionAccuracy` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 91 | JS/TS | `createFraudAlert` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 102 | JS/TS | `getFraudAlerts` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 112 | JS/TS | `resolveFraudAlert` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 122 | JS/TS | `analyzeUserBehavior` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 147 | JS/TS | `getComplianceSnapshot` |
| `apps/backend/src/modules/analytics/analytics.service.ts` | 167 | JS/TS | `generateRecommendations` |
| `apps/backend/src/modules/auth/worldcoin.service.ts` | 21 | JS/TS | `constructor` |
| `apps/backend/src/modules/auth/worldcoin.service.ts` | 100 | JS/TS | `parseProof` |
| `apps/backend/src/modules/auth/worldcoin.service.ts` | 121 | JS/TS | `getNullifierHash` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 20 | JS/TS | `constructor` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 94 | JS/TS | `getSupportedChains` |
| `apps/backend/src/modules/bridge/bridge.controller.ts` | 162 | JS/TS | `getBridgeStats` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 117 | JS/TS | `constructor` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 162 | JS/TS | `getChainConfig` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 169 | JS/TS | `getSupportedChains` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 312 | JS/TS | `getUserBridgeTransactions` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 324 | JS/TS | `getBridgeTransaction` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 337 | JS/TS | `isSupportedChain` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 341 | JS/TS | `assertSupportedToken` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 353 | JS/TS | `mapToBridgeTransaction` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 373 | JS/TS | `monitorBridgeTransaction` |
| `apps/backend/src/modules/bridge/bridge.service.ts` | 390 | JS/TS | `getBridgeStats` |
| `apps/backend/src/modules/defi/defi.controller.ts` | 18 | JS/TS | `constructor` |
| `apps/backend/src/modules/defi/defi.service.ts` | 23 | JS/TS | `constructor` |
| `apps/backend/src/modules/defi/defi.service.ts` | 119 | JS/TS | `getPoolLiquidity` |
| `apps/backend/src/modules/defi/defi.service.ts` | 197 | JS/TS | `getSwapQuote` |
| `apps/backend/src/modules/defi/defi.service.ts` | 282 | JS/TS | `createStake` |
| `apps/backend/src/modules/defi/defi.service.ts` | 303 | JS/TS | `claimStakeRewards` |
| `apps/backend/src/modules/defi/defi.service.ts` | 339 | JS/TS | `getUserStakes` |
| `apps/backend/src/modules/defi/defi.service.ts` | 365 | JS/TS | `unstake` |
| `apps/backend/src/modules/defi/defi.service.ts` | 437 | JS/TS | `getStakeAnalytics` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 18 | JS/TS | `constructor` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 176 | JS/TS | `getSDKList` |
| `apps/backend/src/modules/enterprise/enterprise.controller.ts` | 192 | JS/TS | `getAPIEndpoints` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 51 | JS/TS | `constructor` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 78 | JS/TS | `getWhiteLabel` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 101 | JS/TS | `trackApiUsage` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 138 | JS/TS | `getApiUsage` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 151 | JS/TS | `createDeveloperApp` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 167 | JS/TS | `getDeveloperApps` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 183 | JS/TS | `getDeveloperApp` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 189 | JS/TS | `updateDeveloperApp` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 196 | JS/TS | `verifyApiKey` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 204 | JS/TS | `generateApiKey` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 208 | JS/TS | `generateApiSecret` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 251 | JS/TS | `listPlugins` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 318 | JS/TS | `installPlugin` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 339 | JS/TS | `getInstalledPlugins` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 354 | JS/TS | `uninstallPlugin` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 374 | JS/TS | `getSDKList` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 416 | JS/TS | `generateSDKExample` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 515 | JS/TS | `getAPIEndpoints` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 563 | JS/TS | `getRateLimits` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 604 | JS/TS | `getWhiteLabelFeatures` |
| `apps/backend/src/modules/enterprise/enterprise.service.ts` | 632 | JS/TS | `updateWhiteLabelFeatures` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 39 | JS/TS | `constructor` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 46 | JS/TS | `deposit` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 86 | JS/TS | `withdraw` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 141 | JS/TS | `getUserTransactions` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 164 | JS/TS | `verifyBankAccount` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 274 | JS/TS | `formatPromptPayTag` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 284 | JS/TS | `formatReferenceTag` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 292 | JS/TS | `calculateCRC16` |
| `apps/backend/src/modules/fintech/bank.thai.service.ts` | 320 | JS/TS | `verifyPromptPayPayment` |
| `apps/backend/src/modules/fintech/card.service.ts` | 22 | JS/TS | `constructor` |
| `apps/backend/src/modules/fintech/card.service.ts` | 27 | JS/TS | `issueCard` |
| `apps/backend/src/modules/fintech/card.service.ts` | 75 | JS/TS | `getUserCard` |
| `apps/backend/src/modules/fintech/card.service.ts` | 97 | JS/TS | `generateMockCardNumber` |
| `apps/backend/src/modules/fintech/card.service.ts` | 104 | JS/TS | `maskCardNumber` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 11 | JS/TS | `constructor` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 22 | JS/TS | `getSupportedBanks` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 31 | JS/TS | `getSupportedCurrencies` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 40 | JS/TS | `getExchangeRates` |
| `apps/backend/src/modules/fintech/tradfi.controller.ts` | 49 | JS/TS | `getKYCLevels` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 110 | JS/TS | `constructor` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 118 | JS/TS | `getSupportedBanks` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 125 | JS/TS | `getSupportedCurrencies` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 135 | JS/TS | `getKYCLevels` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 180 | JS/TS | `getUserKYC` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 252 | JS/TS | `getUserBankAccounts` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 283 | JS/TS | `getExchangeRates` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 531 | JS/TS | `getTransaction` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 608 | JS/TS | `getUserCards` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 652 | JS/TS | `getCardTransactions` |
| `apps/backend/src/modules/fintech/tradfi.service.ts` | 679 | JS/TS | `getTradFiStats` |
| `apps/backend/src/modules/game/game.controller.ts` | 18 | JS/TS | `constructor` |
| `apps/backend/src/modules/game/game.controller.ts` | 24 | JS/TS | `getGameTypes` |
| `apps/backend/src/modules/game/game.service.ts` | 38 | JS/TS | `constructor` |
| `apps/backend/src/modules/game/game.service.ts` | 180 | JS/TS | `getLeaderboard` |
| `apps/backend/src/modules/game/game.service.ts` | 204 | JS/TS | `getUserSessions` |
| `apps/backend/src/modules/game/game.service.ts` | 227 | JS/TS | `getSession` |
| `apps/backend/src/modules/game/game.service.ts` | 252 | JS/TS | `generateSlotResult` |
| `apps/backend/src/modules/game/game.service.ts` | 276 | JS/TS | `generateSymbols` |
| `apps/backend/src/modules/game/game.service.ts` | 355 | JS/TS | `generatePokerResult` |
| `apps/backend/src/modules/game/game.service.ts` | 398 | JS/TS | `generatePokerCards` |
| `apps/backend/src/modules/game/game.service.ts` | 520 | JS/TS | `getSportsEvents` |
| `apps/backend/src/modules/game/game.service.ts` | 625 | JS/TS | `getGameTypes` |
| `apps/backend/src/modules/game/game.service.ts` | 639 | JS/TS | `getGameStatistics` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 18 | JS/TS | `constructor` |
| `apps/backend/src/modules/game/gamefi.controller.ts` | 105 | JS/TS | `getActiveTournaments` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 82 | JS/TS | `constructor` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 166 | JS/TS | `generateProvablyFairResult` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 217 | JS/TS | `updateUserStats` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 335 | JS/TS | `getUserNFTs` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 388 | JS/TS | `joinTournament` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 423 | JS/TS | `getTournamentLeaderboard` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 448 | JS/TS | `getActiveTournaments` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 485 | JS/TS | `getGlobalLeaderboard` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 510 | JS/TS | `getUserStats` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 551 | JS/TS | `getGameHistory` |
| `apps/backend/src/modules/game/gamefi.service.ts` | 578 | JS/TS | `verifyProvablyFair` |
| `apps/backend/src/modules/governance/governance.controller.ts` | 18 | JS/TS | `constructor` |
| `apps/backend/src/modules/governance/governance.service.ts` | 40 | JS/TS | `constructor` |
| `apps/backend/src/modules/governance/governance.service.ts` | 182 | JS/TS | `getProposal` |
| `apps/backend/src/modules/governance/governance.service.ts` | 292 | JS/TS | `getVotingPower` |
| `apps/backend/src/modules/governance/governance.service.ts` | 306 | JS/TS | `executeProposal` |
| `apps/backend/src/modules/governance/governance.service.ts` | 337 | JS/TS | `cancelProposal` |
| `apps/backend/src/modules/rewards/rewards.controller.ts` | 18 | JS/TS | `constructor` |
| `apps/backend/src/modules/rewards/rewards.service.ts` | 133 | JS/TS | `getUserRewards` |
| `apps/backend/src/modules/rewards/rewards.service.ts` | 150 | JS/TS | `checkEligibility` |
| `apps/backend/src/modules/social/social.controller.ts` | 6 | JS/TS | `constructor` |
| `apps/backend/src/modules/social/social.service.ts` | 6 | JS/TS | `constructor` |
| `apps/backend/src/modules/social/social.service.ts` | 8 | JS/TS | `createProfile` |
| `apps/backend/src/modules/social/social.service.ts` | 16 | JS/TS | `getProfile` |
| `apps/backend/src/modules/social/social.service.ts` | 22 | JS/TS | `updateProfile` |
| `apps/backend/src/modules/social/social.service.ts` | 29 | JS/TS | `followUser` |
| `apps/backend/src/modules/social/social.service.ts` | 59 | JS/TS | `unfollowUser` |
| `apps/backend/src/modules/social/social.service.ts` | 85 | JS/TS | `getFollowers` |
| `apps/backend/src/modules/social/social.service.ts` | 94 | JS/TS | `getFollowing` |
| `apps/backend/src/modules/social/social.service.ts` | 103 | JS/TS | `createPost` |
| `apps/backend/src/modules/social/social.service.ts` | 109 | JS/TS | `getFeed` |
| `apps/backend/src/modules/social/social.service.ts` | 127 | JS/TS | `likePost` |
| `apps/backend/src/modules/social/social.service.ts` | 134 | JS/TS | `addExperience` |
| `apps/backend/src/modules/social/social.service.ts` | 155 | JS/TS | `unlockAchievement` |
| `apps/backend/src/modules/social/social.service.ts` | 181 | JS/TS | `getAchievements` |
| `apps/backend/src/modules/social/social.service.ts` | 188 | JS/TS | `getLeaderboard` |
| `apps/backend/src/prisma.service.ts` | 18 | JS/TS | `onModuleInit` |
| `apps/backend/src/prisma.service.ts` | 22 | JS/TS | `onModuleDestroy` |
| `apps/frontend-miniapp/src/screens/BridgeScreen.tsx` | 51 | JS/TS | `getQuote` |
| `apps/frontend-miniapp/src/screens/BridgeScreen.tsx` | 76 | JS/TS | `initiateBridge` |
| `apps/frontend-miniapp/src/screens/BridgeScreen.tsx` | 92 | JS/TS | `addLiquidity` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 51 | JS/TS | `loadUserStakes` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 65 | JS/TS | `loadStakeAnalytics` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 76 | JS/TS | `getSwapQuote` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 105 | JS/TS | `executeSwap` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 135 | JS/TS | `createStake` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 172 | JS/TS | `claimRewards` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 191 | JS/TS | `unstakeTokens` |
| `apps/frontend-miniapp/src/screens/DeFiScreen.tsx` | 211 | JS/TS | `formatEther` |
| `apps/frontend-miniapp/src/screens/FinTechScreen.tsx` | 34 | JS/TS | `generatePromptPayQR` |
| `apps/frontend-miniapp/src/screens/FinTechScreen.tsx` | 74 | JS/TS | `checkPaymentStatus` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 60 | JS/TS | `loadStats` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 70 | JS/TS | `loadSportsEvents` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 80 | JS/TS | `playSlots` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 111 | JS/TS | `playPoker` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 135 | JS/TS | `playRoulette` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 162 | JS/TS | `placeSportsBet` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 191 | JS/TS | `renderGameMenu` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 261 | JS/TS | `renderSlotsGame` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 303 | JS/TS | `renderPokerGame` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 346 | JS/TS | `renderRouletteGame` |
| `apps/frontend-miniapp/src/screens/GameScreen.tsx` | 425 | JS/TS | `renderSportsBetting` |
| `apps/frontend-miniapp/src/screens/GovernanceScreen.tsx` | 47 | JS/TS | `loadProposals` |
| `apps/frontend-miniapp/src/screens/GovernanceScreen.tsx` | 59 | JS/TS | `loadVotingPower` |
| `apps/frontend-miniapp/src/screens/GovernanceScreen.tsx` | 72 | JS/TS | `createProposal` |
| `apps/frontend-miniapp/src/screens/GovernanceScreen.tsx` | 113 | JS/TS | `vote` |
| `apps/frontend-miniapp/src/screens/GovernanceScreen.tsx` | 140 | JS/TS | `formatVotes` |
| `apps/frontend-miniapp/src/screens/GovernanceScreen.tsx` | 145 | JS/TS | `renderProposal` |
| `apps/frontend-miniapp/src/screens/RewardScreen.tsx` | 17 | JS/TS | `handleClaimDaily` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 93 | JS/TS | `loadStats` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 103 | JS/TS | `loadNFTs` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 113 | JS/TS | `loadTournaments` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 123 | JS/TS | `loadLeaderboard` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 133 | JS/TS | `spinSlots` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 177 | JS/TS | `joinTournament` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 197 | JS/TS | `renderSlotsTab` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 359 | JS/TS | `renderTournamentsTab` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 404 | JS/TS | `renderNFTsTab` |
| `apps/frontend-miniapp/src/screens/SlotMachineScreen.tsx` | 435 | JS/TS | `renderLeaderboardTab` |
| `apps/frontend-miniapp/src/screens/SocialScreen.tsx` | 16 | JS/TS | `createPost` |
| `apps/frontend-miniapp/src/screens/WorldIDScreen.tsx` | 23 | JS/TS | `handleVerifyWorldID` |
| `install-ubuntu.sh` | 35 | Shell | `print_success` |
| `install-ubuntu.sh` | 39 | Shell | `print_error` |
| `install-ubuntu.sh` | 43 | Shell | `print_warning` |
| `install-ubuntu.sh` | 47 | Shell | `print_info` |
| `install-ubuntu.sh` | 52 | Shell | `check_root` |
| `install-ubuntu.sh` | 61 | Shell | `detect_ubuntu_version` |
| `install.sh` | 37 | Shell | `print_success` |
| `install.sh` | 41 | Shell | `print_error` |
| `install.sh` | 45 | Shell | `print_warning` |
| `install.sh` | 49 | Shell | `print_info` |
| `packages/contracts/scripts/deploy-multichain.ts` | 15 | JS/TS | `main` |
| `packages/contracts/scripts/deploy.ts` | 15 | JS/TS | `main` |

### Automation pipeline files

- `.github/workflows/omega-auto-fix.yml`
- `.github/workflows/omega-auto-test.yml`
- `.github/workflows/omega-autoscale.yml`
- `.github/workflows/omega-heartbeat.yml`
- `.github/workflows/omega-metrics.yml`
- `.github/workflows/omega-onchain.yml`
- `.github/workflows/omega-sync.yml`
- `.github/workflows/omega-threat.yml`
- `INSTALLER_OS_REQUIREMENTS.md`
- `ZeaZDev-Omega_AuditPack/.github/workflows/smoke_test.yml`
- `docker-compose.yml`
- `docs/WORKFLOW_CICD_2026.md`
- `install-ubuntu.sh`
- `install.sh`
- `packages/contracts/scripts/deploy-multichain.ts`
- `packages/contracts/scripts/deploy.ts`

## `ZeaZDev/zeazdev-repo`

Status: **cloned**; commit `5b73116fbde4`.

### API/endpoints

| File | Line | Endpoint/expression |
|---|---:|---|
| `zeaz/app/main.py` | 67 | `/healthz` |
| `zeaz/app/main.py` | 72 | `/auth/token` |
| `zeaz/app/main.py` | 77 | `/wallet/balance/{user_id}/{currency}` |
| `zeaz/app/main.py` | 87 | `/wallet/credit` |
| `zeaz/app/main.py` | 102 | `/wallet/debit` |
| `zeaz/app/main.py` | 117 | `/webhooks/stripe` |
| `zeaz/app/main.py` | 130 | `/tiktok/feed-product-form/generate` |
| `zeaz/app/main.py` | 145 | `/tiktok/video/generate` |
| `zeaz/app/main.py` | 158 | `/tiktok/shop-affiliate/upload` |
| `zeaz/app/main.py` | 171 | `/tiktok/jobs` |
| `zeaz/app/main.py` | 177 | `/tiktok/jobs/{job_id}` |
| `zeaz/app/main.py` | 186 | `/admin/control-panel` |
| `zeaz/app/main.py` | 203 | `/user/control-panel` |

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `zeaz/app/auth.py` | 14 | Python | `mint_jwt` |
| `zeaz/app/auth.py` | 23 | Python | `decode_jwt` |
| `zeaz/app/auth.py` | 33 | Python | `current_principal` |
| `zeaz/app/auth.py` | 39 | Python | `require_role` |
| `zeaz/app/auth.py` | 40 | Python | `_dependency` |
| `zeaz/app/chain_worker.py` | 13 | Python | `process_confirmed_tx` |
| `zeaz/app/chain_worker.py` | 26 | Python | `process_candidates` |
| `zeaz/app/chain_worker.py` | 48 | Python | `run` |
| `zeaz/app/db.py` | 45 | Python | `database_url` |
| `zeaz/app/db.py` | 53 | Python | `db_connection` |
| `zeaz/app/ledger.py` | 21 | Python | `_normalize_currency` |
| `zeaz/app/ledger.py` | 28 | Python | `_write_audit` |
| `zeaz/app/ledger.py` | 33 | Python | `append_entry` |
| `zeaz/app/ledger.py` | 63 | Python | `user_balance` |
| `zeaz/app/main.py` | 30 | Python | `normalize_currency` |
| `zeaz/app/main.py` | 68 | Python | `healthz` |
| `zeaz/app/main.py` | 73 | Python | `create_token` |
| `zeaz/app/main.py` | 78 | Python | `wallet_balance` |
| `zeaz/app/main.py` | 88 | Python | `wallet_credit` |
| `zeaz/app/main.py` | 103 | Python | `wallet_debit` |
| `zeaz/app/main.py` | 118 | Python | `stripe_webhook` |
| `zeaz/app/main.py` | 131 | Python | `generate_feed_product_form` |
| `zeaz/app/main.py` | 146 | Python | `generate_product_video` |
| `zeaz/app/main.py` | 159 | Python | `upload_tiktok_shop_affiliate` |
| `zeaz/app/main.py` | 172 | Python | `get_tiktok_jobs` |
| `zeaz/app/main.py` | 178 | Python | `get_tiktok_job` |
| `zeaz/app/main.py` | 187 | Python | `admin_control_panel` |
| `zeaz/app/main.py` | 204 | Python | `user_control_panel` |
| `zeaz/app/migrate.py` | 6 | Python | `migrate` |
| `zeaz/app/stripe_webhook.py` | 14 | Python | `_extract_v1` |
| `zeaz/app/stripe_webhook.py` | 24 | Python | `verify_signature` |
| `zeaz/app/stripe_webhook.py` | 34 | Python | `process_payment_intent_succeeded` |
| `zeaz/app/tiktok_affiliate.py` | 12 | Python | `_utc_now` |
| `zeaz/app/tiktok_affiliate.py` | 16 | Python | `_create_job` |
| `zeaz/app/tiktok_affiliate.py` | 31 | Python | `create_feed_product_form` |
| `zeaz/app/tiktok_affiliate.py` | 43 | Python | `create_product_video` |
| `zeaz/app/tiktok_affiliate.py` | 57 | Python | `upload_to_tiktok_affiliate` |
| `zeaz/app/tiktok_affiliate.py` | 67 | Python | `get_job` |
| `zeaz/app/tiktok_affiliate.py` | 71 | Python | `list_jobs` |
| `zeaz/app/wallet.py` | 6 | Python | `credit_wallet` |
| `zeaz/app/wallet.py` | 29 | Python | `debit_wallet` |
| `zeaz/preview_project.py` | 14 | Python | `generate_preview` |
| `zeaz/preview_project.py` | 54 | Python | `main` |
| `zeaz/ui/admin/src/dataProvider.ts` | 5 | JS/TS | `httpClient` |
| `zeaz/ui/admin/src/dataProvider.ts` | 14 | JS/TS | `oneItemList` |
| `zeaz/ui/admin/src/main.tsx` | 16 | JS/TS | `TiktokFeedFormCreate` |
| `zeaz/ui/admin/src/main.tsx` | 30 | JS/TS | `TiktokVideoCreate` |
| `zeaz/ui/admin/src/main.tsx` | 40 | JS/TS | `TiktokUploadCreate` |
| `zeaz/ui/admin/src/main.tsx` | 50 | JS/TS | `TiktokJobsList` |
| `zeaz/ui/admin/src/main.tsx` | 61 | JS/TS | `AdminPanelList` |
| `zeaz/ui/admin/src/main.tsx` | 71 | JS/TS | `UserPanelList` |

### Automation pipeline files

- `zeaz/install.sh`

## `ZeaZDev/zeaztools`

Status: **cloned**; commit `7a971921b431`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

None detected or repository blocked.

## `ZeaZDev/zlms-prod`

Status: **cloned**; commit `f26c75d7083b`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

| File | Line | Language | Function/method |
|---|---:|---|---|
| `app/Admin/Default.aspx.cs` | 14 | C# | `Page_Load` |
| `app/CConnect.cs` | 27 | C# | `CConnect` |
| `app/CConnect.cs` | 38 | C# | `connOpen` |
| `app/CConnect.cs` | 45 | C# | `checkConn` |
| `app/CConnect.cs` | 50 | C# | `connClose` |
| `app/CConnect.cs` | 55 | C# | `drClose` |
| `app/CConnect.cs` | 60 | C# | `sqlCmd` |
| `app/CConnect.cs` | 77 | C# | `sqlCmdCheck` |
| `app/CConnect.cs` | 95 | C# | `sqlCmd` |
| `app/CConnect.cs` | 111 | C# | `sqlCmdText` |
| `app/CConnect.cs` | 116 | C# | `sqlCmdAddParam` |
| `app/CConnect.cs` | 121 | C# | `sqlCmdReturn` |
| `app/CConnect.cs` | 140 | C# | `sqlReader` |
| `app/CConnect.cs` | 147 | C# | `sqlReader` |
| `app/CConnect.cs` | 153 | C# | `rsMoveNext` |
| `app/CConnect.cs` | 158 | C# | `Recordset` |
| `app/CConnect.cs` | 163 | C# | `Recordset` |
| `app/CConnect.cs` | 168 | C# | `daFill` |
| `app/CConnect.cs` | 188 | C# | `daDeleteCmd` |
| `app/CConnect.cs` | 205 | C# | `daFill` |
| `app/CConnect.cs` | 225 | C# | `daUpdate` |
| `app/CConnect.cs` | 231 | C# | `daUpdate` |
| `app/CConnect.cs` | 237 | C# | `daUpdate` |
| `app/CConnect.cs` | 243 | C# | `daUpdate` |
| `app/CConnect.cs` | 249 | C# | `getMd5Hash` |
| `app/CConnect.cs` | 258 | C# | `verifyMd5Hash` |
| `app/Certificate/Certificate_adjust.aspx` | 16 | C# | `reportDesigner_EndCallback` |
| `app/Certificate/Certificate_adjust.aspx.cs` | 25 | C# | `Page_Load` |
| `app/Certificate/Certificate_adjust.aspx.cs` | 43 | C# | `reportDesigner_SaveReportLayout` |
| `app/Certificate/Certificate_adjust.aspx.cs` | 47 | C# | `btExportReport_Click` |
| `app/Certificate/Certificate_edit.aspx.cs` | 23 | C# | `Page_Load` |
| `app/Certificate/Certificate_edit.aspx.cs` | 36 | C# | `ADD_COM_CODE` |
| `app/Certificate/Certificate_edit.aspx.cs` | 55 | C# | `coursename` |
| `app/Certificate/Certificate_edit.aspx.cs` | 77 | C# | `Button1_Click` |
| `app/Certificate/Certificatedetail.aspx.cs` | 21 | C# | `Page_Load` |
| `app/Certificate/Certificatedetail_user.aspx.cs` | 21 | C# | `Page_Load` |
| `app/Certificate/Certificatedetail_user.aspx.cs` | 33 | C# | `mapmonth` |
| `app/Certificate/Certificatedetail_user.aspx.cs` | 66 | C# | `GetReportByName` |
| `app/Certificate/Certificatedetail_user.aspx.cs` | 90 | C# | `coursename` |
| `app/Certificate/Default.aspx` | 7 | C# | `openModal` |
| `app/Certificate/Default.aspx.cs` | 27 | C# | `ShowMessage` |
| `app/Certificate/Default.aspx.cs` | 32 | C# | `Page_Load` |
| `app/Certificate/Default.aspx.cs` | 39 | C# | `ADD_COM_CODE` |
| `app/Certificate/Default.aspx.cs` | 58 | C# | `Button1_Click` |
| `app/Certificate/Default.aspx.cs` | 66 | C# | `delete_course` |
| `app/Certificate/Default.aspx.cs` | 85 | C# | `Button2_Click` |
| `app/Certificate/Default.aspx.cs` | 102 | C# | `SetSession` |
| `app/Certificate/Default.aspx.cs` | 107 | C# | `bnAdduser_Click` |
| `app/Certificate/Default.aspx.cs` | 124 | C# | `clearAddnew` |
| `app/Certificate/Default.aspx.cs` | 130 | C# | `addUser` |
| `app/Certificate/Issue.aspx.cs` | 25 | C# | `Page_Load` |
| `app/Certificate/Issue.aspx.cs` | 33 | C# | `ADD_COM_CODE` |
| `app/Certificate/Issue.aspx.cs` | 52 | C# | `ADD_COM_CODE1` |
| `app/Certificate/Issue.aspx.cs` | 72 | C# | `SetSession` |
| `app/Certificate/Issue.aspx.cs` | 77 | C# | `Button1_Click` |
| `app/Certificate/Issue.aspx.cs` | 85 | C# | `bnAdduser_Click` |
| `app/Certificate/Issue.aspx.cs` | 93 | C# | `addUser` |
| `app/Certificate/ViewCertificate.aspx.cs` | 22 | C# | `ShowMessage` |
| `app/Certificate/ViewCertificate.aspx.cs` | 27 | C# | `Page_Load` |
| `app/Certificate/ViewCertificate.aspx.cs` | 51 | C# | `coursename` |
| `app/Certificate/ViewCertificate.aspx.cs` | 71 | C# | `Button1_Click` |
| `app/Certificate/WebForm1.aspx` | 16 | C# | `reportDesigner_EndCallback` |
| `app/Certificate/WebForm1.aspx.cs` | 25 | C# | `Page_Load` |
| `app/Certificate/WebForm1.aspx.cs` | 32 | C# | `reportDesigner_SaveReportLayout` |
| `app/Certificate/WebForm1.aspx.cs` | 36 | C# | `btExportReport_Click` |
| `app/Certificate/XtraReport1.resx.cs` | 19 | C# | `XtraReport1` |
| `app/Certificate/XtraReport1.resx.cs` | 24 | C# | `Dispose` |
| `app/Certificate/XtraReport1.resx.cs` | 31 | C# | `InitializeComponent` |
| `app/Checkservice.aspx.cs` | 22 | C# | `Page_Load` |
| `app/Checkservice.aspx.cs` | 26 | C# | `Button1_Click` |
| `app/Course/Class.aspx` | 6 | C# | `openModal` |
| `app/Course/Class.aspx.cs` | 31 | C# | `ShowMessage` |
| `app/Course/Class.aspx.cs` | 36 | C# | `Page_Load` |
| `app/Course/Class.aspx.cs` | 59 | C# | `coursename` |
| `app/Course/Class.aspx.cs` | 82 | C# | `Button1_Click` |
| `app/Course/Class.aspx.cs` | 90 | C# | `delete_course` |
| `app/Course/Class.aspx.cs` | 109 | C# | `Button2_Click` |
| `app/Course/Class.aspx.cs` | 126 | C# | `SetSession` |
| `app/Course/Class.aspx.cs` | 131 | C# | `bnAdduser_Click` |
| `app/Course/Class.aspx.cs` | 148 | C# | `clearAddnew` |
| `app/Course/Class.aspx.cs` | 154 | C# | `addUser` |
| `app/Course/ClassItemdetail.aspx` | 6 | C# | `ReGen` |
| `app/Course/ClassItemdetail.aspx.cs` | 28 | C# | `ShowMessage` |
| `app/Course/ClassItemdetail.aspx.cs` | 33 | C# | `Page_Load` |
| `app/Course/ClassItemdetail.aspx.cs` | 47 | C# | `coursename` |
| `app/Course/ClassItemdetail.aspx.cs` | 77 | C# | `renderdata` |
| `app/Course/ClassItemdetail.aspx.cs` | 134 | C# | `renderdata1` |
| `app/Course/ClassItemdetail.aspx.cs` | 195 | C# | `ReGenToken` |
| `app/Course/ClassItemdetail.aspx.cs` | 214 | C# | `getquestion` |
| `app/Course/ClassItemdetail.aspx.cs` | 240 | C# | `checkanswer` |
| `app/Course/ClassItemdetail.aspx.cs` | 263 | C# | `addUser` |
| `app/Course/ClassItemdetail.aspx.cs` | 270 | C# | `CConnect` |
| `app/Course/ClassItemdetail.aspx.cs` | 273 | C# | `Button1_Click` |
| `app/Course/Coursecalenda.aspx.cs` | 18 | C# | `Page_Load` |
| `app/Course/Coursecalenda.aspx.cs` | 31 | C# | `SetupMappings` |
| `app/Course/Coursedetail.aspx` | 6 | C# | `openModal` |
| `app/Course/Coursedetail.aspx.cs` | 29 | C# | `ShowMessage` |
| `app/Course/Coursedetail.aspx.cs` | 34 | C# | `Page_Load` |
| `app/Course/Coursedetail.aspx.cs` | 58 | C# | `coursename` |
| `app/Course/Coursedetail.aspx.cs` | 79 | C# | `Button1_Click` |
| `app/Course/Coursedetail.aspx.cs` | 87 | C# | `delete_course` |
| `app/Course/Coursedetail.aspx.cs` | 106 | C# | `Button2_Click` |
| `app/Course/Coursedetail.aspx.cs` | 123 | C# | `SetSession` |
| `app/Course/Coursedetail.aspx.cs` | 128 | C# | `bnAdduser_Click` |
| `app/Course/Coursedetail.aspx.cs` | 145 | C# | `clearAddnew` |
| `app/Course/Coursedetail.aspx.cs` | 151 | C# | `addUser` |
| `app/Course/Default.aspx` | 6 | C# | `openModal` |
| `app/Course/Default.aspx` | 44 | C# | `getDate` |
| `app/Course/Default.aspx.cs` | 28 | C# | `ShowMessage` |
| `app/Course/Default.aspx.cs` | 33 | C# | `Page_Load` |
| `app/Course/Default.aspx.cs` | 38 | C# | `Button1_Click` |
| `app/Course/Default.aspx.cs` | 46 | C# | `Button3_Click` |
| `app/Course/Default.aspx.cs` | 54 | C# | `Button4_Click` |
| `app/Course/Default.aspx.cs` | 62 | C# | `delete_course` |
| `app/Course/Default.aspx.cs` | 81 | C# | `Button2_Click` |
| `app/Course/Default.aspx.cs` | 98 | C# | `SetSession` |
| `app/Course/Default.aspx.cs` | 103 | C# | `bnAdduser_Click` |
| `app/Course/Default.aspx.cs` | 120 | C# | `clearAddnew` |
| `app/Course/Default.aspx.cs` | 126 | C# | `addUser` |
| `app/Course/Question.aspx` | 6 | C# | `openModal` |
| `app/Course/Question.aspx.cs` | 28 | C# | `Page_Load` |
| `app/Course/Question.aspx.cs` | 42 | C# | `ADD_COM_CODE` |
| `app/Course/Question.aspx.cs` | 63 | C# | `coursename` |
| `app/Course/Question.aspx.cs` | 94 | C# | `Button1_Click` |
| `app/Course/Viewresult.aspx.cs` | 23 | C# | `ShowMessage` |
| `app/Course/Viewresult.aspx.cs` | 28 | C# | `Page_Load` |
| `app/Course/Viewresult.aspx.cs` | 52 | C# | `coursename` |
| `app/Course/class_edit.aspx.cs` | 22 | C# | `Page_Load` |
| `app/Course/class_edit.aspx.cs` | 35 | C# | `coursename` |
| `app/Course/class_edit.aspx.cs` | 56 | C# | `Button1_Click` |
| `app/Course/classitem_edit.aspx.cs` | 23 | C# | `Page_Load` |
| `app/Course/classitem_edit.aspx.cs` | 36 | C# | `coursename` |
| `app/Course/classitem_edit.aspx.cs` | 58 | C# | `Button1_Click` |
| `app/Course/classitemupload.aspx` | 11 | C# | `ReGen` |
| `app/Course/classitemupload.aspx.cs` | 25 | C# | `Page_Load` |
| `app/Course/classitemupload.aspx.cs` | 71 | C# | `addUser` |
| `app/Course/classitemupload.aspx.cs` | 86 | C# | `renderdata` |
| `app/Course/classitemupload.aspx.cs` | 127 | C# | `coursename` |
| `app/Course/classitemupload.aspx.cs` | 148 | C# | `ReGenToken` |
| `app/Course/course_edit.aspx.cs` | 25 | C# | `Page_Load` |
| `app/Course/course_edit.aspx.cs` | 38 | C# | `coursename` |
| `app/Course/course_edit.aspx.cs` | 60 | C# | `Button1_Click` |
| `app/Course/period.aspx` | 6 | C# | `openModal` |
| `app/Course/period.aspx.cs` | 35 | C# | `ShowMessage` |
| `app/Course/period.aspx.cs` | 40 | C# | `Page_Load` |
| `app/Course/period.aspx.cs` | 65 | C# | `ADD_COM_CODE` |
| `app/Course/period.aspx.cs` | 84 | C# | `Button5_Click` |
| `app/Course/period.aspx.cs` | 92 | C# | `coursename` |
| `app/Course/period.aspx.cs` | 117 | C# | `Button1_Click` |
| `app/Course/period.aspx.cs` | 125 | C# | `Button3_Click` |
| `app/Course/period.aspx.cs` | 133 | C# | `Button4_Click` |
| `app/Course/period.aspx.cs` | 141 | C# | `gvDP_RowDataBound` |
| `app/Course/period.aspx.cs` | 152 | C# | `delete_course` |
| `app/Course/period.aspx.cs` | 171 | C# | `Button2_Click` |
| `app/Course/period.aspx.cs` | 188 | C# | `SetSession` |
| `app/Course/period.aspx.cs` | 193 | C# | `bnAdduser_Click` |
| `app/Course/period.aspx.cs` | 210 | C# | `clearAddnew` |
| `app/Course/period.aspx.cs` | 216 | C# | `search_actiontype` |
| `app/Course/period.aspx.cs` | 237 | C# | `addUser` |
| `app/Course/period_edit.aspx.cs` | 24 | C# | `Page_Load` |
| `app/Course/period_edit.aspx.cs` | 37 | C# | `coursename` |
| `app/Course/period_edit.aspx.cs` | 60 | C# | `Button1_Click` |
| `app/Course/setusercourse.aspx.cs` | 24 | C# | `ShowMessage` |
| `app/Course/setusercourse.aspx.cs` | 29 | C# | `Page_Load` |
| `app/Course/setusercourse.aspx.cs` | 51 | C# | `ADD_COM_CODE` |
| `app/Course/setusercourse.aspx.cs` | 70 | C# | `Button1_Click` |
| `app/Course/setusercourse.aspx.cs` | 91 | C# | `delete_course` |
| `app/Course/setusercourse.aspx.cs` | 110 | C# | `Button2_Click` |
| `app/Course/setusercourse.aspx.cs` | 127 | C# | `SetSession` |
| `app/Course_user/Course_enroll.aspx.cs` | 18 | C# | `Page_Load` |
| `app/Course_user/Course_enroll.aspx.cs` | 33 | C# | `Button1_Click` |
| `app/Course_user/Course_enroll.aspx.cs` | 41 | C# | `addUser` |
| `app/Course_user/Course_enroll.aspx.cs` | 43 | C# | `CConnect` |
| `app/Course_user/Coursedetail.aspx.cs` | 23 | C# | `ShowMessage` |
| `app/Course_user/Coursedetail.aspx.cs` | 28 | C# | `Page_Load` |
| `app/Course_user/Coursedetail.aspx.cs` | 52 | C# | `coursename` |
| `app/Course_user/Default.aspx.cs` | 20 | C# | `ShowMessage` |
| `app/Course_user/Default.aspx.cs` | 25 | C# | `Page_Load` |
| `app/Course_user/Default.aspx.cs` | 40 | C# | `Button3_Click` |
| `app/Course_user/Default.aspx.cs` | 48 | C# | `Button1_Click` |
| `app/Course_user/Default.aspx.cs` | 59 | C# | `check_enroll` |
| `app/Course_user/Default.aspx.cs` | 81 | C# | `addUser` |
| `app/Course_user/Default.aspx.cs` | 83 | C# | `CConnect` |
| `app/Course_user/View_pollresult.aspx.cs` | 16 | C# | `ShowMessage` |
| `app/Course_user/View_pollresult.aspx.cs` | 21 | C# | `Page_Load` |
| `app/Course_user/View_pollresult.aspx.cs` | 26 | C# | `renderdata1` |
| `app/Course_user/_class.aspx.cs` | 25 | C# | `ShowMessage` |
| `app/Course_user/_class.aspx.cs` | 30 | C# | `Page_Load` |
| `app/Course_user/_class.aspx.cs` | 53 | C# | `coursename` |
| `app/Course_user/classitemdetail.aspx.cs` | 28 | C# | `ShowMessage` |
| `app/Course_user/classitemdetail.aspx.cs` | 33 | C# | `Page_Load` |
| `app/Course_user/classitemdetail.aspx.cs` | 47 | C# | `coursename` |
| `app/Course_user/classitemdetail.aspx.cs` | 77 | C# | `renderdata` |
| `app/Course_user/classitemdetail.aspx.cs` | 134 | C# | `checkquestion` |
| `app/Course_user/classitemdetail.aspx.cs` | 184 | C# | `showrdetailresult` |
| `app/Course_user/classitemdetail.aspx.cs` | 227 | C# | `showrdansresult` |
| `app/Course_user/classitemdetail.aspx.cs` | 250 | C# | `renderdata1` |
| `app/Course_user/classitemdetail.aspx.cs` | 329 | C# | `ReGenToken` |
| `app/Course_user/classitemdetail.aspx.cs` | 348 | C# | `getquestion` |
| `app/Course_user/classitemdetail.aspx.cs` | 409 | C# | `addUser` |
| `app/Course_user/classitemdetail.aspx.cs` | 411 | C# | `CConnect` |
| `app/Course_user/classitemdetail.aspx.cs` | 414 | C# | `addUser1` |
| `app/Course_user/classitemdetail.aspx.cs` | 421 | C# | `CConnect` |
| `app/Course_user/classitemdetail.aspx.cs` | 424 | C# | `fillquestion` |
| `app/Course_user/classitemdetail.aspx.cs` | 444 | C# | `Button1_Click` |
| `app/Course_user/period.aspx.cs` | 31 | C# | `ShowMessage` |
| `app/Course_user/period.aspx.cs` | 36 | C# | `Page_Load` |
| `app/Course_user/period.aspx.cs` | 61 | C# | `checkpoll` |
| `app/Course_user/period.aspx.cs` | 84 | C# | `coursename` |
| `app/Course_user/period.aspx.cs` | 109 | C# | `Button1_Click` |
| `app/Course_user/period.aspx.cs` | 117 | C# | `addPoll` |
| `app/Course_user/period.aspx.cs` | 119 | C# | `CConnect` |
| `app/Course_user/poll.aspx.cs` | 27 | C# | `ShowMessage` |
| `app/Course_user/poll.aspx.cs` | 32 | C# | `Page_Load` |
| `app/Course_user/poll.aspx.cs` | 37 | C# | `renderdata1` |
| `app/Course_user/poll.aspx.cs` | 92 | C# | `renderdata3` |
| `app/Course_user/poll.aspx.cs` | 150 | C# | `renderdata2` |
| `app/Course_user/poll.aspx.cs` | 178 | C# | `addUser` |
| `app/Course_user/poll.aspx.cs` | 193 | C# | `Button1_Click` |
| `app/Course_user/poll1.aspx.cs` | 31 | C# | `ShowMessage` |
| `app/Course_user/poll1.aspx.cs` | 36 | C# | `Page_Load` |
| `app/Course_user/poll1.aspx.cs` | 41 | C# | `renderdata1` |
| `app/Course_user/poll1.aspx.cs` | 96 | C# | `addUser` |
| `app/Course_user/poll1.aspx.cs` | 111 | C# | `Button1_Click` |
| `app/Course_user/profile.aspx.cs` | 54 | C# | `Page_Load` |
| `app/Default.aspx.cs` | 14 | C# | `Page_Load` |
| `app/DropzoneJs_scripts/dropzone.js` | 12 | JS/TS | `require` |
| `app/DropzoneJs_scripts/dropzone.js` | 72 | JS/TS | `Emitter` |
| `app/DropzoneJs_scripts/dropzone.js` | 84 | JS/TS | `mixin` |
| `app/DropzoneJs_scripts/dropzone.js` | 122 | JS/TS | `on` |
| `app/DropzoneJs_scripts/dropzone.js` | 543 | JS/TS | `Dropzone` |
| `app/DropzoneJs_scripts/dropzone.js` | 1837 | JS/TS | `define` |
| `app/Dummy.aspx.cs` | 12 | C# | `Page_Load` |
| `app/MissingExternalDependencies.Stubs.cs` | 19 | C# | `ExpandToLevel` |
| `app/MissingExternalDependencies.Stubs.cs` | 27 | C# | `FromFile` |
| `app/MissingExternalDependencies.Stubs.cs` | 28 | C# | `Dispose` |
| `app/MissingExternalDependencies.Stubs.cs` | 29 | C# | `Dispose` |
| `app/MissingExternalDependencies.Stubs.cs` | 30 | C# | `EndInit` |
| `app/MissingExternalDependencies.Stubs.cs` | 46 | C# | `OpenReport` |
| `app/MissingExternalDependencies.Stubs.cs` | 103 | C# | `AddXY` |
| `app/Police.Master.cs` | 18 | C# | `Page_Load` |
| `app/Police.Master.cs` | 32 | C# | `getuser` |
| `app/Police_noform.Master.cs` | 18 | C# | `Page_Load` |
| `app/Police_noform.Master.cs` | 22 | C# | `getuser` |
| `app/Police_service/GET_NAME.cs` | 111 | C# | `RaisePropertyChanged` |
| `app/Police_service/IPolice_web.cs` | 17 | C# | `Check_user` |
| `app/Police_service/Police_webClient.cs` | 18 | C# | `Police_webClient` |
| `app/Police_service/Police_webClient.cs` | 22 | C# | `Police_webClient` |
| `app/Police_service/Police_webClient.cs` | 27 | C# | `Police_webClient` |
| `app/Police_service/Police_webClient.cs` | 32 | C# | `Police_webClient` |
| `app/Police_service/Police_webClient.cs` | 37 | C# | `Police_webClient` |
| `app/Police_service/Police_webClient.cs` | 42 | C# | `Check_user` |
| `app/Police_user.Master.cs` | 20 | C# | `Page_Load` |
| `app/Police_user.Master.cs` | 25 | C# | `getuser` |
| `app/QA/Activities.aspx.cs` | 27 | C# | `ShowMessage` |
| `app/QA/Activities.aspx.cs` | 32 | C# | `Page_Load` |
| `app/QA/Activities.aspx.cs` | 54 | C# | `ADD_COM_CODE` |
| `app/QA/Activities.aspx.cs` | 73 | C# | `Button4_Click` |
| `app/QA/Activities.aspx.cs` | 81 | C# | `Button1_Click` |
| `app/QA/Activities.aspx.cs` | 89 | C# | `delete_course` |
| `app/QA/Activities.aspx.cs` | 108 | C# | `Button2_Click` |
| `app/QA/Activities.aspx.cs` | 125 | C# | `SetSession` |
| `app/QA/Activities.aspx.cs` | 130 | C# | `bnAdduser_Click` |
| `app/QA/Activities.aspx.cs` | 147 | C# | `clearAddnew` |
| `app/QA/Activities.aspx.cs` | 152 | C# | `addUser` |
| `app/QA/Activityupload.aspx` | 11 | C# | `ReGen` |
| `app/QA/Activityupload.aspx.cs` | 23 | C# | `Page_Load` |
| `app/QA/Activityupload.aspx.cs` | 68 | C# | `addUser` |
| `app/QA/Activityupload.aspx.cs` | 83 | C# | `renderdata` |
| `app/QA/Activityupload.aspx.cs` | 124 | C# | `ReGenToken` |
| `app/QA/Default.aspx.cs` | 25 | C# | `ShowMessage` |
| `app/QA/Default.aspx.cs` | 30 | C# | `Page_Load` |
| `app/QA/Default.aspx.cs` | 35 | C# | `Button1_Click` |
| `app/QA/Default.aspx.cs` | 43 | C# | `Button3_Click` |
| `app/QA/Default.aspx.cs` | 51 | C# | `Button4_Click` |
| `app/QA/Default.aspx.cs` | 59 | C# | `delete_course` |
| `app/QA/Default.aspx.cs` | 78 | C# | `Button2_Click` |
| `app/QA/Default.aspx.cs` | 95 | C# | `SetSession` |
| `app/QA/Default.aspx.cs` | 100 | C# | `bnAdduser_Click` |
| `app/QA/Default.aspx.cs` | 117 | C# | `clearAddnew` |
| `app/QA/Default.aspx.cs` | 122 | C# | `addUser` |
| `app/QA/Default.aspx.cs` | 124 | C# | `CConnect` |
| `app/QA/Poll_detail.aspx.cs` | 26 | C# | `ShowMessage` |
| `app/QA/Poll_detail.aspx.cs` | 31 | C# | `Page_Load` |
| `app/QA/Poll_detail.aspx.cs` | 52 | C# | `delete_course` |
| `app/QA/Poll_detail.aspx.cs` | 71 | C# | `Button2_Click` |
| `app/QA/Poll_detail.aspx.cs` | 88 | C# | `SetSession` |
| `app/QA/Poll_detail.aspx.cs` | 93 | C# | `bnAdduser_Click` |
| `app/QA/Poll_detail.aspx.cs` | 110 | C# | `clearAddnew` |
| `app/QA/Poll_detail.aspx.cs` | 115 | C# | `addUser` |
| `app/QA/Project_edit.aspx.cs` | 14 | C# | `Page_Load` |
| `app/QA/Standard.aspx.cs` | 25 | C# | `ShowMessage` |
| `app/QA/Standard.aspx.cs` | 30 | C# | `Page_Load` |
| `app/QA/Standard.aspx.cs` | 51 | C# | `delete_course` |
| `app/QA/Standard.aspx.cs` | 70 | C# | `Button2_Click` |
| `app/QA/Standard.aspx.cs` | 87 | C# | `SetSession` |
| `app/QA/Standard.aspx.cs` | 92 | C# | `bnAdduser_Click` |
| `app/QA/Standard.aspx.cs` | 109 | C# | `clearAddnew` |
| `app/QA/Standard.aspx.cs` | 114 | C# | `addUser` |
| `app/QA/Standard_detail.aspx.cs` | 26 | C# | `ShowMessage` |

_Display capped at 300 of 6044 detected functions/methods._

### Automation pipeline files

- `.github/workflows/build-and-package.yml`
- `.github/workflows/ci.yml`
- `.github/workflows/deploy-staging.yml`
- `.github/workflows/security-scan.yml`
- `INSTALLER.md`
- `app/Properties/PublishProfiles/peb-lms-staging - Web Deploy.pubxml`
- `app/Properties/PublishProfiles/peb-lms-staging - Web Deploy.pubxml.user`
- `app/courseware/tests/Bootstrap.php`
- `app/courseware/user_guide/database/transactions.html`
- `app/examdb/user_guide/database/transactions.html`
- `app/knowledge/forum/install/mssql/install/azure/InstallCommon.sql`
- `app/knowledge/forum/install/mssql/install/azure/InstallMembership.sql`
- `app/knowledge/forum/install/mssql/install/azure/InstallProfile.SQL`
- `app/knowledge/forum/install/mssql/install/azure/InstallRoles.sql`
- `app/ui__/css/bootstrap-table.css`
- `app/ui__/css/bootstrap-theme.css`
- `app/ui__/css/bootstrap-theme.css.map`
- `app/ui__/css/bootstrap-theme.min.css`
- `app/ui__/css/bootstrap-theme.min.css.map`
- `app/ui__/css/bootstrap.css`
- `app/ui__/css/bootstrap.css.map`
- `app/ui__/css/bootstrap.min.css`
- `app/ui__/css/bootstrap.min.css.map`
- `app/ui__/js/bootstrap-datepicker.js`
- `app/ui__/js/bootstrap.js`
- `app/ui__/js/bootstrap.min.js`
- `app/web/css/bootstrap-table.css`
- `app/web/css/bootstrap-theme.css`
- `app/web/css/bootstrap-theme.css.map`
- `app/web/css/bootstrap-theme.min.css`
- `app/web/css/bootstrap-theme.min.css.map`
- `app/web/css/bootstrap.css`
- `app/web/css/bootstrap.css.map`
- `app/web/css/bootstrap.min.css`
- `app/web/css/bootstrap.min.css.map`
- `app/web/js/bootstrap-datepicker.js`
- `app/web/js/bootstrap.js`
- `app/web/js/bootstrap.min.js`
- `installer.sh`
- `scripts/generate_installer.sh`

