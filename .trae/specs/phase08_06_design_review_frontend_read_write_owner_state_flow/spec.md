# phase08-06 产出前端读写承接位与状态流设计 Spec

## Why

`phase08-03` 已冻结 review 动作必须收敛到单一 `Review action application owner`，`phase08-04` 已冻结 review read model 只是既有 canonical facts 的轻量组合层，`phase08-05` 已冻结 daily / weekly review 的页面流与交互流。当前还缺一份可直接指导实现的前端承接位设计，来回答 review 相关读取落到哪里、写动作由谁统一编排、成功后回哪里、哪些临时页面编排必须回收。
如果这一步不收紧，后续实现很容易回到“review 页面直接拼 `useFeedbackSignalsRead + useDecisionListRead + useMutation`”的模式，最终让 `query` 与 `application` 边界重新散掉，并把 review 编排漂移成第二套 page-local owner。

## What Changes

- 冻结 review 前端正式切片落点为 `frontend/src/features/review/`
- 冻结 daily / weekly review read layer 的唯一正式承接位与 `phase05 / phase06` 数据消费位置
- 冻结 `Review action application owner` 的接口、职责、成功回流与错误归一化边界
- 冻结 review 页面级状态流、区块级状态流与 query 失效策略
- 冻结必须回收的页面级 / 组件级临时编排点，避免把既有 canonical 页面模式复制进 review
- 冻结 caller 与 owner 的一对一映射表，保证 review 编排不散落在 route / 页面 / 展示组件

## Impact

- Affected specs:
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `phase08_04_freeze_review_contract_read_record_boundary`
  - `phase08_05_design_dashboard_review_entry_page_interaction_flow`
  - `phase06_07_design_frontend_write_path_mutation_owners`
  - `phase07_05_design_frontend_generated_client_query_application_migration`
- Affected code:
  - 后续新增 `frontend/src/features/review/data/*`
  - 后续新增 `frontend/src/features/review/application/*`
  - 后续新增 `frontend/src/features/review/pages/*`
  - 后续新增 `frontend/src/features/review/components/*`
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/dashboard/data/use-dashboard-overview-read.ts`
  - `frontend/src/features/dashboard/data/use-feedback-signals-read.ts`
  - `frontend/src/features/dashboard/data/use-recent-activities-read.ts`
  - `frontend/src/features/reuse-summary/data/use-reuse-summary-read.ts`
  - `frontend/src/features/decision-center/data/use-decision-list-read.ts`
  - `frontend/src/features/decision-center/application/use-create-draft-decision.ts`
  - `frontend/src/features/decision-center/application/use-link-decision-to-target.ts`
  - `frontend/src/features/product-registry/application/use-bind-module-to-product.ts`
  - `frontend/src/features/repository-binding/application/use-bind-repository-to-product.ts`
  - `frontend/src/features/repository-binding/application/use-map-module-to-repository.ts`

## ADDED Requirements

### Requirement: review 前端正式承接位必须收敛为单一 `review` 切片

系统 SHALL 将 review 前端正式承接位冻结为单一业务切片 `frontend/src/features/review/`，而不是把 review 读取与写入继续散落在 `dashboard / decision-center / product-registry / repository-binding` 的页面或组件内。

#### Scenario: review 切片的最小物理落点

- **WHEN** 后续实现开始新增 review 前端能力
- **THEN** 必须新增以下最小落点：
  - `frontend/src/features/review/data/use-daily-review-read.ts`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/review/data/review-query-options.ts` 或等价只读 helper
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/features/review/application/review-action-types.ts`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/review/components/review-page-shell.tsx`
- **AND** `review/data/` 只承接读取、queryKey、响应解包与只读派生
- **AND** `review/application/` 只承接写动作编排、成功回流语义、query 失效与错误归一化
- **AND** 不得为 review 另起根级 `services/`、`stores/`、`gateways/` 并列体系

#### Scenario: 既有切片与 review 切片的关系

