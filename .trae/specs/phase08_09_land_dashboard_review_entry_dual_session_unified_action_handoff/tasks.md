# Tasks

- [x] Task 1: 收口 Dashboard review 入口与首页宿主整合
  - [x] SubTask 1.1: 确认 `DashboardPrimaryActionPanel` 继续作为唯一正式 review 入口 caller，并稳定承接 `/reviews/daily` 与 `/reviews/weekly`
  - [x] SubTask 1.2: 对齐 `DashboardHomePage` / `DashboardHomePageShell` 的宿主责任，确保 review 入口与既有 Dashboard 模块最小整合，不引入第二套工作台
  - [x] SubTask 1.3: 冻结入口进入时的 `fromDashboard / dashboardSection / dashboardReturnTo` 透传与返回链

- [x] Task 2: 收口 Daily Review 的独立会话语义
  - [x] SubTask 2.1: 冻结 `DailyReviewPage` 的最小区块顺序为 `current focus / pending decisions / representative signals / action footer`
  - [x] SubTask 2.2: 冻结 `pending decisions` 继续以 `Decision proposed` 摘要作为正式来源，不在页面层重新拼第二套队列
  - [x] SubTask 2.3: 明确 Daily Review 不消费 `overview / recent activity / reuse snapshot` 作为主区，防止与 Weekly Review 混同

- [x] Task 3: 收口 Weekly Review 对 `phase05 / phase06` 读模型的正式消费
  - [x] SubTask 3.1: 冻结 `WeeklyReviewPage` 的最小区块顺序为 `overview / recent activity / representative signals / reuse snapshot / action footer`
  - [x] SubTask 3.2: 冻结 `overview / recent activity / representative signals` 继续消费 `phase05` 既有 canonical 读模型
  - [x] SubTask 3.3: 冻结 `reuse snapshot / module_reuse_summary / capability_summary` 继续消费 `phase06` 既有读模型，并保持局部失败边界
  - [x] SubTask 3.4: 回收 Weekly Review 页面层对 query、retry、query key 的临时编排，使重试能力重新回到 read owner

- [x] Task 4: 收口统一动作承接与共享壳层边界
  - [x] SubTask 4.1: 冻结 `ReviewPageShell` 只承接头部、返回、页面级状态与底部动作区
  - [x] SubTask 4.2: 冻结 `ReviewActionFooter + useReviewAction()` 继续作为 daily / weekly 的唯一正式动作承接位
  - [x] SubTask 4.3: 明确纯 handoff、`next-step result`、以及成功 envelope 透传 Dashboard 来源参数的统一语义

- [x] Task 5: 完成 `phase08-09` 关键路径验收
  - [x] SubTask 5.1: 验证 `/dashboard -> /reviews/daily` 与 `/dashboard -> /reviews/weekly` 双入口真实可达
  - [x] SubTask 5.2: 验证 Daily Review 与 Weekly Review 不以同一套数据装配与完成定义冒充双路径
  - [x] SubTask 5.3: 验证 review 页面进入 `Decision / Product / Module / Repository` canonical path 时继续透传 Dashboard 来源参数
  - [x] SubTask 5.4: 验证 `next-step result` 路径继续命中 `ReviewService.SubmitReviewResult`
  - [x] SubTask 5.5: 执行 `(cd frontend && npm run build)`，确认前端 owner 收敛与页面集成未回退

# 实现总结

## 修复项
- `use-daily-review-read.ts`：新增 `retry` 暴露，与 `useWeeklyReviewRead` 对齐，使重试能力回到 read owner
- `review-page-shell.tsx`：新增 `onRetry` prop，page-error 状态增加重试按钮，对齐 dashboard 错误态
- `daily-review-page.tsx` / `weekly-review-page.tsx`：透传 `onRetry={review.retry}` 到 ReviewPageShell
- `review-action-footer.tsx` / `daily-review-page.tsx` / `weekly-review-page.tsx`：将 footer 内部硬编码的 `dashboardSection` 回收到页面侧显式传入的 `actionSections` 映射，确保 `Weekly Review` 的 `decision` / `submit_next_step` 不再沿用 daily 的 `current-focus` 返回链语义；当前 weekly 成功 envelope 已对齐到 `overview / asset-feedback`

## 验证结果
- API 双路径差异验证：
  - Daily Review: currentFocusSignals(3) + pendingDecisions(1) + representativeSignals(2)，无 overview/recent/reuse ✅
  - Weekly Review: overview(✓) + recentActivities(10) + representativeSignals(2) + moduleReuseSummary(2) + capabilitySummary(2)，无 currentFocus/pendingDecisions ✅
  - 确认 daily / weekly 不以同一套数据装配冒充双路径 ✅
- SubmitReviewResult: ✅ (weekly reviewKind 正确记录)
- frontend: `npm run build` ✅
- backend: `go build ./...` ✅
- Dashboard 入口: `DashboardPrimaryActionPanel` 双入口稳定渲染 ✅
- 统一动作: `ReviewActionFooter + useReviewAction()` 统一承接 ✅
- 双会话动作语义: `Daily Review` 保持 `current-focus / asset-feedback` 映射，`Weekly Review` 已切换为 `overview / asset-feedback` 映射，不再把 weekly 回流语义压回 daily ✅
- phase05 读模型: overview / recentActivity / representativeSignals 已正式消费 ✅
- phase06 读模型: moduleReuseSummary / capabilitySummary 已正式消费 ✅

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 1`
- `Task 5` depends on `Task 2`
- `Task 5` depends on `Task 3`
- `Task 5` depends on `Task 4`
