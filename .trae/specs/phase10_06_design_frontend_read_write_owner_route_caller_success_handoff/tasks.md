# Tasks

- [x] Task 1: 冻结 `phase10` 页面 route caller 与 owner 的一对一映射表
  - [x] SubTask 1.1: 明确 `/onboarding /dashboard /reviews/daily /products/$productId /modules/$moduleId /repositories/$repositoryId /decisions/$decisionId` 七类 route caller 的正式职责边界
  - [x] SubTask 1.2: 明确每个 route caller 只能消费哪个 page read owner 与哪个 page action owner
  - [x] SubTask 1.3: 明确 route caller 只允许保留的导航、toast、局部 UI 与错误展示职责

- [x] Task 2: 冻结 `Onboarding / Dashboard / Review` 的 read owner 与 action owner 设计
  - [x] SubTask 2.1: 明确 `useOnboardingRead` 与 `useOnboardingAction` 的分工，以及 `OnboardingPage` 必须回收的 page-level 编排
  - [x] SubTask 2.2: 明确 `useDashboardHomeRead` 与 `useDashboardPrimaryAction` 的分工，特别是 `primaryCta / secondaryCtas / actionDescriptor` 由 read owner 单值产出，避免 `DashboardHomePage` 继续承担 CTA 判定
  - [x] SubTask 2.3: 明确 `useDailyReviewRead` 与 `useReviewAction` 在 `phase10` 页面中的复用边界，避免 `Dashboard` 复制 `Review` 的动作编排

- [x] Task 3: 冻结四类 Detail 页的 page read owner 与 page action owner
  - [x] SubTask 3.1: 明确 `useProductDetailPageRead` 与 `useProductDetailActions` 的职责、输入输出与回收点，并保留 canonical panel owner 的直接消费边界
  - [x] SubTask 3.2: 明确 `useModuleDetailPageRead` 与 `useModuleDetailActions` 的职责、输入输出与回收点，并区分 CTA handoff 与普通返回
  - [x] SubTask 3.3: 明确 `useRepositoryDetailPageRead` 与 `useRepositoryDetailActions` 的职责、输入输出与回收点，并保留 canonical panel owner 的直接消费边界
  - [x] SubTask 3.4: 明确 `useDecisionDetailPageRead` 与 `useDecisionDetailActions` 的职责、输入输出与回收点，避免状态推进与返回链继续停留在页面层

- [x] Task 4: 冻结成功回流、失效刷新、错误归一化与返回链设计
  - [x] SubTask 4.1: 明确 canonical owner 与 page action owner 的失效分工，禁止页面层继续手写正式 invalidation 主线
  - [x] SubTask 4.2: 明确 `Onboarding / Dashboard / Review / Detail pages` 的 success envelope 最小字段与消费方式
  - [x] SubTask 4.3: 明确 `fromOnboarding / fromDashboard / fromList / returnTo` 等返回链参数在“普通返回”和“动作成功回流”两类场景下分别由哪个 owner 承接
  - [x] SubTask 4.4: 明确错误归一化必须发生在 application owner 内，而不是 page-local 处理

- [x] Task 5: 识别并冻结必须回收的页面级散装 mutation / CTA 编排点
  - [x] SubTask 5.1: 列出 `OnboardingPage` 当前必须回收的直接 owner import、step success 与 detail handoff 编排
  - [x] SubTask 5.2: 列出 `DashboardHomePage` 当前必须回收的 page-level 多 query、整页状态拼装与 retry 编排
  - [x] SubTask 5.3: 列出 `Product / Module / Repository / Decision Detail` 当前必须回收的 page-local invalidation、返回链与 CTA 编排

- [x] Task 6: 完成 `phase10-06` 三件套自检并对齐上游
  - [x] SubTask 6.1: 复核 `spec.md` 是否正确承接 `phase10-04 / phase10-05 / phase06-07 / phase08-06 / phase09-06`
  - [x] SubTask 6.2: 复核 `tasks.md` 与 `checklist.md` 是否完整覆盖 caller-owner 映射、canonical panel owner 边界、CTA 判定落位、成功回流与散装编排回收
  - [x] SubTask 6.3: 确认三件套已能机械回答“谁来读、谁来写、成功后谁负责回流与 reread”

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 2` and `Task 3`
- `Task 6` depends on `Task 1` to `Task 5`
