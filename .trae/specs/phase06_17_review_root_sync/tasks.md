# Tasks

- [x] Task 1: 完成 `phase06` 全链路审核。
  - [x] SubTask 1.1: 检查 `phase06` 三件套、`phase06-12` 正式规格正文、`phase06-13` 合同主线、`phase06-14 / 15` 实现规格与 `phase06-16` 联调验收报告之间的入口关系与状态表达。
  - [x] SubTask 1.2: 确认 `phase06-12` 仍为 `Onboarding + Data Sovereignty + Reuse Awareness` 当前阶段唯一规格入口。
  - [x] SubTask 1.3: 确认 `phase06-13`、`phase06-14 / 15` 与 `phase06-16` 已进入当前阶段收口语义，不再保留“待继续进入 /spec / 实现 / 验收”的旧表达。

- [x] Task 2: 冻结根级状态回写边界。
  - [x] SubTask 2.1: 明确 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 的同步职责。
  - [x] SubTask 2.2: 明确根级回写只更新状态、入口、冻结结论、切换条件与活动文档索引，不复制 phase 正文与验收细节。
  - [x] SubTask 2.3: 识别当前根级文档中仍保留的 `phase06` 旧状态、旧目标或旧入口表达。

- [x] Task 3: 冻结 `plan.md` 的状态更新与下一阶段切换条件。
  - [x] SubTask 3.1: 明确 `phase06` 状态必须更新为与 `phase06-16` 验收结论一致的收口表达。
  - [x] SubTask 3.2: 明确当前目标必须从“正在交付”切换为与收口结果一致的表达。
  - [x] SubTask 3.3: 明确下一阶段只能以“正式 phase 入口建立后切换”的受控条件表达，不得预设 `phase07` 名称。

- [x] Task 4: 冻结下一阶段入口与上游输入表达。
  - [x] SubTask 4.1: 明确下一阶段必须直接承接 `phase06-12` 正式规格正文、`phase06-13` 合同主线与 `phase06-16` 验收结论。
  - [x] SubTask 4.2: 明确根级文档不得凭空写入任何未建立的下一阶段名称。
  - [x] SubTask 4.3: 明确若下一阶段正式 phase 名称或三件套尚未建立，必须保留“待建立后切换”的显式状态。

- [x] Task 5: 冻结 `phase06` 活动文档入口与收口后角色归位。
  - [x] SubTask 5.1: 明确 `docs/README.md` 与 `architecture_map.md` 必须能稳定指向 `phase06` 三件套、`phase06-12`、`phase06-13` 与 `phase06-16`。
  - [x] SubTask 5.2: 明确 `phase06` 三件套在收口后只承担该阶段规划与冻结记录角色，不再承担根级当前状态说明。
  - [x] SubTask 5.3: 明确 `phase06-14 / 15` 只保留为实现过程记录，不提升为新的长期主入口。
  - [x] SubTask 5.4: 识别并消除可能造成活动文档孤岛的入口缺口。

- [x] Task 6: 完成规格校验。
  - [x] SubTask 6.1: 验证 `phase06` 全链路审核范围与通过标准已经明确。
  - [x] SubTask 6.2: 验证根级状态回写边界与目标文档已经明确。
  - [x] SubTask 6.3: 验证下一阶段入口条件与“不得预设 phase07”边界已经明确。
  - [x] SubTask 6.4: 验证活动文档无孤岛与角色归位要求已经明确。
  - [x] SubTask 6.5: 验证本次收口没有引入第二套入口或第二套阶段状态。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`

# 预期同步目标索引

- `AGENTS.md` — 当前阶段状态、当前主目标、当前下一阶段入口
- `plan.md` — `phase06` 状态、当前目标、当前下一阶段切换条件
- `docs/README.md` — 当前活动文档入口与 `phase06` 收口后导航
- `architecture_map.md` — `phase06` 活动文档、规格入口、验收入口与目录落点
- `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_*` — `phase06` 规划与冻结记录的收口后角色
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/` — `phase06` 正式规格入口
- `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/` — `phase06` 合同主线入口
- `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md` — `phase06` 联调验收与收口结论入口
