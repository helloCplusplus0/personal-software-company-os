# Tasks

- [x] Task 1: 对齐 `phase05-11` 的直接上游与现有 proto 主线模式，明确这次任务是“合同主线落地”，不是重新设计一版 Dashboard 合同。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase05_dashboard_feedback_foundation_dev_plan.md` 中 `phase05-11` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase05-08` 已冻结的 `Dashboard + Feedback` 最小 `.proto` 合同设计
  - [x] SubTask 1.3: 对齐 `phase05-10` 正式正文、`phase03-11` 与 `phase04-11` 的 proto 主线落地模式
  - [x] SubTask 1.4: 对齐现有 `proto/README.md`、`proto/Makefile`、`proto/buf.yaml`、`proto/buf.gen.yaml` 的单 workspace 基线

- [x] Task 2: 冻结 `Dashboard + Feedback` 合同进入现有 proto workspace 的单一落点与包名语义。
  - [x] SubTask 2.1: 冻结合同源文件落点为 `proto/psco/dashboard/v1/dashboard.proto`
  - [x] SubTask 2.2: 冻结包名与版本语义为 `psco.dashboard.v1`
  - [x] SubTask 2.3: 冻结 `.proto` 作为 `Dashboard + Feedback` 当前阶段唯一合同定义入口
  - [x] SubTask 2.4: 冻结 `google.protobuf.Timestamp` import 与 `RecentActivityItem.activity_at` 的显式时间语义

- [x] Task 3: 冻结 `dashboard.proto` 必须承接的服务、消息与演进规则。
  - [x] SubTask 3.1: 冻结单一 `DashboardService` 与三类只读 RPC
  - [x] SubTask 3.2: 冻结 `DashboardOverview`、`FeedbackSignal`、`ProductAssetCoverageSummary`、`RecentActivityItem` 的最小落地范围
  - [x] SubTask 3.3: 冻结字段编号、枚举编号、`UNSPECIFIED = 0`、`reserved` 与 `missing both bindings` 的演进边界必须继续对齐 `phase05-08`

- [x] Task 4: 冻结 `buf` 工具链、生成产物与 README 更新要求。
  - [x] SubTask 4.1: 冻结 `buf build / lint / generate / breaking` 必须继续通过既有受控入口运行
  - [x] SubTask 4.2: 冻结 `buf breaking --against '../.git#branch=main,subdir=proto'` 为唯一 breaking 基准路径
  - [x] SubTask 4.3: 冻结 Go 与 TypeScript 生成产物新增落点为 `backend/internal/gen/proto/psco/dashboard/v1/` 与 `frontend/src/gen/proto/psco/dashboard/v1/`
  - [x] SubTask 4.4: 冻结 `proto/README.md` 必须把 Dashboard 纳入目录总览、生成产物表与 RPC → HTTP 映射矩阵

- [x] Task 5: 冻结 Dashboard 过渡传输层与 `.proto` 的单向映射边界。
  - [x] SubTask 5.1: 冻结 `backend/internal/dashboard/types.go`、handler DTO 与前端 adapter 只允许从 `.proto` 派生或显式对齐
  - [x] SubTask 5.2: 冻结三个 Dashboard `GET` 入口必须显式组装空 Proto request，而不是绕过 request 合同
  - [x] SubTask 5.3: 冻结错误状态码、错误包络与局部错误展示仍属于 HTTP / handler 适配层，不进入 `.proto` 本体
  - [x] SubTask 5.4: 冻结 DTO / 页面层不得新增 `.proto` 中不存在的业务字段语义

- [x] Task 6: 冻结 `phase05-11` 的最小化边界与阶段推进关系。
  - [x] SubTask 6.1: 冻结当前阶段只承接合同源落地、生成入口、校验链与映射边界
  - [x] SubTask 6.2: 冻结当前阶段继续允许 `chi + JSON HTTP` 作为过渡传输层
  - [x] SubTask 6.3: 冻结当前阶段不要求立即完成 gRPC / Connect 迁移或全量 DTO 替换
  - [x] SubTask 6.4: 冻结 `phase05-12 / 13 / 14` 必须优先引用已落地合同主线

- [x] Task 7: 完成 `phase05-11` 规格一致性校验。
  - [x] SubTask 7.1: 验证本 spec 已把 `.proto` 收口为仓库内唯一合同源
  - [x] SubTask 7.2: 验证本 spec 已覆盖 `buf` 工具链入口、breaking 基准路径与生成产物落点
  - [x] SubTask 7.3: 验证本 spec 已明确 `proto/README.md` 的总览更新要求
  - [x] SubTask 7.4: 验证本 spec 已明确 DTO/HTTP 适配层与 `.proto` 的单向映射边界
  - [x] SubTask 7.5: 验证本 spec 与 `phase05-08 / 10` 及既有 `phase03-11 / 04-11` 主线模式一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 3`, `Task 4`
- `Task 6` depends on `Task 3`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 1` through `Task 6`
