# Phase03-02 Decision 模板、状态语义与最小展示模型 Spec

## Why

`phase03-01` 已经冻结了 `Decision Center` 的页面边界与信息结构，但页面要真正可进入后续 `/spec` 与实现，还必须把 `Decision` 到底记录哪些字段、`status` 如何表达、列表与详情最少展示什么写成单值结论。否则后续前端、后端、`.proto` 与验收会继续在字段口径、状态解释和展示模型之间来回漂移。

## What Changes

- 冻结 `Decision` 的最小结构化模板字段集合
- 冻结 `Decision` 字段级 `required / optional` 规则
- 冻结 `alternatives` 的最小结构与输入约束
- 冻结 `Decision` 的最小 `status` 枚举与状态语义
- 冻结 `RecordDecision` 的最小创建校验前提
- 冻结 `Decision List` 的最小展示字段集合、`link_count / linked_module_summary` 计算口径与无关联时空值语义
- 冻结 `Decision Detail` 的最小展示字段集合与最小来源上下文承接要求
- 明确当前阶段不引入超出 `v0.1` 的复杂审批、投票或自动化字段

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `frontend/` 中的 `Decision List / Create / Detail` 展示模型、后续 `backend/` 中的 `Decision` DTO 与校验模型、后续 `.proto` 中的 `Decision` 消息字段

## ADDED Requirements

### Requirement: Decision 最小结构化模板冻结

系统 SHALL 将 `Decision` 的最小结构化模板冻结为当前阶段唯一记录模板。

#### Scenario: 判断最小结构化模板

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `Decision` 字段
- **THEN** 必须至少承接 `title / context / problem / alternatives / choice / reason / impact / status`
- **AND** 不得减少上述字段，导致 `Decision` 退化为无法结构化检索的自由备注
- **AND** 不得在当前阶段额外引入审批流、投票流或自动推荐类字段作为前置要求

### Requirement: Decision 字段级必填规则冻结

系统 SHALL 将 `Decision` 模板中的字段级 `required / optional` 规则冻结为单值结论。

#### Scenario: 判断字段级 required / optional

- **WHEN** 用户执行 `RecordDecision`
- **THEN** `title / context / problem / choice / reason / status` 必须为必填字段
- **AND** `alternatives / impact` 必须为可选字段
- **AND** 必填字段在去首尾空白后不得为空字符串
- **AND** 当前阶段不得用隐式默认值替代必填字段输入

### Requirement: alternatives 最小结构冻结

系统 SHALL 将 `alternatives` 冻结为按顺序保留的可重复文本条目集合，而不是嵌套对象结构。

#### Scenario: 判断 alternatives 数据结构

- **WHEN** 前后端或合同层定义 `alternatives`
- **THEN** 必须将其解释为按输入顺序保留的文本条目集合
- **AND** 不得扩写为包含 `label / score / vote / source` 等嵌套对象结构
- **AND** 条目允许为空集合
- **AND** 单个条目在去首尾空白后不得为空字符串

### Requirement: Decision 状态枚举冻结

系统 SHALL 将当前阶段 `Decision` 的最小 `status` 枚举冻结为与上游基线一致的状态集合。

#### Scenario: 判断状态范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `Decision.status`
- **THEN** 当前阶段只允许使用 `proposed / active / superseded / archived`
- **AND** 不得在当前阶段额外引入 `draft / pending_approval / rejected / auto_generated` 等新状态

### Requirement: Decision 状态语义冻结

系统 SHALL 为当前阶段的每个 `Decision.status` 提供页面、数据与合同都可复用的最小语义解释。

#### Scenario: proposed 状态语义

- **WHEN** `Decision.status = proposed`
- **THEN** 必须表示该决策已经被记录，但尚未成为当前生效结论

#### Scenario: active 状态语义

- **WHEN** `Decision.status = active`
- **THEN** 必须表示该决策是当前仍然生效或正在执行的结论

#### Scenario: superseded 状态语义

- **WHEN** `Decision.status = superseded`
- **THEN** 必须表示该决策曾经生效，但已被后续决策替代

#### Scenario: archived 状态语义

- **WHEN** `Decision.status = archived`
- **THEN** 必须表示该决策已归档保留，不再作为当前执行结论

### Requirement: RecordDecision 最小校验前提冻结

系统 SHALL 将 `RecordDecision` 的最小创建校验前提冻结为可直接进入实现的单值规则。

#### Scenario: 创建成功前提

