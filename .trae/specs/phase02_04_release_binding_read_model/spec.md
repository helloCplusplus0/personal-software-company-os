# Phase02-04 版本登记、关联动作归属与最小读模型 Spec

## Why

`phase02-01 ~ 03` 已经冻结了页面边界、实体与创建闭环，但 `Module Registry` 要真正进入可实现状态，还需要继续把版本登记路径、`BindModuleToProduct / MapModuleToRepository` 的动作归属、`Decision` 的最小关联入口，以及列表页/详情页各自到底读取什么内容写成单值结论。否则后续 `/spec` 与实现仍会在“谁负责写、谁负责读、哪些是轻量入口”上继续漂移。

## What Changes

- 冻结 `CreateRelease` 的最小交互路径
- 冻结 `BindModuleToProduct` 与 `MapModuleToRepository` 在页面中的动作拥有者
- 冻结到 `Decision` 的最小关联入口
- 冻结 `Module Registry / List` 与 `Module Detail` 的最小读模型
- 明确当前阶段不把 `phase03+` 的独立主线提前并入

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `frontend/` 中的版本登记入口、模块详情关联面板、列表页读取模型、详情页读取模型与相关读写接口承接

## ADDED Requirements

### Requirement: CreateRelease 最小交互路径冻结

系统 SHALL 将 `CreateRelease` 冻结为从 `Module Detail` 进入、完成后仍留在 `Module Registry` 主线中的最小交互。

#### Scenario: 版本登记入口

- **WHEN** 用户要为某个模块登记新版本
- **THEN** 必须从该模块的 `Module Detail` 进入 `Release Create`
- **AND** 不得要求用户跳转到独立的版本管理主线

#### Scenario: 版本登记完成

- **WHEN** 用户完成 `CreateRelease`
- **THEN** 结果必须回流到当前模块的详情上下文
- **AND** 用户必须能够在 `Module Detail` 中看到更新后的版本主线

### Requirement: 绑定动作归属冻结

系统 SHALL 将 `BindModuleToProduct` 与 `MapModuleToRepository` 冻结为由 `Module Detail` 直接承接的最小写入动作。

#### Scenario: BindModuleToProduct 动作拥有者

- **WHEN** 用户要将某个模块绑定到 `Product`
- **THEN** 当前阶段必须由该模块的 `Module Detail` 承接写入动作
- **AND** 不得仅提供跳转而不提供最小写入触点

#### Scenario: MapModuleToRepository 动作拥有者

- **WHEN** 用户要为某个模块映射 `Repository`
- **THEN** 当前阶段必须由该模块的 `Module Detail` 承接写入动作
- **AND** 不得将该动作提前并入 `Repository Binding` 独立主线

### Requirement: Decision 最小关联入口冻结

系统 SHALL 将 `Decision` 在 `phase02` 中的作用冻结为与当前模块相关的只读展示或跳转入口，而不是独立写入主线。

#### Scenario: Decision 入口范围

- **WHEN** 用户查看某个模块详情
- **THEN** 页面可以展示相关 `Decision` 入口或跳转
- **AND** 当前阶段不得在 `Module Registry` 中直接扩写为完整 `RecordDecision` 主线
- **AND** 不得把 `Decision Center` 的独立实现需求提前并入 `phase02`

### Requirement: 列表页最小读模型冻结

系统 SHALL 将 `Module Registry / List` 的最小读模型冻结为服务模块列表阅读与进入详情的最小数据集合。

#### Scenario: 列表读取模型

- **WHEN** 用户进入 `Module Registry / List`
- **THEN** 页面最小读模型至少应包含 `name / description / status / latest_release / product_bind_count / repository_bind_count`
- **AND** 当前阶段的列表读取只服务于列表展示、筛选入口与进入详情
- **AND** 不得在当前阶段扩写为复杂搜索分析工作台

### Requirement: 详情页最小读模型冻结

系统 SHALL 将 `Module Detail` 的最小读模型冻结为支持详情展示、版本主线、绑定关系与 `Decision` 入口的一组最小读取结构。

#### Scenario: 详情读取模型

- **WHEN** 用户进入 `Module Detail`
- **THEN** 页面最小读模型至少应包含核心对象字段、版本列表、产品绑定、仓库映射与相关 `Decision` 入口
- **AND** 当前阶段必须让用户在同一详情上下文中理解该模块的当前状态
- **AND** 不得要求用户跳到 `phase03+` 的独立主线才能完成当前阶段核心理解

### Requirement: phase03+ 非目标边界冻结

系统 SHALL 明确 `phase02-04` 只冻结当前主线所需的版本、绑定与读取模型，不提前并入后续独立主线。

#### Scenario: 非目标判定

- **WHEN** 后续 `/spec` 或实现讨论 `phase02-04`
- **THEN** 不得把 `Decision Center` 全量主线提前并入
- **AND** 不得把 `Product Registry` 全量主线提前并入
- **AND** 不得把 `Repository Binding` 全量主线提前并入

## MODIFIED Requirements

### Requirement: Module Detail 页面职责

`Module Detail` 在当前阶段 SHALL 被解释为版本登记、绑定动作与相关只读入口的综合承接页，而不是只读展示页。

#### Scenario: 详情页职责解释

- **WHEN** 后续 `/spec` 或实现讨论 `Module Detail`
- **THEN** 必须将其理解为当前阶段版本与绑定动作的直接拥有者
- **AND** 不得退回为仅展示详情而把关键动作散落到其他主线

## REMOVED Requirements

### Requirement: 版本与绑定动作跳转式承接
**Reason**: 当前阶段已经冻结 `Module Detail` 直接承接版本登记与绑定动作，继续保留纯跳转式承接会导致页面归属继续漂移。
**Migration**: 若后续确需将相关动作升级为独立主线，必须在 `phase03+` 或新的 `audit` 流程中重新冻结，不在 `phase02-04` 当前规格中处理。
