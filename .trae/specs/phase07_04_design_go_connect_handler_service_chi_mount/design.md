# Phase07-04 设计产出：Go Connect handler、service implementation 与 chi mount 设计

> 本文档是 phase07-04 spec 的执行产出，定义 Connect handler 与既有 service 的接线方式、chi/Connect 横切逻辑分层、router 结构调整、错误映射方案与 service 分层保持策略。
> 产出日期：2026-08-11
> 上游：`phase07-01 frozen_scope.md`（34 RPC 迁移总表）、`phase07-02 frozen_scope.md`（chi + Connect + buf 组合）、`phase07-03 frozen_scope.md`（compat 策略与退场标准）
> 参考：Context7 `/connectrpc/connect-go` 最新文档

---

## 1. 设计：Generated Connect Handler 与 Service Implementation 正式接线

### 1.1 接线总览

```
┌──────────────────────────────────────────────────────────────────┐
│  platform/server.go（装配根）                                      │
│  ├── buildModuleRegistry(pool) → (querySvc, commandSvc)           │
│  ├── buildDecisionCenter(pool) → (querySvc, commandSvc)           │
│  ├── buildProductRegistry(pool, ...) → (querySvc, commandSvc)     │
│  ├── buildRepositoryBinding(pool) → (querySvc, commandSvc)        │
│  ├── buildDashboard(pool) → (querySvc)                            │
│  ├── buildOnboarding(pool) → (querySvc)                           │
│  ├── buildExport(pool) → (querySvc, commandSvc)                   │
│  ├── buildBackup(pool) → (querySvc, commandSvc)                   │
│  └── buildReuseSummary(pool) → (querySvc)                         │
│                                                                    │
│  ┌─ mountConnect(router, querySvc, commandSvc) ─┐                 │
│  │  moduleRegConnectSvc := &connect.ModuleRegistryServer{         │
│  │      QueryService:   querySvc,               │                 │
│  │      CommandService: commandSvc,             │                 │
│  │  }                                           │                 │
│  │  path, handler := module_registryv1connect.  │                 │
│  │      NewModuleRegistryServiceHandler(        │                 │
│  │          moduleRegConnectSvc,                │                 │
│  │      )                                       │                 │
│  │  router.Handle(path, handler)                │                 │
│  └──────────────────────────────────────────────┘                 │
│                                                                    │
│  ┌─ mountCompat(r, ...) ───────────────────────┐                  │
│  │  // compat 过渡组，按 phase07-03 §2 退场     │                  │
│  │  r.Get("/candidates/products", ...)          │                  │
│  │  r.Get("/candidates/repositories", ...)      │                  │
│  │  r.Post("/modules/{id}/bindings/...",...)    │                  │
│  └──────────────────────────────────────────────┘                  │
└──────────────────────────────────────────────────────────────────┘
```

### 1.2 各层职责

| 层 | 文件落点 | 职责 | 迁移后变化 |
|----|---------|------|-----------|
| **`build*` 函数** | `platform/server.go`（保持） | 构造 repository store → service，解决依赖注入顺序 | 无变化，保持现有 `build*` 模式 |
| **`service` 层** | `backend/internal/<module>/service/`（保持） | 业务编排、校验、跨 store 聚合 | 无变化，保持现有 `QueryService` / `CommandService` |
| **Connect 实现** | `backend/internal/<module>/connect/`（**新增**） | transport 解包、调用 service、返回 proto response / Connect error | 新增，每个 module 一个 `connect/server.go` |
| **Generated handler** | `backend/internal/gen/connect/`（自动生成） | Connect procedure 路由、序列化、协议协商 | 新增，`buf generate` 产出 |
| **`platform` 装配** | `platform/server.go`（修改） | `build*` → Connect 实现 → `New*Handler` → `router.Handle(path, handler)` | 修改：从 `mount*` 手写 JSON 路由改为 `mountConnect` |

### 1.3 文件落点与命名

