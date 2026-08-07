# Phase04-09 联调验收环境、重置基线与兼容迁移设计 Spec

## Why

`phase04-03 / 04 / 06 / 07 / 08` 已分别冻结三类绑定关系、数据与错误边界、前端交互流、后端模块边界与最小 `.proto` 合同，但当前仍缺少一份“如何把这些冻结结果放进同一可重复环境里验证”的收口设计。若没有这一层，`Product / Repository / Binding` 的实现将继续依赖手工 SQL、一次性数据库状态或历史临时入口，无法形成可重复复核的联调主线。

同时，`products / repositories` 目前仍承接 `phase02` 的只读前提语义，`phase04` 还需要把它们升级为正式主线数据、补上 `product_repositories` 关系表，并明确旧入口、旧绑定与历史数据如何兼容迁移。否则 `phase04` 验收会停留在“接口与页面都已写完，但数据库基线和兼容路径未收口”的半完成状态。

## What Changes

- 冻结 `Product / Repository / Binding` 联调环境的可重复建立方式与启动顺序
- 冻结 `products / repositories` 从 `phase02` 只读前提升级为 `phase04` 正式主线的 migration 设计
- 冻结 `product_repositories` 关系表、相关索引与历史绑定兼容边界
- 冻结 `phase04` 专用重置脚本、基线 seed 与 fixture 设计
- 冻结从空状态到首个 `Product`、首个 `Repository` 与首轮三类绑定的冷启动验收路径
- 冻结 `Module Detail` 旧入口兼容跳转、旧 transport 兼容委派与多入口回流验收矩阵
- 明确当前阶段不依赖手工 SQL 建立验收环境，也不额外冻结测试框架或实现工具选型

## Impact

- Affected specs:
  - `phase04_02_product_repository_template_status_read_model`
  - `phase04_03_binding_relation_candidate_scope_context_entry`
  - `phase04_04_product_repository_binding_data_api_error_boundary`
  - `phase04_06_frontend_state_interaction_flow`
  - `phase04_07_backend_module_boundary_interface_grouping`
  - `phase04_08_product_repository_binding_proto_contract`
- Affected code:
  - 预期新增 `database/migrations/0006_product_repository_binding_mainline.sql`
  - 预期新增 `database/scripts/reset_product_repository_mainline.sh`
  - 预期新增 `database/seeds/seed_product_repository_mainline_baseline.sql`
  - 预期更新 `database/seeds/seed_readonly_prereqs.sql`
  - 预期约束 `frontend/src/routes/products/**`、`frontend/src/routes/repositories/**`
  - 预期约束 `backend/internal/productregistry/**`、`backend/internal/repositorybinding/**`

## ADDED Requirements

### Requirement: phase04 联调环境必须可重复建立

系统 SHALL 为 `Product / Repository / Binding` 提供可重复建立的联调环境，使验收可在同一条真实前后端主线上反复执行，而不是依赖手工补库。

#### Scenario: 环境建立入口复用

- **WHEN** 执行 `phase04` 联调验收
- **THEN** 必须继续复用现有单一数据库与脚本主线：`database/scripts/init_db.sh`、`database/scripts/run_seeds.sh`、`database/scripts/reset_module_mainline.sh`
- **AND** 必须新增 `database/scripts/reset_product_repository_mainline.sh` 作为 `phase04` 主线重置入口
- **AND** 不得新建第二个数据库、第二套 `init_db` 入口或第二套 phase04 专用 seed runner

#### Scenario: phase04 联调启动顺序

- **WHEN** 从零建立 `phase04` 联调环境
- **THEN** 启动顺序必须冻结为：`init_db.sh` -> 后端启动并自动执行 migration（含 `0006_product_repository_binding_mainline.sql`）-> `run_seeds.sh` -> `reset_module_mainline.sh` -> `reset_product_repository_mainline.sh` -> 前端启动
- **AND** `run_seeds.sh` 必须发生在 `0006` migration 完成之后，因为 `products / repositories` 的结构已从只读前提升级为正式主线
- **AND** `reset_product_repository_mainline.sh` 必须在 `reset_module_mainline.sh` 之后执行，因为其基线恢复需要依赖 `modules` 基线来重建 `product_modules / module_repositories`
- **AND** 当前阶段不要求 `reset_decision_mainline.sh` 成为 `phase04` 验收前置步骤

