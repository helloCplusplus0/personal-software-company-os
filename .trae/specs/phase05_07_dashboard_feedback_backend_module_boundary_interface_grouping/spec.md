# Phase05-07 Dashboard/Feedback 后端模块边界与接口分组设计 Spec

## Why

`phase05-02` 已冻结 `Feedback Signal Card` 字段模板、优先级与展示模型，`phase05-04` 已冻结 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 的最小接口边界与错误语义前提，`phase05-05 / 06` 已冻结前端页面、组件、状态模型与交互流。但“接口边界停在哪、错误语义如何承接”还不等于“后端代码放在哪、谁拥有 query service、跨模块数据怎么读、Feedback Signal Card 由谁归一化、recent_activity 由谁映射”。如果不把 Dashboard 后端模块的物理边界、文件分层、接口分组、query service owner、跨模块读取方向、归一化与类型映射 owner、DTO 与合同映射策略、错误响应 owner 写成单值结论，后续 `.proto` 合同、Go 实现与验收会继续在“Dashboard 是独立模块还是塞进既有模块”“跨模块读取是直接 SQL 还是 reader 接口”“Feedback Signal Card 归一化在 Dashboard 还是在 canonical 模块”之间漂移。

> 阶段分工约束：本规格只冻结 Dashboard 后端模块的物理边界、文件分层、接口分组、query service owner、跨模块读取方向、归一化与类型映射 owner、DTO 与合同映射 owner、错误响应 owner。`.proto` 服务命名、消息结构、字段编号、包名版本与 `chi + JSON HTTP` 显式映射（路由路径、HTTP 方法、状态码细节、reader 接口具体方法签名）由 `phase05-08` 承接；Go 数据访问层具体工具（sqlx / sqlc / GORM / database/sql）、缓存层、连接池配置、查询超时不在本规格冻结。
>
> 与 `phase05-04` 的边界划分：`phase05-04` 已冻结三类聚合读的接口边界与错误语义前提；本规格承接这些结论，在此基础上冻结这些接口在后端如何落点为模块、service、handler 与 candidate reader，不重新发明第二套接口边界或第二套错误语义。

## What Changes

- 冻结 Dashboard 后端模块的物理边界与 Go package 命名
- 冻结 Dashboard 模块内部文件分层与支撑文件单值化映射
- 冻结 Dashboard 三类聚合读取的接口分组与只读不写边界
- 冻结 Dashboard 模块单一 `QueryService` 与三个读组方法（`ReadOverview / ReadFeedbackSignal / ReadRecentActivity`），不拆散为三套独立 service owner
- 冻结 `Feedback Signal Card` 归一化组装 owner 与 `recent_activity` 类型映射 owner
- 冻结 Dashboard 与四个 canonical 模块的服务侧连接边界、依赖注入装配点与跨模块读取 owner 归属原则
- 冻结跨模块读取沿用 `phase03 / phase04` 已验证的两类受控模式：`candidate` 自拥有读模式，以及 canonical owner 提供的 provider-owned read 模式
- 冻结 DTO/合同映射策略与 owner、错误响应归属与局部失败传递边界
- 明确当前阶段不提前冻结 Go 数据访问层具体工具、`.proto` 细节与 chi 路由显式映射

## Impact

- Affected specs: `phase05_dashboard_feedback_foundation`
- Affected code: 后续 `backend/internal/dashboard/` 模块的新增（handler / service / repository / candidate / types.go / errors.go / validate.go / response.go）
- Affected code: 后续 `backend/internal/dashboard/candidate/` 需新增 Dashboard 自拥有的原始聚合 reader，沿用 `phase03` 已验证的 `candidate` 自拥有模式
- Affected code: 后续 `backend/internal/dashboard/service/query_service.go` 需新增单一 `QueryService` 承接三个读组编排
- Affected code: 后续 `backend/internal/platform/` 路由装配层需新增 Dashboard 模块的 handler 注册与 candidate reader 依赖注入装配
- Affected code: 若复用 canonical owner 已拥有的源语义或摘要读取，现有 canonical 模块可继续通过 owner-owned read 参与装配，无需被强制改写成 Dashboard `candidate` 自实现
- Affected code: 后续 `proto/psco/dashboard/v1/` 的 `.proto` 落点（具体内容由 phase05-08 承接）

