# Phase03-09 Decision Center 联调验收环境与重置基线设计 Spec

## Why

`phase03-02 / 03 / 04` 已冻结 `Decision Center` 的结构化模板、目标范围、数据读写范围与错误语义，`phase03-07 / 08` 已冻结后端模块边界与最小 `.proto` 合同。但当前仍缺少一份联调验收环境与重置基线设计，用来回答"`Decision Center` 联调环境如何可重复建立、`decisions` 表如何从 `phase02` 只读前提原位升级为结构化主线、重置脚本与基线种子的落点与职责是什么、冷启动验收路径如何从空状态走到首条 `Decision` 关联 `Module`、异常路径验证前提如何不依赖手工 SQL"。如果缺少这一层，`phase03-14` 验收会回到 `phase02` 早期"手工补 SQL 才能建立最小联调环境"的补救状态，违反 `shared_baseline §7` 与 `architecture_plan §4.7` 已冻结的前置规划要求。

## What Changes

- 冻结 `Decision Center` 联调环境的可重复建立方式与前置脚本入口
- 冻结 `decisions` 表从 `phase02` 只读前提原位升级为结构化主线的 migration 设计
- 冻结现有示例 `Decision` 数据的兼容回填策略
- 冻结 `Decision Center` 重置脚本落点、职责与清空范围
- 冻结 `Decision Center` 基线种子数据范围与覆盖维度
- 冻结异常路径验证前提与验证要求
- 冻结从空状态到首条 `Decision` 关联 `Module` 的冷启动验收路径
- 明确当前阶段不冻结 Go 数据访问层具体工具、不冻结前端测试框架选型

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `database/migrations/0004_decision_center_mainline.sql`、`database/scripts/reset_decision_mainline.sh`、`database/seeds/seed_decision_mainline_baseline.sql`、`database/seeds/seed_readonly_prereqs.sql`（`decisions` seed 更新）
- Affected code: 预期直接映射到 `phase03-14` 验收的可重复环境建立入口

## ADDED Requirements

### Requirement: 联调环境必须可重复建立

系统 SHALL 为 `Decision Center` 提供可重复建立的联调环境前提，使验收不依赖一次性的临时状态或手工 SQL。本 Requirement 复用 `phase02-12` 已冻结的独立数据库与初始化入口模式，并扩展承接 `Decision Center` 结构化主线。

#### Scenario: 环境初始化入口复用

- **WHEN** 执行 `Decision Center` 联调
- **THEN** 必须复用 `phase02-11` 已冻结的 `database/scripts/init_db.sh` 创建 `PSCO` 独立数据库
- **AND** 必须复用 `database/scripts/run_seeds.sh` 执行只读前提种子（`products / repositories`）
- **AND** 必须复用 `database/scripts/reset_module_mainline.sh` 建立 `Module Registry` 基线（`Decision -> Module` 候选读取依赖 `modules` 基线数据）
- **AND** 不得为 `Decision Center` 新建第二个数据库或第二套 `init_db` 入口

#### Scenario: 启动顺序明确

- **WHEN** 从零建立 `Decision Center` 联调环境
- **THEN** 启动顺序必须为：`init_db.sh` -> 后端启动（自动运行 migration，含 `0004_decision_center_mainline`）-> `run_seeds.sh`（只读前提）-> `reset_module_mainline.sh` -> `reset_decision_mainline.sh` -> 前端启动
- **AND** `run_seeds.sh` 必须在后端 migration 完成后执行，因为 `seed_readonly_prereqs.sql` 依赖 `0002_readonly_prereqs` 已创建的 `products / repositories / decisions` 表与 `0004` 已添加的 `decisions` 结构化字段
- **AND** `reset_decision_mainline.sh` 必须在 `reset_module_mainline.sh` 之后执行，因为 `decision_links` 依赖 `modules` 基线数据
- **AND** 不得颠倒 `reset_module_mainline` 与 `reset_decision_mainline` 的顺序

### Requirement: decisions 表原位升级 migration 必须冻结

