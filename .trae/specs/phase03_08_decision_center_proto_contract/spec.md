# phase03-08 Decision Center 最小 Protocol Buffers 合同设计 Spec

## Why

`phase03-04` 已经把 `Decision Center` 的最小数据读写范围、接口边界与错误语义冻结为单值结论，`phase03-07` 又把后端模块边界、接口分组与文件落点收敛到了可直接进入实现的粒度。当前仍缺少一份 `Decision Center` 的最小 `.proto` 合同设计，用来回答“哪些消息进入合同源、服务接口如何命名、包名与版本如何冻结、字段编号如何分配、以及 `.proto` 与 `chi + JSON HTTP` 过渡传输层如何保持显式映射”，否则后续实现仍可能继续围绕手写 JSON DTO 扩张，把 `Protocol Buffers` 的落地成本与兼容性成本后移。

## What Changes

- 为 `Decision Center` 新增最小 `.proto` 合同设计，覆盖当前阶段全部必须动作与候选读取
- 冻结 `package`、版本语义、文件落点、消息结构、字段编号、枚举与服务接口
- 冻结 `DecisionListRead / DecisionDetailRead / DecisionWrite / DecisionLinkWrite / DecisionModuleCandidateRead` 的 request / response 语义
- 冻结 `.proto` 与 `chi + JSON HTTP` 过渡传输层之间的显式映射策略
- 冻结合同演进规则，包括 `reserved` 约束与 `buf breaking` 前提
- **BREAKING**：后续 `Decision Center` 的实现与验收不得再把手写 JSON 结构视为并列合同源，`.proto` 成为当前阶段唯一合同源

## Impact

- Affected specs:
  - `phase03_04_decision_data_api_error_boundary`
  - `phase03_07_backend_module_boundary_interface_grouping`
  - 后续 `phase03-10` 正式规格正文
- Affected code:
  - 预期新增 `proto/psco/decision_center/v1/decision_center.proto`
  - 预期补充 `proto/README.md` 中关于 `Decision Center` 的合同说明与 RPC → HTTP 映射
  - 预期约束 `backend/internal/decisioncenter/types.go`
  - 预期约束前端 `decision-center` 过渡传输层适配代码

## ADDED Requirements

### Requirement: Decision Center 最小 Proto 合同源

系统 SHALL 为 `Decision Center` 落地单一 `.proto` 合同源，作为当前阶段唯一合同定义入口。

#### Scenario: 合同源落地

- **WHEN** 执行 `phase03-08`
- **THEN** 仓库中必须存在可追踪的 `.proto` 文件落点
- **AND** 该 `.proto` 文件必须覆盖 `DecisionListRead / DecisionDetailRead / DecisionWrite / DecisionLinkWrite / DecisionModuleCandidateRead`
- **AND** 不得再以手写 JSON 结构充当并列合同源

### Requirement: Proto 包名与版本语义

系统 SHALL 为当前阶段 `Decision Center` 的 Proto 合同冻结明确的包名、版本号与文件组织方式。

#### Scenario: 版本与目录冻结

- **WHEN** 新增 `Decision Center` 的 `.proto` 合同源
- **THEN** 必须冻结单一包名与版本语义为 `psco.decision_center.v1`
- **AND** 必须冻结稳定目录落点为 `proto/psco/decision_center/v1/decision_center.proto`
- **AND** 后续新增字段必须在该版本演进规则下进行，而不是临时改写包名

### Requirement: 最小消息结构必须覆盖当前动作矩阵

系统 SHALL 让 Proto 消息结构完整承接 `phase03-02`、`phase03-03`、`phase03-04`、`phase03-06` 与 `phase03-07` 已冻结的最小读写模型。

#### Scenario: 读组消息结构

- **WHEN** 定义 `DecisionListRead` 与 `DecisionDetailRead`
- **THEN** 必须覆盖 `DecisionListItem`、`Decision`、`DecisionDetail`、已关联 `Module` 结果与最小来源上下文
- **AND** `DecisionDetailRead` 不得内嵌 `Decision -> Module` 候选读取结果
- **AND** `DecisionModuleCandidateRead` 必须通过独立 request / response 承接候选读取

#### Scenario: 写组消息结构

- **WHEN** 定义 `DecisionWrite` 与 `DecisionLinkWrite`
- **THEN** 必须覆盖 `RecordDecision` 与 `LinkDecisionToTarget`
- **AND** 必须保持与 `phase03-02` 的模板字段、`status` 枚举与校验语义一致
- **AND** `LinkDecisionToTarget` 当前阶段唯一允许的正式目标类型必须是 `Module`

