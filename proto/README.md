# PSCO Proto 合同源

> **文档定位**：本目录是 PSCO 项目的 Protocol Buffers 合同源入口，统一承接 `Module Registry`（phase02-11A）、`Decision Center`（phase03-11）、`Product / Repository / Binding`（phase04-11）、`Dashboard + Feedback`（phase05-11）与 `Onboarding + Export + Backup + Reuse Summary`（phase06-13）的最小合同源。
> 上游规格：
> - `phase02-11A spec` §"Module Registry 最小 Proto 合同源" + `module_registry_spec_v0.1.md` §6.1 / §6.5
> - `phase03-08 Decision Center 最小 Proto 合同设计 Spec` + `phase03-10 decision_center_spec_v0.1.md` §7
> - `phase03-11 Decision Center 最小 Protocol Buffers 合同主线 Spec`
> - `phase04-08 Product / Repository / Binding 最小 Proto 合同设计 Spec` + `phase04-10 product_repository_binding_spec_v0.1.md` §合同设计
> - `phase04-11 Product / Repository / Binding 最小 Protocol Buffers 合同主线 Spec`
> - `phase05-08 Dashboard + Feedback 最小 Proto 合同设计 Spec` + `phase05-10 dashboard_feedback_spec_v0.1.md` §合同设计
> - `phase05-11 Dashboard + Feedback 最小 Protocol Buffers 合同主线 Spec`
> - `phase06-10 当前阶段最小 Protocol Buffers 合同设计 Spec` + `phase06-12 Onboarding + Sovereignty + Reuse 正式规格正文 Spec` §"合同、传输与演进基线冻结"
> - `phase06-13 落实当前阶段最小 Protocol Buffers 合同主线 Spec`

## 1. 目录结构

```
proto/
├── buf.yaml              # buf 模块配置（lint + breaking 规则），单一 workspace
├── buf.gen.yaml          # 代码生成配置（Go + TypeScript），单一生成入口
├── Makefile              # 生成与校验入口（build / gen / lint / breaking / clean）
├── README.md             # 本文件
└── psco/
    ├── common/
    │   └── v1/
    │       └── common.proto             # 跨 Product / Repository 共享最小公共枚举（phase04-11 落地）
    ├── module_registry/
    │   └── v1/
    │       └── module_registry.proto    # Module Registry 最小合同源
    ├── decision_center/
    │   └── v1/
    │       └── decision_center.proto    # Decision Center 最小合同源（phase03-11 落地）
    ├── product_registry/
    │   └── v1/
    │       └── product_registry.proto   # Product Registry 最小合同源（phase04-11 落地）
    ├── repository_binding/
    │   └── v1/
    │       └── repository_binding.proto # Repository Binding 最小合同源（phase04-11 落地）
    ├── dashboard/
    │   └── v1/
    │       └── dashboard.proto          # Dashboard + Feedback 最小合同源（phase05-11 落地）
    ├── onboarding/
    │   └── v1/
    │       └── onboarding.proto         # Onboarding 最小合同源（phase06-13 落地）
    ├── export/
    │   └── v1/
    │       └── export.proto             # Export 最小合同源（phase06-13 落地）
    ├── backup/
    │   └── v1/
    │       └── backup.proto             # Backup 最小合同源（phase06-13 落地）
    └── reuse_summary/
        └── v1/
            └── reuse_summary.proto      # Reuse Summary 最小合同源（phase06-13 落地）
```

> 约束：不得为单个模块新增第二套 `buf.yaml`、`buf.gen.yaml`、`Makefile` 或并列 proto 根目录。`Module Registry`、`Decision Center`、`Product Registry`、`Repository Binding`、`Dashboard + Feedback`、`Onboarding`、`Export`、`Backup` 与 `Reuse Summary` 必须在同一个 buf workspace 内共同通过 `build / lint / generate / breaking` 校验。

## 2. 包名与版本语义

