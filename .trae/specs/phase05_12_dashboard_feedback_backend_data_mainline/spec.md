# Phase05-12 Dashboard + Feedback 后端与数据主线实现 Spec

## Why

`phase05-07` 已冻结 Dashboard 后端模块边界、文件分层、接口分组、query service owner、跨模块读取方向、归一化 owner 与错误响应 owner；`phase05-08` 已冻结 `.proto` 合同设计；`phase05-11` 已将 `.proto` 合同落地为仓库主线；`phase05-09` 已冻结验收环境、九类 fixture 与局部错误模拟机制。但截至当前，仓库中还没有 Dashboard 后端模块的实际代码、跨模块 candidate reader 实现、platform 装配接线、reset 脚本与 fixture SQL 文件。

`phase05-12` 的目标是把上述已冻结结论推进为仓库内实际存在、可编译、可运行、可被前端直接消费的后端与数据主线，为 `phase05-13` 前端实现与 `phase05-14` 联调验收提供可执行的后端基线。

## What Changes

- 新增 `backend/internal/dashboard/` 模块，承接 Dashboard 三类只读聚合读取
- 新增 `backend/internal/dashboard/types.go` 承接 DTO 与领域类型，从 `dashboard.proto` 单向派生
- 新增 `backend/internal/dashboard/errors.go` 承接业务错误哨兵值
- 新增 `backend/internal/dashboard/service/query_service.go` 承接单一 `QueryService` 与三个读组方法
- 新增 `backend/internal/dashboard/candidate/` 承接跨模块 reader 接口定义与实现
- 新增 `backend/internal/dashboard/handler/query_handler.go` 承接三个 GET HTTP endpoint
- 新增 `backend/internal/dashboard/handler/response.go` 承接统一 JSON 响应
- 更新 `backend/internal/platform/server.go` 与 `router.go` 装配 Dashboard 模块
- 新增 `database/scripts/reset_dashboard_acceptance.sh` 作为 Dashboard 验收统一入口
- 新增 `database/seeds/seed_dashboard_acceptance_baseline.sql` 与九类 fixture SQL 文件
- 新增局部错误模拟机制（通过环境变量在后端实现层承接）
- **BREAKING**：Dashboard 后端实现完成后，`phase05-13` 前端与 `phase05-14` 联调验收必须从 `/api/dashboard/overview`、`/api/dashboard/feedback-signals`、`/api/dashboard/recent-activities` 三个 endpoint 消费数据，不得绕过后端直接 mock

## Impact

- Affected specs:
  - `phase05_07_dashboard_feedback_backend_module_boundary_interface_grouping`
  - `phase05_08_define_dashboard_feedback_proto_contract`
  - `phase05_09_design_dashboard_acceptance_baseline_fixtures`
  - `phase05_10_dashboard_feedback_formal_spec`
  - `phase05_11_dashboard_feedback_proto_mainline`
  - 后续 `phase05-13` 前端主线
  - 后续 `phase05-14` 联调与验收
- Affected code:
  - `backend/internal/dashboard/types.go`（新增）
  - `backend/internal/dashboard/errors.go`（新增）
  - `backend/internal/dashboard/service/query_service.go`（新增）
  - `backend/internal/dashboard/candidate/overview_readers.go`（新增）
  - `backend/internal/dashboard/candidate/feedback_readers.go`（新增）
  - `backend/internal/dashboard/candidate/activity_readers.go`（新增）
  - `backend/internal/dashboard/handler/query_handler.go`（新增）
  - `backend/internal/dashboard/handler/response.go`（新增）
  - `backend/internal/platform/server.go`（更新）
  - `backend/internal/platform/router.go`（更新）
  - `database/scripts/reset_dashboard_acceptance.sh`（新增）
  - `database/seeds/seed_dashboard_acceptance_baseline.sql`（新增）
  - `database/seeds/seed_dashboard_fixture_*.sql`（新增 9 个文件）

## ADDED Requirements

