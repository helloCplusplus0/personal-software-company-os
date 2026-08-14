# Tasks

- [x] Task 1: 完成 `phase10` 收口事实审核。
  - [x] SubTask 1.1: 检查 `phase10` 三件套、`phase10-08 / 09 / 10` 实现结论与 `phase10-11` 验收结论之间的入口关系与状态表达。
  - [x] SubTask 1.2: 确认 `phase10-11` 已给出“可进入 `phase10-12` 根级同步”的明确结论。
  - [x] SubTask 1.3: 确认当前根级文档中哪些位置仍保留 `phase10` 只处于 `/plan`、待继续验收或待继续收口的旧口径。

- [x] Task 2: 冻结根级同步边界。
  - [x] SubTask 2.1: 明确 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 的同步职责。
  - [x] SubTask 2.2: 明确根级回写只更新状态、入口、冻结结论与进入条件，不复制实现细节、浏览器日志与调试正文。
  - [x] SubTask 2.3: 明确 `phase10` 在本次回写中的角色切换边界，以及 `phase09` 在本次回写中的角色保持边界。

- [x] Task 3: 冻结后续 `Agent Consumption Layer` 的进入条件表达。
  - [x] SubTask 3.1: 明确下一阶段必须直接承接 `phase10-11` 验收结论与当前根级真相源。
  - [x] SubTask 3.2: 明确只允许表达 `Agent Consumption Layer` 的进入条件与上游输入，不得提前写成已启动事实或当前阶段。
  - [x] SubTask 3.3: 明确若后续正式入口尚未建立，只能写“`phase10` 收口后再进入”“待正式入口建立后切换”或等价受控表达。

- [x] Task 4: 冻结 `plan.md` 与 `AGENTS.md` 的状态更新与角色切换表达。
  - [x] SubTask 4.1: 明确 `phase10` 状态必须更新为正式业务 phase 已完成收口的表达。
  - [x] SubTask 4.2: 明确 `phase09` 继续保留为最近完成正式支撑能力 phase 的表达。
  - [x] SubTask 4.3: 明确后续目标必须写成 `Agent Consumption Layer` 的进入条件，而不是新的当前阶段事实。

- [x] Task 5: 冻结活动文档入口与收口后角色归位。
  - [x] SubTask 5.1: 明确 `docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 必须稳定指向 `phase10` 三件套与 `phase10-11` 验收入口。
  - [x] SubTask 5.2: 明确 `phase10` 三件套在收口后承担“最近完成正式业务 phase 的规划与冻结记录”角色。
  - [x] SubTask 5.3: 明确 `phase09` 继续保留为最近完成正式支撑能力 phase 的规划、冻结与验收入口。
  - [x] SubTask 5.4: 识别并消除可能造成 `phase10` 活动文档孤岛的入口缺口。

- [x] Task 6: 完成规格校验。
  - [x] SubTask 6.1: 验证 `phase10` 收口事实审核范围与通过标准已经明确。
  - [x] SubTask 6.2: 验证根级同步边界、目标文档与禁止事项已经明确。
  - [x] SubTask 6.3: 验证 `Agent Consumption Layer` 进入条件与“不提前进入下一阶段”边界已经明确。
  - [x] SubTask 6.4: 验证 `docs/README.md`、`architecture_map.md`、`docs/phase/README.md` 的入口互链要求已经明确。
  - [x] SubTask 6.5: 验证本规格未引入第二套根级入口、第二套阶段状态或下一阶段既成事实口径。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`

# 预期同步目标索引

- `AGENTS.md` — `phase10` 的完成状态、当前项目入口与 `Agent Consumption Layer` 进入条件
- `plan.md` — `phase10` 收口状态、当前主目标切换与下一阶段进入条件
- `docs/README.md` — `phase10` 当前/最近完成阶段入口、正式验收入口与 docs 工作流入口
- `architecture_map.md` — `phase10` 文档目录落点、根级入口职责与收口后角色说明
- `docs/phase/README.md` — `phase10` 三件套与最近完成阶段入口的收口后角色
- `docs/phase/phase10_asset_action_closure_foundation_*` — `phase10` 规划与冻结记录的收口后角色
- `.trae/specs/phase10_11_complete_asset_action_closure_integration_browser_regression_validation/` — `phase10` 验收与收口结论入口

# 最小修复任务

- [x] Task 7: 清理 `docs/README.md` 中残留的旧下一阶段口径。
  - [x] SubTask 7.1: 删除 [docs/README.md](file:///home/dell/Projects/personal-software-company-os/docs/README.md#L96-L97) 中仍将 `dry-run` 写为后续进入条件的旧表述，避免与 `phase10` 收口后仅保留 `Agent Consumption Layer` 进入条件的单值口径冲突。
  - [x] SubTask 7.2: 复核 `docs/README.md` 当前阶段状态段，仅保留 `phase10` / `phase09` 收口角色、`phase10-11` 验收入口与 `Agent Consumption Layer` 进入条件，不新增第二套下一阶段叙事。
