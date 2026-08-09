# Phase05-09 联调验收环境、Dashboard 冷启动基线与 fixture 设计 Spec

## Why

`phase05-02 / 03 / 04` 已冻结反馈信号展示模型、跳转返回路径与聚合读错误语义，`phase05-06 / 07 / 08` 已冻结前端状态模型、后端模块边界与 `.proto` 合同设计。但截至当前，仓库主线里还没有 `Dashboard + Feedback` 联调验收所需的数据库重置入口、冷启动基线种子、七类最小 fixture 组合与跳转返回链路验收矩阵。

如果不把这些验收前提写成单值结论，后续 `phase05-12` 实现与 `phase05-14` 联调验收就会继续在"空系统怎么建立""有产品缺仓库的 fixture 长什么样""局部错误怎么模拟""跳转返回路径怎么验收"之间漂移，最终退回到临时手工 SQL 补票。

## What Changes

- 冻结 `Dashboard + Feedback` 验收环境必须复用既有 `database/scripts/` + `database/seeds/` 模式，不发明第二套验收工具链
- 冻结 `reset_dashboard_acceptance.sh` 作为 Dashboard 验收统一入口，编排既有 reset 脚本并提供 fixture 加载能力
- 冻结七类最小 fixture 组合的命名、数据状态、CTA 映射与文件落点
- 冻结空系统验收基线、最小有数据验收基线与局部错误状态验收基线的建立方式
- 冻结 Dashboard 跳转返回链路验收矩阵，覆盖四类 Detail + 四类 List + Create 回流 + 多跳 + 刷新恢复
- 冻结局部错误状态的模拟方式必须受控且可重复，不得依赖破坏数据库结构
- **BREAKING**：后续 Dashboard 联调验收不得再依赖临时手工 SQL 才能建立最小验收环境

## Impact

- Affected specs:
  - `phase05_dashboard_feedback_foundation`
  - `phase05_02_feedback_signal_priority_display_model`
  - `phase05_03_dashboard_navigation_context_return_path`
  - `phase05_04_dashboard_aggregate_api_error_boundary`
  - `phase05_07_dashboard_feedback_backend_module_boundary_interface_grouping`
- Affected code:
  - 后续 `database/scripts/reset_dashboard_acceptance.sh`（新增）
  - 后续 `database/seeds/seed_dashboard_acceptance_baseline.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_empty_system.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_modules_only.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_products_without_modules.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_products_missing_repository.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_products_missing_module.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_pending_decisions.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_recent_activities.sql`（新增）
  - 后续 `database/seeds/seed_dashboard_fixture_products_missing_all_repositories.sql`（新增，CTA 4 扩展 fixture）
  - 后续 `database/seeds/seed_dashboard_fixture_products_missing_both_bindings.sql`（新增，CTA 6 扩展 fixture）

## ADDED Requirements

### Requirement: Dashboard 验收环境必须复用既有 reset/seed 模式

系统 SHALL 将 `Dashboard + Feedback` 的验收环境建立方式冻结为复用仓库现有 `database/scripts/` + `database/seeds/` 模式，不得为 Dashboard 验收发明第二套工具链或第二套目录结构。

#### Scenario: 验收工具链复用边界

- **WHEN** 后续实现或验收讨论 Dashboard 验收环境如何建立
- **THEN** 必须复用现有 `database/scripts/` 目录承接重置脚本
- **AND** 必须复用现有 `database/seeds/` 目录承接基线种子与 fixture SQL
- **AND** 必须复用现有 `podman exec` / `docker exec` / 宿主机 `psql` 自动检测模式执行 SQL
- **AND** 不得在 `backend/`、`frontend/` 或仓库根新增并列验收工具链
- **AND** 不得引入第二套数据库连接配置或第二套 `psql` 执行包装

#### Scenario: 与既有 reset 脚本的编排关系

- **WHEN** Dashboard 验收需要建立空系统或有数据状态
- **THEN** 必须通过编排既有 `reset_module_mainline.sh`、`reset_decision_mainline.sh` 与 `reset_product_repository_mainline.sh` 实现
- **AND** 必须继续依赖 `seed_readonly_prereqs.sql` 提供只读前提数据
- **AND** 不得绕过既有脚本直接 DELETE 或 TRUNCATE 底层表
- **AND** Dashboard 验收脚本只允许在既有脚本之上增加编排层与 fixture 加载层

