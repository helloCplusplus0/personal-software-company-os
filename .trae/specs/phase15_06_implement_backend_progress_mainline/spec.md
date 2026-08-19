# phase15-06 落实后端主线 Spec

## Why

`phase15-03`（DDL / 9+2 错误码 / 派生算法实现序）与 `phase15-04`（proto 合同 / 3 RPC 错误语义 / Go 模块分层 / ProgressReader 接口 / brief `progress = 9` 装配演进点）已将全部实现设计冻结到"逐字转写"粒度，本子任务将其落地为可运行、可测试的后端主线代码：`psco.progress.v1` 合同 + `0013` 迁移 + `backend/internal/progress/` 模块 + `projectcontext` brief 装配 + `connecterrors` / `platform` 接线。`phase15-07`（前端实现）与 `phase15-08`（联调验收）以本子任务的运行时产出为前提。

`phase15-06` 是源代码实现类子任务：全部设计决策已在 phase15-03 / 04 冻结，本 spec 不新增任何设计决策，只承载"逐字转写关系 + 实现顺序 + 测试覆盖 + 门禁验收"。

## What Changes

- 新建 `proto/psco/progress/v1/progress.proto`（phase15-04 §"psco.progress.v1 合同设计"草案逐字转写：3 枚举 + `ProgressEvent` 11 字段 + 3 RPC envelope + `ProgressService` 含无 Update 显式声明）
- 修改 `proto/psco/project_context/v1/project_context.proto`：追加 `import "psco/progress/v1/progress.proto"` + 内联消息 `BriefProgress`（字段号 1-4）+ `GetProjectBriefResponse.progress = 9`（槽位 2/3/4 reserved 不动）
- `buf generate` 产出三端生成物（Go pb / Go Connect / TS；生成目录已 gitignore，`make gen` 重新生成）
- 新建 `database/migrations/0013_phase15_progress_timeline.sql`（phase15-03 DDL 草案逐字转写：单表 + 三列 `TEXT + CHECK(IN ...)` + `idx_progress_events_repository_sort` + 幂等 `IF NOT EXISTS`；落入 `database/migrations/` 即被 `RunMigrations` 按文件名升序自动登记——phase14-07 OBS-01 修复后机制，零手工登记改动）+ 应用到开发库（`backend/.env` DATABASE_URL）
- 新建 `backend/internal/progress/` 模块：根包 `types.go` / `errors.go` / `validate.go` / `derive.go` + 四子包 `connect/` / `service/` / `repository/` / `candidate/`（文件落点与类型签名沿 phase15-04 §"Go 模块分层与文件落点"逐字实现，结构对照 `standard` 模块）
- 修改 `backend/internal/projectcontext/`：`candidate/context_readers.go` 追加 `ProgressReader` 接口（与 `StandardReader` 同文件）；`types.go` 的 `ProjectBriefReadResult` 追加 `Progress progress.ProgressSummary` 值类型字段；`service/query_service.go` 的 `GetProjectBrief` 编排新增步骤 6；`connect/server.go` 组装 `BriefProgress` 块（复用 `progress/connect` 导出的 `DomainProgressEventToProto`）
- 修改 `backend/internal/connecterrors/connect_errors.go`：登记 progress 5 个哨兵错误（CodeNotFound 组 2 个 + CodeInvalidArgument 组 1 个 + CodeInternal 兜底显式引用 2 个）
- 修改 `backend/internal/platform/router.go`：新增 `buildProgress` / `mountProgressConnect`；`buildProjectContext` 签名演进（新增 `progressReader` 参数）
- 修改 `backend/internal/platform/server.go`：装配接线（progress 构造先于 `buildProjectContext`，QueryService 作为 `ProgressReader` 注入）
- 新建测试：`backend/internal/progress/validate_test.go`（校验单元测试，含非法用例矩阵）、`backend/internal/progress/derive_test.go`（派生纯函数单元测试）、`backend/internal/progress/service/service_integration_test.go`（3 RPC 集成）；修改 `backend/internal/projectcontext/connect/server_integration_test.go`（brief progress round-trip 与派生断言，含 `phase_completed` 后当前 phase 为空）

## Impact

- Affected specs:
  - `phase15-03` / `phase15-04`（直接设计上游，本 spec 为其实现落地；全部内容逐字转写，零再决策）
  - `phase15-02`（语义上游：合法矩阵 12 格 / K-1~K-5 正则 / 派生规则 / brief `progress = 9`——经 03/04 间接承接）
  - `phase15-07` / `phase15-08`（下游：前端实现依赖本子任务的合同生成物与运行时 RPC；联调验收依赖 3 RPC 可用）