## ADDED Requirements

### Requirement: Dashboard 后端模块物理边界冻结

系统 SHALL 将 Dashboard 后端逻辑冻结为独立模块 `backend/internal/dashboard/`，不塞入既有任一 canonical 模块。

#### Scenario: Dashboard 模块的物理落点

- **WHEN** 后续实现讨论 Dashboard 后端代码的物理归属
- **THEN** 必须落点为 `backend/internal/dashboard/`
- **AND** 不得把 Dashboard 聚合读取逻辑塞入 `moduleregistry / decisioncenter / productregistry / repositorybinding` 任一既有模块
- **AND** 不得把 Dashboard 聚合读取逻辑塞入 `platform` 路由装配层

#### Scenario: Dashboard 模块的 Go package 命名

- **WHEN** 后续实现讨论 Dashboard 模块的 Go package 路径
- **THEN** 必须为 `github.com/psco/backend/internal/dashboard`
- **AND** 必须沿用既有 `internal/<modulename>` 单值命名模式
- **AND** 不得引入 `dashboardapi / dashboardquery / dashboardfeature / dashboardservice` 等第二套命名

### Requirement: Dashboard 模块内部文件分层冻结

系统 SHALL 将 Dashboard 模块内部文件分层冻结为沿用 `phase02-09 / phase03-10 / phase04-10` 已验证的 `handler / service / repository / candidate / types.go / errors.go / validate.go / handler/response.go` 结构。

#### Scenario: Dashboard 模块文件分层

- **WHEN** 后续实现讨论 Dashboard 模块内部文件组织
- **THEN** 必须至少承接以下分层：
  - `handler/` 承接 HTTP 入口与参数解析
  - `service/` 承接聚合读取编排与 `Feedback Signal Card` 归一化
  - `repository/` 承接 Dashboard 自有的派生查询（如有）
  - `candidate/` 承接跨模块读取接口定义与 owner 归属
  - `types.go` 承接 DTO 与领域类型
  - `errors.go` 承接业务错误哨兵值
  - `validate.go` 承接入参校验
  - `handler/response.go` 承接统一 JSON 响应
- **AND** 必须沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的分层模式
- **AND** 不得为 Dashboard 单独发明第二套文件分层命名

#### Scenario: 支撑文件单值化映射

- **WHEN** 后续实现新增 Dashboard 模块的支撑文件
- **THEN** `errors.go / types.go / validate.go / handler/response.go` 必须按职责单值化映射到唯一文件
- **AND** 禁止把 errors / types / validate / response 逻辑散落在 handler 或 service 内部
- **AND** 必须沿用 phase04 已落地的“支撑文件单值化”约束

### Requirement: Dashboard 三类聚合读取的接口分组冻结

系统 SHALL 将 Dashboard 模块的接口分组冻结为只读三组：`DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead`。

#### Scenario: Dashboard 读接口分组

- **WHEN** 后续实现讨论 Dashboard 模块的接口分组
- **THEN** 必须冻结为三个只读聚合接口：
  - `DashboardOverviewRead` 承接主聚合读取
  - `FeedbackSignalRead` 承接反馈信号附属聚合读取
  - `RecentActivityRead` 承接最近活动附属聚合读取
- **AND** 三个读取分别对应 `service/` 层三个独立 service 方法
- **AND** 三个读取分别对应 `handler/` 层三个独立 HTTP endpoint
- **AND** 不得并列增加趋势分析、外部埋点、通知中心或导出接口
- **AND** 具体路由路径、HTTP 方法、状态码细节由 `phase05-08` 承接，不在本规格冻结

#### Scenario: Dashboard 模块不承接写接口

- **WHEN** 后续实现讨论 Dashboard 模块的接口范围
- **THEN** Dashboard 模块当前阶段只承接读组
- **AND** 不得在 Dashboard 模块新增任何业务写入接口
- **AND** 不得把 `BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository / RecordDecision` 等写动作重新塞回 Dashboard
- **AND** 写动作的 canonical owner 继续由 `Product Registry / Repository Binding / Decision Center` 承接

### Requirement: DashboardOverviewRead query service owner 与跨模块读取边界冻结

