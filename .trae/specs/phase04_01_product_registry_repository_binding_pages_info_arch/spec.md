# Phase04-01 Product Registry 与 Repository Binding 页面边界和信息结构 Spec

## Why

`phase04` 要把 `Product Registry + Repository Binding` 从 `phase01` 的页面范围与 `phase02` 的临时绑定承接推进为第三条正式执行主线，第一步必须先把页面边界、页面之间的入口关系，以及 `PC / 移动浏览器` 下的信息结构冻结成单值结论。只有先收住页面职责与动作 owner，后续 `Product / Repository` 模板、绑定关系、接口分组与实现设计才不会继续漂移。

## What Changes

- 冻结 `Product Registry / List`、`Product Create`、`Product Detail` 的最小页面边界
- 冻结 `Repository Binding / List`、`Repository Create`、`Repository Binding Detail / Workspace` 的最小页面边界
- 冻结 `Product Registry`、`Repository Binding`、`Module Detail` 之间的最小跳转关系与上下文入口
- 冻结五个核心动作与六个页面之间的单值动作归属矩阵
- 冻结 `Product List / Create / Detail` 与 `Repository Binding / List / Create / Detail(or Workspace)` 的最小页面级信息区块组成
- 冻结 `PC` 与移动浏览器下的基础信息密度策略
- 明确当前阶段不引入第二套移动端 UI 方案，不引入独立 `React Native` 客户端，不把完整 `PWA` 写成前置范围

## Impact

- Affected specs: `phase04_product_and_repository_binding_foundation`
- Affected code: 后续 `frontend/` 中的 `Product List`、`Product Create`、`Product Detail`、`Repository Binding / List`、`Repository Create`、`Repository Binding Detail / Workspace` 页面与路由结构，`Module Detail` 中的绑定兼容入口

## ADDED Requirements

### Requirement: Phase04 页面边界冻结

系统 SHALL 将 `phase04` 的页面主线冻结为 `Product Registry / List`、`Product Create`、`Product Detail`、`Repository Binding / List`、`Repository Create`、`Repository Binding Detail / Workspace` 六类页面或页面级功能块。

#### Scenario: 页面范围判定

- **WHEN** 接手者阅读 `phase04-01` 页面规格
- **THEN** 必须得到上述六类页面的单值结论
- **AND** 不得在本阶段额外引入 `Feature / Opportunity / Experiment` 页面
- **AND** 不得把 `Decision Center`、`Dashboard` 扩写为 `phase04-01` 的独立实现主线
- **AND** 不得把独立 `AI Assistant` 一级导航纳入 `phase04` 页面主线

### Requirement: Product Registry 列表页职责冻结

系统 SHALL 将 `Product Registry / List` 冻结为当前阶段 `Product` 主线的默认进入页，承接产品列表读取、筛选入口、创建入口与进入详情入口。

#### Scenario: Product 列表页职责判定

- **WHEN** 用户进入 `Product Registry`
- **THEN** 页面必须承接 `Product` 列表读取
- **AND** 必须提供进入 `Product Create` 的明确入口
- **AND** 必须提供进入 `Product Detail` 的明确入口
- **AND** 当前阶段的筛选只作为列表页入口能力存在，不扩写为复杂检索工作台

### Requirement: Product Create 页面职责冻结

系统 SHALL 将 `Product Create` 冻结为 `CreateProduct` 的唯一页面级承接入口。

#### Scenario: Product 创建页职责判定

- **WHEN** 用户执行 `CreateProduct`
- **THEN** 必须通过 `Product Create` 页面或等价页面级表单完成
- **AND** 当前阶段不得把产品创建分散到列表页内联复杂编辑流或其他独立页面主线

### Requirement: Product Detail 页面职责冻结

系统 SHALL 将 `Product Detail` 冻结为产品详情读取、已绑定模块/仓库读取、`BindModuleToProduct` 与进入仓库绑定流程的最小上下文承接页。

#### Scenario: Product 详情页职责判定

- **WHEN** 用户进入某个 `Product` 详情
- **THEN** 页面必须展示该 `Product` 的核心信息
- **AND** 必须展示当前已绑定 `Module` 列表
- **AND** 必须展示当前已绑定 `Repository` 列表
- **AND** 必须承接 `BindModuleToProduct` 的最小写入触点
- **AND** 必须提供进入 `Repository Binding Detail / Workspace` 的上下文入口
- **AND** 不得在当前阶段并行承接第二套仓库绑定写入流程

### Requirement: Repository Binding 列表页职责冻结

系统 SHALL 将 `Repository Binding / List` 冻结为当前阶段 `Repository` 主线的默认进入页，承接仓库列表读取、筛选入口、创建入口与进入绑定工作台入口。

#### Scenario: Repository 列表页职责判定