- Affected code（全部实现落点）:
  - 新建：`proto/psco/progress/v1/progress.proto`、`database/migrations/0013_phase15_progress_timeline.sql`、`backend/internal/progress/`（根包 4 文件 + 4 子包 + 3 测试文件）
  - 修改：`proto/psco/project_context/v1/project_context.proto`、`backend/internal/projectcontext/`（candidate / types / service / connect 4 文件）、`backend/internal/connecterrors/connect_errors.go`、`backend/internal/platform/router.go`、`backend/internal/platform/server.go`、`backend/internal/projectcontext/connect/server_integration_test.go`
  - 生成物（gitignore，不入版本控制）：`backend/internal/gen/proto/psco/progress/v1/`、`backend/internal/gen/connect/psco/progress/v1/`、`frontend/src/gen/proto/psco/progress/v1/` + project_context 同三端再生成
- 运行时约束：用户已手动开启前后端服务器（后端 :8080），本子任务**禁止**重启或在其他端口重复开启服务器；后端运行时验证（浏览器完整会话）归 `phase15-07/08`，本子任务门禁为工具链与测试层

## ADDED Requirements

### Requirement: proto 合同落地必须逐字转写且门禁全绿

本子任务 SHALL 将 phase15-04 §"psco.progress.v1 合同设计"的注释版草案逐字转写为 `proto/psco/progress/v1/progress.proto`（枚举 / 消息 / 字段号 / RPC 定义 / 注释零再决策），并按其 §"brief progress = 9 装配演进点必须单值冻结"修改 `project_context.proto`（import + `BriefProgress` 字段号 1-4 + `progress = 9`，槽位 2/3/4 reserved 不动）：

- 生成入口：`proto/` 目录执行 `make gen`（`buf generate`，三插件矩阵：protoc-gen-go / connectrpc-gosimple / bufbuild-es）
- 门禁（DoD）：`make lint`（STANDARD 规则集，SERVICE_SUFFIX 豁免既有）、`make build`、`make breaking`（对比 `../.git#branch=main,subdir=proto`；新增文件 + 追加字段 9 均为向后兼容变更，对既有字段 1-8 零破坏）全部通过

#### Scenario: 合同可编译且零破坏

- **WHEN** proto 两文件落地并执行 `make lint && make build && make breaking`
- **THEN** 三命令退出码均为 0；`buf generate` 产出 Go pb（`backend/internal/gen/proto/psco/progress/v1/`）、Go Connect handler 构造器（`backend/internal/gen/connect/psco/progress/v1/`）、TS（`frontend/src/gen/proto/psco/progress/v1/`），且 project_context 三端生成物同步再生成
- **AND** `GetProjectBriefResponse` 既有字段（1、5-8）与 reserved（2/3/4）零改动

### Requirement: 0013 迁移落地与开发库应用

本子任务 SHALL 将 phase15-03 §"progress_events DDL 级设计必须可直接进入迁移实现"的 SQL 草案逐字转写为 `database/migrations/0013_phase15_progress_timeline.sql`（11 列 + 三列 `TEXT + CHECK(IN ...)` + FK `ON DELETE RESTRICT` + `idx_progress_events_repository_sort` + 文件头注释，零再决策）：

- 自动登记：落入 `database/migrations/` 即被 `RunMigrations` 按文件名升序自动登记执行，**零手工登记代码改动**（phase14-07 OBS-01 修复后机制）
- 开发库应用：用户已启动的服务器启动早于本迁移文件存在，`RunMigrations` 不会自动补执行——本子任务以 `psql`（`backend/.env` 的 `DATABASE_URL`）直接执行 `0013` 一次；DDL 全段 `IF NOT EXISTS` 幂等，服务器下次重启时 `RunMigrations` 重放该文件为空过并补登记 `schema_migrations`，无双跑风险
- 集成测试前置：`progress` 与 `projectcontext` 集成测试连接同一开发库（DATABASE_URL 优先，回落 `backend/.env`），表就绪为测试运行前提

#### Scenario: 迁移可重放且表就绪

- **WHEN** `0013` 文件落地并经 `psql` 应用到开发库
- **THEN** `progress_events` 表与 `idx_progress_events_repository_sort` 索引存在；重复执行整文件为空过（幂等）；`schema_migrations` 由下次服务器启动自动补登记，不产生版本漂移

### Requirement: progress 后端模块实现必须沿 standard 模式逐字落地

本子任务 SHALL 实现 `backend/internal/progress/`（`package progress`），文件落点 / 类型签名 / 枚举值 / 错误语义按 phase15-04 冻结逐字实现，与 `standard` 模块逐文件对照结构一致：

