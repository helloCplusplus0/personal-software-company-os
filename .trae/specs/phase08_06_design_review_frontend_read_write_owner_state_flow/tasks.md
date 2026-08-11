# Tasks

- [x] Task 1: 冻结 review 前端切片的正式物理落点
  - [x] SubTask 1.1: 明确 `frontend/src/features/review/data/` 与 `application/` 的最小文件结构
  - [x] SubTask 1.2: 明确 `review` 切片与既有 `dashboard / reuse-summary / decision-center` 切片的边界
  - [x] SubTask 1.3: 明确 route / page / section 只作为 caller，不再并列承接 owner 逻辑

- [x] Task 2: 冻结 daily / weekly review read layer 的正式消费位置
  - [x] SubTask 2.1: 明确 `useDailyReviewRead()` 的输入来源、`pending decisions` 最小口径与 top N 展示边界
  - [x] SubTask 2.2: 明确 `useWeeklyReviewRead()` 对 `overview / recent activity / representative signals / reuse snapshot` 的正式消费位置
  - [x] SubTask 2.3: 明确 review 页面不得直接 import 底层 dashboard / reuse / decision query hooks
  - [x] SubTask 2.4: 明确 review 聚合读层继续建立在底层 canonical query key 之上，不把 `['review']` 冻结成必需主缓存

- [x] Task 3: 冻结 review 页面级与区块级状态流
  - [x] SubTask 3.1: 明确 `initial-loading / ready / page-error` 的页面级状态模型
  - [x] SubTask 3.2: 明确 `ready / empty / error` 的区块级状态模型
  - [x] SubTask 3.3: 明确 daily / weekly 页面如何通过 read owner 消费派生状态，而不是自行重组多 query 状态

- [x] Task 4: 冻结 `Review action application owner` 的接口、职责与成功回流
  - [x] SubTask 4.1: 明确 `submitAction()` 与 `ReviewActionSuccess` 的最小契约
  - [x] SubTask 4.2: 明确 `Decision` 相关动作如何复用既有 canonical application owner
  - [x] SubTask 4.3: 明确实体 handoff、`next-step result` 与页面导航的分工边界
  - [x] SubTask 4.4: 明确所有进入 canonical 页的 success envelope 都必须继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`

- [x] Task 5: 冻结 query 失效策略与错误归一化边界
  - [x] SubTask 5.1: 明确 canonical owner 与 review owner 各自负责的 query 失效范围
  - [x] SubTask 5.2: 明确 `Decision draft`、`Decision link`、纯 route handoff 三类动作的失效矩阵
  - [x] SubTask 5.3: 明确错误归一化留在 `Review action application owner`，页面只展示稳定 review-facing error
  - [x] SubTask 5.4: 明确 `BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository` 三类实体 mutation 的回流与失效矩阵

- [x] Task 6: 识别并冻结必须回收的临时编排点与 caller-owner 映射表
  - [x] SubTask 6.1: 明确 review 不得复制既有 canonical 页面中的页面级 `invalidateDetail` 与组件级 mutation 包装模式
  - [x] SubTask 6.2: 明确 `DashboardPrimaryActionPanel / ReviewPageShell / ReviewActionFooter / BackToDashboardButton` 的 caller-owner 映射
  - [x] SubTask 6.3: 明确 `CurrentFocusSection / PendingDecisionSection / RepresentativeSignalsSection / ReuseSnapshotSection` 的只读消费关系

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 4`
- `Task 6` depends on `Task 2`
- `Task 6` depends on `Task 4`
