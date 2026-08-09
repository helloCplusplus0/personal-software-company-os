# Tasks

- [x] Task 1: 对齐 `phase05-13` 的直接上游、仓库前端骨架与当前技术栈边界，明确这次任务是“Dashboard 前端主线实现”，不是重写 Dashboard 正式规格或联调验收文档。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase05_dashboard_feedback_foundation_dev_plan.md` 中 `phase05-13` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase05-05` 页面/路由/组件树、`phase05-06` 状态模型、`phase05-10` 正式正文、`phase05-11` Proto 主线、`phase05-12` 后端主线
  - [x] SubTask 1.3: 对齐 `phase03-13` 与 `phase04-13` 的既有前端主线规格模式
  - [x] SubTask 1.4: 对齐当前仓库 `frontend/package.json`、`frontend/src/routes/__root.tsx`、既有 feature 结构与 `frontend/src/gen/proto/psco/dashboard/v1/` 的真实落点
  - [x] SubTask 1.5: 使用 Context7 补查 `TanStack Router` 与 `TanStack Query` 的最新路由搜索参数、navigation state 与 query invalidate 主线口径

- [x] Task 2: 冻结 `Dashboard` 路由、页面落点与根导航接入。
  - [x] SubTask 2.1: 冻结 `frontend/src/routes/dashboard.tsx` 与 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` 的单值落点
  - [x] SubTask 2.2: 冻结根导航必须新增 `Dashboard` 一级入口并进入 `/dashboard`
  - [x] SubTask 2.3: 冻结当前阶段不得把 `/` 重写为 `Dashboard Home`
  - [x] SubTask 2.4: 冻结当前阶段不得把 `/dashboard` 拆成第二套路由树或子页面体系

- [x] Task 3: 冻结 `frontend/src/features/dashboard/` 的数据适配层与合同承接边界。
  - [x] SubTask 3.1: 冻结 `types.ts`、`data/api-adapter.ts`、`data/dashboard-adapter.ts` 的最小落点
  - [x] SubTask 3.2: 冻结 `dashboard-adapter.ts` 必须直接导出真实 API 实现，不得并列 `mock-adapter.ts`
  - [x] SubTask 3.3: 冻结适配层字段语义必须对齐 `frontend/src/gen/proto/psco/dashboard/v1/dashboard_pb.ts`
  - [x] SubTask 3.4: 冻结本地 view model 若存在，只能作为 `.proto / HTTP` 的单向派生结果

- [x] Task 4: 冻结 Dashboard 三路读取的前端主线与 Query 边界。
  - [x] SubTask 4.1: 冻结 `/api/dashboard/overview`、`/api/dashboard/feedback-signals`、`/api/dashboard/recent-activities` 三路读取必须显式存在
  - [x] SubTask 4.2: 冻结 `overviewQueryState / feedbackQueryState / activityQueryState` 的最小查询状态
  - [x] SubTask 4.3: 冻结整页重试、Feedback 分区重试与 Recent Activity 分区重试的粒度
  - [x] SubTask 4.4: 冻结 Dashboard 不得发明第二套缓存协议或手工请求协调器

- [x] Task 5: 冻结 `DashboardHomePage` 四区块与主 CTA 面板的前端消费语义。
  - [x] SubTask 5.1: 冻结页面级组件树继续承接 `Current Focus / dashboard_overview / Asset Feedback / Recent Activity / DashboardPrimaryActionPanel`
  - [x] SubTask 5.2: 冻结 `CurrentFocusSection` 最多展示 `5` 条主队列反馈卡片
  - [x] SubTask 5.3: 冻结 `AssetFeedbackSection` 最多展示 `3` 条代表性缺口项，且不升级为第二主队列
  - [x] SubTask 5.4: 冻结 `DashboardStatBar` 与 `RecentActivitySection` 的成功态、骨架占位与局部失败展示语义
  - [x] SubTask 5.5: 冻结 `DashboardPrimaryActionPanel` 的 `computing / ready / hidden / suppressed` 状态机与唯一主 CTA 规则

- [x] Task 6: 冻结 Dashboard 到 canonical owner 页面的跳转、来源参数与返回上下文承接。
  - [x] SubTask 6.1: 冻结卡片与活动项导航必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`
  - [x] SubTask 6.2: 冻结 `dashboardSection` 只允许 `overview / current-focus / asset-feedback / recent-activity / empty-state`
  - [x] SubTask 6.3: 冻结目标页主动返回 Dashboard 必须通过一次性 navigation state 承接区块恢复
  - [x] SubTask 6.4: 冻结刷新后 owner 页面必须继续从 URL 恢复 Dashboard 来源参数

- [x] Task 7: 冻结 Router 与局部状态边界。
  - [x] SubTask 7.1: 冻结来源恢复必须继续通过 `createFileRoute + validateSearch` 与 `useNavigate` 承接
  - [x] SubTask 7.2: 冻结 `fromDashboard / dashboardSection / dashboardReturnTo` 不得提升为全局 Zustand 事实源
  - [x] SubTask 7.3: 冻结 Dashboard 刷新时必须重新读取三路真实接口，不得恢复旧页面快照
  - [x] SubTask 7.4: 冻结 Zustand 如被复用，只能服务局部 UI 草稿或非事实源级状态

- [x] Task 8: 冻结单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略。
  - [x] SubTask 8.1: 冻结桌面宽屏下 `DashboardStatBar` 置顶、`Current Focus` 第一行动优先级与左右双列并列承接
  - [x] SubTask 8.2: 冻结移动浏览器下 `Current Focus / DashboardStatBar / Asset Feedback / Recent Activity` 的单列重排顺序
  - [x] SubTask 8.3: 冻结不得引入独立移动端页面、独立路由树或第二套移动端 UI 架构

- [x] Task 9: 完成 `phase05-13` 规格一致性校验。
  - [x] SubTask 9.1: 验证本 spec 已把 Dashboard 前端从“设计冻结”推进为“可运行主线”
  - [x] SubTask 9.2: 验证本 spec 已直接引用 `phase05-11 / 12` 真实合同与后端主线，而非并列 mock
  - [x] SubTask 9.3: 验证本 spec 继续守住 Dashboard 只负责读与跳转、不形成第二套写入工作台
  - [x] SubTask 9.4: 验证本 spec 继续守住单一 `React Web` 下的双场景布局，不引入第二套移动端 UI 架构
  - [x] SubTask 9.5: 验证本 spec 与 `phase05-05 / 06 / 10 / 11 / 12` 及既有 `phase03-13 / 04-13` 模式一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 3`
- `Task 5` depends on `Task 2`, `Task 3`, `Task 4`
- `Task 6` depends on `Task 2`, `Task 5`
- `Task 7` depends on `Task 2`, `Task 4`, `Task 6`
- `Task 8` depends on `Task 2`, `Task 5`
- `Task 9` depends on `Task 1` through `Task 8`
