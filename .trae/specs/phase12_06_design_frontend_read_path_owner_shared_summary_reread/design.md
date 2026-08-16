# phase12-06 前端读路径 owner、共享摘要与回流设计

> 本设计消费 `phase12-05`（共享 owner / 三类页面矩阵）与 `phase12-07`（后端合同：无需新增）的冻结结论，产出前端可执行的读路径 owner 审计、共享摘要承接位、缓存/reread/回流关系设计。
> 不修改 `phase12-05/07` 已冻结结论，不提前实现代码。

---

## 一、前端 read owner 影响对象清单与分类矩阵

### 1.1 既有 read hook 审计（8 个）

| # | 文件 | 输入锚点 | 缓存键 | 输出 shape | 是否含四实体语义解释 | 是否含共享摘要拼装 | 分类 | 理由 |
|---|---|---|---|---|---|---|---|---|
| R1 | `use-product-detail-read.ts` | `productId` | `['product-detail', productId]` | `ProductDetail`（含 `product`、`bound_modules`、`bound_repositories`） | 否（纯数据映射） | 否 | **no-change** | 纯切片内 detail read，输入/output/shape 不变 |
| R2 | `use-repository-detail-read.ts` | `repositoryId` | `['repository-detail', repositoryId]` | `RepositoryDetail`（含 `repository`、`bound_products`、`mapped_modules`） | 否（纯数据映射） | 否 | **no-change** | 纯切片内 detail read，不变 |
| R3 | `use-module-detail-read.ts` | `moduleId` | `['module-detail', moduleId]` | `ModuleDetail`（含 `module`、`releases`、`product_bindings`、`repository_mappings`、`decision_links`） | 否（纯数据映射） | 否 | **no-change** | 纯切片内 detail read，不变 |
| R4 | `use-decision-detail-read.ts` | `decisionId` | `['decision-detail', decisionId]` | `DecisionDetail`（含 `decision`、`linked_modules`、`source_context`） | 否（纯数据映射） | 否 | **no-change** | 纯切片内 detail read，不变 |
| R5 | `use-dashboard-overview-read.ts` | 无（全量聚合） | `['dashboard-overview']` | `DashboardOverview`（计数聚合） | 否（纯聚合） | 否 | **no-change** | 纯聚合读取，不定义实体语义 |
| R6 | `use-onboarding-read.ts` | 无（全局状态） | `['onboarding-state']` | `OnboardingFullReadResult`（含 `first_run_state`、`chain_state`） | 否（纯状态聚合） | 否 | **no-change** | 纯状态读取，不定义实体语义 |
| R7 | `use-daily-review-read.ts` | 无 | `DAILY_REVIEW_QUERY_KEY` | `DailyReviewReadModel`（含 `currentFocusSignals`、`representativeSignals`、`pendingDecisions`） | 否（纯信号映射） | 否 | **no-change** | 纯 review 信号读取，不定义实体语义 |
| R8 | `use-weekly-review-read.ts` | `returnCandidateId?` | `WEEKLY_REVIEW_QUERY_KEY` | `WeeklyReviewReadModel`（含 `overview`、`recentActivities`、`representativeSignals`、`moduleReuseSummary`、`capabilitySummary`） | 否（纯信号映射） | 否（模板候选为既有数据） | **no-change** | 纯 review 信号读取，不定义实体语义 |

### 1.2 新增 L3 共享只读 owner（`project-context/data/*`）

| # | 文件（待创建） | 输入锚点 | 输出 shape | 分类 | 理由 |
|---|---|---|---|---|---|
| L3-1 | `shared-semantic-constants.ts` | 无 | 四实体冻结语义标签、实体类型-标签映射、共享入口描述常量 | **must-change** | 唯一跨切片共享语义常量源，只回收稳定跨页复用的语义标签与入口描述，不接管切片内 next-action / CTA 文案 |
| L3-2 | `use-project-context-read.ts` | `repositoryId` | TanStack Query `useQuery` 结果，封装 `GetProjectContext` RPC 调用 | **must-change** | 唯一跨切片共享 GetProjectContext query options，避免各页面直接 import transport |
| L3-3 | `entry-location-view-model.ts` | `RuleEntry[]` / `PhaseEntry[]` | `{entryRef, entryKind, title, summary}[]` | **must-change** | 为 Repository / Product / Module / Decision detail 页后续接入共享入口定位时提供唯一 adapter 落点；本阶段不把 Dashboard / Review 误列为现成 consumer |

