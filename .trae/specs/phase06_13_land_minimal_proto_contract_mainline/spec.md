# Phase06-13 Onboarding + Sovereignty + Reuse 最小 Protocol Buffers 合同主线 Spec

## Why

`phase06-10` 已冻结 `Onboarding / Export / Backup / Reuse Summary` 的最小 `.proto` 合同设计，`phase06-12` 又将这些结论收口为当前阶段唯一正式规格正文。但截至现在，仓库主线里仍没有 `phase06` 对应的实际 `.proto` 文件、生成产物落点、`buf build / lint / generate / breaking` 入口覆盖，以及 HTTP DTO / adapter / handler 对 `.proto` 的单向承接主线。

如果继续围绕手写 DTO、handler JSON 结构或前端类型推进后续实现，就会再次长出与 `.proto` 并列的第二套合同源。`phase06-13` 的目标，就是把已经冻结的合同设计推进为仓库内真实存在、可生成、可校验、可被实现直接引用的合同主线。

## What Changes

- 将 `phase06-10` 已冻结的 `Onboarding / Export / Backup / Reuse Summary` 最小 `.proto` 合同正式落地到现有 `proto/` workspace
- 冻结 `proto/psco/onboarding/v1/onboarding.proto`、`proto/psco/export/v1/export.proto`、`proto/psco/backup/v1/backup.proto`、`proto/psco/reuse_summary/v1/reuse_summary.proto` 的单一文件落点、包名与服务归属
- 冻结 `phase06` draft-first 继续复用既有 `Product / Repository / Module / Decision` canonical create 合同，不得在 Onboarding 主线内并列新增第二套 create RPC
- 冻结 `buf build / lint / generate / breaking` 在当前仓库中的最小可运行入口，并要求继续复用 `proto/Makefile`、`proto/buf.yaml`、`proto/buf.gen.yaml`
- 冻结 `proto/README.md` 必须把 `phase06` 新增合同纳入单一合同源总览与 RPC → HTTP 映射矩阵
- 冻结 `backend/internal/*/types.go`、HTTP handler DTO、前端 `types.ts` / `api-adapter.ts` 与 `.proto` 的单向语义映射边界，阻断第二套合同源
- 明确当前阶段只落地合同源、生成入口、校验链与映射边界，不提前完成完整 gRPC / Connect 传输层迁移
- **BREAKING**：后续 `phase06` 的实现、联调与验收不得再把手写 JSON 结构、页面类型或 handler DTO 视为并列合同源，`.proto` 成为仓库内唯一合同定义入口

## Impact

- Affected specs:
  - `phase06_10_design_minimal_protocol_buffers_contract`
  - `phase06_12_onboarding_sovereignty_reuse_formal_spec`
  - 后续 `phase06` 源码实现、联调与验收任务
  - `phase03_11_decision_center_proto_mainline`
  - `phase04_11_product_repository_binding_proto_mainline`
  - `phase05_11_dashboard_feedback_proto_mainline`
- Affected code:
  - `proto/psco/onboarding/v1/onboarding.proto`
  - `proto/psco/export/v1/export.proto`
  - `proto/psco/backup/v1/backup.proto`
  - `proto/psco/reuse_summary/v1/reuse_summary.proto`
  - `proto/psco/product_registry/v1/product_registry.proto`
  - `proto/psco/repository_binding/v1/repository_binding.proto`
  - `proto/psco/module_registry/v1/module_registry.proto`
  - `proto/psco/decision_center/v1/decision_center.proto`
  - `proto/README.md`
  - `proto/Makefile`
  - `proto/buf.yaml`
  - `proto/buf.gen.yaml`
  - `backend/internal/gen/proto/psco/onboarding/v1/`
  - `backend/internal/gen/proto/psco/export/v1/`
  - `backend/internal/gen/proto/psco/backup/v1/`
  - `backend/internal/gen/proto/psco/reuse_summary/v1/`
  - `frontend/src/gen/proto/psco/onboarding/v1/`
  - `frontend/src/gen/proto/psco/export/v1/`
  - `frontend/src/gen/proto/psco/backup/v1/`
  - `frontend/src/gen/proto/psco/reuse_summary/v1/`
  - `backend/internal/onboarding/`、`export/`、`backup/`、`reusesummary/` 或等价模块
  - `frontend/src/features/onboarding/`、`export/`、`backup/`、`reuse-summary/`
  - 现有 `frontend/src/features/product-registry/`、`repository-binding/`、`module-registry/`、`decision-center/` 中承接 canonical create 的 adapter / types

## ADDED Requirements

