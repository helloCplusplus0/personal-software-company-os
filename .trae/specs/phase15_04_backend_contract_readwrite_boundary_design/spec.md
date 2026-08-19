# phase15-04 产出后端合同与读写边界设计 Spec

## Why

`phase15-02` 已冻结事件模型语义（合法矩阵 / K 正则 / 派生规则 / brief 前后对照表），`phase15-03` 已冻结数据模型与校验派生设计（DDL / 9+2 错误码 / 派生算法实现序），但后端合同层仍未落地到"可直接进入实现"粒度：`psco.progress.v1` 合同（消息 + 枚举 + 3 RPC envelope 与错误语义）、Go 模块分层、`ProgressReader` 跨模块读取接口签名、brief `progress = 9` 装配演进点。`phase15-06`（后端实现）以本 spec 为直接设计上游；DP-2（repository_id 存在性校验承接位）在本 spec 裁决。

`phase15-04` 是实现设计类子任务：纯设计文档冻结，不写任何代码、不建 proto 文件、不改迁移；proto 合同以注释版草案承载（实现归 phase15-06）。

## What Changes

- 冻结 `proto/psco/progress/v1/progress.proto` 合同设计（注释版草案，phase15-06 逐字转写）：3 枚举（`WorkflowType` / `EventKind` / `ProgressSource`）+ `ProgressEvent` 消息（11 字段）+ 3 RPC 请求/响应 envelope + `ProgressService`（含无 `Update` 的显式声明，裁决⑨）
- 冻结 3 RPC 错误语义（每 RPC 三要素：请求 / 响应 / 错误）；裁决 DP-2：repository 存在性校验承接位 = `progress/candidate/RepositoryReader`（沿 `standard/candidate/TargetReader` 模式）；新增第 3 类错误码 `REPOSITORY_NOT_FOUND`（跨模块引用校验码）
- 冻结 Go 模块分层：`backend/internal/progress/`（根包 `types.go` / `errors.go` / `validate.go` / `derive.go` + `connect/` / `service/` / `repository/` / `candidate/` 四子包）+ `connecterrors` 登记点 + `platform` 装配点
- 冻结 `ProgressReader` 跨模块读取接口签名（消费方 `projectcontext/candidate` 拥有，`progress/service.QueryService` 实现，platform 装配点注入，沿 `StandardReader` 模式）
- 冻结 brief `progress = 9` 装配演进点：`BriefProgress` 内联消息字段号自然序冻结（1-4）+ `project_context.proto` 跨包 import + Go 装配步骤 + 转换函数复用关系

## Impact

- Affected specs:
  - `phase15-06`（本 spec 为其合同 + 后端模块 + 装配的直接设计上游；proto 草案逐字转写为 `progress.proto`）
  - `phase15-03`（本 spec 消费其错误码总表与执行序；新增 `REPOSITORY_NOT_FOUND` 为显式补充而非改写；DP-2 承接位裁决补齐其执行序第 6 步）
  - `phase15-05`（前端消费本 spec 的 RPC envelope 与错误码作为表单承接输入；DP-1 数据通道经 `GetProjectBrief.progress` 的可行性由本 spec `ProgressReader` 输出形态保证）
  - `phase15-02`（BriefProgress 字段清单语义上游；本 spec 冻结字段号分配——其显式保留项）
- Affected code: 无（设计冻结；`proto/` / `backend/` 零改动）
- 验收产物：本目录 `tasks.md / checklist.md`

## ADDED Requirements

### Requirement: psco.progress.v1 合同设计必须可直接进入 proto 实现

本 spec SHALL 冻结以下 proto 草案（phase15-06 逐字转写为 `proto/psco/progress/v1/progress.proto`，消息 / 枚举 / 字段号 / 注释零再决策）：

