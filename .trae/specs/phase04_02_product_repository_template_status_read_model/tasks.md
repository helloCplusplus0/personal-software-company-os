# Tasks

- [x] Task 1: 冻结 `Product / Repository` 最小结构化模板。将当前阶段唯一允许的记录字段写成单值结论，避免表单、DTO 与合同层继续漂移。
  - [x] SubTask 1.1: 明确 `Product` 最小字段集合为 `name / description / status`
    - 证据：`spec.md` §ADDED Requirements「Product 最小结构化模板冻结」Scenario: 判断 Product 最小结构化模板 L32-35 — 必须至少承接 `name / description / status`
  - [x] SubTask 1.2: 明确 `Repository` 最小字段集合为 `name / url / provider / status`
    - 证据：`spec.md` §ADDED Requirements「Repository 最小结构化模板冻结」Scenario: 判断 Repository 最小结构化模板 L43-46 — 必须至少承接 `name / url / provider / status`
  - [x] SubTask 1.3: 明确当前阶段不引入复杂生命周期、自动扫描或远程导入字段
    - 证据：`spec.md` §REMOVED Requirements「复杂生命周期、自动扫描字段或远程导入字段前置」L263-264 + §ADDED Requirements L35（Product 不得引入 `customer / value_proposition / business_model / metrics / remote_import_source`）与 L46（Repository 不得引入 `oauth_binding / remote_import_status / sync_cursor / scanned_commit`）

- [x] Task 2: 冻结字段级 `required / optional` 与最小创建校验前提。把必填字段、空字符串规则与非法输入判定写成可直接实现的最小规则。
  - [x] SubTask 2.1: 明确 `Product` 的 `name / description / status` 为必填
    - 证据：`spec.md` §ADDED Requirements「Product 字段级必填规则冻结」Scenario: 判断 Product 字段级 required / optional L54-57
  - [x] SubTask 2.2: 明确 `Repository` 的 `name / url / provider / status` 为必填
    - 证据：`spec.md` §ADDED Requirements「Repository 字段级必填规则冻结」Scenario: 判断 Repository 字段级 required / optional L65-68
  - [x] SubTask 2.3: 明确必填字段去首尾空白后不得为空字符串
    - 证据：`spec.md` §ADDED Requirements L57（Product）+ L68（Repository）— 「必填字段在去首尾空白后不得为空字符串」
  - [x] SubTask 2.4: 明确缺少必填字段或非法 `status` 时必须返回明确校验失败
    - 证据：`spec.md` §ADDED Requirements「CreateProduct / CreateRepository 最小校验前提冻结」Scenario: 创建失败前提 L86-90 — 必须返回明确校验失败语义，不得降级为模糊通用错误

- [x] Task 3: 冻结 `Product / Repository.status` 的最小枚举、状态语义与创建写入语义。把持久化枚举、默认值、显式提交规则和 `statusFilter` 写成前后端、页面与合同都可复用的单值结论。
  - [x] SubTask 3.1: 明确最小持久化状态集合为 `active / archived`
    - 证据：`spec.md` §ADDED Requirements「Product / Repository 状态枚举冻结」Scenario: 判断状态范围 L98-100 — 当前阶段只允许使用 `active / archived`
  - [x] SubTask 3.2: 明确 `active` 的状态语义
    - 证据：`spec.md` §ADDED Requirements「Product / Repository 状态语义冻结」Scenario: active 状态语义 L108-109 — 表示可见、可继续绑定、可继续维护的有效状态
  - [x] SubTask 3.3: 明确 `archived` 的状态语义
    - 证据：`spec.md` §ADDED Requirements「Product / Repository 状态语义冻结」Scenario: archived 状态语义 L113-117 — 已归档保留，仍允许作为历史事实被读取和展示，不自动解释为不可见或必须从候选中移除；候选读取的具体过滤策略属于 `phase04-03` 范围，本规格只冻结 `archived` 状态自身的可见性语义
  - [x] SubTask 3.4: 明确创建写入继续按"显式提交 status"处理
    - 证据：`spec.md` §ADDED Requirements「status 创建写入与默认值语义冻结」Scenario: 创建写入 status 提交语义 L125-128 — Create 页面输入模型、HTTP DTO 与 `.proto` 写请求都必须携带 `status`，不得保留"创建请求可省略 status，由服务端隐式补默认值"的并行解释
  - [x] SubTask 3.5: 明确"默认 active"只表示预填并显式提交，不得解释为服务端隐式补默认值
    - 证据：`spec.md` §ADDED Requirements「status 创建写入与默认值语义冻结」Scenario: 默认 active 语义 L132-134 — "默认 active"只表示预填并显式提交，不得解释为服务端静默补值或合同层隐式默认值
  - [x] SubTask 3.6: 明确列表 `statusFilter` 只允许 `all / active / archived`，且 `all` 只存在于 UI/路由层
    - 证据：`spec.md` §ADDED Requirements「statusFilter 语义冻结」Scenario: 判断 statusFilter 取值范围 L142-145 — 只允许 `all / active / archived`，`all` 只存在于 UI 与路由搜索参数层，不得写入数据库/HTTP 持久化 DTO/后端领域模型/`.proto` 持久化字段

