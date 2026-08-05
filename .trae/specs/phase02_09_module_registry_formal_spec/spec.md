# Phase02-09 产出首份 Module Registry 正式规格文档 Spec

## Why

`phase02-01` 到 `phase02-08` 已经分别冻结了 `Module Registry` 的页面边界、实体主线、创建闭环、版本与绑定动作、数据与 API 边界、前端实现设计以及后端实现设计。但这些结论目前仍分散在多个子规格中，后续 `phase02-10 / 11 / 12` 与 `phase03` 若继续直接引用这些零散结论，仍会回到多入口、多解释和实现前临场拼装状态。

因此，当前必须产出一份单一、完整、可互链的 `Module Registry` 正式规格正文，把 `phase01-06` 的正式 MVP 上游与 `phase02-01 ~ 08` 的冻结结论统一收口为后续实现与验收的直接规格来源。

## What Changes

- 冻结 `phase02` 对应 `Module Registry` 正式规格正文的覆盖范围与目标文件落点
- 冻结该正式规格对 `phase01-06` 正式 MVP 规格正文与 `phase02-01 ~ 08` 子规格结论的继承关系
- 冻结页面、动作、数据读写、API、空状态、非目标、实现设计层结果与 Done 标准在正式规格中的统一承接要求
- 冻结 `Module Registry` 正式规格正文作为后续实现与 `phase03` 直接上游规格来源的定位
- 明确正式规格正文必须与根级真相源和当前 `phase02` 阶段文档互链一致，不得形成第二套边界

## Impact

- Affected specs: `phase02_module_registry_foundation`、`phase02_01_module_registry_pages_info_arch`、`phase02_02_module_entity_release_line`、`phase02_03_module_create_empty_state`、`phase02_04_release_binding_read_model`、`phase02_05_data_api_boundary`、`phase02_06_frontend_page_route_component_design`、`phase02_07_frontend_state_interaction_flow`、`phase02_08_backend_module_boundary_interface_grouping`
- Affected code: 当前无代码改动；影响后续正式规格正文编写、`phase02-10 / 11 / 12 / 13` 实现与验收口径，以及 `phase03` 对 `Module Registry` 的引用入口

## ADDED Requirements

### Requirement: Module Registry 正式规格正文必须收口到单一文档

系统 SHALL 将 `phase02` 对应的 `Module Registry` 正式规格收口到单一正文文档，而不是继续以 `phase02-01 ~ 08` 零散子规格作为长期并列入口。

#### Scenario: 正式规格正文文件落点

- **WHEN** 后续实现或验收需要引用 `Module Registry` 的正式规格
- **THEN** 必须存在单一正文文件 `phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
- **AND** 该正文文件必须能够从当前 spec 目录追踪
- **AND** 不得要求接手者先手工拼装多个 `phase02-01 ~ 08` 子规格后才能理解当前阶段正式口径

### Requirement: 正式规格正文覆盖范围必须完整

系统 SHALL 要求 `Module Registry` 正式规格正文在同一处完整覆盖页面、动作、数据读写、API、空状态、非目标、实现设计层结果与 Done 标准。

#### Scenario: 判断正式规格正文是否覆盖完整

- **WHEN** 接手者编写或审阅 `Module Registry` 正式规格正文
- **THEN** 文档必须至少完整覆盖页面矩阵、动作矩阵、数据矩阵、最小 API 分组、空状态与冷启动路径、非目标矩阵、实现设计层结果与 Done 标准
- **AND** 不得只保留零散冻结结论或只写“概要说明”而缺少统一正文

### Requirement: phase01-06 与 phase02-01~08 结论必须统一继承

系统 SHALL 要求 `Module Registry` 正式规格正文完整继承 `phase01-06` 正式 MVP 规格正文与 `phase02-01 ~ 08` 已冻结的单值结论，而不是重新定义第二套边界。

#### Scenario: 判断正式规格正文是否继承前置冻结

- **WHEN** 正式规格涉及技术路线、对象范围、页面职责、动作归属、数据边界、API 分组、空状态或实现设计结果
- **THEN** 必须以 `mvp_spec_v0.1.md` 作为当前阶段唯一执行层上游
- **AND** 必须完整承接 `phase02-01 ~ 08` 已冻结的当前阶段结论
- **AND** 不得绕开既有子规格扩范围、改口径或引入第二套解释

### Requirement: 页面与动作章节必须单值化

系统 SHALL 要求正式规格正文在同一处完整承接 `Module Registry / List`、`Module Create`、`Module Detail`、`Release Create` 四类页面主线及其动作归属。

#### Scenario: 判断页面与动作章节是否合格

- **WHEN** 正式规格定义页面结构与动作矩阵
- **THEN** 必须明确 `Module Registry / List` 承接列表读取、筛选入口、创建入口与进入详情入口
- **AND** 必须明确 `Module Create` 承接 `CreateModule`
- **AND** 必须明确 `Module Detail` 承接详情读取、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository` 与 `Decision` 只读入口
- **AND** 不得把 `Product Registry`、`Repository Binding`、`Decision Center` 的独立主线提前并入当前正式规格正文

