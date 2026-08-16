# Tasks

- [x] Task 1: 盘点 `phase12` 三件套中与 `phase12-01` 直接相关的范围边界、成功标准、共享只读 owner 与非目标表达
  - [x] SubTask 1.1: 审阅 `phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md#L30-L42` 中当前 `phase12-01` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md` 中主交付、共享只读 owner、完成条件与非目标表达
  - [x] SubTask 1.3: 审阅 `phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md` 中单值定义、验收前提、固定样本与非目标表达

- [x] Task 2: 冻结 `Semantic Alignment & Read-Only Consumption Foundation` 的唯一主交付边界
  - [x] SubTask 2.1: 将 `四实体语义一致性收口 + 只读消费深化` 冻结为 `phase12` 的唯一主交付组成
  - [x] SubTask 2.2: 明确语义一致性只承接表达层与消费层对齐，不承接 schema 或关系主线重构
  - [x] SubTask 2.3: 明确只读消费深化只承接 `phase11` 已交付 `Project Context` 主线上的共享只读能力，不承接更重 agent 通道

- [x] Task 3: 冻结 `phase12` 的成功标准、DoD 与阶段收口口径
  - [x] SubTask 3.1: 提炼 `phase12-04 ~ 12` 对范围边界、样本协议、共享只读 owner 与固定入口的前置依赖
  - [x] SubTask 3.2: 将“何时算完成、何时不允许继续扩范围”写成单值口径
  - [x] SubTask 3.3: 保证后续执行者进入 `/spec` 时不需要再猜“本阶段到底做什么、做到什么程度算完成”

- [x] Task 4: 冻结 `phase12` 明确不做清单
  - [x] SubTask 4.1: 固定排除四实体 schema 重写、关系主线重构与第二套 canonical API
  - [x] SubTask 4.2: 固定排除 `MCP / CLI / agent 写回 / Draft / 审批流 / 前端对话式 agent 入口`
  - [x] SubTask 4.3: 明确这些能力只能作为后续阶段候选增强，不得以“顺手增强”形式混入当前阶段

- [x] Task 5: 冻结共享只读 owner 与验收协议前提
  - [x] SubTask 5.1: 固定 `GetProjectContext / ExportProjectContext / frontend/src/features/project-context/` 的 owner 层级关系
  - [x] SubTask 5.2: 固定 `repository_id` 作为唯一结构化输入锚点
  - [x] SubTask 5.3: 固定 `product_id / module_id / decision_id` 只能从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析
  - [x] SubTask 5.4: 固定样本解析失败、结果不唯一或无法回到同一 `repository_id` 时必须直接判定验收失败

- [x] Task 6: 完成三件套一致性校验
  - [x] SubTask 6.1: 校验 `architecture_plan`、`dev_plan`、`shared_baseline` 对唯一主交付、成功标准与非目标的表述单值一致
  - [x] SubTask 6.2: 校验三件套对共享只读 owner、固定样本、固定入口与固定 `6` 问的表述单值一致
  - [x] SubTask 6.3: 校验没有把未来更重能力写成当前事实，也没有允许验收通过额外入口补答案

- [x] Task 7: 将 `phase12-01` 的冻结结果显式回写到目标源文档
  - [x] SubTask 7.1: 在 `phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md` 中冻结唯一主交付、成功标准与非目标口径
  - [x] SubTask 7.2: 在 `phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md` 中冻结主交付边界、设计顺序、共享只读 owner 与阶段完成条件
  - [x] SubTask 7.3: 在 `phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md` 中冻结单值定义、验收前提、固定样本解析协议与非目标口径

- [x] Task 8: 校验本 spec 包的验收对象已经落到 `phase12` 三件套实际回写结果
  - [x] SubTask 8.1: 确认验收不只检查 spec 包自身，而是检查三件套是否已完成正式冻结
  - [x] SubTask 8.2: 确认唯一主交付能力、共享只读 owner 与固定样本解析协议在 spec 与三件套中一致
  - [x] SubTask 8.3: 确认成功标准、DoD 与收口口径已在目标源文档中可直接引用

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 1
- Task 6 depends on Task 2
- Task 6 depends on Task 3
- Task 6 depends on Task 4
- Task 6 depends on Task 5
- Task 7 depends on Task 2
- Task 7 depends on Task 3
- Task 7 depends on Task 4
- Task 7 depends on Task 5
- Task 8 depends on Task 6
- Task 8 depends on Task 7
