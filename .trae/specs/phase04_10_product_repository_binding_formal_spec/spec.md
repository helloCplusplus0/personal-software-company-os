# Phase04-10 产出首份 Product / Repository / Binding 正式规格文档 Spec

## Why

`phase04-01` 到 `phase04-09` 已分别冻结了 `Product Registry` 与 `Repository Binding` 的页面边界、模板与状态语义、三类绑定关系与上下文入口、数据与错误边界、前后端实现设计、最小 `.proto` 合同，以及联调验收环境、重置基线与兼容迁移设计。但这些结论目前仍分散在多个子规格中，后续实现、验收与 `phase05` 如果继续直接引用这些零散结论，仍会回到多入口、多解释和实现前临场拼装状态。

因此，当前必须产出一份单一、完整、可互链的 `Product / Repository / Binding` 正式规格正文，把 `phase01-06` 的正式 MVP 上游、`phase02` 与 `phase03` 的已交付结果，以及 `phase04-01 ~ 09` 的冻结结论统一收口为后续实现与 `phase05` 的直接上游规格来源。

## What Changes

- 冻结 `Product / Repository / Binding` 正式规格正文的目标文件落点与文档定位
- 冻结该正式规格对 `phase01-06` 正式 MVP 规格正文、`phase02` 正式规格与验收结果、`phase03` 正式规格与验收结论、`phase04-01 ~ 09` 子规格的继承关系
- 冻结页面、动作、模板、绑定关系、数据读写、API、合同、验收基线、迁移边界、非目标、实现设计层结果与 Done 标准在正式规格中的统一承接要求
- 冻结 `Product / Repository / Binding` 正式规格正文作为后续实现与 `phase05` 直接上游规格来源的定位
- 明确正式规格正文必须与根级真相源、`phase04` 三件套、`phase02 / phase03` 已交付结果互链一致，不得形成第二套边界

## Impact

- Affected specs:
  - `phase04_product_and_repository_binding_foundation`
  - `phase04_01_product_registry_repository_binding_pages_info_arch`
  - `phase04_02_product_repository_template_status_read_model`
  - `phase04_03_binding_relation_candidate_scope_context_entry`
  - `phase04_04_product_repository_binding_data_api_error_boundary`
  - `phase04_05_frontend_page_route_component_design`
  - `phase04_06_frontend_state_interaction_flow`
  - `phase04_07_backend_module_boundary_interface_grouping`
  - `phase04_08_product_repository_binding_proto_contract`
  - `phase04_09_integration_acceptance_reset_baseline_compat_migration`
  - `phase02_09_module_registry_formal_spec`
  - `phase02_12_module_registry_integration_validation_acceptance`
  - `phase03_10_decision_center_formal_spec`
  - `phase03_14_decision_center_integration_validation_acceptance`
- Affected code:
  - 当前无代码改动；影响后续 `phase04` 正式规格正文编写、`Product / Repository / Binding` 实现与验收口径，以及 `phase05` 对当前主线的引用入口

## ADDED Requirements

### Requirement: Product / Repository / Binding 正式规格正文必须收口到单一文档

系统 SHALL 将 `phase04` 对应的 `Product / Repository / Binding` 正式规格收口到单一正文文档，而不是继续以 `phase04-01 ~ 09` 零散子规格作为长期并列入口。

#### Scenario: 正式规格正文文件落点

