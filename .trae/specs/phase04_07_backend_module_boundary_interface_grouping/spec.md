# Phase04-07 后端模块边界与接口分组设计 Spec

## Why

`phase04-03` 已冻结三类绑定关系、候选读取归属、canonical owner 与 `phase02` 迁移边界，`phase04-04` 又冻结了最小数据读写范围、接口边界与错误语义前提，`phase04-05 / 06` 进一步把前端页面结构与状态流收敛到可直接进入实现的粒度。当前仍缺少一份后端实现前设计结果，用来回答“哪些能力归 `Product Registry`、哪些能力归 `Repository Binding`、读写与候选读取如何分组、旧 `phase02` 接口如何在迁移后保持兼容但不再保留第二 owner”。否则后续实现仍会在模块归属、handler/service 组织、旧接口迁移与 reread 回流上发生漂移。

## What Changes

- 冻结 `Product Registry` 与 `Repository Binding` 在后端的最小模块边界
- 冻结列表读取、详情读取、创建写入、绑定写入与候选读取的接口分组
- 冻结与 `Module Registry`、`Decision Center` 的服务侧连接边界
- 冻结三类绑定写入的 canonical owner、reread owner 与 `phase02` 旧接口兼容策略
- 冻结后端模块的最小包级落点与分层语义，确保后续实现可直接创建文件
- 明确当前阶段不提前冻结 Go 数据访问层具体工具、HTTP/RPC 框架与最终 `.proto` 命名

## Impact

- Affected specs: `phase04_product_and_repository_binding_foundation`
- Affected code: 后续 `backend/internal/productregistry/`、`backend/internal/repositorybinding/`、`backend/internal/platform/router.go`
- Affected code: 现有 `backend/internal/moduleregistry/` 中 `phase02` 临时承接的候选读取与绑定写入迁移

## ADDED Requirements

### Requirement: Product Registry 后端边界必须冻结为单一业务模块

系统 SHALL 将 `Product Registry` 在 `phase04` 的后端承接方式冻结为单一业务模块，由其统一承接当前阶段属于产品主线的读取、创建与 `BindModuleToProduct` 写入能力。

#### Scenario: Product Registry 当前阶段模块归属

- **WHEN** 后续实现定义 `Product Registry` 的后端模块边界
- **THEN** `ProductListRead`、`ProductDetailRead`、`ProductCreateWrite`、`ProductModuleCandidateRead`、`ProductModuleBindingWrite` 必须归属于 `Product Registry` 后端模块
- **AND** 不得把上述能力拆散到 `Repository Binding`、`Module Registry` 或 `Decision Center` 的后端模块中

#### Scenario: Product Registry 当前阶段非归属能力

- **WHEN** 后续实现讨论 `CreateRepository`、`BindRepositoryToProduct`、`MapModuleToRepository`
- **THEN** 必须判定为当前阶段 `Product Registry` 后端模块的非归属能力
- **AND** `Product Registry` 只允许通过最小读连接边界承接仓库绑定结果摘要，不吸收 `Repository Binding` 的主线写入

### Requirement: Repository Binding 后端边界必须冻结为单一业务模块

系统 SHALL 将 `Repository Binding` 在 `phase04` 的后端承接方式冻结为单一业务模块，由其统一承接当前阶段属于仓库主线的读取、创建与两类仓库绑定写入能力。

#### Scenario: Repository Binding 当前阶段模块归属

- **WHEN** 后续实现定义 `Repository Binding` 的后端模块边界
- **THEN** `RepositoryListRead`、`RepositoryDetailRead`、`RepositoryCreateWrite`、`ProductBindingCandidateRead`、`RepositoryModuleCandidateRead`、`RepositoryProductBindingWrite`、`RepositoryModuleMappingWrite` 必须归属于 `Repository Binding` 后端模块
- **AND** 不得把上述能力拆散到 `Product Registry`、`Module Registry` 或 `Decision Center` 的后端模块中

#### Scenario: Repository Binding 当前阶段非归属能力

- **WHEN** 后续实现讨论 `CreateProduct` 或 `BindModuleToProduct`
- **THEN** 必须判定为当前阶段 `Repository Binding` 后端模块的非归属能力
- **AND** `Repository Binding` 不得吸收 `Product Registry` 的产品主线写入

### Requirement: 后端接口必须按读组与写组拆分

系统 SHALL 将 `Product Registry` 与 `Repository Binding` 的后端接口冻结为读组与写组两大类，并在读组内部保留候选读取子组，避免把读取语义与写入语义混入同一入口。本规格沿用 `phase04-04` 已使用的 `*Read` / `*Write` 命名体系，不引入第二套命名。

#### Scenario: Product Registry 最小接口分组