```proto
// PSCO Progress 项目推进时间轴协议合同源
//
// 文档定位：本文件是 phase15 项目推进时间轴（三轨 append-only 事件流）的唯一合同定义入口。
// 上游规格：
//   - phase15-02 事件模型语义边界与 brief 演进对照 Spec（合法矩阵 / K-1~K-5 / 派生语义）
//   - phase15-03 数据模型与校验派生设计 Spec（DDL / 9+2 错误码 / 派生算法实现序）
//   - phase15-04 后端合同与读写边界设计 Spec（本文件逐字源）
//
// 合同约束：
//   - Progress 是 repository 锚定的治理层信息流，不是第六业务主实体（裁决②⑥）
//   - append-only：无 UpdateProgressEvent RPC，事件一经录入只有整条删除一种
//     修正路径（裁决⑨；service 定义处显式声明）
//   - "当前 phase / 最新任务"为读取侧派生值，不落库、不缓存（裁决③）；
//     派生摘要经 brief progress 块消费（BriefProgress），完整流经 ListProgressEvents
//   - ListProgressEvents 不分页（个人规模，反假大空；裁决⑨）
//   - source 预留 manual / git / agent 三值，创建入口仅开放 manual（裁决⑧）
//   - 错误语义三类归一：NotFound / InvalidArgument / Internal（沿 standard 合同约束）
//   - task_key / detail / evidence_ref 可空字段 proto 侧以空串承接"未填"，
//     DB 侧 NULL 与空串读取等价（NULL → 空串，转换在 repository 层）
//
// 生成入口：
//   在 proto/ 目录执行 make gen，或直接执行 buf generate。
//   生成产物：
//     - Go:    backend/internal/gen/proto/psco/progress/v1/
//     - Go Connect: backend/internal/gen/connect/psco/progress/v1/
//     - TS:    frontend/src/gen/proto/psco/progress/v1/

syntax = "proto3";

package psco.progress.v1;

import "google/protobuf/timestamp.proto";

// Go 包路径（生成代码落点）
option go_package = "github.com/psco/backend/internal/gen/proto/psco/progress/v1;progressv1";

// ============================================================================
// 枚举
// ============================================================================

// WorkflowType 三轨 workflow（对齐 docs/ 三目录与三推进链，裁决④）。
// 与 DDL workflow_type CHECK 单值映射。
enum WorkflowType {
  WORKFLOW_TYPE_UNSPECIFIED = 0;
  WORKFLOW_TYPE_PHASE = 1;
  WORKFLOW_TYPE_AUDIT = 2;
  WORKFLOW_TYPE_FIX = 3;
}

// EventKind 事件类型（裁决⑤：audit/fix 轨禁止 phase 边界标记——
// 组合合法性由应用层 V7 承接，DB 不建组合约束）。
enum EventKind {
  EVENT_KIND_UNSPECIFIED = 0;
  EVENT_KIND_PHASE_STARTED = 1;
  EVENT_KIND_PHASE_COMPLETED = 2;
  EVENT_KIND_TASK_COMPLETED = 3;
  EVENT_KIND_NOTE = 4;
}

// ProgressSource 事件来源（预留 manual/git/agent 三值；本阶段创建入口仅 manual，裁决⑧）。
enum ProgressSource {
  PROGRESS_SOURCE_UNSPECIFIED = 0;
  PROGRESS_SOURCE_MANUAL = 1;
  PROGRESS_SOURCE_GIT = 2;
  PROGRESS_SOURCE_AGENT = 3;
}

// ============================================================================
// 核心消息
// ============================================================================

// ProgressEvent 推进事件（progress_events 表投影）。
// List / Create 响应与 brief progress 块统一使用本消息——
// 不另造 ProgressEventSummary（沿 standard "不另造 StandardSummary" 模式）。
message ProgressEvent {
  string id = 1;
  // 锚点：进度事实唯一归属仓库（裁决②）
  string repository_id = 2;
  WorkflowType workflow_type = 3;
  EventKind event_kind = 4;
  // 任务项标识；可空：空串 = 未填（note 轨）
  string task_key = 5;
  // 一句话标题（非空上限 200，应用层承接）
  string title = 6;
  // 展开说明；可空：空串 = 未填（上限 2000，应用层承接）
  string detail = 7;
  // 证据导航引用；可空：空串 = 未填（/ 或 https:// 前缀，裁决⑦）
  string evidence_ref = 8;
  ProgressSource source = 9;
  // 用户声明发生时间（允许补录历史，与 created_at 分离）
  google.protobuf.Timestamp occurred_at = 10;
  // 系统录入时间
  google.protobuf.Timestamp created_at = 11;
}

// ============================================================================
// 请求 / 响应
// ============================================================================

// ListProgressEventsRequest 完整事件流读取请求。
// workflow_type 为过滤参数：UNSPECIFIED（零值/未设置）= 不过滤（三轨全量）。
message ListProgressEventsRequest {
  // 必填：读取锚点
  string repository_id = 1;
  // 可选过滤：设置三值之一 = 只读该轨；UNSPECIFIED = 三轨全量
  WorkflowType workflow_type = 2;
}

// ListProgressEventsResponse 完整事件流响应（不分页，裁决⑨）。
// events 按三键链倒序（occurred_at DESC, created_at DESC, id DESC）排列。
message ListProgressEventsResponse {
  repeated ProgressEvent events = 1;
}

// CreateProgressEventRequest 创建事件请求。
// source 为 optional：未设置 → 服务端归一 manual（与 DDL DEFAULT 对齐）；
// 显式设置非 MANUAL（UNSPECIFIED/GIT/AGENT）→ INVALID_SOURCE（V9d）。
message CreateProgressEventRequest {
  // 必填：写入目标锚点（不存在 → InvalidArgument，shared_baseline §3.2 冻结）
  string repository_id = 1;
  // 必填：UNSPECIFIED → INVALID_WORKFLOW_TYPE（V1a）
  WorkflowType workflow_type = 2;
  // 必填：UNSPECIFIED → INVALID_EVENT_KIND（V1b）
  EventKind event_kind = 3;
  // 按矩阵可空：必填格空串 → TASK_KEY_REQUIRED；格式不符 → TASK_KEY_FORMAT_INVALID
  string task_key = 4;
  // 必填：空串或超 200 字符 → INVALID_TITLE
  string title = 5;
  // 可选：上限 2000 字符 → INVALID_DETAIL
  string detail = 6;
  // 可选：非 / 或 https:// 前缀 → INVALID_EVIDENCE_REF
  string evidence_ref = 7;
  // 可选：未设置归一 manual
  optional ProgressSource source = 8;
  // 必填：nil → INVALID_OCCURRED_AT（envelope 前置；用户声明时间无合法零值）
  google.protobuf.Timestamp occurred_at = 9;
}

// CreateProgressEventResponse 创建响应，返回创建后的完整 ProgressEvent
// （含服务端生成的 id / source 归一值 / created_at；沿 CreateStandard 模式）。
message CreateProgressEventResponse {
  ProgressEvent event = 1;
}

// DeleteProgressEventRequest 删除请求（误录修正；整条删除是唯一修正路径）。
message DeleteProgressEventRequest {
  string id = 1;
}

// DeleteProgressEventResponse 删除响应（空；沿 DeleteStandard 模式）。
message DeleteProgressEventResponse {}

// ============================================================================
// 服务
// ============================================================================

// ProgressService 项目推进时间轴写读服务。
//
// 职责边界（phase15-04 冻结）：
//   - 本服务是 Progress 事件流写读的唯一正式承接位（canonical owner）
//   - append-only 纯净性（裁决⑨）：无 UpdateProgressEvent RPC——
//     事件一经录入只有整条删除一种修正路径，任何设计不得使历史事件
//     因新事件而消失或变形
//   - 9 条校验规则（V1-V9）+ envelope 前置 + repository 存在性校验
//     在写路径统一执行（执行序沿 phase15-03 冻结 6 步）
//   - web 写路径与 agent 读路径同包同 service，不出现第二套 canonical API
service ProgressService {
  // 读取完整事件流（三键链倒序，不分页；workflow_type 可选过滤）。
  // 失败语义：repository_id 非法 UUID → InvalidArgument；仓库不存在 →
  //           NotFound（读锚点语义，沿 GetProjectBrief）；无事件或过滤后
  //           为空 → 空列表（非错误）；读取失败 → Internal。
  rpc ListProgressEvents(ListProgressEventsRequest) returns (ListProgressEventsResponse);

  // 创建推进事件（校验执行序 6 步：envelope 前置 → V1a → V1b → V7 →
  // task_key 矩阵分支 → V9 文本顺序 → repository 存在性）。
  // 失败语义：任一校验失败（INVALID_REPOSITORY_ID / INVALID_OCCURRED_AT /
  //           INVALID_WORKFLOW_TYPE / INVALID_EVENT_KIND /
  //           EVENT_KIND_NOT_ALLOWED / TASK_KEY_REQUIRED /
  //           TASK_KEY_FORMAT_INVALID / INVALID_TITLE / INVALID_DETAIL /
  //           INVALID_EVIDENCE_REF / INVALID_SOURCE / REPOSITORY_NOT_FOUND）
  //           → InvalidArgument；写入失败 → Internal。
  rpc CreateProgressEvent(CreateProgressEventRequest) returns (CreateProgressEventResponse);

  // 删除推进事件（误录修正；无软删除、无 Update）。
  // 失败语义：id 非法 UUID → InvalidArgument；不存在 → NotFound；删除失败 → Internal。
  rpc DeleteProgressEvent(DeleteProgressEventRequest) returns (DeleteProgressEventResponse);

  // 显式声明：本服务不提供 UpdateProgressEvent（append-only 语义纯净，裁决⑨）。
  // 误录修正 = Delete + 重新 Create；本注释为合同级声明，勿删。
}
```

