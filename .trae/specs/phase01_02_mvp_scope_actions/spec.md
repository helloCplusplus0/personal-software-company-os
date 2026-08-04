# Phase01-02 冻结 MVP 对象与动作范围 Spec

## Why

在 `phase01-01` 已经冻结技术路线与系统边界后，下一步必须把 `v0.1` 的对象范围和动作范围冻结成单值结论。只有先明确哪些对象进入主执行范围、哪些对象只是派生层或后移层，以及 MVP 至少承接哪些动作，后续页面、数据、API 与实现才不会扩范围或回到泛化叙述。

## What Changes

- 冻结 `v0.1` 主执行范围内的核心实体
- 冻结 `Capability` 为派生层，不作为重实体
- 冻结 `Feature / Opportunity / Experiment` 为后移对象
- 冻结 MVP 核心动作清单
- 冻结 `Product / Module / Repository` 之间的最小绑定动作
- 冻结 `Decision` 必须进入 MVP，并保留与目标对象的关联动作

## Impact

- Affected specs: `phase01_mvp_spec_convergence`
- Affected code: 当前无代码改动，影响后续对象模型、动作合同、页面职责、数据结构与 API 设计

## ADDED Requirements

### Requirement: MVP 核心实体范围冻结
系统 SHALL 将以下对象冻结为 `v0.1` 主执行范围内的核心实体：`Product`、`Module`、`Release`、`Decision`、`Repository`、`Venture（可选）`。

#### Scenario: 判断对象是否进入 MVP
- **WHEN** 接手者查询某对象是否属于 `v0.1` 主执行范围
- **THEN** 只有 `Product`、`Module`、`Release`、`Decision`、`Repository`、`Venture（可选）` 被视为核心实体
- **AND** 不得再把其他长期理论对象写成 `v0.1` 当前核心实体

### Requirement: Capability 派生层冻结
系统 SHALL 将 `Capability` 冻结为派生层，而不是 `v0.1` 中需要重维护的核心实体。

#### Scenario: 判断 Capability 定位
- **WHEN** 后续规格涉及 `Capability`
- **THEN** 必须将其解释为来自模块、版本、绑定关系与决策沉淀的派生结果
- **AND** 不得要求用户在 `v0.1` 中重录入或重维护 `Capability`

### Requirement: 后移对象范围冻结
系统 SHALL 将 `Feature`、`Opportunity`、`Experiment` 冻结为长期理论模型中的保留对象，但不进入 `v0.1` 主执行范围。

#### Scenario: 判断后移对象边界
- **WHEN** 后续规格涉及 `Feature`、`Opportunity` 或 `Experiment`
- **THEN** 必须明确其属于后移对象
- **AND** 不得把这些对象写成 `v0.1` 首轮必须承接的主范围

### Requirement: MVP 核心动作清单冻结
系统 SHALL 至少冻结以下 MVP 核心动作：`CreateProduct`、`CreateModule`、`CreateRelease`、`CreateRepository`、`RecordDecision`、`LinkDecisionToTarget`。

#### Scenario: 判断动作是否进入 MVP
- **WHEN** 接手者查询 `v0.1` 的核心动作
- **THEN** 必须至少包含 `CreateProduct`、`CreateModule`、`CreateRelease`、`CreateRepository`、`RecordDecision`、`LinkDecisionToTarget`
- **AND** 不得将 MVP 动作重新退化为泛化的“信息登记”表述

### Requirement: Repository 进入系统入口动作冻结
系统 SHALL 明确 `Repository` 作为核心实体，必须提供显式进入动作 `CreateRepository`，使 `Repository` 基础资产可以建立。

#### Scenario: Repository 建立入口
- **WHEN** 接手者查询 `Repository` 如何进入 `v0.1` 系统
- **THEN** 必须通过 `CreateRepository` 动作建立 `Repository` 基础资产
- **AND** 不得把 `Repository` 的建立退化为仅依赖绑定动作间接隐式产生

### Requirement: Repository Binding 最小动作冻结
系统 SHALL 将 `Product / Module / Repository` 之间的最小绑定动作冻结为 `BindRepositoryToProduct`、`BindModuleToProduct`、`MapModuleToRepository`。

#### Scenario: 判断绑定动作边界
- **WHEN** 后续规格涉及 `Repository Binding`
- **THEN** 必须至少细化到 `BindRepositoryToProduct`、`BindModuleToProduct`、`MapModuleToRepository`
- **AND** 不得继续使用无法落地的泛化动作命名替代上述绑定动作

### Requirement: Decision 必须保留在 MVP
系统 SHALL 明确 `Decision` 必须进入 `v0.1`，并允许 `Decision` 与目标对象建立结构化关联。

#### Scenario: 判断 Decision 定位
- **WHEN** 后续规格涉及 `Decision`
- **THEN** 必须将其视为 `v0.1` 的核心实体与核心动作承载点
- **AND** 不得将其从 MVP 中移除或降级为可有可无的辅助记录

## MODIFIED Requirements

### Requirement: MVP 页面与对象动作的一致性前提
后续页面、数据、API 规格 SHALL 以本次冻结的核心实体、派生层、后移层和动作矩阵为唯一上游，不得再独立扩展第二套对象范围或第二套动作范围。

#### Scenario: 后续规格引用边界
- **WHEN** 后续 `/spec` 编写页面、数据、API 或实现前提
- **THEN** 必须引用本次冻结的对象与动作范围
- **AND** 不得绕开本规格重新定义 `v0.1` 的对象与动作边界
