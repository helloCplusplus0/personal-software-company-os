# Phase03-06 前端状态模型与交互流设计 Spec

## Why

`phase03-05` 已经把 `Decision Center` 的页面、路由、文件落点与组件树冻结到可直接映射实现的层级，但“页面怎么摆”并不等于“页面如何流转”。如果不继续把列表查询承接、创建页草稿状态、详情页待关联目标承接、关联动作状态与返回路径恢复规则收紧，后续编码仍会在 `TanStack Router / Query / Zustand` 的边界上出现临场发挥。因此本子任务需要产出一份更强的实现前设计结果，让 `phase03-10` 的前端实现可以直接照此落地。

## What Changes

- 将 `Decision List` 的查询条件、读取状态与空状态收敛到“路由搜索参数 + 页面容器视图状态”层
- 将 `Decision Create` 的来源上下文、表单草稿状态、提交状态与成功回流规则收敛到页面级状态模型
- 将 `Decision Detail` 中详情读取、待关联目标展示、候选读取、关联写入的交互流冻结为单值设计
- 将列表、创建、详情、关联之间的返回路径与列表上下文恢复规则冻结为单值设计
- 明确当前阶段只冻结用户可见状态与交互语义，不把 hook 命名、缓存实现或 store 细节写成既成事实

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `frontend/` 中 `Decision Center` 的列表搜索参数承接、表单草稿管理、Mutation 成功回流、详情页候选读取与关联动作状态
- Affected code: 预期会直接影响 `frontend/src/routes/decisions/*`、`frontend/src/features/decision-center/pages/*`、`frontend/src/features/decision-center/components/*`

## ADDED Requirements

### Requirement: Decision List 查询条件必须冻结到路由搜索参数层

系统 SHALL 将 `Decision List` 的用户可编辑查询条件冻结到路由搜索参数层，使列表上下文可被返回、刷新与浏览器前进后退正确承接。

#### Scenario: 列表最小搜索参数

- **WHEN** 后续实现 `DecisionListRoute` 与 `DecisionListPage`
- **THEN** 列表最小搜索参数必须至少包含 `queryText` 与 `statusFilter`
- **AND** 当前阶段不提前引入排序、分页、多条件组合筛选

#### Scenario: 列表搜索参数默认来源

- **WHEN** `DecisionListPage` 首次根据当前 URL 初始化页面状态
- **THEN** `queryText` 与 `statusFilter` 的默认来源必须为当前路由搜索参数
- **AND** 不得将搜索参数的事实源切换为不可见的全局瞬时状态

#### Scenario: 列表搜索参数默认值

- **WHEN** `DecisionListRoute` 的 URL 中不存在 `queryText` 或 `statusFilter`
- **THEN** `queryText` 必须默认为空字符串（即不执行文本筛选）
- **AND** `statusFilter` 必须默认为"全部状态"（即不执行状态筛选）
- **AND** 不得将默认值解释为 `null` 或跳过列表读取

#### Scenario: 刷新后的列表恢复

- **WHEN** 用户刷新 `DecisionListRoute`
- **THEN** 若当前 URL 仍保留 `queryText` 与 `statusFilter`
- **AND** `DecisionListPage` 必须按该参数恢复读取状态与筛选结果

#### Scenario: 从列表进入其他页面再返回

- **WHEN** 用户从 `DecisionListPage` 进入 `DecisionCreatePage` 或 `DecisionDetailPage` 后再返回
- **THEN** 已确认的 `queryText` 与 `statusFilter` 必须继续保留
- **AND** 不得要求用户重新输入筛选条件

### Requirement: Decision List 视图状态必须由读取结果派生

系统 SHALL 将 `Decision List` 的页面容器状态冻结为“读取状态 + 派生视图状态”的组合，而不是混写成一组彼此重叠的临时布尔值。

#### Scenario: 列表最小读取状态

- **WHEN** 后续实现列表读取
- **THEN** 页面容器必须能够区分 `pending`、`success`、`error` 三种最小读取结果
- **AND** 不得在当前阶段扩写为复杂缓存同步状态机

#### Scenario: 列表最小视图状态

- **WHEN** 页面根据读取结果决定展示内容
- **THEN** 页面最小视图状态必须冻结为 `initial-loading`、`ready`、`empty`、`error`
- **AND** `ready` 与 `empty` 必须由成功读取后的数据是否为空派生

#### Scenario: 列表空状态主动作

- **WHEN** 列表读取成功但没有任何 `Decision`
- **THEN** 页面必须进入 `empty`
- **AND** 空状态主动作必须直接进入 `DecisionCreatePage`

