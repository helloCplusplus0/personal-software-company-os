# PSCO Module Registry 规格 v0.1 — 正式规格正文

> **文档定位**：本文档是 `phase02_module_registry_foundation` 的正式规格正文，作为后续 `phase02-10 / 11 / 12 / 13` 实现与验收以及 `phase03` 引用 `Module Registry` 主线时的直接上游规格来源。
> **上游收敛**：本文档由 `phase02-01` 到 `phase02-08` 的冻结结论收敛而成，不另立第二套边界。
> **互链前提**：本文档以 `phase01-06` 的 `mvp_spec_v0.1.md` 为当前阶段唯一执行层上游，并与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致。

---

## 1. 技术路线

本文档继承 `mvp_spec_v0.1.md` §1 的技术路线，聚焦 `Module Registry` 当前阶段：

- 项目路线：`Durable System Track`
- 正式运行主线：`React Web + Go Backend + PostgreSQL`
- 前端：`React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`
- 前端交付策略：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI；不引入独立 `React Native` 客户端，`PWA` 仅作可兼容增强方向
- 后端：`Go`，模块化单体优先
- 数据库：`PostgreSQL` 为唯一数据库主线
- 合同：`Contract First`，跨语言长期标准为 `Protocol Buffers`
- 部署：`Caddy + systemd`，运行方式 `Single Server First`

> 约束：不得重新解释为 `Product Track`；不得把 `Rust` 写成当前阶段必需项；`Local First` 当前解释为数据所有权优先，不等于切换到 `SQLite`。

---

## 2. 对象范围

### 2.1 核心对象（`Module Registry` 直接承接）

- `Module`：模块登记与版本主线的核心实体
- `Release`：依附 `Module` 的版本登记对象

### 2.2 连接对象（只读或候选读取，不承接写入主线）

- `Product`：作为 `BindModuleToProduct` 的候选读取前提
- `Repository`：作为 `MapModuleToRepository` 的候选读取前提
- `Decision`：作为 `ModuleDetailRead` 的附属只读展示或跳转入口

### 2.3 非归属能力

- `CreateProduct`、`CreateRepository`、`RecordDecision` 不归属于当前阶段 `Module Registry` 后端模块
- `Module Registry` 只能通过最小连接边界读取或跳转这些能力

---

## 3. 页面矩阵

### 3.1 页面主线

