# phase10-07 产出后端合同、服务与读模型承接设计 Spec

## Why

`phase10-04 / 05 / 06` 已经把 `Onboarding` 建链流、关键页面 CTA、前端 owner / caller / 成功回流冻结下来了，但后端这层还没有明确写成“哪些必须复用既有 canonical contract/service，哪些才允许新增最小承接位”。如果这一步不收紧，后续实现很容易在 `Onboarding` 恢复、`Decision` pending reread、`Current Focus / next-step CTA` 上重新长出影子状态表、第二套 pending 字段，或者把页面 convenience 误写成新的业务合同。

## What Changes

- 冻结 `Onboarding` 建链引导所需的最小后端合同、QueryService 与可选恢复辅助承接位
- 冻结 `Decision` 生命周期推进与 pending reread 所需的最小合同、CommandService 与 reread 责任边界
- 冻结 `Current Focus / pending signals / next-step CTA` 相关读模型与 query owner 设计
- 冻结需要沿用与允许新增的 `.proto / Connect / service / store` 承接位
- 冻结“新增 vs 复用”的单值判定规则，禁止为页面 convenience 新增影子状态表或第二套 pending 字段

## Impact

- Affected specs:
  - `phase10_02_freeze_onboarding_first_run_chain_guidance_action_matrix`
  - `phase10_03_freeze_decision_lifecycle_detail_cta_pending_semantics_matrix`
  - `phase10_04_design_onboarding_chain_return_empty_failure_flow`
  - `phase10_05_design_dashboard_review_detail_next_step_cta_handoff`
  - `phase10_06_design_frontend_read_write_owner_route_caller_success_handoff`
  - `phase08_07_design_review_backend_contract_service_data_handoff`
  - `phase09_07_design_backend_service_contract_minimal_data_handoff`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
- Affected code:
  - `proto/psco/onboarding/v1/onboarding.proto`
  - `proto/psco/dashboard/v1/dashboard.proto`
  - `proto/psco/review/v1/review.proto`
  - `proto/psco/decision_center/v1/decision_center.proto`
  - `backend/internal/onboarding/service/query_service.go`
  - `backend/internal/onboarding/candidate/*`
  - 可选新增 `backend/internal/onboarding/repository/*`
  - `backend/internal/dashboard/service/query_service.go`
  - `backend/internal/dashboard/candidate/*`
  - `backend/internal/review/service/query_service.go`
  - `backend/internal/decisioncenter/service/query_service.go`
  - `backend/internal/decisioncenter/service/command_service.go`
  - `backend/internal/platform/router.go`
  - 对应 `connect/server.go` 与 generated Connect handlers

## ADDED Requirements

### Requirement: `phase10` 后端必须优先复用既有 canonical 写合同与业务服务

系统 SHALL 将 `phase10` 的后端写主线冻结为“优先复用既有 canonical contract/service”，不为 `Onboarding`、`Dashboard`、`Review` 或 detail page convenience 复制第二套写 RPC。

#### Scenario: `Onboarding` 的四类创建动作

- **WHEN** `Onboarding` 需要承接 `Product / Repository / Module / Decision` 的首轮创建
- **THEN** 后端正式合同必须继续复用既有 canonical RPC：
  - `ProductRegistryService.CreateProduct`
  - `RepositoryBindingService.CreateRepository`
  - `ModuleRegistryService.CreateModule`
  - `DecisionCenterService.CreateDecision`
- **AND** 不得在 `OnboardingService` 下新增 `CreateOnboardingProduct / CreateDraftRepository / CompleteOnboardingStep` 一类 phase10-local 写 RPC
- **AND** 不得新增 `/api/onboarding/*` 并列 JSON canonical API

#### Scenario: `Onboarding` 的关系闭合动作

- **WHEN** `Onboarding` 或各 detail 页需要承接正式关系闭合
- **THEN** 后端正式写入口必须继续复用既有 canonical RPC：
  - `ProductRegistryService.BindModuleToProduct`
  - `RepositoryBindingService.BindRepositoryToProduct`
  - `RepositoryBindingService.MapModuleToRepository`
  - `DecisionCenterService.LinkDecisionToTarget`
- **AND** 不得新增 `BindRepositoryFromOnboarding`、`MapModuleForPhase10`、`LinkDecisionFromCurrentFocus` 一类旁路写合同

#### Scenario: `Decision` 生命周期推进

