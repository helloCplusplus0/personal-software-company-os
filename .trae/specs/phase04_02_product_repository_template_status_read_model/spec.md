# Phase04-02 Product / Repository 模板、状态语义与最小展示模型 Spec

## Why

`phase04-01` 已经冻结了 `Product Registry` 与 `Repository Binding` 的页面边界和信息结构，但页面要真正可进入后续 `/spec` 与实现，还必须把 `Product` 与 `Repository` 到底记录哪些字段、`status` 如何表达、列表/详情/绑定最少展示什么写成单值结论。否则后续前端、后端、`.proto` 与验收会继续在字段口径、状态解释、`provider` 语义和展示模型之间来回漂移。

## What Changes

- 冻结 `Product` 的最小结构化模板字段集合
- 冻结 `Repository` 的最小结构化模板字段集合
- 冻结 `Product / Repository` 字段级 `required / optional` 规则
- 冻结 `CreateProduct / CreateRepository` 的最小创建校验前提
- 冻结 `Product / Repository.status` 的最小持久化枚举与状态语义
- 冻结 `status` 在创建写入、DTO、`.proto` 与列表 `statusFilter` 中的单值语义
- 冻结 `Repository.provider` 是否采用受控枚举，以及仓库列表是否引入 `providerFilter`
- 冻结 `Product List / Detail`、`Repository List / Detail(or Workspace)` 与绑定展示所需的最小读模型
- 明确当前阶段不引入超出 `v0.1` 的复杂生命周期、自动扫描字段或远程导入字段

## Impact

- Affected specs: `phase04_product_and_repository_binding_foundation`
- Affected code: 后续 `frontend/` 中的 `Product / Repository` 表单与列表详情展示模型、后续 `backend/` 中的 `Product / Repository` DTO 与校验模型、后续 `.proto` 中的 `Product / Repository / Binding` 消息字段

## ADDED Requirements

### Requirement: Product 最小结构化模板冻结

系统 SHALL 将 `Product` 的最小结构化模板冻结为当前阶段唯一记录模板。

#### Scenario: 判断 Product 最小结构化模板

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `Product` 字段
- **THEN** 必须至少承接 `name / description / status`
- **AND** 不得减少上述字段，导致 `Product` 退化为无法完成最小资产登记的空壳记录
- **AND** 不得在当前阶段额外引入 `customer / value_proposition / business_model / metrics / remote_import_source` 等超出 `v0.1` 的前置字段

### Requirement: Repository 最小结构化模板冻结

系统 SHALL 将 `Repository` 的最小结构化模板冻结为当前阶段唯一记录模板。

#### Scenario: 判断 Repository 最小结构化模板

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `Repository` 字段
- **THEN** 必须至少承接 `name / url / provider / status`
- **AND** 不得减少上述字段，导致 `Repository` 无法完成最小实现锚点登记
- **AND** 不得在当前阶段额外引入 `oauth_binding / remote_import_status / sync_cursor / scanned_commit` 等自动化集成字段

### Requirement: Product 字段级必填规则冻结

系统 SHALL 将 `Product` 模板中的字段级 `required / optional` 规则冻结为单值结论。

#### Scenario: 判断 Product 字段级 required / optional

- **WHEN** 用户执行 `CreateProduct`
- **THEN** `name / description / status` 必须为必填字段
- **AND** 当前阶段 `Product` 模板中不额外存在必填之外的用户提交扩展字段
- **AND** 必填字段在去首尾空白后不得为空字符串

### Requirement: Repository 字段级必填规则冻结

系统 SHALL 将 `Repository` 模板中的字段级 `required / optional` 规则冻结为单值结论。

#### Scenario: 判断 Repository 字段级 required / optional

- **WHEN** 用户执行 `CreateRepository`
- **THEN** `name / url / provider / status` 必须为必填字段
- **AND** 当前阶段 `Repository` 模板中不额外存在必填之外的用户提交扩展字段
- **AND** 必填字段在去首尾空白后不得为空字符串

### Requirement: CreateProduct / CreateRepository 最小校验前提冻结

系统 SHALL 将 `CreateProduct` 与 `CreateRepository` 的最小创建校验前提冻结为可直接进入实现的单值规则。

#### Scenario: Product 创建成功前提

- **WHEN** 用户提交 `CreateProduct`
- **THEN** `name / description / status` 必须全部满足最小非空校验
- **AND** `status` 必须属于当前阶段冻结的最小持久化枚举

#### Scenario: Repository 创建成功前提

- **WHEN** 用户提交 `CreateRepository`
- **THEN** `name / url / provider / status` 必须全部满足最小非空校验
- **AND** `status` 必须属于当前阶段冻结的最小持久化枚举

#### Scenario: 创建失败前提

- **WHEN** 用户提交缺少必填字段或非法 `status` 的 `Product / Repository`
- **THEN** 系统必须返回明确校验失败语义
- **AND** 不得降级为模糊通用错误

