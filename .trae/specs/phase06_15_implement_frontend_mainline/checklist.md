- [x] 已明确 `phase06-15` 的目标是把 Onboarding + Sovereignty + Reuse 的前端主线推进为仓库内真实可运行实现，而不是重写 formal spec、proto 或验收脚本
- [x] 已明确 `/onboarding` 是唯一正式首轮业务入口，且 `/` 的默认进入路径继续受 `first_run_state` 控制
- [x] 已明确 Onboarding 必须完整承接 `welcome / product / repository / module / decision / complete` 六段语义
- [x] 已明确四类最小人工必填字段继续对齐 `phase06` draft-first 口径：`Product:name`、`Repository:name+url`、`Module:name`、`Decision:title+choice+reason`
- [x] 已明确四类 create 主线必须统一回收到各自 feature slice 的 `application` owner
- [x] 已明确四个既有 create 页面必须移除 page-level `useMutation` 正式主线
- [x] 已明确 Onboarding 与既有 create 页面必须共享同一套默认补值、错误归一化与 query 失效语义
- [x] 已明确 `query` 层保持纯只读，不得混入 create / update / bind / link
- [x] 已明确四类 detail route 与 detail page 必须正式承接 `fromOnboarding + onboardingStep`
- [x] 已明确 `fromOnboarding` 的回流优先级高于 `fromList / fromDashboard / fromProductDetail / fromModuleDetail`
- [x] 已明确 `ReuseSummaryRead` 以前端单一页面级 query owner 落地，不得拆成两套平行 owner
- [x] 已明确 Dashboard 必须在 `Asset Feedback` 内增加 `Reuse Snapshot` 子区域，而不是新建一级复用导航
- [x] 已明确 Module Detail 与 Product Detail 都必须挂接最小复用摘要，且不承接绑定写入逻辑
- [x] 已明确复用反馈必须"可见且可解释"，不得只显示裸统计数字
- [x] 已明确 Dashboard 必须提供稳定可见的 Export / Backup 正式用户入口
- [x] 已明确 `Backup` 的三类失败语义 `manifest_missing / coverage_incomplete / schema_mismatch` 必须在前端继续可区分
- [x] 已明确 `Export / Backup` 与 `Reuse Summary` 的前端字段语义必须继续从 `.proto -> HTTP DTO` 单向派生
- [x] 已明确 `phase06` 覆盖范围内的既有 create 页面不再保留与 Onboarding 冲突的第二套写入语义
- [x] 已明确 `phase06-15` 完成后，首轮录入、固定 mutation owner、复用反馈显示与 Export / Backup 入口都必须进入可联调状态

# 实现验收补充

## Task 1: 对齐上游与 DoD
- [x] SubTask 1.1: 对齐 `dev_plan.md#L281-L295` 的范围与 DoD
- [x] SubTask 1.2: 对齐 `phase06-06 / 07 / 09 / 12 / 14` 的单值约束
- [x] SubTask 1.3: 对齐当前 `frontend/src/routes/`、四个 canonical feature slice、`dashboard/` 真实落点

## Task 2: Onboarding 路由、根级 first-run 进入路径与 onboarding 切片
- [x] SubTask 2.1: 新增 `frontend/src/routes/index.tsx`，承接 first-run 默认进入路径
- [x] SubTask 2.2: 新增 `frontend/src/routes/onboarding.tsx`，承接 `/onboarding` 唯一路由入口与 `validateSearch`
- [x] SubTask 2.3: 新增 `frontend/src/features/onboarding/` 页面、组件、types 与只读 query owner
- [x] SubTask 2.4: 落实 `welcome / product / repository / module / decision / complete` 六段步骤编排
- [x] SubTask 2.5: 接入 Dashboard 的 `Start Onboarding / Continue Onboarding` 入口（`OnboardingCtaButton`）

## Task 3: 四类 canonical create 的固定 mutation 承接位
- [x] SubTask 3.1: 新增 `product-registry/application/use-create-draft-product.ts`
- [x] SubTask 3.2: 新增 `repository-binding/application/use-create-draft-repository.ts`
- [x] SubTask 3.3: 新增 `module-registry/application/use-create-draft-module.ts`
- [x] SubTask 3.4: 新增 `decision-center/application/use-create-draft-decision.ts`
- [x] SubTask 3.5: 统一 owner 的默认补值、错误归一化与 query 失效语义

