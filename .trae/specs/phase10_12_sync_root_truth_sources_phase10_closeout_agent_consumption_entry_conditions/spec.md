# Phase10-12 完成根级同步、阶段收口与下一阶段进入条件回写 Spec

## Why

`phase10-08` 到 `phase10-11` 已完成 `Onboarding` 首轮建链、`Decision` 生命周期闭环、关键 detail pages 的下一步动作承接，以及联调、浏览器验收与反回归验证。当前项目已经具备完成 `phase10` 根级收口的事实基础，但根级真相源仍停留在“`phase10` 已建立 `/plan` 入口”的状态，会同时保留“阶段仍在推进”与“阶段已通过正式验收”的双口径。

`phase10-12` 的目标，是把 `phase10` 作为当前最近完成正式业务 phase 的完成结论、正式收口入口，以及后续 `Agent Consumption Layer` 的单值进入条件，统一回写到根级入口文档中；同时明确 `phase09` 继续保留为最近完成正式支撑能力 phase，并严格禁止在 `phase10` 未完成根级收口前提前进入下一阶段。

## What Changes

- 冻结 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 的 `phase10` 根级同步边界
- 冻结 `phase10` 在根级入口中的收口定位：`Asset-Action Closure` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 冻结 `phase10-11` 作为 `phase10` 正式验收与收口证据入口进入根级真相源的同步关系
- 冻结 `phase10` 三件套在收口后的角色：保留为最近完成正式业务 phase 的规划与冻结记录，不再承担“当前进行中阶段”表达
- 冻结 `phase09` 在根级入口中的角色保持：继续作为最近完成正式支撑能力 phase 保留
- 冻结后续 `Agent Consumption Layer` 的进入条件表达，只允许在 `phase10` 收口后出现，且只写条件与上游输入
- 冻结 `docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 对 `phase10` 活动文档、正式验收证据与最近完成阶段入口的稳定互链
- **BREAKING**：根级文档不得继续保留“`phase10` 仍只处于 `/plan` 或待继续验收”的旧表达
- **BREAKING**：根级文档不得在 `phase10` 未完成收口前，把 `Agent Consumption Layer` 写成已启动事实、当前阶段或并行主线

## Impact

- Affected specs:
  - `phase10_08_land_onboarding_first_run_chain_guidance_canonical_handoff`
  - `phase10_09_land_decision_lifecycle_detail_cta_pending_reread_unification`
  - `phase10_10_land_key_detail_pages_next_step_cta_handoff_matrix`
  - `phase10_11_complete_asset_action_closure_integration_browser_regression_validation`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
- Affected code:
  - 当前无业务代码改动
  - 影响 `AGENTS.md`
  - 影响 `plan.md`
  - 影响 `docs/README.md`
  - 影响 `architecture_map.md`
  - 影响 `docs/phase/README.md`
  - 影响 `docs/phase/phase10_asset_action_closure_foundation_*`
  - 影响 `.trae/specs/phase10_11_complete_asset_action_closure_integration_browser_regression_validation/`

## ADDED Requirements

### Requirement: Phase10 收口事实必须先经过根级同步审核

系统 SHALL 在 `phase10-12` 中先审核 `phase10` 是否已经完成“`Onboarding` 首轮建链 → `Decision` 生命周期闭环 → detail pages 动作承接矩阵 → 联调、浏览器验收与反回归验证”的全链路，再允许回写根级状态。

#### Scenario: 审核 phase10 收口事实

- **WHEN** 接手者执行 `phase10-12`
- **THEN** 必须检查 `phase10` 三件套、`phase10-08 / 09 / 10` 实现结论与 `phase10-11` 正式验收结论之间的入口关系与状态表达
- **AND** 必须确认 `phase10-11` 已给出“可进入 `phase10-12` 根级同步”的明确结论
- **AND** 若任一链路仍停留在“待实现”“待验收”或“待收口”状态，则不得宣告 `phase10` 已完成根级收口

### Requirement: 根级真相源必须反映 phase10 作为最近完成正式业务 phase 的收口定位

系统 SHALL 在 `phase10-12` 中把 `phase10` 明确表述为当前项目已完成收口的最近正式业务 phase，并作为后续 `Agent Consumption Layer` 建立正式入口前的直接上游，而不是继续把它写成“刚建立 `/plan` 的当前进行中阶段”。

#### Scenario: 回写根级入口

- **WHEN** `phase10` 审核通过并进入根级同步
- **THEN** 必须同步 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 与 `docs/phase/README.md`
- **AND** 回写内容只允许包含：
  - `phase10` 的完成状态
  - `phase10` 作为最近完成正式业务 phase 的角色表达
  - `phase10` 三件套入口与 `phase10-11` 验收入口
  - 后续 `Agent Consumption Layer` 的进入条件
- **AND** 不得把浏览器验收细节、后台日志、调试过程或修复过程正文原样复制到根级文档

### Requirement: 根级真相源必须明确 phase09 的角色保持

系统 SHALL 在 `phase10-12` 中清楚表达 `phase09` 继续作为最近完成正式支撑能力 phase 保留，不得因为 `phase10` 收口而丢失其根级角色入口。

#### Scenario: phase09 角色保持

- **WHEN** 根级文档回写 `phase10` 收口状态
- **THEN** 必须明确 `phase09` 仍是最近完成正式支撑能力 phase
- **AND** 必须明确 `phase09-11` 继续保留为 `Template Reuse + Derived Intelligence` 的正式验收与收口结论入口
- **AND** 不得让 `phase10` 收口覆盖或替换 `phase09` 在“最近完成正式支撑能力 phase”上的角色

### Requirement: 后续 Agent Consumption Layer 进入条件必须明确且单值化

系统 SHALL 在 `phase10-12` 中清楚表达后续 `Agent Consumption Layer` 的进入条件，但只允许写“条件”和“上游输入”，不允许提前把该阶段写成已启动事实、当前阶段或平行主线。

#### Scenario: 表达下一阶段进入条件

- **WHEN** 接手者查询 `phase10` 收口后的下一步
- **THEN** 必须能看到清楚的 `Agent Consumption Layer` 进入条件表达
- **AND** 进入条件必须直接承接：
  - `phase10-11` 正式验收结论
  - `phase10` 根级收口结论
  - 当前根级真相源
- **AND** 该表达只允许使用“`phase10` 收口后再进入 `Agent Consumption Layer`”“待正式入口建立后切换”或等价受控说法
- **AND** 不得在根级文档中提前写入 `Agent Consumption Layer` 已开始实现、已建立 `/plan`、已冻结范围，或其后的更远阶段故事

### Requirement: 根级文档与 docs 入口必须保持单值一致且无孤岛

系统 SHALL 在 `phase10-12` 中完成根级真相源与 `docs` 入口文档的互链复核，确保后续接手者无需依赖对话记忆即可找到 `phase10` 的正式收口证据与最近完成阶段入口。

#### Scenario: 复核入口一致性

- **WHEN** 接手者从根级入口或 `docs/README.md` 查找 `phase10` 文档
- **THEN** 必须能找到 `phase10` 三件套与 `phase10-11` 正式验收入口
- **AND** `docs/phase/README.md` 必须明确 `phase10` 三件套在收口后的角色与下一阶段切换条件
- **AND** `architecture_map.md` 必须明确这些入口在收口后的职责与目录落点
- **AND** 不得存在只能通过搜索结果、历史对话或偶然路径发现的活动文档

### Requirement: Plan 状态必须从“已建立 /plan 入口”切换为“已完成收口”

系统 SHALL 要求 `plan.md` 与 `AGENTS.md` 中关于 `phase10` 的状态、当前主目标与后续目标表达，更新为与真实收口结果一致的单值口径。

#### Scenario: 更新 `plan.md` 与 `AGENTS.md`

- **WHEN** `phase10-12` 完成根级同步
- **THEN** `phase10` 的状态必须从“已建立正式 `/plan` 入口”切换为“已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`”
- **AND** 当前主目标不得继续写成 `phase10` 仍在 `/plan` 或执行中
- **AND** 后续目标必须写成 `Agent Consumption Layer` 的进入条件，而不是新的当前阶段事实

