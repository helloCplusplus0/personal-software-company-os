# Phase07 传输主线完全切换正式规格 v0.1

> **文档定位**：`phase07` 传输主线完全切换的唯一正式规格正文。
> **角色**：`phase07-08 ~ 11`（后端实现、前端实现、compat 退场、验收收口）的唯一直接上游规格入口。
> **冻结日期**：2026-08-11
> **上游继承**：`phase07-01 ~ 06` 的冻结与设计结论（本规格生效后，它们退为冻结来源与证据链，不再承担并列直接执行入口职责）
> **下一阶段**：`phase07-08`（后端 Connect handler 实现）

---

## 1. 文档定位与上游继承

### 1.1 角色

本规格是 `phase07` 传输主线完全切换的**唯一正式规格正文**。后续 `phase07-08 ~ 11` 的实现、退场、验收与收口核销，必须以本规格为唯一直接上游。

### 1.2 上游继承链

```
phase07-01 frozen_scope.md          ─┐
phase07-02 frozen_scope.md           │
phase07-03 frozen_scope.md           ├─ 冻结来源与证据链（本规格生效后退为非直接执行入口）
phase07-04 design.md                 │
phase07-05 design.md                 │
phase07-06 design.md                ─┘
         │
         ▼
transport_mainline_cutover_spec_v0.1.md  ← 唯一正式规格正文
         │
         ▼
phase07-08 ~ 11（实现、退场、验收、收口） ← 唯一直接上游
```

### 1.3 禁止

- 后续实现不得绕开本规格，继续并列引用 `phase07-01 ~ 06` 作为长期直接执行入口
- 验收时不得以"设计文档已有"为由跳过本规格的正式口径
- 本规格不得重新定义与 `phase07-01 ~ 06` 冲突的第二套迁移边界或第二套验收门槛

---

## 2. 迁移范围与当前阶段边界

### 2.1 Canonical 业务模块（9 个）

| # | Canonical 业务模块 | Proto 包 | 对应页面/业务主线 |
|---|-------------------|----------|------------------|
| 1 | Module Registry | `psco.module_registry.v1` | Module List / Detail / Create / Release |
| 2 | Decision Center | `psco.decision_center.v1` | Decision List / Detail / Create / Link |
| 3 | Product Registry | `psco.product_registry.v1` | Product List / Detail / Create / Bind Module |
| 4 | Repository Binding | `psco.repository_binding.v1` | Repository List / Detail / Create / Bind Product / Map Module |
| 5 | Dashboard | `psco.dashboard.v1` | Dashboard Home (Overview / Feedback / Activity) |
| 6 | Onboarding | `psco.onboarding.v1` | Onboarding Page (FirstRunState) |
| 7 | Export | `psco.export.v1` | Dashboard Export (Snapshot / Execute) |
| 8 | Backup | `psco.backup.v1` | Dashboard Backup (Snapshot / Execute) |
| 9 | Reuse Summary | `psco.reuse_summary.v1` | Dashboard / Module Detail / Product Detail (ReuseSummary) |

### 2.2 Canonical RPC 迁移总表（34 条）

