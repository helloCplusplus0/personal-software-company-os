# phase03_decision_center_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase03_decision_center_foundation` 的架构规划文档。

目标不是把 `phase03` 做成只补一页 `Decision` 表单的轻量补丁，而是在 `phase02` 已交付 `Module Registry` 主线的基础上，冻结 `Decision Center` 最小闭环的主交付边界、交互归属、合同前提与实现范围，再继续进入 `/spec`、实现、验收与收口。

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `PSCO-summarize-feedback.md`
7. `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
8. `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
9. `.trae/specs/phase02_11a_module_registry_proto_contract/`
10. `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`

## 3. 本阶段目标

`phase03` 的目标是：

> 在 `phase02` 已交付 `Module Registry` 主线前提下，交付 `Decision Center` 的最小可执行闭环，使 `Decision` 从“存在于共识与附属入口”推进为“可结构化记录、可关联 Module、可直接验收”的第二条执行主线。

本阶段需要回答的核心问题：

1. `Decision Center` 在 `v0.1` 中到底承接哪些对象、动作与页面职责
2. `RecordDecision` 与 `LinkDecisionToTarget` 的最小交互闭环是什么
3. `Decision` 如何与 `Module Registry` 保持直接但不反向重写 `phase02` 的连接
4. 如何把 `Protocol Buffers` 作为当前阶段的合同起点，而不是再走一遍 `phase02-11A` 式补票
5. 如何在不引入第二套前端架构的前提下，同时满足 `PC` 与移动浏览器可用性
6. 哪些实现设计层结果必须在 `/plan` 中冻结，才能让后续 `/spec` 与实现直接落地
7. 当前 `phase03` 结束时，仓库里最少要新增哪些代码、合同、数据与验收结果，才算真正完成本阶段交付

## 4. 架构冻结结论

### 4.1 当前阶段唯一执行层上游

`phase03` 必须直接承接：

- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
- `.trae/specs/phase02_11a_module_registry_proto_contract/`
- `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`

不允许在本阶段重新解释：

