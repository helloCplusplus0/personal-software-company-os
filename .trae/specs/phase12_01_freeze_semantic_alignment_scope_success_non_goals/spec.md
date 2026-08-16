# phase12-01 冻结 Semantic Alignment & Read-Only Consumption Foundation 范围边界、成功标准与非目标 Spec

## Why
`phase12` 已经完成 `/plan` 三件套，但如果不先把“本阶段唯一主交付是什么、何时算完成、哪些能力明确不做”冻结成正式 `/spec` 上游，后续 `phase12-04 ~ 12` 很容易在实现或验收时重新打开范围讨论，把更重 agent 通道、结构重构或影子读模型偷渡进当前阶段。

## What Changes
- 冻结 `Semantic Alignment & Read-Only Consumption Foundation` 为 `phase12` 的唯一主交付能力
- 冻结 `phase12` 的正式成功标准、DoD 与阶段收口口径
- 冻结 `phase12` 的明确不做清单，阻断 schema 重写、MCP / CLI、agent 写回、前端对话入口与第二套 canonical API 混入当前阶段
- 对齐 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中关于边界、非目标、共享只读 owner 与验收目标的表达

## Impact
- Affected specs: `phase12_semantic_alignment_and_readonly_consumption_foundation`
- Affected code:
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结 phase12 唯一主交付能力
系统规格 SHALL 将 `Semantic Alignment & Read-Only Consumption Foundation` 冻结为 `phase12` 的唯一主交付能力，并明确其由两部分同级组成：

- `四实体语义一致性收口`
- `只读消费深化`

补充冻结：

- `四实体语义一致性收口` 只承接表达层与消费层的对齐，不承接 schema、关系主线或 canonical owner 重写；
- `只读消费深化` 只承接 `phase11` 已交付 `Project Context` 主线上的共享只读能力深化，不承接写回、审批流、MCP / CLI 或第二套业务协议出口；
- 上述两部分必须在同一 `phase12` 主交付内推进，不得被重写为“前端文案整理”和“agent 通道探索”两条彼此独立的散线。

#### Scenario: 成功冻结 phase12 主交付边界
- **WHEN** 后续执行者读取 `phase12` 三件套并进入 `phase12-04 ~ 12`
- **THEN** 能明确判断 `phase12` 只承接四实体语义一致性收口与只读消费深化
- **AND** 不会把更重 agent 通道、结构重构或第二套读模型误判为当前阶段并列主交付

### Requirement: 冻结 phase12 成功标准与收口口径
系统规格 SHALL 冻结 `phase12` 的成功标准、DoD 与阶段收口口径，使后续执行者无需再临场判断“做到什么程度才算完成”。

当前阶段的成功标准至少固定为：

1. `Semantic Alignment & Read-Only Consumption Foundation` 已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结；
2. `Product / Repository / Module / Decision` 的正式语义在 Web 端已形成稳定、可回看、可复验的表达；
3. `Dashboard / Onboarding / Daily Review / Weekly Review` 已通过共享语义摘要或固定入口回到同一套四实体解释，而不是继续依赖各切片旧文案各讲一套；
4. `phase11` 已交付的 `ProjectContextService.GetProjectContext / ExportProjectContext` 继续作为当前阶段唯一结构化 canonical owner 与 agent-facing Markdown owner，不存在与其并列的第二套跨切片 canonical 摘要合同；
5. 当前阶段的固定样本、固定入口、固定 `6` 问、固定 `repository_id` 锚点与样本解析协议已冻结为可复跑验收协议；
6. 若后续子任务需要引入 `MCP / CLI / agent 写回 / 前端对话式入口 / 四实体 schema 重写 / 第二套 canonical API / 影子状态表`，必须判定为超出当前阶段范围，而不是在 `phase12` 内临场吸收。

当前阶段的收口口径至少固定为：

- `phase12-01` 的完成标志不是“新增了一份 spec 包”，而是三件套已经完成唯一主交付能力、成功标准、非目标与验收前提的正式冻结并保持单值一致；
- 若三件套之间仍存在主交付冲突、成功标准缺口、共享只读 owner 冲突、样本解析协议缺失或非目标漂移，则 `phase12-01` 不得判定为完成。

