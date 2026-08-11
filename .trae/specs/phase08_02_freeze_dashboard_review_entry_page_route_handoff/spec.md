# phase08-02 冻结 Dashboard review 入口、页面职责与路由承接位 Spec

## Why

`phase08-01` 已经冻结了 `Operating Review Loop` 的范围边界、成功标准与非目标，但要让后续 `/spec` 与实现真正可落地，还必须先把 `Dashboard`、daily / weekly review 页面、正式入口 caller 与路由回流关系冻结成单值结论。  
否则后续实现很容易回到“在 `/dashboard` 里多塞两个按钮”或“多个页面各自长出 review 入口”的模糊状态，最终既无法形成独立可验收的 daily / weekly 会话，也会破坏既有 Dashboard / Decision 主线的职责边界。

## What Changes

- 冻结 `Dashboard` 与 daily / weekly review 的页面边界和正式路由关系
- 冻结 daily / weekly review 的最小入口形态、页面职责、输入优先级与完成定义
- 冻结 `Dashboard` 标题行动区为当前阶段唯一正式 review 入口承接位
- 冻结既有 `FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar / OnboardingCtaButton / SovereigntyPanel` 不升级为正式 review 入口
- 冻结 review 页与 `/dashboard` 之间的单值回流路径
- 冻结当前阶段不采用 Dashboard 内联工作台、弹层主线或多页面重复 review 入口

## Impact

- Affected specs:
  - `phase08_01_freeze_operating_review_loop_scope_success_non_goals`
  - `phase08_operating_review_loop_foundation_dev_plan.md`
  - `phase08_operating_review_loop_foundation_shared_baseline.md`
  - 后续 `phase08-03 / phase08-04 / phase08-05 / phase08-06`
- Affected code:
  - `frontend/src/routes/dashboard.tsx`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
  - 预期新增的 review route 宿主与 review page 壳层
  - `frontend/src/routes/decisions/new.tsx`
  - `frontend/src/features/dashboard/lib/dashboard-source.ts`

## ADDED Requirements

### Requirement: `Dashboard` 与 review 页面主线必须单值分层

系统 SHALL 将 `Dashboard Home` 冻结为 review 入口页与既有聚合读取页，而不是直接承接 review 工作会话的页面宿主；daily / weekly review 必须作为独立页面主线承接。

#### Scenario: 判定 `Dashboard Home` 的页面职责

- **WHEN** 用户进入 `/dashboard`
- **THEN** 页面必须继续承接既有 `overview / current focus / asset feedback / recent activity / sovereignty / onboarding CTA` 的只读编排
- **AND** 必须新增明确的 daily / weekly review 进入入口
- **AND** 不得在 `DashboardHomePage` 内直接承接 review 工作会话、review 结果提交或新的 mutation owner
- **AND** 不得把 `/dashboard` 改写成带 review session 搜索参数的工作台页

#### Scenario: 判定 review 页面主线

- **WHEN** 后续 `/spec` 或实现定义 review 页面
- **THEN** 必须得到 `Daily Review` 与 `Weekly Review` 是两条独立页面级工作会话的单值结论
- **AND** 不得把 review 主线收敛成 Dashboard 内联展开区、抽屉主线、弹层主线或第二套 Dashboard 子页面体系

### Requirement: daily / weekly review 正式路由必须显式独立

系统 SHALL 将 daily / weekly review 的正式业务入口冻结为两条显式独立路由，而不是通过 `/dashboard` 搜索参数或同页切换冒充双路径。

#### Scenario: 判定 review 路由承接位

- **WHEN** 接手者判断 `phase08` 的 review 正式路由
- **THEN** 必须得到以下单值结论：
  - `Daily Review` 通过独立路由承接
  - `Weekly Review` 通过独立路由承接
- **AND** 推荐落点为 `/reviews/daily` 与 `/reviews/weekly` 或与之等价的两条显式独立 route 宿主
- **AND** 不得使用 `/dashboard?review=daily|weekly`、`/dashboard#daily` 或“同一路由 + 两按钮 + 条件渲染”作为正式主线
- **AND** daily / weekly 可以复用同一页面壳层组件，但不得共用同一条 route 身份

### Requirement: `Dashboard` 标题行动区必须成为唯一正式 review 入口承接位

系统 SHALL 将 `Dashboard` 标题行动区冻结为当前阶段唯一正式 review 入口承接位，用于从 `/dashboard` 进入 daily / weekly review。当前源码中的 `DashboardPrimaryActionPanel` 是该承接位的既有真实载体，但 `phase08` 不强制要求后续实现继续保留它当前的单 CTA 状态机形态。

#### Scenario: 判定哪个既有 caller 被升级为 review 正式入口承接位

- **WHEN** 后续 `/spec` 说明 caller 升级关系
- **THEN** 必须将 `DashboardPrimaryActionPanel` 所在的 `Dashboard` 标题行动区升级为 daily / weekly review 的唯一正式入口承接位
- **AND** 该承接位必须单值承接两类动作：进入 `Daily Review`、进入 `Weekly Review`
- **AND** 若后续实现继续复用 `DashboardPrimaryActionPanel` 组件本身，必须先把它从“条件单 CTA 面板”重定义为稳定可见的 review launcher
- **AND** 若后续实现不继续复用该组件本身，也必须保持唯一正式入口仍位于同一标题行动区，而不是迁移到其他 Dashboard 区块形成并列主入口
- **AND** 不得再在其他 Dashboard 区块重复长出第二套正式 review 入口

### Requirement: 既有 Dashboard caller 的职责必须保持单值

系统 SHALL 冻结 `FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar / OnboardingCtaButton / SovereigntyPanel` 在 `phase08-02` 中的职责，不将它们升级为正式 review 入口。

