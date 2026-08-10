# Tasks

- [x] Task 1: 对齐 `phase06-15` 的直接上游、当前 `frontend/` 骨架与本阶段 DoD，明确这次任务是“前端主线实现”，不是重写 formal spec、proto 或验收脚本。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md#L281-L295` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase06-06 / 07 / 09 / 12 / 14` 的单值约束
  - [x] SubTask 1.3: 对齐当前 `frontend/src/routes/`、四个 canonical feature slice、`dashboard/` 真实落点

- [x] Task 2: 实现 Onboarding 路由、根级 first-run 进入路径与 `frontend/src/features/onboarding/` 主线。
  - [x] SubTask 2.1: 新增 `frontend/src/routes/index.tsx`，承接 first-run 默认进入路径
  - [x] SubTask 2.2: 新增 `frontend/src/routes/onboarding.tsx`，承接 `/onboarding` 唯一路由入口与 `validateSearch`
  - [x] SubTask 2.3: 新增 `frontend/src/features/onboarding/` 页面、组件、types 与只读 query owner
  - [x] SubTask 2.4: 落实 `welcome / product / repository / module / decision / complete` 六段步骤编排
  - [x] SubTask 2.5: 接入 Dashboard 的 `Start Onboarding / Continue Onboarding` 入口（`OnboardingCtaButton`）

- [x] Task 3: 实现四类 canonical create 的固定 mutation 承接位。
  - [x] SubTask 3.1: 新增 `product-registry/application/` 并落地 `use-create-draft-product.ts`
  - [x] SubTask 3.2: 新增 `repository-binding/application/` 并落地 `use-create-draft-repository.ts`
  - [x] SubTask 3.3: 新增 `module-registry/application/` 并落地 `use-create-draft-module.ts`
  - [x] SubTask 3.4: 新增 `decision-center/application/` 并落地 `use-create-draft-decision.ts`
  - [x] SubTask 3.5: 统一 owner 的默认补值、错误归一化与 query 失效语义

- [x] Task 4: 回收四类既有 create 页面中的 page-level `useMutation`。
  - [x] SubTask 4.1: 修改 `product-create-page.tsx` 与 `product-create-form.tsx`，改为消费 `useCreateDraftProduct`
  - [x] SubTask 4.2: 修改 `repository-create-page.tsx` 与 `repository-create-form.tsx`，改为消费 `useCreateDraftRepository`
  - [x] SubTask 4.3: 修改 `module-create-page.tsx` 与 `module-create-form.tsx`，改为消费 `useCreateDraftModule`
  - [x] SubTask 4.4: 修改 `decision-create-page.tsx` 与 `decision-create-form.tsx`，改为消费 `useCreateDraftDecision`
  - [x] SubTask 4.5: 验证四个 create 页面不再保留与 Onboarding 冲突的第二套正式写入语义

- [x] Task 5: 实现 `fromOnboarding` 详情回流与返回优先级。
  - [x] SubTask 5.1: 修改四类 detail route 的 `validateSearch`，承接 `fromOnboarding + onboardingStep`（通过 `onboardingSourceSearchSchema`）
  - [x] SubTask 5.2: 修改 `product-detail-page.tsx` 返回逻辑
  - [x] SubTask 5.3: 修改 `repository-binding-detail-page.tsx` 返回逻辑
  - [x] SubTask 5.4: 修改 `module-detail-page.tsx` 返回逻辑
  - [x] SubTask 5.5: 修改 `decision-detail-page.tsx` 返回逻辑
  - [x] SubTask 5.6: 验证 `fromOnboarding` 的优先级高于 `fromList / fromDashboard / fromProductDetail / fromModuleDetail`
  - [x] SubTask 5.7: 新增 `onboarding/lib/onboarding-return.ts` helper 承接返回判定与 search 构造
  - [x] SubTask 5.8: `/onboarding` 路由 `validateSearch` 接收 `onboardingStep`，`OnboardingPage` 用作本地兜底

- [x] Task 6: 实现 `ReuseSummaryRead` 的前端切片与 Dashboard / Detail 挂接。
  - [x] SubTask 6.1: 新增 `frontend/src/features/reuse-summary/types.ts`
  - [x] SubTask 6.2: 新增 `frontend/src/features/reuse-summary/data/api-adapter.ts`
  - [x] SubTask 6.3: 新增页面级 `ReuseSummaryRead` query owner
  - [x] SubTask 6.4: 在 `Dashboard` 的 `Asset Feedback` 内增加 `Reuse Snapshot` 子区域
  - [x] SubTask 6.5: 在 `Module Detail` 挂接最小复用摘要（`ReuseSummaryInline`，scope=`module_detail`）
  - [x] SubTask 6.6: 在 `Product Detail` 挂接最小复用摘要（`ReuseSummaryInline`，scope=`product_detail`）
  - [x] SubTask 6.7: 验证复用反馈可见、可解释，且不升级为第二个一级导航

- [x] Task 7: 实现 Dashboard 内的 Export / Backup 正式用户入口。
  - [x] SubTask 7.1: 接入 `GetExportSnapshot / ExportCoreAssets` 的前端读取与触发
  - [x] SubTask 7.2: 接入 `GetBackupSnapshot / CreateInstanceBackup` 的前端读取与触发
  - [x] SubTask 7.3: 在 Dashboard 中提供稳定可见的 Export / Backup 入口（`sovereigntyPanel` slot）
  - [x] SubTask 7.4: 验证 Backup 三类失败语义在前端仍可单值区分

- [x] Task 8: 完成 `phase06-15` 的编译与行为验收。
  - [x] SubTask 8.1: 验证首轮录入从 `/` 到 `/onboarding` 再到 `/dashboard` 可走通
  - [x] SubTask 8.2: 验证四类 create 页面与 Onboarding 共享同一套 owner 语义
  - [x] SubTask 8.3: 验证 Dashboard / Module Detail / Product Detail 的复用反馈显示正确
  - [x] SubTask 8.4: 验证 Export / Backup 入口与 snapshot 消费主线可运行
  - [x] SubTask 8.5: 运行 `tsc -b --noEmit` 通过，无类型错误
  - [x] SubTask 8.6: 运行 `npm run build` 通过，产物已生成
  - [x] SubTask 8.7: 运行 `npm run lint` 通过，无新增 lint 错误
  - [x] SubTask 8.8: 路由树已通过 `@tanstack/router-cli generate` 重新生成

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`
- `Task 7` depends on `Task 1`
- `Task 8` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, `Task 7`
