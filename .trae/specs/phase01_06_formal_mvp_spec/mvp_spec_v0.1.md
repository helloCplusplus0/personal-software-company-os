# PSCO MVP 规格 v0.1 — 正式 MVP 规格正文

> **文档定位**：本文档是 PSCO 项目 `v0.1` 的正式 MVP 规格正文，作为后续实现与 `phase02` 的唯一上游规格来源。
> **上游收敛**：本文档由 `phase01-01` 到 `phase01-05` 的冻结结论收敛而成，不另立第二套边界。
> **互链前提**：本文档与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致。

---

## 1. 技术路线

- 项目路线：`Durable System Track`
- 正式运行主线：`React Web + Go Backend + PostgreSQL`
- 前端：`React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`
- 前端交付策略：`v0.1` 只交付单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI；当前不引入独立 `React Native` 客户端，`PWA` 仅作为可兼容增强方向
- 后端：`Go`，模块化单体优先，单进程、单主运行面优先
- 数据库：`PostgreSQL` 为 `v0.1` 唯一数据库主线
- 可选计算：`Rust` 仅作计算扩展位，不进入 `v0.1` 首轮实现
- 合同：`Contract First`，跨语言长期标准为 `Protocol Buffers`
- 部署：`Caddy + systemd`，运行方式 `Single Server First`

> 约束：本项目不得重新解释为 `Product Track`；不得把 `Hono` 写成主运行时；不得把 `Rust` 写成 `v0.1` 当前必需项；`Local First` 当前解释为数据所有权优先，不等于切换到 `SQLite`。

---

## 2. 对象范围

### 2.1 核心实体（进入 `v0.1` 主执行范围）

- `Product`
- `Module`
- `Release`
- `Decision`
- `Repository`
- `Venture`（可选实体，不强制创建）

### 2.2 派生层

- `Capability`：作为派生结果层，不单独建核心表，不要求用户主动 CRUD

### 2.3 后移对象（不进入 `v0.1` 主执行范围）

- `Feature`
- `Opportunity`
- `Experiment`

---

## 3. 动作矩阵

`v0.1` 必须承接以下核心动作：

- `CreateProduct`
- `CreateModule`
- `CreateRelease`
- `CreateRepository`
- `RecordDecision`
- `BindRepositoryToProduct`
- `BindModuleToProduct`
- `MapModuleToRepository`
- `LinkDecisionToTarget`

可选补充动作：

- `CreateVenture`
- `BindProductToVenture`

> 动作-数据载体对应：`CreateProduct/CreateModule/CreateRelease/CreateRepository` → 对应核心表；`RecordDecision` → `decisions`；`BindRepositoryToProduct` → `product_repositories`；`BindModuleToProduct` → `product_modules`；`MapModuleToRepository` → `module_repositories`；`LinkDecisionToTarget` → `decision_links`。

---

## 4. 页面与输入路径

### 4.1 页面范围

- `Dashboard`：最小聚合反馈页面，不做复杂驾驶舱
- `Module Registry`：承载模块登记与 `CreateRelease` 等动作
- `Product Registry`：承载产品登记与绑定入口
- `Decision Center`：承载决策记录与目标关联
- `Repository Binding`：承载仓库绑定与实现映射

### 4.2 明确不进入 `v0.1`

- `AI Assistant` 一级导航
- `Feature` / `Opportunity` / `Experiment` 页面

### 4.3 输入路径

- 首轮默认以低摩擦手动录入为主
- 空状态服务于首轮录入路径
- 独立 `AI Assistant` 工作台不纳入 `v0.1`
- 前端交付维持单一 `React Web` 入口，不拆分独立原生客户端；页面交互同时考虑 `PC` 与移动浏览器可用性

---

## 5. 数据模型

### 5.1 核心表

- `ventures`（可选）
- `products`
- `repositories`
- `modules`
- `module_releases`
- `decisions`

### 5.2 关系表

- `product_modules`
- `product_repositories`
- `module_repositories`
- `decision_links`

### 5.3 派生视图