| 页面 | 职责 |
| --- | --- |
| `Module Registry / List` | 承接列表读取、筛选入口、创建入口与进入详情入口 |
| `Module Create` | 承接 `CreateModule` |
| `Module Detail` | 承接详情读取、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository` 与 `Decision` 只读入口 |
| `Release Create` | 承接 `CreateRelease`，依附当前模块上下文 |

### 3.2 轻量跳转或关联入口

- `Product Registry`
- `Repository Binding`
- `Decision Center`

> 约束：上述轻量入口不得扩写为当前阶段第二条页面主线。

### 3.3 页面跳转关系

- `Module Registry / List` → `Module Create`
- `Module Registry / List` → `Module Detail`
- `Module Detail` → `Release Create`
- `Module Detail` → `Product Registry / Repository Binding / Decision Center`（轻量跳转或关联入口）

---

## 4. 动作矩阵

### 4.1 直接承接动作

| 动作 | 页面拥有者 | 写入数据载体 |
| --- | --- | --- |
| `CreateModule` | `Module Create` | `modules` |
| `CreateRelease` | `Module Detail` → `Release Create` | `module_releases` |
| `BindModuleToProduct` | `Module Detail` | `product_modules` |
| `MapModuleToRepository` | `Module Detail` | `module_repositories` |

### 4.2 允许最小入口但不扩写为独立主线

| 动作 | 承接方式 |
| --- | --- |
| `LinkDecisionToTarget` | 只作为只读展示或跳转入口，不扩写为当前阶段独立写入主线 |

### 4.3 非归属动作

- `CreateProduct`、`CreateRepository`、`RecordDecision` 不在当前阶段 `Module Registry` 后端模块中实现

---

## 5. 数据模型

### 5.1 直接承接表

- `modules`
- `module_releases`
- `product_modules`
- `module_repositories`

### 5.2 最小读取或关联前提

- `decisions`
- `decision_links`

### 5.3 候选读取前提（只读，不要求写入主线）

- `products`
- `repositories`

> 约束：候选读取前提表的结构定义属于实现阶段（`phase02-11` 数据主线）需要涵盖的只读前提，其运行时所需最小候选记录可通过最小种子数据或测试 fixture 提供以支撑 `phase02-12` 验收，但不视为 `phase02` 写入主线。

### 5.4 对象最小字段

- **`Module`**：`id`、`name`、`description`、`status`、`created_at`
- **`Release`**：`id`、`module_id`、`version`、`status`、`released_at`

### 5.5 Module 状态表达

- 推荐最小状态集合：`active` / `archived`
- 不得在当前阶段额外引入新的复杂生命周期体系作为实现前提

### 5.6 模块准入规则

1. **已登记**：通过 `CreateModule` 完成手动登记，且 `name` 与 `description` 已定义
2. **名称唯一**：`name` 在系统内唯一
3. **状态明确**：`Module` 必须先进入明确的 `status`，未定型状态不得进入版本主线
4. **版本准入**：存在至少一个可追踪 `Release` 后，才进入版本主线
5. **不因未绑定而阻断**：未绑定 `Product` 或未映射 `Repository` 时允许登记，但不得进入有效复用反馈统计

### 5.7 最小读写模型

- **列表读取**至少承接：`name / description / status / latest_release / product_bind_count / repository_bind_count`
- **详情读取**至少承接：核心对象字段、版本列表、产品绑定、仓库映射与相关 `Decision` 入口
- **创建写入**承接：`CreateModule`（最小字段 `name / description / status`）
- **版本写入**承接：`CreateRelease`（最小字段 `version / status / released_at`，`module_id` 由上下文隐式承接）
- **关联写入**承接：`BindModuleToProduct`、`MapModuleToRepository`

---

## 6. API 边界与接口分组

### 6.1 合同边界

- 遵守 `Contract First`，以结构化合同优先
- 长期跨语言合同方向冻结为 `Protocol Buffers`
- 当前阶段不要求完整 `proto` 工具链落地
- 不得引入与 `Protocol Buffers` 冲突的第二套跨语言合同主线

### 6.2 接口分组

#### 读组

| 接口 | 职责 |
| --- | --- |
| `ModuleListRead` | 只承接列表展示所需最小读取、筛选前提与进入详情所需标识 |
| `ModuleDetailRead` | 统一承接模块核心字段、版本列表、产品绑定、仓库映射与相关 `Decision` 入口读取 |

#### 写组

| 接口 | 职责 |
| --- | --- |
| `ModuleCreateWrite` | 只承接 `CreateModule`，返回新建模块标识 |
| `ModuleReleaseWrite` | 以 `moduleId` 为上下文前提，只承接当前模块下的版本登记 |
| `ModuleBindingWrite` | 统一承接 `BindModuleToProduct` 与 `MapModuleToRepository` |

#### 候选读取（phase02 由 Module Registry 临时承接）

| 接口 | 职责 |
| --- | --- |
| `ProductBindingCandidateRead` | 只服务 `BindModuleToProduct` 的候选 Product 选择 |
| `RepositoryBindingCandidateRead` | 只服务 `MapModuleToRepository` 的候选 Repository 选择 |

#### 方向级最小 API 矩阵

| 动作语义 | 接口分组 | 接口名 | 方向 |
| --- | --- | --- | --- |
| 列表读取 | 读组 | `ModuleListRead` | 读 |
| 详情读取 | 读组 | `ModuleDetailRead` | 读 |
| Product 候选读取 | 候选读取 | `ProductBindingCandidateRead` | 读 |
| Repository 候选读取 | 候选读取 | `RepositoryBindingCandidateRead` | 读 |
| `CreateModule` | 写组 | `ModuleCreateWrite` | 写 |
| `CreateRelease` | 写组 | `ModuleReleaseWrite` | 写 |
| `BindModuleToProduct` | 写组 | `ModuleBindingWrite` | 写 |
| `MapModuleToRepository` | 写组 | `ModuleBindingWrite` | 写 |
| `Decision` 附属读取 | 内嵌于 `ModuleDetailRead` | — | 读 |

> 本矩阵与 `mvp_spec_v0.1.md` §6.2 最小 API 矩阵互链，仅承接 `Module Registry` 当前阶段涉及的语义，不扩写为完整跨模块服务矩阵。

### 6.3 Decision 读取边界

- `Decision` 在当前阶段作为 `ModuleDetailRead` 的附属读取承接，**不设独立读接口组**
- 不得为 `Decision` 单独设立独立 Read 接口、独立 handler 或独立 service 文件
- `phase03` 实现 `Decision` 模块后，`Decision` 数据组装逻辑可迁移为通过 `Decision` 模块读模型协作获取，但 `ModuleDetailRead` 的外部接口契约保持不变

### 6.4 关系写入后的读取语义

- 任一关联写入成功后，默认后端语义为"详情页重新读取当前绑定结果"
- 不得额外设计一套脱离 `ModuleDetailRead` 的回流读取路径

### 6.5 合同与存储解耦

- 不得直接将 `modules`、`module_releases`、`product_modules`、`module_repositories` 的存储模型原样暴露为外部合同
- 必须允许后续在 `Contract First` 路线下独立演进接口消息结构
- 不得复用已删除字段编号或字段语义，必须保持与 `Protocol Buffers` 长期方向兼容

---

## 7. 空状态与冷启动

### 7.1 空状态引导

- 用户首次进入 `Module Registry / List` 且系统中尚无任何模块时，页面必须展示明确的空状态提示
- 空状态主动作必须直接进入 `Module Create`
- 空状态文案必须围绕"先完成首个模块登记"展开
- 不得把导入、自动扫描或 AI 建议写成空状态主入口

### 7.2 冷启动路径

- 首轮必须允许用户从空状态进入 `CreateModule`
- 首轮必须允许用户完成首个 `Release` 登记
- 首轮允许模块先登记后补充 `Product / Repository` 关联
- 当前阶段不依赖导入、自动扫描或 AI 推荐

### 7.3 创建最小闭环

- 最小闭环为 `Module Registry / List → Module Create → Module Detail`
- `CreateModule` 成功后默认回流到 `ModuleDetailPage`
- `CreateRelease` 成功后默认回流到当前模块的 `ModuleDetailPage`
- 绑定动作成功后用户必须停留在当前 `ModuleDetailPage`，并重新读取对应绑定结果

### 7.4 返回路径

- 从 `Module Create` 主动取消或返回：默认返回保留原搜索参数上下文的 `ModuleListPage`
- 从 `Module Detail` 主动返回 `Module List`：必须恢复原有的 `queryText` 与 `statusFilter`，不得要求用户重新输入筛选条件
- 从 `Release Create` 主动取消或返回：默认返回当前模块的 `ModuleDetailPage`

---

## 8. 前端实现设计层结果

### 8.1 页面文件落点

| 页面 | 文件路径 |
| --- | --- |
| `ModuleListPage` | `frontend/src/features/module-registry/pages/module-list-page.tsx` |
| `ModuleCreatePage` | `frontend/src/features/module-registry/pages/module-create-page.tsx` |
| `ModuleDetailPage` | `frontend/src/features/module-registry/pages/module-detail-page.tsx` |
| `ReleaseCreatePage` | `frontend/src/features/module-registry/pages/release-create-page.tsx` |

### 8.2 路由树与 URL 语义

| 路由 | URL | 文件路径 |
| --- | --- | --- |
| `ModuleListRoute` | `/modules` | `frontend/src/routes/modules/index.tsx` |
| `ModuleCreateRoute` | `/modules/new` | `frontend/src/routes/modules/new.tsx` |
| `ModuleDetailRoute` | `/modules/:moduleId` | `frontend/src/routes/modules/$moduleId.tsx` |
| `ReleaseCreateRoute` | `/modules/:moduleId/releases/new` | `frontend/src/routes/modules/$moduleId/releases/new.tsx` |

> 约束：不得把 `Product`、`Repository`、`Decision` 提前扩写为并列主树；当前阶段不得把 `Module Detail` 提前拆成独立的子路由工作台。

### 8.3 组件树

#### 列表页

- `ModuleListPageShell`
  - `ModuleListToolbar`（承接筛选入口与创建入口）
  - `ModuleListContent`（只承接模块列表读取结果）
  - `ModuleListTableOrCards`

#### 创建页

- `ModuleCreatePageShell`
  - `ModuleCreateForm`（只承接 `name / description / status` 最小录入）
  - `ModuleCreateActions`

#### 详情页

- `ModuleDetailPageShell`
  - `ModuleSummaryCard`（只承接模块核心字段与状态表达）
  - `ModuleReleaseListSection`
  - `ModuleBindingPanel`（直接承接 `BindModuleToProduct` 与 `MapModuleToRepository`）
  - `ModuleDecisionEntryPanel`（只承接只读展示或跳转）

#### 版本登记页

- `ReleaseCreatePageShell`
  - `ReleaseCreateForm`（只承接当前模块上下文中的版本登记最小字段）
  - `ReleaseCreateActions`

#### 组件归属原则

- `ModuleSummaryCard`、`ModuleReleaseListSection`、`ModuleBindingPanel`、`ModuleDecisionEntryPanel` 默认先归属于 `ModuleDetailPage`
- `ModuleCreateForm` 默认先归属于 `ModuleCreatePage`
- `ReleaseCreateForm` 默认先归属于 `ReleaseCreatePage`
- 只有在列表页与详情页或多个页面确实共享同一职责时，才允许抽为共享组件
- 当前阶段不得为了“组件纯洁”提前拆出无明确复用证据的通用组件层

### 8.4 状态模型

#### Module List

- 查询条件冻结到路由搜索参数层：`queryText`、`statusFilter`
- 从 `ModuleCreatePage` 或 `ModuleDetailPage` 返回时，必须按原有 `queryText` 与 `statusFilter` 恢复列表上下文
- 读取状态：`pending`、`success`、`error`
- 派生视图状态：`initial-loading`、`ready`、`empty`、`error`
- 错误反馈停留在 `ModuleListPage` 内容区域，不跳转独立错误页

#### Module Create

- 草稿状态：`idle`、`dirty`（至少承接 `name / description / status`）
- 提交状态：`submitting`、`submit-success`、`submit-error`
- 提交失败时停留当前页，保留草稿，错误显示在表单上下文
- 提交成功默认回流到 `ModuleDetailPage`
- 提交成功后用户返回 `ModuleListPage` 时，列表读取必须能够承接新建模块结果

#### Release Create

- 当前模块标识来自路由参数 `moduleId`，不得复制可写全局状态
- 状态：`idle`、`dirty`、`submitting`、`submit-success`、`submit-error`
- 提交成功默认回流到当前模块的 `ModuleDetailPage`
- 提交成功回流后 `ModuleDetailPage` 必须承接最新版本列表读取

#### 绑定动作

- 面板状态：`closed`、`open-idle`、`submitting`、`submit-success`、`submit-error`
- 同一时刻只允许一个绑定面板处于打开态
- 绑定成功后停留在 `ModuleDetailPage`，重新读取绑定结果
- 绑定失败时错误停留在面板上下文，保留当前选择

### 8.5 布局降级策略

#### PC 页面布局

- `ModuleListPage`：高信息密度列表布局
- `ModuleDetailPage`：分区式详情布局，摘要、版本与关联入口可同时可见
  - 至少分为摘要主区、版本列表区、关联动作区、`Decision` 入口区

#### 移动浏览器布局

- `ModuleListPage`：单列列表或卡片重排
- `ModuleDetailPage`：摘要、版本、关联、`Decision` 入口按垂直顺序重排，次级信息可折叠
- `ModuleCreatePage` / `ReleaseCreatePage`：单列垂直布局，主动作按钮无需横向滚动即可见

> 约束：不得新增独立移动端页面体系；不得引入独立 `React Native` 客户端；不得把完整 `PWA` 能力写成当前阶段实现前提。

### 8.6 运行时实现细节不冻结

- 当前阶段不提前冻结具体 hook 命名、Query key 细节、store API 命名、缓存时间、请求取消策略或 optimistic update 方案
- 页面级 UI 状态（草稿、面板开闭、提交错误等瞬时状态）优先归属于当前页面或当前详情页上下文，只在确有跨同页多组件共享需要时再抽为页面作用域状态容器，不得默认升级为跨路由全局状态

---

## 9. 后端实现设计层结果

### 9.1 模块边界

- `Module Registry` 在后端冻结为单一业务模块
- 统一承接 `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository` 及列表/详情读取
- 不得把这些能力拆散到 `Product`、`Repository` 或 `Decision` 的后端模块中

### 9.2 文件落点

```
backend/internal/moduleregistry/
  handler/
    query_handler.go              # 读组入口：ModuleListRead + ModuleDetailRead
    command_handler.go            # 写组入口：ModuleCreateWrite + ModuleReleaseWrite + ModuleBindingWrite
  service/
    query_service.go              # 读组编排（含 Decision 内嵌附属读取）
    command_service.go            # 写组编排（创建 + 版本 + 绑定）
  repository/
    module_store.go               # modules 表读写
    release_store.go              # module_releases 表读写
    binding_store.go              # product_modules + module_repositories 表读写
  candidate/
    product_candidate_read.go     # ProductBindingCandidateRead（临时承接，phase03 可迁移）
    repository_candidate_read.go  # RepositoryBindingCandidateRead（临时承接，phase04 可迁移）