### Requirement: reset_dashboard_acceptance.sh 必须作为 Dashboard 验收统一入口

系统 SHALL 冻结 `database/scripts/reset_dashboard_acceptance.sh` 为 Dashboard 验收的唯一统一入口，提供 `--clean-only`、`--restore-only`、`--fixture <name>` 与默认 `clean+restore` 四种模式。

#### Scenario: 脚本模式矩阵

- **WHEN** 后续实现或验收讨论 Dashboard 验收环境如何重置
- **THEN** `reset_dashboard_acceptance.sh` 必须支持以下模式：
  - 默认（无参数）：先清空所有 Dashboard 相关数据，再恢复完整基线
  - `--clean-only`：仅清空所有 Dashboard 相关数据，用于验证空系统状态
  - `--restore-only`：仅恢复完整基线，用于验证有数据状态
  - `--fixture <name>`：先清空，再加载指定 fixture，用于验证特定 CTA 或区块结果
- **AND** `--fixture` 的 `<name>` 只允许取本规格冻结的七类最小 fixture 或两类 CTA 扩展 fixture 名称之一
- **AND** 不得并列发明 `--reset-all`、`--load-fixture`、`--setup` 等第二套参数命名

#### Scenario: 清空范围

- **WHEN** `reset_dashboard_acceptance.sh` 执行清空操作
- **THEN** 必须通过编排既有 `reset_product_repository_mainline.sh --clean-only`、`reset_decision_mainline.sh --clean-only` 与 `reset_module_mainline.sh --clean-only` 实现清空
- **AND** 清空顺序必须按依赖逆序执行：先清 `product_repository`（依赖 modules），再清 `decision`（decision_links 依赖 modules），最后清 `module`（modules 依赖 readonly prereqs 中的 products / decisions）
- **AND** 三个既有脚本的清空范围必须覆盖以下表：`decision_links` / `product_repositories` / `product_modules` / `module_repositories` / `module_releases` / `modules` / `products` / `repositories` / `decisions`
- **AND** 不得绕过既有脚本直接 `DELETE FROM` 或 `TRUNCATE` 底层表
- **AND** 不得清空 schema 或 migration 元数据表

#### Scenario: 恢复范围

- **WHEN** `reset_dashboard_acceptance.sh` 执行恢复操作
- **THEN** 必须按依赖顺序执行：
  1. `seed_readonly_prereqs.sql`（提供 products / repositories / decisions 只读前提）
  2. `reset_module_mainline.sh --restore-only`（恢复 modules / module_releases / product_modules / module_repositories / decision_links）
  3. `reset_decision_mainline.sh --restore-only`（恢复 decisions 正式基线与 decision_links，依赖 modules）
  4. `reset_product_repository_mainline.sh --restore-only`（恢复 products / repositories 正式基线，依赖 modules）
- **AND** 恢复后必须执行 `seed_dashboard_acceptance_baseline.sql` 补齐 Dashboard 验收所需的额外基线数据
- **AND** 恢复操作必须保证幂等，可重复执行不报错

### Requirement: 七类最小 fixture 组合与两类 CTA 扩展 fixture 必须单值化

系统 SHALL 将 Dashboard 验收的七类最小 fixture 组合与两类 CTA 扩展 fixture 冻结为单值结论，每类 fixture 必须有唯一名称、唯一数据状态与唯一 CTA 或区块结果映射。七类最小 fixture 是 `phase05-09` dev_plan 范围明确要求的；两类 CTA 扩展 fixture 用于覆盖 CTA 4 与 CTA 6，确保所有 CTA 都有正式 `--fixture` 入口，不依赖手工 SQL 变体。

#### Scenario: fixture 命名与 CTA 映射矩阵

- **WHEN** 后续实现或验收讨论 Dashboard 验收 fixture 组合
- **THEN** 必须冻结为以下七类最小 fixture 与两类 CTA 扩展 fixture：

