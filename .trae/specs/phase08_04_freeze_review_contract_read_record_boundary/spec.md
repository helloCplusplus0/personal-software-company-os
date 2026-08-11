# phase08-04 冻结本阶段合同、读模型与记录模型的最小边界 Spec

## Why

`phase08-01 ~ 03` 已经分别冻结了 `Operating Review Loop` 的范围边界、`Dashboard -> Review` 的入口与路由承接位，以及 `Feedback -> Decision -> Update` 的动作边界与 owner。如果此时不继续把 review 相关的最小正式合同、读模型、写模型和可选记录模型边界冻结成单值结论，后续实现很容易把 `phase05` Dashboard、`phase06` Reuse Summary、`Decision Center` 与 review 自己的上下文拼成第二套事实主线，既重写既有 `.proto` 合同职责，也让 review 记录意外膨胀成新的长期核心实体。

## What Changes

- 冻结 review 入口、review 上下文、review 动作与结果回流的最小正式合同边界
- 冻结本阶段必需的 review read model / write model 边界
- 冻结 `phase05` Feedback 与 `phase06` Reuse Awareness 的正式消费范围
- 冻结 review context 与既有 `Dashboard / Reuse Summary / Decision` canonical 事实源之间的关系
- 冻结是否需要 review 记录，以及其轻量化边界
- 冻结当前真实 `caller / route / query owner / application owner` inventory 必须如何进入后续 `/spec`

## Impact

- Affected specs:
  - `phase08_01_freeze_operating_review_loop_scope_success_non_goals`
  - `phase08_02_freeze_dashboard_review_entry_page_route_handoff`
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `phase05_10_dashboard_feedback_formal_spec`
  - `phase06_04_define_reuse_summary_read_model_page_attachment`
  - `phase06_09_design_reuse_summary_query_dashboard_detail_integration`
- Affected code:
  - `proto/psco/dashboard/v1/dashboard.proto`
  - `proto/psco/reuse_summary/v1/reuse_summary.proto`
  - 后续新增的 `proto/psco/review/v1/` 或等价 review 合同包
  - `frontend/src/features/dashboard/data/*`
  - `frontend/src/features/reuse-summary/data/use-reuse-summary-read.ts`
  - `frontend/src/features/decision-center/data/*`
  - 后续新增的 review read owner / review action owner
  - `backend/internal/dashboard/service/query_service.go`
  - `backend/internal/reusesummary/service/query_service.go`
  - 后续新增的 review query / record 承接位

## ADDED Requirements

### Requirement: review 相关正式合同必须保持与既有 canonical contract 解耦

系统 SHALL 将 review 相关正式合同冻结为“新增 review 合同承接 review 上下文与可选过程记录；既有 `Dashboard / Reuse Summary / Decision / Product / Module / Repository` canonical contract 继续承接各自事实”的单值边界。

#### Scenario: `.proto` 单一长期合同源继续成立

- **WHEN** 后续 `/spec` 或实现新增 review 相关接口
- **THEN** `.proto` 必须继续作为唯一长期合同源
- **AND** review 读取与动作承接必须继续走 ConnectRPC
- **AND** 不得为 review 新长出 hand-written JSON canonical contract

#### Scenario: review 合同与既有 contract 的关系

- **WHEN** 接手者设计 review 合同边界
- **THEN** review 合同只允许承接：
  - daily / weekly review context 读取
  - review 动作承接所需的最小命令输入
  - 可选的 review 过程记录提交
- **AND** 不得把 `DashboardOverview / FeedbackSignal / RecentActivity / ReuseSummary / DecisionDetail` 改写为 review-local 并列合同
- **AND** 不得通过扩写 `dashboard.proto` 或 `reuse_summary.proto` 把 review session / review result / review record 吞进既有服务职责

### Requirement: review read model 必须是对既有事实源的轻量组合层

系统 SHALL 将 review read model 冻结为“组合既有 canonical facts 的轻量读层”，而不是新的长期事实源。

#### Scenario: daily review 必需消费范围

- **WHEN** 后续 `/spec` 冻结 `daily review` 的最小 read model
- **THEN** 必须正式消费以下既有事实范围：
  - `phase05` `FeedbackSignalRead.current_focus_signals`
  - `phase05` `FeedbackSignalRead.asset_feedback_summary.representative_signals`
  - 既有 `Decision` canonical fact 中可解释为 pending / backlog 的最小范围
- **AND** 当前阶段不得把 daily review 扩写成独立任务列表事实源
- **AND** daily review read model 只承接“当前要处理什么”的最小上下文，不复刻 Dashboard 全页概览

#### Scenario: weekly review 必需消费范围

- **WHEN** 后续 `/spec` 冻结 `weekly review` 的最小 read model
- **THEN** 必须正式消费以下既有事实范围：
  - `phase05` `DashboardOverview`
  - `phase05` `RecentActivityRead.activities`
  - `phase05` `FeedbackSignalRead.asset_feedback_summary.representative_signals`
  - `phase06` `ReuseSummaryRead(scope = dashboard)` 返回的 `module_reuse_summary / capability_summary`
