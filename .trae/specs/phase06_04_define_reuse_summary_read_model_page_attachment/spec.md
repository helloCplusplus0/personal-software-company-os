# Phase06-04 `module_reuse_summary / capability_summary` 最小读模型与页面挂接位 Spec

## Why

`phase06` 要让用户不只完成录入和数据主权闭合，还要在系统里尽早看见“复用正在发生”的最小反馈。当前如果不先把 `module_reuse_summary / capability_summary` 的字段、聚合口径、页面挂接位和空状态冻结成单值结论，后续 Dashboard、Module Detail、Product Detail 会各自长出一套复用感知语义，`Capability` 也会在实现阶段被临时猜成第二个重实体。

## What Changes

- 冻结 `module_reuse_summary` 的最小字段、统计口径、排序与空状态语义
- 冻结 `capability_summary` 的最小字段、聚合口径、事实来源与空聚合语义
- 冻结未填写 `capability_key` 的 `Module` 在 capability 聚合中的处理方式
- 冻结 Dashboard、Module Detail、Product Detail 的最小挂接位
- 冻结复用感知的新鲜度语义与刷新口径
- 明确当前阶段复用感知是派生读能力，不引入新的一级“复用中心”或 `Capability` 重实体 CRUD

## Impact

- Affected specs: `phase06_onboarding_sovereignty_reuse_foundation`
- Affected code:
  - 后续 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - 后续 `frontend/src/features/dashboard/components/asset-feedback-section.tsx` 或等价 Dashboard 派生反馈子区块
  - 后续 `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - 后续 `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - 后续新增 `ReuseSummaryRead` query owner
  - 后续 `module_reuse_summary / capability_summary` 对应 `.proto` 合同、HTTP DTO 与前端消费模型

## ADDED Requirements

### Requirement: `module_reuse_summary` 最小读模型冻结

系统 SHALL 将 `module_reuse_summary` 冻结为“一个 Module 当前被多少 Product 直接复用”的派生读模型，不引入独立统计表作为唯一事实源。

#### Scenario: `module_reuse_summary` 字段集合

- **WHEN** 接手者定义 `module_reuse_summary` 的最小字段集合
- **THEN** 至少必须带出：
  - `module_id`
  - `reuse_product_count`
  - `latest_reuse_at`：表示“最近一次 `Product` 绑定该 `Module` 的时间”
  - `explanation_text`
- **AND** 当前阶段不得把 `module_reuse_summary` 扩写为完整复用历史时间线或独立分析报表

#### Scenario: `module_reuse_summary` 统计口径

- **WHEN** 系统计算某个 `Module` 的复用情况
- **THEN** 统计口径必须单值冻结为“当前有多少 `Product` 直接绑定了该 `Module`”
- **AND** 当前阶段不得把 `Repository` 映射数、`Decision` 链接数、`Release` 数量混入 `reuse_product_count`

### Requirement: `module_reuse_summary` 列表排序与裁剪冻结

系统 SHALL 为 Dashboard 与 Product Detail 中的 `module_reuse_summary` 列表冻结最小排序与裁剪规则，避免不同页面各自派生不同顺序。

#### Scenario: Dashboard 列表排序

- **WHEN** Dashboard 展示多条 `module_reuse_summary`
- **THEN** 必须按 `reuse_product_count DESC` 排序
- **AND** 当 `reuse_product_count` 相同时，必须按 `latest_reuse_at DESC` 排序
- **AND** 当前阶段 Dashboard 最多只展示前 `5` 条，避免把复用感知扩写为独立列表页

#### Scenario: Product Detail 列表排序

- **WHEN** Product Detail 展示当前 Product 绑定模块的 `module_reuse_summary`
- **THEN** 必须先按“当前 Product 已绑定的 Module”过滤
- **AND** 再按 `reuse_product_count DESC`、`latest_reuse_at DESC` 排序
- **AND** 当前阶段 Product Detail 不设展示上限，全量展示当前 Product 作用域下的 `module_reuse_summary`

