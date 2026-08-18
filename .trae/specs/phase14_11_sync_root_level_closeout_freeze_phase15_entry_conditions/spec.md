# phase14-11 完成根级同步、阶段收口与下一阶段进入条件回写 Spec

## Why

`phase14-10` 已按固定样本、固定入口、固定 6 问、八项裁决门禁与固定 rerun 协议完成正式验收（6/6 达标、八项裁决全绿、16/16 页矩阵、工具链四步门禁全绿、独立复核 PASS 0 阻断），`Standard Entity Foundation` 的阶段主线交付（Standard 五层主线 + 画像七触点系统性退役 + T7 裁决 brief 解耦）已完成。但根级五文档仍停留在 `phase14 in_progress` 状态，且 `phase14-10` 期间触发的 T7 用户裁决（brief 画像残余解耦）使 `phase14` `shared_baseline` 中冻结的 `CON-08` 原承接口径（"`current_phase` 三字段随画像主记录保留在 brief"）已失效，`phase15` 进入条件必须按裁决后口径重新冻结，否则下一阶段规划将被旧口径误导。

`phase14-11` 的目标不是继续扩写功能，而是：完成根级同步与 `phase14` 阶段收口，把 `phase15` 进入条件按 T7 裁决后口径单值冻结，彻底完结 `phase14`。

## What Changes

- 回写根级五文档（`AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`）：`phase14` 状态位更新为已完成收口、最近完成正式业务 phase 更新为 `phase14`、验收与收口入口指向 `phase14_10` acceptance_report、登记本 spec 目录入口
- 冻结 `phase15` 进入条件于本 spec（作为 `phase15 /plan` 的直接输入），含后续项排序：`CON-08` 时间轴承接（**T7 裁决后口径：另行新建正规承接，不复活画像派生形态**）、`standard_bindings` 目标类型扩展、agent 写回、Git 推进跟踪 / 模板仓库自动接入 / 自动同步
- 留档 `CON-08` 口径变更链：`phase13-12 CON-08`（时间轴必须从画像剥离）→ `phase14 shared_baseline` §2.2 裁决④原口径（`current_phase` 三字段保留在 brief，时间轴 phase15 承接）→ `phase14-10 T7` 用户裁决（brief 内联画像消息整体删除 + 画像主表 drop，时间轴如进入须新建正规字段）——本 spec 冻结的 phase15 进入条件以 T7 裁决后口径为准
- 本阶段不改任何前后端与数据库代码；只做文档回写与状态收口

## Impact

- Affected specs:
  - `.trae/specs/phase14_10_validate_standard_entity_integration_dogfooding_regression/acceptance_report.md`（验收结论作为本阶段收口依据；其 §9 CON-08 口径变更留痕为本 spec 输入）
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（L75-78 根级同步范围定义）
  - `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`（§2.2 裁决④、L118 保留项表述的时效性由本 spec 重新界定——正文不改写，口径变更以本 spec 留档为准）
- Affected docs（仅回写状态与入口，对齐式更新，不改写阶段正文结论）:
  - `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/README.md`
- Affected code: 无
- 验收产物：本目录 `tasks.md / checklist.md`（根级回写完成后收口）

## ADDED Requirements

### Requirement: phase14 阶段收口状态必须单值回写根级五文档

系统 SHALL 完成以下根级同步，且不留孤岛文档、不破坏单一真相源分工：

- `AGENTS.md`：当前阶段状态更新为 `phase14` 已完成正式收口；最近完成正式业务 phase 更新为 `phase14`（`phase13` / `phase10` / `phase09` 相应退位为历史保留角色）；验收与收口入口指向 `phase14_10` acceptance_report；登记本 spec 为 `phase15` 进入条件冻结入口
- `plan.md`：`phase14` 状态位更新为 `completed` + 当前收口结果（Standard 五层主线 + 画像七触点退役 + T7 brief 解耦）；`§1 当前状态` 更新；`§3 phase14` 条目补收口结论；`phase15` 进入条件登记（指向本 spec）
- `docs/README.md`：根级阶段状态更新；登记 `phase14_10` 验收入口与本 spec 目录入口
- `architecture_map.md`：登记 `phase14_10` 与本 spec 目录落点；`phase14` 三件套角色表述从"待执行"更新为"已完成收口的规划与冻结记录"
- `docs/phase/README.md`：当前阶段状态更新；登记 `phase14_10` 验收入口与本 spec 入口

#### Scenario: 收口后无孤岛且单值一致

- **WHEN** 收口完成后的任意执行者从根级入口（`AGENTS.md / plan.md / docs/README.md`）出发
- **THEN** 五文档对"phase14 已收口"的表述单值一致，能找到 `phase14_10` acceptance_report 与本 spec（phase15 进入条件）
- **AND** 本 spec 目录与 `phase14_10` 目录均不是孤岛文档

### Requirement: phase15 进入条件必须按 T7 裁决后口径单值冻结

`phase15` 的进入条件 SHALL 冻结为：

