# phase15-05 前端信息架构与维护入口设计 Spec

## Why

`phase15-04` 已冻结后端合同（3 RPC envelope + 错误语义 + `BriefProgress` 装配演进点），但 Repository detail 内嵌"项目进度"区（组件树 / 交互规格 / 误录删除确认）、`frontend/src/features/progress/` 切片结构、DP-1（web 当前卡数据通道，含 phase15-02 Obs-1 移交事项）与 DP-3（`occurred_at` 录入控件与时区口径）裁决尚未落地到"足以直接进入实现"粒度。本子任务产出前端设计冻结，`phase15-07` 不再做任何设计决策。

`phase15-05` 是实现设计类子任务：纯设计文档冻结，不写任何前端代码、不建切片文件、不改 proto；设计对象为未来的 `frontend/src/features/progress/` 切片与 `repository-binding-detail-page.tsx` 挂载修改，均由 phase15-07 落地。

## What Changes

- 裁决 DP-1：当前 phase 派生卡数据通道 = `GetProjectBrief.progress`（BriefProgress 投影，与 agent 消费同源）；空值两情形统一文案不区分（Obs-1 闭环）；`recent_events` 在 web 端不重复消费
- 裁决 DP-3：`occurred_at` 录入控件 = `datetime-local`（浏览器本地时区语义）→ 提交转 UTC pb Timestamp；展示一律浏览器本地时区；不做未来时间校验
- 冻结 `frontend/src/features/progress/` 切片结构（11 文件清单：data 3 / application 2 / components 4 / types.ts / index.ts；query 纯只读 + mutation 固定承接位，project_rules §2.5）
- 冻结进度区组件树与挂载分区关系（`ProgressSection` 挂载 Repository detail 尾部，与 Standard 只读摘要同级全宽堆叠）
- 冻结录入表单交互规格（最小摩擦：`occurred_at` 默认 now、`event_kind` 记住上次选择〔localStorage〕、`workflow_type × event_kind` 非法组合联动禁用矩阵、`task_key` 联动必填与 placeholder 矩阵、提交回流与重置语义）
- 冻结时间轴列表交互规格（三轨过滤、事件行紧凑结构、空态文案、`window.confirm` 误录删除确认文案）
- 冻结移动端适配基线（对齐 Repository detail 既有基线与紧凑化规范）
- 冻结无独立路由与导航项声明（裁决⑥：零 routes 文件、零 NAV_ITEMS 改动、零 Dashboard 改动）

## Impact

- Affected specs:
  - `docs/phase/phase15_project_progress_timeline_foundation_dev_plan.md`（上游，phase15-05 定义 L39-42）
  - `.trae/specs/phase15_01_freeze_progress_timeline_scope_success_non_goals/spec.md`（DP-1 / DP-3 登记与冻结要求——本 spec 完成两项裁决）
  - `.trae/specs/phase15_02_freeze_event_model_semantics_and_brief_evolution/spec.md`（直接语义上游：合法矩阵 12 格 + K-1~K-5 正则〔联动禁用与 placeholder 依据〕、派生空值同型语义〔DP-1 裁决边界〕、Obs-1 移交事项〔本 spec 闭环〕）
  - `.trae/specs/phase15_03_data_model_validation_derivation_design/spec.md`（校验执行序与错误码总表——前端错误回显消费）
  - `.trae/specs/phase15_04_backend_contract_readwrite_boundary_design/spec.md`（直接合同上游：3 RPC envelope 与错误语义——前端消费模型与 mutation 的唯一合同依据；`BriefProgress` 字段号 1-4——DP-1 通道投影依据）
  - `docs/phase/phase15_project_progress_timeline_foundation_shared_baseline.md`（§2.2 裁决⑥⑧⑨⑩ / §3.6 前端承接矩阵 / §2.4 反假大空约束）
- Affected code: 无（零代码改动；设计对象为未来的 `frontend/src/features/progress/`（11 文件）与 `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`（import + 尾部挂载追加），均由 phase15-07 落地）
- 设计参照（本轮实际读取）：`repository-binding-detail-page.tsx`（挂载点现状 L27/L306 与页面布局）、`standard` 切片 22 文件结构、`use-repository-standards-read.ts`（brief 投影 query owner 先例）、`use-create-standard.ts`（mutation owner 与 normalizeError 先例）、`standard/index.ts`（barrel 受控导出先例）、`standard-readonly-summary.tsx`（Repository detail 内嵌区先例）、`standard-detail-page.tsx`（`window.confirm` 删除确认先例 L45）

