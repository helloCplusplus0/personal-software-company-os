# phase08-09 落实 Dashboard review 入口、双路径 review 会话与统一动作承接 Spec

## Why

`phase08-08` 已经把 review 合同、后端承接位、前端 owner 与最小页面 enablement 落到代码树里，但当前还缺一份面向用户可见行为的集成规格，来正式冻结 `Dashboard -> Daily / Weekly Review` 的真实入口、双会话差异与统一动作承接边界。
如果这一步不继续收口，后续实现很容易把 daily / weekly 重新做成“同一套页面换标题”，或把 `phase05 / phase06` 既有读模型消费重新散回页面级临时编排，最终削弱 `phase08` 的经营回路语义。

## What Changes

- 冻结 `DashboardPrimaryActionPanel` 作为当前阶段唯一正式 review 入口 caller 的集成边界
- 冻结 `Daily Review` 与 `Weekly Review` 的真实会话差异、区块顺序与最小消费模型
- 冻结 `phase05 / phase06` 已交付读模型在 review 会话中的正式消费位置
- 冻结 `ReviewPageShell + ReviewActionFooter + useReviewAction()` 的统一动作承接语义
- 冻结 `Dashboard -> /reviews/* -> canonical page / review record` 的最小关键路径验收口径

## Impact

- Affected specs:
  - `phase08_05_design_dashboard_review_entry_page_interaction_flow`
  - `phase08_06_design_review_frontend_read_write_owner_state_flow`
  - `phase08_07_design_review_backend_contract_service_data_handoff`
  - `phase08_08_land_review_contract_backend_frontend_owner_enablement`
  - `phase05_10_dashboard_feedback_formal_spec`
  - `phase05_13_dashboard_frontend_mainline`
  - `phase06_09_design_reuse_summary_query_dashboard_detail_integration`
  - `phase06_15_implement_frontend_mainline`
- Affected code:
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/review/components/review-page-shell.tsx`
  - `frontend/src/features/review/components/review-action-footer.tsx`
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/features/review/data/use-daily-review-read.ts`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/dashboard/components/feedback-signal-card.tsx`
  - `frontend/src/features/dashboard/components/recent-activity-item-card.tsx`
  - `frontend/src/features/dashboard/components/reuse-snapshot-section.tsx`

## ADDED Requirements

### Requirement: Dashboard 必须正式承接 review 双入口，而不是停留在 enablement 占位

系统 SHALL 将 `DashboardPrimaryActionPanel` 冻结为当前阶段唯一正式 review 入口 caller，使用户能够从 `/dashboard` 真实进入 `Daily Review` 与 `Weekly Review` 两条会话路径，而不是只保留文案占位或测试入口。

#### Scenario: Dashboard review 双入口的正式进入方式

- **WHEN** 用户进入 `/dashboard`
- **THEN** 标题行动区必须稳定呈现 `Daily Review` 与 `Weekly Review` 两个显式入口
- **AND** 两个入口都必须导航到 `/reviews/daily` 与 `/reviews/weekly`
- **AND** 两个入口都必须继续透传 `buildDashboardSourceParams('empty-state')`
- **AND** 当前阶段不得把 review 入口继续藏在 `Current Focus / Asset Feedback / Recent Activity` 区块内部按钮里

#### Scenario: Dashboard review 入口与既有模块的最小整合

- **WHEN** review 入口落到 Dashboard 首页
- **THEN** `DashboardHomePage` 仍然只承接总览页与入口宿主，不承担 review read owner 或 review action owner
- **AND** `OnboardingCtaButton`、`SovereigntyPanel` 与既有 Dashboard 区块必须继续保持独立职责
- **AND** review 入口不得把 Dashboard 首屏改造成第二套工作台或 tab 容器

### Requirement: Daily Review 必须成为独立于 Weekly Review 的真实短会话

系统 SHALL 将 `Daily Review` 冻结为“当下焦点处理会话”，正式消费 `phase05` 反馈与决策事实的最小组合，而不是 `Weekly Review` 的裁剪版或同模板页面。

#### Scenario: Daily Review 的最小读模型与区块顺序

- **WHEN** 用户进入 `/reviews/daily`
- **THEN** 页面至少必须按以下顺序正式消费并展示：
  - `current focus`
  - `pending decisions`
  - `representative signals`
  - `ReviewActionFooter`
- **AND** `current focus` 必须继续来自 review context 中承接的 `current_focus_signals`
- **AND** `pending decisions` 必须继续来自 `Decision` canonical facts 的 `proposed` 决策最小摘要
- **AND** `representative signals` 只能作为补充缺口输入，不得反向升级成第二条主队列

#### Scenario: Daily Review 不得冒充 Weekly Review

- **WHEN** 审查 `DailyReviewPage`
- **THEN** 页面不得把 `overview / recent activity / reuse snapshot / capability summary` 升级为正式主区
- **AND** 页面不得通过复用同一套配置把 daily / weekly 渲染成除标题外无差异的双路径

### Requirement: Weekly Review 必须正式消费 `phase05 / phase06` 周期整理读模型

系统 SHALL 将 `Weekly Review` 冻结为“周期整理与归纳会话”，正式消费 `phase05` 的 overview / recent activity / representative signals，以及 `phase06` 的 reuse snapshot 读模型，而不是把这些读取继续停留在 Dashboard 内部。

#### Scenario: Weekly Review 的最小读模型与区块顺序

- **WHEN** 用户进入 `/reviews/weekly`
- **THEN** 页面至少必须按以下顺序正式消费并展示：
  - `overview`
  - `recent activity`
  - `representative signals`
  - `reuse snapshot`
  - `ReviewActionFooter`
- **AND** `overview` 必须继续消费 `DashboardOverview`
- **AND** `recent activity` 必须继续消费 `RecentActivityItem`
- **AND** `representative signals` 必须继续消费 `asset_feedback_summary.representative_signals`
- **AND** `reuse snapshot` 必须继续消费 `module_reuse_summary / capability_summary`

#### Scenario: Weekly Review 的 reuse snapshot 消费边界

- **WHEN** `WeeklyReviewPage` 承接 `reuse snapshot`
- **THEN** 页面只能消费 `useWeeklyReviewRead()` 暴露的只读结果与 retry 能力
- **AND** 页面不得直接 import `useReuseSummaryRead()`、`queryClient` 或 review query key 重新编排第二套读取主线
- **AND** `reuse snapshot` 的局部失败只能维持在该区块内，不得把整页回退为 page error

### Requirement: 两条 review 会话必须继续通过统一 action owner 承接动作

系统 SHALL 将 daily / weekly review 的动作出口继续收敛到 `ReviewActionFooter + useReviewAction()`，形成统一的 handoff / next-step result 承接位，而不是让页面各自长出第二套 action 编排。

#### Scenario: review 完成区的统一动作语义

- **WHEN** 用户在 `Daily Review` 或 `Weekly Review` 点击底部动作
- **THEN** 纯 `decision / product / module / repository` handoff 路径必须继续只返回 success envelope，并导航到既有 canonical 页面
- **AND** `next-step result` 路径必须继续正式调用 `ReviewService.SubmitReviewResult`
- **AND** 所有成功 envelope 都必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`
- **AND** 页面只允许消费 action owner 结果执行 `navigate()` 与 toast，不得自行补写第二套 mutation / invalidation