系统 SHALL 将 `DashboardOverviewRead` 的 query service owner 冻结为 Dashboard 模块单一 `QueryService` 的 `ReadOverview` 读组，并通过 Dashboard 模块 `candidate/` 子包定义并拥有的 reader 接口与实现承接跨模块读取。

#### Scenario: DashboardOverviewRead 的 service owner

- **WHEN** 后续实现讨论 `DashboardOverviewRead` 的 query service owner
- **THEN** 必须为 `backend/internal/dashboard/service/query_service.go` 中的单一 `QueryService`
- **AND** `ReadOverview` 作为 `QueryService` 的一个读组方法，负责编排 `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count` 六个计数的聚合
- **AND** 必须沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的单文件 `service/query_service.go` 统一承接读组编排模式
- **AND** 不得为该读组单独发明 `OverviewReadService` 等独立 service owner
- **AND** 不得把该读组逻辑下沉到 `handler/` 或 `repository/` 层

#### Scenario: DashboardOverviewRead 的跨模块读取边界

- **WHEN** `DashboardOverviewRead` 需要读取四个 canonical 模块的计数
- **THEN** 必须通过 Dashboard 模块 `candidate/` 子包定义并拥有的 reader 接口与实现承接
- **AND** reader 接口至少包括：
  - `ModuleCountReader` 返回 `module_count`
  - `ProductCountReader` 返回 `product_count / product_with_module_count / product_with_repository_count`
  - `RepositoryCountReader` 返回 `repository_count`
  - `DecisionCountReader` 返回 `decision_count`
- **AND** reader 接口的定义与实现均由 Dashboard 模块 `candidate/` 子包自己拥有，沿用 `phase03` 已验证的 `DecisionModuleCandidateRead` 模式
- **AND** canonical 模块不需要为 Dashboard 新增 candidate 实现
- **AND** Dashboard `candidate/` 实现可以直接读取 canonical 模块的表，但必须在 `candidate/` 子包内隔离
- **AND** Dashboard `service/` 层不得直接跨模块写 SQL，跨模块读取必须通过 `candidate/` 子包暴露的 reader 接口承接
- **AND** reader 接口的具体方法签名由 `phase05-08` 承接

#### Scenario: DashboardOverviewRead 整页失败语义的服务侧解释

- **WHEN** `DashboardOverviewRead` 任一组成 reader 失败
- **THEN** `QueryService.ReadOverview` 必须返回整页失败语义
- **AND** 不得在主聚合读取层吞掉局部失败后返回部分计数
- **AND** 该失败语义必须与 `phase05-04` 已冻结的“主聚合失败触发整页失败”一致

### Requirement: FeedbackSignalRead query service owner 与 Feedback Signal Card 归一化组装 owner 冻结

系统 SHALL 将 `FeedbackSignalRead` 的 query service owner 与 `Feedback Signal Card` 归一化组装 owner 冻结为 Dashboard 模块单一 `QueryService` 的 `ReadFeedbackSignal` 读组内部。

#### Scenario: FeedbackSignalRead 的 service owner

- **WHEN** 后续实现讨论 `FeedbackSignalRead` 的 query service owner
- **THEN** 必须为 `backend/internal/dashboard/service/query_service.go` 中的单一 `QueryService`
- **AND** `ReadFeedbackSignal` 作为 `QueryService` 的一个读组方法，负责编排 `pending_decision_signals` 与 `product_asset_coverage` 的读取、归一化、排序与裁剪
- **AND** 必须沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的单文件 `service/query_service.go` 统一承接读组编排模式
- **AND** 不得为该读组单独发明 `FeedbackSignalReadService` 等独立 service owner
- **AND** 不得把该读组逻辑拆散到 `Decision Center` 或 `Product Registry` 各自实现

#### Scenario: Feedback Signal Card 归一化组装 owner

- **WHEN** 后续实现讨论 `Feedback Signal Card` 的归一化组装逻辑归属
- **THEN** 必须冻结为 Dashboard 模块 `QueryService.ReadFeedbackSignal` 内部
- **AND** `pending_decision_signals` 与 `product_asset_coverage` 作为原始数据来源，通过 Dashboard `candidate/` 子包的 reader 接口承接
- **AND** 归一化为统一 `Feedback Signal Card` 的逻辑必须在 Dashboard 模块内
- **AND** 不得让各 canonical 模块自己产 `Feedback Signal Card`
- **AND** 不得把归一化逻辑分散到 `handler/` 或 `repository/` 层

