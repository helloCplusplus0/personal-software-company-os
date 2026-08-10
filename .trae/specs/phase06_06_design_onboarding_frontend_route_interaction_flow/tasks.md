# Tasks

- [x] Task 1: 产出 Onboarding 页面文件映射与组件树设计。把新增文件、修改文件、URL 语义与组件树层级冻结到可直接落地的程度。
  - [x] SubTask 1.1: 明确 `onboarding` 切片下的新增文件清单（路由、页面、组件、query owner、types）
  - [x] SubTask 1.2: 明确需要修改的既有文件清单（`__root.tsx`、`dashboard-home-page.tsx`、`dashboard-primary-action-panel.tsx`、四类 canonical detail route/page）
  - [x] SubTask 1.3: 明确 `/onboarding` 路由的搜索参数 `step` 的合法值、默认回退与自动修正规则
  - [x] SubTask 1.4: 明确组件树层级（OnboardingPage → Shell → Stepper + StepRouter → 6 个步骤组件）
  - [x] SubTask 1.5: 明确每个组件的职责边界与禁止事项（不得内联 useMutation、不得承接写动作）

- [x] Task 2: 产出 Onboarding 步骤分层与自动定位设计。把 6 个步骤的定义、状态与自动定位规则写成单值结论。
  - [x] SubTask 2.1: 明确 6 个步骤的标识、标题、说明文案、最小必填字段与系统填充字段
  - [x] SubTask 2.2: 明确每个对象步骤的三种展示状态（pending / current / completed）与触发条件
  - [x] SubTask 2.3: 明确步骤自动定位规则（根据 OnboardingRead 返回的草稿状态定位到当前步骤）
  - [x] SubTask 2.4: 明确 URL 搜索参数 step 与自动定位结果的同步规则

- [x] Task 3: 产出 Onboarding 交互流设计。把首次录入、草稿保存、继续补全与完成回流的交互流设计到可直接实现的程度。
  - [x] SubTask 3.1: 明确首次录入流（Welcome → Product → Repository → Module → Decision → Complete）
  - [x] SubTask 3.2: 明确草稿保存语义（表单提交即创建草稿，不提供独立保存按钮）
  - [x] SubTask 3.3: 明确继续补全流（Dashboard CTA → /onboarding → 自动定位到第一个未完成步骤）
  - [x] SubTask 3.4: 明确完成回流（Complete 步骤展示摘要 + 前往 Dashboard + first_run_state 进入 completed）
  - [x] SubTask 3.5: 明确草稿编辑回流（OnboardingDraftSummary 编辑链接 → canonical detail → 返回 /onboarding）
  - [x] SubTask 3.6: 明确四类 canonical detail route 的 `validateSearch` 与 detail page 返回逻辑必须承接 `fromOnboarding / onboardingStep`

- [x] Task 4: 产出三类状态入口承接与回流设计。把 cold-start / in_progress / completed 的入口承接与回流写成单值结论。
  - [x] SubTask 4.1: 明确 cold-start（not_started）的根级默认进入、Onboarding 定位，以及直达 `/dashboard` 时仍以 `Start Onboarding` 作为唯一主 CTA
  - [x] SubTask 4.2: 明确 in_progress 的根级默认进入、Dashboard CTA 展示、Onboarding 自动定位、canonical detail 不劫持
  - [x] SubTask 4.3: 明确 completed 的根级默认进入、Dashboard CTA 隐藏、全局导航 Onboarding 入口隐藏

- [x] Task 5: 产出根级路由入口守卫设计。把 first-run 判定的承接位、判定来源与跳转关系写成单值结论。
  - [x] SubTask 5.1: 明确 first-run 判定由 `/` index 路由的 beforeLoad 承接
  - [x] SubTask 5.2: 明确 beforeLoad 通过 OnboardingRead 或等价 API 读取 first_run_state
  - [x] SubTask 5.3: 明确 not_started 重定向到 /onboarding、in_progress / completed 重定向到 /dashboard
  - [x] SubTask 5.4: 明确 beforeLoad 不得使用 sessionStorage / localStorage 作为唯一判定来源

