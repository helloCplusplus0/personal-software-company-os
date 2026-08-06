# Phase03-05 前端页面、路由与组件分层设计 Spec

## Why

`phase03-01 ~ 04` 已经冻结了 `Decision Center` 的页面边界、模板读模型、目标范围、入口上下文、数据读写范围与错误语义前提，但这些结论还不足以直接进入前端实现。为了让 `phase03-06 / 07 / 08 / 10` 与后续正式编码不再依赖临场发挥，当前子任务必须产出一份可直接指导 `frontend/` 落地的页面、路由与组件分层设计。

## What Changes

- 产出 `Decision List / Decision Create / Decision Detail` 的具体页面分层设计
- 产出一组可直接映射到前端实现的最小路由结构与 URL 语义
- 产出列表、模板表单、目标关联面板的组件树与组件职责边界
- 产出 `PC / 移动浏览器` 双场景下的布局降级策略与信息重排规则
- 产出可直接映射到 `frontend/` 的前端文件落点与路由树
- 明确当前阶段不引入第二套移动端 UI 架构、不拆出 `Product / Repository` 的独立前端主线

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `frontend/` 中 `Decision Center` 的路由定义、页面文件组织、布局壳层、页面组件拆分、响应式布局容器
- Affected code: 预期会直接映射到 `frontend/src/routes/decisions/*`、`frontend/src/features/decision-center/pages/*`、`frontend/src/features/decision-center/components/*`

## ADDED Requirements

### Requirement: 前端页面结果必须可直接映射到页面文件

系统 SHALL 将当前阶段前端页面设计收敛为一组可以直接映射到 `frontend/` 页面文件的页面集合，而不是只停留在抽象页面名称。

#### Scenario: 页面集合判定

- **WHEN** 后续实现开始创建 `Decision Center` 页面文件
- **THEN** 页面集合必须至少包含 `DecisionListPage`、`DecisionCreatePage`、`DecisionDetailPage`
- **AND** 不得在当前阶段额外拆出 `DecisionLinkPage`、`ProductPage`、`RepositoryPage` 的独立主线页面

#### Scenario: 页面职责判定

- **WHEN** 上述页面被映射到具体实现文件
- **THEN** `DecisionListPage` 只承接列表读取、筛选入口、创建入口与进入详情
- **AND** `DecisionCreatePage` 只承接 `RecordDecision`
- **AND** `DecisionDetailPage` 只承接详情读取、已关联目标展示、`Decision -> Module` 候选读取与 `LinkDecisionToTarget`

### Requirement: 路由树必须冻结到文件落点层

系统 SHALL 将 `Decision Center` 当前阶段的前端路由设计冻结到“路由树 + 文件落点”层，确保后续实现可以直接按既定结构创建文件。

#### Scenario: 最小文件落点

- **WHEN** 后续实现开始创建前端路由与页面文件
- **THEN** 当前阶段至少应能映射出以下文件层级语义：
- **AND** `frontend/src/routes/decisions/index.tsx`
- **AND** `frontend/src/routes/decisions/new.tsx`
- **AND** `frontend/src/routes/decisions/$decisionId.tsx`
- **AND** `frontend/src/features/decision-center/pages/decision-list-page.tsx`
- **AND** `frontend/src/features/decision-center/pages/decision-create-page.tsx`
- **AND** `frontend/src/features/decision-center/pages/decision-detail-page.tsx`

#### Scenario: 路由树表达

- **WHEN** 后续实现需要查看当前阶段的最小路由树
- **THEN** 路由树至少应表达为：
- **AND** `/decisions`
- **AND** `/decisions/new`
- **AND** `/decisions/:decisionId`
- **AND** 不得把 `Product`、`Repository` 提前扩写为并列主树

#### Scenario: Module Detail 入口到新主线的文件落点映射

- **WHEN** 后续实现需要将 `Module Detail` 的两个决策入口接到新 `Decision Center` 路由主线
- **THEN** 现有 `frontend/src/routes/modules/$moduleId.tsx` 承载的 `ModuleDetailPage` 必须继续作为 `Module Detail` 的页面宿主
- **AND** 现有 `frontend/src/features/module-registry/components/module-decision-entry-panel.tsx` 中的 `ModuleDecisionEntryPanel` 必须从当前只读展示升级为承接两个入口动作的触点组件
- **AND** `ModuleDecisionEntryPanel` 中"为当前 `Module` 记录决策"触点必须导航到 `DecisionCreateRoute`（`/decisions/new`）并携带当前 `Module` 的目标标识与可展示名称作为上下文
- **AND** `ModuleDecisionEntryPanel` 中"查看当前 `Module` 相关决策"触点必须导航到 `DecisionListRoute`（`/decisions`）
- **AND** 不得在 `Module Detail` 侧新增中间分发组件或中间路由层
- **AND** 入口上下文的具体传递机制（路由 state / 搜索参数 / 客户端状态）延后至 `phase03-06` 冻结