| 模块 | 包名 | 版本号 | 文件落点 |
| --- | --- | --- | --- |
| Common | `psco.common.v1` | `v1` | `psco/common/v1/common.proto` |
| Module Registry | `psco.module_registry.v1` | `v1` | `psco/module_registry/v1/module_registry.proto` |
| Decision Center | `psco.decision_center.v1` | `v1` | `psco/decision_center/v1/decision_center.proto` |
| Product Registry | `psco.product_registry.v1` | `v1` | `psco/product_registry/v1/product_registry.proto` |
| Repository Binding | `psco.repository_binding.v1` | `v1` | `psco/repository_binding/v1/repository_binding.proto` |
| Dashboard | `psco.dashboard.v1` | `v1` | `psco/dashboard/v1/dashboard.proto` |
| Onboarding | `psco.onboarding.v1` | `v1` | `psco/onboarding/v1/onboarding.proto` |
| Export | `psco.export.v1` | `v1` | `psco/export/v1/export.proto` |
| Backup | `psco.backup.v1` | `v1` | `psco/backup/v1/backup.proto` |
| Reuse Summary | `psco.reuse_summary.v1` | `v1` | `psco/reuse_summary/v1/reuse_summary.proto` |

### 演进规则（统一适用于所有模块）

- 新增字段必须使用新的递增编号，不得复用已删除字段编号
- 破坏性变更（删除字段、修改字段类型）必须升级版本号（`v2`）
- 后续版本删除字段或废弃字段名时必须使用 `reserved` 保留字段号，必要时同时保留字段名
- `buf breaking` 用于 CI 中自动检测破坏性变更

### 跨包依赖

- `Decision Center` 通过 `import "psco/module_registry/v1/module_registry.proto"` 直接复用 `psco.module_registry.v1.ModuleStatus`，用于 `DecisionModuleCandidate.status`
- `Product Registry` 通过 `import "psco/module_registry/v1/module_registry.proto"` 直接复用 `psco.module_registry.v1.ModuleStatus`，用于 `BoundModuleSummary.module_status` / `ProductModuleCandidate.module_status`
- `Product Registry` 通过 `import "psco/common/v1/common.proto"` 直接复用 `psco.common.v1.ActiveArchivedStatus`，用于 `Product.status` / `ProductListItem.status` / `BoundRepositorySummary.repository_status` / `CreateProductRequest.status` / `ListProductsRequest.status_filter`
- `Repository Binding` 通过 `import "psco/module_registry/v1/module_registry.proto"` 直接复用 `psco.module_registry.v1.ModuleStatus`，用于 `MappedModuleSummary.module_status` / `RepositoryModuleCandidate.module_status`
- `Repository Binding` 通过 `import "psco/common/v1/common.proto"` 直接复用 `psco.common.v1.ActiveArchivedStatus`，用于 `Repository.status` / `RepositoryListItem.status` / `BoundProductSummary.product_status` / `RepositoryProductCandidate.product_status` / `CreateRepositoryRequest.status` / `ListRepositoriesRequest.status_filter`
- `Dashboard + Feedback` 当前阶段不复用 `ActiveArchivedStatus` 或 `ModuleStatus`，自有枚举（`FeedbackSignalFamily` / `FeedbackSignalCode` / `FeedbackSignalPriority` / `DashboardTargetType` / `RecentActivityType`）在 `dashboard.proto` 内单值定义，不通过 import 引入跨包枚举
- `Onboarding` 当前阶段不 import 其他 `psco.*` 包，自有枚举（`FirstRunStatus` / `OnboardingStep`）在 `onboarding.proto` 内单值定义；draft-first 四类创建动作继续复用既有 `ProductRegistryService.CreateProduct` / `RepositoryBindingService.CreateRepository` / `ModuleRegistryService.CreateModule` / `DecisionCenterService.CreateDecision` canonical 合同，不在 `OnboardingService` 下并列新增 `CreateDraft*` RPC
- `Export` 与 `Backup` 各自独立定义资产覆盖矩阵枚举（`ExportAssetScope` / `BackupAssetScope`），语义对齐同一 9 类核心资产但不通过 import 引入跨包枚举依赖，以保持 `Export` 与 `Backup` 模块边界互不耦合（对齐 phase06-08 模块分离冻结）
- `Reuse Summary` 当前阶段不 import 其他 `psco.*` 包，自有枚举（`ReuseSummaryScope`）在 `reuse_summary.proto` 内单值定义；`module_id` 与 `product_id` 以 `string` 形式承接，不引入对其他模块 service 或消息结构的合同依赖
- `common.proto` 只承接跨 `Product / Repository` 共享且不会引入业务 owner 歧义的最小公共枚举，不定义业务消息或服务接口
- 不得在 `psco.product_registry.v1` 或 `psco.repository_binding.v1` 中重定义本地等价 `ActiveArchivedStatus` 或 `ModuleStatus` 枚举
- 该 import 仅限于复用枚举类型，不引入对其他模块 service 或消息结构的合同依赖

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

