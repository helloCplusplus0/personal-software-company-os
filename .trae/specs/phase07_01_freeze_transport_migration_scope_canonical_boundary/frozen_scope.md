# Phase07-01 冻结范围产出

> 本文档是 phase07-01 spec 的执行产出，冻结 phase07 迁移范围、canonical 业务接口清单、非业务端点边界、legacy inventory 与收口判定标准。
> 产出时间：2026-08-11

## 1. 冻结：9 个 Canonical 业务模块

根据 `phase07` shared_baseline §5.1 与 architecture_plan §4.2，以下 9 个 canonical 业务模块必须在本阶段完成正式传输主线切换：

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

**禁止项**：
- 不得把"只迁部分模块、其余继续兼容"解释为 phase07 完成
- 不得把仍依赖 hand-written JSON business handler 的模块留到 mvp0.3 再处理
- 不得把"只迁新增业务接口"或"只迁单个模块试点"解释为完成

---

## 2. 冻结：Canonical 业务接口迁移总表

> 下钻到 `service / RPC / 当前入口路径 / 目标 Connect path / 当前 transport owner / 迁移后正式 owner / 页面或动作 owner` 级别。
> 本节总表以当前 `.proto` 已声明的 34 条业务 RPC 为 transport migration inventory；其中 `ModuleRegistryService` 下 4 条 module-centered binding / candidate RPC 虽仍在 `.proto` 中，但在当前源码里已经带有 compat / delegate 语义，phase07 实施中不得再把其 JSON HTTP 入口误写成 Module Registry 的长期 canonical owner。

### 2.1 Module Registry (`psco.module_registry.v1`)

