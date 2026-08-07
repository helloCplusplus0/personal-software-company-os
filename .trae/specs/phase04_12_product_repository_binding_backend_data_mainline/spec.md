# Phase04-12 Product / Repository / Binding 后端与数据主线 Spec

## Why

`phase04-07` 已冻结后端模块边界、读写分组、跨模块摘要读取 owner 与 `phase02` 临时接口兼容策略，`phase04-09` 已冻结联调环境、重置基线、兼容迁移与冷启动验收路径，`phase04-10` 正式规格正文又把页面、动作、数据/API、迁移与 Done 标准收敛为当前阶段唯一正文，`phase04-11` 则已把最小 `.proto` 合同正式落地为仓库内单一合同源。现在缺的不是再讨论边界，而是把这些冻结结论真正推进成可运行的后端与数据主线。

因此，`phase04-12` 必须实现 `Product Registry` 与 `Repository Binding` 的最小后端读写接口、`products / repositories / product_modules / product_repositories / module_repositories` 的数据主线、兼容委派与可重复联调基线，使后续 `phase04-13` 前端主线与联调验收可以直接消费真实后端，而不再依赖临时 SQL、旧 owner 或第二套合同解释。

## What Changes

- 实现 `backend/internal/productregistry/` 后端模块，承接 `Product List / Detail / Create / BindModuleToProduct / ListProductModuleCandidates`
- 实现 `backend/internal/repositorybinding/` 后端模块，承接 `Repository List / Detail / Create / BindRepositoryToProduct / MapModuleToRepository / 两条候选读取`
- 实现 `products / repositories` 从 `phase02` 只读前提升级到 `phase04` 主线的 migration，并新增 `product_repositories`
- 实现 `Product Registry` 与 `Repository Binding` 的 `candidate/`、`repository/`、`service/`、`handler/`、`types/errors/validate` 文件主线
- 实现 `phase04` 基线 seed、重置脚本与 `seed_readonly_prereqs.sql` 升级，使联调环境可重复恢复
- 实现 `phase02` 旧候选读取与旧模块中心绑定入口的兼容委派，确保旧 transport 若保留也不再持有业务 owner
- **BREAKING**：`Product / Repository / Binding` 的后端与数据主线从本阶段起必须以 `productregistry` / `repositorybinding` 与 `.proto` 单一合同源为唯一实现入口；旧 `Module Registry` 入口若继续保留，只能作为兼容适配层

## Impact

- Affected specs:
  - `phase04_04_product_repository_binding_data_api_error_boundary`
  - `phase04_07_backend_module_boundary_interface_grouping`
  - `phase04_09_integration_acceptance_reset_baseline_compat_migration`
  - `phase04_10_product_repository_binding_formal_spec`
  - `phase04_11_product_repository_binding_proto_mainline`
- Affected code:
  - `backend/internal/productregistry/`
  - `backend/internal/repositorybinding/`
  - `backend/internal/moduleregistry/handler/`
  - `backend/internal/moduleregistry/service/`
  - `backend/internal/platform/router.go`
  - `backend/internal/platform/server.go`
  - `database/migrations/0006_product_repository_binding_mainline.sql`
  - `database/seeds/seed_product_repository_mainline_baseline.sql`
  - `database/seeds/seed_readonly_prereqs.sql`
  - `database/scripts/reset_product_repository_mainline.sh`

## ADDED Requirements

### Requirement: Product Registry 后端模块文件落点必须与既有主线同构

系统 SHALL 在 `backend/internal/productregistry/` 下实现 `Product Registry` 后端模块，文件落点、分层语义与接线方式必须与现有 `moduleregistry`、`decisioncenter` 主线同构，并承接 `phase04-07` 已冻结的模块边界。

#### Scenario: Product Registry 文件落点冻结

- **WHEN** 实现 `Product Registry` 后端模块
- **THEN** 必须按以下落点创建文件：
  - `handler/query_handler.go`
  - `handler/command_handler.go`
  - `handler/response.go`
  - `service/query_service.go`
  - `service/command_service.go`
  - `repository/product_store.go`
  - `repository/binding_store.go`
  - `candidate/module_candidate_read.go`
  - `bound_repository_reader.go`
  - `errors.go`
  - `types.go`
  - `validate.go`