```
backend/internal/
├── platform/
│   ├── server.go                  # 修改：build* → Connect 实现 → Handle(path, handler)
│   ├── router.go                  # 大幅修改：移除手写 JSON 业务路由，保留 infra + compat 分组
│   └── connect_errors.go          # 新增：单值 error → Connect code 映射
├── moduleregistry/
│   ├── service/                   # 保持：QueryService + CommandService
│   │   ├── query_service.go
│   │   └── command_service.go
│   ├── handler/                   # 兼容过渡期保留，phase07-09 退场
│   │   ├── query_handler.go
│   │   ├── command_handler.go
│   │   └── response.go
│   └── connect/                   # 新增：Connect transport 实现
│       └── server.go
├── decisioncenter/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
├── productregistry/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
├── repositorybinding/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
├── dashboard/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
├── onboarding/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
├── export/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
├── backup/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
├── reusesummary/
│   ├── service/                   # 保持
│   ├── handler/                   # 兼容过渡
│   └── connect/                   # 新增
│       └── server.go
└── gen/
    ├── proto/                     # Go protobuf 消息（已有）
    └── connect/                   # Go Connect handler（新增，buf generate 产出）
```

### 1.4 Connect 实现模板（以 Module Registry 为例）

```go
// backend/internal/moduleregistry/connect/server.go
package connect

import (
    "context"

    "connectrpc.com/connect"

    pb "github.com/psco/backend/internal/gen/proto/psco/module_registry/v1"
    "github.com/psco/backend/internal/gen/connect/psco/module_registry/v1/module_registryv1connect"
    "github.com/psco/backend/internal/moduleregistry"
    "github.com/psco/backend/internal/moduleregistry/service"
    "github.com/psco/backend/internal/platform"
)

// ModuleRegistryServer 实现 generated Connect handler interface。
// 嵌入 Unimplemented*Handler 确保新增 RPC 时编译不炸。
type ModuleRegistryServer struct {
    module_registryv1connect.UnimplementedModuleRegistryServiceHandler

    query   *service.QueryService
    command *service.CommandService
}

// NewModuleRegistryServer 构造 Connect service implementation。
func NewModuleRegistryServer(query *service.QueryService, command *service.CommandService) *ModuleRegistryServer {
    return &ModuleRegistryServer{query: query, command: command}
}

// ListModules 实现 ListModules RPC。
func (s *ModuleRegistryServer) ListModules(
    ctx context.Context,
    req *pb.ListModulesRequest,
) (*pb.ListModulesResponse, error) {
    items, err := s.query.ListModules(ctx, moduleregistry.ListQuery{
        QueryText:    req.QueryText,
        StatusFilter: moduleregistry.ModuleStatus(req.StatusFilter),
    })
    if err != nil {
        return nil, platform.MapToConnectError(err)
    }
    // 组装 proto response
    pbItems := make([]*pb.ModuleListItem, 0, len(items))
    for _, item := range items {
        pbItems = append(pbItems, moduleListItemToProto(item))
    }
    return &pb.ListModulesResponse{Modules: pbItems}, nil
}

// GetModuleDetail 实现 GetModuleDetail RPC。
func (s *ModuleRegistryServer) GetModuleDetail(
    ctx context.Context,
    req *pb.GetModuleDetailRequest,
) (*pb.GetModuleDetailResponse, error) {
    detail, err := s.query.GetModuleDetail(ctx, req.ModuleId)
    if err != nil {
        return nil, platform.MapToConnectError(err)
    }
    return &pb.GetModuleDetailResponse{Detail: moduleDetailToProto(detail)}, nil
}

// CreateModule 实现 CreateModule RPC。
func (s *ModuleRegistryServer) CreateModule(
    ctx context.Context,
    req *pb.CreateModuleRequest,
) (*pb.CreateModuleResponse, error) {
    m, err := s.command.CreateModule(ctx, moduleregistry.CreateModuleRequest{
        Name:        req.Name,
        Description: req.Description,
        Status:      req.Status,
    })
    if err != nil {
        return nil, platform.MapToConnectError(err)
    }
    return &pb.CreateModuleResponse{Module: moduleToProto(m)}, nil
}

// ... 其余 RPC 同理
```