### 1.3 Summary 分类

| 分类 | 数量 | 对象 |
|---|---|---|
| must-change | 3 | L3-1、L3-2、L3-3（新增 `project-context/data/*`） |
| no-change | 8 | R1-R8（8 个既有 read hook 不变） |
| follow-regression | 0 | — |

---

## 二、query 层、L3 共享只读 owner 与页面层的分工矩阵

### 2.1 三层职责边界

```
┌─────────────────────────────────────────────────────────────┐
│ 页面 / 组件渲染层（pages/、components/）                      │
│   - 消费切片 read owner 或 L3 共享只读结果                     │
│   - 不得拼装第二套共享摘要合同                                 │
│   - 语义 label 从 L3 常量获取，不再硬编码                       │
└──────────────┬──────────────────────────┬───────────────────┘
               │                          │
    ┌──────────▼──────────┐    ┌──────────▼──────────────────┐
    │ 切片 read owner      │    │ L3 project-context/data/*   │
    │ (use-*-read.ts)      │    │                             │
    │                      │    │ shared-semantic-constants   │
    │ - 原始读取            │    │   - 四实体冻结语义标签       │
    │ - 缓存键              │    │   - 实体类型映射             │
    │ - 响应解包            │    │   - 入口描述常量             │
    │ - 页面级 read model   │    │                             │
    │                      │    │ use-project-context-read    │
    │ - 页面专属字段        │    │   - GetProjectContext 封装   │
    │   (如 Product 绑定    │    │   - 共享 query options       │
    │    仓库、Module 仓库   │    │   - 缓存键                   │
    │    映射、Decision     │    │                             │
    │    linked_modules)    │    │ entry-location-view-model   │
    │                      │    │   - entry_ref/kind 裁剪      │
    └──────────────────────┘    │   - 统一入口定位视图          │
                                └─────────────────────────────┘
```

### 2.2 各层的承接项

| 承接项 | 归属 | 理由 |
|---|---|---|
| 四实体冻结语义标签（"经营目标与交付容器" 等） | L3 `shared-semantic-constants.ts` | 跨 4+ 页面/切片/组件稳定复用，回收重复解释 |
| 实体类型-标签映射（`{Product: "经营目标与交付容器", ...}`） | L3 `shared-semantic-constants.ts` | 同上，为入口定位视图提供统一映射 |
| `GetProjectContext` query options | L3 `use-project-context-read.ts` | 跨切片共享 query 封装，避免各页面直接创建 transport |
| `entry_ref / entry_kind / label / summary` 裁剪 | L3 `entry-location-view-model.ts` | 冻结 detail 页共享入口定位的唯一 adapter 落点，避免未来在各 detail 页面各自裁剪一套 |
| 实体 detail 读取（Product/Module/Repository/Decision） | 切片 read owner（R1-R4） | 页面专属字段与读取逻辑，不跨切片 |
| Dashboard 聚合读取 | 切片 read owner（R5） | 页面专属聚合，不跨切片 |
| Onboarding 状态读取 | 切片 read owner（R6） | 页面专属读取，不跨切片 |
| Daily/Weekly Review 信号读取 | 切片 read owner（R7-R8） | 页面专属读取，不跨切片 |
| 页面专属字段（如 binding 详情、linked modules） | 切片 read owner | 不升格为 L3 共享合同 |

### 2.3 页面专属字段不进入 L3 的清单

