# Tasks

- [x] Task 1: 收口 `Product / Repository / Binding` 合同在现有 `proto/` workspace 中的落点与边界。
  - [x] SubTask 1.1: 明确 `common.proto`、`product_registry.proto`、`repository_binding.proto` 的文件落点、包名与版本语义必须复用现有 `proto/` workspace
  - [x] SubTask 1.2: 明确不得新增第二套 `buf.yaml`、`buf.gen.yaml`、`Makefile` 或并列 proto 根目录
  - [x] SubTask 1.3: 明确 `Product / Repository / Binding` 必须与 `Module Registry`、`Decision Center` 在同一个 buf workspace 内共同通过校验

  执行证据：
  - 新增 `proto/psco/common/v1/common.proto`（包名 `psco.common.v1`）、`proto/psco/product_registry/v1/product_registry.proto`（包名 `psco.product_registry.v1`）、`proto/psco/repository_binding/v1/repository_binding.proto`（包名 `psco.repository_binding.v1`），均复用现有 `proto/buf.yaml`（`modules: - path: .` 自动覆盖整个 `proto/` 目录）。
  - 未新增第二套 `buf.yaml` / `buf.gen.yaml` / `Makefile`；`proto/psco/common/`、`proto/psco/product_registry/`、`proto/psco/repository_binding/` 仅作为包目录而非并列 proto 根。
  - `buf build` 与 `buf lint` 在同一 workspace 内同时覆盖 `Module Registry`、`Decision Center`、`Common`、`Product Registry` 与 `Repository Binding`，均返回 0 退出码。

- [x] Task 2: 收口 `Product / Repository / Binding` 合同源文件的实际落地要求。
  - [x] SubTask 2.1: 明确 `proto/psco/common/v1/common.proto`、`proto/psco/product_registry/v1/product_registry.proto`、`proto/psco/repository_binding/v1/repository_binding.proto` 必须存在并承接 `phase04-08` 已冻结的服务、消息、字段编号与枚举
  - [x] SubTask 2.2: 明确 `ActiveArchivedStatus` 必须在 `psco.common.v1` 中单值定义
  - [x] SubTask 2.3: 明确 `ModuleStatus` 必须通过跨包 import 复用 `psco.module_registry.v1.ModuleStatus`
  - [x] SubTask 2.4: 明确当前阶段不新增第二套本地等价枚举、不改写已冻结字段语义

  执行证据：
  - `common.proto` 已落地，定义单一枚举 `ActiveArchivedStatus`（3 个枚举值：UNSPECIFIED=0 / ACTIVE=1 / ARCHIVED=2），承接 `phase04-08` §"公共 active / archived 状态枚举必须单值化"。
  - `product_registry.proto` 已落地，覆盖 `ProductRegistryService`（5 个 RPC）+ 16 个消息（Product / ProductListItem / BoundModuleSummary / BoundRepositorySummary / ProductDetail / ProductModuleCandidate + 5 组 request/response 共 10 个消息），字段编号与 `phase04-08` §"字段编号方案必须在当前阶段冻结"逐项一致。
  - `repository_binding.proto` 已落地，覆盖 `RepositoryBindingService`（7 个 RPC）+ 21 个消息（Repository / RepositoryListItem / BoundProductSummary / MappedModuleSummary / RepositoryDetail / RepositoryProductCandidate / RepositoryModuleCandidate + 7 组 request/response 共 14 个消息），字段编号与 `phase04-08` 逐项一致。
  - 通过 `import "psco/module_registry/v1/module_registry.proto"` 复用 `psco.module_registry.v1.ModuleStatus`，`BoundModuleSummary.module_status` / `MappedModuleSummary.module_status` / `ProductModuleCandidate.module_status` / `RepositoryModuleCandidate.module_status` 字段类型均为 `psco.module_registry.v1.ModuleStatus`。
  - 通过 `import "psco/common/v1/common.proto"` 复用 `psco.common.v1.ActiveArchivedStatus`，Product / Repository 相关 status 字段均复用该枚举。
  - 未在 `psco.product_registry.v1` 或 `psco.repository_binding.v1` 中重定义本地等价枚举；生成的 Go 产物中 `product_registry.pb.go` 出现 12 处 `module_registry` 引用、20 处 `common` 引用；`repository_binding.pb.go` 出现 12 处 `module_registry` 引用、24 处 `common` 引用；TS 产物中各出现 8 处 `module_registry` 引用，验证跨包解析成功。

- [x] Task 3: 收口 `buf build / lint / generate / breaking` 的仓库执行入口。
  - [x] SubTask 3.1: 明确 `proto/buf.yaml`、`proto/buf.gen.yaml` 与 `proto/Makefile` 是当前阶段唯一受控工具链入口
  - [x] SubTask 3.2: 明确 `build / lint / generate / breaking` 必须同时覆盖 `Module Registry`、`Decision Center`、`Product Registry` 与 `Repository Binding`
  - [x] SubTask 3.3: 明确 `buf breaking` 的 Git 基准路径、失败退出码与禁止绕过的约束

  执行证据：
  - 复用现有 `proto/Makefile`（含 `build / gen / lint / breaking / clean` 五个 target），Makefile 头部已补充 `phase04-08` 与 `phase04-11` 上游规格引用，未新增第二套脚本。
  - `make build` + `make gen` + `make lint` + `make breaking` 全部通过；`buf breaking --against '../.git#branch=main,subdir=proto'` 退出码为 0，对照仓库主线 Git 基准，未吞失败。
  - 校验同时覆盖五个模块（Common / Module Registry / Decision Center / Product Registry / Repository Binding），新增三个 `.proto` 文件属于向后兼容新增，不构成 breaking 变更。
  - `proto/README.md` §7 校验命令说明已同步补充 `phase04-11` 新增 proto 的 breaking 兼容性说明。

