# phase08-03 冻结 `Feedback -> Decision -> Update` 闭环的动作边界与 owner Spec

## Why

`phase08-01` 已冻结 `Operating Review Loop` 的中心主线与成功标准，`phase08-02` 已冻结 `Dashboard -> Review` 的入口、页面职责与路由承接位。接下来如果不继续把 “反馈如何进入决策、决策如何回流既有实体、真正的写路径由谁承接” 冻结成单值结论，后续实现很容易重新散落到页面、卡片、弹层和详情页局部逻辑里，最终既削弱 `Decision` 的经营中心地位，也会把 review loop 做成第二套临时任务系统。

## What Changes

- 冻结 `Feedback -> Decision -> Update` 的正式动作主线与最小成功闭环
- 冻结 review 中 `Decision` 的正式承接位、回写语义与允许动作 / 禁止动作
- 冻结 `Decision / Module / Product / Repository` 的 review result writeback 矩阵
- 冻结前端 `Review action application owner` 的单值边界
- 冻结后端 command owner 的复用边界，不新增 review-local 并列 command 主线
- 冻结 review 动作层的错误归一化边界，禁止页面级临时拼装 raw transport error

## Impact

- Affected specs:
  - `phase08_01_freeze_operating_review_loop_scope_success_non_goals`
  - `phase08_02_freeze_dashboard_review_entry_page_route_handoff`
  - `phase03_10_decision_center_formal_spec`
  - `phase04_10_product_repository_binding_formal_spec`
  - `phase06_07_design_frontend_write_path_mutation_owners`
- Affected code:
  - `frontend/src/features/dashboard/components/feedback-signal-card.tsx`
  - `frontend/src/features/decision-center/application/use-create-draft-decision.ts`
  - `frontend/src/features/decision-center/application/use-link-decision-to-target.ts`
  - `frontend/src/features/product-registry/application/use-bind-module-to-product.ts`
  - `frontend/src/features/repository-binding/application/use-bind-repository-to-product.ts`
  - `frontend/src/features/repository-binding/application/use-map-module-to-repository.ts`
  - `backend/internal/decisioncenter/service/command_service.go`
  - 后续新增的 review action application owner 承接位

## ADDED Requirements

### Requirement: `Feedback -> Decision -> Update` 必须以 `Decision` 作为经营中心，而不是唯一输出分支

系统 SHALL 将 review 中的反馈信号、代表性缺口与当前焦点统一解释为经营动作输入，并保持 `Decision` 作为经营中心；但当前阶段正式输出不得被错误收窄成“所有动作都必须先进入 `Decision`”的单一路径。

#### Scenario: review 中的反馈如何进入正式动作主线

- **WHEN** 用户在 daily / weekly review 中消费 `current_focus_signals / representative_signals / pending decisions`
- **THEN** 正式输出必须继续允许以下三类结果并存：
  - 进入单个 `Decision`
  - 进入单个实体更新或 canonical action handoff
  - 记录一条明确的下一步动作结果
- **AND** `Decision` 必须继续作为正式经营中心，而不是被弱化为可选备注
- **AND** 不得从 review 页面直接对 `Product / Module / Repository` 执行并列结构化写入
- **AND** 不得把反馈卡片上的事实信号直接升级为 review-local 任务状态

#### Scenario: 已存在决策上下文的反馈信号

- **WHEN** 某条反馈已经指向既有 `Decision` 列表或详情上下文
- **THEN** review 动作必须优先进入既有 `Decision Center` canonical 路径
- **AND** 不得绕过既有 `Decision` owner 直接在 review 页复制一套决策编辑状态
- **AND** 这不排除其他 review 输出在当前阶段通过 canonical action handoff 回流既有实体

#### Scenario: 尚无决策上下文的反馈信号

- **WHEN** 某条反馈需要形成新的经营判断，但当前不存在对应的既有 `Decision` 上下文
- **THEN** review 动作可以形成以下最小正式结果之一：
  - 创建 `decision draft` 并回流既有 `Decision Detail`
  - 进入单个实体更新或 canonical action handoff
  - 记录一条明确的下一步动作结果
