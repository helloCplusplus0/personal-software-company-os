# phase15_project_progress_timeline_foundation_architecture_plan

## 1. 文档定位

- 本文档是 `phase15_project_progress_timeline_foundation` 的架构规划与冻结记录，承担该阶段 `/plan` 的架构结论冻结职责。
- 本文档在阶段收口后保留为规划与冻结记录，不覆盖根级当前状态；根级状态只以 `plan.md` 为准。
- 本文档的直接下游是同阶段 `dev_plan` 与 `shared_baseline`；三件套共同构成 `phase15` 全部 `/spec` 子任务的强制边界上游。

## 2. 上游输入

- 唯一直接规划上游：`.trae/specs/phase14_11_sync_root_level_closeout_freeze_phase15_entry_conditions/spec.md`（`phase15` 进入条件冻结：`CON-08` 时间轴按 T7 裁决后口径另行新建正规承接，不复活画像派生形态）与 `phase14-10` acceptance_report（阶段交付实况）
- 排序裁决：用户于 `2026-08-19` 经两轮结构化反思拍板——项目进度时间轴为第一优先级（候选池其余项：`standard_bindings` 目标类型扩展 / agent 写回 / Git 推进跟踪 / 模板仓库自动接入 / 自动同步，全部顺延为后续进入条件）
- 十一项裁决（主裁决 8 项 + 补裁 3 项，2026-08-19 用户拍板）：
  - 裁决①（排序与能力定位）：项目进度时间轴为 `phase15` 唯一主交付；目标是接管用户高频提示词中的"进度说明"段，与 `standards[]`（背景说明）合成 agent 一次调用即可获得的"背景 + 进度"完整上下文
  - 裁决②（承接锚点）：源代码仓库——`progress_events` 以 `repository_id` 为锚点；进度是单仓库事实，与 brief 装配粒度（`GetProjectBrief(repository_id)`）天然对齐
  - 裁决③（数据模型形态）：append-only 推进事件流（单表 `progress_events`）；记录的是推进信息流而非单点信息——历史事件永不丢弃，"当前"仅为派生值不落库（用户强化警惕点：不可把进度理解为只保留最新推进的单点快照）
  - 裁决④（三轨 workflow）：`workflow_type = phase / audit / fix` 三值，与 `docs/` 三目录、`project_rules.md` §4 三套推进链（phase 推进链 / fix 推进链 / audit 推进链）一一对应
  - 裁决⑤（事件颗粒度）：主事件颗粒度 = 任务项级（phase 轨 `phaseNN-MM` / audit 轨 `audit_NNN` / fix 轨 `fix_NNN`，对齐"每个子任务验收通过即 git 提交"的真实节奏）；phase 轨另设 `phase_started / phase_completed` 低频边界标记；audit / fix 轨天然单任务颗粒，仅允许 `task_completed` 与 `note`
  - 裁决⑥（维护方式）：web 手动维护——沿用 Standard 设计策略，在 Repository detail 页内提供维护与展示入口；不建独立 `/progress` 路由与全局导航项，不做第六 CRUD 主实体
  - 裁决⑦（证据引用边界）：`evidence_ref` 为导航引用（沿用 Standard `ref` 规则：`/` 开头仓库内相对路径或 `https://` 开头 URL）；PSCO 不接管、不复制 plan.md / spec 正文
  - 裁决⑧（来源预留）：`source = manual / git / agent`，起步仅实现 `manual`；git 自动采集与 agent 写回为后续 phase 进入条件，本阶段只预留枚举位不加破坏性变更
  - 补裁⑨（RPC 语义）：最小集 3 个 RPC——`ListProgressEvents`（不分页，个人规模，显式冻结）/ `CreateProgressEvent` / `DeleteProgressEvent`；不设 `Update`（误录删了重录，保持 append-only 语义纯净）
  - 补裁⑩（消费分层）：brief 新增 `progress` 块（字段号 9）= 当前派生摘要 + 最近事件；完整事件流走 `ListProgressEvents` 独立 RPC（支持 `workflow_type` 过滤）
  - 补裁⑪（三重边界分离）：与 plan.md（正文事实源在仓库，PSCO 只持事件与引用）、与 phase11 `PhaseEntry`（根级文档只读投影，导航入口）、与 Decision（决策留痕，"为什么"）均不合并不互替
- 参照基线：`PSCO-mvp05-summarize-feedback.md`、`TECH_STACK_BASELINE.md`、`project_rules.md` §2.5 / §2.6、phase14 已验证的建模模式（`proto/psco/<entity>/v1` + `backend/internal/<module>` 分层 + `frontend/src/features/<slice>` + 跨模块 Read 接口）

## 3. 本阶段目标

