# phase04-08 Product / Repository / Binding 最小 Protocol Buffers 合同设计 Spec

## Why

`phase04-04` 已冻结 `Product / Repository / Binding` 的最小数据读写范围、错误语义与详情/候选边界，`phase04-07` 又把后端模块边界、方向级 API 矩阵、跨模块关系摘要读取边界与文件落点收敛到了可直接进入实现的粒度。当前仍缺少一份可以进入仓库 `proto/` 主线的最小 `.proto` 合同设计，用来回答“`Product Registry` 与 `Repository Binding` 各自有哪些消息、字段编号如何冻结、服务接口如何命名、包名和版本如何演进、以及 `.proto` 如何与 `chi + JSON HTTP` 过渡传输层保持显式映射”。否则后续实现仍会回到手写 JSON DTO 扩张，重新制造并列合同源。

## What Changes

- 为 `Product / Repository / Binding` 新增最小 `.proto` 合同设计，覆盖当前阶段全部正式动作与候选读取
- 冻结 `package`、版本语义、文件落点、消息结构、字段编号、枚举与服务接口
- 冻结 `ProductListRead / ProductDetailRead / ProductCreateWrite / ProductModuleBindingWrite / ProductModuleCandidateRead` 的 request / response 语义
- 冻结 `RepositoryListRead / RepositoryDetailRead / RepositoryCreateWrite / RepositoryProductBindingWrite / RepositoryModuleMappingWrite / ProductBindingCandidateRead / RepositoryModuleCandidateRead` 的 request / response 语义
- 冻结 `.proto` 与 `chi + JSON HTTP` 过渡传输层之间的显式映射策略
- 冻结合同演进规则，包括 `reserved` 约束、`buf breaking` 前提与 `v1 -> v2` 升级边界
- **BREAKING**：后续 `Product / Repository / Binding` 的实现与验收不得再把手写 JSON 结构视为并列合同源，`.proto` 成为当前阶段唯一合同源

## Impact

- Affected specs:
  - `phase04_04_product_repository_binding_data_api_error_boundary`
  - `phase04_07_backend_module_boundary_interface_grouping`
  - 后续 `phase04` 正式规格正文
- Affected code:
  - 预期新增 `proto/psco/common/v1/common.proto`
  - 预期新增 `proto/psco/product_registry/v1/product_registry.proto`
  - 预期新增 `proto/psco/repository_binding/v1/repository_binding.proto`
  - 预期补充 `proto/README.md` 中关于 `phase04` 的合同说明与 RPC → HTTP 映射
  - 预期约束后续 `backend/internal/productregistry/types.go`
  - 预期约束后续 `backend/internal/repositorybinding/types.go`

## ADDED Requirements

### Requirement: phase04 Proto 合同源必须进入现有单一 proto workspace

系统 SHALL 为 `phase04` 落地单一 `.proto` 合同源集合，并复用仓库现有 `proto/` 工作区、`buf.yaml`、`buf.gen.yaml` 与 `Makefile`。

#### Scenario: 合同源落地

- **WHEN** 执行 `phase04-08`
- **THEN** 仓库中必须存在可追踪的 `.proto` 文件落点
- **AND** 该合同源集合必须覆盖 `Product Registry` 与 `Repository Binding` 当前阶段全部正式动作与候选读取
- **AND** 不得新增第二套 `buf.yaml`、`buf.gen.yaml`、`Makefile` 或并列 proto 根目录
- **AND** 不得再以手写 JSON 结构充当并列合同源

### Requirement: Proto 包名、目录与版本语义必须冻结

系统 SHALL 为 `phase04` 当前阶段的 Proto 合同冻结明确的包名、目录与版本语义，直接进入现有 `proto/` 主线。

#### Scenario: 目录与包名冻结

- **WHEN** 新增 `phase04` 的 `.proto` 合同源
- **THEN** 必须冻结以下目录与包名：
  - `proto/psco/common/v1/common.proto` → `package psco.common.v1`
  - `proto/psco/product_registry/v1/product_registry.proto` → `package psco.product_registry.v1`
  - `proto/psco/repository_binding/v1/repository_binding.proto` → `package psco.repository_binding.v1`
- **AND** `common.proto` 只允许承接跨 `Product / Repository` 共享且不会引入业务 owner 歧义的最小公共枚举
- **AND** `product_registry.proto` 只允许承接 `Product Registry` 的服务接口与其拥有的消息
- **AND** `repository_binding.proto` 只允许承接 `Repository Binding` 的服务接口与其拥有的消息
- **AND** 后续新增字段必须在各自 `v1` 演进规则下进行，而不是临时改写包名

