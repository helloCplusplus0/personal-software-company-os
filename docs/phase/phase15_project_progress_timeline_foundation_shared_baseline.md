# phase15_project_progress_timeline_foundation_shared_baseline

## 1. 文档定位

- 本文档集中冻结 `phase15_project_progress_timeline_foundation` 的单值基线与能力矩阵，是该阶段全部 `/spec` 子任务的共享参照。
- 三件套不一致时以本文档为准并回改；阶段收口后本文档保留为该阶段冻结记录。

## 2. 当前单值基线

### 2.1 项目路线

- 当前阶段：`phase15_project_progress_timeline_foundation`（`/plan` 三件套已建立，进入子任务执行）
- 直接规划上游：`phase14-11` spec（`phase15` 进入条件，T7 裁决后口径）+ 2026-08-19 用户排序裁决
- 单一主交付：项目推进时间轴最小主线（三轨 append-only 事件流 + web 维护展示 + brief 进度块 + agent 直读完整流）
- 下一阶段进入条件在 `phase15-09` 收口时冻结（预期含 git 自动采集 / agent 写回等 `source` 预留项的排序）

### 2.2 十一项裁决结论（本阶段最高优先级基线）

前 8 项为主裁决（2026-08-19 用户两轮结构化拍板），后 3 项为补裁（随方案整体获批冻结）：

| 裁决项 | 结论 | 对应输入 |
|---|---|---|
| ① 排序与能力定位 | 项目进度时间轴为唯一主交付；接管高频提示词"进度说明"段，与 `standards[]`（背景说明）合成一次调用的"背景 + 进度" | `phase14-11` 候选池 + 用户反思 |
| ② 承接锚点 | 源代码仓库：`progress_events.repository_id` 锚定，与 brief 装配粒度对齐 | 用户想法 1 |
| ③ 数据模型形态 | append-only 推进事件流（单表）；记录信息流而非单点；"当前"为派生值不落库；历史永不丢弃 | 用户想法 5（强化警惕点） |
| ④ 三轨 workflow | `phase / audit / fix` 对齐 `docs/` 三目录与三推进链 | 用户澄清第一点 |
| ⑤ 事件颗粒度 | 主事件 = 任务项级（`phaseNN-MM` / `audit_NNN` / `fix_NNN`，对齐"子任务验收通过即 git 提交"节奏）；phase 轨设边界标记；audit/fix 轨仅 `task_completed` 与 `note` | 用户澄清第二点 |
| ⑥ 维护方式 | web 手动维护，Repository detail 内嵌维护与展示入口；无独立路由 / 导航项 / Dashboard 主卡片 | 用户想法 2 |
| ⑦ 证据引用边界 | `evidence_ref` 导航引用（与 Standard `ref` 规则单值一致）；不托管 plan.md / spec 正文 | 用户想法 3 + Standard 哲学 |
| ⑧ 来源预留 | `source = manual / git / agent`，起步仅实现 `manual`；git 采集与 agent 写回为后续进入条件 | 后续演进预留 |
| ⑨ RPC 语义（补裁） | 最小集 3 RPC：List（不分页）/ Create / Delete；无 Update（append-only 语义纯净） | 方案获批 |
| ⑩ 消费分层（补裁） | brief `progress = 9` 摘要块 + `ListProgressEvents` 完整流（支持 `workflow_type` 过滤） | 方案获批 |
| ⑪ 三重边界分离（补裁） | 与 plan.md / phase11 `PhaseEntry` / Decision 均不合并不互替 | 方案获批 |

### 2.3 当前阶段正式技术主线

- 合同源：`proto/psco/progress/v1/progress.proto`（package `psco.progress.v1`），`.proto` 唯一长期合同
- 传输：ConnectRPC 正式传输；`chi` 仅基础设施端点（project_rules §2.6）
- 后端：Go 单体 `backend/internal/progress/`（connect / service / repository 分层）；brief 装配经 `projectcontext` 侧 `ProgressReader` 独立 Read 接口（沿袭 `StandardReader` 模式）
- 前端：`frontend/src/features/progress/` 切片（application / components / data / types.ts / index.ts，承接形态为组件级——无独立路由，见 §3.6，沿 `standard` 切片结构），承接位为 Repository detail 内嵌进度区；query 纯只读、mutation 收敛切片固定承接位（project_rules §2.5）
- 存储：`0013_phase15_progress_timeline.sql`（单表 `progress_events`，幂等 DDL；枚举承载冻结为 `TEXT + CHECK(IN ...)`，沿 `0011` 已验证模式）
- brief 演进：`project_context.proto` 内联 `BriefProgress`，`GetProjectBriefResponse.progress = 9`（槽位 2/3/4 保持 reserved）

### 2.4 当前阶段特别约束

