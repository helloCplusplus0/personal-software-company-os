# phase05_dashboard_feedback_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase05_dashboard_feedback_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase05` 继续遵守交付型 phase 原则：不是只把 Dashboard 画成一张静态页面，而是要完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路，并最终交付 `Dashboard + Feedback` 最小可执行闭环。

相较于 `phase03` 与 `phase04`，本阶段在任务拆分上显式吸取以下经验：

- Dashboard 必须先冻结信息层级与动作化边界，不能实现时临场拼装
- `Feedback` 必须先被定义为派生信号层，而不是在实现期膨胀成新的重实体系统
- `.proto` 合同、验收环境与空状态基线必须前置进入任务主线
- 聚合页与写入页的 owner 必须保持单值化，Dashboard 只负责读与跳转

## 2. 本阶段目标

在 `phase04` 已交付 `Product Registry + Repository Binding` 最小主线的前提下，交付 `Dashboard + Feedback` 最小可执行闭环，使用户能够在 Dashboard 中看到当前资产状态、关键缺口与下一步动作，并基于结构化跳转回到既有主页面完成处理。

## 3. 子任务清单

### 第一组：冻结类子任务

### phase05-01 冻结 Dashboard 页面边界、信息层级与区块结构

范围：

- 冻结 `Dashboard Home` 的最小页面边界
- 冻结 `Current Focus / Next Action / Recent Activity / Asset Feedback` 的最小区块结构
- 冻结 `dashboard_overview / Current Focus（内部承接 Next Action 主队列） / Asset Feedback / Recent Activity` 四个区块的固定归属关系
- 冻结桌面与移动浏览器下的信息优先级与布局降级策略
- 冻结 Dashboard 与既有主导航的关系
- 冻结 `Dashboard Home` 的正式业务入口路由为 `/dashboard`，以及其与根级布局宿主的关系

DoD：

- `Dashboard Home` 页面职责单值化
- 第一屏优先级、区块层级与可点击入口单值化
- `Current Focus / Next Action` 是唯一主行动队列、`Asset Feedback` 是补充摘要区块、`Recent Activity` 是独立活动流的关系已单值化
- `Dashboard Home` 的正式业务入口、主导航归属与是否承接 `/` 已单值化
- 不形成复杂驾驶舱或第二套首页体系
- 页面边界与 `PSCO-summarize-feedback.md` 的 MVP 页面范围一致

### phase05-02 冻结 Feedback 信号模板、优先级与最小展示模型

范围：

- 冻结 `pending_decision_signals` 的 Dashboard 消费语义
- 冻结 `product_asset_coverage` 的 Dashboard 消费语义
- 冻结 `recent_activity_feed` 与 `dashboard_overview` 的最小展示模型
- 冻结 `product missing both bindings` 在 `product_asset_coverage` 中的独立读模型语义
- 冻结反馈信号的优先级、排序与空状态语义
- 冻结 `recent_activity_feed` 的显式排序字段与时间语义
- 冻结统一 `Feedback Signal Card` 的最小字段模板
- 冻结 `Current Focus` 与 `Asset Feedback` 的最大展示数量

DoD：

- `Feedback` 当前阶段被单值化为派生信号层
- 最小反馈卡片模板明确
- 优先级与排序前提明确，并固定为 `pending_decision_signals > missing both bindings > missing repository binding > missing module binding`
- `product_asset_coverage` 已明确包含 `missing both bindings` 的独立计数或代表项语义
- `recent_activity_feed` 的排序时间字段已显式化，不依赖隐式实现假设
- `Current Focus / Next Action` 最多展示 `5` 条主卡片、`Asset Feedback` 最多展示 `3` 条代表性缺口项
- `recent_activity_feed` 明确为独立活动流，不进入反馈优先级队列
- 不引入独立 `Feedback` 重实体、不引入复杂评分模型

### phase05-03 冻结 Dashboard 跳转目标、来源上下文与返回路径

范围：

- 冻结从 Dashboard 到 `Product / Repository / Module / Decision` 的最小跳转矩阵
- 冻结信号卡片、最近活动列表与空状态 CTA 的目标页面
- 冻结 `recent_activity_feed` 中 `Module / Release / Product / Repository / Decision / Binding` 各类型的单值落点
- 冻结 `fromDashboard` 来源标记、来源区块标记与最小返回路径规则

DoD：

- 任一 Dashboard 卡片都有明确 canonical owner 落点
- `Release` 在当前阶段不新增独立 Detail 页面，其活动项统一回落到所属 `Module Detail`
- `Binding` 活动类型至少拆分为 `product_module_binding / product_repository_binding / module_repository_binding`，不得以笼统 `binding` 保留歧义
- Dashboard 不再承接第二套写入流程
- 返回 Dashboard 时的来源恢复规则明确
- 不形成“从 Dashboard 直接完成绑定/补录”的影子工作台