> 说明：`module_registry.proto` 当前仍声明 8 条 RPC，因此全部纳入 transport migration inventory。
> 其中后 4 条在当前后端源码中已明确为 compat / delegate 语义：
> - `BindModuleToProduct` 实际委派到 Product Registry canonical command owner
> - `MapModuleToRepository` 实际委派到 Repository Binding canonical command owner
> - `ListProductCandidates` / `ListRepositoryCandidates` 仅保留历史兼容读取入口，不再承接长期业务 owner

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `ListModules` | `/api/modules` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListModules` | `moduleregistry/handler/query_handler.go:ListModules` | Connect handler + `service.QueryService` | `modules/index.tsx` (ModuleListPage) |
| 2 | `GetModuleDetail` | `/api/modules/{moduleId}` | GET | `/psco.module_registry.v1.ModuleRegistryService/GetModuleDetail` | `moduleregistry/handler/query_handler.go:GetModuleDetail` | Connect handler + `service.QueryService` | `modules/$moduleId.tsx` (ModuleDetailPage) |
| 3 | `CreateModule` | `/api/modules` | POST | `/psco.module_registry.v1.ModuleRegistryService/CreateModule` | `moduleregistry/handler/command_handler.go:CreateModule` | Connect handler + `service.CommandService` | `modules/new.tsx` → `application/use-create-draft-module.ts` |
| 4 | `CreateRelease` | `/api/modules/{moduleId}/releases` | POST | `/psco.module_registry.v1.ModuleRegistryService/CreateRelease` | `moduleregistry/handler/command_handler.go:CreateRelease` | Connect handler + `service.CommandService` | `module-registry/pages/release-create-page.tsx` |
| 5 | `BindModuleToProduct` | `/api/modules/{moduleId}/bindings/products` | POST | `/psco.module_registry.v1.ModuleRegistryService/BindModuleToProduct` | `moduleregistry/handler/command_handler.go:BindModuleToProduct` (兼容委派) | Connect handler（module-centered compat facade）+ Product Registry canonical command owner | `product-module-binding-panel.tsx` (ProductDetailPage)；`ModuleDetailPage` 仅保留兼容跳转入口 |
| 6 | `MapModuleToRepository` | `/api/modules/{moduleId}/bindings/repositories` | POST | `/psco.module_registry.v1.ModuleRegistryService/MapModuleToRepository` | `moduleregistry/handler/command_handler.go:MapModuleToRepository` (兼容委派) | Connect handler（module-centered compat facade）+ Repository Binding canonical command owner | `repository-module-mapping-panel.tsx` (RepositoryDetailPage)；`ModuleDetailPage` 仅保留兼容跳转入口 |
| 7 | `ListProductCandidates` | `/api/candidates/products` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListProductCandidates` | `moduleregistry/handler/query_handler.go:ListProductCandidates` (legacy compat) | Connect handler（compat facade）+ Product Registry candidate source | 当前前端无 active caller；`ModuleDetailPage` 已退化为兼容跳转入口 |
| 8 | `ListRepositoryCandidates` | `/api/candidates/repositories` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListRepositoryCandidates` | `moduleregistry/handler/query_handler.go:ListRepositoryCandidates` (legacy compat) | Connect handler（compat facade）+ Repository Binding candidate source | 当前前端无 active caller；`ModuleDetailPage` 已退化为兼容跳转入口 |

### 2.2 Decision Center (`psco.decision_center.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `ListDecisions` | `/api/decisions` | GET | `/psco.decision_center.v1.DecisionCenterService/ListDecisions` | `decisioncenter/handler/query_handler.go:ListDecisions` | Connect handler + `service.QueryService` | `decisions/index.tsx` (DecisionListPage) |
| 2 | `GetDecisionDetail` | `/api/decisions/{decisionId}` | GET | `/psco.decision_center.v1.DecisionCenterService/GetDecisionDetail` | `decisioncenter/handler/query_handler.go:GetDecisionDetail` | Connect handler + `service.QueryService` | `decisions/$decisionId.tsx` (DecisionDetailPage) |
| 3 | `ListDecisionModuleCandidates` | `/api/decisions/{decisionId}/candidates/modules` | GET | `/psco.decision_center.v1.DecisionCenterService/ListDecisionModuleCandidates` | `decisioncenter/handler/query_handler.go:ListDecisionModuleCandidates` | Connect handler + `service.QueryService` | `decisions/$decisionId.tsx` (DecisionDetailPage link dialog) |
| 4 | `CreateDecision` | `/api/decisions` | POST | `/psco.decision_center.v1.DecisionCenterService/CreateDecision` | `decisioncenter/handler/command_handler.go:CreateDecision` | Connect handler + `service.CommandService` | `decisions/new.tsx` → `application/use-create-draft-decision.ts` |
| 5 | `LinkDecisionToTarget` | `/api/decisions/{decisionId}/links` | POST | `/psco.decision_center.v1.DecisionCenterService/LinkDecisionToTarget` | `decisioncenter/handler/command_handler.go:LinkDecisionToTarget` | Connect handler + `service.CommandService` | `decision-center/components/decision-module-candidate-panel.tsx` (DecisionDetailPage) |

### 2.3 Product Registry (`psco.product_registry.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `ListProducts` | `/api/products` | GET | `/psco.product_registry.v1.ProductRegistryService/ListProducts` | `productregistry/handler/query_handler.go:ListProducts` | Connect handler + `service.QueryService` | `products/index.tsx` (ProductListPage) |
| 2 | `GetProductDetail` | `/api/products/{productId}` | GET | `/psco.product_registry.v1.ProductRegistryService/GetProductDetail` | `productregistry/handler/query_handler.go:GetProductDetail` | Connect handler + `service.QueryService` | `products/$productId.tsx` (ProductDetailPage) |
| 3 | `ListProductModuleCandidates` | `/api/products/{productId}/candidates/modules` | GET | `/psco.product_registry.v1.ProductRegistryService/ListProductModuleCandidates` | `productregistry/handler/query_handler.go:ListProductModuleCandidates` | Connect handler + `service.QueryService` | `products/$productId.tsx` (ProductDetailPage bind dialog) |
| 4 | `CreateProduct` | `/api/products` | POST | `/psco.product_registry.v1.ProductRegistryService/CreateProduct` | `productregistry/handler/command_handler.go:CreateProduct` | Connect handler + `service.CommandService` | `products/new.tsx` → `application/use-create-draft-product.ts` |
| 5 | `BindModuleToProduct` | `/api/products/{productId}/bindings/modules` | POST | `/psco.product_registry.v1.ProductRegistryService/BindModuleToProduct` | `productregistry/handler/command_handler.go:BindModuleToProduct` | Connect handler + `service.CommandService` | `product-registry/components/product-module-binding-panel.tsx` (ProductDetailPage) |

