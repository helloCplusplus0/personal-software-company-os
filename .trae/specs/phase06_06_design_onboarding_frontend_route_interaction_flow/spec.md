# Phase06-06 Onboarding 前端页面、路由与交互流设计 Spec

## Why

`phase06-01` 已冻结 first-run onboarding 边界、入口与最小完成条件，`phase06-02` 已冻结 draft-first 写路径与唯一 application 承接位。但如果不在进入实现前把 Onboarding 的页面文件映射、URL 语义、组件树、交互流、状态承接与回流和移动端布局降级策略设计到可直接落地的程度，后续 `phase06-15` 前端实现仍会在步骤分层、表单组织、路由守卫和 Dashboard 入口挂接位上各自猜测。

## What Changes

- 产出 `Onboarding Home / Flow` 的页面文件映射、URL 语义与组件树
- 产出首次录入、草稿保存、继续补全与完成回流的交互流设计
- 产出 cold-start / in_progress / completed 三类状态的入口承接与回流设计
- 产出根级路由入口守卫的承接位与跳转关系设计
- 产出 Dashboard `Continue Onboarding` 入口的挂接位设计
- 产出移动浏览器下的布局降级策略

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-01` first-run onboarding 边界
  - `phase06-02` draft-first 写路径
  - `phase06-05` 前端四条约束执行口径
- Affected code:
  - 新增 `frontend/src/routes/onboarding.tsx`（路由文件）
  - 新增 `frontend/src/routes/index.tsx` 或修改 `frontend/src/routes/__root.tsx`（根级路由守卫）
  - 新增 `frontend/src/features/onboarding/` 切片（页面、组件、query owner、types）
  - 修改 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`（`Start Onboarding / Continue Onboarding` 入口）
  - 修改 `frontend/src/features/dashboard/components/`（Onboarding CTA 组件与优先级判定）
  - 修改四类 canonical detail route 的 `validateSearch`：`frontend/src/routes/products/$productId.tsx`、`frontend/src/routes/repositories/$repositoryId.tsx`、`frontend/src/routes/modules/$moduleId.tsx`、`frontend/src/routes/decisions/$decisionId.tsx`
  - 修改四类 canonical detail page 的返回逻辑：`product-detail-page.tsx`、`repository-binding-detail-page.tsx`、`module-detail-page.tsx`、`decision-detail-page.tsx`
  - 后续各 feature slice 的 `application` 写入 owner（由 phase06-07 详细设计，本 spec 只定义 Onboarding 侧消费方式）

---

## 1. 页面文件映射

### 1.1 新增文件

| 文件路径 | 职责 |
| --- | --- |
| `frontend/src/routes/onboarding.tsx` | `/onboarding` 路由注册，承接搜索参数 `step` 的 `validateSearch`，指向 `OnboardingPage` |
| `frontend/src/routes/index.tsx` | `/` index 路由，`beforeLoad` 承接 first-run 判定与默认重定向 |
| `frontend/src/features/onboarding/pages/onboarding-page.tsx` | Onboarding 页面组件，编排 query、派生当前步骤、分发到步骤组件 |
| `frontend/src/features/onboarding/components/onboarding-shell.tsx` | 页面壳层：标题 + 步骤指示器 + 内容出口 |
| `frontend/src/features/onboarding/components/onboarding-stepper.tsx` | 步骤指示器：展示 5 步进度与当前步骤 |
| `frontend/src/features/onboarding/components/onboarding-step-router.tsx` | 步骤内容路由：根据当前步骤渲染对应步骤组件 |
| `frontend/src/features/onboarding/components/welcome-step.tsx` | Welcome 步骤：引导介绍 + "开始录入"按钮 |
| `frontend/src/features/onboarding/components/product-step.tsx` | Product 步骤：草稿表单或草稿摘要 |
| `frontend/src/features/onboarding/components/repository-step.tsx` | Repository 步骤：草稿表单或草稿摘要 |
| `frontend/src/features/onboarding/components/module-step.tsx` | Module 步骤：草稿表单或草稿摘要 |
| `frontend/src/features/onboarding/components/decision-step.tsx` | Decision 步骤：草稿表单或草稿摘要 |
| `frontend/src/features/onboarding/components/complete-step.tsx` | Complete 步骤：完成态 + "前往 Dashboard"按钮 |
| `frontend/src/features/onboarding/components/onboarding-draft-form.tsx` | 通用草稿表单壳：字段收集 + 提交事件 + 局部 loading/error 展示 |
| `frontend/src/features/onboarding/components/onboarding-draft-summary.tsx` | 通用草稿摘要：已创建草稿的名称/标题 + "编辑"链接 |
| `frontend/src/features/onboarding/data/use-onboarding-read.ts` | `OnboardingRead` query owner：读取 `first_run_state` 与四类对象草稿存在性 |
| `frontend/src/features/onboarding/types.ts` | Onboarding 前端类型定义 |

### 1.2 修改文件