| 语言 | 产物类型 | 插件 | 落点 |
| --- | --- | --- | --- |
| Go | Proto 消息类型 | `buf.build/protocolbuffers/go` | `backend/internal/gen/proto/psco/**/*.pb.go` |
| Go | Connect handler + client（simple 模式） | `buf.build/connectrpc/gosimple` | `backend/internal/gen/connect/psco/**/*.connect.go` |
| TypeScript | Proto 消息类型 + service descriptor | `buf.build/bufbuild/es` | `frontend/src/gen/proto/psco/**/*_pb.ts` |

> 生成产物已加入 `.gitignore`，不进入版本控制。每次 `make gen` 重新生成。
> `phase07-08` 已将生成链从 2 插件（Go protobuf + TS）升级为 3 插件正式矩阵（Go protobuf + Go Connect simple + TS）。
> Connect handler 的 simple 模式直接使用 proto 消息类型作为 handler 签名（如 `(ctx, *pb.ListModulesRequest) (*pb.ListModulesResponse, error)`），与现有 Go service 层签名风格一致。

## 4. 过渡传输层映射

当前阶段保留 `chi + JSON HTTP` 作为过渡传输层（`phase02-11` 已实现；`Decision Center` 由 `phase03-12 / 13` 实现；`Product / Repository / Binding` 由 `phase04-12 / 13` 实现；`Dashboard + Feedback` 由 `phase05-12 / 13` 实现；`Onboarding + Export + Backup + Reuse Summary` 由 `phase06-14 / 15` 实现）。
`.proto` 是唯一合同源，JSON 请求与响应语义从 `.proto` 派生。

### RPC → HTTP 映射矩阵

#### Module Registry

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `ListModules` | GET | `/api/modules` | query params |
| `GetModuleDetail` | GET | `/api/modules/{moduleId}` | URL path |
| `CreateModule` | POST | `/api/modules` | JSON body |
| `CreateRelease` | POST | `/api/modules/{moduleId}/releases` | JSON body + URL path |
| `BindModuleToProduct` | POST | `/api/modules/{moduleId}/bindings/products` | JSON body + URL path（兼容委派到 Product Registry） |
| `MapModuleToRepository` | POST | `/api/modules/{moduleId}/bindings/repositories` | JSON body + URL path（兼容委派到 Repository Binding） |

> phase04-12 起 canonical 候选读取已切换到 Product Registry / Repository Binding；
> `/api/candidates/products` 与 `/api/candidates/repositories` 当前仍作为 Module Registry 历史入口的 compat 路由保留，最晚在 `phase07-09` 退场。
> canonical 候选读取由 Product Registry 的 `ListProductModuleCandidates` 与 Repository Binding 的
> `ListRepositoryProductCandidates` / `ListRepositoryModuleCandidates` 各自承接。

#### Decision Center

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `ListDecisions` | GET | `/api/decisions` | query params |
| `GetDecisionDetail` | GET | `/api/decisions/{decisionId}` | URL path |
| `CreateDecision` | POST | `/api/decisions` | JSON body |
| `LinkDecisionToTarget` | POST | `/api/decisions/{decisionId}/links` | JSON body + URL path |
| `ListDecisionModuleCandidates` | GET | `/api/decisions/{decisionId}/candidates/modules` | URL path |

#### Product Registry

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `ListProducts` | GET | `/api/products` | query params |
| `GetProductDetail` | GET | `/api/products/{productId}` | URL path |
| `CreateProduct` | POST | `/api/products` | JSON body |
| `BindModuleToProduct` | POST | `/api/products/{productId}/bindings/modules` | JSON body + URL path |
| `ListProductModuleCandidates` | GET | `/api/products/{productId}/candidates/modules` | URL path |

#### Repository Binding

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `ListRepositories` | GET | `/api/repositories` | query params |
| `GetRepositoryDetail` | GET | `/api/repositories/{repositoryId}` | URL path |
| `CreateRepository` | POST | `/api/repositories` | JSON body |
| `BindRepositoryToProduct` | POST | `/api/repositories/{repositoryId}/bindings/products` | JSON body + URL path |
| `MapModuleToRepository` | POST | `/api/repositories/{repositoryId}/bindings/modules` | JSON body + URL path |
| `ListRepositoryProductCandidates` | GET | `/api/repositories/{repositoryId}/candidates/products` | URL path |
| `ListRepositoryModuleCandidates` | GET | `/api/repositories/{repositoryId}/candidates/modules` | URL path |

