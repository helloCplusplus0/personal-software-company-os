# 落实根级上下文真相源治理 Spec

## Why
`phase11-04` 已经冻结了根级上下文真相源治理的设计边界，但如果不把这些设计真正落实到 `PSCO` 自身仓库的根级入口文档，重复 phase 状态、重复目录落点、重复共识入口与悬空引用仍会继续存在。`phase11-06` 需要把治理矩阵从“可指导实现的设计输入”推进为“逐文件审计并落地完成的一次性校准结果”。

## What Changes
- 落实根级治理矩阵到 `PSCO` 自身仓库的目标入口文件
- 回收根级文档之间重复承载的 phase 状态、目录落点与最终共识入口
- 清理治理矩阵目标文件范围内指向不存在文件 `PSCO-summarize-feedback.md` 的引用
- 将当前阶段 `PSCO` 自身仓库的有效最终共识入口收敛为单值入口；在 `phase11-06` 当前时点，该入口为 `PSCO-mvp05-summarize-feedback.md`
- 冻结逐文件审计与结果记录要求，包括“已审计 / 是否修改 / 不修改原因”
- 冻结 `phase11-06` 的成功标准、DoD 与收口口径

## Impact
- Affected specs: `phase11_project_context_foundation`
- Affected code:
  - `README.md`
  - `AGENTS.md`
  - `plan.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `docs/phase/README.md`
  - `project_rules.md`
  - `global_skills.md`
  - `PSCO-mvp05-summarize-feedback.md`
  - `docs/phase/phase11_project_context_foundation_dev_plan.md`
  - `docs/phase/phase11_project_context_foundation_architecture_plan.md`
  - `docs/phase/phase11_project_context_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 落实根级治理矩阵到目标入口文件
系统规格 SHALL 将 `phase11-04` 已冻结的根级治理矩阵落实到 `PSCO` 自身仓库的目标入口文件，并且仅针对以下文件逐项审计与按需同步：

- `README.md`
- `AGENTS.md`
- `plan.md`
- `architecture_map.md`
- `docs/README.md`
- `docs/phase/README.md`
- `project_rules.md`
- `global_skills.md`
- `PSCO-mvp05-summarize-feedback.md`

补充冻结：

- 本子任务只治理 `PSCO` 自身仓库入口；
- 不定义未来消费侧项目必须具备哪些同名文件；
- 不允许只修改部分显眼入口后，就把根级治理判定为完成。

#### Scenario: 成功限定实施范围
- **WHEN** 后续执行者开始落实根级真相源治理
- **THEN** 能明确只需逐项审计并按需同步上述目标文件
- **AND** 不会把当前仓库的入口文件清单误写成未来所有项目的固定合同

### Requirement: 回收重复主结论与单值化当前有效最终共识入口
系统规格 SHALL 回收根级入口文档之间重复承载的 phase 状态、目录落点与最终共识正文，使主结论重新收敛到各自唯一正式承接位。

补充冻结：

- `plan.md` 继续作为阶段状态与推进路线的唯一正式承接位；
- `architecture_map.md` 继续作为目录结构与文档落点的唯一正式承接位；
- 当前阶段 `PSCO` 自身仓库的有效最终共识入口继续收敛为单值入口；在 `phase11-06` 当前时点，该入口为 `PSCO-mvp05-summarize-feedback.md`；
- `README.md`、`AGENTS.md`、`docs/README.md`、`docs/phase/README.md`、`global_skills.md` 等入口文件只允许保留摘要式引用或受控跳转，不得继续并列承载完整主结论正文。
- 该入口的具体文件名只代表 `PSCO` 当前阶段的有效共识文档，不上升为未来版本推进或其他项目的固定文件合同，也不作为通用 agent 上下文机制的长期锚点。

#### Scenario: 成功回收重复承载
- **WHEN** 后续执行者逐项同步根级入口文档
- **THEN** 能直接判断哪些正文必须保留在唯一正式承接位
- **AND** 不会继续保留第二个 phase 状态正文、第二个目录落点正文或第二个最终共识正文入口

### Requirement: 清理目标文件范围内的悬空引用并统一共识入口指向
系统规格 SHALL 清理治理矩阵目标文件范围内所有指向不存在文件 `PSCO-summarize-feedback.md` 的引用，并统一改写为当前有效入口或直接删除失效引用。

补充冻结：

- 若引用语义指向“根级最终共识”，必须统一收敛到当前有效最终共识入口；在 `phase11-06` 当前时点，该入口为 `PSCO-mvp05-summarize-feedback.md`；
- 不允许保留任何已知失效文件引用作为历史兼容占位；
- “悬空引用清零”是 `phase11-06` 的硬性完成条件之一，但完成范围以治理矩阵目标文件为准。

