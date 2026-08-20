# phase15-07 落实前端主线 Spec

## Why

`phase15-05` 已将前端设计冻结到"足以直接进入实现"粒度（切片 11 文件清单 / DP-1 当前卡数据通道 / DP-3 occurred_at 控件与时区口径 / 组件树与挂载分区 / 表单与时间轴交互规格 / 移动端基线），`phase15-06` 已提供合同生成物（`frontend/src/gen/proto/psco/progress/v1/progress_pb.ts` + project_context 再生成物含 `BriefProgress`）与 3 RPC 后端运行时。本子任务将其落地为可运行前端代码：`frontend/src/features/progress/` 切片 + Repository detail 内嵌进度区。`phase15-08`（联调验收与 dogfooding）以本子任务的浏览器完整会话为前提。

`phase15-07` 是源代码实现类子任务：全部设计决策已在 phase15-05 冻结，本 spec 不新增任何设计决策，只承载"逐字转写关系 + 实现顺序 + 浏览器会话 DoD 验收"。

## What Changes

- 新建 `frontend/src/features/progress/` 切片，11 文件清单（phase15-05 §"切片结构必须冻结"逐字落地）：
  - `data/`（3 文件）：`connect-client.ts`（导出 `progressClient` + `projectContextClient`）、`use-progress-events-read.ts`（query key `['progress-events', repositoryId, filter]`）、`use-repository-progress-read.ts`（query key `['repository-progress', repositoryId]`，DP-1 通道）
  - `application/`（2 文件）：`use-create-progress-event.ts`、`use-delete-progress-event.ts`（mutation 固定承接位，失效矩阵逐字：两 key 前缀 + 当前卡）
  - `components/`（4 文件）：`progress-section.tsx`（外壳与编排）、`progress-current-phase-card.tsx`、`progress-event-form.tsx`、`progress-timeline-list.tsx`
  - `types.ts`（string union + 11 字段模型 + pb 转换 + K-1~K-4 正则常量 + 合法组合判定辅助）、`index.ts`（barrel 仅导出 `ProgressSection`）
- 修改 `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`：import 区追加 `ProgressSection`（from `@/features/progress`），渲染区在 `<StandardReadonlySummary />`（现 L306）之后追加 `<ProgressSection repositoryId={repositoryId} />`
- 零路由 / 零导航 / 零 Dashboard 改动（裁决⑥）；零 `source` 输入位（裁决⑧）；零后端 / proto / 迁移改动（phase15-06 已收口）
- 运行时前提（phase15-06 留档观察 3 的承接）：当前运行中的后端服务器不含 phase15-06 progress RPC——浏览器完整会话验证前**由用户重启后端服务器**（重启即载入新代码 + `RunMigrations` 空过补登记 0013，DDL 幂等无风险）；前端 vite 热更新自动生效；本子任务代理不得擅自重启服务器，实现完成后向用户请求重启再执行浏览器验证

## Impact

- Affected specs:
  - `phase15-05`（直接设计上游，本 spec 为其实现落地；全部内容逐字转写，零再决策）
  - `phase15-02`（语义上游：合法矩阵 12 格 / K-1~K-5 正则——联动禁用与 placeholder 依据；派生空值同型语义——DP-1 统一文案依据；经 05 间接承接）
  - `phase15-03`（校验执行序与错误码总表——前端错误回显消费）
  - `phase15-04`（合同上游：3 RPC envelope 与错误语义——前端消费模型唯一合同依据；`BriefProgress` 字段 1-4——DP-1 通道投影依据）
  - `phase15-06`（运行时上游：3 RPC + 合同生成物 + brief progress = 9 装配；留档观察 3 的补登记前提由本子任务浏览器验证承接）
  - `phase15-08`（下游：dogfooding 固定录入集与反回归矩阵依赖本子任务浏览器主线可用）
- Affected code:
  - 新建：`frontend/src/features/progress/`（11 文件）
  - 修改：`frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`（import 1 行 + 挂载 1 行 + 相关注释条目）
