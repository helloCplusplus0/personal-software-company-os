# phase09-04 冻结本阶段合同、读模型、owner 与候选数据源边界 Spec

## Why

`phase09-01 ~ 03` 已经把范围、模板 handoff 与提示矩阵冻结为单值口径，但如果不继续把 `.proto` 合同归属、前后端 owner、真实 caller inventory，以及模板候选的数据源边界正式冻结，后续实现仍会在 `review / create / detail` 三段链路里各自猜一套。
因此，`phase09-04` 必须把“谁定义合同、谁只读组合、谁承接写动作、候选到底是否允许快照化”一次性拍板，作为 `phase09-05+` 的直接上游。

## What Changes

- 冻结 `phase09` 的最小正式合同归属：`review.proto`、`reuse_summary.proto`、`product_registry.proto` 与新增 `template_reuse.proto` 的职责边界
- 冻结模板候选、组合快照、派生提示与 create 预填的最小正式字段边界
- 冻结 `phase06 reuse_summary`、`phase08 review`、`phase04 product create` 的正式消费边界
- 冻结前端 `read owner / application owner` 与后端 `query owner / command owner` 的承接矩阵
- 冻结当前真实 route / page / query owner / application owner inventory
- 冻结模板候选只允许基于 `product_modules` 读时派生，当前阶段不允许引入轻量快照持久化
- 冻结模板来源参数在 `Product Create -> Product Detail` 成功回流链中的保留方式

## Impact

- Affected specs:
  - `phase09_01_freeze_template_reuse_derived_intelligence_scope_success_non_goals`
  - `phase09_02_freeze_template_reuse_assets_candidates_product_create_prefill_chain`
  - `phase09_03_freeze_derived_insight_hint_set_explanation_cta_handoff`
  - `phase08_04_freeze_review_contract_read_record_boundary`
  - `phase06_04_define_reuse_summary_read_model_page_attachment`
  - `phase06_07_design_frontend_write_path_mutation_owners`
  - `phase07_05_design_frontend_generated_client_query_application_migration`
