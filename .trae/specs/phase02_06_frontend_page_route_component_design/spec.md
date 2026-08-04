# Phase02-06 前端页面、路由与组件分层设计 Spec

## Why

`phase02-01 ~ 05` 已经冻结了 `Module Registry` 的页面边界、实体主线、创建闭环、读写范围与 API 前提，但这些结论还不足以直接进入前端实现。为了让 `phase02-07 / 08`、`phase02-09` 与后续正式编码不再依赖临场发挥，当前子任务必须产出一份可直接指导 `frontend/` 落地的页面、路由与组件分层设计。

## What Changes

- 产出 `Module List / Module Create / Module Detail / Release Create` 的具体页面分层设计
- 产出一组可直接映射到前端实现的最小路由结构与 URL 语义
- 产出列表、表单、详情、关联面板的组件树与组件职责边界
- 产出 `PC / 移动浏览器` 双场景下的布局降级策略与信息重排规则
- 产出可直接映射到 `frontend/` 的前端文件落点与路由树
- 明确当前阶段不引入第二套移动端 UI 架构、不拆出 `Product / Repository / Decision` 的独立前端主线

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `frontend/` 中 `Module Registry` 的路由定义、页面文件组织、布局壳层、页面组件拆分、响应式布局容器
- Affected code: 预期会直接映射到 `frontend/src/routes/modules/*`、`frontend/src/features/module-registry/pages/*`、`frontend/src/features/module-registry/components/*`

## ADDED Requirements

### Requirement: 前端页面结果必须可直接映射到页面文件

系统 SHALL 将当前阶段前端页面设计收敛为一组可以直接映射到 `frontend/` 页面文件的页面集合，而不是只停留在抽象页面名称。

#### Scenario: 页面集合判定

- **WHEN** 后续实现开始创建 `Module Registry` 页面文件
- **THEN** 页面集合必须至少包含 `ModuleListPage`、`ModuleCreatePage`、`ModuleDetailPage`、`ReleaseCreatePage`
- **AND** 不得在当前阶段额外拆出 `ProductPage`、`RepositoryPage`、`DecisionPage` 的独立主线页面

#### Scenario: 页面职责判定

- **WHEN** 上述页面被映射到具体实现文件
- **THEN** `ModuleListPage` 只承接列表读取、筛选入口、创建入口与进入详情
- **AND** `ModuleCreatePage` 只承接 `CreateModule`
- **AND** `ModuleDetailPage` 只承接详情读取、版本列表、绑定动作与 `Decision` 入口
- **AND** `ReleaseCreatePage` 只承接 `CreateRelease`

### Requirement: 路由树必须冻结到文件落点层

系统 SHALL 将 `Module Registry` 当前阶段的前端路由设计冻结到“路由树 + 文件落点”层，确保后续实现可以直接按既定结构创建文件。

#### Scenario: 最小文件落点

- **WHEN** 后续实现开始创建前端路由与页面文件
- **THEN** 当前阶段至少应能映射出以下文件层级语义：
- **AND** `frontend/src/routes/modules/index.tsx`
- **AND** `frontend/src/routes/modules/new.tsx`
- **AND** `frontend/src/routes/modules/$moduleId.tsx`
- **AND** `frontend/src/routes/modules/$moduleId/releases/new.tsx`
- **AND** `frontend/src/features/module-registry/pages/module-list-page.tsx`
- **AND** `frontend/src/features/module-registry/pages/module-create-page.tsx`
- **AND** `frontend/src/features/module-registry/pages/module-detail-page.tsx`
- **AND** `frontend/src/features/module-registry/pages/release-create-page.tsx`

#### Scenario: 路由树表达

- **WHEN** 后续实现需要查看当前阶段的最小路由树
- **THEN** 路由树至少应表达为：
- **AND** `/modules`
- **AND** `/modules/new`
- **AND** `/modules/:moduleId`
- **AND** `/modules/:moduleId/releases/new`
- **AND** 不得把 `Product`、`Repository`、`Decision` 提前扩写为并列主树

