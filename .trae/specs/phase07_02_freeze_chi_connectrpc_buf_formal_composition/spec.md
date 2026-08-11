# phase07-02 冻结 `chi + ConnectRPC + buf` 正式组合方式 Spec

> **执行产出**：`frozen_scope.md` — 包含 chi 职责冻结、Connect handler 挂载方式、buf.gen.yaml 目标形态（3 插件）、前端 Connect 客户端组合、`/api` 全链路承接关系。
> **执行日期**：2026-08-11
> **参考**：Context7 `/connectrpc/connect-go` + `/connectrpc/connect-es` 最新文档
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

当前仓库仍处于“`.proto` 已冻结、`chi + JSON HTTP` 仍承接业务 transport、前端仍以手写 `fetch` 为主”的过渡态。`phase07` 如果不先把 `chi`、Connect handler、`buf` 生成链、前端客户端与 `/api` 访问链的正式组合方式冻结成单值规则，后续实现会很容易长出第二套基址、第二套产物目录或第二套前端调用组织。

## What Changes

- 冻结 `chi` 的唯一正式职责：顶层 `HTTP shell`、统一 middleware、非业务基础设施端点、以及 `/api` 业务入口的 mount 外壳
- 冻结 Connect handler 的挂载方式：canonical 业务 RPC 统一在单一 `/api` 前缀下承接，不新增 `/rpc`、`/connect`、`/grpc` 等并列业务基址
- 冻结 `buf.gen.yaml` 的正式插件矩阵与产物落点：
  - `buf.build/protocolbuffers/go` -> `backend/internal/gen/proto/`
  - `buf.build/connectrpc/gosimple` -> `backend/internal/gen/connect/`
  - `buf.build/bufbuild/es` -> `frontend/src/gen/proto/`
- 冻结前端客户端生成与运行时组合：TypeScript 继续只生成 `Protobuf-ES` 产物，由共享 Connect transport + `createClient()` 消费 `*_pb.ts` 中的服务描述符，不新增第二个 TS 代码生成根
- 冻结浏览器侧 `/api` 前缀、Vite dev proxy、本地后端启动与未来 Caddy 反代的承接关系，保证外部访问面保持单值
- **BREAKING**：`chi` 不再被允许作为 canonical 业务合同定义层；手写 JSON 业务 handler 只能视为迁移中的临时 compat 资产

## Impact

- Affected specs:
  - `phase07-04` Go Connect handler / chi mount 设计
  - `phase07-05` 前端生成客户端与 query / application 迁移设计
  - `phase07-06` transport 实现与回归矩阵
  - `phase07-11` 验收核销
- Affected code:
  - `proto/buf.gen.yaml`
  - `proto/README.md`
  - `proto/Makefile`
  - `backend/internal/platform/server.go`
  - `backend/internal/platform/router.go`
  - `backend/internal/gen/proto/`
  - `backend/internal/gen/connect/`
  - `frontend/src/gen/proto/`
  - `frontend/vite.config.ts`
  - `frontend/.env.example`
  - 前端共享 transport / feature query / application 承接位

## ADDED Requirements

### Requirement: `chi` 只承担 HTTP 外壳与非业务基础设施端点职责

系统 SHALL 将 `chi` 冻结为唯一顶层 HTTP shell，只承接：

- 根级 middleware（如 `request id / logging / recover / timeout / cors`）
- 非业务基础设施端点（`healthz / readyz / metrics / debug / pprof`）
- `/api` 业务前缀的 mount 外壳

`chi` SHALL NOT 再承担 canonical 业务合同定义职责，也不得继续以手写路由路径表充当长期业务 API 真相源。

#### Scenario: Infra endpoints remain on chi

- **WHEN** 服务启动并暴露基础设施端点
- **THEN** `healthz / readyz / metrics / debug / pprof` 继续由原生 `chi + net/http` 承接
- **AND** 这些端点不进入 `.proto`

#### Scenario: Business APIs no longer use chi route declarations as canonical contract