- append-only 纯净性：无 `Update` RPC；事件一经录入只有整条删除一种修正路径；任何设计不得使历史事件因新事件而消失或变形
- 派生不落库：当前 phase / 最新任务 / recent_events 全部为读取侧计算，不持久化、不缓存第二套状态
- 三轨合法矩阵单值：audit / fix 轨禁止 `phase_started / phase_completed`；phase 轨边界标记 `task_key = phaseNN`
- 正文零托管：`evidence_ref` 是导航引用而非正文；PSCO 不复制 / 缓存 plan.md / spec 正文（裁决⑦）
- 单仓库作用域：进度事实只锚定 `repository_id`，不做跨仓库聚合视图
- source 现值约束：本阶段创建入口仅开放 `manual`；`git / agent` 为预留枚举值，不提供写入路径
- 反假大空：录入摩擦必须低于"在提示词里手写进度说明"基线，否则违反裁决①初衷，须回退
- 先消费后维护：agent 写回不进入本阶段（CON-09 不变）

### 2.5 当前阶段交付模式

- 沿袭交付型 phase 模式：每个子任务产出 `.trae/specs/phase15_XX_*/` 三件套（spec / tasks / checklist），实现类子任务附代码与验证证据，验收类子任务附 acceptance_report，全部经独立复核后收口。

## 3. 当前阶段能力矩阵

### 3.1 `Project Progress Timeline Foundation` 单值定义

- 一句话定义：repository 锚定的三轨 append-only 推进事件流——以任务项级颗粒承接 phase / audit / fix 三套 workflow 的推进事实，web 可维护可回看、agent 可直读完整信息流，"当前"仅为派生视图。
- 解决的初心问题：项目推进进度不再散落在提示词手写说明与对话记忆中；一次维护、web / agent 共享消费，与 `standards[]` 合成 agent 的"背景 + 进度"完整上下文。

### 3.2 progress_events 字段矩阵

| 字段 | 类型 | 说明 |
|---|---|---|
| `id` | uuid PK | |
| `repository_id` | uuid FK NOT NULL | 锚点；不存在 → 创建时 `invalid_argument` |
| `workflow_type` | enum `phase / audit / fix` | 对齐三推进链 |
| `event_kind` | enum `phase_started / phase_completed / task_completed / note` | |
| `task_key` | text 可空 | 任务项标识；格式随 workflow×kind 变化（§3.3） |
| `title` | text NOT NULL | 一句话标题，上限 200 字符 |
| `detail` | text 可空 | 展开说明，上限 2000 字符 |
| `evidence_ref` | text 可空 | 导航引用：`/` 开头仓库内路径或 `https://` 开头 URL |
| `source` | enum `manual / git / agent`，默认 `manual` | 本阶段仅开放 `manual` |
| `occurred_at` | timestamptz NOT NULL | 用户声明发生时间；允许补录历史 |
| `created_at` | timestamptz NOT NULL | 系统录入时间；与 `occurred_at` 分离 |

索引：`(repository_id, occurred_at DESC, created_at DESC)`；读取全序为三键链 `(occurred_at DESC, created_at DESC, id DESC)`——`id DESC` 为最终 tiebreak，在补录同刻事件与同事务批量插入（未来 `source=git`）两类碰撞场景下保证时间轴顺序与派生结果确定。

### 3.3 事件校验规则（完整版，phase15-03 实现依据）

1. `workflow_type ∈ {phase, audit, fix}`；`event_kind ∈ {phase_started, phase_completed, task_completed, note}`
2. `event_kind = task_completed` 时 `task_key` 必填
3. `workflow_type = phase` 且 `event_kind ∈ {phase_started, phase_completed}`：`task_key` 必填且格式 `^phase[0-9]{2,}$`
4. `workflow_type = phase` 且 `event_kind = task_completed`：`task_key` 必填且格式 `^phase[0-9]{2,}-[0-9]{2,}$`
5. `workflow_type = audit` 且 `event_kind = task_completed`：`task_key` 必填且格式 `^audit_[0-9]{3,}$`
6. `workflow_type = fix` 且 `event_kind = task_completed`：`task_key` 必填且格式 `^fix_[0-9]{3,}$`
7. `workflow_type ∈ {audit, fix}`：`event_kind` 仅允许 `task_completed / note`（禁止 phase 边界标记）
8. `event_kind = note`：`task_key` 可空；若填写不强制格式（自由标注）
9. `evidence_ref` 若填写必须为以 `/` 开头的仓库内相对路径或 `https://` 开头的 URL（与 Standard `ref` 规则单值一致）；`title` 非空上限 200；`detail` 上限 2000；`source` 现值仅接受 `manual`

注：不做 `task_key` 唯一性约束（误录经 Delete 修正；补录与重录为合法场景）；不对 `occurred_at` 做未来时间校验（补录历史是显式需求，排序按声明时间单值执行）。

### 3.4 派生规则矩阵（读取侧计算，phase15-03 冻结算法细节）

