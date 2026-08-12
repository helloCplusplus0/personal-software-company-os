# Tasks

- [x] Task 1: 对齐 `phase09-03` 的上游边界与最小提示集合
  - [x] SubTask 1.1: 对齐 `phase09-01`、`phase09-02` 与 `phase09` 三件套中关于派生提示的共同边界
  - [x] SubTask 1.2: 冻结当前阶段正式提示只包含 `reuse_opportunity_hint` 与 `capability_gap_hint`
  - [x] SubTask 1.3: 明确本任务不越权冻结 `phase09-04` 的合同字段、caller 清单或 owner inventory

- [x] Task 2: 冻结 `reuse_opportunity_hint` 的事实来源、解释口径与动作承接
  - [x] SubTask 2.1: 明确其事实来源为模板候选聚合结果 + `module_reuse_summary`
  - [x] SubTask 2.2: 明确其触发条件、成功空态、解释文案与模板预填创建 CTA
  - [x] SubTask 2.3: 明确其 target owner 指向 `TemplateReuseRead + ProductCreate` canonical path

- [x] Task 3: 冻结 `capability_gap_hint` 的事实来源、依赖关系与动作承接
  - [x] SubTask 3.1: 明确其事实来源为 `capability_summary + active template candidate`
  - [x] SubTask 3.2: 明确 `Weekly Review` 无 active template candidate 时返回成功空态，不回退到 generic focus；`Product Create` 则改用 `templateCandidateId` 对应的 selected template candidate
  - [x] SubTask 3.3: 明确其解释文案、`Module Registry / Module Detail` CTA 与 target owner

- [x] Task 4: 冻结提示的四元组矩阵与页面职责边界
  - [x] SubTask 4.1: 明确所有正式提示都必须满足 `trigger -> explanation -> CTA -> target owner`
  - [x] SubTask 4.2: 明确没有稳定 canonical CTA 的提示不得进入 `phase09`
  - [x] SubTask 4.3: 明确 `Weekly Review` 与 `Product Create` 的提示职责边界，避免长出第二套提示主线

- [x] Task 5: 冻结提示与既有 canonical 动作链的对接方式
  - [x] SubTask 5.1: 明确 `reuse_opportunity_hint` 只导向模板 handoff 与 `Product Create` canonical 路径
  - [x] SubTask 5.2: 明确 `capability_gap_hint` 只导向 `Module Registry / Module Detail` canonical 路径
  - [x] SubTask 5.3: 明确提示不得在当前阶段内联第二套写动作或独立任务系统

- [x] Task 6: 完成 `phase09-03` 规格自检与一致性校验
  - [x] SubTask 6.1: 校验 `spec.md` 已覆盖提示类型、触发条件、解释文案、CTA、空态与页面职责边界
  - [x] SubTask 6.2: 校验任务拆分与 `phase09` 三件套、`phase09-01 / 02` 上游边界一致
  - [x] SubTask 6.3: 校验本任务未越权冻结合同字段、caller 清单或读写 owner inventory

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 2, Task 3
- Task 5 depends on Task 2, Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
