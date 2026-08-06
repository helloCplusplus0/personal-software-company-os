# PSCO Proto 合同源

> **文档定位**：本目录是 PSCO 项目的 Protocol Buffers 合同源入口，统一承接 `Module Registry`（phase02-11A）与 `Decision Center`（phase03-11）的最小合同源。
> 上游规格：
> - `phase02-11A spec` §"Module Registry 最小 Proto 合同源" + `module_registry_spec_v0.1.md` §6.1 / §6.5
> - `phase03-08 Decision Center 最小 Proto 合同设计 Spec` + `phase03-10 decision_center_spec_v0.1.md` §7
> - `phase03-11 Decision Center 最小 Protocol Buffers 合同主线 Spec`

## 1. 目录结构

```
proto/
├── buf.yaml              # buf 模块配置（lint + breaking 规则），单一 workspace
├── buf.gen.yaml          # 代码生成配置（Go + TypeScript），单一生成入口
├── Makefile              # 生成与校验入口（build / gen / lint / breaking / clean）
├── README.md             # 本文件
└── psco/
    ├── module_registry/
    │   └── v1/
    │       └── module_registry.proto    # Module Registry 最小合同源
    └── decision_center/
        └── v1/
            └── decision_center.proto    # Decision Center 最小合同源（phase03-11 落地）
```

> 约束：不得为单个模块新增第二套 `buf.yaml`、`buf.gen.yaml`、`Makefile` 或并列 proto 根目录。`Module Registry` 与 `Decision Center` 必须在同一个 buf workspace 内共同通过 `build / lint / generate / breaking` 校验。

## 2. 包名与版本语义

| 模块 | 包名 | 版本号 | 文件落点 |
| --- | --- | --- | --- |
| Module Registry | `psco.module_registry.v1` | `v1` | `psco/module_registry/v1/module_registry.proto` |
| Decision Center | `psco.decision_center.v1` | `v1` | `psco/decision_center/v1/decision_center.proto` |

### 演进规则（统一适用于所有模块）

- 新增字段必须使用新的递增编号，不得复用已删除字段编号
- 破坏性变更（删除字段、修改字段类型）必须升级版本号（`v2`）
- 后续版本删除字段或废弃字段名时必须使用 `reserved` 保留字段号，必要时同时保留字段名
- `buf breaking` 用于 CI 中自动检测破坏性变更

### 跨包依赖

- `Decision Center` 通过 `import "psco/module_registry/v1/module_registry.proto"` 直接复用 `psco.module_registry.v1.ModuleStatus`，用于 `DecisionModuleCandidate.status`
- 不得在 `psco.decision_center.v1` 中重定义本地等价枚举
- 该 import 仅限于复用 `ModuleStatus` 枚举类型，不引入对 `Module Registry` service 或其他消息结构的合同依赖

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

| 语言 | 模块 | 落点 | 说明 |
| --- | --- | --- | --- |
| Go | Module Registry | `backend/internal/gen/proto/psco/module_registry/v1/` | 消息类型（service 定义已冻结在 .proto 中，后续迁移时再加回 grpc/connect 插件） |
| Go | Decision Center | `backend/internal/gen/proto/psco/decision_center/v1/` | 消息类型（同上） |
| TypeScript | Module Registry | `frontend/src/gen/proto/psco/module_registry/v1/` | 消息类型（同上） |
| TypeScript | Decision Center | `frontend/src/gen/proto/psco/decision_center/v1/` | 消息类型（同上） |

> 生成产物已加入 `.gitignore`，不进入版本控制。每次 `make gen` 重新生成。
> 当前阶段所有模块只生成消息类型与 service 描述符，不生成 gRPC 服务桩或客户端桩。后续迁移到 gRPC / Connect 时在 `buf.gen.yaml` 中加回对应插件，不需要修改 `.proto` 合同源。

## 4. 过渡传输层映射

当前阶段保留 `chi + JSON HTTP` 作为过渡传输层（`phase02-11` 已实现；`Decision Center` 由 `phase03-12 / 13` 实现）。
`.proto` 是唯一合同源，JSON 请求与响应语义从 `.proto` 派生。

### RPC → HTTP 映射矩阵

#### Module Registry

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

#### Decision Center

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `ListDecisions` | GET | `/api/decisions` | query params |
| `GetDecisionDetail` | GET | `/api/decisions/{decisionId}` | URL path |
| `CreateDecision` | POST | `/api/decisions` | JSON body |
| `LinkDecisionToTarget` | POST | `/api/decisions/{decisionId}/links` | JSON body + URL path |
| `ListDecisionModuleCandidates` | GET | `/api/decisions/{decisionId}/candidates/modules` | URL path |

> 约束：JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 语义显式对齐；HTTP 过渡层使用 URL 路径参数承接 `decisionId` 或 `moduleId` 时，handler 必须在进入业务层前显式组装为对应的 Proto request 字段；不得把 HTTP 路径、状态码或中间件策略误写成 Proto 合同本体。

### 字段映射约定

