# Phase04-06 前端状态模型与交互流设计 Spec

## Why

`phase04-01` 已冻结页面边界与信息区块，`phase04-02` 已冻结模板字段、状态语义、`statusFilter` 取值与"仓库列表不引入 `providerFilter`"决策，`phase04-03` 已冻结三类绑定关系语义、候选范围、上下文入口、canonical owner 与绑定成功后 reread 承接页面，`phase04-04` 已冻结数据读写范围、接口边界与错误语义前提，`phase04-05` 已冻结页面文件落点、路由树、组件树、上下文承接路由参数与布局降级策略。但"页面怎么放"还不等于"页面如何流转"。如果不继续把列表查询条件承接、创建页草稿与提交状态、详情页候选读取与绑定动作状态、多入口来源上下文、返回路径与刷新恢复规则收紧，后续编码仍会在 `TanStack Router / TanStack Query / Zustand` 的边界上出现临场发挥，且 `Product List` 与 `Repository Binding / List` 必须承接 `phase02-09` §7.4/§8.4 与 `phase03-06` 已建立的"查询条件冻结到路由搜索参数层、返回列表时保留原搜索参数上下文"单值模式，不得引入 `sessionStorage` 等持久化层作为缺参回退源。

> 阶段分工约束：本规格只冻结前端页面级状态模型、交互流、路由搜索参数与页面状态之间的承接策略、来源上下文恢复规则与默认返回路径。后端模块边界、接口分组与 `.proto` 合同设计分别由 `phase04-07` 与 `phase04-08` 承接，不在本规格中提前冻结。具体 hook 命名、Query key、store API 命名、缓存时间、请求取消与 optimistic update 方案不在本规格中冻结。
>
> 与 `phase04-05` 边界划分：`phase04-05` 已冻结上下文承接路由参数的命名与传递规则（`moduleId / moduleName / fromModuleDetail`、`productId / productName / fromProductDetail`、目标页身份参数 `productId / repositoryId`）。本规格承接其路由参数命名，在此基础上冻结这些参数的状态语义、交互流转、查询条件承接策略与返回路径恢复规则，不重新发明第二套来源标记命名。
>
> 文档对齐说明：依据 Context7 当前文档，`TanStack Router` 的搜索参数默认值与校验应由 route 级定义承接，列表上下文适合以 URL search 作为可刷新、可返回、可前进后退的事实源；`Zustand` 当前文档强调派生状态优先通过 selector 计算、局部状态优先保留在页面或局部 store。本规格只冻结这些结论在产品行为上的语义，不冻结具体实现 API。

## What Changes

- 冻结 `Product Registry / List` 与 `Repository Binding / List` 的查询状态、筛选状态、空状态与刷新恢复规则
- 冻结列表查询条件在路由搜索参数与页面状态之间的承接策略，含 `fromList` 来源标记与返回列表时保留原搜索参数上下文的规则
- 冻结 `Product Create` 与 `Repository Create` 的来源上下文、草稿状态、提交状态、失败保持与成功回流规则
- 冻结 `Product Detail` 与 `Repository Binding Detail / Workspace` 的详情读取状态、候选读取状态、绑定动作状态与成功后 reread 规则
- 冻结 `Module Detail`、`Product Detail`、`Repository Binding / List` 多入口进入正式主线后的来源标记、默认返回路径与上下文恢复规则
- 承接 `phase04-02` 与 `phase04-04` 已冻结的筛选维度决策（两个列表均只冻结 `queryText / statusFilter`，不引入 `providerFilter`）
- 明确当前阶段只冻结用户可见状态与交互语义，不把 hook、store 或缓存实现写成既成事实

## Impact

- Affected specs: `phase04_product_and_repository_binding_foundation`
- Affected code: 后续 `frontend/src/routes/products/*`、`frontend/src/routes/repositories/*` 的 `validateSearch`、来源上下文搜索参数、返回路径承接与刷新恢复逻辑
- Affected code: 后续 `frontend/src/features/product-registry/*` 与 `frontend/src/features/repository-binding/*` 的列表视图状态、返回列表时保留原搜索参数上下文的导航逻辑、创建页草稿管理、详情页绑定面板状态、候选读取状态与成功后 reread 行为

## ADDED Requirements

### Requirement: 列表查询条件承接策略冻结

