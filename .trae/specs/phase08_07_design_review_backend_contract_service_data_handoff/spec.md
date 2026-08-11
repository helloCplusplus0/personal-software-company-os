# phase08-07 产出后端服务、合同与最小数据承接设计 Spec

## Why

`phase08-04` 已冻结 review 合同只能承接 review context、最小动作命令与可选 review record，`phase08-05 / 06` 又已经把页面流、前端 read/application owner 和成功回流语义收紧为单值。如果后端这一层还不把 `.proto`、Connect service、后端 query / command owner、以及 review result 的最小数据落点冻结清楚，后续实现很容易重新把 Dashboard / Decision / Reuse Summary 的事实复制一遍，或者在 review 模块里长出第二套 command 主线。

## What Changes

- 冻结 review 后端正式合同包、RPC 矩阵与 Connect 服务承接位
- 冻结 review read model 的后端 owner 及其与 `Dashboard / Decision / Reuse Summary` 既有服务的协作边界
- 冻结 review 动作命令与结果回流的后端 owner，明确哪些动作继续直接走既有 canonical 服务
- 冻结可选 `review record` / `next-step result` 的最小数据承接设计
- 冻结 review 模块在 `platform/router.go`、`connecterrors`、`buf` 工具链中的接线方式
- 冻结 review 关键路径的 `buf / go build / API smoke` 验收清单

## Impact

- Affected specs:
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `phase08_04_freeze_review_contract_read_record_boundary`
  - `phase08_05_design_dashboard_review_entry_page_interaction_flow`
  - `phase08_06_design_review_frontend_read_write_owner_state_flow`
  - `phase07_04_design_go_connect_handler_service_chi_mount`
  - `phase07_08_land_buf_generation_connect_contract_mainline`
  - `phase07_09_cut_go_backend_transport_mainline`
- Affected code:
  - 后续新增 `proto/psco/review/v1/review.proto`
  - `proto/buf.gen.yaml`
  - `proto/Makefile`
  - 后续新增 `backend/internal/review/`
  - 后续新增 `backend/internal/review/service/query_service.go`
  - 后续新增 `backend/internal/review/service/command_service.go`
  - 后续新增 `backend/internal/review/connect/server.go`
  - 后续新增 `backend/internal/review/repository/review_record_store.go`
  - `backend/internal/platform/router.go`
  - `backend/internal/platform/server.go`
  - `backend/internal/connecterrors/connect_errors.go`
  - `backend/internal/dashboard/service/query_service.go`
  - `backend/internal/decisioncenter/service/query_service.go`
  - `backend/internal/decisioncenter/service/command_service.go`
  - `backend/internal/reusesummary/service/query_service.go`
  - `database/` 下后续新增 review record migration

## ADDED Requirements

### Requirement: review 后端正式合同必须落到单一 `psco.review.v1` 包

系统 SHALL 将 review 后端正式合同冻结为新的 `proto/psco/review/v1/review.proto`，并通过单一 `ReviewService` 承接 review context 读取与最小 review-local 结果提交。

#### Scenario: review proto 的最小 RPC 矩阵

- **WHEN** 后续实现 review 后端合同
- **THEN** `ReviewService` 至少必须包含以下 RPC：
  - `GetDailyReviewContext`
  - `GetWeeklyReviewContext`
  - `SubmitReviewResult`
- **AND** 当前阶段不得新增 `CreateDecisionFromReview`、`BindModuleFromReview`、`MapRepositoryFromReview` 一类复制既有实体写入的 review-local RPC
- **AND** review 相关对外合同继续只通过 ConnectRPC 暴露，不得新增 hand-written JSON canonical API

#### Scenario: review proto 复用既有 canonical message，而不是重复抄字段

- **WHEN** `review.proto` 需要表达 review context
- **THEN** 必须直接 import 并复用既有 canonical message：
  - `psco.dashboard.v1.DashboardOverview`
  - `psco.dashboard.v1.FeedbackSignal`
  - `psco.dashboard.v1.RecentActivityItem`
  - `psco.decision_center.v1.DecisionListItem`
  - `psco.reuse_summary.v1.ModuleReuseSummary`
  - `psco.reuse_summary.v1.CapabilitySummary`
