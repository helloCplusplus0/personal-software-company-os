# 最小只读项目上下文聚合导出设计 Spec

## Why
`phase11` 已经把 `PSCO` 冻结为“项目上下文系统”，但如果不把最小只读项目上下文导出的输入锚点、聚合范围、结构化输出边界与 Markdown 导出职责边界正式冻结，后续 `/spec` 与实现仍会回到“到底从哪里读、聚合到什么程度、哪些属于结构化输出、哪些属于渲染层”的临场判断。`phase11-05` 需要把这组导出设计冻结成可直接承接的单值输入。

## What Changes
- 冻结项目上下文聚合只读读取的输入锚点与失败语义设计
- 冻结聚合内容范围、边界与与现有 canonical 数据对齐的投影设计
- 冻结结构化只读输出与 AGENTS 风格 Markdown 导出的职责边界
- 冻结结构化读取到 Markdown 导出的单向派生关系
- 冻结当前阶段不承接的协议、写路径与消费侧目录依赖
- 冻结 `phase11-05` 的成功标准、DoD 与收口口径

## Impact
- Affected specs: `phase11_project_context_foundation`
- Affected code:
  - `docs/phase/phase11_project_context_foundation_dev_plan.md`
  - `docs/phase/phase11_project_context_foundation_architecture_plan.md`
  - `docs/phase/phase11_project_context_foundation_shared_baseline.md`
  - `backend/`
  - `proto/`
  - `frontend/`

## ADDED Requirements

### Requirement: 冻结最小只读导出的输入锚点与失败语义
系统规格 SHALL 将当前阶段项目上下文聚合只读读取的输入锚点冻结为 `repository_id`，并同时冻结失败语义口径。

补充冻结：

- `repository_id` 是当前阶段唯一正式结构化输入锚点；
- 当前阶段不以本地路径、Git remote URL、`product_id` 或工作区扫描作为并列主锚点；
- 当前阶段只承接“已完成 Repository Binding”的仓库上下文读取；当前阶段将“绑定完成”明确解释为：目标 `Repository` 至少已有一条 `product_repositories` 绑定，且至少已有一条 `module_repositories` 映射；
- 仓库不存在或仓库绑定不完整都必须返回明确失败态，不允许执行者自行补猜锚点或补猜绑定关系。

#### Scenario: 成功判断如何读取当前项目上下文
- **WHEN** 后续执行者设计项目上下文聚合读取入口
- **THEN** 能直接判断 `repository_id` 是唯一正式结构化输入锚点
- **AND** 不会把本地路径、目录扫描或其他身份对象扩张成并列主锚点

### Requirement: 冻结聚合内容范围与 canonical 投影边界
系统规格 SHALL 冻结当前阶段项目上下文聚合导出的最小内容范围，并要求其继续对齐 `PSCO` 已登记的 canonical 关系与现有正式主线。

当前阶段至少应聚合：

- 当前 `Repository` 身份；
- 关联 `Product` 摘要；
- 关联 `Module` 摘要与状态；
- 关联 `Decision` 摘要与状态；
- 与当前项目直接相关的规则与约束入口；
- 与当前 phase 直接相关的文档入口。

补充冻结：

- 当前阶段的通用能力以 `PSCO` 中已登记的 `Repository / Product / Module / Decision` canonical 关系为准；
- 当前阶段不把消费侧项目目录中的固定文件名或固定目录结构当作必要输入合同；
- 当前阶段不把“统一项目模板”作为前置依赖。

#### Scenario: 成功判断上下文聚合到什么程度
- **WHEN** 后续执行者为最小只读导出设计聚合投影
- **THEN** 能明确当前阶段至少需要聚合哪些 canonical 内容
- **AND** 不会把聚合范围扩张成全库扫描、知识图谱或消费侧目录扫描

### Requirement: 冻结 Decision 聚合口径
系统规格 SHALL 冻结 `Decision` 的聚合范围、去重规则与导出过滤规则。

`Decision` 聚合口径固定为：

- 以当前 `Repository` 为根，只合并基于既有 `Decision -> Module` canonical link 可直接投影出的两类命中：
  - 命中“当前 `Repository` 已映射 `Module`”的 `Decision`
  - 命中“当前 `Repository` 已绑定 `Product` 所属 `Module`”的 `Decision`
- 当前阶段不得把 `Repository` 或 `Product` 伪装成 `Decision` 的直接 link target，也不得继续沿 `Product -> Module -> 其他 Repository` 做递归扩张；超出上述两类命中范围的 `Decision` 不进入当前阶段导出；
- 同一 `Decision` 若同时命中多类关系，必须以 `decision_id` 去重，并保留命中来源摘要；
- 当前阶段结构化只读主列表只承接非 `archived` 的 `Decision`；`archived` 不进入主导出列表。

#### Scenario: 成功判断哪些 Decision 应进入导出
- **WHEN** 后续执行者实现或继续细化 `Decision` 聚合读取
- **THEN** 能直接依据固定口径判断哪些 `Decision` 进入当前阶段导出
- **AND** 不需要临场决定是否递归扩张、是否混入 `archived` 或保留哪条重复命中记录

### Requirement: 冻结结构化只读输出与 Markdown 导出职责边界
系统规格 SHALL 冻结结构化只读输出与 AGENTS 风格 Markdown 导出的职责边界，使二者不形成并列事实源。