系统 SHALL 将 `Product Registry / List` 与 `Repository Binding / List` 的列表查询条件在路由搜索参数与页面状态之间的承接策略冻结为单值结论，承接 `phase02-09` §7.4/§8.4 与 `phase03-06` 已建立的"查询条件冻结到路由搜索参数层、返回列表时保留原搜索参数上下文"模式。

#### Scenario: 判断列表查询条件的路由搜索参数承接

- **WHEN** 后续实现讨论 `Product List` 或 `Repository Binding / List` 的查询条件来源
- **THEN** `queryText` 与 `statusFilter` 必须冻结到路由搜索参数层，作为列表筛选的唯一事实源
- **AND** `queryText` 必须允许空字符串
- **AND** `statusFilter` 只允许 `all / active / archived`，承接 `phase04-02` 冻结结论
- **AND** `all` 只允许存在于 UI 与路由搜索参数层，不得写入持久化层
- **AND** 列表默认查询条件为 `queryText = 空`、`statusFilter = all`
- **AND** 当前阶段不引入排序、分页或其他筛选维度
- **AND** 路由搜索参数是列表查询条件的唯一事实源，不得将事实源切换为不可见的全局瞬时状态或 `sessionStorage` 等持久化层

#### Scenario: 判断 fromList 来源标记与返回列表的参数保留

- **WHEN** 用户从 `Product Create`、`Product Detail` 或其他非列表入口返回 `Product List`
- **AND** 用户从 `Repository Create`、`Repository Binding Detail / Workspace` 或其他非列表入口返回 `Repository Binding / List`
- **THEN** 必须通过 `fromList` 路由搜索参数显式建模"来源列表上下文存在/不存在"
- **AND** 当 `fromList` 标记来源列表上下文存在时，返回列表的导航必须携带原 `queryText` 与 `statusFilter` 作为路由搜索参数，承接 `phase02-09` §7.4"保留原搜索参数上下文"模式
- **AND** 当 `fromList` 标记来源列表上下文不存在（如直链直达、外部链接、从 `Module Detail` / `Product Detail` 等非列表入口进入后再返回列表）时，返回列表必须落到默认筛选参数（`queryText = 空`、`statusFilter = all`）
- **AND** 不得在无来源列表上下文时伪造或套用上一份筛选条件
- **AND** `fromList` 只作为来源标记参与返回导航的参数拼接判定，不得作为独立恢复机制绕过路由搜索参数

#### Scenario: 判断列表刷新行为

- **WHEN** 用户在 `Product List` 或 `Repository Binding / List` 页面直接刷新
- **THEN** 必须以当前路由搜索参数中的 `queryText` 与 `statusFilter` 作为查询条件
- **AND** 若路由搜索参数缺失，必须使用默认筛选参数（`queryText = 空`、`statusFilter = all`）
- **AND** 不得在刷新后回退到任何持久化层中上一次的筛选值
- **AND** 无参 URL（如 `/products`、`/repositories`）必须稳定表现为默认筛选参数，不得出现不确定行为

### Requirement: 列表查询状态与空状态模型冻结

系统 SHALL 将 `Product Registry / List` 与 `Repository Binding / List` 的读取状态、派生视图状态与空状态模型冻结为单值结论，承接 `phase02-09` §8.4 已建立的模式。

#### Scenario: 判断列表读取状态

- **WHEN** 后续实现讨论 `Product List` 或 `Repository Binding / List` 的读取状态
- **THEN** 读取状态必须为 `pending / success / error`
- **AND** 派生视图状态必须为 `initial-loading / ready / empty / error`
- **AND** `ready` 与 `empty` 必须由成功读取后的数据是否为空派生
- **AND** 不得在当前阶段扩写为复杂缓存同步状态机

#### Scenario: 判断 Product List 空状态

- **WHEN** `ProductListRead` 成功返回零条记录
- **THEN** `ProductListPage` 必须进入 `empty`
- **AND** 空状态主动作必须直接进入 `ProductCreatePage`
- **AND** 不得把导入、自动扫描或 AI 建议写成空状态主入口

#### Scenario: 判断 Repository Binding / List 空状态

- **WHEN** `RepositoryListRead` 成功返回零条记录
- **THEN** `RepositoryBindingListPage` 必须进入 `empty`
- **AND** 空状态主动作必须直接进入 `RepositoryCreatePage`
- **AND** 不得把导入、自动扫描或 AI 建议写成空状态主入口

#### Scenario: 判断列表错误呈现位置

