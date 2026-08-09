# Phase05-08 Dashboard + Feedback 最小 Protocol Buffers 合同设计 Spec

## Why

`phase05-02` 已冻结 `Feedback Signal` 与 `Recent Activity` 的最小展示模型，`phase05-03` 已冻结它们的 canonical 落点与导航身份解释，`phase05-04` 已冻结 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 的最小接口边界与错误语义前提，`phase05-07` 又已冻结 Dashboard 后端模块边界、读组 owner 与 DTO/合同映射 owner。但截至当前，仓库主线里还没有 `Dashboard + Feedback` 对应的 `.proto` 合同设计结果。

如果不把 `Dashboard + Feedback` 的文件落点、包名版本、RPC 矩阵、消息边界、字段编号、`missing both bindings` 的合同表达方式、`RecentActivityItem.activity_at` 的时间语义、`chi + JSON HTTP` 过渡映射与 `buf breaking` 前提写成单值结论，后续 `phase05-09` 及实现阶段就会继续在“Dashboard 是否单独成包”“反馈信号用 string 还是 enum”“聚合信号如何表达 list 落点”“哪些变化算 breaking”之间漂移。

## What Changes

- 冻结 `Dashboard + Feedback` 最小 `.proto` 合同源进入现有单一 `proto/` workspace
- 冻结 `proto/psco/dashboard/v1/dashboard.proto` 的文件落点、包名、版本语义与生成产物归属
- 冻结 `DashboardService` 的最小 RPC 矩阵，单值承接 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead`
- 冻结 `DashboardOverview`、`FeedbackSignal`、`ProductAssetCoverageSummary`、`RecentActivityItem` 的最小消息结构与字段编号方案
- 冻结 `FeedbackSignalFamily / FeedbackSignalCode / FeedbackSignalPriority / DashboardTargetType / RecentActivityType` 的最小枚举语义
- 冻结 `product missing both bindings` 在合同中的独立表达方式，不允许回退为隐式组合语义
- 冻结 `RecentActivityItem.activity_at` 为当前阶段唯一显式活动时间字段与排序前提锚点
- 冻结 `.proto` 与 `chi + JSON HTTP` 的显式映射矩阵，阻断第二套合同源
- 冻结 `reserved`、递增编号与 `buf build / lint / generate / breaking` 的最小演进规则
- **BREAKING**：后续 Dashboard 相关 HTTP DTO、前端消费类型与 handler adapter 不得再把手写 JSON 结构视为与 `.proto` 并列的第二套合同源

## Impact

- Affected specs:
  - `phase05_dashboard_feedback_foundation`
  - `phase05_02_feedback_signal_priority_display_model`
  - `phase05_03_dashboard_navigation_context_return_path`
  - `phase05_04_dashboard_aggregate_api_error_boundary`
  - `phase05_07_dashboard_feedback_backend_module_boundary_interface_grouping`
  - `phase04_11_product_repository_binding_proto_mainline`
- Affected code:
  - `proto/psco/dashboard/v1/dashboard.proto`
  - `proto/buf.yaml`
  - `proto/buf.gen.yaml`
  - `proto/Makefile`
  - `proto/README.md`
  - `backend/internal/gen/proto/psco/dashboard/v1/`
  - `frontend/src/gen/proto/psco/dashboard/v1/`
  - 后续 `backend/internal/dashboard/types.go`
  - 后续 `frontend/src/features/dashboard/`

## ADDED Requirements

### Requirement: Dashboard + Feedback 合同必须进入现有 proto workspace

系统 SHALL 将 `Dashboard + Feedback` 的最小 `.proto` 合同落地到现有 `proto/` workspace，而不是为 `phase05` 新建第二个 proto 根目录、第二套 `buf.yaml` 或第二套生成入口。

#### Scenario: 合同源文件落点

- **WHEN** 执行 `phase05-08`
- **THEN** 仓库中必须新增 `proto/psco/dashboard/v1/dashboard.proto`
- **AND** 该文件必须成为当前阶段 `Dashboard + Feedback` 的唯一合同定义入口
- **AND** 必须继续复用现有 `proto/buf.yaml`、`proto/buf.gen.yaml` 与 `proto/Makefile`
- **AND** 不得在 `backend/`、`frontend/` 或仓库根新增并列 proto 工作区

#### Scenario: 包名与版本语义

- **WHEN** `Dashboard + Feedback` 合同进入仓库主线
- **THEN** 包名与版本语义必须冻结为 `psco.dashboard.v1`
- **AND** 文件目录必须冻结为 `proto/psco/dashboard/v1/`
- **AND** Go 生成包路径必须与现有模式保持同构，落到 `backend/internal/gen/proto/psco/dashboard/v1/`
- **AND** TypeScript 生成产物必须落到 `frontend/src/gen/proto/psco/dashboard/v1/`
- **AND** 后续字段演进必须继续在 `v1` 版本语义下进行，不得在当前阶段另起第二套包名或版本口径

### Requirement: DashboardService 最小 RPC 矩阵必须单值化

系统 SHALL 将 `Dashboard + Feedback` 的服务接口冻结为单一 `DashboardService`，由它承接当前阶段三类只读聚合读取，不拆散为第二套 service 或并列 RPC 主线。

#### Scenario: 服务接口矩阵

- **WHEN** 后续 `.proto` 合同或实现讨论 Dashboard 的最小 RPC 矩阵
- **THEN** 必须冻结为：
  - `GetDashboardOverview` → 承接 `DashboardOverviewRead`
  - `GetFeedbackSignals` → 承接 `FeedbackSignalRead`
  - `GetRecentActivities` → 承接 `RecentActivityRead`
- **AND** 三个 RPC 都只承接读，不承接任何业务写入
- **AND** 不得并列增加趋势分析、通知中心、导出或批量操作 RPC
- **AND** 不得把三个读取拆成多个 service 定义

#### Scenario: request 边界

- **WHEN** 当前阶段定义三个读组的 request 消息
- **THEN** `GetDashboardOverviewRequest`、`GetFeedbackSignalsRequest` 与 `GetRecentActivitiesRequest` 必须保持无筛选、无分页、无排序切换的最小请求边界
- **AND** 不得为它们引入 `queryText / statusFilter / page / pageSize / sort / dateRange`
- **AND** 当前阶段允许使用空 request 消息承接默认读取，而不是额外发明 URL 或 body 级过滤合同

### Requirement: DashboardOverview 消息边界必须对齐 phase05-04

系统 SHALL 将 `DashboardOverview` 的消息边界冻结为只服务概览卡片与系统状态判定前提的主聚合读模型，不混入反馈队列或活动流字段。

#### Scenario: DashboardOverview 最小字段

- **WHEN** 后续 `.proto` 合同设计 `DashboardOverview`
- **THEN** 它至少必须承接：
  - `module_count`
  - `product_count`
  - `repository_count`
  - `decision_count`
  - `product_with_repository_count`
  - `product_with_module_count`
- **AND** 这些字段必须足以支撑 `phase05-04` 已冻结的空系统、非空缺口与创建导向 CTA 1-4 判定前提
- **AND** 不得在 `DashboardOverview` 中混入 `FeedbackSignal` 或 `RecentActivityItem`

#### Scenario: DashboardOverview 返回边界

- **WHEN** 定义 `GetDashboardOverviewResponse`
- **THEN** 响应必须只以 `DashboardOverview overview = 1` 这类单一主读模型承接
- **AND** 不得把整页错误、局部失败标记或反馈卡片列表直接并入该响应体
- **AND** 当前阶段错误语义继续由传输层显式映射承接，不在本消息中发明第二套错误包络

### Requirement: FeedbackSignal 合同必须冻结统一卡片模型与资产缺口摘要

系统 SHALL 将 `FeedbackSignalRead` 的合同边界冻结为“统一反馈主队列 + 资产缺口补充摘要”两层结构，并继续以统一 `FeedbackSignal` 作为主队列单值卡片模型。

#### Scenario: FeedbackSignal 最小字段模板

- **WHEN** 后续 `.proto` 合同设计统一反馈卡片
- **THEN** `FeedbackSignal` 至少必须承接：
  - `signal_family`
  - `signal_code`
  - `priority`
  - `title`
  - `summary`
  - `action_label`
  - `target_type`
  - `target_id`
  - `target_label`
- **AND** 上述字段语义必须继续对齐 `phase05-02` 已冻结的最小模板
- **AND** 不得减少字段，导致卡片无法同时承接“解释缺口”与“导航到 canonical owner”
- **AND** 当前阶段不得额外引入 `score / trend / external_metric / recommendation_reason`

#### Scenario: FeedbackSignalFamily 与 FeedbackSignalCode 枚举

- **WHEN** 当前阶段定义 `FeedbackSignalFamily` 与 `FeedbackSignalCode`
- **THEN** 至少必须冻结以下枚举值：
  - `FEEDBACK_SIGNAL_FAMILY_PENDING_DECISION`
  - `FEEDBACK_SIGNAL_FAMILY_PRODUCT_ASSET_COVERAGE`
  - `FEEDBACK_SIGNAL_CODE_PENDING_DECISION`
  - `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS`
  - `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING`
  - `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING`
- **AND** 每个枚举都必须包含首个 `UNSPECIFIED = 0`
- **AND** 不得用裸 string 取代这些冻结后的业务枚举

#### Scenario: FeedbackSignalPriority 枚举必须直接表达冻结后的优先级顺序

- **WHEN** 当前阶段定义 `FeedbackSignalPriority`
- **THEN** 它必须足以单值表达 `phase05-02 / phase05-04` 已冻结的优先级顺序：
  - `P1_PENDING_DECISION`
  - `P2_PRODUCT_MISSING_BOTH_BINDINGS`
  - `P3_PRODUCT_MISSING_REPOSITORY_BINDING`
  - `P4_PRODUCT_MISSING_MODULE_BINDING`
- **AND** 首个枚举值必须为 `UNSPECIFIED = 0`
- **AND** 不得重新压缩成会丢失顺序语义的模糊高/中/低三级口径

#### Scenario: target_type / target_id 的导航语义

- **WHEN** 当前阶段定义 `DashboardTargetType` 与消费 `FeedbackSignal.target_type / target_id`
- **THEN** 它们必须足以确定 `phase05-03` 已冻结的唯一 canonical 跳转目标
- **AND** `DashboardTargetType` 至少必须承接：
  - `DASHBOARD_TARGET_TYPE_DECISION_DETAIL`
  - `DASHBOARD_TARGET_TYPE_DECISION_LIST`
  - `DASHBOARD_TARGET_TYPE_PRODUCT_DETAIL`
  - `DASHBOARD_TARGET_TYPE_MODULE_DETAIL`
  - `DASHBOARD_TARGET_TYPE_REPOSITORY_DETAIL`
- **AND** 不得把 `target_type / target_id` 解释为“在 Dashboard 内直接执行动作”的内部指令
- **AND** 当 `target_type = DASHBOARD_TARGET_TYPE_DECISION_LIST` 时，`target_id` 必须允许为空字符串，以表达聚合决策信号落到 `Decision Center / List` 的语义

#### Scenario: ProductAssetCoverageSummary 的最小结构

- **WHEN** 当前阶段定义 `GetFeedbackSignalsResponse`
- **THEN** 它除了 `repeated FeedbackSignal current_focus_signals` 之外，还必须承接 `ProductAssetCoverageSummary asset_feedback_summary`
- **AND** `ProductAssetCoverageSummary` 至少必须包含：
  - `fully_bound_product_count`
  - `missing_both_bindings_count`
  - `missing_repository_binding_count`
  - `missing_module_binding_count`
  - `repeated FeedbackSignal representative_signals`
- **AND** `representative_signals` 最多展示 `3` 条代表性缺口项
- **AND** `current_focus_signals` 最多展示 `5` 条主队列信号

#### Scenario: missing both bindings 的合同表达方式

- **WHEN** `.proto` 合同表达 `product_asset_coverage`
- **THEN** `product missing both bindings` 必须作为独立 `FeedbackSignalCode` 与独立计数字段 `missing_both_bindings_count` 存在
- **AND** 不得回退为“缺仓库 + 缺模块”两个单缺口的隐式组合语义
- **AND** 同一个双缺口产品不得在代表项语义上同时重复表达为两个单缺口信号

### Requirement: RecentActivity 合同必须冻结活动类型与显式时间字段

系统 SHALL 将 `RecentActivityRead` 的合同边界冻结为独立活动流，并以 `RecentActivityType` 与显式 `activity_at` 字段承接 `phase05-02 / 03 / 04` 已冻结的活动项语义。

#### Scenario: RecentActivityItem 最小字段

- **WHEN** 后续 `.proto` 合同设计 `RecentActivityItem`
- **THEN** 它至少必须承接：
  - `activity_type`
  - `activity_at`
  - `target_type`
  - `target_id`
  - `target_label`
- **AND** `activity_at` 必须使用显式时间字段类型承接，而不是依赖隐式 `created_at` 约定
- **AND** 当前阶段最多返回 `10` 条活动项
- **AND** 不得把活动流字段并入 `FeedbackSignal`

#### Scenario: RecentActivityType 枚举

- **WHEN** 当前阶段定义 `RecentActivityType`
- **THEN** 至少必须冻结以下枚举值：
  - `RECENT_ACTIVITY_TYPE_MODULE`
  - `RECENT_ACTIVITY_TYPE_RELEASE`
  - `RECENT_ACTIVITY_TYPE_PRODUCT`
  - `RECENT_ACTIVITY_TYPE_REPOSITORY`
  - `RECENT_ACTIVITY_TYPE_DECISION`
  - `RECENT_ACTIVITY_TYPE_PRODUCT_MODULE_BINDING`
  - `RECENT_ACTIVITY_TYPE_PRODUCT_REPOSITORY_BINDING`
  - `RECENT_ACTIVITY_TYPE_MODULE_REPOSITORY_BINDING`
- **AND** 首个枚举值必须为 `UNSPECIFIED = 0`
- **AND** 不得保留笼统 `binding` 类型

#### Scenario: RecentActivity 的导航落点解释

- **WHEN** 合同设计消费 `RecentActivityItem.target_type / target_id`
- **THEN** 必须继续对齐 `phase05-03` 已冻结的导航解释：
  - `release` 活动统一回落到 `Module Detail`
  - `product_module_binding` 统一落到 `Product Detail`
  - `product_repository_binding` 与 `module_repository_binding` 统一落到 `Repository Binding Detail / Workspace`
- **AND** `target_type` 表示 canonical owner 页面目标，而不是活动类型本身的简单复制

### Requirement: chi + JSON HTTP 过渡映射必须从 proto 单向承接

系统 SHALL 冻结 `Dashboard + Feedback` 的 HTTP 过渡传输层与 `.proto` 的关系为“`.proto` 定义合同，`chi + JSON HTTP` 显式映射承接”，不允许再并列扩张第二套字段语义。

#### Scenario: RPC 到 HTTP 的映射矩阵

- **WHEN** 当前阶段讨论 Dashboard 三个 RPC 的 HTTP 过渡承接
- **THEN** 必须冻结为：
  - `GetDashboardOverview` → `GET /api/dashboard/overview`
  - `GetFeedbackSignals` → `GET /api/dashboard/feedback-signals`
  - `GetRecentActivities` → `GET /api/dashboard/recent-activities`
- **AND** 当前阶段三个入口都不承接 body、分页参数或排序切换参数
- **AND** 不得把 HTTP 路径、状态码或中间件策略误写成 `.proto` 合同本体

#### Scenario: JSON 映射边界

- **WHEN** 后续实现编写 `backend/internal/dashboard/types.go`、handler DTO 或前端 adapter
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 显式对齐
- **AND** 不得在 DTO、handler 或页面层私自新增 `.proto` 中不存在的业务字段语义
- **AND** 当前阶段错误语义继续由传输层映射承接，状态码细节不进入 `.proto` 合同本体

### Requirement: 合同演进与 Buf 校验规则必须冻结

系统 SHALL 将 `Dashboard + Feedback` 的合同演进规则冻结为与仓库现有 proto 主线一致的单值规则，确保字段编号、枚举语义与 breaking check 有唯一解释。

#### Scenario: 字段与枚举演进规则

- **WHEN** 当前阶段冻结 `dashboard.proto` 的字段编号与枚举编号
- **THEN** 新增字段必须使用新的递增编号
- **AND** 不得复用已删除字段编号或字段语义
- **AND** 每个业务枚举都必须包含首个 `UNSPECIFIED = 0`
- **AND** 不得通过修改既有字段类型、重复性或语义来做兼容性演进

#### Scenario: reserved 约束

- **WHEN** 后续版本删除字段、删除枚举值或废弃字段名
- **THEN** 必须使用 `reserved` 保留字段号
- **AND** 必要时必须同时保留字段名或枚举名
- **AND** 不得删除后再回收复用原编号

#### Scenario: Buf 校验链

- **WHEN** 后续实现者或 CI 校验 `Dashboard + Feedback` 合同
- **THEN** 必须继续复用仓库现有 `proto/Makefile`
- **AND** 最小校验链必须覆盖 `buf build`、`buf lint`、`buf generate`、`buf breaking`
- **AND** `buf breaking` 必须直接对照 `../.git#branch=main,subdir=proto`
- **AND** 不得吞掉 breaking 失败退出码