## ADDED Requirements

### Requirement: DP-1 裁决——当前 phase 派生卡数据通道必须单值冻结

本 spec SHALL 裁决 DP-1（phase15-01 登记："前端当前 phase 派生卡的取数通道：经 `GetProjectBrief.progress` 摘要消费，或前端自 `ListProgressEvents` 结果派生"），并闭环 phase15-02 Obs-1（"空值同型后 phase15-05 若经 brief 通道则前端无法区分'尚未开始/已完结'文案"）：

**裁决结论**：

1. **数据通道 = `GetProjectBrief.progress`（BriefProgress 投影），唯一通道**。当前卡 query owner 为切片内 `data/use-repository-progress-read.ts`（query key `['repository-progress', repositoryId]`，调 `projectContextClient.getProjectBrief` 仅投影 `progress` 块）——与 agent 消费主路径同源，沿 `use-repository-standards-read` 的"切片自有 brief 投影 owner"模式（不跨切片消费、不新增 RPC、不建第二通道）。
2. **前端不得自 `ListProgressEvents` 结果计算派生值用于当前卡**（phase15-02 封死："后端统一派生"约束；`ListProgressEvents` 结果仅用于时间轴原始事件展示）。
3. **空值两情形统一文案**：`current_phase_key` 为空串时（含"从未开始"与"全部完结"两种同型零值情形）统一展示"暂无进行中 phase"，**不做区分**——区分两种情形将要求前端第二套派生语义（phase15-02 封死）或后端第二套 phase 状态字段（违反"同型零值、不引入第二套状态字段或状态枚举"冻结），均被上游禁止。此为 Obs-1 的裁决答案：接受 brief 通道的空值同型约束。
4. **"完结态"用户体验承接口径**（显式留档，防 phase15-07 实现期误解）：dev_plan phase15-07 DoD"phase_completed 后当前卡转完结态"的合规落地 = 当前卡从进行中态（展示 phase key + label）转为空态统一文案"暂无进行中 phase"，且时间轴最新事件行呈现刚录入的 `phase_completed` 边界事件（原始事件展示，非派生）——用户经两者组合确认完结事实，无需当前卡单独区分空值情形。
5. **web 端当前卡消费 `BriefProgress` 的三字段**：`current_phase_key` / `current_phase_label` / `latest_task_completed`；`recent_events` 在 web 端**不重复消费**（时间轴经 `ListProgressEvents` 承接完整流展示，同屏不出现第二份事件列表；`recent_events` 的消费方是 agent brief 通道）。

#### Scenario: 当前卡实现判定

- **WHEN** phase15-07 实现当前卡
- **THEN** 取数唯一经 `use-repository-progress-read`（brief 投影）；组件内无任何从事件集合计算当前 phase / 最新任务的逻辑；空态文案统一为"暂无进行中 phase"
- **AND** `grep` 切片内不出现 `ListProgressEvents` 结果参与派生计算的代码路径（时间轴消费除外）

### Requirement: DP-3 裁决——occurred_at 录入控件与时区口径必须单值冻结

本 spec SHALL 裁决 DP-3 的 phase15-05 部分（录入控件与展示时区；dogfooding 附件时间口径归 phase15-08）：

| 项 | 冻结内容 |
|---|---|
| 输入控件 | `<input type="datetime-local">`（浏览器原生控件，分钟粒度） |
| 默认值 | 表单挂载时与提交成功重置时 = 浏览器本地当前时刻（`YYYY-MM-DDTHH:mm` 格式） |
| 提交转换 | datetime-local 值按浏览器本地时区语义解析（`new Date(value)` 本地解析）→ 构造 pb Timestamp（UTC）；补录历史 = 直接修改输入值，转换规则不变 |
| 展示口径 | 事件 `occurred_at`（及当前卡 `latest_task_completed.occurred_at`）一律浏览器本地时区展示（`toLocaleString` 或等价格式化）；`created_at` 不在时间轴行内展示（见时间轴 Requirement） |
| 未来时间 | 不做未来时间校验（phase15-02 冻结"occurred_at 未来时间不校验"；排序按声明时间单值执行） |
| 附件口径 | dogfooding 固定录入集附件的 `occurred_at` 时区标注归 phase15-08（本 spec 不越界；本裁决保证 rerun 不因前端时区漂移失败——录入与展示同为浏览器本地时区闭环） |

