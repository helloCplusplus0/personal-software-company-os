# Tasks

- [x] Task 1: 对齐 `phase06-11` 的直接上游与验收边界，明确这次任务是“联调验收环境、fixture 与恢复基线设计”，不是源码实现。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md` 中 `phase06-11` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase06` shared baseline / architecture plan 中关于首轮成功会话、数据主权闭合、复用感知最新状态与 `backup verified` 的冻结结论
  - [x] SubTask 1.3: 对齐 `phase06-01 / 02 / 03 / 04 / 06 / 08 / 09 / 10` 中与 fixture 直接相关的状态与合同语义
  - [x] SubTask 1.4: 对齐 `phase05-09 / 14` 已建立的验收环境与 fixture 设计模式，复用而不重造

- [x] Task 2: 冻结 `phase06` 验收环境的统一入口与目录复用策略。
  - [x] SubTask 2.1: 冻结 `database/scripts/ + database/seeds/` 作为唯一验收工具链
  - [x] SubTask 2.2: 冻结 `reset_phase06_acceptance.sh` 作为唯一统一入口
  - [x] SubTask 2.3: 冻结默认、`--clean-only`、`--restore-only`、`--fixture <name>` 四种模式
  - [x] SubTask 2.4: 冻结清空范围、默认恢复顺序与幂等要求

- [x] Task 3: 冻结 `first_run_state` 与首轮成功会话的正式 fixture。
  - [x] SubTask 3.1: 冻结 `cold-start-empty`
  - [x] SubTask 3.2: 冻结 `in-progress-partial-entry`
  - [x] SubTask 3.3: 冻结 `completed-unbound`
  - [x] SubTask 3.4: 冻结 `completed-bound`

- [x] Task 4: 冻结数据主权闭合与恢复前提的正式 fixture。
  - [x] SubTask 4.1: 冻结 `export-ready`
  - [x] SubTask 4.2: 冻结 `backup-verified`
  - [x] SubTask 4.3: 冻结 `backup-manifest-missing`
  - [x] SubTask 4.4: 冻结 `backup-coverage-incomplete`
  - [x] SubTask 4.5: 冻结 `backup-schema-mismatch`

- [x] Task 5: 冻结复用感知可见与最新状态验证 fixture。
  - [x] SubTask 5.1: 冻结 `reuse-latest`
  - [x] SubTask 5.2: 冻结 `reuse-latest-after-binding`
  - [x] SubTask 5.3: 冻结这两类 fixture 与 Dashboard / Module Detail / Product Detail 的验收挂接关系

- [x] Task 6: 冻结首轮成功会话、阶段完成与失败路径验收矩阵。
  - [x] SubTask 6.1: 冻结入口判定与回访继续矩阵
  - [x] SubTask 6.2: 冻结“缺少绑定关系仍完成首轮会话”与“绑定补全后再次验证”矩阵
  - [x] SubTask 6.3: 冻结数据主权闭合矩阵
  - [x] SubTask 6.4: 冻结 `backup verified` 失败路径矩阵
  - [x] SubTask 6.5: 冻结复用感知最新状态矩阵

- [x] Task 7: 冻结"不得依赖手工补数据"的验收门禁并完成规格一致性校验。
  - [x] SubTask 7.1: 明确所有关键状态必须通过白名单 fixture 重复建立
  - [x] SubTask 7.2: 明确不得手工修改 `first_run_state`、绑定关系、导出结果或备份 `manifest`
  - [x] SubTask 7.3: 验证本 spec 与 `phase06-01 / 03 / 04 / 06 / 08 / 09 / 10` 保持单值一致
  - [x] SubTask 7.4: 验证本 spec 足以支撑后续真实 API 联调、浏览器验收与阶段收口

- [x] Task 8: 冻结合同一致性验证、备份失败 fixture 结果语义与验收产物清理边界。
  - [x] SubTask 8.1: 冻结 OnboardingRead（含 `first_run_state`）的 `.proto -> HTTP DTO -> 前端消费模型` 合同一致性纳入验收矩阵
  - [x] SubTask 8.2: 冻结 Export / Backup / Reuse Summary 的合同一致性验证不得遗漏任一能力
  - [x] SubTask 8.3: 冻结 `backup_snapshot` 读取侧合同一致性必须由 `BackupWrite.read_verify` 或等价上游冻结承接位独立验证，不得以写入响应附带代替
  - [x] SubTask 8.4: 冻结 `--restore-only` 模式必须处理同标题 readonly seed 的情况
  - [x] SubTask 8.5: 冻结三类备份失败 fixture 的结果语义与可重复复现要求，不提前冻结单一实现方式
  - [x] SubTask 8.6: 冻结验收产物清理边界，不提前冻结数据库表、文件系统目录或其他存储介质
  - [x] SubTask 8.7: 冻结清理边界不清理 canonical 表与 `schema_migrations`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 6` depends on `Task 3`, `Task 4`, `Task 5`
- `Task 8` depends on `Task 1`, `Task 2`
- `Task 7` depends on `Task 1` through `Task 6`, and `Task 8`