## Task 4: 回收四类既有 create 页面中的 page-level `useMutation`
- [x] SubTask 4.1: `product-create-page.tsx` 改为消费 `useCreateDraftProduct`
- [x] SubTask 4.2: `repository-create-page.tsx` 改为消费 `useCreateDraftRepository`
- [x] SubTask 4.3: `module-create-page.tsx` 改为消费 `useCreateDraftModule`
- [x] SubTask 4.4: `decision-create-page.tsx` 改为消费 `useCreateDraftDecision`
- [x] SubTask 4.5: 验证四个 create 页面不再保留与 Onboarding 冲突的第二套正式写入语义

## Task 5: `fromOnboarding` 详情回流与返回优先级
- [x] SubTask 5.1: 四类 detail route 的 `validateSearch` 承接 `fromOnboarding + onboardingStep`（通过 `onboardingSourceSearchSchema`）
- [x] SubTask 5.2: `product-detail-page.tsx` 返回逻辑加入 `fromOnboarding` 优先级
- [x] SubTask 5.3: `repository-binding-detail-page.tsx` 返回逻辑加入 `fromOnboarding` 优先级
- [x] SubTask 5.4: `module-detail-page.tsx` 返回逻辑加入 `fromOnboarding` 优先级
- [x] SubTask 5.5: `decision-detail-page.tsx` 返回逻辑加入 `fromOnboarding` 优先级
- [x] SubTask 5.6: 验证 `fromOnboarding` 的优先级高于 `fromList / fromDashboard / fromProductDetail / fromModuleDetail`
- [x] SubTask 5.7: 新增 `onboarding/lib/onboarding-return.ts` helper 承接返回判定与 search 构造
- [x] SubTask 5.8: `/onboarding` 路由 `validateSearch` 接收 `onboardingStep`，`OnboardingPage` 用作本地兜底

## Task 6: `ReuseSummaryRead` 前端切片与 Dashboard / Detail 挂接
- [x] SubTask 6.1: 新增 `frontend/src/features/reuse-summary/types.ts`
- [x] SubTask 6.2: 新增 `frontend/src/features/reuse-summary/data/api-adapter.ts`
- [x] SubTask 6.3: 新增 `frontend/src/features/reuse-summary/data/use-reuse-summary-read.ts` 页面级 query owner
- [x] SubTask 6.4: 在 `Dashboard` 的 `Asset Feedback` 内增加 `Reuse Snapshot` 子区域（`ReuseSnapshotSection`）
- [x] SubTask 6.5: 在 `Module Detail` 挂接最小复用摘要（`ReuseSummaryInline`，scope=`module_detail`）
- [x] SubTask 6.6: 在 `Product Detail` 挂接最小复用摘要（`ReuseSummaryInline`，scope=`product_detail`）
- [x] SubTask 6.7: 验证复用反馈可见、可解释，且不升级为第二个一级导航

## Task 7: Dashboard 内的 Export / Backup 正式用户入口
- [x] SubTask 7.1: 接入 `GetExportSnapshot / ExportCoreAssets` 的前端读取与触发（`sovereignty-api-adapter.ts` + `SovereigntyPanel`）
- [x] SubTask 7.2: 接入 `GetBackupSnapshot / CreateInstanceBackup` 的前端读取与触发（`SovereigntyPanel`）
- [x] SubTask 7.3: 在 Dashboard 中提供稳定可见的 Export / Backup 入口（`sovereigntyPanel` slot）
- [x] SubTask 7.4: 验证 Backup 三类失败语义在前端仍可单值区分（`VERIFY_FAILURE_CODE_LABELS`）

## Task 8: 编译与行为验收
- [x] SubTask 8.1: 验证首轮录入从 `/` 到 `/onboarding` 再到 `/dashboard` 可走通（`/` beforeLoad + `OnboardingPage` + `OnboardingCtaButton`）
- [x] SubTask 8.2: 验证四类 create 页面与 Onboarding 共享同一套 owner 语义（Onboarding 步骤 + 四个 create 页面都消费 `application` owner）
- [x] SubTask 8.3: 验证 Dashboard / Module Detail / Product Detail 的复用反馈显示正确（`ReuseSnapshotSection` + `ReuseSummaryInline`）
- [x] SubTask 8.4: 验证 Export / Backup 入口与 snapshot 消费主线可运行（`SovereigntyPanel` + `sovereignty-api-adapter.ts`）
- [x] SubTask 8.5: 运行 `tsc -b --noEmit` 通过，无类型错误
- [x] SubTask 8.6: 运行 `npm run build` 通过，产物已生成
- [x] SubTask 8.7: 运行 `npm run lint` 通过，无新增 lint 错误
- [x] SubTask 8.8: 路由树已通过 `@tanstack/router-cli generate` 重新生成，`/onboarding` 与 `/` 已注册到 `routeTree.gen.ts`
