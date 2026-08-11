# Tasks

- [x] Task 1: 对齐 `phase08-03` 的直接上游、writeback matrix 与真实 owner 事实
  - [x] SubTask 1.1: 对齐 `phase08-01 / phase08-02` 已冻结的边界、页面承接位与 review 入口主线
  - [x] SubTask 1.2: 消费 `shared_baseline` 已冻结的 `Review Result Writeback Matrix`
  - [x] SubTask 1.3: 对齐当前真实 frontend application owner 与 backend command owner

- [x] Task 2: 冻结 `Feedback -> Decision` 的正式动作主线
  - [x] SubTask 2.1: 明确反馈信号、代表性缺口与 pending decision 的正式输出允许 `Decision / entity handoff / next-step result` 三类结果并存
  - [x] SubTask 2.2: 明确已存在决策上下文时优先进入既有 `Decision Center` canonical 路径
  - [x] SubTask 2.3: 明确无既有决策上下文时，只有在需要形成新的正式经营判断时才必须先创建 `decision draft`

- [x] Task 3: 冻结 `Decision / Module / Product / Repository` 的 writeback 矩阵
  - [x] SubTask 3.1: 明确 `Decision` 的允许动作与禁止动作
  - [x] SubTask 3.2: 明确 `Module` 的正式回流方式、允许动作与禁止动作
  - [x] SubTask 3.3: 明确 `Product` 的最小回流为 canonical action handoff，直写仅可复用既有 owner
  - [x] SubTask 3.4: 明确 `Repository` 的最小回流为 canonical action handoff，直写仅可复用既有 owner
  - [x] SubTask 3.5: 明确最小成功闭环为 `Decision` 正式承接 + 至少一种实体回流落地

- [x] Task 4: 冻结前端 `Review action application owner` 的单值边界
  - [x] SubTask 4.1: 明确 review 写路径必须收敛到单一 `Review action application owner`
  - [x] SubTask 4.2: 明确该 owner 只编排既有 canonical application owner 与 canonical action handoff
  - [x] SubTask 4.3: 明确 route / page / panel / card 不能直接成为新的 mutation owner
  - [x] SubTask 4.4: 明确 `FeedbackSignalCard / RecentActivityItemCard` 保持只读 caller 身份

- [x] Task 5: 冻结后端 command owner 与错误归一化边界
  - [x] SubTask 5.1: 明确 `Decision` 写入必须复用既有 `Decision Center` command owner
  - [x] SubTask 5.2: 明确 `Product / Repository / Module` 相关写入不得创建 review-local 并列 command 主线
  - [x] SubTask 5.3: 明确前端错误归一化收敛在 review action owner 边界
  - [x] SubTask 5.4: 明确后端继续保持既有 domain error → proto / Connect error 单值映射

- [x] Task 6: 完成 `phase08-03` 规格自检与一致性校验
  - [x] SubTask 6.1: 校验 `spec.md` 已覆盖动作边界、owner 边界、允许动作 / 禁止动作与错误边界
  - [x] SubTask 6.2: 校验规格与 `phase08` 三件套、`phase08-01 / 02` 规格及当前源码事实一致
  - [x] SubTask 6.3: 校验本任务未越权冻结 `phase08-04+` 的合同命名、读模型结构或实现细节

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 3
- Task 5 depends on Task 1, Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