### 2.4 Repository Binding (`psco.repository_binding.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `ListRepositories` | `/api/repositories` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositories` | `repositorybinding/handler/query_handler.go:ListRepositories` | Connect handler + `service.QueryService` | `repositories/index.tsx` (RepositoryListPage) |
| 2 | `GetRepositoryDetail` | `/api/repositories/{repositoryId}` | GET | `/psco.repository_binding.v1.RepositoryBindingService/GetRepositoryDetail` | `repositorybinding/handler/query_handler.go:GetRepositoryDetail` | Connect handler + `service.QueryService` | `repositories/$repositoryId.tsx` (RepositoryDetailPage) |
| 3 | `ListRepositoryProductCandidates` | `/api/repositories/{repositoryId}/candidates/products` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositoryProductCandidates` | `repositorybinding/handler/query_handler.go:ListRepositoryProductCandidates` | Connect handler + `service.QueryService` | `repositories/$repositoryId.tsx` (RepositoryDetailPage bind dialog) |
| 4 | `ListRepositoryModuleCandidates` | `/api/repositories/{repositoryId}/candidates/modules` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositoryModuleCandidates` | `repositorybinding/handler/query_handler.go:ListRepositoryModuleCandidates` | Connect handler + `service.QueryService` | `repositories/$repositoryId.tsx` (RepositoryDetailPage map dialog) |
| 5 | `CreateRepository` | `/api/repositories` | POST | `/psco.repository_binding.v1.RepositoryBindingService/CreateRepository` | `repositorybinding/handler/command_handler.go:CreateRepository` | Connect handler + `service.CommandService` | `repositories/new.tsx` → `application/use-create-draft-repository.ts` |
| 6 | `BindRepositoryToProduct` | `/api/repositories/{repositoryId}/bindings/products` | POST | `/psco.repository_binding.v1.RepositoryBindingService/BindRepositoryToProduct` | `repositorybinding/handler/command_handler.go:BindRepositoryToProduct` | Connect handler + `service.CommandService` | `repository-binding/components/repository-product-binding-panel.tsx` (RepositoryDetailPage) |
| 7 | `MapModuleToRepository` | `/api/repositories/{repositoryId}/bindings/modules` | POST | `/psco.repository_binding.v1.RepositoryBindingService/MapModuleToRepository` | `repositorybinding/handler/command_handler.go:MapModuleToRepository` | Connect handler + `service.CommandService` | `repository-binding/components/repository-module-mapping-panel.tsx` (RepositoryDetailPage) |

### 2.5 Dashboard (`psco.dashboard.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `GetDashboardOverview` | `/api/dashboard/overview` | GET | `/psco.dashboard.v1.DashboardService/GetDashboardOverview` | `dashboard/handler/query_handler.go:GetOverview` | Connect handler + `service.QueryService` | `dashboard/pages/dashboard-home-page.tsx` via `dashboard/data/api-adapter.ts:fetchDashboardOverview` |
| 2 | `GetFeedbackSignals` | `/api/dashboard/feedback-signals` | GET | `/psco.dashboard.v1.DashboardService/GetFeedbackSignals` | `dashboard/handler/query_handler.go:GetFeedbackSignals` | Connect handler + `service.QueryService` | `dashboard/pages/dashboard-home-page.tsx` via `dashboard/data/api-adapter.ts:fetchFeedbackSignals` |
| 3 | `GetRecentActivities` | `/api/dashboard/recent-activities` | GET | `/psco.dashboard.v1.DashboardService/GetRecentActivities` | `dashboard/handler/query_handler.go:GetRecentActivities` | Connect handler + `service.QueryService` | `dashboard/pages/dashboard-home-page.tsx` via `dashboard/data/api-adapter.ts:fetchRecentActivities` |

### 2.6 Onboarding (`psco.onboarding.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `GetFirstRunState` | `/api/onboarding/state` | GET | `/psco.onboarding.v1.OnboardingService/GetFirstRunState` | `onboarding/handler/query_handler.go:GetFirstRunState` | Connect handler + `service.QueryService` | `onboarding/pages/onboarding-page.tsx` via `onboarding/data/api-adapter.ts` |

