# PSCO Decision Center 规格 v0.1 — 正式规格正文

> **文档定位**：本文档是 `phase03_decision_center_foundation` 的正式规格正文，作为后续 `phase03-11 / 12 / 13 / 14 / 15` 合同落地、实现、验收与收口以及 `phase04` 引用 `Decision Center` 主线时的直接上游规格来源。
> **上游收敛**：本文档由 `phase03-01` 到 `phase03-09` 的冻结结论收敛而成，不另立第二套边界。
> **互链前提**：本文档以 `phase01-06` 的 `mvp_spec_v0.1.md` 为当前阶段唯一执行层上游，完整承接 `module_registry_spec_v0.1.md` 与 `phase02-12` 验收结论中已交付的 `Module Registry` 边界，并与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致。

---

## 1. 技术路线

本文档继承 `mvp_spec_v0.1.md` §1 的技术路线，聚焦 `Decision Center` 当前阶段：

- 项目路线：`Durable System Track`
- 正式运行主线：`React Web + Go Backend + PostgreSQL`
- 前端：`React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`
- 前端交付策略：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI；不引入独立 `React Native` 客户端，`PWA` 仅作可兼容增强方向
- 后端：`Go`，模块化单体优先
- 数据库：`PostgreSQL` 为唯一数据库主线
- 合同：`Contract First`，跨语言标准为 `Protocol Buffers`；当前阶段必须为 `Decision Center` 落地最小 `.proto` 合同源
- 部署：`Caddy + systemd`，运行方式 `Single Server First`

> 约束：不得重新解释为 `Product Track`；不得把 `Rust` 写成当前阶段必需项；`Local First` 当前解释为数据所有权优先，不等于切换到 `SQLite`。

---

## 2. 对象范围

### 2.1 核心对象（`Decision Center` 直接承接）

- `Decision`：决策结构化记录的核心实体
- `DecisionLink`：`Decision` 与目标对象（当前阶段为 `Module`）的关联关系

### 2.2 连接对象（只读或候选读取，不承接写入主线）

- `Module`：作为 `LinkDecisionToTarget` 的唯一正式目标类型与 `DecisionModuleCandidateRead` 的候选来源
- `Product`：保留为受控连接位，当前阶段不承接 `Decision -> Product` 正式写入
- `Repository`：保留为受控连接位，当前阶段不承接 `Decision -> Repository` 正式写入

### 2.3 非归属能力

- `CreateModule`、`CreateProduct`、`CreateRepository` 不归属于当前阶段 `Decision Center` 后端模块
- `Decision Center` 只能通过最小连接边界读取或校验 `Module`，不得吸收其他对象的主线写入

---

## 3. 页面矩阵

### 3.1 页面主线

| 页面 | 职责 |
| --- | --- |
| `Decision Center / List` | 承接决策列表读取、筛选入口、创建入口与进入详情入口 |
| `Decision Create` | 承接 `RecordDecision` |
| `Decision Detail` | 承接详情读取、已关联目标展示、`Decision -> Module` 候选读取与 `LinkDecisionToTarget` |

### 3.2 轻量跳转或关联入口

- `Module Detail`：当前阶段只承接两个单值轻量触点（记录决策 / 查看相关决策），不扩写为第二个 `Decision` 工作台

> 约束：上述轻量入口不得扩写为当前阶段第二条页面主线。

### 3.3 页面跳转关系

- `Decision Center / List` → `Decision Create`
- `Decision Center / List` → `Decision Detail`
- `Decision Create` → `Decision Detail`（`RecordDecision` 成功后默认回流到新建 `Decision` 的详情页）
- `Decision Detail` → `Decision Center / List`
- `Module Detail` "为当前 `Module` 记录决策"触点 → 带上下文的 `Decision Create`
- `Module Detail` "查看当前 `Module` 相关决策"触点 → `Decision Center / List`

### 3.4 页面信息区块

#### 列表页最小信息区块

- **列表工具栏区**：承接搜索输入、状态筛选与进入 `Decision Create` 的入口
- **列表内容区**：至少展示 `title / status / created_at / link_count / linked_module_summary`
- **空状态区**：无决策时引导用户进入 `Decision Create`

#### 创建页最小信息区块

- **结构化表单区**：承接 `title / context / problem / alternatives / choice / reason / impact / status` 字段
- **来源上下文区**：从 `Module Detail` 带上下文进入时展示预填的 `Module` 信息
- **提交取消操作区**：承接 `RecordDecision` 提交与返回列表路径

#### 详情页最小信息区块

- **核心字段区**：展示 `title / context / problem / alternatives / choice / reason / impact / status / created_at`
- **已关联目标区**：展示当前已建立的 `Decision -> Module` 关联结果
- **候选读取及目标关联区**：承接 `Decision -> Module` 候选读取与 `LinkDecisionToTarget` 写入触点

---

## 4. 动作矩阵

### 4.1 直接承接动作

| 动作 | 页面拥有者 | 写入数据载体 |
| --- | --- | --- |
| `RecordDecision` | `Decision Create` | `decisions` |
| `LinkDecisionToTarget` | `Decision Detail` | `decision_links` |

### 4.2 允许最小入口但不扩写为独立主线

| 动作 | 承接方式 |
| --- | --- |
| `Decision -> Product` | 只保留合同保留位或轻量候选读取前提，不扩写为当前阶段正式写入主线 |
| `Decision -> Repository` | 只保留合同保留位或轻量候选读取前提，不扩写为当前阶段正式写入主线 |

### 4.3 非归属动作

- `CreateModule`、`CreateProduct`、`CreateRepository` 不在当前阶段 `Decision Center` 后端模块中实现
- `Decision Center` 不得调用或吸收 `Module Registry` 的主线写入能力

---

## 5. 数据模型

### 5.1 直接承接表

- `decisions`
- `decision_links`

### 5.2 最小读取或校验前提

- `modules`（候选读取与目标存在性校验）

### 5.3 候选读取前提（只读，不要求写入主线）

- `products`（保留位，当前阶段不要求实现 `Decision -> Product` 候选读取）
- `repositories`（保留位，当前阶段不要求实现 `Decision -> Repository` 候选读取）

