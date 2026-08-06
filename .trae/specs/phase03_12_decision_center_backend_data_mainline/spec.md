# Phase03-12 Decision Center 后端与数据主线 Spec

## Why

`phase03-01 ~ 11` 已冻结 `Decision Center` 的页面边界、数据模型、错误语义、后端模块边界、`.proto` 合同主线与联调验收环境设计。`phase03-11` 已将 `.proto` 合同正式落地为仓库唯一合同源。但仓库当前仍缺少 `Decision Center` 的实际后端读写接口、`decisions` 表从 `phase02` 只读前提原位升级为结构化主线的 migration、基线 seed 与可重复重置脚本，后续 `phase03-13` 前端主线与 `phase03-14` 联调验收无法直接消费。

因此，`phase03-12` 必须实现 `Decision Center` 所需的最小后端读写接口、数据主线 migration、基线 seed 与重置脚本，使 `Decision Center` 后端主线可运行、数据主线与已冻结边界一致、联调环境可重复建立。

## What Changes

- 实现 `backend/internal/decisioncenter/` 后端模块的全部文件（handler / service / repository / candidate / 支撑文件）
- 实现 `database/migrations/0004_decision_center_mainline.sql`，将 `decisions` 表从 `phase02` 只读前提原位升级为结构化主线
- 实现 `database/migrations/0005_decision_source_context.sql`，为 `decisions` 表新增 `source_module_id` 字段以承接入口上下文来源（`phase03-10 §5.11`“持续到正式关联完成”要求跨刷新持久化）
- 实现 `database/seeds/seed_decision_mainline_baseline.sql` 基线种子（含 `decisions` + `decision_links`，基线 `Decision 1` 带 `source_module_id` 演示来源上下文）
- 更新 `database/seeds/seed_readonly_prereqs.sql` 中 `decisions` seed 从 `title-only` 升级为结构化字段插入
- 实现 `database/scripts/reset_decision_mainline.sh` 重置脚本
- 在 `backend/internal/platform/router.go` 中新增 `mountDecisionCenter` 装配函数
- 在 `backend/internal/platform/server.go` 中将 `Decision Center` 路由挂载到 `/api`
- 在 `.proto` 合同源 `CreateDecisionRequest` 中新增 `source_module_id` 字段（编号 `9`，非 breaking change），并在 `service` / `repository` 层承接来源校验与持久化读写，`DecisionDetailRead` 通过 `source_context` 持久化返回来源 `Module` 标识
- **BREAKING**：后续 `Decision Center` 的前端实现与联调验收必须直接消费本阶段交付的后端 API 与数据主线，不得继续依赖手工 SQL 或 mock 数据

## Impact

- Affected specs:
  - `phase03_04_decision_data_api_error_boundary`
  - `phase03_07_backend_module_boundary_interface_grouping`
  - `phase03_09_decision_center_integration_baseline`
  - `phase03_10_decision_center_formal_spec`
  - `phase03_11_decision_center_proto_mainline`
  - `phase03_13` 前端主线
  - `phase03-14` 联调与验收
- Affected code:
  - `backend/internal/decisioncenter/` （新增整模块）
  - `backend/internal/platform/router.go` （新增 `mountDecisionCenter`）
  - `backend/internal/platform/server.go` （挂载 `Decision Center` 路由）
  - `database/migrations/0004_decision_center_mainline.sql` （新增）
  - `database/migrations/0005_decision_source_context.sql` （新增，承接 `phase03-10 §5.11` 入口上下文来源持久化）
  - `database/seeds/seed_decision_mainline_baseline.sql` （新增，基线 `Decision 1` 带 `source_module_id`）
  - `database/seeds/seed_readonly_prereqs.sql` （更新 `decisions` seed）
  - `database/scripts/reset_decision_mainline.sh` （新增）
  - `proto/psco/decision_center/v1/decision_center.proto` （`CreateDecisionRequest` 新增 `source_module_id` 字段编号 `9`）

## ADDED Requirements

