# Tasks

- [x] Task 1: 对齐 phase05-12 直接上游与现有后端模块模式
  - [x] SubTask 1.1: 对齐 dev_plan phase05-12 范围与 DoD
  - [x] SubTask 1.2: 对齐 phase05-07 后端模块边界、文件分层与接口分组
  - [x] SubTask 1.3: 对齐 phase05-08/11 已落地 proto 合同的字段编号与枚举
  - [x] SubTask 1.4: 对齐 phase05-04 聚合读范围、接口边界与错误语义
  - [x] SubTask 1.5: 对齐 phase05-09 验收环境、九类 fixture 与局部错误模拟机制
  - [x] SubTask 1.6: 对齐现有 productregistry / decisioncenter 模块分层模式与 candidate 自拥有读模式

- [x] Task 2: 实现 Dashboard DTO 与领域类型
  - [x] SubTask 2.1: 实现 `backend/internal/dashboard/types.go`，承接 DashboardOverview / FeedbackSignal / ProductAssetCoverageSummary / RecentActivityItem DTO
  - [x] SubTask 2.2: 实现枚举类型（FeedbackSignalFamily / FeedbackSignalCode / FeedbackSignalPriority / DashboardTargetType / RecentActivityType）使用 string 类型与 snake_case 常量
  - [x] SubTask 2.3: 实现 `backend/internal/dashboard/errors.go` 承接业务错误哨兵值

- [x] Task 3: 实现 Dashboard candidate reader
  - [x] SubTask 3.1: 实现 `candidate/overview_readers.go` 承接 ModuleCountReader / ProductCountReader / RepositoryCountReader / DecisionCountReader
  - [x] SubTask 3.2: 实现 `candidate/feedback_readers.go` 承接 PendingDecisionSignalReader / ProductAssetCoverageReader
  - [x] SubTask 3.3: 实现 `candidate/activity_readers.go` 承接 ModuleActivityReader / ProductActivityReader / RepositoryActivityReader / DecisionActivityReader
  - [x] SubTask 3.4: 在 candidate reader 中承接环境变量局部错误模拟机制

- [x] Task 4: 实现 Dashboard QueryService
  - [x] SubTask 4.1: 实现 `service/query_service.go` 单一 QueryService 结构体与构造函数
  - [x] SubTask 4.2: 实现 ReadOverview 方法，编排四个计数 reader，任一失败返回整页失败
  - [x] SubTask 4.3: 实现 ReadFeedbackSignal 方法，编排信号 reader，归一化为统一 FeedbackSignal 列表与 ProductAssetCoverageSummary
  - [x] SubTask 4.4: 实现 ReadRecentActivity 方法，编排四个活动 reader，映射为统一 RecentActivityItem 列表，按 activity_at 倒序排序，最多 10 条
  - [x] SubTask 4.5: 实现反馈信号归一化与优先级排序（P1 > P2 > P3 > P4，同优先级按 created_at DESC）

- [x] Task 5: 实现 Dashboard handler
  - [x] SubTask 5.1: 实现 `handler/response.go` 承接 writeJSON / writeError
  - [x] SubTask 5.2: 实现 `handler/query_handler.go` 承接三个 GET endpoint
  - [x] SubTask 5.3: 注册 GET /api/dashboard/overview、GET /api/dashboard/feedback-signals、GET /api/dashboard/recent-activities 路由
  - [x] SubTask 5.4: 实现空态成功响应（200 + 空列表/零计数）与错误响应（500）

- [x] Task 6: 更新 platform 装配层
  - [x] SubTask 6.1: 在 `platform/router.go` 新增 buildDashboard 构造 candidate reader 与 QueryService
  - [x] SubTask 6.2: 在 `platform/router.go` 新增 mountDashboard 注册路由
  - [x] SubTask 6.3: 在 `platform/server.go` 的 NewServer 中调用 buildDashboard + mountDashboard

- [x] Task 7: 实现数据库脚本与 fixture
  - [x] SubTask 7.1: 实现 `database/scripts/reset_dashboard_acceptance.sh` 支持默认/--clean-only/--restore-only/--fixture <name> 四种模式
  - [x] SubTask 7.2: 实现 `database/seeds/seed_dashboard_acceptance_baseline.sql` 基线数据
  - [x] SubTask 7.3: 实现九类 fixture SQL 文件（empty-system / modules-only / products-without-modules / products-missing-repository / products-missing-module / pending-decisions / recent-activities / products-missing-all-repositories / products-missing-both-bindings）

- [x] Task 8: 验证后端可编译与可运行
  - [x] SubTask 8.1: 验证 `go build ./...` 通过
  - [x] SubTask 8.2: 验证三个 GET endpoint 可运行
  - [x] SubTask 8.3: 验证 reset_dashboard_acceptance.sh 可重复执行

- [x] Task 9: 修复独立复核发现的 2 个阻断性问题
  - [x] SubTask 9.1: 修复响应包络未对齐 proto envelope — 在 types.go 新增 `DashboardOverviewReadResult` 与 `RecentActivityReadResult` 响应包络 DTO
  - [x] SubTask 9.2: 修复 handler GetOverview 与 GetRecentActivities 返回包络结构，不再返回裸 DashboardOverview / 裸 []RecentActivityItem
  - [x] SubTask 9.3: 修复 reset_dashboard_acceptance.sh host psql 模式缺失 PGPASSWORD 处理 — 补齐与既有 reset 脚本一致的密码解析段
  - [x] SubTask 9.4: 同步更新 spec.md 补充"响应包络必须对齐 proto envelope"与"host psql 模式密码处理必须对齐既有 reset 脚本"两条 scenario
  - [x] SubTask 9.5: 验证 go build / go vet / bash -n 全部通过
  - [x] SubTask 9.6: 运行时验证三个 endpoint 响应 JSON 顶层字段为 `overview` / `current_focus_signals`+`asset_feedback_summary` / `activities`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 2`, `Task 3`
- `Task 5` depends on `Task 2`, `Task 4`
- `Task 6` depends on `Task 4`, `Task 5`
- `Task 7` depends on `Task 1`
- `Task 8` depends on `Task 6`, `Task 7`
- `Task 9` depends on `Task 8`（复核后修复）