#### Dashboard

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `GetDashboardOverview` | GET | `/api/dashboard/overview` | 无 |
| `GetFeedbackSignals` | GET | `/api/dashboard/feedback-signals` | 无 |
| `GetRecentActivities` | GET | `/api/dashboard/recent-activities` | 无 |

> Dashboard 三个 `GET` 入口当前阶段无 body、无 query 过滤、无路径参数。handler 必须显式组装空 Proto request 再进入业务层，不得因 `GET` 无参数就绕过 request 合同边界。

#### Onboarding

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `GetFirstRunState` | GET | `/api/onboarding/state` | 无 |

> Onboarding 当前阶段只承接 `GetFirstRunState` 读组，不承接 `CreateDraft*` 写组。
> 四类 draft-first 创建动作继续复用既有 canonical create HTTP 映射（`POST /api/products` / `POST /api/repositories` / `POST /api/modules` / `POST /api/decisions`），不在 Onboarding 下发明 `/api/onboarding/drafts/*` 第二套路由分组。
> Onboarding 页面只是既有 canonical create 合同的消费入口，不成为新的写合同 owner。
> `GetFirstRunState` 的 `GET` 入口当前阶段无 body、无 query 过滤、无路径参数。handler 必须显式组装空 Proto request 再进入业务层，不得因 `GET` 无参数就绕过 request 合同边界。

#### Export

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `GetExportSnapshot` | GET | `/api/dashboard/export` | 无 |
| `ExportCoreAssets` | POST | `/api/dashboard/export` | 无（当前阶段承接全部 9 类核心资产） |

> `GetExportSnapshot` 的 `GET` 入口当前阶段无 body、无 query 过滤、无路径参数。handler 必须显式组装空 Proto request 再进入业务层，不得因 `GET` 无参数就绕过 request 合同边界。
> `ExportCoreAssets` 当前阶段不引入按 scope 选择性导出的 request 字段，承接全部 9 类核心资产。

#### Backup

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `GetBackupSnapshot` | GET | `/api/dashboard/backup` | 无 |
| `CreateInstanceBackup` | POST | `/api/dashboard/backup` | 无（当前阶段承接全部 9 类核心资产） |

> `GetBackupSnapshot` 是当前阶段正式 `read / verify` 子路径合同入口，由独立读取 owner（`BackupRead` 或等价读取接口）承接，不得被 `CreateInstanceBackup` 写入响应附带的临时结果代替。
> `GetBackupSnapshot` 的 `GET` 入口当前阶段无 body、无 query 过滤、无路径参数。handler 必须显式组装空 Proto request 再进入业务层，不得因 `GET` 无参数就绕过 request 合同边界。
> `CreateInstanceBackup` 当前阶段不引入按 scope 选择性备份的 request 字段，承接全部 9 类核心资产。

#### Reuse Summary

| RPC 方法 | HTTP 方法 | HTTP 路径 | 请求体来源 |
| --- | --- | --- | --- |
| `GetReuseSummary` | GET | `/api/reuse-summary` | query params（`scope` / `module_id` / `product_id`） |

> `GetReuseSummary` 的 `scope` / `module_id` / `product_id` 必须通过 query 参数映射到 Proto request，handler 必须在进入业务层前显式组装对应的 Proto request 消息，不得因 `GET` 入口无 body 就绕过 Proto request 这一合同边界。
> `scope` 与参数使用关系：`dashboard` 下 `module_id` / `product_id` 均不使用；`module_detail` 下使用 `module_id`，不使用 `product_id`；`product_detail` 下使用 `product_id`，不使用 `module_id`。

> 约束：JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 语义显式对齐；HTTP 过渡层使用 URL 路径参数承接 `decisionId` / `moduleId` / `productId` / `repositoryId` 时，handler 必须在进入业务层前显式组装为对应的 Proto request 字段；不得把 HTTP 路径、状态码或中间件策略误写成 Proto 合同本体。
> `phase06` 的 `GetReuseSummary` 使用 query 参数（而非 URL 路径参数）承接作用域字段，handler 必须从 query string 提取 `scope` / `module_id` / `product_id` 后组装为 Proto request。

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

#### product_id 的传输差异