- [x] Task 4: 冻结 `Repository.provider` 与仓库列表筛选边界。把 `provider` 是否采用受控枚举、仓库列表是否引入 `providerFilter` 写成单值结论。
  - [x] SubTask 4.1: 明确 `Repository.provider` 当前阶段为必填字符串字段，不采用受控枚举
    - 证据：`spec.md` §ADDED Requirements「Repository.provider 语义冻结」Scenario: 判断 provider 字段语义 L153-157 — 必须作为必填字符串承接，不得把 `provider` 冻结为受控枚举
  - [x] SubTask 4.2: 明确 `provider` 只要求最小非空校验，不引入自动远程语义
    - 证据：`spec.md` §ADDED Requirements「Repository.provider 语义冻结」Scenario L155-157 — 只要求去首尾空白后的最小非空校验，不得引入基于 `provider` 的自动远程导入、自动鉴权或自动同步语义
  - [x] SubTask 4.3: 明确 `Repository Binding / List` 当前阶段不引入 `providerFilter`
    - 证据：`spec.md` §ADDED Requirements「Repository 列表不引入 providerFilter」Scenario: 判断 Repository 列表筛选维度 L165-168 — 当前阶段工具栏的筛选维度只承接 `queryText / statusFilter`，不得增加 `providerFilter`
  - [x] SubTask 4.4: 明确仓库列表工具栏当前只承接 `queryText / statusFilter`
    - 证据：`spec.md` §ADDED Requirements「Repository 列表不引入 providerFilter」Scenario L166 + §MODIFIED Requirements「Repository List 筛选结构解释」Scenario: Repository 列表筛选解释 L257-259 — 必须理解为 `queryText / statusFilter` 两个最小筛选维度

