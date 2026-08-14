# Tasks

- [x] Task 1: 落实 `Onboarding` 最小后端合同与恢复锚点
  - [x] SubTask 1.1: 在 `proto/psco/onboarding/v1/onboarding.proto` 新增 `GetOnboardingChainState` 及最小响应字段
  - [x] SubTask 1.2: 生成并接线相应的 Connect / generated 代码
  - [x] SubTask 1.3: 在 `backend/internal/onboarding/service/query_service.go` 落地链路状态读取编排
  - [x] SubTask 1.4: 新增 `backend/internal/onboarding/repository/recovery_store.go`，持久化最小 `current_product_id` 恢复锚点
  - [x] SubTask 1.5: 在现有产品创建成功路径上冻结 `current_product_id`，并补齐服务层测试

- [x] Task 2: 重构 `useOnboardingRead` 为六段式链路读取 owner
  - [x] SubTask 2.1: 扩展 `frontend/src/features/onboarding/data/use-onboarding-read.ts`，组合 `GetFirstRunState + GetOnboardingChainState`
  - [x] SubTask 2.2: 产出单值的 step、step 状态、主 CTA、handoff 目标与返回链只读上下文
  - [x] SubTask 2.3: 移除对草稿摘要 search 的正式依赖，只保留过渡兼容或删除

- [x] Task 3: 新增 `useOnboardingAction` 并回收页面级写编排
  - [x] SubTask 3.1: 新增 `frontend/src/features/onboarding/application/use-onboarding-action.ts`
  - [x] SubTask 3.2: 在该 owner 内统一复用既有 canonical create / bind owner，返回稳定 success envelope
  - [x] SubTask 3.3: 将 `OnboardingPage` 中直接 import 四个 create owner、页面级 `invalidateQueries()`、页面级下一步拼装迁移到 `useOnboardingAction`
  - [x] SubTask 3.4: 为 `product / repository / module / decision` 四步补齐失败归一化与 reread 失效策略

- [x] Task 4: 将 `/onboarding` 页面正式切换为链路状态驱动
  - [x] SubTask 4.1: 重构 `frontend/src/features/onboarding/pages/onboarding-page.tsx`，只保留页面壳、步骤 UI 与 owner 消费
  - [x] SubTask 4.2: 去除 `drafts` search 主线、`navigateToDetail()` 页面级合同拼装与 `resumeServerProgress()` 之类恢复协调
  - [x] SubTask 4.3: 落实六段式 step 的空态、失败态、complete 态与默认下一步展示

- [x] Task 5: 落实 canonical handoff 的新来源合同与 detail 返回链
  - [x] SubTask 5.1: 更新 `frontend/src/features/onboarding/lib/onboarding-source-schema.ts` 与 `onboarding-return.ts`，切换到 `fromOnboarding / onboardingProductId / onboardingStep`
  - [x] SubTask 5.2: 更新 `frontend/src/routes/onboarding.tsx` 与四类 detail route 的 `validateSearch`
  - [x] SubTask 5.3: 更新 `Product / Repository / Module / Decision Detail` 页面返回逻辑，确保 handoff 完成后优先回 `/onboarding`
  - [x] SubTask 5.4: 移除 detail 页对草稿摘要 search 的正式依赖

- [x] Task 6: 完成 `Onboarding` 主线与 handoff 的浏览器级验收
  - [x] SubTask 6.1: 验证冷启动用户可完成 `product -> repository -> module -> decision -> complete`
  - [x] SubTask 6.2: 验证 `repository / module / decision` 的 canonical handoff 成功后回流到正确 step
  - [x] SubTask 6.3: 验证刷新恢复与中途中断再进入不会丢失 `current_product_id`
  - [x] SubTask 6.4: 验证本子任务没有改写 `Dashboard / Review` 的 pending 组装逻辑

- [x] Task 7: 完成 `phase10-08` 自检并同步验收证据
  - [x] SubTask 7.1: 复核实现是否对齐 `phase10-04 / 06 / 07`
  - [x] SubTask 7.2: 复核前后端 owner、合同、返回链与浏览器行为是否单值化
  - [x] SubTask 7.3: 记录剩余非目标项，避免把 `Dashboard / Review` pending 改造混入本子任务

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 1` and `Task 4`
- `Task 6` depends on `Task 1` to `Task 5`
- `Task 7` depends on `Task 1` to `Task 6`
