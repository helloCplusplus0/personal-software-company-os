# Phase03-13 Decision Center 前端主线实现 Spec

## Why

`phase03-05 / 06 / 10` 已经把 `Decision Center` 的页面、路由、组件树、状态模型与布局降级策略冻结到可直接实现的层级，`phase03-12` 又补齐了真实后端、数据主线与 `source_context` 的持久化承接。但当前仓库仍只有 `Module Registry` 的前端可运行主线，`Decision Center` 还没有对应的页面、路由、数据适配与从 `Module Detail` 发起的最小上下文入口，后续 `phase03-14` 冷启动验收无法直接走通。

因此，`phase03-13` 必须把 `Decision Center` 从“已完成前端设计 + 已具备后端接口”推进为“前端主线可运行”：落实 `Decision List / Decision Create / Decision Detail` 页面与路由、接入 `phase03-12` 真实 API、实现最小目标关联与返回路径恢复，并继续坚持单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略。

## What Changes

- 实现 `Decision List / Decision Create / Decision Detail` 前端主线
- 实现 `frontend/src/routes/decisions/` 路由文件，并在根布局中接入 `Decision Center` 主导航入口
- 实现 `frontend/src/features/decision-center/` 下的页面、组件、类型、数据适配与最小页面状态承接
- 接入 `phase03-12` 已落地的真实后端 API，不新增并列 mock 数据主线
- 实现列表搜索参数承接、创建回流、详情返回、候选读取、目标关联与 `source_context` 前端承接
- 将 `Module Detail` 中的 `ModuleDecisionEntryPanel` 从只读展示升级为 `Decision Center` 的最小入口触点
- 落实单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局，不引入第二套移动端 UI 架构
- **BREAKING**：`Decision Center` 前端主线必须直接消费 `phase03-12` 真实 API 与数据语义，不得再通过临时 mock 数据、手工填充状态或与 `.proto / HTTP` 语义并列的第二套对象模型兜底

## Impact

- Affected specs:
  - `phase03_05_frontend_page_route_component_design`
  - `phase03_06_frontend_state_interaction_flow`
  - `phase03_10_decision_center_formal_spec`
  - `phase03_12_decision_center_backend_data_mainline`
- Affected code:
  - `frontend/src/routes/__root.tsx`
  - `frontend/src/routes/decisions/`
  - `frontend/src/features/decision-center/`
  - `frontend/src/features/module-registry/components/module-decision-entry-panel.tsx`
  - `frontend/src/routeTree.gen.ts`

## ADDED Requirements

### Requirement: Decision Center 路由与页面主线必须可运行

系统 SHALL 按 `phase03-10 §9.1 / §9.2` 已冻结的页面与路由结构，落实 `Decision Center` 的前端可运行主线，不得改写为第二套路由树或并列工作台。

#### Scenario: 路由文件落点与页面落点