| 文件路径 | 修改内容 |
| --- | --- |
| `frontend/src/routes/__root.tsx` | 全局导航中追加 `Onboarding` 入口（仅 `first_run_state != completed` 时可见） |
| `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` | 读取 `first_run_state`；`not_started` 时展示 `Start Onboarding` CTA，`in_progress` 时展示 `Continue Onboarding` CTA |
| `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx` | 追加 Onboarding CTA 优先级判定，覆盖 `not_started / in_progress` 两种 label 变体 |
| `frontend/src/routes/products/$productId.tsx` | 扩展 `validateSearch`，承接 `fromOnboarding` 与必要的 `onboardingStep` 返回参数 |
| `frontend/src/routes/repositories/$repositoryId.tsx` | 扩展 `validateSearch`，承接 `fromOnboarding` 与必要的 `onboardingStep` 返回参数 |
| `frontend/src/routes/modules/$moduleId.tsx` | 扩展 `validateSearch`，承接 `fromOnboarding` 与必要的 `onboardingStep` 返回参数 |
| `frontend/src/routes/decisions/$decisionId.tsx` | 扩展 `validateSearch`，承接 `fromOnboarding` 与必要的 `onboardingStep` 返回参数 |
| `frontend/src/features/product-registry/pages/product-detail-page.tsx` | 追加 `fromOnboarding` 返回逻辑；与 `fromList / fromDashboard` 做优先级合并 |
| `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx` | 追加 `fromOnboarding` 返回逻辑；与 `fromList / fromProductDetail / fromModuleDetail / fromDashboard` 做优先级合并 |
| `frontend/src/features/module-registry/pages/module-detail-page.tsx` | 追加 `fromOnboarding` 返回逻辑；与 `fromList / fromDashboard` 做优先级合并 |
| `frontend/src/features/decision-center/pages/decision-detail-page.tsx` | 追加 `fromOnboarding` 返回逻辑；与 `fromList / fromDashboard` 做优先级合并 |

---

## 2. URL 语义层

### 2.1 路由定义

| 路由路径 | 路由文件 | 页面组件 | 说明 |
| --- | --- | --- | --- |
| `/onboarding` | `routes/onboarding.tsx` | `OnboardingPage` | Onboarding 唯一正式业务入口 |
| `/` | `routes/index.tsx` | 无（重定向） | 根级 index 路由，`beforeLoad` 承接 first-run 判定 |

### 2.2 搜索参数

`/onboarding` 路由使用单一搜索参数 `step` 承接当前步骤：

| `step` 值 | 对应步骤 | 说明 |
| --- | --- | --- |
| `welcome` | Welcome | 引导介绍页，默认起始步骤 |
| `product` | Product Draft | 创建 Product 草稿 |
| `repository` | Repository Draft | 创建 Repository 草稿 |
| `module` | Module Draft | 创建 Module 草稿 |
| `decision` | Decision Draft | 创建 Decision 草稿 |
| `complete` | Complete | 完成态，引导回 Dashboard |

搜索参数校验规则（`validateSearch`）：

- `step` 为可选参数，类型为 `enum`
- 合法值：`welcome / product / repository / module / decision / complete`
- 若未提供或值非法，默认回退为 `welcome`
- 若 `OnboardingRead` 返回的草稿状态与 `step` 不一致（如 `step=product` 但 Product 草稿已存在），页面组件负责自动修正 `step` 到正确位置

### 2.3 跳转关系

| 来源 | 目标 | 触发条件 | 携带参数 |
| --- | --- | --- | --- |
| `/` (index beforeLoad) | `/onboarding` | `first_run_state = not_started` | 无 |
| `/` (index beforeLoad) | `/dashboard` | `first_run_state = in_progress` 或 `completed` | 无 |
| Dashboard Onboarding CTA | `/onboarding` | `first_run_state = not_started` 或 `in_progress` | 无 |
| 全局导航 `Onboarding` 入口 | `/onboarding` | `first_run_state != completed` | 无 |
| Onboarding Complete 步骤"前往 Dashboard" | `/dashboard` | 用户点击完成 | 无 |
| Onboarding 草稿摘要"编辑"链接 | canonical detail 页 | 用户点击编辑 | `fromOnboarding=true`, `onboardingStep=<current-step>` |

> 约束：`fromOnboarding=true` 是 Onboarding 外层来源标记，不覆盖既有 `fromList / fromDashboard` 来源模型。canonical detail 页面在检测到 `fromOnboarding=true` 时，返回路径应回到 `/onboarding`，并优先恢复 `onboardingStep` 指定的步骤；若未携带 `onboardingStep`，则回退为由 `OnboardingPage` 自动定位。

---

## 3. 组件树

```
OnboardingPage
├── OnboardingShell
│   ├── 标题区："首轮录入引导"
│   ├── OnboardingStepper
│   │   └── StepperItem × 5（Product / Repository / Module / Decision / Complete）
│   └── OnboardingStepRouter
│       ├── WelcomeStep（step=welcome）
│       │   ├── 引导文案
│       │   └── "开始录入"按钮 → navigate(step=product)
│       ├── ProductStep（step=product）
│       │   ├── pending/current → OnboardingDraftForm（字段：name）
│       │   └── completed → OnboardingDraftSummary（名称 + "编辑"链接 → /products/:id）
│       ├── RepositoryStep（step=repository）
│       │   ├── pending/current → OnboardingDraftForm（字段：name, url）
│       │   └── completed → OnboardingDraftSummary（名称 + "编辑"链接 → /repositories/:id）
│       ├── ModuleStep（step=module）
│       │   ├── pending/current → OnboardingDraftForm（字段：name）
│       │   └── completed → OnboardingDraftSummary（名称 + "编辑"链接 → /modules/:id）
│       ├── DecisionStep（step=decision）
│       │   ├── pending/current → OnboardingDraftForm（字段：title, choice, reason）
│       │   └── completed → OnboardingDraftSummary（标题 + "编辑"链接 → /decisions/:id）
│       └── CompleteStep（step=complete）
│           ├── 完成文案
│           ├── 已创建草稿摘要列表（四类对象）
│           └── "前往 Dashboard"按钮 → navigate(/dashboard)
```

### 3.1 组件职责约束

