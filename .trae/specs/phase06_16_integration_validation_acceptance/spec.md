# Phase06-16 联调、验证与验收 Spec

## Why

`phase06-11` 已冻结联调环境、fixture 与恢复基线，`phase06-12` 已冻结 `Onboarding + Sovereignty + Reuse` 正式规格正文，`phase06-13` 已冻结 `.proto` 合同主线，`phase06-14 / 15` 已分别推进后端与前端主线实现。但当前仍缺少一份把真实环境、真实主线、真实 fixture 与验收结论收成同一闭环的联调验收规格，无法证明 `phase06` 最小主线已经可重复走通、失败语义未被折叠、且没有破坏 `phase05` 已交付边界。

`phase06-16` 的目标，是冻结 `phase06` 的真实联调入口、最小验收矩阵、证据记录方式与收口规则，让后续验收不再停留在“本地能跑通一次”，而是形成可重复复核的阶段交付结论。

## What Changes

- 冻结 `phase06-16` 必须基于真实前端、真实后端、真实数据库与 `phase06-11` 已冻结的 `reset_phase06_acceptance.sh` + fixture 白名单执行
- 冻结首轮录入、部分补全、导出、备份、恢复前提与复用反馈的最小联调验收矩阵
- 冻结 cold-start、`in_progress` 回访继续、`completed` 成功会话与根级入口行为的验收口径
- 冻结 `Export` 覆盖矩阵、`Backup` 校验语义与 `module_reuse_summary / capability_summary` 新鲜度的验收边界
- 冻结局部错误、失败语义与重复执行前提的验收规则，避免把局部失败折叠为整页失败或泛化错误
- 冻结与 `phase05` 已交付 `Dashboard + Feedback` 主线兼容性的回归检查项
- 冻结 `phase06-16` 必须产出单一验收记录，显式记录环境、步骤、结果、问题、复测与 DoD 达成情况
- **BREAKING**：`phase06` 联调验收不得再依赖手工补 SQL、手工改 `first_run_state`、手工伪造导出覆盖矩阵或手工构造备份校验结果

## Impact

- Affected specs:
  - `phase06_11_design_integration_acceptance_fixtures_restore_baseline`
  - `phase06_12_onboarding_sovereignty_reuse_formal_spec`
  - `phase06_13_land_minimal_proto_contract_mainline`
  - `phase06_14_implement_backend_data_script_mainline`
  - `phase06_15_implement_frontend_mainline`
  - `phase05_14_dashboard_feedback_integration_validation_acceptance`
- Affected code:
  - `database/scripts/reset_phase06_acceptance.sh`
  - `database/seeds/seed_phase06_acceptance_baseline.sql`
  - `database/seeds/seed_phase06_fixture_*.sql`
  - `backend/internal/onboarding/`
  - `backend/internal/export/`
  - `backend/internal/backup/`
  - `backend/internal/reusesummary/`
  - `frontend/src/routes/index.tsx`
  - `frontend/src/routes/onboarding.tsx`
  - `frontend/src/features/onboarding/`
  - `frontend/src/features/dashboard/`
  - `frontend/src/features/reuse-summary/`
  - `frontend/src/features/product-registry/`
  - `frontend/src/features/repository-binding/`
  - `frontend/src/features/module-registry/`
  - `frontend/src/features/decision-center/`
  - 后续 `phase06_16_integration_validation_acceptance/acceptance_report.md`（新增）

## ADDED Requirements

### Requirement: 联调验收环境必须复用 phase06-11 冻结的真实基线

系统 SHALL 将 `phase06-16` 的联调验收环境冻结为复用 `phase06-11` 已冻结的 `reset_phase06_acceptance.sh`、默认基线与 fixture 白名单，在真实前端、真实后端、真实数据库上执行，不得并列引入 mock、假页面或手工 SQL 替代路径。

#### Scenario: 真实验收入口

- **WHEN** 执行 `phase06-16` 联调验收
- **THEN** 必须使用真实前端根级入口、`/dashboard`、`/onboarding`、canonical detail 页面与 Dashboard 动作入口完成验收
- **AND** 必须使用真实后端 `/api/onboarding/*`、`/api/dashboard/export`、`/api/dashboard/backup`、`/api/reuse-summary/*` 或等价正式接口
- **AND** 必须复用 `database/scripts/reset_phase06_acceptance.sh` 建立默认基线与 fixture
- **AND** 不得通过 mock adapter、手工拼 JSON、手工 SQL 或浏览器 DevTools 篡改结果替代联调

#### Scenario: 运行时前置条件核对

- **WHEN** 进入正式联调步骤前
- **THEN** 必须显式核对并记录以下前置条件：
  - 前端服务可访问根级入口、`/dashboard` 与 `/onboarding`
  - 后端服务与 `/api` 路由可用
  - `reset_phase06_acceptance.sh` 默认恢复与 fixture 恢复可重复执行
  - `.proto`、HTTP DTO 与前端消费模型保持单值一致
  - 导出 / 备份验收产物目录可重复清理与重建
