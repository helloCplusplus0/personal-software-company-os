# phase10-08 落实 `Onboarding` 首轮建链引导与 canonical handoff Spec

## Why

`phase10-04 / 06 / 07` 已经把 `Onboarding` 的六段主线、前端 owner 分工和后端最小合同承接位冻结下来了，但当前源码仍停留在 `phase06` 的 draft-first 过渡实现：`OnboardingPage` 直接持有四个 create owner、用 URL search 暂存草稿摘要、靠页面级 `invalidateQueries()` 与本地 step 兜底推进。  
如果不把这条主线真正落到前后端承接位、canonical handoff 返回合同和浏览器级交互上，`Onboarding` 仍然会把建链结果留在页面局部状态里，后续补链摩擦也不会真正下降。

## What Changes

- 将 `/onboarding` 从 `phase06` 的 draft-first 过渡页升级为 `phase10` 的六段式首轮建链引导页
- 落实单一 `useOnboardingAction`，回收 `OnboardingPage` 当前直接持有四个 create owner 与页面级失效刷新
- 落实最小后端 `GetOnboardingChainState` 与 `current_product_id` 恢复锚点，支撑刷新、回流与中途中断恢复
- 将 canonical handoff 返回合同正式收敛为 `fromOnboarding / onboardingProductId / onboardingStep`
- 改造 `Product / Repository / Module / Decision Detail` 的 Onboarding 返回链，使 handoff 完成后回到正确 step，而不是继续依赖草稿 search
- 补齐浏览器级验收路径，仅覆盖 `Onboarding` 主线、建链动作与成功后 handoff，不改写 `Dashboard / Review` 的 pending 组装逻辑

## Impact

- Affected specs:
  - `phase10_02_freeze_onboarding_first_run_chain_guidance_action_matrix`
  - `phase10_04_design_onboarding_chain_return_empty_failure_flow`
  - `phase10_06_design_frontend_read_write_owner_route_caller_success_handoff`
  - `phase10_07_design_backend_contract_service_read_model_handoff`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