- **AND** 不得在 `review.proto` 中重新定义等价的 `ReviewOverview`、`ReviewFeedbackSignal`、`ReviewDecisionSummary`、`ReviewReuseSummary`
- **AND** review 合同只允许新增承接 review 流程语义的包装 message 与 enum

### Requirement: review read model 必须由单一后端 QueryService 编排既有 canonical 服务

系统 SHALL 将 review read model 的后端 owner 冻结为 `backend/internal/review/service/query_service.go`，并要求它只编排既有 query service，不直接重写 candidate reader 或 SQL。

#### Scenario: Daily Review 后端读取承接位

- **WHEN** 后续实现 `GetDailyReviewContext`
- **THEN** `review.QueryService` 必须只消费：
  - `dashboard.QueryService.ReadFeedbackSignal`
  - `decisioncenter.QueryService.ListDecisions(status = proposed)` 作为当前阶段 `pending / backlog decisions` 的最小后端事实来源
- **AND** Daily Review 返回的 `pending_decisions` 必须直接建立在 `DecisionListItem` 的 top N 摘要之上
- **AND** 不得在 review 模块里重写第二套 pending decision 聚合 SQL

#### Scenario: Weekly Review 后端读取承接位

- **WHEN** 后续实现 `GetWeeklyReviewContext`
- **THEN** `review.QueryService` 必须只消费：
  - `dashboard.QueryService.ReadOverview`
  - `dashboard.QueryService.ReadRecentActivity`
  - `dashboard.QueryService.ReadFeedbackSignal`
  - `reusesummary.QueryService.ReadReuseSummary(scope = dashboard)`
- **AND** Weekly Review 的 `reuse snapshot` 必须继续来自 `ReuseSummaryService` 的 canonical 结果
- **AND** 不得在 review 模块里复制第二套 `module_reuse_summary / capability_summary` 聚合逻辑

#### Scenario: review QueryService 不成为新的事实主线

- **WHEN** `review.QueryService` 组装 Daily / Weekly context
- **THEN** 它只能做 request 校验、既有 service 调用、只读组合与轻量派生
- **AND** 不得把组合结果反向持久化成新的总表、物化视图或并列缓存真相源
- **AND** 不得绕过 `dashboard / decisioncenter / reusesummary` 既有 query service 直接访问下游 repository

### Requirement: review 动作命令的后端 owner 必须区分“直接复用 canonical command”与“review-local result sink”

系统 SHALL 将 review 动作命令的后端边界冻结为两层：实体写入继续走既有 canonical command owner；review-local 后端命令只承接 `SubmitReviewResult` 这一类流程结果提交。

#### Scenario: Decision / Product / Module / Repository 写入继续走既有 canonical command

- **WHEN** review 结果需要创建 `Decision`、关联 `Decision -> Module`、绑定 `Module -> Product`、绑定 `Repository -> Product` 或映射 `Module -> Repository`
- **THEN** 后端正式 owner 继续分别是：
  - `decisioncenter.CommandService`
  - `moduleregistry.QueryService` 与 `ModuleRegistryService` 继续承接 `Module` canonical 读取与 handoff 目标
  - `productregistry.CommandService`
  - `repositorybinding.CommandService`
- **AND** review 模块不得为这些写入再长出第二套 `review-local command service`
- **AND** `SubmitReviewResult` 不得被设计成代替这些 canonical RPC 的万能写入口
- **AND** 当 review 结果只是进入既有 `Module Detail / Module List` canonical 页面时，后端不需要新增 review-local module command，必须继续复用 `Module Registry` 既有 Connect canonical path

#### Scenario: review-local 后端命令只承接最小流程结果

- **WHEN** 后续实现 `SubmitReviewResult`
- **THEN** 它只允许承接以下流程语义字段：
  - `review_kind`
  - `result_kind`
  - 可选的 `decision_id`
  - 可选的 `target_type / target_id`
  - 最小摘要文本
  - `started_at / completed_at` 或等价时间字段