- **WHEN** `Decision` 需要从 `proposed` 推进到 `active / superseded / archived`
- **THEN** 后端正式写入口必须继续是 `DecisionCenterService.UpdateDecisionStatus`
- **AND** `Dashboard / Review / Current Focus` 不得新增 `ResolvePendingDecision`、`AdvanceDecisionFromReview` 或等价第二套 command RPC

### Requirement: `Onboarding` 只允许新增单值恢复辅助读模型，不得新增第二套草稿真相源

系统 SHALL 冻结 `Onboarding` 在 `phase10` 的后端新增空间：只允许在既有 `psco.onboarding.v1` 包内新增单值恢复辅助读模型，用来表达 `current_product_id` 锚点、最近未完成 step 与未解释 handoff；不得新增第二套 draft state、第二套进度表或并列业务事实源。

#### Scenario: 现有 `GetFirstRunState` 的职责保留

- **WHEN** 后续评估现有 `OnboardingService.GetFirstRunState`
- **THEN** 它必须继续只承接冷启动摘要与首轮大盘状态：
  - `status`
  - `is_first_entry`
  - `current_step`
  - `completion_progress`
- **AND** 它可以继续作为 `Dashboard -> Onboarding` 的入口摘要读取
- **AND** 不得被直接扩写为承接完整 phase10 返回链、handoff 状态与 CTA inventory 的总读模型

#### Scenario: `phase10` 允许新增的 `Onboarding` 最小读合同

- **WHEN** 现有 `GetFirstRunState` 无法稳定表达 `phase10-04 / 06` 所需的恢复语义
- **THEN** 只允许在既有 `proto/psco/onboarding/v1/onboarding.proto` 中新增一个最小 read RPC，例如 `GetOnboardingChainState`
- **AND** 该 RPC 的最小响应边界只允许承接：
  - `current_product_id`
  - `current_step`
  - `resume_status`
  - `next_step_kind`
  - 可选的 `canonical_handoff_target`
  - 可选的 `return_hint`
- **AND** 不得内嵌四类创建表单 payload、完整草稿快照或页面局部 UI 状态
- **AND** 不得为了 phase10 再新建 `psco.onboarding_flow.v1` 或并列 proto 包

#### Scenario: `current_product_id` 的正式事实源层级

- **WHEN** `Onboarding` QueryService 需要返回 `current_product_id`
- **THEN** 正式事实源只允许按以下单值层级解释：
  - 已存在最小恢复锚点 store，且其中存在有效 `current_product_id`
  - 否则，按 canonical facts 派生唯一候选 `Product`
- **AND** 不得同时并列消费 store 与多个候选 `Product` 再由页面层决定谁是当前锚点
- **AND** 不得把来源 search、前端本地缓存或页面最近访问记录视为后端正式事实源

#### Scenario: 何时可仅靠 canonical facts 唯一恢复 `current_product_id`

- **WHEN** 后端试图不新增恢复锚点 store，直接用 canonical facts 恢复 `current_product_id`
- **THEN** 只有在当前 actor scope 内，`Onboarding` QueryService 按 `repository -> module -> decision -> complete` 的 phase10 主链顺序评估后，能够得到且仅得到一个候选 `Product` 时，才允许判定为“可唯一恢复”
- **AND** 该候选 `Product` 必须满足：
  - 已经作为首轮建链主上下文进入 phase10 主链
  - 其下游 canonical facts 能单值解释当前最近未完成 step
  - 不需要额外读取页面来源参数、本地草稿或人工兜底规则
- **AND** 若候选数为 `0` 或 `>1`
- **THEN** 必须判定为“canonical facts 不足以唯一恢复”
- **AND** 后续实现必须新增最小恢复锚点 store，而不是在 service 层继续发明第二套猜测规则

#### Scenario: `Onboarding` 恢复辅助读模型的 store 边界

- **WHEN** 仅靠现有 canonical facts 无法唯一恢复 `current_product_id`
- **THEN** 只允许新增一个最小恢复辅助承接位，例如 `onboarding_recovery_store`
- **AND** 该承接位只允许持久化单值恢复锚点：
  - `current_product_id`
  - 可选的 `updated_at`
- **AND** 不得持久化 `repository/module/decision` 草稿 payload
- **AND** 不得复制 counts、step 完成度、pending 状态或 detail CTA 结果
- **AND** 若现有 canonical facts 已足以唯一恢复，则不得新增该 store
- **AND** 新增该 store 后，它必须成为 `current_product_id` 的正式优先事实源，直到该锚点被新的首轮 `Product` 成功冻结所替代

#### Scenario: `Onboarding` QueryService 的正式承接位