### Requirement: Product / Repository 状态枚举冻结

系统 SHALL 将当前阶段 `Product.status` 与 `Repository.status` 的最小持久化枚举冻结为与上游基线一致的状态集合。

#### Scenario: 判断状态范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `Product.status` 或 `Repository.status`
- **THEN** 当前阶段只允许使用 `active / archived`
- **AND** 不得在当前阶段额外引入 `draft / syncing / disconnected / retired / imported` 等新状态

### Requirement: Product / Repository 状态语义冻结

系统 SHALL 为当前阶段的每个 `Product.status` 与 `Repository.status` 提供页面、数据与合同都可复用的最小语义解释。

#### Scenario: active 状态语义

- **WHEN** `Product.status = active` 或 `Repository.status = active`
- **THEN** 必须表示该记录当前处于可见、可继续绑定、可继续维护的有效状态

#### Scenario: archived 状态语义

- **WHEN** `Product.status = archived` 或 `Repository.status = archived`
- **THEN** 必须表示该记录已归档保留
- **AND** 仍允许作为历史事实被读取和展示
- **AND** 当前阶段不把“已归档”自动解释为记录不可见或必须从候选中移除
- **AND** 候选读取的具体过滤策略属于 `phase04-03` 范围，本规格只冻结 `archived` 状态自身的可见性语义

### Requirement: status 创建写入与默认值语义冻结

系统 SHALL 将 `status` 在创建写入、DTO 与 `.proto` 写模型中的语义冻结为单值结论。

#### Scenario: 创建写入 status 提交语义

- **WHEN** 用户执行 `CreateProduct` 或 `CreateRepository`
- **THEN** 当前阶段必须按“显式提交 `status`”处理
- **AND** Create 页面输入模型、HTTP DTO 与 `.proto` 写请求都必须携带 `status`
- **AND** 不得保留“创建请求可省略 `status`，由服务端隐式补默认值”的并行解释

#### Scenario: 默认 active 语义

- **WHEN** 当前阶段讨论 `CreateProduct` 或 `CreateRepository` 的默认状态
- **THEN** “默认 `active`”只表示创建页、HTTP DTO 与 `.proto` 写模型在未发生用户改动时都应预填并显式提交 `active`
- **AND** 不得把“默认 `active`”解释为服务端静默补值或合同层隐式默认值

### Requirement: statusFilter 语义冻结

系统 SHALL 将 `Product List` 与 `Repository List` 的 `statusFilter` 语义冻结为 UI/路由层枚举，而不是持久化字段值。

#### Scenario: 判断 statusFilter 取值范围

- **WHEN** 后续前端列表状态、路由搜索参数或交互流讨论 `statusFilter`
- **THEN** 当前阶段只允许使用 `all / active / archived`
- **AND** `all` 只允许存在于 UI 与路由搜索参数层
- **AND** `all` 不得写入数据库、HTTP 持久化 DTO、后端领域模型或 `.proto` 持久化字段

### Requirement: Repository.provider 语义冻结

系统 SHALL 将当前阶段 `Repository.provider` 冻结为必填字符串字段，但不采用受控枚举。

#### Scenario: 判断 provider 字段语义

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `Repository.provider`
- **THEN** `provider` 必须作为必填字符串承接
- **AND** 当前阶段只要求去首尾空白后的最小非空校验
- **AND** 不得把 `provider` 冻结为受控枚举
- **AND** 不得在当前阶段引入基于 `provider` 的自动远程导入、自动鉴权或自动同步语义

### Requirement: Repository 列表不引入 providerFilter

系统 SHALL 将当前阶段 `Repository Binding / List` 的筛选维度冻结为 `queryText / statusFilter`，不引入 `providerFilter`。

#### Scenario: 判断 Repository 列表筛选维度

- **WHEN** 后续前端状态模型、路由搜索参数或页面交互讨论 `Repository Binding / List` 的筛选维度
- **THEN** 当前阶段工具栏的筛选维度只承接 `queryText / statusFilter`
- **AND** 不得在当前阶段增加 `providerFilter`
- **AND** 若后续确需新增 `providerFilter`，必须进入新的冻结任务重新单值化

### Requirement: Product List 最小展示模型冻结

系统 SHALL 将 `Product Registry / List` 的最小展示字段冻结为当前阶段唯一列表读模型。

#### Scenario: Product 列表页最小展示字段

- **WHEN** 用户在 `Product Registry / List` 查看产品列表
- **THEN** 每个列表项至少必须展示 `name / description / status / created_at / module_bind_count / repository_bind_count`
- **AND** 不得要求用户进入详情页才能理解最基本的产品状态与绑定概况
- **AND** 当前阶段不得在列表页扩写 `customer / value_proposition / business_model / metrics` 等额外展示字段