系统 SHALL 将 `decisions` 表从 `phase02` 只读前提（仅 `id / title / created_at`）原位升级为 `phase03` 结构化主线，通过新增 migration 完成，不新建替代表，不临时双写。

#### Scenario: migration 文件落点与职责

- **WHEN** 后续实现 `decisions` 表原位升级
- **THEN** migration 文件必须落在 `database/migrations/0004_decision_center_mainline.sql`
- **AND** 该 migration 必须通过 `ALTER TABLE decisions ADD COLUMN` 原位添加结构化字段，不得新建 `decisions_v2` 或替代表
- **AND** 必须添加以下字段并最终达到 `NOT NULL` 约束：`context` / `problem` / `alternatives TEXT[]` / `choice` / `reason` / `impact` / `status`
- **AND** 对于无 `DEFAULT` 的必填字段（`context / problem / choice / reason`），必须按"先 `ADD COLUMN` 允许 `NULL` -> 回填 -> `SET NOT NULL`"三步流程添加，避免在已有数据时因 `NOT NULL` 无默认值而失败
- **AND** 对于有 `DEFAULT` 的字段（`alternatives TEXT[] NOT NULL DEFAULT '{}'` / `impact TEXT NOT NULL DEFAULT ''` / `status TEXT NOT NULL DEFAULT 'proposed'`），可直接 `ADD COLUMN ... NOT NULL DEFAULT ...`
- **AND** 必须添加 `status` 的 `CHECK` 约束：`CHECK (status IN ('proposed', 'active', 'superseded', 'archived'))`
- **AND** 必须为列表读取性能添加索引：`CREATE INDEX idx_decisions_status_created_at ON decisions (status, created_at DESC)`
- **AND** 该 migration 不得删除原有 `title / created_at` 字段，不得破坏既有 `decision_links` 的外键引用

#### Scenario: alternatives 存储方式冻结

- **WHEN** 在 `decisions` 表中存储 `alternatives` 字段
- **THEN** 必须使用 `TEXT[]`（PostgreSQL 原生数组）存储，对齐 `phase03-02` "按顺序保留的文本条目集合"语义
- **AND** 必须使用 `DEFAULT '{}'` 保证空数组语义，不得使用 `NULL`
- **AND** 不得使用 `JSONB` 存储 alternatives，避免引入与 `.proto` `repeated string` 不必要的序列化层

#### Scenario: 必填字段与空字符串约束

- **WHEN** 定义 `decisions` 表的结构化字段约束
- **THEN** `title / context / problem / choice / reason / status` 必须为 `NOT NULL`
- **AND** `alternatives` 必须为 `NOT NULL DEFAULT '{}'`
- **AND** `impact` 必须为 `NOT NULL DEFAULT ''`（对齐 `phase03-02` 创建可选语义）
- **AND** 空字符串不得视为合法必填值的校验由后端 service 层承接，不在 migration 层通过 `CHECK (length(trim(field)) > 0)` 冻结，避免迁移期回填复杂度

### Requirement: 现有示例 Decision 数据兼容回填必须冻结

系统 SHALL 在 `0004` migration 中完成现有示例 `Decision` 数据的兼容回填，保证 `phase02` 中已存在的 `title-only` 数据在升级后仍可正常读取与展示，不依赖手工 SQL 修补。

#### Scenario: 兼容回填策略

- **WHEN** `0004` migration 执行时 `decisions` 表已有 `phase02` seed 插入的 `title-only` 数据
- **THEN** migration 必须通过 `UPDATE decisions SET ... WHERE <新字段> IS NULL` 完成回填
- **AND** 回填必须保留原有 `title / created_at`，不得覆盖
- **AND** `context / problem / choice / reason` 必须回填为明确的占位文本（如 `（历史决策，phase03 升级前无结构化上下文）`），不得回填为空字符串（违反必填非空校验语义）
- **AND** `alternatives` 必须回填为 `'{}'`（空数组）
- **AND** `impact` 必须回填为 `''`（对齐可选语义）
- **AND** `status` 必须回填为 `'proposed'`（对齐 `phase03-02` 默认状态）
- **AND** 回填 `UPDATE` 仅对无 `DEFAULT` 的必填字段（`context / problem / choice / reason`）实际生效；`DEFAULT` 字段（`alternatives / impact / status`）由 `ADD COLUMN ... NOT NULL DEFAULT ...` 在加列时自动填充到已有行，回填 `UPDATE` 对它们为 `no-op`，但保留在回填语句中不影响幂等性
- **AND** 回填必须可重复执行（`WHERE` 条件保证幂等）