### Requirement: products / repositories 主线升级 migration 必须冻结

系统 SHALL 通过新增 migration 将 `products / repositories` 从 `phase02` 只读前提原位升级为 `phase04` 正式主线，并新增 `product_repositories` 关系表，不得通过影子表或双写迁移绕过。

#### Scenario: migration 文件落点与职责

- **WHEN** 后续实现 `phase04` 数据主线升级
- **THEN** migration 文件必须落在 `database/migrations/0006_product_repository_binding_mainline.sql`
- **AND** `products` 必须通过 `ALTER TABLE` 原位新增 `description` 与 `status`
- **AND** `repositories` 必须通过 `ALTER TABLE` 原位新增 `url`、`provider` 与 `status`
- **AND** 必须新增 `product_repositories` 表承接 `Product <-> Repository` 绑定关系
- **AND** 不得新建 `products_v2`、`repositories_v2` 或并行影子绑定表

#### Scenario: Product 主线字段升级

- **WHEN** `0006` migration 升级 `products`
- **THEN** `description` 必须新增为 `TEXT`
- **AND** `status` 必须新增为 `TEXT NOT NULL DEFAULT 'active'`
- **AND** 回填后 `description` 必须进入 `NOT NULL`
- **AND** `status` 必须增加 `CHECK (status IN ('active', 'archived'))`
- **AND** 必须新增列表读取索引 `products(status, created_at DESC)`
- **AND** 不得删除既有 `id / name / created_at`

#### Scenario: Repository 主线字段升级

- **WHEN** `0006` migration 升级 `repositories`
- **THEN** `url` 必须新增为 `TEXT`
- **AND** `provider` 必须新增为 `TEXT`
- **AND** `status` 必须新增为 `TEXT NOT NULL DEFAULT 'active'`
- **AND** 回填后 `url / provider` 必须进入 `NOT NULL`
- **AND** `status` 必须增加 `CHECK (status IN ('active', 'archived'))`
- **AND** 必须新增列表读取索引 `repositories(status, created_at DESC)`
- **AND** 不得删除既有 `id / name / created_at`

#### Scenario: product_repositories 关系表结构冻结

- **WHEN** `0006` migration 新增 `product_repositories`
- **THEN** 表结构至少必须包含 `id / product_id / repository_id / created_at`
- **AND** `product_id` 必须引用 `products(id) ON DELETE RESTRICT`
- **AND** `repository_id` 必须引用 `repositories(id) ON DELETE RESTRICT`
- **AND** `(product_id, repository_id)` 必须唯一
- **AND** 必须新增 `product_id` 与 `repository_id` 方向的读取索引
- **AND** 不得在当前阶段引入额外绑定属性字段

### Requirement: 历史数据兼容回填必须冻结

系统 SHALL 为 `phase02` 遗留的 `products / repositories` 数据提供兼容回填策略，保证历史只读前提记录升级后仍可被 `phase04` 主线读取与绑定。

#### Scenario: Product 历史数据回填

- **WHEN** `0006` migration 执行时 `products` 表中已有历史记录
- **THEN** `description` 必须回填为 `'（历史产品，phase04 升级前无描述）'`
- **AND** `status` 必须回填或默认落为 `active`
- **AND** 回填必须保留原有 `id / name / created_at`
- **AND** 回填语句必须具备幂等性

#### Scenario: Repository 历史数据回填

- **WHEN** `0006` migration 执行时 `repositories` 表中已有历史记录
- **THEN** `url` 必须回填为 `'https://example.com/legacy'`
- **AND** `provider` 必须回填为 `'legacy'`
- **AND** `status` 必须回填或默认落为 `active`
- **AND** 回填必须保留原有 `id / name / created_at`
- **AND** 回填语句必须具备幂等性

#### Scenario: phase02 历史绑定数据兼容