### 5.4 对象最小字段

- **`Decision`**：`id`、`title`、`context`、`problem`、`alternatives`、`choice`、`reason`、`impact`、`status`、`created_at`
- **`DecisionLink`**：`id`、`decision_id`、`module_id`、`created_at`

### 5.5 Decision 结构化模板

最小字段集合冻结为：

- `title`
- `context`
- `problem`
- `alternatives`
- `choice`
- `reason`
- `impact`
- `status`

字段级冻结如下：

- **创建必填**：`title / context / problem / choice / reason / status`
- **创建可选**：`alternatives / impact`
- `alternatives` 冻结为按顺序保留的文本条目集合；当前阶段不引入嵌套对象结构（如 `label / score / vote / source`）
- `alternatives` 条目允许为空集合；单个条目在去首尾空白后不得为空字符串
- 空字符串不得视为合法必填值；写入前必须完成去首尾空白后的最小非空校验
- 不得用隐式默认值替代必填字段输入

### 5.6 status 枚举与状态语义

当前阶段 `status` 最小枚举冻结为：

| 状态 | 语义 |
| --- | --- |
| `proposed` | 该决策已经被记录，但尚未成为当前生效结论 |
| `active` | 该决策是当前仍然生效或正在执行的结论 |
| `superseded` | 该决策曾经生效，但已被后续决策替代 |
| `archived` | 该决策已归档保留，不再作为当前执行结论 |

> 约束：不得额外引入 `draft / pending_approval / rejected / auto_generated` 等新状态。

### 5.7 decisions 表演进兼容基线

- `phase02` 中仅承接只读入口的 `decisions` 表（`id / title / created_at`）在 `phase03` 中必须原位升级为结构化主线，不得并行新建替代表后再临时双写
- 原位升级通过 `ALTER TABLE decisions ADD COLUMN` 完成，新增字段：`context / problem / alternatives TEXT[] / choice / reason / impact / status`
- 必须添加 `status` 的 `CHECK` 约束：`CHECK (status IN ('proposed', 'active', 'superseded', 'archived'))`
- 必须为列表读取性能添加索引：`CREATE INDEX idx_decisions_status_created_at ON decisions (status, created_at DESC)`
- 对于无 `DEFAULT` 的必填字段（`context / problem / choice / reason`），必须按"先 `ADD COLUMN` 允许 `NULL` -> 回填 -> `SET NOT NULL`"三步流程添加
- 对于有 `DEFAULT` 的字段（`alternatives TEXT[] NOT NULL DEFAULT '{}'` / `impact TEXT NOT NULL DEFAULT ''` / `status TEXT NOT NULL DEFAULT 'proposed'`），可直接 `ADD COLUMN ... NOT NULL DEFAULT ...`
- `alternatives` 必须使用 `TEXT[]`（PostgreSQL 原生数组），不得使用 `JSONB`，避免引入与 `.proto` `repeated string` 不必要的序列化层
- 不得删除原有 `title / created_at` 字段，不得破坏既有 `decision_links` 的外键引用
- 现有示例 `Decision` 数据必须通过迁移脚本完成兼容回填，满足当前阶段最小结构化模板的非空约束
- 兼容回填必须保留原有 `title / created_at`；`context / problem / choice / reason` 回填为占位文本（如 `（历史决策，phase03 升级前无结构化上下文）`），`alternatives` 回填为 `'{}'`，`impact` 回填为 `''`，`status` 回填为 `'proposed'`
- 回填 `UPDATE` 仅对无 `DEFAULT` 的必填字段实际生效；`DEFAULT` 字段由 `ADD COLUMN ... NOT NULL DEFAULT ...` 自动填充，回填 `UPDATE` 对它们为 `no-op`，但保留在回填语句中不影响幂等性
- 迁移与 seed 必须可重复执行，当前阶段不得依赖手工 SQL 临时修补历史样例

### 5.8 最小读写模型

#### 列表读取

至少承接：`title / status / created_at / link_count / linked_module_summary`

#### 详情读取

至少承接：核心对象字段（`id / title / context / problem / alternatives / choice / reason / impact / status / created_at`）、结构化模板字段、已关联 `Module` 列表与最小来源上下文

#### 创建写入

承接 `RecordDecision`（最小字段 `title / context / problem / alternatives / choice / reason / impact / status`）

#### 关联写入

承接 `LinkDecisionToTarget`（最小字段 `decision_id / target_type / module_id`）

### 5.9 link_count 与 linked_module_summary 计算口径

- `link_count` 当前阶段仅统计 `decision_links` 中已建立的 `Decision -> Module` 有效关联数
- `linked_module_summary` 仅基于已关联 `Module` 生成，不混入 `Product / Repository`
- `linked_module_summary` 按 `module_name` 升序取前 `3` 个名称；若超出 `3` 个，则在摘要末尾附加 `+N`
- 当 `Decision` 当前没有任何已关联 `Module` 时，`link_count` 返回 `0`，`linked_module_summary` 返回空字符串，不返回 `null`
- 入口上下文中的预填 `Module` 在正式 `LinkDecisionToTarget` 写入前不计入 `link_count` 与 `linked_module_summary`

### 5.10 目标候选读取基线

- 候选来源为当前已存在的 `modules`
- 候选范围同时覆盖 `active` 与 `archived` 的 `Module`，避免历史决策无法关联历史模块
- 候选排序采用 `status(active 优先) -> module_name 升序`
- 已建立 `Decision -> Module` 关联的目标不得再次出现在可关联候选中
- 无可关联候选时，页面必须返回明确空状态，而不是把空结果误报为接口错误

### 5.11 入口上下文与正式关联结果边界

