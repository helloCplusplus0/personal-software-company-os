# Phase01-04 冻结数据与合同基线 Spec

## Why

在 `phase01-01` 到 `phase01-03` 已经完成技术路线、对象动作范围和页面输入路径冻结之后，下一步必须把 `v0.1` 的数据主线、最小数据模型方向和 contract-first 边界冻结成单值结论。只有先明确当前数据库主线、核心表关系、`Repository Binding` 与 `Decision Record` 的最小结构要求，后续 `/spec` 才不会重新引入第二套数据库解释、第二套合同路线或过早落入实现细节分叉。

## What Changes

- 冻结 `v0.1` 的数据库主线为 `PostgreSQL`
- 冻结 `v0.1` 最小数据模型方向，包括核心表、关系表和派生视图
- 冻结 `Repository Binding` 的最小结构要求，包含 `MapModuleToRepository` 的显式关系载体（`module_repositories`）
- 冻结 `Decision Record` 的最小结构要求
- 冻结 `Contract First` 在当前项目中的最小落地边界
- 明确当前阶段不要求完整 `proto` 工具链和具体 Go 数据访问实现选型

## Impact

- Affected specs: `phase01_mvp_spec_convergence`
- Affected code: 当前无代码改动，影响后续数据库 schema 设计、关系模型、API 合同、Go 后端接口边界与 `/spec` 正文编写

## ADDED Requirements

### Requirement: 数据库主线冻结
系统 SHALL 将 `PostgreSQL` 冻结为 `v0.1` 的唯一数据库主线，并与 `Durable System Track` 保持一致。

#### Scenario: 判断当前数据库主线
- **WHEN** 接手者查询 `v0.1` 当前数据库主线
- **THEN** 必须得到 `PostgreSQL` 的单值结论
- **AND** 不得把 `SQLite`、文件数据库或其他数据库重新写成当前默认主线

### Requirement: 最小数据模型方向冻结
系统 SHALL 将 `v0.1` 的最小数据模型方向冻结为核心表、关系表与派生视图三层结构。

#### Scenario: 判断最小数据模型方向
- **WHEN** 后续规格定义 `v0.1` 的数据结构
- **THEN** 核心表至少必须围绕 `ventures（可选）`、`products`、`repositories`、`modules`、`module_releases`、`decisions`
- **AND** 关系表至少必须围绕 `product_modules`、`product_repositories`、`module_repositories`、`decision_links`
- **AND** 派生视图只能承接 `Capability` 与资产反馈类聚合结果，而不得回退为新的重实体主线

### Requirement: Capability 只作为派生结果层
系统 SHALL 明确 `Capability` 在当前数据基线中只作为派生结果层，不进入 `v0.1` 核心表主线。

#### Scenario: 判断 Capability 数据定位
- **WHEN** 后续规格涉及 `Capability` 的数据表示
- **THEN** 必须将其解释为基于模块、版本、绑定关系和决策沉淀生成的派生结果
- **AND** 不得为 `Capability` 额外建立与当前 MVP 主线并列的核心 CRUD 实体要求

### Requirement: Repository Binding 结构要求冻结
系统 SHALL 将 `Repository Binding` 的结构要求冻结为 `Product` 绑定、`Module` 绑定与实现映射三类关系，并采用显式关系承载。

#### Scenario: 判断 Repository Binding 结构
- **WHEN** 后续规格定义 `Repository Binding`
- **THEN** 必须至少承接 `product_repositories` 这类产品与仓库绑定关系
- **AND** 必须承接模块与产品的显式绑定关系
- **AND** 必须承接模块与仓库之间的实现映射关系，通过 `module_repositories` 显式承载，而不得只保留页面层泛化说明

### Requirement: Decision Record 结构要求冻结
系统 SHALL 将 `Decision Record` 冻结为结构化记录，而不是自由文本备注。

#### Scenario: 判断 Decision Record 最小结构
- **WHEN** 后续规格定义 `Decision` 数据结构
- **THEN** 必须至少承接 `title`、`context`、`problem`、`alternatives`、`choice`、`reason`、`impact`、`status`
- **AND** 必须支持通过 `decision_links` 将 `Decision` 与目标对象建立结构化关联
- **AND** 不得把 `Decision` 退化为无法检索、无法关联的散装备注

### Requirement: Contract First 最小边界冻结
系统 SHALL 在当前项目中冻结 `Contract First` 的最小边界，使后续 API 与跨语言边界可以与长期 `Protocol Buffers` 方向保持一致。

#### Scenario: 判断当前合同边界
- **WHEN** 后续规格涉及 API、前后端边界或 TS / Go / Rust 的跨语言边界
- **THEN** 必须以结构化合同优先，而不是由前端猜字段或由实现倒推合同
- **AND** 合同方向必须与 `Protocol Buffers` 长期标准兼容
- **AND** 当前阶段可以暂不落地完整 `proto` 生成链

### Requirement: 当前阶段不冻结第二套合同路线
系统 SHALL 明确当前阶段不得再引入与 `Protocol Buffers` 长期方向冲突的第二套跨语言合同主线。

#### Scenario: 判断第二套合同路线是否允许
- **WHEN** 后续规格讨论跨语言合同
- **THEN** 不得把 `OpenAPI`、`GraphQL`、自定义 JSON 协议或其他方案写成当前项目的并列长期主线
- **AND** 若某些页面或接口需要临时说明格式，也必须从属于 `Contract First` 总方向，而不是形成第二套正式标准

## MODIFIED Requirements

### Requirement: 后续数据与 API 规格引用前提
后续 `/spec` 编写数据模型、关系结构、API 合同与后端接口边界时 SHALL 以本次冻结的数据主线、最小数据模型方向与合同边界为唯一上游，不得重新解释当前数据库主线或重建第二套合同体系。

#### Scenario: 后续数据与 API 规格编写
- **WHEN** 后续 `/spec` 定义表结构、关系结构、接口输入输出或跨语言边界
- **THEN** 必须引用本次冻结的数据与合同基线
- **AND** 不得绕开本规格重新定义 `v0.1` 的数据库和合同方向
