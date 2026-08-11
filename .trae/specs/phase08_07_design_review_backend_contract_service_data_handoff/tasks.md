# Tasks

- [x] Task 1: 冻结 review 后端正式合同包与 RPC 矩阵
  - [x] SubTask 1.1: 明确 `proto/psco/review/v1/review.proto` 的包名、服务名与生成落点
  - [x] SubTask 1.2: 明确 `GetDailyReviewContext / GetWeeklyReviewContext / SubmitReviewResult` 三条 RPC 的最小职责
  - [x] SubTask 1.3: 明确 review proto 直接 import 的 canonical message 清单，避免重复抄字段

- [x] Task 2: 冻结 review read model 的后端 owner 与协作边界
  - [x] SubTask 2.1: 明确 `review.QueryService` 的正式落点与 Daily Review 的读取来源
  - [x] SubTask 2.2: 明确 `review.QueryService` 的 Weekly Review 读取来源与 `phase06` 复用感知消费位置
  - [x] SubTask 2.3: 明确 review query owner 不得绕过 `dashboard / decisioncenter / reusesummary` 既有 query service

- [x] Task 3: 冻结 review 动作命令与结果回流的后端 owner
  - [x] SubTask 3.1: 明确实体写入继续复用 `Decision / Product / Repository` 既有 canonical command owner，并补齐 `Module Registry` 的 canonical handoff 服务边界
  - [x] SubTask 3.2: 明确 `SubmitReviewResult` 只承接最小流程结果，不复制实体写模型
  - [x] SubTask 3.3: 明确 review 成功响应只返回最小结果 envelope 或记录标识

- [x] Task 4: 冻结 `review_record` 的最小数据承接设计
  - [x] SubTask 4.1: 明确 `review_records` 单表的最小字段集合
  - [x] SubTask 4.2: 明确 `backend/internal/review/repository/review_record_store.go` 与 `review.CommandService` 的职责边界
  - [x] SubTask 4.3: 明确 `next-step result` 必须有正式落点，而 `decision handoff / entity handoff` 允许无 record 路径

- [x] Task 5: 冻结 review Connect transport 与 platform 接线方式
  - [x] SubTask 5.1: 明确 `backend/internal/review/connect/server.go` 与 `mountReviewConnect(...)` 的正式落点
  - [x] SubTask 5.2: 明确 review handler 继续通过 generated `(path, handler)` 挂到单一 `/api` 业务树
  - [x] SubTask 5.3: 明确 review 新增错误继续统一收敛到 `connecterrors.MapToConnectError`

- [x] Task 6: 冻结工具链与 API 验收清单
  - [x] SubTask 6.1: 明确 `buf build / gen / lint` 与 `go build ./...` 的正式验收命令
  - [x] SubTask 6.2: 明确 `GetDailyReviewContext / GetWeeklyReviewContext / SubmitReviewResult` 的最小 `/api` smoke 清单，并纳入至少一种实体回流或 canonical action handoff
  - [x] SubTask 6.3: 明确既有 `DecisionCenterService.CreateDecision` 继续作为 review 后续 canonical path 的验收说明

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 1`
- `Task 6` depends on `Task 5`
