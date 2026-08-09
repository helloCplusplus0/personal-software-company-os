# Phase05-13 Dashboard 前端主线实现 Spec

## Why

`phase05-05` 已冻结 Dashboard 前端页面、路由与组件分层，`phase05-06` 已冻结前端状态模型与交互流，`phase05-10` 已产出 `Dashboard + Feedback` 正式规格正文，`phase05-11` 已落地 `.proto` 合同源，`phase05-12` 已落地后端三个 GET endpoint（`/api/dashboard/overview`、`/api/dashboard/feedback-signals`、`/api/dashboard/recent-activities`）与响应包络。但截至当前，仓库中还没有 Dashboard 前端页面、路由、组件、API 适配层与状态编排代码，既有 canonical 路由也还没有承接 `fromDashboard / dashboardSection / dashboardReturnTo` 来源参数。

`phase05-13` 的目标是把上述已冻结结论推进为仓库内实际存在、可编译、可运行、可被联调验收的前端主线，为 `phase05-14` 联调验收提供可执行的前端交付物。

> 阶段分工约束：本规格只冻结"前端主线如何落地为代码"，不重新冻结页面边界、组件树、状态模型、CTA 矩阵、跳转矩阵与来源参数语义。这些已由 `phase05-05 / 06 / 10` 单值化，本规格承接这些结论推进为实现。

## What Changes

- 新增 `frontend/src/routes/dashboard.tsx` 承接 `/dashboard` 路由
- 新增 `frontend/src/features/dashboard/` feature，承接 pages / components / data / types
- 修改 `frontend/src/routes/__root.tsx` 在一级导航中新增 `Dashboard` 入口
- 修改既有四类 List 路由（`modules / products / repositories / decisions` 的 `index.tsx`）扩展 `validateSearch` 承接 Dashboard 来源参数
- 修改既有四类 Detail 路由（`$moduleId / $productId / $repositoryId / $decisionId`）新增 `validateSearch` 承接 Dashboard 来源参数
- 修改既有三类 Create 路由（`modules / products / repositories` 的 `new.tsx`）扩展 `validateSearch` 承接 Dashboard 来源参数
- 修改既有 List / Detail / Create 页面组件，在 `fromDashboard=true` 时提供"返回 Dashboard"导航
- **BREAKING**：Dashboard 前端实现完成后，`phase05-14` 联调验收必须从 `/dashboard` 进入验证 Dashboard 闭环，不得绕过前端直接 mock

## Impact

- Affected specs:
  - `phase05_05_dashboard_frontend_page_route_component_design`
  - `phase05_06_dashboard_frontend_state_interaction_flow`
  - `phase05_10_dashboard_feedback_formal_spec`
  - `phase05_12_dashboard_feedback_backend_data_mainline`
  - 后续 `phase05-14` 联调与验收
- Affected code:
  - `frontend/src/routes/__root.tsx`（修改：新增 Dashboard 导航）
  - `frontend/src/routes/dashboard.tsx`（新增）
  - `frontend/src/features/dashboard/types.ts`（新增）
  - `frontend/src/features/dashboard/data/api-adapter.ts`（新增）
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`（新增）
  - `frontend/src/features/dashboard/components/dashboard-home-page-shell.tsx`（新增）
  - `frontend/src/features/dashboard/components/dashboard-stat-bar.tsx`（新增）
  - `frontend/src/features/dashboard/components/current-focus-section.tsx`（新增）
  - `frontend/src/features/dashboard/components/asset-feedback-section.tsx`（新增）
  - `frontend/src/features/dashboard/components/recent-activity-section.tsx`（新增）
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`（新增）
  - `frontend/src/features/dashboard/components/feedback-signal-card.tsx`（新增）
  - `frontend/src/features/dashboard/components/recent-activity-item-card.tsx`（新增）
  - `frontend/src/features/dashboard/lib/cta-matrix.ts`（新增：CTA 命中计算）
  - `frontend/src/features/dashboard/lib/dashboard-source.ts`（新增：来源参数工具）
  - `frontend/src/routes/modules/index.tsx`、`products/index.tsx`、`repositories/index.tsx`、`decisions/index.tsx`（修改：扩展 validateSearch）
  - `frontend/src/routes/modules/$moduleId.tsx`、`products/$productId.tsx`、`repositories/$repositoryId.tsx`、`decisions/$decisionId.tsx`（修改：新增 validateSearch）
  - `frontend/src/routes/modules/new.tsx`、`products/new.tsx`、`repositories/new.tsx`（修改：扩展 validateSearch）
  - 既有 List / Detail / Create 页面组件（修改：新增"返回 Dashboard"导航）