#### Scenario: 时区闭环判定

- **WHEN** 用户在本地时区（如 Asia/Shanghai）录入 `occurred_at` 为 `2026-08-19 10:30` 并提交
- **THEN** 请求体 Timestamp 为对应 UTC 时刻；刷新页面后时间轴展示仍为本地时区 `2026-08-19 10:30`（无时区漂移）
- **AND** 补录历史（git log 真实提交时间，本地时区）与实时录入走同一转换路径

### Requirement: 切片结构必须冻结（文件级）

`frontend/src/features/progress/` SHALL 按以下 11 文件清单实现（结构沿 `standard` 切片模式；types 对齐后端 snake_case）：

**data/（3 文件，全部纯只读 query owner）**

| 文件 | query key | 承接 |
|---|---|---|
| `connect-client.ts` | — | 导出 `progressClient`（ProgressService）与 `projectContextClient`（brief 读取用，DP-1 通道） |
| `use-progress-events-read.ts` | `['progress-events', repositoryId, filter]` | `ListProgressEvents` → `ProgressEvent[]`（filter：`'all'` 不传过滤 = 三轨全量；`'phase'/'audit'/'fix'` 传对应枚举；结果即后端三键链倒序，前端不重排序） |
| `use-repository-progress-read.ts` | `['repository-progress', repositoryId]` | `GetProjectBrief` → 仅投影 `progress` 块 → `ProgressSummary`（DP-1 通道；空态恒有值——后端恒构造块） |

**application/（2 文件，固定 mutation 承接位）**

| 文件 | RPC | 失效矩阵 |
|---|---|---|
| `use-create-progress-event.ts` | `CreateProgressEvent` | `['progress-events', repositoryId]` 前缀（覆盖全部过滤变体）+ `['repository-progress', repositoryId]` |
| `use-delete-progress-event.ts` | `DeleteProgressEvent` | 同上（删除同样改变事件流与派生摘要） |

**components/（4 文件）**：`progress-section.tsx`（进度区外壳与编排，Repository detail 挂载位）/ `progress-current-phase-card.tsx`（当前 phase 派生卡）/ `progress-event-form.tsx`（录入表单）/ `progress-timeline-list.tsx`（时间轴 + 三轨过滤 + 事件行 + 删除确认触发）

**根文件（2 文件）**：
- `types.ts`：`WorkflowType` / `EventKind` / `ProgressSource` string union + `ProgressEvent`（11 字段）+ `ProgressSummary`（四字段）+ `CreateProgressEventInput`（表单提交模型）+ pb 转换（`pbToProgressEvent` / 枚举映射 / datetime-local ↔ pb Timestamp 转换）+ K-1~K-4 正则常量 + 合法组合判定辅助（联动禁用依据）
- `index.ts`：barrel，仅导出 `ProgressSection`（沿 `standard/index.ts` 受控导出模式——Repository detail 挂载位唯一出口；不导出内部组件 / query / mutation owner，防切片外散装拼装）

**切片纪律（project_rules §2.5）**：

- query 层零写动作；Create / Delete 两个写 RPC 全部收敛 application 2 owner，页面与组件不内联 `useMutation`
- 错误归一化收敛 application owner 内（`normalizeError` 提取 ConnectError rawMessage，沿 `use-create-standard` 模式），组件仅消费归一化后 Error
- `event_kind` 记忆经 localStorage（纯 UX 偏好默认值，见表单 Requirement），不构成数据事实源

#### Scenario: 切片实现判定

- **WHEN** phase15-07 落地切片
- **THEN** 文件清单与本表一一对应（11 文件，无第五目录）；`grep useMutation` 仅命中 application 2 文件；失效矩阵逐字实现；barrel 仅导出 `ProgressSection`

### Requirement: 进度区组件树与挂载分区关系必须冻结