- **WHEN** 业务模块进入 `phase07` 正式 transport 迁移
- **THEN** canonical 业务接口的长期合同定义以 `.proto + Connect procedure` 为准
- **AND** `chi` 仅保留 mount 与 middleware 组织职责

### Requirement: Connect handler 必须在单一 `/api` 前缀下挂载

系统 SHALL 采用“`chi` 顶层根路由 + `/api` 单一业务前缀 + Connect generated handler”组合方式。

canonical 业务 RPC 的对外访问面 SHALL 统一表现为：

- `/api/{Connect procedure path}`

其中 Connect generated handler 返回的 procedure path 与 `http.Handler` 必须通过统一 mount 方案纳入 `/api` 外壳下，不得为 Connect 单独新增第二套公网或本地开发基址。

#### Scenario: Connect procedures mount under shared /api prefix

- **WHEN** 某个 canonical 业务 service 被 Connect 生成并挂载
- **THEN** 浏览器与前端客户端仅通过 `/api/...` 访问该 RPC
- **AND** 不出现 `/rpc/...`、`/connect/...`、`/grpc/...` 等并列业务前缀

#### Scenario: Existing infra endpoints remain outside /api Connect tree

- **WHEN** 请求目标是 `healthz / readyz / metrics / debug / pprof`
- **THEN** 请求继续走原生 `chi + net/http`
- **AND** 不被错误迁入 Connect business tree

### Requirement: `buf` 生成链必须冻结为单一正式插件矩阵与单一产物根

系统 SHALL 继续以 `proto/Makefile` 和 `proto/buf.gen.yaml` 作为唯一 proto 工具链入口，并冻结 `buf.gen.yaml` 的正式插件矩阵为：

- `buf.build/protocolbuffers/go`
- `buf.build/connectrpc/gosimple`
- `buf.build/bufbuild/es`

系统 SHALL 冻结产物落点为：

- Go protobuf messages -> `backend/internal/gen/proto/`
- Go Connect handlers / procedure constants -> `backend/internal/gen/connect/`
- TypeScript protobuf / service descriptors -> `frontend/src/gen/proto/`

系统 SHALL NOT 新增第二个 `buf.gen.yaml`、第二个 proto workspace、第二个 Go 生成根或第二个 TS 生成根。

#### Scenario: Single generation command updates all official artifacts

- **WHEN** 在 `proto/` 目录执行 `make gen` 或 `buf generate`
- **THEN** Go protobuf、Go Connect、TypeScript protobuf 产物都从同一个 `buf.gen.yaml` 生成
- **AND** 产物仅落到冻结的三个正式目录根

#### Scenario: Go and TS output roots stay single-valued

- **WHEN** 后续 `phase07` 子任务继续扩写 service 或字段
- **THEN** 不得把 Go Connect 产物写回 `backend/internal/gen/proto/`
- **AND** 不得额外长出 `frontend/src/gen/connect/` 或其他并列 TS 生成根

### Requirement: 前端客户端必须以 `bufbuild/es` 产物 + 共享 Connect transport 组合承接

系统 SHALL 将前端正式客户端组合冻结为：

- `bufbuild/es` 生成的 `*_pb.ts` 作为唯一 TS 合同产物
- `@connectrpc/connect` + `@connectrpc/connect-web` 作为浏览器侧运行时 transport
- `createClient()` + `*_pb.ts` 中服务描述符作为正式客户端创建方式

系统 SHALL 在前端共享 transport 承接位统一设置 `baseUrl`，并继续维持业务切片中的 `query` 只读与 `application` 写动作边界。

系统 SHALL NOT 以新增手写 `fetch + JSON DTO`、第二个 TS 代码生成插件根或页面级临时 transport 拼装作为长期主线。

#### Scenario: Generated TS descriptors are consumed without a second TS codegen root

- **WHEN** 前端为某个 canonical 业务 service 创建 RPC client
- **THEN** client 基于 `frontend/src/gen/proto/` 下的生成产物与共享 Connect transport 创建
- **AND** 不新增第二个 TS 生成目录来承接同一份业务合同

