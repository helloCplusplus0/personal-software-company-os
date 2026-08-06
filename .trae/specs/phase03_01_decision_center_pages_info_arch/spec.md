# Phase03-01 Decision Center 页面边界与信息结构 Spec

## Why

`phase03` 要把 `Decision Center` 从 `Module Detail` 中的附属只读入口推进为第二条正式执行主线，第一步必须先把页面边界、页面之间的入口关系，以及 `PC / 移动浏览器` 下的信息结构冻结成单值结论。只有先收住页面职责，后续 `Decision` 模板、状态模型、接口分组与实现设计才不会继续漂移。

## What Changes

- 冻结 `Decision Center / List`、`Decision Create`、`Decision Detail` 的最小页面边界
- 冻结 `Decision Center` 主线中的最小跳转关系与默认进入路径
- 冻结 `Module Detail` 与 `Decision Center` 之间的最小入口关系
- 冻结 `Decision Detail` 中 `Decision -> Module` 候选读取与目标关联的页面级归属
- 冻结 `Decision Center / List`、`Decision Create`、`Decision Detail` 三个页面的最小页面级信息区块组成
- 冻结 `PC` 与移动浏览器下的基础信息密度策略
- 明确当前阶段不引入第二套移动端 UI 方案，不引入独立 `React Native` 客户端，不把完整 `PWA` 写成前置范围

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `frontend/` 中的 `Decision List`、`Decision Create`、`Decision Detail` 页面与路由结构，`Module Detail` 中的 `Decision` 上下文入口

## ADDED Requirements

### Requirement: Decision Center 页面边界冻结

系统 SHALL 将 `phase03` 的页面主线冻结为 `Decision Center / List`、`Decision Create`、`Decision Detail` 三类页面或页面级功能块。

#### Scenario: 页面范围判定

- **WHEN** 接手者阅读 `phase03-01` 页面规格
- **THEN** 必须得到上述三类页面的单值结论
- **AND** 不得在本阶段额外引入 `Feature / Opportunity / Experiment` 页面
- **AND** 不得把 `Product / Repository / Dashboard` 扩写为 `phase03-01` 的独立实现主线
- **AND** 不得把独立 `AI Assistant` 一级导航纳入 `phase03` 页面主线

### Requirement: Decision Center 列表页职责冻结

系统 SHALL 将 `Decision Center / List` 冻结为当前阶段的默认进入页，承接决策列表读取、筛选入口、创建入口与进入详情入口。

#### Scenario: 列表页职责判定

- **WHEN** 用户进入 `Decision Center`
- **THEN** 页面必须承接 `Decision` 列表读取
- **AND** 必须提供进入 `Decision Create` 的明确入口
- **AND** 必须提供进入 `Decision Detail` 的明确入口
- **AND** 当前阶段的筛选只作为列表页入口能力存在，不扩写为复杂检索工作台

### Requirement: Decision Create 页面职责冻结

系统 SHALL 将 `Decision Create` 冻结为 `RecordDecision` 的唯一页面级承接入口。

#### Scenario: 创建页职责判定

- **WHEN** 用户执行 `RecordDecision`
- **THEN** 必须通过 `Decision Create` 页面或等价页面级表单完成
- **AND** 当前阶段不得把决策创建分散到列表页内联复杂编辑流或其他独立页面主线

### Requirement: Decision Detail 页面职责冻结

系统 SHALL 将 `Decision Detail` 冻结为详情读取、已关联目标展示、`Decision -> Module` 候选读取与 `LinkDecisionToTarget` 的最小承接页。

#### Scenario: 详情页职责判定

- **WHEN** 用户进入某条 `Decision` 详情
- **THEN** 页面必须展示该 `Decision` 的核心字段与结构化模板字段
- **AND** 必须展示当前已建立的目标关联结果
- **AND** 必须承接面向 `Module` 的候选读取与目标选择
- **AND** 必须承接 `LinkDecisionToTarget` 的最小写入触点
- **AND** 不得把 `Decision Detail` 扩写为新的跨对象复合工作台

### Requirement: 页面跳转关系冻结

系统 SHALL 冻结 `Decision Center` 主线中的最小跳转关系，避免页面职责与交互流在后续规格中继续漂移。

#### Scenario: 最小跳转关系判定

- **WHEN** 用户从 `Decision Center` 进入决策主线
- **THEN** 必须支持 `Decision Center / List -> Decision Create`
- **AND** 必须支持 `Decision Center / List -> Decision Detail`
- **AND** 必须支持 `Decision Detail -> Decision Center / List`
- **AND** `Module Detail` 中的 `Decision` 入口必须冻结为两个单值触点，不得保留"或"式双路线
- **AND** 记录决策触点必须单值指向 `带上下文的 Decision Create`，用于从当前 `Module` 直接发起 `RecordDecision`
- **AND** 查看相关决策触点必须单值指向 `Decision Center / List`，用于查看与当前 `Module` 相关的已有决策
- **AND** 不得把这些入口扩写为当前阶段的第二条页面主线

### Requirement: Module Detail 入口归属冻结

系统 SHALL 将 `Module Detail` 中的 `Decision` 入口冻结为两个单值轻量触点，而不是 `phase03` 的第二个 `Decision` 工作台。

#### Scenario: Module Detail 记录决策触点判定

- **WHEN** 用户在 `Module Detail` 中触发"为当前 `Module` 记录决策"动作
- **THEN** 必须进入带上下文的 `Decision Create`
- **AND** 该上下文必须携带当前 `Module` 作为待关联目标
- **AND** 不得在 `Module Detail` 内承接完整的 `RecordDecision` 表单

