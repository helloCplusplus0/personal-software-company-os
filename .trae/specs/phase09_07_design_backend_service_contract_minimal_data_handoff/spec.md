# phase09-07 产出后端服务、合同与最小数据承接设计 Spec

## Why

`phase09-04` 已冻结模板读合同、owner 与候选数据源边界，`phase09-05 / 06` 又分别冻结了页面流、前端 read/application owner 与状态流。当前仍缺一份直接面向后端实现的规格，来回答 `template_reuse.proto` 应该长什么样、`TemplateReuseService` 如何与既有 `Review / Product / Decision / ReuseSummary` 服务协作、三类核心读取接口与成功回流来源复读各由谁承接，以及为什么当前阶段不需要轻量快照记录。否则后续实现很容易在 proto、Connect service、query owner 与 API smoke 上各写一套。

## What Changes

- 冻结 `template_reuse.proto` 的最小消息、枚举与 Connect RPC 矩阵
- 冻结 `TemplateReuseService`、`templatereuse.QueryService` 与既有 `Review / Product / Decision / ReuseSummary` 服务的协作边界
- 冻结模板候选读取、模板预填详情读取、派生提示读取三类核心接口，以及创建成功后模板来源复读接口的正式责任边界
- 冻结后端 query owner、必要 command owner 与结果回流设计
- 冻结“默认关闭轻量快照记录”的证据链与受控边界
- 冻结 `buf / go build / frontend type-check / browser acceptance` 的关键路径验收口径

## Impact

- Affected specs:
  - `phase09_04_freeze_contract_read_model_owner_candidate_source_boundary`
  - `phase09_05_design_template_reuse_hint_page_flow_interaction_return_chain`
  - `phase09_06_design_frontend_read_write_owner_state_flow`
  - `phase08_07_design_review_backend_contract_service_data_handoff`
  - `phase07_04_design_go_connect_handler_service_chi_mount`
  - `phase06_05_freeze_contract_transport_source_constraints`
  - `phase06_10_design_minimal_protocol_buffers_contract`
- Affected code:
  - 新增 `proto/psco/template_reuse/v1/template_reuse.proto`
  - `proto/Makefile`
  - 后续新增 `backend/internal/templatereuse/`
  - 后续新增 `backend/internal/templatereuse/service/query_service.go`
  - 后续新增 `backend/internal/templatereuse/candidate/`
  - 后续新增 `backend/internal/templatereuse/connect/server.go`
  - `backend/internal/review/service/query_service.go`
  - `backend/internal/reusesummary/service/query_service.go`
  - `backend/internal/productregistry/service/query_service.go`
  - `backend/internal/productregistry/service/command_service.go`
  - `backend/internal/platform/router.go`
  - 后续新增 `frontend/src/features/template-reuse/data/*`
  - 后续新增 `frontend/src/features/template-reuse/application/*`

## ADDED Requirements

### Requirement: `phase09` 后端正式合同必须收敛到单一 `psco.template_reuse.v1` 包与 `TemplateReuseService`

系统 SHALL 将 `phase09` 的模板候选、模板预填、模板来源复读与派生提示合同冻结为单一 `proto/psco/template_reuse/v1/template_reuse.proto`，并由 `TemplateReuseService` 作为唯一 canonical transport owner。

#### Scenario: `template_reuse.proto` 的最小 RPC 矩阵

- **WHEN** 后续实现 `phase09` 后端合同
- **THEN** `TemplateReuseService` 至少必须包含以下 RPC：
  - `ListTemplateCandidates`
  - `GetTemplateCandidatePrefill`
  - `GetDerivedInsightHints`
  - `GetTemplateSourceSummary`
- **AND** 其中前三者构成当前阶段的三类核心读取接口：
  - 模板候选读取
  - 模板预填详情读取
  - 派生提示读取
- **AND** `GetTemplateSourceSummary` 必须作为创建成功后模板来源复读链的唯一正式读取接口
- **AND** 不得在 `ReviewService`、`ReuseSummaryService` 或 `ProductRegistryService` 内再长出并列模板 RPC

#### Scenario: 既有 `.proto` 包保持 canonical 职责

