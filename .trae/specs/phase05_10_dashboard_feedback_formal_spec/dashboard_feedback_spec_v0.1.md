# PSCO Dashboard + Feedback 规格 v0.1 — 正式规格正文

> **文档定位**：本文档是 `phase05_dashboard_feedback_foundation` 的正式规格正文，作为后续 `phase05-11 / 12 / 13 / 14` 合同落地、实现、联调验收与收口以及 `phase06+` 引用 `Dashboard + Feedback` 主线时的直接上游规格来源。
> **上游收敛**：本文档由 `phase05-01` 到 `phase05-09` 的冻结结论收敛而成，不另立第二套边界。`phase05-01 ~ 09` 在本文档生效后退为追溯来源与证据链，不再承担并列直接执行层入口职责。
> **互链前提**：本文档以 `phase01-06` 的 `mvp_spec_v0.1.md` 为当前阶段唯一执行层总上游，完整承接 `module_registry_spec_v0.1.md` 与 `phase02-12` 验收结论中已交付的 `Module Registry` 边界、`decision_center_spec_v0.1.md` 与 `phase03-14` 验收结论中已交付的 `Decision Center` 边界、`product_repository_binding_spec_v0.1.md` 与 `phase04-14` 验收结论中已交付的 `Product Registry + Repository Binding` 边界，并与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致。
> **状态约束**：`phase04` 已完成收口，项目当前根级阶段为 `phase05_dashboard_feedback_foundation`；本文档是 `Dashboard + Feedback` 的正式规格入口，`phase05` 后续实现、复核与下一阶段规格不得再把 `phase05-01 ~ 09` 当作并列直接执行层入口使用。

---

## 1. 技术路线

本文档继承 `mvp_spec_v0.1.md` §1 与 `product_repository_binding_spec_v0.1.md` §1 的技术路线，聚焦 `Dashboard + Feedback` 当前阶段：

- 项目路线：`Durable System Track`
- 正式运行主线：`React Web + Go Backend + PostgreSQL`
- 前端：`React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`
- 前端交付策略：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI；不引入独立 `React Native` 客户端，`PWA` 仅作可兼容增强方向
- 后端：`Go`，模块化单体优先
- 数据库：`PostgreSQL` 为唯一数据库主线
- 合同：`Contract First`，跨语言标准为 `Protocol Buffers`；当前阶段必须为 `Dashboard + Feedback` 落地最小 `.proto` 合同源
- 合同工具链：`buf build / lint / generate / breaking`
- 部署：`Caddy + systemd`，运行方式 `Single Server First`

> 约束：不得重新解释为 `Product Track`；不得把 `Rust` 写成当前阶段必需项；`Local First` 当前解释为数据所有权优先，不等于切换到 `SQLite`。

---

## 2. 对象范围

### 2.1 核心交付对象（进入 `phase05` 主执行范围）

- `Dashboard Home`：承接聚合读取、反馈信号展示、最近活动展示、空状态引导与跳转入口的统一页面
- `Feedback Signal Layer`：基于既有结构化对象与绑定关系派生出的"行动信号层"，不是独立表，也不是人工录入型对象

### 2.2 派生读取对象（当前阶段必须正式消费）

- `dashboard_overview`：概览聚合读模型
- `pending_decision_signals`：待决策信号读模型
- `product_asset_coverage`：资产缺口覆盖读模型
- `recent_activity_feed`：最近活动流读模型

### 2.3 连接对象（只读或候选读取，不承接写入主线）

- `Module`：作为候选来源、已绑定摘要读取与活动流数据来源
- `Release`：只作为最近活动与反馈信号的数据来源，当前阶段不新增独立 Detail 页面
- `Product`：作为资产缺口判定与活动流数据来源
- `Repository`：作为绑定关系判定与活动流数据来源
- `Decision`：作为待决策信号与活动流数据来源
- `ProductModuleBinding` / `ProductRepositoryBinding` / `ModuleRepositoryMapping`：作为资产缺口判定与活动流数据来源

### 2.4 非归属能力

