# phase08-10 落实 `Feedback -> Decision -> Update` 闭环、结果回流与临时编排清理 Spec

## Why

`phase08-08` 已把 review 合同、后端承接位与前端 owner 主线落进仓库，`phase08-09` 又把 Dashboard 双入口、daily / weekly 会话与统一动作承接推进成真实页面行为。但当前还缺一份把“反馈如何进入正式决策、决策如何回流既有实体更新、哪些只是过渡 handoff、哪些必须被清理”冻结成单值结论的实现规格；否则 review 仍可能停留在导航壳层，或把成功回流、错误语义与刷新策略重新散落到页面局部逻辑中。

## What Changes

- 冻结 `Feedback -> Decision -> Update` 的最小可重复执行闭环
- 冻结 `Decision` 作为 review loop 正式中心时的进入方式、成功回流与结果回写矩阵
- 冻结至少一种实体更新路径必须从 review loop 真实落到既有 canonical update，并完成 reread
- 冻结错误语义、成功回流与必要刷新必须继续收敛在既有 owner 主线
- 冻结 `phase08` 过程中新增的临时散装编排点的回收范围，防止前后端保留并列临时主线

## Impact

- Affected specs:
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `phase08_06_design_review_frontend_read_write_owner_state_flow`
  - `phase08_07_design_review_backend_contract_service_data_handoff`
  - `phase08_08_land_review_contract_backend_frontend_owner_enablement`
  - `phase08_09_land_dashboard_review_entry_dual_session_unified_action_handoff`
  - `phase03_10_decision_center_formal_spec`
  - `phase04_10_product_repository_binding_formal_spec`
  - `phase04_14_product_repository_binding_integration_validation_acceptance`
