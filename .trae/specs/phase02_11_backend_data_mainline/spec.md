# phase02-11 后端与数据主线实现 Spec

## Why

`phase02-09` 已经把 `Module Registry` 的后端边界、数据范围与接口分组冻结为单一正式规格入口，`phase02-10` 也已经交付前端可运行主线。当前仍缺少一条可运行的后端与数据库主线，来替换前端临时适配层，并把 `Module Registry` 的最小读写闭环真正落到 Go 后端与 PostgreSQL 数据层。

## What Changes

- 实现 `Module Registry` 所需的最小后端读写接口
- 实现 `modules / module_releases / product_modules / module_repositories` 的数据主线
- 实现 `decisions / decision_links` 的当前阶段只读承接边界
- 实现 `products / repositories` 的最小只读候选读取前提，但不扩写为当前阶段写入主线
- 冻结本地开发数据库策略：复用现有 Podman PostgreSQL 容器，为 PSCO 创建独立数据库，不新开第二个数据库容器
- 明确迁移、最小种子数据或测试 fixture、后端运行入口与验收约束

## Impact

- Affected specs:
  - `phase02_09_module_registry_formal_spec`
  - `phase02_08_backend_module_boundary_interface_grouping`
  - `phase02_05_data_api_boundary`
  - `phase02_10_frontend_module_registry_mainline`
- Affected code:
  - `backend/`
  - `database/`
  - `scripts/`
  - 本地共享 PostgreSQL 开发实例

## ADDED Requirements

### Requirement: 实现 Module Registry 后端最小读写主线

系统 SHALL 在 `backend/internal/moduleregistry/` 下实现 `Module Registry` 的最小可运行后端主线，并严格遵守 `phase02-08 / 09` 已冻结的模块边界、接口分组、文件落点与分层语义。

#### Scenario: 后端读组与写组可运行

- **WHEN** `phase02-11` 开始实现 `Module Registry` 后端主线
- **THEN** 读组必须承接 `ModuleListRead` 与 `ModuleDetailRead`
- **AND** 写组必须承接 `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite`
- **AND** 实现文件必须落在 `backend/internal/moduleregistry/handler/`、`service/`、`repository/`、`candidate/`
- **AND** 不得把 `Module Registry` 能力拆散到 `Product`、`Repository` 或 `Decision` 的后端模块中

### Requirement: 落实 PostgreSQL 数据主线

系统 SHALL 在 PostgreSQL 中实现 `Module Registry` 当前阶段的数据主线，至少覆盖 `modules`、`module_releases`、`product_modules`、`module_repositories` 四张直接承接表，并保证字段语义、约束与 `phase02-09` 的最小读写模型一致。

#### Scenario: 核心表落地

- **WHEN** `phase02-11` 实现数据库迁移
- **THEN** 必须创建 `modules`、`module_releases`、`product_modules`、`module_repositories`
- **AND** `modules.name` 必须具备唯一性约束
- **AND** `module_releases.module_id`、`product_modules.module_id`、`module_repositories.module_id` 必须与 `modules.id` 建立正确的关系约束
- **AND** 不得引入超出当前阶段的新主线实体表来承接 `Feature / Opportunity / Experiment`

### Requirement: 实现当前阶段只读前提表

系统 SHALL 为 `Decision` 附属读取与 `Product / Repository` 候选读取实现最小只读前提表或其等价结构前提，但这些表在 `phase02` 中只服务读取，不扩写为新的写入主线。

#### Scenario: 只读前提表存在

- **WHEN** `ModuleDetailRead` 需要展示 `Decision` 入口，或绑定动作需要读取 `Product / Repository` 候选
- **THEN** 数据库中必须存在支撑这些读取的最小前提结构
- **AND** 至少需要覆盖 `decisions`、`decision_links`、`products`、`repositories`
- **AND** `decisions / decision_links` 只服务详情页附属读取或跳转前提
- **AND** `products / repositories` 只服务候选读取，不要求在当前阶段实现 `CreateProduct` 或 `CreateRepository`

