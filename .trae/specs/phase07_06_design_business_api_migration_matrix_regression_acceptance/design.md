# Phase07-06 设计产出：业务接口迁移矩阵与回归验收设计

> 本文档是 phase07-06 spec 的执行产出，产出 canonical 业务接口迁移矩阵（34 条 RPC）、跨模块回归清单、fixture/联调/退场证据矩阵、前端 mutation owner 验收映射、工具链迁移清单与 phase07 最终证据包结构。
> 产出日期：2026-08-11
> 上游：`phase07-01 frozen_scope.md`（34 RPC 总表 + legacy inventory）、`phase07-03 frozen_scope.md`（compat 退场标准）、`phase07-05 design.md`（前端 owner 清单）

---

## 1. 输入边界与当前源码事实

### 1.1 唯一上游基线

| 输入 | 来源 | 本设计复用方式 |
|------|------|--------------|
| 34 条 canonical RPC 总表 | `phase07-01 frozen_scope.md` §2 | 直接复用为迁移矩阵基础列，每条升级为"迁移波次 + 回归项 + 最终证据" |
| 4 条 legacy/compat endpoint inventory | `phase07-01 frozen_scope.md` §4 | 逐项升级为 endpoint 级退场证据矩阵 |
| 11 项前端 mutation owner | `phase07-05 design.md` §3 | 升级为验收映射表，绑 fixture + 触发位 + 回流检查项 |
| 9 个 business module 波次依赖 | `phase07-04 design.md` §1 | 明确波次划分与依赖关系 |
| 5 个 reset 脚本入口 | 当前仓库真实存在 | 纳入 fixture/联调/验收矩阵 |

### 1.2 当前工具链事实

| 工具链 | 当前状态 | 目标 phase07 状态 |
|--------|---------|------------------|
| `frontend/vite.config.ts` | 单一 `/api` proxy 到 `localhost:8081` | 不变，继续 `/api` 单一基址 |
| `frontend/package.json` | 无 `@connectrpc/connect` / `@connectrpc/connect-web` 依赖 | 新增两个依赖 |
| `proto/Makefile` | 5 个 target（build/gen/lint/breaking/clean） | `gen` target 需产出 Connect 生成物 |
| `proto/buf.gen.yaml` | 2 插件（`protocolbuffers/go` + `bufbuild/es`） | 新增 `connectrpc/go` 插件 |
| `database/scripts/reset_*.sh` | 5 个脚本，编排 `psql` / seed / 既有 reset 脚本，不直接发 HTTP 请求 | 继续作为 fixture / 数据恢复入口；HTTP / Connect 验证由脚本后的联调步骤承担 |
| CI workflow | **不存在** `.github/workflows/` 目录 | 显式建账为缺口 |

---

## 2. Canonical 业务接口迁移矩阵（34 条 RPC）

> 每条 RPC 下钻到 `service / RPC / 当前入口路径 / 方法 / 目标 Connect path / 当前 transport owner / 迁移后 owner / 页面/动作 owner / 迁移波次 / 最小回归项 / 最终收口证据`。

### 波次划分

| 波次 | 模块 | RPC 数 | 依赖 | 说明 |
|------|------|--------|------|------|
| **Wave 1** | Module Registry（前 4）+ Decision Center + Product Registry + Repository Binding | 23 | 无 | 核心 CRUD 模块，迁移后其他模块的 compat 入口才能退场 |
| **Wave 2** | Dashboard + Onboarding + Reuse Summary | 5 | Wave 1 | 读密集型模块，依赖 Wave 1 模块的 Connect 主线先就绪 |
| **Wave 3** | Export + Backup | 4 | Wave 1+2 | Sovereignty 低频操作，最后迁移 |
| **Wave 4** | Module Registry（后 4 compat）+ Legacy 退场 | 4 RPC + 4 endpoint | Wave 1+2+3 | compat facade 与 legacy 退场，最后核销 |

### 2.1 Wave 1：Module Registry（前 4 / 8 canconical）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 1 | `ListModules` | `/api/modules` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListModules` | `query_handler.go:ListModules` | Connect + `service.QueryService` | `modules/index.tsx` | Wave 1 | 列表页加载，筛选 queryText + statusFilter → 列表正确渲染 | `curl` Connect path 返回 200 + 列表页面功能正常 |
| 2 | `GetModuleDetail` | `/api/modules/{moduleId}` | GET | `/psco.module_registry.v1.ModuleRegistryService/GetModuleDetail` | `query_handler.go:GetModuleDetail` | Connect + `service.QueryService` | `modules/$moduleId.tsx` | Wave 1 | 详情页加载，含 Decision list + Reuse Summary inline | `curl` Connect path 返回 200 + 详情页功能正常 |
| 3 | `CreateModule` | `/api/modules` | POST | `/psco.module_registry.v1.ModuleRegistryService/CreateModule` | `command_handler.go:CreateModule` | Connect + `service.CommandService` | `application/use-create-draft-module.ts` | Wave 1 | 创建 module → 成功跳转 Detail → List 列表刷新含新记录 | `curl` Connect path 返回 200 + 创建流程正常 |
| 4 | `CreateRelease` | `/api/modules/{moduleId}/releases` | POST | `/psco.module_registry.v1.ModuleRegistryService/CreateRelease` | `command_handler.go:CreateRelease` | Connect + `service.CommandService` | `application/use-create-release.ts`（新增） | Wave 1 | 创建 release → 成功回流 → Detail 页 release 列表刷新 | `curl` Connect path 返回 200 + release 创建流程正常 |

