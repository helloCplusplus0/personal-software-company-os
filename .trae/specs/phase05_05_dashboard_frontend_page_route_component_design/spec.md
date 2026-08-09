# Phase05-05 Dashboard 前端页面、路由与组件分层设计 Spec

## Why

`phase05-01` 已冻结 `Dashboard Home` 的页面边界、区块归属、点击热区与 `PC / 移动浏览器` 信息层级，`phase05-03` 已冻结 Dashboard 到既有 owner 页面的跳转矩阵与来源参数，`phase05-04` 已冻结三类聚合读与错误语义前提。但要真正进入前端实现，还必须把 `Dashboard` 的页面文件落点、路由树、导航接入、区块级/卡片级/列表级组件职责，以及跨路由来源参数承接方式写成单值结论。否则后续实现会在“页面放哪、路由怎么挂、哪些组件该复用、目标页如何接 Dashboard 来源参数”之间继续漂移。

> 阶段分工约束：本规格只冻结前端页面分层、路由树、组件树、Dashboard 来源参数在路由层的最小承接方式，以及 `PC / 移动浏览器` 下的布局降级策略。前端页面级状态模型、局部 loading / error / empty 的状态机、区块展开/折叠、返回恢复细节统一由 `phase05-06` 承接；后端模块边界、接口分组与 owner 由 `phase05-07` 承接；`.proto` 服务命名、消息结构、包名版本与 `chi + JSON HTTP` 显式映射由 `phase05-08` 承接，不在本规格中提前冻结。
>
> 与 `phase05-06` 的边界划分：本规格只冻结“页面如何组织、路由如何命名、参数如何承接、组件职责如何拆分”，不冻结“页面如何流转”。也就是说，`fromDashboard / dashboardSection / dashboardReturnTo` 在本规格中只冻结为路由搜索参数承接模型；这些参数在列表、详情、Create 场景下的返回行为、刷新恢复、局部失败 fallback、主 CTA 展示/抑制规则，继续由 `phase05-06` 沿用 `phase05-03 / 04` 单值化，不在本规格中重复冻结。

## What Changes

- 冻结 `Dashboard Home` 的前端页面文件落点
- 冻结 `Dashboard` 的路由文件、正式 URL 与主导航接入方式
- 冻结 `Dashboard Home` 的页面级、区块级、卡片级与列表级组件树
- 冻结 `Dashboard` 到既有 `Module / Product / Repository / Decision` 页面时，路由层最小参数承接方式
- 冻结 `PC / 移动浏览器` 双场景下的页面布局降级策略
- 明确当前阶段不引入第二套移动端 UI 架构，不提前冻结运行时 hook、Query key、store API 与页面状态容器

## Impact