- **WHEN** review 需要消费 `Dashboard / Reuse Summary / Decision` 的 canonical facts
- **THEN** `dashboard / reuse-summary / decision-center` 继续保留各自 canonical `data/` 与 `application/` owner
- **AND** `review` 切片只允许消费这些既有 owner，不得把它们复制成第二套 feature-local fetch / mutation
- **AND** review 不得回写既有切片的 `data/` 文件来承接 review-local 状态机

### Requirement: daily / weekly review read layer 必须有唯一正式消费位置

系统 SHALL 把 daily / weekly review 的数据消费位置冻结在 `review/data/` 的两个正式 read owner 中，而不是让 route、page 或 section 组件各自直接消费底层 query hooks。

#### Scenario: Daily Review 的正式 read owner

- **WHEN** `Daily Review` 页面需要读取 review context
- **THEN** 必须只通过 `useDailyReviewRead()` 或等价单一 read owner 承接
- **AND** 该 read owner 必须正式消费：
  - `useFeedbackSignalsRead()` 的 `current_focus_signals`
  - `useFeedbackSignalsRead()` 的 `asset_feedback_summary.representative_signals`
  - `useDecisionListRead({ statusFilter: 'proposed' })` 作为当前阶段 `pending / backlog decisions` 的最小前端事实来源
- **AND** `pending decisions` 在当前阶段的最小展示口径必须是 `proposed` 决策的 top N 摘要列表
- **AND** 若后续要把 `active` 扩展进 backlog 口径，必须在后续 `/spec` 中单独冻结，不得在实现阶段自行放宽

#### Scenario: Weekly Review 的正式 read owner

- **WHEN** `Weekly Review` 页面需要读取 review context
- **THEN** 必须只通过 `useWeeklyReviewRead()` 或等价单一 read owner 承接
- **AND** 该 read owner 必须正式消费：
  - `useDashboardOverviewRead()`
  - `useRecentActivitiesRead()`
  - `useFeedbackSignalsRead()` 的 `asset_feedback_summary.representative_signals`
  - `useReuseSummaryRead({ scope: 'dashboard' })` 返回的 `module_reuse_summary / capability_summary`
- **AND** `reuse snapshot / module_reuse_summary / capability_summary` 的正式消费位置只允许在 `useWeeklyReviewRead()` 内完成组合
- **AND** review 页面与 section 组件不得各自再次直接调用 `useReuseSummaryRead()`

#### Scenario: review 读取不得在页面层重新拼装第二套 query 主线

- **WHEN** `DailyReviewPage` 或 `WeeklyReviewPage` 实现页面编排
- **THEN** 页面层只允许消费 `useDailyReviewRead()` / `useWeeklyReviewRead()` 的聚合结果
- **AND** 不得在页面层再次直接 import `useFeedbackSignalsRead / useDashboardOverviewRead / useRecentActivitiesRead / useReuseSummaryRead / useDecisionListRead`
- **AND** `CurrentFocusSection / AssetFeedbackSection / RecentActivitySection` 若被复用，也只能消费 review slice 传入的只读 props，不得在组件内部自行读取

#### Scenario: review read layer 的缓存口径仍然是轻量组合层

- **WHEN** `useDailyReviewRead()` 或 `useWeeklyReviewRead()` 在前端形成聚合读取
- **THEN** 其缓存语义必须继续建立在底层 canonical query key 之上
- **AND** 当前阶段不得把 `['review']` 冻结成必需的长期 query key
- **AND** 若后续实现确实引入 review slice-local query key，它的身份也只能是组合层缓存别名，而不是新的事实主缓存
- **AND** `Review action application owner` 的正式失效主线必须以底层 canonical keys 为准，而不是依赖一个未经上游冻结的 `['review']`

### Requirement: review read layer 必须导出稳定的页面级状态模型

系统 SHALL 把 review 页面所需的状态流冻结为“read owner 派生页面级状态 + 区块级状态”，而不是把多个 query 的原始状态直接暴露给页面逐段拼接。

#### Scenario: page-level 状态模型

- **WHEN** `useDailyReviewRead()` 或 `useWeeklyReviewRead()` 对外暴露页面状态
- **THEN** 必须至少导出以下页面级状态之一：
  - `initial-loading`
  - `ready`
  - `page-error`