### 2.2 Wave 1：Decision Center（5 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 5 | `ListDecisions` | `/api/decisions` | GET | `/psco.decision_center.v1.DecisionCenterService/ListDecisions` | `query_handler.go:ListDecisions` | Connect + `service.QueryService` | `decisions/index.tsx` | Wave 1 | 列表页加载，筛选 → 列表正确渲染 | `curl` Connect path 返回 200 + 列表页功能正常 |
| 6 | `GetDecisionDetail` | `/api/decisions/{decisionId}` | GET | `/psco.decision_center.v1.DecisionCenterService/GetDecisionDetail` | `query_handler.go:GetDecisionDetail` | Connect + `service.QueryService` | `decisions/$decisionId.tsx` | Wave 1 | 详情页加载 source_context + 已关联 target | `curl` Connect path 返回 200 + 详情页功能正常 |
| 7 | `ListDecisionModuleCandidates` | `/api/decisions/{decisionId}/candidates/modules` | GET | `/psco.decision_center.v1.DecisionCenterService/ListDecisionModuleCandidates` | `query_handler.go:ListDecisionModuleCandidates` | Connect + `service.QueryService` | `decisions/$decisionId.tsx` link dialog | Wave 1 | link dialog 拉取候选列表 | `curl` Connect path 返回 200 + 候选列表正常 |
| 8 | `CreateDecision` | `/api/decisions` | POST | `/psco.decision_center.v1.DecisionCenterService/CreateDecision` | `command_handler.go:CreateDecision` | Connect + `service.CommandService` | `application/use-create-draft-decision.ts` | Wave 1 | 创建 decision → 成功跳转 Detail → List 列表刷新 | `curl` Connect path 返回 200 + 创建流程正常 |
| 9 | `LinkDecisionToTarget` | `/api/decisions/{decisionId}/links` | POST | `/psco.decision_center.v1.DecisionCenterService/LinkDecisionToTarget` | `command_handler.go:LinkDecisionToTarget` | Connect + `service.CommandService` | `application/use-link-decision-to-target.ts`（新增） | Wave 1 | link module → 成功回流 → Detail 页 target 列表刷新 | `curl` Connect path 返回 200 + link 流程正常 |

### 2.3 Wave 1：Product Registry（5 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 10 | `ListProducts` | `/api/products` | GET | `/psco.product_registry.v1.ProductRegistryService/ListProducts` | `query_handler.go:ListProducts` | Connect + `service.QueryService` | `products/index.tsx` | Wave 1 | 列表页加载，筛选 → 列表正确渲染 | `curl` Connect path 返回 200 + 列表页正常 |
| 11 | `GetProductDetail` | `/api/products/{productId}` | GET | `/psco.product_registry.v1.ProductRegistryService/GetProductDetail` | `query_handler.go:GetProductDetail` | Connect + `service.QueryService` | `products/$productId.tsx` | Wave 1 | 详情页加载，含 module bindings + Reuse Summary inline | `curl` Connect path 返回 200 + 详情页正常 |
| 12 | `ListProductModuleCandidates` | `/api/products/{productId}/candidates/modules` | GET | `/psco.product_registry.v1.ProductRegistryService/ListProductModuleCandidates` | `query_handler.go:ListProductModuleCandidates` | Connect + `service.QueryService` | `products/$productId.tsx` bind dialog | Wave 1 | bind dialog 拉取候选 module 列表 | `curl` Connect path 返回 200 + 候选列表正常 |
| 13 | `CreateProduct` | `/api/products` | POST | `/psco.product_registry.v1.ProductRegistryService/CreateProduct` | `command_handler.go:CreateProduct` | Connect + `service.CommandService` | `application/use-create-draft-product.ts` | Wave 1 | 创建 product → 成功跳转 Detail → List 刷新 | `curl` Connect path 返回 200 + 创建流程正常 |
| 14 | `BindModuleToProduct` | `/api/products/{productId}/bindings/modules` | POST | `/psco.product_registry.v1.ProductRegistryService/BindModuleToProduct` | `command_handler.go:BindModuleToProduct` | Connect + `service.CommandService` | `application/use-bind-module-to-product.ts`（新增） | Wave 1 | bind module → 成功回流 → Detail 页 bindings 刷新 | `curl` Connect path 返回 200 + bind 流程正常 |

