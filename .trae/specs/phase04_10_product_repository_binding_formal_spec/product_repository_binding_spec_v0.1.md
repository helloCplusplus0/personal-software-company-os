# PSCO Product / Repository / Binding 规格 v0.1 — 正式规格正文

> **文档定位**：本文档是 `phase04_product_and_repository_binding_foundation` 的正式规格正文，作为后续 `phase04-11 / 12 / 13` 合同落地、实现、验收与收口以及 `phase05` 引用 `Product Registry + Repository Binding` 主线时的直接上游规格来源。
> **上游收敛**：本文档由 `phase04-01` 到 `phase04-09` 的冻结结论收敛而成，不另立第二套边界。
> **互链前提**：本文档以 `phase01-06` 的 `mvp_spec_v0.1.md` 为当前阶段唯一执行层上游，完整承接 `module_registry_spec_v0.1.md` 与 `phase02-12` 验收结论中已交付的 `Module Registry` 边界、`decision_center_spec_v0.1.md` 与 `phase03-14` 验收结论中已交付的 `Decision Center` 边界，并与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致。
> **状态约束**：`phase04` 已完成收口，项目当前根级阶段已切换到 `phase05_dashboard_feedback_foundation`；本文档继续作为 `Product / Repository / Binding` 的已交付正式规格入口使用，但不单独承担根级当前阶段状态表达。

---

## 1. 技术路线

本文档继承 `mvp_spec_v0.1.md` §1 的技术路线，聚焦 `Product / Repository / Binding` 当前阶段：

- 项目路线：`Durable System Track`
- 正式运行主线：`React Web + Go Backend + PostgreSQL`
- 前端：`React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`
- 前端交付策略：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI；不引入独立 `React Native` 客户端，`PWA` 仅作可兼容增强方向
- 后端：`Go`，模块化单体优先
- 数据库：`PostgreSQL` 为唯一数据库主线
- 合同：`Contract First`，跨语言标准为 `Protocol Buffers`；当前阶段必须为 `Product / Repository / Binding` 落地最小 `.proto` 合同源
- 合同工具链：`buf build / lint / generate / breaking`
- 部署：`Caddy + systemd`，运行方式 `Single Server First`

> 约束：不得重新解释为 `Product Track`；不得把 `Rust` 写成当前阶段必需项；`Local First` 当前解释为数据所有权优先，不等于切换到 `SQLite`。

---

## 2. 对象范围

### 2.1 核心对象（`Product / Repository / Binding` 直接承接）

- `Product`：产品登记的核心实体
- `Repository`：仓库登记与绑定关系的核心实体
- `ProductModuleBinding`：`Product` 与 `Module` 的绑定关系（`product_modules`）
- `ProductRepositoryBinding`：`Product` 与 `Repository` 的绑定关系（`product_repositories`）
- `ModuleRepositoryMapping`：`Module` 与 `Repository` 的映射关系（`module_repositories`）

### 2.2 连接对象（只读或候选读取，不承接写入主线）

- `Module`：作为 `BindModuleToProduct` / `MapModuleToRepository` 的候选来源与已绑定摘要读取对象
- `Decision`：保留为受控连接位，当前阶段不承接 `Decision -> Product / Repository` 正式写入

### 2.3 非归属能力

- `CreateModule`、`CreateRelease`、`RecordDecision`、`LinkDecisionToTarget` 不归属于当前阶段 `Product Registry` 与 `Repository Binding` 后端模块
- `Product / Repository / Binding` 模块只能通过最小连接边界读取或校验 `Module`，不得吸收 `Module Registry` 或 `Decision Center` 的主线写入能力

---

## 3. 页面矩阵

### 3.1 页面主线

| 页面 | 职责 |
| --- | --- |
| `Product Registry / List` | 承接产品列表读取、筛选入口、创建入口与进入详情入口 |
| `Product Create` | 承接 `CreateProduct` |
| `Product Detail` | 承接产品详情读取、`BindModuleToProduct`、已绑定模块/仓库摘要读取与进入仓库绑定主线的上下文入口 |
| `Repository Binding / List` | 承接仓库列表读取、筛选入口、创建入口与进入绑定工作台入口 |
| `Repository Create` | 承接 `CreateRepository` |
| `Repository Binding Detail / Workspace` | 承接仓库详情读取、候选读取、`BindRepositoryToProduct` 与 `MapModuleToRepository` |

### 3.2 兼容跳转入口

- `Module Detail`：当前阶段只承接绑定摘要展示与兼容跳转入口，不扩写为第二个绑定工作台；可携带 `moduleId / moduleName / fromModuleDetail` 上下文进入正式主入口，但不得继续停留在本页直接提交绑定写入

> 约束：上述兼容入口不得扩写为当前阶段第二条页面主线，也不得并行拥有第二套主写入流程。

### 3.3 页面跳转关系

- `Product Registry / List` → `Product Create`
- `Product Registry / List` → `Product Detail`
- `Product Create` → `Product Detail`（`CreateProduct` 成功后默认回流到新建 `Product` 的详情页，携带继承来源标记）
- `Product Detail` → `Repository Binding Detail / Workspace`（进入仓库绑定主线的上下文入口，可携带 `productId / productName / fromProductDetail`）
- `Repository Binding / List` → `Repository Create`
- `Repository Binding / List` → `Repository Binding Detail / Workspace`
- `Repository Create` → `Repository Binding Detail / Workspace`（`CreateRepository` 成功后默认回流到新建 `Repository` 的工作台，携带继承来源标记）
- `Module Detail` 兼容跳转 → `Product Detail` 或 `Repository Binding / List`（携带 `moduleId / moduleName / fromModuleDetail`，目标实体已确定时额外携带 `productId` 或 `repositoryId`）

### 3.4 页面信息区块

#### 列表页最小信息区块

- **列表工具栏区**：承接搜索输入、状态筛选与进入 `Create` 的入口
- **列表内容区**：至少展示 `name / status / created_at` 与对应绑定计数
- **空状态区**：无记录时引导用户进入对应 `Create`

#### 创建页最小信息区块

- **结构化表单区**：承接 `Product` 或 `Repository` 的最小字段录入
- **来源上下文区**：从 `Module Detail` 或 `Product Detail` 带上下文进入时展示预填信息
- **提交取消操作区**：承接 `Create` 提交与按真实来源返回的路径

#### 详情页/工作台最小信息区块

- **核心字段区**：展示实体核心字段与状态
- **已绑定列表区**：展示当前已建立的对应绑定结果
- **候选读取及绑定动作区**：承接候选读取与对应绑定写入触点
- **上下文入口区**（仅 `Product Detail`）：承接进入仓库绑定主线的上下文入口

---

## 4. 动作矩阵

### 4.1 直接承接动作

| 动作 | 单值页面 owner | 写入数据载体 | canonical write owner（后端） | reread owner（后端） |
| --- | --- | --- | --- | --- |
| `CreateProduct` | `Product Create` | `products` | `ProductCreateWrite` | `ProductDetailRead` |
| `CreateRepository` | `Repository Create` | `repositories` | `RepositoryCreateWrite` | `RepositoryDetailRead` |
| `BindModuleToProduct` | `Product Detail` | `product_modules` | `ProductModuleBindingWrite` | `ProductDetailRead` |
| `BindRepositoryToProduct` | `Repository Binding Detail / Workspace` | `product_repositories` | `RepositoryProductBindingWrite` | `RepositoryDetailRead` |
| `MapModuleToRepository` | `Repository Binding Detail / Workspace` | `module_repositories` | `RepositoryModuleMappingWrite` | `RepositoryDetailRead` |

> 约束：
> - 五个核心动作的页面 owner 单值化，不得由 `Module Detail`、`Repository Binding` 或其他页面并行拥有第二套主写入流程。
> - `BindRepositoryToProduct` 与 `MapModuleToRepository` 的 reread owner 唯一为 `RepositoryDetailRead`；`ProductDetailRead` 不得成为 `BindRepositoryToProduct` 的第二 reread owner。
> - 三类绑定写入成功后必须回到对应 canonical owner 页面完成 reread，不得只靠 toast 作为成功依据。

### 4.2 允许最小入口但不扩写为独立主线

| 动作 | 承接方式 |
| --- | --- |
| `Decision -> Product` | 只保留合同保留位或轻量候选读取前提，不扩写为当前阶段正式写入主线 |
| `Decision -> Repository` | 只保留合同保留位或轻量候选读取前提，不扩写为当前阶段正式写入主线 |
| `Module Detail` 兼容跳转 | 只允许携带上下文进入正式主入口，不扩写为第二套主写入流程 |

### 4.3 非归属动作

- `CreateModule`、`CreateRelease`、`RecordDecision`、`LinkDecisionToTarget` 不在当前阶段 `Product Registry` 与 `Repository Binding` 后端模块中实现
- `Product / Repository / Binding` 模块不得调用或吸收 `Module Registry` 或 `Decision Center` 的主线写入能力

---

## 5. 数据模型

### 5.1 直接承接表

- `products`
- `repositories`
- `product_modules`
- `product_repositories`
- `module_repositories`

### 5.2 最小读取或校验前提

- `modules`（候选读取、目标存在性校验与已绑定摘要读取）

### 5.3 候选读取前提（只读，不要求写入主线）

- `decisions`（保留位，当前阶段不要求实现 `Decision -> Product / Repository` 候选读取）

### 5.4 对象最小字段