- **AND** page-level 状态由 read owner 内部根据必需 query 的组合结果派生
- **AND** route / page 不得自己基于多个 `isLoading / isError` 再拼一遍第二套状态机

#### Scenario: section-level 状态模型

- **WHEN** review 页面存在局部失败或空状态区块
- **THEN** read owner 还必须导出区块级状态：
  - `ready`
  - `empty`
  - `error`
- **AND** `Weekly Review` 的 `reuse snapshot` 区块失败只允许是区块级 `error`
- **AND** `Daily Review` 的 `pending decisions` 空状态只允许是区块级 `empty`
- **AND** 不得因为单一区块失败把整页一律降级成 `page-error`

### Requirement: review 写路径必须收敛到单一 `Review action application owner`

系统 SHALL 将 review 中所有正式动作统一收敛到 `review/application/use-review-action.ts` 或等价单一 owner，由它消费既有 canonical application owner 并向页面暴露稳定的 success handoff 语义。

#### Scenario: Review action owner 的最小接口

- **WHEN** 后续实现 `Review action application owner`
- **THEN** 它至少必须暴露：
  - `submitAction(input): Promise<ReviewActionSuccess>`
  - `isPending`
  - `error`
  - `reset()`
- **AND** `ReviewActionSuccess` 必须是稳定的 review-facing 成功 envelope，而不是直接把底层 mutation response 裸传给页面
- **AND** 该成功 envelope 至少必须包含：
  - `resultKind`
  - `navigateTo`
  - `params`
  - `search`
  - 可选的 `successMessage`
- **AND** 页面只负责消费该 success envelope 执行 `navigate()` 与展示 toast
- **AND** 只要成功目标是既有 canonical 页面，`search` 就必须继续承接 `fromDashboard / dashboardSection / dashboardReturnTo`

#### Scenario: Review action owner 的正式职责

- **WHEN** 页面触发 review 完成区动作
- **THEN** `Review action application owner` 只允许承接：
  - 选择进入 `Decision` draft
  - 选择进入既有 `Decision Detail / Decision List`
  - 选择进入既有 `Product / Module / Repository` canonical 页面
  - 选择提交 `next-step result` 到后续 `review record` 路径
  - 统一错误归一化
  - 统一 query 失效
- **AND** route / page / section / card 不得直接持有第二套 `useMutation`
- **AND** `Review action application owner` 不得直接持有页面局部 UI 状态，如抽屉开闭、区块展开或按钮 hover

#### Scenario: Review action owner 复用既有 canonical owner

- **WHEN** review 动作需要真正执行写入
- **THEN** 必须优先复用既有 canonical owner：
  - `useCreateDraftDecision`
  - `useLinkDecisionToTarget`
  - `useBindModuleToProduct`
  - `useBindRepositoryToProduct`
  - `useMapModuleToRepository`
- **AND** 对于纯 route handoff 动作，`Review action application owner` 可以直接返回成功 envelope，而无需发起 mutation
- **AND** 不得在 review 切片内复制这些 mutation 的 transport、query 失效或错误处理实现

### Requirement: review 成功回流与 query 失效策略必须由 owner 单值化

系统 SHALL 把 review 写动作的成功回流与 query 失效策略冻结在 `Review action application owner` 中，并以“canonical owner 负责自身切片失效 + review owner 负责 review 与 dashboard 相关失效”的方式组合。

#### Scenario: Decision draft created from review

- **WHEN** review 动作创建新的 `Decision` draft
- **THEN** 必须复用 `useCreateDraftDecision`
- **AND** 由该 canonical owner 继续失效 `decision-list`
- **AND** `Review action application owner` 还必须补充失效：
  - `['dashboard-overview']`
  - `['dashboard-feedback-signals']`
  - `['dashboard-recent-activities']`