| 组件 | 职责 | 禁止 |
| --- | --- | --- |
| `OnboardingPage` | 编排 `OnboardingRead` query、派生当前步骤、分发到步骤组件 | 不得内联 `useMutation` |
| `OnboardingShell` | 页面壳层布局：标题 + 步骤指示器 + 内容出口 | 不得承接写动作 |
| `OnboardingStepper` | 展示步骤进度与当前步骤，已完成步骤可点击回看 | 不得跳过未完成步骤 |
| `OnboardingStepRouter` | 根据当前步骤渲染对应步骤组件 | 不得承接写动作 |
| `WelcomeStep` | 展示引导文案与"开始录入"按钮 | 无表单、无 mutation |
| `ProductStep / RepositoryStep / ModuleStep / DecisionStep` | 根据草稿状态渲染表单或摘要，调用对应 `application` owner 提交 | 不得内联 `useMutation`，不得自行拼装成功回流/失效刷新 |
| `OnboardingDraftForm` | 通用草稿表单壳：字段收集、提交事件、局部 loading/error 展示 | 不得自行定义 mutation，只接收 `onSubmit` 回调 |
| `OnboardingDraftSummary` | 展示已创建草稿的名称/标题与"编辑"链接 | 无写动作 |
| `CompleteStep` | 展示完成态与"前往 Dashboard"按钮 | 无表单、无 mutation |
| `useOnboardingRead` | `OnboardingRead` query owner：读取 `first_run_state` 与草稿存在性 | 纯只读，不得承接写动作 |

---

## 4. 步骤分层设计

### 4.1 步骤定义

| 步骤序号 | 步骤标识 | 步骤标题 | 说明文案 | 最小必填字段 | 系统填充字段 | 提交动作 | 可跳过 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 0 | `welcome` | 欢迎进入 PSCO | 简要介绍系统功能与首轮录入流程（Product → Repository → Module → Decision） | 无 | 无 | 无 | 是 |
| 1 | `product` | 创建你的第一个 Product | Product 是你交付给用户的产品或项目 | `name` | `description=''`, `status='active'` | `useCreateDraftProduct` | 否 |
| 2 | `repository` | 登记一个 Repository | Repository 是你的代码仓库 | `name`, `url` | `provider='manual'`, `status='active'` | `useCreateDraftRepository` | 否 |
| 3 | `module` | 创建一个 Module | Module 是可复用的代码模块 | `name` | `description=''`, `status='active'`, `capability_key=null` | `useCreateDraftModule` | 否 |
| 4 | `decision` | 记录一个 Decision | Decision 是你的技术或业务决策留痕 | `title`, `choice`, `reason` | `context=''`, `problem=''`, `alternatives=[]`, `impact=''`, `status='proposed'` | `useCreateDraftDecision` | 否 |
| 5 | `complete` | 首轮录入完成 | 你已完成首轮最小录入，可以前往 Dashboard 继续经营 | 无 | 无 | 无 | 否 |

### 4.2 步骤状态定义

每个对象步骤（Product / Repository / Module / Decision）有三种展示状态：

| 步骤状态 | 触发条件 | 展示内容 |
| --- | --- | --- |
| `pending` | 草稿尚未创建，且不是当前操作步骤 | 灰色标识，不可操作（在步骤指示器中可见但不可点击） |
| `current` | 草稿尚未创建，且是当前操作步骤 | 草稿表单（`OnboardingDraftForm`） |
| `completed` | 草稿已创建 | 草稿摘要（`OnboardingDraftSummary`），展示名称/标题 + "编辑"链接 |

### 4.3 步骤自动定位规则

`OnboardingPage` 根据 `OnboardingRead` 返回的 `first_run_state` 与草稿状态自动定位到当前步骤：

1. 若 `first_run_state = not_started`，定位到 `welcome`
2. 若 `first_run_state = completed`，定位到 `complete`
3. 若 Product 草稿不存在（且 `first_run_state != not_started`），定位到 `product`
4. 若 Product 草稿存在但 Repository 草稿不存在，定位到 `repository`
5. 若 Product + Repository 草稿存在但 Module 草稿不存在，定位到 `module`
6. 若 Product + Repository + Module 草稿存在但 Decision 草稿不存在，定位到 `decision`
7. 若四类草稿均存在但 `first_run_state != completed`，定位到 `complete`

> 约束：URL 搜索参数 `step` 的值由 `OnboardingPage` 根据上述规则自动同步。用户手动修改 `step` 时，若目标步骤的前置步骤未完成，页面必须修正 `step` 回到正确位置。

---

## 5. 交互流设计

### 5.1 首次录入流

```
用户首次进入应用（first_run_state = not_started）
  → / index beforeLoad 读取 first_run_state
  → 重定向到 /onboarding?step=welcome
  → OnboardingPage 渲染 WelcomeStep
  → 用户点击"开始录入"
  → navigate(/onboarding?step=product)
  → ProductStep 渲染 OnboardingDraftForm（字段：name）
  → 用户填写 name 并提交
  → 调用 useCreateDraftProduct（application owner）
  → 成功后：
    - mutation 失效 OnboardingRead query
    - OnboardingRead 重新读取，返回 Product 草稿已存在
    - 自动前进到 /onboarding?step=repository
  → 重复 Repository → Module → Decision
  → Decision 草稿创建成功后：
    - OnboardingRead 返回四类草稿均存在
    - 自动前进到 /onboarding?step=complete
    - first_run_state 进入 completed（由后端 OnboardingWrite 承接状态跃迁）
  → CompleteStep 展示完成态 + "前往 Dashboard"按钮
  → 用户点击"前往 Dashboard" → navigate(/dashboard)
```

### 5.2 草稿保存与继续补全流

草稿保存语义：

