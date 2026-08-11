# Tasks

- [x] Task 1: 对齐 `phase08-02` 的直接上游与真实 inventory
  - [x] SubTask 1.1: 对齐 `phase08-01` 已冻结的范围边界、成功标准与非目标
  - [x] SubTask 1.2: 消费 `shared_baseline` 已冻结的 `Dashboard / Decision` route、caller、query owner、application owner inventory
  - [x] SubTask 1.3: 对齐当前源码中 `/dashboard`、`DashboardHomePage`、`dashboard-source.ts` 与 `DecisionCreateRoute` 的既有事实

- [x] Task 2: 冻结 Dashboard 与 review 页面主线边界
  - [x] SubTask 2.1: 明确 `Dashboard Home` 在 `phase08` 中升级为"总览 + review 入口页"
  - [x] SubTask 2.2: 明确 `Dashboard Home` 继续保持只读编排页，不承接 review 工作会话或 mutation owner
  - [x] SubTask 2.3: 明确 daily / weekly review 作为独立页面主线，而非 Dashboard 内联工作台

- [x] Task 3: 冻结 daily / weekly review 的正式路由与最小页面职责
  - [x] SubTask 3.1: 冻结 daily / weekly review 为两条显式独立 route 宿主
  - [x] SubTask 3.2: 冻结 daily review 的输入优先级、最小页面区块与完成定义
  - [x] SubTask 3.3: 冻结 weekly review 的输入优先级、最小页面区块与完成定义
  - [x] SubTask 3.4: 明确 daily / weekly 可复用页面壳层，但不得共用同一条 route 身份

- [x] Task 4: 冻结正式 review 入口 caller 与保持只读跳转的既有 caller
  - [x] SubTask 4.1: 明确 `Dashboard` 标题行动区为当前阶段唯一正式 review 入口承接位，并说明 `DashboardPrimaryActionPanel` 只是当前源码载体
  - [x] SubTask 4.2: 明确 `FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar` 继续保持 canonical 跳转
  - [x] SubTask 4.3: 明确 `OnboardingCtaButton / SovereigntyPanel` 不升级为 review 正式入口
  - [x] SubTask 4.4: 明确不在多个页面重复长出第二套 review 入口

- [x] Task 5: 冻结 review 与 Dashboard 的单值回流路径
  - [x] SubTask 5.1: 明确 daily / weekly review 主动离开后统一回到 `/dashboard`
  - [x] SubTask 5.2: 明确只有已接入 `dashboardSourceSearchSchema` 的 canonical route 才能直接复用 `dashboard-source.ts` 与 `BackToDashboardButton`
  - [x] SubTask 5.3: 明确 `/decisions/new` 当前不属于可直接复用 Dashboard 返回链的 route，若后续需要必须单独扩展
  - [x] SubTask 5.4: 明确不得为 review 在 `/dashboard` 上新增第二套持久搜索参数事实源

- [x] Task 6: 完成 `phase08-02` 规格自检与一致性校验
  - [x] SubTask 6.1: 校验 `spec.md` 已覆盖入口形态、页面职责、路由承接位、回流路径与 caller 升级矩阵
  - [x] SubTask 6.2: 校验规格与 `phase08` 三件套、`phase08-01` 规格及当前源码事实一致
  - [x] SubTask 6.3: 校验本任务未越权冻结 `phase08-03+` 的动作 owner、合同细节或实现落点

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 2
- Task 5 depends on Task 2, Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