#### Scenario: Go 包路径冻结

- **WHEN** 冻结 `phase04` `.proto` 合同源的生成落点
- **THEN** `go_package` 必须分别冻结为：
  - `github.com/psco/backend/internal/gen/proto/psco/common/v1;commonv1`
  - `github.com/psco/backend/internal/gen/proto/psco/product_registry/v1;productregistryv1`
  - `github.com/psco/backend/internal/gen/proto/psco/repository_binding/v1;repositorybindingv1`
- **AND** TypeScript 生成产物必须继续落在 `frontend/src/gen/proto/psco/...`

### Requirement: 公共 active / archived 状态枚举必须单值化

系统 SHALL 为 `Product` 与 `Repository` 共享的 `active / archived` 状态语义冻结单一 Proto 枚举，避免在 `product_registry.proto` 与 `repository_binding.proto` 中出现两套等价枚举，也避免形成跨包循环依赖。

#### Scenario: 共享状态枚举冻结

- **WHEN** 定义 `Product` 与 `Repository` 相关消息中的 `status`
- **THEN** 必须在 `psco.common.v1` 中定义单一枚举 `ActiveArchivedStatus`
- **AND** 枚举值必须冻结为：
  - `ACTIVE_ARCHIVED_STATUS_UNSPECIFIED = 0`
  - `ACTIVE_ARCHIVED_STATUS_ACTIVE = 1`
  - `ACTIVE_ARCHIVED_STATUS_ARCHIVED = 2`
- **AND** `Product` / `Repository` 自身的 `status`、列表项 `status`、详情/候选中 Product / Repository 摘要的 `status`（如 `BoundRepositorySummary.repository_status` / `BoundProductSummary.product_status` / `RepositoryProductCandidate.product_status`）与写组 request 中的 `status` 必须统一复用 `ActiveArchivedStatus`
- **AND** Module 摘要的 `module_status`（如 `BoundModuleSummary.module_status` / `MappedModuleSummary.module_status` / `ProductModuleCandidate.module_status` / `RepositoryModuleCandidate.module_status`）通过 import 复用 `psco.module_registry.v1.ModuleStatus`，不在此枚举覆盖范围内
- **AND** 不得在 `psco.product_registry.v1` 或 `psco.repository_binding.v1` 中重定义本地等价枚举
- **AND** `ACTIVE_ARCHIVED_STATUS_UNSPECIFIED` 在列表读取 `status_filter` 字段中表示“不过滤”，对应 UI/路由层的 `all`；`all` 不得作为枚举值写入 `.proto`，承接 `phase04-04` `statusFilter` 语义冻结（`all` 只存在于 UI/路由层，不得写入数据库、HTTP 持久化 DTO 或 `.proto` 持久化字段）
- **AND** `ACTIVE_ARCHIVED_STATUS_UNSPECIFIED` 在 `status_filter` 之外的字段（如 `CreateProductRequest.status` / `CreateRepositoryRequest.status`）中不允许作为合法写入值，由 service 层校验拒绝，承接 `phase04-04` 必填字段约束（`status` 必填且必须属于 `active / archived`）
- **AND** 该 UNSPECIFIED = 不过滤语义与现有 `psco.module_registry.v1.ModuleStatus` 与 `psco.decision_center.v1.DecisionStatus` 的 UNSPECIFIED 语义一致，承接 `phase02-11A` 与 `phase03-11` 已建立模式

### Requirement: Product Registry 最小消息结构必须覆盖当前动作矩阵

系统 SHALL 让 `product_registry.proto` 的消息结构完整承接 `phase04-04 / 06 / 07` 已冻结的最小读写模型、页面动作与 reread 语义。

#### Scenario: Product Registry 核心对象与读组消息

- **WHEN** 定义 `Product Registry` 的核心对象与读组消息
- **THEN** 必须至少冻结以下消息：
  - `Product`
  - `ProductListItem`
  - `BoundModuleSummary`
  - `BoundRepositorySummary`
  - `ProductDetail`
  - `ListProductsRequest / ListProductsResponse`
  - `GetProductDetailRequest / GetProductDetailResponse`