- `CreateModule`、`CreateRelease`、`CreateProduct`、`CreateRepository`、`RecordDecision`、`LinkDecisionToTarget`、`BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 不归属于 `Dashboard + Feedback` 模块
- Dashboard 模块只能通过最小连接边界读取或校验 canonical 模块数据，不得吸收任一 canonical 模块的主线写入能力
- Dashboard 不形成"从 Dashboard 直接完成绑定/补录"的影子工作台

---

## 3. 页面矩阵

### 3.1 页面主线

| 页面 | 职责 |
| --- | --- |
| `Dashboard Home` | 承接概览读取、反馈信号读取、最近活动读取、空状态引导与跳转入口 |

- `Dashboard Home` 的正式业务入口路由冻结为 `/dashboard`
- `Dashboard` 在当前阶段作为既有主导航中的一级入口新增，不替代根级布局宿主本身
- 当前阶段不把 `/` 单值解释为 `Dashboard Home`；如未来需要把 `/` 重定向到 `/dashboard`，必须在后续 `phase / fix` 中单独冻结

### 3.2 跳转目标页面（canonical owner，只读跳转）

| 页面 | 承接来源 |
| --- | --- |
| `Module Registry / List` | 概览卡片、空状态 CTA、活动项跳转 |
| `Module Detail` | 活动项跳转（含 Release 回落） |
| `Product Registry / List` | 概览卡片、空状态 CTA |
| `Product Detail` | 反馈信号卡片、活动项跳转 |
| `Repository Binding / List` | 概览卡片、空状态 CTA |
| `Repository Binding Detail / Workspace` | 活动项跳转 |
| `Decision Center / List` | 概览卡片、聚合决策信号跳转 |
| `Decision Detail` | 单项决策信号跳转、活动项跳转 |
| `Module Registry / Create` | 空状态 CTA |
| `Product Registry / Create` | 空状态 CTA |
| `Repository Binding / Create` | 空状态 CTA |

> 约束：`Release` 当前阶段不新增独立 Detail 页面，其活动项统一回落到所属 `Module Detail`。`Binding` 活动类型必须至少拆分为 `product_module_binding / product_repository_binding / module_repository_binding`，不得以笼统 `binding` 保留歧义。

### 3.3 页面信息区块

`Dashboard Home` 固定归属四个区块与一个主行动面板：

- `dashboard_overview`：概览卡片区块
- `Current Focus / Next Action`：唯一主行动队列
- `Asset Feedback`：资产缺口补充摘要区块
- `Recent Activity`：独立活动流区块
- `DashboardPrimaryActionPanel`：主 CTA 面板，独立于四区块

> 约束：`Current Focus / Next Action` 是第一屏唯一主行动队列；`Asset Feedback` 是补充摘要区块，不形成第二条独立优先级队列；`Recent Activity` 是独立活动流，不与反馈信号共用排序逻辑。

---

## 4. 区块矩阵

### 4.1 dashboard_overview 区块

- 承接概览卡片展示与点击入口
- 只承接概览区块标题、区块容器与卡片集合布局
- 包含组件：`DashboardOverviewSection` → `DashboardOverviewCardGrid` → `DashboardOverviewCard`

#### 概览卡片字段

| 字段 | 可点击 | 跳转目标 |
| --- | --- | --- |
| `module_count` | 是 | `Module Registry / List` |
| `product_count` | 是 | `Product Registry / List` |
| `repository_count` | 是 | `Repository Binding / List` |
| `decision_count` | 是 | `Decision Center / List` |
| `product_with_repository_count` | 否（仅展示） | — |
| `product_with_module_count` | 否（仅展示） | — |

> 约束：派生覆盖率指标（`product_with_repository_count` / `product_with_module_count`）在当前阶段只作为辅助指标展示，不得为它们分配独立 canonical 跳转目标。

### 4.2 Current Focus / Next Action 区块

- 承接主行动队列区块容器与标题语义
- 只消费归一化后的反馈信号卡片
- 包含组件：`CurrentFocusSection` → `FeedbackSignalCardList` → `FeedbackSignalCard`
- 最多展示 `5` 条主队列反馈卡片

### 4.3 Asset Feedback 区块

- 承接补充摘要区块容器与标题语义
- 只消费 `product_asset_coverage` 的补充摘要
- 包含组件：`AssetFeedbackSection` → `AssetFeedbackList` → `AssetFeedbackItemCard`
- 最多展示 `3` 条代表性缺口项
- 不得与 `Current Focus` 复用成一个无语义差别的区块容器
- 不得在页面级升级为并列主 CTA

### 4.4 Recent Activity 区块

- 承接活动流区块容器与标题语义
- 只按活动时间倒序排序，不参与反馈优先级竞争
- 包含组件：`RecentActivitySection` → `RecentActivityList` → `RecentActivityItemCard`
- 最多展示 `10` 条活动项

### 4.5 DashboardPrimaryActionPanel

- 独立于四区块，直接挂载在 `DashboardHomePageShell` 下
- 只承接主 CTA 优先级矩阵的命中与展示
- 遵守空状态主 CTA 按钮可点模式（不采用整卡模式）
- 同一时刻只展示一个主 CTA

---

## 5. 动作矩阵

### 5.1 直接承接动作

| 动作 | 单值页面 owner | 数据载体 | 后端 owner |
| --- | --- | --- | --- |
| `ViewDashboard` | `Dashboard Home` | — | `DashboardQueryService` |
| `ReadDashboardOverview` | `Dashboard Home` | `dashboard_overview` | `QueryService.ReadOverview` |
| `ReadPendingDecisionSignals` | `Dashboard Home` | `pending_decision_signals` | `QueryService.ReadFeedbackSignal` |
| `ReadProductAssetCoverage` | `Dashboard Home` | `product_asset_coverage` | `QueryService.ReadFeedbackSignal` |
| `ReadRecentActivity` | `Dashboard Home` | `recent_activity_feed` | `QueryService.ReadRecentActivity` |
| `NavigateFromDashboardToOwner` | `Dashboard Home` → canonical owner | — | — |

> 约束：
> - Dashboard 模块只承接读组，不承接任何业务写入接口。
> - `FeedbackSignalRead` 在前端同时服务 `Current Focus` 与 `Asset Feedback` 两个区块，不发明第二套读取状态。
> - `NavigateFromDashboardToOwner` 只负责跳转，不等于 Dashboard 承接写入。

### 5.2 非归属动作

- `CreateModule` / `CreateRelease` / `CreateProduct` / `CreateRepository` / `RecordDecision` / `LinkDecisionToTarget` / `BindModuleToProduct` / `BindRepositoryToProduct` / `MapModuleToRepository` 不在 Dashboard 模块中实现
- Dashboard 模块不得调用或吸收任一 canonical 模块的主线写入能力
- 用户从 Dashboard 跳到 canonical owner 页面完成写入时，该写入必须由 canonical owner 模块的 service 承接，Dashboard 模块不得承接任何写入副作用

---

## 6. 数据模型

### 6.1 直接承接表（只读消费）

- `modules`
- `module_releases`
- `products`
- `repositories`
- `decisions`
- `decision_links`
- `product_modules`
- `product_repositories`
- `module_repositories`

### 6.2 当前阶段必须正式消费的派生读取

- `dashboard_overview`
- `pending_decision_signals`
- `product_asset_coverage`
- `recent_activity_feed`

### 6.3 最小读模型

#### dashboard_overview

至少承接：

- `module_count`
- `product_count`
- `repository_count`
- `decision_count`
- `product_with_repository_count`
- `product_with_module_count`

#### pending_decision_signals

至少承接：

- 信号计数
- 最多 `5` 条进入 `Current Focus / Next Action` 的高优先级项
- 每条信号至少包含：`signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label`

#### product_asset_coverage

至少承接：

- 已完整绑定产品数（`fully_bound_product_count`）
- 缺少模块与仓库双绑定的产品数（`missing_both_bindings_count`）
- 缺少仓库绑定但已有模块绑定的产品数（`missing_repository_binding_count`）
- 缺少模块绑定但已有仓库绑定的产品数（`missing_module_binding_count`）
- 最多 `3` 条代表性缺口项（`representative_signals`）
- 每条缺口项至少包含：`signal_code / priority / target_product_id / target_product_name / action_label`

> 约束：`product missing both bindings` 必须作为独立信号与独立计数字段存在，不得回退为"缺仓库 + 缺模块"两个单缺口的隐式组合语义。同一个双缺口产品不得在代表项语义上同时重复表达为两个单缺口信号。

#### recent_activity_feed

至少承接：

- 最近新增或更新的 `Module / Release / Product / Repository / Decision / product_module_binding / product_repository_binding / module_repository_binding`
- 最多 `10` 条活动项
- 显式活动时间字段 `activity_at`（不依赖隐式 `created_at` 推断最近活动顺序）
- 对应对象类型
- 对应对象跳转信息

### 6.4 反馈信号优先级

当前阶段反馈信号优先级的单值顺序：

1. `pending_decision_signals`（`P1`）
2. `product missing both bindings`（`P2`）
3. `product missing repository binding`（`P3`）
4. `product missing module binding`（`P4`）

- 同优先级内排序默认按"最近需要处理时间优先"；若上游读模型暂不提供该字段，则回退为 `created_at DESC`
- `recent_activity_feed` 不参与反馈优先级竞争
- `Asset Feedback` 不形成第二条独立优先级队列

### 6.5 Feedback 最小语义

- `Feedback` 当前阶段不是独立表，也不是人工录入型对象
- `Feedback` 是基于既有结构化对象与绑定关系派生出的"行动信号"
- 信号最少回答两类问题：
  - 当前哪里不完整
  - 现在该去哪里补

### 6.6 空状态语义

- `product_asset_coverage` 的空状态：当不存在资产缺口或代表性缺口项时，必须返回完整结构，三类缺口计数为 `0`，代表项为空列表，区分成功空态与局部读取失败
- `pending_decision_signals` 的空状态：返回空列表与零值计数，不映射为资源不存在
- `recent_activity_feed` 的空状态：返回空列表，不映射为错误

---

## 7. 聚合读与反馈信号

### 7.1 三类聚合读

| 聚合读 | 角色 | 失败语义 |
| --- | --- | --- |
| `DashboardOverviewRead` | 主聚合读取 | 失败触发整页失败 |
| `FeedbackSignalRead` | 附属聚合读取 | 失败只触发局部失败 |
| `RecentActivityRead` | 附属聚合读取 | 失败只触发局部失败 |

> 约束：
> - `DashboardOverviewRead` 作为主聚合读取，决定页面是否可进入整页成功语义
> - `FeedbackSignalRead` 与 `RecentActivityRead` 作为附属聚合读取，不得提升为第二个"主页面生死开关"
> - 附属聚合局部失败时，允许局部错误展示与局部重试，不强制整页失败
> - 未知反馈结果下不得伪造主 CTA
> - 当前阶段不新增第二套 Dashboard 接口主线，不提前冻结外部埋点、趋势分析或通知中心接口

### 7.2 CTA 优先级矩阵

Dashboard 主 CTA 的唯一优先级矩阵，同一时刻只允许一个主 CTA：

| 顺序 | 命中条件 | 主 CTA 目标 |
| --- | --- | --- |
| CTA 1 | `module_count = 0 && product_count = 0 && repository_count = 0 && decision_count = 0` | `Module Registry / Create` |
| CTA 2 | `module_count = 0 && (product_count > 0 \|\| repository_count > 0 \|\| decision_count > 0)` | `Module Registry / Create` |
| CTA 3 | `module_count > 0 && product_count = 0` | `Product Registry / Create` |
| CTA 4 | `module_count > 0 && product_count > 0 && repository_count = 0` | `Repository Binding / Create` |
| CTA 5 | 存在 `pending_decision_signals` | 最高优先级决策信号落点 |
| CTA 6 | 存在 `product missing both bindings` | 对应 `Product Detail` |
| CTA 7 | 存在 `product missing repository binding` | 对应 `Product Detail` |
| CTA 8 | 存在 `product missing module binding` | 对应 `Product Detail` |
| CTA 9 | 无缺口且有活动数据 | 系统已就绪中性状态，不展示强制主 CTA |

> 约束：
> - 同时存在多类缺口时，只允许按上述顺序展示一个主 CTA，其他动作降级到区块内次级入口
> - 当前阶段不得在空状态下并排给出多个同级主 CTA
> - CTA 1 与 CTA 2 必须区分展示（冷启动空系统 vs 非空缺口状态）
> - "非空但无 Module"的状态不得与冷启动空系统混同

### 7.3 错误语义

| 场景 | 语义 |
| --- | --- |
| `DashboardOverviewRead` 失败 | 整页失败，只允许整页级重试 |
| `FeedbackSignalRead` 失败 | 局部失败，`Current Focus` 与 `Asset Feedback` 进入 `error`，overview 与 recent activity 仍成功 |
| `RecentActivityRead` 失败 | 局部失败，`Recent Activity` 进入 `error`，overview 与 feedback 仍成功 |
| 空系统 | 成功语义，所有计数为 0 |
| 某区块无结果 | 成功空态，返回空列表 |
| 非法查询参数 | 明确失败语义，不得静默跳到错误页面 |

---

## 8. 跳转与返回上下文

### 8.1 跳转矩阵

#### 概览卡片跳转

所有概览卡片跳转必须携带：`fromDashboard=true`、`dashboardSection=overview`、`dashboardReturnTo=/dashboard`

| 概览卡片 | 跳转目标 |
| --- | --- |
| `module_count` | `Module Registry / List` |
| `product_count` | `Product Registry / List` |
| `repository_count` | `Repository Binding / List` |
| `decision_count` | `Decision Center / List` |

#### 反馈信号卡片跳转

反馈信号卡片跳转必须携带：`fromDashboard=true`、`dashboardReturnTo=/dashboard`

| 信号类型 | 跳转目标 | `dashboardSection` |
| --- | --- | --- |
| 单项决策信号（绑定 `decision_id`） | `Decision Detail` | `current-focus` 或 `asset-feedback` |
| 聚合决策信号（未绑定单一 `decision_id`） | `Decision Center / List` | `current-focus` 或 `asset-feedback` |
| `product missing both bindings` | `Product Detail`（以 `product_id` 为目标身份） | `current-focus` 或 `asset-feedback` |
| `product missing repository binding` | `Product Detail`（以 `product_id` 为目标身份） | `current-focus` 或 `asset-feedback` |
| `product missing module binding` | `Product Detail`（以 `product_id` 为目标身份） | `current-focus` 或 `asset-feedback` |

> 约束：若卡片位于 `Current Focus / Next Action`，则 `dashboardSection=current-focus`；若卡片位于 `Asset Feedback`，则 `dashboardSection=asset-feedback`。不得伪造不存在的 `Decision Detail`。

#### 最近活动项跳转

活动项跳转必须携带：`fromDashboard=true`、`dashboardSection=recent-activity`、`dashboardReturnTo=/dashboard`

| 活动类型 | 跳转目标 |
| --- | --- |
| `module` | `Module Detail` |
| `release` | 所属 `Module Detail` |
| `product` | `Product Detail` |
| `repository` | `Repository Binding Detail / Workspace` |
| `decision` | `Decision Detail` |
| `product_module_binding` | `Product Detail` |
| `product_repository_binding` | `Repository Binding Detail / Workspace` |
| `module_repository_binding` | `Repository Binding Detail / Workspace` |

### 8.2 来源参数语义

| 参数 | 语义 | 允许取值 |
| --- | --- | --- |
| `fromDashboard` | Dashboard 外层来源标记 | `true` |
| `dashboardSection` | 来源区块标记 | `overview` / `current-focus` / `asset-feedback` / `recent-activity` / `empty-state` |
| `dashboardReturnTo` | 返回 Dashboard 路径 | `/dashboard` |

> 约束：
> - `fromDashboard` 是 Dashboard 外层来源标记，不覆盖 `phase04` 已冻结的页面原生来源模型（如 `fromList`）
> - 不得发明 `dashboardFrom`、`returnToDashboard`、`fromDashboardList` 等第二套参数命名
> - Dashboard 自身路由 `/dashboard` 当前阶段不承接搜索参数

### 8.3 返回路径规则

#### 直接跳转返回

- 从四类 Detail（Module / Product / Repository / Decision）返回 → `/dashboard`
- 从四类 List（Module / Product / Repository / Decision）返回 → `/dashboard`

#### 多跳返回（Dashboard → List → Detail）

- Detail 页必须同时保留 `fromList`（原生列表来源）与 `fromDashboard=true`（外层来源）
- 从 Detail 返回 List 必须使用 `fromList` 上下文
- 从 List 返回 Dashboard 必须使用 `fromDashboard` 上下文
- 不得把两者混写成同一个主来源字段

#### Create 页回流

- Create 页必须携带 `fromDashboard=true`、`dashboardSection=empty-state`、`dashboardReturnTo=/dashboard`
- 取消时必须返回 `/dashboard`，而不是回列表
- 提交成功后进入 Detail 页必须继续保留 `fromDashboard=true`
- 从 Detail 页返回必须能回到 `/dashboard`

#### 主动返回 Dashboard

- 携带 `fromDashboard=true` 的页面主动触发"返回 Dashboard"时，必须返回 `/dashboard`
- 主动返回导航通过一次性路由导航状态承接同名 `dashboardSection` 值，作为落地 `/dashboard` 后恢复来源区块的唯一临时承接机制
- 该一次性路由导航状态只服务本次主动返回落地，不得提升为 `/dashboard` 的搜索参数、持久化层或新的第二套来源参数命名

#### 写入成功后的 reread

- 用户从 Dashboard 进入 canonical owner 页后发起既有写入时，写入成功后必须继续停留在 canonical owner 页面完成 reread
- 不得因为页面携带 `fromDashboard=true` 就自动跳回 Dashboard
- 只有用户后续主动选择返回 Dashboard 时，才允许使用 `dashboardReturnTo=/dashboard`

### 8.4 刷新恢复与非法参数回退

- 刷新 `/dashboard` 时，必须重新执行三类聚合读，基于最新读结果重新派生整页状态
- 刷新携带 `fromDashboard=true` 的目标页时，必须继续保留来源参数，以当前 URL 搜索参数为唯一事实源
- 不得依赖 `sessionStorage / localStorage / 全局瞬时 store` 恢复上一次 Dashboard 页面状态
- `dashboardSection` 非法时回退为 `overview`
- `dashboardReturnTo` 缺失或非法时回退为 `/dashboard`
- 不得静默跳到根路径 `/`

---

## 9. 前端页面与状态模型

### 9.1 页面文件落点

| 文件 | 落点 |
| --- | --- |
| `DashboardHomePage` | `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` |
| `DashboardRoute` | `frontend/src/routes/dashboard.tsx`（URL: `/dashboard`） |
| Dashboard 组件 | `frontend/src/features/dashboard/components/` |

> 约束：当前阶段不得并行引入 `Dashboard Detail`、独立 `Feedback` 页面或第二套 Dashboard 子页面体系；不得把 `/dashboard` 提前拆成 `overview / activity / feedback` 等子路由。

### 9.2 主导航接入

- 必须在 `frontend/src/routes/__root.tsx` 的一级导航中新增 `Dashboard` 入口
- 既有 `Modules / Decisions / Products / Repositories` 一级导航继续保留
- 不得把 `Dashboard` 做成隐藏入口、二级入口或替换根布局宿主语义

### 9.3 组件树

```
DashboardHomePageShell
├── CurrentFocusSection
│   └── FeedbackSignalCardList
│       └── FeedbackSignalCard
├── DashboardOverviewSection
│   └── DashboardOverviewCardGrid
│       └── DashboardOverviewCard
├── AssetFeedbackSection
│   └── AssetFeedbackList
│       └── AssetFeedbackItemCard
├── RecentActivitySection
│   └── RecentActivityList
│       └── RecentActivityItemCard
└── DashboardPrimaryActionPanel
```

> 约束：组件树顺序与第一屏优先级（`Current Focus / Next Action` 优先）保持一致。只有存在跨 `dashboard` 与既有 canonical 页面复用证据时，才允许把组件抽为共享组件。

### 9.4 整页查询状态与整页视图状态

#### 最小查询状态

- `overviewQueryState = pending / success / error`
- `feedbackQueryState = idle / pending / success / error`
- `activityQueryState = idle / pending / success / error`

#### 派生整页视图状态

- `initial-loading`：只允许出现在首轮进入且 `DashboardOverviewRead` 尚未成功前
- `ready`：一旦 `DashboardOverviewRead` 成功，页面必须进入 `ready`，即使附属聚合局部失败也不得回退
- `page-error`：只允许由 `DashboardOverviewRead` 失败触发

### 9.5 分区状态模型

| 区块 | 最小状态 | 说明 |
| --- | --- | --- |
| `DashboardOverviewSection` | `loading / ready / error` | 不单独引入 `empty`；零计数仍为 `ready` |
| `CurrentFocusSection` | `loading / ready / empty / error` | 与 `AssetFeedbackSection` 共享 `feedbackQueryState` |
| `AssetFeedbackSection` | `loading / ready / empty / error` | 与 `CurrentFocusSection` 共享 `feedbackQueryState` |
| `RecentActivitySection` | `loading / ready / empty / error` | 依赖 `activityQueryState` |
| `DashboardPrimaryActionPanel` | `computing / ready / hidden / suppressed` | 见 §9.6 |

### 9.6 主 CTA 状态机

| 状态 | 语义 |
| --- | --- |
| `computing` | `DashboardOverviewRead` 仍未成功，不得猜测任何主 CTA |
| `ready` | 当前存在唯一主 CTA |
| `hidden` | 成功中性态，不展示强制主 CTA |
| `suppressed` | 附属聚合局部失败，禁止猜测或伪造反馈型主 CTA |

#### 命中规则

- `DashboardOverviewRead` 成功且命中 CTA 1-4 → 直接进入 `ready`，不需要等待附属聚合
- `DashboardOverviewRead` 成功且未命中 CTA 1-4，`FeedbackSignalRead` 尚未成功 → 保持 `computing`
- `DashboardOverviewRead` 成功且未命中 CTA 1-4，`FeedbackSignalRead = error` → 进入 `suppressed`
- `DashboardOverviewRead` 成功，未命中 CTA 1-4，`FeedbackSignalRead` 成功且不存在反馈缺口 → 进入 `hidden`

### 9.7 局部重试

- `FeedbackSignalRead` 失败时，`CurrentFocusSection` 与 `AssetFeedbackSection` 都必须在各自区块内容区域内呈现局部错误
- 任一区块触发重试时，必须重新执行同一次 `FeedbackSignalRead`
- `RecentActivityRead` 失败时，错误反馈必须停留在活动区块内容区域内，区块重试只允许重新执行 `RecentActivityRead`
- 不得把局部重试升级为整页刷新

### 9.8 PC 与移动浏览器布局降级

#### PC 桌面环境

- 页面必须优先保证 `Current Focus / Next Action` 的第一行动优先级
- `dashboard_overview` 必须以概览卡片网格形式存在
- `Asset Feedback` 与 `Recent Activity` 作为补充区块与主行动区并列承接
- 允许同屏呈现四区块，但不得牺牲 `Current Focus` 的主行动优先级

#### 移动浏览器窄屏环境

- 必须沿用与桌面端同一套区块语义与动作体系
- 区块顺序必须优先保持 `Current Focus / dashboard_overview / Asset Feedback / Recent Activity`
- 必须通过单列垂直重排、信息裁剪与列表收束降低拥挤度
- 不得引入第二套独立移动端页面或第二套路由树

### 9.9 不冻结的运行时实现细节

以下内容不在当前阶段冻结：具体 hook 命名、Query key、store API 命名、请求取消、缓存时间、局部重试策略、Zustand 容器拆分、预取策略、optimistic update 方案。

---

## 10. 后端模块边界与接口分组

### 10.1 模块物理边界

- Dashboard 后端逻辑独立落点为 `backend/internal/dashboard/`
- Go package：`github.com/psco/backend/internal/dashboard`
- 不得把 Dashboard 聚合读取逻辑塞入 `moduleregistry / decisioncenter / productregistry / repositorybinding` 任一既有模块
- 不得把 Dashboard 聚合读取逻辑塞入 `platform` 路由装配层

### 10.2 模块内部文件分层

| 文件/目录 | 职责 |
| --- | --- |
| `handler/` | 承接 HTTP 入口与参数解析 |
| `service/` | 承接聚合读取编排与 `Feedback Signal Card` 归一化 |
| `repository/` | 承接 Dashboard 自有的派生查询（如有） |
| `candidate/` | 承接跨模块读取接口定义与 owner 归属 |
| `types.go` | 承接 DTO 与领域类型 |
| `errors.go` | 承接业务错误哨兵值 |
| `validate.go` | 承接入参校验 |
| `handler/response.go` | 承接统一 JSON 响应 |

> 约束：`errors.go / types.go / validate.go / handler/response.go` 必须按职责单值化映射到唯一文件，禁止散落在 handler 或 service 内部。

### 10.3 接口分组

Dashboard 模块的接口分组冻结为只读三组：

| 读组 | service 方法 | 职责 |
| --- | --- | --- |
| `DashboardOverviewRead` | `QueryService.ReadOverview` | 编排六个计数的聚合 |
| `FeedbackSignalRead` | `QueryService.ReadFeedbackSignal` | 编排 `pending_decision_signals` 与 `product_asset_coverage` 的读取、归一化、排序与裁剪 |
| `RecentActivityRead` | `QueryService.ReadRecentActivity` | 编排多个 canonical 模块的活动项读取、类型映射、归并与按 `activity_at` 排序 |

> 约束：
> - 三个读取分别对应 `handler/` 层三个独立 HTTP endpoint
> - Dashboard 模块当前阶段只承接读组，不得新增 `handler/command_*.go` 或 `service/command_*.go` 等写组文件
> - 不得并列增加趋势分析、外部埋点、通知中心或导出接口

### 10.4 QueryService owner

- 单一 `QueryService` 落点为 `backend/internal/dashboard/service/query_service.go`
- 沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的单文件 `service/query_service.go` 统一承接读组编排模式
- 不得为任一读组单独发明独立 service owner
- 不得把读组逻辑下沉到 `handler/` 或 `repository/` 层

### 10.5 跨模块读取边界

Dashboard 与四个 canonical 模块的服务侧连接边界冻结为两类已验证受控读模式：

1. **Dashboard candidate 自拥有读**：Dashboard `candidate/` 子包自己定义并拥有接口与实现，用于 Dashboard 自己拥有的原始聚合读取。沿用 `phase03` 已验证的 `DecisionModuleCandidateRead` 模式。
2. **provider-owned read**：canonical owner 提供已拥有的源语义或摘要读取，由 Dashboard `QueryService` 通过接口注入消费。沿用 `phase04` 已验证的注入复用模式。

> 约束：
> - 不得在这两类模式之外再发明第三套跨模块读取架构
> - Dashboard `service/` 层不得直接 import canonical 模块的 `repository/` 包
> - Dashboard `service/` 层不得直接跨模块写 SQL
> - 具体采用哪一种模式，由读语义 owner 决定，而不是按模块名一刀切

#### Dashboard candidate reader 分工（仅适用于 Dashboard 自拥有原始聚合读）

| reader | 返回内容 |
| --- | --- |
| `ModuleCountReader` | `module_count` |
| `ModuleActivityReader` | `Module` 与 `Release` 的最近活动 |
| `DecisionCountReader` | `decision_count` |
| `PendingDecisionSignalReader` | 原始待决策信号 |
| `DecisionActivityReader` | `Decision` 的最近活动 |
| `ProductCountReader` | `product_count / product_with_module_count / product_with_repository_count` |
| `ProductAssetCoverageReader` | 原始资产缺口数据 |
| `ProductActivityReader` | `Product` 与 `product_module_binding` 的最近活动 |
| `RepositoryCountReader` | `repository_count` |
| `RepositoryActivityReader` | `Repository` 与 `product_repository_binding / module_repository_binding` 的最近活动 |

> 约束：
> - 上表只描述“Dashboard 自己拥有的原始聚合读”在采用 `candidate-owned` 模式时的最小 reader 分工，不得被误读为“所有 Dashboard 读取都必须 candidate-owned”
> - reader 接口的定义与实现均由 Dashboard 模块 `candidate/` 子包自己拥有。canonical 模块不需要为 Dashboard 新增 candidate 实现
> - Dashboard `candidate/` 实现可以直接读取 canonical 模块的表，但必须在 `candidate/` 子包内隔离
> - reader 接口的具体方法签名由 `phase05-08` 承接

#### provider-owned read 的正式保留口径

- 当某个读取属于 canonical owner 已拥有的源语义或摘要读取时，正式正文必须继续允许直接采用 `provider-owned read`
- `provider-owned read` 一旦可复用，Dashboard 不得为了“形式统一”再复制一份等价的 `candidate` reader 实现
- `Dashboard candidate reader` 分工表不得被当作 `provider-owned read` 的替代表或强制覆盖表
- `platform` 装配层可以同时承接：
  - Dashboard `candidate/` 子包自拥有的 reader 实现
  - canonical owner 已拥有的 provider-owned read 实现
- 具体采用哪一种模式，继续由读语义 owner 决定，而不是按模块名、目录结构或“是否已经列进 candidate 表”一刀切

### 10.6 归一化与类型映射 owner

| 归一化/映射 | owner |
| --- | --- |
| `Feedback Signal Card` 归一化组装 | Dashboard 模块 `QueryService.ReadFeedbackSignal` 内部 |
| `recent_activity` 类型映射 | Dashboard 模块 `QueryService.ReadRecentActivity` 内部 |
| DTO 定义 | `backend/internal/dashboard/types.go` |
| DTO 与 `.proto` 映射函数 | Dashboard 模块（显式映射函数，不依赖反射或隐式序列化） |

> 约束：不得让各 canonical 模块自己产 `Feedback Signal Card` 或统一活动流结构。不得把归一化/映射逻辑分散到 `handler/` 或 `repository/` 层。

### 10.7 依赖注入装配点

- 装配必须发生在 `backend/internal/platform/` 路由装配层
- `platform` 层负责构造 Dashboard `candidate/` 子包的 reader 实现，或装配 canonical owner 提供的 provider-owned read 实现
- `platform` 层负责把这些读依赖注入到 Dashboard `QueryService` 构造函数
- 沿用 `phase03 / phase04` 已落地的 platform 装配模式
- 不得让 Dashboard 模块在内部自行 new 出跨模块读依赖

### 10.8 错误响应归属

- 错误响应必须通过 `backend/internal/dashboard/handler/response.go` 统一返回
- 沿用 `phase02-09 / phase03-10 / phase04-10` 已落地的 JSON 错误响应格式
- `DashboardOverviewRead` 失败 → handler 返回整页失败响应，不得返回部分计数或伪装成功
- `FeedbackSignalRead` 或 `RecentActivityRead` 失败 → 各自 handler 继续返回既有统一错误响应，不得把附属读失败升级为整页失败
- 空态响应不映射为错误（不得返回 `404` 或 `500`）
- 当前阶段不得在附属读 endpoint 的响应包络里额外发明"局部失败标记"

### 10.9 不冻结的实现细节

以下内容不在当前阶段冻结：Go 数据访问层具体工具（`sqlx / sqlc / GORM / database/sql`）、缓存层、连接池配置、查询超时。

---

## 11. API 边界与合同矩阵

### 11.1 合同源

- `.proto` 是 `Dashboard + Feedback` 的唯一合同源
- 合同文件落点：`proto/psco/dashboard/v1/dashboard.proto`
- 包名与版本语义：`psco.dashboard.v1`
- Go 生成包路径：`backend/internal/gen/proto/psco/dashboard/v1/`
- TypeScript 生成产物路径：`frontend/src/gen/proto/psco/dashboard/v1/`
- 继续复用现有 `proto/buf.yaml`、`proto/buf.gen.yaml` 与 `proto/Makefile`
- 不得在 `backend/`、`frontend/` 或仓库根新增并列 proto 工作区

### 11.2 DashboardService RPC 矩阵

| RPC | 承接读组 | HTTP 映射 |
| --- | --- | --- |
| `GetDashboardOverview` | `DashboardOverviewRead` | `GET /api/dashboard/overview` |
| `GetFeedbackSignals` | `FeedbackSignalRead` | `GET /api/dashboard/feedback-signals` |
| `GetRecentActivities` | `RecentActivityRead` | `GET /api/dashboard/recent-activities` |

> 约束：
> - 三个 RPC 都只承接读，不承接任何业务写入
> - 当前阶段三个入口都不承接 body、分页参数或排序切换参数
> - request 消息保持无筛选、无分页、无排序切换的最小请求边界
> - 不得把 HTTP 路径、状态码或中间件策略误写成 `.proto` 合同本体
> - JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 显式对齐

### 11.3 消息结构与字段编号

#### DashboardOverview

| 字段 | 编号 | 类型 |
| --- | --- | --- |
| `module_count` | 1 | `int32` |
| `product_count` | 2 | `int32` |
| `repository_count` | 3 | `int32` |
| `decision_count` | 4 | `int32` |
| `product_with_repository_count` | 5 | `int32` |
| `product_with_module_count` | 6 | `int32` |

- `GetDashboardOverviewRequest`：空消息（当前版本无字段）
- `GetDashboardOverviewResponse`：`overview=1(DashboardOverview)`

#### FeedbackSignal

| 字段 | 编号 | 类型 |
| --- | --- | --- |
| `signal_family` | 1 | `FeedbackSignalFamily` |
| `signal_code` | 2 | `FeedbackSignalCode` |
| `priority` | 3 | `FeedbackSignalPriority` |
| `title` | 4 | `string` |
| `summary` | 5 | `string` |
| `action_label` | 6 | `string` |
| `target_type` | 7 | `DashboardTargetType` |
| `target_id` | 8 | `string` |
| `target_label` | 9 | `string` |

#### ProductAssetCoverageSummary

| 字段 | 编号 | 类型 |
| --- | --- | --- |
| `fully_bound_product_count` | 1 | `int32` |
| `missing_both_bindings_count` | 2 | `int32` |
| `missing_repository_binding_count` | 3 | `int32` |
| `missing_module_binding_count` | 4 | `int32` |
| `representative_signals` | 5 | `repeated FeedbackSignal` |

- `GetFeedbackSignalsRequest`：空消息（当前版本无字段）
- `GetFeedbackSignalsResponse`：`current_focus_signals=1(repeated FeedbackSignal)` / `asset_feedback_summary=2(ProductAssetCoverageSummary)`
- `current_focus_signals` 最多展示 `5` 条主队列信号
- `representative_signals` 最多展示 `3` 条代表性缺口项

#### RecentActivityItem

| 字段 | 编号 | 类型 |
| --- | --- | --- |
| `activity_type` | 1 | `RecentActivityType` |
| `activity_at` | 2 | `google.protobuf.Timestamp` |
| `target_type` | 3 | `DashboardTargetType` |
| `target_id` | 4 | `string` |
| `target_label` | 5 | `string` |

- `GetRecentActivitiesRequest`：空消息（当前版本无字段）
- `GetRecentActivitiesResponse`：`activities=1(repeated RecentActivityItem)`
- 当前阶段最多返回 `10` 条活动项

### 11.4 枚举编号方案

#### FeedbackSignalFamily

| 枚举值 | 编号 |
| --- | --- |
| `FEEDBACK_SIGNAL_FAMILY_UNSPECIFIED` | 0 |
| `FEEDBACK_SIGNAL_FAMILY_PENDING_DECISION` | 1 |
| `FEEDBACK_SIGNAL_FAMILY_PRODUCT_ASSET_COVERAGE` | 2 |

#### FeedbackSignalCode

| 枚举值 | 编号 |
| --- | --- |
| `FEEDBACK_SIGNAL_CODE_UNSPECIFIED` | 0 |
| `FEEDBACK_SIGNAL_CODE_PENDING_DECISION` | 1 |
| `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS` | 2 |
| `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING` | 3 |
| `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING` | 4 |

#### FeedbackSignalPriority

| 枚举值 | 编号 |
| --- | --- |
| `FEEDBACK_SIGNAL_PRIORITY_UNSPECIFIED` | 0 |
| `FEEDBACK_SIGNAL_PRIORITY_P1_PENDING_DECISION` | 1 |
| `FEEDBACK_SIGNAL_PRIORITY_P2_PRODUCT_MISSING_BOTH_BINDINGS` | 2 |
| `FEEDBACK_SIGNAL_PRIORITY_P3_PRODUCT_MISSING_REPOSITORY_BINDING` | 3 |
| `FEEDBACK_SIGNAL_PRIORITY_P4_PRODUCT_MISSING_MODULE_BINDING` | 4 |

#### DashboardTargetType

| 枚举值 | 编号 |
| --- | --- |
| `DASHBOARD_TARGET_TYPE_UNSPECIFIED` | 0 |
| `DASHBOARD_TARGET_TYPE_DECISION_DETAIL` | 1 |
| `DASHBOARD_TARGET_TYPE_DECISION_LIST` | 2 |
| `DASHBOARD_TARGET_TYPE_PRODUCT_DETAIL` | 3 |
| `DASHBOARD_TARGET_TYPE_MODULE_DETAIL` | 4 |
| `DASHBOARD_TARGET_TYPE_REPOSITORY_DETAIL` | 5 |

> 约束：当 `target_type = DASHBOARD_TARGET_TYPE_DECISION_LIST` 时，`target_id` 必须允许为空字符串，以表达聚合决策信号落到 `Decision Center / List` 的语义。

#### RecentActivityType

| 枚举值 | 编号 |
| --- | --- |
| `RECENT_ACTIVITY_TYPE_UNSPECIFIED` | 0 |
| `RECENT_ACTIVITY_TYPE_MODULE` | 1 |
| `RECENT_ACTIVITY_TYPE_RELEASE` | 2 |
| `RECENT_ACTIVITY_TYPE_PRODUCT` | 3 |
| `RECENT_ACTIVITY_TYPE_REPOSITORY` | 4 |
| `RECENT_ACTIVITY_TYPE_DECISION` | 5 |
| `RECENT_ACTIVITY_TYPE_PRODUCT_MODULE_BINDING` | 6 |
| `RECENT_ACTIVITY_TYPE_PRODUCT_REPOSITORY_BINDING` | 7 |
| `RECENT_ACTIVITY_TYPE_MODULE_REPOSITORY_BINDING` | 8 |

### 11.5 合同演进规则

- 新增字段必须使用新的递增编号，不得复用已删除字段编号或字段语义
- 每个业务枚举都必须包含首个 `UNSPECIFIED = 0`
- 不得通过修改既有字段类型、重复性或语义来做兼容性演进
- 删除字段、删除枚举值或废弃字段名时，必须使用 `reserved` 保留字段号，必要时同时保留字段名或枚举名
- 不得删除后再回收复用原编号
- `buf breaking` 必须直接对照 `../.git#branch=main,subdir=proto`，不得吞掉 breaking 失败退出码
- 最小校验链必须覆盖 `buf build`、`buf lint`、`buf generate`、`buf breaking`

