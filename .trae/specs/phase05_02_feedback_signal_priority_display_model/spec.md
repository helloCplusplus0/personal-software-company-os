# Phase05-02 Feedback 信号模板、优先级与最小展示模型 Spec

## Why

`phase05-01` 已经把 `Dashboard Home` 的页面边界、四区块归属、首页入口与点击热区冻结成单值结论，但 `Dashboard + Feedback` 要真正进入后续 `/spec`、合同与实现，还必须把 `pending_decision_signals`、`product_asset_coverage`、`recent_activity_feed` 与 `dashboard_overview` 的消费语义、最小字段模板、优先级与排序口径写成单值结论。否则后续前端展示、后端聚合、`.proto` 合同与验收会继续在“什么算反馈信号”“怎样排序”“哪些项进入主队列”“活动流按什么时间排序”之间漂移。

## What Changes

- 冻结 `Feedback` 当前阶段的单值解释为派生信号层
- 冻结 `pending_decision_signals` 的 Dashboard 消费语义
- 冻结 `product_asset_coverage` 的 Dashboard 消费语义与空状态结构
- 冻结 `product missing both bindings` 在 `product_asset_coverage` 中的独立读模型语义
- 冻结统一 `Feedback Signal Card` 的最小字段模板
- 冻结反馈信号的优先级、排序前提与空状态语义
- 冻结 `Current Focus / Next Action` 与 `Asset Feedback` 的最大展示数量
- 冻结 `recent_activity_feed` 与 `dashboard_overview` 的最小展示模型
- 冻结 `recent_activity_feed` 的显式排序字段与时间语义

## Impact

- Affected specs: `phase05_dashboard_feedback_foundation`
- Affected code: 后续 `frontend/src/features/dashboard/` 中的反馈卡片、活动流与概览展示模型，后续 `backend/` 中的 Dashboard 聚合读取与派生查询，后续 `proto/` 中的 `FeedbackSignal`、`RecentActivityItem` 与概览读取消息

## ADDED Requirements

### Requirement: Feedback 最小语义冻结

系统 SHALL 将 `Feedback` 在当前阶段单值化为基于既有结构化对象与绑定关系派生出的“行动信号层”，而不是独立重实体、人工录入对象或完整运营系统。

#### Scenario: 判断 Feedback 当前阶段是什么

- **WHEN** 后续 `/spec`、前端展示、后端聚合或 `.proto` 合同讨论 `Feedback`
- **THEN** 必须将其解释为基于 `Decision / Product / Repository / Module Binding` 派生出的行动信号
- **AND** 每条信号至少回答“当前哪里不完整”或“现在该去哪里补”
- **AND** 不得在当前阶段扩写为独立 `Feedback` 表、人工录入流或复杂评分系统

### Requirement: pending_decision_signals Dashboard 消费语义冻结

系统 SHALL 将 `pending_decision_signals` 冻结为进入 `Current Focus / Next Action` 主队列的高优先级反馈来源。

#### Scenario: pending_decision_signals 进入主队列

- **WHEN** Dashboard 读取到 `pending_decision_signals`
- **THEN** 它们必须先归一化为统一的 `Feedback Signal Card`
- **AND** 必须进入 `Current Focus / Next Action` 主队列
- **AND** 不得降级到 `Asset Feedback` 或 `Recent Activity`

#### Scenario: pending_decision_signals 空状态

- **WHEN** 当前系统不存在待决策信号
- **THEN** Dashboard 必须返回空列表语义
- **AND** 不得将“无待决策信号”解释为读取失败或资源不存在

### Requirement: product_asset_coverage Dashboard 消费语义冻结

系统 SHALL 将 `product_asset_coverage` 冻结为 `Dashboard` 中与资产缺口相关的反馈来源，并明确其同时服务 `Current Focus / Next Action` 与 `Asset Feedback` 两个展示层级。

#### Scenario: product_asset_coverage 进入主队列

- **WHEN** `product_asset_coverage` 产生高优先级缺口项
- **THEN** 缺口项必须先归一化为统一的 `Feedback Signal Card`
- **AND** 必须按冻结后的优先级参与 `Current Focus / Next Action` 主队列排序

#### Scenario: product_asset_coverage 进入补充摘要