| 序号 | fixture 名称 | 类别 | 数据状态摘要 | 映射 CTA / 区块结果 |
|------|-------------|------|-------------|-------------------|
| 1 | `empty-system` | 最小 | 所有表为空 | CTA 1：冷启动空系统 → Module Registry / Create |
| 2 | `modules-only` | 最小 | 仅有 modules + module_releases，无 products / repositories / decisions | CTA 3：module_count > 0 && product_count = 0 → Product Registry / Create |
| 3 | `products-without-modules` | 最小 | 仅有 products（+ readonly prereqs 中的 repositories / decisions），无 modules | CTA 2：非空缺口 → Module Registry / Create |
| 4 | `products-missing-repository` | 最小 | modules + products + product_modules 存在，但目标 product 无 product_repositories | CTA 7：product missing repository binding → Product Detail |
| 5 | `products-missing-module` | 最小 | modules + products + product_repositories 存在，但目标 product 无 product_modules | CTA 8：product missing module binding → Product Detail |
| 6 | `pending-decisions` | 最小 | 完整基线 + 存在 pending 状态的 decisions | CTA 5：pending_decision_signals → Decision Detail / List |
| 7 | `recent-activities` | 最小 | 完整基线 + 所有 product 已完成双绑定 + 多类近期活动项 | Recent Activity 区块 + CTA 9：系统已就绪中性状态 |
| 8 | `products-missing-all-repositories` | CTA 扩展 | modules + products + product_modules 存在，但 repositories 表为空 | CTA 4：module_count > 0 && product_count > 0 && repository_count = 0 → Repository Binding / Create |
| 9 | `products-missing-both-bindings` | CTA 扩展 | 目标 product 同时无 product_modules 与 product_repositories 记录 | CTA 6：product missing both bindings → Product Detail |

- **AND** 每类 fixture 必须有对应的 `seed_dashboard_fixture_<name>.sql` 文件
- **AND** 不得在本规格冻结的九类 fixture 之外并列发明额外 fixture
- **AND** 不得把多类 fixture 合并为一个文件

#### Scenario: fixture 1 — empty-system

- **WHEN** 加载 `empty-system` fixture
- **THEN** 数据库中 `modules / module_releases / products / repositories / decisions / decision_links / product_modules / product_repositories / module_repositories` 必须全部为空
- **AND** DashboardOverviewRead 必须返回 `module_count=0 / product_count=0 / repository_count=0 / decision_count=0`
- **AND** Dashboard 必须命中 CTA 1 冷启动空系统状态
- **AND** 主 CTA 必须指向 `Module Registry / Create`，携带 `fromDashboard=true / dashboardSection=empty-state / dashboardReturnTo=/dashboard`

#### Scenario: fixture 2 — modules-only

- **WHEN** 加载 `modules-only` fixture
- **THEN** 数据库中必须存在至少 `2` 条 modules 与 `2` 条 module_releases
- **AND** `products / repositories / decisions / decision_links / product_modules / product_repositories / module_repositories` 必须为空
- **AND** DashboardOverviewRead 必须返回 `module_count >= 2 / product_count=0 / repository_count=0 / decision_count=0`
- **AND** Dashboard 必须命中 CTA 3 状态
- **AND** 主 CTA 必须指向 `Product Registry / Create`

#### Scenario: fixture 3 — products-without-modules

- **WHEN** 加载 `products-without-modules` fixture
- **THEN** 数据库中必须存在至少 `2` 条 products
- **AND** `modules / module_releases / product_modules / module_repositories / product_repositories / decision_links` 必须为空
- **AND** `repositories` 与 `decisions` 允许保留 readonly prereqs 中的最小占位数据
- **AND** DashboardOverviewRead 必须返回 `module_count=0 / product_count >= 2`
- **AND** Dashboard 必须命中 CTA 2 非空缺口状态
- **AND** 主 CTA 必须指向 `Module Registry / Create`
- **AND** 该状态不得与冷启动空系统混同

#### Scenario: fixture 4 — products-missing-repository

- **WHEN** 加载 `products-missing-repository` fixture
- **THEN** 数据库中必须存在：
  - 至少 `1` 条已绑定 module 但未绑定 repository 的 product（目标 product）
  - 至少 `1` 条已完整绑定的 product（对照 product，用于验证非全局缺口）
