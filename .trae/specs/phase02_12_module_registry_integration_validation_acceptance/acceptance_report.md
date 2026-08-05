# phase02-12 Module Registry 联调、验证与验收报告

> 本文件为 `phase02-12` 的可复核验收证据，依据 `spec.md` 与 `tasks.md` 验收路径执行。
> 验收环境：本地共享 PostgreSQL 容器 `rento-preview-postgres` + Go 后端（端口 8081）+ React 前端（端口 5173，真实 API 模式 `VITE_USE_REAL_API=true`）。
> 验收时间：2026-08-05。

## 1. 联调环境前提（Task 1）

| 项 | 验证结果 |
| --- | --- |
| 数据库 | 复用 `rento-preview-postgres` 容器，`psco_development` 库已通过 `init_db.sh` 建立并通过 `migrate.go` 完成迁移 |
| 只读前提种子 | `run_seeds.sh` 执行 `seed_readonly_prereqs.sql`，提供产品（Product A/B/C）、仓库（main-repo/mirror-repo）、决策（auth-service 技术选型）候选，幂等可重复 |
| Decision 关联前提 | `run_seeds.sh` 默认 **不** 执行 `seed_decision_link_fixture.sql`（需 `RUN_DECISION_LINK_FIXTURE=1` 显式启用）。phase02-12 验收基线通过 `reset_module_mainline.sh` 重建 `decision_links`，不依赖 fixture |
| 模块主线基线 | `reset_module_mainline.sh`（默认模式）提供"清空 + 恢复基线"可重复入口，重建 2 个模块 + 3 个版本 + 2 个产品绑定 + 2 个仓库映射 + 1 个决策关联，幂等可重复 |
| 后端 | `./backend/bin/psco-server` 启动，监听 `:8081`，`RUN_SEEDS_ON_BOOT=true` |
| 前端 | `npm run dev` 启动，监听 `:5173`，`.env` 设置 `VITE_USE_REAL_API=true` 与 `VITE_API_BASE_URL=http://localhost:8081` |
| 启动顺序 | 数据库 → 后端（迁移 + 种子）→ 前端（真实 API 模式） |
| 复验入口 | `./database/scripts/init_db.sh` → `./database/scripts/run_seeds.sh` → `./database/scripts/reset_module_mainline.sh` |
| 合同基准 | `proto/psco/module_registry/v1/module_registry.proto` 为唯一合同源 |

## 2. 最小主线正向路径（Task 2）

### 2.1 CreateModule 最小闭环

- 操作：`POST /api/modules`，请求体 `{ "name": "contract-verify-module", "description": "...", "status": "active" }`
- 结果：HTTP 201，响应体返回完整 `Module` 对象（含 `id` 与 `created_at`）
- 前端：列表页可看到新建模块，详情页显示核心字段
- **结论：通过**

### 2.2 CreateRelease 最小闭环

- 操作：`POST /api/modules/{moduleId}/releases`，请求体 `{ "version": "0.1.0", "status": "active", "released_at": "2026-08-05T15:00:00Z" }`
- 结果：HTTP 201，响应体返回完整 `Release` 对象（含 `id` / `module_id` / `released_at`）
- 前端：详情页版本列表承接最新版本
- **结论：通过**

### 2.3 绑定动作最小闭环

- BindModuleToProduct：`POST /api/modules/{moduleId}/bindings/products`，请求体 `{ "product_id": "..." }`，返回 HTTP 204
- MapModuleToRepository：`POST /api/modules/{moduleId}/bindings/repositories`，请求体 `{ "repository_id": "..." }`，返回 HTTP 204
- 前端：绑定后停留详情页，重新读取并显示最新绑定结果
- **结论：通过**

## 3. 关键状态路径（Task 3）

### 3.1 空状态验证

- 操作：执行 `./database/scripts/reset_module_mainline.sh --clean-only` 清空模块主线（依赖 `ON DELETE CASCADE` 级联清空 4 张关联表），访问 `/modules?statusFilter=all`
- 结果：列表页显示"系统中尚无任何模块，先完成首个模块登记"与"完成首个模块登记"按钮（直接进入 `/modules/new`）
- 恢复：执行 `./database/scripts/reset_module_mainline.sh`（默认模式）完整恢复 2 个模块 + 3 个版本 + 2 个产品绑定 + 2 个仓库映射 + 1 个决策关联，幂等可重复
- **结论：通过**