### Requirement: 后端模块文件落点与分层必须与 phase03-10 §10 一致

系统 SHALL 在 `backend/internal/decisioncenter/` 下实现 `Decision Center` 后端模块，文件落点、分层语义与支撑文件组织方式必须与 `phase03-10` §10.2 / §10.3 / §10.6 冻结结论完全一致，并与现有 `moduleregistry` 模块保持同构。

#### Scenario: 文件落点冻结

- **WHEN** 实现 `Decision Center` 后端模块
- **THEN** 必须按以下落点创建文件：
  - `handler/query_handler.go` — 读组入口：`DecisionListRead` + `DecisionDetailRead` + `DecisionModuleCandidateRead`
  - `handler/command_handler.go` — 写组入口：`DecisionWrite` + `DecisionLinkWrite`
  - `handler/response.go` — HTTP 协议层共享工具（JSON 编解码、错误到状态码映射）
  - `service/query_service.go` — 读组编排（列表 + 详情 + 候选读取）
  - `service/command_service.go` — 写组编排（创建 + 关联）
  - `repository/decision_store.go` — `decisions` 表读写
  - `repository/link_store.go` — `decision_links` 表读写
  - `candidate/module_candidate_read.go` — `ModuleCandidateRead`（跨模块候选读取，由 `Decision Center` 拥有）
  - `errors.go` — 业务错误哨兵值
  - `types.go` — 跨层共享 API 消息结构（DTO、枚举、请求/响应类型、列表查询参数）
  - `validate.go` — 输入校验辅助
- **AND** 读组与写组必须在各自单文件内编排，不拆 `list_service.go` / `detail_service.go` / `candidate_service.go` 或两个独立写 service 文件
- **AND** 支撑文件（`errors.go` / `types.go` / `validate.go` / `handler/response.go`）不得散落到 `handler/` 或 `service/` 内部

#### Scenario: 分层语义冻结

- **WHEN** 实现各层职责
- **THEN** `handler/` 只负责承接 HTTP 请求与返回结果
- **AND** `service/` 只负责动作语义、校验顺序与跨连接口编排
- **AND** `repository/` 只负责 `decisions` / `decision_links` 表持久化与读取
- **AND** `candidate/` 只负责跨模块 `Module` 候选读取
- **AND** `service/` 层不得直接写跨模块 SQL

### Requirement: types.go 必须从 .proto 单向承接语义

系统 SHALL 让 `types.go` 中的 DTO 结构体从 `phase03-11` 已落地的 `.proto` 合同源派生或显式对齐，不得形成第二套合同源。

#### Scenario: 枚举与核心对象类型定义

- **WHEN** 定义 `types.go` 中的类型
- **THEN** `DecisionStatus` 必须定义为 `string` 类型，常量值为 `proposed` / `active` / `superseded` / `archived`，对齐 `.proto` `DecisionStatus` 枚举
- **AND** `DecisionLinkTargetType` 必须定义为 `string` 类型，常量值为 `module`，对齐 `.proto` `DecisionLinkTargetType` 枚举
- **AND** `Decision` 结构体必须覆盖 `id / title / context / problem / alternatives / choice / reason / impact / status / created_at`，对齐 `.proto` `Decision` 消息
- **AND** `alternatives` 必须建模为 `[]string`，对齐 `.proto` `repeated string`
- **AND** `created_at` 必须使用 `time.Time`，对齐 `.proto` `google.protobuf.Timestamp`
- **AND** `DecisionListItem` 必须覆盖 `id / title / status / created_at / link_count / linked_module_summary`，对齐 `.proto` `DecisionListItem`
- **AND** `LinkedModule` 必须覆盖 `module_id / module_name`
- **AND** `SourceContext` 必须覆盖 `source_module_id / source_module_name`
- **AND** `DecisionDetail` 必须覆盖 `decision / linked_modules / source_context`
- **AND** `DecisionModuleCandidate` 必须覆盖 `module_id / module_name / status`，`status` 类型为 `moduleregistry.ModuleStatus`（跨包复用，不重定义本地等价枚举）

