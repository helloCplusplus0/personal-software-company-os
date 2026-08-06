# Phase03-03 LinkDecisionToTarget 目标范围与入口上下文 Spec

## Why

`phase03-01` 已经冻结了 `Decision Center` 的页面边界，`phase03-02` 已经冻结了 `Decision` 模板与读模型，但 `LinkDecisionToTarget` 到底能关联哪些目标、`Decision -> Module` 如何形成当前阶段直接闭环、以及从 `Module Detail` 发起记录决策时上下文如何传递，仍需要写成单值结论。否则后续页面交互、接口设计、`.proto` 合同与验收路径仍会在“可关联什么”和“上下文从哪里来”这两个问题上继续漂移。

## What Changes

- 冻结 `LinkDecisionToTarget` 的当前阶段目标范围
- 冻结 `Decision -> Module` 作为当前阶段唯一必交付目标闭环
- 冻结 `Product / Repository` 在当前阶段的受控连接位解释
- 冻结从 `Module Detail` 发起带上下文记录决策的最小入口方式
- 冻结 `Decision Create` 与 `Decision Detail` 对入口上下文的最小承接规则
- 冻结从 `Module Detail` 带上下文创建成功后的默认回流路径与入口上下文在 `Decision Detail` 中的继续承接规则
- 明确当前阶段不把 `Product / Repository` 扩写为 `LinkDecisionToTarget` 的并列写入主线

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `frontend/` 中 `Module Detail` 到 `Decision Center` 的上下文跳转、`Decision Create` 预填上下文承接、`Decision Detail` 目标关联面板；后续 `backend/` 与 `.proto` 中 `LinkDecisionToTarget` 的目标类型与上下文字段

## ADDED Requirements

### Requirement: LinkDecisionToTarget 当前阶段目标范围冻结

系统 SHALL 将 `LinkDecisionToTarget` 的当前阶段可写入目标范围冻结为单值结论。

#### Scenario: 判断当前阶段目标范围

- **WHEN** 后续 `/spec`、前端交互、后端接口或 `.proto` 合同讨论 `LinkDecisionToTarget` 可关联的目标类型
- **THEN** 当前阶段必须将 `Module` 解释为唯一必交付的可写入目标类型
- **AND** 不得将 `Product / Repository` 扩写为当前阶段并列的写入主线
- **AND** 不得在当前阶段引入其他新目标类型

### Requirement: Decision -> Module 直接闭环冻结

系统 SHALL 将 `Decision -> Module` 冻结为 `phase03` 当前阶段唯一必须打通的目标关联闭环。

#### Scenario: 判断直接闭环

- **WHEN** 用户在 `Decision Detail` 中执行目标关联
- **THEN** 当前阶段必须能够完成 `Decision -> Module` 的候选读取、目标选择与关联写入
- **AND** 关联结果必须能够回流到当前 `Decision Detail` 的已关联目标区
- **AND** 关联结果必须能够被 `Decision List` 的 `link_count / linked_module_summary` 消费

### Requirement: Product / Repository 受控连接位冻结

系统 SHALL 将 `Product / Repository` 在当前阶段的定位冻结为受控连接位，而不是 `LinkDecisionToTarget` 的第二主线。

#### Scenario: Product / Repository 当前阶段定位

- **WHEN** 后续规格讨论 `Decision` 是否可以关联 `Product / Repository`
- **THEN** 当前阶段只允许保留合同保留位、轻量候选读取前提或未来扩展语义
- **AND** 不得要求当前阶段实现 `Decision -> Product` 或 `Decision -> Repository` 的正式写入闭环
- **AND** 不得因此扩写新的页面主线、写入动作或验收主线

### Requirement: Module Detail 记录决策入口上下文冻结

系统 SHALL 将从 `Module Detail` 发起“为当前 Module 记录决策”的入口上下文冻结为当前阶段唯一记录决策预填来源。

#### Scenario: 从 Module Detail 发起记录决策

- **WHEN** 用户在 `Module Detail` 中触发“为当前 `Module` 记录决策”
- **THEN** 必须进入带上下文的 `Decision Create`
- **AND** 入口上下文必须至少携带当前 `Module` 的目标标识与可展示名称
- **AND** 该上下文只能作为预填来源与后续关联建议依据
- **AND** 不得在 `Module Detail` 内直接完成 `LinkDecisionToTarget` 写入