### Requirement: 当前版本字段编号方案必须冻结为机械映射

系统 SHALL 像 `phase03-08 / phase04-08` 一样，为 `Dashboard + Feedback` 当前版本冻结明确的枚举编号与消息字段编号方案，使 `phase05-09` 的 `.proto` 落地成为机械映射，而不是只冻结抽象原则后把真实编号留到实现期决定。该 Requirement 与上面的“合同演进与 Buf 校验规则”配合：前者冻结当前版具体编号方案，后者冻结未来演进规则，共同满足 `phase05-08` DoD 中“合同字段语义、字段编号与页面区块单值一致”。

#### Scenario: Feedback 相关枚举编号冻结

- **WHEN** 当前阶段定义 `FeedbackSignalFamily`、`FeedbackSignalCode`、`FeedbackSignalPriority` 与 `DashboardTargetType`
- **THEN** 它们必须按以下编号方案冻结：
  - `FeedbackSignalFamily`：`FEEDBACK_SIGNAL_FAMILY_UNSPECIFIED = 0` / `FEEDBACK_SIGNAL_FAMILY_PENDING_DECISION = 1` / `FEEDBACK_SIGNAL_FAMILY_PRODUCT_ASSET_COVERAGE = 2`
  - `FeedbackSignalCode`：`FEEDBACK_SIGNAL_CODE_UNSPECIFIED = 0` / `FEEDBACK_SIGNAL_CODE_PENDING_DECISION = 1` / `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS = 2` / `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING = 3` / `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING = 4`
  - `FeedbackSignalPriority`：`FEEDBACK_SIGNAL_PRIORITY_UNSPECIFIED = 0` / `FEEDBACK_SIGNAL_PRIORITY_P1_PENDING_DECISION = 1` / `FEEDBACK_SIGNAL_PRIORITY_P2_PRODUCT_MISSING_BOTH_BINDINGS = 2` / `FEEDBACK_SIGNAL_PRIORITY_P3_PRODUCT_MISSING_REPOSITORY_BINDING = 3` / `FEEDBACK_SIGNAL_PRIORITY_P4_PRODUCT_MISSING_MODULE_BINDING = 4`
  - `DashboardTargetType`：`DASHBOARD_TARGET_TYPE_UNSPECIFIED = 0` / `DASHBOARD_TARGET_TYPE_DECISION_DETAIL = 1` / `DASHBOARD_TARGET_TYPE_DECISION_LIST = 2` / `DASHBOARD_TARGET_TYPE_PRODUCT_DETAIL = 3` / `DASHBOARD_TARGET_TYPE_MODULE_DETAIL = 4` / `DASHBOARD_TARGET_TYPE_REPOSITORY_DETAIL = 5`