- **AND** 字段语义必须单值对应 `phase04-04` 已冻结的数据范围：
  - `Product` → `id / name / description / status / created_at`
  - `ProductListItem` → `id / name / description / status / created_at / module_bind_count / repository_bind_count`
  - `BoundModuleSummary` → `module_id / module_name / module_status`
  - `BoundRepositorySummary` → `repository_id / repository_name / provider / repository_status`
  - `ProductDetail` → `product + repeated BoundModuleSummary + repeated BoundRepositorySummary`
- **AND** `BoundModuleSummary.module_status` 必须通过 import 直接复用 `psco.module_registry.v1.ModuleStatus`
- **AND** `BoundRepositorySummary.repository_status` 必须复用 `psco.common.v1.ActiveArchivedStatus`
- **AND** `ProductDetailRead` 不得内嵌候选读取结果

#### Scenario: Product Registry 写组与候选读取消息

- **WHEN** 定义 `Product Registry` 的写组与候选读取消息
- **THEN** 必须至少冻结以下消息：
  - `CreateProductRequest / CreateProductResponse`
  - `BindModuleToProductRequest / BindModuleToProductResponse`
  - `ProductModuleCandidate`
  - `ListProductModuleCandidatesRequest / ListProductModuleCandidatesResponse`
- **AND** `CreateProductRequest` 必须覆盖 `name / description / status`
- **AND** `BindModuleToProductRequest` 必须覆盖 `product_id / module_id`
- **AND** `ProductModuleCandidate` 必须覆盖 `module_id / module_name / module_status`
- **AND** `CreateProductResponse` 必须只返回 `product_id`
- **AND** `BindModuleToProductResponse` 必须为空响应
- **AND** `ListProductModuleCandidatesRequest` 必须显式承接 `product_id`
- **AND** `ProductModuleCandidate.module_status` 必须复用 `psco.module_registry.v1.ModuleStatus`

### Requirement: Repository Binding 最小消息结构必须覆盖当前动作矩阵

系统 SHALL 让 `repository_binding.proto` 的消息结构完整承接 `phase04-04 / 06 / 07` 已冻结的最小读写模型、页面动作与 reread 语义。

#### Scenario: Repository Binding 核心对象与读组消息

- **WHEN** 定义 `Repository Binding` 的核心对象与读组消息
- **THEN** 必须至少冻结以下消息：
  - `Repository`
  - `RepositoryListItem`
  - `BoundProductSummary`
  - `MappedModuleSummary`
  - `RepositoryDetail`
  - `ListRepositoriesRequest / ListRepositoriesResponse`
  - `GetRepositoryDetailRequest / GetRepositoryDetailResponse`
- **AND** 字段语义必须单值对应 `phase04-04` 已冻结的数据范围：
  - `Repository` → `id / name / url / provider / status / created_at`
  - `RepositoryListItem` → `id / name / url / provider / status / created_at / product_bind_count / module_bind_count`
  - `BoundProductSummary` → `product_id / product_name / product_status`
  - `MappedModuleSummary` → `module_id / module_name / module_status`
  - `RepositoryDetail` → `repository + repeated BoundProductSummary + repeated MappedModuleSummary`
- **AND** `BoundProductSummary.product_status` 必须复用 `psco.common.v1.ActiveArchivedStatus`
- **AND** `MappedModuleSummary.module_status` 必须通过 import 直接复用 `psco.module_registry.v1.ModuleStatus`
- **AND** `RepositoryDetailRead` 不得内嵌候选读取结果

#### Scenario: Repository Binding 写组与候选读取消息

- **WHEN** 定义 `Repository Binding` 的写组与候选读取消息
- **THEN** 必须至少冻结以下消息：
  - `CreateRepositoryRequest / CreateRepositoryResponse`
  - `BindRepositoryToProductRequest / BindRepositoryToProductResponse`
  - `MapModuleToRepositoryRequest / MapModuleToRepositoryResponse`
  - `RepositoryProductCandidate`
  - `RepositoryModuleCandidate`
  - `ListRepositoryProductCandidatesRequest / ListRepositoryProductCandidatesResponse`
  - `ListRepositoryModuleCandidatesRequest / ListRepositoryModuleCandidatesResponse`
- **AND** `CreateRepositoryRequest` 必须覆盖 `name / url / provider / status`
- **AND** `BindRepositoryToProductRequest` 必须覆盖 `repository_id / product_id`
- **AND** `MapModuleToRepositoryRequest` 必须覆盖 `repository_id / module_id`
- **AND** `RepositoryProductCandidate` 必须覆盖 `product_id / product_name / product_status`
- **AND** `RepositoryModuleCandidate` 必须覆盖 `module_id / module_name / module_status`
- **AND** `CreateRepositoryResponse` 必须只返回 `repository_id`
- **AND** `BindRepositoryToProductResponse` 与 `MapModuleToRepositoryResponse` 必须为空响应
- **AND** `ListRepositoryProductCandidatesRequest` 与 `ListRepositoryModuleCandidatesRequest` 都必须显式承接 `repository_id`

