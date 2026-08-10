# Phase06-17 审核与根级同步 Spec

## Why

`phase06-01` 到 `phase06-16` 已覆盖 `Onboarding + Data Sovereignty + Reuse Awareness` 的边界冻结、正式规格正文、`.proto` 合同主线、后端与前端实现，以及联调验收收口。但根级真相源当前仍保留“`phase06` 只完成 `/plan`、待进入 `/spec -> 实现 -> 验收 -> 收口`”的旧口径，这会与已经形成的 `phase06-12 / 13 / 16` 收口事实冲突。

`phase06-17` 的目标，是审核 `phase06` 是否已经完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路，把阶段完成事实同步回根级真相源、目录入口与活动文档索引，并明确下一阶段只能以“待正式 phase 入口建立后切换”的受控表达存在，不得在根级文档中越权预设 `phase07` 名称。

## What Changes

- 冻结 `phase06` 收口审核的最小检查范围与通过标准
- 冻结根级状态回写的目标文档、允许更新内容与禁止事项
- 冻结 `plan.md` 中 `phase06` 状态、当前目标与下一阶段切换条件的正式表达
- 冻结 `phase06` 正式规格、合同主线、实现结论与联调验收结论进入根级真相源的同步关系
- 冻结 `docs/README.md`、`architecture_map.md` 与 `phase06` 活动文档之间的入口互链，确保无孤岛
- 冻结“当前阶段规划文档转为已收口阶段记录”的表达边界
- **BREAKING**：根级文档不得继续保留“`phase06` 待进入 `/spec`”或任何等价旧状态表达
- **BREAKING**：根级文档不得凭空写入 `phase07` 或任何未建立正式入口的下一阶段名称

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - 下一阶段正式 phase 入口（若已建立，则作为 `phase06-17` 的直接下游；若未建立，仅允许保留受控切换条件）
- Affected code:
  - 当前无业务代码改动，影响 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`
  - 影响 `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_*`
  - 影响 `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/`
  - 影响 `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/`
  - 影响 `.trae/specs/phase06_14_implement_backend_data_script_mainline/`
  - 影响 `.trae/specs/phase06_15_implement_frontend_mainline/`
  - 影响 `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md`

## ADDED Requirements

### Requirement: Phase06 全链路完成事实必须先经过审核

系统 SHALL 在 `phase06-17` 中先审核 `phase06` 是否已经完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路，再允许回写根级状态。

#### Scenario: 审核 phase06 全链路完成事实

- **WHEN** 接手者执行 `phase06-17`
- **THEN** 必须检查 `phase06` 三件套、`phase06-12` 正式规格正文、`phase06-13` 合同主线、`phase06-14 / 15` 实现结论与 `phase06-16` 联调验收结论之间的入口关系与状态表达
- **AND** 必须确认 `phase06-12` 的正式规格正文仍是当前阶段唯一规格收敛入口
- **AND** 必须确认 `phase06-13` 合同主线、`phase06-14 / 15` 实现结论与 `phase06-16` 联调验收结论已进入 `phase06` 收口语义
- **AND** 若任一链路仍停留在“仅规划”或“仅部分实现”状态，则不得宣告 `phase06` 已收口

### Requirement: 根级状态必须与 phase06 实际收口结论单值一致

系统 SHALL 在 `phase06-17` 中明确根级状态回写的目标文档、同步边界与允许回写内容，确保根级状态与 `phase06` 实际收口结论单值一致。

#### Scenario: 回写根级状态

- **WHEN** `phase06` 审核通过并进入根级同步
- **THEN** 必须同步 `AGENTS.md`、`plan.md`、`docs/README.md` 与 `architecture_map.md`
- **AND** 回写内容只允许包含：
  - 当前阶段状态
  - 当前主目标的完成结论
  - `phase06` 的正式规格入口、合同入口与验收入口
  - 下一阶段切换条件
  - 活动文档入口与归档后角色表达
- **AND** 不得把任务清单、实现细节、异常矩阵或验收正文原样复制到根级文档

### Requirement: Plan 状态更新必须反映 phase06 已收口

系统 SHALL 要求 `plan.md` 中 `phase06` 的状态、当前目标与下一阶段切换条件更新为与真实交付结果一致的单值表达。

#### Scenario: 更新 plan 中 phase06 状态

- **WHEN** `phase06-17` 完成审核与根级同步
- **THEN** `plan.md` 中 `phase06` 的状态必须从“已建立 `/plan` 入口，待继续进入 `/spec`”更新为与 `phase06-16` 验收结论一致的收口状态
- **AND** 当前目标不得继续写成“交付最小主线”这一进行中表达，而应切换为已完成结论或收口后角色表达
- **AND** 当前下一阶段入口必须继续使用受控切换条件表达，不得越权预设新的 phase 名称

### Requirement: 下一阶段入口必须受控且不得预设 phase07 名称

系统 SHALL 在 `phase06-17` 中明确下一阶段入口的表达边界，使后续阶段直接承接 `phase06` 已冻结结果，但在新的正式 phase 入口建立前，不得在根级文档中写入 `phase07` 或等价未来阶段名。

#### Scenario: 确认下一阶段入口

- **WHEN** 接手者查询 `phase06` 收口后的下一阶段如何进入
- **THEN** 必须能得到清楚的切换条件表达
- **AND** 该表达必须直接承接 `phase06-12` 正式规格正文、`phase06-13` 合同主线与 `phase06-16` 验收结论
- **AND** 若下一阶段正式 phase 名称、三件套或 dev plan 尚未建立，根级文档只能写“待下一阶段正式 phase 入口建立后切换”或等价受控表达
- **AND** 不得在 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 中出现 `phase07` 或任何未建立阶段的既成事实口径

### Requirement: 活动文档入口必须无孤岛

系统 SHALL 在 `phase06-17` 中完成 `phase06` 活动文档、根级真相源与 `docs/README.md` / `architecture_map.md` 的入口互链复核，确保收口后不存在孤立文档。

#### Scenario: 复核活动文档入口

- **WHEN** 接手者从根级入口或 `docs/README.md` 查找 `phase06` 文档
- **THEN** 必须能找到 `phase06` 三件套、`phase06-12` 正式规格入口、`phase06-13` 合同入口与 `phase06-16` 验收结论入口
- **AND** `architecture_map.md` 必须明确这些入口在收口后的职责与目录落点
- **AND** 不得存在只能通过对话记忆、搜索结果或历史路径偶然发现的活动文档

### Requirement: phase06 规划、规格与验收结论必须以收口后角色归位

系统 SHALL 在 `phase06-17` 中明确 `phase06` 三件套、正式规格正文、实现规格与验收报告在收口后的角色，不允许它们继续混充“项目当前进行中状态”或互相并列争夺入口。

#### Scenario: 归位 phase06 文档角色

- **WHEN** `phase06` 完成根级同步
- **THEN** `phase06` 三件套必须被表述为该阶段的规划与冻结记录
- **AND** `phase06-12` 必须继续作为该阶段正式规格正文入口
- **AND** `phase06-13` 必须继续作为该阶段合同主线入口
- **AND** `phase06-14 / 15` 必须保留为实现过程记录，不上升为新的根级真相源
- **AND** `phase06-16` 必须作为该阶段联调验收与收口结论入口

## MODIFIED Requirements

### Requirement: Phase06 收口后的合法入口

`phase06` 收口完成后，后续接手与下一阶段 SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `docs/README.md` / `architecture_map.md` 为准，`Onboarding + Data Sovereignty + Reuse Awareness` 的执行层边界以 `phase06-12` 正式规格正文、`phase06-13` 合同主线与 `phase06-16` 验收结论为准。

#### Scenario: 后续阶段读取入口

- **WHEN** 后续接手者在 `phase06` 收口后继续推进项目
- **THEN** 必须先从根级文档读取项目当前状态、阶段入口与切换条件
- **AND** 再从 `phase06-12`、`phase06-13` 与 `phase06-16` 读取 `phase06` 已交付边界
- **AND** 不得把 `phase06-14 / 15` 的实现规格、对话结论或临时验收命令提升为新的长期主入口

## REMOVED Requirements

### Requirement: 保留 phase06 仍处于“待进入 /spec -> 实现 -> 验收 -> 收口”状态的旧表达

**Reason**: `phase06` 已完成正式规格、合同主线、实现与联调验收，继续保留旧状态会导致根级真相源与实际阶段结论冲突。
**Migration**: 改为在根级文档中明确 `phase06` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并以 `phase06-12`、`phase06-13` 与 `phase06-16` 作为当前阶段已交付边界；下一阶段仅在正式入口建立后切换，不预设 `phase07` 名称。
