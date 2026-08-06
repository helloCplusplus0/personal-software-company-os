# Tasks

- [x] Task 1: 收口 `Decision Center` 前端文件落点、路由树与主导航接入。
  - [x] SubTask 1.1: 明确 `frontend/src/routes/decisions/index.tsx`、`new.tsx`、`$decisionId.tsx` 的文件落点与 URL 语义
  - [x] SubTask 1.2: 明确 `frontend/src/features/decision-center/pages/` 下三个页面文件的落点
  - [x] SubTask 1.3: 明确 `frontend/src/routes/__root.tsx` 必须增加 `Decision Center` 主导航入口

- [x] Task 2: 收口 `Decision Center` 前端数据适配边界与类型语义。
  - [x] SubTask 2.1: 明确 `types.ts` 必须承接列表、详情、创建、候选、关联与 `source_context` 字段语义
  - [x] SubTask 2.2: 明确 `data/api-adapter.ts` 必须直接消费 `phase03-12` 真实 API
  - [x] SubTask 2.3: 明确 `data/decision-center-adapter.ts` 不提供并列 mock 主线，只导出真实 API 适配实现

- [x] Task 3: 收口 `Decision List` 前端主线实现要求。
  - [x] SubTask 3.1: 明确列表路由必须使用 `createFileRoute + validateSearch` 承接 `queryText / statusFilter`
  - [x] SubTask 3.2: 明确列表读取、空状态、错误态与进入创建/详情的最小用户可见行为
  - [x] SubTask 3.3: 明确列表上下文恢复继续沿用 `Zustand + sessionStorage` 的"最后一次搜索"缓存模式，不形成第二套事实源

- [x] Task 4: 收口 `Decision Create` 前端主线实现要求。
  - [x] SubTask 4.1: 明确 `DecisionCreateRoute` 必须承接 `sourceModuleId / sourceModuleName` 搜索参数
  - [x] SubTask 4.2: 明确 `DecisionContextSourcePanel`、`DecisionCreateForm`、`DecisionCreateActions` 的最小职责
  - [x] SubTask 4.3: 明确创建成功回流到 `DecisionDetailPage`、失败保留草稿与来源上下文

- [x] Task 5: 收口 `Decision Detail` 前端主线实现要求。
  - [x] SubTask 5.1: 明确详情页必须统一承接详情读取、已关联目标、待关联目标、候选读取与最小关联动作
  - [x] SubTask 5.2: 明确待关联目标必须由 `source_context + linked_modules` 派生，不得靠临时本地假状态兜底
  - [x] SubTask 5.3: 明确关联成功后必须失效并重读详情、候选与列表相关查询

- [x] Task 6: 收口 `Module Detail` 到 `Decision Center` 的入口升级要求。
  - [x] SubTask 6.1: 明确 `ModuleDecisionEntryPanel` 必须提供"为当前 Module 记录决策"与"查看当前 Module 相关决策"两个正式动作
  - [x] SubTask 6.2: 明确面板内已有决策项应可直接进入 `DecisionDetailPage`
  - [x] SubTask 6.3: 明确不得在 `Module Detail` 侧新增中间路由或中间分发组件

- [x] Task 7: 收口前端运行时主线与刷新策略。
  - [x] SubTask 7.1: 明确页面读取必须采用 `TanStack Query useQuery`
  - [x] SubTask 7.2: 明确创建与关联写入必须采用 `useMutation + onSuccess invalidateQueries`
  - [x] SubTask 7.3: 明确不得在 mutation 成功后用手工拼接页面假数据替代正式 reread

- [x] Task 8: 收口 `PC / 移动浏览器` 双场景布局实现要求。
  - [x] SubTask 8.1: 明确 `DecisionListPage` 的桌面高密度布局与移动单列重排
  - [x] SubTask 8.2: 明确 `DecisionCreatePage` 的来源上下文区、表单区、动作区在两类场景下的布局关系
  - [x] SubTask 8.3: 明确 `DecisionDetailPage` 的桌面分区布局与移动垂直重排，不引入第二套移动端 UI 架构