#### Scenario: 请求体与查询参数定义

- **WHEN** 定义写组请求体与列表查询参数
- **THEN** `CreateDecisionRequest` 必须覆盖 `title / context / problem / alternatives / choice / reason / impact / status / source_module_id`
- **AND** `source_module_id` 为可选字段（`phase03-10 §5.11` 入口上下文来源 `Module` 标识），空字符串表示无来源上下文，不参与 `§5.5` 结构化模板字段冻结
- **AND** `LinkDecisionToTargetRequest` 必须覆盖 `target_type / module_id`（`decision_id` 由 URL 路径参数承接，不放在请求体）
- **AND** `ListQuery` 必须覆盖 `QueryText` 与 `StatusFilter`，与 `moduleregistry.ListQuery` 同构

### Requirement: errors.go 必须覆盖 phase03-04 冻结的全部错误语义

系统 SHALL 在 `errors.go` 中定义覆盖 `phase03-04` 全部错误语义的哨兵值，并通过 `handler/response.go` 映射为 HTTP 状态码。

#### Scenario: 错误哨兵值冻结

- **WHEN** 定义 `errors.go`
- **THEN** 必须定义以下哨兵错误：
  - `ErrDecisionNotFound` — `Decision` 不存在（资源不存在，404）
  - `ErrModuleNotFound` — `Module` 不存在（资源不存在，404）
  - `ErrDuplicateLink` — 重复关联（重复冲突，409）
  - `ErrInvalidInput` — 必填字段缺失或空白（校验失败，400）
  - `ErrInvalidStatus` — 非法 `status` 值（校验失败，400）
  - `ErrInvalidTargetType` — 目标类型越界，非 `MODULE`（校验失败，400）
  - `ErrInvalidAlternatives` — `alternatives` 条目空白（校验失败，400）

#### Scenario: HTTP 状态码映射

- **WHEN** `handler/response.go` 中的 `writeError` 映射错误到 HTTP 状态码
- **THEN** `ErrDecisionNotFound` / `ErrModuleNotFound` → 404
- **AND** `ErrDuplicateLink` → 409
- **AND** `ErrInvalidInput` / `ErrInvalidStatus` / `ErrInvalidTargetType` / `ErrInvalidAlternatives` → 400
- **AND** 其他未收口错误 → 500
- **AND** 不得出现 500 级未收口错误替代业务错误

### Requirement: RecordDecision 创建写入必须执行完整校验

系统 SHALL 在 `service/command_service.go` 的 `CreateDecision` 中执行 `phase03-04` 冻结的完整校验顺序，不得跳过或降级。

#### Scenario: 创建校验顺序

- **WHEN** 执行 `RecordDecision`
- **THEN** 校验顺序必须为：
  1. 必填字段非空校验（`title / context / problem / choice / reason / status`），去首尾空白后不得为空字符串
  2. `status` 取值合法性校验（`proposed / active / superseded / archived`）
  3. `alternatives` 条目校验（每个条目去首尾空白后不得为空字符串）
  4. `source_module_id` 可选来源校验（`phase03-10 §5.11`）：非空时校验 `Module` 存在性（跨模块只读校验由 `candidate` 子包承接），不存在返回 `ErrModuleNotFound`；空字符串表示无来源上下文，跳过校验
- **AND** 校验失败必须返回对应的 `Err*` 哨兵错误
- **AND** 不得降级为模糊通用错误
- **AND** 成功后必须返回新建 `Decision` 标识（`decision_id`），支持前端回流到 `DecisionDetailPage`

#### Scenario: alternatives 写入语义

- **WHEN** 写入 `alternatives` 字段
- **THEN** 必须按输入顺序保留
- **AND** 空数组必须写入 `'{}'`，不得写入 `NULL`
- **AND** 不得引入嵌套对象结构

### Requirement: LinkDecisionToTarget 关联写入必须执行完整校验

系统 SHALL 在 `service/command_service.go` 的 `LinkDecisionToTarget` 中执行 `phase03-04` 冻结的完整校验顺序。

