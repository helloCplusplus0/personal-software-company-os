# phase04_product_and_repository_binding_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase04_product_and_repository_binding_foundation` 的架构规划文档。

`phase04` 的目标不是零散补几个 `Product` 与 `Repository` 表单，而是在 `phase02` 已交付 `Module Registry`、`phase03` 已交付 `Decision Center` 的基础上，把 `Product Registry + Repository Binding` 提升为当前阶段的正式主线，并把 `phase02` 中临时承接的绑定动作回收到正确归属。

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
9. `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
10. `.trae/specs/phase03_11_decision_center_proto_mainline/`
11. `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`

## 3. 本阶段目标

`phase04` 的目标是：

> 在 `phase02` 与 `phase03` 已交付主线前提下，交付 `Product Registry + Repository Binding` 的最小可执行闭环，使 `Product`、`Repository` 与三类绑定关系从“存在于 MVP 正式规格与临时入口”推进为“有正式页面、有结构化合同、有稳定读写边界、可独立联调验收”的第三条执行主线。

本阶段需要回答的核心问题：

1. `Product Registry` 与 `Repository Binding` 的页面边界、动作归属与上下文入口是什么
2. `CreateProduct / CreateRepository / BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 的最小闭环如何单值化
3. `phase02` 中临时由 `Module Registry` 承接的绑定写入与候选读取，如何在 `phase04` 中迁回正确归属
4. 如何在不回退重做 `Module Registry` 与 `Decision Center` 的前提下，完成 `Product / Repository / Binding` 数据、合同与交互主线
5. 当前阶段结束时，仓库最少要新增哪些代码、合同、数据与验收证据，才算真正完成 `phase04` 交付

## 4. 架构冻结结论

### 4.1 当前阶段唯一执行层上游

`phase04` 必须直接承接：

- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
- `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
- `.trae/specs/phase03_11_decision_center_proto_mainline/`
- `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`

不允许在本阶段重新解释：

- `v0.1` 对象范围与动作矩阵
- `Module Registry` 已冻结页面、动作、合同与验收结论
- `Decision Center` 已冻结页面、动作、合同与验收结论
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA`

### 4.2 当前阶段主交付对象

`phase04` 的主交付对象是：

- `Product Registry`
- `Repository Binding`

其最小主线必须优先承接：

- `CreateProduct`
- `CreateRepository`
- `BindRepositoryToProduct`
- `BindModuleToProduct`
- `MapModuleToRepository`
- `Product` 与 `Repository` 的列表读取、详情读取与绑定关系读取
- 从 `Module Detail`、`Product Detail` 进入绑定动作的最小上下文入口

### 4.3 当前阶段前端交付策略

前端继续统一遵守：

- 单一 `React Web`
- 同时考虑 `PC` 与移动浏览器 UI
- 不拆分独立原生客户端
- `PWA` 仅保留兼容增强位

当前阶段重点是：

- 让 `Product List / Create / Detail` 成为正式主线页面
- 让 `Repository Binding` 不再只是 MVP 里的名字，而是真正承接仓库创建与绑定关系管理
- 让 `Module Detail` 从临时绑定工作台回落为摘要与上下文入口，而不是继续并行拥有主写入职责
- 不在 `phase04` 引入第二套前端架构

### 4.4 当前阶段数据与合同承接原则

`phase04` 直接承接的最小数据与接口闭环如下：

- `products`
- `repositories`
- `product_modules`
- `product_repositories`
- `module_repositories`
- 与 `modules` 的候选读取、详情摘要与已绑定结果读取

当前阶段关于合同的冻结如下：

