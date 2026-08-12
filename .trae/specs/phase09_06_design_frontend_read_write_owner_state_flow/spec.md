# phase09-06 产出前端读写承接位与状态流设计 Spec

## Why

`phase09-04` 已冻结模板读合同、前后端 owner 与 caller inventory，`phase09-05` 已冻结模板与提示的页面流、交互流和返回链。当前还缺一份可直接指导源码实现的前端承接位设计，来回答模板候选与提示的 read layer 落在哪、`Product Create` 预填 handoff 由谁编排、成功回流与 query 失效由谁负责，以及哪些页面级临时编排必须回收。否则后续实现很容易把模板逻辑重新散落到 `Weekly Review / Product Create / Product Detail` 页面内，长出第二套 page-local owner 与第二套 create form state 主线。

## What Changes

- 冻结模板候选、模板预填、模板来源复读与派生提示的前端 read layer 切片落点
- 冻结 `Product Create` 模板预填 handoff 的 application owner 设计与既有 canonical mutation owner 的边界
- 冻结成功回流、query 失效、错误反馈与 unavailable 成功态的正式承接位
- 冻结必须回收的页面级临时编排点与禁止继续复制的模式
- 冻结 caller 与 owner 的一对一映射表
- 冻结 `Weekly Review` 新增消费位与既有 `reuseSnapshot / representativeSignals` 的边界矩阵
- 冻结 `Product Detail` 模板来源复读上下文的消费方式与 canonical binding path 的继续导向

## Impact

- Affected specs:
  - `phase09_04_freeze_contract_read_model_owner_candidate_source_boundary`
  - `phase09_05_design_template_reuse_hint_page_flow_interaction_return_chain`
  - `phase08_06_design_review_frontend_read_write_owner_state_flow`
  - `phase06_07_design_frontend_write_path_mutation_owners`
  - `phase07_05_design_frontend_generated_client_query_application_migration`