- **AND** 它不得承接完整 `DecisionDraftInput`、完整 `ProductBindingInput`、完整 `RepositoryMappingInput`
- **AND** 对 `decision handoff / entity handoff` 路径，`SubmitReviewResult` 允许成为可选记录提交，而不是强制前置步骤

### Requirement: review 结果回流的后端响应语义必须保持最小化

系统 SHALL 将 review 结果回流的后端响应语义冻结为“返回稳定结果 envelope 或记录标识”，而不是在 review 模块中内嵌 canonical reread 结果。

#### Scenario: `SubmitReviewResult` 的最小响应边界

- **WHEN** `SubmitReviewResult` 成功
- **THEN** 响应只允许返回：
  - 可选的 `review_record_id`
  - `result_kind`
  - 可选的 `decision_id`
  - 可选的 `target_type / target_id`
- **AND** 不得在响应中内嵌完整 `DecisionDetail`、`ProductDetail`、`ModuleDetail`、`RepositoryDetail`
- **AND** 前端后续进入 canonical 页面后的 reread 继续由既有 service 承接

#### Scenario: 无 record 路径的后端边界

- **WHEN** 当前 review 结果是 `decision handoff` 或 `entity handoff`，且产品选择不记录 review record
- **THEN** 后端允许完全不调用 `SubmitReviewResult`
- **AND** 该路径不得被解释为 review 模块缺失
- **AND** 这不改变 `next-step result` 必须拥有正式落点的要求

### Requirement: `next-step result` 的最小数据承接必须冻结为单表轻量过程记录

系统 SHALL 将 `next-step result` 的后端正式落点冻结为单一轻量 `review_records` 数据承接位，而不是把结果写进 Dashboard / Decision / Product / Module / Repository 的影子字段中。

#### Scenario: review record 的最小数据模型

- **WHEN** 后续实现需要为 `next-step result` 提供正式持久化
- **THEN** 数据层只允许新增单一 `review_records` 或等价单表
- **AND** 该表最小字段至少包括：
  - `id`
  - `review_kind`
  - `result_kind`
  - `started_at`
  - `completed_at`
  - 可选的 `decision_id`
  - 可选的 `target_type`
  - 可选的 `target_id`
  - `summary_text`
  - `created_at`
- **AND** 不得复制完整实体快照、完整 review context、完整 Decision / Product / Module / Repository 结构

#### Scenario: review record 的后端承接位

- **WHEN** 后续实现 review 结果持久化
- **THEN** 必须新增：
  - `backend/internal/review/repository/review_record_store.go`
  - `backend/internal/review/service/command_service.go`
- **AND** `review.CommandService` 只允许消费 `ReviewRecordStore`
- **AND** 不得借机把 `review.CommandService` 做成跨实体写入总调度器

### Requirement: review Connect service implementation 与 platform 挂载方式必须沿用 phase07 主线

系统 SHALL 将 review 模块的后端接线方式冻结为“generated Connect handler + review service implementation + platform mount”，与现有 phase07 canonical pattern 保持一致。

#### Scenario: review Connect handler 的物理落点与挂载方式

- **WHEN** 后续实现 review Connect transport
- **THEN** 必须新增 `backend/internal/review/connect/server.go`
- **AND** `platform/router.go` 必须新增 `mountReviewConnect(...)`
- **AND** review handler 必须通过 generated `(path, handler)` 挂到现有单一 `/api` 业务树下
- **AND** 不得新增第二个 `/rpc`、第二个 `/review-api` 或并列路由根

#### Scenario: review transport 继续使用单值错误映射

- **WHEN** review Connect handler 返回 domain error
- **THEN** 必须继续统一走 `connecterrors.MapToConnectError`
- **AND** review 新增错误只允许扩展现有单值映射入口
- **AND** 不得在 review handler 内临时拼装第二套 JSON / Connect 错误语义

### Requirement: `phase05 / phase06` 既有服务必须被明确纳入 review 后端消费边界

系统 SHALL 在后端设计层明确：review 不是新的业务事实模块，而是对既有 `Dashboard / Decision / Reuse Summary` 服务的受控消费层。

