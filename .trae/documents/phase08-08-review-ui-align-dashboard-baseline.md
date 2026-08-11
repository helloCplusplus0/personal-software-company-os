# phase08-08 Review UI 对齐 Dashboard 基线调整 Plan

## Context

phase08-08 已经过 GPT54 源代码与浏览器双重验收通过，但 UI 实现存在三处偏离：

1. **Dashboard 入口 review launcher 与 dashboard 紧凑化基线不一致**：
   - 按钮组缺少响应式降级（移动端会挤压）
   - 主次按钮分级不够清晰（default + outline 平铺，移动端不堆叠）
   - phase08-05 spec §"标题行动区的响应式布局责任"要求"移动浏览器下必须降级为纵向堆叠或紧凑双按钮"

2. **Daily / Weekly Review 详情页用 Card 重型卡片，偏离 dashboard 的 `section + divide-y` 紧凑列表基线**：
   - dashboard 经过多轮优化已建立稳定的紧凑化规范（phase05-13 / phase06-16），review 页面没有对齐
   - Pending Decision 行不可点击，违背 phase08-05 spec §"最小展示形态必须是'pending / backlog 决策摘要列表 + 进入既有 Decision canonical 页的动作入口'"
   - Representative Signals 行不可点击，未复用既有 FeedbackSignalCard
   - Weekly Review 的 overview 用 `grid-cols-3 sm:grid-cols-6` 在移动端挤压
   - Weekly Review 的 reuse snapshot 与 dashboard 的 ReuseSnapshotSection 风格不一致
   - 违背 project_memory §"新增 UI 组件需遵循 phase05-13 已建立的 text-xs/space-y-2/divide-y 紧凑化规范"

3. **移动端溢出问题**：
   - Weekly Review overview 的 6 列网格在移动端撑破容器
   - ReviewActionFooter 的 5 个按钮 `flex-wrap` 在移动端换行混乱
   - 缺少 `min-w-0` 防御内部 flex 内容撑破

**目标**：在不动 phase08-08 已冻结合同、owner、route 承接位的前提下，将 review 入口与两个详情页 UI 看齐 dashboard 紧凑化基线，并修复移动端响应式。

**Workflow 选择**：用户已确认直接调整实现（不走 fix workflow），按最佳实践标准执行 + tsc/build + 浏览器验收。

## 设计基线对齐参考

dashboard 已建立的紧凑化规范（必须对齐）：

- **壳层**：`<div className="space-y-4">` + 标题行 `<h1 className="text-xl font-bold">` + 主 CTA 内联右侧
- **区块容器**：`<section className="space-y-2" aria-label="...">` + `<h2 className="text-base font-semibold">`
- **紧凑列表**：`<div className="divide-y divide-border overflow-hidden rounded-lg border">`
- **单行卡片**：`<button className="flex w-full items-center gap-2 px-3 py-2 text-left hover:bg-muted/40">`
- **空态**：`<div className="rounded-lg border border-dashed p-4 text-center">` + `<p className="text-xs text-muted-foreground">`
- **错误态**：`<div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">` + `<Button variant="outline" size="sm" className="mt-2 h-7">重试</Button>`
- **骨架**：`<Skeleton className="h-5 w-full" />`
- **stat 数字**：`text-lg font-bold leading-none tabular-nums` + `text-[10px] text-muted-foreground`
- **stat 容器**：`flex flex-wrap items-stretch overflow-hidden rounded-lg border bg-card divide-x divide-border`
- **移动端响应式**：`flex flex-col gap-2 sm:flex-row` / `min-w-0` / `truncate`

## 文件改动清单

### 1. DashboardPrimaryActionPanel（review 入口按钮）

**文件**：`frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`

**改动**：
- 按钮文案保持英文 "Daily Review" / "Weekly Review"（用户确认）
- 主次按钮分级：Daily Review 用 `variant="default"`（主入口），Weekly Review 用 `variant="outline"`（次入口）
- 移动端响应式：容器 `flex flex-col gap-2 sm:flex-row sm:items-center`，按钮 `w-full sm:w-auto shrink-0`
- 按钮高度统一 `h-9`（与 OnboardingCtaButton 协调）
- 图标色调与 dashboard stat bar 一致：Clock 用 `text-amber-500`，Calendar 用 `text-blue-500`

### 2. ReviewPageShell（review 页面壳层）

**文件**：`frontend/src/features/review/components/review-page-shell.tsx`