### Requirement: 空状态与冷启动章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接当前阶段空状态、冷启动路径与返回路径的单值结论。

#### Scenario: 判断空状态与冷启动章节是否合格

- **WHEN** 正式规格定义空状态、首轮录入与回流路径
- **THEN** 必须明确用户可从 `Module Registry / List` 空状态直接进入 `CreateModule`
- **AND** 必须明确首轮允许先登记模块、再补充 `Product / Repository` 关联
- **AND** 必须明确 `CreateModule` 成功默认回流到 `ModuleDetail`
- **AND** 必须明确 `CreateRelease` 与绑定动作完成后仍回流或停留在当前模块详情上下文
- **AND** 不得把导入、自动扫描或 AI 建议写成当前阶段空状态主入口

### Requirement: 数据与 API 章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接当前阶段的最小数据矩阵、读写范围、候选读取前提与接口分组。

#### Scenario: 判断数据与 API 章节是否合格

- **WHEN** 正式规格定义数据模型方向或 API 边界
- **THEN** 必须明确直接承接 `modules`、`module_releases`、`product_modules`、`module_repositories`
- **AND** 必须明确最小读取或关联前提包含 `decisions`、`decision_links`
- **AND** 必须明确候选读取前提包含只读的 `products`、`repositories`
- **AND** 必须明确最小读组为 `ModuleListRead`、`ModuleDetailRead`
- **AND** 必须明确最小写组为 `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite`
- **AND** 必须明确 `Decision` 在当前阶段作为 `ModuleDetailRead` 的附属读取承接，不设独立读接口组

### Requirement: 前端实现设计结果必须进入正式规格正文

系统 SHALL 要求正式规格正文显式承接 `phase02-06 / 07` 已冻结的前端实现设计层结果，使后续前端实现不需要再次回到子规格拆读。

#### Scenario: 判断前端实现设计结果是否被承接

- **WHEN** 正式规格定义前端实现设计层结果
- **THEN** 必须明确页面集合、最小路由结构、URL 语义、页面壳层与组件职责
- **AND** 必须明确列表、创建、详情、版本登记与绑定动作的状态模型、默认回流路径与错误呈现位置
- **AND** 必须明确单一 `React Web` 下同时覆盖 `PC` 与移动浏览器的布局降级策略
- **AND** 不得要求后续前端实现再额外回到多个子规格中重新拼装主线

### Requirement: 后端实现设计结果必须进入正式规格正文

系统 SHALL 要求正式规格正文显式承接 `phase02-08` 已冻结的后端实现设计层结果，使后续后端实现可直接按既定边界进入编码。

#### Scenario: 判断后端实现设计结果是否被承接