### Requirement: 核心消息字段语义必须单值映射上游读模型

系统 SHALL 将 `Decision Center` `.proto` 合同的核心消息字段语义单值映射到 `phase03-02 / 03 / 04` 已冻结的读模型与写模型，避免合同落地期出现字段语义漂移或实现者自行决定字段结构。本 Requirement 冻结字段语义与建模约束；具体字段编号方案、写组 response 返回语义与跨包 enum 依赖策略由本 spec 独立 Requirement `字段编号方案与合同边界策略冻结` 承接，不留待 `phase03-11` 实现期决定。

#### Scenario: 核心读模型消息字段语义来源

- **WHEN** 定义 `Decision`、`DecisionListItem`、`DecisionDetail` 与 `LinkedModule` 消息
- **THEN** `Decision` 必须覆盖 `id / title / context / problem / alternatives / choice / reason / impact / status / created_at`，对齐 `phase03-02` 最小结构化模板
- **AND** `DecisionListItem` 必须覆盖 `id / title / status / created_at / link_count / linked_module_summary`，对齐 `phase03-02` 列表读模型冻结与 `shared_baseline §5.1` 计算口径
- **AND** `DecisionDetail` 必须覆盖 `Decision` 全部字段、已关联 `Module` 列表与 `SourceContext`，对齐 `phase03-02` 详情读模型
- **AND** `LinkedModule` 必须覆盖 `module_id / module_name`，对齐 `phase03-02` `linked_module_summary` 计算口径的生成基础
- **AND** `linked_module_summary` 必须建模为 `string`，承接 `phase03-02` 的 `+N` 摘要语义，不得拆成结构化列表字段

#### Scenario: alternatives 合同建模约束

- **WHEN** 在 `.proto` 中定义 `alternatives` 字段
- **THEN** 必须建模为 `repeated string`
- **AND** 不得引入嵌套 `message`（如 `label / score / vote / source`）
- **AND** 必须保持按输入顺序保留的语义，对齐 `phase03-02 alternatives` 最小结构冻结

#### Scenario: SourceContext 字段结构冻结

- **WHEN** 定义 `SourceContext` 消息
- **THEN** 必须包含 `source_module_id` 与 `source_module_name` 两个字段
- **AND** 该字段语义对齐 `phase03-03` 入口上下文“至少携带当前 `Module` 的目标标识与可展示名称”
- **AND** 当 `Decision` 从 `Decision Center / List` 直接发起时，`SourceContext` 必须能表达“无特定来源目标”语义（`source_module_id` 与 `source_module_name` 均为空字符串）
- **AND** `SourceContext` 不得内嵌完整 `Module` 对象，只承接最小来源标识
- **AND** 本 Scenario 同时收口 `phase03-02` 显式后移的“来源上下文字段结构由 `phase03-03` 承接”前提

#### Scenario: 候选读取消息字段语义

- **WHEN** 定义 `DecisionModuleCandidate` 消息
- **THEN** 必须覆盖 `module_id / module_name / status`，对齐 `shared_baseline §5.1` 候选读取排序基线
- **AND** `status` 字段必须复用 `ModuleStatus` 语义（`active / archived`），不得引入第二套状态枚举
- **AND** 候选读取排序语义（`active` 优先 -> `module_name` 升序，已关联不得再次出现）由 `service/repository` 层承接，不进入 `.proto` 合同本体

#### Scenario: 写组 request 字段语义冻结

- **WHEN** 定义 `CreateDecisionRequest` 与 `LinkDecisionToTargetRequest`
- **THEN** `CreateDecisionRequest` 必须覆盖 `title / context / problem / alternatives / choice / reason / impact / status`，对齐 `phase03-02` 模板字段与必填规则
- **AND** `LinkDecisionToTargetRequest` 必须包含 `decision_id / target_type / module_id`，对齐 `phase03-03` 目标范围与 `phase03-04` 错误语义
- **AND** `LinkDecisionToTargetRequest` 的 `target_type` 当前阶段只允许 `MODULE`，不得接受 `PRODUCT / REPOSITORY`

#### Scenario: 字段语义与页面动作单值映射

- **WHEN** 校验合同字段语义与页面动作的一致性
- **THEN** `ListDecisions` 响应字段必须对应 `Decision Center / List` 页面展示字段
- **AND** `GetDecisionDetail` 响应字段必须对应 `Decision Detail` 页面展示字段
- **AND** `CreateDecision` request 字段必须对应 `Decision Create` 页面表单字段
- **AND** `LinkDecisionToTarget` request 字段必须对应 `Decision Detail` 页面关联写入动作
- **AND** `ListDecisionModuleCandidates` 响应字段必须对应 `Decision Detail` 候选读取面板展示字段

