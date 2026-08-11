# phase08_operating_review_loop_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase08` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase08_operating_review_loop_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase08` 已建立正式 `/plan` 入口，作为 `mvp0.3` 的首个正式业务 phase
- `phase08` 的规划上游统一以 `PSCO-mvp03-summarize-feedback.md`、`phase07-07` 正式规格与 `phase07-11` 验收结论为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp03-summarize-feedback.md`
  - `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md`
  - `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md`
  - `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
  - `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`
  - `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
  - `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
  - `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
- 当前阶段只承接 `Operating Review Loop`
- 当前阶段不反向重写 `mvp0.3` 已冻结的方向排序

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web`
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http + ConnectRPC`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- Contract Tooling：`buf build / lint / generate / breaking`
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段新增或演进后的业务接口必须遵守：
  - `.proto` 是唯一长期合同源
  - `ConnectRPC` 是业务接口正式传输层
  - `chi` 只承担 router shell、middleware 与非业务端点承载职责
- 当前阶段新增或演进后的前端业务动作必须遵守：
  - 写路径唯一 `application` 入口
  - `query` 层纯只读
  - mutation 固定承接位
- 当前阶段不允许：
  - 把 review loop 写成通用任务管理器
  - 把 `Template Reuse`、`Derived Intelligence Deepening`、`Real-Project Dry-Run` 偷渡为并列主交付
  - 引入新的长期核心实体
  - 回退 `phase07` 已收口的 ConnectRPC 正式主线

### 2.5 当前阶段交付模式

- `phase08` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 review loop 业务主线代码

## 3. 当前阶段动作矩阵

`phase08` 最少需要直接承接：

- `EnterDailyReviewFromDashboard`
- `EnterWeeklyReviewFromDashboard`
- `AssembleCurrentFocusAndPendingDecisions`
- `RouteFeedbackIntoDecisionAction`
- `WriteBackReviewResultToExistingEntities`
- `ReadReviewContextThroughSingleOwner`
- `KeepDecisionAsOperatingCenter`
- `RunDashboardToDecisionBrowserAcceptance`

当前阶段必须打通的最小闭环：

- `Dashboard -> Review Entry`
- `Current Focus / Feedback Signals / Pending Decisions -> Unified Action`
- `Feedback -> Decision -> Update`
- `Review Result -> Existing Entity Writeback`
- `Browser Validation -> Root Sync`

允许以最小连接位承接但不扩写为独立主线：

- 为 review 提供上下文的轻量读模型
- 必要的 review 记录或结果落点
- 为后续 phase 预留但当前不启用的扩展字段

### 3.1 当前真实入口 / Caller / Owner Inventory

为避免后续 `/spec` 只描述理想路径而漏掉真实调用方，当前阶段冻结以下 inventory 作为最小执行清单：

- `DashboardRoute (/dashboard)` -> `DashboardHomePage`
  - 页面宿主：`frontend/src/routes/dashboard.tsx`
  - 页面编排：`frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - 当前只读 owner：
    - `useDashboardOverviewRead`
    - `useFeedbackSignalsRead`
    - `useRecentActivitiesRead`
  - 当前真实入口组件：
    - `DashboardPrimaryActionPanel`
    - `CurrentFocusSection`
    - `AssetFeedbackSection`
    - `RecentActivitySection`
    - `DashboardStatBar`
    - `OnboardingCtaButton`
    - `SovereigntyPanel`

- Dashboard 到 Decision 的既有跳转 caller
  - `DashboardPrimaryActionPanel`：主 CTA 命中后跳转 `/decisions` 或 `/decisions/$decisionId`
  - `FeedbackSignalCard`：按 `target_type` 跳转 `/decisions`、`/decisions/$decisionId`、`/products/$productId`、`/modules/$moduleId`、`/repositories/$repositoryId`
  - `RecentActivityItemCard`：按 `target_type` 跳转既有详情页或列表页
  - `DashboardStatBar`：从 overview 区块跳转既有 list 页

- `Decision` 切片当前真实 owner
  - route 宿主：
    - `frontend/src/routes/decisions/index.tsx`
    - `frontend/src/routes/decisions/new.tsx`
    - `frontend/src/routes/decisions/$decisionId.tsx`
  - 页面宿主：
    - `DecisionListPage`
    - `DecisionCreatePage`
    - `DecisionDetailPage`
  - 只读 owner：
    - `useDecisionListRead`
    - `useDecisionDetailRead`
    - `useDecisionModuleCandidatesRead`
  - 写入 owner：
    - `useCreateDraftDecision`
    - `useLinkDecisionToTarget`

- 既有“Decision action handoff” caller
  - `ModuleDecisionEntryPanel`：从 Module Detail 进入 `/decisions/new`
  - `OnboardingPage`：首轮决策录入后进入 `/decisions/$decisionId`
  - `DecisionListToolbar`：进入 `/decisions/new`

- Dashboard 来源上下文当前统一承接位
  - `dashboard-source.ts`
  - `BackToDashboardButton`
  - 当前既有 `modules / products / repositories / decisions` route 的 `validateSearch`

当前阶段后续 `/spec` 必须直接消费上述 inventory，逐项说明：

- 哪些 caller 被复用
- 哪些 caller 被升级为 review 正式入口
- 哪些 caller 保持只读跳转，不承担正式写路径
- 哪些 owner 被扩展，哪些 owner 明确禁止改成页面内联编排

### 3.2 Daily / Weekly Review 差异冻结

当前阶段虽然只交付一个 `Operating Review Loop`，但必须明确 daily / weekly review 不是“同一个页面加两个按钮”：