### 11.6 DTO 字段语义对齐

| DTO | 最小字段集 |
| --- | --- |
| `Feedback Signal Card` DTO | `signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label` |
| `Recent Activity Item` DTO | `activity_type / activity_at / target_type / target_id / target_label` |
| `Dashboard Overview` DTO | `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count` |

> 约束：不得减少上述字段；不得在当前阶段额外引入 `score / trend / external_metric / recommendation_reason`。

---

## 12. 冷启动、fixture 与验收基线

### 12.1 验收环境建立方式

- Dashboard 验收环境必须复用既有 `database/scripts/` + `database/seeds/` 模式，不发明第二套工具链
- `reset_dashboard_acceptance.sh` 是 Dashboard 验收的唯一统一入口，落点为 `database/scripts/`
- 必须复用现有 `podman exec` / `docker exec` / 宿主机 `psql` 自动检测模式执行 SQL
- 不得在 `backend/`、`frontend/` 或仓库根新增并列验收工具链

### 12.2 reset_dashboard_acceptance.sh 模式矩阵

| 模式 | 语义 |
| --- | --- |
| 默认（无参数） | 先清空所有 Dashboard 相关数据，再恢复完整基线 |
| `--clean-only` | 仅清空，用于验证空系统状态 |
| `--restore-only` | 仅恢复完整基线，用于验证有数据状态 |
| `--fixture <name>` | 先清空，再加载指定 fixture |

