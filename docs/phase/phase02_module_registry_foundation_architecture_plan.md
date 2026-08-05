# phase02_module_registry_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase02_module_registry_foundation` 的架构规划文档。

目标不是把 `phase02` 做成纯文档冻结阶段，而是先在 `/plan` 中冻结 `Module Registry` 最小可执行主线的架构边界、交付范围、阶段输出物与实现前提，再继续进入 `/spec`、实现、验收与收口。

## 2. 上游输入

本阶段唯一上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `PSCO-summarize-feedback.md`
7. `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`

## 3. 本阶段目标

`phase02` 的目标是：

> 在 `phase01-06` 正式 MVP 规格正文前提下，交付 `Module Registry` 的最小可执行主线，使其成为 `v0.1` 第一条可实现、可验收、可继续扩展的执行入口。

本阶段需要回答的核心问题：

1. `Module Registry` 在 `v0.1` 中到底承接哪些对象、动作与页面职责
2. 模块登记、模块列表、模块详情与版本登记的最小交互闭环是什么
3. `Module` 如何与 `Release / Product / Repository / Decision` 保持最小但清晰的连接
4. 如何在单一 `React Web` 交付策略下，同时满足 `PC` 与移动浏览器可用性
5. 哪些能力必须进入 `phase02`，哪些必须后移到 `phase03+`
6. 在不提前进入实现的前提下，当前阶段最少需要冻结哪些源码实现设计层结果，才能保证 `/spec` 与后续实现可直接落地
7. 当前 `phase02` 结束时，仓库里最少要新增哪些可运行代码与验收结果，才算真正完成本阶段交付

## 4. 架构冻结结论

### 4.1 当前阶段唯一执行层上游

`phase02` 必须直接承接：

`.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`

不允许在本阶段重新解释：

- `v0.1` 对象范围
- `Module` 以外的后移对象
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA` 能力

### 4.2 当前阶段主交付对象

`phase02` 主交付对象是：

`Module Registry`

其最小主线应优先承接：

- `CreateModule`
- `CreateRelease`
- `BindModuleToProduct`（以最小触点承接）
- `MapModuleToRepository`（以最小触点承接）
- 与 `Decision` 的只读或关联入口

### 4.3 当前阶段前端交付策略

前端仍统一遵守：

- 单一 `React Web`
- 同时考虑 `PC` 与移动浏览器 UI
- 不拆分独立原生客户端
- `PWA` 仅保留兼容增强位

当前阶段重点是：

- 让列表、表单、详情和基础关联入口在桌面与窄屏下都可用
- 不在 `phase02` 引入第二套 UI 架构

### 4.4 当前阶段数据与合同承接原则

`phase02` 只承接 `Module Registry` 所需的最小数据与接口闭环：

- `modules`
- `module_releases`
- `product_modules`
- `module_repositories`
- 与 `decisions` 的关联读取或跳转入口

当前阶段关于合同的补充冻结如下：

- `Contract First` 在 `phase02` 中不再只停留在长期方向，必须为 `Module Registry` 落地最小 `.proto` 合同源
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层
- 但过渡传输层不得形成与 `.proto` 并列的第二套合同源，字段语义必须与 `.proto` 保持单值一致

当前阶段不在架构规划中提前冻结：

- Go 数据访问层具体实现工具
- 完整查询接口矩阵
- Dashboard 聚合接口

### 4.5 当前阶段交互归属原则

为了避免 `phase02` 在后续 `/spec` 和实现阶段出现“动作存在但不知道由谁承接”的歧义，当前阶段先冻结以下交互归属原则：

- `Module List` 承接模块读取、筛选入口、创建入口与进入详情入口
- `Module Create` 承接 `CreateModule`
- `Module Detail` 承接 `CreateRelease`
- `Module Detail` 承接 `BindModuleToProduct` 与 `MapModuleToRepository` 的最小写入触点，而不是仅提供跳转
- `Decision` 在 `phase02` 中只承接只读展示或跳转入口，不在本阶段扩写为 `Decision Center` 主线

### 4.6 当前阶段源码设计层输出要求

`phase02` 虽然仍处于 `/plan`，但为了保证后续 `/spec` 可直接进入实现，本阶段必须把以下源码设计层结果纳入任务规划：

- 前端页面与路由分层
- 页面级状态模型与交互状态流转
- 列表读取、详情读取、创建写入、版本写入的最小读写模型
- 后端 `Module Registry` 模块边界与接口分组前提
- `Module Registry` 最小 `Protocol Buffers` 合同落地与过渡传输层承接策略
- `PC / 移动浏览器` 双场景下的布局降级策略

### 4.7 当前阶段交付模式

`phase02` 必须按交付型 phase 推进，而不是只做文档冻结。

这意味着：

- 当前 `/plan` 只负责建立阶段上游与任务拆分
- 后续必须继续进入 `/spec`
- `/spec` 后必须继续进入实际源代码实现
- 实现完成后必须进入验证、验收与收口

## 5. 当前阶段范围冻结

### 5.1 本阶段必须进入范围

- 模块列表页范围
- 模块创建入口
- 模块详情页最小结构
- 版本登记最小流
- 模块与 `Product / Repository` 的最小关联入口
- 面向空状态用户的首轮模块登记路径
- 前端页面/路由/组件分层的最小实现设计
- 后端 `Module Registry` 模块边界与接口分组的最小实现设计
- `Module Registry` 最小 `.proto` 合同源

### 5.2 本阶段允许最小承接但不扩写的连接

- 到 `Decision Center` 的入口或关联展示
- 到 `Product Registry` / `Repository Binding` 的轻量关联

### 5.3 本阶段明确不做

- `Feature / Opportunity / Experiment`
- 独立 `AI Assistant`
- Dashboard 完整反馈逻辑
- Product 全量主线
- Repository 全量主线
- 独立移动端客户端
- 完整 `PWA`

## 6. 本阶段输出物

当前 `/plan` 步骤必须产出：

1. `phase02_module_registry_foundation_architecture_plan.md`
2. `phase02_module_registry_foundation_dev_plan.md`
3. `phase02_module_registry_foundation_shared_baseline.md`

当前 `/plan` 通过审核后，下一步再进入：

- `phase02` 对应 `/spec`

整个 `phase02` 最终还必须产出：

- `Module Registry` 对应的正式规格正文
- `Module Registry` 前端最小可运行主线
- `Module Registry` 后端与数据主线
- 联调、验收与根级同步结果

## 7. 本阶段不做

本阶段明确不做的能力范围：

- Product / Decision / Repository 的独立实现阶段规划
- `Feature / Opportunity / Experiment`
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA`

补充说明：

- 当前 `/plan` 步骤本身不直接写业务代码
- 但整个 `phase02` 不是“禁止代码实现”，而是必须在后续子任务中完成代码交付

## 8. 通过标准

当以下条件满足时，`phase02` 的架构规划才算通过：

1. `Module Registry` 的主交付边界已单值化
2. `phase02` 的动作范围与 `phase01-06` 正式规格正文保持一致
3. 桌面与移动浏览器的单一前端交付策略已写清
4. 绑定动作、版本动作与 `Decision` 入口的交互归属已写清
5. 源码设计层输出要求已进入本阶段规划
6. 后移对象与非目标边界已无冲突
7. 后续 `dev_plan` 与 `shared_baseline` 可以在此基础上继续展开 `/spec`、实现、验收与收口
