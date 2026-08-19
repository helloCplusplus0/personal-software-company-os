# phase15-01 冻结 Project Progress Timeline Foundation 的范围边界、成功标准与非目标 Spec

## Why

`phase15 /plan` 三件套已建立（十一项裁决冻结 + 9 个子任务规划），但裁决结论、成功标准与非目标分散于三件套正文，且与 `phase14-11` 进入条件、根级文档状态存在多处交叉引用。进入子任务执行前，必须先把"做什么 / 做到什么程度 / 明确不做什么"收敛为单一主交付边界声明，作为 `phase15-02` 及后续全部子任务的唯一直接边界上游，防止实现期范围漂移与裁决失真。

`phase15-01` 是边界收敛类子任务：纯文档冻结，不改任何代码，不回写根级文档（根级同步属 `phase15-09`）。

## What Changes

- 冻结 `phase15` 单一主交付边界声明（`Project Progress Timeline Foundation` 单值定义 + 四组成部分 + 治理层信息流定位）
- 完整留档十一项裁决矩阵（8 主裁决 + 3 补裁，每项含结论、对应输入、`phase15-08` 验收验证点）
- 冻结成功标准（4 条可验收条目，对齐 `architecture_plan` §3）
- 冻结非目标（12 项显式排除）
- 留档候选池顺延项（`phase14-11` 候选池中除时间轴外的全部顺延项及其进入条件约束），排序裁决留待 `phase15-09` 冻结
- 留档 `CON-08` 时间轴承接口径继承声明（T7 裁决后口径在本阶段的落地形态逐项对齐）
- 登记 3 个实现级待决点（DP-1 / DP-2 / DP-3）为后续子任务 spec 的必备裁决项（边界收敛复核中识别，属"决策归属声明"，不替后续子任务预先裁决）
- 产出本目录三件套（spec / tasks / checklist）

## Impact

- Affected specs:
  - `phase15-02` 及后续全部子任务（本 spec 为其唯一直接边界上游；`phase15-02` 的语义边界冻结以本 spec 的边界声明为起点）
  - `phase14-11` spec（其冻结的 `phase15` 进入条件由本 spec 承接落地；正文不改）
  - `docs/phase/phase15_*` 三件套（本 spec 从其收敛，不反向改写；不一致时以 `shared_baseline` 单值基线为准并回改本 spec）
- Affected docs: 无根级文档回写（`AGENTS.md / plan.md` 等根级同步属 `phase15-09`）
- Affected code: 无
- 验收产物：本目录 `tasks.md / checklist.md`

## ADDED Requirements

### Requirement: phase15 单一主交付边界必须单值冻结

本 spec SHALL 冻结以下边界声明，且与 `shared_baseline` §3.1 / §2.3、`architecture_plan` §4.2 / §4.3 单值一致：

- **单值定义**：repository 锚定的三轨 append-only 推进事件流——以任务项级颗粒承接 `phase / audit / fix` 三套 workflow 的推进事实，web 可维护可回看、agent 可直读完整信息流，"当前"仅为派生视图。
- **单一主交付**：`Project Progress Timeline Foundation` 最小主线（合同 → 存储 → 后端 → 前端 → agent 消费），一能力贯通，不并列第二主交付；本阶段无退役任务（T7 裁决后时间轴当前为零承接位，纯新建）。
- **四组成部分**：
  1. `progress_events` 合同与存储（新建）
  2. web 端维护与展示入口（Repository detail 内嵌进度区，新建切片 `features/progress/`）
  3. agent 消费链路（brief `progress = 9` 装配 + `ListProgressEvents` 独立 RPC）
  4. 派生逻辑（当前 phase / 最新任务的读取侧计算，web 当前卡与 brief 摘要共用）
- **正式定位**：repository 锚定的治理层信息流，不是第六业务主实体，不进入四实体关系主链；单仓库作用域（与 Standard 的全局作用域形成对照）。
- **初心问题**：项目推进进度不再散落在提示词手写说明与对话记忆中；一次维护、web / agent 共享消费，与 `standards[]` 合成 agent 的"背景 + 进度"完整上下文。

#### Scenario: 边界声明可机械判定

- **WHEN** 任一后续子任务 spec 声明其范围
- **THEN** 其交付内容必须落入四组成部分之一或其直接支撑（校验 / 装配 / 测试 / 文档收口）
- **AND** 落不进边界声明的内容只能以非目标或顺延项身份出现，不得以"顺手实现"身份进入实现类子任务