- **AND** 后续新增反馈相关枚举值必须使用新的递增编号，不得插入到已有编号之间

#### Scenario: DashboardOverview 相关消息编号冻结

- **WHEN** 当前阶段定义 `DashboardOverview` 与 `GetDashboardOverview` 读组消息
- **THEN** 它们必须按以下字段编号方案冻结：
  - `DashboardOverview`：`module_count=1(int32)` / `product_count=2(int32)` / `repository_count=3(int32)` / `decision_count=4(int32)` / `product_with_repository_count=5(int32)` / `product_with_module_count=6(int32)`
  - `GetDashboardOverviewRequest`：空消息（当前版本无字段）
  - `GetDashboardOverviewResponse`：`overview=1(DashboardOverview)`
- **AND** 后续若为 `GetDashboardOverviewRequest` 增加字段，只能从新的递增字段号开始，不得反向改写当前“空 request”边界

#### Scenario: FeedbackSignal 读组消息编号冻结

- **WHEN** 当前阶段定义 `FeedbackSignal`、`ProductAssetCoverageSummary` 与 `GetFeedbackSignals` 读组消息
- **THEN** 它们必须按以下字段编号方案冻结：
  - `FeedbackSignal`：`signal_family=1(FeedbackSignalFamily)` / `signal_code=2(FeedbackSignalCode)` / `priority=3(FeedbackSignalPriority)` / `title=4(string)` / `summary=5(string)` / `action_label=6(string)` / `target_type=7(DashboardTargetType)` / `target_id=8(string)` / `target_label=9(string)`
  - `ProductAssetCoverageSummary`：`fully_bound_product_count=1(int32)` / `missing_both_bindings_count=2(int32)` / `missing_repository_binding_count=3(int32)` / `missing_module_binding_count=4(int32)` / `representative_signals=5(repeated FeedbackSignal)`
  - `GetFeedbackSignalsRequest`：空消息（当前版本无字段）
  - `GetFeedbackSignalsResponse`：`current_focus_signals=1(repeated FeedbackSignal)` / `asset_feedback_summary=2(ProductAssetCoverageSummary)`