## ADDED Requirements

### Requirement: Dashboard 路由与主导航接入必须按 phase05-05 落地

系统 SHALL 将 `Dashboard` 路由与主导航接入落地为仓库内实际代码，对齐 `phase05-05` 已冻结的 `/dashboard` 路由与一级导航接入方式。

#### Scenario: Dashboard 路由文件落点

- **WHEN** 实现 Dashboard 路由
- **THEN** 必须新增 `frontend/src/routes/dashboard.tsx`
- **AND** 必须使用 `createFileRoute('/dashboard')` 注册路由
- **AND** 路由组件必须只承接 `DashboardHomePage`
- **AND** 当前阶段不得为 `/dashboard` 引入 `validateSearch`（`/dashboard` 不承接搜索参数）
- **AND** 不得把 `/dashboard` 拆成 `overview / activity / feedback` 等子路由

#### Scenario: 主导航接入

- **WHEN** 修改 `frontend/src/routes/__root.tsx`
- **THEN** 必须在一级导航中新增 `Dashboard` 入口，链接到 `/dashboard`
- **AND** `Dashboard` 入口必须置于既有 `Modules / Decisions / Products / Repositories` 之前（作为首项）
- **AND** 既有四个一级导航必须继续保留
- **AND** 不得把 `Dashboard` 做成隐藏入口或二级入口

### Requirement: Dashboard 类型与 API 适配层必须从 proto envelope 单向派生

系统 SHALL 将 Dashboard 前端类型与 API 适配层落地为 `frontend/src/features/dashboard/types.ts` 与 `frontend/src/features/dashboard/data/api-adapter.ts`，从 `phase05-12` 已落地的后端响应包络单向派生。

#### Scenario: 类型定义落点

- **WHEN** 实现 Dashboard 前端类型
- **THEN** 必须新增 `frontend/src/features/dashboard/types.ts`
- **AND** 必须定义 `DashboardOverview` 接口承接 `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count`
- **AND** 必须定义 `FeedbackSignal` 接口承接 `signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label`
- **AND** 必须定义 `ProductAssetCoverageSummary` 接口承接 `fully_bound_product_count / missing_both_bindings_count / missing_repository_binding_count / missing_module_binding_count / representative_signals`
- **AND** 必须定义 `RecentActivityItem` 接口承接 `activity_type / activity_at / target_type / target_id / target_label`
- **AND** 必须定义响应 envelope 类型：`DashboardOverviewResponse { overview: DashboardOverview }`、`FeedbackSignalsResponse { current_focus_signals: FeedbackSignal[]; asset_feedback_summary: ProductAssetCoverageSummary }`、`RecentActivitiesResponse { activities: RecentActivityItem[] }`
- **AND** `activity_at` 必须使用 `string` 类型（ISO 8601 时间字符串，由 JSON 反序列化得到）
- **AND** 字段名必须使用 snake_case，与后端 JSON tag 对齐

#### Scenario: API 适配层落点

- **WHEN** 实现 Dashboard API 适配层
- **THEN** 必须新增 `frontend/src/features/dashboard/data/api-adapter.ts`
- **AND** 必须复用现有 `module-registry/data/api-adapter.ts` 的 `request<T>` 封装模式（GET 不设 Content-Type 避免 CORS preflight，非 2xx 抛 `ApiError`）
- **AND** 必须导出 `fetchDashboardOverview(): Promise<DashboardOverview>`、`fetchFeedbackSignals(): Promise<FeedbackSignalsResponse>`、`fetchRecentActivities(): Promise<RecentActivitiesResponse>`
- **AND** 三个函数分别请求 `/api/dashboard/overview`、`/api/dashboard/feedback-signals`、`/api/dashboard/recent-activities`
- **AND** 必须导出 `ApiError` 类型供页面层使用
- **AND** 不得引入第二套 fetch 封装或第二套错误类型

