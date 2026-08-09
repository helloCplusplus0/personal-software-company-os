# Phase05-03 Dashboard 跳转目标、来源上下文与返回路径 Spec

## Why

`phase05-01` 已冻结 `Dashboard Home` 的页面边界、区块归属、首页入口与点击热区，`phase05-02` 已冻结反馈信号、活动流与最小展示模型，但“可以点”还不等于“点了之后去哪里、带什么上下文、怎么返回”。如果不把 `Dashboard` 到既有 canonical owner 页面的跳转矩阵、`fromDashboard` 来源标记、`dashboardSection` 区块标记与返回路径恢复规则写成单值结论，后续前端实现、合同设计与验收会继续在“返回上一页”“回 Dashboard”“写入成功后是否回流”之间产生第二套语义。

## What Changes

- 冻结 `dashboard_overview` 四类独立概览卡片的 canonical 跳转目标
- 冻结 `Current Focus / Next Action`、`Asset Feedback` 与空状态 CTA 的目标页面
- 冻结 `recent_activity_feed` 中 `Module / Release / Product / Repository / Decision / Binding` 各活动类型的单值落点
- 冻结 `fromDashboard=true`、`dashboardSection`、`dashboardReturnTo=/dashboard` 的参数语义与允许取值
- 冻结从 Dashboard 跳出后的主动返回、刷新恢复、多跳进入与 Create 页回流场景下的上下文承接规则
- 明确 `fromDashboard` 是 Dashboard 外层来源标记，不覆盖 `phase04` 已冻结的页面原生来源模型
- 明确 Dashboard 只承接读与跳转，不形成“直接补录/直接绑定”的影子工作台

## Impact

- Affected specs: `phase05_dashboard_feedback_foundation`
- Affected code: 后续 `frontend/src/routes/dashboard/*`、`frontend/src/routes/modules/*`、`frontend/src/routes/products/*`、`frontend/src/routes/repositories/*`、`frontend/src/routes/decisions/*` 的导航参数、返回逻辑与搜索参数恢复
- Affected code: 后续 `frontend/src/features/dashboard/*` 的卡片点击逻辑、空状态 CTA 与最近活动项导航承接
- Affected code: 后续 `proto/` 与 `backend/` 中的 `FeedbackSignal`、`RecentActivityItem`、Dashboard 聚合 DTO 与跳转字段语义

## ADDED Requirements

### Requirement: Dashboard 概览卡片跳转目标冻结

系统 SHALL 将 `dashboard_overview` 中当前阶段允许点击的概览卡片跳转目标冻结为单值结论。

#### Scenario: module_count 概览卡片跳转

- **WHEN** 用户点击 `dashboard_overview` 中的 `module_count` 概览卡片
- **THEN** 必须跳转到 `Module Registry / List`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=overview`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: product_count 概览卡片跳转

- **WHEN** 用户点击 `dashboard_overview` 中的 `product_count` 概览卡片
- **THEN** 必须跳转到 `Product Registry / List`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=overview`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: repository_count 概览卡片跳转

- **WHEN** 用户点击 `dashboard_overview` 中的 `repository_count` 概览卡片
- **THEN** 必须跳转到 `Repository Binding / List`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=overview`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: decision_count 概览卡片跳转

- **WHEN** 用户点击 `dashboard_overview` 中的 `decision_count` 概览卡片
- **THEN** 必须跳转到 `Decision Center / List`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=overview`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: 派生覆盖率指标不形成独立跳转

- **WHEN** `dashboard_overview` 展示 `product_with_repository_count` 或 `product_with_module_count`
- **THEN** 它们在当前阶段只作为辅助指标展示
- **AND** 不得临时发明新的 `Product List` 筛选态
- **AND** 不得为它们分配独立 canonical 跳转目标

### Requirement: Feedback Signal Card 跳转矩阵冻结

系统 SHALL 将 `Current Focus / Next Action` 与 `Asset Feedback` 中反馈卡片的 canonical 落点冻结为既有 owner 页面。

#### Scenario: 单项决策信号跳转

- **WHEN** `pending_decision_signals` 中某条信号已绑定具体 `decision_id`
- **THEN** 该卡片必须跳转到对应 `Decision Detail`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=current-focus`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: 聚合决策信号跳转

- **WHEN** `pending_decision_signals` 中某条信号未绑定单一 `decision_id`
- **THEN** 该卡片必须跳转到 `Decision Center / List`
- **AND** 不得伪造不存在的 `Decision Detail`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=current-focus`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: 产品资产缺口信号跳转