| # | Service | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 页面/动作 Owner |
|---|---------|-----|---------------|------|----------------------|----------------|
| 1 | ModuleRegistry | `ListModules` | `/api/modules` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListModules` | `modules/index.tsx` |
| 2 | ModuleRegistry | `GetModuleDetail` | `/api/modules/{moduleId}` | GET | `/psco.module_registry.v1.ModuleRegistryService/GetModuleDetail` | `modules/$moduleId.tsx` |
| 3 | ModuleRegistry | `CreateModule` | `/api/modules` | POST | `/psco.module_registry.v1.ModuleRegistryService/CreateModule` | `application/use-create-draft-module.ts` |
| 4 | ModuleRegistry | `CreateRelease` | `/api/modules/{moduleId}/releases` | POST | `/psco.module_registry.v1.ModuleRegistryService/CreateRelease` | `application/use-create-release.ts`（新增） |
| 5 | DecisionCenter | `ListDecisions` | `/api/decisions` | GET | `/psco.decision_center.v1.DecisionCenterService/ListDecisions` | `decisions/index.tsx` |
| 6 | DecisionCenter | `GetDecisionDetail` | `/api/decisions/{decisionId}` | GET | `/psco.decision_center.v1.DecisionCenterService/GetDecisionDetail` | `decisions/$decisionId.tsx` |
| 7 | DecisionCenter | `ListDecisionModuleCandidates` | `/api/decisions/{decisionId}/candidates/modules` | GET | `/psco.decision_center.v1.DecisionCenterService/ListDecisionModuleCandidates` | `decisions/$decisionId.tsx` link dialog |
| 8 | DecisionCenter | `CreateDecision` | `/api/decisions` | POST | `/psco.decision_center.v1.DecisionCenterService/CreateDecision` | `application/use-create-draft-decision.ts` |
| 9 | DecisionCenter | `LinkDecisionToTarget` | `/api/decisions/{decisionId}/links` | POST | `/psco.decision_center.v1.DecisionCenterService/LinkDecisionToTarget` | `application/use-link-decision-to-target.ts`（新增） |
| 10 | ProductRegistry | `ListProducts` | `/api/products` | GET | `/psco.product_registry.v1.ProductRegistryService/ListProducts` | `products/index.tsx` |
| 11 | ProductRegistry | `GetProductDetail` | `/api/products/{productId}` | GET | `/psco.product_registry.v1.ProductRegistryService/GetProductDetail` | `products/$productId.tsx` |
| 12 | ProductRegistry | `ListProductModuleCandidates` | `/api/products/{productId}/candidates/modules` | GET | `/psco.product_registry.v1.ProductRegistryService/ListProductModuleCandidates` | `products/$productId.tsx` bind dialog |
| 13 | ProductRegistry | `CreateProduct` | `/api/products` | POST | `/psco.product_registry.v1.ProductRegistryService/CreateProduct` | `application/use-create-draft-product.ts` |
| 14 | ProductRegistry | `BindModuleToProduct` | `/api/products/{productId}/bindings/modules` | POST | `/psco.product_registry.v1.ProductRegistryService/BindModuleToProduct` | `application/use-bind-module-to-product.ts`（新增） |
| 15 | RepositoryBinding | `ListRepositories` | `/api/repositories` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositories` | `repositories/index.tsx` |
| 16 | RepositoryBinding | `GetRepositoryDetail` | `/api/repositories/{repositoryId}` | GET | `/psco.repository_binding.v1.RepositoryBindingService/GetRepositoryDetail` | `repositories/$repositoryId.tsx` |
| 17 | RepositoryBinding | `ListRepositoryProductCandidates` | `/api/repositories/{repositoryId}/candidates/products` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositoryProductCandidates` | `repositories/$repositoryId.tsx` bind dialog |
| 18 | RepositoryBinding | `ListRepositoryModuleCandidates` | `/api/repositories/{repositoryId}/candidates/modules` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositoryModuleCandidates` | `repositories/$repositoryId.tsx` map dialog |
| 19 | RepositoryBinding | `CreateRepository` | `/api/repositories` | POST | `/psco.repository_binding.v1.RepositoryBindingService/CreateRepository` | `application/use-create-draft-repository.ts` |
| 20 | RepositoryBinding | `BindRepositoryToProduct` | `/api/repositories/{repositoryId}/bindings/products` | POST | `/psco.repository_binding.v1.RepositoryBindingService/BindRepositoryToProduct` | `application/use-bind-repository-to-product.ts`（新增） |
| 21 | RepositoryBinding | `MapModuleToRepository` | `/api/repositories/{repositoryId}/bindings/modules` | POST | `/psco.repository_binding.v1.RepositoryBindingService/MapModuleToRepository` | `application/use-map-module-to-repository.ts`（新增） |
| 22 | Dashboard | `GetDashboardOverview` | `/api/dashboard/overview` | GET | `/psco.dashboard.v1.DashboardService/GetDashboardOverview` | `data/use-dashboard-overview-read.ts`（新增） |
| 23 | Dashboard | `GetFeedbackSignals` | `/api/dashboard/feedback-signals` | GET | `/psco.dashboard.v1.DashboardService/GetFeedbackSignals` | `data/use-feedback-signals-read.ts`（新增） |
| 24 | Dashboard | `GetRecentActivities` | `/api/dashboard/recent-activities` | GET | `/psco.dashboard.v1.DashboardService/GetRecentActivities` | `data/use-recent-activities-read.ts`（新增） |
| 25 | Onboarding | `GetFirstRunState` | `/api/onboarding/state` | GET | `/psco.onboarding.v1.OnboardingService/GetFirstRunState` | `data/use-onboarding-read.ts`（替换 transport） |
| 26 | ReuseSummary | `GetReuseSummary` | `/api/reuse-summary` | GET | `/psco.reuse_summary.v1.ReuseSummaryService/GetReuseSummary` | `data/use-reuse-summary-read.ts`（替换 transport） |
| 27 | Export | `GetExportSnapshot` | `/api/dashboard/export` | GET | `/psco.export.v1.ExportService/GetExportSnapshot` | `dashboard/components/sovereignty-panel.tsx` |
| 28 | Export | `ExportCoreAssets` | `/api/dashboard/export` | POST | `/psco.export.v1.ExportService/ExportCoreAssets` | `dashboard/components/sovereignty-panel.tsx`（过渡位） |
| 29 | Backup | `GetBackupSnapshot` | `/api/dashboard/backup` | GET | `/psco.backup.v1.BackupService/GetBackupSnapshot` | `dashboard/components/sovereignty-panel.tsx` |
| 30 | Backup | `CreateInstanceBackup` | `/api/dashboard/backup` | POST | `/psco.backup.v1.BackupService/CreateInstanceBackup` | `dashboard/components/sovereignty-panel.tsx`（过渡位） |
| 31 | ModuleRegistry | `BindModuleToProduct` | `/api/modules/{moduleId}/bindings/products` | POST | `/psco.module_registry.v1.ModuleRegistryService/BindModuleToProduct` | compat facade（非 canonical business owner） |
| 32 | ModuleRegistry | `MapModuleToRepository` | `/api/modules/{moduleId}/bindings/repositories` | POST | `/psco.module_registry.v1.ModuleRegistryService/MapModuleToRepository` | compat facade（非 canonical business owner） |
| 33 | ModuleRegistry | `ListProductCandidates` | `/api/candidates/products` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListProductCandidates` | compat facade（非 canonical business owner） |
| 34 | ModuleRegistry | `ListRepositoryCandidates` | `/api/candidates/repositories` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListRepositoryCandidates` | compat facade（非 canonical business owner） |