- [x] Task 4: 收口生成产物归属与当前阶段生成边界。
  - [x] SubTask 4.1: 明确 Go 生成产物落在 `backend/internal/gen/proto/psco/common/v1/`、`product_registry/v1/`、`repository_binding/v1/`
  - [x] SubTask 4.2: 明确 TypeScript 生成产物落在 `frontend/src/gen/proto/psco/common/v1/`、`product_registry/v1/`、`repository_binding/v1/`
  - [x] SubTask 4.3: 明确当前阶段只要求最小合同产物，不要求完整 gRPC / Connect 传输层迁移

  执行证据：
  - Go 产物落点：
    - `backend/internal/gen/proto/psco/common/v1/common.pb.go`（含 `ActiveArchivedStatus` 枚举 type）
    - `backend/internal/gen/proto/psco/product_registry/v1/product_registry.pb.go`（含 16 个消息 struct + `ProductRegistryService` GenService 描述符）
    - `backend/internal/gen/proto/psco/repository_binding/v1/repository_binding.pb.go`（含 21 个消息 struct + `RepositoryBindingService` GenService 描述符）
  - TS 产物落点：
    - `frontend/src/gen/proto/psco/common/v1/common_pb.ts`（含 `ActiveArchivedStatus` 枚举）
    - `frontend/src/gen/proto/psco/product_registry/v1/product_registry_pb.ts`（含 34 个 exported types/consts + `ProductRegistryService` GenService 描述符）
    - `frontend/src/gen/proto/psco/repository_binding/v1/repository_binding_pb.ts`（含 44 个 exported types/consts + `RepositoryBindingService` GenService 描述符）
  - 未引入 gRPC / Connect 插件，`buf.gen.yaml` 仍只保留 `protocolbuffers/go` 与 `bufbuild/es` 两个消息类型生成插件，与 `phase02-11A` / `phase03-11` 边界一致。

- [x] Task 5: 收口 DTO / HTTP 过渡传输层与 `.proto` 的单向承接关系。
  - [x] SubTask 5.1: 明确 `backend/internal/productregistry/types.go`、`backend/internal/repositorybinding/types.go`、handler DTO 与前端 adapter 的语义必须从 `.proto` 派生或显式对齐
  - [x] SubTask 5.2: 明确 URL 路径参数到 Proto request 字段的组装属于适配层，不构成第二套合同
  - [x] SubTask 5.3: 明确禁止在页面层、DTO 层或 HTTP 层私自新增 `.proto` 中不存在的业务字段语义

  执行证据：
  - `proto/README.md` §4 新增 `Product Registry` 与 `Repository Binding` 的 RPC → HTTP 映射矩阵（共 12 条映射），并新增 `product_id 的传输差异` 与 `repository_id 的传输差异` 小节，明确 URL 路径参数组装属于过渡传输层适配，不构成第二套合同。
  - `proto/README.md` §5 新增 `后端 — Product Registry`、`前端 — Product Registry`、`后端 — Repository Binding`、`前端 — Repository Binding` 衔接小节，列出预期 `types.go` 结构体与 `.proto` 消息的单向承接约束（实现由 `phase04-12 / 13` 承接）。
  - `common.proto`、`product_registry.proto`、`repository_binding.proto` 文件头注释均明确"JSON 请求与响应语义必须从本 .proto 派生或严格对齐，不得形成第二套合同源"。

- [x] Task 6: 完成 `phase04-11` 规格校验与收口。
  - [x] SubTask 6.1: 验证规格已完整承接 `phase04-08` 与 `phase04-10` 的合同主线结论
  - [x] SubTask 6.2: 验证规格已与现有 `proto/` workspace 现实保持一致，不引入第二套入口
  - [x] SubTask 6.3: 验证规格已明确 `phase04-12 / 13` 可直接消费的合同、工具链与映射边界

  执行证据：
  - `common.proto` + `product_registry.proto` + `repository_binding.proto` 的包名、字段编号、枚举、服务接口、跨包 import、RPC → HTTP 映射矩阵与 `phase04-08` Spec 和 `phase04-10` §合同设计逐项一致，可在子代理复核中追溯。
  - 复用现有 `proto/buf.yaml` / `proto/buf.gen.yaml` / `proto/Makefile`，未新增第二套入口；`proto/README.md` 升级为统一合同源入口文档（覆盖五个模块）。
  - `phase04-12` 可直接消费 `.proto` 派生的 `types.go` 单向承接约束（§5 衔接表）；`phase04-13` 可直接消费前端生成类型与 RPC → HTTP 映射矩阵（§4 映射表）；`phase04-14` 可直接通过 `make gen / lint / breaking` 验收合同兼容性。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 2`
- `Task 6` depends on `Task 2`, `Task 3`, `Task 4`, and `Task 5`
