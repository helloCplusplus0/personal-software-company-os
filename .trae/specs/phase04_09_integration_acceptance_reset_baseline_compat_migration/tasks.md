# Tasks

- [x] Task 1: 冻结 `phase04` 联调环境建立方式与启动顺序。
  - [x] SubTask 1.1: 冻结复用 `init_db.sh`、`run_seeds.sh`、`reset_module_mainline.sh` 的单一入口模式
  - [x] SubTask 1.2: 冻结 `init_db -> migration(含 0006) -> run_seeds -> reset_module_mainline -> reset_product_repository_mainline -> 前端` 的启动顺序
  - [x] SubTask 1.3: 明确当前阶段不要求 `reset_decision_mainline.sh` 成为 `phase04` 验收前置步骤

- [x] Task 2: 冻结 `products / repositories / product_repositories` 的主线升级 migration 设计。
  - [x] SubTask 2.1: 冻结 migration 文件落点 `database/migrations/0006_product_repository_binding_mainline.sql`
  - [x] SubTask 2.2: 冻结 `products` 原位新增 `description / status`、`repositories` 原位新增 `url / provider / status`
  - [x] SubTask 2.3: 冻结 `product_repositories` 表结构、唯一约束与读取索引
  - [x] SubTask 2.4: 冻结 `active / archived` 状态约束与列表读取索引
  - [x] SubTask 2.5: 明确不得新建影子表、不得删除既有 `id / name / created_at`

- [x] Task 3: 冻结历史数据与历史绑定的兼容回填策略。
  - [x] SubTask 3.1: 冻结 `products` 历史记录的 `description / status` 回填策略
  - [x] SubTask 3.2: 冻结 `repositories` 历史记录的 `url / provider / status` 回填策略
  - [x] SubTask 3.3: 冻结回填保留既有 `id / name / created_at` 且具备幂等性
  - [x] SubTask 3.4: 冻结 `product_modules / module_repositories` 历史关系继续可读，不通过重建关系表迁移

- [x] Task 4: 冻结 `phase04` 重置脚本与清空策略。
  - [x] SubTask 4.1: 冻结脚本落点 `database/scripts/reset_product_repository_mainline.sh`
  - [x] SubTask 4.2: 冻结三种模式（`--clean-only / --restore-only / 默认`）与 `resolve_psql` 同构实现
  - [x] SubTask 4.3: 冻结清空范围（`product_repositories / product_modules / module_repositories / products / repositories`）
  - [x] SubTask 4.4: 冻结前置校验（数据库存在 + `modules` 基线存在）
  - [x] SubTask 4.5: 冻结继续使用 `DELETE FROM ...`，不改用 `TRUNCATE ... RESTART IDENTITY ... CASCADE`

- [x] Task 5: 冻结 `phase04` 基线 seed 与 fixture 设计。
  - [x] SubTask 5.1: 冻结 seed 文件落点 `database/seeds/seed_product_repository_mainline_baseline.sql`
  - [x] SubTask 5.2: 冻结 `Product / Repository` 基线覆盖维度（数量、状态、provider）
  - [x] SubTask 5.3: 冻结三类绑定关系的最小基线覆盖
  - [x] SubTask 5.4: 冻结 fixture 只允许“基线 seed + reset 模式 + API 操作”两层策略
  - [x] SubTask 5.5: 冻结 `seed_readonly_prereqs.sql` 仅保留历史兼容最小前提定位
  - [x] SubTask 5.6: 冻结 `seed_readonly_prereqs.sql` 中 `products / repositories` seed 的字段升级策略，确保与 `0006` migration `NOT NULL` 约束兼容
    - 证据：`spec.md` §ADDED Requirements「phase04 基线 seed 与 fixture 设计必须冻结」Scenario: seed_readonly_prereqs.sql products / repositories seed 字段升级 L222-232
  - [x] SubTask 5.7: 冻结 `seed_product_repository_mainline_baseline.sql` 与 `seed_module_mainline_baseline.sql` 的 name 兼容策略，确保两条重置链可交叉重复执行
    - 证据：`spec.md` §ADDED Requirements「phase04 基线 seed 与 fixture 设计必须冻结」Scenario: 基线 seed 与 module 基线的 name 兼容必须冻结 L184-191

- [x] Task 6: 冻结冷启动主路径、兼容入口主路径与多入口回流验收矩阵。
  - [x] SubTask 6.1: 冻结从空状态到首个 `Product`、首个 `Repository`、首轮三类绑定的冷启动路径
  - [x] SubTask 6.2: 冻结 `Module Detail -> Product` 与 `Module Detail -> Repository` 的旧入口兼容路径
  - [x] SubTask 6.3: 冻结 `Product Detail -> Repository` 的兼容跳转路径
  - [x] SubTask 6.4: 冻结 `fromList / fromModuleDetail / fromProductDetail / direct-entry` 四类回流矩阵

- [x] Task 7: 冻结旧 transport 兼容与异常路径验证要求。
  - [x] SubTask 7.1: 冻结旧候选读取/旧绑定入口只能作为兼容适配层委派给 canonical owner，并逐项枚举保留的兼容面（`ProductBindingCandidateRead` / `ModuleBindingWrite` / `RepositoryBindingCandidateRead` / `Module Detail` 旧入口兼容跳转）要求逐项验收
    - 证据：`spec.md` §ADDED Requirements「旧 transport 兼容与异常路径验证要求必须冻结」Scenario: 旧 transport 兼容委派 L298-308
  - [x] SubTask 7.2: 冻结最小异常路径清单（创建校验、目标不存在、目标非 active、重复冲突、列表空、候选空、详情不存在）
  - [x] SubTask 7.3: 冻结异常前提通过基线 seed 与受控 API 操作建立，不新增独立 fixture SQL

- [x] Task 8: 完成规格自校验并对齐 DoD。
  - [x] SubTask 8.1: 验证验收环境建立方式已明确（DoD 1）
  - [x] SubTask 8.2: 验证重置脚本、基线数据、旧绑定兼容与异常路径验证要求已进入阶段任务（DoD 2）
  - [x] SubTask 8.3: 验证冷启动主路径、兼容入口主路径与三类绑定 reread 路径已明确（DoD 3）
  - [x] SubTask 8.4: 验证不再依赖临时手工 SQL（DoD 4）
  - [x] SubTask 8.5: 验证与 `phase04-03 / 04 / 06 / 07 / 08` 已冻结结论一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 2` and `Task 4`
- `Task 6` depends on `Task 1`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 2`, `Task 5`, and `Task 6`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, and `Task 7`