> **注**：RPC #31-34 是 ModuleRegistryService 下仍带 compat 语义的 RPC。它们在 `.proto` 中作为 transport inventory 存在，但正式业务 owner 分别归属 Product Registry 和 Repository Binding。Connect handler 实现为 compat 薄壳，不作为 Module Registry 的 canonical business owner。

### 2.3 Legacy / Compat 业务入口（4 条，Phase 收口前必须退场）

| # | Legacy 入口 | 替代 Connect Path | 最晚退场 |
|---|------------|------------------|---------|
| L1 | `GET /api/candidates/products` | `ProductRegistryService.ListProducts` | phase07-09 |
| L2 | `GET /api/candidates/repositories` | `RepositoryBindingService.ListRepositories` | phase07-09 |
| L3 | `POST /api/modules/{moduleId}/bindings/products` | `ProductRegistryService.BindModuleToProduct` | phase07-10 |
| L4 | `POST /api/modules/{moduleId}/bindings/repositories` | `RepositoryBindingService.MapModuleToRepository` | phase07-10 |

### 2.4 非业务基础设施端点 Keep List（不纳入迁移）

| 端点 | 路径 | 方法 | 职责 |
|------|------|------|------|
| healthz | `/healthz` | GET | 简单健康检查，返回 `{"status":"ok"}` |
| readyz | `/readyz` | GET | 就绪检查（含 DB ping），待实现 |
| metrics | `/metrics` | GET | Prometheus metrics，待实现 |
| debug/pprof | `/debug/pprof/*` | GET | Go pprof 性能分析，待实现 |

这些端点**不进入 `.proto`**，**不通过 Connect 传输**，**不承担业务合同定义职责**。

### 2.5 当前阶段边界

