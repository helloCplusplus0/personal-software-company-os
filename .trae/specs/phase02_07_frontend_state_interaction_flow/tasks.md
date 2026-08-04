# Tasks

- [x] Task 1: 收敛 `Module List` 的查询承接与页面状态模型。
  - [x] SubTask 1.1: 明确 `queryText / statusFilter` 必须归属于路由搜索参数
  - [x] SubTask 1.2: 明确读取状态与派生视图状态的分层关系
  - [x] SubTask 1.3: 明确列表错误反馈位置与空状态主动作
  - [x] SubTask 1.4: 明确从创建页或详情页返回时保留列表搜索参数上下文

- [x] Task 2: 收敛 `Module Create` 的草稿状态、提交状态与回流设计。
  - [x] SubTask 2.1: 明确 `name / description / status` 的最小草稿状态与 `idle / dirty` 语义
  - [x] SubTask 2.2: 明确 `submitting / submit-success / submit-error` 的最小提交状态
  - [x] SubTask 2.3: 明确提交失败时保留输入且在当前表单上下文展示错误
  - [x] SubTask 2.4: 明确提交成功后进入新建模块对应的 `ModuleDetailPage`
  - [x] SubTask 2.5: 明确成功写入后列表读模型需要重读的用户可见语义

- [x] Task 3: 收敛 `Release Create` 的路由上下文与版本登记交互流。
  - [x] SubTask 3.1: 明确 `moduleId` 必须来自当前路由参数
  - [x] SubTask 3.2: 明确版本登记页的最小草稿状态与提交状态
  - [x] SubTask 3.3: 明确提交失败后保留输入并停留当前 `moduleId` 上下文
  - [x] SubTask 3.4: 明确提交成功后回到当前 `ModuleDetailPage` 并承接最新版本读取

- [x] Task 4: 收敛 `ModuleBindingPanel` 的状态归属与面板交互规则。
  - [x] SubTask 4.1: 明确面板最小状态 `closed / open-idle / submitting / submit-success / submit-error`
  - [x] SubTask 4.2: 明确同一时刻只允许一个绑定面板或一种绑定模式处于打开态
  - [x] SubTask 4.3: 明确绑定成功后停留当前 `ModuleDetailPage` 并重读绑定结果
  - [x] SubTask 4.4: 明确绑定失败后错误反馈留在面板内并允许修正重试

- [x] Task 5: 收敛页面返回路径与统一回流宿主页。
  - [x] SubTask 5.1: 明确从 `ModuleCreatePage` 主动返回时回到保留原搜索参数的 `ModuleListPage`
  - [x] SubTask 5.2: 明确从 `ReleaseCreatePage` 主动返回时回到当前 `ModuleDetailPage`
  - [x] SubTask 5.3: 明确 `ModuleDetailPage` 为创建成功、版本登记成功与绑定动作后的统一回流宿主页
  - [x] SubTask 5.4: 明确 `Decision` 入口仍以 `ModuleDetailPage` 为主上下文

- [x] Task 6: 收敛局部 UI 状态归属边界。
  - [x] SubTask 6.1: 明确表单草稿、提交错误与面板开闭状态优先归属于当前页面或详情页上下文
  - [x] SubTask 6.2: 明确不得默认升级为跨路由全局状态

- [x] Task 7: 完成规格校验。
  - [x] SubTask 7.1: 验证页面级状态模型、状态归属与错误呈现位置已经明确
  - [x] SubTask 7.2: 验证列表、详情、创建、版本登记之间的交互流已经明确
  - [x] SubTask 7.3: 验证设计结果足以直接进入实现
  - [x] SubTask 7.4: 验证未把 hook 命名、缓存细节或 store API 提前写成既成事实

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
