# Phase02-07 前端状态模型与交互流设计 Spec

## Why

`phase02-06` 已经把 `Module Registry` 的页面、路由、页面壳层与组件树冻结到可直接映射实现的层级，但“页面怎么组织”仍不等于“页面如何流转”。如果不继续把状态归属、查询条件承接、表单提交回流、错误呈现位置与详情页重读规则收紧，后续编码仍会在 `TanStack Router / Query / Zustand` 的承接边界上出现临场发挥。因此本子任务需要产出一份更强的实现前设计结果，让 `phase02-10` 的前端实现可以直接照此落地。

## What Changes

- 将 `Module List` 的查询条件、读取状态与空状态收敛到“路由搜索参数 + 页面容器视图状态”层
- 将 `Module Create`、`Release Create` 的草稿状态、提交状态、错误态与成功回流收敛到页面表单层
- 将 `Module Detail` 内绑定动作的状态归属、单面板规则、成功后重读与错误呈现位置收敛到详情页上下文
- 将列表、创建、详情、版本登记之间的返回路径与上下文保留规则冻结为单值设计
- 明确当前阶段只冻结用户可见状态与交互语义，不把 hook 命名、缓存实现或 store 细节写成既成事实

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `frontend/` 中 `Module Registry` 的列表查询承接、表单草稿管理、Mutation 成功回流、详情页重读与面板开闭状态
- Affected code: 预期会直接影响 `frontend/src/routes/modules/*`、`frontend/src/features/module-registry/pages/*`、`frontend/src/features/module-registry/components/*`

## ADDED Requirements

### Requirement: Module List 查询条件必须冻结到路由搜索参数层

系统 SHALL 将 `Module List` 的用户可编辑查询条件冻结到路由搜索参数层，使列表上下文可被返回、分享与浏览器前进后退正确承接。

#### Scenario: 列表最小搜索参数

- **WHEN** 后续实现 `ModuleListRoute` 与 `ModuleListPage`
- **THEN** 列表最小搜索参数必须至少包含 `queryText` 与 `statusFilter`
- **AND** 当前阶段不提前引入排序、分页、多条件组合筛选

#### Scenario: 从列表进入其他页面再返回

- **WHEN** 用户从 `ModuleListPage` 进入 `ModuleCreatePage` 或 `ModuleDetailPage` 后再返回
- **THEN** 已确认的 `queryText` 与 `statusFilter` 必须继续保留
- **AND** 不得把列表筛选状态藏入不可见的全局瞬时状态

### Requirement: Module List 视图状态必须由读取结果派生

系统 SHALL 将 `Module List` 的页面容器状态冻结为“读取状态 + 派生视图状态”的组合，而不是混写成一组彼此重叠的临时布尔值。

#### Scenario: 列表最小读取状态

- **WHEN** 后续实现列表读取
- **THEN** 页面容器必须能够区分 `pending`、`success`、`error` 三种最小读取结果
- **AND** 不得在当前阶段扩写为复杂缓存同步状态机

#### Scenario: 列表最小视图状态

- **WHEN** 页面根据读取结果决定展示内容
- **THEN** 页面最小视图状态必须冻结为 `initial-loading`、`ready`、`empty`、`error`
- **AND** `ready` 与 `empty` 必须由成功读取后的数据是否为空派生

#### Scenario: 列表错误呈现位置

- **WHEN** 列表读取失败
- **THEN** 错误反馈必须停留在 `ModuleListPage` 的内容区域上下文内
- **AND** 不得跳转到独立错误页

#### Scenario: 列表空状态主动作

- **WHEN** 列表读取成功但没有任何模块
- **THEN** 页面必须进入 `empty`
- **AND** 空状态主动作必须直接进入 `ModuleCreatePage`
- **AND** 不得把导入、自动扫描或 AI 建议写成空状态主入口

### Requirement: Module Create 状态必须区分草稿状态与提交状态