**合同设计决策冻结**（草案注释承载 + 本节展开理由）：

1. **`ProgressEvent` 字段号 1-11 按 DDL 列序自然分配**：与 phase15-03 DDL 草案 11 列一一对应（id → created_at），实现期零再决策。
2. **`workflow_type` 过滤参数为非 optional 枚举**：proto3 零值 `UNSPECIFIED` 即"未设置"，语义定义为"不过滤（三轨全量）"——与"零值 = 默认行为"惯例一致，无需 optional 区分（agent 过滤单轨时显式传三值之一）。
3. **`source` 为 `optional` 枚举**：区分"未设置（→ 归一 manual，与 DDL DEFAULT 'manual' 对齐）"与"显式设置（→ 必须 MANUAL，否则 INVALID_SOURCE）"两种输入形态，保证 V9d 可触发（沿 `CreateStandardRequest.status` optional 模式）。
4. **可空文本三字段（task_key / detail / evidence_ref）proto 侧以空串承接"未填"**：不使用 optional string；DB NULL 与空串读取等价（NULL → 空串，转换在 repository 层承接），单值语义无双态。
5. **Create 响应返回完整 `ProgressEvent`**：前端创建后直接回显（含服务端生成 id / created_at / source 归一值），无需二次 List。
6. **Delete 响应为空消息**：沿 `DeleteStandardResponse {}` 模式。
7. **时间字段用 `google.protobuf.Timestamp`**：沿 `Standard.created_at / updated_at` 模式。

