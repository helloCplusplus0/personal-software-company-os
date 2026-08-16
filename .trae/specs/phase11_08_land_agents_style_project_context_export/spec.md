# AGENTS 风格项目上下文导出 Spec

## Why
`phase11-07` 已经落地了最小结构化只读项目上下文能力，但新接手 agent 仍缺少一份可直接消费的文档化导出结果。`phase11-08` 需要把这份导出能力正式落地，同时继续守住“单向派生、只读、无写回、无第二事实源”的边界。

## What Changes
- 冻结 AGENTS 风格或等价 Markdown 风格项目上下文导出的正式承接边界
- 冻结结构化只读结果到 Markdown 导出的单向派生关系
- 冻结导出内容的最小字段边界、受控引用与可读组织方式
- 冻结 PSCO 仓库自身 dogfooding 的最小验收路径
- 冻结不主动写入外部仓库、不扩写为主动注入能力的硬约束

## Impact
- Affected specs:
  - `phase11_project_context_foundation`
  - `phase11_05_design_minimal_readonly_project_context_export`
  - `phase11_07_land_minimal_readonly_project_context_read_capability`
- Affected code:
  - `docs/phase/phase11_project_context_foundation_dev_plan.md`
  - `docs/phase/phase11_project_context_foundation_architecture_plan.md`
  - `docs/phase/phase11_project_context_foundation_shared_baseline.md`
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/`
  - `backend/internal/platform/`

## ADDED Requirements

### Requirement: 提供 AGENTS 风格项目上下文导出的正式承接位
系统 SHALL 提供一个最小 AGENTS 风格或等价 Markdown 风格的项目上下文导出能力，该能力继续落在 Go backend 的正式只读业务接口层，并服务于新接手 agent 的直接消费。

补充冻结：

- 导出能力的身份是 `phase11-07` 结构化只读结果的文档化投影；
- 导出能力不得成为新的业务事实源、目录扫描器或仓库写入器；
- 导出能力不得要求消费侧项目目录与 `PSCO` 当前仓库具有相同文件布局。

#### Scenario: 新接手 agent 读取导出
- **WHEN** 新接手 agent 需要快速恢复当前项目核心上下文
- **THEN** 系统可以提供一份 AGENTS 风格或等价 Markdown 风格的正式导出结果
- **AND** 该结果可以直接消费，而不要求 agent 再次拼装结构化字段

### Requirement: Markdown 导出必须从结构化只读结果单向派生
系统 SHALL 要求 Markdown 导出只能从 `phase11-07` 的结构化只读结果单向派生，不得绕过结构化读取主线直接从根级文档、目录扫描结果或临时补猜生成并列导出。

补充冻结：

- 导出生成过程中允许做文案组织、标题编排与受控引用整形；
- 不允许额外长出第二套字段语义、第二套状态机或第二套事实判断；
- 不允许 Markdown 导出反向决定结构化读取字段边界。

#### Scenario: 生成 Markdown 导出
- **WHEN** 执行者实现 AGENTS 风格导出
- **THEN** 必须先消费同一 `repository_id` 对应的结构化只读结果
- **AND** 再将其单向渲染为 Markdown

### Requirement: 冻结 Markdown 导出的最小内容边界
系统 SHALL 冻结 Markdown 导出的最小内容边界，保证导出结果既能直接服务 agent 消费，又不会扩张成第二套冗余知识库。

Markdown 导出至少承接：

- 当前 `Repository` 摘要；
- 当前 phase 与单一主交付摘要；
- 当前项目关联的 `Product / Module / Decision` 最小可读摘要；
- 当前明确不做或不承接的边界摘要；
- 规则、约束与文档入口的受控引用。

补充冻结：

- 文档入口只允许承接 `phase11-07` 结构化结果中已存在的入口定位值与定位类型；
- 导出文案可以重组顺序，但不得擅自补充结构化结果中不存在的新事实。

#### Scenario: 检查导出字段是否达标
- **WHEN** 执行者对照 Markdown 导出内容进行验收
- **THEN** 可以直接判断最小内容边界是否完整
- **AND** 不需要再临场决定“哪些内容该出现、哪些内容只能留在结构化层”

### Requirement: 冻结 PSCO 仓库自身 dogfooding 验收路径
系统 SHALL 将 `PSCO` 仓库自身冻结为 `phase11-08` 的第一 dogfooding 场景，并用固定入口数量与固定恢复问题验证导出是否真实降低上下文恢复成本。

补充冻结：

- dogfooding 继续使用当前 `PSCO` 仓库，只作为第一验证对象，不上升为未来所有项目的固定目录模板；
- 新路径固定入口集合继续遵守 `shared_baseline` 已冻结的 `<= 3` 入口约束；
- 导出结果必须能帮助 agent 回答当前 phase、直接上游、单一主交付、明确不做、关联实体摘要入口这 `5` 个恢复问题。

#### Scenario: 使用 PSCO 自身做 dogfooding
- **WHEN** 验收者使用 `PSCO` 仓库执行新路径 dogfooding 剧本
- **THEN** 可以在不超过 `3` 个固定入口的前提下恢复核心上下文
- **AND** AGENTS 风格导出是其中的正式入口之一

### Requirement: 冻结不写入与不主动注入边界
系统 SHALL 明确禁止 AGENTS 风格导出能力扩写为主动注入、仓库写入、外部仓库同步或 agent 写回流程。

补充冻结：

- 当前阶段只承接读取与文档化导出；
- 不主动写入 `AGENTS.md`、不主动写入外部仓库、也不生成需落盘的默认文件；
- 不将导出能力扩写为审批流、草稿流或会话级记忆回写。

#### Scenario: 阻断写回型扩张
- **WHEN** 后续执行者讨论“顺手把导出结果写进仓库”或“自动注入给 agent”
- **THEN** 应直接判定这不属于 `phase11-08` 当前范围
- **AND** 不得作为当前任务的一部分落地

### Requirement: 冻结最小实现与验收结果
系统 SHALL 将 `phase11-08` 的完成标准冻结为“最小文档导出能力已正式存在，并已通过 PSCO 自身 dogfooding 验收”，而不是只停留在设计口径。

当前阶段至少需要形成：

1. 一个正式的只读 Markdown 导出承接位；
2. 与 `phase11-07` 结构化读取之间的单向派生实现；
3. 覆盖“仓库存在且绑定完成”成功路径的导出验证；
4. 覆盖不主动写入外部仓库、不形成第二事实源的验收说明。

#### Scenario: 判断 phase11-08 是否完成
- **WHEN** 执行者准备宣告 `phase11-08` 完成
- **THEN** 必须能指出正式导出承接位、单向派生实现与 dogfooding 验收证据
- **AND** 不得仅以“已有一段 Markdown 模板”判定完成

## MODIFIED Requirements

### Requirement: phase11 项目上下文导出链路
`phase11` 的项目上下文导出链路 SHALL 明确分成两层并保持单向关系：

- 第一层：`phase11-07` 的结构化只读项目上下文读取；
- 第二层：`phase11-08` 的 AGENTS 风格或等价 Markdown 风格导出。

补充冻结：

- 第二层只消费第一层结果；
- 两层都属于只读能力；
- 第二层不得绕开第一层形成并列读取主线。

#### Scenario: 查看 phase11 导出链路
- **WHEN** 读者同时查看 `phase11-07` 与 `phase11-08`
- **THEN** 能明确两者是“结构化读取在前、Markdown 导出在后”的关系
- **AND** 不会把 Markdown 导出误读为独立事实源

## REMOVED Requirements

### Requirement: 将 AGENTS 风格导出等同于仓库写入能力
**Reason**: `phase11-08` 当前目标是提供最小只读文档导出，而不是主动改写仓库或为任意项目注入固定文件。
**Migration**: 所有“自动写入 `AGENTS.md` / 自动同步到外部仓库 / 自动注入到会话上下文”的想法统一下沉为未来候选增强，不在 `phase11-08` 正式范围承接。