- **WHEN** `product missing both bindings`、`product missing repository binding` 或 `product missing module binding` 作为反馈卡片出现
- **THEN** 它们都必须跳转到对应 `Product Detail`
- **AND** 该跳转必须以 `product_id` 作为目标身份参数
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`
- **AND** 若卡片位于 `Current Focus / Next Action`，则 `dashboardSection=current-focus`
- **AND** 若卡片位于 `Asset Feedback`，则 `dashboardSection=asset-feedback`

### Requirement: Recent Activity 活动类型落点冻结

系统 SHALL 将 `recent_activity_feed` 的活动类型落点冻结为单值结论，不得保留笼统 `binding` 歧义。

#### Scenario: 结构化对象活动类型跳转

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `module`
- **THEN** 活动项必须跳转到对应 `Module Detail`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=recent-activity`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `product`
- **THEN** 活动项必须跳转到对应 `Product Detail`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=recent-activity`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `repository`
- **THEN** 活动项必须跳转到对应 `Repository Binding Detail / Workspace`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=recent-activity`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `decision`
- **THEN** 活动项必须跳转到对应 `Decision Detail`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=recent-activity`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: Release 活动回落到 Module Detail

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `release`
- **THEN** 活动项必须统一跳转到所属 `Module Detail`
- **AND** 不得在当前阶段新增独立 `Release Detail` 页面
- **AND** 跳转身份必须以所属 `module_id` 为准，而不是临时拼出新的 `release` 页面路由
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=recent-activity`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

#### Scenario: 三类 Binding 活动落点

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `product_module_binding`
- **THEN** 活动项必须跳转到对应 `Product Detail`
- **AND** 目标身份必须为 `product_id`
- **AND** `module_id` 只作为展示上下文，不作为导航 owner

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `product_repository_binding`
- **THEN** 活动项必须跳转到对应 `Repository Binding Detail / Workspace`
- **AND** 目标身份必须为 `repository_id`
- **AND** `product_id` 只作为展示上下文，不作为导航 owner

- **WHEN** `recent_activity_feed` 的 `activity_type` 为 `module_repository_binding`
- **THEN** 活动项必须跳转到对应 `Repository Binding Detail / Workspace`
- **AND** 目标身份必须为 `repository_id`
- **AND** `module_id` 只作为展示上下文，不作为导航 owner

- **AND** 上述三类 Binding 活动都必须携带 `fromDashboard=true`
- **AND** 都必须携带 `dashboardSection=recent-activity`
- **AND** 都必须携带 `dashboardReturnTo=/dashboard`

### Requirement: 空状态 CTA 跳转目标与优先级冻结

系统 SHALL 将 Dashboard 空状态与主 CTA 的跳转目标冻结为共享基线中唯一优先级矩阵。

#### Scenario: 冷启动与缺口态 CTA

- **WHEN** `module_count = 0 && product_count = 0 && repository_count = 0 && decision_count = 0`
- **THEN** Dashboard 主 CTA 必须跳转到 `Module Registry / Create`
- **AND** 必须携带 `fromDashboard=true`
- **AND** 必须携带 `dashboardSection=empty-state`
- **AND** 必须携带 `dashboardReturnTo=/dashboard`

- **WHEN** `module_count = 0 && (product_count > 0 || repository_count > 0 || decision_count > 0)`
- **THEN** Dashboard 主 CTA 仍必须跳转到 `Module Registry / Create`
- **AND** 不得改写为并列多个主 CTA

- **WHEN** `module_count > 0 && product_count = 0`
- **THEN** Dashboard 主 CTA 必须跳转到 `Product Registry / Create`

- **WHEN** `module_count > 0 && product_count > 0 && repository_count = 0`
- **THEN** Dashboard 主 CTA 必须跳转到 `Repository Binding / Create`

#### Scenario: 高优先级反馈主 CTA

- **WHEN** 存在 `pending_decision_signals`
- **THEN** Dashboard 主 CTA 必须指向最高优先级决策信号落点
- **AND** 若该信号绑定具体 `decision_id`，则跳转 `Decision Detail`
- **AND** 若该信号未绑定具体 `decision_id`，则跳转 `Decision Center / List`

- **WHEN** 不存在待决策信号且存在 `product missing both bindings`
- **THEN** Dashboard 主 CTA 必须跳转到对应 `Product Detail`

