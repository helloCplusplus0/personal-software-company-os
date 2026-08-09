# Phase05-11 Dashboard + Feedback 最小 Protocol Buffers 合同主线 Spec

## Why

`phase05-08` 已冻结 `Dashboard + Feedback` 的最小 `.proto` 合同设计，`phase05-10` 又将这些结论收口进正式规格正文。但截至当前，仓库主线里还没有 `Dashboard + Feedback` 对应的实际 `.proto` 文件、生成产物落点、`buf build / lint / generate / breaking` 执行入口覆盖，以及 DTO/HTTP 对 `.proto` 的单向承接主线。

如果继续围绕手写 DTO、handler 响应结构或页面类型推进 `phase05-12 / 13 / 14`，就会重新长出与 `.proto` 并列的第二套合同源。`phase05-11` 的目标，就是把已经冻结的合同设计推进为仓库内真实存在、可生成、可校验、可被实现直接引用的合同主线。

## What Changes

- 将 `phase05-08` 已冻结的 `Dashboard + Feedback` 最小 `.proto` 合同正式落地到现有 `proto/` workspace
- 冻结 `proto/psco/dashboard/v1/dashboard.proto` 的单一文件落点、包名、跨包 import 与生成产物归属
- 冻结 `buf build / lint / generate / breaking` 在当前仓库中的最小可运行入口，并要求继续复用 `proto/buf.yaml`、`proto/buf.gen.yaml`、`proto/Makefile`
- 冻结 `proto/README.md` 必须把 `Dashboard + Feedback` 纳入单一 proto 合同源总览与 RPC → HTTP 映射矩阵
- 冻结 `backend/internal/dashboard/types.go`、HTTP handler DTO、前端 adapter 与 `.proto` 的单向语义映射边界，阻断第二套合同源
- 明确当前阶段只落地合同源、生成入口、校验链与映射边界，不提前完成完整 gRPC / Connect 传输层迁移
- **BREAKING**：后续 `Dashboard + Feedback` 的实现与验收不得再把手写 JSON 结构、页面类型或 handler DTO 视为并列合同源，`.proto` 成为仓库内唯一合同定义入口

## Impact

- Affected specs:
  - `phase05_08_define_dashboard_feedback_proto_contract`
  - `phase05_10_dashboard_feedback_formal_spec`
  - 后续 `phase05-12` 后端与数据主线
  - 后续 `phase05-13` 前端主线
  - 后续 `phase05-14` 联调与验收
  - `phase03_11_decision_center_proto_mainline`
  - `phase04_11_product_repository_binding_proto_mainline`
- Affected code:
  - `proto/psco/dashboard/v1/dashboard.proto`
  - `proto/README.md`
  - `proto/Makefile`
  - `proto/buf.yaml`
  - `proto/buf.gen.yaml`
  - `backend/internal/gen/proto/psco/dashboard/v1/`
  - `frontend/src/gen/proto/psco/dashboard/v1/`
  - `backend/internal/dashboard/types.go`
  - `backend/internal/dashboard/handler/*.go`
  - `frontend/src/features/dashboard/`

## ADDED Requirements

### Requirement: Dashboard + Feedback 合同必须落地到现有 proto workspace

系统 SHALL 将 `Dashboard + Feedback` 的最小 `.proto` 合同落地到现有 `proto/` workspace，而不是为 `phase05` 新建第二个 proto 根目录、第二套 `buf.yaml` 或第二套生成入口。

#### Scenario: 合同源文件落点

- **WHEN** 执行 `phase05-11`
- **THEN** 仓库中必须存在 `proto/psco/dashboard/v1/dashboard.proto`
- **AND** 该文件必须作为 `Dashboard + Feedback` 当前阶段唯一合同定义入口
- **AND** 必须继续复用现有 `proto/buf.yaml` 作为同一个 buf workspace 的根配置
- **AND** 不得在 `backend/`、`frontend/` 或仓库根新增并列 proto 工作区

#### Scenario: 包名与版本语义落地

- **WHEN** `Dashboard + Feedback` 合同正式进入仓库
- **THEN** 包名与版本语义必须冻结为 `psco.dashboard.v1`
- **AND** 后续字段演进必须继续在该版本语义下进行
- **AND** 不得在落地阶段临时改写为第二套包名、目录层级或版本策略

### Requirement: dashboard.proto 必须完整承接 phase05-08 的最小合同设计

系统 SHALL 要求 `dashboard.proto` 在仓库落地时完整承接 `phase05-08` 已冻结的服务、消息、字段编号、枚举与演进规则，而不是重新发明“更适合实现”的第二套合同。

#### Scenario: 服务与消息落地

- **WHEN** 落地 `proto/psco/dashboard/v1/dashboard.proto`
- **THEN** 文件内必须存在单一 `DashboardService`
- **AND** 必须承接三类只读 RPC：
  - `GetDashboardOverview`
  - `GetFeedbackSignals`
  - `GetRecentActivities`
