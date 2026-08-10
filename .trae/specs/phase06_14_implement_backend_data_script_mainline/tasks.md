# Tasks

- [x] Task 1: 对齐 `phase06-14` 的直接上游、现有后端模块模式与当前仓库数据现场。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md` 中 `phase06-14` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase06-08 / 09 / 11 / 12 / 13` 已冻结的后端、数据、脚本与验收边界
  - [x] SubTask 1.3: 对齐现有 `platform/router.go`、`server.go`、`chi + pgxpool` 装配模式
  - [x] SubTask 1.4: 对齐现有 `reset_*` 脚本、`database/seeds/` 与 `schema_migrations` 现场

- [x] Task 2: 实现 `phase06` 的最小数据主线 migration。
  - [x] SubTask 2.1: 新增 `database/migrations/0007_phase06_backend_data_mainline.sql`
  - [x] SubTask 2.2: 为 `modules` 原位新增 `capability_key`
  - [x] SubTask 2.3: 新增 `instance_exports` 元数据表
  - [x] SubTask 2.4: 新增 `instance_backups` 元数据表，并承接 `verified_status + verify_failure_code`

- [x] Task 3: 实现 `backend/internal/onboarding/` 的只读主线。
  - [x] SubTask 3.1: 实现 `candidate/first_run_readers.go`，读取四类 canonical 对象计数
  - [x] SubTask 3.2: 实现 `service/query_service.go`，推导 `status / current_step / completion_progress`
  - [x] SubTask 3.3: 实现 `handler/query_handler.go` 与 `handler/response.go`
  - [x] SubTask 3.4: 实现 `types.go / errors.go`，并从 `onboarding.proto` 单向对齐

- [x] Task 4: 实现 `backend/internal/export/` 的读写主线。
  - [x] SubTask 4.1: 实现 `candidate/asset_reader.go`，装配 9 类核心资产
  - [x] SubTask 4.2: 实现 `repository/export_store.go`，持久化最新导出快照
  - [x] SubTask 4.3: 实现 `service/query_service.go`，承接"最新快照 / 预览态快照"读取
  - [x] SubTask 4.4: 实现 `service/command_service.go`，装配并写入导出结果
  - [x] SubTask 4.5: 实现 `handler/query_handler.go`、`handler/command_handler.go`、`handler/response.go`、`types.go / errors.go`

- [x] Task 5: 实现 `backend/internal/backup/` 的读写与 `read / verify` 主线。
  - [x] SubTask 5.1: 实现 `candidate/asset_reader.go`，装配 9 类核心资产
  - [x] SubTask 5.2: 实现 `repository/backup_store.go`，持久化备份快照与校验元数据
  - [x] SubTask 5.3: 实现 `service/command_service.go`，写入 `manifest / asset_coverage / schema_version / instance_version`
  - [x] SubTask 5.4: 实现 `service/query_service.go`，按 `manifest_missing / coverage_incomplete / schema_mismatch` 三类失败语义做 `read / verify`
  - [x] SubTask 5.5: 实现 `handler/query_handler.go`、`handler/command_handler.go`、`handler/response.go`、`types.go / errors.go`

- [x] Task 6: 实现 `backend/internal/reusesummary/` 的派生读主线。
  - [x] SubTask 6.1: 实现 `candidate/reuse_readers.go`，承接 `dashboard / module_detail / product_detail` 三种作用域读取
  - [x] SubTask 6.2: 实现 `service/query_service.go`，落实复用聚合、排序、裁剪与最新状态语义
  - [x] SubTask 6.3: 在后端落地单一 `capability_key -> capability_label` 内置映射
  - [x] SubTask 6.4: 实现 `handler/query_handler.go`、`handler/response.go`、`types.go / errors.go`

- [x] Task 7: 更新 `platform` 装配层并接入 `phase06` 路由。
  - [x] SubTask 7.1: 在 `platform/router.go` 新增 `buildOnboarding / mountOnboarding`
  - [x] SubTask 7.2: 在 `platform/router.go` 新增 `buildExport / mountExport`
  - [x] SubTask 7.3: 在 `platform/router.go` 新增 `buildBackup / mountBackup`
  - [x] SubTask 7.4: 在 `platform/router.go` 新增 `buildReuseSummary / mountReuseSummary`
  - [x] SubTask 7.5: 在 `platform/server.go` 按既有 canonical 模块与 `Dashboard` 之后的顺序接入 `phase06` 模块

- [x] Task 8: 实现 `phase06` 验收重置脚本与 baseline / fixture SQL。
  - [x] SubTask 8.1: 实现 `database/scripts/reset_phase06_acceptance.sh`
  - [x] SubTask 8.2: 实现 `database/seeds/seed_phase06_acceptance_baseline.sql`
  - [x] SubTask 8.3: 实现 11 个 `seed_phase06_fixture_*.sql`
  - [x] SubTask 8.4: 在脚本中编排既有 `reset_module_mainline.sh`、`reset_decision_mainline.sh`、`reset_product_repository_mainline.sh`、`reset_dashboard_acceptance.sh`
  - [x] SubTask 8.5: 实现 `instance_exports / instance_backups` 的受控清理，不触碰 `schema_migrations`

- [x] Task 9: 验证 `phase06` 后端、数据与脚本主线可运行。
  - [x] SubTask 9.1: 验证 `go build ./...` 通过
  - [x] SubTask 9.2: 验证 `0007_phase06_backend_data_mainline.sql` 可被现有迁移主线执行
  - [x] SubTask 9.3: 验证 `reset_phase06_acceptance.sh` 默认模式与 `--fixture <name>` 可重复执行
  - [x] SubTask 9.4: 验证 `GET /api/onboarding/state`
  - [x] SubTask 9.5: 验证 `GET/POST /api/dashboard/export`
  - [x] SubTask 9.6: 验证 `GET/POST /api/dashboard/backup`
  - [x] SubTask 9.7: 验证 `GET /api/reuse-summary`
  - [x] SubTask 9.8: 验证关键路径不依赖手工补 SQL、手工改 `first_run_state` 或手工伪造备份结果

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`
- `Task 7` depends on `Task 3`, `Task 4`, `Task 5`, `Task 6`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 4`, `Task 5`, `Task 6`
- `Task 9` depends on `Task 2` through `Task 8`