- **AND** 目标 product 必须在 `product_modules` 中有记录，但在 `product_repositories` 中无记录
- **AND** 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
- **AND** DashboardOverviewRead 的 `product_with_repository_count` 必须小于 `product_count`
- **AND** FeedbackSignalRead 必须产出 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING` 信号
- **AND** FeedbackSignalRead 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION` 信号
- **AND** Dashboard 必须命中 CTA 7 状态
- **AND** 主 CTA 必须指向目标 product 的 `Product Detail`

#### Scenario: fixture 5 — products-missing-module

- **WHEN** 加载 `products-missing-module` fixture
- **THEN** 数据库中必须存在：
  - 至少 `1` 条已绑定 repository 但未绑定 module 的 product（目标 product）
  - 至少 `1` 条已完整绑定的 product（对照 product）
- **AND** 目标 product 必须在 `product_repositories` 中有记录，但在 `product_modules` 中无记录
- **AND** 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
- **AND** DashboardOverviewRead 的 `product_with_module_count` 必须小于 `product_count`
- **AND** FeedbackSignalRead 必须产出 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING` 信号
- **AND** FeedbackSignalRead 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION` 信号
- **AND** Dashboard 必须命中 CTA 8 状态
- **AND** 主 CTA 必须指向目标 product 的 `Product Detail`

#### Scenario: fixture 6 — pending-decisions