#### Scenario: 列表错误呈现位置

- **WHEN** 列表读取失败
- **THEN** 错误反馈必须停留在 `DecisionListPage` 的内容区域上下文内
- **AND** 不得跳转到独立错误页

### Requirement: Decision Create 状态必须区分来源上下文、草稿状态与提交状态

系统 SHALL 将 `Decision Create` 的前端状态拆分为“来源上下文状态 + 表单草稿状态 + 提交状态”，避免后续实现把页面输入、入口上下文与服务端写入状态混成一个黑箱。

#### Scenario: 创建页来源上下文状态

- **WHEN** 用户从 `Module Detail` 带上下文进入 `DecisionCreatePage`
- **THEN** 页面必须能够承接“带来源上下文进入”的最小状态
- **AND** 该上下文至少包含来源 `Module` 的目标标识与可展示名称
- **AND** 该上下文只承接展示与后续回流，不等于已完成正式关联

#### Scenario: 创建页最小草稿状态

- **WHEN** 后续实现 `DecisionCreatePage`
- **THEN** 页面必须至少承接 `title`、`context`、`problem`、`alternatives`、`choice`、`reason`、`impact`、`status` 的最小草稿
- **AND** 草稿状态必须至少能区分 `idle` 与 `dirty`

#### Scenario: 创建页最小提交状态

- **WHEN** 用户提交 `RecordDecision`
- **THEN** 页面必须至少承接 `submitting`、`submit-success`、`submit-error`
- **AND** 提交状态不得反向覆盖用户草稿值

#### Scenario: 创建失败后的状态保持

- **WHEN** `RecordDecision` 提交失败
- **THEN** 页面必须停留在当前 `DecisionCreatePage`
- **AND** 已输入草稿必须原样保留
- **AND** 来源上下文展示必须继续保留
- **AND** 错误反馈必须显示在当前表单上下文中，而不是跳转或清空表单

#### Scenario: 创建成功后的默认回流

- **WHEN** `RecordDecision` 提交成功
- **THEN** 默认回流路径必须进入新建 `Decision` 对应的 `DecisionDetailPage`
- **AND** 不得并列保留“成功后回列表”的第二套默认路径

### Requirement: Decision Detail 必须统一承接读取、待关联目标与关联动作状态

系统 SHALL 将 `Decision Detail` 的前端状态冻结为依附当前 `decisionId` 路由上下文的复合详情流程，而不是拆成多个彼此松散的子状态容器。

#### Scenario: 详情页最小读取状态

- **WHEN** 后续实现 `DecisionDetailPage`
- **THEN** 页面必须至少承接 `pending`、`success`、`error`
- **AND** 读取成功后才能派生概要、已关联目标与候选读取的展示状态

#### Scenario: 待关联目标承接状态

- **WHEN** 用户从 `Module Detail` 带上下文进入 `DecisionCreatePage` 并成功创建后进入 `DecisionDetailPage`
- **THEN** 详情页必须继续承接“存在待关联 `Module`”的显式状态
- **AND** 该状态必须持续到用户完成 `LinkDecisionToTarget` 或主动放弃关联
- **AND** 不得在进入 `DecisionDetailPage` 后静默丢失该上下文

#### Scenario: 候选读取最小状态

- **WHEN** `DecisionDetailPage` 读取 `Decision -> Module` 候选
- **THEN** 候选读取最小状态至少应包含 `pending`、`ready`、`empty`、`error`
- **AND** 候选空结果必须进入 `empty`
- **AND** 不得将空结果误解释为资源不存在

#### Scenario: 关联写入最小状态

- **WHEN** 用户在 `DecisionDetailPage` 执行 `LinkDecisionToTarget`
- **THEN** 关联写入最小状态至少应包含 `idle`、`submitting`、`submit-success`、`submit-error`
- **AND** 该状态只归属于当前详情页上下文，不升级为跨路由全局状态

#### Scenario: 关联成功后的详情页行为

- **WHEN** `LinkDecisionToTarget` 提交成功
- **THEN** 用户必须停留在当前 `DecisionDetailPage`
- **AND** 当前详情页必须承接最新的已关联目标读取结果
- **AND** 若本次关联目标正是入口上下文中的待关联 `Module`，待关联目标状态必须被清除

#### Scenario: 关联失败后的状态保持

- **WHEN** `LinkDecisionToTarget` 提交失败
- **THEN** 页面必须停留在当前 `DecisionDetailPage`
- **AND** 当前选中的候选目标与待关联目标展示必须继续保留
- **AND** 错误反馈必须停留在当前详情页的关联动作上下文内

