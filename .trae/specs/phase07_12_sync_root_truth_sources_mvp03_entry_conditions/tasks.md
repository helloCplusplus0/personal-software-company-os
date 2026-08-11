# Tasks

- [x] Task 1: 完成 `phase07` 收口事实审核。
  - [x] SubTask 1.1: 检查 `phase07` 三件套、`phase07-07` 正式规格、`phase07-08 / 09 / 10` 实现结论与 `phase07-11` 验收报告之间的入口关系与状态表达。
  - [x] SubTask 1.2: 确认 `phase07-07` 仍为传输主线迁移阶段的正式规格入口。
  - [x] SubTask 1.3: 确认 `phase07-11` 已给出“可进入 `phase07-12` 根级收口”的明确结论。

- [x] Task 2: 冻结根级同步边界。
  - [x] SubTask 2.1: 明确 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 的同步职责。
  - [x] SubTask 2.2: 明确根级回写只更新状态、入口、冻结结论与进入条件，不复制实现细节和验收正文。
  - [x] SubTask 2.3: 识别当前根级文档中仍保留的 `phase07` 进行中旧口径。

- [x] Task 3: 冻结 `mvp0.3` 业务阶段进入条件表达。
  - [x] SubTask 3.1: 明确后续 `mvp0.3` 业务阶段必须直接承接 `phase07-07` 正式规格与 `phase07-11` 验收结论。
  - [x] SubTask 3.2: 明确只允许表达进入条件与上游输入，不得提前命名新的业务 phase。
  - [x] SubTask 3.3: 明确若后续正式 phase 入口尚未建立，只能写“待正式 phase 入口建立后切换”或等价受控表达。

- [x] Task 4: 冻结 `plan.md` 的状态更新与角色切换表达。
  - [x] SubTask 4.1: 明确 `phase07` 状态必须更新为基础主线迁移已完成的收口表达。
  - [x] SubTask 4.2: 明确当前主目标不得继续保留 `phase07-08 ~ 12` 的进行中串行任务表述。
  - [x] SubTask 4.3: 明确后续目标必须写成 `mvp0.3` 业务阶段进入条件，而不是新的 phase 名称。

- [x] Task 5: 冻结活动文档入口与收口后角色归位。
  - [x] SubTask 5.1: 明确 `docs/README.md` 与 `architecture_map.md` 必须能稳定指向 `phase07` 三件套、`phase07-07` 与 `phase07-11`。
  - [x] SubTask 5.2: 明确 `phase07` 三件套在收口后只承担该阶段规划与冻结记录角色，不再承担根级当前状态说明。
  - [x] SubTask 5.3: 明确 `phase07-08 / 09 / 10` 只保留为实现结论与迁移记录，不提升为新的根级长期入口。
  - [x] SubTask 5.4: 识别并消除可能造成活动文档孤岛的入口缺口。

- [x] Task 6: 完成规格校验。
  - [x] SubTask 6.1: 验证 `phase07` 收口事实审核范围与通过标准已经明确。
  - [x] SubTask 6.2: 验证根级同步边界、目标文档与禁止事项已经明确。
  - [x] SubTask 6.3: 验证 `mvp0.3` 进入条件与“不提前猜测阶段名称”边界已经明确。
  - [x] SubTask 6.4: 验证 `docs/README.md` 与 `architecture_map.md` 的入口互链要求已经明确。
  - [x] SubTask 6.5: 验证本规格未引入第二套根级入口、第二套阶段状态或未来阶段既成事实口径。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`

# 预期同步目标索引

- `AGENTS.md` — `phase07` 的完成状态、当前项目入口与后续 `mvp0.3` 业务阶段进入条件
- `plan.md` — `phase07` 收口状态、当前主目标切换与后续进入条件
- `docs/README.md` — `phase07` 活动文档入口、正式规格入口与验收入口
- `architecture_map.md` — `phase07` 文档目录落点、根级入口职责与收口后角色说明
- `docs/phase/phase07_transport_contract_mainline_migration_*` — `phase07` 规划与冻结记录的收口后角色
- `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/` — `phase07` 正式规格入口
- `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md` — `phase07` 验收与收口结论入口
