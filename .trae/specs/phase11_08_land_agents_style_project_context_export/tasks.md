# Tasks

- [x] Task 1: 盘点 `phase11-08` 的直接上游与冻结边界
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L223-L237` 中 `phase11-08` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase11-05` 中关于结构化读取与 AGENTS 风格 Markdown 导出职责边界的冻结口径
  - [x] SubTask 1.3: 审阅 `phase11-07` 中关于结构化只读结果、入口字段与失败语义的已交付结果

- [x] Task 2: 冻结 AGENTS 风格导出的正式承接边界
  - [x] SubTask 2.1: 固定导出能力必须落在 Go backend 的正式只读业务接口层
  - [x] SubTask 2.2: 固定导出能力的身份是 `phase11-07` 结构化结果的文档化投影，而不是第二套读取主线
  - [x] SubTask 2.3: 固定导出能力不扩写为主动注入、仓库写入或外部仓库同步

- [x] Task 3: 冻结 Markdown 导出的单向派生关系与内容边界
  - [x] SubTask 3.1: 固定 Markdown 导出只能从同一 `repository_id` 的结构化只读结果单向派生
  - [x] SubTask 3.2: 固定 Markdown 导出的最小内容边界：仓库摘要、phase/主交付、实体摘要、明确不做、规则与文档入口
  - [x] SubTask 3.3: 固定文档入口只允许承接结构化结果中已有的入口字段，不允许临时补充第二套事实

- [x] Task 4: 冻结 PSCO 自身 dogfooding 与验收口径
  - [x] SubTask 4.1: 固定 `PSCO` 仓库自身为第一 dogfooding 场景，但不外推为未来项目模板
  - [x] SubTask 4.2: 固定新路径继续遵守 `<= 3` 固定入口约束
  - [x] SubTask 4.3: 固定导出结果必须帮助回答预设的 `5` 个恢复问题

- [x] Task 5: 将 `phase11-08` 的冻结结果回写到目标源文档
  - [x] SubTask 5.1: 在 `phase11_project_context_foundation_dev_plan.md` 中冻结 `phase11-08` 的范围、产物与 DoD
  - [x] SubTask 5.2: 在 `phase11_project_context_foundation_architecture_plan.md` 中冻结导出 owner、单向派生与 dogfooding 验收边界
  - [x] SubTask 5.3: 在 `phase11_project_context_foundation_shared_baseline.md` 中冻结导出内容边界、入口约束与非目标

- [x] Task 6: 创建导出合同并生成代码
  - [x] SubTask 6.1: 更新 `proto/psco/project_context/v1/project_context.proto`，为 Markdown 导出增加正式只读合同
  - [x] SubTask 6.2: 执行 `buf generate` 生成 Go proto 与 Connect 代码

- [x] Task 7: 实现 Backend 导出路径
  - [x] SubTask 7.1: 在 `backend/internal/projectcontext/` 中增加 Markdown 导出 DTO / renderer 承接位
  - [x] SubTask 7.2: 在 `service/query_service.go` 中编排“结构化读取 -> Markdown 导出”的单向派生
  - [x] SubTask 7.3: 在 `connect/server.go` 中增加正式导出 handler，并保持错误语义与只读边界一致

- [x] Task 8: 完成装配与回归验证
  - [x] SubTask 8.1: 在 `backend/internal/platform/` 中挂载导出接口
  - [x] SubTask 8.2: 补充覆盖成功路径的集成验证，证明导出结果来自结构化读取而不是第二套事实源
  - [x] SubTask 8.3: 补充 PSCO 自身 dogfooding 验收记录所需的验证脚本或测试证据
  - [x] SubTask 8.4: 执行 `go test ./...` 与 `go build ./...` 完成回归验证

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4
- Task 6 depends on Task 5
- Task 7 depends on Task 6
- Task 8 depends on Task 7