- **Proto RPC**：`product_id` 作为请求消息的显式字段（如 `GetProductDetailRequest.product_id` / `BindModuleToProductRequest.product_id` / `ListProductModuleCandidatesRequest.product_id`）
- **HTTP 过渡层**：`product_id` 由 URL 路径参数承接，不放在 JSON 请求体

这是传输层差异，不是合同差异。HTTP handler 从 URL 提取 `product_id` 后组装为 Proto 请求。

#### repository_id 的传输差异

- **Proto RPC**：`repository_id` 作为请求消息的显式字段（如 `GetRepositoryDetailRequest.repository_id` / `BindRepositoryToProductRequest.repository_id` / `MapModuleToRepositoryRequest.repository_id` / `ListRepositoryProductCandidatesRequest.repository_id` / `ListRepositoryModuleCandidatesRequest.repository_id`）
- **HTTP 过渡层**：`repository_id` 由 URL 路径参数承接，不放在 JSON 请求体

这是传输层差异，不是合同差异。HTTP handler 从 URL 提取 `repository_id` 后组装为 Proto 请求。

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
| `CreateModuleRequest` | `CreateModuleRequest` | 语义对齐（name 为最小人工必填；description 默认 ""；status 默认 active） |
| `CreateReleaseRequest` | `CreateReleaseRequest` | 语义对齐（types.go 无 module_id，由 URL 承接；released_at: string ↔ Timestamp） |
| `BindModuleToProductRequest` | `BindModuleToProductRequest` | 语义对齐（types.go 无 module_id，由 URL 承接） |
| `MapModuleToRepositoryRequest` | `MapModuleToRepositoryRequest` | 语义对齐（types.go 无 module_id，由 URL 承接） |

> 后续阶段可逐步将 `types.go` 替换为从 `.proto` 生成的 Go 类型，但当前阶段不做此迁移。

### 前端 — Module Registry（`frontend/src/features/module-registry/types.ts`）

当前 `types.ts` 中的手写 TypeScript 接口与 `.proto` 消息语义对齐。
命名约定按类型分层：
- 响应/域类型（`Module` / `Release` / `ModuleListItem` / `ProductBinding` / `RepositoryMapping` / `DecisionLink` / `ModuleDetail` / `ProductCandidate` / `RepositoryCandidate`）使用 **snake_case**（如 `created_at` / `module_id` / `product_bind_count`），直接承接后端 JSON（snake_case），与 `.proto` 的 snake_case 对齐，无需转换。
- 输入/写入参数（`CreateModuleInput` / `CreateReleaseInput` / `BindModuleToProductInput` / `MapModuleToRepositoryInput`）与搜索参数（`ModuleListSearch`）使用 **camelCase**（如 `moduleId` / `releasedAt` / `queryText`），属前端表单与路由层命名，由 adapter 层在组装 JSON 请求时转换为 snake_case。
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
| `CreateDecisionRequest` | `CreateDecisionRequest` | 语义对齐（title / choice / reason 为最小人工必填；status 默认 proposed） |
| `LinkDecisionToTargetRequest` | `LinkDecisionToTargetRequest` | 语义对齐（types.go 无 decision_id，由 URL 承接） |

> 实现期不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义。

### 前端 — Decision Center（`frontend/src/features/decision-center/`，由 phase03-13 实现）

`Decision Center` 的前端过渡传输层适配代码由 `phase03-13` 实现。
命名约定按类型分层：
- 响应/域类型（`Decision` / `DecisionListItem` / `LinkedModule` / `SourceContext` / `DecisionDetail` / `DecisionModuleCandidate` / `CreateDecisionResponse`）使用 **snake_case**（如 `created_at` / `module_id` / `link_count` / `linked_module_summary`），直接承接后端 JSON（snake_case），与 `.proto` 的 snake_case 对齐，无需转换。`types.ts` 文件头明确标注"后端返回 JSON 使用 snake_case，前端类型直接承接，无需转换"。
- 输入/写入参数（`CreateDecisionInput` / `LinkDecisionToTargetInput`）与搜索参数（`DecisionListSearch`）使用 **camelCase**（如 `decisionId` / `queryText` / `statusFilter`），属前端表单与路由层命名，由 adapter 层在组装 JSON 请求时转换为 snake_case。`CreateDecisionInput.source_module_id` 为 snake_case，属历史兼容字段，后续统一为 camelCase。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

### 后端 — Product Registry（`backend/internal/productregistry/types.go`，由 phase04-12 实现）

