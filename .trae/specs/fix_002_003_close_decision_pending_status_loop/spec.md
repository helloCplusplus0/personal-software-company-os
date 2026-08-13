# Decision Pending 与状态推进闭环修复 Spec

## Why
`fix_002` 与 `fix_003` 已经确认，当前 PSCO 的 `Decision` 主线存在一个同源断裂：`Dashboard / Daily Review / Current Focus` 把 `proposed` 决策当作 pending，但 `Decision Detail` 又没有正式状态推进入口，导致用户即使已经“处理过”，系统仍会持续催促。要真正修复这个问题，不能只改单条 reader 条件，也不能只在详情页加本地按钮，而必须同时冻结 canonical pending 语义与 canonical 状态推进写链。

## What Changes
- 冻结 `pending decision` 的唯一 canonical 判定继续锚定在 `Decision.status`，不得改由 `decision_links` 或 `review_records` 代理。
- 为 `Decision Detail` 增加正式状态推进写链：`.proto -> Connect -> CommandService -> DecisionStore -> 前端 mutation owner`。
- 在既有四态 `proposed / active / superseded / archived` 内冻结最小状态推进矩阵，不新增第五态。
- 要求状态推进成功后统一触发 `Decision Detail / Decision List / Dashboard / Review` 的 reread，避免详情页与 pending 展示语义分叉。
- 保持 `SubmitReviewResult` 与 `LinkDecisionToTarget` 的既有职责边界不变。

## Impact
- Affected specs:
  - `phase03_02_decision_template_status_read_model`
  - `phase03_08_decision_center_proto_contract`
  - `phase03_10_decision_center_formal_spec`
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `phase08_07_design_review_backend_contract_service_data_handoff`
  - `phase08_10_land_feedback_decision_update_closed_loop_result_writeback_cleanup`
  - `fix_002_decision_pending_signal_semantics_analysis`
  - `fix_003_decision_detail_status_advance_analysis`
- Affected code:
  - `proto/psco/decision_center/v1/decision_center.proto`
  - `backend/internal/decisioncenter/connect/server.go`
  - `backend/internal/decisioncenter/service/command_service.go`
  - `backend/internal/decisioncenter/repository/decision_store.go`
  - `backend/internal/decisioncenter/types.go`
  - `backend/internal/dashboard/candidate/feedback_readers.go`
  - `backend/internal/review/service/query_service.go`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
  - `frontend/src/features/decision-center/application/*`
  - `frontend/src/features/review/data/use-daily-review-read.ts`
  - `frontend/src/features/review/pages/daily-review-page.tsx`

## ADDED Requirements
### Requirement: Pending Decision 必须继续锚定 canonical Decision.status
系统 SHALL 将 `pending decision` 的判定继续冻结为 `Decision.status` 的 canonical 语义，不得把 `decision_links`、`review_records` 或页面局部状态提升为并列事实源。

#### Scenario: pending 判定仍基于 proposed
- **WHEN** `Dashboard / Current Focus / Daily Review` 读取待处理决策
- **THEN** 当前阶段必须继续以 `Decision.status = proposed` 作为 pending 的最小判定入口
- **AND** 不得因为存在 `decision_links` 就把该决策自动视为“已处理”
- **AND** 不得因为存在 `review_records` 就把该决策自动视为“已完成”

### Requirement: Decision Detail 必须提供正式状态推进入口
系统 SHALL 将 `Decision Detail` 冻结为当前阶段唯一承接正式 `Decision.status` 推进的 canonical 页面。

#### Scenario: 从详情页推进状态
- **WHEN** 用户进入某条 `Decision` 的详情页
- **AND** 当前 `Decision.status` 处于允许推进的非终态
- **THEN** 页面必须展示正式状态推进 CTA
- **AND** CTA 必须通过单一 mutation owner 走后端正式写链
- **AND** 不得只在前端局部 UI 中制造“已处理”假象

### Requirement: 状态推进必须走正式后端写链
系统 SHALL 为 `Decision.status` 推进提供 `.proto` 驱动的正式写链，而不是旁路式本地补丁。

