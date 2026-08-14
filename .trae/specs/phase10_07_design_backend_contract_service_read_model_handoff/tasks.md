# Tasks

- [x] Task 1: 冻结 `phase10` 后端“复用既有 canonical 写合同”主线
  - [x] SubTask 1.1: 明确 `Onboarding` 四类创建动作继续复用 `Product / Repository / Module / Decision` 既有 canonical RPC
  - [x] SubTask 1.2: 明确关系闭合动作继续复用 `BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository / LinkDecisionToTarget`
  - [x] SubTask 1.3: 明确 `Decision` 生命周期推进继续只复用 `UpdateDecisionStatus`

- [x] Task 2: 冻结 `Onboarding` 的最小恢复辅助读模型设计
  - [x] SubTask 2.1: 明确现有 `GetFirstRunState` 保留为冷启动摘要读取，不扩写为 phase10 全量恢复总读模型
  - [x] SubTask 2.2: 明确若现有合同不足时，允许在 `psco.onboarding.v1` 内新增单一恢复辅助 read RPC 的最小字段边界
  - [x] SubTask 2.3: 明确 `backend/internal/onboarding/service/query_service.go` 与 `candidate/*` 的承接位
  - [x] SubTask 2.4: 明确 `current_product_id` 的正式事实源层级，以及只有在 canonical facts 候选数不为 `1` 时才允许新增最小恢复锚点 store

- [x] Task 3: 冻结 `Decision` pending reread 的后端合同与服务边界
  - [x] SubTask 3.1: 明确 `pending` 继续只锚定 `Decision.status = proposed`
  - [x] SubTask 3.2: 明确 `DecisionCenterService.GetDecisionDetail / ListDecisions / UpdateDecisionStatus` 足以承接 phase10 决策主线
  - [x] SubTask 3.3: 明确状态推进后的 reread 继续回到 `DecisionCenter / Dashboard / Review` 既有读取，不新增专用 reread RPC

- [x] Task 4: 冻结 `Current Focus / pending signals / next-step CTA` 的读模型与 query owner 设计
  - [x] SubTask 4.1: 明确 `Dashboard.QueryService` 继续是 `current_focus_signals` 的正式后端 owner
  - [x] SubTask 4.2: 明确 `Review.QueryService` 继续只作为组合 owner，复用 `Dashboard` 与 `DecisionCenter` 的 canonical 读取
  - [x] SubTask 4.3: 明确 `FeedbackSignal` 的复用门槛 checklist，以及任一门槛不满足时只允许做既有 message 的最小字段演进
  - [x] SubTask 4.4: 明确 `review.proto` 只能复用或透传 dashboard 的 next-step 描述，不复制第二套 review-local next-step message

- [x] Task 5: 冻结 `.proto / Connect / service / store` 的最小物理落点
  - [x] SubTask 5.1: 列出允许新增的最小落点，包括 `onboarding.proto`、`dashboard.proto`、对应 QueryService、candidate reader 与可选恢复锚点 store
  - [x] SubTask 5.2: 列出必须继续复用的 Connect transport 与 `platform/router.go` 挂载位
  - [x] SubTask 5.3: 明确不允许新增 `backend/internal/phase10/`、第二个 pending 服务或第二套 decision lifecycle 服务

- [x] Task 6: 冻结“新增 vs 复用”单值判定规则
  - [x] SubTask 6.1: 明确什么情况下必须直接复用既有 canonical facts / contract / service
  - [x] SubTask 6.2: 明确什么情况下才允许新增最小字段或最小辅助读模型
  - [x] SubTask 6.3: 明确哪些提议应直接判定为越界，例如影子状态表、第二套 pending 字段、phase10-local 读服务

- [x] Task 7: 完成 `phase10-07` 三件套自检并对齐上游
  - [x] SubTask 7.1: 复核 `spec.md` 是否正确承接 `phase10-03 / 04 / 05 / 06` 与 `phase08-07 / phase09-07`
  - [x] SubTask 7.2: 复核 `tasks.md` 与 `checklist.md` 是否完整覆盖合同复用、新增边界、QueryService owner、`current_product_id` 恢复判定与 `FeedbackSignal` 复用门槛
  - [x] SubTask 7.3: 确认三件套已能机械回答“这里该复用既有合同，还是允许新增最小承接位”

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 2` and `Task 4`
- `Task 6` depends on `Task 1` to `Task 5`
- `Task 7` depends on `Task 1` to `Task 6`
