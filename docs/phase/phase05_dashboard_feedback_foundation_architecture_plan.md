# phase05_dashboard_feedback_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase05_dashboard_feedback_foundation` 的架构规划文档。

当前根级真相源已完成 `phase05` 收口。本文档保留为 `phase05` 的架构规划与冻结记录；文中“当前阶段”均指 `phase05` 当时上下文，不覆盖根级当前“`phase05` 已完成收口、下一阶段正式 phase 入口待建立后切换”的状态。

`phase05` 的已交付边界统一以 `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`、`.trae/specs/phase05_11_dashboard_feedback_proto_mainline/` 与 `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md` 为准。

`phase05` 的目标不是补一页“看上去像驾驶舱”的统计大盘，而是在 `phase04` 已交付 `Product Registry + Repository Binding` 主线的基础上，冻结 `Dashboard + Feedback` 最小闭环的主交付边界、聚合归属、合同前提与实现范围，再继续进入 `/spec`、实现、验收与收口。

当前阶段必须解决两个历史风险：

- `Dashboard` 不能退化为无行动价值的统计页
- `Feedback` 不能直接膨胀成一套新的重实体或完整运营系统

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `PSCO-summarize-feedback.md`
7. `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
8. `.trae/specs/phase04_11_product_repository_binding_proto_mainline/`
9. `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md`

补充说明：

- `phase02` 与 `phase03` 的已交付边界继续通过 `phase04` 正式规格与验收结论继承
- 当前阶段不得把 `phase02`、`phase03` 重新拉回为与 `phase04` 并列的第二套直接执行层上游

## 3. 本阶段目标

`phase05` 的目标是：

> 在 `phase04` 已交付 `Product / Repository / Binding` 主线前提下，交付 `Dashboard + Feedback` 的最小可执行闭环，使用户能够在一个统一页面中看到“当前资产状态 + 关键缺口 + 下一步动作”，并通过结构化跳转回到既有 canonical owner 页面完成后续处理。

本阶段需要回答的核心问题：

1. `Dashboard` 在 `v0.1` 中到底承接哪些页面区块、聚合视图与最小交互
2. `Feedback` 在当前阶段是“什么”，以及它如何避免扩写成新的重实体系统
3. `phase03` 预留的 `pending_decision_signals` 与 `phase04` 预留的 `product_asset_coverage` 如何进入统一的最小反馈主线
4. `Dashboard` 的行动卡片如何跳回既有 `Product / Repository / Module / Decision` canonical owner 页面，而不是再造第二套写入入口
5. 当前阶段结束时，仓库最少要新增哪些代码、合同、查询与验收证据，才算真正完成 `phase05`

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase05` 必须直接承接：

- `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
- `.trae/specs/phase04_11_product_repository_binding_proto_mainline/`
- `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md`

不允许在本阶段重新解释：

- `Product Registry` 已冻结页面、动作、数据、合同与验收结论
- `Repository Binding` 已冻结页面、动作、数据、合同与验收结论
- 三类绑定动作的 canonical owner 与 reread 规则
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA`

### 4.2 当前阶段主交付对象

`phase05` 的主交付对象是：

- `Dashboard`
- `Feedback Signal Layer`

其最小主线必须优先承接：

- `Dashboard Home` 的统一聚合读取
- `Current Focus / Next Action` 风格的最小反馈卡片
- `pending_decision_signals` 的正式消费链
- `product_asset_coverage` 的正式消费链
- 最近活动与资产概览的最小只读聚合
- 从 Dashboard 跳回既有 `Product / Repository / Module / Decision` 主页面的上下文入口

当前阶段对 `Feedback` 的单值解释冻结如下：

- `Feedback` 当前阶段是基于既有结构化数据派生出的“反馈信号层”
- 其职责是把“缺什么、下一步该去哪里补”显式展示出来
- 当前阶段不引入独立 `Feedback` 重实体、完整 PMM 指标系统、外部遥测接入或市场反馈采集系统

### 4.3 当前阶段前端交付策略

前端继续统一遵守：

- 单一 `React Web`
- 同时考虑 `PC` 与移动浏览器 UI
- 不拆分独立原生客户端
- `PWA` 仅保留兼容增强位

当前阶段重点是：

- 让 Dashboard 成为“行动导向”的首页，而不是平铺统计页
- 让第一屏优先暴露当前缺口与下一步动作，而不是大量概念化指标
- 让空状态、最近活动、反馈信号与跳转入口在桌面和窄屏下都可用
- 不在 `phase05` 引入第二套前端架构

### 4.4 当前阶段数据与合同承接原则

`phase05` 直接承接的最小数据与接口闭环如下：

- `modules`
- `releases`
- `products`
- `repositories`
- `decisions`
- `decision_links`
- `product_modules`
- `product_repositories`
- `module_repositories`

当前阶段必须正式消费的派生语义如下：

- `pending_decision_signals`
- `product_asset_coverage`
- `recent_activity_feed`
- `dashboard_overview`

其中补充冻结以下最小读模型前提：