- `Contract First` 必须从本阶段一开始进入 `Product / Repository / Binding` 最小 `.proto` 合同源
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层
- 过渡传输层不得形成与 `.proto` 并列的第二套合同源
- `phase02` 中临时承接的 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead` 必须在本阶段明确迁移策略，不得继续长期悬空

当前阶段不在架构规划中提前冻结：

- Go 数据访问层具体实现工具
- `Decision -> Product / Repository` 的正式关联写入主线
- Dashboard 聚合反馈与 `product_asset_coverage` 的完整消费链

补充冻结：

- `phase02` 中已经落地的 `BindModuleToProduct / MapModuleToRepository` 可运行能力，不得通过重建影子表、第二套绑定表或临时双写绕过迁移
- `phase04` 必须明确历史绑定数据、候选读取与既有前端入口在迁移后的兼容前提，保证既有路径升级后仍可继续读取

### 4.5 当前阶段交互归属原则

为了避免 `phase04` 后续 `/spec` 与实现阶段出现“同一个绑定动作由多个页面长期并行拥有”的歧义，当前阶段先冻结以下交互归属原则：

- `Product Registry / List` 承接产品列表读取、筛选入口、创建入口与进入详情入口
- `Product Create` 承接 `CreateProduct`
- `Product Detail` 承接产品详情读取、已绑定模块/仓库读取、`BindModuleToProduct` 与进入仓库绑定流程的上下文入口
- `Repository Binding` 承接仓库列表读取、创建入口、绑定工作台入口与三类关系管理
- `Repository Create` 承接 `CreateRepository`
- `Repository Binding Detail / Workspace` 承接仓库详情读取、候选读取、`BindRepositoryToProduct / MapModuleToRepository`
- `Module Detail` 当前阶段只承接绑定摘要与轻量跳转，不再作为主绑定工作台

当前阶段关于绑定动作归属的补充原则：

- 五个核心动作的页面归属必须单值化：
  - `CreateProduct` → `Product Create`
  - `CreateRepository` → `Repository Create`
  - `BindModuleToProduct` → `Product Detail`
  - `BindRepositoryToProduct` → `Repository Binding Detail / Workspace`
  - `MapModuleToRepository` → `Repository Binding Detail / Workspace`
- `BindModuleToProduct` 的主归属冻结为 `Product Detail`，不得再由 `Repository Binding` 或 `Module Detail` 并行拥有第二套主写入流程
- `BindRepositoryToProduct` 与 `MapModuleToRepository` 的主归属冻结为 `Repository Binding Detail / Workspace`
- `Product Detail` 与 `Module Detail` 允许保留跳转或带上下文入口，但不拥有第二套主写入流程
- 无可绑定候选时，页面必须返回明确空状态，而不是把空结果误报为接口错误
- `Module Detail` 中既有绑定入口在迁移后只允许保留为兼容跳转：可以带上 `moduleId / moduleName / fromModuleDetail` 等上下文进入正式主入口，但不得继续停留在本页直接提交绑定写入

### 4.6 当前阶段源码设计层输出要求

`phase04` 虽然当前处于 `/plan`，但为了保证后续 `/spec` 可直接进入实现，本阶段必须把以下源码设计层结果纳入任务规划：

- 前端 `Product / Repository / Binding` 页面与路由分层
- 产品列表、产品详情、仓库列表、仓库详情/绑定工作台的组件职责
- 产品与仓库最小结构化模板、状态语义与最小读模型
- 三类绑定关系的候选读取、重复绑定排除规则、空状态与上下文恢复规则
- `Product Registry`、`Repository Binding` 与 `Module Registry` 的页面级入口关系
- 后端 `Product Registry` 与 `Repository Binding` 的模块边界、接口分组与迁移边界
- `Product / Repository / Binding` 最小 `.proto` 合同落地与过渡传输层承接策略
- `phase02` 临时绑定承接点迁移后的兼容策略
- 联调环境的可重复建立入口（重置/基线/fixture），不得把这一步后移到验收补救
- `Product List` 与 `Repository List` 的路由搜索参数、刷新恢复与返回路径规则
- `PC / 移动浏览器` 双场景下的布局降级策略
- `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 在迁移完成后的 canonical owner、旧入口保留级别与 reread 承接页面

### 4.7 当前阶段规划吸取的 phase02 / phase03 经验

本阶段必须明确吸取前两阶段经验，避免重复补票：

