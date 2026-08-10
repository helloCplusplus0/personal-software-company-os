# Phase06-10 当前阶段最小 Protocol Buffers 合同设计 Spec

## Why

`phase06-01 / 02 / 03 / 04 / 05 / 08 / 09` 已分别冻结了 Onboarding、draft-first、Export / Backup、Reuse Summary、合同约束与后端边界，但当前仍缺少一份把这些语义统一收进 `.proto` 设计层的单值合同结论。若不先冻结 `OnboardingRead / Export / Backup / Reuse Summary` 的最小文件分组、包名版本、消息结构、服务接口、错误语义，以及 draft-first 对既有 canonical create 合同的影响与 `.proto -> DTO / HTTP JSON` 映射策略，后续 `phase06-12` 正式规格正文与 `phase06-13` 合同主线落地仍会在服务拆分、字段组合和 HTTP 形状上各自长出不同解释。

## What Changes

- 冻结 `Onboarding / Export / Backup / Reuse Summary` 四组最小 `.proto` 合同的文件分组、包名与版本语义
- 冻结四组服务的最小 RPC 矩阵、请求/响应消息边界与必要枚举
- 冻结 draft-first 继续复用既有 `CreateProduct / CreateRepository / CreateModule / CreateDecision` canonical 合同，而不是长出第二套 `Onboarding` create contract
- 冻结 `first_run_state`、`export_snapshot`、`backup_snapshot`、`module_reuse_summary / capability_summary` 的 Proto 承接方式
- 冻结 `.proto -> HTTP DTO / JSON` 的单向映射策略与最小 HTTP 入口分组
- 冻结当前阶段的错误语义、枚举零值、保留字段与 breaking 演进前提
- 明确当前 `/spec` 只负责合同设计冻结，不在同一轮同时要求真实 `.proto`、生成产物与源码实现全部落地；但真实 `.proto` 合同进入主线时点不得晚于 `phase06-12` 正式规格正文

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-01` Onboarding 边界与入口
  - `phase06-02` draft-first / partial-entry 写路径
  - `phase06-03` Export / Backup / restore prerequisites 语义
  - `phase06-04` Reuse Summary 读模型与挂接位
  - `phase06-05` 合同、传输与源码约束
  - `phase06-08` Export / Backup 后端模块边界
  - `phase06-09` Reuse Summary query / Dashboard / Detail 集成
- Affected code:
  - 后续 `proto/psco/onboarding/v1/onboarding.proto`
  - 后续 `proto/psco/export/v1/export.proto`
  - 后续 `proto/psco/backup/v1/backup.proto`
  - 后续 `proto/psco/reuse_summary/v1/reuse_summary.proto`
  - 后续 `proto/psco/product_registry/v1/product_registry.proto`
  - 后续 `proto/psco/repository_binding/v1/repository_binding.proto`
  - 后续 `proto/psco/module_registry/v1/module_registry.proto`
  - 后续 `proto/psco/decision_center/v1/decision_center.proto`
  - 后续 `proto/README.md`
  - 后续 `backend/internal/*/types.go`、handler DTO 与 `frontend/src/features/*/data/api-adapter.ts`
  - 后续 `phase06-13` 的 `buf build / lint / generate / breaking` 主线落地

## ADDED Requirements

### Requirement: Phase06 Proto 文件分组、包名与版本语义冻结

系统 SHALL 将 `phase06` 当前阶段最小合同设计冻结为四个逻辑 Proto 文件组，而不是把 Onboarding、Export、Backup、Reuse Summary 混写进既有模块合同或拆成多套平行文件。

#### Scenario: 合同文件分组

- **WHEN** 接手者设计 `phase06` 的最小 `.proto` 合同
- **THEN** 必须冻结为以下四个逻辑文件组：
  - `proto/psco/onboarding/v1/onboarding.proto`
  - `proto/psco/export/v1/export.proto`
  - `proto/psco/backup/v1/backup.proto`
  - `proto/psco/reuse_summary/v1/reuse_summary.proto`
- **AND** 当前子任务中这些路径只作为正式设计落点
- **AND** 当前子任务不得要求真实文件已经写入仓库主线

#### Scenario: 包名与版本语义

- **WHEN** 接手者冻结四组合同的 package 与 version
- **THEN** 必须分别冻结为：
  - `psco.onboarding.v1`
  - `psco.export.v1`
  - `psco.backup.v1`
  - `psco.reuse_summary.v1`
- **AND** 当前阶段 breaking 变更必须以 `v2` 作为后续主版本演进前提
- **AND** 当前阶段不得在同一业务能力下并列发明第二套包名或目录层级

### Requirement: OnboardingService 最小合同设计冻结

系统 SHALL 将 Onboarding 的最小 `.proto` 合同冻结为一个只承接 `first_run_state` 读取的 `OnboardingService`，而不是并列承接四类 canonical 草稿创建动作。

#### Scenario: Onboarding RPC 矩阵

- **WHEN** 接手者设计 Onboarding 的服务接口
- **THEN** 必须冻结单一 `OnboardingService`
- **AND** 该 service 当前阶段只承接 `GetFirstRunState`
- **AND** 当前阶段不得把 `CreateDraftProduct / CreateDraftRepository / CreateDraftModule / CreateDraftDecision` 再挂到 `OnboardingService` 下形成第二套并列创建合同

#### Scenario: `first_run_state` 消息边界

- **WHEN** 接手者设计 `GetFirstRunState` 的响应消息
- **THEN** `first_run_state` 至少必须承接：
  - `status`
  - `is_first_entry`
  - `current_step`
  - `completion_progress`
- **AND** `status` 必须以枚举承接 `not_started / in_progress / completed`
- **AND** `current_step` 必须显式承接当前引导步骤，而不是要求前端自行推导

### Requirement: Draft-First 继续复用既有 canonical create 合同

系统 SHALL 将 `phase06` 的四类 draft-first 创建动作冻结为“继续复用既有 canonical create contract”，而不是新增 `OnboardingService.CreateDraft*` 及其专属 HTTP 路由。

#### Scenario: canonical service 归属不变

- **WHEN** 接手者为 `Product / Repository / Module / Decision` 设计 phase06 的 draft-first 合同
- **THEN** 必须继续复用以下既有 service 与 create RPC 主线：
  - `ProductRegistryService.CreateProduct`
  - `RepositoryBindingService.CreateRepository`
  - `ModuleRegistryService.CreateModule`
  - `DecisionCenterService.CreateDecision`
- **AND** 当前阶段不得为相同四类创建语义再长出 `OnboardingService.CreateDraft*` 第二套 canonical 写合同

#### Scenario: Draft-First request 最小字段

- **WHEN** 接手者设计 phase06 的四类 draft-first create request
- **THEN** 必须继续复用既有 `CreateProductRequest / CreateRepositoryRequest / CreateModuleRequest / CreateDecisionRequest`
- **AND** 这些既有 request 在 phase06 中必须继续对齐 `phase06-02` 已冻结的最小人工必填字段：
  - `CreateProductRequest` 只把 `name` 作为人工必填
  - `CreateRepositoryRequest` 只把 `name + url` 作为人工必填
  - `CreateModuleRequest` 只把 `name` 作为人工必填
  - `CreateDecisionRequest` 只把 `title + choice + reason` 作为人工必填
- **AND** 当前阶段不得把 `description / provider / capability_key / target_link` 等后补字段重新拉回 request 必填
- **AND** 既有 create request 中其余字段若仍需存在，必须以系统默认值、可选语义或等价非人工阻断方式承接，而不是要求新建一套 `CreateDraft*Request`

#### Scenario: Draft-First response 边界

- **WHEN** 接手者设计 phase06 的四类 draft-first create response
- **THEN** 必须继续复用各 canonical create 的既有 response 形状
- **AND** 当前阶段不得要求四类 create response 内联更新后的 `first_run_state`
- **AND** Onboarding 进度推进必须通过后续 `GetFirstRunState` 重新读取或等价正式读路径承接
- **AND** 当前阶段不得要求 create response 内联完整 canonical detail 读模型

### Requirement: ExportService 最小合同设计冻结

系统 SHALL 将 Export 的最小 `.proto` 合同冻结为围绕 `export_snapshot` 的单一 `ExportService`，同时承接快照读取与导出执行。

#### Scenario: Export RPC 矩阵

- **WHEN** 接手者设计 Export 的服务接口
- **THEN** 必须冻结单一 `ExportService`
- **AND** 该 service 最小必须承接：
  - `GetExportSnapshot`
  - `ExportCoreAssets`
- **AND** 当前阶段不得把 Export 再拆为第二套“预览 service”与“执行 service”

#### Scenario: `export_snapshot` 消息边界

- **WHEN** 接手者设计 `ExportSnapshot`
- **THEN** 该消息至少必须承接：
  - `asset_scope`
  - `created_at`
  - `result_status`
  - `result_summary`
- **AND** `asset_scope` 必须能表达当前阶段 9 类核心资产覆盖矩阵
- **AND** 当前阶段不得在 `.proto` 中提前冻结导出文件格式、存储介质或下载介质细节

### Requirement: BackupService 最小合同设计冻结

系统 SHALL 将 Backup 的最小 `.proto` 合同冻结为围绕 `backup_snapshot` 的单一 `BackupService`，同时承接备份执行与独立读取校验。

#### Scenario: Backup RPC 矩阵

- **WHEN** 接手者设计 Backup 的服务接口
- **THEN** 必须冻结单一 `BackupService`
- **AND** 该 service 最小必须承接：
  - `GetBackupSnapshot`
  - `CreateInstanceBackup`
- **AND** `GetBackupSnapshot` 必须显式承担当前阶段 `read / verify` 子路径语义
- **AND** 当前阶段不得把“写入响应里顺带返回一次 manifest”解释为已满足 snapshot 正式读取侧

#### Scenario: `backup_snapshot` 消息边界

- **WHEN** 接手者设计 `BackupSnapshot`
- **THEN** 该消息至少必须承接：
  - `created_at`
  - `manifest_summary`
  - `asset_coverage`
  - `schema_version_prerequisite`
  - `verified_status`
- **AND** `verified_status` 必须能单值表达“未校验 / 已验证 / 校验失败”这类最小恢复前提状态
- **AND** 当前阶段不得把真正 restore 写回动作写入该 service 合同

### Requirement: ReuseSummaryService 最小合同设计冻结

系统 SHALL 将 Reuse Summary 的最小 `.proto` 合同冻结为一个单一 `ReuseSummaryService` 与单一 `GetReuseSummary` RPC，以支撑 Dashboard / Module Detail / Product Detail 每页一个页面级 query 的设计。

#### Scenario: Reuse Summary RPC 矩阵

- **WHEN** 接手者设计 Reuse Summary 的服务接口
- **THEN** 必须冻结单一 `ReuseSummaryService`
- **AND** 该 service 当前阶段只承接 `GetReuseSummary`
- **AND** 当前阶段不得为 Dashboard / Module Detail / Product Detail 各自再建第二套平行 RPC

#### Scenario: `GetReuseSummaryRequest` 作用域设计

- **WHEN** 接手者设计 `GetReuseSummaryRequest`
- **THEN** request 必须至少承接：
  - `scope`
  - `module_id`
  - `product_id`
- **AND** `scope` 必须以枚举承接 `dashboard / module_detail / product_detail`
- **AND** `module_id` 与 `product_id` 只允许按 scope 条件性使用

#### Scenario: `GetReuseSummaryResponse` 结果组合

- **WHEN** 接手者设计 `GetReuseSummaryResponse`
- **THEN** response 必须同时承接：
  - `module_reuse_summary[]`
  - `capability_summary[]`
- **AND** `module_reuse_summary` 与 `capability_summary` 的字段集合必须继续对齐 `phase06-04`
- **AND** `capability_summary` 在无可聚合数据时必须允许成功空数组或等价成功空态字段

### Requirement: `.proto -> HTTP DTO / JSON` 映射策略冻结

系统 SHALL 将 `phase06` 四组合同的 HTTP 过渡传输层映射策略冻结为“`.proto` 单向定义合同，HTTP DTO / JSON 只做显式映射承接”。

#### Scenario: Onboarding HTTP 映射

- **WHEN** 接手者设计 `OnboardingService` 的 HTTP 适配层
- **THEN** 当前阶段必须至少冻结以下映射：
  - `GetFirstRunState` → `GET /api/onboarding/state`
- **AND** HTTP body / query / path 参数只允许作为 Proto request 字段的传输来源

#### Scenario: Draft-First 沿用既有 canonical create HTTP 映射

- **WHEN** 接手者设计四类 draft-first 创建动作的 HTTP 适配层
- **THEN** 当前阶段必须继续冻结以下既有映射：
  - `CreateProduct` → `POST /api/products`
  - `CreateRepository` → `POST /api/repositories`
  - `CreateModule` → `POST /api/modules`
  - `CreateDecision` → `POST /api/decisions`
- **AND** 当前阶段不得为同一批创建语义再发明 `/api/onboarding/drafts/*` 第二套路由分组
- **AND** Onboarding 页面只作为这些既有 create 合同的消费入口，不成为新的写合同 owner

#### Scenario: Export / Backup / Reuse Summary HTTP 映射

- **WHEN** 接手者设计 `ExportService / BackupService / ReuseSummaryService` 的 HTTP 适配层
- **THEN** 当前阶段必须至少冻结以下映射：
  - `GetExportSnapshot` → `GET /api/dashboard/export`
  - `ExportCoreAssets` → `POST /api/dashboard/export`
  - `GetBackupSnapshot` → `GET /api/dashboard/backup`
  - `CreateInstanceBackup` → `POST /api/dashboard/backup`
  - `GetReuseSummary` → `GET /api/reuse-summary`
- **AND** `GetReuseSummary` 的 `scope / module_id / product_id` 必须通过 query 参数映射到 Proto request
- **AND** 当前阶段不得再为相同语义发明第二套路由分组

#### Scenario: DTO 单向派生

- **WHEN** 后续实现者编写 `types.go`、handler DTO、`types.ts` 与 `api-adapter.ts`
- **THEN** 所有字段语义必须从 `.proto` 单向派生或与 `.proto` 显式对齐
- **AND** 当前阶段不得在 HTTP DTO 或页面层私自新增 `.proto` 中不存在的业务字段语义

### Requirement: 错误语义与 breaking 演进前提冻结

系统 SHALL 将 `phase06` 四组合同的错误语义与 breaking 演进前提冻结为单值规则，而不是留给后续实现自由解释。

#### Scenario: 当前阶段错误语义

- **WHEN** 接手者冻结 `Onboarding / Export / Backup / Reuse Summary` 的错误语义
- **THEN** 必须至少明确以下最小失败家族：
  - Onboarding：状态读取失败
  - Draft-First（沿用既有 canonical create）：创建失败、最小必填不满足、系统默认值补齐后仍不合法
  - Export：资产装配失败、导出产物生成失败、导出结果不可读取
  - Backup：备份产物生成失败、manifest 缺失或不可解析、覆盖矩阵不完整、schema/version 前提不可校验
  - Reuse Summary：作用域参数非法、读取失败、作用域对象不存在
- **AND** 当前阶段不得把这些错误语义混写进成功 response 的业务字段中

#### Scenario: breaking 演进规则

- **WHEN** 接手者冻结四组 `.proto` 的演进前提
- **THEN** 必须继续遵守以下规则：
  - 枚举首值必须为 `*_UNSPECIFIED = 0`
  - 不得复用已删除字段编号
  - 删除字段或枚举值后必须保留 `reserved` 编号，必要时保留名称
  - 不得随意修改既有字段类型、重复性或语义
  - 不得新增 `required` 字段
- **AND** 当前阶段必须明确后续 breaking 变更需要通过新主版本或显式升级策略处理

### Requirement: `phase06-10` 非冻结边界

系统 SHALL 明确当前 `/spec` 轮次只负责合同设计冻结，但 `phase06` 后续推进中真实 `.proto` 合同进入主线的时点不得晚于 `phase06-12` 正式规格正文。

#### Scenario: 非当前阶段目标

- **WHEN** 接手者执行 `phase06-10`
- **THEN** 当前子任务不得要求：
  - 在本次 `/spec` 同一轮里真实 `.proto` 文件、`buf` 产物与源码实现同时全部完成
  - Go / TypeScript 生成产物已落地
  - handler / DTO / adapter 已完成源码实现
- **AND** `phase06-13` 可以继续承接 `buf build / lint / generate / breaking`、DTO 映射与产物收口
- **AND** 但真实 `.proto` 合同进入主线的时点不得晚于 `phase06-12` 正式规格正文收口

## MODIFIED Requirements

### Requirement: `chi + HTTP JSON` 的当前阶段合同职责

`chi + HTTP JSON` 在 `phase06-10` 中 SHALL 继续作为过渡传输层，但其职责必须从抽象“传输适配”推进为对四组正式 Proto service 与既有 canonical create contract 的显式 request / response 映射。

#### Scenario: 传输适配职责收口

- **WHEN** 接手者设计 `phase06` 的 HTTP 入口
- **THEN** 必须先以 `OnboardingService / ExportService / BackupService / ReuseSummaryService` 与既有 `CreateProduct / CreateRepository / CreateModule / CreateDecision` 的 Proto request / response 作为唯一合同前提
- **AND** HTTP 入口只承接参数来源、状态码和 JSON 序列化差异
- **AND** 不得再把“当前还在 HTTP JSON 阶段”解释为可以先不冻结服务接口与消息边界

## REMOVED Requirements

### Requirement: `phase06-10` 当前 `/spec` 必须在同一轮同步完成真实 proto 文件、生成产物与源码适配

**Reason**: 当前轮次是 `/spec` 设计冻结，不是实现落地；若要求在同一轮同时提交真实 `.proto` 文件、生成产物、DTO 与 handler 源码，会把规格冻结与实现收口重新混成一个步骤，破坏 `phase06` 的执行节奏。

**Migration**: `phase06-10` 先冻结 proto-ready 的文件分组、包名版本、RPC / message / error 语义、draft-first 与既有 canonical create 的关系，以及 HTTP 映射策略；后续执行中真实 `.proto` 合同进入主线的时点不得晚于 `phase06-12` 正式规格正文，`phase06-13` 再承接 `buf` 工具链、DTO 映射与生成产物收口。