### Requirement: DashboardHomePage 必须编排三个独立查询并派生整页状态

系统 SHALL 将 `DashboardHomePage` 落地为三个独立 TanStack Query 查询的编排器，派生 `phase05-06` 已冻结的整页视图状态。

#### Scenario: 三个独立查询

- **WHEN** 实现 `DashboardHomePage`
- **THEN** 必须使用三个独立 `useQuery` hook 分别请求 overview / feedback / recent-activities
- **AND** 三个 query 的 `queryKey` 必须相互独立（如 `['dashboard-overview']`、`['dashboard-feedback-signals']`、`['dashboard-recent-activities']`）
- **AND** 三个 query 默认 `enabled: true`，同时发起
- **AND** 不得把三个查询合并为单一 query

#### Scenario: 整页视图状态派生

- **WHEN** `DashboardHomePage` 根据三个查询状态派生整页视图状态
- **THEN** 必须派生 `initial-loading / ready / page-error` 三态
- **AND** `initial-loading` 只允许出现在 overview query 首次 pending 时
- **AND** `page-error` 只允许由 overview query 失败触发
- **AND** 一旦 overview query 成功，整页必须进入 `ready`，即使 feedback 或 recent-activity 局部失败也不得回退
- **AND** 不得把 feedback 或 recent-activity 失败升级为 `page-error`

#### Scenario: 整页重试

- **WHEN** 整页处于 `page-error`
- **THEN** 必须提供整页级重试入口
- **AND** 整页重试必须同时重新触发 overview / feedback / recent-activity 三个 query
- **AND** 不得只重试 overview 而遗漏附属聚合

### Requirement: 四个区块组件必须按 phase05-06 状态模型实现

系统 SHALL 将 Dashboard 主体布局落地为 `DashboardStatBar + CurrentFocusSection + AssetFeedbackSection + RecentActivitySection` 的紧凑实现，各自承接 `phase05-06` 已冻结的分区状态模型与 `phase05-10` 已冻结的区块语义。

#### Scenario: DashboardStatBar 实现

- **WHEN** 实现 `DashboardStatBar`
- **THEN** 必须将 `dashboard_overview` 的 6 个概览数字与资产覆盖的 4 个缺口计数收敛为单一紧凑条带
- **AND** `overviewQueryState` 成功后必须直接进入 `ready`，零计数仍为 `ready`，不得引入 `empty` 态
- **AND** `feedbackQueryState` 未成功时，资产覆盖部分必须独立展示 `loading` 骨架或 `error` 占位，不得把整条 stat bar 升级为整页错误
- **AND** stat bar 必须保持单容器展示，在 PC 桌面优先承接全局概览，在移动浏览器允许折行但不得撑出横向滚动

#### Scenario: DashboardStatBar 概览点击规则

- **WHEN** 渲染 `DashboardStatBar` 的概览组
- **THEN** `module_count / product_count / repository_count / decision_count` 四个概览单元必须可点击
- **AND** 点击 `module_count` 必须导航到 `/modules`，携带 `fromDashboard=true / dashboardSection=overview / dashboardReturnTo=/dashboard`
- **AND** 点击 `product_count` 必须导航到 `/products`，携带同样来源参数
- **AND** 点击 `repository_count` 必须导航到 `/repositories`，携带同样来源参数
- **AND** 点击 `decision_count` 必须导航到 `/decisions`，携带同样来源参数
- **AND** `product_with_repository_count / product_with_module_count` 两个概览单元只展示，不跳转

#### Scenario: CurrentFocusSection 实现

- **WHEN** 实现 `CurrentFocusSection`
- **THEN** 必须承接 `feedbackQueryState` 派生的 `loading / ready / empty / error` 四态
- **AND** `ready` 时必须以 `FeedbackSignalCardList` 展示 `current_focus_signals`（最多 5 条）
- **AND** `empty` 时必须展示"暂无待处理反馈信号"成功空态
- **AND** `error` 时必须在区块内容区域内展示局部错误与重试入口
- **AND** 重试必须重新触发 feedback query

