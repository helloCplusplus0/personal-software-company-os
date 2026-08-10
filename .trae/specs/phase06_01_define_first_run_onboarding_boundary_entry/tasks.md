# Tasks

- [x] Task 1: 冻结 `Onboarding Home / Flow` 的页面主线与页面职责。将 `Onboarding` 收敛为当前阶段唯一新增页面主线，并明确它只承接首轮引导、最小录入入口与继续补全入口。
  - [x] SubTask 1.1: 明确 `Onboarding Home / Flow` 是 `phase06-01` 当前阶段唯一新增页面主线
  - [x] SubTask 1.2: 明确 `Onboarding` 只承接首轮引导、四类对象最小录入入口、当前步骤展示与继续补全入口
  - [x] SubTask 1.3: 明确 `Onboarding` 不替代 `Dashboard` 与既有 canonical owner 页面

- [x] Task 2: 冻结 `Onboarding` 的正式入口路由与根级入口判定承接位。把 `/onboarding`、根级路由入口守卫和默认进入路径之间的关系写成单值结论。
  - [x] SubTask 2.1: 明确 `Onboarding` 的正式业务入口为 `/onboarding`
  - [x] SubTask 2.2: 明确 cold-start 主 CTA 与回访继续入口都单值回到 `/onboarding`
  - [x] SubTask 2.3: 明确应用启动时的 first-run 判定由前端根级路由入口守卫承接，而不是页面级 `useEffect`
  - [x] SubTask 2.4: 明确当前阶段不形成第二套根级宿主或第二套默认首页体系

- [x] Task 3: 冻结 `first_run_state` 的最小状态机与用户进入语义。把 `not_started / in_progress / completed` 的状态集合、状态跃迁与默认落点写清，避免后续实现出现多套入口逻辑。
  - [x] SubTask 3.1: 明确 `first_run_state` 最小状态集合为 `not_started / in_progress / completed`
  - [x] SubTask 3.2: 明确三种状态的进入条件与状态跃迁
  - [x] SubTask 3.3: 明确 cold-start 用户默认导向 `/onboarding`
  - [x] SubTask 3.4: 明确 `in_progress` 用户根级默认进入路径固定为 `/dashboard`，`Dashboard` 必须提供 `Continue Onboarding` 入口回到 `/onboarding`，但不强制劫持所有 canonical detail
  - [x] SubTask 3.5: 明确 `completed` 用户不再默认导向 `/onboarding`

- [x] Task 4: 冻结首轮成功会话的最小完成条件与推荐执行顺序。将“何时算完成 onboarding”写成单值结论，供后续写路径、状态更新与验收直接复用。
  - [x] SubTask 4.1: 明确首轮成功会话必须在一次连续会话内最少各创建 `1` 条已持久化的 `Product / Repository / Module / Decision`
  - [x] SubTask 4.1A: 明确“`一次连续会话`”指同一个 `first_run_state` 生命周期中的同一个 first-run onboarding run，允许中途退出后通过 `Continue Onboarding` 回访继续补全
  - [x] SubTask 4.2: 明确当前阶段允许 `draft-first / partial-entry`
  - [x] SubTask 4.3: 明确当前阶段不要求绑定关系全部完成后才算首轮成功会话
  - [x] SubTask 4.4: 明确 `Decision` 至少必须完成最小可持久化记录
  - [x] SubTask 4.5: 明确推荐执行顺序固定为 `Product -> Repository -> Module -> Decision`

- [x] Task 5: 完成 `phase06-01` 规格一致性校验。确认本次规格与 `phase06` 三件套、`phase05` 已交付边界和当前根级真相源保持一致。
  - [x] SubTask 5.1: 验证页面边界、入口语义与 `phase06_onboarding_sovereignty_reuse_foundation_architecture_plan.md` 一致
  - [x] SubTask 5.2: 验证状态机、默认落点与 `phase06_onboarding_sovereignty_reuse_foundation_shared_baseline.md` 一致
  - [x] SubTask 5.3: 验证任务目标与 DoD 和 `phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md` 一致
  - [x] SubTask 5.4: 验证本次规格未反向重写 `phase05` 已冻结的 Dashboard / canonical owner 边界

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