- **AND** 若当前动作被判定为需要形成新的正式经营判断，必须先创建 `decision draft`
- **AND** 不得把“反馈已被处理”仅停留在 review 页的局部勾选或临时文案里

### Requirement: `Decision` 必须继续作为 review loop 的正式写入中心

系统 SHALL 将 `Decision Center` 的既有 canonical owner 冻结为 review loop 中所有决策写入动作的正式承接位。

#### Scenario: Decision 的允许动作

- **WHEN** 后续 `/spec` 或实现定义 review 中与 `Decision` 相关的正式动作
- **THEN** 当前阶段允许的动作只包括：
  - 新建 `decision draft`
  - 进入既有 `Decision Detail`
  - 在既有 `Decision Detail` 内完成目标关联
- **AND** 不得在当前阶段扩写成 review 内联决策编辑工作台

#### Scenario: Decision 的禁止动作

- **WHEN** review 页面、review 面板或 Dashboard 卡片尝试承接决策写入
- **THEN** 不得在这些位置直接内联一套新的 decision 表单 owner
- **AND** 不得在 review 记录中复制并持久化并列 decision 状态
- **AND** 不得把 `Decision` 弱化为“review 备注”或“可选附属记录”

### Requirement: `Decision -> Module / Product / Repository` writeback 矩阵必须单值化

系统 SHALL 按共享基线中已冻结的 writeback matrix，将 review 结果回流路径冻结为“`Decision` 正式承接 + 至少一种实体回流落地”的单值矩阵。

#### Scenario: Module writeback

- **WHEN** review 结果需要回流 `Module`
- **THEN** 正式回流方式必须是：
  - 通过既有 `Decision -> Module` link 写路径完成正式关联；或
  - 从 review / decision 导航到既有 `Module Detail` canonical 更新入口
- **AND** 当前阶段不得在 review 页直接内联 `Module` 结构化写入
- **AND** 不得长出 review-local module 状态

#### Scenario: Product writeback

- **WHEN** review 结果需要回流 `Product`
- **THEN** 当前阶段最小正式落地必须优先通过 canonical action handoff 回到既有 `Product Detail` 或列表页
- **AND** 若后续确需在本阶段直接执行 `Product` 相关写入，只能复用既有 `product-registry` application owner
- **AND** 不得为 review 单独创建并列 product 写路径
- **AND** 不得在 review 记录中持久化影子 product 状态

#### Scenario: Repository writeback

- **WHEN** review 结果需要回流 `Repository`
- **THEN** 当前阶段最小正式落地必须优先通过 canonical action handoff 回到既有 `Repository Detail` 或列表页
- **AND** 若后续确需在本阶段直接执行 `Repository` 相关写入，只能复用既有 `repository-binding` application owner
- **AND** 不得为 review 单独创建并列 repository 写路径
- **AND** 不得在 review 记录中持久化影子 repository 状态

#### Scenario: 最小成功闭环的写回要求

- **WHEN** `phase08` 进入实现与验收
- **THEN** `review result writeback` 的最小必做必须是：`Decision` 正式承接 + 至少一种实体回流落地
- **AND** `Product / Repository` 在当前阶段允许先以 canonical action handoff 落地
- **AND** 不强制要求本阶段新增 `Product / Repository` 专用直写 API

### Requirement: 前端 review 写路径必须收敛为单一 `Review action application owner`

系统 SHALL 将 review 中的正式写动作冻结到单一 `Review action application owner` 承接位，而不是允许 review page、review panel、decision card、feedback card 各自直接持有 mutation 组合。

#### Scenario: Review action owner 的职责边界

- **WHEN** 后续 `/spec` 或实现定义前端 review 写路径
- **THEN** 必须存在单一 `Review action application owner`
- **AND** 该 owner 只承接以下职责：
  - 选择进入 `Decision` 的正式动作
  - 调用既有 canonical application owner 或 canonical route handoff
  - 统一承接成功回流、query 失效与错误归一化
- **AND** route / page / 展示组件不得直接成为新的 mutation owner

#### Scenario: Review action owner 复用既有 canonical owner