**根包 4 文件**：

| 文件 | 承接内容（逐字源） |
|---|---|
| `types.go` | 受控枚举（`WorkflowType` / `EventKind` / `ProgressSource` string 形态对齐 DDL CHECK）+ `ProgressEventReadResult`（11 字段）+ `ProgressSummary`（四字段，`RecentEvents` 恒非 nil）+ `CreateProgressEventInput`——phase15-04 types.go 草案逐字转写；不 import 生成 pb 包 |
| `errors.go` | 5 个哨兵错误（`ErrProgressEventNotFound` / `ErrRepositoryNotFound` / `ErrInvalidInput` / `ErrProgressReadFailed` / `ErrProgressWriteFailed`）——phase15-04 哨兵清单逐字转写 |
| `validate.go` | V1-V9 + envelope 前置校验（11 错误码：9 业务码 + `INVALID_REPOSITORY_ID` / `INVALID_OCCURRED_AT`）+ 执行序 6 步报第一个错误 + K-1~K-4 正则（`^phase[0-9]{2,}$` / `^phase[0-9]{2,}-[0-9]{2,}$` / `^audit_[0-9]{3,}$` / `^fix_[0-9]{3,}$`）+ TrimSpace 边界（`task_key` trim 后持久化；`title` trim 判定原值入库；rune 计数）——phase15-03 错误码总表与执行序逐字实现，`%w: [CODE] message` 包装格式 |
| `derive.go` | `DeriveProgressSummary(events []ProgressEventReadResult) ProgressSummary` 纯函数（零 I/O / 零时间函数）：recent_events N=10 切片 + latest_task_completed 首个 `task_completed` + current_phase 三态判定（无 `phase_started` → 空值；最新 `phase_started` 后存在同 key `phase_completed` → 空值；否则 key+label）——phase15-03 三派生算法逐字实现 |

**四子包**：

| 子包 | 承接内容 |
|---|---|
| `repository/` | `ProgressEventStore`：单一查询（phase15-03 冻结 SQL 形态：`repository_id = $1` 必选 + `workflow_type = $2` 可选 + `ORDER BY occurred_at DESC, created_at DESC, id DESC`，Go 侧不重排）+ INSERT（`RETURNING` 完整行；NULL → 空串解码）+ DELETE BY id（`RowsAffected()==0` → `ErrProgressEventNotFound`）；读失败 wrap `ErrProgressReadFailed`、写失败 wrap `ErrProgressWriteFailed` |
| `candidate/` | `RepositoryReader`（DP-2 承接位）：`RepositoryExists(ctx, repositoryID) (bool, error)`——`SELECT EXISTS(SELECT 1 FROM repositories WHERE id = $1)`，纯存在性事实查询；查询失败包装原始错误（Internal 兜底） |
| `service/` | `QueryService`（依赖 store + repositoryReader）：`ListProgressEvents`（UUID 校验 → 存在性 `!found` → `ErrRepositoryNotFound`〔NotFound 读锚点语义〕→ store 查询可选过滤）+ `GetProgressSummary`（store 单查询无过滤 → `DeriveProgressSummary`；不做 repository 存在性校验，无行 → 零值摘要）；`CommandService`（依赖 store + repositoryReader）：`CreateProgressEvent`（执行序 6 步：envelope 前置 → V1a → V1b → V7 → task_key 矩阵分支 → V9 文本 → repository 存在性 `!found` → `[REPOSITORY_NOT_FOUND]` wrap `ErrInvalidInput`；source 归一 manual）+ `DeleteProgressEvent`（UUID 校验 `[INVALID_PROGRESS_EVENT_ID]` → store 删除） |
| `connect/` | `Server`（实现 `ProgressServiceHandler`）：proto 解包（枚举 UNSPECIFIED → 对应 `[CODE]` InvalidArgument；`source` nil → 归一 manual、显式非 MANUAL → `INVALID_SOURCE`）→ service 调用 → proto 组装；**导出 `DomainProgressEventToProto(e progress.ProgressEventReadResult) *progressv1.ProgressEvent`**（沿 `DomainStandardToProto` 导出模式，供 projectcontext 组装 BriefProgress 复用）；错误统一 `connecterrors.MapToConnectError` |

#### Scenario: 模块结构与 standard 对照一致

