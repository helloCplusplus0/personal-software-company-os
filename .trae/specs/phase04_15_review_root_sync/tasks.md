# Tasks

- [x] Task 1: 完成 `phase04` 文档互链复核。
  - [x] SubTask 1.1: 检查 `phase04` 三件套、`phase04-10` 正式规格正文、`phase04-11 / 12 / 13 / 14` 相关文档之间的入口关系与状态表达
  - [x] SubTask 1.2: 确认 `product_repository_binding_spec_v0.1.md` 仍为 `Product / Repository / Binding` 当前阶段唯一规格入口
  - [x] SubTask 1.3: 确认 `phase04-14` 的联调验收结论、`phase04-11` 合同主线与 `phase04-12 / 13` 实现结论已纳入当前阶段收口语义

- [x] Task 2: 冻结根级状态回写边界。
  - [x] SubTask 2.1: 明确 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 的同步职责
  - [x] SubTask 2.2: 明确根级回写只更新状态、入口、完成标志与冻结结论，不复制 phase 正文与验收细节
  - [x] SubTask 2.3: 识别当前仍保留旧状态或旧入口的根级文档位置

- [x] Task 3: 冻结 `plan.md` 的状态更新与当前阶段切换要求。
  - [x] SubTask 3.1: 明确 `phase04` 状态必须更新为与交付结果一致的表达
  - [x] SubTask 3.2: 明确 `phase04` 完成标志必须与 `phase04-14` 的运行证据一致
  - [x] SubTask 3.3: 明确 `phase05` 作为下一阶段入口在 `plan.md` 中的表达方式

- [x] Task 4: 冻结 `phase05` 的进入条件与上游输入。
  - [x] SubTask 4.1: 明确 `phase05_dashboard_feedback_foundation` 直接承接 `phase04` 已交付的 `Product / Repository / Binding` 主线
  - [x] SubTask 4.2: 明确 `phase05` 不重复实现 `Product Registry`、`Repository Binding` 与三类绑定动作，而是推进 `Dashboard + Feedback` 最小闭环
  - [x] SubTask 4.3: 明确 `phase05` 的上游输入至少包含 `phase04-10` 正式规格、`phase04-11` 合同主线与 `phase04-14` 验收结论

- [x] Task 5: 完成规格校验。
  - [x] SubTask 5.1: 验证根级文档与 `phase04` 文档的同步目标已经明确
  - [x] SubTask 5.2: 验证 `plan.md` 状态更新要求已经明确
  - [x] SubTask 5.3: 验证 `phase05` 进入条件已经清楚
  - [x] SubTask 5.4: 验证本次收口没有引入第二套入口或第二套阶段状态

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`

# 预期同步目标索引

- `AGENTS.md` — 当前阶段状态、当前主目标、下一阶段入口
- `plan.md` — `phase04` 状态、当前阶段完成标志、`phase05` 入口表达
- `docs/README.md` — 当前活动文档入口与 phase 状态导航
- `architecture_map.md` — 活动文档与 `phase04/phase05` 目录入口
- `docs/phase/phase04_product_and_repository_binding_foundation_*` — `phase04` 收口状态与根级口径一致性
- `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md` — `Product / Repository / Binding` 正式规格入口
- `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/` — `phase04` 联调验收与收口结论入口