#### Scenario: 成功清理悬空引用
- **WHEN** 后续执行者完成根级入口治理同步
- **THEN** 治理矩阵目标文件中不再存在指向 `PSCO-summarize-feedback.md` 的失效引用
- **AND** 最终共识类表达都能指向当前唯一有效入口

### Requirement: 逐文件审计与结果记录
系统规格 SHALL 要求 `phase11-06` 对治理矩阵中的目标文件逐项审计，并记录每个文件的治理结果。

每个目标文件至少必须记录：

- `已审计`
- `是否修改`
- `不修改原因`（如适用）

补充冻结：

- 允许出现“无需改动”的结果；
- 但即使无需改动，也必须留下逐项审计结论；
- 不允许以“只改了最明显的几个入口”替代完整审计。

#### Scenario: 成功记录逐文件治理结果
- **WHEN** 后续执行者完成某个目标文件的治理判断
- **THEN** 能明确记录该文件是否已审计、是否修改以及不修改原因
- **AND** 后续验收者可以据此判断治理矩阵是否已被完整执行

### Requirement: 同步活动入口并保持单值一致
系统规格 SHALL 在本次实施中同步治理矩阵中已冻结的活动入口，包括 `README.md` 与 `global_skills.md`，确保新接手 agent 从根级入口读到的上下文保持单值一致。

#### Scenario: 新接手 agent 成功读取单值上下文
- **WHEN** 新接手 agent 从根级入口开始恢复项目上下文
- **THEN** 不会读到互相冲突的 phase 状态、目录落点或最终共识入口
- **AND** 可以通过根级入口摘要与受控跳转恢复到同一套正式结论

### Requirement: 冻结 phase11-06 成功标准与收口口径
系统规格 SHALL 冻结 `phase11-06` 的成功标准、DoD 与子任务收口口径，使后续执行者无需再临场判断“根级治理落实到什么程度才算完成”。

当前子任务的成功标准至少固定为：

1. 治理矩阵中的目标文件已全部完成逐项审计，并形成逐文件结果记录；
2. 根级入口不再互相复制主结论，phase 状态、目录落点与最终共识入口已重新收敛到唯一正式承接位；
3. 治理矩阵目标文件范围内指向不存在文件 `PSCO-summarize-feedback.md` 的悬空引用已清零；
4. 当前阶段 `PSCO` 自身仓库的有效最终共识入口已完成单值化；在 `phase11-06` 当前时点，该入口为 `PSCO-mvp05-summarize-feedback.md`；
5. 新接手 agent 从根级入口读到的上下文已保持单值一致，而不是依赖额外补猜或多份冲突文档交叉求解。

当前子任务的收口口径至少固定为：

- `phase11-06` 的完成标志不是“修改了几个根级文件”，而是治理矩阵中的目标文件已经全部审计完毕，目标文件范围内的悬空引用清零，且根级入口重新达到单值一致；
- 若仍存在未审计目标文件、仍存在治理矩阵目标文件范围内的 `PSCO-summarize-feedback.md` 引用、仍存在并列最终共识入口、或根级入口之间仍相互复制主结论，则 `phase11-06` 不得判定为完成。

#### Scenario: 成功判断子任务是否完成
- **WHEN** 后续执行者进入实现验收、handoff 或阶段收口
- **THEN** 可以依据逐文件审计结果与根级入口实际状态直接判断 `phase11-06` 是否达标
- **AND** 不需要重新争论“治理矩阵到底有没有落实完”

## MODIFIED Requirements

### Requirement: phase11 根级治理从设计输入进入落实状态
`phase11` 中根级治理相关表达 SHALL 从“冻结治理设计”推进为“按治理矩阵落实到目标入口文件并完成逐项审计”。

#### Scenario: 设计与实施闭环
- **WHEN** 读者对照 `phase11-04` 与 `phase11-06`
- **THEN** 能看到前者负责冻结治理设计，后者负责落实到目标根级入口
- **AND** 不会把 `phase11-06` 误解为再次讨论治理原则，而不是实施校准

## REMOVED Requirements

### Requirement: 允许根级入口继续并列承载最终共识或保留悬空引用
**Reason**: 这会直接破坏根级上下文真相源治理的单值一致性目标。
**Migration**: 将所有并列最终共识正文、失效共识入口引用与重复主结论正文回收回唯一正式承接位或改写为摘要式引用。