- **AND** 页面成功后必须导航到既有 `Decision Detail` 或既有 `Decision` canonical 路径
- **AND** 成功 envelope 的 `search` 必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`

#### Scenario: Decision link completed from review

- **WHEN** review 动作执行 `Decision -> Module` 关联
- **THEN** 必须复用 `useLinkDecisionToTarget`
- **AND** 由该 canonical owner 继续失效 `decision-detail` 与 `decision-module-candidates`
- **AND** `Review action application owner` 还必须补充失效：
  - `['dashboard-feedback-signals']`
  - `['dashboard-recent-activities']`
- **AND** 页面成功后必须导航到既有 `Decision Detail` 或目标实体 canonical 页面，而不是停留在 review 页面上自行渲染“已处理”
- **AND** 成功 envelope 的 `search` 必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`

#### Scenario: Module 绑定到 Product 由 review 触发

- **WHEN** review 动作执行 `BindModuleToProduct`
- **THEN** 必须复用 `useBindModuleToProduct`
- **AND** 由该 canonical owner 继续失效 `product-detail` 与 `product-module-candidates`
- **AND** `Review action application owner` 还必须补充失效：
  - `['product-list']`
  - `['dashboard-overview']`
  - `['dashboard-feedback-signals']`
  - `['dashboard-recent-activities']`
  - `['reuse-summary', 'dashboard']`
- **AND** 页面成功后必须导航到既有 `Product Detail` 或既有 `Product` canonical 路径
- **AND** 成功 envelope 的 `search` 必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`

#### Scenario: Repository 绑定到 Product 由 review 触发

- **WHEN** review 动作执行 `BindRepositoryToProduct`
- **THEN** 必须复用 `useBindRepositoryToProduct`
- **AND** 由该 canonical owner 继续失效 `repository-detail` 与 `repository-product-candidates`
- **AND** `Review action application owner` 还必须补充失效：
  - `['dashboard-overview']`
  - `['dashboard-feedback-signals']`
  - `['dashboard-recent-activities']`
- **AND** 页面成功后必须导航到既有 `Repository Detail`、`Product Detail` 或既有对应 canonical 路径
- **AND** 成功 envelope 的 `search` 必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`

#### Scenario: Module 映射到 Repository 由 review 触发

- **WHEN** review 动作执行 `MapModuleToRepository`
- **THEN** 必须复用 `useMapModuleToRepository`
- **AND** 由该 canonical owner 继续失效 `repository-detail` 与 `repository-module-candidates`
- **AND** `Review action application owner` 还必须补充失效：
  - `['dashboard-overview']`
  - `['dashboard-feedback-signals']`
  - `['dashboard-recent-activities']`
- **AND** 页面成功后必须导航到既有 `Repository Detail`、`Module Detail` 或既有对应 canonical 路径
- **AND** 成功 envelope 的 `search` 必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`

#### Scenario: Entity handoff without immediate mutation

- **WHEN** review 动作当前只要求进入既有 `Product / Module / Repository` canonical 页面
- **THEN** `Review action application owner` 只返回 route handoff success envelope
- **AND** 不得为了“统一接口”强行制造空 mutation
- **AND** 这类 handoff 不要求立即失效 query，但必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`

#### Scenario: Query invalidation follows TanStack Query v5 canonical pattern

- **WHEN** review owner 需要在 mutation 成功后刷新相关读模型
- **THEN** 必须通过 `queryClient.invalidateQueries()` 承接，而不是在页面层手写多处 queryClient 操作
- **AND** 若一次动作涉及多组 query，owner 必须按官方推荐在 `onSuccess` 中统一触发并等待相关 invalidation 完成
- **AND** 正式失效目标必须优先指向底层 canonical query key，而不是依赖未冻结的 review-local 主缓存 key
- **AND** 当前阶段不引入 review 专用 optimistic update

### Requirement: review 页面不得复制既有 canonical 页面中的临时编排模式

系统 SHALL 明确当前仓库中哪些页面级 / 组件级临时编排点只能留在既有 canonical 页面，review 实现不得继续复制这些模式。

#### Scenario: 必须避免复制的临时编排点