- **WHEN** 后续实现 `phase10` 的 `Onboarding` 后端读取
- **THEN** 正式 owner 继续必须是 `backend/internal/onboarding/service/query_service.go`
- **AND** 允许在 `backend/internal/onboarding/candidate/` 下新增最小 reader 来读取：
  - `current_product_id` 锚点
  - 当前主上下文下的 `repository / module / decision` 完成度
  - 未解释 handoff 的 canonical 事实
- **AND** service 层不得直接跨模块写 SQL
- **AND** 不得新增并列 `OnboardingCommandService`

### Requirement: `Decision` pending reread 必须继续锚定 canonical `Decision.status`

系统 SHALL 冻结 `phase10` 下所有 pending reread 相关的后端语义继续完全锚定 canonical `Decision.status`，并且只通过既有 `DecisionCenter` / `Dashboard` / `Review` 服务承接，不新增第二套 pending 字段或 pending 总表。

#### Scenario: `Decision` 写合同的最小新增空间

- **WHEN** 评估 `Decision` 生命周期推进是否需要新合同
- **THEN** 必须优先判定“无需新增”
- **AND** 现有 `UpdateDecisionStatus / GetDecisionDetail / ListDecisions` 必须继续作为正式合同主线
- **AND** 不得新增 `GetPendingDecisionState`、`ResolveDecisionPending`、`GetDecisionNextAction` 一类并列 Decision RPC

#### Scenario: `pending` 的后端事实来源

- **WHEN** `Dashboard / Review / Current Focus` 需要读取 pending decision
- **THEN** 后端事实来源必须继续以 `Decision.status = proposed` 为唯一正式条件
- **AND** `decision_links` 只允许影响 target handoff 目标，不得影响 pending 是否退出
- **AND** `review_records`、页面已读状态或来源参数不得成为 pending 的并列事实来源

#### Scenario: 状态推进后的 reread 承接位

- **WHEN** `Decision` 状态推进成功后需要 reread
- **THEN** reread 必须继续分别回到既有服务：
  - `DecisionCenterService.GetDecisionDetail`
  - `DecisionCenterService.ListDecisions(status = proposed)`
  - `DashboardService.GetFeedbackSignals`
  - `ReviewService.GetDailyReviewContext`
- **AND** 不得新增专门的“reread after update”后端接口

### Requirement: `Current Focus / pending signals / next-step CTA` 必须优先演进既有 Dashboard/Review 读模型

系统 SHALL 将 `Current Focus / pending signals / next-step CTA` 的后端读模型优先收敛到既有 `DashboardService` 与 `ReviewService`，只有当现有合同无法稳定表达“下一步动作”时，才允许做最小字段级演进；不得新建 phase10-local 读服务。

#### Scenario: `Current Focus` 的 query owner 归属

- **WHEN** 后续实现 `Current Focus / pending signals / next-step CTA`
- **THEN** `Dashboard.QueryService` 必须继续是 `current_focus_signals` 的正式后端 owner
- **AND** `Review.QueryService` 必须继续只作为 `Daily Review / Weekly Review` 的组合 owner
- **AND** 不得新增 `phase10.QueryService`、`CurrentFocusService` 或并列模块承接相同事实

#### Scenario: 现有 `DashboardService.GetFeedbackSignals` 足以表达时

- **WHEN** 前端在当前页面上只需要：
  - `signal_family`
  - `signal_code`
  - `priority`
  - `target_type`
  - `target_id`
  - `action_label`
- **AND** page read owner 仍可仅靠这些字段单值推出：
  - `primaryCta`
  - `actionDescriptor`
  - canonical target 去向
- **AND** 不需要额外的：
  - `next_step_kind`
  - `canonical_owner`
  - `return_hint`
- **THEN** 必须直接复用现有 `DashboardService.GetFeedbackSignals`
- **AND** 不得为了“方便前端少判断”新增第二个 next-step RPC

#### Scenario: 现有反馈信号不足以稳定表达下一步动作时

- **WHEN** 既有 `FeedbackSignal` 出现以下任一情况：
  - 无法单值表达 `phase10-05 / 06` 已冻结的 `primaryCta / actionDescriptor`
  - 同一 `signal_code` 仍可能映射到两个以上 canonical owner
  - 页面无法仅靠现有字段稳定推出成功回流所需的 `return_hint`
