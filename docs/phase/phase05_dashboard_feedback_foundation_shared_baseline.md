# phase05_dashboard_feedback_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase05` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase05_dashboard_feedback_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase04` 已完成收口，项目当前入口已切换为 `phase05_dashboard_feedback_foundation`
- 当前阶段下一步：先完成 `phase05` 三件套审核，再进入对应 `/spec`

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
  - `.trae/specs/phase04_11_product_repository_binding_proto_mainline/`
  - `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md`
- 当前阶段只承接 `phase04` 已冻结并验收的资产、绑定与合同边界
- 当前阶段不反向重写 `phase04` 已冻结的 `Product Registry` 与 `Repository Binding` 边界
- `phase02` 与 `phase03` 的已交付语义继续经由 `phase04` 上游继承，不作为并列第二入口

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- Client State：`Zustand`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- Contract Tooling：`buf build / lint / generate / breaking`
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段不得重新引入 `Feature / Opportunity / Experiment`
- 当前阶段不得引入独立 `AI Assistant` 一级导航
- 当前阶段不得引入独立 `React Native` 客户端
- 当前阶段不得把完整 `PWA` 能力写成前置范围
- 当前阶段不得把 Dashboard 实现成无行动价值的统计页
- 当前阶段不得把 `Feedback` 扩写为独立重实体或完整运营系统
- 当前阶段不得让 Dashboard 并行拥有 `Product / Repository / Decision / Binding` 的第二套写入入口
- 当前阶段不得把 GitHub OAuth / 自动导入写成当前阶段阻断项

### 2.5 当前阶段交付模式

- `phase05` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 `Dashboard + Feedback` 最小主线代码

## 3. 当前阶段动作矩阵

`phase05` 最少需要直接承接：

- `ViewDashboard`
- `ReadDashboardOverview`
- `ReadPendingDecisionSignals`
- `ReadProductAssetCoverage`
- `ReadRecentActivity`
- `NavigateFromDashboardToOwner`

当前阶段必须打通的最小反馈闭环：

- `Decision -> pending_decision_signals -> Dashboard -> Decision canonical owner`
- `Product / Repository / Module Binding -> product_asset_coverage -> Dashboard -> Product/Repository canonical owner`
- `结构化资产新增或变化 -> recent_activity_feed -> Dashboard -> 对应详情页`

允许以最小连接位承接但不扩写为独立主线：

- `Release` 在 Dashboard 中的最近活动语义
- `Capability` 的派生视角
- `Venture` 的可选背景语义

## 4. 当前阶段页面矩阵

- `Dashboard Home`
- `Dashboard Home` 的正式业务入口路由冻结为 `/dashboard`
- `Dashboard` 在当前阶段作为既有主导航中的一级入口新增，不替代根级布局宿主本身
- 当前阶段不把 `/` 单值解释为 `Dashboard Home`；如未来需要把 `/` 重定向到 `/dashboard`，必须在后续 `phase / fix` 中单独冻结

允许存在最小跳转目标：

- `Module Registry / List / Detail`
- `Product Registry / List / Detail`
- `Repository Binding / List / Detail`
- `Decision Center / List / Detail`
- `Release` 当前阶段不新增独立 Detail 页面；Dashboard 中涉及 `Release` 的活动项统一跳转到所属 `Module Detail`

### 4.1 当前阶段交互归属矩阵

- `Dashboard Home`：承接概览读取、反馈信号读取、最近活动读取、空状态引导与跳转入口
- `Dashboard Home` 中的 `Next Action` 卡片：承接“下一步去哪里做”，不承接“在本页直接做”
- `Product Detail`：继续承接产品详情读取与 `BindModuleToProduct`
- `Repository Binding Detail / Workspace`：继续承接仓库详情读取与 `BindRepositoryToProduct / MapModuleToRepository`
- `Decision Detail / Decision Center`：继续承接决策详情读取与后续决策处理
- `Module Detail`：继续承接模块摘要与既有上下文入口

## 5. 当前阶段数据矩阵

直接承接：

- `modules`
- `releases`
- `products`
- `repositories`
- `decisions`
- `decision_links`
- `product_modules`
- `product_repositories`
- `module_repositories`

当前阶段必须正式消费的派生读取：

- `dashboard_overview`
- `pending_decision_signals`
- `product_asset_coverage`
- `recent_activity_feed`

