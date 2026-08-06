# Phase03-10 产出首份 Decision Center 正式规格文档 Spec

## Why

`phase03-01` 到 `phase03-09` 已经分别冻结了 `Decision Center` 的页面边界、结构化模板、入口上下文、数据与错误语义、前后端实现设计、最小 `.proto` 合同与联调验收基线。但这些结论当前仍分散在多个子规格中，后续 `phase03-11 / 12 / 13 / 14 / 15` 若继续直接引用这些零散规格，仍会回到多入口、多解释、实现前临场拼装的状态。

因此，当前必须产出一份单一、完整、可互链的 `Decision Center` 正式规格正文，把 `phase01-06` 的正式 MVP 上游、`phase02` 已交付结果，以及 `phase03-01 ~ 09` 的冻结结论统一收口为后续实现、验收与 `phase04` 的直接上游规格来源。

## What Changes

- 冻结 `Decision Center` 正式规格正文的目标文件落点与文档定位
- 冻结该正式规格对 `phase01-06` 正式 MVP 规格正文、`phase02` 正式规格与验收结果、`phase03-01 ~ 09` 子规格的继承关系
- 冻结页面、动作、数据读写、API、合同、验收基线、非目标、实现设计层结果与 Done 标准在正式规格中的统一承接要求
- 冻结 `Decision Center` 正式规格正文作为后续实现与 `phase04` 直接上游规格来源的定位
- 明确正式规格正文必须与根级真相源、`phase03` 三件套、`phase02` 已交付结果互链一致，不得形成第二套边界

## Impact

- Affected specs:
  - `phase03_decision_center_foundation`
  - `phase03_01_decision_center_pages_info_arch`
  - `phase03_02_decision_template_status_read_model`
  - `phase03_03_link_target_scope_entry_context`
  - `phase03_04_decision_data_api_error_boundary`
  - `phase03_05_frontend_page_route_component_design`
  - `phase03_06_frontend_state_interaction_flow`
  - `phase03_07_backend_module_boundary_interface_grouping`
  - `phase03_08_decision_center_proto_contract`
  - `phase03_09_decision_center_integration_baseline`
  - `phase02_09_module_registry_formal_spec`
  - `phase02_12_module_registry_integration_validation_acceptance`
- Affected code:
  - 当前无代码改动；影响后续 `phase03-11 / 12 / 13 / 14 / 15` 的合同、实现、验收与收口口径

## ADDED Requirements

### Requirement: Decision Center 正式规格正文必须收口到单一文档

系统 SHALL 将 `phase03` 对应的 `Decision Center` 正式规格收口到单一正文文档，而不是继续以 `phase03-01 ~ 09` 零散子规格作为长期并列入口。

#### Scenario: 正式规格正文文件落点

