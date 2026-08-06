# Phase03-04 Decision Center 数据读写范围、接口边界与错误语义前提 Spec

## Why

`phase03-01 ~ 03` 已经把 `Decision Center` 的页面边界、模板读模型、目标范围与入口上下文收敛到可执行状态。为了让后续前端状态设计、后端模块设计、`.proto` 合同与正式规格正文可以直接落地，当前还需要把 `Decision Center` 所需的最小数据读写范围、列表/详情/创建/关联的接口边界，以及 `RecordDecision` / `LinkDecisionToTarget` 的错误语义前提冻结成单值结论。

## What Changes

- 冻结 `Decision Center` 当前阶段最小数据读写范围
- 冻结列表读取、详情读取、创建写入、候选读取、关联写入的最小接口边界
- 冻结 `RecordDecision` 与 `LinkDecisionToTarget` 的请求校验与失败语义归属
- 冻结关键异常路径的最小错误语义前提
- 明确当前阶段不提前冻结聚合分析接口或 `Dashboard` 消费接口

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `frontend/` 中 `Decision List / Detail / Create` 的数据请求承接；后续 `backend/` 中 `Decision Center` 读写接口分组、校验逻辑与错误映射；后续 `.proto` 中读写 RPC 与消息错误语义设计

## ADDED Requirements

### Requirement: Decision Center 当前阶段数据读写范围冻结

系统 SHALL 将 `Decision Center` 当前阶段所需的数据读写范围冻结为最小可执行集合，只承接当前主线真正需要的数据对象。

#### Scenario: 当前阶段写入范围

- **WHEN** 后续 `/spec`、实现设计或代码实现讨论 `Decision Center` 的写入范围
- **THEN** 当前阶段只承接 `RecordDecision` 与 `LinkDecisionToTarget`
- **AND** 写入数据范围只对应 `decisions` 与 `decision_links`
- **AND** 不得把 `Decision -> Product`、`Decision -> Repository` 或其他目标类型提前扩写为当前阶段写入主线

#### Scenario: 当前阶段读取范围

- **WHEN** 后续 `/spec`、实现设计或代码实现讨论 `Decision Center` 的读取范围
- **THEN** 当前阶段只承接 `Decision List`、`Decision Detail`、`Decision -> Module` 候选读取与创建后的最小回流读取
- **AND** 读取前提只包含 `decisions`、`decision_links` 与 `modules`
- **AND** 创建后的最小回流读取由 `DecisionDetailRead` 承接
- **AND** 当前阶段必须承接 `phase02` 既有 `decisions` 表向结构化主线升级后的兼容读取

### Requirement: 最小接口承接前提冻结

系统 SHALL 将 `Decision Center` 的最小接口承接前提冻结为“按当前页面主线承接读写动作”，而不是提前展开完整服务矩阵。

#### Scenario: 页面到接口的最小映射

- **WHEN** 后续 `/spec` 或实现需要定义 `Decision Center` 的最小接口前提
- **THEN** `Decision Center / List` 必须承接列表读取接口与进入详情入口
- **AND** `Decision Create` 必须承接 `RecordDecision` 创建写入接口
- **AND** `Decision Detail` 必须承接详情读取、`Decision -> Module` 候选读取与 `LinkDecisionToTarget` 关联写入接口
- **AND** 不得把这些动作提前拆成并列独立业务工作台

### Requirement: 读动作接口分组前提冻结

系统 SHALL 将当前阶段读动作接口分为最小读取组，而不是提前冻结完整查询矩阵。

#### Scenario: 读动作最小接口分组

- **WHEN** 后续 `/spec`、实现设计或 `.proto` 合同讨论当前阶段读接口
- **THEN** 最小读动作接口分组至少应包含 `DecisionListRead`、`DecisionDetailRead` 与 `DecisionModuleCandidateRead`
- **AND** `DecisionListRead` 只服务列表展示、筛选入口与进入详情
- **AND** `DecisionDetailRead` 只服务详情展示、已关联目标展示与最小来源上下文承接
- **AND** `DecisionModuleCandidateRead` 只服务 `Decision Detail` 中 `Decision -> Module` 的候选读取
- **AND** 不得提前冻结完整聚合查询、趋势分析查询或跨页面反馈接口

### Requirement: 写动作接口分组前提冻结

系统 SHALL 将当前阶段写动作接口分为创建与关联两组最小写入接口，而不是提前冻结更复杂的写模型体系。

#### Scenario: 写动作最小接口分组

- **WHEN** 后续 `/spec`、实现设计或 `.proto` 合同讨论当前阶段写接口
- **THEN** 最小写动作接口分组至少应包含 `DecisionWrite` 与 `DecisionLinkWrite`
- **AND** `DecisionWrite` 只承接 `RecordDecision`
- **AND** `DecisionLinkWrite` 只承接 `LinkDecisionToTarget`
- **AND** 不得在当前阶段引入额外审批、投票、批量变更或自动化写入接口

### Requirement: RecordDecision 请求校验与失败语义归属冻结

系统 SHALL 将 `RecordDecision` 的请求校验与失败语义归属冻结为可直接进入实现的单值规则。

#### Scenario: RecordDecision 必填校验失败