### 2.4 Wave 1：Repository Binding（7 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 15 | `ListRepositories` | `/api/repositories` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositories` | `query_handler.go:ListRepositories` | Connect + `service.QueryService` | `repositories/index.tsx` | Wave 1 | 列表页加载，筛选 → 列表正确渲染 | `curl` Connect path 返回 200 + 列表页正常 |
| 16 | `GetRepositoryDetail` | `/api/repositories/{repositoryId}` | GET | `/psco.repository_binding.v1.RepositoryBindingService/GetRepositoryDetail` | `query_handler.go:GetRepositoryDetail` | Connect + `service.QueryService` | `repositories/$repositoryId.tsx` | Wave 1 | 详情页加载，含 product bindings + module mappings | `curl` Connect path 返回 200 + 详情页正常 |
| 17 | `ListRepositoryProductCandidates` | `/api/repositories/{repositoryId}/candidates/products` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositoryProductCandidates` | `query_handler.go:ListRepositoryProductCandidates` | Connect + `service.QueryService` | `repositories/$repositoryId.tsx` bind dialog | Wave 1 | bind dialog 拉取候选 product 列表 | `curl` Connect path 返回 200 |
| 18 | `ListRepositoryModuleCandidates` | `/api/repositories/{repositoryId}/candidates/modules` | GET | `/psco.repository_binding.v1.RepositoryBindingService/ListRepositoryModuleCandidates` | `query_handler.go:ListRepositoryModuleCandidates` | Connect + `service.QueryService` | `repositories/$repositoryId.tsx` map dialog | Wave 1 | map dialog 拉取候选 module 列表 | `curl` Connect path 返回 200 |
| 19 | `CreateRepository` | `/api/repositories` | POST | `/psco.repository_binding.v1.RepositoryBindingService/CreateRepository` | `command_handler.go:CreateRepository` | Connect + `service.CommandService` | `application/use-create-draft-repository.ts` | Wave 1 | 创建 repository → 成功跳转 Detail → List 刷新 | `curl` Connect path 返回 200 + 创建流程正常 |
| 20 | `BindRepositoryToProduct` | `/api/repositories/{repositoryId}/bindings/products` | POST | `/psco.repository_binding.v1.RepositoryBindingService/BindRepositoryToProduct` | `command_handler.go:BindRepositoryToProduct` | Connect + `service.CommandService` | `application/use-bind-repository-to-product.ts`（新增） | Wave 1 | bind product → 成功回流 → Detail 页 bindings 刷新 | `curl` Connect path 返回 200 + bind 流程正常 |
| 21 | `MapModuleToRepository` | `/api/repositories/{repositoryId}/bindings/modules` | POST | `/psco.repository_binding.v1.RepositoryBindingService/MapModuleToRepository` | `command_handler.go:MapModuleToRepository` | Connect + `service.CommandService` | `application/use-map-module-to-repository.ts`（新增） | Wave 1 | map module → 成功回流 → Detail 页 mappings 刷新 | `curl` Connect path 返回 200 + map 流程正常 |

### 2.5 Wave 2：Dashboard（3 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 22 | `GetDashboardOverview` | `/api/dashboard/overview` | GET | `/psco.dashboard.v1.DashboardService/GetDashboardOverview` | `query_handler.go:GetOverview` | Connect + `service.QueryService` | `data/use-dashboard-overview-read.ts`（新增） | Wave 2 | 首页加载，聚合计数正确 | `curl` Connect path 返回 200 + 首页计数正常 |
| 23 | `GetFeedbackSignals` | `/api/dashboard/feedback-signals` | GET | `/psco.dashboard.v1.DashboardService/GetFeedbackSignals` | `query_handler.go:GetFeedbackSignals` | Connect + `service.QueryService` | `data/use-feedback-signals-read.ts`（新增） | Wave 2 | 首页加载，Feedback 区块正确渲染 | `curl` Connect path 返回 200 + Feedback 正常 |
| 24 | `GetRecentActivities` | `/api/dashboard/recent-activities` | GET | `/psco.dashboard.v1.DashboardService/GetRecentActivities` | `query_handler.go:GetRecentActivities` | Connect + `service.QueryService` | `data/use-recent-activities-read.ts`（新增） | Wave 2 | 首页加载，Recent Activities 正确渲染 | `curl` Connect path 返回 200 + Activities 正常 |

### 2.6 Wave 2：Onboarding（1 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 25 | `GetFirstRunState` | `/api/onboarding/state` | GET | `/psco.onboarding.v1.OnboardingService/GetFirstRunState` | `query_handler.go:GetFirstRunState` | Connect + `service.QueryService` | `data/use-onboarding-read.ts`（替换 transport） | Wave 2 | first_run_state 三态正确 → CTA 分流正确 | `curl` Connect path 返回 200 + Onboarding 页面正常 |

### 2.7 Wave 2：Reuse Summary（1 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 26 | `GetReuseSummary` | `/api/reuse-summary` | GET | `/psco.reuse_summary.v1.ReuseSummaryService/GetReuseSummary` | `query_handler.go:GetReuseSummary` | Connect + `service.QueryService` | `data/use-reuse-summary-read.ts`（替换 transport） | Wave 2 | Dashboard / Module Detail / Product Detail 三处 ReuseSummary 正确渲染 | `curl` Connect path 返回 200 + 三处页面正常 |

### 2.8 Wave 3：Export（2 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 27 | `GetExportSnapshot` | `/api/dashboard/export` | GET | `/psco.export.v1.ExportService/GetExportSnapshot` | `query_handler.go:GetExportSnapshot` | Connect + `service.QueryService` | `dashboard/components/sovereignty-panel.tsx` | Wave 3 | 快照信息正确显示 | `curl` Connect path 返回 200 + 快照信息正常 |
| 28 | `ExportCoreAssets` | `/api/dashboard/export` | POST | `/psco.export.v1.ExportService/ExportCoreAssets` | `command_handler.go:ExportCoreAssets` | Connect + `service.CommandService` | `dashboard/components/sovereignty-panel.tsx`（过渡位） | Wave 3 | 点击导出 → 成功返回 | 导出流程正常 |

### 2.9 Wave 3：Backup（2 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 最小回归项 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|-----------|---------|
| 29 | `GetBackupSnapshot` | `/api/dashboard/backup` | GET | `/psco.backup.v1.BackupService/GetBackupSnapshot` | `query_handler.go:GetBackupSnapshot` | Connect + `service.QueryService` | `dashboard/components/sovereignty-panel.tsx` | Wave 3 | 快照信息正确显示 | `curl` Connect path 返回 200 + 快照正常 |
| 30 | `CreateInstanceBackup` | `/api/dashboard/backup` | POST | `/psco.backup.v1.BackupService/CreateInstanceBackup` | `command_handler.go:CreateInstanceBackup` | Connect + `service.CommandService` | `dashboard/components/sovereignty-panel.tsx`（过渡位） | Wave 3 | 点击备份 → 成功返回 | 备份流程正常 |

### 2.10 Wave 4：Module Registry 后 4 条 Compat RPC（4 RPC）

| # | RPC | 当前路径 | 方法 | 目标 Connect Procedure | 当前 Transport | 迁移后 Owner | 页面/动作 Owner | 波次 | 特殊说明 | 最终证据 |
|---|-----|---------|------|----------------------|---------------|-------------|----------------|------|---------|---------|
| 31 | `BindModuleToProduct` | `/api/modules/{moduleId}/bindings/products` | POST | `/psco.module_registry.v1.ModuleRegistryService/BindModuleToProduct` | `command_handler.go:BindModuleToProduct`（compat 委派） | Connect handler（compat facade）+ Product Registry canonical | 无 active caller（仅作为 compat 语义存在于 .proto） | Wave 4 | **transport inventory 不等于 canonical business owner**；Connect handler 实现为 compat 薄壳，正式写入由 Product Registry 承接 | Wave 1+2+3 完成后该 compat facade 只作为传输层提供协议兼容，不作为业务 owner |
| 32 | `MapModuleToRepository` | `/api/modules/{moduleId}/bindings/repositories` | POST | `/psco.module_registry.v1.ModuleRegistryService/MapModuleToRepository` | `command_handler.go:MapModuleToRepository`（compat 委派） | Connect handler（compat facade）+ Repository Binding canonical | 无 active caller | Wave 4 | 同上 | 同上 |
| 33 | `ListProductCandidates` | `/api/candidates/products` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListProductCandidates` | `query_handler.go:ListProductCandidates`（legacy compat） | Connect handler（compat facade）+ Product Registry canonical | 无 active caller | Wave 4 | 同上 | 同上 |
| 34 | `ListRepositoryCandidates` | `/api/candidates/repositories` | GET | `/psco.module_registry.v1.ModuleRegistryService/ListRepositoryCandidates` | `query_handler.go:ListRepositoryCandidates`（legacy compat） | Connect handler（compat facade）+ Repository Binding canonical | 无 active caller | Wave 4 | 同上 | 同上 |

