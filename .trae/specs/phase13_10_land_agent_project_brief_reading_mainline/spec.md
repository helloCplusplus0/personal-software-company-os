# Phase13-10 落实 agent 项目简报读取主线 Spec

## Why

`phase13-07` 已冻结 `project brief for agent` 的 7 顶层字段 schema、`repository_id` 唯一锚点、数组摘要协议与 `ProjectContextService` 受控演进关系；`phase13-08` 已交付治理画像后端写读主线；`phase13-09` 已交付前端治理画像承接与手工维护入口。但 agent 侧至今仍只能通过 `phase11/12` 的 `GetProjectContext`（含硬编码 `rules / phases / boundaries` 投影与 singular `product`）恢复上下文，无法读取以治理画像为准的正式项目简报。

本次 `/spec` 的目标，是把 `GetProjectBrief` 的合同落点、消息 schema、读取编排、失败语义与旧 RPC 兼容策略冻结为可直接实施的后端实现规格。

## What Changes

- 在现有 `ProjectContextService` 主线内新增 `GetProjectBrief` RPC（受控演进，不新建第二个 repository-scoped 只读聚合 service）
- 冻结 brief 的 proto 消息 schema（7 顶层字段，复用 `governance_profile.proto` 既有消息）
- 冻结 brief 读取编排与失败语义
- 冻结 `GetProjectContext / ExportProjectContext` 的兼容窗口与退役策略
- 完成 `product -> products[]` 与 `rules / phases / boundaries -> governance_profile / global_assets / current_phase` 的合同映射收口

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - `phase13-07`（brief schema 与 ProjectContextService 关系的直接下游实现）
  - `phase13-08`（治理画像后端主线，被 brief 聚合消费）
  - 后续 `phase13-11` 联调验收
