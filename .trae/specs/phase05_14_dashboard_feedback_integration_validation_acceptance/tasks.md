# Tasks

- [x] Task 1: 建立 `phase05-14` 联调验收环境与前置条件记录。
  - [x] SubTask 1.1: 复用 `reset_dashboard_acceptance.sh` 与正式 fixture，明确空状态、有数据状态、局部错误状态所使用的基线入口。
  - [x] SubTask 1.2: 核对前端 `/dashboard`、后端 `/api/dashboard/*`、真实 API 模式与 `.proto -> HTTP -> 前端` 单值链路。
  - [x] SubTask 1.3: 冻结验收环境记录格式，明确后续 `acceptance_report.md` 必须承接的环境与前置条件章节。

- [x] Task 2: 执行 Dashboard 状态矩阵的最小联调验收。
  - [x] SubTask 2.1: 验收 `empty-system` 空状态，确认全零概览、空态文案与空状态主 CTA。
  - [x] SubTask 2.2: 验收默认恢复或 `recent-activities` 基线，确认非零概览、反馈信号/中性态与最近活动展示。
  - [x] SubTask 2.3: 验收局部错误状态，确认 overview 不被附属聚合失败拖垮，局部错误与局部重试语义正确。
  - [x] SubTask 2.4: 验收 `Recent Activity` 的活动类型覆盖、排序、空态与错误态。
  - [x] SubTask 2.5: 若状态矩阵联调暴露实现或 fixture 缺口，必须先在当前阶段修复并复测通过后再继续验收。

- [x] Task 3: 执行三类聚合读取与前端消费闭环验收。
  - [x] SubTask 3.1: 验收 `GET /api/dashboard/overview` 的 response envelope、零值/非零值成功态与前端消费结果。
  - [x] SubTask 3.2: 验收 `GET /api/dashboard/feedback-signals` 的 response envelope、Current Focus / Asset Feedback 共享消费与 CTA 命中结果。
  - [x] SubTask 3.3: 验收 `GET /api/dashboard/recent-activities` 的 response envelope、活动列表展示与空数组成功语义。

- [x] Task 4: 执行 Dashboard 到四类 canonical owner 的跳转与返回路径验收。
  - [x] SubTask 4.1: 覆盖 `Dashboard -> Module / Product / Repository / Decision` 的直接跳转与主动返回。
  - [x] SubTask 4.2: 覆盖 `Dashboard -> List -> Detail -> List -> Dashboard` 多跳返回链，确认 `fromList` 与 `fromDashboard` 并存恢复。
  - [x] SubTask 4.3: 验收 `dashboardSection` 的一次性恢复机制，确认回到 `/dashboard` 后来源区块上下文正确。
  - [x] SubTask 4.4: 若跳转返回联调暴露“返回 Dashboard 后上下文未真正恢复”等缺口，必须先修复并复测通过。

- [x] Task 5: 产出单一验收记录并收口联调问题。
  - [x] SubTask 5.1: 编写 `acceptance_report.md`，记录环境、步骤、结果、问题与 DoD 达成情况。
  - [x] SubTask 5.2: 显式记录联调中发现的问题、级别、修复状态与复测结论。
  - [x] SubTask 5.3: 若存在未收口阻断，回写任务并中止通过结论；若全部收口，给出阶段通过结论。

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4