#### Scenario: 判定既有 caller 是否继续保持只读跳转

- **WHEN** 后续 `/spec` 判断既有 Dashboard caller 的职责
- **THEN** 必须保持以下单值结论：
  - `FeedbackSignalCard` 继续承接到既有 canonical 页面或 `Decision` 的事实跳转
  - `RecentActivityItemCard` 继续承接到既有详情页或列表页的事实跳转
  - `DashboardStatBar` 继续承接概览到既有 list 页的跳转
  - `OnboardingCtaButton` 继续承接 onboarding 入口
  - `SovereigntyPanel` 继续承接数据主权相关入口
- **AND** 不得在这些 caller 上并行增加 daily / weekly review 正式入口按钮
- **AND** 不得把 `FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar / OnboardingCtaButton / SovereigntyPanel` 与标题行动区并列解释为第二套正式 review 入口

### Requirement: daily / weekly review 页面职责必须分别单值化

系统 SHALL 将 daily / weekly review 的页面职责、输入优先级、最小页面区块与完成定义分别冻结，避免退化成“同页双按钮”。

#### Scenario: Daily Review 页面职责判定

- **WHEN** 用户进入 `Daily Review`
- **THEN** 页面必须优先承接 `current_focus_signals`
- **AND** 必须承接 `pending decisions`
- **AND** 必须在需要时补充 `representative_signals`
- **AND** 页面完成语义必须偏向“形成当下一个动作并立即进入处理”
- **AND** 当前阶段不得把 `Daily Review` 扩写成周期性整理工作台

#### Scenario: Weekly Review 页面职责判定

- **WHEN** 用户进入 `Weekly Review`
- **THEN** 页面必须优先承接 `overview / recent activities / representative signals / reuse snapshot`
- **AND** 必须体现周期性整理语义，而不是 `Daily Review` 的简单放大版
- **AND** 页面完成语义必须偏向“形成一组后续动作或决策入口”
- **AND** 当前阶段不得把 `Weekly Review` 收窄为只展示最近活动的只读页

#### Scenario: 判定 daily / weekly 的最小页面级区块

- **WHEN** 后续 `/spec` 定义 review 页的最小页面结构
- **THEN** `Daily Review` 至少必须包含：页面头部、current focus 区、pending decision 区、代表性缺口补充区、离开或进入下一步动作区
- **AND** `Weekly Review` 至少必须包含：页面头部、overview / recent activity 区、代表性反馈区、reuse snapshot 区、离开或进入下一步动作区
- **AND** 两者允许复用页面壳层，但不得复用完全相同的区块归属与标题语义

### Requirement: review 与 Dashboard 的回流路径必须单值化

系统 SHALL 将 review 页面回到 `Dashboard` 的路径冻结为单值回流，而不是在多个页面分别定义不同返回语义。

#### Scenario: 判定 review 返回 Dashboard 的路径

- **WHEN** 用户从 `Daily Review` 或 `Weekly Review` 主动离开 review 页面
- **THEN** 必须回到 `/dashboard`
- **AND** 必须保持 `Dashboard` 作为统一经营动作入口页
- **AND** 不得为 daily / weekly 分别长出不同的首页落点

#### Scenario: 判定 Dashboard 来源上下文工具的继续使用边界

- **WHEN** 后续 `/spec` 设计从 review 进入既有 canonical 页面后的返回链
- **THEN** 只允许在已接入 `dashboardSourceSearchSchema` 或已明确消费 `BackToDashboardButton` 的 canonical route 上继续复用既有 `dashboard-source.ts` 机制
- **AND** 当前直接可复用的最小范围至少包括：`/decisions`、`/decisions/$decisionId`、`/modules/$moduleId`、`/products/$productId`、`/repositories/$repositoryId` 及其已承接的既有列表/详情返回链
- **AND** 不得把 `/decisions/new` 默认视为已具备同等返回链能力
- **AND** 若后续 review 需要进入 `Decision Create` 并保留 Dashboard 返回链，必须在后续子任务中单独扩展 `DecisionCreateRoute` 的 `validateSearch` 与页面消费边界后，才允许复用该机制
- **AND** 不得为了 review 入口在 `/dashboard` 上新增第二套持久搜索参数事实源
- **AND** 若 review 需要额外来源上下文，只能在 review route 自身或既有来源机制上最小扩展

## MODIFIED Requirements

### Requirement: `Dashboard Home` 页面职责解释

`Dashboard Home` 在 `phase08` 中 SHALL 不再只被解释为“总览页”，而是被解释为“总览 + 正式 review 入口页”；但它仍然不是 review 工作会话的正式页面宿主。

#### Scenario: 判定 Dashboard 的新语义

- **WHEN** 后续 `/spec`、实现或验收描述 `/dashboard`
- **THEN** 必须明确 `Dashboard` 已从纯总览页升级为经营动作入口页
- **AND** 必须明确它通过正式入口把用户送入 `Daily Review / Weekly Review`
- **AND** 不得把 `Dashboard` 重新解释成 review 工作台本身

## REMOVED Requirements

### Requirement: 在多个页面重复长出 review 正式入口

**Reason**: 这会让 `phase08` 失去单值入口，导致 `Dashboard`、反馈卡片、活动卡片、详情页工具栏分别生长 review 入口，最终回到理想化而不可验收的多入口状态。  
**Migration**: 当前阶段统一由 `Dashboard` 标题行动区承接 daily / weekly review 正式入口；当前源码中的 `DashboardPrimaryActionPanel` 可以作为该承接位的演进起点，但不得把它当前的单 CTA 语义直接等同于最终正式入口。其他既有 caller 继续保留当前 canonical 跳转职责，后续如确需新增 review 次级入口，必须在新的 `phase08` 子任务中单独冻结并说明不与正式入口并列。