- **AND** weekly review read model 不得跳过 `reuse snapshot`
- **AND** 不得把 weekly review 写成新的统计中心或独立 intelligence feed

#### Scenario: review context 不是第二套事实主线

- **WHEN** review read model 需要展示标题、解释文案、优先级、排序结果或分组结果
- **THEN** 这些字段只允许作为 review 上下文派生层存在
- **AND** 不得成为新的 canonical count / status / score 事实源
- **AND** 不得把 review context 持久化后再反向覆盖 `Dashboard / Reuse Summary / Decision` 的原始事实语义

### Requirement: `phase05 / phase06` 既有数据的正式消费边界必须单值化

系统 SHALL 将 `phase05 Feedback` 与 `phase06 Reuse Awareness` 的正式消费范围冻结为 `phase08-04` 的必答项，而不是留给后续实现自由发挥。

#### Scenario: `phase05 Feedback` 的正式消费边界

- **WHEN** 接手者定义 review context 的最小输入
- **THEN** `phase05` 当前正式消费边界至少必须包括：
  - `current_focus_signals`
  - `asset_feedback_summary.representative_signals`
  - `DashboardOverview`
  - `RecentActivity`
- **AND** 不得修改这些读模型在 `phase05` 中已冻结的字段语义
- **AND** 不得把 review 需要的额外解释性字段直接回填到 `dashboard.proto` 作为并列长期语义

#### Scenario: `phase06 Reuse Awareness` 的正式消费边界

- **WHEN** 接手者定义 weekly review 的复用感知输入
- **THEN** 必须正式消费 `ReuseSummaryRead(scope = dashboard)` 返回的：
  - `module_reuse_summary`
  - `capability_summary`
- **AND** 当前阶段不得把 `ReuseSummaryRead` 合并回 `FeedbackSignalRead`
- **AND** 不得为 review 另建第二套 `reuse snapshot` 读模型去替代 `phase06` 已交付的 canonical query owner

### Requirement: review read owner 与 write owner 的边界必须继续分离

系统 SHALL 延续当前项目 `query` 纯只读、`application` 唯一写路径的边界，将 review read owner 与 review action owner 冻结为两类不同承接位。

#### Scenario: review read owner 的职责

- **WHEN** 后续实现 review 相关前端读取
- **THEN** 必须存在单一 review read owner 或等价单一 read layer
- **AND** 该 read owner 只承接：
  - review context 读取
  - 对既有 `Dashboard / Reuse Summary / Decision` 响应的解包与组合
  - review 页面所需的只读状态派生
- **AND** 不得在该 owner 中混入 `create / update / bind / link / submit` 一类写动作

#### Scenario: review action owner 的职责

- **WHEN** 后续实现 review 写路径
- **THEN** review action owner 继续只承接：
  - decision action
  - 结果提交
  - 实体回流
  - 失效刷新
  - 错误归一化
- **AND** 该边界必须与 `phase08-03` 已冻结结论保持一致
- **AND** 不得让 review read owner 退化成 page-local mutation owner

### Requirement: review-local write model 不得复制既有实体写模型

系统 SHALL 将 review-local write model 冻结为“最多承接 review 过程输入与可选记录提交”，而不是复制 `Decision / Product / Module / Repository` 的结构化写模型。

#### Scenario: review 动作与实体写回的关系

- **WHEN** review 动作需要回流既有实体
- **THEN** 必须继续复用 `phase08-03` 已冻结的 canonical owner / canonical action handoff
- **AND** review-local write model 不得重新定义 `DecisionDraftInput`、`ProductBindingInput`、`RepositoryMappingInput` 之类与既有实体等价的并列请求结构
- **AND** 若后续需要新增 review command 合同，其作用只能是承接 review 流程语义，而不是成为新的实体写入真相源

### Requirement: review 记录是可选的轻量过程记录，而不是新的长期核心实体

系统 SHALL 将 review 记录冻结为“可选、轻量、过程性”的模型；`phase08` 不以引入 review 记录为前置阻断项。

#### Scenario: 当前阶段不强制要求 review record

- **WHEN** 后续实现选择直接通过 route handoff + canonical action owner 闭合 review loop
- **THEN** 当前阶段允许不引入独立 review 记录
- **AND** 该无 record 路径当前只覆盖：
  - `decision handoff`
  - `entity handoff`
- **AND** 这不得阻断 `phase08` 最小闭环成立

#### Scenario: `next-step result` 的最小正式落点

- **WHEN** 后续实现选择把 `next-step result` 作为 review 的正式输出
- **THEN** 该结果必须落到轻量 `review record`
- **AND** 不得把 `next-step result` 做成仅存在于页面局部状态、toast、临时勾选或一次性文案中的瞬时行为
- **AND** 当前阶段不得另外发明第二套与 `review record` 并列的 `next-step` 专用 sink
- **AND** 这不改变 `review record` 对 `decision handoff / entity handoff` 路径仍然保持可选