| 派生项 | 规则 |
|---|---|
| 当前 phase | 最新 `phase_started`（序：`occurred_at DESC, created_at DESC, id DESC` 三键链）的 `task_key`；若存在同 key 且序更晚的 `phase_completed` → "全部完结"态（空值） |
| 当前 phase label | 该最新 `phase_started` 的 `title` |
| 最新完成任务项 | 该 repository 最新一条 `task_completed` 事件（不限 phase；同三键链取最新） |
| recent_events | 最近 N=10 条三轨混合事件（同上序） |
| 派生执行位 | 后端 service 层统一计算；web 当前卡与 brief 摘要共用同一实现；不落库、不缓存 |

### 3.5 brief 演进矩阵（GetProjectBrief）

| 字段 | 演进 | 说明 |
|---|---|---|
| `progress` | 新增（字段号 9） | 内联轻量消息 `BriefProgress`（定义于 `project_context.proto`）：`current_phase_key / current_phase_label / latest_task_completed / recent_events[]`；`latest_task_completed` 与 `recent_events[]` 元素同型复用 `psco.progress.v1.ProgressEvent`（不建第二套摘要消息） |
| 空态语义 | 无条件装配 | Go 装配侧恒构造 `progress` 块（非 nil）：0 事件时字段为零值、`recent_events` 为空数组——前端进度区与 agent 单值消费，无双套判空逻辑 |
| 槽位 2/3/4 | 保持 reserved | 不复用画像退役槽位；`progress = 9` 为下一可用号 |
| 顶层块口径 | 5 顶层块 + progress 摘要块 | `repository / products / modules / decisions / standards[] + progress`；`progress = 9` 为 phase15-02 冻结的正式演进 |
| 装配来源 | `projectcontext` 经 `ProgressReader` | 沿袭 `StandardReader` 跨模块 Read 模式；brief 与 List RPC 同源同派生 |

注（口径辨析，phase15-02 冻结 brief 前后对照表时须留档）：`phase14-11` spec 中"`brief` 中不回填时间轴字段，时间轴经独立入口消费"的"回填"，指向 reserved 槽位 2/3/4 的画像派生字段恢复；phase15 以新字段号 9 新建 `BriefProgress` 正规消息，完整事件流仍经 `ListProgressEvents` 独立入口消费，属该 spec 明示的 `phase15 /plan` 裁决范围内正当演进（裁决⑩，2026-08-19 用户批准）。

### 3.6 前端承接矩阵

| 承接位 | 形态 | 说明 |
|---|---|---|
| Repository detail 进度区 | 新增内嵌区 | 时间轴倒序列表 + 当前 phase 派生卡 + 录入表单 + 误录删除（带确认） |
| 切片 | `frontend/src/features/progress/` | application / components / data / types.ts / index.ts（沿 `standard` 切片结构） |
| 录入表单交互 | 最小摩擦 | `occurred_at` 默认 now；`event_kind` 记住上次选择；`workflow_type × event_kind` 联动过滤非法组合 |
| 路由与导航 | 无独立路由 | 不加 `/progress` 路由、全局导航项、Dashboard 主卡片 |
| 移动端适配 | 对齐既有基线 | 响应式单列布局 |
| mutation 承接 | 切片内固定承接位 | Create / Delete 收敛 application owner；query 纯只读（§2.5） |

### 3.7 非目标矩阵

git 自动采集 / agent 写回 / MCP / CLI / 模板仓库自动接入 / 自动同步 / `standard_bindings` 目标类型扩展 / 进度事件与 Decision 互链 / 第六 CRUD 主实体化 / `UpdateProgressEvent` / phase11 `PhaseEntry` 投影改动 / plan.md 正文接管——全部为 phase15 非目标，进入条件在 `phase15-09` 冻结。

## 4. 当前阶段验收前提

- 验收环境：phase13-11 固定 dogfooding 样本（`repository_id: ca261521-8daf-4248-8f12-43525326e759`）上录入进度事件；恢复脚本可重复执行。
- dogfooding 固定录入集：PSCO 自身 phase14 真实历史回放（`phase_started phase14` + `phase14-01 ~ phase14-11` 共 11 条 `task_completed` + `phase_completed phase14` + 穿插 `note`）；逐条明细（task_key / title / occurred_at，occurred_at 取仓库 git log 中 phase14 各子任务真实提交时间）为 phase15-08 spec 必备附件，保证 rerun 可复现。
- 验收协议：固定样本 / 固定入口 / 固定问题 / 固定 rerun 记录格式，沿袭 phase13-11 / phase14-10 协议。
- 验收门禁：十一项裁决逐条可验证（①brief 一次调用含背景+进度 / ②repository 锚定 / ③append-only 断言 + 派生不落库 / ④三轨可录可滤 / ⑤任务项颗粒 + phase 边界标记 + audit/fix 无边界事件 / ⑥维护入口仅在 Repository detail / ⑦evidence_ref 导航且正文零托管 / ⑧source 仅 manual 可写 / ⑨无 Update RPC / ⑩brief 摘要与完整流分离 / ⑪与 PhaseEntry、Decision、plan.md 三重边界不被破坏）。
