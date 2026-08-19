# phase15-02 冻结事件模型语义边界与 brief 演进对照 Spec

## Why

`phase15-01` 已冻结单一主交付边界，但事件模型的语义细节（三轨 × event_kind 合法矩阵、`task_key` 格式、派生规则空态语义）与 brief 演进对照（`progress = 9`、`BriefProgress` 字段清单）仍散落于 `shared_baseline` §3，未被确认为可直接进入实现设计的单值语义基线；同时三重边界（plan.md / phase11 `PhaseEntry` / Decision）只有结论级声明、缺少可机械判定的显式断言。`phase15-03 / 04 / 05` 三个设计类子任务全部以本 spec 为语义上游，必须先完成语义单值冻结，防止设计期分叉。

`phase15-02` 是边界收敛类子任务：纯文档冻结，不改任何代码、不改 proto 文件、不回写根级文档（根级同步属 `phase15-09`）。

## What Changes

- 前置确认并冻结三轨 × event_kind 合法矩阵（3×4 全 12 格判定）与 `task_key` 格式规则（4 组必填正则 + note 可空规则）
- 前置确认 shared_baseline §3.3 的 9 条校验规则为清单级语义基线（形式化〔判定逻辑 + 稳定错误码〕归 phase15-03，本 spec 不偷渡）
- 冻结派生规则语义（4 项派生 + 三键链全序 + 派生执行位 + `current_phase_key` 空值双情形同型语义）
- 留档 brief 演进前后对照表（前侧 = `project_context.proto` 现状；后侧 = `progress = 9` + `BriefProgress` 字段清单 + 空态恒构造 + 槽位 2/3/4 保持 reserved）
- 冻结三重边界分离的显式断言语句（plan.md / phase11 `PhaseEntry` / Decision 每重含可机械判定断言）
- 产出本目录三件套（spec / tasks / checklist）

## Impact

- Affected specs:
  - `phase15-03 / 04 / 05`（本 spec 为其唯一直接语义上游：合法矩阵与校验规则 → phase15-03 形式化；brief 对照表与字段清单 → phase15-04 合同设计；派生空态语义与三重边界 → phase15-05 前端承接）
  - `phase15-01` spec（本 spec 承接其边界声明，不冲突不扩界）
  - `docs/phase/phase15_*` 三件套（本 spec 从 shared_baseline §3.2-3.5 收敛确认，不反向改写；不一致时以 shared_baseline 单值基线为准并回改本 spec）
- Affected code: 无（proto / Go / TS 文件零改动；`project_context.proto` 仅供前侧对照读取）
- 验收产物：本目录 `tasks.md / checklist.md`

## ADDED Requirements

### Requirement: 三轨 × event_kind 合法矩阵与 task_key 格式规则必须单值冻结

本 spec SHALL 前置确认并冻结以下合法矩阵（与 `shared_baseline` §3.3 规则 1-8、`architecture_plan` §4.5 矩阵单值一致），作为 phase15-03 校验形式化与 phase15-05 表单联动禁用的唯一语义依据：

| workflow_type | phase_started | phase_completed | task_completed | note |
|---|---|---|---|---|
| `phase` | 合法（task_key 必填，格式 `phaseNN`） | 合法（task_key 必填，格式 `phaseNN`） | 合法（task_key 必填，格式 `phaseNN-MM`） | 合法（task_key 可空） |
| `audit` | 禁止 | 禁止 | 合法（task_key 必填，格式 `audit_NNN`） | 合法（task_key 可空） |
| `fix` | 禁止 | 禁止 | 合法（task_key 必填，格式 `fix_NNN`） | 合法（task_key 可空） |

`task_key` 格式规则（正则冻结，phase15-03 直接转译为实现）：

| 规则 | 适用 | 格式 |
|---|---|---|
| K-1 | phase 轨 `phase_started` / `phase_completed` | `^phase[0-9]{2,}$` |
| K-2 | phase 轨 `task_completed` | `^phase[0-9]{2,}-[0-9]{2,}$` |
| K-3 | audit 轨 `task_completed` | `^audit_[0-9]{3,}$` |
| K-4 | fix 轨 `task_completed` | `^fix_[0-9]{3,}$` |
| K-5 | 三轨通用 `note` | 可空；若填写不强制格式（自由标注） |