### Requirement: 枚举与字段语义必须单值化

系统 SHALL 为 `Decision Center` 的 `.proto` 合同冻结稳定的枚举定义与字段语义。

#### Scenario: status 枚举冻结

- **WHEN** 定义 `DecisionStatus`
- **THEN** 必须覆盖 `proposed / active / superseded / archived`
- **AND** 必须保留 `UNSPECIFIED` 作为零值
- **AND** JSON 过渡层中的字符串语义必须与该枚举单值对应

#### Scenario: link target type 枚举冻结

- **WHEN** 定义 `DecisionLinkTargetType`
- **THEN** 当前阶段只允许 `MODULE`
- **AND** 必须保留 `UNSPECIFIED` 作为零值
- **AND** 不得把 `Product / Repository` 提前写成当前阶段正式可用枚举值

### Requirement: 字段编号与兼容性规则

系统 SHALL 为 `Decision Center` 的 `.proto` 合同建立可演进的字段编号规则，避免后续演进破坏兼容性。

#### Scenario: 字段编号分配

- **WHEN** 为消息定义字段
- **THEN** 必须为每个字段分配稳定编号
- **AND** 不得复用已删除字段编号或字段语义
- **AND** 必须为未来演进保留合理空间

#### Scenario: 删除字段后的 reserved 约束

- **WHEN** 后续版本删除字段或废弃字段名
- **THEN** 必须使用 `reserved` 保留字段号
- **AND** 必要时同时保留字段名
- **AND** 不得在未声明 `reserved` 的情况下复用旧编号或旧名称

### Requirement: 字段编号方案与合同边界策略冻结

系统 SHALL 为 `Decision Center` `.proto` 合同冻结具体字段编号方案、写组 response 最小返回语义与跨包 enum 依赖策略，使 `phase03-11` `.proto` 落地成为机械映射，不在实现期产生分叉。本 Requirement 与 `字段编号与兼容性规则` Requirement 配合：后者冻结演进规则（`reserved` / 不复用），前者冻结当前编号方案与合同边界策略，共同满足 `phase03-08` DoD "合同字段语义、字段编号与页面动作单值一致"。

#### Scenario: 枚举字段编号冻结

- **WHEN** 定义 `DecisionStatus` 与 `DecisionLinkTargetType` 枚举
- **THEN** `DecisionStatus` 必须按以下编号冻结：
  - `DECISION_STATUS_UNSPECIFIED = 0`
  - `DECISION_STATUS_PROPOSED = 1`
  - `DECISION_STATUS_ACTIVE = 2`
  - `DECISION_STATUS_SUPERSEDED = 3`
  - `DECISION_STATUS_ARCHIVED = 4`
- **AND** `DecisionLinkTargetType` 必须按以下编号冻结：
  - `DECISION_LINK_TARGET_TYPE_UNSPECIFIED = 0`
  - `DECISION_LINK_TARGET_TYPE_MODULE = 1`
- **AND** 后续新增枚举值必须使用递增编号，不得插入到已有编号之间

#### Scenario: 核心对象字段编号冻结

- **WHEN** 定义 `Decision` / `DecisionListItem` / `LinkedModule` / `SourceContext` / `DecisionDetail` / `DecisionModuleCandidate` 消息
- **THEN** 各消息字段编号必须按以下方案冻结：
  - `Decision`：`id=1(string)` / `title=2(string)` / `context=3(string)` / `problem=4(string)` / `alternatives=5(repeated string)` / `choice=6(string)` / `reason=7(string)` / `impact=8(string)` / `status=9(DecisionStatus)` / `created_at=10(google.protobuf.Timestamp)`
  - `DecisionListItem`：`id=1(string)` / `title=2(string)` / `status=3(DecisionStatus)` / `created_at=4(google.protobuf.Timestamp)` / `link_count=5(int32)` / `linked_module_summary=6(string)`
  - `LinkedModule`：`module_id=1(string)` / `module_name=2(string)`
  - `SourceContext`：`source_module_id=1(string)` / `source_module_name=2(string)`
  - `DecisionDetail`：`decision=1(Decision)` / `linked_modules=2(repeated LinkedModule)` / `source_context=3(SourceContext)`
  - `DecisionModuleCandidate`：`module_id=1(string)` / `module_name=2(string)` / `status=3(ModuleStatus, 跨包 import)`
- **AND** 后续新增字段必须使用新的递增编号，不得复用已删除字段编号

#### Scenario: 读组 request / response 字段编号冻结