> 约束：
> - 所有 `--fixture` 模式统一遵守"先清空，再加载指定 fixture"语义，不允许任何 fixture 叠加已有数据
> - `--fixture` 的 `<name>` 只允许取本规格冻结的九类 fixture 名称之一
> - 不得并列发明 `--reset-all`、`--load-fixture`、`--setup` 等第二套参数命名

### 12.3 清空与恢复范围

#### 清空范围

通过编排既有 `reset_product_repository_mainline.sh --clean-only`、`reset_decision_mainline.sh --clean-only` 与 `reset_module_mainline.sh --clean-only` 实现，按依赖逆序执行：

1. 先清 `product_repository`（依赖 modules）
2. 再清 `decision`（decision_links 依赖 modules）
3. 最后清 `module`（modules 依赖 readonly prereqs 中的 products / decisions）

覆盖表：`decision_links` / `product_repositories` / `product_modules` / `module_repositories` / `module_releases` / `modules` / `products` / `repositories` / `decisions`

> 不得绕过既有脚本直接 `DELETE FROM` 或 `TRUNCATE` 底层表；不得清空 schema 或 migration 元数据表。

#### 恢复范围

按依赖顺序执行：

1. `seed_readonly_prereqs.sql`（提供 products / repositories / decisions 只读前提）
2. `reset_module_mainline.sh --restore-only`（恢复 modules / module_releases / product_modules / module_repositories / decision_links）
3. `reset_decision_mainline.sh --restore-only`（恢复 decisions 正式基线与 decision_links，依赖 modules）
4. `reset_product_repository_mainline.sh --restore-only`（恢复 products / repositories 正式基线，依赖 modules）
5. `seed_dashboard_acceptance_baseline.sql`（补齐 Dashboard 验收所需的额外基线数据）