1. `phase14-11` 根级回写与阶段收口完成（本 spec 收口即满足）
2. `phase15 /plan` MUST 在进入实现拆分前完成后续项排序裁决，候选池与既有约束如下（排序本身属 `phase15 /plan` 裁决范围，本条件只冻结候选池与边界）：
   - **`CON-08` 时间轴承接（T7 裁决后口径）**：开发推进跟踪（当前阶段 / 历史阶段）SHALL 以新建正规承接位落地——独立数据模型 + `.proto` 合同 + web 时间轴展示 + agent 可读；**禁止复活画像派生形态**（`BriefGovernanceProfile / BriefCurrentPhase / BriefTrackType / BriefPhaseStatus` 已删除且字段号 reserved，`governance_profiles` 主表已 drop）；`brief` 中不回填时间轴字段，时间轴经独立入口消费
   - **`standard_bindings` 目标类型扩展**：`CON-02` 可扩展设计已就位（`target_type` CHECK 约束可扩展、不为目标类型建分表）；扩展由真实绑定需求驱动，不预先扩枚举
   - **agent 写回**：`CON-09` 先消费后维护顺序不变——agent 写回在时间轴等消费面稳定且经用户显式裁决后才可进入；不因 `phase14` 完成、不因时间轴进入而自动解锁
   - **Git 推进跟踪 / 模板仓库自动接入 / 自动同步**：继续按既有条件约束（自 `phase11` 起持续顺延的非目标池），须先经独立 `/plan` 裁决排期，不得搭车进入时间轴 phase
3. `phase15 /plan` 直接上游 = 本 spec（phase15 进入条件）+ `phase14_10` acceptance_report（阶段交付实况）；不回放 `phase13-12` 缺口记录（其 CON-01 ~ CON-05 / CON-07 已由 phase14 落地消化，CON-06 / CON-08 / CON-09 的未完成部分已按本 spec口径承接）

#### Scenario: 进入条件可机械判定且口径无冲突

- **WHEN** 用户或执行者评估是否可以开启 `phase15 /plan`
- **THEN** 以上条件可逐条机械判定；时间轴承接的口径唯一指向"新建正规承接、不复活画像派生形态"
- **AND** `phase14` `shared_baseline` L118 旧口径（三字段保留在 brief）不再产生误导——其时效性已被本 spec 显式界定

### Requirement: CON-08 口径变更链必须完整留档

本 spec SHALL 留档以下口径变更链，使后续执行者可追溯 `current_phase` 承接口径的三次演变而无需回放对话：

1. `phase13-12 CON-08`：判定时间轴必须从治理画像剥离，独立承接（后端持久化 + web 展示 + agent 可读）
2. `phase14 shared_baseline` §2.2 裁决④原口径：`current_phase` 三字段随画像主记录收缩后保留在 brief 内联轻量消息，时间轴 phase15 承接
3. `phase14-10 T7` 用户裁决（2026-08-18）：画像为错误设计——brief 内联画像消息（含 `current_phase` 块）整体删除并 reserved、画像主表 drop；时间轴如进入 phase15 须另行新建正规字段，不复活画像派生形态
4. 本 spec：以 3 为最终口径冻结 phase15 进入条件；`phase14` 三件套正文不改写（历史冻结记录），口径时效性以本 spec 为准

#### Scenario: 口径追溯

- **WHEN** 后续执行者对 `current_phase` / 时间轴承接口径产生疑问
- **THEN** 能从本变更链直接判定最终有效口径为 T7 裁决后口径
- **AND** 不需要重新解释 `shared_baseline` 旧表述的适用范围

### Requirement: 根级回写必须为对齐式更新

- 五文档回写 SHALL 采用对齐式更新（更新状态位、入口与收口结论），不推翻重写既有叙事结构
- `phase14` 三件套（architecture_plan / dev_plan / shared_baseline）正文不改写，保留为该阶段 `/plan` 的规划与冻结记录（含已被 T7 取代的旧口径，时效性由本 spec 界定）
- 回写后 `AGENTS.md` 不新增实现细节或 task 清单；`plan.md` 不展示 task 级拆分
- `phase13-12` 缺口记录的角色随本收口退位为历史输入（其未完成项已由本 spec 承接），不再作为活跃直接上游登记

## REMOVED Requirements

### Requirement: `CON-08` 原承接口径（`current_phase` 三字段随画像主记录保留在 brief）

**Reason**: `phase14-10` Task 2 首验实证画像残余依赖阻断 brief 双消费路径（画像写路径退役后主记录行无正式入口可重建 → 全新部署仓库永久 404），用户裁决画像为错误设计、残余依赖必须系统性移除；T7 已落地（brief 内联消息删除 + reserved + 主表 drop）。

**Migration**: 已由 `phase14-10` T7 完成迁移（brief 5 顶层块，槽位 2/3/4 reserved）；时间轴未来承接形态由本 spec 冻结的 phase15 进入条件第 2 条第 1 项约束（新建正规承接、不复活画像派生形态）。