### Requirement: 字段语义与页面动作必须单值映射

系统 SHALL 将 `.proto` 合同字段语义与当前页面动作、HTTP DTO 以及 reread 语义保持单值一致，不留给实现阶段临场决定。

#### Scenario: Product Registry 字段与页面动作映射

- **WHEN** 校验 `Product Registry` 的合同字段语义
- **THEN** `ListProductsResponse` 必须只对应 `Product Registry / List` 页面展示字段
- **AND** `GetProductDetailResponse` 必须只对应 `Product Detail` 页面详情展示字段
- **AND** `CreateProductRequest` 必须只对应 `Product Create` 页面表单字段
- **AND** `CreateProductResponse.product_id` 必须支撑前端成功后进入新 `Product Detail`
- **AND** `BindModuleToProductRequest` 必须只对应 `Product Detail` 上的绑定写入动作
- **AND** `ListProductModuleCandidatesResponse` 必须只对应 `Product Detail` 的候选读取面板

#### Scenario: Repository Binding 字段与页面动作映射

- **WHEN** 校验 `Repository Binding` 的合同字段语义
- **THEN** `ListRepositoriesResponse` 必须只对应 `Repository Binding / List` 页面展示字段
- **AND** `GetRepositoryDetailResponse` 必须只对应 `Repository Binding Detail / Workspace` 页面展示字段
- **AND** `CreateRepositoryRequest` 必须只对应 `Repository Create` 页面表单字段
- **AND** `CreateRepositoryResponse.repository_id` 必须支撑前端成功后进入新 `Repository Binding Detail / Workspace`
- **AND** `BindRepositoryToProductRequest` 与 `MapModuleToRepositoryRequest` 必须只对应 `Repository Binding Detail / Workspace` 上的两个绑定动作
- **AND** `ListRepositoryProductCandidatesResponse` 与 `ListRepositoryModuleCandidatesResponse` 必须只对应当前工作台的两个候选读取面板

#### Scenario: 排序规则不进入 .proto 合同本体

- **WHEN** 后续实现讨论列表读取与候选读取的排序语义（`phase04-04` 冻结的列表读取与三条候选读取均按 `created_at` 降序）
- **THEN** 排序规则由 service / repository 层承接，不进入 `.proto` 合同本体
- **AND** `.proto` 合同只定义消息结构与字段，不定义 `order_by` 字段或排序注解
- **AND** 该边界承接现有 `psco.decision_center.v1.decision_center.proto` 已建立模式（候选读取排序语义由 service/repository 层承接，不进入 `.proto` 合同本体）

### Requirement: 字段编号方案必须在当前阶段冻结

系统 SHALL 为 `phase04` `.proto` 合同冻结明确的字段编号方案，使后续 `.proto` 落地成为机械映射，而不是在实现期再发明编号。

#### Scenario: Product Registry 枚举与核心对象字段编号冻结

- **WHEN** 定义 `Product Registry` 的消息编号
- **THEN** 必须至少冻结以下字段编号方案：
  - `Product`：`id=1` / `name=2` / `description=3` / `status=4(ActiveArchivedStatus)` / `created_at=5(google.protobuf.Timestamp)`
  - `ProductListItem`：`id=1` / `name=2` / `description=3` / `status=4(ActiveArchivedStatus)` / `created_at=5(Timestamp)` / `module_bind_count=6(int32)` / `repository_bind_count=7(int32)`
  - `BoundModuleSummary`：`module_id=1` / `module_name=2` / `module_status=3(ModuleStatus)`
  - `BoundRepositorySummary`：`repository_id=1` / `repository_name=2` / `provider=3` / `repository_status=4(ActiveArchivedStatus)`
  - `ProductDetail`：`product=1(Product)` / `bound_modules=2(repeated BoundModuleSummary)` / `bound_repositories=3(repeated BoundRepositorySummary)`

#### Scenario: Product Registry request / response 字段编号冻结