- **WHEN** 定义 `ListDecisions` / `GetDecisionDetail` / `ListDecisionModuleCandidates` 的 request 与 response
- **THEN** 各消息字段编号必须按以下方案冻结：
  - `ListDecisionsRequest`：`query_text=1(string)` / `status_filter=2(DecisionStatus)`
  - `ListDecisionsResponse`：`decisions=1(repeated DecisionListItem)`
  - `GetDecisionDetailRequest`：`decision_id=1(string)`
  - `GetDecisionDetailResponse`：`decision_detail=1(DecisionDetail)`
  - `ListDecisionModuleCandidatesRequest`：`decision_id=1(string)`
  - `ListDecisionModuleCandidatesResponse`：`candidates=1(repeated DecisionModuleCandidate)`

#### Scenario: 写组 request / response 字段编号与返回语义冻结

- **WHEN** 定义 `CreateDecision` / `LinkDecisionToTarget` 的 request 与 response
- **THEN** `CreateDecisionRequest` 字段编号必须按以下方案冻结：
  - `title=1(string)` / `context=2(string)` / `problem=3(string)` / `alternatives=4(repeated string)` / `choice=5(string)` / `reason=6(string)` / `impact=7(string)` / `status=8(DecisionStatus)`
- **AND** `LinkDecisionToTargetRequest` 字段编号必须按以下方案冻结：
  - `decision_id=1(string)` / `target_type=2(DecisionLinkTargetType)` / `module_id=3(string)`
- **AND** `CreateDecisionResponse` 必须只包含 `decision_id=1(string)`，对齐 `phase03-07` "返回新建 `Decision` 标识以支持前端回流到 `DecisionDetailPage`"
- **AND** `LinkDecisionToTargetResponse` 必须为空响应（无字段），对齐 `phase03-07` "详情页重新读取当前已关联目标结果"的回流语义
- **AND** 不得在 `LinkDecisionToTargetResponse` 中返回 link 结果或 detail reread 标识，避免形成脱离 `DecisionDetailRead` 的第二套回流读取路径

#### Scenario: ModuleStatus 跨包依赖策略冻结

- **WHEN** 定义 `DecisionModuleCandidate.status` 的枚举类型
- **THEN** 必须通过 `import "psco/module_registry/v1/module_registry.proto"` 直接复用 `psco.module_registry.v1.ModuleStatus`
- **AND** 不得在 `psco.decision_center.v1` 中重定义本地等价枚举，避免引入两套 `status` 语义
- **AND** 该 import 仅限于复用 `ModuleStatus` 枚举类型，不得因此引入对 `Module Registry` service 或其他消息结构的合同依赖
- **AND** 理由：`ModuleStatus` 是 `Module` 的固有属性，候选读取返回的 `status` 必须与 `Module Registry` 单值一致；重定义本地等价枚举会违反 `shared_baseline §6` "不允许出现并列定义"；`module_registry.proto` 已是仓库内稳定合同源（`phase02-11A` 已落地）

### Requirement: 服务接口必须对齐已冻结接口分组

系统 SHALL 将 Proto 服务接口与 `phase03-04 / 07` 已冻结的读写接口分组保持单值一致，不引入第二套命名体系。

#### Scenario: 最小服务接口矩阵

- **WHEN** 定义 `Decision Center` 的 Proto service
- **THEN** 最小服务接口至少应包含 `ListDecisions`、`GetDecisionDetail`、`CreateDecision`、`LinkDecisionToTarget`、`ListDecisionModuleCandidates`
- **AND** `ListDecisions` 对齐 `DecisionListRead`
- **AND** `GetDecisionDetail` 对齐 `DecisionDetailRead`
- **AND** `CreateDecision` 对齐 `DecisionWrite`
- **AND** `LinkDecisionToTarget` 对齐 `DecisionLinkWrite`
- **AND** `ListDecisionModuleCandidates` 对齐 `DecisionModuleCandidateRead`

### Requirement: 过渡传输层必须从 Proto 单向承接

系统 SHALL 明确当前阶段 `chi + JSON HTTP` 与 `.proto` 的关系，避免形成第二套合同源。

#### Scenario: 过渡层保留

- **WHEN** 当前阶段继续保留 `chi + JSON HTTP`
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 语义显式对齐
- **AND** 不得在 `handler`、HTTP DTO 或前端 adapter 中私自新增 `.proto` 中不存在的字段语义
- **AND** 不得把 HTTP 路径、状态码或中间件策略误写成 Proto 合同本体

#### Scenario: 路径参数与消息字段映射

