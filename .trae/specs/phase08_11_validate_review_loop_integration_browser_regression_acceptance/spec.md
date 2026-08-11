# phase08-11 完成 review loop 联调、浏览器验收与反回归验证 Spec

## Why

`phase08-08` 已完成 review 合同、后端承接与前端 owner 收敛，`phase08-09` 已完成 Dashboard review 入口、双路径会话与统一动作承接，`phase08-10` 已把 `Feedback -> Decision -> Update` 的最小闭环推进到真实可交互状态。当前仍缺少一份把这些能力放回统一运行环境中重新拉通的正式验收规格；否则我们只能证明“局部能力能工作”，还不能证明 `Operating Review Loop` 已达到可收口的阶段验收标准。

## What Changes

- 新增 `phase08-11` 联调、浏览器验收与反回归验证规格，作为 `phase08` 的统一运行验证入口
- 冻结 `Dashboard -> Daily Review -> Decision -> Update` 与 `Dashboard -> Weekly Review -> Decision -> Update` 两条关键经营路径的独立验收矩阵
- 冻结 `phase05` Dashboard/Feedback 与 `phase06` Reuse Awareness 在 review loop 中的正式消费与反回归口径
- 冻结 `buf build / lint / generate`、`go build ./...`、`npx tsc -b --noEmit`、`frontend build`、关键 Connect procedure 与 `/api` smoke 的统一验证顺序
- 冻结浏览器端关键交互与 `phase03 ~ phase06` 相关页面的最小反回归范围
- 冻结本阶段边界证据：`Template Reuse / Derived Intelligence / dry-run` 继续明确不做
- **BREAKING**：`phase08` 的通过结论不再接受“单条路径源码正确”“局部构建通过”或“单一页面浏览器点击通过”作为替代证据，必须经由本阶段统一联调、浏览器验收与反回归验证

## Impact

- Affected specs:
  - `phase08_08_land_review_contract_backend_frontend_owner_enablement`
  - `phase08_09_land_dashboard_review_entry_dual_session_unified_action_handoff`
  - `phase08_10_land_feedback_decision_update_closed_loop_result_writeback_cleanup`
  - `phase05_14_dashboard_feedback_integration_validation_acceptance`
  - `phase06_16_integration_validation_acceptance`
  - `phase03_14_decision_center_integration_validation_acceptance`
  - `phase04_14_product_repository_binding_integration_validation_acceptance`
- Affected code:
  - `frontend/src/features/review/**`
  - `frontend/src/features/dashboard/**`
  - `frontend/src/features/decision-center/**`
  - `frontend/src/features/module-registry/**`
  - `frontend/src/features/product-registry/**`
  - `frontend/src/features/repository-binding/**`
  - `frontend/src/features/onboarding/**`
  - `frontend/src/features/reuse-summary/**`
  - `backend/internal/review/**`
  - `backend/internal/dashboard/**`
  - `backend/internal/decisioncenter/**`
  - `backend/internal/moduleregistry/**`
  - `backend/internal/productregistry/**`
  - `backend/internal/repositorybinding/**`
  - `backend/internal/onboarding/**`
  - `backend/internal/reusesummary/**`
  - `proto/**`

## ADDED Requirements

### Requirement: `phase08-11` 必须在单一正式运行环境中重建 review loop 验收基线

系统 SHALL 在 `phase08-11` 中把 `phase08-08 / 09 / 10` 已交付能力统一放回同一套真实运行环境中执行，要求前端、后端、数据库、Connect procedure 与 `/api` 访问都运行在当前正式主线上，不得回退到 mock、旧适配层或临时旁路。

#### Scenario: 统一环境前置检查

- **WHEN** 执行 `phase08-11` 联调与浏览器验收
- **THEN** 必须先确认前端、后端、数据库与 `/api` 主线处于可运行状态
- **AND** 必须以当前仓库正式 `ConnectRPC + React Web` 主线作为唯一运行口径
- **AND** 不得为本阶段验收临时恢复第二套页面入口、第二套 transport 主线或 review-local 影子写路径