### 1.5 Platform 装配模板

```go
// backend/internal/platform/server.go（修改后）

func NewServer(cfg Config, pool *pgxpool.Pool) *Server {
    r := chi.NewRouter()

    // 根级 middleware（不变）
    r.Use(chimw.RequestID)
    r.Use(chimw.RealIP)
    r.Use(chimw.Logger)
    r.Use(chimw.Recoverer)
    r.Use(chimw.Timeout(60 * time.Second))
    r.Use(corsMiddleware)

    // 健康检查（不变）
    r.Get("/healthz", healthz)

    // 装配业务模块
    r.Route("/api", func(r chi.Router) {
        // 依赖注入（保持 build* 模式）
        repoQuerySvc, repoCommandSvc, repoBindingStore := buildRepositoryBinding(pool)
        productQuerySvc, productCommandSvc := buildProductRegistry(pool, repoBindingStore)

        // === Canonical Connect 主线 ===
        mountConnectModuleRegistry(r, pool, productQuerySvc, productCommandSvc, repoQuerySvc, repoCommandSvc)
        mountConnectDecisionCenter(r, pool)
        mountConnectProductRegistry(r, productQuerySvc, productCommandSvc)
        mountConnectRepositoryBinding(r, repoQuerySvc, repoCommandSvc)

        dashboardQuerySvc := buildDashboard(pool)
        mountConnectDashboard(r, dashboardQuerySvc)

        onboardingQuerySvc := buildOnboarding(pool)
        mountConnectOnboarding(r, onboardingQuerySvc)

        exportQuerySvc, exportCommandSvc := buildExport(pool)
        mountConnectExport(r, exportQuerySvc, exportCommandSvc)

        backupQuerySvc, backupCommandSvc := buildBackup(pool)
        mountConnectBackup(r, backupQuerySvc, backupCommandSvc)

        reuseSummaryQuerySvc := buildReuseSummary(pool)
        mountConnectReuseSummary(r, reuseSummaryQuerySvc)

        // === Compat 过渡组（phase07-09/10 退场） ===
        mountCompatRoutes(r, pool, productQuerySvc, productCommandSvc, repoQuerySvc, repoCommandSvc)
    })

    return &Server{...}
}

// mountConnectModuleRegistry 在 /api 下挂载 Module Registry Connect handler
func mountConnectModuleRegistry(r chi.Router, pool *pgxpool.Pool, ...) {
    // 构造既有 service 层
    moduleStore := moduleregistryrepo.NewModuleStore(pool)
    releaseStore := moduleregistryrepo.NewReleaseStore(pool)
    bindingStore := moduleregistryrepo.NewBindingStore(pool)
    querySvc := moduleregistrysvc.NewQueryService(moduleStore, releaseStore, bindingStore)
    commandSvc := moduleregistrysvc.NewCommandService(moduleStore, releaseStore)

    // 构造 Connect implementation
    connectSvc := moduleregistryconnect.NewModuleRegistryServer(querySvc, commandSvc)

    // 挂载 generated handler
    // generated constructor 返回 (path string, handler http.Handler)
    path, handler := module_registryv1connect.NewModuleRegistryServiceHandler(connectSvc)
    // 在 /api 子路由内显式消费 generated path，避免同一路径重复 Mount 冲突
    r.Handle(path, handler)
}
```

---

## 2. 设计：`chi` Middleware、Connect Interceptor 与错误处理链承接

### 2.1 两层横切架构

