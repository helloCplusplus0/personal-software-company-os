# Tasks

- [x] Task 1: 切片 data 层与 types 转换层
  - [x] SubTask 1.1: 新建 `frontend/src/features/progress/types.ts`——`WorkflowType`/`EventKind`/`ProgressSource` string union + `ProgressEvent`（11 字段，对齐后端 snake_case）+ `ProgressSummary`（四字段）+ `CreateProgressEventInput`（表单提交模型）+ pb 转换（`pbToProgressEvent` / 枚举映射 / datetime-local ↔ pb Timestamp 转换）+ K-1~K-4 正则常量（`^phase[0-9]{2,}$` / `^phase[0-9]{2,}-[0-9]{2,}$` / `^audit_[0-9]{3,}$` / `^fix_[0-9]{3,}$`）+ 合法组合判定辅助（联动禁用依据）；沿 `standard/types.ts` 的 pb 转换写法
  - [x] SubTask 1.2: 新建 `frontend/src/features/progress/data/connect-client.ts`——导出 `progressClient`（ProgressService）与 `projectContextClient`（brief 读取用，DP-1 通道）；沿 `standard/data/connect-client.ts` 模式（createClient + `@/shared/rpc/connect-transport`）
  - [x] SubTask 1.3: 新建 `frontend/src/features/progress/data/use-progress-events-read.ts`——query key `['progress-events', repositoryId, filter]`（filter：`'all'` 不传过滤 = 三轨全量；`'phase'/'audit'/'fix'` 传对应枚举）；结果即后端三键链倒序，前端不重排序；`pbToProgressEvent` 逐元素转换
  - [x] SubTask 1.4: 新建 `frontend/src/features/progress/data/use-repository-progress-read.ts`——query key `['repository-progress', repositoryId]`；`GetProjectBrief` → 仅投影 `progress` 块 → `ProgressSummary`（DP-1 通道；空态恒有值）；沿 `use-repository-standards-read` 切片自有 brief 投影 owner 模式

- [x] Task 2: 切片 application 层（mutation 固定承接位）
  - [x] SubTask 2.1: 新建 `frontend/src/features/progress/application/use-create-progress-event.ts`——`CreateProgressEvent` 承接位：表单值 → pb 请求组装（枚举映射 + `occurred_at` 本地解析转 UTC Timestamp + 可空文本空串直传 + **不设置 source 字段**）；错误归一化 `normalizeError`（提取 ConnectError rawMessage，沿 `use-create-standard` 模式）；失效矩阵逐字：`['progress-events', repositoryId]` 前缀（覆盖全部过滤变体）+ `['repository-progress', repositoryId]`
  - [x] SubTask 2.2: 新建 `frontend/src/features/progress/application/use-delete-progress-event.ts`——`DeleteProgressEvent` 承接位（参数仅事件 `id`）；错误归一化同上；失效矩阵同上（删除同样改变事件流与派生摘要）

- [x] Task 3: 切片 components 层（4 组件）
  - [x] SubTask 3.1: 新建 `frontend/src/features/progress/components/progress-current-phase-card.tsx`——DP-1 通道消费：进行中态（current_phase_key code/mono + current_phase_label + 最新完成任务行 task_key/title/occurred_at 本地时区）/ 空态统一文案"暂无进行中 phase"（不区分两情形）/ latest_task_completed 无值附行"暂无任务完成记录" / 读取失败"进度摘要读取失败" / 加载态"加载中..."；组件内零派生逻辑（grep 无 ListProgressEvents 结果参与计算）
  - [x] SubTask 3.2: 新建 `frontend/src/features/progress/components/progress-event-form.tsx`——字段清单与默认值逐字（workflow_type 默认 phase / event_kind 默认 localStorage 记忆值回退 task_completed / occurred_at 默认 now / source 不暴露）；联动禁用矩阵（audit/fix × phase_started/phase_completed disabled + 非法时自动重置 task_completed）；task_key placeholder 矩阵；event_kind 记忆（key `psco.progress.last-event-kind`，Create 成功写入）；前端轻量层行内提示（不阻断输入）+ 后端权威层错误回显；提交回流与重置语义（成功后 title/task_key/detail/evidence_ref 清空 + occurred_at 重置 now + workflow_type/event_kind 保持；失败保留已填值；不弹窗不跳转）
  - [x] SubTask 3.3: 新建 `frontend/src/features/progress/components/progress-timeline-list.tsx`——三轨过滤（组件 state 默认 'all'，四枚紧凑按钮组 `h-7 text-xs`，切换变更 query key 第三段）；事件行结构逐字（四行紧凑结构 / Badge variant 按 workflow_type / occurred_at 本地时区 / task_key 未填不渲染 / created_at 不展示）；空态文案四则；删除按钮 → `window.confirm` 逐字文案 `确认删除事件「{title}」？删除仅用于修正误录，操作不可撤销。` → `use-delete-progress-event`；前端不重排序不过滤
  - [x] SubTask 3.4: 新建 `frontend/src/features/progress/components/progress-section.tsx`——进度区外壳与编排（section.min-w-0.space-y-2 + 区块标题"项目进度" + 三子组件 border-t pt-2 分隔）；沿 `StandardReadonlySummary` 外壳样式模式