- Affected specs: `phase05_dashboard_feedback_foundation`
- Affected code: `frontend/src/routes/__root.tsx`、新增 `frontend/src/routes/dashboard.tsx`、既有 `frontend/src/routes/modules/*`、`frontend/src/routes/products/*`、`frontend/src/routes/repositories/*`、`frontend/src/routes/decisions/*` 的 Dashboard 来源参数承接
- Affected code: 新增 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` 与 `frontend/src/features/dashboard/components/*`

## ADDED Requirements

### Requirement: Dashboard 页面文件落点冻结

系统 SHALL 将 `Dashboard` 模块的页面文件落点冻结为单值结论，遵循仓库既有 `features/<module>/pages/` 模式。

#### Scenario: 判断 Dashboard 页面文件落点

- **WHEN** 后续实现讨论 `Dashboard` 的页面文件位置
- **THEN** 必须得到以下单值结论：
- **AND** `DashboardHomePage` → `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
- **AND** 当前阶段不得并行引入 `Dashboard Detail`、独立 `Feedback` 页面或第二套 Dashboard 子页面体系

### Requirement: Dashboard 路由树与主导航接入冻结

系统 SHALL 将 `Dashboard` 的路由树、正式 URL 与主导航接入方式冻结为单值结论。

#### Scenario: 判断 Dashboard 路由文件与 URL

- **WHEN** 后续实现讨论 `Dashboard` 的路由结构
- **THEN** 必须得到以下单值结论：
- **AND** `DashboardRoute` → `/dashboard` → `frontend/src/routes/dashboard.tsx`
- **AND** `frontend/src/routes/dashboard.tsx` 必须只承接 `DashboardHomePage`
- **AND** 当前阶段不得把 `/dashboard` 提前拆成 `overview / activity / feedback` 等子路由
- **AND** 当前阶段不得把根路径 `/` 重新解释为 `Dashboard Home`

#### Scenario: 判断 Dashboard 与既有主导航的接入方式

- **WHEN** 后续实现讨论 `Dashboard` 在前端导航中的接入方式
- **THEN** 必须在 `frontend/src/routes/__root.tsx` 的一级导航中新增 `Dashboard` 入口
- **AND** 既有 `Modules / Decisions / Products / Repositories` 一级导航继续保留
- **AND** 不得把 `Dashboard` 做成隐藏入口、二级入口或替换根布局宿主语义

### Requirement: Dashboard Home 页面级组件树冻结

系统 SHALL 将 `Dashboard Home` 的页面级组件树冻结为单值结论，确保区块层次与 `phase05-01` 四区块归属保持一致。

#### Scenario: 判断 Dashboard Home 页面级组件树

- **WHEN** 后续实现讨论 `DashboardHomePage` 的页面结构
- **THEN** 必须得到以下单值组件树：
- **AND** `DashboardHomePageShell`
  - `CurrentFocusSection`
  - `DashboardOverviewSection`
  - `AssetFeedbackSection`
  - `RecentActivitySection`
  - `DashboardPrimaryActionPanel`
- **AND** 组件树顺序与 `phase05-01` 已冻结的第一屏优先级（`Current Focus / Next Action` 优先）保持一致
- **AND** 当前阶段四个区块必须继续对应 `dashboard_overview / Current Focus / Asset Feedback / Recent Activity`
- **AND** `DashboardPrimaryActionPanel` 承接 `phase05-04` 已冻结的主 CTA 优先级矩阵与 `phase05-01` 已冻结的空状态主 CTA 按钮可点模式
- **AND** 不得把 `Dashboard Home` 拆成第二套页面级工作台

### Requirement: Dashboard 区块级、卡片级与列表级组件职责冻结

系统 SHALL 将 `Dashboard` 的区块级、卡片级与列表级组件职责冻结为单值结论，避免后续实现把不同语义揉成一个巨型组件。

#### Scenario: 判断 Dashboard Overview 组件职责

- **WHEN** 后续实现讨论 `dashboard_overview` 区块的组件拆分
- **THEN** 必须得到以下单值组件树：
- **AND** `DashboardOverviewSection`
  - `DashboardOverviewCardGrid`
    - `DashboardOverviewCard`
- **AND** `DashboardOverviewSection` 只承接概览区块标题、区块容器与卡片集合布局
- **AND** `DashboardOverviewCard` 只承接单个概览指标展示与点击入口
- **AND** `DashboardOverviewCard` 必须继续遵守 `phase05-01 / 03` 已冻结的“仅四类一级概览卡片可点击、派生覆盖率指标只展示不跳转”规则

#### Scenario: 判断 Current Focus 组件职责

- **WHEN** 后续实现讨论 `Current Focus / Next Action` 区块的组件拆分
- **THEN** 必须得到以下单值组件树：
- **AND** `CurrentFocusSection`
  - `FeedbackSignalCardList`
    - `FeedbackSignalCard`
- **AND** `CurrentFocusSection` 只承接主行动队列区块容器与标题语义
- **AND** `FeedbackSignalCardList` 只承接反馈卡片列表编排
- **AND** `FeedbackSignalCard` 只承接单条高优先级反馈信号的标题、摘要、优先级与跳转入口展示

#### Scenario: 判断 Asset Feedback 组件职责

- **WHEN** 后续实现讨论 `Asset Feedback` 区块的组件拆分
- **THEN** 必须得到以下单值组件树：
- **AND** `AssetFeedbackSection`
  - `AssetFeedbackList`
    - `AssetFeedbackItemCard`
- **AND** `AssetFeedbackSection` 只承接补充摘要区块容器与标题语义
- **AND** `AssetFeedbackList` 只承接代表性缺口项集合编排
- **AND** `AssetFeedbackItemCard` 只承接单条资产缺口摘要与跳转入口展示
- **AND** 不得把 `Asset Feedback` 与 `Current Focus` 复用成一个无语义差别的区块容器

#### Scenario: 判断 Recent Activity 组件职责

- **WHEN** 后续实现讨论 `Recent Activity` 区块的组件拆分
- **THEN** 必须得到以下单值组件树：
- **AND** `RecentActivitySection`
  - `RecentActivityList`
    - `RecentActivityItemCard`
- **AND** `RecentActivitySection` 只承接活动流区块容器与标题语义
- **AND** `RecentActivityList` 只承接活动项集合编排
- **AND** `RecentActivityItemCard` 只承接单条活动的对象标签、时间与跳转入口展示

#### Scenario: 判断 DashboardPrimaryActionPanel 组件职责

- **WHEN** 后续实现讨论主 CTA 与空状态的组件拆分
- **THEN** 必须得到以下单值结论：
- **AND** `DashboardPrimaryActionPanel` 独立于四区块，直接挂载在 `DashboardHomePageShell` 下
- **AND** `DashboardPrimaryActionPanel` 只承接 `phase05-04` 已冻结的主 CTA 优先级矩阵的命中与展示
- **AND** `DashboardPrimaryActionPanel` 必须遵守 `phase05-01` 已冻结的空状态主 CTA 按钮可点模式，不采用整卡模式
- **AND** `DashboardPrimaryActionPanel` 必须遵守 `phase05-04` 已冻结的单主 CTA 约束，同一时刻只展示一个主 CTA
- **AND** 当前阶段不冻结主 CTA 的具体状态机（命中、抑制、loading），由 `phase05-06` 承接

### Requirement: Dashboard 来源参数的路由承接冻结

系统 SHALL 将 `Dashboard` 来源参数在前端路由层的最小承接方式冻结为单值结论，沿用 `phase05-03` 已冻结的命名，不发明第二套参数体系。

#### Scenario: Dashboard 自身路由参数边界

- **WHEN** 后续实现讨论 `frontend/src/routes/dashboard.tsx` 的搜索参数
- **THEN** 当前阶段不得为 `Dashboard Home` 引入新的查询参数、筛选参数或区块级 URL 状态
- **AND** `/dashboard` 当前阶段默认作为无搜索参数的首页入口存在

#### Scenario: List 路由承接 Dashboard 来源参数

- **WHEN** 后续实现讨论 `Module Registry / List`、`Product Registry / List`、`Repository Binding / List`、`Decision Center / List` 的路由搜索参数
- **THEN** 这些路由在保留各自原生搜索参数的同时，必须允许附加承接 `fromDashboard`、`dashboardSection`、`dashboardReturnTo`
- **AND** 不得以新的 `dashboardFrom`、`returnToDashboard`、`fromDashboardList` 替代既有命名
- **AND** 不得因为接入 Dashboard 来源参数而移除各列表页原生的搜索参数模型

#### Scenario: Detail 路由承接 Dashboard 来源参数

- **WHEN** 后续实现讨论 `Module Detail`、`Product Detail`、`Repository Binding Detail / Workspace`、`Decision Detail` 的路由参数
- **THEN** 详情页必须继续使用路径参数承接对象身份（如 `$moduleId / $productId / $repositoryId / $decisionId`）
- **AND** 必须通过搜索参数承接 `fromDashboard`、`dashboardSection`、`dashboardReturnTo`
- **AND** 不得把 `dashboardSection` 或 `dashboardReturnTo` 挪到路径参数层

#### Scenario: Create 路由承接 Dashboard 来源参数

- **WHEN** 后续实现讨论 `Module Registry / Create`、`Product Registry / Create`、`Repository Binding / Create` 的路由搜索参数
- **THEN** Create 路由必须允许承接 `fromDashboard`、`dashboardSection`、`dashboardReturnTo`
- **AND** 当前阶段只允许空状态主 CTA 通过这些参数接入 Create 页
- **AND** 不得在 Create 路由层额外发明 Dashboard 私有路径

### Requirement: Dashboard 组件归属原则冻结

系统 SHALL 将 `Dashboard` 组件归属原则冻结为单值结论，默认组件归属于 `dashboard` feature，不提前抽象新的共享层。

#### Scenario: 判断 Dashboard 组件归属

- **WHEN** 后续实现讨论 `Dashboard` 组件的归属与复用
- **THEN** `DashboardOverviewSection / DashboardOverviewCardGrid / DashboardOverviewCard / CurrentFocusSection / FeedbackSignalCardList / FeedbackSignalCard / AssetFeedbackSection / AssetFeedbackList / AssetFeedbackItemCard / RecentActivitySection / RecentActivityList / RecentActivityItemCard / DashboardPrimaryActionPanel` 默认归属于 `frontend/src/features/dashboard/components/`
- **AND** 只有存在跨 `dashboard` 与既有 `modules / products / repositories / decisions` 页面复用证据时，才允许再抽为共享组件
- **AND** 当前阶段不得为了“组件纯洁”提前拆出新的 Dashboard 专属基础组件层

### Requirement: PC 与移动浏览器布局降级策略冻结

系统 SHALL 在单一 `React Web` 前端交付策略下，冻结 `Dashboard Home` 在 `PC / 移动浏览器` 双场景下的布局降级策略，不引入第二套移动端 UI 架构。

#### Scenario: 判断 Dashboard Home 在 PC 桌面环境的布局

- **WHEN** `DashboardHomePage` 在 `PC` 桌面环境展示
- **THEN** 页面必须优先保证 `Current Focus / Next Action` 的第一行动优先级
- **AND** `dashboard_overview` 必须以概览卡片网格形式存在
- **AND** `Asset Feedback` 与 `Recent Activity` 必须作为补充区块与主行动区并列承接
- **AND** 当前阶段允许同屏呈现四区块，但不得牺牲 `Current Focus` 的主行动优先级

#### Scenario: 判断 Dashboard Home 在移动浏览器窄屏环境的布局

- **WHEN** `DashboardHomePage` 在移动浏览器窄屏环境展示
- **THEN** 必须沿用与桌面端同一套区块语义与动作体系
- **AND** 区块顺序必须优先保持 `Current Focus / dashboard_overview / Asset Feedback / Recent Activity`
- **AND** 必须通过单列垂直重排、信息裁剪与列表收束降低拥挤度
- **AND** 不得引入第二套独立移动端页面或第二套路由树

### Requirement: 运行时实现细节不冻结

系统 SHALL 明确当前阶段只冻结前端页面分层、路由树、组件树、参数承接与布局降级策略，不冻结运行时实现细节。

#### Scenario: 判断不冻结的运行时实现细节

- **WHEN** 后续实现讨论 hook 命名、Query key、store API 命名、请求取消、缓存时间、局部重试策略或 Zustand 容器拆分
- **THEN** 它们都不属于本规格冻结范围
- **AND** 页面级 UI 状态、局部 loading / error / empty 状态机与返回恢复细节统一由 `phase05-06` 承接

## MODIFIED Requirements

### Requirement: phase05-01 布局降级策略的前端实现解释

`phase05-01` 已冻结 `Dashboard Home` 的四区块归属与 `PC / 移动浏览器` 信息层级，`phase05-05` 在此基础上 SHALL 将这些结论进一步解释为前端页面、区块组件、列表组件与卡片组件的实现分层，而不是仅停留在页面信息结构描述。

#### Scenario: 页面结构到组件结构的解释

- **WHEN** 后续实现根据 `phase05-01` 编写 Dashboard 前端页面
- **THEN** 必须同时满足 `phase05-01` 的页面/区块边界与本规格冻结的组件树
- **AND** 不得只保留页面概念而缺失区块级、卡片级与列表级分层

### Requirement: phase05-03 来源参数的前端路由解释

`phase05-03` 已冻结 `fromDashboard / dashboardSection / dashboardReturnTo` 的语义，`phase05-05` 在此基础上 SHALL 将其进一步解释为前端路由层唯一允许承接的 Dashboard 来源参数集合。

#### Scenario: 来源参数的路由承接解释

- **WHEN** 后续实现修改 `modules / products / repositories / decisions` 路由搜索参数
- **THEN** 必须沿用 `fromDashboard / dashboardSection / dashboardReturnTo` 作为唯一 Dashboard 来源参数集合
- **AND** 不得在前端路由层另起一套 Dashboard 参数命名

## REMOVED Requirements

### Requirement: 第二套移动端 UI 架构

**Reason**: 当前阶段已冻结为单一 `React Web` 同时覆盖 `PC` 与移动浏览器，通过布局降级而不是第二套路由树或第二套移动端页面体系解决窄屏适配。引入独立移动端 UI 架构会直接违反 `shared_baseline` §2.3/§2.4 与 `architecture_plan` §4.3 的单一前端交付策略。

**Migration**: 若后续确需独立移动端客户端或第二套 UI 体系，必须进入新的 `phase / audit` 流程，不在 `phase05-05` 当前规格中处理。
