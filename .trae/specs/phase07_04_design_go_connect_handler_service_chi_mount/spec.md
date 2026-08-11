# phase07-04 产出 Go Connect handler、service implementation 与 chi mount 设计 Spec

> **执行产出**：`design.md` — 包含 7 个设计区域：Connect handler 接线（5 层链 + 9 模块文件落点）、chi/Connect 横切分层（两层架构图 + 8 项职责边界表）、router 结构调整（迁移前后对比 + 3.3 职责分工）、错误映射（30+ sentinel error → 5 Connect Code 对照 + 完整 MapToConnectError 代码）、service 分层保持（4 条禁止项）、compat 过渡组（显式分组函数 + 退场时点标注）、一致性声明（9 项上游对齐）。
> **执行日期**：2026-08-11
> **参考**：Context7 `/connectrpc/connect-go` 最新文档
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

`phase07-01` 已冻结 34 条业务 RPC 迁移范围，`phase07-02` 已冻结 `chi + ConnectRPC + buf` 的正式组合方式，`phase07-03` 已冻结 compat 退场标准。当前仓库仍是 `platform/router.go` 直接装配手写 JSON handler，如果不先把 Connect handler、service implementation、middleware / interceptor 和错误链的组合方式冻结成单值规则，后续实现很容易长出第二套路由组织、第二套横切逻辑或第二套错误映射。

## What Changes

- 冻结 generated Connect handler 与既有 service implementation 的正式组合方式
- 冻结 `chi` middleware、Connect interceptor 与错误处理链的唯一承接位
- 冻结 procedure path、`/api` mount、route group 与 `platform` 装配结构的调整方案
- 冻结 domain error -> proto error code -> Connect error 的单值映射方案
- 冻结 compat JSON handler 在迁移阶段的允许位置与收口边界
- **BREAKING**：`platform/router.go` 不再允许继续以手写 `r.Get / r.Post` 业务路径表作为长期正式路由组织模式

## Impact

- Affected specs:
  - `phase07-05` 前端生成客户端与 query / application 迁移设计
  - `phase07-06` transport 实现与双栈过渡
  - `phase07-08` 后端 Connect handler 实现
  - `phase07-09` 后端 compat 路由退场
  - `phase07-11` 验收与收口核销
- Affected code:
  - `backend/internal/platform/server.go`
  - `backend/internal/platform/router.go`
  - `backend/internal/gen/connect/`
  - `backend/internal/*/service/`
  - `backend/internal/*/handler/`
  - `proto/buf.gen.yaml`
  - `proto/Makefile`

## ADDED Requirements

### Requirement: Generated Connect handler 必须通过 platform 装配层与既有 service implementation 组合

系统 SHALL 将 `platform` 装配层定义为 Go Connect handler 与既有 service implementation 的唯一正式接线位。

正式组合方式 SHALL 满足：

- `build*` 负责构造 repository / candidate / service 层依赖
- 每个 canonical 业务模块提供一个 Connect service implementation 承接位，消费既有 `service` 层，而不是在 transport 层重复实现业务逻辑
- `platform` 装配层消费 `New<Service>Handler(...)` 返回的 `(path, handler)`，把 generated handler 纳入 `/api` 下的单一业务树
- compat JSON handler 若仍临时保留，只能作为 phase07 迁移期过渡资产，不得与 Connect handler 并列为长期正式 owner

#### Scenario: Platform composes generated handler with existing service layer

- **WHEN** 团队为某个 canonical 模块接入 Connect transport
- **THEN** `platform` 装配层必须先构造既有 query / command service，再把它们接入对应的 Connect service implementation
- **AND** generated handler 必须通过统一 `/api` 业务树挂载
- **AND** 不得在 router 层重新拼装业务逻辑

#### Scenario: Transport implementation does not duplicate service logic

- **WHEN** 某个 RPC 已存在稳定的 `service` 层实现
- **THEN** Connect service implementation 只能做 transport 解包、调用 service、返回 proto 结果或 Connect error
- **AND** 不得复制第二份 command / query 业务规则

### Requirement: chi 与 Connect 的横切逻辑承接位必须单值化

系统 SHALL 将横切逻辑拆分为两层且各自职责固定：

- `chi` middleware：HTTP 壳层治理，只承接 request id、real IP、logging、recovery、timeout、CORS、顶层 auth/context 注入入口
- Connect interceptor：RPC 级治理，只承接 RPC 级 metadata 读取、请求校验、错误归一化、必要的审计/埋点扩展

系统 SHALL NOT：

- 在 `chi` middleware 与 Connect interceptor 中复制两套长期并列的 request id / logging / recovery 体系
- 把业务权限判断、业务错误映射散落回单个 handler 方法
- 让 JSON compat handler 继续保留独立错误处理主线

#### Scenario: Request lifecycle uses one HTTP shell and one RPC chain

- **WHEN** 某个 Connect RPC 请求进入后端
- **THEN** 请求先经过 `chi` 根级 middleware 链
- **AND** 再进入 Connect interceptor 链与具体 RPC 实现
- **AND** request id、日志关联与 panic recovery 的唯一主承接位保持明确

#### Scenario: Compat JSON handlers do not define a second cross-cutting stack

- **WHEN** phase07 迁移期临时保留 compat JSON handler
- **THEN** 它们只能复用同一 `chi` 外层治理基线
- **AND** 不得新增第二套长期错误包装或日志治理方式

### Requirement: router 结构调整必须保持单一 `/api` 业务树且不引入第二套路由组织模式

系统 SHALL 把 `platform/router.go` 从“手写 JSON 路径表注册器”调整为“Connect handler + infra keep list + compat 过渡入口”的组合根。

