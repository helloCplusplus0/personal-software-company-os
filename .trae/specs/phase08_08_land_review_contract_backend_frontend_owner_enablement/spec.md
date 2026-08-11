# phase08-08 落实 review 相关合同、后端承接与前端 owner 收敛 Spec

## Why

`phase08-05 ~ 07` 已经把 review 的页面流、前端 owner、后端合同与数据承接位冻结成单值设计，但仓库里目前还没有真正可运行的 `review` 合同、后端模块、前端切片、独立路由与 Dashboard launcher。若这一步不把正式承接位落到代码树中，后续实现很容易退回到“页面临时拼 query / mutation / route state”的过渡写法，直接违背 `phase08` 的 owner 收敛目标。

## What Changes

- 落实 `psco.review.v1` proto 合同并接入现有 `buf` 生成主线
- 落实 `backend/internal/review/` 后端模块、Connect handler、最小 `review_records` 数据承接与 `/api` 挂载
- 落实 `frontend/src/features/review/` 前端切片，包括 read owner、application owner、slice-local client、页面壳层与最小页面
- 落实 `/reviews/daily` 与 `/reviews/weekly` 两条独立 route，并正式承接 Dashboard 来源参数
- 落实 `DashboardPrimaryActionPanel` 从旧单 CTA 命中器收敛为 dual review launcher
- 冻结本轮实现的最小验证：`buf`、`go build`、`frontend build`、review API smoke、review route smoke

## Impact

- Affected specs:
  - `phase08_02_freeze_dashboard_review_entry_page_route_handoff`
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `phase08_04_freeze_review_contract_read_record_boundary`
  - `phase08_05_design_dashboard_review_entry_page_interaction_flow`
  - `phase08_06_design_review_frontend_read_write_owner_state_flow`
  - `phase08_07_design_review_backend_contract_service_data_handoff`
  - `phase07_08_land_buf_generation_connect_contract_mainline`
  - `phase07_09_cut_go_backend_transport_mainline`
  - `phase07_10_cut_frontend_client_slice_mainline`
- Affected code:
  - `proto/psco/review/v1/review.proto`
  - `proto/buf.gen.yaml`
  - `proto/Makefile`
  - `backend/internal/review/`
  - `backend/internal/platform/router.go`
  - `backend/internal/connecterrors/connect_errors.go`
  - `database/` 下新增 review record migration
  - `frontend/src/gen/proto/psco/review/v1/*`
  - `frontend/src/features/review/`
  - `frontend/src/routes/reviews/daily.tsx`
  - `frontend/src/routes/reviews/weekly.tsx`
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
  - `frontend/src/features/dashboard/components/dashboard-home-page-shell.tsx`

## ADDED Requirements

### Requirement: review 合同与生成产物必须沿用现有单一 buf 主线落地

系统 SHALL 在现有 `proto/` workspace 内新增 `proto/psco/review/v1/review.proto`，并通过既有 `buf build / gen / lint` 主线同时生成后端 Go proto、Go Connect 与前端 TypeScript 产物。

#### Scenario: review proto 的物理落点与生成产物

- **WHEN** 实现 `phase08-08`
- **THEN** 必须新增 `proto/psco/review/v1/review.proto`
- **AND** 生成产物必须继续只落到：
  - `backend/internal/gen/proto/psco/review/v1/*`
  - `backend/internal/gen/connect/psco/review/v1/*`
  - `frontend/src/gen/proto/psco/review/v1/*`
- **AND** 不得为 review 单独新增第二份 `buf.yaml`、第二份 `buf.gen.yaml`、第二套 TS 生成脚本或第二套前端 RPC 客户端生成链

#### Scenario: review proto 只承接 phase08-07 已冻结的最小 RPC

- **WHEN** 编写 `review.proto`
- **THEN** `ReviewService` 只允许实现：
  - `GetDailyReviewContext`
  - `GetWeeklyReviewContext`
  - `SubmitReviewResult`
- **AND** 当前阶段不得新增 `CreateDecisionFromReview`、`BindModuleFromReview`、`BindRepositoryFromReview`、`MapModuleFromReview` 一类复制 canonical 写路径的 review-local RPC
- **AND** review context message 必须复用 `Dashboard / Decision / Reuse Summary` 既有 canonical message，而不是重新抄一套等价字段

### Requirement: review 后端正式承接位必须落到单一 `backend/internal/review/`