- **WHEN** 任一列表读取失败
- **THEN** 错误反馈必须停留在当前列表页内容区域上下文内
- **AND** 不得跳转到独立错误页

### Requirement: 列表筛选维度冻结

系统 SHALL 将 `Product List` 与 `Repository Binding / List` 的筛选维度冻结为单值结论，承接 `phase04-02` 与 `phase04-04` 已冻结的筛选决策。

#### Scenario: 判断 Product List 筛选维度

- **WHEN** 后续实现讨论 `Product List` 的筛选维度
- **THEN** 工具栏筛选维度只冻结 `queryText / statusFilter`
- **AND** `queryText` 只匹配 `name` 字段（模糊匹配），承接 `phase04-04`
- **AND** 不得引入超出 `queryText / statusFilter` 的筛选维度

#### Scenario: 判断 Repository Binding / List 筛选维度

- **WHEN** 后续实现讨论 `Repository Binding / List` 的筛选维度
- **THEN** 工具栏筛选维度只冻结 `queryText / statusFilter`
- **AND** 不得在当前阶段引入 `providerFilter`
- **AND** 承接 `phase04-02` "Repository 列表不引入 providerFilter"与 `phase04-04` 列表读取筛选参数冻结结论
- **AND** 若后续确需新增 `providerFilter`，必须进入新的冻结任务重新单值化

### Requirement: Product Create 交互状态流转冻结

系统 SHALL 将 `Product Create` 的前端状态拆分为"来源上下文 + 表单草稿状态 + 提交状态"，承接 `phase02-09` §8.4 Create 状态模式。

#### Scenario: 判断 Product Create 来源上下文

- **WHEN** 用户进入 `ProductCreatePage`
- **THEN** 来源上下文必须由路由搜索参数派生，只允许以下三种之一：
  - `fromList` 存在 → 来自 `Product List`，承接 `queryText / statusFilter`（经路由搜索参数保留）
  - `fromModuleDetail` 存在 → 来自 `Module Detail`，承接 `moduleId / moduleName`（承接 `phase04-05` 路由参数）
  - 无来源参数 → `direct-entry`
- **AND** 不得同时并行持有两个主来源上下文
- **AND** 不得发明第二套来源标记命名

#### Scenario: 判断 Product Create 草稿与提交状态

- **WHEN** 用户在 `Product Create` 录入或提交 `name / description / status`
- **THEN** 草稿状态必须为 `idle / dirty`
- **AND** `status` 默认预填 `active`，承接 `phase04-02` 显式提交语义
- **AND** 提交状态必须为 `submitting / submit-success / submit-error`
- **AND** 提交状态不得反向覆盖用户草稿值

#### Scenario: 判断 Product Create 提交失败处理

- **WHEN** `CreateProduct` 提交返回校验失败或 `submit-error`
- **THEN** 必须停留当前 `ProductCreatePage`
- **AND** 已输入草稿必须原样保留
- **AND** 来源上下文必须继续保留
- **AND** 错误必须显示在表单上下文，不得跳转独立错误页
- **AND** 不得降级为模糊通用错误，承接 `phase04-04` 创建失败错误语义

#### Scenario: 判断 Product Create 提交成功回流

- **WHEN** `CreateProduct` 提交成功
- **THEN** 默认回流路径必须进入新建 `Product` 对应的 `ProductDetailPage`
- **AND** 回流时必须继续携带创建页已有的来源标记与必要上下文参数，承接 `phase03-06` 创建成功后来源上下文继承模式：
  - `fromList` 存在 → 必须继续携带 `fromList` 与原 `queryText / statusFilter`
  - `fromModuleDetail` 存在 → 必须继续携带 `fromModuleDetail` 与原 `moduleId / moduleName`
  - 无来源参数 → `direct-entry`，不携带来源标记
- **AND** 回流后 `ProductDetailPage` 必须承接新建产品读取
- **AND** 后续用户返回时，`ProductDetailPage` 必须能基于继承的来源标记按真实来源返回
- **AND** 不得并列保留"成功后默认回列表"的第二套路径
- **AND** 不得在创建成功回流后丢失来源标记导致退化为 `direct-entry`

#### Scenario: 判断 Product Create 取消返回