### Requirement: `capability_summary` 最小读模型冻结

系统 SHALL 将 `capability_summary` 冻结为基于 `Module.capability_key` 派生聚合得到的能力视角，不引入独立 `Capability` 重实体。

#### Scenario: `capability_summary` 字段集合

- **WHEN** 接手者定义 `capability_summary` 的最小字段集合
- **THEN** 至少必须带出：
  - `capability_key`
  - `capability_label`
  - `supporting_module_count`
  - `latest_capability_update_at`
  - `empty_state_text`
- **AND** 当前阶段不得要求用户单独创建、编辑或删除 `Capability`

#### Scenario: `capability_summary` 聚合口径

- **WHEN** 系统计算 `capability_summary`
- **THEN** 聚合单位必须单值冻结为 `capability_key`
- **AND** `supporting_module_count` 必须表示“当前参与该 `capability_key` 聚合的 Module 数量”
- **AND** `latest_capability_update_at` 必须表示“当前参与该聚合的 Module 中，最新一条已提交更新时间”

### Requirement: `capability_summary` 事实来源冻结

系统 SHALL 将 `capability_summary` 的事实来源冻结为 `Module` 写模型中的轻量 `capability_key` 与系统内置 `capability_label` 映射。

#### Scenario: 事实来源判定

- **WHEN** 接手者实现 `capability_summary`
- **THEN** 必须以 `Module.capability_key` 作为唯一聚合主键来源
- **AND** 必须以系统内置 `capability_label` 映射为当前阶段唯一 label 来源
- **AND** 当前阶段不得引入独立 `capabilities` 表、可编辑能力字典或第二套手工映射事实源

### Requirement: `capability_summary` 列表排序与裁剪冻结

系统 SHALL 为 Dashboard 与 Product Detail 中的 `capability_summary` 列表冻结最小排序与裁剪规则，避免不同页面各自派生不同顺序。

#### Scenario: Dashboard 列表排序

- **WHEN** Dashboard 展示多条 `capability_summary`
- **THEN** 必须按 `supporting_module_count DESC` 排序
- **AND** 当 `supporting_module_count` 相同时，必须按 `latest_capability_update_at DESC` 排序
- **AND** 当前阶段 Dashboard 最多只展示前 `5` 条，避免把能力视角扩写为独立列表页

#### Scenario: Product Detail 列表排序

- **WHEN** Product Detail 展示当前 Product 绑定模块形成的 `capability_summary`
- **THEN** 必须先按“当前 Product 已绑定且填写了 `capability_key` 的 Module”派生
- **AND** 再按 `supporting_module_count DESC`、`latest_capability_update_at DESC` 排序
- **AND** 当前阶段 Product Detail 不设展示上限，全量展示当前 Product 作用域下的 capability 聚合结果

### Requirement: 未声明 capability 的 Module 处理规则冻结

系统 SHALL 冻结未填写 `capability_key` 的 `Module` 在 `capability_summary` 聚合中的处理方式，避免实现阶段继续猜测。

#### Scenario: 未填写 `capability_key` 的 Module

- **WHEN** 某个 `Module` 未填写 `capability_key`
- **THEN** 该 `Module` 不参与当前阶段 `capability_summary` 聚合
- **AND** 该行为不得阻断该 `Module` 参与 `module_reuse_summary` 的计算
- **AND** 该行为不得阻断首轮成功会话成立

#### Scenario: 全部 Module 都未填写 `capability_key`

- **WHEN** 当前读取范围内所有 `Module` 都未填写 `capability_key`
- **THEN** `capability_summary` 必须返回成功空态，而不是错误态
- **AND** 页面必须展示“当前还没有可归纳的 Capability Summary”语义，而不是把空态解释成读取失败

### Requirement: Dashboard 挂接位冻结

系统 SHALL 将 Dashboard 中的复用感知挂接位冻结为现有 `Asset Feedback` 区块内的追加子区块，不引入新的一级 Dashboard 主区块或独立复用页面。