- **入口上下文**：从 `Module Detail` 带上下文进入 `Decision Create` 时，该 `Module` 信息只代表预填来源或待关联目标，不等于已落库关联
- **正式关联结果**：只有在 `Decision Detail` 中完成 `LinkDecisionToTarget` 后，才能视为正式建立 `Decision -> Module` 关联
- 在正式关联写入前，`Decision List` 的 `link_count / linked_module_summary` 不得将该预填来源计入已关联结果
- 带 `Module` 上下文进入 `Decision Create` 后，创建成功默认进入新建 `Decision` 的 `Decision Detail`，不得回流到 `Decision Center / List`
- 该入口上下文中的 `Module` 必须带入 `Decision Detail` 作为显式待关联目标继续承接，持续到用户完成正式 `LinkDecisionToTarget` 或主动放弃关联
- 无特定目标上下文进入创建页时（直接从 `Decision Center / List` 进入），必须承接"无特定来源目标"的语义，不得伪造默认 `Module` 目标

---

## 6. API 边界与接口分组

### 6.1 合同边界

- 遵守 `Contract First`，以结构化合同优先
- 跨语言合同标准冻结为 `Protocol Buffers`
- 当前阶段必须落地 `Decision Center` 最小 `.proto` 合同源
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层，但不得形成与 `.proto` 并列的第二套合同源
- 不得引入与 `Protocol Buffers` 冲突的第二套跨语言合同主线

### 6.2 接口分组

#### 读组

| 接口 | 职责 |
| --- | --- |
| `DecisionListRead` | 只承接列表展示所需最小读取、筛选前提与进入详情所需标识 |
| `DecisionDetailRead` | 统一承接决策核心字段、结构化模板字段、最小来源上下文与已关联 `Module` 列表读取 |
| `DecisionModuleCandidateRead` | 只承接 `Decision Detail` 中 `Decision -> Module` 的候选读取 |

#### 写组

| 接口 | 职责 |
| --- | --- |
| `DecisionWrite` | 只承接 `RecordDecision`，返回新建 `Decision` 标识 |
| `DecisionLinkWrite` | 只承接 `LinkDecisionToTarget`，当前阶段唯一允许的正式目标类型为 `Module` |

#### 方向级最小 API 矩阵

| 动作语义 | 接口分组 | 接口名 | 方向 |
| --- | --- | --- | --- |
| 列表读取 | 读组 | `DecisionListRead` | 读 |
| 详情读取 | 读组 | `DecisionDetailRead` | 读 |
| `Module` 候选读取 | 读组 | `DecisionModuleCandidateRead` | 读 |
| `RecordDecision` | 写组 | `DecisionWrite` | 写 |
| `LinkDecisionToTarget` | 写组 | `DecisionLinkWrite` | 写 |

> 本矩阵与 `mvp_spec_v0.1.md` §6.2 最小 API 矩阵互链，仅承接 `Decision Center` 当前阶段涉及的语义，不扩写为完整跨模块服务矩阵。

### 6.3 详情读取与候选读取边界

- `DecisionDetailRead` 只承接详情本体、最小来源上下文与已建立关联结果
- `DecisionModuleCandidateRead` 必须通过独立 request / response 承接候选读取
- 不得把候选读取结果直接并入 `DecisionDetailRead` 的最小合同边界
- 不得把候选读取拆成需要前端自行拼装的多个独立业务入口

### 6.4 关系写入后的读取语义

- `LinkDecisionToTarget` 提交成功后，默认后端语义为"详情页重新读取当前已关联目标结果"
- 不得额外设计一套脱离 `DecisionDetailRead` 的回流读取路径

### 6.5 错误语义

| 场景 | 语义 | 归属接口 |
| --- | --- | --- |
| `RecordDecision` 必填字段缺失 | 校验失败 | `DecisionWrite` |
| `RecordDecision` 字段值非法（含 `alternatives` 条目空白、非法 `status`） | 校验失败 | `DecisionWrite` |
| `DecisionDetailRead` 不存在的 `decision_id` | 资源不存在 | `DecisionDetailRead` |
| `DecisionModuleCandidateRead` 无可关联候选 | 空列表语义（非错误） | `DecisionModuleCandidateRead` |
| `LinkDecisionToTarget` 目标类型越界（非 `MODULE`） | 校验失败 | `DecisionLinkWrite` |
| `LinkDecisionToTarget` 目标 `Decision` 或 `Module` 不存在 | 资源不存在 | `DecisionLinkWrite` |
| `LinkDecisionToTarget` 重复关联 | 重复冲突 | `DecisionLinkWrite` |

> 约束：不得出现 `500` 级未收口错误替代业务错误；不得把空候选结果误报为接口失败。

### 6.6 合同与存储解耦

- 不得直接将 `decisions`、`decision_links` 的存储模型原样暴露为外部合同
- `.proto` 必须作为当前阶段唯一合同源；若同时存在 `JSON HTTP` 过渡接口，其消息语义必须从 `.proto` 派生或与 `.proto` 严格对齐
- 必须允许后续在 `Contract First` 路线下独立演进接口消息结构
- 不得复用已删除字段编号或字段语义，必须保持与当前 `.proto` 合同主线兼容

---

## 7. 合同设计（Proto 主线）

### 7.1 单一 Proto 合同源

- `.proto` 是当前阶段唯一合同源
- 不得再以手写 JSON 结构充当并列合同源
- 稳定目录落点为 `proto/psco/decision_center/v1/decision_center.proto`
- 单一包名与版本语义冻结为 `psco.decision_center.v1`

### 7.2 最小服务接口矩阵

| RPC | 对齐接口 | 对齐页面动作 |
| --- | --- | --- |
| `ListDecisions` | `DecisionListRead` | `Decision Center / List` 列表读取 |
| `GetDecisionDetail` | `DecisionDetailRead` | `Decision Detail` 详情读取 |
| `CreateDecision` | `DecisionWrite` | `Decision Create` 创建写入 |
| `LinkDecisionToTarget` | `DecisionLinkWrite` | `Decision Detail` 关联写入 |
| `ListDecisionModuleCandidates` | `DecisionModuleCandidateRead` | `Decision Detail` 候选读取 |

### 7.3 核心消息结构与字段编号

