# phase12-09 落实只读消费深化与共享只读入口 Spec

## Why
`phase11` 已经交付最小只读项目上下文读取与导出能力，`phase12-05 / 06 / 07 / 08` 已分别冻结共享入口、前端读路径 owner、后端合同边界与前端语义表达收口。`phase12-09` 需要把这些设计真正落到共享只读入口与消费链路上，让 Web 与 agent 复用同一套 project-context 事实源，而不是各自继续拼第二套解释结果。

## What Changes
- 冻结 `phase12-09` 为“落实共享只读入口与只读消费深化”的实现型子任务，而不是新一轮边界讨论
- 落实 `frontend/src/features/project-context/` 作为唯一合法的 Web 跨切片共享只读承接位
- 落实 Web / agent 共用的规则、约束、文档入口与最小摘要消费结果，优先复用既有 `GetProjectContext / ExportProjectContext`
- 落实必要的共享只读 adapter、query options、入口定位 view model 与消费接入，不引入第二服务、第二事实源或新协议层
- 对齐 detail pages、dashboard、onboarding、review 与 agent 导出对同一 project-context 事实源的消费方式

## Impact
- Affected specs:
  - `phase12_semantic_alignment_and_readonly_consumption_foundation`
  - `phase11_07_land_minimal_readonly_project_context_read_capability`
  - `phase11_08_land_agents_style_project_context_export`
  - `phase12_05_design_readonly_consumption_shared_entry`
  - `phase12_06_design_frontend_read_path_owner_shared_summary_reread`
  - `phase12_07_design_backend_contract_export_shared_read_views`
  - `phase12_08_land_frontend_four_entity_semantic_alignment`