### Requirement: Phase06 合同必须落地到现有 proto workspace

系统 SHALL 将 `phase06` 的最小 `.proto` 合同落地到现有 `proto/` workspace，而不是为 `Onboarding / Export / Backup / Reuse Summary` 新建第二个 proto 根目录、第二套 `buf.yaml` 或第二套生成入口。

#### Scenario: 合同源文件落点

- **WHEN** 执行 `phase06-13`
- **THEN** 仓库中必须存在：
  - `proto/psco/onboarding/v1/onboarding.proto`
  - `proto/psco/export/v1/export.proto`
  - `proto/psco/backup/v1/backup.proto`
  - `proto/psco/reuse_summary/v1/reuse_summary.proto`
- **AND** 上述文件必须作为 `phase06` 当前阶段唯一合同定义入口
- **AND** 必须继续复用现有 `proto/buf.yaml` 作为同一个 buf workspace 的根配置
- **AND** 不得在 `backend/`、`frontend/` 或仓库根新增并列 proto 工作区

#### Scenario: 包名与版本语义落地

- **WHEN** `phase06` 合同正式进入仓库
- **THEN** 包名与版本语义必须冻结为：
  - `psco.onboarding.v1`
  - `psco.export.v1`
  - `psco.backup.v1`
  - `psco.reuse_summary.v1`
- **AND** 后续 breaking 变更必须继续以 `v2` 作为主版本演进前提
- **AND** 不得在落地阶段临时改写为第二套包名、目录层级或版本策略

### Requirement: Draft-First 必须继续复用既有 canonical create 合同主线

系统 SHALL 要求 `phase06` 的 draft-first 写入在合同主线落地时继续复用既有 `Product / Repository / Module / Decision` 的 canonical create RPC，而不是在 `Onboarding` 主线内并列新增第二套 create 合同。

#### Scenario: Onboarding 合同归属不变

- **WHEN** 执行 `phase06-13`
- **THEN** `onboarding.proto` 当前阶段只允许承接 `OnboardingService.GetFirstRunState`
- **AND** 当前阶段不得在 `OnboardingService` 下新增 `CreateDraftProduct / CreateDraftRepository / CreateDraftModule / CreateDraftDecision`
- **AND** `ProductRegistryService.CreateProduct`、`RepositoryBindingService.CreateRepository`、`ModuleRegistryService.CreateModule`、`DecisionCenterService.CreateDecision` 必须继续作为四类 draft-first 写入的正式合同主线

#### Scenario: 既有 canonical create request 必须同步进入 phase06 主线

- **WHEN** `phase06-13` 将 draft-first 写入推进为仓库主线合同
- **THEN** `product_registry.proto`、`repository_binding.proto`、`module_registry.proto`、`decision_center.proto` 中对应的 `CreateProductRequest / CreateRepositoryRequest / CreateModuleRequest / CreateDecisionRequest` 必须同步对齐 `phase06-02` 与 `phase06-10` 已冻结的最小人工必填语义
- **AND** `CreateProductRequest` 必须允许仅由用户提供 `name`，并由系统语义承接 `description = ''`、`status = active`
- **AND** `CreateRepositoryRequest` 必须允许仅由用户提供 `name + url`，并由系统语义承接 `provider = manual`、`status = active`
- **AND** `CreateModuleRequest` 必须允许仅由用户提供 `name`，并由系统语义承接 `description = ''`、`status = active`
- **AND** `CreateDecisionRequest` 必须允许仅由用户提供 `title + choice + reason`，并由系统语义承接 `context = ''`、`problem = ''`、`impact = ''`、`alternatives = []`、`status = proposed`
- **AND** 当前阶段不得因为继续复用既有 canonical create 合同，就保留一套与 `phase06` draft-first 语义冲突的旧 request 必填口径

#### Scenario: HTTP 路由分组不扩写

- **WHEN** `phase06` 合同主线进入仓库
- **THEN** 当前阶段必须继续沿用以下既有 HTTP 映射：
  - `POST /api/products`
  - `POST /api/repositories`
  - `POST /api/modules`
  - `POST /api/decisions`
- **AND** 当前阶段不得发明 `/api/onboarding/drafts/*` 第二套路由分组
- **AND** `Onboarding` 页面只作为既有 canonical create 合同的消费入口，不成为新的写合同 owner

### Requirement: Phase06 新增 proto 文件必须完整承接 phase06-10 与 phase06-12 的最小服务矩阵

系统 SHALL 要求 `phase06` 新增 `.proto` 文件在仓库落地时完整承接 `phase06-10` 与 `phase06-12` 已冻结的服务、消息、字段语义与演进规则，而不是在实现阶段重新发明“更适合当前代码”的第二套合同。

