# Phase06-01 First-Run Onboarding 边界、入口与最小完成条件 Spec

## Why

`phase06` 要把 `Onboarding + Data Sovereignty + Reuse Awareness` 从 `phase05` 已交付的 Dashboard / Feedback 主线上推进为新的最小闭环，但第一步不能直接进入 draft-first 写路径、合同或实现，而必须先把 first-run onboarding 的页面边界、入口判定、回访语义与首轮成功会话条件冻结成单值结论。只有先收住“谁进入 onboarding、何时继续、何时完成”，后续 `first_run_state`、前端根级路由守卫、四类对象的最小录入与验收 fixture 才不会继续漂移。

## What Changes

- 冻结 `Onboarding Home / Flow` 为 `phase06-01` 当前阶段唯一新增页面主线
- 冻结 `Onboarding` 与既有 `Dashboard / Product / Repository / Module / Decision` canonical owner 页面之间的职责边界
- 冻结 `Onboarding` 的正式业务入口路由为 `/onboarding`
- 冻结 cold-start 用户、`in_progress` 回访用户与 `completed` 用户的默认进入语义
- 冻结 `first_run_state` 的最小状态集合、状态跃迁与根级入口判定承接位
- 冻结前端应用启动时的 first-run 判定由根级路由入口守卫承接，而不是分散到页面级 `useEffect`
- 冻结首轮成功会话的推荐执行顺序与最小完成条件
- 明确当前阶段不形成第二套根级宿主、不形成复杂向导系统、不把 `Onboarding` 扩写为新的业务工作台

## Impact

- Affected specs: `phase06_onboarding_sovereignty_reuse_foundation`
- Affected code: 后续 `frontend/src/routes/__root.tsx` 或等价根级路由入口守卫、新增 `frontend/src/routes/onboarding.tsx` 或等价路由文件、`frontend/src/features/onboarding/` 页面与状态承接、后续 `OnboardingRead / OnboardingWrite` owner、`first_run_state` 读写模型与验收 fixture

## ADDED Requirements

### Requirement: Phase06 Onboarding 页面主线冻结

系统 SHALL 将 `phase06-01` 的页面主线冻结为单一 `Onboarding Home / Flow`，并明确它只承接首轮引导、最小录入入口与继续补全入口，不承接新的根级宿主语义，也不替代既有 canonical owner 页面。

#### Scenario: 页面范围判定

- **WHEN** 接手者阅读 `phase06-01` 页面规格
- **THEN** 必须得到 `Onboarding Home / Flow` 是当前子任务唯一新增页面主线的单值结论
- **AND** 不得在本子任务并行引入新的 `Onboarding Workspace`、独立 `Operating Console`、第二套首页体系或复杂多宿主向导系统
- **AND** 不得把 `Product Detail`、`Repository Binding Detail / Workspace`、`Module Detail`、`Decision Center / Detail` 的既有 canonical owner 写入职责迁入 `Onboarding`

### Requirement: Onboarding 页面职责冻结

系统 SHALL 将 `Onboarding Home / Flow` 冻结为当前阶段 first-run 主线的默认承接页，承接首轮引导、四类对象最小录入入口、当前步骤展示与继续补全入口。

#### Scenario: Onboarding 页面职责判定

- **WHEN** 用户进入 `Onboarding Home / Flow`
- **THEN** 页面必须承接 first-run 引导与当前步骤展示
- **AND** 必须承接 `Product / Repository / Module / Decision` 四类对象最小录入入口
- **AND** 必须承接继续补全与完成回流入口
- **AND** 当前阶段不得在本页直接扩写为完整 `Dashboard`、完整详情编辑工作台或完整运维中心

### Requirement: Onboarding 与既有页面关系冻结

系统 SHALL 冻结 `Onboarding` 与既有 `Dashboard / Product / Repository / Module / Decision` 页面之间的职责关系，避免后续实现阶段出现两套页面语义并存。

#### Scenario: 页面关系判定

- **WHEN** 用户完成某一步最小录入后继续编辑
- **THEN** `Onboarding` 可以继续承接下一步首轮引导
- **AND** 正式详情写入仍应回到既有 canonical owner 页面或后续阶段单值写路径
- **AND** `Dashboard` 继续承接经营入口与反馈展示，不被 `Onboarding` 替代
- **AND** 当前阶段不得把 `Onboarding` 设计成所有后续编辑动作的永久宿主

### Requirement: Onboarding 正式入口路由冻结

系统 SHALL 将 `Onboarding` 的正式业务入口冻结为 `/onboarding`，并明确它是当前阶段 first-run 主线的唯一正式路由入口。

#### Scenario: 路由入口判定

