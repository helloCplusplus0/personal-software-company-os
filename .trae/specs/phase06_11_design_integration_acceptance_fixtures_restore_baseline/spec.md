# Phase06-11 联调验收环境、Fixture 与恢复基线设计 Spec

## Why

`phase06-01 ~ 10` 已分别冻结了 first-run 入口、draft-first、Export / Backup 语义、Reuse Summary 读模型、前端交互流与最小 Proto 合同，但当前仍缺少一份把“如何重复建立验收环境、如何用单值 fixture 证明关键状态、如何验证恢复前提与复用感知最新状态”收成同一套基线的正式规格。若不先冻结联调验收环境、fixture 组合、恢复验证入口与阶段完成验收矩阵，后续实现与验收仍会退回到手工补数据、手工切状态和一次性临时验证。

## What Changes

- 冻结 `phase06` 验收环境必须复用既有 `database/scripts/ + database/seeds/` 模式，不发明第二套验收工具链
- 冻结 `reset_phase06_acceptance.sh` 作为 `phase06` 联调验收统一入口
- 冻结 `cold-start / in-progress / completed` 三类 `first_run_state` fixture
- 冻结“缺少绑定关系仍完成首轮会话”与“绑定补全后再次验证”两类 fixture
- 冻结 `export-ready`、`backup-verified`、`reuse-latest` 三类关键能力 fixture
- 冻结 `backup verified` 的失败路径 fixture：`manifest` 缺失、覆盖矩阵不完整、`schema/version` 不可校验
- 冻结首轮成功会话、数据主权闭合、复用感知最新状态与回访继续的验收矩阵
- 冻结 Onboarding / Export / Backup / Reuse Summary 的合同一致性验证必须纳入验收矩阵
- 冻结 `backup_snapshot` 读取侧合同一致性必须独立验证，不得以 `BackupWrite` 响应附带代替
- 冻结导出/备份验收产物的清理边界
- 冻结 `--restore-only` 模式必须处理同标题 readonly seed 的情况
- **BREAKING**：`phase06` 联调验收不得再依赖手工补 SQL、手工改 `first_run_state` 或手工伪造备份校验结果

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-01` first-run 入口与状态机
  - `phase06-02` draft-first / partial-entry
  - `phase06-03` Export / Backup / restore prerequisites
  - `phase06-04` Reuse Summary 读模型与挂接位
  - `phase06-06` Onboarding 前端路由与交互流
  - `phase06-08` Export / Backup 后端模块边界
  - `phase06-09` Reuse Summary query / Dashboard / Detail 集成
  - `phase06-10` 最小 Proto 合同设计
- Affected code:
  - 后续 `database/scripts/reset_phase06_acceptance.sh`
  - 后续 `database/seeds/seed_phase06_acceptance_baseline.sql`
  - 后续 `database/seeds/seed_phase06_fixture_cold_start_empty.sql`
  - 后续 `database/seeds/seed_phase06_fixture_in_progress_partial_entry.sql`
  - 后续 `database/seeds/seed_phase06_fixture_completed_unbound.sql`
  - 后续 `database/seeds/seed_phase06_fixture_completed_bound.sql`
  - 后续 `database/seeds/seed_phase06_fixture_export_ready.sql`
  - 后续 `database/seeds/seed_phase06_fixture_backup_verified.sql`
  - 后续 `database/seeds/seed_phase06_fixture_backup_manifest_missing.sql`
  - 后续 `database/seeds/seed_phase06_fixture_backup_coverage_incomplete.sql`
  - 后续 `database/seeds/seed_phase06_fixture_backup_schema_mismatch.sql`
  - 后续 `database/seeds/seed_phase06_fixture_reuse_latest.sql`
  - 后续 `database/seeds/seed_phase06_fixture_reuse_latest_after_binding.sql`
  - 后续联调验收脚本、前后端真实 API 验收与浏览器验证用例

## ADDED Requirements

### Requirement: Phase06 验收环境必须复用既有 reset/seed 模式

系统 SHALL 将 `phase06` 的联调验收环境建立方式冻结为复用仓库既有 `database/scripts/` + `database/seeds/` 模式，不得为 Onboarding / Export / Backup / Reuse Summary 验收再发明第二套工具链。

#### Scenario: 验收工具链复用边界

- **WHEN** 后续实现或验收讨论 `phase06` 联调环境如何建立
- **THEN** 必须复用 `database/scripts/` 承接重置入口
- **AND** 必须复用 `database/seeds/` 承接基线种子与 fixture SQL
- **AND** 必须继续复用既有 `psql` 自动检测执行模式
- **AND** 不得在 `backend/`、`frontend/` 或仓库根新增并列验收工具链
- **AND** 不得通过手工修改数据库记录来切换 `first_run_state`、导出结果、备份校验结果或复用感知状态

#### Scenario: 与既有脚本的编排关系

- **WHEN** `phase06` 验收需要建立冷启动、有数据或失败路径状态
- **THEN** 必须通过编排既有 `reset_module_mainline.sh`、`reset_decision_mainline.sh`、`reset_product_repository_mainline.sh` 与 `reset_dashboard_acceptance.sh` 实现
- **AND** 必须继续依赖 `seed_readonly_prereqs.sql` 或等价只读前提数据
- **AND** 不得绕过既有脚本直接 `DELETE / TRUNCATE` canonical 表
- **AND** `phase06` 验收脚本只允许在既有脚本之上增加编排层、fixture 加载层与恢复前提验证层

### Requirement: `reset_phase06_acceptance.sh` 必须作为统一验收入口

系统 SHALL 冻结 `database/scripts/reset_phase06_acceptance.sh` 为 `phase06` 联调验收的唯一统一入口。

#### Scenario: 脚本模式矩阵

- **WHEN** 后续实现或验收讨论 `phase06` 环境如何重置
- **THEN** `reset_phase06_acceptance.sh` 最小必须支持以下模式：
  - 默认（无参数）：先清空，再恢复 `phase06` 默认验收基线
  - `--clean-only`：仅清空当前阶段相关数据与验收产物
  - `--restore-only`：仅恢复默认验收基线
  - `--fixture <name>`：先清空，再加载指定 fixture
- **AND** `<name>` 只允许取本规格冻结的 fixture 白名单之一
- **AND** 不得并列发明第二套 `setup / seed / replay` 脚本入口命名

#### Scenario: 清空范围

- **WHEN** `reset_phase06_acceptance.sh` 执行清空
- **THEN** 必须通过编排既有 mainline / dashboard reset 脚本完成 canonical 数据清空
- **AND** 还必须清理当前阶段验收过程中生成的导出 / 备份验收产物与对应元数据
- **AND** 不得清空 `schema_migrations` 或 migration 基线
- **AND** 清空后环境必须可立即重复执行 `--restore-only` 或 `--fixture <name>`

#### Scenario: 恢复默认基线

- **WHEN** `reset_phase06_acceptance.sh` 执行默认恢复
- **THEN** 必须先恢复 `phase05` 已冻结的 Dashboard / Feedback 验收基线
- **AND** 再加载 `seed_phase06_acceptance_baseline.sql`
- **AND** 恢复后必须形成一套可继续覆盖 Onboarding、Export、Backup 与 Reuse Summary 的默认联调环境
- **AND** 默认恢复必须幂等，可重复执行不报错

#### Scenario: 同标题 readonly seed 处理

- **WHEN** `--restore-only` 模式遇到与正式基线 seed 同标题的 readonly seed 记录
- **THEN** 必须沿用 `phase05` 已冻结的 `UPDATE` 收口 + `INSERT WHERE NOT EXISTS` 补缺策略
- **AND** 不得跳过同标题 readonly seed 的恢复
- **AND** 不得产生重复记录
- **AND** 该行为必须与 `seed_decision_mainline_baseline.sql` 的处理模式保持一致

### Requirement: Fixture 白名单与命名必须单值化

系统 SHALL 将 `phase06` 联调验收使用的 fixture 冻结为单值白名单，避免后续继续新增散装状态。

#### Scenario: Fixture 白名单

- **WHEN** 后续实现或验收讨论 `phase06` 的 fixture 组合
- **THEN** 必须冻结为以下白名单：
  - `cold-start-empty`
  - `in-progress-partial-entry`
  - `completed-unbound`
  - `completed-bound`
  - `export-ready`
  - `backup-verified`
  - `backup-manifest-missing`
  - `backup-coverage-incomplete`
  - `backup-schema-mismatch`
  - `reuse-latest`
  - `reuse-latest-after-binding`
- **AND** 每个 fixture 必须有唯一 SQL 文件或唯一受控生成路径
- **AND** 不得在上述白名单之外临时追加匿名 fixture

### Requirement: `first_run_state` 三类正式 fixture 必须冻结

系统 SHALL 将 `cold-start / in-progress / completed` 三类 `first_run_state` 验收基线冻结为正式 fixture。

#### Scenario: `cold-start-empty`

- **WHEN** 加载 `cold-start-empty`
- **THEN** `Product / Repository / Module / Decision` 四类首轮对象必须全部不存在
- **AND** `first_run_state` 必须为 `not_started`
- **AND** 根级默认进入路径必须回落到 `/onboarding`
- **AND** 当前 fixture 不得伪造任何已完成首轮对象记录

#### Scenario: `in-progress-partial-entry`

- **WHEN** 加载 `in-progress-partial-entry`
- **THEN** 四类首轮对象中至少存在 `1` 条、但未满四类对象最小持久化集合
- **AND** `first_run_state` 必须为 `in_progress`
- **AND** 根级默认进入路径必须为 `/dashboard`
- **AND** `Dashboard` 必须可重复验证 `Continue Onboarding -> /onboarding`
- **AND** 该 fixture 必须能把 `Onboarding` 自动定位到第一个未完成步骤

#### Scenario: `completed-unbound`

- **WHEN** 加载 `completed-unbound`
- **THEN** `Product / Repository / Module / Decision` 四类对象必须都已至少持久化 `1` 条
- **AND** `first_run_state` 必须为 `completed`
- **AND** 当前 fixture 明确允许缺少以下绑定关系：
  - `Product -> Module`
  - `Product -> Repository`
  - `Module -> Repository`
  - `Decision -> target`
- **AND** 该 fixture 必须证明“缺少绑定关系仍完成首轮成功会话”

### Requirement: 绑定补全后再次验证 fixture 必须冻结

系统 SHALL 将“绑定补全后再次验证”冻结为独立 fixture，而不是在 `completed-unbound` 上手工改库。

#### Scenario: `completed-bound`

- **WHEN** 加载 `completed-bound`
- **THEN** 必须延续 `completed-unbound` 的四类最小持久化对象
- **AND** 同时补齐当前阶段冻结的核心绑定关系：
  - `product_modules`
  - `product_repositories`
  - `module_repositories`
  - 需要时的 `decision_links`
- **AND** 该 fixture 必须作为数据主权闭合、导出覆盖矩阵与复用感知验证的正式前置状态
- **AND** 不得要求验收人员在 `completed-unbound` 基础上手工补数据

### Requirement: 数据主权闭合 fixture 必须独立存在

系统 SHALL 将导出可验与备份可验的最小前置状态冻结为独立 fixture 证据，不得混在首轮成功会话 fixture 中模糊处理。

#### Scenario: `export-ready`

- **WHEN** 加载 `export-ready`
- **THEN** 该 fixture 必须至少满足 `completed-bound`
- **AND** 必须具备 `Export` 当前阶段最小覆盖矩阵所需的 9 类 canonical 数据集
- **AND** 验收时必须可以直接验证：
  - `GET /api/dashboard/export` 可读取 `export_snapshot`
  - `POST /api/dashboard/export` 可生成正式导出结果
- **AND** 当前 fixture 不得缺失任何核心绑定关系后仍被判定为“数据主权闭合”

#### Scenario: `backup-verified`

- **WHEN** 加载 `backup-verified`
- **THEN** 该 fixture 必须至少满足 `completed-bound`
- **AND** 必须具备可重复触发 `CreateInstanceBackup` 与 `GetBackupSnapshot` 的正式前置数据
- **AND** 验收时必须可以直接验证：
  - 备份产物可生成
  - `manifest` 可重新读取
  - 覆盖矩阵完整
  - `schema / version` 前提可校验
- **AND** 只有上述条件同时成立，当前 fixture 才允许作为 `backup verified` 的正式证据

### Requirement: `backup verified` 失败路径 fixture 必须单值化

系统 SHALL 将 `backup verified` 的失败路径冻结为三类独立 fixture，避免继续把失败原因混成一个“备份失败”。

#### Scenario: `backup-manifest-missing`

- **WHEN** 加载 `backup-manifest-missing`
- **THEN** 当前 fixture 必须能稳定复现“备份产物存在，但 `manifest` 缺失或不可读取”
- **AND** `backup verified` 不得成立
- **AND** 失败语义必须单值回落到 `manifest` 缺失 / 不可解析

#### Scenario: `backup-coverage-incomplete`

- **WHEN** 加载 `backup-coverage-incomplete`
- **THEN** 当前 fixture 必须能稳定复现“`manifest` 可读，但核心覆盖矩阵不完整”
- **AND** `backup verified` 不得成立
- **AND** 失败语义必须单值回落到覆盖矩阵缺失 / 不完整

#### Scenario: `backup-schema-mismatch`

- **WHEN** 加载 `backup-schema-mismatch`
- **THEN** 当前 fixture 必须能稳定复现“`manifest` 与覆盖矩阵可读，但 `schema / version` 前提不可校验”
- **AND** `backup verified` 不得成立
- **AND** 失败语义必须单值回落到 `schema / version` 不匹配或不可验证

### Requirement: 复用感知最新状态 fixture 必须冻结

系统 SHALL 将“复用感知可见”和“绑定补全后读取到最新状态”冻结为两类独立 fixture。

#### Scenario: `reuse-latest`

- **WHEN** 加载 `reuse-latest`
- **THEN** 必须存在至少 `1` 个被多个 `Product` 直接绑定的 `Module`
- **AND** 必须存在至少 `1` 个填写了 `capability_key` 的 `Module`
- **AND** 验收时必须可以在 Dashboard、Module Detail 或 Product Detail 中看到非空 `module_reuse_summary / capability_summary`
- **AND** 当前 fixture 不得依赖异步统计表或离线任务才能展示复用反馈

#### Scenario: `reuse-latest-after-binding`

- **WHEN** 加载 `reuse-latest-after-binding`
- **THEN** 该 fixture 必须在 `reuse-latest` 基础上额外体现一次已提交绑定变化
- **AND** 再次读取 `ReuseSummaryRead` 时，必须返回更新后的 `reuse_product_count`、`latest_reuse_at` 或 `capability_summary`
- **AND** 该 fixture 必须作为“读取时反映最新已提交状态”的正式验收证据

### Requirement: 首轮成功会话与阶段完成验收矩阵必须冻结

系统 SHALL 将 `phase06` 的最小验收矩阵冻结为一组可重复复核的单值检查项。

#### Scenario: 首轮成功会话矩阵

- **WHEN** 验收人员执行 `phase06` 首轮主线验收
- **THEN** 必须至少覆盖以下矩阵：
  - `cold-start-empty` → 验证入口判定为 `/onboarding`
  - `in-progress-partial-entry` → 验证回访继续为 `Dashboard -> Continue Onboarding`
  - `completed-unbound` → 验证四类对象最小持久化已完成，但绑定仍可后补
- **AND** 不得把“只完成部分对象录入”判定为首轮成功会话

#### Scenario: 阶段完成矩阵

- **WHEN** 验收人员判断 `phase06` 是否达到阶段完成条件
- **THEN** 必须至少覆盖以下矩阵：
  - `export-ready` → 数据主权导出闭合
  - `backup-verified` → 恢复前提校验成立
  - `reuse-latest-after-binding` → 复用感知读取最新已提交状态
- **AND** 上述三个矩阵必须各自有独立 fixture 证据
- **AND** 不得用单个“大而全” fixture 替代这些独立证据

### Requirement: 验收不得依赖手工补数据

系统 SHALL 明确 `phase06` 联调验收不得依赖手工补 SQL、手工写文件、手工改状态来完成。

#### Scenario: 禁止手工补票

- **WHEN** 验收人员执行 `phase06` 联调验收
- **THEN** 必须通过 `reset_phase06_acceptance.sh` 与 fixture 白名单重复建立环境
- **AND** 不得手工更新 `first_run_state`
- **AND** 不得手工插入绑定关系来替代正式 fixture
- **AND** 不得手工修改备份 `manifest` 来伪造 `backup verified` 成立或失败

### Requirement: 合同一致性验证必须纳入验收矩阵

系统 SHALL 将 Onboarding / Export / Backup / Reuse Summary 的合同一致性验证纳入 `phase06` 验收矩阵，不得只验证功能行为而遗漏合同层一致性。

#### Scenario: OnboardingRead 合同一致性

- **WHEN** 验收人员执行 `phase06` 合同一致性验收
- **THEN** 必须验证 `OnboardingRead`（含 `first_run_state`）的 `.proto -> HTTP DTO -> 前端消费模型` 单值一致
- **AND** 不得只验证 `CreateDraft*` 写入响应而遗漏 `GetFirstRunState` 读取侧合同一致性
- **AND** 该验证必须覆盖 `cold-start / in-progress / completed` 三类 `first_run_state`

#### Scenario: Export / Backup / Reuse Summary 合同一致性

- **WHEN** 验收人员执行 `phase06` 合同一致性验收
- **THEN** 必须验证 Export、Backup、Reuse Summary 三类能力的 `.proto -> HTTP DTO -> 前端消费模型` 单值一致
- **AND** 不得只验证其中部分能力而遗漏其他
- **AND** 每类能力的合同一致性必须有独立 fixture 证据

#### Scenario: `backup_snapshot` 读取侧合同独立性

- **WHEN** 验收人员验证 Backup 合同一致性
- **THEN** 必须独立验证 `GetBackupSnapshot`（读取侧）的合同一致性
- **AND** 不得以 `CreateInstanceBackup`（写入侧）响应附带代替读取侧合同验证
- **AND** 读取侧合同验证必须由 `BackupWrite` owner 内的正式 `read/verify` 子路径或与上游冻结口径一致的等价读取承接位承接

### Requirement: 备份失败 fixture 结果语义冻结

系统 SHALL 将三类备份失败 fixture 冻结为可重复复现的结果语义，但不提前冻结其底层实现方式、存储介质或环境目录细节。

#### Scenario: 备份失败 fixture 结果要求

- **WHEN** 接手者实现 `backup-manifest-missing / backup-coverage-incomplete / backup-schema-mismatch` fixture
- **THEN** `backup-manifest-missing` 必须稳定复现“manifest 缺失 / 不可解析”
- **AND** `backup-coverage-incomplete` 必须稳定复现“覆盖矩阵不完整”
- **AND** `backup-schema-mismatch` 必须稳定复现“schema / version 前提不可校验”
- **AND** 每类失败 fixture 必须可通过 `--fixture <name>` 重复加载并稳定复现失败语义
- **AND** 当前阶段不得把失败 fixture 的底层实现方式冻结为 SQL 模拟、文件删除、环境变量或其他单一机制

### Requirement: 导出/备份验收产物清理边界冻结

系统 SHALL 将 `phase06` 验收过程中产生的导出/备份验收产物清理边界冻结为单值结论，但不提前冻结产物存储介质、目录或具体表名。

#### Scenario: 清理边界

- **WHEN** `reset_phase06_acceptance.sh` 执行清空
- **THEN** 必须清理当前阶段验收过程中生成的导出 / 备份产物及其关联元数据
- **AND** 不得清理 canonical 表数据（由既有 mainline reset 脚本承接）
- **AND** 不得清理 `schema_migrations` 记录
- **AND** 当前阶段不得把清理边界解释为已经冻结了产物必须落在数据库表、文件系统目录或其他单一介质上

## MODIFIED Requirements

### Requirement: `phase06` 验收环境的建立职责

`phase06` 的验收环境在当前子任务中 SHALL 从抽象“后续需要 fixture”推进为“必须通过统一 reset 脚本 + 白名单 fixture 反复建立”，并显式继承 `phase05` 的 reset / seed 模式。

#### Scenario: 验收入口职责收口

- **WHEN** 接手者讨论 `phase06` 验收环境如何落地
- **THEN** 必须先通过统一 reset 入口恢复 baseline 或加载 fixture
- **AND** 验收环境必须同时支撑真实 API 联调与浏览器侧重复复核
- **AND** 不得继续停留在“到时再手工准备一批数据”的模糊表述

## REMOVED Requirements

### Requirement: 通过一次性手工补数据完成 `phase06` 验收

**Reason**: 这种做法无法支撑 `first_run_state`、回访继续、`backup verified` 与复用感知最新状态的重复验证，也会直接破坏当前阶段“可重复建立验收环境”的 DoD。

**Migration**: 后续统一改为通过 `reset_phase06_acceptance.sh` 与白名单 fixture 建立环境；所有关键状态、失败路径与阶段完成结论都必须回落到正式 fixture 证据。
