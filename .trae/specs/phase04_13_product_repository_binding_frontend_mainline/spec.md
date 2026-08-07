# Phase04-13 前端 Product Registry 与 Repository Binding 主线实现 Spec

## Why

`phase04-05 / 06 / 10` 已经把 `Product Registry` 与 `Repository Binding` 的页面、路由、组件树、状态模型与正式规格正文冻结到可直接实现的层级，`phase04-12` 又补齐了真实后端、数据主线、重置脚本与旧候选读取兼容链路。但当前前端仓库仍只有 `Module Registry` 与 `Decision Center` 的可运行主线，`Product / Repository / Binding` 还没有对应的页面、路由、数据适配与最小上下文入口，后续阶段无法直接在前端走通创建、详情与三类绑定关系。

因此，`phase04-13` 必须把 `Product Registry` 与 `Repository Binding` 从“已完成前端设计 + 已具备后端接口”推进为“前端主线可运行”：落实 `Product List / Create / Detail` 与 `Repository Binding / List / Create / Detail` 页面及路由、接入 `phase04-12` 真实 API、实现三类绑定关系面板与最小上下文入口，并把 `Module Detail` 中旧绑定入口正式收敛为兼容入口或轻量跳转。

## What Changes

- 实现 `Product List / Create / Detail` 前端主线
- 实现 `Repository Binding / List / Create / Detail(or Workspace)` 前端主线
- 实现 `frontend/src/routes/products/` 与 `frontend/src/routes/repositories/` 路由文件，并在根布局接入导航入口
- 实现 `frontend/src/features/product-registry/` 与 `frontend/src/features/repository-binding/` 下的页面、组件、类型、数据适配与最小页面状态承接
- 直接接入 `phase04-12` 已落地的真实后端 API，不新增并列 mock 数据主线
- 实现三类绑定关系的候选读取、选择、提交、成功 reread 与错误停留规则
- 实现从 `Module Detail / Product Detail` 发起的最小上下文入口与返回路径承接
- 将 `Module Detail` 中旧绑定入口从直接写入承接位收敛为兼容入口或轻量跳转，不再形成第二主工作台
- 落实单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略
- **BREAKING**：`Module Detail` 不再继续承接 `BindModuleToProduct / MapModuleToRepository` 的正式前端写入流程；该页只保留兼容入口、只读摘要或轻量跳转
- **BREAKING**：`Product / Repository / Binding` 前端主线必须直接消费 `phase04-12` 真实 API 与正式字段语义，不得保留并列 mock / 临时对象语义作为第二数据主线

## Impact

- Affected specs:
  - `phase04_05_frontend_page_route_component_design`
  - `phase04_06_frontend_state_interaction_flow`
  - `phase04_10_product_repository_binding_formal_spec`
  - `phase04_12_product_repository_binding_backend_data_mainline`
