# Phase01-07 审核与根级同步 Spec

## Why

`phase01-01` 到 `phase01-06` 已经分别完成技术边界、对象动作、页面输入、数据合同、冷启动导入导出与正式 MVP 规格正文的冻结。最后一步必须完成 `phase01` 文档互链复核、根级状态回写与下一阶段入口确认，确保项目不会同时保留多套入口状态或把新增补充共识遗落在对话中。

## What Changes

- 冻结 `phase01` 文档互链复核的最小检查范围
- 冻结根级状态回写的目标文档与允许更新内容
- 冻结 `phase02_module_registry_foundation` 的进入条件表达要求
- 冻结正式 MVP 规格正文与根级真相源之间的同步关系
- 将“前端正式交付物为单一 React Web，同时考虑 PC 与移动浏览器 UI，`React Native` 不进入 `v0.1` 首轮，`PWA` 仅为可兼容增强方向”的补充共识纳入本次根级同步范围

## Impact

- Affected specs: `phase01_mvp_spec_convergence`
- Affected code: 当前无代码改动，影响 `AGENTS.md`、`plan.md`、`architecture_map.md`、`project_rules.md`、`docs/phase/*` 与 `phase02` 进入前提的文档同步

## ADDED Requirements

### Requirement: Phase01 文档互链必须完成复核
系统 SHALL 在 `phase01-07` 中完成 `phase01` 阶段文档与根级真相源之间的互链复核，确保不存在冲突口径、过期状态或孤立规格入口。

#### Scenario: 复核 phase01 文档互链
- **WHEN** 接手者执行 `phase01-07`
- **THEN** 必须检查 `phase01` 三件套、`phase01-01` 到 `phase01-06` 子规格与根级真相源之间的引用关系
- **AND** 必须确认 `phase01-06` 的正式 MVP 规格正文已成为执行层唯一规格入口
- **AND** 不得保留与根级文档冲突的旧状态或并列入口

### Requirement: 根级状态必须完成回写
系统 SHALL 在 `phase01-07` 中明确根级状态回写的目标文档、回写边界与同步内容。

#### Scenario: 回写根级状态
- **WHEN** `phase01` 审核通过并进入根级同步
- **THEN** 必须明确哪些根级文档需要同步 `phase01` 完成状态、正式规格入口与下一阶段入口
- **AND** 必须保证回写内容只更新状态、入口与冻结结论，不重写根级文档职责
- **AND** 不得在根级文档中复制整份正式 MVP 规格正文

### Requirement: Plan 状态更新必须单值化
系统 SHALL 要求 `plan.md` 中 `phase01` 的状态更新为明确、单值且与 `phase01` 验收结果一致的表达。

#### Scenario: 更新 plan 中 phase01 状态
- **WHEN** `phase01-07` 完成审核与根级同步
- **THEN** `plan.md` 中的 `phase01` 状态必须更新正确
- **AND** 当前阶段完成标志必须与已形成的正式规格入口保持一致
- **AND** 不得继续保留“`phase01` 仍处于未审核状态”的旧表达

### Requirement: Phase02 进入条件必须清楚
系统 SHALL 在 `phase01-07` 中明确 `phase02_module_registry_foundation` 的进入条件，使其直接承接正式 MVP 规格正文，而不是重新解释边界。

#### Scenario: 确认 phase02 进入条件
- **WHEN** 接手者查询 `phase02` 是否可以开始
- **THEN** 必须能明确得到 `phase02` 以 `phase01-06` 正式 MVP 规格正文为唯一上游规格来源
- **AND** 必须明确 `Module Registry` 的实现范围只承接 `v0.1` 已冻结边界
- **AND** 不得在进入 `phase02` 时重新引入 `Feature / Opportunity / Experiment`、独立 `AI Assistant`、`React Native` 客户端或完整 `PWA` 能力作为前置范围

### Requirement: 前端端策略补充共识必须在本次同步中冻结
系统 SHALL 在 `phase01-07` 中把当前新增的前端端策略补充共识纳入正式同步范围，避免其只停留在对话层。

#### Scenario: 同步前端端策略补充共识
- **WHEN** 执行 `phase01-07` 的根级同步
- **THEN** 必须明确 `v0.1` 前端正式交付物为单一 `React Web`
- **AND** 必须明确该前端方案同时考虑 `PC` 与移动浏览器 UI
- **AND** 必须明确当前不引入独立 `React Native` 客户端
- **AND** 必须明确 `PWA` 仅作为可兼容增强方向，而不是 `v0.1` 首轮阻断项

### Requirement: 根级真相源与 phase 文档必须保持单值一致
系统 SHALL 要求 `AGENTS.md`、`plan.md`、`architecture_map.md`、`project_rules.md` 与 `phase01` 文档在阶段状态、入口路径与当前冻结结论上保持单值一致。

#### Scenario: 校验根级文档与 phase 文档一致性
- **WHEN** 接手者对根级文档和 `phase01` 文档做一致性检查
- **THEN** 必须确认项目当前阶段、下一阶段入口、目录口径、正式规格入口与当前冻结结论一致
- **AND** 不得出现根级文档已经更新但 `phase01` 文档仍保留旧结论，或相反情况

## MODIFIED Requirements

### Requirement: Phase01 收口后的正式入口
`phase01` 收口完成后，后续实现与 `phase02` SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `architecture_map.md` 为准，执行层规格以 `phase01-06` 的正式 MVP 规格正文为准。

#### Scenario: 后续阶段读取入口
- **WHEN** 后续接手者开始 `phase02`
- **THEN** 必须先从根级文档读取项目当前状态与目录入口
- **AND** 再从 `phase01-06` 正式 MVP 规格正文读取执行层边界
- **AND** 不得把任何单个子规格或对话结论当作新的长期主入口