- **AND** `missing_both_bindings_count=2` 必须与 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS = 2` 一起作为当前阶段双缺口语义的单值锚点冻结
- **AND** 后续若为 `GetFeedbackSignalsRequest` 增加字段，只能从新的递增字段号开始，不得反向改写当前“空 request”边界

#### Scenario: RecentActivity 读组消息与枚举编号冻结

- **WHEN** 当前阶段定义 `RecentActivityType`、`RecentActivityItem` 与 `GetRecentActivities` 读组消息
- **THEN** 它们必须按以下编号方案冻结：
  - `RecentActivityType`：`RECENT_ACTIVITY_TYPE_UNSPECIFIED = 0` / `RECENT_ACTIVITY_TYPE_MODULE = 1` / `RECENT_ACTIVITY_TYPE_RELEASE = 2` / `RECENT_ACTIVITY_TYPE_PRODUCT = 3` / `RECENT_ACTIVITY_TYPE_REPOSITORY = 4` / `RECENT_ACTIVITY_TYPE_DECISION = 5` / `RECENT_ACTIVITY_TYPE_PRODUCT_MODULE_BINDING = 6` / `RECENT_ACTIVITY_TYPE_PRODUCT_REPOSITORY_BINDING = 7` / `RECENT_ACTIVITY_TYPE_MODULE_REPOSITORY_BINDING = 8`
  - `RecentActivityItem`：`activity_type=1(RecentActivityType)` / `activity_at=2(google.protobuf.Timestamp)` / `target_type=3(DashboardTargetType)` / `target_id=4(string)` / `target_label=5(string)`
  - `GetRecentActivitiesRequest`：空消息（当前版本无字段）
  - `GetRecentActivitiesResponse`：`activities=1(repeated RecentActivityItem)`
- **AND** `activity_at=2(google.protobuf.Timestamp)` 必须作为当前阶段唯一显式活动时间字段编号冻结
- **AND** 后续若为 `GetRecentActivitiesRequest` 增加字段，只能从新的递增字段号开始，不得反向改写当前“空 request”边界

## MODIFIED Requirements

### Requirement: phase05 Contract First 的 Dashboard 合同解释

`Contract First` 在 `phase05` 当前阶段 SHALL 被进一步解释为“先将 Dashboard 三类聚合读与反馈/活动模型冻结为单一 `.proto` 合同源，再由 `chi + JSON HTTP` 过渡层单向承接”，而不是继续由页面模型、HTTP DTO 或 handler adapter 并列定义字段语义。

#### Scenario: Contract First 承接方式

- **WHEN** 后续实现或验收讨论 `Dashboard + Feedback` 的合同主线
- **THEN** 必须优先引用 `proto/psco/dashboard/v1/dashboard.proto`
- **AND** 必须同时满足 `phase05-02 / 03 / 04 / 07` 已冻结的字段模板、导航语义、读取边界与后端 owner 规则
- **AND** 不得在 `.proto` 之外再发明并列合同源

## REMOVED Requirements

### Requirement: Dashboard 合同继续停留在 DTO / 页面模型层

**Reason**: `phase05-08` 的目标就是把 `Dashboard + Feedback` 从“前置规格已冻结字段和边界”推进到“仓库内有单一 `.proto` 合同设计入口”的状态。若继续让 DTO、页面类型或 handler 结构承担并列合同职责，后续实现会再次长出第二套语义。

**Migration**: 后续 `Dashboard + Feedback` 的消息结构、枚举、字段编号、HTTP 映射与 breaking 校验，都必须统一从 `proto/psco/dashboard/v1/dashboard.proto` 进入；`backend/internal/dashboard/types.go`、前端 adapter 与 JSON DTO 只允许作为从 `.proto` 单向映射出的过渡实现层存在。
