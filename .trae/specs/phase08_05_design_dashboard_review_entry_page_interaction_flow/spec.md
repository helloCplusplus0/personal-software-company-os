# phase08-05 产出 Dashboard review 入口、页面流与交互流设计 Spec

## Why

`phase08-01 ~ 04` 已经冻结了 review loop 的范围、入口边界、动作 owner 与合同边界，但当前仍缺少一份可直接进入实现的页面流与交互流设计，来说明 `Dashboard` 如何稳定承接 daily / weekly review 双入口，以及 review 完成后如何回流既有 canonical 页面而不长出第二套宿主。
如果这一步继续模糊，后续实现很容易回到“在 `/dashboard` 临时多塞两块内容”或“review 页面自己发明一套返回与完成语义”的状态，最终既破坏 `Dashboard` 的信息密度，也无法形成 daily / weekly 两条可单独验收的成功会话。

## What Changes

- 冻结 `Dashboard` 标题行动区到 daily / weekly review 的正式入口形态与进入交互
- 冻结 `Daily Review` 与 `Weekly Review` 的最小页面骨架、信息编排顺序与共享壳层边界
- 冻结 `current focus / representative signals / pending decisions / reuse snapshot` 在 review 页中的消费方式
- 冻结 review 完成、主动离开、进入 canonical 页面后的回流链与返回链
- 冻结移动浏览器下的最小降级策略，保证不引入第二套移动端页面体系
- 冻结真实 inventory 的复用 / 扩展 / 禁止升级矩阵，避免后续实现重新发明宿主或写路径 owner

## Impact

- Affected specs:
  - `phase08_01_freeze_operating_review_loop_scope_success_non_goals`
  - `phase08_02_freeze_dashboard_review_entry_page_route_handoff`
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `phase08_04_freeze_review_contract_read_record_boundary`
  - `phase05_05_dashboard_frontend_page_route_component_design`
  - `phase05_06_dashboard_frontend_state_interaction_flow`
- Affected code:
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
  - `frontend/src/features/dashboard/components/current-focus-section.tsx`
  - `frontend/src/features/dashboard/components/asset-feedback-section.tsx`
  - `frontend/src/features/dashboard/components/recent-activity-section.tsx`
  - `frontend/src/features/dashboard/lib/dashboard-source.ts`
  - `frontend/src/features/dashboard/components/back-to-dashboard-button.tsx`
  - 后续新增的 review route、review page 与 review shell
  - `frontend/src/features/decision-center/pages/*`
  - `frontend/src/features/module-registry/pages/*`
  - `frontend/src/features/product-registry/pages/*`
  - `frontend/src/features/repository-binding/pages/*`

## ADDED Requirements

### Requirement: Dashboard 标题行动区必须演进为稳定可见的 review launcher

系统 SHALL 将 `Dashboard` 标题行动区设计为 stable review launcher，用于从 `/dashboard` 进入 `Daily Review` 与 `Weekly Review`，同时不破坏既有总览页的信息密度。

#### Scenario: Dashboard 正式 review 入口形态

- **WHEN** 用户进入 `/dashboard`
- **THEN** 标题行动区必须同时提供 `开始 Daily Review` 与 `开始 Weekly Review` 两个显式入口
- **AND** 这两个入口必须继续位于 `DashboardPrimaryActionPanel` 所在的标题行动区，而不是迁移到 `CurrentFocusSection / AssetFeedbackSection / RecentActivitySection`
- **AND** 入口文案必须直接表达“进入 review 会话”，不得继续沿用 `phase05` 单主 CTA 的创建导向语义冒充 review 入口
- **AND** review launcher 必须保持稳定可见，不得重新退化成依赖 overview / feedback 命中结果才出现的条件按钮

#### Scenario: Dashboard 标题行动区的信息密度约束

- **WHEN** 后续实现调整标题行动区
- **THEN** 必须继续与 `OnboardingCtaButton`、`SovereigntyPanel` 共存于 `Dashboard` 首屏语义内
- **AND** 不得因为新增 daily / weekly review 双入口而把首屏改造成大块工作台、第二行工具栏或多层 tab
- **AND** PC 下允许双按钮并排或主次按钮组合
- **AND** 移动浏览器下必须降级为纵向堆叠或紧凑双按钮，不得压缩成难以点击的微型入口