#### `Decision`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `id` | 1 | `string` | 决策标识 |
| `title` | 2 | `string` | 决策标题（必填） |
| `context` | 3 | `string` | 决策上下文（必填） |
| `problem` | 4 | `string` | 决策问题（必填） |
| `alternatives` | 5 | `repeated string` | 候选方案（可选，按顺序保留） |
| `choice` | 6 | `string` | 最终选择（必填） |
| `reason` | 7 | `string` | 选择理由（必填） |
| `impact` | 8 | `string` | 影响评估（可选） |
| `status` | 9 | `DecisionStatus` | 决策状态（必填） |
| `created_at` | 10 | `google.protobuf.Timestamp` | 创建时间 |

#### `DecisionListItem`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `id` | 1 | `string` | 决策标识 |
| `title` | 2 | `string` | 决策标题 |
| `status` | 3 | `DecisionStatus` | 决策状态 |
| `created_at` | 4 | `google.protobuf.Timestamp` | 创建时间 |
| `link_count` | 5 | `int32` | 已关联 `Module` 数 |
| `linked_module_summary` | 6 | `string` | 已关联 `Module` 摘要（含 `+N` 语义） |

#### `LinkedModule`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `module_id` | 1 | `string` | 模块标识 |
| `module_name` | 2 | `string` | 模块名称 |

#### `SourceContext`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `source_module_id` | 1 | `string` | 来源模块标识（无来源时为空字符串） |
| `source_module_name` | 2 | `string` | 来源模块名称（无来源时为空字符串） |

#### `DecisionDetail`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `decision` | 1 | `Decision` | 决策核心对象 |
| `linked_modules` | 2 | `repeated LinkedModule` | 已关联 `Module` 列表 |
| `source_context` | 3 | `SourceContext` | 最小来源上下文 |

#### `DecisionModuleCandidate`

| 字段 | 编号 | 类型 | 说明 |
| --- | --- | --- | --- |
| `module_id` | 1 | `string` | 候选模块标识 |
| `module_name` | 2 | `string` | 候选模块名称 |
| `status` | 3 | `ModuleStatus`（跨包 import） | 候选模块状态 |

### 7.4 枚举冻结

#### `DecisionStatus`

| 枚举值 | 编号 |
| --- | --- |
| `DECISION_STATUS_UNSPECIFIED` | 0 |
| `DECISION_STATUS_PROPOSED` | 1 |
| `DECISION_STATUS_ACTIVE` | 2 |
| `DECISION_STATUS_SUPERSEDED` | 3 |
| `DECISION_STATUS_ARCHIVED` | 4 |

#### `DecisionLinkTargetType`

| 枚举值 | 编号 |
| --- | --- |
| `DECISION_LINK_TARGET_TYPE_UNSPECIFIED` | 0 |
| `DECISION_LINK_TARGET_TYPE_MODULE` | 1 |

> 约束：当前阶段只允许 `MODULE`，不得把 `Product / Repository` 提前写成当前阶段正式可用枚举值。

### 7.5 Request / Response 字段编号

#### 读组

| 消息 | 字段 | 编号 | 类型 |
| --- | --- | --- | --- |
| `ListDecisionsRequest` | `query_text` | 1 | `string` |
| `ListDecisionsRequest` | `status_filter` | 2 | `DecisionStatus` |
| `ListDecisionsResponse` | `decisions` | 1 | `repeated DecisionListItem` |
| `GetDecisionDetailRequest` | `decision_id` | 1 | `string` |
| `GetDecisionDetailResponse` | `decision_detail` | 1 | `DecisionDetail` |
| `ListDecisionModuleCandidatesRequest` | `decision_id` | 1 | `string` |
| `ListDecisionModuleCandidatesResponse` | `candidates` | 1 | `repeated DecisionModuleCandidate` |

#### 写组

| 消息 | 字段 | 编号 | 类型 |
| --- | --- | --- | --- |
| `CreateDecisionRequest` | `title` | 1 | `string` |
| `CreateDecisionRequest` | `context` | 2 | `string` |
| `CreateDecisionRequest` | `problem` | 3 | `string` |
| `CreateDecisionRequest` | `alternatives` | 4 | `repeated string` |
| `CreateDecisionRequest` | `choice` | 5 | `string` |
| `CreateDecisionRequest` | `reason` | 6 | `string` |
| `CreateDecisionRequest` | `impact` | 7 | `string` |
| `CreateDecisionRequest` | `status` | 8 | `DecisionStatus` |
| `CreateDecisionResponse` | `decision_id` | 1 | `string` |
| `LinkDecisionToTargetRequest` | `decision_id` | 1 | `string` |
| `LinkDecisionToTargetRequest` | `target_type` | 2 | `DecisionLinkTargetType` |
| `LinkDecisionToTargetRequest` | `module_id` | 3 | `string` |
| `LinkDecisionToTargetResponse` | — | — | 空响应（无字段） |

> `LinkDecisionToTargetResponse` 必须为空响应，对齐"详情页重新读取当前已关联目标结果"的回流语义，不得返回 link 结果或 detail reread 标识。

### 7.6 ModuleStatus 跨包依赖策略

- 必须通过 `import "psco/module_registry/v1/module_registry.proto"` 直接复用 `psco.module_registry.v1.ModuleStatus`
- 不得在 `psco.decision_center.v1` 中重定义本地等价枚举，避免引入两套 `status` 语义
- 该 import 仅限于复用 `ModuleStatus` 枚举类型，不得因此引入对 `Module Registry` service 或其他消息结构的合同依赖

### 7.7 RPC → HTTP 映射矩阵

| RPC | HTTP 方法 | 路径 |
| --- | --- | --- |
| `ListDecisions` | `GET` | `/api/decisions` |
| `GetDecisionDetail` | `GET` | `/api/decisions/{decisionId}` |
| `CreateDecision` | `POST` | `/api/decisions` |
| `LinkDecisionToTarget` | `POST` | `/api/decisions/{decisionId}/links` |
| `ListDecisionModuleCandidates` | `GET` | `/api/decisions/{decisionId}/candidates/modules` |

> 约束：JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 语义显式对齐；HTTP 过渡层使用 URL 路径参数承接 `decisionId` 或 `moduleId` 时，handler 必须在进入业务层前显式组装为对应的 Proto request 字段；不得把 HTTP 路径、状态码或中间件策略误写成 Proto 合同本体。