- **WHEN** 后续实现设计 `Product Registry` 的后端接口分组
- **THEN** 读组至少必须包含 `ProductListRead`、`ProductDetailRead`、`ProductModuleCandidateRead`
- **AND** 写组至少必须包含 `ProductCreateWrite`、`ProductModuleBindingWrite`
- **AND** `ProductListRead` 只承接列表语义
- **AND** `ProductDetailRead` 只承接详情本体、已绑定模块列表与已绑定仓库摘要读取
- **AND** `ProductModuleCandidateRead` 只承接 `BindModuleToProduct` 的候选 `Module` 读取
- **AND** 不得把创建写入与绑定写入混成单个无边界的大接口

#### Scenario: Repository Binding 最小接口分组

- **WHEN** 后续实现设计 `Repository Binding` 的后端接口分组
- **THEN** 读组至少必须包含 `RepositoryListRead`、`RepositoryDetailRead`、`ProductBindingCandidateRead`、`RepositoryModuleCandidateRead`
- **AND** 写组至少必须包含 `RepositoryCreateWrite`、`RepositoryProductBindingWrite`、`RepositoryModuleMappingWrite`
- **AND** `RepositoryListRead` 只承接列表语义
- **AND** `RepositoryDetailRead` 只承接详情本体、已绑定产品列表与已映射模块列表读取
- **AND** `ProductBindingCandidateRead` 只承接 `BindRepositoryToProduct` 的候选 `Product` 读取
- **AND** `RepositoryModuleCandidateRead` 只承接 `MapModuleToRepository` 的候选 `Module` 读取

### Requirement: 方向级 API 矩阵冻结

系统 SHALL 将 `Product Registry` 与 `Repository Binding` 的方向级 API 矩阵冻结为单值结论，承接 `phase02-09` §6.2 与 `phase03-10` §6.2 已建立的方向级矩阵模式，使后续实现与 `phase04-08` `.proto` 合同设计可直接照此对齐。本矩阵与 `phase04-04` 已冻结的接口名一一对齐，不引入第二套命名。

#### Scenario: Product Registry 方向级 API 矩阵

- **WHEN** 后续实现或 `phase04-08` 合同设计讨论 `Product Registry` 的接口矩阵
- **THEN** 方向级 API 矩阵必须为：
  - 列表读取 → 读组 → `ProductListRead` → 读
  - 详情读取 → 读组 → `ProductDetailRead` → 读
  - `Module` 候选读取 → 候选读取 → `ProductModuleCandidateRead` → 读
  - `CreateProduct` → 写组 → `ProductCreateWrite` → 写
  - `BindModuleToProduct` → 写组 → `ProductModuleBindingWrite` → 写
- **AND** 不得在当前阶段引入超出该矩阵的并列接口
- **AND** 候选读取虽归读组语义，但必须通过独立 `candidate/` 子包拥有，不得并入 `ProductDetailRead`

#### Scenario: Repository Binding 方向级 API 矩阵

- **WHEN** 后续实现或 `phase04-08` 合同设计讨论 `Repository Binding` 的接口矩阵
- **THEN** 方向级 API 矩阵必须为：
  - 列表读取 → 读组 → `RepositoryListRead` → 读
  - 详情读取 → 读组 → `RepositoryDetailRead` → 读
  - `Product` 候选读取 → 候选读取 → `ProductBindingCandidateRead` → 读
  - `Module` 候选读取 → 候选读取 → `RepositoryModuleCandidateRead` → 读
  - `CreateRepository` → 写组 → `RepositoryCreateWrite` → 写
  - `BindRepositoryToProduct` → 写组 → `RepositoryProductBindingWrite` → 写
  - `MapModuleToRepository` → 写组 → `RepositoryModuleMappingWrite` → 写
- **AND** 不得在当前阶段引入超出该矩阵的并列接口
- **AND** 候选读取虽归读组语义，但必须通过独立 `candidate/` 子包拥有，不得并入 `RepositoryDetailRead`

### Requirement: Product Detail 与 Repository Detail 必须分别作为 reread owner

系统 SHALL 将三类绑定写入的默认 reread 承接方式冻结为对应详情读组重新读取，不再保留脱离详情读组的第二套回流读取路径。

#### Scenario: BindModuleToProduct 的 owner 与 reread owner

- **WHEN** 后续实现 `BindModuleToProduct`
- **THEN** `ProductModuleBindingWrite` 必须是唯一 canonical write owner
- **AND** `ProductDetailRead` 必须是成功写入后的唯一 reread owner
- **AND** 不得额外设计脱离 `ProductDetailRead` 的并列回流读取接口

#### Scenario: BindRepositoryToProduct 的 owner 与 reread owner

