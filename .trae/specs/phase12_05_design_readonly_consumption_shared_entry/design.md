# phase12-05 只读消费深化与共享入口设计

> 本设计产出作为 `phase12-07`（后端合同、导出结果与共享只读视图设计）与 `phase12-06`（前端读路径 owner、共享摘要与回流设计）的正式输入前提。
> 所有结论均继承 `phase12-03` 已冻结的只读消费边界，并以当前真实源码合同与读路径实现为准，不预支不存在的字段或 resolver。

---

## 一、影响对象清单

### 1.1 后端合同与服务对象

| # | 对象 | 分类 | 理由 |
|---|---|---|---|
| B1 | `ProjectContextService.GetProjectContext` | follow-regression | 继续作为唯一结构化 canonical owner，不在 `phase12-05` 改写身份，但必须被真实字段能力重新盘点 |
| B2 | `ProjectContextService.ExportProjectContext` | follow-regression | 继续作为 agent-facing Markdown 导出 owner，不在 `phase12-05` 改写身份，但必须与 L1 单向派生关系继续冻结 |
| B3 | `proto/psco/project_context/v1/project_context.proto` | must-change | `phase12-05` 必须把哪些字段已存在、哪些只能作为 `phase12-07` 候选说清，不能再把不存在字段当作已覆盖 |

### 1.2 前端共享入口与读路径对象

| # | 对象 | 分类 | 理由 |
|---|---|---|---|
| F1 | `frontend/src/features/project-context/` | must-change | 已被冻结为唯一允许的新 Web 跨切片共享只读 owner，`phase12-05` 必须明确它未来承接什么、不承接什么 |
| F2 | `frontend/src/features/repository-binding/data/use-repository-detail-read.ts` | follow-regression | 直接 repository-scoped 页面可直接承接 `repository_id`，本读路径只需保持与 L1/L3 关系一致 |
| F3 | `frontend/src/features/product-registry/data/use-product-detail-read.ts` | must-change | 需要冻结 Product 详情页如何从现有 detail read 回到 `repository_id` 的正式规则 |
| F4 | `frontend/src/features/module-registry/data/use-module-detail-read.ts` | must-change | 需要冻结 Module 详情页如何从现有 detail read 回到 `repository_id` 的正式规则 |
| F5 | `frontend/src/features/decision-center/data/use-decision-detail-read.ts` | must-change | 当前不直接返回 `repository_id`，`phase12-05` 必须把 Decision 路径的正式 resolver 方案与候选补充条件写清 |
| F6 | `frontend/src/features/dashboard/data/use-dashboard-overview-read.ts` | follow-regression | 衍生消费页继续消费受控派生摘要或共享语义，不直接承接结构化锚点解析 |
| F7 | `frontend/src/features/onboarding/data/use-onboarding-read.ts` | follow-regression | 同上 |
| F8 | `frontend/src/features/review/data/use-daily-review-read.ts` | follow-regression | 同上 |
| F9 | `frontend/src/features/review/data/use-weekly-review-read.ts` | follow-regression | 同上 |

### 1.3 文档与设计输入对象

| # | 对象 | 分类 | 理由 |
|---|---|---|---|
| D1 | `phase12-05 spec.md` | must-change | 必须补齐“不得预判 phase12-07 无新增”的约束 |
| D2 | `phase12-05 tasks.md` | must-change | 必须把“真实已覆盖 / L3 单向派生 / phase12-07 候选”拆成可验收任务 |
| D3 | `phase12-05 checklist.md` | must-change | 必须把 resolver 歧义、字段真实性与模板完整性纳入显式验收 |

---

## 二、结论矩阵

### 2.1 共享只读 owner 结论

| 对象 | 当前承接 | 需要改成什么 | 为什么需要改 |
|---|---|---|---|
| `GetProjectContext` | 结构化 `repository_id` 驱动聚合读取 | 继续保持 L1 canonical owner，不额外承接页面私有 shape | 与 `phase12-03`、当前 proto 合同一致，避免长出第二套事实源 |
| `ExportProjectContext` | Markdown 导出入口 | 继续保持 L2 单向派生 owner，不承担 Web 专属字段真相源 | 避免 agent/Web 双轨语义 |
| `frontend/src/features/project-context/` | 仅在文档层被冻结为唯一共享入口 | 明确未来只承接受控 adapter、query options、共享语义来源与入口定位视图 | 避免“共享入口”沦为第二个页面 data 层或第二 canonical |
| 切片内 `pages/components/data` | 当前各自消费本切片读结果 | 明确只负责本地渲染与页面回流，不得自造跨切片共享摘要 | 与“切片内展示 owner”身份对齐 |

### 2.2 三类页面锚点结论

