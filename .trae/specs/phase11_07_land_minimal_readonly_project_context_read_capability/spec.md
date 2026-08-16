# 落实最小只读项目上下文聚合读取能力 Spec

## Why
`phase11-05` 已经冻结了最小只读项目上下文导出的输入锚点、聚合边界与输出职责，但如果没有一个正式承接位把这些设计落实为可调用的只读读取能力，agent 仍然无法稳定消费当前项目上下文。`phase11-07` 需要把“设计边界”推进为“最小可用的只读能力”，同时继续守住 `.proto + ConnectRPC`、只读聚合投影与非固定目录合同三条主线。

## What Changes
- 落实一个最小只读“项目上下文聚合读取”正式承接位
- 冻结该能力的输入合同、输出边界与失败语义
- 明确该能力继续复用现有 `.proto + ConnectRPC` 主线，不引入第二套 canonical API 或新协议层
- 明确该能力只依赖 `repository_id` 与其在 `PSCO` 中已登记的 canonical 关系，不依赖消费侧项目目录结构
- 冻结 `phase11-07` 的成功标准、DoD 与收口口径

## Impact
- Affected specs:
  - `phase11_project_context_foundation`
  - `phase11_05_design_minimal_readonly_project_context_export`
- Affected code:
  - `docs/phase/phase11_project_context_foundation_dev_plan.md`
  - `docs/phase/phase11_project_context_foundation_shared_baseline.md`
  - `docs/phase/phase11_project_context_foundation_architecture_plan.md`
  - backend 只读业务接口正式承接位（后续实现）
  - `.proto + ConnectRPC` 业务合同主线（后续实现）

## ADDED Requirements

### Requirement: 提供最小只读项目上下文聚合读取承接位
系统规格 SHALL 落实一个最小只读“项目上下文聚合读取”正式承接位，用于让 agent 按单一稳定合同读取当前项目核心上下文。

补充冻结：

- 该承接位的身份是聚合只读读取能力，而不是新实体主线；
- 该承接位必须复用既有 `.proto + ConnectRPC` 正式主线；
- 该承接位不得偷渡 MCP、CLI、前端对话式入口或任何第二协议层；
- 该承接位的实现目标是让执行者不必再临场决定“current project”如何绑定。

#### Scenario: 成功建立最小只读承接位
- **WHEN** 后续执行者开始落实 `phase11-07`
- **THEN** 能明确存在一个正式只读承接位用于读取项目上下文
- **AND** 不会把该能力误做成前端页面查询拼装、临时脚本或第二套 API

### Requirement: 冻结唯一输入锚点与失败语义
系统规格 SHALL 继续冻结 `repository_id` 为当前阶段唯一正式结构化输入锚点，并明确未绑定仓库的失败态与返回语义。

补充冻结：

- 当前阶段不允许把本地路径、Git remote URL、`product_id` 或工作区扫描升格为并列主锚点；
- 当前阶段只承接“已完成 Repository Binding”的仓库上下文读取；当前阶段将“绑定完成”明确解释为：目标 `Repository` 至少已有一条 `product_repositories` 绑定，且至少已有一条 `module_repositories` 映射；
- 若目标仓库不存在或尚未完成 `Repository Binding`，必须返回明确失败态，而不是由执行者自行补猜绑定关系；
- 失败语义必须属于正式合同的一部分，而不是散落在实现说明中的口头约定。

#### Scenario: 成功固定输入与失败态
- **WHEN** 后续执行者定义项目上下文读取合同
- **THEN** 能直接判断 `repository_id` 是唯一正式结构化输入锚点
- **AND** 能直接判断未绑定仓库的失败态与返回语义

### Requirement: 保持只读聚合投影与 canonical 一致性
系统规格 SHALL 冻结该能力为基于既有 canonical 数据的只读聚合投影，不得引入第二套业务事实源。

补充冻结：

- 该能力只读取已登记的 `Repository / Product / Module / Decision` canonical 关系；
- `Decision` 聚合口径必须继续遵守 `phase11-05` 已冻结的两类 module-link 派生命中范围；
- 返回结果只承接结构化只读字段边界内的信息；
- 不允许在读取侧额外拼装与 backend canonical contracts 并列的第二套字段语义、状态机或事实源。

#### Scenario: 成功保持只读聚合边界
- **WHEN** 后续执行者实现项目上下文聚合读取
- **THEN** 返回结果仍然只是既有 canonical 数据上的只读投影
- **AND** 不会因为导出方便而长出第二套业务事实源

