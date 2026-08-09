# Tasks

- [x] Task 1: 冻结 `dashboard_overview`、反馈卡片、最近活动与空状态 CTA 的最小跳转矩阵。把每类可点击入口的 canonical owner 页面写成单值结论。
  - [x] SubTask 1.1: 明确 `module_count / product_count / repository_count / decision_count` 四类概览卡片分别跳转到对应 List 页面
  - [x] SubTask 1.2: 明确 `pending_decision_signals` 的单项信号与聚合信号分别落到 `Decision Detail` 与 `Decision Center / List`
  - [x] SubTask 1.3: 明确三类 `product_asset_coverage` 缺口信号统一落到对应 `Product Detail`
  - [x] SubTask 1.4: 明确冷启动空状态、非空缺口态与高优先级主 CTA 的跳转目标与优先级顺序

- [x] Task 2: 冻结 `recent_activity_feed` 的活动类型落点与 `Binding` 子类型拆分。确保 `Release` 与 `Binding` 不再保留跳转歧义。
  - [x] SubTask 2.1: 明确 `module / product / repository / decision` 四类活动分别落到既有详情页
  - [x] SubTask 2.2: 明确 `release` 活动统一回落到所属 `Module Detail`，当前阶段不新增 `Release Detail`
  - [x] SubTask 2.3: 明确 `product_module_binding / product_repository_binding / module_repository_binding` 三类绑定活动的单值落点与目标身份字段
  - [x] SubTask 2.4: 明确移除笼统 `binding` 活动类型，禁止保留未拆分语义

- [x] Task 3: 冻结 `fromDashboard` 来源标记、来源区块标记与最小返回参数模型。把 Dashboard 来源上下文写成后续实现唯一可用的参数集合。
  - [x] SubTask 3.1: 明确 `fromDashboard=true`、`dashboardSection`、`dashboardReturnTo=/dashboard` 是最小参数集合
  - [x] SubTask 3.2: 明确 `dashboardSection` 允许值仅为 `overview / current-focus / asset-feedback / recent-activity / empty-state`
  - [x] SubTask 3.3: 明确 `fromDashboard` 是 Dashboard 外层来源标记，不替代目标页原生来源上下文
  - [x] SubTask 3.4: 明确不得发明 `dashboardFrom`、`returnToDashboard` 等第二套命名

- [x] Task 4: 冻结从 Dashboard 跳出后的主动返回、刷新恢复与多跳上下文承接规则。把“返回 Dashboard”与“写入成功后 reread”区分清楚。
  - [x] SubTask 4.1: 明确从 Dashboard 直接进入目标页后，主动返回 Dashboard 时必须回到 `/dashboard`
  - [x] SubTask 4.2: 明确 `Dashboard -> Module/Product/Repository/Decision List -> Detail` 多跳场景下，`fromList` 负责立即返回路径，`fromDashboard` 负责外层返回路径
  - [x] SubTask 4.3: 明确带 `fromDashboard` 的目标页刷新后必须恢复来源标记与返回路径
  - [x] SubTask 4.4: 明确写入成功后继续遵守 `phase04` 的 canonical owner + reread 规则，不自动跳回 Dashboard
  - [x] SubTask 4.5: 明确非法 `dashboardSection` 或缺失/非法 `dashboardReturnTo` 的回退语义
  - [x] SubTask 4.6: 明确 Dashboard 来源下 Create 页的取消返回 `/dashboard`、提交成功后 Detail 保留 `fromDashboard` 回流与外层来源标记保留规则
  - [x] SubTask 4.7: 明确高优先级反馈主 CTA 的 `dashboardSection` 必须回落到真实来源 `current-focus`，不得误写成 `empty-state`

- [x] Task 5: 完成 `phase05-03` 规格一致性校验。确认本次冻结与 `phase05` 三件套、`phase05-01 / 02` 及 `phase04-03 / 06` 上游规则保持一致。
  - [x] SubTask 5.1: 验证跳转矩阵与 `phase05_dashboard_feedback_foundation_architecture_plan.md`、`shared_baseline.md` 保持一致
  - [x] SubTask 5.2: 验证来源参数命名与返回规则不引入第二套状态命名体系
  - [x] SubTask 5.3: 验证 `fromDashboard` 未覆盖 `phase04-06` 已冻结的页面原生来源模型
  - [x] SubTask 5.4: 验证 Dashboard 仍只承接读与跳转，不形成第二套写入工作台

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`
