# phase10-10 落实关键 detail pages 的下一步动作承接矩阵 Spec

## Why

`phase10-09` 已将 `Decision Detail` 落实为唯一正式状态推进承接位，`Dashboard / Daily Review / Current Focus` 的 pending handoff 与 reread 主线也已统一。但 `Product Detail / Module Detail / Repository Detail` 三个关键详情页当前仍停留在"信息展示壳层"状态：有返回按钮、有绑定面板，但缺少"下一步做什么"的页面级主 CTA 承接位，用户在详情页看完后不知道下一步该做什么。

`phase10-05` 已冻结各 detail page 的 CTA inventory 与主次优先级，`phase10-10` 必须把这些设计落地为代码实现，让三个 detail page 都能回答"下一步做什么"，且 CTA 指向 canonical path。

## What Changes

- 为 `Product Detail` 新增 Decision 入口面板与页面级下一步动作区
- 为 `Module Detail` 将只读绑定面板升级为指向 canonical binding owner 的 CTA handoff
- 为 `Repository Detail` 新增 Decision 入口面板与页面级下一步动作区
- 三个 detail page 统一落实"结构缺口 → Decision → 返回 reread"的页面级主/次 CTA 优先级
- 落实从 detail page 返回 Dashboard / Review 后的 reread 语义

## Impact

- Affected specs:
  - `phase10_05_design_dashboard_review_detail_next_step_cta_handoff` — 本 spec 是其 Product/Module/Repository Detail 段的实现
  - `phase10_06_design_frontend_read_write_owner_route_caller_success_handoff` — 本 spec 遵守其单值 owner 约束
  - `phase10_09_land_decision_lifecycle_detail_cta_pending_reread_unification` — 本 spec 不修改其 Decision pending 主线
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
- Affected code:
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/product-registry/data/use-product-detail-read.ts`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/module-registry/components/module-binding-panel.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/repository-binding/data/use-repository-detail-read.ts`

## ADDED Requirements

### Requirement: 三个 detail page 必须具备页面级下一步动作承接区

系统 SHALL 为 `Product Detail / Module Detail / Repository Detail` 各新增一个页面级下一步动作承接区，能够在页面任意状态下回答"当前最应该做什么"，而不是只堆叠一组无优先级按钮。

#### Scenario: 判断 detail page 是否具备正式下一步语义

- **WHEN** 用户进入任一 detail page
- **THEN** 页面必须展示一个明确的下一步动作区
- **AND** 该动作区必须优先展示当前最高优先级的 canonical 动作
- **AND** 不得把"返回列表"或"返回 Dashboard"作为兜底主 CTA（除非无结构缺口且无 pending decision）

### Requirement: `Product Detail` 下一步动作承接矩阵

系统 SHALL 为 `Product Detail` 落实页面级 CTA 矩阵，优先级为：Repository 缺口 > Module 缺口 > 相关 Decision > 返回 reread。

#### Scenario: Product 无 Repository 绑定

- **WHEN** `Product Detail` 读取到当前 Product 没有已绑定 Repository
- **THEN** 主 CTA 必须为"绑定仓库"，指向 canonical Repository Binding 流程
- **AND** 该 CTA 必须在页面级下一步动作区中以主按钮样式展示

#### Scenario: Product 有 Repository 但无 Module 绑定

- **WHEN** `Product Detail` 有已绑定 Repository 但没有已绑定 Module
- **THEN** 主 CTA 必须为"绑定模块"，指向当前页面已有的 ProductModuleBindingPanel
- **AND** 该 CTA 必须在页面级下一步动作区中以主按钮样式展示

#### Scenario: Product 结构完整但存在相关 pending decision

- **WHEN** `Product Detail` 结构缺口已闭合，但存在与当前 Product 相关的 pending decision
- **THEN** 主 CTA 必须为"查看相关决策"，指向相关 Decision Detail
- **AND** 该 CTA 在页面级下一步动作区中以主按钮样式展示

#### Scenario: Product 无结构缺口且无 pending decision

- **WHEN** `Product Detail` 无结构缺口且无 pending decision
- **THEN** 页面级下一步动作区允许展示"返回 Dashboard"作为兜底 CTA
- **AND** 不得继续展示已闭合的"补仓库/补模块" CTA

### Requirement: `Product Detail` Decision 入口面板

系统 SHALL 为 `Product Detail` 新增 Decision 入口面板，承接"为当前 Product 记录决策"与"查看当前 Product 相关决策"两类正式入口。

#### Scenario: 从 Product Detail 进入 Decision Create

- **WHEN** 用户点击 Product Detail 中"记录决策"按钮
- **THEN** 必须导航到带 `sourceProductId / sourceProductName` 的 `/decisions/new`
- **AND** 不得在 Product Detail 侧新增中间路由

#### Scenario: Product Detail 展示已关联决策

- **WHEN** Product Detail 读取到已关联的 decision_links
- **THEN** Decision 入口面板必须展示已关联决策列表
- **AND** 每个决策项必须可点击进入对应 DecisionDetailPage

### Requirement: `Module Detail` 下一步动作承接矩阵

系统 SHALL 为 `Module Detail` 落实页面级 CTA 矩阵，优先级为：Product 绑定 > Repository 映射 > 相关 Decision > 返回 reread。

#### Scenario: Module 无 Product 绑定