### 7.8 合同演进规则

- 必须为每个字段分配稳定编号，不得复用已删除字段编号或字段语义
- 后续版本删除字段或废弃字段名时必须使用 `reserved` 保留字段号，必要时同时保留字段名
- 后续新增枚举值必须使用递增编号，不得插入到已有编号之间
- `buf` 校验链必须至少覆盖 `buf build`、`buf lint`、`buf generate`、`buf breaking`
- `buf breaking` 必须直接对照仓库主线 `.git` 基准，不得吞掉失败退出码

---

## 8. 空状态与冷启动

### 8.1 空状态引导

- 用户首次进入 `Decision Center / List` 且系统中尚无任何决策时，页面必须展示明确的空状态提示
- 空状态主动作必须直接进入 `Decision Create`
- 不得因空状态返回接口错误

### 8.2 冷启动验收路径

冷启动验收路径冻结为以下 `8` 步：

1. 执行 `reset_decision_mainline.sh --clean-only`，使 `decisions` 进入空状态
2. 进入 `Decision Center / List`，验证空状态入口展示
3. 从空状态进入 `Decision Create`
4. 填写最小结构化模板，提交 `RecordDecision`
5. 验证回流到 `DecisionDetailPage`，详情页显示新建 `Decision` 核心字段
6. 在 `DecisionDetailPage` 触发 `DecisionModuleCandidateRead`，验证候选列表展示
7. 选择一个 `Module`，执行 `LinkDecisionToTarget`
8. 验证详情页重新读取并显示已关联 `Module` 结果

> 约束：该路径必须在不执行任何手工 `SQL` 的前提下完成；前提是 `reset_module_mainline.sh` 已建立 `modules` 基线（候选读取需要 `modules`）。

### 8.3 创建最小闭环

- 最小闭环为 `Decision Center / List → Decision Create → Decision Detail`
- `RecordDecision` 成功后默认回流到新建 `Decision` 的 `DecisionDetailPage`，不得回流到 `Decision Center / List`
- `LinkDecisionToTarget` 成功后用户必须停留在当前 `DecisionDetailPage`，并重新读取已关联目标结果

### 8.4 返回路径

- 从 `Decision Create` 主动取消或返回：默认返回保留原搜索参数上下文的 `DecisionListPage`
- 从 `Decision Detail` 主动返回 `Decision Center / List`：必须恢复原有的 `queryText` 与 `statusFilter`，不得要求用户重新输入筛选条件
- 从创建成功后的详情页再返回列表：详情页已继承来源列表上下文时，必须恢复进入创建前原有的 `queryText` 与 `statusFilter`；来源列表上下文不存在（如从 `Module Detail` 直接进入创建）时，返回列表必须落到默认列表参数

---

## 9. 前端实现设计层结果

### 9.1 页面文件落点

| 页面 | 文件路径 |
| --- | --- |
| `DecisionListPage` | `frontend/src/features/decision-center/pages/decision-list-page.tsx` |
| `DecisionCreatePage` | `frontend/src/features/decision-center/pages/decision-create-page.tsx` |
| `DecisionDetailPage` | `frontend/src/features/decision-center/pages/decision-detail-page.tsx` |

### 9.2 路由树与 URL 语义

| 路由 | URL | 文件路径 |
| --- | --- | --- |
| `DecisionListRoute` | `/decisions` | `frontend/src/routes/decisions/index.tsx` |
| `DecisionCreateRoute` | `/decisions/new` | `frontend/src/routes/decisions/new.tsx` |
| `DecisionDetailRoute` | `/decisions/:decisionId` | `frontend/src/routes/decisions/$decisionId.tsx` |

> 约束：不得把 `Product`、`Repository` 提前扩写为并列主树；当前阶段不得把 `Decision Detail` 提前拆成独立的子路由工作台。

### 9.3 Module Detail 入口映射

- `frontend/src/routes/modules/$moduleId.tsx` 承载的 `ModuleDetailPage` 继续作为 `Module Detail` 的页面宿主
- `frontend/src/features/module-registry/components/module-decision-entry-panel.tsx` 中的 `ModuleDecisionEntryPanel` 必须从只读展示升级为承接两个入口动作的触点组件：
  - "为当前 `Module` 记录决策"触点导航到 `/decisions/new`，携带当前 `Module` 的目标标识与可展示名称作为上下文
  - "查看当前 `Module` 相关决策"触点导航到 `/decisions`
- 不得在 `Module Detail` 侧新增中间分发组件或中间路由层

### 9.4 组件树

#### 列表页

- `DecisionListPageShell`
  - `DecisionListToolbar`（承接搜索输入、状态筛选与创建入口）
  - `DecisionListContent`（只承接列表读取结果）
  - `DecisionListTableOrCards`
  - `DecisionListEmptyState`（只承接无决策时进入 `Decision Create` 的引导）

#### 创建页

- `DecisionCreatePageShell`
  - `DecisionContextSourcePanel`（只承接从 `Module Detail` 带入的来源 `Module` 展示）
  - `DecisionCreateForm`（只承接结构化模板字段录入）
  - `DecisionCreateActions`（只承接提交与取消动作）

#### 详情页

- `DecisionDetailPageShell`
  - `DecisionDetailSummaryCard`（只承接决策核心字段、结构化模板字段与最小来源上下文展示）
  - `DecisionLinkedTargetsSection`（只承接已建立的 `Decision -> Module` 关联结果）
  - `DecisionPendingLinkTargetCard`（只承接入口上下文中尚未完成正式关联的待关联 `Module`，必须作为显式待关联目标持续展示，直到用户完成 `LinkDecisionToTarget` 或主动放弃关联）
  - `DecisionModuleCandidatePanel`（只承接 `Decision -> Module` 候选读取与目标选择）
  - `DecisionLinkActions`（只承接 `LinkDecisionToTarget` 的最小写入触点）

#### 组件归属原则

- 页面专属组件默认归属于对应页面
- 只有在多个页面确实共享同一职责时，才允许抽为共享组件
- 当前阶段不得为了"组件纯洁"提前拆出无明确复用证据的通用组件层