- 不做（边界冻结）：无 `routes/` 改动、无 `__root.tsx` NAV_ITEMS 改动、无 Dashboard 改动、无 shared/ 晋升（切片优先，project_rules §2.5）、无 dogfooding 固定录入集与验收报告（phase15-08）、无根级文档回写（phase15-09）

## ADDED Requirements

### Requirement: 切片结构必须按 11 文件清单逐字落地

本子任务 SHALL 按 phase15-05 §"切片结构必须冻结"实现 `frontend/src/features/progress/`（结构沿 `standard` 切片模式；types 对齐后端 snake_case）：

**data/（3 文件，全部纯只读 query owner）**：

| 文件 | query key | 承接 |
|---|---|---|
| `connect-client.ts` | — | 导出 `progressClient`（ProgressService）与 `projectContextClient`（brief 读取用，DP-1 通道）；沿 `standard/data/connect-client.ts` 模式（createClient + shared transport，页面组件不得直接 createClient） |
| `use-progress-events-read.ts` | `['progress-events', repositoryId, filter]` | `ListProgressEvents` → `ProgressEvent[]`（filter：`'all'` 不传过滤 = 三轨全量；`'phase'/'audit'/'fix'` 传对应枚举；结果即后端三键链倒序，前端不重排序） |
| `use-repository-progress-read.ts` | `['repository-progress', repositoryId]` | `GetProjectBrief` → 仅投影 `progress` 块 → `ProgressSummary`（DP-1 通道；空态恒有值——后端恒构造块；沿 `use-repository-standards-read` 切片自有 brief 投影 owner 模式） |

**application/（2 文件，固定 mutation 承接位）**：

| 文件 | RPC | 失效矩阵 |
|---|---|---|
| `use-create-progress-event.ts` | `CreateProgressEvent` | `['progress-events', repositoryId]` 前缀（覆盖全部过滤变体）+ `['repository-progress', repositoryId]` |
| `use-delete-progress-event.ts` | `DeleteProgressEvent` | 同上（删除同样改变事件流与派生摘要） |

**components/（4 文件）**：`progress-section.tsx`（进度区外壳与编排，Repository detail 挂载位）/ `progress-current-phase-card.tsx`（当前 phase 派生卡）/ `progress-event-form.tsx`（录入表单）/ `progress-timeline-list.tsx`（时间轴 + 三轨过滤 + 事件行 + 删除确认触发）

**根文件（2 文件）**：
- `types.ts`：`WorkflowType` / `EventKind` / `ProgressSource` string union + `ProgressEvent`（11 字段）+ `ProgressSummary`（四字段）+ `CreateProgressEventInput`（表单提交模型）+ pb 转换（`pbToProgressEvent` / 枚举映射 / datetime-local ↔ pb Timestamp 转换）+ K-1~K-4 正则常量 + 合法组合判定辅助（联动禁用依据）
- `index.ts`：barrel，仅导出 `ProgressSection`（沿 `standard/index.ts` 受控导出模式；不导出内部组件 / query / mutation owner）

**切片纪律（project_rules §2.5）**：query 层零写动作；Create / Delete 两个写 RPC 全部收敛 application 2 owner，页面与组件不内联 `useMutation`；错误归一化收敛 application owner 内（`normalizeError` 提取 ConnectError rawMessage，沿 `use-create-standard` 模式）；`event_kind` 记忆经 localStorage（纯 UX 偏好默认值，不构成数据事实源）。

#### Scenario: 切片实现判定

- **WHEN** 切片落地
- **THEN** 文件清单与本表一一对应（11 文件，无第五目录）；`grep useMutation` 仅命中 application 2 文件；失效矩阵逐字实现；barrel 仅导出 `ProgressSection`

### Requirement: DP-1 当前卡数据通道必须单值落地

本子任务 SHALL 按 phase15-05 §"DP-1 裁决"实现当前 phase 派生卡数据通道：

