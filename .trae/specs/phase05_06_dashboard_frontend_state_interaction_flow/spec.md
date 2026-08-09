# Phase05-06 Dashboard 前端状态模型与交互流设计 Spec

## Why

`phase05-03` 已冻结 Dashboard 的跳转目标、来源参数与返回路径，`phase05-04` 已冻结三类聚合读、空态与错误语义前提，`phase05-05` 已冻结页面、路由与组件分层。但“页面放哪、组件怎么拆”还不等于“页面如何流转”。如果不继续把 `Dashboard Home` 的整页查询状态、分区加载状态、局部错误状态、主 CTA 命中/抑制规则、返回恢复与刷新恢复规则收紧，后续实现仍会在 `TanStack Router / TanStack Query / Zustand` 的边界上出现临场发挥，尤其会在“什么时候整页失败、什么时候只局部失败、什么时候可以显示主 CTA、从目标页回来如何恢复 Dashboard 上下文”上分叉。

> 阶段分工约束：本规格只冻结前端页面级状态模型、分区级状态模型、局部交互流、主 CTA 命中/抑制语义、来源上下文恢复规则与刷新恢复策略。后端模块边界、接口分组与 owner 由 `phase05-07` 承接；`.proto` 服务命名、消息结构、包名版本与 `chi + JSON HTTP` 显式映射由 `phase05-08` 承接，不在本规格中提前冻结。具体 hook 命名、Query key、store API、缓存时间、请求取消、预取策略与 optimistic update 不在本规格中冻结。
>
> 与 `phase05-05` 的边界划分：`phase05-05` 已冻结页面、路由、组件树与来源参数在路由层的最小承接方式；本规格承接这些结论，在此基础上冻结这些组件的状态语义、交互流转、返回恢复与刷新恢复，不重新发明第二套来源参数命名，也不重写页面/组件归属。

## What Changes

- 冻结 `Dashboard Home` 的整页查询状态、整页视图状态与整页失败语义
- 冻结 `DashboardOverviewSection / CurrentFocusSection / AssetFeedbackSection / RecentActivitySection / DashboardPrimaryActionPanel` 的分区级状态模型
- 冻结空系统、有数据、无反馈信号、无最近活动、局部失败与整页失败时的页面行为
- 冻结主 CTA 的命中、等待、抑制与隐藏规则，以及多缺口并存时的区块降级策略
- 冻结从 Dashboard 跳出到 List / Detail / Create 后再返回的上下文恢复策略
- 冻结刷新恢复、路由搜索参数与来源透传策略，明确 URL 为外层来源上下文的唯一事实源
- 明确当前阶段只冻结用户可见状态与交互语义，不把 hook、缓存或 store 实现写成既成事实

## Impact

