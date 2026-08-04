# Phase02-01 Module Registry 页面边界与信息结构 Spec

## Why

`phase02` 要让 `Module Registry` 成为 `v0.1` 第一条可直接进入实现的主线，第一步必须先把页面边界、页面之间的跳转关系，以及 `PC / 移动浏览器` 下的信息结构冻结成单值结论。只有先收住页面职责，后续的对象、状态、API 与源码设计层规格才不会继续漂移。

## What Changes

- 冻结 `Module Registry / List`、`Module Create`、`Module Detail`、`Release Create` 的页面边界
- 冻结这四类页面之间的最小跳转关系
- 冻结 `Decision Center / Product Registry / Repository Binding` 在 `phase02` 中的入口方式
- 冻结 `PC` 与移动浏览器下的基础信息密度策略
- 明确 `phase02` 不引入第二套移动端 UI 方案，不引入独立 `React Native` 客户端，不把完整 `PWA` 写成前置范围

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `frontend/` 中的模块列表页、模块创建页、模块详情页、版本登记入口与相关路由结构

## ADDED Requirements

### Requirement: Module Registry 页面边界冻结

系统 SHALL 将 `phase02` 的页面主线冻结为 `Module Registry / List`、`Module Create`、`Module Detail`、`Release Create` 四类页面或页面级功能块。

#### Scenario: 页面范围判定

- **WHEN** 接手者阅读 `phase02` 页面规格
- **THEN** 必须得到上述四类页面的单值结论
- **AND** 不得在本阶段额外引入 `Feature / Opportunity / Experiment` 页面
- **AND** 不得把 `Decision Center`、`Product Registry`、`Repository Binding` 扩写为 `phase02` 的独立实现主线
- **AND** 不得把独立 `AI Assistant` 一级导航纳入 `phase02` 页面主线

### Requirement: Module Registry 列表页职责冻结

系统 SHALL 将 `Module Registry / List` 冻结为当前阶段的默认进入页，承接模块读取、筛选入口、创建入口与进入详情入口。

#### Scenario: 列表页职责判定

- **WHEN** 用户进入 `Module Registry`
- **THEN** 页面必须承接模块列表读取
- **AND** 必须提供进入 `Module Create` 的明确入口
- **AND** 必须提供进入 `Module Detail` 的明确入口
- **AND** 当前阶段的筛选只作为列表页入口能力存在，不扩写为复杂检索工作台

### Requirement: Module Create 页面职责冻结

系统 SHALL 将 `Module Create` 冻结为 `CreateModule` 的唯一页面级承接入口。

#### Scenario: 创建页职责判定

- **WHEN** 用户执行 `CreateModule`
- **THEN** 必须通过 `Module Create` 页面或等价页面级表单完成
- **AND** 当前阶段不得把模块创建分散到列表页内联复杂编辑流或其他独立页面主线

### Requirement: Module Detail 页面职责冻结

系统 SHALL 将 `Module Detail` 冻结为详情读取、`CreateRelease`、`BindModuleToProduct` 与 `MapModuleToRepository` 的最小承接页。

#### Scenario: 详情页职责判定

- **WHEN** 用户进入某个模块详情
- **THEN** 页面必须展示该模块的核心信息
- **AND** 必须承接版本列表或版本入口
- **AND** 必须承接产品绑定与仓库映射的最小写入触点
- **AND** 与 `Decision` 的关系在当前阶段只能以只读展示或跳转入口承接

### Requirement: Release Create 页面职责冻结

系统 SHALL 将 `Release Create` 冻结为版本登记的最小页面级入口，并与 `Module Detail` 保持直接上下文关联。

#### Scenario: 版本登记入口判定

- **WHEN** 用户要为某个 `Module` 创建 `Release`
- **THEN** 必须从该模块详情上下文进入 `Release Create`
- **AND** 不得要求用户跳出 `Module Registry` 主线去其他独立主线完成版本登记

### Requirement: 页面跳转关系冻结

系统 SHALL 冻结 `Module Registry` 主线中的最小跳转关系，避免页面职责与交互流在后续规格中继续漂移。

#### Scenario: 最小跳转关系判定

- **WHEN** 用户从列表进入模块主线
- **THEN** 必须支持 `Module Registry / List -> Module Create`
- **AND** 必须支持 `Module Registry / List -> Module Detail`
- **AND** 必须支持 `Module Detail -> Release Create`
- **AND** 必须支持 `Module Detail -> Product Registry / Repository Binding / Decision Center` 的轻量跳转或关联入口
- **AND** 不得把这些轻量入口扩写为当前阶段的第二条页面主线

### Requirement: PC 与移动浏览器信息密度策略冻结

系统 SHALL 在单一 `React Web` 前端交付策略下，同时定义 `PC` 与移动浏览器的基础信息密度规则。

#### Scenario: 桌面端信息密度

- **WHEN** 页面在 `PC` 桌面环境展示
- **THEN** 列表页应优先承接更高信息密度
- **AND** 详情页可在同屏承接核心信息、版本区块与关联入口

#### Scenario: 移动浏览器信息密度

- **WHEN** 页面在移动浏览器窄屏环境展示
- **THEN** 必须采用同一套页面语义与动作体系
- **AND** 必须通过信息裁剪、区块折叠、垂直重排或分层展示降低拥挤度
- **AND** 不得引入第二套独立移动端 UI 方案
- **AND** 不得要求独立 `React Native` 客户端或完整 `PWA` 作为当前阶段前提

## MODIFIED Requirements

### Requirement: Module Registry 页面承接方式

`Module Registry` 在 `phase02` 中 SHALL 不只被解释为一个“模块列表概念”，而是被解释为包含列表、创建、详情、版本登记入口在内的一条完整页面主线。

#### Scenario: 页面主线解释

- **WHEN** 后续 `/spec` 或实现讨论 `Module Registry`
- **THEN** 必须将其理解为一条完整页面主线
- **AND** 不得只剩单独列表页而把创建、详情、版本登记散落为无归属入口

## REMOVED Requirements

### Requirement: 第二套移动端 UI 方案
**Reason**: 当前阶段已冻结为单一 `React Web` 同时覆盖 `PC` 与移动浏览器，不需要并行维护第二套移动端页面方案。
**Migration**: 若后续确有独立客户端需要，必须进入新的 `phase / audit` 流程，不在 `phase02-01` 当前规格中处理。