#### Scenario: Module Detail 查看相关决策触点判定

- **WHEN** 用户在 `Module Detail` 中触发"查看当前 `Module` 相关决策"动作
- **THEN** 必须跳转到 `Decision Center / List`
- **AND** 可以携带当前 `Module` 上下文用于列表筛选
- **AND** 不得在 `Module Detail` 内承接完整的 `Decision List` 复合能力

#### Scenario: Module Detail 入口边界判定

- **WHEN** 评估 `Module Detail` 的 `Decision` 承接范围
- **THEN** 两个触点都必须作为轻量入口或上下文跳转存在
- **AND** 不得在 `Module Detail` 内承接 `Decision Detail` 或完整的 `Decision List / Create` 复合工作台

### Requirement: Decision Detail 目标候选读取归属冻结

系统 SHALL 将 `Decision -> Module` 的候选读取冻结为 `Decision Detail` 的附属读取能力，仅服务 `LinkDecisionToTarget`。

#### Scenario: 目标候选读取归属判定

- **WHEN** 用户在 `Decision Detail` 中准备执行目标关联
- **THEN** 页面必须在当前详情上下文中承接 `Module` 候选读取
- **AND** 候选读取的职责只能服务目标选择与关联写入
- **AND** 不得把候选读取扩写为独立 `Module` 浏览工作台

### Requirement: Decision Center / List 最小页面级信息区块冻结

系统 SHALL 将 `Decision Center / List` 的最小页面级信息区块冻结为单值结论，使后续前端设计可直接进入实现。

#### Scenario: 列表页信息区块判定

- **WHEN** 渲染 `Decision Center / List`
- **THEN** 页面至少必须包含列表工具栏区、列表内容区与空状态区
- **AND** 列表工具栏区必须承接搜索输入、状态筛选与进入 `Decision Create` 的入口
- **AND** 列表内容区必须至少展示 `title / status / created_at / link_count / linked_module_summary`
- **AND** 空状态区必须在无决策时引导用户进入 `Decision Create`
- **AND** 不得在列表页内联完整的 `RecordDecision` 表单

### Requirement: Decision Create 最小页面级信息区块冻结

系统 SHALL 将 `Decision Create` 的最小页面级信息区块冻结为单值结论，使后续前端设计可直接进入实现。

#### Scenario: 创建页信息区块判定

- **WHEN** 渲染 `Decision Create`
- **THEN** 页面至少必须包含结构化表单区、来源上下文区与提交取消操作区
- **AND** 结构化表单区必须承接 `title / context / problem / alternatives / choice / reason / impact / status` 字段
- **AND** 来源上下文区必须在从 `Module Detail` 带上下文进入时展示预填的 `Module` 信息
- **AND** 提交取消操作区必须承接 `RecordDecision` 提交与返回列表路径
- **AND** 不得在创建页内联 `Decision Detail` 或目标关联写入能力

### Requirement: Decision Detail 最小页面级信息区块冻结

系统 SHALL 将 `Decision Detail` 的最小页面级信息区块冻结为单值结论，使后续前端设计可直接进入实现。

#### Scenario: 详情页信息区块判定

- **WHEN** 渲染 `Decision Detail`
- **THEN** 页面至少必须包含核心字段区、已关联目标区与候选读取及目标关联区
- **AND** 核心字段区必须展示 `title / context / problem / alternatives / choice / reason / impact / status`
- **AND** 已关联目标区必须展示当前已建立的 `Decision -> Module` 关联结果
- **AND** 候选读取及目标关联区必须承接 `Decision -> Module` 候选读取与 `LinkDecisionToTarget` 写入触点
- **AND** 不得把候选读取区扩写为独立 `Module` 浏览工作台

### Requirement: PC 与移动浏览器信息密度策略冻结

系统 SHALL 在单一 `React Web` 前端交付策略下，同时定义 `PC` 与移动浏览器的基础信息密度规则。

#### Scenario: 桌面端信息密度

- **WHEN** 页面在 `PC` 桌面环境展示
- **THEN** 列表页应优先承接更高信息密度
- **AND** 详情页可在同屏承接结构化模板、已关联目标与候选关联区块

#### Scenario: 移动浏览器信息密度

- **WHEN** 页面在移动浏览器窄屏环境展示
- **THEN** 必须采用同一套页面语义与动作体系
- **AND** 必须通过信息裁剪、区块折叠、垂直重排或分层展示降低拥挤度
- **AND** 不得引入第二套独立移动端 UI 方案
- **AND** 不得要求独立 `React Native` 客户端或完整 `PWA` 作为当前阶段前提

## MODIFIED Requirements

### Requirement: Decision Center 最小职责解释

`Decision Center` 在 `phase03` 中 SHALL 不只被解释为一个“决策记录概念”，而是被解释为包含列表、创建、详情与目标关联入口在内的一条完整页面主线。

#### Scenario: 页面主线解释

- **WHEN** 后续 `/spec` 或实现讨论 `Decision Center`
- **THEN** 必须将其理解为一条完整页面主线
- **AND** 不得只剩单独列表页而把创建、详情与关联入口散落为无归属入口

## REMOVED Requirements

### Requirement: 第二套移动端 UI 方案
**Reason**: 当前阶段已冻结为单一 `React Web` 同时覆盖 `PC` 与移动浏览器，不需要并行维护第二套移动端页面方案。
**Migration**: 若后续确有独立客户端需要，必须进入新的 `phase / audit` 流程，不在 `phase03-01` 当前规格中处理。