---

## 3. 跨模块回归清单

### 3.1 模块内回归

| 模块 | 回归项 | 覆盖 RPC # | 验证方式 |
|------|--------|-----------|---------|
| **Module Registry** | 列表页加载 + 筛选 | 1 | 页面渲染检查 |
| | 详情页加载（含 Decision list + Reuse Summary） | 2 | 页面渲染检查 |
| | 创建 Module → 跳转 Detail → List 刷新 | 3 | 端到端操作 |
| | 创建 Release → Detail 页 release 列表刷新 | 4 | 端到端操作 |
| **Decision Center** | 列表页加载 + 筛选 | 5 | 页面渲染检查 |
| | 详情页加载（含 source_context + 已关联 target） | 6 | 页面渲染检查 |
| | 创建 Decision → 跳转 Detail → List 刷新 | 8 | 端到端操作 |
| | Link Decision → Detail 页 target 列表刷新 | 7, 9 | 端到端操作 |
| **Product Registry** | 列表页加载 + 筛选 | 10 | 页面渲染检查 |
| | 详情页加载（含 bindings + Reuse Summary） | 11 | 页面渲染检查 |
| | 创建 Product → 跳转 Detail → List 刷新 | 13 | 端到端操作 |
| | Bind Module → Detail 页 bindings 刷新 | 12, 14 | 端到端操作 |
| **Repository Binding** | 列表页加载 + 筛选 | 15 | 页面渲染检查 |
| | 详情页加载（含 bindings + mappings） | 16 | 页面渲染检查 |
| | 创建 Repository → 跳转 Detail → List 刷新 | 19 | 端到端操作 |
| | Bind Product → Detail 页 bindings 刷新 | 17, 20 | 端到端操作 |
| | Map Module → Detail 页 mappings 刷新 | 18, 21 | 端到端操作 |
| **Dashboard** | 首页加载（Overview + Feedback + Recent Activities） | 22, 23, 24 | 页面渲染检查 |
| | SovereigntyPanel（Export + Backup 快照与触发） | 27, 28, 29, 30 | 页面渲染 + 操作 |
| | Reuse Summary 挂接 | 26 | 页面渲染检查 |
| **Onboarding** | first_run_state 三态 → CTA 分流 | 25 | 页面渲染 + route guard |
| | 冷启动 → 完成 Onboarding → 回到 Dashboard | 25 | 端到端流程 |