- **WHEN** 接手者判断 `Onboarding` 的正式业务入口
- **THEN** 必须得到 `/onboarding` 的单值结论
- **AND** 冷启动空系统主 CTA 必须落到 `/onboarding`
- **AND** 首轮未完成用户回访主 CTA 必须为 `Continue Onboarding -> /onboarding`
- **AND** 当前阶段不得并行发明第二个 `Onboarding` 业务入口路由

### Requirement: 应用入口判定承接位冻结

系统 SHALL 将 first-run 的默认进入路径判定冻结为前端根级路由入口守卫承接，而不是分散到页面组件 `useEffect` 中各自判断。

#### Scenario: 根级入口守卫判定

- **WHEN** 应用在浏览器中首次装载并需要判断默认进入路径
- **THEN** 判定逻辑必须由前端根级路由入口守卫（`beforeLoad` 或等价根级 loader）承接
- **AND** 页面组件不得各自通过 `useEffect` 发起并行重定向判断
- **AND** 当前阶段不得在不同页面中保留多套 first-run 默认进入语义

### Requirement: `first_run_state` 最小状态机冻结

系统 SHALL 将 `first_run_state` 的最小状态集合冻结为 `not_started / in_progress / completed`，并明确最小状态跃迁语义。

#### Scenario: 状态集合判定

- **WHEN** 接手者定义当前阶段 `first_run_state`
- **THEN** 必须得到 `not_started / in_progress / completed` 三态单值结论
- **AND** 当前阶段不得额外引入会改变入口语义的第四种业务状态

#### Scenario: 状态跃迁判定

- **WHEN** 用户尚未开始任何首轮对象写入
- **THEN** `first_run_state` 必须为 `not_started`

- **WHEN** 用户已至少创建 `1` 条首轮对象记录，但 `Product / Repository / Module / Decision` 四类对象尚未全部持久化
- **THEN** `first_run_state` 必须为 `in_progress`

- **WHEN** 用户已在一次连续会话中至少各创建 `1` 条已持久化记录：`Product / Repository / Module / Decision`
- **THEN** `first_run_state` 必须进入 `completed`

### Requirement: “一次连续会话” 的操作化定义冻结

系统 SHALL 将当前阶段“`一次连续会话`”的含义冻结为同一个 first-run onboarding run，而不是要求用户必须在单个浏览器 tab、单次页面停留或一次不退出的连续交互中完成全部录入。

#### Scenario: 连续会话边界判定

- **WHEN** 接手者判断当前阶段“`一次连续会话`”的成立边界
- **THEN** 必须明确它指的是同一个 `first_run_state` 生命周期中的同一个 first-run onboarding run
- **AND** 不要求用户必须始终停留在单个浏览器 tab
- **AND** 不要求用户不能刷新页面
- **AND** 不要求用户不能中途离开再通过 `Continue Onboarding` 回到 `/onboarding`

#### Scenario: 允许退出后继续补全

- **WHEN** 用户已经创建了部分首轮对象记录，使 `first_run_state = in_progress`
- **THEN** 当前阶段必须允许用户中途退出
- **AND** 必须允许用户在后续回访中通过 `Continue Onboarding` 继续补全剩余对象
- **AND** 只要用户仍处于同一个 `first_run_state = in_progress` 的 first-run onboarding run 中，后续补齐四类对象记录后仍可进入 `completed`

#### Scenario: 不允许跨已完成 run 重新累计

- **WHEN** 用户已经满足 `completed`
- **THEN** 当前阶段不得再要求或允许通过新的回访会话重新累计同一轮 first-run 完成条件
- **AND** `completed` 之后的正常使用主线应回到 `Dashboard` 与既有 canonical owner 页面

### Requirement: Cold-Start 进入语义冻结

系统 SHALL 冻结 cold-start 用户的默认进入语义，确保空系统首次进入时不会落到错误页面。

#### Scenario: cold-start 用户进入应用

- **WHEN** 用户首次进入应用，且 `first_run_state = not_started`
- **THEN** 根级默认进入路径必须回落到 `/onboarding`
- **AND** 用户必须在 `Onboarding` 中看到首轮引导与第一步录入入口
- **AND** 当前阶段不得把 cold-start 用户默认导向 `Dashboard` 或任一 canonical detail 页面

### Requirement: `in_progress` 回访语义冻结

系统 SHALL 冻结首轮未完成用户的回访语义，避免后续实现阶段出现“强制打断”与“完全放任”两套解释。

#### Scenario: `in_progress` 用户从根级默认进入

- **WHEN** 用户回访应用，且 `first_run_state = in_progress`
- **THEN** 根级默认进入路径必须固定为 `/dashboard`
- **AND** `Dashboard` 必须提供稳定可见的 `Continue Onboarding` 入口
- **AND** 该入口必须单值回到 `/onboarding`

#### Scenario: `in_progress` 用户进入既有 canonical detail