#### Scenario: proto 草案可逐字转写

- **WHEN** phase15-06 实现 proto 文件
- **THEN** 以上草案逐字转写为 `proto/psco/progress/v1/progress.proto`（仅可调整：无），枚举 / 消息 / 字段号 / RPC 定义 / 注释零再决策
- **AND** `buf lint / build / breaking` 全通过；生成链（make gen）产出 Go pb / Go Connect / TS 三端产物

### Requirement: 3 RPC 错误语义逐个三要素冻结且 DP-2 承接位单值裁决

本 spec SHALL 冻结 3 RPC 的错误语义（每 RPC 三要素齐全），并裁决 DP-2：

**RPC 1：`ListProgressEvents`（读路径）**

| 要素 | 冻结内容 |
|---|---|
| 请求 | `repository_id`（必填）+ `workflow_type`（可选过滤，UNSPECIFIED = 不过滤） |
| 响应 | `events[]` 全量三键链倒序；不分页 |
| 错误 | `repository_id` 非 UUID → InvalidArgument `[INVALID_REPOSITORY_ID]`；仓库不存在 → **NotFound**；无事件 / 过滤后空 → 空列表（非错误）；读取失败 → Internal |

**RPC 2：`CreateProgressEvent`（写路径）**

| 要素 | 冻结内容 |
|---|---|
| 请求 | 9 字段（repository_id / workflow_type / event_kind / task_key / title / detail / evidence_ref / source / occurred_at） |
| 响应 | 创建后的完整 `ProgressEvent`（含服务端生成值） |
| 错误 | 校验执行序 6 步逐项失败 → InvalidArgument（错误码总表见下）；写入失败 → Internal |

**RPC 3：`DeleteProgressEvent`（写路径）**

| 要素 | 冻结内容 |
|---|---|
| 请求 | `id`（事件 id，必填） |
| 响应 | 空消息 |
| 错误 | `id` 非 UUID → InvalidArgument `[INVALID_PROGRESS_EVENT_ID]`；不存在 → **NotFound**；删除失败 → Internal |

**错误码总表（phase15-03 冻结 9 业务码 + 2 envelope 码 + 本 spec 新增 2 码）**：

| 错误码 | 类别 | 判定逻辑 | RPC |
|---|---|---|---|
| `INVALID_WORKFLOW_TYPE` ~ `INVALID_SOURCE`（9 个） | 业务码（phase15-03 冻结，零改写） | 见 phase15-03 错误码总表 | Create |
| `INVALID_REPOSITORY_ID` | envelope 前置（phase15-03 冻结） | UUID 格式层 | List / Create / Delete（Delete 为 `INVALID_PROGRESS_EVENT_ID`） |
| `INVALID_OCCURRED_AT` | envelope 前置（phase15-03 冻结） | Timestamp 已设置 | Create |
| `REPOSITORY_NOT_FOUND` | **跨模块引用校验码（本 spec 新增）** | repository 存在性查询返回 false | Create（InvalidArgument 语义） |
| `INVALID_PROGRESS_EVENT_ID` | **ID 格式码（本 spec 新增）** | Delete 的 id 非 UUID | Delete |

- `REPOSITORY_NOT_FOUND` 是 phase15-03 执行序第 6 步"repository 存在性校验（语义 invalid_argument）"的稳定错误码承接——phase15-03 显式留位（"承接位归 phase15-04 DP-2 裁决"），本 spec 补齐码值不改写其执行序。
- `INVALID_PROGRESS_EVENT_ID` 与 `INVALID_REPOSITORY_ID` 同为 ID 格式码，按字段命名（Delete 以事件 id 定位，无 repository_id 输入）。

**DP-2 裁决（repository_id 存在性校验承接位）**：

- **承接位 = `backend/internal/progress/candidate/RepositoryReader`**（progress 模块自拥有 candidate 子包，沿 `standard/candidate/TargetReader` 模式）：
  - 签名：`RepositoryExists(ctx context.Context, repositoryID string) (bool, error)`——**纯存在性事实查询**（返回 bool），不承载业务错误语义
  - 查询形态：`SELECT EXISTS(SELECT 1 FROM repositories WHERE id = $1)`（沿 TargetReader.EnsureTargetExists 内层查询形态）
  - 查询失败 → 包装原始错误（Internal，由 connecterrors 兜底）
