# Tasks

- [x] Task 1: 收敛 `Decision Center` 的后端模块归属边界。
  - [x] SubTask 1.1: 明确 `RecordDecision / LinkDecisionToTarget / DecisionListRead / DecisionDetailRead / DecisionModuleCandidateRead` 归属于 `Decision Center` 后端模块
  - [x] SubTask 1.2: 明确 `CreateModule / CreateProduct / CreateRepository` 不归属于当前阶段的 `Decision Center` 后端模块

- [x] Task 2: 收敛 `Decision Center` 的最小读接口分组。
  - [x] SubTask 2.1: 明确 `DecisionListRead` 的职责边界
  - [x] SubTask 2.2: 明确 `DecisionDetailRead` 作为统一详情读模型宿主
  - [x] SubTask 2.3: 明确 `DecisionModuleCandidateRead` 作为独立候选读取子组
  - [x] SubTask 2.4: 明确列表读取、详情读取与候选读取不得扩写为 Dashboard 或跨对象分析读取

- [x] Task 3: 收敛 `Decision Center` 的最小写接口分组。
  - [x] SubTask 3.1: 明确 `DecisionWrite` 只承接 `RecordDecision`
  - [x] SubTask 3.2: 明确 `DecisionLinkWrite` 只承接 `LinkDecisionToTarget`
  - [x] SubTask 3.3: 明确关联写入成功后回流到 `DecisionDetailRead`

- [x] Task 4: 收敛 `Module Registry` 的服务侧连接边界。
  - [x] SubTask 4.1: 明确 `Module Registry` 侧只提供 `Decision -> Module` 所需的最小候选读取与目标校验边界
  - [x] SubTask 4.2: 明确 `Decision Center` 不吸收 `Module Registry` 的主线写入
  - [x] SubTask 4.3: 明确 `Product / Repository` 在当前阶段仍然保持后移边界
  - [x] SubTask 4.4: 明确 `ModuleCandidateRead` 接口与实现由 `Decision Center` 的 `candidate/` 子包自己定义和拥有，`Module Registry` 不暴露专门服务契约，接线在应用装配点完成

- [x] Task 5: 收敛后端模块内部的实现前分层语义。
  - [x] SubTask 5.1: 明确入口层、业务编排层、数据访问/外部连接层的职责边界
  - [x] SubTask 5.2: 明确当前阶段不冻结 Go HTTP/RPC 框架、ORM 或数据访问层具体工具
  - [x] SubTask 5.3: 明确合同边界与存储模型解耦

- [x] Task 6: 冻结后端文件/目录落点到实现结构层。
  - [x] SubTask 6.1: 明确模块根包落在 `backend/internal/decisioncenter/`，按 `handler/service/repository/candidate` 四子包组织
  - [x] SubTask 6.2: 明确读组文件落点（`handler/query_handler.go` + `service/query_service.go`）
  - [x] SubTask 6.3: 明确写组文件落点（`handler/command_handler.go` + `service/command_service.go`）
  - [x] SubTask 6.4: 明确数据访问层文件落点（`repository/decision_store.go` + `link_store.go`）
  - [x] SubTask 6.5: 明确跨模块候选读取文件落点（`candidate/module_candidate_read.go`）
  - [x] SubTask 6.6: 明确文件落点只冻结职责分工，不冻结 Go HTTP/RPC/ORM 等实现工具
  - [x] SubTask 6.7: 明确支撑文件落点（`errors.go` + `types.go` + `validate.go` + `handler/response.go`），与现有 `moduleregistry` 同构

- [x] Task 7: 完成规格校验。
  - [x] SubTask 7.1: 验证后端模块边界已经明确
  - [x] SubTask 7.2: 验证读写接口分组已经明确
  - [x] SubTask 7.3: 验证与 `Module Registry` 的服务侧连接边界已经明确
  - [x] SubTask 7.4: 验证后端文件/目录落点已冻结到可直接创建文件的层级（含支撑文件）
  - [x] SubTask 7.5: 验证未把 Go 数据访问层具体工具提前写成既成事实
  - [x] SubTask 7.6: 验证设计结果足以直接进入实现
  - [x] SubTask 7.7: 验证 `ModuleCandidateRead` 接口拥有者与接线位置已单值化

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 2`, `Task 3`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