### Requirement: 路由设计必须冻结到 URL 语义层

系统 SHALL 将当前阶段前端最小路由结构冻结到可直接实现的 URL 语义层，而不是只保留抽象 route name。

#### Scenario: 最小路由结构

- **WHEN** 后续实现创建前端路由
- **THEN** 最小路由结构至少应承接以下 URL 语义：
- **AND** `DecisionListRoute` -> `/decisions`
- **AND** `DecisionCreateRoute` -> `/decisions/new`
- **AND** `DecisionDetailRoute` -> `/decisions/:decisionId`
- **AND** `DecisionListRoute` 的 `/decisions` 必须承接 `queryText` 与 `statusFilter` 作为路由搜索参数
- **AND** 路由搜索参数与页面状态之间的详细映射规则延后至 `phase03-06` 冻结

#### Scenario: 路由进入关系

- **WHEN** 页面之间的进入与返回关系被实现
- **THEN** `DecisionListRoute` 必须能进入 `DecisionCreateRoute`
- **AND** `DecisionListRoute` 必须能进入 `DecisionDetailRoute`
- **AND** `DecisionCreateRoute` 在 `RecordDecision` 成功后必须能进入 `DecisionDetailRoute`
- **AND** 从 `Module Detail` 带上下文进入 `DecisionCreateRoute` 并成功创建后，必须默认进入新建 `Decision` 的 `DecisionDetailRoute`，不得回流到 `Decision Center / List`（承接 `phase03-03` 已冻结结论，默认行为的详细交互流延后至 `phase03-06` 冻结）
- **AND** `DecisionDetailRoute` 必须能返回 `DecisionListRoute`
- **AND** `Module Detail` 中“为当前 `Module` 记录决策”触点必须进入带上下文的 `DecisionCreateRoute`
- **AND** `Module Detail` 中“查看当前 `Module` 相关决策”触点必须进入 `DecisionListRoute`

#### Scenario: 当前阶段不拆子工作台

- **WHEN** 路由结构被进一步展开
- **THEN** 当前阶段不得把 `Decision Detail` 提前拆成独立的子路由工作台
- **AND** 不得增加 `Product`、`Repository` 的独立路由主线

### Requirement: 页面壳层与组件树必须冻结到实现结构层

系统 SHALL 为每个页面产出可直接进入实现的页面壳层与组件树设计，而不是只描述“有一个区块”。

#### Scenario: 列表页组件树

- **WHEN** 后续实现 `DecisionListPage`
- **THEN** 页面壳层至少应包含 `DecisionListPageShell`
- **AND** 主组件树至少应包含 `DecisionListToolbar`、`DecisionListContent`、`DecisionListTableOrCards`、`DecisionListEmptyState`
- **AND** `DecisionListToolbar` 承接搜索输入、状态筛选与进入 `Decision Create` 的入口
- **AND** `DecisionListContent` 只承接列表读取结果，不承接创建写入或目标关联逻辑
- **AND** `DecisionListEmptyState` 只承接无决策时进入 `Decision Create` 的引导

#### Scenario: 创建页组件树

- **WHEN** 后续实现 `DecisionCreatePage`
- **THEN** 页面壳层至少应包含 `DecisionCreatePageShell`
- **AND** 主组件树至少应包含 `DecisionContextSourcePanel`、`DecisionCreateForm`、`DecisionCreateActions`
- **AND** `DecisionContextSourcePanel` 只承接从 `Module Detail` 带入的来源 `Module` 展示
- **AND** `DecisionCreateForm` 只承接结构化模板字段录入
- **AND** `DecisionCreateActions` 只承接提交与取消动作
- **AND** 不得在当前阶段扩写出自动推荐、AI 建议或目标关联写入面板

#### Scenario: 详情页组件树

- **WHEN** 后续实现 `DecisionDetailPage`
- **THEN** 页面壳层至少应包含 `DecisionDetailPageShell`
- **AND** 主组件树至少应包含 `DecisionDetailSummaryCard`、`DecisionLinkedTargetsSection`、`DecisionPendingLinkTargetCard`、`DecisionModuleCandidatePanel`、`DecisionLinkActions`
- **AND** `DecisionDetailSummaryCard` 只承接决策核心字段、结构化模板字段与最小来源上下文展示
- **AND** `DecisionLinkedTargetsSection` 只承接已建立的 `Decision -> Module` 关联结果
- **AND** `DecisionPendingLinkTargetCard` 只承接从 `Module Detail` 带入的入口上下文中尚未完成正式关联的待关联 `Module`，必须作为显式待关联目标持续展示，直到用户完成 `LinkDecisionToTarget` 或主动放弃关联（承接 `phase03-03` 已冻结结论）
- **AND** `DecisionModuleCandidatePanel` 只承接 `Decision -> Module` 候选读取与目标选择，待关联 `Module` 在候选读取中应作为首选候选或显式待确认目标出现
- **AND** `DecisionLinkActions` 只承接 `LinkDecisionToTarget` 的最小写入触点
- **AND** 不得把候选读取区扩写为独立 `Module` 浏览工作台