- **WHEN** 后续实现 `phase09` 合同
- **THEN** 必须保持以下单值归属：
  - `psco.review.v1`：只承接 `Weekly Review / Daily Review` 页面级组合上下文
  - `psco.reuse_summary.v1`：只承接 `module_reuse_summary / capability_summary`
  - `psco.product_registry.v1`：只承接 canonical `Product Create / Product Detail / Product <-> Module Binding`
  - `psco.template_reuse.v1`：只承接模板候选、模板预填、模板来源复读与派生提示读取
- **AND** 不得把模板候选与提示重新定义进 `review.proto` 或把模板来源复读偷渡到 `product_registry.proto`

### Requirement: `template_reuse.proto` 的消息与枚举必须足以单值表达模板候选、active candidate、提示与来源复读语义

系统 SHALL 在 `template_reuse.proto` 中冻结最小消息、枚举与请求/响应字段边界，确保前后端不会分别拼装不同的 phase09 语义。

#### Scenario: 四个读取 RPC 的最小 request 合同

- **WHEN** 后续定义 `TemplateReuseService` 的 request message
- **THEN** 至少必须冻结以下最小请求边界：
  - `ListTemplateCandidatesRequest`：`consumer_surface`
  - `GetTemplateCandidatePrefillRequest`：`template_candidate_id`、`consumer_surface`
  - `GetDerivedInsightHintsRequest`：`template_candidate_id`、`consumer_surface`
  - `GetTemplateSourceSummaryRequest`：`template_candidate_id`、`template_source`、`consumer_surface`
- **AND** `consumer_surface` 必须使用 `TemplateConsumerSurface` 枚举，而不是页面自定义字符串
- **AND** request message 不得退化为“空请求 + 服务端自行猜当前页面语义”

#### Scenario: 模板候选读取合同的最小字段

- **WHEN** 后续定义 `ListTemplateCandidates`
- **THEN** `ListTemplateCandidatesRequest` 至少必须承接：
  - `consumer_surface`
- **THEN** `ListTemplateCandidatesResponse` 至少必须承接：
  - `candidates[]`
  - `default_active_candidate_id`
- **AND** `TemplateCandidateSummary` 至少必须承接：
  - `template_candidate_id`
  - `template_title`
  - `template_description`
  - `modules[]`
  - `source_product_count`
  - `total_reuse_product_count`
  - `latest_source_product_updated_at`
- **AND** 当 `candidates[]` 非空时，`default_active_candidate_id` 必须等于排序第一项的 `template_candidate_id`
- **AND** 当 `candidates[]` 为空时，`default_active_candidate_id` 必须为空字符串
- **AND** 当前阶段 `ListTemplateCandidatesRequest.consumer_surface` 只允许正式取值 `WEEKLY_REVIEW`

#### Scenario: 模板预填详情合同的最小字段

- **WHEN** 后续定义 `GetTemplateCandidatePrefill`
- **THEN** `GetTemplateCandidatePrefillRequest` 至少必须承接：
  - `template_candidate_id`
  - `consumer_surface`
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
- **AND** `capability_gap_hints[]` 只允许包含 `CAPABILITY_GAP`
- **AND** 不得把 `CreateProductRequest` 直接扩写成内嵌模板 payload 的第二套写合同
- **AND** 当前阶段 `GetTemplateCandidatePrefillRequest.consumer_surface` 只允许正式取值 `PRODUCT_CREATE`

#### Scenario: 派生提示合同的最小字段

- **WHEN** 后续定义 `GetDerivedInsightHints`
- **THEN** `GetDerivedInsightHintsRequest` 至少必须承接：
  - `template_candidate_id`
  - `consumer_surface`
  - `review_scope_key`
- **AND** `review_scope_key` 在 `consumer_surface = WEEKLY_REVIEW` 时必须为必填
- **AND** `review_scope_key` 在 `consumer_surface = PRODUCT_CREATE` 时必须为空字符串
- **AND** `review_scope_key` 必须来自 canonical review context 的正式作用域标识，不得来自页面本地临时拼装
- **THEN** `GetDerivedInsightHintsResponse` 至少必须承接：
  - `resolution_status`
  - `unavailable_reason_text`
  - `hints[]`
- **THEN** `DerivedInsightHint` 至少必须承接：
  - `hint_type`
  - `title`
  - `explanation_text`
  - `cta_kind`
  - `template_candidate_id`
  - 可选的 `capability_key`
  - 可选的 `module_id`
