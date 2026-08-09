# Phase05-10 Dashboard + Feedback 正式规格正文收敛 Spec

## Why

`phase05-01` 到 `phase05-09` 已分别冻结了 `Dashboard` 的页面边界、反馈信号模板、跳转返回、聚合读错误语义、前端页面/状态模型、后端模块边界、`.proto` 合同设计与验收基线，但当前仍缺少一份可直接作为实现与后续阶段上游的正式规格正文。若继续让实现直接并列引用九份子规格，`phase05` 就会重复 `phase02/03/04` 之前“局部 spec 都是对的，但执行入口不够单一”的老问题。

## What Changes

- 冻结 `phase05` 正式规格正文的唯一文档落点与文件命名
- 冻结该正式正文必须完整收敛 `phase05-01` 到 `phase05-09` 的结论，不另立第二套边界
- 冻结正式正文的章节骨架，必须完整覆盖页面、区块、聚合读、反馈信号、API、合同、验收基线、非目标与 Done 标准
- 冻结该正文与 `phase01-06` 正式 MVP 规格、`phase04-10` 正式规格正文、`phase04-14` 验收结论的互链前提
- 冻结“正式正文成为直接执行层上游、前九个子 spec 退为追溯来源”的角色切换规则
- **BREAKING**：`phase05` 后续实现、实现复核与下一阶段规格，不得再把 `phase05-01 ~ 09` 当作并列直接执行层入口使用

## Impact

- Affected specs:
  - `phase01_06_formal_mvp_spec`
  - `phase04_10_product_repository_binding_formal_spec`
  - `phase04_14_product_repository_binding_integration_validation_acceptance`
  - `phase05_01_dashboard_pages_info_arch`
  - `phase05_02_feedback_signal_priority_display_model`
  - `phase05_03_dashboard_navigation_context_return_path`
  - `phase05_04_dashboard_aggregate_api_error_boundary`
  - `phase05_05_dashboard_frontend_page_route_component_design`
  - `phase05_06_dashboard_frontend_state_interaction_flow`
  - `phase05_07_dashboard_feedback_backend_module_boundary_interface_grouping`
  - `phase05_08_define_dashboard_feedback_proto_contract`
  - `phase05_09_design_dashboard_acceptance_baseline_fixtures`
- Affected code:
  - `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`（已产出正式正文）
  - 后续 `phase05-11 / 12 / 13 / 14` 实现、联调与验收任务都必须以该正式正文为直接规格入口

## ADDED Requirements

### Requirement: phase05 正式正文必须成为 Dashboard + Feedback 的唯一规格入口

系统 SHALL 为 `Dashboard + Feedback` 产出一份正式规格正文，使其成为 `phase05` 后续实现、验收与下一阶段引用时的唯一直接执行层入口。

#### Scenario: 正式正文落点与命名