- **WHEN** 定义 `Product Registry` 的 request / response 编号
- **THEN** 必须至少冻结以下编号方案：
  - `ListProductsRequest`：`query_text=1(string)` / `status_filter=2(ActiveArchivedStatus)`
  - `ListProductsResponse`：`products=1(repeated ProductListItem)`
  - `GetProductDetailRequest`：`product_id=1(string)`
  - `GetProductDetailResponse`：`product_detail=1(ProductDetail)`
  - `CreateProductRequest`：`name=1(string)` / `description=2(string)` / `status=3(ActiveArchivedStatus)`
  - `CreateProductResponse`：`product_id=1(string)`
  - `BindModuleToProductRequest`：`product_id=1(string)` / `module_id=2(string)`
  - `ProductModuleCandidate`：`module_id=1(string)` / `module_name=2(string)` / `module_status=3(ModuleStatus)`
  - `ListProductModuleCandidatesRequest`：`product_id=1(string)`
  - `ListProductModuleCandidatesResponse`：`candidates=1(repeated ProductModuleCandidate)`

#### Scenario: Repository Binding 核心对象字段编号冻结

- **WHEN** 定义 `Repository Binding` 的消息编号
- **THEN** 必须至少冻结以下字段编号方案：
  - `Repository`：`id=1` / `name=2` / `url=3` / `provider=4` / `status=5(ActiveArchivedStatus)` / `created_at=6(google.protobuf.Timestamp)`
  - `RepositoryListItem`：`id=1` / `name=2` / `url=3` / `provider=4` / `status=5(ActiveArchivedStatus)` / `created_at=6(Timestamp)` / `product_bind_count=7(int32)` / `module_bind_count=8(int32)`
  - `BoundProductSummary`：`product_id=1` / `product_name=2` / `product_status=3(ActiveArchivedStatus)`
  - `MappedModuleSummary`：`module_id=1` / `module_name=2` / `module_status=3(ModuleStatus)`
  - `RepositoryDetail`：`repository=1(Repository)` / `bound_products=2(repeated BoundProductSummary)` / `mapped_modules=3(repeated MappedModuleSummary)`

#### Scenario: Repository Binding request / response 字段编号冻结

- **WHEN** 定义 `Repository Binding` 的 request / response 编号
- **THEN** 必须至少冻结以下编号方案：
  - `ListRepositoriesRequest`：`query_text=1(string)` / `status_filter=2(ActiveArchivedStatus)`
  - `ListRepositoriesResponse`：`repositories=1(repeated RepositoryListItem)`
  - `GetRepositoryDetailRequest`：`repository_id=1(string)`
  - `GetRepositoryDetailResponse`：`repository_detail=1(RepositoryDetail)`
  - `CreateRepositoryRequest`：`name=1(string)` / `url=2(string)` / `provider=3(string)` / `status=4(ActiveArchivedStatus)`
  - `CreateRepositoryResponse`：`repository_id=1(string)`
  - `BindRepositoryToProductRequest`：`repository_id=1(string)` / `product_id=2(string)`
  - `MapModuleToRepositoryRequest`：`repository_id=1(string)` / `module_id=2(string)`
  - `RepositoryProductCandidate`：`product_id=1(string)` / `product_name=2(string)` / `product_status=3(ActiveArchivedStatus)`
  - `RepositoryModuleCandidate`：`module_id=1(string)` / `module_name=2(string)` / `module_status=3(ModuleStatus)`
  - `ListRepositoryProductCandidatesRequest`：`repository_id=1(string)`
  - `ListRepositoryProductCandidatesResponse`：`candidates=1(repeated RepositoryProductCandidate)`
  - `ListRepositoryModuleCandidatesRequest`：`repository_id=1(string)`
  - `ListRepositoryModuleCandidatesResponse`：`candidates=1(repeated RepositoryModuleCandidate)`

### Requirement: 服务接口必须对齐 phase04-07 已冻结的接口分组

系统 SHALL 将 Proto service 接口与 `phase04-07` 已冻结的方向级 API 矩阵保持单值一致，不引入第二套命名体系。

#### Scenario: Product Registry service 矩阵

- **WHEN** 定义 `Product Registry` 的 Proto service
- **THEN** 必须冻结单一服务 `ProductRegistryService`
- **AND** 最小 RPC 矩阵必须为：
  - `ListProducts` → 对齐 `ProductListRead`
  - `GetProductDetail` → 对齐 `ProductDetailRead`
  - `CreateProduct` → 对齐 `ProductCreateWrite`
  - `BindModuleToProduct` → 对齐 `ProductModuleBindingWrite`
  - `ListProductModuleCandidates` → 对齐 `ProductModuleCandidateRead`