| 字段 | 归属切片 read owner | 不进入 L3 的理由 |
|---|---|---|
| `ProductDetail.bound_modules` | R1 `use-product-detail-read` | 仅 Product Detail 页消费 |
| `ProductDetail.bound_repositories` | R1 `use-product-detail-read` | 仅 Product Detail 页消费 |
| `RepositoryDetail.bound_products` | R2 `use-repository-detail-read` | 仅 Repository Detail 页消费 |
| `RepositoryDetail.mapped_modules` | R2 `use-repository-detail-read` | 仅 Repository Detail 页消费 |
| `ModuleDetail.releases` | R3 `use-module-detail-read` | 仅 Module Detail 页消费 |
| `ModuleDetail.product_bindings` | R3 `use-module-detail-read` | 仅 Module Detail 页消费 |
| `DecisionDetail.linked_modules` | R4 `use-decision-detail-read` | 仅 Decision Detail 页消费 |
| `DecisionDetail.source_context` | R4 `use-decision-detail-read` | 仅 Decision Detail 页消费 |

---

## 三、页面读取、缓存、成功回流与 reread 关系

### 3.1 直接 repository-scoped 页面：`repositories/$repositoryId`

```
RepositoryDetailPage
  │
  ├─ R2 use-repository-detail-read(repositoryId)
  │     └─ 缓存键: ['repository-detail', repositoryId]
  │     └─ 失败: 页面级错误态
  │
  └─ L3-2 use-project-context-read(repositoryId)  [可选]
        └─ 缓存键: ['project-context', repositoryId]
        └─ 失败: 共享摘要区局部错误态，不阻断页面
        └─ 条件: 页面需要展示共享摘要时才启用
```

**reread**: 无 — 共享语义常量是静态的；`GetProjectContext` 在 Repository 变更后由 mutation owner 主动 invalidate `['project-context', repositoryId]`

### 3.2 间接 repository-scoped 页面：`products/$productId`、`modules/$moduleId`、`decisions/$decisionId`

```
ProductDetailPage / ModuleDetailPage / DecisionDetailPage
  │
  ├─ R1/R3/R4 use-*-detail-read(entityId)
  │     └─ 缓存键: ['*-detail', entityId]
  │     └─ 失败: 页面级错误态
  │
  └─ L3-2 use-project-context-read(repositoryId)  [可选]
        └─ repositoryId 从 detail read 结果中解析
        └─ 若无法解析: 不展示共享摘要，不阻断页面
        └─ 缓存键: ['project-context', repositoryId]
```

**reread**:
- 间接页面若在成功回调时**能够拿到唯一的 `repositoryId`**，则其 mutation owner 必须同时 invalidate `['project-context', repositoryId]`
- 间接页面若**拿不到唯一 `repositoryId`**，则不得把共享摘要写成依赖成功回流即时刷新的核心信息；此时继续依赖 detail read reread，或在重新进入页面时按 `staleTime` 重新获取
- 不允许为了图省事直接 invalidate 全量 `['project-context']`

### 3.3 衍生消费页：`dashboard`、`onboarding`、`reviews/daily`、`reviews/weekly`

```
DashboardHomePage / OnboardingPage / DailyReviewPage / WeeklyReviewPage
  │
  ├─ R5/R6/R7/R8 各自的 read hook
  │     └─ 各自的缓存键
  │     └─ 失败: 页面级或局部错误态
  │
  └─ L3-1 shared-semantic-constants（静态导入）
        └─ 不涉及 query/reread
        └─ 用于四实体语义 label、入口描述等
```

**reread**: 衍生页不调用 `GetProjectContext`，不涉及 `['project-context', *]` 缓存

### 3.4 写路径成功后的 reread 策略