### Requirement: Dashboard 后端模块必须按 phase05-07 冻结的物理边界与文件分层落地

系统 SHALL 将 Dashboard 后端逻辑落地为独立模块 `backend/internal/dashboard/`，沿用 `phase02-09 / phase03-10 / phase04-10` 已验证的 `handler / service / candidate / types.go / errors.go / handler/response.go` 分层结构。

#### Scenario: 模块物理落点

- **WHEN** 执行 `phase05-12`
- **THEN** 仓库中必须存在 `backend/internal/dashboard/` 目录
- **AND** 必须包含 `types.go / errors.go / service/query_service.go / handler/query_handler.go / handler/response.go / candidate/` 文件分层
- **AND** Go package 必须为 `github.com/psco/backend/internal/dashboard`
- **AND** 不得把 Dashboard 聚合读取逻辑塞入既有任一 canonical 模块

#### Scenario: 支撑文件单值化

- **WHEN** 新增 Dashboard 模块支撑文件
- **THEN** `errors.go / types.go / handler/response.go` 必须按职责单值化映射到唯一文件
- **AND** 禁止把 errors / types / response 逻辑散落在 handler 或 service 内部

### Requirement: Dashboard DTO 必须从 dashboard.proto 单向派生

系统 SHALL 将 Dashboard DTO 落点在 `backend/internal/dashboard/types.go`，字段语义从 `proto/psco/dashboard/v1/dashboard.proto` 单向派生或显式对齐。

#### Scenario: DTO 字段对齐

- **WHEN** 实现 Dashboard DTO
- **THEN** `DashboardOverview` DTO 至少必须承接 `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count`
- **AND** `FeedbackSignal` DTO 至少必须承接 `signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label`
- **AND** `ProductAssetCoverageSummary` DTO 至少必须承接 `fully_bound_product_count / missing_both_bindings_count / missing_repository_binding_count / missing_module_binding_count / representative_signals`
- **AND** `RecentActivityItem` DTO 至少必须承接 `activity_type / activity_at / target_type / target_id / target_label`
- **AND** DTO 使用 snake_case JSON tag，与 `.proto` 的 snake_case 对齐
- **AND** 枚举使用 Go `string` 类型与 snake_case 常量值（对齐现有 `ProductStatus` 模式）
- **AND** `activity_at` 使用 `time.Time` 类型，对应 `.proto` 的 `google.protobuf.Timestamp`

#### Scenario: DTO 不得新增 proto 中不存在的字段

- **WHEN** 编写 Dashboard DTO
- **THEN** 不得在 DTO 中新增 `score / trend / external_metric / recommendation_reason` 等 `.proto` 中不存在的业务字段
- **AND** 不得在 DTO 中新增错误标记字段（错误语义由 HTTP 状态码承载）

### Requirement: 单一 QueryService 必须承接三个读组编排

系统 SHALL 将 Dashboard 三类聚合读取的 query service owner 冻结为 `backend/internal/dashboard/service/query_service.go` 中的单一 `QueryService`，包含 `ReadOverview / ReadFeedbackSignal / ReadRecentActivity` 三个方法。

#### Scenario: QueryService 结构

- **WHEN** 实现 Dashboard QueryService
- **THEN** 必须为单一 `QueryService` 结构体
- **AND** 必须包含 `ReadOverview(ctx) (*DashboardOverview, error)` 方法
- **AND** 必须包含 `ReadFeedbackSignal(ctx) (*FeedbackSignalReadResult, error)` 方法
- **AND** 必须包含 `ReadRecentActivity(ctx) ([]RecentActivityItem, error)` 方法
- **AND** 不得为三个读组分别发明独立 service owner
- **AND** 不得把读组逻辑下沉到 handler 层

#### Scenario: ReadOverview 整页失败语义

- **WHEN** `ReadOverview` 任一组成 reader 失败
- **THEN** 必须返回整页失败语义（error）
- **AND** 不得吞掉局部失败后返回部分计数
- **AND** 该语义与 `phase05-04` 已冻结的"主聚合失败触发整页失败"一致

