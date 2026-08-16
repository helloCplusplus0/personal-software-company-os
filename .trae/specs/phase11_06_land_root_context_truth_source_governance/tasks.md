# Tasks

- [x] Task 1: 盘点 `phase11-06` 的实施输入与目标文件范围
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L165-L194` 中当前 `phase11-06` 的范围、补充说明与 DoD
  - [x] SubTask 1.2: 审阅 `phase11_project_context_foundation_architecture_plan.md` 中根级治理矩阵、单一写者规则与最终共识入口的冻结口径
  - [x] SubTask 1.3: 审阅 `phase11_project_context_foundation_shared_baseline.md` 中根级真相源治理矩阵与"悬空引用清零"的冻结口径

- [x] Task 2: 建立根级治理逐文件审计基线
  - [x] SubTask 2.1: 逐项盘点 `README.md / AGENTS.md / plan.md / architecture_map.md / docs/README.md / docs/phase/README.md / project_rules.md / global_skills.md / PSCO-mvp05-summarize-feedback.md`
  - [x] SubTask 2.2: 为每个目标文件记录 `已审计 / 是否修改 / 不修改原因`
  - [x] SubTask 2.3: 明确每个目标文件当前是否存在重复 phase 状态、重复目录落点、重复共识入口或悬空引用

- [x] Task 3: 回收根级入口中的重复主结论
  - [x] SubTask 3.1: 回收重复承载的 phase 状态正文，收敛回 `plan.md`（README.md / docs/README.md / docs/phase/README.md 已改为摘要式引用）
  - [x] SubTask 3.2: 回收重复承载的目录落点正文，收敛回 `architecture_map.md`（审计确认无重复）
  - [x] SubTask 3.3: 回收重复承载的最终共识正文，收敛回当前有效最终共识入口（`phase11-06` 当前时点为 `PSCO-mvp05-summarize-feedback.md`，审计确认已统一指向）

- [x] Task 4: 清理悬空引用并统一最终共识入口
  - [x] SubTask 4.1: 清理治理矩阵目标文件范围内指向不存在文件 `PSCO-summarize-feedback.md` 的引用（审计确认 9 个目标文件中无悬空引用）
  - [x] SubTask 4.2: 将最终共识类入口统一改写为当前有效最终共识入口（`phase11-06` 当前时点为 `PSCO-mvp05-summarize-feedback.md`，但该文件名不作为未来固定合同）
  - [x] SubTask 4.3: 确保不保留历史兼容名义下的失效共识入口（审计确认无）

- [x] Task 5: 同步治理矩阵中的活动入口文件
  - [x] SubTask 5.1: 按单一写者规则同步 `README.md`（移除内联 phase 状态，改为指向 plan.md 的摘要式引用）
  - [x] SubTask 5.2: 按单一写者规则同步 `docs/README.md`（§4 从 24 行详细 phase 状态压缩为 6 行摘要式入口）
  - [x] SubTask 5.3: 按单一写者规则同步 `docs/phase/README.md`（§2 从 17 行详细 phase 状态压缩为 5 行摘要式入口）

- [x] Task 6: 保证"不修改"文件也完成正式审计闭环
  - [x] SubTask 6.1: 对无需改动的目标文件保留明确的不修改原因
  - [x] SubTask 6.2: 确保"不修改"不等于"未审计"
  - [x] SubTask 6.3: 确保治理矩阵中的目标文件全部进入审计结果记录

- [x] Task 7: 冻结 `phase11-06` 的成功标准、DoD 与收口口径
  - [x] SubTask 7.1: 将"何时算完成、何时不得判定完成"写成可执行口径
  - [x] SubTask 7.2: 保证后续执行者无需再猜"是不是只改几个显眼入口就够了"
  - [x] SubTask 7.3: 固定"目标文件必须全部逐项审计"的硬性完成条件

- [x] Task 8: 校验根级入口治理的实际结果
  - [x] SubTask 8.1: 校验根级入口不再互相复制主结论（README.md / docs/README.md / docs/phase/README.md 已改为摘要式引用）
  - [x] SubTask 8.2: 校验治理矩阵目标文件范围内对 `PSCO-summarize-feedback.md` 的悬空引用已清零（9 个目标文件中无悬空引用）
  - [x] SubTask 8.3: 校验新接手 agent 从根级入口读到的上下文已单值一致

- [x] Task 9: 将 `phase11-06` 的实现验收对象显式落到目标根级文件
  - [x] SubTask 9.1: 确认验收不只检查 spec 包自身，而是检查目标根级文件是否已实际完成治理
  - [x] SubTask 9.2: 确认治理矩阵中的目标文件已全部完成逐项审计
  - [x] SubTask 9.3: 确认“当前有效最终共识入口单值化 + 目标文件范围内悬空引用清零 + 单值一致”已在目标文件实际状态中成立，且未把当前版本号文件名上升为未来通用合同

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
- Task 4 depends on Task 2
- Task 5 depends on Task 2
- Task 6 depends on Task 2
- Task 7 depends on Task 3
- Task 7 depends on Task 4
- Task 7 depends on Task 5
- Task 7 depends on Task 6
- Task 8 depends on Task 3
- Task 8 depends on Task 4
- Task 8 depends on Task 5
- Task 8 depends on Task 6
- Task 9 depends on Task 8