### phase05-04 冻结 Dashboard 聚合读范围、接口边界与错误语义前提

范围：

- 冻结 Dashboard 所需的最小聚合读范围
- 冻结概览读取、信号读取与最近活动读取的最小接口边界
- 冻结空状态、局部失败、整页失败与无信号语义
- 冻结空系统、缺 Product、缺 Repository、待决策、资产缺口并存时的主 CTA 优先级矩阵
- 冻结“空系统”与“已有结构化资产但仍无 Module”的区分规则

DoD：

- 当前阶段聚合读边界明确
- Dashboard 不承接新的业务写入接口
- 空信号、空活动、空系统与错误语义明确
- 非空但无 Module 的系统状态不再与冷启动空系统混同
- 空状态 CTA 优先级明确，且不允许同时并排出现多个同级主 CTA
- 不提前冻结外部埋点、趋势分析或通知中心接口

### 第二组：实现设计产出类子任务

### phase05-05 产出 Dashboard 前端页面、路由与组件分层设计

范围：

- 产出 `Dashboard Home` 页面分层
- 产出区块级组件、卡片级组件与列表级组件分层
- 产出 Dashboard 与既有导航、详情页之间的最小路由与参数承接方式
- 产出 `PC / 移动浏览器` 双场景下的布局降级策略

DoD：

- 前端页面与路由分层明确
- 区块级与卡片级组件职责明确
- 无第二套移动端 UI 架构
- 设计结果足以直接进入实现

### phase05-06 产出 Dashboard 前端状态模型与交互流设计

范围：

- 产出 Dashboard 整页查询状态、分区加载状态与局部错误状态模型
- 产出空状态、有数据状态、无反馈信号状态与最近活动为空时的页面行为
- 产出主 CTA 优先级命中规则与多缺口并存时的区块降级策略
- 产出从 Dashboard 跳出后返回的上下文恢复策略
- 产出刷新恢复、路由搜索参数与来源透传策略

DoD：

- 页面级状态模型明确
- 分区级状态模型明确
- Dashboard 跳转与返回链路明确
- 设计结果足以直接进入实现

### phase05-07 产出 Dashboard/Feedback 后端模块边界与接口分组设计

范围：

- 产出 Dashboard 聚合读取模块边界
- 产出概览读取、信号读取与最近活动读取的接口分组
- 产出与 `Module Registry`、`Decision Center`、`Product Registry`、`Repository Binding` 的服务侧连接边界
- 产出聚合读取 owner、query service owner 与 DTO/合同映射策略
- 产出 `Feedback Signal Card` 的归一化组装 owner 与 `recent_activity` 类型映射 owner

DoD：

- 后端模块边界明确
- Dashboard 聚合查询 owner 明确
- 读接口分组明确
- 不提前冻结 Go 数据访问层具体工具
- 设计结果足以直接进入实现

### phase05-08 产出 Dashboard + Feedback 最小 Protocol Buffers 合同设计

范围：

- 基于前置冻结结果产出 `Dashboard + Feedback` 最小 `.proto` 合同设计
- 明确 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 的消息结构、服务接口、包名与版本语义
- 明确当前阶段 `.proto` 与 `chi + JSON HTTP` 过渡传输层的显式映射策略
- 明确 `FeedbackSignal` 与 `RecentActivityItem` 的最小字段模板、枚举值与最大返回数量
- 明确 `product_asset_coverage` 中 `missing both bindings` 的合同表达方式，以及 `RecentActivityItem` 的显式时间字段

DoD：

- `.proto` 合同不晚于正式规格正文进入阶段主线
- 合同字段语义、字段编号与页面区块单值一致
- 聚合读取、反馈信号读取与最近活动读取的消息边界明确
- `FeedbackSignal` 的最小字段、优先级枚举与 `RecentActivityItem` 的类型枚举已单值化
- `RecentActivityItem` 的时间字段与排序前提已单值化
- 合同演进规则明确，包括 `reserved` 与 breaking check 前提

### phase05-09 产出联调验收环境、Dashboard 冷启动基线与 fixture 设计

范围：

- 产出 `Dashboard + Feedback` 联调所需的数据库重置、基线种子与 fixture 设计
- 产出 Dashboard 空状态、最小有数据状态与局部错误状态的验收基线
- 产出 Dashboard 跳转到各 canonical owner 页面后的返回链路验收矩阵
- 产出最近活动与反馈信号的最小 fixture 组合
- 产出“空系统 / 仅有模块 / 有产品无模块 / 有产品缺仓库 / 有产品缺模块 / 有待决策 / 有最近活动”七类最小 fixture 组合

