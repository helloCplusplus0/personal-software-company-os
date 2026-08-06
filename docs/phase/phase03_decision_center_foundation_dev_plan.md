# phase03_decision_center_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase03_decision_center_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase03` 继续遵守交付型 phase 原则：不是只把 `Decision` 写成一份规格，而是要完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路，并最终交付 `Decision Center` 最小可执行闭环。

相较于 `phase02`，本阶段在任务拆分上显式吸取以下经验：

- `Protocol Buffers` 合同不再后补，必须前置进入任务主线
- 验收环境、重置脚本与基线 fixture 不再等到联调阶段补票
- 主交付对象与“允许最小连接”的对象分开写清，避免实现期反复扩 scope
- 正式规格正文要在前端、后端、合同、验收设计都冻结后再进入，减少后续返工

## 2. 本阶段目标

在 `phase02` 已交付 `Module Registry` 主线前提下，交付 `Decision Center` 最小可执行闭环，使 `Decision` 从“附属入口与共识要求”进入“可结构化记录、可关联目标、可独立验收”的正式执行主线。

## 3. 子任务清单

### 第一组：冻结类子任务

### phase03-01 冻结 Decision Center 页面边界与信息结构

范围：

- 冻结 `Decision List / Decision Create / Decision Detail` 的最小页面结构
- 冻结 `Decision Center` 与 `Module Detail` 之间的入口关系
- 冻结桌面与移动浏览器下的基础信息密度策略

DoD：

- `Decision Center` 页面职责单值化
- 不形成第二套移动端 UI 方案
- 页面边界与 `phase01-06` 正式规格正文一致

### phase03-02 冻结 Decision 模板、状态语义与最小展示模型

范围：

- 冻结 `Decision` 最小结构化模板
- 冻结 `status` 的最小状态语义
- 冻结列表展示与详情展示所需的最小字段集合
- 冻结字段级 required / optional、`alternatives` 的最小结构与创建校验前提

DoD：

- `Decision` 记录模板明确
- `Decision` 最小字段集合与字段级 required / optional 已明确
- `alternatives` 的最小结构已明确为可重复文本条目集合，不允许前后端各自发明不同嵌套结构
- `status` 的当前阶段枚举已明确并与后续 `.proto` 合同单值一致
- 创建校验前提已明确，包括必填字段、空字符串处理与最小合法输入要求
- 列表与详情的最小读模型明确
- 不引入超出 `v0.1` 的复杂审批、投票或自动化字段

### phase03-03 冻结 LinkDecisionToTarget 的目标范围与入口上下文

范围：

- 冻结 `LinkDecisionToTarget` 的当前阶段目标范围
- 冻结 `Decision -> Module` 的直接闭环
- 冻结从 `Module Detail` 发起带上下文记录决策的最小入口方式

DoD：

- `Decision -> Module` 关联路径明确
- `Module Detail` 与 `Decision Center` 的交互归属明确
- `Product / Repository` 在当前阶段只保留受控连接位，不扩写为第二主线

### phase03-04 冻结数据读写范围、接口边界与错误语义前提

范围：

- 冻结 `Decision Center` 所需的最小数据读写范围
- 冻结列表读取、详情读取、创建写入、关联写入的最小接口边界
- 冻结创建失败、目标不存在、重复关联等错误语义前提
- 冻结 `RecordDecision` 与 `LinkDecisionToTarget` 的请求校验与失败语义归属

DoD：

- 当前阶段数据与接口边界明确
- `RecordDecision` 与 `LinkDecisionToTarget` 的校验失败类型已进入阶段规划
- 关键异常路径已进入阶段规划，而不是联调时临时补写
- 不提前冻结超出当前阶段的聚合分析接口

### 第二组：实现设计产出类子任务

### phase03-05 产出前端页面、路由与组件分层设计

范围：

- 产出 `Decision List / Decision Create / Decision Detail` 的页面分层
- 产出页面之间与 `Module Detail` 入口之间的最小路由结构
- 产出列表、模板表单、目标关联面板的组件分层
- 产出 `PC / 移动浏览器` 双场景下的布局降级策略

DoD：

- 前端页面与路由分层明确
- 页面级组件职责明确
- 无第二套移动端 UI 架构
- 设计结果足以直接进入实现

### phase03-06 产出前端状态模型与交互流设计

范围：

- 产出列表查询状态、筛选状态与空状态的最小模型
- 产出记录决策表单、目标关联面板与详情页的交互状态流转
- 产出成功、失败、空状态与返回路径
- 产出列表查询条件在路由搜索参数与页面状态之间的承接策略
- 产出从 `Decision Create / Decision Detail` 返回 `Decision List` 时的上下文恢复规则

DoD：

- 页面级状态模型明确
- 列表、创建、详情、关联之间的交互流明确
- 列表筛选维度已明确，至少冻结 `queryText` 与 `statusFilter`
- 查询条件的默认来源、路由持久化策略与刷新恢复策略已明确
- 从 `Decision Create` 或 `Decision Detail` 返回 `Decision List` 时，必须恢复原有 `queryText` 与 `statusFilter`
- 设计结果足以直接进入实现
- 不把运行时实现细节提前写成既成事实

### phase03-07 产出后端模块边界与接口分组设计