| Proto 字段类型 | JSON 序列化 | 说明 |
| --- | --- | --- |
| `string` | string | 直接映射 |
| `int32` | number | 直接映射 |
| `optional string` | string \| null | nullable 字段 |
| `enum` | string | 使用枚举名的小写形式（如 `"active"` / `"proposed"`） |
| `google.protobuf.Timestamp` | RFC3339 string | 标准 JSON 映射 |
| `google.protobuf.Empty` | 无请求体 / 空响应 | 绑定/映射动作返回 204 |
| `repeated T` | array | 直接映射（如 `repeated string alternatives` 映射为 `string[]`） |

### 路径参数的传输差异

#### module_id 的传输差异

- **Proto RPC**：`module_id` 作为请求消息的显式字段（如 `CreateReleaseRequest.module_id`）
- **HTTP 过渡层**：`module_id` 由 URL 路径参数承接，不放在 JSON 请求体

这是传输层差异，不是合同差异。HTTP handler 从 URL 提取 `module_id` 后组装为 Proto 请求。

#### decision_id 的传输差异

- **Proto RPC**：`decision_id` 作为请求消息的显式字段（如 `GetDecisionDetailRequest.decision_id` / `LinkDecisionToTargetRequest.decision_id`）
- **HTTP 过渡层**：`decision_id` 由 URL 路径参数承接，不放在 JSON 请求体

这是传输层差异，不是合同差异。HTTP handler 从 URL 提取 `decision_id` 后组装为 Proto 请求。

## 5. 与现有实现的衔接关系

> `.proto` 是唯一合同源。`types.go` / `types.ts` 是过渡传输层的 HTTP DTO，通过显式映射与 `.proto` 保持语义一致。
> 两者不是"字段严格一致"——存在命名约定差异（snake_case vs PascalCase vs camelCase）、时间表示差异（Timestamp vs time.Time vs RFC3339 string）和传输层字段裁剪差异（路径参数在 HTTP 层由 URL 承接）。
> 禁止在页面层、DTO 层或 HTTP 层私自新增 `.proto` 中不存在的业务字段语义。

### 后端 — Module Registry（`backend/internal/moduleregistry/types.go`）

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

### 前端 — Module Registry（`frontend/src/features/module-registry/types.ts`）

当前 `types.ts` 中的手写 TypeScript 接口与 `.proto` 消息语义对齐。
命名使用 camelCase（如 `moduleId` / `releasedAt`），与 `.proto` 的 snake_case 通过显式映射转换。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

### 后端 — Decision Center（`backend/internal/decisioncenter/types.go`，由 phase03-12 实现）

`Decision Center` 的 `.proto` 合同源已在 `phase03-11` 落地，后端过渡传输层 DTO 由 `phase03-12` 实现。
实现时必须遵守的单向承接约束：

| 预期 `types.go` 结构体 | `.proto` 消息 | 单向承接约束 |
| --- | --- | --- |
| `Decision` | `Decision` | 语义对齐（created_at: time.Time ↔ Timestamp；alternatives: []string ↔ repeated string） |
| `DecisionListItem` | `DecisionListItem` | 语义对齐 |
| `LinkedModule` | `LinkedModule` | 语义对齐 |
| `SourceContext` | `SourceContext` | 语义对齐 |
| `DecisionDetail` | `DecisionDetail` | 语义对齐 |
| `DecisionModuleCandidate` | `DecisionModuleCandidate` | 语义对齐（status 复用 ModuleStatus，不重定义本地枚举） |
| `CreateDecisionRequest` | `CreateDecisionRequest` | 语义对齐 |
| `LinkDecisionToTargetRequest` | `LinkDecisionToTargetRequest` | 语义对齐（types.go 无 decision_id，由 URL 承接） |

> 实现期不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义。

### 前端 — Decision Center（`frontend/src/features/decision-center/`，由 phase03-13 实现）

`Decision Center` 的前端过渡传输层适配代码由 `phase03-13` 实现。
命名使用 camelCase（如 `decisionId` / `createdAt` / `linkedModuleSummary`），与 `.proto` 的 snake_case 通过显式映射转换。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

## 6. 当前阶段不做

- 不完成完整 gRPC / Connect 传输层迁移
- 不替换 `chi + JSON HTTP` 为 gRPC 服务器
- 不将生成代码集成到现有 handler / adapter 中
- 不引入 gRPC 网关或连接池
- 以上属于后续 phase 的范围

## 7. 校验命令

```bash
# 编译 proto 文件为 buf image，验证合同源可被正确解析与编译
make build

# 生成 Go + TypeScript 合同产物
make gen

# 规范校验
make lint

# 破坏性变更检测（需要 main 分支已有 .proto 基准）
make breaking
```

> `make build` 是 buf 校验链的最基础入口，`lint / generate / breaking` 都隐式依赖 build 成功。
> `make breaking` 对照仓库 `main` 分支的 `proto/` 子目录基准。
> 当前 `main` 分支已有 `phase02-11A` 提交的 `Module Registry` proto 基准，`phase03-11` 新增的 `Decision Center` proto 属于向后兼容的新增，`make breaking` 退出码为 0。
> 后续任何对已存在字段编号、字段类型或字段语义的删除或修改都会被检测为破坏性变更并返回非零退出码。