- **AND** 必须继续承接 `DashboardOverview`、`FeedbackSignal`、`ProductAssetCoverageSummary`、`RecentActivityItem`
- **AND** 不得额外并列增加趋势分析、通知中心、导出、批量操作或写入 RPC

#### Scenario: 编号与演进规则落地

- **WHEN** `dashboard.proto` 进入仓库主线
- **THEN** 当前版本字段编号、枚举编号、`UNSPECIFIED = 0`、`reserved` 与 breaking 演进规则必须继续对齐 `phase05-08`
- **AND** 不得在实现阶段自行重排字段编号、修改枚举顺序或回收已删除字段号
- **AND** 不得把 `missing both bindings` 回退为两个单缺口的隐式组合语义

#### Scenario: import 与时间字段落地

- **WHEN** 落地 `RecentActivityItem.activity_at`
- **THEN** `dashboard.proto` 必须显式使用 `google.protobuf.Timestamp`
- **AND** 该 import 必须在同一个 `proto/` workspace 下通过 `buf build` 成功解析
- **AND** 不得改用字符串时间字段绕开 `.proto` 时间语义

### Requirement: buf 工具链入口必须可运行且复用现有入口

系统 SHALL 为 `Dashboard + Feedback` 合同落地继续复用并扩展当前仓库已有的 `buf` 工具链入口，使 `build / lint / generate / breaking` 可以在同一 proto workspace 中执行，而不是为单个模块发明第二套脚本体系。

#### Scenario: build lint generate 入口

- **WHEN** 后续实现者在 `proto/` 目录校验 `Dashboard + Feedback` 合同
- **THEN** 必须能够通过既有入口运行 `buf build`、`buf lint` 与 `buf generate`
- **AND** 这些入口必须同时覆盖 `Module Registry`、`Decision Center`、`Product Registry`、`Repository Binding` 与 `Dashboard + Feedback`
- **AND** 不得要求实现者手工拼接单文件命令绕过 `proto/buf.yaml` 或 `proto/buf.gen.yaml`

#### Scenario: breaking 基准路径

- **WHEN** 后续实现者或 CI 对 `Dashboard + Feedback` 合同执行破坏性变更校验
- **THEN** 必须通过既有 `proto/Makefile` 或等价受控入口运行 `buf breaking`
- **AND** `buf breaking` 必须直接对照仓库主线 Git 基准，路径口径与 `proto/` 子目录保持一致
- **AND** `buf breaking` 的基准路径必须继续冻结为 `../.git#branch=main,subdir=proto`
- **AND** 失败时必须保留非零退出码，不得吞掉错误
- **AND** 不得改为对临时导出文件、临时副本目录或手工拼接镜像做 breaking 基准

### Requirement: 生成产物落点必须与现有 proto 主线同构

系统 SHALL 要求 `Dashboard + Feedback` 的代码生成产物继续落在现有 Go / TypeScript 生成目录主线上，使后续后端与前端实现可以复用既有生成模式。

#### Scenario: Go 与 TypeScript 生成产物

- **WHEN** 对 `Dashboard + Feedback` 合同执行 `buf generate`
- **THEN** Go 生成产物必须落在 `backend/internal/gen/proto/psco/dashboard/v1/`
- **AND** TypeScript 生成产物必须落在 `frontend/src/gen/proto/psco/dashboard/v1/`
- **AND** 当前阶段只要求生成消息类型与 service 定义对应的最小合同产物
- **AND** 不得因为 `phase05-11` 额外引入完整 gRPC 服务端、Connect 网关或第二套客户端生成主线

### Requirement: proto/README.md 必须把 Dashboard 纳入单一合同源总览

系统 SHALL 要求 `proto/README.md` 在 `phase05-11` 完成后继续承担仓库 proto 合同源总览入口，并把 `Dashboard + Feedback` 纳入其中，而不是让 `dashboard.proto` 成为 README 之外的孤立合同文件。

#### Scenario: 目录总览更新

- **WHEN** `phase05-11` 更新 `proto/README.md`
- **THEN** 目录结构、包名与版本语义表必须新增 `psco/dashboard/v1/dashboard.proto`
- **AND** 生成产物落点表必须新增 `backend/internal/gen/proto/psco/dashboard/v1/` 与 `frontend/src/gen/proto/psco/dashboard/v1/`
- **AND** 不得遗漏 `Dashboard + Feedback`，导致 README 与真实 proto 主线脱节

#### Scenario: RPC → HTTP 映射总览更新

- **WHEN** `proto/README.md` 继续维护过渡传输层映射总览
- **THEN** 必须新增 `DashboardService` 三个 RPC 的 HTTP 映射：
  - `GetDashboardOverview` → `GET /api/dashboard/overview`
  - `GetFeedbackSignals` → `GET /api/dashboard/feedback-signals`
  - `GetRecentActivities` → `GET /api/dashboard/recent-activities`