- **WHEN** `0006` migration 执行后
- **THEN** 既有 `product_modules` 与 `module_repositories` 历史数据必须保持可读
- **AND** `phase04` 不得通过重建这些历史关系表来完成 owner 迁移
- **AND** `phase04` 只新增 `product_repositories` 关系，不回写第二套旧绑定表

### Requirement: phase04 重置脚本必须与既有 mainline 脚本同构

系统 SHALL 为 `phase04` 提供与 `reset_module_mainline.sh`、`reset_decision_mainline.sh` 同构的重置脚本，承接 `Product / Repository / Binding` 的清空与基线恢复。

#### Scenario: 脚本落点与模式

- **WHEN** 后续实现 `phase04` 重置脚本
- **THEN** 脚本必须落在 `database/scripts/reset_product_repository_mainline.sh`
- **AND** 必须支持 `--clean-only`、`--restore-only` 与默认模式（清空 + 恢复）
- **AND** 必须复用既有 `resolve_psql` 与环境变量覆盖模式
- **AND** 必须支持 `-h / --help`

#### Scenario: 清空范围冻结

- **WHEN** 执行 `reset_product_repository_mainline.sh --clean-only` 或默认模式
- **THEN** 清空范围必须覆盖 `product_repositories`、`product_modules`、`module_repositories`、`products`、`repositories`
- **AND** 不得清空 `modules / module_releases / decisions / decision_links`
- **AND** 清空必须按以下受控顺序执行，避免触发 `ON DELETE RESTRICT`：先 `product_repositories`，再 `product_modules` 与 `module_repositories`，最后 `products` 与 `repositories`
- **AND** 清空完成后必须输出 `products / repositories / product_repositories / product_modules / module_repositories` 的计数

#### Scenario: DELETE 策略冻结

- **WHEN** 设计 `phase04` 重置脚本的清空方式
- **THEN** 必须继续沿用 `DELETE FROM ...` 的受控清空模式
- **AND** 不得切换为 `TRUNCATE ... RESTART IDENTITY ... CASCADE` 作为默认方案
- **AND** 原因必须明确为：当前主键采用 `UUID` 默认值，不依赖 sequence 重置；同时 `DELETE` 更便于承接现有脚本的受控清空范围与跨阶段一致性

#### Scenario: 前置校验冻结

- **WHEN** 执行 `reset_product_repository_mainline.sh`
- **THEN** 必须校验目标数据库已存在
- **AND** 在 `--restore-only` 与默认模式下，必须校验 `modules` 基线已存在
- **AND** 若 `modules` 基线不存在，必须提示先执行 `reset_module_mainline.sh`

### Requirement: phase04 基线 seed 与 fixture 设计必须冻结

系统 SHALL 冻结 `phase04` 的基线 seed 与 fixture 策略，使联调与验收不再依赖零散手工 SQL。

#### Scenario: 基线 seed 文件落点

- **WHEN** 后续实现 `phase04` 基线数据
- **THEN** seed 文件必须落在 `database/seeds/seed_product_repository_mainline_baseline.sql`
- **AND** 必须通过 `BEGIN / COMMIT` 包裹
- **AND** 必须同时承担“清空 + 恢复”职责，默认脚本直接执行整份 SQL
- **AND** `--restore-only` 模式必须通过跳过清空语句并依赖幂等 `INSERT` 恢复基线

#### Scenario: 基线 Product / Repository 覆盖维度

- **WHEN** 定义 `phase04` 基线 `Product / Repository` 数据
- **THEN** 至少必须包含 `3` 条 `Product` 与 `2` 条 `Repository`
- **AND** `Product` 必须覆盖 `active` 与 `archived` 两种状态
- **AND** `Repository` 必须覆盖 `active` 与 `archived` 两种状态
- **AND** `Repository.provider` 必须至少覆盖两个不同字符串值
- **AND** 所有基线记录都必须通过业务字段查找，不得硬编码 `UUID`

#### Scenario: 基线 seed 与 module 基线的 name 兼容必须冻结