### Requirement: 路由设计必须冻结到 URL 语义层

系统 SHALL 将当前阶段前端最小路由结构冻结到可直接实现的 URL 语义层，而不是只保留抽象 route name。

#### Scenario: 最小路由结构

- **WHEN** 后续实现创建前端路由
- **THEN** 最小路由结构至少应承接以下 URL 语义：
- **AND** `ModuleListRoute` -> `/modules`
- **AND** `ModuleCreateRoute` -> `/modules/new`
- **AND** `ModuleDetailRoute` -> `/modules/:moduleId`
- **AND** `ReleaseCreateRoute` -> `/modules/:moduleId/releases/new`

#### Scenario: 路由进入关系

- **WHEN** 页面之间的进入与返回关系被实现
- **THEN** `ModuleListRoute` 必须能进入 `ModuleCreateRoute`
- **AND** `ModuleListRoute` 必须能进入 `ModuleDetailRoute`
- **AND** `ModuleDetailRoute` 必须能进入 `ReleaseCreateRoute`
- **AND** `ModuleDetailRoute` 必须保留到 `Product Registry / Repository Binding / Decision Center` 的轻量跳转或入口

#### Scenario: 当前阶段不拆子工作台

- **WHEN** 路由结构被进一步展开
- **THEN** 当前阶段不得把 `Module Detail` 提前拆成独立的子路由工作台
- **AND** 不得增加 `Product`、`Repository`、`Decision` 的独立路由主线

### Requirement: 页面壳层与组件树必须冻结到实现结构层

系统 SHALL 为每个页面产出可直接进入实现的页面壳层与组件树设计，而不是只描述“有一个区块”。

#### Scenario: 列表页组件树

- **WHEN** 后续实现 `ModuleListPage`
- **THEN** 页面壳层至少应包含 `ModuleListPageShell`
- **AND** 主组件树至少应包含 `ModuleListToolbar`、`ModuleListContent`、`ModuleListTableOrCards`
- **AND** `ModuleListToolbar` 承接筛选入口与创建入口
- **AND** `ModuleListContent` 只承接模块列表读取结果，不承接详情或创建逻辑

#### Scenario: 创建页组件树

- **WHEN** 后续实现 `ModuleCreatePage`
- **THEN** 页面壳层至少应包含 `ModuleCreatePageShell`
- **AND** 主组件树至少应包含 `ModuleCreateForm`、`ModuleCreateActions`
- **AND** `ModuleCreateForm` 只承接 `name / description / status` 的最小录入结构
- **AND** 不得在当前阶段扩写出导入向导、自动扫描面板或 AI 建议面板

#### Scenario: 详情页组件树

- **WHEN** 后续实现 `ModuleDetailPage`
- **THEN** 页面壳层至少应包含 `ModuleDetailPageShell`
- **AND** 主组件树至少应包含 `ModuleSummaryCard`、`ModuleReleaseListSection`、`ModuleBindingPanel`、`ModuleDecisionEntryPanel`
- **AND** `ModuleBindingPanel` 直接承接 `BindModuleToProduct` 与 `MapModuleToRepository`
- **AND** `ModuleDecisionEntryPanel` 只承接只读展示或跳转，不扩写成独立写入面板
- **AND** `ModuleSummaryCard` 只承接模块核心字段与状态表达，不扩写成可编辑工作区

#### Scenario: 版本登记页组件树

- **WHEN** 后续实现 `ReleaseCreatePage`
- **THEN** 页面壳层至少应包含 `ReleaseCreatePageShell`
- **AND** 主组件树至少应包含 `ReleaseCreateForm`、`ReleaseCreateActions`
- **AND** `ReleaseCreateForm` 只承接当前模块上下文中的版本登记最小字段
- **AND** 不得将其扩写为独立版本管理工作台

### Requirement: 组件归属必须冻结到共享与页面专属边界

