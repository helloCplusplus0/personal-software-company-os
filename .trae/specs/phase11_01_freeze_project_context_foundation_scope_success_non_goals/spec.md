# Project Context Foundation 范围边界与成功标准冻结 Spec

## Why
`phase11` 已经完成 `/plan` 三件套，但 `phase11-01` 仍需要一个正式 `/spec` 承接位，把“本阶段到底做什么、不做什么、何时算完成”冻结成后续执行不可漂移的单值输入。这样后续 `phase11-04 ~ 10` 才不会在实现或验收时重新打开范围讨论。

## What Changes
- 冻结 `Project Context Foundation` 作为 `phase11` 的唯一主交付能力
- 冻结 `phase11` 的正式成功标准、DoD 与收口口径
- 冻结 `phase11` 的明确不做清单，阻断 MCP / CLI / agent 写回等能力偷渡进入当前阶段
- 对齐 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中关于边界、非目标、验收目标的表达

## Impact
- Affected specs: `phase11_project_context_foundation`
- Affected code: 
  - `docs/phase/phase11_project_context_foundation_architecture_plan.md`
  - `docs/phase/phase11_project_context_foundation_dev_plan.md`
  - `docs/phase/phase11_project_context_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结 phase11 唯一主交付能力
系统规格 SHALL 将 `Project Context Foundation` 冻结为 `phase11` 的唯一主交付能力，并明确其由两类主交付主体与两类必备承接结果构成：

- 主交付主体：
  - `根级上下文真相源治理`
  - `最小只读项目上下文导出`
- 必备承接结果：
  - `AGENTS` 风格 Markdown 导出
  - `PSCO` 仓库自身 dogfooding 验证

补充冻结：

- `AGENTS` 风格导出是“最小只读项目上下文导出”的正式导出形态，不是与其并列的第二主交付；
- `PSCO` 仓库自身 dogfooding 验证是当前阶段的正式验收承接位，不是并列的新业务主线；
- 执行者不得把上述四项重新解释成“两个主交付 + 两个可选增强”。

#### Scenario: 成功冻结主交付边界
- **WHEN** 执行者读取 `phase11` 三件套并进入后续 `/spec`
- **THEN** 能明确判断 `phase11` 只承接根级治理与最小只读导出
- **AND** 不会把其他候选能力误当作当前并列主交付

### Requirement: 冻结 phase11 成功标准与收口口径
系统规格 SHALL 冻结 `phase11` 的成功标准、DoD 与阶段收口口径，使后续执行者无需再临场判断“做到什么程度算完成”。

当前阶段的成功标准至少固定为：

1. `Project Context Foundation` 已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结；
2. `phase11` 的明确不做清单已在三件套中单值一致，且没有把更重协议层、agent 写回或重型集成写成当前事实；
3. `PSCO` 当前仓库样本、根级文件结构与未来跨项目通用能力之间的边界已写清，不允许把当前样本外推为统一模板；
4. 后续 `phase11-04 ~ 10` 的执行者在进入 `/spec`、实现与验收前，不需要重新判断“本阶段到底做什么、何时算完成、哪些能力必须后移”；
5. 若后续子任务需要引入 `MCP / CLI / agent 写回 / 前端对话式入口 / 四实体结构重构 / 知识图谱 / 重型 GitHub / Gitea 集成 / 主动注入`，必须判定为超出当前阶段范围，而不是在 `phase11` 内临场吸收。

当前阶段的收口口径至少固定为：

- `phase11-01` 的完成标志不是“新增了一份 spec 包”，而是三件套已经完成边界、成功标准与非目标的正式冻结并保持单值一致；
- 若三件套之间仍存在主交付定义冲突、成功标准缺口或非目标漂移，则 `phase11-01` 不得判定为完成。

#### Scenario: 成功判断阶段是否完成
- **WHEN** 执行者进入 `phase11-09` 验收或 `phase11-10` 根级同步
- **THEN** 可以依据三件套中的固定完成条件判断 `phase11` 是否达标
- **AND** 不需要重新解释阶段目标或额外补充隐含完成条件

### Requirement: 冻结 phase11 明确不做清单
系统规格 SHALL 显式列出 `phase11` 明确不做的能力集合，并要求这些非目标在三件套之间保持单值一致。

#### Scenario: 阻断后续能力偷渡
- **WHEN** 后续执行者讨论或设计 `MCP / CLI / agent 写回 / 前端对话式入口 / 知识图谱 / 重型集成`
- **THEN** 能直接判断这些内容不属于 `phase11`
- **AND** 不会把这些能力以“顺手增强”形式写入当前阶段

### Requirement: 冻结本阶段与未来候选能力的边界
系统规格 SHALL 明确区分“当前 phase11 必做内容”与“未来候选增强能力”，并禁止通过当前仓库样本、目录结构或 dogfooding 结果反向冻结未来所有项目模板。

#### Scenario: 约束未来能力不被提前写成事实
- **WHEN** 执行者利用 `PSCO` 自身仓库做 dogfooding 或治理根级入口
- **THEN** 只能将其视为当前样本与当前阶段验证对象
- **AND** 不得将现有根级文件清单或样本结果上升为未来所有项目的前置合同

## MODIFIED Requirements

### Requirement: phase11 三件套中的边界表达
`phase11` 三件套中的边界表达 SHALL 对齐为同一口径：

- 主交付主体：`根级上下文真相源治理 + 最小只读项目上下文导出`
- 必备承接结果：`AGENTS` 风格导出 + `PSCO` 仓库自身 dogfooding 验证

不得再出现“唯一主交付只由两部分组成”与“四项并列同级组成”并存的表达。

#### Scenario: 三件套单值一致
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 获取到的主交付、非目标与成功标准一致
- **AND** 不会出现一个文档允许、另一个文档排除的冲突表达

## REMOVED Requirements

### Requirement: 将更重消费通道视为当前阶段默认范围
**Reason**: `phase11` 当前目标是先完成上下文系统的边界冻结与最小只读导出，不应提前承接更重协议层或受控维护能力。
**Migration**: 将 `MCP / CLI / agent 写回 / Draft / 审批流 / 主动注入` 明确保留为未来阶段候选增强，不在 `phase11-01` 的正式范围内承接。