1. 数据通道 = `GetProjectBrief.progress`（BriefProgress 投影），唯一通道——query owner 为 `use-repository-progress-read`，与 agent 消费主路径同源
2. 前端不得自 `ListProgressEvents` 结果计算派生值用于当前卡（`ListProgressEvents` 结果仅用于时间轴原始事件展示）
3. 空值两情形统一文案：`current_phase_key` 为空串时（含"从未开始"与"全部完结"两种同型零值情形）统一展示"暂无进行中 phase"，不做区分
4. "完结态"用户体验承接口径：`phase_completed` 后当前卡从进行中态（phase key + label）转为空态统一文案"暂无进行中 phase"，且时间轴最新事件行呈现刚录入的 `phase_completed` 边界事件（原始事件展示，非派生）——用户经两者组合确认完结事实
5. web 端当前卡消费 `BriefProgress` 三字段：`current_phase_key` / `current_phase_label` / `latest_task_completed`；`recent_events` 在 web 端不重复消费

#### Scenario: 当前卡实现判定

- **WHEN** 当前卡实现
- **THEN** 取数唯一经 `use-repository-progress-read`（brief 投影）；组件内无任何从事件集合计算当前 phase / 最新任务的逻辑；空态文案统一为"暂无进行中 phase"
- **AND** `grep` 切片内不出现 `ListProgressEvents` 结果参与派生计算的代码路径（时间轴消费除外）

### Requirement: DP-3 occurred_at 控件与时区口径必须单值落地

本子任务 SHALL 按 phase15-05 §"DP-3 裁决"实现：

| 项 | 冻结内容 |
|---|---|
| 输入控件 | `<input type="datetime-local">`（浏览器原生控件，分钟粒度） |
| 默认值 | 表单挂载时与提交成功重置时 = 浏览器本地当前时刻（`YYYY-MM-DDTHH:mm` 格式） |
| 提交转换 | datetime-local 值按浏览器本地时区语义解析（`new Date(value)` 本地解析）→ 构造 pb Timestamp（UTC）；补录历史 = 直接修改输入值，转换规则不变 |
| 展示口径 | 事件 `occurred_at`（及当前卡 `latest_task_completed.occurred_at`）一律浏览器本地时区展示（`toLocaleString` 或等价格式化）；`created_at` 不在时间轴行内展示 |
| 未来时间 | 不做未来时间校验（排序按声明时间单值执行） |

#### Scenario: 时区闭环判定

- **WHEN** 用户在本地时区录入 `occurred_at` 并提交
- **THEN** 请求体 Timestamp 为对应 UTC 时刻；刷新页面后时间轴展示仍为本地时区同一时刻（无时区漂移）；补录历史与实时录入走同一转换路径

### Requirement: 挂载与组件树必须按冻结落地

本子任务 SHALL 按 phase15-05 §"进度区组件树与挂载分区关系必须冻结"实现，容器形态与挂载分区按 **2026-08-19 用户两轮 UI 反馈裁决**修订：

**挂载点（2026-08-19 第二轮修订，覆盖 phase15-05 原"页面级第三全宽区块"冻结）**：`repository-binding-detail-page.tsx`——`ProgressSection` 挂载于**绑定工作台右列**（`space-y-4 lg:col-span-2`）内"相关决策"面板之后，成为工作台第四卡片，与"已绑定产品 / 已映射模块 / 相关决策"同列堆叠保持风格一致性（用户反馈：全宽区块游离于主体 grid 之后、Standard 注释区之前是糟糕设计）。页面最终渲染顺序：grid 三列区（左列仓库摘要 / 右列 已绑定产品 → 已映射模块 → 相关决策 → **项目进度**）→ StandardReadonlySummary（页面底部注释风格收尾，用户裁决保持）。

**分区关系（修订后）**：进度区为工作台右列第四卡片（Card 容器）；Standard 只读摘要是页面底部的外部关联注释区（裸 section 先例保持不动，风格断层不再存在——两类区块各归其位）。