#### Scenario: AssetFeedbackSection 实现

- **WHEN** 实现 `AssetFeedbackSection`
- **THEN** 必须承接 `feedbackQueryState` 派生的 `loading / ready / empty / error` 四态（与 CurrentFocusSection 共享）
- **AND** `ready` 时必须以 `AssetFeedbackList` 展示 `representative_signals`（最多 3 条）
- **AND** 四类缺口计数（`fully_bound_product_count / missing_both_bindings_count / missing_repository_binding_count / missing_module_binding_count`）必须由 `DashboardStatBar` 统一展示，不得在本区块再复制第二份计数网格
- **AND** `empty` 时必须展示"暂无代表性缺口项"成功空态
- **AND** `error` 时必须在区块内容区域内展示局部错误与重试入口

#### Scenario: RecentActivitySection 实现

- **WHEN** 实现 `RecentActivitySection`
- **THEN** 必须承接 `activityQueryState` 派生的 `loading / ready / empty / error` 四态
- **AND** `ready` 时必须以 `RecentActivityList` 展示 `activities`（最多 10 条，后端已排序）
- **AND** `empty` 时必须展示"暂无最近活动"成功空态
- **AND** `error` 时必须在区块内容区域内展示局部错误与重试入口
- **AND** 重试必须重新触发 recent-activity query

### Requirement: FeedbackSignalCard 必须按信号类型跳转

系统 SHALL 将 `FeedbackSignalCard` 落地为按 `target_type` 与 `signal_code` 决定跳转目标的实现。

#### Scenario: 决策信号跳转

- **WHEN** 渲染 `target_type = decision_detail` 的反馈卡片
- **THEN** 点击必须导航到 `/decisions/$decisionId`，`target_id` 作为 `decisionId`
- **AND** 必须携带 `fromDashboard=true / dashboardReturnTo=/dashboard`
- **AND** `dashboardSection` 必须为 `current-focus`（若卡片位于 CurrentFocusSection）或 `asset-feedback`（若位于 AssetFeedbackSection）

#### Scenario: 聚合决策信号跳转

- **WHEN** 渲染 `target_type = decision_list` 的反馈卡片
- **THEN** 点击必须导航到 `/decisions`（Decision Center / List）
- **AND** 必须携带 `fromDashboard=true / dashboardSection=current-focus|asset-feedback / dashboardReturnTo=/dashboard`

#### Scenario: 产品缺口信号跳转

- **WHEN** 渲染 `target_type = product_detail` 的反馈卡片
- **THEN** 点击必须导航到 `/products/$productId`，`target_id` 作为 `productId`
- **AND** 必须携带 `fromDashboard=true / dashboardReturnTo=/dashboard`
- **AND** `dashboardSection` 必须为 `current-focus` 或 `asset-feedback`

### Requirement: RecentActivityItemCard 必须按活动类型跳转

系统 SHALL 将 `RecentActivityItemCard` 落地为按 `target_type` 决定跳转目标的实现。

#### Scenario: 活动项跳转映射

- **WHEN** 渲染 `RecentActivityItemCard`
- **AND** 用户点击卡片
- **THEN** 必须按 `target_type` 导航：
  - `module_detail` → `/modules/$moduleId`，`target_id` 作为 `moduleId`
  - `product_detail` → `/products/$productId`，`target_id` 作为 `productId`
  - `repository_detail` → `/repositories/$repositoryId`，`target_id` 作为 `repositoryId`
  - `decision_detail` → `/decisions/$decisionId`，`target_id` 作为 `decisionId`
- **AND** 必须携带 `fromDashboard=true / dashboardSection=recent-activity / dashboardReturnTo=/dashboard`

### Requirement: DashboardPrimaryActionPanel 必须按 CTA 矩阵命中

系统 SHALL 将 `DashboardPrimaryActionPanel` 落地为按 `phase05-04 / phase05-10` 已冻结的 CTA 优先级矩阵命中的实现。

#### Scenario: CTA 命中计算