#### Scenario: FeedbackSignalRead 跨模块读取边界

- **WHEN** `QueryService.ReadFeedbackSignal` 需要读取 `pending_decision_signals` 与 `product_asset_coverage`
- **THEN** 必须通过 Dashboard 模块 `candidate/` 子包定义并拥有的 reader 接口与实现承接
- **AND** reader 接口至少包括：
  - `PendingDecisionSignalReader` 返回原始待决策信号
  - `ProductAssetCoverageReader` 返回原始资产缺口数据
- **AND** reader 接口的定义与实现均由 Dashboard 模块 `candidate/` 子包自己拥有，沿用 `phase03` 已验证的 `DecisionModuleCandidateRead` 模式
- **AND** canonical 模块不需要为 Dashboard 新增 candidate 实现
- **AND** Dashboard `candidate/` 实现可以直接读取 `Decision Center` 与 `Product Registry` 的表，但必须在 `candidate/` 子包内隔离
- **AND** Dashboard `service/` 层不得直接跨模块写 SQL

#### Scenario: FeedbackSignalRead 局部失败语义的服务侧解释

- **WHEN** `FeedbackSignalRead` 的任一 reader 失败
- **THEN** `QueryService.ReadFeedbackSignal` 必须返回局部失败语义
- **AND** 不得拖垮 `DashboardOverviewRead` 的整页成功语义
- **AND** 该失败语义必须与 `phase05-04` 已冻结的“附属聚合失败只触发局部失败”一致

### Requirement: RecentActivityRead query service owner 与 recent_activity 类型映射 owner 冻结

系统 SHALL 将 `RecentActivityRead` 的 query service owner 与 `recent_activity` 类型映射 owner 冻结为 Dashboard 模块单一 `QueryService` 的 `ReadRecentActivity` 读组内部。

#### Scenario: RecentActivityRead 的 service owner

- **WHEN** 后续实现讨论 `RecentActivityRead` 的 query service owner
- **THEN** 必须为 `backend/internal/dashboard/service/query_service.go` 中的单一 `QueryService`
- **AND** `ReadRecentActivity` 作为 `QueryService` 的一个读组方法，负责编排多个 canonical 模块的活动项读取、类型映射、归并与按 `activity_at` 排序
- **AND** 必须沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的单文件 `service/query_service.go` 统一承接读组编排模式
- **AND** 不得为该读组单独发明 `RecentActivityReadService` 等独立 service owner
- **AND** 不得把该读组逻辑拆散到各 canonical 模块各自实现

#### Scenario: recent_activity 类型映射 owner

- **WHEN** 后续实现讨论 `module / release / product / repository / decision / product_module_binding / product_repository_binding / module_repository_binding` 等活动项的类型映射逻辑归属
- **THEN** 必须冻结为 Dashboard 模块 `QueryService.ReadRecentActivity` 内部
- **AND** 各 canonical 模块的原始活动数据通过 Dashboard `candidate/` 子包的 reader 接口承接
- **AND** 统一映射为 `activity_type / activity_at / target_type / target_id / target_label` 的逻辑必须在 Dashboard 模块内
- **AND** 不得让各 canonical 模块自己产统一活动流结构
- **AND** 不得把类型映射逻辑分散到 `handler/` 或 `repository/` 层

#### Scenario: RecentActivityRead 跨模块读取边界

- **WHEN** `QueryService.ReadRecentActivity` 需要读取多个 canonical 模块的最近活动
- **THEN** 必须通过 Dashboard 模块 `candidate/` 子包定义并拥有的 reader 接口与实现承接
- **AND** reader 接口至少包括：
  - `ModuleActivityReader` 返回 `Module` 与 `Release` 的最近活动
  - `ProductActivityReader` 返回 `Product` 与 `product_module_binding` 的最近活动
  - `RepositoryActivityReader` 返回 `Repository` 与 `product_repository_binding / module_repository_binding` 的最近活动
  - `DecisionActivityReader` 返回 `Decision` 的最近活动