### Requirement: 不以消费侧目录结构作为输入前提
系统规格 SHALL 明确该能力不要求消费侧项目目录与 `PSCO` 当前仓库拥有相同结构，也不要求存在固定文件名作为必要输入合同。

补充冻结：

- 当前阶段的通用项目上下文能力以 `repository_id` 与其在 `PSCO` 中已绑定的 canonical 关系为准；
- `PSCO` 当前仓库中的根级文件清单只用于自身治理与第一轮 dogfooding，不上升为未来所有项目的必备目录模板；
- 即使未来出现最佳实践项目模板，其身份也只能是增强型 convention/profile，而不是 `phase11-07` 的前置依赖。

#### Scenario: 成功避免目录结构合同化
- **WHEN** 后续执行者为 agent 说明该能力的使用前提
- **THEN** 不需要要求消费侧项目与 `PSCO` 当前仓库拥有相同文件布局
- **AND** 不会把 `README.md / AGENTS.md / rules` 等固定文件名误写成读取前置条件

### Requirement: 冻结输入合同、输出边界与一致性说明
系统规格 SHALL 为该最小只读读取能力冻结正式产物边界，至少包括输入合同、输出边界、失败语义与 canonical 一致性说明。

正式产物至少包括：

- 一个最小结构化只读读取承接位；
- 对应的输入合同、输出边界与失败语义；
- 与既有 canonical 数据的一致性说明。

补充冻结：

- 规则、约束与文档入口字段必须提供可定位入口，而不只是摘要文本；
- 当前阶段允许以 `entry_ref + entry_kind` 一类定位字段承接入口，不强制固定为 path/url/id 某一种单一形态。

#### Scenario: 成功形成可直接实施的正式产物
- **WHEN** 后续执行者进入实现准备或实现验收
- **THEN** 能直接知道应该交付哪些正式结果
- **AND** 不需要再次争论“到底是设计说明完成了，还是正式读取能力已经落地了”

### Requirement: 冻结 phase11-07 成功标准与收口口径
系统规格 SHALL 冻结 `phase11-07` 的成功标准、DoD 与收口口径，使后续执行者无需再临场判断“最小只读上下文能力做到什么程度才算完成”。

当前子任务的成功标准至少固定为：

1. 已存在可供 agent 消费的最小只读项目上下文能力；
2. 该能力的只读边界、输入合同、输出边界与失败语义已形成正式承接结果；
3. `repository_id` 仍是唯一正式结构化输入锚点，执行者不需要再临场决定“current project”如何绑定；
4. 该能力未引入 agent 写回、第二套 canonical API、新协议层或消费侧固定目录合同；
5. 该能力与既有 canonical 数据的一致性已可直接说明与验证。

当前子任务的收口口径至少固定为：

- `phase11-07` 的完成标志不是“有一段读取逻辑能跑”，而是最小只读项目上下文能力已经有正式承接位、正式合同和正式失败语义；
- 若仍依赖执行者临场补猜当前项目绑定关系、仍依赖消费侧目录结构、仍引入第二套 API 或仍缺少明确失败态，则 `phase11-07` 不得判定为完成。

#### Scenario: 成功判断子任务是否完成
- **WHEN** 后续执行者进入实现验收、handoff 或阶段收口
- **THEN** 可以根据正式承接位、合同边界与失败语义直接判断 `phase11-07` 是否达标
- **AND** 不需要重新争论“最小只读能力是不是已经正式落地”

## MODIFIED Requirements

### Requirement: phase11 最小只读项目上下文能力从设计进入实现承接
`phase11` 中关于“最小只读项目上下文导出”的表达 SHALL 从设计冻结推进为最小结构化只读读取能力的正式承接。

#### Scenario: 设计与实现承接闭环
- **WHEN** 读者对照 `phase11-05` 与 `phase11-07`
- **THEN** 能看到前者负责冻结导出设计与边界，后者负责落实最小只读读取能力
- **AND** 不会把 `phase11-07` 误解为重新讨论导出边界，而不是开始正式承接实现

## REMOVED Requirements

### Requirement: 允许执行者用本地路径或目录结构临场补锚 current project
**Reason**: 这会直接破坏 `repository_id` 作为唯一正式结构化输入锚点的冻结口径，并把通用项目上下文能力重新绑回当前仓库文件布局。
**Migration**: 统一通过 `repository_id` 与既有 canonical 关系完成项目上下文读取；若未绑定，则返回正式失败态。