- **WHEN** review 动作需要执行写入
- **THEN** 必须优先复用既有 canonical application owner：
  - `useCreateDraftDecision`
  - `useLinkDecisionToTarget`
  - `useBindModuleToProduct`
  - `useBindRepositoryToProduct`
  - `useMapModuleToRepository`
- **AND** review action owner 只允许编排这些既有 owner，不得复制第二套等价 mutation 实现

#### Scenario: 页面级临时编排禁止项

- **WHEN** review 页面、review 面板或 Dashboard 卡片实现动作按钮
- **THEN** 不得直接在页面级内联 `useMutation`
- **AND** 不得在组件内零散拼装失效刷新、成功跳转与错误归一化
- **AND** 不得让 `FeedbackSignalCard`、`RecentActivityItemCard` 等只读 caller 演化为正式写路径 owner

### Requirement: 后端 command owner 必须复用既有领域 command 主线

系统 SHALL 将 review loop 的后端写入承接边界冻结为“复用既有领域 command owner”，而不是新增 review-local 并列 command service。

#### Scenario: Decision 写入的后端 owner

- **WHEN** review 动作需要创建 `Decision` 或完成 `Decision -> Module` 关联
- **THEN** 必须复用 `Decision Center` 既有 command owner
- **AND** 当前已冻结的最小后端承接位是 `backend/internal/decisioncenter/service/command_service.go`
- **AND** 不得为了 review 额外创建并列的 `review decision command service`

#### Scenario: Product / Repository / Module 相关写入的后端 owner

- **WHEN** review 动作后续需要触发 `Product / Repository / Module` 相关写入
- **THEN** 必须复用各自既有领域切片的 command owner / application owner 主线
- **AND** 不得创建 `review product command`、`review repository command` 或 `review module command` 作为并列业务主线
- **AND** 若当前阶段只需 canonical action handoff，则无需新增后端 review command

### Requirement: review 动作错误归一化必须收敛在 owner 边界

系统 SHALL 将 review 动作中的错误归一化冻结在 owner 边界，而不是让页面组件直接消费 raw transport / raw Connect error。

#### Scenario: 前端错误归一化边界

- **WHEN** `Review action application owner` 调用既有 canonical application owner 或 route handoff
- **THEN** 必须由该 owner 统一归一化错误并向页面暴露稳定的 review-facing error 语义
- **AND** page / panel / card 只负责展示归一化后的错误
- **AND** 不得在页面层直接分支解析 raw `ConnectError` 或 transport error

#### Scenario: 后端错误语义边界

- **WHEN** review 动作命中既有 `Decision / Product / Repository / Module` command 主线
- **THEN** 后端仍必须保持既有 domain error → proto / Connect error 的单值映射
- **AND** review 不得引入第二套错误码体系
- **AND** 不得为了 review 写路径在 handler 层临时拼装新的 JSON 错误语义

## MODIFIED Requirements

### Requirement: `FeedbackSignalCard` 与 `RecentActivityItemCard` 的职责解释

`FeedbackSignalCard` 与 `RecentActivityItemCard` 在 `phase08` 中 SHALL 继续保持“事实输入与 canonical 导航 caller”的身份，而不是升级为正式写路径 owner。

#### Scenario: Dashboard 卡片与 review owner 的边界

- **WHEN** 后续实现 review loop
- **THEN** `FeedbackSignalCard` 与 `RecentActivityItemCard` 可以继续作为事实跳转 caller
- **AND** review 正式写路径必须迁移到 review 页与 `Review action application owner`
- **AND** 不得把这些卡片直接改造成并列的 review 写入入口

## REMOVED Requirements

### Requirement: 页面级临时 mutation 编排可以作为 review 落地方式

**Reason**: 这会直接违背 `phase08` 已冻结的 `application owner` 单值原则，并让 `Feedback -> Decision -> Update` 的动作编排重新散落到页面、卡片与弹层局部逻辑中。  
**Migration**: review 正式写路径统一收敛到单一 `Review action application owner`；该 owner 只编排既有 canonical owner 与 canonical action handoff，不复制第二套 mutation / command 实现。