系统 SHALL 将 review 后端正式实现收敛为单一 `backend/internal/review/` 模块，并沿用现有 phase07 Connect transport 主线，不得在 platform 或现有业务模块里散落实现。

#### Scenario: review 后端最小模块结构

- **WHEN** 实现 review 后端
- **THEN** 至少必须新增以下落点：
  - `backend/internal/review/service/query_service.go`
  - `backend/internal/review/service/command_service.go`
  - `backend/internal/review/connect/server.go`
  - `backend/internal/review/repository/review_record_store.go`
- **AND** 若需要 domain types / errors，只允许继续收敛在 `backend/internal/review/` 模块内部
- **AND** 不得把 review query / command 直接塞进 `dashboard/service`、`decisioncenter/service` 或 `platform/router.go`

#### Scenario: review QueryService 只组合既有 canonical service

- **WHEN** `GetDailyReviewContext` 与 `GetWeeklyReviewContext` 落地
- **THEN** `review.QueryService` 只允许消费：
  - `dashboard.QueryService`
  - `decisioncenter.QueryService`
  - `reusesummary.QueryService`
- **AND** 不得绕过这些既有 service 直接新写第二套 SQL、candidate read 或并列 repository
- **AND** Daily Review 的 `pending decisions` 继续固定为 `DecisionListItem(status = proposed)` 的 top N 摘要
- **AND** Weekly Review 的 `reuse snapshot` 继续固定为 `ReuseSummary(scope = dashboard)` 的 canonical 结果

#### Scenario: review CommandService 只承接最小 review result sink

- **WHEN** 实现 `SubmitReviewResult`
- **THEN** `review.CommandService` 只允许承接 `next-step result` 和可选 review 过程留痕
- **AND** `decision handoff / entity handoff` 路径允许完全不调用 `SubmitReviewResult`
- **AND** review 后端不得变成跨 `Decision / Product / Module / Repository` 的总调度写入口

### Requirement: review 记录的数据承接必须以单表轻量过程记录落地

系统 SHALL 为 `next-step result` 落实单一轻量 `review_records` 数据承接位，并保持与 `phase08-04 / 07` 一致的最小字段边界。

#### Scenario: review record 的数据库落点

- **WHEN** 实现 `review_records`
- **THEN** 必须新增单一 migration 与单一 store
- **AND** 最小字段必须覆盖：
  - `id`
  - `review_kind`
  - `result_kind`
  - `started_at`
  - `completed_at`
  - 可选 `decision_id`
  - 可选 `target_type`
  - 可选 `target_id`
  - `summary_text`
  - `created_at`
- **AND** 不得把完整 review context、完整 decision detail 或实体快照冗余写入该表

### Requirement: review Connect transport 必须正式挂到现有 `/api` 业务树

系统 SHALL 将 review 的对外传输正式接入现有 `/api` Connect 主线，保持与其他业务模块一致的 handler 构造与路由挂载方式。

#### Scenario: review Connect handler 的挂载方式

- **WHEN** 落实 review Connect transport
- **THEN** `backend/internal/platform/router.go` 必须新增 `mountReviewConnect(...)` 或等价单值装配位
- **AND** handler 必须继续通过 generated `(path, handler)` 挂到单一 `/api` 业务树
- **AND** review 新增错误必须继续统一收敛到 `connecterrors.MapToConnectError`
- **AND** 不得新增 `/review-api`、`/rpc`、`/internal-review` 等第二套路由根

### Requirement: review 前端正式承接位必须以单一 `frontend/src/features/review/` 切片落地

系统 SHALL 将 review 前端承接位真正实现为单一 `review` 切片，而不是让 route、page、Dashboard 组件继续直接拼底层 query / mutation。

#### Scenario: review 前端最小切片结构