| 页面类别 | 当前承接 | 需要改成什么 | 为什么需要改 |
|---|---|---|---|
| 直接 `repository-scoped` 页面 | 可直接拿到路由 `repositoryId` | 继续直接用 `repository_id` 进入 L1，再由 L3/L4 消费 | 这是唯一无须解析回锚点的正式路径 |
| 间接 `repository-scoped` Product 页面 | 已有 Product detail read，且 `bound_repositories[*].repository_id` 可见 | 冻结为“先从 detail read 提取候选 `repository_id`，仅在唯一候选成立时回到共享主线” | 既避免临场猜锚点，也避免擅自选一个 repository |
| 间接 `repository-scoped` Module 页面 | 已有 Module detail read，且 `repository_mappings[*].repository_id` 可见 | 冻结为“先从 detail read 提取候选 `repository_id`，仅在唯一候选成立时回到共享主线” | 与 Product 同理 |
| 间接 `repository-scoped` Decision 页面 | 当前 detail read 只有 `linked_modules` 与 `source_context`，没有直接 `repository_id` | 冻结为“经 linked module/source module 回到 Module detail，再按唯一候选规则解析 `repository_id`；若仍无法唯一确定，显式进入 `phase12-07` 候选” | 不能再假设“既有 detail read 已覆盖” |
| 衍生消费页 | 当前不依赖固定 `repository_id` | 继续只消费 L3 共享语义来源、固定入口链接或既有派生摘要，不得升格为锚点入口 | 保持读路径只读与单一锚点规则 |

### 2.3 phase12-07 输入结论

| 输入类别 | 当前状态 | 正式结论 | 为什么这样收口 |
|---|---|---|---|
| 四实体最小共享摘要 | 当前 proto 已有按实体拆分的摘要对象 | 先按真实已有字段复用；跨实体统一 label 由 L3 单向派生；不存在的通用字段不得假装已覆盖 | 避免把 adapter 字段误写成 L1 字段 |
| 规则/约束/文档入口定位 | 当前 proto 已有 `entry_ref / entry_kind`，以及 `label / summary` | 继续复用真实已有定位字段；如后续确需 source-entity 级回溯，再作为 `phase12-07` 候选 | 真实合同只到这一步 |
| resolver | Product / Module 可见 repository 候选，Decision 仍需链式解析 | 明确 Product / Module / Decision 三条不同 resolver 规则；Decision 不再被笼统写成“已由 detail read 覆盖” | 这是本轮最核心的机械冻结点 |

---

## 三、承接位矩阵

### 3.1 共享只读 owner 分层

| 层级 | Owner | 输入 | 输出 | 最终承接位 | 分类 |
|---|---|---|---|---|---|
| L1 | `ProjectContextService.GetProjectContext` | `repository_id` | `RepositorySummary / ProductSummary / ModuleSummary[] / DecisionSummary[] / RuleEntry[] / PhaseEntry[] / BoundaryEntry[]` | Go backend canonical owner | follow-regression |
| L2 | `ProjectContextService.ExportProjectContext` | `repository_id` | Markdown 导出结果 | Go backend Markdown owner | follow-regression |
| L3 | `frontend/src/features/project-context/` | L1 结构化结果 + 冻结语义来源 | 共享 query options、受控 adapter、共享语义来源、入口定位视图 | Web 跨切片共享只读 owner | must-change |
| L4 | 各 feature `pages/components/data` | L3 共享结果或切片本地只读结果 | 页面级展示、回流与局部空态 | 切片内展示 owner | follow-regression |

### 3.2 三类页面承接矩阵

| 页面类别 | 输入 | 解析链 | 复用方式 | 失败语义 | 最终承接位 |
|---|---|---|---|---|---|
| `repositories/$repositoryId` | 路由 `repositoryId` | 无需解析，直接映射为 `repository_id` | `repository_id -> L1 -> L3 -> L4` | 无效 ID → 404；L1 失败 → 页面错误态；结果为空 → 合法空态 | Repository detail 切片通过 L3 消费 |
| `products/$productId` | 路由 `productId` | `use-product-detail-read` 提取 `bound_repositories[*].repository_id`，仅当唯一候选成立时回到 L1 | `product detail -> repository candidate -> L1 -> L3 -> L4` | 无候选或多候选 → 不展示共享摘要并记录为 resolver 未满足；L1 失败 → 局部错误态 | Product detail 切片通过 L3 消费 |
| `modules/$moduleId` | 路由 `moduleId` | `use-module-detail-read` 提取 `repository_mappings[*].repository_id`，仅当唯一候选成立时回到 L1 | `module detail -> repository candidate -> L1 -> L3 -> L4` | 无候选或多候选 → 不展示共享摘要并记录为 resolver 未满足；L1 失败 → 局部错误态 | Module detail 切片通过 L3 消费 |
| `decisions/$decisionId` | 路由 `decisionId` | `use-decision-detail-read` 提取 `source_module_id` 与 `linked_modules[*].module_id`，再经 Module detail 提取 `repository_mappings[*].repository_id`；仅当唯一候选成立时回到 L1 | `decision detail -> module candidate(s) -> module detail -> repository candidate -> L1 -> L3 -> L4` | 无 module 候选、module 链失败、repository 候选不唯一 → 不展示共享摘要，并明确列入 `phase12-07` 候选 | Decision detail 切片通过 L3 消费 |
| `dashboard / onboarding / reviews/*` | 页面既有只读结果 + L3 共享语义来源 | 不解析 `repository_id` | 只消费共享语义来源、入口定位链接或受控派生摘要 | L3 不可用 → 局部降级，不阻断主体页面 | 各衍生页切片继续本地渲染 |

