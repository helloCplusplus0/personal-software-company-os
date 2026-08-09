# Tasks

- [x] Task 1: 完成 `phase05` 文档互链复核。
  - [x] SubTask 1.1: 检查 `phase05` 三件套、`phase05-10` 正式规格正文、`phase05-11 / 12 / 13 / 14` 相关文档之间的入口关系与状态表达。
  - [x] SubTask 1.2: 确认 `dashboard_feedback_spec_v0.1.md` 仍为 `Dashboard + Feedback` 当前阶段唯一规格入口。
  - [x] SubTask 1.3: 确认 `phase05-11` 合同主线、`phase05-12 / 13` 实现结论与 `phase05-14` 联调验收结论已纳入当前阶段收口语义。

- [x] Task 2: 冻结根级状态回写边界。
  - [x] SubTask 2.1: 明确 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 的同步职责。
  - [x] SubTask 2.2: 明确根级回写只更新状态、入口、完成标志与冻结结论，不复制 phase 正文与验收细节。
  - [x] SubTask 2.3: 识别当前仍保留旧状态、旧目标或旧入口的根级文档位置。

- [x] Task 3: 冻结 `plan.md` 的状态更新与阶段切换要求。
  - [x] SubTask 3.1: 明确 `phase05` 状态必须更新为与交付结果一致的表达。
  - [x] SubTask 3.2: 明确 `phase05` 完成标志必须与 `phase05-14` 的运行证据一致。
  - [x] SubTask 3.3: 明确不得继续保留“当前目标是完成 phase05 三件套规划”的旧表达。

- [x] Task 4: 冻结下一阶段入口与上游输入表达。
  - [x] SubTask 4.1: 明确下一阶段必须直接承接 `phase05-10` 正式规格正文、`phase05-11` 合同主线与 `phase05-14` 验收结论。
  - [x] SubTask 4.2: 明确下一阶段入口不得在根级文档中凭空猜测未建立的 phase 名称。
  - [x] SubTask 4.3: 明确若下一阶段正式 phase 入口尚未建立，必须把“待建立后切换”写成显式状态。

- [x] Task 5: 完成规格校验。
  - [x] SubTask 5.1: 验证根级文档与 `phase05` 文档的同步目标已经明确。
  - [x] SubTask 5.2: 验证 `plan.md` 状态更新要求已经明确。
  - [x] SubTask 5.3: 验证下一阶段入口条件与禁止猜测边界已经清楚。
  - [x] SubTask 5.4: 验证本次收口没有引入第二套入口或第二套阶段状态。

- [x] Task 6: 修复 `phase05` 文档中与“下一阶段入口待建立后切换”不一致的未来阶段表达。
  - [x] SubTask 6.1: 将 `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md` 中 `phase06+` 改为不预设阶段名称的受控表达。
  - [x] SubTask 6.2: 复核 `phase05` 相关文档，确保未来引用统一使用“下一阶段正式 phase 入口建立后切换”或等价受控表达。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 5`

# 预期同步目标索引

- `AGENTS.md` — 当前阶段状态、当前主目标、当前下一阶段入口
- `plan.md` — `phase05` 状态、当前阶段完成标志、下一阶段入口表达
- `docs/README.md` — 当前活动文档入口与 phase 状态导航
- `architecture_map.md` — 活动文档与 `phase05/下一阶段` 目录入口
- `docs/phase/phase05_dashboard_feedback_foundation_*` — `phase05` 收口状态与根级口径一致性
- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md` — `Dashboard + Feedback` 正式规格入口
- `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/` — `phase05` 合同主线入口
- `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md` — `phase05` 联调验收与收口结论入口
