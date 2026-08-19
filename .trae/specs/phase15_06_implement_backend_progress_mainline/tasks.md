# Tasks

- [x] Task 1: proto 合同落地与生成门禁
  - [x] SubTask 1.1: 新建 `proto/psco/progress/v1/progress.proto`——phase15-04 §"psco.progress.v1 合同设计"注释版草案逐字转写（3 枚举 + `ProgressEvent` 11 字段 + 3 RPC 请求/响应 envelope + `ProgressService` 含无 Update 显式声明注释；含文件头文档定位与生成入口注释）
  - [x] SubTask 1.2: 修改 `proto/psco/project_context/v1/project_context.proto`——追加 `import "psco/progress/v1/progress.proto"`（沿既有 standard import 模式）+ 内联消息 `BriefProgress`（字段号 1-4 自然序：current_phase_key / current_phase_label / latest_task_completed / recent_events，后两者同型复用 `psco.progress.v1.ProgressEvent`）+ `GetProjectBriefResponse` 追加 `BriefProgress progress = 9`（reserved 2/3/4 不动；顶层块注释同步 5+progress 口径）
  - [x] SubTask 1.3: `proto/` 目录执行 `make gen`（三插件矩阵产出 Go pb / Go Connect / TS，progress 新包 + project_context 再生成）；执行 `make lint && make build && make breaking` 三门禁，全部零退出码（breaking 对比 main 基准：新增文件 + 追加字段为向后兼容变更，既有字段 1-8 与 reserved 零改动）

- [x] Task 2: 0013 迁移落地与开发库应用
  - [x] SubTask 2.1: 新建 `database/migrations/0013_phase15_progress_timeline.sql`——phase15-03 §"progress_events DDL 级设计"SQL 草案逐字转写（11 列 + `workflow_type`/`event_kind`/`source` 三列 `TEXT + CHECK(IN ...)` + FK `repositories(id) ON DELETE RESTRICT` + `idx_progress_events_repository_sort (repository_id, occurred_at DESC, created_at DESC)` + 文件头注释；全段 `IF NOT EXISTS` 幂等；无 DO 块无 DROP 段）
  - [x] SubTask 2.2: 以 `backend/.env` 的 `DATABASE_URL` 经 `psql -f` 应用 0013 到开发库一次（用户已启动的服务器早于本文件存在，RunMigrations 不会自动补执行；DDL 幂等保证服务器下次启动重放空过 + schema_migrations 自动补登记）；验证 `progress_events` 表与索引存在（`\d progress_events`）