### 12.4 九类 fixture 矩阵

| 序号 | fixture 名称 | 类别 | 数据状态摘要 | 映射 CTA / 区块结果 |
|------|-------------|------|-------------|-------------------|
| 1 | `empty-system` | 最小 | 所有表为空 | CTA 1：冷启动空系统 → Module Registry / Create |
| 2 | `modules-only` | 最小 | 仅有 modules + module_releases | CTA 3：module_count > 0 && product_count = 0 → Product Registry / Create |
| 3 | `products-without-modules` | 最小 | 仅有 products | CTA 2：非空缺口 → Module Registry / Create |
| 4 | `products-missing-repository` | 最小 | 目标 product 有 module 无 repository | CTA 7：product missing repository binding → Product Detail |
| 5 | `products-missing-module` | 最小 | 目标 product 有 repository 无 module | CTA 8：product missing module binding → Product Detail |
| 6 | `pending-decisions` | 最小 | 完整基线 + pending 状态 decisions | CTA 5：pending_decision_signals → Decision Detail / List |
| 7 | `recent-activities` | 最小 | 完整基线 + 所有 product 双绑定 + 多类活动项 | Recent Activity 区块 + CTA 9：系统已就绪中性状态 |
| 8 | `products-missing-all-repositories` | CTA 扩展 | modules + products + product_modules 存在，但 repositories 表为空 | CTA 4：module_count > 0 && product_count > 0 && repository_count = 0 → Repository Binding / Create |
| 9 | `products-missing-both-bindings` | CTA 扩展 | 目标 product 同时无 product_modules 与 product_repositories | CTA 6：product missing both bindings → Product Detail |