**改动**：
- 头部添加副标题/说明文案承接位（`subtitle?: string` prop），让 daily/weekly 各自传达会话语义
- 头部紧凑化：`flex flex-col gap-1 sm:flex-row sm:items-center sm:justify-between`
- 返回按钮保持 `variant="ghost" size="sm" h-8 px-2`（与 dashboard stat bar 风格一致）
- 标题保持 `text-xl font-bold`，副标题用 `text-xs text-muted-foreground`
- 整体 `space-y-4`（对齐 dashboard）
- 错误态/加载态样式与 dashboard 一致（`rounded-lg border border-destructive/50 bg-destructive/10 p-3`）

### 3. DailyReviewPage（Daily Review 详情页）

**文件**：`frontend/src/features/review/pages/daily-review-page.tsx`

**改动**：
- 移除所有 `Card / CardHeader / CardContent / CardTitle` 重型卡片
- 改用 dashboard 的 `<section className="space-y-2">` + `<h2 className="text-base font-semibold">` + `<div className="divide-y divide-border overflow-hidden rounded-lg border">` 紧凑列表模式
- 区块顺序保持 phase08-05 spec 冻结的优先级：current focus → pending decisions → representative signals
- **PendingDecision 行可点击**：整行 `<button>` 进入既有 `/decisions/$decisionId`，携带 `buildDashboardSourceParams('current-focus')` 来源参数（spec §"最小展示形态必须是 pending / backlog 决策摘要列表 + 进入既有 Decision canonical 页的动作入口"）
- **Representative Signals 直接复用 `FeedbackSignalCard`**（`section="asset-feedback"`），跳转后回到 dashboard 的 asset-feedback 区块
- **Current Focus 直接复用 `FeedbackSignalCard`**（`section="current-focus"`），跳转后回到 dashboard 的 current-focus 区块
- 区块底部追加"查看全部"链接（如 pending decisions 底部"进入决策列表 →"），调用 `useReviewAction({ actionType: 'go_to_decision', dashboardSection: 'current-focus' })`
- 区块状态空态/错误态/骨架样式与 dashboard 完全一致（border-dashed / border-destructive / Skeleton）
- 移除 `AlertCircle / Lightbulb / FileText` 等大图标，section 标题保持纯文本（与 dashboard 的 `<h2>` 一致）

### 4. WeeklyReviewPage（Weekly Review 详情页）

**文件**：`frontend/src/features/review/pages/weekly-review-page.tsx`

**改动**：
- 移除所有 `Card / CardHeader / CardContent / CardTitle` 重型卡片
- 改用 dashboard 的 `<section>` + `<h2>` + `divide-y` 紧凑列表模式
- **overview 区块**改为紧凑 stat bar 风格：`flex flex-wrap items-stretch overflow-hidden rounded-lg border bg-card divide-x divide-border`，每个 cell `min-w-[68px] flex flex-col justify-center px-3 py-2`，数字 `text-lg font-bold leading-none tabular-nums` + label `text-[10px] text-muted-foreground`（对齐 DashboardStatBar，但不含可点击跳转与 coverage 组，因为 weekly review overview 是只读摘要）
- **recent activity 直接复用 `RecentActivityItemCard`**（先扩展该组件接受可选 `section` prop）
- **representative signals 直接复用 `FeedbackSignalCard`**（`section="asset-feedback"`）
- **reuse snapshot 直接复用 dashboard 的 `ReuseSnapshotSection`**（路径：`frontend/src/features/dashboard/components/reuse-snapshot-section.tsx`），传入 weekly review read owner 的 moduleReuseSummary / capabilitySummary
- 区块顺序保持 phase08-05 spec 冻结：overview → recent activity → representative signals → reuse snapshot → 完成区
- 移动端响应式：所有 flex 容器加 `min-w-0`，长文本 `truncate`
- 移除 `BarChart3 / Clock / Lightbulb / RefreshCw` 等大图标

### 5. ReviewActionFooter（review 完成区）

**文件**：`frontend/src/features/review/components/review-action-footer.tsx`

**改动**：
- 重新分层为主行动区 + 完成区：
  - 主行动区：`flex flex-col gap-2 sm:flex-row sm:flex-wrap sm:items-center`
  - 主按钮「进入决策」（`go_to_decision`，default 风格，最显眼）
  - 次按钮组「产品」「模块」「仓库」（outline 风格，紧凑 `h-8 px-2 text-xs`）
  - 完成区：`mt-3 pt-3 border-t`，"完成 Review" 按钮（submit_next_step，secondary 风格）+ 说明文案"记录一条 next-step 结果到 review_records"
- 移动端按钮纵向堆叠 `w-full sm:w-auto`
- 错误态样式对齐 dashboard 错误态