- [x] Task 3: progress 根包实现与单元测试
  - [x] SubTask 3.1: 新建 `backend/internal/progress/types.go`——phase15-04 types.go 草案逐字转写：受控枚举（`WorkflowType`/`EventKind`/`ProgressSource`，string 形态对齐 DDL CHECK 值）+ `ProgressEventReadResult`（11 字段）+ `ProgressSummary`（四字段；`LatestTaskCompleted` 指针可空、`RecentEvents` 恒非 nil）+ `CreateProgressEventInput`；包注释沿 standard 分层语义；不 import 生成 pb 包
  - [x] SubTask 3.2: 新建 `backend/internal/progress/errors.go`——phase15-04 哨兵清单逐字转写（5 哨兵：`ErrProgressEventNotFound`/`ErrRepositoryNotFound` → NotFound；`ErrInvalidInput` → InvalidArgument；`ErrProgressReadFailed`/`ErrProgressWriteFailed` → Internal）
  - [x] SubTask 3.3: 新建 `backend/internal/progress/validate.go`——phase15-03 错误码总表与执行序 6 步逐字实现：envelope 前置（`INVALID_REPOSITORY_ID` UUID 格式 / `INVALID_OCCURRED_AT` 已设置）→ V1a `INVALID_WORKFLOW_TYPE` → V1b `INVALID_EVENT_KIND` → V7 `EVENT_KIND_NOT_ALLOWED`（audit/fix × phase_started/phase_completed）→ task_key 矩阵分支（必填 `TASK_KEY_REQUIRED` 先于格式 `TASK_KEY_FORMAT_INVALID`；K-1 `^phase[0-9]{2,}$` / K-2 `^phase[0-9]{2,}-[0-9]{2,}$` / K-3 `^audit_[0-9]{3,}$` / K-4 `^fix_[0-9]{3,}$`；note 格 K-5 可空不校验格式）→ V9 文本顺序（`INVALID_TITLE`〔trim 非空 + rune ≤200〕→ `INVALID_DETAIL`〔≤2000〕→ `INVALID_EVIDENCE_REF`〔空或 `/`|`https://` 前缀〕→ `INVALID_SOURCE`〔仅 manual〕）；报第一个错误；`%w: [CODE] message` 包装格式；错误信息携带期望格式说明（如 `expected ^phase[0-9]{2,}$`）
  - [x] SubTask 3.4: 新建 `backend/internal/progress/derive.go`——`DeriveProgressSummary` 纯函数（零 I/O 零时间函数）：输入三键链 DESC 全序切片（索引小 = 事件晚），输出四字段——recent_events 取 `events[0:min(10,len)]`、latest_task_completed 自 i=0 找首个 `task_completed`（无 → nil）、current_phase 三态（无 `phase_started` → 空值；`latestStartedIdx` 前存在同 key `phase_completed` → 空值；否则 key=`task_key` label=`title`）；空集 → `RecentEvents` 空切片非 nil
  - [x] SubTask 3.5: 新建 `backend/internal/progress/validate_test.go`——表驱动：合法矩阵 12 格正例 + 11 错误码逐码反例（`EVENT_KIND_NOT_ALLOWED` 至少 audit×phase_started 与 fix×phase_completed 两用例；`TASK_KEY_FORMAT_INVALID` K-1~K-4 每正则至少 1 不符值；`INVALID_TITLE` 空与超 200 rune 两用例）+ 执行序首错断言（多错输入只报第一个）；按错误码断言非信息全文
  - [x] SubTask 3.6: 新建 `backend/internal/progress/derive_test.go`——纯内存边界全覆盖：空集 / 仅 note / 从未开始 / 进行中 / 全部完结（最新 phase_started 后同 key phase_completed）/ 补录同刻 tiebreak（同 occurred_at+created_at 下 id DESC 序）/ recent 截断（>10 条取前 10）

- [x] Task 4: progress 四子包实现
  - [x] SubTask 4.1: 新建 `backend/internal/progress/repository/progress_event_store.go`——`ProgressEventStore`：`ListByRepository(ctx, repositoryID, workflowType *WorkflowType)` 单一查询（phase15-03 冻结 SQL 形态：repository_id 必选 + workflow_type 可选过滤 + `ORDER BY occurred_at DESC, created_at DESC, id DESC`，Go 侧不重排；空结果返回空切片非 nil）+ `Insert(ctx, input)`（`RETURNING` 完整行；`task_key`/`detail`/`evidence_ref` 可空 TEXT NULL↔空串解码在 scan 层）+ `DeleteByID(ctx, id)`（`RowsAffected()==0` → `ErrProgressEventNotFound`）；读 wrap `ErrProgressReadFailed`、写 wrap `ErrProgressWriteFailed`
  - [x] SubTask 4.2: 新建 `backend/internal/progress/candidate/repository_reader.go`——`RepositoryReader`（DP-2 承接位）：`RepositoryExists(ctx, repositoryID) (bool, error)`，`SELECT EXISTS(SELECT 1 FROM repositories WHERE id = $1)`；纯存在性事实查询，失败包装原始错误（Internal 兜底），不承载业务错误语义
  - [x] SubTask 4.3: 新建 `backend/internal/progress/service/query_service.go`——`QueryService`（依赖 store + repositoryReader）：`ListProgressEvents`（UUID 校验 → `RepositoryExists` `!found` → `ErrRepositoryNotFound`〔NotFound 读锚点语义〕→ store 查询，`workflow_type` UNSPECIFIED → nil 不过滤）+ `GetProgressSummary`（store 单查询无过滤 → `DeriveProgressSummary`；ProgressReader 实现位；不做存在性校验，无行 → 零值摘要 + 空 RecentEvents）
  - [x] SubTask 4.4: 新建 `backend/internal/progress/service/command_service.go`——`CommandService`（依赖 store + repositoryReader）：`CreateProgressEvent`（validate 执行序 6 步全量 → `RepositoryExists` `!found` → `[REPOSITORY_NOT_FOUND] repository %s does not exist` wrap `ErrInvalidInput`〔InvalidArgument 写入目标引用语义〕→ store Insert）+ `DeleteProgressEvent`（UUID 校验 `[INVALID_PROGRESS_EVENT_ID]` → store Delete，`ErrProgressEventNotFound` 透传）
  - [x] SubTask 4.5: 新建 `backend/internal/progress/connect/server.go`——`Server`（实现 `ProgressServiceHandler`，`NewServer(querySvc, commandSvc)`）：proto 解包（枚举 UNSPECIFIED/越界 → 对应 `[CODE]` InvalidArgument；`source` nil → 归一 manual，显式设置非 MANUAL → `INVALID_SOURCE`；`occurred_at` nil → `INVALID_OCCURRED_AT`）→ service 调用 → proto 组装；**导出 `DomainProgressEventToProto`**（沿 `DomainStandardToProto` 导出模式，供 projectcontext 组装 BriefProgress 复用）；内部枚举转换函数小写私有；错误统一 `connecterrors.MapToConnectError`

