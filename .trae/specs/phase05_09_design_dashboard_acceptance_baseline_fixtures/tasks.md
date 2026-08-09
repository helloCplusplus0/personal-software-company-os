# Tasks

- [x] Task 1: 对齐 `phase05-09` 直接上游与既有验收环境基线，确定本次验收环境只复用现有 `database/scripts/` + `database/seeds/` 模式。
  - [x] SubTask 1.1: 对齐 `phase05-02 / 03 / 04 / 07 / 08` 已冻结的展示模型、导航返回、错误语义、后端边界与 `.proto` 合同
  - [x] SubTask 1.2: 对齐 `phase02-12 / phase03-14 / phase04-14` 已落地的 reset/seed 模式（`--clean-only / --restore-only / default`、`ON CONFLICT DO NOTHING`、`WHERE NOT EXISTS`、`UPDATE-then-INSERT`）
  - [x] SubTask 1.3: 对齐 `phase05` 共享基线 §7 冷启动与验收基线、CTA 优先级矩阵 1-9
  - [x] SubTask 1.4: 冻结 Dashboard 验收环境不发明第二套工具链，只新增编排层与 fixture 层

- [x] Task 2: 冻结 `reset_dashboard_acceptance.sh` 的模式矩阵、清空范围与恢复范围。
  - [x] SubTask 2.1: 冻结四种模式：默认 `clean+restore`、`--clean-only`、`--restore-only`、`--fixture <name>`
  - [x] SubTask 2.2: 冻结清空范围通过编排既有 `reset_product_repository_mainline.sh --clean-only`、`reset_decision_mainline.sh --clean-only` 与 `reset_module_mainline.sh --clean-only` 实现，按依赖逆序执行
  - [x] SubTask 2.3: 冻结恢复范围按依赖顺序：`seed_readonly_prereqs.sql` → `reset_module_mainline.sh --restore-only` → `reset_decision_mainline.sh --restore-only` → `reset_product_repository_mainline.sh --restore-only` → `seed_dashboard_acceptance_baseline.sql`
  - [x] SubTask 2.4: 冻结 `--fixture <name>` 只允许取七类最小 fixture 或两类 CTA 扩展 fixture 名称之一
  - [x] SubTask 2.5: 冻结 `decisions` 表清空通过编排既有 `reset_decision_mainline.sh --clean-only` 实现，不得绕过既有脚本直接 DELETE
  - [x] SubTask 2.6: 冻结所有 `--fixture` 模式统一遵守"先清空，再加载指定 fixture"语义，不允许任何 fixture 叠加已有数据

- [x] Task 3: 冻结七类最小 fixture 与两类 CTA 扩展 fixture 的命名、数据状态与 CTA 映射。
  - [x] SubTask 3.1: 冻结 fixture 1 `empty-system`：所有表为空，映射 CTA 1
  - [x] SubTask 3.2: 冻结 fixture 2 `modules-only`：仅有 modules + module_releases，映射 CTA 3
  - [x] SubTask 3.3: 冻结 fixture 3 `products-without-modules`：仅有 products，映射 CTA 2
  - [x] SubTask 3.4: 冻结 fixture 4 `products-missing-repository`：目标 product 有 module 无 repository，且不得同时命中 `pending_decision_signals`，映射 CTA 7
  - [x] SubTask 3.5: 冻结 fixture 5 `products-missing-module`：目标 product 有 repository 无 module，且不得同时命中 `pending_decision_signals`，映射 CTA 8
  - [x] SubTask 3.6: 冻结 fixture 6 `pending-decisions`：存在 pending 状态 decisions，映射 CTA 5
  - [x] SubTask 3.7: 冻结 fixture 7 `recent-activities`：所有 product 双绑定 + 覆盖八类活动类型 + 不得产生 `pending_decision_signals`，映射 Recent Activity 区块与 CTA 9
  - [x] SubTask 3.8: 冻结 fixture 8 `products-missing-all-repositories`（CTA 扩展）：modules + products 存在但 repositories 为空，映射 CTA 4
  - [x] SubTask 3.9: 冻结 fixture 9 `products-missing-both-bindings`（CTA 扩展）：目标 product 同时无 product_modules 与 product_repositories，且不得同时命中 `pending_decision_signals`，映射 CTA 6

- [x] Task 4: 冻结 CTA 4、CTA 6 与 CTA 9 的验收覆盖方式必须通过正式 `--fixture` 入口建立。
  - [x] SubTask 4.1: 冻结 CTA 4 通过 `--fixture products-missing-all-repositories` 正式建立，有独立 SQL 文件
  - [x] SubTask 4.2: 冻结 CTA 6 通过 `--fixture products-missing-both-bindings` 正式建立，有独立 SQL 文件
  - [x] SubTask 4.3: 冻结 CTA 9 直接由 `--fixture recent-activities` 覆盖（完整基线 + 所有 product 双绑定 + 多类活动项）
  - [x] SubTask 4.4: 冻结 CTA 4 / 6 不得通过 fixture 变体或临时手工 SQL 建立