| 写操作 | 影响的 read owner | 失效动作 | 触发方 |
|---|---|---|---|
| `Product` 创建/更新 | R1 `['product-detail', productId]` | invalidate | 切片 mutation owner |
| `Repository` 创建/更新 | R2 `['repository-detail', repositoryId]` | invalidate | 切片 mutation owner |
| `Module` 创建/更新 | R3 `['module-detail', moduleId]` | invalidate | 切片 mutation owner |
| `Decision` 创建/更新/状态变更 | R4 `['decision-detail', decisionId]` | invalidate | 切片 mutation owner |
| `Binding` 变更（影响 repository 关联） | R2 `['repository-detail', repositoryId]` + `['project-context', repositoryId]` | invalidate | 切片 mutation owner |
| `Product / Module / Decision` 写操作，且成功回调可拿到唯一 `repositoryId` | 对应 detail query + `['project-context', repositoryId]` + 既有 dashboard/review query | invalidate | 切片 mutation owner |
| `Product / Module / Decision` 写操作，但成功回调拿不到唯一 `repositoryId` | 对应 detail query + 既有 dashboard/review query | invalidate | 切片 mutation owner |
| `Feedback` 提交 | R5 `['dashboard-overview']` + R7/R8 | invalidate | 切片 mutation owner |
| `Review` 完成 | R5 + R7/R8 | invalidate | 切片 mutation owner |

**关键规则**：
- 谁能在成功回调中拿到**唯一 `repositoryId`**，谁就负责精确 invalidate `['project-context', repositoryId]`
- 若 mutation owner 无法拿到唯一 `repositoryId`，则不得伪造 project-context 新鲜度闭环，也不得全量扫失效
- mutation owner 只负责自己正式承接的 reread 集，不由页面或组件临场追加第二套 query 失效
- 页面不各自随意 invalidate 一批 query

### 3.5 初次加载 / 局部重试 / 整页重试 / 成功回流 区分

| 场景 | 行为 | 触发条件 |
|---|---|---|
| 初次加载 | 按页面执行全部 query（detail read + 可选共享摘要） | 页面首次渲染 |
| 局部重试 | 只重试失败的 query，不连带已成功的 query | 单个 query 失败后用户点击重试 |
| 整页重试 | 同时重试所有 query | 页面级 `page-error` 状态 |
| 成功回流 | 写操作成功后，mutation owner invalidate 相关 query，TanStack Query 自动重读 | 写操作成功回调 |

---

## 四、散装解释逻辑回收清单

### 4.1 四实体语义标签（5 处散装 → 1 处共享）

| 散装位置 | 当前表达 | 回收目标 | 回收方式 |
|---|---|---|---|
| `onboarding-page.tsx` WelcomeStep | "创建一个产品（Product）" 等 | L3-1 `PRODUCT_SEMANTIC_LABEL` | 导入常量替换硬编码字符串 |
| `product-summary-card.tsx` | 无语义标签 | L3-1 `PRODUCT_SEMANTIC_LABEL` | 在标题区引入常量（phase12-08 实现） |
| `repository-summary-card.tsx` | 无语义标签 | L3-1 `REPOSITORY_SEMANTIC_LABEL` | 同上 |
| `module-summary-card.tsx` | 无语义标签 | L3-1 `MODULE_SEMANTIC_LABEL` | 同上 |
| `decision-detail-summary-card.tsx` | 无语义框架 | L3-1 `DECISION_SEMANTIC_LABEL` | 同上 |

### 4.2 入口定位视图 adapter（先冻结唯一落点，不假装当前已有 3 处散装）

| 候选消费者 | 当前状态 | 统一承接位 | 使用规则 |
|---|---|---|---|
| `repository-binding-detail-page.tsx` 共享上下文区 | 当前尚未消费 `RuleEntry[] / PhaseEntry[]` | L3-3 `entry-location-view-model.ts` | 若后续展示规则 / phase 入口链接，必须经该 adapter 裁剪 |
| `product-detail-page.tsx` / `module-detail-page.tsx` / `decision-detail-page.tsx` 共享摘要区 | 当前尚未消费入口定位字段 | L3-3 | 仅在页面已通过唯一 `repositoryId` 接入 `use-project-context-read` 后使用 |
| `dashboard / onboarding / review` | 当前没有 `GetProjectContext` 输入来源，也不直接消费入口定位字段 | 不接入 L3-3 | 本阶段明确不列为 consumer，避免假覆盖 |