- **错误语义包装归 service 层**（各路径按冻结语义各自包装）：
  - Create 路径（CommandService）：`!found` → `fmt.Errorf("%w: [REPOSITORY_NOT_FOUND] repository %s does not exist", ErrInvalidInput, repositoryID)`（InvalidArgument，沿 standard TargetReader "target 不存在归 InvalidArgument" 模式）
  - List 路径（QueryService）：`!found` → `progress.ErrRepositoryNotFound`（CodeNotFound，读锚点语义）
- **DB FK RESTRICT 为存储层兜底，非校验承接位**（沿 phase15-03 冻结口径；service 层校验先行保证错误语义确定性，FK 只拦截绕过 service 的写入）

**Create（InvalidArgument）与 List（NotFound）错误语义辨析留档**：

- 同一 `repository_id` 在两 RPC 中错误码不同是**语义正确**的：
  - Create 的 repository_id 是**写入目标外键引用**——"引用了一个不存在的实体"是输入构造错误 → InvalidArgument（shared_baseline §3.2 冻结："不存在 → 创建时 invalid_argument"）
  - List 的 repository_id 是**读取锚点**——"要读的资源不存在"是资源定位失败 → NotFound（沿 `GetProjectBrief` / `GetStandard` 读路径语义）
- 两语义各自有冻结上游，phase15-04 不做统一化改写（统一为同一码将违反其一上游）。

**哨兵错误清单（`backend/internal/progress/errors.go`，phase15-06 逐字转写）**：

```go
// 业务错误哨兵值（由 connecterrors.MapToConnectError 统一映射 Connect code）。
var (
    // ErrProgressEventNotFound 目标事件不存在（Delete 定位失败）。
    // 映射为 connect.CodeNotFound。
    ErrProgressEventNotFound = errors.New("progress event not found")

    // ErrRepositoryNotFound 读路径锚点仓库不存在（ListProgressEvents）。
    // 映射为 connect.CodeNotFound（读锚点语义，沿 GetProjectBrief）。
    ErrRepositoryNotFound = errors.New("repository not found")

    // ErrInvalidInput 事件输入违反校验规则（V1-V9 + envelope 前置 +
    // REPOSITORY_NOT_FOUND 跨模块引用校验统一包装）。
    // 映射为 connect.CodeInvalidArgument。
    ErrInvalidInput = errors.New("invalid progress input")

    // ErrProgressReadFailed 事件流读取失败。
    // 映射为 connect.CodeInternal。
    ErrProgressReadFailed = errors.New("progress read failed")

    // ErrProgressWriteFailed 事件写入/删除失败。
    // 映射为 connect.CodeInternal。
    ErrProgressWriteFailed = errors.New("progress write failed")
)
```

**`connecterrors` 登记点冻结**（`backend/internal/connecterrors/connect_errors.go` 修改）：

- CodeNotFound 组追加：`progress.ErrProgressEventNotFound`、`progress.ErrRepositoryNotFound`
- CodeInvalidArgument 组追加：`progress.ErrInvalidInput`
- CodeInternal 兜底组追加显式引用：`_ = progress.ErrProgressReadFailed`、`_ = progress.ErrProgressWriteFailed`

#### Scenario: 错误语义可直接进入实现与测试

- **WHEN** phase15-06 实现 Connect handler 与错误映射
- **THEN** 3 RPC 错误分支按本节表格逐项实现；DP-2 承接位按 `RepositoryReader.RepositoryExists` 签名逐字实现；connecterrors 登记按清单修改
- **AND** 集成测试可按错误码（非错误信息全文）断言每条错误分支（List 不存在 → NotFound / Create 不存在 → InvalidArgument 含 REPOSITORY_NOT_FOUND / Delete 不存在 → NotFound）

### Requirement: Go 模块分层与文件落点必须冻结且沿 standard 模式

本 spec SHALL 冻结 `backend/internal/progress/` 模块分层（沿 `standard` 模块模式，phase15-06 直接实现）：

**根包（`backend/internal/progress/`，`package progress`）**：

| 文件 | 承接内容 |
|---|---|
| `types.go` | 受控枚举（`WorkflowType` / `EventKind` / `ProgressSource`，string 形态对齐 DDL CHECK 值）+ 读模型 `ProgressEventReadResult`（11 字段）+ 派生摘要 `ProgressSummary`（四字段）+ 写模型 `CreateProgressEventInput`；消息结构从 proto 单向派生或显式对齐，不 import 生成 pb 包 |
| `errors.go` | 5 个哨兵错误（见上一 Requirement 清单） |
| `validate.go` | V1-V9 + envelope 前置校验（phase15-03 错误码总表与执行序 6 步逐字实现；`%w: [CODE] message` 包装格式） |
| `derive.go` | `DeriveProgressSummary(events []ProgressEventReadResult) ProgressSummary` 纯函数（phase15-03 三算法逐字实现；零 I/O 零时间函数） |