- **WHEN** 后续实现 `BindRepositoryToProduct`
- **THEN** `RepositoryProductBindingWrite` 必须是唯一 canonical write owner
- **AND** `RepositoryDetailRead` 必须是成功写入后的唯一 reread owner
- **AND** `ProductDetailRead` 只允许在用户回到产品详情时读取已更新结果，不得成为该动作的第二 reread owner

#### Scenario: MapModuleToRepository 的 owner 与 reread owner

- **WHEN** 后续实现 `MapModuleToRepository`
- **THEN** `RepositoryModuleMappingWrite` 必须是唯一 canonical write owner
- **AND** `RepositoryDetailRead` 必须是成功写入后的唯一 reread owner
- **AND** 不得额外设计脱离 `RepositoryDetailRead` 的并列回流读取接口

### Requirement: Module Registry 连接边界必须冻结为只读候选与存在性校验

系统 SHALL 将 `Product Registry` 与 `Repository Binding` 对 `Module Registry` 的服务侧连接边界冻结为最小候选读取、存在性校验与展示摘要读取，不承接任何 `Module Registry` 主线写入。

#### Scenario: Product Registry 到 Module Registry 的连接边界

- **WHEN** `Product Registry` 需要支持 `BindModuleToProduct`
- **THEN** 它只允许依赖 `ProductModuleCandidateRead` 一类的最小候选读取接口
- **AND** 只允许依赖 `module_id / module_name / module_status / created_at` 与存在性校验所需的最小读取
- **AND** 不得在当前阶段调用或吸收 `CreateModule`、`CreateRelease` 或 `Module Registry` 其他主线写入

#### Scenario: Repository Binding 到 Module Registry 的连接边界

- **WHEN** `Repository Binding` 需要支持 `MapModuleToRepository`
- **THEN** 它只允许依赖 `RepositoryModuleCandidateRead` 一类的最小候选读取接口
- **AND** 只允许依赖 `module_id / module_name / module_status / created_at` 与存在性校验所需的最小读取
- **AND** 不得在当前阶段调用或吸收 `CreateModule`、`CreateRelease` 或 `BindModuleToProduct`

### Requirement: 跨模块已绑定关系摘要读取边界冻结

系统 SHALL 将 `ProductDetailRead` 与 `RepositoryDetailRead` 中“已绑定关系列表”的跨模块读取边界冻结为单值结论，承接 `phase04-04` 详情读取数据范围冻结（`ProductDetailRead` 必须承接已绑定模块列表与已绑定仓库列表，`RepositoryDetailRead` 必须承接已绑定产品列表与已映射模块列表）与 `phase04-03` 关系表数据访问 owner 冻结（`product_modules` 归属 `Product Registry`，`product_repositories` / `module_repositories` 归属 `Repository Binding`），避免实现阶段在“直接越界读跨模块关系表 / 新造未冻结接口 / 复制第二套关系读取 owner”之间自行选择。

> owner 确定原则：跨模块关系摘要读取接口的 owner = 关系表 owner。`product_repositories` 与 `module_repositories` 的 owner 为 `Repository Binding`，`product_modules` 的 owner 为 `Product Registry`。展示字段需要跨模块读取实体表时（如 `modules` / `products`），由 owner 模块在其 `repository/binding_store.go` 中通过 SQL join 承接，与候选读取的“消费方拥有跨模块读”模式在 owner 确定原则上一致，差别仅在文件落点（关系摘要读取落 `repository/binding_store.go`，候选读取落 `candidate/` 子包）。

#### Scenario: 判断跨模块读接口 owner 与接口形态

- **WHEN** `ProductDetailRead` 需要承接“已绑定仓库列表”
- **THEN** 必须通过 `Repository Binding` 模块拥有的独立跨模块读接口 `ProductRepositorySummaryRead` 承接
- **AND** `ProductRepositorySummaryRead` 的语义为“给定 `productId`，返回已绑定的 `Repository` 摘要列表”
- **AND** 每条摘要至少包含 `repository_id / repository_name / provider / repository_status`，承接 `phase04-04` 详情读取数据范围
- **AND** `Product Registry` 模块不得直接读 `product_repositories` 表
- **AND** 不得新造未冻结的跨模块读接口或复制第二套关系读取 owner

#### Scenario: 判断跨模块读接口包级归属与文件落点

- **WHEN** 后续实现组织 `ProductRepositorySummaryRead` 的代码落点
- **THEN** 接口定义必须落在 `Repository Binding` 模块的 `types.go`（模块层级，不暴露 `repository/` 子包）
- **AND** 实现必须落在 `Repository Binding` 模块的 `repository/binding_store.go`（`product_repositories` 表数据访问本来就在此文件）
- **AND** 不得把该接口放入 `candidate/` 子包（`candidate/` 语义为候选读取，不承接关系摘要读取）
- **AND** 不得在 `Product Registry` 模块内复制该接口的第二套实现