### 4.3 明确不回收的散装逻辑

| 逻辑 | 位置 | 不回收的理由 |
|---|---|---|
| `ModuleNextActionBar` 动作描述（"绑定产品"/"映射仓库"） | `module-next-action-bar.tsx` | 仅 Module Detail 页消费，属切片内专属 UI 文案 |
| `DecisionDetailSummaryCard` 状态推进 CTA | `decision-detail-summary-card.tsx` | 仅 Decision Detail 页消费，属切片内专属 UI |
| `DashboardPrimaryActionPanel` review 入口 | `dashboard-primary-action-panel.tsx` | 仅 Dashboard 页消费，属切片内专属 UI |
| `ReviewActionFooter` 完成 Review | `review-action-footer.tsx` | 仅 Review 页消费，属切片内专属 UI |
| `OnboardingCtaButton` 开始首轮录入 | `onboarding-cta-button.tsx` | 仅 Dashboard 页消费，属切片内专属 UI |

### 4.4 回收清单与 phase12-04 的一致性

本回收清单与 `phase12-04` 的 surface 承接矩阵一致：
- 摘要卡片 4 个全部 `must-change` → 对应 §4.1 四行回收
- 说明文案 Onboarding WelcomeStep `must-change` → 对应 §4.1 第一行回收
- 下一步动作说明 `ModuleNextActionBar` 四个场景 `must-change` → 对应 §4.3 不回收行（属切片内专属）
- 入口定位 `no-change` → 对应 §4.2：当前仅冻结唯一 adapter 落点，不把 Dashboard / Review 误写成既有 consumer

---

## 五、`project-context/data/*` 最小文件边界

### 5.1 目录结构

```
frontend/src/features/project-context/
├── data/
│   ├── shared-semantic-constants.ts    # 四实体冻结语义标签、实体类型映射、入口描述常量
│   ├── use-project-context-read.ts     # TanStack Query 封装 GetProjectContext
│   └── entry-location-view-model.ts    # entry_ref/entry_kind/label/summary 裁剪 adapter
├── index.ts                            # barrel export
```

### 5.2 各文件最小设计

#### `shared-semantic-constants.ts`

```ts
// 四实体冻结语义标签（phase12-02 冻结）
export const PRODUCT_SEMANTIC_LABEL = "经营目标与交付容器"
export const REPOSITORY_SEMANTIC_LABEL = "代码仓库身份对象与项目锚点"
export const MODULE_SEMANTIC_LABEL = "可复用能力资产"
export const DECISION_SEMANTIC_LABEL = "规则、约束、选择与依据的索引对象"

// 实体类型-标签映射
export const ENTITY_SEMANTIC_LABEL_MAP: Record<string, string> = {
  Product: PRODUCT_SEMANTIC_LABEL,
  Repository: REPOSITORY_SEMANTIC_LABEL,
  Module: MODULE_SEMANTIC_LABEL,
  Decision: DECISION_SEMANTIC_LABEL,
}

// 入口描述常量
export const RULE_ENTRY_LABEL = "规则与约束"
export const PHASE_ENTRY_LABEL = "当前阶段"
export const BOUNDARY_ENTRY_LABEL = "当前阶段边界"
```

**启用条件**：已有 4+ 页面/切片/组件稳定复用（Onboarding + 4 个 SummaryCard），符合 phase12-03 的 3+ 门槛。

**禁止事项**：不承接写路径、页面私有状态、并列 canonical 字段语义。

#### `use-project-context-read.ts`

- 输入：`repositoryId: string`
- 输出：`UseQueryResult<ProjectContext>`（基于 `GetProjectContext` RPC 的 TanStack Query 封装）
- 缓存键：`['project-context', repositoryId]`
- 失败语义：以 `UseQueryResult.error` 暴露，不回退到页面级错误
- 使用方式：各页面可选启用，不强制

