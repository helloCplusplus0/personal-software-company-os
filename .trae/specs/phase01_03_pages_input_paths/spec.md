# Phase01-03 冻结页面与输入路径 Spec

## Why

在技术路线、对象范围和动作范围已经冻结后，下一步必须把 `v0.1` 的页面范围、每个页面的最小职责，以及首轮录入路径冻结成单值结论。只有页面和输入路径先收口，后续数据、API 和实现才不会重新引入未来页面、独立 AI 工作台或高摩擦录入流程。

## What Changes

- 冻结 `v0.1` 的页面级主范围
- 冻结每个页面的最小职责
- 冻结页面与动作范围的一致性约束
- 冻结冷启动场景下的首轮录入路径
- 冻结低摩擦录入原则与空状态前提
- 明确独立 `AI Assistant` 工作台和未来页面不进入 `v0.1`

## Impact

- Affected specs: `phase01_mvp_spec_convergence`
- Affected code: 当前无代码改动，影响后续页面信息架构、导航、空状态设计、表单流和录入流程

## ADDED Requirements

### Requirement: MVP 页面范围冻结
系统 SHALL 将 `v0.1` 页面级主范围冻结为 `Dashboard`、`Module Registry`、`Product Registry`、`Decision Center`、`Repository Binding`。

#### Scenario: 判断页面是否进入 MVP
- **WHEN** 接手者查询某页面是否属于 `v0.1` 主范围
- **THEN** 只有 `Dashboard`、`Module Registry`、`Product Registry`、`Decision Center`、`Repository Binding` 被视为 MVP 页面
- **AND** 不得把其他未来页面提前写入 `v0.1`

### Requirement: Dashboard 最小职责冻结
系统 SHALL 将 `Dashboard` 冻结为最小聚合反馈页面，而不是复杂驾驶舱。

#### Scenario: Dashboard 职责判定
- **WHEN** 后续规格涉及 `Dashboard`
- **THEN** 必须将其职责限定为 MVP 资产反馈、关键状态聚合与最小概览
- **AND** 不得把它扩展为复杂分析驾驶舱或独立 AI 工作台入口

### Requirement: Module Registry 最小职责冻结
系统 SHALL 将 `Module Registry` 冻结为模块资产登记、查看、筛选与后续动作入口页面。

#### Scenario: Module Registry 职责判定
- **WHEN** 后续规格涉及 `Module Registry`
- **THEN** 必须承接 `CreateModule`、`CreateRelease`、模块查看与模块相关后续动作入口
- **AND** 不得让其承担不属于 `v0.1` 的未来对象流程

### Requirement: Product Registry 最小职责冻结
系统 SHALL 将 `Product Registry` 冻结为产品资产登记、查看、绑定入口与上下文聚合页面。

#### Scenario: Product Registry 职责判定
- **WHEN** 后续规格涉及 `Product Registry`
- **THEN** 必须承接 `CreateProduct` 以及产品与模块、仓库、决策关系的查看入口
- **AND** 不得让其承接不属于 `v0.1` 的未来业务对象流程

### Requirement: Decision Center 最小职责冻结
系统 SHALL 将 `Decision Center` 冻结为 `Decision` 记录、查看、筛选与关联入口页面。

#### Scenario: Decision Center 职责判定
- **WHEN** 后续规格涉及 `Decision Center`
- **THEN** 必须承接 `RecordDecision` 与 `LinkDecisionToTarget`
- **AND** 不得把 `Decision` 降级为散落在其他页面中的附属备注

### Requirement: Repository Binding 最小职责冻结
系统 SHALL 将 `Repository Binding` 冻结为 `Repository` 创建与 `Product / Module / Repository` 绑定关系管理页面。

#### Scenario: Repository Binding 职责判定
- **WHEN** 后续规格涉及 `Repository Binding`
- **THEN** 必须承接 `CreateRepository`、`BindRepositoryToProduct`、`BindModuleToProduct`、`MapModuleToRepository`
- **AND** 不得退回为只有泛化说明、没有显式动作入口的页面

### Requirement: 页面与动作范围一致性冻结
系统 SHALL 保证每个 MVP 页面至少对应一组已冻结的核心动作，不允许出现页面范围与动作范围断裂。

#### Scenario: 页面动作一致性校验
- **WHEN** 后续规格定义页面职责或导航
- **THEN** 必须能把页面职责映射到已冻结的 MVP 动作
- **AND** 不得新增没有动作承载的空页面

### Requirement: 冷启动录入路径冻结
系统 SHALL 冻结首轮冷启动的最小录入路径，使用户可以从零开始建立 `Product / Module / Repository / Decision` 基础资产并形成第一版资产反馈。

#### Scenario: 首轮冷启动路径
- **WHEN** 用户第一次进入系统且尚无资产数据
- **THEN** 系统必须允许用户按最小路径完成创建 `Product`、创建 `Repository`、登记 `Module`、记录 `Decision` 与基础绑定关系
- **AND** 后续可在 `Dashboard` 看到第一版资产反馈

### Requirement: 低摩擦录入原则冻结
系统 SHALL 将低摩擦录入作为 `v0.1` 的页面和表单设计前提。

#### Scenario: 录入流程设计
- **WHEN** 后续规格定义页面表单、空状态或入口动作
- **THEN** 必须优先采用低摩擦手动录入路径
- **AND** 不得把复杂导入、自动扫描或多层级配置作为首轮必经流程

### Requirement: 空状态与未来页面边界冻结
系统 SHALL 要求每个 MVP 页面后续都提供与冷启动一致的空状态入口，并明确独立 `AI Assistant` 工作台、`Feature` 页面、`Opportunity` 页面、`Experiment` 页面不进入 `v0.1`。

#### Scenario: 空状态与非目标页面判定
- **WHEN** 后续规格涉及页面空状态或新增页面
- **THEN** 空状态必须服务于首轮录入路径
- **AND** 不得把独立 `AI Assistant` 工作台、`Feature` 页面、`Opportunity` 页面或 `Experiment` 页面写入 `v0.1`

## MODIFIED Requirements

### Requirement: 后续页面规格引用前提
后续页面、表单、导航和 UX 规格 SHALL 以本次冻结的页面范围、最小职责和输入路径为唯一上游，不得重新发明第二套导航结构或第二套路径解释。

#### Scenario: 后续页面与导航规格编写
- **WHEN** 后续 `/spec` 编写页面、导航或 UX 细节
- **THEN** 必须引用本次冻结的页面范围和输入路径
- **AND** 不得绕开本规格重新定义 `v0.1` 的页面体系