- Affected code:
  - `proto/psco/review/v1/review.proto`
  - `proto/psco/reuse_summary/v1/reuse_summary.proto`
  - `proto/psco/product_registry/v1/product_registry.proto`
  - 新增 `proto/psco/template_reuse/v1/template_reuse.proto`
  - `backend/internal/review/service/query_service.go`
  - `backend/internal/reusesummary/service/query_service.go`
  - `backend/internal/productregistry/service/command_service.go`
  - 后续新增 `backend/internal/templatereuse/*`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`
  - `frontend/src/features/product-registry/application/use-create-draft-product.ts`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - 后续新增 `frontend/src/features/template-reuse/*`

## ADDED Requirements

### Requirement: `phase09` 的正式合同必须按 `.proto` 包职责单值归属

系统 SHALL 继续以 `.proto` 作为唯一长期合同源，并把 `phase09` 的新增正式消息与 RPC 归属冻结为“既有 canonical 合同继续承接既有事实；新增模板复用读合同独立承接 phase09 支撑读能力”。

#### Scenario: 判断既有 `.proto` 包是否保持 canonical 职责

- **WHEN** 后续 `/spec`、实现或验收需要定义 `phase09` 的正式合同
- **THEN** `psco.review.v1` 只允许继续承接 `Weekly Review / Daily Review` 页面级组合上下文
- **AND** `psco.reuse_summary.v1` 只允许继续承接 `module_reuse_summary / capability_summary`
- **AND** `psco.product_registry.v1` 只允许继续承接 `Product Create / Product Detail / Product <-> Module Binding` 的 canonical 写读合同
- **AND** 不得把模板候选、模板预填或模板来源复读直接塞进 `reuse_summary.proto` 或 `product_registry.proto` 作为并列事实主线

#### Scenario: 判断 `phase09` 新增读合同的包归属

- **WHEN** 后续需要为模板候选、组合快照、派生提示与 create 预填补齐正式读合同
- **THEN** 必须新增单一读包 `psco.template_reuse.v1`
- **AND** `template_reuse.proto` 必须成为以下消息的唯一定义源：
  - `TemplateCandidateSummary`
  - `TemplateModuleRef`
  - `TemplateCandidatePrefill`
  - `DerivedInsightHint`
- **AND** `review.proto` 只允许 import 并消费这些消息，不得重新定义一套等价字段

#### Scenario: 判断模板读合同的正式传输 owner

- **WHEN** 后续需要为模板候选列表、模板预填与模板来源复读补齐正式 RPC
- **THEN** 必须由 `TemplateReuseService` 作为模板读能力的唯一 canonical transport owner
- **AND** `TemplateReuseService` 至少必须承接：
  - `ListTemplateCandidates`
  - `GetTemplateCandidatePrefill`
  - `GetTemplateSourceSummary`
- **AND** `review.proto` 只允许通过 `WeeklyReviewContext` import 并组合模板相关消息，不得新增 `GetWeeklyTemplateCandidates`、`GetWeeklyDerivedHints` 一类并列模板 RPC
- **AND** `ReviewService` 在 `Weekly Review` 场景中的身份只允许是页面级组合 transport owner，而不是模板读能力的第二 canonical service

### Requirement: 模板候选、组合快照、派生提示与 create 预填必须冻结最小正式字段边界

系统 SHALL 在 `template_reuse.proto` 中冻结 `phase09` 最小正式字段边界，确保前后端不会分别拼出不同的模板与提示语义。

#### Scenario: 判断 `TemplateCandidateSummary` 的最小字段

- **WHEN** 后续定义模板候选列表的正式合同
- **THEN** `TemplateCandidateSummary` 至少必须承接：
  - `template_candidate_id`
  - `template_title`
  - `template_description`
  - `modules[]`
  - `source_product_count`
  - `total_reuse_product_count`
  - `latest_source_product_updated_at`
- **AND** `modules[]` 必须引用 `TemplateModuleRef`
- **AND** 不得把完整 Product 实体快照、页面本地排序态或 UI 展示字段直接编码进候选合同

#### Scenario: 判断 `TemplateModuleRef` 的最小字段

- **WHEN** 后续定义模板候选内的模块组合摘要
- **THEN** `TemplateModuleRef` 至少必须承接：
  - `module_id`
  - `module_name`
  - `capability_key`
  - `capability_label`
- **AND** 不得把模块完整详情、仓库映射或第二套模块状态语义复制进模板组合摘要

#### Scenario: 判断 `TemplateCandidatePrefill` 的最小字段

- **WHEN** 后续定义 `Product Create` 的模板预填正式合同
- **THEN** `TemplateCandidatePrefill` 至少必须承接：
  - `template_candidate_id`
  - `resolution_status`
  - `unavailable_reason_text`
  - `template_title`
  - `template_description`
  - `suggested_product_name`
  - `suggested_product_description`
  - `modules[]`
  - `capability_gap_hints[]`
- **AND** `capability_gap_hints[]` 只允许包含 `capability_gap_hint`
- **AND** 不得把 `CreateProductRequest` 直接扩写为内嵌模板 payload 的写合同

#### Scenario: 判断 `TemplateSourceSummary` 的最小字段

- **WHEN** 后续定义 `Product Detail` 的模板来源复读正式合同
- **THEN** `TemplateSourceSummary` 至少必须承接：
  - `template_candidate_id`
  - `resolution_status`
  - `unavailable_reason_text`
  - `template_title`
  - `template_description`
  - `modules[]`
  - `template_source`
- **AND** 不得把 `Product Detail` 对模板来源的复读偷渡为 `product_registry.proto` 内联字段

#### Scenario: 判断 `DerivedInsightHint` 的最小字段

- **WHEN** 后续定义正式派生提示合同
- **THEN** `DerivedInsightHint` 至少必须承接：
  - `hint_type`
  - `title`
  - `explanation_text`
  - `cta_kind`
  - `template_candidate_id`
  - 可选的 `capability_key`
  - 可选的 `module_id`
- **AND** `hint_type` 只允许取值：
  - `REUSE_OPPORTUNITY`
  - `CAPABILITY_GAP`
- **AND** `cta_kind` 只允许表达既有 canonical 动作，不得返回任意页面 URL 字符串

### Requirement: 模板候选必须保持读时派生，当前阶段不允许轻量快照持久化

系统 SHALL 在 `phase09-04` 最终拍板：模板候选、组合快照与 create 预填上下文当前阶段必须全部基于 `product_modules` 已持久化事实读时派生，不允许新增独立模板快照表或等价持久化记录。

#### Scenario: 判断模板候选的数据源与派生方式

- **WHEN** 后续实现模板候选、候选详情或预填读取
- **THEN** canonical 来源必须是 `product_modules` 已持久化绑定事实
- **AND** 去重键必须继续是去重并升序排序后的 `module_id` 集合
- **AND** 候选排序必须继续遵守 `phase09-02` 已冻结顺序
- **AND** `templateCandidateId` 必须由后端根据 normalized module-set key 单向派生

#### Scenario: 判断是否允许引入轻量快照记录

- **WHEN** 后续方案试图新增模板快照表、候选缓存表、预填持久化记录或等价中间事实表
- **THEN** 当前阶段必须判定为不允许
- **AND** 不得保留“实现时再决定是否落表”的灰区
- **AND** 只有新的独立 `/spec` 明确推翻本规格后，未来阶段才允许重新仲裁

#### Scenario: 判断 `templateCandidateId` 在读时派生下的漂移语义

- **WHEN** `templateCandidateId` 因底层 `product_modules` 事实变化而无法再被当前系统重新派生
- **THEN** 模板相关读取必须返回**可恢复 unavailable 成功态**
- **AND** 不得把这类事实漂移解释为页面级 fatal error
- **AND** 不得静默把失效模板重新映射为另一个候选

#### Scenario: 判断 Product Create 中失效 `templateCandidateId` 的行为

- **WHEN** `/products/new` 以 `fromTemplateReuse=true` 进入，但 `templateCandidateId` 已无法重新解析
- **THEN** `GetTemplateCandidatePrefill` 必须返回：
  - `resolution_status = UNAVAILABLE`
  - `unavailable_reason_text`
  - 空的 `modules[]`
  - 空的 `capability_gap_hints[]`
- **AND** `Product Create` 必须退化为可继续编辑的空白 create 表单
- **AND** 页面必须展示“模板来源已失效但当前仍可继续创建”的可恢复提示

#### Scenario: 判断 Product Detail 中失效 `templateCandidateId` 的行为

- **WHEN** `/products/$productId` 成功回流后，需要基于 `templateCandidateId` 复读模板来源，但该候选已无法重新解析
- **THEN** `GetTemplateSourceSummary` 必须返回：
  - `resolution_status = UNAVAILABLE`
  - `unavailable_reason_text`
  - 空的 `modules[]`
- **AND** `Product Detail` 必须展示模板来源不可复读的可恢复空态
- **AND** 页面仍必须继续提供 canonical `Product <-> Module Binding` CTA

### Requirement: `phase06 reuse_summary`、`phase08 review` 与 `phase04 product create` 的正式消费边界必须单值

系统 SHALL 冻结 `phase09` 对三条既有主线的消费边界，避免 phase09 把既有能力改写成自己的并列 owner。

#### Scenario: 判断 `phase06 reuse_summary` 的正式消费边界

- **WHEN** `phase09` 需要消费 `module_reuse_summary / capability_summary`
- **THEN** 只能通过既有 `ReuseSummaryService` / `useReuseSummaryRead` 或其后端 query owner 消费
- **AND** `ReuseSummaryService` 只继续负责摘要事实，不承接模板候选、模板预填与提示编排
- **AND** 不得把 `phase09` 的模板语义直接写回 `reuse_summary.proto`

#### Scenario: 判断 `phase08 review` 的正式消费边界

- **WHEN** `Weekly Review` 需要显示模板候选或提示
- **THEN** `ReviewService` 只允许继续承担页面级组合 owner
- **AND** 它可以组合 `Dashboard / ReuseSummary / TemplateReuse` 的只读结果
- **AND** 不得在 `review.QueryService` 中直接写跨模块 SQL 或复制模板候选 derivation 逻辑
- **AND** 不得把 `ReviewService` 扩写成模板读能力的第二 canonical transport owner

#### Scenario: 判断 `phase04 product create` 的正式消费边界

- **WHEN** `Product Create` 通过模板进入预填
- **THEN** `CreateProductRequest / Response` 必须保持 canonical create 写合同不变
- **AND** `useCreateDraftProduct` 必须继续是前端唯一正式 create mutation owner
- **AND** 不得在 create mutation 内自动写入 `product_modules`
- **AND** 模板相关逻辑只允许作为预填读取与成功回流上下文存在

### Requirement: 后端 query owner / command owner 承接矩阵必须单值

系统 SHALL 冻结 `phase09` 后端 owner 矩阵，确保模板与提示不会复制第二套业务主线。

#### Scenario: 判断后端 query owner 的正式分工

- **WHEN** 后续实现 `phase09` 后端读能力
- **THEN** 必须保持以下单值分工：
  - `review.QueryService`：承接 `Weekly Review` 页面级组合读取
  - `reusesummary.QueryService`：承接 `module_reuse_summary / capability_summary`
  - `templatereuse.QueryService`：承接模板候选派生、`templateCandidateId` 解析、create 预填读取、模板相关提示读取
  - `productregistry.QueryService`：承接 `Product Detail` 与 canonical binding reread
- **AND** `templatereuse.QueryService` 必须通过 candidate/reader 子包读取 canonical 表
- **AND** 不得把 `templatereuse.QueryService` 退化为 `review.QueryService` 内的一组私有 helper

#### Scenario: 判断后端 command owner 的正式分工

- **WHEN** 后续实现 `phase09` 写路径
- **THEN** 必须保持以下单值分工：
  - `review.CommandService` 只承接 review result sink
  - `productregistry.CommandService` 只承接 create product 与 canonical binding
- **AND** 当前阶段不得新增 `TemplateReuseCommandService`
- **AND** 不得新增 `CreateProductFromTemplate`、`ApplyTemplateToProduct` 或等价第二套业务写 RPC

### Requirement: 前端 read owner / application owner 承接矩阵必须单值

系统 SHALL 冻结 `phase09` 前端 owner 矩阵，避免 route、page 与表单组件各自长出第二套 query 或 mutation。

#### Scenario: 判断 Weekly Review 的前端 owner 分工

- **WHEN** `/reviews/weekly` 消费 `phase09` 能力
- **THEN** 页面层必须继续只消费：
  - `useWeeklyReviewRead`
  - `useReviewAction`
- **AND** `useWeeklyReviewRead` 可以在切片内部组合模板候选与提示读取
- **AND** `WeeklyReviewPage` 不得直接新增第二个 page-level `useQuery` 或 `useMutation`

#### Scenario: 判断 Product Create 的前端 owner 分工

- **WHEN** `/products/new` 通过模板来源进入
- **THEN** 页面层必须只新增模板相关只读 owner
- **AND** `useCreateDraftProduct` 必须继续是唯一正式写路径 owner
- **AND** 模板预填正式读取必须通过单一 `template reuse` read owner 承接
- **AND** `ProductCreatePage` 与 `ProductCreateForm` 不得直接持有 `queryClient` 或模板相关写动作

#### Scenario: 判断 Product Detail 的前端 owner 分工

- **WHEN** `/products/$productId` 需要复读模板来源摘要
- **THEN** 页面层只允许新增只读来源摘要消费位
- **AND** 模板来源复读必须通过单一 `template reuse` read owner 承接
- **AND** canonical 下一步动作仍然只能导向既有 `ProductModuleBindingPanel`
- **AND** 不得在 `ProductDetailPage` 内长出模板应用写路径

### Requirement: 当前真实 caller / route / owner inventory 必须成为后续设计强制输入

系统 SHALL 把当前真实调用面冻结为 `phase09-05+` 的强制输入，避免后续设计跳过真实入口直接抽象出第二套宿主。

#### Scenario: 当前真实 caller inventory

- **WHEN** 后续 `/spec` 或实现设计列出 `phase09` 消费面
- **THEN** 至少必须显式覆盖以下当前真实 inventory：
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
  - `frontend/src/routes/reviews/weekly.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/routes/products/new.tsx`
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`
  - `frontend/src/features/product-registry/components/product-create-form.tsx`
  - `frontend/src/features/product-registry/application/use-create-draft-product.ts`
  - `frontend/src/routes/products/$productId.tsx`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/reuse-summary/data/use-reuse-summary-read.ts`
- **AND** 后续设计必须说明哪些文件被扩展、哪些只被消费、哪些明确禁止升级为新的写 owner

### Requirement: 模板来源参数必须在 `Product Create -> Product Detail` 成功回流链中保持单值

系统 SHALL 冻结模板来源参数在成功回流链中的保留方式，确保模板语义不依赖页面局部状态或一次性内存。

#### Scenario: 判断 Product Create 的模板来源读取入口

- **WHEN** `/products/new` 从模板候选进入
- **THEN** 正式 search 参数必须继续只使用：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=weekly-review | dashboard | product-detail`
- **AND** `templateCandidateId` 必须成为模板预填与 create 场景下 `capability_gap_hint` 的唯一读取入口

#### Scenario: 判断创建成功后的模板来源复读链

- **WHEN** 用户通过模板预填创建 Product 成功
- **THEN** 跳转到 `/products/$productId` 时必须继续携带：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
- **AND** `Product Detail` 只能通过这些参数重新读取模板来源摘要
- **AND** 不得依赖 localStorage、session state 或页面内存保留模板上下文

## MODIFIED Requirements

### Requirement: `Weekly Review` 的 phase09 消费方式

`phase09-02 / 03` 已冻结 `Weekly Review` 需要消费模板候选与提示。

自 `phase09-04` 起，系统必须把该 requirement 修改为：

- `Weekly Review` 的页面宿主继续保持单一 `useWeeklyReviewRead`
- 模板候选与提示的消息定义归属 `template_reuse.proto`
- `review.proto` 只承担周报页面级组合合同，不拥有模板候选与提示的原生消息定义
- `ReviewService` 只承担周报页面组合读取；模板读能力的 canonical transport owner 仍然是 `TemplateReuseService`

#### Scenario: 判断 Weekly Review 是否仍然是单一页面宿主

- **WHEN** 后续实现 `Weekly Review` 的模板候选与提示展示
- **THEN** 页面不得直接并排消费 `dashboard / reuse_summary / template_reuse` 多个 page-level query owner
- **AND** 必须继续以 `useWeeklyReviewRead` 作为页面级只读入口

### Requirement: `Product Create` 的模板承接解释

`phase09-02` 已冻结 `Product Create` 是模板预填的唯一动作入口。

自 `phase09-04` 起，系统必须把该 requirement 修改为：

- `Product Create` 只新增模板相关只读 owner
- 正式写动作继续只通过 `useCreateDraftProduct`
- 模板上下文只能经由 `templateCandidateId` 读取与成功回流链延续
- `templateCandidateId` 无法重解析时，create 页必须退化为可恢复 unavailable 成功态，而不是整页错误

#### Scenario: 判断 create 页是否长出第二套写主线

- **WHEN** 后续实现需要在 create 页承接模板
- **THEN** 只能新增 read owner 与展示/解释承接位
- **AND** 不得把模板确认、模板应用或自动绑定实现成新的 mutation owner

## REMOVED Requirements

### Requirement: 允许 `phase09` 同时保留“读时派生”与“轻量快照记录”两套长期稳态

**Reason**: 当前项目规则与 `phase09` 上游规格都已明确要求模板候选不得建立独立数据库表，且 `phase09-04` 必须在进入实现设计前拍板唯一稳态。继续保留两套候选稳态只会把合同、owner 与验收口径重新拉回灰区。

**Migration**: 后续实现必须统一按“`product_modules` 读时派生 + `templateCandidateId` 读时解析”推进；若未来确需引入快照化方案，必须通过新的独立 `/spec` 显式替换本 requirement，而不是在当前阶段实现中偷偷并存。
