# Tasks

- [x] Task 1: 产出 `Decision Center` 的页面文件级页面集合与职责设计。
  - [x] SubTask 1.1: 明确 `DecisionListPage`、`DecisionCreatePage`、`DecisionDetailPage` 的页面集合
  - [x] SubTask 1.2: 明确每个页面对应的单一职责，不拆出 `DecisionLinkPage`、`ProductPage`、`RepositoryPage` 的独立页面主线
  - [x] SubTask 1.3: 将页面集合映射到预期的 `frontend/src/features/decision-center/pages/*` 文件落点

- [x] Task 2: 产出最小前端路由结构、URL 语义与路由树。
  - [x] SubTask 2.1: 明确 `/decisions`、`/decisions/new`、`/decisions/:decisionId`
  - [x] SubTask 2.2: 明确 `DecisionListRoute` 的 `/decisions` 承接 `queryText` 与 `statusFilter` 作为路由搜索参数，详细状态映射延后至 `phase03-06`
  - [x] SubTask 2.3: 明确 `List -> Create / Detail`、`Create -> Detail (RecordDecision 成功后)` 与 `Detail -> List` 的进入关系
  - [x] SubTask 2.4: 明确 `Module Detail -> DecisionCreateRoute` 与 `Module Detail -> DecisionListRoute` 的入口关系
  - [x] SubTask 2.5: 将上述路由映射到预期的 `frontend/src/routes/decisions/*` 文件落点
  - [x] SubTask 2.6: 验证未把 `Product / Repository` 扩写为并列主树
  - [x] SubTask 2.7: 明确 `Module Detail` 侧现有入口（`frontend/src/routes/modules/$moduleId.tsx` 与 `module-decision-entry-panel.tsx`）到新 `DecisionCreateRoute` / `DecisionListRoute` 的文件落点映射

- [x] Task 3: 产出页面壳层、组件树与组件归属设计。
  - [x] SubTask 3.1: 明确 `DecisionListPageShell` 与列表页组件树
  - [x] SubTask 3.2: 明确 `DecisionCreatePageShell` 与创建页组件树
  - [x] SubTask 3.3: 明确 `DecisionDetailPageShell` 与详情页组件树
  - [x] SubTask 3.4: 明确页面专属组件与共享组件的边界
  - [x] SubTask 3.5: 验证组件职责已覆盖列表、模板表单、目标关联面板
  - [x] SubTask 3.6: 明确 `DecisionPendingLinkTargetCard` 承接从 `Module Detail` 带入的待关联 `Module` 的显式状态，承接 `phase03-03` 已冻结结论

- [x] Task 4: 产出 `PC / 移动浏览器` 布局降级策略。
  - [x] SubTask 4.1: 明确桌面端的高信息密度列表布局策略
  - [x] SubTask 4.2: 明确桌面端详情页的分区布局（含待关联目标区）
  - [x] SubTask 4.3: 明确移动浏览器下的单列、折叠与垂直重排策略（含待关联目标区）
  - [x] SubTask 4.4: 明确移动浏览器下创建页的来源上下文区、表单区与动作区布局
  - [x] SubTask 4.5: 验证未引入第二套移动端 UI 架构、独立 `React Native` 或完整 `PWA`

- [x] Task 5: 完成规格校验。
  - [x] SubTask 5.1: 验证前端页面与路由分层已经明确到可实现层
  - [x] SubTask 5.2: 验证页面级组件职责已经明确到壳层、组件树与组件归属层
  - [x] SubTask 5.3: 验证无第二套移动端 UI 架构
  - [x] SubTask 5.4: 验证设计结果足以直接进入实现，而不是只停留在一致性复述
  - [x] SubTask 5.5: 验证 `phase03-03` 待关联目标语义已在 `Decision Detail` 组件树中单值冻结

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