- **纳入**：34 条 canonical RPC 全部切换到 ConnectRPC 正式传输主线
- **纳入**：4 条 legacy / compat 入口全部退场
- **纳入**：前端 8 个 hand-written `api-adapter.ts` 全部删除
- **不纳入**：readyz / metrics / debug/pprof 实现
- **不纳入**：Caddy 部署
- **不纳入**：CI workflow 正式配置（当前缺口显式建账，phase 收口时提供本地等价证据）

---

## 3. 合同与工具链

### 3.1 `.proto` 唯一合同源

- `.proto` 是唯一长期合同源；任何对外字段、枚举、响应 envelope 与错误语义都必须以 `.proto` 为准
- 业务 transport 正式切换到 `ConnectRPC`，不再保留 hand-written JSON business handler 作为长期主线
- `chi` 不再承担业务合同定义职责

### 3.2 单一 `/api` 访问基址

```
浏览器
  │
  │ GET/POST /api/psco.module_registry.v1.ModuleRegistryService/ListModules
  ▼
开发环境                       生产/验收环境
Vite dev server                Caddy 反代
proxy: /api → :8081            /api → :8081
  │                              │
  └──────────────┬───────────────┘
                 ▼
          Go HTTP Server (:8081)
          chi → /api mount → Connect handlers
```

| 环境 | /api 承接方式 | 文件 |
|------|-------------|------|
| 开发 | Vite dev proxy `/api` → `localhost:8081` | `frontend/vite.config.ts` |
| 验收 | 本地后端直接监听 8081，curl 直连 `/api` | 验收脚本 |
| 部署 | Caddy 反代 `/api` → 后端 8081 | 待部署 |

### 3.3 `buf` 生成链（3 插件矩阵）

| 插件 | 用途 | 产物落点 |
|------|------|---------|
| `buf.build/protocolbuffers/go` | Go protobuf 消息 | `backend/internal/gen/proto/` |
| `buf.build/connectrpc/gosimple` | Go Connect handler + client（simple 模式） | `backend/internal/gen/connect/` |
| `buf.build/bufbuild/es` | TypeScript protobuf + service descriptor | `frontend/src/gen/proto/` |

**工具链入口**：`proto/Makefile`（五个 target：`build / gen / lint / breaking / clean`），`make gen` 触发 `buf generate` 产出全部产物。

### 3.4 前端运行时依赖

| 包 | 用途 |
|----|------|
| `@connectrpc/connect` | 核心运行时：`createClient()`、错误类型 |
| `@connectrpc/connect-web` | 浏览器 transport：`createConnectTransport()` |
| `@bufbuild/protobuf` | protobuf-ES 运行时（已有依赖） |

### 3.5 CI 缺口

当前仓库不存在 `.github/workflows/` 目录。本阶段显式建账为缺口，不因 CI 缺失而阻断 phase 收口。替代方案：本地执行 `(cd proto && make build && make gen && make lint)` + `(cd backend && go build ./...)` + `(cd frontend && tsc -b --noEmit)` + `(cd frontend && npm run build)` 作为等价证据。

---

## 4. 后端正式传输主线

### 4.1 `chi` 的唯一正式职责

```
chi 根路由器
├── 根级 middleware（RequestID / RealIP / Logger / Recoverer / Timeout / CORS）
├── GET /healthz                          ← 保留 chi + net/http
├── GET /readyz                           ← 保留 chi + net/http（待实现）
├── GET /metrics                          ← 保留 chi + net/http（待实现）
├── GET /debug/pprof/*                    ← 保留 chi + net/http（待实现）
└── /api mount（r.Route("/api", ...)）    ← 仅作为 mount 外壳
    ├── Connect handler mount（业务接口）   ← 迁移后唯一业务 transport
    └── compat 过渡组（phase07-09/10 退场）← 临时过渡，不得长期保留
```

### 4.2 Connect Handler 挂载方式

Connect-Go 生成物为每个 service 产生 `New<Service>Handler(svc, opts...)` → 返回 `(path string, handler http.Handler)`。挂载时必须消费生成的 `path`：

```go
// 正式推荐：按 service 消费 generated path 后注册到 /api 子路由
r.Route("/api", func(r chi.Router) {
    path, handler := module_registryv1connect.NewModuleRegistryServiceHandler(moduleRegistrySvc, opts...)
    r.Handle(path, handler)

    path, handler = decision_centerv1connect.NewDecisionCenterServiceHandler(decisionCenterSvc, opts...)
    r.Handle(path, handler)

    // ... 其余 7 个 service 同理
})
```