#### Scenario: Dashboard 页面挂接位

- **WHEN** Dashboard 展示当前阶段复用感知
- **THEN** 正式挂接位必须位于现有 `Asset Feedback` 区块内部
- **AND** 该挂接位至少同时承接：
  - `Top Reused Modules`（来自 `module_reuse_summary`）
  - `Capability Snapshot`（来自 `capability_summary`）
- **AND** 当前阶段不得为复用感知新增独立一级导航、独立 `/reuse` 页面或新的 Dashboard 一级 section

### Requirement: Module Detail 挂接位冻结

系统 SHALL 将 Module Detail 中的复用感知挂接位冻结为 `Module Summary` 所在的详情主区，用于展示当前 Module 的直接复用反馈与关联 capability 视角。

#### Scenario: Module Detail 页面挂接位

- **WHEN** 用户进入 `Module Detail`
- **THEN** 页面必须在 `Module Summary` 邻近区域展示当前 `module_id` 的 `module_reuse_summary`
- **AND** 若该 Module 已填写 `capability_key`，页面还必须展示该 Module 所属 `capability_summary` 的最小摘要
- **AND** 当前阶段不得要求用户跳转到独立“复用中心”才能看见这些反馈

### Requirement: Product Detail 挂接位冻结

系统 SHALL 将 Product Detail 中的复用感知挂接位冻结为当前 Product 绑定模块相关区域的邻近子区块，用于展示“本 Product 已绑定模块的复用情况”和“这些模块形成的 capability 快照”。

#### Scenario: Product Detail 页面挂接位

- **WHEN** 用户进入 `Product Detail`
- **THEN** 页面必须在已绑定模块相关区域附近展示当前 Product 作用域下的 `module_reuse_summary`
- **AND** 页面必须展示基于“当前 Product 已绑定且填写了 `capability_key` 的 Module”派生出的 `capability_summary`
- **AND** 当前阶段不得把 Product Detail 的 capability 视图解释为 Product 自己拥有独立 Capability 实体

### Requirement: 最小解释文案冻结

系统 SHALL 为两类派生读模型冻结最小解释文案，避免不同页面各写一套口径。

#### Scenario: `module_reuse_summary` 解释文案

- **WHEN** 页面展示某条 `module_reuse_summary`
- **THEN** 最小解释文案必须表达“该 Module 当前被多少个 Product 直接复用”
- **AND** 当 `reuse_product_count <= 1` 时，文案必须表达“当前尚未形成跨 Product 复用”

#### Scenario: `capability_summary` 解释文案

- **WHEN** 页面展示某条 `capability_summary`
- **THEN** 最小解释文案必须表达“该 capability 由多少个 Module 共同支撑”
- **AND** 当当前读取范围内没有任何参与 capability 聚合的 Module 时，必须展示成功空态解释文案，而不是留空白容器

### Requirement: 新鲜度与刷新语义冻结

系统 SHALL 将当前阶段复用感知的新鲜度语义冻结为“读取时反映最新已提交状态”，不引入异步离线刷新前提。

#### Scenario: 新鲜度判定

- **WHEN** `Product / Module / Repository / Decision` 的已提交状态发生变化并影响复用感知
- **THEN** 后续读取 `module_reuse_summary / capability_summary` 时必须反映最新已提交状态
- **AND** 当前阶段不得要求离线聚合作业、异步统计表刷新或后台批处理成功后才可见

### Requirement: 空状态与错误状态边界冻结

系统 SHALL 冻结复用感知的成功空态与读取失败态边界，避免页面把“暂无数据”和“读失败”混写。

#### Scenario: 成功空态

- **WHEN** 读取成功，但当前作用域下不存在复用数据或 capability 聚合结果
- **THEN** 页面必须展示成功空态文案
- **AND** 不得显示错误提示、重试按钮或红色错误容器

#### Scenario: 读取失败态

- **WHEN** `ReuseSummaryRead` 查询失败
- **THEN** 页面必须展示读取失败语义与重试入口
- **AND** 不得把读取失败伪装成“暂无复用反馈”

