# Phase07-11 完成 phase01 ~ phase06 联调、回归与退场验收 Spec

## Why

`phase07-08`、`phase07-09`、`phase07-10` 已分别完成生成链、Go 后端传输主线与前端调用主线切换，但目前仍缺少一份把 `phase01 ~ phase06` 真实业务能力重新拉通的统一验收规格。没有这一层，`phase07` 只能证明“传输切换已经落地”，却还不能证明“旧主线已经退场、34 条 RPC 与 9 个业务模块在单一 `/api` + Connect 主线上可重复运行并可供后续 `mvp0.3` 直接承接”。

## What Changes

- 新增 `phase07-11` 联调、回归与退场验收规格，作为 `phase07` 最终运行验证与证据收口入口
- 冻结覆盖 `Module / Decision / Product / Repository / Dashboard / Onboarding / Export / Backup / Reuse Summary` 的跨 phase 联调矩阵
- 冻结本地启动链、`proto / backend / frontend` 构建链与 CI 等价生成链的统一验证口径
- 冻结 `34` 条 canonical RPC、`4` 条 legacy / compat endpoint inventory、`11` 项 frontend mutation owner inventory 的逐项核销要求
- 冻结正式验收记录与收口证据包结构，要求同一轮验收内显式记录环境、步骤、结果、问题与复测结论
- **BREAKING**：`phase07` 收口不再接受“单侧 build 通过”“局部页面看起来可用”或“前序 phase 曾各自验收通过”作为替代证据，必须通过本阶段统一联调与退场核销

## Impact

- Affected specs:
  - `phase07_06_design_business_api_migration_matrix_regression_acceptance`
  - `phase07_07_formal_transport_mainline_cutover_spec`
  - `phase07_08_land_buf_generation_connect_contract_mainline`
  - `phase07_09_cut_go_backend_transport_mainline`
  - `phase07_10_cut_frontend_client_slice_mainline`
  - `phase02_12_module_registry_integration_validation_acceptance`
  - `phase03_14_decision_center_integration_validation_acceptance`
  - `phase04_14_product_repository_binding_integration_validation_acceptance`
  - `phase05_14_dashboard_feedback_integration_validation_acceptance`
  - `phase06_16_integration_validation_acceptance`
- Affected code:
  - `frontend/src/**`
  - `backend/internal/**`
  - `proto/**`
  - `database/scripts/reset_*.sh`
  - `database/seeds/**`
  - `frontend/vite.config.ts`
  - `plan.md`
  - `docs/phase/phase07_transport_contract_mainline_migration_dev_plan.md`

## ADDED Requirements

### Requirement: phase07-11 必须以单一 `/api` + Connect 主线重建统一联调环境

系统 SHALL 在 `phase07-11` 中把 `phase01 ~ phase06` 的验收入口统一到同一套真实环境中执行，要求前端、后端、数据库、重置脚本与生成链都运行在单一 `/api` + Connect 主线上，不得回退到旧 hand-written JSON 业务主线、mock adapter 或临时兼容入口。

#### Scenario: 环境重建与前置条件核对

- **WHEN** 执行 `phase07-11` 联调与回归
- **THEN** 必须先核对 `proto/Makefile`、`backend`、`frontend`、数据库 reset 脚本与 `/api` 路由均处于可运行状态
- **AND** 必须明确开发环境、验收环境与部署等价环境都使用单一 `/api` 基址
- **AND** 不得为验收临时恢复旧 `api-adapter.ts`、旧 JSON handler 或第二套 API 基址
- **AND** 若当前仓库仍无正式 CI workflow，必须以本地等价链路作为唯一可复核替代证据，而不是跳过验证

### Requirement: phase07-11 必须完成 9 个业务模块的统一回归矩阵

系统 SHALL 把 `Module / Decision / Product / Repository / Dashboard / Onboarding / Export / Backup / Reuse Summary` 纳入同一轮回归矩阵，证明 `phase01 ~ phase06` 的正式业务能力在传输主线切换后仍保持等价运行。

#### Scenario: 模块内与跨模块最小回归

- **WHEN** 执行 `phase07-11` 回归
- **THEN** 必须至少覆盖 `phase07-06` 已冻结的模块内最小路径与跨模块联动路径
- **AND** `CreateModule / CreateRelease / CreateDecision / LinkDecisionToTarget / CreateProduct / CreateRepository / BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository` 必须逐项证明成功回流与 reread 仍然成立
- **AND** `Dashboard`、`Onboarding`、`Export`、`Backup`、`Reuse Summary` 必须证明在 Connect 主线下继续保持既有页面语义
- **AND** 回归通过的判定必须以真实页面行为、真实请求返回与 reread 结果为准，而不是仅凭代码审查

### Requirement: phase07-11 必须逐项核销 legacy / compat 业务入口 inventory

系统 SHALL 在本阶段对 `phase07-07` 冻结的 `L1 ~ L4` legacy / compat 入口逐项核销，确保旧业务入口已经退出正式运行面，且不存在未声明残留。

#### Scenario: legacy / compat 入口退场核销

