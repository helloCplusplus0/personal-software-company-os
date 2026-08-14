# Tasks

- [x] Task 1: 对齐 `phase10-03` 的直接上游与 `Decision` 新语义
  - [x] SubTask 1.1: 对齐 `phase10-01`、`phase10` 三件套与 `fix_002 / fix_003` 已收口结论的共同边界
  - [x] SubTask 1.2: 明确 `Decision.status` 在 `phase10` 中同时承接生命周期推进、pending 判定与 reread 解释
  - [x] SubTask 1.3: 明确当前子任务不新增第五态、不引入第二事实源

- [x] Task 2: 冻结 `Decision` 四态生命周期矩阵
  - [x] SubTask 2.1: 冻结 `proposed / active / superseded / archived` 四态范围
  - [x] SubTask 2.2: 冻结 `proposed` 与 `active` 的允许迁移矩阵
  - [x] SubTask 2.3: 冻结 `superseded / archived` 的终态解释

- [x] Task 3: 冻结 `Decision Detail` 在各状态下的 CTA 矩阵
  - [x] SubTask 3.1: 明确 `Decision Detail` 是唯一正式状态推进承接位
  - [x] SubTask 3.2: 冻结 `proposed` 状态下允许展示的 CTA
  - [x] SubTask 3.3: 冻结 `active` 状态下允许展示的 CTA
  - [x] SubTask 3.4: 冻结终态下只保留结果消费与返回动作
  - [x] SubTask 3.5: 冻结 `Decision Detail` 非状态型 CTA 的完整 inventory 与边界

- [x] Task 4: 冻结 `Dashboard / Daily Review / Current Focus` 与 `Decision.status` 的统一 pending 语义
  - [x] SubTask 4.1: 明确 `pending decision` 继续完全锚定 canonical `Decision.status`
  - [x] SubTask 4.2: 明确 `decision_links / review_records` 不是退出 pending 的代理条件
  - [x] SubTask 4.3: 明确跨消费面的统一解释与禁止分裂语义
  - [x] SubTask 4.4: 冻结 `proposed / active / superseded / archived` 四态在三类消费面上的完整消费语义

- [x] Task 5: 冻结退出 pending 的正式条件与 reread 规则
  - [x] SubTask 5.1: 明确只有 `Decision.status` 正式离开 `proposed` 才能退出 pending
  - [x] SubTask 5.2: 明确状态推进后的 reread 必须回答的核心问题
  - [x] SubTask 5.3: 明确禁止通过页面局部隐藏、前端临时过滤或 toast 假象掩盖未完成状态

- [x] Task 6: 完成 `phase10-03` 规格自检与一致性校验
  - [x] SubTask 6.1: 校验 `spec.md` 已覆盖 Why / What Changes / Impact / Requirements / Migration
  - [x] SubTask 6.2: 校验本规格与 `phase10` 三件套、`fix_002 / fix_003` 收口结论口径一致
  - [x] SubTask 6.3: 校验本规格未越权冻结后续接口名、实现细节或页面布局细节

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 2
- Task 5 depends on Task 2, Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