#### Scenario: ReadFeedbackSignal 局部失败语义

- **WHEN** `ReadFeedbackSignal` 的任一 reader 失败
- **THEN** 必须返回局部失败语义（error）
- **AND** 不得拖垮 `ReadOverview` 的整页成功语义
- **AND** 该语义与 `phase05-04` 已冻结的"附属聚合失败只触发局部失败"一致

#### Scenario: ReadRecentActivity 局部失败语义

- **WHEN** `ReadRecentActivity` 的任一 reader 失败
- **THEN** 必须返回局部失败语义（error）
- **AND** 不得拖垮 `ReadOverview` 的整页成功语义

### Requirement: 跨模块读取必须通过 candidate 子包隔离

系统 SHALL 将 Dashboard 跨模块读取通过 `backend/internal/dashboard/candidate/` 子包定义并拥有的 reader 接口与实现承接，`service/` 层不得直接跨模块写 SQL。

#### Scenario: candidate reader 分工

- **WHEN** 实现 Dashboard candidate reader
- **THEN** 必须至少承接以下 reader 分工：
  - Overview 相关：`ModuleCountReader / ProductCountReader / RepositoryCountReader / DecisionCountReader`
  - Feedback 相关：`PendingDecisionSignalReader / ProductAssetCoverageReader`
  - Activity 相关：`ModuleActivityReader / ProductActivityReader / RepositoryActivityReader / DecisionActivityReader`
- **AND** reader 接口定义与实现均由 Dashboard `candidate/` 子包自己拥有
- **AND** reader 实现可以直接读取 canonical 模块的表，但必须在 `candidate/` 子包内隔离
- **AND** Dashboard `service/` 层不得直接跨模块写 SQL

#### Scenario: candidate reader 文件落点

- **WHEN** 新增 Dashboard candidate 文件
- **THEN** overview 相关 reader 落点在 `candidate/overview_readers.go`
- **AND** feedback 相关 reader 落点在 `candidate/feedback_readers.go`
- **AND** activity 相关 reader 落点在 `candidate/activity_readers.go`
- **AND** 不得把所有 reader 塞进单文件或散落到 service 层

### Requirement: DashboardOverviewRead 聚合逻辑必须对齐 phase05-04

系统 SHALL 将 `ReadOverview` 的聚合逻辑冻结为编排四个 canonical 模块的计数 reader，返回 `DashboardOverview`。

#### Scenario: ReadOverview 计数来源

- **WHEN** `ReadOverview` 执行聚合
- **THEN** `module_count` 来自 `ModuleCountReader`
- **AND** `product_count / product_with_module_count / product_with_repository_count` 来自 `ProductCountReader`
- **AND** `repository_count` 来自 `RepositoryCountReader`
- **AND** `decision_count` 来自 `DecisionCountReader`
- **AND** 任一 reader 失败时必须返回整页失败

#### Scenario: product_with_module_count 与 product_with_repository_count 语义

- **WHEN** 计算 `product_with_module_count`
- **THEN** 必须统计在 `product_modules` 表中存在记录的去重 product 数量
- **AND** 计算 `product_with_repository_count` 必须统计在 `product_repositories` 表中存在记录的去重 product 数量

### Requirement: FeedbackSignalRead 必须归一化 pending_decision_signals 与 product_asset_coverage

系统 SHALL 将 `ReadFeedbackSignal` 冻结为编排 `PendingDecisionSignalReader` 与 `ProductAssetCoverageReader`，在 Dashboard 模块内归一化为统一 `FeedbackSignal` 列表与 `ProductAssetCoverageSummary`。

#### Scenario: 反馈信号归一化与排序