- [x] Task 5: 冻结 `Product / Repository` 的最小展示模型。将列表、详情与绑定展示至少展示哪些字段写成可直接进入后续页面、状态和合同设计的读模型。
  - [x] SubTask 5.1: 明确 `Product List` 至少展示 `name / description / status / created_at / module_bind_count / repository_bind_count`
    - 证据：`spec.md` §ADDED Requirements「Product List 最小展示模型冻结」Scenario: Product 列表页最小展示字段 L176-179
  - [x] SubTask 5.2: 明确 `Product Detail` 至少展示 `name / description / status / created_at` 与已绑定模块、已绑定仓库
    - 证据：`spec.md` §ADDED Requirements「Product Detail 最小展示模型冻结」Scenario: Product 详情页最小展示字段 L187-191
  - [x] SubTask 5.3: 明确 `Product Detail` 的已绑定模块最小展示模型为 `module_id / module_name / module_status`
    - 证据：`spec.md` §ADDED Requirements「Product Detail 最小展示模型冻结」Scenario: Product 详情页已绑定模块展示模型 L195-197
  - [x] SubTask 5.4: 明确 `Product Detail` 的已绑定仓库最小展示模型为 `repository_id / repository_name / provider / repository_status`
    - 证据：`spec.md` §ADDED Requirements「Product Detail 最小展示模型冻结」Scenario: Product 详情页已绑定仓库展示模型 L201-203
  - [x] SubTask 5.5: 明确 `Repository List` 至少展示 `name / url / provider / status / created_at / product_bind_count / module_bind_count`
    - 证据：`spec.md` §ADDED Requirements「Repository List 最小展示模型冻结」Scenario: Repository 列表页最小展示字段 L211-214
  - [x] SubTask 5.6: 明确 `Repository Detail / Workspace` 至少展示 `name / url / provider / status / created_at` 与已绑定产品、已映射模块
    - 证据：`spec.md` §ADDED Requirements「Repository Detail / Workspace 最小展示模型冻结」Scenario: Repository 工作台最小展示字段 L222-225
  - [x] SubTask 5.7: 明确 `Repository Detail / Workspace` 的已绑定产品最小展示模型为 `product_id / product_name / product_status`
    - 证据：`spec.md` §ADDED Requirements「Repository Detail / Workspace 最小展示模型冻结」Scenario: Repository 工作台已绑定产品展示模型 L229-231
  - [x] SubTask 5.8: 明确 `Repository Detail / Workspace` 的已映射模块最小展示模型为 `module_id / module_name / module_status`
    - 证据：`spec.md` §ADDED Requirements「Repository Detail / Workspace 最小展示模型冻结」Scenario: Repository 工作台已映射模块展示模型 L235-237

- [x] Task 6: 完成规格校验。确认本次 `phase04-02` 规格可以直接作为后续页面、接口与合同设计的上游。
  - [x] SubTask 6.1: 验证 `Product / Repository` 模板字段已经单值化
    - 证据：`spec.md` §ADDED Requirements「Product 最小结构化模板冻结」+「Repository 最小结构化模板冻结」— 两个模板各自单值化为 `name / description / status` 与 `name / url / provider / status`
  - [x] SubTask 6.2: 验证 `required / optional` 与最小创建校验规则已经明确
    - 证据：`spec.md` §ADDED Requirements「Product 字段级必填规则冻结」+「Repository 字段级必填规则冻结」+「CreateProduct / CreateRepository 最小校验前提冻结」— 必填字段、空字符串规则与非法输入判定已单值化
  - [x] SubTask 6.3: 验证 `status` 枚举、状态语义、显式提交规则与 `statusFilter` 语义已经明确
    - 证据：`spec.md` §ADDED Requirements「Product / Repository 状态枚举冻结」+「状态语义冻结」+「status 创建写入与默认值语义冻结」+「statusFilter 语义冻结」— 枚举 `active / archived`、语义、显式提交、`statusFilter` `all / active / archived` 已单值化
  - [x] SubTask 6.4: 验证 `provider` 是否受控枚举、是否引入 `providerFilter` 已经明确
    - 证据：`spec.md` §ADDED Requirements「Repository.provider 语义冻结」（不采用受控枚举）+「Repository 列表不引入 providerFilter」— 两个决策已单值化
  - [x] SubTask 6.5: 验证列表、详情与绑定展示的最小读模型已经明确
    - 证据：`spec.md` §ADDED Requirements「Product List 最小展示模型冻结」+「Product Detail 最小展示模型冻结」+「Repository List 最小展示模型冻结」+「Repository Detail / Workspace 最小展示模型冻结」— 四组读模型含绑定展示已单值化
  - [x] SubTask 6.6: 验证未引入超出 `v0.1` 的复杂生命周期、自动扫描字段或远程导入字段
    - 证据：`spec.md` §REMOVED Requirements「复杂生命周期、自动扫描字段或远程导入字段前置」+ §ADDED Requirements L100（不得引入 `draft / syncing / disconnected / retired / imported`）+ L35/L46（不得引入超出 `v0.1` 的前置字段）

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
