# Tasks

- [x] Task 1: 盘点 `phase11` 三件套中与 `phase11-02` 直接相关的系统定位与渠道分工表达
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L45-L58` 中当前 `phase11-02` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase11_project_context_foundation_architecture_plan.md` 中关于"上下文系统 / 开发流程控制器 / web / agent"的正式表达
  - [x] SubTask 1.3: 审阅 `phase11_project_context_foundation_shared_baseline.md` 中 PSCO 单值定位与 Web / Agent 分工矩阵

- [x] Task 2: 冻结 PSCO 作为"上下文系统"的正式定位
  - [x] SubTask 2.1: 明确 PSCO 的职责是提供上下文、关系、约束、决策依据与回看入口
  - [x] SubTask 2.2: 明确 PSCO 不是 IDE 现场流程控制器、开发编排器或 agent 工作台
  - [x] SubTask 2.3: 明确 IDE / agent 负责项目内微观执行推进，PSCO 不规定下一步如何开发

- [x] Task 3: 冻结 web 与 agent 的职责分工边界
  - [x] SubTask 3.1: 固定 web 继续承接全局查看、回顾、校对、历史查阅、人工修正与最终确认
  - [x] SubTask 3.2: 固定 agent 当前只承接现场上下文消费
  - [x] SubTask 3.3: 明确 web 不退化、agent 不对称并行进入，不得各自长出第二套正式流程

- [x] Task 4: 冻结共享 Go backend canonical core 与当前阶段明确不做
  - [x] SubTask 4.1: 固定 web 与 agent 共享同一套 Go backend canonical core
  - [x] SubTask 4.2: 固定当前阶段不承接 agent 写回、审批流、前端对话式入口或 agent 专属一级业务对象
  - [x] SubTask 4.3: 明确不允许形成第二套语义、第二套流程或第二套事实源

- [x] Task 5: 完成三件套一致性校验
  - [x] SubTask 5.1: 校验 `architecture_plan`、`dev_plan`、`shared_baseline` 对系统定位的表述单值一致
  - [x] SubTask 5.2: 校验 web / agent 分工边界在三件套中单值一致
  - [x] SubTask 5.3: 校验后续执行者不会再把 PSCO 理解成 IDE 现场流程编排器

- [x] Task 6: 将 `phase11-02` 的冻结结果显式回写到目标源文档
  - [x] SubTask 6.1: 在 `phase11_project_context_foundation_dev_plan.md` 中冻结系统定位、web / agent 分工边界与 DoD 口径
  - [x] SubTask 6.2: 在 `phase11_project_context_foundation_architecture_plan.md` 中冻结 PSCO 作为上下文系统的定位、共享后端约束与当前明确不做事项
  - [x] SubTask 6.3: 在 `phase11_project_context_foundation_shared_baseline.md` 中冻结单值定位、Web / Agent 分工矩阵与共享 Go backend canonical core 口径

- [x] Task 7: 校验本 spec 包的验收对象已经落到 `phase11` 三件套实际回写结果
  - [x] SubTask 7.1: 确认验收不只检查 spec 包自身，而是检查三件套是否已完成正式冻结
  - [x] SubTask 7.2: 确认 `PSCO 是上下文系统`、`web / agent` 分工与共享后端约束已在目标源文档中可直接引用
  - [x] SubTask 7.3: 确认 `agent` 当前只做现场上下文消费、当前明确不做事项与收口口径已在目标源文档中可直接引用

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
- Task 7 depends on Task 6