### Requirement: `phase08-11` 必须分别验证 daily / weekly review 两条关键经营路径

系统 SHALL 将 `Dashboard -> Daily Review -> Decision -> Update` 与 `Dashboard -> Weekly Review -> Decision -> Update` 冻结为两条独立关键路径，分别验证其页面入口、正式承接、结果回流与返回链，而不是让一条路径的通过结论替代另一条。

#### Scenario: Daily Review 路径验收

- **WHEN** 团队执行 Daily Review 浏览器与联调验收
- **THEN** 必须验证 `Dashboard -> Daily Review -> Decision` 正式进入成立
- **AND** 必须验证至少一种实体 canonical update 或正式 action handoff 已真实成立
- **AND** 必须验证成功后的 reread、返回 Dashboard 与来源参数语义保持一致

#### Scenario: Weekly Review 路径验收

- **WHEN** 团队执行 Weekly Review 浏览器与联调验收
- **THEN** 必须验证 `Dashboard -> Weekly Review -> Decision` 正式进入成立
- **AND** 必须验证 `Weekly Review` 不复用 `Daily Review` 的数据装配或完成定义冒充通过
- **AND** 必须验证 weekly 成功路径、返回链与 Dashboard 入口语义保持一致

### Requirement: `phase08-11` 必须验证 weekly review 对 `phase05 / phase06` 读模型的正式消费

系统 SHALL 在本阶段把 `overview / recent activity / representative signals / reuse snapshot / module_reuse_summary / capability_summary` 的消费验证纳入 Weekly Review 正式验收，而不是只证明页面能打开。

#### Scenario: Weekly Review 读模型消费验收

- **WHEN** 执行 Weekly Review 联调与浏览器验收
- **THEN** 必须验证 `phase05` 提供的 `overview / recent activity / representative signals` 仍按正式语义显示
- **AND** 必须验证 `phase06` 提供的 `reuse snapshot / module_reuse_summary / capability_summary` 已被正式消费
- **AND** 必须验证局部失败边界未回退成整页崩溃或错误语义漂移

### Requirement: `phase08-11` 必须完成工具链、API smoke 与关键 Connect procedure 验证

系统 SHALL 将 `buf`、`backend`、`frontend` 构建链，以及 review loop 关键 procedure 与 `/api` 访问 smoke 纳入统一验收，而不是只凭浏览器点击通过给出阶段结论。

#### Scenario: 工具链与接口主线验证

- **WHEN** 执行 `phase08-11` 验收
- **THEN** 必须验证 `(cd proto && make build && make gen && make lint)` 通过
- **AND** 必须验证 `(cd backend && go build ./...)` 通过
- **AND** 必须验证 `(cd frontend && npx tsc -b --noEmit)` 与 `(cd frontend && npm run build)` 通过
- **AND** 必须验证 review loop 关键 Connect procedure 与 `/api` 路径 smoke 成立
- **AND** 不得把“局部命令跑过一次”写成“整体验收已经完成”

### Requirement: `phase08-11` 必须覆盖浏览器端关键交互与 `phase03 ~ phase06` 最小反回归范围

系统 SHALL 在 `phase08-11` 中验证浏览器端不存在“API 成功但 UI 崩溃”的收口缺口，并对 `phase03 ~ phase06` 相关页面执行最小反回归检查，证明 review loop 的引入没有破坏既有 canonical 页面。

#### Scenario: 浏览器关键交互验收

- **WHEN** 执行浏览器验收
- **THEN** 必须验证 Dashboard、Daily Review、Weekly Review、Decision、Module、Product、Repository 等关键页面可正常进入与返回
- **AND** 必须验证关键点击、跳转、成功回流与 reread 结果可在真实浏览器中观察到
- **AND** 必须确认 console 中不存在阻断级 runtime error

#### Scenario: phase03 ~ phase06 最小反回归

