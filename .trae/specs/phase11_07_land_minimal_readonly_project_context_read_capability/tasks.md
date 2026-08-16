# Tasks

- [x] Task 1: 盘点 `phase11-07` 的实现输入与冻结边界
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L197-L222` 中 `phase11-07` 的范围、正式产物与 DoD
  - [x] SubTask 1.2: 审阅 `phase11_project_context_foundation_shared_baseline.md` 中项目上下文导出矩阵、字段边界与验收前提
  - [x] SubTask 1.3: 审阅 `phase11_project_context_foundation_architecture_plan.md` 中项目上下文导出的正式边界、输入合同与不做清单

- [x] Task 2: 明确最小只读项目上下文能力的正式承接位
  - [x] SubTask 2.1: 固定最小只读项目上下文能力必须落在 Go backend 的正式只读业务接口承接位
  - [x] SubTask 2.2: 固定该能力继续复用 `.proto + ConnectRPC` 主线，不引入第二套 canonical API 或新协议层
  - [x] SubTask 2.3: 固定该能力的身份是聚合投影，而不是新实体主线或前端侧拼装查询

- [x] Task 3: 冻结输入合同与未绑定仓库失败语义
  - [x] SubTask 3.1: 固定 `repository_id` 为唯一正式结构化输入锚点
  - [x] SubTask 3.2: 固定只承接"已完成 Repository Binding"的仓库上下文读取，并明确以 `product_repositories + module_repositories` 同时存在作为完成判定
  - [x] SubTask 3.3: 固定仓库不存在/绑定不完整两类失败态与返回语义，避免执行者临场补猜 current project

- [x] Task 4: 冻结最小结构化只读输出边界
  - [x] SubTask 4.1: 固定最小输出至少覆盖 `Repository / Product / Module / Decision` 摘要与状态
  - [x] SubTask 4.2: 固定规则、约束与文档入口字段的最小承接边界，并要求提供可定位入口
  - [x] SubTask 4.3: 固定 `Decision` 聚合、去重与归档过滤继续遵守 `phase11-05` 已冻结的两类 module-link 派生命中口径

- [x] Task 5: 保持 canonical 一致性与只读边界
  - [x] SubTask 5.1: 明确该能力只读取既有 canonical 数据，不引入第二套业务事实源
  - [x] SubTask 5.2: 明确不得把 agent 写回、Draft、审批流或 agent 专属一级业务对象偷渡进 `phase11-07`
  - [x] SubTask 5.3: 明确不得把读取承接位扩成 MCP、CLI、前端对话式入口或其他新协议层

- [x] Task 6: 明确通用能力与消费侧目录结构的边界
  - [x] SubTask 6.1: 固定该能力不要求消费侧项目目录与 `PSCO` 当前仓库拥有相同结构
  - [x] SubTask 6.2: 固定 `README.md / AGENTS.md / rules` 等固定文件名不是必要输入合同
  - [x] SubTask 6.3: 固定未来最佳实践项目模板仅属候选增强，不是 `phase11-07` 的前置依赖

- [x] Task 7: 将 `phase11-07` 的实施对象显式落到正式产物
  - [x] SubTask 7.1: 明确正式产物至少包括最小结构化只读读取承接位（落在 Go backend 的 `.proto + ConnectRPC` 正式主线）
  - [x] SubTask 7.2: 明确正式产物至少包括输入合同、输出边界与失败语义
  - [x] SubTask 7.3: 明确正式产物至少包括与既有 canonical 数据的一致性说明

- [x] Task 8: 冻结 `phase11-07` 的成功标准、DoD 与收口口径
  - [x] SubTask 8.1: 固定"何时算最小只读能力已经正式落地"（正式承接位 + 正式合同 + 正式失败语义）
  - [x] SubTask 8.2: 固定"何时不得判定完成"，包括临场补锚、目录结构依赖与第二套 API
  - [x] SubTask 8.3: 保证后续执行者无需再猜"读取能力做到什么程度就够了"

- [x] Task 9: 将验收对象显式落到后续实现结果，而不是只检查 spec 包自身
  - [x] SubTask 9.1: 确认验收对象包括正式只读承接位、正式输入合同与正式失败语义
  - [x] SubTask 9.2: 确认验收对象包括与既有 canonical 数据一致的最小输出结果
  - [x] SubTask 9.3: 确认验收对象包括"未依赖消费侧目录结构、未引入第二套 canonical API、未引入 agent 写回"的实际实现结果

- [x] Task 10: 将 `phase11-07` 的冻结结果显式回写到目标源文档
  - [x] SubTask 10.1: 在 `phase11_project_context_foundation_dev_plan.md` 中冻结正式承接位、输入合同、输出边界与 DoD 口径
  - [x] SubTask 10.2: 在 `phase11_project_context_foundation_architecture_plan.md` 中冻结结构化只读输出边界、输入合同与不做清单（§4.4 已存在，无需修改）
  - [x] SubTask 10.3: 在 `phase11_project_context_foundation_shared_baseline.md` 中冻结导出矩阵、字段边界、失败语义与非目标边界（§3.6 已存在，无需修改）

- [x] Task 11: 创建 Proto 合同并生成 Go 代码
  - [x] SubTask 11.1: 创建 `proto/psco/project_context/v1/project_context.proto`，定义 `GetProjectContext` RPC 及 6 类输出消息
  - [x] SubTask 11.2: 执行 `buf generate` 生成 Go proto 与 Connect 代码

- [x] Task 12: 实现 Backend 模块（5 个源文件）
  - [x] SubTask 12.1: 创建 `types.go`（跨层共享 DTO，对齐 Proto 结构）
  - [x] SubTask 12.2: 创建 `errors.go`（仓库不存在 / 绑定不完整 / 聚合读取失败的业务错误哨兵）
  - [x] SubTask 12.3: 创建 `candidate/context_readers.go`（跨模块 reader：ReadRepository / ReadProduct / ReadModules / ReadDecisions / ReadRules / ReadPhases）
  - [x] SubTask 12.4: 创建 `service/query_service.go`（只读聚合编排 GetProjectContext）
  - [x] SubTask 12.5: 创建 `connect/server.go`（Connect handler + DTO→Proto 转换）

- [x] Task 13: 路由装配与错误映射注册
  - [x] SubTask 13.1: 在 `router.go` 中添加 `buildProjectContext` / `mountProjectContextConnect` 函数
  - [x] SubTask 13.2: 在 `server.go` 的 `NewServer` 中调用装配与挂载
  - [x] SubTask 13.3: 在 `connect_errors.go` 中注册仓库不存在与绑定不完整的正式错误映射

- [x] Task 14: 编译验证与代码复核修复
  - [x] SubTask 14.1: `go build ./...` 编译通过
  - [x] SubTask 14.2: 子代理独立复核通过
  - [x] SubTask 14.3: 修复阻断性问题：`ReadRepository` 区分 `pgx.ErrNoRows` 与其他错误

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4
- Task 6 depends on Task 1
- Task 7 depends on Task 2
- Task 7 depends on Task 3
- Task 7 depends on Task 4
- Task 7 depends on Task 5
- Task 7 depends on Task 6
- Task 8 depends on Task 7
- Task 9 depends on Task 8
- Task 10 depends on Task 8
- Task 11 depends on Task 10
- Task 12 depends on Task 11
- Task 13 depends on Task 12
- Task 14 depends on Task 13