- [x] Task 5: projectcontext ProgressReader 接入与 brief 装配
  - [x] SubTask 5.1: 修改 `backend/internal/projectcontext/candidate/context_readers.go`——与 `StandardReader` 同文件追加 `ProgressReader` 接口（phase15-04 接口定义逐字转写：`GetProgressSummary(ctx, repositoryID string) (progress.ProgressSummary, error)`；失败 → `progress.ErrProgressReadFailed`；无事件 → 零值摘要非错误；文件头注释同步 progress 读取条目）
  - [x] SubTask 5.2: 修改 `backend/internal/projectcontext/types.go`——`ProjectBriefReadResult` 追加 `Progress progress.ProgressSummary` 值类型字段（空态恒构造零值；沿 `Standards []standard.StandardReadResult` domain 透传模式；import progress 包）
  - [x] SubTask 5.3: 修改 `backend/internal/projectcontext/service/query_service.go`——`QueryService` 依赖新增 `progressReader candidate.ProgressReader`（`NewQueryService(contextReaders, standardReader, progressReader)` 签名演进）；`GetProjectBrief` 编排新增步骤 6（standards 之后、return 之前）：`GetProgressSummary` 失败直接透传（wrapped `ErrProgressReadFailed` → CodeInternal connecterrors 统一映射）；编排注释同步 6 步
  - [x] SubTask 5.4: 修改 `backend/internal/projectcontext/connect/server.go`——`GetProjectBrief` 组装 `BriefProgress` 块：`CurrentPhaseKey`/`CurrentPhaseLabel` 直映射 + `LatestTaskCompleted` nil 不设置 + `RecentEvents` 经 `DomainProgressEventToProto` 逐元素转换（空态组装为空数组非 nil）；import progresspb 与 progressconnect

- [x] Task 6: connecterrors 登记与 platform 装配
  - [x] SubTask 6.1: 修改 `backend/internal/connecterrors/connect_errors.go`——CodeNotFound 组追加 `progress.ErrProgressEventNotFound`、`progress.ErrRepositoryNotFound`；CodeInvalidArgument 组追加 `progress.ErrInvalidInput`；CodeInternal 兜底组追加显式引用 `_ = progress.ErrProgressReadFailed`、`_ = progress.ErrProgressWriteFailed`；import progress 包
  - [x] SubTask 6.2: 修改 `backend/internal/platform/router.go`——新增 `buildProgress(pool)`（`candidate.NewRepositoryReader` → `repository.NewProgressEventStore` → 双 service，沿 `buildStandard` 模式）+ `mountProgressConnect(r, querySvc, commandSvc)`（`progressv1connect.NewProgressServiceHandler` + `r.Handle(path+"*", http.StripPrefix("/api", handler))`）+ `buildProjectContext` 签名演进（新增 `progressReader projectcontextcandidate.ProgressReader` 参数并注入 `NewQueryService`）；import progress 各子包与生成 connect 包
  - [x] SubTask 6.3: 修改 `backend/internal/platform/server.go`——装配接线：`buildProgress` + `mountProgressConnect` 置于 standard 块之后、projectcontext 块之前；`buildProjectContext(pool, standardQuerySvc, progressQuerySvc)` 三参调用（progress QueryService 作为 ProgressReader 实现注入）；装配注释同步 phase15-06 条目