### 2.7 Export (`psco.export.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `GetExportSnapshot` | `/api/dashboard/export` | GET | `/psco.export.v1.ExportService/GetExportSnapshot` | `export/handler/query_handler.go:GetExportSnapshot` | Connect handler + `service.QueryService` | `dashboard/components/sovereignty-panel.tsx` via `dashboard/data/sovereignty-api-adapter.ts` |
| 2 | `ExportCoreAssets` | `/api/dashboard/export` | POST | `/psco.export.v1.ExportService/ExportCoreAssets` | `export/handler/command_handler.go:ExportCoreAssets` | Connect handler + `service.CommandService` | `dashboard/components/sovereignty-panel.tsx` via `dashboard/data/sovereignty-api-adapter.ts` |

### 2.8 Backup (`psco.backup.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `GetBackupSnapshot` | `/api/dashboard/backup` | GET | `/psco.backup.v1.BackupService/GetBackupSnapshot` | `backup/handler/query_handler.go:GetBackupSnapshot` | Connect handler + `service.QueryService` | `dashboard/components/sovereignty-panel.tsx` via `dashboard/data/sovereignty-api-adapter.ts` |
| 2 | `CreateInstanceBackup` | `/api/dashboard/backup` | POST | `/psco.backup.v1.BackupService/CreateInstanceBackup` | `backup/handler/command_handler.go:CreateInstanceBackup` | Connect handler + `service.CommandService` | `dashboard/components/sovereignty-panel.tsx` via `dashboard/data/sovereignty-api-adapter.ts` |

### 2.9 Reuse Summary (`psco.reuse_summary.v1`)

| # | RPC | 当前 HTTP 路径 | 方法 | 目标 Connect Procedure | 当前 Transport Owner | 迁移后 Owner | 页面/动作 Owner |
|---|-----|---------------|------|----------------------|---------------------|-------------|----------------|
| 1 | `GetReuseSummary` | `/api/reuse-summary` | GET | `/psco.reuse_summary.v1.ReuseSummaryService/GetReuseSummary` | `reusesummary/handler/query_handler.go:GetReuseSummary` | Connect handler + `service.QueryService` | `dashboard/pages/dashboard-home-page.tsx` + `module-registry/pages/module-detail-page.tsx` + `product-registry/pages/product-detail-page.tsx` via `reuse-summary/data/use-reuse-summary-read.ts` |

---

### 2.10 迁移总表统计

| 指标 | 数量 |
|------|------|
| Canonical 业务模块 | 9 |
| 总 RPC 数 | 34 |
| 读组 RPC | 21 |
| 写组 RPC | 13 |
| 含路径参数 URL | 15 |
| 无参数 URL | 19 |

---

## 3. 冻结：非业务基础设施端点 Keep List

以下端点**不纳入** phase07 迁移范围，继续保留在 `chi + net/http`：

| 端点 | 路径 | 方法 | 当前 Handler | 职责 |
|------|------|------|-------------|------|
| healthz | `/healthz` | GET | `platform/router.go:healthz` | 简单健康检查，返回 `{"status":"ok"}` |
| readyz | (待实现) | GET | (待实现) | 就绪检查（含 DB ping） |
| metrics | (待实现) | GET | (待实现) | Prometheus metrics 端点 |
| debug / pprof | (待实现) | GET | (待实现) | Go pprof 性能分析端点 |

**约束**：
- 这些端点不得被解释为 canonical 业务接口
- 这些端点不得纳入 `.proto` 合同
- 这些端点的职责只限基础设施，不承担业务合同定义
- 除上述端点外，不得再把其他业务主线接口列入长期 keep list

---

## 4. 冻结：Legacy / Compat 业务入口 Inventory

以下 4 个 legacy/compat 业务入口必须在 phase07 收口前退场：