- **AND** `hint_type` 只允许：
  - `REUSE_OPPORTUNITY`
  - `CAPABILITY_GAP`
- **AND** `cta_kind` 只允许表达既有 canonical 动作，不得返回任意页面 URL
- **AND** `GetDerivedInsightHintsResponse.resolution_status = RESOLVED` 且 `hints[]` 为空时，才表示“当前成功但无可展示提示”

#### Scenario: 模板来源复读合同的最小字段

- **WHEN** 后续定义 `GetTemplateSourceSummary`
- **THEN** `GetTemplateSourceSummaryRequest` 至少必须承接：
  - `template_candidate_id`
  - `template_source`
  - `consumer_surface`
- **THEN** `TemplateSourceSummary` 至少必须承接：
  - `template_candidate_id`
  - `resolution_status`
  - `unavailable_reason_text`
  - `template_title`
  - `template_description`
  - `modules[]`
  - `template_source`
- **AND** 不得把该来源复读结果退化为 `ProductDetail` 内联字段
- **AND** 当前阶段 `GetTemplateSourceSummaryRequest.consumer_surface` 只允许正式取值 `PRODUCT_DETAIL`

#### Scenario: 关键枚举的最小边界

- **WHEN** 后续定义 `template_reuse.proto` 枚举
- **THEN** 至少必须冻结：
  - `TemplateResolutionStatus`：`RESOLVED / UNAVAILABLE`
  - `DerivedInsightHintType`：`REUSE_OPPORTUNITY / CAPABILITY_GAP`
  - `TemplateSource`：`WEEKLY_REVIEW / DASHBOARD / PRODUCT_DETAIL`
  - `TemplateConsumerSurface`：至少覆盖 `WEEKLY_REVIEW / PRODUCT_CREATE / PRODUCT_DETAIL`
- **AND** 不得把这些长期语义退回 string 常量

### Requirement: 三类核心读取接口与来源复读接口的责任边界必须单值化

系统 SHALL 把 `TemplateReuseService` 的四个读取接口职责冻结清楚，确保 API smoke 与前端 owner 实现不会重复消费或串位。

#### Scenario: `ListTemplateCandidates` 的正式职责

- **WHEN** 后续实现模板候选读取
- **THEN** `ListTemplateCandidates` 只允许负责：
  - 基于 `product_modules` 已持久化事实派生候选集合
  - 应用已冻结的去重键与排序规则
  - 返回排序后的候选列表与 `default_active_candidate_id`
- **AND** 它不得负责：
  - Product Create 预填详情
  - 派生提示计算
  - 创建成功后的模板来源复读

#### Scenario: `GetTemplateCandidatePrefill` 的正式职责

- **WHEN** 后续实现模板预填详情读取
- **THEN** `GetTemplateCandidatePrefill` 只允许负责：
  - 解析 `template_candidate_id`
  - 组装 `Product Create` 预填字段
  - 返回 create 场景下的 `capability_gap_hints[]`
  - 返回 `RESOLVED / UNAVAILABLE` 两种解析状态
- **AND** 它不得负责候选列表排序、active candidate 选择或 `Product Detail` 来源复读

#### Scenario: `GetDerivedInsightHints` 的正式职责

- **WHEN** 后续实现派生提示读取
- **THEN** `GetDerivedInsightHints` 只允许负责：
  - 围绕传入的 `template_candidate_id` 计算当前消费面的正式提示
  - 围绕 `consumer_surface + review_scope_key` 解释当前提示计算作用域
  - 在 `Weekly Review` 场景返回 `reuse_opportunity_hint` 与可选 `capability_gap_hint`
  - 在无可用提示时返回成功空列表
- **AND** 它不得负责返回模板预填字段或来源复读摘要
- **AND** `Product Create` 场景下所需 `capability_gap_hint` 由 `GetTemplateCandidatePrefill` 承接，不复制为第二次提示读取

#### Scenario: `GetDerivedInsightHints` 的 request 作用域语义

- **WHEN** `GetDerivedInsightHints` 在 `Weekly Review` 场景被调用
- **THEN** request 必须同时携带：
  - `template_candidate_id`
  - `consumer_surface = WEEKLY_REVIEW`
  - `review_scope_key`