### Requirement: 十一项裁决矩阵必须完整留档为阶段最高优先级基线

本 spec SHALL 完整留档十一项裁决（前 8 项为 2026-08-19 用户两轮结构化拍板的主裁决，后 3 项为随方案整体获批冻结的补裁），作为全部子任务的强制边界：

| # | 裁决项 | 结论（单值） | 对应输入 | phase15-08 验收验证点 |
|---|---|---|---|---|
| ① | 排序与能力定位 | 项目进度时间轴为唯一主交付；接管高频提示词"进度说明"段，与 `standards[]`（背景说明）合成一次调用的"背景 + 进度" | `phase14-11` 候选池 + 用户反思 | brief 一次调用含背景 + 进度 |
| ② | 承接锚点 | 源代码仓库：`progress_events.repository_id` 锚定，与 brief 装配粒度对齐 | 用户想法 1 | repository 锚定 |
| ③ | 数据模型形态 | append-only 推进事件流（单表）；记录信息流而非单点；"当前"为派生值不落库；历史永不丢弃 | 用户想法 5（强化警惕点） | append-only 断言 + 派生不落库 |
| ④ | 三轨 workflow | `phase / audit / fix` 对齐 `docs/` 三目录与 `project_rules.md` §4 三推进链 | 用户澄清第一点 | 三轨可录可滤 |
| ⑤ | 事件颗粒度 | 主事件 = 任务项级（`phaseNN-MM` / `audit_NNN` / `fix_NNN`，对齐"子任务验收通过即 git 提交"节奏）；phase 轨设边界标记；audit / fix 轨仅 `task_completed` 与 `note` | 用户澄清第二点 | 任务项颗粒 + phase 边界标记 + audit/fix 无边界事件 |
| ⑥ | 维护方式 | web 手动维护，Repository detail 内嵌维护与展示入口；无独立路由 / 导航项 / Dashboard 主卡片 | 用户想法 2 | 维护入口仅在 Repository detail |
| ⑦ | 证据引用边界 | `evidence_ref` 导航引用（与 Standard `ref` 规则单值一致）；不托管 plan.md / spec 正文 | 用户想法 3 + Standard 哲学 | evidence_ref 导航且正文零托管 |
| ⑧ | 来源预留 | `source = manual / git / agent`，起步仅实现 `manual`；git 采集与 agent 写回为后续进入条件 | 后续演进预留 | source 仅 manual 可写 |
| ⑨ | RPC 语义（补裁） | 最小集 3 RPC：List（不分页）/ Create / Delete；无 Update（append-only 语义纯净） | 方案获批 | 无 Update RPC |
| ⑩ | 消费分层（补裁） | brief `progress = 9` 摘要块 + `ListProgressEvents` 完整流（支持 `workflow_type` 过滤） | 方案获批 | brief 摘要与完整流分离 |
| ⑪ | 三重边界分离（补裁） | 与 plan.md / phase11 `PhaseEntry` / Decision 均不合并不互替 | 方案获批 | 三重边界不被破坏 |

裁决矩阵的展开细节（字段矩阵、9 条校验规则、派生规则、brief 演进矩阵、前端承接矩阵）以 `shared_baseline` §3 为单值来源；本矩阵只冻结裁决结论本身，不复制展开细节（防双源漂移）。

#### Scenario: 任一裁决可追溯且不被后续子任务改写

- **WHEN** 后续执行者对任一裁决含义产生疑问
- **THEN** 能从本矩阵定位裁决结论、原始输入与验收验证点，并经 `shared_baseline` §2.2 / `architecture_plan` §2 交叉确认
- **AND** 除用户显式补裁外，任何子任务不得改写、放宽或收窄裁决结论；发现冲突时以 `shared_baseline` 单值基线为准并回改

### Requirement: 成功标准必须冻结为可验收条目

本 spec SHALL 冻结以下 4 条成功标准（与 `architecture_plan` §3 逐条对齐），作为 `phase15-08` 验收门禁的判定依据：

