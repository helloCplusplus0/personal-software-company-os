# Phase02-03 模块创建与空状态路径 Spec

## Why

`phase02-01` 已冻结页面边界，`phase02-02` 已冻结 `Module / Release` 的最小主线，但 `Module Registry` 要真正成为 `v0.1` 第一条可执行入口，还必须明确首轮用户如何从空状态进入 `CreateModule`、最小表单到底填什么、以及列表、创建、详情之间如何形成最小闭环。

## What Changes

- 冻结首轮 `CreateModule` 的最小表单字段
- 冻结从空状态进入首个模块登记的引导路径
- 冻结 `Module Registry / List -> Module Create -> Module Detail` 的最小闭环
- 冻结创建成功、创建失败与取消返回的最小路径
- 明确当前阶段不引入复杂导入、自动扫描或 AI 建议作为创建前提

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `frontend/` 中的模块创建页、空状态区、创建成功回流逻辑与列表页入口

## ADDED Requirements

### Requirement: CreateModule 最小表单冻结

系统 SHALL 将首轮 `CreateModule` 冻结为一组最小必填表单字段，用于完成模块登记主线的第一步。

#### Scenario: 最小表单字段

- **WHEN** 用户进入 `Module Create`
- **THEN** 当前阶段最小表单字段至少应包含 `name / description / status`
- **AND** 不得要求用户在创建时同步完成 `Product` 绑定、`Repository` 映射或 `Release` 登记
- **AND** 不得把超出正式规格正文的复杂字段作为创建前置条件

### Requirement: 首轮空状态引导冻结

系统 SHALL 将 `Module Registry` 的空状态冻结为引导用户完成首个模块登记的入口，而不是纯展示空白信息。

#### Scenario: 列表空状态

- **WHEN** 用户首次进入 `Module Registry / List` 且系统中尚无任何模块
- **THEN** 页面必须展示明确的空状态提示
- **AND** 必须提供进入 `Module Create` 的主入口
- **AND** 空状态文案必须围绕“先完成首个模块登记”展开

#### Scenario: 非空列表状态

- **WHEN** 系统中已存在至少一个模块
- **THEN** 列表页不再展示首轮空状态主提示
- **AND** 仍必须保留明确的 `CreateModule` 入口

### Requirement: 模块创建最小闭环冻结

系统 SHALL 将当前阶段的最小闭环冻结为 `Module Registry / List -> Module Create -> Module Detail`。

#### Scenario: 创建成功闭环

- **WHEN** 用户完成 `CreateModule`
- **THEN** 系统必须能够让该模块进入列表主线
- **AND** 用户必须能够从创建结果进入该模块的 `Module Detail`
- **AND** 不得把创建结果留在无后续承接的成功提示中断点

#### Scenario: 从列表进入创建

- **WHEN** 用户位于 `Module Registry / List`
- **THEN** 必须始终存在进入 `Module Create` 的明确入口

#### Scenario: 从创建回到列表

- **WHEN** 用户取消创建或创建失败
- **THEN** 系统必须提供返回列表的清晰路径
- **AND** 不得使用户陷入无回路的中间页面

### Requirement: 创建成功后的最小后续动作冻结

系统 SHALL 将创建成功后的下一步承接冻结为查看 `Module Detail`，而不是要求用户立即完成更多对象录入。

#### Scenario: 创建后后续动作

- **WHEN** 用户成功创建模块
- **THEN** 默认后续承接应是进入 `Module Detail` 或回到列表并可进入该模块详情
- **AND** 当前阶段允许后续再补 `Release / Product / Repository` 关联
- **AND** 不得在创建完成时强制执行复杂导入、自动扫描或 AI 推荐

### Requirement: 当前主线非目标冻结

系统 SHALL 明确 `phase02-03` 只承接手动创建与空状态路径，不引入第二条资产进入主线。

#### Scenario: 非目标判定

- **WHEN** 后续 `/spec` 或实现讨论模块创建入口
- **THEN** 不得把复杂导入写成当前主线
- **AND** 不得把自动扫描代码写成创建前提
- **AND** 不得把 AI 建议写成首轮必需入口

## MODIFIED Requirements

### Requirement: Module Registry 首轮输入路径

`Module Registry` 在当前阶段 SHALL 不仅是模块读取入口，也必须成为首轮模块登记的直接起点。

#### Scenario: 首轮输入路径解释

- **WHEN** 接手者讨论 `Module Registry` 的首轮用户路径
- **THEN** 必须把它理解为从空状态进入创建，再进入详情的最小主线
- **AND** 不得只保留读取语义而缺失首轮录入主线

## REMOVED Requirements

### Requirement: 复杂导入或 AI 建议作为首轮创建前提
**Reason**: 当前阶段的目标是先跑通手动创建与空状态闭环，而不是引入更高复杂度的资产进入方式。
**Migration**: 若后续需要轻量导入说明或 AI 增强入口，必须在后续 `phase / audit` 中单独冻结，不在 `phase02-03` 当前规格中处理。