#### Scenario: Repository Binding service 矩阵

- **WHEN** 定义 `Repository Binding` 的 Proto service
- **THEN** 必须冻结单一服务 `RepositoryBindingService`
- **AND** 最小 RPC 矩阵必须为：
  - `ListRepositories` → 对齐 `RepositoryListRead`
  - `GetRepositoryDetail` → 对齐 `RepositoryDetailRead`
  - `CreateRepository` → 对齐 `RepositoryCreateWrite`
  - `BindRepositoryToProduct` → 对齐 `RepositoryProductBindingWrite`
  - `MapModuleToRepository` → 对齐 `RepositoryModuleMappingWrite`
  - `ListRepositoryProductCandidates` → 对齐 `ProductBindingCandidateRead`
  - `ListRepositoryModuleCandidates` → 对齐 `RepositoryModuleCandidateRead`

#### Scenario: phase02 已落地 RPC 与 phase04 新增 RPC 的迁移承接

- **WHEN** `phase04` 新增 `psco.product_registry.v1.ProductRegistryService.BindModuleToProduct` 与 `psco.repository_binding.v1.RepositoryBindingService.MapModuleToRepository`
- **THEN** 必须显式承接：现有 `psco.module_registry.v1.ModuleRegistryService.BindModuleToProduct` / `MapModuleToRepository`（`phase02-11A` 落地）与 phase04 新增 RPC 在当前阶段共存
- **AND** 共存期间旧 RPC 视为兼容适配入口，canonical owner 以 phase04 新 RPC 为准（承接 `phase04-07` "旧 transport 入口若保留只能作为兼容适配层"口径）
- **AND** `.proto` 合同源层面不主动删除 `module_registry.proto` 中的旧 RPC 定义，避免触发 `buf breaking` 失败
- **AND** 后续 `phase04-11` 落地时由实现任务决定是否对旧 RPC 标注 `deprecated`，本阶段不提前冻结删除时机
- **AND** 新 RPC 的 request 字段顺序反映各自 owner 上下文，属不同包内独立编号，不构成 wire 兼容问题：
  - `ProductRegistryService.BindModuleToProductRequest` 以 `product_id=1` 为主键（spec L230），`ModuleRegistryService.BindModuleToProductRequest` 以 `module_id=1` 为主键（phase02-11A `module_registry.proto` L160-163）
  - `RepositoryBindingService.MapModuleToRepositoryRequest` 以 `repository_id=1` 为主键（spec L256），`ModuleRegistryService.MapModuleToRepositoryRequest` 以 `module_id=1` 为主键（phase02-11A `module_registry.proto` L165-169）

### Requirement: 详情读取与候选读取必须保持消息边界分离

系统 SHALL 保持 `ProductDetailRead / RepositoryDetailRead` 与三条候选读取在合同层的边界分离，不得为了图省事把候选结果塞进详情响应。

#### Scenario: Product Detail 与候选读取边界

- **WHEN** 定义 `GetProductDetailResponse` 与 `ListProductModuleCandidatesResponse`
- **THEN** `GetProductDetailResponse` 只承接详情本体、已绑定模块列表与已绑定仓库列表
- **AND** `ListProductModuleCandidatesResponse` 只承接候选 `Module` 列表
- **AND** 不得把候选结果直接塞进 `ProductDetail`

#### Scenario: Repository Detail 与候选读取边界

- **WHEN** 定义 `GetRepositoryDetailResponse`、`ListRepositoryProductCandidatesResponse` 与 `ListRepositoryModuleCandidatesResponse`
- **THEN** `GetRepositoryDetailResponse` 只承接详情本体、已绑定产品列表与已映射模块列表
- **AND** 两个候选读取响应只承接对应候选结果
- **AND** 不得把候选结果直接塞进 `RepositoryDetail`

#### Scenario: 详情内已绑定列表摘要消息与 phase04-07 跨模块关系摘要读取接口的对应关系