- **WHEN** `seed_product_repository_mainline_baseline.sql` 定义 `products / repositories` 基线记录
- **THEN** 必须保留 `phase02` 历史基线中已存在的既有记录名：`Product A / Product B / Product C / main-repo / mirror-repo`，其中 `Product A` 与 `main-repo` 被 `seed_module_mainline_baseline.sql` 通过 `name` 查找 ID 用于重建 `product_modules / module_repositories`
- **AND** `phase04` 新增的基线 `Product / Repository` 记录只能在这些既有 name 之外扩展，不得替换或重命名既有记录
- **AND** 此约束的目的是保证 `reset_product_repository_mainline.sh` 与 `reset_module_mainline.sh` 可交叉重复执行：`reset_product_repository_mainline.sh` 清空并重建 `products / repositories` 后，后续执行 `reset_module_mainline.sh` 仍能通过 `name` 找到 `product_id / repository_id` 来重建 `product_modules / module_repositories`
- **AND** 若后续阶段需要变更既有 name，必须同步改写 `seed_module_mainline_baseline.sql` 的查找目标，并在对应阶段 spec 中显式冻结
- **AND** 当前阶段不得单方面变更既有 name

#### Scenario: 基线三类绑定覆盖维度

- **WHEN** 定义 `phase04` 基线绑定数据
- **THEN** 必须同时覆盖三类关系：
  - `product_modules`
  - `product_repositories`
  - `module_repositories`
- **AND** 至少包含 `2` 条 `product_modules`
- **AND** 至少包含 `1` 条 `product_repositories`
- **AND** 至少包含 `2` 条 `module_repositories`
- **AND** 必须保留至少 `1` 条“无已绑定仓库的 Product”与至少 `1` 条“无已绑定 Product 的 Repository”，以验证空区块语义
- **AND** `product_modules / product_repositories / module_repositories` 的 INSERT 必须通过 `name` 查找 `product_id / repository_id / module_id`，不得硬编码 `UUID`

#### Scenario: fixture 分层策略冻结

- **WHEN** 讨论 `phase04` 验收所需 fixture
- **THEN** 只允许使用两层 fixture：
  - 基线 fixture：由 `seed_product_repository_mainline_baseline.sql` 提供
  - 临时场景 fixture：由 `--clean-only`、`--restore-only` 与受控 API 操作建立
- **AND** 不得为单个异常路径再新增独立 `fixture SQL` 文件
- **AND** 不得通过手工 SQL 临时改库制造场景

#### Scenario: seed_readonly_prereqs 的兼容定位冻结

- **WHEN** `phase04` 联调环境仍执行 `run_seeds.sh`
- **THEN** `seed_readonly_prereqs.sql` 中的 `products / repositories` 只能继续承担“历史兼容最小前提”角色
- **AND** `phase04` 最终验收状态必须以 `seed_product_repository_mainline_baseline.sql` 恢复后的数据为准
- **AND** 不得再把 `seed_readonly_prereqs.sql` 解释为 `phase04` 正式主线基线

#### Scenario: seed_readonly_prereqs.sql products / repositories seed 字段升级

- **WHEN** `phase04` 执行 `run_seeds.sh`
- **THEN** `seed_readonly_prereqs.sql` 必须将 `products` 的 INSERT 从 `name-only` 升级为 `name + description + status` 完整字段插入
- **AND** `seed_readonly_prereqs.sql` 必须将 `repositories` 的 INSERT 从 `name-only` 升级为 `name + url + provider + status` 完整字段插入
- **AND** 必须保持原有 `name`（`Product A / Product B / Product C / main-repo / mirror-repo`）以兼容 `seed_module_mainline_baseline.sql` 中通过 `name` 查找 ID 的 `product_modules / module_repositories` INSERT
- **AND** `status` 默认落为 `'active'`
- **AND** `description` 必须填入 `'（历史产品，phase04 升级前无描述）'`
- **AND** `url` 必须填入 `'https://example.com/legacy'`
- **AND** `provider` 必须填入 `'legacy'`
- **AND** 必须保持 `ON CONFLICT (name) DO NOTHING` 以保证 `run_seeds.sh` 可重复执行
- **AND** 此升级模式必须对齐 `phase03-12` 已建立的 `decisions` seed 字段升级先例

### Requirement: 冷启动验收路径必须冻结

系统 SHALL 冻结 `phase04` 从空状态到首轮主线完成的冷启动验收路径。

