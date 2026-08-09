# Tasks

- [x] Task 1: 冻结 Dashboard 后端模块物理边界、Go package 命名与内部文件分层。把 Dashboard 代码落点、package 路径与文件分层写成单值结论。
  - [x] SubTask 1.1: 明确 Dashboard 模块物理落点为 `backend/internal/dashboard/`
  - [x] SubTask 1.2: 明确 Go package 路径为 `github.com/psco/backend/internal/dashboard`，沿用既有 `internal/<modulename>` 命名模式
  - [x] SubTask 1.3: 明确模块内部文件分层为 `handler/ / service/ / repository/ / candidate/ / types.go / errors.go / validate.go / handler/response.go`
  - [x] SubTask 1.4: 明确 `errors.go / types.go / validate.go / handler/response.go` 按职责单值化映射到唯一文件，禁止散落

- [x] Task 2: 冻结 Dashboard 三类聚合读取的接口分组与只读不写边界。把读接口分组、读组编排 owner 与写接口排除写成单值结论。
  - [x] SubTask 2.1: 明确读接口分组为 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 三组
  - [x] SubTask 2.2: 明确三个读组统一收口到 `backend/internal/dashboard/service/query_service.go` 的单一 `QueryService`，对应 `ReadOverview / ReadFeedbackSignal / ReadRecentActivity` 三个读组方法
  - [x] SubTask 2.3: 明确 Dashboard 模块不承接任何业务写入接口，写动作 canonical owner 继续由 canonical 模块承接
  - [x] SubTask 2.4: 明确具体路由路径、HTTP 方法、状态码细节由 `phase05-08` 承接

- [x] Task 3: 冻结 `DashboardOverviewRead` 的 QueryService 读组 owner 与跨模块读取边界。把主聚合读取的读组方法、reader 接口集合与整页失败语义写成单值结论。
  - [x] SubTask 3.1: 明确 `ReadOverview` 作为 `QueryService` 的读组方法，落点为 `backend/internal/dashboard/service/query_service.go`
  - [x] SubTask 3.2: 明确 reader 接口集合为 `ModuleCountReader / ProductCountReader / RepositoryCountReader / DecisionCountReader`
  - [x] SubTask 3.3: 明确 reader 接口定义与实现均由 Dashboard 模块 `candidate/` 子包自己拥有，沿用 phase03 `DecisionModuleCandidateRead` 模式
  - [x] SubTask 3.4: 明确原始聚合读允许由 Dashboard `candidate/` 自己拥有，但不得排除 canonical owner 已拥有的 provider-owned read 模式
  - [x] SubTask 3.5: 明确 Dashboard `candidate/` 实现可以直接读取 canonical 模块的表，但必须在 `candidate/` 子包内隔离；`service/` 层不得直接跨模块写 SQL
  - [x] SubTask 3.6: 明确任一 reader 失败时 `QueryService.ReadOverview` 必须返回整页失败，与 `phase05-04` 错误语义一致

- [x] Task 4: 冻结 `FeedbackSignalRead` 的 QueryService 读组 owner 与 `Feedback Signal Card` 归一化组装 owner。把反馈信号读取的读组方法、归一化归属、reader 接口集合与局部失败语义写成单值结论。
  - [x] SubTask 4.1: 明确 `ReadFeedbackSignal` 作为 `QueryService` 的读组方法，落点为 `backend/internal/dashboard/service/query_service.go`
  - [x] SubTask 4.2: 明确 `Feedback Signal Card` 归一化组装 owner 为 `QueryService.ReadFeedbackSignal` 内部，不让 canonical 模块产卡片
  - [x] SubTask 4.3: 明确 reader 接口集合为 `PendingDecisionSignalReader / ProductAssetCoverageReader`，并按读语义 owner 选择 Dashboard `candidate` 自拥有或 provider-owned read 模式
  - [x] SubTask 4.4: 明确已拥有源语义的 canonical owner read 不得被强制改写为 Dashboard `candidate` 复制实现
  - [x] SubTask 4.5: 明确任一 reader 失败时 `QueryService.ReadFeedbackSignal` 必须返回局部失败，不拖垮整页