1. **web 维护与回看**：用户可在 web 端（Repository detail 进度区）录入三轨进度事件（含 `occurred_at` 补录历史），并以时间轴倒序完整回看推进信息流
2. **agent 双通道消费**：agent 经 `GetProjectBrief` 直读进度摘要（当前 phase + 最新完成任务 + 最近事件），经 `ListProgressEvents` 获取完整事件流（可按 `workflow_type` 过滤）——与 `standards[]` 合成"背景说明 + 进度说明"一次调用
3. **append-only 断言**：无 `Update` RPC；新事件录入不使历史事件消失或变形；"当前 phase / 最新任务"为读取侧派生值，不落库
4. **真实历史回放验证**：用 PSCO 自身真实历史回放（phase14 全程：开启 → 11 个子任务完成 → 收口）验证可被录入为事件流，并完整重建"截至上一对话我们推进到哪"的进度语义

#### Scenario: 成功标准与验收门禁对应

- **WHEN** `phase15-08` 执行验收
- **THEN** 4 条成功标准逐条映射到固定问题取证与十一项裁决门禁（`shared_baseline` §4），无成功标准落空
- **AND** 成功标准 4 对应的 dogfooding 固定录入集逐条明细（task_key / title / occurred_at，取 git log 真实提交时间）为 `phase15-08` spec 必备附件

### Requirement: 非目标必须显式冻结且候选池顺延项留档

本 spec SHALL 显式冻结以下 12 项非目标（与 `dev_plan` §4 / `shared_baseline` §3.7 单值一致）：

1. git 自动采集（`source=git` 仅预留枚举位，不提供写入路径）
2. agent 写回（`source=agent` 仅预留枚举位；CON-09 先消费后维护不变）
3. MCP / CLI 消费通道
4. 模板仓库自动接入
5. 自动同步
6. `standard_bindings` 目标类型扩展
7. 进度事件与 Decision 互链
8. 第六 CRUD 主实体化（独立路由 / 全局导航 / Dashboard 主卡片）
9. `UpdateProgressEvent`（append-only 无更新语义）
10. phase11 `GetProjectContextResponse` 的 `PhaseEntry` / 规则投影改动
11. plan.md 正文接管
12. plan.md 自动同步

本 spec SHALL 同时留档候选池顺延项——`phase14-11` 冻结的候选池经 2026-08-19 用户排序裁决选中"时间轴"（第一优先级）后，其余全部顺延：

| 顺延项 | 承接约束（沿 `phase14-11` 冻结口径） |
|---|---|
| `standard_bindings` 目标类型扩展 | CON-02 可扩展设计已就位（`target_type` CHECK 可扩展、不为目标类型建分表）；扩展由真实绑定需求驱动，不预先扩枚举 |
| agent 写回 | CON-09 先消费后维护顺序不变——在时间轴等消费面稳定且经用户显式裁决后才可进入；不因 phase14 完成、不因时间轴进入而自动解锁 |
| Git 推进跟踪 | 自 `phase11` 起持续顺延的非目标池；须先经独立 `/plan` 裁决排期，不得搭车进入时间轴 phase |
| 模板仓库自动接入 | 同上 |
| 自动同步 | 同上 |

顺延项的最终排序在 `phase15-09` 收口时作为下一阶段进入条件冻结。

#### Scenario: 非目标不被偷渡

- **WHEN** 任一实现类子任务（phase15-03 ~ 07）的设计或代码中出现 12 项非目标对应的能力
- **THEN** 视为超出冻结边界，必须移除或经用户显式补裁后方可保留
- **AND** 顺延项不因本阶段完成而自动解锁——每项均须独立裁决排期

### Requirement: CON-08 时间轴承接口径必须按 T7 裁决后口径继承

本 spec SHALL 留档 `CON-08` 口径继承声明。`phase14-11` 冻结的进入条件第 2 条第 1 项要求：时间轴承接 SHALL 以**新建正规承接位**落地，**禁止复活画像派生形态**。`phase15` 落地形态与该口径逐项对齐：

| `phase14-11` 冻结要求 | phase15 落地形态 |
|---|---|
| 独立数据模型 | `progress_events` 单表（`0013_phase15_progress_timeline.sql`） |
| `.proto` 合同 | `psco.progress.v1` 新合同包（`progress.proto`） |
| web 时间轴展示 | Repository detail 内嵌进度区 |
| agent 可读 | brief `progress = 9` 装配 + `ListProgressEvents` 独立 RPC |
| 禁止复活画像派生形态 | `BriefProgress` 以新字段号 9 新建；槽位 2/3/4 保持 reserved；不重建 `BriefGovernanceProfile / BriefCurrentPhase / BriefTrackType / BriefPhaseStatus` 等已删除消息 |