- **WHEN** `ReadFeedbackSignal` 执行归一化
- **THEN** `pending_decision_signals` 与 `product_asset_coverage` 必须归一化为统一 `FeedbackSignal` 列表
- **AND** 归一化后的主队列（`current_focus_signals`）最多展示 `5` 条
- **AND** 主队列必须按优先级排序：`P1 > P2 > P3 > P4`
- **AND** 同优先级内按 `created_at DESC` 回退排序
- **AND** `ProductAssetCoverageSummary` 的 `representative_signals` 最多展示 `3` 条代表性缺口项

#### Scenario: pending_decision_signals 信号生成

- **WHEN** 读取到 pending 状态的 decisions
- **THEN** 若存在已绑定具体 `decision_id` 的 decision_link，必须生成单项 `FeedbackSignal`（`target_type = DECISION_DETAIL`，`target_id = decision_id`）
- **AND** 若存在未绑定单一 `decision_id` 的聚合决策信号，必须生成聚合 `FeedbackSignal`（`target_type = DECISION_LIST`，`target_id = ""`）
- **AND** pending decision 的状态判定必须沿用 `phase03` Decision Center 已冻结的 `proposed` status 语义

#### Scenario: product_asset_coverage 信号生成

- **WHEN** 读取到资产缺口数据
- **THEN** 同时缺少 `product_modules` 与 `product_repositories` 记录的 product 必须生成 `PRODUCT_MISSING_BOTH_BINDINGS` 信号
- **AND** 仅有 `product_modules` 但无 `product_repositories` 的 product 必须生成 `PRODUCT_MISSING_REPOSITORY_BINDING` 信号
- **AND** 仅有 `product_repositories` 但无 `product_modules` 的 product 必须生成 `PRODUCT_MISSING_MODULE_BINDING` 信号
- **AND** `missing_both_bindings_count` 必须作为独立计数字段，不回退为两个单缺口的隐式组合

#### Scenario: 空态响应

- **WHEN** 不存在 `pending_decision_signals` 且 `product_asset_coverage` 不产生缺口项
- **THEN** `ReadFeedbackSignal` 必须返回成功语义
- **AND** `current_focus_signals` 必须为空列表
- **AND** `ProductAssetCoverageSummary` 必须返回完整结构，三类缺口计数为 `0`，`representative_signals` 为空列表
- **AND** 不得将空态映射为错误

### Requirement: RecentActivityRead 必须归并多模块活动并按 activity_at 排序

系统 SHALL 将 `ReadRecentActivity` 冻结为编排四个 canonical 模块的活动 reader，在 Dashboard 模块内映射为统一 `RecentActivityItem` 列表，按 `activity_at` 倒序排序，最多返回 `10` 条。

#### Scenario: 活动类型映射

- **WHEN** `ReadRecentActivity` 执行类型映射
- **THEN** 必须覆盖以下活动类型：
  - `module`：来自 `modules` 表的 `created_at`
  - `release`：来自 `module_releases` 表的 `released_at`，target 回落到所属 Module Detail
  - `product`：来自 `products` 表的 `created_at`
  - `repository`：来自 `repositories` 表的 `created_at`
  - `decision`：来自 `decisions` 表的 `created_at`
  - `product_module_binding`：来自 `product_modules` 表的 `created_at`
  - `product_repository_binding`：来自 `product_repositories` 表的 `created_at`
  - `module_repository_binding`：来自 `module_repositories` 表的 `created_at`
- **AND** 每条活动项必须带显式 `activity_at` 时间字段
- **AND** 不得依赖隐式 `created_at` 推断（release 使用 `released_at`）

#### Scenario: 活动项跳转目标映射

- **WHEN** 映射活动项的 `target_type / target_id / target_label`
- **THEN** `module` 活动 → `MODULE_DETAIL`，`target_id = module_id`，`target_label = module_name`
- **AND** `release` 活动 → `MODULE_DETAIL`，`target_id = module_id`，`target_label = module_name`
- **AND** `product` 活动 → `PRODUCT_DETAIL`，`target_id = product_id`，`target_label = product_name`
- **AND** `repository` 活动 → `REPOSITORY_DETAIL`，`target_id = repository_id`，`target_label = repository_name`
- **AND** `decision` 活动 → `DECISION_DETAIL`，`target_id = decision_id`，`target_label = decision_title`
- **AND** `product_module_binding` 活动 → `PRODUCT_DETAIL`，`target_id = product_id`，`target_label = product_name`
- **AND** `product_repository_binding` 活动 → `REPOSITORY_DETAIL`，`target_id = repository_id`，`target_label = repository_name`
- **AND** `module_repository_binding` 活动 → `REPOSITORY_DETAIL`，`target_id = repository_id`，`target_label = repository_name`