```
请求进入
    │
    ▼
┌─────────────────────────────────────────────────────────┐
│  chi 根级 middleware（HTTP 壳层治理）                      │
│  ├── RequestID    → 注入 request_id 到 context            │
│  ├── RealIP       → 解析真实客户端 IP                      │
│  ├── Logger       → 请求日志（method / path / status / duration）│
│  ├── Recoverer    → panic 恢复、500 兜底                   │
│  ├── Timeout      → 请求超时控制                           │
│  └── CORS         → 跨域头设置                             │
├─────────────────────────────────────────────────────────┤
│  chi /api 子路由器（仅 mount shell，不承担业务路由定义）      │
│    │                                                     │
│    ├── Connect handler mount（canonical 业务主线）          │
│    │     │                                               │
│    │     ▼                                               │
│    │  ┌─────────────────────────────────────────────┐    │
│    │  │  Connect interceptor 链（RPC 级治理）         │    │
│    │  │  ├── Metadata 读取（request_id 等 context 值）│    │
│    │  │  ├── 请求校验（option: protovalidate）        │    │
│    │  │  ├── 错误归一化（domain error → Connect code）│    │
│    │  │  └── 审计扩展位（未来）                       │    │
│    │  ├─────────────────────────────────────────────┤    │
│    │  │  Connect service implementation              │    │
│    │  │  ├── 解包 proto request                       │    │
│    │  │  ├── 调用既有 service 层                       │    │
│    │  │  └── 返回 proto response / Connect error      │    │
│    │  └─────────────────────────────────────────────┘    │
│    │                                                     │
│    └── compat 过渡组（仅迁移期，复用 chi 外壳治理）          │
│          └── 旧 JSON handler（phase07-09/10 退场）        │
└─────────────────────────────────────────────────────────┘
```

### 2.2 职责边界

| 横切关注点 | 唯一承接位 | 禁止 |
|-----------|-----------|------|
| Request ID | `chi` middleware `chimw.RequestID` | 不得在 Connect interceptor 中重新生成第二套 request_id |
| 日志 | `chi` middleware `chimw.Logger` | 不得在 Connect implementation 中独立打印第二套请求日志 |
| Panic 恢复 | `chi` middleware `chimw.Recoverer` | 不得在 Connect interceptor 中重新实现 panic recovery |
| 超时 | `chi` middleware `chimw.Timeout` | 不得在 Connect 层重复设置 context deadline |
| CORS | `chi` middleware `corsMiddleware` | 不得在 Connect 层重复设置 CORS 头 |
| RPC 级错误归一化 | `platform.MapToConnectError()` | 不得在模块 handler 中各自维护独立错误映射表 |
| 请求校验 | `Connect interceptor`（option） | 不得在 chi middleware 和 Connect interceptor 中重复校验 |
| 审计/埋点 | `Connect interceptor`（扩展位） | 不得在 chi middleware 中实现 RPC 级审计 |

### 2.3 Connect Interceptor 配置

```go
// backend/internal/platform/server.go

import (
    "connectrpc.com/connect"
    "connectrpc.com/validate"
)

func mountConnectModuleRegistry(r chi.Router, ...) {
    connectSvc := moduleregistryconnect.NewModuleRegistryServer(querySvc, commandSvc)

    path, handler := module_registryv1connect.NewModuleRegistryServiceHandler(
        connectSvc,
        // 全局 interceptor：错误归一化
        connect.WithInterceptors(&errorNormalizerInterceptor{}),
        // 可选：protovalidate 校验
        // connect.WithInterceptors(validate.NewInterceptor()),
    )
    r.Handle(path, handler)
}
```

### 2.4 Compat Handler 约束

Compat JSON handler 在迁移期复用同一 `chi` 外壳治理基线，但：
- 不得自建第二套错误包装
- 不得自建第二套日志治理
- 其 `writeError` 函数仅作为过渡兼容，不升级为长期错误映射入口

---

## 3. 设计：Procedure Path、Route Group 与 Router 结构调整

### 3.1 当前 router.go 结构 → 迁移后