显式不做的约束（沿 `shared_baseline` §3.3 注，防止实现期加戏）：

- 不做 `task_key` 唯一性约束（误录经 Delete 修正；补录与重录为合法场景）
- 不对 `occurred_at` 做未来时间校验（补录历史是显式需求，排序按声明时间单值执行）

#### Scenario: 矩阵与格式规则可机械转译

- **WHEN** phase15-03 形式化校验规则或 phase15-05 设计表单联动禁用
- **THEN** 12 格矩阵判定与 K-1 ~ K-5 正则可直接转译为判定逻辑与禁用矩阵，无需二次解释
- **AND** 任何子任务不得新增矩阵未声明的合法组合（如 audit 轨 phase_started）或收窄已声明组合（如 note 轨填写 task_key）

### Requirement: 9 条校验规则完成清单级前置确认

本 spec SHALL 前置确认 `shared_baseline` §3.3 的 9 条校验规则为清单级语义基线（逐条引用，不改写）；每条的判定逻辑细化与稳定错误码归 `phase15-03` 形式化，本 spec 不提前产出：

1. `workflow_type ∈ {phase, audit, fix}`；`event_kind ∈ {phase_started, phase_completed, task_completed, note}`
2. `event_kind = task_completed` 时 `task_key` 必填
3. `workflow_type = phase` 且 `event_kind ∈ {phase_started, phase_completed}`：`task_key` 必填且格式 K-1
4. `workflow_type = phase` 且 `event_kind = task_completed`：`task_key` 必填且格式 K-2
5. `workflow_type = audit` 且 `event_kind = task_completed`：`task_key` 必填且格式 K-3
6. `workflow_type = fix` 且 `event_kind = task_completed`：`task_key` 必填且格式 K-4
7. `workflow_type ∈ {audit, fix}`：`event_kind` 仅允许 `task_completed / note`
8. `event_kind = note`：`task_key` 可空；若填写不强制格式（K-5）
9. `evidence_ref` 若填写必须为 `/` 开头仓库内相对路径或 `https://` 开头 URL（与 Standard `ref` 规则单值一致）；`title` 非空上限 200；`detail` 上限 2000；`source` 现值仅接受 `manual`

#### Scenario: 清单与形式化边界清晰

- **WHEN** phase15-03 产出校验形式化设计
- **THEN** 其规则集合与本清单 1:1 对应（条数、语义、顺序均不增减），仅补充判定逻辑与错误码
- **AND** 本 spec 不含任何错误码定义或错误语义细分（`invalid_argument` 等对外语义归 phase15-04 合同设计）

### Requirement: 派生规则语义必须冻结含空值边界

本 spec SHALL 冻结以下派生规则语义（与 `shared_baseline` §3.4 单值一致，并补齐空值边界语义）：

| 派生项 | 规则 |
|---|---|
| 当前 phase key | 最新 `phase_started`（三键链序）的 `task_key`；若存在同 key 且序更晚的 `phase_completed` → 全部完结态（空值） |
| 当前 phase label | 该最新 `phase_started` 的 `title` |
| 最新完成任务项 | 该 repository 最新一条 `task_completed` 事件（不限 phase 轨；三轨同序取最新） |
| recent_events | 最近 N=10（冻结）条三轨混合事件（同三键链序） |

- **三键链全序**：`(occurred_at DESC, created_at DESC, id DESC)`；`id DESC` 为最终 tiebreak，在补录同刻事件与同事务批量插入（未来 `source=git`）两类碰撞场景下保证时间轴顺序与派生结果确定。
- **派生执行位**：后端 service 层统一计算；web 当前卡与 brief 摘要共用同一实现；不落库、不缓存第二套状态（裁决③）。
- **`current_phase_key` 空值双情形同型语义**（本 spec 补齐冻结）：
  - 从未开始：无任何 `phase_started` 事件 → `current_phase_key = ""`、`current_phase_label = ""`
  - 全部完结：最新 `phase_started` 存在，但存在同 key 且序更晚的 `phase_completed` → 同样 `current_phase_key = ""`、`current_phase_label = ""`
  - 两种空值情形在派生输出层面**同型**（均为空字符串零值），不引入第二套 phase 状态字段或状态枚举（与已退役 `BriefPhaseStatus` 零关系，CON-08 禁止复活口径）
  - web 当前卡如需区分两种空值情形的展示文案（如"尚未开始 / 已完结"），属 phase15-05 前端承接语义；其数据通道必须满足"后端统一派生"约束（DP-1，phase15-05 裁决），不得形成前端第二套派生语义