正式 router 结构 SHALL 满足：

- 根路由继续由 `server.go` 创建并挂载根级 middleware 与 infra keep list
- `/api` 继续是唯一业务前缀
- canonical 业务 service 通过 generated procedure path 挂入 `/api`
- compat JSON 入口若仍存在，必须被显式分组为迁移期过渡资产，而不是与 canonical Connect tree 混写
- 不得新增第二个业务 router 根、第二个 `/rpc` 或第二套“按模块继续手写 REST group”的正式组织模式

#### Scenario: Connect services mount under one business tree

- **WHEN** `platform` 装配 9 个 canonical 业务模块
- **THEN** 这些 service 都必须收敛到单一 `/api` 业务树下
- **AND** procedure path 与 generated path 成为业务路由的正式承接方式

#### Scenario: Compat routes are visibly transitional

- **WHEN** 某些 compat JSON 路径在 phase07-06 到 phase07-10 期间仍暂时存在
- **THEN** 它们必须在 router 结构中被显式标识为 compat / legacy group
- **AND** 不得和 canonical Connect service mount 混成同一长期组织模式

### Requirement: Domain error -> proto error code -> Connect error 必须维持单值映射

系统 SHALL 冻结后端错误链为：

`domain/service error -> proto 语义对应的错误类别 -> Connect error code`

最小设计要求：

- domain 层继续保留现有业务错误类型或 sentinel error 语义
- transport 层只在单一映射位完成 Connect error 构造
- `invalid_argument / not_found / already_exists / failed_precondition / unauthenticated / permission_denied / internal` 等长期语义必须单值映射
- 不得同时保留“旧 JSON HTTP status 映射表”和“新 Connect code 映射表”作为长期双主线

#### Scenario: Validation error maps consistently

- **WHEN** service 返回参数校验或输入非法类错误
- **THEN** transport 层必须稳定映射到对应 proto 语义与 `connect.CodeInvalidArgument`
- **AND** 前后端不得再依赖另一套长期 JSON error body 解释

#### Scenario: Not found and conflict remain distinguishable after migration

- **WHEN** service 返回 not found 或 already exists / conflict 类错误
- **THEN** transport 层必须稳定区分对应 Connect code
- **AND** 不得在迁移后把这些错误统一折叠为 `internal`

### Requirement: Service implementation 分层不得因 Connect 迁移而重写

系统 SHALL 保持既有 repository / candidate / service 分层作为业务实现主线。

Connect service implementation 的职责 SHALL 仅包括：

- 接收 proto request
- 调用既有 query / command service
- 进行结果装配
- 在固定映射位返回 Connect response / error

系统 SHALL NOT：

- 把跨模块 SQL、跨模块 candidate 构造搬进 Connect handler
- 让 generated handler 构造器直接依赖 repository 而跳过 service 层
- 因 transport 迁移引入第二套 `service` 或 `application service` 命名体系

#### Scenario: Existing service layering remains canonical

- **WHEN** 团队从当前 JSON handler 迁移到 Connect service implementation
- **THEN** canonical 业务规则仍以现有 `service` 层为准
- **AND** Connect implementation 只是新的 transport 承接位

## MODIFIED Requirements

### Requirement: phase07-02 的 Connect mount 规则补充为可直接进入后端实现的组合设计

`phase07-02` 已冻结 `chi + ConnectRPC + buf` 的正式组合方式。

自 `phase07-04` 起，系统必须把这一 requirement 修改为：

- 不只定义 `/api + generated path` 的 mount 口径，还要定义 `build* / service implementation / generated handler / chi` 的接线关系
- `chi` middleware 与 Connect interceptor 的职责边界必须写到可直接编码的粒度
- compat JSON handler 的存在位置必须被限制在明确过渡组内，而不是散落在 canonical mount 旁边

#### Scenario: Mount rule becomes implementation-ready

- **WHEN** 团队进入 `phase07-08` 后端实现
- **THEN** 必须可以直接根据 `phase07-04` 决定 handler 承接位、service implementation 组织和 router 结构
- **AND** 不再需要重新设计 transport 结构

### Requirement: phase07-03 的退场标准补充为后端传输接线前提

`phase07-03` 已冻结 compat 入口的退场窗口与双门槛收口标准。

自 `phase07-04` 起，系统必须把这一 requirement 修改为：

- Connect canonical tree 与 compat JSON group 在 router 结构中必须可区分
- 任何仍保留的 compat handler 都不得阻碍 canonical Connect service implementation 成为正式 owner
- 错误处理链必须从现在起朝单值 Connect 映射收敛，避免 phase07-09/11 再出现双主线错误语义

#### Scenario: Compat retirement is prepared by router and error design

- **WHEN** 团队为某模块先引入 Connect handler，再暂时保留 compat JSON 入口
- **THEN** router 结构和错误链必须已经体现“canonical owner 已切换、compat 仅待退场”
- **AND** 不得把 compat 入口继续设计成同等级正式主线

## REMOVED Requirements

### Requirement: 每个业务模块继续在 `router.go` 中以手写 `r.Get / r.Post` 直接注册为正式主线

**Reason**: 这会让 `.proto + ConnectRPC` 只是新增承接层，而不是正式 transport mainline，也会把 `platform/router.go` 永久保留成第二套业务路由真相源。

**Migration**:

- 将 canonical 业务 service 改为 generated Connect handler 承接
- `platform` 装配层只负责依赖接线、middleware / mount 组织和 compat 过渡分组
- 旧 JSON 路由仅在迁移窗口内短时保留，并按 `phase07-03` 退场