- **`Product`**：`id`、`name`、`description`、`status`、`created_at`
- **`Repository`**：`id`、`name`、`url`、`provider`、`status`、`created_at`
- **`ProductModuleBinding`**：`id`、`product_id`、`module_id`、`created_at`
- **`ProductRepositoryBinding`**：`id`、`product_id`、`repository_id`、`created_at`
- **`ModuleRepositoryMapping`**：`id`、`module_id`、`repository_id`、`created_at`

### 5.5 Product 最小模板

字段级冻结如下：

- 最小字段集合：`name / description / status`
- **创建必填**：`name / description / status`
- **创建可选**：无
- 不得额外引入 `customer / value_proposition / business_model / metrics / remote_import_source` 等超出 `v0.1` 的字段
- 空字符串不得视为合法必填值；写入前必须完成去首尾空白后的最小非空校验
- 不得用隐式默认值替代必填字段输入

### 5.6 Repository 最小模板

字段级冻结如下：

- 最小字段集合：`name / url / provider / status`
- **创建必填**：`name / url / provider / status`
- **创建可选**：无
- `provider` 为必填字符串字段，不采用受控枚举
- 不得额外引入 `oauth_binding / remote_import_status / sync_cursor / scanned_commit` 等自动化集成字段
- 空字符串不得视为合法必填值；写入前必须完成去首尾空白后的最小非空校验

### 5.7 status 枚举与状态语义

当前阶段 `status` 最小枚举冻结为：

| 状态 | 语义 |
| --- | --- |
| `active` | 可见、可继续绑定、可继续维护的有效状态 |
| `archived` | 已归档保留，仍允许作为历史事实被读取和展示，但不进入候选范围 |

> 约束：
> - 不得额外引入 `draft / syncing / disconnected / retired / imported` 等新状态。
> - 创建写入必须显式提交 `status`，`Create` 页面、HTTP DTO、`.proto` 写请求都必须携带 `status`。
> - "默认 `active`" 仅表示未发生用户改动时预填并显式提交 `active`，不解释为服务端静默补值或合同层隐式默认值。

### 5.8 statusFilter（UI 层枚举）

- UI/路由层枚举只允许 `all / active / archived`
- `all` 只允许存在于 UI 与路由搜索参数层
- `all` 不得写入数据库、HTTP 持久化 DTO、后端领域模型或 `.proto` 持久化字段
- `.proto` 中 `status_filter` 字段的 `ACTIVE_ARCHIVED_STATUS_UNSPECIFIED` 表示"不过滤"，对应 UI 的 `all`

### 5.9 三类绑定关系语义

| 关系 | 语义 | 表 | 唯一约束 |
| --- | --- | --- | --- |
| `ProductModuleBinding` | `Module M 被 Product P 使用` | `product_modules` | `(product_id, module_id)` 唯一 |
| `ProductRepositoryBinding` | `Repository R 是 Product P 的实现锚点` | `product_repositories` | `(product_id, repository_id)` 唯一 |
| `ModuleRepositoryMapping` | `Module M 实现于 Repository R` | `module_repositories` | `(module_id, repository_id)` 唯一 |

> 约束：
> - 不引入额外绑定属性字段（绑定权重、绑定备注等）。
> - `product_repositories` 为 `phase04` 阶段新增表，不存在 `phase02` 历史数据，不需要兼容回填。
> - `product_modules` 与 `module_repositories` 表已存在于 `phase02` migration 中，迁移后表结构不变，仅数据访问 owner 切换。

### 5.10 候选读取基线

- `BindModuleToProduct` 候选 = `status=active` 的 `Module`，排除已绑定到当前 `Product` 的 `Module`，按 `created_at` 降序
- `BindRepositoryToProduct` 候选 = `status=active` 的 `Product`，排除已绑定到当前 `Repository` 的 `Product`，按 `created_at` 降序
- `MapModuleToRepository` 候选 = `status=active` 的 `Module`，排除已映射到当前 `Repository` 的 `Module`，按 `created_at` 降序

候选展示模型：

- 候选 `Product` 每条至少 `product_id / product_name / product_status`
- 候选 `Module` 每条至少 `module_id / module_name / module_status`
- 已建立绑定关系的目标不得再次出现在可关联候选中
- 无可关联候选时，页面必须返回明确空状态，而不是把空结果误报为接口错误

### 5.11 最小读写模型

#### 列表读取

- `Product List` 至少承接：`name / description / status / created_at / module_bind_count / repository_bind_count`
- `Repository List` 至少承接：`name / url / provider / status / created_at / product_bind_count / module_bind_count`
- 列表筛选维度：两个列表均只冻结 `queryText / statusFilter`，`queryText` 只匹配 `name` 字段（模糊匹配）
- `Repository` 列表不引入 `providerFilter`
- 列表按 `created_at` 降序，不引入分页

#### 详情读取

- `Product Detail` 至少承接：核心对象字段 + 已绑定模块列表（`module_id / module_name / module_status`）+ 已绑定仓库列表（`repository_id / repository_name / provider / repository_status`）
- `Repository Detail / Workspace` 至少承接：核心对象字段 + 已绑定产品列表（`product_id / product_name / product_status`）+ 已映射模块列表（`module_id / module_name / module_status`）

#### 创建写入

- `CreateProduct`（最小字段 `name / description / status`）
- `CreateRepository`（最小字段 `name / url / provider / status`）

#### 绑定写入

- `BindModuleToProduct`（最小字段 `product_id / module_id`，`product_id` 由 `Product Detail` 上下文隐式承接，`module_id` 由候选列表选择）
- `BindRepositoryToProduct`（最小字段 `repository_id / product_id`，`repository_id` 由工作台上下文隐式承接，`product_id` 由候选选择）
- `MapModuleToRepository`（最小字段 `repository_id / module_id`，`repository_id` 由工作台上下文隐式承接，`module_id` 由候选选择）

### 5.12 绑定计数计算口径

- `module_bind_count`（`Product List` 项）：当前 `Product` 已建立的 `product_modules` 有效关联数
- `repository_bind_count`（`Product List` 项）：当前 `Product` 已建立的 `product_repositories` 有效关联数
- `product_bind_count`（`Repository List` 项）：当前 `Repository` 已建立的 `product_repositories` 有效关联数
- `module_bind_count`（`Repository List` 项）：当前 `Repository` 已建立的 `module_repositories` 有效关联数
- 当无任何已建立绑定时，对应计数返回 `0`，不返回 `null`

---

## 6. API 边界与接口分组

### 6.1 合同边界

- 遵守 `Contract First`，以结构化合同优先
- 跨语言合同标准冻结为 `Protocol Buffers`
- 当前阶段必须落地 `Product / Repository / Binding` 最小 `.proto` 合同源
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层，但不得形成与 `.proto` 并列的第二套合同源
- `chi + JSON HTTP` 的请求/响应语义必须从 `.proto` 派生或显式对齐；不得在 handler / HTTP DTO / 前端 adapter 中私自新增 `.proto` 中不存在的业务字段语义
- 不得引入与 `Protocol Buffers` 冲突的第二套跨语言合同主线

### 6.2 Product Registry 接口分组

#### 读组

| 接口 | 职责 |
| --- | --- |
| `ProductListRead` | 只承接列表展示所需最小读取、筛选前提与进入详情所需标识 |
| `ProductDetailRead` | 统一承接产品核心字段、已绑定模块列表、已绑定仓库列表与最小上下文入口读取 |
| `ProductModuleCandidateRead` | 只承接 `Product Detail` 中 `BindModuleToProduct` 的候选 `Module` 读取 |

#### 写组

| 接口 | 职责 |
| --- | --- |
| `ProductCreateWrite` | 只承接 `CreateProduct`，返回新建 `Product` 标识 |
| `ProductModuleBindingWrite` | 只承接 `BindModuleToProduct`，写入 `product_modules` |

#### 跨模块关系摘要读取（服务详情展示，owner = 关系表 owner）

| 接口 | owner | 数据访问文件 | 摘要字段 |
| --- | --- | --- | --- |
| `ProductModuleSummaryRead` | `Product Registry` | `productregistry/repository/binding_store.go` | `module_id / module_name / module_status` |
| `ProductRepositorySummaryRead` | `Repository Binding` | `repositorybinding/repository/binding_store.go` | `repository_id / repository_name / provider / repository_status` |

> 注：`ProductDetailRead` 通过注入 `ProductModuleSummaryRead` 与 `ProductRepositorySummaryRead` 接口承接跨模块读，不得直接读 `product_modules` / `product_repositories` 表。

#### 方向级最小 API 矩阵

| 动作语义 | 接口分组 | 接口名 | 方向 |
| --- | --- | --- | --- |
| 列表读取 | 读组 | `ProductListRead` | 读 |
| 详情读取 | 读组 | `ProductDetailRead` | 读 |
| `Module` 候选读取 | 候选读取 | `ProductModuleCandidateRead` | 读 |
| `CreateProduct` | 写组 | `ProductCreateWrite` | 写 |
| `BindModuleToProduct` | 写组 | `ProductModuleBindingWrite` | 写 |

### 6.3 Repository Binding 接口分组

#### 读组

| 接口 | 职责 |
| --- | --- |
| `RepositoryListRead` | 只承接列表展示所需最小读取、筛选前提与进入工作台所需标识 |
| `RepositoryDetailRead` | 统一承接仓库核心字段、已绑定产品列表、已映射模块列表与绑定工作台上下文读取 |
| `ProductBindingCandidateRead` | 只承接 `Repository Binding Detail / Workspace` 中 `BindRepositoryToProduct` 的候选 `Product` 读取 |
| `RepositoryModuleCandidateRead` | 只承接 `Repository Binding Detail / Workspace` 中 `MapModuleToRepository` 的候选 `Module` 读取 |

