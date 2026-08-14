# Tasks

- [x] Task 1: 冻结 `Onboarding` 单一路由宿主与六段式页面版执行清单
  - [x] SubTask 1.1: 将 `welcome / product / repository / module / decision / complete` 六段式页面流写成单值主线，逐步明确 canonical owner、最小动作与成功后默认下一步
  - [x] SubTask 1.2: 明确 `repository / module / decision` 在"可同页闭合"与"必须 canonical handoff"两类场景下的页面落点与 reread 结果
  - [x] SubTask 1.3: 冻结 `complete` 的进入条件、结果回看边界与返回 `Dashboard` 主动作

- [x] Task 2: 冻结 canonical detail handoff 的统一返回合同
  - [x] SubTask 2.1: 明确从 `Onboarding` 进入 detail handoff 时必须携带参数级冻结的来源合同：`fromOnboarding`、`onboardingProductId`、`onboardingStep`
  - [x] SubTask 2.2: 明确 detail handoff 成功返回后的优先级：显式 step 线索优先，其次才是基于 `current_product_id` 的恢复读模型
  - [x] SubTask 2.3: 明确返回线索失效、参数缺失或不可解释时的回退规则，禁止跳回列表页或 `Dashboard`
  - [x] SubTask 2.4: 明确“进入 handoff 但未完成承接就返回/退出”时必须回到原所属未完成 step，且不得错误推进到下一步

- [x] Task 3: 冻结空态、已存在实体场景与中途中断恢复规则
  - [x] SubTask 3.1: 明确完全冷启动、部分已存在与多实体并存时的单值页面落点
  - [x] SubTask 3.2: 明确"实体已存在但关系未解释"不得视为 step 完成的规则
  - [x] SubTask 3.3: 明确从 `Dashboard` 继续、浏览器刷新重载与中途中断再进入时的统一恢复主线
  - [x] SubTask 3.4: 按 `welcome / product / repository / module / decision / complete` 六段式逐步补齐空态、失败态与 reread 落点，不再保留全局泛化描述

- [x] Task 4: 冻结失败态与移动浏览器最小降级策略
  - [x] SubTask 4.1: 明确 step 内写失败、恢复读模型失败、handoff 承接失败与返回参数失效四类失败态的停留位置与重试规则
  - [x] SubTask 4.2: 明确移动浏览器下允许降级的仅是布局与信息密度，不允许改变主线语义、默认下一步与返回合同
  - [x] SubTask 4.3: 将失败态与移动端语义写入后续浏览器验收可直接执行的观察点

- [x] Task 5: 完成 `phase10-04` 三件套自检并对齐上游
  - [x] SubTask 5.1: 复核 `spec.md` 是否直接承接 `phase10-01 / phase10-02 / phase10` 三件套，而未回退到 `phase06` draft-first 叙事
  - [x] SubTask 5.2: 复核 `tasks.md` 与 `checklist.md` 是否完整覆盖建链流、参数级返回合同、逐步空态/失败态/reread 落点、已存在实体场景、中断恢复与移动端降级
  - [x] SubTask 5.3: 确认文档语言、术语与页面语义保持单值一致，可直接进入后续实现规格

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` to `Task 4`
