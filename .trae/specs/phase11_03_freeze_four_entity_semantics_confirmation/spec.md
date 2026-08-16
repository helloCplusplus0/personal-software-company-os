# 四实体语义确认口径冻结 Spec

## Why
`phase11` 已经进入“项目上下文系统”冻结阶段，但 `Product / Repository / Module / Decision` 这四个 canonical 实体如果只停留在松散共识层，后续 `/spec` 与实现仍可能把“语义澄清”误做成“结构重构”。`phase11-03` 需要把四实体的正式语义、当前阶段解释与明确不做事项冻结成不可漂移的单值输入。

## What Changes
- 冻结 `Product / Repository / Module / Decision` 的正式语义说明
- 冻结当前阶段四实体只做语义确认、不做 schema 重构或实体关系重写的边界
- 冻结 `Module` 与 `Decision` 在当前阶段的解释口径
- 冻结 `phase11-03` 的成功标准、DoD 与收口口径
- 对齐 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中关于四实体语义的表达

## Impact
- Affected specs: `phase11_project_context_foundation`
- Affected code:
  - `docs/phase/phase11_project_context_foundation_dev_plan.md`
  - `docs/phase/phase11_project_context_foundation_architecture_plan.md`
  - `docs/phase/phase11_project_context_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结四实体正式语义说明
系统规格 SHALL 将 `Product / Repository / Module / Decision` 在 `phase11` 当前阶段的正式语义冻结为单值口径：

- `Product`：经营目标与交付容器；
- `Repository`：代码仓库身份对象与项目锚点；
- `Module`：可复用能力资产，允许后置提炼；
- `Decision`：规则、约束、选择与依据的索引对象。

#### Scenario: 成功读取四实体语义
- **WHEN** 后续执行者读取 `phase11` 三件套并进入后续 `/spec`
- **THEN** 能直接引用四实体的正式语义说明
- **AND** 不需要再临场补猜实体含义或重新发明第二套定义

### Requirement: 冻结当前阶段只做语义确认
系统规格 SHALL 冻结 `phase11-03` 当前只承接四实体语义澄清，不承接 schema 重构、实体拆并、关系主线重写或第二套实体体系引入。

#### Scenario: 阻断语义确认偷渡为结构重构
- **WHEN** 后续执行者讨论四实体字段、表结构、聚合关系或实体边界调整
- **THEN** 应直接判断 `phase11-03` 的职责是语义确认
- **AND** 不得以“先明确语义”为名把结构重构写成当前阶段事实

### Requirement: 冻结 Module 与 Decision 的当前阶段解释
系统规格 SHALL 明确 `Module` 与 `Decision` 在当前阶段的解释边界：

- `Module` 当前代表可复用能力资产，允许在后续真实复用沉淀中继续提炼，但当前阶段不要求重写其 schema、层级或注册主线；
- `Decision` 当前代表规则、约束、选择与依据的索引对象，用于支撑项目上下文恢复与只读导出，不在本子任务内扩写为新的审批流、流程引擎或结构重构入口。

#### Scenario: 成功约束 Module 与 Decision 的阶段解释
- **WHEN** 后续执行者设计上下文读取、导出、回看入口或相关 `/spec`
- **THEN** 能明确 `Module` 与 `Decision` 当前阶段分别承接什么语义
- **AND** 不会把它们误解为“马上要重写的数据结构主线”

### Requirement: 冻结 phase11-03 成功标准与收口口径
系统规格 SHALL 冻结 `phase11-03` 的成功标准、DoD 与子任务收口口径，使后续执行者无需再临场判断“语义确认做到什么程度才算完成”。

当前子任务的成功标准至少固定为：

1. 四实体正式语义说明已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结；
2. `phase11-03` 已明确四实体当前只做语义确认，不做 schema 重构、实体拆并或关系主线重写；
3. `Module` 与 `Decision` 的当前阶段解释已在三件套中可直接引用，不需要后续执行者二次解释；
4. 后续 `/spec`、实现与验收不会再把“四实体语义确认”误解为“应该顺手重构 schema”；
5. 若后续子任务讨论实体重命名、表结构重写、关系大改或新增第二套实体体系，必须能直接判定为超出 `phase11-03` 当前冻结范围。

当前子任务的收口口径至少固定为：

- `phase11-03` 的完成标志不是“新增了一份 spec 包”，而是三件套已经完成四实体语义、阶段解释与非结构重构边界的正式冻结并保持单值一致；
- 若三件套之间仍存在四实体语义漂移、`Module / Decision` 解释冲突，或仍给 schema 重构留出当前阶段偷渡空间，则 `phase11-03` 不得判定为完成。

#### Scenario: 成功判断子任务是否完成
- **WHEN** 后续执行者进入后续 `/spec`、实现、验收或 handoff
- **THEN** 可以依据三件套中的固定完成条件直接判断 `phase11-03` 是否达标
- **AND** 不需要重新解释四实体含义、阶段解释或非目标边界

## MODIFIED Requirements

### Requirement: phase11 三件套中的四实体语义表达
`phase11` 三件套中的四实体语义表达 SHALL 对齐为同一口径：

- `Product` 是经营目标与交付容器；
- `Repository` 是代码仓库身份对象与项目锚点；
- `Module` 是可复用能力资产，允许后置提炼；
- `Decision` 是规则、约束、选择与依据的索引对象；
- 当前阶段只确认语义，不重写结构。

#### Scenario: 三件套单值一致
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 获取到的四实体语义、阶段解释与非目标口径一致
- **AND** 不会出现某个文档允许“schema 重构”而另一个文档排除的冲突表达

## REMOVED Requirements

### Requirement: 将四实体语义确认视为结构重构起点
**Reason**: `phase11` 当前目标是冻结上下文系统的可读语义边界，而不是提前打开 canonical 实体重构工程。
**Migration**: 将所有“趁语义确认顺手重写 schema / 关系 / 实体层级”的表达收敛为“先冻结语义，结构重构如有必要必须在未来独立 phase / fix / audit 中承接”。
