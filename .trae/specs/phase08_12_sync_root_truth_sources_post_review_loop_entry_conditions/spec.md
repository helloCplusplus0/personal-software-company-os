# Phase08-12 完成根级同步与后续进入条件回写 Spec

## Why

`phase08-08` 到 `phase08-11` 已完成 review 合同落地、Dashboard 双入口与双会话承接、`Feedback -> Decision -> Update` 最小闭环，以及统一联调、浏览器验收与反回归验证。现在项目已经具备完成 `phase08` 根级收口的事实基础，但如果根级真相源不及时同步，仓库仍会同时保留“phase08 正在推进中”与“phase08 已完成经营回路验收”两套状态。

`phase08-12` 的目标，是把 `phase08` 作为当前项目首个正式业务 phase 的完成结论、正式证据入口、以及后续支撑能力 phase / dry-run phase 的进入条件，统一回写到根级入口文档中，同时严格避免提前猜测任何尚未建立的后续 phase 名称、编号或路线故事。

## What Changes

- 冻结 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 的根级同步边界
- 冻结 `phase08` 在根级入口中的收口定位：`Operating Review Loop` 已完成正式交付、验证与收口，当前作为最近完成的正式业务 phase 保留
- 冻结 `phase08-11` 正式验收结论进入根级真相源的同步关系
- 冻结 `phase07` 在根级入口中的角色回退：只保留为最近完成的前置基础阶段规划与冻结记录，不再覆盖根级当前业务状态
- 冻结后续支撑能力 phase 与 dry-run phase 的进入条件表达，只允许写条件与上游输入，不允许预设未建立的正式 phase 名称
- 冻结 `docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 对 `phase08` 活动文档、正式验收结论与最近完成阶段入口的稳定互链
- **BREAKING**：根级文档不得继续保留“phase08 仍待继续联调/继续验收/继续收口”的旧表达
- **BREAKING**：根级文档不得凭空写入任何尚未建立的支撑能力 phase、dry-run phase 名称、编号或拆分方案

## Impact

- Affected specs:
  - `phase08_operating_review_loop_foundation`
  - `phase08_08_land_review_contract_backend_frontend_owner_enablement`
  - `phase08_09_land_dashboard_review_entry_dual_session_unified_action_handoff`
  - `phase08_10_land_feedback_decision_update_closed_loop_result_writeback_cleanup`
  - `phase08_11_validate_review_loop_integration_browser_regression_acceptance`
  - `phase07_12_sync_root_truth_sources_mvp03_entry_conditions`
- Affected code:
  - 当前无业务代码改动，影响 `AGENTS.md`
  - 影响 `plan.md`
  - 影响 `docs/README.md`
  - 影响 `architecture_map.md`
  - 影响 `docs/phase/README.md`
  - 影响 `docs/phase/phase08_operating_review_loop_foundation_*`
  - 影响 `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`

## ADDED Requirements

### Requirement: Phase08 收口事实必须先经过根级同步审核

系统 SHALL 在 `phase08-12` 中先审核 `phase08` 是否已经完成“review 合同与 owner 落地 → Dashboard 双入口与双路径会话 → `Feedback -> Decision -> Update` 最小闭环 → 统一联调、浏览器验收与反回归”的全链路，再允许回写根级状态。

#### Scenario: 审核 phase08 收口事实

- **WHEN** 接手者执行 `phase08-12`
- **THEN** 必须检查 `phase08` 三件套、`phase08-08 / 09 / 10` 实现结论与 `phase08-11` 正式验收结论之间的入口关系与状态表达
- **AND** 必须确认 `phase08-11` 已给出“可进入 `phase08-12` 根级同步”的明确结论
- **AND** 若任一链路仍停留在“待实现”“待验收”或“待收口”状态，则不得宣告 `phase08` 已完成根级收口

### Requirement: 根级真相源必须反映 phase08 作为最近完成正式业务 phase 的收口定位

系统 SHALL 在 `phase08-12` 中把 `phase08` 明确表述为当前项目已完成收口的正式业务 phase，并作为后续支撑能力 phase / dry-run phase 建立正式入口前的直接上游，而不是继续把它写成“当前仍在执行中的未收口阶段”。

#### Scenario: 回写根级入口

- **WHEN** `phase08` 审核通过并进入根级同步
- **THEN** 必须同步 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 与 `docs/phase/README.md`
- **AND** 回写内容只允许包含：
  - `phase08` 的完成状态
  - `phase08` 作为最近完成正式业务 phase 的角色表达
  - `phase08` 的活动三件套入口与 `phase08-11` 验收入口
  - 后续支撑能力 phase / dry-run phase 的进入条件
- **AND** 不得把浏览器验收细节、运行时日志、问题复测过程等正文原样复制到根级文档

### Requirement: 根级真相源必须明确 phase07 的角色回退

系统 SHALL 在 `phase08-12` 中清楚表达 `phase07` 已退回最近完成前置基础阶段的规划与冻结记录角色，不再覆盖当前业务主线状态。

#### Scenario: phase07 角色归位

- **WHEN** 根级文档回写 `phase08` 收口状态
- **THEN** 必须明确 `phase07` 三件套只保留为最近完成阶段的规划与冻结记录
- **AND** 必须明确 `phase07-07` 正式规格与 `phase07-11` 验收结论继续作为执行层正式证据入口
- **AND** 不得再把 `phase07-08 / 09 / 10` 提升为根级长期主入口

### Requirement: 后续支撑能力 phase 与 dry-run phase 进入条件必须明确但不得预设正式名称

系统 SHALL 在 `phase08-12` 中清楚表达后续支撑能力 phase 与 dry-run phase 的进入条件，但只允许写“条件”和“上游输入”，不允许猜测下一阶段的正式名称、编号或拆分方案。

#### Scenario: 表达后续进入条件

- **WHEN** 接手者查询 `phase08` 收口后的下一步
- **THEN** 必须能看到清楚的后续进入条件表达
- **AND** 条件必须直接承接 `phase08-11` 验收结论、`phase08` 根级收口结论与当前根级真相源
- **AND** 该表达只允许使用“后续支撑能力 phase”“后续 dry-run phase”“待正式 phase 入口建立后切换”或等价受控说法
- **AND** 不得在根级文档中提前写入任何未建立的阶段名称

### Requirement: 根级文档与 docs 入口必须保持单值一致且无孤岛

系统 SHALL 在 `phase08-12` 中完成根级真相源与 `docs` 入口文档的互链复核，确保后续接手者无需依赖对话记忆即可找到 `phase08` 的正式收口证据与最近完成阶段入口。

#### Scenario: 复核入口一致性

- **WHEN** 接手者从根级入口或 `docs/README.md` 查找 `phase08` 文档
- **THEN** 必须能找到 `phase08` 三件套、`phase08-11` 验收结论与 `phase08` 当前/最近完成阶段入口
- **AND** `docs/phase/README.md` 必须明确 `phase08` 三件套的角色与后续阶段切换条件
- **AND** `architecture_map.md` 必须明确这些入口在收口后的职责与目录落点
- **AND** 不得存在只能通过搜索结果、历史对话或偶然路径发现的活动文档

### Requirement: Plan 状态必须从进行中业务 phase 切换为收口后角色表达

系统 SHALL 要求 `plan.md` 中关于 `phase08` 的状态、当前主目标与后续目标表达，更新为与真实收口结果一致的单值口径。

#### Scenario: 更新 `plan.md`

- **WHEN** `phase08-12` 完成根级同步
- **THEN** `plan.md` 中 `phase08` 的状态必须从“当前正式业务 phase 推进中”切换为“已完成收口的最近正式业务 phase”的表达
- **AND** 当前主目标不得继续写成 `phase08-05 ~ 12` 的进行中串行任务
- **AND** 后续目标必须写成支撑能力 phase / dry-run phase 的进入条件，而不是新的 phase 名称

## MODIFIED Requirements

### Requirement: Phase08 收口后的合法入口

`phase08` 收口完成后，后续接手与下一阶段推进 SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `docs/README.md` / `architecture_map.md` / `docs/phase/README.md` 为准，`phase08` 的执行层验收边界以 `phase08-11` 正式验收结论为准。

#### Scenario: 后续阶段读取入口

- **WHEN** 后续接手者在 `phase08` 收口后继续推进项目
- **THEN** 必须先从根级文档读取项目当前状态、`phase08` 的完成结论与后续进入条件
- **AND** 再从 `phase08-11` 读取当前 review loop 的正式验收边界与收口结论
- **AND** 不得把 `phase08-08 / 09 / 10` 的局部实现结论或零散复测记录提升为新的长期主入口

## REMOVED Requirements

### Requirement: 保留 phase08 仍处于“继续做联调/继续做验收”的旧表达

**Reason**: `phase08` 已完成正式实现、统一联调、双路径浏览器验收与反回归验证，继续保留进行中表述会导致根级真相源与实际阶段状态冲突。

**Migration**: 改为在根级文档中明确 `phase08` 已完成 `Operating Review Loop` 的正式交付与统一验收，并以 `phase08-11` 作为正式验收入口；后续只表达进入条件，不预设阶段名称。

### Requirement: 直接在根级文档中猜测支撑能力 phase 或 dry-run phase 的正式名称

**Reason**: 当前只需要为后续阶段建立进入条件，不需要也不允许在没有正式入口时自由命名未来阶段。

**Migration**: 改为使用“后续支撑能力 phase”“后续 dry-run phase”与“待正式 phase 入口建立后切换”的受控表达。