- **THEN** 必须判定为“现有字段不足”
- **AND** 只允许在既有 `proto/psco/dashboard/v1/dashboard.proto` 中为 `FeedbackSignal` 或其最小包装 message 增加受控字段，例如：
  - `next_step_kind`
  - `canonical_owner`
  - `return_hint`
- **AND** 不得继续把这类场景判定为“无需新增”
- **AND** 不得新建 `GetNextStepCta`、`GetCurrentFocusAction` 一类并列 RPC
- **AND** 不得把前端 route/search 细节直接下沉为后端 URL 字段

#### Scenario: 判断 `FeedbackSignal` 的复用门槛是否满足

- **WHEN** 后续执行者评估某个 `Current Focus / pending signal` 场景应复用现有 `FeedbackSignal` 还是做最小字段演进
- **THEN** 必须逐项检查以下门槛：
  - 当前 signal 是否已能单值确定主 CTA
  - 当前 signal 是否已能单值确定 canonical owner
  - 当前 signal 是否已能单值确定 canonical target
  - 当前 signal 是否已能支撑前端生成稳定 `actionDescriptor`
  - 当前 signal 是否不需要额外 `return_hint`
- **AND** 只有以上各项全部满足时，才允许判定为“直接复用现有字段”
- **AND** 只要任一门槛不满足，就必须进入最小字段演进路径，只允许在既有 `proto/psco/dashboard/v1/dashboard.proto` 中为 `FeedbackSignal` 或其最小包装 message 增加受控字段，例如：
  - `next_step_kind`
  - `canonical_owner`
  - `return_hint`

#### Scenario: `ReviewService` 复用 Dashboard/Decision 读模型

- **WHEN** `Daily Review` 需要消费 `Current Focus / pending decisions / representative signals`
- **THEN** `Review.QueryService` 必须继续复用：
  - `dashboard.QueryService.ReadFeedbackSignal`
  - `decisioncenter.QueryService.ListDecisions(status = proposed)`
- **AND** 若 `Dashboard` 的 `FeedbackSignal` 增加了最小 `next_step` 描述，`review.proto` 必须直接复用或透传该 canonical message
- **AND** 不得在 `review.proto` 中复制第二套 `ReviewNextStepAction`

### Requirement: “新增 vs 复用”判定规则必须单值冻结

系统 SHALL 冻结 `phase10-07` 的新增判定规则，使后续执行者能够机械判断“这里该复用既有合同，还是允许新增最小承接位”。

#### Scenario: 可以由既有 canonical facts 组合得到

- **WHEN** 某个页面动作语义已经可以由以下既有 canonical facts 稳定组合得到：
  - `Decision.status`
  - `Dashboard FeedbackSignal`
  - `Product / Module / Repository / Decision Detail`
  - `ReuseSummary`
- **THEN** 必须判定为“复用既有合同/服务”
- **AND** 不得新增 `.proto`
- **AND** 不得新增并列表或影子字段
- **AND** 对 `Onboarding current_product_id` 与 `Current Focus next-step`，还必须分别满足前述“唯一恢复候选数 = 1”与“FeedbackSignal 复用门槛全部满足”的前置条件

#### Scenario: 页面层会重复编排，但现有合同仍无法稳定表达

- **WHEN** 页面层在多个入口上重复编排同一类“下一步动作”语义
- **AND** 现有合同无法稳定表达该语义
- **THEN** 只允许新增最小合同字段或最小辅助读模型
- **AND** 新增必须满足：
  - 落在既有 canonical proto 包内
  - 有明确单一 QueryService owner
  - 不复制既有 canonical facts
  - 不携带页面局部 UI 状态

#### Scenario: 只为了页面 convenience 想新增状态

- **WHEN** 新增提议只是为了让页面少做判断、少做 search 组装或少做 CTA 文案切换
- **THEN** 默认必须判定为“不允许新增”
- **AND** 这类问题应优先由既有 contract 字段复用、前端 page read owner 或 page action owner 解决

#### Scenario: 试图新增影子状态表或第二套 pending 字段

- **WHEN** 后续实现提出新增：
  - pending summary table
  - current focus snapshot table
  - onboarding draft progress table
  - 第二套 `is_pending` / `next_action_status`
- **THEN** 必须判定为未满足 `phase10-07`
- **AND** 除前述单值 `onboarding` 恢复辅助锚点外，不得新增影子状态持久化

### Requirement: Connect transport 与 router 挂载必须沿用既有模块主线

系统 SHALL 将 `phase10` 的后端 transport 继续冻结为既有模块内的 generated Connect handler + service implementation + `platform/router.go` 挂载，不新建 phase10-local transport 根。

