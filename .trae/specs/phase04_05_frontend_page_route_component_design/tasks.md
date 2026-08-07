# Tasks

- [x] Task 1: 冻结 Product Registry 与 Repository Binding 页面文件落点
  - [x] SubTask 1.1: 明确 `ProductListPage / ProductCreatePage / ProductDetailPage` 文件路径遵循 `features/product-registry/pages/` 模式
    - 证据：`spec.md` §ADDED Requirements「Product Registry 页面文件落点冻结」Scenario: 判断 Product Registry 页面文件落点 L34-41
  - [x] SubTask 1.2: 明确 `RepositoryBindingListPage / RepositoryCreatePage / RepositoryBindingDetailPage` 文件路径遵循 `features/repository-binding/pages/` 模式，显式保留上游 `Repository Binding / List` 与 `Repository Binding Detail / Workspace` 语义
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 页面文件落点冻结」Scenario: 判断 Repository Binding 页面文件落点 L47-54

- [x] Task 2: 冻结 Product Registry 与 Repository Binding 路由树与 URL 语义
  - [x] SubTask 2.1: 明确 `ProductListRoute / ProductCreateRoute / ProductDetailRoute` 的 URL 与路由文件路径，遵循 `TanStack Router` 文件路由约定
    - 证据：`spec.md` §ADDED Requirements「Product Registry 路由树与 URL 语义冻结」Scenario: 判断 Product Registry 路由树 L60-68
  - [x] SubTask 2.2: 明确 `RepositoryBindingListRoute / RepositoryCreateRoute / RepositoryBindingDetailRoute` 的 URL 与路由文件路径，显式保留上游语义
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 路由树与 URL 语义冻结」Scenario: 判断 Repository Binding 路由树 L74-82

- [x] Task 3: 冻结六个页面的组件树与组件职责
  - [x] SubTask 3.1: 明确 `ProductListPageShell` 组件树（Toolbar + Content + TableOrCards + EmptyState），筛选维度只冻结 `queryText / statusFilter`
    - 证据：`spec.md` §ADDED Requirements「Product List 组件树冻结」Scenario: 判断 Product List 组件树 L88-98
  - [x] SubTask 3.2: 明确 `ProductCreatePageShell` 组件树（Form + Actions），只承接 `name / description / status`
    - 证据：`spec.md` §ADDED Requirements「Product Create 组件树冻结」Scenario: 判断 Product Create 组件树 L104-111
  - [x] SubTask 3.3: 明确 `ProductDetailPageShell` 组件树（SummaryCard + BoundModuleListSection + ModuleBindingPanel + BoundRepositoryListSection + RepositoryBindingEntry）
    - 证据：`spec.md` §ADDED Requirements「Product Detail 组件树冻结」Scenario: 判断 Product Detail 组件树 L117-129
  - [x] SubTask 3.4: 明确 `RepositoryBindingListPageShell` 组件树，筛选维度只冻结 `queryText / statusFilter`
    - 证据：`spec.md` §ADDED Requirements「Repository Binding / List 组件树冻结」Scenario: 判断 Repository Binding / List 组件树 L135-144
  - [x] SubTask 3.5: 明确 `RepositoryCreatePageShell` 组件树（Form + Actions），只承接 `name / url / provider / status`
    - 证据：`spec.md` §ADDED Requirements「Repository Create 组件树冻结」Scenario: 判断 Repository Create 组件树 L150-157
  - [x] SubTask 3.6: 明确 `RepositoryBindingDetailPageShell` 组件树（SummaryCard + BoundProductListSection + ProductBindingPanel + MappedModuleListSection + ModuleMappingPanel）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 组件树冻结」Scenario: 判断 Repository Binding Detail / Workspace 组件树 L163-176

- [x] Task 4: 冻结上下文入口路由与上下文承接
  - [x] SubTask 4.1: 明确 `Module Detail` 兼容跳转到 `Product Detail` 的路由参数（`moduleId / moduleName / fromModuleDetail` + 目标页身份参数 `productId`），区分目标实体未确定/已确定
    - 证据：`spec.md` §ADDED Requirements「Module Detail 兼容入口路由与上下文承接冻结」Scenario: 判断 Module Detail 兼容跳转到 Product Detail L182-190
  - [x] SubTask 4.2: 明确 `Module Detail` 兼容跳转到 `Repository Binding Detail / Workspace` 的路由参数（`moduleId / moduleName / fromModuleDetail` + 目标页身份参数 `repositoryId`），区分目标实体未确定/已确定
    - 证据：`spec.md` §ADDED Requirements「Module Detail 兼容入口路由与上下文承接冻结」Scenario: 判断 Module Detail 兼容跳转到 Repository Binding Detail / Workspace L192-200
  - [x] SubTask 4.3: 明确 `Product Detail` 上下文跳转到 `Repository Binding Detail / Workspace` 的路由参数（`productId / productName / fromProductDetail` + 目标页身份参数 `repositoryId`），区分目标实体未确定/已确定
    - 证据：`spec.md` §ADDED Requirements「Product Detail 上下文入口路由与上下文承接冻结」Scenario: 判断 Product Detail 上下文跳转到 Repository Binding Detail / Workspace L206-214