DoD：

- 验收环境建立方式明确
- 空系统与有数据系统都可重复建立
- Dashboard 主路径、空状态路径与跳转返回路径验收矩阵明确
- 七类最小 fixture 组合明确，且每类都能映射到单值 CTA 或单值区块结果
- 不再依赖临时手工 SQL 才能完成验收

### 第三组：规格、实现与验收子任务

### phase05-10 产出首份 Dashboard + Feedback 正式规格文档

范围：

- 基于前九个子任务产出 `phase05` 对应的 `/spec`
- 作为后续实现与下一阶段的直接上游规格来源

DoD：

- 文档完整覆盖页面、区块、聚合读、反馈信号、API、合同、验收基线、非目标与 Done 标准
- 与 `phase01-06` 正式 MVP 规格正文、`phase04` 正式规格/验收结论互链一致

### phase05-11 落实 Dashboard + Feedback 最小 Protocol Buffers 合同主线

范围：

- 将 `phase05-08` 已冻结的 `.proto` 合同正式落地为仓库内唯一合同源
- 落地 `buf build / lint / generate / breaking` 的最小工具链入口
- 明确当前阶段 DTO/HTTP 适配层与 `.proto` 的语义映射落点

DoD：

- `Dashboard + Feedback` 最小 `.proto` 合同已落地为单一合同源
- `buf` 校验链可运行，breaking check 基准路径正确
- 过渡传输层不得形成与 `.proto` 并列的第二套合同源

### phase05-12 实现后端与数据主线

范围：

- 实现 Dashboard 所需的最小后端聚合读接口
- 实现反馈信号所需的最小派生查询
- 实现最近活动、资产概览与反馈信号的数据主线
- 实现联调所需的重置脚本、基线 seed 与最小 fixture

DoD：

- 后端聚合读接口可运行
- 数据主线与已冻结边界一致
- 聚合查询与既有主线可重复执行
- 联调环境可重复建立，不依赖手工补数据

### phase05-13 实现前端 Dashboard 主线

范围：

- 实现 `Dashboard Home` 前端主线
- 实现 `Current Focus / Next Action / Recent Activity / Asset Feedback` 最小区块
- 实现 Dashboard 到既有 canonical owner 页面之间的跳转与返回上下文
- 落实单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略

DoD：

- Dashboard 前端主线可运行
- 最小反馈信号在前端可走通
- Dashboard 只负责读与跳转，不形成第二套写入工作台
- 无第二套移动端 UI 架构

### phase05-14 联调、验证与验收

范围：

- 完成前后端联调
- 完成 Dashboard 空状态、有数据状态、局部错误状态与最近活动展示验收
- 完成反馈信号与概览聚合的最小验收路径
- 完成从 Dashboard 到 `Product / Repository / Module / Decision` 的跳转与返回路径验证

DoD：

- `Dashboard + Feedback` 最小主线可完整走通
- 最小反馈信号、聚合读取与跳转返回路径可被重复复核
- 验收结果可重复复核，并可明确证明当前阶段已形成可运行交付物
- 发现的问题已收口到当前阶段，不遗留隐性阻断

### phase05-15 审核与根级同步

范围：

- 完成 `phase05` 文档互链复核
- 回写根级状态
- 确认下一阶段入口

DoD：

- 根级文档与 `phase05` 文档保持单值一致
- `plan.md` 中 `phase05` 状态更新正确
- 下一阶段的进入条件清楚

## 4. 明确不做

本阶段不做：

- 第二套资产注册或绑定主线
- 复杂驾驶舱 / BI 分析页
- 外部遥测、自动指标回流与通知中心
- 独立 `Feedback` 重实体 CRUD
- `Feature / Opportunity / Experiment`
- GitHub OAuth / 自动导入
- 自动扫描代码

## 5. 依赖关系

执行顺序固定为：

1. `phase05-01`
2. `phase05-02`
3. `phase05-03`
4. `phase05-04`
5. `phase05-05`
6. `phase05-06`
7. `phase05-07`
8. `phase05-08`
9. `phase05-09`
10. `phase05-10`
11. `phase05-11`
12. `phase05-12`
13. `phase05-13`
14. `phase05-14`
15. `phase05-15`

不允许跳过前置冻结、合同设计与验收基线设计，直接进入 `Dashboard + Feedback` 实现；也不允许只完成文档冻结而不完成本阶段的代码交付、验收与收口。