### Requirement: 页面返回路径必须冻结为单值规则

系统 SHALL 将当前阶段主要页面之间的主动返回路径冻结为单值规则，避免后续实现为同一动作设计两套默认返回行为。

#### Scenario: 从创建页主动返回

- **WHEN** 用户主动取消或返回离开 `DecisionCreatePage`
- **THEN** 默认返回路径必须为 `DecisionListPage`
- **AND** 当来源列表上下文存在时，必须保留原 `queryText` 与 `statusFilter`
- **AND** 当来源列表上下文不存在时，必须落到默认列表参数

#### Scenario: 从详情页主动返回

- **WHEN** 用户主动返回离开 `DecisionDetailPage`
- **THEN** 默认返回路径必须为 `DecisionListPage`
- **AND** 当来源列表上下文存在时，必须保留原 `queryText` 与 `statusFilter`
- **AND** 当来源列表上下文不存在时，必须落到默认列表参数

#### Scenario: 从创建成功后的详情页再返回列表

- **WHEN** 用户从 `DecisionCreatePage` 成功进入新建 `Decision` 的 `DecisionDetailPage` 后再返回列表
- **AND** 该详情页已继承来源列表上下文
- **THEN** 系统必须恢复进入创建前原有的 `queryText` 与 `statusFilter`
- **AND** 若该路径中来源列表上下文不存在（如从 `Module Detail` 直接进入创建），返回列表必须落到默认列表参数

### Requirement: 跨页面列表上下文必须冻结为单值承接策略

系统 SHALL 将 `DecisionCreatePage` 与 `DecisionDetailPage` 承接来源列表上下文的方式冻结为单值页面状态语义，避免后续实现出现路由层显式承接与页面内瞬时状态兜底两套分叉。

#### Scenario: 创建页来源列表上下文状态

- **WHEN** 用户进入 `DecisionCreatePage`
- **THEN** 页面必须持有"来源列表上下文存在"或"来源列表上下文不存在"的最小页面状态
- **AND** 从 `DecisionListPage` 进入时，来源列表上下文必须存在，且至少包含当前 `queryText` 与 `statusFilter`
- **AND** 从 `Module Detail` 进入时，来源列表上下文不存在
- **AND** 当来源列表上下文存在时，返回列表必须恢复原有 `queryText` 与 `statusFilter`
- **AND** 当来源列表上下文不存在时，返回列表必须落到默认列表参数

#### Scenario: 详情页来源列表上下文状态

- **WHEN** 用户进入 `DecisionDetailPage`
- **THEN** 页面必须持有"来源列表上下文存在"或"来源列表上下文不存在"的最小页面状态
- **AND** 从 `DecisionListPage` 进入详情时，来源列表上下文必须存在，且至少包含当前 `queryText` 与 `statusFilter`
- **AND** 从 `DecisionCreatePage` 成功创建后进入详情时，来源列表上下文必须继承自创建页
- **AND** 从 `Module Detail` 直接进入或从外部链接直达时，来源列表上下文不存在
- **AND** 当来源列表上下文存在时，返回列表必须恢复原有 `queryText` 与 `statusFilter`
- **AND** 当来源列表上下文不存在时，返回列表必须落到默认列表参数

### Requirement: 页面级 UI 状态应优先保持局部归属

系统 SHALL 将当前阶段的局部 UI 状态优先冻结为页面级或组件级归属，只在确有同页多组件共享需要时再抽为页面作用域状态容器。

#### Scenario: 草稿与详情页局部状态归属

- **WHEN** 后续实现 `DecisionCreatePage` 与 `DecisionDetailPage`
- **THEN** 表单草稿、提交错误、待关联目标展示、候选读取状态与关联动作状态应优先归属于当前页面或当前详情页上下文
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

### Requirement: Decision List 前端承接方式

`Decision List` 在当前阶段不仅 SHALL 承接读取结果，还应统一承接列表搜索参数、派生视图状态与从其他页面返回后的上下文恢复。

#### Scenario: 列表页上下文收口

- **WHEN** 用户从创建页或详情页回到 `DecisionListPage`
- **THEN** 页面必须能够按原有搜索参数恢复列表上下文
- **AND** 不得要求用户重新输入筛选条件

## REMOVED Requirements

### Requirement: 并列默认回流路径
**Reason**: 当前阶段需要单值交互流，避免实现阶段为同一动作保留两套默认返回路径。
**Migration**: 若后续确有多回流模式需求，必须在新的 `phase` 或 `audit` 中重新定义，不在当前阶段保留并列默认路径。