> 约束：
> - 每类 fixture 必须有对应的 `seed_dashboard_fixture_<name>.sql` 文件，落点在 `database/seeds/`
> - 不得在 `database/seeds/` 下新建 `dashboard/` 或 `fixtures/` 等子目录
> - CTA 1-9 全部有正式 `--fixture` 入口覆盖，不依赖变体或手工 SQL
> - 不得在验收过程中临时执行手工 `INSERT`、`DELETE`、`UPDATE` 语句

#### CTA 6 / 7 / 8 / 9 的单值命中前提

- `products-missing-repository` 必须额外满足：
  - 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
  - `FeedbackSignalRead` 必须产出 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING`
  - `FeedbackSignalRead` 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION`
- `products-missing-module` 必须额外满足：
  - 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
  - `FeedbackSignalRead` 必须产出 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING`
  - `FeedbackSignalRead` 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION`
- `products-missing-both-bindings` 必须额外满足：
  - 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
  - `FeedbackSignalRead` 必须产出 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS`
  - `FeedbackSignalRead` 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION`
- `recent-activities` 必须额外满足：
  - 所有 product 必须已完成模块与仓库双绑定，使 Dashboard 不命中任何缺口 CTA
  - `decision` 类型活动只能来自非 pending decision
  - 当前 fixture 不得存在进入 `pending_decision_signals` 主队列的待决策信号
  - `FeedbackSignalRead` 不得产出 `FEEDBACK_SIGNAL_CODE_PENDING_DECISION`、`FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS`、`FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING` 或 `FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING`