- **WHEN** 实现 review 前端
- **THEN** 至少必须新增以下落点：
  - `frontend/src/features/review/data/connect-client.ts`
  - `frontend/src/features/review/data/review-query-options.ts`
  - `frontend/src/features/review/data/use-daily-review-read.ts`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/review/application/review-action-types.ts`
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/features/review/components/review-page-shell.tsx`
  - `frontend/src/features/review/components/review-action-footer.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
- **AND** 当前阶段不得新增 review 专用 `zustand` store、第二套共享 transport、或页面内部自管的 mutation 编排器

#### Scenario: review 前端 read owner 的正式职责

- **WHEN** `useDailyReviewRead()` 与 `useWeeklyReviewRead()` 实现
- **THEN** 它们必须优先消费 review service 生成客户端，而不是在页面层直接并排消费 `dashboard / decision-center / reuse-summary` 的底层 query hook
- **AND** 它们仍需对外暴露 `pageStatus` 与 `sectionStatus` 的稳定视图模型
- **AND** 若 review service 当前阶段返回的 message 仍直接复用 canonical message，前端只允许在 read owner 内完成最小解包与只读派生
- **AND** 页面与复用 section 不得再直接 import 底层 canonical query hook

#### Scenario: review action owner 的正式职责

- **WHEN** `useReviewAction()` 实现
- **THEN** 它必须成为 review 完成区唯一正式写路径 owner
- **AND** 纯 `decision / entity handoff` 路径只返回 success envelope 并导航到既有 canonical 页面
- **AND** `next-step result` 路径必须正式调用 review service 的 `SubmitReviewResult`
- **AND** 所有成功 envelope 必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`
- **AND** query invalidation 必须统一收敛在 `onSuccess` 内通过 `queryClient.invalidateQueries()` 触发

### Requirement: review route 必须以两条独立文件路由落地

系统 SHALL 将 daily / weekly review 正式落为两条独立文件路由，并正式承接 Dashboard 来源参数，而不是继续停留在“将来再加 route”的状态。

#### Scenario: review route 的最小落点

- **WHEN** 实现 review route
- **THEN** 必须新增：
  - `frontend/src/routes/reviews/daily.tsx`
  - `frontend/src/routes/reviews/weekly.tsx`
- **AND** 两条 route 都必须定义 `validateSearch`
- **AND** `validateSearch` 必须继续复用 `dashboardSourceSearchSchema` 或与之等价的单值 schema
- **AND** review route 的正式路径继续固定为 `/reviews/daily` 与 `/reviews/weekly` 或与之等价的两条显式独立 route

#### Scenario: review route 页面宿主

- **WHEN** review route 进入页面实现
- **THEN** `DailyReviewRoute` 只承接 `DailyReviewPage`
- **AND** `WeeklyReviewRoute` 只承接 `WeeklyReviewPage`
- **AND** route 文件不得直接拼 query、mutation、toast、section-level retry

### Requirement: Dashboard 标题行动区必须在本阶段完成 dual review launcher 收敛

系统 SHALL 在本阶段把 `DashboardPrimaryActionPanel` 从旧的单 CTA 命中器正式收敛为 dual review launcher，使 review route 具备真实入口，而不是继续依赖页面临时按钮。

#### Scenario: DashboardPrimaryActionPanel 的职责切换

- **WHEN** 落实 Dashboard review 入口
- **THEN** `DashboardPrimaryActionPanel` 必须稳定渲染 `Daily Review` 与 `Weekly Review` 两个显式入口
- **AND** 它只负责组装 `/reviews/daily` 或 `/reviews/weekly` 的 route search，并继续透传 `buildDashboardSourceParams('empty-state')`
- **AND** 不得继续保留 `computePrimaryCta()` 驱动的旧单主 CTA 作为 formal review 入口 owner
- **AND** 若仍需保留旧创建导向动作，只能作为降级次级入口，不得与 dual review launcher 并列为同级正式入口

#### Scenario: Dashboard 首屏布局责任不下沉到 review owner

- **WHEN** 为 dual review launcher 调整 Dashboard 标题区
- **THEN** 响应式换行与堆叠责任必须继续落在 `DashboardHomePageShell` 或等价标题壳层
- **AND** `DashboardPrimaryActionPanel` 不得额外持有 review read owner、review action owner、或页面级状态机

### Requirement: 本阶段最小实现必须提供 review 页面与动作流的正式使能位

系统 SHALL 在本阶段交付“可进入 route、可读取 review context、可触发最小完成动作、可回流 canonical 页面”的正式 enablement，而不是只交付空文件或占位接口。

#### Scenario: Daily / Weekly Review 最小 enablement

- **WHEN** 实现 `DailyReviewPage` 与 `WeeklyReviewPage`
- **THEN** 两个页面都必须通过 `ReviewPageShell` 承接头部、离开入口、页面级状态与底部动作区
- **AND** `DailyReviewPage` 至少必须显示 `current focus / pending decisions / representative signals` 的最小只读组织
- **AND** `WeeklyReviewPage` 至少必须显示 `overview / recent activity / representative signals / reuse snapshot` 的最小只读组织
- **AND** 两个页面都必须接上 `ReviewActionFooter` 或等价完成区，能够正式触发 `useReviewAction()`
- **AND** 不得把“下一步再接 action owner”留成 page-local TODO

