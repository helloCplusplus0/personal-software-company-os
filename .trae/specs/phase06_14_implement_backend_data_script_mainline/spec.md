# Phase06-14 Onboarding + Sovereignty + Reuse 后端、数据与脚本主线 Spec

## Why

`phase06-08` 已冻结 `Export / Backup` 的后端模块边界，`phase06-09` 已冻结 `ReuseSummaryRead` 的 query owner 与页面集成语义，`phase06-11` 已冻结 `reset_phase06_acceptance.sh` 与 11 个 fixture 白名单，`phase06-12` 已把页面、交互、写路径、导出 / 备份、复用读模型、合同与验收基线收口为正式正文，`phase06-13` 又把 `.proto` 合同推进成仓库唯一合同源。但仓库当前仍缺少 `Onboarding / Export / Backup / Reuse Summary` 的真实后端模块、数据承接位、脚本编排与 fixture 实现。

如果继续停留在 formal spec 和 `.proto` 层，`phase06` 的关键路径仍无法真正跑通：`/onboarding` 没有后端状态源，导出 / 备份无法形成可重复验证的快照主线，复用感知没有真实派生读，`reset_phase06_acceptance.sh` 也无法把 11 个 fixture 可靠建立起来。`phase06-14` 的目标，就是把这些冻结结论推进为可编译、可运行、可重复执行、且不依赖手工补票的数据与后端主线。

## What Changes

- 新增 `backend/internal/onboarding/`，承接 `GetFirstRunState` 的最小只读后端主线
- 新增 `backend/internal/export/`，承接 `GetExportSnapshot / ExportCoreAssets` 的后端读写主线
- 新增 `backend/internal/backup/`，承接 `GetBackupSnapshot / CreateInstanceBackup` 的后端读写主线，并显式落实 `read / verify` 子路径
- 新增 `backend/internal/reusesummary/`，承接 `GetReuseSummary` 的只读派生查询主线
- 新增 `database/migrations/0007_phase06_backend_data_mainline.sql`，为 `phase06` 补齐 `modules.capability_key`、导出元数据与备份元数据承接位
- 新增 `database/scripts/reset_phase06_acceptance.sh`，复用既有 `database/scripts/` 工具链建立 `phase06` 默认基线与 11 个 fixture
- 新增 `database/seeds/seed_phase06_acceptance_baseline.sql` 与 11 个 `seed_phase06_fixture_*.sql`
- 更新 `backend/internal/platform/router.go` 与 `server.go`，把 `Onboarding / Export / Backup / Reuse Summary` 主线装配到现有 `chi + pgx` 组合根
- 明确 `phase06` 当前阶段以数据库元数据表承接导出 / 备份快照与恢复前提校验，不再依赖一次性写入响应或手工构造校验结果
- **BREAKING**：`phase06` 后续前端实现、联调与验收必须直接消费这条后端与数据主线；不得继续依赖手工补 SQL、手工改状态、手工伪造备份结果或第二套临时 DTO 语义

## Impact

- Affected specs:
  - `phase06_08_design_export_backup_backend_module_boundary`
  - `phase06_09_design_reuse_summary_query_dashboard_detail_integration`
  - `phase06_11_design_integration_acceptance_fixtures_restore_baseline`
  - `phase06_12_onboarding_sovereignty_reuse_formal_spec`
  - `phase06_13_land_minimal_proto_contract_mainline`
  - 后续 `phase06` 前端实现与联调验收
- Affected code:
  - `backend/internal/onboarding/`
  - `backend/internal/export/`
  - `backend/internal/backup/`
  - `backend/internal/reusesummary/`
  - `backend/internal/platform/router.go`
  - `backend/internal/platform/server.go`
  - `database/migrations/0007_phase06_backend_data_mainline.sql`
  - `database/scripts/reset_phase06_acceptance.sh`
  - `database/seeds/seed_phase06_acceptance_baseline.sql`
  - `database/seeds/seed_phase06_fixture_*.sql`

