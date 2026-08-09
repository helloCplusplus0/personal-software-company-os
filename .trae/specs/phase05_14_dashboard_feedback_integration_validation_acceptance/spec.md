# Phase05-14 Dashboard + Feedback 联调、验证与验收 Spec

## Why

`phase05-09` 已冻结 Dashboard 验收环境、基线与 fixture，`phase05-10` 已冻结 `Dashboard + Feedback` 正式规格正文，`phase05-11` 已落地 `.proto` 合同主线，`phase05-12` 与 `phase05-13` 已分别落地后端/数据主线与前端主线。但截至当前，还没有把这些交付物串成一份可重复复核的联调验收方案，明确证明 `Dashboard + Feedback` 最小主线已经真实可运行、可重复验证、且当前阶段发现的问题能被收口。

`phase05-14` 的目标是冻结真实联调入口、最小验收矩阵、验收记录方式与问题收口规则，让后续验收不再停留在“本地看起来能用”，而是形成可重复复核的阶段交付结论。

## What Changes

- 冻结 `phase05-14` 必须基于真实前端、真实后端、真实数据库与 `phase05-09` 验收基线执行，不允许 mock、伪页面或手工 SQL 替代
- 冻结 `Dashboard` 空状态、有数据状态、局部错误状态与 `Recent Activity` 展示的最小验收矩阵
- 冻结 `overview / feedback-signals / recent-activities` 三类聚合读取的联调验证范围与成功语义
- 冻结从 Dashboard 到 `Product / Repository / Module / Decision` 的跳转、主动返回与多跳返回路径验收矩阵
- 冻结 `phase05-14` 必须产出单一验收记录，明确环境、步骤、结果、问题与 DoD 达成情况
- 冻结联调中发现的问题必须在当前阶段显式收口，不得把隐性阻断留到下一阶段

## Impact

- Affected specs:
  - `phase05_09_design_dashboard_acceptance_baseline_fixtures`
  - `phase05_10_dashboard_feedback_formal_spec`
  - `phase05_11_dashboard_feedback_proto_mainline`
  - `phase05_12_dashboard_feedback_backend_data_mainline`
  - `phase05_13_dashboard_frontend_mainline`
- Affected code:
  - `database/scripts/reset_dashboard_acceptance.sh`
  - `database/seeds/seed_dashboard_acceptance_baseline.sql`
  - `database/seeds/seed_dashboard_fixture_*.sql`
  - `backend/internal/dashboard/`
  - `backend/internal/platform/router.go`
  - `frontend/src/routes/dashboard.tsx`
  - `frontend/src/features/dashboard/`
  - 后续 `phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`（新增）

## ADDED Requirements

### Requirement: 联调验收环境必须复用 phase05-09 冻结的真实基线

系统 SHALL 将 `phase05-14` 联调验收环境冻结为复用 `phase05-09` 已冻结的 `reset_dashboard_acceptance.sh`、验收基线种子与 fixture 入口，在真实前端、真实后端、真实数据库上执行，不得并列引入 mock 数据、伪接口或手工 SQL 替代路径。

#### Scenario: 验收环境真实入口

- **WHEN** 执行 `phase05-14` 联调验收
- **THEN** 必须使用真实前端 `/dashboard` 作为浏览器验收入口
- **AND** 必须使用真实后端 `/api/dashboard/overview`、`/api/dashboard/feedback-signals`、`/api/dashboard/recent-activities` 三个 endpoint
- **AND** 必须复用 `database/scripts/reset_dashboard_acceptance.sh` 建立基线与 fixture
- **AND** 不得使用 mock adapter、假页面、手工拼 JSON 或手工 SQL 替代联调

#### Scenario: 运行时前置条件核对

- **WHEN** 进入正式联调步骤前
- **THEN** 必须显式核对并记录以下前置条件：
  - 前端开发服务可访问 `/dashboard`
  - 后端服务与 `/api` 路由可用
  - `frontend/.env` 处于真实 API 模式
  - 验收重置脚本与 fixture 可重复执行
  - `.proto`、后端 HTTP envelope 与前端消费模型保持单值一致
- **AND** 任一前置条件不成立时，不得直接宣告验收通过

### Requirement: Dashboard 状态矩阵必须完成最小联调验收

系统 SHALL 将 `Dashboard + Feedback` 的最小联调验收矩阵冻结为至少覆盖空状态、有数据状态、局部错误状态与最近活动展示四类核心页面状态，且每类状态都必须可由正式基线或 fixture 重复建立。

#### Scenario: 空状态验收

- **WHEN** 使用 `reset_dashboard_acceptance.sh --fixture empty-system` 建立空系统
- **THEN** `/dashboard` 必须进入冷启动空状态
- **AND** `DashboardOverview` 必须展示全零计数成功态，而不是整页错误
- **AND** 主 CTA 必须指向 `Module Registry / Create`
- **AND** `Current Focus`、`Asset Feedback` 与 `Recent Activity` 必须展示各自成功空态或受控空态

#### Scenario: 有数据状态验收

- **WHEN** 使用默认恢复或 `--fixture recent-activities` 建立最小有数据基线
- **THEN** `/dashboard` 必须展示非零概览聚合
- **AND** 至少一类反馈信号或中性态必须可见
- **AND** `Recent Activity` 必须展示真实活动项或受控空列表成功态
- **AND** 页面不得因为附属聚合局部失败之外的原因退化为整页错误

