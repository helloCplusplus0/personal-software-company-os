# PSCO 上下文系统定位与 Web / Agent 分工冻结 Spec

## Why
`phase11` 已经明确要把 PSCO 推进为“可被 agent 稳定读取的项目上下文系统”，但 `phase11-02` 仍需要一个正式 `/spec` 承接位，把“PSCO 是什么、web 做什么、agent 做什么、哪些事情当前明确不做”冻结成不可漂移的单值输入。这样后续 `phase11-04 ~ 10` 才不会把 PSCO 重新解释为开发流程控制器，或在 web / agent 两条渠道上长出第二套语义。

## What Changes
- 冻结 PSCO 在 `phase11` 的正式定位为“上下文系统”，而不是“开发流程控制器”
- 冻结 web 与 agent 的职责分工、承接边界与共享后端约束
- 冻结 `phase11-02` 的成功标准、DoD 与收口口径
- 冻结“web 不退化、agent 不对称并行进入”的当前阶段口径
- 冻结当前阶段明确不做的渠道能力，阻断第二套语义、第二套流程与第二套事实源

## Impact
- Affected specs: `phase11_project_context_foundation`
- Affected code:
  - `docs/phase/phase11_project_context_foundation_dev_plan.md`
  - `docs/phase/phase11_project_context_foundation_architecture_plan.md`
  - `docs/phase/phase11_project_context_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结 PSCO 为上下文系统
系统规格 SHALL 将 PSCO 在 `phase11` 的正式定位冻结为“上下文系统”，其职责是提供上下文、关系、约束、决策依据与回看入口，而不是在 IDE 现场承担开发流程编排职责。

#### Scenario: 明确系统定位
- **WHEN** 后续执行者读取 `phase11` 三件套并进入 `/spec`
- **THEN** 能明确判断 PSCO 的职责是提供可消费上下文
- **AND** 不会把 PSCO 重新理解为 IDE 现场流程控制器或开发编排器

### Requirement: 冻结 Web / Agent 分工矩阵
系统规格 SHALL 冻结 web 与 agent 的正式分工边界：

- `web` 继续承接全局查看、关系校对、回顾、历史查阅、人工修正与最终确认；
- `agent` 当前只承接项目现场的上下文消费，读取与当前项目直接相关的规则、决策与文档入口，以降低人工重复解释成本。

#### Scenario: 成功判断渠道职责
- **WHEN** 执行者为 `phase11` 设计读取、导出、验收或 handoff
- **THEN** 可以直接判断哪些能力属于 web 渠道，哪些能力属于 agent 渠道
- **AND** 不会让 web 与 agent 同时承接同一职责的第二套正式流程

### Requirement: 冻结共享 Go backend canonical core 约束
系统规格 SHALL 冻结 web 与 agent 共用同一套 Go backend canonical core 的约束，当前阶段不允许任何一方长出第二套领域语义、第二套流程或第二套事实源。

#### Scenario: 阻断第二套主线
- **WHEN** 后续执行者讨论 agent 渠道、导出能力或读取入口
- **THEN** 必须继续复用既有 Go backend canonical core
- **AND** 不得通过前端对话式入口、agent 专属协议对象或独立语义层绕开正式主线

### Requirement: 冻结 agent 当前只做只读消费
系统规格 SHALL 冻结 agent 在 `phase11` 当前阶段只承接现场上下文消费，不承接写回、审批、自动维护或新的正式业务写路径。

#### Scenario: 阻断 agent 写回偷渡
- **WHEN** 执行者讨论 `MCP / CLI / agent 写回 / Draft / 审批流 / 主动注入`
- **THEN** 应直接判断这些内容不属于 `phase11-02` 的正式范围
- **AND** 不得以“为 agent 方便”之名把它们写成当前阶段事实

### Requirement: 冻结 phase11-02 成功标准与收口口径
系统规格 SHALL 冻结 `phase11-02` 的成功标准、DoD 与子任务收口口径，使后续执行者无需再临场判断“定位与分工冻结到什么程度才算完成”。

当前子任务的成功标准至少固定为：

1. `PSCO 是上下文系统、不是开发流程控制器` 的正式定位，已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结；
2. `web` 与 `agent` 的职责分工边界已单值化，且不会让两条渠道并行长出第二套正式流程；
3. `web` 与 `agent` 共享同一套 `Go backend canonical core` 的约束已在三件套中明确，且没有允许任何一方长出第二套语义、第二套流程或第二套事实源；
4. `agent` 当前只做现场上下文消费、`web` 不退化的阶段口径已在三件套中可直接引用，不需要执行者二次解释；
5. 若后续子任务讨论 `agent 写回 / 审批流 / 前端对话式入口 / agent 专属一级业务对象 / 主动注入`，必须能直接判定为超出 `phase11-02` 当前冻结范围，而不是在本子任务内临场吸收。

当前子任务的收口口径至少固定为：

- `phase11-02` 的完成标志不是“新增了一份 spec 包”，而是三件套已经完成系统定位、渠道分工、共享后端约束与当前明确不做事项的正式冻结并保持单值一致；
- 若三件套之间仍存在“PSCO 是否是流程控制器”的表达漂移、`web / agent` 分工冲突、`agent` 写回偷渡空间或共享后端约束缺口，则 `phase11-02` 不得判定为完成。

#### Scenario: 成功判断子任务是否完成
- **WHEN** 后续执行者进入 `phase11-04 ~ 10` 的 `/spec`、实现、验收或 handoff
- **THEN** 可以依据三件套中的固定完成条件直接判断 `phase11-02` 是否达标
- **AND** 不需要重新解释定位、分工、共享后端约束或当前明确不做事项

## MODIFIED Requirements

### Requirement: phase11 三件套中的定位与分工表达
`phase11` 三件套中的定位与分工表达 SHALL 对齐为同一口径：

- PSCO 是上下文系统，不是开发流程控制器；
- web 不退化，继续作为正式交互渠道之一；
- agent 当前只承接现场上下文消费；
- web 与 agent 共享同一套 Go backend canonical core；
- 当前阶段不允许 web 与 agent 各自长出第二套语义、第二套流程或第二套事实源。

#### Scenario: 三件套单值一致
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 获取到的系统定位、职责分工与非目标口径一致
- **AND** 不会出现某个文档允许“流程控制器”或“agent 写回”而另一个文档排除的冲突表达

## REMOVED Requirements

### Requirement: 将 PSCO 视为 IDE 现场流程编排器
**Reason**: `phase11` 当前的中心任务是建立最小只读项目上下文系统，而不是重新定义开发现场的流程控制层。
**Migration**: 将所有“PSCO 指挥 IDE / agent 下一步如何开发”的表达收敛为“PSCO 提供上下文与约束，IDE / agent 负责项目内微观推进”。