### Requirement: Product Detail 最小展示模型冻结

系统 SHALL 将 `Product Detail` 的最小展示字段与绑定展示要求冻结为当前阶段唯一详情读模型。

#### Scenario: Product 详情页最小展示字段

- **WHEN** 用户进入 `Product Detail`
- **THEN** 页面至少必须展示 `name / description / status / created_at`
- **AND** 必须展示当前已绑定模块列表
- **AND** 必须展示当前已绑定仓库列表
- **AND** 不得要求用户跳转其他独立主线才能理解该产品的当前绑定概况

#### Scenario: Product 详情页已绑定模块展示模型

- **WHEN** 页面展示 `Product -> Module` 绑定结果
- **THEN** 每条已绑定模块至少必须承接 `module_id / module_name / module_status`
- **AND** 当前阶段不得扩写 `latest_release_version / reuse_score / capability_summary`

#### Scenario: Product 详情页已绑定仓库展示模型

- **WHEN** 页面展示 `Product -> Repository` 绑定结果
- **THEN** 每条已绑定仓库至少必须承接 `repository_id / repository_name / provider / repository_status`
- **AND** 当前阶段不得扩写远程同步状态、分支状态或扫描结果

### Requirement: Repository List 最小展示模型冻结

系统 SHALL 将 `Repository Binding / List` 的最小展示字段冻结为当前阶段唯一列表读模型。

#### Scenario: Repository 列表页最小展示字段

- **WHEN** 用户在 `Repository Binding / List` 查看仓库列表
- **THEN** 每个列表项至少必须展示 `name / url / provider / status / created_at / product_bind_count / module_bind_count`
- **AND** 不得要求用户进入工作台页才能理解最基本的仓库状态与绑定概况
- **AND** 当前阶段不得扩写同步时间、远程权限状态或扫描统计

### Requirement: Repository Detail / Workspace 最小展示模型冻结

系统 SHALL 将 `Repository Binding Detail / Workspace` 的最小展示字段与绑定展示要求冻结为当前阶段唯一详情读模型。

#### Scenario: Repository 工作台最小展示字段

- **WHEN** 用户进入 `Repository Binding Detail / Workspace`
- **THEN** 页面至少必须展示 `name / url / provider / status / created_at`
- **AND** 必须展示当前已绑定产品列表
- **AND** 必须展示当前已映射模块列表

#### Scenario: Repository 工作台已绑定产品展示模型

- **WHEN** 页面展示 `Repository -> Product` 绑定结果
- **THEN** 每条已绑定产品至少必须承接 `product_id / product_name / product_status`
- **AND** 当前阶段不得扩写业务指标、价值主张或客户画像

#### Scenario: Repository 工作台已映射模块展示模型

- **WHEN** 页面展示 `Repository -> Module` 映射结果
- **THEN** 每条已映射模块至少必须承接 `module_id / module_name / module_status`
- **AND** 当前阶段不得扩写实现扫描、依赖分析或自动识别结果

## MODIFIED Requirements

### Requirement: Product / Repository Create 表单字段解释

`Product Create` 与 `Repository Create` 在当前阶段 SHALL 不只被解释为“任意文本输入页”，而必须被解释为承接冻结后的最小结构化模板、字段级校验规则与 `status` 单值语义的页面级表单。

#### Scenario: Create 页面字段解释

- **WHEN** 后续 `/spec` 或实现讨论 `Product Create` 或 `Repository Create`
- **THEN** 必须以本次冻结的最小结构化模板、`required / optional` 规则、`status` 语义与 `provider` 语义为唯一上游
- **AND** 不得重新发明第二套字段定义、默认值策略或筛选解释

### Requirement: Repository List 筛选结构解释

`Repository Binding / List` 在当前阶段 SHALL 不只被解释为“未来可能扩展的过滤工作台”，而必须被解释为当前只承接 `queryText / statusFilter` 的最小列表筛选入口。

#### Scenario: Repository 列表筛选解释

- **WHEN** 后续 `/spec` 或实现讨论 `Repository Binding / List` 的筛选结构
- **THEN** 必须将其理解为 `queryText / statusFilter` 两个最小筛选维度
- **AND** 不得在没有新增冻结任务前，把 `providerFilter`、远程同步状态筛选或扫描状态筛选并入当前阶段列表工具栏

## REMOVED Requirements

### Requirement: 复杂生命周期、自动扫描字段或远程导入字段前置
**Reason**: 当前阶段只需要支撑 `Product Registry + Repository Binding` 最小可执行闭环，不需要引入超出 `v0.1` 的复杂生命周期、自动扫描或远程导入模型。
**Migration**: 若后续确需更复杂的生命周期、扫描状态、远程导入来源或自动同步字段，必须进入新的 `phase / audit` 流程单独冻结，不在 `phase04-02` 当前规格中处理。
