# Phase02-13 审核与根级同步 Spec

## Why

`phase02-01` 到 `phase02-12` 已经完成 `Module Registry` 的冻结、正式规格、前后端实现、数据主线、Proto 合同与联调验收。最后一步必须完成 `phase02` 文档互链复核、根级状态回写与下一阶段入口确认，确保项目不会同时保留“phase02 仍在 /plan”与“phase02 已交付验收通过”两套状态，也不会把 `phase03` 的进入条件遗留在对话层。

## What Changes

- 冻结 `phase02` 文档互链复核的最小检查范围
- 冻结根级状态回写的目标文档、允许更新内容与禁止事项
- 冻结 `plan.md` 中 `phase02` 状态更新与完成标志表达
- 冻结 `phase03_decision_center_foundation` 的进入条件表达
- 冻结 `Module Registry` 正式规格、联调验收结果与根级真相源之间的同步关系

## Impact

- Affected specs:
  - `phase02_module_registry_foundation`
  - `phase03_decision_center_foundation`
- Affected code:
  - 当前无业务代码改动，影响 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/*` 与 `phase02-09 ~ 12` 文档同步

## ADDED Requirements

### Requirement: Phase02 文档互链必须完成复核

系统 SHALL 在 `phase02-13` 中完成 `phase02` 阶段文档与根级真相源之间的互链复核，确保不存在冲突口径、过期状态或孤立入口。

#### Scenario: 复核 phase02 文档互链

- **WHEN** 接手者执行 `phase02-13`
- **THEN** 必须检查 `phase02` 三件套、`phase02-09` 正式规格正文、`phase02-10 / 11 / 11A / 12` 的实现与验收文档、以及根级真相源之间的引用关系
- **AND** 必须确认 `phase02-09` 的 `module_registry_spec_v0.1.md` 仍是 `Module Registry` 当前阶段的唯一规格收敛入口
- **AND** 不得保留“`Protocol Buffers` 仅为长期方向、不在当前阶段落地”的旧表达

### Requirement: 根级状态必须完成回写

系统 SHALL 在 `phase02-13` 中明确根级状态回写的目标文档、回写边界与同步内容。

#### Scenario: 回写根级状态

- **WHEN** `phase02` 审核通过并进入根级同步
- **THEN** 必须明确哪些根级文档需要同步 `phase02` 的完成状态、交付结果、正式规格入口与下一阶段入口
- **AND** 必须保证回写内容只更新状态、入口、冻结结论与完成标志，不复制整份 phase 正文
- **AND** 不得把验收报告原样扩写进根级文档

### Requirement: Plan 状态更新必须单值化

系统 SHALL 要求 `plan.md` 中 `phase02` 的状态更新为明确、单值且与当前验收结果一致的表达。

#### Scenario: 更新 plan 中 phase02 状态

- **WHEN** `phase02-13` 完成审核与根级同步
- **THEN** `plan.md` 中的 `phase02` 状态必须从旧的 `/plan` 口径更新为与交付结果一致的状态
- **AND** 当前阶段完成标志必须与 `phase02-12` 已形成的运行证据保持一致
- **AND** 不得继续保留“当前阶段：`phase02_module_registry_foundation (/plan)`”这类旧表达

### Requirement: Phase03 进入条件必须清楚

系统 SHALL 在 `phase02-13` 中明确 `phase03_decision_center_foundation` 的进入条件，使下一阶段直接承接 `phase02` 已交付结果，而不是重新解释边界。

#### Scenario: 确认 phase03 进入条件

- **WHEN** 接手者查询 `phase03` 是否可以开始
- **THEN** 必须能明确得到 `phase03` 以前一阶段已交付的 `Module Registry` 主线、`Decision` 仍属 MVP 必入对象、以及 `phase02-09` / `phase02-11A` 的规格与合同结果为直接上游输入
- **AND** 必须明确 `phase03` 不重复实现 `Module Registry` 主线，而是承接 `Decision Center` 的最小闭环
- **AND** 不得在进入 `phase03` 时撤销 `phase02` 已冻结的前后端、合同与联调结论

### Requirement: 根级真相源与 phase02 文档必须保持单值一致

系统 SHALL 要求 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 与 `phase02` 文档在阶段状态、入口路径与当前冻结结论上保持单值一致。

#### Scenario: 校验根级文档与 phase02 文档一致性

- **WHEN** 接手者对根级文档和 `phase02` 文档做一致性检查
- **THEN** 必须确认项目当前阶段状态、`phase02` 是否完成、`phase03` 的下一阶段入口、目录入口与规格入口一致
- **AND** 不得出现根级文档已经进入下一阶段，但 `phase02` 文档仍保留旧状态，或相反情况

## MODIFIED Requirements

### Requirement: Phase02 收口后的合法入口

`phase02` 收口完成后，后续接手与 `phase03` SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `docs/README.md` / `architecture_map.md` 为准，`Module Registry` 执行层规格以 `phase02-09` 的正式规格正文与 `phase02-12` 验收结论为准。

#### Scenario: 后续阶段读取入口

- **WHEN** 后续接手者开始 `phase03`
- **THEN** 必须先从根级文档读取项目当前状态、阶段入口与目录入口
- **AND** 再从 `phase02-09` 正式规格正文与 `phase02-12` 验收结果读取 `Module Registry` 已交付边界
- **AND** 不得把任何单个子任务文档或对话结论当作新的长期主入口

## REMOVED Requirements

### Requirement: 保留 phase02 仍处于规划或未收口状态的旧表达

**Reason**: `phase02` 已完成从 `/plan -> /spec -> 实现 -> 验收` 的交付链路，继续保留旧状态会导致根级入口与实际交付结果冲突。
**Migration**: 改为在根级文档中明确 `phase02` 已收口完成，并把 `phase03_decision_center_foundation` 标成当前下一阶段入口。
