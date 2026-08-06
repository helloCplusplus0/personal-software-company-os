# Tasks

- [x] Task 1: 冻结 `Decision Center` 联调环境可重复建立方式与前置脚本入口。
  - [x] SubTask 1.1: 冻结复用 `init_db.sh` / `run_seeds.sh` / `reset_module_mainline.sh` 的入口模式
  - [x] SubTask 1.2: 冻结启动顺序（`init_db` -> 后端 migration -> `run_seeds` -> `reset_module_mainline` -> `reset_decision_mainline` -> 前端），`run_seeds` 必须在 migration 完成后执行
  - [x] SubTask 1.3: 明确不得新建第二个数据库或第二套 `init_db` 入口

- [x] Task 2: 冻结 `decisions` 表原位升级 migration 设计。
  - [x] SubTask 2.1: 冻结 migration 文件落点 `database/migrations/0004_decision_center_mainline.sql`
  - [x] SubTask 2.2: 冻结 `ALTER TABLE` 原位添加字段（`context / problem / alternatives / choice / reason / impact / status`），无 `DEFAULT` 必填字段按三步流程（`ADD COLUMN` 允许 `NULL` -> 回填 -> `SET NOT NULL`）
  - [x] SubTask 2.3: 冻结 `alternatives` 存储方式为 `TEXT[]`（`DEFAULT '{}'`）
  - [x] SubTask 2.4: 冻结 `status` 的 `CHECK` 约束与列表读取索引
  - [x] SubTask 2.5: 冻结必填字段 `NOT NULL` 约束与空字符串校验归属（service 层）
  - [x] SubTask 2.6: 明确不得新建替代表、不得删除原有字段、不得破坏 `decision_links` 外键

- [x] Task 3: 冻结现有示例 `Decision` 数据兼容回填策略。
  - [x] SubTask 3.1: 冻结回填必须保留原有 `title / created_at`
  - [x] SubTask 3.2: 冻结 `context / problem / choice / reason` 回填为占位文本
  - [x] SubTask 3.3: 冻结 `alternatives` 回填为 `'{}'`、`impact` 回填为 `''`、`status` 回填为 `'proposed'`
  - [x] SubTask 3.4: 冻结回填必须可重复执行（`WHERE` 条件幂等）
  - [x] SubTask 3.5: 冻结既有 `decision_links` 兼容性保证

- [x] Task 4: 冻结 `Decision Center` 重置脚本落点与职责。
  - [x] SubTask 4.1: 冻结脚本落点 `database/scripts/reset_decision_mainline.sh`
  - [x] SubTask 4.2: 冻结三种模式（`--clean-only` / `--restore-only` / 默认）
  - [x] SubTask 4.3: 冻结复用 `resolve_psql` 模式与环境变量参数
  - [x] SubTask 4.4: 冻结清空范围（`DELETE FROM decisions`，级联清空 `decision_links`，不影响 `modules / products / repositories`）
  - [x] SubTask 4.5: 冻结前置校验（数据库存在 + `modules` 基线存在）

- [x] Task 5: 冻结 `Decision Center` 基线种子数据范围。
  - [x] SubTask 5.1: 冻结 seed 文件落点 `database/seeds/seed_decision_mainline_baseline.sql`
  - [x] SubTask 5.2: 冻结基线 `Decision` 数据覆盖维度（`3` 条、覆盖 `proposed / active / archived|superseded`）
  - [x] SubTask 5.3: 冻结 `alternatives` 数组与空数组覆盖、`impact` 空字符串覆盖
  - [x] SubTask 5.4: 冻结保留 `phase02` 原有 `title` 以兼容 `decision_links`
  - [x] SubTask 5.5: 冻结基线 `decision_links` 数据（至少 `2` 条、通过 `name` 查找不硬编码 `UUID`、含 `1` 条无关联 `Decision`）
  - [x] SubTask 5.6: 冻结 `seed_readonly_prereqs.sql` 中 `decisions` seed 更新为结构化字段

- [x] Task 6: 冻结异常路径验证前提与要求。
  - [x] SubTask 6.1: 冻结异常路径验证清单（必填缺失 / 字段值非法含 alternatives 空白与非法 status / 目标类型越界 / 目标不存在 / 重复关联 / 候选空结果 / 详情不存在）
  - [x] SubTask 6.2: 冻结异常路径通过 `API` 层测试触发，不通过 `seed` 异常数据
  - [x] SubTask 6.3: 冻结异常前提数据由基线 `seed` 提供（重复关联 / 候选排除）
  - [x] SubTask 6.4: 明确不得新建单独的 `fixture SQL` 文件

- [x] Task 7: 冻结冷启动验收路径。
  - [x] SubTask 7.1: 冻结冷启动路径 `8` 步（`clean-only` -> 空状态 -> `Create` -> `Record` -> `Detail` -> 候选 -> `Link` -> 重新读取）
  - [x] SubTask 7.2: 冻结该路径不依赖任何手工 `SQL`
  - [x] SubTask 7.3: 冻结前提是 `reset_module_mainline.sh` 已建立 `modules` 基线
  - [x] SubTask 7.4: 冻结空状态验收前提（空状态入口 + 主动作进入 `Create`）

- [x] Task 8: 完成规格校验。
  - [x] SubTask 8.1: 验证验收环境建立方式已明确（DoD 1）
  - [x] SubTask 8.2: 验证重置脚本、基线数据与异常路径验证要求已进入阶段任务（DoD 2）
  - [x] SubTask 8.3: 验证不再依赖临时手工 `SQL`（DoD 3）
  - [x] SubTask 8.4: 验证与 `phase02-12` 验收环境模式同构
  - [x] SubTask 8.5: 验证与 `shared_baseline §5.1 / §7` 与 `architecture_plan §4.6 / §4.7` 一致
  - [x] SubTask 8.6: 验证 `decisions` 表原位升级与 `phase03-02` 结构化模板字段一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 2` and `Task 4`
- `Task 6` depends on `Task 5`
- `Task 7` depends on `Task 4` and `Task 5`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, and `Task 7`