```
当前 (router.go ~2000 lines)              迁移后
─────────────────────────────────         ─────────────────────────
server.go                                  server.go（装配根）
├── chi 根 middleware                      ├── chi 根 middleware（不变）
├── GET /healthz                          ├── GET /healthz（不变）
└── /api Route                            └── /api Route
    ├── mountModuleRegistry                   ├── mountConnectModuleRegistry
    │   ├── r.Get("/modules", ...)            │   └── r.Handle(path, connectHandler)
    │   ├── r.Get("/modules/{id}", ...)       ├── mountConnectDecisionCenter
    │   ├── r.Post("/modules", ...)           │   └── r.Handle(path, connectHandler)
    │   ├── r.Post("/modules/{id}/release"...)│   ├── mountConnectProductRegistry
    │   ├── r.Post("/modules/{id}/bindings/...│   │   └── r.Handle(path, connectHandler)
    │   ├── r.Get("/candidates/products",...) │   ├── mountConnectRepositoryBinding
    │   └── r.Get("/candidates/repos",...)    │   │   └── r.Handle(path, connectHandler)
    ├── mountDecisionCenter                   │   ├── mountConnectDashboard
    │   ├── r.Get("/decisions", ...)          │   │   └── r.Handle(path, connectHandler)
    │   ├── r.Get("/decisions/{id}", ...)     │   ├── mountConnectOnboarding
    │   ├── r.Post("/decisions", ...)         │   │   └── r.Handle(path, connectHandler)
    │   └── r.Post("/decisions/{id}/link"...) │   ├── mountConnectExport
    ├── mountProductRegistry                  │   │   └── r.Handle(path, connectHandler)
    │   ├── r.Get("/products", ...)           │   ├── mountConnectBackup
    │   ├── r.Get("/products/{id}", ...)      │   │   └── r.Handle(path, connectHandler)
    │   ├── r.Post("/products", ...)          │   ├── mountConnectReuseSummary
    │   └── r.Post("/products/{id}/bind"...)  │   │   └── r.Handle(path, connectHandler)
    ├── mountRepositoryBinding                │   └── mountCompatRoutes（过渡）
    │   ├── r.Get("/repositories", ...)       │       ├── r.Get("/candidates/products", ...)
    │   ├── r.Get("/repositories/{id}", ...)  │       ├── r.Get("/candidates/repositories", ...)
    │   ├── r.Post("/repositories", ...)      │       ├── r.Post("/modules/{id}/bindings/products", ...)
    │   ├── r.Post("/repositories/{id}/bind"..│       └── r.Post("/modules/{id}/bindings/repositories", ...)
    │   └── r.Post("/repositories/{id}/map"..)│
    ├── mountDashboard                        │
    ├── mountOnboarding                       │
    ├── mountExport                           │
    ├── mountBackup                           │
    └── mountReuseSummary                     │
```

### 3.2 Procedure Path 完整映射

| 模块 | 外部访问 URL | 内部 Connect procedure |
|------|-------------|----------------------|
| Module Registry | `/api/psco.module_registry.v1.ModuleRegistryService/ListModules` | generated path |
| Module Registry | `/api/psco.module_registry.v1.ModuleRegistryService/GetModuleDetail` | generated path |
| Module Registry | `/api/psco.module_registry.v1.ModuleRegistryService/CreateModule` | generated path |
| ... | ... | ... |

全部 34 条业务 RPC 的完整映射见 `phase07-01 frozen_scope.md` §2。

### 3.3 Router 职责分工

| 文件 | 迁移前 | 迁移后 |
|------|--------|--------|
| `server.go` | 构造 `chi.NewRouter()`、应用 middleware、创建 `/api` Route、调用 `mount*` | 新增 `mountConnect*` 调用，保留 `build*` 模式 |
| `router.go` | 2000+ 行手写 `r.Get/Post` 业务路径注册 | 缩减为：`mountCompatRoutes`（4 条 compat 过渡路由）+ `healthz` + infra 扩展位 |

### 3.4 禁止

- 不得在 `router.go` 中继续新增手写 `r.Get / r.Post` 业务路径
- 不得新增第二个 `/rpc`、`/grpc`、`/connect` 业务前缀
- 不得把 compat 过渡路由与 Connect canonical mount 混写在同一个 `mount*` 函数中
- 不得为每个模块单独创建第二层 `chi.Router` 子路由

