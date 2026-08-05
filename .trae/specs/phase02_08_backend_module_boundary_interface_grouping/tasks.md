# Tasks

- [x] Task 1: 收敛 `Module Registry` 的后端模块归属边界。
  - [x] SubTask 1.1: 明确 `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository` 归属于 `Module Registry` 后端模块
  - [x] SubTask 1.2: 明确 `CreateProduct`、`CreateRepository`、`RecordDecision` 不归属于当前阶段的 `Module Registry` 后端模块

- [x] Task 2: 收敛 `Module Registry` 的最小读接口分组。
  - [x] SubTask 2.1: 明确 `ModuleListRead` 的职责边界
  - [x] SubTask 2.2: 明确 `ModuleDetailRead` 作为统一详情读模型宿主
  - [x] SubTask 2.3: 明确列表读取与详情读取不得扩写为 Dashboard 或跨对象分析读取

- [x] Task 3: 收敛 `Module Registry` 的最小写接口分组。
  - [x] SubTask 3.1: 明确 `ModuleCreateWrite` 只承接模块自身登记
  - [x] SubTask 3.2: 明确 `ModuleReleaseWrite` 只承接依附 `moduleId` 的版本登记
  - [x] SubTask 3.3: 明确 `ModuleBindingWrite` 统一承接产品绑定与仓库映射

- [x] Task 4: 收敛 `Product / Repository / Decision` 的服务侧连接边界。
  - [x] SubTask 4.1: 明确 `Product` 侧只提供候选读取与关系校验边界
  - [x] SubTask 4.2: 明确 `Repository` 侧只提供候选读取与关系校验边界
  - [x] SubTask 4.3: 明确 `Decision` 侧作为 `ModuleDetailRead` 内嵌附属读取承接，不设独立读接口组
  - [x] SubTask 4.4: 明确 `phase02` 阶段 `Product / Repository` 跨模块候选读取由 `Module Registry` 临时承接并通过接口边界与独立代码落点隔离

- [x] Task 5: 收敛后端模块内部的实现前分层语义。
  - [x] SubTask 5.1: 明确入口层、业务编排层、数据访问/外部连接层的职责边界
  - [x] SubTask 5.2: 明确当前阶段不冻结 Go HTTP/RPC 框架、ORM 或数据访问层具体工具
  - [x] SubTask 5.3: 明确合同边界与存储模型解耦

- [x] Task 6: 冻结后端文件/目录落点到实现结构层。
  - [x] SubTask 6.1: 明确模块根包落在 `backend/internal/moduleregistry/`，按 `handler/service/repository/candidate` 四子包组织
  - [x] SubTask 6.2: 明确读组文件落点（`handler/query_handler.go` + `service/query_service.go`）
  - [x] SubTask 6.3: 明确写组文件落点（`handler/command_handler.go` + `service/command_service.go`）
  - [x] SubTask 6.4: 明确数据访问层文件落点（`repository/module_store.go` + `release_store.go` + `binding_store.go`）
  - [x] SubTask 6.5: 明确跨模块候选读取文件落点（`candidate/product_candidate_read.go` + `repository_candidate_read.go`）
  - [x] SubTask 6.6: 明确 `Decision` 读取不设独立文件落点，内嵌于 `service/query_service.go`
  - [x] SubTask 6.7: 明确文件落点只冻结职责分工，不冻结 Go HTTP/RPC/ORM 等实现工具

- [x] Task 7: 完成规格校验。
  - [x] SubTask 7.1: 验证后端模块边界已经明确
  - [x] SubTask 7.2: 验证读写接口分组已经明确
  - [x] SubTask 7.3: 验证与 `Product / Repository / Decision` 的连接边界已经明确且与 `phase02-05` 单值性一致
  - [x] SubTask 7.4: 验证后端文件/目录落点已冻结到可直接创建文件的层级
  - [x] SubTask 7.5: 验证未把 Go 数据访问层具体工具提前写成既成事实
  - [x] SubTask 7.6: 验证设计结果足以直接进入实现
  - [x] SubTask 7.7: 验证 `products / repositories` 候选读取前提已同步到 `shared_baseline`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 2`, `Task 3`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