- **WHEN** 用户进入 `Repository Binding`
- **THEN** 页面必须承接 `Repository` 列表读取
- **AND** 必须提供进入 `Repository Create` 的明确入口
- **AND** 必须提供进入 `Repository Binding Detail / Workspace` 的明确入口
- **AND** 当前阶段的筛选只作为列表页入口能力存在，不扩写为复杂检索工作台

### Requirement: Repository Create 页面职责冻结

系统 SHALL 将 `Repository Create` 冻结为 `CreateRepository` 的唯一页面级承接入口。

#### Scenario: Repository 创建页职责判定

- **WHEN** 用户执行 `CreateRepository`
- **THEN** 必须通过 `Repository Create` 页面或等价页面级表单完成
- **AND** 当前阶段不得把仓库创建分散到列表页内联复杂编辑流或其他独立页面主线

### Requirement: Repository Binding Detail / Workspace 页面职责冻结

系统 SHALL 将 `Repository Binding Detail / Workspace` 冻结为仓库详情读取、候选读取、`BindRepositoryToProduct` 与 `MapModuleToRepository` 的最小承接页。

#### Scenario: Repository 工作台职责判定

- **WHEN** 用户进入某个 `Repository` 的绑定工作台
- **THEN** 页面必须展示该 `Repository` 的核心信息
- **AND** 必须展示当前已绑定 `Product` 列表
- **AND** 必须展示当前已映射 `Module` 列表
- **AND** 必须承接 `BindRepositoryToProduct` 的最小写入触点
- **AND** 必须承接 `MapModuleToRepository` 的最小写入触点
- **AND** 不得把 `BindModuleToProduct` 迁入该页面形成并列主写入流程

### Requirement: 五个核心动作 owner 冻结

系统 SHALL 将 `phase04` 的五个核心动作冻结为单值页面 owner，避免后续 `/spec` 与实现阶段出现并行 owner。

#### Scenario: 动作 owner 判定

- **WHEN** 接手者判断 `phase04` 的动作归属
- **THEN** 必须得到以下单值结论：
- **AND** `CreateProduct` → `Product Create`
- **AND** `CreateRepository` → `Repository Create`
- **AND** `BindModuleToProduct` → `Product Detail`
- **AND** `BindRepositoryToProduct` → `Repository Binding Detail / Workspace`
- **AND** `MapModuleToRepository` → `Repository Binding Detail / Workspace`

### Requirement: 页面跳转关系冻结

系统 SHALL 冻结 `Product Registry`、`Repository Binding` 与 `Module Detail` 之间的最小跳转关系，避免页面职责与交互流在后续规格中继续漂移。

#### Scenario: Product 主线最小跳转关系判定

- **WHEN** 用户从 `Product Registry` 进入产品主线
- **THEN** 必须支持 `Product Registry / List -> Product Create`
- **AND** 必须支持 `Product Registry / List -> Product Detail`
- **AND** 必须支持 `Product Detail -> Product Registry / List`
- **AND** 必须支持 `Product Detail -> Repository Binding Detail / Workspace` 的带上下文跳转

#### Scenario: Repository 主线最小跳转关系判定

- **WHEN** 用户从 `Repository Binding` 进入仓库主线
- **THEN** 必须支持 `Repository Binding / List -> Repository Create`
- **AND** 必须支持 `Repository Binding / List -> Repository Binding Detail / Workspace`
- **AND** 必须支持 `Repository Binding Detail / Workspace -> Repository Binding / List`

#### Scenario: Module Detail 兼容入口判定

- **WHEN** 用户从 `Module Detail` 发起与 `Product / Repository` 绑定相关的后续动作
- **THEN** 当前阶段只允许通过轻量跳转或带上下文入口进入正式主入口
- **AND** 不得继续在 `Module Detail` 内直接提交 `BindModuleToProduct`、`BindRepositoryToProduct` 或 `MapModuleToRepository`
- **AND** 不得把 `Module Detail` 扩写为第二个绑定工作台

### Requirement: Product 页面级信息区块冻结

系统 SHALL 将 `Product Registry / List`、`Product Create`、`Product Detail` 的最小页面级信息区块冻结为单值结论，使后续前端设计可直接进入实现。

#### Scenario: Product 列表页信息区块判定

- **WHEN** 渲染 `Product Registry / List`
- **THEN** 页面至少必须包含列表工具栏区、列表内容区与空状态区
- **AND** 列表工具栏区必须承接搜索输入、状态筛选与进入 `Product Create` 的入口
- **AND** 列表内容区必须至少展示 `name / description / status / created_at / module_bind_count / repository_bind_count`
- **AND** 空状态区必须在无产品时引导用户进入 `Product Create`

#### Scenario: Product 创建页信息区块判定

- **WHEN** 渲染 `Product Create`
- **THEN** 页面至少必须包含结构化表单区与提交取消操作区
- **AND** 结构化表单区必须承接 `name / description / status`
- **AND** 提交取消操作区必须承接 `CreateProduct` 提交与返回列表路径

#### Scenario: Product 详情页信息区块判定