### 4.3 分层保持

| 层 | 文件落点 | 迁移后变化 |
|----|---------|-----------|
| **repository** | `backend/internal/<module>/repository/` | 无变化 |
| **candidate** | `backend/internal/<module>/candidate/` | 无变化 |
| **service** | `backend/internal/<module>/service/` | 无变化，保持 QueryService / CommandService 分层 |
| **connect** | `backend/internal/<module>/connect/` | **新增**，Connect transport 实现 |
| **handler** | `backend/internal/<module>/handler/` | compat 过渡期保留，phase07-09/10 退场 |
| **platform** | `backend/internal/platform/` | 新增 `connect_errors.go`，修改 `server.go` |

### 4.4 Connect Implementation 职责边界

Connect implementation（`<module>/connect/server.go`）的职责**仅限于**：
1. 接收 proto request
2. 调用既有 `service.QueryService` / `service.CommandService`
3. 进行 proto ↔ domain 类型转换
4. 调用 `platform.MapToConnectError(err)` 返回 Connect error

**禁止**：在 Connect implementation 中直接调用 repository store、拼装跨模块 SQL、实现业务校验规则，或因 transport 迁移引入第二套 service 命名体系。

### 4.5 Domain Error → Connect Error 单值映射

| 错误类别 | Connect Code | 哨兵错误示例 |
|---------|-------------|-------------|
| 资源不存在 | `CodeNotFound` | `ErrModuleNotFound`、`ErrProductNotFound`、`ErrDecisionNotFound` 等 |
| 重复/冲突 | `CodeAlreadyExists` | `ErrDuplicateModuleName`、`ErrDuplicateBinding` 等 |
| 非法输入 | `CodeInvalidArgument` | `ErrInvalidInput`、`ErrInvalidStatus` 等 |
| 前置条件不满足 | `CodeFailedPrecondition` | `ErrModuleNotActive`、`ErrProductNotActive` 等 |
| 内部错误 | `CodeInternal` | 未分类错误兜底 |

**唯一映射承接位**：`platform/connect_errors.go:MapToConnectError`。各模块 Connect implementation 统一调用此函数，不得各自维护独立错误映射表。

### 4.6 Compat 路由过渡组

```go
// mountCompatRoutes 承接 phase07 迁移期的 compat 过渡路由。
// 按 phase07-03 §2 退场清单在 phase07-09/10 删除。
func mountCompatRoutes(r chi.Router) {
    // 候选读取 — 最晚 phase07-09 退场
    r.Get("/candidates/products", compatHandler.ListProductCandidates)
    r.Get("/candidates/repositories", compatHandler.ListRepositoryCandidates)
    // Module-centered 绑定 — 最晚 phase07-10 退场
    r.Post("/modules/{moduleId}/bindings/products", compatHandler.BindModuleToProduct)
    r.Post("/modules/{moduleId}/bindings/repositories", compatHandler.MapModuleToRepository)
}
```

---

## 5. 前端正式传输主线

### 5.1 调用组织架构

```
shared/rpc/
└── connect-transport.ts                     ← 唯一 cross-slice transport

features/<slice>/data/
├── connect-client.ts                        ← slice-local generated client
└── use-<entity>-read.ts                     ← slice-local read owner（纯只读）

features/<slice>/application/
└── use-<entity>-draft.ts                    ← slice-local mutation owner

pages / components
└── 消费 useXxxRead() / useXxxDraft()        ← 不得直接 import transport / client / api-adapter
```

### 5.2 共享 Transport（唯一落点）

```typescript
// frontend/src/shared/rpc/connect-transport.ts
import { createConnectTransport } from '@connectrpc/connect-web';

export const transport = createConnectTransport({ baseUrl: '/api' });
```

### 5.3 Slice-local Client（9 个切片）