- Affected specs: `phase05_dashboard_feedback_foundation`
- Affected code: 后续 `frontend/src/routes/dashboard.tsx` 与 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` 的页面容器状态、重试入口、主 CTA 命中逻辑与返回恢复逻辑
- Affected code: 后续 `frontend/src/features/dashboard/components/*` 的分区状态展示、局部错误呈现、空态呈现与 CTA 展示/抑制行为
- Affected code: 后续 `frontend/src/routes/modules/*`、`products/*`、`repositories/*`、`decisions/*` 的 Dashboard 来源参数保留与返回 Dashboard 逻辑

## ADDED Requirements

### Requirement: Dashboard 整页查询状态与整页视图状态冻结

系统 SHALL 将 `Dashboard Home` 的整页查询状态冻结为“主聚合读取状态 + 附属聚合读取状态 + 派生整页视图状态”的单值模型。

#### Scenario: Dashboard 最小查询状态

- **WHEN** 后续实现讨论 `DashboardHomePage` 的最小查询状态
- **THEN** 页面必须至少同时承接：
  - `overviewQueryState = pending / success / error`
  - `feedbackQueryState = idle / pending / success / error`
  - `activityQueryState = idle / pending / success / error`
- **AND** `DashboardOverviewRead` 作为主聚合读取，决定页面是否可进入整页成功语义
- **AND** `FeedbackSignalRead` 与 `RecentActivityRead` 作为附属聚合读取，不得提升为第二个“主页面生死开关”

#### Scenario: Dashboard 整页视图状态

- **WHEN** 页面根据三类读取结果决定整页展示
- **THEN** 派生整页视图状态必须冻结为 `initial-loading / ready / page-error`
- **AND** `initial-loading` 只允许出现在首轮进入且 `DashboardOverviewRead` 尚未成功前
- **AND** `page-error` 只允许由 `DashboardOverviewRead` 失败触发
- **AND** 一旦 `DashboardOverviewRead` 成功，页面必须进入 `ready`
- **AND** 即使附属聚合局部失败，页面也不得回退为 `page-error`

### Requirement: DashboardOverviewSection 状态模型冻结

系统 SHALL 将 `DashboardOverviewSection` 的状态模型冻结为只依赖主聚合读取的分区状态。

#### Scenario: DashboardOverviewSection 最小状态

- **WHEN** 后续实现讨论 `DashboardOverviewSection` 的分区状态
- **THEN** 它的最小状态必须为 `loading / ready / error`
- **AND** `loading` 对应 `overviewQueryState = pending`
- **AND** `ready` 对应 `overviewQueryState = success`
- **AND** `error` 对应 `overviewQueryState = error`
- **AND** 当前阶段不得为概览区块单独引入 `empty`

#### Scenario: 概览成功零计数的解释

- **WHEN** `DashboardOverviewRead` 成功且所有计数为 `0`
- **THEN** `DashboardOverviewSection` 仍必须进入 `ready`
- **AND** 冷启动空系统属于成功态，不得被解释为概览区块错误或空数据失败

### Requirement: CurrentFocusSection 与 AssetFeedbackSection 状态模型冻结

系统 SHALL 将 `CurrentFocusSection` 与 `AssetFeedbackSection` 冻结为共享 `FeedbackSignalRead` 的两块分区状态，而不是各自发明第二套读取状态。

#### Scenario: Feedback 分区最小状态

- **WHEN** 后续实现讨论 `CurrentFocusSection` 与 `AssetFeedbackSection` 的状态模型
- **THEN** 两个区块必须共享 `feedbackQueryState`
- **AND** 最小状态必须为 `loading / ready / empty / error`
- **AND** `loading` 对应 `feedbackQueryState = pending`
- **AND** `error` 对应 `feedbackQueryState = error`
- **AND** `ready` 与 `empty` 必须由 `feedbackQueryState = success` 后的数据结果派生

#### Scenario: Current Focus 成功空态

- **WHEN** `FeedbackSignalRead` 成功
- **AND** 不存在进入 `Current Focus / Next Action` 主队列的反馈卡片
- **THEN** `CurrentFocusSection` 必须进入 `empty`
- **AND** “无反馈信号”必须解释为成功空态
- **AND** 不得将该状态映射为错误或资源不存在

#### Scenario: Asset Feedback 成功空态

- **WHEN** `FeedbackSignalRead` 成功
- **AND** `Asset Feedback` 的代表性缺口项为空列表且缺口计数显式为 `0`
- **THEN** `AssetFeedbackSection` 必须进入 `empty`
- **AND** 必须继续展示成功空态语义，而不是局部失败

#### Scenario: Feedback 分区局部重试

- **WHEN** `feedbackQueryState = error`
- **THEN** `CurrentFocusSection` 与 `AssetFeedbackSection` 都必须在各自区块内容区域内呈现局部错误
- **AND** 任一区块触发重试时，必须重新执行同一次 `FeedbackSignalRead`
- **AND** 不得把局部重试升级为整页刷新

### Requirement: RecentActivitySection 状态模型冻结

系统 SHALL 将 `RecentActivitySection` 的分区状态冻结为依赖 `RecentActivityRead` 的独立活动流状态模型。

#### Scenario: Recent Activity 最小状态

- **WHEN** 后续实现讨论 `RecentActivitySection` 的状态模型
- **THEN** 最小状态必须为 `loading / ready / empty / error`
- **AND** `loading` 对应 `activityQueryState = pending`
- **AND** `error` 对应 `activityQueryState = error`
- **AND** `ready` 与 `empty` 必须由 `activityQueryState = success` 后的活动项是否为空派生

#### Scenario: Recent Activity 局部重试

- **WHEN** `RecentActivitySection` 进入 `error`
- **THEN** 错误反馈必须停留在活动区块内容区域内
- **AND** 区块重试只允许重新执行 `RecentActivityRead`
- **AND** 不得把活动流局部失败升级为整页失败

### Requirement: DashboardPrimaryActionPanel 状态模型冻结

系统 SHALL 将 `DashboardPrimaryActionPanel` 的状态冻结为单主 CTA 约束下的主动作状态机，而不是四个区块各自决定一个主按钮。

#### Scenario: 主 CTA 最小状态

- **WHEN** 后续实现讨论 `DashboardPrimaryActionPanel` 的状态模型
- **THEN** 最小状态必须为 `computing / ready / hidden / suppressed`
- **AND** `ready` 表示当前存在唯一主 CTA
- **AND** `hidden` 表示当前为成功中性态，不展示强制主 CTA
- **AND** `suppressed` 表示由于附属聚合局部失败，当前阶段禁止猜测或伪造反馈型主 CTA

#### Scenario: 主 CTA 的计算前提

- **WHEN** `DashboardOverviewRead` 仍未成功
- **THEN** `DashboardPrimaryActionPanel` 必须进入 `computing`
- **AND** 不得在主聚合未成功前猜测任何主 CTA

- **WHEN** `DashboardOverviewRead` 成功且已命中 `phase05-04` 顺序 1-4 的创建导向 CTA
- **THEN** `DashboardPrimaryActionPanel` 必须直接进入 `ready`
- **AND** 不需要等待 `FeedbackSignalRead` 或 `RecentActivityRead` 再决定是否显示

- **WHEN** `DashboardOverviewRead` 成功且未命中顺序 1-4
- **AND** `FeedbackSignalRead` 尚未成功
- **THEN** `DashboardPrimaryActionPanel` 必须继续保持 `computing`
- **AND** 不得提前猜测顺序 5-8 的反馈型主 CTA

#### Scenario: 主 CTA 的抑制与隐藏

- **WHEN** `DashboardOverviewRead` 成功且未命中顺序 1-4
- **AND** `FeedbackSignalRead = error`
- **THEN** `DashboardPrimaryActionPanel` 必须进入 `suppressed`
- **AND** 不得伪造顺序 5-8 的反馈型主 CTA
- **AND** 当前阶段只允许反馈区块内提供局部重试入口

- **WHEN** `DashboardOverviewRead` 成功
- **AND** 当前系统未命中 `phase05-04` 已冻结的顺序 1-4 创建导向 CTA
- **AND** `FeedbackSignalRead` 成功且不存在反馈缺口
- **AND** `RecentActivityRead` 成功且有活动数据或空活动数据
- **THEN** `DashboardPrimaryActionPanel` 必须进入 `hidden`
- **AND** 不得为了保持动作感而新增第二套主 CTA

### Requirement: Dashboard 成功态页面行为冻结

系统 SHALL 将 Dashboard 在冷启动空系统、非空缺口、有数据但无信号、无活动等成功态下的页面行为冻结为单值结论。

#### Scenario: 冷启动空系统行为

- **WHEN** `module_count = 0 && product_count = 0 && repository_count = 0 && decision_count = 0`
- **THEN** 整页必须进入 `ready`
- **AND** `DashboardOverviewSection` 为 `ready`
- **AND** `DashboardPrimaryActionPanel` 必须 `ready` 并命中 `Module Registry / Create`
- **AND** `CurrentFocusSection` 与 `AssetFeedbackSection` 可以根据 `FeedbackSignalRead` 进入 `loading / empty / error`
- **AND** 当前阶段不得把冷启动空系统解释为整页失败

#### Scenario: 非空但无 Module 行为

- **WHEN** `module_count = 0 && (product_count > 0 || repository_count > 0 || decision_count > 0)`
- **THEN** 整页必须进入 `ready`
- **AND** `DashboardPrimaryActionPanel` 必须 `ready` 并继续命中 `Module Registry / Create`
- **AND** 该状态必须与冷启动空系统区分展示

#### Scenario: 无反馈信号行为

- **WHEN** `DashboardOverviewRead` 成功
- **AND** `FeedbackSignalRead` 成功且 `Current Focus / Next Action` 结果为空
- **THEN** `CurrentFocusSection` 必须进入 `empty`
- **AND** `AssetFeedbackSection` 可独立进入 `ready` 或 `empty`
- **AND** 不得因为 `Current Focus` 为空就把整页回退到空系统语义

#### Scenario: 最近活动为空行为

- **WHEN** `RecentActivityRead` 成功但活动列表为空
- **THEN** `RecentActivitySection` 必须进入 `empty`
- **AND** 不得因为“暂无最近活动”而触发整页失败或创建导向主 CTA

### Requirement: Dashboard 局部失败与整页失败交互流冻结

系统 SHALL 将 Dashboard 的整页失败、局部失败、整页重试与分区重试交互流冻结为单值结论。

#### Scenario: 整页失败与整页重试

- **WHEN** `DashboardOverviewRead = error`
- **THEN** 页面必须进入 `page-error`
- **AND** `DashboardOverviewSection / CurrentFocusSection / AssetFeedbackSection / RecentActivitySection / DashboardPrimaryActionPanel` 都不得继续渲染成功语义
- **AND** 整页只允许提供整页级重试入口
- **AND** 整页重试必须重新执行 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead`

#### Scenario: 附属聚合局部失败

- **WHEN** `DashboardOverviewRead = success`
- **AND** `FeedbackSignalRead = error`
- **THEN** 页面必须保持 `ready`
- **AND** 只允许 `CurrentFocusSection / AssetFeedbackSection` 进入 `error`
- **AND** `DashboardPrimaryActionPanel` 必须遵守 `suppressed` 规则

- **WHEN** `DashboardOverviewRead = success`
- **AND** `RecentActivityRead = error`
- **THEN** 页面必须保持 `ready`
- **AND** 只允许 `RecentActivitySection` 进入 `error`
- **AND** 不得影响 `DashboardPrimaryActionPanel` 已命中的合法 CTA

### Requirement: 多缺口并存时的区块降级策略冻结

系统 SHALL 将多缺口并存时的区块降级策略冻结为“唯一主 CTA + 其他动作降级到区块内次级入口”。

#### Scenario: 多类缺口并存

- **WHEN** 同时存在创建导向缺口、决策缺口与资产缺口
- **THEN** `DashboardPrimaryActionPanel` 必须只展示 `phase05-04` 已冻结顺序命中的唯一主 CTA
- **AND** 其他反馈动作必须留在 `CurrentFocusSection` 或 `AssetFeedbackSection` 内部作为次级入口
- **AND** 不得并排展示多个同级主 CTA

#### Scenario: Asset Feedback 不升级为第二主队列

- **WHEN** `AssetFeedbackSection` 存在代表性缺口项
- **THEN** 这些缺口项只能作为区块内补充摘要入口存在
- **AND** 不得在页面级升级为并列主 CTA

### Requirement: Dashboard 返回链路与上下文恢复冻结

系统 SHALL 将从 Dashboard 跳出到 List / Detail / Create 后的返回链路冻结为“立即返回路径”和“外层返回 Dashboard”并存但不混写的单值规则。

#### Scenario: Dashboard 到列表再到详情的返回恢复

- **WHEN** 用户从 Dashboard 进入 `Module / Product / Repository / Decision` 列表
- **AND** 再从列表进入详情页
- **THEN** 详情页必须同时允许保留：
  - 列表原生来源模型（如 `fromList + queryText + statusFilter`）作为立即返回列表的路径
  - `fromDashboard + dashboardSection + dashboardReturnTo` 作为外层返回 Dashboard 的路径
- **AND** 不得把两者混写成同一个来源字段

#### Scenario: 从 owner 页面主动返回 Dashboard

- **WHEN** 用户在携带 `fromDashboard=true` 的 List / Detail / Create 页面主动触发“返回 Dashboard”
- **THEN** 必须返回 `/dashboard`
- **AND** 主动返回导航必须通过一次性路由导航状态承接同名 `dashboardSection` 值，作为落地 `/dashboard` 后恢复来源区块的唯一临时承接机制
- **AND** 必须继续以 `dashboardSection` 作为来源区块恢复依据
- **AND** 该一次性路由导航状态只服务本次主动返回落地，不得提升为 `/dashboard` 的搜索参数、持久化层或新的第二套来源参数命名
- **AND** 不得退化为回列表、回根路径 `/` 或回最近浏览器历史中的不确定页面

#### Scenario: 写入成功后的 reread 不等于返回 Dashboard

- **WHEN** 用户从 Dashboard 进入 `Product Detail`、`Repository Binding Detail / Workspace` 或 Create 页后发起既有写入
- **THEN** 写入成功后必须继续停留在 canonical owner 页面完成 reread
- **AND** 不得因为页面携带 `fromDashboard=true` 就自动跳回 Dashboard
- **AND** 只有用户后续主动选择返回 Dashboard 时，才允许使用 `dashboardReturnTo=/dashboard`

### Requirement: 刷新恢复与来源透传策略冻结

系统 SHALL 将 Dashboard 与目标页之间的刷新恢复与来源透传策略冻结为“URL 搜索参数是外层来源上下文唯一事实源”的单值结论。

#### Scenario: Dashboard Home 刷新恢复

- **WHEN** 用户刷新 `/dashboard`
- **THEN** 页面必须重新执行 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead`
- **AND** 必须基于最新读结果重新派生整页状态、分区状态与主 CTA 状态
- **AND** 若此前是通过主动返回落地到 `/dashboard`，一次性路由导航状态中的 `dashboardSection` 不得作为刷新后的恢复事实源继续保留
- **AND** 不得依赖 `sessionStorage / localStorage / 全局瞬时 store` 恢复上一次 Dashboard 页面状态

#### Scenario: 目标页刷新恢复

- **WHEN** 用户刷新携带 `fromDashboard=true` 的 List / Detail / Create 页面
- **THEN** 页面必须继续保留 `fromDashboard / dashboardSection / dashboardReturnTo`
- **AND** 必须以当前 URL 中的这些搜索参数作为 Dashboard 外层来源的唯一事实源
- **AND** 不得在刷新后回退到隐藏的内存状态或持久化层恢复 Dashboard 来源

#### Scenario: 来源参数缺失或非法时的恢复

- **WHEN** 页面携带 `fromDashboard=true`
- **AND** `dashboardSection` 非法或 `dashboardReturnTo` 缺失/非法
- **THEN** 必须沿用 `phase05-03` 已冻结的回退规则
- **AND** `dashboardSection` 非法时回退为 `overview`
- **AND** `dashboardReturnTo` 缺失或非法时回退为 `/dashboard`
- **AND** 不得静默跳到根路径 `/`

### Requirement: 页面级 UI 状态局部归属原则冻结

系统 SHALL 将当前阶段 Dashboard 的局部 UI 状态优先冻结为页面级或区块级归属，而不是直接升级为跨路由全局状态。

#### Scenario: 页面与区块局部状态归属

- **WHEN** 后续实现讨论分区展开/收起、局部重试按钮状态、CTA 临时禁用、区块空态提示等 UI 状态
- **THEN** 它们应优先归属于 `DashboardHomePage` 或对应分区/组件上下文
- **AND** 不得默认升级为跨路由全局状态

### Requirement: 运行时实现细节不冻结

系统 SHALL 明确区分“状态与交互语义”与“具体实现手段”，避免当前 spec 过早冻结实现细节。

#### Scenario: 当前阶段允许冻结的内容

- **WHEN** 当前 spec 讨论 Dashboard 状态与交互流
- **THEN** 可以冻结状态名称、状态归属、派生规则、错误呈现位置、重试粒度、返回路径与刷新恢复规则

#### Scenario: 当前阶段不得冻结的内容

- **WHEN** 后续实现尚未开始
- **THEN** 不得提前冻结具体 hook 命名、Query key、store API、缓存时间、请求取消、预取策略或 optimistic update 方案

## MODIFIED Requirements

### Requirement: phase05-05 页面与组件分层的状态语义解释

`phase05-05` 已冻结 `DashboardHomePageShell`、四区块与 `DashboardPrimaryActionPanel` 的组件树，`phase05-06` 在此基础上 SHALL 将这些组件进一步解释为各自承接的页面级与分区级状态模型，而不是仅停留在静态结构描述。

#### Scenario: 组件树到状态树的解释

- **WHEN** 后续实现根据 `phase05-05` 编写 Dashboard 前端页面
- **THEN** 必须同时满足 `phase05-05` 的组件归属与本规格冻结的状态模型
- **AND** 不得只保留静态组件树而缺失整页与分区状态语义

### Requirement: phase05-03 来源参数的状态语义解释

`phase05-03` 已冻结 `fromDashboard / dashboardSection / dashboardReturnTo` 的参数语义，`phase05-06` 在此基础上 SHALL 将其进一步解释为返回链路、刷新恢复与外层来源上下文的唯一事实源。

#### Scenario: 来源参数到交互流的解释

- **WHEN** 后续实现修改 List / Detail / Create 页的返回逻辑或刷新恢复逻辑
- **THEN** 必须沿用 `fromDashboard / dashboardSection / dashboardReturnTo` 作为唯一 Dashboard 外层来源参数集合
- **AND** 不得在页面状态层另起第二套 Dashboard 来源命名

### Requirement: phase05-04 聚合错误语义的前端解释

`phase05-04` 已冻结主聚合失败触发整页失败、附属聚合失败只触发局部失败的错误语义前提，`phase05-06` 在此基础上 SHALL 将其进一步解释为前端整页状态、分区状态、主 CTA 抑制与重试粒度。

#### Scenario: 错误语义到前端状态的解释

- **WHEN** 后续实现讨论整页错误、分区错误与主 CTA 展示
- **THEN** 必须同时满足 `phase05-04` 的错误语义前提与本规格冻结的状态模型
- **AND** 不得在前端额外发明第二套“局部失败也整页失败”或“反馈失败仍猜测主 CTA”的解释

## REMOVED Requirements

### Requirement: 以持久化层恢复 Dashboard 页面状态

**Reason**: `phase05` 当前阶段已经冻结 URL 搜索参数作为 Dashboard 外层来源上下文的唯一事实源，`/dashboard` 本身又不承接筛选参数；继续引入 `sessionStorage / localStorage` 作为刷新恢复或返回恢复的事实源，会形成第二套不可见状态模型。

**Migration**: 若后续确需引入持久化恢复能力，必须进入新的 `phase / fix / audit` 重新冻结，不在 `phase05-06` 当前规格中处理。