- Affected code:
  - 后续新增 `frontend/src/features/template-reuse/data/*`
  - 后续新增 `frontend/src/features/template-reuse/application/*`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/product-registry/application/use-create-draft-product.ts`
  - 后续新增 `frontend/src/features/product-registry/application/use-product-create-form-state.ts`
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`
  - `frontend/src/features/product-registry/components/product-create-form.tsx`
  - 后续新增 `frontend/src/features/product-registry/application/use-product-create-template-handoff.ts`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/reuse-summary/data/use-reuse-summary-read.ts`

## ADDED Requirements

### Requirement: phase09 前端读能力必须收敛到单一 `template-reuse` 切片与既有 canonical read owner

系统 SHALL 将 `phase09` 新增模板读能力冻结为单一业务切片 `frontend/src/features/template-reuse/`，并要求它只承接模板候选、模板预填、模板来源复读与派生提示读取，不复制既有 `review / product-registry / reuse-summary` 的 canonical read owner。

#### Scenario: template-reuse 切片的最小物理落点

- **WHEN** 后续实现开始落地 `phase09` 前端读能力
- **THEN** 必须至少新增以下正式落点：
  - `frontend/src/features/template-reuse/data/use-template-candidates-read.ts`
  - `frontend/src/features/template-reuse/data/use-template-prefill-read.ts`
  - `frontend/src/features/template-reuse/data/use-template-source-read.ts`
  - `frontend/src/features/template-reuse/data/use-derived-insight-hints-read.ts`
  - `frontend/src/features/template-reuse/data/template-reuse-query-options.ts` 或等价只读 helper
- **AND** `template-reuse/data/` 只承接读取、queryKey、响应解包、成功空态与 unavailable 成功态的只读派生
- **AND** 不得在 `template-reuse/data/` 混入 create / bind / apply 一类写动作
- **AND** `product-registry/data/` 不得新增任何模板预填、模板来源复读或提示读取 owner

#### Scenario: template-reuse 与既有 canonical read owner 的边界

- **WHEN** `phase09` 需要消费 `module_reuse_summary / capability_summary`
- **THEN** 只能继续通过既有 `useReuseSummaryRead()` 消费
- **AND** `template-reuse` 不得复制第二套 `reuse snapshot` query owner
- **WHEN** `phase09` 需要消费 `Weekly Review` 页面级上下文
- **THEN** 页面宿主仍必须继续只通过 `useWeeklyReviewRead()` 消费
- **AND** `template-reuse` 只能作为其内部被组合的 slice-local read contract，而不是页面层并排直接调用的第二 page-level query 主线

### Requirement: Weekly Review 的模板新增消费位必须继续收敛到单一 `useWeeklyReviewRead`

系统 SHALL 冻结 `Weekly Review` 的新增模板消费位只能由 `useWeeklyReviewRead()` 组合承接，页面层不得直接并排消费 `template-reuse` 切片的底层 read hooks。

#### Scenario: Weekly Review 的新增消费位组合方式

- **WHEN** `Weekly Review` 需要新增模板候选、`reuse_opportunity_hint` 与 `capability_gap_hint`
- **THEN** `useWeeklyReviewRead()` 必须继续是页面唯一正式 read owner
- **AND** 它可以在切片内部组合：
  - `template-reuse` 的模板候选读取
  - `template-reuse` 的派生提示读取
  - 既有 `review` 页面组合结果
- **AND** `WeeklyReviewPage` 不得直接新增第二个 page-level `useQuery`

#### Scenario: Weekly Review 与既有 reuseSnapshot / representativeSignals 的边界矩阵

- **WHEN** `Weekly Review` 同时展示既有 `reuseSnapshot / representativeSignals` 与新增模板能力
- **THEN** 边界必须保持为：
  - `representativeSignals`：继续由 `review` 既有上下文承接
  - `reuseSnapshot`：继续由既有 `module_reuse_summary / capability_summary` 承接
  - `templateCandidates / reuse_opportunity_hint / capability_gap_hint`：只由 `template-reuse` 读层承接
- **AND** 模板候选不得反向改写 `reuseSnapshot` 的只读语义
- **AND** `reuseSnapshot` 不得升级为模板候选排序、切换或 handoff 的 owner

### Requirement: Product Create 的模板 handoff 必须由单一 application owner 编排，但不得侵入 canonical create mutation owner

系统 SHALL 为 `Product Create` 模板预填 handoff 冻结单一 application owner，用来承接 search 参数解释、模板预填只读结果拼装、返回链上下文与 create 成功回流上下文；但正式创建写动作仍必须只通过既有 `useCreateDraftProduct()`。

#### Scenario: Product Create handoff application owner 的最小落点

- **WHEN** 后续实现 `Product Create` 的模板 handoff 编排
- **THEN** 必须新增单一正式承接位，例如：
  - `frontend/src/features/product-registry/application/use-product-create-template-handoff.ts`
- **AND** 该 owner 至少必须承接：
  - 解析 `fromTemplateReuse / templateCandidateId / templateSource`
  - 组合 `use-template-prefill-read` 的只读结果
  - 组合 create 页面需要的模板摘要 view model
  - 透传 `capability_gap_hint` 的补齐返回链上下文
  - 组装创建成功后进入 `Product Detail` 的模板来源 search 参数

#### Scenario: Product Create canonical mutation owner 不被模板逻辑侵入

- **WHEN** 用户在模板预填上下文中提交创建
- **THEN** 正式写动作必须继续只通过 `useCreateDraftProduct()`
- **AND** `use-product-create-template-handoff` 不得内联第二套 `useMutation`
- **AND** `useCreateDraftProduct()` 不得被模板逻辑扩写为同时承接模板读取、模板提示读取或模板来源摘要拼装
- **AND** 不得新增 `useCreateProductFromTemplate`、`useApplyTemplateToProduct` 或等价第二套 create mutation owner

### Requirement: Product Create 页面不得因为模板预填而长出第二套 create form state 主线

系统 SHALL 冻结模板预填只允许作为既有 create form state 的初始值或受控外部输入，不得在页面内并行再长出一套独立模板表单状态机。

#### Scenario: 模板预填如何进入 create form

- **WHEN** `Product Create` 通过 `templateCandidateId` 读取到模板预填
- **THEN** 预填值只能进入既有 `ProductCreateForm` 所消费的单一表单状态主线
- **AND** 模板摘要、模板来源标记与 `capability_gap_hint` 必须作为表单外部上下文展示
- **AND** 不得在页面内并行维护“模板原始值 store + 表单编辑值 store”两套长期状态

#### Scenario: create form state 的唯一正式承接位

- **WHEN** 后续实现需要满足“从 Module 补齐页返回后恢复当前表单草稿”
- **THEN** 必须将当前 `ProductCreateForm` 的字段级本地 `useState` 升级为单一正式 create form state owner，例如：
  - `frontend/src/features/product-registry/application/use-product-create-form-state.ts`
- **AND** 该 owner 必须成为 `name / description / status` 以及后续 create 字段的唯一正式状态主线
- **AND** `ProductCreatePage`、`ProductCreateForm` 与 `use-product-create-template-handoff` 只能共同消费这一条状态主线
- **AND** 不得保留“组件本地 `useState` + 额外临时 store”并行双轨

#### Scenario: 用户编辑预填后的 create 草稿

- **WHEN** 用户修改预填后的名称、描述或状态
- **THEN** 修改结果必须直接落在既有 create form state 中
- **AND** 模板 handoff owner 只允许保留模板来源上下文，不得接管字段级表单编辑状态
- **AND** 页面刷新、失败重试或从缺口补齐页返回时，仍必须恢复到同一条 create form state 主线

#### Scenario: 从 Module 补齐页返回时如何恢复 create 草稿

- **WHEN** 用户从 `Product Create` 进入 `Module Registry / Module Detail`，随后返回原 create 会话
- **THEN** 草稿恢复必须只依赖上述单一 create form state owner
- **AND** 该 owner 可以使用单一受控的临时状态承载跨页面返回期间的字段值
- **AND** 该受控临时状态若存在，也必须被视为 formal create form state 本身，而不是附加在组件本地 state 之外的第二套草稿源
- **AND** `templateCandidateId / templateSource / fromTemplateReuse` 继续只承担模板上下文与返回链语义，不承担字段级草稿存储
- **AND** 不得改为把字段草稿分散存进 search 参数、页面局部 state 与额外 store 的混合方案

### Requirement: Product Create 的 query 失效、成功回流与错误反馈必须由 owner 分层承接

系统 SHALL 把模板 create 场景的失效、回流与错误处理冻结为“canonical mutation owner 负责 create 相关失效，handoff owner 负责模板来源回流参数与模板只读上下文恢复”的分层组合。

#### Scenario: 创建成功后的失效与回流分工

- **WHEN** 用户通过模板预填成功创建 Product
- **THEN** `useCreateDraftProduct()` 继续只负责既有 create 成功后的 query 失效
- **AND** `use-product-create-template-handoff` 只负责生成跳转到 `/products/$productId` 时的：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
  - 可选的 `fromDashboard / dashboardSection / dashboardReturnTo`
- **AND** `ProductCreatePage` 只消费这两类 owner 的结果执行最终导航与 toast

#### Scenario: 模板预填 unavailable 成功态与请求失败态的承接位

- **WHEN** `templateCandidateId` 不可解析并返回 `UNAVAILABLE`
- **THEN** `use-template-prefill-read()` 必须将其派生为成功可恢复状态
- **AND** `use-product-create-template-handoff` 必须把该状态转换为“模板摘要 unavailable + 空白 create 仍可编辑”的页面级 view model
- **WHEN** 模板预填读取发生真实失败
- **THEN** `template-reuse` 读层必须只导出局部 error 状态
- **AND** `ProductCreatePage` 不得把整页降级为 page error

### Requirement: Product Detail 的模板来源复读必须保持只读，并继续导向 canonical binding path

系统 SHALL 冻结 `Product Detail` 对模板来源复读的消费方式为“新增单一 read owner + 页面级只读摘要消费”，不得长出第二套详情写路径。

#### Scenario: Product Detail 的模板来源 read owner

- **WHEN** `/products/$productId` 需要根据成功回流参数复读模板来源摘要
- **THEN** 页面必须只通过单一正式 read owner 消费，例如：
  - `frontend/src/features/template-reuse/data/use-template-source-read.ts`
- **AND** 页面层不得直接拼 `templateCandidateId` 请求
- **AND** 该 read owner 只承接模板来源摘要与 unavailable 成功态派生，不承接任何写动作

#### Scenario: Product Detail 继续导向 canonical binding path

- **WHEN** 页面成功展示模板来源摘要或 unavailable 空态
- **THEN** 后续正式下一步动作仍然只能导向既有 `ProductModuleBindingPanel`
- **AND** 模板来源摘要区不得内联自动绑定、批量写入或第二套 binding mutation
- **AND** `ProductDetailPage` 不得因为新增模板来源复读而再声明第二套 detail-local mutation owner

### Requirement: caller 与 owner 必须形成一对一映射，禁止跨页面漂移

系统 SHALL 为 `phase09` 前端消费链产出稳定的 caller-owner 映射表，保证每个 caller 都能追溯到唯一正式 owner。

#### Scenario: phase09 caller-owner 映射表

- **WHEN** 后续实现 `phase09` 前端
- **THEN** 至少必须满足以下单值映射：
  - `WeeklyReviewPage` -> `useWeeklyReviewRead`
  - `WeeklyReview` 内模板候选区 / 提示区 -> 只消费 `useWeeklyReviewRead` 派生 props
  - `ReviewActionFooter` 或等价完成区 -> `useReviewAction`
  - `ProductCreatePage` -> `useCreateDraftProduct` + `use-product-create-template-handoff` + `use-product-create-form-state`
  - `ProductCreateForm` -> 只消费 `use-product-create-form-state` 提供的字段值、字段更新函数与 submit callback
  - `Product Detail` 模板来源摘要区 -> `use-template-source-read`
  - `ProductModuleBindingPanel` -> 继续只消费既有 canonical binding owner
- **AND** 不得出现“一个 caller 同时直接依赖多个底层 query/mutation owner，再由页面补第二层编排”的并列状态

### Requirement: 必须识别并回收 page-level 临时编排点，避免复制第二套 owner

系统 SHALL 明确 `phase09` 当前实现中哪些页面级临时编排点必须回收，后续不得把同样模式继续复制到模板链路。

#### Scenario: 必须回收或禁止继续扩写的临时编排点

- **WHEN** 后续实现开始落地 `phase09-06`
- **THEN** 至少必须识别并回收或限制以下模式：
  - `ProductCreatePage` 中直接组装模板 search 参数、成功回流 search 与模板上下文拼装的页面级逻辑
  - `ProductCreatePage` 中通过 `mutation.mutate(..., { onSuccess, onError })` 继续扩写模板语义的模式
  - `ProductCreateForm` 中继续以本地 `useState` 单独保存正式 create 草稿字段的模式
  - `ProductDetailPage` 中直接新增第二个 page-level 模板读取 query 的模式
  - `WeeklyReviewPage` 中直接并排新增模板相关 query hook 的模式
- **AND** 页面层只允许保留：
  - 导航消费
  - toast 展示
  - 组件编排
  - UI 局部展开/收起状态
- **AND** 不得继续把模板 handoff、query 失效与错误归一化停留在 page-local 逻辑中

### Requirement: query 与 application 边界必须保持实现可落地但不过度冻结

系统 SHALL 保持 `query` 与 `application` 的职责边界清晰，同时不提前冻结具体库 API 的实现写法。

#### Scenario: query 层的正式职责

- **WHEN** 某个 `phase09` 前端 owner 被放入 `data/`
- **THEN** 它只允许承接：
  - queryKey
  - 只读请求
  - 响应解包
  - success empty / unavailable / error 的只读派生
- **AND** 不得在 `data/` 层混入导航、toast、失效矩阵或 mutation 编排

#### Scenario: application 层的正式职责

- **WHEN** 某个 `phase09` 前端 owner 被放入 `application/`
- **THEN** 它只允许承接：
  - 多个 read owner 的消费编排
  - success handoff search 组装
  - query 失效策略组合
  - 错误归一化
  - 页面回流语义
- **AND** 不得在 `application/` 层重新声明第二套长期表单字段状态

## MODIFIED Requirements

### Requirement: Product Create 的模板承接方式

`phase09-05` 已冻结 `Product Create` 只承接模板预填与解释延续，不重新承接候选选择主线。

自 `phase09-06` 起，系统必须把该 requirement 修改为：

- `Product Create` 页面只能消费：
  - `use-product-create-template-handoff` 这一条模板编排 owner
  - `useCreateDraftProduct` 这一条 canonical create mutation owner
  - `use-product-create-form-state` 这一条正式 create form state owner
- 表单字段状态继续只保留在这条单一 create form state 主线中

#### Scenario: 判断 Product Create 是否仍保持单一写路径与单一表单状态

- **WHEN** 后续实现或验收检查模板 create 场景
- **THEN** 必须同时满足：
  - 只有一条正式 create mutation owner
  - 只有一条正式 form state 主线
  - 模板 handoff 只编排上下文，不接管写动作
- **AND** 从 Module 补齐页返回时的草稿恢复只来自同一条 form state 主线
- **AND** 不得通过“页面里只是多写一点临时逻辑”绕过该边界

### Requirement: Weekly Review 的模板消费方式

`phase09-05` 已冻结 `Weekly Review` 是模板候选与提示的唯一主消费宿主。

自 `phase09-06` 起，系统必须把该 requirement 修改为：

- `Weekly Review` 的页面宿主仍只消费 `useWeeklyReviewRead`
- 模板候选与提示即使来自 `template-reuse` 切片，也只能在 `useWeeklyReviewRead` 内部被组合

#### Scenario: 判断 Weekly Review 是否长出第二套 query 主线

- **WHEN** 后续实现新增模板候选与提示消费
- **THEN** `WeeklyReviewPage` 不得直接 import `use-template-candidates-read`、`use-derived-insight-hints-read` 或等价底层 read hook
- **AND** 页面层只允许拿到 `useWeeklyReviewRead` 已派生好的单值结果

## REMOVED Requirements

### Requirement: 允许模板预填、模板来源复读或提示消费继续停留在页面级临时编排逻辑

**Reason**: 这会让 `Weekly Review / Product Create / Product Detail` 各自生长半套 owner，重新打散 `phase09-04` 已冻结的 caller-owner 边界，并直接诱发第二套 create form state 或第二套 detail 写路径。

**Migration**: 模板只读能力统一回收到 `template-reuse/data/`；`Product Create` 模板 handoff 回收到单一 application owner；`Product Detail` 模板来源复读回收到单一只读 owner；页面层只保留导航消费、toast 与组件编排。
