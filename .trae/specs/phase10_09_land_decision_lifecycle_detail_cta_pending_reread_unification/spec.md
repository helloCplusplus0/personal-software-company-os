# phase10-09 落实 `Decision` 生命周期闭环、detail CTA 与 pending reread 统一 Spec

## Why

`phase10-03` 已冻结 `Decision` 四态生命周期、`Decision Detail` CTA 边界与 `Dashboard / Daily Review / Current Focus` 的 canonical pending 语义，但当前系统仍缺少一条真正落地的实现主线：`Decision Detail` 还没有成为唯一正式状态推进承接位，`Dashboard / Daily Review` 也还没有在成功后用同一套 reread 结果完成 pending 收口。  
因此，`phase10-09` 必须把“状态推进入口、详情页 CTA、成功回流、跨页面 reread 与浏览器行为”一起落成单值实现规格，避免继续出现“有提示但无动作出口”或“已处理却仍误报”。

## What Changes

- 落实 `Decision Detail` 作为唯一正式 `Decision.status` 推进承接位
- 落实 `Decision Detail` 在 `proposed / active / superseded / archived` 四态下的真实 CTA 矩阵与页面行为
- 落实 `Dashboard / Daily Review / Current Focus` 对 pending decision 的统一 handoff 与统一 reread
- 落实状态推进成功后的失效刷新、返回来源与成功回流语义
- 落实失败态、终态、重复点击与浏览器刷新时的最小行为约束
- 明确本子任务不改写 `Product / Module / Repository Detail` 的独立 CTA inventory

## Impact

- Affected specs:
  - `phase10_03_freeze_decision_lifecycle_detail_cta_pending_semantics_matrix`
  - `phase10_05_design_dashboard_review_detail_next_step_cta_handoff`
  - `phase10_06_design_frontend_read_write_owner_route_caller_success_handoff`
  - `phase10_07_design_backend_contract_service_read_model_handoff`
  - `docs/fix/fix_002_decision_pending_signal_semantics_analysis.md`
  - `docs/fix/fix_003_decision_detail_status_advance_analysis.md`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