#### 写组

| 接口 | 职责 |
| --- | --- |
| `RepositoryCreateWrite` | 只承接 `CreateRepository`，返回新建 `Repository` 标识 |
| `RepositoryProductBindingWrite` | 只承接 `BindRepositoryToProduct`，写入 `product_repositories` |
| `RepositoryModuleMappingWrite` | 只承接 `MapModuleToRepository`，写入 `module_repositories` |

#### 跨模块关系摘要读取（服务详情展示，owner = 关系表 owner）

| 接口 | owner | 数据访问文件 | 摘要字段 |
| --- | --- | --- | --- |
| `RepositoryProductSummaryRead` | `Repository Binding` | `repositorybinding/repository/binding_store.go` | `product_id / product_name / product_status` |
| `RepositoryModuleSummaryRead` | `Repository Binding` | `repositorybinding/repository/binding_store.go` | `module_id / module_name / module_status` |

> 注：`RepositoryDetailRead` 通过注入 `RepositoryProductSummaryRead` 与 `RepositoryModuleSummaryRead` 接口承接跨模块读，不得直接读 `product_repositories` / `module_repositories` 表。

#### 方向级最小 API 矩阵

| 动作语义 | 接口分组 | 接口名 | 方向 |
| --- | --- | --- | --- |
| 列表读取 | 读组 | `RepositoryListRead` | 读 |
| 详情读取 | 读组 | `RepositoryDetailRead` | 读 |
| `Product` 候选读取 | 候选读取 | `ProductBindingCandidateRead` | 读 |
| `Module` 候选读取 | 候选读取 | `RepositoryModuleCandidateRead` | 读 |
| `CreateRepository` | 写组 | `RepositoryCreateWrite` | 写 |
| `BindRepositoryToProduct` | 写组 | `RepositoryProductBindingWrite` | 写 |
| `MapModuleToRepository` | 写组 | `RepositoryModuleMappingWrite` | 写 |

> 本矩阵与 `mvp_spec_v0.1.md` §6.2 最小 API 矩阵互链，仅承接 `Product / Repository / Binding` 当前阶段涉及的语义，不扩写为完整跨模块服务矩阵。

### 6.4 详情读取与候选读取边界

- `ProductDetailRead` / `RepositoryDetailRead` 只承接详情本体与已建立绑定结果
- 候选读取必须通过独立 request / response 承接
- 不得把候选读取结果直接并入详情读取的最小合同边界
- 不得把候选读取拆成需要前端自行拼装的多个独立业务入口
- 候选 `Module` / `Product` 读取必须各自独立于详情读取

### 6.5 跨模块读边界双轨

- 跨模块关系摘要读取（4 条，服务详情展示，落 `repository/binding_store.go`）与候选读取（3 条，服务绑定动作，落 `candidate/` 子包）是两类独立 Read 接口
- 两类接口必须通过独立 Read 接口隔离，不得混用或合并
- 跨模块关系摘要读取接口定义均落在模块根包 `types.go`（模块层级，不暴露 `repository/` 子包），实现落 `repository/binding_store.go`
- 跨模块关系摘要读取不得放入 `candidate/` 子包
- 具体接线必须在应用装配点（`backend/internal/platform/`）完成，`service/` 层不得直接写跨模块 SQL

### 6.6 关系写入后的读取语义

- `BindModuleToProduct` 提交成功后，默认后端语义为"`ProductDetailRead` 重新读取当前已绑定 `Module` 列表"
- `BindRepositoryToProduct` 提交成功后，默认后端语义为"`RepositoryDetailRead` 重新读取当前已绑定 `Product` 列表"
- `MapModuleToRepository` 提交成功后，默认后端语义为"`RepositoryDetailRead` 重新读取当前已映射 `Module` 列表"
- 不得额外设计一套脱离对应详情读组的并列回流读取接口

### 6.7 合同与存储解耦

- 不得直接将 `products`、`repositories`、`product_modules`、`product_repositories`、`module_repositories` 的存储模型原样暴露为外部合同
- `.proto` 必须作为当前阶段唯一合同源；若同时存在 `JSON HTTP` 过渡接口，其消息语义必须从 `.proto` 派生或与 `.proto` 严格对齐
- 必须允许后续在 `Contract First` 路线下独立演进接口消息结构
- 不得复用已删除字段编号或字段语义，必须保持与当前 `.proto` 合同主线兼容

---

## 7. 合同设计（Proto 主线）

### 7.1 单一 Proto 合同源

- `.proto` 是当前阶段唯一合同源
- 不得再以手写 JSON 结构充当并列合同源
- 三段式文件落点：
  - `proto/psco/common/v1/common.proto` → `package psco.common.v1`
  - `proto/psco/product_registry/v1/product_registry.proto` → `package psco.product_registry.v1`
  - `proto/psco/repository_binding/v1/repository_binding.proto` → `package psco.repository_binding.v1`
- `common.proto` 只承接跨 `Product / Repository` 共享且不会引入业务 owner 歧义的最小公共枚举
- `product_registry.proto` 与 `repository_binding.proto` 各自承接对应模块的服务接口与消息
- 不新增第二套 `buf.yaml` / `buf.gen.yaml` / `Makefile` 或并列 proto 根目录

### 7.2 包名与版本语义

- Go 包路径冻结：
  - `github.com/psco/backend/internal/gen/proto/psco/common/v1;commonv1`
  - `github.com/psco/backend/internal/gen/proto/psco/product_registry/v1;productregistryv1`
  - `github.com/psco/backend/internal/gen/proto/psco/repository_binding/v1;repositorybindingv1`
- TypeScript 生成产物继续落在 `frontend/src/gen/proto/psco/...`
- 后续新增字段必须在各自 `v1` 演进规则下进行，不临时改写包名

### 7.3 共享枚举与跨包复用

#### `ActiveArchivedStatus`（在 `psco.common.v1` 中单值定义）

| 枚举值 | 编号 |
| --- | --- |
| `ACTIVE_ARCHIVED_STATUS_UNSPECIFIED` | 0 |
| `ACTIVE_ARCHIVED_STATUS_ACTIVE` | 1 |
| `ACTIVE_ARCHIVED_STATUS_ARCHIVED` | 2 |

> 约束：
> - `Product` / `Repository` 自身的 `status`、列表项 `status`、详情/候选摘要 `status`、写组 request `status` 必须统一复用 `ActiveArchivedStatus`。
> - `Module` 摘要 `module_status` 通过 import 复用 `psco.module_registry.v1.ModuleStatus`。
> - `ACTIVE_ARCHIVED_STATUS_UNSPECIFIED` 在 `status_filter` 中表示"不过滤"对应 UI 的 `all`；在 `CreateProductRequest.status` / `CreateRepositoryRequest.status` 等写组 request 中由 service 层校验拒绝。

### 7.4 最小服务接口矩阵

#### `ProductRegistryService`

| RPC | 对齐接口 | 对齐页面动作 |
| --- | --- | --- |
| `ListProducts` | `ProductListRead` | `Product Registry / List` 列表读取 |
| `GetProductDetail` | `ProductDetailRead` | `Product Detail` 详情读取 |
| `CreateProduct` | `ProductCreateWrite` | `Product Create` 创建写入 |
| `BindModuleToProduct` | `ProductModuleBindingWrite` | `Product Detail` 绑定写入 |
| `ListProductModuleCandidates` | `ProductModuleCandidateRead` | `Product Detail` 候选读取 |

#### `RepositoryBindingService`

| RPC | 对齐接口 | 对齐页面动作 |
| --- | --- | --- |
| `ListRepositories` | `RepositoryListRead` | `Repository Binding / List` 列表读取 |
| `GetRepositoryDetail` | `RepositoryDetailRead` | `Repository Binding Detail / Workspace` 详情读取 |
| `CreateRepository` | `RepositoryCreateWrite` | `Repository Create` 创建写入 |
| `BindRepositoryToProduct` | `RepositoryProductBindingWrite` | `Repository Binding Detail / Workspace` 绑定写入 |
| `MapModuleToRepository` | `RepositoryModuleMappingWrite` | `Repository Binding Detail / Workspace` 映射写入 |
| `ListRepositoryProductCandidates` | `ProductBindingCandidateRead` | `Repository Binding Detail / Workspace` 候选读取 |
| `ListRepositoryModuleCandidates` | `RepositoryModuleCandidateRead` | `Repository Binding Detail / Workspace` 候选读取 |

### 7.5 核心消息结构与字段编号

#### Product Registry

##### `Product`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `id` | 1 | `string` | 产品标识 |
| `name` | 2 | `string` | 产品名称（必填） |
| `description` | 3 | `string` | 产品描述（必填） |
| `status` | 4 | `ActiveArchivedStatus` | 产品状态（必填） |
| `created_at` | 5 | `google.protobuf.Timestamp` | 创建时间 |

##### `ProductListItem`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `id` | 1 | `string` | 产品标识 |
| `name` | 2 | `string` | 产品名称 |
| `description` | 3 | `string` | 产品描述 |
| `status` | 4 | `ActiveArchivedStatus` | 产品状态 |
| `created_at` | 5 | `google.protobuf.Timestamp` | 创建时间 |
| `module_bind_count` | 6 | `int32` | 已绑定 `Module` 数 |
| `repository_bind_count` | 7 | `int32` | 已绑定 `Repository` 数 |