- `capability_summary`
- `module_reuse_summary`
- `product_asset_coverage`
- `pending_decision_signals`

### 5.4 对象最小字段

以下为 `v0.1` 各核心对象的最小字段集，用于登记与检索闭环，不包含后续可扩展字段：

- **`Product`**：`id`、`name`、`description`、`status`、`created_at`
- **`Module`**：`id`、`name`、`description`、`status`、`created_at`
- **`Release`**：`id`、`module_id`、`version`、`status`、`released_at`
- **`Repository`**：`id`、`name`、`url`、`provider`、`status`、`created_at`
- **`Venture`（可选）**：`id`、`name`、`description`、`status`、`created_at`
- **`Decision`**：`id`、`title`、`context`、`problem`、`alternatives`、`choice`、`reason`、`impact`、`status`、`created_at`

> 说明：派生视图与 `Capability` 不参与对象最小字段定义；对象最小字段只服务于核心对象的手动登记、检索与基础反馈闭环。

### 5.5 模块准入规则

一个 `Module` 进入注册与版本主线，必须满足以下准入条件：

1. **已登记**：通过 `CreateModule` 完成手动登记，且 `name` 与 `description` 已定义，未定义的 `Module` 不得进入注册主线
2. **名称唯一**：`name` 在系统内唯一，避免重复登记同一模块
3. **状态明确**：`Module` 必须先进入明确的 `status`（如 `active` / `archived`），未定型状态不得进入版本主线
4. **版本准入**：`Module` 只有在存在至少一个 `Release`（通过 `CreateRelease` 创建）且 `Release` 状态可追踪时，才进入版本主线
5. **不因未绑定而阻断**：`Module` 未绑定 `Product` 或未映射 `Repository` 时允许登记，但不得进入可复用反馈（`capability_summary` / `module_reuse_summary`）的有效统计

> 说明：准入规则只约束"何时进入注册与版本主线"，不扩大 `v0.1` 范围；未达标模块仅保留登记，不参与复用反馈统计。

### 5.6 结构要求

- **Repository Binding**：冻结为产品绑定、模块绑定与实现映射三类关系，采用显式关系承载；模块与仓库的实现映射通过 `module_repositories` 显式承载
- **Decision Record**：冻结为结构化记录，最小字段包含 `title / context / problem / alternatives / choice / reason / impact / status`，并通过 `decision_links` 与目标对象结构化关联

---

## 6. API 边界

### 6.1 合同边界

- 遵守 `Contract First`，以结构化合同优先，不由前端猜字段或由实现倒推合同
- 长期跨语言合同方向冻结为 `Protocol Buffers`
- 当前阶段不要求完整 `proto` 工具链落地
- 不得引入与 `Protocol Buffers` 冲突的第二套跨语言合同主线（如并列的 `OpenAPI` / `GraphQL` / 自定义 JSON 协议）

### 6.2 最小 API 矩阵

将 `v0.1` 核心动作映射到最小接口集合。以下为方向级映射，具体端点命名与入参出参以最终 `Contract First` 合同为准，不作为成品协议：

| 核心动作 | 最小 API |
| --- | --- |
| `CreateProduct` | `POST /products` |
| `CreateModule` | `POST /modules` |
| `CreateRelease` | `POST /modules/{module_id}/releases` |
| `CreateRepository` | `POST /repositories` |
| `RecordDecision` | `POST /decisions` |
| `BindRepositoryToProduct` | `POST /products/{product_id}/repositories` |
| `BindModuleToProduct` | `POST /products/{product_id}/modules` |
| `MapModuleToRepository` | `POST /modules/{module_id}/repositories` |
| `LinkDecisionToTarget` | `POST /decisions/{decision_id}/links` |

可选补充动作映射：

| 可选动作 | 最小 API |
| --- | --- |
| `CreateVenture` | `POST /ventures` |
| `BindProductToVenture` | `POST /ventures/{venture_id}/products` |