#### Scenario: 关联校验顺序

- **WHEN** 执行 `LinkDecisionToTarget`
- **THEN** 校验顺序必须为：
  1. `target_type` 取值合法性（当前阶段只允许 `module`），越界返回 `ErrInvalidTargetType`
  2. `Decision` 存在性校验，不存在返回 `ErrDecisionNotFound`
  3. `Module` 存在性校验，不存在返回 `ErrModuleNotFound`
  4. 重复关联检测，已存在返回 `ErrDuplicateLink`
- **AND** 成功后必须返回空响应（无返回体），前端通过 `DecisionDetailRead` 重新读取
- **AND** 不得返回 link 结果或 detail reread 标识

### Requirement: 列表读取必须实现 link_count 与 linked_module_summary 计算口径

系统 SHALL 在 `service/query_service.go` 的 `ListDecisions` 中实现 `phase03-10` §5.9 冻结的 `link_count` 与 `linked_module_summary` 计算口径。

#### Scenario: link_count 计算口径

- **WHEN** 列表读取返回 `DecisionListItem`
- **THEN** `link_count` 必须仅统计 `decision_links` 中已建立的 `Decision -> Module` 有效关联数
- **AND** 当 `Decision` 无关联时，`link_count` 必须返回 `0`

#### Scenario: linked_module_summary 计算口径

- **WHEN** 列表读取返回 `DecisionListItem`
- **THEN** `linked_module_summary` 必须仅基于已关联 `Module` 生成，不混入 `Product / Repository`
- **AND** 必须按 `module_name` 升序取前 `3` 个名称
- **AND** 若超出 `3` 个，必须在摘要末尾附加 `+N`（如 `moduleA, moduleB, moduleC +2`）
- **AND** 当 `Decision` 无关联时，`linked_module_summary` 必须返回空字符串，不返回 `null`

### Requirement: 候选读取必须实现排序与排除规则

系统 SHALL 在 `candidate/module_candidate_read.go` 中实现 `phase03-10` §5.10 冻结的候选读取排序与排除规则。

#### Scenario: 候选范围与排序

- **WHEN** 执行 `DecisionModuleCandidateRead`
- **THEN** 候选来源必须为当前已存在的 `modules`
- **AND** 候选范围必须同时覆盖 `active` 与 `archived` 的 `Module`
- **AND** 排序必须采用 `status(active 优先) -> module_name 升序`
- **AND** 已建立 `Decision -> Module` 关联的目标不得再次出现在候选中
- **AND** `status` 必须复用 `moduleregistry.ModuleStatus` 类型，不重定义本地等价枚举

#### Scenario: 空候选结果语义

- **WHEN** 无可关联候选
- **THEN** 必须返回空列表（`[]`），不得返回 `null`
- **AND** 不得将空结果误报为接口错误或资源不存在

### Requirement: ModuleCandidateRead 接口必须由 candidate 子包拥有

系统 SHALL 将 `ModuleCandidateRead` 的接口与实现由 `Decision Center` 的 `candidate/` 子包自己定义和拥有，`Module Registry` 不为 `Decision Center` 暴露专门的服务契约。

#### Scenario: 接口拥有者与接线

- **WHEN** 实现 `ModuleCandidateRead`
- **THEN** 接口与实现必须在 `candidate/module_candidate_read.go` 中定义
- **AND** `candidate/` 子包通过独立 Read 接口隔离，`service/` 层不得直接写跨模块 SQL
- **AND** 具体接线（构造与注入）必须在应用装配点（`platform/router.go`）完成
- **AND** 不得在 `service/` 或 `handler/` 内部自行构造 `ModuleCandidateRead`

### Requirement: source_context 入口上下文来源必须持久化承接

