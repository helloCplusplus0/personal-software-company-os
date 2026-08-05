# phase02-11A Module Registry 最小 Protocol Buffers 合同主线 Spec

## Why

`phase02-09` 已经把 `Module Registry` 的页面动作、读写模型与接口分组冻结为单值结论，但当前仓库仍缺少可作为跨语言单一合同源的 `.proto` 定义。若继续沿用仅有 `chi + JSON HTTP` 的过渡实现，后续随着前后端与模块数量增加，再回补 `Protocol Buffers` 的字段编号、消息边界与演进约束，成本会明显放大。

## What Changes

- 为 `Module Registry` 新增最小 `.proto` 合同源，覆盖当前阶段全部必须动作与候选读取
- 冻结 `package`、版本语义、消息结构、字段编号、枚举和服务接口
- 冻结 `.proto` 与 `chi + JSON HTTP` 过渡传输层之间的单向承接关系
- 明确 `Decision` 仍为 `ModuleDetailRead` 的内嵌附属读取，不新增独立 Proto 服务
- **BREAKING**：后续实现与验收不得再把手写 JSON 结构视为并列合同源，`.proto` 成为当前阶段唯一合同源

## Impact

- Affected specs:
  - `phase02-09 module_registry_spec_v0.1`
  - `phase02-11 backend_data_mainline`
  - `phase02-12` 联调验收入口
- Affected code:
  - 预期新增 `proto/psco/module_registry/v1/*.proto`
  - 预期补充 Go / TS 合同承接与生成脚本入口
  - 预期约束 `backend/internal/moduleregistry/types.go`
  - 预期约束 `frontend/src/features/module-registry/data/api-adapter.ts`

## ADDED Requirements

### Requirement: Module Registry 最小 Proto 合同源

系统 SHALL 为 `Module Registry` 落地单一 `.proto` 合同源，作为当前阶段唯一合同定义入口。

#### Scenario: 合同源落地

- **WHEN** 实现 `phase02-11A`
- **THEN** 仓库中必须存在可追踪的 `.proto` 文件落点
- **AND** 该 `.proto` 文件必须覆盖 `ModuleListRead / ModuleDetailRead / ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite / Candidate Read`
- **AND** 不得再以手写 JSON 结构充当并列合同源

### Requirement: Proto 包名与版本语义

系统 SHALL 为当前阶段 Proto 合同冻结明确的包名、版本号与文件组织方式。

#### Scenario: 版本与目录冻结

- **WHEN** 新增 `.proto` 合同源
- **THEN** 必须冻结单一包名与版本语义（如 `psco.module_registry.v1`）
- **AND** 必须冻结稳定的目录落点
- **AND** 后续新增字段必须在该版本演进规则下进行，而不是临时改写包名

### Requirement: 最小消息结构覆盖当前动作矩阵

系统 SHALL 让 Proto 消息结构完整承接 `phase02-09` 已冻结的最小读写模型。

#### Scenario: 读组消息结构

- **WHEN** 定义 `ModuleListRead` 与 `ModuleDetailRead`
- **THEN** 必须覆盖列表项、详情对象、版本列表、产品绑定、仓库映射与 `Decision` 附属读取
- **AND** `Decision` 不得被拆成独立 Proto 读服务

#### Scenario: 写组消息结构

- **WHEN** 定义 `ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite`
- **THEN** 必须覆盖 `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository`
- **AND** 必须保持与 `phase02-09` 的最小字段语义一致

#### Scenario: 候选读取消息结构

- **WHEN** 定义 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead`
- **THEN** 只承接候选读取所需最小字段
- **AND** 不得借机把 `Product` 或 `Repository` 扩写为独立写入主线

### Requirement: 字段编号与兼容性规则

系统 SHALL 为 `.proto` 合同建立可演进的字段编号规则，避免后续演进破坏兼容性。

#### Scenario: 字段演进

- **WHEN** 为消息定义字段
- **THEN** 必须为每个字段分配稳定编号
- **AND** 不得复用已删除字段编号或字段语义
- **AND** 必须为未来保留字段演进空间

### Requirement: 过渡传输层与 Proto 的承接关系

系统 SHALL 明确当前阶段 `chi + JSON HTTP` 与 `.proto` 的关系，避免形成第二套合同源。

#### Scenario: 过渡层保留

- **WHEN** 当前阶段继续保留 `chi + JSON HTTP`
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 语义对齐（通过显式映射保持一致，非字段严格一致）
- **AND** 不得在 `handler` 或前端适配层私自新增 Proto 中不存在的字段语义
- **AND** 不得把 HTTP 路径、状态码或中间件策略误写成 Proto 合同本体

### Requirement: 生成链与实现边界

系统 SHALL 明确 `phase02-11A` 当前需要落地到什么程度，避免把最小合同主线膨胀成完整通信栈改造。

#### Scenario: MVP 阶段合同落地边界

- **WHEN** 执行 `phase02-11A`
- **THEN** 必须落地 `.proto` 合同源与最小生成入口约定
- **AND** 当前阶段可以不完成完整 gRPC / Connect / 网关迁移
- **AND** 当前阶段可以保留 `chi` 作为 HTTP 过渡传输层

## MODIFIED Requirements

### Requirement: Contract First 阶段要求

系统 SHALL 将 `Contract First` 从“长期方向冻结”推进为“当前阶段最小落地”，且落地点限定在 `Module Registry`。

#### Scenario: phase02 合同口径更新

- **WHEN** `phase02-11A` 进入执行链
- **THEN** `Module Registry` 的 `.proto` 必须成为当前阶段唯一合同源
- **AND** 后续 `phase02-12` 联调验收必须基于该合同源核对实现
- **AND** 不得继续沿用“当前阶段不要求 proto 落地”的旧口径

## REMOVED Requirements

### Requirement: 当前阶段仅保留 Protocol Buffers 长期方向

**Reason**: 该口径会导致实现继续围绕手写 JSON 合同增长，把 Proto 落地成本后移到更高阶段。
**Migration**: 改为“当前阶段必须落地 `Module Registry` 最小 `.proto` 合同源；完整传输层迁移可后续推进”。 