- **WHEN** 实现 CTA 命中逻辑
- **THEN** 必须新增 `frontend/src/features/dashboard/lib/cta-matrix.ts` 导出 `computePrimaryCta(overview, feedbackSignals)` 函数
- **AND** 函数必须按以下顺序命中：
  1. `module_count=0 && product_count=0 && repository_count=0 && decision_count=0` → CTA 1，目标 `/modules/new`
  2. `module_count=0 && (product_count>0 || repository_count>0 || decision_count>0)` → CTA 2，目标 `/modules/new`
  3. `module_count>0 && product_count=0` → CTA 3，目标 `/products/new`
  4. `module_count>0 && product_count>0 && repository_count=0` → CTA 4，目标 `/repositories/new`
  5. 存在 `pending_decision` 信号 → CTA 5，目标最高优先级决策信号落点
  6. 存在 `product_missing_both_bindings` 信号 → CTA 6，目标对应 Product Detail
  7. 存在 `product_missing_repository_binding` 信号 → CTA 7，目标对应 Product Detail
  8. 存在 `product_missing_module_binding` 信号 → CTA 8，目标对应 Product Detail
  9. 无缺口 → CTA 9，无主 CTA
- **AND** 命中 CTA 1-4 时必须返回 `{ state: 'ready', cta: CTA_1_TO_4 }`
- **AND** 命中 CTA 5-8 时必须返回 `{ state: 'ready', cta: CTA_5_TO_8 }`
- **AND** 命中 CTA 9 时必须返回 `{ state: 'hidden' }`

#### Scenario: 主 CTA 状态机

- **WHEN** `DashboardPrimaryActionPanel` 根据查询状态与 CTA 命中结果派生面板状态
- **THEN** overview query 未成功时必须为 `computing`，不渲染任何 CTA
- **AND** overview 成功且命中 CTA 1-4 时必须为 `ready`，渲染对应创建导向 CTA 按钮
- **AND** overview 成功未命中 CTA 1-4，feedback query 未成功时必须为 `computing`
- **AND** overview 成功未命中 CTA 1-4，feedback query 失败时必须为 `suppressed`，不渲染 CTA
- **AND** overview 成功未命中 CTA 1-4，feedback 成功且命中 CTA 5-8 时必须为 `ready`
- **AND** overview 成功未命中 CTA 1-4，feedback 成功且无缺口（CTA 9）时必须为 `hidden`
- **AND** 同一时刻只允许展示一个主 CTA

#### Scenario: CTA 1-4 跳转参数

- **WHEN** 命中 CTA 1-4 并渲染创建导向 CTA 按钮
- **THEN** 点击必须导航到对应 Create 路由
- **AND** 必须携带 `fromDashboard=true / dashboardSection=empty-state / dashboardReturnTo=/dashboard`

### Requirement: 既有路由必须承接 Dashboard 来源参数

系统 SHALL 修改既有四类 List / 四类 Detail / 三类 Create 路由，使其 `validateSearch` 承接 `fromDashboard / dashboardSection / dashboardReturnTo` 三个可选搜索参数。

#### Scenario: List 路由扩展 validateSearch

- **WHEN** 修改 `modules/index.tsx`、`products/index.tsx`、`repositories/index.tsx`、`decisions/index.tsx`
- **THEN** 必须在既有原生搜索参数 schema 基础上扩展三个可选字段：
  - `fromDashboard: z.boolean().optional()`
  - `dashboardSection: z.enum(['overview', 'current-focus', 'asset-feedback', 'recent-activity', 'empty-state']).optional()`
  - `dashboardReturnTo: z.string().optional()`
- **AND** 不得移除各列表页原生的搜索参数（如 `queryText / statusFilter`）
- **AND** 不得改变原生搜索参数的默认值与 catch 行为

#### Scenario: Detail 路由新增 validateSearch

- **WHEN** 修改 `modules/$moduleId.tsx`、`products/$productId.tsx`、`repositories/$repositoryId.tsx`、`decisions/$decisionId.tsx`
- **THEN** 必须新增 `validateSearch` schema 承接三个 Dashboard 来源参数（同 List 路由）
- **AND** 不得影响既有路径参数（`$moduleId / $productId` 等）的承接
- **AND** 当前阶段不得为详情页引入除 Dashboard 来源参数之外的新搜索参数