- 每个步骤的表单提交即创建草稿，不提供独立的"保存草稿"按钮
- 草稿创建成功后，`application` owner 负责：
  - 调用对应 `CreateDraft*` mutation
  - 失效 `OnboardingRead` query（使页面重新读取最新草稿状态）
  - 失效对应 canonical 列表 query（如 `useProductsQuery` 等）
  - 触发步骤前进（通过 `OnboardingRead` 重新读取后自动定位）

继续补全流：

```
用户中途退出（first_run_state = in_progress）
  → 后续回访应用
  → / index beforeLoad 读取 first_run_state = in_progress
  → 重定向到 /dashboard
  → Dashboard 展示 Continue Onboarding CTA
  → 用户点击"继续录入"
  → navigate(/onboarding)
  → OnboardingPage 根据 OnboardingRead 自动定位到第一个未完成步骤
  → 用户继续创建剩余草稿
  → 四类草稿全部完成后进入 Complete 步骤
```

### 5.3 完成回流

完成态语义：

- `CompleteStep` 展示四类已创建草稿的摘要列表
- 每条摘要提供"编辑"链接，跳转到对应 canonical detail 页面（携带 `fromOnboarding=true` 与 `onboardingStep`）
- `first_run_state` 进入 `completed` 后：
  - 全局导航中的 `Onboarding` 入口隐藏
  - Dashboard 中的 `Continue Onboarding` CTA 隐藏
  - 用户后续正常使用 Dashboard 与 canonical owner 主线

### 5.4 草稿编辑回流

用户从 Onboarding 草稿摘要点击"编辑"链接：

```
OnboardingDraftSummary "编辑"链接
  → navigate(/products/:id?fromOnboarding=true&onboardingStep=product)
  → canonical detail 页面展示
  → 用户编辑完成后返回
  → canonical detail 页面检测 fromOnboarding=true
  → 返回到 /onboarding（优先恢复 onboardingStep 指定的 step）
```

> 约束：`fromOnboarding=true` 不覆盖既有 `fromList / fromDashboard` 来源模型。若 canonical detail 页面同时收到 `fromOnboarding` 与 `fromList`，以 `fromList` 为优先返回路径（原生列表来源优先），`fromOnboarding` 作为次级回退。若同时收到 `fromOnboarding` 与 `fromDashboard`，以 `fromDashboard` 为优先返回路径（Dashboard 外层来源优先），`fromOnboarding` 作为次级回退。仅当 `fromOnboarding` 成为实际返回来源时，才恢复 `onboardingStep`。

### 5.5 步骤间导航

步骤间导航规则：

| 导航方向 | 触发方式 | 行为 |
| --- | --- | --- |
| 前进到下一步 | 当前步骤草稿创建成功 | 自动前进（由 OnboardingRead 重新读取后自动定位） |
| 返回上一步 | 用户点击步骤指示器中已完成的步骤 | 展示该步骤的草稿摘要（OnboardingDraftSummary），不展示表单 |
| 跳到未完成步骤 | 用户点击步骤指示器中未完成的步骤 | 禁止点击（灰色不可交互） |
| Welcome 前进 | 用户点击"开始录入"按钮 | 前进到 product 步骤（或自动定位的第一个未完成步骤） |

步骤间导航约束：

- 当前操作步骤不提供"返回上一步"按钮（用户可通过步骤指示器回看已完成步骤）
- 已完成步骤点击后展示草稿摘要，不展示空白表单（防止重复创建）
- 用户想修改已创建草稿时，必须通过草稿摘要中的"编辑"链接跳转到 canonical detail 页面修改，Onboarding 内不提供草稿编辑表单
- Welcome 步骤可跳过（点击"开始录入"即前进），其他步骤不可跳过

---

## 6. 状态承接与回流

### 6.1 cold-start（`first_run_state = not_started`）

| 维度 | 承接 |
| --- | --- |
| 根级默认进入 | `/` index `beforeLoad` 重定向到 `/onboarding` |
| Onboarding 页面 | 自动定位到 `welcome` 步骤 |
| Dashboard | 若用户通过默认主线进入则不承接主 CTA；若用户绕过根级路由直达 `/dashboard`，必须展示 `Start Onboarding` CTA（最高优先级） |
| 全局导航 | 展示 `Onboarding` 入口 |
| canonical detail | 不强制劫持（但 cold-start 用户默认不会进入 canonical detail） |

> 边界行为：若 `not_started` 用户绕过根级路由直接访问 `/dashboard`（如手动输入 URL），Dashboard 仍必须把 `/onboarding` 作为当前状态下唯一正式主 CTA，展示 `Start Onboarding`，并抑制 phase05 CTA 1-9。这样既保持“其他路由不被劫持”，也不让 cold-start 用户重新长出 `/modules/new` 的第二条正式主线。

### 6.2 in_progress 回访

| 维度 | 承接 |
| --- | --- |
| 根级默认进入 | `/` index `beforeLoad` 重定向到 `/dashboard` |
| Dashboard | 在 `DashboardPrimaryActionPanel` 中展示 `Continue Onboarding` CTA（最高优先级） |
| Onboarding 入口 | Dashboard CTA + 全局导航均可进入 `/onboarding` |
| Onboarding 页面 | 自动定位到第一个未完成步骤 |
| canonical detail | 不强制劫持，用户可自由访问 |
| canonical detail 返回 | 若 `fromOnboarding=true`，返回 `/onboarding`，并优先恢复 `onboardingStep` |

### 6.3 completed