##### `BoundModuleSummary`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `module_id` | 1 | `string` | 模块标识 |
| `module_name` | 2 | `string` | 模块名称 |
| `module_status` | 3 | `ModuleStatus`（跨包 import） | 模块状态 |

##### `BoundRepositorySummary`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `repository_id` | 1 | `string` | 仓库标识 |
| `repository_name` | 2 | `string` | 仓库名称 |
| `provider` | 3 | `string` | 仓库提供者 |
| `repository_status` | 4 | `ActiveArchivedStatus` | 仓库状态 |

##### `ProductDetail`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `product` | 1 | `Product` | 产品核心对象 |
| `bound_modules` | 2 | `repeated BoundModuleSummary` | 已绑定 `Module` 列表 |
| `bound_repositories` | 3 | `repeated BoundRepositorySummary` | 已绑定 `Repository` 列表 |

##### `ProductModuleCandidate`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `module_id` | 1 | `string` | 候选模块标识 |
| `module_name` | 2 | `string` | 候选模块名称 |
| `module_status` | 3 | `ModuleStatus`（跨包 import） | 候选模块状态 |

##### Request / Response

| 消息 | 字段 | 编号 | 类型 |
| --- | --- | --- | --- |
| `ListProductsRequest` | `query_text` | 1 | `string` |
| `ListProductsRequest` | `status_filter` | 2 | `ActiveArchivedStatus` |
| `ListProductsResponse` | `products` | 1 | `repeated ProductListItem` |
| `GetProductDetailRequest` | `product_id` | 1 | `string` |
| `GetProductDetailResponse` | `product_detail` | 1 | `ProductDetail` |
| `CreateProductRequest` | `name` | 1 | `string` |
| `CreateProductRequest` | `description` | 2 | `string` |
| `CreateProductRequest` | `status` | 3 | `ActiveArchivedStatus` |
| `CreateProductResponse` | `product_id` | 1 | `string` |
| `BindModuleToProductRequest` | `product_id` | 1 | `string` |
| `BindModuleToProductRequest` | `module_id` | 2 | `string` |
| `BindModuleToProductResponse` | （空响应） | — | — |
| `ListProductModuleCandidatesRequest` | `product_id` | 1 | `string` |
| `ListProductModuleCandidatesResponse` | `candidates` | 1 | `repeated ProductModuleCandidate` |

#### Repository Binding

##### `Repository`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `id` | 1 | `string` | 仓库标识 |
| `name` | 2 | `string` | 仓库名称（必填） |
| `url` | 3 | `string` | 仓库 URL（必填） |
| `provider` | 4 | `string` | 仓库提供者（必填） |
| `status` | 5 | `ActiveArchivedStatus` | 仓库状态（必填） |
| `created_at` | 6 | `google.protobuf.Timestamp` | 创建时间 |

##### `RepositoryListItem`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `id` | 1 | `string` | 仓库标识 |
| `name` | 2 | `string` | 仓库名称 |
| `url` | 3 | `string` | 仓库 URL |
| `provider` | 4 | `string` | 仓库提供者 |
| `status` | 5 | `ActiveArchivedStatus` | 仓库状态 |
| `created_at` | 6 | `google.protobuf.Timestamp` | 创建时间 |
| `product_bind_count` | 7 | `int32` | 已绑定 `Product` 数 |
| `module_bind_count` | 8 | `int32` | 已映射 `Module` 数 |

##### `BoundProductSummary`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `product_id` | 1 | `string` | 产品标识 |
| `product_name` | 2 | `string` | 产品名称 |
| `product_status` | 3 | `ActiveArchivedStatus` | 产品状态 |

##### `MappedModuleSummary`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `module_id` | 1 | `string` | 模块标识 |
| `module_name` | 2 | `string` | 模块名称 |
| `module_status` | 3 | `ModuleStatus`（跨包 import） | 模块状态 |

##### `RepositoryDetail`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `repository` | 1 | `Repository` | 仓库核心对象 |
| `bound_products` | 2 | `repeated BoundProductSummary` | 已绑定 `Product` 列表 |
| `mapped_modules` | 3 | `repeated MappedModuleSummary` | 已映射 `Module` 列表 |

##### `RepositoryProductCandidate` / `RepositoryModuleCandidate`

| 消息 | 字段 | 编号 | 类型 |
| --- | --- | --- | --- |
| `RepositoryProductCandidate` | `product_id` | 1 | `string` |
| `RepositoryProductCandidate` | `product_name` | 2 | `string` |
| `RepositoryProductCandidate` | `product_status` | 3 | `ActiveArchivedStatus` |
| `RepositoryModuleCandidate` | `module_id` | 1 | `string` |
| `RepositoryModuleCandidate` | `module_name` | 2 | `string` |
| `RepositoryModuleCandidate` | `module_status` | 3 | `ModuleStatus`（跨包 import） |

##### Request / Response

| 消息 | 字段 | 编号 | 类型 |
| --- | --- | --- | --- |
| `ListRepositoriesRequest` | `query_text` | 1 | `string` |
| `ListRepositoriesRequest` | `status_filter` | 2 | `ActiveArchivedStatus` |
| `ListRepositoriesResponse` | `repositories` | 1 | `repeated RepositoryListItem` |
| `GetRepositoryDetailRequest` | `repository_id` | 1 | `string` |
| `GetRepositoryDetailResponse` | `repository_detail` | 1 | `RepositoryDetail` |
| `CreateRepositoryRequest` | `name` | 1 | `string` |
| `CreateRepositoryRequest` | `url` | 2 | `string` |
| `CreateRepositoryRequest` | `provider` | 3 | `string` |
| `CreateRepositoryRequest` | `status` | 4 | `ActiveArchivedStatus` |
| `CreateRepositoryResponse` | `repository_id` | 1 | `string` |
| `BindRepositoryToProductRequest` | `repository_id` | 1 | `string` |
| `BindRepositoryToProductRequest` | `product_id` | 2 | `string` |
| `BindRepositoryToProductResponse` | （空响应） | — | — |
| `MapModuleToRepositoryRequest` | `repository_id` | 1 | `string` |
| `MapModuleToRepositoryRequest` | `module_id` | 2 | `string` |
| `MapModuleToRepositoryResponse` | （空响应） | — | — |
| `ListRepositoryProductCandidatesRequest` | `repository_id` | 1 | `string` |
| `ListRepositoryProductCandidatesResponse` | `candidates` | 1 | `repeated RepositoryProductCandidate` |
| `ListRepositoryModuleCandidatesRequest` | `repository_id` | 1 | `string` |
| `ListRepositoryModuleCandidatesResponse` | `candidates` | 1 | `repeated RepositoryModuleCandidate` |

> 约束：排序规则不进入 `.proto` 合同本体，由 service / repository 层承接。

### 7.6 RPC → HTTP 映射

#### Product Registry

| RPC | HTTP |
| --- | --- |
| `ListProducts` | `GET /api/products` |
| `GetProductDetail` | `GET /api/products/{productId}` |
| `CreateProduct` | `POST /api/products` |
| `BindModuleToProduct` | `POST /api/products/{productId}/bindings/modules` |
| `ListProductModuleCandidates` | `GET /api/products/{productId}/candidates/modules` |

#### Repository Binding

| RPC | HTTP |
| --- | --- |
| `ListRepositories` | `GET /api/repositories` |
| `GetRepositoryDetail` | `GET /api/repositories/{repositoryId}` |
| `CreateRepository` | `POST /api/repositories` |
| `BindRepositoryToProduct` | `POST /api/repositories/{repositoryId}/bindings/products` |
| `MapModuleToRepository` | `POST /api/repositories/{repositoryId}/bindings/modules` |
| `ListRepositoryProductCandidates` | `GET /api/repositories/{repositoryId}/candidates/products` |
| `ListRepositoryModuleCandidates` | `GET /api/repositories/{repositoryId}/candidates/modules` |

> 约束：
> - handler 必须在进入业务层前显式组装为对应 Proto request 字段；`product_id / repository_id / module_id` 在 Proto RPC 中必须作为请求消息的显式字段存在。
> - HTTP 路径、状态码或中间件策略不得误写成 Proto 合同本体；差异视为传输层差异而非并列合同定义。

### 7.7 详情读取与候选读取消息边界

- `GetProductDetailResponse` / `GetRepositoryDetailResponse` 不得内嵌候选读取结果
- `.proto` 只定义摘要消息结构，不定义跨模块关系摘要读取接口的 RPC
- 四个跨模块关系摘要读取接口是后端模块内部接口，不直接暴露为 RPC
- 前端只通过 `GetProductDetail` 与 `GetRepositoryDetail` RPC 获取已绑定列表

### 7.8 合同演进与 Buf 校验

#### 字段与枚举演进规则

- 新增字段或枚举值必须使用新的递增编号
- 不得插入到已有编号之间
- 不得复用已删除字段编号或枚举值编号

#### 删除字段或枚举值后的 reserved 约束

- 必须使用 `reserved` 保留编号
- 必要时同时 `reserved` 对应名称
- 不得在未声明 `reserved` 的情况下复用旧编号或旧名称

#### v1 breaking 升级边界

- 删除已有字段 / 修改已有字段类型/编号/JSON/wire 兼容语义 / 删除已有 RPC 或修改其 request/response 兼容边界 → 视为 v1 breaking 变更
- 不得直接在现有 v1 包内覆盖
- 必须通过新版本目录与包名（如 `v2`）承接