#### Scenario: 成功判断 phase12 是否达到进入后续 `/spec` 的条件
- **WHEN** 执行者准备开始 `phase12-04 ~ 07` 设计类子任务
- **THEN** 可以直接依据三件套中的固定成功标准判断当前阶段边界是否充分冻结
- **AND** 不需要重新解释 `phase12` 的目标、完成条件或验收协议

### Requirement: 冻结 phase12 明确不做清单
系统规格 SHALL 显式列出 `phase12` 明确不做的能力集合，并要求这些非目标在三件套之间保持单值一致。

当前阶段至少必须明确以下内容属于非目标：

- `Product / Repository / Module / Decision` 的 schema 重写或关系主线重构
- `MCP` 协议层正式实现
- `CLI` 工具正式实现
- agent 自动写回、Draft 接口、审批流
- 前端对话式 agent 入口
- 把 Web 做成 agent 工作台
- 第二套 canonical API
- 影子状态表或第二套语义字段

#### Scenario: 阻断更重能力偷渡进入 phase12
- **WHEN** 后续执行者讨论更重消费通道、受控维护能力或结构重构
- **THEN** 能直接判断这些内容不属于 `phase12`
- **AND** 不会把这些能力以“顺手增强”或“为了提高消费效率”形式写入当前阶段

### Requirement: 冻结 phase12 的共享只读 owner 与验收协议前提
系统规格 SHALL 要求 `phase12` 在范围边界冻结阶段就明确共享只读 owner 与验收协议前提，避免后续设计和验收再次长出临时入口或第二套解释。

当前阶段至少固定如下前提：

- `ProjectContextService.GetProjectContext` 是唯一结构化 canonical owner；
- `ProjectContextService.ExportProjectContext` 是 agent-facing Markdown 导出 owner；
- `frontend/src/features/project-context/` 是唯一允许的新 Web 跨切片共享只读 owner；
- `repository_id` 是唯一结构化输入锚点；
- `product_id / module_id / decision_id` 只允许从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析；
- 名称解析失败、结果不唯一或无法回到同一 `repository_id` 时，验收必须直接失败，不允许通过额外入口补救。

#### Scenario: 成功冻结共享只读 owner 与样本解析前提
- **WHEN** 执行者设计 `phase12-05 ~ 07` 的共享只读入口、字段或 resolver
- **THEN** 必须以前述 owner 与样本解析协议为前提
- **AND** 不会在前端页面、临时脚本或额外服务中再长出第二套结构化事实源

## MODIFIED Requirements

### Requirement: phase12 三件套中的边界表达
`phase12` 三件套中的边界表达 SHALL 对齐为同一口径：

- 唯一主交付能力：`Semantic Alignment & Read-Only Consumption Foundation`
- 同级组成：`四实体语义一致性收口 + 只读消费深化`
- 当前阶段非目标：结构重构、写回通道、重协议出口、对话式 agent 工作台与第二套事实源
- 当前阶段验收前提：固定样本、固定入口、固定 `6` 问、固定 `repository_id` 锚点与固定样本解析协议

不得再出现：

- 一个文档只强调“语义一致性收口”，另一个文档又把“只读消费深化”写成可选增强；
- 一个文档允许 `frontend/src/features/project-context/` 作为共享 owner，另一个文档又允许页面本地摘要成为并列事实源；
- 一个文档冻结固定样本协议，另一个文档仍允许通过额外入口补齐验收答案。

#### Scenario: 三件套单值一致
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 获取到的主交付、成功标准、共享只读 owner、样本解析协议与非目标口径一致
- **AND** 不会出现一个文档允许、另一个文档排除的冲突表达

## REMOVED Requirements

### Requirement: 将更重 agent 消费通道视为当前阶段默认范围
**Reason**: `phase12` 当前目标是先完成语义一致性与共享只读消费的正式冻结，不应提前承接更重协议层、对话入口或受控维护能力。
**Migration**: 将 `MCP / CLI / agent 写回 / Draft / 审批流 / 前端对话式入口 / 第二套 canonical API` 明确保留为后续阶段候选增强，不在 `phase12-01` 的正式范围内承接。