| # | Legacy 入口 | 当前调用方 | 存在原因 | 替代 RPC / Connect Path | 允许并存最晚时点 | 删除证据 |
|---|------------|-----------|---------|------------------------|----------------|---------|
| 1 | `GET /api/candidates/products` | 当前前端无 active caller；仅余 `module-registry/data/api-adapter.ts` 导出 | phase02 Module Registry 曾临时承接 Product 候选读取；phase04 起 canonical owner 已迁移到 Product Registry | `ProductRegistryService.ListProducts` (`status_filter=ACTIVE`) 或等价候选 RPC；不得再复活 module-centered HTTP 入口 | phase07-09 后端切换完成 | `router.go` 中移除该路由；handler 中删除 `ListProductCandidates`；前端删除未使用 adapter 导出 |
| 2 | `GET /api/candidates/repositories` | 当前前端无 active caller；仅余 `module-registry/data/api-adapter.ts` 导出 | phase02 Module Registry 曾临时承接 Repository 候选读取；phase04 起 canonical owner 已迁移到 Repository Binding | `RepositoryBindingService.ListRepositories` (`status_filter=ACTIVE`) 或等价候选 RPC；不得再复活 module-centered HTTP 入口 | phase07-09 后端切换完成 | `router.go` 中移除该路由；handler 中删除 `ListRepositoryCandidates`；前端删除未使用 adapter 导出 |
| 3 | `POST /api/modules/{moduleId}/bindings/products` | 当前前端无 active caller；仅余 `module-registry/data/api-adapter.ts` 导出 | phase04 起为兼容委派，实际写动作委派到 Product Registry | `ProductRegistryService.BindModuleToProduct`（或等价 module-centered Connect facade）；正式写入 owner 仍归 Product Registry | phase07-10 adapter 清理完成 | `router.go` 中移除该路由；`command_handler.go` 中删除 `BindModuleToProduct` 兼容委派；前端删除未使用 adapter 导出 |
| 4 | `POST /api/modules/{moduleId}/bindings/repositories` | 当前前端无 active caller；仅余 `module-registry/data/api-adapter.ts` 导出 | phase04 起为兼容委派，实际写动作委派到 Repository Binding | `RepositoryBindingService.MapModuleToRepository`（或等价 module-centered Connect facade）；正式写入 owner 仍归 Repository Binding | phase07-10 adapter 清理完成 | `router.go` 中移除该路由；`command_handler.go` 中删除 `MapModuleToRepository` 兼容委派；前端删除未使用 adapter 导出 |

**约束**：
- 这 4 个入口不得作为 phase07 收口后的长期兼容层继续保留
- 每个入口的退场必须伴随：路由删除 + handler 代码删除 + 前端调用切换 + 回归验证
- 允许在迁移过程中短时并存（作为临时 adapter），但不得写入阶段完成态

---

## 5. 冻结：Phase07 收口判定标准

`phase07` 完成时，必须同时满足以下全部条件：

1. **Canonical 业务接口全部切换**：phase01 ~ phase06 的 34 条 proto-defined business RPC 均已切到 ConnectRPC，不再存在 hand-written JSON business handler 作为正式主线；其中 4 条 module-centered compat RPC 不得继续以 JSON HTTP 入口承担正式 owner 语义
2. **迁移总表覆盖**：每条 RPC 的 `service / RPC / 当前入口路径 / 目标 Connect procedure path / 页面动作 owner` 均已核销
3. **单一 `/api` 基址**：浏览器侧、Vite dev proxy、验收脚本与部署链路仍通过单一 `/api` 基址访问业务接口
4. **Legacy inventory 逐项核销**：4 个 legacy/compat 业务入口均已删除，不存在未声明残留
5. **不双主线并列**：phase 收口后不存在"新 Connect 主线 + 旧 JSON 主线"并列正式状态

**禁止**：
- 不得把"新 Connect 主线已增加，但旧 JSON 主线仍可继续工作"解释为完成
- 不得把"Connect 生成链已跑通，但业务接口未全部切换"解释为完成

---

## 6. 冻结：前端页面/动作 Owner 清单

### 6.1 前端 API Adapter 清单（迁移前当前状态）