- Affected code:
  - `frontend/src/routes/__root.tsx`
  - `frontend/src/routes/products/`
  - `frontend/src/routes/repositories/`
  - `frontend/src/features/product-registry/`
  - `frontend/src/features/repository-binding/`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/module-registry/components/module-binding-panel.tsx`
  - `frontend/src/routeTree.gen.ts`

## ADDED Requirements

### Requirement: Product Registry 与 Repository Binding 路由及页面主线必须可运行

系统 SHALL 按 `phase04-05 / 10` 已冻结的页面文件落点与路由语义，落实 `Product Registry` 与 `Repository Binding` 的前端可运行主线，不得改写为第二套路由树或并列工作台。

#### Scenario: Product Registry 页面与路由主线

- **WHEN** 用户进入 `/products`
- **THEN** 系统必须能渲染 `ProductListPage`
- **AND** 用户必须能继续进入 `/products/new` 与 `/products/:productId`
- **AND** 页面文件与路由文件落点必须与 `phase04-05` 已冻结结论一致

#### Scenario: Repository Binding 页面与路由主线

- **WHEN** 用户进入 `/repositories`
- **THEN** 系统必须能渲染 `RepositoryBindingListPage`
- **AND** 用户必须能继续进入 `/repositories/new` 与 `/repositories/:repositoryId`
- **AND** 页面文件与路由文件落点必须与 `phase04-05` 已冻结结论一致
- **AND** 不得把 `Repository Binding Detail / Workspace` 再拆成独立子路由工作台

#### Scenario: 根布局导航接入

- **WHEN** 用户进入前端应用根布局
- **THEN** 根导航必须同时提供 `Module Registry`、`Decision Center`、`Product Registry` 与 `Repository Binding` 的可见入口
- **AND** `Product Registry` 导航入口必须进入 `/products`
- **AND** `Repository Binding` 导航入口必须进入 `/repositories`
- **AND** 不得通过隐藏调试链接或仅靠地址栏直输承接正式主线

### Requirement: 前端数据适配必须直接消费 phase04-12 真实 API

系统 SHALL 为 `Product Registry` 与 `Repository Binding` 提供最小前端数据适配层，但该适配层必须直接消费 `phase04-12` 已落地的真实 HTTP API，而不是复制 `phase02-10` 的 mock / real 双轨主线。

#### Scenario: Product Registry 数据适配落点

- **WHEN** 实现 `Product Registry` 前端数据层
- **THEN** 必须至少落地：
- **AND** `frontend/src/features/product-registry/types.ts`
- **AND** `frontend/src/features/product-registry/data/api-adapter.ts`
- **AND** `frontend/src/features/product-registry/data/product-registry-adapter.ts`
- **AND** `product-registry-adapter.ts` 必须直接导出真实 API 实现，不提供并列 `mock-adapter.ts`

#### Scenario: Repository Binding 数据适配落点

- **WHEN** 实现 `Repository Binding` 前端数据层
- **THEN** 必须至少落地：
- **AND** `frontend/src/features/repository-binding/types.ts`
- **AND** `frontend/src/features/repository-binding/data/api-adapter.ts`
- **AND** `frontend/src/features/repository-binding/data/repository-binding-adapter.ts`
- **AND** `repository-binding-adapter.ts` 必须直接导出真实 API 实现，不提供并列 `mock-adapter.ts`

#### Scenario: HTTP 与字段语义承接

- **WHEN** 页面通过数据适配层读取列表、详情、候选或提交创建/绑定写入
- **THEN** 所有字段、状态枚举、返回路径语义必须对齐 `phase04-10 / 12` 的正式规格与真实 API
- **AND** 不得在前端适配层新增后端未定义的业务字段或第二套状态枚举
- **AND** 不得用临时对象结构把 `Repository Binding Detail / Workspace` 重写成第二套页面语义

### Requirement: Product Registry 前端必须承接列表、创建与详情闭环

系统 SHALL 在 `Product Registry` 前端承接列表读取、创建写入、详情读取、模块绑定与进入仓库绑定主线的最小交互闭环。

#### Scenario: Product List 闭环

- **WHEN** 用户进入 `/products`
- **THEN** 页面必须通过正式适配层读取 `ProductListRead`
- **AND** 页面必须承接 `queryText / statusFilter` 搜索参数
- **AND** 页面必须区分 `initial-loading / ready / empty / error`
- **AND** 空状态主动作必须直接进入 `ProductCreatePage`

#### Scenario: Product Create 闭环

- **WHEN** 用户从 `ProductListPage` 或 `Module Detail` 进入 `ProductCreatePage` 并成功提交最小字段
- **THEN** 系统必须默认回流到对应的 `ProductDetailPage`
- **AND** 回流时必须继续携带创建页已有的来源标记与必要上下文参数（承接 `phase04-06` 创建成功后来源上下文继承模式）：
  - `fromList` 存在 → 必须继续携带 `fromList` 与原 `queryText / statusFilter`
  - `fromModuleDetail` 存在 → 必须继续携带 `fromModuleDetail` 与原 `moduleId / moduleName`
  - 无来源参数 → `direct-entry`，不携带来源标记
- **AND** 回流后 `ProductDetailPage` 必须能基于继承的来源标记按真实来源返回
- **AND** 创建失败时必须保留当前草稿与来源上下文，错误必须显示在表单上下文
- **AND** 用户主动取消返回时必须按真实来源决定（`fromList` → 回 `Product List` + 原 `queryText / statusFilter`；`fromModuleDetail` → 回原 `ModuleDetailPage`；`direct-entry` → 回 `Product List` 默认筛选参数）

#### Scenario: Product Detail 闭环

- **WHEN** 用户进入 `ProductDetailPage`
- **THEN** 页面必须承接产品详情读取、已绑定模块摘要、已绑定仓库摘要与 `BindModuleToProduct` 写入触点
- **AND** 用户必须能够从当前页进入 `Repository Binding` 正式主入口
- **AND** `BindModuleToProduct` 成功后必须停留在当前 `ProductDetailPage` 并重新读取详情结果
- **AND** `BindModuleToProduct` 失败时必须停留在当前页，当前已选候选 `Module` 必须继续保留，错误必须停留在面板上下文
- **AND** 候选 `Module` 为空时必须展示明确的无可绑定候选空状态提示，不得误报为接口错误
- **AND** 用户主动返回时必须基于继承的来源标记按真实来源返回（`fromList` → 回列表 + 原 `queryText / statusFilter`；`fromModuleDetail` → 回原 `ModuleDetailPage`；`direct-entry` → 回列表默认筛选参数）

### Requirement: Repository Binding 前端必须承接列表、创建、详情与两类绑定闭环

系统 SHALL 在 `Repository Binding` 前端承接列表读取、创建写入、详情读取、`BindRepositoryToProduct` 与 `MapModuleToRepository` 的最小交互闭环。

#### Scenario: Repository Binding / List 闭环

- **WHEN** 用户进入 `/repositories`
- **THEN** 页面必须通过正式适配层读取 `RepositoryListRead`
- **AND** 页面必须承接 `queryText / statusFilter` 搜索参数
- **AND** 页面必须区分 `initial-loading / ready / empty / error`
- **AND** 空状态主动作必须直接进入 `RepositoryCreatePage`

#### Scenario: Repository Create 闭环

- **WHEN** 用户从 `Repository Binding / List`、`Product Detail` 或 `Module Detail` 进入 `RepositoryCreatePage` 并成功提交最小字段
- **THEN** 系统必须默认回流到对应的 `RepositoryBindingDetailPage`
- **AND** 回流时必须继续携带创建页已有的来源标记与必要上下文参数（承接 `phase04-06` 创建成功后来源上下文继承模式）：
  - `fromList` 存在 → 必须继续携带 `fromList` 与原 `queryText / statusFilter`
  - `fromProductDetail` 存在 → 必须继续携带 `fromProductDetail` 与原 `productId / productName`
  - `fromModuleDetail` 存在 → 必须继续携带 `fromModuleDetail` 与原 `moduleId / moduleName`
  - 无来源参数 → `direct-entry`，不携带来源标记
- **AND** 回流后 `RepositoryBindingDetailPage` 必须能基于继承的来源标记按真实来源返回
- **AND** 创建失败时必须保留当前草稿与来源上下文，错误必须显示在表单上下文
- **AND** 用户主动取消返回时必须按真实来源决定（`fromList` → 回 `Repository Binding / List` + 原 `queryText / statusFilter`；`fromProductDetail` → 回原 `ProductDetailPage`；`fromModuleDetail` → 回原 `ModuleDetailPage`；`direct-entry` → 回 `Repository Binding / List` 默认筛选参数）

#### Scenario: Repository Binding Detail / Workspace 闭环

- **WHEN** 用户进入 `RepositoryBindingDetailPage`
- **THEN** 页面必须承接仓库详情读取、已绑定产品摘要、已映射模块摘要、`BindRepositoryToProduct` 与 `MapModuleToRepository` 的候选读取与写入面板
- **AND** 同一时刻只允许一个绑定面板处于打开态（互斥展开），承接 `phase04-06`
- **AND** `BindRepositoryToProduct` 成功后必须停留在当前页并重新读取详情结果，当前活动面板必须回到 `closed`
- **AND** `MapModuleToRepository` 成功后必须停留在当前页并重新读取详情结果，当前活动面板必须回到 `closed`
- **AND** 两类绑定失败时必须停留在当前页，当前已选候选目标必须继续保留，错误必须停留在对应面板上下文
- **AND** 候选 `Product` 或候选 `Module` 为空时必须展示明确的无可绑定候选空状态提示，不得误报为接口错误
- **AND** 用户主动返回时必须基于继承的来源标记按真实来源返回（`fromList` → 回列表 + 原 `queryText / statusFilter`；`fromProductDetail` → 回原 `ProductDetailPage`；`fromModuleDetail` → 回原 `ModuleDetailPage`；`direct-entry` → 回列表默认筛选参数）

### Requirement: Module Detail 旧绑定入口必须收敛为兼容入口或轻量跳转

系统 SHALL 将 `Module Detail` 中当前仍承接正式绑定写入的旧入口收敛为兼容入口或轻量跳转，不再形成第二主工作台。

#### Scenario: Module Detail 发起 Product 绑定动作

- **WHEN** 用户在 `Module Detail` 发起“绑定到 Product”相关动作
- **THEN** 页面只能提供跳转到 `Product Registry` 正式主入口的兼容入口
- **AND** 若目标 `Product` 未确定，则必须进入 `/products` 并携带 `moduleId / moduleName / fromModuleDetail`
- **AND** 若目标 `Product` 已确定，则必须进入 `/products/:productId` 并携带 `moduleId / moduleName / fromModuleDetail`
- **AND** 不得继续在 `Module Detail` 内直接提交 `BindModuleToProduct`

#### Scenario: Module Detail 发起 Repository 映射动作

- **WHEN** 用户在 `Module Detail` 发起“映射到 Repository”相关动作
- **THEN** 页面只能提供跳转到 `Repository Binding` 正式主入口的兼容入口
- **AND** 若目标 `Repository` 未确定，则必须进入 `/repositories` 并携带 `moduleId / moduleName / fromModuleDetail`
- **AND** 若目标 `Repository` 已确定，则必须进入 `/repositories/:repositoryId` 并携带 `moduleId / moduleName / fromModuleDetail`
- **AND** 不得继续在 `Module Detail` 内直接提交 `MapModuleToRepository`

#### Scenario: Module Detail 兼容入口展示语义

- **WHEN** `Module Detail` 展示现有产品绑定与仓库映射摘要
- **THEN** 页面可以继续展示只读摘要
- **AND** 当前页中的绑定交互区必须回落为兼容入口或轻量跳转组件
- **AND** 不得继续保留候选读取、选择器、提交按钮组成的第二主工作台

### Requirement: Product Detail 与 Repository Binding Detail 必须承接上下文入口

系统 SHALL 承接从 `Module Detail` 与 `Product Detail` 发起的最小上下文入口，使正式主线页面能够预填候选目标并保持返回路径一致。

#### Scenario: Module Detail -> Product Detail 上下文承接

- **WHEN** 用户携带 `moduleId / moduleName / fromModuleDetail` 进入 `ProductDetailPage`
- **THEN** `ProductModuleBindingPanel` 必须能够承接该来源上下文
- **AND** 用户成功绑定后，页面必须继续停留在 `ProductDetailPage`
- **AND** 返回动作必须按 `phase04-06` 已冻结的来源规则回到真实来源

#### Scenario: Module Detail -> Repository Binding Detail 上下文承接

- **WHEN** 用户携带 `moduleId / moduleName / fromModuleDetail` 进入 `RepositoryBindingDetailPage`
- **THEN** `RepositoryModuleMappingPanel` 必须能够承接该来源上下文
- **AND** 用户成功映射后，页面必须继续停留在 `RepositoryBindingDetailPage`
- **AND** 返回动作必须按 `phase04-06` 已冻结的来源规则回到真实来源

#### Scenario: Product Detail -> Repository Binding Detail 上下文承接

- **WHEN** 用户携带 `productId / productName / fromProductDetail` 进入 `RepositoryBindingDetailPage`
- **THEN** `RepositoryProductBindingPanel` 必须能够承接该来源上下文
- **AND** 用户成功绑定后，页面必须继续停留在 `RepositoryBindingDetailPage`
- **AND** 返回动作必须按 `phase04-06` 已冻结的来源规则回到真实来源

#### Scenario: 来源上下文刷新恢复与无来源默认行为

- **WHEN** 用户刷新带有 `fromModuleDetail` 或 `fromProductDetail` 的正式主线页面
- **THEN** 页面必须继续恢复对应来源标记，承接 `phase04-06`
- **AND** 不得在刷新后静默丢失来源标记
- **AND** 当用户从外部链接或刷新后直接进入任一创建页或详情页且不存在来源列表上下文时，返回列表必须落到默认列表参数（`queryText = 空`、`statusFilter = all`），不得伪造上一份列表筛选条件，不得回退到任何持久化层中上一次的筛选值

### Requirement: 页面读取、写入与 reread 必须沿用 TanStack Query 主线

系统 SHALL 继续沿用仓库当前 `TanStack Query` 主线，通过 `useQuery` 与 `useMutation` 收口列表、详情、候选与写入后的 reread，不扩写第二套客户端缓存协议。

#### Scenario: 创建成功后的查询失效

- **WHEN** `CreateProduct` 或 `CreateRepository` 提交成功
- **THEN** 前端至少必须失效对应列表查询
- **AND** 回流到详情页后必须读取新建实体详情
- **AND** 不得通过手工拼接局部假数据替代正式 reread

#### Scenario: 绑定成功后的查询失效

- **WHEN** 任一绑定动作提交成功
- **THEN** 前端至少必须失效当前详情查询、对应候选列表查询与相关列表查询
- **AND** 成功结果必须由后端 reread 驱动展示
- **AND** 不得仅靠 toast 或局部假状态宣告成功

### Requirement: 单一 React Web 布局主线必须覆盖 PC 与移动浏览器

系统 SHALL 在同一套 `React Web` 页面与组件主线中，同时承接 `PC` 与移动浏览器布局，不得通过新增第二套移动端页面或第二套交互树满足适配要求。

#### Scenario: PC 布局

- **WHEN** 页面在桌面宽屏展示
- **THEN** `Product List` 与 `Repository Binding / List` 必须优先保持高信息密度列表布局
- **AND** `Product Detail` 与 `Repository Binding Detail / Workspace` 必须优先采用分区式布局，使摘要、已绑定结果与当前绑定动作同页可见
- **AND** `Product Create` 与 `Repository Create` 必须确保表单区与主动作区同屏可见

#### Scenario: 移动浏览器布局

- **WHEN** 页面在移动浏览器展示
- **THEN** `Product List` 与 `Repository Binding / List` 必须重排为单列列表或卡片布局
- **AND** `Product Detail` 与 `Repository Binding Detail / Workspace` 必须按摘要、已绑定结果、动作区的垂直顺序重排
- **AND** `Product Create` 与 `Repository Create` 必须采用单列垂直布局，主动作按钮无需横向滚动即可见
- **AND** 不得引入独立移动端页面体系、独立 `React Native` 客户端或第二套移动端 UI 架构

## MODIFIED Requirements

### Requirement: phase04-10 前端实现设计进入运行态

`phase04-10` 已冻结的前端页面、路由、组件树、状态模型与布局降级要求，在本阶段 SHALL 从“实现前设计”推进为“仓库内可运行的前端主线”，并成为后续联调与验收的直接前置条件。

#### Scenario: 前端主线进入可运行状态

- **WHEN** `phase04-13` 完成
- **THEN** 用户必须能够在前端走通 `Product List -> Product Create -> Product Detail -> BindModuleToProduct` 最小闭环
- **AND** 用户必须能够在前端走通 `Repository Binding / List -> Repository Create -> Repository Binding Detail -> BindRepositoryToProduct / MapModuleToRepository` 最小闭环
- **AND** 从 `Module Detail / Product Detail` 发起的上下文入口必须前后贯通
- **AND** 实现结果不得偏离 `phase04-05 / 06 / 10 / 12` 已冻结的数据语义、回流规则与非目标边界

## REMOVED Requirements

### Requirement: Module Detail 继续承接正式绑定写入工作台

**Reason**: `phase04-10` 已明确 `Module Detail` 在当前阶段只承接绑定摘要展示与兼容跳转入口，不得继续停留在本页直接提交绑定写入。继续保留当前 `ModuleBindingPanel` 的正式写入语义，会让 `Module Detail` 与 `Product Detail / Repository Binding Detail` 并列拥有第二套主工作台。

**Migration**: 本阶段必须把 `ModuleDetailPage` 中当前的 `ModuleBindingPanel` 收敛为兼容入口或轻量跳转组件；正式绑定写入统一迁移到 `ProductDetailPage` 与 `RepositoryBindingDetailPage` 承接，旧候选读取与旧写入入口只保留 transport / 页面层兼容语义，不再作为 canonical 前端主线。