#### Scenario: Create 路由扩展 validateSearch

- **WHEN** 修改 `modules/new.tsx`、`products/new.tsx`、`repositories/new.tsx`
- **THEN** 必须扩展 `validateSearch` 承接三个 Dashboard 来源参数
- **AND** 不得移除既有 Create 路由的原生搜索参数（如 `fromList`）

#### Scenario: 来源参数工具函数

- **WHEN** 实现来源参数的构造与解析
- **THEN** 必须新增 `frontend/src/features/dashboard/lib/dashboard-source.ts`
- **AND** 必须导出 `buildDashboardSourceParams(section)` 返回 `{ fromDashboard: true, dashboardSection: section, dashboardReturnTo: '/dashboard' }`
- **AND** 必须导出 `useDashboardSource()` hook 从当前路由搜索参数读取来源上下文
- **AND** 必须导出 `navigateBackToDashboard(navigate, section)` 工具，使用 TanStack Router 的 `navigate` 跳转到 `/dashboard`，并通过路由 state 承接 `dashboardSection` 作为一次性恢复标记

### Requirement: 既有页面必须支持返回 Dashboard 导航

系统 SHALL 修改既有 List / Detail / Create 页面组件，在 `fromDashboard=true` 时提供"返回 Dashboard"导航。

#### Scenario: List 页面返回 Dashboard

- **WHEN** List 页面检测到 `fromDashboard=true`
- **THEN** 必须在页面顶部或工具栏区域展示"返回 Dashboard"按钮
- **AND** 点击必须导航到 `/dashboard`
- **AND** 必须通过 TanStack Router 的 `navigate({ to: '/dashboard', state: { dashboardSection: <当前 section> } })` 承接一次性恢复标记
- **AND** 不得移除 List 页面原生的"返回"或筛选控件

#### Scenario: Detail 页面返回 Dashboard

- **WHEN** Detail 页面检测到 `fromDashboard=true`
- **THEN** 必须展示"返回 Dashboard"按钮
- **AND** 点击必须导航到 `/dashboard`，通过路由 state 承接 `dashboardSection`
- **AND** 当 Detail 页同时携带 `fromList=true` 时，必须同时保留"返回列表"与"返回 Dashboard"两个导航入口
- **AND** "返回列表"使用 `fromList` 上下文，"返回 Dashboard"使用 `fromDashboard` 上下文

#### Scenario: Create 页面取消返回 Dashboard

- **WHEN** Create 页面检测到 `fromDashboard=true`
- **AND** 用户点击取消
- **THEN** 必须导航到 `/dashboard`，而不是回列表
- **AND** 必须通过路由 state 承接 `dashboardSection=empty-state`
- **AND** 提交成功后进入 Detail 页时必须继续保留 `fromDashboard=true`
- **AND** 不得因为 `fromDashboard=true` 就在提交成功后自动跳回 Dashboard

### Requirement: Dashboard 主动返回的一次性状态承接必须用路由 state

系统 SHALL 将"主动返回 Dashboard"时的 `dashboardSection` 恢复标记通过 TanStack Router 的路由 state 一次性承接，不引入持久化层。

#### Scenario: 一次性路由 state 承接

- **WHEN** 用户在携带 `fromDashboard=true` 的目标页主动触发"返回 Dashboard"
- **THEN** 必须使用 `navigate({ to: '/dashboard', state: { dashboardSection: <section> } })`
- **AND** `DashboardHomePage` 必须通过 `useRouterState().location.state` 读取该一次性 `dashboardSection`
- **AND** 读取后该 state 不得持久化，刷新 `/dashboard` 后不得继续保留
- **AND** 不得使用 `sessionStorage / localStorage / Zustand store` 承接该恢复标记
- **AND** 不得把该 state 提升为 `/dashboard` 的搜索参数

### Requirement: PC 与移动浏览器布局降级必须用 Tailwind 响应式

系统 SHALL 通过 Tailwind CSS 响应式类实现 `phase05-05 / phase05-10` 已冻结的 PC / 移动浏览器布局降级，不引入第二套移动端 UI 架构。

#### Scenario: PC 桌面布局

