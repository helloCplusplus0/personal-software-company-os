# Phase04-05 前端页面、路由与组件分层设计 Spec

## Why

`phase04-01` 已冻结页面边界与信息区块，`phase04-02` 已冻结模板字段、状态语义与最小展示模型，`phase04-03` 已冻结三类绑定关系语义、候选范围、上下文入口、canonical owner 与 reread 承接页面，`phase04-04` 已冻结数据读写范围、最小接口边界与错误语义前提。但要真正进入前端实现，还必须把页面文件落点、路由树与 URL 语义、组件树、上下文承接路由参数与 `PC / 移动浏览器` 布局降级策略写成单值结论。否则后续前端实现会在文件组织、路由参数与组件职责划分之间继续漂移。

> 阶段分工约束：本规格只冻结前端页面分层、路由树、组件树、上下文承接路由参数与布局降级策略。后端模块边界、接口分组与 `.proto` 合同设计分别由 `phase04-07` 与 `phase04-08` 承接，不在本规格中提前冻结。运行时实现细节（hook 命名、Query key、store API 命名、缓存时间等）不在本规格中冻结。
>
> 与 `phase04-06` 边界划分：本规格只冻结"上下文承接路由参数"（路由搜索参数的命名与传递规则），不冻结"返回路径与上下文恢复规则"（如 `sessionStorage` 持久化策略、`fromList` 行为语义、Create 提交成功回流、Create 取消返回等）。页面级状态模型（读取状态、草稿状态、绑定面板开闭默认值、面板互斥展开规则、面板开闭状态机、派生视图状态、页面级 UI 状态容器归属）与列表查询条件在路由搜索参数与页面状态之间的承接策略，统一由 `phase04-06` 承接，不在本规格中提前冻结。本规格布局降级策略只冻结 `PC` 与移动浏览器下的布局结构（分区式详情、垂直重排、单列布局等），不冻结面板开闭默认值与互斥展开等状态与交互规则。

## What Changes

- 冻结 `Product Registry` 与 `Repository Binding` 两个模块的页面文件落点
- 冻结 `Product Registry` 与 `Repository Binding` 的路由树与 URL 语义
- 冻结 `Product List / Create / Detail` 的组件树与组件职责
- 冻结 `Repository Binding / List`、`Repository Create`、`Repository Binding Detail / Workspace` 的组件树与组件职责
- 冻结 `Module Detail` 兼容入口的路由参数与上下文承接方式
- 冻结 `Product Detail` 上下文入口的路由参数与上下文承接方式
- 冻结 `PC` 与移动浏览器布局降级策略
- 明确当前阶段不提前冻结运行时实现细节
- 明确列表搜索上下文持久化、返回路径规则与页面级状态模型推迟到 `phase04-06` 冻结

## Impact

- Affected specs: `phase04_product_and_repository_binding_foundation`
- Affected code: 后续 `frontend/src/features/product-registry/` 与 `frontend/src/features/repository-binding/` 的页面、组件、状态模型，后续 `frontend/src/routes/products/` 与 `frontend/src/routes/repositories/` 的路由文件，`Module Detail` 既有绑定面板的兼容入口改造

## ADDED Requirements

### Requirement: Product Registry 页面文件落点冻结

系统 SHALL 将 `Product Registry` 模块的页面文件落点冻结为单值结论，遵循 `phase02-09` 已建立的 `features/<module>/pages/` 模式。

#### Scenario: 判断 Product Registry 页面文件落点

- **WHEN** 后续实现讨论 `Product Registry` 的页面文件位置
- **THEN** 必须得到以下单值结论：
- **AND** `ProductListPage` → `frontend/src/features/product-registry/pages/product-list-page.tsx`
- **AND** `ProductCreatePage` → `frontend/src/features/product-registry/pages/product-create-page.tsx`
- **AND** `ProductDetailPage` → `frontend/src/features/product-registry/pages/product-detail-page.tsx`
- **AND** 不得在当前阶段引入超出上述三页的 `Product Registry` 页面文件

### Requirement: Repository Binding 页面文件落点冻结