- `v0.1` 对象范围
- `Module Registry` 已冻结页面、动作、数据、合同与验收结论
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA` 能力

### 4.2 当前阶段主交付对象

`phase03` 主交付对象是：

`Decision Center`

其最小主线应优先承接：

- `Decision` 结构化记录
- `Decision` 列表读取与详情读取
- `RecordDecision`
- `LinkDecisionToTarget` 的最小闭环，其中 `Decision -> Module` 为当前阶段必交付连接
- 从 `Module Detail` 进入 `Decision Center` 或进入带上下文的 `RecordDecision` 入口

### 4.3 当前阶段前端交付策略

前端仍统一遵守：

- 单一 `React Web`
- 同时考虑 `PC` 与移动浏览器 UI
- 不拆分独立原生客户端
- `PWA` 仅保留兼容增强位

当前阶段重点是：

- 让 `Decision List / Create / Detail` 在桌面与窄屏下都可用
- 让 `Module Detail` 中已有的 `Decision` 入口从只读/跳转提升为真实执行入口
- 不在 `phase03` 引入第二套 UI 架构

### 4.4 当前阶段数据与合同承接原则

`phase03` 直接承接的最小数据与接口闭环如下：

- `decisions`
- `decision_links`
- 与 `modules` 的最小目标关联读取
- `phase02` 既有 `decisions` 只读前提表向 `phase03` 结构化主线的原位演进

当前阶段关于合同的冻结如下：

- `Contract First` 从本阶段一开始就必须落到 `Decision Center` 最小 `.proto` 合同源
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层
- 但过渡传输层不得形成与 `.proto` 并列的第二套合同源，字段语义必须与 `.proto` 保持单值一致

当前阶段不在架构规划中提前冻结：

- Go 数据访问层具体实现工具
- `Decision` 与 `Product / Repository` 全量目标体系
- Dashboard 聚合接口与 `pending_decision_signals` 的完整消费链

补充冻结：

- `phase02` 中仅用于 `Module Detail` 只读展示的 `decisions` 表，不得在 `phase03` 中通过新建并列表或临时影子表绕过升级
- `phase03` 必须明确现有示例 `Decision` 与既有 `decision_links` 的迁移兼容前提，保证历史示例在升级后仍可继续读取与展示

### 4.5 当前阶段交互归属原则

为了避免 `phase03` 在后续 `/spec` 和实现阶段出现“决策存在但不知道由谁承接”的歧义，当前阶段先冻结以下交互归属原则：

- `Decision Center / List` 承接决策列表读取、筛选入口、创建入口与进入详情入口
- `Decision Create` 承接 `RecordDecision`
- `Decision Detail` 承接决策详情读取、目标关联展示与 `Decision -> Module` 候选读取
- `Decision Detail` 承接 `LinkDecisionToTarget` 的最小写入触点
- `Module Detail` 当前阶段只承接轻量入口与上下文跳转，不在 `phase03` 中反向扩写为第二个 `Decision` 工作台

当前阶段关于目标候选读取的补充原则：

- `Decision -> Module` 的候选读取属于 `Decision Detail` 的附属读取，不单独扩写为新的主工作台
- 候选读取的职责是支撑 `LinkDecisionToTarget`，而不是把 `Module Registry` 反向并入 `Decision Center`
- 当前阶段必须明确候选范围、排序、已关联目标排除规则与无候选时的页面行为，避免实现期各自发明

### 4.6 当前阶段源码设计层输出要求

`phase03` 虽然仍处于 `/plan`，但为了保证后续 `/spec` 可直接进入实现，本阶段必须把以下源码设计层结果纳入任务规划：

- 前端页面与路由分层
- 页面级状态模型与交互状态流转
- `Decision` 列表读取、详情读取、创建写入与目标关联写入的最小读写模型
- `Decision` 最小结构化模板的字段级约束、`status` 枚举与创建校验前提
- `Decision List` 最小读模型中 `link_count` 与 `linked_module_summary` 的计算口径与空值语义
- 后端 `Decision Center` 模块边界与接口分组前提
- `Decision Center` 最小 `Protocol Buffers` 合同落地与过渡传输层承接策略
- `Decision Detail` 中目标候选读取的接口归属、返回模型与前端面板承接方式
- `phase02` 既有 `Decision` 示例数据与 `decision_links` 的迁移兼容、回填与保留策略
- 验收环境的可重复建立入口（重置/基线/fixture），不得把这一步后移到验收补救
- `Decision List` 的路由搜索参数、刷新恢复与返回列表上下文恢复规则
- `PC / 移动浏览器` 双场景下的布局降级策略

### 4.7 当前阶段规划吸取的 phase02 经验

本阶段必须明确吸取 `phase02` 的规划与执行经验，避免重复补票：

- `.proto` 合同主线必须从阶段任务一开始纳入，不再通过中途追加 `11A` 类补充任务收口
- 验收环境的重置脚本、基线数据与联调证据要求必须在实现任务中前置规划，不能等到验收时再手工补 SQL
- `phase03` 只承接 `Decision Center` 最小闭环，不反向重写 `Module Registry` 已冻结边界
- 当前阶段必须明确“什么是本阶段直接交付，什么只是允许最小连接”，避免像 `phase02` 一样在候选读取与主写入之间反复扩 scope

### 4.8 当前阶段交付模式

`phase03` 必须按交付型 phase 推进，而不是只做文档冻结。

这意味着：

- 当前 `/plan` 只负责建立阶段上游与任务拆分
- 后续必须继续进入 `/spec`
- `/spec` 后必须继续进入实际源代码实现
- 实现完成后必须进入验证、验收与收口

## 5. 当前阶段范围冻结

### 5.1 本阶段必须进入范围

- `Decision Center / List`
- `Decision Create`
- `Decision Detail`
- `RecordDecision` 最小结构化模板
- `Decision -> Module` 最小目标关联闭环
- `Decision Detail` 中面向 `Module` 的最小候选读取与选择交互
- 从 `Module Detail` 进入 `Decision Center` 的最小上下文入口
- 前端页面/路由/组件分层的最小实现设计
- 后端 `Decision Center` 模块边界、接口分组与最小 `.proto` 合同源
- `phase02` 既有 `Decision` 示例数据与 `decision_links` 的迁移兼容方案
- 支撑 `phase03-14` 验收的可重复环境重置与基线恢复策略

### 5.2 本阶段允许最小承接但不扩写的连接

- `Decision -> Product / Repository` 的合同保留位或轻量候选读取前提
- `pending_decision_signals` 作为后续 Dashboard 或提醒能力的上游数据语义
- 从 `Module Registry` 进入 `Decision Center` 的跳转或预填上下文

### 5.3 本阶段明确不做

- Product 全量主线
- Repository 全量主线
- Dashboard 完整反馈逻辑
- `Feature / Opportunity / Experiment`
- 独立 `AI Assistant`
- 独立移动端客户端
- 完整 `PWA`
- AI 自动推荐或自动生成决策

## 6. 本阶段输出物

当前 `/plan` 步骤必须产出：

1. `phase03_decision_center_foundation_architecture_plan.md`
2. `phase03_decision_center_foundation_dev_plan.md`
3. `phase03_decision_center_foundation_shared_baseline.md`

当前 `/plan` 通过审核后，下一步再进入：

- `phase03` 对应 `/spec`

整个 `phase03` 最终还必须产出：

- `Decision Center` 对应的正式规格正文
- `Decision Center` 前端最小可运行主线
- `Decision Center` 后端、数据与最小 `.proto` 合同主线
- 联调、验收与根级同步结果

## 7. 本阶段不做

本阶段明确不做的能力范围：

- Product / Repository / Dashboard 的独立实现阶段规划
- `Feature / Opportunity / Experiment`
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA`

补充说明：

- 当前 `/plan` 步骤本身不直接写业务代码
- 但整个 `phase03` 不是“禁止代码实现”，而是必须在后续子任务中完成代码交付

## 8. 通过标准

当以下条件满足时，`phase03` 的架构规划才算通过：

1. `Decision Center` 的主交付边界已单值化
2. `phase03` 的动作范围与 `phase01-06` 正式规格正文保持一致，并直接承接 `phase02` 已交付结果
3. 桌面与移动浏览器的单一前端交付策略已写清
4. `RecordDecision`、`LinkDecisionToTarget` 与 `Module Detail` 入口的交互归属已写清
5. 源码设计层输出要求与验收环境要求已进入本阶段规划
6. 后移对象与非目标边界已无冲突
7. 后续 `dev_plan` 与 `shared_baseline` 可以在此基础上继续展开 `/spec`、实现、验收与收口