#### Scenario: `Onboarding / Dashboard / Review / DecisionCenter` 的 transport 接线

- **WHEN** 后续为 `phase10-07` 演进后端接口
- **THEN** 必须继续分别落到既有模块：
  - `backend/internal/onboarding/connect/server.go`
  - `backend/internal/dashboard/connect/server.go`
  - `backend/internal/review/connect/server.go`
  - `backend/internal/decisioncenter/connect/server.go`
- **AND** `platform/router.go` 继续只挂既有 `OnboardingService / DashboardService / ReviewService / DecisionCenterService`
- **AND** 不得新增第二个 `/phase10`、`/workflow` 或并列 Connect 根

### Requirement: `phase10-07` 必须明确最小物理落点

系统 SHALL 为允许新增的后端承接位冻结最小物理落点，避免“理论上可以新增”但实现时落到任意模块。

#### Scenario: 允许新增的最小物理落点

- **WHEN** 后续开始实现 `phase10-07`
- **THEN** 只允许优先考虑以下最小落点：
  - `proto/psco/onboarding/v1/onboarding.proto`：新增 `GetOnboardingChainState` 或等价最小 read RPC
  - `backend/internal/onboarding/service/query_service.go`：新增链路恢复编排
  - `backend/internal/onboarding/candidate/chain_state_readers.go`：新增恢复 reader
  - 可选 `backend/internal/onboarding/repository/recovery_store.go`：仅当 canonical facts 无法唯一恢复 `current_product_id`
  - `proto/psco/dashboard/v1/dashboard.proto`：最小 `next_step` 字段演进
  - `backend/internal/dashboard/service/query_service.go`：产出 canonical `current_focus` 的最小 next-step 描述
  - `proto/psco/review/v1/review.proto`：仅做对既有 dashboard next-step message 的透传或复用
  - `backend/internal/review/service/query_service.go`：延续组合，不复制 signal 事实
- **AND** 不得新增并列 `backend/internal/phase10/`
- **AND** 不得在 `DecisionCenter` 之外新建第二个 decision lifecycle 服务

## MODIFIED Requirements

### Requirement: `OnboardingService` 在 `phase10` 中的解释

`phase10-07` 修改了现有 `OnboardingService` 的消费方式：它不再只服务 `first_run_state` 的冷启动摘要，还必须允许在同一 proto 包与同一模块内承接单值恢复辅助读模型；但它仍然不是新的写合同 owner。

#### Scenario: 判断 `OnboardingService` 是否越界

- **WHEN** 后续实现让 `OnboardingService` 同时承接恢复读取与业务写入
- **THEN** 必须判定为未满足 `phase10-07`
- **AND** `OnboardingService` 仍然只能承接读

### Requirement: `Dashboard / Review` 在 `phase10` 中的 next-step 解释

`phase10-07` 修改了 `Dashboard / Review` 对 next-step 语义的承接方式：默认继续复用既有 `FeedbackSignal` 与 Review context，只有在既有字段无法稳定表达已冻结的 CTA 语义时，才允许做最小字段级演进，而不是新建第二个读服务。

#### Scenario: 判断 next-step 设计是否越界

- **WHEN** 后续实现提出新建专门的 `NextStepService`、`CurrentFocusService` 或 phase10-local proto 包
- **THEN** 必须判定为未满足 `phase10-07`
- **AND** 应先回到 `dashboard.proto / review.proto` 的最小字段演进方案

## REMOVED Requirements

### Requirement: 为 `Onboarding` 新增并列写服务或第二套草稿状态表

**Reason**: 这会直接把 `Onboarding` 重新做成 workflow engine，并制造第二套与 canonical entities 并列的真相源，违背 `phase10` 的单值主链。
**Migration**: 四类创建与关系闭合继续复用既有 canonical contract/service；恢复语义若确实无法表达，只允许新增单值恢复辅助读模型或最小锚点 store。

### Requirement: 为 `Current Focus / pending signals` 新增第二套 pending 字段、影子表或专用读服务

**Reason**: `pending` 语义已经由 `Decision.status = proposed` 单值冻结，`Current Focus` 也已由 `Dashboard / Review` 服务承接；再新增专用状态层只会复制既有 canonical facts。
**Migration**: 继续复用 `DashboardService.GetFeedbackSignals`、`ReviewService.GetDailyReviewContext`、`DecisionCenterService.ListDecisions(status = proposed)`；如需更稳定的 next-step 描述，只允许对既有 message 做最小字段演进。