#### Scenario: 标题行动区的响应式布局责任

- **WHEN** 后续实现为标题行动区落地双 review 入口
- **THEN** `DashboardHomePageShell` 或等价标题壳层必须成为响应式重排责任位
- **AND** 不得把“是否换行、是否堆叠、是否缩小按钮”下放给 `DashboardPrimaryActionPanel`、`OnboardingCtaButton` 各自临场处理
- **AND** PC 下允许标题与行动区同一行呈现，但行动区必须保持 `Daily Review / Weekly Review` 为主按钮组
- **AND** 移动浏览器下允许标题区扩展为两层：第一层只承接页面标题，第二层承接 review launcher
- **AND** `OnboardingCtaButton` 若在移动浏览器下仍需显示，必须降级为 review launcher 下方的次级入口，不得与双 review 按钮同排并列
- **AND** `SovereigntyPanel` 必须继续保持为标题区下方的独立全宽区块，不得被吸入 review launcher 行内

### Requirement: Daily Review 页面必须围绕“当下要处理什么”组织

系统 SHALL 将 `Daily Review` 设计为短会话页面，围绕当前焦点、待处理决策与代表性补充信号组织信息，而不是 Dashboard 的另一种排版。

#### Scenario: Daily Review 最小页面骨架

- **WHEN** 用户进入 `Daily Review`
- **THEN** 页面至少必须包含以下区块顺序：
  - review 页面头部
  - `current focus` 主区
  - `pending decisions` 区
  - `representative signals` 补充区
  - `完成 / 进入下一步动作` 区
- **AND** 页面头部必须明确显示当前为 `Daily Review`
- **AND** 页面头部必须提供返回 `/dashboard` 的主动离开入口
- **AND** 当前阶段不得把 `Recent Activity` 或 `reuse snapshot` 升级为 `Daily Review` 主区

#### Scenario: Daily Review 的信息组织优先级

- **WHEN** `Daily Review` 渲染 review context
- **THEN** `current_focus_signals` 必须占据页面最高信息优先级
- **AND** `pending decisions` 必须作为紧邻主区的第二优先级区块存在
- **AND** `representative_signals` 只作为补充缺口输入存在，不得与 `current focus` 形成第二条同级主队列
- **AND** 页面完成语义必须偏向“形成一个明确下一动作并立即进入 canonical 页面”

#### Scenario: Daily Review 中 pending decisions 的真实来源与最小展示口径

- **WHEN** 后续实现为 `Daily Review` 落地 `pending decisions` 区
- **THEN** 该区块必须来源于 `phase08-04` 已冻结的“既有 `Decision` canonical fact 中可解释为 pending / backlog 的最小范围”
- **AND** 该区块不得从 `current_focus_signals`、`representative_signals` 或 `recent activity` 反推生成
- **AND** 最小展示形态必须是“pending / backlog 决策摘要列表（top N）+ 进入既有 `Decision` canonical 页的动作入口”或与之等价的紧凑摘要区
- **AND** 该摘要列表至少必须承接 `decision_id`、标题或等价主标签、最小状态语义
- **AND** 不得把完整 `Decision Detail` 表单、完整决策上下文或新的 review-local decision shadow state 直接内嵌进 `Daily Review`
- **AND** 若需“查看全部”，只能继续进入既有 `Decision List`，不得在 `Daily Review` 内并排长出第二套 pending decision 工作台

### Requirement: Weekly Review 页面必须围绕“周期整理与归纳”组织

系统 SHALL 将 `Weekly Review` 设计为周期性整理页面，正式消费 overview、recent activity、representative signals 与 reuse snapshot，而不是 `Daily Review` 的简单扩容版。

#### Scenario: Weekly Review 最小页面骨架

- **WHEN** 用户进入 `Weekly Review`
- **THEN** 页面至少必须包含以下区块顺序：
  - review 页面头部
  - `overview` 摘要区
  - `recent activity` 区
  - `representative signals` 区
  - `reuse snapshot` 区
  - `完成 / 进入下一步动作` 区