- 单一主交付：建立项目推进时间轴最小主线——以 repository 锚定的 append-only 事件流承接三轨 workflow 推进事实，web 端可维护可回看，agent 经 brief 与独立 RPC 直读完整信息流；无退役任务（T7 裁决后时间轴当前为零承接位，纯新建）。
- 成功标准：
  1. 用户可在 web 端（Repository detail 进度区）录入三轨进度事件（含 `occurred_at` 补录历史），并以时间轴倒序完整回看推进信息流
  2. agent 经 `GetProjectBrief` 直读进度摘要（当前 phase + 最新完成任务 + 最近事件），经 `ListProgressEvents` 获取完整事件流（可按 `workflow_type` 过滤）——与 `standards[]` 合成"背景说明 + 进度说明"一次调用
  3. append-only 断言成立：无 `Update` RPC；新事件录入不使历史事件消失或变形；"当前 phase / 最新任务"为读取侧派生值，不落库
  4. 用 PSCO 自身真实历史回放验证：phase14 全程（开启 → 11 个子任务完成 → 收口）可被录入为事件流，并完整重建"截至上一对话我们推进到哪"的进度语义

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

- 规划上游：`phase14-11` spec（唯一）+ 2026-08-19 排序裁决
- 执行层上游：本三件套冻结后，`phase15` 全部 `/spec` 子任务以三件套为强制边界上游
- 既有不可回退项继续有效：四实体 + Standard 语义、`.proto + ConnectRPC` 传输主线、canonical owner 单值化、`chi` 仅承担基础设施端点、brief 5 顶层块与槽位 2/3/4 reserved 现状

### 4.2 当前阶段单一主交付能力

- `Project Progress Timeline Foundation`：repository 锚定的三轨 append-only 推进事件流最小完整主线（合同 → 存储 → 后端 → 前端 → agent 消费），一能力贯通，不并列第二主交付。
- 交付边界内的四个组成部分：
  1. `progress_events` 合同与存储（新建）
  2. web 端维护与展示入口（Repository detail 内嵌进度区，新建切片 `features/progress/`）
  3. agent 消费链路（brief `progress = 9` 装配 + `ListProgressEvents` 独立 RPC）
  4. 派生逻辑（当前 phase / 最新任务的读取侧计算，含 web 当前卡与 brief 摘要共用）

### 4.3 Progress 事件流的正式定位

- 正式地位：repository 锚定的治理层信息流，不是第六业务主实体，不进入四实体关系主链。
- 作用域：单仓库作用域（`repository_id` 锚点）；与 Standard 的全局作用域形成对照——规范跨项目复用，进度是单仓库事实。
- 交互入口：Repository detail 页内嵌"项目进度"区（时间轴 + 当前卡 + 录入/删除）；不建独立路由、不加全局导航项、不建 Dashboard 主卡片。
- 信息形态：append-only 事件流（信息流而非单点）；所有消费方（web / agent）读到的是同一持久化事件集合 + 同一派生规则，无第二套进度语义。

### 4.4 数据模型与 pg 承载（裁决③②）

单表 `progress_events`，落在一张编号迁移文件（`0013_phase15_progress_timeline.sql`）：

- `id`（uuid PK）
- `repository_id`（uuid FK → repositories，NOT NULL）
- `workflow_type`（enum：`phase / audit / fix`）
- `event_kind`（enum：`phase_started / phase_completed / task_completed / note`）
- `task_key`（text，可空；kind 与格式约束见 shared_baseline §3.3）
- `title`（text，NOT NULL，一句话）
- `detail`（text，可空，展开说明）
- `evidence_ref`（text，可空，导航引用；格式约束与 Standard `ref` 单值一致）
- `source`（enum：`manual / git / agent`，NOT NULL DEFAULT `manual`；本阶段创建入口仅开放 `manual`）
- `occurred_at`（timestamptz，NOT NULL；用户声明发生时间，允许补录历史，与 `created_at` 分离）
- `created_at`（timestamptz，NOT NULL DEFAULT now()）

索引：`(repository_id, occurred_at DESC, created_at DESC)`（时间轴回看与派生计算的统一读取序）；读取全序为三键链 `(occurred_at DESC, created_at DESC, id DESC)`，`id DESC` 为最终 tiebreak（补录同刻与同事务批量插入两类碰撞场景下保证顺序确定）。
枚举承载冻结为 `TEXT + CHECK(IN ...)`（沿 `0011` 已验证模式，全仓单值，不再选型）；幂等 DDL 细节由 phase15-03 落地。

### 4.5 事件模型语义（裁决④⑤③）

- 三轨 × event_kind 合法矩阵（冻结）：

| workflow_type | phase_started | phase_completed | task_completed | note |
|---|---|---|---|---|
| `phase` | 合法（task_key=`phaseNN`） | 合法（task_key=`phaseNN`） | 合法（task_key=`phaseNN-MM`） | 合法（task_key 可空） |
| `audit` | 禁止 | 禁止 | 合法（task_key=`audit_NNN`） | 合法 |
| `fix` | 禁止 | 禁止 | 合法（task_key=`fix_NNN`） | 合法 |

- 派生规则（读取侧计算，不落库、不缓存）：
  - 当前 phase = 最新 `phase_started` 的 `task_key`（序：`occurred_at DESC, created_at DESC, id DESC` 三键链）；若存在同 key 且序更晚的 `phase_completed`，则为"全部完结"态（当前 phase 为空）
  - 最新完成任务项 = 该 repository 最新一条 `task_completed`（不限 phase，同三键链取最新）
  - brief `recent_events` = 最近 N 条（N=10 冻结）三轨混合事件
