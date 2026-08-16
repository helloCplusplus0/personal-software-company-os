# phase12-07 产出后端合同、导出结果与共享只读视图设计 Spec

## Why
`phase12-05` 已经冻结了共享只读 owner、三类页面承接矩阵与 `phase12-07` 候选清单，但还没有把“哪些继续复用 `GetProjectContext`、哪些才允许进入受控派生读取”正式单值化。`phase12-07` 的职责就是把后端合同、导出结果与共享只读视图的最小设计冻结到足以直接进入 `/spec` 与代码实现。

## What Changes
- 冻结 `phase12-07` 在 `phase12` 内的正式职责、范围与非目标
- 冻结 `.proto / Connect / service / renderer` 的最小演进判定规则
- 冻结“继续复用 `GetProjectContext` vs 进入 `ProjectContextService` 下的受控派生读取”的正式判断矩阵
- 冻结 `phase12-05` 提交的字段 / 入口定位 / resolver 候选中，哪些继续复用、哪些允许进入受控补充、哪些明确不做
- 冻结任何新增只读承接位必须回收的前端 / agent 重复解释逻辑
- 对齐 `phase12` 三件套、`phase12-03` 与 `phase12-05` 的后端合同口径

## Impact
- Affected specs:
  - `phase12_semantic_alignment_and_readonly_consumption_foundation`
  - `phase12_03_freeze_readonly_consumption_boundary_entry_conditions`
  - `phase12_05_design_readonly_consumption_shared_entry`
  - `phase11_07_land_minimal_readonly_project_context_read_capability`
  - `phase11_08_land_agents_style_project_context_export`