#### Scenario: 既有 decision_links 兼容性

- **WHEN** `0004` migration 执行后
- **THEN** 既有 `decision_links` 必须仍可正常读取
- **AND** `decision_links.decision_id` 外键引用不得因 migration 失效
- **AND** `Module Detail` 页面中 `Decision` 入口展示必须仍能读取到原有 `decision_links` 数据

### Requirement: Decision Center 重置脚本必须与 module_mainline 同构

系统 SHALL 为 `Decision Center` 提供与 `reset_module_mainline.sh` 同构的重置脚本，支持可重复执行的"清空 -> 恢复基线"入口，使验收不依赖手工 SQL。

#### Scenario: 重置脚本落点与模式

- **WHEN** 后续实现 `Decision Center` 重置脚本
- **THEN** 脚本必须落在 `database/scripts/reset_decision_mainline.sh`
- **AND** 必须支持三种模式：`--clean-only`（仅清空）、`--restore-only`（仅恢复）、默认（清空 + 恢复）
- **AND** 必须复用 `reset_module_mainline.sh` 的 `resolve_psql` 模式（宿主机 psql -> docker exec -> podman exec 自动检测）
- **AND** 必须复用相同的环境变量覆盖参数（`PG_HOST / PG_PORT / PG_USER / PSCO_DB / PG_CONTAINER / SEEDS_DIR`）
- **AND** 必须支持 `-h / --help` 输出用法说明

#### Scenario: 清空范围冻结

- **WHEN** 执行 `reset_decision_mainline.sh --clean-only` 或默认模式
- **THEN** 清空范围必须是 `DELETE FROM decisions`，依赖 `decision_links.decision_id` 的 `ON DELETE CASCADE` 级联清空 `decision_links`
- **AND** 不得清空 `modules / products / repositories` 表
- **AND** 不得清空 `module_releases / product_modules / module_repositories` 表
- **AND** 清空后必须输出当前 `decisions` 与 `decision_links` 计数以供确认

#### Scenario: 前置校验

- **WHEN** 执行 `reset_decision_mainline.sh`
- **THEN** 必须校验 `PSCO` 数据库已存在（否则提示先执行 `init_db.sh`）
- **AND** 在 `--restore-only` 与默认模式下，必须校验 `modules` 基线数据已存在（`SELECT COUNT(*) FROM modules >= 1`），否则提示先执行 `reset_module_mainline.sh`
- **AND** 不得在 `modules` 为空时恢复 `decision` 基线，避免 `decision_links` 外键引用失败

### Requirement: Decision Center 基线种子数据范围必须冻结

系统 SHALL 为 `Decision Center` 冻结基线种子数据范围，覆盖 `phase03-14` 验收所需的最小展示、筛选、详情与关联场景。

#### Scenario: 基线 seed 文件落点

- **WHEN** 后续实现 `Decision Center` 基线种子
- **THEN** seed 文件必须落在 `database/seeds/seed_decision_mainline_baseline.sql`
- **AND** 该文件必须同时承担"清空 + 恢复"职责（与 `seed_module_mainline_baseline.sql` 模式一致），开头包含 `DELETE FROM decisions`，后续 `INSERT` 使用 `ON CONFLICT DO NOTHING`
- **AND** 必须通过 `BEGIN / COMMIT` 事务包裹

#### Scenario: 基线 Decision 数据覆盖维度

- **WHEN** 定义 `Decision Center` 基线数据
- **THEN** 必须至少包含 `3` 条结构化 `Decision`，覆盖以下 `status` 维度：
  - `1` 条 `proposed`
  - `1` 条 `active`
  - `1` 条 `archived` 或 `superseded`