- **WHEN** Dashboard 展示 `Asset Feedback`
- **THEN** 该区块只承接 `product_asset_coverage` 的补充摘要结果
- **AND** 最多展示 `3` 条代表性缺口项
- **AND** 不得形成与 `Current Focus / Next Action` 并列的第二条全局优先级队列

#### Scenario: product_asset_coverage 空状态

- **WHEN** 当前系统不存在资产缺口（所有产品均已完成模块与仓库双绑定）或当前不存在代表性缺口项
- **THEN** `product_asset_coverage` 必须返回完整的读模型结构，不得因无缺口而省略字段或返回空对象
- **AND** 以下计数必须显式为 `0`：缺少模块与仓库双绑定的产品数、缺少仓库绑定但已有模块绑定的产品数、缺少模块绑定但已有仓库绑定的产品数
- **AND** 代表性缺口项列表必须为空列表，而不是 `null` 或省略字段
- **AND** 此结果必须以成功语义返回（如 HTTP 200），不得映射为资源不存在或读取失败
- **AND** 此类“成功空态”必须与局部读取失败显式区分：成功空态返回完整结构 + 零值计数 + 空代表项列表；局部失败返回局部错误标记并允许局部重试，两者不得互相替代

### Requirement: missing both bindings 独立读模型语义冻结

系统 SHALL 在 `product_asset_coverage` 中显式冻结 `product missing both bindings` 的独立计数与代表项语义，不得再把它隐含地拆散到两个单缺口类型中。

#### Scenario: 判断双缺口产品

- **WHEN** 某个 `Product` 同时缺少 `Repository` 绑定与 `Module` 绑定
- **THEN** 它必须被归类为 `product missing both bindings`
- **AND** 不得同时重复计入 `product missing repository binding` 与 `product missing module binding` 两类代表项语义

#### Scenario: 双缺口计数与代表项

- **WHEN** Dashboard 读取 `product_asset_coverage`
- **THEN** 读模型至少必须显式承接“缺少模块与仓库双绑定的产品数”
- **AND** 代表性缺口项必须允许出现 `product missing both bindings` 这一独立 `signal_code`

### Requirement: Feedback Signal Card 最小字段模板冻结

系统 SHALL 将统一 `Feedback Signal Card` 的最小字段模板冻结为当前阶段唯一反馈卡片模型。

#### Scenario: 判断 Feedback Signal Card 最小字段

- **WHEN** 后续前端组件、后端 DTO 或 `.proto` 合同讨论反馈卡片字段
- **THEN** 每条卡片至少必须承接 `signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label`
- **AND** 不得减少上述字段，导致卡片既无法解释缺口也无法承接跳转
- **AND** 当前阶段不得额外引入 `score / trend / external_metric / recommendation_reason`

### Requirement: 反馈信号优先级冻结

系统 SHALL 将当前阶段反馈信号的优先级顺序冻结为唯一单值结论。

#### Scenario: 判断优先级顺序

- **WHEN** Dashboard 对归一化后的反馈信号进行排序
- **THEN** 必须按以下顺序处理：`pending_decision_signals > product missing both bindings > product missing repository binding > product missing module binding`
- **AND** 不得在当前阶段引入第二套优先级口径

### Requirement: 反馈信号排序前提冻结

系统 SHALL 将反馈信号同优先级内的排序前提冻结为“最近需要处理时间优先”，并定义缺省回退规则。

#### Scenario: 同优先级排序

- **WHEN** 两条或多条反馈信号具有相同优先级
- **THEN** 系统必须优先按“最近需要处理时间”排序
- **AND** 若当前读模型暂未提供显式处理时间，则回退为 `created_at DESC`

### Requirement: Current Focus 与 Asset Feedback 展示数量冻结

系统 SHALL 将 `Current Focus / Next Action` 与 `Asset Feedback` 的最大展示数量冻结为当前阶段唯一展示上限。

#### Scenario: Current Focus 展示上限

- **WHEN** Dashboard 展示 `Current Focus / Next Action`
- **THEN** 主队列最多展示 `5` 条反馈卡片
- **AND** 多出的信号不得在当前阶段继续扩写成第二条主队列

#### Scenario: Asset Feedback 展示上限