- `.proto` 合同主线必须从阶段任务一开始纳入，不再中途补票
- 验收环境的重置脚本、基线数据与联调证据要求必须在实现任务中前置规划
- 绑定动作的归属必须在 `/plan` 阶段收口，不再把临时承接长期保留成既成事实
- 正式规格正文要在页面边界、数据边界、合同边界与验收基线冻结后再进入
- 根级同步必须作为本阶段内建任务，而不是在实现完成后再临时想起

### 4.8 当前阶段交付模式

`phase04` 必须按交付型 phase 推进，而不是只做文档冻结。

这意味着：

- 当前 `/plan` 只负责建立阶段上游与任务拆分
- 后续必须继续进入 `/spec`
- `/spec` 后必须继续进入实际源代码实现
- 实现完成后必须进入验证、验收与收口

## 5. 当前阶段范围冻结

### 5.1 本阶段必须进入范围

- `Product Registry / List`
- `Product Create`
- `Product Detail`
- `Repository Binding / List`
- `Repository Create`
- `Repository Binding Detail / Workspace`
- `CreateProduct`
- `CreateRepository`
- `BindRepositoryToProduct`
- `BindModuleToProduct`
- `MapModuleToRepository`
- `products / repositories / product_modules / product_repositories / module_repositories`
- `Product` 与 `Repository` 最小 `.proto` 合同主线
- 支撑 `phase04` 联调验收的可重复环境重置与基线恢复策略

### 5.2 本阶段允许最小承接但不扩写的连接

- `Module Detail` 中的绑定摘要、跳转与上下文入口
- `Decision Center` 作为上游决策背景的只读入口或未来连接位
- `Product / Repository` 与 `Decision` 的轻量展示前提
- `Venture` 保留为可选实体，但不进入当前阶段主线

### 5.3 本阶段明确不做

- Dashboard 最小反馈闭环实现
- `Decision -> Product / Repository` 正式关联写入主线
- `Feature / Opportunity / Experiment`
- GitHub OAuth / 自动导入
- 独立 `AI Assistant`
- 独立移动端客户端
- 完整 `PWA`
- 自动扫描代码或自动识别仓库关系

## 6. 本阶段输出物

当前 `/plan` 步骤必须产出：

1. `phase04_product_and_repository_binding_foundation_architecture_plan.md`
2. `phase04_product_and_repository_binding_foundation_dev_plan.md`
3. `phase04_product_and_repository_binding_foundation_shared_baseline.md`

当前 `/plan` 通过审核后，下一步再进入：

- `phase04` 对应 `/spec`

整个 `phase04` 最终还必须产出：

- `Product / Repository / Binding` 对应的正式规格正文
- `Product Registry` 前端最小可运行主线
- `Repository Binding` 前端最小可运行主线
- `Product / Repository / Binding` 后端、数据与最小 `.proto` 合同主线
- 联调、验收与根级同步结果

## 7. 本阶段不做

本阶段明确不做的能力范围：

- Dashboard 与反馈聚合实现
- `Decision Center` 的返工重做
- `Module Registry` 的重构式返工
- `Feature / Opportunity / Experiment`
- 独立 `AI Assistant`
- 独立 `React Native` 客户端
- 完整 `PWA`

补充说明：

- 当前 `/plan` 步骤本身不直接写业务代码
- 但整个 `phase04` 不是“禁止代码实现”，而是必须在后续子任务中完成代码交付

## 8. 通过标准

当以下条件满足时，`phase04` 的架构规划才算通过：

1. `Product Registry` 与 `Repository Binding` 的主交付边界已单值化
2. `phase04` 的动作范围与 `phase01-06` 正式规格正文保持一致，并直接承接 `phase02 + phase03` 已交付结果
3. 桌面与移动浏览器的单一前端交付策略已写清
4. `CreateProduct / CreateRepository / BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 的交互归属已写清
5. 临时绑定承接点的迁移策略已进入本阶段规划
6. 源码设计层输出要求与验收环境要求已进入本阶段规划
7. 后移对象与非目标边界已无冲突
8. 后续 `dev_plan` 与 `shared_baseline` 可以在此基础上继续展开 `/spec`、实现、验收与收口