### 3.2 错误态验证

- 操作：在 `/modules/new` 提交重复名称 `auth-service`
- 结果：表单上方显示 "module name already exists"，URL 保持 `/modules/new` 不跳转，名称与描述草稿保留，显示"草稿未保存"提示
- **结论：通过**

### 3.3 返回路径验证

- 操作：在列表设置 `queryText=auth`，进入 `auth-service` 详情页，点击"返回列表"
- 结果：URL 恢复为 `/modules?queryText=auth&statusFilter=all`，搜索框值恢复为 `auth`，状态筛选保持"全部状态"
- 依据：`module-list-search-store.ts` 使用 `persist` middleware + `sessionStorage` 持久化 `queryText` 与 `statusFilter`
- **结论：通过**

### 3.4 筛选上下文恢复验证

- 已在 3.3 一并验证：返回列表时 `queryText` 与 `statusFilter` 同时恢复
- **结论：通过**

## 4. 关键异常路径与边界语义（Task 4）

通过 curl 直接验证后端错误语义，所有错误均返回当前阶段已冻结的错误码与消息，无 500 级未收口错误。

| 异常路径 | 请求 | 实际状态码 | 错误消息 | 结论 |
| --- | --- | --- | --- | --- |
| 无效 moduleId（详情） | `GET /api/modules/not-a-uuid` | 404 | `module not found` | 通过 |
| 不存在的模块（详情） | `GET /api/modules/<uuid-not-exist>` | 404 | `module not found` | 通过 |
| 无效 moduleId（创建 release） | `POST /api/modules/not-a-uuid/releases` | 404 | `module not found` | 通过 |
| 模块不存在（创建 release） | `POST /api/modules/<uuid-not-exist>/releases` | 404 | `module not found` | 通过 |
| 重复模块名 | `POST /api/modules` name=`auth-service` | 409 | `module name already exists` | 通过 |
| 非法 status | `POST /api/modules` status=`unknown` | 400 | `invalid status` | 通过 |
| 缺失必填字段 | `POST /api/modules` 缺 description | 400 | `invalid input` | 通过 |
| 非法 release status | `POST /api/.../releases` status=`unknown` | 400 | `invalid release status` | 通过 |
| 非法 released_at | `POST /api/.../releases` released_at=`not-a-date` | 400 | `invalid input: invalid released_at (expect RFC3339)...` | 通过 |
| 重复 release version | `POST /api/.../releases` version=`1.1.0` | 409 | `release version already exists for this module` | 通过 |
| 重复产品绑定 | `POST /api/.../bindings/products` 已绑定 | 409 | `binding already exists` | 通过 |
| 不存在的 product | `POST /api/.../bindings/products` product_id=`<uuid-not-exist>` | 404 | `product not found` | 通过 |
| 重复仓库映射 | `POST /api/.../bindings/repositories` 已映射 | 409 | `binding already exists` | 通过 |
| 不存在的 repository | `POST /api/.../bindings/repositories` repository_id=`<uuid-not-exist>` | 404 | `repository not found` | 通过 |
| 无效 repository_id 格式 | `POST /api/.../bindings/repositories` repository_id=`not-a-uuid` | 404 | `repository not found` | 通过 |
| 模块不存在（绑定仓库） | `POST /api/modules/<uuid-not-exist>/bindings/repositories` | 404 | `module not found` | 通过 |

> 实现说明：`ValidateModuleID`（`backend/internal/moduleregistry/validate.go` L20-25）将"无效 ID 格式"与"模块不存在"统一映射为 `ErrModuleNotFound`（404 + `module not found`）；`ValidateProductID` / `ValidateRepositoryID` 同理。当前阶段 spec 未要求区分"ID 格式非法"与"资源不存在"，且无 500 级未收口错误，符合 spec "已冻结的错误语义"+"不得出现 500 级未收口错误"。若后续阶段需区分二者，可在此处增加 `ErrInvalidModuleID` / `ErrInvalidProductID` / `ErrInvalidRepositoryID` 哨兵并映射为 400。