- **WHEN** 模块实现完成
- **THEN** 根包 4 文件 + 四子包与 `backend/internal/standard/` 逐文件对照结构一致（types/errors/validate〔+derive〕→ connect/service/repository/candidate）；无第五子包、无第二套转换函数、无 service 直写跨模块 SQL
- **AND** 3 RPC 错误分支按 phase15-04 §"3 RPC 错误语义逐个三要素冻结"表格逐项实现（List 不存在 → NotFound / Create 不存在 → InvalidArgument 含 `REPOSITORY_NOT_FOUND` / Delete 不存在 → NotFound / List 过滤空 → 空列表非错误）

### Requirement: projectcontext ProgressReader 接入与 brief progress = 9 装配

本子任务 SHALL 按 phase15-04 §"ProgressReader 跨模块读取接口签名单值冻结"与 §"brief progress = 9 装配演进点必须单值冻结"实现 projectcontext 侧接入：

1. `candidate/context_readers.go` 追加 `ProgressReader` 接口（与 `StandardReader` 同文件，接口注释逐字转写：`GetProgressSummary(ctx, repositoryID) (progress.ProgressSummary, error)`；失败 → `progress.ErrProgressReadFailed`；无事件 → 零值摘要非错误）
2. `types.go`：`ProjectBriefReadResult` 追加 `Progress progress.ProgressSummary` 值类型字段（空态恒构造零值，沿 `Standards` domain 类型直接透传模式）
3. `service/query_service.go`：`QueryService` 依赖新增 `progressReader candidate.ProgressReader`（`NewQueryService` 签名演进）；`GetProjectBrief` 编排新增步骤 6（standards 之后）：`GetProgressSummary` 失败透传（`ErrProgressReadFailed` → CodeInternal）；空态恒构造零值
4. `connect/server.go`：组装 `BriefProgress`——`CurrentPhaseKey` / `CurrentPhaseLabel` 直映射；`LatestTaskCompleted` nil 不设置；`RecentEvents` 经 `DomainProgressEventToProto` 逐元素转换，空态组装为空数组（非 nil）

#### Scenario: brief 装配零第二套派生

- **WHEN** GetProjectBrief 实现
- **THEN** progress 读取唯一经 `ProgressReader` 接口（progress/service.QueryService 实现，platform 装配点注入）；projectcontext 不出现 `progress_events` 表 SQL 或派生逻辑复制
- **AND** 0 事件仓库的 brief 响应 `progress` 块非 nil（恒构造）：`current_phase_key` 空串、`latest_task_completed` 未设置、`recent_events` 空数组

### Requirement: connecterrors 登记与 platform 装配接线

本子任务 SHALL 完成错误映射登记与组合根装配：

- `connecterrors/connect_errors.go`：CodeNotFound 组追加 `progress.ErrProgressEventNotFound`、`progress.ErrRepositoryNotFound`；CodeInvalidArgument 组追加 `progress.ErrInvalidInput`；CodeInternal 兜底组追加显式引用 `_ = progress.ErrProgressReadFailed`、`_ = progress.ErrProgressWriteFailed`（沿 standard 登记模式）
- `platform/router.go`：新增 `buildProgress(pool)`（构造 `candidate.RepositoryReader` → `repository.ProgressEventStore` → 双 service，沿 `buildStandard` 模式）与 `mountProgressConnect(r, querySvc, commandSvc)`（`progressv1connect.NewProgressServiceHandler` + `r.Handle(path+"*", http.StripPrefix("/api", handler))`，沿 `mountStandardConnect` 模式）；`buildProjectContext` 签名演进：新增 `progressReader projectcontextcandidate.ProgressReader` 参数
- `platform/server.go`：装配顺序——`buildProgress` + `mountProgressConnect` 置于 standard 块之后、`buildProjectContext` 之前；`buildProjectContext(pool, standardQuerySvc, progressQuerySvc)`（progress QueryService 作为 `ProgressReader` 实现注入）

#### Scenario: 装配后 RPC 可达

- **WHEN** 后端编译启动（用户环境由既有服务器承载，本子任务以编译 + 测试验证）
- **THEN** `/api/psco.progress.v1.ProgressService/` 三 RPC 经 Connect handler 可路由；brief 的 standards[] 与 progress 两 reader 注入互不影响

### Requirement: 测试覆盖必须含非法用例矩阵与派生断言

本子任务 SHALL 新增三层测试（DoD 冻结）：