- [x] Task 7: 集成测试与全量门禁
  - [x] SubTask 7.1: 新建 `backend/internal/progress/service/service_integration_test.go`——沿 standard 集成测试模式（真实 PostgreSQL：DATABASE_URL 优先回落 `backend/.env`，无 DB 时 `t.Skipf`；独立 fixture〔uuid 后缀 repositories 行 + progress_events〕+ `t.Cleanup` 显式清理）：Create→List round-trip（多事件三键链倒序断言，含补录历史 occurred_at 排序）+ 三轨过滤（workflow_type 各轨断言 + UNSPECIFIED 全量）+ Delete 后 List 不含 + Delete 不存在 → NotFound + Create repository 不存在 → InvalidArgument 错误信息含 `REPOSITORY_NOT_FOUND` + 校验失败码抽查（`TASK_KEY_FORMAT_INVALID` / `EVENT_KIND_NOT_ALLOWED`）
  - [x] SubTask 7.2: 修改 `backend/internal/projectcontext/connect/server_integration_test.go`——追加 brief progress 场景（沿既有 harness 与 fixture 模式，独立 repository fixture + cleanup，不破坏既有场景）：①round-trip——Create phase_started（task_key=phaseNN）+ task_completed 后 GetProjectBrief，断言 `progress.current_phase_key` = 录入 key、`current_phase_label` = 录入 title、`recent_events` 含录入事件、`latest_task_completed` 非空；②派生断言（DoD 冻结）——Create phase_completed（同 task_key）后 GetProjectBrief，断言 `progress.current_phase_key` 为空串（"phase_completed 后当前 phase 为空"）；③空态恒构造——0 事件仓库 GetProjectBrief，断言 `progress` 块非 nil、`recent_events` 空数组、`latest_task_completed` 未设置
  - [x] SubTask 7.3: 全量门禁执行——`backend/` 目录 `go build ./...` && `go vet ./...` && `go test ./...` 全部零退出码（集成测试连开发库执行非 skip）；`proto/` 目录 `make lint && make build && make breaking` 复验零退出码；确认零前端源码改动（`frontend/src/` 仅 `gen/` 生成物再生成，无手写文件改动）

- [x] Task 8: 一致性校验、独立复核与收口
  - [x] SubTask 8.1: 一致性校验——proto 与 phase15-04 草案逐字比对（枚举/字段号/RPC/注释）；0013 与 phase15-03 DDL 逐字比对（列/约束/索引/注释）；validate/derive 与 phase15-03 错误码总表和三算法逐条比对；模块结构与 standard 模式逐文件对照；确认 spec §What Changes 文件清单与 git 工作区实际改动一一对应（无清单外文件）
  - [x] SubTask 8.2: 子代理独立复核（实现与上游零漂移 / 测试覆盖 DoD 全项含非法用例矩阵与 phase_completed 派生断言 / 门禁七命令实测全绿 / 无 phase15-07/08 偷渡〔零前端手写代码、零 dogfooding、零根级文档〕/ 服务器未被重启）
  - [x] SubTask 8.3: 修复独立复核发现的阻断性问题（如有）并复验；tasks.md / checklist.md 全部勾选附执行记录；变更保持未提交，待用户最终确认后手动提交

# Task Dependencies

- Task 1（proto）与 Task 2（迁移）相互独立，可并行
- Task 3 depends on Task 1（根包不依赖 pb，但 validate/derive 的枚举值与 proto 值域同源冻结，先落合同保证一致性）
- Task 4 depends on Task 3（service 消费 validate/derive/types；connect 消费生成 pb——Task 1 已产出）
- Task 5 depends on Task 4（projectcontext 消费 progress 包与 DomainProgressEventToProto）
- Task 6 depends on Task 4 + Task 5（装配接线消费 progress 四子包与演进后的 buildProjectContext 签名）
- Task 7 depends on Task 1 ~ Task 6（集成测试与门禁为全量收口）+ Task 2（集成测试需表就绪）
- Task 8 depends on Task 7
- 后续：phase15-07 depends on 本 spec（合同生成物 + 3 RPC 运行时）+ phase15-05（切片设计）；phase15-08 depends on 本 spec + phase15-07

