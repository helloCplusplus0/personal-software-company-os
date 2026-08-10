# Tasks

- [x] Task 1: 冻结 `Export` 与 `Backup` 的职责边界。把“带走核心资产数据”和“保留当前实例并校验恢复前提”写成单值结论。
  - [x] SubTask 1.1: 明确 `Export` 的正式语义是面向用户带走核心资产数据
  - [x] SubTask 1.2: 明确 `Backup` 的正式语义是面向当前实例保留与恢复前提校验
  - [x] SubTask 1.3: 验证当前规格没有再把 `Export / Backup` 写成同义词

- [x] Task 2: 冻结最小覆盖矩阵。把 `Export` 与 `Backup` 必须覆盖的数据集和元信息压缩成可执行前提。
  - [x] SubTask 2.1: 明确 `Export` 最小覆盖 `products / modules / releases / repositories / decisions / decision_links / product_modules / product_repositories / module_repositories`
  - [x] SubTask 2.2: 明确 `Backup` 的最小覆盖范围不小于 `Export`
  - [x] SubTask 2.3: 明确 `Backup` 必须额外带出 `manifest / created_at / schema_version prerequisites`
  - [x] SubTask 2.4: 验证当前规格不允许“只导出主实体，不导出绑定关系”

- [x] Task 3: 冻结正式入口位与执行路由。把 `Dashboard` 动作区与 `/dashboard/export`、`/dashboard/backup` 写成单值口径。
  - [x] SubTask 3.1: 明确 `Export` 正式用户入口位于 `Dashboard` 动作区，执行路由固定为 `/dashboard/export`
  - [x] SubTask 3.2: 明确 `Backup` 正式用户入口位于 `Dashboard` 动作区，执行路由固定为 `/dashboard/backup`
  - [x] SubTask 3.3: 验证当前规格没有把完整流程内联塞回 `Dashboard Home` 主内容区

- [x] Task 4: 冻结恢复前提与 `backup verified` 最小验证链。把“可验证恢复前提成立”细化成可执行动作。
  - [x] SubTask 4.1: 明确当前阶段“恢复前提”只承接 restore prerequisites read / verify，不要求真正 restore 写回
  - [x] SubTask 4.2: 明确 `backup verified` 至少要求产物可读、manifest 可读、覆盖矩阵可核对、schema / version 前提可校验
  - [x] SubTask 4.3: 明确“仅文件写出成功”不得视为 `backup verified`

- [x] Task 5: 冻结最小错误语义与第三方依赖边界。确保后续实现和验收不会把错误与非目标混写。
  - [x] SubTask 5.1: 明确 `Export` 的最小失败归类
  - [x] SubTask 5.2: 明确 `Backup` / 恢复前提校验的最小失败归类
  - [x] SubTask 5.3: 明确当前阶段不依赖 GitHub 或其他第三方平台作为唯一前提
  - [x] SubTask 5.4: 明确自动同步、连续备份、多端同步、复杂灾备、完整 restore 写回不进入当前范围

- [x] Task 6: 完成规格一致性校验。验证本次 `phase06-03` 规格已和三件套及历史导出 / 备份基线对齐。
  - [x] SubTask 6.1: 验证本次规格与 `phase06` shared baseline、architecture plan、dev plan 保持一致
  - [x] SubTask 6.2: 验证本次规格延续 `phase01-05` 的导出 / 备份语义，但把 `backup verified` 补细到动作级
  - [x] SubTask 6.3: 验证当前规格足以支持后续 `.proto`、后端 owner、Dashboard 入口和验收 fixture 继续推进

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
