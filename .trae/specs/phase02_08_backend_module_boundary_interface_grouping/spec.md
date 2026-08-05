# Phase02-08 后端模块边界与接口分组设计 Spec

## Why

`phase02-05` 已经冻结了 `Module Registry` 的最小数据读写范围与接口前提，`phase02-06 / 07` 又把前端页面结构与交互流收敛到了可直接进入实现的粒度。当前仍缺少一份后端实现前设计结果，用来回答“哪些能力归 `Module Registry` 后端模块自己负责、哪些能力只通过连接边界承接、读写接口应该如何分组、这些接口落到 `backend/` 的哪些文件上”，否则后续实现仍会在模块归属、接口入口、跨模块调用与代码组织上发生漂移。

## What Changes

- 冻结 `Module Registry` 在后端的最小模块边界
- 冻结列表读取、详情读取、创建写入、版本写入、关联写入的接口分组
- 冻结 `Module Registry` 与 `Product / Repository / Decision` 的服务侧连接边界
- 冻结读接口、写接口与关系写入接口的职责拆分
- 冻结后端文件/目录落点到可直接创建文件的层级
- 明确当前阶段不提前冻结 Go HTTP/RPC 框架、数据访问层工具或完整合同文件结构

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `backend/` 中 `Module Registry` 的 handler/service/repository 边界、读写接口组织与跨模块连接方式
- Affected code: 预期会直接映射到 `backend/internal/moduleregistry/` 下的 handler、service、repository 与 candidate 文件落点

## ADDED Requirements

### Requirement: Module Registry 后端边界必须冻结为单一业务模块

系统 SHALL 将 `Module Registry` 在 `phase02` 的后端承接方式冻结为单一业务模块，由其统一承接当前阶段属于模块主线的读取、写入与关系写入能力。

#### Scenario: 当前阶段模块归属

- **WHEN** 后续实现定义 `Module Registry` 的后端模块边界
- **THEN** `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository` 必须归属于 `Module Registry` 后端模块
- **AND** 列表读取与详情读取也必须归属于该模块
- **AND** 不得把这些能力拆散到 `Product`、`Repository` 或 `Decision` 的后端模块中

#### Scenario: 当前阶段非归属能力

- **WHEN** 后续实现讨论 `RecordDecision`、`CreateProduct` 或 `CreateRepository`
- **THEN** 必须判定为当前阶段 `Module Registry` 后端模块的非归属能力
- **AND** `Module Registry` 只能通过最小连接边界读取或跳转这些能力

### Requirement: 后端接口必须按读组与写组拆分

系统 SHALL 将 `Module Registry` 的后端接口冻结为读组与写组两大类，避免把读取语义与写入语义混入同一入口。本 spec 的接口命名与分组沿用 `phase02-05` 已冻结的 `*Read` / `*Write` 命名体系，不引入第二套命名。

#### Scenario: 最小读接口分组

- **WHEN** 后续实现设计后端读取入口
- **THEN** 最小读接口分组至少应包含 `ModuleListRead` 与 `ModuleDetailRead`
- **AND** `ModuleListRead` 只承接列表展示所需最小读取
- **AND** `ModuleDetailRead` 只承接模块详情、版本列表、绑定关系与 `Decision` 入口所需最小读取

#### Scenario: 最小写接口分组

- **WHEN** 后续实现设计后端写入入口
- **THEN** 最小写接口分组至少应包含 `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite`
- **AND** 不得把创建、版本登记与关联写入混成单个无边界的大接口

### Requirement: Module List 读组必须只承接列表语义

系统 SHALL 将列表读组冻结为只服务 `Module Registry / List` 的读取接口，而不是提前膨胀成通用搜索中心。

#### Scenario: 列表读取职责

- **WHEN** 后续实现 `ModuleListRead`
- **THEN** 其职责只承接列表最小字段、筛选前提与进入详情所需标识
- **AND** 最小读取字段必须与 `phase02-02` 与 `phase02-04` 已冻结的列表读模型保持一致
- **AND** 当前阶段不提前扩写分页统计、Dashboard 聚合或跨对象分析读取

### Requirement: Module Detail 读组必须作为后端读模型宿主

系统 SHALL 将 `ModuleDetailRead` 冻结为 `Module Detail` 页面的统一后端读模型宿主，集中承接详情页当前阶段所有必要读取。

#### Scenario: 详情读取承接范围

- **WHEN** 后续实现 `ModuleDetailRead`
- **THEN** 它必须统一承接模块核心字段、版本列表、产品绑定、仓库映射与相关 `Decision` 入口读取
- **AND** 不得把这些详情读取拆成需要前端自行拼装的多个独立业务入口

