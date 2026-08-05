# PSCO Proto 合同源

> **文档定位**：本目录是 PSCO 项目 `Module Registry` 的 Protocol Buffers 合同源入口。
> 上游规格：`phase02-11A spec` §"Module Registry 最小 Proto 合同源" + `module_registry_spec_v0.1.md` §6.1 / §6.5

## 1. 目录结构

```
proto/
├── buf.yaml              # buf 模块配置（lint + breaking 规则）
├── buf.gen.yaml          # 代码生成配置（Go + TypeScript）
├── Makefile              # 生成与校验入口
├── README.md             # 本文件
└── psco/
    └── module_registry/
        └── v1/
            └── module_registry.proto   # Module Registry 最小合同源
```

## 2. 包名与版本语义

- **包名**：`psco.module_registry.v1`
- **版本号**：`v1`
- **演进规则**：
  - 新增字段必须使用新的递增编号，不得复用已删除字段编号
  - 破坏性变更（删除字段、修改字段类型）必须升级版本号（`v2`）
  - `buf breaking` 用于 CI 中自动检测破坏性变更

## 3. 生成入口

### 前置条件

- `buf` CLI（已安装，版本 ≥ 1.55）
- 无需本地安装 protoc 插件，buf 通过远程插件自动下载

### 生成命令

```bash
# 在 proto/ 目录执行
make gen

# 或直接使用 buf
buf generate
```

### 生成产物落点

| 语言 | 落点 | 说明 |
| --- | --- | --- |
| Go | `backend/internal/gen/proto/psco/module_registry/v1/` | 消息类型（service 定义已冻结在 .proto 中，后续迁移时再加回 grpc/connect 插件） |
| TypeScript | `frontend/src/gen/proto/psco/module_registry/v1/` | 消息类型（同上） |

> 生成产物已加入 `.gitignore`，不进入版本控制。每次 `make gen` 重新生成。

## 4. 过渡传输层映射

当前阶段保留 `chi + JSON HTTP` 作为过渡传输层（`phase02-11` 已实现）。
`.proto` 是唯一合同源，JSON 请求与响应语义从 `.proto` 派生。

### RPC → HTTP 映射矩阵

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `ListModules` | GET | `/api/modules` | query params |
| `GetModuleDetail` | GET | `/api/modules/{moduleId}` | URL path |
| `CreateModule` | POST | `/api/modules` | JSON body |
| `CreateRelease` | POST | `/api/modules/{moduleId}/releases` | JSON body + URL path |
| `BindModuleToProduct` | POST | `/api/modules/{moduleId}/bindings/products` | JSON body + URL path |
| `MapModuleToRepository` | POST | `/api/modules/{moduleId}/bindings/repositories` | JSON body + URL path |
| `ListProductCandidates` | GET | `/api/candidates/products` | — |
| `ListRepositoryCandidates` | GET | `/api/candidates/repositories` | — |

### 字段映射约定

| Proto 字段类型 | JSON 序列化 | 说明 |
| --- | --- | --- |
| `string` | string | 直接映射 |
| `int32` | number | 直接映射 |
| `optional string` | string \| null | nullable 字段 |
| `enum` | string | 使用枚举名的小写形式（如 `"active"`） |
| `google.protobuf.Timestamp` | RFC3339 string | 标准 JSON 映射 |
| `google.protobuf.Empty` | 无请求体 / 空响应 | 绑定/映射动作返回 204 |
| `repeated T` | array | 直接映射 |

### module_id 的传输差异

- **Proto RPC**：`module_id` 作为请求消息的显式字段（如 `CreateReleaseRequest.module_id`）
- **HTTP 过渡层**：`module_id` 由 URL 路径参数承接，不放在 JSON 请求体

这是传输层差异，不是合同差异。HTTP handler 从 URL 提取 `module_id` 后组装为 Proto 请求。

## 5. 与现有实现的衔接关系

> `.proto` 是唯一合同源。`types.go` / `types.ts` 是过渡传输层的 HTTP DTO，通过显式映射与 `.proto` 保持语义一致。
> 两者不是"字段严格一致"——存在命名约定差异（snake_case vs PascalCase vs camelCase）、时间表示差异（Timestamp vs time.Time vs RFC3339 string）和传输层字段裁剪差异（module_id 在 HTTP 层由 URL 承接）。

### 后端（`backend/internal/moduleregistry/types.go`）

当前 `types.go` 中的手写 Go 结构体是过渡传输层的 JSON DTO，与 `.proto` 消息语义对齐：

| `types.go` 结构体 | `.proto` 消息 | 对齐状态 |
| --- | --- | --- |
| `Module` | `Module` | 语义对齐（created_at: time.Time ↔ Timestamp） |
| `Release` | `Release` | 语义对齐（released_at: time.Time ↔ Timestamp） |
| `ModuleListItem` | `ModuleListItem` | 语义对齐（latest_release: *string ↔ optional string） |
| `ModuleDetail` | `ModuleDetail` | 语义对齐 |
| `ProductBinding` | `ProductBinding` | 语义对齐 |
| `RepositoryMapping` | `RepositoryMapping` | 语义对齐 |
| `DecisionLink` | `DecisionLink` | 语义对齐 |
| `ProductCandidate` | `ProductCandidate` | 语义对齐 |
| `RepositoryCandidate` | `RepositoryCandidate` | 语义对齐 |
| `CreateModuleRequest` | `CreateModuleRequest` | 语义对齐 |
| `CreateReleaseRequest` | `CreateReleaseRequest` | 语义对齐（types.go 无 module_id，由 URL 承接；released_at: string ↔ Timestamp） |
| `BindModuleToProductRequest` | `BindModuleToProductRequest` | 语义对齐（types.go 无 module_id，由 URL 承接） |
| `MapModuleToRepositoryRequest` | `MapModuleToRepositoryRequest` | 语义对齐（types.go 无 module_id，由 URL 承接） |

> 后续阶段可逐步将 `types.go` 替换为从 `.proto` 生成的 Go 类型，但当前阶段不做此迁移。

### 前端（`frontend/src/features/module-registry/types.ts`）

当前 `types.ts` 中的手写 TypeScript 接口与 `.proto` 消息语义对齐。
命名使用 camelCase（如 `moduleId` / `releasedAt`），与 `.proto` 的 snake_case 通过显式映射转换。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

## 6. 当前阶段不做

- 不完成完整 gRPC / Connect 传输层迁移
- 不替换 `chi + JSON HTTP` 为 gRPC 服务器
- 不将生成代码集成到现有 handler / adapter 中
- 不引入 gRPC 网关或连接池
- 以上属于后续 phase 的范围

## 7. 校验命令

```bash
# 生成
make gen

# 规范校验
make lint

# 破坏性变更检测（需要 main 分支已有 .proto 基准）
make breaking
```

> `make breaking` 在首次提交前会失败（main 分支无 .proto 基准），这是预期行为。
> 提交后 CI 中运行将正确检测字段编号与语义的破坏性变更。
