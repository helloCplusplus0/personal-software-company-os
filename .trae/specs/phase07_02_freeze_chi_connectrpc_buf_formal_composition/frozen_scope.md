# Phase07-02 冻结产出：`chi + ConnectRPC + buf` 正式组合方式

> 本文档是 phase07-02 spec 的执行产出，冻结 `chi`、Connect handler、`buf` 生成链、前端客户端与 `/api` 访问链的正式组合方式。
> 产出日期：2026-08-11
> 上游：`phase07-01 frozen_scope.md`（34 条业务 RPC 迁移总表）、`shared_baseline.md`、`architecture_plan.md`
> 参考：Connect-Go v2 (`/connectrpc/connect-go`) + Connect-ES (`/connectrpc/connect-es`) 最新文档

---

## 1. 冻结：`chi` 的唯一正式职责

### 1.1 当前事实

从 [server.go](file:///home/dell/Projects/personal-software-company-os/backend/internal/platform/server.go) 和 [router.go](file:///home/dell/Projects/personal-software-company-os/backend/internal/platform/router.go) 中提取的当前状态：

| 事实 | 当前状态 | 迁移后 |
|------|---------|--------|
| 根级 middleware | `RequestID / RealIP / Logger / Recoverer / Timeout / CORS`（`server.go:41-46`） | 保持不变 |
| 健康检查 | `GET /healthz` → `platform/router.go:healthz`（`server.go:49`） | 保持不变 |
| `/api` 子路由 | `r.Route("/api", ...)` 包裹所有业务路由（`server.go:52`） | 保持为 mount 外壳 |
| 业务路由注册 | 8 个 `mount*` 函数直接注册 `r.Get/Post` 路径（`router.go`） | 改为 Connect handler mount |
| `readyz / metrics / debug` | 当前未实现 | 不纳入 phase07，保留为未来 `chi + net/http` 扩展位 |

### 1.2 冻结结论

`chi` 在 `phase07` 收口后的**唯一正式职责**：

```
chi 根路由器
├── 根级 middleware（RequestID / RealIP / Logger / Recoverer / Timeout / CORS）
├── GET /healthz                          ← 保留 chi + net/http（infra keep list）
├── GET /readyz                           ← 保留 chi + net/http（infra keep list，待实现）
├── GET /metrics                          ← 保留 chi + net/http（infra keep list，待实现）
├── GET /debug/pprof/*                    ← 保留 chi + net/http（infra keep list，待实现）
└── /api 子路由器（r.Route("/api", ...)）  ← 仅作为 mount 外壳
    ├── Connect handler mount（业务接口）   ← 迁移后唯一业务 transport
    └── （不再包含 hand-written JSON 业务 handler）
```

**禁止**：
- `chi` 不得继续作为 canonical 业务合同定义层
- 手写 `r.Get("/modules", ...)` 等路径声明不得作为长期业务 API 真相源
- `chi` 的 middleware 链不得被 Connect interceptor 复制为第二套横切逻辑体系

### 1.3 Infra Keep List 详细

| 端点 | 路径 | 方法 | 当前状态 | 职责 |
|------|------|------|---------|------|
| **healthz** | `/healthz` | GET | ✅ 已实现 | 简单健康检查，返回 `{"status":"ok"}` |
| **readyz** | `/readyz` | GET | ⬜ 待实现 | 就绪检查（含 DB ping），不纳入 phase07 |
| **metrics** | `/metrics` | GET | ⬜ 待实现 | Prometheus metrics 端点，不纳入 phase07 |
| **debug/pprof** | `/debug/pprof/*` | GET | ⬜ 待实现 | Go pprof 性能分析，不纳入 phase07 |

这些端点**不进入 `.proto`**，**不承担业务合同定义职责**，**不通过 Connect 传输**。

---

## 2. 冻结：Connect Handler 在单一 `/api` 前缀下的正式挂载方式

### 2.1 Connect-Go 生成物与挂载机制

基于 [Connect-Go 最新文档](https://github.com/connectrpc/connect-go)，Connect-Go 生成物会为每个 service 产生：

```
backend/internal/gen/connect/psco/<module>/v1/<module>v1connect/<module>.connect.go
```

关键生成物：
- `New<Service>Handler(svc, opts...)` → 返回 `(path string, handler http.Handler)` — 可直接挂载到 `net/http` mux
- `Unimplemented<Service>Handler` → 嵌入用空实现，返回 `CodeUnimplemented`
- **Procedure 常量**：`<Service>ListModulesProcedure` = `"/psco.module_registry.v1.ModuleRegistryService/ListModules"`

### 2.2 挂载到 `chi` 的方式

Connect handler 是标准 `http.Handler`，但 `phase07` 的正式挂载方式**不能**依赖“把每个 service handler 都 `Mount("/")` 到同一个 `chi` 子路由器”。

原因：

- `New<Service>Handler(...)` 已返回 `(path string, handler http.Handler)`，正式实现必须消费返回的 `path`
- `chi.Mount()` 在同一路径重复 mount 会冲突，因此不适合作为 9 个业务 service 的逐个复用承接位

**冻结结论**：在 `/api` 外壳下，以“**按 service 消费 generated path**”为正式挂载模式；若实现层希望只做一次 `Mount("/api", mux)`，也必须先把多个 Connect handlers 组合到单个 `http.ServeMux` 或等价聚合承接位，再统一挂入 `chi`

```go
// 正式推荐：按 service 消费 generated path，再注册到 /api 子路由
r.Route("/api", func(r chi.Router) {
    path, handler := module_registryv1connect.NewModuleRegistryServiceHandler(moduleRegistrySvc, opts...)
    r.Handle(path, handler)

    path, handler = decision_centerv1connect.NewDecisionCenterServiceHandler(decisionCenterSvc, opts...)
    r.Handle(path, handler)
})

// 允许的等价实现：先聚合，再统一挂入 /api
apiMux := http.NewServeMux()
path, handler := module_registryv1connect.NewModuleRegistryServiceHandler(moduleRegistrySvc, opts...)
apiMux.Handle(path, handler)
path, handler = decision_centerv1connect.NewDecisionCenterServiceHandler(decisionCenterSvc, opts...)
apiMux.Handle(path, handler)
r.Mount("/api", apiMux)
```

**禁止**：

- 把多个 business service handler 逐个 `Mount("/")` 到同一个 `chi` 子路由器
- 丢弃 generated `path`，再手写另一套路由字符串承接 canonical RPC

### 2.3 冻结：外部访问路径映射

| 层 | 路径形式 | 说明 |
|----|---------|------|
| `.proto` RPC | `ListModules` | 合同定义 |
| Connect procedure | `/psco.module_registry.v1.ModuleRegistryService/ListModules` | 生成常量 |
| chi mount shell | `/api` | 单一业务前缀 |
| **浏览器外部访问** | `/api/psco.module_registry.v1.ModuleRegistryService/ListModules` | **最终对外 URL** |

**禁止**：
- 新增 `/rpc`、`/connect`、`/grpc` 等并列业务前缀
- 把 Connect procedure path 的前缀剥离后再映射到旧 JSON 路径（如把 `/psco.../ListModules` 映射为 `/api/modules`）

### 2.4 冻结：`simple` 模式选择

Connect-Go 的 `protoc-gen-connect-go` 支持 `simple` 模式，生成更简洁的 handler 签名：

```go
// simple=false（默认）：需要 connect.Request/Response 包装
func (s *Server) ListModules(ctx context.Context, req *connect.Request[pb.ListModulesRequest]) (*connect.Response[pb.ListModulesResponse], error)

// simple=true：直接使用 proto 消息类型，与现有 service 层签名兼容
func (s *Server) ListModules(ctx context.Context, req *pb.ListModulesRequest) (*pb.ListModulesResponse, error)
```

**冻结选择**：采用 simple 模式，理由是与现有 Go service 层签名风格一致，减少迁移摩擦。

---

## 3. 冻结：`buf.gen.yaml` 正式插件矩阵与产物落点

### 3.1 当前状态

当前 [buf.gen.yaml](file:///home/dell/Projects/personal-software-company-os/proto/buf.gen.yaml) 只有两个插件：
- `buf.build/protocolbuffers/go` → `../backend/internal/gen/proto`
- `buf.build/bufbuild/es` → `../frontend/src/gen/proto`

**缺失**：Connect-Go simple 模式的正式远程插件配置。

### 3.2 冻结：正式插件矩阵（3 插件）

| 插件 | 用途 | 输出目录 | 产物 |
|------|------|---------|------|
| `buf.build/protocolbuffers/go` | Go protobuf 消息类型 | `../backend/internal/gen/proto` | `*.pb.go` |
| `buf.build/connectrpc/gosimple` | Go Connect handler + client（simple 模式） | `../backend/internal/gen/connect` | `*.connect.go` |
| `buf.build/bufbuild/es` | TypeScript protobuf + service descriptor | `../frontend/src/gen/proto` | `*_pb.ts` |

### 3.3 冻结：`buf.gen.yaml` 目标形态

```yaml
version: v2

managed:
  enabled: true
  override:
    - file_option: go_package_prefix
      value: github.com/psco/backend/internal/gen/proto

plugins:
  # 1. Go protobuf 消息类型（保持现有）
  - remote: buf.build/protocolbuffers/go
    out: ../backend/internal/gen/proto
    opt: paths=source_relative

  # 2. Go Connect handler + client（新增，simple 模式）
  - remote: buf.build/connectrpc/gosimple
    out: ../backend/internal/gen/connect
    opt:
      - paths=source_relative
      - simple

  # 3. TypeScript protobuf + service descriptor（保持现有）
  - remote: buf.build/bufbuild/es
    out: ../frontend/src/gen/proto
    opt:
      - target=ts
      - import_extension=js
```

### 3.4 冻结：产物目录结构

```
backend/internal/gen/
├── proto/                          # ← Go protobuf 消息（已有）
│   └── psco/
│       ├── module_registry/v1/     # *.pb.go
│       ├── decision_center/v1/
│       ├── product_registry/v1/
│       ├── repository_binding/v1/
│       ├── dashboard/v1/
│       ├── onboarding/v1/
│       ├── export/v1/
│       ├── backup/v1/
│       ├── reuse_summary/v1/
│       └── common/v1/
└── connect/                        # ← Go Connect handler（新增）
    └── psco/
        ├── module_registry/v1/module_registryv1connect/
        │   └── module_registry.connect.go
        ├── decision_center/v1/decision_centerv1connect/
        │   └── decision_center.connect.go
        ├── product_registry/v1/productregistryv1connect/
        │   └── product_registry.connect.go
        ├── repository_binding/v1/repositorybindingv1connect/
        │   └── repository_binding.connect.go
        ├── dashboard/v1/dashboardv1connect/
        │   └── dashboard.connect.go
        ├── onboarding/v1/onboardingv1connect/
        │   └── onboarding.connect.go
        ├── export/v1/exportv1connect/
        │   └── export.connect.go
        ├── backup/v1/backupv1connect/
        │   └── backup.connect.go
        └── reuse_summary/v1/reusesummaryv1connect/
            └── reuse_summary.connect.go

frontend/src/gen/proto/             # ← TS protobuf（已有，service descriptor 已包含）
└── psco/
    ├── module_registry/v1/         # *_pb.ts（含 service descriptor）
    ├── decision_center/v1/
    ├── ...
```

### 3.5 冻结：工具链入口

- `proto/Makefile`：继续作为唯一生成入口，`build / gen / lint / breaking / clean` 五个 target 不变
- `make gen` 或 `buf generate`：从 `proto/` 目录执行，产出上述全部产物
- 生成产物已加入 `.gitignore`，不进入版本控制
- **禁止**新增第二个 `buf.gen.yaml`、第二个 proto workspace、第二个 Go 生成根或第二个 TS 生成根

---

## 4. 冻结：前端客户端正式生成与运行时组合方式

### 4.1 当前状态

前端当前通过手写 `fetch + JSON DTO` adapter 消费业务接口：
- 8 个 `api-adapter.ts` 文件（phase07-01 §6.1 清单）
- 使用 `fetch()` + 手动 JSON 解析
- 无 Connect transport 运行时

### 4.2 冻结：正式前端客户端组合

基于 [Connect-ES 最新文档](https://github.com/connectrpc/connect-es)：

```
┌─────────────────────────────────────────────────────┐
│  前端业务切片（query / application）                  │
│  ├── query 层：调用 Connect client 读方法（纯只读）    │
│  └── application 层：调用 Connect client 写方法        │
├─────────────────────────────────────────────────────┤
│  共享 Connect transport 层                           │
│  ├── createConnectTransport({ baseUrl: '/api' })     │
│  └── @connectrpc/connect-web（浏览器 fetch transport）│
├─────────────────────────────────────────────────────┤
│  TS 生成产物（bufbuild/es）                           │
│  ├── *_pb.ts：消息类型 + service 描述符               │
│  └── createClient(ServiceDescriptor, transport)      │
└─────────────────────────────────────────────────────┘
```

### 4.3 冻结：创建客户端代码模式

```typescript
// frontend/src/shared/connect-transport.ts（共享 transport 承接位）
import { createConnectTransport } from '@connectrpc/connect-web';

export const transport = createConnectTransport({
  baseUrl: '/api',  // 与 Vite dev proxy 的 /api 前缀一致
});

// 业务切片中的使用示例（module-registry/data/）
import { createClient } from '@connectrpc/connect';
import { ModuleRegistryService } from '@/gen/proto/psco/module_registry/v1/module_registry_pb';
import { transport } from '@/shared/connect-transport';

const client = createClient(ModuleRegistryService, transport);

// 读：query 层
const { modules } = await client.listModules({ queryText: '', statusFilter: 0 });

// 写：application 层
const { module } = await client.createModule({ name: 'my-module' });
```

### 4.4 冻结：npm 依赖

| 包 | 版本约束 | 用途 |
|----|---------|------|
| `@connectrpc/connect` | `^2.0` | 核心运行时：`createClient()`、错误类型 |
| `@connectrpc/connect-web` | `^2.0` | 浏览器 transport：`createConnectTransport()` |
| `@bufbuild/protobuf` | `^2.0` | 已存在（protobuf-ES 运行时，已有依赖） |

### 4.5 冻结：迁移边界

- 前端 transport 迁移**只允许**发生在切片的 `query / application` 固定承接位
- **禁止**在 route 组件或展示组件中直接持有 `createClient()` 或 `transport`
- **禁止**新增第二个 TS 生成目录（如 `frontend/src/gen/connect/`）
- 旧 `api-adapter.ts` 在对应切片完成迁移后删除

---

## 5. 冻结：浏览器、Vite、本地启动链与 Caddy 的 `/api` 承接关系

### 5.1 当前事实

| 组件 | 当前状态 | 文件 |
|------|---------|------|
| 前端 API base URL | `VITE_API_BASE_URL` 默认为空 → 走同源 `/api` | [frontend/.env.example](file:///home/dell/Projects/personal-software-company-os/frontend/.env.example) |
| Vite dev proxy | `/api` → `http://localhost:8081` | [vite.config.ts](file:///home/dell/Projects/personal-software-company-os/frontend/vite.config.ts) |
| 后端监听 | 单一端口（默认 8081），`/api` 子路由在服务内承接 | [server.go](file:///home/dell/Projects/personal-software-company-os/backend/internal/platform/server.go) |
| Caddy | 未部署，计划中 | 待 phase07 后部署 |

### 5.2 冻结：各环境 `/api` 承接方式

```
                    浏览器
                      │
                      │ GET/POST /api/psco.module_registry.v1.ModuleRegistryService/ListModules
                      ▼
        ┌─────────────┴─────────────┐
        │                           │
    开发环境                    生产/验收环境
        │                           │
   Vite dev server              Caddy 反代
   proxy: /api → :8081          /api → :8081
        │                           │
        └─────────────┬─────────────┘
                      ▼
               Go HTTP Server (:8081)
               ┌─────────────────────┐
               │ chi 根路由器         │
               │ ├── /healthz (chi)  │
               │ └── /api mount      │
               │     └── Connect     │
               │         handlers    │
               └─────────────────────┘
```

### 5.3 冻结约束

- 浏览器始终向 `/api` 发起业务请求
- Vite dev proxy：开发期将 `/api` 转发到本地后端监听端口（`:8081`）
- 本地后端启动链：只负责监听单一 HTTP 端口，在服务内部通过 `chi` 的 `/api` 子路由承接
- Caddy：生产或验收环境继续以反代方式暴露同一 `/api`
- `VITE_API_BASE_URL` 只允许表达 origin（如 `https://psco.example.com`），不得把 Connect 迁移扩写成第二个业务 base path

---

## 6. 冻结：Connect Interceptor 与 Middleware 关系

### 6.1 职责边界

| 层 | 职责 | 实现 |
|----|------|------|
| `chi` middleware | HTTP 级横切：RequestID、日志、恢复、CORS、超时 | 保持现有 `chi` middleware 链 |
| Connect interceptor | RPC 级横切：校验、错误归一化、元数据处理 | `connect.WithInterceptors(...)` |

### 6.2 冻结

- `chi` middleware 继续作为统一 HTTP 外壳
- Connect interceptor 只承接 RPC 级横切逻辑，**不得复制**第二套请求治理体系
- domain error → proto error code → Connect error 的映射必须单值化

---

## 7. 与上游文档一致性声明

| 上游文档 | 关键结论 | 本产出对齐 |
|---------|---------|-----------|
| `shared_baseline.md` §2.3 | `chi + ConnectRPC` 后端主线 | §1 chi 职责冻结 |
| `shared_baseline.md` §2.4 | 单一 `/api` 前缀、infra keep list | §2, §5 |
| `shared_baseline.md` §5.3 | 3 插件矩阵 + 单一生成入口 | §3 |
| `architecture_plan.md` §4.3 | `chi` 装配层 + Connect 业务传输 | §1, §2 |
| `architecture_plan.md` §4.4 | 前端单一 transport + query/application 边界 | §4 |
| `architecture_plan.md` §4.5 | `buf.gen.yaml` 补齐 Connect-Go 生成链 + bufbuild/es | §3 |
| `phase07-01 frozen_scope.md` | 34 条业务 RPC 迁移总表 | 所有 RPC 通过 §2 的挂载方式承接 |
| `project_rules.md` §2.6 | `.proto` 唯一合同源 + ConnectRPC 正式传输 | 全文 |