候选读取与 Decision 入口边界：
- `GET /api/candidates/products` 与 `/repositories` 返回最小候选字段（id / name），未扩写为写入主线
- Decision 仅作为详情读取中的附属入口（`decision_links`），无独立 RPC，未扩写为新的写入主线
- Decision 关联数据由 `reset_module_mainline.sh` 在恢复基线时通过 name/title 查找建立，不依赖 `seed_decision_link_fixture.sql` 的显式启用

## 5. 合同与传输层一致性（Task 5）

### 5.1 `.proto` 与 HTTP DTO 语义对齐

以 `proto/psco/module_registry/v1/module_registry.proto` 为基准，核对后端 `types.go` 与前端 `types.ts`：

| `.proto` 消息 | 后端 `types.go` | 前端 `types.ts` | 对齐结论 |
| --- | --- | --- | --- |
| `Module` (id/name/description/status/created_at) | ✓ snake_case | ✓ snake_case | 通过 |
| `Release` (id/module_id/version/status/released_at) | ✓ | ✓ | 通过 |
| `ModuleListItem` (id/name/description/status/latest_release/product_bind_count/repository_bind_count) | ✓ | ✓ | 通过 |
| `ProductBinding` (product_id/product_name) | ✓ | ✓ | 通过 |
| `RepositoryMapping` (repository_id/repository_name) | ✓ | ✓ | 通过 |
| `DecisionLink` (decision_id/decision_title) | ✓ | ✓ | 通过 |
| `ModuleDetail` (module/releases/product_bindings/repository_mappings/decision_links) | ✓ | ✓ | 通过 |
| `ProductCandidate` (id/name) | ✓ | ✓ | 通过 |
| `RepositoryCandidate` (id/name) | ✓ | ✓ | 通过 |
| `CreateModuleRequest` (name/description/status) | ✓ | ✓ input 字段一致 | 通过 |
| `CreateReleaseRequest` (module_id/version/status/released_at) | ✓ module_id 由 URL 承接 | ✓ input 用 camelCase + adapter 转换 | 通过 |
| `BindModuleToProductRequest` (module_id/product_id) | ✓ module_id 由 URL 承接 | ✓ adapter 转换 productId→product_id | 通过 |
| `MapModuleToRepositoryRequest` (module_id/repository_id) | ✓ module_id 由 URL 承接 | ✓ adapter 转换 repositoryId→repository_id | 通过 |

### 5.2 RPC 到 HTTP 路由映射

| `.proto` RPC | HTTP 路由 | 状态码 | 对齐结论 |
| --- | --- | --- | --- |
| `ListModules` | `GET /api/modules` | 200 | 通过 |
| `GetModuleDetail` | `GET /api/modules/{moduleId}` | 200 | 通过 |
| `CreateModule` | `POST /api/modules` | 201 | 通过 |
| `CreateRelease` | `POST /api/modules/{moduleId}/releases` | 201 | 通过 |
| `BindModuleToProduct` | `POST /api/modules/{moduleId}/bindings/products` | 204 | 通过 |
| `MapModuleToRepository` | `POST /api/modules/{moduleId}/bindings/repositories` | 204 | 通过 |
| `ListProductCandidates` | `GET /api/candidates/products` | 200 | 通过 |
| `ListRepositoryCandidates` | `GET /api/candidates/repositories` | 200 | 通过 |

### 5.3 字段映射策略

- 响应字段：后端 `types.go` JSON 标签使用 snake_case，与 `.proto` 字段名一致，前端 `types.ts` 直接消费，无第二套语义
- 请求体字段：前端 input 类型使用 camelCase（`releasedAt` / `productId` / `repositoryId`），在 `api-adapter.ts` 显式转换为 snake_case 后发送，与后端 `types.go` 期望对齐
- `module_id`：在 `CreateRelease` / `BindModuleToProduct` / `MapModuleToRepository` 中由 URL 路径参数隐式承接（与 `.proto` 注释一致），不放在请求体

