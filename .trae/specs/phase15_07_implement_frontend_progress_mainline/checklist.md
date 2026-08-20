# phase15-07 Checklist

## 实现与设计一致性

- [x] 切片 11 文件清单与 phase15-05 §"切片结构必须冻结"一一对应（data 3 / application 2 / components 4 / types.ts / index.ts；无第五目录）
- [x] `types.ts`：string union 三枚举 + `ProgressEvent` 11 字段（snake_case 对齐后端）+ `ProgressSummary` 四字段 + pb 转换（`pbToProgressEvent` / 枚举映射 / datetime-local ↔ pb Timestamp）+ K-1~K-4 正则常量逐字（与 phase15-02 L43-46 核对一致）+ 合法组合判定辅助
- [x] `connect-client.ts` 沿 standard 模式（`progressClient` + `projectContextClient`；shared transport；页面组件不直接 createClient）
- [x] `use-progress-events-read.ts`：query key `['progress-events', repositoryId, filter]`；filter 'all' 不传过滤 / 三轨传枚举；前端不重排序
- [x] `use-repository-progress-read.ts`：query key `['repository-progress', repositoryId]`；brief 投影唯一通道（DP-1）；空态恒有值（undefined → 零值摘要防御）
- [x] application 2 文件失效矩阵逐字（`['progress-events', repositoryId]` 前缀 + `['repository-progress', repositoryId]`）；`normalizeError` 收敛 owner 内（与 use-create-standard 逐字同型）；Create 请求不设置 source 字段
- [x] 切片纪律（grep 实测）：`useMutation` 仅命中 application 2 文件；`ListProgressEvents` 仅命中 data owner（时间轴经 hook 间接消费，无派生计算路径）；barrel 仅导出 `ProgressSection`

## 组件与交互规格

- [x] `ProgressCurrentPhaseCard`：进行中态三要素（key code/mono + label + latest_task_completed 行本地时区）；空态统一文案"暂无进行中 phase"（不区分两情形）；"暂无任务完成记录"附行；失败/加载态文案；组件内零派生逻辑（独立复核 grep 确认无 ListProgressEvents 引用）
- [x] `ProgressEventForm`：字段清单与默认值逐字（source 无输入位）；联动禁用矩阵（audit/fix × phase_started/phase_completed disabled + 非法自动重置 task_completed——浏览器 S4 实证 aria-disabled×2 与切轨自动重置）；task_key placeholder 矩阵五格逐字（S4 getAttribute 实测）；event_kind localStorage 记忆（key `psco.progress.last-event-kind` / 成功写入 / 非法回退——浏览器 S1 实证记忆默认值）
- [x] 表单提交回流：成功后 title/task_key/detail/evidence_ref 清空 + occurred_at 重置 now + workflow_type/event_kind 保持（S1 实证 title 重置）；失败保留已填值行内回显；不弹窗不跳转；前端轻量提示不阻断输入
- [x] `ProgressTimelineList`：三轨过滤（默认 'all'，四枚按钮组 h-7 text-xs，URL 不变——S6 四向实证）；事件行四行结构逐字（Badge variant 按 workflow_type / task_key 未填不渲染——S1 实证 / created_at 不展示）；**evidence_ref 双形态（2026-08-19 第三轮修订）：https:// 可点击外链（target=_blank rel=noreferrer），/ 开头仓库内路径纯文本标注不跳转（前端无对应路由，只读验证 4 条全 <p> 渲染零 <a> 残留）**；**列表容器 max-h-80 overflow-y-auto 限高滚动（第三轮修订，只读验证类名在位）**；空态文案四则逐字（S7 实证全量空文案）
- [x] 删除确认：`window.confirm` 逐字文案（S2 浏览器 dialog.message() 实证逐字一致）；成功后时间轴 + 当前卡同步刷新（S2 实证 latest 由 phase15-02 变 phase15-01）；无软删除无二次修正入口
- [x] `ProgressSection` 外壳（**Card 化 + 移入工作台右列，2026-08-19 用户两轮 UI 反馈裁决**：Card min-w-0 + CardTitle"项目进度" text-base font-semibold；挂载于绑定工作台右列"相关决策"之后为第四卡片，与主体工作台卡片风格一致；内部 border-t pt-3 分隔）；页面渲染顺序：grid 三列区（右列 已绑定产品 → 已映射模块 → 相关决策 → 项目进度）→ StandardReadonlySummary（页面底部注释风格收尾，保持 phase14 先例）
- [x] 移动端基线（2026-08-19 第二轮修订）：表单 `grid grid-cols-2 gap-2 lg:grid-cols-4`（PC 四短字段一行 + title/evidence_ref 并排 + detail 全宽）+ 提交按钮 h-8；事件行 flex flex-wrap + min-w-0 + truncate；当前卡 text-xs 紧凑