**容器形态（2026-08-19 第一轮修订）**：进度区是维护工作台（表单录入 + 删除 + 过滤），容器为 Card 组件（`data-slot="card"`，rounded-lg border bg-card py-4）+ `CardTitle`（text-base leading-none font-semibold，与工作台卡片同款）。

**组件树（逐层冻结，2026-08-19 两轮修订版）**：

```
ProgressSection（Card min-w-0：flex flex-col gap-3 rounded-lg border bg-card py-4）
├── CardHeader / CardTitle："项目进度"（text-base leading-none font-semibold）
├── CardContent（px-4 space-y-3）
│   ├── ProgressCurrentPhaseCard
│   │   ├── 进行中态：current_phase_key（code/mono 样式）+ current_phase_label（text-xs）
│   │   │   └── 最新完成任务行：latest_task_completed 的 task_key + title + occurred_at（本地时区，text-xs text-muted-foreground）
│   │   ├── 空态（current_phase_key 空串）：统一文案"暂无进行中 phase"（text-xs text-muted-foreground，不区分两情形）
│   │   └── latest_task_completed 无值时：附行"暂无任务完成记录"（text-xs text-muted-foreground）
│   │   └── 读取失败："进度摘要读取失败"（text-xs text-destructive）；加载态："加载中..."（text-xs text-muted-foreground）
│   ├── ProgressEventForm（border-t pt-3 分隔）
│   └── ProgressTimelineList（border-t pt-3 分隔）
```

内部子区域保持紧凑密度（text-xs），录入摩擦不增。

无路由声明（裁决⑥）：零 `routes/` 文件新增、零 `__root.tsx` NAV_ITEMS 改动、零 Dashboard 卡片改动——进度区唯一入口 = Repository detail 内嵌。

#### Scenario: 挂载实现判定

- **WHEN** 挂载落地
- **THEN** detail 页渲染顺序为 …grid 三列区 → StandardReadonlySummary → ProgressSection；全站 `grep` 无 `/progress` 路由与导航项；`tsc --noEmit` 零错误

### Requirement: 录入表单交互规格必须按冻结落地

本子件 SHALL 按 phase15-05 §"录入表单交互规格必须冻结"实现 `ProgressEventForm`：

**字段清单与默认值**：

| 字段 | 控件 | 必填性 | 默认值 / placeholder |
|---|---|---|---|
| `workflow_type` | select（phase / audit / fix） | 必填 | 默认 `phase` |
| `event_kind` | select（四值，非法组合 option disabled） | 必填 | 默认 = localStorage 记忆值（无记忆或非法时回退 `task_completed`） |
| `task_key` | text input | 按联动矩阵 | placeholder 按联动矩阵 |
| `title` | text input | 必填 | placeholder"一句话标题" |
| `detail` | textarea（rows=2） | 可选 | — |
| `evidence_ref` | text input | 可选 | placeholder"https://… 或 /仓库内路径" |
| `occurred_at` | datetime-local | 必填 | 默认浏览器本地当前时刻（DP-3） |
| `source` | **不暴露输入位**（请求不设置 source 字段 → 后端归一 manual） | — | — |

**`workflow_type × event_kind` 联动禁用矩阵**：

| workflow_type | event_kind 选项态 | 切换时重置规则 |
|---|---|---|
| `phase` | 四选项全部可选 | — |
| `audit` | `phase_started` / `phase_completed` **disabled**；`task_completed` / `note` 可选 | 当前 event_kind 变为非法时自动重置为 `task_completed` |
| `fix` | 同 `audit` | 同上 |

**`task_key` 联动必填与 placeholder 矩阵**：

| workflow_type × event_kind | task_key | placeholder |
|---|---|---|
| `phase` × `phase_started` / `phase_completed` | 必填 | `phaseNN（如 phase15）` |
| `phase` × `task_completed` | 必填 | `phaseNN-MM（如 phase15-05）` |
| `audit` × `task_completed` | 必填 | `audit_NNN（如 audit_001）` |
| `fix` × `task_completed` | 必填 | `fix_NNN（如 fix_001）` |
| 三轨 × `note` | 可选 | `自由标注，可留空` |