- **WHEN** 加载 `pending-decisions` fixture
- **THEN** 数据库中必须存在至少 `1` 条 pending 状态的 decision
- **AND** 必须存在至少 `1` 条已绑定具体 `decision_id` 的 decision_link（用于验证单项信号跳转 `Decision Detail`）
- **AND** 允许存在 `1` 条未绑定单一 `decision_id` 的聚合决策信号（用于验证聚合信号跳转 `Decision Center / List`）
- **AND** FeedbackSignalRead 必须产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION` 信号
- **AND** Dashboard 必须命中 CTA 5 状态
- **AND** 主 CTA 必须指向最高优先级决策信号落点
- **AND** pending decision 的状态判定必须沿用 `phase03` Decision Center 已冻结的 status 语义

#### Scenario: fixture 7 — recent-activities

- **WHEN** 加载 `recent-activities` fixture
- **THEN** 数据库中必须存在覆盖以下活动类型的近期数据：
  - `module` 类型：至少 `1` 条近期创建的 module
  - `release` 类型：至少 `1` 条近期创建的 module_release
  - `product` 类型：至少 `1` 条近期创建的 product
  - `repository` 类型：至少 `1` 条近期创建的 repository
  - `decision` 类型：至少 `1` 条近期创建的非 pending decision
  - `product_module_binding` 类型：至少 `1` 条近期创建的 product_modules 记录
  - `product_repository_binding` 类型：至少 `1` 条近期创建的 product_repositories 记录
  - `module_repository_binding` 类型：至少 `1` 条近期创建的 module_repositories 记录
- **AND** 所有 product 必须已完成模块与仓库双绑定，使 Dashboard 不命中任何缺口 CTA
- **AND** 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
- **AND** 所有活动项必须带有显式时间字段，且时间分布在可排序的近期范围内
- **AND** RecentActivityRead 必须返回最多 `10` 条活动项，按 `activity_at` 倒序排序
- **AND** FeedbackSignalRead 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION`、`FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS`、`FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING` 或 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING`
- **AND** Dashboard 必须命中 CTA 9 系统已就绪中性状态，不展示强制主 CTA

#### Scenario: fixture 8 — products-missing-all-repositories（CTA 4 扩展）

- **WHEN** 加载 `products-missing-all-repositories` fixture
- **THEN** 数据库中必须存在至少 `1` 条 module 与至少 `2` 条 products
- **AND** 目标 products 必须在 `product_modules` 中有记录
- **AND** `repositories` 表必须为空，`product_repositories` 与 `module_repositories` 必须为空
- **AND** DashboardOverviewRead 必须返回 `module_count > 0 / product_count > 0 / repository_count = 0`
- **AND** Dashboard 必须命中 CTA 4 状态
- **AND** 主 CTA 必须指向 `Repository Binding / Create`，携带 `fromDashboard=true / dashboardSection=empty-state / dashboardReturnTo=/dashboard`

#### Scenario: fixture 9 — products-missing-both-bindings（CTA 6 扩展）

- **WHEN** 加载 `products-missing-both-bindings` fixture
- **THEN** 数据库中必须存在：
  - 至少 `1` 条同时缺少 `product_modules` 与 `product_repositories` 记录的 product（目标 product）
  - 至少 `1` 条已完整绑定的 product（对照 product，用于验证非全局缺口）
- **AND** 目标 product 必须在 `product_modules` 与 `product_repositories` 中均无记录
- **AND** `modules` 与 `repositories` 必须存在，使缺口判定为"product 缺绑定"而非"系统缺实体"
- **AND** 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
- **AND** FeedbackSignalRead 必须产出 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS` 信号
- **AND** FeedbackSignalRead 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION` 信号
- **AND** Dashboard 必须命中 CTA 6 状态
- **AND** 主 CTA 必须指向目标 product 的 `Product Detail`

### Requirement: CTA 4 与 CTA 6 的覆盖方式必须明确

系统 SHALL 明确 CTA 4（module_count > 0 && product_count > 0 && repository_count = 0）与 CTA 6（product missing both bindings）的验收覆盖方式必须通过正式 `--fixture` 入口建立，不得依赖 fixture 变体或临时手工 SQL。

#### Scenario: CTA 4 覆盖方式

- **WHEN** 验收需要覆盖 CTA 4（无仓库但有模块与产品）
- **THEN** 必须通过 `reset_dashboard_acceptance.sh --fixture products-missing-all-repositories` 正式建立
- **AND** 该 fixture 必须有对应的 `seed_dashboard_fixture_products_missing_all_repositories.sql` 文件
- **AND** DashboardOverviewRead 必须返回 `module_count > 0 / product_count > 0 / repository_count = 0`
- **AND** 主 CTA 必须指向 `Repository Binding / Create`
- **AND** 不得通过手工清空 `repositories` 表的方式建立该状态

#### Scenario: CTA 6 覆盖方式

- **WHEN** 验收需要覆盖 CTA 6（product missing both bindings）
- **THEN** 必须通过 `reset_dashboard_acceptance.sh --fixture products-missing-both-bindings` 正式建立
- **AND** 该 fixture 必须有对应的 `seed_dashboard_fixture_products_missing_both_bindings.sql` 文件
- **AND** FeedbackSignalRead 必须产出 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS` 信号
- **AND** FeedbackSignalRead 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION` 信号
- **AND** Dashboard 必须命中 CTA 6 状态
- **AND** 主 CTA 必须指向目标 product 的 `Product Detail`
- **AND** 不得通过手工删除 `product_modules` 记录的方式建立该状态

#### Scenario: CTA 9 覆盖方式

- **WHEN** 验收需要覆盖 CTA 9（无缺口且有活动数据）
- **THEN** 必须通过 `reset_dashboard_acceptance.sh --fixture recent-activities` 正式建立
- **AND** 该 fixture 基于完整基线且所有 product 均已完成双绑定
- **AND** 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
- **AND** `decision` 类型活动只能来自非 pending decision
- **AND** Dashboard 必须进入"系统已就绪"中性状态
- **AND** 不得展示强制主 CTA

### Requirement: fixture 文件落点与命名必须单值化

系统 SHALL 将 Dashboard fixture 文件的落点与命名冻结为单值结论，沿用既有 `database/seeds/` 扁平命名模式。

#### Scenario: fixture 文件命名规则

- **WHEN** 后续实现新增 Dashboard fixture 文件
- **THEN** 基线文件必须命名为 `seed_dashboard_acceptance_baseline.sql`
- **AND** 九类 fixture 文件必须命名为 `seed_dashboard_fixture_<fixture-name>.sql`，其中 `<fixture-name>` 为本规格冻结的九类名称之一
- **AND** 所有 fixture 文件必须落点在 `database/seeds/` 目录
- **AND** 不得在 `database/seeds/` 下新建 `dashboard/` 或 `fixtures/` 等子目录
- **AND** 不得使用 `fixture-`、`dash-fixture-`、`acceptance_` 等第二套前缀命名

#### Scenario: fixture SQL 幂等约束

- **WHEN** 后续实现编写 fixture SQL 文件
- **THEN** 每个 fixture 文件必须可重复执行不报错
- **AND** 必须使用 `ON CONFLICT DO NOTHING`、`WHERE NOT EXISTS` 或 `UPDATE-then-INSERT` 模式保证幂等
- **AND** 必须在文件头部以注释说明：用途、依赖、幂等语义、使用方式
- **AND** 必须沿用既有 `seed_module_mainline_baseline.sql` 与 `seed_product_repository_mainline_baseline.sql` 的注释结构

### Requirement: 空系统验收基线必须可重复建立

系统 SHALL 将 Dashboard 空系统验收基线冻结为通过 `reset_dashboard_acceptance.sh --clean-only` 或 `--fixture empty-system` 建立，且可重复执行。

#### Scenario: 空系统建立方式

- **WHEN** 验收需要建立空系统状态
- **THEN** 必须执行 `reset_dashboard_acceptance.sh --clean-only` 或 `reset_dashboard_acceptance.sh --fixture empty-system`
- **AND** 执行后所有 Dashboard 相关表必须为空
- **AND** 重复执行不得报错
- **AND** 不得依赖手工 `DELETE FROM` 或 `TRUNCATE` 语句

#### Scenario: 空系统验证点

- **WHEN** 空系统建立后进行 Dashboard 验收
- **THEN** DashboardOverviewRead 必须以成功语义返回所有计数为 `0`
- **AND** FeedbackSignalRead 必须以成功语义返回空列表与零值计数
- **AND** RecentActivityRead 必须以成功语义返回空列表
- **AND** Dashboard 必须命中 CTA 1 冷启动空系统状态
- **AND** 不得把空系统误判为整页失败或资源不存在

### Requirement: 最小有数据验收基线必须可重复建立

系统 SHALL 将 Dashboard 最小有数据验收基线冻结为通过 `reset_dashboard_acceptance.sh`（默认模式）或 `--restore-only` 建立，且可重复执行。

#### Scenario: 最小有数据基线建立方式

- **WHEN** 验收需要建立最小有数据状态
- **THEN** 必须执行 `reset_dashboard_acceptance.sh`（默认）或 `reset_dashboard_acceptance.sh --restore-only`
- **AND** 执行后数据库必须包含：
  - 至少 `2` 条 modules、`3` 条 module_releases
  - 至少 `3` 条 products、`2` 条 repositories
  - 至少 `2` 条 product_modules、`1` 条 product_repositories、`2` 条 module_repositories
  - 至少 `1` 条 decision（来自 readonly prereqs）
- **AND** DashboardOverviewRead 必须返回非零计数
- **AND** 重复执行不得报错

#### Scenario: 最小有数据基线验证点

- **WHEN** 最小有数据基线建立后进行 Dashboard 验收
- **THEN** DashboardOverviewRead 必须返回 `module_count >= 2 / product_count >= 3 / repository_count >= 2 / decision_count >= 1`
- **AND** FeedbackSignalRead 必须能识别资产缺口（因基线中存在缺仓库或缺模块的 product）
- **AND** RecentActivityRead 必须能返回至少 `1` 条活动项
- **AND** Dashboard 不得进入冷启动空系统状态

### Requirement: 局部错误状态验收基线必须可模拟且可重复

系统 SHALL 将 Dashboard 局部错误状态的模拟方式冻结为受控环境变量单一入口，不得依赖破坏数据库结构、临时手工操作或 SQL fixture 文件。局部错误模拟不产生 `seed_dashboard_fixture_local_errors.sql` 文件，只通过环境变量在后端实现层承接。

#### Scenario: 局部错误模拟机制

- **WHEN** 验收需要模拟 DashboardOverviewRead 失败
- **THEN** 必须通过受控机制（如环境变量 `DASHBOARD_SIMULATE_OVERVIEW_ERROR=true`）触发 overview reader 返回错误
- **AND** 该机制必须由后端在实现阶段承接，不要求 fixture SQL 层模拟
- **AND** 该机制必须可重复启用与关闭
- **AND** 不得通过删除表、重命名列或破坏 schema 来模拟错误

- **WHEN** 验收需要模拟 FeedbackSignalRead 失败
- **THEN** 必须通过受控机制（如环境变量 `DASHBOARD_SIMULATE_FEEDBACK_ERROR=true`）触发 feedback reader 返回错误
- **AND** DashboardOverviewRead 必须仍然成功
- **AND** Dashboard 只允许 `Current Focus / Next Action` 与 `Asset Feedback` 进入局部失败语义
- **AND** 不得强制整页失败

- **WHEN** 验收需要模拟 RecentActivityRead 失败
- **THEN** 必须通过受控机制（如环境变量 `DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR=true`）触发 recent activity reader 返回错误
- **AND** DashboardOverviewRead 必须仍然成功
- **AND** Dashboard 只允许 `Recent Activity` 区块进入局部失败语义
- **AND** 不得强制整页失败

#### Scenario: 局部错误验证点

- **WHEN** 局部错误启用后进行 Dashboard 验收
- **THEN** 必须验证 `phase05-04` 已冻结的以下语义：
  - overview 失败 → 整页失败
  - feedback 失败 → 局部失败，overview 与 recent activity 仍成功
  - recent activity 失败 → 局部失败，overview 与 feedback 仍成功
- **AND** 局部失败区块必须展示局部重试入口，而不是整页主 CTA
- **AND** 局部错误模拟机制的详细实现由 `phase05-12` 承接，本规格只冻结模拟方式必须受控且可重复

### Requirement: Dashboard 跳转返回链路验收矩阵必须覆盖完整路径

系统 SHALL 将 Dashboard 跳转返回链路的验收矩阵冻结为覆盖四类 Detail + 四类 List + Create 回流 + 多跳 + 刷新恢复的单值结论。

#### Scenario: 直接跳转到 Detail 页后返回

- **WHEN** 验收 Dashboard 直接跳转到 canonical owner Detail 页后返回
- **THEN** 必须覆盖以下四类 Detail 路径：
  - Dashboard → Module Detail → 返回 `/dashboard`
  - Dashboard → Product Detail → 返回 `/dashboard`
  - Dashboard → Repository Binding Detail / Workspace → 返回 `/dashboard`
  - Dashboard → Decision Detail → 返回 `/dashboard`
- **AND** 每条路径必须验证目标页携带 `fromDashboard=true / dashboardSection / dashboardReturnTo=/dashboard`
- **AND** 每条路径必须验证返回后回到 `/dashboard`
- **AND** 每条路径必须验证返回后恢复到对应来源区块的浏览上下文

#### Scenario: 直接跳转到 List 页后返回

- **WHEN** 验收 Dashboard 直接跳转到 canonical owner List 页后返回
- **THEN** 必须覆盖以下四类 List 路径：
  - Dashboard → Module Registry / List → 返回 `/dashboard`
  - Dashboard → Product Registry / List → 返回 `/dashboard`
  - Dashboard → Repository Binding / List → 返回 `/dashboard`
  - Dashboard → Decision Center / List → 返回 `/dashboard`
- **AND** 每条路径必须验证目标页携带 `fromDashboard=true / dashboardSection=overview / dashboardReturnTo=/dashboard`
- **AND** 每条路径必须验证返回后回到 `/dashboard`

#### Scenario: 多跳路径（Dashboard → List → Detail）

- **WHEN** 验收 Dashboard 来源下从 List 再进入 Detail 的多跳路径
- **THEN** 必须覆盖四类 canonical owner 的 List → Detail 路径
- **AND** Detail 页必须同时保留 `fromList`（原生列表来源）与 `fromDashboard=true`（外层来源）
- **AND** 从 Detail 返回 List 必须使用 `fromList` 上下文
- **AND** 从 List 返回 Dashboard 必须使用 `fromDashboard` 上下文
- **AND** 不得把两者混写成同一个主来源字段

#### Scenario: Create 页回流路径

- **WHEN** 验收 Dashboard 来源下进入 Create 页后的回流路径
- **THEN** 必须覆盖以下三类 Create 路径：
  - Dashboard → Module Registry / Create → 取消 → 返回 `/dashboard`
  - Dashboard → Module Registry / Create → 提交 → Module Detail → 返回 `/dashboard`
  - Dashboard → Product Registry / Create → 取消 → 返回 `/dashboard`
  - Dashboard → Product Registry / Create → 提交 → Product Detail → 返回 `/dashboard`
  - Dashboard → Repository Binding / Create → 取消 → 返回 `/dashboard`
  - Dashboard → Repository Binding / Create → 提交 → Repository Binding Detail → 返回 `/dashboard`
- **AND** Create 页必须携带 `fromDashboard=true / dashboardSection=empty-state / dashboardReturnTo=/dashboard`
- **AND** 取消时必须返回 `/dashboard`，而不是回列表
- **AND** 提交成功后进入 Detail 页必须继续保留 `fromDashboard=true`
- **AND** 从 Detail 页返回必须能回到 `/dashboard`

#### Scenario: 刷新恢复路径

- **WHEN** 验收 Dashboard 来源下目标页刷新后的上下文恢复
- **THEN** 必须验证刷新后 `fromDashboard / dashboardSection / dashboardReturnTo` 参数不丢失
- **AND** 必须覆盖至少一类 Detail 页与一类 List 页的刷新场景

#### Scenario: 非法参数回退路径

- **WHEN** 验收 Dashboard 来源参数缺失或非法时的回退
- **THEN** `dashboardSection` 非法时必须回退为 `overview`
- **AND** `dashboardReturnTo` 缺失或非法时必须回退为 `/dashboard`
- **AND** 不得静默跳到错误页面或根路径 `/`

### Requirement: Dashboard 验收不得依赖临时手工 SQL

系统 SHALL 冻结 Dashboard 验收环境的建立必须通过受控脚本与 fixture 文件完成，不得依赖临时手工 SQL。

#### Scenario: 验收环境建立方式

- **WHEN** 后续 `phase05-12` 实现或 `phase05-14` 联调验收需要建立 Dashboard 验收环境
- **THEN** 必须通过 `reset_dashboard_acceptance.sh` 的四种模式建立
- **AND** 必须通过 `seed_dashboard_fixture_<name>.sql` 加载特定 fixture
- **AND** 局部错误状态必须通过受控环境变量模拟
- **AND** 不得在验收过程中临时执行手工 `INSERT`、`DELETE`、`UPDATE` 语句
- **AND** 不得在验收报告中引用未纳入仓库的临时 SQL 片段作为验收证据

## MODIFIED Requirements

### Requirement: phase05 冷启动与验收基线解释

`phase05` 共享基线 §7 中的"当前阶段必须提供可重复执行的重置脚本、基线种子与异常路径验证前提"在 `phase05-09` SHALL 被进一步解释为"必须通过 `reset_dashboard_acceptance.sh` 统一入口编排既有 reset 脚本，并通过九类 fixture 文件（七类最小 + 两类 CTA 扩展）与受控环境变量局部错误模拟机制覆盖所有验收路径"。

#### Scenario: 验收基线承接方式

- **WHEN** 后续实现或验收讨论 Dashboard 冷启动与验收基线
- **THEN** 必须优先引用 `reset_dashboard_acceptance.sh` 与 `seed_dashboard_fixture_*.sql`
- **AND** 必须同时满足 `phase05-02 / 03 / 04 / 07 / 08` 已冻结的字段模板、导航语义、错误语义与后端 owner 规则
- **AND** 不得在受控脚本与 fixture 之外再发明并列验收入口

## REMOVED Requirements

### Requirement: Dashboard 验收依赖临时手工 SQL

**Reason**: `phase05-09` 的目标就是把 Dashboard 验收环境从"联调时临时补 SQL"推进到"仓库内有受控脚本与 fixture 文件"的状态。若继续允许临时手工 SQL，验收环境不可重复，验收证据不可复验。

**Migration**: 后续 Dashboard 验收环境的建立、重置与 fixture 加载，必须统一通过 `reset_dashboard_acceptance.sh` 与 `seed_dashboard_fixture_*.sql` 完成；局部错误状态必须通过受控环境变量模拟。
