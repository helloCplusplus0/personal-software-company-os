# Tasks

- [x] Task 1: 为 `Product Detail` 落实下一步动作 CTA 与 Decision 入口面板
  - [x] SubTask 1.1: 为 `Product Detail` 新增 `ProductDecisionEntryPanel` 组件（复用 `ModuleDecisionEntryPanel` 模式，适配 `sourceProductId / sourceProductName`）
  - [x] SubTask 1.2: 为 `Product Detail` 新增页面级 `ProductNextActionBar` 组件，实现"Repository 缺口 → Module 缺口 → 返回"的 CTA 优先级
  - [x] SubTask 1.3: 在 `product-detail-page.tsx` 中集成 `ProductDecisionEntryPanel` 与 `ProductNextActionBar`，确保不破坏既有布局

- [x] Task 2: 为 `Module Detail` 落实下一步动作 CTA（绑定 handoff 转 canonical owner）
  - [x] SubTask 2.1: 升级 `ModuleBindingPanel` 的"绑定产品"CTA 为指向 canonical Product Detail 的导航 handoff（已在 phase04-13 实现）
  - [x] SubTask 2.2: 升级 `ModuleBindingPanel` 的"映射仓库"CTA 为指向 canonical Repository Detail 的导航 handoff（已在 phase04-13 实现）
  - [x] SubTask 2.3: 为 `Module Detail` 新增页面级 `ModuleNextActionBar` 组件，实现"Product 绑定 → Repository 映射 → 返回"的 CTA 优先级

- [x] Task 3: 为 `Repository Detail` 落实下一步动作 CTA 与 Decision 入口面板
  - [x] SubTask 3.1: 为 `Repository Detail` 新增 `RepositoryDecisionEntryPanel` 组件（复用 `ModuleDecisionEntryPanel` 模式，适配 `sourceRepositoryId / sourceRepositoryName`）
  - [x] SubTask 3.2: 为 `Repository Detail` 新增页面级 `RepositoryNextActionBar` 组件，实现"Product 绑定 → Module 映射 → 返回"的 CTA 优先级
  - [x] SubTask 3.3: 在 `repository-binding-detail-page.tsx` 中集成 `RepositoryDecisionEntryPanel` 与 `RepositoryNextActionBar`

- [x] Task 4: 落实从 detail page 返回 Dashboard / Review 的 reread 语义
  - [x] SubTask 4.1: 确保三个 detail page 的既有 `BackToDashboardButton` 返回后，Dashboard 的 feedback signals 与 overview 查询被正确失效刷新
  - [x] SubTask 4.2: 确保从 detail page 通过返回按钮返回来源页后，来源页的 pending / Current Focus 基于 canonical reread 正确收口
  - [x] SubTask 4.3: 补齐 product-detail 和 repository-binding-detail 的 `invalidateDetail` 语义，确保绑定成功后的失效范围覆盖 Dashboard 与 Review 相关 query

- [x] Task 5: 完成 `phase10-10` 浏览器级闭环验收
  - [x] SubTask 5.1: 验证 Product Detail 在各结构缺口状态下展示正确的下一步动作 CTA
  - [x] SubTask 5.2: 验证 Module Detail 的绑定 CTA 正确 handoff 到 canonical Product/Repository Detail
  - [x] SubTask 5.3: 验证 Repository Detail 在各结构缺口状态下展示正确的下一步动作 CTA
  - [x] SubTask 5.4: 验证从 detail page 返回 Dashboard 后 Current Focus 正确收口
  - [x] SubTask 5.5: 验证本子任务未改写 Decision pending 主线或 Dashboard/Daily Review 的 canonical 解释

- [x] Task 6: 完成 `phase10-10` 交付自检与边界复核
  - [x] SubTask 6.1: 复核三个 detail page 的 CTA 是否指向 canonical path，不做局部临时跳转
  - [x] SubTask 6.2: 复核 Decision 入口面板是否正确适配各实体（Product → sourceProductId、Repository → sourceRepositoryId）
  - [x] SubTask 6.3: 记录仍属于非目标的事项，确保未改写 Decision pending 主线

# Task Dependencies

- `Task 2` 可在 `Task 1` 之后独立执行
- `Task 3` 可在 `Task 1` 之后独立执行
- `Task 4` depends on `Task 1`、`Task 2`、`Task 3`
- `Task 5` depends on `Task 1` to `Task 4`
- `Task 6` depends on `Task 1` to `Task 5`