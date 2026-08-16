# phase12-07 后端合同、导出结果与共享只读视图设计

> 本设计承接 `phase12-05` 已冻结的共享 owner、候选清单与判断题，负责回答：哪些继续复用 `GetProjectContext` 真实合同，哪些只停留在 L3 单向派生，哪些当前不做，以及是否存在足以进入 `ProjectContextService` 受控派生读取的候选。
> 所有结论以当前真实 `.proto`、Connect server、query service 与 renderer 实现为准，不把前端 adapter、页面 detail read 或文档占位误写成后端既有合同。

---

## 一、既有合同源现状

### 1.1 真实已有结构化合同

源自 `proto/psco/project_context/v1/project_context.proto`：

| 合同对象 | 真实字段 |
|---|---|
| `RepositorySummary` | `id / name / provider / url / description` |
| `ProductSummary` | `id / name / description / status` |
| `ModuleSummary` | `id / name / description / status` |
| `DecisionSummary` | `id / title / status / context / hit_sources / created_at` |
| `RuleEntry` | `key / label / summary / entry_ref / entry_kind` |
| `PhaseEntry` | `phase / label / status_summary / entry_ref / entry_kind` |
| `BoundaryEntry` | `key / label / summary` |
| `GetProjectContext` | 输入锚点仅 `repository_id` |
| `ExportProjectContext` | 仅返回 `markdown`，单向派生自 `GetProjectContext` 结构化结果 |

### 1.2 真实已有后端模块

| 文件 | 职责 | 分类 |
|---|---|---|
| `backend/internal/projectcontext/types.go` | project context 内部类型与 domain 映射 | follow-regression |
| `backend/internal/projectcontext/errors.go` | 错误语义 | follow-regression |
| `backend/internal/projectcontext/candidate/context_readers.go` | 聚合只读读取 | follow-regression |
| `backend/internal/projectcontext/service/query_service.go` | `GetProjectContext / ExportProjectContext` 查询编排 | follow-regression |
| `backend/internal/projectcontext/connect/server.go` | Connect server 承接位 | follow-regression |
| `backend/internal/projectcontext/renderer/markdown.go` | Markdown renderer | follow-regression |
| `backend/internal/gen/proto/psco/project_context/v1/*` | proto 生成代码 | no-change |
| `backend/internal/gen/connect/psco/project_context/v1/*` | Connect 生成代码 | no-change |

### 1.3 当前不存在的后端合同能力

以下内容在当前 `.proto` 与 `ProjectContextService` 中都**不存在**：

1. 跨实体统一的 `entity_type / entity_name / entity_status` 字段
2. `related_modules_count / related_products_count / decision_count / binding_status` 等聚合字段
3. `source_entity_type / source_entity_id` 级入口回溯字段
4. Product / Module / Decision 专用 `resolved_repository_id`
5. 除 `GetProjectContext / ExportProjectContext` 之外的 `ProjectContextService` 只读 RPC

---

## 二、影响对象清单

| # | 对象 | 分类 | 理由 |
|---|---|---|---|
| B1 | `proto/psco/project_context/v1/project_context.proto` | must-change | 本设计必须准确回答哪些继续复用真实字段、哪些不进入本阶段，不能再把不存在字段写成已覆盖 |
| B2 | `backend/internal/projectcontext/service/query_service.go` | follow-regression | 继续作为 `GetProjectContext / ExportProjectContext` 编排位，本设计需要明确是否新增受控派生读取 |
| B3 | `backend/internal/projectcontext/connect/server.go` | follow-regression | 继续作为 Connect 承接位，本设计需要明确当前不新增 RPC |
| B4 | `backend/internal/projectcontext/renderer/markdown.go` | follow-regression | 继续作为 Markdown renderer，本设计需要冻结其与 Web 共享视图的边界 |
| B5 | `backend/internal/projectcontext/candidate/context_readers.go` | follow-regression | 继续作为底层只读读取位，本设计需要明确当前不新增专用 resolver 读取 |
| B6 | `frontend/src/features/project-context/*` | follow-regression | L3 共享只读 owner 仍由 `phase12-06` 继续承接，本设计只冻结其可消费的后端事实边界 |
| B7 | `phase12-07 spec.md` | follow-regression | 需求与设计口径需继续一致 |
| B8 | `phase12-07 tasks.md` | must-change | 任务必须改为核验真实字段 / L3 派生 / 当前不做 / 候选评估四类结果 |
| B9 | `phase12-07 checklist.md` | must-change | 验收必须显式覆盖“不得假装已覆盖”“Decision 候选已完成正式评估”“真实文件路径准确” |

---

## 三、结论矩阵

### 3.1 共享摘要与定位需求分组