- **WHEN** HTTP 过渡层使用 URL 路径参数承接 `decisionId` 或 `moduleId`
- **THEN** handler 必须在进入业务层前显式组装为对应的 Proto request 字段
- **AND** 该差异必须被视为传输层差异，而不是并列合同定义

### Requirement: RPC 到 HTTP 的映射矩阵必须明确

系统 SHALL 为当前阶段的最小动作矩阵冻结单值的 RPC → HTTP 映射矩阵。

#### Scenario: 最小映射矩阵

- **WHEN** 设计 `Decision Center` 的过渡传输层映射
- **THEN** 至少必须明确以下映射：
- **AND** `ListDecisions` -> `GET /api/decisions`
- **AND** `GetDecisionDetail` -> `GET /api/decisions/{decisionId}`
- **AND** `CreateDecision` -> `POST /api/decisions`
- **AND** `LinkDecisionToTarget` -> `POST /api/decisions/{decisionId}/links`
- **AND** `ListDecisionModuleCandidates` -> `GET /api/decisions/{decisionId}/candidates/modules`

### Requirement: 详情读取与候选读取必须保持消息边界分离

系统 SHALL 保持 `DecisionDetailRead` 与 `DecisionModuleCandidateRead` 在合同层的边界分离，不得为了图省事把两类读取并成单条消息。

#### Scenario: Detail 与 Candidate 边界

- **WHEN** 定义 `GetDecisionDetailResponse` 与 `ListDecisionModuleCandidatesResponse`
- **THEN** `GetDecisionDetailResponse` 只承接详情本体、最小来源上下文与已建立关联结果
- **AND** `ListDecisionModuleCandidatesResponse` 只承接候选 `Module` 列表
- **AND** 不得把候选结果直接塞进详情响应中

### Requirement: 错误语义必须在合同设计中保留显式映射前提

系统 SHALL 在 `.proto` 合同设计中为当前阶段已冻结的错误语义保留稳定承接前提，但不把 HTTP 状态码本身写进 `.proto`。

#### Scenario: 错误语义承接

- **WHEN** 设计 `RecordDecision`、`DecisionDetailRead`、`DecisionModuleCandidateRead`、`LinkDecisionToTarget` 的合同边界
- **THEN** 必须显式保留校验失败、资源不存在、重复冲突与空结果语义的承接空间
- **AND** `DecisionModuleCandidateRead` 的空候选结果必须表现为正常空列表响应，不得设计为空错误
- **AND** HTTP 状态码映射继续由过渡传输层负责

### Requirement: 生成链与校验链必须冻结

系统 SHALL 明确 `Decision Center` 当前阶段的最小合同落地链路，避免把合同设计膨胀成完整通信栈改造。

#### Scenario: 当前阶段合同落地边界

- **WHEN** 执行 `phase03-08`
- **THEN** 必须冻结 `.proto` 合同源与最小生成入口约定
- **AND** 当前阶段可以不完成完整 gRPC / Connect / 网关迁移
- **AND** 当前阶段可以保留 `chi` 作为 HTTP 过渡传输层

#### Scenario: buf 校验链

- **WHEN** 冻结当前阶段合同工具链
- **THEN** 必须至少覆盖 `buf build`、`buf lint`、`buf generate`、`buf breaking`
- **AND** `buf breaking` 必须直接对照仓库主线 `.git` 基准
- **AND** 不得吞掉 `buf breaking` 的失败退出码

## MODIFIED Requirements

### Requirement: Contract First 当前阶段解释

`Contract First` 在 `phase03` 当前阶段 SHALL 被解释为“先冻结最小数据读写范围、接口边界与错误语义，再将其落地为最小 `.proto` 合同源”，而不是继续停留在长期方向。

#### Scenario: phase03 合同口径更新

- **WHEN** `phase03-08` 进入执行链
- **THEN** `Decision Center` 的 `.proto` 必须成为当前阶段唯一合同源
- **AND** 后续 `phase03-10` 正式规格正文与 `phase03-11` 实现必须基于该合同源展开
- **AND** 不得继续沿用“当前阶段只冻结接口边界、不要求 proto 设计”的旧口径

## REMOVED Requirements

### Requirement: 当前阶段只保留 Decision Center 的 JSON DTO 设计
**Reason**: 该口径会导致 `Decision Center` 的前后端继续围绕手写 JSON DTO 扩张，把字段编号、消息边界与兼容性约束继续后移。
**Migration**: 改为“当前阶段必须冻结 `Decision Center` 最小 `.proto` 合同设计；`chi + JSON HTTP` 继续作为过渡传输层，通过显式映射承接 `.proto` 语义”。