- **WHEN** 不存在更高优先级缺口且存在 `product missing repository binding`
- **THEN** Dashboard 主 CTA 必须跳转到对应 `Product Detail`

- **WHEN** 不存在更高优先级缺口且存在 `product missing module binding`
- **THEN** Dashboard 主 CTA 必须跳转到对应 `Product Detail`

- **AND** 上述主 CTA 都必须携带 `fromDashboard=true`
- **AND** 都必须携带 `dashboardSection=current-focus`
- **AND** 都必须携带 `dashboardReturnTo=/dashboard`
- **AND** 不得把来自 `Current Focus / Next Action` 的高优先级反馈主 CTA 伪装成 `empty-state` 来源

#### Scenario: 系统已就绪时不强制给主 CTA

- **WHEN** 无缺口且有活动数据
- **THEN** Dashboard 必须进入“系统已就绪”中性状态
- **AND** 不再展示强制主 CTA
- **AND** 不得为了保持按钮存在而发明额外写入入口

### Requirement: Dashboard 来源上下文参数语义冻结

系统 SHALL 将 `fromDashboard`、`dashboardSection` 与 `dashboardReturnTo` 冻结为当前阶段唯一 Dashboard 来源参数集合。

#### Scenario: Dashboard 来源参数的最小集合

- **WHEN** 任一页面由 Dashboard 发起跳转
- **THEN** 路由搜索参数至少必须包含 `fromDashboard=true`、`dashboardSection`、`dashboardReturnTo=/dashboard`
- **AND** `dashboardSection` 只允许以下取值之一：`overview / current-focus / asset-feedback / recent-activity / empty-state`
- **AND** `dashboardReturnTo` 当前阶段只允许为 `/dashboard`
- **AND** 不得发明 `fromDashboardList`、`dashboardFrom`、`returnToDashboard` 等第二套命名

#### Scenario: fromDashboard 的角色解释

- **WHEN** 接手者判断 `fromDashboard` 的语义
- **THEN** 它必须被解释为 Dashboard 外层来源标记
- **AND** 它只服务“回到 Dashboard”的上下文恢复
- **AND** 不得替代目标页面原生的主来源上下文模型

### Requirement: Dashboard 多跳返回与上下文恢复规则冻结

系统 SHALL 将从 Dashboard 跳出后的主动返回、刷新恢复与多跳路径中的上下文承接冻结为单值结论。

> 关键区分：用户主动“返回 Dashboard”与目标页面上的写入成功后 reread 不是同一动作。`phase04-03 / 04-06` 已冻结的 canonical owner + reread 规则继续有效，不因 `fromDashboard` 而改写。

#### Scenario: 从 Dashboard 直接进入 canonical owner 页后返回

- **WHEN** 用户从 Dashboard 直接进入 `Module Detail`、`Product Detail`、`Repository Binding Detail / Workspace`、`Decision Detail`、`Module Registry / List`、`Product Registry / List`、`Repository Binding / List` 或 `Decision Center / List`
- **THEN** 页面必须保留 `fromDashboard=true`、`dashboardSection` 与 `dashboardReturnTo=/dashboard`
- **AND** 用户主动触发“返回 Dashboard”时，必须回到 `/dashboard`
- **AND** 返回后应恢复到对应来源区块的浏览上下文

#### Scenario: Dashboard 来源进入列表后再进入详情

- **WHEN** 用户从 Dashboard 进入 `Module Registry / List`、`Product Registry / List`、`Repository Binding / List` 或 `Decision Center / List`
- **AND** 再从列表进入某个详情页
- **THEN** 详情页可以继续承接列表原生来源模型（如 `fromList + queryText + statusFilter`）
- **AND** 同时允许保留 `fromDashboard=true + dashboardSection + dashboardReturnTo=/dashboard` 作为外层来源标记
- **AND** `fromList` 继续服务“返回列表”的立即返回路径
- **AND** `fromDashboard` 继续服务“返回 Dashboard”的外层返回路径
- **AND** 不得把两者混写成同一个主来源字段

#### Scenario: Dashboard 来源下的写入成功与 reread

- **WHEN** 用户从 Dashboard 进入 `Product Detail` 或 `Repository Binding Detail / Workspace` 后发起既有绑定写入
- **THEN** 写入成功后必须继续停留在对应 canonical owner 页面完成 reread
- **AND** 不得因为携带 `fromDashboard` 就自动跳回 Dashboard
- **AND** reread 完成后，如用户主动选择返回 Dashboard，才允许使用 `dashboardReturnTo=/dashboard`