**受控枚举与核心类型（types.go 草案，phase15-06 逐字转写）**：

```go
// WorkflowType 三轨 workflow（受控枚举，对齐 DDL workflow_type CHECK）。
type WorkflowType string

const (
    WorkflowTypePhase WorkflowType = "phase"
    WorkflowTypeAudit WorkflowType = "audit"
    WorkflowTypeFix   WorkflowType = "fix"
)

// EventKind 事件类型（受控枚举，对齐 DDL event_kind CHECK）。
type EventKind string

const (
    EventKindPhaseStarted   EventKind = "phase_started"
    EventKindPhaseCompleted EventKind = "phase_completed"
    EventKindTaskCompleted  EventKind = "task_completed"
    EventKindNote           EventKind = "note"
)

// ProgressSource 事件来源（受控枚举；本阶段创建入口仅 manual）。
type ProgressSource string

const (
    ProgressSourceManual ProgressSource = "manual"
    ProgressSourceGit    ProgressSource = "git"
    ProgressSourceAgent  ProgressSource = "agent"
)

// ProgressEventReadResult 推进事件读取结果（progress_events 表投影）。
type ProgressEventReadResult struct {
    ID           string         `json:"id"`
    RepositoryID string         `json:"repository_id"`
    WorkflowType WorkflowType   `json:"workflow_type"`
    EventKind    EventKind      `json:"event_kind"`
    TaskKey      string         `json:"task_key"`
    Title        string         `json:"title"`
    Detail       string         `json:"detail"`
    EvidenceRef  string         `json:"evidence_ref"`
    Source       ProgressSource `json:"source"`
    OccurredAt   time.Time      `json:"occurred_at"`
    CreatedAt    time.Time      `json:"created_at"`
}

// ProgressSummary 派生摘要（与 BriefProgress 四字段 1:1；空态恒构造零值）。
type ProgressSummary struct {
    CurrentPhaseKey     string                   `json:"current_phase_key"`
    CurrentPhaseLabel   string                   `json:"current_phase_label"`
    LatestTaskCompleted *ProgressEventReadResult  `json:"latest_task_completed"` // 可空：无任务完成时 nil
    RecentEvents        []ProgressEventReadResult `json:"recent_events"`        // min(10, len)；恒非 nil（空态为空切片）
}

// CreateProgressEventInput 创建事件输入（source 已由 service 归一 manual）。
type CreateProgressEventInput struct {
    RepositoryID string         `json:"repository_id"`
    WorkflowType WorkflowType   `json:"workflow_type"`
    EventKind    EventKind      `json:"event_kind"`
    TaskKey      string         `json:"task_key"`
    Title        string         `json:"title"`
    Detail       string         `json:"detail"`
    EvidenceRef  string         `json:"evidence_ref"`
    Source       ProgressSource `json:"source"`
    OccurredAt   time.Time      `json:"occurred_at"`
}
```

**四子包**：

| 子包 | 承接内容 |
|---|---|
| `connect/` | Connect handler（proto request 解包 → service 调用 → proto response 组装）；**导出 `DomainProgressEventToProto(e progress.ProgressEventReadResult) *progressv1.ProgressEvent`**（沿 `standard/connect.DomainStandardToProto` 导出模式）供 projectcontext/connect 组装 BriefProgress 复用；模块内部转换函数小写私有 |
| `service/` | `QueryService`（`ListProgressEvents` + `GetProgressSummary`）+ `CommandService`（`CreateProgressEvent` + `DeleteProgressEvent`）；校验触发与 repository 存在性包装在此层 |
| `repository/` | `ProgressEventStore`：单一查询（phase15-03 冻结 SQL 形态，repository_id 必选 + workflow_type 可选过滤 + 三键链 ORDER BY）+ INSERT（返回完整行）+ DELETE BY id |
| `candidate/` | `RepositoryReader`（DP-2 承接位，见上一 Requirement） |

**service 依赖注入（沿 standard 装配模式）**：

- `QueryService`：依赖 `store *repository.ProgressEventStore` + `repositoryReader *candidate.RepositoryReader`
- `CommandService`：依赖 `store` + `repositoryReader`

**platform 装配点冻结**（`backend/internal/platform/router.go` 修改）：