#### Scenario: 状态推进写链落地
- **WHEN** 用户触发 `Decision Detail` 中的状态推进动作
- **THEN** 写入必须通过 `.proto` 新增的正式写接口进入 Connect server
- **AND** Connect server 必须调用 `CommandService` 的单一承接位
- **AND** `CommandService` 必须调用 `DecisionStore` 的单一持久化入口
- **AND** 最终数据库中的 `decisions.status` 必须成为唯一持久化结果

### Requirement: 最小状态推进矩阵必须单值化
系统 SHALL 在当前 fix 中继续使用既有四态，并冻结最小可用状态推进矩阵。

#### Scenario: allowed transitions
- **WHEN** `Decision.status = proposed`
- **THEN** 当前阶段至少必须允许推进到：
  - `active`
  - `superseded`
  - `archived`

#### Scenario: active 后续退出
- **WHEN** `Decision.status = active`
- **THEN** 当前阶段至少必须允许推进到：
  - `superseded`
  - `archived`

#### Scenario: 终态不可继续推进
- **WHEN** `Decision.status = superseded` 或 `archived`
- **THEN** 页面不得继续暴露新的状态推进 CTA
- **AND** 必须把它们视为当前阶段的终态

### Requirement: 状态推进成功后必须统一 reread
系统 SHALL 在状态推进成功后统一刷新详情、列表、review 与 dashboard 相关读取，避免用户看到分裂语义。

#### Scenario: reread after status update
- **WHEN** 某条 `Decision` 在详情页中成功完成状态推进
- **THEN** 至少必须失效并重新读取：
  - 当前 `decision-detail`
  - `decision-list`
  - `Daily Review / Weekly Review`
  - `dashboard-feedback-signals`
  - `dashboard-overview`
- **AND** `Dashboard / Current Focus / Daily Review` 不得继续把该条已退出 pending 的决策显示为待处理

## MODIFIED Requirements
### Requirement: Decision Detail 页面职责冻结
`Decision Detail` 在当前阶段 SHALL 继续作为单一详情页壳层，同时承接：

1. 详情读取
2. 已关联目标展示
3. 待关联目标承接
4. 候选读取
5. `LinkDecisionToTarget`
6. `Decision.status` 的正式状态推进

补充约束：

- 状态推进必须收敛到详情页内部的固定 mutation owner；
- 不得把 `Decision Detail` 扩写为第二套 review-local 工作台；
- 不得拆出新的“决策处理中心”子路由或并列页面；
- `LinkDecisionToTarget` 与状态推进必须是两条不同的 canonical 写动作，不得互相代理语义。

#### Scenario: detail 仍保持 canonical owner
- **WHEN** 用户从 `Dashboard` 或 `Daily Review` 进入 `Decision Detail`
- **THEN** 必须能在同一页面内看到当前决策状态与正式推进动作
- **AND** 不需要回到 review 页面才能完成状态推进
- **AND** 也不得在 review 页面内联一套新的 decision 编辑能力

### Requirement: Review 完成区职责边界
`SubmitReviewResult` 在本 fix 后 SHALL 继续只承接 review 留痕与 next-step 建议，不得升级为 `Decision.status` 的代理写入口。

#### Scenario: review result 不能代替 decision status update
- **WHEN** 用户执行 `submit_next_step`
- **THEN** 系统只允许写入 review result
- **AND** 不得同时偷偷推进 `Decision.status`
- **AND** 若需要正式改变决策状态，必须进入既有 `Decision Detail` canonical 页面完成

### Requirement: LinkDecisionToTarget 职责边界
`LinkDecisionToTarget` 在本 fix 后 SHALL 继续只承接正式目标关联，不得被解释为“决策已处理完成”。

#### Scenario: linked target is not status completion
- **WHEN** 用户在 `Decision Detail` 中完成 `LinkDecisionToTarget`
- **THEN** 系统只允许更新关联结果
- **AND** 不得因为存在 link 就自动把 `Decision.status` 从 `proposed` 改为其他状态
- **AND** 若该决策还需退出 pending，必须由用户显式执行状态推进动作

## REMOVED Requirements
### Requirement: 使用过程记录或关联记录代理 pending 退出
**Reason**: `fix_002` 已确认，`decision_links` 与 `review_records` 都不是 `Decision.status` 的 canonical 来源；继续允许这种代理语义会制造第二事实源。
**Migration**: 所有 pending 退出与状态完成语义统一回收至 `Decision.status` 的正式推进写链，不保留读侧补丁逻辑。