- Affected code:
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/decision-center/data/use-decision-detail-read.ts`
  - `frontend/src/features/decision-center/application/use-update-decision-status.ts`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/dashboard/components/current-focus-section.tsx`
  - `frontend/src/features/dashboard/data/use-feedback-signals-read.ts`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/data/use-daily-review-read.ts`
  - `frontend/src/features/review/application/use-review-action.ts`
  - `backend/internal/decisioncenter/service/command_service.go`
  - `backend/internal/dashboard/candidate/feedback_readers.go`
  - `backend/internal/dashboard/service/query_service.go`
  - `backend/internal/review/service/query_service.go`

## ADDED Requirements

### Requirement: `Decision Detail` 必须成为唯一正式生命周期推进承接位

系统 SHALL 将 `Decision Detail` 落实为当前阶段唯一正式承接 `Decision.status` 推进的页面与写动作入口。`Dashboard / Daily Review / Current Focus` 只负责把用户 handoff 到该详情页，不得继续内联第二套状态推进写路径。

#### Scenario: 从 pending 入口进入 `Decision Detail`

- **WHEN** `Dashboard / Daily Review / Current Focus` 读取到 `Decision.status = proposed`
- **THEN** 该入口的正式主动作必须进入对应 `Decision Detail`
- **AND** 不得在卡片、列表项或 review footer 里直接把 `Decision` 推进到新状态
- **AND** 不得把 `submitReviewResult`、`decision_links` 或页面局部隐藏动作解释为正式状态推进

#### Scenario: `proposed` 状态下的详情页 CTA

- **WHEN** 用户进入一条 `status = proposed` 的 `Decision Detail`
- **THEN** 页面必须提供以下正式状态推进 CTA：
  - `Mark Active`
  - `Mark Superseded`
  - `Archive`
- **AND** 页面允许保留 canonical 导航、目标消费与来源返回类次 CTA
- **AND** 页面不得再缺少正式退出 pending 的动作出口

#### Scenario: `active` 状态下的详情页 CTA

- **WHEN** 用户进入一条 `status = active` 的 `Decision Detail`
- **THEN** 页面只允许继续提供：
  - `Mark Superseded`
  - `Archive`
- **AND** 不得继续展示 `Mark Active`
- **AND** 页面展示的其他 CTA 只能承接 canonical 导航、结果消费与来源返回

#### Scenario: 终态下的详情页 CTA

- **WHEN** 用户进入一条 `status = superseded` 或 `archived` 的 `Decision Detail`
- **THEN** 页面不得继续展示状态推进 CTA
- **AND** 页面只允许保留 canonical 导航、结果消费与来源返回类 CTA
- **AND** 不得继续给出会让用户误判为“当前仍待处理”的按钮或文案

### Requirement: pending reread 必须完全锚定 canonical `Decision.status`

系统 SHALL 让 `Dashboard / Daily Review / Current Focus` 对 pending decision 的解释继续完全锚定 canonical `Decision.status`，并在 `Decision Detail` 状态推进成功后通过同一套 reread 结果完成收口。

#### Scenario: `proposed` 继续视为 pending

- **WHEN** `Decision.status = proposed`
- **THEN** `Dashboard / Daily Review / Current Focus` 必须继续把该决策解释为 pending
- **AND** 若展示下一步动作，正式动作必须统一指向该 `Decision Detail`

#### Scenario: `active / superseded / archived` 不再视为 pending

- **WHEN** 某条 `Decision` 在 `Decision Detail` 中被正式推进到 `active / superseded / archived`
- **THEN** `Dashboard / Daily Review / Current Focus` 在下一次 reread 中都必须停止将其显示为 pending
- **AND** 不得出现某个页面已收口、另一个页面仍误报 pending 的分裂语义

#### Scenario: 禁止使用代理条件退出 pending

- **WHEN** 用户只完成以下任一动作：
  - 建立 `decision_links`
  - 提交 `review_records`
  - 浏览 `Decision Detail`
  - 页面局部隐藏一条 pending 卡片
- **THEN** 这些动作都不得单独被解释为退出 pending 的正式条件
- **AND** pending 的退出仍只能来自 canonical `Decision.status` 离开 `proposed`

### Requirement: 状态推进成功后的回流与失效刷新必须单值化

系统 SHALL 将 `Decision Detail` 状态推进成功后的回流、query 失效与 reread 语义单值化，使详情页与来源页共同消费同一结果，而不是靠 toast 或局部 optimistic hide 制造“已经处理”的假象。

#### Scenario: 详情页内的成功回流

- **WHEN** 用户在 `Decision Detail` 点击任一正式状态推进 CTA 并成功
- **THEN** 当前详情页必须立即 reread 并显示最新 `Decision.status`
- **AND** CTA 矩阵必须同步切换到新状态对应的正式集合
- **AND** 若当前详情页带有来源上下文，页面仍必须保留稳定的“返回来源”入口

#### Scenario: 返回来源后的 reread 收口

- **WHEN** 用户从完成状态推进后的 `Decision Detail` 返回 `Dashboard / Daily Review / Current Focus` 来源页
- **THEN** 来源页必须基于重新读取后的 canonical facts 更新：
  - pending count
  - pending card / current focus 主动作
  - 已生效或历史状态的展示语义
- **AND** 原本那条 pending decision 不得继续残留在待处理区块

#### Scenario: 失败与重复点击行为

- **WHEN** 状态推进请求失败、终态重复触发，或用户在 pending 中重复点击同一 CTA
- **THEN** 页面必须保持当前 canonical 状态不变
- **AND** 不得提前从 `Dashboard / Daily Review / Current Focus` 中移除该决策
- **AND** 错误提示与 pending 态恢复必须由正式 owner 统一承接

### Requirement: 浏览器级验收必须覆盖 detail、dashboard 与 review 的一致性

系统 SHALL 在浏览器级验收中覆盖 `Decision Detail`、`Dashboard`、`Daily Review` 与 `Current Focus` 的完整闭环，确保本子任务交付的是可运行行为，而不是只在静态代码中存在的规则。

#### Scenario: 从 `Dashboard` 完成一条 pending decision

- **WHEN** 用户从 `Dashboard` 或 `Current Focus` 进入一条 `proposed` 的 `Decision Detail`
- **AND** 在详情页将其推进到 `active / superseded / archived` 之一
- **THEN** 返回 `Dashboard` 后，该决策必须不再作为 pending 展示
- **AND** `Dashboard` 主 CTA / `Current Focus` 必须切换到新的 canonical 下一步动作

#### Scenario: 从 `Daily Review` 完成一条 pending decision

- **WHEN** 用户从 `Daily Review` 进入一条 `proposed` 的 `Decision Detail`
- **AND** 在详情页完成正式状态推进
- **THEN** 返回 `Daily Review` 后，该决策必须从 pending 区域消失
- **AND** `submitReviewResult` 相关流程不得再与 `Decision.status` 形成第二套冲突语义

#### Scenario: 非目标边界保持成立

- **WHEN** 完成 `phase10-09` 实现与浏览器验收
- **THEN** `Product / Module / Repository Detail` 的独立 CTA inventory 不得被本子任务顺带改写
- **AND** 本子任务只允许影响它们对 `Decision` 结果消费的既有 canonical 行为，不得扩写出新的页面级动作矩阵

## MODIFIED Requirements

### Requirement: `Decision Detail` 页面职责从“只读详情页”升级为“生命周期推进页”

系统 SHALL 将当前 `Decision Detail` 的正式职责修改为：同时承接详情读取、状态推进 CTA、成功后 reread、来源返回与 canonical 导航，而不是继续停留在“只读摘要 + 建立关联”的半成品状态。

#### Scenario: 页面职责升级后的最小组成

- **WHEN** 后续实现 `Decision Detail`
- **THEN** 页面至少必须稳定承接：
  - `Decision` 当前状态展示
  - 四态对应 CTA 矩阵
  - 状态推进提交中的 pending / error / success
  - 来源返回入口
  - 成功后的 reread
- **AND** 不得继续把状态推进散落到其他页面或 review footer

### Requirement: `Dashboard / Daily Review / Current Focus` 必须从“提示 pending”升级为“提示并 handoff 到唯一动作出口”

系统 SHALL 将这三个消费面修改为：不仅继续提示 pending decision，还必须把用户稳定 handoff 到唯一动作出口 `Decision Detail`，并在回流后用统一 reread 结果关闭 pending。

#### Scenario: 页面职责升级后的正式边界

- **WHEN** 后续实现 `Dashboard / Daily Review / Current Focus`
- **THEN** 它们只允许承接：
  - pending 决策的读取与展示
  - 指向 `Decision Detail` 的主动作 handoff
  - 返回后的 reread 结果消费
- **AND** 不得继续承接新的决策状态写路径

## REMOVED Requirements

### Requirement: 页面局部代理退出 pending

**Reason**: 通过 `decision_links`、`review_records`、本地 dismiss 或局部 optimistic hide 代理 pending 退出，会制造第二套语义真相源，并让 `Decision Detail / Dashboard / Daily Review / Current Focus` 出现分裂解释。  
**Migration**: 统一回收到 canonical `Decision.status` 主线；所有 pending 退出都必须通过 `Decision Detail` 的正式状态推进完成，并由详情页与来源页共同 reread 收口。
