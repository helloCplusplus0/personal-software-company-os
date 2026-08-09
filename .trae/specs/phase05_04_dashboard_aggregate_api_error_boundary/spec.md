# Phase05-04 Dashboard 聚合读范围、接口边界与错误语义前提 Spec

## Why

`phase05-01` 已冻结 `Dashboard Home` 页面边界，`phase05-02` 已冻结反馈信号与活动流展示模型，`phase05-03` 已冻结跳转与返回上下文，但“页面长什么样、点到哪里去”还不等于“后台最少要读什么、接口边界停在哪、空态和错误怎么解释”。如果不把 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 的最小聚合读范围、错误语义与主 CTA 优先级矩阵写成单值结论，后续前端状态设计、后端模块分组、`.proto` 合同与验收就会继续在“哪些数据算主聚合”“局部失败是否拖垮整页”“空系统和非空缺口如何区分”之间漂移。

> 阶段分工约束：本规格只冻结 `phase05-04` 范围内的聚合读范围、最小接口边界与错误语义前提。前端页面级状态模型与交互流由 `phase05-06` 承接；后端模块边界、接口分组、query service owner 与 DTO 映射策略由 `phase05-07` 承接；`.proto` 服务接口命名、消息结构、包名版本与 `chi + JSON HTTP` 显式映射由 `phase05-08` 承接，不在本规格中提前冻结。

## What Changes

