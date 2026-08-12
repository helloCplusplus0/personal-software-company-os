# Tasks

- [x] Task 1: 完成 `phase09` 收口事实审核。
  - [x] SubTask 1.1: 检查 `phase09` 三件套、`phase09-08 / 09 / 10` 实现结论与 `phase09-11` 验收报告之间的入口关系与状态表达。
  - [x] SubTask 1.2: 确认 `phase09-11` 已给出“可进入 `phase09-12` 根级同步”的明确结论。
  - [x] SubTask 1.3: 确认当前根级文档中哪些位置仍保留 `phase09` 进行中旧口径。

- [x] Task 2: 冻结根级同步边界。
  - [x] SubTask 2.1: 明确 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 的同步职责。
  - [x] SubTask 2.2: 明确根级回写只更新状态、入口、冻结结论与进入条件，不复制实现细节和验收正文。
  - [x] SubTask 2.3: 明确 `phase08` 在本次回写中的角色归位边界。

- [x] Task 3: 冻结后续 `dry-run` 的进入条件表达。
  - [x] SubTask 3.1: 明确后续 `dry-run` 必须直接承接 `phase09-11` 验收结论与当前根级真相源。
  - [x] SubTask 3.2: 明确只允许表达 `dry-run` 进入条件与上游输入，不得提前扩大 `mvp0.4+`。
  - [x] SubTask 3.3: 明确若后续正式入口尚未建立，只能写“待正式入口建立后切换”或等价受控表达。

- [x] Task 4: 冻结 `plan.md` 与 `AGENTS.md` 的状态更新与角色切换表达。
  - [x] SubTask 4.1: 明确 `phase09` 状态必须更新为正式支撑能力 phase 已完成收口的表达。
  - [x] SubTask 4.2: 明确 `phase08` 在根级入口中退回最近完成正式业务 phase 的角色表达。
  - [x] SubTask 4.3: 明确后续目标必须写成 `dry-run` 进入条件，而不是新的 `mvp0.4+` 范围故事。

- [x] Task 5: 冻结活动文档入口与收口后角色归位。
  - [x] SubTask 5.1: 明确 `docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 必须稳定指向 `phase09` 三件套与 `phase09-11` 验收结论。
  - [x] SubTask 5.2: 明确 `phase09` 三件套在收口后承担“最近完成正式支撑能力 phase 的规划与冻结记录”角色。
  - [x] SubTask 5.3: 明确 `phase08` 继续保留为最近完成正式业务 phase 的规划、冻结与验收入口，不再覆盖当前支撑能力状态。
  - [x] SubTask 5.4: 识别并消除可能造成活动文档孤岛的入口缺口。

- [x] Task 6: 完成规格校验。
  - [x] SubTask 6.1: 验证 `phase09` 收口事实审核范围与通过标准已经明确。
  - [x] SubTask 6.2: 验证根级同步边界、目标文档与禁止事项已经明确。
  - [x] SubTask 6.3: 验证后续 `dry-run` 进入条件与“不提前扩大 `mvp0.4+`”边界已经明确。
  - [x] SubTask 6.4: 验证 `docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 的入口互链要求已经明确。
  - [x] SubTask 6.5: 验证本规格未引入第二套根级入口、第二套阶段状态或未来阶段既成事实口径。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`

# 预期同步目标索引

- `AGENTS.md` — `phase09` 的完成状态、当前项目入口与后续 `dry-run` 进入条件
- `plan.md` — `phase09` 收口状态、当前主目标切换与 `dry-run` 进入条件
- `docs/README.md` — `phase09` 当前/最近完成阶段入口、正式验收入口与 docs 工作流入口
- `architecture_map.md` — `phase09` 文档目录落点、根级入口职责与收口后角色说明
- `docs/phase/README.md` — `phase09` 三件套与最近完成阶段入口的收口后角色
- `docs/phase/phase09_template_reuse_derived_intelligence_foundation_*` — `phase09` 规划与冻结记录的收口后角色
- `.trae/specs/phase09_11_validate_template_reuse_derived_hint_integration_browser_regression_acceptance/acceptance_report.md` — `phase09` 验收与收口结论入口

# 实现总结

## 已更新文档

- `AGENTS.md`
- `plan.md`
- `docs/README.md`
- `architecture_map.md`
- `docs/phase/README.md`

## 回写结果

- 根级状态已统一切换为：`phase09` 已完成正式交付、统一验收与收口，当前作为最近完成正式支撑能力 phase 保留
- `phase09-11` 已进入根级与 docs 入口，作为 `phase09` 正式验收与收口结论入口
- `phase08` 已明确退回最近完成正式业务 phase；`phase08-11` 继续保留为 `Operating Review Loop` 的正式验收与收口结论入口
- 后续 `dry-run` 只保留进入条件表达，不预设任何 `mvp0.4+` 既成事实、未来阶段名称或范围拆分
- `docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 已完成互链同步，不再存在 `phase09` 收口证据只能靠搜索定位的入口缺口
