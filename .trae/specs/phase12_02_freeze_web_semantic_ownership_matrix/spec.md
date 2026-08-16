# phase12-02 冻结四实体在 Web 端的正式语义承接矩阵 Spec

## Why
`phase12-01` 已经冻结了本阶段的主交付、成功标准与非目标，但如果不继续把四实体在 Web 端到底由哪些页面、组件和共享承接位正式表达冻结下来，后续 `phase12-04` 仍会在设计时重新猜“哪一页负责讲语义、哪一页只是跟随回归、哪些对象允许无需改动”。

## What Changes
- 冻结 `Product / Repository / Module / Decision` 在 Web 端的正式解释口径
- 冻结必须显式承接四实体冻结语义的 primary owner 页面、摘要组件与跟随回归对象
- 冻结空态、提示文案与下一步动作说明的显式审计面，要求与页面、组件同等级逐项审计
- 冻结 route shell、列表页、toolbar 与搜索 store 在当前阶段的非 primary owner 身份
- 冻结共享承接位、禁止散落重复解释的边界，以及 `no-change` 结果的显式记录规则
- 对齐 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中关于四实体语义承接矩阵的表达

## Impact
- Affected specs: `phase12_semantic_alignment_and_readonly_consumption_foundation`
- Affected code:
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`
  - `frontend/src/routes/*`
  - `frontend/src/features/*/pages/*`
  - `frontend/src/features/*/components/*`

## ADDED Requirements

### Requirement: 冻结四实体在 Web 端的正式解释口径
系统规格 SHALL 冻结 `Product / Repository / Module / Decision` 在 Web 端的正式解释口径，并要求后续页面、摘要卡片、空态、提示文案与下一步动作说明都回到同一套冻结语义。

当前阶段至少固定如下语义：

- `Product`：经营目标与交付容器
- `Repository`：代码仓库身份对象与项目锚点
- `Module`：可复用能力资产，允许后置提炼
- `Decision`：规则、约束、选择与依据的索引对象

补充冻结：

- 当前阶段只承接表达层与消费层对齐，不承接 schema、关系主线或 canonical owner 重写；
- `Module` 在 Web 端不得再主要被解释为普通模块登记对象；
- `Decision` 在 Web 端不得再主要被解释为孤立文本卡片；
- 若现有字段命名、页面命名或旧文案与上述语义存在张力，优先通过解释性语言、摘要表达与共享只读语义承接解决，而不是先动结构。

#### Scenario: 用户从 Web 端读取四实体语义
- **WHEN** 用户查看四实体详情页、摘要卡片或相关说明文案
- **THEN** 能直接得到与 `phase11` 冻结口径一致的四实体解释
- **AND** 不需要执行者在实现时再临场解释“这个对象到底意味着什么”

### Requirement: 冻结 Web 端 primary owner 与跟随回归矩阵
系统规格 SHALL 冻结四实体语义在 Web 端的 primary owner、跟随回归对象与非 owner 对象，使后续执行者可以机械判断“哪些页面必须显式承接语义、哪些页面只负责跟随回归、哪些对象默认不作为语义主承接位”。

当前阶段至少固定如下矩阵：

1. **Primary owner 页面**
   - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
   - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
   - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
   - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
2. **Primary owner 摘要组件**
   - `frontend/src/features/product-registry/components/product-summary-card.tsx`
   - `frontend/src/features/repository-binding/components/repository-summary-card.tsx`
   - `frontend/src/features/module-registry/components/module-summary-card.tsx`
   - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
3. **跟随回归页面**
   - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
   - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
   - `frontend/src/features/review/pages/daily-review-page.tsx`
   - `frontend/src/features/review/pages/weekly-review-page.tsx`
4. **跟随回归组件**
   - `frontend/src/features/dashboard/components/current-focus-section.tsx`
   - `frontend/src/features/dashboard/components/recent-activity-section.tsx`
   - `frontend/src/features/review/components/review-page-shell.tsx`
5. **Route shell**
   - `frontend/src/routes/dashboard.tsx`
   - `frontend/src/routes/onboarding.tsx`
   - `frontend/src/routes/reviews/daily.tsx`
   - `frontend/src/routes/reviews/weekly.tsx`
   - `frontend/src/routes/products/$productId.tsx`
   - `frontend/src/routes/repositories/$repositoryId.tsx`
   - `frontend/src/routes/modules/$moduleId.tsx`
   - `frontend/src/routes/decisions/$decisionId.tsx`
   - 这些对象只负责稳定路由承接与交接，不作为独立语义 primary owner
6. **默认非 primary owner 对象**
   - `List pages`
   - toolbar
   - 搜索 store
   - 当前阶段默认只做跟随回归，不承担四实体正式语义冻结职责

#### Scenario: 执行者判断页面责任归属
- **WHEN** 后续执行者为 `phase12-04` 产出语义一致性设计
- **THEN** 能机械判断每个对象属于 primary owner、跟随回归还是默认非 primary owner
- **AND** 不会把 route shell、列表页或临时容器误当成四实体语义的正式主承接位

### Requirement: 冻结空态、提示文案与下一步动作说明的显式审计面

系统规格 SHALL 将空态、提示文案与下一步动作说明纳入与页面、组件同等级的显式审计面，要求后续 `phase12-04` 必须逐项审计、分类并留档，不得只以“哪些页面必须展示什么语义”做总括性覆盖。

当前阶段至少固定如下审计面：

1. **空态（Empty States）**
   - 各实体 List 页空态文案（`ProductListPage` / `RepositoryListPage` / `ModuleListPage` / `DecisionListPage`）
   - Dashboard 无数据空态
   - Review 无待处理项空态（`DailyReviewPage` / `WeeklyReviewPage`）
   - Onboarding 各步骤空态
   - 空态文案必须回到四实体冻结语义，不得回落为“暂无数据”一类无语义占位

2. **提示文案（Hint Text）**
   - Onboarding `WelcomeStep` 中对四实体的介绍文案
   - Dashboard 各区块标题与说明文案
   - Review 页面头部说明文案
   - Detail 页中 entity 描述性标题与引导文案
   - 提示文案不得再使用与冻结语义冲突的旧解释（如 Module 被描述为“模块登记”）

3. **下一步动作说明（Next-Action Descriptions）**
   - `ModuleNextActionBar` 中的动作标签与说明
   - `DecisionDetailSummaryCard` 中的状态推进 CTA 文案
   - `DashboardPrimaryActionPanel` 中的主 CTA 说明
   - `OnboardingCtaButton` 中的引导文案
   - `ReviewActionFooter` 中的完成动作说明
   - 下一步动作说明必须与四实体冻结语义一致，不得把 Module 的动作描述为“登记模块”，不得把 Decision 的动作描述为“提交决策”

补充冻结：

- 这三类 surface 必须与页面、组件、route shell 一起进入 `phase12-04` 审计范围；
- 三类 surface 均必须被标记为 `must-change`、`follow-regression` 或 `no-change` 之一；
- `no-change` 必须逐项留档并写明理由；
- 不允许用“页面已经在 owner 矩阵里了”跳过对这三类 surface 的独立审计。

#### Scenario: 执行者审计空态/提示文案/下一步动作说明
- **WHEN** 后续执行者为 `phase12-04` 产出语义一致性设计
- **THEN** 能机械判断每个空态、每条提示文案、每个下一步动作说明属于 `must-change`、`follow-regression` 还是 `no-change`
- **AND** 不会因为“owner 矩阵已经覆盖了页面”而漏掉对这些 surface 的独立审计

### Requirement: 冻结共享承接位与禁止散落重复解释的边界
系统规格 SHALL 冻结四实体解释在切片页面、切片组件与跨切片共享只读承接位之间的边界，避免在多个页面散落复制第二套解释。

当前阶段至少固定如下边界：

- 四实体语义的默认承接位仍是各 feature 自己的 `pages/`、`components/` 与 `data/`；
- 当 `3+` 页面或切片稳定复用同一份 repository-scoped 语义摘要、规则入口或文档入口解释时，才允许晋升到 `frontend/src/features/project-context/`；
- `frontend/src/features/project-context/` 只允许承接受控只读 adapter、共享语义摘要与入口定位视图，不得并列定义第二套 canonical 字段语义；
- 切片页面与组件不得各自重新拼装跨切片共享语义摘要；若已经达到稳定复用条件，必须回收到唯一共享承接位；
- 不允许把“共享呈现”偷换成新的写路径 owner、页面私有状态 owner 或第二套事实源。

#### Scenario: 多页面复用同一段四实体解释
- **WHEN** `3+` 页面需要复用同一段四实体语义摘要或入口说明
- **THEN** 该解释必须晋升到 `frontend/src/features/project-context/` 或保持单向派生自同一只读事实源
- **AND** 不会继续散落在多个切片页面中各讲一套

### Requirement: 冻结 no-change 显式记录规则
系统规格 SHALL 要求 `phase12-02` 把“允许无需改动”的口径冻结为显式记录规则，而不是默认跳过。

当前阶段至少固定如下规则：

- 所有进入 `phase12-04` 审计范围的 route / page / component / data owner / 空态 / 提示文案 / 下一步动作说明 都必须被标记为 `must-change`、`follow-regression` 或 `no-change` 之一；
- `no-change` 必须记录“不改仍满足当前阶段冻结口径”的理由；
- `follow-regression` 不等于“无需审计”，而是表示其语义承接跟随 primary owner 或共享承接位回归；
- 若某对象没有被显式分类或没有留下 `no-change` 理由，则不得视为 `phase12-02` 已完成冻结。

#### Scenario: 审计结果包含无需改动对象
- **WHEN** 执行者判断某页面或组件当前无需改动
- **THEN** 必须显式记录其 `no-change` 身份与理由
- **AND** 不允许用“本轮先跳过”替代正式冻结结果

## MODIFIED Requirements

### Requirement: phase12 三件套中的四实体语义承接表达
`phase12` 三件套中的四实体语义承接表达 SHALL 对齐为同一口径：

- 四实体正式语义：`Product / Repository / Module / Decision` 必须与 `phase11` 冻结口径一致
- primary owner：四类详情页与四类摘要组件
- 跟随回归：`Dashboard / Onboarding / Daily Review / Weekly Review` 页面与其容器组件
- 非 primary owner：route shell、list pages、toolbar 与搜索 store
- 空态 / 提示文案 / 下一步动作说明：必须与页面、组件同等级进入审计范围，逐项标记 `must-change / follow-regression / no-change` 并留档
- 显式记录规则：允许 `no-change`，但必须逐项留档并说明理由

不得再出现：

- 一个文档要求详情页是 primary owner，另一个文档又允许列表页或 route shell 承担同等级语义冻结职责；
- 一个文档允许共享语义晋升到 `frontend/src/features/project-context/`，另一个文档又默认页面本地重复解释；
- 一个文档允许 `no-change` 逐项留档，另一个文档又把未审计对象默认为“无需改动”。

#### Scenario: 三件套对齐四实体语义承接矩阵
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 能获得同一套四实体解释、owner 分类与 `no-change` 记录规则
- **AND** 后续 `/spec` 与实现不需要再次猜测页面责任边界

## REMOVED Requirements

### Requirement: 将 Web 端语义承接分类留给后续设计临场决定
**Reason**: `phase12-02` 的目标就是先把四实体在 Web 端由谁正式承接冻结下来，避免 `phase12-04` 再次把 owner 分类、共享承接位与 `no-change` 规则留到实现前临场补判断。
**Migration**: 后续设计与实现必须直接承接 `phase12-02` 冻结的四实体语义、owner 矩阵与显式记录规则，不再新增第二套页面责任解释。
