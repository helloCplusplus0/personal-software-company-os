# phase12-03 冻结只读消费深化边界与更重通道进入条件 Spec

## Why
`phase12-01` 已冻结主交付边界，`phase12-02` 已冻结 Web 端四实体语义承接矩阵，但如果不继续把“只读消费到底可以深化到哪里、共享只读 owner 如何分层、`repository_id` 在不同页面如何承接、哪些能力必须留到后续阶段”正式冻结，后续 `phase12-05 ~ 07` 仍会在设计时反复补判断。

## What Changes
- 冻结 `phase11` 已交付只读项目上下文能力在 `phase12` 可深化的正式范围
- 冻结 Web / agent 共享只读语义的 owner 分层与承接边界
- 冻结 `repository_id` 在直接 `repository-scoped`、间接 `repository-scoped` 与衍生消费页三类页面中的承接规则
- 冻结“复用既有只读主线 vs 新增最小受控承接位”的判定边界
- 冻结更重消费通道、受控维护能力与额外专家讨论的进入条件
- 对齐 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中关于只读消费深化的表达

## Impact
- Affected specs: `phase12_semantic_alignment_and_readonly_consumption_foundation`
- Affected code:
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/*`
  - `frontend/src/features/project-context/*`
  - `frontend/src/features/*/data/use-*-read.ts`

## ADDED Requirements

### Requirement: 冻结只读消费深化的正式边界
系统规格 SHALL 冻结 `phase11` 已交付只读项目上下文能力在 `phase12` 可以如何深化，并要求所有后续设计继续保持“先消费、后维护”的单一主线。

当前阶段至少固定如下边界：

- 只读消费深化继续以 `phase11` 已交付的 `.proto + ConnectRPC + ProjectContextService` 为正式基础；
- 深化目标只包括更稳定、可定位、可复用的共享只读语义摘要、规则入口、文档入口与最小解释结果；
- 允许在现有只读主线内演进最小字段、受控摘要视图或受控 adapter，但不得借机打开写回、审批流、MCP / CLI 或第二套业务协议出口；
- Web 与 agent 共享同一套只读事实源，不允许前端本地再拼一套“页面版真相源”；
- 后续执行者不得再临场决定“这里该新建前端共享读模型，还是直接拼页面解释”，必须先回答是否已满足本阶段冻结的共享承接条件。

#### Scenario: 执行者判断某只读增强是否属于 phase12
- **WHEN** 执行者准备新增只读字段、共享摘要或消费入口
- **THEN** 能直接判断该增强是否属于 `phase12` 的只读消费深化
- **AND** 不会把更重的维护能力或新协议层误写成当前阶段默认范围

### Requirement: 冻结共享只读 owner 分层
系统规格 SHALL 冻结后端 canonical owner、Web 跨切片共享只读 owner 与切片内展示 owner 的单值分层，避免只读消费深化时长出并列事实源。

当前阶段至少固定如下 owner：

1. **结构化 canonical owner**
   - `ProjectContextService.GetProjectContext`
   - `.proto` 与 `backend/internal/projectcontext/*`
   - 承接项目上下文聚合只读结果、规则/约束/文档入口定位与四实体最小摘要事实
2. **agent-facing Markdown owner**
   - `ProjectContextService.ExportProjectContext`
   - 只承接 agent-facing Markdown 导出，不承接 Web 页面专属真相源职责
3. **Web 跨切片共享只读 owner**
   - `frontend/src/features/project-context/`
   - 只允许承接受控只读 adapter、query options、共享语义摘要与入口定位视图
   - 启用条件：当 `3+` 页面或切片稳定复用同一份 repository-scoped 语义摘要、规则入口或文档入口时
4. **切片内展示 owner**
   - 各 feature 自己的 `pages/`、`components/` 与 `data/`
   - 只负责把共享只读结果映射为页面文案、摘要卡片、空态和说明文案

补充冻结：

- `frontend/src/features/project-context/` 不得承接写路径、页面私有状态或并列 canonical 字段语义；
- 切片内不得重新拼装跨切片共享语义摘要；若达到稳定复用条件，必须回收到唯一共享只读 owner；
- 若现有 `GetProjectContext` 已足够承接，则前端必须直接复用其结构化结果或受控派生视图，不得先长出影子摘要合同。

#### Scenario: 执行者判断共享只读 owner 落点
- **WHEN** 后续设计需要为 Web 或 agent 提供共享只读解释
- **THEN** 能机械判断该结果应落在 canonical owner、Markdown owner、跨切片共享只读 owner 还是切片内展示 owner
- **AND** 不会在前端页面或临时脚本中再长出第二套事实源

### Requirement: 冻结 repository-scoped 承接边界矩阵
系统规格 SHALL 冻结 `repository_id` 在不同消费页面中的承接边界，使后续执行者无需临场决定“直接传入、先解析回 repository_id，还是只能消费派生摘要”。

当前阶段至少固定如下矩阵：

1. **直接 repository-scoped 页面**
   - `repositories/$repositoryId`
   - 允许直接消费 `GetProjectContext(repository_id)`
2. **间接 repository-scoped 详情页**
   - `products/$productId`
   - `modules/$moduleId`
   - `decisions/$decisionId`
   - 必须先通过稳定解析链路回到同一 `repository_id`，再复用共享只读语义摘要
3. **衍生消费页**
   - `dashboard`
   - `onboarding`
   - `reviews/daily`
   - `reviews/weekly`
   - 只允许消费受控共享摘要、固定入口链接或派生只读解释，不得伪装成新的结构化主锚点入口

补充冻结：

- `repository_id` 是唯一正式结构化输入锚点；
- `product_id / module_id / decision_id` 只允许从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析；
- 不允许通过页面内临场搜索、额外脚本、手工数据库查询或工作区扫描补齐锚点；
- 任何新增 Web 读路径都必须先回答自己属于三类页面中的哪一类，否则不得进入实现。

#### Scenario: 执行者为某页面设计只读承接链路
- **WHEN** 执行者设计某个页面的只读读取或共享摘要接入
- **THEN** 能明确该页面属于直接、间接或衍生消费页
- **AND** 能据此判断是否允许直接使用 `repository_id`、必须先解析回 `repository_id`，或只能消费派生摘要

### Requirement: 冻结复用与最小新增的判定边界
系统规格 SHALL 冻结“继续复用 `phase11` 既有只读主线”与“新增最小受控承接位”的判定边界，避免后续设计把局部便利扩张成第二服务或第二事实源。

当前阶段至少固定如下判定规则：

- 若现有 `GetProjectContext` 已能承接所需共享事实，应优先复用，而不是新增第二服务；
- 若确需新增字段或视图，唯一合法身份是 `ProjectContextService` 下的受控派生读取，且必须说明回收了哪段前端 / agent 重复解释逻辑；
- 面向 Web 的共享只读 adapter 必须保留 `repository_id`、`entry_ref`、`entry_kind` 或其单向派生定位关系；
- 任何新增承接位都必须能在验收记录中说明其输入锚点、输出结果与失败条件；
- 不允许把“前端先临时拼摘要，后面再治理”视为当前阶段合法过渡。

#### Scenario: 执行者评估是否需要新增只读视图
- **WHEN** 现有只读主线无法直接承接某个共享摘要需求
- **THEN** 执行者必须先判断是否可在 `ProjectContextService` 下增加最小受控派生读取
- **AND** 不会直接新增并列业务域、第二套 API 或前端影子摘要合同

### Requirement: 冻结更重通道与受控维护能力进入条件
系统规格 SHALL 冻结更重消费通道、受控维护能力与额外专家讨论的进入条件，使这些能力只作为后续阶段候选，而不是在 `phase12` 内被顺手吸收。

当前阶段至少固定如下进入条件：

- 只有在 `phase12` 正式完成语义一致性收口、共享只读 owner 单值化、固定样本与固定入口验收闭环后，才允许讨论更重通道；
- 更具体地，必须同时满足：固定 `6` 问验收闭环、固定 `repository_id` 锚点解析协议闭环、固定样本解析协议闭环，三者缺一不可；
- 新能力若包含 `MCP / CLI / agent 写回 / Draft / 审批流 / 前端对话式入口 / 第二套 canonical API / 影子状态表` 任一项，默认判定为超出当前阶段范围；
- 若确有进入后续讨论的必要，必须先显式说明：为什么现有只读消费已不足、为什么不能继续通过受控只读深化解决、以及它准备承接哪类新增职责；
- 额外专家讨论只允许围绕已压缩候选方向展开，不重开一轮广义方案发散。

#### Scenario: 团队讨论引入更重消费通道
- **WHEN** 后续执行者提出引入更重消费通道或受控维护能力
- **THEN** 能直接判断这些内容不属于 `phase12-03` 的正式范围
- **AND** 只有满足冻结进入条件后，才允许作为后续阶段候选增强继续讨论

## MODIFIED Requirements

### Requirement: phase12 三件套中的只读消费深化表达
`phase12` 三件套中的只读消费深化表达 SHALL 对齐为同一口径：

- 只读消费深化继续复用 `phase11` 的正式只读主线
- `ProjectContextService.GetProjectContext / ExportProjectContext / frontend/src/features/project-context/` 的 owner 层级保持单值一致
- `repository_id` 是唯一结构化输入锚点，并按直接 / 间接 / 衍生三类页面承接
- 更重消费通道与受控维护能力只保留为后续进入条件，不得写成当前事实

不得再出现：

- 一个文档允许前端本地拼接并列共享摘要，另一个文档又要求单一事实源；
- 一个文档允许页面临场决定锚点解析方式，另一个文档又要求三类页面承接矩阵；
- 一个文档把更重通道写成“当前顺手增强”，另一个文档又将其冻结为后续条件。

#### Scenario: 三件套对齐只读消费深化边界
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 能获得同一套只读消费深化边界、owner 分层、锚点承接矩阵与进入条件
- **AND** 后续 `/spec` 与实现不需要再次补充第二套边界解释

## REMOVED Requirements

### Requirement: 将更重消费通道视为当前阶段默认可并入能力
**Reason**: `phase12-03` 的目标是先冻结只读消费深化边界，而不是提前把更重维护能力、协议层或对话入口并入当前阶段。
**Migration**: 后续设计与实现必须先在当前 spec 冻结的只读边界内回答共享 owner、锚点承接与复用判定；超出部分统一转为后续阶段候选增强。