| 维度 | 承接 |
| --- | --- |
| 根级默认进入 | `/` index `beforeLoad` 重定向到 `/dashboard` |
| Dashboard | 不展示 `Continue Onboarding` CTA |
| 全局导航 | 隐藏 `Onboarding` 入口 |
| Onboarding 页面 | 仍可通过直接访问 `/onboarding` 进入，但展示 `Complete` 步骤（只读完成态） |
| 正常使用主线 | 回到 Dashboard 与 canonical owner 页面 |

---

## 7. 根级路由入口守卫设计

### 7.1 承接位

first-run 判定由 `frontend/src/routes/index.tsx` 的 `beforeLoad` 承接：

```
routes/index.tsx
  beforeLoad:
    1. 调用 OnboardingRead（或等价 fetchFirstRunState API）读取 first_run_state
    2. 若读取成功：
       - 若 status = not_started → throw redirect({ to: '/onboarding' })
       - 若 status = in_progress 或 completed → throw redirect({ to: '/dashboard' })
    3. 若读取失败 → throw redirect({ to: '/dashboard' })（降级为正常 Dashboard 主线）
  component: 无（纯重定向路由）
```

### 7.2 设计约束

- `beforeLoad` 中不得分散到页面组件 `useEffect` 中各自判断
- `beforeLoad` 中读取 `first_run_state` 必须通过正式 `OnboardingRead` query 或等价 API，不得通过 `sessionStorage` / `localStorage` 等本地存储作为唯一判定来源
- `/` index 路由不承接任何页面渲染，只承接重定向
- 其他路由（如 `/dashboard`、`/onboarding`、`/modules/*` 等）的 `beforeLoad` 不重复承接 first-run 判定
- `OnboardingRead` 读取失败时，`beforeLoad` 必须降级重定向到 `/dashboard`，与 §8.4（Dashboard CTA 不展示）和 §12.2（全局导航 Onboarding 入口隐藏）的降级策略保持一致
- 降级到 `/dashboard` 后，用户在 Dashboard 上看到的是正常 CTA 矩阵（phase05 CTA 1-9），不展示 `Continue Onboarding` CTA
- 降级行为不得抛出整页错误页阻断用户进入应用

---

## 8. Dashboard Onboarding 入口设计

### 8.1 挂接位

Onboarding CTA 挂接在 `DashboardPrimaryActionPanel` 中，作为 `first_run_state != completed` 且用户落在 Dashboard 时的最高优先级 CTA。

### 8.2 CTA 优先级

当 Dashboard 读取到 `first_run_state` 后，CTA 优先级矩阵补充：

| 顺序 | 命中条件 | 主 CTA 目标 |
| --- | --- | --- |
| CTA 0（新增） | `first_run_state = not_started` | `/onboarding`（Start Onboarding） |
| CTA 1（新增） | `first_run_state = in_progress` | `/onboarding`（Continue Onboarding） |

> 约束：Onboarding CTA 优先于 phase05 已冻结的 CTA 1-9。`not_started` 时只展示 `Start Onboarding`，`in_progress` 时只展示 `Continue Onboarding`，二者都不得与其他 CTA 并排展示。

### 8.3 可见性控制

| `first_run_state` | Dashboard Onboarding CTA |
| --- | --- |
| `not_started` | 展示 `Start Onboarding`（仅当用户实际落在 Dashboard 时） |
| `in_progress` | 展示 `Continue Onboarding`（作为最高优先级 CTA） |
| `completed` | 不展示 |

### 8.4 Dashboard 读取 first_run_state

`DashboardHomePage` 需要读取 `first_run_state` 以决定是否展示 Onboarding CTA：

- 通过 `OnboardingRead` query 或 `DashboardOverviewRead` 扩展字段读取 `first_run_state`
- `first_run_state` 的读取状态独立于 `DashboardOverviewRead` / `FeedbackSignalRead` / `RecentActivityRead`
- `first_run_state` 读取失败时，Onboarding CTA 不展示（降级为正常 Dashboard CTA 矩阵）

---

## 9. 移动浏览器布局降级策略

### 9.1 断点定义

| 断点 | 宽度范围 | 布局策略 |
| --- | --- | --- |
| PC | `>= 1024px` | 标准布局 |
| Tablet | `768px - 1023px` | 紧凑布局 |
| Mobile | `< 768px` | 降级布局 |

### 9.2 步骤指示器降级

| 断点 | Stepper 展示 |
| --- | --- |
| PC | 水平展示全部 5 步 + 当前步骤标题 + 已完成/未完成图标 |
| Tablet | 水平展示全部 5 步，隐藏步骤标题，只展示图标 + 进度条 |
| Mobile | 只展示当前步骤序号 + 进度条（如"步骤 2/5"），不展示完整步骤条 |

### 9.3 表单降级

| 断点 | 表单布局 |
| --- | --- |
| PC | 标签与输入框水平排列，输入框最大宽度 `480px` |
| Tablet | 标签与输入框水平排列，输入框最大宽度 `100%` |
| Mobile | 标签在输入框上方，输入框全宽，最小触控高度 `44px` |

### 9.4 内容区降级

| 断点 | 内容区布局 |
| --- | --- |
| PC | 内容区最大宽度 `640px`，居中 |
| Tablet | 内容区最大宽度 `100%`，左右 padding `24px` |
| Mobile | 内容区全宽，左右 padding `16px` |

### 9.5 草稿摘要降级

| 断点 | 摘要布局 |
| --- | --- |
| PC | 名称/标题 + "编辑"链接水平排列 |
| Mobile | 名称/标题在上，"编辑"链接在下 |

### 9.6 最小支持宽度

- 最小支持宽度：`320px`
- 在 `320px` 宽度下，所有步骤的表单、摘要和按钮必须可用且不溢出

---

