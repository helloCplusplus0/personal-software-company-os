# phase12-05 产出只读消费深化与共享入口设计 Spec

## Why
`phase12-03` 已经冻结了只读消费深化边界、共享只读 owner 分层与三类页面承接矩阵，但后续 `phase12-07 / 06` 仍缺少一份可直接进入实现的共享入口设计结果。`phase12-05` 的职责就是把 `phase11` 已交付只读项目上下文能力推进为“Web / agent 共享消费时到底怎么承接、哪些继续复用、哪些才允许最小新增”的正式设计输入。

## What Changes
- 冻结 `phase11` 项目上下文能力在 `phase12` 的深化方向与共享消费入口设计职责
- 冻结 `GetProjectContext / ExportProjectContext / frontend/src/features/project-context/` 与切片内展示 owner 的共享入口矩阵
- 冻结直接 `repository-scoped` / 间接 `repository-scoped` / 衍生消费页三类页面的承接规则与 resolver 责任边界
- 冻结“继续复用既有只读合同 vs 新增最小只读承接位”的判定规则
- 冻结供 `phase12-07` 继续承接的共享摘要字段需求、入口定位需求与最小 resolver 需求
- 对齐 `phase12` 三件套、`phase12-03` 与 `phase11-07 / 08` 的只读消费表达

## Impact
- Affected specs:
  - `phase12_semantic_alignment_and_readonly_consumption_foundation`
  - `phase12_03_freeze_readonly_consumption_boundary_entry_conditions`
  - `phase11_07_land_minimal_readonly_project_context_read_capability`
  - `phase11_08_land_agents_style_project_context_export`