#### Scenario: 详情读模型的连接边界

- **WHEN** `ModuleDetailRead` 需要承接 `Product`、`Repository` 或 `Decision` 关联信息
- **THEN** 只允许读取当前详情页确实需要的最小关联信息
- **AND** 不得顺势扩写为 `Product Registry`、`Repository Binding` 或 `Decision Center` 的全量读模型

### Requirement: 创建写组必须只负责 Module 自身登记

系统 SHALL 将 `ModuleCreateWrite` 冻结为只承接 `CreateModule` 的写入入口，不混入版本或关联写入。

#### Scenario: 创建写入职责

- **WHEN** 后续实现 `CreateModule`
- **THEN** `ModuleCreateWrite` 只负责模块自身最小字段写入
- **AND** 成功结果必须能返回新建模块标识，以支持前端回流到 `ModuleDetailPage`

### Requirement: 版本写组必须依附 Module 上下文

系统 SHALL 将 `ModuleReleaseWrite` 冻结为依附当前模块上下文的版本登记入口，而不是独立版本中心。

#### Scenario: 版本写入职责

- **WHEN** 后续实现 `CreateRelease`
- **THEN** `ModuleReleaseWrite` 必须以 `moduleId` 作为当前模块上下文前提
- **AND** 只负责当前模块下的版本登记
- **AND** 不得扩写为独立 `Release` 管理主线

### Requirement: 关联写组必须统一承接关系写入

系统 SHALL 将 `BindModuleToProduct` 与 `MapModuleToRepository` 冻结为统一的关系写入组，由 `ModuleBindingWrite` 承接。

#### Scenario: 关系写入职责

- **WHEN** 后续实现模块与产品、仓库的关联写入
- **THEN** `ModuleBindingWrite` 必须统一承接 `BindModuleToProduct` 与 `MapModuleToRepository`
- **AND** 这两个动作的拥有者仍然是 `Module Registry` 后端模块
- **AND** 不得把写入责任转交给 `Product` 或 `Repository` 模块

#### Scenario: 关系写入后的读取语义

- **WHEN** 任一关联写入成功
- **THEN** 当前阶段必须以“详情页重新读取当前绑定结果”为默认后端语义
- **AND** 不得额外设计一套脱离 `ModuleDetailRead` 的回流读取路径

### Requirement: Product 连接边界必须冻结为候选读取与关系校验

系统 SHALL 将 `Module Registry` 与 `Product` 的后端连接边界冻结为最小候选读取与关系写入校验，不承接产品主线写入。

#### Scenario: Product 最小连接边界

- **WHEN** `Module Registry` 需要支持 `BindModuleToProduct`
- **THEN** 它只允许依赖 `ProductBindingCandidateRead` 一类的最小候选读取接口
- **AND** 只允许依赖绑定所需的最小存在性与合法性校验
- **AND** 不得在当前阶段调用或吸收 `CreateProduct`、`UpdateProduct` 一类产品主线写入

### Requirement: Repository 连接边界必须冻结为候选读取与关系校验

系统 SHALL 将 `Module Registry` 与 `Repository` 的后端连接边界冻结为最小候选读取与关系写入校验，不承接仓库主线写入。

#### Scenario: Repository 最小连接边界

- **WHEN** `Module Registry` 需要支持 `MapModuleToRepository`
- **THEN** 它只允许依赖 `RepositoryBindingCandidateRead` 一类的最小候选读取接口
- **AND** 只允许依赖映射所需的最小存在性与合法性校验
- **AND** 不得在当前阶段调用或吸收 `CreateRepository`、仓库同步或仓库扫描类能力

### Requirement: Decision 连接边界必须冻结为 ModuleDetailRead 内嵌附属读取

系统 SHALL 将 `Module Registry` 与 `Decision` 的后端连接边界冻结为 `ModuleDetailRead` 的内嵌附属读取或跳转前提，不设独立读接口组，不扩写为独立决策写入主线。此约束沿用 `phase02-05` 已冻结的“`Decision` 入口作为 `ModuleDetailRead` 的附属读取承接，不设独立读接口组”单值结论，不在本 spec 引入第二套说法。

#### Scenario: Decision 最小连接边界

- **WHEN** `ModuleDetailRead` 需要展示相关 `Decision` 入口
- **THEN** `Decision` 关联信息必须作为 `ModuleDetailRead` 的内嵌附属读取承接，由读组业务编排层在组装详情读模型时一并读取
- **AND** 不得为 `Decision` 单独设立独立 Read 接口、独立 handler 或独立 service 文件
- **AND** 该读取只服务详情页的只读展示或跳转前提
- **AND** 不得在当前阶段新增 `RecordDecision` 或 `LinkDecisionToTarget` 的后端写入入口到 `Module Registry`

