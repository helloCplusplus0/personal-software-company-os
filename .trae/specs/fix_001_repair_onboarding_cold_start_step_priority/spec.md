# Onboarding 冷启动步骤优先级修复 Spec

## Why
`fix_001` 已确认当前 `/onboarding` 欢迎页存在首击无响应阻断：用户点击“开始首轮录入”后，页面仍停留在 `welcome`。问题根因已定位为前端 `OnboardingPage` 对本地起步状态与服务端步骤事实源的优先级编排错误，因此需要通过一个最小、可验证的修复规格，把修复边界、非目标与验收口径正式冻结。

## What Changes
- 修复 `OnboardingPage` 中 welcome 首次点击后的本地起步步骤优先级。
- 冻结 `startStep` 的一次性兜底语义与收敛规则，避免其长期覆盖服务端 `first_run_state.current_step`。
- 保持 detail 页返回时 `focusedStep / onboardingStep` 的既有回流语义不变。
- 保持后端 `GetFirstRunState`、`/onboarding` 路由 `validateSearch`、根级 `/` 默认进入路径合同不变。

## Impact
- Affected specs:
  - `phase06_12_onboarding_sovereignty_reuse_formal_spec`
  - `phase06_15_implement_frontend_mainline`
  - `fix_001_onboarding_cold_start_state_analysis`
- Affected code:
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/routes/onboarding.tsx`
  - `frontend/src/features/onboarding/data/use-onboarding-read.ts`
  - 验证涉及 `frontend/src/routes/index.tsx`

## ADDED Requirements
### Requirement: Welcome 首次起步本地兜底
系统 SHALL 在用户位于 `/onboarding` 欢迎页并点击“开始首轮录入”后，立即进入 `product` 步骤，而不是继续停留在 `welcome`。

#### Scenario: 冷启动欢迎页首次点击成功
- **WHEN** 用户首次进入 `/onboarding`，且服务端 `first_run_state.status = not_started`、`current_step = welcome`
- **AND** 用户点击“开始首轮录入”
- **THEN** 页面必须立即展示 `product` 步骤
- **AND** 不得继续停留在 `welcome`

### Requirement: 本地起步状态的一次性收敛
系统 SHALL 将 welcome 点击后产生的本地起步状态视为一次性兜底，而不是长期事实源。

#### Scenario: 服务端步骤追平后回归事实源
- **WHEN** 用户已从 welcome 进入 `product` 步骤
- **AND** 后续 create 成功触发 `onboarding-state` 重新读取
- **AND** 服务端 `current_step` 已推进到 `product` 或后续步骤
- **THEN** 页面必须清空 welcome 点击产生的本地起步兜底
- **AND** 后续步骤展示必须重新回到服务端 `first_run_state.current_step` 驱动

### Requirement: 起步兜底不得污染 detail 回流语义
系统 SHALL 保持 canonical detail 页返回 `/onboarding` 时的 `focusedStep / onboardingStep` 语义不变。

#### Scenario: detail 返回优先级继续成立
- **WHEN** 用户从 `Product / Repository / Module / Decision` detail 页面带 `onboardingStep` 返回 `/onboarding`
- **THEN** 页面必须继续优先恢复该一次性 `focusedStep`
- **AND** welcome 点击产生的本地起步兜底不得覆盖 detail 回流语义

## MODIFIED Requirements
### Requirement: Onboarding 页面当前步骤编排
系统 SHALL 在 `OnboardingPage` 内同时承接三类步骤来源，但必须遵守以下固定优先级：

1. detail 页返回的一次性 `focusedStep`
2. welcome 首次点击后的一次性 `startStep`
3. 服务端 `first_run_state.current_step`
4. 最终兜底 `welcome`

补充约束：

- `startStep` 只允许用于 welcome 点击后的本地起步，不得扩展为通用 URL 语义；
- 当服务端步骤已经追平或超过本地起步步骤时，`startStep` 必须被清空；
- 当前修复不得改变 `/onboarding` 路由对 `onboardingStep` 的既有解释；
- 当前修复不得要求前端自行重新定义 `first_run_state` 领域语义。

#### Scenario: welcome 起步与服务端 welcome 并存
- **WHEN** 页面当前同时存在 `serverStep = welcome` 与本地 `startStep = product`
- **THEN** 页面必须优先展示 `product`
- **AND** 不得因为服务端仍返回 `welcome` 而压回欢迎页

#### Scenario: 后续步骤由服务端继续主导
- **WHEN** 本地起步已经消费完成
- **AND** 服务端 `current_step` 已推进到 `repository / module / decision / complete`
- **THEN** 页面必须继续按服务端步骤推进
- **AND** 不得被过期的本地 `startStep` 长期卡住

## REMOVED Requirements
### Requirement: 无
**Reason**: 本次修复不删除任何既有正式能力，只修正前端步骤编排优先级。
**Migration**: 不适用。