**挂载点**：`frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`——import 区追加 `ProgressSection`（from `@/features/progress`），渲染区在 `<StandardReadonlySummary />`（现 L306）之后追加 `<ProgressSection repositoryId={repositoryId} />`。

**分区关系（单值）**：进度区为 Repository detail 页面级**第三全宽区块**（与 Standard 只读摘要同级，位于其后），不进入 `grid lg:grid-cols-3` 摘要/工作台区；区间垂直堆叠由页面容器 `space-y-4` 承接（与 Standard 摘要区现状一致），进度区内部子区域以 `border-t pt-2` 分隔。理由：进度区含表单与列表需要全宽空间；Standard 摘要已是全宽同级先例。

**组件树（逐层冻结）**：

```
ProgressSection（section.min-w-0.space-y-2）
├── 区块标题："项目进度"（text-xs font-medium text-muted-foreground，沿 StandardReadonlySummary 标题模式）
├── ProgressCurrentPhaseCard
│   ├── 进行中态：current_phase_key（code/mono 样式）+ current_phase_label（text-xs）
│   │   └── 最新完成任务行：latest_task_completed 的 task_key + title + occurred_at（本地时区，text-xs text-muted-foreground）
│   ├── 空态（current_phase_key 空串）：统一文案"暂无进行中 phase"（text-xs text-muted-foreground，不区分两情形）
│   └── latest_task_completed 无值时：附行"暂无任务完成记录"（text-xs text-muted-foreground）
│   └── 读取失败："进度摘要读取失败"（text-xs text-destructive）；加载态："加载中..."（text-xs text-muted-foreground）
├── ProgressEventForm（border-t pt-2 分隔）
└── ProgressTimelineList（border-t pt-2 分隔）
```

无路由声明（裁决⑥）：零 `routes/` 文件新增、零 `__root.tsx` NAV_ITEMS 改动、零 Dashboard 卡片改动——进度区唯一入口 = Repository detail 内嵌。

#### Scenario: 挂载实现判定

- **WHEN** phase15-07 落地挂载
- **THEN** detail 页渲染顺序为 …grid 三列区 → StandardReadonlySummary → ProgressSection；全站 `grep` 无 `/progress` 路由与导航项；`tsc --noEmit` 零错误

### Requirement: 录入表单交互规格必须冻结（最小摩擦 + 联动禁用）

`ProgressEventForm` SHALL 满足以下规格（反假大空约束：录入摩擦必须低于"在提示词里手写进度说明"基线）：

**字段清单与默认值**：

| 字段 | 控件 | 必填性 | 默认值 / placeholder |
|---|---|---|---|
| `workflow_type` | select（phase / audit / fix） | 必填 | 默认 `phase` |
| `event_kind` | select（四值，非法组合 option disabled） | 必填 | 默认 = localStorage 记忆值（无记忆或非法时回退 `task_completed`） |
| `task_key` | text input | 按联动矩阵 | placeholder 按联动矩阵（下表） |
| `title` | text input | 必填 | placeholder"一句话标题" |
| `detail` | textarea（rows=2） | 可选 | — |
| `evidence_ref` | text input | 可选 | placeholder"https://… 或 /仓库内路径" |
| `occurred_at` | datetime-local | 必填 | 默认浏览器本地当前时刻（DP-3） |
| `source` | **不暴露输入位**（裁决⑧：创建入口仅 manual；请求不设置 source 字段 → 后端归一 manual，phase15-04 合同决策 3） | — | — |

**`workflow_type × event_kind` 联动禁用矩阵**（phase15-02 合法矩阵 12 格的 UI 投影）：

| workflow_type | event_kind 选项态 | 切换时重置规则 |
|---|---|---|
| `phase` | 四选项全部可选 | — |
| `audit` | `phase_started` / `phase_completed` **disabled**；`task_completed` / `note` 可选 | 当前 event_kind 变为非法时自动重置为 `task_completed` |
| `fix` | 同 `audit` | 同上 |

**`task_key` 联动必填与 placeholder 矩阵**（K-1~K-4 正则的 UI 投影）：