- **AND** 各 reader 必须返回带有显式 `activity_at` 字段的原始活动项
- **AND** reader 接口的定义与实现均由 Dashboard 模块 `candidate/` 子包自己拥有，沿用 `phase03` 已验证的 `DecisionModuleCandidateRead` 模式
- **AND** canonical 模块不需要为 Dashboard 新增 candidate 实现
- **AND** Dashboard `candidate/` 实现可以直接读取 canonical 模块的表，但必须在 `candidate/` 子包内隔离
- **AND** Dashboard `service/` 层不得直接跨模块写 SQL
- **AND** 当前阶段允许在归并与排序阶段裁剪到最多 `10` 条活动项

#### Scenario: RecentActivityRead 局部失败语义的服务侧解释

- **WHEN** `RecentActivityRead` 的任一 reader 失败
- **THEN** `QueryService.ReadRecentActivity` 必须返回局部失败语义
- **AND** 不得拖垮 `DashboardOverviewRead` 的整页成功语义
- **AND** 该失败语义必须与 `phase05-04` 已冻结的“附属聚合失败只触发局部失败”一致

### Requirement: Dashboard 与四个 canonical 模块的服务侧连接边界冻结

系统 SHALL 将 Dashboard 与 `Module Registry / Decision Center / Product Registry / Repository Binding` 的服务侧连接边界冻结为两类受控读模式，而不是一刀切只允许一种跨模块读取架构。

#### Scenario: 跨模块读取方向

- **WHEN** 后续实现讨论 Dashboard 如何读取四个 canonical 模块的数据
- **THEN** 必须只允许以下两类已验证模式：
  - Dashboard `candidate/` 子包自己定义并拥有接口与实现，用于 Dashboard 自己拥有的原始聚合读取
  - canonical owner 提供 provider-owned read，由 Dashboard `QueryService` 通过接口注入消费，用于 canonical owner 已拥有的源语义或摘要读取
- **AND** 不得在这两类模式之外再发明第三套跨模块读取架构
- **AND** 不得让 Dashboard `service/` 层直接 import canonical 模块的 `repository/` 包
- **AND** 不得让 Dashboard `service/` 层直接跨模块写 SQL
- **AND** 具体采用哪一种模式，必须由读语义 owner 决定，而不是按模块名一刀切

#### Scenario: Dashboard candidate 自拥有读的适用范围

- **WHEN** 某个 Dashboard 读取属于 Dashboard 自己拥有的原始聚合读
- **AND** 当前仓库中不存在已冻结的 canonical owner read 可直接复用
- **THEN** 可以由 Dashboard `candidate/` 子包自己定义并拥有接口与实现
- **AND** 该模式沿用 `phase03` 已验证的 `DecisionModuleCandidateRead` 模式
- **AND** 实现可以直接读取 canonical 模块的表，但必须在 Dashboard `candidate/` 子包内隔离

#### Scenario: provider-owned read 的适用范围

- **WHEN** 某个 Dashboard 读取要复用 canonical owner 已拥有的源语义或摘要读取
- **THEN** 允许继续采用 provider-owned read 模式
- **AND** Dashboard 通过接口注入消费该读能力
- **AND** 不得为了形式统一而强制把这类读取重写成 Dashboard `candidate` 自实现
- **AND** 该模式与仓库现有 `productregistry.QueryService` 注入 `BoundRepositoryReader` 的接线方式一致

#### Scenario: Dashboard candidate reader 落点与分工

- **WHEN** Dashboard 的 `candidate/` 子包承接 Dashboard 自己拥有的原始聚合 reader
- **THEN** reader 接口定义与实现必须落点在 Dashboard 模块自己的 `candidate/` 子包
- **AND** `candidate/` 至少承接以下 reader 分工：
  - Module 相关：`ModuleCountReader / ModuleActivityReader`
  - Decision 相关：`DecisionCountReader / PendingDecisionSignalReader / DecisionActivityReader`
  - Product 相关：`ProductCountReader / ProductAssetCoverageReader / ProductActivityReader`
  - Repository 相关：`RepositoryCountReader / RepositoryActivityReader`
- **AND** reader 实现可以直接读取 canonical 模块的表，但必须在 Dashboard `candidate/` 子包内隔离
- **AND** 该落点必须沿用 `phase03` 已验证的 `decisioncenter/candidate/module_candidate_read.go` 模式
- **AND** 具体接口签名由 `phase05-08` 承接，不在本规格冻结