#### Scenario: 局部错误状态验收

- **WHEN** 使用 `phase05-12` 已落地的局部错误模拟入口触发 `feedback-signals` 或 `recent-activities` 失败
- **THEN** `/dashboard` 必须保持整页 `ready`
- **AND** 对应区块必须展示局部错误与局部重试入口
- **AND** `overview` 成功区块不得被局部错误拖垮
- **AND** 主 CTA 不得在数据不完整时伪造反馈型决策

#### Scenario: Recent Activity 验收

- **WHEN** 使用 `--fixture recent-activities` 建立活动基线
- **THEN** `Recent Activity` 必须展示真实活动列表
- **AND** 活动类型映射必须覆盖 `module / release / product / repository / decision / product_module_binding / product_repository_binding / module_repository_binding`
- **AND** 列表顺序必须按后端返回结果稳定展示
- **AND** 空列表与错误态必须与 `phase05-10 / 13` 已冻结语义一致

### Requirement: 聚合读取与反馈信号最小路径必须可重复复核

系统 SHALL 将 `overview / feedback-signals / recent-activities` 三类聚合读取的最小联调路径冻结为可重复复核的验收步骤，显式证明后端聚合、HTTP 包络与前端消费编排已经形成单一闭环。

#### Scenario: DashboardOverview 最小验收路径

- **WHEN** 验收 `GET /api/dashboard/overview`
- **THEN** 必须验证其返回包络为 `{"overview": ...}` 成功语义
- **AND** 前端 `DashboardStatBar` 必须正确消费 `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count`
- **AND** 零计数与非零计数都必须被视为成功态

#### Scenario: FeedbackSignals 最小验收路径

- **WHEN** 验收 `GET /api/dashboard/feedback-signals`
- **THEN** 必须验证其返回包络同时包含 `current_focus_signals` 与 `asset_feedback_summary`
- **AND** `Current Focus` 与 `Asset Feedback` 必须共享同一次读取结果
- **AND** CTA 命中、代表性缺口项与资产覆盖计数必须与正式规格正文一致

#### Scenario: RecentActivities 最小验收路径

- **WHEN** 验收 `GET /api/dashboard/recent-activities`
- **THEN** 必须验证其返回包络为 `{"activities": [...]}` 成功语义
- **AND** 前端必须按真实活动结果展示 `Recent Activity`
- **AND** 空数组必须作为成功空态消费，不得映射为错误

### Requirement: Dashboard 到四类 canonical owner 的跳转与返回路径必须完成验收

系统 SHALL 将 Dashboard 到 `Product / Repository / Module / Decision` 四类 canonical owner 的跳转与返回路径冻结为本阶段必验项，覆盖直接返回、多跳返回与来源上下文恢复。

#### Scenario: Dashboard 直接跳转返回

- **WHEN** 用户从 Dashboard 直接进入四类 List 或 Detail 页面
- **THEN** 目标页必须携带 `fromDashboard=true / dashboardSection / dashboardReturnTo=/dashboard`
- **AND** 页面必须提供“返回 Dashboard”入口
- **AND** 主动返回后必须落回 `/dashboard`

#### Scenario: Dashboard -> List -> Detail 多跳返回

- **WHEN** 用户走 `Dashboard -> List -> Detail`
- **THEN** Detail 页必须同时保留页面原生 `fromList` 上下文与 Dashboard 外层来源上下文
- **AND** 从 Detail 返回 List 必须恢复列表原生来源
- **AND** 从 List 返回 Dashboard 必须继续恢复 Dashboard 外层来源
- **AND** 不得在列表筛选变化、详情跳转或返回列表后丢失 Dashboard 返回能力

#### Scenario: Dashboard 来源区块恢复

- **WHEN** 用户从目标页主动返回 Dashboard
- **THEN** 必须通过一次性 navigation state 承接 `dashboardSection`
- **AND** 落回 `/dashboard` 后必须恢复对应来源区块的视图上下文
- **AND** 不得把该一次性状态提升为新的持久来源事实源

### Requirement: phase05-14 必须产出单一验收记录并收口问题

系统 SHALL 将 `phase05-14` 的验收结果冻结为单一验收记录，显式记录环境、步骤、结果、问题与 DoD 达成情况；联调中发现的问题必须在当前阶段显式收口，不得遗留隐性阻断。

#### Scenario: 验收记录结构

- **WHEN** `phase05-14` 实施完成
- **THEN** 必须产出单一 `acceptance_report.md`
- **AND** 该报告至少包含：
  - 验收环境与前置条件
  - 最小主线端到端验收结果
  - Dashboard 状态矩阵验收结果
  - 跳转与返回路径验收结果
  - 合同与正式规格一致性核对
  - 联调中发现的问题与修复
  - 阶段收口结论
  - DoD 达成情况
- **AND** 不得把同一结论拆散到多个并列验收文档

#### Scenario: 问题收口规则

- **WHEN** `phase05-14` 联调中发现问题
- **THEN** 必须显式记录问题、级别、定位与收口结果
- **AND** 若问题已修复，必须记录复测通过结论
- **AND** 若问题未修复，不得宣告 `phase05-14` 通过
- **AND** 不得把阻断问题默认为“留待下一阶段优化”