## 10. OnboardingRead query owner 设计

### 10.1 职责

`useOnboardingRead` 是 Onboarding 切片的唯一 query owner，承接：

- 读取 `first_run_state`（`status / current_step`）
- 读取四类对象的草稿存在性（是否已创建 + 草稿摘要）
- 提供 `queryKey` 与只读解包

### 10.2 前端消费模型

```
OnboardingReadResponse {
  first_run_state: {
    status: 'not_started' | 'in_progress' | 'completed'
    current_step: string | null
  }
  drafts: {
    product: { id: string, name: string } | null
    repository: { id: string, name: string } | null
    module: { id: string, name: string } | null
    decision: { id: string, title: string } | null
  }
}
```

与 `shared_baseline` §5.1 `first_run_state` 字段集合的映射关系：

| shared_baseline §5.1 字段 | 前端消费模型映射 | 说明 |
| --- | --- | --- |
| `status` | `first_run_state.status` | 直接映射，三态 `not_started / in_progress / completed` |
| 是否首次进入 | 由 `status = not_started` 推导 | `not_started` 即首次进入，无需独立字段 |
| 当前引导步骤 | `first_run_state.current_step` | 直接映射，表示当前应定位的步骤标识 |
| 首轮完成条件 | 由 `status = completed` 或 `drafts` 四类对象存在性推导 | `completed` 即首轮完成；`drafts` 四类对象非 null 即对应草稿已创建 |

> 约束：前端消费模型不额外新增 `is_first_entry` 或 `is_first_run_completed` 等冗余字段，这些语义由 `status` 与 `drafts` 推导得到。

### 10.3 queryKey

```
['onboarding', 'read']
```

### 10.4 约束

- `useOnboardingRead` 纯只读，不得承接任何写动作
- 前端消费模型必须与后端 HTTP JSON 形状单值一致，并与 `.proto` 语义对齐（由 phase06-10/13 承接合同落地）
- `first_run_state` 的状态枚举必须与 phase06-01 冻结的 `not_started / in_progress / completed` 三态单值一致
- 草稿摘要字段（`id / name / title`）必须从 `.proto` 单向承接，不得在前端类型中补出第二套字段

---

## 11. 写路径消费设计

### 11.1 Onboarding 侧写路径消费方式

Onboarding 步骤组件（`ProductStep / RepositoryStep / ModuleStep / DecisionStep`）不内联 `useMutation`，而是调用各 feature slice 的 `application` owner：

| 步骤 | 调用的 application owner | owner 归属 |
| --- | --- | --- |
| ProductStep | `useCreateDraftProduct` | `product-registry` slice |
| RepositoryStep | `useCreateDraftRepository` | `repository-binding` slice |
| ModuleStep | `useCreateDraftModule` | `module-registry` slice |
| DecisionStep | `useCreateDraftDecision` | `decision-center` slice |

### 11.2 application owner 职责

每个 `useCreateDraft*` application owner 负责：

- 调用对应 `CreateDraft*` mutation（通过 `api-adapter.ts` 发起 HTTP 请求）
- 成功后失效 `OnboardingRead` query（`['onboarding', 'read']`）
- 成功后失效对应 canonical 列表 query（如 `['products', 'list']`）
- 错误归一化（将 HTTP 错误转换为前端可展示的错误信息）
- 返回 `mutate` 函数与 `isPending / error` 状态给步骤组件

### 11.3 表单组件与 application owner 的边界

```
OnboardingDraftForm
  ├── 接收字段配置（字段名、类型、占位文本）
  ├── 接收 onSubmit 回调（由步骤组件从 application owner 传入）
  ├── 管理表单局部状态（字段值、局部校验）
  ├── 展示局部 loading / error
  └── 提交时调用 onSubmit（不直接调用 mutation）

ProductStep
  ├── 从 useOnboardingRead 读取 Product 草稿状态
  ├── 若草稿不存在 → 渲染 OnboardingDraftForm，传入 useCreateDraftProduct.mutate
  └── 若草稿已存在 → 渲染 OnboardingDraftSummary
```

> 约束：`OnboardingDraftForm` 不得直接调用 `useMutation`，只通过 `onSubmit` 回调提交。`ProductStep` 等步骤组件负责将 `application` owner 的 `mutate` 函数传入 `onSubmit`。

---

## 12. 全局导航设计

### 12.1 Onboarding 导航入口

在 `__root.tsx` 的全局导航中追加 `Onboarding` 入口：

| `first_run_state` | Onboarding 导航入口 |
| --- | --- |
| `not_started` | 可见，指向 `/onboarding` |
| `in_progress` | 可见，指向 `/onboarding` |
| `completed` | 隐藏 |

### 12.2 导航入口可见性控制

- 全局导航需要读取 `first_run_state` 以决定 `Onboarding` 入口的可见性
- `first_run_state` 的读取通过 `OnboardingRead` query 承接
- `first_run_state` 读取失败时，`Onboarding` 导航入口默认隐藏（降级为正常导航）

---

## ADDED Requirements

### Requirement: Onboarding 页面文件映射与组件树冻结

系统 SHALL 将 Onboarding 的页面文件映射、URL 语义与组件树冻结为本 spec §1-§3 定义的设计，后续实现不得偏离。

#### Scenario: 页面文件落点

- **WHEN** 接手者实现 Onboarding 前端主线
- **THEN** 必须按 §1.1 的文件映射新增路由与组件文件
- **AND** 必须按 §1.2 的修改清单更新既有文件
- **AND** 不得在 `onboarding` 切片之外新增第二套 Onboarding 页面或组件

#### Scenario: URL 语义

