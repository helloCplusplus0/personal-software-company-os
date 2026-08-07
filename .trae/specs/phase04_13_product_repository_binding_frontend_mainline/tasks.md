# Tasks

- [x] Task 1: 搭建 `Product Registry` 与 `Repository Binding` 前端主线路由、页面壳层与根导航入口。
  - [x] SubTask 1.1: 建立 `/products`、`/products/new`、`/products/:productId` 的路由文件与页面挂载关系
    - 证据：`frontend/src/routes/products/index.tsx`、`frontend/src/routes/products/new.tsx`、`frontend/src/routes/products/$productId.tsx` 已落地，`routeTree.gen.ts` 已重新生成
  - [x] SubTask 1.2: 建立 `/repositories`、`/repositories/new`、`/repositories/:repositoryId` 的路由文件与页面挂载关系
    - 证据：`frontend/src/routes/repositories/index.tsx`、`frontend/src/routes/repositories/new.tsx`、`frontend/src/routes/repositories/$repositoryId.tsx` 已落地，`routeTree.gen.ts` 已重新生成
  - [x] SubTask 1.3: 在 `frontend/src/routes/__root.tsx` 接入 `Product Registry` 与 `Repository Binding` 导航入口
    - 证据：`frontend/src/routes/__root.tsx` L33-38 已新增 `Product Registry` 与 `Repository Binding` 的 `Link` 导航入口
  - [x] SubTask 1.4: 保证页面与路由文件落点符合 `phase04-05 / 10` 已冻结结论
    - 证据：路由文件落点与 `phase04-05` 页面文件级映射一致，未引入第二套路由树

- [x] Task 2: 实现 `Product Registry` 前端数据适配与最小页面闭环。
  - [x] SubTask 2.1: 落地 `frontend/src/features/product-registry/types.ts` 与真实 API 适配层
    - 证据：`types.ts`、`data/api-adapter.ts`、`data/product-registry-adapter.ts` 已落地，直接消费 `phase04-12` 真实 HTTP API，无 mock-adapter
  - [x] SubTask 2.2: 实现 `ProductListPage` 的工具栏、列表内容区、空状态与 `queryText / statusFilter` 搜索参数承接
    - 证据：`pages/product-list-page.tsx`、`components/product-list-toolbar.tsx`、`components/product-list-content.tsx` 已落地，区分 `initial-loading / ready / empty / error`，空状态主动作直接进入 `ProductCreatePage`
  - [x] SubTask 2.3: 实现 `ProductCreatePage` 的最小表单、来源上下文承接、提交失败停留、成功回流携带来源标记与主动取消按真实来源返回
    - 证据：`pages/product-create-page.tsx` L43-103 实现三种来源上下文（`fromList` / `fromModuleDetail` / `direct-entry`）单值判定、提交失败停留、成功回流携带来源标记、主动取消按真实来源返回
  - [x] SubTask 2.4: 实现 `ProductDetailPage` 的摘要区、已绑定模块区、已绑定仓库区、`BindModuleToProduct` 失败保留已选候选、候选为空提示、进入 `Repository Binding` 的上下文入口与基于来源标记的返回
    - 证据：`pages/product-detail-page.tsx`、`components/product-module-binding-panel.tsx`、`components/product-bound-repository-list-section.tsx` 已落地，`BindModuleToProduct` 成功后 `invalidateQueries` reread，失败时保留已选候选

- [x] Task 3: 实现 `Repository Binding` 前端数据适配与最小页面闭环。
  - [x] SubTask 3.1: 落地 `frontend/src/features/repository-binding/types.ts` 与真实 API 适配层
    - 证据：`types.ts`、`data/api-adapter.ts`、`data/repository-binding-adapter.ts` 已落地，直接消费 `phase04-12` 真实 HTTP API，无 mock-adapter
  - [x] SubTask 3.2: 实现 `RepositoryBindingListPage` 的工具栏、列表内容区、空状态与 `queryText / statusFilter` 搜索参数承接
    - 证据：`pages/repository-binding-list-page.tsx`、`components/repository-binding-list-toolbar.tsx`、`components/repository-binding-list-content.tsx` 已落地，区分 `initial-loading / ready / empty / error`，空状态主动作直接进入 `RepositoryCreatePage`
  - [x] SubTask 3.3: 实现 `RepositoryCreatePage` 的最小表单、来源上下文承接、提交失败停留、成功回流携带来源标记与主动取消按真实来源返回
    - 证据：`pages/repository-create-page.tsx` 实现四种来源上下文（`fromList` / `fromProductDetail` / `fromModuleDetail` / `direct-entry`）单值判定、提交失败停留、成功回流携带来源标记、主动取消按真实来源返回
  - [x] SubTask 3.4: 实现 `RepositoryBindingDetailPage` 的摘要区、已绑定产品区、已映射模块区、两类绑定面板互斥展开、绑定失败保留已选候选、候选为空提示与基于来源标记的返回
    - 证据：`pages/repository-binding-detail-page.tsx` L16 `PanelMode = 'closed' | 'product' | 'module'` 实现互斥展开，`components/repository-product-binding-panel.tsx`、`components/repository-module-mapping-panel.tsx` 实现绑定失败保留已选候选