#### Scenario: review 与既有服务的协作边界

- **WHEN** 审核 review 后端模块设计
- **THEN** 必须看到以下单值边界：
  - `DashboardService` 继续只承接 `overview / feedback / recent activity`
  - `ReuseSummaryService` 继续只承接 `reuse snapshot`
  - `DecisionCenterService` 继续只承接 decision canonical 读写
  - `ModuleRegistryService` 继续只承接 module canonical 读与 module detail handoff 目标
  - `ProductRegistryService / RepositoryBindingService` 继续只承接实体 canonical 写入
- **AND** review 只消费这些服务，不反向吞并它们的职责

### Requirement: review 关键路径的工具链与 API 验收口径必须冻结

系统 SHALL 将 `phase08-07` 的后端设计验收口径冻结为“合同工具链 + Go 构建 + review Connect smoke”三层，而不是实现时临场补测试口径。

#### Scenario: 合同与构建验收口径

- **WHEN** 团队验证 review 后端设计是否可直接进入实现
- **THEN** 最小工具链验收命令必须冻结为：
  - `(cd proto && make build && make gen && make lint)`
  - `(cd backend && go build ./...)`
- **AND** 当前阶段不得把根目录直接执行 `buf` 或 `go build ./...` 写成新的正式口径

#### Scenario: review Connect / `/api` smoke 清单

- **WHEN** 团队执行 review 后端关键路径 smoke
- **THEN** 至少必须覆盖：
  - `POST /api/psco.review.v1.ReviewService/GetDailyReviewContext`
  - `POST /api/psco.review.v1.ReviewService/GetWeeklyReviewContext`
  - `POST /api/psco.review.v1.ReviewService/SubmitReviewResult` 的 `next-step result` 成功路径
  - 既有 `DecisionCenterService.CreateDecision` 仍可作为 review 后续 canonical path
- **AND** 还必须覆盖至少一种实体回流或 canonical action handoff 路径：
  - `ProductRegistryService.BindModuleToProduct`
  - `RepositoryBindingService.BindRepositoryToProduct`
  - `RepositoryBindingService.MapModuleToRepository`
  - `ModuleRegistryService.GetModuleDetail`
- **AND** 若当前阶段未实现 `SubmitReviewResult` 的 decision/entity handoff 记录路径，smoke 中必须显式说明该路径为“直接复用 canonical service，无 review-local RPC”

## MODIFIED Requirements

### Requirement: review 合同与后端 owner 的职责解释

自 `phase08-07` 起，review 合同与后端 owner SHALL 被解释为“消费既有 canonical service 的流程编排层 + 可选结果承接层”，而不是并列的业务事实模块。

#### Scenario: review backend 不是第二套业务主线

- **WHEN** 后续实现 review 后端模块
- **THEN** `review.QueryService` 只能组合读取
- **AND** `review.CommandService` 只能承接最小 review result sink
- **AND** 实体 canonical 读写仍必须回到既有模块服务

### Requirement: review record 的存在方式

`phase08-04` 已冻结 review 记录是可选轻量过程记录。自 `phase08-07` 起，系统必须进一步把这一 requirement 修改为“若存在 review record，其后端承接只能是单表 + 单一 repository / command owner”。

#### Scenario: review record 不膨胀为核心实体

- **WHEN** 后续实现 `review_records`
- **THEN** 它只能服务于 `next-step result` 或可选 review 过程留痕
- **AND** 不得扩展为一级导航、独立详情 CRUD 或新的聚合事实源

## REMOVED Requirements

### Requirement: review 后端通过复制既有 Dashboard / Decision / Reuse Summary 字段或写入接口来实现闭环

**Reason**: 这会直接破坏 `.proto` 单一合同源、phase07 的 Connect 主线，以及 `phase08-04 / 08-06` 已冻结的 canonical service / owner 边界。
**Migration**: review 后端统一新增 `psco.review.v1` 合同与 `backend/internal/review/` 模块；读取只组合既有 query service，实体写入继续直连既有 canonical service，`next-step result` 只落到轻量 `review_records` 承接位。