- **WHEN** 用户提交 `RecordDecision`
- **THEN** 必填字段必须全部满足最小非空校验
- **AND** `status` 必须属于当前阶段冻结的最小状态枚举
- **AND** `alternatives` 若存在则每个条目都必须满足最小非空校验

#### Scenario: 创建失败前提

- **WHEN** 用户提交缺少必填字段、非法 `status` 或包含空白 `alternatives` 条目的记录
- **THEN** 系统必须返回明确校验失败语义
- **AND** 不得降级为模糊通用错误

### Requirement: Decision List 最小展示模型冻结

系统 SHALL 将 `Decision Center / List` 的最小展示字段与关键字段语义冻结为当前阶段唯一列表读模型，与 `shared_baseline §5.1` 列表读模型计算口径保持一致。

#### Scenario: 列表页最小展示字段

- **WHEN** 用户在 `Decision Center / List` 查看决策列表
- **THEN** 每个列表项至少必须展示 `title / status / created_at / link_count / linked_module_summary`
- **AND** 不得要求用户进入详情页才能理解最基本的决策状态
- **AND** 当前阶段不得在列表页扩写完整结构化模板全文展示

#### Scenario: link_count 计算口径冻结

- **WHEN** 列表读模型计算 `link_count`
- **THEN** 当前阶段仅统计 `decision_links` 中已建立的 `Decision -> Module` 有效关联数
- **AND** 不得混入 `Decision -> Product / Repository` 等未冻结目标类型的关联计数

#### Scenario: linked_module_summary 计算口径冻结

- **WHEN** 列表读模型计算 `linked_module_summary`
- **THEN** 仅基于已关联 `Module` 生成，不混入 `Product / Repository`
- **AND** 必须按 `module_name` 升序取前 `3` 个名称
- **AND** 若超出 `3` 个，必须在摘要末尾附加 `+N` 表示超出数量

#### Scenario: 无关联时空值语义冻结

- **WHEN** `Decision` 当前没有任何已关联 `Module`
- **THEN** `link_count` 必须返回 `0`
- **AND** `linked_module_summary` 必须返回空字符串
- **AND** 不得返回 `null` 或将空结果误报为接口错误

### Requirement: Decision Detail 最小展示模型冻结

系统 SHALL 将 `Decision Detail` 的最小展示字段与承接要求冻结为当前阶段唯一详情读模型，与 `shared_baseline §5.1` 详情读模型要求保持一致。

#### Scenario: 详情页最小展示字段

- **WHEN** 用户进入 `Decision Detail`
- **THEN** 页面至少必须展示 `title / context / problem / alternatives / choice / reason / impact / status / created_at`
- **AND** 必须展示当前已建立的目标关联结果
- **AND** 必须展示最小来源上下文
- **AND** 必须为后续 `Decision -> Module` 候选读取与 `LinkDecisionToTarget` 预留承接空间
- **AND** 不得要求用户跳转其他独立主线才能理解这条决策的当前含义

#### Scenario: 最小来源上下文承接要求

- **WHEN** 渲染 `Decision Detail` 的来源上下文区
- **THEN** 必须承接 `Decision` 创建时的最小来源上下文信息
- **AND** 当 `Decision` 从 `Module Detail` 带上下文发起时，必须能够展示该来源 `Module` 信息
- **AND** 当 `Decision` 从 `Decision Center / List` 直接发起时，必须能够表达"无特定来源目标"的语义
- **AND** 来源上下文的具体字段结构与入口上下文冻结由 `phase03-03` 承接，本阶段不提前定义

## MODIFIED Requirements

### Requirement: Decision Create 表单字段解释

`Decision Create` 在当前阶段 SHALL 不只被解释为“任意文本输入页”，而必须被解释为承接冻结后的 `Decision` 最小结构化模板与字段级校验规则的页面级表单。

#### Scenario: 创建页字段解释

- **WHEN** 后续 `/spec` 或实现讨论 `Decision Create`
- **THEN** 必须以本次冻结的最小结构化模板、`required / optional` 规则与状态语义为唯一上游
- **AND** 不得重新发明第二套字段定义或状态解释

## REMOVED Requirements

### Requirement: 复杂审批、投票或自动化字段前置
**Reason**: 当前阶段只需要支撑 `Decision Center` 最小可执行闭环，不需要引入超出 `v0.1` 的复杂治理字段。
**Migration**: 若后续确需更复杂的审批、投票、自动化建议或评分结构，必须进入新的 `phase / audit` 流程单独冻结，不在 `phase03-02` 当前规格中处理。