| 模块 | Adapter 文件 | 涉及 RPC 数 | 当前 Transport |
|------|-------------|-----------|---------------|
| Module Registry | `module-registry/data/api-adapter.ts` | 8 | hand-written fetch + JSON |
| Decision Center | `decision-center/data/api-adapter.ts` | 5 | hand-written fetch + JSON |
| Product Registry | `product-registry/data/api-adapter.ts` | 5 | hand-written fetch + JSON |
| Repository Binding | `repository-binding/data/api-adapter.ts` | 7 | hand-written fetch + JSON |
| Dashboard | `dashboard/data/api-adapter.ts` | 3 | hand-written fetch + JSON |
| Dashboard (Sovereignty) | `dashboard/data/sovereignty-api-adapter.ts` | 4 (Export×2 + Backup×2) | hand-written fetch + JSON |
| Onboarding | `onboarding/data/api-adapter.ts` | 1 | hand-written fetch + JSON |
| Reuse Summary | `reuse-summary/data/api-adapter.ts` | 1 | hand-written fetch + JSON |

### 6.2 前端 Application Owner 清单（正式写动作承接位）

| 模块 | Application Owner 文件 | 承接写动作 | 迁移后状态 |
|------|------------------------|-----------|-----------|
| Module Registry | `module-registry/application/use-create-draft-module.ts` | CreateModule | 回收至 application owner |
| Decision Center | `decision-center/application/use-create-draft-decision.ts` | CreateDecision | 回收至 application owner |
| Product Registry | `product-registry/application/use-create-draft-product.ts` | CreateProduct | 回收至 application owner |
| Repository Binding | `repository-binding/application/use-create-draft-repository.ts` | CreateRepository | 回收至 application owner |

### 6.3 页面/组件内仍存在的 Mutation 清单（需回收或标记过渡）

| 位置 | Mutation | 当前状态 | 迁移后处理 |
|------|----------|---------|-----------|
| `module-registry/pages/release-create-page.tsx` | CreateRelease | 页面内 `useMutation` | 回收至 Module Registry 切片内固定承接位 |
| `product-registry/components/product-module-binding-panel.tsx` | BindModuleToProduct | 组件内 `useMutation` | 回收至 Product Registry 切片内固定承接位 |
| `repository-binding/components/repository-product-binding-panel.tsx` | BindRepositoryToProduct | 组件内 `useMutation` | 回收至 Repository Binding 切片内固定承接位 |
| `repository-binding/components/repository-module-mapping-panel.tsx` | MapModuleToRepository | 组件内 `useMutation` | 回收至 Repository Binding 切片内固定承接位 |
| `decision-center/components/decision-module-candidate-panel.tsx` | LinkDecisionToTarget | 组件内 `useMutation` | 回收至 Decision Center 切片内固定承接位 |
| `dashboard/components/sovereignty-panel.tsx` | ExportCoreAssets | 组件内 `useMutation` | 保留在 SovereigntyPanel 内（低频操作，允许短时过渡） |
| `dashboard/components/sovereignty-panel.tsx` | CreateInstanceBackup | 组件内 `useMutation` | 保留在 SovereigntyPanel 内（低频操作，允许短时过渡） |

**约束**：
- 前 5 项页面/组件内 mutation 必须在 phase07-10 中回收至各自切片内固定承接位
- Export/Backup 的 mutation 作为低频操作允许在 SovereigntyPanel 内短时过渡，但需在 phase07-10 中显式声明为"允许的过渡位"
- 不得在 phase07 收口后仍保留未声明的页面/组件级长期正式 mutation 主线

---

## 7. 与上游文档一致性声明

本文档已与以下上游文档完成逐项对齐：

- ✅ `phase07_transport_contract_mainline_migration_shared_baseline.md`：9 个业务模块 (§5.1)、keep list (§2.4)、验收基线 (§6) 一致
- ✅ `phase07_transport_contract_mainline_migration_architecture_plan.md`：迁移边界原则 (§4.6)、前后端交付策略 (§4.3/4.4)、完成条件 (§5) 一致
- ✅ `phase07_transport_contract_mainline_migration_dev_plan.md`：phase07-01 范围 (§L30-44) 全覆盖
- ✅ `audit_001_transport_contract_mainline_analysis.md`：escalate-phase 结论已落实为一次性正式切换
- ✅ `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/`：所有 9 个 proto 文件的 RPC 定义均已纳入迁移总表
- ✅ 根级规则 `project_rules.md` §2.6：`.proto` 唯一合同源 + ConnectRPC 正式传输层约束已遵守