- **WHEN** 实现 `Decision Center` 前端主线
- **THEN** 必须至少落地以下文件语义：
- **AND** `frontend/src/routes/decisions/index.tsx`
- **AND** `frontend/src/routes/decisions/new.tsx`
- **AND** `frontend/src/routes/decisions/$decisionId.tsx`
- **AND** `frontend/src/features/decision-center/pages/decision-list-page.tsx`
- **AND** `frontend/src/features/decision-center/pages/decision-create-page.tsx`
- **AND** `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
- **AND** 不得把 `Decision Detail` 拆成独立子工作台路由

#### Scenario: 根布局主导航接入

- **WHEN** 用户进入前端应用根布局
- **THEN** 根导航必须同时提供 `Module Registry` 与 `Decision Center` 的可见入口
- **AND** `Decision Center` 导航入口必须进入 `/decisions`
- **AND** 不得通过隐藏调试链接或仅靠地址栏直输来承接正式主线

### Requirement: 列表搜索参数必须继续冻结到路由搜索参数层

系统 SHALL 继续使用 `TanStack Router createFileRoute + validateSearch` 承接 `Decision List` 的查询条件，使列表上下文可被刷新、返回与浏览器导航正确恢复。

#### Scenario: DecisionListRoute 搜索参数

- **WHEN** 实现 `frontend/src/routes/decisions/index.tsx`
- **THEN** `DecisionListRoute` 必须通过 `validateSearch` 冻结 `queryText` 与 `statusFilter`
- **AND** `queryText` 默认值必须为 `''`
- **AND** `statusFilter` 默认值必须为 `'all'`
- **AND** 不得把列表查询条件提升为不可见的全局事实源

#### Scenario: 列表上下文跨页面恢复

- **WHEN** 用户从 `DecisionListPage` 进入 `DecisionCreatePage` 或 `DecisionDetailPage` 后主动返回列表
- **THEN** 系统必须恢复原有 `queryText` 与 `statusFilter`
- **AND** 当前阶段允许使用与 `module-list-search-store.ts` 同构的 `Zustand + sessionStorage` 页面上下文缓存承接“最后一次列表搜索”
- **AND** 该上下文缓存只服务返回路径恢复，不得形成第二套列表事实源
- **AND** 必须显式单值化“来源列表上下文存在 / 不存在”：仅当用户从 `DecisionListPage` 进入 `Create / Detail` 时才恢复 `lastSearch`
- **AND** 从 `Module Detail` 入口或外部直达进入 `Create / Detail` 时，返回列表必须落到默认参数（`statusFilter: 'all'`），不得恢复历史筛选
- **AND** 该“是否来自列表”的判定必须通过显式路由搜索参数（`fromList`）单值化承接，不得依赖全局 store 的隐式状态

### Requirement: 前端数据适配必须直接消费 phase03-12 真实 API

系统 SHALL 为 `Decision Center` 提供最小前端数据适配层，但该适配层必须直接消费 `phase03-12` 已落地的真实 HTTP API，而不是复制 `phase02-10` 的 mock / real 双轨主线。

#### Scenario: 数据适配落点

- **WHEN** 实现 `Decision Center` 前端数据层
- **THEN** 必须至少落地：
- **AND** `frontend/src/features/decision-center/types.ts`
- **AND** `frontend/src/features/decision-center/data/api-adapter.ts`
- **AND** `frontend/src/features/decision-center/data/decision-center-adapter.ts`
- **AND** `decision-center-adapter.ts` 必须直接导出真实 API 实现，不提供并列 `mock-adapter.ts`

#### Scenario: HTTP 与合同语义承接

- **WHEN** 数据适配层与页面交互
- **THEN** 列表、详情、创建、候选读取、目标关联的字段语义必须对齐 `phase03-12` 的 HTTP / `.proto` 语义
- **AND** `CreateDecisionRequest` 必须承接 `source_module_id`
- **AND** `DecisionDetailRead` 必须承接 `linked_modules` 与 `source_context`
- **AND** 不得在前端适配层新增后端未定义的业务字段或第二套状态枚举

### Requirement: Decision List 前端必须承接最小读取与空状态闭环

系统 SHALL 在 `DecisionListPage` 中承接列表读取、搜索筛选、空状态引导与进入创建/详情的最小前端闭环。

#### Scenario: 列表读取与内容状态

- **WHEN** 用户进入 `/decisions`
- **THEN** 页面必须通过 `TanStack Query useQuery` 发起列表读取
- **AND** 页面必须至少区分 `initial-loading`、`ready`、`empty`、`error`
- **AND** 错误反馈必须停留在当前页面内容区域
- **AND** 空状态主动作必须直接进入 `DecisionCreatePage`

#### Scenario: 列表内容密度与进入详情

- **WHEN** 列表读取成功
- **THEN** 页面必须展示 `title / status / created_at / link_count / linked_module_summary`
- **AND** 用户必须可从当前列表项进入对应的 `DecisionDetailPage`
- **AND** `PC` 下优先采用高信息密度表格布局，移动浏览器下重排为单列卡片或紧凑列表

### Requirement: Decision Create 前端必须承接来源上下文与创建回流

系统 SHALL 在 `DecisionCreatePage` 中承接从 `Module Detail` 带入的来源上下文、结构化模板录入、提交状态与成功回流。

#### Scenario: 来源上下文进入创建页

- **WHEN** 用户从 `Module Detail` 点击“为当前 Module 记录决策”
- **THEN** 系统必须导航到 `/decisions/new`
- **AND** 必须通过 `DecisionCreateRoute` 的搜索参数承接 `sourceModuleId` 与 `sourceModuleName`
- **AND** `DecisionContextSourcePanel` 必须展示该来源 `Module`
- **AND** 该来源上下文只表示“待关联来源”，不等于已建立正式 `decision_links`

#### Scenario: 创建页提交流

- **WHEN** 用户在 `DecisionCreatePage` 提交 `RecordDecision`
- **THEN** 页面必须通过 `TanStack Query useMutation` 发起创建写入
- **AND** 成功后必须默认回流到新建 `Decision` 对应的 `DecisionDetailPage`
- **AND** 创建失败时必须保留当前草稿与来源上下文展示
- **AND** 错误反馈必须停留在表单上下文内

### Requirement: Decision Detail 前端必须统一承接详情、待关联目标与关联动作

系统 SHALL 在 `DecisionDetailPage` 中统一承接详情读取、已关联目标展示、待关联目标展示、候选读取与最小目标关联动作。

#### Scenario: 详情页读取与待关联目标展示

- **WHEN** 用户进入 `/decisions/:decisionId`
- **THEN** 页面必须通过 `useQuery` 读取 `DecisionDetailRead`
- **AND** `DecisionDetailSummaryCard` 必须展示结构化模板字段与 `source_context`
- **AND** 当 `source_context.source_module_id` 存在且当前 `linked_modules` 中尚无该 `Module` 时，`DecisionPendingLinkTargetCard` 必须显式展示该待关联目标
- **AND** 不得在进入详情页后静默丢失该待关联目标
- **AND** 待关联目标仅在正式 `LinkDecisionToTarget` 写入后由 `reread` 驱动消失
- **AND** 当前阶段不提供“主动放弃关联”出口，`source_context` 作为入口历史记录保留，不因用户放弃而清除；后端无清除接口时不得用前端临时假状态兜底隐藏待关联目标

#### Scenario: 候选读取与最小目标关联

- **WHEN** 页面需要承接 `Decision -> Module` 关联
- **THEN** 页面必须通过 `useQuery` 读取 `ListDecisionModuleCandidates`
- **AND** 页面必须区分 `pending`、`ready`、`empty`、`error`
- **AND** 页面必须通过 `useMutation` 触发 `LinkDecisionToTarget`
- **AND** 关联成功后必须停留在当前 `DecisionDetailPage`
- **AND** 必须在 `onSuccess` 中失效并重新读取当前详情、候选列表与决策列表相关查询
- **AND** 若本次关联目标正是待关联目标，待关联目标卡片必须在重新读取后消失

### Requirement: Module Detail 入口组件必须升级为正式触点

系统 SHALL 将 `ModuleDecisionEntryPanel` 从只读展示升级为 `Decision Center` 的正式入口触点，同时保留当前 `Module Detail` 作为单一宿主页面。

#### Scenario: 入口动作升级

- **WHEN** 实现 `frontend/src/features/module-registry/components/module-decision-entry-panel.tsx`
- **THEN** 面板必须至少提供两个动作：
- **AND** “为当前 Module 记录决策” -> 导航到带 `sourceModuleId / sourceModuleName` 的 `/decisions/new`
- **AND** “查看当前 Module 相关决策” -> 导航到 `/decisions`
- **AND** 当前已展示的相关决策列表项应可直接进入对应的 `DecisionDetailPage`
- **AND** 不得在 `Module Detail` 侧新增中间路由或中间分发组件

### Requirement: 页面写入与读取刷新必须采用 TanStack Query 主线

系统 SHALL 继续沿用仓库当前 `TanStack Query` 主线，通过 `useMutation` 成功后的 `invalidateQueries` 收口列表、详情与候选刷新，不扩写第二套客户端缓存协议。

#### Scenario: 创建成功后的查询失效

- **WHEN** `CreateDecision` 提交成功
- **THEN** 前端至少必须失效 `DecisionList` 相关查询
- **AND** 回流到 `DecisionDetailPage` 后必须读取新建详情
- **AND** 不得通过手工拼接局部页面假数据替代正式 reread

#### Scenario: 关联成功后的查询失效

- **WHEN** `LinkDecisionToTarget` 提交成功
- **THEN** 前端至少必须失效当前 `DecisionDetail`、当前候选列表与 `DecisionList` 相关查询
- **AND** 若页面来自 `Module Detail` 入口，后续刷新结果必须仍由后端 `source_context` / `linked_modules` 驱动

### Requirement: 单一 React Web 布局主线必须覆盖 PC 与移动浏览器

系统 SHALL 在同一套 `React Web` 页面与组件主线中，同时承接 `PC` 与移动浏览器布局，不得通过新增第二套移动端页面或第二套交互树满足适配要求。

#### Scenario: PC 布局

- **WHEN** 页面在桌面宽屏展示
- **THEN** `DecisionListPage` 必须优先保持高信息密度列表布局
- **AND** `DecisionCreatePage` 必须使来源上下文区、表单区与动作区同屏可见
- **AND** `DecisionDetailPage` 必须优先采用分区式布局，使概要、已关联目标、待关联目标与候选关联区同页可见

#### Scenario: 移动浏览器布局

- **WHEN** 页面在移动浏览器展示
- **THEN** `DecisionListPage` 必须重排为单列
- **AND** `DecisionCreatePage` 必须重排为来源上下文区、表单区、动作区的单列垂直布局
- **AND** `DecisionDetailPage` 必须按概要、已关联目标、待关联目标、候选读取与目标关联的垂直顺序重排
- **AND** 不得引入独立移动端页面体系、独立 `React Native` 客户端或把完整 `PWA` 作为当前阶段前提

## MODIFIED Requirements

### Requirement: phase03-10 前端实现设计进入运行态

`phase03-10` 已冻结的前端页面、路由、组件树、状态模型与布局降级要求，在本阶段 SHALL 从“实现前设计”推进为“仓库内可运行的前端主线”，并成为 `phase03-14` 联调与验收的直接前置条件。

#### Scenario: 前端主线进入可运行状态

- **WHEN** `phase03-13` 完成
- **THEN** 用户必须能够在前端走通 `Decision List -> Decision Create -> Decision Detail -> LinkDecisionToTarget` 最小闭环
- **AND** 从 `Module Detail` 发起的来源上下文入口必须前后贯通
- **AND** 实现结果不得偏离 `phase03-10 / 12` 已冻结的数据语义、回流规则与非目标边界

## REMOVED Requirements

### Requirement: Decision Center 继续依赖 mock 数据或临时页面占位

**Reason**: `phase03-12` 已经补齐真实后端与数据主线，`phase03-13` 的职责就是让前端主线直接消费正式接口，而不是继续靠 mock 或手工状态兜底。

**Migration**: 当前阶段前端统一以 `Decision Center` 真实 API 适配层为入口；若后续需要演示数据或 isolated preview，必须在新的 `fix / audit / phase` 中重新冻结，不在 `phase03-13` 当前实现范围中保留并列数据主线。