## ADDED Requirements

### Requirement: Phase06 后端模块必须按现有主线结构落地

系统 SHALL 将 `Onboarding / Export / Backup / Reuse Summary` 落地到现有 `backend/internal/` 主线，沿用仓库已验证的 `handler / service / candidate / repository / types.go / errors.go / response.go` 结构，而不是把 `phase06` 逻辑散落回既有 canonical 模块或 `platform` 包。

#### Scenario: Onboarding 模块物理落点

- **WHEN** 实现 `GetFirstRunState`
- **THEN** 仓库中必须存在 `backend/internal/onboarding/`
- **AND** 必须至少包含：
  - `handler/query_handler.go`
  - `handler/response.go`
  - `service/query_service.go`
  - `candidate/first_run_readers.go`
  - `types.go`
  - `errors.go`
- **AND** 当前阶段不得为 `Onboarding` 新增独立写组 service
- **AND** draft-first 写入继续复用既有 canonical create 模块

#### Scenario: Export / Backup / Reuse Summary 模块物理落点

- **WHEN** 实现 `Export / Backup / Reuse Summary`
- **THEN** 必须存在：
  - `backend/internal/export/`
  - `backend/internal/backup/`
  - `backend/internal/reusesummary/`
- **AND** `export/` 至少包含 `handler/query_handler.go`、`handler/command_handler.go`、`handler/response.go`、`service/query_service.go`、`service/command_service.go`、`candidate/asset_reader.go`、`repository/export_store.go`、`types.go`、`errors.go`
- **AND** `backup/` 至少包含 `handler/query_handler.go`、`handler/command_handler.go`、`handler/response.go`、`service/query_service.go`、`service/command_service.go`、`candidate/asset_reader.go`、`repository/backup_store.go`、`types.go`、`errors.go`
- **AND** `reusesummary/` 至少包含 `handler/query_handler.go`、`handler/response.go`、`service/query_service.go`、`candidate/reuse_readers.go`、`types.go`、`errors.go`
- **AND** `service/` 层不得直接跨模块写 SQL，跨模块只读必须由各自 `candidate/` 子包拥有

### Requirement: Phase06 数据主线必须通过 0007 migration 补齐最小承接位

系统 SHALL 通过 `database/migrations/0007_phase06_backend_data_mainline.sql` 为 `phase06` 补齐最小数据承接位，使复用聚合、导出快照、备份快照与恢复前提校验都能落在仓库主线内。

#### Scenario: `modules.capability_key` 原位升级

- **WHEN** 执行 `0007_phase06_backend_data_mainline.sql`
- **THEN** 必须为 `modules` 表原位新增 `capability_key TEXT NULL`
- **AND** 不得引入独立 `capabilities` 表、可编辑能力字典或第二套聚合事实源
- **AND** 未填写 `capability_key` 的 `Module` 必须继续允许存在
- **AND** 当前阶段不得把未填写 `capability_key` 误判为写入失败

#### Scenario: 导出与备份元数据表落地

- **WHEN** 执行 `0007_phase06_backend_data_mainline.sql`
- **THEN** 必须新增 `instance_exports` 与 `instance_backups` 两张表
- **AND** `instance_exports` 至少必须承接：
  - `id`
  - `created_at`
  - `result_status`
  - `result_summary`
  - `asset_scope_json`
  - `artifact_payload_json`
- **AND** `instance_backups` 至少必须承接：
  - `id`
  - `created_at`
  - `manifest_json`
  - `asset_coverage_json`
  - `schema_version`
  - `instance_version`
  - `verified_status`
  - `verify_failure_code`
  - `backup_payload_json`
- **AND** `verify_failure_code` 只允许承接：
  - `manifest_missing`
  - `coverage_incomplete`
  - `schema_mismatch`
- **AND** 当前阶段不得把导出 / 备份元数据继续悬空在一次性响应体中