- **WHEN** Dashboard 展示 `Asset Feedback`
- **THEN** 该区块最多展示 `3` 条代表性缺口项
- **AND** 该区块只作为补充摘要，而不是主行动区

### Requirement: recent_activity_feed 最小展示模型冻结

系统 SHALL 将 `recent_activity_feed` 冻结为独立活动流读模型，而不是反馈信号的另一种表现形式。

#### Scenario: 活动流最小展示模型

- **WHEN** Dashboard 展示 `recent_activity_feed`
- **THEN** 每条活动项至少必须承接 `activity_type / activity_at / target_type / target_id / target_label`
- **AND** 当前阶段最多展示 `10` 条活动项
- **AND** 活动项类型至少承接 `Module / Release / Product / Repository / Decision / product_module_binding / product_repository_binding / module_repository_binding`

#### Scenario: 活动流与反馈队列关系

- **WHEN** Dashboard 同时展示反馈信号与最近活动
- **THEN** `recent_activity_feed` 必须作为独立活动流区块存在
- **AND** 不得进入反馈优先级队列
- **AND** 不得与 `Current Focus / Next Action` 混合排序

### Requirement: recent_activity_feed 排序字段与时间语义冻结

系统 SHALL 将 `recent_activity_feed` 的排序字段冻结为显式活动时间字段，并禁止依赖隐式实现假设推断活动顺序。

#### Scenario: 活动流排序字段

- **WHEN** Dashboard 对最近活动进行排序
- **THEN** 必须以显式活动时间字段（如 `activity_at`）为唯一排序基准
- **AND** 默认按活动时间倒序展示
- **AND** 不得依赖隐式 `created_at` 或不透明数据库默认值推断最近活动顺序

### Requirement: dashboard_overview 最小展示模型冻结

系统 SHALL 将 `dashboard_overview` 冻结为概览辅助读模型，而不是反馈信号来源。

#### Scenario: 概览最小展示字段

- **WHEN** Dashboard 展示 `dashboard_overview`
- **THEN** 概览至少必须承接 `module_count / product_count / repository_count / decision_count / product_with_repository_count / product_with_module_count`
- **AND** `dashboard_overview` 只服务概览卡片区块
- **AND** 不得进入反馈优先级队列

## MODIFIED Requirements

### Requirement: Dashboard 四类读取的职责解释

`dashboard_overview / pending_decision_signals / product_asset_coverage / recent_activity_feed` 在当前阶段 SHALL 不再被解释为“都属于 Dashboard 的一组并列聚合读”，而必须被解释为职责不同、排序口径不同、进入区块不同的四类读取。

#### Scenario: 判断四类读取职责

- **WHEN** 后续 `/spec`、前端展示、后端聚合或合同设计讨论四类 Dashboard 读取
- **THEN** 必须理解为：`dashboard_overview` 承接概览、`pending_decision_signals` 承接高优先级决策信号、`product_asset_coverage` 承接资产缺口信号、`recent_activity_feed` 承接独立活动流
- **AND** 不得把四类读取重新混合为统一排序队列

### Requirement: Asset Feedback 与 Current Focus 的关系解释

`Asset Feedback` 在当前阶段 SHALL 不再被解释为另一条与 `Current Focus / Next Action` 并列的主行动队列，而必须被解释为 `product_asset_coverage` 的补充摘要区块。

#### Scenario: 判断 Asset Feedback 角色

- **WHEN** 后续设计或实现讨论 `Asset Feedback`
- **THEN** 必须将其理解为资产缺口摘要区块
- **AND** 不得把它重写成第二条主行动队列

## REMOVED Requirements

### Requirement: 将 recent_activity_feed 作为反馈优先级输入

**Reason**: `recent_activity_feed` 的职责是承接结构化资产变化的活动流，而不是判断“下一步先去哪里补”的反馈信号；若把活动流并入反馈队列，会重新引入 Dashboard 无行动价值统计页与混合排序歧义。

**Migration**: 当前阶段统一将 `recent_activity_feed` 作为独立活动流区块处理，并以显式活动时间字段排序；如后续需要从活动流派生新的反馈信号，应在新的 `phase / fix / audit` 中单独冻结派生规则，而不是直接复用活动流本身。