#### Scenario: 空态响应

- **WHEN** 系统暂无可展示的最近活动
- **THEN** `ReadRecentActivity` 必须返回空列表
- **AND** 不得映射为资源不存在或接口错误

### Requirement: Dashboard handler 必须承接三个 GET endpoint

系统 SHALL 将 Dashboard handler 冻结为三个独立 GET HTTP endpoint，对齐 `phase05-08` 已冻结的 RPC → HTTP 映射矩阵。

#### Scenario: HTTP 路由注册

- **WHEN** 实现 Dashboard handler 路由注册
- **THEN** 必须注册以下路由：
  - `GET /api/dashboard/overview` → 承接 `DashboardOverviewRead`
  - `GET /api/dashboard/feedback-signals` → 承接 `FeedbackSignalRead`
  - `GET /api/dashboard/recent-activities` → 承接 `RecentActivityRead`
- **AND** 三个 endpoint 当前阶段无 body、无 query 过滤、无路径参数
- **AND** handler 必须显式组装空 Proto request 语义（直接调用 service 方法），不得绕过 service 合同边界

#### Scenario: 整页失败响应

- **WHEN** `ReadOverview` 返回失败
- **THEN** handler 必须返回 `500` 错误响应
- **AND** 不得返回部分计数或伪装成功

#### Scenario: 局部失败响应

- **WHEN** `ReadFeedbackSignal` 或 `ReadRecentActivity` 返回失败
- **THEN** 各自 handler 必须返回 `500` 错误响应
- **AND** 不得把附属读失败升级为整页失败
- **AND** 当前阶段不在附属读 endpoint 响应包络里额外发明"局部失败标记"
- **AND** Dashboard 页面级"主聚合成功、附属聚合局部失败"语义由前端基于三次独立请求结果组合派生

#### Scenario: 空态成功响应

- **WHEN** 读取返回零计数或空列表
- **THEN** handler 必须返回 `200` 成功响应
- **AND** 不得把空态映射为 `404` 或 `500`

#### Scenario: 响应包络必须对齐 proto envelope

- **WHEN** 三个 GET endpoint 返回成功响应
- **THEN** `GET /api/dashboard/overview` 响应体必须为 `{"overview": {...}}`，对齐 `GetDashboardOverviewResponse { overview = 1 }`
- **AND** `GET /api/dashboard/feedback-signals` 响应体必须为 `{"current_focus_signals": [...], "asset_feedback_summary": {...}}`，对齐 `GetFeedbackSignalsResponse`
- **AND** `GET /api/dashboard/recent-activities` 响应体必须为 `{"activities": [...]}`，对齐 `GetRecentActivitiesResponse { activities = 1 }`
- **AND** handler 不得直接返回裸 `DashboardOverview` 或裸 `[]RecentActivityItem`
- **AND** 响应 DTO 落点在 `backend/internal/dashboard/types.go`，由 `DashboardOverviewReadResult / FeedbackSignalReadResult / RecentActivityReadResult` 三个结构体承接
- **AND** 该约束保证 HTTP JSON 形状与 `proto/psco/dashboard/v1/dashboard.proto` 冻结的唯一合同源一致，不形成第二套合同源

### Requirement: platform 装配层必须接线 Dashboard 模块

系统 SHALL 在 `backend/internal/platform/` 装配层构造 Dashboard candidate reader、QueryService 与 handler，并将路由挂到 `/api` 下。

#### Scenario: 装配顺序