#### Scenario: 派生语义无第二套解释

- **WHEN** phase15-03 设计派生算法或 phase15-04 设计 ProgressReader 输出
- **THEN** 派生项、三键链序、空值语义与本冻结单值一致；"最新"一律指三键链序最新，不存在按 created_at 单键或按录入顺序的第二套"最新"定义
- **AND** 前端不得自行从事件集合计算派生值用于当前卡（web 当前卡数据通道归 DP-1 在 phase15-05 裁决，且裁决空间已被"后端统一派生"约束封死）

### Requirement: brief 演进前后对照表必须留档

本 spec SHALL 留档 `GetProjectBriefResponse` 演进前后对照表。前侧为 `proto/psco/project_context/v1/project_context.proto` 现状（L191-204），后侧为 phase15 演进（裁决⑩）：

**前侧（现状）**：

| 字段号 | 字段 | 类型 | 状态 |
|---|---|---|---|
| 1 | `repository` | `RepositorySummary` | 在用 |
| 2 | `governance_profile` | — | reserved（T7 裁决画像退役） |
| 3 | `global_assets` | — | reserved |
| 4 | `current_phase` | — | reserved |
| 5 | `products` | `ProductSummary[]` | 在用 |
| 6 | `modules` | `ModuleSummary[]` | 在用 |
| 7 | `decisions` | `DecisionSummary[]` | 在用 |
| 8 | `standards` | `psco.standard.v1.Standard[]` | 在用（phase14-02 冻结演进） |

**后侧（phase15 演进）**：

| 字段号 | 字段 | 类型 | 演进 |
|---|---|---|---|
| 1 ~ 8 | （全部不变） | — | 零改动 |
| 9 | `progress` | `BriefProgress`（内联轻量消息，定义于 `project_context.proto`） | 新增 |
| 2 / 3 / 4 | — | — | 保持 reserved（不复用画像退役槽位；`progress = 9` 为下一可用号） |

**`BriefProgress` 字段清单**（语义冻结；字段号分配与消息定义落点细节归 phase15-04 合同设计）：

| 字段 | 类型 | 语义 |
|---|---|---|
| `current_phase_key` | `string` | 当前 phase 的 `task_key`（`phaseNN`）；空值含从未开始 / 全部完结两种情形（同型零值） |
| `current_phase_label` | `string` | 当前 phase 标题（该最新 `phase_started` 的 `title`） |
| `latest_task_completed` | `psco.progress.v1.ProgressEvent`（可选） | 最新完成任务项事件（三轨同序取最新；无任务完成时为零值不设置） |
| `recent_events` | `psco.progress.v1.ProgressEvent[]` | 最近 N=10 条三轨混合事件（三键链倒序） |

**装配约束冻结**：

- 空态语义：Go 装配侧恒构造 `progress` 块（非 nil）；0 事件时字段为零值、`recent_events` 为空数组——前端进度区与 agent 单值消费，无双套判空逻辑
- 元素同型复用：`latest_task_completed` 与 `recent_events[]` 元素同型复用 `psco.progress.v1.ProgressEvent`，不建第二套摘要消息
- 跨包导入：`project_context.proto` 需新增 `import "psco/progress/v1/progress.proto"`（沿既有 `import "psco/standard/v1/standard.proto"` 模式；import 细节与生成链影响归 phase15-04）
- 装配来源：`projectcontext` 经 `ProgressReader` 跨模块 Read 接口（沿袭 `StandardReader` 模式）；brief 与 `ListProgressEvents` RPC 同源同派生
- 顶层块口径演进：5 顶层块 → 5 顶层块 + `progress` 摘要块

口径辨析已在 `phase15-01` spec 留档（`phase14-11`"不回填"指向 reserved 槽位 2/3/4 恢复；phase15 以新字段号 9 新建 `BriefProgress` 属裁决⑩正当演进），本 spec 引用不重复展开。

#### Scenario: 对照表是 phase15-04 合同设计的唯一前侧依据