### 9.5 状态模型

#### Decision List

- 查询条件冻结到路由搜索参数层：`queryText`、`statusFilter`
- `queryText` 默认为空字符串（不执行文本筛选）
- `statusFilter` 默认为"全部状态"（不执行状态筛选）
- 不得将默认值解释为 `null` 或跳过列表读取
- 从 `DecisionCreatePage` 或 `DecisionDetailPage` 返回时，必须按原有 `queryText` 与 `statusFilter` 恢复列表上下文
- 刷新页面后，若路由搜索参数仍在，列表必须按该参数恢复读取状态
- 读取状态：`pending`、`success`、`error`
- 派生视图状态：`initial-loading`、`ready`、`empty`、`error`
- 列表读取成功但没有任何 `Decision` 时进入 `empty`，空状态主动作必须直接进入 `DecisionCreatePage`
- 错误反馈停留在 `DecisionListPage` 内容区域，不跳转独立错误页

#### Decision Create

- 来源上下文状态：至少包含来源 `Module` 的目标标识与可展示名称；只承接展示与后续回流，不等于已完成正式关联
- 草稿状态：至少承接 `title / context / problem / alternatives / choice / reason / impact / status`；至少能区分 `idle` 与 `dirty`
- 提交状态：`submitting`、`submit-success`、`submit-error`
- 提交失败时停留当前页，保留草稿与来源上下文展示，错误显示在表单上下文
- 提交成功默认回流到新建 `Decision` 的 `DecisionDetailPage`

#### Decision Detail

- 读取状态：`pending`、`success`、`error`
- 待关联目标承接状态：必须继续承接"存在待关联 `Module`"的显式状态，持续到用户完成 `LinkDecisionToTarget` 或主动放弃关联；不得在进入 `DecisionDetailPage` 后静默丢失该上下文
- 候选读取状态：`pending`、`ready`、`empty`、`error`；候选空结果必须进入 `empty`，不得将空结果误解释为资源不存在
- 关联写入状态：`idle`、`submitting`、`submit-success`、`submit-error`；只归属于当前详情页上下文，不升级为跨路由全局状态
- 关联成功后用户必须停留在当前 `DecisionDetailPage`，当前详情页必须承接最新的已关联目标读取结果；若本次关联目标正是入口上下文中的待关联 `Module`，待关联目标状态必须被清除
- 关联失败时页面必须停留在当前 `DecisionDetailPage`，当前选中的候选目标与待关联目标展示必须继续保留，错误反馈停留在当前详情页的关联动作上下文内

#### 跨页面列表上下文单值承接

- 创建页来源列表上下文状态：
  - 从 `DecisionListPage` 进入时，来源列表上下文必须存在，至少包含当前 `queryText` 与 `statusFilter`
  - 从 `Module Detail` 进入时，来源列表上下文不存在
  - 存在时返回列表必须恢复原有参数；不存在时落到默认列表参数
- 详情页来源列表上下文状态：
  - 从 `DecisionListPage` 进入详情时，来源列表上下文必须存在
  - 从 `DecisionCreatePage` 成功创建后进入详情时，来源列表上下文必须继承自创建页
  - 从 `Module Detail` 直接进入或从外部链接直达时，来源列表上下文不存在
  - 存在时返回列表必须恢复原有参数；不存在时落到默认列表参数

#### 页面级 UI 状态局部归属

- 表单草稿、提交错误、待关联目标展示、候选读取状态与关联动作状态应优先归属于当前页面或当前详情页上下文
- 不得默认升级为跨路由全局状态

### 9.6 布局降级策略

#### PC 页面布局

- `DecisionListPage`：高信息密度列表布局
- `DecisionDetailPage`：分区式详情布局，至少分为概要主区、已关联目标区、待关联目标区、候选读取与目标关联区
  - 待关联目标区在存在入口上下文待关联 `Module` 时必须可见，在无待关联目标时不占用布局空间
  - 候选读取与目标关联区应优先保持同页可见，而不是依赖额外导航切换
- `DecisionCreatePage`：来源上下文区、表单区与动作区同屏可见

#### 移动浏览器布局

- `DecisionListPage`：单列列表或卡片重排
- `DecisionDetailPage`：概要、已关联目标、待关联目标、候选读取与目标关联按垂直顺序重排；可对次级信息采用折叠或延迟展开
- `DecisionCreatePage`：来源上下文区、表单区与动作区采用单列垂直布局；主动作按钮必须在无需横向滚动的前提下可见

> 约束：不得新增独立移动端页面体系；不得引入独立 `React Native` 客户端；不得把完整 `PWA` 能力写成当前阶段实现前提。

### 9.7 运行时实现细节不冻结

- 当前阶段不提前冻结具体 hook 命名、Query key 细节、store API 命名、缓存时间、请求取消策略或 optimistic update 方案
- 页面级 UI 状态优先归属于当前页面或当前详情页上下文，只在确有跨同页多组件共享需要时再抽为页面作用域状态容器

---

## 10. 后端实现设计层结果

### 10.1 模块边界

- `Decision Center` 在后端冻结为单一业务模块
- 统一承接 `RecordDecision`、`LinkDecisionToTarget`、`DecisionListRead`、`DecisionDetailRead`、`DecisionModuleCandidateRead`
- 不得把这些能力拆散到 `Module Registry`、`Product` 或 `Repository` 的后端模块中
- `Decision Center` 只能通过最小连接边界读取或校验 `Module`，不得吸收其他对象的主线写入

### 10.2 文件落点

```
backend/internal/decisioncenter/
  handler/
    query_handler.go              # 读组入口：DecisionListRead + DecisionDetailRead + DecisionModuleCandidateRead
    command_handler.go            # 写组入口：DecisionWrite + DecisionLinkWrite
  service/
    query_service.go              # 读组编排（列表 + 详情 + 候选读取）
    command_service.go            # 写组编排（创建 + 关联）
  repository/
    decision_store.go             # decisions 表读写
    link_store.go                 # decision_links 表读写
  candidate/
    module_candidate_read.go      # ModuleCandidateRead（跨模块候选读取，由 Decision Center 拥有）
  errors.go                       # 业务错误哨兵值
  types.go                        # 跨层共享 API 消息结构（DTO、枚举、请求/响应类型、列表查询参数）
  validate.go                     # 输入校验辅助
  handler/response.go             # HTTP 协议层共享工具（JSON 编解码、错误到状态码映射）
```