- 完整校验规则（9 条）与派生算法细节以 `shared_baseline` §3.3 / §3.4 为单值来源。

### 4.6 内容与边界（裁决⑦⑪）

- PSCO 承接：推进事件事实（何时发生了什么）、任务项标识、一句话标题与展开说明、证据导航引用、来源标识。
- 仓库承接：plan.md / 三件套 / spec 正文。正文事实源在仓库；`evidence_ref` 是导航引用而非正文。
- 与 phase11 `PhaseEntry` 的边界：`GetProjectContextResponse.phases` 是根级文档的只读投影（导航入口），本阶段不改动；`progress_events` 是 PSCO-native 持久化事实（事件流）。两者并存、不合并。
- 与 Decision 的边界：进度事件 = 客观时间轴事实（"发生了什么"）；Decision = 决策留痕（"为什么这么决定"）。不合并、不互替，起步不做互链。

### 4.7 agent 消费链路（GetProjectBrief 演进 + 独立 RPC）

- `GetProjectBriefResponse` 新增 `progress` 字段（字段号 9，下一可用号；槽位 2/3/4 保持 reserved）：
  - 内联轻量消息 `BriefProgress`（定义于 `project_context.proto`，导入 `psco.progress.v1.ProgressEvent`）：`current_phase_key / current_phase_label / latest_task_completed / recent_events[]`；`latest_task_completed` 与 `recent_events[]` 元素同型复用 `ProgressEvent`（不建第二套摘要消息）
  - 空态语义：Go 装配侧恒构造 `progress` 块（非 nil），0 事件时字段为零值、`recent_events` 为空数组——双端单值消费，无双套判空逻辑
  - 装配来源：`projectcontext` 侧经跨模块 Read 接口 `ProgressReader`（沿袭 `StandardReader` 模式）读取派生
- 完整事件流：`psco.progress.v1.ProgressService.ListProgressEvents`（输入 `repository_id` + 可选 `workflow_type` 过滤，不分页，按 `occurred_at DESC`）——agent 需要全量信息流时走本 RPC，不膨胀 brief。
- 读取路径：web 与 agent 同源（同一事件表 + 同一派生规则），符合 phase12 共享只读消费基线。
- brief 顶层块口径演进：5 顶层块 → 5 顶层块 + `progress` 摘要块（`progress = 9` 为 phase15-02 冻结的正式演进）。

### 4.8 合同与传输主线（project_rules §2.6 对齐）

- 新合同包：`proto/psco/progress/v1/progress.proto`，package `psco.progress.v1`。
- RPC 最小集（3 个，裁决⑨）：`ListProgressEvents`（不分页，显式冻结）/ `CreateProgressEvent`（校验规则执行位，`repository_id` 不存在 → `invalid_argument` 语义错误）/ `DeleteProgressEvent`（误录修正；无 `Update`）；全部走 ConnectRPC。
- web 写路径与 agent 读路径同包不同 RPC，不出现第二套 canonical API。

### 4.9 当前阶段验收协议前提

- 固定样本：沿用 `phase13-11` 冻结的 dogfooding 样本 repository（`ca261521-8daf-4248-8f12-43525326e759`），在其上录入进度事件。
- dogfooding 数据：用 PSCO 自身 phase14 真实历史回放（`phase_started phase14` + 11 条 `task_completed` + `phase_completed phase14` + 穿插 `note`）作为固定录入集——PSCO 用自己的进度时间轴讲完自己的 phase14 故事；逐条明细（含 occurred_at，取 git log 真实提交时间）为 phase15-08 spec 必备附件。
- 固定问题取证：agent 直答当前 phase / 最新任务 / 完整事件流（brief 摘要 + ListProgressEvents）；append-only 断言（无 Update RPC、历史零丢失）；派生正确性（phase_completed 后当前 phase 为空）；三轨过滤可用性。
- 工具链门禁：`buf lint / build / breaking`、`go build / vet / test`、前端 `tsc --noEmit`、浏览器反回归矩阵（Repository detail 新增进度区 + 既有页面抽查）。
- 十一项裁决逐条可验证（验收门禁清单见 `shared_baseline` §4）。

### 4.10 当前阶段明确不做

- git 自动采集（`source=git` 预留不实现）与 agent 写回（`source=agent` 预留不实现；CON-09 先消费后维护不变）
- MCP / CLI 消费通道、模板仓库自动接入、自动同步
- `standard_bindings` 目标类型扩展
- 进度事件与 Decision 互链
- 第六 CRUD 主实体化（独立路由 / 全局导航 / Dashboard 主卡片）
- `UpdateProgressEvent`（append-only 无更新语义）
- phase11 `GetProjectContextResponse` 的 `PhaseEntry` / 规则投影改动
- plan.md 正文接管与自动同步（正文事实源在仓库）