# 执行记录（2026-08-19）

- Task 1（并行批次 1）：progress.proto 逐字落地（3 枚举 + 11 字段 + 3 RPC + 无 Update 显式声明）；project_context.proto 四处演进（import + BriefProgress 1-4 + progress=9 + 注释）；`make gen / lint / build / breaking` 四门禁退出码全 0；三端生成物齐备（Go pb / Go Connect progressv1connect / TS + project_context 再生成）。
- Task 2（并行批次 1）：0013 逐字落地（11 列 + 三 CHECK + FK RESTRICT + 三键索引 + IF NOT EXISTS）；本机无 psql，经 docker 容器等效 psql 应用到 `psco_development`（与 backend/.env DATABASE_URL 同库）；幂等重放空过（NOTICE skipping）；`\d progress_events` 验证 11 列 / 3 CHECK / FK / 索引齐备；零 Go 代码改动（RunMigrations 自动登记机制）。
- Task 3：根包 4 文件 + 2 测试文件落地；`go build / vet / test ./internal/progress/...` 全绿（71 用例：validate 57 + derive 13 + trim 专项 1）。标注：spec"合法矩阵 12 格正例"与 phase15-02 语义矩阵冲突（audit/fix × phase_started/phase_completed 为 V7 禁止格），按语义上游实现为 8 格正例 + 4 禁止格反例（12 格全判定）——独立复核裁决该处理正确（笔误级冲突，spec 单值一致声明规定以上游为准）。
- Task 4：四子包 5 文件落地（repository/candidate/service×2/connect），与 standard 同层逐文件对照一致；导出 `DomainProgressEventToProto`；3 RPC 错误分支与 phase15-04 三要素表逐项对应；根包零改动、无签名适配。
- Task 5+6（合并批次，签名演进耦合）：projectcontext 4 文件接入（ProgressReader 同文件追加 / Progress 值字段 / 编排步骤 6 / BriefProgress 组装）+ connecterrors 5 哨兵登记 + platform router/server 装配（buildProgress/mountProgressConnect 置 standard 后 projectcontext 前）；额外最小适配：既有 server_integration_test.go 调用点随 NewQueryService 三参签名演进注入真实 progress service（该文件本就在 spec 修改清单内，独立复核裁决在允许范围内）；全仓 `go build / vet` 零错、既有集成测试不回归。
- Task 7：service_integration_test.go 5 集成用例（round-trip 含补录与同刻 tiebreak / 三轨过滤 / Delete 语义 / Create 错误分支 / List 读锚点 NotFound）+ brief 3 场景（round-trip / phase_completed 后 current_phase_key 空串〔DoD 冻结〕/ 空态恒构造）；七条门禁复验退出码全 0（go build/vet/test ./... -count=1 + make lint/build/breaking），集成测试真实连库执行非 skip。过程修复：一次 Edit 误删既有 helper 当场恢复并经 git diff 核实零残留；空态 recent_events 客户端断言按 proto wire 语义修正为 len==0（服务端组装侧恒构造非 nil 语义保持，独立复核裁决与 spec 一致）。
- Task 8：独立复核六维度全 PASS、0 阻断项：①实现与 phase15-02/03/04 逐字零漂移；②模块结构与 standard 逐文件对照无第二套；③测试覆盖 DoD 全项；④门禁六命令亲自复验退出码全 0（含 -v 确认集成测试真实执行）；⑤无偷渡（git 工作区与 spec 清单一一对应、frontend 零改动、无 UpdateProgressEvent 实现、无 phase11 改动、无根级回写、progress_events 0 行、服务器进程自 11:26 持续运行未重启）；⑥三项实现标注全部裁决正确。非阻断观察 3 项留档：spec"12 格正例"表述建议 phase15-09 澄清为"12 格全判定（8 正例 + 4 禁止格反例）"；spec"后端 :8080"与用户环境实际 8081 不符（环境观察，佐证无重复开服）；schema_migrations 尚未登记 0013 属 spec 预期内（服务器下次重启 RunMigrations 空过补登记，phase15-07/08 前提需确认）。
- 收口状态：全部变更未提交，待用户最终确认后手动提交。