系统 SHALL 将 `Module Create` 的前端状态拆分为“表单草稿状态”与“提交状态”，避免后续实现把本地输入与服务端写入状态混成一个黑箱。

#### Scenario: 创建页最小草稿状态

- **WHEN** 后续实现 `ModuleCreatePage`
- **THEN** 页面必须至少承接 `name`、`description`、`status` 三个最小字段草稿
- **AND** 草稿状态必须至少能区分 `idle` 与 `dirty`

#### Scenario: 创建页最小提交状态

- **WHEN** 用户提交 `CreateModule`
- **THEN** 页面必须至少承接 `submitting`、`submit-success`、`submit-error`
- **AND** 提交状态不得反向覆盖用户草稿值

#### Scenario: 创建失败后的状态保持

- **WHEN** `CreateModule` 提交失败
- **THEN** 页面必须停留在当前 `ModuleCreatePage`
- **AND** 已输入草稿必须原样保留
- **AND** 错误反馈必须显示在当前表单上下文中，而不是跳转或清空表单

#### Scenario: 创建成功后的默认回流

- **WHEN** `CreateModule` 提交成功
- **THEN** 默认回流路径必须进入新建模块对应的 `ModuleDetailPage`
- **AND** 不得并列保留“成功后回列表”的第二套默认路径

#### Scenario: 创建成功后的列表读取预期

- **WHEN** 用户稍后返回 `ModuleListPage`
- **THEN** 列表读取必须能够承接新建模块结果
- **AND** 当前阶段将其冻结为“成功写入后触发相关读模型重读”的用户可见语义

### Requirement: Release Create 状态必须依附当前模块路由上下文

系统 SHALL 将 `Release Create` 的状态模型冻结为依附 `moduleId` 路由上下文的版本登记流程，而不是一个脱离详情上下文的独立工作台。

#### Scenario: 版本登记上下文来源

- **WHEN** 后续实现 `ReleaseCreatePage`
- **THEN** 当前模块标识必须来自当前路由参数 `moduleId`
- **AND** 不得再复制一份可写的全局“当前模块”状态作为事实源

#### Scenario: 版本登记最小状态

- **WHEN** 用户在 `ReleaseCreatePage` 录入版本
- **THEN** 页面必须至少承接 `idle`、`dirty`、`submitting`、`submit-success`、`submit-error`

#### Scenario: 版本登记失败后的状态保持

- **WHEN** `CreateRelease` 提交失败
- **THEN** 页面必须停留在当前 `ReleaseCreatePage`
- **AND** 必须保留当前输入
- **AND** 不得跳出当前 `moduleId` 上下文

#### Scenario: 版本登记成功后的默认回流

- **WHEN** `CreateRelease` 提交成功
- **THEN** 默认回流路径必须返回当前模块的 `ModuleDetailPage`
- **AND** `ModuleDetailPage` 返回后必须承接最新版本列表读取

### Requirement: 绑定动作状态必须冻结到 Module Detail 上下文

系统 SHALL 将 `BindModuleToProduct` 与 `MapModuleToRepository` 的状态冻结到 `ModuleDetailPage` 上下文，而不是拆成独立页面或全局工作流。

#### Scenario: 绑定面板最小状态

- **WHEN** 后续实现 `ModuleBindingPanel`
- **THEN** 面板状态至少应包含 `closed`、`open-idle`、`submitting`、`submit-success`、`submit-error`
- **AND** 面板必须分别承接 `Product` 候选与 `Repository` 候选的选择上下文

#### Scenario: 同时只允许一个绑定面板处于打开态

- **WHEN** 用户在 `ModuleDetailPage` 中执行绑定动作
- **THEN** 同一时刻只允许一个绑定面板或一种绑定模式处于打开态
- **AND** 不得同时打开两个并行提交上下文

#### Scenario: 绑定成功后的详情页行为

- **WHEN** 任一绑定动作提交成功
- **THEN** 用户必须停留在当前 `ModuleDetailPage`
- **AND** 当前详情页必须重新读取或刷新对应绑定结果
- **AND** 不得把绑定成功回流设计成跳转到 `Product Registry` 或 `Repository Binding`