- **AND** `handler/` 只负责 HTTP 请求/响应
- **AND** `service/` 只负责校验顺序、动作编排与 reread 所需读取聚合
- **AND** `repository/` 只负责 `products / product_modules` 与 `ProductModuleSummaryRead` 摘要读取的 SQL；`ProductRepositorySummaryRead` 的 owner 是 `Repository Binding`（phase04-07 L162-181 冻结），由 `Repository Binding` 模块的 `repository/binding_store.go` 实现，`Product Registry` 通过 `productregistry.BoundRepositoryReader` 接口在 platform 装配点注入消费，不得在 `productregistry/repository/` 内复制第二套实现
- **AND** `candidate/` 只负责 `BindModuleToProduct` 所需候选读取与存在性校验
- **AND** 允许在 `repository/` 下新增辅助文件（如 `util.go` 封装 `isUniqueViolation` 等通用工具），不构成第二套命名体系或第二套错误映射
- **AND** `BoundRepositorySummary.RepositoryStatus` 与 `BoundProductSummary.ProductStatus` / `RepositoryProductCandidate.ProductStatus` 允许借用 `ProductStatus` / `RepositoryStatus` 字符串别名类型，因二者均为 `string` 别名且取值集合为 `active / archived`，与 `psco.common.v1.ActiveArchivedStatus` 单值对齐，不构成第二套枚举

#### Scenario: Product Registry 合同承接边界

- **WHEN** 定义 `backend/internal/productregistry/types.go`
- **THEN** `Product`、`ProductListItem`、`ProductDetail`、`BoundModuleSummary`、`BoundRepositorySummary`、`ProductModuleCandidate`、`CreateProductRequest`、`CreateProductResponse`、`BindModuleToProductRequest`、`ListProductsQuery`
- **AND** 必须从 `proto/psco/product_registry/v1/product_registry.proto` 单向派生或显式对齐
- **AND** 不得在 `types.go` 或 handler DTO 中新增 `.proto` 中不存在的业务字段语义
- **AND** `status` 语义必须继续复用 `active / archived`

### Requirement: Repository Binding 后端模块文件落点必须与既有主线同构

系统 SHALL 在 `backend/internal/repositorybinding/` 下实现 `Repository Binding` 后端模块，文件落点、分层语义与接线方式必须与既有后端主线同构，并承接 `phase04-07` 已冻结的模块边界。

#### Scenario: Repository Binding 文件落点冻结

- **WHEN** 实现 `Repository Binding` 后端模块
- **THEN** 必须按以下落点创建文件：
  - `handler/query_handler.go`
  - `handler/command_handler.go`
  - `handler/response.go`
  - `service/query_service.go`
  - `service/command_service.go`
  - `repository/repository_store.go`
  - `repository/binding_store.go`
  - `candidate/product_candidate_read.go`
  - `candidate/module_candidate_read.go`
  - `errors.go`
  - `types.go`
  - `validate.go`
- **AND** `repository/repository_store.go` 负责 `repositories` 本体读写
- **AND** `repository/binding_store.go` 负责 `product_repositories / module_repositories` 与详情摘要读取 SQL
- **AND** 两条候选读取必须继续落在独立 `candidate/` 子包，不并入 `RepositoryDetailRead`

#### Scenario: Repository Binding 合同承接边界

- **WHEN** 定义 `backend/internal/repositorybinding/types.go`
- **THEN** `Repository`、`RepositoryListItem`、`RepositoryDetail`、`BoundProductSummary`、`MappedModuleSummary`、`RepositoryProductCandidate`、`RepositoryModuleCandidate`、`CreateRepositoryRequest`、`CreateRepositoryResponse`、`BindRepositoryToProductRequest`、`MapModuleToRepositoryRequest`、`ListRepositoriesQuery`
- **AND** 必须从 `proto/psco/repository_binding/v1/repository_binding.proto` 单向派生或显式对齐
- **AND** 不得在 `types.go` 或 handler DTO 中新增 `.proto` 中不存在的业务字段语义
- **AND** `provider` 继续保持必填字符串，不升级为本阶段受控枚举

### Requirement: Query Service 必须承接已冻结的读取范围与 summary read 接线

系统 SHALL 在 `productregistry/service/query_service.go` 与 `repositorybinding/service/query_service.go` 中落实列表读取、详情读取、候选读取与跨模块摘要读取的编排边界，不得让消费方 service 直接越界写跨模块 SQL。

#### Scenario: Product Registry Query Service 读取边界