#### Scenario: onboarding.proto 服务与消息落地

- **WHEN** 落地 `proto/psco/onboarding/v1/onboarding.proto`
- **THEN** 文件内必须存在单一 `OnboardingService`
- **AND** 当前阶段只承接 `GetFirstRunState`
- **AND** `first_run_state` 响应至少必须承接 `status`、`is_first_entry`、`current_step`、`completion_progress`
- **AND** `status` 必须继续以枚举承接 `not_started / in_progress / completed`

#### Scenario: export.proto 服务与消息落地

- **WHEN** 落地 `proto/psco/export/v1/export.proto`
- **THEN** 文件内必须存在单一 `ExportService`
- **AND** 当前阶段最小必须承接：
  - `GetExportSnapshot`
  - `ExportCoreAssets`
- **AND** `ExportSnapshot` 必须继续承接 `asset_scope`、`created_at`、`result_status`、`result_summary`
- **AND** `asset_scope` 必须能表达 `products / modules / releases / repositories / decisions / decision_links / product_modules / product_repositories / module_repositories`

#### Scenario: backup.proto 服务与消息落地

- **WHEN** 落地 `proto/psco/backup/v1/backup.proto`
- **THEN** 文件内必须存在单一 `BackupService`
- **AND** 当前阶段最小必须承接：
  - `GetBackupSnapshot`
  - `CreateInstanceBackup`
- **AND** `GetBackupSnapshot` 必须显式承担当前阶段 `read / verify` 子路径语义
- **AND** `BackupSnapshot` 必须继续承接 `created_at`、`manifest_summary`、`asset_coverage`、`schema_version_prerequisite`、`verified_status`
- **AND** 当前阶段不得把“写入响应里顺带返回一次 manifest”解释为已满足 snapshot 正式读取侧

#### Scenario: backup 失败语义必须继续保持单值化

- **WHEN** `phase06-13` 将 `Backup` 合同推进为 proto mainline
- **THEN** `GetBackupSnapshot` 的读取侧、HTTP DTO 映射侧与前端消费侧必须继续能单值区分以下失败家族：
  - `backup-manifest-missing`
  - `backup-coverage-incomplete`
  - `backup-schema-mismatch`
- **AND** `manifest` 缺失 / 不可解析、覆盖矩阵不完整、`schema / version` 前提不可校验三类失败语义不得被折叠为泛化的 `backup failed`
- **AND** 当前阶段不得把这些失败语义混写进 `BackupSnapshot` 成功字段本体
- **AND** 后续 `phase06` 验收必须仍能基于这三类单值失败语义验证 `phase06-11` 已冻结的 fixture 结果

#### Scenario: reuse_summary.proto 服务与消息落地

- **WHEN** 落地 `proto/psco/reuse_summary/v1/reuse_summary.proto`
- **THEN** 文件内必须存在单一 `ReuseSummaryService`
- **AND** 当前阶段只承接 `GetReuseSummary`
- **AND** `GetReuseSummaryRequest` 必须继续承接 `scope`、`module_id`、`product_id`
- **AND** `scope` 必须继续以枚举承接 `dashboard / module_detail / product_detail`
- **AND** `GetReuseSummaryResponse` 必须同时承接 `module_reuse_summary[]` 与 `capability_summary[]`

#### Scenario: 演进规则落地

- **WHEN** `phase06` 合同正式进入仓库
- **THEN** 当前版本字段编号、枚举编号、`*_UNSPECIFIED = 0`、`reserved` 与 breaking 演进规则必须继续对齐 `phase06-10` 与 `phase06-12`
- **AND** 不得在落地阶段自行重排字段编号、修改枚举顺序、复用已删除 tag 或新增 `required`
- **AND** 删除字段或枚举值后必须保留对应 `reserved` 编号，必要时同时保留名称

### Requirement: buf 工具链入口必须可运行且复用现有入口

系统 SHALL 为 `phase06` 合同落地继续复用并扩展当前仓库已有的 `buf` 工具链入口，使 `build / lint / generate / breaking` 可以在同一 proto workspace 中执行，而不是为单个模块发明第二套脚本体系。

#### Scenario: build lint generate 入口

- **WHEN** 后续实现者在 `proto/` 目录校验 `phase06` 合同
- **THEN** 必须能够通过既有入口运行 `buf build`、`buf lint` 与 `buf generate`
- **AND** 这些入口必须同时覆盖现有 `Module Registry / Decision Center / Product Registry / Repository Binding / Dashboard` 以及本阶段新增 `Onboarding / Export / Backup / Reuse Summary`
- **AND** 不得要求实现者手工拼接单文件命令绕过 `proto/buf.yaml` 或 `proto/buf.gen.yaml`