- 新增 `buildProgress(pool *pgxpool.Pool) (*progressservice.QueryService, *progressservice.CommandService)`：构造 RepositoryReader → ProgressEventStore → 双 service（沿 `buildStandard` 模式）
- 新增 `mountProgressConnect(r chi.Router, querySvc, commandSvc)`：`progressv1connect.NewProgressServiceHandler` + `r.Handle(path+"*", http.StripPrefix("/api", handler))`（沿 `mountStandardConnect` 模式）
- `buildProjectContext` 签名演进：新增 `progressReader projectcontextcandidate.ProgressReader` 参数（沿 standardReader 注入模式，phase15-06 实现）

#### Scenario: 分层与落点可直接进入实现

- **WHEN** phase15-06 实现 Go 模块
- **THEN** 文件落点 / 类型签名 / 枚举值 / 装配函数按本节逐字实现；与 standard 模块（`backend/internal/standard/`）逐文件对照结构一致
- **AND** 不出现第五子包、第二套转换函数、或 service 直写跨模块 SQL

### Requirement: ProgressReader 跨模块读取接口签名单值冻结

本 spec SHALL 冻结 `ProgressReader` 接口（沿 `StandardReader` 模式：消费方拥有的 candidate 接口 + provider service 实现 + platform 装配点注入）：

**接口定义**（落点 `backend/internal/projectcontext/candidate/context_readers.go`，与 `StandardReader` 同文件追加，phase15-06 逐字转写）：

```go
// ProgressReader 项目进度派生摘要读取接口（消费方拥有的 candidate 接口）。
//
// phase15-04 冻结：brief 对 Progress 摘要的读取必须通过本接口承接，
// 由 platform 装配点注入 progress/service.QueryService 作为实现；
// projectcontext 不得直接书写 progress_events 表 SQL 或复制其派生逻辑。
type ProgressReader interface {
    // GetProgressSummary 读取该仓库进度派生摘要（当前 phase + 最新任务 +
    // 最近事件；三键链倒序派生，算法沿 phase15-03 冻结）。
    // 失败语义：读取失败 → progress.ErrProgressReadFailed；
    //           仓库无事件 → 零值摘要 + 空 RecentEvents（非错误，空态恒构造）。
    GetProgressSummary(ctx context.Context, repositoryID string) (progress.ProgressSummary, error)
}
```

**实现位**：`progress/service.QueryService.GetProgressSummary(ctx, repositoryID)`——内部调用 store 单一查询（repository_id 过滤、无 workflow_type 过滤）取全量事件 → `DeriveProgressSummary` 纯函数计算四字段；与 `ListProgressEvents` RPC 同源同派生（shared_baseline §3.5"brief 与 List RPC 同源同派生"的接口形态落地）。

**装配接线**：`buildProjectContext(pool, standardReader, progressReader)`——progressReader 由 `buildProgress` 产出的 QueryService 注入（沿 standardReader 由 buildStandard 的 QueryService 注入模式）。

**设计决策冻结**：

1. **Reader 输出 = `ProgressSummary` 四字段，不含全量事件集**：brief 摘要块只需要当前 phase / 最新任务 / 最近 10 条；全量事件流唯一经 `ListProgressEvents` RPC 消费（裁决⑩消费分层在接口形态上的落地——Reader 不膨胀为第二全量通道）。
2. **`RecentEvents` 恒非 nil**（空态为空切片）：空态恒构造约束（shared_baseline §3.5）在 domain 层的承接；brief 装配侧与前端消费侧无双套判空。
3. **`LatestTaskCompleted` 为指针可空**：无任务完成时 nil，proto 组装时不设置（phase15-02"零值不设置"语义的 Go 承接）。
4. **Reader 不做 repository 存在性校验**：brief 主流程的 ReadRepository 已先行承接仓库存在性（NotFound）；Reader 对不存在仓库返回零值摘要（EXISTS 语义由查询天然承接：无行 → 空集 → 零值）。

#### Scenario: 接口签名可直接进入实现

- **WHEN** phase15-06 实现 projectcontext candidate 接口与 progress service 实现
- **THEN** 接口定义逐字转写至 `context_readers.go`；`GetProgressSummary` 实现经 store 单查询 + DeriveProgressSummary，无第二套派生路径
- **AND** brief 装配不出现 projectcontext 直读 progress_events 表 SQL

### Requirement: brief progress = 9 装配演进点必须单值冻结

本 spec SHALL 冻结 brief 装配演进的全部设计点（phase15-06 直接实现）：

**proto 侧（`project_context.proto` 修改）**：

1. 新增 `import "psco/progress/v1/progress.proto";`（沿既有 `import "psco/standard/v1/standard.proto"` 模式）
2. 新增内联消息 `BriefProgress`（字段号按 phase15-02 保留项自然序冻结）：

