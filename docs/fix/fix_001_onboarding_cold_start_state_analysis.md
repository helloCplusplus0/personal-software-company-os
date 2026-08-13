# Fix 001 Analysis - Onboarding 冷启动欢迎页首次点击无响应

## 1. 问题摘要
- 对应问题：`fix_001`
- 问题级别：`P0`
- 是否阻断修复：是

## 2. 根因结论
- 根因一：前端 `OnboardingPage` 的 `currentStep` 优先级错误，本地 `startStep` 被服务端 `serverStep=welcome` 长期压住，导致首次点击后仍停留在欢迎页。
- 根因二：`startStep` 虽然被设计为 welcome 页点击后的本地起步兜底，但当前实现没有定义“何时让位给服务端步骤”的收敛规则，导致本地起步状态模型不完整。
- 根因三：后端 `first_run_state` 返回 `not_started + current_step=welcome` 是符合 `phase06` 冻结规格的；本问题不是后端状态推导错误，而是前端在“本地一次性起步”与“服务端事实源”之间的编排失配。

## 3. 证据链
- 页面预展示链路：
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx:71-83`
  - 当前页面同时维护 `startStep`、`focusedStep`、`serverStep` 三套步骤来源。
  - `currentStep = focusedStep ?? serverStep ?? startStep ?? 'welcome'`，导致 welcome 页点击后 `setStartStep('product')` 仍会被 `serverStep='welcome'` 抢先。
- 服务端实际创建链路：
  - `backend/internal/onboarding/service/query_service.go:75-91`
  - 当四类对象计数都为 `0` 时，服务端明确返回：
    - `status = not_started`
    - `current_step = welcome`
  - 这是 phase06 正式口径的一部分，不是异常返回。
- 数据模型映射：
  - `frontend/src/features/onboarding/data/use-onboarding-read.ts:52-60`
  - 前端只读 hook 只是把 Connect 返回解包为：
    - `status`
    - `is_first_entry`
    - `current_step`
    - `completion_progress`
  - 未发现解包层对 `welcome` 做额外篡改。
- 路由合同与本地兜底语义：
  - `frontend/src/routes/onboarding.tsx:8-12`
  - `/onboarding` 路由明确冻结：步骤状态以服务端 `first_run_state` 为事实源；`onboardingStep` 仅用于 canonical detail 页返回时的本地兜底。
  - `phase06_12` 正式规格同样冻结：
    - `spec.md:100-105`：`onboardingStep` 只承接 detail 页返回
    - `spec.md:123-134`：`not_started -> /onboarding`，并冻结 `welcome / product / repository / module / decision / complete` 六段语义
- 关键代码位置：
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx:71-83`
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx:251-256`
  - `backend/internal/onboarding/service/query_service.go:75-91`
  - `frontend/src/routes/onboarding.tsx:8-12`
- 是否存在历史脏数据：
  - 当前判断：否
  - 原因：该问题发生在冷启动交互层，不涉及历史数据写入错误或存量脏数据修补。

## 4. 影响面分析
- 首次冷启动进入：
  - 直接受影响，是当前最主要的阻断点。
- Dashboard 中的 Start / Continue Onboarding 入口：
  - 间接受影响。
  - 根级与 Dashboard 都会把 `not_started` 用户导向 `/onboarding`，进入后仍会撞上同一个 welcome 首击问题。
- Onboarding 后续四步 create 主线：
  - 数据写入链路本身未被证明错误，但用户被挡在 welcome 步骤外，导致后续主线无法触发。
- canonical detail 页回流：
  - 当前未发现直接回归问题。
  - 但修复时必须避免误伤 `focusedStep / onboardingStep` 的既有回流优先级。
- 后端 `first_run_state` 推导：
  - 当前不受影响，且不建议在本 fix 中修改。

## 5. 候选方案对比
### 方案 A
- 做法：
  - 保留现有 `startStep` 设计；
  - 将 `currentStep` 优先级调整为 `focusedStep ?? startStep ?? serverStep ?? 'welcome'`；
  - 补充 `startStep` 的收敛规则：当服务端步骤追平或超过本地起步步骤时，清空 `startStep`，让页面重新回到服务端事实源。
- 优点：
  - 变更最小，直接修复当前阻断点；
  - 不改变 `/onboarding` 路由冻结的 URL 语义；
  - 不混淆 detail 页回流专用的 `focusedStep / onboardingStep` 语义；
  - 符合“服务端为事实源，本地只做一次性起步兜底”的既有架构。
- 风险：
  - 需要小心定义“何时清空 `startStep`”，避免本地步骤长期覆盖服务端步骤；
  - 需要补充验证，确保 product 创建成功后页面能自然推进到 repository，而不是卡在 product。

### 方案 B
- 做法：
  - 弱化或删除 `startStep`；
  - 让 welcome 点击后也走 `onboardingStep=product` 这类统一本地兜底机制，复用 `focusedStep` 逻辑。
- 优点：
  - 本地步骤来源更少，表面上状态模型更统一。
- 风险：
  - 与当前 `/onboarding` 路由合同冲突：`onboardingStep` 当前只为 detail 页返回设计；
  - 容易把“从 detail 返回”和“welcome 首次开始”两种不同语义混成一类；
  - 改动范围更大，更容易误伤已冻结的回流链。

## 6. 推荐方案
- 推荐原因：
  - 推荐 **方案 A**。
  - 它最符合当前 phase06 冻结语义：服务端 `first_run_state` 仍是正式事实源，本地 `startStep` 只负责承接 welcome 页点击后的“一次性起步”。
  - 同时它能把 bug 修复限制在 `OnboardingPage` 内部编排层，不需要改动后端合同、路由合同或 detail 回流链。
- 实施边界：
  - 只修复 welcome 首次点击后的本地步骤优先级与收敛规则；
  - 保持 `focusedStep` 继续只承接 detail 页返回；
  - 保持后端 `not_started -> current_step=welcome` 不变；
  - 保持 `/onboarding` 路由 `onboardingStep` 的既有合同不扩义。
- 明确不在本次修复范围内的内容：
  - 不在本 fix 中把 onboarding 重写为完整“首轮建链引擎”；
  - 不在本 fix 中引入新的 URL 参数或第二套路由语义；
  - 不在本 fix 中改动根级 `/` 路由判定或 Dashboard CTA 合同；
  - 不在本 fix 中处理 `fix_002 / fix_003` 的决策状态链问题。

## 7. 数据修复策略
- 是否需要修历史数据：否
- 若需要，修复范围：不适用
- 若不需要，原因：
  - 问题发生在页面本地步骤编排层，不涉及存量业务数据被错误写入。

## 8. 验收标准
- 用户从冷启动欢迎页点击“开始首轮录入”后，页面必须立即进入 `product` 步骤，而不是继续停留在 `welcome`。
- 在 `product` 创建成功并失效 `onboarding-state` 后，页面必须能自然推进到服务端派生的下一步，而不是被本地 `startStep` 长期卡住。
- 从 canonical detail 页带 `onboardingStep` 返回时，既有 `focusedStep` 回流语义必须保持不变。
- 后端 `GetFirstRunState` 的合同、`/onboarding` 路由的 `validateSearch` 合同、根级 `/` 默认进入路径合同必须保持不变。

## 9. 回滚条件
- 若修复后导致 product 创建成功后页面无法继续推进到 repository / module / decision，则必须回滚；
- 若修复后导致 detail 页返回 `/onboarding` 时的 `onboardingStep` 恢复语义失效，则必须回滚。
