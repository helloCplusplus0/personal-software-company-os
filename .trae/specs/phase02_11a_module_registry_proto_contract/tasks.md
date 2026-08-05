# Tasks

- [x] Task 1: 冻结 Proto 合同源落点与版本语义。
  - [x] SubTask 1.1: 冻结 `Module Registry` Proto 文件目录、文件名、package 与版本号
  - [x] SubTask 1.2: 明确 `.proto` 作为当前阶段唯一合同源，手写 JSON 结构不得再并列扩张
  - [x] SubTask 1.3: 明确当前阶段只要求最小合同落地，不要求立即完成完整传输层迁移

- [x] Task 2: 冻结核心消息结构与字段编号。
  - [x] SubTask 2.1: 定义 `Module`、`Release`、`ModuleListItem`、`ModuleDetail` 等核心消息
  - [x] SubTask 2.2: 定义产品绑定、仓库映射、`Decision` 附属读取与候选读取消息
  - [x] SubTask 2.3: 为字段分配稳定编号，并明确保留/删除的兼容性规则

- [x] Task 3: 冻结服务接口与请求响应语义。
  - [x] SubTask 3.1: 定义 `ModuleListRead` 与 `ModuleDetailRead` 的 request / response
  - [x] SubTask 3.2: 定义 `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite` 的 request / response
  - [x] SubTask 3.3: 定义 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead`
  - [x] SubTask 3.4: 明确 `Decision` 只内嵌在 `ModuleDetailRead` 中，不新增独立 Proto 服务

- [x] Task 4: 冻结 Proto 与过渡传输层的衔接规则。
  - [x] SubTask 4.1: 明确 `chi + JSON HTTP` 仅作为过渡传输层，不能形成第二套合同源
  - [x] SubTask 4.2: 明确后端 `types.go` / handler response 与前端 `api-adapter` 必须对齐 Proto 消息语义
  - [x] SubTask 4.3: 明确 HTTP 路径、状态码与中间件不属于 Proto 合同本体

- [x] Task 5: 冻结最小生成与验收前提。
  - [x] SubTask 5.1: 明确当前阶段是否需要生成 Go / TS 合同产物，至少冻结生成入口约定
  - [x] SubTask 5.2: 明确 `phase02-12` 联调验收需要如何基于 `.proto` 校对实现
  - [x] SubTask 5.3: 明确后续新增接口必须优先改 `.proto`，而不是先改 JSON 结构

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 3` and `Task 4`

# 实现产物索引

## Proto 合同源（proto/）
- `proto/psco/module_registry/v1/module_registry.proto` — Module Registry 最小 .proto 合同源（唯一合同定义入口）
- `proto/buf.yaml` — buf 模块配置（lint + breaking 规则）
- `proto/buf.gen.yaml` — 代码生成配置（Go 消息类型 + TypeScript 消息类型）
- `proto/Makefile` — 生成与校验入口（make gen / make lint / make breaking / make clean）
- `proto/README.md` — 合同源说明（包名、版本语义、生成入口、过渡传输层映射、与现有实现对齐关系）

## 生成产物（已 gitignore，由 make gen 重新生成）
- `backend/internal/gen/proto/psco/module_registry/v1/module_registry.pb.go` — Go 消息类型
- `frontend/src/gen/proto/psco/module_registry/v1/module_registry_pb.ts` — TypeScript 消息类型

## 依赖更新
- `backend/go.mod` — 新增 `google.golang.org/protobuf` 依赖（生成 Go 代码运行时）
- `frontend/package.json` — 新增 `@bufbuild/protobuf` 依赖（生成 TS 代码运行时）
- `frontend/tsconfig.app.json` — 排除 `src/gen` 目录（生成代码尚未集成到现有实现）

## .gitignore 更新
- 新增 `backend/internal/gen/` 和 `frontend/src/gen/` 排除规则

## 合同覆盖矩阵

| 动作语义 | 接口分组 | Proto RPC | Proto Request | Proto Response |
| --- | --- | --- | --- | --- |
| 列表读取 | 读组 | `ListModules` | `ListModulesRequest` | `ListModulesResponse` |
| 详情读取 | 读组 | `GetModuleDetail` | `GetModuleDetailRequest` | `GetModuleDetailResponse` |
| CreateModule | 写组 | `CreateModule` | `CreateModuleRequest` | `CreateModuleResponse` |
| CreateRelease | 写组 | `CreateRelease` | `CreateReleaseRequest` | `CreateReleaseResponse` |
| BindModuleToProduct | 写组 | `BindModuleToProduct` | `BindModuleToProductRequest` | `BindModuleToProductResponse` |
| MapModuleToRepository | 写组 | `MapModuleToRepository` | `MapModuleToRepositoryRequest` | `MapModuleToRepositoryResponse` |
| Product 候选读取 | 候选读取 | `ListProductCandidates` | `ListProductCandidatesRequest` | `ListProductCandidatesResponse` |
| Repository 候选读取 | 候选读取 | `ListRepositoryCandidates` | `ListRepositoryCandidatesRequest` | `ListRepositoryCandidatesResponse` |

## 运行时验证记录

- `buf lint` 通过（STANDARD 规则集）
- `buf generate` 成功生成 Go + TypeScript 产物
- `go build ./...` 通过（含生成代码）
- `go vet ./...` 通过
- `npm run build` 通过（生成代码已排除出 TS 编译范围）
- `npm run lint` 通过（0 errors，仅 shadcn/ui 预存警告）

## 生成策略说明

当前阶段只生成消息类型（`protocolbuffers/go` + `bufbuild/es`），不生成 gRPC 服务桩。
原因：
1. spec 明确"当前阶段可以不完成完整 gRPC / Connect / 网关迁移"
2. 当前传输层为 `chi + JSON HTTP`，不需要 gRPC 桩
3. service 定义已在 .proto 合同源中冻结，后续迁移到 gRPC 时加回 `grpc/go` 插件即可

## 后续新增接口约束

后续任何新增接口必须优先修改 `.proto` 合同源，再派生 JSON HTTP 实现。
不得先改 JSON 结构再回补 `.proto`，避免形成第二套合同源。