### 3.3 resolver 候选唯一性规则

| resolver | 唯一候选成立条件 | 不成立时怎么处理 | 是否已被 phase12-05 判定为必须新增后端字段 |
|---|---|---|---|
| Product → repository | `bound_repositories` 去重后恰好 1 个 `repository_id` | 不展示共享摘要；进入 `phase12-07` 候选评估，但不在 `phase12-05` 直接宣布新增 | 否 |
| Module → repository | `repository_mappings` 去重后恰好 1 个 `repository_id` | 不展示共享摘要；进入 `phase12-07` 候选评估，但不在 `phase12-05` 直接宣布新增 | 否 |
| Decision → repository | 由 `source_module_id` / `linked_modules` 回到 Module detail 后，去重后恰好 1 个 `repository_id` | 不展示共享摘要；显式列为 `phase12-07` 的优先 resolver 候选 | 否，但需进入候选评估 |

---

## 四、共享语义来源 vs 切片内渲染矩阵

| 内容 | 应收敛到哪里 | 继续保留在哪里 | 原因 |
|---|---|---|---|
| 四实体冻结语义单行解释 | `frontend/src/features/project-context/` | 各 detail / dashboard / onboarding / review 页面本地渲染 | 这是高频共享语义来源，不等于共享页面结构 |
| `repository_id` 驱动的共享 query options | `frontend/src/features/project-context/` | 切片内页面调用结果 | query options 属于 L3，不属于页面私有 data 逻辑 |
| `entry_ref / entry_kind` 等入口定位视图 | `frontend/src/features/project-context/` | 各页面本地决定如何展示入口链接 | 定位数据可共享，渲染结构不共享 |
| Product / Module / Decision 各自的详情页布局与空态 | 不共享 | 各自切片内继续保留 | 页面结构是切片专属，不应被 `project-context/` 吞并 |
| Dashboard / Onboarding / Review 的说明文案布局 | 不共享 | 各自切片内继续保留 | 这些页面只消费共享语义来源，不承担结构化锚点职责 |

---

## 五、Before / After 样例

### 5.1 phase12-07 输入判断

| Before | After |
|---|---|
| “当前既有 `GetProjectContext` 已覆盖四实体最小共享摘要需求，无需 phase12-07 新增字段。” | “当前真实 proto 已覆盖 Repository / Product / Module / Decision 的基础摘要字段；跨实体统一 label 由 L3 单向派生；若后续需要计数、source-entity 回溯或稳定 resolver 字段，再进入 `phase12-07` 候选评估。” |
| “规则/约束/文档入口定位需求已覆盖，无需 phase12-07 新增字段。” | “当前已确认可复用 `entry_ref / entry_kind / label / summary`；若后续需要 source-entity 级定位回溯，再由 `phase12-07` 评估是否增加受控派生字段。” |

### 5.2 间接 repository-scoped 页面解析

| Before | After |
|---|---|
| “间接页面先通过既有 detail read 解析关联的 `repository_id`。” | “Product 与 Module 页面按现有 detail read 的唯一候选规则解析；Decision 页面必须先经 module 链回到 repository 候选，若不唯一则不得硬猜，并显式进入 `phase12-07` resolver 候选。” |
| “若 detail read 未返回 `repository_id`，则该页面不提供共享摘要，不为此新增 RPC 字段。” | “若当前链路无法唯一回到 `repository_id`，页面不得硬猜；是否需要新增受控派生读取或字段，由 `phase12-07` 依据重复解释回收收益决定。” |

---

## 六、供 phase12-07 继续承接的正式输入需求

### 6.1 已确认可直接复用的真实结构化字段

| 需求 | 当前真实来源 | phase12-07 是否必须新增 |
|---|---|---|
| Repository 最小摘要 | `repository.id / name / provider / url / description` | 否 |
| Product 最小摘要 | `product.id / name / description / status` | 否 |
| Module 最小摘要 | `modules[].id / name / description / status` | 否 |
| Decision 最小摘要 | `decisions[].id / title / status / context / hit_sources / created_at` | 否 |
| 规则入口定位 | `rules[].label / summary / entry_ref / entry_kind` | 否 |
| phase 文档入口定位 | `phases[].label / status_summary / entry_ref / entry_kind` | 否 |
| boundary 摘要 | `boundaries[].label / summary` | 否 |