#### Scenario: 依赖注入装配点

- **WHEN** 后续实现讨论 Dashboard `QueryService` 如何获得跨模块读依赖
- **THEN** 装配必须发生在 `backend/internal/platform/` 路由装配层
- **AND** `platform` 层负责构造 Dashboard `candidate/` 子包的 reader 实现，或装配 canonical owner 提供的 provider-owned read 实现
- **AND** `platform` 层负责把这些读依赖注入到 Dashboard `QueryService` 构造函数
- **AND** 必须沿用 `phase03 / phase04` 已落地的 platform 装配模式
- **AND** 不得让 Dashboard 模块在内部自行 new 出跨模块读依赖

### Requirement: 跨模块读取接口的 owner 归属原则冻结

系统 SHALL 将跨模块读取接口的 owner 归属原则冻结为两类受控模式下的单值选择规则，而不是把所有读取强行归并为一种实现 owner。

#### Scenario: 接口定义 owner

- **WHEN** 后续实现讨论跨模块 reader 接口的定义归属
- **THEN** 若采用 Dashboard `candidate` 自拥有读模式，接口定义必须由消费方模块（Dashboard）的 `candidate/` 子包拥有
- **AND** 若采用 provider-owned read 模式，接口定义应继续归属于 canonical owner 已拥有的读契约
- **AND** 不得在同一条读语义上同时并列维护两套接口定义

#### Scenario: 接口实现 owner

- **WHEN** 后续实现讨论跨模块 reader 接口的实现归属
- **THEN** 若采用 Dashboard `candidate` 自拥有读模式，接口实现必须由消费方模块（Dashboard）的 `candidate/` 子包自己拥有
- **AND** 若采用 provider-owned read 模式，接口实现继续由 canonical owner 自己拥有
- **AND** 无论采用哪一种模式，Dashboard `service/` 层都不得直接跨模块写 SQL

#### Scenario: 接口演进规则

- **WHEN** Dashboard 后续需要新增跨模块读取能力
- **THEN** 必须先判断该读取是 Dashboard 自己拥有的原始聚合读，还是 canonical owner 已拥有的源语义/摘要读取
- **AND** 只有前者才允许在 Dashboard 的 `candidate/` 子包新增或扩展接口定义与实现
- **AND** 后者应优先复用 canonical owner 已拥有的 provider-owned read，而不是复制为第二份 Dashboard `candidate` 实现
- **AND** 不得在未冻结接口前让 `service/` 层直接跨模块写 SQL

### Requirement: Dashboard 模块只读不写边界冻结

系统 SHALL 将 Dashboard 模块的职责边界冻结为只承接读，不承接任何业务写入。

#### Scenario: Dashboard 不承接写

- **WHEN** 后续实现讨论 Dashboard 模块的职责
- **THEN** Dashboard 模块只承接 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 三个只读聚合接口
- **AND** 不得在 Dashboard 模块新增 `handler/command_*.go` 或 `service/command_*.go` 等写组文件
- **AND** 不得让 Dashboard 模块直接修改 `modules / releases / products / repositories / decisions / decision_links / product_modules / product_repositories / module_repositories` 任一表
- **AND** 写动作的 canonical owner 继续由 `Product Registry / Repository Binding / Decision Center` 承接

#### Scenario: Dashboard 跳转不等于 Dashboard 写入

- **WHEN** 用户从 Dashboard 跳到 canonical owner 页面完成写入
- **THEN** 该写入必须由 canonical owner 模块的 service 承接
- **AND** Dashboard 模块不得承接任何写入副作用
- **AND** 该原则与 `phase05-04` 已冻结的“Dashboard 不承接写入接口”一致

### Requirement: DTO/合同映射策略与 owner 冻结

系统 SHALL 将 Dashboard 模块的 DTO/合同映射策略冻结为“DTO 在 `types.go`，映射函数 owner 在 Dashboard 模块，`.proto` 消息由 `phase05-08` 承接”。

#### Scenario: DTO owner