#### Buf 校验链冻结

- 必须至少覆盖 `buf build` / `buf lint` / `buf generate` / `buf breaking`
- 必须复用仓库现有 `proto/Makefile`
- `buf breaking` 必须继续对照 `../.git#branch=main,subdir=proto`
- 不得吞掉 `buf breaking` 的失败退出码
- 当前阶段新增 `common / product_registry / repository_binding` 三个 `.proto` 文件属于向后兼容新增，不构成 breaking

#### 当前阶段合同落地边界

- 必须冻结 `.proto` 合同源、最小消息结构、服务接口、字段编号、RPC → HTTP 映射与生成/校验入口
- 当前阶段可不完成完整 gRPC / Connect / 网关迁移
- 当前阶段可继续保留 `chi` 作为 HTTP 过渡传输层
- 当前阶段不要求立即用生成类型替换全部手写 DTO

### 7.9 错误语义承接

- `.proto` 合同设计必须保留校验失败/资源不存在/重复冲突/空结果的承接空间
- 不把 HTTP 状态码本身写进 `.proto`
- 候选读取空结果必须表现为正常空列表响应，不得设计为空错误

---

## 8. 空状态与冷启动

### 8.1 空状态引导

- 用户首次进入 `Product Registry / List` 且系统中尚无任何 `Product` 时，页面必须展示明确的空状态提示
- 空状态主动作必须直接进入 `Product Create`
- 空状态文案必须围绕"先完成首个产品登记"展开
- 用户首次进入 `Repository Binding / List` 且系统中尚无任何 `Repository` 时，页面必须展示明确的空状态提示
- 空状态主动作必须直接进入 `Repository Create`
- 不得把导入、自动扫描或 AI 建议写成空状态主入口

### 8.2 冷启动路径

- 首轮必须允许用户从空状态进入 `Product Create`
- 首轮必须允许用户从空状态进入 `Repository Create`
- 首轮必须允许用户完成首个 `Product -> Repository` 绑定
- 首轮必须允许用户将已存在 `Module` 绑定到 `Product`
- 首轮必须允许用户将已存在 `Module` 映射到 `Repository`
- 首轮必须允许用户从 `Module Detail` 的兼容入口跳转进入正式绑定主入口，并完成至少一条绑定写入
- 当前阶段不依赖导入、自动扫描或 AI 推荐

### 8.3 创建最小闭环

- `Product` 最小闭环为 `Product Registry / List → Product Create → Product Detail`
- `Repository` 最小闭环为 `Repository Binding / List → Repository Create → Repository Binding Detail / Workspace`
- `CreateProduct` 成功后默认回流到新建 `Product` 的 `Product Detail`，携带继承来源标记
- `CreateRepository` 成功后默认回流到新建 `Repository` 的 `Repository Binding Detail / Workspace`，携带继承来源标记
- 不并列保留"成功后默认回列表"的第二套路径
- 绑定动作成功后用户必须停留在当前 canonical owner 页面，并重新读取对应已绑定列表完成 reread

### 8.4 多入口回流矩阵

#### `Product Create` 来源上下文

- `fromList` / `fromModuleDetail` / `direct-entry`（不得并行持有两个主来源）

#### `Repository Create` 来源上下文

- `fromList` / `fromProductDetail` / `fromModuleDetail` / `direct-entry`（不得并行持有两个主来源）

#### `Product Detail` 来源上下文

- `fromList` / `fromModuleDetail` / `direct-entry`（从 `ProductCreatePage` 成功创建后进入时来源上下文必须继承自创建页）

#### `Repository Binding Detail` 来源上下文

- `fromList` / `fromProductDetail` / `fromModuleDetail` / `direct-entry`（从 `RepositoryCreatePage` 成功创建后进入时来源上下文必须继承自创建页）

#### 返回路径规则（按真实来源决定，不统一伪造成回列表）

| 来源标记 | 返回路径 |
| --- | --- |
| `fromList` | 返回对应 List 并携带原 `queryText / statusFilter` |
| `fromModuleDetail` | 返回原 `ModuleDetailPage` |
| `fromProductDetail` | 返回原 `ProductDetailPage`（仅 `Repository Create` / `Repository Binding Detail` 适用） |
| `direct-entry` | 返回对应 List 默认筛选参数（`queryText=空`、`statusFilter=all`） |

> 约束：
> - `fromList` 显式建模"来源列表上下文存在/不存在"；存在时返回列表的导航必须携带原 `queryText / statusFilter`；不存在时必须落默认筛选参数。
> - 不得在无来源列表上下文时伪造或套用上一份筛选条件。
> - `fromList` 只作为来源标记参与返回导航的参数拼接判定，不得作为独立恢复机制绕过路由搜索参数。
> - 列表查询条件唯一事实源是路由搜索参数；不得引入 `sessionStorage` 或其他持久化层作为缺参回退源。

### 8.5 刷新恢复

- 直接刷新列表页时必须以当前路由搜索参数中的 `queryText / statusFilter` 作为查询条件
- 路由搜索参数缺失时使用默认筛选参数
- 不得在刷新后回退到任何持久化层中上一次的筛选值
- 无参 URL（`/products`、`/repositories`）必须稳定表现为默认筛选参数
- 带 `fromModuleDetail` / `fromProductDetail` 的页面刷新后必须继续恢复来源标记，不得静默丢失

### 8.6 Module Detail 兼容跳转

- `Module Detail` 仅允许兼容跳转，跳转参数冻结为 `moduleId / moduleName / fromModuleDetail`
- 上下文参数只表示来源模块身份与来源页面标记，不表示目标实体身份
- 目标实体未确定时先进入列表页（`/products` 或 `/repositories`）选择目标
- 目标实体已确定时额外携带目标页身份参数 `productId` 或 `repositoryId`，与上下文参数拆开传递
- 接收方页面基于上下文参数预填对应绑定面板的候选 `Module` 选择
- `Module Detail` 内 `ModuleBindingPanel` 从写入承接位回落为只读摘要 + 兼容跳转入口

---

## 9. 前端实现设计层结果

### 9.1 页面文件落点

| 页面 | 文件路径 |
| --- | --- |
| `ProductListPage` | `frontend/src/features/product-registry/pages/product-list-page.tsx` |
| `ProductCreatePage` | `frontend/src/features/product-registry/pages/product-create-page.tsx` |
| `ProductDetailPage` | `frontend/src/features/product-registry/pages/product-detail-page.tsx` |
| `RepositoryBindingListPage` | `frontend/src/features/repository-binding/pages/repository-binding-list-page.tsx` |
| `RepositoryCreatePage` | `frontend/src/features/repository-binding/pages/repository-create-page.tsx` |
| `RepositoryBindingDetailPage` | `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx` |

### 9.2 路由树与 URL 语义

#### Product Registry

| 路由 | URL | 文件路径 |
| --- | --- | --- |
| `ProductListRoute` | `/products` | `frontend/src/routes/products/index.tsx` |
| `ProductCreateRoute` | `/products/new` | `frontend/src/routes/products/new.tsx` |
| `ProductDetailRoute` | `/products/:productId` | `frontend/src/routes/products/$productId.tsx` |

#### Repository Binding

| 路由 | URL | 文件路径 |
| --- | --- | --- |
| `RepositoryBindingListRoute` | `/repositories` | `frontend/src/routes/repositories/index.tsx` |
| `RepositoryCreateRoute` | `/repositories/new` | `frontend/src/routes/repositories/new.tsx` |
| `RepositoryBindingDetailRoute` | `/repositories/:repositoryId` | `frontend/src/routes/repositories/$repositoryId.tsx` |

> 约束：
> - 不允许把 `Detail` 拆成独立子路由工作台。
> - 不引入与 `TanStack Router` 文件路由约定冲突的第二套路由体系。

### 9.3 组件树

#### Product List

- `ProductListPageShell`
  - `ProductListToolbar`（承接筛选入口与创建入口）
  - `ProductListContent`（只承接产品列表读取结果）
  - `ProductListTableOrCards`
  - `ProductListEmptyState`

#### Product Create

- `ProductCreatePageShell`
  - `ProductCreateForm`（只承接 `name / description / status` 最小录入）
  - `ProductCreateActions`

#### Product Detail

- `ProductDetailPageShell`
  - `ProductSummaryCard`（只承接产品核心字段与状态表达）
  - `ProductBoundModuleListSection`（含 `ProductModuleBindingPanel`）
  - `ProductBoundRepositoryListSection`
  - `ProductRepositoryBindingEntry`（承接进入仓库绑定主线的上下文入口）

#### Repository Binding List

- `RepositoryBindingListPageShell`
  - `RepositoryBindingListToolbar`（承接筛选入口与创建入口）
  - `RepositoryBindingListContent`（只承接仓库列表读取结果）
  - `RepositoryBindingListTableOrCards`
  - `RepositoryBindingListEmptyState`

#### Repository Create

- `RepositoryCreatePageShell`
  - `RepositoryCreateForm`（只承接 `name / url / provider / status` 最小录入）
  - `RepositoryCreateActions`

#### Repository Binding Detail / Workspace

- `RepositoryBindingDetailPageShell`
  - `RepositorySummaryCard`（只承接仓库核心字段与状态表达）
  - `RepositoryBoundProductListSection`（含 `RepositoryProductBindingPanel`）
  - `RepositoryMappedModuleListSection`（含 `RepositoryModuleMappingPanel`）