---

## 4. 设计：Domain Error → Proto Error Code → Connect Error 单值映射

### 4.1 当前错误模式盘点

当前各模块通过 `handler/response.go` 中的 `writeError` 函数将 sentinel error 映射为 HTTP 状态码：

| 错误类别 | 哨兵错误示例 | 当前 HTTP 状态 | 目标 Connect Code |
|---------|-------------|---------------|-------------------|
| 资源不存在 | `ErrModuleNotFound`、`ErrProductNotFound`、`ErrRepositoryNotFound`、`ErrDecisionNotFound` | 404 | `CodeNotFound` |
| 重复/冲突 | `ErrDuplicateModuleName`、`ErrDuplicateBinding`、`ErrDuplicateReleaseVersion`、`ErrDuplicateLink` | 409 | `CodeAlreadyExists` |
| 非法输入 | `ErrInvalidInput`、`ErrInvalidStatus`、`ErrInvalidTargetType` | 400 | `CodeInvalidArgument` |
| 前置条件不满足 | `ErrModuleNotActive`、`ErrProductNotActive` | 400 | `CodeFailedPrecondition` |
| 内部错误 | 其他未分类错误 | 500 | `CodeInternal` |

### 4.2 单值映射函数

```go
// backend/internal/platform/connect_errors.go

package platform

import (
    "errors"

    "connectrpc.com/connect"

    "github.com/psco/backend/internal/backup"
    "github.com/psco/backend/internal/decisioncenter"
    "github.com/psco/backend/internal/export"
    "github.com/psco/backend/internal/moduleregistry"
    "github.com/psco/backend/internal/onboarding"
    "github.com/psco/backend/internal/productregistry"
    "github.com/psco/backend/internal/repositorybinding"
    "github.com/psco/backend/internal/reusesummary"
)

// MapToConnectError 将 domain/service error 映射为 Connect error。
// 这是整个后端唯一正式的错误映射承接位，各模块 Connect implementation
// 统一调用此函数，不得各自维护独立映射表。
func MapToConnectError(err error) *connect.Error {
    // Not Found
    if isNotFound(err) {
        return connect.NewError(connect.CodeNotFound, err)
    }
    // Already Exists / Conflict
    if isAlreadyExists(err) {
        return connect.NewError(connect.CodeAlreadyExists, err)
    }
    // Invalid Argument
    if isInvalidArgument(err) {
        return connect.NewError(connect.CodeInvalidArgument, err)
    }
    // Failed Precondition
    if isFailedPrecondition(err) {
        return connect.NewError(connect.CodeFailedPrecondition, err)
    }
    // Internal（兜底）
    return connect.NewError(connect.CodeInternal, err)
}

func isNotFound(err error) bool {
    return errors.Is(err, moduleregistry.ErrModuleNotFound) ||
        errors.Is(err, moduleregistry.ErrProductNotFound) ||
        errors.Is(err, moduleregistry.ErrRepositoryNotFound) ||
        errors.Is(err, decisioncenter.ErrDecisionNotFound) ||
        errors.Is(err, decisioncenter.ErrModuleNotFound) ||
        errors.Is(err, productregistry.ErrProductNotFound) ||
        errors.Is(err, productregistry.ErrModuleNotFound) ||
        errors.Is(err, repositorybinding.ErrRepositoryNotFound) ||
        errors.Is(err, repositorybinding.ErrProductNotFound) ||
        errors.Is(err, repositorybinding.ErrModuleNotFound) ||
        errors.Is(err, onboarding.ErrNotFound) ||
        errors.Is(err, export.ErrNotFound) ||
        errors.Is(err, backup.ErrNotFound) ||
        errors.Is(err, reusesummary.ErrNotFound)
}

func isAlreadyExists(err error) bool {
    return errors.Is(err, moduleregistry.ErrDuplicateModuleName) ||
        errors.Is(err, moduleregistry.ErrDuplicateBinding) ||
        errors.Is(err, moduleregistry.ErrDuplicateReleaseVersion) ||
        errors.Is(err, decisioncenter.ErrDuplicateLink) ||
        errors.Is(err, productregistry.ErrDuplicateBinding) ||
        errors.Is(err, repositorybinding.ErrDuplicateBinding) ||
        errors.Is(err, repositorybinding.ErrDuplicateMapping)
}

func isInvalidArgument(err error) bool {
    return errors.Is(err, moduleregistry.ErrInvalidInput) ||
        errors.Is(err, moduleregistry.ErrInvalidStatus) ||
        errors.Is(err, moduleregistry.ErrInvalidReleaseStatus) ||
        errors.Is(err, decisioncenter.ErrInvalidInput) ||
        errors.Is(err, decisioncenter.ErrInvalidStatus) ||
        errors.Is(err, decisioncenter.ErrInvalidTargetType) ||
        errors.Is(err, decisioncenter.ErrInvalidAlternatives) ||
        errors.Is(err, productregistry.ErrInvalidInput) ||
        errors.Is(err, productregistry.ErrInvalidStatus) ||
        errors.Is(err, repositorybinding.ErrInvalidInput) ||
        errors.Is(err, repositorybinding.ErrInvalidStatus) ||
        errors.Is(err, backup.ErrInvalidInput) ||
        errors.Is(err, export.ErrInvalidInput)
}

func isFailedPrecondition(err error) bool {
    return errors.Is(err, productregistry.ErrModuleNotActive) ||
        errors.Is(err, repositorybinding.ErrProductNotActive) ||
        errors.Is(err, repositorybinding.ErrModuleNotActive)
}
```