- **WHEN** 用户在 `Product Create` 主动取消或返回
- **THEN** 默认返回路径必须按真实来源决定，不得统一伪造成回列表路径：
  - `fromList` 存在 → 返回 `Product List`，并携带原 `queryText` 与 `statusFilter` 作为路由搜索参数，承接 `phase02-09` §7.4
  - `fromModuleDetail` 存在 → 返回原 `ModuleDetailPage`
  - 无来源参数 → `direct-entry`，返回 `Product List` 默认筛选参数
- **AND** 不得要求用户重新输入筛选条件
- **AND** 不得在非列表来源下伪造 `fromList` 标记或回列表路径

### Requirement: Repository Create 交互状态流转冻结

系统 SHALL 将 `Repository Create` 的前端状态拆分为"来源上下文 + 表单草稿状态 + 提交状态"，承接 `phase02-09` §8.4 Create 状态模式。

#### Scenario: 判断 Repository Create 来源上下文

- **WHEN** 用户进入 `RepositoryCreatePage`
- **THEN** 来源上下文必须由路由搜索参数派生，只允许以下四种之一：
  - `fromList` 存在 → 来自 `Repository Binding / List`，承接 `queryText / statusFilter`（经路由搜索参数保留）
  - `fromProductDetail` 存在 → 来自 `Product Detail`，承接 `productId / productName`（承接 `phase04-05` 路由参数）
  - `fromModuleDetail` 存在 → 来自 `Module Detail`，承接 `moduleId / moduleName`（承接 `phase04-05` 路由参数）
  - 无来源参数 → `direct-entry`
- **AND** 不得同时并行持有两个主来源上下文
- **AND** 不得发明第二套来源标记命名

#### Scenario: 判断 Repository Create 草稿与提交状态

- **WHEN** 用户在 `Repository Create` 录入或提交 `name / url / provider / status`
- **THEN** 草稿状态必须为 `idle / dirty`
- **AND** `status` 默认预填 `active`
- **AND** 提交状态必须为 `submitting / submit-success / submit-error`
- **AND** 提交状态不得反向覆盖用户草稿值

#### Scenario: 判断 Repository Create 提交失败处理

- **WHEN** `CreateRepository` 提交返回校验失败或 `submit-error`
- **THEN** 必须停留当前 `RepositoryCreatePage`
- **AND** 已输入草稿必须原样保留
- **AND** 来源上下文必须继续保留
- **AND** 错误必须显示在表单上下文，不得跳转独立错误页
- **AND** 不得降级为模糊通用错误，承接 `phase04-04` 创建失败错误语义

#### Scenario: 判断 Repository Create 提交成功回流

- **WHEN** `CreateRepository` 提交成功
- **THEN** 默认回流路径必须进入新建 `Repository` 对应的 `RepositoryBindingDetailPage`
- **AND** 回流时必须继续携带创建页已有的来源标记与必要上下文参数，承接 `phase03-06` 创建成功后来源上下文继承模式：
  - `fromList` 存在 → 必须继续携带 `fromList` 与原 `queryText / statusFilter`
  - `fromProductDetail` 存在 → 必须继续携带 `fromProductDetail` 与原 `productId / productName`
  - `fromModuleDetail` 存在 → 必须继续携带 `fromModuleDetail` 与原 `moduleId / moduleName`
  - 无来源参数 → `direct-entry`，不携带来源标记
- **AND** 回流后 `RepositoryBindingDetailPage` 必须承接新建仓库读取
- **AND** 后续用户返回时，`RepositoryBindingDetailPage` 必须能基于继承的来源标记按真实来源返回
- **AND** 不得并列保留"成功后默认回列表"的第二套路径
- **AND** 不得在创建成功回流后丢失来源标记导致退化为 `direct-entry`

#### Scenario: 判断 Repository Create 取消返回

- **WHEN** 用户在 `Repository Create` 主动取消或返回
- **THEN** 默认返回路径必须按真实来源决定，不得统一伪造成回列表路径：
  - `fromList` 存在 → 返回 `Repository Binding / List`，并携带原 `queryText` 与 `statusFilter` 作为路由搜索参数，承接 `phase02-09` §7.4
  - `fromProductDetail` 存在 → 返回原 `ProductDetailPage`
  - `fromModuleDetail` 存在 → 返回原 `ModuleDetailPage`
  - 无来源参数 → `direct-entry`，返回 `Repository Binding / List` 默认筛选参数
- **AND** 不得要求用户重新输入筛选条件
- **AND** 不得在非列表来源下伪造 `fromList` 标记或回列表路径