- **WHEN** 后续实现、验收或 `phase04` 需要引用 `Decision Center` 当前阶段正式规格
- **THEN** 必须存在单一正文文件 `phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
- **AND** 该正文文件必须能够从当前 spec 目录直接追踪
- **AND** 不得要求接手者先手工拼装 `phase03-01 ~ 09` 后才能理解当前阶段正式口径

### Requirement: 正式规格正文覆盖范围必须完整

系统 SHALL 要求 `Decision Center` 正式规格正文在同一处完整覆盖页面、动作、数据读写、API、合同、验收基线、非目标、实现设计层结果与 Done 标准。

#### Scenario: 判断正式规格正文是否覆盖完整

- **WHEN** 接手者编写或审阅 `Decision Center` 正式规格正文
- **THEN** 文档必须至少完整覆盖页面矩阵、动作矩阵、数据矩阵、最小 API 分组、合同章节、空状态与冷启动路径、异常路径前提、非目标矩阵、实现设计层结果与 Done 标准
- **AND** 不得只保留零散冻结结论、概要说明或“见各子 spec”式拼接文本

### Requirement: phase01-06、phase02 已交付结果与 phase03-01~09 结论必须统一继承

系统 SHALL 要求 `Decision Center` 正式规格正文完整继承 `phase01-06` 正式 MVP 规格正文、`phase02` 已交付的 `Module Registry` 正式规格与验收结果，以及 `phase03-01 ~ 09` 已冻结的当前阶段结论，而不是重新定义第二套边界。

#### Scenario: 判断正式规格正文是否继承前置冻结

- **WHEN** 正式规格涉及对象范围、页面职责、动作归属、数据边界、错误语义、合同边界、验收环境或实现设计结果
- **THEN** 必须以 `mvp_spec_v0.1.md` 作为当前阶段唯一执行层上游
- **AND** 必须完整承接 `module_registry_spec_v0.1.md` 与 `phase02-12` 验收结论中对 `Module Registry` 的已交付边界
- **AND** 必须完整承接 `phase03-01 ~ 09` 已冻结的当前阶段结论
- **AND** 不得绕开既有子规格扩范围、改口径或引入第二套解释

### Requirement: 页面与动作章节必须单值化

系统 SHALL 要求正式规格正文在同一处完整承接 `Decision Center / List`、`Decision Create`、`Decision Detail` 三类页面主线及其动作归属。

#### Scenario: 判断页面与动作章节是否合格

- **WHEN** 正式规格定义页面结构与动作矩阵
- **THEN** 必须明确 `Decision Center / List` 承接列表读取、筛选入口、创建入口与进入详情入口
- **AND** 必须明确 `Decision Create` 承接 `RecordDecision`
- **AND** 必须明确 `Decision Detail` 承接详情读取、候选读取与 `LinkDecisionToTarget`
- **AND** 必须明确 `Module Detail` 只承接轻量入口与上下文跳转，不扩写为第二个 `Decision` 工作台
- **AND** 不得把 `Product / Repository` 扩写为当前阶段并列主交付页面主线

### Requirement: 结构化模板与状态语义章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接 `Decision` 最小结构化模板、字段级 required / optional、`alternatives` 结构、`status` 枚举与最小读模型。

#### Scenario: 判断模板与状态章节是否合格

- **WHEN** 正式规格定义 `Decision` 模板、状态或展示模型
- **THEN** 必须明确 `title / context / problem / alternatives / choice / reason / impact / status` 的最小结构化模板
- **AND** 必须明确 `title / context / problem / choice / reason / status` 为创建必填
- **AND** 必须明确 `alternatives` 为按顺序保留的文本条目集合，不引入嵌套对象结构
- **AND** 必须明确 `status` 最小枚举为 `proposed / active / superseded / archived`
- **AND** 必须明确列表与详情的最小读模型，以及 `link_count / linked_module_summary` 的计算口径

### Requirement: 目标范围与入口上下文章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接 `LinkDecisionToTarget` 的当前阶段目标范围、`Decision -> Module` 直接闭环以及 `Module Detail` 带上下文入口规则。

#### Scenario: 判断目标范围与入口上下文章节是否合格

- **WHEN** 正式规格定义目标范围、入口上下文与正式关联动作
- **THEN** 必须明确当前阶段唯一正式目标类型是 `Module`
- **AND** 必须明确 `Product / Repository` 只保留受控连接位，不扩写为第二主线
- **AND** 必须明确从 `Module Detail` 带上下文进入 `Decision Create` 后，创建成功默认进入新建 `Decision` 的详情页
- **AND** 必须明确入口上下文中的 `Module` 在 `Decision Detail` 中作为显式待关联目标继续承接

### Requirement: 数据、API 与错误语义章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接当前阶段的最小数据矩阵、读写范围、候选读取前提、接口分组与错误语义。

#### Scenario: 判断数据与 API 章节是否合格

- **WHEN** 正式规格定义数据模型方向或 API 边界
- **THEN** 必须明确直接承接 `decisions`、`decision_links`
- **AND** 必须明确直接读取或校验 `modules`
- **AND** 必须明确最小读组为 `DecisionListRead / DecisionDetailRead / DecisionModuleCandidateRead`
- **AND** 必须明确最小写组为 `DecisionWrite / DecisionLinkWrite`
- **AND** 必须明确 `RecordDecision` 校验失败、候选空结果、资源不存在、重复关联等错误语义归属
- **AND** 不得提前冻结超出当前阶段的聚合分析接口

### Requirement: 前端实现设计结果必须进入正式规格正文

系统 SHALL 要求正式规格正文显式承接 `phase03-05 / 06` 已冻结的前端实现设计层结果，使后续前端实现不需要再次回到多个子规格拆读。

#### Scenario: 判断前端实现设计结果是否被承接

- **WHEN** 正式规格定义前端实现设计层结果
- **THEN** 必须明确页面集合、最小路由结构、URL 语义、组件树与组件职责
- **AND** 必须明确列表、创建、详情、待关联目标、候选读取与关联动作的状态模型与交互流
- **AND** 必须明确 `queryText / statusFilter` 的搜索参数承接、返回列表恢复与刷新恢复规则
- **AND** 必须明确单一 `React Web` 下同时覆盖 `PC` 与移动浏览器的布局降级策略
- **AND** 不得要求后续前端实现再额外回到多个子规格中重新拼装主线

### Requirement: 后端实现设计结果必须进入正式规格正文

系统 SHALL 要求正式规格正文显式承接 `phase03-07` 已冻结的后端实现设计层结果，使后续后端实现可直接按既定边界进入编码。

#### Scenario: 判断后端实现设计结果是否被承接

- **WHEN** 正式规格定义后端实现设计层结果
- **THEN** 必须明确 `Decision Center` 为单一后端业务模块
- **AND** 必须明确 `handler / service / repository / candidate` 的分层语义与文件落点
- **AND** 必须明确 `errors.go / types.go / validate.go / handler/response.go` 的支撑文件落点
- **AND** 必须明确 `ModuleCandidateRead` 的拥有者与接线位置由 `Decision Center candidate/` 子包定义并在应用装配点接线
- **AND** 不得把 Go HTTP/RPC 框架、ORM、SQL Builder 或数据访问层具体工具写成当前正式规格阻断项

### Requirement: 合同章节必须完整承接 Proto 主线

系统 SHALL 要求正式规格正文显式承接 `phase03-08` 已冻结的 `.proto` 合同设计结果，使后续合同落地成为机械映射。

#### Scenario: 判断合同章节是否合格

- **WHEN** 正式规格定义 `Decision Center` 合同边界
- **THEN** 必须明确 `.proto` 是当前阶段唯一合同源
- **AND** 必须明确 `psco.decision_center.v1` 包名、最小服务矩阵、核心消息结构、写组 response 语义与 RPC -> HTTP 映射矩阵
- **AND** 必须明确 `ModuleStatus` 通过跨包 import 复用，不在本地重定义
- **AND** 必须明确字段编号、`reserved` 约束与 `buf breaking` 前提

### Requirement: 验收基线章节必须完整承接

系统 SHALL 要求正式规格正文显式承接 `phase03-09` 已冻结的联调环境、重置脚本、基线种子、异常路径前提与冷启动验收路径。

#### Scenario: 判断验收基线章节是否合格

- **WHEN** 正式规格定义验收环境与联调基线
- **THEN** 必须明确环境建立顺序、脚本入口与基线 seed 分层职责
- **AND** 必须明确 `reset_decision_mainline` 的 clean-only / restore-only / 默认模式
- **AND** 必须明确从空状态到首条 `Decision` 再到正式关联 `Module` 的冷启动验收路径
- **AND** 必须明确当前阶段验收不依赖临时手工 SQL

### Requirement: 非目标与 Done 标准章节必须存在

系统 SHALL 要求正式规格正文显式包含当前阶段非目标矩阵与正式 Done 标准，而不是只保留功能边界。

#### Scenario: 判断非目标与 Done 标准章节是否存在

- **WHEN** 正式规格进入验收阶段
- **THEN** 文档必须明确 `Product / Repository` 全量主线、Dashboard 聚合反馈、自动扫描代码、自动知识图谱、独立 `React Native` 客户端、完整 `PWA` 等为当前阶段非目标
- **AND** 必须明确何时可判定 `Decision Center` 当前阶段规格足以进入 `phase03-11 / 12 / 13 / 14 / 15`
- **AND** 不得缺失验收口径或把 Done 标准分散到多个文档中

### Requirement: 正式规格正文必须成为后续直接上游

系统 SHALL 将 `Decision Center` 正式规格正文定位为后续合同落地、实现、验收、收口与 `phase04` 进入条件判断时的直接上游规格来源。

#### Scenario: 判断后续阶段是否引用直接上游

- **WHEN** `phase03-11` 到 `phase03-15` 或 `phase04` 需要引用 `Decision Center` 当前阶段边界
- **THEN** 必须优先引用 `decision_center_spec_v0.1.md`
- **AND** 不得继续并列引用多个 `phase03-01 ~ 09` 子规格作为长期正式入口
- **AND** 不得绕开正式规格正文直接以对话结论替代规格入口

### Requirement: 正式规格正文必须与根级真相源和阶段文档互链一致

系统 SHALL 要求正式规格正文与根级真相源、`phase03` 三件套及 `phase02` 已交付结果保持单值一致，不得形成第二套正式真相源。

#### Scenario: 判断正式规格正文与真相源关系

- **WHEN** 正式规格引用项目定位、技术基线、目录规则、最终共识、`Module Registry` 已交付边界或当前阶段范围
- **THEN** 必须与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致
- **AND** 必须与 `phase03_decision_center_foundation_architecture_plan.md`、`phase03_decision_center_foundation_dev_plan.md`、`phase03_decision_center_foundation_shared_baseline.md` 保持单值一致
- **AND** 必须与 `module_registry_spec_v0.1.md` 与 `phase02-12` 验收结论保持单值一致
- **AND** 不得在正式规格正文中重写一套与上述文档冲突的主结论

## MODIFIED Requirements

### Requirement: phase03 后续实现与 phase04 的规格入口

后续 `phase03-11` 到 `phase03-15` 的合同落地、实现、验收与收口，以及 `phase04` 对 `Decision Center` 的引用，SHALL 以本次产出的 `decision_center_spec_v0.1.md` 作为直接上游规格入口；而该正文自身继续以 `mvp_spec_v0.1.md`、`module_registry_spec_v0.1.md` 与 `phase02-12` 验收结果作为当前阶段已冻结上游。

#### Scenario: 后续阶段引用正式规格

- **WHEN** 后续阶段需要引用 `Decision Center` 当前阶段正式规格
- **THEN** 必须从 `decision_center_spec_v0.1.md` 进入
- **AND** 该正文仍必须显式继承 `mvp_spec_v0.1.md`、`module_registry_spec_v0.1.md` 与 `phase02-12` 验收结果的已冻结边界
- **AND** 不得把 `phase01-06`、`phase02` 与 `phase03-01 ~ 09` 并列改写成多个长期直接入口

## REMOVED Requirements

### Requirement: 以后续实现继续直接拼装 phase03-01~09 子规格
**Reason**: `phase03-10` 的目标就是把前九个子任务的冻结结论收口为单一正式规格正文，避免后续实现与验收继续依赖多文档临场拼装。
**Migration**: 后续合同落地、实现、验收与 `phase04` 若需要引用 `Decision Center` 当前阶段边界，应统一从 `decision_center_spec_v0.1.md` 进入；前置子规格仅作为该正文的冻结来源与追踪依据。
