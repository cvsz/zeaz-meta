# Function, API, and Pipeline Inventory

Generated complete static index of detected functions, HTTP/API strings, and automation pipeline files.

## `zgitcp`

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

## `zwallet`

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
| `infra/installer/install.prod.sh` | 20 | Shell | `fail` |
| `infra/installer/install.prod.sh` | 22 | Shell | `require_root` |
| `infra/installer/install.prod.sh` | 26 | Shell | `install_base` |
| `infra/installer/install.prod.sh` | 33 | Shell | `setup_dirs` |
| `infra/installer/install.prod.sh` | 38 | Shell | `generate_secrets` |
| `infra/installer/install.prod.sh` | 50 | Shell | `generate_tls` |
| `infra/installer/install.prod.sh` | 58 | Shell | `run_postgres` |
| `infra/installer/install.prod.sh` | 66 | Shell | `run_redis` |
| `infra/installer/install.prod.sh` | 74 | Shell | `run_elasticsearch` |
| `infra/installer/install.prod.sh` | 84 | Shell | `health_check` |
| `infra/installer/install.prod.sh` | 88 | Shell | `main` |
| `infra/installer/install.sh` | 14 | Shell | `log` |
| `infra/installer/install.sh` | 15 | Shell | `err` |
| `infra/installer/install.sh` | 17 | Shell | `require_root` |
| `infra/installer/install.sh` | 21 | Shell | `install_base` |
| `infra/installer/install.sh` | 32 | Shell | `install_docker` |
| `infra/installer/install.sh` | 48 | Shell | `install_k3s` |
| `infra/installer/install.sh` | 53 | Shell | `setup_dirs` |
| `infra/installer/install.sh` | 58 | Shell | `generate_certs` |
| `infra/installer/install.sh` | 67 | Shell | `run_core_services` |
| `infra/installer/install.sh` | 90 | Shell | `run_edge_waf` |
| `infra/installer/install.sh` | 96 | Shell | `run_siem` |
| `infra/installer/install.sh` | 114 | Shell | `post_checks` |
| `infra/installer/install.sh` | 120 | Shell | `main` |
| `infra/observability/alerts/webhook_alert.py` | 9 | Python | `send_alert` |
| `infra/observability/structured_logger.py` | 9 | Python | `log_event` |
| `mobile/App.tsx` | 52 | JS/TS | `createIntent` |
| `packages/crypto-core/src/evmSigner.ts` | 20 | JS/TS | `sign` |
| `packages/crypto-core/src/hdWallet.ts` | 13 | JS/TS | `constructor` |
| `packages/crypto-core/src/hdWallet.ts` | 17 | JS/TS | `create` |
| `packages/crypto-core/src/hdWallet.ts` | 23 | JS/TS | `fromMnemonic` |
| `packages/crypto-core/src/hdWallet.ts` | 32 | JS/TS | `deriveEvm` |
| `packages/crypto-core/src/hdWallet.ts` | 51 | JS/TS | `wipe` |
| `packages/crypto-core/src/txPipeline.ts` | 11 | JS/TS | `simulate` |
| `packages/crypto-core/src/txPipeline.ts` | 12 | JS/TS | `estimateGas` |
| `packages/crypto-core/src/txPipeline.ts` | 13 | JS/TS | `sendRawTransaction` |
| `packages/crypto-core/src/txPipeline.ts` | 19 | JS/TS | `constructor` |
| `packages/crypto/src/adapters.ts` | 10 | JS/TS | `signByChain` |
| `packages/crypto/src/btc.ts` | 10 | JS/TS | `buildUnsignedBtcTx` |
| `packages/crypto/src/btc.ts` | 21 | JS/TS | `txidFromRaw` |
| `packages/crypto/src/btc.ts` | 27 | JS/TS | `u32` |
| `packages/crypto/src/btc.ts` | 28 | JS/TS | `u64` |
| `packages/crypto/src/btc.ts` | 29 | JS/TS | `varSlice` |
| `packages/crypto/src/encryption.ts` | 9 | JS/TS | `encryptAes256Gcm` |
| `packages/crypto/src/encryption.ts` | 20 | JS/TS | `decryptAes256Gcm` |
| `packages/crypto/src/evm.ts` | 11 | JS/TS | `signEvmMessageDeterministic` |
| `packages/crypto/src/evm.ts` | 18 | JS/TS | `signEip712TypedData` |
| `packages/crypto/src/hdWallet.ts` | 11 | JS/TS | `createMnemonic` |
| `packages/crypto/src/hdWallet.ts` | 15 | JS/TS | `deriveBip32Node` |
| `packages/crypto/src/hdWallet.ts` | 21 | JS/TS | `deriveSeed` |
| `packages/crypto/src/memory.ts` | 1 | JS/TS | `wipeBuffer` |
| `packages/crypto/src/mpcWallet.ts` | 15 | JS/TS | `constructor` |
| `packages/crypto/src/mpcWallet.ts` | 22 | JS/TS | `addShare` |
| `packages/crypto/src/mpcWallet.ts` | 27 | JS/TS | `combineShares` |
| `packages/crypto/src/mpcWallet.ts` | 46 | JS/TS | `sign` |
| `packages/crypto/src/mpcWallet.ts` | 59 | JS/TS | `encryptShare` |
| `packages/crypto/src/mpcWallet.ts` | 70 | JS/TS | `decryptShare` |
| `packages/crypto/src/solana.ts` | 9 | JS/TS | `signSolanaMessageDeterministic` |
| `packages/crypto/src/solana.ts` | 14 | JS/TS | `getSolanaPublicKey` |
| `packages/crypto/src/solana.ts` | 18 | JS/TS | `verifySolanaSignature` |
| `packages/events/src/eventBus.ts` | 7 | JS/TS | `connect` |
| `packages/events/src/eventBus.ts` | 15 | JS/TS | `connect` |
| `packages/events/src/eventBus.ts` | 34 | JS/TS | `constructor` |
| `packages/events/src/eventBus.ts` | 40 | JS/TS | `connect` |
| `packages/events/src/eventBus.ts` | 83 | JS/TS | `shouldProcess` |
| `packages/events/src/eventBus.ts` | 92 | JS/TS | `createEventBusFromEnv` |
| `packages/rpc/src/index.ts` | 72 | JS/TS | `getProviderHealth` |
| `packages/rpc/src/index.ts` | 83 | JS/TS | `pickCandidates` |
| `packages/rpc/src/index.ts` | 109 | JS/TS | `onSuccess` |
| `packages/rpc/src/index.ts` | 115 | JS/TS | `onFailure` |
| `packages/rpc/src/index.ts` | 125 | JS/TS | `createFetchRpcClient` |
| `packages/shared-types/src/wallet.ts` | 6 | JS/TS | `isSupportedChain` |
| `packages/shared-types/src/wallet.ts` | 10 | JS/TS | `assertAddress` |
| `packages/shared-types/src/wallet.ts` | 16 | JS/TS | `assertTransferRequest` |
| `scripts/migrate.sh` | 6 | Shell | `apply` |
| `scripts/rollback.sh` | 6 | Shell | `rollback_ledger` |
| `scripts/rollback.sh` | 11 | Shell | `rollback_idempotency` |
| `security/controller/main.go` | 30 | Go | `initK8s` |
| `security/controller/main.go` | 42 | Go | `quarantinePod` |
| `security/controller/main.go` | 60 | Go | `sanitizeLogField` |
| `security/controller/main.go` | 66 | Go | `handler` |
| `security/controller/main.go` | 89 | Go | `main` |
| `services/card/src/index.ts` | 35 | JS/TS | `createTokenizedCard` |
| `services/card/src/index.ts` | 36 | JS/TS | `freezeCard` |
| `services/card/src/index.ts` | 37 | JS/TS | `unfreezeCard` |
| `services/card/src/index.ts` | 40 | JS/TS | `assertNonEmpty` |
| `services/card/src/index.ts` | 47 | JS/TS | `constructor` |
| `services/card/src/index.ts` | 49 | JS/TS | `issueCard` |
| `services/card/src/index.ts` | 80 | JS/TS | `freeze` |
| `services/card/src/index.ts` | 85 | JS/TS | `unfreeze` |
| `services/compliance/src/index.ts` | 25 | JS/TS | `verifyIdentity` |
| `services/compliance/src/index.ts` | 42 | JS/TS | `evaluate` |
| `services/compliance/src/index.ts` | 64 | JS/TS | `constructor` |
| `services/compliance/src/index.ts` | 66 | JS/TS | `runKycAndRisk` |
| `services/event-workers/src/index.ts` | 19 | JS/TS | `fakeRoutes` |
| `services/event-workers/src/index.ts` | 26 | JS/TS | `handleTxRequested` |
| `services/event-workers/src/index.ts` | 39 | JS/TS | `handleSwapRequested` |
| `services/indexer/src/index.ts` | 69 | JS/TS | `process` |
| `services/indexer/src/index.ts` | 83 | JS/TS | `hasSeen` |
| `services/indexer/src/index.ts` | 87 | JS/TS | `buildCursor` |
| `services/indexer/src/index.ts` | 102 | JS/TS | `buildDedupeKey` |
| `services/liquidity/src/index.ts` | 28 | JS/TS | `getQuote` |
| `services/liquidity/src/index.ts` | 29 | JS/TS | `execute` |
| `services/liquidity/src/index.ts` | 42 | JS/TS | `parsePositiveAtomic` |
| `services/liquidity/src/index.ts` | 54 | JS/TS | `constructor` |
| `services/liquidity/src/index.ts` | 56 | JS/TS | `quote` |
| `services/liquidity/src/index.ts` | 65 | JS/TS | `execute` |
| `services/router/src/intentRouter.ts` | 32 | JS/TS | `constructor` |
| `services/router/src/intentRouter.ts` | 36 | JS/TS | `fetchQuotes` |
| `services/router/src/intentRouter.ts` | 60 | JS/TS | `pickBest` |
| `services/swap-engine/src/index.ts` | 41 | JS/TS | `assertNonEmpty` |
| `services/swap-engine/src/index.ts` | 47 | JS/TS | `parseAtomic` |
| `services/swap-engine/src/index.ts` | 58 | JS/TS | `parseGweiToWei` |
| `services/swap-engine/src/index.ts` | 77 | JS/TS | `validateQuoteRequest` |
| `services/swap-engine/src/index.ts` | 90 | JS/TS | `scoreRoute` |
| `services/swap-engine/src/index.ts` | 117 | JS/TS | `selectBestRoute` |
| `services/wallet-engine/src/walletEngine.ts` | 6 | JS/TS | `buildTransferDigest` |
| `terraform/aws/lambda/stop_instances.py` | 7 | Python | `lambda_handler` |
| `tests/ledger/stress.ts` | 3 | JS/TS | `run` |

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

## `ABTPi18n`

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
| `apps/backend/src/worker/tasks.py` | 92 | Python | `check_secret_rotation` |
| `apps/backend/src/worker/tasks.py` | 97 | Python | `check_secret_rotation_async` |
| `apps/backend/src/worker/tasks.py` | 138 | Python | `cleanup_audit_logs` |
| `apps/backend/src/worker/tasks.py` | 143 | Python | `cleanup_audit_logs_async` |
| `apps/frontend/src/app/[lng]/dashboard/page.tsx` | 49 | JS/TS | `Strategies` |
| `apps/frontend/src/app/[lng]/dashboard/page.tsx` | 66 | JS/TS | `BotControl` |
| `apps/frontend/src/app/[lng]/dashboard/page.tsx` | 70 | JS/TS | `startBot` |
| `apps/frontend/src/app/[lng]/dashboard/page.tsx` | 78 | JS/TS | `stopBot` |
| `apps/frontend/src/app/[lng]/login/page.tsx` | 19 | JS/TS | `handleSuccess` |
| `apps/frontend/src/app/[lng]/login/page.tsx` | 25 | JS/TS | `handleError` |
| `apps/frontend/src/app/[lng]/settings/page.tsx` | 26 | JS/TS | `save` |
| `apps/frontend/src/app/i18n/client.ts` | 18 | JS/TS | `initI18n` |
| `apps/frontend/src/components/LanguageSelector.tsx` | 19 | JS/TS | `LanguageSelector` |
| `apps/frontend/src/components/LanguageSelector.tsx` | 26 | JS/TS | `handleLanguageChange` |
| `apps/frontend/src/components/Navigation.tsx` | 24 | JS/TS | `Navigation` |
| `apps/frontend/src/components/Navigation.tsx` | 33 | JS/TS | `isActive` |
| `apps/frontend/src/components/auth/GoogleSignIn.tsx` | 16 | JS/TS | `GoogleSignIn` |
| `apps/frontend/src/components/auth/GoogleSignIn.tsx` | 20 | JS/TS | `handleGoogleSignIn` |
| `apps/frontend/src/components/settings/NotificationPreferences.tsx` | 22 | JS/TS | `NotificationPreferences` |
| `apps/frontend/src/components/settings/NotificationPreferences.tsx` | 39 | JS/TS | `loadPreferences` |
| `apps/frontend/src/components/settings/NotificationPreferences.tsx` | 49 | JS/TS | `handleSave` |
| `apps/frontend/src/components/settings/NotificationPreferences.tsx` | 84 | JS/TS | `togglePreference` |
| `apps/frontend/src/components/settings/TelegramLink.tsx` | 22 | JS/TS | `TelegramLink` |
| `apps/frontend/src/components/settings/TelegramLink.tsx` | 37 | JS/TS | `loadStatus` |
| `apps/frontend/src/components/settings/TelegramLink.tsx` | 47 | JS/TS | `handleLink` |
| `apps/frontend/src/components/settings/TelegramLink.tsx` | 87 | JS/TS | `handleUnlink` |
| `apps/frontend/src/components/settings/TelegramLink.tsx` | 122 | JS/TS | `handleTestNotification` |
| `apps/frontend/src/components/settings/ThemeCustomizer.tsx` | 12 | JS/TS | `ThemeCustomizer` |
| `apps/frontend/src/components/settings/ThemeCustomizer.tsx` | 29 | JS/TS | `handleApplyColors` |
| `apps/frontend/src/components/settings/ThemeCustomizer.tsx` | 35 | JS/TS | `handleResetColors` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 26 | JS/TS | `ThemeProvider` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 58 | JS/TS | `updateActualTheme` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 79 | JS/TS | `handleChange` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 85 | JS/TS | `setTheme` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 92 | JS/TS | `handleSetPrimaryColor` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 99 | JS/TS | `handleSetSecondaryColor` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 106 | JS/TS | `handleSetAccentColor` |
| `apps/frontend/src/contexts/ThemeContext.tsx` | 132 | JS/TS | `useTheme` |
| `core/strategy_autoload.py` | 20 | Python | `load_external_strategies` |
| `core/strategy_base.py` | 25 | Python | `execute` |
| `core/strategy_registry.py` | 33 | Python | `register` |
| `core/strategy_registry.py` | 51 | Python | `create` |
| `core/strategy_registry.py` | 72 | Python | `list_names` |
| `core/strategy_registry.py` | 82 | Python | `get_class` |
| `fix.sh` | 40 | Shell | `log_info` |
| `fix.sh` | 44 | Shell | `log_success` |
| `fix.sh` | 48 | Shell | `log_warn` |
| `fix.sh` | 52 | Shell | `log_error` |
| `fix.sh` | 56 | Shell | `log_debug` |
| `fix.sh` | 63 | Shell | `command_exists` |
| `fix.sh` | 68 | Shell | `check_disk_space` |
| `fix.sh` | 85 | Shell | `prune_docker` |
| `fix.sh` | 109 | Shell | `install_nodejs_in_container` |
| `fix.sh` | 148 | Shell | `generate_prisma_client` |
| `fix.sh` | 174 | Shell | `run_migrations` |
| `fix.sh` | 193 | Shell | `main` |
| `install.sh` | 12 | Shell | `need_cmd` |
| `scripts/export-docs-to-wiki.sh` | 40 | Shell | `inplace_sed` |
| `scripts/export-docs-to-wiki.sh` | 58 | Shell | `convert_to_wiki_name` |
| `scripts/export-docs-to-wiki.sh` | 75 | Shell | `detect_collisions` |
| `scripts/export-docs-to-wiki.sh` | 109 | Shell | `rewrite_markdown_links` |
| `scripts/export-docs-to-wiki.sh` | 137 | Shell | `collect_assets` |
| `scripts/export-docs-to-wiki.sh` | 148 | Shell | `copy_assets` |
| `scripts/export-docs-to-wiki.sh` | 186 | Shell | `rewrite_asset_paths` |
| `scripts/export-docs-to-wiki.sh` | 278 | Shell | `copy_doc_file` |
| `scripts/zeaz_meta_installer.sh` | 12 | Shell | `log` |
| `scripts/zeaz_meta_installer.sh` | 18 | Shell | `usage` |
| `scripts/zeaz_meta_installer.sh` | 63 | Shell | `plan` |
| `scripts/zeaz_meta_installer.sh` | 68 | Shell | `generate` |
| `scripts/zeaz_meta_installer.sh` | 78 | Shell | `install_files` |
| `scripts/zeaz_meta_installer.sh` | 97 | Shell | `release` |
| `tests/test_metaultra_example.py` | 12 | Python | `test_example_strategy_signals` |
| `tests/test_metaultra_example.py` | 17 | Python | `test_on_tick_prints` |
| `tools/drive_sync.py` | 27 | Python | `extract_folder_id` |
| `tools/drive_sync.py` | 63 | Python | `download_drive_folder` |
| `tools/drive_sync.py` | 113 | Python | `main` |
| `tools/generate_demo.js` | 24 | JS/TS | `createPlaceholderSVG` |
| `tools/generate_demo.js` | 308 | JS/TS | `openModal` |
| `tools/generate_demo.js` | 315 | JS/TS | `closeModal` |
| `tools/integrate_drive_assets.py` | 29 | Python | `load_mapping_config` |
| `tools/integrate_drive_assets.py` | 57 | Python | `match_files` |
| `tools/integrate_drive_assets.py` | 90 | Python | `compute_destination` |
| `tools/integrate_drive_assets.py` | 133 | Python | `integrate_assets` |
| `tools/integrate_drive_assets.py` | 219 | Python | `main` |
| `tools/load_gdrive_strategies.py` | 20 | Python | `main` |
| `tools/metaultra/example_module.py` | 12 | Python | `__init__` |
| `tools/metaultra/example_module.py` | 16 | Python | `on_tick` |
| `tools/metaultra/example_module.py` | 20 | Python | `generate_signals` |
| `tools/metaultra/example_module.ts` | 7 | JS/TS | `constructor` |
| `tools/metaultra/example_module.ts` | 12 | JS/TS | `onTick` |
| `tools/metaultra/example_module.ts` | 16 | JS/TS | `generateSignals` |
| `tools/screenshot_pages.js` | 35 | JS/TS | `ensureDir` |
| `tools/screenshot_pages.js` | 44 | JS/TS | `captureScreenshot` |
| `tools/screenshot_pages.js` | 68 | JS/TS | `captureAllPages` |
| `tools/screenshot_pages.js` | 126 | JS/TS | `generateIndexHtml` |
| `tools/screenshot_pages.js` | 309 | JS/TS | `openModal` |
| `tools/screenshot_pages.js` | 316 | JS/TS | `closeModal` |
| `tools/screenshot_pages.js` | 344 | JS/TS | `generateLanguageSections` |
| `tools/test_tradingview_integration.py` | 17 | Python | `check_pydantic_models` |
| `tools/test_tradingview_integration.py` | 45 | Python | `check_strategy_logic` |
| `tools/test_tradingview_integration.py` | 53 | Python | `__init__` |
| `tools/test_tradingview_integration.py` | 57 | Python | `execute` |
| `tools/test_tradingview_integration.py` | 125 | Python | `check_yaml_config` |
| `tools/test_tradingview_integration.py` | 155 | Python | `check_gdrive_loader_structure` |
| `tools/test_tradingview_integration.py` | 181 | Python | `check_endpoint_structure` |
| `tools/test_tradingview_integration.py` | 218 | Python | `test_pydantic_models` |
| `tools/test_tradingview_integration.py` | 222 | Python | `test_strategy_logic` |
| `tools/test_tradingview_integration.py` | 226 | Python | `test_yaml_config` |
| `tools/test_tradingview_integration.py` | 230 | Python | `test_gdrive_loader_structure` |
| `tools/test_tradingview_integration.py` | 234 | Python | `test_endpoint_structure` |
| `tools/test_tradingview_integration.py` | 238 | Python | `main` |

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

## `zypto`

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
| `monopoly/mna.py` | 14 | Python | `score` |
| `monopoly/mna.py` | 19 | Python | `shortlist` |
| `monopoly/pricing.py` | 1 | Python | `segment_price` |
| `ops/control_plane/edge_cache.py` | 6 | Python | `set_kv` |
| `ops/control_plane/edge_cache.py` | 10 | Python | `get_kv` |
| `ops/control_plane/jwks_service.py` | 6 | Python | `publish` |
| `ops/control_plane/jwks_service.py` | 10 | Python | `get` |
| `ops/control_plane/policy_service.py` | 1 | Python | `allow` |
| `ops/control_plane/revocation_service.py` | 6 | Python | `revoke` |
| `ops/control_plane/revocation_service.py` | 10 | Python | `is_revoked` |
| `ops/recovery/auto.py` | 1 | Python | `recover` |
| `p2p/go-node/main.go` | 21 | Go | `HandlePeerFound` |
| `p2p/go-node/main.go` | 27 | Go | `main` |
| `p2p/gossipsub/main.go` | 18 | Go | `main` |
| `partners/incentives.py` | 4 | Python | `partner_commission` |
| `partners/marketplace.py` | 4 | Python | `list_marketplace` |
| `partners/onboarding.py` | 4 | Python | `onboard_partner` |
| `payments/stripe/billing.py` | 4 | Python | `compute_bill` |
| `payments/stripe/checkout.py` | 10 | Python | `create_subscription` |
| `payments/stripe/metered.py` | 8 | Python | `report_usage` |
| `payments/stripe/tax.py` | 4 | Python | `apply_tax` |
| `payments/stripe/webhook.py` | 9 | Python | `handle` |
| `pilot/scripts/smoke_test.py` | 6 | Python | `test` |
| `pitchdeck/demo/app.py` | 24 | Python | `generate_signal` |
| `pitchdeck/demo/app.py` | 69 | Python | `tenants` |
| `pitchdeck/demo/app.py` | 74 | Python | `tenant_snapshot` |
| `pitchdeck/demo/app.py` | 81 | Python | `ws_tenant_stream` |
| `pitchdeck/demo/app.py` | 101 | Python | `home` |
| `policy_sim/scenarios.py` | 1 | Python | `simulate` |
| `post_ipo/capital_allocation.py` | 1 | Python | `allocate` |
| `post_ipo/metrics.py` | 1 | Python | `metrics` |
| `public_sector_eu/controls/logging.py` | 3 | Python | `log` |
| `ratification/telemetry/collector.py` | 4 | Python | `record` |
| `ratification/telemetry/collector.py` | 8 | Python | `summary` |
| `regulatory_moat/engine/attestation.py` | 1 | Python | `generate` |
| `regulatory_moat/engine/control_engine.py` | 5 | Python | `__init__` |
| `regulatory_moat/engine/control_engine.py` | 8 | Python | `run` |
| `regulatory_moat/engine/evidence_store.py` | 8 | Python | `put` |
| `regulatory_moat/engine/evidence_store.py` | 16 | Python | `query` |
| `regulatory_moat/engine/policy_engine.py` | 9 | Python | `evaluate` |
| `regulatory_moat/engine/reporting.py` | 1 | Python | `report` |
| `regulatory_moat/integrations/aws.py` | 1 | Python | `validate_bucket` |
| `regulatory_moat/integrations/idp.py` | 1 | Python | `validate_identity_policy` |
| `regulatory_moat/integrations/k8s.py` | 1 | Python | `validate_pod` |
| `revenue/billing/aggregator.py` | 1 | Python | `total_multi_cloud` |
| `revenue/billing/aws_cur_ingest.py` | 4 | Python | `load_cur` |
| `revenue/billing/aws_cur_ingest.py` | 9 | Python | `total_cost` |
| `revenue/billing/azure_cost_ingest.py` | 1 | Python | `parse_azure` |
| `revenue/billing/gcp_billing_ingest.py` | 1 | Python | `parse_gcp` |
| `revenue/billing/invoice.py` | 4 | Python | `generate_invoice` |
| `revenue/billing/pricing.py` | 17 | Python | `price` |
| `revenue/marketplace/demand.py` | 1 | Python | `add_demand` |
| `revenue/marketplace/matching.py` | 1 | Python | `match` |
| `revenue/marketplace/supply.py` | 1 | Python | `add_supply` |
| `revenue/metering/collector.py` | 4 | Python | `record` |
| `revenue/metering/usage_store.py` | 4 | Python | `get_usage` |
| `revenue/sales/contracts.py` | 1 | Python | `contract` |
| `revenue/sales/pipeline.py` | 4 | Python | `advance` |
| `revenue/sla/enforcement.py` | 7 | Python | `check` |
| `revenue/sla/penalties.py` | 1 | Python | `penalty` |
| `revenue/workloads/batch.py` | 1 | Python | `run_batch` |
| `revenue/workloads/inference.py` | 1 | Python | `run_inference` |
| `revenue/workloads/streaming.py` | 1 | Python | `stream` |
| `revops/board_metrics.py` | 3 | Python | `metrics` |
| `revops/compensation.py` | 3 | Python | `commission` |
| `revops/forecasting_ai.py` | 3 | Python | `weighted_pipeline` |
| `revops/forecasting_ai.py` | 6 | Python | `forecast_next_month` |
| `revops/quota.py` | 3 | Python | `assign_quota` |
| `saas/backend/auth.py` | 10 | Python | `create_token` |
| `saas/backend/auth.py` | 20 | Python | `verify_token` |
| `saas/backend/main.py` | 25 | Python | `signup` |
| `saas/backend/main.py` | 31 | Python | `me` |
| `saas/backend/main.py` | 39 | Python | `ws_metrics` |
| `saas/backend/main.py` | 58 | Python | `healthz` |
| `saas/backend/stripe_api.py` | 16 | Python | `create_checkout` |
| `saas/frontend/app/billing/page.tsx` | 4 | JS/TS | `subscribe` |
| `saas/frontend/app/page.tsx` | 12 | JS/TS | `Sparkline` |
| `sales/pipeline.py` | 4 | Python | `next_stage` |
| `sales_crm/api.py` | 10 | Python | `pipeline` |
| `sales_crm/api.py` | 15 | Python | `forecast` |
| `sales_crm/automation.py` | 1 | Python | `next_action` |
| `sales_crm/forecasting.py` | 4 | Python | `forecast_revenue` |
| `sales_crm/forecasting.py` | 9 | Python | `forecast_growth` |
| `sales_crm/pipeline.py` | 6 | Python | `advance` |
| `sales_crm/pipeline.py` | 15 | Python | `pipeline_value` |
| `sales_org/comp_plan.py` | 4 | Python | `commission` |
| `sales_org/quotas.py` | 4 | Python | `quota` |
| `scale/traffic/replay.py` | 1 | Python | `replay` |
| `scaling/kpis.py` | 4 | Python | `metrics` |
| `scaling_50m/metrics.py` | 4 | Python | `metrics` |
| `scripts/zeaz-enterprise-install.sh` | 16 | Shell | `usage` |
| `scripts/zeaz-enterprise-install.sh` | 42 | Shell | `run` |
| `scripts/zeaz-enterprise-install.sh` | 50 | Shell | `run_bash` |
| `scripts/zeaz-enterprise-install.sh` | 58 | Shell | `need` |
| `scripts/zeaz-enterprise-install.sh` | 62 | Shell | `preflight` |
| `scripts/zeaz-enterprise-install.sh` | 72 | Shell | `generate_values` |
| `scripts/zeaz-enterprise-install.sh` | 102 | Shell | `install_argocd` |
| `scripts/zeaz-enterprise-install.sh` | 114 | Shell | `apply_platform_manifests` |
| `scripts/zeaz-enterprise-install.sh` | 121 | Shell | `post_install_report` |
| `scripts/zeaz-enterprise-install.sh` | 134 | Shell | `main` |
| `scripts/zeaz-os-bootstrap.sh` | 14 | Shell | `usage` |
| `scripts/zeaz-os-bootstrap.sh` | 46 | Shell | `run` |
| `scripts/zeaz-os-bootstrap.sh` | 54 | Shell | `run_bash` |
| `scripts/zeaz-os-bootstrap.sh` | 62 | Shell | `require_ubuntu_2404` |
| `scripts/zeaz-os-bootstrap.sh` | 72 | Shell | `clean_environment` |
| `scripts/zeaz-os-bootstrap.sh` | 94 | Shell | `install_base_packages` |
| `scripts/zeaz-os-bootstrap.sh` | 105 | Shell | `install_container_runtime` |
| `scripts/zeaz-os-bootstrap.sh` | 118 | Shell | `install_k8s_tools` |
| `scripts/zeaz-os-bootstrap.sh` | 129 | Shell | `harden_defaults` |
| `scripts/zeaz-os-bootstrap.sh` | 140 | Shell | `write_state` |
| `scripts/zeaz-os-bootstrap.sh` | 144 | Shell | `main` |
| `security/risk/scoring.py` | 1 | Python | `risk_score` |
| `security/risk/scoring.py` | 19 | Python | `allow` |
| `security/soc2_check.py` | 33 | Python | `run_checks` |
| `security/soc2_check.py` | 44 | Python | `compliance_score` |
| `services/bridge/main.go` | 14 | Go | `publishHandler` |
| `services/bridge/main.go` | 33 | Go | `main` |
| `services/tss-signing/cmd/server/main.go` | 13 | Go | `main` |
| `services/tss-signing/internal/audit/audit.go` | 27 | Go | `Emit` |
| `services/tss-signing/internal/config/config.go` | 19 | Go | `Load` |
| `services/tss-signing/internal/config/config.go` | 35 | Go | `getenv` |
| `services/tss-signing/internal/config/config.go` | 42 | Go | `getenvInt` |
| `services/tss-signing/internal/config/config.go` | 54 | Go | `getenvBool` |
| `services/tss-signing/internal/signer/signer.go` | 33 | Go | `Sign` |
| `services/tss-signing/internal/signer/signer.go` | 44 | Go | `Sign` |
| `services/tss-signing/internal/signer/signer.go` | 61 | Go | `stableDigest` |
| `sovereign/compliance/audit.py` | 6 | Python | `log` |
| `sovereign/compliance/audit.py` | 10 | Python | `query` |
| `sovereign/compliance/dsr.py` | 1 | Python | `export_user` |
| `sovereign/compliance/dsr.py` | 5 | Python | `delete_user` |
| `sovereign/compliance/reports.py` | 1 | Python | `compliance_summary` |
| `sovereign/control_plane/geo_router.py` | 6 | Python | `route` |
| `sovereign/control_plane/identity.py` | 11 | Python | `assert_eu_scope` |
| `sovereign/control_plane/policy_engine.py` | 6 | Python | `allow` |
| `sovereign/data_plane/encryption.py` | 9 | Python | `__init__` |
| `sovereign/data_plane/encryption.py` | 12 | Python | `encrypt` |
| `sovereign/data_plane/encryption.py` | 15 | Python | `decrypt` |
| `sovereign/data_plane/partitioning.py` | 1 | Python | `partition_key` |
| `sovereign/data_plane/retention.py` | 4 | Python | `ttl_expire` |
| `sovereign_ai/governance/audit.py` | 6 | Python | `log` |
| `sovereign_ai/governance/model_registry.py` | 4 | Python | `register` |
| `sovereign_ai/governance/risk_classification.py` | 1 | Python | `classify` |
| `sovereign_ai/serving/guardrails.py` | 1 | Python | `enforce` |
| `sovereign_ai/serving/inference.py` | 6 | Python | `run` |
| `sovereign_ai/serving/router.py` | 1 | Python | `route` |
| `tender_ai/generator/proposal_builder.py` | 1 | Python | `build` |
| `tender_ai/generator/rfp_parser.py` | 1 | Python | `parse_rfp` |
| `tender_ai/optimization/bid_optimizer.py` | 1 | Python | `optimize` |
| `tender_ai/optimization/learning.py` | 4 | Python | `record` |
| `tender_ai/optimization/learning.py` | 8 | Python | `adjust` |
| `tender_ai/optimization/scoring_model.py` | 1 | Python | `score` |
| `tender_ai/pipeline/evaluator.py` | 4 | Python | `evaluate` |
| `tender_ai/pipeline/workflow.py` | 5 | Python | `run` |
| `testing/locustfile.py` | 8 | Python | `hit_api` |
| `tests/contract/test_openapi.py` | 11 | Python | `test_openapi_contract` |
| `tests/contract/test_orders_contract.py` | 8 | Python | `test_metrics_contract_shape` |
| `tests/revenue/billing/test_pricing.py` | 6 | Python | `test_price_sums_supported_metrics` |
| `tests/revenue/billing/test_pricing.py` | 15 | Python | `test_price_rejects_unknown_metric` |
| `tests/security/test_tenant_isolation.py` | 8 | Python | `test_requires_auth_header` |
| `tests/security/test_tenant_isolation.py` | 13 | Python | `test_tenant_isolation_violation` |
| `tools/arch_guard.py` | 12 | Python | `load_rules` |
| `tools/arch_guard.py` | 17 | Python | `top_package` |
| `tools/arch_guard.py` | 23 | Python | `collect_imports` |
| `tools/arch_guard.py` | 39 | Python | `resolve_target_pkg` |
| `tools/arch_guard.py` | 43 | Python | `scan` |
| `tools/deep_scan.py` | 72 | Python | `run` |
| `tools/deep_scan.py` | 77 | Python | `safe_clone_or_pull` |
| `tools/deep_scan.py` | 90 | Python | `git_configure` |
| `tools/deep_scan.py` | 97 | Python | `create_autofix_branch` |
| `tools/deep_scan.py` | 104 | Python | `commit_changes` |
| `tools/deep_scan.py` | 115 | Python | `push_branch` |
| `tools/deep_scan.py` | 129 | Python | `create_gitea_pr` |
| `tools/deep_scan.py` | 157 | Python | `collect_python_entities` |
| `tools/deep_scan.py` | 178 | Python | `safe_backup_file` |
| `tools/deep_scan.py` | 187 | Python | `ast_autofix_file` |
| `tools/deep_scan.py` | 195 | Python | `visit_ImportFrom` |
| `tools/deep_scan.py` | 203 | Python | `visit_Call` |
| `tools/deep_scan.py` | 212 | Python | `visit_FunctionDef` |
| `tools/deep_scan.py` | 249 | Python | `build_embeddings` |
| `tools/deep_scan.py` | 263 | Python | `detect_anomalies` |
| `tools/deep_scan.py` | 275 | Python | `write_bug_report` |
| `tools/deep_scan.py` | 286 | Python | `write_graph` |
| `tools/deep_scan.py` | 292 | Python | `append_changelog` |
| `tools/deep_scan.py` | 302 | Python | `main` |
| `tools/dependency_graph.py` | 10 | Python | `top` |
| `tools/tests/test_autofix.py` | 8 | Python | `setup_module` |
| `tools/tests/test_autofix.py` | 14 | Python | `teardown_module` |
| `tools/tests/test_autofix.py` | 19 | Python | `test_ast_autofix_replaces_random_and_backups` |
| `trillion/ai_plane/inference.py` | 2 | Python | `run` |
| `trillion/ai_plane/model_router.py` | 9 | Python | `route` |
| `trillion/ai_plane/model_router.py` | 15 | Python | `score` |
| `trillion/ai_plane/training.py` | 2 | Python | `fit` |
| `trillion/control_plane/identity.py` | 4 | Python | `__init__` |
| `trillion/control_plane/identity.py` | 7 | Python | `issue` |
| `trillion/control_plane/identity.py` | 10 | Python | `verify` |
| `trillion/control_plane/policy_engine.py` | 6 | Python | `__init__` |
| `trillion/control_plane/policy_engine.py` | 9 | Python | `allow` |
| `trillion/control_plane/scheduler.py` | 9 | Python | `choose` |
| `trillion/control_plane/scheduler.py` | 16 | Python | `score` |
| `trillion/data_plane/feature_store.py` | 6 | Python | `__init__` |
| `trillion/data_plane/feature_store.py` | 10 | Python | `upsert` |
| `trillion/data_plane/feature_store.py` | 14 | Python | `get` |
| `trillion/data_plane/lakehouse.py` | 6 | Python | `__init__` |
| `trillion/data_plane/lakehouse.py` | 9 | Python | `append` |
| `trillion/data_plane/lakehouse.py` | 12 | Python | `query` |
| `trillion/data_plane/stream.py` | 9 | Python | `__init__` |
| `trillion/data_plane/stream.py` | 12 | Python | `publish` |
| `trillion/data_plane/stream.py` | 15 | Python | `consume` |
| `trillion/economics/marketplace.py` | 1 | Python | `match_supply` |
| `trillion/economics/pricing.py` | 1 | Python | `value_based_price` |
| `trillion/governance/audit.py` | 6 | Python | `log` |
| `trillion/governance/audit.py` | 10 | Python | `query` |
| `trust_fabric/attestation/attestor.py` | 1 | Python | `attest` |
| `trust_fabric/attestation/registry.py` | 4 | Python | `register` |
| `trust_fabric/attestation/registry.py` | 8 | Python | `get` |
| `trust_fabric/attestation/verifier.py` | 1 | Python | `verify` |
| `trust_fabric/identity/federation.py` | 4 | Python | `add_issuer` |
| `trust_fabric/identity/federation.py` | 8 | Python | `is_trusted` |
| `trust_fabric/identity/issuer.py` | 6 | Python | `issue` |
| `trust_fabric/identity/verifier.py` | 5 | Python | `verify` |
| `trust_fabric/policy/risk.py` | 1 | Python | `score` |
| `trust_fabric/policy/trust_policy.py` | 1 | Python | `allow` |
| `trust_fabric/transport/mtls.py` | 1 | Python | `enforce` |
| `v1/app/ai/agent.py` | 9 | Python | `__init__` |
| `v1/app/ai/agent.py` | 19 | Python | `forward` |
| `v1/app/ai/agent.py` | 24 | Python | `__init__` |
| `v1/app/ai/agent.py` | 31 | Python | `act` |
| `v1/app/ai/agent.py` | 36 | Python | `store` |
| `v1/app/ai/agent.py` | 41 | Python | `train_step` |
| `v1/app/ai/env.py` | 5 | Python | `__init__` |
| `v1/app/ai/env.py` | 11 | Python | `reset` |
| `v1/app/ai/env.py` | 17 | Python | `_get_obs` |
| `v1/app/ai/env.py` | 20 | Python | `step` |
| `v1/app/ai/inference.py` | 11 | Python | `__init__` |
| `v1/app/ai/inference.py` | 27 | Python | `predict` |
| `v1/app/ai/trainer.py` | 8 | Python | `train` |
| `v1/app/aiops/ab.py` | 5 | Python | `__init__` |
| `v1/app/aiops/ab.py` | 8 | Python | `set` |
| `v1/app/aiops/ab.py` | 11 | Python | `route` |
| `v1/app/aiops/adwin.py` | 8 | Python | `__init__` |
| `v1/app/aiops/adwin.py` | 13 | Python | `update` |
| `v1/app/aiops/agent_orchestrator.py` | 5 | Python | `__init__` |
| `v1/app/aiops/agent_orchestrator.py` | 11 | Python | `step` |
| `v1/app/aiops/agents/availability.py` | 5 | Python | `propose` |
| `v1/app/aiops/agents/bidder.py` | 5 | Python | `__init__` |
| `v1/app/aiops/agents/bidder.py` | 8 | Python | `bid` |
| `v1/app/aiops/agents/cost.py` | 5 | Python | `__init__` |
| `v1/app/aiops/agents/cost.py` | 9 | Python | `allow` |
| `v1/app/aiops/agents/governor.py` | 7 | Python | `__init__` |
| `v1/app/aiops/agents/governor.py` | 10 | Python | `cast` |
| `v1/app/aiops/agents/governor.py` | 13 | Python | `approved` |
| `v1/app/aiops/agents/latency.py` | 5 | Python | `propose` |
| `v1/app/aiops/agents/planner.py` | 7 | Python | `__init__` |
| `v1/app/aiops/agents/planner.py` | 11 | Python | `act` |
| `v1/app/aiops/agents/region.py` | 5 | Python | `propose` |
| `v1/app/aiops/agents/router.py` | 5 | Python | `decide` |
| `v1/app/aiops/agents/safety.py` | 5 | Python | `__init__` |
| `v1/app/aiops/agents/safety.py` | 8 | Python | `allow` |
| `v1/app/aiops/agents/seller.py` | 7 | Python | `__init__` |
| `v1/app/aiops/agents/seller.py` | 10 | Python | `ask` |
| `v1/app/aiops/audit/logger.py` | 10 | Python | `__init__` |
| `v1/app/aiops/audit/logger.py` | 13 | Python | `log` |
| `v1/app/aiops/audit/replay.py` | 7 | Python | `run` |
| `v1/app/aiops/audit/trail.py` | 7 | Python | `verify_chain` |
| `v1/app/aiops/autoscale.py` | 5 | Python | `set_replicas` |
| `v1/app/aiops/autoscale.py` | 8 | Python | `decide` |
| `v1/app/aiops/bandit.py` | 7 | Python | `__init__` |
| `v1/app/aiops/bandit.py` | 12 | Python | `select` |
| `v1/app/aiops/bandit.py` | 21 | Python | `update` |
| `v1/app/aiops/bandit_hb.py` | 9 | Python | `__init__` |
| `v1/app/aiops/bandit_hb.py` | 18 | Python | `_post` |
| `v1/app/aiops/bandit_hb.py` | 24 | Python | `_fuse` |
| `v1/app/aiops/bandit_hb.py` | 31 | Python | `select` |
| `v1/app/aiops/bandit_hb.py` | 45 | Python | `update` |
| `v1/app/aiops/bandit_hierarchical.py` | 7 | Python | `__init__` |
| `v1/app/aiops/bandit_hierarchical.py` | 11 | Python | `select` |
| `v1/app/aiops/bandit_hierarchical.py` | 16 | Python | `update` |
| `v1/app/aiops/bandit_nl.py` | 9 | Python | `__init__` |
| `v1/app/aiops/bandit_nl.py` | 15 | Python | `transform` |
| `v1/app/aiops/bandit_nl.py` | 24 | Python | `__init__` |
| `v1/app/aiops/bandit_nl.py` | 30 | Python | `embed` |
| `v1/app/aiops/bandit_nl.py` | 33 | Python | `select` |
| `v1/app/aiops/bandit_nl.py` | 50 | Python | `update` |
| `v1/app/aiops/bandit_nl.py` | 55 | Python | `increase_exploration` |
| `v1/app/aiops/bandit_nl.py` | 58 | Python | `reset_partial` |
| `v1/app/aiops/bandit_ts.py` | 7 | Python | `__init__` |
| `v1/app/aiops/bandit_ts.py` | 14 | Python | `_posterior` |
| `v1/app/aiops/bandit_ts.py` | 20 | Python | `select` |
| `v1/app/aiops/bandit_ts.py` | 32 | Python | `update` |
| `v1/app/aiops/bandit_ts_ns.py` | 7 | Python | `__init__` |
| `v1/app/aiops/bandit_ts_ns.py` | 15 | Python | `_posterior` |
| `v1/app/aiops/bandit_ts_ns.py` | 21 | Python | `select` |
| `v1/app/aiops/bandit_ts_ns.py` | 33 | Python | `update` |
| `v1/app/aiops/bayesian_rca.py` | 37 | Python | `infer` |
| `v1/app/aiops/canary.py` | 6 | Python | `__init__` |
| `v1/app/aiops/canary.py` | 9 | Python | `set_replicas` |
| `v1/app/aiops/canary.py` | 16 | Python | `rollout` |
| `v1/app/aiops/canary.py` | 29 | Python | `check` |
| `v1/app/aiops/changepoint.py` | 5 | Python | `__init__` |
| `v1/app/aiops/changepoint.py` | 8 | Python | `add` |
| `v1/app/aiops/changepoint.py` | 11 | Python | `mean` |
| `v1/app/aiops/changepoint.py` | 18 | Python | `__init__` |
| `v1/app/aiops/changepoint.py` | 24 | Python | `update` |
| `v1/app/aiops/compliance/engine.py` | 8 | Python | `__init__` |
| `v1/app/aiops/compliance/engine.py` | 13 | Python | `check` |
| `v1/app/aiops/coordination.py` | 6 | Python | `__init__` |
| `v1/app/aiops/coordination.py` | 9 | Python | `acquire` |
| `v1/app/aiops/coordination.py` | 13 | Python | `done` |
| `v1/app/aiops/coordination.py` | 16 | Python | `dedupe` |
| `v1/app/aiops/cost_model.py` | 10 | Python | `estimate` |
| `v1/app/aiops/cost_model.py` | 13 | Python | `budget_penalty` |
| `v1/app/aiops/counterfactual.py` | 1 | Python | `ips` |
| `v1/app/aiops/counterfactual.py` | 9 | Python | `doubly_robust` |
| `v1/app/aiops/detector.py` | 5 | Python | `__init__` |
| `v1/app/aiops/detector.py` | 10 | Python | `ingest` |
| `v1/app/aiops/detector.py` | 15 | Python | `anomalies` |
| `v1/app/aiops/diagnoser.py` | 5 | Python | `diagnose` |
| `v1/app/aiops/dreamer.py` | 8 | Python | `__init__` |
| `v1/app/aiops/dreamer.py` | 16 | Python | `forward` |
| `v1/app/aiops/dreamer.py` | 21 | Python | `__init__` |
| `v1/app/aiops/dreamer.py` | 29 | Python | `forward` |
| `v1/app/aiops/drift.py` | 4 | Python | `ks_stat` |
| `v1/app/aiops/drift.py` | 23 | Python | `__init__` |
| `v1/app/aiops/drift.py` | 26 | Python | `drifted` |
| `v1/app/aiops/enterprise_pipeline.py` | 16 | Python | `process` |
| `v1/app/aiops/evaluator.py` | 2 | Python | `better` |
| `v1/app/aiops/executor.py` | 4 | Python | `_sh` |
| `v1/app/aiops/executor.py` | 9 | Python | `scale_api` |
| `v1/app/aiops/executor.py` | 12 | Python | `restart_pods` |
| `v1/app/aiops/executor.py` | 15 | Python | `rollback_deploy` |
| `v1/app/aiops/executor.py` | 18 | Python | `scale_consumers` |
| `v1/app/aiops/executor.py` | 21 | Python | `restart_kafka` |
| `v1/app/aiops/fairness/monitor.py` | 5 | Python | `check_bias` |
| `v1/app/aiops/fairness/pricing_guard.py` | 2 | Python | `clamp` |
| `v1/app/aiops/features.py` | 6 | Python | `__init__` |
| `v1/app/aiops/features.py` | 10 | Python | `add` |
| `v1/app/aiops/features.py` | 13 | Python | `mean_std` |
| `v1/app/aiops/features_auto.py` | 8 | Python | `__init__` |
| `v1/app/aiops/features_auto.py` | 16 | Python | `forward` |
| `v1/app/aiops/features_auto.py` | 20 | Python | `feature_importance` |
| `v1/app/aiops/features_ctx.py` | 4 | Python | `feature_dict` |
| `v1/app/aiops/features_ctx.py` | 17 | Python | `build_context` |
| `v1/app/aiops/governance/policy_guard.py` | 5 | Python | `__init__` |
| `v1/app/aiops/governance/policy_guard.py` | 9 | Python | `allow` |
| `v1/app/aiops/governance/voting.py` | 5 | Python | `__init__` |
| `v1/app/aiops/governance/voting.py` | 10 | Python | `vote` |
| `v1/app/aiops/governance/voting.py` | 13 | Python | `result` |
| `v1/app/aiops/market/clearing.py` | 18 | Python | `clear_market` |
| `v1/app/aiops/market/exchange.py` | 15 | Python | `__init__` |
| `v1/app/aiops/market/exchange.py` | 19 | Python | `run_once` |
| `v1/app/aiops/market/exchange.py` | 36 | Python | `execute_trades` |
| `v1/app/aiops/market/orderbook.py` | 16 | Python | `__init__` |
| `v1/app/aiops/market/orderbook.py` | 20 | Python | `add_bid` |
| `v1/app/aiops/market/orderbook.py` | 26 | Python | `add_ask` |
| `v1/app/aiops/market/orderbook.py` | 32 | Python | `get` |
| `v1/app/aiops/market/orderbook.py` | 37 | Python | `clear` |
| `v1/app/aiops/market/pricing.py` | 4 | Python | `dynamic_price` |
| `v1/app/aiops/negotiation.py` | 15 | Python | `__init__` |
| `v1/app/aiops/negotiation.py` | 18 | Python | `negotiate` |
| `v1/app/aiops/online_learning.py` | 10 | Python | `__init__` |
| `v1/app/aiops/online_learning.py` | 13 | Python | `log` |
| `v1/app/aiops/online_learning.py` | 18 | Python | `replay` |
| `v1/app/aiops/orchestrator.py` | 12 | Python | `__init__` |
| `v1/app/aiops/orchestrator.py` | 20 | Python | `ingest_metrics` |
| `v1/app/aiops/orchestrator.py` | 23 | Python | `step` |
| `v1/app/aiops/orchestrator.py` | 48 | Python | `run_loop` |
| `v1/app/aiops/orchestrator_v11.py` | 10 | Python | `__init__` |
| `v1/app/aiops/orchestrator_v11.py` | 15 | Python | `_gates_passed` |
| `v1/app/aiops/orchestrator_v11.py` | 26 | Python | `step` |
| `v1/app/aiops/orchestrator_v2.py` | 12 | Python | `__init__` |
| `v1/app/aiops/orchestrator_v2.py` | 23 | Python | `step` |
| `v1/app/aiops/orchestrator_v3.py` | 12 | Python | `__init__` |
| `v1/app/aiops/orchestrator_v3.py` | 23 | Python | `step` |
| `v1/app/aiops/orchestrator_v3.py` | 92 | Python | `_reward` |
| `v1/app/aiops/orchestrator_v4.py` | 18 | Python | `__init__` |
| `v1/app/aiops/orchestrator_v4.py` | 34 | Python | `_context` |
| `v1/app/aiops/orchestrator_v4.py` | 46 | Python | `step` |
| `v1/app/aiops/orchestrator_v4.py` | 112 | Python | `_reward` |
| `v1/app/aiops/orchestrator_v5.py` | 19 | Python | `__init__` |
| `v1/app/aiops/orchestrator_v5.py` | 38 | Python | `step` |
| `v1/app/aiops/orchestrator_v5.py` | 100 | Python | `_reward` |
| `v1/app/aiops/patcher.py` | 7 | Python | `apply_patch` |
| `v1/app/aiops/planner.py` | 11 | Python | `plan` |
| `v1/app/aiops/policy.py` | 12 | Python | `__init__` |
| `v1/app/aiops/policy.py` | 15 | Python | `register` |
| `v1/app/aiops/policy.py` | 18 | Python | `get` |
| `v1/app/aiops/priors_learned.py` | 10 | Python | `__init__` |
| `v1/app/aiops/priors_learned.py` | 16 | Python | `ingest_incident` |
| `v1/app/aiops/priors_learned.py` | 22 | Python | `recompute` |
| `v1/app/aiops/quorum.py` | 12 | Python | `__init__` |
| `v1/app/aiops/quorum.py` | 17 | Python | `vote` |
| `v1/app/aiops/quorum.py` | 24 | Python | `approved` |
| `v1/app/aiops/rewrite.py` | 10 | Python | `_coerce_scalar` |
| `v1/app/aiops/rewrite.py` | 22 | Python | `_simple_yaml_load` |
| `v1/app/aiops/rewrite.py` | 43 | Python | `propose_patch` |
| `v1/app/aiops/risk/circuit.py` | 2 | Python | `__init__` |
| `v1/app/aiops/risk/circuit.py` | 6 | Python | `record` |
| `v1/app/aiops/risk/circuit.py` | 12 | Python | `open` |
| `v1/app/aiops/risk/exposure.py` | 2 | Python | `check` |
| `v1/app/aiops/risk/limits.py` | 5 | Python | `allow` |
| `v1/app/aiops/rl_planner.py` | 14 | Python | `__init__` |
| `v1/app/aiops/rl_planner.py` | 20 | Python | `_state` |
| `v1/app/aiops/rl_planner.py` | 23 | Python | `select` |
| `v1/app/aiops/rl_planner.py` | 29 | Python | `update` |
| `v1/app/aiops/rssm.py` | 10 | Python | `__init__` |
| `v1/app/aiops/rssm.py` | 18 | Python | `forward` |
| `v1/app/aiops/safeguards.py` | 8 | Python | `__init__` |
| `v1/app/aiops/safeguards.py` | 15 | Python | `_refresh_window` |
| `v1/app/aiops/safeguards.py` | 21 | Python | `allow` |
| `v1/app/aiops/scaler_cloud.py` | 5 | Python | `__init__` |
| `v1/app/aiops/scaler_cloud.py` | 8 | Python | `scale_cluster` |
| `v1/app/aiops/scaler_cloud.py` | 22 | Python | `shift_traffic` |
| `v1/app/aiops/shadow.py` | 23 | Python | `simulate` |
| `v1/app/aiops/simulator.py` | 5 | Python | `simulate_step` |
| `v1/app/aiops/simulator.py` | 23 | Python | `rollout` |
| `v1/app/aiops/sla.py` | 2 | Python | `__init__` |
| `v1/app/aiops/sla.py` | 9 | Python | `get` |
| `v1/app/aiops/sla.py` | 14 | Python | `allow` |
| `v1/app/aiops/slo_planner.py` | 20 | Python | `choose` |
| `v1/app/aiops/state.py` | 2 | Python | `__init__` |
| `v1/app/aiops/state.py` | 6 | Python | `transition` |
| `v1/app/aiops/state_store.py` | 9 | Python | `lease` |
| `v1/app/aiops/state_store.py` | 14 | Python | `release` |
| `v1/app/aiops/state_store.py` | 18 | Python | `idem_key` |
| `v1/app/aiops/tenant_policy.py` | 7 | Python | `__init__` |
| `v1/app/aiops/tenant_policy.py` | 10 | Python | `get` |
| `v1/app/aiops/tenant_policy.py` | 13 | Python | `snapshot` |
| `v1/app/aiops/tenant_policy.py` | 16 | Python | `load` |
| `v1/app/api/deps.py` | 7 | Python | `current_user` |
| `v1/app/api/routes_admin.py` | 12 | Python | `toggle_strategy` |
| `v1/app/api/routes_admin.py` | 19 | Python | `assign_role` |
| `v1/app/api/routes_admin.py` | 25 | Python | `list_roles` |
| `v1/app/api/routes_admin.py` | 31 | Python | `audit_log` |
| `v1/app/api/routes_admin.py` | 37 | Python | `update_slo` |
| `v1/app/api/routes_admin.py` | 43 | Python | `slo_dashboard` |
| `v1/app/api/routes_admin.py` | 49 | Python | `upsert_growth_campaign` |
| `v1/app/api/routes_admin.py` | 55 | Python | `growth_actions` |
| `v1/app/api/routes_admin.py` | 61 | Python | `growth_analytics` |
| `v1/app/api/routes_aiops.py` | 12 | Python | `kill` |
| `v1/app/api/routes_aiops.py` | 19 | Python | `set_limit` |
| `v1/app/api/routes_auth.py` | 9 | Python | `login` |
| `v1/app/api/routes_god.py` | 16 | Python | `kill` |
| `v1/app/api/routes_god.py` | 23 | Python | `set_budget` |
| `v1/app/api/routes_god.py` | 31 | Python | `set_circuit` |
| `v1/app/api/routes_user.py` | 16 | Python | `portfolio` |
| `v1/app/api/routes_user.py` | 26 | Python | `onboarding_stage` |
| `v1/app/arb/detector.py` | 4 | Python | `best_pair` |
| `v1/app/arb/detector.py` | 19 | Python | `net_edge` |
| `v1/app/arb/detector.py` | 24 | Python | `is_actionable` |
| `v1/app/arb/engine.py` | 13 | Python | `__init__` |
| `v1/app/arb/engine.py` | 18 | Python | `run` |
| `v1/app/arb/executor.py` | 18 | Python | `__init__` |
| `v1/app/arb/executor.py` | 22 | Python | `run` |
| `v1/app/auth/jwt.py` | 10 | Python | `sign` |
| `v1/app/auth/jwt.py` | 16 | Python | `verify` |
| `v1/app/auth/rbac.py` | 7 | Python | `require` |
| `v1/app/auth/rbac.py` | 8 | Python | `_check` |
| `v1/app/billing/engine.py` | 6 | Python | `invoice` |
| `v1/app/core/event_bus.py` | 6 | Python | `__init__` |
| `v1/app/core/event_bus.py` | 10 | Python | `start` |
| `v1/app/core/event_bus.py` | 13 | Python | `publish` |
| `v1/app/core/event_bus.py` | 16 | Python | `stop` |
| `v1/app/core/event_bus.py` | 21 | Python | `__init__` |
| `v1/app/core/event_bus.py` | 28 | Python | `start` |
| `v1/app/core/event_bus.py` | 31 | Python | `listen` |
| `v1/app/core/orchestrator.py` | 10 | Python | `__init__` |
| `v1/app/core/orchestrator.py` | 16 | Python | `run` |
| `v1/app/data/market.py` | 6 | Python | `get_data` |
| `v1/app/data/producer.py` | 6 | Python | `__init__` |
| `v1/app/data/producer.py` | 10 | Python | `run` |
| `v1/app/data/producer_multi.py` | 17 | Python | `__init__` |
| `v1/app/data/producer_multi.py` | 20 | Python | `_binance` |
| `v1/app/data/producer_multi.py` | 27 | Python | `_okx` |
| `v1/app/data/producer_multi.py` | 37 | Python | `_bybit` |
| `v1/app/data/producer_multi.py` | 47 | Python | `run` |
| `v1/app/data/stream.py` | 7 | Python | `__init__` |
| `v1/app/data/stream.py` | 10 | Python | `stream` |
| `v1/app/enterprise/saas.py` | 22 | Python | `__init__` |
| `v1/app/enterprise/saas.py` | 27 | Python | `assign_role` |
| `v1/app/enterprise/saas.py` | 33 | Python | `list_roles` |
| `v1/app/enterprise/saas.py` | 37 | Python | `add_audit_event` |
| `v1/app/enterprise/saas.py` | 42 | Python | `get_audit_events` |
| `v1/app/enterprise/saas.py` | 46 | Python | `upsert_slo` |
| `v1/app/enterprise/saas.py` | 56 | Python | `slo_dashboard` |
| `v1/app/enterprise/saas.py` | 60 | Python | `_append_audit` |
| `v1/app/execution/base.py` | 7 | Python | `get_best_bid_ask` |
| `v1/app/execution/base.py` | 11 | Python | `place_market` |
| `v1/app/execution/base.py` | 15 | Python | `fee_rate` |
| `v1/app/execution/binance.py` | 11 | Python | `__init__` |
| `v1/app/execution/binance.py` | 17 | Python | `sign` |
| `v1/app/execution/binance.py` | 20 | Python | `get_best_bid_ask` |
| `v1/app/execution/binance.py` | 27 | Python | `place_market` |
| `v1/app/execution/binance.py` | 41 | Python | `fee_rate` |
| `v1/app/execution/binance.py` | 44 | Python | `place_order` |
| `v1/app/execution/bybit.py` | 7 | Python | `__init__` |
| `v1/app/execution/bybit.py` | 13 | Python | `get_best_bid_ask` |
| `v1/app/execution/bybit.py` | 20 | Python | `place_market` |
| `v1/app/execution/bybit.py` | 29 | Python | `fee_rate` |
| `v1/app/execution/idempotency.py` | 8 | Python | `make_key` |
| `v1/app/execution/okx.py` | 9 | Python | `__init__` |
| `v1/app/execution/okx.py` | 16 | Python | `get_best_bid_ask` |
| `v1/app/execution/okx.py` | 23 | Python | `place_market` |
| `v1/app/execution/okx.py` | 44 | Python | `fee_rate` |
| `v1/app/feature_store/offline.py` | 10 | Python | `__init__` |
| `v1/app/feature_store/offline.py` | 13 | Python | `write` |
| `v1/app/feature_store/online.py` | 10 | Python | `__init__` |
| `v1/app/feature_store/online.py` | 13 | Python | `put` |
| `v1/app/feature_store/online.py` | 18 | Python | `get` |
| `v1/app/features/builder.py` | 5 | Python | `build` |
| `v1/app/growth/automation.py` | 18 | Python | `__init__` |
| `v1/app/growth/automation.py` | 22 | Python | `track_stage` |
| `v1/app/growth/automation.py` | 28 | Python | `upsert_campaign` |
| `v1/app/growth/automation.py` | 34 | Python | `generate_actions` |
| `v1/app/growth/automation.py` | 46 | Python | `analytics` |
| `v1/app/guardrails/budgets.py` | 5 | Python | `__init__` |
| `v1/app/guardrails/budgets.py` | 9 | Python | `set` |
| `v1/app/guardrails/budgets.py` | 12 | Python | `get` |
| `v1/app/guardrails/circuit.py` | 5 | Python | `__init__` |
| `v1/app/guardrails/circuit.py` | 9 | Python | `set` |
| `v1/app/guardrails/circuit.py` | 12 | Python | `get` |
| `v1/app/guardrails/circuit.py` | 15 | Python | `tripped` |
| `v1/app/guardrails/kill_switch.py` | 2 | Python | `__init__` |
| `v1/app/guardrails/kill_switch.py` | 5 | Python | `set` |
| `v1/app/guardrails/kill_switch.py` | 8 | Python | `is_on` |
| `v1/app/infra/redis_client.py` | 7 | Python | `__init__` |
| `v1/app/infra/redis_client.py` | 10 | Python | `set_price` |
| `v1/app/infra/redis_client.py` | 13 | Python | `get_price` |
| `v1/app/infra/redis_client.py` | 16 | Python | `set_book` |
| `v1/app/infra/redis_client.py` | 20 | Python | `get_books` |
| `v1/app/intelligence/funding.py` | 1 | Python | `funding_signal` |
| `v1/app/intelligence/fusion.py` | 1 | Python | `fuse` |
| `v1/app/intelligence/portfolio.py` | 1 | Python | `size` |
| `v1/app/limits/quotas.py` | 7 | Python | `__init__` |
| `v1/app/limits/quotas.py` | 11 | Python | `consume` |
| `v1/app/limits/rate_limit.py` | 11 | Python | `__init__` |
| `v1/app/limits/rate_limit.py` | 14 | Python | `allow` |
| `v1/app/main.py` | 10 | Python | `startup` |
| `v1/app/main.py` | 15 | Python | `health` |
| `v1/app/metering/usage.py` | 7 | Python | `__init__` |
| `v1/app/metering/usage.py` | 11 | Python | `inc_req` |
| `v1/app/metering/usage.py` | 14 | Python | `inc_trade` |
| `v1/app/risk/engine.py` | 2 | Python | `__init__` |
| `v1/app/risk/engine.py` | 8 | Python | `size` |
| `v1/app/risk/engine.py` | 11 | Python | `check_drawdown` |
| `v1/app/signal/consumer.py` | 9 | Python | `__init__` |
| `v1/app/signal/consumer.py` | 17 | Python | `run` |
| `v1/app/signal/engine.py` | 5 | Python | `compute` |
| `v1/app/slo/metrics.py` | 7 | Python | `init_metrics` |
| `v1/app/slo/metrics.py` | 11 | Python | `observe` |
| `v1/app/storage/clickhouse.py` | 7 | Python | `__init__` |
| `v1/app/storage/clickhouse.py` | 10 | Python | `insert` |
| `v1/app/tenancy/context.py` | 6 | Python | `set_tenant` |
| `v1/app/tenancy/context.py` | 10 | Python | `get_tenant` |
| `v1/app/tenancy/middleware.py` | 7 | Python | `dispatch` |
| `v1/app/utils/logger.py` | 6 | Python | `get_logger` |
| `v1/infra/hft/ebpf/exporter.py` | 7 | Python | `main` |
| `v1/infra/hft/ebpf/run_probe.py` | 4 | Python | `main` |
| `v1/infra/scripts/ui.py` | 15 | Python | `_run_deploy_script` |
| `v1/infra/scripts/ui.py` | 28 | Python | `deploy` |
| `v1/infra/scripts/ui.py` | 42 | Python | `status` |
| `v1/src/zcyptobot/__main__.py` | 4 | Python | `main` |
| `v1/src/zcyptobot/core/arbitrage.py` | 23 | Python | `find` |
| `v1/src/zcyptobot/core/arbitrage_v6.py` | 26 | Python | `__init__` |
| `v1/src/zcyptobot/core/arbitrage_v6.py` | 31 | Python | `plan` |
| `v1/src/zcyptobot/core/event_bus.py` | 13 | Python | `__init__` |
| `v1/src/zcyptobot/core/event_bus.py` | 18 | Python | `start` |
| `v1/src/zcyptobot/core/event_bus.py` | 21 | Python | `publish` |
| `v1/src/zcyptobot/core/event_bus.py` | 24 | Python | `consume` |
| `v1/src/zcyptobot/core/event_bus.py` | 39 | Python | `stop` |
| `v1/src/zcyptobot/core/infra_v7.py` | 19 | Python | `terraform_vars` |
| `v1/src/zcyptobot/core/infra_v7.py` | 28 | Python | `helm_values` |
| `v1/src/zcyptobot/core/orchestrator_v2.py` | 18 | Python | `__init__` |
| `v1/src/zcyptobot/core/orchestrator_v2.py` | 24 | Python | `process` |
| `v1/src/zcyptobot/core/orchestrator_v2.py` | 39 | Python | `run` |
| `v1/src/zcyptobot/core/orderbook.py` | 19 | Python | `snapshot` |
| `v1/src/zcyptobot/core/orderbook.py` | 37 | Python | `spread_bps` |
| `v1/src/zcyptobot/core/orderbook.py` | 42 | Python | `imbalance` |
| `v1/src/zcyptobot/core/rl_cluster.py` | 17 | Python | `resources` |
| `v1/src/zcyptobot/core/rl_v5.py` | 20 | Python | `__init__` |
| `v1/src/zcyptobot/core/rl_v5.py` | 24 | Python | `score` |
| `v1/src/zcyptobot/core/rl_v5.py` | 27 | Python | `action` |
| `v1/src/zcyptobot/core/streams.py` | 13 | Python | `__init__` |
| `v1/src/zcyptobot/core/streams.py` | 17 | Python | `publish` |
| `v1/src/zcyptobot/core/streams.py` | 21 | Python | `consume_group` |
| `v1/src/zcyptobot/core/streams.py` | 48 | Python | `ack` |
| `v1/src/zcyptobot/core/streams.py` | 51 | Python | `close` |
| `v1/src/zcyptobot/data/stream.py` | 10 | Python | `__init__` |
| `v1/src/zcyptobot/data/stream.py` | 14 | Python | `listen` |
| `v1/src/zcyptobot/engines.py` | 11 | Python | `__init__` |
| `v1/src/zcyptobot/engines.py` | 14 | Python | `accept` |
| `v1/src/zcyptobot/engines.py` | 25 | Python | `__init__` |
| `v1/src/zcyptobot/engines.py` | 34 | Python | `on_tick` |
| `v1/src/zcyptobot/engines.py` | 68 | Python | `__init__` |
| `v1/src/zcyptobot/engines.py` | 73 | Python | `market_value` |
| `v1/src/zcyptobot/engines.py` | 81 | Python | `__init__` |
| `v1/src/zcyptobot/engines.py` | 84 | Python | `propose_order` |
| `v1/src/zcyptobot/engines.py` | 112 | Python | `__init__` |
| `v1/src/zcyptobot/engines.py` | 115 | Python | `execute` |
| `v1/src/zcyptobot/execution/base.py` | 5 | Python | `place_order` |
| `v1/src/zcyptobot/execution/binance.py` | 13 | Python | `__init__` |
| `v1/src/zcyptobot/execution/binance.py` | 18 | Python | `_sign` |
| `v1/src/zcyptobot/execution/binance.py` | 21 | Python | `place_order` |
| `v1/src/zcyptobot/orchestrator.py` | 22 | Python | `__init__` |
| `v1/src/zcyptobot/orchestrator.py` | 30 | Python | `run_cycle` |
| `v1/src/zcyptobot/pipeline.py` | 18 | Python | `__init__` |
| `v1/src/zcyptobot/pipeline.py` | 27 | Python | `_apply_fill` |
| `v1/src/zcyptobot/pipeline.py` | 43 | Python | `on_tick` |
| `v1/src/zcyptobot/risk/engine.py` | 5 | Python | `__init__` |
| `v1/src/zcyptobot/risk/engine.py` | 11 | Python | `size` |
| `v1/src/zcyptobot/risk/engine.py` | 14 | Python | `check_drawdown` |
| `v1/src/zcyptobot/services.py` | 16 | Python | `__init__` |
| `v1/src/zcyptobot/services.py` | 19 | Python | `next_tick` |
| `v1/src/zcyptobot/services.py` | 32 | Python | `__init__` |
| `v1/src/zcyptobot/services.py` | 35 | Python | `trending_symbols` |
| `v1/src/zcyptobot/signal/engine.py` | 9 | Python | `compute` |
| `v1/src/zcyptobot/simulator.py` | 18 | Python | `generate_ticks` |
| `v1/src/zcyptobot/simulator.py` | 33 | Python | `run_simulation` |
| `v1/src/zcyptobot/v14/future_stack.py` | 35 | Python | `__init__` |
| `v1/src/zcyptobot/v14/future_stack.py` | 40 | Python | `build_hermetic` |
| `v1/src/zcyptobot/v14/future_stack.py` | 44 | Python | `attest_sbom` |
| `v1/src/zcyptobot/v14/future_stack.py` | 47 | Python | `set_dependency_trust` |
| `v1/src/zcyptobot/v14/future_stack.py` | 50 | Python | `risk_gate` |
| `v1/src/zcyptobot/v14/future_stack.py` | 53 | Python | `register_artifact` |
| `v1/src/zcyptobot/v14/future_stack.py` | 60 | Python | `enforce_provenance` |
| `v1/src/zcyptobot/v14/future_stack.py` | 67 | Python | `__init__` |
| `v1/src/zcyptobot/v14/future_stack.py` | 71 | Python | `sync_feature_store` |
| `v1/src/zcyptobot/v14/future_stack.py` | 75 | Python | `infer_realtime` |
| `v1/src/zcyptobot/v14/future_stack.py` | 80 | Python | `world_model_plan` |
| `v1/src/zcyptobot/v14/future_stack.py` | 83 | Python | `negotiate_agents` |
| `v1/src/zcyptobot/v14/future_stack.py` | 86 | Python | `automl_train_and_deploy` |
| `v1/src/zcyptobot/v14/future_stack.py` | 90 | Python | `causal_uplift_gate` |
| `v1/src/zcyptobot/v14/future_stack.py` | 95 | Python | `__init__` |
| `v1/src/zcyptobot/v14/future_stack.py` | 99 | Python | `set_budget` |
| `v1/src/zcyptobot/v14/future_stack.py` | 102 | Python | `record_usage` |
| `v1/src/zcyptobot/v14/future_stack.py` | 105 | Python | `enforce_throttling` |
| `v1/src/zcyptobot/v14/future_stack.py` | 108 | Python | `revenue_metrics` |
| `v1/src/zcyptobot/v14/future_stack.py` | 111 | Python | `dynamic_price` |
| `v1/src/zcyptobot/v14/future_stack.py` | 116 | Python | `evaluate_error_budget` |
| `v1/src/zcyptobot/v14/future_stack.py` | 119 | Python | `trace_context` |
| `v1/src/zcyptobot/v14/future_stack.py` | 122 | Python | `inject_chaos` |
| `v1/src/zcyptobot/v14/future_stack.py` | 125 | Python | `auto_remediate` |
| `v1/src/zcyptobot/v14/future_stack.py` | 128 | Python | `forecast_capacity` |
| `v1/src/zcyptobot/v14/future_stack.py` | 135 | Python | `__init__` |
| `v1/src/zcyptobot/v14/future_stack.py` | 138 | Python | `record_attestation` |
| `v1/src/zcyptobot/v14/future_stack.py` | 141 | Python | `map_regulation` |
| `v1/src/zcyptobot/v14/future_stack.py` | 144 | Python | `classify_data` |
| `v1/src/zcyptobot/v14/future_stack.py` | 149 | Python | `provision_service` |
| `v1/src/zcyptobot/v14/future_stack.py` | 152 | Python | `publish_sdk` |
| `v1/src/zcyptobot/v14/future_stack.py` | 155 | Python | `contract_test` |
| `v1/src/zcyptobot/v14/future_stack.py` | 160 | Python | `place_workload` |
| `v1/src/zcyptobot/v14/future_stack.py` | 163 | Python | `orchestrate_spot` |
| `v1/src/zcyptobot/v14/future_stack.py` | 166 | Python | `schedule_gpu` |
| `v1/src/zcyptobot/v14/future_stack.py` | 170 | Python | `storage_tier` |
| `v1/src/zcyptobot/v14/future_stack.py` | 179 | Python | `register_plugin` |
| `v1/src/zcyptobot/v14/future_stack.py` | 182 | Python | `list_data_asset` |
| `v1/src/zcyptobot/v14/future_stack.py` | 185 | Python | `list_compute_asset` |
| `v1/src/zcyptobot/v14/future_stack.py` | 190 | Python | `self_heal` |
| `v1/src/zcyptobot/v14/future_stack.py` | 193 | Python | `self_optimize` |
| `v1/src/zcyptobot/v14/future_stack.py` | 196 | Python | `run_agent` |
| `v1/src/zcyptobot/v14/future_stack.py` | 199 | Python | `ai_refactor` |
| `v1/src/zcyptobot/v14/future_stack.py` | 204 | Python | `dao_vote` |
| `v1/src/zcyptobot/v14/future_stack.py` | 207 | Python | `policy_consensus` |
| `v1/src/zcyptobot/v14/future_stack.py` | 212 | Python | `treaty_route` |
| `v1/src/zcyptobot/v14/future_stack.py` | 217 | Python | `zk_proof` |
| `v1/src/zcyptobot/v14/future_stack.py` | 220 | Python | `fhe_compute` |
| `v1/src/zcyptobot/v14/future_stack.py` | 223 | Python | `federated_round` |
| `v1/src/zcyptobot/v14/future_stack.py` | 226 | Python | `quantum_safe_cipher` |
| `v1/src/zcyptobot/v14/future_stack.py` | 233 | Python | `__init__` |
| `v1/src/zcyptobot/v14/future_stack.py` | 246 | Python | `bootstrap` |
| `v1/src/zcyptobot/v4/governance.py` | 16 | Python | `__init__` |
| `v1/src/zcyptobot/v4/governance.py` | 20 | Python | `requires_human_approval` |
| `v1/src/zcyptobot/v4/portfolio.py` | 16 | Python | `allocate` |
| `v1/src/zcyptobot/v4/surveillance.py` | 16 | Python | `detect_quote_instability` |
| `v1/tests/test_aiops_v11.py` | 6 | Python | `__init__` |
| `v1/tests/test_aiops_v11.py` | 9 | Python | `propose` |
| `v1/tests/test_aiops_v11.py` | 13 | Python | `test_negotiation_filters_constraints` |
| `v1/tests/test_aiops_v11.py` | 23 | Python | `test_orchestrator_rewrite_requires_gates` |
| `v1/tests/test_aiops_v11.py` | 30 | Python | `test_orchestrator_rewrite_when_ready` |
| `v1/tests/test_aiops_v12_market_governance.py` | 6 | Python | `test_double_auction_clears_when_prices_cross` |
| `v1/tests/test_aiops_v12_market_governance.py` | 15 | Python | `test_weighted_vote_threshold` |
| `v1/tests/test_aiops_v12_market_governance.py` | 22 | Python | `test_policy_guard_cost_ceiling` |
| `v1/tests/test_aiops_v3.py` | 5 | Python | `test_bayesian_rca_returns_ranked_probabilities` |
| `v1/tests/test_aiops_v3.py` | 14 | Python | `test_rl_planner_updates_q_value` |
| `v1/tests/test_aiops_v6_v7_nonstationary.py` | 10 | Python | `test_lints_ns_discounted_update_changes_state` |
| `v1/tests/test_aiops_v6_v7_nonstationary.py` | 18 | Python | `test_hierarchical_ts_select_and_update` |
| `v1/tests/test_aiops_v6_v7_nonstationary.py` | 26 | Python | `test_counterfactual_estimators_return_float` |
| `v1/tests/test_aiops_v6_v7_nonstationary.py` | 33 | Python | `test_policy_registry_roundtrip_and_drift_monitor` |
| `v1/tests/test_enterprise_growth_upgrade.py` | 5 | Python | `test_multi_tenant_role_isolation_and_audit` |
| `v1/tests/test_enterprise_growth_upgrade.py` | 20 | Python | `test_slo_dashboard_and_growth_analytics` |
| `v1/tests/test_institutional_upgrade.py` | 6 | Python | `test_orderbook_depth_model_features` |
| `v1/tests/test_institutional_upgrade.py` | 15 | Python | `test_multi_exchange_arbitrage_detects_net_edge` |
| `v1/tests/test_institutional_upgrade.py` | 29 | Python | `test_rl_training_planner_resources` |
| `v1/tests/test_orchestrator.py` | 4 | Python | `test_orchestrator_cycle_produces_stats` |
| `v1/tests/test_orchestrator.py` | 12 | Python | `test_orchestrator_can_run_multiple_cycles` |
| `v1/tests/test_pipeline.py` | 7 | Python | `test_bot_accepts_valid_ticks_and_updates_equity` |
| `v1/tests/test_pipeline.py` | 16 | Python | `test_bot_rejects_invalid_symbol` |
| `v1/tests/test_pipeline.py` | 23 | Python | `test_simulation_runs_end_to_end` |
| `v1/tests/test_v2_core.py` | 5 | Python | `test_signal_engine_long` |
| `v1/tests/test_v2_core.py` | 10 | Python | `test_risk_engine_size_and_drawdown` |
| `v1/tests/test_v4_stack.py` | 6 | Python | `test_approval_policy_threshold` |
| `v1/tests/test_v4_stack.py` | 12 | Python | `test_portfolio_allocator_respects_gross_limit` |
| `v1/tests/test_v4_stack.py` | 25 | Python | `test_surveillance_detects_quote_instability` |
| `v1/tests/test_v5_v6_v7_upgrade.py` | 7 | Python | `test_v5_rl_policy_actions` |
| `v1/tests/test_v5_v6_v7_upgrade.py` | 13 | Python | `test_v6_arbitrage_router_builds_plan` |
| `v1/tests/test_v5_v6_v7_upgrade.py` | 31 | Python | `test_v7_infra_blueprint_materializes_configs` |
| `v1/zeaz-quant/app/data/collector.py` | 7 | Python | `get_market` |
| `v1/zeaz-quant/app/execution/engine.py` | 2 | Python | `execute` |
| `v1/zeaz-quant/app/features/builder.py` | 5 | Python | `build` |
| `v1/zeaz-quant/app/main.py` | 21 | Python | `loop` |
| `v1/zeaz-quant/app/risk/manager.py` | 2 | Python | `__init__` |
| `v1/zeaz-quant/app/risk/manager.py` | 7 | Python | `validate` |
| `v1/zeaz-quant/app/risk/manager.py` | 16 | Python | `update` |
| `v1/zeaz-quant/app/signal/engine.py` | 2 | Python | `generate` |
| `v1/zeaz-quant/app/storage/db.py` | 5 | Python | `__init__` |
| `v1/zeaz-quant/app/storage/db.py` | 11 | Python | `save` |
| `war_simulation/metrics.py` | 1 | Python | `kpis` |
| `war_simulation/response_engine.py` | 4 | Python | `respond` |
| `war_simulation/strategies.py` | 1 | Python | `choose` |
| `warfare/engine.py` | 4 | Python | `step` |
| `warfare/metrics.py` | 1 | Python | `kpis` |
| `warfare/response.py` | 1 | Python | `respond` |
| `zcyptobot-hyperscale/scripts/bootstrap.py` | 11 | Python | `write` |
| `zcyptobot-hyperscale/scripts/bootstrap.py` | 16 | Python | `service_scaffold` |
| `zcyptobot-hyperscale/scripts/bootstrap.py` | 28 | Python | `infra_scaffold` |
| `zcyptobot-hyperscale/scripts/bootstrap.py` | 36 | Python | `observability_scaffold` |
| `zcyptobot-hyperscale/scripts/bootstrap.py` | 49 | Python | `ai_scaffold` |
| `zcyptobot-hyperscale/scripts/bootstrap.py` | 54 | Python | `main` |
| `zcyptobot-hyperscale/services/service_0001/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0001/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0002/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0002/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0003/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0003/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0004/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0004/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0005/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0005/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0006/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0006/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0007/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0007/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0008/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0008/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0009/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0009/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_001/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_001/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0010/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0010/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0011/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0011/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0012/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0012/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0013/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0013/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0014/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0014/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0015/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0015/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0016/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0016/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0017/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0017/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0018/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0018/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0019/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0019/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_002/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_002/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0020/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0020/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0021/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0021/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0022/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0022/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0023/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0023/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0024/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0024/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0025/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0025/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0026/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0026/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0027/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0027/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0028/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0028/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0029/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0029/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_003/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_003/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0030/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0030/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0031/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0031/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0032/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0032/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0033/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0033/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0034/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0034/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0035/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0035/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0036/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0036/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0037/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0037/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0038/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0038/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0039/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0039/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_004/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_004/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0040/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0040/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0041/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0041/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0042/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0042/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0043/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0043/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0044/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0044/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0045/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0045/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0046/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0046/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0047/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0047/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0048/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0048/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0049/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0049/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_005/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_005/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0050/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0050/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0051/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0051/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0052/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0052/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0053/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0053/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0054/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0054/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0055/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0055/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0056/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0056/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0057/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0057/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0058/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0058/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0059/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0059/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_006/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_006/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0060/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0060/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0061/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0061/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0062/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0062/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0063/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0063/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0064/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0064/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0065/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0065/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0066/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0066/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0067/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0067/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0068/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0068/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0069/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0069/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_007/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_007/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0070/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0070/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0071/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0071/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0072/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0072/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0073/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0073/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0074/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0074/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0075/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0075/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0076/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0076/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0077/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0077/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0078/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0078/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0079/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0079/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_008/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_008/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0080/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0080/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0081/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0081/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0082/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0082/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0083/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0083/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0084/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0084/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0085/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0085/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0086/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0086/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0087/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0087/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0088/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0088/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0089/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0089/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_009/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_009/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0090/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0090/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0091/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0091/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0092/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0092/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0093/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0093/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0094/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0094/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0095/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0095/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0096/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0096/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0097/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0097/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0098/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0098/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0099/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0099/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_010/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_010/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0100/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0100/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0101/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0101/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0102/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0102/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0103/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0103/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0104/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0104/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0105/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0105/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0106/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0106/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0107/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0107/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0108/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0108/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0109/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0109/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_011/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_011/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0110/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0110/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0111/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0111/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0112/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0112/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0113/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0113/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0114/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0114/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0115/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0115/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0116/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0116/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0117/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0117/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0118/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0118/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0119/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0119/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_012/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_012/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0120/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0120/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0121/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0121/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0122/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0122/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0123/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0123/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0124/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0124/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0125/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0125/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0126/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0126/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0127/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0127/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0128/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0128/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0129/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0129/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_013/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_013/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0130/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0130/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0131/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0131/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0132/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0132/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0133/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0133/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0134/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0134/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0135/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0135/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0136/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0136/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0137/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0137/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0138/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0138/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0139/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0139/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_014/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_014/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0140/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0140/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0141/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0141/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0142/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0142/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0143/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0143/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0144/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0144/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0145/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0145/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0146/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0146/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0147/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0147/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0148/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0148/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0149/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0149/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_015/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_015/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0150/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0150/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0151/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0151/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0152/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0152/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0153/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0153/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0154/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0154/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0155/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0155/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0156/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0156/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0157/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0157/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0158/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0158/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0159/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0159/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_016/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_016/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0160/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0160/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0161/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0161/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0162/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0162/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0163/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0163/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0164/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0164/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0165/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0165/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0166/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0166/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0167/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0167/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0168/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0168/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0169/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0169/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_017/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_017/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0170/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0170/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0171/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0171/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0172/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0172/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0173/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0173/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0174/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0174/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0175/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0175/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0176/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0176/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0177/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0177/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0178/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0178/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0179/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0179/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_018/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_018/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0180/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0180/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0181/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0181/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0182/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0182/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0183/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0183/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0184/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0184/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0185/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0185/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0186/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0186/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0187/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0187/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0188/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0188/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0189/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0189/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_019/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_019/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0190/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0190/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0191/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0191/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0192/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0192/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0193/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0193/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0194/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0194/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0195/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0195/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0196/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0196/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0197/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0197/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0198/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0198/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0199/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0199/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_020/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_020/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0200/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0200/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0201/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0201/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0202/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0202/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0203/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0203/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0204/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0204/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0205/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0205/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0206/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0206/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0207/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0207/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0208/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0208/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0209/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0209/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_021/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_021/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0210/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0210/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0211/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0211/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0212/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0212/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0213/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0213/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0214/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0214/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0215/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0215/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0216/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0216/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0217/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0217/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0218/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0218/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_0219/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0219/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_022/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_022/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_0220/src/main.py` | 6 | Python | `health` |
| `zcyptobot-hyperscale/services/service_0220/tests/test_basic.py` | 1 | Python | `test_health_payload_shape` |
| `zcyptobot-hyperscale/services/service_023/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_023/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_024/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_024/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_025/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_025/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_026/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_026/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_027/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_027/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_028/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_028/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_029/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_029/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_030/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_030/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_031/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_031/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_032/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_032/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_033/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_033/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_034/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_034/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_035/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_035/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_036/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_036/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_037/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_037/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_038/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_038/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_039/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_039/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_040/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_040/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_041/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_041/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_042/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_042/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_043/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_043/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_044/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_044/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_045/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_045/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_046/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_046/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_047/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_047/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_048/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_048/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_049/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_049/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_050/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_050/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_051/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_051/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_052/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_052/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_053/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_053/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_054/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_054/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_055/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_055/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_056/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_056/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_057/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_057/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_058/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_058/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_059/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_059/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_060/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_060/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_061/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_061/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_062/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_062/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_063/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_063/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_064/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_064/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_065/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_065/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_066/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_066/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_067/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_067/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_068/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_068/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_069/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_069/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_070/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_070/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_071/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_071/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_072/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_072/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_073/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_073/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_074/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_074/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_075/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_075/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_076/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_076/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_077/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_077/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_078/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_078/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_079/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_079/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_080/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_080/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_081/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_081/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_082/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_082/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_083/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_083/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_084/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_084/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_085/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_085/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_086/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_086/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_087/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_087/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_088/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_088/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_089/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_089/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_090/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_090/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_091/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_091/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_092/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_092/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_093/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_093/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_094/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_094/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_095/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_095/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_096/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_096/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_097/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_097/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_098/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_098/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_099/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_099/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/services/service_100/src/main.py` | 1 | Python | `handler` |
| `zcyptobot-hyperscale/services/service_100/tests/test_basic.py` | 1 | Python | `test_basic` |
| `zcyptobot-hyperscale/tests/test_smoke.py` | 1 | Python | `test_smoke` |
| `zeaz-enterprise/backend/auth_sso.py` | 12 | Python | `verify` |
| `zeaz-enterprise/backend/db.py` | 7 | Python | `__init__` |
| `zeaz-enterprise/backend/db.py` | 11 | Python | `__enter__` |
| `zeaz-enterprise/backend/db.py` | 16 | Python | `__exit__` |
| `zeaz-enterprise/backend/main.py` | 22 | Python | `get_state` |
| `zeaz-enterprise/backend/main.py` | 29 | Python | `ws` |
| `zeaz-enterprise/backend/rl/agents.py` | 2 | Python | `__init__` |
| `zeaz-enterprise/backend/rl/agents.py` | 5 | Python | `propose` |
| `zeaz-enterprise/backend/rl/agents.py` | 10 | Python | `allow` |
| `zeaz-enterprise/backend/rl/agents.py` | 15 | Python | `allow` |
| `zeaz-enterprise/backend/rl/agents.py` | 20 | Python | `decide` |
| `zeaz-enterprise/backend/rl/planner.py` | 7 | Python | `onehot` |
| `zeaz-enterprise/backend/rl/planner.py` | 14 | Python | `__init__` |
| `zeaz-enterprise/backend/rl/planner.py` | 17 | Python | `plan` |
| `zeaz-enterprise/backend/rl/world_model.py` | 6 | Python | `__init__` |
| `zeaz-enterprise/backend/rl/world_model.py` | 10 | Python | `forward` |
| `zeaz-enterprise/backend/stripe_billing.py` | 13 | Python | `_ensure_api_key` |
| `zeaz-enterprise/backend/stripe_billing.py` | 18 | Python | `create_customer` |
| `zeaz-enterprise/backend/stripe_billing.py` | 23 | Python | `create_subscription` |
| `zeaz-enterprise/backend/stripe_billing.py` | 28 | Python | `report_usage` |
| `zeaz-enterprise/backend/tenancy.py` | 6 | Python | `schema_for` |
| `zeaz-enterprise/backend/tenancy.py` | 12 | Python | `set_schema` |
| `zeaz-hyperscale/agents/evolver.py` | 9 | Python | `propose_policy` |
| `zeaz-hyperscale/agents/evolver.py` | 16 | Python | `propose_code_patch` |
| `zeaz-hyperscale/agents/evolver.py` | 24 | Python | `commit_patch` |
| `zeaz-hyperscale/agents/guardrails.py` | 2 | Python | `ips_dr_gate` |
| `zeaz-hyperscale/agents/guardrails.py` | 5 | Python | `cost_cap` |
| `zeaz-hyperscale/agents/guardrails.py` | 8 | Python | `quorum` |
| `zeaz-hyperscale/agents/patcher.py` | 6 | Python | `argocd_sync` |
| `zeaz-hyperscale/agents/scorer.py` | 1 | Python | `rank` |
| `zeaz-hyperscale/agents/validator.py` | 5 | Python | `run_unit` |
| `zeaz-hyperscale/agents/validator.py` | 9 | Python | `run_lint` |
| `zeaz-hyperscale/agents/validator.py` | 13 | Python | `simulate` |
| `zeaz-hyperscale/agents/validator.py` | 18 | Python | `validate` |
| `zeaz-live/backend/auth.py` | 6 | Python | `create` |
| `zeaz-live/backend/auth.py` | 10 | Python | `verify` |
| `zeaz-live/backend/billing.py` | 4 | Python | `track` |
| `zeaz-live/backend/billing.py` | 8 | Python | `get_usage` |
| `zeaz-live/backend/engine.py` | 7 | Python | `decide` |
| `zeaz-live/backend/feature_store/offline.py` | 6 | Python | `insert` |
| `zeaz-live/backend/feature_store/online.py` | 8 | Python | `put` |
| `zeaz-live/backend/feature_store/online.py` | 12 | Python | `get` |
| `zeaz-live/backend/i18n.py` | 50 | Python | `normalize_locale` |
| `zeaz-live/backend/i18n.py` | 59 | Python | `get_messages` |
| `zeaz-live/backend/main.py` | 26 | Python | `require_auth` |
| `zeaz-live/backend/main.py` | 32 | Python | `get_state` |
| `zeaz-live/backend/main.py` | 39 | Python | `metrics` |
| `zeaz-live/backend/main.py` | 46 | Python | `admin_usage` |
| `zeaz-live/backend/main.py` | 53 | Python | `ws` |
| `zeaz-live/backend/main.py` | 79 | Python | `i18n_messages` |
| `zeaz-live/backend/rbac.py` | 1 | Python | `check` |
| `zeaz-live/backend/rl/model.py` | 8 | Python | `__init__` |
| `zeaz-live/backend/rl/model.py` | 16 | Python | `forward` |
| `zeaz-live/backend/rl/train.py` | 8 | Python | `train` |
| `zeaz-live/backend/stream.py` | 4 | Python | `update_state` |
| `zeaz-live/backend/tenant.py` | 4 | Python | `get_tenant_state` |
| `zeaz-live/frontend/src/api.ts` | 1 | JS/TS | `getMetrics` |
| `zeaz-live/frontend/src/api.ts` | 6 | JS/TS | `getDecision` |
| `zeaz-live/frontend/src/i18n.ts` | 13 | JS/TS | `fetchMessages` |
| `zeaz-live/frontend/src/ws.ts` | 1 | JS/TS | `connectWS` |
| `zeaz-prod/backend/auth.py` | 12 | Python | `create_token` |
| `zeaz-prod/backend/auth.py` | 21 | Python | `verify` |
| `zeaz-prod/backend/auth.py` | 28 | Python | `verify_credentials` |
| `zeaz-prod/backend/db.py` | 15 | Python | `init_db` |
| `zeaz-prod/backend/db.py` | 35 | Python | `save_decision` |
| `zeaz-prod/backend/explain.py` | 1 | Python | `explain` |
| `zeaz-prod/backend/kafka.py` | 7 | Python | `_serialize` |
| `zeaz-prod/backend/kafka.py` | 17 | Python | `publish` |
| `zeaz-prod/backend/main.py` | 19 | Python | `startup` |
| `zeaz-prod/backend/main.py` | 24 | Python | `token` |
| `zeaz-prod/backend/main.py` | 28 | Python | `get_state` |
| `zeaz-prod/backend/main.py` | 35 | Python | `ws` |
| `zeaz-prod/backend/redis_client.py` | 13 | Python | `publish` |
| `zeaz-prod/backend/redis_client.py` | 17 | Python | `cache_set` |
| `zeaz-prod/backend/rl/agent.py` | 11 | Python | `__init__` |
| `zeaz-prod/backend/rl/agent.py` | 18 | Python | `key` |
| `zeaz-prod/backend/rl/agent.py` | 21 | Python | `_ensure` |
| `zeaz-prod/backend/rl/agent.py` | 25 | Python | `act` |
| `zeaz-prod/backend/rl/agent.py` | 34 | Python | `update` |
| `zeaz-prod/backend/rl/agent.py` | 41 | Python | `replay_train` |
| `zeaz-prod/backend/rl/replay.py` | 6 | Python | `__init__` |
| `zeaz-prod/backend/rl/replay.py` | 10 | Python | `add` |
| `zeaz-prod/backend/rl/replay.py` | 15 | Python | `sample` |
| `zeaz-quant/app/data/collector.py` | 7 | Python | `get_market` |
| `zeaz-quant/app/execution/engine.py` | 2 | Python | `execute` |
| `zeaz-quant/app/features/builder.py` | 5 | Python | `build` |
| `zeaz-quant/app/main.py` | 21 | Python | `loop` |
| `zeaz-quant/app/risk/manager.py` | 2 | Python | `__init__` |
| `zeaz-quant/app/risk/manager.py` | 7 | Python | `validate` |
| `zeaz-quant/app/risk/manager.py` | 16 | Python | `update` |
| `zeaz-quant/app/signal/engine.py` | 2 | Python | `generate` |
| `zeaz-quant/app/storage/db.py` | 5 | Python | `__init__` |
| `zeaz-quant/app/storage/db.py` | 11 | Python | `save` |
| `zeaz-v12/app/aiops/carbon.py` | 8 | Python | `best_region` |
| `zeaz-v12/app/aiops/carbon.py` | 11 | Python | `score` |
| `zeaz-v12/app/aiops/chaos.py` | 7 | Python | `inject` |
| `zeaz-v12/app/aiops/cost_optimizer.py` | 4 | Python | `__init__` |
| `zeaz-v12/app/aiops/cost_optimizer.py` | 10 | Python | `estimate` |
| `zeaz-v12/app/aiops/cost_optimizer.py` | 13 | Python | `recommend` |
| `zeaz-v12/app/aiops/cost_optimizer.py` | 18 | Python | `enforce_budget` |
| `zeaz-v12/app/aiops/economic_router.py` | 7 | Python | `__init__` |
| `zeaz-v12/app/aiops/economic_router.py` | 12 | Python | `route` |
| `zeaz-v12/app/aiops/global_router.py` | 4 | Python | `split` |
| `zeaz-v12/app/aiops/global_router.py` | 15 | Python | `failover` |
| `zeaz-v12/app/aiops/orchestrator.py` | 2 | Python | `step` |
| `zeaz-v12/app/aiops/pricing_rt.py` | 2 | Python | `__init__` |
| `zeaz-v12/app/aiops/pricing_rt.py` | 5 | Python | `price` |
| `zeaz-v12/app/aiops/pricing_rt.py` | 8 | Python | `surge` |
| `zeaz-v12/app/aiops/sla_negotiation.py` | 2 | Python | `negotiate` |
| `zeaz-v12/app/aiops/spot_scheduler.py` | 4 | Python | `choose_pool` |
| `zeaz-v12/app/aiops/spot_scheduler.py` | 9 | Python | `can_evict` |
| `zeaz-v12/app/infra_client/k8s.py` | 11 | Python | `apply_action` |
| `zeaz-v12/app/sim/simulator.py` | 5 | Python | `__init__` |
| `zeaz-v12/app/sim/simulator.py` | 12 | Python | `step` |
| `zeaz/app/aiops/dataset.py` | 5 | Python | `__init__` |
| `zeaz/app/aiops/dataset.py` | 9 | Python | `add` |
| `zeaz/app/aiops/dataset.py` | 14 | Python | `sample` |
| `zeaz/app/aiops/evolution.py` | 5 | Python | `__init__` |
| `zeaz/app/aiops/evolution.py` | 8 | Python | `propose` |
| `zeaz/app/aiops/evolution.py` | 18 | Python | `promote` |
| `zeaz/app/aiops/evolution.py` | 21 | Python | `rollback` |
| `zeaz/app/aiops/meta.py` | 4 | Python | `maml_step` |
| `zeaz/app/aiops/planner_mpc.py` | 9 | Python | `onehot` |
| `zeaz/app/aiops/planner_mpc.py` | 16 | Python | `__init__` |
| `zeaz/app/aiops/planner_mpc.py` | 22 | Python | `plan` |
| `zeaz/app/aiops/policy_actor.py` | 6 | Python | `__init__` |
| `zeaz/app/aiops/policy_actor.py` | 14 | Python | `forward` |
| `zeaz/app/aiops/world_model.py` | 6 | Python | `__init__` |
| `zeaz/app/aiops/world_model.py` | 14 | Python | `forward` |
| `zeaz/app/aiops/world_model.py` | 19 | Python | `__init__` |
| `zeaz/app/aiops/world_model.py` | 27 | Python | `forward` |
| `zeaz/app/aiops/world_model.py` | 33 | Python | `__init__` |
| `zeaz/app/aiops/world_model.py` | 41 | Python | `forward` |
| `zeaz/app/aiops/world_model.py` | 46 | Python | `__init__` |
| `zeaz/app/aiops/world_model.py` | 52 | Python | `forward` |
| `zeaz/app/control_main.py` | 7 | Python | `main` |
| `zeaz/app/execution/engine.py` | 5 | Python | `execute` |
| `zeaz/app/ingestion/market.py` | 11 | Python | `run` |
| `zeaz/app/ingestion/sentiment.py` | 4 | Python | `run` |
| `zeaz/app/main.py` | 8 | Python | `main` |
| `zeaz/app/observability/metrics.py` | 12 | Python | `start_metrics_server` |
| `zeaz/app/risk/engine.py` | 5 | Python | `__init__` |
| `zeaz/app/risk/engine.py` | 10 | Python | `validate` |
| `zeaz/app/risk/engine.py` | 18 | Python | `update` |
| `zeaz/app/signal/engine.py` | 5 | Python | `generate` |
| `zeaz/app/storage/clickhouse.py` | 7 | Python | `__init__` |
| `zeaz/app/storage/clickhouse.py` | 10 | Python | `insert` |
| `zeaz/app/stream/processor.py` | 14 | Python | `run` |
| `zk/commit_trace.py` | 6 | Python | `trace_commit` |
| `zk/prover/main.go` | 10 | Go | `VerifyTransition` |
| `zk/prover/main.go` | 16 | Go | `main` |
| `zk/prover_stub.py` | 1 | Python | `generate_proof` |
| `zk/prover_stub.py` | 6 | Python | `verify_proof` |

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

## `ZeaZDev-Omega`

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

## `ztsaff`

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

## `zLinebot`

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
| `app/src/services/referral.ts` | 5 | JS/TS | `registerReferral` |
| `app/src/services/risk.infer.ts` | 1 | JS/TS | `finalRisk` |
| `app/src/services/risk.ts` | 7 | JS/TS | `ruleScore` |
| `app/src/services/rl.ts` | 3 | JS/TS | `rankRL` |
| `app/src/services/rlhf.ts` | 13 | JS/TS | `normalizeReward` |
| `app/src/services/rlhf.ts` | 21 | JS/TS | `saveFeedback` |
| `app/src/services/rlhf.ts` | 59 | JS/TS | `logFeedback` |
| `app/src/services/routing.ts` | 8 | JS/TS | `chooseWarehouse` |
| `app/src/services/runpod.client.ts` | 6 | JS/TS | `assertEnv` |
| `app/src/services/runpod.client.ts` | 11 | JS/TS | `sign` |
| `app/src/services/runpod.client.ts` | 15 | JS/TS | `runpodInfer` |
| `app/src/services/scheduler.ts` | 3 | JS/TS | `chooseCompute` |
| `app/src/services/session.ts` | 7 | JS/TS | `keyFor` |
| `app/src/services/session.ts` | 9 | JS/TS | `getSessionVec` |
| `app/src/services/session.ts` | 22 | JS/TS | `setSessionVec` |
| `app/src/services/sessionUpdate.ts` | 4 | JS/TS | `normalizeAlpha` |
| `app/src/services/sessionUpdate.ts` | 8 | JS/TS | `ema` |
| `app/src/services/sessionUpdate.ts` | 19 | JS/TS | `updateSession` |
| `app/src/services/stripe.ts` | 6 | JS/TS | `getStripeClient` |
| `app/src/services/stripe.ts` | 16 | JS/TS | `createCheckout` |
| `app/src/services/tenantProvision.ts` | 3 | JS/TS | `safeTenantSchema` |
| `app/src/services/tenantProvision.ts` | 7 | JS/TS | `createTenantSchema` |
| `app/src/services/tiktok.shop.ts` | 53 | JS/TS | `readString` |
| `app/src/services/tiktok.shop.ts` | 57 | JS/TS | `readNumber` |
| `app/src/services/tiktok.shop.ts` | 65 | JS/TS | `loadLocalProducts` |
| `app/src/services/tiktok.shop.ts` | 91 | JS/TS | `loadTikTokShowcaseFromApi` |
| `app/src/services/tiktok.shop.ts` | 120 | JS/TS | `loadTikTokUserProfilesFromApi` |
| `app/src/services/tiktok.shop.ts` | 151 | JS/TS | `loadLocalUserProfiles` |
| `app/src/services/tiktok.shop.ts` | 158 | JS/TS | `COUNT` |
| `app/src/services/tiktok.shop.ts` | 180 | JS/TS | `fetchShowcaseProducts` |
| `app/src/services/tiktok.shop.ts` | 194 | JS/TS | `fetchShopUserProfiles` |
| `app/src/services/tiktok.shop.ts` | 208 | JS/TS | `buildShopIntelligenceReport` |
| `app/src/services/tiktok.shop.ts` | 239 | JS/TS | `escapeCsv` |
| `app/src/services/tiktok.shop.ts` | 244 | JS/TS | `exportShopIntelligenceCsv` |
| `app/src/services/tiktok.shop.ts` | 283 | JS/TS | `parseHashtags` |
| `app/src/services/tiktok.shop.ts` | 288 | JS/TS | `parseShotPlan` |
| `app/src/services/tiktok.shop.ts` | 297 | JS/TS | `generateVideoDrafts` |
| `app/src/services/tiktok.shop.ts` | 345 | JS/TS | `listVideoJobs` |
| `app/src/services/tiktok.stream.ts` | 24 | JS/TS | `normalizeString` |
| `app/src/services/tiktok.stream.ts` | 33 | JS/TS | `enqueueTikTokWebhookEvent` |
| `app/src/services/tiktok.stream.ts` | 75 | JS/TS | `ensureConsumerGroup` |
| `app/src/services/tiktok.stream.ts` | 87 | JS/TS | `mapFieldsToRecord` |
| `app/src/services/tiktok.stream.ts` | 102 | JS/TS | `isStreamReadResult` |
| `app/src/services/tiktok.stream.ts` | 128 | JS/TS | `persistTikTokEvent` |
| `app/src/services/tiktok.stream.ts` | 147 | JS/TS | `moveToDlq` |
| `app/src/services/tiktok.stream.ts` | 168 | JS/TS | `processStreamEntry` |
| `app/src/services/tiktok.stream.ts` | 186 | JS/TS | `startTikTokStreamWorker` |
| `app/src/services/tiktok.ts` | 35 | JS/TS | `asJsonRecord` |
| `app/src/services/tiktok.ts` | 43 | JS/TS | `normalizeScope` |
| `app/src/services/tiktok.ts` | 55 | JS/TS | `createState` |
| `app/src/services/tiktok.ts` | 59 | JS/TS | `buildAuthorizationUrl` |
| `app/src/services/tiktok.ts` | 76 | JS/TS | `exchangeCodeForToken` |
| `app/src/services/tiktok.ts` | 118 | JS/TS | `fetchUserInfo` |
| `app/src/services/tiktok.ts` | 135 | JS/TS | `verifyWebhookSignature` |
| `app/src/services/traffic.ts` | 5 | JS/TS | `variant` |
| `app/src/services/traffic.ts` | 11 | JS/TS | `setTrafficSplit` |
| `app/src/services/twotower.ts` | 5 | JS/TS | `loadTwoTower` |
| `app/src/services/twotower.ts` | 15 | JS/TS | `scoreTwoTower` |
| `app/src/services/userEmbedding.ts` | 4 | JS/TS | `updateUserEmbedding` |
| `app/src/services/vector.search.ts` | 1 | JS/TS | `search` |
| `app/src/services/vector.ts` | 3 | JS/TS | `resolveCloudEndpoint` |
| `app/src/services/vector.ts` | 10 | JS/TS | `endpoint` |
| `app/src/services/vector.ts` | 14 | JS/TS | `upsert` |
| `app/src/services/warehouse.consumer.ts` | 5 | JS/TS | `startWarehouseConsumer` |
| `app/src/types/onnxruntime-node.d.ts` | 13 | JS/TS | `create` |
| `app/src/types/onnxruntime-node.d.ts` | 14 | JS/TS | `run` |
| `app/src/utils/env.ts` | 1 | JS/TS | `parsePort` |
| `app/src/utils/env.ts` | 15 | JS/TS | `parsePositiveInt` |
| `app/src/utils/env.ts` | 28 | JS/TS | `parseBoolean` |
| `app/src/utils/env.ts` | 46 | JS/TS | `parseWorkerMode` |
| `app/src/utils/env.ts` | 54 | JS/TS | `requireNonEmpty` |
| `app/src/utils/env.ts` | 63 | JS/TS | `ensureSecretStrength` |
| `app/src/worker.ts` | 18 | JS/TS | `shutdown` |
| `app/src/world/basic.ts` | 3 | JS/TS | `step` |
| `app/src/world/mixture.ts` | 1 | JS/TS | `mixture` |
| `app/src/world/self_improve.basic.ts` | 15 | JS/TS | `predict` |
| `app/src/world/self_improve.basic.ts` | 20 | JS/TS | `update` |
| `app/src/world/self_improve.basic.ts` | 28 | JS/TS | `guard` |
| `app/src/world/self_improve.ts` | 10 | JS/TS | `setTrafficSplit` |
| `app/src/ws.ts` | 7 | JS/TS | `startWS` |
| `apps/admin/src/pages/Automations.tsx` | 7 | JS/TS | `addAction` |
| `apps/admin/src/pages/Automations.tsx` | 11 | JS/TS | `saveAutomation` |
| `apps/admin/src/pages/Builder.tsx` | 23 | JS/TS | `onConnect` |
| `apps/admin/src/pages/Builder.tsx` | 27 | JS/TS | `saveFlow` |
| `apps/api/src/middleware/auth.ts` | 3 | JS/TS | `authMiddleware` |
| `apps/api/src/plugins/rateLimit.ts` | 3 | JS/TS | `rateLimitPlugin` |
| `apps/api/src/routes/analytics.ts` | 3 | JS/TS | `analyticsRoutes` |
| `apps/api/src/routes/auth.ts` | 5 | JS/TS | `authRoutes` |
| `apps/api/src/routes/automation.ts` | 4 | JS/TS | `automationRoutes` |
| `apps/api/src/routes/health.ts` | 3 | JS/TS | `healthRoutes` |
| `apps/api/src/routes/logs.ts` | 3 | JS/TS | `logsRoutes` |
| `apps/api/src/routes/stripeWebhook.ts` | 10 | JS/TS | `stripeWebhook` |
| `apps/api/src/routes/webhook.ts` | 74 | JS/TS | `safeEqual` |
| `apps/api/src/routes/webhook.ts` | 84 | JS/TS | `verifyTikTokSignature` |
| `apps/api/src/routes/webhook.ts` | 106 | JS/TS | `verifyLineSignature` |
| `apps/api/src/routes/webhook.ts` | 119 | JS/TS | `parseTikTokSignature` |
| `apps/api/src/routes/webhook.ts` | 147 | JS/TS | `isFreshTimestamp` |
| `apps/api/src/routes/webhook.ts` | 158 | JS/TS | `parseTenantId` |
| `apps/api/src/routes/webhook.ts` | 167 | JS/TS | `validatePayload` |
| `apps/api/src/routes/webhook.ts` | 183 | JS/TS | `validateLinePayload` |
| `apps/api/src/routes/webhook.ts` | 200 | JS/TS | `headerValue` |
| `apps/api/src/routes/webhook.ts` | 210 | JS/TS | `getRawBody` |
| `apps/api/src/routes/webhook.ts` | 214 | JS/TS | `rememberConversation` |
| `apps/api/src/routes/webhook.ts` | 231 | JS/TS | `parseAutoReplyOverrides` |
| `apps/api/src/routes/webhook.ts` | 244 | JS/TS | `sanitizeText` |
| `apps/api/src/routes/webhook.ts` | 248 | JS/TS | `autoReplyIntent` |
| `apps/api/src/routes/webhook.ts` | 261 | JS/TS | `tenantAutoReplyProfile` |
| `apps/api/src/routes/webhook.ts` | 281 | JS/TS | `buildTenantAutoReply` |
| `apps/api/src/routes/webhook.ts` | 333 | JS/TS | `webhookRoutes` |
| `apps/api/src/tracing.ts` | 3 | JS/TS | `initTracer` |
| `apps/worker/src/executor.ts` | 7 | JS/TS | `executeAutomation` |
| `apps/worker/src/executor.ts` | 35 | JS/TS | `resolveValue` |
| `apps/worker/src/executor.ts` | 42 | JS/TS | `evalCondition` |
| `apps/worker/src/executor.ts` | 56 | JS/TS | `runAction` |
| `apps/worker/src/tracing.ts` | 3 | JS/TS | `initWorkerTracer` |
| `apps/worker/src/worker.ts` | 9 | JS/TS | `runKafkaWorker` |
| `bootstrap.sh` | 9 | Shell | `log` |
| `bootstrap.sh` | 13 | Shell | `require_root` |
| `bootstrap.sh` | 20 | Shell | `wait_for_deploy` |
| `cloud/worker/main.py` | 12 | Python | `infer` |
| `cloudflare/router.js` | 2 | JS/TS | `fetch` |
| `cloudflare/worker.js` | 1 | JS/TS | `computeLineSignature` |
| `cloudflare/worker.js` | 13 | JS/TS | `verifyLineSignature` |
| `cloudflare/worker.js` | 24 | JS/TS | `forwardLineWebhook` |
| `cloudflare/worker.js` | 42 | JS/TS | `fetch` |
| `ml/anomaly.py` | 4 | Python | `fit_anomaly_model` |
| `ml/diffusion_rec.py` | 10 | Python | `__init__` |
| `ml/diffusion_rec.py` | 25 | Python | `forward` |
| `ml/diffusion_rec.py` | 30 | Python | `sample` |
| `ml/export_onnx.py` | 11 | Python | `_verify_model_file` |
| `ml/export_onnx.py` | 24 | Python | `export_onnx` |
| `ml/federated_client.py` | 1 | Python | `local_train` |
| `ml/federated_server.py` | 4 | Python | `aggregate` |
| `ml/federated_server.py` | 11 | Python | `apply` |
| `ml/fhe_secure.py` | 8 | Python | `secure_add` |
| `ml/foundation_ranker.py` | 14 | Python | `__init__` |
| `ml/foundation_ranker.py` | 28 | Python | `forward` |
| `ml/he.py` | 8 | Python | `encrypt` |
| `ml/he.py` | 12 | Python | `add` |
| `ml/he.py` | 17 | Python | `decrypt` |
| `ml/market_env.py` | 17 | Python | `__init__` |
| `ml/market_env.py` | 43 | Python | `_sample_seller_state` |
| `ml/market_env.py` | 51 | Python | `_obs` |
| `ml/market_env.py` | 60 | Python | `reset` |
| `ml/market_env.py` | 73 | Python | `step` |
| `ml/ppo.py` | 4 | Python | `ppo_step` |
| `ml/reward_model.py` | 6 | Python | `__init__` |
| `ml/reward_model.py` | 14 | Python | `forward` |
| `ml/reward_model.py` | 18 | Python | `train_reward_model` |
| `ml/shap_explain.py` | 4 | Python | `explain` |
| `ml/simulate.py` | 5 | Python | `simulate` |
| `ml/train.py` | 6 | Python | `__init__` |
| `ml/train.py` | 14 | Python | `forward` |
| `ml/train.py` | 19 | Python | `train` |
| `ml/train_multi_agent.py` | 21 | Python | `__init__` |
| `ml/train_multi_agent.py` | 25 | Python | `score` |
| `ml/train_multi_agent.py` | 32 | Python | `act` |
| `ml/train_multi_agent.py` | 46 | Python | `generate_logs` |
| `ml/transformer_ranker.py` | 14 | Python | `__init__` |
| `ml/transformer_ranker.py` | 23 | Python | `forward` |
| `ml/two_tower.py` | 6 | Python | `__init__` |
| `ml/two_tower.py` | 14 | Python | `forward` |
| `ml/two_tower.py` | 19 | Python | `__init__` |
| `ml/two_tower.py` | 24 | Python | `forward` |
| `ml/two_tower.py` | 30 | Python | `train` |
| `ml/world_model.py` | 6 | Python | `__init__` |
| `ml/world_model.py` | 14 | Python | `forward` |
| `ml/world_model.py` | 22 | Python | `__init__` |
| `ml/world_model.py` | 27 | Python | `forward` |
| `mobile/Admin.tsx` | 123 | JS/TS | `applyGuardrailPreset` |
| `mobile/Admin.tsx` | 160 | JS/TS | `acknowledgeIncident` |
| `mobile/Admin.tsx` | 169 | JS/TS | `toggleExperimentLocal` |
| `mobile/Admin.tsx` | 187 | JS/TS | `runRetrain` |
| `mobile/Admin.tsx` | 427 | JS/TS | `TabButton` |
| `mobile/Admin.tsx` | 449 | JS/TS | `PanelToggle` |
| `mobile/Admin.tsx` | 465 | JS/TS | `Stepper` |
| `mobile/App.tsx` | 149 | JS/TS | `addToCart` |
| `mobile/App.tsx` | 162 | JS/TS | `checkout` |
| `mobile/App.tsx` | 186 | JS/TS | `advanceOrder` |
| `mobile/App.tsx` | 201 | JS/TS | `setPreference` |
| `mobile/App.tsx` | 205 | JS/TS | `saveProfile` |
| `mobile/App.tsx` | 209 | JS/TS | `runAssistant` |
| `mobile/App.tsx` | 427 | JS/TS | `TabButton` |
| `mobile/App.tsx` | 449 | JS/TS | `PanelToggle` |
| `mobile/ai.ts` | 1 | JS/TS | `localAI` |
| `mobile/api.ts` | 1 | JS/TS | `triggerRetrain` |
| `mobile/api.ts` | 5 | JS/TS | `toggleExperiment` |
| `packages/auth/apiKey.ts` | 4 | JS/TS | `createApiKey` |
| `packages/auth/encryption.ts` | 5 | JS/TS | `encrypt` |
| `packages/auth/jwt.ts` | 5 | JS/TS | `signToken` |
| `packages/auth/jwt.ts` | 9 | JS/TS | `verifyToken` |
| `packages/automation/ai-builder.ts` | 29 | JS/TS | `generateFlow` |
| `packages/automation/ai.ts` | 7 | JS/TS | `generateReply` |
| `packages/automation/compiler.ts` | 1 | JS/TS | `compileFlow` |
| `packages/automation/kafka.ts` | 18 | JS/TS | `emitEvent` |
| `packages/automation/logger.ts` | 19 | JS/TS | `log` |
| `packages/automation/plugins.ts` | 5 | JS/TS | `registerPlugin` |
| `packages/automation/plugins.ts` | 9 | JS/TS | `runPlugin` |
| `packages/automation/runner.ts` | 12 | JS/TS | `processEvent` |
| `packages/automation/stream.ts` | 21 | JS/TS | `startStream` |
| `packages/automation/stream.ts` | 46 | JS/TS | `stopStream` |
| `packages/billing/cost.ts` | 3 | JS/TS | `trackCost` |
| `packages/billing/subscription.ts` | 4 | JS/TS | `createCustomer` |
| `packages/billing/subscription.ts` | 21 | JS/TS | `createSubscription` |
| `packages/billing/usage.ts` | 3 | JS/TS | `trackUsage` |
| `packages/config/generate-env.ts` | 5 | JS/TS | `generateSecret` |
| `packages/monitoring/anomaly.ts` | 10 | JS/TS | `parsePrometheusValue` |
| `packages/monitoring/anomaly.ts` | 18 | JS/TS | `detectAnomaly` |
| `scripts/bootstrap-saas-foundation.sh` | 11 | Shell | `write_if_missing` |
| `scripts/configure_cloudflare.sh` | 7 | Shell | `require_env` |
| `scripts/configure_cloudflare.sh` | 83 | Shell | `upsert_record` |
| `scripts/configure_cloudflare_zero_trust_full.sh` | 20 | Shell | `require_value` |
| `scripts/configure_cloudflare_zero_trust_full.sh` | 40 | Shell | `cf_call` |
| `scripts/configure_cloudflare_zero_trust_full.sh` | 71 | Shell | `configure_tunnel_ingress` |
| `scripts/configure_cloudflare_zero_trust_full.sh` | 93 | Shell | `find_dns_record_id` |
| `scripts/configure_cloudflare_zero_trust_full.sh` | 116 | Shell | `upsert_dns_cname` |
| `scripts/deploy-k8s.sh` | 8 | Shell | `log` |
| `scripts/deploy-k8s.sh` | 10 | Shell | `check_command` |
| `scripts/deploy-k8s.sh` | 19 | Shell | `main` |
| `scripts/deploy.sh` | 13 | Shell | `log` |
| `scripts/deploy.sh` | 15 | Shell | `check_docker` |
| `scripts/deploy.sh` | 41 | Shell | `main` |
| `scripts/generate-secrets.sh` | 35 | Shell | `rand_base64` |
| `scripts/generate-secrets.sh` | 39 | Shell | `rand_hex` |
| `scripts/generate-secrets.sh` | 43 | Shell | `rand_hex_64` |
| `scripts/generate-secrets.sh` | 50 | Shell | `prompt_value` |
| `scripts/install_full.sh` | 6 | Shell | `log` |
| `scripts/install_full.sh` | 10 | Shell | `resolve_repo_root` |
| `scripts/install_full.sh` | 49 | Shell | `install_tenant_namespace` |
| `scripts/install_full.sh` | 60 | Shell | `run_npm_install_autoheal` |
| `scripts/install_full.sh` | 89 | Shell | `configure_cloudflare_zero_trust_autoheal` |
| `scripts/install_secure.sh` | 4 | Shell | `gen_secret` |
| `scripts/master_meta_deep_impact_dive_audit_scan.sh` | 20 | Shell | `count_files` |
| `scripts/master_meta_deep_impact_dive_audit_scan.sh` | 26 | Shell | `count_loc` |
| `scripts/master_meta_full_source_ci_cd.sh` | 12 | Shell | `run_ci` |
| `scripts/master_meta_full_source_ci_cd.sh` | 20 | Shell | `run_cd` |
| `scripts/spiffe.sh` | 9 | Shell | `cleanup` |
| `tools/codex_engine/analyzers.py` | 30 | Python | `run` |
| `tools/codex_engine/analyzers.py` | 67 | Python | `run` |
| `tools/codex_engine/analyzers.py` | 97 | Python | `run` |
| `tools/codex_engine/analyzers.py` | 130 | Python | `run` |
| `tools/codex_engine/analyzers.py` | 164 | Python | `run` |
| `tools/codex_engine/analyzers.py` | 200 | Python | `__init__` |
| `tools/codex_engine/analyzers.py` | 203 | Python | `run` |
| `tools/codex_engine/cli.py` | 12 | Python | `main` |
| `tools/codex_engine/models.py` | 47 | Python | `blocking_findings` |
| `tools/codex_engine/pipeline.py` | 21 | Python | `__init__` |
| `tools/codex_engine/pipeline.py` | 31 | Python | `_target_files` |
| `tools/codex_engine/pipeline.py` | 48 | Python | `_ollama_review` |
| `tools/codex_engine/pipeline.py` | 71 | Python | `run` |
| `watchdog.sh` | 8 | Shell | `log` |
| `zlinebot-master-orchestrator.sh` | 13 | Shell | `log` |
| `zlinebot-master-orchestrator.sh` | 15 | Shell | `check_docker` |
| `zlinebot-master-orchestrator.sh` | 29 | Shell | `setup_cloudflare_tunnel` |
| `zlinebot-master-orchestrator.sh` | 46 | Shell | `validate_secrets` |
| `zlinebot-master-orchestrator.sh` | 58 | Shell | `main` |
| `zlinebot-master-selfheal.sh` | 16 | Shell | `log` |
| `zlinebot-master-selfheal.sh` | 20 | Shell | `retry` |
| `zlinebot-master-selfheal.sh` | 33 | Shell | `wait_k8s` |
| `zlinebot-master-selfheal.sh` | 44 | Shell | `wait_pods` |
| `zlinebot-master-selfheal.sh` | 62 | Shell | `wait_rollouts` |
| `zlinebot-master-selfheal.sh` | 69 | Shell | `health_check` |
| `zlinebot-master-selfheal.sh` | 83 | Shell | `rollback` |
| `zlinebot-master-selfheal.sh` | 92 | Shell | `main` |
| `zlinebot-master.sh` | 31 | Shell | `install_docker_engine` |
| `zlinebot-master.sh` | 71 | Shell | `wait_for_k3s_api` |

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

## `zlttbots`

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
| `services/account-farm/src/core/proxy.js` | 3 | JS/TS | `getProxy` |
| `services/account-farm/src/monitor/health.js` | 3 | JS/TS | `farmHealth` |
| `services/account-farm/src/scheduler/scheduler.js` | 3 | JS/TS | `scheduler` |
| `services/account-farm/src/scheduler/worker.js` | 7 | JS/TS | `worker` |
| `services/account-farm/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/account-farm/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/account-farm/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/account-farm/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/affiliate-webhook/src/main.py` | 35 | Python | `verify` |
| `services/affiliate-webhook/src/main.py` | 40 | Python | `db_connection` |
| `services/affiliate-webhook/src/main.py` | 47 | Python | `healthz` |
| `services/affiliate-webhook/src/main.py` | 66 | Python | `conversion` |
| `services/affiliate-webhook/src/main.py` | 115 | Python | `get_server_config` |
| `services/affiliate-webhook/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/affiliate-webhook/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/affiliate-webhook/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/affiliate-webhook/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/ai-fix/src/engine.py` | 5 | Python | `generate_patch` |
| `services/ai-fix/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/ai-fix/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/ai-fix/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/ai-fix/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/ai-orchestrator/src/main.py` | 11 | Python | `health` |
| `services/ai-orchestrator/src/main.py` | 15 | Python | `run_growth_cycle` |
| `services/ai-orchestrator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/ai-orchestrator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/ai-orchestrator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/ai-orchestrator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/ai-video-generator/src/core/generator.js` | 5 | JS/TS | `generateVideo` |
| `services/ai-video-generator/src/core/worker.js` | 4 | JS/TS | `worker` |
| `services/ai-video-generator/src/index.js` | 3 | JS/TS | `main` |
| `services/ai-video-generator/src/render/video.js` | 3 | JS/TS | `renderVideo` |
| `services/ai-video-generator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/ai-video-generator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/ai-video-generator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/ai-video-generator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/ai-video-generator/src/subtitle/subtitle.js` | 7 | JS/TS | `safeSubtitlePath` |
| `services/ai-video-generator/src/subtitle/subtitle.js` | 21 | JS/TS | `generateSubtitle` |
| `services/ai-video-generator/src/tts/voice.js` | 10 | JS/TS | `validateFilename` |
| `services/ai-video-generator/src/tts/voice.js` | 22 | JS/TS | `resolveOutputPath` |
| `services/ai-video-generator/src/tts/voice.js` | 32 | JS/TS | `hasKnownAudioHeader` |
| `services/ai-video-generator/src/tts/voice.js` | 43 | JS/TS | `toSafeAudioBuffer` |
| `services/ai-video-generator/src/tts/voice.js` | 54 | JS/TS | `generateVoice` |
| `services/analytics/ltv_model.py` | 7 | Python | `predict` |
| `services/analytics/roi.py` | 9 | Python | `fetch_revenue` |
| `services/analytics/roi.py` | 11 | Python | `fetch_cost` |
| `services/analytics/roi.py` | 22 | Python | `profitable` |
| `services/analytics/roi.py` | 29 | Python | `__init__` |
| `services/analytics/roi.py` | 32 | Python | `evaluate` |
| `services/analytics/src/metrics/campaign.js` | 3 | JS/TS | `campaignROI` |
| `services/analytics/src/metrics/conversion.js` | 3 | JS/TS | `conversionRate` |
| `services/analytics/src/metrics/products.js` | 3 | JS/TS | `productPerformance` |
| `services/analytics/src/metrics/revenue.js` | 3 | JS/TS | `totalRevenue` |
| `services/analytics/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/analytics/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/analytics/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/analytics/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/analytics/unit_economics.py` | 16 | Python | `profit` |
| `services/analytics/unit_economics.py` | 20 | Python | `cpa` |
| `services/analytics/unit_economics.py` | 26 | Python | `ltv` |
| `services/analytics/unit_economics.py` | 32 | Python | `roas` |
| `services/analytics/unit_economics.py` | 38 | Python | `profitable` |
| `services/arbitrage-engine/src/api/server.py` | 91 | Python | `get_db` |
| `services/arbitrage-engine/src/api/server.py` | 95 | Python | `_probe_database` |
| `services/arbitrage-engine/src/api/server.py` | 106 | Python | `startup_probe` |
| `services/arbitrage-engine/src/api/server.py` | 118 | Python | `lifespan` |
| `services/arbitrage-engine/src/api/server.py` | 127 | Python | `health_live` |
| `services/arbitrage-engine/src/api/server.py` | 132 | Python | `health_ready` |
| `services/arbitrage-engine/src/api/server.py` | 141 | Python | `healthz` |
| `services/arbitrage-engine/src/api/server.py` | 156 | Python | `metrics` |
| `services/arbitrage-engine/src/api/server.py` | 161 | Python | `list_arbitrage_events` |
| `services/arbitrage-engine/src/api/server.py` | 170 | Python | `scan_arbitrage` |
| `services/arbitrage-engine/src/api/server.py` | 196 | Python | `ingest_payout` |
| `services/arbitrage-engine/src/api/server.py` | 216 | Python | `sync_affiliate` |
| `services/arbitrage-engine/src/api/server.py` | 238 | Python | `enqueue_publish_job` |
| `services/arbitrage-engine/src/api/server.py` | 246 | Python | `run_daily` |
| `services/arbitrage-engine/src/api/server.py` | 247 | Python | `publish_fn` |
| `services/arbitrage-engine/src/api/server.py` | 255 | Python | `get_counter` |
| `services/arbitrage-engine/src/api/server.py` | 261 | Python | `ingest_video` |
| `services/arbitrage-engine/src/api/server.py` | 267 | Python | `ingest_performance` |
| `services/arbitrage-engine/src/api/server.py` | 277 | Python | `posted_product_reporting` |
| `services/arbitrage-engine/src/bakup/commission.py` | 1 | Python | `commission_rate` |
| `services/arbitrage-engine/src/connectors/client.py` | 8 | Python | `__init__` |
| `services/arbitrage-engine/src/connectors/client.py` | 11 | Python | `fetch_network_snapshot` |
| `services/arbitrage-engine/src/connectors/contracts.py` | 29 | Python | `normalize_currency` |
| `services/arbitrage-engine/src/connectors/providers.py` | 21 | Python | `fetch_snapshot` |
| `services/arbitrage-engine/src/connectors/providers.py` | 26 | Python | `__init__` |
| `services/arbitrage-engine/src/connectors/providers.py` | 30 | Python | `_request` |
| `services/arbitrage-engine/src/connectors/providers.py` | 42 | Python | `fetch_snapshot` |
| `services/arbitrage-engine/src/connectors/providers.py` | 71 | Python | `default_connectors` |
| `services/arbitrage-engine/src/core/database.py` | 32 | Python | `__init__` |
| `services/arbitrage-engine/src/core/database.py` | 44 | Python | `seed_products` |
| `services/arbitrage-engine/src/core/database.py` | 53 | Python | `get_db` |
| `services/arbitrage-engine/src/core/database.py` | 57 | Python | `fetch_products` |
| `services/arbitrage-engine/src/core/database.py` | 82 | Python | `insert_event` |
| `services/arbitrage-engine/src/core/database.py` | 106 | Python | `list_events` |
| `services/arbitrage-engine/src/core/database.py` | 137 | Python | `upsert_product_payout` |
| `services/arbitrage-engine/src/core/database.py` | 157 | Python | `get_product_payout` |
| `services/arbitrage-engine/src/core/database.py` | 176 | Python | `enqueue_publish_job` |
| `services/arbitrage-engine/src/core/database.py` | 180 | Python | `dequeue_publish_jobs` |
| `services/arbitrage-engine/src/core/database.py` | 187 | Python | `increase_daily_counter` |
| `services/arbitrage-engine/src/core/database.py` | 193 | Python | `get_daily_counter` |
| `services/arbitrage-engine/src/core/database.py` | 197 | Python | `put_dead_letter` |
| `services/arbitrage-engine/src/core/database.py` | 201 | Python | `record_publish_result` |
| `services/arbitrage-engine/src/core/database.py` | 205 | Python | `record_video` |
| `services/arbitrage-engine/src/core/database.py` | 209 | Python | `record_performance` |
| `services/arbitrage-engine/src/core/database.py` | 213 | Python | `get_posted_product_reporting` |
| `services/arbitrage-engine/src/core/database.py` | 241 | Python | `utc_now` |
| `services/arbitrage-engine/src/engine/arbitrage.py` | 3 | Python | `detect` |
| `services/arbitrage-engine/src/main.py` | 4 | Python | `_safe_port` |
| `services/arbitrage-engine/src/main.py` | 19 | Python | `get_server_config` |
| `services/arbitrage-engine/src/publishing/controller.py` | 24 | Python | `__init__` |
| `services/arbitrage-engine/src/publishing/controller.py` | 28 | Python | `run_for_tenant` |
| `services/arbitrage-engine/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/arbitrage-engine/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/arbitrage-engine/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/arbitrage-engine/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/arbitrage-engine/src/workers/worker.py` | 7 | Python | `worker` |
| `services/auto-ml/loop.py` | 19 | Python | `_log` |
| `services/auto-ml/loop.py` | 24 | Python | `_run_command` |
| `services/auto-ml/loop.py` | 48 | Python | `train` |
| `services/auto-ml/loop.py` | 54 | Python | `evaluate` |
| `services/auto-ml/loop.py` | 66 | Python | `load_metric` |
| `services/auto-ml/loop.py` | 74 | Python | `save_metric` |
| `services/auto-ml/loop.py` | 80 | Python | `deploy` |
| `services/auto-ml/loop.py` | 86 | Python | `loop` |
| `services/billing-service/src/main.py` | 19 | Python | `healthz` |
| `services/billing-service/src/main.py` | 24 | Python | `charge` |
| `services/billing-service/src/main.py` | 36 | Python | `get_server_config` |
| `services/billing-service/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/billing-service/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/billing-service/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/billing-service/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/blockchain/bridge_secure.py` | 6 | Python | `verify_tx` |
| `services/budget-allocator/src/bandit.py` | 7 | Python | `__init__` |
| `services/budget-allocator/src/bandit.py` | 11 | Python | `select` |
| `services/budget-allocator/src/bandit.py` | 19 | Python | `update` |
| `services/budget-allocator/src/global_allocator.py` | 5 | Python | `__init__` |
| `services/budget-allocator/src/global_allocator.py` | 9 | Python | `allocate` |
| `services/budget-allocator/src/global_allocator.py` | 16 | Python | `update` |
| `services/budget-allocator/src/main.py` | 45 | Python | `healthz` |
| `services/budget-allocator/src/main.py` | 50 | Python | `allocate` |
| `services/budget-allocator/src/main.py` | 76 | Python | `get_server_config` |
| `services/budget-allocator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/budget-allocator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/budget-allocator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/budget-allocator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/campaign-optimizer/src/main.py` | 10 | Python | `_required_env` |
| `services/campaign-optimizer/src/main.py` | 17 | Python | `db` |
| `services/campaign-optimizer/src/main.py` | 40 | Python | `optimize` |
| `services/campaign-optimizer/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/campaign-optimizer/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/campaign-optimizer/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/campaign-optimizer/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/capital-allocator/src/capital_engine.py` | 7 | Python | `__init__` |
| `services/capital-allocator/src/capital_engine.py` | 10 | Python | `reinvest` |
| `services/capital-allocator/src/capital_engine.py` | 16 | Python | `allocate` |
| `services/capital-allocator/src/main.py` | 43 | Python | `healthz` |
| `services/capital-allocator/src/main.py` | 53 | Python | `allocate` |
| `services/capital-allocator/src/main.py` | 75 | Python | `get_server_config` |
| `services/capital-allocator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/capital-allocator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/capital-allocator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/capital-allocator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/click-tracker/src/core/tracker.js` | 5 | JS/TS | `normalizeRedirectTarget` |
| `services/click-tracker/src/core/tracker.js` | 16 | JS/TS | `trackClick` |
| `services/click-tracker/src/core/tracker.js` | 33 | JS/TS | `resolveProductRedirect` |
| `services/click-tracker/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/click-tracker/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/click-tracker/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/click-tracker/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/click-tracker/src/server/server.js` | 9 | JS/TS | `isAllowedRedirect` |
| `services/click-tracker/src/server/server.js` | 25 | JS/TS | `sendJson` |
| `services/click-tracker/src/utils/fingerprint.js` | 1 | JS/TS | `fingerprint` |
| `services/core/profit_loop.py` | 13 | Python | `create_campaign` |
| `services/core/profit_loop.py` | 15 | Python | `pause_campaign` |
| `services/core/profit_loop.py` | 17 | Python | `scale_campaign` |
| `services/core/profit_loop.py` | 28 | Python | `snapshot` |
| `services/core/profit_loop.py` | 32 | Python | `__init__` |
| `services/core/profit_loop.py` | 44 | Python | `launch` |
| `services/core/profit_loop.py` | 49 | Python | `run_once` |
| `services/cost/control.py` | 10 | Python | `allow` |
| `services/cost/control.py` | 19 | Python | `reset` |
| `services/cost/controls.py` | 14 | Python | `enable_kill_switch` |
| `services/cost/controls.py` | 17 | Python | `disable_kill_switch` |
| `services/cost/controls.py` | 20 | Python | `allow_spend` |
| `services/cost/controls.py` | 32 | Python | `should_stop_for_roi` |
| `services/cost/controls.py` | 35 | Python | `is_cost_spike` |
| `services/cost/guard.py` | 12 | Python | `allow` |
| `services/cost/guard.py` | 20 | Python | `reset` |
| `services/cost/guard.py` | 24 | Python | `from_env` |
| `services/creative-generator/src/ab_test.py` | 7 | Python | `__init__` |
| `services/creative-generator/src/ab_test.py` | 10 | Python | `update` |
| `services/creative-generator/src/ab_test.py` | 17 | Python | `select` |
| `services/creative-generator/src/llm.py` | 10 | Python | `generate_ad` |
| `services/creative-generator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/creative-generator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/creative-generator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/creative-generator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/data/pipeline.py` | 23 | Python | `validate` |
| `services/data/pipeline.py` | 28 | Python | `label` |
| `services/data/pipeline.py` | 34 | Python | `run` |
| `services/drift-detector/src/main.py` | 12 | Python | `loop` |
| `services/drift-detector/src/psi.py` | 6 | Python | `psi` |
| `services/drift-detector/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/drift-detector/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/drift-detector/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/drift-detector/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/edge-worker/index.js` | 2 | JS/TS | `fetch` |
| `services/exchange/src/main.py` | 40 | Python | `place` |
| `services/exchange/src/orderbook.py` | 16 | Python | `__post_init__` |
| `services/exchange/src/orderbook.py` | 25 | Python | `__init__` |
| `services/exchange/src/orderbook.py` | 29 | Python | `add` |
| `services/exchange/src/orderbook.py` | 35 | Python | `match` |
| `services/exchange/src/security.py` | 16 | Python | `verify_signature` |
| `services/exchange/src/security.py` | 24 | Python | `check_replay` |
| `services/exchange/src/security.py` | 31 | Python | `rate_limit` |
| `services/exchange/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/exchange/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/exchange/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/exchange/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/execution-engine/src/main.py` | 39 | Python | `http_post` |
| `services/execution-engine/src/main.py` | 59 | Python | `acquire_token` |
| `services/execution-engine/src/main.py` | 76 | Python | `_sanitize_log_value` |
| `services/execution-engine/src/main.py` | 81 | Python | `healthz` |
| `services/execution-engine/src/main.py` | 95 | Python | `publish` |
| `services/execution-engine/src/main.py` | 116 | Python | `status` |
| `services/execution-engine/src/main.py` | 120 | Python | `get_server_config` |
| `services/execution-engine/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/execution-engine/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/execution-engine/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/execution-engine/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/feature-store/src/main.py` | 14 | Python | `__init__` |
| `services/feature-store/src/main.py` | 17 | Python | `hgetall` |
| `services/feature-store/src/main.py` | 20 | Python | `hset` |
| `services/feature-store/src/main.py` | 23 | Python | `ping` |
| `services/feature-store/src/main.py` | 53 | Python | `validate_daily_cap` |
| `services/feature-store/src/main.py` | 72 | Python | `_load_features` |
| `services/feature-store/src/main.py` | 90 | Python | `_save_features` |
| `services/feature-store/src/main.py` | 110 | Python | `healthz` |
| `services/feature-store/src/main.py` | 120 | Python | `get_features` |
| `services/feature-store/src/main.py` | 127 | Python | `update_features` |
| `services/feature-store/src/main.py` | 163 | Python | `replace_all_features` |
| `services/feature-store/src/main.py` | 172 | Python | `replace_all_campaign_features` |
| `services/feature-store/src/main.py` | 183 | Python | `get_server_config` |
| `services/feature-store/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/feature-store/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/feature-store/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/feature-store/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/federation/src/main.py` | 22 | Python | `_safe_int_env` |
| `services/federation/src/main.py` | 43 | Python | `_safe_float_env` |
| `services/federation/src/main.py` | 87 | Python | `encode_signed_claims` |
| `services/federation/src/main.py` | 93 | Python | `db_connection` |
| `services/federation/src/main.py` | 97 | Python | `ensure_schema` |
| `services/federation/src/main.py` | 119 | Python | `startup` |
| `services/federation/src/main.py` | 153 | Python | `healthz` |
| `services/federation/src/main.py` | 168 | Python | `list_nodes` |
| `services/federation/src/main.py` | 198 | Python | `register` |
| `services/federation/src/main.py` | 276 | Python | `get_server_config` |
| `services/federation/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/federation/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/federation/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/federation/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/gpu-renderer/src/api/server.py` | 21 | Python | `healthz` |
| `services/gpu-renderer/src/api/server.py` | 40 | Python | `metrics` |
| `services/gpu-renderer/src/api/server.py` | 45 | Python | `render_video` |
| `services/gpu-renderer/src/core/queue.py` | 7 | Python | `enqueue` |
| `services/gpu-renderer/src/core/queue.py` | 10 | Python | `dequeue` |
| `services/gpu-renderer/src/core/render.py` | 6 | Python | `_should_use_cuda` |
| `services/gpu-renderer/src/core/render.py` | 20 | Python | `build_ffmpeg_command` |
| `services/gpu-renderer/src/core/render.py` | 55 | Python | `render` |
| `services/gpu-renderer/src/main.py` | 4 | Python | `get_server_config` |
| `services/gpu-renderer/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/gpu-renderer/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/gpu-renderer/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/gpu-renderer/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/gpu-renderer/src/worker/worker.py` | 5 | Python | `worker` |
| `services/gpu/adapter.py` | 12 | Python | `create_pod` |
| `services/identity/did.py` | 16 | Python | `_b64` |
| `services/identity/did.py` | 20 | Python | `_b64decode` |
| `services/identity/did.py` | 25 | Python | `load_or_create_key` |
| `services/identity/did.py` | 41 | Python | `get_did` |
| `services/identity/did.py` | 49 | Python | `export_public_key` |
| `services/identity/did.py` | 57 | Python | `sign` |
| `services/identity/did.py` | 63 | Python | `verify` |
| `services/identity/keystore.py` | 15 | Python | `__init__` |
| `services/identity/keystore.py` | 20 | Python | `_load_or_create` |
| `services/identity/keystore.py` | 43 | Python | `get` |
| `services/jwt-auth/src/main.py` | 27 | Python | `_b64encode` |
| `services/jwt-auth/src/main.py` | 31 | Python | `_b64decode` |
| `services/jwt-auth/src/main.py` | 36 | Python | `encode` |
| `services/jwt-auth/src/main.py` | 47 | Python | `decode` |
| `services/jwt-auth/src/main.py` | 106 | Python | `_parse_positive_int_env` |
| `services/jwt-auth/src/main.py` | 130 | Python | `healthz` |
| `services/jwt-auth/src/main.py` | 135 | Python | `jwks` |
| `services/jwt-auth/src/main.py` | 149 | Python | `issue_token` |
| `services/jwt-auth/src/main.py` | 176 | Python | `verify_bearer_token` |
| `services/jwt-auth/src/main.py` | 201 | Python | `introspect` |
| `services/jwt-auth/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/jwt-auth/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/jwt-auth/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/jwt-auth/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/landing-service/src/main.py` | 10 | Python | `healthz` |
| `services/landing-service/src/main.py` | 15 | Python | `landing` |
| `services/landing-service/src/main.py` | 23 | Python | `get_server_config` |
| `services/landing-service/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/landing-service/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/landing-service/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/landing-service/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/market-crawler/src/api/server.py` | 20 | Python | `healthz` |
| `services/market-crawler/src/api/server.py` | 39 | Python | `metrics` |
| `services/market-crawler/src/api/server.py` | 44 | Python | `crawl` |
| `services/market-crawler/src/core/database.py` | 6 | Python | `get_db` |
| `services/market-crawler/src/core/database.py` | 10 | Python | `insert_product` |
| `services/market-crawler/src/core/queue.py` | 9 | Python | `enqueue` |
| `services/market-crawler/src/core/queue.py` | 12 | Python | `dequeue` |
| `services/market-crawler/src/main.py` | 4 | Python | `get_server_config` |
| `services/market-crawler/src/parsers/amazon.py` | 4 | Python | `crawl_amazon` |
| `services/market-crawler/src/parsers/lazada.py` | 4 | Python | `crawl_lazada` |
| `services/market-crawler/src/parsers/shopee.py` | 3 | Python | `crawl_shopee` |
| `services/market-crawler/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/market-crawler/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/market-crawler/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/market-crawler/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/market-crawler/src/workers/worker.py` | 10 | Python | `process` |
| `services/market-crawler/src/workers/worker.py` | 23 | Python | `worker` |
| `services/market-creator/engine.py` | 13 | Python | `discover` |
| `services/market-creator/engine.py` | 21 | Python | `validate` |
| `services/market-creator/engine.py` | 26 | Python | `launch` |
| `services/market-creator/engine.py` | 36 | Python | `loop` |
| `services/market-orchestrator/src/main.py` | 20 | Python | `healthz` |
| `services/market-orchestrator/src/main.py` | 25 | Python | `launch` |
| `services/market-orchestrator/src/main.py` | 47 | Python | `get_server_config` |
| `services/market-orchestrator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/market-orchestrator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/market-orchestrator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/market-orchestrator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/master-orchestrator/src/deployment_controller.py` | 18 | Python | `_sanitize_log_value` |
| `services/master-orchestrator/src/deployment_controller.py` | 84 | Python | `_now_iso` |
| `services/master-orchestrator/src/deployment_controller.py` | 87 | Python | `create` |
| `services/master-orchestrator/src/deployment_controller.py` | 121 | Python | `apply_event` |
| `services/master-orchestrator/src/deployment_controller.py` | 149 | Python | `get` |
| `services/master-orchestrator/src/deployment_controller.py` | 155 | Python | `_log` |
| `services/master-orchestrator/src/distributed_loop.py` | 14 | Python | `_assert_internal_url` |
| `services/master-orchestrator/src/distributed_loop.py` | 25 | Python | `safe_call` |
| `services/master-orchestrator/src/distributed_loop.py` | 41 | Python | `run_cycle` |
| `services/master-orchestrator/src/economy_loop.py` | 13 | Python | `safe_call` |
| `services/master-orchestrator/src/economy_loop.py` | 25 | Python | `run_economy` |
| `services/master-orchestrator/src/execution.py` | 4 | Python | `execute_campaign` |
| `services/master-orchestrator/src/federated_loop.py` | 23 | Python | `_assert_internal_url` |
| `services/master-orchestrator/src/federated_loop.py` | 34 | Python | `safe_call` |
| `services/master-orchestrator/src/federated_loop.py` | 50 | Python | `build_task_token` |
| `services/master-orchestrator/src/federated_loop.py` | 60 | Python | `run_global_task` |
| `services/master-orchestrator/src/kafka_producer.py` | 14 | Python | `emit_decision` |
| `services/master-orchestrator/src/main.py` | 23 | Python | `_load_emit_decision` |
| `services/master-orchestrator/src/main.py` | 103 | Python | `_validate_base_url` |
| `services/master-orchestrator/src/main.py` | 121 | Python | `safe_call` |
| `services/master-orchestrator/src/main.py` | 179 | Python | `healthz` |
| `services/master-orchestrator/src/main.py` | 184 | Python | `create_deployment` |
| `services/master-orchestrator/src/main.py` | 200 | Python | `get_deployment` |
| `services/master-orchestrator/src/main.py` | 209 | Python | `handle_deployment_event` |
| `services/master-orchestrator/src/main.py` | 234 | Python | `run_cycle_endpoint` |
| `services/master-orchestrator/src/main.py` | 239 | Python | `run_economy_endpoint` |
| `services/master-orchestrator/src/main.py` | 244 | Python | `run_federated_task_endpoint` |
| `services/master-orchestrator/src/main.py` | 248 | Python | `build_tracked_destination` |
| `services/master-orchestrator/src/main.py` | 254 | Python | `create_checkout` |
| `services/master-orchestrator/src/main.py` | 274 | Python | `run_campaign` |
| `services/master-orchestrator/src/main.py` | 315 | Python | `activate_profit_mode` |
| `services/master-orchestrator/src/main.py` | 378 | Python | `get_server_config` |
| `services/master-orchestrator/src/rl_loop.py` | 4 | Python | `decide_and_scale` |
| `services/master-orchestrator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/master-orchestrator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/master-orchestrator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/master-orchestrator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/model-registry/src/main.py` | 19 | Python | `_ensure_writable_directory` |
| `services/model-registry/src/main.py` | 35 | Python | `_allowed_source_roots` |
| `services/model-registry/src/main.py` | 42 | Python | `_safe_model_component` |
| `services/model-registry/src/main.py` | 49 | Python | `_resolve_source_file` |
| `services/model-registry/src/main.py` | 73 | Python | `_resolve_destination` |
| `services/model-registry/src/main.py` | 81 | Python | `_is_within_roots` |
| `services/model-registry/src/main.py` | 103 | Python | `_atomic_copy` |
| `services/model-registry/src/main.py` | 129 | Python | `healthz` |
| `services/model-registry/src/main.py` | 138 | Python | `register` |
| `services/model-registry/src/main.py` | 154 | Python | `latest` |
| `services/model-registry/src/main.py` | 159 | Python | `get_server_config` |
| `services/model-registry/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/model-registry/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/model-registry/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/model-registry/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/model-service/src/async_queue.py` | 23 | Python | `from_env` |
| `services/model-service/src/async_queue.py` | 32 | Python | `enqueue` |
| `services/model-service/src/main.py` | 24 | Python | `_load_local_metrics_module` |
| `services/model-service/src/main.py` | 67 | Python | `lifespan` |
| `services/model-service/src/main.py` | 75 | Python | `featurize` |
| `services/model-service/src/main.py` | 81 | Python | `fallback_predict` |
| `services/model-service/src/main.py` | 87 | Python | `predict` |
| `services/model-service/src/main.py` | 97 | Python | `predict_from_payload` |
| `services/model-service/src/main.py` | 107 | Python | `healthz` |
| `services/model-service/src/main.py` | 118 | Python | `metrics` |
| `services/model-service/src/main.py` | 123 | Python | `predict_api` |
| `services/model-service/src/main.py` | 128 | Python | `predict_async` |
| `services/model-service/src/main.py` | 154 | Python | `fetch_result` |
| `services/model-service/src/main.py` | 162 | Python | `wait_result` |
| `services/model-service/src/main.py` | 174 | Python | `ws_result` |
| `services/model-service/src/main.py` | 188 | Python | `get_server_config` |
| `services/model-service/src/metrics.py` | 27 | Python | `inc` |
| `services/model-service/src/metrics.py` | 30 | Python | `observe` |
| `services/model-service/src/metrics.py` | 33 | Python | `labels` |
| `services/model-service/src/onnx_model.py` | 12 | Python | `__init__` |
| `services/model-service/src/onnx_model.py` | 22 | Python | `_load` |
| `services/model-service/src/onnx_model.py` | 48 | Python | `warm_up` |
| `services/model-service/src/onnx_model.py` | 55 | Python | `input_width` |
| `services/model-service/src/onnx_model.py` | 64 | Python | `predict` |
| `services/model-service/src/queue_runtime.py` | 15 | Python | `_load_local_metrics_module` |
| `services/model-service/src/queue_runtime.py` | 42 | Python | `_consumer` |
| `services/model-service/src/queue_runtime.py` | 56 | Python | `_decode_message` |
| `services/model-service/src/queue_runtime.py` | 63 | Python | `_publish` |
| `services/model-service/src/queue_runtime.py` | 68 | Python | `_request_loop` |
| `services/model-service/src/queue_runtime.py` | 103 | Python | `_response_loop` |
| `services/model-service/src/queue_runtime.py` | 127 | Python | `start_background_consumers` |
| `services/model-service/src/ray_serve.py` | 18 | Python | `__init__` |
| `services/model-service/src/ray_serve.py` | 21 | Python | `__call__` |
| `services/model-service/src/result_store.py` | 17 | Python | `_sanitize_log_value` |
| `services/model-service/src/result_store.py` | 22 | Python | `__init__` |
| `services/model-service/src/result_store.py` | 25 | Python | `_build_client` |
| `services/model-service/src/result_store.py` | 31 | Python | `_key` |
| `services/model-service/src/result_store.py` | 34 | Python | `set_result` |
| `services/model-service/src/result_store.py` | 52 | Python | `get_result` |
| `services/model-service/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/model-service/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/model-service/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/model-service/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/model-sync/src/main.py` | 16 | Python | `sync` |
| `services/model-sync/src/monitor.py` | 9 | Python | `loop` |
| `services/model-sync/src/rollback.py` | 11 | Python | `backup` |
| `services/model-sync/src/rollback.py` | 21 | Python | `rollback` |
| `services/model-sync/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/model-sync/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/model-sync/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/model-sync/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/model-sync/src/sync.py` | 15 | Python | `upload` |
| `services/model-sync/src/sync.py` | 19 | Python | `download` |
| `services/model-sync/src/sync.py` | 23 | Python | `loop` |
| `services/network-egress/src/client.py` | 17 | Python | `__init__` |
| `services/network-egress/src/client.py` | 38 | Python | `_is_blocked_address` |
| `services/network-egress/src/client.py` | 48 | Python | `_validate_url` |
| `services/network-egress/src/client.py` | 69 | Python | `post` |
| `services/network-egress/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/network-egress/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/network-egress/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/network-egress/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/omniscient/engine.py` | 29 | Python | `analyze` |
| `services/omniscient/engine.py` | 51 | Python | `reason` |
| `services/omniscient/engine.py` | 64 | Python | `fix` |
| `services/omniscient/engine.py` | 71 | Python | `verify` |
| `services/omniscient/engine.py` | 75 | Python | `defend` |
| `services/omniscient/engine.py` | 79 | Python | `run` |
| `services/omniscient/engine.py` | 104 | Python | `_log_event` |
| `services/omniscient/security_ops.py` | 20 | Python | `simulate_attack` |
| `services/omniscient/security_ops.py` | 38 | Python | `deploy_defense` |
| `services/omniscient/security_ops.py` | 45 | Python | `learn` |
| `services/omniscient/security_ops.py` | 52 | Python | `risk_score` |
| `services/org/legal_gate.py` | 5 | Python | `check` |
| `services/org/treasury.py` | 9 | Python | `__init__` |
| `services/org/treasury.py` | 12 | Python | `can_spend` |
| `services/org/treasury.py` | 15 | Python | `spend` |
| `services/p2p-node/src/node.js` | 9 | JS/TS | `startNode` |
| `services/p2p-node/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/p2p-node/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/p2p-node/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/p2p-node/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/p2p/libp2p_bridge.py` | 10 | Python | `start` |
| `services/payment/adapter.py` | 17 | Python | `send` |
| `services/payment/audit.py` | 13 | Python | `log` |
| `services/payment/circuit.py` | 13 | Python | `allow` |
| `services/payment/circuit.py` | 24 | Python | `success` |
| `services/payment/circuit.py` | 31 | Python | `fail` |
| `services/payment/main.py` | 30 | Python | `healthz` |
| `services/payment/main.py` | 35 | Python | `checkout` |
| `services/payment/main.py` | 57 | Python | `get_server_config` |
| `services/payment/stripe_adapter.py` | 11 | Python | `charge` |
| `services/payment/webhook.py` | 17 | Python | `execute` |
| `services/payment/webhook.py` | 21 | Python | `_verify_event` |
| `services/payment/webhook.py` | 41 | Python | `_persist_order` |
| `services/payment/webhook.py` | 59 | Python | `_default_db` |
| `services/payment/webhook.py` | 70 | Python | `execute` |
| `services/payment/webhook.py` | 80 | Python | `stripe_webhook` |
| `services/pr-bot/src/autofix.py` | 15 | Python | `redact_patch_secrets` |
| `services/pr-bot/src/autofix.py` | 19 | Python | `create_auto_fix_pr` |
| `services/pr-bot/src/bot.py` | 10 | Python | `_require_env` |
| `services/pr-bot/src/bot.py` | 17 | Python | `create_pr` |
| `services/pr-bot/src/config.py` | 12 | Python | `_require_env` |
| `services/pr-bot/src/config.py` | 26 | Python | `from_env` |
| `services/pr-bot/src/github_app.py` | 24 | Python | `__init__` |
| `services/pr-bot/src/github_app.py` | 28 | Python | `generate_jwt` |
| `services/pr-bot/src/github_app.py` | 33 | Python | `get_installation_token` |
| `services/pr-bot/src/memory.py` | 21 | Python | `__init__` |
| `services/pr-bot/src/memory.py` | 25 | Python | `_validate_embedding` |
| `services/pr-bot/src/memory.py` | 31 | Python | `insert_memory` |
| `services/pr-bot/src/memory.py` | 54 | Python | `search_similar` |
| `services/pr-bot/src/opa.py` | 6 | Python | `evaluate_gate` |
| `services/pr-bot/src/rl.py` | 18 | Python | `__init__` |
| `services/pr-bot/src/rl.py` | 25 | Python | `select_agents` |
| `services/pr-bot/src/rl.py` | 31 | Python | `update` |
| `services/pr-bot/src/rl.py` | 38 | Python | `snapshot` |
| `services/pr-bot/src/security.py` | 7 | Python | `verify_webhook_signature` |
| `services/pr-bot/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/pr-bot/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/pr-bot/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/pr-bot/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/pr-bot/src/server.py` | 21 | Python | `health` |
| `services/pr-bot/src/server.py` | 26 | Python | `webhook` |
| `services/pr-bot/src/server.py` | 70 | Python | `get_server_config` |
| `services/product-discovery/src/main.py` | 32 | Python | `format` |
| `services/product-discovery/src/main.py` | 59 | Python | `_is_blocked_ip` |
| `services/product-discovery/src/main.py` | 70 | Python | `_validate_market_crawler_url` |
| `services/product-discovery/src/main.py` | 98 | Python | `fetch_products` |
| `services/product-discovery/src/main.py` | 120 | Python | `request_logging_middleware` |
| `services/product-discovery/src/main.py` | 151 | Python | `http_exception_handler` |
| `services/product-discovery/src/main.py` | 156 | Python | `unhandled_exception_handler` |
| `services/product-discovery/src/main.py` | 162 | Python | `healthz` |
| `services/product-discovery/src/main.py` | 182 | Python | `discover` |
| `services/product-discovery/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/product-discovery/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/product-discovery/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/product-discovery/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/product-generator/src/main.py` | 20 | Python | `healthz` |
| `services/product-generator/src/main.py` | 25 | Python | `generate` |
| `services/product-generator/src/main.py` | 42 | Python | `get_server_config` |
| `services/product-generator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/product-generator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/product-generator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/product-generator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/retraining-loop/src/main.py` | 17 | Python | `trigger_once` |
| `services/retraining-loop/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/retraining-loop/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/retraining-loop/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/retraining-loop/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/reward-collector/src/kafka_producer.py` | 14 | Python | `emit_feedback` |
| `services/reward-collector/src/main.py` | 28 | Python | `safe_call` |
| `services/reward-collector/src/main.py` | 41 | Python | `healthz` |
| `services/reward-collector/src/main.py` | 46 | Python | `reward` |
| `services/reward-collector/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/reward-collector/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/reward-collector/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/reward-collector/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/reward-collector/src/server.py` | 4 | Python | `get_server_config` |
| `services/reward-collector/src/server.py` | 12 | Python | `run` |
| `services/rl-coordinator/src/main.py` | 35 | Python | `call_agent` |
| `services/rl-coordinator/src/main.py` | 49 | Python | `healthz` |
| `services/rl-coordinator/src/main.py` | 71 | Python | `decide` |
| `services/rl-coordinator/src/main.py` | 96 | Python | `update` |
| `services/rl-coordinator/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/rl-coordinator/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/rl-coordinator/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/rl-coordinator/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/rl-coordinator/src/server.py` | 4 | Python | `get_agents` |
| `services/rl-coordinator/src/server.py` | 12 | Python | `run` |
| `services/rl-engine/src/linucb.py` | 7 | Python | `__init__` |
| `services/rl-engine/src/linucb.py` | 21 | Python | `add_arm` |
| `services/rl-engine/src/linucb.py` | 28 | Python | `select` |
| `services/rl-engine/src/linucb.py` | 45 | Python | `update` |
| `services/rl-engine/src/main.py` | 53 | Python | `to_vector` |
| `services/rl-engine/src/main.py` | 59 | Python | `load_agent` |
| `services/rl-engine/src/main.py` | 72 | Python | `persist_agent` |
| `services/rl-engine/src/main.py` | 81 | Python | `load_policy_model` |
| `services/rl-engine/src/main.py` | 93 | Python | `score_with_policy` |
| `services/rl-engine/src/main.py` | 107 | Python | `healthz` |
| `services/rl-engine/src/main.py` | 123 | Python | `select` |
| `services/rl-engine/src/main.py` | 154 | Python | `update` |
| `services/rl-engine/src/policy_model.py` | 10 | Python | `__init__` |
| `services/rl-engine/src/policy_model.py` | 16 | Python | `forward` |
| `services/rl-engine/src/policy_model.py` | 25 | Python | `load_state_dict` |
| `services/rl-engine/src/policy_model.py` | 32 | Python | `load` |
| `services/rl-engine/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/rl-engine/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/rl-engine/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/rl-engine/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/rl-engine/src/server.py` | 4 | Python | `run` |
| `services/rl-engine/src/store.py` | 15 | Python | `db_connection` |
| `services/rl-engine/src/store.py` | 24 | Python | `save_state` |
| `services/rl-engine/src/store.py` | 28 | Python | `load_state` |
| `services/rl-engine/src/store.py` | 33 | Python | `log_decision` |
| `services/rl-policy/src/export_onnx.py` | 16 | Python | `export_policy_to_onnx` |
| `services/rl-policy/src/main.py` | 26 | Python | `healthz` |
| `services/rl-policy/src/main.py` | 31 | Python | `train` |
| `services/rl-policy/src/model.py` | 10 | Python | `__init__` |
| `services/rl-policy/src/model.py` | 16 | Python | `forward` |
| `services/rl-policy/src/model.py` | 25 | Python | `train_step` |
| `services/rl-policy/src/model.py` | 31 | Python | `save` |
| `services/rl-policy/src/model.py` | 42 | Python | `load` |
| `services/rl-policy/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/rl-policy/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/rl-policy/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/rl-policy/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/rl-policy/src/train.py` | 13 | Python | `train_episode` |
| `services/rl-trainer/src/agent_replicator.py` | 28 | Python | `__init__` |
| `services/rl-trainer/src/agent_replicator.py` | 38 | Python | `deploy` |
| `services/rl-trainer/src/agent_replicator.py` | 48 | Python | `replicate` |
| `services/rl-trainer/src/agent_replicator.py` | 60 | Python | `parse_targets` |
| `services/rl-trainer/src/agent_replicator.py` | 78 | Python | `build_remote_command` |
| `services/rl-trainer/src/agent_replicator.py` | 90 | Python | `build_ssh_command` |
| `services/rl-trainer/src/agent_replicator.py` | 107 | Python | `_validate_target_inputs` |
| `services/rl-trainer/src/agent_replicator.py` | 115 | Python | `_validate_runtime` |
| `services/rl-trainer/src/agent_replicator.py` | 120 | Python | `_validate_image` |
| `services/rl-trainer/src/agent_replicator.py` | 132 | Python | `_execute_ssh_command` |
| `services/rl-trainer/src/autonomy/ai_sre.py` | 24 | Python | `detect` |
| `services/rl-trainer/src/autonomy/ai_sre.py` | 34 | Python | `evaluate` |
| `services/rl-trainer/src/autonomy/infra_tuner_pipeline.py` | 42 | Python | `parse_args` |
| `services/rl-trainer/src/autonomy/infra_tuner_pipeline.py` | 50 | Python | `load_states` |
| `services/rl-trainer/src/autonomy/infra_tuner_pipeline.py` | 78 | Python | `to_matrix` |
| `services/rl-trainer/src/autonomy/infra_tuner_pipeline.py` | 95 | Python | `choose_action` |
| `services/rl-trainer/src/autonomy/infra_tuner_pipeline.py` | 109 | Python | `build_policy` |
| `services/rl-trainer/src/autonomy/infra_tuner_pipeline.py` | 144 | Python | `main` |
| `services/rl-trainer/src/autonomy/p2p.py` | 13 | Python | `default_envelope` |
| `services/rl-trainer/src/autonomy/redteam.py` | 20 | Python | `allow` |
| `services/rl-trainer/src/autonomy/redteam.py` | 24 | Python | `sample_payload` |
| `services/rl-trainer/src/autonomy/redteam.py` | 38 | Python | `plan` |
| `services/rl-trainer/src/autonomy/simulation.py` | 14 | Python | `policy` |
| `services/rl-trainer/src/autonomy/simulation.py` | 22 | Python | `__init__` |
| `services/rl-trainer/src/autonomy/simulation.py` | 25 | Python | `step` |
| `services/rl-trainer/src/autonomy/simulation.py` | 29 | Python | `metrics` |
| `services/rl-trainer/src/causal_rl.py` | 4 | Python | `doubly_robust` |
| `services/rl-trainer/src/civilization.py` | 7 | Python | `__init__` |
| `services/rl-trainer/src/civilization.py` | 11 | Python | `act` |
| `services/rl-trainer/src/civilization.py` | 18 | Python | `__init__` |
| `services/rl-trainer/src/civilization.py` | 21 | Python | `step` |
| `services/rl-trainer/src/civilization.py` | 25 | Python | `metrics` |
| `services/rl-trainer/src/compute_market.py` | 16 | Python | `__init__` |
| `services/rl-trainer/src/compute_market.py` | 19 | Python | `register` |
| `services/rl-trainer/src/compute_market.py` | 29 | Python | `bid` |
| `services/rl-trainer/src/compute_market.py` | 44 | Python | `assign` |
| `services/rl-trainer/src/global_strategy.py` | 7 | Python | `__init__` |
| `services/rl-trainer/src/global_strategy.py` | 12 | Python | `update` |
| `services/rl-trainer/src/global_strategy.py` | 19 | Python | `equilibrium` |
| `services/rl-trainer/src/global_strategy.py` | 22 | Python | `coordinate` |
| `services/rl-trainer/src/hierarchical_rl.py` | 8 | Python | `_initial_weights` |
| `services/rl-trainer/src/hierarchical_rl.py` | 11 | Python | `__init__` |
| `services/rl-trainer/src/hierarchical_rl.py` | 18 | Python | `_dot` |
| `services/rl-trainer/src/hierarchical_rl.py` | 22 | Python | `_vector` |
| `services/rl-trainer/src/hierarchical_rl.py` | 25 | Python | `select` |
| `services/rl-trainer/src/hierarchical_rl.py` | 33 | Python | `update` |
| `services/rl-trainer/src/hybrid_rl.py` | 10 | Python | `_initial_policy_weights` |
| `services/rl-trainer/src/hybrid_rl.py` | 13 | Python | `__init__` |
| `services/rl-trainer/src/hybrid_rl.py` | 19 | Python | `select_arm` |
| `services/rl-trainer/src/hybrid_rl.py` | 28 | Python | `update` |
| `services/rl-trainer/src/long_term_reward.py` | 4 | Python | `long_term_reward` |
| `services/rl-trainer/src/main.py` | 16 | Python | `__init__` |
| `services/rl-trainer/src/main.py` | 19 | Python | `subscribe` |
| `services/rl-trainer/src/main.py` | 22 | Python | `poll` |
| `services/rl-trainer/src/main.py` | 65 | Python | `build_autonomous_snapshot` |
| `services/rl-trainer/src/main.py` | 106 | Python | `_redact_snapshot` |
| `services/rl-trainer/src/main.py` | 118 | Python | `_snapshot_fingerprint` |
| `services/rl-trainer/src/main.py` | 123 | Python | `loop` |
| `services/rl-trainer/src/meta_learning.py` | 7 | Python | `__init__` |
| `services/rl-trainer/src/meta_learning.py` | 12 | Python | `adapt` |
| `services/rl-trainer/src/p2p_agent.py` | 11 | Python | `__init__` |
| `services/rl-trainer/src/p2p_agent.py` | 23 | Python | `_validate_host` |
| `services/rl-trainer/src/p2p_agent.py` | 35 | Python | `start_server` |
| `services/rl-trainer/src/p2p_agent.py` | 36 | Python | `run` |
| `services/rl-trainer/src/p2p_agent.py` | 56 | Python | `gossip` |
| `services/rl-trainer/src/ppo.py` | 7 | Python | `__init__` |
| `services/rl-trainer/src/ppo.py` | 12 | Python | `policy` |
| `services/rl-trainer/src/ppo.py` | 16 | Python | `update` |
| `services/rl-trainer/src/reward.py` | 4 | Python | `compute_reward` |
| `services/rl-trainer/src/rllib_ppo.py` | 6 | Python | `train_distributed_ppo` |
| `services/rl-trainer/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/rl-trainer/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/rl-trainer/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/rl-trainer/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/rl-trainer/src/sovereign_identity.py` | 15 | Python | `_b64` |
| `services/rl-trainer/src/sovereign_identity.py` | 19 | Python | `load_or_create_key` |
| `services/rl-trainer/src/sovereign_identity.py` | 35 | Python | `get_did` |
| `services/rl-trainer/src/sovereign_identity.py` | 43 | Python | `export_public_key` |
| `services/rl-trainer/src/sovereign_identity.py` | 51 | Python | `sign` |
| `services/rl-trainer/src/sovereign_identity.py` | 56 | Python | `build_heartbeat_identity` |
| `services/rl-trainer/src/token_engine.py` | 5 | Python | `__init__` |
| `services/rl-trainer/src/token_engine.py` | 9 | Python | `stake` |
| `services/rl-trainer/src/token_engine.py` | 13 | Python | `reward` |
| `services/rl-trainer/src/token_engine.py` | 17 | Python | `slash` |
| `services/rl-trainer/src/treasury.py` | 7 | Python | `__init__` |
| `services/rl-trainer/src/treasury.py` | 11 | Python | `hedge` |
| `services/rl-trainer/src/treasury.py` | 15 | Python | `allocate` |
| `services/rl-trainer/src/uplift_model.py` | 8 | Python | `__init__` |
| `services/rl-trainer/src/uplift_model.py` | 12 | Python | `fit` |
| `services/rl-trainer/src/uplift_model.py` | 20 | Python | `predict_uplift` |
| `services/rl-trainer/src/world_model.py` | 7 | Python | `__init__` |
| `services/rl-trainer/src/world_model.py` | 11 | Python | `predict_next` |
| `services/rl-trainer/src/world_model.py` | 15 | Python | `update` |
| `services/rl-trainer/tests/test_security_hardening.py` | 16 | Python | `_parse_module` |
| `services/rl-trainer/tests/test_security_hardening.py` | 21 | Python | `_iter_calls` |
| `services/rl-trainer/tests/test_security_hardening.py` | 25 | Python | `test_parse_targets_rejects_invalid_host` |
| `services/rl-trainer/tests/test_security_hardening.py` | 30 | Python | `test_build_remote_command_rejects_invalid_image` |
| `services/rl-trainer/tests/test_security_hardening.py` | 35 | Python | `test_sample_payload_is_deterministic` |
| `services/rl-trainer/tests/test_security_hardening.py` | 43 | Python | `test_plan_uses_deterministic_payloads` |
| `services/rl-trainer/tests/test_security_hardening.py` | 50 | Python | `test_agent_replicator_does_not_use_subprocess_run` |
| `services/rl-trainer/tests/test_security_hardening.py` | 61 | Python | `test_redteam_does_not_use_non_deterministic_random_calls` |
| `services/rtb-engine/src/causal_rl.py` | 4 | Python | `doubly_robust` |
| `services/rtb-engine/src/hierarchical_rl.py` | 7 | Python | `__init__` |
| `services/rtb-engine/src/hierarchical_rl.py` | 14 | Python | `_dot` |
| `services/rtb-engine/src/hierarchical_rl.py` | 18 | Python | `_vector` |
| `services/rtb-engine/src/hierarchical_rl.py` | 21 | Python | `select` |
| `services/rtb-engine/src/hierarchical_rl.py` | 29 | Python | `update` |
| `services/rtb-engine/src/latency_bid.py` | 1 | Python | `latency_adjusted_bid` |
| `services/rtb-engine/src/long_term_reward.py` | 4 | Python | `long_term_reward` |
| `services/rtb-engine/src/ltv_model.py` | 7 | Python | `predict` |
| `services/rtb-engine/src/main.py` | 75 | Python | `healthz` |
| `services/rtb-engine/src/main.py` | 80 | Python | `bid` |
| `services/rtb-engine/src/main.py` | 117 | Python | `openrtb_bid` |
| `services/rtb-engine/src/main.py` | 135 | Python | `_env_flag` |
| `services/rtb-engine/src/main.py` | 142 | Python | `resolve_bind_host` |
| `services/rtb-engine/src/rtb.py` | 5 | Python | `__init__` |
| `services/rtb-engine/src/rtb.py` | 9 | Python | `compute_bid` |
| `services/rtb-engine/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/rtb-engine/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/rtb-engine/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/rtb-engine/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/scaling-engine/src/main.py` | 24 | Python | `healthz` |
| `services/scaling-engine/src/main.py` | 29 | Python | `scale` |
| `services/scaling-engine/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/scaling-engine/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/scaling-engine/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/scaling-engine/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/scheduler/src/main.py` | 39 | Python | `healthz` |
| `services/scheduler/src/main.py` | 54 | Python | `fetch_nodes` |
| `services/scheduler/src/main.py` | 64 | Python | `decode_task_token` |
| `services/scheduler/src/main.py` | 83 | Python | `verify_task_token` |
| `services/scheduler/src/main.py` | 93 | Python | `assign` |
| `services/scheduler/src/main.py` | 113 | Python | `score` |
| `services/scheduler/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/scheduler/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/scheduler/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/scheduler/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/security-api/src/app/deps.py` | 22 | Python | `__init__` |
| `services/security-api/src/app/deps.py` | 26 | Python | `save` |
| `services/security-api/src/app/deps.py` | 30 | Python | `get` |
| `services/security-api/src/app/deps.py` | 37 | Python | `get_scan_service` |
| `services/security-api/src/app/deps.py` | 41 | Python | `validate_repo_url` |
| `services/security-api/src/app/deps.py` | 54 | Python | `require_api_key` |
| `services/security-api/src/app/deps.py` | 61 | Python | `_run_scan` |
| `services/security-api/src/app/deps.py` | 66 | Python | `enqueue_scan` |
| `services/security-api/src/app/deps.py` | 86 | Python | `get_scan_payload` |
| `services/security-api/src/app/main.py` | 19 | Python | `_configure_logging` |
| `services/security-api/src/app/main.py` | 28 | Python | `_security_headers_middleware` |
| `services/security-api/src/app/main.py` | 29 | Python | `security_headers` |
| `services/security-api/src/app/main.py` | 39 | Python | `create_app` |
| `services/security-api/src/app/routes.py` | 8 | Python | `build_router` |
| `services/security-api/src/app/routes.py` | 13 | Python | `healthz` |
| `services/security-api/src/app/routes.py` | 18 | Python | `run_scan` |
| `services/security-api/src/app/routes.py` | 29 | Python | `get_scan` |
| `services/security-api/src/container/di.py` | 10 | Python | `__init__` |
| `services/security-api/src/container/di.py` | 13 | Python | `register` |
| `services/security-api/src/container/di.py` | 16 | Python | `resolve` |
| `services/security-api/src/container/di.py` | 25 | Python | `init_container` |
| `services/security-api/src/domain/interfaces.py` | 13 | Python | `scan` |
| `services/security-api/src/domain/runtime_policy.py` | 10 | Python | `normalize_language` |
| `services/security-api/src/domain/runtime_policy.py` | 16 | Python | `assert_allowed` |
| `services/security-api/src/infra/backpressure.py` | 12 | Python | `run_limited` |
| `services/security-api/src/infra/config.py` | 6 | Python | `_ensure_under_base` |
| `services/security-api/src/infra/config.py` | 16 | Python | `get_secure_tmp_dir` |
| `services/security-api/src/infra/config.py` | 28 | Python | `configure_process_umask` |
| `services/security-api/src/infra/config.py` | 32 | Python | `get_server_config` |
| `services/security-api/src/infra/process.py` | 11 | Python | `safe_exec` |
| `services/security-api/src/infra/rate_limit.py` | 16 | Python | `rate_limit` |
| `services/security-api/src/infra/rate_limit.py` | 17 | Python | `wrapper` |
| `services/security-api/src/infra/rate_limit.py` | 27 | Python | `rate_limit` |
| `services/security-api/src/infra/rate_limit.py` | 28 | Python | `wrapper` |
| `services/security-api/src/infra/rate_limit.py` | 30 | Python | `passthrough` |
| `services/security-api/src/infra/scan_service.py` | 25 | Python | `resolve_executable` |
| `services/security-api/src/infra/scan_service.py` | 35 | Python | `safe_exec_async` |
| `services/security-api/src/infra/scan_service.py` | 61 | Python | `scan` |
| `services/security-api/src/runtime_engine/base.py` | 17 | Python | `__post_init__` |
| `services/security-api/src/runtime_engine/base.py` | 49 | Python | `execute` |
| `services/security-api/src/runtime_engine/manager.py` | 10 | Python | `__init__` |
| `services/security-api/src/runtime_engine/manager.py` | 16 | Python | `resolve` |
| `services/security-api/src/runtime_engine/manager.py` | 23 | Python | `execute` |
| `services/security-api/src/runtime_engine/node.py` | 12 | Python | `execute` |
| `services/security-api/src/runtime_engine/python.py` | 12 | Python | `execute` |
| `services/security-api/src/runtime_engine/runner.py` | 10 | Python | `_resolve_entrypoint` |
| `services/security-api/src/runtime_engine/runner.py` | 27 | Python | `run_exec` |
| `services/security-api/src/server.py` | 7 | Python | `run` |
| `services/security-api/src/shared/event_bus/interface.py` | 10 | Python | `publish` |
| `services/security-api/src/shared/event_bus/interface.py` | 14 | Python | `subscribe` |
| `services/security-api/src/shared/event_bus/nats_bus.py` | 11 | Python | `__init__` |
| `services/security-api/src/shared/event_bus/nats_bus.py` | 15 | Python | `connect` |
| `services/security-api/src/shared/event_bus/nats_bus.py` | 23 | Python | `publish` |
| `services/security-api/src/shared/event_bus/nats_bus.py` | 29 | Python | `subscribe` |
| `services/security-api/src/shared/event_bus/nats_bus.py` | 34 | Python | `callback` |
| `services/security-api/src/shared/queue/async_queue.py` | 10 | Python | `producer` |
| `services/security-api/src/shared/queue/async_queue.py` | 14 | Python | `consumer` |
| `services/security-api/src/shared/security/auth.py` | 7 | Python | `verify_internal_request` |
| `services/security-api/src/shared/security/crypto.py` | 8 | Python | `sign` |
| `services/security-api/src/shared/security/crypto_impl.py` | 10 | Python | `__init__` |
| `services/security-api/src/shared/security/crypto_impl.py` | 13 | Python | `sign` |
| `services/security-api/src/shared/security/secrets.py` | 7 | Python | `get` |
| `services/security-api/src/shared/security/validation.py` | 8 | Python | `validate_target` |
| `services/self-modifier/engine.py` | 10 | Python | `apply_patch_request` |
| `services/self-modifier/kms_verify.py` | 12 | Python | `verify` |
| `services/shared/event_bus.py` | 11 | Python | `publish` |
| `services/shared/kafka_producer.py` | 7 | Python | `__init__` |
| `services/shared/kafka_producer.py` | 13 | Python | `publish` |
| `services/shared/kafka_secure.py` | 17 | Python | `safe_produce` |
| `services/shared/region_sync.py` | 18 | Python | `broadcast` |
| `services/shared/schema.py` | 6 | Python | `validate` |
| `services/shopee-crawler/src/core/affiliate.js` | 1 | JS/TS | `generateAffiliateLink` |
| `services/shopee-crawler/src/core/crawler.js` | 6 | JS/TS | `crawl` |
| `services/shopee-crawler/src/core/scheduler.js` | 11 | JS/TS | `scheduler` |
| `services/shopee-crawler/src/crawler/search.js` | 3 | JS/TS | `searchProducts` |
| `services/shopee-crawler/src/index.js` | 3 | JS/TS | `main` |
| `services/shopee-crawler/src/parser/product.js` | 1 | JS/TS | `parseProduct` |
| `services/shopee-crawler/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/shopee-crawler/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/shopee-crawler/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/shopee-crawler/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/singularity/src/agents.py` | 25 | Python | `handle` |
| `services/singularity/src/agents.py` | 32 | Python | `handle` |
| `services/singularity/src/agents.py` | 54 | Python | `handle` |
| `services/singularity/src/agents.py` | 75 | Python | `handle` |
| `services/singularity/src/agents.py` | 95 | Python | `handle` |
| `services/singularity/src/agents.py` | 109 | Python | `handle` |
| `services/singularity/src/agents.py` | 118 | Python | `handle` |
| `services/singularity/src/agents.py` | 137 | Python | `handle` |
| `services/singularity/src/agents.py` | 152 | Python | `handle` |
| `services/singularity/src/orchestrator.py` | 17 | Python | `__init__` |
| `services/singularity/src/orchestrator.py` | 21 | Python | `publish` |
| `services/singularity/src/orchestrator.py` | 30 | Python | `consume` |
| `services/singularity/src/orchestrator.py` | 34 | Python | `history` |
| `services/singularity/src/orchestrator.py` | 41 | Python | `handle` |
| `services/singularity/src/orchestrator.py` | 45 | Python | `__init__` |
| `services/singularity/src/orchestrator.py` | 49 | Python | `dispatch` |
| `services/singularity/src/orchestrator.py` | 54 | Python | `run` |
| `services/singularity/src/risk.py` | 4 | Python | `risk` |
| `services/singularity/src/run.py` | 19 | Python | `main` |
| `services/singularity/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/singularity/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/singularity/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/singularity/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/singularity/src/verify.py` | 10 | Python | `_is_url_allowed` |
| `services/singularity/src/verify.py` | 15 | Python | `ssrf_probe` |
| `services/singularity/src/verify.py` | 26 | Python | `path_traversal_probe` |
| `services/stream-consumer/src/main.py` | 34 | Python | `db_connection` |
| `services/stream-consumer/src/main.py` | 43 | Python | `shutdown` |
| `services/stream-consumer/src/main.py` | 49 | Python | `update_online_features` |
| `services/stream-consumer/src/main.py` | 84 | Python | `persist_offline_metrics` |
| `services/stream-consumer/src/main.py` | 101 | Python | `process` |
| `services/stream-consumer/src/main.py` | 108 | Python | `create_consumer` |
| `services/stream-consumer/src/main.py` | 121 | Python | `main` |
| `services/stream-consumer/src/producer.py` | 10 | Python | `emit_decision` |
| `services/stream-consumer/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/stream-consumer/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/stream-consumer/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/stream-consumer/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/tenant-service/src/main.py` | 22 | Python | `hash_api_key` |
| `services/tenant-service/src/main.py` | 39 | Python | `db_connection` |
| `services/tenant-service/src/main.py` | 59 | Python | `healthz` |
| `services/tenant-service/src/main.py` | 73 | Python | `create_tenant` |
| `services/tenant-service/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/tenant-service/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/tenant-service/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/tenant-service/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/tiktok-farm/src/core/accounts.js` | 3 | JS/TS | `getAccount` |
| `services/tiktok-farm/src/core/proxies.js` | 3 | JS/TS | `getProxy` |
| `services/tiktok-farm/src/core/queue.js` | 9 | JS/TS | `enqueue` |
| `services/tiktok-farm/src/core/queue.js` | 18 | JS/TS | `dequeue` |
| `services/tiktok-farm/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/tiktok-farm/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/tiktok-farm/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/tiktok-farm/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/tiktok-farm/src/uploader/uploader.js` | 3 | JS/TS | `upload` |
| `services/tiktok-farm/src/workers/worker.js` | 6 | JS/TS | `worker` |
| `services/tiktok-shop-miner/src/core/miner.js` | 6 | JS/TS | `mine` |
| `services/tiktok-shop-miner/src/core/scheduler.js` | 11 | JS/TS | `scheduler` |
| `services/tiktok-shop-miner/src/core/trending.js` | 1 | JS/TS | `scoreProduct` |
| `services/tiktok-shop-miner/src/crawler/search.js` | 3 | JS/TS | `searchProducts` |
| `services/tiktok-shop-miner/src/index.js` | 3 | JS/TS | `main` |
| `services/tiktok-shop-miner/src/parser/product.js` | 1 | JS/TS | `parseProduct` |
| `services/tiktok-shop-miner/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/tiktok-shop-miner/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/tiktok-shop-miner/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/tiktok-shop-miner/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/tiktok-uploader/src/api.ts` | 5 | JS/TS | `uploadVideo` |
| `services/tiktok-uploader/src/automation/login.js` | 1 | JS/TS | `login` |
| `services/tiktok-uploader/src/automation/upload.js` | 1 | JS/TS | `uploadVideo` |
| `services/tiktok-uploader/src/core/browser.js` | 3 | JS/TS | `createBrowser` |
| `services/tiktok-uploader/src/core/uploader.js` | 5 | JS/TS | `runUploader` |
| `services/tiktok-uploader/src/core/worker.js` | 4 | JS/TS | `worker` |
| `services/tiktok-uploader/src/index.js` | 3 | JS/TS | `main` |
| `services/tiktok-uploader/src/models/account.js` | 3 | JS/TS | `constructor` |
| `services/tiktok-uploader/src/models/video.js` | 3 | JS/TS | `constructor` |
| `services/tiktok-uploader/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/tiktok-uploader/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/tiktok-uploader/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/tiktok-uploader/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/tracking/server.py` | 21 | Python | `_connect` |
| `services/tracking/server.py` | 32 | Python | `_assert_campaign_id` |
| `services/tracking/server.py` | 41 | Python | `_sanitize_log_value` |
| `services/tracking/server.py` | 45 | Python | `_validated_affiliate_base_url` |
| `services/tracking/server.py` | 60 | Python | `healthz` |
| `services/tracking/server.py` | 65 | Python | `track` |
| `services/viral-predictor/src/api/server.py` | 24 | Python | `healthz` |
| `services/viral-predictor/src/api/server.py` | 48 | Python | `metrics` |
| `services/viral-predictor/src/api/server.py` | 53 | Python | `predict_viral` |
| `services/viral-predictor/src/core/database.py` | 4 | Python | `get_db` |
| `services/viral-predictor/src/features/features.py` | 3 | Python | `extract_features` |
| `services/viral-predictor/src/main.py` | 6 | Python | `_safe_port` |
| `services/viral-predictor/src/model/model.py` | 15 | Python | `__init__` |
| `services/viral-predictor/src/model/model.py` | 19 | Python | `forward` |
| `services/viral-predictor/src/model/model.py` | 23 | Python | `save` |
| `services/viral-predictor/src/model/model.py` | 27 | Python | `load` |
| `services/viral-predictor/src/model/model.py` | 37 | Python | `predict` |
| `services/viral-predictor/src/security/rate_limit.py` | 9 | Python | `rate_limit` |
| `services/viral-predictor/src/security/rate_limit.py` | 10 | Python | `wrapper` |
| `services/viral-predictor/src/security/rate_limit.py` | 20 | Python | `rate_limit` |
| `services/viral-predictor/src/security/rate_limit.py` | 21 | Python | `wrapper` |
| `services/viral-predictor/src/training/train.py` | 12 | Python | `_sigmoid` |
| `services/viral-predictor/src/training/train.py` | 16 | Python | `train` |
| `start-zlttbots.sh` | 62 | Shell | `start_background_process` |
| `tests/omniscient/test_engine.py` | 9 | Python | `test_run_returns_actions` |
| `tests/omniscient/test_engine.py` | 19 | Python | `test_simulate_attack_requires_http_scheme` |
| `tests/omniscient/test_engine.py` | 23 | Python | `test_security_helpers` |
| `tests/runtime/test_service_health_import_safe.py` | 8 | Python | `_load_app` |
| `tests/runtime/test_service_health_import_safe.py` | 20 | Python | `test_reward_collector_healthz` |
| `tests/runtime/test_service_health_import_safe.py` | 27 | Python | `test_rl_engine_healthz` |
| `tests/runtime/test_service_health_import_safe.py` | 34 | Python | `test_rl_coordinator_healthz` |
| `tests/security_platform/test_ai_fix_engine.py` | 13 | Python | `test_generate_patch` |
| `tests/security_platform/test_pr_bot_extension.py` | 28 | Python | `test_gate_blocks_critical_and_warns_medium` |
| `tests/security_platform/test_pr_bot_extension.py` | 39 | Python | `test_rl_bandit_updates_weights_deterministically` |
| `tests/security_platform/test_pr_bot_extension.py` | 53 | Python | `test_verify_signature_uses_hmac_sha256` |
| `tests/security_platform/test_pr_bot_extension.py` | 62 | Python | `test_memory_store_validates_embedding_dimensions` |
| `tests/security_platform/test_runtime_engine.py` | 11 | Python | `test_python_runtime_executes_script` |
| `tests/security_platform/test_runtime_engine.py` | 31 | Python | `test_runtime_allowlist_rejects_unknown` |
| `tests/security_platform/test_runtime_request_validation.py` | 10 | Python | `_load_runtime_request` |
| `tests/security_platform/test_runtime_request_validation.py` | 15 | Python | `test_runtime_request_rejects_non_positive_timeout` |
| `tests/security_platform/test_runtime_request_validation.py` | 26 | Python | `test_runtime_request_rejects_non_string_env` |
| `tests/security_platform/test_secure_tmp_dir_validation.py` | 11 | Python | `test_secure_tmp_dir_rejects_traversal` |
| `tests/security_platform/test_security_api.py` | 18 | Python | `test_healthz` |
| `tests/security_platform/test_security_api.py` | 25 | Python | `test_validate_repo_url_rejects_http` |
| `tests/services/security_api_runtime_import.py` | 13 | Python | `_load_module` |
| `tests/services/security_api_runtime_import.py` | 23 | Python | `load_runtime_engine_manager` |
| `tests/test_advanced_rl_features.py` | 10 | Python | `_load_module` |
| `tests/test_advanced_rl_features.py` | 20 | Python | `test_rtb_engine_openrtb_and_reward_fields` |
| `tests/test_advanced_rl_features.py` | 55 | Python | `test_capital_allocator_reports_reinvestment_and_allocations` |
| `tests/test_affiliate_publishing_flow.py` | 14 | Python | `load_module` |
| `tests/test_affiliate_publishing_flow.py` | 24 | Python | `test_affiliate_connector_payout_ingestion_daily_controller_and_reporting` |
| `tests/test_affiliate_publishing_flow.py` | 38 | Python | `fake_snapshot` |
| `tests/test_affiliate_publishing_flow.py` | 122 | Python | `test_dead_letter_after_retry_budget` |
| `tests/test_affiliate_publishing_flow.py` | 155 | Python | `test_arbitrage_scan_and_listing_work_inmemory` |
| `tests/test_ai_economy_saas.py` | 11 | Python | `_load_app` |
| `tests/test_ai_economy_saas.py` | 20 | Python | `test_product_generator_is_deterministic` |
| `tests/test_ai_economy_saas.py` | 34 | Python | `test_billing_service_rates_by_plan` |
| `tests/test_ai_economy_saas.py` | 43 | Python | `test_market_orchestrator_launches_supported_markets` |
| `tests/test_ai_economy_saas.py` | 50 | Python | `__init__` |
| `tests/test_ai_economy_saas.py` | 53 | Python | `raise_for_status` |
| `tests/test_ai_economy_saas.py` | 56 | Python | `json` |
| `tests/test_ai_economy_saas.py` | 59 | Python | `fake_post` |
| `tests/test_ai_economy_saas.py` | 69 | Python | `test_master_orchestrator_economy_endpoint` |
| `tests/test_ai_economy_saas.py` | 78 | Python | `__init__` |
| `tests/test_ai_economy_saas.py` | 81 | Python | `raise_for_status` |
| `tests/test_ai_economy_saas.py` | 84 | Python | `json` |
| `tests/test_ai_economy_saas.py` | 87 | Python | `fake_get` |
| `tests/test_ai_economy_saas.py` | 91 | Python | `fake_post` |
| `tests/test_auto_ml_loop_security.py` | 15 | Python | `test_run_command_uses_argument_list_without_shell` |
| `tests/test_auto_ml_loop_security.py` | 18 | Python | `fake_run` |
| `tests/test_auto_ml_loop_security.py` | 32 | Python | `test_evaluate_rejects_non_numeric_output` |
| `tests/test_auto_ml_loop_security.py` | 47 | Python | `test_loop_returns_evaluate_failed_when_eval_command_fails` |
| `tests/test_autonomous_rl_services.py` | 13 | Python | `_load_module` |
| `tests/test_autonomous_rl_services.py` | 23 | Python | `_load_app` |
| `tests/test_autonomous_rl_services.py` | 27 | Python | `test_model_registry_registers_and_returns_latest` |
| `tests/test_autonomous_rl_services.py` | 48 | Python | `test_reward_collector_clips_profit_reward` |
| `tests/test_autonomous_rl_services.py` | 54 | Python | `fake_safe_call` |
| `tests/test_autonomous_rl_services.py` | 66 | Python | `test_rl_engine_uses_policy_model_when_present` |
| `tests/test_autonomous_rl_services.py` | 75 | Python | `eval` |
| `tests/test_autonomous_rl_services.py` | 78 | Python | `__call__` |
| `tests/test_autonomous_rl_services.py` | 94 | Python | `test_budget_allocator_exposes_market_budget` |
| `tests/test_autonomous_rl_services.py` | 112 | Python | `test_rtb_engine_uses_floor_bid` |
| `tests/test_autonomous_rl_services.py` | 131 | Python | `test_tlearner_predicts_uplift` |
| `tests/test_autonomous_rl_services.py` | 145 | Python | `test_identity_signatures_round_trip` |
| `tests/test_autonomous_rl_services.py` | 160 | Python | `test_exchange_matches_orders_with_signature` |
| `tests/test_autonomous_rl_services.py` | 169 | Python | `signed_order` |
| `tests/test_autonomous_rl_services.py` | 192 | Python | `test_rl_trainer_autonomous_snapshot_includes_new_subsystems` |
| `tests/test_distributed_rl_stack.py` | 11 | Python | `_load_app` |
| `tests/test_distributed_rl_stack.py` | 20 | Python | `test_budget_allocator_enforces_step_and_cap` |
| `tests/test_distributed_rl_stack.py` | 44 | Python | `test_rtb_engine_clamps_bid_and_applies_pacing` |
| `tests/test_distributed_rl_stack.py` | 66 | Python | `test_rl_coordinator_picks_highest_score` |
| `tests/test_distributed_rl_stack.py` | 73 | Python | `__init__` |
| `tests/test_distributed_rl_stack.py` | 76 | Python | `raise_for_status` |
| `tests/test_distributed_rl_stack.py` | 79 | Python | `json` |
| `tests/test_distributed_rl_stack.py` | 82 | Python | `fake_post` |
| `tests/test_distributed_rl_stack.py` | 96 | Python | `test_rtb_engine_applies_latency_penalty` |
| `tests/test_enterprise_maturity.py` | 13 | Python | `test_secret_manager_rotation_due` |
| `tests/test_enterprise_maturity.py` | 21 | Python | `test_rbac_policy_and_audit_pipeline` |
| `tests/test_enterprise_maturity.py` | 32 | Python | `test_retry_and_idempotency_controls` |
| `tests/test_enterprise_maturity.py` | 35 | Python | `flaky` |
| `tests/test_enterprise_maturity.py` | 47 | Python | `test_circuit_breaker_opens` |
| `tests/test_enterprise_maturity.py` | 61 | Python | `test_slo_alerting_and_probe` |
| `tests/test_enterprise_maturity.py` | 75 | Python | `test_governance_scaling_and_queue_admission` |
| `tests/test_enterprise_maturity.py` | 89 | Python | `test_roadmap_has_all_items` |
| `tests/test_enterprise_maturity.py` | 95 | Python | `test_full_upgrade_blueprint_covers_requested_capabilities` |
| `tests/test_enterprise_maturity.py` | 104 | Python | `test_full_upgrade_blueprint_validation_rejects_small_layout` |
| `tests/test_extract_gemini_source.py` | 6 | Python | `test_extracts_patch_and_heredoc` |
| `tests/test_extract_gemini_source.py` | 31 | Python | `test_extracts_hinted_file` |
| `tests/test_feature_impact_dive.py` | 22 | Python | `test_extract_compose_services_reads_services_block` |
| `tests/test_feature_impact_dive.py` | 41 | Python | `test_extract_app_features_discovers_endpoints_from_multiple_files` |
| `tests/test_feature_impact_dive.py` | 63 | Python | `test_extract_runtime_service_features_supports_python_decorators` |
| `tests/test_feature_impact_dive.py` | 69 | Python | `healthz` |
| `tests/test_feature_impact_dive.py` | 82 | Python | `test_extract_documented_features_reads_titles` |
| `tests/test_feature_impact_dive.py` | 92 | Python | `test_build_surface_manifest_contains_write_endpoint_policy` |
| `tests/test_feature_impact_dive.py` | 107 | Python | `test_validate_manifest_detects_drift` |
| `tests/test_feature_impact_dive.py` | 126 | Python | `test_validate_manifest_reports_json_decode_errors` |
| `tests/test_feature_impact_dive.py` | 140 | Python | `test_validate_manifest_rejects_non_string_policy_fields` |
| `tests/test_feature_impact_dive.py` | 159 | Python | `test_format_markdown_contains_counts` |
| `tests/test_feature_impact_dive.py` | 182 | Python | `test_manifest_to_dict_is_json_serializable` |
| `tests/test_feature_impact_dive.py` | 194 | Python | `test_validate_manifest_rejects_missing_required_keys` |
| `tests/test_feature_impact_dive.py` | 208 | Python | `test_resolve_output_path_rejects_traversal` |
| `tests/test_frontend_auth_security_defaults.py` | 4 | Python | `test_frontend_does_not_ship_seeded_credentials` |
| `tests/test_frontend_auth_security_defaults.py` | 10 | Python | `test_frontend_requires_credentials_before_auth_calls` |
| `tests/test_go_to_market_system.py` | 13 | Python | `test_generate_daily_plan_phases_and_validation` |
| `tests/test_go_to_market_system.py` | 25 | Python | `test_generate_post_is_deterministic_by_slot` |
| `tests/test_go_to_market_system.py` | 37 | Python | `test_pitch_deck_and_status_contains_metrics` |
| `tests/test_gpu_renderer_no_cost.py` | 13 | Python | `test_build_ffmpeg_command_defaults_to_cpu_when_no_gpu` |
| `tests/test_gpu_renderer_no_cost.py` | 25 | Python | `test_build_ffmpeg_command_uses_cuda_when_forced` |
| `tests/test_infra_hardening_modules.py` | 11 | Python | `load_module` |
| `tests/test_infra_hardening_modules.py` | 21 | Python | `test_payment_circuit_opens_after_repeated_failures` |
| `tests/test_infra_hardening_modules.py` | 31 | Python | `test_payment_audit_writes_json_line` |
| `tests/test_infra_hardening_modules.py` | 40 | Python | `test_cost_control_enforces_daily_budget` |
| `tests/test_infra_hardening_modules.py` | 49 | Python | `test_legal_gate_filters_jurisdictions` |
| `tests/test_infra_hardening_modules.py` | 55 | Python | `test_treasury_caps_spending` |
| `tests/test_jwt_auth_service.py` | 7 | Python | `_load_jwt_auth_app` |
| `tests/test_jwt_auth_service.py` | 16 | Python | `test_token_and_introspect_round_trip` |
| `tests/test_jwt_auth_service.py` | 31 | Python | `test_token_rejects_invalid_subject` |
| `tests/test_model_service_async.py` | 10 | Python | `__init__` |
| `tests/test_model_service_async.py` | 13 | Python | `__getitem__` |
| `tests/test_model_service_async.py` | 16 | Python | `__len__` |
| `tests/test_model_service_async.py` | 19 | Python | `item` |
| `tests/test_model_service_async.py` | 29 | Python | `tensor` |
| `tests/test_model_service_async.py` | 33 | Python | `dot` |
| `tests/test_model_service_async.py` | 38 | Python | `__init__` |
| `tests/test_model_service_async.py` | 41 | Python | `produce` |
| `tests/test_model_service_async.py` | 44 | Python | `flush` |
| `tests/test_model_service_async.py` | 49 | Python | `__init__` |
| `tests/test_model_service_async.py` | 52 | Python | `subscribe` |
| `tests/test_model_service_async.py` | 55 | Python | `poll` |
| `tests/test_model_service_async.py` | 60 | Python | `__init__` |
| `tests/test_model_service_async.py` | 63 | Python | `set_result` |
| `tests/test_model_service_async.py` | 68 | Python | `get_result` |
| `tests/test_model_service_async.py` | 73 | Python | `inc` |
| `tests/test_model_service_async.py` | 76 | Python | `observe` |
| `tests/test_model_service_async.py` | 79 | Python | `labels` |
| `tests/test_model_service_async.py` | 83 | Python | `load_main_module` |
| `tests/test_model_service_async.py` | 103 | Python | `test_predict_async_accepts_idempotency_key` |
| `tests/test_model_service_async.py` | 119 | Python | `test_fetch_result_returns_pending_when_missing` |
| `tests/test_network_bind_hardening.py` | 10 | Python | `_load_module` |
| `tests/test_network_bind_hardening.py` | 20 | Python | `test_p2p_agent_rejects_wildcard_bind_by_default` |
| `tests/test_network_bind_hardening.py` | 26 | Python | `test_p2p_agent_accepts_loopback_address` |
| `tests/test_network_bind_hardening.py` | 32 | Python | `test_resolve_bind_host_disables_wildcard_without_explicit_opt_in` |
| `tests/test_network_bind_hardening.py` | 39 | Python | `test_resolve_bind_host_allows_wildcard_with_explicit_opt_in` |
| `tests/test_network_egress_client.py` | 17 | Python | `__init__` |
| `tests/test_network_egress_client.py` | 21 | Python | `test_retries_on_server_error_and_succeeds` |
| `tests/test_network_egress_client.py` | 24 | Python | `fake_post` |
| `tests/test_network_egress_client.py` | 39 | Python | `test_raises_after_retry_exhaustion` |
| `tests/test_network_egress_client.py` | 40 | Python | `failing_post` |
| `tests/test_network_egress_client.py` | 51 | Python | `test_blocks_disallowed_host` |
| `tests/test_network_egress_client.py` | 57 | Python | `test_blocks_private_address` |
| `tests/test_payment_adapter_resilience.py` | 10 | Python | `test_payment_modules_importable` |
| `tests/test_payment_adapter_resilience.py` | 15 | Python | `test_circuit_opens_at_failure_threshold` |
| `tests/test_payment_adapter_resilience.py` | 27 | Python | `test_send_rejects_invalid_retry_count` |
| `tests/test_payment_adapter_resilience.py` | 32 | Python | `test_send_returns_failed_after_http_5xx` |
| `tests/test_platform_core_auth_context_guard.py` | 4 | Python | `test_platform_core_missing_user_context_returns_unauthorized` |
| `tests/test_profit_blueprint_modules.py` | 17 | Python | `__init__` |
| `tests/test_profit_blueprint_modules.py` | 21 | Python | `fetch_revenue` |
| `tests/test_profit_blueprint_modules.py` | 24 | Python | `fetch_cost` |
| `tests/test_profit_blueprint_modules.py` | 29 | Python | `__init__` |
| `tests/test_profit_blueprint_modules.py` | 33 | Python | `create_campaign` |
| `tests/test_profit_blueprint_modules.py` | 36 | Python | `pause_campaign` |
| `tests/test_profit_blueprint_modules.py` | 39 | Python | `scale_campaign` |
| `tests/test_profit_blueprint_modules.py` | 44 | Python | `snapshot` |
| `tests/test_profit_blueprint_modules.py` | 54 | Python | `test_roi_engine_calculates_ratio_with_decimal_precision` |
| `tests/test_profit_blueprint_modules.py` | 61 | Python | `test_unit_economics_computes_profit_cpa_ltv_and_roas` |
| `tests/test_profit_blueprint_modules.py` | 75 | Python | `test_profit_loop_scales_when_roi_reaches_threshold` |
| `tests/test_profit_blueprint_modules.py` | 83 | Python | `test_profit_loop_pauses_when_roi_is_non_positive` |
| `tests/test_profit_blueprint_modules.py` | 91 | Python | `test_daily_cost_guard_blocks_when_budget_exceeded` |
| `tests/test_profit_blueprint_modules.py` | 100 | Python | `test_daily_cost_guard_rejects_negative_cost` |
| `tests/test_profit_blueprint_modules.py` | 106 | Python | `test_profit_loop_enforces_stop_loss_from_unit_economics` |
| `tests/test_profit_blueprint_modules.py` | 118 | Python | `test_profit_loop_respects_global_kill_switch` |
| `tests/test_profit_blueprint_modules.py` | 126 | Python | `test_spend_control_plane_validates_caps_and_anomaly_and_roi` |
| `tests/test_profit_mode_pipeline.py` | 16 | Python | `load_module` |
| `tests/test_profit_mode_pipeline.py` | 27 | Python | `__init__` |
| `tests/test_profit_mode_pipeline.py` | 30 | Python | `ping` |
| `tests/test_profit_mode_pipeline.py` | 33 | Python | `hgetall` |
| `tests/test_profit_mode_pipeline.py` | 36 | Python | `hset` |
| `tests/test_profit_mode_pipeline.py` | 42 | Python | `test_feature_store_supports_budget_and_revenue_updates` |
| `tests/test_profit_mode_pipeline.py` | 77 | Python | `test_feature_store_can_replace_all_campaign_features` |
| `tests/test_profit_mode_pipeline.py` | 119 | Python | `test_feature_store_rejects_daily_cap_above_budget` |
| `tests/test_profit_mode_pipeline.py` | 140 | Python | `test_affiliate_webhook_updates_feature_store_and_reward_collector` |
| `tests/test_profit_mode_pipeline.py` | 145 | Python | `__enter__` |
| `tests/test_profit_mode_pipeline.py` | 148 | Python | `__exit__` |
| `tests/test_profit_mode_pipeline.py` | 151 | Python | `cursor` |
| `tests/test_profit_mode_pipeline.py` | 154 | Python | `execute` |
| `tests/test_profit_mode_pipeline.py` | 161 | Python | `fake_post` |
| `tests/test_profit_mode_pipeline.py` | 178 | Python | `test_profit_mode_activation_wires_tracking_payment_and_execution` |
| `tests/test_profit_mode_pipeline.py` | 194 | Python | `fake_safe_call` |
| `tests/test_rebrand_zlttbots_to_zlttbots.py` | 16 | Python | `test_audit_detects_legacy_tokens` |
| `tests/test_rebrand_zlttbots_to_zlttbots.py` | 28 | Python | `test_apply_mode_rewrites_tokens` |
| `tests/test_repo_cleanup_audit.py` | 10 | Python | `test_candidate_local_module_targets_resolves_index_and_suffix` |
| `tests/test_repo_cleanup_audit.py` | 21 | Python | `test_traverse_reachable_returns_transitive_nodes` |
| `tests/test_repo_cleanup_audit.py` | 37 | Python | `test_normalize_path_list_ignores_non_existing_entries` |
| `tests/test_safe_edge_pack.py` | 4 | Python | `test_tunnel_config_contains_expected_hosts` |
| `tests/test_safe_edge_pack.py` | 19 | Python | `test_kong_config_exposes_summary_and_redirect_routes` |
| `tests/test_safe_edge_pack.py` | 26 | Python | `test_go_live_doc_has_safety_note` |
| `tests/test_safe_edge_pack.py` | 31 | Python | `test_global_tunnel_config_supports_wildcard_host` |
| `tests/test_safe_edge_pack.py` | 36 | Python | `test_global_deploy_script_includes_multi_region_and_worker_steps` |
| `tests/test_safe_edge_pack.py` | 43 | Python | `test_interactive_domain_env_generator_supports_wildcard_zeaz_domain` |
| `tests/test_script_lifecycle_smoke.py` | 10 | Python | `run_command` |
| `tests/test_script_lifecycle_smoke.py` | 14 | Python | `test_feature_impact_dive_help_smoke` |
| `tests/test_script_lifecycle_smoke.py` | 19 | Python | `test_docs_snippet_validator_smoke` |
| `tests/test_script_lifecycle_smoke.py` | 24 | Python | `test_node_dependency_scan_shell_syntax` |
| `tests/test_singularity_engine.py` | 24 | Python | `test_singularity_orchestration_generates_verify_event` |
| `tests/test_singularity_engine.py` | 58 | Python | `test_risk_score_stable` |
| `tests/test_tracking_and_payment_webhook.py` | 10 | Python | `__init__` |
| `tests/test_tracking_and_payment_webhook.py` | 13 | Python | `execute` |
| `tests/test_tracking_and_payment_webhook.py` | 16 | Python | `__enter__` |
| `tests/test_tracking_and_payment_webhook.py` | 19 | Python | `__exit__` |
| `tests/test_tracking_and_payment_webhook.py` | 24 | Python | `__init__` |
| `tests/test_tracking_and_payment_webhook.py` | 27 | Python | `cursor` |
| `tests/test_tracking_and_payment_webhook.py` | 30 | Python | `commit` |
| `tests/test_tracking_and_payment_webhook.py` | 33 | Python | `__enter__` |
| `tests/test_tracking_and_payment_webhook.py` | 36 | Python | `__exit__` |
| `tests/test_tracking_and_payment_webhook.py` | 40 | Python | `test_tracking_redirects_with_click_id` |
| `tests/test_tracking_and_payment_webhook.py` | 54 | Python | `test_payment_webhook_records_payment` |
| `tests/test_tracking_and_payment_webhook.py` | 58 | Python | `execute` |
| `tests/test_v3_runtime.py` | 26 | Python | `test_api_gateway_uses_service_discovery_round_robin` |
| `tests/test_v3_runtime.py` | 38 | Python | `test_central_queue_publish_consume` |
| `tests/test_v3_runtime.py` | 48 | Python | `test_autoscaling_worker_pool` |
| `tests/test_v3_runtime.py` | 63 | Python | `test_gpu_scheduler_assigns_best_fit` |
| `tests/test_v3_runtime.py` | 73 | Python | `test_distributed_crawler_cluster_reconcile_and_dispatch` |
| `tests/test_v3_runtime.py` | 90 | Python | `test_enterprise_runtime_bootstrap_and_components` |
| `tests/test_v3_runtime.py` | 103 | Python | `test_multi_region_canary_auto_rollback_and_pipeline_optimization` |
| `tests/test_v3_runtime.py` | 130 | Python | `test_policy_analytics_persistence_manifest_healing_and_notifications` |
| `tools/auto_fix_optional_deps.sh` | 6 | Shell | `create_rate_limit_module` |
| `tools/auto_fix_optional_deps.sh` | 43 | Shell | `fix_slowapi_imports` |
| `tools/auto_fix_optional_deps.sh` | 58 | Shell | `fix_uvicorn` |
| `tools/check_imports.py` | 12 | Python | `_check_file` |
| `tools/check_no_runtime_imports.py` | 8 | Python | `is_excluded` |
| `tools/check_no_runtime_imports.py` | 12 | Python | `is_main_guard` |
| `tools/check_no_runtime_imports.py` | 26 | Python | `check_nodes` |
| `tools/check_no_runtime_imports.py` | 46 | Python | `check_file` |
| `tools/check_runtime_dependency_guards.py` | 14 | Python | `_blocked_import` |
| `tools/check_runtime_dependency_guards.py` | 18 | Python | `check_file` |

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

## `zttlbots`

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

## `zTTato-Platform`

Status: **blocked**; commit `blocked`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

None detected or repository blocked.

## `zvath`

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
| `packages/revenue/engine.py` | 41 | Python | `enforce_limit` |
| `packages/revenue/engine.py` | 52 | Python | `dynamic_pricing` |
| `packages/revenue/funnel.py` | 6 | Python | `upsell_decision` |
| `packages/revenue/funnel.py` | 15 | Python | `next_step` |
| `packages/revenue/funnel_v2.py` | 6 | Python | `funnel` |
| `packages/revenue/growth.py` | 6 | Python | `growth_decision` |
| `packages/revenue/ltv.py` | 4 | Python | `calculate_ltv` |
| `packages/revenue/optimizer.py` | 4 | Python | `optimize` |
| `packages/revenue/pricing_ai.py` | 6 | Python | `price_optimize` |
| `packages/revenue/rl_agent.py` | 24 | Python | `_q` |
| `packages/revenue/rl_agent.py` | 28 | Python | `choose_action` |
| `packages/revenue/rl_agent.py` | 40 | Python | `observe` |
| `packages/revenue/rl_agent.py` | 58 | Python | `build_default_agent` |
| `packages/sre/autonomous.py` | 38 | Python | `__init__` |
| `packages/sre/autonomous.py` | 43 | Python | `detect_issues` |
| `packages/sre/autonomous.py` | 61 | Python | `propose_plan` |
| `packages/sre/chaos.py` | 23 | Python | `success_rate` |
| `packages/sre/chaos.py` | 40 | Python | `__init__` |
| `packages/sre/chaos.py` | 43 | Python | `propose_actions` |
| `packages/sre/chaos.py` | 76 | Python | `maybe_fail` |
| `packages/sre/chaos.py` | 82 | Python | `auto_heal` |
| `packages/sre/chaos.py` | 104 | Python | `run_chaos_simulation` |
| `packages/tenant/enforcer.py` | 6 | Python | `enforce` |
| `packages/vector/repo.py` | 8 | Python | `search_similar` |
| `scripts/install.sh` | 12 | Shell | `pick_python` |
| `tests/helpers.py` | 7 | Python | `__init__` |
| `tests/helpers.py` | 11 | Python | `llen` |
| `tests/helpers.py` | 14 | Python | `set` |
| `tests/helpers.py` | 21 | Python | `rpush` |
| `tests/helpers.py` | 25 | Python | `lpop` |
| `tests/helpers.py` | 31 | Python | `setup_oauth_env` |
| `tests/helpers.py` | 40 | Python | `__init__` |
| `tests/helpers.py` | 49 | Python | `json` |
| `tests/test_ai_safety_and_compliance.py` | 14 | Python | `test_ai_red_teaming_report` |
| `tests/test_ai_safety_and_compliance.py` | 29 | Python | `test_adversarial_defense_flags_attack` |
| `tests/test_ai_safety_and_compliance.py` | 38 | Python | `test_model_watermark_round_trip` |
| `tests/test_ai_safety_and_compliance.py` | 46 | Python | `test_regulatory_compliance_frameworks` |
| `tests/test_ai_security_expansions.py` | 13 | Python | `test_ai_security_soc_detects_elevated_threats` |
| `tests/test_ai_security_expansions.py` | 25 | Python | `test_opa_policy_engine_blocks_disallowed_values` |
| `tests/test_ai_security_expansions.py` | 35 | Python | `test_robust_cryptographic_watermark_detects_tampering` |
| `tests/test_ai_security_expansions.py` | 43 | Python | `test_compliance_certification_pipeline_requires_stage_gates` |
| `tests/test_api_security_hardening.py` | 17 | Python | `_build_request` |
| `tests/test_api_security_hardening.py` | 42 | Python | `receive` |
| `tests/test_api_security_hardening.py` | 55 | Python | `__init__` |
| `tests/test_api_security_hardening.py` | 59 | Python | `get` |
| `tests/test_api_security_hardening.py` | 62 | Python | `delete` |
| `tests/test_api_security_hardening.py` | 65 | Python | `execute` |
| `tests/test_api_security_hardening.py` | 73 | Python | `__init__` |
| `tests/test_api_security_hardening.py` | 76 | Python | `set` |
| `tests/test_api_security_hardening.py` | 83 | Python | `getdel` |
| `tests/test_api_security_hardening.py` | 86 | Python | `pipeline` |
| `tests/test_api_security_hardening.py` | 90 | Python | `test_generate_requires_auth_header` |
| `tests/test_api_security_hardening.py` | 96 | Python | `test_generate_enforces_role_and_tenant` |
| `tests/test_api_security_hardening.py` | 108 | Python | `test_oauth_uses_s256_and_validates_state` |
| `tests/test_api_security_hardening.py` | 132 | Python | `test_oauth_callback_rejects_replayed_state` |
| `tests/test_api_security_hardening.py` | 154 | Python | `test_stripe_webhook_requires_secret_in_production` |
| `tests/test_api_security_hardening.py` | 167 | Python | `test_stripe_webhook_requires_signature_in_non_local` |
| `tests/test_api_security_hardening.py` | 180 | Python | `test_auth_context_accepts_valid_bearer` |
| `tests/test_api_security_hardening.py` | 187 | Python | `test_oauth_state_is_shared_across_api_instances` |
| `tests/test_autonomous_sre_and_sovereign_ml.py` | 9 | Python | `test_autonomous_sre_agent_detects_failover_and_prioritizes_tasks` |
| `tests/test_autonomous_sre_and_sovereign_ml.py` | 26 | Python | `test_homomorphic_inference_engine_linear_infer` |
| `tests/test_autonomous_sre_and_sovereign_ml.py` | 35 | Python | `test_homomorphic_addition_preserves_sum` |
| `tests/test_health_logging.py` | 8 | Python | `test_health_degraded_logging_is_rate_limited` |
| `tests/test_health_logging.py` | 12 | Python | `ping` |
| `tests/test_health_logging.py` | 19 | Python | `warning` |
| `tests/test_health_logging.py` | 22 | Python | `info` |
| `tests/test_health_logging.py` | 37 | Python | `test_health_logs_recovery_once` |
| `tests/test_health_logging.py` | 44 | Python | `warning` |
| `tests/test_health_logging.py` | 47 | Python | `info` |
| `tests/test_health_logging.py` | 53 | Python | `ping` |
| `tests/test_input_validation.py` | 14 | Python | `test_compute_ctr_rejects_invalid_inputs` |
| `tests/test_input_validation.py` | 24 | Python | `test_detect_winner_rejects_empty_list` |
| `tests/test_input_validation.py` | 30 | Python | `test_cac_and_ltv_reject_negative_values` |
| `tests/test_input_validation.py` | 43 | Python | `test_optimizer_handles_zero_cac_and_negative_inputs` |
| `tests/test_mlops_stack.py` | 23 | Python | `test_drift_detector_flags_shift` |
| `tests/test_mlops_stack.py` | 30 | Python | `test_model_monitor_emits_alerts` |
| `tests/test_mlops_stack.py` | 37 | Python | `test_shadow_deployment_routes_deterministically` |
| `tests/test_mlops_stack.py` | 45 | Python | `test_pipeline_spec_is_unified` |
| `tests/test_mlops_stack.py` | 53 | Python | `test_safe_rl_controller_clamps_action` |
| `tests/test_mlops_stack.py` | 60 | Python | `test_estimate_uplift` |
| `tests/test_mlops_stack.py` | 65 | Python | `test_lineage_payloads` |
| `tests/test_mlops_stack.py` | 75 | Python | `test_privacy_layer_detects_and_redacts` |
| `tests/test_mlops_stack.py` | 86 | Python | `test_dp_sgd_clips_and_noises` |
| `tests/test_mlops_stack.py` | 94 | Python | `test_federated_aggregation_weighted_average` |
| `tests/test_mlops_stack.py` | 111 | Python | `test_explainability_shap_and_lime` |
| `tests/test_mlops_stack.py` | 122 | Python | `test_compliance_summary_risk_level` |
| `tests/test_production_upgrade.py` | 12 | Python | `test_crypto_requires_strong_key` |
| `tests/test_production_upgrade.py` | 20 | Python | `test_queue_enqueue_is_idempotent` |
| `tests/test_production_upgrade.py` | 31 | Python | `test_queue_sends_invalid_json_to_dlq` |
| `tests/test_production_upgrade.py` | 36 | Python | `stop_after_first_sleep` |
| `tests/test_production_upgrade.py` | 49 | Python | `test_oauth_refresh_normalizes_payload` |
| `tests/test_quantum_governance_federation.py` | 20 | Python | `test_quantum_safe_encrypt_decrypt_round_trip` |
| `tests/test_quantum_governance_federation.py` | 33 | Python | `test_quantum_safe_rejects_tampered_payload` |
| `tests/test_quantum_governance_federation.py` | 46 | Python | `test_ai_governance_board_enforces_critical_threshold_and_signoff` |
| `tests/test_quantum_governance_federation.py` | 75 | Python | `test_global_ai_federation_network_filters_by_policy` |
| `tests/test_revenue_optimizer.py` | 8 | Python | `test_calculate_ltv_handles_zero_churn` |
| `tests/test_revenue_optimizer.py` | 13 | Python | `test_calculate_cac_handles_zero_users` |
| `tests/test_revenue_optimizer.py` | 18 | Python | `test_optimize_returns_expected_strategy` |
| `tests/test_rl_and_recommendations.py` | 9 | Python | `test_rl_agent_updates_q_and_prefers_best_action` |
| `tests/test_rl_and_recommendations.py` | 29 | Python | `test_vector_clustering_and_recommendations` |
| `tests/test_service_smoke.py` | 11 | Python | `__init__` |
| `tests/test_service_smoke.py` | 15 | Python | `__iter__` |
| `tests/test_service_smoke.py` | 18 | Python | `commit` |
| `tests/test_service_smoke.py` | 23 | Python | `__init__` |
| `tests/test_service_smoke.py` | 27 | Python | `send` |
| `tests/test_service_smoke.py` | 30 | Python | `flush` |
| `tests/test_service_smoke.py` | 38 | Python | `test_kafka_worker_smoke_routes_failed_job_to_dlq` |
| `tests/test_service_smoke.py` | 53 | Python | `test_redis_worker_smoke_sends_retry_exhausted_job_to_dlq` |
| `tests/test_service_smoke.py` | 58 | Python | `stop_after_second_idle` |
| `tests/test_ultra_expansion.py` | 8 | Python | `test_retention_rate` |
| `tests/test_ultra_expansion.py` | 14 | Python | `test_price_optimize` |
| `tests/test_ultra_expansion.py` | 19 | Python | `test_growth_decision_paths` |
| `tests/test_ultra_extensions.py` | 21 | Python | `__init__` |
| `tests/test_ultra_extensions.py` | 25 | Python | `fetch_one` |
| `tests/test_ultra_extensions.py` | 30 | Python | `test_crypto_encrypt_decrypt_and_rotate` |
| `tests/test_ultra_extensions.py` | 44 | Python | `test_crypto_rejects_tampered_ciphertext` |
| `tests/test_ultra_extensions.py` | 58 | Python | `test_crypto_decrypt_supports_legacy_format` |
| `tests/test_ultra_extensions.py` | 70 | Python | `test_budget_manager_and_guarded_call` |
| `tests/test_ultra_extensions.py` | 82 | Python | `test_circuit_breaker_opens_after_threshold` |
| `tests/test_ultra_extensions.py` | 86 | Python | `fail` |
| `tests/test_ultra_extensions.py` | 97 | Python | `test_chaos_injection_paths` |
| `tests/test_ultra_extensions.py` | 107 | Python | `test_auto_heal_recovers_transient_failures` |
| `tests/test_ultra_extensions.py` | 110 | Python | `flaky` |
| `tests/test_ultra_extensions.py` | 120 | Python | `test_run_chaos_simulation_tracks_recovery` |
| `tests/test_ultra_modules.py` | 8 | Python | `test_viral_score_increases_with_keywords` |
| `tests/test_ultra_modules.py` | 15 | Python | `test_compute_ctr_zero_division_safe` |
| `tests/test_ultra_modules.py` | 20 | Python | `test_detect_winner_returns_max_score` |
| `tests/test_ultra_modules.py` | 26 | Python | `test_upsell_decision_paths` |
| `tests/test_ultra_system.py` | 9 | Python | `test_extract_trend_patterns_filters_by_viral_keyword` |
| `tests/test_ultra_system.py` | 15 | Python | `test_funnel_v2_paths` |
| `tests/test_worker_resilience.py` | 14 | Python | `test_competitor_scraper_uses_https_proxy` |
| `tests/test_worker_resilience.py` | 24 | Python | `raise_for_status` |
| `tests/test_worker_resilience.py` | 27 | Python | `fake_get` |
| `tests/test_worker_resilience.py` | 44 | Python | `test_competitor_scraper_allows_empty_proxy_list` |
| `tests/test_worker_resilience.py` | 54 | Python | `raise_for_status` |
| `tests/test_worker_resilience.py` | 57 | Python | `fake_get` |
| `tests/test_worker_resilience.py` | 71 | Python | `test_stealth_jitter_handles_reversed_bounds` |
| `tests/test_worker_resilience.py` | 89 | Python | `test_queue_dependency_failure_retries_then_dlq` |
| `tests/test_worker_resilience.py` | 97 | Python | `labels` |
| `tests/test_worker_resilience.py` | 101 | Python | `inc` |
| `tests/test_worker_resilience.py` | 104 | Python | `fail_with_timeout` |
| `tests/test_worker_resilience.py` | 107 | Python | `stop_after_idle` |
| `tests/test_zero_trust_confidential_and_healing.py` | 11 | Python | `test_spiffe_parsing_and_workload_identity_allowlist` |
| `tests/test_zero_trust_confidential_and_healing.py` | 21 | Python | `test_confidential_compute_policy_evaluates_attestation` |
| `tests/test_zero_trust_confidential_and_healing.py` | 44 | Python | `test_ai_threat_intel_feed_deduplicates_and_ranks` |
| `tests/test_zero_trust_confidential_and_healing.py` | 60 | Python | `test_autonomous_healer_generates_remediation_actions` |

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

## `tiktok-shop-bot`

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

## `tiktokshop-api-client`

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

## `tiktok-shop-sdk`

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
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4500 | JS/TS | `hydrateOnIdle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4504 | JS/TS | `elementIsVisibleInViewport` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4509 | JS/TS | `hydrateOnVisible` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4529 | JS/TS | `hydrateOnMediaQuery` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4540 | JS/TS | `hydrateOnInteraction` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4543 | JS/TS | `doHydrate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4551 | JS/TS | `teardown` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4565 | JS/TS | `forEachElement` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4588 | JS/TS | `isAsyncWrapper` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4589 | JS/TS | `defineAsyncComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4607 | JS/TS | `retry` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4612 | JS/TS | `load` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4618 | JS/TS | `userRetry` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4619 | JS/TS | `userFail` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4647 | JS/TS | `__asyncHydrate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4650 | JS/TS | `performHydrate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4679 | JS/TS | `setup` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4685 | JS/TS | `onError` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4749 | JS/TS | `createInnerComp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4757 | JS/TS | `isKeepAlive` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4769 | JS/TS | `setup` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4844 | JS/TS | `unmount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4848 | JS/TS | `pruneCache` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4856 | JS/TS | `pruneCacheEntry` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4876 | JS/TS | `cacheSubtree` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4965 | JS/TS | `matches` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4976 | JS/TS | `onActivated` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4979 | JS/TS | `onDeactivated` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 4982 | JS/TS | `registerKeepAliveHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5004 | JS/TS | `injectToKeepAliveRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5016 | JS/TS | `resetShapeFlag` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5020 | JS/TS | `getInnerChild` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5023 | JS/TS | `injectHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5047 | JS/TS | `createHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5067 | JS/TS | `onErrorCaptured` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5072 | JS/TS | `resolveComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5076 | JS/TS | `resolveDynamicComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5083 | JS/TS | `resolveDirective` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5086 | JS/TS | `resolveAsset` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5120 | JS/TS | `resolve` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5123 | JS/TS | `renderList` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5175 | JS/TS | `createSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5192 | JS/TS | `renderSlot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5235 | JS/TS | `ensureValidVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5244 | JS/TS | `toHandlers` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5255 | JS/TS | `getPublicInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5283 | JS/TS | `isReservedPrefix` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5284 | JS/TS | `hasSetupBinding` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5286 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5364 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5403 | JS/TS | `defineProperty` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5421 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5427 | JS/TS | `has` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5439 | JS/TS | `createDevRenderContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5458 | JS/TS | `exposePropsOnRenderContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5474 | JS/TS | `exposeSetupStateOnRenderContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5495 | JS/TS | `warnRuntimeUsage` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5498 | JS/TS | `defineProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5504 | JS/TS | `defineEmits` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5510 | JS/TS | `defineExpose` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5515 | JS/TS | `defineOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5520 | JS/TS | `defineSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5526 | JS/TS | `defineModel` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5531 | JS/TS | `withDefaults` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5537 | JS/TS | `useSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5540 | JS/TS | `useAttrs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5543 | JS/TS | `getContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5550 | JS/TS | `normalizePropsOrEmits` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5556 | JS/TS | `mergeDefaults` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5578 | JS/TS | `mergeModels` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5583 | JS/TS | `createPropsRestProxy` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5595 | JS/TS | `withAsyncContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5612 | JS/TS | `createDuplicateChecker` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5623 | JS/TS | `applyOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5773 | JS/TS | `registerLifecycleHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5818 | JS/TS | `resolveInjections` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5853 | JS/TS | `callHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5860 | JS/TS | `createWatcher` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5890 | JS/TS | `resolveMergedOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5920 | JS/TS | `mergeOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5973 | JS/TS | `mergeDataFn` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5987 | JS/TS | `mergeInject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 5990 | JS/TS | `normalizeInject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6000 | JS/TS | `mergeAsArray` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6003 | JS/TS | `mergeObjectOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6006 | JS/TS | `mergeEmitsOrPropsOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6020 | JS/TS | `mergeWatchOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6029 | JS/TS | `createAppContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6051 | JS/TS | `createAppAPI` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6082 | JS/TS | `use` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6098 | JS/TS | `mixin` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6112 | JS/TS | `component` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6125 | JS/TS | `directive` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6138 | JS/TS | `mount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6180 | JS/TS | `onUnmount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6188 | JS/TS | `unmount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6205 | JS/TS | `provide` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6220 | JS/TS | `runWithContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6234 | JS/TS | `provide` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6249 | JS/TS | `inject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6264 | JS/TS | `hasInjectionContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6268 | JS/TS | `useSSRContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6279 | JS/TS | `watchEffect` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6282 | JS/TS | `watchPostEffect` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6289 | JS/TS | `watchSyncEffect` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6296 | JS/TS | `watch2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6304 | JS/TS | `doWatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6332 | JS/TS | `watchStopHandle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6379 | JS/TS | `instanceWatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6394 | JS/TS | `createPathGetter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6404 | JS/TS | `useModel` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6429 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6433 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6456 | JS/TS | `next` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6467 | JS/TS | `getModelModifiers` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6470 | JS/TS | `emit` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6556 | JS/TS | `normalizeEmitsOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6566 | JS/TS | `extendEmits` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6599 | JS/TS | `isEmitListener` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6607 | JS/TS | `markAttrsAccessed` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6610 | JS/TS | `renderComponentRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6638 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6756 | JS/TS | `getChildRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6767 | JS/TS | `setRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6779 | JS/TS | `filterSingleRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6800 | JS/TS | `getFunctionalFallthrough` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6809 | JS/TS | `filterModelListeners` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6818 | JS/TS | `isElementRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6821 | JS/TS | `shouldUpdateComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6868 | JS/TS | `hasPropsChanged` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6881 | JS/TS | `updateHOCHostEl` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6896 | JS/TS | `createInternalObject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6897 | JS/TS | `isInternalObject` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6898 | JS/TS | `initProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6922 | JS/TS | `isInHmrContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 6928 | JS/TS | `updateProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7020 | JS/TS | `setFullProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7062 | JS/TS | `resolvePropValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7104 | JS/TS | `normalizePropsOptions` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7115 | JS/TS | `extendProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7193 | JS/TS | `validatePropName` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7201 | JS/TS | `getType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7213 | JS/TS | `validateProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7229 | JS/TS | `validateProp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7259 | JS/TS | `assertType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7282 | JS/TS | `getInvalidTypeMessage` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7300 | JS/TS | `styleValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7309 | JS/TS | `isExplicable` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7313 | JS/TS | `isBoolean` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7316 | JS/TS | `isInternalKey` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7317 | JS/TS | `normalizeSlotValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7318 | JS/TS | `normalizeSlot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7333 | JS/TS | `normalizeObjectSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7351 | JS/TS | `normalizeVNodeSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7360 | JS/TS | `assignSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7367 | JS/TS | `initSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7383 | JS/TS | `updateSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7417 | JS/TS | `startMeasure` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7425 | JS/TS | `endMeasure` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7440 | JS/TS | `isSupported` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7452 | JS/TS | `initFeatureFlags` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7476 | JS/TS | `createRenderer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7479 | JS/TS | `createHydrationRenderer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7482 | JS/TS | `baseCreateRenderer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7505 | JS/TS | `patch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7607 | JS/TS | `processText` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7621 | JS/TS | `processCommentNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7632 | JS/TS | `mountStaticNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7642 | JS/TS | `patchStaticNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7657 | JS/TS | `moveStaticNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7666 | JS/TS | `removeStaticNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7675 | JS/TS | `processElement` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7714 | JS/TS | `mountElement` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7775 | JS/TS | `setScopeId` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7801 | JS/TS | `mountChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7817 | JS/TS | `patchElement` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7908 | JS/TS | `patchBlockChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7939 | JS/TS | `patchProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 7968 | JS/TS | `processFragment` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8044 | JS/TS | `processComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8070 | JS/TS | `mountComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8119 | JS/TS | `updateComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8140 | JS/TS | `setupRenderEffect` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8141 | JS/TS | `componentUpdateFn` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8156 | JS/TS | `hydrateSubTree` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8332 | JS/TS | `updateComponentPreRender` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8343 | JS/TS | `patchChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8420 | JS/TS | `patchUnkeyedChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8464 | JS/TS | `patchKeyedChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8630 | JS/TS | `move` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8664 | JS/TS | `remove22` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8671 | JS/TS | `performLeave` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8693 | JS/TS | `unmount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8772 | JS/TS | `remove2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8792 | JS/TS | `performRemove` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8800 | JS/TS | `performLeave` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8810 | JS/TS | `removeFragment` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8819 | JS/TS | `unmountComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8844 | JS/TS | `unmountChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8849 | JS/TS | `getNextHostNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8861 | JS/TS | `render2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8910 | JS/TS | `resolveChildrenNamespace` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8913 | JS/TS | `toggleRecurse` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8922 | JS/TS | `needTransition` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8925 | JS/TS | `traverseStaticChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8953 | JS/TS | `getSequence` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 8993 | JS/TS | `locateNonHydratedAsyncRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9003 | JS/TS | `invalidateMount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9009 | JS/TS | `isSuspense` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9018 | JS/TS | `process` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9055 | JS/TS | `triggerEvent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9061 | JS/TS | `mountSuspense` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9108 | JS/TS | `patchSuspense` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9269 | JS/TS | `createSuspenseBoundary` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9312 | JS/TS | `resolve` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9396 | JS/TS | `fallback` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9403 | JS/TS | `mountFallback` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9438 | JS/TS | `move` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9442 | JS/TS | `next` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9445 | JS/TS | `registerDep` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9494 | JS/TS | `unmount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9516 | JS/TS | `hydrateSuspense` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9544 | JS/TS | `normalizeSuspenseChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9552 | JS/TS | `normalizeSuspenseSlot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9580 | JS/TS | `queueEffectWithSuspense` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9591 | JS/TS | `setActiveBranch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9605 | JS/TS | `isVNodeSuspensible` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9615 | JS/TS | `openBlock` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9618 | JS/TS | `closeBlock` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9623 | JS/TS | `setBlockTracking` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9629 | JS/TS | `setupBlock` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9637 | JS/TS | `createElementBlock` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9650 | JS/TS | `createBlock` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9662 | JS/TS | `isVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9665 | JS/TS | `isSameVNodeType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9677 | JS/TS | `transformVNodeArgs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9680 | JS/TS | `createVNodeWithArgsTransform` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9685 | JS/TS | `normalizeKey` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9696 | JS/TS | `createBaseVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9751 | JS/TS | `_createVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9815 | JS/TS | `guardReactiveProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9819 | JS/TS | `cloneVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9874 | JS/TS | `deepCloneVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9881 | JS/TS | `createTextVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9884 | JS/TS | `createStaticVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9889 | JS/TS | `createCommentVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9892 | JS/TS | `normalizeVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9908 | JS/TS | `cloneIfMounted` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9911 | JS/TS | `normalizeChildren` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9956 | JS/TS | `mergeProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9980 | JS/TS | `invokeVNodeHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 9988 | JS/TS | `createComponentInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10080 | JS/TS | `getCurrentInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10085 | JS/TS | `registerGlobalSetter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10103 | JS/TS | `setCurrentInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10112 | JS/TS | `unsetCurrentInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10117 | JS/TS | `validateComponentName` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10124 | JS/TS | `isStatefulComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10128 | JS/TS | `setupComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10138 | JS/TS | `setupStatefulComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10211 | JS/TS | `handleSetupResult` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10240 | JS/TS | `registerRuntimeCompiler` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10248 | JS/TS | `isRuntimeOnly` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10249 | JS/TS | `finishComponentSetup` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10302 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10307 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10311 | JS/TS | `deleteProperty` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10316 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10321 | JS/TS | `getSlotsProxy` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10323 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10329 | JS/TS | `createSetupContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10330 | JS/TS | `expose` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10377 | JS/TS | `getComponentPublicInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10380 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10387 | JS/TS | `has` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10396 | JS/TS | `classify` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10397 | JS/TS | `getComponentName` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10400 | JS/TS | `formatComponentName` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10409 | JS/TS | `inferFromRegistry` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10422 | JS/TS | `isClassComponent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10425 | JS/TS | `computed2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10435 | JS/TS | `h` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10460 | JS/TS | `initCustomFormatter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10470 | JS/TS | `header` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10509 | JS/TS | `hasBody` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10512 | JS/TS | `body` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10522 | JS/TS | `formatInstance` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10555 | JS/TS | `createInstanceBlock` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10586 | JS/TS | `formatValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10599 | JS/TS | `extractKeys` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10612 | JS/TS | `isKeyOfType` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10624 | JS/TS | `genRefFlag` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10639 | JS/TS | `withMemo` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10649 | JS/TS | `isMemoSame` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10731 | JS/TS | `setScopeId` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10738 | JS/TS | `insertStaticContent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10793 | JS/TS | `decorate$1` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10801 | JS/TS | `callHook2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10808 | JS/TS | `hasExplicitCallback` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10811 | JS/TS | `resolveTransitionProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10848 | JS/TS | `finishEnter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10854 | JS/TS | `finishLeave` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10861 | JS/TS | `makeEnterHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10864 | JS/TS | `resolve2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10876 | JS/TS | `onBeforeEnter` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10881 | JS/TS | `onBeforeAppear` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10888 | JS/TS | `onLeave` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10890 | JS/TS | `resolve2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10911 | JS/TS | `onEnterCancelled` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10915 | JS/TS | `onAppearCancelled` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10919 | JS/TS | `onLeaveCancelled` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10925 | JS/TS | `normalizeDuration` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10935 | JS/TS | `NumberOf` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10942 | JS/TS | `addTransitionClass` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10946 | JS/TS | `removeTransitionClass` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10956 | JS/TS | `nextFrame` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10962 | JS/TS | `whenTransitionEnds` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10964 | JS/TS | `resolveIfNotStale` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10978 | JS/TS | `end` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10982 | JS/TS | `onEnd` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10994 | JS/TS | `getTransitionInfo` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 10996 | JS/TS | `getStyleProperties` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11033 | JS/TS | `getTimeout` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11039 | JS/TS | `toMs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11043 | JS/TS | `forceReflow` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11047 | JS/TS | `patchClass` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11065 | JS/TS | `beforeMount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11073 | JS/TS | `mounted` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11078 | JS/TS | `updated` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11094 | JS/TS | `beforeUnmount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11098 | JS/TS | `setDisplay` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11102 | JS/TS | `initVShowForSSR` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11110 | JS/TS | `useCssVars` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11124 | JS/TS | `setVars` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11143 | JS/TS | `setVarsOnVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11169 | JS/TS | `setVarsOnNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11182 | JS/TS | `patchStyle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11232 | JS/TS | `setStyle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11262 | JS/TS | `autoPrefix` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11281 | JS/TS | `patchAttr` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11299 | JS/TS | `patchDOMProp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11349 | JS/TS | `addEventListener` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11352 | JS/TS | `removeEventListener` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11356 | JS/TS | `patchEvent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11376 | JS/TS | `parseName` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11391 | JS/TS | `getNow` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11392 | JS/TS | `createInvoker` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11393 | JS/TS | `invoker` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11410 | JS/TS | `sanitizeEventValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11420 | JS/TS | `patchStopImmediatePropagation` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11434 | JS/TS | `isNativeOn` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11436 | JS/TS | `patchProp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11465 | JS/TS | `shouldSetAsProp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11502 | JS/TS | `defineCustomElement` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11506 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11513 | JS/TS | `defineSSRCustomElement` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11519 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11555 | JS/TS | `connectedCallback` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11583 | JS/TS | `_setParent` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11589 | JS/TS | `_inheritParentContext` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11597 | JS/TS | `disconnectedCallback` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11615 | JS/TS | `_processMutations` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11623 | JS/TS | `_resolveDef` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11632 | JS/TS | `resolve2` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11669 | JS/TS | `_mount` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11693 | JS/TS | `_resolveProps` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11703 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11706 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11712 | JS/TS | `_setAttr` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11725 | JS/TS | `_getProp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11731 | JS/TS | `_setProp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11762 | JS/TS | `_update` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11767 | JS/TS | `_createVNode` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11789 | JS/TS | `dispatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11808 | JS/TS | `_applyStyles` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11841 | JS/TS | `_parseSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11853 | JS/TS | `_renderSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11883 | JS/TS | `_getSlots` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11900 | JS/TS | `_injectChildStyle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11906 | JS/TS | `_beginPatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11913 | JS/TS | `_endPatch` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11922 | JS/TS | `_removeChildStyle` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11935 | JS/TS | `useHost` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11953 | JS/TS | `useShadowRoot` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11957 | JS/TS | `useCssModule` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11981 | JS/TS | `decorate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 11991 | JS/TS | `setup` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12075 | JS/TS | `callPendingCbs` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12084 | JS/TS | `recordPosition` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12090 | JS/TS | `applyTranslation` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12102 | JS/TS | `hasCSSTransform` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12118 | JS/TS | `getModelAssigner` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12122 | JS/TS | `onCompositionStart` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12125 | JS/TS | `onCompositionEnd` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12133 | JS/TS | `castValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12139 | JS/TS | `created` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12158 | JS/TS | `mounted` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12161 | JS/TS | `beforeUpdate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12183 | JS/TS | `created` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12215 | JS/TS | `beforeUpdate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12220 | JS/TS | `setChecked` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12236 | JS/TS | `created` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12243 | JS/TS | `beforeUpdate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12253 | JS/TS | `created` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12271 | JS/TS | `mounted` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12274 | JS/TS | `beforeUpdate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12277 | JS/TS | `updated` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12283 | JS/TS | `setSelected` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12315 | JS/TS | `getValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12318 | JS/TS | `getCheckboxValue` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12323 | JS/TS | `created` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12326 | JS/TS | `mounted` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12329 | JS/TS | `beforeUpdate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12332 | JS/TS | `updated` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12336 | JS/TS | `resolveDynamicModel` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12353 | JS/TS | `callModelHook` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12361 | JS/TS | `initVModelForSSR` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12409 | JS/TS | `withModifiers` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12429 | JS/TS | `withKeys` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12447 | JS/TS | `ensureRenderer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12450 | JS/TS | `ensureHydrationRenderer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12455 | JS/TS | `render` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12458 | JS/TS | `hydrate` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12461 | JS/TS | `createApp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12487 | JS/TS | `createSSRApp` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12502 | JS/TS | `resolveRootNamespace` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12510 | JS/TS | `injectNativeTagCheck` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12516 | JS/TS | `injectCompilerOptionsCheck` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12520 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12523 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12535 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12539 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12545 | JS/TS | `normalizeContainer` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12563 | JS/TS | `initDirectivesForSSR` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12572 | JS/TS | `initDev` |
| `apps/docs/.vitepress/cache/deps/chunk-KK2TUMQT.js` | 12580 | JS/TS | `compile2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 8 | JS/TS | `__esm` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 11 | JS/TS | `__commonJS` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 14 | JS/TS | `__copyProps` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 22 | JS/TS | `__toESM` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 40 | JS/TS | `copyBuffer` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 46 | JS/TS | `rfdc2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 60 | JS/TS | `cloneArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 78 | JS/TS | `clone` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 100 | JS/TS | `cloneProto` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 122 | JS/TS | `rfdcCircles` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 136 | JS/TS | `cloneArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 159 | JS/TS | `clone` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 190 | JS/TS | `cloneProto` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 236 | JS/TS | `toUpper` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 239 | JS/TS | `classify` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 242 | JS/TS | `basename` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 256 | JS/TS | `isUrlString` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 265 | JS/TS | `debounce` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 275 | JS/TS | `applyFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 314 | JS/TS | `_applyPromised` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 319 | JS/TS | `flatHooks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 332 | JS/TS | `_createTask` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 334 | JS/TS | `serialTaskCaller` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 342 | JS/TS | `parallelTaskCaller` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 347 | JS/TS | `callEachWith` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 353 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 363 | JS/TS | `hook` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 405 | JS/TS | `hookOnce` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 407 | JS/TS | `_function` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 418 | JS/TS | `removeHook` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 429 | JS/TS | `deprecateHook` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 437 | JS/TS | `deprecateHooks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 443 | JS/TS | `addHooks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 454 | JS/TS | `removeHooks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 460 | JS/TS | `removeAllHooks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 465 | JS/TS | `callHook` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 469 | JS/TS | `callHookParallel` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 473 | JS/TS | `callHookWith` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 494 | JS/TS | `beforeEach` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 506 | JS/TS | `afterEach` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 519 | JS/TS | `createHooks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 534 | JS/TS | `__esm2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 537 | JS/TS | `__commonJS2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 540 | JS/TS | `__copyProps2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 548 | JS/TS | `__toESM2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2071 | JS/TS | `define` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2103 | JS/TS | `getComponentTypeName` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2111 | JS/TS | `getComponentFileName` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2116 | JS/TS | `saveComponentGussedName` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2120 | JS/TS | `getAppRecord` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2126 | JS/TS | `isFragment` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2135 | JS/TS | `getInstanceName` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2155 | JS/TS | `getUniqueComponentId` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2161 | JS/TS | `getComponentInstance` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2166 | JS/TS | `createRect` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2182 | JS/TS | `getTextRect` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2188 | JS/TS | `getFragmentRect` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2209 | JS/TS | `mergeRects` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2228 | JS/TS | `getComponentBoundingRect` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2243 | JS/TS | `getRootElementsFromComponentInstance` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2250 | JS/TS | `getFragmentRootElements` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2297 | JS/TS | `getContainerElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2300 | JS/TS | `getCardElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2303 | JS/TS | `getIndicatorElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2306 | JS/TS | `getNameElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2309 | JS/TS | `getStyles` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2317 | JS/TS | `create` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2345 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2362 | JS/TS | `highlight` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2370 | JS/TS | `unhighlight` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2376 | JS/TS | `inspectFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2392 | JS/TS | `selectComponentFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2401 | JS/TS | `cancelInspectComponentHighLighter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2407 | JS/TS | `inspectComponentHighLighter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2410 | JS/TS | `onSelect` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2427 | JS/TS | `scrollToComponent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2469 | JS/TS | `waitForInspectorInit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2482 | JS/TS | `setupInspector` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2490 | JS/TS | `getComponentInspector` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2492 | JS/TS | `setup` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2507 | JS/TS | `isReadonly` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2513 | JS/TS | `isReactive` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2525 | JS/TS | `isRef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2528 | JS/TS | `toRaw` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2537 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2540 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2566 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2580 | JS/TS | `has` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2593 | JS/TS | `createDefaultSetCallback` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2619 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2646 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2649 | JS/TS | `isRef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2658 | JS/TS | `getTimelineLayersStateFromStorage` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2686 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2690 | JS/TS | `addTimelineLayer` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2702 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2709 | JS/TS | `addInspector` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2722 | JS/TS | `getActiveInspectors` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2738 | JS/TS | `getInspector` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2741 | JS/TS | `DevToolsV6PluginAPIHookKeys` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2753 | JS/TS | `DevToolsContextHookKeys` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2767 | JS/TS | `DevToolsMessagingHookKeys` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2778 | JS/TS | `createDevToolsCtxHooks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2920 | JS/TS | `initStateFactory` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2945 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2952 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2960 | JS/TS | `updateAllStates` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2969 | JS/TS | `setActiveAppRecord` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2973 | JS/TS | `setActiveAppRecordId` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2978 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2990 | JS/TS | `deleteProperty` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 2994 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3001 | JS/TS | `onDevToolsConnected` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3015 | JS/TS | `resolveIcon` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3025 | JS/TS | `addCustomTab` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3035 | JS/TS | `addCustomCommand` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3049 | JS/TS | `removeCustomCommand` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3057 | JS/TS | `openInEditor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3085 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3089 | JS/TS | `_getSettings` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3096 | JS/TS | `getPluginLocalKey` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3099 | JS/TS | `getPluginSettingsOptions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3101 | JS/TS | `item` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3107 | JS/TS | `getPluginSettings` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3117 | JS/TS | `item` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3122 | JS/TS | `initPluginSettings` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3129 | JS/TS | `setPluginSettings` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3167 | JS/TS | `vueAppInit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3170 | JS/TS | `vueAppUnmount` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3173 | JS/TS | `vueAppConnected` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3176 | JS/TS | `componentAdded` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3179 | JS/TS | `componentEmit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3182 | JS/TS | `componentUpdated` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3185 | JS/TS | `componentRemoved` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3188 | JS/TS | `setupDevtoolsPlugin` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3191 | JS/TS | `perfStart` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3194 | JS/TS | `perfEnd` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3200 | JS/TS | `setupDevToolsPlugin` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3205 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3245 | JS/TS | `notifyComponentUpdate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3270 | JS/TS | `addInspector` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3276 | JS/TS | `sendInspectorTree` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3282 | JS/TS | `sendInspectorState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3288 | JS/TS | `selectInspectorNode` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3291 | JS/TS | `visitComponentTree` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3295 | JS/TS | `now` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3301 | JS/TS | `addTimelineLayer` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3304 | JS/TS | `addTimelineEvent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3311 | JS/TS | `getSettings` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3315 | JS/TS | `getComponentInstances` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3318 | JS/TS | `getComponentBounds` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3321 | JS/TS | `getComponentName` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3324 | JS/TS | `highlightElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3328 | JS/TS | `unhighlightElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3364 | JS/TS | `setupDevToolsPlugin` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3367 | JS/TS | `callDevToolsPluginSetupFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3385 | JS/TS | `registerDevToolsPlugin` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3411 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3416 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3422 | JS/TS | `getRoutes` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3426 | JS/TS | `filterRoutes` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3439 | JS/TS | `filterCurrentRoute` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3455 | JS/TS | `normalizeRouterInfo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3456 | JS/TS | `init` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3482 | JS/TS | `createDevToolsApi` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3485 | JS/TS | `getInspectorTree` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3504 | JS/TS | `getInspectorState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3526 | JS/TS | `editInspectorState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3544 | JS/TS | `sendInspectorState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3552 | JS/TS | `inspectComponentInspector` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3556 | JS/TS | `cancelInspectComponentInspector` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3560 | JS/TS | `getComponentRenderCode` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3566 | JS/TS | `scrollToComponent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3574 | JS/TS | `toggleApp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3585 | JS/TS | `inspectDOM` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3594 | JS/TS | `updatePluginSettings` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3597 | JS/TS | `getPluginSettings` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3636 | JS/TS | `onDevToolsClientConnected` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3652 | JS/TS | `toggleHighPerfMode` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3661 | JS/TS | `updateDevToolsClientDetected` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3680 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3684 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3688 | JS/TS | `getByKey` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3691 | JS/TS | `getByValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3694 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3700 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3704 | JS/TS | `register` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3713 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3716 | JS/TS | `getIdentifier` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3719 | JS/TS | `getValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3724 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3728 | JS/TS | `register` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3738 | JS/TS | `getAllowedProps` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3744 | JS/TS | `valuesOfObj` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3756 | JS/TS | `find` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3770 | JS/TS | `forEach` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3773 | JS/TS | `includes` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3776 | JS/TS | `findArr` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3786 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3789 | JS/TS | `register` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3792 | JS/TS | `findApplicable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3795 | JS/TS | `findByName` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3801 | JS/TS | `getType` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3802 | JS/TS | `isUndefined` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3803 | JS/TS | `isNull` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3804 | JS/TS | `isPlainObject2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3813 | JS/TS | `isEmptyObject` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3814 | JS/TS | `isArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3815 | JS/TS | `isString` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3816 | JS/TS | `isNumber` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3817 | JS/TS | `isBoolean` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3818 | JS/TS | `isRegExp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3819 | JS/TS | `isMap` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3820 | JS/TS | `isSet` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3821 | JS/TS | `isSymbol` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3822 | JS/TS | `isDate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3823 | JS/TS | `isError` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3824 | JS/TS | `isNaNValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3825 | JS/TS | `isPrimitive2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3826 | JS/TS | `isBigint` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3827 | JS/TS | `isInfinite` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3828 | JS/TS | `isTypedArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3829 | JS/TS | `isURL` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3831 | JS/TS | `escapeKey` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3832 | JS/TS | `stringifyPath` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3833 | JS/TS | `parsePath` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3857 | JS/TS | `simpleTransformation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3922 | JS/TS | `compositeTransformation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 3967 | JS/TS | `isInstanceOfRegisteredClass` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4010 | JS/TS | `transformValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4031 | JS/TS | `untransformValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4054 | JS/TS | `getNthKey` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4064 | JS/TS | `validatePath` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4075 | JS/TS | `getDeep` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4099 | JS/TS | `setDeep` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4168 | JS/TS | `traverse` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4184 | JS/TS | `applyValueAnnotations` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4190 | JS/TS | `applyReferentialEqualityAnnotations` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4191 | JS/TS | `apply` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4210 | JS/TS | `isDeep` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4211 | JS/TS | `addIdentity` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4219 | JS/TS | `generateReferentialEqualityAnnotations` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4308 | JS/TS | `getType2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4311 | JS/TS | `isArray2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4314 | JS/TS | `isPlainObject3` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4320 | JS/TS | `isNull2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4323 | JS/TS | `isOneOf` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4326 | JS/TS | `isUndefined2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4330 | JS/TS | `assignProp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4343 | JS/TS | `copy` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4366 | JS/TS | `constructor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4376 | JS/TS | `serialize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4397 | JS/TS | `deserialize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4408 | JS/TS | `stringify` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4411 | JS/TS | `parse` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4414 | JS/TS | `registerClass` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4417 | JS/TS | `registerSymbol` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4420 | JS/TS | `registerCustom` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vue_devtools-api.js` | 4426 | JS/TS | `allowErrorProps` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 41 | JS/TS | `computedEager` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 52 | JS/TS | `computedWithControl` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 57 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 68 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 76 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 85 | JS/TS | `tryOnScopeDispose` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 92 | JS/TS | `createEventHook` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 94 | JS/TS | `off` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 97 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 100 | JS/TS | `on` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 102 | JS/TS | `offFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 108 | JS/TS | `trigger` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 118 | JS/TS | `createGlobalState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 131 | JS/TS | `injectLocal` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 141 | JS/TS | `provideLocal` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 152 | JS/TS | `createInjectionState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 155 | JS/TS | `useProvidingState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 160 | JS/TS | `useInjectedState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 163 | JS/TS | `createRef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 170 | JS/TS | `createSharedComposable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 174 | JS/TS | `dispose` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 192 | JS/TS | `extendRef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 198 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 201 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 212 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 217 | JS/TS | `isDefined` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 220 | JS/TS | `makeDestructurable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 225 | JS/TS | `value` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 240 | JS/TS | `reactify` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 246 | JS/TS | `reactifyObject` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 268 | JS/TS | `toReactive` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 272 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 275 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 282 | JS/TS | `deleteProperty` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 285 | JS/TS | `has` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 288 | JS/TS | `ownKeys` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 291 | JS/TS | `getOwnPropertyDescriptor` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 300 | JS/TS | `reactiveComputed` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 303 | JS/TS | `reactiveOmit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 310 | JS/TS | `isDef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 311 | JS/TS | `notNullish` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 312 | JS/TS | `assert` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 317 | JS/TS | `isObject` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 318 | JS/TS | `now` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 319 | JS/TS | `timestamp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 320 | JS/TS | `clamp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 321 | JS/TS | `noop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 323 | JS/TS | `rand` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 328 | JS/TS | `hasOwn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 330 | JS/TS | `getIsIOS` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 334 | JS/TS | `createFilterWrapper` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 335 | JS/TS | `wrapper` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 342 | JS/TS | `bypassFilter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 345 | JS/TS | `debounceFilter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 349 | JS/TS | `_clearTimeout` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 355 | JS/TS | `filter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 388 | JS/TS | `throttleFilter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 402 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 410 | JS/TS | `filter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 413 | JS/TS | `invoke2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 442 | JS/TS | `pausableFilter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 447 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 450 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 453 | JS/TS | `eventFilter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 459 | JS/TS | `cacheStringFunction` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 472 | JS/TS | `promiseTimeout` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 480 | JS/TS | `identity` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 483 | JS/TS | `createSingletonPromise` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 485 | JS/TS | `wrapper` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 498 | JS/TS | `invoke` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 501 | JS/TS | `containsProp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 504 | JS/TS | `increaseWithUnit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 515 | JS/TS | `pxValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 518 | JS/TS | `objectPick` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 527 | JS/TS | `objectOmit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 532 | JS/TS | `objectEntries` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 535 | JS/TS | `getLifeCycleTarget` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 538 | JS/TS | `toArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 541 | JS/TS | `toRef2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 548 | JS/TS | `reactivePick` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 553 | JS/TS | `refAutoReset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 557 | JS/TS | `resetAfter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 565 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 569 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 578 | JS/TS | `useDebounceFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 584 | JS/TS | `refDebounced` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 592 | JS/TS | `refDefault` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 594 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 598 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 603 | JS/TS | `useThrottleFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 609 | JS/TS | `refThrottled` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 619 | JS/TS | `refWithControl` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 627 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 630 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 635 | JS/TS | `get2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 640 | JS/TS | `set2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 652 | JS/TS | `untrackedGet` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 653 | JS/TS | `silentSet` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 654 | JS/TS | `peek` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 655 | JS/TS | `lay` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 670 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 680 | JS/TS | `watchWithFilter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 694 | JS/TS | `watchPausable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 711 | JS/TS | `syncRef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 744 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 749 | JS/TS | `syncRefs` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 762 | JS/TS | `toRefs2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 768 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 771 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 794 | JS/TS | `tryOnBeforeMount` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 803 | JS/TS | `tryOnBeforeUnmount` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 808 | JS/TS | `tryOnMounted` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 817 | JS/TS | `tryOnUnmounted` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 822 | JS/TS | `createUntil` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 823 | JS/TS | `toMatch` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 852 | JS/TS | `toBe` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 887 | JS/TS | `toBeTruthy` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 890 | JS/TS | `toBeNull` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 893 | JS/TS | `toBeUndefined` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 896 | JS/TS | `toBeNaN` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 899 | JS/TS | `toContains` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 905 | JS/TS | `changed` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 908 | JS/TS | `changedTimes` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 943 | JS/TS | `until` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 946 | JS/TS | `defaultComparator` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 949 | JS/TS | `useArrayDifference` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 969 | JS/TS | `useArrayEvery` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 972 | JS/TS | `useArrayFilter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 975 | JS/TS | `useArrayFind` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 980 | JS/TS | `useArrayFindIndex` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 983 | JS/TS | `findLast` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 991 | JS/TS | `useArrayFindLast` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 996 | JS/TS | `isArrayIncludesOptions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 999 | JS/TS | `useArrayIncludes` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1021 | JS/TS | `useArrayJoin` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1024 | JS/TS | `useArrayMap` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1027 | JS/TS | `useArrayReduce` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1028 | JS/TS | `reduceCallback` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1034 | JS/TS | `useArraySome` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1037 | JS/TS | `uniq` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1040 | JS/TS | `uniqueElementsBy` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1047 | JS/TS | `useArrayUnique` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1053 | JS/TS | `useCounter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1060 | JS/TS | `inc` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1061 | JS/TS | `dec` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1062 | JS/TS | `get2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1063 | JS/TS | `set2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1064 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1072 | JS/TS | `defaultMeridiem` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1078 | JS/TS | `formatOrdinal` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1083 | JS/TS | `formatDate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1094 | JS/TS | `stripTimeZone` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1141 | JS/TS | `normalizeDate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1158 | JS/TS | `useDateFormat` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1161 | JS/TS | `useIntervalFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1168 | JS/TS | `clean` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1174 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1178 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1205 | JS/TS | `useInterval` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1212 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1213 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1234 | JS/TS | `useLastChanged` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1244 | JS/TS | `useTimeoutFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1251 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1257 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1261 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1284 | JS/TS | `useTimeout` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1304 | JS/TS | `useToNumber` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1321 | JS/TS | `useToString` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1324 | JS/TS | `useToggle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1331 | JS/TS | `toggle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1346 | JS/TS | `watchArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1368 | JS/TS | `watchAtMost` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1386 | JS/TS | `watchDebounced` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1401 | JS/TS | `watchDeep` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1411 | JS/TS | `watchIgnorable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1481 | JS/TS | `watchImmediate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1491 | JS/TS | `watchOnce` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1498 | JS/TS | `watchThrottled` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1514 | JS/TS | `watchTriggerable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1516 | JS/TS | `onEffect` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1523 | JS/TS | `onCleanup` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1526 | JS/TS | `_cb` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1532 | JS/TS | `trigger` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1544 | JS/TS | `getWatchSources` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1551 | JS/TS | `getOldValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1554 | JS/TS | `whenever` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1573 | JS/TS | `computedAsync` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1630 | JS/TS | `computedInject` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1645 | JS/TS | `createReusableTemplate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1651 | JS/TS | `setup` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1660 | JS/TS | `setup` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1678 | JS/TS | `keysToCamelKebabCase` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1684 | JS/TS | `createTemplatePromise` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1687 | JS/TS | `create` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1714 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1720 | JS/TS | `renderList` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1731 | JS/TS | `createUnrefFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1740 | JS/TS | `unrefElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1745 | JS/TS | `useEventListener` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1747 | JS/TS | `cleanup` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1751 | JS/TS | `register` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1785 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1793 | JS/TS | `onClickOutside` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1805 | JS/TS | `shouldIgnore` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1815 | JS/TS | `hasMultipleRoots` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1819 | JS/TS | `checkMultipleRoots` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1826 | JS/TS | `listener` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1867 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1883 | JS/TS | `useMounted` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1893 | JS/TS | `useSupported` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1900 | JS/TS | `useMutationObserver` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1904 | JS/TS | `cleanup` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1926 | JS/TS | `takeRecords` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1929 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1940 | JS/TS | `onElementRemoval` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1949 | JS/TS | `cleanupAndUpdate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1973 | JS/TS | `stopHandle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1980 | JS/TS | `createKeyPredicate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 1989 | JS/TS | `onKeyStroke` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2017 | JS/TS | `listener` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2025 | JS/TS | `onKeyDown` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2028 | JS/TS | `onKeyPressed` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2031 | JS/TS | `onKeyUp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2036 | JS/TS | `onLongPress` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2043 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2052 | JS/TS | `onRelease` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2069 | JS/TS | `onDown` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2091 | JS/TS | `onMove` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2116 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2119 | JS/TS | `isFocusedElementEditable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2132 | JS/TS | `isTypedCharValid` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2146 | JS/TS | `onStartTyping` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2148 | JS/TS | `keydown` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2156 | JS/TS | `templateRef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2158 | JS/TS | `_trigger` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2163 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2168 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2176 | JS/TS | `useActiveElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2184 | JS/TS | `getDeepActiveElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2194 | JS/TS | `trigger` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2225 | JS/TS | `useRafFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2238 | JS/TS | `loop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2257 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2264 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2280 | JS/TS | `useAnimate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2316 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2319 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2326 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2329 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2338 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2341 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2348 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2351 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2357 | JS/TS | `play` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2370 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2379 | JS/TS | `reverse` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2391 | JS/TS | `finish` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2400 | JS/TS | `cancel` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2431 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2465 | JS/TS | `syncResume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2469 | JS/TS | `syncPause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2492 | JS/TS | `useAsyncQueue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2515 | JS/TS | `updateResult` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2555 | JS/TS | `whenAborted` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2564 | JS/TS | `useAsyncState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2578 | JS/TS | `execute` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2612 | JS/TS | `waitUntilIsLoaded` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2619 | JS/TS | `then` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2631 | JS/TS | `getDefaultSerialization` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2643 | JS/TS | `useBase64` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2646 | JS/TS | `execute` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2699 | JS/TS | `imgLoaded` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2711 | JS/TS | `blobToBase64` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2721 | JS/TS | `useBattery` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2730 | JS/TS | `updateBatteryInfo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2751 | JS/TS | `useBluetooth` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2766 | JS/TS | `requestDevice` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2784 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2789 | JS/TS | `connectToBluetoothGATTServer` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2824 | JS/TS | `useSSRWidth` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2828 | JS/TS | `provideSSRWidth` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2835 | JS/TS | `useMediaQuery` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2841 | JS/TS | `handler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2951 | JS/TS | `useBreakpoints` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2952 | JS/TS | `getValue2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2966 | JS/TS | `match` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2974 | JS/TS | `greaterOrEqual` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2977 | JS/TS | `smallerOrEqual` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2988 | JS/TS | `current` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2995 | JS/TS | `greater` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 2998 | JS/TS | `smaller` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3001 | JS/TS | `between` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3004 | JS/TS | `isGreater` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3007 | JS/TS | `isGreaterOrEqual` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3010 | JS/TS | `isSmaller` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3013 | JS/TS | `isSmallerOrEqual` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3016 | JS/TS | `isInBetween` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3020 | JS/TS | `active` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3026 | JS/TS | `useBroadcastChannel` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3036 | JS/TS | `post` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3040 | JS/TS | `close` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3086 | JS/TS | `useBrowserLocation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3098 | JS/TS | `buildState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3120 | JS/TS | `useCached` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3129 | JS/TS | `usePermission` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3138 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3169 | JS/TS | `useClipboard` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3184 | JS/TS | `updateText` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3199 | JS/TS | `copy` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3216 | JS/TS | `legacyCopy` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3226 | JS/TS | `legacyRead` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3230 | JS/TS | `isAllowed` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3240 | JS/TS | `useClipboardItems` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3251 | JS/TS | `updateContent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3260 | JS/TS | `copy` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3275 | JS/TS | `cloneFnJSON` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3278 | JS/TS | `useCloned` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3299 | JS/TS | `sync` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3318 | JS/TS | `getHandlers` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3323 | JS/TS | `getSSRHandler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3326 | JS/TS | `setSSRHandler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3329 | JS/TS | `usePreferredDark` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3332 | JS/TS | `guessSerializerType` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3370 | JS/TS | `useStorage` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3421 | JS/TS | `dispatchWriteEvent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3434 | JS/TS | `write` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3451 | JS/TS | `read` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3470 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3492 | JS/TS | `updateFromCustomEvent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3498 | JS/TS | `useColorMode` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3564 | JS/TS | `defaultOnChanged` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3568 | JS/TS | `onChanged` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3577 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3580 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3586 | JS/TS | `useConfirmDialog` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3591 | JS/TS | `reveal` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3598 | JS/TS | `confirm` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3603 | JS/TS | `cancel` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3618 | JS/TS | `useCountdown` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3631 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3635 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3639 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3646 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3660 | JS/TS | `useCssVar` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3667 | JS/TS | `updateCssVar` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3706 | JS/TS | `useCurrentElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3716 | JS/TS | `useCycleList` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3720 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3728 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3732 | JS/TS | `set2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3740 | JS/TS | `shift` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3743 | JS/TS | `next` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3746 | JS/TS | `prev` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3749 | JS/TS | `getInitialValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3762 | JS/TS | `useDark` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3783 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3786 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3796 | JS/TS | `fnBypass` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3799 | JS/TS | `fnSetSource` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3802 | JS/TS | `defaultDump` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3805 | JS/TS | `defaultParse` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3808 | JS/TS | `useManualRefHistory` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3815 | JS/TS | `_createHistoryRecord` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3824 | JS/TS | `_setSource` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3828 | JS/TS | `commit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3836 | JS/TS | `clear` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3840 | JS/TS | `undo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3847 | JS/TS | `redo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3854 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3875 | JS/TS | `useRefHistory` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3896 | JS/TS | `setSource` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3904 | JS/TS | `commit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3908 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3913 | JS/TS | `batch` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3915 | JS/TS | `cancel` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3922 | JS/TS | `dispose` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3936 | JS/TS | `useDebouncedRefHistory` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3943 | JS/TS | `useDeviceMotion` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3960 | JS/TS | `init` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 3987 | JS/TS | `ensurePermissions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4023 | JS/TS | `useDeviceOrientation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4046 | JS/TS | `useDevicePixelRatio` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4061 | JS/TS | `useDevicesList` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4075 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4085 | JS/TS | `ensurePermissions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4124 | JS/TS | `useDisplayMedia` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4136 | JS/TS | `_start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4144 | JS/TS | `_stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4149 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4153 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4177 | JS/TS | `useDocumentVisibility` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4187 | JS/TS | `useDraggable` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4208 | JS/TS | `filterEvent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4213 | JS/TS | `handleEvent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4219 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4239 | JS/TS | `move` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4264 | JS/TS | `end` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4274 | JS/TS | `config` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4294 | JS/TS | `useDropZone` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4304 | JS/TS | `getFiles` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4309 | JS/TS | `checkDataTypes` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4321 | JS/TS | `checkValidity` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4327 | JS/TS | `isSafari` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4328 | JS/TS | `handleDragEvent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4381 | JS/TS | `useResizeObserver` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4385 | JS/TS | `cleanup` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4409 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4419 | JS/TS | `useElementBounding` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4435 | JS/TS | `recalculate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4460 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4491 | JS/TS | `useElementByPoint` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4506 | JS/TS | `cb` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4517 | JS/TS | `useElementHover` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4526 | JS/TS | `toggle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4549 | JS/TS | `useElementSize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4595 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4605 | JS/TS | `useIntersectionObserver` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4644 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4653 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4657 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4663 | JS/TS | `useElementVisibility` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4700 | JS/TS | `useEventBus` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4702 | JS/TS | `on` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4707 | JS/TS | `_off` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4711 | JS/TS | `once` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4712 | JS/TS | `_listener` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4718 | JS/TS | `off` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4726 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4729 | JS/TS | `emit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4735 | JS/TS | `resolveNestedOptions$1` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4740 | JS/TS | `useEventSource` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4756 | JS/TS | `close` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4764 | JS/TS | `_init` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4805 | JS/TS | `open` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4829 | JS/TS | `useEyeDropper` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4833 | JS/TS | `open` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4843 | JS/TS | `useFavicon` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4850 | JS/TS | `applyIcon` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4878 | JS/TS | `isFetchOptions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4882 | JS/TS | `isAbsoluteURL` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4885 | JS/TS | `headersToObject` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4890 | JS/TS | `combineCallbacks` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4914 | JS/TS | `createFetch` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4918 | JS/TS | `useFactoryFetch` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 4959 | JS/TS | `useFetch` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5002 | JS/TS | `abort` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5013 | JS/TS | `loading` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5020 | JS/TS | `execute` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5156 | JS/TS | `setMethod` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5174 | JS/TS | `then` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5182 | JS/TS | `waitUntilFinished` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5187 | JS/TS | `setType` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5193 | JS/TS | `then` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5205 | JS/TS | `then` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5210 | JS/TS | `joinPaths` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5225 | JS/TS | `prepareInitialFiles` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5236 | JS/TS | `useFileDialog` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5256 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5263 | JS/TS | `open` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5288 | JS/TS | `useFileSystemAccess` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5314 | JS/TS | `open` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5321 | JS/TS | `create` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5328 | JS/TS | `save` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5340 | JS/TS | `saveAs` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5351 | JS/TS | `updateFile` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5355 | JS/TS | `updateData` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5382 | JS/TS | `useFocus` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5395 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5415 | JS/TS | `useFocusWithin` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5432 | JS/TS | `useFps` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5459 | JS/TS | `useFullscreen` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5506 | JS/TS | `isCurrentElementFullScreen` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5511 | JS/TS | `isElementFullScreen` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5524 | JS/TS | `exit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5538 | JS/TS | `enter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5549 | JS/TS | `toggle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5550 | JS/TS | `await` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5552 | JS/TS | `handlerCallback` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5570 | JS/TS | `mapGamepadToXbox360Controller` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5613 | JS/TS | `useGamepad` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5621 | JS/TS | `stateFromGamepad` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5640 | JS/TS | `updateGamepadState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5648 | JS/TS | `onGamepadConnected` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5655 | JS/TS | `onGamepadDisconnected` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5680 | JS/TS | `useGeolocation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5700 | JS/TS | `updatePosition` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5706 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5721 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5739 | JS/TS | `useIdle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5750 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5781 | JS/TS | `loadImage` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5814 | JS/TS | `useImage` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5830 | JS/TS | `resolveElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5838 | JS/TS | `useScroll` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5863 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5866 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5871 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5874 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5878 | JS/TS | `scrollTo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5909 | JS/TS | `onScrollEnd` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5920 | JS/TS | `setArrivedState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5956 | JS/TS | `onScrollHandler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 5994 | JS/TS | `measure` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6001 | JS/TS | `useInfiniteScroll` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6024 | JS/TS | `checkAndLoad` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6050 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6056 | JS/TS | `useKeyModifier` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6073 | JS/TS | `useLocalStorage` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6087 | JS/TS | `useMagicKeys` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6097 | JS/TS | `toJSON` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6105 | JS/TS | `setRefs` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6113 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6118 | JS/TS | `updateRefs` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6156 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6177 | JS/TS | `usingElRef` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6181 | JS/TS | `timeRangeToArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6187 | JS/TS | `tracksToArray` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6194 | JS/TS | `useMediaControls` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6221 | JS/TS | `disableTrack` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6233 | JS/TS | `enableTrack` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6242 | JS/TS | `togglePictureInPicture` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6488 | JS/TS | `useMemoize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6489 | JS/TS | `initCache` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6495 | JS/TS | `generateKey` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6496 | JS/TS | `_loadData` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6500 | JS/TS | `loadData` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6501 | JS/TS | `deleteData` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6504 | JS/TS | `clearData` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6507 | JS/TS | `memoized` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6520 | JS/TS | `useMemory` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6537 | JS/TS | `useMouse` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6555 | JS/TS | `mouseHandler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6567 | JS/TS | `touchHandler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6576 | JS/TS | `scrollHandler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6585 | JS/TS | `reset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6609 | JS/TS | `useMouseInElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6624 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6674 | JS/TS | `useMousePressed` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6690 | JS/TS | `onPressed` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6696 | JS/TS | `onReleased` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6722 | JS/TS | `useNavigatorLanguage` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6736 | JS/TS | `useNetwork` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6750 | JS/TS | `updateNetworkInformation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6792 | JS/TS | `useNow` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6798 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6809 | JS/TS | `useObjectUrl` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6811 | JS/TS | `release` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6828 | JS/TS | `useClamp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6833 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6836 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6841 | JS/TS | `useOffsetPagination` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6868 | JS/TS | `prev` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6871 | JS/TS | `next` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6894 | JS/TS | `useOnline` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6898 | JS/TS | `usePageLeave` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6901 | JS/TS | `handler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6916 | JS/TS | `useScreenOrientation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6930 | JS/TS | `lockOrientation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6935 | JS/TS | `unlockOrientation` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 6947 | JS/TS | `useParallax` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7021 | JS/TS | `useParentElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7023 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7032 | JS/TS | `usePerformanceObserver` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7040 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7043 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7072 | JS/TS | `usePointer` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7079 | JS/TS | `handler` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7095 | JS/TS | `usePointerLock` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7121 | JS/TS | `lock` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7132 | JS/TS | `unlock` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7147 | JS/TS | `usePointerSwipe` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7157 | JS/TS | `updatePosStart` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7162 | JS/TS | `updatePosEnd` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7181 | JS/TS | `eventIsAllowed` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7230 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7241 | JS/TS | `usePreferredColorScheme` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7252 | JS/TS | `usePreferredContrast` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7266 | JS/TS | `usePreferredLanguages` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7277 | JS/TS | `usePreferredReducedMotion` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7285 | JS/TS | `usePreferredReducedTransparency` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7293 | JS/TS | `usePrevious` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7308 | JS/TS | `useScreenSafeArea` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7325 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7339 | JS/TS | `getValue` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7342 | JS/TS | `useScriptTag` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7357 | JS/TS | `loadScript` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7358 | JS/TS | `resolveWithElement` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7402 | JS/TS | `load` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7407 | JS/TS | `unload` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7423 | JS/TS | `checkOverflowScroll` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7434 | JS/TS | `preventDefault` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7446 | JS/TS | `useScrollLock` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7466 | JS/TS | `lock` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7483 | JS/TS | `unlock` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7495 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7498 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7505 | JS/TS | `useSessionStorage` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7509 | JS/TS | `useShare` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7513 | JS/TS | `share` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7531 | JS/TS | `defaultSortFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7532 | JS/TS | `defaultCompare` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7533 | JS/TS | `useSorted` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7564 | JS/TS | `useSpeechRecognition` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7577 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7580 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7583 | JS/TS | `toggle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7644 | JS/TS | `useSpeechSynthesis` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7658 | JS/TS | `toggle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7661 | JS/TS | `bindEventsForUtterance` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7694 | JS/TS | `speak` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7699 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7735 | JS/TS | `useStepper` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7744 | JS/TS | `at` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7749 | JS/TS | `get2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7754 | JS/TS | `goTo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7758 | JS/TS | `goToNext` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7763 | JS/TS | `goToPrevious` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7768 | JS/TS | `goBackTo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7772 | JS/TS | `isNext` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7775 | JS/TS | `isPrevious` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7778 | JS/TS | `isCurrent` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7781 | JS/TS | `isBefore` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7784 | JS/TS | `isAfter` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7809 | JS/TS | `useStorageAsync` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7838 | JS/TS | `read` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7887 | JS/TS | `useStyleTag` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7896 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7898 | JS/TS | `load` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7919 | JS/TS | `unload` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7938 | JS/TS | `useSwipe` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7962 | JS/TS | `getTouchEventCoords` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7963 | JS/TS | `updateCoordsStart` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7967 | JS/TS | `updateCoordsEnd` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 7972 | JS/TS | `onTouchEnd` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8000 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8013 | JS/TS | `useTemplateRefsList` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8024 | JS/TS | `useTextDirection` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8031 | JS/TS | `getValue2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8045 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8048 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8060 | JS/TS | `getRangesFromSelection` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8065 | JS/TS | `useTextSelection` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8076 | JS/TS | `onSelectionChange` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8090 | JS/TS | `tryRequestAnimationFrame` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8097 | JS/TS | `useTextareaAutosize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8105 | JS/TS | `triggerResize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8140 | JS/TS | `useThrottledRefHistory` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8170 | JS/TS | `DEFAULT_FORMATTER` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8173 | JS/TS | `useTimeAgo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8189 | JS/TS | `formatTimeAgo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8202 | JS/TS | `getValue2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8205 | JS/TS | `format` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8211 | JS/TS | `applyFormat` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8222 | JS/TS | `unitMax` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8235 | JS/TS | `useTimeoutPoll` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8242 | JS/TS | `loop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8248 | JS/TS | `resume` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8256 | JS/TS | `pause` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8268 | JS/TS | `useTimestamp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8277 | JS/TS | `update` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8292 | JS/TS | `useTitle` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8301 | JS/TS | `format` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8361 | JS/TS | `createEasingFunction` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8362 | JS/TS | `a` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8363 | JS/TS | `b` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8364 | JS/TS | `c` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8365 | JS/TS | `calcBezier` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8366 | JS/TS | `getSlope` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8367 | JS/TS | `getTforX` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8380 | JS/TS | `lerp` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8383 | JS/TS | `toVec` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8386 | JS/TS | `executeTransition` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8399 | JS/TS | `tick` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8425 | JS/TS | `useTransition` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8427 | JS/TS | `sourceVal` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8463 | JS/TS | `useUrlSearchParams` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8475 | JS/TS | `getRawParams` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8486 | JS/TS | `constructQuery` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8498 | JS/TS | `read` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8501 | JS/TS | `updateState` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8529 | JS/TS | `write` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8548 | JS/TS | `onChanged` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8564 | JS/TS | `useUserMedia` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8575 | JS/TS | `getDeviceOptions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8589 | JS/TS | `_start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8598 | JS/TS | `_stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8603 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8607 | JS/TS | `start` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8613 | JS/TS | `restart` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8648 | JS/TS | `useVModel` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8665 | JS/TS | `cloneFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8666 | JS/TS | `getValue2` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8667 | JS/TS | `triggerEmit` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8700 | JS/TS | `get` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8703 | JS/TS | `set` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8709 | JS/TS | `useVModels` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8721 | JS/TS | `useVibrate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8730 | JS/TS | `vibrate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8734 | JS/TS | `stop` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8757 | JS/TS | `useVirtualList` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8772 | JS/TS | `useVirtualListResources` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8780 | JS/TS | `createGetViewCapacity` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8797 | JS/TS | `createGetOffset` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8814 | JS/TS | `createCalculateRange` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8833 | JS/TS | `createGetDistance` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8843 | JS/TS | `useWatchForSizes` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8848 | JS/TS | `createComputedTotalSize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8859 | JS/TS | `createScrollTo` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8867 | JS/TS | `useHorizontalVirtualList` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8899 | JS/TS | `useVerticalVirtualList` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8930 | JS/TS | `useWakeLock` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8953 | JS/TS | `forceRequest` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8958 | JS/TS | `request` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8964 | JS/TS | `release` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 8979 | JS/TS | `useWebNotification` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9003 | JS/TS | `ensurePermissions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9017 | JS/TS | `show` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9028 | JS/TS | `close` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9059 | JS/TS | `resolveNestedOptions` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9064 | JS/TS | `useWebSocket` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9086 | JS/TS | `_sendBuffer` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9093 | JS/TS | `resetRetry` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9099 | JS/TS | `resetHeartbeat` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9103 | JS/TS | `close` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9113 | JS/TS | `send` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9123 | JS/TS | `_init` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9200 | JS/TS | `open` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9221 | JS/TS | `useWebWorker` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9227 | JS/TS | `post` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9259 | JS/TS | `depsParser` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9275 | JS/TS | `jobRunner` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9285 | JS/TS | `createWorkerBlobUrl` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9291 | JS/TS | `useWebWorkerFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9302 | JS/TS | `workerTerminate` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9314 | JS/TS | `generateWorker` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9349 | JS/TS | `callWorker` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9358 | JS/TS | `workerFn` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9374 | JS/TS | `useWindowFocus` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9388 | JS/TS | `useWindowScroll` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9392 | JS/TS | `useWindowSize` |
| `apps/docs/.vitepress/cache/deps/vitepress___@vueuse_core.js` | 9403 | JS/TS | `update` |
| `apps/docs/.vitepress/theme/index.ts` | 14 | JS/TS | `enhanceApp` |
| `apps/examples/src/affiliate_partner/createConversationwithCreator.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/editAffiliatePartnerCampaign.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/generateAffiliatePartnerCampaignProductLink.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/getAffiliateCampaignCreatorFulfillmentStatusInfo.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/getAffiliateCampaignCreatorFulfillmentStatusList.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/getAffiliateCampaignCreatorProductContentStatistics.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/getAffiliateCampaignCreatorProductSampleStatus.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/getAffiliatePartnerCampaignDetail.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/getAffiliatePartnerCampaignList.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/getAffiliatePartnerCampaignProductList.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/partnerGenerateMultiAffiliateCampaignProductLink.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/publishAffiliatePartnerCampaign.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/reviewAffiliatePartnerCampaign.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/searchCAPAffiliateOrders.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_partner/searchTapAffiliateOrders.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/affiliate_seller/createConversationwithCreator.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopPerformance.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopProductPerformance.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopProductPerformanceList.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopSKUPerformance.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopSKUPerformanceList.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopVideoPerformanceDetails.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopVideoPerformanceList.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopVideoPerformanceOverview.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/analytics/getShopVideoProductPerformanceList.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/authentication/generateToken.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/authentication/refreshToken.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/cli.ts` | 11 | JS/TS | `getAllModules` |
| `apps/examples/src/cli.ts` | 26 | JS/TS | `main` |
| `apps/examples/src/event/deleteShopWebhook.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/event/getShopWebhooks.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/event/updateShopWebhook.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/finances/getPayments.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/finances/getStatements.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/finances/getTransactionsByOrder.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/finances/getTransactionsByStatement.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/finances/getWithdrawals.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/combinablePackage.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/createFirstMileBundle.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/createPackage.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/getEligibleShippingService.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/getOrderSplitAttributes.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/getPackageHandoverTimeSlots.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/searchCombinablePackages.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/searchPackage.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/splitOrders.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/fulfillment/uncombinePackages.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/logistic/getGlobalSellerWarehouse.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/logistic/getShippingProviders.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/logistic/getWarehouseDeliveryOptions.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/logistic/getWarehouseList.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/orders/addExternalOrderReferences.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/orders/getExternalOrderReferences.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/orders/getOrderDetail.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/orders/getOrderList.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/orders/getPriceDetail.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/orders/searchOrderByExternalOrderReference.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/SearchResponsiblePersons.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/activateProducts.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/checkProductListing.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/createCategoryUpgradeTask.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/createCustomBrands.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/createGlobalProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/createImageTranslationTasks.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/createManufacturer.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/createProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/createResponsiblePerson.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/deactivateProducts.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/deleteGlobalProducts.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/deleteProducts.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/editGlobalProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/editPartialManufacturer.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/editProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/editResponsiblePersons.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getAttributes.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getBrands.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getCategories.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getCategoryRules.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getGlobalAttributes.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getGlobalCategories.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getGlobalCategoryRules.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getGlobalProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getImageTranslationTasks.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getProductPrerequisites.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getProductsSEOWords.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/getRecommendedProductTitleAndDescription.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/optimizedImages.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/productInformationIssueDiagnosis.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/publishGlobalProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/recommendCategory.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/recommendGlobalCategory.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/recoverProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/searchGlobalProducts.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/searchInventory.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/searchManufacturer.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/searchProducts.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/searchSizeCharts.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/updateGlobalInventory.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/updatePrice.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/updateProductInventory.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/product/uploadProductFile.ts` | 12 | JS/TS | `main` |
| `apps/examples/src/product/uploadProductImage.ts` | 12 | JS/TS | `main` |
| `apps/examples/src/promotion/createActivity.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/deactivateActivity.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/getActivity.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/getCoupon.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/removeActivityProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/searchActivity.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/searchCoupon.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/updateActivity.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/promotion/updateActivityProduct.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/approveCancellation.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/approveReturn.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/calculateCancellation.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/cancelOrder.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/createReturn.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/getAftersaleEligibility.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/getRejectReasons.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/getReturnRecord.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/rejectCancellation.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/rejectReturn.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/searchCancellation.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/return_refund/searchReturn.ts` | 9 | JS/TS | `main` |
| `apps/examples/src/seller/getActiveShops.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/seller/getSellerPermissions.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/shop/getAuthorizedShops.ts` | 10 | JS/TS | `main` |
| `apps/examples/src/shop/getCategoryAssets.ts` | 10 | JS/TS | `main` |
| `apps/landing/scripts/generate-i18n-types.ts` | 14 | JS/TS | `flattenKeys` |
| `apps/landing/scripts/generate-i18n-types.ts` | 25 | JS/TS | `generateTypes` |
| `apps/landing/scripts/generate-i18n-types.ts` | 79 | JS/TS | `capitalize` |
| `apps/landing/src/components/Navbar.tsx` | 19 | JS/TS | `handleScroll` |
| `apps/landing/src/components/Navbar.tsx` | 29 | JS/TS | `switchLang` |
| `apps/landing/src/components/contributors/ContributorCard.tsx` | 5 | JS/TS | `ContributorCard` |
| `apps/landing/src/components/contributors/ContributorSkeleton.tsx` | 1 | JS/TS | `ContributorSkeleton` |
| `apps/landing/src/components/pages/QuickLinksSection.tsx` | 21 | JS/TS | `copyToClipboard` |
| `apps/landing/src/constants/footerSections.ts` | 16 | JS/TS | `getFooterSections` |
| `apps/landing/src/context/theme/ThemeContext.tsx` | 14 | JS/TS | `useTheme` |
| `apps/landing/src/context/theme/ThemeProvider.tsx` | 4 | JS/TS | `ThemeProvider` |
| `apps/landing/src/context/theme/ThemeProvider.tsx` | 9 | JS/TS | `initTheme` |
| `apps/landing/src/context/theme/ThemeProvider.tsx` | 25 | JS/TS | `toggleTheme` |
| `apps/landing/src/hooks/useI18n.ts` | 32 | JS/TS | `useI18n` |
| `apps/landing/src/hooks/useI18n.ts` | 38 | JS/TS | `useI18n` |
| `apps/landing/src/pages/_app.tsx` | 6 | JS/TS | `App` |
| `apps/landing/src/pages/index.tsx` | 32 | JS/TS | `getStaticProps` |
| `apps/landing/src/services/contributors.service.ts` | 3 | JS/TS | `getContributors` |
| `packages/sdk/src/client/RequestMultipart.ts` | 19 | JS/TS | `requestMultipart` |
| `packages/sdk/src/client/__tests__/Request.test.ts` | 14 | JS/TS | `constructor` |
| `packages/sdk/src/modules/AffiliatePartnerModule.ts` | 44 | JS/TS | `constructor` |
| `packages/sdk/src/modules/AffiliateSellerModule.ts` | 70 | JS/TS | `constructor` |
| `packages/sdk/src/modules/AffiliateSellerModule.ts` | 357 | JS/TS | `getOpenCollaborationSettings` |
| `packages/sdk/src/modules/AffiliateSellerModule.ts` | 556 | JS/TS | `getLatestUnreadMessages` |
| `packages/sdk/src/modules/AnalyticsModule.ts` | 34 | JS/TS | `constructor` |
| `packages/sdk/src/modules/AuthModule.ts` | 15 | JS/TS | `constructor` |
| `packages/sdk/src/modules/EventModule.ts` | 18 | JS/TS | `constructor` |
| `packages/sdk/src/modules/EventModule.ts` | 26 | JS/TS | `getShopWebhooks` |
| `packages/sdk/src/modules/FinanceModule.ts` | 23 | JS/TS | `constructor` |
| `packages/sdk/src/modules/FulfillmentModule.ts` | 30 | JS/TS | `constructor` |
| `packages/sdk/src/modules/LogisticModule.ts` | 22 | JS/TS | `constructor` |
| `packages/sdk/src/modules/LogisticModule.ts` | 29 | JS/TS | `getWarehouseList` |
| `packages/sdk/src/modules/LogisticModule.ts` | 41 | JS/TS | `getGlobalSellerWarehouse` |
| `packages/sdk/src/modules/OrderModule.ts` | 22 | JS/TS | `constructor` |
| `packages/sdk/src/modules/ProductModule.ts` | 110 | JS/TS | `getProductPrerequisites` |
| `packages/sdk/src/modules/ProductModule.ts` | 284 | JS/TS | `createCategoryUpgradeTask` |
| `packages/sdk/src/modules/ProductModule.ts` | 427 | JS/TS | `deleteProducts` |
| `packages/sdk/src/modules/PromotionModule.ts` | 31 | JS/TS | `constructor` |
| `packages/sdk/src/modules/PromotionModule.ts` | 156 | JS/TS | `getCoupon` |
| `packages/sdk/src/modules/ReturnRefundModule.ts` | 42 | JS/TS | `constructor` |
| `packages/sdk/src/modules/ReturnRefundModule.ts` | 126 | JS/TS | `rejectReturn` |
| `packages/sdk/src/modules/SellerModule.ts` | 17 | JS/TS | `constructor` |
| `packages/sdk/src/modules/SellerModule.ts` | 28 | JS/TS | `getActiveShops` |
| `packages/sdk/src/modules/SellerModule.ts` | 46 | JS/TS | `getSellerPermissions` |
| `packages/sdk/src/modules/ShopModule.ts` | 14 | JS/TS | `constructor` |
| `packages/sdk/src/modules/ShopModule.ts` | 22 | JS/TS | `getAuthorizedShops` |
| `packages/sdk/src/modules/ShopModule.ts` | 37 | JS/TS | `getCategoryAssets` |
| `packages/sdk/src/modules/__tests__/productModule.test.ts` | 3 | JS/TS | `append` |
| `packages/sdk/src/modules/__tests__/productModule.test.ts` | 6 | JS/TS | `get` |
| `packages/sdk/src/sdk/TikTokShopSDK.ts` | 51 | JS/TS | `constructor` |
| `packages/sdk/src/sdk/TikTokShopSDK.ts` | 122 | JS/TS | `setAccessToken` |
| `packages/sdk/src/sdk/TikTokShopSDK.ts` | 126 | JS/TS | `setShopCipher` |
| `packages/sdk/src/sdk/TikTokShopSDK.ts` | 130 | JS/TS | `setCategoryAssetsCipher` |
| `packages/sdk/src/utils/handleResponse.ts` | 7 | JS/TS | `constructor` |
| `packages/sdk/src/utils/signature.ts` | 15 | JS/TS | `generateSignature` |

### Automation pipeline files

- `.github/workflows/ci.yml`
- `.github/workflows/publish.yml`
- `apps/docs/core-api/finance/get-transactions-by-order.md`
- `apps/docs/core-api/finance/get-transactions-by-statement.md`
- `apps/examples/src/finances/getTransactionsByOrder.ts`
- `apps/examples/src/finances/getTransactionsByStatement.ts`

## `tiktokshop-php`

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
| `tests/Resources/FulfillmentTest.php` | 47 | PHP | `testSearchPackage` |
| `tests/Resources/FulfillmentTest.php` | 53 | PHP | `testCombinePackage` |
| `tests/Resources/FulfillmentTest.php` | 59 | PHP | `testFulfillmentUploadDeliveryFile` |
| `tests/Resources/FulfillmentTest.php` | 65 | PHP | `testUpdateShippingInfo` |
| `tests/Resources/FulfillmentTest.php` | 72 | PHP | `testMarkPackageAsShipped` |
| `tests/Resources/FulfillmentTest.php` | 79 | PHP | `testGetTracking` |
| `tests/Resources/FulfillmentTest.php` | 86 | PHP | `testGetOrderSplitAttributes` |
| `tests/Resources/FulfillmentTest.php` | 92 | PHP | `testUncombinePackages` |
| `tests/Resources/FulfillmentTest.php` | 99 | PHP | `testGetPackageDetail` |
| `tests/Resources/FulfillmentTest.php` | 106 | PHP | `testBatchShipPackages` |
| `tests/Resources/FulfillmentTest.php` | 112 | PHP | `testSplitOrders` |
| `tests/Resources/FulfillmentTest.php` | 119 | PHP | `testCreatePackages` |
| `tests/Resources/FulfillmentTest.php` | 125 | PHP | `testShipPackage` |
| `tests/Resources/FulfillmentTest.php` | 132 | PHP | `testUpdatePackageShippingInfo` |
| `tests/Resources/FulfillmentTest.php` | 139 | PHP | `testFulfillmentUploadDeliveryImage` |
| `tests/Resources/FulfillmentTest.php` | 145 | PHP | `testGetEligibleShippingService` |
| `tests/Resources/GlobalProductTest.php` | 23 | PHP | `testGetGlobalCategories` |
| `tests/Resources/GlobalProductTest.php` | 29 | PHP | `testGetGlobalAttributes` |
| `tests/Resources/GlobalProductTest.php` | 35 | PHP | `testGetGlobalCategoryRules` |
| `tests/Resources/GlobalProductTest.php` | 41 | PHP | `testUpdateGlobalInventory` |
| `tests/Resources/GlobalProductTest.php` | 47 | PHP | `testEditGlobalProduct` |
| `tests/Resources/GlobalProductTest.php` | 53 | PHP | `testSearchGlobalProducts` |
| `tests/Resources/GlobalProductTest.php` | 59 | PHP | `testGetGlobalProduct` |
| `tests/Resources/GlobalProductTest.php` | 65 | PHP | `testDeleteGlobalProducts` |
| `tests/Resources/GlobalProductTest.php` | 71 | PHP | `testCreateGlobalProduct` |
| `tests/Resources/GlobalProductTest.php` | 77 | PHP | `testRecommendGlobalCategories` |
| `tests/Resources/GlobalProductTest.php` | 83 | PHP | `testPublishGlobalProduct` |
| `tests/Resources/LogisticTest.php` | 21 | PHP | `testGetGlobalSellerWarehouse` |
| `tests/Resources/LogisticTest.php` | 27 | PHP | `testGetShippingProvider` |
| `tests/Resources/LogisticTest.php` | 34 | PHP | `testGetWarehouseDeliveryOptions` |
| `tests/Resources/LogisticTest.php` | 41 | PHP | `testGetWarehouseList` |
| `tests/Resources/OrderTest.php` | 20 | PHP | `testGetOrderDetail` |
| `tests/Resources/OrderTest.php` | 26 | PHP | `testGetOrderList` |
| `tests/Resources/ProductTest.php` | 23 | PHP | `testDeactivateProducts` |
| `tests/Resources/ProductTest.php` | 29 | PHP | `testGetBrands` |
| `tests/Resources/ProductTest.php` | 35 | PHP | `testCreateProduct` |
| `tests/Resources/ProductTest.php` | 41 | PHP | `testSearchProducts` |
| `tests/Resources/ProductTest.php` | 47 | PHP | `testGetAttributes` |
| `tests/Resources/ProductTest.php` | 53 | PHP | `testUploadProductImage` |
| `tests/Resources/ProductTest.php` | 59 | PHP | `testGetProduct` |
| `tests/Resources/ProductTest.php` | 65 | PHP | `testUploadProductFile` |
| `tests/Resources/ProductTest.php` | 71 | PHP | `testUpdateInventory` |
| `tests/Resources/ProductTest.php` | 77 | PHP | `testInventorySearch` |
| `tests/Resources/ProductTest.php` | 83 | PHP | `testDeleteProducts` |
| `tests/Resources/ProductTest.php` | 89 | PHP | `testEditProduct` |
| `tests/Resources/ProductTest.php` | 95 | PHP | `testActivateProducts` |
| `tests/Resources/ProductTest.php` | 101 | PHP | `testPartialEditProduct` |
| `tests/Resources/ProductTest.php` | 107 | PHP | `testRecommendCategory` |
| `tests/Resources/ProductTest.php` | 113 | PHP | `testCheckListingPrerequisites` |
| `tests/Resources/ProductTest.php` | 119 | PHP | `testUpdatePrice` |
| `tests/Resources/ProductTest.php` | 125 | PHP | `testGetCategoryRules` |
| `tests/Resources/ProductTest.php` | 131 | PHP | `testGetCategories` |
| `tests/Resources/ProductTest.php` | 137 | PHP | `testRecoverProducts` |
| `tests/Resources/ProductTest.php` | 143 | PHP | `testCreateCustomBrand` |
| `tests/Resources/PromotionTest.php` | 21 | PHP | `testSearchActivities` |
| `tests/Resources/PromotionTest.php` | 27 | PHP | `testGetActivity` |
| `tests/Resources/PromotionTest.php` | 34 | PHP | `testUpdateActivity` |
| `tests/Resources/PromotionTest.php` | 41 | PHP | `testDeactivateActivity` |
| `tests/Resources/PromotionTest.php` | 48 | PHP | `testUpdateActivityProduct` |
| `tests/Resources/PromotionTest.php` | 55 | PHP | `testCreateActivity` |
| `tests/Resources/PromotionTest.php` | 61 | PHP | `testRemoveActivityProduct` |
| `tests/Resources/PromotionTest.php` | 68 | PHP | `testSearchCoupons` |
| `tests/Resources/PromotionTest.php` | 74 | PHP | `testGetCoupon` |
| `tests/Resources/ReturnRefundTest.php` | 21 | PHP | `testRejectCancellation` |
| `tests/Resources/ReturnRefundTest.php` | 27 | PHP | `testCreateReturn` |
| `tests/Resources/ReturnRefundTest.php` | 33 | PHP | `testApproveCancellation` |
| `tests/Resources/ReturnRefundTest.php` | 39 | PHP | `testGetRejectReasons` |
| `tests/Resources/ReturnRefundTest.php` | 45 | PHP | `testCancelOrder` |
| `tests/Resources/ReturnRefundTest.php` | 51 | PHP | `testGetAftersaleEligibility` |
| `tests/Resources/ReturnRefundTest.php` | 57 | PHP | `testSearchReturns` |
| `tests/Resources/ReturnRefundTest.php` | 63 | PHP | `testApproveReturn` |
| `tests/Resources/ReturnRefundTest.php` | 69 | PHP | `testRejectReturn` |
| `tests/Resources/ReturnRefundTest.php` | 75 | PHP | `testSearchCancellations` |
| `tests/Resources/ReturnRefundTest.php` | 81 | PHP | `testCalculateRefund` |
| `tests/Resources/ReturnRefundTest.php` | 87 | PHP | `testGetReturnRecords` |
| `tests/Resources/SellerTest.php` | 21 | PHP | `testGetActiveShopList` |
| `tests/Resources/SellerTest.php` | 27 | PHP | `testGetSellerPermissions` |
| `tests/Resources/SupplychainTest.php` | 20 | PHP | `testConfirmPackageShipment` |
| `tests/TestResource.php` | 30 | PHP | `setUp` |
| `tests/TestResource.php` | 55 | PHP | `tiktokShopClientForTest` |
| `tests/TestResource.php` | 63 | PHP | `assertPreviousRequest` |
| `tests/WebhookTest.php` | 24 | PHP | `setUpBeforeClass` |
| `tests/WebhookTest.php` | 37 | PHP | `testGetShopId` |
| `tests/WebhookTest.php` | 42 | PHP | `testVerify` |
| `tests/WebhookTest.php` | 58 | PHP | `testGetType` |
| `tests/WebhookTest.php` | 63 | PHP | `testGetTimestamp` |
| `tests/WebhookTest.php` | 68 | PHP | `testGetData` |
| `tests/WebhookTest.php` | 74 | PHP | `test__construct` |
| `tests/WebhookTest.php` | 86 | PHP | `testCapture` |

### Automation pipeline files

- `.github/workflows/ci.yml`

## `zLinebot-automos`

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
| `frontend/src/App.jsx` | 8 | JS/TS | `resolveRoute` |
| `frontend/src/components/AuthForm.jsx` | 3 | JS/TS | `messageFromResult` |
| `frontend/src/components/AuthForm.jsx` | 31 | JS/TS | `handleSubmit` |
| `frontend/src/components/ChatDemo.jsx` | 9 | JS/TS | `onAsk` |
| `frontend/src/lib/api.js` | 7 | JS/TS | `tenantHeader` |
| `frontend/src/lib/api.js` | 11 | JS/TS | `parseResponse` |
| `frontend/src/lib/api.js` | 31 | JS/TS | `apiRequest` |
| `frontend/src/lib/api.js` | 53 | JS/TS | `apiPost` |
| `frontend/src/lib/api.js` | 65 | JS/TS | `apiGet` |
| `frontend/src/lib/api.js` | 74 | JS/TS | `apiPatch` |
| `frontend/src/lib/api.js` | 86 | JS/TS | `postRegister` |
| `frontend/src/lib/api.js` | 90 | JS/TS | `postLogin` |
| `frontend/src/lib/api.js` | 94 | JS/TS | `postChat` |
| `frontend/src/lib/api.js` | 98 | JS/TS | `getMe` |
| `frontend/src/lib/api.js` | 102 | JS/TS | `getStats` |
| `frontend/src/lib/api.js` | 106 | JS/TS | `getLeads` |
| `frontend/src/lib/api.js` | 111 | JS/TS | `patchLead` |
| `frontend/src/lib/api.js` | 115 | JS/TS | `getRevenueDaily` |
| `frontend/src/lib/api.js` | 119 | JS/TS | `getTemplates` |
| `frontend/src/lib/api.js` | 123 | JS/TS | `createTemplate` |
| `frontend/src/lib/api.js` | 127 | JS/TS | `getCampaigns` |
| `frontend/src/lib/api.js` | 131 | JS/TS | `createBroadcast` |
| `frontend/src/lib/api.js` | 135 | JS/TS | `createCheckout` |
| `frontend/src/pages/DashboardPage.jsx` | 17 | JS/TS | `formatMoney` |
| `frontend/src/pages/DashboardPage.jsx` | 25 | JS/TS | `readSession` |
| `frontend/src/pages/DashboardPage.jsx` | 34 | JS/TS | `actionsForStatus` |
| `frontend/src/pages/DashboardPage.jsx` | 43 | JS/TS | `formatDateTime` |
| `frontend/src/pages/DashboardPage.jsx` | 201 | JS/TS | `updateLeadStatus` |
| `frontend/src/pages/DashboardPage.jsx` | 206 | JS/TS | `submitTemplate` |
| `frontend/src/pages/DashboardPage.jsx` | 214 | JS/TS | `submitBroadcast` |
| `frontend/src/pages/DashboardPage.jsx` | 225 | JS/TS | `startBilling` |
| `frontend/src/pages/DashboardPage.jsx` | 239 | JS/TS | `logout` |
| `frontend/src/pages/DashboardPage.jsx` | 538 | JS/TS | `MetricCard` |
| `frontend/src/pages/DashboardPage.jsx` | 547 | JS/TS | `Panel` |
| `frontend/src/pages/DashboardPage.jsx` | 559 | JS/TS | `EmptyState` |
| `frontend/src/pages/DashboardPage.jsx` | 567 | JS/TS | `StatusBadge` |
| `frontend/src/pages/LoginPage.jsx` | 10 | JS/TS | `submitLogin` |
| `frontend/src/pages/SignupPage.jsx` | 10 | JS/TS | `submitSignup` |
| `infra/monitor/health.sh` | 7 | Shell | `check` |
| `install_and_run.sh` | 7 | Shell | `log` |
| `install_and_run.sh` | 8 | Shell | `warn` |
| `install_and_run.sh` | 9 | Shell | `err` |
| `install_and_run.sh` | 11 | Shell | `require_cmd` |
| `install_full_stack.sh` | 7 | Shell | `log` |
| `install_full_stack.sh` | 8 | Shell | `require_root` |
| `installer/install.sh` | 21 | Shell | `require_option_value` |
| `installer/install.sh` | 29 | Shell | `usage` |
| `installer/lib/common.sh` | 13 | Shell | `log` |
| `installer/lib/common.sh` | 17 | Shell | `die` |
| `installer/lib/common.sh` | 22 | Shell | `require_root` |
| `installer/lib/common.sh` | 28 | Shell | `require_cmds` |
| `installer/lib/common.sh` | 39 | Shell | `ensure_parent_dir` |
| `installer/lib/common.sh` | 43 | Shell | `sanitize_domain` |
| `installer/lib/common.sh` | 51 | Shell | `is_public_domain` |
| `installer/lib/common.sh` | 56 | Shell | `write_file` |
| `installer/lib/runtime.sh` | 10 | Shell | `preflight_checks` |
| `installer/lib/runtime.sh` | 16 | Shell | `install_system_packages` |
| `installer/lib/stack.sh` | 10 | Shell | `stage_source_tree` |
| `installer/lib/stack.sh` | 49 | Shell | `generate_stack_env` |
| `installer/lib/stack.sh` | 119 | Shell | `configure_tls_assets` |
| `installer/lib/stack.sh` | 164 | Shell | `export_stack_archive` |
| `installer/lib/stack.sh` | 173 | Shell | `prepare_stack` |
| `investor/portal.jsx` | 5 | JS/TS | `hasInvestorAccess` |
| `landing/src/App.jsx` | 8 | JS/TS | `startCheckout` |
| `landing/src/App.jsx` | 16 | JS/TS | `scrollToDemo` |
| `landing/src/components/DemoChat.jsx` | 9 | JS/TS | `sendMsg` |
| `marketplace/service.ts` | 5 | JS/TS | `listBots` |
| `marketplace/service.ts` | 16 | JS/TS | `purchaseBot` |
| `mobile/App.tsx` | 16 | JS/TS | `connect` |
| `one-click-deploy-config-installer-starter.sh` | 17 | Shell | `usage` |
| `one-click-deploy-config-installer-starter.sh` | 50 | Shell | `log` |
| `one-click-deploy-config-installer-starter.sh` | 51 | Shell | `err` |
| `one-click-deploy-config-installer-starter.sh` | 53 | Shell | `require_cmd` |
| `portfolio/engine.ts` | 3 | JS/TS | `aggregate` |
| `portfolio/optimizer.js` | 3 | JS/TS | `normalize` |
| `portfolio/optimizer.js` | 12 | JS/TS | `covarianceMatrix` |
| `portfolio/optimizer.js` | 31 | JS/TS | `optimize` |
| `rl/agent.py` | 12 | Python | `__init__` |
| `rl/agent.py` | 38 | Python | `act` |
| `rl/agent.py` | 47 | Python | `remember` |
| `rl/agent.py` | 50 | Python | `train` |
| `rl/agent.py` | 87 | Python | `sync_target` |
| `rl/infer.py` | 9 | Python | `main` |
| `rl/model.py` | 5 | Python | `__init__` |
| `rl/model.py` | 15 | Python | `forward` |
| `rl/train.py` | 11 | Python | `step_environment` |
| `rl/train.py` | 21 | Python | `shape_reward` |
| `rl/train.py` | 25 | Python | `main` |
| `scripts/build-images.sh` | 7 | Shell | `usage` |
| `scripts/codemods/fix_backend.py` | 5 | Python | `process_file` |
| `scripts/codemods/fix_backend.py` | 20 | Python | `run` |
| `scripts/codemods/fix_frontend.js` | 4 | JS/TS | `processFile` |
| `scripts/codemods/fix_frontend.js` | 13 | JS/TS | `walk` |
| `scripts/deploy-images.sh` | 13 | Shell | `usage` |
| `scripts/run-stack.sh` | 8 | Shell | `ensure_env` |
| `scripts/runner_ultra_v4.3.sh` | 24 | Shell | `fail` |
| `scripts/runner_ultra_v4.3.sh` | 25 | Shell | `ok` |
| `scripts/runner_ultra_v4.3.sh` | 27 | Shell | `retry` |
| `scripts/runner_ultra_v5.sh` | 29 | Shell | `fail` |
| `scripts/runner_ultra_v5.sh` | 30 | Shell | `ok` |
| `scripts/runner_ultra_v5.sh` | 32 | Shell | `cleanup` |
| `scripts/runner_ultra_v5.sh` | 40 | Shell | `retry` |
| `scripts/zlinebot_autonomos.sh` | 12 | Shell | `log` |
| `scripts/zlinebot_autonomos.sh` | 14 | Shell | `display_path` |
| `scripts/zlinebot_autonomos.sh` | 30 | Shell | `ensure_cmd` |
| `scripts/zlinebot_autonomos.sh` | 37 | Shell | `write_if_missing` |
| `scripts/zlinebot_autonomos.sh` | 49 | Shell | `main` |
| `security/middleware.py` | 5 | Python | `security_layer` |
| `stack-workflow-manager.sh` | 23 | Shell | `usage` |
| `stack-workflow-manager.sh` | 44 | Shell | `setup_logging` |
| `stack-workflow-manager.sh` | 60 | Shell | `log` |
| `stack-workflow-manager.sh` | 61 | Shell | `warn` |
| `stack-workflow-manager.sh` | 63 | Shell | `parse_args` |
| `stack-workflow-manager.sh` | 105 | Shell | `collect_signatures` |
| `stack-workflow-manager.sh` | 114 | Shell | `script_inventory` |
| `stack-workflow-manager.sh` | 125 | Shell | `inspect_exact_duplicates` |
| `stack-workflow-manager.sh` | 168 | Shell | `inspect_overlap` |
| `stack-workflow-manager.sh` | 190 | Shell | `print_priority` |
| `stack-workflow-manager.sh` | 199 | Shell | `confirm_run` |
| `stack-workflow-manager.sh` | 211 | Shell | `execute_workflow` |
| `stack-workflow-manager.sh` | 238 | Shell | `main` |
| `start-zLineBot-automos.sh` | 17 | Shell | `usage` |
| `start-zLineBot-automos.sh` | 37 | Shell | `ensure_runtime` |
| `start-zLineBot-automos.sh` | 44 | Shell | `ensure_install_exists` |
| `start-zLineBot-automos.sh` | 48 | Shell | `current_lock_version` |
| `start-zLineBot-automos.sh` | 54 | Shell | `resolve_target_ref` |
| `start-zLineBot-automos.sh` | 64 | Shell | `clone_or_update_repo` |
| `start-zLineBot-automos.sh` | 91 | Shell | `create_service` |
| `start-zLineBot-automos.sh` | 115 | Shell | `prepare_install_tree` |
| `start-zLineBot-automos.sh` | 131 | Shell | `start_service` |
| `start-zLineBot-automos.sh` | 138 | Shell | `stop_service` |
| `start-zLineBot-automos.sh` | 144 | Shell | `restart_service` |
| `start-zLineBot-automos.sh` | 150 | Shell | `status_service` |
| `start-zLineBot-automos.sh` | 155 | Shell | `logs_service` |
| `start-zLineBot-automos.sh` | 160 | Shell | `install_app` |
| `start-zLineBot-automos.sh` | 207 | Shell | `uninstall_app` |
| `start-zLineBot-automos.sh` | 221 | Shell | `update_app` |
| `start-zLineBot-automos.sh` | 236 | Shell | `upgrade_app` |
| `start-zLineBot-automos.sh` | 246 | Shell | `lock_version` |
| `start-zLineBot-automos.sh` | 259 | Shell | `main` |
| `tools/control_panel/app.py` | 29 | Python | `_available_scripts` |
| `tools/control_panel/app.py` | 33 | Python | `_safe_log_path` |
| `tools/control_panel/app.py` | 43 | Python | `_timestamp` |
| `tools/control_panel/app.py` | 47 | Python | `_spawn_script` |
| `tools/control_panel/app.py` | 74 | Python | `_iter_log_lines` |
| `tools/control_panel/app.py` | 81 | Python | `index` |
| `tools/control_panel/app.py` | 91 | Python | `run_script` |
| `tools/control_panel/app.py` | 98 | Python | `view_log` |
| `tools/control_panel/app.py` | 111 | Python | `stream_log` |
| `tools/control_panel/app.py` | 116 | Python | `event_stream` |
| `zBAutomos/ai/cnn_ob.py` | 15 | Python | `__init__` |
| `zBAutomos/ai/cnn_ob.py` | 31 | Python | `forward` |
| `zBAutomos/ai/infer.py` | 11 | Python | `predict` |
| `zBAutomos/ai/infer_cnn.py` | 12 | Python | `predict` |
| `zBAutomos/ai/model.py` | 5 | Python | `__init__` |
| `zBAutomos/ai/model.py` | 10 | Python | `forward` |
| `zBAutomos/ai/train.py` | 4 | Python | `train` |
| `zBAutomos/ai/train_cnn.py` | 8 | Python | `train` |
| `zBAutomos/analytics/core.ts` | 1 | JS/TS | `sharpe` |
| `zBAutomos/analytics/core.ts` | 12 | JS/TS | `drawdown` |
| `zBAutomos/analytics/core.ts` | 26 | JS/TS | `var95` |
| `zBAutomos/api/src/auth/routes.ts` | 5 | JS/TS | `registerAuthRoutes` |
| `zBAutomos/api/src/copy/routes.ts` | 9 | JS/TS | `registerCopyRoutes` |
| `zBAutomos/api/src/market/routes.ts` | 3 | JS/TS | `registerMarketRoutes` |
| `zBAutomos/api/src/server.ts` | 35 | JS/TS | `start` |
| `zBAutomos/api/src/trading/propagation.ts` | 12 | JS/TS | `propagate` |
| `zBAutomos/api/src/trading/ranking.ts` | 7 | JS/TS | `rank` |
| `zBAutomos/api/src/ws/routes.ts` | 3 | JS/TS | `registerWsRoutes` |
| `zBAutomos/backtest/data_loader.ts` | 3 | JS/TS | `loadTicksFromJson` |
| `zBAutomos/backtest/tick_engine.ts` | 9 | JS/TS | `constructor` |
| `zBAutomos/data/ob_collector.ts` | 8 | JS/TS | `toImbalance` |
| `zBAutomos/install_zba.sh` | 9 | Shell | `require_cmd` |
| `zBAutomos/options/vol_surface.py` | 4 | Python | `implied_volatility` |
| `zBAutomos/options/vol_surface.py` | 13 | Python | `surface` |
| `zBAutomos/worker/engine.ts` | 1 | JS/TS | `runWorkerLoop` |
| `zBAutomos/worker/execution.ts` | 13 | JS/TS | `execute` |
| `zBAutomos/worker/risk.ts` | 1 | JS/TS | `enforceRisk` |
| `zBAutomos/worker/src/arb.ts` | 15 | JS/TS | `latencyArb` |
| `zBAutomos/worker/src/arb.ts` | 49 | JS/TS | `execBuy` |
| `zBAutomos/worker/src/arb.ts` | 53 | JS/TS | `execSell` |
| `zBAutomos/worker/src/arb.ts` | 57 | JS/TS | `placeOrder` |
| `zBAutomos/worker/src/execution.ts` | 8 | JS/TS | `executeAtomicHedge` |
| `zBAutomos/worker/src/market_maker.ts` | 3 | JS/TS | `makeMarket` |
| `zBAutomos/worker/src/queue.ts` | 7 | JS/TS | `enqueue` |
| `zBAutomos/worker/src/queue.ts` | 12 | JS/TS | `drain` |
| `zBAutomos/worker/src/risk.ts` | 9 | JS/TS | `constructor` |
| `zBAutomos/worker/src/risk.ts` | 11 | JS/TS | `updatePnL` |
| `zBAutomos/worker/src/risk.ts` | 15 | JS/TS | `updateReturns` |
| `zBAutomos/worker/src/risk.ts` | 22 | JS/TS | `checkPreTrade` |
| `zBAutomos/worker/src/risk.ts` | 37 | JS/TS | `exposure` |
| `zBAutomos/worker/src/risk.ts` | 41 | JS/TS | `var95` |

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

## `zeaz-platform`

Status: **cloned**; commit `b22f68a3f6b8`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

- `install-ubuntu-dev-stack.sh`

## `zeapay`

Status: **blocked**; commit `blocked`.

### API/endpoints

None detected or repository blocked.

### Functions/methods

None detected or repository blocked.

### Automation pipeline files

None detected or repository blocked.