| 切片 | Client 文件 | 生成的 Service 来源 |
|------|-----------|-------------------|
| Module Registry | `features/module-registry/data/connect-client.ts` | `@/gen/proto/psco/module_registry/v1/module_registry_pb` |
| Decision Center | `features/decision-center/data/connect-client.ts` | `@/gen/proto/psco/decision_center/v1/decision_center_pb` |
| Product Registry | `features/product-registry/data/connect-client.ts` | `@/gen/proto/psco/product_registry/v1/product_registry_pb` |
| Repository Binding | `features/repository-binding/data/connect-client.ts` | `@/gen/proto/psco/repository_binding/v1/repository_binding_pb` |
| Dashboard | `features/dashboard/data/connect-client.ts` | `@/gen/proto/psco/dashboard/v1/dashboard_pb` + export + backup |
| Onboarding | `features/onboarding/data/connect-client.ts` | `@/gen/proto/psco/onboarding/v1/onboarding_pb` |
| Reuse Summary | `features/reuse-summary/data/connect-client.ts` | `@/gen/proto/psco/reuse_summary/v1/reuse_summary_pb` |

### 5.4 Query 层约束

- **纯只读**：query 层只承接 `useQuery`，不混入 `useMutation`
- **单一 owner**：同一读接口的 queryKey / queryFn 只在 slice-local read owner 中定义一次
- **页面消费**：页面通过 `useXxxRead()` 消费，不重新拼装 queryKey/queryFn

### 5.5 前端正式写动作 Owner（11 项）

| # | 写动作 | Owner 位置 | 类型 | 最晚核销 |
|---|--------|-----------|------|---------|
| 1 | `CreateModule` | `application/use-create-draft-module.ts`（替换 transport） | ✅ Application Owner | phase07-07 |
| 2 | `CreateDecision` | `application/use-create-draft-decision.ts`（替换 transport） | ✅ Application Owner | phase07-07 |
| 3 | `CreateProduct` | `application/use-create-draft-product.ts`（替换 transport） | ✅ Application Owner | phase07-07 |
| 4 | `CreateRepository` | `application/use-create-draft-repository.ts`（替换 transport） | ✅ Application Owner | phase07-07 |
| 5 | `CreateRelease` | `application/use-create-release.ts`（**新增**） | 🔄 回收到 Owner | phase07-07 |
| 6 | `BindModuleToProduct` | `application/use-bind-module-to-product.ts`（**新增**） | 🔄 回收到 Owner | phase07-07 |
| 7 | `BindRepositoryToProduct` | `application/use-bind-repository-to-product.ts`（**新增**） | 🔄 回收到 Owner | phase07-07 |
| 8 | `MapModuleToRepository` | `application/use-map-module-to-repository.ts`（**新增**） | 🔄 回收到 Owner | phase07-07 |
| 9 | `LinkDecisionToTarget` | `application/use-link-decision-to-target.ts`（**新增**） | 🔄 回收到 Owner | phase07-07 |
| 10 | `ExportCoreAssets` | `dashboard/components/sovereignty-panel.tsx`（替换 transport） | ⏸️ 短时过渡位 | phase07-10 |
| 11 | `CreateInstanceBackup` | `dashboard/components/sovereignty-panel.tsx`（替换 transport） | ⏸️ 短时过渡位 | phase07-10 |

### 5.6 旧 Adapter 回收清单

| 优先级 | 文件 | 当前角色 |
|--------|------|---------|
| 🟢 优先 | `dashboard/data/api-adapter.ts` | 3 个 hand-written fetch |
| 🟢 优先 | `dashboard/data/sovereignty-api-adapter.ts` | 4 个 hand-written fetch |
| 🟢 优先 | `onboarding/data/api-adapter.ts` | 1 个 hand-written fetch |
| 🟢 优先 | `reuse-summary/data/api-adapter.ts` | 1 个 hand-written fetch |
| 🟠 中等 | `decision-center/data/api-adapter.ts` + adapter 壳 | 5 个 hand-written fetch |
| 🟠 中等 | `product-registry/data/api-adapter.ts` + adapter 壳 | 5 个 hand-written fetch |
| 🟠 中等 | `repository-binding/data/api-adapter.ts` + adapter 壳 | 7 个 hand-written fetch |
| 🔴 最后 | `module-registry/data/api-adapter.ts` + adapter 壳 + mock | 8 个 hand-written fetch + compat switch |

### 5.7 禁止