```

### 9.3 分层语义

| 层 | 子包 | 职责 |
| --- | --- | --- |
| 入口层 | `handler/` | 只负责承接请求与返回结果 |
| 业务编排层 | `service/` | 只负责动作语义、校验顺序与跨连接口编排 |
| 数据访问/外部连接层 | `repository/` + `candidate/` | 只负责持久化与依赖调用 |

### 9.4 读组与写组代码组织

- 读组（`ModuleListRead` + `ModuleDetailRead`）→ `handler/query_handler.go` + `service/query_service.go`（单文件编排，不拆 `list_service.go` / `detail_service.go`）
- 写组（`ModuleCreateWrite` + `ModuleReleaseWrite` + `ModuleBindingWrite`）→ `handler/command_handler.go` + `service/command_service.go`（单文件编排，不拆三个独立 service 文件）

### 9.5 Decision 读取实现

- `Decision` 读取不设独立文件落点
- `Decision` 关联数据的读取逻辑内嵌在 `service/query_service.go` 的详情读取编排中
- 不得在 `candidate/` 子包中新增 `decision_*` 文件

### 9.6 跨模块候选读取临时承接

- `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead` 在 `phase02` 阶段由 `Module Registry` 后端模块临时承接
- 必须通过独立 Read 接口定义与独立代码落点（`candidate/` 子包）隔离，不得在 `service/` 中直接写跨模块 SQL
- `phase03 / phase04` 实现对应模块后，这两个文件可整体迁移到对应模块，但接口契约保持不变

### 9.7 不冻结的实现工具

- 不得提前冻结 Go HTTP 框架、RPC 框架、ORM、SQL Builder、Repository 模板或目录生成器
- 文件内部的函数签名、中间件选型、数据库访问方式不作为当前阶段冻结内容
- 只冻结职责分工、接口归属与文件落点，不冻结实现手段

---

## 10. 非目标矩阵

当前阶段明确不做：

- `Product` 全量主线
- `Decision Center` 全量主线
- `Repository Binding` 全量主线
- `Dashboard` 聚合反馈
- 自动扫描代码
- 自动知识图谱
- 独立 `AI Assistant` 工作台
- 独立 `React Native` 客户端
- 完整 `PWA` 能力落地
- 复杂导入向导或 AI 建议作为首轮创建前提
- 完整跨模块服务矩阵（若后续需要，应在对应新 phase 或 audit 中单独冻结）
- 第二套移动端 UI 架构

---

## 11. Done 标准

当以下条件成立时，`Module Registry` 当前阶段规格成立并可进入 `phase02-10 / 11 / 12`：

1. 页面矩阵与动作矩阵已单值化且一一对应
2. 数据模型与读写模型已明确
3. 接口分组（读组 + 写组 + 候选读取）已明确
4. 空状态与冷启动路径已明确
5. 前端实现设计层结果（页面文件落点、路由树、组件树、状态模型、布局降级）已冻结到可直接进入实现的深度
6. 后端实现设计层结果（模块边界、文件落点、分层语义、候选读取临时承接）已冻结到可直接进入实现的深度
7. `Decision` 边界已收口为 `ModuleDetailRead` 内嵌附属读取，不设独立读接口组
8. 非目标矩阵已明确
9. 本文档与 `mvp_spec_v0.1.md` 及根级真相源互链一致，无第二套边界

---

## 12. 与根级真相源互链

- 项目定位与入口：以 `AGENTS.md` 为准
- 阶段路线：以 `plan.md` 为准
- 技术栈标准：以 `TECH_STACK_BASELINE.md` 为准
- 规则门禁：以 `project_rules.md` 为准
- 目录与迁移落点：以 `architecture_map.md` 为准
- 最终共识：以 `PSCO-summarize-feedback.md` 为准
- 当前阶段唯一执行层上游：`mvp_spec_v0.1.md`
- 当前阶段架构规划：`phase02_module_registry_foundation_architecture_plan.md`
- 当前阶段共享基线：`phase02_module_registry_foundation_shared_baseline.md`
- 当前阶段开发计划：`phase02_module_registry_foundation_dev_plan.md`

> 本文档不重写上述根级真相源与阶段文档的主结论，仅作为 `Module Registry` 当前阶段的唯一规格收敛入口。前置子规格 `phase02-01 ~ 08` 仅作为本文档的冻结来源与追踪依据，不再作为后续实现与验收的长期并列入口。