### Requirement: Product Detail 交互状态流转冻结

系统 SHALL 将 `Product Detail` 的详情读取状态与 `BindModuleToProduct` 绑定面板交互状态流转冻结为单值结论，承接 `phase04-03` canonical owner 与 reread 决策。

#### Scenario: 判断 Product Detail 读取状态

- **WHEN** 用户进入 `ProductDetailPage`
- **THEN** 详情读取状态必须为 `pending / success / error`
- **AND** 资源不存在时必须派生 `not-found` 视图状态，承接 `phase04-04` 目标不存在错误语义
- **AND** 错误反馈必须停留在详情页内容区域，不得跳转独立错误页

#### Scenario: 判断 Product Detail 来源上下文

- **WHEN** 用户进入 `ProductDetailPage`
- **THEN** 来源上下文必须由路由搜索参数派生，只允许以下三种之一：
  - `fromList` 存在 → 来自 `Product List`，承接 `queryText / statusFilter`（经路由搜索参数保留）
  - `fromModuleDetail` 存在 → 来自 `Module Detail`，承接 `moduleId / moduleName`（承接 `phase04-05` 路由参数）
  - 无来源参数 → `direct-entry`
- **AND** 从 `ProductCreatePage` 成功创建后进入时，来源上下文必须继承自创建页，承接 `phase03-06` 创建成功后来源上下文继承模式
- **AND** 不得同时并行持有两个主来源上下文

#### Scenario: 判断 ProductModuleBindingPanel 候选读取状态

- **WHEN** 用户在 `ProductDetailPage` 打开 `ProductModuleBindingPanel`
- **THEN** 候选读取最小状态必须为 `closed / pending / ready / empty / error`
- **AND** `empty` 必须表示候选列表为空，而不是接口失败或资源不存在
- **AND** 候选 `Module` 读取必须独立于详情读取，承接 `phase04-04` 候选读取独立性

#### Scenario: 判断 ProductModuleBindingPanel 写入状态

- **WHEN** 用户在 `ProductDetailPage` 提交 `BindModuleToProduct`
- **THEN** 写入状态必须为 `idle / submitting / submit-success / submit-error`
- **AND** 该状态只归属于当前详情页上下文

#### Scenario: 判断 BindModuleToProduct 提交成功 reread

- **WHEN** `BindModuleToProduct` 提交成功
- **THEN** 用户必须停留在当前 `ProductDetailPage`
- **AND** 必须重新读取已绑定 `Module` 列表完成 reread，承接 `phase04-03` reread 承接页面冻结结论
- **AND** 不得只靠 `toast` 或局部通知作为成功依据
- **AND** `ProductModuleBindingPanel` 必须回到 `closed`

#### Scenario: 判断 BindModuleToProduct 候选为空状态

- **WHEN** `ProductModuleCandidateRead` 返回零条候选 `Module`
- **THEN** 面板必须展示明确的无可绑定候选空状态提示
- **AND** 不得把空候选结果误报为接口错误，承接 `phase04-03` 与 `phase04-04` 候选空结果语义

#### Scenario: 判断 BindModuleToProduct 提交失败处理

- **WHEN** `BindModuleToProduct` 提交返回重复冲突、目标不存在或 `submit-error`
- **THEN** 必须停留在当前 `ProductDetailPage`
- **AND** 当前已选候选 `Module` 必须继续保留
- **AND** 错误必须停留在面板上下文，不得跳转独立错误页
- **AND** 重复绑定不得降级为静默成功，承接 `phase04-03` 与 `phase04-04` 重复绑定语义

### Requirement: Repository Binding Detail / Workspace 绑定工作台交互状态流转冻结

系统 SHALL 将 `Repository Binding Detail / Workspace` 的详情读取状态、`BindRepositoryToProduct` 与 `MapModuleToRepository` 两类绑定面板交互状态流转冻结为单值结论，承接 `phase04-03` canonical owner 与 reread 决策。

#### Scenario: 判断 Repository Binding Detail 读取状态

- **WHEN** 用户进入 `RepositoryBindingDetailPage`
- **THEN** 详情读取状态必须为 `pending / success / error`
- **AND** 资源不存在时必须派生 `not-found` 视图状态，承接 `phase04-04` 目标不存在错误语义
- **AND** 错误反馈必须停留在工作台内容区域，不得跳转独立错误页