- **AND** 必须明确这三类读取当前阶段无 body、无 query 过滤、无路径参数
- **AND** 不得把 HTTP 路径、状态码或中间件策略误写成 `.proto` 合同本体

### Requirement: Dashboard 过渡传输层必须从 proto 单向承接

系统 SHALL 冻结 `Dashboard + Feedback` 的 DTO / HTTP adapter / handler 与 `.proto` 的关系为“`.proto` 单向定义合同，过渡传输层显式映射承接”，不允许继续并列扩张第二套字段语义。

#### Scenario: Dashboard DTO 与 adapter 映射边界

- **WHEN** 后续实现者编写 `backend/internal/dashboard/types.go`、HTTP handler DTO 或前端 adapter
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 显式对齐
- **AND** 不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义
- **AND** 时间格式、枚举字符串形式、局部错误展示状态等传输差异必须被明确视为适配层差异，而不是第二套合同定义

#### Scenario: Dashboard GET 入口与 proto request 组装

- **WHEN** HTTP 过渡层承接 `GetDashboardOverview`、`GetFeedbackSignals`、`GetRecentActivities`
- **THEN** handler 必须显式组装对应的空 Proto request 消息，再进入业务层
- **AND** 当前阶段不得因为 HTTP 侧没有 body、query 或路径参数，就跳过 Proto request 这一合同边界
- **AND** 不得因为 `GET` 无参数而在 `.proto` 之外保留另一份请求合同

#### Scenario: 错误语义与合同边界

- **WHEN** 过渡传输层承接 `DashboardOverviewRead` 整页失败、`FeedbackSignalRead` 局部失败或 `RecentActivityRead` 局部失败
- **THEN** 错误状态码、错误响应包络与局部错误展示语义继续属于 HTTP / handler 适配层
- **AND** `.proto` 只承接成功路径下的合同结构，不并列发明第二套错误业务字段
- **AND** 不得为了表达局部失败而向 `DashboardOverview`、`GetFeedbackSignalsResponse` 或 `GetRecentActivitiesResponse` 私自新增 `.proto` 中未冻结的错误标记字段

### Requirement: 当前阶段合同落地边界必须保持最小化

系统 SHALL 明确 `phase05-11` 的目标是“合同源落地 + 工具链入口 + 映射边界”，而不是把整个 `Dashboard + Feedback` 传输栈一次性改写完成。

#### Scenario: 当前阶段允许保留的实现边界

- **WHEN** 执行 `phase05-11`
- **THEN** 当前阶段可以继续保留 `chi + JSON HTTP` 作为过渡传输层
- **AND** 当前阶段可以只生成消息类型和最小合同产物
- **AND** 当前阶段不要求完成完整 gRPC / Connect 传输层迁移
- **AND** 当前阶段不要求立即用生成类型替换全部现有手写 DTO

## MODIFIED Requirements

### Requirement: phase05 Contract First 的 Dashboard 合同解释从“设计冻结”推进为“仓库主线落地”

系统 SHALL 将 `phase05-08` 与 `phase05-10` 中已经冻结的 `Dashboard + Feedback` Proto 设计，从“规格层定义”推进为“仓库内实际存在且可被 buf 工具链消费的单一合同源”。

#### Scenario: 合同阶段推进

- **WHEN** `phase05-11` 开始执行
- **THEN** `Dashboard + Feedback` `.proto` 不再只停留在 `phase05-08` 与 `phase05-10` 文档正文中
- **AND** 必须在仓库 `proto/` workspace 内拥有实际文件落点、生成入口与校验入口
- **AND** 后续 `phase05-12 / 13 / 14` 必须优先引用该已落地合同源，而不是回到文档层手工解释字段

### Requirement: buf 校验链从“基线要求”推进为“仓库执行入口”

系统 SHALL 将 `phase05` 共享基线与正式规格正文中关于 `buf build / lint / generate / breaking` 的要求，从抽象校验前提推进为仓库中的受控执行入口。

#### Scenario: 工具链收口

- **WHEN** `phase05-11` 完成
- **THEN** `buf` 校验链必须能够在仓库中通过受控入口执行
- **AND** 这些入口必须与现有 `proto/` workspace 保持单一真相源
- **AND** 不得继续保留“工具链要求存在于文档中，但仓库内没有对应执行入口”的状态

## REMOVED Requirements

### Requirement: Dashboard 合同继续停留在 DTO / handler / 页面模型层

**Reason**: `phase05-11` 的目标就是把 `Dashboard + Feedback` 合同从“已经设计好”推进到“仓库中已经落地、可生成、可校验、可被实现引用”的状态。若继续让 DTO、页面类型或 handler 结构承担并列合同职责，后续实现会再次长出第二套语义。

**Migration**: 后续 `Dashboard + Feedback` 的实现、联调与验收应统一从 `proto/psco/dashboard/v1/dashboard.proto` 与现有 `proto/` 工具链入口进入；`phase05-08` 与 `phase05-10` 继续作为设计与正式规格上游，不再承担仓库内合同主线入口职责。