### 5.1 最小读模型

- `dashboard_overview` 至少承接：
  - `module_count`
  - `product_count`
  - `repository_count`
  - `decision_count`
  - `product_with_repository_count`
  - `product_with_module_count`
- `pending_decision_signals` 至少承接：
  - 信号计数
  - 最多 `5` 条进入 `Current Focus / Next Action` 的高优先级项
  - 每条信号至少包含：`signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label`
- `product_asset_coverage` 至少承接：
  - 已完整绑定产品数
  - 缺少模块与仓库双绑定的产品数
  - 缺少仓库绑定但已有模块绑定的产品数
  - 缺少模块绑定但已有仓库绑定的产品数
  - 最多 `3` 条代表性缺口项
  - 每条缺口项至少包含：`signal_code / priority / target_product_id / target_product_name / action_label`
- `recent_activity_feed` 至少承接：
  - 最近新增或更新的 `Module / Release / Product / Repository / Decision / product_module_binding / product_repository_binding / module_repository_binding`
  - 最多 `10` 条活动项
  - 显式活动时间字段（如 `activity_at`）
  - 对应对象类型
  - 对应对象跳转信息

当前阶段关于 Dashboard 区块与排序的补充冻结如下：

- `Current Focus / Next Action` 只消费归一化后的反馈信号卡片
- `Asset Feedback` 只消费 `product_asset_coverage` 的补充摘要，不单独参与全局优先级排序
- `Recent Activity` 只按活动时间倒序排序，不参与反馈优先级竞争
- 同优先级反馈信号默认按“最近需要处理时间优先”；若上游读模型暂不提供该字段，则回退为 `created_at DESC`

当前阶段关于反馈信号优先级的单值顺序如下：

1. `pending_decision_signals`
2. `product missing both bindings`
3. `product missing repository binding`
4. `product missing module binding`

当前阶段 `Feedback` 最小语义冻结如下：

- `Feedback` 当前阶段不是独立表，也不是人工录入型对象
- `Feedback` 是基于既有结构化对象与绑定关系派生出的“行动信号”
- 信号最少回答两类问题：
  - 当前哪里不完整
  - 现在该去哪里补

### 5.2 最小接口归属前提

- `DashboardOverviewRead` 在 `phase05` 中按 `Dashboard Home` 的主聚合读取处理
- `FeedbackSignalRead` 在 `phase05` 中按 `Dashboard Home` 的附属聚合读取处理
- `RecentActivityRead` 在 `phase05` 中按 `Dashboard Home` 的附属聚合读取处理
- Dashboard 不直接承接任何新的业务写入接口
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层，但不得形成与 `.proto` 并列的第二套合同源

最小错误语义前提：

- 空系统时，Dashboard 返回明确空状态语义
- 某个反馈区块无结果时，返回空列表语义，不映射为资源不存在
- 某个聚合区块局部失败时，允许局部错误展示与局部重试，不强制整页失败
- 非法查询参数或不存在的跳转目标，必须返回明确失败语义，不得静默跳到错误页面

当前阶段关于活动项与反馈卡片的最小跳转矩阵冻结如下：

- `pending_decision_signals`：
  - 单项信号已绑定具体 `decision_id` → `Decision Detail`
  - 聚合信号未绑定单一 `decision_id` → `Decision Center / List`
- `product missing module binding` → `Product Detail`
- `product missing repository binding` → `Product Detail`
- `product missing both bindings` → `Product Detail`
- `recent_activity: module` → `Module Detail`
- `recent_activity: release` → 所属 `Module Detail`
- `recent_activity: product` → `Product Detail`
- `recent_activity: repository` → `Repository Binding Detail / Workspace`
- `recent_activity: decision` → `Decision Detail`
- `recent_activity: product_module_binding` → `Product Detail`
- `recent_activity: product_repository_binding` → `Repository Binding Detail / Workspace`
- `recent_activity: module_repository_binding` → `Repository Binding Detail / Workspace`

### 5.3 当前阶段源码设计层基线

