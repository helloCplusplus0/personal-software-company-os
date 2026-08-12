# phase09-09 落实模板候选选择、Product Create 预填与结果回流 Spec

## Why

`phase09-08` 已经将 `template_reuse.proto`、后端 `QueryService`、Connect handler 与前端四个 read owner 真实落地，并通过了工具链和 API smoke。但当前 `Weekly Review` 页面还不存在模板候选区，`Product Create` 页面还不认识 `fromTemplateReuse / templateCandidateId / templateSource` 参数，`Product Detail` 也不展示模板来源摘要。`phase09-09` 的职责就是把这些已落地的 read owner 正式接入页面，完成"从模板候选 → 预填创建 → 成功回流 → 来源摘要"的完整页面闭环。

## What Changes

- `Weekly Review` 页面新增模板候选选择区，包括候选列表、单选切换、CTA 与空态/失败态
- `Product Create` 页面新增 `fromTemplateReuse` 搜索参数接受、模板预填 handoff application owner、模板来源摘要展示与预填字段标记
- `Product Create` 新增单一 `use-product-create-form-state` owner，回收组件本地 `useState` 为正式 form state 主线
- `Product Create` 新增 `use-product-create-template-handoff` application owner，编排模板预填、返回路径与成功回流
- `Product Detail` 页面新增模板来源摘要区与 canonical binding CTA
- 路由新增 `fromTemplateReuse / templateCandidateId / templateSource` 搜索参数，并在 `templateSource=product-detail` 时补充 `templateSourceProductId`
- 实现 `templateSource` 驱动的返回链、非法参数回退与空候选回退

## Impact

- Affected specs:
  - `phase09_05_design_template_reuse_hint_page_flow_interaction_return_chain`
  - `phase09_06_design_frontend_read_write_owner_state_flow`
  - `phase09_08_land_template_reuse_contract_backend_frontend_read_enablement`
- Affected code:
  - `frontend/src/features/review/data/use-weekly-review-read.ts`（内部组合 template-reuse read hooks）
  - `frontend/src/features/review/pages/weekly-review-page.tsx`（新增模板候选区块）
  - `frontend/src/features/product-registry/application/use-product-create-template-handoff.ts`（新增）
  - `frontend/src/features/product-registry/application/use-product-create-form-state.ts`（新增）
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`（接入模板 handoff）
  - `frontend/src/features/product-registry/components/product-create-form.tsx`（接收预填初始值）
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`（新增模板来源摘要）
  - `frontend/src/routes/products/new.tsx`（新增模板搜索参数）
  - `frontend/src/routes/products/$productId.tsx`（新增模板搜索参数）

## ADDED Requirements

### Requirement: Weekly Review 必须新增模板候选选择区

系统 SHALL 在 `Weekly Review` 页面中上部新增模板候选选择区，承接 `phase09-08` 已落地的 `use-template-candidates-read` 与 `use-derived-insight-hints-read`，并通过 `useWeeklyReviewRead` 内部组合的方式保持页面只消费单一 read owner。

#### Scenario: Weekly Review 模板候选区的内部组合方式

- **WHEN** 实现 `Weekly Review` 的模板候选消费
- **THEN** `useWeeklyReviewRead` 必须在其内部组合 `use-template-candidates-read` 与 `use-derived-insight-hints-read`
- **AND** 必须导出以下新增字段：
  - `templateCandidates: TemplateCandidateSummary[]`
  - `defaultActiveCandidateId: string`
  - `activeCandidateId: string`（当前选中态）
  - `setActiveCandidateId(id: string): void`
  - `templateSectionStatus: 'ready' | 'empty' | 'error'`
  - `hints: DerivedInsightHint[]`
  - `hintsSectionStatus: 'ready' | 'empty' | 'error'`
- **AND** `WeeklyReviewPage` 不得直接 import `use-template-candidates-read` 或 `use-derived-insight-hints-read`

#### Scenario: 模板候选区的 UI 落点

- **WHEN** 用户进入 `Weekly Review`
- **THEN** 模板候选区必须位于 `Overview` 与 `Recent Activity` 之间的中上部主要动作区域
- **AND** 候选区标题为 `"模板候选"`，副标题为 `"基于现有复用事实的标准组合"` 或等价文案
- **AND** 候选区支持以下状态：
  - **ready**：展示候选卡片列表 + active candidate 高亮
  - **empty**：展示成功空态 `"当前没有可复用模板候选"`
  - **error**：展示局部失败态 + 重试按钮