### Requirement: Decision Create 入口上下文承接冻结

系统 SHALL 将 `Decision Create` 对入口上下文的承接规则冻结为单值结论。

#### Scenario: 带 Module 上下文进入创建页

- **WHEN** 用户从 `Module Detail` 带上下文进入 `Decision Create`
- **THEN** 页面必须展示该来源 `Module` 的最小上下文信息
- **AND** 页面必须将该 `Module` 解释为“待关联目标”，而不是已经完成写入的正式关联结果
- **AND** 创建成功后必须默认进入新建 `Decision` 的 `Decision Detail`，不得回流到 `Decision Center / List`
- **AND** 该入口上下文中的 `Module` 必须带入 `Decision Detail` 作为显式待关联目标继续承接

#### Scenario: 无特定目标上下文进入创建页

- **WHEN** 用户直接从 `Decision Center / List` 进入 `Decision Create`
- **THEN** 页面必须承接“无特定来源目标”的语义
- **AND** 不得伪造默认 `Module` 目标

### Requirement: Decision Detail 目标写入入口冻结

系统 SHALL 将 `Decision Detail` 冻结为当前阶段唯一承接正式 `LinkDecisionToTarget` 写入的页面上下文。

#### Scenario: 正式目标关联写入

- **WHEN** 用户准备将某条 `Decision` 正式关联到目标对象
- **THEN** 当前阶段必须在 `Decision Detail` 中完成目标选择与 `LinkDecisionToTarget` 写入
- **AND** 即使该 `Decision` 来自 `Module Detail` 带上下文创建，也不得跳过 `Decision Detail` 直接视为已建立正式关联
- **AND** 不得在 `Decision Create` 中隐式完成正式目标关联写入

#### Scenario: 入口上下文在 Decision Detail 中的继续承接

- **WHEN** 用户从 `Module Detail` 带上下文创建 `Decision` 并进入 `Decision Detail`
- **THEN** 该入口上下文中的 `Module` 必须在 `Decision Detail` 中作为显式待关联目标继续承接
- **AND** 在候选读取面板中，该 `Module` 必须作为首选候选或显式待确认目标出现，降低正式关联摩擦
- **AND** 该显式待关联状态必须持续到用户完成正式 `LinkDecisionToTarget` 或主动放弃关联
- **AND** 不得在进入 `Decision Detail` 后丢失入口上下文，迫使用户重新在候选列表中查找目标

### Requirement: 来源上下文与正式关联结果边界冻结

系统 SHALL 将“入口上下文”和“正式关联结果”解释为两个不同层级，避免后续实现把预填来源误当成已落库关联。

#### Scenario: 区分入口上下文与正式关联

- **WHEN** 用户从 `Module Detail` 带上下文进入 `Decision Create`
- **THEN** 该 `Module` 信息必须只代表入口上下文或待关联目标
- **AND** 只有在 `Decision Detail` 中完成 `LinkDecisionToTarget` 后，才能视为正式建立 `Decision -> Module` 关联
- **AND** 在正式关联写入前，`Decision List` 的 `link_count / linked_module_summary` 不得将该预填来源计入已关联结果

## MODIFIED Requirements

### Requirement: Decision Center 入口上下文解释

`Decision Center` 在当前阶段 SHALL 不只承接“进入哪个页面”，还必须承接“入口上下文来自哪里、如何与正式关联写入衔接”的单值解释。

#### Scenario: 入口上下文解释

- **WHEN** 后续 `/spec` 或实现讨论 `Decision Center` 的入口上下文
- **THEN** 必须以“`Module Detail` 带上下文进入 `Decision Create`，`Decision Detail` 承接正式写入”作为当前阶段唯一主解释
- **AND** 不得重新发明第二套并行入口上下文与写入路径

## REMOVED Requirements

### Requirement: Product / Repository 当前阶段并列写入主线
**Reason**: 当前阶段必须先打通 `Decision -> Module` 最小闭环，若同时把 `Product / Repository` 提前扩写为并列写入主线，会直接放大范围并削弱 `Decision Center` 收敛性。
**Migration**: 若后续确需实现 `Decision -> Product` 或 `Decision -> Repository` 正式写入，必须进入新的 `phase / audit` 流程单独冻结，不在 `phase03-03` 当前规格中处理。