## MODIFIED Requirements

### Requirement: Phase10 收口后的合法入口

`phase10` 收口完成后，后续接手与下一阶段推进 SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `docs/README.md` / `architecture_map.md` / `docs/phase/README.md` 为准，`phase10` 的执行层验收边界以 `phase10-11` 正式验收结论为准。

#### Scenario: 后续阶段读取入口

- **WHEN** 后续接手者在 `phase10` 收口后继续推进项目
- **THEN** 必须先从根级文档读取项目当前状态、`phase10` 的完成结论与 `Agent Consumption Layer` 进入条件
- **AND** 再从 `phase10-11` 读取当前 `Asset-Action Closure` 主线的正式验收边界与收口结论
- **AND** 不得把 `phase10-08 / 09 / 10` 的局部实现结论、浏览器复测片段或零散日志提升为新的长期主入口

## REMOVED Requirements

### Requirement: 保留 phase10 仍处于“仅建立 /plan 入口、待继续验收”的旧表达

**Reason**: `phase10` 已完成正式实现、联调、浏览器验收、反回归验证与根级收口准备，继续保留进行中表述会导致根级真相源与实际阶段状态冲突。

**Migration**: 改为在根级文档中明确 `phase10` 已完成 `Asset-Action Closure` 的正式交付与统一验收，并以 `phase10-11` 作为正式验收入口；后续只表达 `Agent Consumption Layer` 进入条件，不提前写成当前阶段事实。

### Requirement: 直接在根级文档中把 Agent Consumption Layer 写成当前既成事实

**Reason**: 当前只需要为后续 `Agent Consumption Layer` 建立进入条件，不需要也不允许在没有正式入口时自由扩写下一阶段内容。

**Migration**: 改为使用“`phase10` 收口后再进入 `Agent Consumption Layer`”“待正式入口建立后切换”的受控表达，只保留进入条件与上游输入。