#### 组件归属原则

- 默认归属于当前 Page，确有跨页复用证据才抽为共享组件
- 当前阶段不得为了"组件纯洁"提前拆出无明确复用证据的通用组件层

### 9.4 状态模型

#### Product List / Repository List

- 查询条件冻结到路由搜索参数层：`queryText`、`statusFilter` 为唯一事实源
- 默认值：`queryText=空`、`statusFilter=all`
- 从 `Create` 或 `Detail` 返回时，必须按原有 `queryText` 与 `statusFilter` 恢复列表上下文
- 读取状态：`pending`、`success`、`error`
- 派生视图状态：`initial-loading`、`ready`、`empty`、`error`
- `ready` 与 `empty` 必须由成功读取后的数据是否为空派生
- 错误反馈停留在列表页内容区域，不跳转独立错误页

#### Product Create / Repository Create

- 草稿状态：`idle`、`dirty`
- `status` 默认预填 `active`
- 提交状态：`submitting`、`submit-success`、`submit-error`
- 提交状态不得反向覆盖用户草稿值
- 提交失败时停留当前页，已输入草稿原样保留，来源上下文继续保留，错误显示在表单上下文
- 提交成功默认回流到新建实体的详情/工作台页，携带继承来源标记
- 取消返回按真实来源决定返回路径，不统一伪造成回列表路径

#### Product Detail

- 详情读取状态：`pending`、`success`、`error`
- 资源不存在时派生 `not-found` 视图状态
- `ProductModuleBindingPanel` 候选读取状态：`closed`、`pending`、`ready`、`empty`、`error`
- `empty` 必须表示候选列表为空而非接口失败或资源不存在
- 候选 `Module` 读取必须独立于详情读取
- `ProductModuleBindingPanel` 写入状态：`idle`、`submitting`、`submit-success`、`submit-error`，只归属于当前详情页上下文
- `BindModuleToProduct` 提交成功 reread：用户必须停留在当前 `ProductDetailPage`，必须重新读取已绑定 `Module` 列表完成 reread，面板必须回到 `closed`
- `BindModuleToProduct` 提交失败：停留当前页，当前已选候选 `Module` 必须继续保留，错误停留在面板上下文，重复绑定不得降级为静默成功

#### Repository Binding Detail / Workspace

- 详情读取状态：`pending`、`success`、`error`
- 资源不存在时派生 `not-found` 视图状态
- 两类绑定面板状态：候选读取 `closed / pending / ready / empty / error`，写入 `idle / submitting / submit-success / submit-error`
- 同一时刻只允许一个绑定面板处于打开态（互斥展开）
- 候选 `Product` 与候选 `Module` 读取必须各自独立于详情读取
- `BindRepositoryToProduct` 提交成功 reread：用户必须停留在当前 `RepositoryBindingDetailPage`，必须重新读取已绑定 `Product` 列表完成 reread，当前活动面板必须回到 `closed`
- `MapModuleToRepository` 提交成功 reread：用户必须停留在当前 `RepositoryBindingDetailPage`，必须重新读取已映射 `Module` 列表完成 reread，当前活动面板必须回到 `closed`
- 两类绑定提交失败：停留当前页，当前已选候选目标继续保留，错误停留在对应面板上下文，重复绑定不得降级为静默成功

#### 页面级 UI 状态局部归属原则

- 列表与创建页局部状态：搜索输入草稿、表单草稿、提交错误与空状态展示应优先归属于当前页面，不默认升级为跨路由全局状态
- 详情页局部状态：候选读取状态、当前选中候选、活动面板状态、提交错误与来源上下文展示应优先归属于当前详情页上下文
- 派生视图状态应优先由当前读模型结果计算，不被重复持久化为独立全局字段

### 9.5 布局降级策略

#### PC 页面布局

- `ProductListPage` / `RepositoryBindingListPage`：高信息密度列表布局，工具栏与列表同屏
- `ProductDetailPage` / `RepositoryBindingDetailPage`：分区式详情布局，核心信息与绑定结果同屏可见
- `ProductCreatePage` / `RepositoryCreatePage`：单列垂直布局，主动作按钮无需横向滚动即可见

#### 移动浏览器布局

- `ProductListPage` / `RepositoryBindingListPage`：单列列表或卡片重排，信息裁剪保留核心字段
- `ProductDetailPage` / `RepositoryBindingDetailPage`：摘要、已绑定列表、候选与绑定动作按垂直顺序重排，次级信息可折叠
- `ProductCreatePage` / `RepositoryCreatePage`：单列垂直布局

> 约束：不得新增独立移动端页面体系；不得引入独立 `React Native` 客户端；不得把完整 `PWA` 能力写成当前阶段实现前提。

### 9.6 运行时实现细节不冻结

- 当前阶段不提前冻结具体 hook 命名、Query key 细节、store API 命名、缓存时间、请求取消策略或 optimistic update 方案
- 页面级 UI 状态（草稿、面板开闭、提交错误等瞬时状态）优先归属于当前页面或当前详情页上下文，只在确有跨同页多组件共享需要时再抽为页面作用域状态容器，不得默认升级为跨路由全局状态
- 不得引入 `sessionStorage` 或其他持久化层作为列表查询条件的事实源或缺参回退源

---

## 10. 后端实现设计层结果

### 10.1 模块边界

- `Product Registry` 在后端冻结为单一业务模块
- 统一承接 `CreateProduct`、`BindModuleToProduct`、`ProductListRead`、`ProductDetailRead`、`ProductModuleCandidateRead`、`ProductModuleSummaryRead`
- `Repository Binding` 在后端冻结为单一业务模块
- 统一承接 `CreateRepository`、`BindRepositoryToProduct`、`MapModuleToRepository`、`RepositoryListRead`、`RepositoryDetailRead`、`ProductBindingCandidateRead`、`RepositoryModuleCandidateRead`、`ProductRepositorySummaryRead`、`RepositoryProductSummaryRead`、`RepositoryModuleSummaryRead`
- 不得把这些能力拆散到 `Module`、`Product`、`Repository` 之外的其他后端模块中

### 10.2 文件落点

#### Product Registry

```
backend/internal/productregistry/
  handler/
    query_handler.go              # 读组入口：ProductListRead + ProductDetailRead
    command_handler.go            # 写组入口：ProductCreateWrite + ProductModuleBindingWrite
  service/
    query_service.go              # 读组编排（ProductDetailRead 通过注入 ProductModuleSummaryRead / ProductRepositorySummaryRead 接口承接跨模块读）
    command_service.go            # 写组编排（创建 + 绑定）
  repository/
    product_store.go              # products 表读写
    binding_store.go              # product_modules 关系写入 + ProductModuleSummaryRead 实现
  candidate/
    module_candidate_read.go      # ProductModuleCandidateRead
  errors.go                       # 模块错误语义收口
  types.go                        # 模块级类型与跨模块关系摘要读取接口定义
  validate.go                     # 模块输入校验
  handler/
    response.go                   # 统一响应组装
```

#### Repository Binding

```
backend/internal/repositorybinding/
  handler/
    query_handler.go              # 读组入口：RepositoryListRead + RepositoryDetailRead
    command_handler.go            # 写组入口：RepositoryCreateWrite + RepositoryProductBindingWrite + RepositoryModuleMappingWrite
  service/
    query_service.go              # 读组编排（RepositoryDetailRead 通过注入 RepositoryProductSummaryRead / RepositoryModuleSummaryRead 接口承接跨模块读）
    command_service.go            # 写组编排（创建 + 两类绑定）
  repository/
    repository_store.go           # repositories 表读写
    binding_store.go              # product_repositories + module_repositories 关系写入 + 三个跨模块关系摘要读取接口实现
  candidate/
    product_candidate_read.go     # ProductBindingCandidateRead
    module_candidate_read.go      # RepositoryModuleCandidateRead
  errors.go                       # 模块错误语义收口
  types.go                        # 模块级类型与跨模块关系摘要读取接口定义
  validate.go                     # 模块输入校验
  handler/
    response.go                   # 统一响应组装
```

### 10.3 分层语义

| 层 | 子包 | 职责 |
| --- | --- | --- |
| 入口层 | `handler/` | 只负责承接请求与返回结果 |
| 业务编排层 | `service/` | 只负责动作语义、校验顺序与跨连接口编排 |
| 数据访问/外部连接层 | `repository/` + `candidate/` | 只负责持久化与跨模块依赖调用 |

### 10.4 读组与写组代码组织

- 读组编排必须落在单文件 `service/query_service.go`，不得拆为 `list_service.go` / `detail_service.go` / `candidate_service.go`
- 写组编排必须落在单文件 `service/command_service.go`，不得按动作拆为多个独立 service 文件
- 支撑文件 `errors.go` / `types.go` / `validate.go` / `handler/response.go` 按职责单值化映射到唯一文件，禁止散落在 handler 或 service 内部

### 10.5 候选读取接口拥有者与接线原则

- `candidate/` 子包自己定义和拥有接口与实现
- `Module Registry` 不为 `Product Registry` 或 `Repository Binding` 暴露专门的服务契约或服务方法
- 具体接线必须在应用装配点（`backend/internal/platform/`）完成
- `service/` 层不得直接写跨模块 SQL

### 10.6 Module Registry 连接边界

- 仅最小候选读取、存在性校验与展示摘要读取
- 不承接 `CreateModule` / `CreateRelease` 等主线写入