- **WHEN** 实现 `Product Registry` 的 `ListProducts`、`GetProductDetail`、`ListProductModuleCandidates`
- **THEN** `ListProducts` 必须承接 `name / description / status / created_at / module_bind_count / repository_bind_count`
- **AND** 筛选只允许 `queryText / statusFilter`
- **AND** 排序必须按 `created_at DESC`
- **AND** `GetProductDetail` 必须通过注入的 `ProductModuleSummaryRead` 与 `ProductRepositorySummaryRead` 拼装已绑定模块与已绑定仓库列表
- **AND** `service/` 层不得直接写跨模块 SQL 读取 `modules`、`repositories` 或关系表

#### Scenario: Repository Binding Query Service 读取边界

- **WHEN** 实现 `Repository Binding` 的 `ListRepositories`、`GetRepositoryDetail`、`ListRepositoryProductCandidates`、`ListRepositoryModuleCandidates`
- **THEN** `ListRepositories` 必须承接 `name / url / provider / status / created_at / product_bind_count / module_bind_count`
- **AND** 筛选只允许 `queryText / statusFilter`
- **AND** 排序必须按 `created_at DESC`
- **AND** `GetRepositoryDetail` 必须通过注入的 `RepositoryProductSummaryRead` 与 `RepositoryModuleSummaryRead` 拼装已绑定产品与已映射模块列表
- **AND** `service/` 层不得直接写跨模块 SQL 读取 `products`、`modules` 或关系表

### Requirement: Command Service 必须落实创建与三类绑定写入校验

系统 SHALL 在 `productregistry/service/command_service.go` 与 `repositorybinding/service/command_service.go` 中落实创建与绑定写入校验顺序，承接 `phase04-04` 已冻结的输入边界、错误语义与 reread owner。

#### Scenario: CreateProduct 与 CreateRepository 校验顺序

- **WHEN** 执行 `CreateProduct` 或 `CreateRepository`
- **THEN** 校验顺序必须至少为：
  1. 必填字段去首尾空白后非空
  2. `status` 只允许 `active / archived`
  3. `Product.name`、`Repository.name`、`Repository.url`、`Repository.provider` 的最小格式合法性
- **AND** 校验失败必须返回明确业务错误，不得降级为模糊 500
- **AND** 成功后必须返回新建实体 `id`，支撑前端回流到 `Product Detail` 或 `Repository Binding Detail / Workspace`

#### Scenario: BindModuleToProduct 校验顺序

- **WHEN** 执行 `BindModuleToProduct`
- **THEN** 校验顺序必须为：
  1. `product_id / module_id` 格式合法
  2. `Product` 存在且状态合法
  3. `Module` 存在且状态为 `active`
  4. 重复绑定检测
- **AND** 候选目标 `Module` 的存在性与 active 校验必须通过 candidate 子包三态返回（`exists / active / err`）分流：不存在返回 `ErrModuleNotFound`（404），存在但非 active 返回 `ErrModuleNotActive`（400）
- **AND** 成功后默认 reread owner 必须是 `ProductDetailRead`
- **AND** 不得返回脱离详情 reread 的第二套结果模型

#### Scenario: BindRepositoryToProduct 与 MapModuleToRepository 校验顺序

- **WHEN** 执行 `BindRepositoryToProduct` 或 `MapModuleToRepository`
- **THEN** 校验顺序必须分别覆盖：
  - `repository_id / product_id` 或 `repository_id / module_id` 格式合法
  - `Repository` 存在且状态合法
  - 目标 `Product` 或 `Module` 存在
  - 候选目标必须处于 `active`
  - 重复绑定/重复映射检测
- **AND** 候选目标的存在性与 active 校验必须通过 candidate 子包三态返回分流：不存在返回对应 `Err*NotFound`（404），存在但非 active 返回 `ErrProductNotActive` / `ErrModuleNotActive`（400）
- **AND** 成功后默认 reread owner 必须是 `RepositoryDetailRead`
- **AND** `ProductDetailRead` 不得成为 `BindRepositoryToProduct` 的第二 reread owner

### Requirement: 错误语义必须显式落到哨兵错误与 HTTP 状态码

系统 SHALL 为 `Product Registry` 与 `Repository Binding` 定义显式业务错误与 HTTP 状态码映射，承接 `phase04-04` 已冻结的错误语义前提。

#### Scenario: Product Registry 错误语义