- Affected code:
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/features/review/components/review-action-footer.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-create-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/decision-center/application/use-create-draft-decision.ts`
  - `frontend/src/features/decision-center/application/use-link-decision-to-target.ts`
  - `frontend/src/features/product-registry/application/use-bind-module-to-product.ts`
  - `frontend/src/features/repository-binding/application/use-bind-repository-to-product.ts`
  - `frontend/src/features/repository-binding/application/use-map-module-to-repository.ts`
  - `backend/internal/decisioncenter/service/command_service.go`
  - `backend/internal/review/service/command_service.go`

## ADDED Requirements

### Requirement: `phase08-10` 必须把 review 从“可进入”推进为“可重复闭环”

系统 SHALL 将 `phase08-10` 的正式目标冻结为：用户能够从 review 会话中的反馈事实进入 `Decision` 正式承接位，再至少完成一种既有实体更新路径，并看到成功回流与 reread 结果，而不是停留在只会跳转页面的半成品状态。

#### Scenario: 最小正式闭环的组成

- **WHEN** 团队审查 `phase08-10` 是否达成 DoD
- **THEN** 最小正式闭环必须同时包含：
  - 从 `Daily Review` 或 `Weekly Review` 进入 `Decision` 正式承接位
  - 在 `Decision` canonical 路径中形成新的 `Decision` 或消费既有 `Decision`
  - 至少一种既有实体更新动作真实执行成功
  - 成功后由既有 canonical 页面 reread 最新结果
- **AND** 不得把“只进入了 `Decision Create` 或 `Decision Detail`”视为闭环完成
- **AND** 不得把“只提交了 `SubmitReviewResult`”视为实体更新已完成

#### Scenario: 闭环可重复执行

- **WHEN** 用户多次从 review 会话进入 `Feedback -> Decision -> Update`
- **THEN** 每次都必须沿用同一套 canonical owner、成功回流与刷新语义
- **AND** 不得要求通过手工 SQL、临时脚本或页面刷新补完闭环
- **AND** 不得让一次闭环成功依赖 review-local 临时状态残留

### Requirement: `Decision` 必须继续作为 review loop 的正式中心

系统 SHALL 将 `Decision` 冻结为 `Feedback -> Decision -> Update` 闭环中的正式中心；review 可以引导进入实体更新，但不得绕过 `Decision` 直接把反馈事实解释成并列长期任务系统。

#### Scenario: 从反馈进入正式决策

- **WHEN** 用户在 review 中消费 `current_focus_signals / representative_signals / pending decisions`
- **THEN** 若当前反馈需要形成新的正式经营判断，必须先进入 `Decision Create` 或既有 `Decision Detail`
- **AND** 创建新决策时，成功后必须回流到新建 `Decision` 的 `DecisionDetailPage`
- **AND** 既有决策路径必须继续优先进入 `Decision Center` canonical 页面，而不是在 review 页内联决策编辑工作台

#### Scenario: 决策中心地位不可被实体 handoff 替代

- **WHEN** review footer 暴露 `Product / Module / Repository` handoff
- **THEN** 这些 handoff 只能作为既有 canonical 更新路径的入口
- **AND** 不得把“直接跳到实体页”解释成新的经营判断已经正式承接
- **AND** 当前阶段至少一条成功验收路径必须显式经过 `Decision` 正式承接位

### Requirement: `Decision -> Update` 的最小结果回流必须真实落到既有 canonical update

系统 SHALL 将 `phase08-10` 的最小结果回流冻结为“`Decision` 正式承接 + 至少一种实体 canonical update 成功 + reread 可见”，而不是只停留在 route handoff 或 review record。

#### Scenario: 最小实体更新路径

- **WHEN** 团队实现 `phase08-10`
- **THEN** 至少必须真实打通以下任一最小路径：
  - `Decision -> Module` 关联后，继续进入既有 `Product Detail` 执行 `BindModuleToProduct`
  - 进入既有 `Repository Binding Detail` 执行 `BindRepositoryToProduct`
  - 进入既有 `Repository Binding Detail` 执行 `MapModuleToRepository`
- **AND** 所选路径必须继续复用既有 canonical application owner 与后端 command owner
- **AND** 成功后必须停留在既有 canonical 页面并 reread 最新结果
- **AND** 不得在 review 模块中复制一套影子更新表单或影子写入接口

#### Scenario: route handoff 与正式更新的边界

- **WHEN** review action owner 返回 `go_to_product / go_to_module / go_to_repository`
- **THEN** 这些 success envelope 仍然可以作为进入既有页面的正式入口
- **AND** 但对 `phase08-10` 的验收而言，只有后续 canonical update 真正执行成功并 reread，才算完成 `Update`
- **AND** 单纯到达页面本身不再构成闭环完成证据

### Requirement: 成功回流、错误语义与必要刷新必须继续服从 owner 单值化

系统 SHALL 将 review loop 中新增的成功回流、错误归一化与刷新策略继续收敛到既有 canonical owner 与 `Review action application owner`，而不是在 review page、decision page、实体 detail page 各自补第二套临时编排。

#### Scenario: review 到 decision 的成功回流

- **WHEN** review 触发创建 `Decision draft`
- **THEN** 必须继续复用 `useCreateDraftDecision`
- **AND** 成功后由 canonical owner 负责 `decision-list` 失效，由 review owner 补充 dashboard / review 相关失效
- **AND** 页面只消费 success envelope 执行导航与 toast，不得在 route 里再补一层重复失效

#### Scenario: decision 到实体更新的成功回流

- **WHEN** 既有实体 canonical update 成功
- **THEN** 必须继续由各自 canonical owner 负责 detail / candidate reread 与必要列表失效
- **AND** review 相关页面不得再持有同语义的 page-local `invalidateDetail()`、`mutation.mutate(..., { onSuccess })` 或零散 query key 编排
- **AND** 用户必须能够在成功后直接看到 canonical 页面正文中的最新结果

#### Scenario: 错误语义保持单值

- **WHEN** 闭环中任一步骤失败
- **THEN** 失败语义必须继续沿用对应 canonical owner 暴露的稳定错误
- **AND** review / decision / 实体 detail 页面只负责展示归一化后的错误
- **AND** 不得在页面层直接解析 raw `ConnectError`、transport error 或临时拼装 toast 文案树

### Requirement: `SubmitReviewResult` 必须退回到“过程记录”身份，而不是并列写入主线

系统 SHALL 将 `SubmitReviewResult` 在 `phase08-10` 中继续解释为轻量流程结果记录，而不是代替 `Decision` 或实体 canonical update 的并列业务主线。

#### Scenario: next-step result 的边界

- **WHEN** 用户选择 `submit_next_step`
- **THEN** 该路径仍然只负责记录一条轻量 `review record`
- **AND** 它可以构成“本次 review 有明确后续动作”的证据
- **AND** 但不得被视为实体更新已经完成
- **AND** 不得在 `SubmitReviewResult` 中偷偷承接 `CreateDecision`、`BindModuleToProduct`、`MapModuleToRepository` 一类结构化写入

### Requirement: `phase08-10` 必须回收前后端新增的临时散装编排点

系统 SHALL 将 `phase08` 落地过程中为了 enablement、双会话与临时验收而出现的散装编排点回收到单值 owner 主线，保证前后端都不保留并列临时主线。

#### Scenario: 前端必须回收的临时编排点

- **WHEN** 审查 `phase08-10` 前端实现
- **THEN** 不得继续保留以下模式作为正式主线：
  - review page / route 中额外拼装第二套 query 失效或 success handling
  - review footer、decision page 或实体 detail page 中为 review 临时补写一套 mutation 回调树
  - 通过页面局部状态模拟“已处理 / 已闭环”而不是依赖 canonical reread 结果
- **AND** 所有新增编排都必须能追溯到唯一 owner：`useReviewAction` 或既有 canonical application owner

#### Scenario: 后端必须回收的临时编排点

- **WHEN** 审查 `phase08-10` 后端实现
- **THEN** review 模块不得继续长出代替 `Decision / Product / Repository / Module` canonical command 的临时 service
- **AND** `review.CommandService` 仍然只能承接 `review_records` 这一类过程记录
- **AND** 不得在 handler 或 transport 层临时拼出第二套成功 envelope 或错误语义

### Requirement: `phase08-10` 验收必须覆盖闭环与清理两类证据

系统 SHALL 将 `phase08-10` 的验收口径冻结为“闭环走通 + 编排清理完成”的组合证据，而不是只看构建通过或局部 API 可用。

#### Scenario: 最小验收矩阵

- **WHEN** 团队验证 `phase08-10` 是否通过
- **THEN** 至少必须覆盖：
  - 一条 `Daily Review` 或 `Weekly Review` 进入 `Decision` 的成功路径
  - 一条 `Decision` 正式承接后的成功回流到 `DecisionDetailPage`
  - 一条实体 canonical update 成功并 reread 的路径
  - `SubmitReviewResult` 仍保持轻量记录，不冒充实体更新
  - 前端不再保留第二套 review-local mutation / invalidation 主线
  - 后端不再保留 review-local 并列 command 主线
- **AND** 不得只以 `npm run build`、`go build` 或单次 route 导航作为通过证据

## MODIFIED Requirements

### Requirement: `phase08-09` 的统一动作承接在 `phase08-10` 中必须升级为可闭环动作承接

自 `phase08-10` 起，系统 SHALL 不再只满足“review 页面能统一跳转到 canonical 页面或提交 next-step result”，而必须进一步满足“至少一条动作链已经从 review 经由 `Decision` 正式承接后落到 canonical update 并 reread”的闭环要求。

#### Scenario: 统一动作承接不再停留在 handoff 层

- **WHEN** 审查 `useReviewAction()` 与 review footer 的实现
- **THEN** 可以继续保留 route handoff success envelope
- **AND** 但 `phase08-10` 的通过结论必须建立在至少一条真实 canonical update 闭环之上
- **AND** 不得把 `go_to_product / go_to_repository / go_to_module` 的导航成功本身误判为 `Update` 已完成

### Requirement: `Decision Center` 的最小闭环在 `phase08-10` 中必须承接 review 来源语义

自 `phase08-10` 起，系统 SHALL 要求 `Decision Create / Detail` 继续保持 `Decision Center` 既有 canonical 结构，同时能够承接来自 review 的来源语义、成功回流与后续实体更新入口，而不是要求 review 单独长出第二套决策工作台。

#### Scenario: review 来源下的 Decision canonical 路径

- **WHEN** 用户从 review 进入 `Decision Create` 或 `Decision Detail`
- **THEN** `Decision Center` 仍必须保持既有页面结构、canonical owner 与 reread 模式
- **AND** review 来源参数只能作为来源语义与返回链补充，不得迫使 `Decision Center` 演化为 review 专用页面

## REMOVED Requirements

### Requirement: route handoff 或 review record 可以单独充当 `Feedback -> Decision -> Update` 闭环完成证据

**Reason**: 这会把 `phase08-10` 退化成“可以跳转、可以记一条流程备注”的半闭环，既无法证明 `Decision` 保持经营中心，也无法证明结果已经回流既有实体。
**Migration**: 保留 route handoff 与 `SubmitReviewResult` 作为正式组成部分，但通过结论必须追加至少一条真实 canonical update 成功 + reread 证据。