- 冻结 `Dashboard` 当前阶段的最小聚合读范围，只承接 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead`
- 冻结概览读取、信号读取与最近活动读取的最小接口边界
- 冻结空系统、无信号、无活动、局部失败与整页失败的错误语义前提
- 冻结“冷启动空系统”与“已有结构化资产但仍无 Module”的区分规则
- 冻结 Dashboard 主 CTA 的唯一优先级矩阵与单主 CTA 约束
- 明确当前阶段不提前冻结外部埋点、趋势分析、通知中心与新的业务写入接口

## Impact

- Affected specs: `phase05_dashboard_feedback_foundation`
- Affected code: 后续 `frontend/src/features/dashboard/*` 的整页读取承接、空态/局部错误/整页错误解释与 CTA 命中逻辑
- Affected code: 后续 `backend/` 中 Dashboard 聚合查询的最小数据范围校验、局部失败归类与错误哨兵值设计
- Affected code: 后续 `proto/` 中 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 的消息范围与错误语义前提

## ADDED Requirements

### Requirement: Dashboard 当前阶段最小聚合读范围冻结

系统 SHALL 将 `Dashboard Home` 当前阶段最小聚合读范围冻结为三个只读聚合入口，不得在当前阶段扩写第二套 Dashboard 接口主线。

#### Scenario: 判断当前阶段最小聚合读集合

- **WHEN** 后续 `/spec`、前端展示、后端实现或 `.proto` 合同讨论 Dashboard 当前阶段最少需要哪些读取
- **THEN** 必须得到 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 三个聚合读取的单值结论
- **AND** `DashboardOverviewRead` 只承接概览与系统状态判定前提
- **AND** `FeedbackSignalRead` 只承接归一化后的反馈信号与资产缺口补充摘要
- **AND** `RecentActivityRead` 只承接独立活动流
- **AND** 不得在当前阶段并列增加趋势分析、外部埋点、通知中心或导出接口

#### Scenario: 判断 Dashboard 当前阶段是否承接写入

- **WHEN** 后续 `/spec` 或实现讨论 Dashboard 接口范围
- **THEN** Dashboard 当前阶段不得承接任何新的业务写入接口
- **AND** 不得把 `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository`、`RecordDecision` 或其他写动作重新塞回 Dashboard

### Requirement: DashboardOverviewRead 最小接口边界冻结

系统 SHALL 将 `DashboardOverviewRead` 的最小接口边界冻结为只服务概览卡片与空系统判定的主聚合读取。

#### Scenario: 判断 DashboardOverviewRead 输出范围

- **WHEN** 后续 `/spec`、前端消费、后端 DTO 或 `.proto` 合同讨论 `DashboardOverviewRead` 的输出范围
- **THEN** 它至少必须承接 `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count`
- **AND** 不得把 `pending_decision_signals`、`product_asset_coverage` 或 `recent_activity_feed` 的结果直接并入该读取
- **AND** `DashboardOverviewRead` 必须足以支撑空系统、非空但无 `Module`、缺 `Product`、缺 `Repository` 的状态判定

#### Scenario: 判断 DashboardOverviewRead 输入边界

- **WHEN** 后续实现讨论 `DashboardOverviewRead` 是否接受业务筛选、分页、排序或时间范围参数
- **THEN** 当前阶段不得为该读取引入 `queryText / statusFilter / page / pageSize / sort / dateRange`
- **AND** 当前阶段只允许它作为 `Dashboard Home` 默认主聚合读取存在

### Requirement: FeedbackSignalRead 最小接口边界冻结

系统 SHALL 将 `FeedbackSignalRead` 的最小接口边界冻结为“统一反馈信号主队列 + 资产缺口补充摘要”的附属聚合读取。

#### Scenario: 判断 FeedbackSignalRead 输入来源

- **WHEN** 后续 `/spec`、后端实现或 `.proto` 合同讨论 `FeedbackSignalRead` 的上游来源
- **THEN** 它只允许消费 `pending_decision_signals` 与 `product_asset_coverage`
- **AND** 不得把 `recent_activity_feed` 混入信号读取
- **AND** 不得在当前阶段临时发明第二套隐式反馈数据源

#### Scenario: 判断 FeedbackSignalRead 输出范围

- **WHEN** 后续前端消费、后端 DTO 或 `.proto` 合同讨论 `FeedbackSignalRead` 的输出范围
- **THEN** 它至少必须承接：
  - 进入 `Current Focus / Next Action` 的统一反馈卡片列表（最多 `5` 条）
  - `Asset Feedback` 的补充摘要结果（最多 `3` 条代表性缺口项）
- **AND** 反馈卡片必须继续遵守 `phase05-02` 已冻结的最小字段模板
- **AND** 主队列优先级必须继续遵守 `pending_decision_signals > product missing both bindings > product missing repository binding > product missing module binding`
- **AND** 不得在当前阶段为该读取增加分页、手动排序切换、信号家族筛选或批量操作参数

### Requirement: RecentActivityRead 最小接口边界冻结

系统 SHALL 将 `RecentActivityRead` 的最小接口边界冻结为独立活动流读取，不与反馈优先级队列混合。

#### Scenario: 判断 RecentActivityRead 输出范围

- **WHEN** 后续 `/spec`、前端消费、后端 DTO 或 `.proto` 合同讨论 `RecentActivityRead` 的输出范围
- **THEN** 它至少必须承接 `activity_type / activity_at / target_type / target_id / target_label`
- **AND** 当前阶段最多返回 `10` 条活动项
- **AND** 活动项类型至少承接 `module / release / product / repository / decision / product_module_binding / product_repository_binding / module_repository_binding`
- **AND** 活动流必须只按显式活动时间倒序排序
- **AND** 不得把活动流并入 `FeedbackSignalRead`

#### Scenario: 判断 RecentActivityRead 输入边界

- **WHEN** 后续实现讨论 `RecentActivityRead` 是否接受类型筛选、时间范围、分页或排序切换
- **THEN** 当前阶段不得引入 `activityTypeFilter / dateRange / page / pageSize / sort`
- **AND** 当前阶段只允许它作为 `Dashboard Home` 的独立附属聚合读取存在

### Requirement: 空系统与非空缺口系统状态语义冻结

系统 SHALL 将冷启动空系统与“已有结构化资产但仍无 Module”的系统状态语义冻结为两个不同的成功状态。

#### Scenario: 冷启动空系统语义

- **WHEN** `module_count = 0 && product_count = 0 && repository_count = 0 && decision_count = 0`
- **THEN** Dashboard 必须进入冷启动空系统状态
- **AND** 此状态必须以成功语义返回，而不是整页失败
- **AND** 主 CTA 必须指向 `Module Registry / Create`

#### Scenario: 非空但无 Module 语义

- **WHEN** `module_count = 0 && (product_count > 0 || repository_count > 0 || decision_count > 0)`
- **THEN** Dashboard 必须进入“已有结构化资产但仍无 Module”的非空缺口状态
- **AND** 该状态不得与冷启动空系统混同
- **AND** 此状态必须以成功语义返回，而不是整页失败
- **AND** 主 CTA 仍必须指向 `Module Registry / Create`

### Requirement: 空信号、空活动与成功空态语义冻结

系统 SHALL 将无反馈信号、无活动与局部成功空态解释为成功读取结果，不得误判为资源不存在或接口失败。

#### Scenario: 无待处理反馈信号

- **WHEN** 当前系统不存在 `pending_decision_signals`，且 `product_asset_coverage` 不产生进入主队列的缺口项
- **THEN** `FeedbackSignalRead` 必须返回成功语义
- **AND** `Current Focus / Next Action` 必须返回空列表或零项语义
- **AND** 不得将“无信号”解释为读取失败或资源不存在

#### Scenario: Asset Feedback 成功空态

- **WHEN** 当前系统不存在资产缺口或不存在代表性缺口项
- **THEN** `FeedbackSignalRead` 仍必须返回完整补充摘要结构
- **AND** 缺口计数必须显式为 `0`
- **AND** 代表性缺口项必须为空列表，而不是 `null` 或省略字段
- **AND** 不得把该结果解释为局部失败

#### Scenario: Recent Activity 成功空态

- **WHEN** 当前系统暂无可展示的最近活动
- **THEN** `RecentActivityRead` 必须返回空列表语义
- **AND** 不得映射为资源不存在或接口错误

#### Scenario: 无缺口且无活动的非空成功态

- **WHEN** `DashboardOverviewRead` 显示系统已存在结构化资产
- **AND** `FeedbackSignalRead` 成功且不存在任何反馈缺口
- **AND** `RecentActivityRead` 成功但返回空列表
- **THEN** Dashboard 必须进入非空成功态
- **AND** 该状态不得与冷启动空系统混同
- **AND** 不得因“暂无活动”而回退到创建导向主 CTA
- **AND** 当前阶段不展示强制主 CTA

### Requirement: 局部失败与整页失败语义冻结

系统 SHALL 将 Dashboard 的局部失败与整页失败前提冻结为单值结论，避免后续实现阶段对失败范围做两套解释。

#### Scenario: 主聚合失败触发整页失败

- **WHEN** `DashboardOverviewRead` 失败
- **THEN** Dashboard 必须进入整页失败语义
- **AND** 因为当前阶段空系统判定、系统状态区分与主 CTA 优先级都依赖该主聚合读取
- **AND** 不得把该失败伪装为“成功但空系统”

#### Scenario: 附属聚合失败不拖垮整页

- **WHEN** `DashboardOverviewRead` 成功，但 `FeedbackSignalRead` 失败
- **THEN** Dashboard 只允许 `Current Focus / Next Action` 与 `Asset Feedback` 进入局部失败语义
- **AND** 不得强制整页失败
- **AND** 若 `DashboardOverviewRead` 的结果未命中主 CTA 顺序 1-4，则当前阶段必须抑制强制主 CTA
- **AND** 不得在未知反馈结果下伪造 5-8 任一反馈型主 CTA
- **AND** 局部恢复动作必须降级为反馈区块内的重试入口，而不是整页主 CTA

- **WHEN** `DashboardOverviewRead` 成功，但 `RecentActivityRead` 失败
- **THEN** Dashboard 只允许 `Recent Activity` 区块进入局部失败语义
- **AND** 不得强制整页失败

#### Scenario: 非法参数错误语义

- **WHEN** 当前阶段任一 Dashboard 聚合读取收到超出冻结边界的非法参数
- **THEN** 系统必须返回明确校验失败语义
- **AND** 不得静默忽略非法参数后继续返回不确定结果

### Requirement: Dashboard 主 CTA 优先级矩阵冻结

系统 SHALL 将 Dashboard 主 CTA 的命中顺序冻结为唯一优先级矩阵，且同一时刻只允许存在一个主 CTA。

#### Scenario: 主 CTA 命中顺序

- **WHEN** Dashboard 基于成功的主聚合与附属聚合结果决定主 CTA
- **THEN** 必须按以下顺序命中：
  1. `module_count = 0 && product_count = 0 && repository_count = 0 && decision_count = 0`
  2. `module_count = 0 && (product_count > 0 || repository_count > 0 || decision_count > 0)`
  3. `module_count > 0 && product_count = 0`
  4. `module_count > 0 && product_count > 0 && repository_count = 0`
  5. 存在 `pending_decision_signals`
  6. 存在 `product missing both bindings`
  7. 存在 `product missing repository binding`
  8. 存在 `product missing module binding`
  9. 无缺口且有活动数据
- **AND** 顺序 1-4 命中时主 CTA 分别指向 `Module Registry / Create`、`Module Registry / Create`、`Product Registry / Create`、`Repository Binding / Create`
- **AND** 顺序 5-8 命中时主 CTA 必须回到 `phase05-03` 已冻结的 canonical 落点
- **AND** 顺序 9 命中时 Dashboard 进入“系统已就绪”中性状态，不再展示强制主 CTA

#### Scenario: 无缺口且无活动时的 CTA

- **WHEN** 系统不存在反馈缺口
- **AND** `RecentActivityRead` 返回空列表
- **AND** 当前系统不属于冷启动空系统，也不属于“已有结构化资产但仍无 Module”
- **THEN** Dashboard 必须进入非空中性状态
- **AND** 当前阶段不展示强制主 CTA
- **AND** 不得为了保持页面动作感而发明新的创建导向 CTA

#### Scenario: 单主 CTA 约束

- **WHEN** 同时存在多类缺口或多类候选 CTA
- **THEN** Dashboard 同一时刻只允许展示一个主 CTA
- **AND** 其他动作必须降级到区块内次级入口
- **AND** 不得并排展示多个同级主 CTA

### Requirement: 非目标接口边界冻结

系统 SHALL 明确当前阶段不提前冻结超出 Dashboard 最小反馈闭环的聚合接口与外部连接接口。

#### Scenario: 判断非目标接口边界

- **WHEN** 后续 `/spec`、实现设计或代码实现讨论 Dashboard 接口范围
- **THEN** 不得提前冻结外部埋点、趋势分析、通知中心、消息中心、导出接口或外部数据接入接口
- **AND** 不得提前冻结 AI 自动建议、自动评分或自动补全接口
- **AND** 不得提前冻结超出当前阶段的分页、复杂检索、筛选器矩阵或多页面聚合分析接口

## MODIFIED Requirements

### Requirement: Contract First 当前阶段解释

`Contract First` 在 `phase05` 当前阶段 SHALL 被进一步解释为“先冻结 Dashboard 三类聚合读范围、接口边界与错误语义前提，再进入后续 `.proto` 合同设计”，而不是在范围未收敛前提前扩写完整服务矩阵。

#### Scenario: Contract First 承接方式

- **WHEN** 后续 `/spec`、实现设计或 `.proto` 合同讨论 `Dashboard + Feedback` 的 Contract First
- **THEN** 必须优先沿用本规格已冻结的三类聚合读取范围、空态与错误语义前提
- **AND** 不得提前发明与 `.proto` 单一合同源并列的第二套接口定义体系

## REMOVED Requirements

### Requirement: 外部埋点、趋势分析与通知中心接口

**Reason**: `phase05-04` 当前目标是建立 `Dashboard + Feedback` 的最小聚合读闭环，而不是提前完成外部遥测、BI 趋势或通知中心的完整接口规划。

**Migration**: 若后续确需引入外部埋点、趋势分析、通知中心或消息中心接口，必须进入新的 `phase / fix / audit` 任务重新单值化，不在 `phase05-04` 当前规格中处理。