#### Scenario: 判断跨模块读接口装配方式与 service 约束

- **WHEN** 后续实现装配 `ProductDetailRead` 的跨模块仓库摘要读取
- **THEN** 具体接线（构造与注入）必须在应用装配点（如 `backend/internal/platform/`）完成
- **AND** `Product Registry` 的 `service/query_service.go` 必须通过注入的 `ProductRepositorySummaryRead` 接口调用，不得在 `service/` 或 `handler/` 内部自行构造
- **AND** `Product Registry` 的 `service/` 层不得直接写跨模块 SQL 读取 `product_repositories` 表
- **AND** 该接线原则与候选读取接线原则一致，承接 `phase03-10` §10.5 与 `phase04-03` 候选读取接口归属冻结结论

#### Scenario: 判断跨模块读接口与候选读取的关系

- **WHEN** 后续实现讨论 `Repository Binding` 模块的跨模块读边界
- **THEN** `ProductRepositorySummaryRead`（关系摘要读取）与 `ProductBindingCandidateRead` / `RepositoryModuleCandidateRead`（候选读取）是两类不同的跨模块读边界
- **AND** 两类读边界必须通过独立 Read 接口隔离，不得混用或合并
- **AND** `ProductRepositorySummaryRead` 服务 `ProductDetailRead` 的详情展示，不服务绑定动作
- **AND** 候选读取服务绑定动作的候选选择，不服务详情展示

#### Scenario: ProductDetailRead 的“已绑定模块列表”跨模块读取边界

- **WHEN** `ProductDetailRead` 需要承接“已绑定模块列表”
- **THEN** 必须通过 `Product Registry` 模块拥有的独立跨模块读接口 `ProductModuleSummaryRead` 承接
- **AND** `ProductModuleSummaryRead` 的 owner 为 `Product Registry`，因为 `product_modules` 关系表 owner 已在 `phase04-03` 冻结为 `Product Registry`
- **AND** `ProductModuleSummaryRead` 的语义为“给定 `productId`，返回已绑定的 `Module` 摘要列表”
- **AND** 每条摘要至少包含 `module_id / module_name / module_status`，承接 `phase04-04` `ProductDetailRead` 详情读取数据范围
- **AND** 接口定义必须落在 `backend/internal/productregistry/types.go`（模块层级，不暴露 `repository/` 子包）
- **AND** 实现必须落在 `backend/internal/productregistry/repository/binding_store.go`（`product_modules` 表数据访问本来就在此文件，展示字段 `module_name / module_status` 通过 SQL join `modules` 表承接，与候选读取的“消费方拥有跨模块读”模式一致，差别仅在文件落点）
- **AND** 不得把该接口放入 `candidate/` 子包（`candidate/` 语义为候选读取，不承接关系摘要读取）
- **AND** 具体接线（构造与注入）必须在应用装配点（如 `backend/internal/platform/`）完成，`Product Registry` 的 `service/query_service.go` 必须通过注入的 `ProductModuleSummaryRead` 接口调用，不得在 `service/` 或 `handler/` 内部自行构造
- **AND** `Product Registry` 的 `service/` 层不得直接写跨模块 SQL 读取 `modules` 表，承接本规格候选读取接线原则通用规则
- **AND** 该接口为关系摘要读取（服务详情展示），与 `ProductModuleCandidateRead`（候选读取，服务绑定动作）是两类不同的跨模块读边界，必须通过独立 Read 接口隔离
- **AND** 不得新造未冻结的跨模块读接口或复制第二套关系读取 owner

#### Scenario: RepositoryDetailRead 的“已绑定产品列表”跨模块读取边界

- **WHEN** `RepositoryDetailRead` 需要承接“已绑定产品列表”
- **THEN** 必须通过 `Repository Binding` 模块拥有的独立跨模块读接口 `RepositoryProductSummaryRead` 承接
- **AND** `RepositoryProductSummaryRead` 的 owner 为 `Repository Binding`，因为 `product_repositories` 关系表 owner 已在 `phase04-03` 冻结为 `Repository Binding`
- **AND** `RepositoryProductSummaryRead` 的语义为“给定 `repositoryId`，返回已绑定的 `Product` 摘要列表”
- **AND** 每条摘要至少包含 `product_id / product_name / product_status`，承接 `phase04-04` `RepositoryDetailRead` 详情读取数据范围
- **AND** 接口定义必须落在 `backend/internal/repositorybinding/types.go`（模块层级，不暴露 `repository/` 子包）
- **AND** 实现必须落在 `backend/internal/repositorybinding/repository/binding_store.go`（`product_repositories` 表数据访问本来就在此文件，展示字段 `product_name / product_status` 通过 SQL join `products` 表承接，与候选读取的“消费方拥有跨模块读”模式一致，差别仅在文件落点）
- **AND** 不得把该接口放入 `candidate/` 子包（`candidate/` 语义为候选读取，不承接关系摘要读取）
- **AND** 具体接线（构造与注入）必须在应用装配点（如 `backend/internal/platform/`）完成，`Repository Binding` 的 `service/query_service.go` 必须通过注入的 `RepositoryProductSummaryRead` 接口调用，不得在 `service/` 或 `handler/` 内部自行构造
- **AND** `Repository Binding` 的 `service/` 层不得直接写跨模块 SQL 读取 `products` 表，承接本规格候选读取接线原则通用规则
- **AND** 该接口为关系摘要读取（服务详情展示），与 `ProductBindingCandidateRead`（候选读取，服务绑定动作）是两类不同的跨模块读边界，必须通过独立 Read 接口隔离
- **AND** 不得新造未冻结的跨模块读接口或复制第二套关系读取 owner