系统 SHALL 通过 `decisions.source_module_id` 字段承接 `phase03-10 §5.11` 冻结的入口上下文来源 `Module` 标识，支持“持续到用户完成正式 `LinkDecisionToTarget`”的跨刷新承接语义（当前阶段不提供“主动放弃关联”出口，`source_context` 作为入口历史记录保留），`DecisionDetailRead` 必须通过 `source_context` 持久化返回来源 `Module` 标识，不得固化为空值或仅由前端路由状态承接。

#### Scenario: source_module_id 字段持久化

- **WHEN** 从 `Module Detail` 带上下文进入 `Decision Create` 并提交
- **THEN** `CreateDecisionRequest.source_module_id` 必须被持久化到 `decisions.source_module_id` 字段
- **AND** `source_module_id` 必须定义为 `UUID REFERENCES modules(id) ON DELETE SET NULL`（`Module` 被删除时来源标识自动置空，不级联删除 `Decision`）
- **AND** `source_module_id` 为可选字段（无来源时为 `NULL`），不参与 `§5.5` 结构化模板字段冻结
- **AND** 该字段必须通过独立 migration 文件 `database/migrations/0005_decision_source_context.sql` 添加，幂等执行（`ADD COLUMN IF NOT EXISTS` + `CREATE INDEX IF NOT EXISTS`）

#### Scenario: DecisionDetailRead 持久化返回 source_context

- **WHEN** 读取 `Decision` 详情
- **THEN** `DecisionDetailRead` 必须通过 `source_context` 返回来源 `Module` 标识（`source_module_id` + `source_module_name`）
- **AND** `source_module_name` 必须通过 `LEFT JOIN modules` 获取，不得在 `decisions` 表中冗余存储
- **AND** `source_module_id` 为 `NULL` 时，`source_context.source_module_id` 与 `source_context.source_module_name` 必须返回空字符串
- **AND** 入口上下文中的预填 `Module` 在正式 `LinkDecisionToTarget` 写入前不得计入 `linked_modules`

#### Scenario: .proto 合同源对齐

- **WHEN** 在 `.proto` 合同源中定义 `source_module_id`
- **THEN** `CreateDecisionRequest` 必须包含 `source_module_id` 字段（编号 `9`，`string` 类型）
- **AND** 该字段添加必须为非 breaking change（新增字段编号，不修改既有字段编号或语义）
- **AND** `buf lint` / `buf breaking` 必须通过

### Requirement: decisions 表原位升级 migration 必须与 phase03-09 一致

系统 SHALL 通过 `database/migrations/0004_decision_center_mainline.sql` 将 `decisions` 表从 `phase02` 只读前提原位升级为结构化主线，migration 设计必须与 `phase03-09` 冻结结论完全一致。

#### Scenario: migration 文件落点与字段升级

- **WHEN** 实现 `decisions` 表原位升级
- **THEN** migration 文件必须落在 `database/migrations/0004_decision_center_mainline.sql`
- **AND** 必须通过 `ALTER TABLE decisions ADD COLUMN` 原位添加结构化字段，不得新建替代表
- **AND** 对于无 `DEFAULT` 的必填字段（`context / problem / choice / reason`），必须按"先 `ADD COLUMN` 允许 `NULL` -> 回填 -> `SET NOT NULL`"三步流程添加
- **AND** 对于有 `DEFAULT` 的字段（`alternatives TEXT[] NOT NULL DEFAULT '{}'` / `impact TEXT NOT NULL DEFAULT ''` / `status TEXT NOT NULL DEFAULT 'proposed'`），可直接 `ADD COLUMN ... NOT NULL DEFAULT ...`
- **AND** 必须添加 `status` 的 `CHECK` 约束：`CHECK (status IN ('proposed', 'active', 'superseded', 'archived'))`
- **AND** 必须为列表读取性能添加索引：`CREATE INDEX idx_decisions_status_created_at ON decisions (status, created_at DESC)`
- **AND** `alternatives` 必须使用 `TEXT[]`，不得使用 `JSONB`
- **AND** 不得删除原有 `title / created_at` 字段，不得破坏既有 `decision_links` 的外键引用

#### Scenario: 现有示例数据兼容回填

