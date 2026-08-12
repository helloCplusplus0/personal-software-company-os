# phase09-11 完成模板复用与派生提示联调、浏览器验收与反回归验证 Spec

## Why

`phase09-08 / 09 / 10` 已分别完成模板复用合同、模板候选与预填回流、派生提示展示与 handoff 的正式实现，但当前仍缺少一份把这些能力放回同一运行环境重新拉通的统一验收规格。否则我们只能证明局部实现成立，不能证明 `Template Reuse + Derived Intelligence` 已达到可收口的阶段验收标准。

## What Changes

- 新增 `phase09-11` 联调、浏览器验收与反回归验证规格，作为 `phase09` 当前实现链的统一运行验证入口
- 冻结 `buf / frontend type-check / build / backend build` 的统一工具链验证顺序与通过口径
- 冻结模板候选读取、模板预填详情读取、派生提示读取三类最小 API smoke 矩阵
- 冻结 `Weekly Review -> Product Create 预填 -> Product Detail` 的最小浏览器验收闭环、字段级预填断言与成功/空态/失败态判定标准
- 冻结 `Dashboard / Review / Product Detail / ReuseSummary` 的最小反回归页面清单与 reread 观察断言
- 冻结本阶段边界证据：`dry-run / AI Context Enhancement / Venture` 继续明确不做
- **BREAKING**：`phase09` 的通过结论不再接受“单个子任务已通过”“单条 API 能返回”或“局部浏览器点击过一次”作为替代证据，必须经由本阶段统一联调、浏览器验收与反回归验证

## Impact

- Affected specs:
  - `phase09_08_land_template_reuse_contract_backend_frontend_read_enablement`
  - `phase09_09_land_template_candidate_selection_product_create_prefill_result_handoff`
  - `phase09_10_land_derived_hint_display_action_handoff_explanatory_reread`
  - `phase08_11_validate_review_loop_integration_browser_regression_acceptance`
  - `phase05_14_dashboard_feedback_integration_validation_acceptance`
  - `phase06_16_integration_validation_acceptance`
- Affected code:
  - `frontend/src/features/review/**`
  - `frontend/src/features/product-registry/**`
  - `frontend/src/features/template-reuse/**`
  - `frontend/src/features/dashboard/**`
  - `frontend/src/features/reuse-summary/**`
  - `backend/internal/templatereuse/**`
  - `backend/internal/review/**`
  - `backend/internal/productregistry/**`
  - `backend/internal/dashboard/**`
  - `backend/internal/reusesummary/**`
  - `proto/**`

## ADDED Requirements

### Requirement: `phase09-11` 必须在单一正式运行环境中重建模板复用与派生提示验收基线

系统 SHALL 在 `phase09-11` 中把 `phase09-08 / 09 / 10` 已交付能力统一放回同一套真实运行环境中执行，要求前端、后端、数据库、Connect procedure 与 `/api` 访问都运行在当前正式主线上，不得回退到 mock、旧适配层或临时旁路。

#### Scenario: 统一环境前置检查

- **WHEN** 执行 `phase09-11` 联调与浏览器验收
- **THEN** 必须先确认前端、后端、数据库与 `/api` 主线处于可运行状态
- **AND** 必须以当前仓库正式 `ConnectRPC + React Web` 主线作为唯一运行口径
- **AND** 不得为本阶段验收临时恢复第二套模板主线、第二套提示主线或第二套 create 写路径

### Requirement: `phase09-11` 必须完成统一工具链验证与最小 API smoke 矩阵

系统 SHALL 将 `buf`、`backend`、`frontend` 构建链，以及模板候选、模板预填详情、派生提示三类正式接口验证纳入统一验收，而不是只凭浏览器点击给出阶段结论。

#### Scenario: 工具链与 API smoke 验证

- **WHEN** 执行 `phase09-11` 验收
- **THEN** 必须验证 `buf`、`frontend type-check`、`frontend build`、`backend build` 按正式顺序通过
- **AND** 必须验证模板候选读取、模板预填详情读取、派生提示读取三类 API smoke 成立
- **AND** 必须确认 API smoke 使用当前正式合同与 transport 主线，而不是临时脚本私有语义
- **AND** 不得把“局部命令跑过一次”写成“整体验收已经完成”

### Requirement: `phase09-11` 必须冻结单一浏览器闭环与字段级可观察断言

系统 SHALL 将 `Weekly Review -> Product Create 预填 -> Product Detail` 冻结为本阶段唯一正式浏览器闭环，并明确模板候选、active candidate、派生提示、预填字段与结果回流的机械判定标准。

#### Scenario: 浏览器闭环验收

- **WHEN** 执行本阶段浏览器验收
- **THEN** 必须先判断 `Weekly Review` 是否出现模板候选或成功空态
- **AND** 必须判断是否存在单值 active candidate 或成功空态
- **AND** 必须判断是否出现派生提示或成功空态
- **AND** 必须验证是否通过 `templateCandidateId` 进入了可编辑预填创建页
- **AND** 必须验证 `Product Create` 预填字段真实缩短创建路径，而不是仍然要求手工重填模板信息

