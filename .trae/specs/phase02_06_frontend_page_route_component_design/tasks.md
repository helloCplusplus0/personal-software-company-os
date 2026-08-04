# Tasks

- [x] Task 1: 产出 `Module Registry` 的页面文件级页面集合与职责设计。
  - [x] SubTask 1.1: 明确 `ModuleListPage`、`ModuleCreatePage`、`ModuleDetailPage`、`ReleaseCreatePage` 的页面集合
  - [x] SubTask 1.2: 明确每个页面对应的单一职责，不拆出 `Product / Repository / Decision` 的独立页面主线
  - [x] SubTask 1.3: 将页面集合映射到预期的 `frontend/src/features/module-registry/pages/*` 文件落点

- [x] Task 2: 产出最小前端路由结构、URL 语义与路由树。
  - [x] SubTask 2.1: 明确 `/modules`、`/modules/new`、`/modules/:moduleId`、`/modules/:moduleId/releases/new`
  - [x] SubTask 2.2: 明确 `List -> Create / Detail` 与 `Detail -> ReleaseCreate` 的进入关系
  - [x] SubTask 2.3: 将上述路由映射到预期的 `frontend/src/routes/modules/*` 文件落点
  - [x] SubTask 2.4: 明确 `Module Detail` 到 `Product Registry / Repository Binding / Decision Center` 的轻量入口

- [x] Task 3: 产出页面壳层、组件树与组件归属设计。
  - [x] SubTask 3.1: 明确 `ModuleListPageShell` 与列表页组件树
  - [x] SubTask 3.2: 明确 `ModuleCreatePageShell` 与创建页组件树
  - [x] SubTask 3.3: 明确 `ModuleDetailPageShell` 与详情页组件树
  - [x] SubTask 3.4: 明确 `ReleaseCreatePageShell` 与版本登记页组件树
  - [x] SubTask 3.5: 明确页面专属组件与共享组件的边界

- [x] Task 4: 产出 `PC / 移动浏览器` 布局降级策略。
  - [x] SubTask 4.1: 明确桌面端的高信息密度布局策略
  - [x] SubTask 4.2: 明确桌面端详情页的分区布局
  - [x] SubTask 4.3: 明确移动浏览器下的单列、折叠与垂直重排策略
  - [x] SubTask 4.4: 明确移动浏览器下创建页与版本登记页的表单布局
  - [x] SubTask 4.5: 验证未引入第二套移动端 UI 架构、独立 `React Native` 或完整 `PWA`

- [x] Task 5: 完成规格校验。
  - [x] SubTask 5.1: 验证前端页面与路由分层已经明确到可实现层
  - [x] SubTask 5.2: 验证页面级组件职责已经明确到壳层、组件树与组件归属层
  - [x] SubTask 5.3: 验证无第二套移动端 UI 架构
  - [x] SubTask 5.4: 验证设计结果足以直接进入实现，而不是只停留在一致性复述

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