- **AND** 其中 `1` 条必须保留 `phase02` 原有 `title`（`关于 auth-service 技术选型的决策`），并补全结构化字段，以保证 `phase02` `decision_links` 基线兼容
- **AND** 至少 `1` 条 `Decision` 必须包含 `alternatives` 数组（`2` 个以上条目），验证数组保序语义
- **AND** 至少 `1` 条 `Decision` 的 `alternatives` 必须为空数组，验证空数组语义
- **AND** 至少 `1` 条 `Decision` 的 `impact` 必须为空字符串，验证可选字段空值语义

#### Scenario: 基线 decision_links 数据

- **WHEN** 定义 `Decision Center` 基线 `decision_links`
- **THEN** 必须至少包含 `2` 条 `decision_links`，关联到 `reset_module_mainline` 已建立的 `modules` 基线（如 `auth-service`）
- **AND** 其中 `1` 条必须复用 `phase02` 原有的 `关于 auth-service 技术选型的决策` -> `auth-service` 关联，保证兼容性
- **AND** 至少 `1` 条 `Decision` 必须没有任何 `decision_links`，验证 `link_count = 0` 与 `linked_module_summary = ''` 的空值语义
- **AND** `decision_links` 的 `INSERT` 必须通过 `module_name` 与 `decision title` 查找 `ID`，不得硬编码 `UUID`
- **AND** `seed_decision_mainline_baseline.sql` 承担 `decision_links` 的最终基线职责；`seed_module_mainline_baseline.sql` 中既有的 `decision_link` 插入在 `reset_decision_mainline.sh` 执行时会被 `DELETE FROM decisions` 级联删除后由 `decision` 基线 `seed` 重建，最终状态以 `seed_decision_mainline_baseline.sql` 为准

#### Scenario: seed_readonly_prereqs.sql decisions seed 更新

- **WHEN** `phase03` 升级 `decisions` 表后执行 `seed_readonly_prereqs.sql`
- **THEN** 该文件中的 `decisions` seed 必须从 `title-only` 更新为结构化字段插入
- **AND** 必须保持原有 `title`（`关于 auth-service 技术选型的决策`）以兼容 `phase02` `decision_links`
- **AND** 必须补全 `context / problem / choice / reason / status` 必填字段
- **AND** `alternatives` 必须设为 `'{}'`，`impact` 必须设为 `''`
- **AND** 该更新必须保证 `run_seeds.sh` 可重复执行（`ON CONFLICT` 或守卫块）

### Requirement: 异常路径验证前提与要求必须冻结

系统 SHALL 为 `Decision Center` 冻结异常路径验证要求，使 `phase03-14` 验收能覆盖关键异常场景，且异常路径验证不依赖手工 SQL 补数据。

#### Scenario: 异常路径验证要求

- **WHEN** 执行 `Decision Center` 异常路径验证
- **THEN** 必须至少覆盖以下异常路径（对齐 `phase03-04` 已冻结的全部错误语义）：
  - `RecordDecision` 必填字段缺失 -> 返回校验失败
  - `RecordDecision` 字段值非法（含 `alternatives` 条目空白、非法 `status` 值）-> 返回校验失败
  - `LinkDecisionToTarget` 目标类型越界（非 `MODULE`）-> 返回校验失败
  - `LinkDecisionToTarget` 目标 `Decision` 或 `Module` 不存在 -> 返回资源不存在语义
  - `LinkDecisionToTarget` 重复关联 -> 返回重复冲突语义
  - `DecisionModuleCandidateRead` 无可关联候选 -> 返回空列表语义（非错误）
  - `DecisionDetailRead` 不存在的 `decision_id` -> 返回资源不存在语义
- **AND** 异常路径验证通过 `API` 层测试触发，不通过 `seed` 异常数据
- **AND** 不得出现 `500` 级未收口错误替代业务错误

#### Scenario: 异常路径验证前提数据

