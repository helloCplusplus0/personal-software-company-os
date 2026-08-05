# Tasks

- [x] Task 1: 冻结本地 PostgreSQL 共享实例的使用策略。
  - [x] SubTask 1.1: 明确 `phase02-11` 复用现有 Podman PostgreSQL 容器，不新开第二个本地数据库容器
  - [x] SubTask 1.2: 明确 PSCO 在共享实例中使用独立数据库（推荐 `psco_development`），而不是复用 `rento_production`
  - [x] SubTask 1.3: 明确本地 `DATABASE_URL` 必须是显式最终值，且敏感凭据不写入仓库源码

- [x] Task 2: 建立数据库迁移主线。
  - [x] SubTask 2.1: 建立 `modules / module_releases / product_modules / module_repositories` 的迁移定义
  - [x] SubTask 2.2: 建立 `decisions / decision_links / products / repositories` 的最小只读前提结构
  - [x] SubTask 2.3: 保证字段、关系约束与 `phase02-09` 正式规格正文一致

- [x] Task 3: 建立支撑联调的最小数据库初始化能力。
  - [x] SubTask 3.1: 提供 PSCO 独立数据库的初始化入口
  - [x] SubTask 3.2: 提供 `products / repositories` 候选读取所需的最小种子数据或测试 fixture
  - [x] SubTask 3.3: 提供 `decisions / decision_links` 只读入口所需的最小示例数据

- [x] Task 4: 搭建 `Module Registry` 后端包结构与运行入口。
  - [x] SubTask 4.1: 建立 `backend/internal/moduleregistry/handler`、`service`、`repository`、`candidate` 四层结构
  - [x] SubTask 4.2: 落地 `query_handler.go`、`command_handler.go`、`query_service.go`、`command_service.go`
  - [x] SubTask 4.3: 提供后端运行入口与数据库连接接线

- [x] Task 5: 实现读组与候选读取接口。
  - [x] SubTask 5.1: 实现 `ModuleListRead`
  - [x] SubTask 5.2: 实现 `ModuleDetailRead`，并内嵌承接 `Decision` 附属读取
  - [x] SubTask 5.3: 实现 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead`

- [x] Task 6: 实现写组接口。
  - [x] SubTask 6.1: 实现 `ModuleCreateWrite`
  - [x] SubTask 6.2: 实现 `ModuleReleaseWrite`
  - [x] SubTask 6.3: 实现 `ModuleBindingWrite`，统一承接产品绑定与仓库映射

- [x] Task 7: 完成前后端切换与运行验证。
  - [x] SubTask 7.1: 让 `phase02-10` 前端具备切换到真实后端的明确对接点
  - [x] SubTask 7.2: 验证后端读写接口可运行
  - [x] SubTask 7.3: 验证数据主线与已冻结边界一致，且未引入第二套数据主线

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 7` depends on `Task 5` and `Task 6`

# 实现产物索引

## 数据库层（database/）
- `database/migrations/0001_module_registry_mainline.sql` — modules / module_releases 主线表
- `database/migrations/0002_readonly_prereqs.sql` — products / repositories / decisions / decision_links 只读前提
- `database/migrations/0003_module_registry_bindings.sql` — product_modules / module_repositories 关联写入表
- `database/seeds/seed_readonly_prereqs.sql` — products / repositories / decisions 最小种子数据（幂等）
- `database/seeds/seed_decision_link_fixture.sql` — decision_links 联调 fixture（可选，需已有模块）
- `database/scripts/init_db.sh` — PSCO 独立数据库初始化入口（幂等建库，复用共享容器）
- `database/scripts/run_seeds.sh` — 种子数据统一执行入口（幂等，支持 fixture 可选启用）