#### Scenario: 判断 Repository Binding Detail 来源上下文

- **WHEN** 用户进入 `RepositoryBindingDetailPage`
- **THEN** 来源上下文必须由路由搜索参数派生，只允许以下四种之一：
  - `fromList` 存在 → 来自 `Repository Binding / List`，承接 `queryText / statusFilter`（经路由搜索参数保留）
  - `fromProductDetail` 存在 → 来自 `Product Detail`，承接 `productId / productName`（承接 `phase04-05` 路由参数）
  - `fromModuleDetail` 存在 → 来自 `Module Detail`，承接 `moduleId / moduleName`（承接 `phase04-05` 路由参数）
  - 无来源参数 → `direct-entry`
- **AND** 从 `RepositoryCreatePage` 成功创建后进入时，来源上下文必须继承自创建页，承接 `phase03-06` 创建成功后来源上下文继承模式
- **AND** 不得同时并行持有两个主来源上下文

#### Scenario: 判断两类绑定面板状态与互斥展开

- **WHEN** 用户在 `RepositoryBindingDetailPage` 发起 `BindRepositoryToProduct` 或 `MapModuleToRepository`
- **THEN** `RepositoryProductBindingPanel` 候选读取状态必须为 `closed / pending / ready / empty / error`，写入状态必须为 `idle / submitting / submit-success / submit-error`
- **AND** `RepositoryModuleMappingPanel` 候选读取状态必须为 `closed / pending / ready / empty / error`，写入状态必须为 `idle / submitting / submit-success / submit-error`
- **AND** 同一时刻只允许一个绑定面板处于打开态（互斥展开）
- **AND** 候选 `Product` 与候选 `Module` 读取必须各自独立于详情读取，承接 `phase04-04` 候选读取独立性

#### Scenario: 判断两类绑定候选为空状态

- **WHEN** `ProductBindingCandidateRead` 返回零条候选 `Product`
- **AND** `RepositoryModuleCandidateRead` 返回零条候选 `Module`
- **THEN** 对应面板必须展示明确的无可绑定候选空状态提示
- **AND** 不得把空候选结果误报为接口错误，承接 `phase04-03` 与 `phase04-04` 候选空结果语义

#### Scenario: 判断 BindRepositoryToProduct 提交成功 reread

- **WHEN** `BindRepositoryToProduct` 提交成功
- **THEN** 用户必须停留在当前 `RepositoryBindingDetailPage`
- **AND** 必须重新读取已绑定 `Product` 列表完成 reread，承接 `phase04-03` reread 承接页面冻结结论
- **AND** 不得只靠 `toast` 作为成功依据
- **AND** 当前活动面板必须回到 `closed`

#### Scenario: 判断 MapModuleToRepository 提交成功 reread

- **WHEN** `MapModuleToRepository` 提交成功
- **THEN** 用户必须停留在当前 `RepositoryBindingDetailPage`
- **AND** 必须重新读取已映射 `Module` 列表完成 reread，承接 `phase04-03` reread 承接页面冻结结论
- **AND** 不得只靠 `toast` 作为成功依据
- **AND** 当前活动面板必须回到 `closed`

#### Scenario: 判断两类绑定提交失败处理

- **WHEN** `BindRepositoryToProduct` 或 `MapModuleToRepository` 提交返回重复冲突、目标不存在或 `submit-error`
- **THEN** 必须停留在当前 `RepositoryBindingDetailPage`
- **AND** 当前已选候选目标必须继续保留
- **AND** 错误必须停留在对应面板上下文，不得跳转独立错误页
- **AND** 重复绑定不得降级为静默成功，承接 `phase04-03` 与 `phase04-04` 重复绑定语义

### Requirement: 多入口返回路径与上下文恢复规则冻结

系统 SHALL 将从 `Module Detail`、`Product Detail`、`Repository Binding / List` 多入口进入正式绑定主入口时的来源标记、返回路径与默认回流页面冻结为单值结论，承接 `phase04-03` canonical owner 与 reread 决策、`phase04-05` 上下文承接路由参数。

> 关键区分：绑定写入成功后的行为是停留在 canonical owner 页面完成 reread（承接 `phase04-03`），与用户主动发起的"返回来源页"是两个独立动作。本规格冻结的"默认返回路径"指用户主动返回时的目标页，不影响绑定成功后必须先完成 reread 的约束。

#### Scenario: 判断从 Module Detail 兼容入口进入后的来源标记与 reread