- **WHEN** 后续实现讨论 `ProductDetail` / `RepositoryDetail` 中已绑定列表摘要字段（`bound_modules` / `bound_repositories` / `bound_products` / `mapped_modules`）的后端填充方式
- **THEN** `.proto` 合同只定义摘要消息结构（`BoundModuleSummary` / `BoundRepositorySummary` / `BoundProductSummary` / `MappedModuleSummary`），不定义跨模块关系摘要读取接口的 RPC
- **AND** 这四个摘要消息对应 `phase04-07` 冻结的四条跨模块关系摘要读取链路，后端填充方式如下：
  - `BoundModuleSummary` → 由 `ProductModuleSummaryRead` 接口填充（owner = `Product Registry`，实现在 `productregistry/repository/binding_store.go`）
  - `BoundRepositorySummary` → 由 `ProductRepositorySummaryRead` 接口填充（owner = `Repository Binding`，实现在 `repositorybinding/repository/binding_store.go`）
  - `BoundProductSummary` → 由 `RepositoryProductSummaryRead` 接口填充（owner = `Repository Binding`，实现在 `repositorybinding/repository/binding_store.go`）
  - `MappedModuleSummary` → 由 `RepositoryModuleSummaryRead` 接口填充（owner = `Repository Binding`，实现在 `repositorybinding/repository/binding_store.go`）
- **AND** 这四个跨模块关系摘要读取接口是后端模块内部的跨模块读接口，不直接暴露为 RPC，前端只通过 `GetProductDetail` 与 `GetRepositoryDetail` RPC 获取已绑定列表
- **AND** `.proto` 合同的字段范围与 `phase04-07` 冻结的四个跨模块关系摘要读取接口字段范围单值一致，承接 `phase04-04` 详情读取数据范围冻结

### Requirement: 错误语义必须在合同设计中保留显式映射前提

系统 SHALL 在 `.proto` 合同设计中为当前阶段已冻结的错误语义保留稳定承接前提，但不把 HTTP 状态码本身写进 `.proto`。

#### Scenario: Product Registry 错误语义承接

- **WHEN** 设计 `CreateProduct`、`GetProductDetail`、`BindModuleToProduct`、`ListProductModuleCandidates` 的合同边界
- **THEN** 必须显式保留校验失败、资源不存在、重复冲突与空结果语义的承接空间
- **AND** `ListProductModuleCandidates` 的空候选结果必须表现为正常空列表响应，不得设计为空错误
- **AND** `CreateProductResponse` 不得通过额外错误字段复制 HTTP 状态码语义

#### Scenario: Repository Binding 错误语义承接

- **WHEN** 设计 `CreateRepository`、`GetRepositoryDetail`、`BindRepositoryToProduct`、`MapModuleToRepository`、两条候选读取的合同边界
- **THEN** 必须显式保留校验失败、资源不存在、重复冲突与空结果语义的承接空间
- **AND** 两条候选读取的空结果都必须表现为正常空列表响应，不得设计为空错误
- **AND** HTTP 状态码映射继续由过渡传输层负责

### Requirement: chi + JSON HTTP 必须从 Proto 单向承接

系统 SHALL 明确当前阶段 `chi + JSON HTTP` 与 `.proto` 的关系，避免形成第二套合同源。

#### Scenario: 过渡层保留

- **WHEN** 当前阶段继续保留 `chi + JSON HTTP`
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 语义显式对齐
- **AND** 不得在 `handler`、HTTP DTO 或前端 adapter 中私自新增 `.proto` 中不存在的业务字段语义
- **AND** 不得把 HTTP 路径、状态码或中间件策略误写成 Proto 合同本体

#### Scenario: 路径参数与消息字段映射

- **WHEN** HTTP 过渡层使用 URL 路径参数承接 `productId`、`repositoryId` 或 `moduleId`
- **THEN** handler 必须在进入业务层前显式组装为对应 Proto request 字段
- **AND** `product_id / repository_id / module_id` 在 Proto RPC 中必须作为请求消息的显式字段存在
- **AND** 该差异必须被视为传输层差异，而不是并列合同定义

### Requirement: RPC 到 HTTP 的映射矩阵必须明确

系统 SHALL 为当前阶段的最小动作矩阵冻结单值的 RPC → HTTP 映射矩阵。

#### Scenario: Product Registry 映射矩阵

- **WHEN** 设计 `Product Registry` 的过渡传输层映射
- **THEN** 至少必须明确以下映射：
  - `ListProducts` → `GET /api/products`
  - `GetProductDetail` → `GET /api/products/{productId}`
  - `CreateProduct` → `POST /api/products`
  - `BindModuleToProduct` → `POST /api/products/{productId}/bindings/modules`
  - `ListProductModuleCandidates` → `GET /api/products/{productId}/candidates/modules`

#### Scenario: Repository Binding 映射矩阵