### 6. RecentActivityItemCard 扩展（让 review 可复用）

**文件**：`frontend/src/features/dashboard/components/recent-activity-item-card.tsx`

**改动**：
- 扩展 props 接受可选 `section?: DashboardSection`（默认 `'recent-activity'`）
- 内部 `handleClick` 使用传入的 section 构造 `buildDashboardSourceParams(section)`
- 让 Weekly Review 复用时可以传入 `section="recent-activity"`（保持原行为，不破坏 dashboard 内调用方）

### 7. 移动端响应式修复清单

逐文件应用以下模式：
- flex 容器：`flex flex-col gap-2 sm:flex-row sm:items-center`（移动端纵向堆叠）
- 按钮宽度：`w-full sm:w-auto`
- 列表容器：`divide-y divide-border overflow-hidden rounded-lg border`（不撑破容器）
- 长文本容器：`min-w-0` + `truncate`
- stat bar：`flex flex-wrap`（不用固定列数 grid）
- 防御性 `min-w-0` 加在所有 flex/grid 子项上

## 复用的既有组件与工具

| 复用组件 | 路径 | 用途 |
|---------|------|------|
| `FeedbackSignalCard` | `frontend/src/features/dashboard/components/feedback-signal-card.tsx` | Daily Review 的 current focus / representative signals 行；Weekly Review 的 representative signals 行 |
| `RecentActivityItemCard` | `frontend/src/features/dashboard/components/recent-activity-item-card.tsx` | Weekly Review 的 recent activity 行（扩展接受 section prop） |
| `ReuseSnapshotSection` | `frontend/src/features/dashboard/components/reuse-snapshot-section.tsx` | Weekly Review 的 reuse snapshot 区块 |
| `buildDashboardSourceParams` | `frontend/src/features/dashboard/lib/dashboard-source.ts` | review 行跳转与 footer 动作的来源参数构造 |
| `useNavigateBackToDashboard` | `frontend/src/features/dashboard/lib/dashboard-source.ts` | review 页面返回 Dashboard |
| `useReviewAction` | `frontend/src/features/review/application/use-review-action.ts` | review 完成区动作触发（不改 owner 接口） |

## 不改动项（phase08-08 已冻结）

- 不改 `ReviewService` proto 合同与 backend 实现
- 不改 `useDailyReviewRead` / `useWeeklyReviewRead` 数据契约
- 不改 `useReviewAction` owner 接口与 success envelope 语义
- 不改 `/reviews/daily` 与 `/reviews/weekly` 路由 search schema
- 不改 `ReviewActionInput` / `ReviewActionSuccess` 类型定义
- 不引入新组件文件（所有改动收敛在既有 6 个文件内）

## 验证清单

### 静态验证
- `cd frontend && npx tsc --noEmit`（无类型错误）
- `cd frontend && npm run build`（构建通过）

### 浏览器验证（启动 dev server 后用 integrated_browser MCP）
1. **Dashboard 入口验证**（`http://localhost:5173/dashboard`）：
   - review launcher 两个按钮稳定可见，主次分级清晰
   - 移动端尺寸（resize 到 375px 宽度）下按钮纵向堆叠，不挤压
   - 桌面端尺寸下按钮并排，与 OnboardingCtaButton 协调

2. **Daily Review 验证**（`http://localhost:5173/reviews/daily`）：
   - 区块视觉对齐 dashboard 紧凑化基线（section + divide-y，无重型 Card）
   - current focus / representative signals 行可点击跳转到对应 canonical 详情页
   - pending decisions 行可点击进入 Decision Detail，跳转后 BackToDashboardButton 可用
   - 移动端尺寸下不溢出，所有按钮可点击

3. **Weekly Review 验证**（`http://localhost:5173/reviews/weekly`）：
   - overview 改为紧凑 stat bar 风格，移动端不挤压
   - recent activity / representative signals 行可点击跳转
   - reuse snapshot 区块视觉与 dashboard AssetFeedbackSection 内的 ReuseSnapshotSection 一致
   - 移动端尺寸下不溢出

4. **ReviewActionFooter 验证**：
   - 主行动区 + 完成区分层清晰
   - 移动端按钮纵向堆叠，桌面端紧凑排列
   - 错误态样式对齐 dashboard

### 回归验证
- Dashboard 页面无视觉变化（仅 primary action panel 容器响应式调整）
- 既有 FeedbackSignalCard / RecentActivityItemCard / ReuseSnapshotSection 在 dashboard 内的行为不变（RecentActivityItemCard 扩展 section prop 默认值保持原行为）
- review 路由 search schema 与 source 参数透传链路不变