- **WHEN** 在 `platform/server.go` 中装配 Dashboard
- **THEN** 必须在既有四个 canonical 模块装配之后装配 Dashboard
- **AND** 必须通过 `buildDashboard(pool)` 构造 candidate reader 与 QueryService
- **AND** 必须通过 `mountDashboard(r, querySvc)` 注册路由
- **AND** Dashboard QueryService 的跨模块读依赖必须在 platform 装配点注入，不得在 Dashboard 模块内部自行 new

#### Scenario: 路由挂载

- **WHEN** `mountDashboard` 注册路由
- **THEN** 必须在 `/api` 子路由器上注册三个 GET 路径
- **AND** 不得与其他模块路由冲突

### Requirement: 局部错误模拟必须通过环境变量受控承接

系统 SHALL 将 Dashboard 局部错误模拟通过环境变量在后端实现层承接，不要求 fixture SQL 层模拟。

#### Scenario: 环境变量模拟机制

- **WHEN** 验收需要模拟 `DashboardOverviewRead` 失败
- **THEN** 必须通过环境变量 `DASHBOARD_SIMULATE_OVERVIEW_ERROR=true` 触发 overview reader 返回错误
- **AND** 该机制必须可重复启用与关闭
- **AND** 不得通过删除表、重命名列或破坏 schema 来模拟错误

- **WHEN** 验收需要模拟 `FeedbackSignalRead` 失败
- **THEN** 必须通过环境变量 `DASHBOARD_SIMULATE_FEEDBACK_ERROR=true` 触发 feedback reader 返回错误

- **WHEN** 验收需要模拟 `RecentActivityRead` 失败
- **THEN** 必须通过环境变量 `DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR=true` 触发 recent activity reader 返回错误

#### Scenario: 环境变量读取方式

- **WHEN** Dashboard 后端启动
- **THEN** 必须在初始化时读取 `DASHBOARD_SIMULATE_OVERVIEW_ERROR / DASHBOARD_SIMULATE_FEEDBACK_ERROR / DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR` 环境变量
- **AND** 必须在 candidate reader 实现层检查这些环境变量
- **AND** 环境变量为 `true` 时必须返回模拟错误
- **AND** 环境变量未设置或为其他值时必须正常执行

### Requirement: reset_dashboard_acceptance.sh 必须作为 Dashboard 验收统一入口

系统 SHALL 冻结 `database/scripts/reset_dashboard_acceptance.sh` 为 Dashboard 验收的唯一统一入口，提供 `--clean-only`、`--restore-only`、`--fixture <name>` 与默认 `clean+restore` 四种模式。

#### Scenario: 脚本模式矩阵

- **WHEN** 执行 Dashboard 验收环境重置
- **THEN** `reset_dashboard_acceptance.sh` 必须支持以下模式：
  - 默认（无参数）：先清空所有 Dashboard 相关数据，再恢复完整基线
  - `--clean-only`：仅清空所有 Dashboard 相关数据
  - `--restore-only`：仅恢复完整基线
  - `--fixture <name>`：先清空，再加载指定 fixture
- **AND** `--fixture` 的 `<name>` 只允许取九类 fixture 名称之一

#### Scenario: 清空范围

- **WHEN** 执行清空操作
- **THEN** 必须通过编排既有 `reset_product_repository_mainline.sh --clean-only`、`reset_decision_mainline.sh --clean-only` 与 `reset_module_mainline.sh --clean-only` 实现
- **AND** 清空顺序必须按依赖逆序：先清 `product_repository`，再清 `decision`，最后清 `module`
- **AND** 不得绕过既有脚本直接 `DELETE FROM` 或 `TRUNCATE` 底层表

#### Scenario: 恢复范围

- **WHEN** 执行恢复操作
- **THEN** 必须按依赖顺序执行既有 reset 脚本的 `--restore-only` 模式
- **AND** 恢复后必须执行 `seed_dashboard_acceptance_baseline.sql` 补齐 Dashboard 验收所需的额外基线数据
- **AND** 恢复操作必须保证幂等