- **AND** 页面头部必须明确显示当前为 `Weekly Review`
- **AND** 页面头部必须提供返回 `/dashboard` 的主动离开入口
- **AND** 页面不得退化为只读报表页，仍必须承接进入下一步动作的交互出口

#### Scenario: Weekly Review 对 reuse snapshot 的正式消费方式

- **WHEN** `Weekly Review` 渲染 `ReuseSummaryRead(scope = dashboard)` 的结果
- **THEN** 必须正式消费 `module_reuse_summary` 与 `capability_summary`
- **AND** 二者必须作为同一个 `reuse snapshot` 区块内的两个摘要分组被呈现
- **AND** 不得为 weekly review 再发明第二套 `reuse feed` 或复制 `AssetFeedbackSection` 的完整区块壳层
- **AND** 若 `reuse snapshot` 读取失败，只允许在 weekly review 的该区块内呈现局部失败，不得把整页回退为 page error

### Requirement: Daily / Weekly Review 必须复用统一 page shell，但保持区块语义差异

系统 SHALL 为 `Daily Review` 与 `Weekly Review` 设计统一的 review page shell，用来承接公共头部、离开动作、页面级状态与底部完成区；但两条页面不得复用成同一份无差别内容模板。

#### Scenario: 共享 shell 的允许复用边界

- **WHEN** 后续实现定义 review 页面宿主
- **THEN** 允许 daily / weekly 复用统一的 `ReviewPageShell` 或等价共享壳层
- **AND** 共享部分只允许包括：页面头部、说明文案承接位、页面级 loading/error 布局、底部完成区、主动返回 Dashboard 入口
- **AND** 不得把 daily / weekly 的主体区块收敛成同一份配置驱动的无语义模板，导致页面只剩标题不同

#### Scenario: 移动浏览器最小降级策略

- **WHEN** review 页面在移动浏览器窄屏环境展示
- **THEN** 必须继续沿用单一 `React Web` 页面主线
- **AND** 区块顺序必须保持与 PC 相同的主次关系
- **AND** 必须通过单列纵向编排、摘要裁剪、区块折叠阈值或紧凑按钮布局降低拥挤度
- **AND** 不得引入 review 专用第二套路由、第二套移动端页面或横向多列工作台

### Requirement: review 页面中的真实 inventory 必须按复用 / 扩展 / 禁止升级单值化

系统 SHALL 直接消费 `phase08-04` 冻结的真实 inventory，并对每个关键节点给出复用 / 扩展 / 禁止升级的单值结论。

#### Scenario: 真实 inventory 的消费矩阵

- **WHEN** 后续实现进入页面与交互落地
- **THEN** 必须得到以下单值结论：
  - `DashboardHomePage`：复用，继续承接总览页与 review launcher 宿主
  - `DashboardPrimaryActionPanel`：扩展，改造为稳定可见的 dual review launcher
  - `CurrentFocusSection`：复用其“只读主行动队列”语义给 `Daily Review` 主区，不升级为写路径 owner
  - `AssetFeedbackSection`：部分复用其“代表性缺口 + reuse snapshot”语义给 `Weekly Review`，但不得整块嵌入 review 页面
  - `RecentActivitySection`：复用其“紧凑活动流 + 局部失败”语义给 `Weekly Review`
  - `dashboard-source.ts`：复用，继续作为 canonical 页面回 Dashboard 的来源链事实源
  - `BackToDashboardButton`：复用，用于 review 导航进入的 canonical 页面回流
  - `FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar`：禁止升级为正式 review 入口或新的写路径 owner
  - `OnboardingCtaButton / SovereigntyPanel`：禁止升级为正式 review 入口
- **AND** 不得在缺少上述矩阵的情况下发明新的 `ReviewHubSection`、`ReviewStore` 或并列 dashboard-return 机制

### Requirement: review 完成与主动离开的交互流必须单值化

系统 SHALL 将 review 页面上的“完成”“进入下一步动作”“主动离开”设计为三类清晰且彼此不混写的交互出口。

#### Scenario: 主动离开 review 页面