> 说明：`v0.1` 最小接口集合以读（查询）与写（登记/绑定）为主，不包含导出/备份之外的重型接口；导出与备份以独立基础路径承接，不进入核心 CRUD 接口集合。

---

## 7. 冷启动、导入与导出

### 7.1 冷启动

- 首轮必须允许用户从零手动创建 `Product`、`Module`、`Repository`、记录 `Decision`
- 首轮必须允许用户手动完成基础绑定关系
- 冷启动不依赖外部集成、前置配置或自动导入

### 7.2 导入

- `v0.1` 当前不冻结任何正式导入能力，首轮资产进入系统统一以手动录入为主
- `GitHub OAuth / 自动导入` 不进入 `v0.1` 首轮
- 后续 `/spec` 只能在不破坏手动录入闭环的前提下，补充"非阻断型轻量导入说明"

### 7.3 导出

- 语义：**面向用户带走核心资产数据**，供用户独立保存、迁移或另作它用
- 范围至少覆盖 `Product / Module / Release / Repository / Decision` 及基础绑定关系
- 不得把导出留成"后续再说"

### 7.4 备份

- 语义：**面向当前实例保留与恢复**，供用户在更换设备或重新部署时恢复数据
- 至少提供一种面向当前实例的基础备份路径
- 不要求自动连续备份、多端同步或复杂灾备体系
- 导出与备份不得依赖 GitHub 或第三方平台作为唯一前提

---

## 8. 非目标

`v0.1` 明确不做：

- 微服务
- Kubernetes
- Docker 全流程
- GraphQL
- Kafka
- Redis 缓存层
- Elasticsearch
- GitHub OAuth 自动导入
- Rust 计算引擎接入
- 自动扫描代码
- 自动知识图谱
- AI 自动判断方案
- 完整 PMM / PCP 正式标准
- 独立 AI Assistant 工作台
- 独立 `React Native` 客户端
- 完整 `PWA` 能力落地

---

## 9. Done 标准

当以下条件成立时，`v0.1` 规格成立并可进入稳定实现：

1. 核心实体范围明确
2. 冷启动路径明确
3. 模块准入规则明确
4. `Decision` 模板明确
5. 页面与动作一一对应
6. 表结构与关系表明确
7. 非目标列表明确
8. 导出 / 备份要求明确
9. 至少一个真实项目 dry-run 走通

**建议的 MVP 先导指标**：

- 模块复用率（被 2+ Product 使用的 Module 占比）
- 决策复用率（新建相似决策时引用历史 Decision 的比例）
- 资产导入耗时（从创建 Product 到完成首轮 Module 绑定的时间）
- 录入摩擦感知（每个核心对象完成最小登记所需步骤数 / 字段数）

---

## 9.5 遗留与非阻断提醒

以下两项为已评估但有意后移的非阻断项，不在 `phase01-06` 范围内实现，作为后续承接点保留：

1. **空状态设计细化**：当前正文对空状态仅作轻量表达（§4.3）。进入实现前，应在页面级规格或 `phase02+` 中，把每个页面的空状态入口与首轮录入路径进一步落细。
2. **查询类与 `Dashboard` 聚合接口**：当前最小 API 矩阵以写操作（登记 / 绑定）为主。进入实现前，应展开查询类接口与 `Dashboard` 聚合读取接口，使其与 Done 标准中的反馈闭环衔接。

> 说明：以上为"后续在 `phase02+` 或页面级规格处理"的承接点，不构成 `v0.1` 当前范围，不允许写成当前既成事实。

---

## 10. 与根级真相源互链

- 项目定位与入口：以 `AGENTS.md` 为准
- 阶段路线：以 `plan.md` 为准
- 技术栈标准：以 `TECH_STACK_BASELINE.md` 为准
- 规则门禁：以 `project_rules.md` 为准
- 目录与迁移落点：以 `architecture_map.md` 为准
- 最终共识：以 `PSCO-summarize-feedback.md` 为准

> 本文档不重写上述根级真相源的主结论，仅作为 `v0.1` 执行层面的唯一规格收敛入口。