| workflow_type × event_kind | task_key | placeholder |
|---|---|---|
| `phase` × `phase_started` / `phase_completed` | 必填 | `phaseNN（如 phase15）` |
| `phase` × `task_completed` | 必填 | `phaseNN-MM（如 phase15-05）` |
| `audit` × `task_completed` | 必填 | `audit_NNN（如 audit_001）` |
| `fix` × `task_completed` | 必填 | `fix_NNN（如 fix_001）` |
| 三轨 × `note` | 可选 | `自由标注，可留空` |

**`event_kind` 记住上次选择（localStorage，单值冻结）**：

- key：`psco.progress.last-event-kind`
- 写入时机：Create 成功后写入当次 `event_kind`
- 读取时机：表单挂载初始化 form state 时读取
- 非法值回退：记忆值不在四枚举内 → 回退 `task_completed`
- 合规声明：记忆值仅为表单默认值（纯 UX 偏好），不构成数据事实源，不参与任何派生

**校验反馈双层模型**（沿 phase14-05 模式）：

- 前端轻量层：行内即时提示（title 必填与上限 200 / detail 上限 2000 / evidence_ref 前缀 `/` 或 `https://` / task_key 按矩阵必填与格式）——仅 UX 提示，不阻断输入
- 后端权威层：phase15-03 校验执行序 6 步与错误码在 Create 响应错误中返回；前端经 `normalizeError` 行内回显（含 `REPOSITORY_NOT_FOUND` 等跨模块校验错误——前端不做 repository 存在性预校验，交后端权威判定）

**提交回流与重置语义（单值冻结）**：

- 提交：经 `use-create-progress-event`（owner 组装 pb 请求：枚举映射 + `occurred_at` 本地解析转 UTC Timestamp + 可空文本空串直传）
- 成功后：失效矩阵生效（时间轴 + 当前卡刷新）；表单重置 = `title` / `task_key` / `detail` / `evidence_ref` 清空 + `occurred_at` 重置为 now + `workflow_type` / `event_kind` **保持当前选择**（连续录入同轨同类事件是主场景，如回放 11 条 task_completed）；`event_kind` 记忆已随成功写入
- 失败：停留表单上下文保留已填值，错误行内回显（`text-xs text-destructive`）
- 成功后不弹窗不跳转（内嵌区就地刷新——与独立 Create 页回流详情页模式的差异是设计内差异：进度录入是高频短操作）

#### Scenario: 表单实现判定

- **WHEN** phase15-07 落地表单
- **THEN** 联动禁用矩阵与 placeholder 矩阵逐格可验证；切换 audit/fix 后 `phase_started` 不可选且当前值自动重置；刷新页面后 `event_kind` 默认值 = 上次成功录入值；提交成功后表单按重置语义复位且时间轴顶部出现新事件
- **AND** `source` 无输入位；请求体不含 source 字段

### Requirement: 时间轴列表与误录删除交互规格必须冻结

`ProgressTimelineList` SHALL 满足以下规格：

**三轨过滤**：

- 状态：组件 state，`'all' | 'phase' | 'audit' | 'fix'`，默认 `'all'`
- 控件：四枚紧凑按钮组（全部 / Phase / Audit / Fix；`h-7 text-xs`；active 态高亮，inactive `ghost`）
- 联动：切换即变更 query key 第三段（`['progress-events', repositoryId, filter]`）自动重新查询；`'all'` 不传过滤参数（UNSPECIFIED = 三轨全量，phase15-04 合同决策 2）

**事件行结构（紧凑化规范，`divide-y` 列表，每行 `p-2 space-y-1 min-w-0`）**：

```
第一行（flex flex-wrap items-center gap-1.5 text-xs）
├── event_kind Badge（variant 按 workflow_type 区分：phase=default / audit=secondary / fix=outline）
├── workflow_type 文本标签（text-muted-foreground）
├── task_key（code 样式 text-[10px]，未填不渲染）
├── occurred_at（ml-auto，本地时区，text-muted-foreground）
├── source Badge（variant=outline，text-[10px]；现值恒 manual）
└── 删除按钮（ghost icon 按钮，Trash 图标，h-6 w-6）
第二行：title（text-xs font-medium truncate）
第三行（可选，detail 非空时）：detail（text-xs text-muted-foreground line-clamp-2）
第四行（可选，evidence_ref 非空时）：<a> 链接（text-xs underline truncate；https:// 开头 target=_blank rel=noreferrer；/ 开头普通 href 导航——导航引用零托管，仅链接不解析内容，裁决⑦）
```

