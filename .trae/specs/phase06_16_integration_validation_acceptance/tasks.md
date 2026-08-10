# Tasks

- [x] Task 1: 冻结 `phase06-16` 的联调环境、前置条件与正式验收入口。
  - [x] SubTask 1.1: 明确 `reset_phase06_acceptance.sh`、默认恢复与 11 个 fixture 白名单是唯一正式验收环境入口。
  - [x] SubTask 1.2: 明确真实前端、真实后端、真实数据库与正式 `/api` 接口是唯一允许的联调对象。
  - [x] SubTask 1.3: 明确验收前必须核对的前置条件、失败中止条件与重复执行前提。

- [x] Task 2: 冻结首轮录入、回访继续与成功会话的验收矩阵。
  - [x] SubTask 2.1: 明确 `cold-start-empty` 下的根级默认进入路径、`/onboarding` 主 CTA 与冷启动空系统语义。
  - [x] SubTask 2.2: 明确 `in-progress-partial-entry` 下的根级进入 `/dashboard`、`Continue Onboarding` 与 `/onboarding` 自动定位未完成步骤语义。
  - [x] SubTask 2.3: 明确一次首轮成功会话必须验证四类对象都已真实持久化，且成功会话可重复走通。
  - [x] SubTask 2.4: 明确 `Onboarding -> canonical detail -> /onboarding` 回流链与 `fromOnboarding` 返回优先级也属于本阶段必验项。

- [x] Task 3: 冻结 `Export / Backup / backup verified` 的联调验收矩阵。
  - [x] SubTask 3.1: 明确 `export-ready` 的读取、触发与覆盖矩阵检查项，要求绑定 / 关联关系必须进入导出结果。
  - [x] SubTask 3.2: 明确 `backup-verified` 的成功判定必须同时验证产物、manifest、覆盖矩阵与 schema/version 前提。
  - [x] SubTask 3.3: 明确 `backup-manifest-missing / backup-coverage-incomplete / backup-schema-mismatch` 三类 fixture 的单值失败语义验证规则。
  - [x] SubTask 3.4: 明确 `.proto -> HTTP DTO -> 前端展示` 的失败语义一致性也要纳入本阶段验收。

- [x] Task 4: 冻结复用反馈与最新状态的验收矩阵。
  - [x] SubTask 4.1: 明确 Dashboard 的 `Reuse Snapshot` 子区域必须独立验证局部 `loading / ready / empty / error` 状态。
  - [x] SubTask 4.2: 明确 Module Detail 与 Product Detail 的复用摘要挂接位、单 query owner 约束与只读边界。
  - [x] SubTask 4.3: 明确 `reuse-latest-after-binding` 必须证明 `module_reuse_summary / capability_summary` 反映最新已提交状态。

- [x] Task 5: 冻结 `phase05` 兼容性、局部错误边界与重复执行前提的回归验收。
  - [x] SubTask 5.1: 明确 `overview / feedback-signals / recent-activities` 三个既有 Dashboard 查询仍需继续验证。
  - [x] SubTask 5.2: 明确 `phase06` 新增的 `Onboarding CTA / Export / Backup / Reuse Snapshot` 不得破坏 `phase05` 的状态模型、返回路径与局部错误边界。
  - [x] SubTask 5.3: 明确导出、备份、fixture 恢复与默认恢复都必须可重复执行，不得依赖手工清理。

- [x] Task 6: 冻结验收记录、问题收口与 DoD 判定方式。
  - [x] SubTask 6.1: 明确 `acceptance_report.md` 是本阶段唯一正式验收记录。
  - [x] SubTask 6.2: 明确报告必须记录环境、步骤、结果、问题、修复、复测与 DoD 达成情况。
  - [x] SubTask 6.3: 明确若存在未修复阻断问题，则不得宣告 `phase06-16` 通过。

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 1
- Task 6 depends on Task 2, Task 3, Task 4, Task 5