#### Scenario: Product Detail 回流观察

- **WHEN** 浏览器闭环完成创建并进入 `Product Detail`
- **THEN** 必须看到模板来源摘要
- **AND** 必须看到候选 `Module` 组合摘要
- **AND** 必须看到 canonical binding CTA
- **AND** 若任一要素缺失，则本阶段浏览器验收不得通过

### Requirement: `phase09-11` 必须完成 `Dashboard / Review / Product Detail / ReuseSummary` 的最小反回归验证

系统 SHALL 在本阶段验证模板复用与派生提示的引入没有破坏既有页面主链，并将 `Dashboard / Review / Product Detail / ReuseSummary` 冻结为最小反回归页面清单。

#### Scenario: 反回归页面验证

- **WHEN** 执行最小反回归验证
- **THEN** 必须至少覆盖 `Dashboard / Review / Product Detail / ReuseSummary`
- **AND** 必须确认这些页面在当前主线下仍保持既有正式语义
- **AND** 必须验证成功 reread 成立，且未把“无统计变化”误判为失败
- **AND** 若发现页面崩溃、来源链丢失、reread 语义漂移或 owner 越界，则 `phase09-11` 不得通过

### Requirement: `phase09-11` 必须记录本阶段边界未漂移的证据

系统 SHALL 在正式验收中显式记录 `dry-run / AI Context Enhancement / Venture` 仍不属于本阶段交付范围，避免把未来能力误写成当前既成事实。

#### Scenario: 非目标边界核对

- **WHEN** 形成 `phase09-11` 正式结论
- **THEN** 必须显式记录本阶段只验证模板复用与派生提示当前已冻结能力
- **AND** 不得把 `dry-run / AI Context Enhancement / Venture` 写成已进入本阶段交付
- **AND** 不得因为浏览器验收顺利就扩写新的业务范围

### Requirement: `phase09-11` 必须形成单一正式验收结论

系统 SHALL 将本阶段的环境、步骤、结果、问题、复测与 DoD 判定收敛为单一正式验收结论，供 `phase09` 后续收口直接承接。

#### Scenario: 正式结论收口

- **WHEN** `phase09-11` 验收完成
- **THEN** 必须形成单一正式结论，至少覆盖工具链、API smoke、浏览器闭环、最小反回归、reread 断言与边界证据
- **AND** 必须明确当前实现是否具备进入后续阶段收口的条件
- **AND** 不得把同一轮结论拆散为多个并列临时说明

## MODIFIED Requirements

### Requirement: `phase09` 当前阶段的通过口径

`phase09` SHALL 将“通过验收”的判定从“`phase09-08 / 09 / 10` 各自局部通过”推进为“模板复用与派生提示已完成统一联调、浏览器闭环验收、API smoke 与关键反回归验证”。

#### Scenario: `phase09-11` 进入统一验收链

- **WHEN** 团队执行 `phase09-11`
- **THEN** `phase09-08 / 09 / 10` 的实现成果必须进入同一轮统一验收
- **AND** 前序子任务的局部通过结论只能作为上游证据来源，不再单独构成 `phase09` 当前最终通过依据
- **AND** 只有在本阶段统一验收通过后，当前实现链才具备进入后续收口的条件

### Requirement: `phase09-10` 的通过结论在 `phase09-11` 中的定位

`phase09-10` SHALL 在 `phase09-11` 中被继承为“已具备提示展示与 handoff 实现”的上游输入，而不是代替本阶段对工具链、浏览器闭环、反回归与边界证据的统一验证。

#### Scenario: 上游输入与本阶段职责分离

- **WHEN** 团队引用 `phase09-10` 已通过的实现结论
- **THEN** 可以将其作为浏览器闭环与 reread 断言的直接上游证据
- **AND** 但仍必须在 `phase09-11` 中补齐统一工具链、API smoke、浏览器闭环与最小反回归验证
- **AND** 不得把 `phase09-10` 的局部通过，直接写成整个模板复用与派生提示链已完成最终验收

## REMOVED Requirements

### Requirement: 仅凭局部构建通过或单条 API 返回即可视为 `phase09` 联调验收完成

**Reason**: 这类证据无法同时证明模板候选、预填创建、派生提示、浏览器回流链与 reread 语义在真实运行环境中一起成立。

**Migration**: 改为在 `phase09-11` 中统一执行“工具链 + API smoke + 浏览器闭环 + 最小反回归 + 边界证据”组合验证。

### Requirement: 仅凭单页浏览器点击通过即可视为模板复用与派生提示已完成正式验收

**Reason**: 单页点击不能覆盖 `templateCandidateId` 进入预填创建页、`Product Detail` 回流观察、`Dashboard / Review / ReuseSummary` reread 与“无统计变化不等于失败”的判定语义。

**Migration**: 保留浏览器点击作为必要组成部分，但必须追加字段级断言、结果页断言与反回归验证后才可给出通过结论。