### Requirement: 组件归属必须冻结到共享与页面专属边界

系统 SHALL 将当前阶段组件归属冻结到“页面专属组件”和“可复用共享组件”边界，避免后续实现时无序抽象。

#### Scenario: 页面专属组件

- **WHEN** 后续实现讨论页面专属组件
- **THEN** `DecisionListToolbar`、`DecisionListTableOrCards`、`DecisionListEmptyState` 应默认先归属于 `DecisionListPage`
- **AND** `DecisionContextSourcePanel`、`DecisionCreateForm`、`DecisionCreateActions` 默认先归属于 `DecisionCreatePage`
- **AND** `DecisionDetailSummaryCard`、`DecisionLinkedTargetsSection`、`DecisionPendingLinkTargetCard`、`DecisionModuleCandidatePanel`、`DecisionLinkActions` 默认先归属于 `DecisionDetailPage`

#### Scenario: 共享组件边界

- **WHEN** 后续实现讨论共享组件
- **THEN** 只有在多个页面确实共享同一职责时，才允许抽为共享组件
- **AND** 当前阶段不得为了“组件纯洁”提前拆出无明确复用证据的通用组件层

### Requirement: 布局降级策略必须冻结到页面结构层

系统 SHALL 将 `PC / 移动浏览器` 双场景下的布局降级策略冻结到页面结构层，确保同一套 `React Web` 页面在两种场景下都可直接实现。

#### Scenario: PC 页面布局

- **WHEN** 页面在桌面端实现
- **THEN** `DecisionListPage` 应优先采用高信息密度列表布局
- **AND** `DecisionDetailPage` 应优先采用分区式详情布局，使概要、已关联目标、待关联目标与候选关联区同页可见
- **AND** 不得因为桌面端实现而产生第二套专用页面体系

#### Scenario: PC 详情页布局分区

- **WHEN** `DecisionDetailPage` 在桌面端实现
- **THEN** 页面至少应分为概要主区、已关联目标区、待关联目标区、候选读取与目标关联区
- **AND** 待关联目标区在存在入口上下文待关联 `Module` 时必须可见，在无待关联目标时不占用布局空间
- **AND** 候选读取与目标关联区应优先保持同页可见，而不是依赖额外导航切换

#### Scenario: 移动浏览器页面布局

- **WHEN** 页面在移动浏览器实现
- **THEN** `DecisionListPage` 应采用单列列表或卡片重排
- **AND** `DecisionDetailPage` 应将概要、已关联目标、待关联目标、候选读取与目标关联按垂直顺序重排
- **AND** 可对次级信息采用折叠或延迟展开
- **AND** 不得新增独立移动端页面体系

#### Scenario: 移动浏览器创建页布局

- **WHEN** `DecisionCreatePage` 在移动浏览器实现
- **THEN** 来源上下文区、表单区与动作区应采用单列垂直布局
- **AND** 主动作按钮必须在无需横向滚动的前提下可见
- **AND** 不得通过新增移动端专用页面来规避当前布局问题

#### Scenario: 当前阶段移动端边界

- **WHEN** 后续实现讨论移动端承接方式
- **THEN** 不得引入独立 `React Native` 客户端
- **AND** 不得把完整 `PWA` 能力写成当前阶段实现前提
- **AND** 必须继续遵守单一 `React Web` 同时覆盖 `PC` 与移动浏览器的交付策略

## MODIFIED Requirements

### Requirement: Decision Detail 前端承接方式

`Decision Detail` 在当前阶段 SHALL 被解释为一个复合详情页，并在同一页面壳层中统一承接读取、已关联目标展示、待关联目标承接、候选读取与目标关联写入。

#### Scenario: 详情页结构收口

- **WHEN** 后续 `/spec` 或实现讨论 `Decision Detail` 的前端组织方式
- **THEN** 必须保持在同一详情页壳层中整合概要、已关联目标、待关联目标与候选关联能力
- **AND** 不得提前拆出独立子工作台或并列主页面

## REMOVED Requirements

### Requirement: 第二套移动端 UI 架构
**Reason**: 当前阶段已经冻结为单一 `React Web` 交付策略，并要求同时覆盖 `PC` 与移动浏览器 UI。
**Migration**: 若后续确需独立移动端客户端，必须在新的 `phase` 或 `audit` 流程中重新冻结，不在 `phase03-05` 当前设计中处理。