#### Scenario: 统一动作承接不抹平双会话差异

- **WHEN** 两条 review 会话复用同一套 `ReviewActionFooter`
- **THEN** footer 只允许复用动作承接语义与交互骨架
- **AND** daily / weekly 的 `reviewKind`、页面说明文案与区块语义仍必须保持区分
- **AND** 不得为了复用 footer 而把 daily / weekly 都回退成“同一套完成定义”

### Requirement: `phase08-09` 验收必须覆盖双路径会话与既有读模型消费

系统 SHALL 将 `phase08-09` 的验收口径冻结为“Dashboard 入口 + 双路径会话 + 统一动作承接 + `phase05 / phase06` 读模型消费”的组合验收，而不是只验证路由存在或前端可构建。

#### Scenario: 最小关键路径验收

- **WHEN** 验证 `phase08-09` 是否达成 DoD
- **THEN** 至少必须覆盖：
  - `/dashboard -> /reviews/daily`
  - `/dashboard -> /reviews/weekly`
  - `Daily Review` 真实显示 `current focus / pending decisions / representative signals`
  - `Weekly Review` 真实显示 `overview / recent activity / representative signals / reuse snapshot`
  - review 页面通过统一动作承接位进入 `Decision / Product / Module / Repository` canonical path
  - `next-step result` 路径继续命中 `ReviewService.SubmitReviewResult`
- **AND** 不得只用 `npm run build` 冒充双会话验收

## MODIFIED Requirements

### Requirement: `phase08-08` 的“最小 enablement”在 `phase08-09` 中必须升级为真实会话集成

自 `phase08-09` 起，系统 SHALL 不再只满足“可进入 route、可读取 review context、可触发最小完成动作”的 enablement，而必须进一步满足“Dashboard 真实入口可用、daily / weekly 会话不冒充双路径、`phase05 / phase06` 读模型已正式消费进 review”的集成要求。

#### Scenario: review enablement 不再停留在最小占位

- **WHEN** 审查 `phase08-09` 实现结果
- **THEN** 必须同时看到 Dashboard 入口、daily / weekly 差异化主体与统一动作承接都已落到真实页面行为
- **AND** 不得把 `phase08-08` 的最小页面壳层、空列表或占位数据视为 `phase08-09` 完成

### Requirement: review 页面消费边界必须继续服从 owner 单值化

自 `phase08-09` 起，系统 SHALL 进一步要求 review 页面只消费 review slice 暴露的稳定 read / action owner 结果；即便为了 UI 复用 Dashboard 组件，也不得重新把 query、retry、mutation 或 query key 编排拉回页面层。

#### Scenario: 复用 Dashboard 展示组件时的 owner 边界

- **WHEN** `DailyReviewPage` 或 `WeeklyReviewPage` 复用 `FeedbackSignalCard / RecentActivityItemCard / ReuseSnapshotSection`
- **THEN** 页面只能向这些展示组件传只读 props 与 owner 暴露的回调
- **AND** 不得因为复用展示组件而直接 import 底层 canonical hooks、`queryClient`、query key 或 page-local mutation

## REMOVED Requirements

### Requirement: daily / weekly review 可以先通过同一套主体内容落地，再在后续阶段细分

**Reason**: 这会直接削弱 `phase08-09` 的双路径会话目标，让 `Daily Review` 与 `Weekly Review` 退化为“同模板不同标题”，并无法证明 `phase05 / phase06` 读模型已经被正式消费进 review。
**Migration**: 当前阶段必须直接冻结两条会话的最小区块差异、读模型差异与统一动作承接位；若后续需要新增更深层整理能力，必须在新的 `/spec` 中单独扩展。