### Requirement: `GetFirstRunState` 必须由 canonical 数据读时派生

系统 SHALL 将 `first_run_state` 实现为读时派生结果，而不是引入独立 `first_run_state` 表或通过脚本手工写状态。

#### Scenario: `first_run_state` 状态推导

- **WHEN** `Onboarding QueryService` 读取首轮状态
- **THEN** 必须基于 `products / repositories / modules / decisions` 四类 canonical 数据的当前持久化数量推导状态
- **AND** 四类都为 `0` 时返回 `not_started`
- **AND** 至少存在 `1` 类、但未满四类时返回 `in_progress`
- **AND** 四类都至少存在 `1` 条时返回 `completed`
- **AND** 当前阶段不得通过独立状态表、手工 SQL 或缓存值承接该状态

#### Scenario: `current_step` 与 `completion_progress` 推导

- **WHEN** `first_run_state` 为 `not_started` 或 `in_progress`
- **THEN** `current_step` 必须按 `Product -> Repository -> Module -> Decision` 找到第一个尚未完成的步骤
- **AND** `completion_progress` 必须冻结为 `0 / 25 / 50 / 75 / 100`
- **AND** `completed` 时 `current_step` 必须返回 `complete`
- **AND** 当前阶段不得要求前端自行推导下一步

### Requirement: Export 主线必须装配 9 类核心资产并形成可重复读取的快照

系统 SHALL 将 `Export` 实现为“读取快照 + 触发导出”双路径主线，并通过数据库元数据表形成最新可读取快照，而不是只返回一次性导出响应。

#### Scenario: `ExportCoreAssets` 数据装配

- **WHEN** 执行 `POST /api/dashboard/export`
- **THEN** `export.candidate.AssetReader` 必须装配以下 9 类 canonical 数据：
  - `products`
  - `modules`
  - `releases`
  - `repositories`
  - `decisions`
  - `decision_links`
  - `product_modules`
  - `product_repositories`
  - `module_repositories`
- **AND** `export.command_service` 必须把装配结果持久化到 `instance_exports`
- **AND** 不得只导出主实体而遗漏绑定关系

#### Scenario: `GetExportSnapshot` 读取语义

- **WHEN** 执行 `GET /api/dashboard/export`
- **THEN** 若已存在 `instance_exports` 记录，必须返回最新一条快照摘要
- **AND** 若尚无历史导出记录，必须返回基于当前 canonical 数据现算的预览态 `ExportSnapshot`
- **AND** 预览态不得被误判为错误

### Requirement: Backup 主线必须实现 `read / verify` 子路径与三类失败语义

系统 SHALL 将 `Backup` 实现为“创建备份 + 重新读取并校验恢复前提”的双路径主线，并把三类失败语义稳定保留在后端与数据主线内。

#### Scenario: `CreateInstanceBackup` 持久化语义

- **WHEN** 执行 `POST /api/dashboard/backup`
- **THEN** `backup.command_service` 必须装配与 `Export` 相同的 9 类核心资产
- **AND** 必须同时生成 `manifest_json`、`asset_coverage_json`、`schema_version`、`instance_version`
- **AND** 初始写入时 `verified_status` 必须为 `unverified`
- **AND** 写入结果必须持久化到 `instance_backups`

#### Scenario: `GetBackupSnapshot` 校验链

- **WHEN** 执行 `GET /api/dashboard/backup`
- **THEN** `backup.query_service` 必须读取最新一条 `instance_backups`
- **AND** 必须依次校验：
  - `manifest_json` 是否存在且可承接摘要
  - `asset_coverage_json` 是否覆盖 9 类核心资产
  - `schema_version` 是否与当前 `schema_migrations` 最新版本对齐，且 `instance_version` 可读取
- **AND** 三步全部通过时才允许返回 `verified`
- **AND** 当前阶段不得以“写出成功”直接代替 `read / verify`

#### Scenario: 三类失败语义保持单值化