- [x] Task 5: 冻结 fixture 文件落点与命名规则。
  - [x] SubTask 5.1: 冻结基线文件命名为 `seed_dashboard_acceptance_baseline.sql`
  - [x] SubTask 5.2: 冻结九类 fixture 文件命名为 `seed_dashboard_fixture_<name>.sql`
  - [x] SubTask 5.3: 冻结所有文件落点在 `database/seeds/`，不新建子目录
  - [x] SubTask 5.4: 冻结 fixture SQL 幂等约束：`ON CONFLICT DO NOTHING` / `WHERE NOT EXISTS` / `UPDATE-then-INSERT` + 文件头注释结构

- [x] Task 6: 冻结空系统与最小有数据验收基线的建立方式与验证点。
  - [x] SubTask 6.1: 冻结空系统通过 `--clean-only` 或 `--fixture empty-system` 建立
  - [x] SubTask 6.2: 冻结空系统验证点：所有计数为 0、成功语义、CTA 1 命中
  - [x] SubTask 6.3: 冻结最小有数据通过默认模式或 `--restore-only` 建立
  - [x] SubTask 6.4: 冻结最小有数据验证点：非零计数、缺口识别、活动项可见

- [x] Task 7: 冻结局部错误状态验收基线的模拟方式与验证点。
  - [x] SubTask 7.1: 冻结局部错误通过受控环境变量单一入口模拟，不破坏数据库结构
  - [x] SubTask 7.2: 冻结三个环境变量：`DASHBOARD_SIMULATE_OVERVIEW_ERROR` / `DASHBOARD_SIMULATE_FEEDBACK_ERROR` / `DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR`
  - [x] SubTask 7.3: 冻结局部错误不产生 `seed_dashboard_fixture_local_errors.sql` 文件，只通过环境变量承接
  - [x] SubTask 7.4: 冻结 overview 失败 → 整页失败语义
  - [x] SubTask 7.5: 冻结 feedback 失败 → 局部失败，overview 与 recent activity 仍成功
  - [x] SubTask 7.6: 冻结 recent activity 失败 → 局部失败，overview 与 feedback 仍成功
  - [x] SubTask 7.7: 冻结局部错误模拟的详细实现由 `phase05-12` 承接，本规格只冻结模拟方式必须受控且可重复

- [x] Task 8: 冻结 Dashboard 跳转返回链路验收矩阵。
  - [x] SubTask 8.1: 冻结四类 Detail 直接跳转返回路径（Module / Product / Repository / Decision Detail）
  - [x] SubTask 8.2: 冻结四类 List 直接跳转返回路径（Module / Product / Repository / Decision List）
  - [x] SubTask 8.3: 冻结多跳路径（Dashboard → List → Detail → List → Dashboard）的 fromList + fromDashboard 共存规则
  - [x] SubTask 8.4: 冻结三类 Create 页回流路径（Module / Product / Repository Binding Create），含取消返回与提交后 Detail 返回
  - [x] SubTask 8.5: 冻结刷新恢复路径（参数不丢失）与非法参数回退路径（dashboardSection 回退 overview，dashboardReturnTo 回退 /dashboard）

- [x] Task 9: 冻结 Dashboard 验收不得依赖临时手工 SQL 的约束。
  - [x] SubTask 9.1: 冻结验收环境必须通过 `reset_dashboard_acceptance.sh` 建立
  - [x] SubTask 9.2: 冻结 fixture 必须通过 `seed_dashboard_fixture_*.sql` 加载
  - [x] SubTask 9.3: 冻结局部错误必须通过受控环境变量模拟，不产生 SQL fixture 文件
  - [x] SubTask 9.4: 冻结验收报告中不得引用未纳入仓库的临时 SQL 作为证据

- [x] Task 10: 完成 `phase05-09` 规格一致性校验。
  - [x] SubTask 10.1: 验证九类 fixture 每类都能映射到单值 CTA 或单值区块结果
  - [x] SubTask 10.2: 验证 CTA 1-9 全部有正式 `--fixture` 入口覆盖，不依赖变体或手工 SQL
  - [x] SubTask 10.2A: 验证 CTA 6 / 7 / 8 / 9 对应 fixture 都显式排除了更高优先级的 `pending_decision_signals` 抢占
  - [x] SubTask 10.3: 验证验收矩阵覆盖四类 Detail + 四类 List + Create 回流 + 多跳 + 刷新 + 非法回退
  - [x] SubTask 10.4: 验证与 `phase05-02 / 03 / 04 / 07 / 08` 已冻结结论一致
  - [x] SubTask 10.5: 验证与既有 `reset_module_mainline.sh` / `reset_decision_mainline.sh` / `reset_product_repository_mainline.sh` 模式同构

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 3`
- `Task 6` depends on `Task 2` and `Task 5`
- `Task 7` depends on `Task 1` and `Task 5`
- `Task 8` depends on `Task 1`
- `Task 9` depends on `Task 2` and `Task 5` and `Task 7`
- `Task 10` depends on `Task 1` through `Task 9`