### 3.2 跨模块联动回归

| # | 联动路径 | 覆盖模块 | 验证方式 |
|---|---------|---------|---------|
| CR1 | 创建 Module → 在 Product Detail 中 bind 该 module → Module Detail 中 product_bind_count 更新 | Module Registry + Product Registry | 端到端操作 + 刷新检查 |
| CR2 | 创建 Repository → 在 Product Detail 中 bind 该 repository → Product Detail 中 product_with_repository 更新 | Repository Binding + Product Registry | 端到端操作 + 刷新检查 |
| CR3 | 创建 Module → 在 Repository Detail 中 map 该 module → Module Detail 中 repository_bind_count 更新 | Module Registry + Repository Binding | 端到端操作 + 刷新检查 |
| CR4 | 创建 Decision 并 link module → Decision Detail 中 target 列表更新 → Module Detail 中 decision 列表更新 | Decision Center + Module Registry | 端到端操作 + 刷新检查 |
| CR5 | Dashboard 首页 → Overview 计数与实际情况一致（module_count / product_count / decision_count 等） | Dashboard + 全部 CRUD 模块 | 复位 → 创建 → 回到 Dashboard 检查计数 |
| CR6 | Onboarding first_run → 完成 → Dashboard 首页不再显示 Onboarding CTA | Onboarding + Dashboard | 端到端流程 |
| CR7 | Dashboard 首页 → 点击 CTA 进入各模块列表 → 列表页加载 → 进入详情 → 返回 | Dashboard + 全部 List/Detail 页面 | 路径遍历 |
| CR8 | Module Detail → Reuse Summary inline 显示 → Product Detail → Reuse Summary inline 显示 | Module Registry + Product Registry + Reuse Summary | 页面渲染检查 |
| CR9 | mutation 成功后 → 对应 read owner 的 queryClient.invalidateQueries 触发 → 页面数据刷新 | 所有 mutation ↔ read owner 对 | 操作 + 刷新检查 |

### 3.3 Route 级回归

| # | 路径 | 验证方式 |
|---|------|---------|
| R1 | `/` → 根据 `first_run_state` 重定向到 `/onboarding` 或 `/dashboard` | `beforeLoad` 分流验证 |
| R2 | `/onboarding` → first_run_state 路由守卫正确分流 | route guard 验证 |
| R3 | 各模块 `/modules`、`/products`、`/repositories`、`/decisions` → 列表页正常 | 页面加载 |
| R4 | 各模块 `/$id` → 详情页正常 | 页面加载 |
| R5 | 各模块 `/new` → 创建页正常 | 页面加载 |
| R6 | 创建成功 → 跳转 Detail 页 → fromList/fromModuleDetail/fromDashboard 等来源参数正确传递 | 来源链路验证 |

---

## 4. Fixture、联调、退场证据与最终收口证据矩阵

### 4.1 Fixture / 环境入口到回归项的映射

| 脚本入口 | 对应业务模块 | 覆盖 RPC # | 默认恢复 | 联调步骤 | 期望结果 |
|---------|-------------|-----------|---------|---------|---------|
| `reset_module_mainline.sh` | Module Registry | 1-4 | 恢复基线 module 数据 | 1) 执行脚本 2) 访问 Module List 3) 访问 Module Detail 4) 创建 Module 5) 创建 Release | 列表/详情/创建/Release 均正常；#31-34 改由 §4.2 legacy/compat 退场矩阵单独核销 |
| `reset_decision_mainline.sh` | Decision Center | 5-9 | 恢复基线 decision 数据 | 1) 执行脚本 2) 访问 Decision List 3) 访问 Decision Detail 4) 创建 Decision 5) Link Decision | 列表/详情/创建/Link 均正常 |
| `reset_product_repository_mainline.sh` | Product Registry + Repository Binding | 10-21 | 恢复基线 product/repository 数据 | 1) 执行脚本 2) 访问 Product List/Detail 3) 访问 Repository List/Detail 4) 创建 Product/Repository 5) Bind/Map | 所有操作正常 |
| `reset_dashboard_acceptance.sh` | Dashboard | 22-24 | 恢复 dashboard 验收数据 | 1) 执行脚本 2) 访问 Dashboard 首页 3) 检查 Overview/Feedback/Activities | 三个区块均正常渲染 |
| `reset_phase06_acceptance.sh` | Onboarding + Export + Backup + Reuse | 25-30 | 恢复 phase06 验收数据 | 1) 执行脚本 2) 访问 Onboarding 3) 检查 Export/Backup 4) 检查 Reuse Summary | 全部正常 |

### 4.2 Legacy / Compat Endpoint 退场证据矩阵