- **WHEN** 后续实现或下一阶段文档引用 `phase05` 的正式规格正文
- **THEN** 必须引用 `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
- **AND** 文档标题必须明确为 `PSCO Dashboard + Feedback 规格 v0.1 — 正式规格正文` 或等价单值表达
- **AND** 不得把 `phase05-01 ~ 09` 中任意一份子规格继续当作并列正式正文使用

#### Scenario: 正式正文的文档定位

- **WHEN** 编写 `dashboard_feedback_spec_v0.1.md`
- **THEN** 开头必须明确以下三类定位：
  - 文档定位：它是 `phase05_dashboard_feedback_foundation` 的正式规格正文
  - 上游收敛：它由 `phase05-01 ~ 09` 的冻结结论收敛而成
  - 互链前提：它与 `AGENTS.md`、`plan.md`、`TECH_STACK_BASELINE.md`、`project_rules.md`、`architecture_map.md`、`PSCO-summarize-feedback.md` 保持单值一致
- **AND** 必须显式说明 `phase04` 已完成收口，而本文档继续作为 `Dashboard + Feedback` 的已交付正式规格入口使用

### Requirement: 正式正文必须完整收敛前九个子任务的冻结结果

系统 SHALL 将 `phase05-01` 到 `phase05-09` 的冻结结果收口到同一份正式正文中，避免实现期再去并列拼接多个 spec。

#### Scenario: 收敛来源矩阵

- **WHEN** 正式正文声明其直接来源
- **THEN** 必须显式承接以下九个子规格：
  - `phase05-01` 页面边界与信息结构
  - `phase05-02` 反馈信号模板、优先级与最小展示模型
  - `phase05-03` 跳转目标、来源上下文与返回路径
  - `phase05-04` 聚合读范围、接口边界与错误语义
  - `phase05-05` 前端页面、路由与组件分层
  - `phase05-06` 前端状态模型与交互流
  - `phase05-07` 后端模块边界与接口分组
  - `phase05-08` 最小 `.proto` 合同设计
  - `phase05-09` 联调验收环境、冷启动基线与 fixture 设计
- **AND** 不得遗漏其中任意一项，使正式正文失去直接承接实现的完整性

#### Scenario: 子规格到正式正文的角色切换

- **WHEN** `phase05-10` 完成后讨论 `phase05-01 ~ 09` 的角色
- **THEN** 它们必须退回为正式正文的追溯来源与证据链
- **AND** 正式正文必须承担后续实现与引用时的单一入口职责
- **AND** 不得要求执行者继续并列阅读九份子规格才能知道正式范围

### Requirement: 正式正文的章节骨架必须覆盖 DoD 的全部范围

系统 SHALL 将 `phase05-10` 的正式正文章节骨架冻结为足以完整覆盖 dev plan DoD 的单值结构。

#### Scenario: 必备章节矩阵

- **WHEN** 编写 `dashboard_feedback_spec_v0.1.md`
- **THEN** 至少必须覆盖以下章节：
  - `技术路线`
  - `对象范围`
  - `页面矩阵`
  - `区块矩阵`
  - `动作矩阵`
  - `数据模型`
  - `聚合读与反馈信号`
  - `跳转与返回上下文`
  - `前端页面/状态模型`
  - `后端模块边界与接口分组`
  - `API 边界与合同矩阵`
  - `冷启动、fixture 与验收基线`
  - `非目标`
  - `Done 标准`
- **AND** 不得把页面、区块、聚合读、反馈信号、API、合同、验收基线、非目标与 Done 标准拆散到正文之外的第二入口文档

#### Scenario: 页面与区块正文收敛

- **WHEN** 正式正文描述 `Dashboard Home`
- **THEN** 必须继续冻结四个区块的固定归属：`dashboard_overview / Current Focus / Asset Feedback / Recent Activity`
- **AND** 必须继续冻结 `Current Focus / Next Action` 为第一屏唯一主行动队列
- **AND** 必须继续冻结 `Dashboard Home` 的正式业务入口路由为 `/dashboard`
- **AND** 不得把 `/` 重新解释为 `Dashboard Home`

#### Scenario: 聚合读、反馈信号与 CTA 正文收敛

- **WHEN** 正式正文描述 Dashboard 的读取、反馈与 CTA
- **THEN** 必须继续冻结三类读取：`DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead`
- **AND** 必须继续冻结反馈优先级：`pending_decision_signals > product missing both bindings > product missing repository binding > product missing module binding`
- **AND** 必须继续冻结 CTA 1-9 的唯一命中顺序与单主 CTA 约束
- **AND** 必须继续冻结 `hidden / suppressed / ready / computing` 的主 CTA 状态语义或等价单值解释

#### Scenario: API、合同与验收基线正文收敛

- **WHEN** 正式正文描述 API、合同与验收基线
- **THEN** 必须继续冻结 `.proto` 为唯一合同源，`chi + JSON HTTP` 只作为显式映射承接
- **AND** 必须继续冻结 `proto/psco/dashboard/v1/dashboard.proto` 与 `DashboardService`
- **AND** 必须继续冻结 `buf build / lint / generate / breaking` 校验链与 `reserved`/编号演进规则
- **AND** 必须继续冻结 `reset_dashboard_acceptance.sh`、九类 fixture、局部错误环境变量入口与返回链路验收矩阵

### Requirement: 正式正文必须与 phase01-06、phase04 正式文档和验收结论互链一致

系统 SHALL 确保 `phase05` 正式正文不是孤立新文档，而是明确承接 `phase01-06` 与 `phase04` 已完成交付的互链规格。

#### Scenario: 与 phase01-06 的互链关系

- **WHEN** 正式正文声明上游规格关系
- **THEN** 必须明确 `mvp_spec_v0.1.md` 是当前阶段唯一执行层总上游
- **AND** 必须说明 `Dashboard` 继续服务 `v0.1` 的主线目标：资产状态可见、决策留痕可见、基础复用反馈可见
- **AND** 不得把 `phase05` 扩写为偏离 `v0.1` 范围的新产品定义

#### Scenario: 与 phase04-10 的互链关系

- **WHEN** 正式正文声明直接执行层上游
- **THEN** 必须明确承接 `product_repository_binding_spec_v0.1.md`
- **AND** 必须继续尊重 `Product Detail`、`Repository Binding Detail / Workspace`、`Module Detail`、`Decision Detail / Decision Center` 的 canonical owner 边界
- **AND** 不得把 Dashboard 解释为第二套写入主线或第二套绑定工作台

#### Scenario: 与 phase04-14 的互链关系

- **WHEN** 正式正文声明验收前提与回流规则
- **THEN** 必须显式继承 `phase04-14` 已验证通过的真实前后端、真实数据库、真实 reread 与返回路径结论
- **AND** 必须将其解释为 `phase05` 的直接运行前提，而不是重新验证 `phase04` 主线是否存在
- **AND** 不得在 `phase05` 正式正文中回退 `phase04-14` 已收口的 canonical owner、返回路径与 reread 语义

### Requirement: 正式正文必须显式写出非目标与 Done 标准

系统 SHALL 在 `phase05` 正式正文中显式写出当前阶段的非目标边界与 Done 标准，避免实现期把 Dashboard 扩成第二产品线。

#### Scenario: 非目标矩阵

- **WHEN** 正式正文进入非目标章节
- **THEN** 至少必须继续排除以下内容：
  - 第二套写入主线
  - 复杂驾驶舱 / BI 分析
  - 外部遥测、通知中心与自动消息回流
  - GitHub OAuth / 自动导入
  - 独立 `AI Assistant` 一级导航
  - 独立 `React Native` 客户端
  - 完整 `PWA`
- **AND** 不得把这些长期方向误写成当前阶段既成事实

#### Scenario: Done 标准矩阵

- **WHEN** 正式正文定义 `phase05` Done 标准
- **THEN** 至少必须覆盖以下达成条件：
  - 页面、区块与路由边界完整
  - 三类聚合读与反馈信号语义完整
  - 跳转、返回与刷新恢复语义完整
  - 后端模块边界与 `.proto` 合同边界完整
  - 验收环境、fixture 与局部错误基线完整
  - 非目标、验收证据要求与下一阶段引用前提完整
- **AND** Done 标准必须足以支撑后续 `phase05-11 ~ 14` 的实现、联调、验收与收口

## MODIFIED Requirements

### Requirement: phase05 共享基线中的“进入 /spec”解释

`phase05_dashboard_feedback_foundation_shared_baseline.md` 中“当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步”的要求，在 `phase05-10` SHALL 被进一步解释为：`phase05` 的 `/spec` 不再停留在并列子规格冻结，而必须收口为一份正式规格正文，作为实现与下一阶段的直接上游。

#### Scenario: /spec 收口方式

- **WHEN** `phase05-10` 完成后讨论 `/spec` 是否已经形成正式入口
- **THEN** 必须以 `dashboard_feedback_spec_v0.1.md` 为判定标准
- **AND** 不得仅以 `phase05-01 ~ 09` 子规格齐全就视为已经形成正式规格正文

## REMOVED Requirements

### Requirement: phase05 后续实现并列依赖 phase05-01 ~ 09 子规格

**Reason**: `phase05-10` 的目标就是将分散的冻结结论收口为正式正文入口。若后续实现仍需要并列依赖九份子规格，正式正文就失去了“唯一直接执行层上游”的意义。
**Migration**: 后续实现、实现复核、验收与下一阶段文档应默认先引用 `dashboard_feedback_spec_v0.1.md`；`phase05-01 ~ 09` 继续保留为追溯来源、证据链与设计演化记录，不再承担并列直接执行层入口职责。