#### Scenario: breaking 基准路径

- **WHEN** 后续实现者或 CI 对 `phase06` 合同执行破坏性变更校验
- **THEN** 必须通过既有 `proto/Makefile` 或等价受控入口运行 `buf breaking`
- **AND** `buf breaking` 必须直接对照仓库主线 Git 基准，路径口径与 `proto/` 子目录保持一致
- **AND** `buf breaking` 的基准路径必须继续冻结为 `../.git#branch=main,subdir=proto`
- **AND** 失败时必须保留非零退出码，不得吞掉错误
- **AND** 不得改为对临时导出文件、临时副本目录或手工拼接镜像做 breaking 基准

### Requirement: 生成产物落点必须与现有 proto 主线同构

系统 SHALL 要求 `phase06` 的代码生成产物继续落在现有 Go / TypeScript 生成目录主线上，使后续后端与前端实现可以复用既有生成模式。

#### Scenario: Go 与 TypeScript 生成产物

- **WHEN** 对 `phase06` 合同执行 `buf generate`
- **THEN** Go 生成产物必须落在：
  - `backend/internal/gen/proto/psco/onboarding/v1/`
  - `backend/internal/gen/proto/psco/export/v1/`
  - `backend/internal/gen/proto/psco/backup/v1/`
  - `backend/internal/gen/proto/psco/reuse_summary/v1/`
- **AND** TypeScript 生成产物必须落在：
  - `frontend/src/gen/proto/psco/onboarding/v1/`
  - `frontend/src/gen/proto/psco/export/v1/`
  - `frontend/src/gen/proto/psco/backup/v1/`
  - `frontend/src/gen/proto/psco/reuse_summary/v1/`
- **AND** 当前阶段只要求生成消息类型与 service 定义对应的最小合同产物
- **AND** 不得因为 `phase06-13` 额外引入完整 gRPC 服务端、Connect 网关或第二套客户端生成主线

### Requirement: proto/README.md 必须把 Phase06 纳入单一合同源总览

系统 SHALL 要求 `proto/README.md` 在 `phase06-13` 完成后继续承担仓库 proto 合同源总览入口，并把 `Onboarding / Export / Backup / Reuse Summary` 纳入其中，而不是让新增 `.proto` 成为 README 之外的孤立合同文件。

#### Scenario: 目录总览更新

- **WHEN** `phase06-13` 更新 `proto/README.md`
- **THEN** 目录结构、包名与版本语义表必须新增：
  - `psco/onboarding/v1/onboarding.proto`
  - `psco/export/v1/export.proto`
  - `psco/backup/v1/backup.proto`
  - `psco/reuse_summary/v1/reuse_summary.proto`
- **AND** 生成产物落点表必须同步新增对应的 Go / TypeScript 输出目录
- **AND** 不得遗漏 `phase06` 新增合同，导致 README 与真实 proto 主线脱节

#### Scenario: RPC → HTTP 映射总览更新

- **WHEN** `proto/README.md` 继续维护过渡传输层映射总览
- **THEN** 必须新增以下映射：
  - `GetFirstRunState` → `GET /api/onboarding/state`
  - `GetExportSnapshot` → `GET /api/dashboard/export`
  - `ExportCoreAssets` → `POST /api/dashboard/export`
  - `GetBackupSnapshot` → `GET /api/dashboard/backup`
  - `CreateInstanceBackup` → `POST /api/dashboard/backup`
  - `GetReuseSummary` → `GET /api/reuse-summary`
- **AND** 必须同时标注 `phase06` draft-first 继续复用既有 canonical create HTTP 映射，而不是由 `Onboarding` 新增专属 create 路由
- **AND** 不得把 HTTP 路径、状态码或中间件策略误写成 `.proto` 合同本体

### Requirement: Phase06 过渡传输层必须从 proto 单向承接

系统 SHALL 冻结 `phase06` 的 DTO / HTTP adapter / handler 与 `.proto` 的关系为“`.proto` 单向定义合同，过渡传输层显式映射承接”，不允许继续并列扩张第二套字段语义。

#### Scenario: DTO 与 adapter 映射边界