#### Scenario: 若新增 review 记录，其最小身份边界

- **WHEN** 后续实现确实需要 review 记录或 review session 落点
- **THEN** 该记录只能被解释为“经营回路过程记录”
- **AND** 不得升级为新的长期核心实体
- **AND** 不得要求一级导航、列表页、详情页、独立 CRUD 或第二套生命周期状态机

#### Scenario: 若新增 review 记录，其最小字段边界

- **WHEN** 接手者为 review 记录定义最小字段集合
- **THEN** 至少只允许承接以下过程性字段：
  - `review_kind`（daily / weekly）
  - `started_at / completed_at` 或等价时间字段
  - `result_kind`（如 decision handoff / entity handoff / next-step result）
  - 可选的 `decision_id`
  - 可选的 `target_type / target_id`
  - 最小摘要文本
- **AND** 不得复制实体影子快照、完整 Decision 内容、完整 Product / Module / Repository 状态副本

### Requirement: 既有领域实体事实源与 review 新读模型之间的关系必须明确

系统 SHALL 明确 review 新读模型与既有领域实体事实源之间是“读时组合关系”，而不是“主从同步关系”。

#### Scenario: `Dashboard / Reuse Summary / Decision` 与 review context 的关系

- **WHEN** review context 需要显示 overview、feedback、recent activity、reuse snapshot 或 decision backlog
- **THEN** 必须把这些数据解释为来自既有 canonical fact 的读时组合
- **AND** 不得把 review context 存成新的总表或聚合表，再由其他页面反向消费
- **AND** 不得让 review context 反向取代 `DashboardService`、`ReuseSummaryService` 或 `Decision Center` 的事实职责

### Requirement: 当前真实 caller / route / owner inventory 必须直接进入后续 `/spec`

系统 SHALL 将 `shared_baseline` 中已经冻结的真实 `caller / route / query owner / application owner` inventory 视为后续 `/spec` 的强制输入，而不是参考说明。

#### Scenario: caller / route / owner inventory 的最小执行清单

- **WHEN** 后续 `/spec` 进入实现设计
- **THEN** 必须逐项说明以下 inventory 如何被消费：
  - `DashboardRoute (/dashboard)` 与 `DashboardHomePage`
  - `useDashboardOverviewRead`
  - `useFeedbackSignalsRead`
  - `useRecentActivitiesRead`
  - `useReuseSummaryRead`
  - `DashboardPrimaryActionPanel`
  - `CurrentFocusSection`
  - `AssetFeedbackSection`
  - `FeedbackSignalCard`
  - `RecentActivitySection`
  - `RecentActivityItemCard`
  - `DashboardStatBar`
  - `OnboardingCtaButton`
  - `SovereigntyPanel`
  - `Decision` 相关 route / page / read owner / write owner
  - `dashboard-source.ts`
  - `BackToDashboardButton`
- **AND** 后续 `/spec` 必须明确哪些 inventory 被复用、哪些被扩展、哪些明确禁止升级为新的写路径 owner
- **AND** 必须继续明确 `OnboardingCtaButton / SovereigntyPanel` 在当前阶段保持既有入口职责，不升级为正式 review 入口
- **AND** 必须继续明确 `CurrentFocusSection / AssetFeedbackSection / RecentActivitySection` 只承接页面编排与只读展示，不升级为新的写路径 owner
- **AND** 不得在缺少真实 inventory 对照的情况下直接发明抽象的 `ReviewPageOwner / ReviewGateway / ReviewStore`

## MODIFIED Requirements

### Requirement: `Dashboard / Reuse Summary / Decision` 现有合同职责

`phase05` 的 `DashboardService`、`phase06` 的 `ReuseSummaryService` 与既有 `Decision Center` 合同在 `phase08-04` 中 SHALL 继续保持各自 canonical 职责，不因 review loop 而被重写为并列的 review-native 事实总线。

#### Scenario: 既有服务职责保持不变

- **WHEN** review loop 引入新的 context read 或 record write
- **THEN** `DashboardService` 继续只承接 overview / feedback / recent activity 读取
- **AND** `ReuseSummaryService` 继续只承接 reuse snapshot 读取
- **AND** `Decision Center` 继续承接 decision canonical 读写
- **AND** 不得通过“顺手扩字段”把 review 过程语义塞回这些既有服务

## REMOVED Requirements

### Requirement: 为 review 新长出第二套事实主线或并列状态体系

**Reason**: 这会直接破坏 `phase05 / phase06 / phase08-03` 已冻结的 canonical contract、query owner 与 application owner 边界，并让 review loop 从“消费既有交付”退化成“复制既有交付”。  
**Migration**: review 相关新增内容只允许落在新增 review context / optional review record 承接位；既有 overview / feedback / recent activity / reuse snapshot / decision / entity writeback 继续由各自 canonical contract 与 owner 承接。