- Affected code:
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/types.go`
  - `backend/internal/projectcontext/candidate/context_readers.go`
  - `backend/internal/projectcontext/service/query_service.go`
  - `backend/internal/projectcontext/connect/server.go`
  - `backend/internal/platform/router.go`（装配点注入治理画像读取依赖）
  - `backend/internal/governanceprofile/connect/server.go`（导出 domain→proto 转换函数供复用，避免第二套映射）

## ADDED Requirements

### Requirement: 冻结 brief 后端合同落点为 ProjectContextService 内新增 RPC

系统 SHALL 将 `project brief for agent` 的后端合同落点冻结为：在现有 `ProjectContextService`（`proto/psco/project_context/v1/project_context.proto`）内新增 `GetProjectBrief` RPC。

补充冻结：

1. 不得新建第二个 repository-scoped 只读聚合 service 或第二个 `.proto` 合同包
2. brief 消息定义落在 `project_context.proto` 内，通过 `import` 复用 `governance_profile.proto` 的 `GovernanceProfile / GlobalAssetBinding / PhaseStatus` 消息，保证后端、前端与 agent 消费侧共享同一份 schema
3. `GetProjectBrief` 只承接结构化只读读取，不承接任何写入
4. 第一版不新增 `ExportProjectBrief`、markdown 渲染、MCP、CLI 或 agent 写回入口

#### Scenario: 执行者决定 brief 合同落在哪里

- **WHEN** 执行者实现 `GetProjectBrief` 的后端合同
- **THEN** 必须落在现有 `project_context.proto` 的 `ProjectContextService` 内
- **AND** 不得另起第二个只读聚合 service 或第二份并列 `.proto` 合同源

### Requirement: 冻结 brief 的 proto 消息 schema

系统 SHALL 将 `GetProjectBrief` 的响应消息冻结为以下 7 个顶层字段：

1. `repository`（复用既有 `RepositorySummary`）
2. `governance_profile`（复用 `psco.governance_profile.v1.GovernanceProfile` 完整消息）
3. `global_assets`（`repeated psco.governance_profile.v1.GlobalAssetBinding`）
4. `current_phase`（新增最小消息：`name / entry_ref / status`，`status` 复用 `psco.governance_profile.v1.PhaseStatus` 枚举）
5. `products`（`repeated ProductSummary`，复用既有消息）
6. `modules`（`repeated ModuleSummary`，复用既有消息）
7. `decisions`（`repeated DecisionSummary`，复用既有消息）

补充冻结：

- 不得新增与上述 7 个顶层字段并列的第 8 个顶层块
- `governance_profile` 与 `global_assets` 数据同源于同一份治理画像读取结果；两者并存是 `phase13-07` 冻结的 schema 结构，不是第二套事实源
- `current_phase` 从治理画像主记录的 `current_phase_name / current_phase_ref / current_phase_status` 单向派生
- `products[] / modules[] / decisions[]` 即使长度为 1 也保持数组形式，不得退化为 singular schema
- brief 中不得混入任何"给 agent 的自然语言指导词"字段

#### Scenario: 执行者组装 brief 响应

- **WHEN** 执行者实现 `GetProjectBrief` 响应组装
- **THEN** 必须按上述 7 顶层字段填充
- **AND** 必须复用 `governance_profile.proto` 与既有 summary 消息，不得另造第二套字段语义

### Requirement: 冻结 brief 读取编排与失败语义

系统 SHALL 将 `GetProjectBrief` 的服务端编排冻结为：

1. 读取 Repository 身份；不存在 → `CodeNotFound`
2. 读取治理画像聚合（复用 `governanceprofile` 模块读取主线）；画像未创建 → `CodeNotFound`（复用既有 `ErrGovernanceProfileNotFound` 哨兵，语义与 `GetGovernanceProfile` 单值一致）
3. 读取 `products[]`（通过 `product_repositories`，数组语义）
4. 读取 `modules[]`（通过 `module_repositories`）
5. 读取 `decisions[]`（沿用两类 module-link 派生命中 + 去重 + archived 过滤）
6. `current_phase` 从步骤 2 的治理画像主记录派生
7. `global_assets` 从步骤 2 的 `global_asset_bindings` 同源填充

补充冻结：

- brief 不做 Repository Binding 完整性强制校验：`products[] / modules[] / decisions[]` 为空数组是合法状态（对齐 `phase13-07` "数组缺省语义"冻结），绑定不完整不返回 `FailedPrecondition`
- brief 数据全部来自 PSCO 结构化存储（治理画像表 + 四实体表）；`phase12` 遗留的硬编码 `rules / phases / boundaries` 投影不进入 brief
- 治理画像读取必须复用 `governanceprofile` 既有聚合读取能力（通过 candidate 接口注入），不得在 `projectcontext` 内复制第二套治理画像 SQL
- 其他读取失败 → `CodeInternal`

#### Scenario: 执行者实现 brief 聚合编排

- **WHEN** 执行者实现 `GetProjectBrief` 的 service 编排
- **THEN** 必须按上述顺序聚合且以 `repository_id` 为唯一锚点
- **AND** 不得伪造单一"当前 product / module / decision"
- **AND** 不得把硬编码规则投影或目录扫描结果混入 brief

### Requirement: 冻结 GetProjectContext / ExportProjectContext 的兼容与退役策略

系统 SHALL 将现有 `GetProjectContext / ExportProjectContext` 的处置冻结为：

1. 两个 RPC 在 `phase13` 内继续保留，作为兼容读取层与兼容导出层
2. proto 注释显式标注 deprecated 兼容窗口：它们不得再承接新的演进需求，不得与 brief 并列演化为两套 canonical agent 输入协议
3. 退役讨论进入条件：`phase13-11` 验收确认 brief 读取主线稳定后，退役窗口作为下一阶段进入条件讨论，不在 `phase13` 内直接移除
4. `phase13` 内不修改两个旧 RPC 的行为与响应结构（含其硬编码投影，作为兼容现状保留）

#### Scenario: 执行者处理旧 RPC

- **WHEN** 执行者新增 `GetProjectBrief`
- **THEN** 保留旧 RPC 行为不变并标注 deprecated 兼容注释
- **AND** 不得在 phase13 内删除或改写旧 RPC 响应结构

### Requirement: 冻结 brief 的跨模块读取承接方式

系统 SHALL 将 brief 对治理画像的读取承接冻结为：

1. `projectcontext/candidate` 子包定义 `GovernanceProfileReader` 接口（消费方拥有接口）
2. 平台装配点（`router.go`）注入 `governanceprofile/service.QueryService` 作为该接口的实现
3. `projectcontext` 不得直接书写治理画像 SQL，不得复制 governanceprofile 的存储读取逻辑

补充冻结：

- `products[]` 的数组读取由 `projectcontext/candidate` 新增 `ReadProducts`（数组版）承接；旧 `ReadProduct`（singular，LIMIT 1）保留仅供兼容层 `GetProjectContext` 使用
- domain → proto 组装复用 `governanceprofile/connect` 导出的转换函数（导出 `DomainResultToProto` 等既有私有转换），不得在 `projectcontext/connect` 重写第二套治理画像字段映射

#### Scenario: 执行者设计跨模块依赖

- **WHEN** 执行者实现 brief 的治理画像读取依赖
- **THEN** 必须通过 candidate 接口 + 平台装配注入
- **AND** 不得在 projectcontext 内复制治理画像存储 SQL 或字段映射逻辑

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的 agent 读取主线实现前提

`phase13_project_governance_profile_foundation` MUST 在 `phase13-10` 中完成 `GetProjectBrief` 的合同落点、消息 schema、读取编排、失败语义与旧 RPC 兼容策略的实现；若 brief 读取主线未落地，则 agent 仍无法在不目录扫描的前提下恢复 PSCO 正式上下文，`phase13` 不得视为收口。

## REMOVED Requirements

### Requirement: 允许新建第二个 repository-scoped 只读聚合 service 承接 brief

**Reason**: 这会直接违背 `phase13-07` 已冻结的 `ProjectContextService` 受控演进关系，让 repository-scoped 只读聚合合同长出两套长期并列的 canonical 主线。

**Migration**: brief 统一落在现有 `ProjectContextService` 内的 `GetProjectBrief` RPC；旧 `GetProjectContext / ExportProjectContext` 作为 deprecated 兼容层保留，退役窗口由后续阶段裁决。