- **WHEN** 执行 `phase07-11` 退场验收
- **THEN** 必须验证 `GET /api/candidates/products`、`GET /api/candidates/repositories`、`POST /api/modules/{moduleId}/bindings/products`、`POST /api/modules/{moduleId}/bindings/repositories` 已全部返回 `404`
- **AND** 必须同时提供路由删除、旧 handler/adapter 删除与替代 Connect path 正常工作的证据
- **AND** 不得接受“前端已经不调用，所以后端可继续保留”的残留状态
- **AND** 若发现任何旧入口仍可访问，`phase07-11` 不得通过

### Requirement: phase07-11 必须逐项核销 frontend mutation owner inventory

系统 SHALL 在本阶段对 `11` 项 frontend mutation owner inventory 执行逐项验收，证明 canonical 写动作 owner 已按计划收口，且不存在未声明的页面 / 组件级长期正式 mutation 主线。

#### Scenario: mutation owner 核销

- **WHEN** 执行 `phase07-11` owner 验收
- **THEN** `CreateModule / CreateDecision / CreateProduct / CreateRepository / CreateRelease / BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository / LinkDecisionToTarget` 必须逐项验证正式 owner 承接位与成功回流
- **AND** `ExportCoreAssets / CreateInstanceBackup` 只能以 `SovereigntyPanel` 中显式允许过渡位的身份存在
- **AND** 页面与组件中不得继续存在未声明的长期正式 `useMutation`
- **AND** 若存在新的页面级或组件级 canonical mutation 主线，`phase07-11` 不得通过

### Requirement: phase07-11 必须验证本地启动链、生成链与 CI 等价链路

系统 SHALL 把运行链与工具链验证纳入正式验收，而不是只验证浏览器页面结果。

#### Scenario: 本地与 CI 等价验证

- **WHEN** 执行 `phase07-11` 工具链验收
- **THEN** 必须验证 `(cd proto && make build && make gen && make lint)` 通过
- **AND** 必须验证 `(cd backend && go build ./...)` 通过
- **AND** 必须验证 `(cd frontend && npx tsc -b --noEmit)` 与 `(cd frontend && npm run build)` 通过
- **AND** 若仓库中不存在正式 CI workflow，必须把上述链路标记为“CI 等价验证”，并明确这只是当前阶段的替代证据，而不是 CI 已落地事实

### Requirement: phase07-11 必须形成单一正式验收记录与结论

系统 SHALL 将本阶段的环境、步骤、结果、问题、复测与收口判定沉淀为单一正式验收记录，供 `phase07-12` 根级同步与后续 `mvp0.3` phase 直接承接。

#### Scenario: 单一证据包与正式结论

- **WHEN** `phase07-11` 完成
- **THEN** 必须形成单一正式验收记录，至少包含环境前置条件、回归矩阵结果、legacy/compat 核销、mutation owner 核销、工具链验证、问题与复测结论、以及 DoD 达成情况
- **AND** 必须明确 `phase07` 是否具备进入 `phase07-12` 根级收口的条件
- **AND** 不得把同一轮验收结论拆散到多个并列“临时记录”中

## MODIFIED Requirements

### Requirement: phase07 当前阶段完成条件

`phase07` SHALL 将完成条件从“生成链、后端主线、前端主线已切换”推进为“`phase01 ~ phase06` 已在单一 `/api` + Connect 主线上完成统一联调、回归、退场核销与正式证据收口”。

#### Scenario: phase07-11 进入正式收口链

- **WHEN** 团队执行 `phase07-11`
- **THEN** `phase07-08 / 09 / 10` 的实现成果必须进入同一轮统一验收
- **AND** 前序 `phase02 / 03 / 04 / 05 / 06` 的局部验收结论只能作为上游证据来源，不再单独构成 `phase07` 通过依据
- **AND** 只有在统一回归、退场核销与工具链验证都通过后，`phase07` 才能进入 `phase07-12`

### Requirement: phase07-11 对 CI 的表述口径

`phase07-11` SHALL 明确区分“正式 CI 已落地”和“本地 CI 等价验证已完成”，不得把本地命令链通过写成 CI 已建设完成的既成事实。

#### Scenario: CI 缺口表述

- **WHEN** 当前仓库仍不存在 `.github/workflows/`
- **THEN** 正式验收记录必须把 `proto / backend / frontend` 命令链标记为 CI 等价验证
- **AND** 不得把 `phase07-11` 结果描述为“CI 已完成接入”
- **AND** 不得让 CI 缺口掩盖当前阶段真实通过的联调与退场结论

## REMOVED Requirements

### Requirement: 仅凭前序 phase 各自验收通过即可视为 phase07 回归完成

**Reason**: `phase07` 变更的是统一传输主线；前序 phase 的局部验收结论不足以证明 9 个业务模块在新主线下仍能协同运行。

**Migration**: 改为在 `phase07-11` 内统一执行跨模块回归、退场核销、owner 核销与工具链验证，并形成单一正式验收记录。

### Requirement: 仅凭单侧构建、局部 curl 或页面静态观察即可视为 phase07 收口完成

**Reason**: 这类证据不能同时证明真实页面行为、跨模块联动、旧入口退场与 owner 收口，无法支撑 `mvp0.3` 直接承接。

**Migration**: 改为采用“回归矩阵 + inventory 核销 + 工具链验证 + 单一正式结论”的组合验收方式。