- **WHEN** 用户在 `Daily Review` 或 `Weekly Review` 页面选择暂时离开
- **THEN** 必须直接返回 `/dashboard`
- **AND** 不得因为页面上下文不同而把主动离开分别导向 `Decision List`、浏览器返回历史或根路由 `/`

#### Scenario: review route 的 Dashboard 返回链合同

- **WHEN** 后续实现新增 `Daily Review` 与 `Weekly Review` route
- **THEN** review route 必须显式定义自身的 `validateSearch` 或等价 search schema
- **AND** 该 schema 必须正式承接 `fromDashboard / dashboardSection / dashboardReturnTo`
- **AND** review route 从 `/dashboard` 进入时，必须通过 URL search 参数携带 `buildDashboardSourceParams('empty-state')` 或等价单值来源参数，而不是只依赖内存态或组件局部变量
- **AND** review route 主动返回 `/dashboard` 时，允许继续复用既有 `navigate({ to: '/dashboard', state: { dashboardSection } })` 一次性 route state 机制恢复焦点
- **AND** review route 进入既有 canonical 页时，必须继续透传同一组 Dashboard 来源 search 参数，确保 `BackToDashboardButton` 仍然可用
- **AND** 不得为 review route 发明第二套 `reviewReturnTo / fromReviewDashboard` 命名，也不得把持久返回链改写为仅靠 `state`

#### Scenario: review 形成 decision handoff

- **WHEN** review 结果需要进入 `Decision`
- **THEN** 页面完成区必须把用户送入既有 `Decision List / Decision Detail / Decision Create` canonical 路径之一
- **AND** 若进入 `Decision Create`，该路径的 Dashboard 返回链扩展必须在后续子任务中单独定义后才可视为正式能力
- **AND** 当前 `phase08-05` 只冻结“进入 canonical Decision 页面”的页面流，不提前假定 `Decision Create` 已具备 Dashboard 返回链

#### Scenario: review 形成 entity handoff

- **WHEN** review 结果需要回流 `Product / Module / Repository`
- **THEN** 页面完成区必须把用户送入对应的既有 detail 或 list 页面
- **AND** 这些页面后续返回 `/dashboard` 时必须继续复用既有 `dashboard-source.ts + BackToDashboardButton` 链路
- **AND** review 页面本身不得承接实体写入表单或并列 mutation owner

#### Scenario: review 形成 next-step result

- **WHEN** review 结果被定义为 `next-step result`
- **THEN** 页面完成区必须通过后续 `review record` 路径承接该结果
- **AND** 当前阶段不得把它实现为仅停留在页面局部状态的“已处理”勾选
- **AND** `phase08-05` 只冻结其页面出口语义，不提前重写 `phase08-04` 已冻结的 record 边界

## MODIFIED Requirements

### Requirement: DashboardPrimaryActionPanel 的职责解释

`DashboardPrimaryActionPanel` 在 `phase08` 中 SHALL 不再被解释为 `phase05` 的单主 CTA 命中器，而是被解释为 `Dashboard` 标题行动区内的正式 review launcher 承接位。

#### Scenario: phase05 主 CTA 语义如何迁移

- **WHEN** 后续实现重写 `DashboardPrimaryActionPanel`
- **THEN** 允许移除其基于 `overview / feedback` 的单 CTA 命中状态机
- **AND** 必须改为稳定承接 `Daily Review / Weekly Review` 双入口
- **AND** 若仍需展示原有创建导向动作，只能降级为 review 页面或 canonical 页面内的次级动作，不得与正式 review 入口并列为同级主入口

## REMOVED Requirements

### Requirement: 通过复用 Dashboard 现有区块直接承接 review 工作会话

**Reason**: 这会让 `/dashboard` 重新退化为内联工作台，并把 `CurrentFocusSection / AssetFeedbackSection / RecentActivitySection` 从只读编排区块推成 review 会话宿主，直接破坏 `phase08-02` 与 `phase08-04` 已冻结的页面边界和 inventory 负向约束。
**Migration**: `Dashboard` 继续只承接总览与正式 review launcher；`Daily Review / Weekly Review` 作为独立 route 承接工作会话，只复用既有区块语义、列表样式与返回链，而不整块内嵌既有 Dashboard section。