- **WHEN** `GetBackupSnapshot` 校验失败
- **THEN** 失败原因必须单值回落到以下之一：
  - `manifest_missing`
  - `coverage_incomplete`
  - `schema_mismatch`
- **AND** `instance_backups.verify_failure_code` 必须与之保持一致
- **AND** DTO / HTTP / 前端消费侧不得把这三类失败折叠为泛化的 `backup failed`

### Requirement: Reuse Summary 必须通过读时聚合返回最新已提交状态

系统 SHALL 将 `ReuseSummaryRead` 实现为读时聚合，不引入异步统计表或离线聚合作业。

#### Scenario: `module_reuse_summary` 聚合口径

- **WHEN** 读取 `module_reuse_summary`
- **THEN** `reuse_product_count` 必须表示“当前被多少 Product 直接复用”
- **AND** 只允许基于 `product_modules` 与 `modules` 当前已提交数据聚合
- **AND** 不得把 `Repository` 映射数、`Decision` 链接数或 `Release` 数量混入该统计

#### Scenario: `capability_summary` 聚合与映射

- **WHEN** 读取 `capability_summary`
- **THEN** 必须以 `modules.capability_key` 作为唯一聚合主键来源
- **AND** `capability_label` 必须来自后端内置的单一 `capability_key -> capability_label` 映射
- **AND** 当前阶段不得让后端、前端与 fixture 各自维护三套不同映射表
- **AND** 若当前作用域没有任何可聚合 capability，必须返回成功空态而不是错误

#### Scenario: 排序、裁剪与作用域

- **WHEN** `GetReuseSummary` 按不同 `scope` 返回结果
- **THEN** `dashboard` 作用域下的 `module_reuse_summary` 与 `capability_summary` 都必须按“数量优先、时间次级”排序并最多返回前 `5` 条
- **AND** `module_detail` 作用域必须围绕单一 `module_id` 返回该模块的直接复用反馈，并在存在 `capability_key` 时返回对应 capability 摘要
- **AND** `product_detail` 作用域必须围绕单一 `product_id` 先限定在已绑定模块范围内，再返回全量复用 / capability 摘要

### Requirement: Phase06 路由装配必须接入现有 chi 组合根

系统 SHALL 将 `Onboarding / Export / Backup / Reuse Summary` 装配到现有 `backend/internal/platform/router.go` 与 `server.go`，延续 `chi + pgxpool` 的组合根模式。

#### Scenario: 路由注册矩阵

- **WHEN** 更新 `platform/router.go`
- **THEN** 必须注册以下 HTTP 路由：
  - `GET /api/onboarding/state`
  - `GET /api/dashboard/export`
  - `POST /api/dashboard/export`
  - `GET /api/dashboard/backup`
  - `POST /api/dashboard/backup`
  - `GET /api/reuse-summary`
- **AND** `GetReuseSummary` 的 `scope / module_id / product_id` 必须通过 query 参数显式映射到 Proto request
- **AND** handler 必须在进入 service 前显式组装对应 Proto request

#### Scenario: 装配顺序

- **WHEN** 更新 `platform/server.go`
- **THEN** 必须先装配既有四个 canonical 模块与 `Dashboard`
- **AND** 再装配 `Onboarding / Export / Backup / Reuse Summary`
- **AND** 当前阶段不得把 `phase06` 模块塞回既有 canonical 模块的 mount 函数内部

### Requirement: `reset_phase06_acceptance.sh` 必须把 11 个 fixture 变成可重复执行的正式主线

系统 SHALL 实现 `database/scripts/reset_phase06_acceptance.sh` 及其配套 seed，使 `phase06` 的默认基线与 11 个 fixture 都能通过受控脚本建立，不依赖手工补票。

#### Scenario: 脚本模式与编排关系

- **WHEN** 实现 `reset_phase06_acceptance.sh`
- **THEN** 必须支持：
  - 默认模式
  - `--clean-only`
  - `--restore-only`
  - `--fixture <name>`