- `product_asset_coverage` 至少显式区分：`product missing both bindings`、`product missing repository binding`、`product missing module binding`
- `recent_activity_feed` 必须带出显式活动时间字段，并以该字段作为活动流排序基准；当前阶段不得依赖隐式 `created_at` 推断最近活动顺序

当前阶段关于合同的冻结如下：

- `Contract First` 必须从本阶段一开始进入 `Dashboard + Feedback` 最小 `.proto` 合同源
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层
- 过渡传输层不得形成与 `.proto` 并列的第二套合同源
- 派生反馈信号必须来自既有结构化对象与绑定关系，不得在实现阶段临时发明第二套隐式数据源

当前阶段不在架构规划中提前冻结：

- 外部行为埋点
- 自动化消息推送或通知中心
- 完整 Capability Growth 评分系统
- 长周期 BI 分析与趋势预测
- 外部 GitHub / 使用数据接入

当前阶段关于四类读取在 Dashboard 中的单值归属补充冻结如下：

- `dashboard_overview` 只承接概览卡片区块，不进入反馈优先级队列
- `pending_decision_signals` 与 `product_asset_coverage` 必须先归一化为同一套 `Feedback Signal Card`，再进入 `Current Focus / Next Action` 主队列
- `Asset Feedback` 是 `product_asset_coverage` 的补充摘要区块，不再形成第二条独立优先级队列
- `recent_activity_feed` 是独立活动流区块，不与反馈信号共用排序逻辑

当前阶段关于 `Feedback Signal Card` 的最小模板补充冻结如下：

- 每条卡片至少包含：`signal_family`、`signal_code`、`priority`、`title`、`summary`、`action_label`、`target_type`、`target_id`、`target_label`
- `Current Focus / Next Action` 主队列最多展示 `5` 条卡片
- `Asset Feedback` 补充区块最多展示 `3` 条代表性缺口项
- 同优先级内排序默认按“最近需要处理时间优先”；若上游读模型暂不提供显式处理时间，则回退为 `created_at DESC`

当前阶段关于反馈信号优先级的单值顺序补充冻结如下：

1. `pending_decision_signals`
2. `product missing both bindings`
3. `product missing repository binding`
4. `product missing module binding`

不允许在实现阶段把 `recent_activity_feed` 与反馈信号混合排序，也不允许把 `Asset Feedback` 再做成第二套 `Next Action` 队列

### 4.5 当前阶段交互归属原则

为了避免 `phase05` 在后续 `/spec` 与实现阶段出现“Dashboard 既展示又偷偷承接写入”的歧义，当前阶段先冻结以下交互归属原则：

- `Dashboard Home` 承接聚合读取、反馈信号展示、最近活动展示与空状态引导
- `Dashboard Home` 中的反馈卡片只负责展示与跳转，不负责直接提交写入
- `Product Detail` 继续承接 `BindModuleToProduct`
- `Repository Binding Detail / Workspace` 继续承接 `BindRepositoryToProduct / MapModuleToRepository`
- `Decision Detail` 或 `Decision Center` 继续承接决策查看与后续决策处理动作
- `Module Detail` 继续承接模块摘要与既有上下文入口

当前阶段关于 Dashboard 跳转的补充原则：

- 任一 `Next Action` 卡片必须能落到一个既有 canonical owner 页面
- Dashboard 不得生成“补全绑定”“补录决策”的第二套快捷写入工作台
- 从 Dashboard 跳出后，返回路径、来源标记与最小上下文恢复规则必须在后续 `/spec` 中单值化

当前阶段关于 Dashboard 正式入口的补充冻结如下：

- `Dashboard Home` 的正式业务入口路由冻结为 `/dashboard`
- `Dashboard` 在当前阶段作为既有主导航中的一级入口新增，不替代根级布局宿主本身
- 当前阶段不把 `/` 单值解释为 `Dashboard Home`；如未来需要把 `/` 重定向到 `/dashboard`，必须在后续 `phase / fix` 中单独冻结

当前阶段关于 `Feedback Signal` 与 `Recent Activity` 的最小跳转矩阵补充冻结如下：

- `pending_decision_signals`：
  - 单项信号已绑定具体 `decision_id` → 跳转 `Decision Detail`
  - 聚合信号未绑定单一 `decision_id` → 跳转 `Decision Center / List`
- `product missing module binding` → 跳转 `Product Detail`
- `product missing repository binding` → 跳转 `Product Detail`
- `product missing both bindings` → 跳转 `Product Detail`
- `recent_activity: module` → 跳转 `Module Detail`
- `recent_activity: release` → 跳转所属 `Module Detail`
- `recent_activity: product` → 跳转 `Product Detail`
- `recent_activity: repository` → 跳转 `Repository Binding Detail / Workspace`
- `recent_activity: decision` → 跳转 `Decision Detail`
- `recent_activity: product_module_binding` → 跳转 `Product Detail`
- `recent_activity: product_repository_binding` → 跳转 `Repository Binding Detail / Workspace`
- `recent_activity: module_repository_binding` → 跳转 `Repository Binding Detail / Workspace`