## 6. 问题收口清单

| 编号 | 问题描述 | 影响等级 | 处理结论 | 当前状态 |
| --- | --- | --- | --- | --- |
| P-01 | 前端请求体使用 camelCase，后端期望 snake_case | P0 | 修正请求体字段为 snake_case（`released_at` / `product_id` / `repository_id`），由 `api-adapter.ts` 显式转换 | 已收口 |
| P-02 | Vite dev server 偶发僵死导致前端 `ERR_CONNECTION_REFUSED` | P2 | 重启 `npm run dev` 恢复；属本地开发环境偶发，不进入代码 | 已收口 |
| P-03 | Makefile `buf breaking` 基准路径错误且吞失败 | P0 | 修正基准路径为仓库根 `.git`，移除 `\|\| echo` 吞失败逻辑（phase02-11A 已修复并复核通过） | 已收口 |
| P-04 | 文档将 `.proto` 与 `types.go/types.ts` 关系表述为"严格对齐/字段一致" | P1 | 修正为"语义对齐（HTTP DTO 通过显式映射保持语义一致）"（phase02-11A 已修复并复核通过） | 已收口 |
| P-05 | 空状态验收依赖手工 SQL，违反 spec "联调环境必须可重复建立" | P1 | 新增 `database/scripts/reset_module_mainline.sh` + `database/seeds/seed_module_mainline_baseline.sql`，提供"清空 + 恢复基线"可重复入口，支持 `--clean-only` / `--restore-only` / 默认三种模式，幂等可重复 | 已收口 |
| P-06 | 报告将 Decision 关联写成默认环境已具备，但 `run_seeds.sh` 默认不执行 `seed_decision_link_fixture.sql` | P1 | 修正 §1 环境段明确说明 `run_seeds.sh` 默认不建立 `decision_links`；phase02-12 验收基线通过 `reset_module_mainline.sh` 重建 `decision_links`（通过 name/title 查找），不依赖 fixture | 已收口 |
| P-07 | `MapModuleToRepository` 异常路径未被验收记录覆盖，但报告宣称"无剩余风险" | P1 | 补充 4 条仓库映射异常路径验证（重复映射 409 / 不存在 repository 404 / 无效 repository_id 格式 404 / 模块不存在 404），见 §4 表格 | 已收口 |

剩余风险：
- 当前阶段已冻结的错误语义全部验证通过，无 500 级未收口错误
- `ValidateModuleID` / `ValidateProductID` / `ValidateRepositoryID` 将"无效 ID 格式"与"资源不存在"统一映射为 404，spec 未要求区分二者；若后续阶段需区分，可增加 `ErrInvalid*ID` 哨兵映射为 400（P2，不阻断当前阶段）
- 所有发现的问题均已在当前阶段收口，未遗留至 `phase02-13`

## 7. 最终验收结论

`phase02-12` 验收路径全部通过：

1. **联调环境可重复建立**：`init_db.sh` → `run_seeds.sh` → `reset_module_mainline.sh` 三段式入口可重复执行；空状态验证与基线恢复均通过脚本完成，不依赖手工 SQL
2. **最小主线端到端走通**：`CreateModule` → `CreateRelease` → `BindModuleToProduct` → `MapModuleToRepository` 全链路在真实后端模式下验证通过
3. **关键状态路径已验证**：空状态入口、错误态停留、返回路径筛选恢复均符合 `phase02-09` 已冻结规则
4. **关键异常路径已覆盖**：16 条异常路径覆盖 404 / 400 / 409 错误语义（含产品绑定与仓库映射两组 404/409），无 500 级未收口错误
5. **合同与传输层一致**：`.proto` 与 HTTP DTO / 前端类型语义对齐，无第二套合同源
6. **问题已收口**：7 个问题全部在当前阶段处理完毕（P-01~P-04 历史 + P-05~P-07 本轮 GPT-5.4 复核发现），无遗留

**`Module Registry` 当前阶段最小主线已形成可运行交付物，`phase02-12` 通过验收。**
