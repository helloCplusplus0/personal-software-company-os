# Tasks

- [x] Task 1: 冻结 `chi` 的唯一职责与 keep list。把当前 `server.go / router.go / project_rules.md / phase07` 规划中的共识压成单值规则，明确 `chi` 只承接 HTTP 外壳、middleware、infra 端点与 `/api` mount shell。
  - [x] SubTask 1.1: 核对当前 `backend/internal/platform/server.go` 与 `router.go` 的 `/api` 装配方式、根级 middleware 与 `healthz` 承接事实 → `frozen_scope.md` §1.1 当前事实表格
  - [x] SubTask 1.2: 在规格中冻结 `healthz / readyz / metrics / debug / pprof` 为唯一允许长期留在 `chi + net/http` 的 infra keep list → `frozen_scope.md` §1.3 表格
  - [x] SubTask 1.3: 在规格中明确禁止 `chi` 继续承担 canonical 业务合同定义职责 → `frozen_scope.md` §1.2 禁止项

- [x] Task 2: 冻结 Connect handler 在单一 `/api` 前缀下的正式挂载方式。把 generated handler、procedure path 与 `chi` 外壳的组合边界写清，避免后续实现长出第二套业务基址。
  - [x] SubTask 2.1: 根据 `connect-go` 最新文档冻结"generated constructor 返回 path + http.Handler"的使用前提 → `frozen_scope.md` §2.1（基于 Context7 `/connectrpc/connect-go` 最新文档）
  - [x] SubTask 2.2: 在规格中明确 canonical 业务 RPC 对外统一承接为 `/api/{Connect procedure path}` → `frozen_scope.md` §2.3 路径映射表
  - [x] SubTask 2.3: 在规格中明确禁止新增 `/rpc`、`/connect`、`/grpc` 等并列业务前缀 → `frozen_scope.md` §2.3 禁止项

- [x] Task 3: 冻结 `buf.gen.yaml` 的正式插件矩阵与产物落点。把 `phase07` 生成链从"消息类型生成"升级为"Go protobuf + Go Connect + TS protobuf/service descriptor"的单值配置。
  - [x] SubTask 3.1: 冻结正式插件矩阵为 `protocolbuffers/go`、`connectrpc/gosimple`、`bufbuild/es` → `frozen_scope.md` §3.2 表格
  - [x] SubTask 3.2: 冻结 Go protobuf、Go Connect、TypeScript 产物的唯一目录根 → `frozen_scope.md` §3.3 目标 buf.gen.yaml + §3.4 目录树
  - [x] SubTask 3.3: 明确 `proto/Makefile` 与 `buf.gen.yaml` 继续作为唯一生成入口，不允许新增第二个 proto workspace 或第二个生成配置 → `frozen_scope.md` §3.5

- [x] Task 4: 冻结前端客户端的正式生成与运行时组合方式。保证前端只长出一套 Connect 主线，而不是再增一层手写 `fetch + JSON` 或第二个 TS 生成根。
  - [x] SubTask 4.1: 根据当前 `frontend/src/gen/proto/*_pb.ts` 与 `connect-es` 最新文档，冻结 `bufbuild/es` 产物 + `createClient()` 的组合方式 → `frozen_scope.md` §4.1-4.3（基于 Context7 `/connectrpc/connect-es` 最新文档）
  - [x] SubTask 4.2: 冻结 `@connectrpc/connect` + `@connectrpc/connect-web` 为浏览器侧共享 transport 运行时 → `frozen_scope.md` §4.4 npm 依赖表
  - [x] SubTask 4.3: 明确迁移只允许发生在前端切片的 `query / application` 固定承接位，不得让 route 或展示组件长期直接拼 transport → `frozen_scope.md` §4.5 迁移边界

- [x] Task 5: 冻结浏览器、Vite dev proxy、本地启动链与 Caddy 的 `/api` 承接关系。把不同环境如何共同维持单一外部访问面写成实现前提。
  - [x] SubTask 5.1: 基于 `frontend/vite.config.ts` 与 `.env.example` 冻结开发环境 `/api -> localhost:8081` 的代理前提 → `frozen_scope.md` §5.1 当前事实表格
  - [x] SubTask 5.2: 基于 `backend/cmd/server/main.go` 与 `backend/internal/platform/server.go` 冻结本地启动链只监听单一 HTTP 端口、在服务内承接 `/api` 的事实 → `frozen_scope.md` §5.2 架构图
  - [x] SubTask 5.3: 在规格中冻结 Caddy 继续只反代单一 `/api` 基址，不得因 Connect 迁移新增第二条公网业务前缀 → `frozen_scope.md` §5.3 冻结约束

- [x] Task 6: 完成一致性校验并形成后续子任务直接上游。确保本次 `phase07-02` 规格可以直接指导 `phase07-04 / 05 / 06 / 11` 的设计、实现与验收。
  - [x] SubTask 6.1: 校验本次规格与 `phase07` architecture plan / shared baseline 对 `chi`、ConnectRPC、`buf` 与 `/api` 的冻结结论一致 → `frozen_scope.md` §7 一致性声明
  - [x] SubTask 6.2: 校验本次规格与 `phase07-01` 的范围冻结不冲突，且不把 compat inventory 重新解释为长期正式入口 → `frozen_scope.md` §7（34 条业务 RPC 通过 §2 的挂载方式承接，不与 phase07-01 冲突）
  - [x] SubTask 6.3: 校验本次规格已把"单一生成链、单一 `/api`、单一前端 transport 主线、单一产物落点"写成明确门禁 → `frozen_scope.md` §3.5, §2.3, §4.5, §5.3 禁止项

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 2` ✅
- `Task 4` depends on `Task 3` ✅
- `Task 5` depends on `Task 1` and `Task 2` ✅
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5` ✅