- **WHEN** 后续实现讨论 Dashboard 模块的 DTO 归属
- **THEN** DTO 必须落点在 `backend/internal/dashboard/types.go`
- **AND** DTO 字段语义必须与 `phase05-02 / phase05-04` 已冻结的最小字段模板单值一致
- **AND** 不得在 `handler/ / service/ / repository/` 内部散落 DTO 定义

#### Scenario: DTO 与 .proto 映射策略

- **WHEN** 后续实现讨论 Dashboard DTO 与 `.proto` 消息的映射
- **THEN** 映射必须通过显式映射函数实现，不得依赖反射或隐式序列化
- **AND** 映射函数 owner 归属 Dashboard 模块
- **AND** `.proto` 消息结构、字段编号、服务命名由 `phase05-08` 承接
- **AND** 当前阶段只冻结“映射 owner 在 Dashboard 模块、映射方式为显式函数”两条原则

#### Scenario: DTO 字段语义对齐

- **WHEN** 后续实现编写 Dashboard DTO
- **THEN** `Feedback Signal Card` DTO 至少必须承接 `signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label`
- **AND** `Recent Activity Item` DTO 至少必须承接 `activity_type / activity_at / target_type / target_id / target_label`
- **AND** `Dashboard Overview` DTO 至少必须承接 `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count`
- **AND** 不得减少上述字段
- **AND** 不得在当前阶段额外引入 `score / trend / external_metric / recommendation_reason`

### Requirement: 错误响应归属与局部失败传递边界冻结

系统 SHALL 将 Dashboard 模块的错误响应归属与局部失败传递边界冻结为沿用既有 `handler/response.go` 模式，并显式区分整页失败与局部失败。

#### Scenario: 错误响应统一入口

- **WHEN** Dashboard handler 需要返回错误响应
- **THEN** 必须通过 `backend/internal/dashboard/handler/response.go` 统一返回
- **AND** 必须沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的 JSON 错误响应格式
- **AND** 不得在 handler 内部散落错误响应构造逻辑

#### Scenario: 整页失败传递

- **WHEN** `DashboardOverviewRead` 返回失败
- **THEN** handler 必须返回整页失败响应
- **AND** 不得返回部分计数或伪装成功
- **AND** 该语义与 `phase05-04 / phase05-06` 已冻结的整页失败语义一致

#### Scenario: 局部失败传递

- **WHEN** `FeedbackSignalRead` 或 `RecentActivityRead` 返回失败
- **THEN** 各自的 handler 必须继续返回既有统一错误响应
- **AND** 不得把附属读失败升级为 `DashboardOverviewRead` 的整页失败
- **AND** Dashboard 页面级“主聚合成功、附属聚合局部失败”语义必须由调用侧基于三次独立请求结果组合派生
- **AND** 当前阶段不得在附属读 endpoint 的响应包络里额外发明“局部失败标记”
- **AND** 该语义与 `phase05-04 / phase05-06` 已冻结的局部失败语义一致

#### Scenario: 空态响应不映射为错误

- **WHEN** `DashboardOverviewRead` 返回零计数
- **OR** `FeedbackSignalRead` / `RecentActivityRead` 返回空列表
- **THEN** handler 必须返回成功响应
- **AND** 不得把空态映射为 `404` 或 `500`
- **AND** 该语义与 `phase05-04` 已冻结的空态语义一致

### Requirement: 不提前冻结 Go 数据访问层具体工具

系统 SHALL 明确当前阶段不冻结 Go 数据访问层具体工具与 `.proto` / `chi` 显式映射细节，避免过早绑定实现细节。

#### Scenario: 当前阶段允许冻结的内容

- **WHEN** 当前 spec 讨论 Dashboard 后端模块边界
- **THEN** 可以冻结模块物理边界、文件分层、接口分组、service owner、reader 接口 owner、归一化 owner、类型映射 owner、DTO owner、错误响应 owner、跨模块读取方向

#### Scenario: 当前阶段不得冻结的内容

- **WHEN** 后续实现尚未开始
- **THEN** 不得提前冻结 Go 数据访问层具体工具（如 `sqlx / sqlc / GORM / database/sql` 的选择）
- **AND** 不得提前冻结缓存层、连接池配置、查询超时
- **AND** 不得提前冻结 `.proto` 服务命名、消息结构、字段编号、包名版本
- **AND** 不得提前冻结 `chi` 路由路径、HTTP 方法、状态码细节
- **AND** 不得提前冻结 reader 接口的具体方法签名与参数类型
- **AND** 上述内容由 `phase05-08` 与后续实现阶段承接