#### Scenario: Decision 读取的实现归属与可迁移性

- **WHEN** `phase02` 阶段实现 `ModuleDetailRead` 的 `Decision` 附属读取
- **THEN** 其数据组装逻辑物理归属于 `Module Registry` 后端模块的读组业务编排层
- **AND** `phase03` 实现 `Decision` 模块后，该数据组装逻辑可迁移为通过 `Decision` 模块的读模型协作获取，但 `ModuleDetailRead` 的外部接口契约保持不变
- **AND** 不得因为迁移而把 `Decision` 读取升级为独立读接口组

### Requirement: phase02 阶段跨模块候选读取必须通过接口边界临时承接

系统 SHALL 在 `phase02` 阶段，将 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead` 的实现物理归属于 `Module Registry` 后端模块，但必须通过独立 Read 接口边界与独立代码落点隔离，确保后续阶段可迁移到对应模块而不改变接口契约。`Decision` 读取不适用本节，因其已冻结为 `ModuleDetailRead` 内嵌附属读取，不构成独立候选读取接口。

#### Scenario: phase02 阶段候选读取实现归属

- **WHEN** `phase02` 阶段实现绑定动作
- **THEN** `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead` 的实现由 `Module Registry` 后端模块临时承接
- **AND** 必须通过独立 Read 接口定义与独立代码落点隔离，不得在写组业务编排层直接写跨模块 SQL
- **AND** `phase03 / phase04` 实现对应模块后，这些 Read 接口实现可迁移，但接口契约保持不变

#### Scenario: phase02 阶段候选读取的数据前提

- **WHEN** `phase02` 阶段需要读取 `Product / Repository` 候选
- **THEN** 对应表结构（`products / repositories`）在实现阶段必须在数据库中存在
- **AND** 这些表的结构定义属于实现阶段（`phase02-11` 数据主线）需要涵盖的只读前提，其运行时所需最小候选记录可通过最小种子数据或测试 fixture 提供以支撑 `phase02-12` 验收，但不视为 `phase02` 写入主线
- **AND** `phase02` 阶段这些表只支持只读候选读取，不要求完整写入主线
- **AND** 不得因为这些表存在而扩写 `phase02` 的写入范围

### Requirement: 后端文件落点必须冻结到实现结构层

系统 SHALL 将 `Module Registry` 的后端设计冻结到“包 + 文件落点”层级，确保后续实现可直接按既定结构创建文件，而不是由实现者临场决定代码组织方式。本 Requirement 冻结职责到文件的映射，但不冻结文件内部的具体 Go HTTP 框架、RPC 框架、ORM 或 SQL Builder 选型。

#### Scenario: 后端模块根包落点

- **WHEN** 后续实现开始创建 `Module Registry` 后端文件
- **THEN** 模块根包必须落在 `backend/internal/moduleregistry/`
- **AND** 模块内部必须按 `handler/`、`service/`、`repository/`、`candidate/` 四个子包组织
- **AND** 不得把 `Module Registry` 的代码散落到 `backend/` 根目录或其他模块目录中

#### Scenario: 读组文件落点

- **WHEN** 后续实现 `ModuleListRead` 与 `ModuleDetailRead`
- **THEN** 入口层必须落在 `backend/internal/moduleregistry/handler/query_handler.go`
- **AND** 业务编排层必须落在 `backend/internal/moduleregistry/service/query_service.go`
- **AND** `query_service.go` 必须统一承接列表读取与详情读取编排，`Decision` 附属读取作为详情读取的内嵌数据组装逻辑在此承接
- **AND** 不得把读组拆成独立 `list_service.go` 与 `detail_service.go`，保持读组单文件编排

#### Scenario: 写组文件落点

- **WHEN** 后续实现 `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite`
- **THEN** 入口层必须落在 `backend/internal/moduleregistry/handler/command_handler.go`
- **AND** 业务编排层必须落在 `backend/internal/moduleregistry/service/command_service.go`
- **AND** `command_service.go` 必须统一承接创建、版本与绑定三类写入编排
- **AND** 不得把写组拆成三个独立 service 文件，保持写组单文件编排

#### Scenario: 数据访问层文件落点

- **WHEN** 后续实现 `Module Registry` 的数据访问层
- **THEN** `modules` 表访问必须落在 `backend/internal/moduleregistry/repository/module_store.go`
- **AND** `module_releases` 表访问必须落在 `backend/internal/moduleregistry/repository/release_store.go`
- **AND** `product_modules` 与 `module_repositories` 表访问必须落在 `backend/internal/moduleregistry/repository/binding_store.go`
- **AND** 数据访问层只负责持久化与读取，不承接业务校验或跨模块编排

#### Scenario: 跨模块候选读取文件落点

- **WHEN** 后续实现 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead`
- **THEN** `ProductBindingCandidateRead` 必须落在 `backend/internal/moduleregistry/candidate/product_candidate_read.go`
- **AND** `RepositoryBindingCandidateRead` 必须落在 `backend/internal/moduleregistry/candidate/repository_candidate_read.go`
- **AND** `candidate/` 子包必须通过独立 Read 接口定义隔离，不得在 `service/` 中直接写跨模块 SQL
- **AND** `phase03 / phase04` 迁移时，这两个文件可整体迁移到对应模块，但接口契约保持不变