- **AND** `review_scope_key` 必须来自 `ReviewService` 已冻结的 canonical review 作用域，而不是页面临时状态
- **WHEN** `GetDerivedInsightHints` 在 `Product Create` 场景被调用
- **THEN** 当前阶段必须判定为不需要单独调用该 RPC
- **AND** create 场景下的 `capability_gap_hint` 继续由 `GetTemplateCandidatePrefill` 内联返回

#### Scenario: 区分“无提示成功空态”和“候选已漂移 unavailable 成功态”

- **WHEN** `GetDerivedInsightHints` 成功解析了 `template_candidate_id`，但当前没有可展示提示
- **THEN** 必须返回：
  - `resolution_status = RESOLVED`
  - 空的 `hints[]`
  - 空的 `unavailable_reason_text`
- **WHEN** `template_candidate_id` 因读时派生漂移而无法再被当前系统重新解析
- **THEN** 必须返回：
  - `resolution_status = UNAVAILABLE`
  - `unavailable_reason_text`
  - 空的 `hints[]`
- **AND** 不得把“候选已失效”静默退化成普通空提示成功态

#### Scenario: `GetTemplateSourceSummary` 的正式职责

- **WHEN** 后续实现创建成功后的模板来源复读
- **THEN** `GetTemplateSourceSummary` 只允许负责：
  - 基于成功回流链中的 `templateCandidateId + templateSource`
  - 为 `Product Detail` 返回模板来源摘要或 unavailable 成功态
- **AND** 它不得退化为候选列表或预填详情读取的别名接口

### Requirement: `templatereuse.QueryService` 必须成为模板读能力的单一后端 owner

系统 SHALL 将 `phase09` 后端模板读能力冻结为单一 `backend/internal/templatereuse/service/query_service.go`，并要求其通过 `candidate/` 子包读取 canonical 表，不把逻辑散落到 `review / productregistry / reusesummary` 里。

#### Scenario: 后端 query owner 的单值分工

- **WHEN** 后续实现 `phase09` 后端读能力
- **THEN** 必须保持以下单值分工：
  - `review.QueryService`：只承接 `Weekly Review` 页面级组合读取
  - `reusesummary.QueryService`：只承接 `module_reuse_summary / capability_summary`
  - `templatereuse.QueryService`：承接模板候选派生、`templateCandidateId` 解析、模板预填、派生提示与模板来源复读
  - `productregistry.QueryService`：只承接 canonical `Product Detail` 与绑定 reread
- **AND** 不得把模板候选 derivation 藏回 `review.QueryService` 私有 helper

#### Scenario: `templatereuse.QueryService` 的最小物理落点

- **WHEN** 后续实现 `phase09` 后端模块
- **THEN** 至少必须新增：
  - `backend/internal/templatereuse/service/query_service.go`
  - `backend/internal/templatereuse/candidate/template_candidate_readers.go`
  - `backend/internal/templatereuse/connect/server.go`
- **AND** `candidate/` 子包可以直接读取 `products / modules / product_modules`
- **AND** service 层不得直接写跨模块 SQL

### Requirement: `Review / Product / Decision / ReuseSummary` 的协作边界必须保持 canonical 主线不被复制

系统 SHALL 明确 `TemplateReuseService` 对既有服务的协作边界，避免 `phase09` 把现有业务事实复制成第二套。

#### Scenario: Review 不是模板候选 canonical 事实源

- **WHEN** `Weekly Review` 消费模板候选或提示
- **THEN** `ReviewService` 只允许提供：
  - 页面级消费作用域
  - 返回链元数据
  - 页面组合 transport
- **AND** 模板候选 canonical 事实源必须继续是 `product_modules`
- **AND** `review.QueryService` 不得直接承接模板候选 SQL、模板预填解析或提示派生逻辑

#### Scenario: ReuseSummary 只提供复用摘要，不承接模板候选主线

- **WHEN** `phase09` 需要复用 `module_reuse_summary / capability_summary`
- **THEN** 只能继续通过既有 `ReuseSummaryService`
- **AND** `TemplateReuseService` 可以消费其结果或对齐其排序/说明语义
- **AND** 不得把模板候选、active candidate 或预填详情写回 `reuse_summary.proto`