- Affected code:
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/*`
  - `frontend/src/features/project-context/*`
  - `frontend/src/features/*/data/use-*-read.ts`
  - `frontend/src/features/dashboard/*`
  - `frontend/src/features/onboarding/*`
  - `frontend/src/features/review/*`
  - `frontend/src/features/product-registry/*`
  - `frontend/src/features/repository-binding/*`
  - `frontend/src/features/module-registry/*`
  - `frontend/src/features/decision-center/*`

## ADDED Requirements

### Requirement: 冻结 phase12-05 的只读消费深化设计职责
系统规格 SHALL 将 `phase12-05` 冻结为“产出只读消费深化与共享入口设计”的正式子任务，而不是直接实现、也不是重新讨论 `phase12-03` 已冻结边界。

当前子任务至少必须回答：

- `phase11` 已交付只读项目上下文能力在 `phase12` 继续深化到哪里；
- Web / agent 共享消费时，各层 owner 分别承接什么；
- 三类页面如何进入共享只读主线；
- 哪些继续复用 `phase11` 既有合同，哪些才允许新增最小承接位；
- 哪些共享摘要字段、入口定位字段与 resolver 需求必须交给 `phase12-07` 承接。

补充冻结：

- `phase12-05` 只产出设计结论，不直接越权实现后端合同、前端读路径 owner 或页面改造；
- `phase12-05` 必须继续遵守 `phase12-03` 已冻结的只读消费边界与更重通道进入条件；
- `phase12-05` 的完成标准是“足以直接进入 `/spec` 或后续实现设计链条”，而不是“只写一个方向摘要”。

#### Scenario: 执行者开始承接 phase12-05
- **WHEN** 执行者开始编写 `phase12-05` 设计结果
- **THEN** 能明确本任务的职责是产出共享入口与只读消费深化设计
- **AND** 不会把实现、协议扩张或更重通道一并吸收进来

### Requirement: 冻结共享只读 owner 与共享入口矩阵
系统规格 SHALL 为 Web / agent 共享消费冻结单值的 owner 与共享入口矩阵，使后续执行者无需再猜“哪个入口是 canonical、哪个入口只是共享消费层、哪个入口只能做展示映射”。

当前阶段至少必须明确：

1. **结构化 canonical owner**
   - `ProjectContextService.GetProjectContext`
   - 继续作为唯一结构化事实源
2. **agent-facing Markdown owner**
   - `ProjectContextService.ExportProjectContext`
   - 继续作为 agent-facing Markdown 导出入口
3. **Web 跨切片共享只读 owner**
   - `frontend/src/features/project-context/`
   - 继续作为唯一允许的新 Web 跨切片共享只读入口
4. **切片内展示 owner**
   - 各 feature 的 `pages/`、`components/` 与 `data/`
   - 仅消费共享只读结果并映射为切片内展示

补充冻结：

- `frontend/src/features/project-context/` 只能承接受控只读 adapter、共享语义摘要、入口定位视图与 query options；
- `frontend/src/features/project-context/` 不得承接写路径、页面私有状态、并列 canonical 字段语义或第二套业务事实；
- 切片页与切片组件不得直接长出跨切片共享摘要真相源；达到稳定复用条件后，必须回收到唯一共享入口。

#### Scenario: 执行者判断共享入口落点
- **WHEN** 某段共享只读解释需要同时被 Web 和 agent 或多个切片消费
- **THEN** 能机械判断它应落在 canonical owner、Markdown owner、Web 跨切片共享只读 owner 或切片内展示 owner
- **AND** 不会在页面层或临时 adapter 中再长出并列事实源

### Requirement: 冻结三类页面的共享承接与 resolver 矩阵
系统规格 SHALL 为直接 `repository-scoped`、间接 `repository-scoped` 与衍生消费页三类页面冻结共享承接方式与 resolver 责任边界。

当前阶段至少必须明确：

1. **直接 repository-scoped 页面**
   - `repositories/$repositoryId`
   - 允许直接承接 `repository_id` 并消费共享只读主线
2. **间接 repository-scoped 页面**
   - `products/$productId`
   - `modules/$moduleId`
   - `decisions/$decisionId`
   - 必须先通过稳定 resolver 回到同一 `repository_id`，再复用共享摘要
3. **衍生消费页**
   - `dashboard`
   - `onboarding`
   - `reviews/daily`
   - `reviews/weekly`
   - 只允许消费受控派生共享摘要、入口定位或解释性结果，不得升格为新的结构化主锚点入口

补充冻结：

- `repository_id` 仍是唯一正式结构化输入锚点；
- `product_id / module_id / decision_id` 只允许从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析；
- `phase12-05` 必须显式写出三类页面各自的输入、解析、复用与失败语义责任边界；
- 不允许把“页面内临场搜索、手工查询或工作区扫描补锚点”写成合法承接方案。

#### Scenario: 执行者为某个页面设计共享只读接入
- **WHEN** 执行者准备给某页面接入共享只读摘要
- **THEN** 能先判断该页面属于三类页面中的哪一类
- **AND** 能据此机械判断是否允许直接使用 `repository_id`、必须先解析回 `repository_id`，或只能消费派生摘要

### Requirement: 冻结复用既有合同与最小新增承接位的判定规则
系统规格 SHALL 冻结“复用 `phase11` 既有只读合同”与“新增最小只读承接位”的判定规则，避免局部便利被扩大成第二服务或第二共享入口。

当前阶段至少必须明确：

- 若 `GetProjectContext` 已可提供所需共享事实，应优先复用，而不是新增第二服务；
- 若确需新增字段或视图，唯一合法身份是 `ProjectContextService` 下的受控派生读取或其单向派生的共享视图；
- 若某项需求只涉及页面渲染或切片内解释，应继续留在切片展示层，不得越权升格为共享入口字段；
- 任何新增承接位都必须说明：输入锚点是什么、回收了哪段重复解释逻辑、失败条件是什么；
- 不允许把“前端先拼一版共享摘要，之后再治理”视为当前阶段的合法过渡。

#### Scenario: 执行者评估是否需要新增共享承接位
- **WHEN** 某个共享消费需求无法直接由既有 `GetProjectContext` 结构化结果满足
- **THEN** 执行者必须先判断是否可通过最小受控派生读取或单向派生共享视图承接
- **AND** 不会直接新增第二套共享 API、影子摘要合同或页面本地真相源

### Requirement: 冻结供 phase12-07 承接的共享摘要字段与入口定位需求
系统规格 SHALL 要求 `phase12-05` 显式输出供 `phase12-07` 继续承接的共享摘要字段需求、入口定位需求与最小 resolver 需求，而不是把这些内容留给执行者临场猜测。

当前阶段至少必须显式列出：

- 四实体最小共享摘要需要哪些结构化字段；
- 规则、约束、文档入口需要哪些定位字段与定位类型；
- 三类页面共享接入所需的最小 resolver 输入、输出与失败语义；
- 哪些字段应继续复用现有结构化结果，哪些字段才允许进入 `phase12-07` 的受控补充。

补充冻结：

- 所有新增字段需求都必须单向指向 canonical owner，不得长出第二套字段语义；
- 入口定位需求必须延续 `entry_ref / entry_kind` 或其同等定位关系，不得退回无定位的纯文案摘要；
- `phase12-05` 必须显式区分三类结果：`当前真实已覆盖`、`可由 L3 单向派生`、`仍需进入 phase12-07 候选评估`；
- 若当前源码合同尚未提供某字段或 resolver，`phase12-05` 不得越权写成“已覆盖”或“无需 phase12-07 新增”；
- 若某个字段只服务单页临时展示，不得误提炼成 `phase12-07` 的共享字段需求。

#### Scenario: 执行者为 phase12-07 准备输入
- **WHEN** `phase12-05` 设计结果交接给 `phase12-07`
- **THEN** 后续执行者能直接拿到共享摘要字段、入口定位字段与 resolver 最小需求
- **AND** 不需要重新猜“哪些字段该补、哪些入口该定位、解析链该落在哪里”

### Requirement: 冻结间接 repository-scoped 页面的正式 resolver 规则
系统规格 SHALL 要求 `phase12-05` 将 `products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 三类间接 `repository-scoped` 页面各自的 resolver 规则写成机械结论，而不是只写一句“通过 detail read 解析”。

当前阶段至少必须明确：

- Product 页面从哪里提取 `repository_id` 候选；
- Module 页面从哪里提取 `repository_id` 候选；
- Decision 页面如何经 `source_module_id / linked_modules` 链式回到 `repository_id` 候选；
- “唯一候选成立” 的正式条件是什么；
- 无候选、多候选、链式解析失败时，页面如何降级，以及何时进入 `phase12-07` 候选评估。

补充冻结：

- `phase12-05` 不得允许任何间接页面在多候选场景下静默选择一个 `repository_id`；
- `phase12-05` 也不得在缺少源码证据时直接否决 `phase12-07` 的受控派生 resolver 可能性；
- Decision 页面的 resolver 规则必须显式独立写出，不得与 Product / Module 合并成一条笼统结论。

#### Scenario: 执行者处理间接 repository-scoped 页面
- **WHEN** 执行者需要让 Product、Module 或 Decision 页面进入共享只读主线
- **THEN** 能机械判断该页面的 resolver 输入、链路、唯一候选条件与失败语义
- **AND** 不会在页面内临场猜测 `repository_id`

## MODIFIED Requirements

### Requirement: phase12 三件套中的只读消费深化表达
`phase12` 三件套中的只读消费深化表达 SHALL 在 `phase12-05` 上对齐为同一口径：

- `phase12-05` 负责产出共享入口与只读消费深化设计；
- `phase12-07` 承接后端合同、View 与最小 resolver 需求；
- `phase12-06` 承接前端读路径 owner、共享摘要接入与回流设计；
- 三者继续遵守 `05 -> 07 -> 06` 的内部顺序，不得倒置。

不得再出现：

- 一个文档要求 `phase12-05` 只做方向描述，另一个文档又要求其直接给出可进入实现的共享入口结论；
- 一个文档允许页面临场决定 resolver 方式，另一个文档又要求三类页面承接矩阵固定；
- 一个文档把 `frontend/src/features/project-context/` 写成可选承接位，另一个文档又要求其为唯一允许的新 Web 跨切片共享只读 owner；
- 一个文档在缺少源码证据时就替 `phase12-07` 判定“无需新增字段 / resolver”，另一个文档又要求 `phase12-07` 负责该判断。

#### Scenario: 读者对齐 phase12-05 的职责
- **WHEN** 读者分别查看 `dev_plan`、`architecture_plan`、`shared_baseline` 与 `phase12-03`
- **THEN** 能获得同一套 `phase12-05` 职责、顺序、共享入口矩阵与交接边界
- **AND** 后续 `/spec` 与实现不需要补造第二套说明

## REMOVED Requirements

### Requirement: 允许执行者临场决定共享入口、resolver 链或锚点回收方式
**Reason**: 这会让 `phase12-05` 失去“为 `phase12-07 / 06` 提供可机械承接设计输入”的职责，并重新打开只读消费深化的关键边界。
**Migration**: 后续设计与实现必须直接消费 `phase12-05` 冻结的 owner 矩阵、三类页面承接矩阵、复用判定规则与 `phase12-07` 输入需求，不再允许页面或执行者各自补一套解释。