- [x] Task 9: 收口实现验证与验收证据要求。
  - [x] SubTask 9.1: 明确至少需要 `npm run build` 通过，确保路由与页面主线可编译
  - [x] SubTask 9.2: 明确需要验证 `Decision List -> Decision Create -> Decision Detail -> LinkDecisionToTarget` 最小前端闭环
  - [x] SubTask 9.3: 明确需要验证从 `Module Detail` 发起创建时，`source_context` 能贯通到 `DecisionDetailPage`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 3`, `Task 4`, and `Task 5`
- `Task 8` depends on `Task 3`, `Task 4`, and `Task 5`
- `Task 9` depends on `Task 3`, `Task 4`, `Task 5`, `Task 6`, `Task 7`, and `Task 8`

# 执行证据

## 编译与静态检查
- `tsc -b` 通过（无类型错误）
- `vite build` 通过（2162 modules transformed，dist 产物正常）
- `oxlint` 通过（0 errors，4 warnings 均为预先存在的 shadcn/ui 组件导出 warning）

## 路由树生成
- `tanstackRouter` Vite 插件自动扫描 `src/routes/decisions/` 并生成 routeTree.gen.ts
- 3 个新路由文件（index.tsx / new.tsx / $decisionId.tsx）已注册到路由树
- `__root.tsx` 导航已增加 Decision Center 入口

## 文件落点验证
- `frontend/src/routes/decisions/index.tsx` — DecisionListRoute（validateSearch 承接 queryText / statusFilter）
- `frontend/src/routes/decisions/new.tsx` — DecisionCreateRoute（validateSearch 承接 sourceModuleId / sourceModuleName）
- `frontend/src/routes/decisions/$decisionId.tsx` — DecisionDetailRoute
- `frontend/src/features/decision-center/types.ts` — DTO / 枚举 / 请求响应类型
- `frontend/src/features/decision-center/data/api-adapter.ts` — 真实 API 适配层（5 个函数）
- `frontend/src/features/decision-center/data/decision-center-adapter.ts` — 统一导出（不提供 mock）
- `frontend/src/features/decision-center/stores/decision-list-search-store.ts` — Zustand + sessionStorage
- `frontend/src/features/decision-center/pages/decision-list-page.tsx` — 列表页
- `frontend/src/features/decision-center/pages/decision-create-page.tsx` — 创建页
- `frontend/src/features/decision-center/pages/decision-detail-page.tsx` — 详情页
- `frontend/src/features/decision-center/components/decision-list-toolbar.tsx` — 搜索工具栏
- `frontend/src/features/decision-center/components/decision-list-content.tsx` — 列表内容（PC 表格 / 移动卡片）
- `frontend/src/features/decision-center/components/decision-context-source-panel.tsx` — 来源上下文展示
- `frontend/src/features/decision-center/components/decision-create-form.tsx` — 结构化模板表单
- `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx` — 概要卡片
- `frontend/src/features/decision-center/components/decision-linked-targets-section.tsx` — 已关联目标区
- `frontend/src/features/decision-center/components/decision-pending-link-target-card.tsx` — 待关联目标卡片
- `frontend/src/features/decision-center/components/decision-module-candidate-panel.tsx` — 候选读取与关联面板
- `frontend/src/routes/__root.tsx` — 增加 Decision Center 导航入口
- `frontend/src/features/module-registry/components/module-decision-entry-panel.tsx` — 升级为正式入口触点
- `frontend/src/features/module-registry/pages/module-detail-page.tsx` — 传入 moduleId / moduleName

## 前端闭环逻辑验证（代码审查）
- Decision List → Decision Create：列表页“记录决策”按钮导航到 `/decisions/new`，显式带 `fromList: true` ✅
- Decision Create → Decision Detail：创建成功后 `navigate` 回流详情，透传 `fromList: search.fromList` ✅
- Decision Detail → LinkDecisionToTarget：候选面板 `useMutation` 触发关联，成功后 `invalidateQueries` ✅
- Module Detail → Decision Create（带 source_context）：`ModuleDecisionEntryPanel` 导航到 `/decisions/new?sourceModuleId=...&sourceModuleName=...`，不带 `fromList`（返回列表落默认参数）✅
- source_context 贯通到 DecisionDetailPage：创建页通过 `source_module_id` 持久化到后端，详情页通过 `source_context` 读取并展示待关联目标 ✅

## 复核修复（GPT-5.4 第二轮独立复核阻断问题修复）

### 阻断问题 1：返回列表时错误继承旧筛选，上下文“存在/不存在”未单值化
- 根因：`DecisionCreatePage / DecisionDetailPage` 无条件读取全局 `lastSearch` 作为返回参数，从 `Module Detail` 或外部直达进入时也会错误恢复历史筛选
- 修复：引入显式路由搜索参数 `fromList`（`z.boolean().optional()`）单值化“来源列表上下文存在 / 不存在”
  - `routes/decisions/new.tsx`、`routes/decisions/$decisionId.tsx` 的 `validateSearch` 增加 `fromList`
  - `DecisionListPage`、`DecisionListContent` 导航到 `Create / Detail` 时显式带 `fromList: true`
  - `DecisionCreatePage` 创建成功回流 `Detail` 时透传 `fromList: search.fromList`
  - `DecisionCreatePage / DecisionDetailPage` 返回列表统一为 `search.fromList ? lastSearch : { statusFilter: 'all' }`
  - `ModuleDecisionEntryPanel` 进入 `Create / Detail` 不带 `fromList`，返回列表落默认参数
- 验证：`npm run build` 通过，路由树重新生成包含 `fromList` schema

### 阻断问题 2：待关联目标缺少“主动放弃关联”出口，结束条件少了一半
- 根因：`DecisionPendingLinkTargetCard` 注释承诺“直到用户完成 LinkDecisionToTarget 或主动放弃关联”，但实现只有“关联后消失”，且 `spec.md` 只承诺“关联后消失”，后端无清除 `source_context` 接口
- 修复（采用复核选项 B：收敛到规格承诺）：把结束条件收敛为“仅在正式 `LinkDecisionToTarget` 写入后消失”
  - `DecisionPendingLinkTargetCard` 注释删除“或主动放弃关联”，收敛为“仅在正式关联后消失”
  - `types.ts` `source_context` 注释对齐
  - `DecisionDetailPage` 注释对齐
  - `spec.md` 详情页 Scenario 显式声明：待关联目标仅在正式关联后消失，当前阶段不提供主动放弃出口，`source_context` 作为入口历史记录保留
- 验证：`npm run build` 通过，`oxlint` 0 errors

### 上游规格同步（phase03-10 正式规格口径漂移修复）
- 触发：第三轮复核指出 phase03-10 正式规格 `decision_center_spec_v0.1.md` 仍保留“或主动放弃关联”旧承诺，与 phase03-13 已收敛边界不一致
- 同步范围：全量扫描确认旧承诺仅存在于 phase03-10 spec 三处（L227 / L553 / L589），三件套基线文档（shared_baseline / architecture_plan / dev_plan）无残留
- 对齐式更新（不推翻叙事）：
  - `§5.11` L227：收敛为“持续到完成正式 LinkDecisionToTarget；当前阶段不提供主动放弃出口，source_context 作为入口历史记录保留，待关联目标仅在正式关联写入后消失”
  - `§9.4` L553：`DecisionPendingLinkTargetCard` 收敛为“直到完成正式 LinkDecisionToTarget；不提供主动放弃出口，仅在正式关联写入后由详情页 reread 驱动消失”
  - `§9.5` L589：待关联目标承接状态收敛为“持续到完成正式 LinkDecisionToTarget；不提供主动放弃出口，待关联目标仅在正式关联写入后消失”
- 一致性校验：phase03-10 spec §9.5「跨页面列表上下文单值承接」（L595-605）已定义 fromList 单值化语义，与 phase03-13 实现一致
- 验证：`grep "或主动放弃" decision_center_spec_v0.1.md` 返回 0，旧表述已全部清除

### P2 横向漂移收口（子代理第三轮复核建议）
- P2-1（已收口）：phase03-12 spec `source_context` Requirement 两处旧表述同步
  - `spec.md` L13：“持续到完成或放弃” → “持续到正式关联完成”
  - `spec.md` L216：“支持持续到...或主动放弃关联的跨刷新承接语义” → “支持持续到...的跨刷新承接语义（当前阶段不提供主动放弃出口，source_context 作为入口历史记录保留）”
  - 验证：`grep "或主动放弃" phase03_12 spec.md` 返回 0
- P2-2（不修，记录决策）：前置子规格 phase03-03 / phase03-05 / phase03-06 共 6 处历史残留
  - 决策依据：phase03-10 `decision_center_spec_v0.1.md` L828 已声明“前置子规格 phase03-01~09 仅作为冻结来源与追踪依据，不再作为后续实现与验收的长期并列入口”
  - 影响评估：phase03-14/15 以 phase03-10 正式规格为准，前置子规格旧表述不构成误导
  - 处置：当前阶段不单独修复，若后续做 phase03 系列规格一致性收口可统一对齐
- 活动规格链一致性结论：phase03-10 / phase03-12 / phase03-13 三个活动规格的“待关联目标结束条件”口径已完全对齐为“仅在正式 `LinkDecisionToTarget` 写入后消失，不提供主动放弃出口”
