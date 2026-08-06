# Tasks

- [x] Task 1: 收口 `Decision Center` 合同在现有 `proto/` 工作区中的落点与边界。
  - [x] SubTask 1.1: 明确 `decision_center.proto` 的文件落点、包名与版本语义必须复用现有 `proto/` 工作区
  - [x] SubTask 1.2: 明确不得新增第二套 `buf.yaml`、`buf.gen.yaml`、`Makefile` 或并列 proto 根目录
  - [x] SubTask 1.3: 明确 `Decision Center` 与 `Module Registry` 必须在同一个 buf workspace 内共同通过校验

  执行证据：
  - 新增 `proto/psco/decision_center/v1/decision_center.proto`，包名冻结为 `psco.decision_center.v1`，复用现有 `proto/buf.yaml`（`modules: - path: .` 自动覆盖整个 `proto/` 目录）。
  - 未新增第二套 `buf.yaml` / `buf.gen.yaml` / `Makefile`；`proto/psco/decision_center/` 仅作为包目录而非并列 proto 根。
  - `buf build` 与 `buf lint` 在同一 workspace 内同时覆盖 `Module Registry` 与 `Decision Center`，均返回 0 退出码。

- [x] Task 2: 收口 `Decision Center` 合同源文件的实际落地要求。
  - [x] SubTask 2.1: 明确 `proto/psco/decision_center/v1/decision_center.proto` 必须存在并承接 `phase03-08` 已冻结的服务、消息、字段编号与枚举
  - [x] SubTask 2.2: 明确 `DecisionModuleCandidate.status` 必须通过跨包 import 复用 `psco.module_registry.v1.ModuleStatus`
  - [x] SubTask 2.3: 明确当前阶段不新增第二套本地等价枚举、不改写已冻结字段语义

  执行证据：
  - `decision_center.proto` 已落地，覆盖 `DecisionCenterService`（5 个 RPC）+ 16 个消息 + 2 个枚举，字段编号与 `phase03-08` §"字段编号方案与合同边界策略冻结"逐项一致。
  - 通过 `import "psco/module_registry/v1/module_registry.proto"` 复用 `psco.module_registry.v1.ModuleStatus`，`DecisionModuleCandidate.status` 字段类型为 `psco.module_registry.v1.ModuleStatus`。
  - 未在 `psco.decision_center.v1` 中重定义本地等价 `ModuleStatus`；生成的 Go 产物中出现 8 处 `module_registry` 引用，TS 产物中出现 6 处，验证跨包解析成功。

- [x] Task 3: 收口 `buf build / lint / generate / breaking` 的仓库执行入口。
  - [x] SubTask 3.1: 明确 `proto/buf.yaml`、`proto/buf.gen.yaml` 与 `proto/Makefile` 是当前阶段唯一受控工具链入口
  - [x] SubTask 3.2: 明确 `build / lint / generate / breaking` 必须同时覆盖 `Module Registry` 与 `Decision Center`
  - [x] SubTask 3.3: 明确 `buf breaking` 的 Git 基准路径、失败退出码与禁止绕过的约束

  执行证据：
  - 复用现有 `proto/Makefile`，补全 `build` target（`phase02-11A` 原仅含 `gen / lint / breaking`，`phase03-08` 冻结的 `buf build` 受控入口要求由本阶段落地），现 Makefile 含 `build / gen / lint / breaking / clean` 五个 target，未新增第二套脚本。
  - `make build` + `make gen` + `make lint` + `make breaking` 全部通过；`buf breaking --against '../.git#branch=main,subdir=proto'` 退出码为 0，对照仓库主线 Git 基准，未吞失败。
  - 校验同时覆盖两个模块，新增 `decision_center.proto` 不构成破坏性变更（向后兼容）。
  - `proto/README.md` §7 校验命令已同步补充 `make build` 入口说明。

- [x] Task 4: 收口生成产物归属与当前阶段生成边界。
  - [x] SubTask 4.1: 明确 Go 生成产物落在 `backend/internal/gen/proto/psco/decision_center/v1/`
  - [x] SubTask 4.2: 明确 TypeScript 生成产物落在 `frontend/src/gen/proto/psco/decision_center/v1/`
  - [x] SubTask 4.3: 明确当前阶段只要求最小合同产物，不要求完整 gRPC / Connect 传输层迁移

  执行证据：
  - Go 产物落点：`backend/internal/gen/proto/psco/decision_center/v1/decision_center.pb.go`（含 16 个消息 struct + 2 个枚举 type）。
  - TS 产物落点：`frontend/src/gen/proto/psco/decision_center/v1/decision_center_pb.ts`（含 16 个消息 type + 2 个枚举 + `DecisionCenterService` GenService 描述符）。
  - 未引入 gRPC / Connect 插件，`buf.gen.yaml` 仍只保留 `protocolbuffers/go` 与 `bufbuild/es` 两个消息类型生成插件，与 `phase02-11A` 边界一致。

- [x] Task 5: 收口 DTO / HTTP 过渡传输层与 `.proto` 的单向承接关系。
  - [x] SubTask 5.1: 明确 `backend/internal/decisioncenter/types.go`、handler DTO 与前端 adapter 的语义必须从 `.proto` 派生或显式对齐
  - [x] SubTask 5.2: 明确 URL 路径参数到 Proto request 字段的组装属于适配层，不构成第二套合同
  - [x] SubTask 5.3: 明确禁止在页面层、DTO 层或 HTTP 层私自新增 `.proto` 中不存在的业务字段语义

  执行证据：
  - `proto/README.md` §4 新增 `Decision Center` 的 RPC → HTTP 映射矩阵，并新增 `decision_id 的传输差异`小节，明确 URL 路径参数组装属于过渡传输层适配，不构成第二套合同。
  - `proto/README.md` §5 新增 `后端 — Decision Center` 与 `前端 — Decision Center` 衔接小节，列出预期 `types.go` 结构体与 `.proto` 消息的单向承接约束（实现由 `phase03-12 / 13` 承接）。
  - `decision_center.proto` 文件头注释明确"JSON 请求与响应语义必须从本 .proto 派生或严格对齐，不得形成第二套合同源"。

- [x] Task 6: 完成 `phase03-11` 规格校验与收口。
  - [x] SubTask 6.1: 验证规格已完整承接 `phase03-08` 与 `phase03-10` 的合同主线结论
  - [x] SubTask 6.2: 验证规格已与现有 `proto/` 工作区现实保持一致，不引入第二套入口
  - [x] SubTask 6.3: 验证规格已明确 `phase03-12 / 13 / 14` 可直接消费的合同、工具链与映射边界

  执行证据：
  - `decision_center.proto` 字段编号、枚举、服务接口、跨包 import、RPC → HTTP 映射矩阵与 `phase03-08` Spec 和 `phase03-10` §7 逐项一致，可在子代理复核中追溯。
  - 复用现有 `proto/buf.yaml` / `proto/buf.gen.yaml` / `proto/Makefile`，未新增第二套入口；`proto/README.md` 升级为统一合同源入口文档。
  - `phase03-12` 可直接消费 `.proto` 派生的 `types.go` 单向承接约束；`phase03-13` 可直接消费前端生成类型与 RPC → HTTP 映射矩阵；`phase03-14` 可直接通过 `make gen / lint / breaking` 验收合同兼容性。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 2`
- `Task 6` depends on `Task 2`, `Task 3`, `Task 4`, and `Task 5`