#### Scenario: Decision 读取不设独立文件落点

- **WHEN** 后续实现 `ModuleDetailRead` 的 `Decision` 附属读取
- **THEN** 不得为 `Decision` 读取单独创建 handler、service 或 repository 文件
- **AND** `Decision` 关联数据的读取逻辑必须内嵌在 `service/query_service.go` 的详情读取编排中
- **AND** 不得在 `candidate/` 子包中新增 `decision_*` 文件

#### Scenario: 当前阶段不冻结的实现工具

- **WHEN** 当前 spec 讨论后端文件落点
- **THEN** 不得提前冻结 Go HTTP 框架、RPC 框架、ORM、SQL Builder、Repository 模板或目录生成器
- **AND** 文件内部的函数签名、中间件选型、数据库访问方式不作为当前阶段冻结内容
- **AND** 只冻结职责分工、接口归属与文件落点，不冻结实现手段

### Requirement: 后端模块内部应保持分层而不提前钉死工具

系统 SHALL 将 `Module Registry` 的后端实现前结构冻结为“入口层（handler）-> 业务编排层（service）-> 数据访问/外部连接层（repository + candidate）”的分层语义，但不提前冻结具体框架与工具。

#### Scenario: 后端分层语义

- **WHEN** 后续实现开始定义 `Module Registry` 后端内部结构
- **THEN** 入口层（`handler/`）只负责承接请求与返回结果
- **AND** 业务编排层（`service/`）只负责动作语义、校验顺序与跨连接口编排
- **AND** 数据访问层（`repository/`）与外部连接层（`candidate/`）只负责持久化与依赖调用

#### Scenario: 当前阶段不得冻结的实现工具

- **WHEN** 当前 spec 讨论后端边界与接口分组
- **THEN** 不得提前冻结 Go HTTP 框架、RPC 框架、ORM、SQL Builder、Repository 模板或目录生成器
- **AND** 只冻结职责分工与接口归属，不冻结实现手段

### Requirement: 合同边界必须与存储模型解耦

系统 SHALL 将当前阶段的接口分组解释为业务合同边界，而不是数据库表结构的直接暴露。

#### Scenario: 合同与存储解耦

- **WHEN** 后续实现开始定义接口请求与响应结构
- **THEN** 不得直接将 `modules`、`module_releases`、`product_modules`、`module_repositories` 的存储模型原样暴露为外部合同
- **AND** 必须允许后续在 `Contract First` 路线下独立演进接口消息结构

#### Scenario: 合同演进前提

- **WHEN** 后续阶段需要调整接口消息结构
- **THEN** 不得复用已删除字段编号或字段语义
- **AND** 必须保持与 `Protocol Buffers` 长期方向兼容

## MODIFIED Requirements

### Requirement: Module Detail 后端承接方式

`Module Detail` 在当前阶段 SHALL 不仅是前端复合详情页，也必须在后端对应为一个统一读取宿主与动作回流宿主。

#### Scenario: Detail 后端收口

- **WHEN** 后续实现讨论 `Module Detail` 的后端承接方式
- **THEN** 必须保持 `ModuleDetailRead` 作为详情页统一读模型入口
- **AND** 版本登记与绑定动作成功后的默认后端回流语义也必须以该读组为中心
- **AND** `Decision` 附属读取必须内嵌于 `ModuleDetailRead`，不设独立读接口

## REMOVED Requirements

### Requirement: 当前阶段完整跨模块服务矩阵
**Reason**: `phase02-08` 的目标是冻结最小可执行后端边界，而不是提前完成 `Product / Repository / Decision` 全量服务规划。
**Migration**: 若后续需要完整跨模块服务矩阵，应在对应新 phase 或 audit 中单独冻结，不在当前阶段扩写。