- **WHEN** 后续实现者编写 `backend/internal/onboarding/`、`export/`、`backup/`、`reusesummary/` 中的 `types.go`、HTTP handler DTO 或前端 `types.ts` / `api-adapter.ts`
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 显式对齐
- **AND** 不得在 `types.go`、handler DTO、前端页面类型或 adapter 层私自新增 `.proto` 中不存在的业务字段语义
- **AND** 时间格式、枚举字符串形式、局部错误展示状态等传输差异必须被明确视为适配层差异，而不是第二套合同定义

#### Scenario: HTTP request 到 proto request 的显式组装

- **WHEN** HTTP 过渡层承接 `GetFirstRunState`、`GetExportSnapshot`、`ExportCoreAssets`、`GetBackupSnapshot`、`CreateInstanceBackup`、`GetReuseSummary`
- **THEN** handler 必须在进入业务层前显式组装对应的 Proto request 消息
- **AND** `GetReuseSummary` 的 `scope / module_id / product_id` 必须通过 query 参数映射到 Proto request
- **AND** 当前阶段不得因为某些 `GET` 入口无 body 或 path 参数，就绕过 Proto request 这一合同边界

#### Scenario: backup_snapshot 读取侧合同一致性

- **WHEN** 过渡传输层承接 `GetBackupSnapshot`
- **THEN** 该读取侧必须继续作为 `BackupWrite.read_verify` 或等价上游冻结承接位的正式读时合同存在
- **AND** `manifest`、覆盖矩阵与 `schema / version` 前提字段必须从 `.proto` 单向承接到 HTTP DTO 与前端消费模型
- **AND** 当前阶段不得以 `CreateInstanceBackup` 写入响应附带的临时结果代替 `GetBackupSnapshot` 的正式读取合同

### Requirement: 当前阶段合同落地边界必须保持最小化

系统 SHALL 明确 `phase06-13` 的目标是“合同源落地 + 工具链入口 + 映射边界”，而不是把整个 `phase06` 传输栈一次性改写完成。

#### Scenario: 当前阶段允许保留的实现边界

- **WHEN** 执行 `phase06-13`
- **THEN** 当前阶段可以继续保留 `chi + JSON HTTP` 作为过渡传输层
- **AND** 当前阶段可以只生成消息类型和最小合同产物
- **AND** 当前阶段不要求完成完整 gRPC / Connect 传输层迁移
- **AND** 当前阶段不要求立即用生成类型替换全部现有手写 DTO

## MODIFIED Requirements

### Requirement: Phase06 Proto 合同源从“设计冻结”推进为“仓库主线落地”

系统 SHALL 将 `phase06-10` 与 `phase06-12` 中已经冻结的 `Onboarding + Sovereignty + Reuse` Proto 设计，从“规格层定义”推进为“仓库内实际存在且可被 buf 工具链消费的单一合同源”。

#### Scenario: 合同阶段推进

- **WHEN** `phase06-13` 开始执行
- **THEN** `phase06` `.proto` 不再只停留在 `phase06-10` 与 `phase06-12` 文档正文中
- **AND** 必须在仓库 `proto/` workspace 内拥有实际文件落点、生成入口与校验入口
- **AND** 后续 `phase06` 的实现、联调与验收必须优先引用该已落地合同源，而不是回到文档层手工解释字段

### Requirement: buf 校验链从“基线要求”推进为“仓库执行入口”

系统 SHALL 将 `phase06` 共享基线与正式规格正文中关于 `buf build / lint / generate / breaking` 的要求，从抽象校验前提推进为仓库中的受控执行入口。

#### Scenario: 工具链收口

- **WHEN** `phase06-13` 完成
- **THEN** `buf` 校验链必须能够在仓库中通过受控入口执行
- **AND** 这些入口必须与现有 `proto/` workspace 保持单一真相源
- **AND** 不得继续保留“工具链要求存在于文档中，但仓库内没有对应执行入口”的状态

## REMOVED Requirements

### Requirement: Phase06 合同继续停留在 formal spec、子规格或 DTO 层

**Reason**: `phase06-13` 的目标就是把 `Onboarding + Sovereignty + Reuse` 合同从“已经设计好”推进到“仓库中已经落地、可生成、可校验、可被实现引用”的状态。若继续让 DTO、页面类型或 handler 结构承担并列合同职责，后续实现会再次长出第二套语义。

**Migration**: 后续 `phase06` 的实现、联调与验收应统一从 `proto/psco/onboarding/v1/onboarding.proto`、`proto/psco/export/v1/export.proto`、`proto/psco/backup/v1/backup.proto`、`proto/psco/reuse_summary/v1/reuse_summary.proto` 与现有 `proto/` 工具链入口进入；`phase06-10` 与 `phase06-12` 继续作为设计与正式规格上游，不再承担仓库内合同主线入口职责。