- **WHEN** phase15-04 设计 brief 合同演进
- **THEN** 其字段号、字段清单、空态语义、导入关系与本对照表单值一致；字段号分配（BriefProgress 内部 1-4）为本 spec 未冻结项，由 phase15-04 按自然序冻结
- **AND** 不出现 reserved 槽位复用、第二套摘要消息或 brief 之外的第三消费通道定义

### Requirement: 三重边界分离必须各有显式断言语句

本 spec SHALL 冻结三重边界分离的可机械判定断言（裁决⑪展开；每重边界含定位差异 + 断言语句）：

**边界一：与 plan.md**

- 定位差异：`plan.md` 是阶段路线唯一真相源（`project_rules.md` §1），正文事实源在仓库；`progress_events` 是 PSCO-native 推进事件事实
- 显式断言：
  1. `progress_events` 不复制、不解析、不缓存 `plan.md` 正文任何内容
  2. `plan.md` 修订不触发进度事件任何自动变更；进度事件录入不回写 `plan.md`
  3. 唯一允许的关联形态是 `evidence_ref` 导航引用（`/` 开头路径指向 plan.md 属合法引用，仅导航零托管）

**边界二：与 phase11 `PhaseEntry`**

- 定位差异：`GetProjectContextResponse.phases` 是根级文档只读投影（导航入口，自文档扫描派生）；`progress_events` 是用户手动录入的 PSCO-native 持久化事件事实（事件流）
- 显式断言：
  1. phase15 不改动 `GetProjectContextResponse` 的 `PhaseEntry` / 规则投影（非目标 10；两个 deprecated 兼容 RPC 行为同样不变）
  2. `PhaseEntry.phase` 与进度事件 `task_key`（`phaseNN`）语义相关但数据无关：不建立自动映射、不同步、不互为校验依据
  3. 两者并存各司其职：`PhaseEntry` 承接"根级文档说什么"，`progress_events` 承接"用户记录发生了什么"；内容冲突时各自独立呈现，不裁决不合并

**边界三：与 Decision**

- 定位差异：进度事件 = 客观时间轴事实（"发生了什么"）；Decision = 决策留痕（"为什么这么决定"）
- 显式断言：
  1. 不合并、不互替：进度事件不升格为决策留痕，Decision 不降格为时间轴条目
  2. 不互链（起步）：进度事件无 `decision_id` 字段，Decision 无 `progress_event_id` 字段；`evidence_ref` 可指向 Decision detail 路径属合法导航引用，但不构成结构化互链
  3. 数据零交叉：派生规则不消费 Decision 数据；brief `decisions[]` 装配不消费 `progress_events`

#### Scenario: 边界断言可进验收

- **WHEN** phase15-08 执行十一项裁决门禁验证（裁决⑪"三重边界不被破坏"）
- **THEN** 本节 9 条断言（每重 3 条）逐条可经代码检视 + 运行时行为取证验证
- **AND** 任一断言被实现破坏（如派生规则读取 Decision 数据）即验收 FAIL，无裁量空间

### Requirement: 本 spec 与 phase15-01 边界上游及三件套单值一致

本 spec 的全部冻结项 SHALL 落入 `phase15-01` 冻结的四组成部分边界，且与 `shared_baseline` §3.2-3.5 / §2.2 单值一致：

- 合法矩阵 + task_key 格式 + 9 条校验规则清单 → 组成部分 1（合同与存储）的语义前提（直接承接裁决④⑤）
- 派生规则语义 → 组成部分 4（派生逻辑）的语义前提（直接承接裁决③）
- brief 前后对照表 → 组成部分 3（agent 消费链路）的语义前提（直接承接裁决⑩）
- 三重边界断言 → 边界声明的裁决⑪展开
- 本 spec 不承载 phase15-03 的形式化（错误码 / 判定逻辑细化）、phase15-04 的合同细节（DDL / RPC envelope / 字段号分配 / ProgressReader 接口签名）、phase15-05 的前端细节（组件树 / 交互规格 / DP-1 裁决），防止范围前置与双源漂移

#### Scenario: 一致性可校验

- **WHEN** 独立复核执行
- **THEN** 本 spec 各冻结项与 `shared_baseline` §3.2-3.5 / §2.2、`architecture_plan` §4.4-4.7、`phase15-01` spec 对应内容比对一致；前侧对照表与 `project_context.proto` 实际内容逐字段一致
- **AND** git 工作区中本 spec 仅为目录新增，零 proto / 代码 / 根级文档改动
