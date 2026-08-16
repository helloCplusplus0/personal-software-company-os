# Tasks

- [x] Task 1: 盘点 `phase11` 三件套中与 `phase11-04` 直接相关的根级治理表达
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L82-L103` 中当前 `phase11-04` 的范围、正式产物与 DoD
  - [x] SubTask 1.2: 审阅 `phase11_project_context_foundation_architecture_plan.md` 中关于根级真相源治理边界、单一写者规则与最终共识入口的正式表达
  - [x] SubTask 1.3: 审阅 `phase11_project_context_foundation_shared_baseline.md` 中根级真相源治理矩阵、单一写者规则与非目标口径

- [x] Task 2: 冻结根级治理对象与单一写者规则
  - [x] SubTask 2.1: 固定 `README.md / plan.md / AGENTS.md / architecture_map.md / docs/README.md / docs/phase/README.md / project_rules.md / global_skills.md` 的治理对象范围
  - [x] SubTask 2.2: 固定"谁是单一写者、谁只保留摘要式引用"的正式规则
  - [x] SubTask 2.3: 固定这一治理清单只针对 `PSCO` 自身仓库，不外推为未来所有项目的固定目录模板

- [x] Task 3: 冻结根级入口治理矩阵与清理清单设计
  - [x] SubTask 3.1: 明确根级入口治理矩阵的正式组成项
  - [x] SubTask 3.2: 明确重复承载清单与目标落点清单的正式组成项
  - [x] SubTask 3.3: 明确悬空引用清理清单与收口后的单一写者规则表的正式组成项

- [x] Task 4: 冻结最终共识入口统一改写策略与不允许的重复表达模式
  - [x] SubTask 4.1: 明确 `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识入口的统一改写策略
  - [x] SubTask 4.2: 明确不再允许出现的重复表达模式
  - [x] SubTask 4.3: 明确哪些入口文档只能保留摘要式引用或受控跳转

- [x] Task 5: 冻结当前阶段治理路线与非目标边界
  - [x] SubTask 5.1: 明确当前阶段承接的是根级治理设计与一次性校准策略
  - [x] SubTask 5.2: 明确当前阶段不承接静态文件全量 backend 派生
  - [x] SubTask 5.3: 保证后续执行者不需要继续开放争论"是全量派生还是一次性校准"

- [x] Task 6: 冻结 `phase11-04` 的成功标准、DoD 与收口口径
  - [x] SubTask 6.1: 将"何时算完成、何时不得判定完成"写成单值口径
  - [x] SubTask 6.2: 保证后续执行者无需再猜"治理设计做到什么程度才算完成"
  - [x] SubTask 6.3: 固定超出范围的治理路线与未来模板外推必须后移

- [x] Task 7: 完成三件套一致性校验
  - [x] SubTask 7.1: 校验 `architecture_plan`、`dev_plan`、`shared_baseline` 对治理对象、单一写者规则与统一改写策略的表述单值一致
  - [x] SubTask 7.2: 校验根级入口治理矩阵、清理清单与非目标口径在三件套中单值一致
  - [x] SubTask 7.3: 校验后续执行者不会再把治理设计误解为静态文件全量派生方案

- [x] Task 8: 将 `phase11-04` 的冻结结果显式回写到目标源文档
  - [x] SubTask 8.1: 在 `phase11_project_context_foundation_dev_plan.md` 中冻结根级治理设计范围、正式产物与 DoD 口径
  - [x] SubTask 8.2: 在 `phase11_project_context_foundation_architecture_plan.md` 中补齐 `project_rules.md / docs/phase/README.md` 的治理定位，并冻结正式设计产物、统一改写策略与治理路线边界
  - [x] SubTask 8.3: 在 `phase11_project_context_foundation_shared_baseline.md` 中补齐 `project_rules.md / docs/phase/README.md` 的治理矩阵定位，并冻结清理清单口径、单一写者规则与一次性校准路线

- [x] Task 9: 校验本 spec 包的验收对象已经落到 `phase11` 三件套实际回写结果
  - [x] SubTask 9.1: 确认验收不只检查 spec 包自身，而是检查三件套是否已完成正式冻结
  - [x] SubTask 9.2: 确认根级治理对象、统一改写策略与单一写者规则已在目标源文档中可直接引用
  - [x] SubTask 9.3: 确认"只做治理设计与一次性校准、不做全量派生"的边界与收口口径已在目标源文档中可直接引用

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
- Task 9 depends on Task 7
- Task 9 depends on Task 8
