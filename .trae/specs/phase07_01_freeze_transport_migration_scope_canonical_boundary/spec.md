# Phase07-01 冻结本阶段迁移范围、Canonical 业务接口清单与非业务端点边界 Spec

> **执行产出**：`frozen_scope.md` — 包含完整的 34 条 proto-defined business RPC 迁移总表、legacy inventory、keep list、前端 owner 清单与收口判定标准。
> **执行日期**：2026-08-11
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

`phase07` 的目标不是“以后新增接口默认走 ConnectRPC”，而是把 `phase01 ~ phase06` 已交付的 canonical 业务接口一次性从 `chi + JSON HTTP` 过渡主线切到 `.proto + ConnectRPC` 正式主线。若在进入实现前不先冻结“哪些接口必须迁、哪些端点明确不迁、哪些 legacy 入口必须退场、何时才算真的完成切换”，后续 `/spec -> 实现 -> 验收` 仍会退化成边迁边猜、边做边补。

## What Changes

- 冻结 `phase01 ~ phase06` 必须纳入 `phase07` 的 canonical 业务模块与 RPC 清单
- 冻结到 `service / RPC / 当前外部访问路径 / 页面或动作 owner` 级别的迁移总表要求
- 冻结 `Module Registry` 下 module-centered compat RPC 与 Product / Repository canonical owner 的边界说明
- 冻结当前阶段允许继续保留在 `chi + net/http` 的非业务基础设施端点边界
- 冻结当前真实 `legacy / compat` 业务入口 inventory 及其退场要求
- 冻结 `phase07` 收口时“业务主线已切换”的判定标准与最小验收证据

## Impact

- Affected specs:
  - `phase07_transport_contract_mainline_migration`
  - `phase06_12_onboarding_sovereignty_reuse_formal_spec`
  - `phase06_13_land_minimal_proto_contract_mainline`
  - `phase06_16_integration_validation_acceptance`
- Affected code:
  - `proto/psco/*`
  - `proto/buf.gen.yaml`
  - `proto/Makefile`
  - `backend/internal/platform/router.go`
  - `backend/internal/*/handler/`
  - `frontend/src/features/*/data/api-adapter.ts`
  - `frontend/src/features/*/application/`
  - `frontend/src/features/*/components/`
  - `frontend/vite.config.ts`
  - `database/scripts/reset_phase06_acceptance.sh`

## ADDED Requirements

### Requirement: `phase07` 迁移范围必须冻结到 canonical 业务模块级别

系统 SHALL 在 `phase07-01` 中明确冻结本阶段必须完成正式传输主线切换的 canonical 业务模块，不允许以“先迁一部分、其余继续兼容”的方式通过本阶段。

#### Scenario: 冻结必须迁移的 canonical 业务模块

- **WHEN** 接手者定义 `phase07` 的迁移范围
- **THEN** 必须至少覆盖：
  - `psco.module_registry.v1`
  - `psco.decision_center.v1`
  - `psco.product_registry.v1`
  - `psco.repository_binding.v1`
  - `psco.dashboard.v1`
  - `psco.onboarding.v1`
  - `psco.export.v1`
  - `psco.backup.v1`
  - `psco.reuse_summary.v1`
- **AND** 不得把“只迁新增业务接口”或“只迁单个模块试点”解释为 `phase07` 完成
- **AND** 不得把仍依赖 hand-written JSON business handler 的模块留到 `mvp0.3` 业务 phase 再处理

### Requirement: 迁移总表必须下钻到 `service / RPC / 当前入口路径 / 页面或动作 owner`

系统 SHALL 将 `phase07` 的迁移清单冻结为可执行的总表，而不是停留在模块名或方向性描述。

#### Scenario: 冻结 canonical 迁移总表

- **WHEN** 接手者输出 `phase07-01` 迁移范围结果
- **THEN** 每个 canonical 业务模块下的每个 RPC 都必须至少明确：
  - 所属 service
  - RPC 名称
  - 当前外部访问路径
  - 当前页面或动作 owner
  - 目标 Connect procedure path
  - 当前 transport owner
  - 迁移后的正式 owner
- **AND** 不得只写“迁移 Module Registry / Decision Center / Dashboard”等模块名而不下钻
- **AND** 后续 `phase07-06`、`phase07-11` 的回归与验收必须直接复用这份总表

### Requirement: 非业务基础设施端点边界必须冻结为 keep list