- Affected code:
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/*`
  - `frontend/src/features/project-context/*`
  - `frontend/src/features/product-registry/*`
  - `frontend/src/features/repository-binding/*`
  - `frontend/src/features/module-registry/*`
  - `frontend/src/features/decision-center/*`
  - `frontend/src/features/dashboard/*`
  - `frontend/src/features/onboarding/*`
  - `frontend/src/features/review/*`

## ADDED Requirements

### Requirement: 冻结 phase12-09 的实现职责与非目标
系统 SHALL 将 `phase12-09` 冻结为“共享只读入口与只读消费深化”的实现型子任务，只消费 `phase12-05 / 06 / 07 / 08` 已冻结结论，不反向改写它们。

当前子任务至少必须完成：

- 让 `phase11` 的 project-context 能力在 `phase12` 下形成更稳定、可复用、可定位的共享只读入口；
- 让 Web 与 agent 共用同一套规则、约束、文档入口与最小摘要事实源；
- 让 Web 侧跨切片共享只读入口只落在 `frontend/src/features/project-context/`；
- 让 detail pages、dashboard、onboarding、review 与 agent 导出不再各自拼第二套解释性结果。

补充冻结：

- `phase12-09` 不引入写回、Draft、审批流、第二服务或新协议层；
- `phase12-09` 不重新讨论 `phase12-05 / 06 / 07` 已冻结的 owner 分层、resolver 规则与后端候选去留；
- `phase12-09` 若需要补充实现文件，只能沿既有 `ProjectContextService` 与 `frontend/src/features/project-context/` 主线继续承接。

#### Scenario: 执行者开始承接 phase12-09
- **WHEN** 执行者开始实现 `phase12-09`
- **THEN** 能明确这是共享只读入口落地任务
- **AND** 不会把结构性重构或新协议扩张偷渡进来

### Requirement: 落实 Web 唯一跨切片共享只读入口
系统 SHALL 将 `frontend/src/features/project-context/` 落实为唯一合法的 Web 跨切片共享只读承接位，并把稳定复用的 project-context 查询封装、共享语义来源与入口定位视图收敛到该路径下。

当前阶段至少必须落实：

- 基于 `repository_id` 的共享只读 query options / read hook；
- 四实体共享语义来源；
- `entry_ref / entry_kind / label / summary / status_summary` 的受控入口定位 view model；
- `frontend/src/features/project-context/` 的对外导出入口。

补充冻结：

- `frontend/src/features/project-context/` 不得承接写路径、页面私有状态或第二套 canonical facts；
- 不允许在 `dashboard / onboarding / review / detail page` 各自的数据层再保留一套跨切片共享解释逻辑；
- 页面仍然负责本地渲染结构，L3 只负责共享只读查询与 adapter。

#### Scenario: 页面需要消费共享只读 project-context
- **WHEN** 某个页面需要复用 project-context 摘要、规则入口或共享语义来源
- **THEN** 只会从 `frontend/src/features/project-context/` 接入
- **AND** 不会在页面私有 data 层临时再拼一套共享结果

### Requirement: 落实 Web / agent 共用的 project-context 事实源
系统 SHALL 让 Web 与 agent 继续共用 `ProjectContextService.GetProjectContext` 与 `ProjectContextService.ExportProjectContext` 所承载的同一事实源，而不是分别维护两套解释结果。

当前阶段至少必须落实：

- Web 共享只读消费优先复用 `GetProjectContext` 的真实结构化结果；
- agent 导出继续复用 `ExportProjectContext` 对同一事实源的 Markdown 表达；
- 规则、约束、文档入口定位继续基于 `entry_ref / entry_kind` 或其单向派生 view model；
- 四实体最小摘要不再被 Web / agent 各自转换成两套不同语义。

补充冻结：

- renderer 不是第二事实源；
- L3 adapter 不是第二事实源；
- 若某项字段只停留在 L3 / renderer 单向派生，就不得伪装成后端真实字段。

#### Scenario: Web 与 agent 同时消费 project-context
- **WHEN** Web 页面与 agent 导出都需要展示同一仓库上下文
- **THEN** 二者共享同一 project-context 事实源
- **AND** 不会出现两套互相漂移的规则解释或摘要结果

### Requirement: 落实三类页面的共享只读接入边界
系统 SHALL 按 `phase12-05 / 06` 已冻结的三类页面矩阵落地共享只读接入，不让页面临场决定锚点与承接方式。

当前阶段至少必须落实：

1. `repositories/$repositoryId` 可直接使用 `repository_id` 接入共享只读主线；
2. `products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 只能在满足既有 resolver 规则时接入共享只读主线；
3. `dashboard / onboarding / reviews/*` 只消费共享语义来源、受控派生摘要或入口定位视图，不升格为新的结构化锚点入口。

补充冻结：

- 间接页面若拿不到唯一 `repository_id`，不得伪造共享摘要闭环；
- 页面失败语义必须与 `phase12-05 / 06` 已冻结的局部降级策略一致；
- 不允许通过工作区扫描、手工查询或额外脚本补锚点。

#### Scenario: 间接页面尝试消费共享只读
- **WHEN** Product / Module / Decision 页面需要展示共享摘要
- **THEN** 必须先满足既有 resolver 规则
- **AND** 若无法唯一回到 `repository_id`，则按冻结设计降级，而不是硬猜

### Requirement: 落实只读消费的稳定性、可复用性与可定位性
系统 SHALL 让 `phase12-09` 的只读消费能力相较 `phase11` 具有更稳定、可复用、可定位的特征，并将这些特征落到实际消费链路中。

当前阶段至少必须体现：

- 稳定性：共享只读 query、adapter 与入口定位不再散落在多个页面私有实现中；
- 可复用性：四实体共享语义来源、最小摘要与入口定位可以被多切片稳定复用；
- 可定位性：规则、约束、文档入口仍保留 `entry_ref / entry_kind` 的定位能力，Web / agent 均能消费。

补充冻结：

- 不允许为了“复用方便”而复制既有 canonical facts；
- 不允许为了“定位方便”而退化成无定位的纯文案摘要；
- 不允许把只读消费深化写成 agent 专用能力或 Web 专用能力。

#### Scenario: 执行者完成 phase12-09 后复核只读消费能力
- **WHEN** 执行者复核 `phase12-09` 的交付结果
- **THEN** 能看到共享只读比 `phase11` 更稳定、可复用、可定位
- **AND** 能证明 Web / agent 已停止各自拼第二套解释性结果

## MODIFIED Requirements

### Requirement: phase12 内部顺序与 phase12-09 的承接表达
`phase12` 内部实现链路 SHALL 在 `phase12-09` 上继续对齐为同一口径：

- `phase12-05` 冻结共享入口与三类页面承接矩阵；
- `phase12-07` 冻结后端合同、导出结果与共享只读视图边界；
- `phase12-06` 冻结前端读路径 owner、共享摘要与 reread / 回流关系；
- `phase12-08` 收口前端四实体语义表达；
- `phase12-09` 负责把以上冻结结论真正落成共享只读入口与消费结果。

不得再出现：

- 一份文档要求 `phase12-09` 复用 `phase12-05 ~ 08` 结论，另一份文档又让它重开边界讨论；
- 一份文档要求 Web 共享只读入口唯一，另一份文档又允许多个切片各自落地 data owner；
- 一份文档要求 Web / agent 共用同一事实源，另一份文档又允许各自补造摘要结果。

#### Scenario: 读者对齐 phase12-09 的职责
- **WHEN** 读者同时查看 `phase12-05 ~ 08` 与 `phase12-09`
- **THEN** 能得到同一套“承接既有冻结结论并落地共享只读入口”的职责定义
- **AND** 后续实现不需要再发明第二套说明

## REMOVED Requirements

### Requirement: 允许 Web 与 agent 分别拼装 project-context 的解释性结果
**Reason**: 这会让 `phase12` 的共享只读深化目标失效，并重新长出第二套解释链、第二套入口定位和第二套摘要口径。
**Migration**: 后续实现必须统一复用 `ProjectContextService` 与 `frontend/src/features/project-context/` 的正式承接链路，让 Web / agent 从同一事实源出发消费共享结果。