#### Scenario: 空状态到首轮主线闭环

- **WHEN** 执行 `phase04` 冷启动验收
- **THEN** 路径至少必须为：
  1. 执行 `reset_module_mainline.sh`
  2. 执行 `reset_product_repository_mainline.sh --clean-only`
  3. 进入 `Product Registry / List`，验证空状态入口
  4. 创建首个 `Product`
  5. 进入 `Repository Binding / List`，验证空状态入口
  6. 创建首个 `Repository`
  7. 在 `Product Detail` 执行 `BindModuleToProduct`，reread 必须回到 `Product Detail` 反映已绑定模块
  8. 在 `Repository Binding Detail / Workspace` 执行 `BindRepositoryToProduct`，reread 必须回到 `Repository Binding Detail / Workspace` 反映已绑定产品
  9. 在同一 `Repository Binding Detail / Workspace` 执行 `MapModuleToRepository`，reread 必须回到 `Repository Binding Detail / Workspace` 反映已映射模块
  10. 从 `Module Detail` 兼容入口跳转进入正式绑定主入口，完成至少一条绑定写入并 reread 回 canonical owner 页面
- **AND** 三类绑定写入后的 reread 必须分别落到对应 canonical owner 页面，不得只靠 toast 作为成功依据
- **AND** 全路径不得依赖手工 SQL
- **AND** 创建成功后的默认页面流必须承接 `phase04-06` 已冻结的返回与来源上下文规则

### Requirement: 兼容入口与多入口回流验收矩阵必须冻结

系统 SHALL 将 `phase04` 的旧入口兼容与多入口回流矩阵冻结为可执行验收项，避免联调时再临场解释。

#### Scenario: Module Detail 兼容入口主路径

- **WHEN** 用户从 `Module Detail` 发起“绑定到 Product”相关动作
- **THEN** 必须通过兼容跳转进入 `Product` 正式主入口
- **AND** 必须携带 `moduleId / moduleName / fromModuleDetail`
- **AND** 若目标 `Product` 未知，必须先落到 `Product Registry / List` 或 `Product Create`
- **AND** 若目标 `Product` 已知，必须进入对应 `Product Detail`

#### Scenario: Module Detail 到 Repository 主线兼容路径

- **WHEN** 用户从 `Module Detail` 发起“绑定到 Repository”相关动作
- **THEN** 必须通过兼容跳转进入 `Repository` 正式主入口
- **AND** 必须携带 `moduleId / moduleName / fromModuleDetail`
- **AND** 若目标 `Repository` 未知，必须先落到 `Repository Binding / List` 或 `Repository Create`
- **AND** 若目标 `Repository` 已知，必须进入对应 `Repository Binding Detail / Workspace`

#### Scenario: Product Detail 到 Repository 主线兼容路径

- **WHEN** 用户从 `Product Detail` 发起“为当前 Product 绑定 Repository”
- **THEN** 必须进入 `Repository Binding` 正式主入口
- **AND** 必须携带 `productId / productName / fromProductDetail`
- **AND** 目标 `Repository` 未知时先走列表或创建路径，目标已知时直接进入对应工作台

#### Scenario: 多入口回流验收矩阵

- **WHEN** 执行创建、绑定成功后的回流验收
- **THEN** 至少必须覆盖以下矩阵：
  - `fromList` -> 创建成功 -> 详情页 -> 返回列表并恢复原 `queryText / statusFilter`
  - `fromModuleDetail` -> 创建成功 / 绑定成功 -> 对应详情页或工作台 -> 返回原 `Module Detail`
  - `fromProductDetail` -> 仓库创建成功 / 仓库绑定成功 -> `Repository Binding Detail / Workspace` -> 返回原 `Product Detail`
  - `direct-entry` -> 创建成功 / 绑定成功 -> 对应详情页或工作台 -> 返回默认列表参数
- **AND** 不得出现回流后来源上下文丢失、伪造 `fromList` 或回到错误页面的情况

### Requirement: 旧 transport 兼容与异常路径验证要求必须冻结

系统 SHALL 冻结 `phase02` 旧 transport 兼容与 `phase04` 异常路径验证要求，保证迁移期间不出现第二套 owner 或隐性 500。

