# phase02-10 前端 Module Registry 主线实现 Spec

## Why

`phase02-09` 已经把 `Module Registry` 的正式规格正文冻结为单一实现入口，但当前仓库还没有对应的前端可运行主线。`phase02-10` 需要把页面、状态、交互与响应式布局从“可实现设计”推进到“可运行前端”，同时保证不引入第二套移动端 UI 架构，也不因为后端尚未完成而临时发明第二套数据语义。

## What Changes

- 实现 `Module List / Module Create / Module Detail / Release Create` 前端主线
- 实现列表读取、模块创建、详情展示、版本登记、产品绑定、仓库映射与 `Decision` 只读入口的最小前端交互
- 冻结前端数据适配边界：在 `phase02-11` 后端主线完成前，允许使用最小前端适配层承接演示与状态流，但不得偏离 `phase02-09` 已冻结的数据与 API 语义
- 落实单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略
- 明确实现验收以“可运行、可走通、无第二套移动端 UI 架构”为准

## Impact

- Affected specs:
  - `phase02_09_module_registry_formal_spec`
  - `phase02_06_frontend_page_route_component_design`
  - `phase02_07_frontend_state_interaction_flow`
- Affected code:
  - `frontend/src/routes/modules/`
  - `frontend/src/features/module-registry/pages/`
  - `frontend/src/features/module-registry/components/`
  - `frontend/src/features/module-registry/`

## ADDED Requirements

### Requirement: 实现单一前端主线页面与路由

系统 SHALL 按 `phase02-09` 已冻结的页面文件落点与路由语义，产出 `Module Registry` 的前端可运行主线，不得改写为第二套路由树或并列工作台。

#### Scenario: 页面主线可运行

- **WHEN** 用户进入 `/modules`
- **THEN** 系统必须能渲染 `ModuleListPage`
- **AND** 用户必须能继续进入 `/modules/new`、`/modules/:moduleId`、`/modules/:moduleId/releases/new`
- **AND** 页面文件落点必须与 `phase02-09` 正式规格正文保持一致

### Requirement: 承接最小前端交互闭环

系统 SHALL 在前端承接 `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository` 的最小交互闭环，并保持与 `phase02-09` 冻结的回流规则、错误呈现位置和状态模型一致。

#### Scenario: 模块创建闭环

- **WHEN** 用户从 `ModuleListPage` 进入 `ModuleCreatePage` 并成功提交最小字段
- **THEN** 系统必须默认回流到对应的 `ModuleDetailPage`
- **AND** 用户后续返回 `ModuleListPage` 时必须能够承接新建模块结果

#### Scenario: 版本登记闭环

- **WHEN** 用户从 `ModuleDetailPage` 进入 `ReleaseCreatePage` 并成功提交版本信息
- **THEN** 系统必须默认回流到当前模块的 `ModuleDetailPage`
- **AND** 详情页必须承接最新版本列表读取结果

#### Scenario: 绑定动作闭环

- **WHEN** 用户在 `ModuleDetailPage` 执行 `BindModuleToProduct` 或 `MapModuleToRepository`
- **THEN** 系统必须停留在当前详情页
- **AND** 必须重新读取并展示最新绑定结果
- **AND** 绑定失败时错误必须停留在面板上下文

### Requirement: 冻结前端数据适配边界

系统 SHALL 在 `phase02-11` 后端主线完成前，为前端主线提供最小数据适配边界，以支撑页面运行与交互演示；但该适配层不得发明第二套对象字段、状态含义或返回路径语义。

#### Scenario: 后端未完成时的前端运行

- **WHEN** `phase02-10` 实现时后端与数据库主线尚未完成
- **THEN** 系统允许使用前端最小数据适配层、mock 数据或本地演示数据支撑页面运行
- **AND** 适配层输出的数据结构必须对齐 `phase02-09` 的 `Module`、`Release`、列表读取、详情读取与绑定读取语义
- **AND** 不得把该适配层提升为长期并列数据主线

### Requirement: 落实单一 React Web 的响应式布局

系统 SHALL 在同一套 `React Web` 页面与组件主线中，同时承接 `PC` 与移动浏览器布局，不得通过新增第二套移动端页面或第二套交互树来满足适配要求。

#### Scenario: PC 与移动浏览器共用单一主线

- **WHEN** 用户分别在桌面宽屏和移动浏览器访问 `Module List`、`Module Detail`、`Module Create`、`Release Create`
- **THEN** 系统必须通过同一套页面与组件结构完成布局降级
- **AND** `ModuleDetailPage` 在窄屏下必须按摘要、版本、关联、`Decision` 入口的垂直顺序重排
- **AND** 不得引入独立移动端 UI 架构

### Requirement: 保持列表上下文与错误停留规则

系统 SHALL 继续遵守 `phase02-07` 与 `phase02-09` 已冻结的状态模型，尤其是列表搜索参数、表单错误停留、绑定面板状态与返回路径。

#### Scenario: 返回列表恢复上下文

- **WHEN** 用户从 `ModuleCreatePage` 或 `ModuleDetailPage` 返回 `ModuleListPage`
- **THEN** 系统必须恢复原有 `queryText` 与 `statusFilter`
- **AND** 不得要求用户重新输入筛选条件

#### Scenario: 错误反馈停留在原上下文

- **WHEN** 列表读取失败、创建失败、版本提交失败或绑定动作失败
- **THEN** 错误反馈必须停留在当前页面或当前面板上下文
- **AND** 不得跳转到独立错误页

## MODIFIED Requirements

### Requirement: `phase02-09` 正式规格正文进入实现态

`phase02-09` 已冻结的前端页面、组件、路由、状态模型与布局降级要求，在本阶段必须从“可直接实现的设计结果”推进为“可运行前端主线”，并作为 `phase02-12` 联调与验收的直接前置条件。

#### Scenario: 进入 phase02-12 的前置满足

- **WHEN** `phase02-10` 完成
- **THEN** 前端主线必须已经可以独立运行
- **AND** 列表、创建、详情、版本登记与关联动作必须在前端走通
- **AND** 实现结果不得偏离 `phase02-09` 已冻结的页面职责、回流规则、接口语义与非目标边界

## REMOVED Requirements