补充约束：

- 当前阶段不新增 `Release Detail` 页面；`release` 活动统一折叠到所属 `Module Detail`
- 当前阶段不允许把 `Binding` 活动笼统写成一个未区分类型的对象；必须至少区分 `product_module_binding / product_repository_binding / module_repository_binding`

### 4.6 当前阶段源码设计层输出要求

`phase05` 虽然当前处于 `/plan`，但为了保证后续 `/spec` 可直接进入实现，本阶段必须把以下源码设计层结果纳入任务规划：

- `Dashboard Home` 页面、区块与路由分层
- `Current Focus / Next Action / Recent Activity / Asset Feedback` 的最小信息层级
- 反馈信号的优先级、排序与最小展示模型
- `Dashboard` 到 `Product / Repository / Module / Decision` 的跳转与返回上下文规则
- Dashboard 空状态、局部错误、局部 loading 与整页 loading 的交互策略
- Dashboard 聚合读取与反馈信号读取的后端边界、接口分组与 owner
- `Dashboard + Feedback` 最小 `.proto` 合同落地与过渡传输层承接策略
- 验收环境的可重复建立入口、Dashboard 冷启动空状态与有数据状态 fixture
- `PC / 移动浏览器` 双场景下的信息层级降级策略

### 4.7 当前阶段规划吸取的 phase03 / phase04 经验

本阶段必须明确吸取前两阶段经验，避免重复补票：

- `Dashboard` 必须动作化，不能停留在“有系统感但不驱动工作”的概念驾驶舱
- 必须优先复用 `pending_decision_signals`、`product_asset_coverage` 等既有上游语义，而不是实现时临时发明新口径
- 聚合页必须遵守单一真相源，不得把写入职责重新塞回 Dashboard
- `.proto` 合同主线必须从阶段任务一开始纳入，不再后补
- 验收环境、空状态与最近活动基线必须前置进入阶段任务，不再等联调时补票

### 4.8 当前阶段交付模式

`phase05` 必须按交付型 phase 推进，而不是只做文档冻结。

这意味着：

- 当前 `/plan` 只负责建立阶段上游与任务拆分
- 后续必须继续进入 `/spec`
- `/spec` 后必须继续进入实际源代码实现
- 实现完成后必须进入验证、验收与收口

## 5. 当前阶段范围冻结

### 5.1 本阶段必须进入范围

- `Dashboard Home`
- `dashboard_overview` 最小聚合读取
- `pending_decision_signals` 正式消费链
- `product_asset_coverage` 正式消费链
- `recent_activity_feed` 最小聚合读取
- `Current Focus / Next Action / Recent Activity / Asset Feedback` 最小页面区块
- 从 Dashboard 跳回 `Product / Repository / Module / Decision` 的结构化入口
- 支撑 `phase05` 联调验收的可重复 fixture、空状态与有数据基线
- `Dashboard + Feedback` 最小 `.proto` 合同主线

### 5.2 本阶段允许最小承接但不扩写的连接

- `Release` 只作为最近活动与反馈信号的数据来源，不在当前阶段扩写新页面
- `Capability` 只允许作为派生视角或标签语义，不新建表、不新建主页面
- `Venture` 保留为可选语义，不进入当前阶段主线
- `Decision -> Product / Repository` 继续按上游已冻结边界承接，不在当前阶段扩写新的写入主线

### 5.3 本阶段明确不做

- 复杂驾驶舱或 BI 大盘
- 外部遥测、埋点采集与自动指标回流
- 独立 `Feedback` 重实体 CRUD
- 实验复盘系统、机会池、Feature 主线
- 自动推荐、自动打分、AI 自动建议
- 消息中心、通知中心、Inbox
- GitHub OAuth / 自动导入
- 独立 `AI Assistant`
- 独立移动端客户端
- 完整 `PWA`

## 6. 本阶段输出物

当前 `/plan` 步骤必须产出：

1. `phase05_dashboard_feedback_foundation_architecture_plan.md`
2. `phase05_dashboard_feedback_foundation_dev_plan.md`
3. `phase05_dashboard_feedback_foundation_shared_baseline.md`

当前 `/plan` 通过审核后，下一步再进入：

- `phase05` 对应 `/spec`

整个 `phase05` 最终还必须产出：

- `Dashboard + Feedback` 对应的正式规格正文
- `Dashboard + Feedback` 最小 `.proto` 合同主线
- `Dashboard + Feedback` 后端聚合读取与数据主线
- `Dashboard Home` 前端最小可运行主线
- 联调、验收与根级同步结果

## 7. 本阶段不做

本阶段明确不做的能力范围：

- 第二套资产注册或绑定主线
- 完整 Capability Growth 评分体系
- 趋势分析与长期 BI
- 自动化提醒中心
- 外部数据接入
- `Feature / Opportunity / Experiment`
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA`

补充说明：

- 当前 `/plan` 步骤本身不直接写业务代码
- 但整个 `phase05` 不是“禁止代码实现”，而是必须在后续子任务中完成代码交付
