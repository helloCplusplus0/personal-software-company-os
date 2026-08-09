# Tasks

- [x] Task 1: 冻结 Dashboard 整页查询状态、整页视图状态与主/附属聚合的前端角色。把整页是否成功、何时整页失败、何时进入 ready 写成单值结论。
  - [x] SubTask 1.1: 明确 `overviewQueryState / feedbackQueryState / activityQueryState` 三类最小查询状态
  - [x] SubTask 1.2: 明确整页视图状态冻结为 `initial-loading / ready / page-error`
  - [x] SubTask 1.3: 明确 `DashboardOverviewRead` 是主聚合读取，`FeedbackSignalRead / RecentActivityRead` 是附属聚合读取

- [x] Task 2: 冻结四个分区与 `DashboardPrimaryActionPanel` 的状态模型。把概览、反馈、活动和主 CTA 在前端如何派生状态写成可直接实现的结果。
  - [x] SubTask 2.1: 明确 `DashboardOverviewSection` 的最小状态为 `loading / ready / error`
  - [x] SubTask 2.2: 明确 `CurrentFocusSection` 与 `AssetFeedbackSection` 共享 `FeedbackSignalRead`，最小状态为 `loading / ready / empty / error`
  - [x] SubTask 2.3: 明确 `RecentActivitySection` 的最小状态为 `loading / ready / empty / error`
  - [x] SubTask 2.4: 明确 `DashboardPrimaryActionPanel` 的最小状态为 `computing / ready / hidden / suppressed`

- [x] Task 3: 冻结 Dashboard 各类成功态下的页面行为。把冷启动空系统、非空但无 Module、无反馈信号、无活动等场景的页面行为写成单值结论。
  - [x] SubTask 3.1: 明确冷启动空系统与非空但无 Module 的整页行为与主 CTA 表现
  - [x] SubTask 3.2: 明确无反馈信号时 `CurrentFocusSection` 的成功空态行为
  - [x] SubTask 3.3: 明确无最近活动时 `RecentActivitySection` 的成功空态行为
  - [x] SubTask 3.4: 明确无缺口且有活动/无活动时 `DashboardPrimaryActionPanel` 必须隐藏

- [x] Task 4: 冻结整页失败、局部失败、整页重试与分区重试交互流。保证 `phase05-04` 的错误语义被前端状态机正确承接。
  - [x] SubTask 4.1: 明确 `DashboardOverviewRead` 失败触发 `page-error` 与整页级重试
  - [x] SubTask 4.2: 明确 `FeedbackSignalRead` 失败只触发 `CurrentFocusSection / AssetFeedbackSection` 的局部错误
  - [x] SubTask 4.3: 明确 `RecentActivityRead` 失败只触发 `RecentActivitySection` 的局部错误
  - [x] SubTask 4.4: 明确 Feedback 分区重试复用同一次 `FeedbackSignalRead`，Recent Activity 分区重试只重试 `RecentActivityRead`

- [x] Task 5: 冻结主 CTA 优先级命中规则、多缺口并存时的区块降级策略与 CTA 抑制规则。确保页面级只存在一个主 CTA。
  - [x] SubTask 5.1: 明确主 CTA 在 `overview` 成功前处于 `computing`
  - [x] SubTask 5.2: 明确顺序 1-4 的创建导向 CTA 一旦命中即可直接 `ready`
  - [x] SubTask 5.3: 明确顺序 5-8 的反馈型主 CTA 必须等待 `FeedbackSignalRead` 成功后才能命中
  - [x] SubTask 5.4: 明确 `FeedbackSignalRead` 失败且 overview 未命中 1-4 时，主 CTA 必须进入 `suppressed`
  - [x] SubTask 5.5: 明确多类缺口并存时只允许一个主 CTA，其他动作降级为区块内次级入口
  - [x] SubTask 5.6: 明确 `hidden` 只适用于未命中顺序 1-4 创建导向 CTA 的成功中性态，不得误吞冷启动、无 Module、缺 Product 或缺 Repository 场景

- [x] Task 6: 冻结从 Dashboard 跳出后的返回链路、上下文恢复与刷新恢复策略。把 `fromDashboard / dashboardSection / dashboardReturnTo` 的状态语义与 URL 事实源模型写成单值结论。
  - [x] SubTask 6.1: 明确 Dashboard 到 List 再到 Detail 时，列表原生来源模型与 Dashboard 外层来源模型可以并存但不得混写
  - [x] SubTask 6.2: 明确携带 `fromDashboard=true` 的 List / Detail / Create 页面主动返回时必须回 `/dashboard`
  - [x] SubTask 6.3: 明确写入成功后的 reread 不等于自动返回 Dashboard
  - [x] SubTask 6.4: 明确 `/dashboard` 刷新时必须重跑三类读取，不依赖持久化层恢复页面状态
  - [x] SubTask 6.5: 明确目标页刷新时，`fromDashboard / dashboardSection / dashboardReturnTo` 仍以当前 URL 为唯一事实源
  - [x] SubTask 6.6: 明确非法 `dashboardSection / dashboardReturnTo` 的回退规则沿用 `phase05-03`
  - [x] SubTask 6.7: 明确主动返回 `/dashboard` 时，来源区块恢复通过一次性路由导航状态承接同名 `dashboardSection`，且该承接位不提升为 `/dashboard` 搜索参数或持久化层事实源

- [x] Task 7: 冻结局部 UI 状态归属原则与运行时实现细节不冻结边界。避免把当前规格写成具体 hook/store 方案。
  - [x] SubTask 7.1: 明确分区展开/收起、局部重试、CTA 临时禁用等 UI 状态优先归属于页面或区块上下文
  - [x] SubTask 7.2: 明确当前阶段不冻结 hook 命名、Query key、store API、缓存时间、请求取消与 optimistic update
  - [x] SubTask 7.3: 明确不允许以 `sessionStorage / localStorage` 作为 Dashboard 页面状态恢复的事实源

- [x] Task 8: 完成 phase05-06 规格一致性校验。确认本次冻结与 phase05 三件套、phase05-03 / 04 / 05 保持一致，并足以直接进入实现。
  - [x] SubTask 8.1: 验证整页失败/局部失败与 `phase05-04` 的错误语义前提一致
  - [x] SubTask 8.2: 验证返回链路、刷新恢复与 `phase05-03` 的来源参数命名与回退规则一致
  - [x] SubTask 8.3: 验证状态模型与 `phase05-05` 的页面/组件 owner 保持一致，不另起第二套组件归属
  - [x] SubTask 8.4: 验证主 CTA 命中、抑制与单主 CTA 约束和 `shared_baseline`、`phase05-04` 一致
  - [x] SubTask 8.5: 验证当前规格未越界冻结 `phase05-07 / 08` 的后端边界与 `.proto` 设计职责
  - [x] SubTask 8.6: 验证设计结果足以直接进入实现

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`
- `Task 6` depends on `Task 1`, `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 6`
- `Task 8` depends on `Task 1` through `Task 7`