- **AND** 任一前置条件不成立时，不得直接宣告 `phase06-16` 通过

### Requirement: 首轮录入、回访继续与成功会话必须完成最小联调验收

系统 SHALL 将 `Onboarding` 最小主线的联调验收矩阵冻结为覆盖 cold-start、`in_progress` 回访继续与 `completed` 成功会话三类正式状态，显式证明首轮录入主线与根级入口守卫已形成闭环。

#### Scenario: cold-start 根级入口验收

- **WHEN** 使用 `reset_phase06_acceptance.sh --fixture cold-start-empty` 建立冷启动环境
- **THEN** 根级默认进入路径必须回落到 `/onboarding`
- **AND** `/dashboard` 中不得展示与 `completed` 用户相同的首轮状态语义
- **AND** `Start Onboarding` 必须是稳定可见的首轮主 CTA
- **AND** 当前 fixture 不得伪造任何已完成首轮对象记录

#### Scenario: `in_progress` 回访继续验收

- **WHEN** 使用 `reset_phase06_acceptance.sh --fixture in-progress-partial-entry` 建立部分录入环境
- **THEN** 根级默认进入路径必须为 `/dashboard`
- **AND** Dashboard 必须展示 `Continue Onboarding`
- **AND** 进入 `/onboarding` 后必须自动定位到第一个未完成步骤
- **AND** 从 `Onboarding -> canonical detail -> /onboarding` 的回流路径必须恢复对应步骤，而不是丢失回访语义

#### Scenario: 首轮成功会话验收

- **WHEN** 使用默认恢复或从 cold-start 真实完成一次首轮录入主线
- **THEN** 必须显式验证 `Product / Repository / Module / Decision` 四类对象都已落持久化记录
- **AND** `first_run_state` 必须进入 `completed`
- **AND** 成功会话不允许以“只完成部分对象”或“只创建草稿但未持久化”代替
- **AND** 成功会话走通后必须可重复再次执行，不得依赖一次性脏环境

### Requirement: 数据主权路径必须完成 Export / Backup 联调验收

系统 SHALL 将 `Export` 与 `Backup` 的联调验收矩阵冻结为覆盖读取、触发、覆盖矩阵校验与恢复前提校验的最小主线，显式证明数据主权路径真实可验证。

#### Scenario: Export 覆盖矩阵验收

- **WHEN** 使用 `export-ready` 或等价已绑定 fixture 执行导出验收
- **THEN** 必须验证 `ExportSnapshot` 可读取
- **AND** 必须验证正式导出动作可执行
- **AND** 导出结果必须同时覆盖核心主实体与绑定 / 关联关系，而不是只包含主实体
- **AND** 至少要验证以下集合进入导出覆盖矩阵：
  - `products`
  - `modules`
  - `releases`
  - `repositories`
  - `decisions`
  - `decision_links`
  - `product_modules`
  - `product_repositories`
  - `module_repositories`

#### Scenario: Backup Verified 成功语义验收

- **WHEN** 使用 `backup-verified` fixture 执行备份验收
- **THEN** 必须验证 `BackupSnapshot` 可读取
- **AND** 必须验证备份动作可执行并生成可读取产物
- **AND** 只有 `manifest` 可读、覆盖矩阵完整、`schema / version` 前提可校验时，才允许判定 `backup verified`
- **AND** 不得把“文件写出成功”直接等价为 `backup verified`

#### Scenario: Backup 失败语义验收

- **WHEN** 分别使用 `backup-manifest-missing`、`backup-coverage-incomplete`、`backup-schema-mismatch` fixture 验证失败路径
- **THEN** 必须分别看到单值失败语义：
  - `manifest_missing`
  - `coverage_incomplete`
  - `schema_mismatch`
- **AND** 失败语义必须在 `.proto -> HTTP DTO -> 前端展示` 边界保持单值一致
- **AND** 不得把三类失败折叠为泛化的“备份失败”

### Requirement: 复用反馈路径必须验证最新已提交状态

系统 SHALL 将 `module_reuse_summary / capability_summary` 的联调验收冻结为覆盖 Dashboard、Module Detail、Product Detail 三类正式挂接位，并显式验证读取结果反映最新已提交状态。

#### Scenario: Dashboard 复用快照验收

- **WHEN** 使用 `reuse-latest` fixture 进入 Dashboard
- **THEN** `Asset Feedback` 内的 `Reuse Snapshot` 子区域必须独立展示自己的 `loading / ready / empty / error` 状态
- **AND** `module_reuse_summary` 与 `capability_summary` 必须可见且可解释
- **AND** 复用读取失败不得把整页打回 `page-error`

