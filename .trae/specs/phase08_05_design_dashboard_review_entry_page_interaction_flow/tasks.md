# Tasks

- [x] Task 1: 冻结 Dashboard 标题行动区的 review launcher 形态
  - [x] SubTask 1.1: 明确 `DashboardPrimaryActionPanel` 从 `phase05` 单主 CTA 命中器迁移为 `Daily Review / Weekly Review` 双入口承接位
  - [x] SubTask 1.2: 明确双入口与 `OnboardingCtaButton / SovereigntyPanel` 的共存方式，不破坏 `/dashboard` 首屏信息密度
  - [x] SubTask 1.3: 明确 PC 与移动浏览器下标题行动区的最小布局降级
  - [x] SubTask 1.4: 明确标题行动区的响应式重排责任由页面壳层承接，而不是散落到各按钮组件临场处理

- [x] Task 2: 冻结 Daily Review 的页面流、区块顺序与完成语义
  - [x] SubTask 2.1: 明确 `Daily Review` 的页面头部、主动离开入口与页面级壳层要求
  - [x] SubTask 2.2: 明确 `current focus / pending decisions / representative signals` 的区块顺序与主次关系
  - [x] SubTask 2.3: 明确 `Daily Review` 的最小成功会话如何进入 canonical 下一步动作
  - [x] SubTask 2.4: 明确 `pending decisions` 来源于既有 `Decision` canonical fact 的最小范围，并冻结 top N 摘要展示口径

- [x] Task 3: 冻结 Weekly Review 的页面流、区块顺序与 reuse snapshot 消费方式
  - [x] SubTask 3.1: 明确 `Weekly Review` 的页面头部、主动离开入口与页面级壳层要求
  - [x] SubTask 3.2: 明确 `overview / recent activity / representative signals / reuse snapshot` 的区块顺序与主次关系
  - [x] SubTask 3.3: 明确 `module_reuse_summary / capability_summary` 在同一 `reuse snapshot` 区块内的正式消费方式

- [x] Task 4: 冻结 daily / weekly 的共享壳层与移动浏览器降级策略
  - [x] SubTask 4.1: 明确允许复用的 shared shell 边界
  - [x] SubTask 4.2: 明确禁止把 daily / weekly 退化成同一份仅标题不同的模板
  - [x] SubTask 4.3: 明确移动浏览器下的单列重排、信息裁剪与紧凑按钮约束

- [x] Task 5: 冻结 review 完成、handoff 与返回链
  - [x] SubTask 5.1: 明确主动离开 review 时统一返回 `/dashboard`
  - [x] SubTask 5.2: 明确 `decision handoff` 进入既有 `Decision` canonical 页的页面流
  - [x] SubTask 5.3: 明确 `entity handoff` 进入既有 `Product / Module / Repository` canonical 页的页面流
  - [x] SubTask 5.4: 明确 `next-step result` 只保留为后续 `review record` 出口，不在本任务中发明第二套 sink
  - [x] SubTask 5.5: 明确 review route 必须承接并透传 `fromDashboard / dashboardSection / dashboardReturnTo`，确保 canonical 页返回链可执行

- [x] Task 6: 校验真实 inventory 的复用 / 扩展 / 禁止升级矩阵
  - [x] SubTask 6.1: 明确 `DashboardHomePage / DashboardPrimaryActionPanel / dashboard-source.ts / BackToDashboardButton` 的复用或扩展身份
  - [x] SubTask 6.2: 明确 `CurrentFocusSection / AssetFeedbackSection / RecentActivitySection` 只复用语义与展示模式，不升级为会话宿主或写路径 owner
  - [x] SubTask 6.3: 明确 `FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar / OnboardingCtaButton / SovereigntyPanel` 的负向约束

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 2`
- `Task 5` depends on `Task 3`
- `Task 6` depends on `Task 1`
