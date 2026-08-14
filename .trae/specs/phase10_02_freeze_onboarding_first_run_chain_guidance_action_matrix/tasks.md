# Tasks

- [x] Task 1: 对齐 `phase10-02` 的直接上游与 `Onboarding` 新语义
  - [x] SubTask 1.1: 对齐 `phase10-01`、`phase10` 三件套与 `phase06` onboarding 既有规格的共同边界
  - [x] SubTask 1.2: 明确 `Onboarding` 在 `phase10` 中从“首轮登记入口”升级为“首轮建链引导主线”
  - [x] SubTask 1.3: 明确当前子任务不新增第二套草稿系统、不把 `Onboarding` 演化为工作流引擎

- [x] Task 2: 冻结六段式主线与逐步建链矩阵
  - [x] SubTask 2.1: 冻结 `welcome / product / repository / module / decision / complete` 六段式主线
  - [x] SubTask 2.2: 冻结每一步的 canonical owner、最小动作与默认下一步
  - [x] SubTask 2.3: 冻结哪些关系必须在当前步骤直接承接，哪些关系允许延后到 canonical handoff 或 detail CTA
  - [x] SubTask 2.4: 明确 `module` 步骤若无法同页完成最小正式关系时，必须提供单值 canonical handoff
  - [x] SubTask 2.5: 明确 `complete` 是否允许存在未即时完成但已单值解释的 canonical handoff

- [x] Task 3: 冻结 `Onboarding` 与 canonical 写路径的承接关系
  - [x] SubTask 3.1: 明确 `Product / Repository / Module / Decision` 既有 canonical owner 继续作为正式写路径
  - [x] SubTask 3.2: 明确 `Onboarding` 只承接首轮建链引导，不新增并列写路径
  - [x] SubTask 3.3: 明确无法即时完成的关系必须通过单值 canonical handoff 承接

- [x] Task 4: 冻结恢复语义与唯一主上下文锚点
  - [x] SubTask 4.1: 明确 `current_product_id` 是唯一主上下文锚点
  - [x] SubTask 4.2: 明确 `current_repository_id / current_module_id / current_decision_id` 只作为 step 级辅助恢复线索
  - [x] SubTask 4.3: 明确多实体并存时不得按“全局最新实体”猜测当前 onboarding 主线
  - [x] SubTask 4.4: 明确显式 step 返回线索与恢复读模型冲突时的优先级

- [x] Task 5: 完成 `phase10-02` 规格自检与一致性校验
  - [x] SubTask 5.1: 校验 `spec.md` 已覆盖 Why / What Changes / Impact / Requirements / Migration
  - [x] SubTask 5.2: 校验本规格与 `phase10` 三件套的逐步建链矩阵口径一致
  - [x] SubTask 5.3: 校验本规格未越权冻结后续页面命名、接口名或实现细节

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 2
- Task 5 depends on Task 1, Task 2, Task 3, Task 4