- **WHEN** 用户从 `Module Detail` 携带 `moduleId / moduleName / fromModuleDetail` 跳入 `Product Detail` 或 `Repository Binding Detail / Workspace`
- **THEN** 来源标记 `fromModuleDetail` 必须保留在路由搜索参数中，承接 `phase04-05`
- **AND** 接收方页面必须能基于上下文参数预填对应绑定面板的候选 `Module` 选择
- **AND** 绑定写入成功后必须停留在 canonical owner 页面完成 reread，承接 `phase04-03`，不得回流回 `Module Detail`
- **AND** 用户主动返回时的默认返回路径为原 `ModuleDetailPage`
- **AND** 不得在 `Module Detail` 内直接提交绑定写入

#### Scenario: 判断从 Product Detail 上下文入口进入后的来源标记与 reread

- **WHEN** 用户从 `Product Detail` 携带 `productId / productName / fromProductDetail` 跳入 `Repository Binding Detail / Workspace`
- **THEN** 来源标记 `fromProductDetail` 必须保留在路由搜索参数中，承接 `phase04-05`
- **AND** `RepositoryBindingDetailPage` 必须能基于上下文参数预填 `RepositoryProductBindingPanel` 的候选 `Product` 选择
- **AND** `BindRepositoryToProduct` 写入成功后必须停留在 `Repository Binding Detail / Workspace` 完成 reread，承接 `phase04-03`，不得回流回 `Product Detail`
- **AND** 用户主动返回时的默认返回路径为原 `ProductDetailPage`
- **AND** `Product Detail` 自身不得承接第二套仓库绑定写入流程

#### Scenario: 判断从 Product List 进入 Detail 后返回 List 的上下文恢复

- **WHEN** 用户从 `Product List` 进入 `Product Detail` 后主动返回列表
- **THEN** 必须携带 `fromList` 标记返回 `Product List`
- **AND** 返回导航必须携带原 `queryText` 与 `statusFilter` 作为路由搜索参数，承接 `phase02-09` §7.4
- **AND** 不得要求用户重新输入筛选条件

#### Scenario: 判断从 Repository Binding / List 进入 Detail 后返回 List 的上下文恢复

- **WHEN** 用户从 `Repository Binding / List` 进入 `Repository Binding Detail / Workspace` 后主动返回列表
- **THEN** 必须携带 `fromList` 标记返回 `Repository Binding / List`
- **AND** 返回导航必须携带原 `queryText` 与 `statusFilter` 作为路由搜索参数，承接 `phase02-09` §7.4
- **AND** 不得要求用户重新输入筛选条件

#### Scenario: 判断从创建成功后的 Detail 页返回来源页

- **WHEN** 用户从 `ProductCreatePage` 成功进入 `ProductDetailPage` 后主动返回
- **AND** 用户从 `RepositoryCreatePage` 成功进入 `RepositoryBindingDetailPage` 后主动返回
- **THEN** `Detail` 页必须基于从创建页继承的来源标记按真实来源返回，承接 `phase03-06` 创建成功后来源上下文继承模式
- **AND** 若来源为 `fromList`，返回列表必须携带原 `queryText` 与 `statusFilter` 作为路由搜索参数
- **AND** 若来源为 `fromModuleDetail`，返回原 `ModuleDetailPage`
- **AND** 若来源为 `fromProductDetail`，返回原 `ProductDetailPage`
- **AND** 若来源为 `direct-entry`，返回列表必须落到默认筛选参数
- **AND** 不得在创建成功回流后丢失来源标记导致退化为 `direct-entry`

#### Scenario: 判断无来源列表上下文时返回列表的默认行为

- **WHEN** 用户从外部链接或刷新后直接进入任一创建页或详情页，且不存在来源列表上下文
- **THEN** 返回列表必须落到默认列表参数（`queryText = 空`、`statusFilter = all`）
- **AND** 不得伪造上一份列表筛选条件
- **AND** 不得回退到任何持久化层中上一次的筛选值

#### Scenario: 判断来源上下文刷新恢复

- **WHEN** 用户刷新带有 `fromModuleDetail` 或 `fromProductDetail` 的正式主线页面
- **THEN** 页面必须继续恢复对应来源标记
- **AND** 不得在刷新后静默丢失来源标记

### Requirement: 页面级 UI 状态局部归属原则冻结