- **WHEN** 渲染 `Product Detail`
- **THEN** 页面至少必须包含核心信息区、已绑定模块区、已绑定仓库区与绑定入口区
- **AND** 核心信息区必须展示 `Product` 的核心字段
- **AND** 已绑定模块区必须承接 `BindModuleToProduct` 的当前结果与写入触点
- **AND** 已绑定仓库区必须展示已建立的 `Product -> Repository` 结果摘要
- **AND** 绑定入口区必须承接进入 `Repository Binding Detail / Workspace` 的上下文跳转

### Requirement: Repository 页面级信息区块冻结

系统 SHALL 将 `Repository Binding / List`、`Repository Create`、`Repository Binding Detail / Workspace` 的最小页面级信息区块冻结为单值结论，使后续前端设计可直接进入实现。

#### Scenario: Repository 列表页信息区块判定

- **WHEN** 渲染 `Repository Binding / List`
- **THEN** 页面至少必须包含列表工具栏区、列表内容区与空状态区
- **AND** 列表工具栏区必须承接搜索输入（`queryText`）、状态筛选（`statusFilter`）与进入 `Repository Create` 的入口
- **AND** 当前阶段工具栏只冻结上述三项；是否引入 `providerFilter` 不在 `phase04-01` 冻结范围，留给 `phase04-02 / phase04-06` 再正式单值化
- **AND** 列表内容区必须至少展示 `name / url / provider / status / created_at / product_bind_count / module_bind_count`
- **AND** 空状态区必须在无仓库时引导用户进入 `Repository Create`

#### Scenario: Repository 创建页信息区块判定

- **WHEN** 渲染 `Repository Create`
- **THEN** 页面至少必须包含结构化表单区与提交取消操作区
- **AND** 结构化表单区必须承接 `name / url / provider / status`
- **AND** 提交取消操作区必须承接 `CreateRepository` 提交与返回列表路径

#### Scenario: Repository 工作台信息区块判定

- **WHEN** 渲染 `Repository Binding Detail / Workspace`
- **THEN** 页面至少必须包含核心信息区、已绑定产品区、已映射模块区与绑定工作台区
- **AND** 核心信息区必须展示 `Repository` 的核心字段
- **AND** 已绑定产品区必须展示当前 `Product -> Repository` 结果
- **AND** 已映射模块区必须展示当前 `Module -> Repository` 结果
- **AND** 绑定工作台区必须承接候选读取、空状态、`BindRepositoryToProduct` 与 `MapModuleToRepository` 写入触点

### Requirement: PC 与移动浏览器信息密度策略冻结

系统 SHALL 在单一 `React Web` 前端交付策略下，同时定义 `PC` 与移动浏览器的基础信息密度规则。

#### Scenario: 桌面端信息密度

- **WHEN** 页面在 `PC` 桌面环境展示
- **THEN** 列表页应优先承接更高信息密度
- **AND** 详情页或工作台页可在同屏承接核心信息、当前绑定结果与写入入口

#### Scenario: 移动浏览器信息密度

- **WHEN** 页面在移动浏览器窄屏环境展示
- **THEN** 必须采用同一套页面语义与动作体系
- **AND** 必须通过信息裁剪、区块折叠、垂直重排或分层展示降低拥挤度
- **AND** 不得引入第二套独立移动端 UI 方案
- **AND** 不得要求独立 `React Native` 客户端或完整 `PWA` 作为当前阶段前提

## MODIFIED Requirements

### Requirement: Repository Binding 页面职责解释

`Repository Binding` 在 `phase04` 中 SHALL 不只被解释为一个“仓库绑定概念”，而是被解释为包含列表、创建、详情/工作台在内的一条完整页面主线。

#### Scenario: Repository Binding 主线解释

- **WHEN** 后续 `/spec` 或实现讨论 `Repository Binding`
- **THEN** 必须将其理解为一条完整页面主线
- **AND** 不得只剩单独列表页而把创建、详情与绑定工作台散落为无归属入口

### Requirement: Module Detail 的绑定承接方式

`Module Detail` 在 `phase04` 中 SHALL 从 `phase02` 的绑定写入承接位回落为摘要与上下文入口，而不是继续保留第二套正式绑定工作台。

#### Scenario: Module Detail 绑定承接解释

- **WHEN** 后续 `/spec` 或实现讨论 `Module Detail` 的绑定能力
- **THEN** 必须将其理解为摘要展示与正式主入口跳转位
- **AND** 不得继续把 `Module Detail` 实现为 `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 的并列写入 owner

## REMOVED Requirements

### Requirement: 第二套移动端 UI 方案
**Reason**: 当前阶段已冻结为单一 `React Web` 同时覆盖 `PC` 与移动浏览器，不需要并行维护第二套移动端页面方案。
**Migration**: 若后续确有独立客户端需要，必须进入新的 `phase / audit` 流程，不在 `phase04-01` 当前规格中处理。