#### Scenario: Product 与 Decision 保持既有 canonical 主线

- **WHEN** `phase09` 触发创建 Product、绑定 Module、创建 Decision 或模块补齐
- **THEN** 正式写路径仍然只能分别回到：
  - `productregistry.CommandService`
  - `decisioncenter.CommandService`
  - 既有 `Module Registry / Repository Binding` canonical path
- **AND** 当前阶段不得新增 `TemplateReuseCommandService`
- **AND** 不得新增 `CreateProductFromTemplate`、`ApplyTemplateToProduct` 或等价第二套写 RPC

### Requirement: 当前阶段必须明确“无需新增轻量快照记录”，并写出纯读时派生足够的证据

系统 SHALL 在 `phase09-07` 正式判定：当前阶段无需新增轻量快照记录，模板候选、预填与来源复读全部继续基于读时派生完成。

#### Scenario: 纯读时派生足够的证据链

- **WHEN** 评估是否需要模板快照表、候选缓存表或预填持久化记录
- **THEN** 当前阶段必须明确以下证据已经成立：
  - 模板候选 canonical 来源已冻结为 `product_modules` 已持久化绑定事实
  - 去重键、排序规则与 `templateCandidateId` 派生规则都已冻结
  - `templateCandidateId` 漂移时的 unavailable 成功态已经冻结
  - active candidate 只是排序结果上的默认选中与前端会话切换，不是需要后端持久化的新事实
  - `Product Create -> Product Detail` 的模板来源复读链已通过 search 参数单值保留，无需新增服务器侧会话表
- **AND** 因此当前阶段必须判定“无需新增记录”

#### Scenario: 若未来有人提议引入轻量快照记录

- **WHEN** 后续方案试图新增模板快照记录
- **THEN** 必须先给出“纯读时派生不足”的明确证据，例如：
  - 性能瓶颈已被实际测量并证明确实无法接受
  - 读时派生无法满足已冻结的可恢复 unavailable 语义
  - 现有回流链无法稳定支持来源复读
- **AND** 若没有这些证据，方案必须被判定为不允许
- **AND** 即使未来允许，也只能作为“受控支撑资产”，不得升级为新的事实主线

### Requirement: 模板候选、active candidate 与来源复读链必须在后端合同层保持单值语义

系统 SHALL 在后端合同层继续冻结模板候选、active candidate 与来源复读链的单值语义，避免前后端各猜一套。

#### Scenario: active candidate 的后端合同语义

- **WHEN** `ListTemplateCandidates` 返回候选列表
- **THEN** 后端必须只返回排序后的 `candidates[]` 与 `default_active_candidate_id`
- **AND** 后端不得持久化“当前用户 active candidate”
- **AND** 前端切换 active candidate 后，后续提示读取只能通过显式传入 `template_candidate_id` 继续读取

#### Scenario: 创建成功后的模板来源复读链

- **WHEN** 用户通过模板预填创建 Product 成功
- **THEN** 成功回流链必须继续只保留：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
- **AND** `GetTemplateSourceSummary` 必须仅依赖这组参数完成复读
- **AND** 不得依赖服务端 session、额外落表或页面内存

### Requirement: `TemplateReuseService` 的 Connect 装配方式必须沿用 phase07 主线

系统 SHALL 将模板复用模块的后端接线方式冻结为“generated Connect handler + query service + chi `/api` shell”。

#### Scenario: Connect handler 的物理落点与挂载方式

- **WHEN** 后续实现 `TemplateReuseService`
- **THEN** 必须新增 `backend/internal/templatereuse/connect/server.go`
- **AND** `backend/internal/platform/router.go` 必须新增 `mountTemplateReuseConnect(...)`
- **AND** handler 必须通过 generated `(path, handler)` 挂到单一 `/api` 业务树下
- **AND** 不得新增第二个 `/rpc`、`/template-api` 或手写 JSON canonical API

#### Scenario: 错误映射继续收敛到单值入口

- **WHEN** `TemplateReuseService` 返回 domain error
- **THEN** 必须继续统一走既有 Connect 错误映射入口
- **AND** 对于 `templateCandidateId` 漂移这类可恢复场景，必须返回成功 envelope 中的 `UNAVAILABLE`，而不是 transport error
- **AND** 不得在 handler 内拼第二套 JSON/Connect 错误语义