- Affected code:
  - `frontend/src/routes/onboarding.tsx`
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/features/onboarding/data/use-onboarding-read.ts`
  - `frontend/src/features/onboarding/application/use-onboarding-action.ts`
  - `frontend/src/features/onboarding/lib/onboarding-return.ts`
  - `frontend/src/features/onboarding/lib/onboarding-source-schema.ts`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `proto/psco/onboarding/v1/onboarding.proto`
  - `backend/internal/onboarding/service/query_service.go`
  - `backend/internal/onboarding/connect/server.go`
  - `backend/internal/onboarding/candidate/chain_state_readers.go`
  - `backend/internal/onboarding/repository/recovery_store.go`
  - generated proto / connect 代码与相应测试文件

## ADDED Requirements

### Requirement: `phase10-08` 必须选定单一实现路径，而不是继续保留过渡写法

系统 SHALL 将 `phase10-08` 实现路径冻结为：`/onboarding` 继续作为单一路由宿主，前端正式读取承接位为 `useOnboardingRead`，正式写动作承接位为新增 `useOnboardingAction`，后端正式恢复读合同为新增 `OnboardingService.GetOnboardingChainState`。当前 `OnboardingPage` 内部直接持有 create owner、本地草稿摘要 search 与页面级 step 协调逻辑，只允许被回收，不允许继续扩展。

#### Scenario: 判断 `phase10-08` 是否仍停留在 `phase06` 过渡写法

- **WHEN** 后续实现仍让 `OnboardingPage` 直接 import 四个 create owner、直接调用 `queryClient.invalidateQueries()`、直接在页面里拼成功后下一步
- **THEN** 必须判定为未满足 `phase10-08`
- **AND** 页面级草稿摘要 search 只允许退场，不得作为正式恢复主线继续保留

### Requirement: 后端必须落地 `GetOnboardingChainState` 与最小恢复锚点

系统 SHALL 在 `psco.onboarding.v1` 中正式新增 `GetOnboardingChainState`，并在 `backend/internal/onboarding/service/query_service.go` 落地对应读取编排。由于当前代码库缺少稳定的 product-scoped 恢复锚点，本子任务必须同时落地最小 `onboarding_recovery_store`，把 `current_product_id` 作为后端正式恢复事实源，而不是继续依赖页面 search 或“全局唯一候选”猜测。

#### Scenario: `GetOnboardingChainState` 的最小响应边界

- **WHEN** `/onboarding` 初始化、刷新、从 canonical detail 返回或中途中断再进入
- **THEN** 前端必须通过 `GetOnboardingChainState` 获取最小恢复读模型
- **AND** 该响应至少必须包含：
  - `current_product_id`
  - `current_step`
  - `resume_status`
  - `next_step_kind`
  - 可选的 `canonical_handoff_target`
  - 可选的 `return_hint`
- **AND** 不得把表单局部草稿、页面展开态或临时 UI 状态塞进该读模型

#### Scenario: `current_product_id` 的恢复锚点落地

- **WHEN** `product` 步骤首个 `Product` 创建成功
- **THEN** 后端必须在同一成功路径上冻结最小恢复锚点 `current_product_id`
- **AND** 后续 `repository / module / decision / complete` 的 reread 都必须围绕该锚点解释
- **AND** 不得要求前端继续通过 `productDraftId` 一类 search 参数维持主上下文

#### Scenario: 刷新或回流时的恢复事实源优先级

- **WHEN** `/onboarding` 重新读取链路状态
- **THEN** `onboarding_recovery_store.current_product_id` 必须作为正式优先事实源
- **AND** 仅在该锚点尚不存在时，才允许回退为 `GetFirstRunState` 的冷启动摘要
- **AND** 不得按“最近访问实体”或“全局最新 Product”猜测当前主线

### Requirement: 前端 `useOnboardingRead` 必须升级为六段式链路读取 owner

系统 SHALL 保留 `useOnboardingRead` 作为 `/onboarding` route caller 的单一 page read owner，但其内部实现必须从“只解包 `GetFirstRunState`”升级为“组合 `GetFirstRunState + GetOnboardingChainState` 的六段式链路读取 owner”，并单值产出：

- 当前 step
- 当前主 CTA
- 当前 step 的空态/失败态/完成态 view model
- 当前 step 是否需要 canonical handoff
- 返回链与 resume 所需的只读上下文

#### Scenario: `useOnboardingRead` 的正式职责

- **WHEN** `OnboardingPage` 渲染六段主线
- **THEN** 页面层只允许消费 `useOnboardingRead()` 的只读输出
- **AND** 不得再在页面内并排读取 `GetFirstRunState`、解析草稿 search、再自行拼 step 与状态机

### Requirement: 前端必须落地单一 `useOnboardingAction`

系统 SHALL 新增 `frontend/src/features/onboarding/application/use-onboarding-action.ts`，作为 `/onboarding` 唯一正式写动作承接位。它必须复用既有 canonical create / bind owner，不新增第二套 transport，并返回稳定的 success envelope 供 route caller / page shell 消费。

#### Scenario: `product` 步骤提交

- **WHEN** 用户在 `product` 步骤提交首个 `Product`
- **THEN** `useOnboardingAction` 必须复用既有 `Product` canonical create owner 完成写入
- **AND** 成功后必须：
  - 冻结 `current_product_id`
  - 失效 `ONBOARDING_STATE_QUERY_KEY` 与链路读模型 query key
  - 返回“前进到 `repository`”的 success envelope
- **AND** 页面层不得自己决定产品创建后去哪里

#### Scenario: `repository` 步骤提交

- **WHEN** 用户在 `repository` 步骤提交 `Repository`
- **THEN** `useOnboardingAction` 必须在同一动作内优先尝试完成最小 `Repository -> Product` 关系闭合
- **AND** 若当前承接位无法同页闭合，则 success envelope 只能指向单值 canonical handoff
- **AND** 不得让页面层再判断“留在本页还是跳 detail”

#### Scenario: `module` 步骤提交

- **WHEN** 用户在 `module` 步骤提交 `Module`
- **THEN** `useOnboardingAction` 必须优先尝试完成最小 `Module -> Product` 关系闭合
- **AND** 若 `Repository` 映射仍未闭合，该缺口只能通过 detail CTA 或 canonical handoff 暴露，不得阻断主线进入 `decision`

#### Scenario: `decision` 步骤提交

- **WHEN** 用户在 `decision` 步骤提交首个 `Decision`
- **THEN** `useOnboardingAction` 必须优先尝试完成最小正式承接
- **AND** 若无法同页闭合，则 success envelope 只能指向单值 `Decision Detail` handoff
- **AND** handoff 完成返回后必须进入 `complete`

#### Scenario: `useOnboardingAction` 的最小输出

- **WHEN** 任一步骤写动作成功
- **THEN** `useOnboardingAction` 必须统一返回稳定 envelope，至少包含：
  - `resultKind`
  - `nextStep`
  - 可选的 `navigateTo`
  - 可选的 `params`
  - 可选的 `search`
  - 可选的 `successMessage`
- **AND** 页面不得再直接消费底层 mutation response

### Requirement: canonical handoff 返回合同必须正式切换到 `onboardingProductId`

系统 SHALL 将 `Onboarding` 相关 detail handoff 的正式来源合同冻结为：

- `fromOnboarding=true`
- `onboardingProductId=<current_product_id>`
- `onboardingStep=<repository|module|decision>`

当前 `productDraftId / repositoryDraftId / moduleDraftId / decisionDraftId` 一类 search 参数只允许从 `/onboarding` 主线中退场，不再作为正式恢复合同继续扩散到 detail route 与返回 helper。

#### Scenario: 从 `Onboarding` 进入 canonical detail handoff

- **WHEN** `repository / module / decision` 任一步骤需要进入 canonical detail handoff
- **THEN** 跳转必须带上 `fromOnboarding + onboardingProductId + onboardingStep`
- **AND** detail route `validateSearch` 与返回 helper 必须统一承接这三项
- **AND** 不得继续要求 detail 页面透传草稿摘要 search 才能返回 `/onboarding`

#### Scenario: canonical detail 完成后返回 `Onboarding`

- **WHEN** 用户在 `Product / Repository / Module / Decision Detail` 完成来自 `Onboarding` 的正式 handoff 并点击返回
- **THEN** detail 页必须优先返回 `/onboarding`
- **AND** 返回时必须携带 `onboardingStep`，以便 `OnboardingPage` 在服务端 reread 未完全追平时仍可做一次性 step 恢复
- **AND** 一旦 reread 追平，页面必须重新回到服务端链路状态驱动

### Requirement: `OnboardingPage` 必须从页面级草稿协调改为链路状态驱动

系统 SHALL 重构 `frontend/src/features/onboarding/pages/onboarding-page.tsx`，使它只保留页面壳、步骤 UI 与 owner 消费，不再承担正式业务编排。

#### Scenario: 页面层允许保留的职责

- **WHEN** `OnboardingPage` 实现 `phase10-08`
- **THEN** 页面层只允许保留：
  - 六段 step UI 渲染
  - 成功/失败/空态展示
  - 消费 `useOnboardingRead` 的只读 view model
  - 调用 `useOnboardingAction` 的稳定 action 接口
  - 消费 success envelope 执行最终导航

#### Scenario: 页面层必须回收的职责

- **WHEN** 重构当前 `OnboardingPage`
- **THEN** 至少必须回收：
  - 页面直接 import 四个 create owner
  - `handleStepSuccess()` 内 page-local `invalidateQueries`
  - `drafts` 本地 search 持久化主线
  - `navigateToDetail()` 自行拼 detail 返回合同
  - `resumeServerProgress()` 之类页面级恢复协调

### Requirement: 浏览器级交互必须覆盖主线建链与 handoff

系统 SHALL 为 `phase10-08` 提供浏览器级验收闭环，至少覆盖冷启动、成功建链、canonical handoff、刷新恢复和返回 reread。

#### Scenario: 冷启动首轮建链浏览器流

- **WHEN** 冷启动用户从 `Dashboard` 进入 `/onboarding`
- **THEN** 必须能够顺畅完成 `product -> repository -> module -> decision -> complete`
- **AND** 每一步成功后都只有一个正式下一步或一个正式 canonical handoff

#### Scenario: canonical handoff 浏览器流

- **WHEN** `repository / module / decision` 任一步骤进入 canonical detail handoff
- **THEN** 用户完成 handoff 后必须回到 `/onboarding`
- **AND** 页面必须落到预期 step 或下一正式 step
- **AND** 不得跳到列表页或 `Dashboard`

#### Scenario: 刷新恢复浏览器流

- **WHEN** 用户在 `repository / module / decision` 任一步骤刷新浏览器
- **THEN** 系统必须通过后端链路读模型恢复当前主线
- **AND** 不得丢失 `current_product_id`
- **AND** 不得回退为 `welcome`

## MODIFIED Requirements

### Requirement: `Onboarding` 在 `phase10-08` 中的实现解释

`phase10-08` 修改了当前 `Onboarding` 实现口径：它不再是“phase06 draft-first 录入页 + detail 返回兜底”的过渡方案，而必须成为首轮建链引导与 canonical handoff 的正式实现落点。

#### Scenario: 判断 `Onboarding` 是否仍是草稿过渡页

- **WHEN** 后续实现仍以“先创建四类草稿，再由 detail 慢慢补齐”为核心流程
- **THEN** 必须判定为未满足 `phase10-08`
- **AND** 建链结果必须回到 canonical owner，而不是继续留在页面局部状态

### Requirement: `phase10-08` 的非目标边界

`phase10-08` 只负责 `Onboarding` 主线、建链动作与成功后 handoff，不负责改写 `Dashboard / Review` 的 pending 组装逻辑。

#### Scenario: 判断实现是否越界

- **WHEN** 后续实现试图在同一子任务内同步重写 `Dashboard / Review` 的 pending bucket、pending count 或消费矩阵
- **THEN** 必须判定为超出 `phase10-08` 范围
- **AND** 本子任务只允许做 `Onboarding` 入口、建链、handoff 与返回链所需的最小对接改动

## REMOVED Requirements

### Requirement: `/onboarding` 继续依赖草稿摘要 search 作为正式恢复主线

**Reason**: 这会让建链主上下文继续停留在前端页面局部状态，无法支撑刷新恢复、detail handoff 回流与多实体并存场景。
**Migration**: 用 `GetOnboardingChainState + onboarding_recovery_store.current_product_id + fromOnboarding/onboardingProductId/onboardingStep` 取代现有草稿摘要 search 主线；草稿摘要只保留为过渡兼容字段，随后删除。

### Requirement: `OnboardingPage` 直接持有四个 create owner 与页面级失效刷新

**Reason**: 这会破坏 `phase10-06` 已冻结的 owner 边界，让页面继续承担第二套业务编排。
**Migration**: 新增 `useOnboardingAction`，统一复用既有 canonical create / bind owner，并由 owner 返回 success envelope 与失效策略。