- 页面/组件直接 `import` transport / `createClient()` / `api-adapter`
- 在 `query` 层混入写动作
- 新增 `services/`、`sdk/`、`clients/` 根级目录与 `features/*/data|application` 长期并列
- 把只在单一切片内使用的能力过早提升到 `shared/`

---

## 6. 验收、退场与 Phase 收口标准

### 6.1 迁移核销（34 条 RPC）

34 条 canonical RPC（§2.2）必须逐条核销。每条 RPC 的核销方式：
- **后端**：Connect handler 已实现 → `curl` Connect procedure path 返回 200
- **前端**：对应页面/组件已消费 Connect client → 页面功能正常
- **最终**：在 `canonical_rpc_migration.csv` 中逐条标记 ✅

### 6.2 跨模块回归清单

| # | 联动路径 | 验证方式 |
|---|---------|---------|
| CR1 | 创建 Module → Product Detail bind → Module Detail bind_count 更新 | 端到端 + 刷新 |
| CR2 | 创建 Repository → Product Detail bind → Product Detail repository 更新 | 端到端 + 刷新 |
| CR3 | 创建 Module → Repository Detail map → Module Detail bind_count 更新 | 端到端 + 刷新 |
| CR4 | 创建 Decision + link → Decision Detail target 更新 → Module Detail decision 更新 | 端到端 + 刷新 |
| CR5 | Dashboard Overview 计数与实际情况一致 | 复位 → 创建 → 计数检查 |
| CR6 | Onboarding first_run → 完成 → Dashboard 不再显示 CTA | 端到端流程 |
| CR7 | Dashboard CTA → 各模块列表 → 详情 → 返回 | 路径遍历 |
| CR8 | Module Detail + Product Detail 的 Reuse Summary inline | 页面渲染 |
| CR9 | mutation 成功 → queryClient.invalidateQueries → 页面刷新 | 操作 + 刷新 |

### 6.3 Legacy / Compat Endpoint 退场证据

| # | Legacy 入口 | 退场时点 | 必需证据 |
|---|------------|---------|---------|
| L1 | `GET /api/candidates/products` | phase07-09 | 路由删除 + `ListProductCandidates` / `ProductCandidateReader` 删除 + `ProductRegistryService.ListProducts` Connect path 返回 200 + 仓库中无真实页面 / 组件继续通过 `module-registry` compat 候选读取导出承接正式读取 + `curl` 旧路径返回 404 |
| L2 | `GET /api/candidates/repositories` | phase07-09 | 路由删除 + `ListRepositoryCandidates` / `RepositoryCandidateReader` 删除 + `RepositoryBindingService.ListRepositories` Connect path 返回 200 + 仓库中无真实页面 / 组件继续通过 `module-registry` compat 候选读取导出承接正式读取 + `curl` 旧路径返回 404 |
| L3 | `POST /api/modules/{moduleId}/bindings/products` | phase07-10 | 路由删除 + `BindModuleToProduct` compat handler 删除 + `module-registry` `api-adapter.ts` 导出与 `module-registry-adapter.ts` real-api switch 删除 + `ProductModuleBindingPanel` 通过 `ProductRegistryService.BindModuleToProduct` Connect client 执行成功 + `curl` 旧路径返回 404 |
| L4 | `POST /api/modules/{moduleId}/bindings/repositories` | phase07-10 | 路由删除 + `MapModuleToRepository` compat handler 删除 + `module-registry` `api-adapter.ts` 导出与 `module-registry-adapter.ts` real-api switch 删除 + `RepositoryModuleMappingPanel` 通过 `RepositoryBindingService.MapModuleToRepository` Connect client 执行成功 + `curl` 旧路径返回 404 |

### 6.4 Phase07 收口双门槛

```
phase07 收口 = 门槛一（全部满足）AND 门槛二（全部满足）
```

**门槛一：Connect 主线已存在**

| 条件 | 验证方式 |
|------|---------|
| 34 条 RPC 均有 Connect handler | 逐条核销 |
| `buf.gen.yaml` 已包含 3 插件 | `buf generate` 成功 |
| 前端 Connect transport 已建立 | `createConnectTransport` 承接位存在 |
| 所有页面动作通过 Connect client 执行 | 联调验收 |