- [x] Task 6: 产出 Dashboard Onboarding 入口设计。把 CTA 挂接位、优先级、label 变体与可见性写成单值结论。
  - [x] SubTask 6.1: 明确 Onboarding CTA 挂接在 DashboardPrimaryActionPanel
  - [x] SubTask 6.2: 明确 `not_started -> Start Onboarding`、`in_progress -> Continue Onboarding` 都作为最高优先级 CTA
  - [x] SubTask 6.3: 明确 CTA 可见性控制（`not_started / in_progress` 命中时展示，`completed / 读取失败` 隐藏）
  - [x] SubTask 6.4: 明确 Dashboard 读取 first_run_state 的方式与独立读取状态

- [x] Task 7: 产出移动浏览器布局降级策略。把步骤指示器、表单、内容区与草稿摘要的降级规则写成单值结论。
  - [x] SubTask 7.1: 明确断点定义（PC >= 1024px、Tablet 768-1023px、Mobile < 768px）
  - [x] SubTask 7.2: 明确步骤指示器降级（Mobile 只展示当前步骤序号 + 进度条）
  - [x] SubTask 7.3: 明确表单降级（Mobile 标签在输入框上方、全宽、最小触控高度 44px）
  - [x] SubTask 7.4: 明确内容区降级（PC 最大宽度 640px 居中、Mobile 全宽 padding 16px）
  - [x] SubTask 7.5: 明确最小支持宽度 320px

- [x] Task 8: 产出 OnboardingRead query owner 设计。把职责、消费模型与约束写成单值结论。
  - [x] SubTask 8.1: 明确 useOnboardingRead 承接 first_run_state 与四类对象草稿存在性读取
  - [x] SubTask 8.2: 明确前端消费模型字段集合（first_run_state.status / current_step + drafts 四类对象）
  - [x] SubTask 8.3: 明确 queryKey 设计（['onboarding', 'read']）
  - [x] SubTask 8.4: 明确 query owner 纯只读约束与 .proto 单向承接要求

- [x] Task 9: 产出写路径消费设计。把 Onboarding 步骤组件的写路径消费方式与 application owner 职责写成单值结论。
  - [x] SubTask 9.1: 明确四个步骤组件分别调用各 feature slice 的 application owner
  - [x] SubTask 9.2: 明确 application owner 职责（mutation、失效 OnboardingRead、失效 canonical 列表、错误归一化）
  - [x] SubTask 9.3: 明确 OnboardingDraftForm 通过 onSubmit 回调提交，不直接调用 mutation
  - [x] SubTask 9.4: 明确步骤组件负责将 application owner 的 mutate 函数传入 onSubmit

- [x] Task 10: 产出全局导航设计。把 Onboarding 导航入口的可见性控制写成单值结论。
  - [x] SubTask 10.1: 明确全局导航中 Onboarding 入口的可见性（not_started / in_progress 可见、completed 隐藏）
  - [x] SubTask 10.2: 明确全局导航读取 first_run_state 的方式与读取失败时的降级

- [x] Task 11: 完成规格一致性校验。验证本次设计与 phase06-01/02/05 已冻结语义、shared_baseline、architecture_plan 保持一致。
  - [x] SubTask 11.1: 验证步骤最小必填字段与 phase06-02 冻结的 draft-first 字段一致
  - [x] SubTask 11.2: 验证 first_run_state 状态机与 phase06-01 冻结的三态一致
  - [x] SubTask 11.3: 验证写路径消费方式与 phase06-02/05 冻结的 application 唯一入口、query 纯只读约束一致
  - [x] SubTask 11.4: 验证路由入口守卫与 phase06-01 冻结的 beforeLoad 承接位一致
  - [x] SubTask 11.5: 验证 Dashboard CTA 优先级与 phase05 已冻结的 CTA 矩阵兼容，且不重新长出 cold-start 第二主线
  - [x] SubTask 11.6: 验证 `fromOnboarding` 回流链已覆盖 detail route `validateSearch` 与 detail page 返回逻辑两层改动面

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 2`
- `Task 5` depends on `Task 4`
- `Task 6` depends on `Task 4`
- `Task 7` depends on `Task 1`
- `Task 8` depends on `Task 1`
- `Task 9` depends on `Task 1` and `Task 8`
- `Task 10` depends on `Task 4`
- `Task 11` depends on `Task 1` through `Task 10`
