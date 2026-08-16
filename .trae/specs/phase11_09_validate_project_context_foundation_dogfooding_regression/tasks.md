# Tasks

- [x] Task 1: 盘点 `phase11-09` 的直接上游、验收输入与固定样本
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L241-L278`，冻结 `phase11-09` 的范围、验收协议与 DoD
  - [x] SubTask 1.2: 审阅 `phase11-06`、`phase11-07`、`phase11-08` 的正式输入与当前实现状态，确认本轮只承接联调、dogfooding 与反回归验证
  - [x] SubTask 1.3: 冻结 `PSCO` 仓库自身作为第一 dogfooding 样本，并记录“样本 != 模板”的前置说明

- [x] Task 2: 冻结最小工具链验证矩阵
  - [x] SubTask 2.1: 明确 `proto/` 下的正式生成命令、执行顺序与通过标准
  - [x] SubTask 2.2: 明确 `backend/` 下的正式测试命令、执行顺序与通过标准
  - [x] SubTask 2.3: 明确 `backend/` 下的正式构建命令、执行顺序与通过标准
  - [x] SubTask 2.4: 冻结命令失败、环境异常与局部通过的归类口径，不允许临场解释

- [x] Task 3: 冻结双路径 dogfooding 剧本与固定提问集合
  - [x] SubTask 3.1: 固定旧路径基线的 `6` 个入口与禁止额外读取的约束
  - [x] SubTask 3.2: 固定新路径目标的 `3` 个入口与禁止增加第 `4` 个入口的约束
  - [x] SubTask 3.3: 冻结 `5` 个恢复问题、回答格式与达标口径
  - [x] SubTask 3.4: 冻结旧路径 / 新路径共同使用的记录模板：入口清单、回答结果、失败点与是否达标

- [x] Task 4: 冻结根级真相源治理的反回归检查矩阵
  - [x] SubTask 4.1: 明确需要逐项复核的根级入口与 docs 入口范围
  - [x] SubTask 4.2: 冻结重复 phase 状态、重复目录落点、重复技术栈正文与悬空引用回流的判定口径
  - [x] SubTask 4.3: 冻结“根级治理已被真实验证”的留档格式，保证后续验收者可复核

- [x] Task 5: 冻结边界证据与样本/合同分离口径
  - [x] SubTask 5.1: 明确本阶段不做 MCP / CLI / agent 写回 / 前端对话入口的证据要求
  - [x] SubTask 5.2: 明确 `PSCO` 自身治理样本与跨项目通用能力合同分离的验收说明
  - [x] SubTask 5.3: 冻结“dogfooding 结果不能上升为未来所有项目模板”的结论表达

- [x] Task 6: 将 `phase11-09` 的冻结结果回写到 spec 包并完成自检
  - [x] SubTask 6.1: 复核 `spec.md` 是否已提供单值命令、单值入口、单值问题与单值达标条件
  - [x] SubTask 6.2: 复核 `tasks.md` 是否足以让后续独立执行者不再补造 dogfooding 主路径
  - [x] SubTask 6.3: 复核 `checklist.md` 是否覆盖工具链验证、双路径 dogfooding、根级反回归与边界证据
  - [x] SubTask 6.4: 复核 `acceptance_report.md` 是否与 `spec.md`、`tasks.md` 的冻结口径保持一致

- [x] Task 7: 补齐 `PSCO` 自身仓库样本的 `repository_id binding` 证据并重跑最终 dogfooding 验收
  - [x] SubTask 7.1: 提供 `PSCO` 自身仓库已登记并完成 `product_repositories + module_repositories` 绑定的正式可复验证据
  - [x] SubTask 7.2: 基于同一 `repository_id` 重新执行新路径固定 `3` 入口 dogfooding
  - [x] SubTask 7.3: 重新记录固定 `5` 问的回答结果、失败点与是否达标
  - [x] SubTask 7.4: 回写 `acceptance_report.md` 的最终判定，将 `phase11-09` 从“不通过”更新为“通过”或继续留在未通过状态

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 2`
- `Task 6` depends on `Task 3`
- `Task 6` depends on `Task 4`
- `Task 6` depends on `Task 5`
- `Task 7` depends on `Task 6`