**`event_kind` 记住上次选择（localStorage）**：key `psco.progress.last-event-kind`；写入时机 = Create 成功后写入当次 `event_kind`；读取时机 = 表单挂载初始化 form state；非法值回退 `task_completed`；记忆值仅为表单默认值（纯 UX 偏好），不构成数据事实源。

**校验反馈双层模型**：前端轻量层行内即时提示（title 必填与上限 200 / detail 上限 2000 / evidence_ref 前缀 `/` 或 `https://` / task_key 按矩阵必填与格式）——仅 UX 提示不阻断输入；后端权威层错误经 `normalizeError` 行内回显（含 `REPOSITORY_NOT_FOUND` 等跨模块校验错误——前端不做 repository 存在性预校验）。

**提交回流与重置语义**：提交经 `use-create-progress-event`（owner 组装 pb 请求：枚举映射 + `occurred_at` 本地解析转 UTC Timestamp + 可空文本空串直传）；成功后失效矩阵生效（时间轴 + 当前卡刷新），表单重置 = `title` / `task_key` / `detail` / `evidence_ref` 清空 + `occurred_at` 重置为 now + `workflow_type` / `event_kind` **保持当前选择**（连续录入同轨同类事件是主场景）；失败停留表单上下文保留已填值，错误行内回显（`text-xs text-destructive`）；成功后不弹窗不跳转（内嵌区就地刷新）。

#### Scenario: 表单实现判定

- **WHEN** 表单落地
- **THEN** 联动禁用矩阵与 placeholder 矩阵逐格可验证；切换 audit/fix 后 `phase_started` 不可选且当前值自动重置；刷新页面后 `event_kind` 默认值 = 上次成功录入值；提交成功后表单按重置语义复位且时间轴顶部出现新事件
- **AND** `source` 无输入位；请求体不含 source 字段

### Requirement: 时间轴列表与误录删除交互规格必须按冻结落地

本子任务 SHALL 按 phase15-05 §"时间轴列表与误录删除交互规格必须冻结"实现 `ProgressTimelineList`：

**三轨过滤**：组件 state `'all' | 'phase' | 'audit' | 'fix'`，默认 `'all'`；四枚紧凑按钮组（全部 / Phase / Audit / Fix；`h-7 text-xs`；active 态高亮，inactive `ghost`）；切换即变更 query key 第三段自动重新查询；`'all'` 不传过滤参数；过滤切换 URL 不变（纯组件 state，无路由参数）。