## MODIFIED Requirements

### Requirement: phase05-04 聚合读接口边界的服务侧解释

`phase05-04` 已冻结 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 的最小接口边界与错误语义前提，`phase05-07` 在此基础上 SHALL 将其进一步解释为后端模块边界、query service owner、跨模块读取方向与归一化组装 owner。

#### Scenario: 接口边界到服务侧的解释

- **WHEN** 后续实现讨论 Dashboard 三类聚合读取的后端实现
- **THEN** 必须同时满足 `phase05-04` 的接口边界与错误语义前提
- **AND** 必须同时满足本规格冻结的模块边界、service owner 与跨模块读取方向
- **AND** 不得在后端额外发明第二套“局部失败也整页失败”或“反馈失败仍猜测主 CTA”的实现

### Requirement: phase03 跨模块候选读取模式的 Dashboard 应用解释

`phase03-10` 已冻结 `DecisionModuleCandidateRead` 模式（消费方 `candidate/` 子包自己定义并拥有接口与实现，`service/` 层不直接写跨模块 SQL），`phase04-07` 又已冻结 canonical owner 提供 provider-owned read 的注入复用模式。`phase05-07` 在此基础上 SHALL 将两类已验证模式应用到 Dashboard 与四个 canonical 模块之间，并按读语义 owner 选择单值模式。

#### Scenario: 跨模块读取模式应用

- **WHEN** 后续实现讨论 Dashboard 与四个 canonical 模块的跨模块读取
- **THEN** 必须沿用 `phase03 / phase04` 已验证的两类受控模式：
  - Dashboard 自己拥有的原始聚合读，沿用 `phase03-10` 已验证的“消费方 `candidate/` 子包自己定义并拥有接口与实现”模式
  - canonical owner 已拥有的源语义或摘要读取，沿用 `phase04-07` 已验证的 provider-owned read 注入复用模式
- **AND** 必须沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的单文件 `service/query_service.go` 统一承接读组编排模式
- **AND** provider-owned read 属于仓库已验证的受控模式，不得被误判为 Dashboard 额外发明的第三套或第二套未冻结架构
- **AND** 不得为 Dashboard 拆散为多个独立 service owner
- **AND** Dashboard `service/` 层不得直接跨模块写 SQL
- **AND** 若采用 Dashboard `candidate` 自拥有读模式，跨模块 SQL 必须在 `candidate/` 子包内隔离
- **AND** 若采用 provider-owned read 模式，Dashboard 只能通过注入的 canonical owner 读契约消费，不得在 Dashboard 内复制第二份实现

## REMOVED Requirements

### Requirement: Dashboard service 层直接跨模块 SQL

**Reason**: Dashboard `service/` 层直接跨模块写 SQL 会破坏 `phase03-10` 已验证的“跨模块读取在 `candidate/` 子包内隔离”原则，导致 Dashboard `service/` 与四个 canonical 模块的内部表结构强耦合，并偏离 `phase02-09 / phase03-10 / phase04-10` 已落地的单文件 `service/query_service.go` 统一承接读组编排模式。

**Migration**: Dashboard `service/` 层不得直接跨模块写 SQL。对于 Dashboard 自己拥有的原始聚合读，继续通过自身 `candidate/` 子包暴露的 reader 接口承接，跨模块 SQL 只允许出现在 Dashboard `candidate/` 子包内部；对于 canonical owner 已拥有的源语义或摘要读取，继续通过注入的 provider-owned read 契约消费，不得在 Dashboard 内复制第二份实现。

### Requirement: Dashboard 承接业务写入

**Reason**: `phase05-04` 已冻结 Dashboard 不承接任何新的业务写入接口，`phase05-06` 已冻结 Dashboard 跳转不等于 Dashboard 写入。如果 Dashboard 模块承接写入，会形成与 canonical owner 平行的第二套写入入口，违反 `phase05` 已冻结的交互归属原则。

**Migration**: 写动作的 canonical owner 继续由 `Product Registry / Repository Binding / Decision Center` 承接，Dashboard 模块只承接读组。