| # | Legacy 入口 | 替代 Connect Path | 退场时点 | 路由删除证据 | Handler/Adapter 删除证据 | 替代 Connect 回归证据 |
|---|------------|------------------|---------|-------------|------------------------|---------------------|
| L1 | `GET /api/candidates/products` | `ProductRegistryService.ListProducts` | phase07-09 | `router.go` 中无 `r.Get("/api/candidates/products", ...)` | `query_handler.go` 中无 `ListProductCandidates` 方法 + `ProductCandidateReader` 字段 | `curl` Connect path 返回 200 + 候选列表正确 |
| L2 | `GET /api/candidates/repositories` | `RepositoryBindingService.ListRepositories` | phase07-09 | `router.go` 中无 `r.Get("/api/candidates/repositories", ...)` | `query_handler.go` 中无 `ListRepositoryCandidates` 方法 + `RepositoryCandidateReader` 字段 | `curl` Connect path 返回 200 + 候选列表正确 |
| L3 | `POST /api/modules/{moduleId}/bindings/products` | `ProductRegistryService.BindModuleToProduct` | phase07-10 | `router.go` 中无 `r.Post("/api/modules/{moduleId}/bindings/products", ...)` | `command_handler.go` 中无 `BindModuleToProduct` 方法 + 前端 `api-adapter.ts` 无 `bindModuleToProduct` 导出 | `curl` 旧路径返回 404 + Connect 替代路径正常 |
| L4 | `POST /api/modules/{moduleId}/bindings/repositories` | `RepositoryBindingService.MapModuleToRepository` | phase07-10 | `router.go` 中无 `r.Post("/api/modules/{moduleId}/bindings/repositories", ...)` | `command_handler.go` 中无 `MapModuleToRepository` 方法 + 前端 `api-adapter.ts` 无 `mapModuleToRepository` 导出 | `curl` 旧路径返回 404 + Connect 替代路径正常 |

### 4.3 Phase07-11 最终证据包结构

```
phase07 收口证据包
├── canonical_rpc_migration.csv            # 34 条 RPC 的核销结果（逐条 ✅/❌）
├── legacy_endpoint_retirement.csv         # 4 条 legacy endpoint 的退场核销
├── frontend_mutation_owner_acceptance.csv # 11 项 mutation owner 的验收结果
├── frontend_adapter_retirement.csv        # 13 项 adapter 的删除核销
├── cross_module_regression.csv            # 9 项跨模块联动回归结果
├── toolchain_migration.csv                # 工具链迁移核销（见 §5）
├── build_verification.txt                 # `(cd backend && go build ./...)` + `(cd frontend && tsc -b --noEmit)` + `(cd frontend && npm run build)` 通过
└── phase07_closure_decision.md            # 收口判定：双门槛全部满足 → 收口
```

### 4.4 阻断条件

以下任一条件不满足 → `phase07` 不得收口：

| 阻断条件 | 检查方式 |
|---------|---------|
| 34 条 canonical RPC 未全部迁移至 Connect | 逐条核销 canonical_rpc_migration.csv |
| 4 条 legacy endpoint 未全部退场 | 逐条核销 legacy_endpoint_retirement.csv |
| 前端 11 项 mutation owner 未全部收口 | 逐条核销 frontend_mutation_owner_acceptance.csv |
| 旧 JSON business handler 主线仍存在 | `grep` `router.go` 中 `r.Get/Post` 业务路径 |
| 前端 adapter 文件未全部删除 | 13 项 adapter 删除核销 |
| 跨模块回归 CR1-CR9 未全部通过 | 逐条核销 cross_module_regression.csv |
| `(cd backend && go build ./...)` 或 `(cd frontend && tsc -b --noEmit)` 失败 | 编译检查 |
| `/api` 基址出现第二套并列前缀 | 代码审查 |

---

## 5. 前端 Mutation Owner 验收映射表

### 5.1 验收映射（11 项）