| # | 需求 | 当前真实来源 | 结论分组 | 理由 |
|---|---|---|---|---|
| S1 | Repository 最小摘要 | `RepositorySummary` | 继续复用 L1 真实字段 | 当前字段已足够被 Web / agent 同时消费 |
| S2 | Product 最小摘要 | `ProductSummary` | 继续复用 L1 真实字段 | 同上 |
| S3 | Module 最小摘要 | `ModuleSummary[]` | 继续复用 L1 真实字段 | 同上 |
| S4 | Decision 最小摘要 | `DecisionSummary[]` | 继续复用 L1 真实字段 | 当前应继续使用 `title`，不得虚构通用 `name` 字段 |
| S5 | 规则入口定位 | `RuleEntry.label / summary / entry_ref / entry_kind` | 继续复用 L1 真实字段 | 已满足当前入口定位要求 |
| S6 | phase 文档入口定位 | `PhaseEntry.label / status_summary / entry_ref / entry_kind` | 继续复用 L1 真实字段 | 已满足当前入口定位要求 |
| S7 | 四实体统一语义 label | L3 共享语义来源 | 仅停留在 L3 单向派生 | 这是共享解释来源，不是后端事实字段 |
| S8 | 跨实体展示 adapter（如 `entity_type / entity_primary_text / entity_status_text`） | L3 基于 L1 各实体字段单向派生 | 仅停留在 L3 / renderer 映射 | `Repository` 无 `status`，`Decision` 用 `title` 而非 `name`，不能伪造统一后端字段 |
| S9 | `entry_title / entry_summary` 统一 view model | L3 / renderer 对 `label / summary / status_summary` 的单向映射 | 仅停留在 L3 / renderer 映射 | 允许展示层统一命名，但不新增第二套合同字段 |
| S10 | `source_entity_type / source_entity_id` 级入口回溯 | 当前无字段 | 当前不做 | 当前没有 Web / agent 双侧稳定复用场景，不足以支撑新增合同 |
| S11 | 聚合计数字段（`module_count / decision_count / has_product` 等） | 当前仅可由数组长度或存在性局部推导 | 当前不做 | 若写入合同会复制 canonical facts；当前也没有稳定跨 Web / agent 共享痛点 |
| S12 | `binding_status` | 来自 feature detail read，而非 `ProjectContextService` | 当前不做 | 它属于切片专属读取结果，不应被误提升为 project-context 共享合同 |

### 3.2 Resolver 与受控派生读取候选分组

| # | 候选 | 当前状态 | 结论分组 | 理由 |
|---|---|---|---|---|
| R1 | 直接 `repository-scoped` 页面 | 路由参数已提供 `repository_id` | 继续复用 L1 真实锚点 | 不需要新增任何受控派生读取 |
| R2 | Product 页回到 `repository_id` | `use-product-detail-read` 已返回 `bound_repositories[*].repository_id` | 当前不做 | 当前只服务 Product detail 单页接入，不构成 Web / agent 双侧共享后端候选 |
| R3 | Module 页回到 `repository_id` | `use-module-detail-read` 已返回 `repository_mappings[*].repository_id` | 当前不做 | 同上 |
| R4 | Decision 页回到 `repository_id` | 仅能通过 `source_module_id / linked_modules -> Module detail -> repository_id` 链式解析 | 当前不做，但保留为最高优先级候选 | 当前仍缺少 Web / agent 双侧稳定消费收益，且新增 resolver 会主要优化单一前端页；但它必须被显式评估，不得笼统写成“既有 detail read 已覆盖” |
| R5 | Product / Module / Decision 稳定 `resolved_repository_id` 字段 | 当前无字段 | 当前不做 | 现阶段未证明新增字段比现有 detail read / 页面降级更能回收重复解释逻辑 |

### 3.3 受控派生读取候选正式判断

| 候选 | 是否进入 `ProjectContextService` 受控派生读取 | 当前判断 | 回收的重复解释逻辑 | 当前结论 |
|---|---|---|---|---|
| Product 稳定 `resolved_repository_id` | 否 | 页面仍可在 detail read 内按唯一候选规则处理 | 仅回收单页本地判断，未形成 Web / agent 共享痛点 | 当前不做 |
| Module 稳定 `resolved_repository_id` | 否 | 同 Product | 同上 | 当前不做 |
| Decision 稳定 `resolved_repository_id` / resolver | 否 | 已完成正式评估，但暂不进入后端 | 当前只会回收单页链式解析，不会同时回收 agent 导出或其他 3+ 页面稳定重复逻辑 | 当前不做，保留为最高优先级候选 |
| `source_entity_type / source_entity_id` | 否 | 当前缺少明确消费方 | 暂无已证明的重复解释回收收益 | 当前不做 |
| 聚合计数字段 | 否 | 当前可由 L1 结果局部推导 | 若写入合同将复制既有 facts，多于回收收益 | 当前不做 |

### 3.4 分组汇总

| 分组 | 数量 | 内容 |
|---|---|---|
| 继续复用 L1 真实字段 | 6 | S1-S6 |
| 仅停留在 L3 / renderer 单向映射 | 3 | S7-S9 |
| 当前不做 | 7 | S10-S12、R2-R5 |
| 进入受控派生读取 | 0 | 无 |

---

## 四、承接位矩阵

### 4.1 后端合同与视图承接位