系统 SHALL 将当前阶段组件归属冻结到“页面专属组件”和“可复用共享组件”边界，避免后续实现时无序抽象。

#### Scenario: 页面专属组件

- **WHEN** 后续实现讨论页面专属组件
- **THEN** `ModuleSummaryCard`、`ModuleReleaseListSection`、`ModuleBindingPanel`、`ModuleDecisionEntryPanel` 应默认先归属于 `ModuleDetailPage`
- **AND** `ModuleCreateForm` 默认先归属于 `ModuleCreatePage`
- **AND** `ReleaseCreateForm` 默认先归属于 `ReleaseCreatePage`

#### Scenario: 共享组件边界

- **WHEN** 后续实现讨论共享组件
- **THEN** 只有在列表页与详情页或多个页面确实共享同一职责时，才允许抽为共享组件
- **AND** 当前阶段不得为了“组件纯洁”提前拆出无明确复用证据的通用组件层

### Requirement: 布局降级策略必须冻结到页面结构层

系统 SHALL 将 `PC / 移动浏览器` 双场景下的布局降级策略冻结到页面结构层，确保同一套 `React Web` 页面在两种场景下都可直接实现。

#### Scenario: PC 页面布局

- **WHEN** 页面在桌面端实现
- **THEN** `ModuleListPage` 应优先采用高信息密度列表布局
- **AND** `ModuleDetailPage` 应优先采用分区式详情布局，使摘要、版本与关联入口可同时可见
- **AND** 不得因为桌面端实现而产生第二套专用页面体系

#### Scenario: PC 详情页布局分区

- **WHEN** `ModuleDetailPage` 在桌面端实现
- **THEN** 页面至少应分为摘要主区、版本列表区、关联动作区、`Decision` 入口区
- **AND** 版本与关联面板应优先保持同页可见，而不是依赖额外导航切换

#### Scenario: 移动浏览器页面布局

- **WHEN** 页面在移动浏览器实现
- **THEN** `ModuleListPage` 应采用单列列表或卡片重排
- **AND** `ModuleDetailPage` 应将摘要、版本、关联、`Decision` 入口按垂直顺序重排
- **AND** 可对次级信息采用折叠或延迟展开
- **AND** 不得新增独立移动端页面体系

#### Scenario: 移动浏览器的创建与版本登记页

- **WHEN** `ModuleCreatePage` 或 `ReleaseCreatePage` 在移动浏览器实现
- **THEN** 表单区与动作区应采用单列垂直布局
- **AND** 主动作按钮必须在无需横向滚动的前提下可见
- **AND** 不得通过新增移动端专用页面来规避当前布局问题

#### Scenario: 当前阶段移动端边界

- **WHEN** 后续实现讨论移动端承接方式
- **THEN** 不得引入独立 `React Native` 客户端
- **AND** 不得把完整 `PWA` 能力写成当前阶段实现前提
- **AND** 必须继续遵守单一 `React Web` 同时覆盖 `PC` 与移动浏览器的交付策略

## MODIFIED Requirements

### Requirement: Module Detail 前端承接方式

`Module Detail` 在当前阶段 SHALL 被解释为一个复合详情页，并在同一页面壳层中统一承接读取、版本入口、绑定动作与 `Decision` 入口。

#### Scenario: 详情页结构收口

- **WHEN** 后续 `/spec` 或实现讨论 `Module Detail` 的前端组织方式
- **THEN** 必须保持在同一详情页壳层中整合版本、绑定与 `Decision` 入口
- **AND** 不得提前拆出独立子工作台或并列主页面

## REMOVED Requirements

### Requirement: 第二套移动端 UI 架构
**Reason**: 当前阶段已经冻结为单一 `React Web` 交付策略，并要求同时覆盖 `PC` 与移动浏览器 UI。
**Migration**: 若后续确需独立移动端客户端，必须在新的 `phase` 或 `audit` 流程中重新冻结，不在 `phase02-06` 当前设计中处理。