- **AND** 必须编排既有：
  - `reset_module_mainline.sh`
  - `reset_decision_mainline.sh`
  - `reset_product_repository_mainline.sh`
  - `reset_dashboard_acceptance.sh`
- **AND** 不得直接用散装 `DELETE / TRUNCATE` 取代既有 reset 主线

#### Scenario: Phase06 清空与恢复边界

- **WHEN** 脚本执行清空
- **THEN** 除编排既有 reset 脚本外，还必须清理 `instance_exports` 与 `instance_backups`
- **AND** 不得清理 `schema_migrations`
- **AND** 清空后必须可立即再次执行默认恢复或任一 fixture 恢复

#### Scenario: Fixture 白名单文件落点

- **WHEN** 实现 `phase06` 基线与 fixture SQL
- **THEN** 必须存在以下文件：
  - `seed_phase06_acceptance_baseline.sql`
  - `seed_phase06_fixture_cold_start_empty.sql`
  - `seed_phase06_fixture_in_progress_partial_entry.sql`
  - `seed_phase06_fixture_completed_unbound.sql`
  - `seed_phase06_fixture_completed_bound.sql`
  - `seed_phase06_fixture_export_ready.sql`
  - `seed_phase06_fixture_backup_verified.sql`
  - `seed_phase06_fixture_backup_manifest_missing.sql`
  - `seed_phase06_fixture_backup_coverage_incomplete.sql`
  - `seed_phase06_fixture_backup_schema_mismatch.sql`
  - `seed_phase06_fixture_reuse_latest.sql`
  - `seed_phase06_fixture_reuse_latest_after_binding.sql`
- **AND** 文件名与白名单必须和 `phase06-11` 完全一致
- **AND** 不得在实现阶段再长出匿名 fixture

### Requirement: Phase06 关键路径必须可通过真实主线重复验证

系统 SHALL 保证 `phase06-14` 完成后，后端、数据与脚本主线可以支撑 `phase06` 的关键路径重复执行。

#### Scenario: 最小运行验收

- **WHEN** 验证 `phase06-14`
- **THEN** 至少必须通过：
  - `go build ./...`
  - `reset_phase06_acceptance.sh` 默认模式可重复执行
  - `reset_phase06_acceptance.sh --fixture <name>` 可切换到 11 个白名单 fixture
  - `GET /api/onboarding/state`
  - `GET/POST /api/dashboard/export`
  - `GET/POST /api/dashboard/backup`
  - `GET /api/reuse-summary`
- **AND** 当前阶段不得依赖手工补 SQL、手工修改 `first_run_state` 或手工伪造备份校验结果完成数据主权验证

## MODIFIED Requirements

### Requirement: Phase06 从“合同主线已落地”推进到“后端与数据主线可运行”

系统 SHALL 将 `phase06-13` 已落地的 Proto 合同，从“仓库内单一合同源”推进为“后端模块、数据承接位与脚本基线均已可运行”的真实实现主线。

#### Scenario: 阶段推进关系

- **WHEN** `phase06-14` 完成
- **THEN** `phase06` 不得再停留在“已有 `.proto`、但后端与数据仍为空壳”的状态
- **AND** 后续前端与联调必须直接引用本阶段交付的真实 API、真实数据装配与真实 fixture 基线

## REMOVED Requirements

### Requirement: 继续通过手工补票建立 Phase06 验收环境

**Reason**: `phase06-11` 已明确禁止通过手工补 SQL、手工改状态或手工伪造备份结果完成验收；`phase06-14` 必须把这一约束推进成真实脚本与数据主线。

**Migration**: 后续统一通过 `reset_phase06_acceptance.sh`、`seed_phase06_acceptance_baseline.sql` 与 11 个 `seed_phase06_fixture_*.sql` 建立环境；导出 / 备份快照与恢复前提统一通过后端主线和数据承接位读取验证，不再依赖一次性手工状态。