- **WHEN** 后续实现 review 页面与动作区
- **THEN** 不得复制以下模式进入 review：
  - `ProductDetailPage / RepositoryBindingDetailPage` 中的页面级 `invalidateDetail()` 回调拼装
  - `DecisionModuleCandidatePanel` 中组件级包装 `mutation.mutate(..., { onSuccess, onError })`
  - `DashboardHomePage` 中页面级直写多个 retry invalidation 回调作为 review 主线
  - `SovereigntyPanel` 中组件内直接声明 mutation 的过渡写法
- **AND** review 页面只允许消费 `Review action application owner` 暴露的稳定动作接口
- **AND** `DashboardPrimaryActionPanel` 在进入 review 后只保留 route launcher 身份，不升级为 review state coordinator

### Requirement: caller 与 owner 的一对一映射必须可执行

系统 SHALL 产出 review 前端 caller 与 owner 的一对一映射表，保证每个 caller 都能追溯到唯一正式 owner。

#### Scenario: review caller-owner 映射表

- **WHEN** 后续实现 review 前端
- **THEN** caller 与 owner 至少必须满足以下单值映射：
  - `DashboardPrimaryActionPanel` -> `DailyReviewRoute / WeeklyReviewRoute`（只负责 route launcher）
  - `DailyReviewPage` -> `useDailyReviewRead`
  - `WeeklyReviewPage` -> `useWeeklyReviewRead`
  - `ReviewPageShell` -> 只消费 read owner 派生状态与 action owner 结果，不直连底层 hooks
  - `CurrentFocusSection`（若复用） -> `useDailyReviewRead` 传入 props
  - `PendingDecisionSection` -> `useDailyReviewRead` 传入 props
  - `RepresentativeSignalsSection` -> `useDailyReviewRead / useWeeklyReviewRead` 传入 props
  - `ReuseSnapshotSection` -> `useWeeklyReviewRead` 传入 props
  - `ReviewActionFooter` 或等价完成区 -> `useReviewAction`
  - `BackToDashboardButton` -> 继续只消费 `dashboard-source.ts`
- **AND** 不得出现 “同一个 caller 可选两个 owner” 或 “同一个 owner 需要页面再补第二段 query / mutation” 的并列状态

## MODIFIED Requirements

### Requirement: `DashboardPrimaryActionPanel` 的前端职责解释

自 `phase08-06` 起，`DashboardPrimaryActionPanel` SHALL 被解释为 review route launcher caller，而不是 review 读层或写层的编排 owner。

#### Scenario: Dashboard 标题行动区不持有 review 状态

- **WHEN** `DashboardPrimaryActionPanel` 增加 daily / weekly review 双入口
- **THEN** 它只能负责组装 route search 并导航到对应 review route
- **AND** 不得在该组件中直接读取 `pending decisions`、`reuse snapshot` 或持有 `Review action application owner`
- **AND** review session 的正式状态必须在 review slice 内建立

### Requirement: canonical application owner 的消费方式

`phase08-03` 已冻结既有 canonical application owner 必须被复用。
自 `phase08-06` 起，系统必须进一步要求这些 canonical owner 在 review 场景下只能被 `Review action application owner` 消费，而不是被 review page / review section 直接调用。

#### Scenario: review 不直接消费 canonical mutation hooks

- **WHEN** `DailyReviewPage` 或 `WeeklyReviewPage` 需要触发决策或实体回流
- **THEN** 页面不得直接 import `useCreateDraftDecision / useLinkDecisionToTarget / useBindModuleToProduct / useBindRepositoryToProduct / useMapModuleToRepository`
- **AND** 这些 hook 只允许由 `useReviewAction()` 在内部调用或编排

## REMOVED Requirements

### Requirement: review 页面可直接组合既有 query hooks 与 canonical mutation hooks 作为正式实现

**Reason**: 这会让 review 编排重新散落在 route / page / section / card 中，直接违背本项目已冻结的 `query` 纯只读、`application` 唯一写路径，以及 `caller -> owner` 单值映射约束。
**Migration**: review 页面一律退回为 caller；读取统一收敛到 `review/data/use-daily-review-read.ts` 与 `use-weekly-review-read.ts`，写动作统一收敛到 `review/application/use-review-action.ts`，成功回流由 success envelope + 页面导航消费完成。