口径辨析留档（沿 `shared_baseline` §3.5 注）：`phase14-11` spec 中"`brief` 中不回填时间轴字段"的"回填"指向 reserved 槽位 2/3/4 的画像派生字段恢复；phase15 以新字段号 9 新建 `BriefProgress` 正规消息、完整事件流仍经 `ListProgressEvents` 独立入口消费，属该 spec 明示的 `phase15 /plan` 裁决范围内正当演进（裁决⑩，2026-08-19 用户批准）。

#### Scenario: 画像派生形态零复活路径

- **WHEN** 审视 phase15 任一交付物（proto 合同 / DDL / 后端 / 前端 / brief 装配）
- **THEN** 不存在 reserved 槽位 2/3/4 的复用，不存在已删除画像消息的重建，不存在从画像类结构派生时间轴的路径
- **AND** 时间轴全部事实唯一来自 `progress_events` 事件表 + 读取侧派生

### Requirement: 实现级待决点必须登记并归属到子任务

本 spec SHALL 登记边界收敛复核中识别的 3 个实现级待决点，作为对应子任务 spec 的必备裁决项（本 spec 只登记归属与冻结要求，不预先裁决）：

| # | 待决点 | 内容 | 归属子任务 | 冻结要求 |
|---|---|---|---|---|
| DP-1 | web 当前卡数据源通道 | 前端当前 phase 派生卡的取数通道：经 `GetProjectBrief.progress` 摘要消费，或前端自 `ListProgressEvents` 结果派生 | phase15-05 | 必须与 `shared_baseline` §3.4"派生执行位在后端 service 层、web 当前卡与 brief 摘要共用同一实现"单值对齐；防止前端形成第二套派生语义（违反"无第二套进度语义"约束） |
| DP-2 | repository_id 存在性校验承接位 | `CreateProgressEvent` 校验 `repository_id` 不存在 → `invalid_argument` 的实现承接位：progress 模块自拥有 reader / repository 层查询 / FK 23503 错误映射 | phase15-04 | 沿 standard 模式选型并单值冻结（standard 写路径的跨模块校验经 candidate 子包承接，可参照）；无论选何种承接位，`invalid_argument` 对外语义不变 |
| DP-3 | occurred_at 输入与时区口径 | 前端录入控件格式（`datetime-local` 类控件无时区语义）与展示时区；dogfooding 附件 `occurred_at`（git log 提交时间）的时区口径 | phase15-05（录入控件与展示）/ phase15-08（附件时间口径） | 保证补录历史可精确复现，rerun 不因时区漂移失败；附件逐条明细须注明时区 |

#### Scenario: 待决点不丢失

- **WHEN** 对应子任务（phase15-04 / phase15-05 / phase15-08）产出 spec
- **THEN** 该 spec 必须显式冻结所归属的待决点并给出单值结论
- **AND** 待决点的冻结结论不得与十一项裁决及 `shared_baseline` 单值基线冲突

### Requirement: 本 spec 与三件套、根级文档必须单值一致

本 spec 的全部结论 SHALL 与 `phase15` 三件套单值一致；根级文档当前状态（`AGENTS.md` / `plan.md` 中 phase15 已建立 `/plan` 入口、待从 phase15-01 开始执行的表述）保持不变：

- 本 spec 不反向改写三件套正文；发现不一致时以 `shared_baseline` 单值基线为准并回改本 spec
- 本 spec 不回写根级文档（根级状态同步属 `phase15-09`；`phase15-01` 冻结的边界以本目录 spec 为承接位，根级入口可经三件套与 `docs/phase/README.md` 既有登记到达）
- 本 spec 不承载后续子任务的设计细节（合法矩阵展开、DDL、RPC 签名、组件树等），防止范围前置与双源漂移

#### Scenario: 一致性可校验

- **WHEN** 独立复核执行
- **THEN** 本 spec 的边界声明 / 裁决矩阵 / 成功标准 / 非目标 / 顺延项 / CON-08 继承逐项与 `architecture_plan` / `dev_plan` / `shared_baseline` 对应章节比对一致
- **AND** 根级文档无因本 spec 产生的改动（git 工作区中仅本 spec 目录新增）
