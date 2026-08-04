# Phase01-06 产出首份正式 MVP 规格文档 Spec

## Why

在 `phase01-01` 到 `phase01-05` 已经分别冻结技术路线、对象动作范围、页面输入路径、数据合同基线和冷启动导入导出要求之后，下一步必须把这些单点结论收束为一份正式的 MVP 规格正文。只有形成一份完整、单值、可互链的正式规格，后续实现与 `phase02` 才不会重新回到多入口、多解释和多版本并存状态。

## What Changes

- 冻结 `phase01` 对应的正式 MVP 规格正文应覆盖的完整章节范围
- 冻结该正式规格对前五个子任务结论的继承关系
- 冻结对象、动作、页面、数据、API、非目标与 Done 标准的统一承接要求
- 冻结冷启动、空状态、导入路径、导出 / 备份要求与基础度量在正式规格中的必备位置
- 冻结该正式规格作为后续实现与 `phase02` 唯一上游规格来源的定位
- 明确正式规格必须与根级真相源互链一致，不得形成第二套真相源

## Impact

- Affected specs: `phase01_mvp_spec_convergence`
- Affected code: 当前无代码改动，影响后续正式 MVP 规格正文编写、`phase02` 入口、实现顺序和验收口径

## ADDED Requirements

### Requirement: 正式 MVP 规格正文覆盖范围冻结
系统 SHALL 要求 `phase01` 对应的正式 MVP 规格正文完整覆盖对象、动作、页面、数据、API、非目标与 Done 标准。

#### Scenario: 判断正式规格是否覆盖完整
- **WHEN** 接手者编写或审阅 `phase01` 对应的正式 MVP 规格正文
- **THEN** 文档必须至少完整覆盖对象范围、动作矩阵、页面矩阵、数据结构、API 边界、非目标矩阵与 Done 标准
- **AND** 不得只保留零散冻结结论而缺少正式规格正文的统一收口

### Requirement: 前五个子任务结论统一继承
系统 SHALL 要求正式 MVP 规格正文完整继承 `phase01-01` 到 `phase01-05` 已冻结的单值结论，而不是重新定义第二套边界。

#### Scenario: 判断正式规格是否继承前置冻结
- **WHEN** 正式规格涉及技术路线、对象、动作、页面、数据、合同、冷启动、导入或导出要求
- **THEN** 必须以 `phase01-01` 到 `phase01-05` 为唯一直接上游
- **AND** 不得绕开既有子规格重新扩范围、改口径或引入第二套解释

### Requirement: 对象与动作章节必须单值化
系统 SHALL 要求正式 MVP 规格正文在同一处完整承接核心实体、派生层、后移对象与核心动作矩阵。

#### Scenario: 判断对象与动作章节是否合格
- **WHEN** 正式规格定义对象与动作
- **THEN** 必须明确 `Product / Module / Release / Decision / Repository / Venture（可选）` 的正式地位
- **AND** 必须明确 `Capability` 为派生层、`Feature / Opportunity / Experiment` 为后移对象
- **AND** 必须完整承接 `CreateProduct / CreateModule / CreateRelease / CreateRepository / RecordDecision / BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository / LinkDecisionToTarget`

### Requirement: 页面与空状态章节必须完整承接
系统 SHALL 要求正式 MVP 规格正文在同一处完整承接页面范围、页面职责、空状态与首轮冷启动路径。

#### Scenario: 判断页面与冷启动章节是否合格
- **WHEN** 正式规格定义页面、空状态和冷启动流程
- **THEN** 必须完整覆盖 `Dashboard / Module Registry / Product Registry / Decision Center / Repository Binding`
- **AND** 必须明确每个页面的最小职责与对应动作
- **AND** 必须明确零数据用户如何从空状态进入第一版可用状态

### Requirement: 数据与 API 章节必须完整承接
系统 SHALL 要求正式 MVP 规格正文在同一处完整承接最小数据模型方向、关系结构、`Decision Record` / `Repository Binding` 结构要求与 API 边界。

#### Scenario: 判断数据与 API 章节是否合格
- **WHEN** 正式规格定义表结构方向、关系结构或 API 边界
- **THEN** 必须完整承接核心表、关系表、派生视图和结构化合同前提
- **AND** 必须与 `PostgreSQL` 主线和 `Protocol Buffers` 长期方向保持一致
- **AND** 不得把具体 Go 数据访问实现选型或完整 `proto` 工具链落地写成当前规格阻断项

### Requirement: 导入与导出章节必须完整承接
系统 SHALL 要求正式 MVP 规格正文在同一处完整承接导入边界、手动录入优先级、最小导出要求与最小备份要求。

#### Scenario: 判断导入与导出章节是否合格
- **WHEN** 正式规格定义导入说明、导出能力或备份策略
- **THEN** 必须明确 `v0.1` 当前不冻结任何正式导入能力，首轮以手动录入为主
- **AND** 必须明确导出是面向用户带走核心资产数据
- **AND** 必须明确备份是面向当前实例保留与恢复
- **AND** 不得把 GitHub OAuth、自动导入、自动同步或连续备份写成首轮必需项

### Requirement: 基础度量与 Done 标准章节必须存在
系统 SHALL 要求正式 MVP 规格正文显式包含基础度量指标与正式 Done 标准，而不是只保留功能边界。

#### Scenario: 判断度量与 Done 标准是否存在
- **WHEN** 正式规格进入验收阶段
- **THEN** 文档必须明确基础度量指标
- **AND** 必须明确何时可判定 `v0.1` 规格成立并进入稳定实现
- **AND** 不得缺失验收口径或把 Done 标准分散在多个文档中

### Requirement: 正式规格成为唯一上游
系统 SHALL 将 `phase01` 对应的正式 MVP 规格正文定位为后续实现与 `phase02` 的唯一上游规格来源。

#### Scenario: 判断后续是否引用唯一上游
- **WHEN** 后续 `phase02` 或实现任务需要引用 `v0.1` MVP 规格
- **THEN** 必须优先引用 `phase01` 对应的正式规格正文
- **AND** 不得继续并列引用多个子规格作为长期正式入口

### Requirement: 与根级真相源互链一致
系统 SHALL 要求正式 MVP 规格正文与根级真相源保持互链一致，不得形成第二套正式真相源。

#### Scenario: 判断正式规格与根级真相源关系
- **WHEN** 正式规格引用项目定位、技术基线、目录规则或最终共识
- **THEN** 必须与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致
- **AND** 不得在正式规格正文中重写一套与根级真相源冲突的主结论

## MODIFIED Requirements

### Requirement: 后续 phase 与实现引用前提
后续 `phase02` 规划、实现与验收 SHALL 以本次定义的正式 MVP 规格正文为唯一上游规格入口，并把 `phase01-01` 到 `phase01-05` 视为该正文的收敛前置，而不是继续作为并列长期入口。

#### Scenario: 后续 phase 引用正式规格
- **WHEN** 后续阶段需要引用 MVP 规格边界
- **THEN** 必须从正式规格正文进入
- **AND** 不得绕开正式规格正文直接把前置子规格当作长期并列主入口
