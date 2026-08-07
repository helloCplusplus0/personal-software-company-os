# Tasks

- [x] Task 1: 实现 `Product Registry` 后端模块骨架与支撑文件。
  - [x] SubTask 1.1: 创建 `backend/internal/productregistry/` 下的 `handler/`、`service/`、`repository/`、`candidate/`、`types.go`、`errors.go`、`validate.go`
  - [x] SubTask 1.2: 在 `types.go` 中定义与 `proto/psco/product_registry/v1/product_registry.proto` 对齐的对象、请求、响应与查询参数
  - [x] SubTask 1.3: 在 `errors.go` 与 `handler/response.go` 中收口 `Product Registry` 的业务错误与 HTTP 状态码映射

- [x] Task 2: 实现 `Repository Binding` 后端模块骨架与支撑文件。
  - [x] SubTask 2.1: 创建 `backend/internal/repositorybinding/` 下的 `handler/`、`service/`、`repository/`、`candidate/`、`types.go`、`errors.go`、`validate.go`
  - [x] SubTask 2.2: 在 `types.go` 中定义与 `proto/psco/repository_binding/v1/repository_binding.proto` 对齐的对象、请求、响应与查询参数
  - [x] SubTask 2.3: 在 `errors.go` 与 `handler/response.go` 中收口 `Repository Binding` 的业务错误与 HTTP 状态码映射

- [x] Task 3: 实现 `Product Registry` 的 repository / candidate / service / handler 主线。
  - [x] SubTask 3.1: 在 `repository/product_store.go` 中实现 `products` 的 `Create / GetByID / List / Exists`
  - [x] SubTask 3.2: 在 `repository/binding_store.go` 中实现 `ProductModuleSummaryRead` 与 `ProductRepositorySummaryRead`
  - [x] SubTask 3.3: 在 `candidate/module_candidate_read.go` 中实现 `ListProductModuleCandidates` 与 `ModuleExists`
  - [x] SubTask 3.4: 在 `service/query_service.go` 中实现 `ListProducts / GetProductDetail / ListProductModuleCandidates`
  - [x] SubTask 3.5: 在 `service/command_service.go` 中实现 `CreateProduct / BindModuleToProduct`
  - [x] SubTask 3.6: 在 `handler/query_handler.go` 与 `handler/command_handler.go` 中挂接对应 HTTP 入口

- [x] Task 4: 实现 `Repository Binding` 的 repository / candidate / service / handler 主线。
  - [x] SubTask 4.1: 在 `repository/repository_store.go` 中实现 `repositories` 的 `Create / GetByID / List / Exists`
  - [x] SubTask 4.2: 在 `repository/binding_store.go` 中实现 `RepositoryProductSummaryRead`、`RepositoryModuleSummaryRead`、`BindRepositoryToProduct`、`MapModuleToRepository`
  - [x] SubTask 4.3: 在 `candidate/product_candidate_read.go` 与 `candidate/module_candidate_read.go` 中实现两条候选读取与存在性校验
  - [x] SubTask 4.4: 在 `service/query_service.go` 中实现 `ListRepositories / GetRepositoryDetail / ListRepositoryProductCandidates / ListRepositoryModuleCandidates`
  - [x] SubTask 4.5: 在 `service/command_service.go` 中实现 `CreateRepository / BindRepositoryToProduct / MapModuleToRepository`
  - [x] SubTask 4.6: 在 `handler/query_handler.go` 与 `handler/command_handler.go` 中挂接对应 HTTP 入口

- [x] Task 5: 实现 `0006_product_repository_binding_mainline.sql` 数据主线 migration。
  - [x] SubTask 5.1: 原位升级 `products`，新增 `description / status`、检查约束与索引
  - [x] SubTask 5.2: 原位升级 `repositories`，新增 `url / provider / status`、检查约束与索引
  - [x] SubTask 5.3: 新增 `product_repositories` 表、唯一约束与读取索引
  - [x] SubTask 5.4: 为历史 `products / repositories` 数据实现幂等兼容回填

- [x] Task 6: 实现 `phase04` 基线 seed 与重置脚本。
  - [x] SubTask 6.1: 创建 `database/seeds/seed_product_repository_mainline_baseline.sql`
  - [x] SubTask 6.2: 更新 `database/seeds/seed_readonly_prereqs.sql` 的 `products / repositories` 完整字段插入
  - [x] SubTask 6.3: 创建 `database/scripts/reset_product_repository_mainline.sh`，支持 `--clean-only / --restore-only / default`
  - [x] SubTask 6.4: 确认 seed 与 reset 对 `Product A / Product B / Product C / main-repo / mirror-repo` 的兼容恢复不被破坏

- [x] Task 7: 实现旧 `Module Registry` 入口的兼容委派。
  - [x] SubTask 7.1: 若保留旧候选读取入口，则委派到 `Repository Binding` canonical 实现
  - [x] SubTask 7.2: 若保留旧模块中心绑定入口，则把 `BindModuleToProduct` 委派到 `Product Registry`、把 `MapModuleToRepository` 委派到 `Repository Binding`
  - [x] SubTask 7.3: 清理旧 `moduleregistry/service/` 中不应继续保留的长期 owner 逻辑

- [x] Task 8: 完成 `chi` 路由装配与模块接线。
  - [x] SubTask 8.1: 在 `backend/internal/platform/router.go` 中新增 `mountProductRegistry` 与 `mountRepositoryBinding`
  - [x] SubTask 8.2: 在 `server.go` 的 `/api` 子路由下挂接两个新模块
  - [x] SubTask 8.3: 对齐 `.proto` 已冻结的 RPC -> HTTP 映射与现有 `chi.Route / Mount` 组织模式

- [x] Task 9: 验证后端、migration、reset 与兼容委派主线可运行。
  - [x] SubTask 9.1: 运行 `go build ./...` 与必要静态检查，确认后端可编译
  - [x] SubTask 9.2: 验证 `0006` migration 可在现有 `0001-0005` 基础上运行
  - [x] SubTask 9.3: 验证 `reset_product_repository_mainline.sh` 三种模式可重复执行
  - [x] SubTask 9.4: 通过真实 HTTP 请求验证列表、详情、创建、三类绑定与候选读取接口
  - [x] SubTask 9.5: 验证旧 transport 入口若保留，其行为确实是兼容委派而非独立 owner

# Task Dependencies

- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 5`
- `Task 7` depends on `Task 3` and `Task 4`
- `Task 8` depends on `Task 3` and `Task 4`
- `Task 9` depends on `Task 5`, `Task 6`, `Task 7`, and `Task 8`
