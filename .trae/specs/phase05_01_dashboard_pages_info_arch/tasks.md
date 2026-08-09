# Tasks

- [x] Task 1: 冻结 `Dashboard Home` 页面主线与页面职责。把 `Dashboard Home` 收敛为当前阶段唯一页面主线，并明确它只承接聚合读取、空状态引导与跳转入口。
  - [x] SubTask 1.1: 明确 `Dashboard Home` 是 `phase05-01` 当前阶段唯一页面主线
  - [x] SubTask 1.2: 明确 `Dashboard Home` 只承接概览读取、主行动队列、补充摘要、最近活动与空状态 CTA
  - [x] SubTask 1.3: 明确 `Dashboard Home` 不承接新的业务写入入口

- [x] Task 2: 冻结 `Dashboard Home` 的四区块归属与第一屏信息层级。把 `dashboard_overview / Current Focus（内部承接 Next Action 主队列） / Asset Feedback / Recent Activity` 写成单值结论。
  - [x] SubTask 2.1: 明确 `dashboard_overview` 只承接概览卡片区块
  - [x] SubTask 2.2: 明确 `Current Focus` 内部承接 `Next Action` 主队列，不形成第二个并列区块
  - [x] SubTask 2.3: 明确 `Asset Feedback` 只承接 `product_asset_coverage` 的补充摘要
  - [x] SubTask 2.4: 明确 `Recent Activity` 只承接独立活动流，不与反馈信号混合排序
  - [x] SubTask 2.5: 明确第一屏优先暴露 `Current Focus / Next Action`

- [x] Task 3: 冻结 `Dashboard` 的正式入口路由与主导航归属。把 `/dashboard`、`/` 与根级布局宿主之间的关系写成单值结论。
  - [x] SubTask 3.1: 明确 `Dashboard Home` 的正式业务入口为 `/dashboard`
  - [x] SubTask 3.2: 明确 `Dashboard` 作为既有主导航中的一级入口新增
  - [x] SubTask 3.3: 明确当前阶段不把 `/` 单值解释为 `Dashboard Home`
  - [x] SubTask 3.4: 明确根级布局宿主继续只承接全局导航与页面容器语义

- [x] Task 4: 冻结 `PC / 移动浏览器` 双场景下的信息密度与布局降级策略。保持单一 `React Web` 语义，同时明确窄屏的降级方式。
  - [x] SubTask 4.1: 明确桌面端可以同屏承接四类区块的较高信息密度
  - [x] SubTask 4.2: 明确移动浏览器采用区块垂直重排、信息裁剪、折叠或延后展示
  - [x] SubTask 4.3: 明确当前阶段不引入第二套移动端 UI、独立 `React Native` 客户端或完整 `PWA`

- [x] Task 5: 完成 `phase05-01` 规格一致性校验。确认本次规格与 `phase05` 三件套、`phase04` 直接上游与 `phase01-06` MVP 页面范围保持一致。
  - [x] SubTask 5.1: 验证页面范围、区块归属与 `phase05_dashboard_feedback_foundation_architecture_plan.md` 一致
  - [x] SubTask 5.2: 验证任务目标与 DoD 和 `phase05_dashboard_feedback_foundation_dev_plan.md` 一致
  - [x] SubTask 5.3: 验证首页入口、导航归属、空状态前提与 `phase05_dashboard_feedback_foundation_shared_baseline.md` 一致
  - [x] SubTask 5.4: 验证本次规格未超出 `PSCO-summarize-feedback.md` 与 `phase01-06` 的 MVP 页面范围

- [x] Task 6: 冻结 `Dashboard Home` 四类区块的可点击入口热区规则。把"整卡可点"与"仅 `action_label` 可点"的单值结论写入规格。
  - [x] SubTask 6.1: 明确 `dashboard_overview` 中仅 `module_count / product_count / repository_count / decision_count` 作为独立整卡可点击概览卡片，并分别跳转到对应 List 页面
  - [x] SubTask 6.1a: 明确 `product_with_repository_count / product_with_module_count` 当前阶段只作为辅助指标展示，不形成独立可点击卡片，也不引入新的 `Product List` 筛选态
  - [x] SubTask 6.2: 明确 `Current Focus / Next Action` 反馈信号卡片整卡可点，`action_label` 作为视觉锚点
  - [x] SubTask 6.3: 明确 `Asset Feedback` 缺口项整卡可点，`action_label` 作为视觉锚点
  - [x] SubTask 6.4: 明确 `Recent Activity` 活动项整卡可点，对象名称作为视觉锚点
  - [x] SubTask 6.5: 明确空状态主 CTA 作为按钮可点，不采用整卡模式
  - [x] SubTask 6.6: 明确"整卡可点"是四类区块的统一基础规则，不得在同一区块内混合两种模式

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 6`
