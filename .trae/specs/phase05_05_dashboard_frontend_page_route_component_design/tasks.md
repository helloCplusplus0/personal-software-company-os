# Tasks

- [x] Task 1: 冻结 Dashboard 的页面文件落点与主导航接入方式。把 `Dashboard Home` 的页面文件、路由文件和导航入口写成单值结论。
  - [x] SubTask 1.1: 明确 `DashboardHomePage` 文件路径遵循 `frontend/src/features/dashboard/pages/` 模式
  - [x] SubTask 1.2: 明确 `DashboardRoute -> /dashboard -> frontend/src/routes/dashboard.tsx` 的单值路由结论
  - [x] SubTask 1.3: 明确 `frontend/src/routes/__root.tsx` 需要新增一级 `Dashboard` 导航入口，且不替换既有导航主线

- [x] Task 2: 冻结 Dashboard Home 的页面级、区块级、卡片级与列表级组件树。把四区块结构落到可直接实现的前端组件分层。
  - [x] SubTask 2.1: 明确 `DashboardHomePageShell`、四个区块级组件 `CurrentFocusSection / DashboardOverviewSection / AssetFeedbackSection / RecentActivitySection` 与 `DashboardPrimaryActionPanel` 的页面级组件树顺序
  - [x] SubTask 2.2: 明确 `DashboardOverviewCardGrid / DashboardOverviewCard` 的职责边界
  - [x] SubTask 2.3: 明确 `FeedbackSignalCardList / FeedbackSignalCard` 的职责边界
  - [x] SubTask 2.4: 明确 `AssetFeedbackList / AssetFeedbackItemCard` 的职责边界
  - [x] SubTask 2.5: 明确 `RecentActivityList / RecentActivityItemCard` 的职责边界
  - [x] SubTask 2.6: 明确 `DashboardPrimaryActionPanel` 的职责边界，承接主 CTA 优先级矩阵与空状态按钮可点模式

- [x] Task 3: 冻结 Dashboard 到既有 List / Detail / Create 页面时的最小路由与参数承接方式。把 Dashboard 来源参数在路由层如何进入目标页写成单值结论。
  - [x] SubTask 3.1: 明确 `/dashboard` 当前阶段不引入新的搜索参数或区块级 URL 状态
  - [x] SubTask 3.2: 明确 `modules / products / repositories / decisions` 的 List 路由在保留原生搜索参数的同时承接 `fromDashboard / dashboardSection / dashboardReturnTo`
  - [x] SubTask 3.3: 明确 `Module / Product / Repository / Decision` Detail 路由继续用路径参数承接对象身份，用搜索参数承接 Dashboard 来源参数
  - [x] SubTask 3.4: 明确 `Module / Product / Repository` Create 路由允许承接 Dashboard 来源参数，但不新增 Dashboard 私有路径

- [x] Task 4: 冻结 Dashboard 组件归属原则与复用边界。避免在实现前过早抽象新的共享组件层。
  - [x] SubTask 4.1: 明确 Dashboard 组件默认归属于 `frontend/src/features/dashboard/components/`
  - [x] SubTask 4.2: 明确只有存在跨 feature 复用证据时才允许抽为共享组件
  - [x] SubTask 4.3: 明确当前阶段不提前拆出 Dashboard 专属基础组件层

- [x] Task 5: 冻结 `PC / 移动浏览器` 双场景下的布局降级策略。确保单一 `React Web` 页面可直接进入实现。
  - [x] SubTask 5.1: 明确 `PC` 环境下四区块可同屏承接，但 `Current Focus` 始终保持第一行动优先级
  - [x] SubTask 5.2: 明确移动浏览器下的区块顺序为 `Current Focus / dashboard_overview / Asset Feedback / Recent Activity`
  - [x] SubTask 5.3: 明确通过单列垂直重排、信息裁剪与列表收束解决窄屏拥挤，不引入第二套移动端 UI

- [x] Task 6: 冻结 phase05-05 的非目标边界。确认本规格不越界抢 `phase05-06 / 07 / 08` 的职责。
  - [x] SubTask 6.1: 明确页面级状态模型、局部 loading/error/empty 状态机与返回恢复细节由 `phase05-06` 承接
  - [x] SubTask 6.2: 明确后端模块边界、接口分组与 owner 由 `phase05-07` 承接
  - [x] SubTask 6.3: 明确 `.proto` 服务命名、消息结构与过渡传输层映射由 `phase05-08` 承接
  - [x] SubTask 6.4: 明确 hook 命名、Query key、store API 与缓存时间等运行时细节不在本规格冻结

- [x] Task 7: 完成 phase05-05 规格一致性校验。确认本次冻结与 phase05 三件套、phase05-01 / 03 / 04 和现有前端目录模式保持一致。
  - [x] SubTask 7.1: 验证页面文件落点、路由文件与 `frontend/src/routes` / `frontend/src/features` 的既有命名模式一致
  - [x] SubTask 7.2: 验证四区块组件树与 `phase05-01` 的固定归属关系一致
  - [x] SubTask 7.3: 验证路由参数承接方式与 `phase05-03` 的来源参数命名一致，不形成第二套参数体系
  - [x] SubTask 7.4: 验证布局降级策略与 `phase05-01` 的信息层级结论一致，且未引入第二套移动端 UI
  - [x] SubTask 7.5: 验证本规格未越界冻结 `phase05-06 / 07 / 08` 的职责

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 5` depends on `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 7` depends on `Task 1` through `Task 6`