## 后端（backend/）
- `backend/go.mod` / `backend/go.sum` — Go 模块定义（chi v5 + pgx v5 + godotenv + uuid）
- `backend/cmd/server/main.go` — 运行入口（加载 .env → 连接池 → 迁移 → 可选种子 → HTTP 服务 → 优雅关闭）
- `backend/internal/platform/config.go` — 环境配置加载
- `backend/internal/platform/db.go` — pgxpool 连接池
- `backend/internal/platform/migrate.go` — 文件系统迁移运行器（schema_migrations 表 + 事务执行）+ RunSeeds 种子执行器
- `backend/internal/platform/server.go` — chi 路由装配 + HTTP 服务器 + CORS 中间件
- `backend/internal/platform/router.go` — Module Registry 路由挂载（组合根，避免导入循环）
- `backend/internal/moduleregistry/types.go` — API 消息结构（合同与存储解耦）
- `backend/internal/moduleregistry/errors.go` — 业务错误哨兵值
- `backend/internal/moduleregistry/validate.go` — UUID 格式校验辅助
- `backend/internal/moduleregistry/handler/response.go` — JSON 响应与错误到 HTTP 状态码映射
- `backend/internal/moduleregistry/handler/query_handler.go` — 读组 + 候选读取 HTTP 入口
- `backend/internal/moduleregistry/handler/command_handler.go` — 写组 HTTP 入口
- `backend/internal/moduleregistry/service/query_service.go` — 读组编排（含 Decision 内嵌附属读取）
- `backend/internal/moduleregistry/service/command_service.go` — 写组编排（创建 + 版本 + 绑定）
- `backend/internal/moduleregistry/repository/module_store.go` — modules 表数据访问
- `backend/internal/moduleregistry/repository/release_store.go` — module_releases 表数据访问
- `backend/internal/moduleregistry/repository/binding_store.go` — product_modules / module_repositories / decision_links 数据访问
- `backend/internal/moduleregistry/candidate/product_candidate_read.go` — Product 候选读取（phase02 临时承接）
- `backend/internal/moduleregistry/candidate/repository_candidate_read.go` — Repository 候选读取（phase02 临时承接）
- `backend/.env.example` — 后端环境变量样例（DATABASE_URL 显式最终值）

## 前端切换对接点（frontend/）
- `frontend/src/features/module-registry/data/mock-adapter.ts` — phase02-10 mock 适配层（原 module-registry-adapter.ts 内容）
- `frontend/src/features/module-registry/data/api-adapter.ts` — 真实后端 API 适配层（函数签名与 mock 一致）
- `frontend/src/features/module-registry/data/module-registry-adapter.ts` — 切换入口（VITE_USE_REAL_API 控制）
- `frontend/src/vite-env.d.ts` — Vite 环境变量类型声明
- `frontend/.env.example` / `frontend/.env` — 前端环境变量（默认 mock 模式）

## API 端点矩阵

| 动作语义 | 接口分组 | HTTP 方法 | 路径 |
| --- | --- | --- | --- |
| 列表读取 | 读组 | GET | `/api/modules` |
| 详情读取 | 读组 | GET | `/api/modules/{moduleId}` |
| Product 候选读取 | 候选读取 | GET | `/api/candidates/products` |
| Repository 候选读取 | 候选读取 | GET | `/api/candidates/repositories` |
| CreateModule | 写组 | POST | `/api/modules` |
| CreateRelease | 写组 | POST | `/api/modules/{moduleId}/releases` |
| BindModuleToProduct | 写组 | POST | `/api/modules/{moduleId}/bindings/products` |
| MapModuleToRepository | 写组 | POST | `/api/modules/{moduleId}/bindings/repositories` |
| 健康检查 | — | GET | `/healthz` |

## 运行时验证记录

- `go build ./...` 通过，`go vet ./...` 无告警
- 3 个迁移文件成功应用（幂等可重复执行）
- `database/scripts/init_db.sh` 幂等验证通过（数据库已存在时正确跳过）
- `database/scripts/run_seeds.sh` 幂等验证通过（ON CONFLICT 跳过已有数据；fixture 守卫块在无模块时跳过）
- 后端 `RUN_SEEDS_ON_BOOT=true` 启动流程验证通过（迁移后自动执行 2 个种子文件）
- 后端 `RUN_SEEDS_ON_BOOT=false` 启动流程验证通过（跳过种子并打印清晰提示）
- API 端到端验证：healthz / 空列表 / 种子候选 / 创建模块（含重复名 409、非法状态 400）/ 列表计数 / 详情读取 / 版本登记（含重复版本 409）/ 产品绑定（含重复 409、不存在 404）/ 仓库映射 / Decision fixture / 筛选（queryText / statusFilter）
- 非 UUID 路径参数返回 404（而非 500）
- 前端 `npm run build` 在 mock 与 real API 两种模式下均通过
- 浏览器端到端：列表页读取真实后端数据 → 点击列表项以 UUID 进入详情页 → 显示模块、版本、产品绑定、仓库映射