- **WHEN** migration 执行时 `decisions` 表已有 `phase02` seed 插入的 `title-only` 数据
- **THEN** migration 必须通过 `UPDATE decisions SET ... WHERE <新字段> IS NULL` 完成回填
- **AND** 回填必须保留原有 `title / created_at`
- **AND** `context / problem / choice / reason` 必须回填为占位文本 `（历史决策，phase03 升级前无结构化上下文）`
- **AND** `alternatives` 回填为 `'{}'`，`impact` 回填为 `''`，`status` 回填为 `'proposed'`
- **AND** 回填必须可重复执行（`WHERE` 条件保证幂等）

### Requirement: 基线 seed 必须与 phase03-09 覆盖维度一致

系统 SHALL 实现 `database/seeds/seed_decision_mainline_baseline.sql`，基线数据覆盖维度必须与 `phase03-09` 冻结结论完全一致。

#### Scenario: 基线 seed 文件落点与结构

- **WHEN** 实现 `Decision Center` 基线种子
- **THEN** seed 文件必须落在 `database/seeds/seed_decision_mainline_baseline.sql`
- **AND** 该文件必须同时承担"清空 + 恢复"职责，开头包含 `DELETE FROM decisions`，后续 `INSERT` 使用幂等守卫
- **AND** 必须通过 `BEGIN / COMMIT` 事务包裹
- **AND** 幂等语义必须与正式实现一致：复用 `phase02` 标题的 `Decision 1` 必须先 `UPDATE` 收口已有 `placeholder`（含 `seed_readonly_prereqs.sql` 插入的占位内容）为正式基线内容，再 `INSERT` 补缺；`Decision 2/3`（新标题）使用 `WHERE NOT EXISTS` 守卫；`decision_links` 使用 `ON CONFLICT (decision_id, module_id) DO NOTHING`

#### Scenario: 基线 Decision 数据覆盖维度

- **WHEN** 定义基线数据
- **THEN** 必须至少包含 `3` 条结构化 `Decision`，覆盖 `1 proposed + 1 active + 1 archived 或 superseded`
- **AND** 其中 `1` 条必须保留 `phase02` 原有 `title`（`关于 auth-service 技术选型的决策`），并补全结构化字段
- **AND** 至少 `1` 条必须包含 `alternatives` 数组（`2` 个以上条目）
- **AND** 至少 `1` 条的 `alternatives` 必须为空数组
- **AND** 至少 `1` 条的 `impact` 必须为空字符串
- **AND** 保留 `phase02` 原有 `title` 的 `Decision 1` 必须带 `source_module_id`（关联 `auth-service` 模块），演示从 `Module Detail` 带上下文创建的来源承接；其余 `Decision` 的 `source_module_id` 为 `NULL`，验证无来源上下文的空字符串返回语义

#### Scenario: 基线 decision_links 数据

- **WHEN** 定义基线 `decision_links`
- **THEN** 必须至少包含 `2` 条 `decision_links`，关联到 `reset_module_mainline` 已建立的 `modules` 基线
- **AND** 其中 `1` 条必须复用 `phase02` 原有的 `关于 auth-service 技术选型的决策` -> `auth-service` 关联
- **AND** 至少 `1` 条 `Decision` 必须没有任何 `decision_links`，验证 `link_count = 0` 与 `linked_module_summary = ''`
- **AND** `decision_links` 的 `INSERT` 必须通过 `module_name` 与 `decision title` 查找 `ID`，不得硬编码 `UUID`

#### Scenario: seed_readonly_prereqs.sql decisions seed 更新

- **WHEN** 更新 `seed_readonly_prereqs.sql`
- **THEN** `decisions` seed 必须从 `title-only` 更新为结构化字段插入
- **AND** 必须保持原有 `title`（`关于 auth-service 技术选型的决策`）
- **AND** 必须补全 `context / problem / choice / reason / status` 必填字段
- **AND** `alternatives` 必须设为 `'{}'`，`impact` 必须设为 `''`
- **AND** 该更新必须保证 `run_seeds.sh` 可重复执行