- **WHEN** 接手者实现 `/onboarding` 路由
- **THEN** 必须使用单一搜索参数 `step` 承接当前步骤
- **AND** `step` 的合法值必须为 `welcome / product / repository / module / decision / complete`
- **AND** 未提供或非法时必须回退为 `welcome`

#### Scenario: 组件树层级

- **WHEN** 接手者实现 Onboarding 组件
- **THEN** 必须按 §3 的组件树组织页面壳层、步骤指示器、步骤路由与步骤组件
- **AND** 每个组件的职责必须符合 §3.1 的职责约束
- **AND** 不得在步骤组件中内联 `useMutation`

### Requirement: Onboarding 步骤分层与自动定位冻结

系统 SHALL 将 Onboarding 的步骤分层与自动定位规则冻结为本 spec §4 定义的设计。

#### Scenario: 步骤定义

- **WHEN** 接手者实现 Onboarding 步骤
- **THEN** 必须按 §4.1 的步骤定义实现 6 个步骤（welcome / product / repository / module / decision / complete）
- **AND** 每个步骤的最小必填字段必须与 phase06-02 冻结的 draft-first 字段一致
- **AND** 每个步骤的系统填充字段必须与 phase06-02 冻结的占位值一致

#### Scenario: 步骤自动定位

- **WHEN** `OnboardingPage` 渲染时
- **THEN** 必须按 §4.3 的自动定位规则根据 `OnboardingRead` 返回的草稿状态定位到当前步骤
- **AND** URL 搜索参数 `step` 必须与自动定位结果同步
- **AND** 用户手动修改 `step` 到前置步骤未完成的位置时，页面必须修正 `step` 回到正确位置

### Requirement: Onboarding 交互流冻结

系统 SHALL 将 Onboarding 的首次录入、草稿保存、继续补全与完成回流的交互流冻结为本 spec §5 定义的设计。

#### Scenario: 首次录入流

- **WHEN** 用户首次进入应用并完成首轮录入
- **THEN** 必须按 §5.1 的流程从 Welcome 步骤开始
- **AND** 每步表单提交即创建草稿，不提供独立的"保存草稿"按钮
- **AND** 草稿创建成功后必须自动前进到下一步骤

#### Scenario: 继续补全流

- **WHEN** `first_run_state = in_progress` 的用户回访应用
- **THEN** 必须按 §5.2 的流程通过 Dashboard `Continue Onboarding` CTA 进入 `/onboarding`
- **AND** `OnboardingPage` 必须自动定位到第一个未完成步骤
- **AND** 用户完成剩余草稿后必须进入 Complete 步骤

#### Scenario: 完成回流

- **WHEN** 四类草稿全部创建完成
- **THEN** 必须按 §5.3 的流程进入 Complete 步骤
- **AND** Complete 步骤必须展示四类已创建草稿的摘要列表
- **AND** `first_run_state` 必须进入 `completed`
- **AND** `completed` 后全局导航中的 Onboarding 入口必须隐藏

### Requirement: 三类状态入口承接与回流冻结

系统 SHALL 将 cold-start / in_progress / completed 三类状态的入口承接与回流冻结为本 spec §6 定义的设计。

#### Scenario: cold-start 承接

- **WHEN** `first_run_state = not_started` 的用户进入应用
- **THEN** 必须按 §6.1 的承接表由 `/` index `beforeLoad` 重定向到 `/onboarding`
- **AND** Onboarding 页面必须自动定位到 `welcome` 步骤

#### Scenario: in_progress 承接

- **WHEN** `first_run_state = in_progress` 的用户回访应用
- **THEN** 必须按 §6.2 的承接表由 `/` index `beforeLoad` 重定向到 `/dashboard`
- **AND** Dashboard 必须在 `DashboardPrimaryActionPanel` 中展示 `Continue Onboarding` CTA
- **AND** 用户通过 CTA 或全局导航进入 `/onboarding` 时，页面必须自动定位到第一个未完成步骤

#### Scenario: completed 承接

- **WHEN** `first_run_state = completed` 的用户进入应用
- **THEN** 必须按 §6.3 的承接表由 `/` index `beforeLoad` 重定向到 `/dashboard`
- **AND** Dashboard 不得展示 `Continue Onboarding` CTA
- **AND** 全局导航必须隐藏 `Onboarding` 入口

### Requirement: 根级路由入口守卫设计冻结

系统 SHALL 将 first-run 判定的根级路由入口守卫冻结为本 spec §7 定义的设计。

#### Scenario: beforeLoad 承接

- **WHEN** 接手者实现 first-run 判定
- **THEN** 必须在 `frontend/src/routes/index.tsx` 的 `beforeLoad` 中承接
- **AND** `beforeLoad` 必须通过 `OnboardingRead` 或等价 API 读取 `first_run_state`
- **AND** `not_started` 必须重定向到 `/onboarding`
- **AND** `in_progress` 或 `completed` 必须重定向到 `/dashboard`

#### Scenario: 判定来源约束

- **WHEN** `beforeLoad` 读取 `first_run_state`
- **THEN** 必须通过正式 `OnboardingRead` query 或等价 API 读取
- **AND** 不得以 `sessionStorage` / `localStorage` 等本地存储作为唯一判定来源
- **AND** 其他路由的 `beforeLoad` 不得重复承接 first-run 判定

### Requirement: Dashboard Continue Onboarding 入口冻结

系统 SHALL 将 Dashboard `Continue Onboarding` CTA 的挂接位、优先级与可见性冻结为本 spec §8 定义的设计。

#### Scenario: CTA 挂接位