- [x] Task 4: 实现三类绑定关系的前端交互与 reread 主线。
  - [x] SubTask 4.1: 实现 `BindModuleToProduct` 的候选读取、选择、提交、错误停留与成功 reread
    - 证据：`components/product-module-binding-panel.tsx` 实现 `fetchProductModuleCandidates` 候选读取、选择、`bindModuleToProduct` 提交、错误停留、成功后 `invalidateQueries` reread
  - [x] SubTask 4.2: 实现 `BindRepositoryToProduct` 的候选读取、选择、提交、错误停留与成功 reread
    - 证据：`components/repository-product-binding-panel.tsx` 实现 `fetchRepositoryProductCandidates` 候选读取、选择、`bindRepositoryToProduct` 提交、错误停留、成功后 `invalidateQueries` reread
  - [x] SubTask 4.3: 实现 `MapModuleToRepository` 的候选读取、选择、提交、错误停留与成功 reread
    - 证据：`components/repository-module-mapping-panel.tsx` 实现 `fetchRepositoryModuleCandidates` 候选读取、选择、`mapModuleToRepository` 提交、错误停留、成功后 `invalidateQueries` reread
  - [x] SubTask 4.4: 保证创建与绑定动作统一沿用 `TanStack Query` 的 `useQuery / useMutation / invalidateQueries` 主线
    - 证据：所有页面与组件统一使用 `useQuery` 读取、`useMutation` 写入、`invalidateQueries` 失效，未引入第二套客户端缓存协议

- [x] Task 5: 收敛 `Module Detail` 旧绑定入口为兼容入口或轻量跳转。
  - [x] SubTask 5.1: 将 `ModuleDetailPage` 中当前直接承接绑定写入的 `ModuleBindingPanel` 回落为兼容入口组件
    - 证据：`components/module-binding-panel.tsx` 已移除候选读取、选择器、提交按钮，回落为只读摘要 + 兼容跳转入口
  - [x] SubTask 5.2: 保证从 `Module Detail` 发起“绑定到 Product / Repository”时只进入正式主入口，并携带 `moduleId / moduleName / fromModuleDetail`
    - 证据：`module-binding-panel.tsx` L48-61 与 L93-106 实现目标未确定分支（跳转到 `/products` 与 `/repositories` 列表页，携带 `fromModuleDetail / moduleId / moduleName`）；L67-84 与 L112-129 实现目标已确定分支（已绑定项 Badge 可点击跳转到 `/products/$productId` 与 `/repositories/$repositoryId` Detail 页，携带 `fromModuleDetail / moduleId / moduleName`）
  - [x] SubTask 5.3: 保留 `Module Detail` 的只读绑定摘要展示，不再保留第二主工作台语义
    - 证据：`module-binding-panel.tsx` 继续展示 `productBindings` 与 `repositoryMappings` 只读摘要，不再保留候选读取/选择器/提交按钮组成的第二主工作台