### Requirement: 验收口径必须同时覆盖合同、后端、前端与关键路径 smoke

系统 SHALL 将 `phase08-08` 的实现验收冻结为“合同工具链 + 后端构建 + 前端构建 + review smoke”四层。

#### Scenario: 最小验证命令

- **WHEN** 验证 `phase08-08` 是否达成 DoD
- **THEN** 至少必须通过：
  - `(cd proto && make build && make gen && make lint)`
  - `(cd backend && go build ./...)`
  - `(cd frontend && npm run build)`
- **AND** 不得只验证后端构建而跳过前端 owner 收敛的类型构建

#### Scenario: review 关键路径 smoke

- **WHEN** 执行最小 smoke
- **THEN** 至少必须覆盖：
  - `POST /api/psco.review.v1.ReviewService/GetDailyReviewContext`
  - `POST /api/psco.review.v1.ReviewService/GetWeeklyReviewContext`
  - `POST /api/psco.review.v1.ReviewService/SubmitReviewResult` 的 `next-step result` 路径
  - `/dashboard -> /reviews/daily`
  - `/dashboard -> /reviews/weekly`
  - review 页面进入 `Decision / Product / Module / Repository` 既有 canonical path 时继续透传 Dashboard 来源参数

## MODIFIED Requirements

### Requirement: `DashboardPrimaryActionPanel` 的正式身份

自 `phase08-08` 起，`DashboardPrimaryActionPanel` SHALL 被正式解释为 review route launcher caller，而不再是 `phase05` 的单 CTA 命中器。

#### Scenario: 旧 CTA 矩阵退场

- **WHEN** 实现 `phase08-08`
- **THEN** `computePrimaryCta()` 不再作为 `DashboardPrimaryActionPanel` 的正式主线依赖
- **AND** review 入口必须始终稳定可见
- **AND** 不得再用 `overview / feedback` 命中结果决定 review 入口是否显示

### Requirement: review 前端 owner 的消费边界

`phase08-06` 已冻结 review 编排不得散落在 route / 页面 / 展示组件。自 `phase08-08` 起，系统必须进一步要求这些 owner 直接落到 review 切片与 route 文件中，而不是继续以“后续实现再抽离”为理由保留页面级临时编排。

#### Scenario: 页面不再是临时 owner

- **WHEN** 审查 `DailyReviewPage / WeeklyReviewPage`
- **THEN** 页面只能消费 `useDailyReviewRead / useWeeklyReviewRead / useReviewAction`
- **AND** 页面不得再直接 `createClient()`、`createConnectTransport()`、`useMutation()` 或并排拼多组底层 query hook

### Requirement: review read owner 的正式 transport 路径

`phase08-06` 已冻结 `useDailyReviewRead / useWeeklyReviewRead` 必须是前端唯一正式 read owner。自 `phase08-08` 起，系统必须进一步将其实现路径修改为“优先消费 `ReviewService` 生成客户端”，而不是继续在前端页面树中组合多个 canonical hook 作为正式主线。

#### Scenario: review read owner 的实现路径从前端本地组合收敛到 review service

- **WHEN** 实现 `useDailyReviewRead()` 与 `useWeeklyReviewRead()`
- **THEN** 它们必须以 `frontend/src/features/review/data/connect-client.ts` 的 generated client 为正式 transport 入口
- **AND** `dashboard / decision-center / reuse-summary` 的既有 hook 继续服务各自切片，不再承担 review 正式读主线
- **AND** 若仍需复用既有前端 mapper / type helper，只允许在 review slice 内最小复用，不得把页面重新拉回“多 hook 本地组合”模式

## REMOVED Requirements

### Requirement: review enablement 可以先通过页面级临时编排达成

**Reason**: 这会直接破坏 `phase08-06` 已冻结的 owner 单值化目标，并让 `phase08-08` 的实现退化成新的技术债入口。
**Migration**: 本阶段必须直接落 `review` 切片、review route、review Connect contract、review 后端 owner 与 `review_records` 正式承接位；页面只消费这些正式 owner，不再承担临时 orchestration。
