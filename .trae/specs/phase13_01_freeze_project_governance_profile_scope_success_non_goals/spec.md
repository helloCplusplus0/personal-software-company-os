# Phase13-01 冻结 Project Governance Profile Foundation 的范围边界、成功标准与非目标 Spec

## Why

`phase13_project_governance_profile_foundation` 已完成 `/plan` 三件套，但后续进入 `/spec`、实现与验收前，仍需要把本阶段“到底做什么、不做什么、何时算完成”进一步冻结成可直接执行的正式规格，避免执行者再次把 Git 推进跟踪、模板仓库接入、自动同步或 agent 写回混入本阶段主线。

本次 `/spec` 的目标不是新增实现，而是把 `phase13` 的单一主交付能力、成功标准、非目标、验收口径与进入条件压成单值规格，让后续执行者不再需要临场猜测。

## What Changes

- 冻结 `phase13` 的单一主交付能力为 `Project Governance Profile Foundation`
- 冻结 `PSCO-native facts / IDE-accessible context / controlled synced projection` 三层边界在本阶段的适用方式
- 冻结本阶段成功标准、正式完成条件、固定验收关注点与后续进入条件
- 冻结本阶段明确不做 Git 推进跟踪、模板仓库接入、自动同步、MCP / CLI / agent 写回、目录全文扫描入库与第五个业务主实体

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - 后续 `phase13-02 ~ phase13-12` 子任务的进入前提
- Affected code:
  - 无直接源代码改动
  - 直接影响后续 `/spec` 与实现的文档边界：
    - `docs/phase/phase13_project_governance_profile_foundation_architecture_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_dev_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结 Phase13 单一主交付能力

系统 SHALL 将 `phase13` 的单一主交付能力冻结为：

`项目级治理画像 + 全局规范资产 + agent 项目简报输入`

并明确这三部分共同构成 `Project Governance Profile Foundation`，不得被拆解为并列的第二主线，也不得在后续执行中临时扩写为 Git 推进跟踪、模板平台或 IDE 目录扫描能力。

#### Scenario: 冻结 phase13 主目标

- **WHEN** 执行者读取 `phase13` 的 `/plan` 三件套与本 spec
- **THEN** 必须将本阶段理解为“先让 PSCO 正式管理项目级治理画像、全局规范资产与 agent 项目简报输入”
- **AND** 不得再把 `phase13` 理解为“同时建设 Git 推进跟踪、模板接入、自动同步或 agent 写回”

### Requirement: 冻结本阶段范围边界

系统 SHALL 将以下信息层级边界冻结为本阶段正式执行边界：

- `PSCO-native facts`：本阶段正式承接对象
- `IDE-accessible context`：继续由 IDE / agent 现场读取，不默认上升为 PSCO 正式事实
- `Controlled synced projection`：只作为后续进入条件，不在本阶段起点直接实现

#### Scenario: 判断某项能力是否属于 phase13

- **WHEN** 后续执行者评估某项候选能力是否进入 `phase13`
- **THEN** 若该能力属于 Git 推进摘要、模板来源自动接入、自动存在性校验、自动状态建议等 `Controlled synced projection`
- **THEN** 应判定其不属于 `phase13-01` 当前要交付的正式范围

### Requirement: 冻结本阶段成功标准与正式完成条件

系统 SHALL 在 `phase13-01` 中明确 `phase13` 的成功标准至少覆盖以下内容：

1. PSCO 已能正式承接项目级治理画像，而不新增第五个业务主实体
2. 当前项目范式 v1 与 canonical 根级文件集合已被结构化承接
3. 全局规范资产已形成结构化摘要与入口关系的正式承接策略
4. agent 项目简报输入已具备单一 schema 与单一解析协议
5. 第一版前端正式承接位已冻结为 `Repository detail`
6. `Web / agent` 仍共享同一套 `PSCO-native facts`
7. 本阶段明确不做的能力未被偷渡进主线

#### Scenario: 判断 phase13 是否可以进入后续实现

- **WHEN** 执行者准备从 `phase13-01` 进入后续 `/spec` 与实现子任务
- **THEN** 必须能用单值标准回答“做什么、做到什么算完成、哪些事情明确不做”
- **AND** 若仍需要临场决定主交付能力、前端承接位、agent brief 边界或是否允许进入 Git/模板/自动同步
- **THEN** 视为 `phase13-01` 规格未完成

### Requirement: 冻结本阶段非目标与禁止事项

系统 SHALL 将以下事项冻结为 `phase13` 的明确非目标与禁止事项：

1. Git 推进跟踪主线
2. 模板仓库 bootstrap / clone / pull 自动化
3. 目录或文件全文自动扫描入库
4. MCP / CLI / agent 写回 / Draft / 审批流
5. IDE 插件化、会话劫持或开发流程控制台
6. 第五个业务主实体
7. 第二套与四实体并列的事实源
8. 继续放大 `phase12` 的共享只读 UI 表达

#### Scenario: 评审实现提案是否越界

- **WHEN** 后续执行者提出把 Git 推进状态、模板仓库、目录全文、继续放大 `phase12` 共享只读 UI 表达或 agent 写回直接加入 `phase13`
- **THEN** 本 spec 必须将其判定为越出 `phase13-01` 已冻结范围
- **AND** 只能在 `phase13` 正式收口后，作为下一阶段进入条件重新讨论

### Requirement: 冻结后续阶段进入条件

系统 SHALL 明确 `phase13` 正式收口前，后续只允许继续推进：

- 项目级治理画像
- 全局规范资产
- agent 项目简报输入

而 Git 推进跟踪、模板仓库接入、自动同步与更重受控维护能力，只允许在 `phase13` 正式收口后，再依据新条件讨论或进入。

#### Scenario: 评估 phase13 之后的扩展方向

- **WHEN** 执行者评估 `phase13` 收口后的下一阶段能力
- **THEN** 应把 Git 推进跟踪、模板仓库接入、自动同步与更重受控维护能力视为后续候选
- **AND** 不得将这些能力回写成 `phase13` 的既成事实

## MODIFIED Requirements

### Requirement: Phase13 规划文档的执行约束

`phase13_project_governance_profile_foundation` 的 `/plan` 三件套 MUST 被解释为本阶段后续 `/spec` 与实现的单一边界来源；若 `architecture_plan`、`dev_plan`、`shared_baseline` 之间出现冲突，以“单一主交付能力 + 三层边界 + 非目标冻结 + 固定验收口径”的一致性为优先修正方向。

## REMOVED Requirements

### Requirement: 把 phase13 解释为更重同步与自动化阶段

**Reason**: 该解释与当前共识冲突，会让 `phase13` 重新滑向目录扫描器、模板平台或 Git 跟踪平台。

**Migration**: 将 Git 推进跟踪、模板仓库接入、自动同步与更重受控维护能力统一下沉为 `phase13` 正式收口后的下一阶段进入条件，不再视为本阶段默认交付内容。
