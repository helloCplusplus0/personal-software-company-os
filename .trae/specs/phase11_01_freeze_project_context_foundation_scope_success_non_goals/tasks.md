# Tasks

- [x] Task 1: 盘点 `phase11` 三件套中与 `phase11-01` 直接相关的边界、成功标准与非目标表达
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L30-L43` 中当前 `phase11-01` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase11_project_context_foundation_architecture_plan.md` 中主交付、非目标与阶段完成条件表达
  - [x] SubTask 1.3: 审阅 `phase11_project_context_foundation_shared_baseline.md` 中单值定义、非目标与验收前提表达

- [x] Task 2: 冻结 `Project Context Foundation` 的唯一主交付边界
  - [x] SubTask 2.1: 将"根级上下文真相源治理 + 最小只读项目上下文导出"冻结为唯一主交付主体
  - [x] SubTask 2.2: 明确 `AGENTS` 风格导出与 PSCO 自身 dogfooding 作为主交付内的必备组成，而非平行新主线
  - [x] SubTask 2.3: 明确当前仓库样本与未来项目通用能力之间的边界，不允许把当前根级文件清单外推为统一模板

- [x] Task 3: 冻结 `phase11` 的成功标准、DoD 与阶段收口口径
  - [x] SubTask 3.1: 提炼 `phase11-01` 对后续 `phase11-04 ~ 10` 的前置约束
  - [x] SubTask 3.2: 将"何时算完成、何时不允许继续扩范围"写成单值口径
  - [x] SubTask 3.3: 保证后续执行者进入 `/spec` 时不需要再猜"本阶段到底做什么"

- [x] Task 4: 冻结 `phase11` 明确不做清单
  - [x] SubTask 4.1: 固定排除 `MCP / CLI / agent 写回 / 前端对话式入口`
  - [x] SubTask 4.2: 固定排除四实体结构重构、知识图谱、重型 GitHub / Gitea 集成、主动注入
  - [x] SubTask 4.3: 明确这些能力只能作为未来阶段候选增强，不得以"顺手增强"形式混入当前阶段

- [x] Task 5: 完成三件套一致性校验
  - [x] SubTask 5.1: 校验 `architecture_plan`、`dev_plan`、`shared_baseline` 对主交付、非目标与成功标准的表述单值一致
  - [x] SubTask 5.2: 校验没有把未来能力写成当前事实，也没有把当前样本写成未来所有项目模板
  - [x] SubTask 5.3: 校验 `phase11-01` 已足以作为后续 `phase11-04 ~ 10` 的正式边界前提

- [x] Task 6: 将 `phase11-01` 的冻结结果显式回写到目标源文档
  - [x] SubTask 6.1: 在 `phase11_project_context_foundation_dev_plan.md` 中冻结唯一主交付、明确不做与 DoD 口径
  - [x] SubTask 6.2: 在 `phase11_project_context_foundation_architecture_plan.md` 中冻结主交付边界与阶段完成判定口径
  - [x] SubTask 6.3: 在 `phase11_project_context_foundation_shared_baseline.md` 中冻结单值定义、非目标与验收前提口径

- [x] Task 7: 校验本 spec 包的验收对象已经落到 `phase11` 三件套实际回写结果
  - [x] SubTask 7.1: 确认验收不只检查 spec 包自身，而是检查三件套是否已完成正式冻结
  - [x] SubTask 7.2: 确认 `Project Context Foundation` 的组成口径在三件套与 spec 中一致
  - [x] SubTask 7.3: 确认成功标准、DoD 与收口口径已在目标源文档中可直接引用

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