- **WHEN** 异常路径验证需要前提数据
- **THEN** "重复关联"验证的前提（已有 `decision_links`）由 `seed_decision_mainline_baseline.sql` 基线提供
- **AND** "候选排除已关联目标"验证的前提（已有 `decision_links`）由基线提供
- **AND** "无可关联候选"验证通过创建一个已关联全部 `active Module` 的 `Decision` 触发，不需要 `seed` 额外数据
- **AND** 不得为异常路径验证新建单独的 `fixture SQL` 文件，异常前提由基线 `seed` 与 `API` 操作共同建立

### Requirement: 冷启动验收路径必须冻结

系统 SHALL 冻结从空状态到首条 `Decision` 关联 `Module` 的冷启动验收路径，对齐 `shared_baseline §7` 冷启动基线。

#### Scenario: 冷启动验收路径

- **WHEN** 执行 `Decision Center` 冷启动验收
- **THEN** 验收路径必须为：
  1. 执行 `reset_decision_mainline.sh --clean-only`，使 `decisions` 进入空状态
  2. 进入 `Decision Center / List`，验证空状态入口展示
  3. 从空状态进入 `Decision Create`
  4. 填写最小结构化模板，提交 `RecordDecision`
  5. 验证回流到 `DecisionDetailPage`，详情页显示新建 `Decision` 核心字段
  6. 在 `DecisionDetailPage` 触发 `DecisionModuleCandidateRead`，验证候选列表展示
  7. 选择一个 `Module`，执行 `LinkDecisionToTarget`
  8. 验证详情页重新读取并显示已关联 `Module` 结果
- **AND** 该路径必须在不执行任何手工 `SQL` 的前提下完成
- **AND** 该路径的前提是 `reset_module_mainline.sh` 已建立 `modules` 基线（候选读取需要 `modules`）

#### Scenario: 空状态验收前提

- **WHEN** `decisions` 表为空
- **THEN** `Decision Center / List` 必须显示空状态入口，对齐 `phase03-05 / 06` 前端空状态设计
- **AND** 空状态主动作必须直接进入 `Decision Create`
- **AND** 不得因空状态返回接口错误

### Requirement: 当前阶段不冻结的实现工具

系统 SHALL 明确当前阶段不冻结 `Go` 数据访问层具体工具、前端测试框架选型与 `API` 测试工具选型，只冻结验收环境的建立方式、脚本落点、基线数据范围与异常路径验证要求。

#### Scenario: 不冻结的实现工具

- **WHEN** 当前 spec 讨论联调验收环境
- **THEN** 不得提前冻结 `Go ORM / SQL Builder / Repository 模板`
- **AND** 不得提前冻结前端测试框架（如 `Vitest / Playwright`）
- **AND** 不得提前冻结 `API` 测试工具（如 `curl / httpie / Postman`）
- **AND** 只冻结 `SQL` migration、`bash` 重置脚本、`SQL` 基线 seed 的落点与职责

## MODIFIED Requirements

### Requirement: phase03 验收环境前置规划

`phase03` 的验收环境 SHALL 在 `phase03-09` 中前置规划，不得后移到 `phase03-14` 验收时再手工补 `SQL`。

#### Scenario: 前置规划收口

- **WHEN** `phase03-09` 完成验收
- **THEN** `phase03-14` 验收必须能直接复用本 spec 冻结的脚本、seed 与验收路径
- **AND** 不得在 `phase03-14` 验收时新建临时 `SQL` 或临时脚本
- **AND** 不得在 `phase03-14` 验收时发现缺少重置入口而阻塞

## REMOVED Requirements

### Requirement: 依赖手工 SQL 建立验收环境

**Reason**: `phase02-12` 已验证可重复重置脚本模式可行，`phase03` 必须延续该模式，不得回到手工 `SQL` 补救状态。
**Migration**: 改为通过 `reset_decision_mainline.sh` + `seed_decision_mainline_baseline.sql` 建立可重复验收环境，对齐 `shared_baseline §7` "当前阶段验收不得依赖手工补 SQL"。
