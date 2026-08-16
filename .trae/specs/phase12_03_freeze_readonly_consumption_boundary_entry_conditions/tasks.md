# Tasks

- [x] Task 1: 盘点 `phase12-03` 直接上游中的只读消费深化口径
  - [x] SubTask 1.1: 审阅 `phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md#L60-L75` 中 `phase12-03` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `architecture_plan` 中只读消费深化边界、共享只读 owner 与 repository-scoped 承接矩阵
  - [x] SubTask 1.3: 审阅 `shared_baseline` 中只读消费深化矩阵、owner 分层与阶段完成条件

- [x] Task 2: 冻结只读消费深化的正式边界
  - [x] SubTask 2.1: 明确 `phase12` 只读消费深化继续复用 `phase11` 的正式只读主线
  - [x] SubTask 2.2: 明确当前阶段只允许深化共享只读语义摘要、规则入口、文档入口与最小解释结果
  - [x] SubTask 2.3: 明确当前阶段不承接写回、审批流、MCP / CLI 或第二套业务协议出口

- [x] Task 3: 冻结共享只读 owner 分层
  - [x] SubTask 3.1: 固定 `GetProjectContext` 为结构化 canonical owner
  - [x] SubTask 3.2: 固定 `ExportProjectContext` 为 agent-facing Markdown owner
  - [x] SubTask 3.3: 固定 `frontend/src/features/project-context/` 为唯一允许的新 Web 跨切片共享只读 owner
  - [x] SubTask 3.4: 固定各 feature 的 `pages/`、`components/` 与 `data/` 只承接切片内展示映射
  - [x] SubTask 3.5: 冻结 `project-context/` 的启用条件（`3+` 页面/切片稳定复用）与禁止事项（不得承接写路径、页面私有状态、并列 canonical 字段语义）

- [x] Task 4: 冻结 `repository_id` 与三类页面承接矩阵
  - [x] SubTask 4.1: 固定 `repositories/$repositoryId` 为直接 `repository-scoped` 页面
  - [x] SubTask 4.2: 固定 `products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 为间接 `repository-scoped` 页面
  - [x] SubTask 4.3: 固定 `dashboard`、`onboarding`、`reviews/daily`、`reviews/weekly` 为衍生消费页
  - [x] SubTask 4.4: 明确三类页面各自允许的锚点承接方式与禁止事项

- [x] Task 5: 冻结复用与最小新增的判定边界
  - [x] SubTask 5.1: 明确何时必须继续复用 `GetProjectContext`
  - [x] SubTask 5.2: 明确何时只允许新增 `ProjectContextService` 下的受控派生读取
  - [x] SubTask 5.3: 明确任何新增承接位都必须保留定位关系并写清回收的重复解释逻辑

- [x] Task 6: 冻结更重通道与受控维护能力进入条件
  - [x] SubTask 6.1: 明确 `MCP / CLI / agent 写回 / Draft / 审批流 / 前端对话式入口 / 第二套 canonical API / 影子状态表` 不属于当前阶段正式范围
  - [x] SubTask 6.2: 明确讨论更重能力前必须先满足的阶段完成前提
  - [x] SubTask 6.3: 明确额外专家讨论只允许围绕已压缩候选方向展开
  - [x] SubTask 6.4: 明确更重通道进入条件必须完整继承 phase-wide 验收门槛（固定 `6` 问闭环、固定 `repository_id` 锚点闭环、样本解析协议闭环）

- [x] Task 7: 完成三件套一致性校验
  - [x] SubTask 7.1: 校验 `dev_plan`、`architecture_plan` 与 `shared_baseline` 对只读消费深化边界的表达单值一致
  - [x] SubTask 7.2: 校验三件套对 owner 分层、三类页面承接矩阵与锚点规则的表达单值一致
  - [x] SubTask 7.3: 校验三件套对更重通道进入条件与当前非目标的表达单值一致

- [x] Task 8: 校验本 spec 包可直接作为 `phase12-05 ~ 07` 的输入前提
  - [x] SubTask 8.1: 确认 spec 已回答“这里该新建前端共享读模型，还是直接拼页面解释”
  - [x] SubTask 8.2: 确认 spec 已回答三类页面如何承接 `repository_id`
  - [x] SubTask 8.3: 确认 spec 已回答哪些能力必须留到后续阶段处理

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 1
- Task 6 depends on Task 1
- Task 7 depends on Task 2
- Task 7 depends on Task 3
- Task 7 depends on Task 4
- Task 7 depends on Task 5
- Task 7 depends on Task 6
- Task 8 depends on Task 7