### 10.3 分层语义

| 层 | 子包 | 职责 |
| --- | --- | --- |
| 入口层 | `handler/` | 只负责承接请求与返回结果 |
| 业务编排层 | `service/` | 只负责动作语义、校验顺序与跨连接口编排 |
| 数据访问/外部连接层 | `repository/` + `candidate/` | 只负责持久化与依赖调用 |

### 10.4 读组与写组代码组织

- 读组（`DecisionListRead` + `DecisionDetailRead` + `DecisionModuleCandidateRead`）→ `handler/query_handler.go` + `service/query_service.go`（单文件编排，不拆 `list_service.go` / `detail_service.go` / `candidate_service.go`）
- 写组（`DecisionWrite` + `DecisionLinkWrite`）→ `handler/command_handler.go` + `service/command_service.go`（单文件编排，不拆两个独立 service 文件）

### 10.5 ModuleCandidateRead 接口拥有者与接线

- `ModuleCandidateRead` 的接口与实现必须由 `Decision Center` 的 `candidate/` 子包自己定义和拥有
- `Module Registry` 不为 `Decision Center` 暴露专门的服务契约或服务方法
- `candidate/` 子包通过独立 Read 接口隔离，`service/` 层不得直接写跨模块 SQL
- 具体接线（构造与注入）必须在应用装配点完成，不得在 `service/` 或 `handler/` 内部自行构造
- 后续若 `Module Registry` 提供更稳定的读模型协作实现，可迁移 `candidate/` 内的具体实现，但接口契约与拥有者保持不变

### 10.6 支撑文件落点

- `errors.go`：业务错误哨兵值
- `types.go`：跨层共享 API 消息结构（DTO、枚举、请求/响应类型、列表查询参数）
- `validate.go`：输入校验辅助
- `handler/response.go`：HTTP 协议层共享工具（JSON 编解码、错误到状态码映射）
- 支撑文件的组织方式必须与现有 `moduleregistry` 模块的 `errors.go / types.go / validate.go / handler/response.go` 保持同构
- 不得把错误定义、共享类型或校验逻辑散落到 `handler/` 或 `service/` 内部

### 10.7 Module Registry 连接边界

- `Decision Center` 只允许依赖 `DecisionModuleCandidateRead` 一类的最小候选读取接口
- 只允许依赖目标存在性、可展示名称与最小状态语义所需的读取/校验
- 不得在当前阶段调用或吸收 `CreateModule`、`CreateRelease` 或其他 `Module Registry` 主线写入
- `Module Registry` 仍然只作为被连接方提供最小只读协作前提

### 10.8 不冻结的实现工具

- 不得提前冻结 Go HTTP 框架、RPC 框架、ORM、SQL Builder、Repository 模板或目录生成器
- 文件内部的函数签名、中间件选型、数据库访问方式不作为当前阶段冻结内容
- 只冻结职责分工、接口归属与文件落点，不冻结实现手段

---

## 11. 验收基线

### 11.1 联调环境可重复建立

- 复用 `phase02-11` 已冻结的 `database/scripts/init_db.sh` 创建 `PSCO` 独立数据库
- 复用 `database/scripts/run_seeds.sh` 执行只读前提种子（`products / repositories`）
- 复用 `database/scripts/reset_module_mainline.sh` 建立 `Module Registry` 基线
- 不得为 `Decision Center` 新建第二个数据库或第二套 `init_db` 入口

#### 启动顺序

启动顺序必须为：

1. `init_db.sh`
2. 后端启动（自动运行 migration，含 `0004_decision_center_mainline`）
3. `run_seeds.sh`（只读前提，必须在 migration 完成后执行，因为 `seed_readonly_prereqs.sql` 依赖 `0002` 已创建的表与 `0004` 已添加的 `decisions` 结构化字段）
4. `reset_module_mainline.sh`
5. `reset_decision_mainline.sh`（必须在 `reset_module_mainline.sh` 之后执行，因为 `decision_links` 依赖 `modules` 基线数据）
6. 前端启动

### 11.2 Decision Center 重置脚本

- 脚本落点：`database/scripts/reset_decision_mainline.sh`
- 支持三种模式：`--clean-only`（仅清空）、`--restore-only`（仅恢复）、默认（清空 + 恢复）
- 复用 `reset_module_mainline.sh` 的 `resolve_psql` 模式（宿主机 psql -> docker exec -> podman exec 自动检测）
- 清空范围：`DELETE FROM decisions`，依赖 `decision_links.decision_id` 的 `ON DELETE CASCADE` 级联清空 `decision_links`；不得清空 `modules / products / repositories` 等表
- 复用 `reset_module_mainline.sh` 的环境变量覆盖参数（`PG_HOST / PG_PORT / PG_USER / PSCO_DB / PG_CONTAINER / SEEDS_DIR`）与 `-h / --help` 用法说明
- 清空后必须输出当前 `decisions` 与 `decision_links` 计数以供确认
- 前置校验：数据库必须已存在；`--restore-only` 与默认模式下必须校验 `modules` 基线数据已存在

### 11.3 基线种子数据