**门槛二：旧 JSON 主线已退场**

| 条件 | 验证方式 |
|------|---------|
| 4 条 legacy compat 入口全部核销 | 逐条检查删除证据 + 替代 Connect 回归证据 |
| `router.go` 中不再保留 compat 业务路由 | `grep` 验证 |
| `moduleregistry/handler/` 中不再保留 compat handler | `grep` 验证 |
| 前端真实业务主线不再经过 `module-registry` compat adapter | 代码审查 |
| 前端 13 项 adapter 资产已全部删除或按 mock 历史演示位显式隔离并完成核销 | 文件系统检查 + `frontend_adapter_retirement.csv` |
| `(cd backend && go build ./...)` 通过 | 编译检查 |
| `(cd frontend && tsc -b --noEmit)` 通过 | 编译检查 |
| `(cd frontend && npm run build)` 通过 | 构建检查 |

### 6.5 最终证据包结构

```
phase07 收口证据包
├── canonical_rpc_migration.csv            # 34 条 RPC 核销结果
├── legacy_endpoint_retirement.csv         # 4 条 legacy endpoint 退场核销
├── frontend_mutation_owner_acceptance.csv # 11 项 mutation owner 验收结果
├── frontend_adapter_retirement.csv        # 13 项 adapter 删除核销
├── cross_module_regression.csv            # 9 项跨模块联动回归结果
├── toolchain_migration.csv                # Vite / proto 生成链 / 本地脚本 / CI 缺口迁移核销
├── build_verification.txt                 # `(cd backend && go build ./...)` + `(cd frontend && tsc -b --noEmit)` + `(cd frontend && npm run build)` 通过
└── phase07_closure_decision.md            # 收口判定：双门槛全部满足 → 收口
```

### 6.6 阻断条件

以下任一条件不满足 → `phase07` 不得收口：

1. 34 条 canonical RPC 未全部迁移至 Connect
2. 4 条 legacy endpoint 未全部退场
3. 前端 11 项 mutation owner 未全部收口
4. 旧 JSON business handler 主线仍存在
5. 前端 13 项 adapter 资产未按 formal spec 完成删除或显式隔离核销
6. 跨模块回归 CR1-CR9 未全部通过
7. `(cd backend && go build ./...)` 或 `(cd frontend && tsc -b --noEmit)` 失败
8. `/api` 基址出现第二套并列前缀

---

## 7. 迁移执行顺序

```
phase07-07（本规格）          ← 正式规格正文收敛，后续唯一直接上游
       │
       ├── phase07-08（后端）  ← 9 模块 Connect handler 实现、platform 装配、router 调整
       │
       ├── phase07-09（后端退场）← L1/L2 候选读取 compat 退场
       │
       ├── phase07-10（前端退场）← L3/L4 绑定 compat 退场 + 13 项 adapter 删除核销
       │
       └── phase07-11（验收收口）← 34 RPC 核销 + 4 legacy 核销 + 11 owner 核销 + 双门槛判定
```

---

## 8. 一致性声明

本规格与以下文档保持单值一致：

| 文档 | 对齐结论 |
|------|---------|
| `phase07-01 frozen_scope.md` | 9 模块、34 RPC、4 legacy、infra keep list、收口标准 |
| `phase07-02 frozen_scope.md` | chi 职责、Connect mount、buf 3 插件、前端 client、/api 全链路 |
| `phase07-03 frozen_scope.md` | compat 策略、legacy 退场证据、双门槛、最晚时点 |
| `phase07-04 design.md` | 5 层接线链、chi/Connect 横切分层、router 结构、错误映射 |
| `phase07-05 design.md` | shared transport、slice-local client、query/application 边界、adapter 回收 |
| `phase07-06 design.md` | 34 RPC 迁移矩阵、跨模块回归、证据包、阻断条件 |
| `phase07 shared_baseline.md` | 技术栈、模块清单、验收基线 |
| `phase07 architecture_plan.md` | 迁移边界原则、前后端交付策略 |
| `TECH_STACK_BASELINE.md` | Durable System Track、Go + chi + ConnectRPC |
| `project_rules.md` §2.6 | `.proto` 唯一合同源、ConnectRPC 正式传输层 |