### Requirement: Decision 边界必须保持为内嵌附属读取

系统 SHALL 将 `Decision` 当前阶段实现保持为 `ModuleDetailRead` 的内嵌附属读取，不设独立读接口组，不在 `Module Registry` 中新增独立决策写入主线。

#### Scenario: 详情读取承接 Decision 入口

- **WHEN** 后端实现 `ModuleDetailRead`
- **THEN** `Decision` 关联信息必须在详情读取编排中一并组装
- **AND** 不得单独新增 `Decision` 的 handler、service 或独立读接口
- **AND** 不得新增 `RecordDecision` 或 `LinkDecisionToTarget` 的写入接口到 `Module Registry`

### Requirement: 本地开发环境必须复用共享 PostgreSQL 容器

系统 SHALL 在本地开发环境下复用现有 Podman PostgreSQL 容器，而不是为 PSCO 再启动第二个数据库容器；但 PSCO 必须在该共享实例中拥有独立数据库，避免与其他项目混用同一数据库。

#### Scenario: 本地共享实例中的 PSCO 独立数据库

- **WHEN** `phase02-11` 在本地开发环境中初始化数据库
- **THEN** 必须复用当前已运行的 PostgreSQL 容器实例（宿主机端口 `55432`）
- **AND** 必须为 PSCO 创建独立数据库，推荐命名为 `psco_development`
- **AND** 不得直接复用 `rento_production` 作为 PSCO 的工作数据库
- **AND** 不得因为当前项目进入实现而新开第二个本地 PostgreSQL 容器

#### Scenario: DATABASE_URL 显式且可直接使用

- **WHEN** 本地环境需要配置 PSCO 的数据库连接
- **THEN** 必须提供显式最终值的 `DATABASE_URL`
- **AND** 若密码含有 `/`、`=`、`@` 等特殊字符，URL 中必须完成编码
- **AND** 不得依赖嵌套环境变量拼接出最终 `DATABASE_URL`
- **AND** 凭据应通过本地环境文件或本地运行环境提供，不应把敏感值硬编码进仓库源码

### Requirement: 迁移与最小种子数据必须可支撑 phase02-12 验收

系统 SHALL 提供可重复执行的数据库迁移主线，并提供支撑 `phase02-12` 联调验收的最小候选数据或测试 fixture。

#### Scenario: 数据迁移与验收前提可复现

- **WHEN** 开发者在本地准备 `phase02-11` 或 `phase02-12` 环境
- **THEN** 必须能够通过统一的迁移入口创建所需表结构
- **AND** 必须能够初始化最小的 `products / repositories` 候选记录与必要的 `decisions / decision_links` 示例数据
- **AND** 这些初始化数据只服务当前阶段联调与验收，不扩写为第二套业务主线

### Requirement: 前端临时适配层必须能切换到真实后端

系统 SHALL 让 `phase02-10` 前端主线能够在不改变页面职责与交互语义的前提下，从 mock 适配层切换到真实后端读写接口。

#### Scenario: phase02-10 接入真实后端

- **WHEN** `phase02-11` 后端与数据主线完成
- **THEN** 前端必须能够把列表、详情、创建、版本登记、绑定动作与候选读取切换到真实后端
- **AND** 不得借机引入第二套对象字段语义、第二套返回路径或第二套数据主线

## MODIFIED Requirements

### Requirement: phase02-09 正式规格正文进入后端实现态

`phase02-09` 已冻结的后端实现设计与数据边界，在本阶段必须从“可直接实现的设计结果”推进为“可运行的后端与数据库主线”，并成为 `phase02-12` 联调与验收的直接前置条件。

#### Scenario: 进入 phase02-12 的后端前置满足

- **WHEN** `phase02-11` 完成
- **THEN** 后端读写接口必须已经可以运行
- **AND** 数据主线必须与已冻结边界一致
- **AND** 不得引入超出当前阶段的新对象解释或第二套数据主线
- **AND** 前端临时数据适配层必须具备切换到真实后端的明确落点

## REMOVED Requirements