| 层级 | Owner | 输入 | 输出 | 本阶段结论 |
|---|---|---|---|---|
| L1 | `ProjectContextService.GetProjectContext` | `repository_id` | 结构化 project context 事实 | 保持唯一 canonical owner，不新增字段 |
| L2 | `ProjectContextService.ExportProjectContext` | `repository_id` | Markdown 导出结果 | 继续单向派生自 L1，不新增第二结构化合同 |
| L3 | `frontend/src/features/project-context/` | L1 真实字段 + 冻结语义来源 | 共享 query options、adapter、view model | 继续消费 L1，不倒逼后端新增伪统一字段 |
| L4 | feature pages/components/data | L3 结果或切片本地只读结果 | 页面展示与降级 | 保持切片内承接，不反推第二服务 |

### 4.2 Web / agent 共享边界

| 需求 | Web 如何消费 | agent 如何消费 | 是否需要新增后端合同 | 结论 |
|---|---|---|---|---|
| Repository / Product / Module / Decision 摘要 | L3 直接消费 L1 字段 | renderer 直接消费 L1 字段 | 否 | 继续复用 L1 |
| 规则 / phase 入口定位 | L3 读取 `entry_ref / entry_kind / label / summary` | renderer 读取同一字段 | 否 | 继续复用 L1 |
| 统一语义 label / entry view model | L3 做 adapter | renderer 做文本渲染映射 | 否 | 保持单向映射，不新增后端字段 |
| Product / Module / Decision 回锚 `repository_id` | 页面切片结合 detail read 处理 | agent 导出当前不消费反向回锚 | 否 | 当前不做后端 resolver |
| 聚合计数字段 / binding_status | 页面局部或 renderer 局部自行推导 | renderer 若需要也从 L1 局部推导 | 否 | 当前不做后端聚合字段 |

---

## 五、Before / After 样例

### 5.1 共享摘要字段

| Before | After |
|---|---|
| “`entity_name / entity_status / decision_count / binding_status` 都已被后端真实覆盖。” | “后端真实覆盖的只有实体原始摘要字段；跨实体统一命名、计数和 binding 状态要么停留在 L3 / renderer 映射，要么当前不做，不再假装已进入后端合同。” |
| “`entry_title / entry_summary` 已是后端真实字段。” | “后端真实字段仍是 `label / summary / status_summary`；若消费侧想要统一 view model，只能做单向映射，不新增第二套合同命名。” |

### 5.2 Resolver 候选

| Before | After |
|---|---|
| “Decision resolver 已被既有 detail read 覆盖，因此直接当前不做。” | “Decision resolver 已完成正式评估：当前仍只能经 module 链式解析；由于尚未证明它能回收 Web / agent 双侧稳定重复解释逻辑，所以暂不进入后端受控派生读取，但保留为最高优先级候选。” |
| “Product / Module / Decision 专用 `resolved_repository_id` 都无需再讨论。” | “这三类候选都已逐项评估并留档，但当前都不足以进入 `ProjectContextService` 新承接位。” |

### 5.3 真实影响对象

| Before | After |
|---|---|
| “影响对象为 `backend/internal/projectcontext/query_service.go`、`server.go`、`markdown_renderer.go`。” | “影响对象使用真实路径：`service/query_service.go`、`connect/server.go`、`renderer/markdown.go`、`candidate/context_readers.go`。” |

---

## 六、明确不做清单

1. 不新增 `.proto` 字段、消息、RPC 或第二服务
2. 不新增 `ProjectContextService` 下的受控派生读取方法
3. 不把 L3 adapter / renderer 映射结果误写成后端真实字段
4. 不把 Product / Module / Decision detail read 的页面私有结果误提升为 `GetProjectContext` 共享合同
5. 不为 `source_entity_type / source_entity_id` 新增字段
6. 不为聚合计数或 `binding_status` 新增 project-context 共享字段
7. 不把 `ExportProjectContext` 或 renderer 升格为并列事实源
8. 不把 `Decision` resolver 写成“既有能力已覆盖”，除非未来有真实后端承接位落地

---

## 七、与上游文档的一致性声明

| 上游文档 | 关键冻结内容 | 本设计对齐 |
|---|---|---|
| `phase12-05 design §2.3` | 通用字段不得假装已覆盖 | §3.1-§3.4 明确拆成 L1 / L3 / 当前不做 / 候选评估 |
| `phase12-05 design §3.2-§3.3` | Decision resolver 必须进入 `phase12-07` 候选评估 | §3.2-§3.3 已完成正式评估，但未越权写成“既有 detail read 已覆盖” |
| `phase12-05 design §6.3-§6.4` | 是否进入受控派生读取要看重复解释回收收益 | §3.3 逐项记录是否回收双侧重复逻辑 |
| `phase12-03 spec` | 不新增第二服务、第二事实源 | §4-§6 全部继承 |
| `architecture_plan / shared_baseline` | `05 -> 07 -> 06` 顺序与统一设计模板 | §1-§6 满足模板，且不逆行改写 `phase12-05` |