### 10.7 Decision Center 连接边界

- 当前阶段不接入正式主线
- 不得为 `Decision -> Product / Repository` 设计正式候选/详情/绑定接口

### 10.8 不冻结的实现工具

- 不得提前冻结 Go HTTP 框架、RPC 框架、ORM、SQL Builder、Repository 模板或目录生成器
- 文件内部的函数签名、中间件选型、数据库访问方式不作为当前阶段冻结内容
- `chi` 路由组织方式只作为实现兼容前提，不是新的合同源
- 只冻结职责分工、接口归属与文件落点，不冻结实现手段

---

## 11. 迁移边界与兼容策略

### 11.1 products / repositories 主线升级 migration

- migration 文件必须落在 `database/migrations/0006_product_repository_binding_mainline.sql`
- `products` 必须通过 `ALTER TABLE` 原位新增 `description`（`TEXT`）与 `status`（`TEXT NOT NULL DEFAULT 'active'`，`CHECK (status IN ('active', 'archived'))`）
- `repositories` 必须通过 `ALTER TABLE` 原位新增 `url`（`TEXT`）、`provider`（`TEXT`）与 `status`（`TEXT NOT NULL DEFAULT 'active'`，`CHECK (status IN ('active', 'archived'))`）
- 回填后 `products.description` 与 `repositories.url / provider` 必须进入 `NOT NULL`
- 必须为列表读取性能添加索引：`products(status, created_at DESC)` 与 `repositories(status, created_at DESC)`
- 不得删除既有 `id / name / created_at`
- 不得新建 `products_v2`、`repositories_v2` 或并行影子绑定表

### 11.2 product_repositories 关系表结构

- 表结构至少必须包含 `id / product_id / repository_id / created_at`
- `product_id` 必须引用 `products(id) ON DELETE RESTRICT`
- `repository_id` 必须引用 `repositories(id) ON DELETE RESTRICT`
- `(product_id, repository_id)` 必须唯一
- 必须新增 `product_id` 与 `repository_id` 方向的读取索引
- 不得在当前阶段引入额外绑定属性字段
- 该表不存在 `phase02` 历史数据，不需要兼容回填

### 11.3 历史数据兼容回填

#### Product 历史数据回填

- `description` 必须回填为 `'（历史产品，phase04 升级前无描述）'`
- `status` 必须回填或默认落为 `active`
- 回填必须保留原有 `id / name / created_at`
- 回填语句必须具备幂等性

#### Repository 历史数据回填

- `url` 必须回填为 `'https://example.com/legacy'`
- `provider` 必须回填为 `'legacy'`
- `status` 必须回填或默认落为 `active`
- 回填必须保留原有 `id / name / created_at`
- 回填语句必须具备幂等性

#### phase02 历史绑定数据兼容

- 既有 `product_modules` 与 `module_repositories` 历史数据必须保持可读
- `phase04` 不得通过重建这些历史关系表来完成 owner 迁移
- `phase04` 只新增 `product_repositories` 关系，不回写第二套旧绑定表

### 11.4 旧 transport 兼容委派策略

#### `ProductBindingCandidateRead`

- canonical owner 切换为 `Repository Binding` 读组
- 旧 transport 入口最多作为兼容适配层委派给 canonical 实现
- 不得继续由 `Module Registry` 业务 service 保留长期 owner

#### `RepositoryBindingCandidateRead`

- 已废弃，不保留新 canonical 对应接口
- 旧 transport 入口若保留也不得承接新业务实现或新增消费方

#### `ModuleBindingWrite`

- 拆分迁移为 `ProductModuleBindingWrite`（`Product Registry`）与 `RepositoryModuleMappingWrite`（`Repository Binding`）
- 旧模块中心写入入口若保留只能作为兼容适配层委派

#### `binding_store.go` 拆分迁移

- `product_modules` 数据访问迁到 `Product Registry` 的 `repository/binding_store.go`
- `module_repositories` 数据访问迁到 `Repository Binding` 的 `repository/binding_store.go`
- `product_repositories` 数据访问作为 `Repository Binding` 新增能力

#### `Module Detail` 旧入口后端 endpoint

- `phase04` 后端不得为 `Module Detail` 提供新的绑定写入 endpoint
- 旧 endpoint 若保留只能作为兼容适配层委派给新 canonical write owner

#### `.proto` 合同源层面

- `phase04-08` `.proto` 合同源层面不主动删除 `module_registry.proto` 中的旧 RPC
- 旧 RPC 视为兼容适配入口，避免触发 `buf breaking` 失败

### 11.5 历史绑定数据兼容原则

- `phase02` 中已存在的 `product_modules` 与 `module_repositories` 历史数据必须保持可读
- 不得通过重建影子表、第二套绑定表或临时双写绕过迁移
- 表结构不变，仅数据访问 owner 切换

---

## 12. 错误语义

### 12.1 创建校验失败

| 场景 | 语义 | 归属接口 |
| --- | --- | --- |
| `CreateProduct` 必填字段缺失（含 `status` 缺失） | 校验失败 | `ProductCreateWrite` |
| `CreateProduct` 字段值非法（含非法 `status`） | 校验失败 | `ProductCreateWrite` |
| `CreateRepository` 必填字段缺失（含 `status` / `provider` 缺失） | 校验失败 | `RepositoryCreateWrite` |
| `CreateRepository` 字段值非法（含非法 `status`） | 校验失败 | `RepositoryCreateWrite` |

> 约束：不得降级为模糊通用错误；不得出现 `500` 级未收口错误替代业务错误。

### 12.2 资源不存在

| 场景 | 归属接口 |
| --- | --- |
| `ProductDetailRead` 接收到不存在的 `product_id` | `ProductDetailRead` |
| `RepositoryDetailRead` 接收到不存在的 `repository_id` | `RepositoryDetailRead` |
| `BindModuleToProduct` 的目标 `Product` 或 `Module` 不存在 | `ProductModuleBindingWrite` |
| `BindRepositoryToProduct` 的目标 `Product` 或 `Repository` 不存在 | `RepositoryProductBindingWrite` |
| `MapModuleToRepository` 的目标 `Module` 或 `Repository` 不存在 | `RepositoryModuleMappingWrite` |
| `ProductModuleCandidateRead` 接收到不存在的 `product_id` | `ProductModuleCandidateRead` |
| `ProductBindingCandidateRead` 接收到不存在的 `repository_id` | `ProductBindingCandidateRead` |
| `RepositoryModuleCandidateRead` 接收到不存在的 `repository_id` | `RepositoryModuleCandidateRead` |

### 12.3 重复冲突

| 场景 | 归属接口 |
| --- | --- |
| `BindModuleToProduct` 重复绑定 | `ProductModuleBindingWrite` |
| `BindRepositoryToProduct` 重复绑定 | `RepositoryProductBindingWrite` |
| `MapModuleToRepository` 重复映射 | `RepositoryModuleMappingWrite` |

> 约束：必须返回明确的重复冲突语义，不得降级为静默成功，不得降级为模糊通用错误，不得通过 `ON CONFLICT DO NOTHING` 隐式吞掉重复冲突。

### 12.4 候选空

- 任一候选读取接口返回零条候选记录 → 必须返回空列表语义
- 不得将空结果映射为资源不存在
- 不得将空结果映射为接口错误
- 不得把空候选结果误报为接口失败
- 页面必须展示明确的无可绑定候选空状态提示
- `.proto` 层承接：候选读取的空结果必须表现为正常空列表响应，不得设计为空错误

### 12.5 列表空

- `ProductListRead` 或 `RepositoryListRead` 返回零条记录 → 必须返回空列表语义
- 不得将空结果映射为资源不存在或接口错误
- 页面必须展示明确的空状态引导
- 空状态主动作必须直接进入对应 `Create` 页面，不得把导入、自动扫描或 AI 建议写成空状态主入口

### 12.6 详情不存在

- `ProductDetailRead` / `RepositoryDetailRead` 接收到不存在的资源 ID → 必须返回资源不存在语义
- 前端派生视图状态：`ProductDetailPage` / `RepositoryBindingDetailPage` 资源不存在时必须派生 `not-found` 视图状态
- 错误反馈必须停留在详情页/工作台内容区域，不得跳转独立错误页

### 12.7 三类绑定写入校验失败类型单值矩阵

| 失败类型 | 语义 |
| --- | --- |
| 必填字段缺失（`product_id / module_id / repository_id` 缺失） | 校验失败 |
| 目标实体不存在 | 资源不存在 |
| 目标实体非 `active` 状态 | 校验失败 |
| 重复绑定 | 重复冲突 |

> 约束：不得出现 `500` 级未收口错误替代上述业务错误。

### 12.8 非目标（错误语义相关）

- 不提前冻结 Dashboard 聚合接口、`product_asset_coverage` 或类似跨模块聚合读取接口
- 不提前冻结超出当前阶段的分页、复杂检索或批量操作接口
- 不提前冻结 `Decision -> Product / Repository` 正式关联写入接口

---

## 13. 验收基线

### 13.1 联调环境可重复建立