- [x] Task 5: 冻结 `RecentActivityRead` 的 QueryService 读组 owner 与 `recent_activity` 类型映射 owner。把活动流读取的读组方法、类型映射归属、reader 接口集合与局部失败语义写成单值结论。
  - [x] SubTask 5.1: 明确 `ReadRecentActivity` 作为 `QueryService` 的读组方法，落点为 `backend/internal/dashboard/service/query_service.go`
  - [x] SubTask 5.2: 明确 `recent_activity` 类型映射 owner 为 `QueryService.ReadRecentActivity` 内部，不让 canonical 模块产统一活动流结构
  - [x] SubTask 5.3: 明确 reader 接口集合为 `ModuleActivityReader / ProductActivityReader / RepositoryActivityReader / DecisionActivityReader`，并按读语义 owner 选择受控跨模块读取模式
  - [x] SubTask 5.4: 明确各 reader 必须返回带显式 `activity_at` 字段的原始活动项
  - [x] SubTask 5.5: 明确 provider-owned read 可继续承接 canonical owner 已拥有的摘要读取，不被一刀切排除
  - [x] SubTask 5.6: 明确任一 reader 失败时 `QueryService.ReadRecentActivity` 必须返回局部失败，不拖垮整页

- [x] Task 6: 冻结 Dashboard 与四个 canonical 模块的服务侧连接边界、依赖注入装配点与跨模块读取 owner 归属原则。把跨模块读取方向、candidate reader 落点与装配点写成单值结论。
  - [x] SubTask 6.1: 明确跨模块读取只允许两类受控模式：Dashboard `candidate` 自拥有读，或 canonical owner 提供的 provider-owned read
  - [x] SubTask 6.2: 明确 Dashboard `service/` 层不得直接 import canonical 模块 `repository/` 包，不得直接跨模块写 SQL
  - [x] SubTask 6.3: 明确 Dashboard `candidate/` 子包的 reader 分工（Module / Decision / Product / Repository 四组）只覆盖 Dashboard 自己拥有的原始聚合读
  - [x] SubTask 6.4: 明确 canonical owner 已拥有的读语义允许继续通过 provider-owned read 参与装配，不被强制复制为 Dashboard `candidate` 实现
  - [x] SubTask 6.5: 明确依赖注入装配点为 `backend/internal/platform/` 路由装配层，沿用 phase03/04 装配模式
  - [x] SubTask 6.6: 明确接口演进规则：先判断读语义 owner，再决定新增 Dashboard `candidate` 还是复用 canonical owner read；不得在未冻结接口前让 `service/` 层直接跨模块写 SQL

- [x] Task 7: 冻结 DTO/合同映射策略、错误响应归属与局部失败传递边界。把 DTO owner、映射方式、错误响应统一入口与空态/失败语义写成单值结论。
  - [x] SubTask 7.1: 明确 DTO 落点为 `backend/internal/dashboard/types.go`
  - [x] SubTask 7.2: 明确 DTO 与 `.proto` 映射方式为显式映射函数，owner 在 Dashboard 模块，`.proto` 消息由 `phase05-08` 承接
  - [x] SubTask 7.3: 明确 DTO 字段语义对齐 `phase05-02 / phase05-04` 已冻结的最小字段模板
  - [x] SubTask 7.4: 明确错误响应统一入口为 `backend/internal/dashboard/handler/response.go`
  - [x] SubTask 7.5: 明确整页失败、局部失败、空态响应的传递规则与 `phase05-04 / phase05-06` 一致，且“局部失败”由调用侧基于三次独立请求组合派生，不在 handler 响应中新增专用标记

- [x] Task 8: 冻结不提前冻结边界，完成 phase05-07 规格一致性校验。确认本次冻结与 phase05 三件套、`phase05-02 / 04 / 06`、既有 phase02/03/04 后端模式保持一致，并足以直接进入实现。
  - [x] SubTask 8.1: 明确不冻结 Go 数据访问层具体工具（`sqlx / sqlc / GORM / database/sql`）
  - [x] SubTask 8.2: 明确不冻结缓存层、连接池配置、查询超时
  - [x] SubTask 8.3: 明确不冻结 `.proto` 服务命名、消息结构、字段编号、包名版本
  - [x] SubTask 8.4: 明确不冻结 `chi` 路由路径、HTTP 方法、状态码细节、reader 接口具体方法签名
  - [x] SubTask 8.5: 验证模块边界、文件分层、接口分组与 `phase05` 三件套、`phase05-04` 接口边界一致
  - [x] SubTask 8.6: 验证单一 `QueryService` 模式与 `phase02-09 / phase03-10 / phase04-10` 已落地的单文件 `service/query_service.go` 模式一致
  - [x] SubTask 8.7: 验证跨模块读取模式与 `phase03 / 04` 已验证的两类受控模式一致，未一刀切排除 provider-owned read
  - [x] SubTask 8.8: 验证错误语义与 `phase05-04 / phase05-06` 一致，未发明第二套
  - [x] SubTask 8.9: 验证设计结果足以直接进入实现

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`
- `Task 8` depends on `Task 1` through `Task 7`