- **WHEN** 执行最小反回归验证
- **THEN** 必须至少覆盖 Decision、Product / Repository Binding、Dashboard / Feedback、Onboarding、Reuse Summary 等与 review loop 直接关联的既有正式页面
- **AND** 必须证明这些页面在当前主线下仍保持既有正式语义
- **AND** 若发现 review loop 引入导致的页面崩溃、来源链丢失或 owner 越界，则 `phase08-11` 不得通过

### Requirement: `phase08-11` 必须记录本阶段边界未漂移的证据

系统 SHALL 在正式验收中显式记录 `Template Reuse / Derived Intelligence / dry-run` 仍不属于本阶段交付范围，避免在验收结论中把未来能力误写成当前既成事实。

#### Scenario: 非目标边界核对

- **WHEN** 形成 `phase08-11` 正式结论
- **THEN** 必须显式记录本阶段只验证 operating review loop 当前已冻结能力
- **AND** 不得把 `Template Reuse / Derived Intelligence / dry-run` 写成已进入本阶段交付
- **AND** 不得因为浏览器验收顺利就扩写新的业务范围

### Requirement: `phase08-11` 必须形成单一正式验收结论并作为 `phase08-12` 的直接上游

系统 SHALL 将本阶段的环境、步骤、结果、问题、复测与 DoD 判定收敛为单一正式验收结论，供 `phase08-12` 根级同步直接承接。

#### Scenario: 正式结论收口

- **WHEN** `phase08-11` 验收完成
- **THEN** 必须形成单一正式结论，至少覆盖工具链、API smoke、daily / weekly 浏览器路径、反回归、边界证据与问题收口
- **AND** 必须明确 `phase08` 是否具备进入 `phase08-12` 根级同步的条件
- **AND** 不得把同一轮结论拆散为多个并列临时说明

## MODIFIED Requirements

### Requirement: `phase08` 当前阶段的通过口径

`phase08` SHALL 将“通过验收”的判定从“局部子任务已分别落地”推进为“review loop 已完成统一联调、双路径浏览器验收、关键 API smoke 与最小反回归验证”。

#### Scenario: `phase08-11` 进入正式收口链

- **WHEN** 团队执行 `phase08-11`
- **THEN** `phase08-08 / 09 / 10` 的实现成果必须进入同一轮统一验收
- **AND** 前序子任务的局部通过结论只能作为上游证据来源，不再单独构成 `phase08` 最终通过依据
- **AND** 只有在本阶段统一验收通过后，`phase08` 才能进入 `phase08-12`

### Requirement: `phase08-10` 的闭环验证在 `phase08-11` 中的定位

`phase08-10` SHALL 在 `phase08-11` 中被继承为“已具备最小闭环实现”的上游输入，而不是代替本阶段对 daily / weekly 双路径、反回归与边界证据的统一验证。

#### Scenario: 上游输入与本阶段职责分离

- **WHEN** 团队引用 `phase08-10` 已通过的闭环结论
- **THEN** 可以将其作为 Daily 路径的直接上游证据
- **AND** 但仍必须在 `phase08-11` 中补齐 Weekly 路径、最小反回归与阶段边界证据
- **AND** 不得把 `phase08-10` 的单路径闭环通过，直接写成整个 review loop 已完成最终验收

## REMOVED Requirements

### Requirement: 仅凭单条 review 路径或单页浏览器点击通过即可视为 `phase08` 联调验收完成

**Reason**: 这类证据无法同时证明 daily / weekly 双路径独立成立、Weekly Review 正式消费 `phase05 / phase06` 读模型、以及 review loop 没有破坏既有 canonical 页面。

**Migration**: 改为在 `phase08-11` 中统一执行“工具链 + API smoke + 双路径浏览器验收 + 最小反回归 + 边界证据”组合验证。

### Requirement: 仅凭构建通过或局部 API smoke 即可视为 review loop 收口完成

**Reason**: 这类证据不能覆盖真实浏览器交互、返回链、reread 结果与 UI 运行时稳定性，无法支撑 `phase08` 正式收口。

**Migration**: 保留构建与 API smoke 作为必要组成部分，但必须追加真实浏览器与反回归证据后才可给出通过结论。