#### Scenario: Dashboard 来源下 Create 页的返回与回流语义

- **WHEN** 用户从 Dashboard 空状态主 CTA 进入 `Module Registry / Create`、`Product Registry / Create` 或 `Repository Binding / Create`
- **THEN** Create 页必须携带 `fromDashboard=true`、`dashboardSection=empty-state`、`dashboardReturnTo=/dashboard`
- **AND** 用户主动取消 Create 时，必须返回到 `/dashboard`，而不是回列表或回其他默认页
- **AND** Create 提交成功后，必须继续保留 `fromDashboard=true + dashboardSection + dashboardReturnTo=/dashboard` 进入新建对象的 Detail 页，沿用 `phase02 / phase04` 已冻结的 Create 成功进入 Detail 的 reread 规则
- **AND** 从新建后的 Detail 页主动返回时，必须能回到 `/dashboard`，而不是只回列表
- **AND** 不得因为进入 Create 页就丢失 Dashboard 外层来源标记

#### Scenario: Dashboard 来源上下文刷新恢复

- **WHEN** 用户刷新携带 `fromDashboard=true` 的目标页面
- **THEN** 页面必须继续恢复 `fromDashboard`、`dashboardSection` 与 `dashboardReturnTo`
- **AND** 不得在刷新后静默丢失 Dashboard 来源标记

#### Scenario: 缺失或非法 Dashboard 返回参数

- **WHEN** 页面携带 `fromDashboard=true` 但 `dashboardSection` 非法或 `dashboardReturnTo` 缺失/非法
- **THEN** 系统必须返回明确失败语义或回退到冻结默认值
- **AND** `dashboardSection` 非法时必须回退为 `overview`
- **AND** `dashboardReturnTo` 缺失或非法时必须回退为 `/dashboard`
- **AND** 不得静默跳到错误页面或把用户带到根路径 `/`

### Requirement: Dashboard 不形成第二套写入主线

系统 SHALL 明确 Dashboard 在当前阶段只承接读与跳转，不承接任何直接写入完成动作。

#### Scenario: 判断 Dashboard 行为边界

- **WHEN** 后续实现讨论在 Dashboard 上添加“直接绑定”“直接补录”“直接提交决策”等动作
- **THEN** 必须判定为超出当前阶段范围
- **AND** Dashboard 只允许跳转到既有 canonical owner 页面完成后续操作
- **AND** 不得把 Dashboard 扩写为影子工作台

## MODIFIED Requirements

### Requirement: phase05-01 可点击入口热区的落点解释

`phase05-01` 已冻结 `Dashboard Home` 四类区块采用整卡可点统一规则，`phase05-03` 在此基础上 SHALL 将“可点”进一步解释为“必须落到既有 canonical owner 页面，并携带 Dashboard 来源参数”，而不是仅冻结视觉热区。

#### Scenario: 点击热区的落点解释

- **WHEN** 后续实现处理 `dashboard_overview`、`Current Focus / Next Action`、`Asset Feedback` 与 `Recent Activity` 的点击事件
- **THEN** 必须同时满足 `phase05-01` 的热区规则与本规格的 canonical 跳转矩阵
- **AND** 不得只实现可点击样式而缺失单值落点与来源参数

### Requirement: phase05-02 target_type / target_id 的跳转解释

`phase05-02` 已冻结 `Feedback Signal Card` 至少承接 `target_type / target_id / target_label`，`phase05-03` 在此基础上 SHALL 将其进一步解释为指向既有 canonical owner 页面所需的最小导航身份。

#### Scenario: target_type / target_id 的导航解释

- **WHEN** 后续合同或前端实现消费 `Feedback Signal Card` 的 `target_type / target_id`
- **THEN** 它们必须足以确定本规格冻结的唯一跳转目标
- **AND** 不得把 `target_type / target_id` 解释为“在 Dashboard 内直接完成动作”的内部指令

## REMOVED Requirements

### Requirement: recent_activity 使用笼统 binding 类型

**Reason**: `phase05` 当前阶段已经明确 `Binding` 活动至少拆分为 `product_module_binding / product_repository_binding / module_repository_binding`，否则无法唯一落到既有 canonical owner 页面，也无法保持合同与导航语义单值化。

**Migration**: 当前阶段统一移除笼统 `binding` 活动类型；若历史实现存在该类型，必须在进入实现前映射到三类明确子类型之一，再按本规格的落点规则消费。