#### Scenario: 绑定失败后的错误呈现

- **WHEN** 任一绑定动作提交失败
- **THEN** 错误反馈必须停留在当前面板上下文内
- **AND** 面板必须保留当前选择，允许用户修正后重试

### Requirement: Module Detail 必须作为统一回流宿主页

系统 SHALL 将 `ModuleDetailPage` 冻结为当前阶段前端主线中的统一回流宿主页，承接创建完成后的落点、版本登记后的回流以及绑定动作后的停留上下文。

#### Scenario: 详情页作为动作回流宿主

- **WHEN** 创建成功后进入详情，或版本登记/绑定动作在详情上下文完成
- **THEN** `ModuleDetailPage` 必须承担统一回流宿主角色
- **AND** 后续实现不得再新增第二套默认宿主页

#### Scenario: Decision 入口的上下文保持

- **WHEN** 用户在 `ModuleDetailPage` 使用 `Decision` 入口
- **THEN** 当前阶段必须保持 `ModuleDetailPage` 为主上下文页
- **AND** 不得把当前交互改造成多页面往返工作流

### Requirement: 页面返回路径必须冻结为单值规则

系统 SHALL 将当前阶段主要页面之间的主动返回路径冻结为单值规则，避免后续实现为同一动作设计两套默认返回行为。

#### Scenario: 从创建页主动返回

- **WHEN** 用户主动取消或返回离开 `ModuleCreatePage`
- **THEN** 默认返回路径应为保留原搜索参数上下文的 `ModuleListPage`

#### Scenario: 从版本登记页主动返回

- **WHEN** 用户主动取消或返回离开 `ReleaseCreatePage`
- **THEN** 默认返回路径应为当前模块的 `ModuleDetailPage`

### Requirement: 页面级 UI 状态应优先保持局部归属

系统 SHALL 将当前阶段的局部 UI 状态优先冻结为页面级或组件级归属，只在确有跨同页多组件共享需要时再抽为页面作用域状态容器。

#### Scenario: 草稿与面板状态的归属

- **WHEN** 后续实现 `ModuleCreatePage`、`ReleaseCreatePage` 与 `ModuleBindingPanel`
- **THEN** 表单草稿、提交错误、面板开闭等瞬时 UI 状态应优先归属于当前页面或当前详情页上下文
- **AND** 不得默认升级为跨路由全局状态

### Requirement: 运行时实现细节不得写成当前既成事实

系统 SHALL 明确区分“状态与交互设计”与“具体实现手段”，避免当前 spec 过早冻结实现细节。

#### Scenario: 当前阶段允许冻结的内容

- **WHEN** 当前 spec 讨论页面状态与交互流
- **THEN** 可以冻结状态名称、状态归属、返回路径、错误呈现位置、成功后的读模型刷新预期与用户可见行为

#### Scenario: 当前阶段不得冻结的内容

- **WHEN** 后续实现尚未开始
- **THEN** 不得提前冻结具体 hook 命名、Query key 细节、store API 命名、缓存时间、请求取消策略或 optimistic update 方案

## MODIFIED Requirements

### Requirement: Module List 前端承接方式

`Module List` 在当前阶段不仅 SHALL 承接读取结果，还应统一承接列表搜索参数、派生视图状态与从其他页面返回后的上下文恢复。

#### Scenario: 列表页上下文收口

- **WHEN** 用户从创建页或详情页回到 `ModuleListPage`
- **THEN** 页面必须能够按原有搜索参数恢复列表上下文
- **AND** 不得要求用户重新输入筛选条件

## REMOVED Requirements

### Requirement: 并列默认回流路径
**Reason**: 当前阶段需要单值交互流，避免实现阶段为同一动作保留两套默认返回路径。
**Migration**: 若后续确有多回流模式需求，必须在新的 `phase` 或 `audit` 中重新定义，不在当前阶段保留并列默认路径。
