# Phase04-15 审核与根级同步 Spec

## Why

`phase04-01` 到 `phase04-14` 已经完成 `Product Registry + Repository Binding` 的冻结、正式规格、合同主线、前后端实现与联调验收。最后一步必须完成 `phase04` 文档互链复核、根级状态回写与 `phase05` 进入条件确认，避免仓库同时保留“`phase04` 仍在 `/plan`”与“`Product / Repository / Module Binding` 已交付验收通过”两套状态，也避免把 `phase05` 的进入条件继续留在对话层。

## What Changes

- 冻结 `phase04` 文档互链复核的最小检查范围
- 冻结根级状态回写的目标文档、允许更新内容与禁止事项
- 冻结 `plan.md` 中 `phase04` 状态更新、当前阶段切换与完成标志表达
- 冻结 `phase05_dashboard_feedback_foundation` 的进入条件表达
- 冻结 `Product / Repository / Binding` 正式规格、联调验收结果与根级真相源之间的同步关系

## Impact

- Affected specs:
  - `phase04_product_and_repository_binding_foundation`
  - `phase05_dashboard_feedback_foundation`
- Affected code:
  - 当前无业务代码改动，影响 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/*` 与 `phase04-10 ~ 14` 文档同步

## ADDED Requirements

### Requirement: Phase04 文档互链必须完成复核

系统 SHALL 在 `phase04-15` 中完成 `phase04` 阶段文档与根级真相源之间的互链复核，确保不存在冲突口径、过期状态或孤立入口。

#### Scenario: 复核 phase04 文档互链

- **WHEN** 接手者执行 `phase04-15`
- **THEN** 必须检查 `phase04` 三件套、`phase04-10` 正式规格正文、`phase04-11 / 12 / 13 / 14` 的实现与验收文档、以及根级真相源之间的引用关系
- **AND** 必须确认 `phase04-10` 的 `product_repository_binding_spec_v0.1.md` 仍是 `Product / Repository / Binding` 当前阶段的唯一规格收敛入口
- **AND** 必须确认 `phase04-14` 的联调验收结论已经进入当前阶段收口语义，不再保留“尚待联调”或“仍处于 `/plan`”的旧表达

### Requirement: 根级状态必须完成回写

系统 SHALL 在 `phase04-15` 中明确根级状态回写的目标文档、回写边界与同步内容。

#### Scenario: 回写根级状态

- **WHEN** `phase04` 审核通过并进入根级同步
- **THEN** 必须明确哪些根级文档需要同步 `phase04` 的完成状态、交付结果、正式规格入口与下一阶段入口
- **AND** 必须保证回写内容只更新状态、入口、冻结结论与完成标志，不复制整份 phase 正文
- **AND** 不得把联调过程记录、任务清单、异常矩阵或验收细节原样扩写进根级文档

### Requirement: Plan 状态更新必须单值化

系统 SHALL 要求 `plan.md` 中 `phase04` 的状态更新为明确、单值且与当前验收结果一致的表达。

#### Scenario: 更新 plan 中 phase04 状态

- **WHEN** `phase04-15` 完成审核与根级同步
- **THEN** `plan.md` 中的 `phase04` 状态必须从旧的 `current` 或“已进入 `/plan`”口径更新为与交付结果一致的状态
- **AND** 当前阶段完成标志必须与 `phase04-14` 已形成的运行证据保持一致
- **AND** 不得继续保留“当前阶段：`phase04_product_and_repository_binding_foundation`”这类旧表达作为项目当前状态

### Requirement: Phase05 进入条件必须清楚

系统 SHALL 在 `phase04-15` 中明确 `phase05_dashboard_feedback_foundation` 的进入条件，使下一阶段直接承接 `phase04` 已交付结果，而不是重新解释 `Product / Repository / Binding` 当前边界。

#### Scenario: 确认 phase05 进入条件

- **WHEN** 接手者查询 `phase05` 是否可以开始
- **THEN** 必须能明确得到 `phase05` 以前一阶段已交付的 `Product Registry + Repository Binding` 主线、`phase04-10` 正式规格正文、`phase04-11` 合同主线与 `phase04-14` 联调验收结论为直接上游输入
- **AND** 必须明确 `phase05` 不重复实现 `Product Registry`、`Repository Binding` 与三类绑定动作主线，而是推进 `Dashboard + Feedback` 的最小闭环
- **AND** 不得在进入 `phase05` 时撤销 `phase04` 已冻结的页面、动作、数据、合同与联调结论

### Requirement: 根级真相源与 phase04 文档必须保持单值一致

系统 SHALL 要求 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 与 `phase04` 文档在阶段状态、入口路径与当前冻结结论上保持单值一致。

#### Scenario: 校验根级文档与 phase04 文档一致性

- **WHEN** 接手者对根级文档和 `phase04` 文档做一致性检查
- **THEN** 必须确认项目当前阶段状态、`phase04` 是否完成、`phase05` 的下一阶段入口、目录入口与规格入口一致
- **AND** 不得出现根级文档已经进入 `phase05`，但 `phase04` 文档仍保留旧状态，或相反情况

## MODIFIED Requirements

### Requirement: Phase04 收口后的合法入口

`phase04` 收口完成后，后续接手与 `phase05` SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `docs/README.md` / `architecture_map.md` 为准，`Product / Repository / Binding` 执行层规格以 `phase04-10` 的正式规格正文与 `phase04-14` 验收结论为准。

#### Scenario: 后续阶段读取入口

- **WHEN** 后续接手者开始 `phase05`
- **THEN** 必须先从根级文档读取项目当前状态、阶段入口与目录入口
- **AND** 再从 `phase04-10` 正式规格正文与 `phase04-14` 验收结果读取 `Product / Repository / Binding` 已交付边界
- **AND** 不得把任何单个子任务文档、实现注释或对话结论当作新的长期主入口

## REMOVED Requirements

### Requirement: 保留 phase04 仍处于当前进行中或未收口状态的旧表达

**Reason**: `phase04` 已完成从 `/plan -> /spec -> 实现 -> 验收` 的交付链路，继续保留旧状态会导致根级入口与实际交付结果冲突。
**Migration**: 改为在根级文档中明确 `phase04` 已收口完成，并把 `phase05_dashboard_feedback_foundation` 标成当前下一阶段入口，同时保留 `phase04-10` 与 `phase04-14` 作为已交付边界入口。