- **WHEN** 用户提交 `RecordDecision` 且缺少 `title / context / problem / choice / reason / status` 中任一必填字段
- **THEN** 系统必须返回明确校验失败语义
- **AND** 不得降级为模糊通用错误
- **AND** 不得错误映射为资源不存在或重复冲突
- **AND** 错误归属必须落在 `DecisionWrite` 对应的创建写入动作

#### Scenario: RecordDecision 字段值非法

- **WHEN** 用户提交 `RecordDecision` 且 `status` 不在当前阶段冻结枚举内，或必填字段在去首尾空白后为空，或 `alternatives` 条目在去首尾空白后为空
- **THEN** 系统必须返回明确校验失败语义
- **AND** 错误归属必须落在 `DecisionWrite` 对应的创建写入动作

### Requirement: Decision -> Module 候选读取错误语义冻结

系统 SHALL 将 `Decision -> Module` 候选读取的返回语义冻结为“读结果或空结果”，而不是混入关联写入错误。

#### Scenario: 候选读取返回空结果

- **WHEN** `Decision Detail` 发起 `Decision -> Module` 候选读取且当前无可关联 `Module`
- **THEN** 系统必须返回空列表语义
- **AND** 不得错误映射为资源不存在
- **AND** 不得将空结果解释为接口失败

#### Scenario: 详情读取或候选读取前置资源不存在

- **WHEN** `Decision Detail` 发起详情读取或候选读取，但目标 `Decision` 本身不存在
- **THEN** 系统必须返回资源不存在语义
- **AND** 错误归属必须落在 `DecisionDetailRead`

### Requirement: LinkDecisionToTarget 请求校验与失败语义归属冻结

系统 SHALL 将 `LinkDecisionToTarget` 的请求校验与失败语义归属冻结为当前阶段可直接实现的最小错误模型。

#### Scenario: 关联写入目标类型越界

- **WHEN** 用户发起 `LinkDecisionToTarget` 且目标类型不是当前阶段允许的 `Module`
- **THEN** 系统必须返回明确校验失败语义
- **AND** 不得隐式降级为未来目标类型保留位
- **AND** 错误归属必须落在 `DecisionLinkWrite`

#### Scenario: 关联写入目标不存在

- **WHEN** 用户发起 `LinkDecisionToTarget`，但目标 `Decision` 或目标 `Module` 不存在
- **THEN** 系统必须返回资源不存在语义
- **AND** 错误归属必须落在 `DecisionLinkWrite`

#### Scenario: 重复关联

- **WHEN** 用户对同一 `Decision` 重复关联同一 `Module`
- **THEN** 系统必须返回重复冲突语义
- **AND** 不得错误映射为校验失败或资源不存在
- **AND** 错误归属必须落在 `DecisionLinkWrite`

### Requirement: 关键异常路径必须进入当前阶段规划

系统 SHALL 将当前阶段关键异常路径纳入规格前置冻结，而不是把异常语义留到联调阶段临时补写。

#### Scenario: 当前阶段关键异常路径

- **WHEN** 后续 `/spec`、实现设计、`.proto` 合同或验收设计讨论当前阶段异常路径
- **THEN** 至少必须显式覆盖 `RecordDecision` 必填缺失、`RecordDecision` 字段值非法、候选读取空结果、详情读取资源不存在、`LinkDecisionToTarget` 目标类型越界、`LinkDecisionToTarget` 目标不存在、`LinkDecisionToTarget` 重复关联
- **AND** 不得把这些路径后移到联调报告中再补定义

### Requirement: 非目标接口边界冻结

系统 SHALL 明确当前阶段不提前冻结超出 `Decision Center` 最小闭环的聚合分析接口。

#### Scenario: 非目标接口判定

- **WHEN** 后续 `/spec`、实现设计或代码实现讨论跨页面聚合查询、趋势分析查询、`pending_decision_signals` 消费接口或 `Dashboard` 反馈接口
- **THEN** 必须判定为当前阶段非目标
- **AND** 不得把它们写成 `phase03-04` 的必需接口前提

## MODIFIED Requirements

### Requirement: Contract First 当前阶段解释

`Contract First` 在 `phase03` 当前阶段 SHALL 被解释为“先冻结最小数据读写范围、接口边界与错误语义前提，再进入后续 `.proto` 合同设计”，而不是在接口未收敛前提前扩写完整协议矩阵。

#### Scenario: Contract First 承接方式

- **WHEN** 后续 `/spec`、实现设计或 `.proto` 合同讨论 `Contract First`
- **THEN** 必须优先沿用当前已冻结的最小读写接口分组与错误语义归属
- **AND** 不得在当前阶段额外引入与 `.proto` 单一合同源冲突的第二套接口定义体系

## REMOVED Requirements

### Requirement: 当前阶段完整聚合分析接口冻结
**Reason**: `phase03` 当前目标是建立 `Decision Center` 最小可执行闭环，而不是提前完成 `Dashboard`、趋势分析或提醒消费链的完整接口规划。
**Migration**: 若后续需要完整聚合查询、`pending_decision_signals` 消费接口或 `Dashboard` 反馈接口，应在后续 `phase / audit` 中单独冻结，不在 `phase03-04` 当前规格中处理。