- 前端必须明确 `Dashboard Home` 页面分层、最小路由结构与组件职责
- 前端必须明确整页状态、分区状态、空状态与返回路径模型
- 前端必须明确从 Dashboard 跳出再返回时的来源标记与上下文恢复规则
- 前端必须冻结最小来源上下文参数：`fromDashboard=true`、`dashboardSection`、`dashboardReturnTo=/dashboard`
- 后端必须明确 Dashboard 聚合读取的模块边界与接口分组
- 当前阶段必须为 `Dashboard + Feedback` 落地最小 `.proto` 合同源
- 当前阶段必须提前定义联调重置脚本、基线种子、Dashboard 冷启动与有数据 fixture
- 当前阶段不提前冻结 Go 数据访问层具体工具
- 当前阶段必须明确 `pending_decision_signals` 与 `product_asset_coverage` 的消费 owner 与读取边界

## 6. 当前阶段合同与演进基线

- `.proto` 是 `Dashboard + Feedback` 的唯一合同源
- `buf` 校验链至少覆盖：`build`、`lint`、`generate`、`breaking`
- `.proto` 字段语义必须与正式规格正文、HTTP DTO 与前端消费模型保持单值一致
- 合同演进必须遵守兼容性约束；删除字段后必须保留 `reserved` 字段号，必要时同时保留字段名
- `breaking` 校验必须直接对照仓库主线基准，不允许吞掉失败退出码

## 7. 当前阶段冷启动与验收基线

- 首轮必须允许用户在空系统中进入 Dashboard，并看到明确空状态与下一步 CTA
- 首轮必须允许用户在有结构化资产数据时看到最小概览聚合
- 首轮必须允许用户看到至少一类反馈信号并跳到对应 canonical owner 页面
- 首轮必须允许用户从 Dashboard 跳到 `Product / Repository / Module / Decision` 页面后再返回 Dashboard
- 最近活动在有数据状态下必须可见，在无数据状态下必须有明确空区块语义
- 当前阶段验收不得依赖手工补 SQL 才能建立最小联调环境
- 当前阶段必须提供可重复执行的重置脚本、基线种子与异常路径验证前提

当前阶段关于空状态与 CTA 优先级的补充冻结如下：

1. `module_count = 0 && product_count = 0 && repository_count = 0 && decision_count = 0` → Dashboard 进入冷启动空系统状态，主 CTA 指向 `Module Registry / Create`
2. `module_count = 0 && (product_count > 0 || repository_count > 0 || decision_count > 0)` → Dashboard 进入“已有结构化资产但仍无 Module”的非空缺口状态，主 CTA 仍指向 `Module Registry / Create`
3. `module_count > 0 && product_count = 0` → 主 CTA 指向 `Product Registry / Create`
4. `module_count > 0 && product_count > 0 && repository_count = 0` → 主 CTA 指向 `Repository Binding / Create`
5. 存在 `pending_decision_signals` → 主 CTA 指向最高优先级决策信号落点
6. 存在 `product missing both bindings` → 主 CTA 指向对应 `Product Detail`
7. 存在 `product missing repository binding` → 主 CTA 指向对应 `Product Detail`
8. 存在 `product missing module binding` → 主 CTA 指向对应 `Product Detail`
9. 无缺口且有活动数据 → Dashboard 进入“系统已就绪”中性状态，不再展示强制主 CTA

补充约束：

- 同时存在多类缺口时，只允许按上述顺序展示一个主 CTA，其他动作降级到区块内次级入口
- 当前阶段不得在空状态下并排给出多个同级主 CTA

## 8. 非目标矩阵

- 第二套资产注册或绑定主线
- 复杂驾驶舱 / BI 分析
- 外部遥测、通知中心与自动消息回流
- 独立 `Feedback` 重实体 CRUD
- Capability Growth 完整评分系统
- `Feature / Opportunity / Experiment`
- GitHub OAuth / 自动导入
- 自动扫描代码
- AI 自动建议或自动补全
- 独立 `React Native` 客户端
- 完整 `PWA`

## 9. 本阶段校验清单

进入 `phase05` 后续 `/spec` 前，必须再次确认：

1. 当前阶段直接执行层上游是否仍为 `phase04-10 / 11 / 14`
2. `Dashboard` 与 `Feedback Signal Layer` 是否仍然是当前阶段唯一主交付对象
3. `Feedback` 是否仍保持“派生信号层”而不是独立重实体
4. Dashboard 是否仍只承接读与跳转，而不承接第二套写入主线
5. 是否仍采用单一 `React Web` 前端交付策略