#### Scenario: RepositoryDetailRead 的“已映射模块列表”跨模块读取边界

- **WHEN** `RepositoryDetailRead` 需要承接“已映射模块列表”
- **THEN** 必须通过 `Repository Binding` 模块拥有的独立跨模块读接口 `RepositoryModuleSummaryRead` 承接
- **AND** `RepositoryModuleSummaryRead` 的 owner 为 `Repository Binding`，因为 `module_repositories` 关系表 owner 已在 `phase04-03` 冻结为 `Repository Binding`
- **AND** `RepositoryModuleSummaryRead` 的语义为“给定 `repositoryId`，返回已映射的 `Module` 摘要列表”
- **AND** 每条摘要至少包含 `module_id / module_name / module_status`，承接 `phase04-04` `RepositoryDetailRead` 详情读取数据范围
- **AND** 接口定义必须落在 `backend/internal/repositorybinding/types.go`（模块层级，不暴露 `repository/` 子包）
- **AND** 实现必须落在 `backend/internal/repositorybinding/repository/binding_store.go`（`module_repositories` 表数据访问本来就在此文件，展示字段 `module_name / module_status` 通过 SQL join `modules` 表承接，与候选读取的“消费方拥有跨模块读”模式一致，差别仅在文件落点）
- **AND** 不得把该接口放入 `candidate/` 子包（`candidate/` 语义为候选读取，不承接关系摘要读取）
- **AND** 具体接线（构造与注入）必须在应用装配点（如 `backend/internal/platform/`）完成，`Repository Binding` 的 `service/query_service.go` 必须通过注入的 `RepositoryModuleSummaryRead` 接口调用，不得在 `service/` 或 `handler/` 内部自行构造
- **AND** `Repository Binding` 的 `service/` 层不得直接写跨模块 SQL 读取 `modules` 表，承接本规格候选读取接线原则通用规则
- **AND** 该接口为关系摘要读取（服务详情展示），与 `RepositoryModuleCandidateRead`（候选读取，服务绑定动作）是两类不同的跨模块读边界，必须通过独立 Read 接口隔离
- **AND** 不得新造未冻结的跨模块读接口或复制第二套关系读取 owner

#### Scenario: 判断跨模块关系摘要读取接口的整体边界隔离

- **WHEN** 后续实现讨论跨模块关系摘要读取与候选读取的整体关系
- **THEN** `ProductRepositorySummaryRead` / `ProductModuleSummaryRead` / `RepositoryProductSummaryRead` / `RepositoryModuleSummaryRead`（四条关系摘要读取链路）与 `ProductModuleCandidateRead` / `ProductBindingCandidateRead` / `RepositoryModuleCandidateRead`（三条候选读取链路）是两类不同的跨模块读边界
- **AND** 两类读边界必须通过独立 Read 接口隔离，不得混用或合并
- **AND** 关系摘要读取服务详情展示，不服务绑定动作
- **AND** 候选读取服务绑定动作的候选选择，不服务详情展示
- **AND** 不得把任一关系摘要读取接口实现为消费方模块 `service/` 层直接越界读跨模块关系表或实体表

### Requirement: Decision Center 连接边界必须保持后移

系统 SHALL 将 `Product Registry` 与 `Repository Binding` 对 `Decision Center` 的服务侧连接边界冻结为“当前阶段不接入正式主线”，不得把 `Decision Center` 扩写为并列绑定主线或强制依赖。

#### Scenario: 当前阶段不接入 Decision Center 主线

