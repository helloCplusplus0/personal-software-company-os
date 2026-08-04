# Phase02-05 数据读写范围与 API 承接前提 Spec

## Why

`phase02-01 ~ 04` 已经把页面边界、实体主线、创建闭环、版本与绑定动作归属收敛到可执行状态。为了让后续源码实现设计与正式规格正文能够直接落地，当前还需要把 `Module Registry` 所需的数据读写范围、最小接口承接前提、读写接口分组，以及 `Decision` 在当前阶段的接口边界冻结为单值结论。

## What Changes

- 冻结 `Module Registry` 当前阶段所需的数据读写范围
- 冻结最小接口承接前提与读写接口分组
- 冻结列表读取、详情读取、创建写入、版本写入、关联写入各自的最小接口边界
- 冻结 `Decision` 在当前阶段只读/跳转而不扩写为独立写主线的接口边界
- 明确当前阶段不提前冻结完整查询矩阵或 `Dashboard` 聚合接口

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `backend/` 中 `Module Registry` 读写接口分组、`frontend/` 中列表/详情/创建/版本/关联动作的数据请求承接

## ADDED Requirements

### Requirement: Module Registry 数据读写范围冻结

系统 SHALL 将 `phase02` 当前阶段 `Module Registry` 所需的数据读写范围冻结为最小可执行集合，只承接当前主线真正需要的数据对象。

#### Scenario: 当前阶段写入范围

- **WHEN** 后续 `/spec` 或实现讨论 `Module Registry` 的写入范围
- **THEN** 当前阶段只承接 `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository`
- **AND** 写入数据范围只对应 `modules`、`module_releases`、`product_modules`、`module_repositories`

#### Scenario: 当前阶段读取范围

- **WHEN** 后续 `/spec` 或实现讨论 `Module Registry` 的读取范围
- **THEN** 当前阶段只承接列表读取、详情读取与相关只读关联入口
- **AND** 读取前提只包含 `modules`、`module_releases`、`product_modules`、`module_repositories`、`decisions`、`decision_links`

### Requirement: 最小接口承接前提冻结

系统 SHALL 将 `Module Registry` 的最小接口承接前提冻结为“按当前页面主线承接读写动作”，而不是提前展开完整服务矩阵。

#### Scenario: 页面到接口的最小映射

- **WHEN** 后续 `/spec` 或实现需要定义最小接口前提
- **THEN** `Module Registry / List` 必须承接列表读取接口
- **AND** `Module Detail` 必须承接详情读取、版本写入与关联写入接口
- **AND** `Module Create` 必须承接创建写入接口
- **AND** 不得将这些动作提前拆成独立主线服务入口

### Requirement: 读动作接口分组冻结

系统 SHALL 将当前阶段读动作接口分为最小读取组，而不是提前冻结完整查询矩阵。

#### Scenario: 读动作最小接口分组

- **WHEN** 后续 `/spec` 或实现讨论当前阶段读接口
- **THEN** 最小读动作接口分组至少应包含 `ModuleListRead` 与 `ModuleDetailRead`
- **AND** `ModuleListRead` 只服务列表展示、筛选入口与进入详情
- **AND** `ModuleDetailRead` 只服务详情展示、版本主线、绑定关系与 `Decision` 入口

### Requirement: 绑定动作候选目标读取边界冻结

系统 SHALL 为 `BindModuleToProduct` 与 `MapModuleToRepository` 提供最小候选目标读取前提，使绑定动作在 `Module Detail` 中可真正闭环，而不扩写为 `Product Registry` / `Repository Binding` 的独立主线。

#### Scenario: 绑定候选目标最小读取

- **WHEN** 后续 `/spec` 或实现讨论绑定动作的读前提
- **THEN** 当前阶段允许 `Module Detail` 读取最小候选目标集合，接口分组至少应包含 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead`
- **AND** `ProductBindingCandidateRead` 只服务 `BindModuleToProduct` 的候选 Product 选择
- **AND** `RepositoryBindingCandidateRead` 只服务 `MapModuleToRepository` 的候选 Repository 选择
- **AND** 不得把这两类候选读取扩写为 `Product Registry` / `Repository Binding` 的独立主线读取接口

### Requirement: 写动作接口分组冻结

系统 SHALL 将当前阶段写动作接口分为创建、版本与关联三组最小写入接口，而不是提前冻结更复杂的写模型体系。

#### Scenario: 写动作最小接口分组

- **WHEN** 后续 `/spec` 或实现讨论当前阶段写接口
- **THEN** 最小写动作接口分组至少应包含 `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite`
- **AND** `ModuleCreateWrite` 只承接 `CreateModule`
- **AND** `ModuleReleaseWrite` 只承接 `CreateRelease`
- **AND** `ModuleBindingWrite` 只承接 `BindModuleToProduct` 与 `MapModuleToRepository`

### Requirement: Decision 当前阶段接口边界冻结

系统 SHALL 将 `Decision` 在 `phase02` 中的接口边界冻结为只读展示或跳转前提，不扩写为独立写入主线。

#### Scenario: Decision 接口边界

- **WHEN** 后续 `/spec` 或实现讨论 `Decision` 在 `Module Registry` 中的接口
- **THEN** 当前阶段只允许承接相关 `Decision` 的只读读取或跳转前提
- **AND** `Decision` 入口作为 `ModuleDetailRead` 的附属读取承接，不设独立读接口组
- **AND** 不得在 `phase02` 中新增独立 `RecordDecision` 写接口主线
- **AND** 不得把 `Decision Center` 的全量接口提前并入当前阶段

### Requirement: 非目标接口边界冻结

系统 SHALL 明确当前阶段不提前冻结完整查询矩阵或 `Dashboard` 聚合接口。

#### Scenario: 非目标接口判定

- **WHEN** 后续 `/spec` 或实现讨论跨页面聚合查询、复杂分析查询或 `Dashboard` 反馈接口
- **THEN** 必须判定为当前阶段非目标
- **AND** 不得把它们写成 `phase02-05` 的必需接口前提

## MODIFIED Requirements

### Requirement: Contract First 当前阶段解释

`Contract First` 在 `phase02` 当前阶段 SHALL 被解释为“先冻结最小接口边界与接口分组，再进入后续详细合同设计”，而不是一步到位冻结完整协议矩阵。

#### Scenario: Contract First 承接方式

- **WHEN** 后续 `/spec` 或实现讨论 `Contract First`
- **THEN** 必须优先沿用当前已冻结的最小读写接口分组
- **AND** 不得在当前阶段额外引入与 `Protocol Buffers` 长期方向冲突的第二套跨语言路线

## REMOVED Requirements

### Requirement: 当前阶段完整查询矩阵冻结
**Reason**: `phase02` 当前目标是建立 `Module Registry` 最小可执行主线，而不是提前完成所有查询与聚合接口规划。
**Migration**: 若后续需要完整查询矩阵或 `Dashboard` 聚合接口，应在 `phase03+` 或对应新阶段中单独冻结，不在 `phase02-05` 当前规格中处理。
