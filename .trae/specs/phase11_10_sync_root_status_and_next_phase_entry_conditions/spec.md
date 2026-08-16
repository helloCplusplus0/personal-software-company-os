# Phase11-10 根级同步、阶段状态回写与下一阶段进入条件留档 Spec

## Why

`phase11-09` 已完成正式验收，但当前根级入口与 docs 入口仍停留在“`phase11` 处于 `/plan` 待进入 `/spec`”的旧状态。需要用最小范围回写正式状态、验收入口与下一阶段进入条件，避免根级真相源再次漂移。

## What Changes

- 回写 `phase11` 在根级与 docs 入口中的正式状态，使其与已完成验收的事实一致
- 冻结 `phase11` 的正式验收入口与收口入口，保证后续接手者有单值回看路径
- 单值化“下一阶段进入条件”，只把更重的消费通道或受控维护能力保留为后续进入条件，不写成当前既成事实
- 保持 `plan.md` 继续作为阶段状态唯一正式承接位，其他入口只做摘要式同步

## Impact

- Affected specs: `phase11_project_context_foundation`、`phase11-09_validate_project_context_foundation_dogfooding_regression`
- Affected code: `AGENTS.md`、`plan.md`、`docs/README.md`、`docs/phase/README.md`

## ADDED Requirements

### Requirement: Phase11 正式收口状态回写

系统 SHALL 在 `phase11-10` 实施后，把 `phase11` 的正式状态回写到根级真相源，并明确本阶段已完成正式验收与收口。

#### Scenario: `plan.md` 作为唯一 canonical 状态入口

- **WHEN** 执行 `phase11-10` 的根级同步
- **THEN** `plan.md` 必须明确 `phase11` 的正式状态、当前收口结果与下一阶段进入条件
- **AND** 其他根级或 docs 入口只允许摘要式引用该状态，不得形成第二个 canonical writer

### Requirement: Phase11 正式验收入口留档

系统 SHALL 为 `phase11` 提供单值、可回看的正式验收入口，使接手者可以直接定位本阶段的正式验收与收口证据。

#### Scenario: 根级与 docs 入口可定位 phase11 验收结果

- **WHEN** 接手者从 `AGENTS.md`、`docs/README.md` 或 `docs/phase/README.md` 恢复当前项目状态
- **THEN** 这些入口必须能指向 `phase11` 的正式验收/收口入口
- **AND** 不得在多个根级文档重复承载整段验收正文

### Requirement: 下一阶段进入条件单值化

系统 SHALL 单值表达 `phase11` 收口后的下一阶段进入条件，并把更重的消费通道或受控维护能力保留为进入条件，而不是当前事实。

#### Scenario: 后续能力只作为进入条件保留

- **WHEN** 根级与 docs 入口描述 `phase11` 之后的推进方向
- **THEN** 只允许出现一条与 `phase11` 收口结论一致的进入条件表述
- **AND** 不得把 `MCP / CLI / agent 写回 / 更重的消费通道` 写成已经进入的当前主线

## MODIFIED Requirements

### Requirement: 根级入口摘要职责

`AGENTS.md`、`docs/README.md` 与 `docs/phase/README.md` 必须继续只承接摘要式入口职责。它们可以同步 `phase11` 的正式收口状态、验收入口与下一阶段进入条件，但不得替代 `plan.md` 成为阶段状态的正式承接位。

### Requirement: 当前阶段状态摘要

当前根级状态摘要必须从“`phase11` 已完成 `/plan`、待复核进入 `/spec`”更新为与正式验收结果一致的收口状态，并同步保留最近完成正式业务 phase 与最近完成正式支撑能力 phase 的摘要定位。

## REMOVED Requirements

### Requirement: `phase11` 待进入 `/spec` 的旧状态表述

**Reason**: 该表述已与 `phase11-09` 的正式验收结果不一致，会造成接手者误判当前阶段状态。

**Migration**: 将受影响入口统一更新为“`phase11` 已完成正式验收与收口”，并把后续更重能力改写为进入条件而非当前事实。