## 边界与纪律

- [x] 零路由 / 零导航 / 零 Dashboard 改动（裁决⑥；全站 grep 无 `/progress` 路由与导航项——routes 内 "progress" 仅命中无关的 first_run_state 枚举）
- [x] 零 shared/ 晋升（切片优先）；零第二套 UI 框架 / 路由 / 状态管理引入
- [x] 零后端 / proto / 迁移 / 根级文档改动（git status 仅 11 新文件 + 1 修改 + spec 三件套目录）
- [x] 运行时约束遵守：代理未擅自重启任何服务器；浏览器验证前经用户重启后端（载入 phase15-06 代码；RunMigrations 补登记 0013）

## 门禁与验证

- [x] `frontend/` 目录 `tsc --noEmit` 零错误（实现完成时与 Task 5/6 收口三次复验，均退出码 0）
- [x] 浏览器完整会话 DoD 七步全部完成（真实浏览器驱动，33 项断言全 PASS）：①录入 phase_started → ②多条 task_completed（含补录历史 occurred_at，展示倒序实证）→ ③录入 note → ④派生当前卡正确（key/label/latest 三要素）→ ⑤删除误录（confirm 逐字 + 同步刷新）→ ⑥phase_completed 后当前卡转完结态（空态统一文案 + brief currentPhaseKey 派生空串 DoD 冻结断言 + 时间轴最新行呈现 phase_completed）→ ⑦audit/fix 轨录入与过滤（联动禁用 aria-disabled 实证 + 四向过滤 + URL 不变 + DOM 序=API 三键链序）
- [x] 时区闭环抽查：录入 07:31~07:33 UTC（补录）→ 展示本地 15:3x（DP-3 无漂移实证）；补录历史与实时录入同一转换路径
- [x] 浏览器验证结束清理测试数据（progress_events 恢复 0 行 dogfooding 干净态，psql 实证；phase15-08 前提）
- [x] git 工作区改动与 spec §What Changes 文件清单一一对应（+ 本 spec 三件套目录）
- [x] 独立复核通过（0 阻断：七维度全 PASS——切片结构/交互规格/纪律 grep/DP-1·DP-3/门禁实测/留证核对/无偷渡；非阻断观察 4 项留档）
- [x] tasks.md / checklist.md 全部勾选并附执行记录（见 tasks.md §执行记录）
- [x] 变更未提交，待用户最终确认后手动提交

## 附：验证方式说明（留档）

- 本机无 Chrome 且 Chrome DevTools MCP 无法指定 executablePath（仅查找 /opt/google/chrome/chrome，无 sudo 不可创建）；改用 playwright-core + 本地已有 chromium 二进制（~/.cache/ms-playwright/chromium-1234/chrome-linux64/chrome，HOME 重定向 /tmp 规避沙箱写限制）驱动真实浏览器执行完整会话，驱动脚本与成功截图留档 /tmp/psco-browser-test/（drive.mjs + final-empty.png）。
- 过程迭代修正的 3 处问题均为驱动脚本定位误差（非产品缺陷）：section 首个匹配误中 Standard 摘要区、:has-text 子串误中表单 Select trigger、"全部"期望序未按同分钟 created_at DESC tiebreak 推算；修正后产品代码零改动全量通过。