- **WHEN** `first_run_state = in_progress` 的用户进入 Dashboard
- **THEN** `DashboardPrimaryActionPanel` 必须展示 `Continue Onboarding` CTA
- **AND** 该 CTA 必须作为最高优先级 CTA（CTA 0），优先于 phase05 已冻结的 CTA 1-9
- **AND** 该 CTA 的目标路由必须为 `/onboarding`

#### Scenario: CTA 可见性

- **WHEN** `first_run_state` 发生变化
- **THEN** `first_run_state = in_progress` 时必须展示 CTA
- **AND** `first_run_state = not_started` 或 `completed` 时必须隐藏 CTA
- **AND** `first_run_state` 读取失败时必须隐藏 CTA（降级为正常 Dashboard CTA 矩阵）

### Requirement: 移动浏览器布局降级冻结

系统 SHALL 将 Onboarding 在移动浏览器下的布局降级策略冻结为本 spec §9 定义的设计。

#### Scenario: 步骤指示器降级

- **WHEN** Onboarding 在移动浏览器（`< 768px`）下渲染
- **THEN** 步骤指示器必须降级为只展示当前步骤序号 + 进度条
- **AND** 不得展示完整水平步骤条

#### Scenario: 表单降级

- **WHEN** Onboarding 草稿表单在移动浏览器下渲染
- **THEN** 标签必须在输入框上方
- **AND** 输入框必须全宽
- **AND** 最小触控高度必须为 `44px`

#### Scenario: 最小支持宽度

- **WHEN** Onboarding 在 `320px` 宽度下渲染
- **THEN** 所有步骤的表单、摘要和按钮必须可用且不溢出

### Requirement: OnboardingRead query owner 设计冻结

系统 SHALL 将 `useOnboardingRead` query owner 的职责、消费模型与约束冻结为本 spec §10 定义的设计。

#### Scenario: query owner 职责

- **WHEN** 接手者实现 `useOnboardingRead`
- **THEN** 必须承接 `first_run_state` 与四类对象草稿存在性的读取
- **AND** 必须提供 `queryKey` 与只读解包
- **AND** 不得承接任何写动作

#### Scenario: 前端消费模型

- **WHEN** 接手者定义 OnboardingRead 的前端类型
- **THEN** 必须按 §10.2 的消费模型定义 `first_run_state` 与 `drafts` 字段
- **AND** `first_run_state.status` 的枚举必须为 `not_started / in_progress / completed`
- **AND** 草稿摘要字段必须从 `.proto` 单向承接

### Requirement: Onboarding 写路径消费冻结

系统 SHALL 将 Onboarding 步骤组件的写路径消费方式冻结为本 spec §11 定义的设计。

#### Scenario: application owner 消费

- **WHEN** 接手者实现 Onboarding 步骤组件
- **THEN** 步骤组件必须调用各 feature slice 的 `application` owner（如 `useCreateDraftProduct`）
- **AND** 步骤组件不得内联 `useMutation`
- **AND** `OnboardingDraftForm` 必须通过 `onSubmit` 回调提交，不直接调用 mutation

#### Scenario: application owner 职责

- **WHEN** `useCreateDraft*` application owner 执行成功
- **THEN** 必须失效 `OnboardingRead` query
- **AND** 必须失效对应 canonical 列表 query
- **AND** 必须返回 `mutate` 函数与 `isPending / error` 状态

### Requirement: 全局导航 Onboarding 入口冻结

系统 SHALL 将全局导航中 `Onboarding` 入口的可见性控制冻结为本 spec §12 定义的设计。

#### Scenario: 导航入口可见性

- **WHEN** `first_run_state` 为 `not_started` 或 `in_progress`
- **THEN** 全局导航必须展示 `Onboarding` 入口，指向 `/onboarding`
- **AND** 当 `first_run_state = completed` 时，全局导航必须隐藏 `Onboarding` 入口

#### Scenario: 导航入口降级

- **WHEN** `first_run_state` 读取失败
- **THEN** 全局导航必须默认隐藏 `Onboarding` 入口（降级为正常导航）
- **AND** 不得在读取失败时展示 `Onboarding` 入口

## MODIFIED Requirements

### Requirement: Dashboard Primary Action Panel CTA 优先级

`DashboardPrimaryActionPanel` 在 `phase06-06` 中 SHALL 在 phase05 已冻结的 CTA 优先级矩阵前追加 CTA 0（`Continue Onboarding`），当 `first_run_state = in_progress` 时作为最高优先级 CTA 展示。

#### Scenario: CTA 0 命中

- **WHEN** `first_run_state = in_progress`
- **THEN** `DashboardPrimaryActionPanel` 必须只展示 `Continue Onboarding` CTA
- **AND** 不得并排展示 phase05 CTA 1-9 中的任一 CTA

#### Scenario: CTA 0 未命中

- **WHEN** `first_run_state != in_progress`（`not_started` / `completed` / 读取失败）
- **THEN** `DashboardPrimaryActionPanel` 必须回退到 phase05 已冻结的 CTA 1-9 优先级矩阵
- **AND** 不得展示 `Continue Onboarding` CTA

## REMOVED Requirements

### Requirement: Onboarding 使用多路由路径承接步骤（如 `/onboarding/product`、`/onboarding/repository`）

**Reason**: 多路由路径会把 Onboarding 扩写为复杂向导系统，与 phase06-01 冻结的"不形成复杂向导系统"约束冲突。使用单一搜索参数 `step` 承接步骤更简单，且不会产生多个路由文件。

**Migration**: Onboarding 后续实现统一改为：使用单一路由 `/onboarding` + 搜索参数 `?step=` 承接步骤切换，步骤切换通过 `navigate` 更新搜索参数实现，不引入多级路由路径。
