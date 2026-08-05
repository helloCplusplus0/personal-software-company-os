# Tasks

- [x] Task 1: 搭建 `Module Registry` 前端主线路由与页面壳层。
  - [x] SubTask 1.1: 建立 `/modules`、`/modules/new`、`/modules/:moduleId`、`/modules/:moduleId/releases/new` 的路由文件与页面挂载关系
  - [x] SubTask 1.2: 落地 `ModuleListPage`、`ModuleCreatePage`、`ModuleDetailPage`、`ReleaseCreatePage` 的页面壳层
  - [x] SubTask 1.3: 保证路由与页面文件落点符合 `phase02-09` 正式规格正文

- [x] Task 2: 实现列表页与创建页的最小前端主线。
  - [x] SubTask 2.1: 实现 `ModuleListPage` 的工具栏、列表内容区与空状态入口
  - [x] SubTask 2.2: 实现 `queryText`、`statusFilter` 的路由搜索参数承接
  - [x] SubTask 2.3: 实现 `ModuleCreatePage` 的最小表单、提交动作、错误停留与成功回流

- [x] Task 3: 实现详情页、版本登记页与绑定面板的最小前端交互。
  - [x] SubTask 3.1: 实现 `ModuleDetailPage` 的摘要区、版本区、绑定区与 `Decision` 入口区
  - [x] SubTask 3.2: 实现 `ReleaseCreatePage` 的最小表单、提交动作与回流
  - [x] SubTask 3.3: 实现 `BindModuleToProduct` 与 `MapModuleToRepository` 的最小交互面板及其成功/失败状态

- [x] Task 4: 实现前端数据适配边界。
  - [x] SubTask 4.1: 建立仅服务 `phase02-10` 的最小前端数据适配层或 mock 承接
  - [x] SubTask 4.2: 保证列表、详情、版本、绑定与候选读取数据结构对齐 `phase02-09` 已冻结语义
  - [x] SubTask 4.3: 保证适配层不会演变成第二套长期数据主线

- [x] Task 5: 落实页面状态模型与返回路径规则。
  - [x] SubTask 5.1: 实现列表读取、创建提交、版本提交、绑定动作的最小状态流
  - [x] SubTask 5.2: 保证从创建页或详情页返回列表时恢复 `queryText` 与 `statusFilter`
  - [x] SubTask 5.3: 保证错误反馈停留在当前页面或当前面板上下文，不跳转独立错误页

- [x] Task 6: 落实单一 `React Web` 的响应式布局策略。
  - [x] SubTask 6.1: 完成 `ModuleListPage` 的 PC 高密度布局与移动端单列/卡片降级
  - [x] SubTask 6.2: 完成 `ModuleDetailPage` 的分区布局与移动端纵向重排
  - [x] SubTask 6.3: 保证 `ModuleCreatePage` 与 `ReleaseCreatePage` 在移动浏览器下主动作可见且无需第二套页面体系

- [x] Task 7: 验证前端主线达到 `phase02-10` DoD。
  - [x] SubTask 7.1: 验证前端主线可运行
  - [x] SubTask 7.2: 验证列表、创建、详情、版本登记与关联动作在前端可走通
  - [x] SubTask 7.3: 验证实现中未引入第二套移动端 UI 架构

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