- **WHEN** `Module Detail` 读取到当前 Module 没有已绑定 Product
- **THEN** 主 CTA 必须为"绑定产品"，指向 canonical Product Detail（承接绑定写入）
- **AND** 导航到 Product Detail 时必须携带 moduleId / moduleName 来源上下文

#### Scenario: Module 有 Product 绑定但无 Repository 映射

- **WHEN** `Module Detail` 有已绑定 Product 但没有已映射 Repository
- **THEN** 主 CTA 必须为"映射仓库"，指向 canonical Repository Detail（承接映射写入）
- **AND** 导航到 Repository Detail 时必须携带 moduleId / moduleName 来源上下文

#### Scenario: Module 结构完整但存在相关 pending decision

- **WHEN** `Module Detail` 结构缺口已闭合，但存在相关 pending decision
- **THEN** 主 CTA 必须为"查看相关决策"，指向相关 Decision Detail
- **AND** ModuleBindingPanel 中已有的 Decision 入口面板继续保留，但不再作为主 CTA 宿主

#### Scenario: Module 无结构缺口且无 pending decision

- **WHEN** `Module Detail` 无结构缺口且无 pending decision
- **THEN** 页面级下一步动作区允许展示"返回 Dashboard"作为兜底 CTA

### Requirement: `Module Detail` 绑定面板 CTA handoff 升级

系统 SHALL 将 `ModuleDetailPage` 中当前只读的 `ModuleBindingPanel` 升级为指向 canonical binding owner 的 CTA handoff，使其不再只是信息展示。

#### Scenario: 从 Module Detail 去 Product Detail 绑定

- **WHEN** Module 当前无 Product 绑定
- **THEN** ModuleBindingPanel 的"绑定产品" CTA 必须导航到 `/products/$productId`（若已有候选 Product）或 `/products`（若无候选）
- **AND** 导航参数必须携带 `fromModuleDetail=true + moduleId + moduleName`

#### Scenario: 从 Module Detail 去 Repository Detail 映射

- **WHEN** Module 当前无 Repository 映射
- **THEN** ModuleBindingPanel 的"映射仓库" CTA 必须导航到 `/repositories/$repositoryId`（若已有候选）或 `/repositories`（若无候选）
- **AND** 导航参数必须携带 `fromModuleDetail=true + moduleId + moduleName`

### Requirement: `Repository Detail` 下一步动作承接矩阵

系统 SHALL 为 `Repository Detail` 落实页面级 CTA 矩阵，优先级为：Product 绑定 > Module 映射 > 相关 Decision > 返回 reread。

#### Scenario: Repository 无 Product 绑定

- **WHEN** `Repository Detail` 读取到当前 Repository 没有已绑定 Product
- **THEN** 主 CTA 必须为"绑定产品"，展开当前页面已有的 RepositoryProductBindingPanel
- **AND** 该 CTA 在页面级下一步动作区中以主按钮样式展示

#### Scenario: Repository 有 Product 绑定但无 Module 映射

- **WHEN** `Repository Detail` 有已绑定 Product 但没有已映射 Module
- **THEN** 主 CTA 必须为"映射模块"，展开当前页面已有的 RepositoryModuleMappingPanel
- **AND** 该 CTA 在页面级下一步动作区中以主按钮样式展示

#### Scenario: Repository 结构完整但存在相关 pending decision

- **WHEN** `Repository Detail` 结构缺口已闭合，但存在相关 pending decision
- **THEN** 主 CTA 必须为"查看相关决策"，指向相关 Decision Detail

### Requirement: `Repository Detail` Decision 入口面板

系统 SHALL 为 `Repository Detail` 新增 Decision 入口面板，承接"为当前 Repository 记录决策"与"查看当前 Repository 相关决策"两类正式入口。

#### Scenario: 从 Repository Detail 进入 Decision Create

- **WHEN** 用户点击 Repository Detail 中"记录决策"按钮
- **THEN** 必须导航到带 `sourceRepositoryId / sourceRepositoryName` 的 `/decisions/new`
- **AND** 不得在 Repository Detail 侧新增中间路由

### Requirement: 从 detail page 返回 Dashboard / Review 的 reread 语义

系统 SHALL 确保从任一 detail page 返回 Dashboard / Daily Review 后，来源页的 Current Focus、pending signals 与主 CTA 都基于 canonical reread 结果同步更新。

#### Scenario: 从 detail page 完成结构闭合后返回 Dashboard

- **WHEN** 用户从 Dashboard 进入某个 detail page 并完成结构缺口闭合
- **AND** 用户通过 BackToDashboardButton 或返回按钮回到 Dashboard
- **THEN** Dashboard 的 Current Focus 与 pending signals 必须重新评估
- **AND** 若原结构缺口已闭合，Dashboard 主 CTA 必须切换为新的最高优先级事项

#### Scenario: 从 detail page 返回 Daily Review

- **WHEN** 用户从 Daily Review 进入某个 detail page 并完成动作
- **AND** 用户返回 Daily Review
- **THEN** Daily Review 的 Current Focus 必须重新评估
- **AND** 不得继续展示已闭合的结构缺口

## Non-Goals

- 不修改 `Decision Detail` pending 主线或 `Dashboard / Daily Review` 的 canonical 解释
- 不新增独立后端 RPC 或 Proto 合同
- 不修改 `Product / Module / Repository Detail` 既有的绑定写入逻辑（只新增 CTA 展示与 handoff）
- 不修改 `BackToDashboardButton` 的既有行为
- 不新增前端路由

## Open Questions

无。所有设计决策已在 `phase10-05` 中冻结。