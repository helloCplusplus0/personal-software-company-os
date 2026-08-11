- [x] 已明确 `phase07-09` 的目标是 Go 后端业务传输主线切换，不把前端 phase07-10 工作并入当前范围
- [x] 已明确 `phase07-09` 直接继承 `phase07-07 formal spec` 的后端正式传输主线规则
- [x] 已明确 34 条业务 RPC 必须全部由 Go Connect handler 承接
- [x] 已明确 handler implementation 使用 simple 模式签名
- [x] 已明确 `server.go` 继续以 `r.Route("/api", ...)` 作为唯一业务前缀壳
- [x] 已明确 canonical Connect handler 必须显式消费 generated `path`
- [x] 已明确 compat 路由必须集中在单一 `mountCompatRoutes` 或等价承接位
- [x] 已明确不得在 `/api` 子路由内部再次写入 `"/api/..."` 绝对路径
- [x] 已明确 `chi` middleware 与 Connect interceptor 的职责边界
- [x] 已明确 domain error → Connect code 的单值映射承接位
- [x] 已明确 transport 层只做 request 解包、service 调用、response 组装与错误映射
- [x] 已明确继续复用既有 `QueryService / CommandService` 与 candidate/repository owner
- [x] 已明确 `router.go` 需从长期 JSON canonical 路由组织重构为 canonical Connect mount + compat 过渡组
- [x] 已明确 hand-written JSON canonical handler 必须退出主线，仅允许 L3/L4 compat 薄壳暂留
- [x] 已明确 `chi` 只保留 shell 与非业务端点职责
- [x] 已明确 `GET /api/candidates/products` 与 `GET /api/candidates/repositories` 必须在 `phase07-09` 退场
- [x] 已明确需要删除 `ListProductCandidates` / `ListRepositoryCandidates` 及对应 reader 字段
- [x] 已明确替代 Connect path 回归证据：`ProductRegistryService.ListProducts` 与 `RepositoryBindingService.ListRepositories` 返回 200
- [x] 已明确 L1/L2 旧路径必须返回 404
- [x] 已明确 L3/L4 绑定 compat 路由在本阶段只允许作为过渡薄壳保留，并标注最晚 `phase07-10` 退场
- [x] 已明确后端回归以 Connect procedure path 为主线
- [x] 已明确 `(cd backend && go build ./...)`、Connect path `curl`、legacy 路径 404 与 `/api` 单一基址验证要求
- [x] 已验证本 spec 对齐 `phase07-03`、`phase07-04`、`phase07-06`、`phase07-07`、`phase07-08` 与当前源码现状

## 2026-08-11 补充验证与修复

- [x] 修复 `moduleregistry/handler/query_handler.go`：删除 `ListProductCandidates`、`ListRepositoryCandidates` 方法及 `ProductCandidateReader`/`RepositoryCandidateReader` 类型定义（spec 要求删除，原实现遗漏）
- [x] 修复 `NewQueryHandler` 签名：从 3 参数降为 1 参数（移除 `pc`/`rc`）
- [x] 验证 `go build ./...` 通过
- [x] 验证 9 个 Connect 模块全部读端点（List/Get）返回 200 或 Connect 错误码（非路由 404）
- [x] 验证 Connect 写端点 `CreateModule`/`CreateProduct`/`CreateRepository`/`ExportCoreAssets`/`CreateInstanceBackup` 返回 200
- [x] 验证 `CreateDecision` 返回 `{"code":"invalid_argument"}`（Connect CodeInvalidArgument，非路由 404）
- [x] 验证 Connect RPC #33 `ListProductCandidates`（via ModuleRegistry Connect）返回 200
- [x] 验证 Connect RPC #34 `ListRepositoryCandidates`（via ModuleRegistry Connect）返回 200
- [x] 验证 L1 `GET /api/candidates/products` 返回 404
- [x] 验证 L2 `GET /api/candidates/repositories` 返回 404
- [x] 验证 L3 `POST /api/modules/{moduleId}/bindings/products` compat 路由挂载正常
- [x] 验证 L4 `POST /api/modules/{moduleId}/bindings/repositories` compat 路由挂载正常
- [x] 验证 `/api` 单一基址保持，`/healthz` 非业务端点正常

### 2026-08-11 运行时验收实证

- [x] 为避免本机常驻 `:8081` 进程污染结果，使用当前工作树在隔离端口 `:18081` 启动后端实例，并在验收完成后关闭该实例
- [x] 验证 `GET http://127.0.0.1:18081/healthz` 返回 `200 {"status":"ok"}`
- [x] 验证 `POST http://127.0.0.1:18081/api/psco.module_registry.v1.ModuleRegistryService/ListModules` 返回 `200`
- [x] 验证 `POST http://127.0.0.1:18081/api/psco.product_registry.v1.ProductRegistryService/ListProducts` 返回 `200`
- [x] 验证 `POST http://127.0.0.1:18081/api/psco.repository_binding.v1.RepositoryBindingService/ListRepositories` 返回 `200`
- [x] 验证 `GET http://127.0.0.1:18081/api/candidates/products` 返回 `404 page not found`
- [x] 验证 `GET http://127.0.0.1:18081/api/candidates/repositories` 返回 `404 page not found`
- [x] 验证 `POST http://127.0.0.1:18081/api/modules/test-module/bindings/products` 携带 malformed JSON 时返回 `400 {"error":"invalid json body"}`，确认 compat 薄壳已恢复旧 JSON 解析错误语义
- [x] 验证 `POST http://127.0.0.1:18081/api/modules/test-module/bindings/products` 携带不存在的资源标识时返回 `404 {"error":"module not found"}`，确认 compat 薄壳不再把 owner 错误粗化为 `500`
- [x] 通过 Connect 写路径创建临时资源：`module_id=69facd3b-2a76-4821-97b7-681bdaa9e089`、`product_id=4e153734-6f55-455c-84bc-94e0f07e7aef`、`repository_id=22fcd18b-99bd-4c62-a5bb-3e1bf9571436`
- [x] 验证 `POST http://127.0.0.1:18081/api/modules/69facd3b-2a76-4821-97b7-681bdaa9e089/bindings/products` 返回 `204 No Content`
- [x] 验证 `POST http://127.0.0.1:18081/api/modules/69facd3b-2a76-4821-97b7-681bdaa9e089/bindings/repositories` 返回 `204 No Content`