#### Scenario: 旧 transport 兼容委派

- **WHEN** 迁移窗口内仍保留 `phase02` 的旧候选读取或旧绑定入口
- **THEN** 它们只能作为兼容适配层委派给 `phase04` canonical owner
- **AND** 不得继续拥有独立业务实现
- **AND** 验收必须逐项覆盖以下仍保留的兼容面，不得只验证其中一条即宣告通过：
  - `ProductBindingCandidateRead` 旧 transport 入口必须正确委派到 `Repository Binding` canonical 实现，不得返回过时数据
  - `ModuleBindingWrite` 旧模块中心写入入口必须正确委派到 `Product Registry` 的 `ProductModuleBindingWrite` 或 `Repository Binding` 的 `RepositoryModuleMappingWrite`，不得返回过时数据或独立写入
  - `RepositoryBindingCandidateRead` 若迁移窗口内仍保留旧 transport 入口，不得返回过时数据或承接新的业务实现与消费方
  - `Module Detail` 旧入口兼容跳转必须正确携带 `moduleId / moduleName / fromModuleDetail` 并进入正式主入口，不得在本页直接提交绑定写入
- **AND** 若某类旧入口在实现中被完全移除而非保留为兼容适配层，则该类无需进入验收，但必须在实现阶段 spec 或验收报告中显式记录其已移除

#### Scenario: phase04 最小异常路径验证

- **WHEN** 执行 `phase04` 异常路径验证
- **THEN** 至少必须覆盖以下场景：
  - `CreateProduct / CreateRepository` 必填字段缺失
  - `CreateProduct / CreateRepository` 非法 `status`
  - 三类绑定目标不存在
  - 三类绑定目标非 `active`
  - 三类绑定重复冲突
  - `ProductListRead / RepositoryListRead` 空列表（空状态入口）
  - 三条候选读取为空列表
  - `ProductDetailRead / RepositoryDetailRead` 目标不存在
- **AND** 异常前提必须通过基线 seed 与受控 API 操作建立
- **AND** 不得出现未收口的 `500` 替代业务错误

### Requirement: 当前阶段不冻结的实现工具必须明确

系统 SHALL 明确 `phase04-09` 当前只冻结环境、迁移、脚本、seed 与验收矩阵，不冻结测试框架或具体实现工具。

#### Scenario: 当前阶段非目标

- **WHEN** 讨论 `phase04-09` 的实现边界
- **THEN** 不得提前冻结 `Go ORM / SQL Builder / repository helper`
- **AND** 不得提前冻结前端 E2E 测试框架或 API 测试工具
- **AND** 不得把完整自动化测试平台建设写成当前阶段前置条件

## MODIFIED Requirements

### Requirement: products / repositories 的项目定位

`products / repositories` 在当前阶段 SHALL 不再只被解释为 `phase02` 候选读取前提，而必须被解释为 `phase04` 的正式主线实体。

#### Scenario: 正式主线解释

- **WHEN** `phase04-09` 完成后再讨论 `products / repositories`
- **THEN** 必须以 `phase04` 正式写入、详情读取、绑定关系与验收基线为唯一主线解释
- **AND** `phase02` 的只读前提语义只保留为历史兼容背景

### Requirement: phase04 验收环境建立方式

`phase04` 验收环境 SHALL 从“依赖现成数据库状态或手工补 SQL”修改为“通过 migration + reset script + baseline seed 重复建立”。

#### Scenario: 验收环境收口

- **WHEN** 开始 `phase04` 联调验收
- **THEN** 必须能通过脚本重复建立环境
- **AND** 不得再要求操作者手工插入 `products / repositories / bindings`

## REMOVED Requirements

### Requirement: 依赖手工 SQL 才能完成 phase04 联调

**Reason**: 这会让 `phase04` 无法形成可重复复核的最小主线，也会让旧入口兼容和回流矩阵无法稳定复现。
**Migration**: 改为通过 `0006_product_repository_binding_mainline.sql`、`reset_product_repository_mainline.sh` 与 `seed_product_repository_mainline_baseline.sql` 建立统一验收环境。