- [x] Task 5: 冻结 PC 与移动浏览器布局降级策略
  - [x] SubTask 5.1: 明确 `Product List` 布局降级（PC 高密度列表，移动单列/卡片重排）
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器布局降级策略冻结」Scenario: 判断 Product List 布局降级 L220-227
  - [x] SubTask 5.2: 明确 `Product Create` 布局降级（PC/移动单列垂直布局）
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器布局降级策略冻结」Scenario: 判断 Product Create 布局降级 L229-234
  - [x] SubTask 5.3: 明确 `Product Detail` 布局降级（PC 分区式详情，移动垂直重排）
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器布局降级策略冻结」Scenario: 判断 Product Detail 布局降级 L236-242
  - [x] SubTask 5.4: 明确 `Repository Binding / List` 布局降级（PC 高密度列表，移动单列/卡片重排）
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器布局降级策略冻结」Scenario: 判断 Repository Binding / List 布局降级 L244-251
  - [x] SubTask 5.5: 明确 `Repository Create` 布局降级（PC/移动单列垂直布局）
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器布局降级策略冻结」Scenario: 判断 Repository Create 布局降级 L253-258
  - [x] SubTask 5.6: 明确 `Repository Binding Detail / Workspace` 布局降级（PC 分区式详情，移动垂直重排）
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器布局降级策略冻结」Scenario: 判断 Repository Binding Detail / Workspace 布局降级 L260-266

- [x] Task 6: 冻结组件归属原则与运行时实现细节不冻结
  - [x] SubTask 6.1: 明确组件归属原则（默认归属于所属页面，只在确有复用证据时抽为共享组件）
    - 证据：`spec.md` §ADDED Requirements「组件归属原则冻结」Scenario: 判断组件归属 L272-280
  - [x] SubTask 6.2: 明确运行时实现细节不冻结（hook 命名、Query key、store API、缓存时间、optimistic update 等不提前冻结；页面级 UI 状态归属与状态容器选择不冻结，由 phase04-06 承接）
    - 证据：`spec.md` §ADDED Requirements「运行时实现细节不冻结」Scenario: 判断不冻结的运行时实现细节 L286-290

- [x] Task 7: 冻结 Module Detail 绑定面板回落解释（MODIFIED）
  - [x] SubTask 7.1: 明确 `ModuleBindingPanel` 从直接写入承接位回落为摘要展示与兼容跳转入口，提供进入 `Product Detail` 与 `Repository Binding Detail / Workspace` 的跳转
    - 证据：`spec.md` §MODIFIED Requirements「Module Detail 绑定面板解释」Scenario: Module Detail 绑定面板回落解释 L298-305

- [x] Task 8: 移除第二套移动端 UI 架构（REMOVED）
  - [x] SubTask 8.1: 明确不引入第二套移动端 UI 架构，通过布局降级策略解决窄屏适配
    - 证据：`spec.md` §REMOVED Requirements「第二套移动端 UI 架构」L309-312

- [x] Task 9: 完成规格校验
  - [x] SubTask 9.1: 验证页面文件落点与路由树已单值化且遵循 `phase02-09` 模式，`Repository Binding` 命名体系与上游 `Repository Binding / List`、`Repository Binding Detail / Workspace` 一致
    - 证据：`spec.md` §ADDED Requirements「Product Registry 页面文件落点冻结」+「Repository Binding 页面文件落点冻结」+「Product Registry 路由树与 URL 语义冻结」+「Repository Binding 路由树与 URL 语义冻结」
  - [x] SubTask 9.2: 验证六个页面的组件树已单值化且组件职责明确
    - 证据：`spec.md` §ADDED Requirements「Product List / Create / Detail 组件树冻结」+「Repository Binding / List / Create / Detail / Workspace 组件树冻结」
  - [x] SubTask 9.3: 验证上下文入口路由参数已单值化且与 `phase04-03` 兼容入口规则一致
    - 证据：`spec.md` §ADDED Requirements「Module Detail 兼容入口路由与上下文承接冻结」+「Product Detail 上下文入口路由与上下文承接冻结」
  - [x] SubTask 9.4: 验证布局降级策略只冻结布局结构，未越界冻结面板开闭默认值、互斥展开与状态容器归属等 phase04-06 范围内容
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器布局降级策略冻结」+ Why 部分「与 `phase04-06` 边界划分」说明 + What Changes「明确列表搜索上下文持久化、返回路径规则与页面级状态模型推迟到 `phase04-06` 冻结」
  - [x] SubTask 9.5: 验证本规格未越界冻结后端模块边界（`phase04-07`）与 `.proto` 合同设计（`phase04-08`）的设计职责
    - 证据：`spec.md` Why 部分阶段分工约束 + 运行时实现细节不冻结 Requirement — 后端模块边界、接口分组与 `.proto` 合同设计均未在本规格冻结

# Task Dependencies

- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 2`
- `Task 5` depends on `Task 3`
- `Task 7` depends on `Task 3`
- `Task 9` depends on `Task 1` through `Task 8`
