# phase10-06 产出前端读写 owner、route caller 与成功回流设计 Spec

## Why

`phase10-04` 已冻结 `Onboarding` 建链流、返回链与失败恢复，`phase10-05` 已冻结 `Dashboard / Daily Review / Detail pages` 的下一步动作与 CTA 优先级。当前还缺一份可以直接指导前端实现的 owner / caller 设计，来回答这些页面到底由谁读取、由谁承接写动作、成功后由谁负责回流与 reread、哪些 page-level 临时编排必须回收。
如果这一步不收紧，后续实现很容易继续把 `queryClient.invalidateQueries()`、返回链 search 组装、多个 query 的页面级拼装和直接 `useMutation` 停留在 route / page / panel 内，重新长出第二套 page-local 编排主线。

## What Changes

- 产出 `Onboarding / Dashboard / Review / Detail pages` 的 route caller 与 owner inventory
- 产出需要新增或回收的 read owner / application owner 设计
- 产出失效刷新、成功回流、错误归一化与返回链的正式承接位
- 识别必须回收的页面级散装 mutation / CTA 编排点
- 冻结 caller 与 owner 的一对一映射表，避免页面层继续内联第二套动作编排

## Impact

- Affected specs:
  - `phase10_04_design_onboarding_chain_return_empty_failure_flow`
  - `phase10_05_design_dashboard_review_detail_next_step_cta_handoff`
  - `phase06_07_design_frontend_write_path_mutation_owners`
  - `phase08_06_design_review_frontend_read_write_owner_state_flow`
  - `phase09_06_design_frontend_read_write_owner_state_flow`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