| # | 写动作 | 当前 Owner | 目标 Owner | 类型 | 触发页面/组件 | 前置 Fixture | 成功回流检查 | 最晚核销 |
|---|--------|-----------|-----------|------|-------------|-------------|-------------|---------|
| 1 | `CreateModule` | `application/use-create-draft-module.ts` | 同文件（替换 transport） | ✅ Application Owner | `modules/new.tsx` | `reset_module_mainline.sh` | List 刷新 + Detail 跳转 | phase07-07 |
| 2 | `CreateDecision` | `application/use-create-draft-decision.ts` | 同文件（替换 transport） | ✅ Application Owner | `decisions/new.tsx` | `reset_decision_mainline.sh` | List 刷新 + Detail 跳转 | phase07-07 |
| 3 | `CreateProduct` | `application/use-create-draft-product.ts` | 同文件（替换 transport） | ✅ Application Owner | `products/new.tsx` | `reset_product_repository_mainline.sh` | List 刷新 + Detail 跳转 | phase07-07 |
| 4 | `CreateRepository` | `application/use-create-draft-repository.ts` | 同文件（替换 transport） | ✅ Application Owner | `repositories/new.tsx` | `reset_product_repository_mainline.sh` | List 刷新 + Detail 跳转 | phase07-07 |
| 5 | `CreateRelease` | `pages/release-create-page.tsx`（页面内） | `application/use-create-release.ts`（新增） | 🔄 回收到 Owner | `release-create-page.tsx` | `reset_module_mainline.sh` | Detail 页 release 列表刷新 | phase07-07 |
| 6 | `BindModuleToProduct` | `components/product-module-binding-panel.tsx`（组件内） | `application/use-bind-module-to-product.ts`（新增） | 🔄 回收到 Owner | `product-module-binding-panel.tsx` | `reset_product_repository_mainline.sh` | Product Detail bindings 刷新 | phase07-07 |
| 7 | `BindRepositoryToProduct` | `components/repository-product-binding-panel.tsx`（组件内） | `application/use-bind-repository-to-product.ts`（新增） | 🔄 回收到 Owner | `repository-product-binding-panel.tsx` | `reset_product_repository_mainline.sh` | Repository Detail bindings 刷新 | phase07-07 |
| 8 | `MapModuleToRepository` | `components/repository-module-mapping-panel.tsx`（组件内） | `application/use-map-module-to-repository.ts`（新增） | 🔄 回收到 Owner | `repository-module-mapping-panel.tsx` | `reset_product_repository_mainline.sh` | Repository Detail mappings 刷新 | phase07-07 |
| 9 | `LinkDecisionToTarget` | `components/decision-module-candidate-panel.tsx`（组件内） | `application/use-link-decision-to-target.ts`（新增） | 🔄 回收到 Owner | `decision-module-candidate-panel.tsx` | `reset_decision_mainline.sh` | Decision Detail target 列表刷新 | phase07-07 |
| 10 | `ExportCoreAssets` | `components/sovereignty-panel.tsx`（组件内） | 同组件内（替换 transport） | ⏸️ 短时过渡位 | `sovereignty-panel.tsx` | `reset_phase06_acceptance.sh` | 导出成功 | phase07-10 |
| 11 | `CreateInstanceBackup` | `components/sovereignty-panel.tsx`（组件内） | 同组件内（替换 transport） | ⏸️ 短时过渡位 | `sovereignty-panel.tsx` | `reset_phase06_acceptance.sh` | 备份成功 | phase07-10 |

### 5.2 Candidate Read + Mutation 联合验收

| 组件 | Candidate Read | Mutation | 验收要求 |
|------|---------------|---------|---------|
| `product-module-binding-panel.tsx` | `ListProductModuleCandidates`（RPC #12） | `BindModuleToProduct`（RPC #14） | 候选列表通过 Connect 加载 → bind 成功 → 候选列表刷新 |
| `repository-product-binding-panel.tsx` | `ListRepositoryProductCandidates`（RPC #17） | `BindRepositoryToProduct`（RPC #20） | 候选列表通过 Connect 加载 → bind 成功 → 候选列表刷新 |
| `repository-module-mapping-panel.tsx` | `ListRepositoryModuleCandidates`（RPC #18） | `MapModuleToRepository`（RPC #21） | 候选列表通过 Connect 加载 → map 成功 → 候选列表刷新 |
| `decision-module-candidate-panel.tsx` | `ListDecisionModuleCandidates`（RPC #7） | `LinkDecisionToTarget`（RPC #9） | 候选列表通过 Connect 加载 → link 成功 → 候选列表刷新 |

---

## 6. 工具链迁移清单

### 6.1 Vite / Frontend

| 项目 | 当前事实 | 目标状态 | 修改者 | 验证命令 | 阻断条件 |
|------|---------|---------|--------|---------|---------|
| `vite.config.ts` | 单一 `/api` proxy → `localhost:8081` | 不变 | 无须修改 | `npm run dev` → 前端可访问 | `/api` 前缀被改为非单一基址 |
| `package.json` dependencies | 无 `@connectrpc/connect` / `@connectrpc/connect-web` | 新增两个依赖 | phase07-07 | `npm install` → `npm run build` 通过 | `npm install` 失败 |
| `tsc -b --noEmit` | 当前通过 | 迁移后通过 | phase07-07 | `tsc -b --noEmit` | TypeScript 编译失败 |

### 6.2 Proto 生成链

| 项目 | 当前事实 | 目标状态 | 修改者 | 验证命令 | 阻断条件 |
|------|---------|---------|--------|---------|---------|
| `proto/buf.gen.yaml` | 2 插件：`protocolbuffers/go` + `bufbuild/es` | 新增 `connectrpc/go` 插件 | phase07-08 | `make gen` → Go Connect 产物生成 | `buf generate` 失败 |
| `proto/Makefile` `gen` target | `buf generate` | 不变（`buf.gen.yaml` 已含新插件） | 无须修改 Makefile | `make gen` | 生成产物不完整 |
| `proto/Makefile` `build` target | `buf build` | 不变 | 无须修改 | `make build` | `buf build` 失败 |
| `proto/Makefile` `lint` target | `buf lint` | 不变 | 无须修改 | `make lint` | Lint 失败 |
| `proto/Makefile` `breaking` target | `buf breaking --against '.git#branch=main'` | 不变 | 无须修改 | `make breaking` | 破坏性变更检测失败 |

### 6.3 本地脚本