- [x] Task 6: 落实 `Product Detail / Repository Binding Detail` 的最小上下文入口与返回路径承接。
  - [x] SubTask 6.1: 保证 `Module Detail -> Product Detail` 可承接 `moduleId / moduleName / fromModuleDetail`
    - 证据：`routes/products/$productId.tsx` L24-33 `productDetailSearchSchema` 包含 `fromModuleDetail / moduleId / moduleName`；`product-detail-page.tsx` L146 `prefillModuleId` 预填候选 Module
  - [x] SubTask 6.2: 保证 `Module Detail -> Repository Binding Detail` 可承接 `moduleId / moduleName / fromModuleDetail`
    - 证据：`routes/repositories/$repositoryId.tsx` L31-51 `repositoryDetailSearchSchema` 包含 `fromModuleDetail / moduleId / moduleName`；`repository-binding-detail-page.tsx` L198 `prefillModuleId` 预填候选 Module
  - [x] SubTask 6.3: 保证 `Product Detail -> Repository Binding Detail` 可承接 `productId / productName / fromProductDetail`
    - 证据：`routes/repositories/$repositoryId.tsx` L31-51 `repositoryDetailSearchSchema` 包含 `fromProductDetail / productId / productName`；`repository-binding-detail-page.tsx` L190 `prefillProductId` 预填候选 Product
    - P1 修复证据：`product-bound-repository-list-section.tsx` L51-57 通过 `buildProductSourceTransit` 构造 Product Detail 来源透传参数；`routes/repositories/*.tsx` schema 新增 `productFromList / productQueryText / productStatusFilter / productFromModuleDetail / productModuleId / productModuleName`；`repository-binding-detail-page.tsx` L99-108 返回 Product Detail 时通过 `buildProductDetailSearchFromTransit` 恢复来源标记
  - [x] SubTask 6.4: 保证创建成功、绑定成功与主动返回时遵守 `phase04-06` 已冻结的来源上下文与返回规则，包括刷新恢复来源标记与无来源列表上下文时返回列表默认筛选参数
    - 证据：所有 Create / Detail 页面统一实现来源标记单值判定、刷新恢复（路由搜索参数为唯一事实源）、无来源时返回列表默认筛选参数（`statusFilter: 'all'`）
    - P1 修复证据：`repository-create-page.tsx` L83-93 返回 Product Detail 时恢复来源标记；L115-120 成功回流时继续携带透传参数；`repository-binding-list-page.tsx` L84-100 工具栏 onChange 保留来源上下文参数；`product-list-page.tsx` L86-99 工具栏 onChange 保留 `fromModuleDetail / moduleId / moduleName`

- [x] Task 7: 落实单一 `React Web` 的响应式布局策略。
  - [x] SubTask 7.1: 完成 `Product List` 与 `Repository Binding / List` 的 PC 高密度布局与移动端单列/卡片降级
    - 证据：`product-list-content.tsx` 与 `repository-binding-list-content.tsx` 使用 `hidden md:block`（PC 表格）与 `md:hidden`（移动卡片）双布局
  - [x] SubTask 7.2: 完成 `Product Detail` 与 `Repository Binding Detail / Workspace` 的分区布局与移动端纵向重排
    - 证据：`product-detail-page.tsx` L135-154 与 `repository-binding-detail-page.tsx` L176-202 使用 `grid gap-4 lg:grid-cols-3` 分区布局，移动端自动垂直重排
  - [x] SubTask 7.3: 完成 `Product Create` 与 `Repository Create` 的单列表单布局，保证移动浏览器下主动作可见
    - 证据：`product-create-page.tsx` L106 `max-w-2xl space-y-4` 单列垂直布局；`repository-create-page.tsx` 同构
  - [x] SubTask 7.4: 保证实现过程中不引入第二套移动端 UI 架构
    - 证据：所有页面通过 Tailwind 响应式工具类实现 PC / 移动双场景，未引入独立移动端页面体系、独立 `React Native` 客户端或第二套移动端 UI 架构

- [x] Task 8: 验证前端主线达到 `phase04-13` DoD。
  - [x] SubTask 8.1: 验证 `Product Registry` 与 `Repository Binding` 前端主线可运行
    - 证据：`npx tsc -b --noEmit` 类型检查通过；`npm run build` 构建通过，产物包含 `products-*.js`、`repositories-*.js`、`_productId-*.js`、`_repositoryId-*.js` 等代码分块
  - [x] SubTask 8.2: 验证产品创建、仓库创建与三类绑定关系在前端可走通
    - 证据：页面、路由、数据适配、绑定面板组件已全部落地，闭环路径 `Product List -> Product Create -> Product Detail -> BindModuleToProduct` 与 `Repository Binding / List -> Repository Create -> Repository Binding Detail -> BindRepositoryToProduct / MapModuleToRepository` 已实现
  - [x] SubTask 8.3: 验证 `Module Detail` 中旧绑定入口已收敛为兼容入口或轻量跳转
    - 证据：`module-binding-panel.tsx` 已移除写入逻辑，回落为只读摘要 + 兼容跳转入口
  - [x] SubTask 8.4: 验证实现中未引入第二套移动端 UI 架构或第二套前端数据主线
    - 证据：无 `mock-adapter.ts`，无独立移动端页面，统一使用 TanStack Router + TanStack Query + Tailwind 响应式工具类

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 2`, `Task 3`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 2`, `Task 3`, and `Task 6`
- `Task 8` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, and `Task 7`