系统 SHALL 将当前阶段的局部 UI 状态优先冻结为页面级或详情页上下文归属，只在确有同页多组件共享需要时再抽为页面作用域状态容器。

#### Scenario: 判断列表与创建页局部状态归属

- **WHEN** 后续实现 `ProductListPage`、`RepositoryBindingListPage`、`ProductCreatePage` 与 `RepositoryCreatePage`
- **THEN** 搜索输入草稿、表单草稿、提交错误与空状态展示应优先归属于当前页面
- **AND** 不得默认升级为跨路由全局状态

#### Scenario: 判断详情页局部状态归属

- **WHEN** 后续实现 `ProductDetailPage` 与 `RepositoryBindingDetailPage`
- **THEN** 候选读取状态、当前选中候选、活动面板状态、提交错误与来源上下文展示应优先归属于当前详情页上下文
- **AND** 派生视图状态应优先由当前读模型结果计算，而不是被重复持久化为独立全局字段

### Requirement: 运行时实现细节不冻结

系统 SHALL 明确当前阶段只冻结状态语义与交互流转，不冻结运行时实现细节。

#### Scenario: 判断当前阶段允许冻结的内容

- **WHEN** 当前 spec 讨论页面状态与交互流
- **THEN** 可以冻结状态名称、状态归属、默认返回路径、来源上下文、刷新恢复规则、错误呈现位置与成功后的 reread 预期

#### Scenario: 判断当前阶段不得冻结的内容

- **WHEN** 后续实现尚未开始
- **THEN** 不得提前冻结具体 hook 命名、Query key 细节、store API 命名、缓存时间、请求取消策略或 optimistic update 方案
- **AND** 不得引入 `sessionStorage` 或其他持久化层作为列表查询条件的事实源或缺参回退源

## MODIFIED Requirements

### Requirement: phase04-05 上下文承接路由参数的状态语义解释

`phase04-05` 已冻结的上下文承接路由参数（`moduleId / moduleName / fromModuleDetail`、`productId / productName / fromProductDetail`、目标页身份参数）在 `phase04-06` 中 SHALL 被进一步解释为带有来源标记与上下文恢复语义的状态入口，而不只是命名传递规则。

#### Scenario: 上下文承接路由参数的状态语义解释

- **WHEN** 后续实现讨论 `phase04-05` 冻结的上下文承接路由参数
- **THEN** 必须将其理解为携带来源标记与预填语义的状态入口
- **AND** 来源标记（`fromModuleDetail / fromProductDetail`）必须参与返回路径与默认回流页面判定
- **AND** `fromList` 标记必须参与返回列表时的搜索参数拼接判定（存在则保留原参数，不存在则落默认参数），不得作为独立恢复机制绕过路由搜索参数
- **AND** 不得把上下文承接参数只解释为无状态的 URL 透传

### Requirement: phase04-03 canonical owner 与 reread 的前端状态承接解释

`phase04-03` 已冻结的 canonical owner 与绑定成功后 reread 承接页面在 `phase04-06` 中 SHALL 被进一步解释为前端状态流转的终态约束，即绑定写入成功后必须停留在 canonical owner 页面并触发对应已绑定列表的重新读取。

#### Scenario: canonical owner 与 reread 的前端状态承接解释

- **WHEN** 后续实现讨论三类绑定动作成功后的前端行为
- **THEN** 必须将"回到 canonical owner 完成 reread"实现为"停留当前 canonical owner 页面 + 重新读取对应已绑定列表 + 面板回到 `closed`"
- **AND** 不得把 reread 实现为跳转离开 canonical owner 或只靠 `toast` 提示

## REMOVED Requirements

### Requirement: providerFilter 作为当前阶段筛选维度

**Reason**: `phase04-02` 已冻结"Repository 列表不引入 providerFilter"，`phase04-04` 再次冻结列表读取筛选参数为 `queryText / statusFilter`。本规格承接该结论，不在当前阶段引入 `providerFilter`。

**Migration**: 若后续确需新增 `providerFilter`，必须进入新的冻结任务重新单值化，不在 `phase04-06` 当前规格中处理。

### Requirement: 并列默认回流路径

**Reason**: 当前阶段需要单值交互流，避免实现阶段为同一动作保留两套默认返回行为。

**Migration**: 若后续确有多回流模式需求，必须在新的 `phase` 或 `audit` 中重新定义，不在当前阶段保留并列默认路径。