### 4.3 错误映射约束

| 规则 | 说明 |
|------|------|
| 单值映射位 | `platform/connect_errors.go:MapToConnectError` 是唯一映射函数 |
| 禁止模块级映射 | 各模块 Connect implementation 不得自建独立错误映射 |
| 禁止双主线 | 迁移后不得同时保留"旧 JSON HTTP 状态映射"和"新 Connect code 映射" |
| 新增错误 | 新增 domain error 时先在 `connect_errors.go` 的对应分类函数中注册 |
| Internal 兜底 | 未匹配的 domain error 统一映射为 `CodeInternal`，记录日志 |

---

## 5. 设计：Connect 迁移后 Service 分层保持策略

### 5.1 分层不变

| 层 | 迁移前 | 迁移后 | 变化 |
|----|--------|--------|------|
| `repository` | `backend/internal/<module>/repository/` | 不变 | 无 |
| `candidate` | `backend/internal/<module>/candidate/` | 不变 | 无 |
| `service` | `backend/internal/<module>/service/` | 不变 | 无 |
| `handler` | `backend/internal/<module>/handler/` | compat 过渡期保留，phase07-09/10 退场 | 退场 |
| `connect` | 不存在 | `backend/internal/<module>/connect/` | 新增 |

### 5.2 Connect Implementation 职责边界

Connect implementation（`<module>/connect/server.go`）的职责**仅限于**：

1. 接收 proto request
2. 调用既有 `service.QueryService` / `service.CommandService`
3. 进行 proto ↔ domain 类型转换
4. 在固定映射位返回 Connect response / error

**禁止**：
- 在 Connect implementation 中直接调用 repository store
- 在 Connect implementation 中拼装跨模块 SQL
- 在 Connect implementation 中实现业务校验规则
- 因 transport 迁移引入第二套 `service` 或 `application service` 命名体系

### 5.3 类型转换层

proto ↔ domain 类型转换函数放在 `<module>/connect/` 包内作为私有 helper：