补充冻结：

- 结构化只读读取继续落在 Go backend 的 `.proto + ConnectRPC` 正式主线；
- 结构化只读输出负责承接输入锚点、失败语义、聚合字段边界与 canonical 关系结果；
- Markdown 导出负责从同一结构化读取结果单向派生为 AGENTS 风格或等价 Markdown 风格文档；
- Markdown 导出不形成第二套事实源，不反向决定结构化输出字段语义。

#### Scenario: 成功判断哪些字段属于结构化输出、哪些属于 Markdown 渲染
- **WHEN** 后续执行者设计导出接口、字段边界或渲染层
- **THEN** 能明确哪些内容属于结构化只读输出，哪些内容属于 Markdown 导出渲染
- **AND** 不会让 Markdown 导出反向长出第二套字段语义或事实源

### Requirement: 冻结结构化读取到 Markdown 的单向派生关系
系统规格 SHALL 冻结“结构化只读读取 -> Markdown 导出”的单向派生关系说明，使后续设计与实现不能绕开结构化主线直接拼装第二套导出输入。

#### Scenario: 阻断第二套导出主线
- **WHEN** 后续执行者讨论 AGENTS 风格导出如何生成
- **THEN** 应直接判断 Markdown 导出必须从同一结构化只读结果单向派生
- **AND** 不得通过根级静态文档拼装、消费侧目录扫描或 agent 临时补猜形成并列导出主线

### Requirement: 冻结当前阶段明确不做的协议与写路径
系统规格 SHALL 冻结 `phase11-05` 当前明确不承接的协议层与写路径能力。

当前阶段明确不做：

- `MCP / CLI` 协议层；
- agent 写回、审批流或主动注入；
- 前端对话式入口；
- 任何第二套 canonical API 或第二套事实源；
- 把消费侧项目目录中的固定文件名/固定目录结构写成必要输入合同。

#### Scenario: 阻断协议与写路径偷渡
- **WHEN** 后续执行者讨论导出能力如何扩展
- **THEN** 能直接判断哪些协议层与写路径不属于 `phase11-05`
- **AND** 不会把这些能力以“顺手补齐导出体验”之名写成当前阶段事实

### Requirement: 冻结 phase11-05 成功标准与收口口径
系统规格 SHALL 冻结 `phase11-05` 的成功标准、DoD 与子任务收口口径，使后续执行者无需再临场判断“上下文到底聚合到什么程度才算完成”。

当前子任务的成功标准至少固定为：

1. 输入锚点、失败语义、聚合内容范围、`Decision` 聚合口径与结构化/Markdown 职责边界已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结；
2. `repository_id` 已被固定为唯一正式结构化输入锚点，且未绑定仓库失败语义已可直接引用；
3. 结构化只读输出字段边界、Markdown 导出字段边界与“结构化读取 -> Markdown 单向派生”关系已单值化；
4. 当前阶段不依赖消费侧项目目录结构、不依赖统一项目模板、也不承接协议层与写路径扩张的边界已写清；
5. 后续执行者可以直接据此进入 `/spec`，而不需要再猜“哪些属于聚合、哪些属于渲染、上下文到底聚合到什么程度”。

当前子任务的收口口径至少固定为：

- `phase11-05` 的完成标志不是“新增了一份 spec 包”，而是三件套已经完成最小只读项目上下文导出设计的正式冻结并保持单值一致；
- 若三件套之间仍存在输入锚点冲突、聚合范围漂移、`Decision` 聚合规则不一致、或 Markdown 导出反向长出第二套事实源的空间，则 `phase11-05` 不得判定为完成。

#### Scenario: 成功判断子任务是否完成
- **WHEN** 后续执行者进入后续 `/spec`、实现、验收或 handoff
- **THEN** 可以依据三件套中的固定完成条件直接判断 `phase11-05` 是否达标
- **AND** 不需要重新解释输入、输出与职责边界

## MODIFIED Requirements

### Requirement: phase11 三件套中的项目上下文导出表达
`phase11` 三件套中的项目上下文导出表达 SHALL 对齐为同一口径：

- `repository_id` 是唯一正式结构化输入锚点；
- 最小导出继续对齐 `Repository / Product / Module / Decision` canonical 关系；
- 结构化只读读取继续落在 `.proto + ConnectRPC` 正式主线；
- Markdown 导出从同一结构化结果单向派生；
- 当前阶段不依赖消费侧目录结构，也不承接协议与写路径扩张。

#### Scenario: 三件套单值一致
- **WHEN** 读者分别查看 `architecture_plan`、`dev_plan` 与 `shared_baseline`
- **THEN** 获取到的输入锚点、聚合范围、导出职责边界与非目标口径一致
- **AND** 不会出现某个文档允许目录扫描或第二套事实源，而另一个文档排除的冲突表达

## REMOVED Requirements

### Requirement: 将消费侧项目目录结构视为最小导出的必要输入合同
**Reason**: `phase11` 当前目标是先建立基于 canonical 关系的最小只读项目上下文能力，而不是把消费侧项目目录结构提前冻结为硬前提。
**Migration**: 将所有“依赖固定文件名、固定目录结构或统一项目模板才能读取项目上下文”的表达收敛为未来候选增强，不在 `phase11-05` 的正式范围内承接。