- 必须继续复用现有单一数据库与脚本主线：`database/scripts/init_db.sh`、`database/scripts/run_seeds.sh`、`database/scripts/reset_module_mainline.sh`
- 必须新增 `database/scripts/reset_product_repository_mainline.sh` 作为 `phase04` 主线重置入口
- 不得新建第二个数据库、第二套 `init_db` 入口或第二套 `phase04` 专用 seed runner
- 启动顺序必须冻结为：`init_db.sh` → 后端启动并自动执行 migration（含 `0006_product_repository_binding_mainline.sql`）→ `run_seeds.sh` → `reset_module_mainline.sh` → `reset_product_repository_mainline.sh` → 前端启动
- `run_seeds.sh` 必须发生在 `0006` migration 完成之后
- `reset_product_repository_mainline.sh` 必须在 `reset_module_mainline.sh` 之后执行
- 当前阶段不要求 `reset_decision_mainline.sh` 成为 `phase04` 验收前置步骤

### 13.2 重置脚本

- 脚本必须落在 `database/scripts/reset_product_repository_mainline.sh`
- 必须支持 `--clean-only`、`--restore-only` 与默认模式（清空 + 恢复）
- 必须复用既有 `resolve_psql` 与环境变量覆盖模式
- 必须支持 `-h / --help`
- 清空范围必须覆盖 `product_repositories`、`product_modules`、`module_repositories`、`products`、`repositories`
- 不得清空 `modules / module_releases / decisions / decision_links`
- 清空必须按以下受控顺序执行：先 `product_repositories`，再 `product_modules` 与 `module_repositories`，最后 `products` 与 `repositories`
- 必须继续沿用 `DELETE FROM ...` 的受控清空模式，不得切换为 `TRUNCATE ... RESTART IDENTITY ... CASCADE` 作为默认方案
- 必须校验目标数据库已存在
- 在 `--restore-only` 与默认模式下，必须校验 `modules` 基线已存在；若不存在必须提示先执行 `reset_module_mainline.sh`

### 13.3 基线 seed 与 fixture

- seed 文件必须落在 `database/seeds/seed_product_repository_mainline_baseline.sql`
- 必须通过 `BEGIN / COMMIT` 包裹
- 必须同时承担"清空 + 恢复"职责，默认脚本直接执行整份 SQL
- `--restore-only` 模式必须通过跳过清空语句并依赖幂等 `INSERT` 恢复基线
- 至少必须包含 `3` 条 `Product` 与 `2` 条 `Repository`
- `Product` 必须覆盖 `active` 与 `archived` 两种状态
- `Repository` 必须覆盖 `active` 与 `archived` 两种状态
- `Repository.provider` 必须至少覆盖两个不同字符串值
- 所有基线记录都必须通过业务字段查找，不得硬编码 `UUID`
- 必须保留 `phase02` 历史基线中已存在的既有记录名：`Product A / Product B / Product C / main-repo / mirror-repo`
- `phase04` 新增的基线 `Product / Repository` 记录只能在这些既有 name 之外扩展，不得替换或重命名既有记录
- 必须同时覆盖三类关系：`product_modules`、`product_repositories`、`module_repositories`
- 三类关系最小数量约束：至少 `2` 条 `product_modules`、至少 `1` 条 `product_repositories`、至少 `2` 条 `module_repositories`
- 必须保留至少 `1` 条"无已绑定仓库的 `Product`"与至少 `1` 条"无已绑定 `Product` 的 `Repository`"，以验证空区块语义
- 必须更新 `database/seeds/seed_readonly_prereqs.sql` 以兼容 `0006` NOT NULL 约束

### 13.4 冷启动验收路径

> 环境重置步骤（`reset_module_mainline.sh` → `reset_product_repository_mainline.sh --clean-only`）由 §13.1 联调环境可重复建立承接；空状态入口验证由 §8.1 空状态引导承接。本节聚焦从空状态到完成首轮绑定的最小验收主路径。

- 从空状态进入 `Product Create`，完成首个 `Product` 创建
- 从空状态进入 `Repository Create`，完成首个 `Repository` 创建
- 完成首个 `Product -> Repository` 绑定
- 将已存在 `Module` 绑定到 `Product`
- 将已存在 `Module` 映射到 `Repository`
- 从 `Module Detail` 的兼容入口跳转进入正式绑定主入口，并完成至少一条绑定写入
- 三类绑定成功后必须回到对应 canonical owner 页面完成 reread

### 13.5 异常路径验证

- `CreateProduct` / `CreateRepository` 必填字段缺失 → 校验失败
- `CreateProduct` / `CreateRepository` 非法 `status` → 校验失败
- `BindModuleToProduct` / `BindRepositoryToProduct` / `MapModuleToRepository` 目标不存在 → 资源不存在
- `BindModuleToProduct` / `BindRepositoryToProduct` / `MapModuleToRepository` 目标实体非 `active` 状态 → 校验失败
- 三类绑定重复 → 重复冲突
- 候选读取返回空结果 → 空列表语义（非错误）
- 列表读取返回空结果 → 空列表语义（非错误）
- 详情读取不存在资源 → 资源不存在

### 13.6 多入口回流验收矩阵

- `fromList` 返回列表并携带原 `queryText / statusFilter`
- `fromModuleDetail` 返回原 `ModuleDetailPage`
- `fromProductDetail` 返回原 `ProductDetailPage`
- `direct-entry` 返回对应 List 默认筛选参数
- `Product Create` 成功回流到 `Product Detail`，携带继承来源标记
- `Repository Create` 成功回流到 `Repository Binding Detail`，携带继承来源标记
- 三类绑定写入成功后停留 canonical owner 页面 + 重新读取对应已绑定列表 + 面板回到 `closed`

### 13.7 验收约束

- 当前阶段验收不得依赖临时手工 SQL 才能建立最小联调环境
- 必须提供可重复执行的重置脚本、基线种子与异常路径验证前提
- `--restore-only` 模式必须稳定恢复到规格声明的正式基线，包括处理同标题 readonly seed 的情况

---

## 14. 非目标矩阵

当前阶段明确不做：

- `Feature / Opportunity / Experiment` 主线
- Dashboard 聚合反馈与 `product_asset_coverage` 完整消费链
- `Decision -> Product / Repository` 正式关联写入主线
- GitHub OAuth / 自动导入
- 自动扫描代码
- 自动知识图谱
- AI 自动建议或自动绑定
- 独立 `AI Assistant` 工作台
- 独立 `React Native` 客户端
- 完整 `PWA` 能力落地
- 复杂导入向导或 AI 建议作为首轮创建前提
- 超出当前阶段的分页、复杂检索或批量操作接口
- 完整跨模块服务矩阵（若后续需要，应在对应新 `phase` 或 `audit` 中单独冻结）
- 第二套移动端 UI 架构
- 完整自动化测试平台建设
- `Module Registry` 的重构式返工
- `Decision Center` 的返工重做

---

## 15. Done 标准

当以下条件成立时，`Product / Repository / Binding` 当前阶段规格成立并可进入 `phase04-11 / 12 / 13`：

1. 页面矩阵与动作矩阵已单值化且一一对应，五个核心动作的页面 owner / canonical write owner / reread owner 三链贯通
2. 数据模型与读写模型已明确，`Product / Repository` 最小模板、`status` 枚举、三类绑定关系语义与候选读取基线已冻结
3. 接口分组（读组 + 写组 + 候选读取 + 跨模块关系摘要读取）已明确，跨模块读边界双轨已收口
4. 空状态、冷启动路径与多入口回流矩阵已明确
5. 前端实现设计层结果（页面文件落点、路由树、组件树、状态模型、布局降级）已冻结到可直接进入实现的深度
6. 后端实现设计层结果（模块边界、文件落点、分层语义、候选读取临时承接、跨模块关系摘要读取链路）已冻结到可直接进入实现的深度
7. `.proto` 合同主线已冻结为唯一合同源，包名、服务矩阵、消息结构、字段编号、RPC → HTTP 映射、合同演进与 Buf 校验链已明确
8. 迁移边界与兼容策略已明确，`0006` migration、重置脚本、基线 seed、历史数据兼容回填与旧 transport 兼容委派已冻结
9. 错误语义已单值化，创建校验失败、资源不存在、重复冲突、候选空、列表空与详情不存在已收口
10. 验收基线已明确，联调环境可重复建立，冷启动路径、异常路径与多入口回流验收矩阵已完整
11. 非目标矩阵已明确
12. 本文档与 `mvp_spec_v0.1.md`、`module_registry_spec_v0.1.md`、`phase02-12` 验收结论、`decision_center_spec_v0.1.md`、`phase03-14` 验收结论、根级真相源及 `phase04` 三件套互链一致，无第二套边界

---

## 16. 与根级真相源互链

- 项目定位与入口：以 `AGENTS.md` 为准
- 阶段路线：以 `plan.md` 为准
- 技术栈标准：以 `TECH_STACK_BASELINE.md` 为准
- 规则门禁：以 `project_rules.md` 为准
- 目录与迁移落点：以 `architecture_map.md` 为准
- 最终共识：以 `PSCO-summarize-feedback.md` 为准
- 当前阶段唯一执行层上游：`mvp_spec_v0.1.md`
- `Module Registry` 已交付边界：`module_registry_spec_v0.1.md` + `phase02-12` 验收结论
- `Decision Center` 已交付边界：`decision_center_spec_v0.1.md` + `phase03-14` 验收结论
- 当前阶段三件套：`phase04_product_and_repository_binding_foundation_architecture_plan.md` / `phase04_product_and_repository_binding_foundation_dev_plan.md` / `phase04_product_and_repository_binding_foundation_shared_baseline.md`
- 当前阶段子规格追踪来源：`phase04-01` 到 `phase04-09`（仅作为本文档冻结来源与追踪依据，不再作为长期并列入口）