- **WHEN** 用户回访应用，且 `first_run_state = in_progress`，并主动进入某个既有 canonical detail 页面
- **THEN** 当前阶段不要求强制劫持其离开当前页面
- **AND** 但系统仍必须提供稳定可见的 `Continue Onboarding` 返回入口

### Requirement: `completed` 用户进入语义冻结

系统 SHALL 冻结首轮已完成用户的默认进入语义，确保 `Onboarding` 不再长期劫持正常使用主线。

#### Scenario: `completed` 用户进入应用

- **WHEN** 用户进入应用，且 `first_run_state = completed`
- **THEN** 系统不得再将其默认进入路径强制导向 `/onboarding`
- **AND** 用户应继续沿用既有 `Dashboard` 与 canonical owner 页面主线
- **AND** `Onboarding` 在当前阶段仅作为已知入口存在，不再承担默认回访主线

### Requirement: 首轮成功会话最小完成条件冻结

系统 SHALL 将当前阶段的首轮成功会话冻结为一次连续会话内最少各创建 `1` 条已持久化的 `Product / Repository / Module / Decision` 记录。

#### Scenario: 首轮成功会话成立

- **WHEN** 用户在一次连续会话内最少各创建 `1` 条已持久化记录：`Product / Repository / Module / Decision`
- **AND** 该“连续会话”按当前阶段冻结语义，指同一个 `first_run_state` 生命周期中的同一个 first-run onboarding run，可包含中途退出后的回访继续补全
- **THEN** 当前阶段必须判定首轮成功会话成立
- **AND** 会话结束时用户必须能够回到 `Dashboard` 或任一 canonical owner 页面继续补全
- **AND** 允许当前阶段这些记录以 `draft-first / partial-entry` 形式存在

#### Scenario: 首轮成功会话未成立

- **WHEN** 用户只完成了部分对象创建，或四类对象中任一对象尚未持久化
- **THEN** 当前阶段不得把该会话判定为首轮成功会话
- **AND** `first_run_state` 不得进入 `completed`

### Requirement: 首轮成功会话推荐顺序冻结

系统 SHALL 将当前阶段首轮成功会话的推荐执行顺序冻结为 `Product -> Repository -> Module -> Decision`。

#### Scenario: 引导顺序判定

- **WHEN** `Onboarding` 展示当前阶段的推荐执行顺序
- **THEN** 必须得到 `Product -> Repository -> Module -> Decision` 的单值结论
- **AND** 当前阶段不得在不同页面或不同入口中并行给出不同的推荐顺序

### Requirement: 首轮成功会话的非前置条件冻结

系统 SHALL 明确当前阶段首轮成功会话不强制要求四类对象之间的绑定关系全部完成。

#### Scenario: 绑定关系非前置判定

- **WHEN** 用户已满足四类对象各自最小持久化条件
- **THEN** 当前阶段不得要求以下绑定全部完成后才算首轮成功会话：
  - `Product` 已完成全部绑定
  - `Repository` 已绑定 `Product`
  - `Module` 已映射 `Repository`
  - `Decision` 已完成对象链接
- **AND** 但 `Decision` 至少必须完成最小可持久化记录，不允许把“尚未打开 Decision 写路径”也算作 onboarding 已完成

## MODIFIED Requirements

### Requirement: Dashboard 默认进入语义

`Dashboard` 在 `phase06-01` 中 SHALL 继续作为已完成用户与正常经营主线的默认入口，但不再对 `not_started` 的 cold-start 用户承担默认进入职责。

#### Scenario: Dashboard 与 first-run 的关系判定

- **WHEN** 接手者判断 `Dashboard` 与 `Onboarding` 的默认进入关系
- **THEN** 必须明确 `Dashboard` 继续承接经营入口与反馈展示
- **AND** 必须明确 `first_run_state = not_started` 时默认进入路径优先回落到 `/onboarding`
- **AND** 必须明确 `first_run_state = in_progress` 时，根级默认进入路径固定为 `/dashboard`，且 `Dashboard` 必须提供 `Continue Onboarding` 入口回到 `/onboarding`
- **AND** 必须明确 `first_run_state = completed` 后，`Dashboard` 继续恢复为正常默认主线

## REMOVED Requirements

### Requirement: 将首轮未完成用户一律强制劫持回 `/onboarding`

**Reason**: 该解释会把 `Onboarding` 扩写为新的永久工作台，并与现有 canonical detail 页面、`Dashboard` 经营入口及“继续补全而非强制拦截”的冻结方向冲突。

**Migration**: 当前阶段统一改为：`in_progress` 用户根级默认进入路径固定为 `/dashboard`，`Dashboard` 中必须能稳定看到 `Continue Onboarding -> /onboarding`，但不要求对所有 canonical detail 页面进行强制重定向劫持。
