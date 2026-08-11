# Phase07-09 落实 Go 后端业务传输主线切换 Spec

## Why

`phase07-07 formal spec` 已经冻结了后端正式传输主线的目标形态：`.proto` 作为唯一合同源、`chi` 只保留 `/api` shell 与非业务端点职责、业务流量统一进入 Connect handler、错误语义通过单值映射进入 Connect code、legacy/compat 路由按 inventory 在 `phase07-09 / 10` 退场。`phase07-08` 又已经把 `buf` 三插件生成链和 Go Connect 产物主线落实到了仓库里。

当前真正缺的，不再是设计或生成链，而是把现有 `server.go / router.go / backend/internal/*/handler` 为中心的手写 JSON 业务主线切换到 Connect handler / service implementation，并按冻结 inventory 删除后端候选 compat 入口。`phase07-09` 因此是“后端 transport 正式切换”本身，而不是再写一轮设计文档。

## What Changes

- 将 `phase01 ~ phase06` 的 34 条业务 RPC 在 Go 后端切换到 Connect handler / service implementation 主线
- 在 `backend/internal/platform` 收敛为单一 `/api` shell + generated procedure path 挂载模式
- 将 `server.go` 与 `router.go` 从手写 JSON 业务主线重构为 Connect canonical mount + compat 过渡组
- 引入单值 Connect interceptor / 错误映射承接位，统一 domain error → Connect code
- 删除 `phase07-03` / `phase07-07` 明确要求在本阶段退场的 2 条候选 compat 路由：`GET /api/candidates/products` 与 `GET /api/candidates/repositories`
- 清退 canonical 业务接口对应的手写 JSON handler 主线，保留的 handler 仅允许是 `phase07-10` 前仍需存在的绑定 compat 薄壳
- **BREAKING**：后端浏览器对外业务路径不再以 `/api/modules`、`/api/products` 等手写 JSON 路由为 canonical 业务入口，而统一切换为 `/api/<Connect procedure path>`

## Impact

- Affected specs:
  - `phase07_03_freeze_compat_migration_exit_criteria`
  - `phase07_04_design_go_connect_handler_service_chi_mount`
  - `phase07_06_design_business_api_migration_matrix_regression_acceptance`
  - `phase07_07_formal_transport_mainline_cutover_spec`
  - `phase07_08_land_buf_generation_connect_contract_mainline`
- Affected code:
  - `backend/internal/platform/server.go`
  - `backend/internal/platform/router.go`
  - `backend/internal/platform/` 下新增 Connect transport 装配 / interceptor / 错误映射文件
  - `backend/internal/*/handler/`
  - `backend/internal/*/service/`
  - `backend/internal/gen/connect/**`
  - `backend/cmd/**`、后端启动与验收入口

## ADDED Requirements

### Requirement: Go 后端必须切换到 Connect canonical business transport

系统 SHALL 将 `phase01 ~ phase06` 的 canonical 业务接口切换到 Connect handler / service implementation 主线，不再让手写 JSON handler 继续作为正式业务入口。

#### Scenario: 34 条业务 RPC 接入 Connect handler

- **WHEN** 团队执行 `phase07-09`
- **THEN** 34 条 proto-defined business RPC 必须全部有对应的 Go Connect handler 承接
- **AND** handler 必须基于 `backend/internal/gen/connect/**` 的 generated service handler 构造
- **AND** handler implementation 必须使用 simple 模式签名（`context.Context, *pb.Request -> *pb.Response, error`）
- **AND** 不得并存第二套 hand-written JSON canonical business handler

### Requirement: `/api` 子路由必须只承担 Connect shell 与 compat 过渡组

系统 SHALL 保持 `chi` 根路由器只负责 middleware、healthz 与 `/api` shell；`/api` 下的业务承接必须收敛为 Connect procedure path 与显式 compat 过渡组，不得继续散落多套长期 JSON 业务路由。

#### Scenario: `/api` 子路由组织收敛

- **WHEN** `server.go` 装配业务路由
- **THEN** 必须继续使用 `r.Route("/api", ...)` 作为单一业务前缀壳
- **AND** Connect handler 必须显式消费 generated `path`，例如 `r.Handle(path, handler)` 或先聚合到单一 `http.ServeMux` 再挂入 `/api`
- **AND** 任何 compat 路由都必须集中在 `mountCompatRoutes` 或等价单一承接位中
- **AND** 不得在 `/api` 子路由内部再次写入 `"/api/..."` 绝对路径

### Requirement: Connect interceptor 与错误映射必须单值化

系统 SHALL 为 Connect 业务主线提供统一的 interceptor 与错误映射承接位，避免模块各自维护第二套横切行为或错误语义。

#### Scenario: 错误语义统一

- **WHEN** Connect implementation 返回 domain error
- **THEN** 必须通过单一 `MapToConnectError` 或等价函数映射为 Connect code
- **AND** `NotFound / InvalidArgument / FailedPrecondition / AlreadyExists / Internal` 等错误语义必须保持单值映射
- **AND** module/product/repository/decision/dashboard/onboarding/export/backup/reuse 等模块不得各自维护独立错误表

### Requirement: 现有 service 分层必须保留，transport 层只做解包与组装

系统 SHALL 在切换 transport 主线时继续保留既有 `repository / candidate / service / platform` 分层，Connect implementation 只负责 transport 适配，不得把业务逻辑重新塞回 handler。

#### Scenario: Connect implementation 边界

