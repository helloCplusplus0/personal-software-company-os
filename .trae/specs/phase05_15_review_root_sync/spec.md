# Phase05-15 审核与根级同步 Spec

## Why

`phase05-01` 到 `phase05-14` 已完成 `Dashboard + Feedback` 的冻结、正式规格、`.proto` 合同主线、前后端实现与联调验收。最后一步必须完成 `phase05` 文档互链复核、根级状态回写与下一阶段入口确认，避免仓库同时保留“`phase05` 仍在 `/plan`”与“`Dashboard + Feedback` 已交付验收通过”两套状态，也避免把下一阶段入口继续留在对话层或临时记忆里。

`phase05-15` 的目标是把 `phase05` 的收口事实同步到根级真相源，并明确下一阶段入口的表达边界：该入口必须以已存在或本次明确建立的正式 phase 入口为准，不得在根级文档中凭空猜测一个未命名、未建档的下一阶段。

## What Changes

- 冻结 `phase05` 文档互链复核的最小检查范围
- 冻结根级状态回写的目标文档、允许更新内容与禁止事项
- 冻结 `plan.md` 中 `phase05` 状态更新、当前阶段切换与完成标志表达
- 冻结下一阶段入口的确认规则、上游输入表达与禁止猜测边界
- 冻结 `Dashboard + Feedback` 正式规格、合同主线、联调验收结论与根级真相源之间的同步关系

## Impact

- Affected specs:
  - `phase05_dashboard_feedback_foundation`
  - 下一阶段正式 phase 入口（若已存在，则作为 `phase05-15` 的直接下游）
- Affected code:
  - 当前无业务代码改动，影响 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/*` 与 `phase05-10 ~ 14` 文档同步

## ADDED Requirements

### Requirement: Phase05 文档互链必须完成复核

系统 SHALL 在 `phase05-15` 中完成 `phase05` 阶段文档与根级真相源之间的互链复核，确保不存在冲突口径、过期状态或孤立入口。

#### Scenario: 复核 phase05 文档互链

- **WHEN** 接手者执行 `phase05-15`
- **THEN** 必须检查 `phase05` 三件套、`phase05-10` 正式规格正文、`phase05-11 / 12 / 13 / 14` 的合同、实现与验收文档、以及根级真相源之间的引用关系
- **AND** 必须确认 `phase05-10` 的 `dashboard_feedback_spec_v0.1.md` 仍是 `Dashboard + Feedback` 当前阶段的唯一规格收敛入口
- **AND** 必须确认 `phase05-11` 的 `.proto` 合同主线、`phase05-12 / 13` 的实现结论与 `phase05-14` 的联调验收结论已经进入当前阶段收口语义，不再保留“仍在 `/plan`”或“仍待联调”的旧表达

### Requirement: 根级状态必须完成回写

系统 SHALL 在 `phase05-15` 中明确根级状态回写的目标文档、回写边界与同步内容。

#### Scenario: 回写根级状态

- **WHEN** `phase05` 审核通过并进入根级同步
- **THEN** 必须明确哪些根级文档需要同步 `phase05` 的完成状态、交付结果、正式规格入口、合同入口、验收入口与下一阶段入口
- **AND** 必须保证回写内容只更新状态、入口、冻结结论与完成标志，不复制整份 phase 正文
- **AND** 不得把联调过程记录、任务清单、异常矩阵或验收细节原样扩写进根级文档

### Requirement: Plan 状态更新必须单值化

系统 SHALL 要求 `plan.md` 中 `phase05` 的状态更新为明确、单值且与当前验收结果一致的表达。

#### Scenario: 更新 plan 中 phase05 状态

- **WHEN** `phase05-15` 完成审核与根级同步
- **THEN** `plan.md` 中的 `phase05` 状态必须从旧的“已进入 `/plan`”或 `current` 口径更新为与交付结果一致的状态
- **AND** 当前阶段完成标志必须与 `phase05-14` 已形成的运行证据保持一致
- **AND** 不得继续保留“当前目标是完成 `phase05` 三件套规划”这类已过期表达作为项目当前状态

### Requirement: 下一阶段入口必须清楚且受控

系统 SHALL 在 `phase05-15` 中明确下一阶段入口的表达边界，使后续阶段直接承接 `phase05` 已交付结果，而不是重新解释 `Dashboard + Feedback` 当前边界或在根级文档中自由猜测下一阶段名称。

#### Scenario: 确认下一阶段入口

- **WHEN** 接手者查询下一阶段是否可以开始
- **THEN** 必须能明确得到下一阶段的正式入口表达
- **AND** 该入口必须直接承接 `phase05` 已交付的 `Dashboard + Feedback` 主线、`phase05-10` 正式规格正文、`phase05-11` 合同主线与 `phase05-14` 联调验收结论为直接上游输入
- **AND** 若下一阶段的正式 phase 名称、三件套或 dev plan 尚未建立，根级文档不得凭空杜撰一个新阶段名充当正式入口
- **AND** 在这种情况下，必须把“下一阶段入口待以正式 phase 文档建立后切换”写成显式状态，而不是隐含假设

### Requirement: 根级真相源与 phase05 文档必须保持单值一致

系统 SHALL 要求 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 与 `phase05` 文档在阶段状态、入口路径与当前冻结结论上保持单值一致。

#### Scenario: 校验根级文档与 phase05 文档一致性

- **WHEN** 接手者对根级文档和 `phase05` 文档做一致性检查
- **THEN** 必须确认项目当前阶段状态、`phase05` 是否完成、`Dashboard + Feedback` 当前阶段唯一规格入口、合同入口、验收入口与下一阶段入口一致
- **AND** 不得出现根级文档已经宣告 `phase05` 完成，但 `phase05` 文档仍保留旧状态，或相反情况

## MODIFIED Requirements

### Requirement: Phase05 收口后的合法入口

`phase05` 收口完成后，后续接手与下一阶段 SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `docs/README.md` / `architecture_map.md` 为准，`Dashboard + Feedback` 执行层规格以 `phase05-10` 的正式规格正文、`phase05-11` 合同主线与 `phase05-14` 验收结论为准。

#### Scenario: 后续阶段读取入口

- **WHEN** 后续接手者进入下一阶段
- **THEN** 必须先从根级文档读取项目当前状态、阶段入口与目录入口
- **AND** 再从 `phase05-10` 正式规格正文、`phase05-11` 合同主线与 `phase05-14` 验收结果读取 `Dashboard + Feedback` 已交付边界
- **AND** 不得把任何单个子任务文档、实现注释、联调临时命令或对话结论当作新的长期主入口

## REMOVED Requirements

### Requirement: 保留 phase05 仍处于当前进行中、仅完成三件套规划或尚未收口状态的旧表达

**Reason**: `phase05` 已完成从 `/plan -> /spec -> 实现 -> 验收` 的交付链路，继续保留旧状态会导致根级入口与实际交付结果冲突。
**Migration**: 改为在根级文档中明确 `phase05` 已收口完成，并以 `phase05-10` 正式规格正文、`phase05-11` 合同主线与 `phase05-14` 验收结论作为已交付边界；下一阶段入口仅在正式建立后切换为新的 phase 入口。