- **WHEN** 后续实现讨论 `Product Registry`、`Repository Binding` 与 `Decision Center` 的后端连接
- **THEN** 不得为 `Decision -> Product / Repository` 设计正式候选读取、详情读取或绑定写入接口
- **AND** 不得让 `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 依赖 `Decision Center` 成为前置主线
- **AND** 若后续阶段需要 `Decision` 与 `Product / Repository` 建立正式连接，必须进入新的冻结任务单独收敛

### Requirement: phase02 临时承接接口必须迁移为 canonical owner 下的兼容适配

系统 SHALL 将 `phase02` 中由 `Module Registry` 临时承接的候选读取与绑定写入接口解释为历史兼容入口，而不是长期 owner。迁移完成后只允许保留“兼容适配层”，不允许保留第二套业务 owner。

#### Scenario: ProductBindingCandidateRead 的兼容策略

- **WHEN** 迁移 `phase02` 的 `ProductBindingCandidateRead`
- **THEN** canonical owner 必须切换为 `Repository Binding` 读组
- **AND** 若迁移窗口内仍需保留旧 transport 入口，它也只能作为兼容适配层委派给 `Repository Binding` 的 canonical 实现
- **AND** 不得继续由 `Module Registry` 的业务 service 保留长期 owner 身份

#### Scenario: RepositoryBindingCandidateRead 的兼容策略

- **WHEN** 迁移 `phase02` 的 `RepositoryBindingCandidateRead`
- **THEN** 必须判定为已废弃接口
- **AND** 不得在 `phase04` 中保留新的 canonical 对应接口
- **AND** 若迁移窗口内仍保留旧 transport 入口，也不得再承接新的业务实现或新增消费方

#### Scenario: ModuleBindingWrite 的兼容策略

- **WHEN** 迁移 `phase02` 的 `ModuleBindingWrite`
- **THEN** `BindModuleToProduct` 必须迁移为 `Product Registry` 的 `ProductModuleBindingWrite`
- **AND** `MapModuleToRepository` 必须迁移为 `Repository Binding` 的 `RepositoryModuleMappingWrite`
- **AND** 旧的模块中心写入入口若在迁移窗口内暂时保留，也只能作为兼容适配层委派给新的 canonical write owner
- **AND** 不得继续把 `Module Registry` 解释为这两类写入的长期 owner

### Requirement: phase02 历史绑定数据兼容前提冻结

系统 SHALL 将 `phase02` 中已存在的 `product_modules` 与 `module_repositories` 历史绑定数据在 `phase04` 迁移后的兼容前提冻结为单值结论，承接 `phase04-03` 历史绑定数据兼容冻结，确保迁移不破坏既有数据可读性。

#### Scenario: 历史绑定数据可读性

- **WHEN** `phase04` 完成后端模块迁移后
- **THEN** `phase02` 中已存在的 `product_modules` 历史绑定数据必须保持可读
- **AND** `phase02` 中已存在的 `module_repositories` 历史绑定数据必须保持可读
- **AND** 不得通过重建影子表、第二套绑定表或临时双写绕过迁移
- **AND** 既有前端入口升级后必须仍可继续读取历史绑定结果
- **AND** 数据访问 owner 切换到 `Product Registry` 与 `Repository Binding` 后，表结构与历史数据保持不变

#### Scenario: Module Detail 旧入口后端 endpoint 兼容策略

- **WHEN** 后续实现讨论 `Module Detail` 既有绑定入口的后端对应
- **THEN** `phase04` 后端不得为 `Module Detail` 提供新的绑定写入 endpoint
- **AND** `phase02` 中由 `Module Registry` 承接的旧绑定写入 endpoint 若迁移窗口内保留，只能作为兼容适配层委派给 `Product Registry` 的 `ProductModuleBindingWrite` 或 `Repository Binding` 的 `RepositoryModuleMappingWrite`
- **AND** 不得在 `Module Registry` 业务 service 中继续保留绑定写入的长期 owner 实现
- **AND** 该兼容策略承接 `phase04-03` `Module Detail` 旧入口兼容跳转冻结结论，前后端口径一致

### Requirement: 后端模块文件落点必须冻结到实现结构层

系统 SHALL 将 `Product Registry` 与 `Repository Binding` 的后端设计冻结到“包 + 关键文件落点”层级，确保后续实现可直接按既定结构创建文件，而不是由实现者临场决定代码组织方式。本 Requirement 只冻结职责到文件的映射，不冻结文件内部的具体 Go 数据访问层工具、HTTP/RPC 框架或最终合同消息命名。

#### Scenario: Product Registry 文件落点

- **WHEN** 后续实现开始创建 `Product Registry` 后端文件
- **THEN** 模块根包必须落在 `backend/internal/productregistry/`
- **AND** 模块内部必须按 `handler/`、`service/`、`repository/`、`candidate/` 四个子包组织
- **AND** 读组入口必须落在 `backend/internal/productregistry/handler/query_handler.go`
- **AND** 写组入口必须落在 `backend/internal/productregistry/handler/command_handler.go`
- **AND** 读组业务编排必须落在 `backend/internal/productregistry/service/query_service.go`，其中 `ProductDetailRead` 的“已绑定模块列表”必须通过注入的 `ProductModuleSummaryRead` 接口承接、“已绑定仓库列表”必须通过注入的 `ProductRepositorySummaryRead` 接口承接，不得直接读 `product_modules` 或 `product_repositories` 表
- **AND** 写组业务编排必须落在 `backend/internal/productregistry/service/command_service.go`
- **AND** `products` 表访问必须落在 `backend/internal/productregistry/repository/product_store.go`
- **AND** `product_modules` 关系写入、产品侧关系读取与 `ProductModuleSummaryRead` 跨模块读接口实现必须落在 `backend/internal/productregistry/repository/binding_store.go`
- **AND** `ProductModuleCandidateRead` 必须落在 `backend/internal/productregistry/candidate/module_candidate_read.go`

#### Scenario: Repository Binding 文件落点

- **WHEN** 后续实现开始创建 `Repository Binding` 后端文件
- **THEN** 模块根包必须落在 `backend/internal/repositorybinding/`
- **AND** 模块内部必须按 `handler/`、`service/`、`repository/`、`candidate/` 四个子包组织
- **AND** 读组入口必须落在 `backend/internal/repositorybinding/handler/query_handler.go`
- **AND** 写组入口必须落在 `backend/internal/repositorybinding/handler/command_handler.go`
- **AND** 读组业务编排必须落在 `backend/internal/repositorybinding/service/query_service.go`，其中 `RepositoryDetailRead` 的“已绑定产品列表”必须通过注入的 `RepositoryProductSummaryRead` 接口承接、“已映射模块列表”必须通过注入的 `RepositoryModuleSummaryRead` 接口承接，不得直接读 `product_repositories` 或 `module_repositories` 表
- **AND** 写组业务编排必须落在 `backend/internal/repositorybinding/service/command_service.go`
- **AND** `repositories` 表访问必须落在 `backend/internal/repositorybinding/repository/repository_store.go`
- **AND** `product_repositories` 与 `module_repositories` 关系写入、仓库侧关系读取与 `ProductRepositorySummaryRead` / `RepositoryProductSummaryRead` / `RepositoryModuleSummaryRead` 三个跨模块读接口实现必须落在 `backend/internal/repositorybinding/repository/binding_store.go`
- **AND** `ProductBindingCandidateRead` 必须落在 `backend/internal/repositorybinding/candidate/product_candidate_read.go`
- **AND** `RepositoryModuleCandidateRead` 必须落在 `backend/internal/repositorybinding/candidate/module_candidate_read.go`

#### Scenario: 分层语义与单文件编排原则

- **WHEN** 后续实现组织 `Product Registry` 与 `Repository Binding` 的内部分层
- **THEN** 分层语义必须为：
  - 入口层 → `handler/` → 只负责承接请求与返回结果
  - 业务编排层 → `service/` → 只负责动作语义、校验顺序与跨连接口编排
  - 数据访问/外部连接层 → `repository/` + `candidate/` → 只负责持久化与跨模块依赖调用
- **AND** 读组编排必须落在单文件 `service/query_service.go`，不得拆为 `list_service.go` / `detail_service.go` / `candidate_service.go`
- **AND** 写组编排必须落在单文件 `service/command_service.go`，不得按动作拆为多个独立 service 文件
- **AND** 该单文件编排原则承接 `phase02-09` §9.4 与 `phase03-10` §10.4 已建立模式

#### Scenario: 支撑文件落点

- **WHEN** 后续实现创建 `Product Registry` 与 `Repository Binding` 的支撑文件
- **THEN** `Product Registry` 支撑文件必须落在 `backend/internal/productregistry/`：
  - `errors.go` → 业务错误哨兵值
  - `types.go` → 跨层共享 API 消息结构（DTO、枚举、请求/响应类型、列表查询参数）
  - `validate.go` → 输入校验辅助
  - `handler/response.go` → HTTP 协议层共享工具（JSON 编解码、错误到状态码映射）
- **AND** `Repository Binding` 支撑文件必须落在 `backend/internal/repositorybinding/`，文件名与职责同上
- **AND** 支撑文件组织方式必须与现有 `moduleregistry` 与 `decisioncenter` 模块同构，承接 `phase03-10` §10.6 已建立模式（`phase02-09` §9.2 已建立 `handler/` / `service/` / `repository/` / `candidate/` 子包组织模式，支撑文件落点由 `phase03-10` §10.6 显式冻结）
- **AND** 不得把错误定义、共享类型或校验逻辑散落到 `handler/` 或 `service/` 内部

#### Scenario: 候选读取接口拥有者与接线原则

- **WHEN** 后续实现 `Product Registry` 与 `Repository Binding` 的候选读取
- **THEN** `ProductModuleCandidateRead` 的接口与实现必须由 `Product Registry` 的 `candidate/` 子包自己定义和拥有
- **AND** `ProductBindingCandidateRead` 与 `RepositoryModuleCandidateRead` 的接口与实现必须由 `Repository Binding` 的 `candidate/` 子包自己定义和拥有
- **AND** `Module Registry` 不为 `Product Registry` 或 `Repository Binding` 暴露专门的服务契约或服务方法
- **AND** `candidate/` 子包通过独立 Read 接口隔离，`service/` 层不得直接写跨模块 SQL 读取 `modules` 或 `products` 表
- **AND** 具体接线（构造与注入）必须在应用装配点（如 `backend/internal/platform/`）完成，不得在 `service/` 或 `handler/` 内部自行构造
- **AND** 该接线原则承接 `phase03-10` §10.5 与 `phase04-03` 候选读取接口归属冻结结论
- **AND** 该接线原则同样适用于 `ProductRepositorySummaryRead` / `ProductModuleSummaryRead` / `RepositoryProductSummaryRead` / `RepositoryModuleSummaryRead` 四个跨模块关系摘要读取接口，详见「跨模块已绑定关系摘要读取边界冻结」Requirement

#### Scenario: product_repositories 表为 phase04 新增

- **WHEN** 后续实现讨论 `product_repositories` 表的来源
- **THEN** `product_repositories` 表必须理解为 `phase04` 阶段新增表，承接 `BindRepositoryToProduct` 关系
- **AND** 该表不存在 `phase02` 历史数据，不需要兼容回填
- **AND** 该表的 migration 落地由后续实现任务承接，本规格只冻结其归属 `Repository Binding` 模块的 `repository/binding_store.go`
- **AND** `product_modules` 与 `module_repositories` 表已存在于 `phase02` migration 中，迁移后表结构保持不变，只是数据访问 owner 切换

#### Scenario: 当前阶段不冻结的实现工具

- **WHEN** 当前规格讨论后端模块边界与接口分组
- **THEN** 不得提前冻结 Go 数据访问层具体工具、ORM、SQL Builder、Repository 模板或目录生成器
- **AND** 不得提前冻结 HTTP/RPC 框架或最终 `.proto` 消息命名
- **AND** 只冻结职责分工、接口归属、兼容迁移与文件落点，不冻结实现手段

### Requirement: Chi 路由组织方式只作为实现兼容前提

系统 SHALL 将当前仓库已使用的 `chi` 子路由组织方式视为实现兼容前提，而不是新的合同源；后续实现可以继续通过 `Route / Mount` 组织模块级子路由，但接口 owner 与分组必须以本规格为准。

#### Scenario: 现有 router 与模块子路由兼容

- **WHEN** 后续实现在 `backend/internal/platform/router.go` 接入 `Product Registry` 与 `Repository Binding`
- **THEN** 可以沿用当前 `chi` 的模块子路由组织方式进行装配
- **AND** 不得把 `router.go` 中的装配便利性反向解释为跨模块业务 owner
- **AND** canonical owner、读写分组与兼容适配策略仍必须以本规格冻结结论为准

## MODIFIED Requirements

### Requirement: phase02 临时绑定 owner 的解释

`phase02` 中以 `Module Registry` 为中心的候选读取与绑定写入 SHALL 不再被解释为当前阶段长期后端 owner，而必须被解释为具有明确迁移目标的历史兼容入口。

#### Scenario: 历史 owner 解释修正

- **WHEN** 后续 `/spec` 或实现讨论 `phase02` 遗留后端接口的 owner
- **THEN** 必须将其理解为“旧 transport 入口最多保留兼容适配层，业务 owner 已迁移到 `Product Registry` 或 `Repository Binding`”
- **AND** 不得继续把 `Module Registry` 实现为当前阶段并列绑定主线

## REMOVED Requirements

### Requirement: Module Registry 继续作为绑定写入长期 owner

**Reason**: `phase04-03` 已将三类绑定动作的 canonical owner 冻结为 `Product Detail` 与 `Repository Binding Detail / Workspace`，后端也必须随之迁移到 `Product Registry` 与 `Repository Binding` 两个模块，不再允许 `Module Registry` 长期保留对应写入 owner。

**Migration**: 迁移窗口内若仍需保留旧模块中心接口，只允许在 transport 层作为兼容适配入口委派到新的 canonical service；迁移完成后旧接口应退出主线矩阵。

### Requirement: Decision Center 作为 Product / Repository 绑定主线参与者

**Reason**: 当前阶段 `Decision Center` 不在 `Product / Repository / Binding` 的正式主线范围内，提前引入会扩写范围并破坏既有单值主线。

**Migration**: 若后续确需建立 `Decision -> Product / Repository` 正式后端连接，必须进入新的冻结任务单独收敛，不在 `phase04-07` 当前规格中处理。