#### Scenario: Detail 页复用反馈验收

- **WHEN** 使用 `reuse-latest` 或 `reuse-latest-after-binding` fixture 进入 Module Detail 与 Product Detail
- **THEN** 每个详情页都必须只通过一个页面级 `ReuseSummaryRead` query 展示复用摘要
- **AND** 复用摘要不得承接绑定写入、解绑写入或候选筛选逻辑
- **AND** 挂接位必须继续对齐 `phase06-12 / 15` 已冻结的页面位置

#### Scenario: 最新状态新鲜度验收

- **WHEN** 使用 `reuse-latest-after-binding` fixture 或等价真实已提交变化执行复测
- **THEN** `module_reuse_summary / capability_summary` 必须反映最新已提交状态
- **AND** 不得继续展示变更前的旧统计结果
- **AND** 当前阶段不得依赖异步统计表或离线任务才能看到更新后的复用反馈

### Requirement: 局部错误、重复执行与 phase05 兼容性必须纳入验收

系统 SHALL 将局部错误、重复执行前提与 `phase05` 已交付 `Dashboard + Feedback` 主线兼容性冻结为本阶段必验项，避免 phase06 实现破坏既有已交付边界。

#### Scenario: 局部错误边界验收

- **WHEN** 复用读取、导出读取或备份读取中的某一个子查询失败
- **THEN** 失败必须局限在对应子区域
- **AND** 已成功的 Dashboard 既有读区块不得被拖垮为整页错误
- **AND** 局部重试入口必须继续可用

#### Scenario: 重复执行前提验收

- **WHEN** 验收人员重复执行默认恢复、fixture 恢复、导出动作或备份动作
- **THEN** 环境必须可重复重置与复测
- **AND** 重复执行不得依赖手工清库、手工删文件或手工回填状态
- **AND** 验收结论必须建立在可重复运行前提上，而不是一次性偶然成功

#### Scenario: `phase05` 兼容性验收

- **WHEN** 执行 `phase06-16` 联调验收
- **THEN** 必须继续验证 `overview / feedback-signals / recent-activities` 三个既有 Dashboard 查询仍按 `phase05` 口径工作
- **AND** `phase06` 新增的 `Onboarding CTA / Export / Backup / Reuse Snapshot` 不得破坏 `phase05` 已交付的状态模型、返回路径与局部错误边界
- **AND** 若发现回归，不得宣告 `phase06-16` 通过

### Requirement: phase06-16 必须产出单一验收记录并完成问题收口

系统 SHALL 将 `phase06-16` 的联调结果冻结为单一验收记录，显式记录环境、步骤、结果、问题、复测与 DoD 达成情况；联调中发现的问题必须在当前阶段显式收口。

#### Scenario: 验收记录结构

- **WHEN** `phase06-16` 实施完成
- **THEN** 必须产出单一 `acceptance_report.md`
- **AND** 该报告至少包含：
  - 验收环境与前置条件
  - fixture 与恢复入口
  - 首轮录入 / 回访继续 / 成功会话验收结果
  - Export / Backup / backup verified 验收结果
  - Reuse Summary 最新状态验收结果
  - `phase05` 兼容性复核结果
  - 问题、修复与复测
  - DoD 达成情况
- **AND** 不得把同一结论拆散到多个并列验收文档

#### Scenario: 问题收口规则

- **WHEN** 联调中发现阻断问题或语义偏差
- **THEN** 必须显式记录问题、级别、定位、修复与复测结论
- **AND** 若问题未修复，不得宣告 `phase06-16` 通过
- **AND** 不得把本阶段阻断问题默认为“留到下一阶段再看”

## MODIFIED Requirements

### Requirement: Phase06 验收矩阵必须从基线设计升级为真实联调结论

系统 SHALL 在 `phase06-11` 已冻结 fixture 与恢复基线的前提下，将 `phase06-16` 升级为真实联调、真实验证与真实验收结论阶段，而不是停留在“验收环境可建立”的设计层。

#### Scenario: 从基线到验收结论的升级

- **WHEN** `phase06-16` 进入执行
- **THEN** 必须基于 `phase06-11` 已冻结的 11 个 fixture 白名单执行真实联调
- **AND** 必须形成通过 / 不通过的阶段结论
- **AND** 不得只验证脚本存在、fixture 可加载就判定本阶段完成

## REMOVED Requirements

### Requirement: 以单次手工演示替代正式联调验收
**Reason**: 单次手工演示无法证明首轮成功会话、数据主权路径、复用反馈路径与失败语义可重复验证，也无法支撑 `phase06` 的正式收口结论。
**Migration**: 统一改为基于 `reset_phase06_acceptance.sh`、fixture 白名单、真实前后端与单一 `acceptance_report.md` 的重复联调验收流程。