#### Scenario: fixture 加载

- **WHEN** 执行 `--fixture <name>` 模式
- **THEN** 必须先执行清空操作
- **AND** 再加载 `seed_dashboard_fixture_<name>.sql`
- **AND** 所有 fixture 加载必须遵守"先清空再加载指定 fixture"的统一语义，禁止叠加已有数据

#### Scenario: host psql 模式密码处理必须对齐既有 reset 脚本

- **WHEN** `reset_dashboard_acceptance.sh` 在宿主机检测到 `psql`（`PSQL_MODE=host`）
- **THEN** 必须在 `resolve_psql` 之后立即补齐密码处理段，与 `reset_module_mainline.sh / reset_decision_mainline.sh / reset_product_repository_mainline.sh` 保持一致
- **AND** 当 `PG_PASSWORD` 未设置时，必须尝试从 `$PG_CONTAINER` 容器读取 `POSTGRES_PASSWORD`
- **AND** 容器未运行且 `PG_PASSWORD` 未设置时必须显式报错退出
- **AND** 必须通过 `export PGPASSWORD` 把密码暴露给本脚本直接执行的 `seed_readonly_prereqs.sql`、`seed_dashboard_acceptance_baseline.sql` 与 fixture SQL
- **AND** 不得依赖既有 reset 脚本间接处理本脚本直接执行 SQL 的密码，因为既有脚本只在自身进程内 `export PGPASSWORD`，不会反传给父 shell

### Requirement: 九类 fixture SQL 文件必须可重复执行且幂等

系统 SHALL 将九类 Dashboard fixture SQL 文件落地到 `database/seeds/`，每个文件必须可重复执行不报错。

#### Scenario: fixture 文件命名与落点

- **WHEN** 新增 Dashboard fixture 文件
- **THEN** 基线文件必须命名为 `seed_dashboard_acceptance_baseline.sql`
- **AND** 九类 fixture 文件必须命名为 `seed_dashboard_fixture_<fixture-name>.sql`
- **AND** 所有文件必须落点在 `database/seeds/` 目录
- **AND** 不得在 `database/seeds/` 下新建子目录

#### Scenario: fixture SQL 幂等约束

- **WHEN** 编写 fixture SQL 文件
- **THEN** 每个 fixture 文件必须可重复执行不报错
- **AND** 必须使用 `ON CONFLICT DO NOTHING`、`WHERE NOT EXISTS` 或 `UPDATE-then-INSERT` 模式保证幂等
- **AND** 必须在文件头部以注释说明：用途、依赖、幂等语义、使用方式

## MODIFIED Requirements

### Requirement: phase05-07 后端模块边界从"设计冻结"推进为"仓库实现落地"

系统 SHALL 将 `phase05-07` 已冻结的 Dashboard 后端模块边界、文件分层、接口分组、query service owner、跨模块读取方向、归一化 owner 与错误响应 owner，从"规格层定义"推进为"仓库内实际存在且可编译运行的代码"。

#### Scenario: 模块边界落地

- **WHEN** `phase05-12` 完成
- **THEN** Dashboard 后端模块不再只停留在 `phase05-07` 文档正文中
- **AND** 必须在仓库 `backend/internal/dashboard/` 内拥有实际文件、可编译代码与可运行 endpoint
- **AND** 后续 `phase05-13 / 14` 必须优先引用该已落地后端模块，而不是回到文档层手工解释

## REMOVED Requirements

### Requirement: Dashboard 后端模块继续停留在规格层

**Reason**: `phase05-12` 的目标就是把 Dashboard 后端从"已经设计好"推进到"仓库中已经落地、可编译、可运行、可被前端引用"的状态。

**Migration**: 后续 Dashboard 前端实现与联调验收应统一从 `backend/internal/dashboard/` 模块与 `/api/dashboard/*` endpoint 进入；`phase05-07` 继续作为设计上游，不再承担仓库内后端主线入口职责。
