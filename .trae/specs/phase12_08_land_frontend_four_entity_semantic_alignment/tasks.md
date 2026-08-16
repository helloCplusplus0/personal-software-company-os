# Tasks

- [x] Task 1: 盘点 `phase12-08` 的上游冻结输入与实际表达面
  - [x] SubTask 1.1: 审阅 `dev_plan#L223-L237` 中 `phase12-08` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase12-04` 中 page / component / empty state / hint / next-action 的 must-change / follow-regression / no-change 矩阵
  - [x] SubTask 1.3: 审阅 `phase12-05 ~ 07` 中对共享语义来源、读路径 owner、后端真实字段与"不做项"的冻结边界
  - [x] SubTask 1.4: 审阅四个 detail page、四个 summary card、`module-next-action-bar.tsx`、`onboarding-page.tsx` 与四个列表页的当前表达

- [x] Task 2: 落实四个 primary owner detail page 的页面级语义导语
  - [x] SubTask 2.1: 在 `product-detail-page.tsx` 落地 Product 页面级语义导语，并与 `product-summary-card.tsx` 保持单值表达
  - [x] SubTask 2.2: 在 `repository-binding-detail-page.tsx` 落地 Repository 页面级语义导语，并与 `repository-summary-card.tsx` 保持单值表达
  - [x] SubTask 2.3: 在 `module-detail-page.tsx` 落地 Module 页面级语义导语，并与 `module-summary-card.tsx` / `module-next-action-bar.tsx` 保持单值表达
  - [x] SubTask 2.4: 在 `decision-detail-page.tsx` 落地 Decision 页面级语义导语，并与 `decision-detail-summary-card.tsx` 保持单值表达

- [x] Task 3: 落实四个 summary card 与 Decision 语义框架说明
  - [x] SubTask 3.1: 为 `product-summary-card.tsx`、`repository-summary-card.tsx`、`module-summary-card.tsx` 增加冻结语义标签与对齐后的字段文案
  - [x] SubTask 3.2: 为 `decision-detail-summary-card.tsx` 增加 Decision 语义框架说明，但不改动其结构化字段读模型
  - [x] SubTask 3.3: 创建 `frontend/src/features/project-context/data/shared-semantic-constants.ts` 作为唯一共享语义来源
  - [x] SubTask 3.4: 校验没有额外实现 `use-project-context-read.ts`、`entry-location-view-model.ts` 或第二套共享 data owner

- [x] Task 4: 落实 Onboarding、列表页空态与 Module 下一步动作文案
  - [x] SubTask 4.1: 将 `onboarding-page.tsx` WelcomeStep 的四步介绍改为冻结语义表达，同时保留 onboarding 引导语气
  - [x] SubTask 4.2: 将四个列表页空态改为与 Product / Repository / Module / Decision 冻结语义一致的表述
  - [x] SubTask 4.3: 将 `module-next-action-bar.tsx` 的四条 next-action 文案改为冻结语义表达
  - [x] SubTask 4.4: 校验列表页仍然不是 primary owner，改动仅限空态 / 提示文案 surface

- [x] Task 5: 完成 follow-regression / no-change surface 验证
  - [x] SubTask 5.1: 验证 Dashboard / Review 跟随页在新表达落地后没有语义回退（未修改，跟随回归）
  - [x] SubTask 5.2: 验证 `DashboardPrimaryActionPanel`、`ReviewActionFooter`、`OnboardingCtaButton` 与 `DecisionDetailSummaryCard` 状态推进区保持 no-change（未修改）
  - [x] SubTask 5.3: 验证没有改动 query owner、mutation owner、路由结构、切片目录边界与后端合同（git diff 仅涉及 14 个 must-change 文件 + 1 新增常量文件）

- [x] Task 6: 完成验收与回归验证
  - [x] SubTask 6.1: 对照 `phase12-04` 的 surface 矩阵逐项核验 must-change / follow-regression / no-change 结果
  - [x] SubTask 6.2: 运行前端校验（tsc --noEmit 通过，oxlint 0 errors / 4 pre-existing warnings）
  - [x] SubTask 6.3: 记录 `Module` 不再主要被理解为普通模块登记对象、`Decision` 不再主要被理解为孤立文本卡片的证据位
  - [x] SubTask 6.4: 校验本子任务只负责表达层与只读呈现层收口，没有落成结构重构

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4
- Task 6 depends on Task 2
- Task 6 depends on Task 3
- Task 6 depends on Task 4
- Task 6 depends on Task 5