| 项目 | 当前事实 | 目标状态 | 修改者 | 验证命令 | 阻断条件 |
|------|---------|---------|--------|---------|---------|
| `reset_module_mainline.sh` | 仅做 module 基线清空 / 恢复，不发 HTTP 请求 | 保持脚本逻辑不变，继续作为 DB reset 入口；Connect 验证由脚本后的联调步骤承担 | phase07-11 | 执行脚本 → DB 基线恢复成功 | 脚本执行失败 |
| `reset_decision_mainline.sh` | 仅做 decision 基线清空 / 恢复，不发 HTTP 请求 | 保持脚本逻辑不变，继续作为 DB reset 入口；Connect 验证由脚本后的联调步骤承担 | phase07-11 | 执行脚本 → DB 基线恢复成功 | 脚本执行失败 |
| `reset_product_repository_mainline.sh` | 仅做 product / repository 基线清空 / 恢复，不发 HTTP 请求 | 保持脚本逻辑不变，继续作为 DB reset 入口；Connect 验证由脚本后的联调步骤承担 | phase07-11 | 执行脚本 → DB 基线恢复成功 | 脚本执行失败 |
| `reset_dashboard_acceptance.sh` | 编排 dashboard 验收基线恢复，不发 HTTP 请求 | 保持脚本逻辑不变，继续作为 fixture 入口；Connect 验证由脚本后的联调步骤承担 | phase07-11 | 执行脚本 → Dashboard 验收基线恢复成功 | 脚本执行失败 |
| `reset_phase06_acceptance.sh` | 编排 phase06 验收基线恢复，不发 HTTP 请求 | 保持脚本逻辑不变，继续作为 fixture 入口；Connect 验证由脚本后的联调步骤承担 | phase07-11 | 执行脚本 → phase06 验收基线恢复成功 | 脚本执行失败 |

### 6.4 CI 缺口

| 项目 | 当前事实 | 目标状态 | 缺口处理 |
|------|---------|---------|---------|
| CI workflow | `.github/workflows/` **不存在** | 显式建账为缺口 | 在 phase07 收口证据中记录"当前仓库无正式 CI 入口"，phase07 不因 CI 缺失而阻断，但必须显式声明，并在后续阶段补齐。替代方案：本地 `(cd proto && make build && make gen && make lint)` + `(cd backend && go build ./...)` + `(cd frontend && tsc -b --noEmit)` 作为等价证据 |

---

## 7. 迁移执行顺序总览

```
phase07-07（前端实现）          phase07-08（后端实现）          phase07-09（后端 compat 退场）
┌───────────────────────┐     ┌───────────────────────┐     ┌──────────────────────────┐
│ Wave 1 前端:           │     │ Wave 1-3 后端:         │     │ L1: GET /api/candidates/  │
│ - 5 个新增 application  │     │ - 9 模块 Connect 实现   │     │     products → 退场       │
│   owner                │     │ - platform 装配         │     │ L2: GET /api/candidates/  │
│ - 4 个既有 application  │     │ - connect_errors.go    │     │     repositories → 退场    │
│   owner transport 替换  │     │ - router.go 调整        │     │ router.go 清理            │
│ - 12 个 read owner     │     │                         │     │ query_handler.go 清理     │
│ - 各切片 connect-client │     │                         │     │                            │
└───────────────────────┘     └───────────────────────┘     └──────────────────────────┘
           │                            │                            │
           └────────────┬───────────────┘                            │
                        │                                            │
                   phase07-10（前端 compat 退场）                      │
                   ┌──────────────────────────┐                      │
                   │ L3: POST /api/modules/    │                      │
                   │     {id}/bindings/products│                      │
                   │ L4: POST /api/modules/    │                      │
                   │     {id}/bindings/repos   │                      │
                   │ 13 项 adapter 删除         │                      │
                   │ compat switch 最终收口     │                      │
                   └──────────────────────────┘                      │
                              │                                      │
                         phase07-11（验收与收口）                        │
                         ┌──────────────────────────┐                │
                         │ 34 RPC 核销               │                │
                         │ 4 legacy endpoint 核销    │                │
                         │ 11 mutation owner 核销    │                │
                         │ 9 跨模块回归核销           │                │
                         │ 工具链迁移核销             │                │
                         │ 收口判定 → 双门槛全部满足   │                │
                         └──────────────────────────┘                │
```

---

## 8. 与上游文档一致性声明

| 上游文档 | 关键结论 | 本设计对齐 |
|---------|---------|-----------|
| `phase07-01 frozen_scope.md` §2 | 34 条 RPC 总表 + legacy inventory | §2 迁移矩阵逐条复用 34 条 RPC，升级为可执行条目 |
| `phase07-01 frozen_scope.md` §4 | 4 条 legacy endpoint inventory | §4.2 升级为 endpoint 级退场证据矩阵 |
| `phase07-01 frozen_scope.md` §5 | Phase07 收口判定标准 | §4.3 证据包 + §4.4 阻断条件 |
| `phase07-03 frozen_scope.md` §2-5 | compat 退场窗口 + 双门槛 | §4.2 退场证据 + §4.4 阻断条件 |
| `phase07-04 design.md` §1-3 | Connect handler 接线 + router 结构 | §2 波次划分 + 迁移矩阵 |
| `phase07-05 design.md` §3-4 | 11 mutation owner + 13 adapter 回收 | §5 验收映射 + §6.3 工具链 |
| `phase06-16 acceptance_report.md` | Dashboard / Onboarding 验收边界 | §3.2 跨模块回归 CR5-CR8 复用既有验收边界 |
| `project_rules.md` §2.6 | `.proto` 唯一合同源 | 全文迁移矩阵以 Connect procedure path 为唯一目标路径 |