1. **校验单元测试**（`backend/internal/progress/validate_test.go`，表驱动，无需 DB）：合法矩阵 12 格正例（phase 轨 4 格 + audit/fix 轨各 4 格，task_key 按矩阵给合法值）+ 逐错误码反例（11 码每码至少 1 个触发用例：`INVALID_WORKFLOW_TYPE` / `INVALID_EVENT_KIND` / `EVENT_KIND_NOT_ALLOWED`〔audit×phase_started 与 fix×phase_completed 至少各 1〕/ `TASK_KEY_REQUIRED`〔必填格空串与纯空白〕/ `TASK_KEY_FORMAT_INVALID`〔K-1~K-4 每正则至少 1 个不符值〕/ `INVALID_TITLE`〔空与超 200 rune〕/ `INVALID_DETAIL` / `INVALID_EVIDENCE_REF` / `INVALID_SOURCE` / `INVALID_REPOSITORY_ID` / `INVALID_OCCURRED_AT`）+ 执行序首错断言（构造多错输入断言只报第一个，如非法枚举 + 空 title → 只报 `INVALID_WORKFLOW_TYPE`）；断言按错误码非错误信息全文
2. **派生单元测试**（`backend/internal/progress/derive_test.go`，纯内存，无需 DB）：空集 / 仅 note / 从未开始（无 phase_started）/ 进行中 / 全部完结（最新 phase_started 后同 key phase_completed）/ 补录同刻 tiebreak（同 occurred_at+created_at 下 id DESC 序判定）/ recent 截断（>10 条取前 10）全部边界
3. **集成测试**（真实 PostgreSQL，DATABASE_URL 优先回落 `backend/.env`，无 DB 时 `t.Skipf`——沿 standard 集成测试冻结模式；独立 fixture + `t.Cleanup` 清理）：
   - `backend/internal/progress/service/service_integration_test.go`：Create→List round-trip（三键链倒序断言）+ 三轨过滤 + Delete 后 NotFound + Create 各错误分支（repository 不存在 → InvalidArgument 含 `REPOSITORY_NOT_FOUND`；校验失败码抽查）
   - `backend/internal/projectcontext/connect/server_integration_test.go` 追加 brief progress 场景：**round-trip**（Create phase_started + task_completed 后 GetProjectBrief，`progress.current_phase_key` = 录入 key、`recent_events` 含录入事件）+ **派生断言**（录入 `phase_completed`（同 task_key）后 `progress.current_phase_key` 为空串——"phase_completed 后当前 phase 为空" DoD 断言）+ **空态恒构造**（0 事件仓库 progress 块非 nil + 空数组 + 零值字段）

#### Scenario: DoD 测试断言全绿

- **WHEN** `backend/` 执行 `go test ./...`
- **THEN** 全部测试通过（有 DB 时集成测试执行，无 DB 时 skip 非 fail）；校验规则每条有单元测试覆盖且含非法用例矩阵；brief 集成测试含 progress round-trip 与派生断言（含 phase_completed 后当前 phase 为空）

### Requirement: 工具链门禁与运行时约束

本子任务 SHALL 满足 dev_plan phase15-06 DoD 门禁，并遵守运行时约束：

- 门禁命令（全部零退出码）：`proto/` 目录 `make lint` / `make build` / `make breaking` / `make gen`；`backend/` 目录 `go build ./...` / `go vet ./...` / `go test ./...`
- 运行时约束：用户已手动开启前后端服务器——本子任务**禁止**重启后端、禁止在其他端口重复开启服务器；运行中的服务器不热加载本子任务代码（预期内，浏览器完整会话验证归 phase15-07/08）；`0013` 经 `psql` 幂等应用到开发库以满足集成测试前提
- 不做（dev_plan §4 冻结 + 本子任务边界）：git 自动采集与 agent 写回、前端切片实现（phase15-07）、dogfooding 固定录入集与验收报告（phase15-08）、根级文档回写（phase15-09）、`UpdateProgressEvent`、phase11 `PhaseEntry` 改动

#### Scenario: 门禁全绿收口

- **WHEN** 全部实现与测试完成
- **THEN** 七条门禁命令退出码全 0；git 工作区改动仅含本 spec §What Changes 文件清单（+ 本 spec 三件套目录）；变更保持未提交，待用户最终确认后手动提交

## 与上游单值一致声明

- proto 合同、错误语义、模块分层、ProgressReader、brief 装配：全部逐字源 = phase15-04 对应 Requirement；DDL、错误码总表、执行序、派生算法、SQL 形态：全部逐字源 = phase15-03 对应 Requirement
- 合法矩阵 12 格 / K-1~K-5 正则 / 派生语义 / brief `progress = 9` 的语义上游为 phase15-02（经 03/04 间接承接，本 spec 不重复展开）
- 本 spec 不承载任何新设计决策；实现中发现上游文档与既有代码冲突时，以上游 spec 冻结口径为准并回改实现，不得反向改写上游
