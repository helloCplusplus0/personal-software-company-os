# Tasks

- [x] Task 1: 盘点 `phase11` 三件套中与 `phase11-03` 直接相关的四实体语义表达
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L60-L71` 中当前 `phase11-03` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase11_project_context_foundation_architecture_plan.md` 中关于四实体语义冻结与非结构重构边界的正式表达
  - [x] SubTask 1.3: 审阅 `phase11_project_context_foundation_shared_baseline.md` 中四实体语义矩阵与补充冻结口径

- [x] Task 2: 冻结 `Product / Repository / Module / Decision` 的正式语义说明
  - [x] SubTask 2.1: 固定 `Product` 作为经营目标与交付容器
  - [x] SubTask 2.2: 固定 `Repository` 作为代码仓库身份对象与项目锚点
  - [x] SubTask 2.3: 固定 `Module` 与 `Decision` 的正式语义基础定义

- [x] Task 3: 冻结当前阶段只做语义确认、不做结构重构的边界
  - [x] SubTask 3.1: 明确 `phase11-03` 只承接语义澄清，不承接 schema 重构
  - [x] SubTask 3.2: 明确不承接实体拆并、关系主线重写或第二套实体体系引入
  - [x] SubTask 3.3: 明确后续 `/spec` 与实现不得把语义确认偷渡为结构改造

- [x] Task 4: 冻结 `Module` 与 `Decision` 的当前阶段解释
  - [x] SubTask 4.1: 明确 `Module` 当前作为可复用能力资产，允许后置提炼但不在本子任务内重写 schema
  - [x] SubTask 4.2: 明确 `Decision` 当前作为规则、约束、选择与依据的索引对象
  - [x] SubTask 4.3: 明确 `Decision` 当前不扩写为审批流、流程引擎或新的结构重构入口

- [x] Task 5: 冻结 `phase11-03` 的成功标准、DoD 与收口口径
  - [x] SubTask 5.1: 将"何时算完成、何时不得判定完成"写成单值口径
  - [x] SubTask 5.2: 保证后续执行者无需再猜"四实体语义确认做到什么程度才算完成"
  - [x] SubTask 5.3: 固定超出范围的结构重构类讨论必须后移

- [x] Task 6: 完成三件套一致性校验
  - [x] SubTask 6.1: 校验 `architecture_plan`、`dev_plan`、`shared_baseline` 对四实体语义的表述单值一致
  - [x] SubTask 6.2: 校验 `Module / Decision` 的阶段解释在三件套中单值一致
  - [x] SubTask 6.3: 校验后续执行者不会再把语义确认误解为 schema 重构

- [x] Task 7: 将 `phase11-03` 的冻结结果显式回写到目标源文档
  - [x] SubTask 7.1: 在 `phase11_project_context_foundation_dev_plan.md` 中冻结四实体语义确认范围与 DoD 口径
  - [x] SubTask 7.2: 在 `phase11_project_context_foundation_architecture_plan.md` 中补齐 `Module / Decision` 的当前阶段解释，并冻结四实体语义说明与非结构重构边界
  - [x] SubTask 7.3: 在 `phase11_project_context_foundation_shared_baseline.md` 中补齐 `Module / Decision` 的当前阶段解释，并冻结四实体语义矩阵与补充冻结口径

- [x] Task 8: 校验本 spec 包的验收对象已经落到 `phase11` 三件套实际回写结果
  - [x] SubTask 8.1: 确认验收不只检查 spec 包自身，而是检查三件套是否已完成正式冻结
  - [x] SubTask 8.2: 确认四实体正式语义已在目标源文档中可直接引用
  - [x] SubTask 8.3: 确认"只做语义确认、不做 schema 重构"的边界与收口口径已在目标源文档中可直接引用

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4
- Task 6 depends on Task 2
- Task 6 depends on Task 3
- Task 6 depends on Task 4
- Task 7 depends on Task 5
- Task 8 depends on Task 6
- Task 8 depends on Task 7