- **WHEN** 为某模块实现 Connect handler
- **THEN** implementation 必须调用现有 `QueryService / CommandService` 或等价 service 层
- **AND** transport 层只负责 request 解包、service 调用、response 组装、错误映射
- **AND** 不得在 Connect implementation 中直接新增跨模块 SQL 或绕过既有 service owner

### Requirement: phase07-09 必须退场 2 条候选 compat 路由

系统 SHALL 在本阶段完成 `GET /api/candidates/products` 与 `GET /api/candidates/repositories` 两条候选 compat 入口的后端退场，不得推迟到 `phase07-10` 或 `phase07-11`。

#### Scenario: 候选 compat 入口退场

- **WHEN** `phase07-09` 完成
- **THEN** `router.go` 中不得再注册 `/candidates/products` 与 `/candidates/repositories`
- **AND** `moduleregistry/handler/query_handler.go` 中不得再保留 `ListProductCandidates`、`ListRepositoryCandidates` 及对应 `ProductCandidateReader` / `RepositoryCandidateReader`
- **AND** `ProductRegistryService.ListProducts` 与 `RepositoryBindingService.ListRepositories` 的 Connect path 必须成为唯一后端替代路径
- **AND** 旧路径必须返回 404

### Requirement: Module-centered 绑定 compat 入口只允许作为过渡薄壳保留

系统 SHALL 允许 `POST /api/modules/{moduleId}/bindings/products` 与 `POST /api/modules/{moduleId}/bindings/repositories` 在 `phase07-09` 仍保留为后端 compat 薄壳，但不得再把它们当作 canonical business owner。

#### Scenario: 绑定 compat 入口边界

- **WHEN** 团队完成 `phase07-09`
- **THEN** 两条绑定 compat 路由如果继续存在，必须只出现在单一 compat 过渡组中
- **AND** 必须显式标注最晚在 `phase07-10` 退场
- **AND** 对应 canonical owner 必须已经是 `ProductRegistryService.BindModuleToProduct` 与 `RepositoryBindingService.MapModuleToRepository` 的 Connect 实现

### Requirement: chi 只保留 shell 与非业务端点职责

系统 SHALL 在后端主线切换后，把 `chi` 的职责收敛为 middleware shell 与非业务基础设施端点，不再承担业务合同定义或 hand-written canonical route 组织职责。

#### Scenario: chi shell 化

- **WHEN** 审核 `server.go` / `router.go`
- **THEN** `chi` 必须只保留 `RequestID / RealIP / Logger / Recoverer / Timeout / CORS` 等 HTTP middleware
- **AND** 保留 `GET /healthz`
- **AND** 未来可扩展的 `readyz / metrics / debug` 可继续作为 infra keep list
- **AND** canonical business path 不能再通过 `r.Get("/modules", ...)`、`r.Post("/products", ...)` 等 JSON 路由定义

### Requirement: 后端验收必须以 Connect path 为主线

系统 SHALL 要求 `phase07-09` 的后端验收以 Connect procedure path 为主线，而不是继续以旧 JSON 业务路径作为“默认正常路径”。

#### Scenario: 后端回归验证

- **WHEN** 团队验证后端切换结果
- **THEN** 必须逐条验证 canonical RPC 的 Connect path 返回 200
- **AND** 必须验证 `/api` 仍是浏览器侧唯一基址
- **AND** 必须验证候选 compat 旧路径 404
- **AND** 必须保留 `phase07-10` 仍未退场的绑定 compat 入口状态说明，避免提前误删

## MODIFIED Requirements

### Requirement: `router.go` 的业务组织方式

`router.go` SHALL 从“每个模块直接注册手写 JSON 业务路径”的组织方式，修改为“canonical Connect mount + 单一 compat 过渡组”的组织方式。

#### Scenario: `router.go` 重构后

- **WHEN** 实现者查看 `router.go`
- **THEN** 必须能看到 canonical Connect mount 的集中装配逻辑
- **AND** 必须能看到单一 compat 过渡组
- **AND** 不得继续由 `mountModuleRegistry / mountDecisionCenter / mountProductRegistry ...` 直接注册长期 JSON canonical 路由

### Requirement: `phase07` 后端阶段职责划分

`phase07-09` SHALL 负责后端主线切换与候选 compat 路由退场；`phase07-10` 继续负责前端 adapter / mutation owner 退场，不得把前端工作并入本阶段 DoD。

#### Scenario: 阶段边界

- **WHEN** 团队执行 `phase07-09`
- **THEN** 当前阶段完成条件必须聚焦于后端 transport 切换、canonical Connect 可用、L1/L2 compat 退场
- **AND** 不得把前端 13 项 adapter 删除作为 `phase07-09` 完成条件
- **AND** 不得把 L3/L4 绑定 compat 路由的删除提前并入本阶段

## REMOVED Requirements

### Requirement: 手写 JSON 业务 handler 继续作为 canonical 主线

**Reason**: `phase07-07 formal spec` 已明确业务 transport 的唯一正式主线是 Connect handler；继续保留 hand-written JSON canonical business handler 会形成与 `.proto + Connect` 并列的第二套主线。

**Migration**: 将现有 JSON canonical handler 替换为 Connect implementation；保留的 compat 入口统一回收到单一 compat 过渡组，并按 `phase07-03` 冻结时点删除。

### Requirement: 候选 compat 入口可延后到收口时再看

**Reason**: `phase07-03` 已冻结候选读取 compat 入口必须在 `phase07-09` 退场，推迟到 `phase07-10/11` 会破坏子任务链职责边界。

**Migration**: 在本阶段完成 `GET /api/candidates/products` 与 `GET /api/candidates/repositories` 的路由删除、handler 删除、替代 Connect path 验证与 404 回归验证。