- **WHEN** 实现 `productregistry/errors.go` 与 `handler/response.go`
- **THEN** 至少必须覆盖：
  - 资源不存在
  - 必填缺失/非法输入
  - 非法 `status`
  - 非法 UUID
  - 非 `active` 候选目标
  - 重复绑定
- **AND** 404 / 400 / 409 / 500 的映射必须单值稳定
- **AND** 候选空结果与列表空结果必须返回空列表，不得误报错误

#### Scenario: Repository Binding 错误语义

- **WHEN** 实现 `repositorybinding/errors.go` 与 `handler/response.go`
- **THEN** 至少必须覆盖：
  - 资源不存在
  - 必填缺失/非法输入
  - 非法 `status`
  - 非法 UUID
  - 非 `active` 候选目标
  - 重复绑定
  - 重复映射
- **AND** 404 / 400 / 409 / 500 的映射必须单值稳定
- **AND** 两条候选读取空结果必须返回空列表，不得误报错误

### Requirement: 0006 migration 必须将 products / repositories 升级为 phase04 正式主线

系统 SHALL 通过 `database/migrations/0006_product_repository_binding_mainline.sql` 将 `products / repositories` 从历史只读前提升级为 `phase04` 正式主线，并新增 `product_repositories` 关系表。

#### Scenario: Product 与 Repository 原位升级

- **WHEN** 执行 `0006_product_repository_binding_mainline.sql`
- **THEN** `products` 必须原位新增 `description / status`
- **AND** `repositories` 必须原位新增 `url / provider / status`
- **AND** `status` 必须使用 `CHECK (status IN ('active', 'archived'))`
- **AND** 列表读取索引必须分别覆盖 `(status, created_at DESC)`
- **AND** 不得创建 `products_v2`、`repositories_v2` 或影子表

#### Scenario: product_repositories 表结构

- **WHEN** `0006` 新增 `product_repositories`
- **THEN** 表结构至少必须包含 `id / product_id / repository_id / created_at`
- **AND** `(product_id, repository_id)` 必须唯一
- **AND** 必须新增按 `product_id` 与 `repository_id` 的读取索引
- **AND** 当前阶段不得引入额外绑定属性字段

#### Scenario: 历史数据兼容回填

- **WHEN** `0006` 执行时 `products / repositories` 已存在历史记录
- **THEN** `products.description` 必须回填为 `（历史产品，phase04 升级前无描述）`
- **AND** `repositories.url` 必须回填为 `https://example.com/legacy`
- **AND** `repositories.provider` 必须回填为 `legacy`
- **AND** 历史 `product_modules / module_repositories` 数据必须保持可读
- **AND** 回填语句必须幂等

### Requirement: phase04 基线 seed 与 reset 脚本必须可重复执行

系统 SHALL 实现 `phase04` 的基线 seed 与重置脚本，使联调环境不依赖手工补库。

#### Scenario: 基线 seed 落点与覆盖维度

- **WHEN** 实现 `database/seeds/seed_product_repository_mainline_baseline.sql`
- **THEN** 必须同时覆盖：
  - 至少 `3` 条 `Product`
  - 至少 `2` 条 `Repository`
  - 至少 `2` 条 `product_modules`
  - 至少 `1` 条 `product_repositories`
  - 至少 `2` 条 `module_repositories`
- **AND** 必须保留 `Product A / Product B / Product C / main-repo / mirror-repo`
- **AND** 所有关系恢复必须通过 `name` 查找，不得硬编码 UUID
- **AND** 至少保留 `1` 条无已绑定仓库的 `Product` 与 `1` 条无已绑定产品的 `Repository`

#### Scenario: reset_product_repository_mainline.sh 模式与清空范围

- **WHEN** 实现 `database/scripts/reset_product_repository_mainline.sh`
- **THEN** 必须支持 `--clean-only`、`--restore-only` 与默认模式
- **AND** 清空范围必须覆盖 `product_repositories / product_modules / module_repositories / products / repositories`
- **AND** 不得清空 `modules / decisions`
- **AND** 必须继续沿用 `DELETE FROM ...` 的受控清空模式，不切换为默认 `TRUNCATE ... CASCADE`
- **AND** 在恢复模式下必须先校验 `modules` 基线存在，否则提示先执行 `reset_module_mainline.sh`

#### Scenario: seed_readonly_prereqs 升级