- **WHEN** 正式规格定义后端实现设计层结果
- **THEN** 必须明确 `Module Registry` 为单一后端业务模块
- **AND** 必须明确 `handler / service / repository / candidate` 的分层语义与文件落点
- **AND** 必须明确 `Product / Repository` 候选读取由 `Module Registry` 临时承接并通过独立接口边界隔离
- **AND** 必须明确 `Decision` 读取内嵌于 `ModuleDetailRead`，不设独立读接口组或独立文件落点
- **AND** 不得把 Go HTTP/RPC 框架、ORM、SQL Builder 或数据访问层具体工具写成当前正式规格阻断项

### Requirement: 非目标与 Done 标准章节必须存在

系统 SHALL 要求正式规格正文显式包含当前阶段非目标矩阵与正式 Done 标准，而不是只保留功能边界。

#### Scenario: 判断非目标与 Done 标准章节是否存在

- **WHEN** 正式规格进入验收阶段
- **THEN** 文档必须明确 Product 全量主线、Decision Center 全量主线、Repository Binding 全量主线、Dashboard 聚合反馈、自动扫描代码、自动知识图谱、独立 `React Native` 客户端、完整 `PWA` 为当前阶段非目标
- **AND** 必须明确何时可判定 `Module Registry` 当前阶段规格足以进入 `phase02-10 / 11 / 12`
- **AND** 不得缺失验收口径或把 Done 标准分散到多个文档中

### Requirement: 正式规格正文必须成为后续直接上游

系统 SHALL 将 `Module Registry` 正式规格正文定位为后续实现、验收与 `phase03` 引用 `Module Registry` 主线时的直接上游规格来源。

#### Scenario: 判断后续阶段是否引用直接上游

- **WHEN** `phase02-10` 到 `phase02-13` 或 `phase03` 需要引用 `Module Registry` 当前阶段边界
- **THEN** 必须优先引用 `module_registry_spec_v0.1.md`
- **AND** 不得继续并列引用多个 `phase02-01 ~ 08` 子规格作为长期正式入口
- **AND** 不得绕开正式规格正文直接以对话结论替代规格入口

### Requirement: 正式规格正文必须与根级真相源和 phase 文档互链一致

系统 SHALL 要求正式规格正文与根级真相源及当前 `phase02` 阶段文档保持单值一致，不得形成第二套正式真相源。

#### Scenario: 判断正式规格正文与真相源关系

- **WHEN** 正式规格引用项目定位、技术基线、目录规则、最终共识或当前阶段范围
- **THEN** 必须与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致
- **AND** 必须与 `phase02_module_registry_foundation_architecture_plan.md`、`phase02_module_registry_foundation_dev_plan.md`、`phase02_module_registry_foundation_shared_baseline.md` 保持单值一致
- **AND** 不得在正式规格正文中重写一套与上述文档冲突的主结论

## MODIFIED Requirements

### Requirement: phase02 后续实现与 phase03 的规格入口

后续 `phase02-10` 到 `phase02-13` 的实现、验收与收口，以及 `phase03` 对 `Module Registry` 的引用，SHALL 以本次产出的 `module_registry_spec_v0.1.md` 作为直接上游规格入口；而该正文自身继续以 `phase01-06` 的 `mvp_spec_v0.1.md` 作为当前阶段唯一执行层上游。

#### Scenario: 后续阶段引用正式规格

- **WHEN** 后续阶段需要引用 `Module Registry` 当前阶段正式规格
- **THEN** 必须从 `module_registry_spec_v0.1.md` 进入
- **AND** 该正文仍必须显式继承 `mvp_spec_v0.1.md` 的已冻结边界
- **AND** 不得把 `phase01-06` 与 `phase02-01 ~ 08` 并列改写成多个长期直接入口

## REMOVED Requirements

### Requirement: 以后续实现继续直接拼装 phase02-01~08 子规格
**Reason**: `phase02-09` 的目标就是把前八个子任务的冻结结论收口为单一正式规格正文，避免后续实现与验收继续依赖多文档临场拼装。
**Migration**: 后续实现、验收与 `phase03` 若需要引用 `Module Registry` 当前阶段边界，应统一从 `module_registry_spec_v0.1.md` 进入；前置子规格仅作为该正文的冻结来源与追踪依据。