- **WHEN** `DashboardHomePage` 在 PC 桌面环境（`md:` 断点以上）展示
- **THEN** `DashboardStatBar` 必须位于页面主体顶部，优先承接全局概览与资产覆盖信息
- **AND** `CurrentFocusSection` 必须在第一屏优先位置
- **AND** `CurrentFocusSection` 与 `AssetFeedbackSection` 必须位于左列，`RecentActivitySection` 必须位于右列且允许跨两行展示
- **AND** `RecentActivitySection` 必须通过限高与内部滚动控制高度，避免无限撑高页面
- **AND** 四个区块允许同屏呈现
- **AND** 不得牺牲 `CurrentFocus` 的主行动优先级

#### Scenario: 移动浏览器窄屏布局

- **WHEN** `DashboardHomePage` 在移动浏览器窄屏环境（`md:` 断点以下）展示
- **THEN** 所有区块必须单列垂直重排
- **AND** 区块顺序必须保持 `CurrentFocus / DashboardStatBar(dashboard_overview + asset coverage) / AssetFeedback / RecentActivity`
- **AND** 必须通过信息裁剪与列表收束降低拥挤度
- **AND** `DashboardStatBar` 在窄屏下允许折行，但不得撑出 `main` 内容区或造成 header / main 横向错位
- **AND** 不得引入第二套独立移动端页面或第二套路由树

### Requirement: Dashboard 前端必须只读不写

系统 SHALL 确保 Dashboard 前端模块只承接读与跳转，不承接任何业务写入。

#### Scenario: 只读边界

- **WHEN** 实现 Dashboard 前端
- **THEN** `frontend/src/features/dashboard/data/api-adapter.ts` 只允许导出 `fetch*` 读函数
- **AND** 不得导出 `create / update / delete / bind / link` 等写入函数
- **AND** Dashboard 组件不得调用既有 canonical 模块的写入 API
- **AND** Dashboard 跳转到 Create 页后，写入由 Create 页所属 canonical 模块承接，Dashboard 不承接写入副作用

## MODIFIED Requirements

### Requirement: phase05-05 页面与组件分层从"设计冻结"推进为"仓库实现落地"

系统 SHALL 将 `phase05-05` 已冻结的 Dashboard 前端页面、路由、组件分层、来源参数承接与布局降级策略，从"规格层定义"推进为"仓库内实际存在且可编译运行的代码"。

#### Scenario: 前端分层落地

- **WHEN** `phase05-13` 完成
- **THEN** Dashboard 前端不再只停留在 `phase05-05` 文档正文中
- **AND** 必须在仓库 `frontend/src/routes/dashboard.tsx` 与 `frontend/src/features/dashboard/` 内拥有实际文件、可编译代码与可运行页面
- **AND** 后续 `phase05-14` 必须优先引用该已落地前端模块

### Requirement: phase05-06 状态模型与交互流从"设计冻结"推进为"仓库实现落地"

系统 SHALL 将 `phase05-06` 已冻结的整页查询状态、整页视图状态、分区状态模型、主 CTA 状态机、返回恢复与刷新恢复规则，从"规格层定义"推进为"仓库内实际存在且可编译运行的代码"。

#### Scenario: 状态模型落地

- **WHEN** `phase05-13` 完成
- **THEN** Dashboard 前端状态模型不再只停留在 `phase05-06` 文档正文中
- **AND** 必须在 `DashboardHomePage` 与区块组件中实际编排三个独立查询并派生整页/分区状态
- **AND** 主 CTA 状态机必须在 `DashboardPrimaryActionPanel` 与 `cta-matrix.ts` 中实际落地

## REMOVED Requirements

### Requirement: Dashboard 前端主线继续停留在设计层

**Reason**: `phase05-13` 的目标就是把 Dashboard 前端从"已经设计好"推进到"仓库中已经落地、可编译、可运行、可被联调验收"的状态。

**Migration**: 后续 `phase05-14` 联调验收应统一从 `frontend/src/routes/dashboard.tsx` 与 `frontend/src/features/dashboard/` 进入；`phase05-05 / 06` 继续作为设计上游，不再承担仓库内前端主线入口职责。