- Affected code:
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/onboarding/data/*`
  - `frontend/src/features/onboarding/application/*`
  - `frontend/src/features/dashboard/data/*`
  - `frontend/src/features/dashboard/application/*`
  - `frontend/src/features/review/data/*`
  - `frontend/src/features/review/application/*`
  - `frontend/src/features/product-registry/data/*`
  - `frontend/src/features/product-registry/application/*`
  - `frontend/src/features/module-registry/data/*`
  - `frontend/src/features/module-registry/application/*`
  - `frontend/src/features/repository-binding/data/*`
  - `frontend/src/features/repository-binding/application/*`
  - `frontend/src/features/decision-center/data/*`
  - `frontend/src/features/decision-center/application/*`

## ADDED Requirements

### Requirement: route caller 必须退回为薄适配层

系统 SHALL 将 `Onboarding / Dashboard / Review / Detail pages` 的 route caller 冻结为薄适配层：只承接 `params / search` 解释、调用单一页面 read owner 与单一页面 action owner、消费成功 envelope 做导航与 toast，不得直接在 route / page 内继续拼第二套 query 或 mutation 主线。

#### Scenario: route caller 的正式职责边界

- **WHEN** 后续实现某个 page route caller
- **THEN** 它只允许承接：
  - 读取 route params / search
  - 将 params / search 传给单一页面 read owner 与单一页面 action owner
  - 消费 action owner 返回的成功 envelope 执行 `navigate()`、toast 或局部 UI 回流
  - 渲染 page shell / error boundary / skeleton
- **AND** 不得在 route caller 内直接声明第二组业务 query
- **AND** 不得在 route caller 内直接声明 `useMutation`
- **AND** 不得在 route caller 内直接手写 `queryClient.invalidateQueries()` 作为正式业务编排

#### Scenario: `phase10` route caller 与 owner 的一对一映射

- **WHEN** 后续实现 `phase10` 相关页面
- **THEN** 至少必须满足以下 caller-owner 映射：
  - `/onboarding` route caller -> `useOnboardingRead` + `useOnboardingAction`
  - `/dashboard` route caller -> `useDashboardHomeRead` + `useDashboardPrimaryAction`
  - `/reviews/daily` route caller -> `useDailyReviewRead` + `useReviewAction`
  - `/products/$productId` route caller -> `useProductDetailPageRead` + `useProductDetailActions`
  - `/modules/$moduleId` route caller -> `useModuleDetailPageRead` + `useModuleDetailActions`
  - `/repositories/$repositoryId` route caller -> `useRepositoryDetailPageRead` + `useRepositoryDetailActions`
  - `/decisions/$decisionId` route caller -> `useDecisionDetailPageRead` + `useDecisionDetailActions`
- **AND** 一个 route caller 不得同时直接依赖多个底层 query / mutation owner 再由页面补一层临时编排

### Requirement: `query` 层必须只承接页面级读取与只读派生

系统 SHALL 将 `query` 层继续冻结为只读承接位。若某个页面当前已经存在多个 query 的 page-level 编排，则必须回收到单一 page read owner 中，但该 owner 仍只能承接读取、queryKey、状态派生与只读 view model，不得混入 mutation 语义。

#### Scenario: page read owner 的最小职责

- **WHEN** 某个页面需要新增或回收 page read owner
- **THEN** 该 owner 只允许承接：
  - 消费底层 canonical read owner
  - 派生 `initial-loading / ready / page-error` 或等价页面级状态
  - 派生区块级 `ready / empty / error` 状态
  - 产出页面所需的只读 view model
  - 产出当前页面的 `primaryCta / secondaryCtas / actionDescriptor`
  - 产出 route caller 可直接消费的只读返回链上下文
- **AND** 不得在 `data/` 层组装导航目标
- **AND** 不得在 `data/` 层承接 query invalidation
- **AND** 不得在 `data/` 层混入 `create / update / bind / link` 等写动作

#### Scenario: `Dashboard` 的 page read owner 收口

- **WHEN** `Dashboard Home` 继续需要消费 `overview / feedback signals / recent activities / reuse summary`
- **THEN** 页面层不得继续直接并排调用四个 query hook
- **AND** 必须新增单一 `useDashboardHomeRead`
- **AND** 该 owner 必须继续复用既有 `useDashboardOverviewRead / useFeedbackSignalsRead / useRecentActivitiesRead / useReuseSummaryRead`
- **AND** `useDashboardHomeRead()` 必须单值产出当前 `Dashboard` 的 `primaryCta / secondaryCtas / actionDescriptor`
- **AND** `DashboardHomePage` 只允许消费 `useDashboardHomeRead()` 的页面级状态与只读 props

#### Scenario: 各 Detail 页的 page read owner 收口

- **WHEN** `Product / Module / Repository / Decision Detail` 页面需要同时消费详情读取、来源 search 派生、复用摘要或模板来源摘要
- **THEN** 页面层必须只通过对应单一 page read owner 消费这些结果：
  - `useProductDetailPageRead`
  - `useModuleDetailPageRead`
  - `useRepositoryDetailPageRead`
  - `useDecisionDetailPageRead`
- **AND** page read owner 必须继续复用既有 canonical read owner，例如 `useProductDetailRead / useModuleDetailRead / useRepositoryDetailRead / useDecisionDetailRead`
- **AND** page read owner 必须单值产出当前 detail 页的 `primaryCta / secondaryCtas / actionDescriptor`
- **AND** `ProductDetailPage` 不得继续页面级并排调用 `useProductDetailRead + useReuseSummaryRead + useTemplateSourceRead`
- **AND** `ModuleDetailPage` 不得继续页面级并排调用 `useModuleDetailRead + useReuseSummaryRead`

### Requirement: `application` 层必须成为页面动作与成功回流的唯一正式承接位

系统 SHALL 将页面动作、成功回流、错误归一化与 query 失效策略收敛到单一 page action owner 中。页面只消费稳定的 action 接口与成功 envelope，不再直接 import 多个 canonical mutation owner 或自己组装返回链。

#### Scenario: page action owner 的最小接口

- **WHEN** 后续实现某个 page action owner
- **THEN** 它至少必须暴露：
  - 一个稳定的主动作提交接口，例如 `submitAction()` 或等价命名
  - `isPending`
  - `error`
  - `reset()`
- **AND** 成功结果必须是页面可消费的稳定 envelope，而不是裸传底层 mutation response
- **AND** 该成功 envelope 至少必须包含：
  - `resultKind`
  - `navigateTo`
  - `params`
  - `search`
  - 可选的 `successMessage`
- **AND** 页面不得自己把多个 mutation response 再拼成第二套回流协议

#### Scenario: `Onboarding` 的 action owner 只能编排既有 canonical 写路径

- **WHEN** `Onboarding` 需要在 `product / repository / module / decision` 四个 step 触发创建、绑定或 handoff
- **THEN** 必须由单一 `useOnboardingAction` 承接页面动作编排
- **AND** 该 owner 只能复用既有 canonical application owner：
  - `useCreateDraftProduct`
  - `useCreateDraftRepository`
  - `useCreateDraftModule`
  - `useCreateDraftDecision`
  - 以及对应 canonical binding owner
- **AND** 不得为 `Onboarding` 新增第二套 transport 或第二套 feature-local mutation owner
- **AND** `OnboardingPage` 不得继续直接持有四个 create owner 与 `queryClient.invalidateQueries()`

#### Scenario: `Dashboard` 与 `Review` 的 action owner 分工

- **WHEN** `Dashboard` 或 `Daily Review` 需要响应主 CTA 或正式下一步动作
- **THEN** `Dashboard` 只允许通过 `useDashboardPrimaryAction` 产出 route handoff success envelope
- **AND** `Daily Review` 继续只允许通过 `useReviewAction` 承接写动作与 route handoff
- **AND** `Dashboard` 不得为了“主 CTA 一致性”复制 `Review action owner`
- **AND** `Review` 不得回退为页面级直接拼 `navigate + mutation + invalidateQueries`

#### Scenario: 各 Detail 页的 action owner 分工

- **WHEN** `Product / Module / Repository / Decision Detail` 需要承接 CTA、状态推进、关系绑定、返回链与 reread
- **THEN** 页面层只允许通过对应单一 action owner 承接：
  - `useProductDetailActions`
  - `useModuleDetailActions`
  - `useRepositoryDetailActions`
  - `useDecisionDetailActions`
- **AND** 这些 owner 只承接页面级 CTA handoff、状态推进、成功回流、返回链与 reread 编排
- **AND** 这些 owner 必须继续复用既有 canonical mutation owner，而不是复制 transport
- **AND** 已冻结为 canonical action surface 的子组件或 panel，允许继续直接消费其既有 canonical owner
- **AND** page action owner 不得为了形式统一而再包一层重复 owner 去替代这些 canonical panel owner
- **AND** 页面层不得继续直接 import `useUpdateDecisionStatus`
- **AND** 页面层不得继续自己声明 `invalidateDetail()` 作为正式 reread 主线

### Requirement: caller 与 owner 的文件落点必须最小化且可追溯

系统 SHALL 为 `phase10-06` 冻结最小新增 / 回收文件落点，保证 caller 与 owner 都能被单值追溯，而不是通过分散 helper、临时 callback 与 page-local function 混合完成。

#### Scenario: 最小新增 / 回收落点

- **WHEN** 后续开始实现 `phase10-06`
- **THEN** 至少必须满足以下最小落点：
  - `frontend/src/features/onboarding/application/use-onboarding-action.ts`
  - `frontend/src/features/dashboard/data/use-dashboard-home-read.ts`
  - `frontend/src/features/dashboard/application/use-dashboard-primary-action.ts`
  - `frontend/src/features/product-registry/data/use-product-detail-page-read.ts`
  - `frontend/src/features/product-registry/application/use-product-detail-actions.ts`
  - `frontend/src/features/module-registry/data/use-module-detail-page-read.ts`
  - `frontend/src/features/module-registry/application/use-module-detail-actions.ts`
  - `frontend/src/features/repository-binding/data/use-repository-detail-page-read.ts`
  - `frontend/src/features/repository-binding/application/use-repository-detail-actions.ts`
  - `frontend/src/features/decision-center/data/use-decision-detail-page-read.ts`
  - `frontend/src/features/decision-center/application/use-decision-detail-actions.ts`
- **AND** 已有 `useOnboardingRead / useDailyReviewRead / useReviewAction` 优先复用
- **AND** 不得为了形式对称新增无职责的新切片

### Requirement: 成功回流、失效刷新与返回链必须由 owner 单值化

系统 SHALL 将成功回流、失效刷新、错误归一化与返回链单值化到对应 owner 中，并冻结为“canonical owner 负责自身切片失效，page action owner 负责页面级 reread 与返回链 envelope”的分层模式。

#### Scenario: canonical owner 与 page action owner 的失效分工

- **WHEN** 某个页面动作复用既有 canonical mutation owner 成功完成
- **THEN** canonical owner 必须继续负责自身切片的正式失效
- **AND** page action owner 只允许补充页面级 reread 所需的额外失效
- **AND** 页面层不得再直接写第二组 `invalidateQueries`
- **AND** 不得把所有失效都提升到 route caller 中统一手写

#### Scenario: `Onboarding` 成功回流与 reread

- **WHEN** `Onboarding` 某一步创建、绑定或 handoff 成功
- **THEN** `useOnboardingAction` 必须返回单值 success envelope
- **AND** 该 envelope 只能指向：
  - 下一正式 step
  - 对应 canonical detail handoff
  - `complete`
  - 返回 `Dashboard`
- **AND** `useOnboardingAction` 必须统一补充 `ONBOARDING_STATE_QUERY_KEY` 或等价恢复读模型的失效
- **AND** 页面层不得自己决定成功后落下一个 step

#### Scenario: `Dashboard / Review` 的成功回流

- **WHEN** `Dashboard` 或 `Review` 的正式动作成功
- **THEN** 对应 action owner 必须统一返回：
  - 目标 canonical route
  - 需要透传的 `fromDashboard / dashboardSection / dashboardReturnTo`
  - 对应 reread 所需的失效目标
- **AND** 页面层只负责执行最终导航与 toast
- **AND** 不得在 `DashboardHomePage` 或 `DailyReviewPage` 内重新拼接第二套来源 search

#### Scenario: 普通返回链与成功回流的承接位分工

- **WHEN** 页面只是处理普通返回按钮、返回列表或返回上游 detail 的被动返回
- **THEN** 返回链上下文必须由对应 page read owner 单值产出
- **AND** route caller 或 page shell 只消费这份只读返回上下文执行导航
- **AND** 这类普通返回不得要求先经过 page action owner 产出 success envelope
- **WHEN** 页面是在某个正式动作成功后离开当前页
- **THEN** success search 与回流 envelope 才允许由 page action owner 统一组装
- **AND** 不得把普通返回与动作成功回流混成同一承接位

#### Scenario: 各 Detail 页的成功回流

- **WHEN** `Product / Module / Repository / Decision Detail` 中的正式动作成功
- **THEN** 对应 action owner 必须统一回答：
  - 当前页是否应 stay on page reread
  - 是否应切换新的主 CTA
  - 是否需要返回 `Dashboard / Review / Onboarding`
  - 返回时应透传的 search 上下文
- **AND** `ProductDetailPage / RepositoryBindingDetailPage` 不得继续以 page-local `invalidateDetail()` 作为唯一 reread 方案
- **AND** `DecisionDetailPage` 不得继续直接在页面里组装状态推进与返回链

#### Scenario: 错误归一化的正式承接位

- **WHEN** 页面动作失败
- **THEN** 错误归一化必须发生在对应 application owner 内
- **AND** route caller / page 只允许消费归一化后的 `error`
- **AND** 页面不得再直接拼接原始 `error.message` 形成第二套错误语义

### Requirement: 页面级散装 mutation / CTA 编排点必须被识别并回收

系统 SHALL 明确 `phase10-06` 当前必须回收的 page-level 散装 mutation / CTA 编排点，避免这些过渡写法继续扩散到新的 `phase10` 页面实现中。

#### Scenario: 必须回收的 `Onboarding` 页面级编排

- **WHEN** 后续开始实现 `phase10-06`
- **THEN** 必须回收 `frontend/src/features/onboarding/pages/onboarding-page.tsx` 中至少以下 page-level 编排：
  - 直接 import 四个 create owner 的模式
  - `handleStepSuccess()` 中页面级 `invalidateQueries`
  - 页面级 `navigateToDetail()` 与 step 成功后默认下一步拼装
  - 页面级 `resumeServerProgress()` 与一次性 step 恢复协调
- **AND** 页面只允许保留本地 UI 细节与布局编排

#### Scenario: 必须回收的 `Dashboard` 页面级编排

- **WHEN** 后续实现 `Dashboard` 的 `phase10` 主 CTA 主线
- **THEN** 必须回收 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` 中至少以下 page-level 编排：
  - 并排调用四个 read owner 的模式
  - 页面级整页状态拼装
  - `handleFullRetry()` 内直接 `invalidateQueries`
  - 页面级主 CTA 选择与返回区块恢复的混合编排

#### Scenario: 必须回收的各 Detail 页编排

- **WHEN** 后续实现 `Product / Module / Repository / Decision Detail`
- **THEN** 至少必须回收以下模式：
  - `ProductDetailPage` 的 page-local `invalidateDetail()` 与多 query 并排读取
  - `RepositoryBindingDetailPage` 的 page-local `invalidateDetail()` 与动作成功后的来源回流拼装
  - `DecisionDetailPage` 页面直接 import `useUpdateDecisionStatus` 并声明 `handleStatusChange()`
  - `ModuleDetailPage` 页面级返回链分支与 CTA 去向拼装
- **AND** 页面层仍允许保留纯 UI 状态，例如 `panelMode`、展开/收起、scroll/focus
- **AND** 这些 UI 状态不得重新承担正式业务动作编排

### Requirement: 页面层不得继续内联第二套动作编排

系统 SHALL 冻结一个跨页面统一规则：页面层可以保留布局、局部 UI 状态、成功提示展示与最终导航消费，但不得继续内联第二套动作编排。

#### Scenario: 页面层允许保留的职责

- **WHEN** 实现 `Onboarding / Dashboard / Review / Detail pages`
- **THEN** 页面层只允许保留：
  - page shell 与布局
  - skeleton / empty / error 展示
  - 局部 UI state，例如 tab、panel、accordion、scroll、focus
  - 消费 owner 返回的 success envelope 执行导航
  - toast 展示

#### Scenario: 页面层禁止保留的职责

- **WHEN** 页面层试图继续承接业务动作
- **THEN** 必须判定为未满足 `phase10-06`，包括但不限于：
  - 直接 import 两个以上 canonical mutation owner 再由页面做主次 CTA 选择
  - 直接写 `queryClient.invalidateQueries()` 作为正式业务刷新主线
  - 直接拼装动作成功后的 `fromDashboard / fromOnboarding / fromList / returnTo` 业务回流协议
  - 直接根据多个 query 原始状态重新拼页面级第二套状态机

## MODIFIED Requirements

### Requirement: `phase10-05` 中“页面如何显示 CTA”的实现承接方式

`phase10-05` 已冻结了 `Dashboard / Daily Review / Detail pages` 的 CTA inventory、优先级与 reread 语义。
自 `phase10-06` 起，这些 CTA 的正式实现承接方式必须修改为：页面不再直接实现 CTA 编排，而是通过 page read owner + page action owner 的一对一组合消费这些规则。

#### Scenario: 判断 `phase10-05` 是否仍停留在页面直写

- **WHEN** 后续实现或验收发现某个页面虽然显示了正确 CTA，但 CTA 判定、跳转目标、失效刷新与成功回流仍停留在 page-local 逻辑
- **THEN** 不得视为满足 `phase10-06`
- **AND** `primaryCta / secondaryCtas / actionDescriptor` 必须由 page read owner 单值产出
- **AND** page action owner 只消费该 `actionDescriptor` 执行正式动作与成功回流

### Requirement: `phase06-07 / phase08-06 / phase09-06` 中既有 owner 的复用方式

`phase10-06` 修改了既有 owner 设计在 `Asset-Action Closure` 页面中的消费方式：既有 canonical owner 继续保留其 transport、query 失效与错误归一化职责，但页面层不得再直接消费它们，必须由 `phase10` 的 page action owner 做一层正式页面编排。

#### Scenario: 判断既有 canonical owner 是否被页面直接消费

- **WHEN** `Onboarding / Dashboard / Review / Detail pages` 仍在 page route caller 或 page 组件层直接 import 多个既有 canonical owner
- **THEN** 必须判定为未满足 `phase10-06`
- **AND** 这些既有 owner 只能作为 `phase10` page action owner 的内部依赖
- **AND** 已冻结为 canonical action surface 的子组件或 panel 不受该限制，允许继续直接消费其既有 canonical owner

## REMOVED Requirements

### Requirement: 用 page-local `invalidateDetail()` 或 `handleFullRetry()` 维持正式 reread 主线

**Reason**: 这会把页面再次变成隐式业务编排宿主，破坏 caller-owner 一对一映射与 query / application 边界。
**Migration**: 将页面级 invalidation 迁移到对应 page action owner；canonical owner 继续负责自身切片失效，page action owner 只补充页面级 reread 所需额外失效。

### Requirement: 由页面自己拼装 `fromOnboarding / fromDashboard / fromList / returnTo` 返回链协议

**Reason**: 这会让成功回流协议散落在多个 page 与 panel 中，后续无法机械验证。
**Migration**: 将动作成功后的返回链 search 组装统一迁移到对应 page action owner；普通返回按钮所需的只读返回上下文由 page read owner 产出，route caller 或 page shell 直接消费。
