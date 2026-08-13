# Tasks

- [x] Task 1: 冻结 `OnboardingPage` 的步骤优先级与一次性收敛规则。
  - [x] SubTask 1.1: 调整当前步骤编排，确保 `focusedStep > startStep > serverStep > welcome`
  - [x] SubTask 1.2: 为 welcome 点击产生的 `startStep` 增加明确的收敛条件，避免其长期覆盖服务端步骤
  - [x] SubTask 1.3: 保持 `focusedStep / onboardingStep` 的 detail 回流优先级不变

- [x] Task 2: 验证修复不破坏既有路由与只读合同。
  - [x] SubTask 2.1: 校验 `/onboarding` 路由 `validateSearch` 语义保持不变
  - [x] SubTask 2.2: 校验 `useOnboardingRead` 与后端 `GetFirstRunState` 合同无需改动
  - [x] SubTask 2.3: 校验根级 `/` 默认进入路径与 Dashboard Onboarding CTA 合同不受影响

- [x] Task 3: 完成 fix_001 的回归验证。
  - [x] SubTask 3.1: 验证冷启动 welcome 首次点击立即进入 `product`（前端构建通过）
  - [x] SubTask 3.2: 验证 `product` 创建成功后页面可自然推进到服务端派生的下一步（收敛规则已实现）
  - [x] SubTask 3.3: 验证 canonical detail 返回 `/onboarding` 时 `onboardingStep` 恢复语义不变（focusedStep 优先级保持最高）

# Task Dependencies
- Task 2 depends on Task 1
- Task 3 depends on Task 1
