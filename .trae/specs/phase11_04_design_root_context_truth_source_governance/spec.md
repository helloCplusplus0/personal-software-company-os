# 根级上下文真相源治理设计 Spec

## Why
`phase11` 已经明确“根级上下文真相源治理”是主交付主体之一，但如果不把根级入口的单一写者、重复承载清理策略与最终共识入口改写策略正式冻结，后续根级同步实现仍会回到“哪些文档能写正文、哪些只能摘要引用、到底是一次性校准还是全量派生”的反复争论。`phase11-04` 需要把这组治理设计冻结成可直接指导根级同步的单值输入。

## What Changes
- 冻结仅针对 `PSCO` 自身仓库的根级上下文真相源治理范围
- 冻结根级入口治理矩阵、重复承载清单、悬空引用清理清单与单一写者规则表的正式设计要求
- 冻结 `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识入口的统一改写策略
- 冻结不再允许出现的重复表达模式与非目标边界
- 冻结 `phase11-04` 的成功标准、DoD 与收口口径

## Impact
- Affected specs: `phase11_project_context_foundation`
- Affected code:
  - `README.md`
  - `plan.md`
  - `AGENTS.md`
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

### Requirement: 冻结根级治理对象与范围边界
系统规格 SHALL 将 `phase11-04` 的治理对象冻结为仅针对 `PSCO` 自身仓库的根级/入口文档集合：

- `README.md`
- `plan.md`
- `AGENTS.md`
- `architecture_map.md`
- `docs/README.md`
- `docs/phase/README.md`
- `project_rules.md`
- `global_skills.md`

补充冻结：

- `PSCO-mvp05-summarize-feedback.md` 是本子任务必须统一指向的根级最终共识入口；
- 当前子任务只承接 `PSCO` 自身仓库治理，不外推为未来所有项目的固定目录模板或默认入口清单。

#### Scenario: 成功限定治理范围
- **WHEN** 后续执行者为 `phase11` 设计根级同步实现
- **THEN** 能明确哪些文档属于本次治理对象
- **AND** 不会把当前仓库清单误写成未来所有项目的隐含合同

### Requirement: 冻结根级入口治理矩阵与单一写者规则
系统规格 SHALL 冻结根级入口治理矩阵与单一写者规则，使每一类主结论都能直接定位唯一正式承接位，其他入口只保留摘要式引用或受控跳转。

治理矩阵至少覆盖：

- 哪个文件是某类主结论的唯一正式承接位；
- 哪些文件只允许保留摘要式引用；
- 哪些文件不得重复承载 phase 状态、目录落点、技术栈正文或最终共识正文；
- 收口后的单一写者规则表。

#### Scenario: 成功判断单一写者
- **WHEN** 后续执行者盘点根级入口中的重复承载内容
- **THEN** 能直接判断每类主结论应收敛到哪个正式承接位
- **AND** 不需要重新争论“这个结论该写在哪个入口文档里”

### Requirement: 冻结重复承载与悬空引用清理设计
系统规格 SHALL 冻结重复承载清单、目标落点清单与悬空引用清理清单的设计要求，用于直接指导根级同步实现。

补充冻结：

- 必须显式列出当前不再允许保留的重复表达模式；
- 必须显式列出需要改写为摘要式引用、受控跳转或删除的入口内容；
- 必须显式列出指向不存在文件或错误共识入口的悬空引用清理对象。

#### Scenario: 成功指导清理动作
- **WHEN** 后续执行者进入根级入口治理实现
- **THEN** 可以直接根据清单执行重复承载清理与悬空引用清理
- **AND** 不需要临场补猜“哪些重复可以保留、哪些引用必须清掉”

### Requirement: 冻结最终共识入口统一改写策略
系统规格 SHALL 冻结 `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识入口的统一改写策略，要求相关根级入口统一指向该文档而不是继续保留并列结论正文或失效入口引用。

#### Scenario: 成功统一最终共识入口
- **WHEN** 后续执行者治理 `README.md`、`AGENTS.md`、`docs/README.md` 或其他根级入口
- **THEN** 能明确将最终共识类表达统一指向 `PSCO-mvp05-summarize-feedback.md`
- **AND** 不会继续保留并列的第二个最终共识正文入口

### Requirement: 冻结一次性校准而非全量派生的阶段口径
系统规格 SHALL 冻结 `phase11-04` 当前承接的是根级入口治理设计与一次性校准策略，而不是静态文件全量 backend 派生方案。

#### Scenario: 阻断治理路线反复争论
- **WHEN** 后续执行者讨论根级入口治理应采用何种实现路线
- **THEN** 应直接判断当前阶段不承接“静态文件全量派生”路线
- **AND** 不需要继续开放争论“是全量派生还是一次性校准”

### Requirement: 冻结 phase11-04 成功标准与收口口径
系统规格 SHALL 冻结 `phase11-04` 的成功标准、DoD 与子任务收口口径，使后续执行者无需再临场判断“治理设计做到什么程度才算完成”。

当前子任务的成功标准至少固定为：

1. 根级治理对象、单一写者规则、重复承载清理策略与悬空引用清理策略已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结；
2. `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识入口的统一改写策略已在三件套中可直接引用；
3. “当前只做治理设计与一次性校准策略，不做静态文件全量派生”的阶段口径已单值化；
4. 后续执行者可以直接据此产出根级入口治理矩阵、重复承载清单、悬空引用清理清单与单一写者规则表；
5. 当前仓库治理清单与未来跨项目模板之间的边界已写清，不会把 `PSCO` 当前根级文件结构上升为所有项目的固定合同。

当前子任务的收口口径至少固定为：

- `phase11-04` 的完成标志不是“新增了一份 spec 包”，而是三件套已经完成根级治理策略、统一改写策略与非目标边界的正式冻结并保持单值一致；
- 若三件套之间仍存在单一写者冲突、重复承载清理标准不一致、最终共识入口并列、或仍给“全量派生”路线留出当前阶段偷渡空间，则 `phase11-04` 不得判定为完成。

#### Scenario: 成功判断子任务是否完成
- **WHEN** 后续执行者进入根级同步实现、验收或 handoff
- **THEN** 可以依据三件套中的固定完成条件直接判断 `phase11-04` 是否达标
- **AND** 不需要重新解释治理对象、治理策略或治理路线

## MODIFIED Requirements

### Requirement: phase11 三件套中的根级治理表达
`phase11` 三件套中的根级治理表达 SHALL 对齐为同一口径：

- 当前阶段仅针对 `PSCO` 自身仓库治理根级/入口文档；
- 每类主结论必须有唯一正式承接位；
- 其他入口只保留摘要式引用或受控跳转；
- `PSCO-mvp05-summarize-feedback.md` 是根级最终共识单值入口；
- 当前阶段不承接静态文件全量 backend 派生。

#### Scenario: 三件套单值一致
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 获取到的治理对象、单一写者规则、清理策略与非目标口径一致
- **AND** 不会出现某个文档允许“全量派生”或“并列最终共识入口”而另一个文档排除的冲突表达

## REMOVED Requirements

### Requirement: 将根级入口治理视为静态文件全量派生前置阶段
**Reason**: `phase11` 当前目标是先完成根级真相源治理设计与一次性校准策略冻结，而不是提前建设静态文件派生系统。
**Migration**: 将所有“顺手切到 backend 全量派生”的表达收敛为未来候选增强路线，不在 `phase11-04` 的正式范围内承接。