**事件行结构（紧凑化规范，`divide-y` 列表，每行 `p-2 space-y-1 min-w-0`；列表容器 `max-h-80 overflow-y-auto` 限高滚动——2026-08-19 第三轮 UI 反馈修订：事件增多后垂直滚动预览，不再无限下撑页面）**：

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
第四行（可选，evidence_ref 非空时）——**2026-08-19 第三轮 UI 反馈修订：仅指明仓库目录大概位置**：`https://` 开头渲染为可点击外链（target=_blank rel=noreferrer）；`/` 开头仓库内路径前端无对应路由，渲染为纯文本标注（text-xs text-muted-foreground，不提供跳转）——导航引用零托管语义不变（裁决⑦）
```

`created_at` 不在行内展示；前端不重排序（后端三键链倒序直渲）、不过滤（过滤唯一经 query 参数）。

**空态文案**：全量空 `暂无推进事件，从上方表单录入第一条。`；过滤后空 `该轨暂无事件。`；读取失败 `进度事件读取失败`（text-xs text-destructive）；加载态 `加载中...`。

**误录删除确认**：事件行删除按钮 → `window.confirm`（沿 `standard-detail-page.tsx` 删除确认先例，不引入 Dialog 组件）；确认文案逐字：`确认删除事件「{title}」？删除仅用于修正误录，操作不可撤销。`；确认后经 `use-delete-progress-event`（参数仅事件 `id`）；成功后失效矩阵生效；失败行内错误回显；无软删除、无二次修正入口（append-only 语义，裁决⑨）。

#### Scenario: 时间轴实现判定

- **WHEN** 时间轴落地
- **THEN** 过滤切换触发重新查询且 URL 不变；删除确认文案逐字一致；删除成功后事件从列表消失且当前卡同步刷新
- **AND** 列表渲染顺序与 `ListProgressEvents` 响应顺序一致（无前端排序代码）

### Requirement: 移动端适配、运行时前提与工具链门禁

本子任务 SHALL 满足以下基线与门禁：

**移动端适配基线**（对齐 Repository detail 既有基线与紧凑化规范）：
- 表单字段（2026-08-19 第二轮 UI 反馈修订，短值紧凑化）：`grid grid-cols-2 gap-2 lg:grid-cols-4`——PC 四列（workflow_type / event_kind / task_key / occurred_at 四短字段一行；title 与 evidence_ref 各占 2 列并排一行；detail 全宽占 4 列）；移动端两列（四短字段两两一行，title / evidence_ref / detail 全宽）；提交按钮 `h-8`
- 时间轴事件行：`flex flex-wrap` + `min-w-0` + `truncate`（防横向溢出）；过滤按钮组 `flex flex-wrap`
- 当前卡：单列紧凑（text-xs 密度）
- 区块标题与子区域分隔沿紧凑化规范（`text-xs` / `border-t pt-2`）；容器型元素仅 `:focus-visible`

**运行时前提（phase15-06 留档观察 3 承接）**：
- 当前运行中的后端服务器（用户手动启动，早于 phase15-06 代码）不含 progress RPC——浏览器完整会话验证前**由用户重启后端服务器**（重启即载入 phase15-06 代码 + `RunMigrations` 空过补登记 0013，DDL 幂等无风险）
- 前端 vite 热更新自动生效（新切片文件与挂载改动无需重启前端）
- 本子任务代理不得擅自重启任何服务器；实现与静态门禁（tsc）先行，浏览器会话验证在用户重启后端后执行

**门禁（DoD 冻结）**：
- `frontend/` 目录 `tsc --noEmit` 零错误（实现完成即验证）
- 浏览器完整会话 DoD（用户重启后端后执行）：录入 `phase_started` → 录入多条 `task_completed`（含补录历史 occurred_at）→ 录入 `note` → 派生当前卡正确 → 删除误录 → `phase_completed` 后当前卡转完结态（空态统一文案 + 时间轴最新事件行呈现 phase_completed）→ audit/fix 轨录入与过滤

**不做（边界冻结）**：无 dogfooding 固定录入集与验收报告（phase15-08）、无根级文档回写（phase15-09）、无 shared/ 晋升、无第二套路由 / 状态管理 / UI 框架（project_rules §2.4）。

#### Scenario: 门禁全绿收口

- **WHEN** 全部实现与验证完成
- **THEN** `tsc --noEmit` 零错误；浏览器完整会话 DoD 七步全部可完成；git 工作区改动仅含本 spec §What Changes 文件清单（+ 本 spec 三件套目录）；变更保持未提交，待用户最终确认后手动提交

## 与上游单值一致声明

- 切片结构 / DP-1 / DP-3 / 组件树 / 表单与时间轴交互 / 移动端基线：全部逐字源 = phase15-05 对应 Requirement
- 联动禁用矩阵 ↔ phase15-02 合法矩阵 12 格（audit/fix 禁 phase 边界标记 = 规则 7；note 可选 task_key = 规则 8）；task_key placeholder 矩阵 ↔ K-1~K-4 正则（提示语与格式一一对应，不收窄不放宽）
- 错误回显 ↔ phase15-03 错误码总表 + phase15-04 3 RPC 错误语义（前端按归一化 message 回显，不自行翻译错误码）
- 切片纪律 ↔ project_rules §2.5（query 纯只读 / mutation 固定承接位 / 切片优先不晋升 shared）
- 本 spec 不承载任何新设计决策；实现中发现上游文档与既有代码冲突时，以上游 spec 冻结口径为准并回改实现，不得反向改写上游