### Requirement: 复用感知读取 owner 与读取状态边界冻结

系统 SHALL 将复用感知的读取 owner 与读取状态边界与 `phase05` 已冻结的 Dashboard 读取模型显式对齐，避免实现阶段在 `Asset Feedback` 区块内产生两套冲突的读取状态语义。

#### Scenario: `ReuseSummaryRead` owner 单值化

- **WHEN** 接手者实现复用感知的后端读取
- **THEN** 必须以单数 `ReuseSummaryRead` 作为唯一 query owner
- **AND** 该 owner 承接 `ReadModuleReuseSummary` 与 `ReadCapabilitySummary` 两个读取动作
- **AND** 当前阶段不得为 `module_reuse_summary` 与 `capability_summary` 分别建立两个独立 owner 模块

#### Scenario: 复用感知不合并到 `FeedbackSignalRead`

- **WHEN** 接手者设计 Dashboard 的读取模型
- **THEN** 复用感知数据必须通过独立 `ReuseSummaryRead` query 承接
- **AND** 不得把 `module_reuse_summary / capability_summary` 合并到 `phase05` 已冻结的 `FeedbackSignalRead` 响应中
- **AND** `phase05` 已冻结的 `FeedbackSignalRead` 职责（服务 `Current Focus` 与 `Asset Feedback` 的反馈信号部分）不得被修改

#### Scenario: `Asset Feedback` 区块内读取状态分层

- **WHEN** `Asset Feedback` 区块同时展示反馈信号与复用快照
- **THEN** 区块内部必须分为两个独立读取状态的子区域：
  - 反馈信号子区域：由 `FeedbackSignalRead` 服务，沿用 `phase05` 已冻结的局部失败语义
  - 复用快照子区域：由 `ReuseSummaryRead` 服务，拥有独立的成功 / 失败 / loading 状态
- **AND** `ReuseSummaryRead` 失败不得影响 `FeedbackSignalRead` 的结果展示
- **AND** `FeedbackSignalRead` 失败不得影响 `ReuseSummaryRead` 的结果展示
- **AND** 两个子区域的失败语义各自独立，不互为前提

## MODIFIED Requirements

### Requirement: Dashboard Asset Feedback 区块职责

`Dashboard` 中现有 `Asset Feedback` 区块在 `phase06-04` 中 SHALL 继续承接反馈主线，但额外承担复用感知的最小派生反馈挂接位。该扩展必须与 `phase05` 已冻结的 `FeedbackSignalRead` 读取状态边界显式对齐。

#### Scenario: Asset Feedback 区块职责扩展

- **WHEN** 接手者扩展 `Asset Feedback` 区块
- **THEN** 必须继续保留 `phase05` 已冻结的反馈信号展示职责（`FeedbackSignalRead` 服务 `Current Focus` 与 `Asset Feedback` 的反馈信号部分，包括 `product_asset_coverage` 代表性缺口项，最多 `3` 条）
- **AND** 可以在该区块内部追加 `Top Reused Modules / Capability Snapshot`（由独立 `ReuseSummaryRead` query 服务）
- **AND** 反馈信号子区域与复用快照子区域的读取状态必须分层独立（见“复用感知读取 owner 与读取状态边界冻结” Requirement）
- **AND** 不得因此重写 `phase05` 已冻结的 Dashboard 一级结构与返回路径语义
- **AND** 不得把 `phase05` 已冻结的 `FeedbackSignalRead` 响应结构扩写为同时承载 `module_reuse_summary / capability_summary` 数据

## REMOVED Requirements

### Requirement: 为复用感知新增独立一级“复用中心”页面

**Reason**: 当前阶段目标是让复用反馈尽早进入既有页面可见能力，而不是再造新的一级工作台或重实体中心。

**Migration**: phase06 后续实现统一改为：Dashboard 在 `Asset Feedback` 内追加复用快照，`Module Detail / Product Detail` 在各自既有详情主区追加最小复用感知子区块。