`Product Registry` 的 `.proto` 合同源已在 `phase04-11` 落地，后端过渡传输层 DTO 由 `phase04-12` 实现。
实现时必须遵守的单向承接约束：

| 预期 `types.go` 结构体 | `.proto` 消息 | 单向承接约束 |
| --- | --- | --- |
| `Product` | `Product` | 语义对齐（created_at: time.Time ↔ Timestamp；status 复用 ActiveArchivedStatus） |
| `ProductListItem` | `ProductListItem` | 语义对齐 |
| `BoundModuleSummary` | `BoundModuleSummary` | 语义对齐（module_status 复用 ModuleStatus，不重定义本地枚举） |
| `BoundRepositorySummary` | `BoundRepositorySummary` | 语义对齐（repository_status 复用 ActiveArchivedStatus） |
| `ProductDetail` | `ProductDetail` | 语义对齐 |
| `ProductModuleCandidate` | `ProductModuleCandidate` | 语义对齐（module_status 复用 ModuleStatus） |
| `CreateProductRequest` | `CreateProductRequest` | 语义对齐（name 为最小人工必填；status 复用 ActiveArchivedStatus，默认 active） |
| `BindModuleToProductRequest` | `BindModuleToProductRequest` | 语义对齐（types.go 无 product_id，由 URL 承接） |

> 实现期不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义。

### 前端 — Product Registry（`frontend/src/features/product-registry/`，由 phase04-13 实现）

`Product Registry` 的前端过渡传输层适配代码由 `phase04-13` 实现。
命名约定按类型分层（与现有 Module Registry / Decision Center 模式同构）：
- 响应/域类型（`Product` / `ProductListItem` / `BoundModuleSummary` / `BoundRepositorySummary` / `ProductDetail` / `ProductModuleCandidate`）应使用 **snake_case**（如 `created_at` / `module_id` / `module_bind_count`），直接承接后端 JSON（snake_case），与 `.proto` 的 snake_case 对齐，无需转换。
- 输入/写入参数与搜索参数应使用 **camelCase**（如 `productId` / `queryText` / `statusFilter`），属前端表单与路由层命名，由 adapter 层在组装 JSON 请求时转换为 snake_case。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

### 后端 — Repository Binding（`backend/internal/repositorybinding/types.go`，由 phase04-12 实现）

`Repository Binding` 的 `.proto` 合同源已在 `phase04-11` 落地，后端过渡传输层 DTO 由 `phase04-12` 实现。
实现时必须遵守的单向承接约束：

| 预期 `types.go` 结构体 | `.proto` 消息 | 单向承接约束 |
| --- | --- | --- |
| `Repository` | `Repository` | 语义对齐（created_at: time.Time ↔ Timestamp；status 复用 ActiveArchivedStatus） |
| `RepositoryListItem` | `RepositoryListItem` | 语义对齐 |
| `BoundProductSummary` | `BoundProductSummary` | 语义对齐（product_status 复用 ActiveArchivedStatus） |
| `MappedModuleSummary` | `MappedModuleSummary` | 语义对齐（module_status 复用 ModuleStatus，不重定义本地枚举） |
| `RepositoryDetail` | `RepositoryDetail` | 语义对齐 |
| `RepositoryProductCandidate` | `RepositoryProductCandidate` | 语义对齐（product_status 复用 ActiveArchivedStatus） |
| `RepositoryModuleCandidate` | `RepositoryModuleCandidate` | 语义对齐（module_status 复用 ModuleStatus） |
| `CreateRepositoryRequest` | `CreateRepositoryRequest` | 语义对齐（name + url 为最小人工必填；status 复用 ActiveArchivedStatus，默认 active） |
| `BindRepositoryToProductRequest` | `BindRepositoryToProductRequest` | 语义对齐（types.go 无 repository_id，由 URL 承接） |
| `MapModuleToRepositoryRequest` | `MapModuleToRepositoryRequest` | 语义对齐（types.go 无 repository_id，由 URL 承接） |

> 实现期不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义。

### 前端 — Repository Binding（`frontend/src/features/repository-binding/`，由 phase04-13 实现）