系统 SHALL 在 `phase07-01` 中明确列出继续保留在 `chi + net/http` 的非业务基础设施端点，避免实现阶段误把所有 HTTP 端点一并迁入 `.proto`。

#### Scenario: 冻结非业务 keep list

- **WHEN** 接手者定义不纳入 `phase07` 迁移范围的端点
- **THEN** 允许继续保留在 `chi + net/http` 的端点只包括：
  - `healthz`
  - `readyz`
  - `metrics`
  - `debug / pprof`
- **AND** 这些端点不得被解释为 canonical 业务接口
- **AND** 除上述端点外，不得再把其他业务主线接口列入长期 keep list

### Requirement: 当前真实 `legacy / compat` 业务入口必须显式建账

系统 SHALL 在 `phase07-01` 中显式冻结当前仓库已存在的 `legacy / compat` 业务入口 inventory，为后续退场与验收提供 endpoint 级核销依据。

#### Scenario: 冻结 legacy / compat endpoint inventory

- **WHEN** 接手者建立当前真实 transport inventory
- **THEN** 至少必须点名以下兼容入口：
  - `/api/candidates/products`
  - `/api/candidates/repositories`
  - `/api/modules/{moduleId}/bindings/products`
  - `/api/modules/{moduleId}/bindings/repositories`
- **AND** 每个入口都必须明确：
  - 当前调用方
  - 当前存在原因
  - 替代 RPC / Connect path
  - 允许并存的最晚时点
  - 最终删除证据
- **AND** 不得把这些入口以“兼容层”名义默认保留到 `phase07` 收口后

### Requirement: `phase07` 收口判定必须冻结为“正式主线已切换”

系统 SHALL 在 `phase07-01` 中把“业务主线已切换”的完成判定写成验收门禁，而不是实现结束后的主观判断。

#### Scenario: 判定 `phase07` 是否可收口

- **WHEN** 接手者判断 `phase07` 是否达到完成条件
- **THEN** 至少必须同时满足：
  - `phase01 ~ phase06` canonical 业务接口均已切到 ConnectRPC
  - canonical 迁移总表已覆盖到 `service / RPC / 当前入口路径 / 页面动作 owner`
  - 浏览器侧、Vite dev proxy、验收脚本与部署链路仍通过单一 `/api` 基址访问业务接口
  - `legacy / compat` 业务入口 inventory 已逐项核销
  - phase 收口后不存在“新 Connect 主线 + 旧 JSON 主线”并列正式状态
- **AND** 不得把“新主线已增加，但旧 JSON 主线仍可继续工作”解释为完成

### Requirement: 前端动作 owner 必须进入迁移范围表

系统 SHALL 把前端页面或动作 owner 作为 `phase07-01` 迁移范围的一部分冻结下来，避免后续只迁 transport 而忽略正式写动作承接位。

#### Scenario: 冻结页面或动作 owner

- **WHEN** 接手者为每个 RPC 建立迁移总表
- **THEN** 必须明确对应的页面、面板、query owner、application owner 或其他正式动作承接位
- **AND** 对仍存在于页面 / 组件中的正式 mutation，必须标记为：
  - 将回收到 `application` owner
  - 或允许存在的短时过渡位
- **AND** 不得把“保持 query / application 边界”只写成抽象原则而没有对应 owner 清单

## MODIFIED Requirements

### Requirement: `phase07` 的“迁移范围”定义

`phase07` 的迁移范围 SHALL 被解释为“已交付 canonical 业务接口的一次性正式切换”，而不是“为未来实现建立一般性 ConnectRPC 原则”。

#### Scenario: 解释 `phase07-01` 范围

- **WHEN** 接手者阅读 `phase07-01`
- **THEN** 应理解其输出是正式迁移范围、canonical 业务接口清单、非业务 keep list、legacy inventory 与收口判定标准
- **AND** 不得把 `phase07-01` 理解成仅做概念说明或接口方向讨论

## REMOVED Requirements

### Requirement: 可在 `phase07` 中保留部分 canonical 模块继续走 hand-written JSON 业务主线

**Reason**: 这会直接破坏 `audit_001`、`phase07` shared baseline 和 architecture plan 已冻结的“一次性正式切换”目标，让 `phase07` 重新退化成兼容并存阶段。

**Migration**: 统一改为在 `phase07-01` 中显式列出所有必须迁移的 canonical 模块与 RPC；允许短时存在的仅能是实施过程中的临时 adapter 或排障开关，且不得写入阶段完成态。