范围：

- 产出 `Decision Center` 在后端的模块边界
- 产出列表读取、详情读取、创建写入、关联写入的接口分组
- 产出 `Decision Center` 与 `Module Registry` 的服务侧连接边界

DoD：

- 后端模块边界明确
- 读写接口分组明确
- 不提前冻结 Go 数据访问层具体工具
- 设计结果足以直接进入实现

### phase03-08 产出 Decision Center 最小 Protocol Buffers 合同设计

范围：

- 基于前置冻结结果产出 `Decision Center` 最小 `.proto` 合同设计
- 明确 `DecisionListRead / DecisionDetailRead / DecisionWrite / DecisionLinkWrite` 的消息结构、服务接口、包名与版本语义
- 明确当前阶段 `.proto` 与 `chi + JSON HTTP` 过渡传输层的显式映射策略

DoD：

- `.proto` 合同不晚于正式规格正文进入阶段主线
- 合同字段语义、字段编号与页面动作单值一致
- 合同演进规则明确，包括删除字段后的 `reserved` 约束与 breaking check 前提

### phase03-09 产出联调验收环境与重置基线设计

范围：

- 产出 `Decision Center` 联调所需的数据库重置、基线种子与 fixture 设计
- 产出联调前置脚本入口、验收基线数据范围与异常路径验证前提
- 产出从空状态到首条 `Decision` 的冷启动验收路径

DoD：

- 验收环境建立方式明确
- 重置脚本、基线数据与异常路径验证要求已进入阶段任务
- 不再依赖临时手工 SQL 才能完成验收

### 第三组：规格、实现与验收子任务

### phase03-10 产出首份 Decision Center 正式规格文档

范围：

- 基于前九个子任务产出 `phase03` 对应的 `/spec`
- 作为后续实现与 `phase04` 的直接上游规格来源

DoD：

- 文档完整覆盖页面、动作、数据读写、API、合同、验收基线、非目标、实现设计层结果与 Done 标准
- 与 `phase01-06` 正式 MVP 规格正文、`phase02` 正式规格与验收结论互链一致

### phase03-11 落实 Decision Center 最小 Protocol Buffers 合同主线

范围：

- 将 `phase03-08` 已冻结的 `.proto` 合同正式落地为仓库内唯一合同源
- 落地 `buf build / lint / generate / breaking` 的最小工具链入口
- 明确当前阶段 DTO/HTTP 适配层与 `.proto` 的语义映射落点

DoD：

- `Decision Center` 最小 `.proto` 合同已落地为单一合同源
- `buf` 校验链可运行，breaking check 基准路径正确
- 过渡传输层不得形成与 `.proto` 并列的第二套合同源

### phase03-12 实现后端与数据主线

范围：

- 实现 `Decision Center` 所需的最小后端读写接口
- 实现 `decisions / decision_links` 与 `Decision -> Module` 关联的数据主线
- 实现 `Decision Center` 联调所需的重置脚本、基线 seed 与最小 fixture

DoD：

- 后端读写接口可运行
- 数据主线与已冻结边界一致
- 联调环境可重复建立，不依赖手工补数据

### phase03-13 实现前端 Decision Center 主线

范围：

- 实现 `Decision List / Decision Create / Decision Detail` 前端主线
- 实现列表、创建、详情、目标关联与从 `Module Detail` 发起的最小上下文入口
- 落实单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略

DoD：

- 前端主线可运行
- 列表、创建、详情与最小目标关联在前端可走通
- 无第二套移动端 UI 架构

### phase03-14 联调、验证与验收

范围：

- 完成前后端联调
- 完成 `RecordDecision` 与 `LinkDecisionToTarget` 的最小验收路径
- 完成空状态、错误态、目标不存在、重复关联与返回路径验证

DoD：

- `Decision Center` 最小主线可完整走通
- 验收结果可重复复核，并可明确证明当前阶段已形成可运行交付物
- 发现的问题已收口到当前阶段，不遗留隐性阻断

### phase03-15 审核与根级同步

范围：

- 完成 `phase03` 文档互链复核
- 回写根级状态
- 确认下一阶段入口

DoD：

- 根级文档与 `phase03` 文档保持单值一致
- `plan.md` 中 `phase03` 状态更新正确
- `phase04` 的进入条件清楚

## 4. 明确不做

本阶段不做：

- Product 全量主线
- Repository 全量主线
- Dashboard 聚合反馈实现
- `Feature / Opportunity / Experiment` 的重新引入
- 自动扫描代码
- 自动知识图谱
- AI 自动生成或自动判断决策
- 独立 `React Native` 客户端
- 完整 `PWA`

## 5. 依赖关系

执行顺序固定为：

1. `phase03-01`
2. `phase03-02`
3. `phase03-03`
4. `phase03-04`
5. `phase03-05`
6. `phase03-06`
7. `phase03-07`
8. `phase03-08`
9. `phase03-09`
10. `phase03-10`
11. `phase03-11`
12. `phase03-12`
13. `phase03-13`
14. `phase03-14`
15. `phase03-15`

不允许跳过前置收敛、合同设计与验收基线设计，直接进入 `Decision Center` 实现；也不允许只完成文档冻结而不完成本阶段的代码交付、验收与收口。