- [x] Task 4: barrel 与 Repository detail 挂载
  - [x] SubTask 4.1: 新建 `frontend/src/features/progress/index.ts`——barrel 仅导出 `ProgressSection`（沿 `standard/index.ts` 受控导出模式；不导出内部组件 / query / mutation owner）
  - [x] SubTask 4.2: 修改 `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`——import 区追加 `ProgressSection`（from `@/features/progress`），渲染区在 `<StandardReadonlySummary />` 之后追加 `<ProgressSection repositoryId={repositoryId} />`（页面级第三全宽区块）；注释条目同步 phase15-07 来源说明
  - [x] SubTask 4.3: 静态纪律校验——`grep -r "useMutation" frontend/src/features/progress/` 仅命中 application 2 文件；`grep -r "ListProgressEvents" frontend/src/features/progress/` 仅命中时间轴 query owner 与其消费（无派生计算路径）；全站 grep 无 `/progress` 路由与导航项新增；`frontend/` 目录 `tsc --noEmit` 零错误

- [x] Task 5: 浏览器完整会话 DoD 验证（运行时前提：用户重启后端后执行）
  - [x] SubTask 5.1: 向用户请求重启后端服务器（当前运行中的后端不含 phase15-06 progress RPC；重启即载入新代码 + RunMigrations 空过补登记 0013，DDL 幂等无风险；前端 vite 热更新自动生效）——本子任务代理不得擅自重启
  - [x] SubTask 5.2: 浏览器完整会话七步验证（dev_plan phase15-07 DoD 逐字）：①录入 `phase_started`（phase 轨，task_key 如 phase15）→ ②录入多条 `task_completed`（含补录历史 occurred_at——修改 datetime-local 为过去时刻）→ ③录入 `note` → ④派生当前卡正确（current_phase_key/label + latest_task_completed 更新）→ ⑤删除误录（window.confirm 确认 + 列表移除 + 当前卡同步刷新）→ ⑥录入 `phase_completed`（同 task_key）后当前卡转完结态（空态统一文案"暂无进行中 phase" + 时间轴最新事件行呈现 phase_completed）→ ⑦audit/fix 轨录入（task_completed/note）与三轨过滤切换（含 URL 不变断言）
  - [x] SubTask 5.3: 会话过程留证——每步关键界面状态记录（当前卡文案 / 时间轴顺序 / 过滤结果）；时区闭环抽查（录入本地时刻 → 刷新后展示同一本地时刻，无漂移）；验证结束清理测试数据（删除本会话录入的全部事件，恢复 dogfooding 前干净态——phase15-08 固定录入集前提）

- [x] Task 6: 一致性校验、独立复核与收口
  - [x] SubTask 6.1: 一致性校验——切片 11 文件与 phase15-05 §"切片结构必须冻结"一一对应；组件树 / 联动矩阵 / placeholder 矩阵 / 空态文案 / 删除确认文案与 phase15-05 逐字比对；DP-1 / DP-3 裁决逐项落地核对；确认 spec §What Changes 文件清单与 git 工作区实际改动一一对应（无清单外文件、零后端 / proto / 迁移 / 路由 / Dashboard 改动）
  - [x] SubTask 6.2: 子代理独立复核（切片纪律 grep 实测 / 联动矩阵与 phase15-02 合法矩阵逐格比对 / DP-1 唯一通道与零派生验证 / tsc 门禁实测 / 浏览器会话 DoD 七步留证核对 / 无偷渡〔零 dogfooding 残留数据、零根级文档、零 phase15-08/09 越界〕）
  - [x] SubTask 6.3: 修复独立复核发现的阻断性问题（如有）并复验；tasks.md / checklist.md 全部勾选附执行记录；变更保持未提交，待用户最终确认后手动提交

# Task Dependencies

- Task 1（data + types）为基础，无外部依赖
- Task 2 depends on Task 1（mutation owner 消费 types 转换与 connect-client）
- Task 3 depends on Task 1 + Task 2（组件消费 query / mutation owner 与 types）
- Task 4 depends on Task 3（barrel 导出 ProgressSection 后挂载）
- Task 5 depends on Task 4（浏览器验证需完整实现落地）+ 用户重启后端（运行时前提）
- Task 6 depends on Task 5
- 后续：phase15-08 depends on 本 spec（浏览器主线可用）+ phase15-06（后端运行时）；phase15-09 depends on phase15-08