#### `entry-location-view-model.ts`

- 输入：`RuleEntry[] | PhaseEntry[]`
- 输出：`{entryRef: string, entryKind: string, title: string, summary: string}[]`
- 职责：将 `entry_ref/entry_kind/label/summary` 统一裁剪为 detail 页共享上下文区可消费视图
- 当前 consumer：无
- 启用条件：仅当 Repository / Product / Module / Decision detail 页已通过唯一 `repositoryId` 接入 `use-project-context-read`，且需要展示规则 / phase 入口定位时才启用

---

## 六、Before / After 样例

### 6.1 Onboarding WelcomeStep 散装解释 → L3 共享常量

**Before**（`onboarding-page.tsx:L281-287`）：
```tsx
"创建一个产品（Product）"
"创建一个仓库（Repository）"
"创建一个模块（Module）"
"记录一条决策（Decision）"
```

**After**：
```tsx
import { PRODUCT_SEMANTIC_LABEL, REPOSITORY_SEMANTIC_LABEL, MODULE_SEMANTIC_LABEL, DECISION_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'

`登记一个${PRODUCT_SEMANTIC_LABEL}（Product）`
`登记一个${REPOSITORY_SEMANTIC_LABEL}（Repository）`
`登记一个${MODULE_SEMANTIC_LABEL}（Module）`
`记录一条${DECISION_SEMANTIC_LABEL}（Decision）`
```

### 6.2 Summary Card 无语义标签 → 共享常量

**Before**（`module-summary-card.tsx`）：
```tsx
<h3>模块名称</h3>
```

**After**：
```tsx
import { MODULE_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'

<h3>模块名称</h3>
<span className="text-xs text-muted-foreground">{MODULE_SEMANTIC_LABEL}</span>
```

---

## 七、明确不做清单

1. 不修改 8 个既有 `use-*-read.ts` 的输入锚点、缓存键、输出 shape 或数据映射逻辑
2. 不在切片 read owner 中新增跨切片共享摘要拼装
3. 不把页面专属字段（如 `bound_modules`、`linked_modules`、`releases`）升格为 L3 共享合同
4. 不新增 `GetProjectContext` 以外的第二套共享 query（已由 phase12-07 判定无需新增）
5. 不让页面直接创建 `createConnectTransport()` 或 `createClient()` 来调用 `GetProjectContext`
6. 不让 mutation owner 跨切片 invalidate 其他切片的 query
7. 不让衍生消费页（dashboard/onboarding/review）直接调用 `GetProjectContext`
8. 不反向改写 `phase12-04` 的 primary owner 页面职责
9. 不反向改写 `phase12-05` 的共享 owner 分层
10. 不反向改写 `phase12-07` 的后端合同"无需新增"结论

---

## 八、与上游文档的一致性声明

| 上游文档 | 关键冻结内容 | 本设计对齐 |
|---|---|---|
| `phase12-05 design §1` | 共享只读 owner 四层分层（L1-L4） | 本设计作为 L3 consumer + L3 implementation，不越权改写 L1/L2 |
| `phase12-05 design §2` | 三类页面承接矩阵 | §3.1-3.3 分别设计直接/间接/衍生页面的读路径 |
| `phase12-05 design §3` | 复用 vs 新增判定规则 | §2.1 三层职责边界 + §4.3 不回收清单，遵守判定规则 |
| `phase12-07 design §4` | 当前无需新增后端合同 | §1.2 仅新增 L3 前端文件，不新增后端合同 |
| `phase12-07 design §5` | 导出结果与共享只读视图关系 | L3-1 为前端共享语义常量，不创建第二套 agent 合同 |
| `phase12-04 design §1-2` | 影响对象清单与 surface 承接矩阵 | §4.4 回收清单与 phase12-04 一致 |
| `architecture_plan §4.6B` | 设计推进顺序 05→07→06 | 本设计消费 05/07 结论，不逆行修改 |