#### Scenario: Query and application boundaries remain intact during transport migration

- **WHEN** 某个业务切片从手写 JSON adapter 切换到 Connect client
- **THEN** 只允许在切片的 query / application 固定承接位完成切换
- **AND** route、页面展示组件与 query 层不得长期直接持有第二套 transport 组织

### Requirement: 浏览器、Vite、本地启动链与 Caddy 必须共同维持单一 `/api` 外部访问面

系统 SHALL 冻结浏览器侧业务访问前缀为单一 `/api`。

系统 SHALL 冻结链路职责如下：

- 浏览器：始终向 `/api` 发起业务请求
- Vite dev proxy：开发期将 `/api` 转发到本地后端监听端口
- 本地后端启动链：只负责监听单一 HTTP 端口，并在服务内部承接 `/api`
- Caddy：生产或验收环境继续以反代方式暴露同一 `/api`

`VITE_API_BASE_URL` 若存在，只允许表达 origin 或部署入口，不得把 Connect 迁移扩写成第二个业务 base path。

#### Scenario: Development and deployed environments expose the same business base path

- **WHEN** 前端在本地开发、验收环境或未来生产环境访问业务接口
- **THEN** 外部可见业务入口始终是 `/api`
- **AND** Vite / Caddy / 本地后端启动链只是在不同环境承接这同一入口

#### Scenario: Environment configuration does not create a parallel business path

- **WHEN** 通过环境变量注入前端 API 地址
- **THEN** 环境变量只能改变 origin 或 host
- **AND** 不得把 `/api` 之外的第二条业务前缀写成正式入口

## MODIFIED Requirements

### Requirement: 既有 proto 生成链从“消息类型生成”升级为“消息类型 + Go Connect 主线生成”

当前仓库已有 `proto/Makefile + buf.gen.yaml` 单一生成入口，但历史阶段只生成 Go / TS 消息类型。

自 `phase07-02` 起，系统必须把该 requirement 修改为：

- Go 侧正式生成链同时生成 protobuf message 产物与 Connect handler / procedure 常量
- TS 侧继续保持 `bufbuild/es` 单一生成根，不为同一份合同新增第二个 TS 生成根
- `proto/README.md`、`buf.gen.yaml` 与后续实现必须对这一组合方式保持单值一致

#### Scenario: Historical message-only generation is no longer sufficient

- **WHEN** `phase07` 实施业务 transport 迁移
- **THEN** 仅生成 `*.pb.go` 与 `*_pb.ts` 不再满足正式实现要求
- **AND** Go Connect 生成产物必须进入正式工具链基线

### Requirement: 既有 `/api` JSON 路由说明不再等同于长期业务合同

历史阶段的 `proto/README.md` 与 `router.go` 记录了 RPC 到 JSON HTTP 的过渡映射。

自 `phase07-02` 起，系统必须把该 requirement 修改为：

- 这些 JSON 路由说明只视为迁移基线与 compat inventory
- 长期 canonical 业务合同解释以 `.proto + Connect procedure` 为准
- `router.go` 的业务路径声明不得继续作为正式完成态的第一真相源

#### Scenario: JSON route tables are demoted to migration reference

- **WHEN** 团队需要确认某条 canonical 业务接口的正式合同
- **THEN** 应先查看 `.proto` 与 Connect procedure
- **AND** JSON HTTP 路由表仅作为迁移过程参考或 compat 清单

## REMOVED Requirements

### Requirement: `chi + JSON HTTP` 可以继续作为长期 canonical 业务 transport 主线

**Reason**: 该假设会让 `phase07` 退化成“新增一套 Connect、旧主线继续长期保留”的并列正式状态，直接违背 `audit_001` 与 `phase07` shared baseline。

**Migration**:

- 业务合同定义迁移到 `.proto + Connect procedure`
- `chi` 回收到 mount shell、middleware 与 infra keep list
- 旧 JSON 业务 handler 只允许作为短时 compat 资产存在，并在 `phase07` 收口前退场
