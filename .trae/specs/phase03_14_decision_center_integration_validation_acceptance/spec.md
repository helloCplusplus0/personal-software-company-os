# Phase03-14 Decision Center 联调、验证与验收 Spec

## Why

`phase03-11`、`phase03-12` 与 `phase03-13` 已分别完成 `.proto` 合同主线、后端与数据主线、前端主线，但当前仍需要一份独立的联调验收规格，把“各侧已实现”收口为“最小主线已在同一环境中被验证可运行”。如果缺少这一层，`Decision Center` 会停留在局部自证完成，无法形成可重复复核、可明确证明当前阶段已形成交付物的运行证据。

同时，`phase03-14` 还必须承担阶段收口职责：空状态、错误态、返回路径、目标不存在、重复关联与合同一致性都要进入同一轮验收；联调中发现的实现偏差或正式规格漂移，必须在当前阶段明确收口，不能把隐性阻断带到 `phase03-15`。

## What Changes

- 为 `Decision Center` 新增一份联调、验证与验收 spec
- 冻结联调环境重建顺序、真实前后端联调入口与验收前置条件
- 冻结 `Decision List -> Decision Create -> Decision Detail -> LinkDecisionToTarget` 的最小主线验收路径
- 冻结从 `Module Detail` 发起的来源上下文入口、创建回流、详情承接与返回路径验收
- 冻结空状态、错误态、目标不存在、重复关联、候选为空与返回路径验证要求
- 冻结 `.proto`、HTTP 过渡层、前端适配层与正式规格正文的一致性核对要求
- 明确联调中发现的问题必须在当前阶段收口，不得遗留隐性阻断或上游规格漂移

## Impact

- Affected specs:
  - `phase03_09_decision_center_integration_baseline`
  - `phase03_10_decision_center_formal_spec`
  - `phase03_11_decision_center_proto_mainline`
  - `phase03_12_decision_center_backend_data_mainline`
  - `phase03_13_decision_center_frontend_mainline`
- Affected code:
  - `frontend/src/features/decision-center/**`
  - `frontend/src/routes/decisions/**`
  - `backend/internal/decisioncenter/**`
  - `backend/internal/platform/**`
  - `database/migrations/**`
  - `database/scripts/reset_decision_mainline.sh`
  - `database/seeds/seed_decision_mainline_baseline.sql`
  - `proto/psco/decision_center/v1/decision_center.proto`

## ADDED Requirements

### Requirement: 联调环境必须可重复建立并以真实主线运行

系统 SHALL 为 `phase03-14` 提供可重复建立的联调环境前提，使验收在真实前后端与真实数据主线上执行，而不是依赖 mock、手工补数据或一次性会话状态。

#### Scenario: 环境初始化与启动顺序

- **WHEN** 执行 `phase03-14` 联调与验收
- **THEN** 必须复用 `phase03-09` 已冻结的环境顺序：`init_db.sh` -> 后端启动（自动 migration）-> `run_seeds.sh` -> `reset_module_mainline.sh` -> `reset_decision_mainline.sh` -> 前端启动
- **AND** 前端必须直接连接 `phase03-12` 已落地的真实后端 API，不得切回 mock 数据主线
- **AND** 不得在验收过程中新增临时 SQL、临时 seed 或临时脚本弥补环境缺口

#### Scenario: 联调前置条件核对

- **WHEN** 开始执行验收
- **THEN** 必须先明确当前 `.proto` 合同、后端路由、数据库 migration、重置脚本与前端页面主线都已处于可运行状态
- **AND** 必须先确认 `Decision Center` 前端导航可达、后端 `/api` 已挂载、`reset_decision_mainline.sh` 可执行
- **AND** 若任何前置条件不满足，不得跳过并继续手工联调

### Requirement: 最小主线必须端到端完整走通

系统 SHALL 证明 `Decision Center` 当前阶段最小主线已经在同一环境中形成真实可运行交付物，而不是停留在单侧 build、单个接口自测或局部页面演示。

#### Scenario: 冷启动最小闭环

- **WHEN** 执行冷启动验收路径
- **THEN** 必须按以下顺序走通：
- **AND** `reset_decision_mainline.sh --clean-only`
- **AND** 进入 `Decision Center / List` 验证空状态入口
- **AND** 从空状态进入 `Decision Create`
- **AND** 提交 `RecordDecision`
- **AND** 默认回流到 `DecisionDetailPage`
- **AND** 在详情页读取候选列表并执行 `LinkDecisionToTarget`
- **AND** 停留当前详情页并重新读取出最新关联结果
- **AND** 全路径不得依赖手工 SQL

#### Scenario: Module Detail 来源上下文闭环

- **WHEN** 用户从 `Module Detail` 发起“为当前 Module 记录决策”
- **THEN** 创建页必须承接 `sourceModuleId / sourceModuleName`
- **AND** 创建成功后详情页必须通过后端 `source_context` 承接来源 `Module`
- **AND** 待关联目标必须在详情页显式展示，直到完成正式 `LinkDecisionToTarget`
- **AND** 完成正式关联后，待关联目标卡片必须由 reread 驱动消失

### Requirement: 空状态、错误态与返回路径必须被统一验证

系统 SHALL 将关键 UI 状态纳入本轮验收，而不是只验证 happy path。

#### Scenario: 空状态与空候选状态

- **WHEN** `decisions` 表为空或候选读取为空
- **THEN** `Decision Center / List` 必须展示围绕“先记录首条决策”的空状态入口
- **AND** 空状态主动作必须直接进入 `Decision Create`
- **AND** `DecisionModuleCandidateRead` 为空时必须表现为可解释的空列表状态，不得误报资源不存在或接口错误