`Repository Binding` 的前端过渡传输层适配代码由 `phase04-13` 实现。
命名约定按类型分层（与现有 Module Registry / Decision Center 模式同构）：
- 响应/域类型（`Repository` / `RepositoryListItem` / `BoundProductSummary` / `MappedModuleSummary` / `RepositoryDetail` / `RepositoryProductCandidate` / `RepositoryModuleCandidate`）应使用 **snake_case**（如 `created_at` / `product_id` / `product_bind_count`），直接承接后端 JSON（snake_case），与 `.proto` 的 snake_case 对齐，无需转换。
- 输入/写入参数与搜索参数应使用 **camelCase**（如 `repositoryId` / `queryText` / `statusFilter`），属前端表单与路由层命名，由 adapter 层在组装 JSON 请求时转换为 snake_case。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

### 后端 — Dashboard（`backend/internal/dashboard/types.go`，由 phase05-12 实现）

`Dashboard + Feedback` 的 `.proto` 合同源已在 `phase05-11` 落地，后端过渡传输层 DTO 由 `phase05-12` 实现。
实现时必须遵守的单向承接约束：

| 预期 `types.go` 结构体 | `.proto` 消息 | 单向承接约束 |
| --- | --- | --- |
| `DashboardOverview` | `DashboardOverview` | 语义对齐（纯 int32 计数字段，无时间字段） |
| `FeedbackSignal` | `FeedbackSignal` | 语义对齐（枚举使用自有枚举，不引入 ActiveArchivedStatus / ModuleStatus） |
| `ProductAssetCoverageSummary` | `ProductAssetCoverageSummary` | 语义对齐（missing_both_bindings_count 作为独立计数字段，不回退为隐式组合） |
| `RecentActivityItem` | `RecentActivityItem` | 语义对齐（activity_at: time.Time ↔ Timestamp） |
| `GetDashboardOverviewResponse` | `GetDashboardOverviewResponse` | 语义对齐 |
| `GetFeedbackSignalsResponse` | `GetFeedbackSignalsResponse` | 语义对齐（主队列 + 资产缺口摘要两层结构） |
| `GetRecentActivitiesResponse` | `GetRecentActivitiesResponse` | 语义对齐 |

> 实现期不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义。
> 错误状态码、错误包络与局部错误展示仍属于 HTTP / handler 适配层，不进入 `.proto` 合同本体。

### 前端 — Dashboard（`frontend/src/features/dashboard/`，由 phase05-13 实现）

`Dashboard + Feedback` 的前端过渡传输层适配代码由 `phase05-13` 实现。
命名约定按类型分层（与现有 Module Registry / Decision Center / Product Registry / Repository Binding 模式同构）：
- 响应/域类型（`DashboardOverview` / `FeedbackSignal` / `ProductAssetCoverageSummary` / `RecentActivityItem`）应使用 **snake_case**（如 `signal_family` / `signal_code` / `target_id` / `activity_at` / `missing_both_bindings_count`），直接承接后端 JSON（snake_case），与 `.proto` 的 snake_case 对齐，无需转换。
- Dashboard 三个读组当前阶段无输入/写入参数与搜索参数，不涉及 camelCase 转换。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

### 后端 — Onboarding / Export / Backup / Reuse Summary（由 phase06-14 实现）

`Onboarding / Export / Backup / Reuse Summary` 的 `.proto` 合同源已在 `phase06-13` 落地，后端过渡传输层 DTO 由 `phase06-14` 实现。
实现时必须遵守的单向承接约束：

| 预期 `types.go` 结构体 | `.proto` 消息 | 单向承接约束 |
| --- | --- | --- |
| `FirstRunState` | `FirstRunState` | 语义对齐（status: 枚举 ↔ FirstRunStatus；current_step: 枚举 ↔ OnboardingStep） |
| `GetFirstRunStateResponse` | `GetFirstRunStateResponse` | 语义对齐 |
| `ExportSnapshot` | `ExportSnapshot` | 语义对齐（created_at: time.Time ↔ Timestamp；asset_scope: []string ↔ repeated ExportAssetScope） |
| `ExportResultStatus` | `ExportResultStatus` | 语义对齐（自有枚举，不引入跨包枚举） |
| `GetExportSnapshotResponse` / `ExportCoreAssetsResponse` | 同名 `.proto` 消息 | 语义对齐 |
| `BackupSnapshot` | `BackupSnapshot` | 语义对齐（created_at: time.Time ↔ Timestamp；verified_status: 枚举 ↔ BackupVerifiedStatus；verify_failure_code: 枚举 ↔ VerifyFailureCode） |
| `ManifestSummary` / `AssetCoverageEntry` / `SchemaVersionPrerequisite` | 同名 `.proto` 消息 | 语义对齐 |
| `GetBackupSnapshotResponse` / `CreateInstanceBackupResponse` | 同名 `.proto` 消息 | 语义对齐（GetBackupSnapshot 读取侧由独立读取 owner 承接，不与 CreateInstanceBackup 写入响应耦合） |
| `ModuleReuseSummary` | `ModuleReuseSummary` | 语义对齐（latest_reuse_at: time.Time ↔ Timestamp） |
| `CapabilitySummary` | `CapabilitySummary` | 语义对齐（latest_capability_update_at: time.Time ↔ Timestamp；empty_state_text 承接成功空态解释） |
| `GetReuseSummaryRequest` | `GetReuseSummaryRequest` | 语义对齐（scope: 枚举 ↔ ReuseSummaryScope；module_id / product_id 由 query 参数承接） |
| `GetReuseSummaryResponse` | `GetReuseSummaryResponse` | 语义对齐（同时承接 module_reuse_summary[] + capability_summary[]） |