- `created_at` 不在行内展示（occurred_at 是时间轴主序与用户声明时间；紧凑化防过载）
- 前端不重排序（后端三键链倒序直渲）、不过滤（过滤唯一经 query 参数）

**空态文案**：

- 全量空（无任何事件）：`暂无推进事件，从上方表单录入第一条。`
- 过滤后空：`该轨暂无事件。`
- 读取失败：`进度事件读取失败`（text-xs text-destructive）
- 加载态：`加载中...`（text-xs text-muted-foreground）

**误录删除确认（DoD 要求文案冻结，单值）**：

- 触发：事件行删除按钮 → `window.confirm`（沿 `standard-detail-page.tsx` L45 删除确认先例，不引入 Dialog 组件）
- 确认文案（逐字冻结）：`确认删除事件「{title}」？删除仅用于修正误录，操作不可撤销。`
- 确认后：经 `use-delete-progress-event`（参数仅事件 `id`）；成功后失效矩阵生效（时间轴 + 当前卡刷新）；失败行内错误回显
- 无软删除、无二次修正入口（append-only 语义，裁决⑨：整条删除是唯一修正路径）

#### Scenario: 时间轴实现判定

- **WHEN** phase15-07 落地时间轴
- **THEN** 过滤切换触发重新查询且 URL 不变（纯组件 state，无路由参数）；删除确认文案逐字一致；删除成功后事件从列表消失且当前卡同步刷新
- **AND** 列表渲染顺序与 `ListProgressEvents` 响应顺序一致（无前端排序代码）

### Requirement: 移动端适配基线与上游单值一致声明必须冻结

**移动端适配基线**（对齐 Repository detail 既有基线与紧凑化规范）：

- 表单字段：`grid grid-cols-1 gap-2 sm:grid-cols-2`（workflow_type / event_kind / task_key / occurred_at 两列区；title / detail / evidence_ref 全宽）；提交按钮 `h-8`
- 时间轴事件行：`flex flex-wrap` + `min-w-0` + `truncate`（防横向溢出）；过滤按钮组 `flex flex-wrap`
- 当前卡：单列紧凑（text-xs 密度）
- 区块标题与子区域分隔沿紧凑化规范（`text-xs` / `border-t pt-2`）；容器型元素仅 `:focus-visible`

**与语义 / 合同上游单值一致声明**：

- 联动禁用矩阵 ↔ phase15-02 合法矩阵 12 格（audit/fix 禁 phase 边界标记 = 规则 7；note 可选 task_key = 规则 8）
- task_key placeholder 矩阵 ↔ K-1~K-4 正则（提示语与格式一一对应，不收窄不放宽）
- 错误回显 ↔ phase15-03 错误码总表 + phase15-04 3 RPC 错误语义（Create InvalidArgument / Internal；Delete InvalidArgument〔`INVALID_PROGRESS_EVENT_ID`〕/ NotFound / Internal——前端按归一化 message 回显，不自行翻译错误码）
- 切片结构与承接形态 ↔ shared_baseline §3.6 前端承接矩阵 + project_rules §2.5（query 纯只读 / mutation 固定承接位 / 切片优先不晋升 shared）
- 无独立路由 ↔ 裁决⑥；source 不暴露 ↔ 裁决⑧；删除唯一修正路径 ↔ 裁决⑨；DP-1 通道分层 ↔ 裁决⑩
- 不偷渡声明：本 spec 不承载 phase15-07 的实现代码、后端 / proto / 迁移任何改动（phase15-06 范围）、dogfooding 附件时间明细（phase15-08 范围）

#### Scenario: 一致性可校验

- **WHEN** 独立复核执行
- **THEN** 联动矩阵 / placeholder / 错误语义 / 切片纪律与 phase15-02 / 03 / 04 及 shared_baseline §3.6 逐项比对一致；DP-1 / DP-3 裁决与 phase15-01 登记的冻结要求逐项对齐（DP-1："后端统一派生"约束满足；DP-3："补录历史可精确复现，rerun 不因时区漂移失败"满足）
- **AND** git 工作区中本 spec 仅为目录新增，零代码 / 零 proto / 零迁移文件 / 零根级文档改动