- **AND** 候选区失败不得把整页 `Weekly Review` 回退为 page error

#### Scenario: 模板候选的默认选中与单选切换

- **WHEN** 模板候选列表加载成功
- **THEN** 排名第一的候选必须默认进入 active 状态
- **AND** 点击另一候选时，新候选成为唯一 active，旧候选立即退回非选中态
- **AND** active candidate 必须展示清晰的选中态（如 `ring-2 ring-primary` 或 `bg-accent`）
- **AND** 与候选关联的 `template_title`、`template_description`、模块列表与 `source_product_count` 必须跟随 active candidate 同步刷新

#### Scenario: 模板候选的 CTA

- **WHEN** 用户查看 active candidate
- **THEN** 候选区必须提供 `"以该模板创建产品"` CTA 按钮
- **AND** 点击 CTA 时必须导航到 `/products/new`，携带：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<active candidate 的 template_candidate_id>`
  - `templateSource=weekly-review`
  - 保留已有的 `fromDashboard / dashboardSection / dashboardReturnTo` 元数据
- **AND** 无候选时不展示 CTA 按钮

### Requirement: Product Create 必须新增模板 handoff application owner

系统 SHALL 新增 `use-product-create-template-handoff` 作为 `Product Create` 模板预填的单一 application owner，负责解析搜索参数、组合模板预填只读结果、拼装模板摘要 view model 与生成返回/回流参数。

#### Scenario: use-product-create-template-handoff 的最小落点

- **WHEN** 实现 `phase09-09`
- **THEN** 必须新增 `frontend/src/features/product-registry/application/use-product-create-template-handoff.ts`
- **AND** 该 owner 必须承接以下职责：
  - 解析 `fromTemplateReuse`, `templateCandidateId`, `templateSource` 搜索参数
  - 调用 `use-template-prefill-read(templateCandidateId, consumerSurface=PRODUCT_CREATE)`
  - 导出 `templateSummary` view model（包含模板标题、描述、模块列表、来源标记）
  - 导出 `prefillInitialValues`（`{ name: string; description: string }`）
  - 导出 `resolutionStatus`（`'resolved' | 'unavailable' | 'error'`）
  - 导出 `handleReturn()` 函数（基于 `templateSource` 决定返回路径）
  - 导出 `buildSuccessSearch()` 函数（创建成功后进 Detail 页的搜索参数）
  - 导出 `templateSourceLabel`（来源展示文案）
- **AND** 不得内联 `useMutation` 或 `createProduct` 调用
- **AND** 不得接管 `useCreateDraftProduct` 的写动作

#### Scenario: Product Create 的 form state 回收

- **WHEN** 实现 `Product Create` 的正式表单状态
- **THEN** 必须新增 `frontend/src/features/product-registry/application/use-product-create-form-state.ts`
- **AND** 该 owner 必须承接：
  - `name`, `description`, `status` 字段的单值状态
  - `setName`, `setDescription`, `setStatus` 更新函数
  - 接受初始值 `initialValues?: { name?: string; description?: string; status?: ProductStatus }`
  - 导出 `isDirty`（是否有未保存草稿）
  - 导出 `buildSubmitInput()`（组装 `CreateProductInput`）
- **AND** `ProductCreateForm` 必须改为接收 `name / description / status / onChange*` props，不再使用组件本地 `useState`
- **AND** 不得在 `ProductCreatePage` 或 `ProductCreateForm` 中继续保留并行 `useState` 草稿存储

#### Scenario: 预填字段的可见标记

- **WHEN** 预填值已生效
- **THEN** 表单中被预填的字段必须展示可见标记，使浏览器验收可以明确判断"预填已生效"
- **AND** 标记方式为：在字段 Label 旁显示 `"（来自模板）"` 或等价 `text-xs text-muted-foreground` 文案
- **AND** 用户编辑后标记仍可保留，因为模板来源信息不会因编辑而丢失

### Requirement: Product Create 页面必须接入模板 handoff 并保持单一 canonical 创建主线

系统 SHALL 修改 `ProductCreatePage` 以消费 `use-product-create-template-handoff` 和 `use-product-create-form-state`，同时保持 `useCreateDraftProduct` 为唯一 canonical create mutation owner。

#### Scenario: ProductCreatePage 的 owner 消费清单

- **WHEN** 审查 `ProductCreatePage` 的 import
- **THEN** 页面必须只消费以下 owner：
  - `useCreateDraftProduct`（canonical create mutation）
  - `use-product-create-template-handoff`（模板 handoff 编排）
  - `use-product-create-form-state`（正式 form state）
- **AND** 不得在页面内直接 import `use-template-prefill-read`、`TemplateReuseService` 或 `templateReuseClient`
- **AND** 不得在页面内直接拼装 `fromTemplateReuse` → `search` 的转换逻辑

#### Scenario: 模板来源摘要的展示

- **WHEN** 用户通过 `fromTemplateReuse=true` 进入 `/products/new`
- **THEN** 页面必须在表单上方展示模板来源摘要区
- **AND** 摘要区必须包含：
  - 模板标题
  - 模板来源标记（如 `"来源：Weekly Review"`）
  - 模块组合列表
- **AND** 摘要区使用紧凑样式（`border-t pt-2`），不作为独立重型卡片

#### Scenario: 模板 unavailable 成功态

- **WHEN** `templateCandidateId` 不可解析返回 `UNAVAILABLE`
- **THEN** 模板来源摘要区必须展示 `"模板来源已失效，但仍可继续创建"` 的可恢复提示
- **AND** 表单退化为空白但仍可编辑提交
- **AND** 页面不得因 unavailable 而阻断用户操作

#### Scenario: 模板预填请求失败态

- **WHEN** 模板预填读取发生网络/服务失败
- **THEN** 模板来源摘要区必须展示局部失败提示 + 重试按钮
- **AND** 表单仍可编辑提交
- **AND** 页面不得整页回退为 page error

#### Scenario: 取消返回路径

- **WHEN** 用户在 `Product Create` 主动取消或点击返回
- **THEN** 返回路径必须按 `templateSource` 决定：
  - `templateSource=weekly-review` → 返回 `Weekly Review`
  - `templateSource=dashboard` → 返回 `Dashboard`
  - `templateSource=product-detail` → 返回原 `Product Detail`
- **AND** 当 `templateSource=product-detail` 时，进入 `/products/new` 的 search 参数必须同时携带 `templateSourceProductId`
- **AND** 返回原 `Product Detail` 时必须使用该 `templateSourceProductId`，而不是模糊回退到 `/products`
- **AND** 若无模板来源（direct-entry），保持原有返回逻辑
- **AND** 不得统一退回浏览器历史或根路由

#### Scenario: 创建成功回流

- **WHEN** 用户通过模板预填成功创建 Product
- **THEN** 导航到 `/products/$productId` 时必须携带：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
  - 保留已有的 `fromDashboard / dashboardSection / dashboardReturnTo` 元数据
- **AND** 模板来源语义不得因为创建成功而丢失

### Requirement: Product Detail 必须新增模板来源摘要区

系统 SHALL 在 `ProductDetailPage` 中新增模板来源摘要区，消费 `phase09-08` 已落地的 `use-template-source-read`。

#### Scenario: 模板来源摘要的展示条件

- **WHEN** `/products/$productId` 的搜索参数包含 `fromTemplateReuse=true` 且 `templateCandidateId` 非空
- **THEN** 页面必须通过 `use-template-source-read` 读取模板来源摘要
- **AND** 读取条件为 `enabled = fromTemplateReuse && templateCandidateId !== ''`

#### Scenario: 模板来源摘要的 UI 落点

- **WHEN** 模板来源摘要读取成功
- **THEN** 摘要区必须位于 `ProductSummaryCard` 与 `ReuseSummaryInline` 之间
- **AND** 摘要区必须包含：
  - 模板标题
  - 模板来源标记
  - 模块组合列表
- **AND** 摘要区使用紧凑样式（`border-t pt-2`）

#### Scenario: 模板来源摘要的 canonical binding CTA

- **WHEN** 模板来源摘要展示成功
- **THEN** 摘要区必须同时提供 `"为模板模块绑定仓库"` 或等价 CTA
- **AND** CTA 行为为：滚动或聚焦到 `ProductModuleBindingPanel`
- **AND** 不得在摘要区内联自动绑定或第二套绑定写路径

#### Scenario: 模板来源摘要 unavailable

- **WHEN** `GetTemplateSourceSummary` 返回 `UNAVAILABLE`
- **THEN** 摘要区必须展示 `"模板来源已不可复读"` 的可恢复空态
- **AND** canonical binding CTA 仍必须可见

### Requirement: 路由搜索参数必须新增模板相关字段

系统 SHALL 在 `/products/new` 与 `/products/$productId` 路由中新增 `fromTemplateReuse`, `templateCandidateId`, `templateSource` 三个可选搜索参数。

#### Scenario: 路由搜索参数 schema 更新

- **WHEN** 实现 `phase09-09`
- **THEN** `/products/new` 的 `productCreateSearchSchema` 必须新增：
  - `fromTemplateReuse: z.boolean().optional()`
  - `templateCandidateId: z.string().optional()`
  - `templateSource: z.enum(['weekly-review', 'dashboard', 'product-detail']).optional()`
  - `templateSourceProductId: z.string().optional()`
- **AND** `/products/$productId` 的 `productDetailSearchSchema` 必须新增：
  - `fromTemplateReuse: z.boolean().optional()`
  - `templateCandidateId: z.string().optional()`
  - `templateSource: z.enum(['weekly-review', 'dashboard', 'product-detail']).optional()`
- **AND** `templateSourceProductId` 仅在 `templateSource=product-detail` 时有值，其余来源保持为空
- **AND** 其余字段均为可选，不影响 direct-entry 的既有行为

### Requirement: 非法参数与空候选必须回退为可恢复状态

系统 SHALL 确保非法模板参数组合、无效 `templateCandidateId` 与空候选场景不会把页面打成不可恢复错误。

#### Scenario: 非法参数组合回退

- **WHEN** `fromTemplateReuse=true` 但 `templateCandidateId` 为空
- **THEN** `ProductCreatePage` 必须回退为普通 direct-entry 创建
- **AND** 页面不得展示模板来源摘要区
- **AND** 页面不得进入错误状态

#### Scenario: 空候选回退（Weekly Review）

- **WHEN** `ListTemplateCandidates` 返回空列表
- **THEN** 模板候选区展示成功空态
- **AND** 页面其余区块继续正常可用
- **AND** 不得把整页回退为 page error

## MODIFIED Requirements

### Requirement: ProductCreatePage 的表单状态管理方式

`phase04-06` 已冻结 `ProductCreateForm` 以组件本地 `useState` 承接草稿字段。

自 `phase09-09` 起，系统必须把该 requirement 修改为：

- `ProductCreateForm` 改为接收 `name / description / status / onChangeName / onChangeDescription / onChangeStatus` props
- 表单字段状态统一由 `use-product-create-form-state` 这条正式 form state 主线承接
- `ProductCreatePage` 必须消费 `use-product-create-form-state` 并传入 `initialValues` 以支持模板预填

#### Scenario: 判断 ProductCreateForm 是否仍保持单一表单状态

- **WHEN** 审查 `ProductCreateForm` 组件
- **THEN** 不得再包含组件本地 `useState('')`、`useState<ProductStatus>('active')` 等草稿字段声明
- **AND** 所有字段值与更新函数必须来自 props

### Requirement: ProductCreatePage 的来源上下文判定

`phase04-06` 已冻结 `ProductCreatePage` 的来源上下文为 `fromList / fromModuleDetail / direct-entry` 三种。

自 `phase09-09` 起，系统必须把该 requirement 修改为：

- 新增 `fromTemplateReuse` 作为第四种来源上下文
- `fromTemplateReuse` 优先级高于 `fromList`，但低于 `fromDashboard` 的返回链语义
- 取消返回时按 `templateSource` 决定目的地

#### Scenario: 判断来源上下文优先级

- **WHEN** 搜索参数同时存在 `fromTemplateReuse=true` 和 `fromList=true`
- **THEN** 页面主来源语义必须显示为模板创建
- **AND** 取消返回必须按 `templateSource` 决定，而不是按 `fromList` 回列表

## REMOVED Requirements

### Requirement: 允许 ProductCreateForm 继续以组件本地 useState 单独保存正式草稿字段

**Reason**: 这会直接破坏 `phase09-06` 已冻结的"单一正式 create form state 主线"约束，并让模板预填只能通过 hack 方式注入（如 `useEffect` 后 `setName`），无法形成可追溯的正式状态流。

**Migration**: 表单字段状态统一回收至 `use-product-create-form-state`；`ProductCreateForm` 改为接收 props；`ProductCreatePage` 统一消费 form state owner 并传入预填初始值。
