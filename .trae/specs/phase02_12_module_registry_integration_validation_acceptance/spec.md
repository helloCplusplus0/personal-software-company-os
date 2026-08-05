# phase02-12 Module Registry 联调、验证与验收 Spec

## Why

`phase02-10`、`phase02-11` 与 `phase02-11A` 已分别完成前端主线、后端与数据主线、最小 `.proto` 合同主线，但当前仍需要一份单独的联调验收规格，把“能实现”收口为“已被验证可运行”。如果缺少这一层，`phase02` 会停留在局部自证完成，无法形成可复核、可追踪、可交付的当前阶段运行证据。

## What Changes

- 为 `Module Registry` 新增一份联调、验证与验收 spec
- 冻结前后端联调环境、数据前提、合同核对基准与最小验收路径
- 冻结 `CreateModule / CreateRelease / BindModuleToProduct / MapModuleToRepository` 的正向与关键异常路径
- 冻结空状态、错误态、返回路径与 `.proto` 合同一致性的验收要求
- 明确发现的问题必须在当前阶段收口，不得带入 `phase02-13`

## Impact

- Affected specs:
  - `phase02-09 module_registry_spec_v0.1`
  - `phase02-10 frontend_module_registry_mainline`
  - `phase02-11 backend_data_mainline`
  - `phase02-11A module_registry_proto_contract`
- Affected code:
  - `frontend/src/features/module-registry/**`
  - `backend/internal/moduleregistry/**`
  - `database/migrations/**`
  - `database/seeds/**`
  - `proto/psco/module_registry/v1/module_registry.proto`

## ADDED Requirements

### Requirement: 联调环境必须可重复建立

系统 SHALL 为 `phase02-12` 提供可重复建立的联调环境前提，使验收不依赖一次性的临时状态。

#### Scenario: 环境初始化

- **WHEN** 执行 `phase02-12` 联调
- **THEN** 必须明确前端、后端、数据库与种子数据的启动顺序
- **AND** 必须复用 `phase02-11` 已冻结的独立数据库与初始化入口
- **AND** 必须明确前端当前处于真实后端模式，而不是 mock 模式

### Requirement: 最小主线必须端到端走通

系统 SHALL 证明 `Module Registry` 当前阶段最小主线已形成真实可运行交付物。

#### Scenario: CreateModule 最小闭环

- **WHEN** 用户从空状态或列表页进入 `Module Create`
- **THEN** 必须能够成功创建模块并回流到 `ModuleDetailPage`
- **AND** 回流后的详情页必须显示新建模块核心字段

#### Scenario: CreateRelease 最小闭环

- **WHEN** 用户在详情页进入 `Release Create`
- **THEN** 必须能够成功创建版本并回流到当前模块的详情页
- **AND** 详情页版本列表必须承接最新版本结果

#### Scenario: 绑定动作最小闭环

- **WHEN** 用户在详情页执行 `BindModuleToProduct` 或 `MapModuleToRepository`
- **THEN** 必须停留在当前详情页
- **AND** 必须重新读取并显示最新绑定结果

### Requirement: 空状态、错误态与返回路径必须被验证

系统 SHALL 为当前阶段的关键 UI 状态提供明确验收结果，避免只验证 happy path。

#### Scenario: 空状态验证

- **WHEN** 系统中尚无任何模块
- **THEN** `Module Registry / List` 必须显示围绕“先完成首个模块登记”的空状态入口
- **AND** 空状态主动作必须直接进入 `Module Create`

#### Scenario: 错误态验证

- **WHEN** 创建、版本登记或绑定动作失败
- **THEN** 错误必须停留在当前页面或当前面板上下文
- **AND** 不得跳转独立错误页
- **AND** 草稿或当前选择必须按 `phase02-09` 已冻结规则保留

#### Scenario: 返回路径验证

- **WHEN** 用户从创建页、详情页或版本登记页主动返回
- **THEN** 必须符合 `phase02-09` 已冻结的返回路径规则
- **AND** 从详情页或创建页返回列表时必须恢复原有 `queryText` 与 `statusFilter`

### Requirement: 合同与传输层必须一致

系统 SHALL 在联调与验收中核对 `.proto` 合同源、HTTP 过渡传输层与前端适配层之间的单值一致性。

#### Scenario: Proto 作为验收基准

- **WHEN** 执行接口联调验收
- **THEN** 必须以 `.proto` 作为当前阶段唯一合同源
- **AND** HTTP 请求与响应语义必须与 `.proto` 对齐
- **AND** 不得在联调过程中发现第二套独立 JSON 合同语义

### Requirement: 关键异常路径必须覆盖当前阶段边界

系统 SHALL 验证当前阶段最容易掩盖阻断的关键异常路径。

#### Scenario: 参数与资源异常

- **WHEN** 用户访问无效 `moduleId`、重复提交唯一约束、或传入非法状态值
- **THEN** 系统必须返回当前阶段已冻结的错误语义
- **AND** 不得出现 500 级未收口错误替代业务错误

#### Scenario: 候选读取与只读连接边界

- **WHEN** 用户打开产品候选、仓库候选或 `Decision` 入口
- **THEN** 只允许承接 `phase02` 已冻结的最小读取或跳转语义
- **AND** 不得在联调过程中扩写为新的写入主线

### Requirement: 验收结果必须形成可追踪证据

系统 SHALL 为 `phase02-12` 输出可复核的验收结果，而不是停留在口头“已通过”。

#### Scenario: 验收收口

- **WHEN** 完成联调与验证
- **THEN** 必须明确每条最小验收路径的结果
- **AND** 必须记录发现的问题与处理结论
- **AND** 未解决问题不得以“后续再看”形式遗留到 `phase02-13`

## MODIFIED Requirements

### Requirement: phase02 当前阶段完成条件

系统 SHALL 将 `phase02` 的完成条件从“前后端与合同已实现”推进为“前后端、数据与合同已经在同一环境中被验证可运行”。

#### Scenario: phase02-12 进入验收链

- **WHEN** `phase02-12` 开始执行
- **THEN** `phase02-10 / 11 / 11A` 的实现成果必须进入统一联调
- **AND** 当前阶段 DoD 必须以运行证据而不是实现声明为准

## REMOVED Requirements

### Requirement: 仅以单侧构建或局部接口自测视为阶段验收完成

**Reason**: 单侧构建通过或局部接口成功并不能证明 `Module Registry` 当前阶段最小主线已经形成完整可运行交付物。
**Migration**: 改为以前后端联调、关键动作验收、状态验证与 `.proto` 合同核对共同作为 `phase02-12` 验收标准。