```go
// backend/internal/moduleregistry/connect/convert.go

func moduleListItemToProto(item moduleregistry.ModuleListItem) *pb.ModuleListItem {
    return &pb.ModuleListItem{
        Id:                  item.ID,
        Name:                item.Name,
        Description:         item.Description,
        Status:              item.Status,
        LatestRelease:       item.LatestRelease,
        ProductBindCount:    int32(item.ProductBindCount),
        RepositoryBindCount: int32(item.RepositoryBindCount),
    }
}

func moduleToProto(m *moduleregistry.Module) *pb.Module {
    return &pb.Module{
        Id:          m.ID,
        Name:        m.Name,
        Description: m.Description,
        Status:      m.Status,
        CreatedAt:   timestamppb.New(m.CreatedAt),
        UpdatedAt:   timestamppb.New(m.UpdatedAt),
    }
}
```

---

## 6. 设计：Compat 过渡组在 Router 中的显式分组

### 6.1 分组方式

```go
// backend/internal/platform/router.go（修改后）

// mountCompatRoutes 承接 phase07 迁移期的 compat 过渡路由。
// 这些路由按 phase07-03 frozen_scope.md §2 退场清单在 phase07-09/10 删除。
// 不得与 canonical Connect mount 混写为同一长期组织模式。
func mountCompatRoutes(
    r chi.Router,
    pool *pgxpool.Pool,
    productQuerySvc *productregistrysvc.QueryService,
    productCommandSvc *productregistrysvc.CommandService,
    repoQuerySvc *repositorybindingsvc.QueryService,
    repoCommandSvc *repositorybindingsvc.CommandService,
) {
    // 构造 compat handler（复用既有 handler 模式）
    compatHandler := newCompatHandler(pool, productQuerySvc, productCommandSvc, repoQuerySvc, repoCommandSvc)

    // 候选读取 compat — 最晚 phase07-09 退场
    r.Get("/candidates/products", compatHandler.ListProductCandidates)
    r.Get("/candidates/repositories", compatHandler.ListRepositoryCandidates)

    // Module-centered 绑定 compat — 最晚 phase07-10 退场
    r.Post("/modules/{moduleId}/bindings/products", compatHandler.BindModuleToProduct)
    r.Post("/modules/{moduleId}/bindings/repositories", compatHandler.MapModuleToRepository)
}
```

### 6.2 分组约束

- compat 路由**必须**集中在一个 `mountCompatRoutes` 函数中
- compat 路由**不得**散落在各 `mountConnect*` 函数中
- 每个 compat 路由必须带有注释标注退场时点（phase07-09 或 phase07-10）
- `mountCompatRoutes` 函数整体在 phase07-11 收口前删除

---

## 7. 与上游文档一致性声明

| 上游文档 | 关键结论 | 本设计对齐 |
|---------|---------|-----------|
| `phase07-01 frozen_scope.md` §2 | 34 RPC 迁移总表 | §3.2 全部 34 条业务 RPC 通过 procedure path 映射 |
| `phase07-02 frozen_scope.md` §1-2 | chi 职责 + 单一 /api mount | §2 两层横切 + §3 单一 /api 业务树 |
| `phase07-02 frozen_scope.md` §3 | 3 插件 buf.gen.yaml | §1.1 接线图引用 generated handler |
| `phase07-02 frozen_scope.md` §2.4 | `simple=true` 选项 | §1.4 Connect 实现签名使用 `*pb.Request` / `*pb.Response` |
| `phase07-03 frozen_scope.md` §2-4 | compat inventory + 退场时点 | §6 compat 分组 + 注释标注退场时点 |
| `phase07-03 frozen_scope.md` §5 | 双门槛收口 | §4 单值错误映射为退场提供前提 |
| `shared_baseline.md` §2.3 | chi + ConnectRPC 后端主线 | 全文 |
| `architecture_plan.md` §4.3-4.5 | chi 装配层 + Connect 传输 | §1 接线 + §3 router 结构 |
| Context7 connect-go | `New*ServiceHandler` 返回 `(path, handler)` | §1.5 `r.Handle(path, handler)` |
| 现有代码 | `build*` 模式 + `QueryService`/`CommandService` | §1.5 保持 `build*` 模式，Connect 实现只做 transport 解包 |