- 上述四类 fixture 的目的都是保证 CTA 6 / 7 / 8 / 9 为单值命中结果，不允许被更高优先级 CTA 5 或其他缺口 CTA 抢占

### 12.5 fixture SQL 幂等约束

- 每个 fixture 文件必须可重复执行不报错
- 必须使用 `ON CONFLICT DO NOTHING`、`WHERE NOT EXISTS` 或 `UPDATE-then-INSERT` 模式保证幂等
- 必须在文件头部以注释说明：用途、依赖、幂等语义、使用方式
- 必须沿用既有 `seed_module_mainline_baseline.sql` 与 `seed_product_repository_mainline_baseline.sql` 的注释结构

### 12.6 局部错误模拟

局部错误状态的模拟方式冻结为受控环境变量单一入口，不得依赖破坏数据库结构、临时手工操作或 SQL fixture 文件。

| 环境变量 | 触发效果 |
| --- | --- |
| `DASHBOARD_SIMULATE_OVERVIEW_ERROR=true` | overview reader 返回错误 → 整页失败 |
| `DASHBOARD_SIMULATE_FEEDBACK_ERROR=true` | feedback reader 返回错误 → 局部失败，overview 与 recent activity 仍成功 |
| `DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR=true` | recent activity reader 返回错误 → 局部失败，overview 与 feedback 仍成功 |