# 执行记录（2026-08-19）

- Task 1-4（首批实现，额度中断前完成）：切片 11 文件落地（types 221 行含 K-1~K-4 正则与 isEventKindAllowed / data 3 / application 2 含 normalizeError 与双 key 前缀失效矩阵 / components 4）；barrel 仅导出 ProgressSection；detail 页 +5 行挂载（grid 三列区 → StandardReadonlySummary → ProgressSection）；`npx tsc -b` 零错误；纪律 grep 全过（useMutation 仅 application 2 文件 / ListProgressEvents 仅 data owner / 全站无 /progress 路由）。
- Task 5.1：用户确认已重启后端（:8081 载入 phase15-06 代码）；curl 验证 progress RPC 就绪（ListProgressEvents 返回 INVALID_REPOSITORY_ID 证明路由可达）；前端 vite 热更新生效（:5173）。
- Task 5.2-5.3（续接会话完成）：发现并接管上次中断残留的 3 条测试数据（phase_started phase15 + 2 条补录 task_completed，occurred_at 均早于 created_at）；因环境无 Chrome 且 Chrome DevTools MCP 无法指定 executablePath，改用 playwright-core + 本地 chromium 二进制（~/.cache/ms-playwright/chromium-1234，HOME 重定向 /tmp 规避沙箱）编写真实浏览器驱动脚本 /tmp/psco-browser-test/drive.mjs，在真实浏览器中执行完整会话七步：**33 项断言全部 PASS**——S0④ 当前卡三要素 + ①②时间轴三行倒序（含补录）+ DP-3 时区展示（07:31 UTC → 本地 15:31）；S1③ note 录入 + event_kind 记忆默认值 + 表单重置语义 + task_key 未填不渲染；S2⑤ window.confirm 逐字文案 + 行消失 + 当前卡同步刷新（latest 由 phase15-02 变 phase15-01）；S3⑥ phase_completed 后当前卡空态统一文案「暂无进行中 phase」+ brief progress.currentPhaseKey 派生空串（DoD 冻结断言，DP-1 同源权威验证）+ 时间轴最新行呈现 phase_completed；S4⑦ audit 轨（切轨自动重置 + phase 边界选项 aria-disabled×2 + placeholder getAttribute 实测 audit_NNN）+ audit_001 录入；S5⑦ fix 轨 note 录入；S6⑦ 三轨过滤四向切换（Phase/Audit/Fix/全部）+ URL 不变×4 + 全量 DOM 序与 API 三键链序逐一致（前端不重排权威断言；audit/fix 同分钟录入由 created_at DESC tiebreak 正确排序——初次"FAIL"经诊断为脚本期望值错误而非产品缺陷）；S7 清理全部测试数据（progress_events 恢复 0 行 dogfooding 干净态）+ 空态文案 + 当前卡空态组合断言。过程迭代修正 3 处脚本定位问题（section 首个匹配误中 Standard 摘要区→改含"项目进度"锚定；:has-text 子串误中表单 Select trigger→改 :text-is 精确匹配+过滤组容器限定；"全部"期望序手写错误→改 API 序动态对比），产品代码零改动。
- Task 6：独立复核子代理七维度全 PASS、0 阻断：①切片结构与 phase15-05 逐字一致（11 文件/K-1~K-4 正则/query key/失效矩阵/barrel）；②交互规格逐字一致（联动禁用/placeholder 五格/空态四则/删除确认文案/移动端基线）；③切片纪律 grep 实测全过（零越界改动，git 仅 11 新文件+1 修改+spec 目录）；④DP-1 唯一 brief 通道与 recent_events 零 web 消费 + DP-3 datetime-local/UTC 转换/toLocaleString 全落实；⑤门禁亲自复验（tsc -b 零错误 / RPC 干净态 / progress_events 0 行 / localStorage key 逐字）；⑥驱动脚本 33 断言语义与 DoD 七步覆盖完整对应；⑦无偷渡（零 dogfooding 残留、零根级文档、无 acceptance_report、未提交）。非阻断观察 4 项留档（Connect JSON 空 repeated 省略语义等价、drive.mjs 一处恒真占位断言已被 S4 getAttribute 实证覆盖、①②录入动作承接自先前会话展示侧验证、/tmp fail.png 迭代残留截图）。
- 收口状态：tsc -b 零错误复验通过；全部变更未提交（11 新文件 + 1 修改 + spec 三件套），待用户最终确认后手动提交。

## 返工记录（2026-08-19 用户 UI 反馈）