- `daily review`
  - 目标：快速消费 current focus、pending decision 与代表性缺口，形成当天下一步动作
  - 数据优先级：`current_focus_signals` > `pending decisions` > `representative_signals`
  - 预期输出：进入单个 `Decision`、进入单个实体更新、或记录一条明确的下一步动作结果
  - 第一阶段允许与 Dashboard 共用同一页面骨架，但动作语义必须偏向“马上处理”

- `weekly review`
  - 目标：对最近活动、资产覆盖缺口、决策积压与复用快照做周期性整理
  - 数据优先级：`overview / recent activities / representative signals / reuse snapshot`
  - 预期输出：形成一组需要进入 `Decision` 或回流既有实体的后续动作
  - 第一阶段允许共用 review 骨架，但必须体现周期性整理语义，而不只是 daily queue 的放大版

- 当前阶段冻结要求
  - daily / weekly 可以共用 route 宿主或页面壳层
  - 但不得共用完全相同的数据装配、标题语义、验收口径与完成定义
  - `/spec` 必须明确两者的输入、输出与最小 UI 差异

### 3.3 Review Result Writeback Matrix

当前阶段冻结以下“review 结果回流矩阵”，避免实现时把 action handoff 做轻成纯跳转，或做重成长出并列状态体系：

- `Decision`
  - 正式回流方式：复用 `Decision Center` canonical owner
  - 当前允许动作：
    - 新建 decision draft
    - 进入既有 decision detail
    - 在既有 decision detail 内完成目标关联
  - 当前禁止动作：
    - 在 Dashboard 或 review 页直接内联一套 decision 写表单 owner
    - 在 review 记录中复制一套并列 decision 状态

- `Module`
  - 正式回流方式：通过既有 `Decision -> Module` link 写路径，或跳回既有 module detail canonical 更新入口
  - 当前允许动作：
    - 从 decision detail 关联既有 module
    - 从 review 导航到 module detail 承接后续更新
  - 当前禁止动作：
    - 在 review 页直接内联 module 结构化写入
    - 新长出 review-local module 状态

- `Product`
  - 正式回流方式：优先回到既有 product canonical detail / update 入口
  - 当前允许动作：
    - 从 review 导航到 product detail 或 list 以承接后续更新
    - 若后续确需直接更新，必须复用既有 product application owner
  - 当前禁止动作：
    - 为 review 单独创建并列 product 写路径
    - 在 review 记录中持久化影子 product 状态

- `Repository`
  - 正式回流方式：优先回到既有 repository canonical detail / update 入口
  - 当前允许动作：
    - 从 review 导航到 repository detail 或 list 以承接后续更新
    - 若后续确需直接更新，必须复用既有 repository application owner
  - 当前禁止动作：
    - 为 review 单独创建并列 repository 写路径
    - 在 review 记录中持久化影子 repository 状态

- 全局约束
  - `review result writeback` 的最小必做是：`Decision` 正式承接 + 至少一种实体回流落地
  - `Product / Repository` 在 `phase08` 可先以 canonical action handoff 落地，不强制要求本阶段新增专用直写 API
  - 若本阶段新增任何实体直写，必须复用既有 application owner，不能长出 review-local owner

## 4. 当前阶段页面与业务矩阵

本阶段必须保持可用并完成主线演进的页面 / 模块：

- `Dashboard Home`
- `Decision Center`
- `Product Registry`
- `Module Registry`
- `Repository Binding`
- `Feedback Signals`

### 4.1 当前阶段交互归属矩阵

- `Dashboard`：承接 review 入口与经营动作入口
- `Decision Center`：承接 decision 创建、更新与回流语义
- `Product / Module / Repository`：继续承接 review 结果回流后的既有实体更新
- `Feedback Signals`：继续承接代表性反馈事实输入

本阶段变化不在于新增大量页面，而在于把既有页面与动作重组为正式经营回路。

## 5. 当前阶段合同与承接矩阵

### 5.1 当前阶段合同基线

- review 相关接口继续以 `.proto` 为唯一长期合同源
- review 读取与动作承接继续走 ConnectRPC
- 浏览器侧业务访问前缀继续统一为 `/api`
- domain error、proto error code 与 Connect error 维持单值映射

### 5.2 当前阶段读写边界基线

- `Dashboard review read`：只承接 current focus、feedback signals、pending decisions 与 review context 的读取和解包
- `Review action application owner`：只承接 decision action、结果提交、实体回流、失效刷新与错误归一化
- 展示组件：只消费已经收敛的 read / application owner，不自行编排业务动作
- route：只承接页面装配与参数传递，不编排正式写路径
- page：只承接状态装配与交互组合，不直接成为新的 mutation owner
- 当前阶段若需要 review 记录或 review session 落点，必须明确其身份是“经营回路过程记录”，而不是新的长期核心实体

## 6. 当前阶段验收基线

- `buf build / lint / generate` 通过
- `go build ./...` 通过
- `npx tsc -b --noEmit` 通过
- `frontend build` 通过
- review 相关 Connect / `/api` 关键路径 smoke 通过
- 用户已能从 Dashboard 进入 review
- 用户已能从 review 进入 `Decision`
- 用户已能将 review 结果回流到既有实体或既有 canonical action handoff
- 浏览器端关键路径已通过
- 根级规则、`phase08` 三件套与根级真相源口径一致
- 后续支撑能力 phase 与 dry-run phase 尚未提前写成正式既成事实

## 7. 当前阶段非目标

- 不在本阶段直接实现 `Template Reuse`
- 不在本阶段直接实现 `Derived Intelligence Deepening`
- 不在本阶段直接执行真实项目 `dry-run`
- 不在本阶段引入 `Venture / Decision Intelligence / AI Context Enhancement`
- 不在本阶段引入第二套路由、第二套状态管理或第二套 UI 事实源