- Affected code:
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/*`
  - `backend/internal/gen/proto/psco/project_context/v1/*`
  - `backend/internal/gen/connect/psco/project_context/v1/*`
  - `frontend/src/features/project-context/*`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结 phase12-07 的后端合同设计职责
系统规格 SHALL 将 `phase12-07` 冻结为“产出后端合同、导出结果与共享只读视图设计”的正式子任务，而不是直接编码，也不是重新打开 `phase12-05` 已冻结的页面 owner 与 resolver 分类。

当前子任务至少必须回答：

- 哪些共享消费需求继续复用 `GetProjectContext`；
- 哪些候选项允许进入 `ProjectContextService` 下的受控派生读取；
- 哪些候选项当前明确不做；
- 若新增字段、视图或导出结果，它们如何保持 `repository_id` 与 `entry_ref / entry_kind` 的定位能力；
- 它们具体回收了哪段前端 / agent 重复解释逻辑。

补充冻结：

- `phase12-07` 只产出设计结论，不直接修改 `.proto`、Connect、service 或 renderer；
- `phase12-07` 必须继续遵守 `.proto` 是唯一长期合同源、`ConnectRPC` 是业务接口正式传输层、`ProjectContextService` 是唯一业务域载体的上游规则；
- `phase12-07` 的完成标准是“足以直接进入 `/spec` 与实现”，而不是“先给几个备选方向”。

#### Scenario: 执行者开始承接 phase12-07
- **WHEN** 执行者开始编写 `phase12-07` 设计结果
- **THEN** 能明确本任务的职责是冻结后端合同、导出结果与共享只读视图的最小设计
- **AND** 不会把前端读路径 owner、页面渲染或更重通道一并吸收进来

### Requirement: 冻结复用 GetProjectContext vs 受控派生读取的判断矩阵
系统规格 SHALL 为 `phase12-05` 提交的字段、入口定位与 resolver 候选冻结单值判断矩阵，使后续执行者不需要再猜“这里是不是应该新建第二服务”。

当前阶段至少必须明确：

1. 哪些需求继续由 `GetProjectContext` 真实字段承接；
2. 哪些需求只允许作为 `frontend/src/features/project-context/` 的单向派生 adapter 或 view model 存在；
3. 哪些需求允许进入 `ProjectContextService` 下的受控派生读取；
4. 哪些候选项当前明确不进入本阶段实现。

补充冻结：

- 若现有 `GetProjectContext` 已足以承接，则不得为了“前端更顺手”而新增第二服务；
- 若确需新增只读视图，其唯一合法身份是 `ProjectContextService` 下的受控派生读取，不得长成并列业务域；
- 不允许同时存在“L1 已覆盖”和“必须新增受控派生读取”两种互相冲突的结论；
- 所有判断都必须绑定到具体候选项，而不是停留在抽象原则。

#### Scenario: 执行者判断某个共享需求是否需要新增后端承接位
- **WHEN** 执行者面对某个字段、入口定位或 resolver 候选
- **THEN** 能机械判断它是继续复用 `GetProjectContext`、停留在 L3 单向派生，还是进入受控派生读取
- **AND** 不会再提出第二服务或第二事实源

### Requirement: 冻结受控派生读取的最小合同约束
系统规格 SHALL 要求 `phase12-07` 对任何允许进入受控派生读取的候选，显式写出最小合同约束。

当前阶段至少必须明确：

- 该承接位是否继续复用 `GetProjectContext`；
- 若不复用，为什么仍属于 `ProjectContextService` 下的受控派生读取；
- 它的最小输入锚点是什么；
- 它的最小输出字段或视图 shape 是什么；
- 它回收了哪段前端 / agent 重复解释逻辑；
- 它如何继续保留 `repository_id` 与 `entry_ref / entry_kind` 或其单向派生定位关系。

补充冻结：

- 若无法说明回收了哪段重复解释逻辑，则不得进入受控派生读取；
- 若会复制既有 canonical facts，只是换一个 shape 暴露，也不得进入受控派生读取；
- 若某候选只服务单页临时便利，不得误升级为后端合同变更。

#### Scenario: 执行者评估一个候选是否具备新增资格
- **WHEN** 某个候选无法直接由现有合同与 L3 单向派生满足
- **THEN** 执行者必须逐项写清输入锚点、输出字段、回收逻辑与定位保留方式
- **AND** 只有满足这些条件时，才允许把它列入受控派生读取设计

### Requirement: 冻结导出结果与共享只读视图的承接边界
系统规格 SHALL 冻结 `ExportProjectContext` 与可能新增的共享只读视图之间的承接边界，避免 Markdown 导出和 Web 共享结果互相漂移。

当前阶段至少必须明确：

- `ExportProjectContext` 继续作为 agent-facing Markdown 导出 owner；
- Web 共享只读结果优先来自 `GetProjectContext` 真实字段或其受控派生读取；
- renderer 只负责把同一事实源渲染成导出结果，不得成为第二套字段语义来源；
- 若新增只读视图，它与 Markdown 导出之间是“共享事实源，不共享真相定义”的关系。

补充冻结：

- 不允许为了给 Web 让路，把 Markdown renderer 变成并列 canonical owner；
- 不允许为了给 agent 让路，再造一套只服务导出的结构化合同；
- 不允许让 Web 与 agent 各自长出不同版本的规则 / 约束 / 文档入口解释。

#### Scenario: 执行者设计导出结果或共享只读视图
- **WHEN** 执行者需要设计 Web 共享结果与 agent 导出结果的关系
- **THEN** 能明确二者共享同一事实源
- **AND** 不会把 renderer 或导出结果写成第二套合同源

### Requirement: 冻结 phase12-07 的最小设计产物模板
系统规格 SHALL 要求 `phase12-07` 的设计结果继续满足 `phase12-04 ~ 07` 的统一最小模板，并补充适用于后端合同设计的最小审计面。

当前子任务至少必须显式产出：

1. 影响对象清单  
2. 结论矩阵  
3. 承接位矩阵  
4. 共享语义来源 vs 切片内渲染 / 后端合同边界矩阵  
5. Before / After 样例  
6. 明确不做清单  

同时至少补充以下后端设计审计面：

- 真实已有合同字段清单；
- 允许进入受控派生读取的候选清单；
- 明确不进入本阶段的候选清单；
- 每个候选对应的“重复解释回收收益”说明。

#### Scenario: 执行者准备把 phase12-07 交给后续 /spec
- **WHEN** `phase12-07` 设计结果完成
- **THEN** 后续执行者能直接拿到单值的后端合同判断、候选去留、最小字段与定位约束
- **AND** 不需要再补造第二套说明

## MODIFIED Requirements

### Requirement: phase12 三件套中的 phase12-07 顺序与职责表达
`phase12` 三件套中的 `phase12-07` 表达 SHALL 在本 spec 中继续对齐为同一口径：

- `phase12-05` 先冻结共享 owner、页面承接分类与候选清单；
- `phase12-07` 再判断继续复用 `GetProjectContext` 还是进入受控派生读取；
- `phase12-06` 最后消费 `phase12-05 / 07` 已冻结的 owner、字段与入口结果。

不得再出现：

- 一个文档要求 `phase12-07` 判断是否新增受控派生读取，另一个文档又提前把结果写死；
- 一个文档允许新增第二服务，另一个文档又要求所有新增都留在 `ProjectContextService` 下；
- 一个文档把 renderer 当作事实源，另一个文档又要求 `.proto` 仍是唯一长期合同源。

#### Scenario: 读者对齐 phase12-07 的职责
- **WHEN** 读者分别查看 `dev_plan`、`architecture_plan`、`shared_baseline`、`phase12-05`
- **THEN** 能获得同一套 `phase12-07` 职责、顺序、合同判断规则与交接边界
- **AND** 后续 `/spec` 与实现不需要补造第二套解释

## REMOVED Requirements

### Requirement: 允许执行者凭实现便利新建第二服务或影子合同
**Reason**: 这会直接破坏 `phase12` 的单一 canonical facts 约束，并使 Web / agent 重新长出并列事实源。
**Migration**: 后续设计与实现必须以 `GetProjectContext`、`ExportProjectContext` 与 `ProjectContextService` 下的受控派生读取为唯一合法承接位，不再允许新建第二服务或影子合同。