> 实现期不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义。
> `OnboardingRead`（含 `first_run_state`）的 `.proto -> HTTP DTO -> 前端消费模型` 单值一致必须进入验收门禁。
> `backup_snapshot` 读取侧合同一致性必须由 `BackupWrite.read_verify` 或等价上游冻结承接位独立验证，不得只验证 `BackupWrite` 写入响应而遗漏 `GetBackupSnapshot` 读取侧合同一致性。
> 错误状态码、错误包络与局部错误展示仍属于 HTTP / handler 适配层，不进入 `.proto` 合同本体。

### 前端 — Onboarding / Export / Backup / Reuse Summary（由 phase06-15 实现）

`Onboarding / Export / Backup / Reuse Summary` 的前端过渡传输层适配代码由 `phase06-15` 实现。
命名约定按类型分层（与现有 Module Registry / Decision Center / Product Registry / Repository Binding / Dashboard 模式同构）：
- 响应/域类型（`FirstRunState` / `ExportSnapshot` / `BackupSnapshot` / `ManifestSummary` / `AssetCoverageEntry` / `SchemaVersionPrerequisite` / `ModuleReuseSummary` / `CapabilitySummary`）应使用 **snake_case**（如 `is_first_entry` / `current_step` / `completion_progress` / `asset_scope` / `created_at` / `result_status` / `manifest_summary` / `asset_coverage` / `schema_version_prerequisite` / `verified_status` / `reuse_product_count` / `latest_reuse_at` / `explanation_text` / `capability_key` / `capability_label` / `supporting_module_count` / `latest_capability_update_at` / `empty_state_text`），直接承接后端 JSON（snake_case），与 `.proto` 的 snake_case 对齐，无需转换。
- 输入/写入参数与搜索参数应使用 **camelCase**（如 `scope` / `moduleId` / `productId`），属前端表单与路由层命名，由 adapter 层在组装 query 请求时转换为 snake_case。
- Onboarding / Export / Backup 的读组与写组当前阶段无 query 过滤参数；Reuse Summary 的 `GetReuseSummary` 使用 `scope` / `moduleId` / `productId` 作为 query 参数。
后续阶段可逐步替换为从 `.proto` 生成的 TypeScript 类型。

## 6. 当前阶段不做

- 不完成完整 Connect 传输层迁移（Go handler + 前端 client 实现由后续 phase07 子任务承接）
- 不替换 `chi + JSON HTTP` 为 Connect 服务器
- 不将 Connect 生成代码集成到现有 handler / adapter 中（由 phase07 后续实现子任务承接）
- 不引入 gRPC 网关或连接池
- 以上属于后续 phase07 子任务的范围

> `phase07-08` 已落地 3 插件正式生成链与 Connect runtime 依赖，后续 `phase07` 实现、退场与验收必须以此生成链为唯一上游。

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
> 当前 `main` 分支已有 `phase02-11A` 提交的 `Module Registry` proto 基准，`phase03-11` 新增的 `Decision Center` proto、`phase04-11` 新增的 `Common / Product Registry / Repository Binding` proto、`phase05-11` 新增的 `Dashboard + Feedback` proto 与 `phase06-13` 新增的 `Onboarding / Export / Backup / Reuse Summary` proto 均属于向后兼容的新增，`make breaking` 退出码为 0。
> 后续任何对已存在字段编号、字段类型或字段语义的删除或修改都会被检测为破坏性变更并返回非零退出码。