系统 SHALL 将 `Repository Binding` 模块的页面文件落点冻结为单值结论，遵循 `phase02-09` 已建立的 `features/<module>/pages/` 模式。

#### Scenario: 判断 Repository Binding 页面文件落点

- **WHEN** 后续实现讨论 `Repository Binding` 的页面文件位置
- **THEN** 必须得到以下单值结论：
- **AND** `RepositoryBindingListPage`（对应上游 `Repository Binding / List` 语义）→ `frontend/src/features/repository-binding/pages/repository-binding-list-page.tsx`
- **AND** `RepositoryCreatePage`（对应上游 `Repository Create` 语义）→ `frontend/src/features/repository-binding/pages/repository-create-page.tsx`
- **AND** `RepositoryBindingDetailPage`（对应上游 `Repository Binding Detail / Workspace` 语义，保留绑定工作台语义）→ `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
- **AND** 不得在当前阶段引入超出上述三页的 `Repository Binding` 页面文件

### Requirement: Product Registry 路由树与 URL 语义冻结

系统 SHALL 将 `Product Registry` 的路由树与 URL 语义冻结为单值结论，遵循 `phase02-09` 已建立的 `routes/<module>/` 模式与 `TanStack Router` 文件路由约定。

#### Scenario: 判断 Product Registry 路由树

- **WHEN** 后续实现讨论 `Product Registry` 的路由结构
- **THEN** 必须得到以下单值结论：
- **AND** `ProductListRoute` → `/products` → `frontend/src/routes/products/index.tsx`
- **AND** `ProductCreateRoute` → `/products/new` → `frontend/src/routes/products/new.tsx`
- **AND** `ProductDetailRoute` → `/products/$productId` → `frontend/src/routes/products/$productId.tsx`
- **AND** 不得把 `Product Detail` 提前拆成独立的子路由工作台
- **AND** 不得引入与 `TanStack Router` 文件路由约定冲突的第二套路由体系

### Requirement: Repository Binding 路由树与 URL 语义冻结

系统 SHALL 将 `Repository Binding` 的路由树与 URL 语义冻结为单值结论，遵循 `phase02-09` 已建立的 `routes/<module>/` 模式与 `TanStack Router` 文件路由约定。

#### Scenario: 判断 Repository Binding 路由树

- **WHEN** 后续实现讨论 `Repository Binding` 的路由结构
- **THEN** 必须得到以下单值结论：
- **AND** `RepositoryBindingListRoute`（对应上游 `Repository Binding / List` 语义）→ `/repositories` → `frontend/src/routes/repositories/index.tsx`
- **AND** `RepositoryCreateRoute`（对应上游 `Repository Create` 语义）→ `/repositories/new` → `frontend/src/routes/repositories/new.tsx`
- **AND** `RepositoryBindingDetailRoute`（对应上游 `Repository Binding Detail / Workspace` 语义，保留绑定工作台语义）→ `/repositories/$repositoryId` → `frontend/src/routes/repositories/$repositoryId.tsx`
- **AND** 不得把 `Repository Binding Detail / Workspace` 提前拆成独立的子路由工作台
- **AND** 不得引入与 `TanStack Router` 文件路由约定冲突的第二套路由体系

### Requirement: Product List 组件树冻结

系统 SHALL 将 `Product Registry / List` 的组件树冻结为单值结论，使后续实现可直接按组件职责进入编码。

#### Scenario: 判断 Product List 组件树

- **WHEN** 后续实现讨论 `Product Registry / List` 的组件结构
- **THEN** 必须得到以下单值组件树：
- **AND** `ProductListPageShell`
  - `ProductListToolbar`（承接 `queryText` 搜索输入、`statusFilter` 状态筛选与进入 `Product Create` 的入口）
  - `ProductListContent`（只承接产品列表读取结果）
  - `ProductListTableOrCards`（承接列表项展示，每项至少展示 `name / description / status / created_at / module_bind_count / repository_bind_count`）
  - `ProductListEmptyState`（无产品时引导用户进入 `Product Create`）
- **AND** `ProductListToolbar` 的筛选维度只冻结 `queryText / statusFilter`，不引入 `providerFilter` 或其他筛选维度
- **AND** 组件默认归属于 `ProductListPage`，只有在确有跨页复用证据时才允许抽为共享组件

### Requirement: Product Create 组件树冻结

系统 SHALL 将 `Product Create` 的组件树冻结为单值结论。

#### Scenario: 判断 Product Create 组件树

- **WHEN** 后续实现讨论 `Product Create` 的组件结构
- **THEN** 必须得到以下单值组件树：
- **AND** `ProductCreatePageShell`
  - `ProductCreateForm`（只承接 `name / description / status` 最小录入，`status` 默认预填 `active`）
  - `ProductCreateActions`（承接 `CreateProduct` 提交与返回列表路径）
- **AND** 不得在当前阶段扩写 `customer / value_proposition / business_model / metrics` 等超出 `v0.1` 的表单字段

### Requirement: Product Detail 组件树冻结

系统 SHALL 将 `Product Detail` 的组件树冻结为单值结论，使详情读取、已绑定模块列表、已绑定仓库列表、`BindModuleToProduct` 写入触点与进入仓库绑定的上下文入口各有明确组件承接。

#### Scenario: 判断 Product Detail 组件树

- **WHEN** 后续实现讨论 `Product Detail` 的组件结构
- **THEN** 必须得到以下单值组件树：
- **AND** `ProductDetailPageShell`
  - `ProductSummaryCard`（只承接 `Product` 核心字段 `id / name / description / status / created_at`）
  - `ProductBoundModuleListSection`（承接已绑定 `Module` 列表展示与 `BindModuleToProduct` 写入触点）
    - `ProductModuleBindingPanel`（承接候选 `Module` 读取、空状态、选择与 `BindModuleToProduct` 写入）
  - `ProductBoundRepositoryListSection`（承接已绑定 `Repository` 列表展示，只读摘要）
  - `ProductRepositoryBindingEntry`（承接进入 `Repository Binding Detail / Workspace` 的上下文跳转入口）
- **AND** `ProductModuleBindingPanel` 是 `Product Detail` 中唯一直接承接 `BindModuleToProduct` 写入的组件
- **AND** `ProductRepositoryBindingEntry` 只提供上下文跳转入口，不得在 `Product Detail` 内承接第二套仓库绑定写入流程
- **AND** 不得把 `Product Detail` 提前拆成独立的子路由工作台

### Requirement: Repository Binding / List 组件树冻结

系统 SHALL 将 `Repository Binding / List` 的组件树冻结为单值结论。

#### Scenario: 判断 Repository Binding / List 组件树

- **WHEN** 后续实现讨论 `Repository Binding / List` 的组件结构
- **THEN** 必须得到以下单值组件树：
- **AND** `RepositoryBindingListPageShell`
  - `RepositoryBindingListToolbar`（承接 `queryText` 搜索输入、`statusFilter` 状态筛选与进入 `Repository Create` 的入口）
  - `RepositoryBindingListContent`（只承接仓库列表读取结果）
  - `RepositoryBindingListTableOrCards`（承接列表项展示，每项至少展示 `name / url / provider / status / created_at / product_bind_count / module_bind_count`）
  - `RepositoryBindingListEmptyState`（无仓库时引导用户进入 `Repository Create`）
- **AND** `RepositoryBindingListToolbar` 的筛选维度只冻结 `queryText / statusFilter`，不引入 `providerFilter`

### Requirement: Repository Create 组件树冻结

系统 SHALL 将 `Repository Create` 的组件树冻结为单值结论。

#### Scenario: 判断 Repository Create 组件树

- **WHEN** 后续实现讨论 `Repository Create` 的组件结构
- **THEN** 必须得到以下单值组件树：
- **AND** `RepositoryCreatePageShell`
  - `RepositoryCreateForm`（只承接 `name / url / provider / status` 最小录入，`status` 默认预填 `active`）
  - `RepositoryCreateActions`（承接 `CreateRepository` 提交与返回列表路径）
- **AND** 不得在当前阶段扩写 `oauth_binding / remote_import_status / sync_cursor / scanned_commit` 等自动化集成字段

### Requirement: Repository Binding Detail / Workspace 组件树冻结

系统 SHALL 将 `Repository Binding Detail / Workspace` 的组件树冻结为单值结论，使详情读取、已绑定产品列表、已映射模块列表、`BindRepositoryToProduct` 与 `MapModuleToRepository` 写入触点各有明确组件承接。

#### Scenario: 判断 Repository Binding Detail / Workspace 组件树

- **WHEN** 后续实现讨论 `Repository Binding Detail / Workspace` 的组件结构
- **THEN** 必须得到以下单值组件树：
- **AND** `RepositoryBindingDetailPageShell`
  - `RepositorySummaryCard`（只承接 `Repository` 核心字段 `id / name / url / provider / status / created_at`）
  - `RepositoryBoundProductListSection`（承接已绑定 `Product` 列表展示与 `BindRepositoryToProduct` 写入触点）
    - `RepositoryProductBindingPanel`（承接候选 `Product` 读取、空状态、选择与 `BindRepositoryToProduct` 写入）
  - `RepositoryMappedModuleListSection`（承接已映射 `Module` 列表展示与 `MapModuleToRepository` 写入触点）
    - `RepositoryModuleMappingPanel`（承接候选 `Module` 读取、空状态、选择与 `MapModuleToRepository` 写入）
- **AND** `RepositoryProductBindingPanel` 是该页面中唯一直接承接 `BindRepositoryToProduct` 写入的组件
- **AND** `RepositoryModuleMappingPanel` 是该页面中唯一直接承接 `MapModuleToRepository` 写入的组件
- **AND** 不得把 `BindModuleToProduct` 迁入该页面形成并列主写入流程
- **AND** 不得把 `Repository Binding Detail / Workspace` 提前拆成独立的子路由工作台

### Requirement: Module Detail 兼容入口路由与上下文承接冻结

系统 SHALL 将 `Module Detail` 旧入口的兼容跳转路由参数与上下文承接方式冻结为单值结论，承接 `phase04-03` 已冻结的兼容入口规则。

#### Scenario: 判断 Module Detail 兼容跳转到 Product Detail

- **WHEN** 用户从 `Module Detail` 发起与 `Product` 绑定相关的后续动作
- **THEN** 必须进入 `BindModuleToProduct` 的正式主入口（ `Product Detail` 或 `Product Registry / List` ）
- **AND** 必须携带 `moduleId / moduleName / fromModuleDetail` 作为上下文搜索参数
- **AND** 若目标 `Product` 尚未确定，必须先导航到 `/products` 并携带上下文搜索参数，由用户选择目标 `Product` 后进入 `/products/$productId`
- **AND** 若目标 `Product` 已确定，必须直接导航到 `/products/$productId` 并额外携带 `productId` 作为目标页身份参数，与上下文搜索参数拆开传递
- **AND** `ProductDetailPage` 接收上下文搜索参数后，必须能预填 `ProductModuleBindingPanel` 的候选 `Module` 选择
- **AND** 不得在 `Module Detail` 内直接提交 `BindModuleToProduct`

#### Scenario: 判断 Module Detail 兼容跳转到 Repository Binding Detail / Workspace

- **WHEN** 用户从 `Module Detail` 发起与 `Repository` 映射相关的后续动作
- **THEN** 必须进入 `MapModuleToRepository` 的正式主入口（ `Repository Binding Detail / Workspace` 或 `Repository Binding / List` ）
- **AND** 必须携带 `moduleId / moduleName / fromModuleDetail` 作为上下文搜索参数
- **AND** 若目标 `Repository` 尚未确定，必须先导航到 `/repositories` 并携带上下文搜索参数，由用户选择目标 `Repository` 后进入 `/repositories/$repositoryId`
- **AND** 若目标 `Repository` 已确定，必须直接导航到 `/repositories/$repositoryId` 并额外携带 `repositoryId` 作为目标页身份参数，与上下文搜索参数拆开传递
- **AND** `RepositoryBindingDetailPage` 接收上下文搜索参数后，必须能预填 `RepositoryModuleMappingPanel` 的候选 `Module` 选择
- **AND** 不得在 `Module Detail` 内直接提交 `MapModuleToRepository`

### Requirement: Product Detail 上下文入口路由与上下文承接冻结

系统 SHALL 将 `Product Detail` 发起 `Repository` 绑定相关动作时的上下文入口路由参数冻结为单值结论，承接 `phase04-03` 已冻结的上下文入口规则。

#### Scenario: 判断 Product Detail 上下文跳转到 Repository Binding Detail / Workspace

- **WHEN** 用户从 `Product Detail` 发起与 `Repository` 绑定相关的后续动作
- **THEN** 必须进入 `BindRepositoryToProduct` 的正式主入口（ `Repository Binding Detail / Workspace` 或 `Repository Binding / List` ）
- **AND** 必须携带 `productId / productName / fromProductDetail` 作为上下文搜索参数
- **AND** 若目标 `Repository` 尚未确定，必须先导航到 `/repositories` 并携带上下文搜索参数，由用户选择目标 `Repository` 后进入 `/repositories/$repositoryId`
- **AND** 若目标 `Repository` 已确定，必须直接导航到 `/repositories/$repositoryId` 并额外携带 `repositoryId` 作为目标页身份参数，与上下文搜索参数拆开传递
- **AND** `RepositoryBindingDetailPage` 接收上下文搜索参数后，必须能预填 `RepositoryProductBindingPanel` 的候选 `Product` 选择
- **AND** `Product Detail` 自身不得承接第二套仓库绑定写入流程

### Requirement: PC 与移动浏览器布局降级策略冻结

系统 SHALL 在单一 `React Web` 前端交付策略下，冻结 `PC` 与移动浏览器的布局降级策略，不引入第二套移动端 UI 架构。

#### Scenario: 判断 Product List 布局降级

- **WHEN** `ProductListPage` 在 `PC` 桌面环境展示
- **THEN** 必须采用高信息密度列表布局，工具栏与列表内容同屏可见
- **WHEN** `ProductListPage` 在移动浏览器窄屏环境展示
- **THEN** 必须采用单列列表或卡片重排
- **AND** 必须通过信息裁剪降低拥挤度，保留 `name / status / module_bind_count / repository_bind_count` 核心展示
- **AND** 不得引入第二套独立移动端页面

#### Scenario: 判断 Product Create 布局降级

- **WHEN** `ProductCreatePage` 在 `PC` 或移动浏览器环境展示
- **THEN** 必须采用单列垂直布局
- **AND** 主动作按钮无需横向滚动即可见
- **AND** 不得引入第二套独立移动端页面

#### Scenario: 判断 Product Detail 布局降级

- **WHEN** `ProductDetailPage` 在 `PC` 桌面环境展示
- **THEN** 必须采用分区式详情布局，摘要、已绑定模块列表、已绑定仓库列表与绑定入口可同时可见
- **WHEN** `ProductDetailPage` 在移动浏览器窄屏环境展示
- **THEN** 摘要、已绑定模块、已绑定仓库与绑定入口必须按垂直顺序重排
- **AND** 不得引入第二套独立移动端页面

#### Scenario: 判断 Repository Binding / List 布局降级

- **WHEN** `RepositoryBindingListPage` 在 `PC` 桌面环境展示
- **THEN** 必须采用高信息密度列表布局，工具栏与列表内容同屏可见
- **WHEN** `RepositoryBindingListPage` 在移动浏览器窄屏环境展示
- **THEN** 必须采用单列列表或卡片重排
- **AND** 必须通过信息裁剪降低拥挤度，保留 `name / provider / status / product_bind_count / module_bind_count` 核心展示
- **AND** 不得引入第二套独立移动端页面

#### Scenario: 判断 Repository Create 布局降级

- **WHEN** `RepositoryCreatePage` 在 `PC` 或移动浏览器环境展示
- **THEN** 必须采用单列垂直布局
- **AND** 主动作按钮无需横向滚动即可见
- **AND** 不得引入第二套独立移动端页面

#### Scenario: 判断 Repository Binding Detail / Workspace 布局降级

- **WHEN** `RepositoryBindingDetailPage` 在 `PC` 桌面环境展示
- **THEN** 必须采用分区式详情布局，摘要、已绑定产品列表、已映射模块列表与绑定工作台区可同时可见
- **WHEN** `RepositoryBindingDetailPage` 在移动浏览器窄屏环境展示
- **THEN** 摘要、已绑定产品、已映射模块与绑定面板必须按垂直顺序重排
- **AND** 不得引入第二套独立移动端页面

### Requirement: 组件归属原则冻结

系统 SHALL 将组件归属原则冻结为单值结论，避免过早抽象共享组件层。

#### Scenario: 判断组件归属

- **WHEN** 后续实现讨论组件的归属与复用
- **THEN** `ProductSummaryCard`、`ProductBoundModuleListSection`、`ProductModuleBindingPanel`、`ProductBoundRepositoryListSection`、`ProductRepositoryBindingEntry` 默认归属于 `ProductDetailPage`
- **AND** `ProductCreateForm` 默认归属于 `ProductCreatePage`
- **AND** `RepositorySummaryCard`、`RepositoryBoundProductListSection`、`RepositoryProductBindingPanel`、`RepositoryMappedModuleListSection`、`RepositoryModuleMappingPanel` 默认归属于 `RepositoryBindingDetailPage`
- **AND** `RepositoryCreateForm` 默认归属于 `RepositoryCreatePage`
- **AND** 只有在列表页与详情页或多个页面确实共享同一职责时，才允许抽为共享组件
- **AND** 当前阶段不得为了"组件纯洁"提前拆出无明确复用证据的通用组件层

### Requirement: 运行时实现细节不冻结

系统 SHALL 明确当前阶段只冻结页面分层、路由树、组件树、上下文承接与布局降级策略，不冻结运行时实现细节。

#### Scenario: 判断不冻结的运行时实现细节

- **WHEN** 后续实现讨论前端运行时实现细节
- **THEN** 不得提前冻结具体 hook 命名、Query key 细节、store API 命名、缓存时间、请求取消策略或 optimistic update 方案
- **AND** 页面级 UI 状态（草稿、面板开闭、提交错误等瞬时状态）的归属与状态容器选择不作为本规格冻结内容，统一由 `phase04-06` 状态模型与交互流设计承接

## MODIFIED Requirements

### Requirement: Module Detail 绑定面板解释

`Module Detail` 中的 `ModuleBindingPanel` 在 `phase04` 中 SHALL 从 `phase02-09` 的直接写入承接位回落为摘要展示与兼容跳转入口，不再直接承接 `BindModuleToProduct` 或 `MapModuleToRepository` 写入。

#### Scenario: Module Detail 绑定面板回落解释

- **WHEN** 后续实现讨论 `Module Detail` 中的 `ModuleBindingPanel`
- **THEN** 必须将其理解为只读摘要展示与兼容跳转入口
- **AND** 必须提供进入 `Product Detail` 的兼容跳转入口（携带 `moduleId / moduleName / fromModuleDetail` ）
- **AND** 必须提供进入 `Repository Binding Detail / Workspace` 的兼容跳转入口（携带 `moduleId / moduleName / fromModuleDetail` ）
- **AND** 不得继续在 `ModuleBindingPanel` 内直接提交绑定写入
- **AND** 不得把 `ModuleBindingPanel` 扩写为第二个绑定工作台

## REMOVED Requirements

### Requirement: 第二套移动端 UI 架构

**Reason**: 当前阶段已冻结为单一 `React Web` 同时覆盖 `PC` 与移动浏览器，通过布局降级策略而非独立移动端页面体系解决窄屏适配。引入第二套移动端 UI 架构会违反 `shared_baseline` §2.4 与 `architecture_plan` §4.3 的单一前端交付策略。
**Migration**: 若后续确有独立客户端需要，必须进入新的 `phase / audit` 流程，不在 `phase04-05` 当前规格中处理。
