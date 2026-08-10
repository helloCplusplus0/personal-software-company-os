# Phase06-09 复用感知派生读、Dashboard 挂接与详情挂接设计 Spec

## Why

`phase06-04` 已经冻结了 `module_reuse_summary / capability_summary` 的字段、聚合口径、事实来源与页面挂接位，但它仍停在“读模型语义冻结”层。进入实现前，还需要把 `ReuseSummaryRead` 的 query owner、Dashboard 与详情页的最小集成方式、刷新语义、空态/零复用态差异和它与 `phase05` 已冻结 `FeedbackSignalRead` 的边界收成单值结论，否则后续前端与后端会分别长出不同的作用域 query、不同的 section 状态模型和不同的 capability 映射逻辑。

## What Changes

- 冻结 `ReuseSummaryRead` 的 query owner、最小查询作用域与 query key 设计
- 冻结 Dashboard / Module Detail / Product Detail 的最小集成方式与局部状态模型
- 冻结复用反馈与 `phase05` 既有反馈信号的边界关系
- 冻结最小统计字段、时间字段、新鲜度与刷新语义
- 冻结成功空态、零复用态、有复用态和读取失败态的展示差异
- 冻结 `capability_summary` 的最小事实来源、字段映射与缺省值策略在实现层的承接方式
- 冻结前端每页单个 `ReuseSummaryRead` 页面级 query 的查询组织
- 明确当前子任务不提前冻结 `ReuseSummaryRead` 的后端模块、HTTP 路由数量、platform 装配函数与 `.proto` 文件分组

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-04` 复用读模型与页面挂接位冻结
  - `phase06-05` 合同、传输与前端 `query` 只读边界
  - `phase05-10` Dashboard + Feedback formal spec
- Affected code:
  - 后续 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - 后续 `frontend/src/features/dashboard/components/asset-feedback-section.tsx` 或等价复用快照子区
  - 后续 `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - 后续 `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - 后续新增 `frontend/src/features/reuse-summary/` 或等价切片
  - 后续 `ReuseSummaryRead` 对应前端 `types.ts / api-adapter.ts`
  - 后续 `phase06-10` 承接的 `.proto`、HTTP DTO 与后端模块落点

## ADDED Requirements

### Requirement: `ReuseSummaryRead` query owner 冻结

系统 SHALL 将复用感知的正式 query owner 冻结为单数 `ReuseSummaryRead`，并由该 owner 统一承接 `module_reuse_summary / capability_summary` 的不同页面作用域读取。

#### Scenario: owner 单值化

- **WHEN** 接手者实现复用感知查询
- **THEN** 必须以单数 `ReuseSummaryRead` 作为唯一 query owner
- **AND** 该 owner 统一承接：
  - Dashboard 作用域复用快照
  - Module Detail 作用域复用快照
  - Product Detail 作用域复用快照
- **AND** 当前阶段不得为 `module_reuse_summary` 与 `capability_summary` 各自建立第二套平行 query owner

#### Scenario: 固定数量场景的前端查询组织

- **WHEN** 前端在 Dashboard / Module Detail / Product Detail 中消费 `ReuseSummaryRead`
- **THEN** 每个页面必须只新增一个页面级 `ReuseSummaryRead` `useQuery`
- **AND** 当前阶段不要求为了“统一”而强制改写为 `useQueries`
- **AND** 该页面级 query 的成功结果必须同时承接：
  - `module_reuse_summary`
  - `capability_summary`
- **AND** 只有当同一页面需要动态数量的复用感知 query owner 时，才允许切换到 `useQueries`

#### Scenario: 页面级结果组合

- **WHEN** 接手者定义 `ReuseSummaryRead` 的页面级消费模型
- **THEN** 前端必须消费一个页面级结果对象
- **AND** 该结果对象至少同时包含：
  - `module_reuse_summary[]`
  - `capability_summary[]`
- **AND** `capability_summary` 允许以成功空数组返回
- **AND** 当前子任务只冻结页面级读模型组合，不提前冻结其背后的 HTTP 路由数量或 `.proto` 拆分方式

### Requirement: `ReuseSummaryRead` 作用域与 query key 设计冻结

系统 SHALL 冻结 `ReuseSummaryRead` 的最小作用域模型与 query key 规则，避免后续失效与重读策略发散。

#### Scenario: Dashboard 作用域

- **WHEN** Dashboard 读取复用感知
- **THEN** `ReuseSummaryRead` 必须以 `scope = dashboard` 读取全局复用快照
- **AND** query key 必须至少包含：
  - `'reuse-summary'`
  - `'dashboard'`
- **AND** 当前阶段不得在 Dashboard query key 中混入 `FeedbackSignalRead` 的 key 前缀

#### Scenario: Module Detail 作用域

- **WHEN** Module Detail 读取复用感知
- **THEN** `ReuseSummaryRead` 必须以 `scope = module-detail` + `moduleId` 读取当前 Module 作用域复用快照
- **AND** query key 必须至少包含：
  - `'reuse-summary'`
  - `'module-detail'`
  - `moduleId`

#### Scenario: Product Detail 作用域

- **WHEN** Product Detail 读取复用感知
- **THEN** `ReuseSummaryRead` 必须以 `scope = product-detail` + `productId` 读取当前 Product 作用域复用快照
- **AND** query key 必须至少包含：
  - `'reuse-summary'`
  - `'product-detail'`
  - `productId`

#### Scenario: 前缀失效约束

- **WHEN** 前端写路径变更已提交状态并可能影响复用感知
- **THEN** 允许通过 `invalidateQueries({ queryKey: ['reuse-summary'] })` 进行前缀失效
- **AND** 也允许按页面作用域使用更精确 key 做局部失效
- **AND** 当前阶段不得把复用感知失效混入 `['dashboard-feedback-signals']` 或其他既有 query key

### Requirement: Dashboard 最小集成方式冻结

系统 SHALL 将 Dashboard 中的复用感知集成方式冻结为 `Asset Feedback` 区块内部的独立复用快照子区域，而不是重写 `phase05` 的页面主线。

#### Scenario: Dashboard 页面编排

- **WHEN** `DashboardHomePage` 编排查询
- **THEN** 必须继续保留 `overview / feedback / recent-activity` 三个既有独立查询
- **AND** 额外增加一个独立的 `ReuseSummaryRead` query
- **AND** 该新增 query 只影响 `Asset Feedback` 内的复用快照子区域
- **AND** 不得让 `ReuseSummaryRead` 失败把整页从 `ready` 打回 `page-error`

#### Scenario: `Asset Feedback` 区块内部结构

- **WHEN** Dashboard 渲染 `Asset Feedback`
- **THEN** 区块内部必须拆为两个并列子区域：
  - `Representative Feedback Signals`
  - `Reuse Snapshot`
- **AND** 两个子区域共享同一个区块标题与外层 section 身份
- **AND** 当前阶段不得把复用快照抬升为新的 Dashboard 一级 section

#### Scenario: Dashboard 局部状态分层

- **WHEN** `FeedbackSignalRead` 或 `ReuseSummaryRead` 任一查询状态变化
- **THEN** `Representative Feedback Signals` 与 `Reuse Snapshot` 必须各自派生独立的 `loading / ready / empty / error`
- **AND** 一个子区域失败不得覆盖另一个子区域的成功结果
- **AND** 区块级重试必须允许分别重试两个子区域，而不是只能整区块一起重试

### Requirement: Module Detail 最小集成方式冻结

系统 SHALL 将 Module Detail 中的复用感知集成方式冻结为 `Module Summary` 邻近的只读摘要区，而不是新开页面或改写既有详情主区结构。

#### Scenario: Module Detail query 编排

- **WHEN** `ModuleDetailPage` 读取当前 `moduleId`
- **THEN** 必须保留既有 `module-detail` 详情查询不变
- **AND** 额外增加一个独立 `ReuseSummaryRead` query，作用域为当前 `moduleId`
- **AND** 当前阶段不得把复用快照合并回 `fetchModuleDetail` 的 canonical 响应中

#### Scenario: Module Detail 挂接位

- **WHEN** 页面进入 `ready`
- **THEN** 复用感知摘要必须挂接在 `ModuleSummaryCard` 邻近区域
- **AND** 若 `module.capability_key` 已填写，页面必须同时展示该 Module 所属 capability 的最小摘要
- **AND** 若未填写 `capability_key`，页面必须只展示 `module_reuse_summary`，并以成功空态解释 capability 缺失

### Requirement: Product Detail 最小集成方式冻结

系统 SHALL 将 Product Detail 中的复用感知集成方式冻结为“已绑定模块相关区域附近的只读复用摘要区”，避免改写既有绑定写入主线。

#### Scenario: Product Detail query 编排

- **WHEN** `ProductDetailPage` 读取当前 `productId`
- **THEN** 必须保留既有 `product-detail` 详情查询不变
- **AND** 额外增加一个独立 `ReuseSummaryRead` query，作用域为当前 `productId`
- **AND** 当前阶段不得把复用快照合并回 `fetchProductDetail` 的 canonical 响应中

#### Scenario: Product Detail 挂接位

- **WHEN** 页面进入 `ready`
- **THEN** 复用感知摘要必须挂接在 `ProductModuleBindingPanel` 或其邻近区域
- **AND** 必须同时展示：
  - 当前 Product 已绑定模块的 `module_reuse_summary`
  - 由这些已绑定且填写了 `capability_key` 的 Module 派生出的 `capability_summary`
- **AND** 当前阶段不得让该摘要区承接绑定写入、解绑写入或 Module 候选筛选逻辑

### Requirement: 最小统计字段、时间字段与新鲜度语义冻结

系统 SHALL 将 `module_reuse_summary / capability_summary` 的统计字段、时间字段和新鲜度语义在实现层继续冻结为单值结论。

#### Scenario: `module_reuse_summary` 统计字段

- **WHEN** 接手者实现 `module_reuse_summary`
- **THEN** 统计字段必须至少承接：
  - `module_id`
  - `reuse_product_count`
  - `latest_reuse_at`
  - `explanation_text`
- **AND** `reuse_product_count` 只表示“当前被多少 Product 直接复用”

#### Scenario: `capability_summary` 统计字段

- **WHEN** 接手者实现 `capability_summary`
- **THEN** 统计字段必须至少承接：
  - `capability_key`
  - `capability_label`
  - `supporting_module_count`
  - `latest_capability_update_at`
  - `empty_state_text`
- **AND** `supporting_module_count` 只表示当前参与该 `capability_key` 聚合的 Module 数量

#### Scenario: 新鲜度语义

- **WHEN** `Product / Module / Repository / Decision` 的已提交状态发生变化并影响复用感知
- **THEN** 复用感知查询再次读取时必须反映最新已提交状态
- **AND** 当前阶段不依赖异步统计表、离线聚合作业或后台批处理才能更新页面

### Requirement: Dashboard 与详情页刷新语义冻结

系统 SHALL 将复用感知的刷新语义冻结为“按作用域局部重读”，不让复用快照的 refetch 拖垮整页或 canonical detail 主读。

#### Scenario: Dashboard 局部重试

- **WHEN** Dashboard 中 `ReuseSummaryRead` 失败
- **THEN** 重试入口只允许失效或重读 `['reuse-summary', 'dashboard']`
- **AND** 不得因此强制重试 `overview / feedback / recent-activity`

#### Scenario: Detail 页局部重试

- **WHEN** Module Detail 或 Product Detail 中 `ReuseSummaryRead` 失败
- **THEN** 重试入口只允许重读当前 detail 作用域下的 `ReuseSummaryRead`
- **AND** 不得把 canonical detail 主查询一起打回 loading

#### Scenario: 稳定 UI

- **WHEN** 当前作用域已有上一版复用快照数据，且 `ReuseSummaryRead` 因局部刷新再次 pending
- **THEN** 前端允许使用 `placeholderData` 延续上一版数据以避免区块闪烁
- **AND** 该能力只作用于复用快照子区域，不改变整页或 canonical detail 的状态模型

### Requirement: `phase06-09` 的非冻结边界

系统 SHALL 明确当前子任务只冻结复用感知的页面级 query owner、集成方式与展示语义，不提前替后续合同/后端子任务做决定。

#### Scenario: 后端模块与 platform 装配

- **WHEN** 接手者讨论 `ReuseSummaryRead` 的后端模块目录、candidate reader 文件名、`platform/router.go` 装配函数名
- **THEN** 当前子任务不得把这些细节写成单值冻结项
- **AND** 这些结论必须留给后续后端/合同子任务承接

#### Scenario: `.proto` 与 HTTP 路由数量

- **WHEN** 接手者讨论 `ReuseSummaryRead` 的 `.proto` 文件路径、service 分组、HTTP 路由数量或 query 参数设计
- **THEN** 当前子任务不得把这些细节写成单值冻结项
- **AND** 当前子任务只要求后续合同设计必须能单值支撑“每页一个 `ReuseSummaryRead` 页面级 query 同时承接两类 summary”

### Requirement: 复用反馈与既有反馈信号边界冻结

系统 SHALL 将复用反馈与 `phase05` 既有反馈信号边界冻结为“两类不同意图的读模型并列存在”，避免相互吞并。

#### Scenario: 与 `FeedbackSignalRead` 的边界

- **WHEN** 接手者设计 Dashboard 数据流
- **THEN** `FeedbackSignalRead` 继续只服务：
  - `Current Focus`
  - `Asset Feedback` 中的 representative feedback signals
- **AND** `ReuseSummaryRead` 只服务复用快照子区域
- **AND** 当前阶段不得把“缺口反馈优先级”与“复用快照统计”合并为同一排序序列

#### Scenario: 与 Dashboard 主 CTA 的边界

- **WHEN** 接手者扩展 Dashboard 主内容
- **THEN** `ReuseSummaryRead` 不得生成新的主 CTA、一级行动队列或替代 `Current Focus`
- **AND** 复用快照只作为感知反馈，不升级为经营动作总线

### Requirement: 空状态、零复用状态与有复用状态差异冻结

系统 SHALL 明确区分“成功但无数据”“成功但当前尚未形成跨 Product 复用”和“成功且已有复用”三类展示差异。

#### Scenario: Dashboard 成功空态

- **WHEN** `ReuseSummaryRead` 成功，但当前系统内还不足以形成任何 `module_reuse_summary` 或 `capability_summary`
- **THEN** Dashboard 的 `Reuse Snapshot` 必须展示成功空态
- **AND** 空态文案必须表达“当前还没有可展示的复用反馈”

#### Scenario: 零复用状态

- **WHEN** 当前 Module 已存在，但 `reuse_product_count <= 1`
- **THEN** 页面必须展示“当前尚未形成跨 Product 复用”的零复用解释
- **AND** 该状态属于成功读取，不得显示错误样式或重试入口

#### Scenario: 有复用状态

- **WHEN** `reuse_product_count > 1` 或当前存在 capability 聚合结果
- **THEN** 页面必须展示最小统计值、时间字段和解释文案
- **AND** 不得退化成只有数字没有语义的裸计数

### Requirement: `capability_summary` 事实来源、映射与缺省策略冻结

系统 SHALL 将 `capability_summary` 的实现层事实来源、显示映射与缺省策略继续冻结为单值结论。

#### Scenario: 事实来源

- **WHEN** 接手者实现 capability 聚合
- **THEN** 必须以 `Module.capability_key` 作为唯一聚合主键来源
- **AND** 不得引入独立 `Capability` 实体、表或可编辑字典

#### Scenario: 显示映射

- **WHEN** 页面展示 `capability_label`
- **THEN** 必须来自系统内置 `capability_key -> capability_label` 映射
- **AND** 当前阶段不得由后端、前端和 fixture 各自维护三套不同映射表

#### Scenario: 缺省值策略

- **WHEN** `Module.capability_key` 为空或当前作用域没有任何可聚合 capability
- **THEN** `capability_summary` 必须返回成功空态所需字段
- **AND** `empty_state_text` 必须作为正式字段承接缺省解释文案
- **AND** 当前阶段不得把“未填写 capability_key”误报为读取失败

### Requirement: 前端 `useQuery` 组织单值化冻结

系统 SHALL 将前端在 Dashboard / Module Detail / Product Detail 中的 `ReuseSummaryRead` 查询组织冻结为“每页单个页面级 `useQuery` + 统一成功结果对象”。

#### Scenario: 页面级单 query

- **WHEN** 前端在 Dashboard / Module Detail / Product Detail 中消费 `ReuseSummaryRead`
- **THEN** 每个页面必须只发起 1 个页面级 `ReuseSummaryRead` `useQuery`
- **AND** 该 query 的成功结果必须同时提供 `module_reuse_summary[]` 与 `capability_summary[]`
- **AND** `capability_summary` 无可聚合数据时，必须以成功空数组或成功空态字段返回
- **AND** 当前阶段不得把同一页面的复用感知再拆成第二个页面级 query owner

## MODIFIED Requirements

### Requirement: Dashboard Asset Feedback 区块职责

`Dashboard` 中现有 `Asset Feedback` 区块在 `phase06-09` SHALL 继续承接 `phase05` 已冻结的反馈信号主线，同时扩展为“共享外层区块 + 内部分层子区域”的结构，以容纳复用快照。

#### Scenario: Asset Feedback 区块职责扩展

- **WHEN** 接手者扩展 `Asset Feedback` 区块
- **THEN** 必须继续保留 `phase05` 已冻结的 representative feedback signals 子区域
- **AND** 可以在同一区块内追加 `Reuse Snapshot` 子区域
- **AND** 两个子区域的 query、状态、重试入口与空态文案必须彼此独立
- **AND** 不得因此改变 `Dashboard Home` 的一级区块数量、标题体系或返回路径语义

## REMOVED Requirements

### Requirement: 复用感知必须合并进既有 `FeedbackSignalRead` 才算最小实现

**Reason**: `phase05` 已冻结 `FeedbackSignalRead` 的职责边界；把复用快照硬并回去会破坏既有 Dashboard 主线，也会让缺口反馈和复用统计混成第二套模糊排序语义。

**Migration**: 后续实现统一改为：`FeedbackSignalRead` 继续只服务反馈信号，复用感知由独立 `ReuseSummaryRead` 以局部子区域方式接入 Dashboard 与详情页。