### Requirement: `phase09` 关键路径的工具链、API smoke 与浏览器验收口径必须冻结

系统 SHALL 将 `phase09-07` 的后端验收口径冻结为“合同工具链 + Go 构建 + 前端类型校验 + API smoke + 浏览器验收”五层。

#### Scenario: 合同与构建验收口径

- **WHEN** 团队验证 `phase09-07` 设计是否可直接进入实现
- **THEN** 最小工具链口径必须冻结为：
  - `(cd proto && make build && make gen && make lint)`
  - `(cd proto && make breaking)` 作为有 `main` 基准时的 breaking gate
  - `(cd backend && go build ./...)`
  - `(cd frontend && npm run build)` 作为当前仓库的正式前端 type-check 口径
- **AND** 当前阶段不得重新发明第二套 proto 生成脚本或前端类型检查入口

#### Scenario: `TemplateReuseService` API smoke 清单

- **WHEN** 团队执行 `phase09` 后端关键路径 smoke
- **THEN** 至少必须覆盖：
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/ListTemplateCandidates`
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/GetTemplateCandidatePrefill`
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/GetDerivedInsightHints`
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/GetTemplateSourceSummary`
  - `POST /api/psco.review.v1.ReviewService/GetWeeklyReviewContext` 继续可用
  - `POST /api/psco.product_registry.v1.ProductRegistryService/CreateProduct` 继续可用
- **AND** smoke 中必须至少显式覆盖一种 `UNAVAILABLE` 成功态

#### Scenario: 浏览器验收清单的最小闭环

- **WHEN** 团队执行 `phase09` 浏览器验收
- **THEN** 至少必须覆盖以下闭环：
  1. `Weekly Review` 成功读取模板候选，并默认选中第一候选
  2. 切换 active candidate 后，派生提示随之更新
  3. 从模板进入 `Product Create`，成功读取预填详情
  4. `templateCandidateId` 失效时，`Product Create` 退化为空白 create 但仍可提交
  5. 创建成功后进入 `Product Detail`，成功复读模板来源摘要或 unavailable 空态
- **AND** 不得只凭 API 成功就判定 phase09 闭环成立

## MODIFIED Requirements

### Requirement: `phase09-04` 中 `TemplateReuseService` 的解释口径

`phase09-04` 已冻结 `TemplateReuseService` 是模板读能力的 canonical transport owner。

自 `phase09-07` 起，系统必须把该 requirement 进一步修改为：

- `TemplateReuseService` 的三类核心读取接口固定为：
  - `ListTemplateCandidates`
  - `GetTemplateCandidatePrefill`
  - `GetDerivedInsightHints`
- `GetTemplateSourceSummary` 作为创建成功后的来源复读接口单独存在
- `Product Create` 场景下的 `capability_gap_hint` 继续由 `GetTemplateCandidatePrefill` 返回，不再并列新增第二次提示读取

#### Scenario: 判断 `TemplateReuseService` 是否仍保持单一职责

- **WHEN** 后续实现或验收检查模板复用后端合同
- **THEN** 必须同时满足：
  - 候选列表、预填详情、派生提示、来源复读四类读取位职责不重叠
  - 不存在第二个 canonical 模板 transport owner
  - 不存在模板写入 RPC
- **AND** 不得通过“先放到 ReviewService 里，后面再拆”绕过该边界

## REMOVED Requirements

### Requirement: 在 `Review / ReuseSummary / ProductRegistry` 中分散扩写模板读取能力，或先引入轻量快照记录再决定事实源

**Reason**: 这会直接破坏 `.proto` 单一合同源、phase07 的 Connect 主线，以及 `phase09-04 / 05 / 06` 已冻结的 caller-owner 与返回链边界，并把模板候选事实源重新拉回多中心状态。

**Migration**: 模板候选、预填、提示与来源复读统一回收到 `template_reuse.proto + TemplateReuseService + templatereuse.QueryService`；候选数据继续只从 `product_modules` 读时派生；若未来真的需要轻量快照，只能通过新的独立 `/spec` 以受控支撑资产身份重新仲裁。