### 6.2 可由 L3 单向派生、但不属于 L1 真实字段的共享结果

| 共享结果 | 派生来源 | 是否需要 phase12-07 |
|---|---|---|
| 四实体统一语义 label | L3 共享语义来源 + L1 各实体摘要 | 否 |
| 跨实体统一的 `entity_type / entity_name` 展示 adapter | L3 adapter 由各实体真实字段单向派生 | 否 |
| Web 入口定位 view model | L3 将 `label / summary / entry_ref / entry_kind` 裁剪为页面可消费视图 | 否 |

### 6.3 明确进入 phase12-07 候选评估的字段 / resolver

| 候选项 | 当前缺口 | 进入条件 | 当前阶段结论 |
|---|---|---|---|
| Product 页稳定 `resolved_repository_id` | 当前只在唯一候选成立时可解析；多候选仍无 canonical 规则 | 若 `products/$productId` 必须稳定接入共享摘要，且多候选成为重复痛点 | 列为候选，不预判新增 |
| Module 页稳定 `resolved_repository_id` | 当前只在唯一候选成立时可解析；多候选仍无 canonical 规则 | 若 `modules/$moduleId` 必须稳定接入共享摘要，且多候选成为重复痛点 | 列为候选，不预判新增 |
| Decision 页稳定 `resolved_repository_id` / 受控派生 resolver | 当前需经 module 链回到 repository，且没有直接字段 | 若 `decisions/$decisionId` 需要稳定接入共享摘要，且链式解析无法唯一落定 | 列为优先候选，不预判新增 |
| `source_entity_type / source_entity_id` 级入口回溯 | 当前 `project_context.proto` 无此类字段 | 若 Web / agent 同时需要稳定回溯某条规则/入口来自哪个实体 | 列为候选，不预判新增 |
| 聚合计数字段（如复用计数 / 决策计数 / 绑定状态） | 当前 `GetProjectContext` 无这些统一字段 | 若 Web / agent 均出现稳定共享消费且能回收重复解释逻辑 | 列为候选，不预判新增 |

### 6.4 phase12-07 必须承接的判断题

`phase12-07` 必须基于本设计继续回答：

1. Product / Module / Decision 页面是否需要一个受控派生 resolver 或受控派生字段，来稳定回到唯一 `repository_id`
2. 若需要新增，是否仍可保持 `ProjectContextService` 下的受控派生读取身份
3. 该新增到底回收了哪段前端 / agent 重复解释逻辑
4. 新增后如何继续保留 `repository_id` 与 `entry_ref / entry_kind` 的定位能力

---

## 七、明确不做清单

1. 不在 `phase12-05` 直接修改 `.proto`、Connect 接口、service 实现或前端读路径代码
2. 不在 `phase12-05` 直接宣布“phase12-07 无需新增任何字段 / resolver”，除非已有真实源码合同证明
3. 不把 `frontend/src/features/project-context/` 扩写成写路径 owner、页面私有状态 owner 或第二事实源
4. 不允许页面通过工作区扫描、手工查询、额外脚本或临场猜测补齐 `repository_id`
5. 不在 `phase12-05` 把 Dashboard / Onboarding / Review 升格为新的结构化锚点入口

---

## 八、设计推进顺序与一致性声明

### 8.1 设计推进顺序

```text
phase12-05（本设计）
  ├─ 冻结：共享 owner 矩阵、三类页面承接规则、真实字段复用边界、phase12-07 候选输入
  ├─→ phase12-07：判断哪些候选需要进入受控派生读取或字段补充
  └─→ phase12-06：在 phase12-05 / 07 已冻结结果上设计前端读路径 owner、adapter、query options 与 reread 关系
```

### 8.2 与上游文档的一致性声明

| 上游文档 | 关键冻结内容 | 本设计对齐 |
|---|---|---|
| `architecture_plan §4.4A` | 共享只读 owner 分层（L1-L4） | §2.1 / §3.1 |
| `architecture_plan §4.6A` | 三类页面承接矩阵 | §2.2 / §3.2 |
| `architecture_plan §4.6B` | `05 -> 07 -> 06` 顺序 | §8.1 |
| `shared_baseline §3.4` | `repository_id` 为唯一结构化输入锚点、共享只读 owner 与消费边界 | §2-§4 |
| `phase12-03 spec` | 共享只读 owner 分层、三类页面、复用判定、更重通道进入条件 | §2-§7 全部继承，不越权改写 |