#### Scenario: 错误态与表单保留

- **WHEN** 创建、详情读取或目标关联失败
- **THEN** 错误必须停留在当前页面或当前面板上下文
- **AND** 不得跳转独立错误页
- **AND** 创建失败时草稿与来源上下文必须继续保留
- **AND** 关联失败时当前详情、待关联目标与候选上下文必须继续保留

#### Scenario: 返回路径验证

- **WHEN** 用户从 `DecisionCreatePage` 或 `DecisionDetailPage` 主动返回列表
- **THEN** 必须符合 `phase03-06 / 10 / 13` 已冻结的返回路径规则
- **AND** 从 `DecisionListPage` 进入时必须恢复原有 `queryText / statusFilter`
- **AND** 从 `Module Detail` 入口或外部直达进入时必须落到默认列表参数（`statusFilter: 'all'`），不得恢复历史筛选

### Requirement: 关键异常路径必须覆盖当前阶段边界

系统 SHALL 验证 `Decision Center` 当前阶段最容易掩盖阻断的异常路径，并以当前冻结的业务错误语义作为唯一验收标准。

#### Scenario: 最小异常路径覆盖

- **WHEN** 执行 `phase03-14` 异常路径验收
- **THEN** 必须至少覆盖以下 `8` 类异常/边界路径：
- **AND** `RecordDecision` 必填字段缺失 -> 400 校验失败
- **AND** `RecordDecision` 字段值非法（含空白 `alternatives` 条目、非法 `status`）-> 400 校验失败
- **AND** `RecordDecision` 传入无效 `source_module_id` -> 404 资源不存在
- **AND** `LinkDecisionToTarget` 目标类型越界（非 `MODULE`）-> 400 校验失败
- **AND** `LinkDecisionToTarget` 的 `decision_id` 不存在 -> 404 资源不存在
- **AND** `LinkDecisionToTarget` 的 `module_id` 不存在 -> 404 资源不存在
- **AND** `LinkDecisionToTarget` 重复关联 -> 409 重复冲突
- **AND** `DecisionModuleCandidateRead` 无可关联候选 -> 空列表语义（非错误）
- **AND** `DecisionDetailRead` 不存在的 `decision_id` -> 404 资源不存在
- **AND** 不得出现 500 级未收口错误替代业务错误

#### Scenario: 异常前提建立方式

- **WHEN** 异常路径需要前提数据
- **THEN** 必须优先通过基线 seed 与受控 API 操作建立
- **AND** 不得为某个异常路径新建独立 fixture SQL 文件
- **AND** 不得通过手工改库制造异常前提

### Requirement: 合同、传输层、前端适配层与正式规格必须一致

系统 SHALL 在联调与验收中同时核对 `.proto` 合同源、HTTP 过渡传输层、前端适配层和正式规格正文，确保当前阶段不存在第二套合同语义或未收口的规格漂移。

#### Scenario: 合同与 HTTP 语义核对

- **WHEN** 执行接口联调验收
- **THEN** 必须以 `.proto` 作为当前阶段唯一合同源
- **AND** HTTP 请求与响应语义必须与 `.proto` 对齐
- **AND** 前端 `types / api-adapter / decision-center-adapter` 的对象语义必须与 HTTP / `.proto` 对齐
- **AND** 不得在联调过程中发现第二套并列 JSON 合同语义

#### Scenario: 正式规格正文一致性核对

- **WHEN** 联调中发现 `phase03-10` 正式规格、`phase03-12 / 13` 实现 spec 与仓库现实不一致
- **THEN** 必须在当前阶段给出单值结论并完成收口
- **AND** 不得把“formal spec 未同步”“实现已改但上游未回写”视为可忽略瑕疵
- **AND** 当前阶段通过的前提是：`phase03-10` 正式规格正文与已验收实现边界一致，不再保留误导后续阶段的旧承诺

### Requirement: 验收结果必须形成可重复复核证据

系统 SHALL 为 `phase03-14` 输出可追踪、可重复复核的验收结果，而不是停留在口头“已通过”。

#### Scenario: 证据收口

- **WHEN** 完成联调与验证
- **THEN** 必须明确每条最小验收路径的结果
- **AND** 必须明确环境重建入口、执行顺序、关键请求/响应结果与页面行为结果
- **AND** 必须记录发现的问题、修复结论与剩余风险
- **AND** 未解决问题不得以“后续再看”形式遗留到 `phase03-15`

## MODIFIED Requirements

### Requirement: phase03 当前阶段完成条件

系统 SHALL 将 `phase03` 当前阶段的完成条件从“合同、后端、前端都已实现”推进为“合同、后端、前端已经在同一环境中被验证可运行，且本阶段发现的问题已全部收口”。

#### Scenario: phase03-14 进入统一验收链

- **WHEN** `phase03-14` 开始执行
- **THEN** `phase03-11 / 12 / 13` 的成果必须进入统一联调
- **AND** 当前阶段 DoD 必须以可重复运行证据而不是实现声明为准
- **AND** `phase03-14` 不仅验证功能可跑通，还必须验证当前阶段不存在隐性阻断或正式规格漂移

## REMOVED Requirements

### Requirement: 仅以单侧 build 通过或局部接口成功视为 phase03 验收完成

**Reason**: 单侧 build、静态页面可见或局部接口成功，并不能证明 `Decision Center` 当前阶段最小主线已经形成完整可运行交付物。

**Migration**: 改为以前后端联调、冷启动路径、关键异常路径、返回路径、合同一致性与问题收口结论共同作为 `phase03-14` 的验收标准。