> 约束：
> - 局部错误模拟不产生 `seed_dashboard_fixture_local_errors.sql` 文件
> - 该机制必须可重复启用与关闭
> - 不得通过删除表、重命名列或破坏 schema 来模拟错误
> - 局部错误模拟的详细实现由 `phase05-12` 承接，本规格只冻结模拟方式必须受控且可重复

### 12.7 跳转返回链路验收矩阵

| 场景 | 覆盖路径 |
| --- | --- |
| 直接跳转 Detail 返回 | Dashboard → Module/Product/Repository/Decision Detail → 返回 `/dashboard` |
| 直接跳转 List 返回 | Dashboard → Module/Product/Repository/Decision List → 返回 `/dashboard` |
| 多跳路径 | Dashboard → List → Detail → List → Dashboard（`fromList` + `fromDashboard` 共存） |
| Create 回流 | Dashboard → Module/Product/Repository Create → 取消返回 `/dashboard`；提交后 Detail → 返回 `/dashboard` |
| 刷新恢复 | 刷新后 `fromDashboard / dashboardSection / dashboardReturnTo` 参数不丢失 |
| 非法参数回退 | `dashboardSection` 回退 `overview`；`dashboardReturnTo` 回退 `/dashboard` |

---

## 13. 非目标

`phase05` 明确不做：

- 第二套资产注册或绑定主线
- 第二套写入主线（Dashboard 不承接业务写入）
- 复杂驾驶舱 / BI 分析页
- 外部遥测、通知中心与自动消息回流
- 独立 `Feedback` 重实体 CRUD
- Capability Growth 完整评分系统
- `Feature / Opportunity / Experiment`
- GitHub OAuth / 自动导入
- 自动扫描代码
- AI 自动建议或自动补全
- 独立 `AI Assistant` 一级导航
- 独立 `React Native` 客户端
- 完整 `PWA`

> 约束：不得把上述长期方向误写成当前阶段既成事实。

---

## 14. Done 标准

当以下条件成立时，`phase05` 规格成立并可进入稳定实现：

1. **页面、区块与路由边界完整**：`Dashboard Home` 页面职责单值化，四区块固定归属，路由 `/dashboard` 与主导航接入明确
2. **三类聚合读与反馈信号语义完整**：`DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 边界明确，反馈优先级与 CTA 矩阵单值化
3. **跳转、返回与刷新恢复语义完整**：跳转矩阵、来源参数、返回路径与刷新恢复规则单值化
4. **前端页面与状态模型完整**：页面文件落点、组件树、整页状态、分区状态与主 CTA 状态机明确
5. **后端模块边界与接口分组完整**：模块物理边界、文件分层、QueryService owner 与跨模块读取边界明确
6. **`.proto` 合同边界完整**：合同源、RPC 矩阵、消息结构、字段编号、枚举与演进规则明确
7. **验收环境、fixture 与局部错误基线完整**：`reset_dashboard_acceptance.sh`、九类 fixture 与局部错误环境变量入口明确
8. **非目标与下一阶段引用前提完整**：非目标矩阵与 Done 标准足以支撑后续 `phase05-11 ~ 14` 的实现、联调、验收与收口

---

## 15. 与根级真相源互链

- 项目定位与入口：以 `AGENTS.md` 为准
- 阶段路线：以 `plan.md` 为准
- 技术栈标准：以 `TECH_STACK_BASELINE.md` 为准
- 规则门禁：以 `project_rules.md` 为准
- 目录与迁移落点：以 `architecture_map.md` 为准
- 最终共识：以 `PSCO-summarize-feedback.md` 为准

> 本文档不重写上述根级真相源的主结论，仅作为 `Dashboard + Feedback` 执行层面的唯一规格收敛入口。

---

## 16. 追溯来源矩阵

本文档的每条冻结结论均可追溯到以下子规格：

| 章节 | 主要追溯来源 |
| --- | --- |
| §2 对象范围 | `phase05-01` / `phase05-02` / `phase05-04` |
| §3 页面矩阵 | `phase05-01` / `phase05-03` / `phase05-05` |
| §4 区块矩阵 | `phase05-01` / `phase05-02` / `phase05-05` |
| §5 动作矩阵 | `phase05-04` / `phase05-07` |
| §6 数据模型 | `phase05-02` / `phase05-04` |
| §7 聚合读与反馈信号 | `phase05-02` / `phase05-04` / `phase05-06` |
| §8 跳转与返回上下文 | `phase05-03` / `phase05-06` |
| §9 前端页面与状态模型 | `phase05-05` / `phase05-06` |
| §10 后端模块边界与接口分组 | `phase05-07` |
| §11 API 边界与合同矩阵 | `phase05-08` |
| §12 冷启动、fixture 与验收基线 | `phase05-09` |
| §13 非目标 | `phase05` shared_baseline §8 |
| §14 Done 标准 | `phase05-10` spec.md |

> 约束：`phase05-01 ~ 09` 在本文档生效后退为追溯来源与证据链，不再承担并列直接执行层入口职责。后续实现、实现复核、验收与下一阶段文档应默认先引用本文档。