- **WHEN** 设计 `Repository Binding` 的过渡传输层映射
- **THEN** 至少必须明确以下映射：
  - `ListRepositories` → `GET /api/repositories`
  - `GetRepositoryDetail` → `GET /api/repositories/{repositoryId}`
  - `CreateRepository` → `POST /api/repositories`
  - `BindRepositoryToProduct` → `POST /api/repositories/{repositoryId}/bindings/products`
  - `MapModuleToRepository` → `POST /api/repositories/{repositoryId}/bindings/modules`
  - `ListRepositoryProductCandidates` → `GET /api/repositories/{repositoryId}/candidates/products`
  - `ListRepositoryModuleCandidates` → `GET /api/repositories/{repositoryId}/candidates/modules`

### Requirement: 合同演进与 breaking check 规则必须冻结

系统 SHALL 为 `phase04` `.proto` 合同冻结稳定的演进规则，直接复用当前仓库 `buf` 工具链与 `WIRE_JSON` breaking 语义。

#### Scenario: 字段与枚举演进规则

- **WHEN** 后续版本新增字段或枚举值
- **THEN** 必须使用新的递增编号
- **AND** 不得插入到已有编号之间
- **AND** 不得复用已删除字段编号或枚举值编号

#### Scenario: 删除字段或枚举值后的 reserved 约束

- **WHEN** 后续版本删除字段、废弃字段名、删除枚举值或废弃枚举名
- **THEN** 必须使用 `reserved` 保留编号
- **AND** 必要时必须同时 `reserved` 对应名称
- **AND** 不得在未声明 `reserved` 的情况下复用旧编号或旧名称

#### Scenario: v1 breaking 升级边界

- **WHEN** 后续修改触及以下任一情况：
  - 删除已有字段
  - 修改已有字段类型、编号或 JSON / wire 兼容语义
  - 删除已有 RPC 或修改其 request / response 兼容边界
- **THEN** 必须视为 `v1` breaking 变更
- **AND** 不得直接在现有 `v1` 包内覆盖
- **AND** 必须通过新版本目录与包名（如 `v2`）承接

#### Scenario: Buf 校验链冻结

- **WHEN** 冻结当前阶段合同工具链
- **THEN** 必须至少覆盖 `buf build`、`buf lint`、`buf generate`、`buf breaking`
- **AND** 必须复用仓库现有 `proto/Makefile`
- **AND** `buf breaking` 必须继续对照 `../.git#branch=main,subdir=proto`
- **AND** 不得吞掉 `buf breaking` 的失败退出码
- **AND** 当前阶段新增 `common / product_registry / repository_binding` 三个 `.proto` 文件属于向后兼容新增，不构成 breaking

### Requirement: 当前阶段合同落地边界必须明确

系统 SHALL 明确 `phase04-08` 当前需要落地到什么程度，避免把最小合同设计膨胀成完整通信栈改造。

#### Scenario: 当前阶段合同落地边界

- **WHEN** 执行 `phase04-08`
- **THEN** 必须冻结 `.proto` 合同源、最小消息结构、服务接口、字段编号、RPC → HTTP 映射与生成/校验入口
- **AND** 当前阶段可以不完成完整 gRPC / Connect / 网关迁移
- **AND** 当前阶段可以继续保留 `chi` 作为 HTTP 过渡传输层
- **AND** 当前阶段不要求立即用生成类型替换全部手写 DTO

## MODIFIED Requirements

### Requirement: Contract First 在 phase04 的阶段解释

`Contract First` 在 `phase04` 当前阶段 SHALL 被解释为“先冻结 `Product / Repository / Binding` 的最小数据读写范围与后端接口矩阵，再将其落地为最小 `.proto` 合同源”，而不是继续停留在长期方向。

#### Scenario: phase04 合同口径更新

- **WHEN** `phase04-08` 进入执行链
- **THEN** `Product / Repository / Binding` 的 `.proto` 必须成为当前阶段唯一合同源
- **AND** 后续正式规格正文与实现必须基于该合同源展开
- **AND** 不得继续沿用“当前阶段只冻结接口边界、不要求 proto 设计”的旧口径

## REMOVED Requirements

### Requirement: 当前阶段继续以手写 JSON DTO 作为 Product / Repository / Binding 合同主线

**Reason**: 该口径会导致 `phase04` 的前后端继续围绕手写 JSON 结构增长，把字段编号、消息边界、跨包依赖与兼容性约束继续后移。

**Migration**: 改为“当前阶段必须冻结并落地 `common + product_registry + repository_binding` 最小 `.proto` 合同源；`chi + JSON HTTP` 继续作为过渡传输层，通过显式映射承接 `.proto` 语义”。