```proto
// BriefProgress — agent 项目简报进度摘要块（phase15 裁决⑩）。
// 定义于 project_context.proto（内联轻量消息）；latest_task_completed 与
// recent_events 元素同型复用 psco.progress.v1.ProgressEvent（不建第二套摘要消息）。
// 空态语义：Go 装配侧恒构造本块（非 nil）——0 事件时字段为零值、
// recent_events 为空数组，前端进度区与 agent 单值消费，无双套判空逻辑。
message BriefProgress {
  // 当前 phase 的 task_key（phaseNN）；空值含从未开始 / 全部完结两种情形（同型零值）
  string current_phase_key = 1;
  // 当前 phase 标题（该最新 phase_started 的 title）
  string current_phase_label = 2;
  // 最新完成任务项事件（三轨同序取最新；无任务完成时不设置）
  psco.progress.v1.ProgressEvent latest_task_completed = 3;
  // 最近 N=10 条三轨混合事件（三键链倒序）
  repeated psco.progress.v1.ProgressEvent recent_events = 4;
}
```

3. `GetProjectBriefResponse` 追加 `BriefProgress progress = 9;`（槽位 2/3/4 保持 reserved 不动；字段号 9 为 phase15-02 冻结的正式演进）

**Go 装配侧（projectcontext 模块修改）**：

1. `types.go`：`ProjectBriefReadResult` 新增 `Progress progress.ProgressSummary` 字段（值类型——空态恒构造零值；domain 类型直接透传，沿 `Standards []standard.StandardReadResult` 模式）
2. `service/query_service.go`：`GetProjectBrief` 编排新增步骤 6——

```go
// 6. progress 摘要（candidate 接口承接，不复制 progress_events 表 SQL；
//    空态恒构造：0 事件时零值摘要 + 空 RecentEvents，非 nil）
progressSummary, err := s.progressReader.GetProgressSummary(ctx, repositoryID)
if err != nil {
    return nil, err
}
```

（编排顺序：ReadRepository → products[] → modules[] → decisions[] → standards[] → progress；失败语义沿 reader 冻结：progress 读取失败 → `progress.ErrProgressReadFailed`（CodeInternal，由 connecterrors 统一映射））

3. `connect/server.go`：组装 `BriefProgress.progress` 时复用 `progress/connect` 导出的 `DomainProgressEventToProto` 转换事件元素（沿"standards[] 复用 standard/connect.DomainStandardToProto"既有模式）；`RecentEvents` 空态组装为空数组（非 nil 切片）

**brief 顶层块口径**：5 顶层块 + `progress` 摘要块（repository / products / modules / decisions / standards[] + progress；`progress = 9`）

#### Scenario: 装配演进可直接进入实现

- **WHEN** phase15-06 修改 project_context.proto 与 projectcontext 模块
- **THEN** import / BriefProgress 字段号（1-4）/ progress = 9 / Go 编排步骤 6 / 转换函数复用按本节逐字实现
- **AND** `buf breaking` 对既有字段（1-8）零破坏；brief 集成测试含 progress round-trip 与空态断言（恒构造 + 空数组 + 零值字段）

### Requirement: 本 spec 与 phase15-02/03 语义上游单值一致且不偷渡

本 spec 的全部设计内容 SHALL 与语义上游单值一致：

- 3 枚举值域与 DDL CHECK 三列值域单值（phase15-03 DDL 草案）；`ProgressEvent` 11 字段与 DDL 11 列一一对应
- Create 校验执行序与错误码完全引用 phase15-03 冻结（零改写）；本 spec 仅新增 `REPOSITORY_NOT_FOUND`（其执行序第 6 步显式留位的承接码）与 `INVALID_PROGRESS_EVENT_ID`（Delete 的 id 格式码，不在 phase15-03 执行序范围——Delete 路径无该 spec 的输入结构）
- `ProgressSummary` 四字段与 phase15-02 `BriefProgress` 字段清单 1:1（current_phase_key / current_phase_label / latest_task_completed / recent_events）；空值双情形同型零值语义不变
- `BriefProgress` 字段号 1-4 自然序为本 spec 冻结项（phase15-02 显式保留："字段号分配与消息定义落点细节归 phase15-04"）
- 本 spec 不承载：DDL 细节、校验判定逻辑细化、派生算法（phase15-03 已冻结，仅引用）；前端组件树与交互规格（phase15-05）；实现代码与 proto 文件实体（phase15-06）
- 本 spec 不承载：DP-1（web 当前卡数据通道，归 phase15-05）/ DP-3（occurred_at 输入与时区口径，归 phase15-05 / phase15-08）

#### Scenario: 一致性可校验

- **WHEN** 独立复核执行
- **THEN** proto 草案与 phase15-02 矩阵 / phase15-03 DDL+错误码逐项比对一致；ProgressReader / BriefProgress / 装配点与 StandardReader / standards[] / buildProjectContext 既有模式逐项对照一致
- **AND** git 工作区中本 spec 仅为目录新增，零代码 / 零 proto / 零迁移文件 / 零根级文档改动