- **WHEN** 后续实现、验收或 `phase05` 需要引用 `Product / Repository / Binding` 当前阶段正式规格
- **THEN** 必须存在单一正文文件 `phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
- **AND** 该正文文件必须能够从当前 spec 目录直接追踪
- **AND** 不得要求接手者先手工拼装 `phase04-01 ~ 09` 后才能理解当前阶段正式口径

### Requirement: 正式规格正文覆盖范围必须完整

系统 SHALL 要求 `Product / Repository / Binding` 正式规格正文在同一处完整覆盖页面、动作、模板、绑定关系、数据读写、API、合同、验收基线、迁移边界、非目标、实现设计层结果与 Done 标准。

#### Scenario: 判断正式规格正文是否覆盖完整

- **WHEN** 接手者编写或审阅 `Product / Repository / Binding` 正式规格正文
- **THEN** 文档必须至少完整覆盖页面矩阵、动作矩阵、模板与状态语义、绑定关系与上下文入口、最小数据矩阵、最小 API 分组、`.proto` 合同章节、冷启动与异常路径前提、迁移与兼容边界、非目标矩阵、实现设计层结果与 Done 标准
- **AND** 不得只保留零散冻结结论、概要说明或“见各子 spec”式拼接文本

### Requirement: phase01-06、phase02、phase03 与 phase04-01~09 结论必须统一继承

系统 SHALL 要求 `Product / Repository / Binding` 正式规格正文完整继承 `phase01-06` 正式 MVP 规格正文、`phase02` 的正式规格与验收结果、`phase03` 的正式规格与验收结论，以及 `phase04-01 ~ 09` 已冻结的当前阶段结论，而不是重新定义第二套边界。

#### Scenario: 判断正式规格正文是否继承前置冻结

- **WHEN** 正式规格涉及技术路线、对象范围、页面职责、动作归属、数据边界、错误语义、合同边界、验收环境、迁移兼容或实现设计结果
- **THEN** 必须以 `mvp_spec_v0.1.md` 作为当前阶段唯一执行层上游
- **AND** 必须完整承接 `module_registry_spec_v0.1.md` 与 `phase02-12` 验收结论中对 `Module Registry` 的已交付边界
- **AND** 必须完整承接 `decision_center_spec_v0.1.md` 与 `phase03-14` 验收结论中对 `Decision Center` 的已交付边界
- **AND** 必须完整承接 `phase04-01 ~ 09` 已冻结的当前阶段结论
- **AND** 不得绕开既有子规格扩范围、改口径或引入第二套解释

### Requirement: 页面与动作章节必须单值化

系统 SHALL 要求正式规格正文在同一处完整承接 `Product Registry / List`、`Product Create`、`Product Detail`、`Repository Binding / List`、`Repository Create`、`Repository Binding Detail / Workspace` 六类页面主线及其动作归属。

#### Scenario: 判断页面与动作章节是否合格

- **WHEN** 正式规格定义页面结构与动作矩阵
- **THEN** 必须明确 `Product Registry / List` 承接产品列表读取、筛选入口、创建入口与进入详情入口
- **AND** 必须明确 `Product Create` 承接 `CreateProduct`
- **AND** 必须明确 `Product Detail` 承接详情读取、`BindModuleToProduct`、已绑定仓库摘要读取与进入仓库绑定主线的上下文入口
- **AND** 必须明确 `Repository Binding / List` 承接仓库列表读取、筛选入口、创建入口与进入工作台入口
- **AND** 必须明确 `Repository Create` 承接 `CreateRepository`
- **AND** 必须明确 `Repository Binding Detail / Workspace` 承接详情读取、`BindRepositoryToProduct` 与 `MapModuleToRepository`
- **AND** 必须明确 `Module Detail` 只承接兼容入口与上下文跳转，不扩写为第二个绑定工作台

### Requirement: 模板、状态与展示模型章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接 `Product` 与 `Repository` 的最小模板、字段级 required / optional、状态语义与最小读模型。

#### Scenario: 判断模板、状态与展示模型章节是否合格

- **WHEN** 正式规格定义 `Product / Repository` 模板、状态或展示模型
- **THEN** 必须明确 `Product` 最小模板为 `name / description / status`
- **AND** 必须明确 `Repository` 最小模板为 `name / url / provider / status`
- **AND** 必须明确上述字段在创建写入中的必填语义、去首尾空白后的非空校验与 `status` 显式提交规则
- **AND** 必须明确 `status` 最小枚举为 `active / archived`
- **AND** 必须明确列表与详情的最小读模型，以及 `module_bind_count / repository_bind_count / product_bind_count / module_bind_count` 的展示语义

### Requirement: 绑定关系、候选范围与入口上下文章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接三类绑定关系、候选范围、入口上下文与 canonical owner / reread 语义。

#### Scenario: 判断绑定关系、候选范围与入口上下文章节是否合格

- **WHEN** 正式规格定义三类绑定关系、候选读取、入口上下文与回流语义
- **THEN** 必须明确 `BindModuleToProduct` 的 canonical owner 为 `Product Detail`
- **AND** 必须明确 `BindRepositoryToProduct` 与 `MapModuleToRepository` 的 canonical owner 为 `Repository Binding Detail / Workspace`
- **AND** 必须明确三条候选读取的目标范围、排除规则与排序语义
- **AND** 必须明确 `Module Detail` 与 `Product Detail` 进入正式主入口时的上下文参数与目标身份参数拆分规则
- **AND** 必须明确三类绑定成功后的 reread 分别落到对应 canonical owner 页面

### Requirement: 数据、API 与错误语义章节必须完整承接

系统 SHALL 要求正式规格正文在同一处完整承接当前阶段的最小数据矩阵、读写范围、候选读取前提、接口分组与错误语义。

#### Scenario: 判断数据与 API 章节是否合格

- **WHEN** 正式规格定义数据模型方向或 API 边界
- **THEN** 必须明确直接承接 `products`、`repositories`、`product_modules`、`product_repositories`、`module_repositories`
- **AND** 必须明确 `Product Registry` 与 `Repository Binding` 的最小读组、写组与候选读取组
- **AND** 必须明确列表读取、详情读取、创建写入、绑定写入与候选读取的数据范围与最小输入输出前提
- **AND** 必须明确创建失败、资源不存在、目标非 `active`、重复绑定、列表空结果与候选空结果的错误语义归属
- **AND** 不得提前冻结超出当前阶段的 Dashboard 聚合接口或跨模块分析接口

### Requirement: 前端实现设计结果必须进入正式规格正文

系统 SHALL 要求正式规格正文显式承接 `phase04-05 / 06` 已冻结的前端实现设计层结果，使后续前端实现不需要再次回到多个子规格拆读。

#### Scenario: 判断前端实现设计结果是否被承接

- **WHEN** 正式规格定义前端实现设计层结果
- **THEN** 必须明确页面集合、最小路由结构、URL 语义、组件树与组件职责
- **AND** 必须明确列表查询条件在路由搜索参数层的唯一事实源语义
- **AND** 必须明确创建页来源上下文、草稿状态、提交状态、取消返回与成功回流规则
- **AND** 必须明确详情页候选读取状态、绑定动作状态、reread 规则与多入口回流规则
- **AND** 必须明确单一 `React Web` 下同时覆盖 `PC` 与移动浏览器的布局降级策略

### Requirement: 后端实现设计结果必须进入正式规格正文

系统 SHALL 要求正式规格正文显式承接 `phase04-07` 已冻结的后端实现设计层结果，使后续后端实现可直接按既定边界进入编码。

#### Scenario: 判断后端实现设计结果是否被承接

- **WHEN** 正式规格定义后端实现设计层结果
- **THEN** 必须明确 `Product Registry` 与 `Repository Binding` 的后端模块边界
- **AND** 必须明确 `handler / service / repository / candidate` 的分层语义与关键文件落点
- **AND** 必须明确四条关系摘要读取链路与三条候选读取链路的 owner、接口形态与装配原则
- **AND** 必须明确 `phase02` 旧候选读取与旧绑定写入口只允许作为兼容适配层，不再保留第二 owner
- **AND** 不得把 Go HTTP/RPC 框架、ORM、SQL Builder 或数据访问层具体工具写成当前正式规格阻断项

### Requirement: 合同章节必须完整承接 Proto 主线

系统 SHALL 要求正式规格正文显式承接 `phase04-08` 已冻结的 `.proto` 合同设计结果，使后续合同落地成为机械映射。

#### Scenario: 判断合同章节是否合格

- **WHEN** 正式规格定义 `Product / Repository / Binding` 合同边界
- **THEN** 必须明确 `.proto` 是当前阶段唯一合同源
- **AND** 必须明确 `psco.common.v1`、`psco.product_registry.v1`、`psco.repository_binding.v1` 的包名、最小服务矩阵、核心消息结构与 RPC -> HTTP 映射
- **AND** 必须明确共享 `ActiveArchivedStatus` 枚举、跨包 `ModuleStatus` 复用、字段编号与 `reserved` 约束
- **AND** 必须明确合同演进遵循现有 Buf workspace 主线与 `buf build / lint / generate / breaking` 校验链

### Requirement: 验收基线与迁移边界章节必须完整承接

系统 SHALL 要求正式规格正文显式承接 `phase04-09` 已冻结的联调环境、重置脚本、基线 seed、冷启动路径、兼容入口与迁移边界。

#### Scenario: 判断验收基线与迁移边界章节是否合格

- **WHEN** 正式规格定义验收环境、迁移兼容与联调基线
- **THEN** 必须明确环境建立顺序、`0006_product_repository_binding_mainline.sql`、`reset_product_repository_mainline.sh` 与 `seed_product_repository_mainline_baseline.sql` 的职责
- **AND** 必须明确 `products / repositories` 从 `phase02` 只读前提升级为 `phase04` 正式主线实体的迁移边界
- **AND** 必须明确从空状态到首个 `Product`、首个 `Repository` 与首轮三类绑定的冷启动路径
- **AND** 必须明确 `Module Detail` 旧入口兼容、旧 transport 委派与多入口回流矩阵
- **AND** 必须明确当前阶段验收不依赖临时手工 SQL

### Requirement: 非目标与 Done 标准章节必须存在

系统 SHALL 要求正式规格正文显式包含当前阶段非目标矩阵与正式 Done 标准，而不是只保留功能边界。

#### Scenario: 判断非目标与 Done 标准章节是否存在

- **WHEN** 正式规格进入验收阶段
- **THEN** 文档必须明确 `Feature / Opportunity / Experiment` 主线、Dashboard 聚合分析、Decision 对 Product / Repository 的正式主线接入、自动扫描代码、自动知识图谱、独立 `React Native` 客户端、完整 `PWA` 与完整自动化测试平台建设为当前阶段非目标
- **AND** 必须明确何时可判定 `Product / Repository / Binding` 当前阶段规格足以进入实现与后续 `phase05`
- **AND** 不得缺失验收口径或把 Done 标准分散到多个文档中

### Requirement: 正式规格正文必须成为后续直接上游

系统 SHALL 将 `Product / Repository / Binding` 正式规格正文定位为后续实现、验收与 `phase05` 判断进入条件时的直接上游规格来源。

#### Scenario: 判断后续阶段是否引用直接上游

- **WHEN** 后续实现、验收或 `phase05` 需要引用 `Product / Repository / Binding` 当前阶段边界
- **THEN** 必须优先引用 `product_repository_binding_spec_v0.1.md`
- **AND** 不得继续并列引用多个 `phase04-01 ~ 09` 子规格作为长期正式入口
- **AND** 不得绕开正式规格正文直接以对话结论替代规格入口

### Requirement: 正式规格正文必须与根级真相源和阶段文档互链一致

系统 SHALL 要求正式规格正文与根级真相源、`phase04` 三件套及 `phase02 / phase03` 已交付结果保持单值一致，不得形成第二套正式真相源。

#### Scenario: 判断正式规格正文与真相源关系

- **WHEN** 正式规格引用项目定位、技术基线、目录规则、最终共识、已交付模块边界或当前阶段范围
- **THEN** 必须与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致
- **AND** 必须与 `phase04_product_and_repository_binding_foundation_architecture_plan.md`、`phase04_product_and_repository_binding_foundation_dev_plan.md`、`phase04_product_and_repository_binding_foundation_shared_baseline.md` 保持单值一致
- **AND** 必须与 `module_registry_spec_v0.1.md`、`phase02-12` 验收结论、`decision_center_spec_v0.1.md` 与 `phase03-14` 验收结论保持单值一致
- **AND** 不得在正式规格正文中重写一套与上述文档冲突的主结论

## MODIFIED Requirements

### Requirement: phase04 后续实现与 phase05 的规格入口

后续 `Product / Repository / Binding` 的实现、验收与 `phase05` 对当前主线的引用，SHALL 以本次产出的 `product_repository_binding_spec_v0.1.md` 作为直接上游规格入口；而该正文自身继续以 `mvp_spec_v0.1.md`、`module_registry_spec_v0.1.md`、`phase02-12` 验收结论、`decision_center_spec_v0.1.md` 与 `phase03-14` 验收结论作为当前阶段已冻结上游。

#### Scenario: 后续阶段引用正式规格

- **WHEN** 后续阶段需要引用 `Product / Repository / Binding` 当前阶段正式规格
- **THEN** 必须从 `product_repository_binding_spec_v0.1.md` 进入
- **AND** 该正文仍必须显式继承 `mvp_spec_v0.1.md`、`module_registry_spec_v0.1.md`、`phase02-12` 验收结论、`decision_center_spec_v0.1.md` 与 `phase03-14` 验收结论的已冻结边界
- **AND** 不得把 `phase01-06`、`phase02`、`phase03` 与 `phase04-01 ~ 09` 并列改写成多个长期直接入口

## REMOVED Requirements

### Requirement: 以后续实现继续直接拼装 phase04-01~09 子规格
**Reason**: `phase04-10` 的目标就是把前九个子任务的冻结结论收口为单一正式规格正文，避免后续实现与验收继续依赖多文档临场拼装。
**Migration**: 后续实现、验收与 `phase05` 若需要引用 `Product / Repository / Binding` 当前阶段边界，应统一从 `product_repository_binding_spec_v0.1.md` 进入；前置子规格仅作为该正文的冻结来源与追踪依据。