- **WHEN** 继续通过 `run_seeds.sh` 建立历史兼容最小前提
- **THEN** `database/seeds/seed_readonly_prereqs.sql` 中的 `products / repositories` INSERT 必须升级为完整字段插入
- **AND** 必须继续保留历史 name 与 `ON CONFLICT (name) DO NOTHING`
- **AND** 其定位仍然只是“历史兼容最小前提”，不是 `phase04` 正式基线

### Requirement: phase02 旧 transport 入口若保留，必须只做兼容委派

系统 SHALL 将 `phase02` 遗留在 `Module Registry` 中的旧候选读取与旧绑定入口降级为兼容适配层；若这些入口在本阶段实现中继续保留，也只能委派给新的 canonical owner。

#### Scenario: 旧候选读取兼容委派

- **WHEN** 迁移窗口内仍保留 `ProductBindingCandidateRead` 或 `RepositoryBindingCandidateRead` 的旧 transport 入口
- **THEN** 它们只能分别委派到 `Repository Binding` 的 canonical 候选读取实现
- **AND** 不得在 `moduleregistry/service/` 内继续保留长期业务 owner 逻辑
- **AND** 返回数据不得落回旧字段语义或过时读取范围

#### Scenario: 旧模块中心绑定入口兼容委派

- **WHEN** 迁移窗口内仍保留 `ModuleBindingWrite` 旧模块中心写入口
- **THEN** `BindModuleToProduct` 必须委派到 `Product Registry` 的 `ProductModuleBindingWrite`
- **AND** `MapModuleToRepository` 必须委派到 `Repository Binding` 的 `RepositoryModuleMappingWrite`
- **AND** 旧入口本身不得独立写库
- **AND** 兼容委派不得制造第二套 reread owner 或第二套错误语义

### Requirement: 路由装配必须沿用 chi 子路由主线

系统 SHALL 继续沿用当前仓库已建立的 `chi` 子路由装配模式，将 `Product Registry` 与 `Repository Binding` 路由挂到 `/api` 下，而不新造第二套服务入口。

#### Scenario: platform 装配方式

- **WHEN** 实现 `backend/internal/platform/router.go` 与 `server.go`
- **THEN** 必须新增 `mountProductRegistry` 与 `mountRepositoryBinding`
- **AND** 两个模块都必须在 `r.Route("/api", ...)` 下装配
- **AND** 可继续使用 `chi.Route` / `chi.Mount` 的既有组织模式
- **AND** 路由矩阵必须对齐 `phase04-10` 与已落地 `.proto` 的 RPC -> HTTP 映射

## MODIFIED Requirements

### Requirement: products / repositories 从只读前提推进为正式主线实体

`products / repositories` SHALL 不再只被解释为 `phase02` 候选读取前提，而必须在 `phase04-12` 中推进为当前阶段正式读写主线实体。

#### Scenario: 主线解释切换

- **WHEN** `phase04-12` 实现完成
- **THEN** `CreateProduct / CreateRepository / List / Detail / 三类绑定` 必须直接以升级后的 `products / repositories / product_modules / product_repositories / module_repositories` 为事实源
- **AND** `phase02` 的只读前提语义只保留为历史兼容背景
- **AND** 不得继续依赖手工 SQL 补齐 `description / url / provider / status`

### Requirement: phase02 临时 owner 从“业务实现”推进为“兼容适配层”

`Module Registry` 中由 `phase02` 临时承接的候选读取与绑定写入 SHALL 从本阶段起只保留兼容适配语义，不再被解释为长期业务实现。

#### Scenario: 旧 owner 解释修正

- **WHEN** 仍需保留旧入口以服务迁移窗口
- **THEN** 必须将其理解为 transport compatibility shim
- **AND** 真实业务 owner 已迁移到 `Product Registry` 或 `Repository Binding`
- **AND** 迁移完成后旧入口应退出主线矩阵

## REMOVED Requirements

### Requirement: phase04 后端与联调继续依赖手工 SQL 或旧模块中心主写入

**Reason**: 这会让 `phase04` 无法形成可重复建立的后端/数据主线，也会让 `phase02` 临时 owner 持续污染当前阶段边界。

**Migration**: 后续实现、联调与验收统一从 `productregistry` / `repositorybinding` 模块、`0006_product_repository_binding_mainline.sql`、`seed_product_repository_mainline_baseline.sql` 与 `reset_product_repository_mainline.sh` 进入；旧 `Module Registry` 入口若保留，只作为兼容适配层委派到新的 canonical 实现。