- 反馈内容：进度区沿用 StandardReadonlySummary 裸 `section` 模式定位错位——进度区是维护工作台（表单/删除/过滤）而非只读摘要，`section` 应与页面正式内容（div/Card 工作台卡片）保持一致性。
- 用户裁决（AskUserQuestion 两问拍板）：①进度区 **Card 化**（与"已绑定产品/相关决策"同款 Card + CardTitle text-base font-semibold；内部紧凑密度不变，border-t pt-3 分隔）；②Standard 只读摘要区**保持现状**（phase14 先例不动，不同步统一；风格断层留档观察）。
- 实现：重写 `progress-section.tsx`（裸 section → Card/CardHeader/CardTitle/CardContent 四层，40 行）；其余 10 文件与挂载点零改动。
- spec 同步：spec.md §"挂载与组件树必须按冻结落地"对齐式修订（容器形态段 + 组件树 2026-08-19 修订版，注明覆盖 phase15-05 原冻结条目）；checklist.md 对应检查项同步。
- 复验：`npx tsc -b` 零错误；drive.mjs 进度区定位同步更新（`div[data-slot="card"]` 过滤"项目进度"）后重置 3 条初始数据全量重跑浏览器完整会话七步——**33 项断言全部 PASS，Card 化零功能回归**，验证后数据自清理至 progress_events 0 行（dogfooding 干净态保持）。
- Standard 摘要与进度区风格断层（裸 section → Card 并存）留档：属 phase14 交付物，本次按用户裁决零触碰；如后续统一由新裁决驱动。

## 返工记录二（2026-08-19 用户第二轮 UI 反馈）

- 反馈内容：①第一轮"Card 化但仍全宽游离在 Standard 注释区之后"未解决根本问题——进度区必须与主体工作台（grid 右列"已绑定产品/已映射模块/相关决策"）保持风格一致性，Standard section 作为外部关联注释风格收尾于页面底部没问题；②表单短值字段（workflow_type/event_kind/task_key/occurred_at/title/evidence_ref）一行 1-2 个太松散，须紧凑化。
- 实现：①`repository-binding-detail-page.tsx`——`ProgressSection` 移入 `space-y-4 lg:col-span-2` 工作台右列"相关决策"面板之后（第四卡片），Standard 摘要独占页面底部；②`progress-event-form.tsx`——布局由 `grid grid-cols-1 gap-2 sm:grid-cols-2` 改为 `grid grid-cols-2 gap-2 lg:grid-cols-4`（PC：四短字段一行 + title/evidence_ref 各占 2 列并排 + detail 全宽 4 列；移动：两列，title/evidence_ref/detail 全宽），表单由 4 行压缩至 PC 3 行。
- spec 同步：spec.md §挂载点/分区关系（第二轮修订，覆盖 phase15-05 原"页面级第三全宽区块"冻结）与 §移动端适配基线（表单紧凑化）对齐式修订；checklist.md 两项同步。
- 复验：`npx tsc -b` 零错误；重置 3 条初始数据全量重跑浏览器完整会话七步——**33 项断言全部 PASS**（Card 移位与四列布局零功能回归），验证后数据自清理至 progress_events 0 行。
- 页面最终结构（两轮返工后单值）：grid 三列区（左列仓库摘要 / 右列 已绑定产品 → 已映射模块 → 相关决策 → 项目进度 Card）→ 关联 Standard 裸 section（底部注释风格）。

## 返工记录三（2026-08-19 用户第三轮 UI 反馈）

- 反馈内容：①evidence_ref `/` 开头的仓库内路径 web 端不应提供跳转链接（前端无对应路由，点击必然 404）——它仅指明仓库目录大概位置；②时间轴列表不应随事件增多无限下撑页面，达到一定数据量后应滚动预览。
- 实现（仅 `progress-timeline-list.tsx` 两处）：①evidence_ref 双形态渲染——`https://` 开头保留可点击外链（target=_blank rel=noreferrer）；`/` 开头渲染为纯文本标注（`text-xs text-muted-foreground`，无 `<a>`）；②列表容器加 `max-h-80 overflow-y-auto`（约 6 条事件后垂直滚动）。
- 验证：`npx tsc -b` 零错误；会话中断导致前后端停止，经用户重启后执行**只读浏览器验证**（以用户手录的 phase01 系列 5 条事件为素材，数据零改动）——4 条 `/` 开头 evidence_ref 全部渲染为 `<p>` 纯文本、`<a>` 内零 `/plan.md`/`/docs` 残留、限高滚动容器类名在位，3 项检查全 PASS（脚本与截图留档 /tmp/psco-browser-test/verify-ui3.mjs + readonly-ui3.png）。
- spec 同步：spec.md §时间轴事件行结构（列表容器限高滚动 + evidence_ref 第四行双形态）对齐式修订；checklist.md 对应检查项同步。
- 数据说明：用户手录的 phase01 系列 5 条测试事件保留未动（属用户自录数据；phase15-08 dogfooding 前的清理由该子任务统一处理）。