- seed 文件落点：`database/seeds/seed_decision_mainline_baseline.sql`
- 该文件同时承担"清空 + 恢复"职责，开头包含 `DELETE FROM decisions`，后续 `INSERT` 使用 `ON CONFLICT DO NOTHING`，通过 `BEGIN / COMMIT` 事务包裹
- 基线 `Decision` 数据至少包含 `3` 条，覆盖 `proposed / active / archived|superseded` 三个 `status` 维度
- 其中 `1` 条必须保留 `phase02` 原有 `title`（`关于 auth-service 技术选型的决策`），并补全结构化字段
- 至少 `1` 条 `Decision` 必须包含 `alternatives` 数组（`2` 个以上条目），验证数组保序语义
- 至少 `1` 条 `Decision` 的 `alternatives` 必须为空数组，验证空数组语义
- 至少 `1` 条 `Decision` 的 `impact` 必须为空字符串，验证可选字段空值语义
- 基线 `decision_links` 至少包含 `2` 条，关联到 `reset_module_mainline` 已建立的 `modules` 基线
- 至少 `1` 条 `Decision` 必须没有任何 `decision_links`，验证 `link_count = 0` 与 `linked_module_summary = ''` 的空值语义
- `decision_links` 的 `INSERT` 必须通过 `module_name` 与 `decision title` 查找 `ID`，不得硬编码 `UUID`
- `seed_decision_mainline_baseline.sql` 承担 `decision_links` 的最终基线职责；`seed_module_mainline_baseline.sql` 中既有的 `decision_link` 插入在 `reset_decision_mainline.sh` 执行时会被 `DELETE FROM decisions` 级联删除后由 `decision` 基线 `seed` 重建，最终状态以 `seed_decision_mainline_baseline.sql` 为准
- `seed_readonly_prereqs.sql` 中的 `decisions` seed 必须从 `title-only` 更新为结构化字段插入，保持原有 `title` 以兼容 `phase02` `decision_links`

### 11.4 异常路径验证

异常路径验证必须至少覆盖以下 `7` 类（对齐 `phase03-04` 已冻结的全部错误语义）：

1. `RecordDecision` 必填字段缺失 → 返回校验失败
2. `RecordDecision` 字段值非法（含 `alternatives` 条目空白、非法 `status` 值）→ 返回校验失败
3. `LinkDecisionToTarget` 目标类型越界（非 `MODULE`）→ 返回校验失败
4. `LinkDecisionToTarget` 目标 `Decision` 或 `Module` 不存在 → 返回资源不存在语义
5. `LinkDecisionToTarget` 重复关联 → 返回重复冲突语义
6. `DecisionModuleCandidateRead` 无可关联候选 → 返回空列表语义（非错误）
7. `DecisionDetailRead` 不存在的 `decision_id` → 返回资源不存在语义

> 约束：异常路径验证通过 `API` 层测试触发，不通过 `seed` 异常数据；不得为异常路径验证新建单独的 `fixture SQL` 文件；不得出现 `500` 级未收口错误替代业务错误。"无可关联候选"验证通过 `API` 操作建立前置条件（创建一个已关联全部 `active Module` 的 `Decision`），不需要 `seed` 额外数据。

### 11.5 验收环境前置规划

- `phase03-14` 验收必须能直接复用本规格冻结的脚本、seed 与验收路径
- 不得在 `phase03-14` 验收时新建临时 `SQL` 或临时脚本
- 不得在 `phase03-14` 验收时发现缺少重置入口而阻塞

---

## 12. 非目标矩阵

当前阶段明确不做：

- `Product` 全量主线
- `Repository` 全量主线
- `Dashboard` 聚合反馈
- `pending_decision_signals` 的完整消费链
- 跨页面聚合查询、趋势分析查询
- 自动扫描代码
- 自动知识图谱
- `AI` 自动建议或自动裁决
- 独立 `AI Assistant` 工作台
- 独立 `React Native` 客户端
- 完整 `PWA` 能力落地
- 复杂审批流、投票流或自动推荐类字段
- `Decision -> Product` 或 `Decision -> Repository` 的正式写入闭环
- 第二套移动端 UI 架构
- 完整跨模块服务矩阵（若后续需要，应在对应新 `phase` 或 `audit` 中单独冻结）

---

## 13. Done 标准

当以下条件成立时，`Decision Center` 当前阶段规格成立并可进入 `phase03-11 / 12 / 13 / 14 / 15`：

1. 页面矩阵与动作矩阵已单值化且一一对应
2. 数据模型与读写模型已明确，`decisions` 表原位升级与兼容回填策略已冻结
3. 接口分组（读组 + 写组）已明确，错误语义归属已单值化
4. `.proto` 合同设计已冻结，字段编号、枚举、`RPC -> HTTP` 映射已冻结，`buf` 校验链要求已明确；正式落地为仓库唯一合同源由 `phase03-11` 承接
5. 空状态与冷启动路径已明确，不依赖手工 `SQL`
6. 前端实现设计层结果（页面文件落点、路由树、组件树、状态模型、布局降级）已冻结到可直接进入实现的深度
7. 后端实现设计层结果（模块边界、文件落点、分层语义、candidate 接线、支撑文件落点）已冻结到可直接进入实现的深度
8. 验收基线（联调环境、重置脚本、基线种子、异常路径、冷启动路径）已冻结
9. 非目标矩阵已明确
10. 本文档与 `mvp_spec_v0.1.md`、`module_registry_spec_v0.1.md`、`phase02-12` 验收结论及根级真相源互链一致，无第二套边界

---

## 14. 与根级真相源互链

- 项目定位与入口：以 `AGENTS.md` 为准
- 阶段路线：以 `plan.md` 为准
- 技术栈标准：以 `TECH_STACK_BASELINE.md` 为准
- 规则门禁：以 `project_rules.md` 为准
- 目录与迁移落点：以 `architecture_map.md` 为准
- 最终共识：以 `PSCO-summarize-feedback.md` 为准
- 当前阶段唯一执行层上游：`mvp_spec_v0.1.md`
- `Module Registry` 已交付边界上游：`module_registry_spec_v0.1.md` + `phase02-12` 验收结论
- 当前阶段架构规划：`phase03_decision_center_foundation_architecture_plan.md`
- 当前阶段共享基线：`phase03_decision_center_foundation_shared_baseline.md`
- 当前阶段开发计划：`phase03_decision_center_foundation_dev_plan.md`

> 本文档不重写上述根级真相源与阶段文档的主结论，仅作为 `Decision Center` 当前阶段的唯一规格收敛入口。前置子规格 `phase03-01 ~ 09` 仅作为本文档的冻结来源与追踪依据，不再作为后续实现与验收的长期并列入口。