### Requirement: 重置脚本必须与 reset_module_mainline.sh 同构

系统 SHALL 实现 `database/scripts/reset_decision_mainline.sh`，脚本结构与 `reset_module_mainline.sh` 同构。

#### Scenario: 重置脚本落点与模式

- **WHEN** 实现 `Decision Center` 重置脚本
- **THEN** 脚本必须落在 `database/scripts/reset_decision_mainline.sh`
- **AND** 必须支持三种模式：`--clean-only`、`--restore-only`、默认（清空 + 恢复）
- **AND** 必须复用 `reset_module_mainline.sh` 的 `resolve_psql` 模式
- **AND** 必须复用相同的环境变量覆盖参数

#### Scenario: 清空范围冻结

- **WHEN** 执行清空
- **THEN** 清空范围必须是 `DELETE FROM decisions`，依赖 `decision_links.decision_id` 的 `ON DELETE CASCADE` 级联清空 `decision_links`
- **AND** 不得清空 `modules / products / repositories` 表
- **AND** 清空后必须输出当前 `decisions` 与 `decision_links` 计数

#### Scenario: 前置校验

- **WHEN** 执行重置脚本
- **THEN** 必须校验 `PSCO` 数据库已存在
- **AND** 在 `--restore-only` 与默认模式下，必须校验 `modules` 基线数据已存在（`SELECT COUNT(*) FROM modules >= 1`）
- **AND** 不得在 `modules` 为空时恢复 `decision` 基线

### Requirement: 应用装配必须将 Decision Center 挂载到 /api

系统 SHALL 在 `backend/internal/platform/router.go` 中新增 `mountDecisionCenter` 装配函数，并在 `server.go` 中将 `Decision Center` 路由挂载到 `/api`。

#### Scenario: 装配点与路由注册

- **WHEN** 装配 `Decision Center` 模块
- **THEN** 必须在 `platform/router.go` 中新增 `mountDecisionCenter(r chi.Router, pool *pgxpool.Pool)` 函数
- **AND** 装配顺序必须为：repository 层 → candidate 层 → service 层 → handler 层 → 路由注册
- **AND** 必须在 `server.go` 的 `r.Route("/api", ...)` 中调用 `mountDecisionCenter(r, pool)`
- **AND** 路由注册必须对齐 `phase03-10` §7.7 RPC → HTTP 映射矩阵：
  - `GET /api/decisions` — `ListDecisions`
  - `GET /api/decisions/{decisionId}` — `GetDecisionDetail`
  - `POST /api/decisions` — `CreateDecision`
  - `POST /api/decisions/{decisionId}/links` — `LinkDecisionToTarget`
  - `GET /api/decisions/{decisionId}/candidates/modules` — `ListDecisionModuleCandidates`

## MODIFIED Requirements

### Requirement: Decision Center 后端从"设计冻结"推进为"可运行实现"

系统 SHALL 将 `phase03-07 / 10` 中已冻结的 `Decision Center` 后端模块边界、接口分组与文件落点，从"设计层冻结"推进为"仓库内可运行的实际实现"。

#### Scenario: 后端阶段推进

- **WHEN** `phase03-12` 完成
- **THEN** `Decision Center` 后端不再只停留在 `phase03-07 / 10` 文档正文中
- **AND** 必须在仓库 `backend/internal/decisioncenter/` 内拥有实际可运行的读写接口
- **AND** 后续 `phase03-13 / 14` 必须优先消费本阶段交付的 API，而不是回到文档层手工解释接口

## REMOVED Requirements

### Requirement: Decision Center 后端只停留在设计层

**Reason**: `phase03-12` 的目标就是把 `Decision Center` 后端从"已经设计好"推进到"仓库中已实现、可运行、可被前端消费"的状态。

**Migration**: 后续 `Decision Center` 的前端实现与联调验收应统一从 `backend/internal/decisioncenter/` 模块进入；`phase03-07 / 10` 继续作为设计上游，不再承担仓库内后端主线入口职责。
