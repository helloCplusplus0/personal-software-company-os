# phase12-08 落实前端四实体语义一致性收口 Spec

## Why
`phase12-04` 已冻结四实体语义承接矩阵，`phase12-05 ~ 07` 已分别冻结共享 owner、前端读路径边界与后端真实字段能力。`phase12-08` 需要把这些冻结结论真正落到 Web 表达层与只读呈现层上，让用户在详情页、摘要卡片、空态和说明文案里稳定感知 `Product / Repository / Module / Decision` 的正式语义。

## What Changes
- 冻结 `phase12-08` 只负责表达层与只读呈现层收口，不做结构重构、不改写路径 owner、不新开第二套读写边界
- 冻结四个 primary owner detail page 的页面级语义导语、summary card 语义标签与必要说明文案的落地范围
- 冻结 `Onboarding WelcomeStep`、四个列表页空态、`ModuleNextActionBar` 等关键表达面如何回收到冻结语义
- 冻结 `frontend/src/features/project-context/data/shared-semantic-constants.ts` 作为唯一共享语义来源的落地边界；**BREAKING**：不允许继续在多个页面硬编码互相不一致的四实体解释
- 冻结 follow-regression / no-change surface 的实施约束，确保本子任务不借机改写切片边界、query owner、mutation owner 或后端合同

## Impact
- Affected specs:
  - `phase12_semantic_alignment_and_readonly_consumption_foundation`
  - `phase12_04_design_frontend_four_entity_semantic_alignment`
  - `phase12_05_design_readonly_consumption_shared_entry`
  - `phase12_06_design_frontend_read_path_owner_shared_summary_reread`
  - `phase12_07_design_backend_contract_export_shared_read_views`
- Affected code:
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/product-registry/components/product-summary-card.tsx`
  - `frontend/src/features/repository-binding/components/repository-summary-card.tsx`
  - `frontend/src/features/module-registry/components/module-summary-card.tsx`
  - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
  - `frontend/src/features/module-registry/components/module-next-action-bar.tsx`
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/features/product-registry/pages/product-list-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-list-page.tsx`
  - `frontend/src/features/module-registry/pages/module-list-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-list-page.tsx`
  - `frontend/src/features/project-context/data/shared-semantic-constants.ts`

## ADDED Requirements

### Requirement: 冻结 phase12-08 的实现职责与非目标
系统 SHALL 将 `phase12-08` 冻结为“前端四实体语义一致性表达收口”的实现子任务，只承接 `phase12-04` 已冻结的 must-change / follow-regression surface 与 `phase12-06` 已冻结的共享语义来源边界。

当前子任务至少必须满足：

- 至少完成 `repositories/$repositoryId`、`products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 四类 detail page 的 primary owner 收口；
- 落实四个 summary card、四个 detail page 页面级语义导语、`Onboarding WelcomeStep`、四个列表页空态、`ModuleNextActionBar` 的冻结语义；
- 若需要共享四实体语义短语，只允许落到 `frontend/src/features/project-context/data/shared-semantic-constants.ts`；
- 不改写 `phase12-05 / 06 / 07` 已冻结的 read owner、query owner、resolver 或后端合同结论。

补充冻结：

- `phase12-08` 不负责结构重构，不新增共享 query，不实现 `use-project-context-read.ts` 或 `entry-location-view-model.ts`；
- `phase12-08` 不改变写路径 owner、mutation owner、切片目录边界与路由结构；
- `phase12-08` 可新增或消费共享语义常量，但不得把 `frontend/src/features/project-context/` 扩张成第二套页面 data 层。

#### Scenario: 执行者开始承接 phase12-08
- **WHEN** 执行者开始实现 `phase12-08`
- **THEN** 能明确本任务只负责表达层与只读呈现层收口
- **AND** 不会借机改写 `phase12-05 / 06 / 07` 的结构性结论

### Requirement: 落实四个 primary owner detail page 的正式语义导语
系统 SHALL 在四个 primary owner detail page 上显式落地页面级语义导语，并与各自 summary card 形成单值表达，不再让页面只作为被动容器。

当前阶段至少必须实现：

1. `product-detail-page.tsx` 落地 “经营目标与交付容器” 页面级导语
2. `repository-binding-detail-page.tsx` 落地 “代码仓库身份对象与项目锚点” 页面级导语
3. `module-detail-page.tsx` 落地 “可复用能力资产” 页面级导语
4. `decision-detail-page.tsx` 落地 “规则、约束、选择与依据的索引对象” 页面级导语

补充冻结：

- 页面级导语必须与 `phase12-04 §2.1A` 的页面级承接位一致；
- 页面标题仍由真实业务字段承接，不允许把数据标题替换成抽象概念标题；
- 导语可引用共享语义常量，但渲染位置与排版继续保留在各切片页面内。

#### Scenario: 用户进入任一 primary owner detail page
- **WHEN** 用户打开 Product / Repository / Module / Decision 详情页
- **THEN** 页面级导语与摘要卡片一起稳定表达该实体的冻结语义
- **AND** 用户不会只把页面看成普通登记详情或孤立文本展示

### Requirement: 落实四个 summary card 与 Decision 语义框架说明
系统 SHALL 在四个 summary card 中落地 `phase12-04` 已冻结的语义标签与说明，使摘要卡片本身成为实体正式语义的稳定承接位。

当前阶段至少必须实现：

- `product-summary-card.tsx` 增加 Product 语义标签与对齐后的字段文案
- `repository-summary-card.tsx` 增加 Repository 语义标签，并保持 URL/provider 作为身份属性
- `module-summary-card.tsx` 增加 Module 语义标签，并让描述/时间等字段更贴近“可复用能力资产”语义
- `decision-detail-summary-card.tsx` 增加 Decision 语义框架说明，但不改写其现有结构化字段区

补充冻结：

- `DecisionDetailSummaryCard` 的状态推进 CTA 仍属于业务动作区，不纳入四实体语义标签承接位；
- summary card 只收口表达，不新增第二套数据 shape；
- 若字段标题需要改名，只能围绕 `phase12-04` 已冻结的表达改清晰度，不得发明新实体定义。

#### Scenario: 用户浏览摘要卡片
- **WHEN** 用户查看四个实体的摘要卡片
- **THEN** 能直接读到实体的正式语义标签或语义框架说明
- **AND** 不会把 Module 主要理解成普通模块登记对象，也不会把 Decision 主要理解成孤立文本卡片

### Requirement: 落实 Onboarding、列表页空态与 Module 下一步动作文案
系统 SHALL 将 `phase12-04` 已冻结的 must-change 文案 surface 真正落到 Onboarding Welcome、四个列表页空态与 `ModuleNextActionBar`，回收当前偏 CRUD 或偏登记表单的表达。

当前阶段至少必须实现：

- `onboarding-page.tsx` WelcomeStep 把四步介绍改成冻结语义表达，不再写成“创建一个模块/记录一条决策”的普通录入步骤
- 四个列表页空态改为与 Product / Repository / Module / Decision 冻结语义一致的表述
- `module-next-action-bar.tsx` 四条 next-action 文案回到“可复用能力资产 / 经营目标 / 代码仓库身份对象与项目锚点 / 规则与决策记录”的正式说法

补充冻结：

- `DashboardPrimaryActionPanel`、`ReviewActionFooter`、`OnboardingCtaButton` 继续保持 no-change；
- 列表页仍不是 primary owner；其 `must-change` 仅表示文件内的空态与提示位必须落盘，不代表 owner 身份升级；
- Onboarding Welcome 的引导语气可以保留，但不得再把四实体整体解释成 CRUD 流程。

#### Scenario: 用户查看 onboarding 或列表空态
- **WHEN** 用户进入 WelcomeStep 或遇到四个列表页空态
- **THEN** 能看到与冻结语义一致的表达
- **AND** 不会把 PSCO 的核心对象误读为普通表单登记项

### Requirement: 冻结共享语义来源的实施边界
系统 SHALL 要求 `phase12-08` 只在必要处落地 `shared-semantic-constants.ts` 这一类共享语义来源，不把共享语义来源误实现成共享 query、共享页面布局或第二数据层。

当前阶段至少必须明确：

- 共享语义来源只包括四实体冻结语义标签、实体类型到标签的映射、以及稳定复用的入口描述常量；
- 页面如何排版、空态如何组合、导语放在哪里，继续由各切片本地渲染；
- `phase12-08` 不实现 `use-project-context-read.ts` 与 `entry-location-view-model.ts`，这些仍属于 `phase12-06` 的后续能力位；
- 不允许在页面内继续硬编码另一套互相冲突的四实体解释。

#### Scenario: 实现者需要复用四实体语义短语
- **WHEN** 实现者在多个页面 / 组件中需要复用同一条四实体语义短语
- **THEN** 只会从 `shared-semantic-constants.ts` 取值
- **AND** 不会顺手新建另一套共享只读 data owner 或页面私有常量组

### Requirement: 冻结 follow-regression 与 no-change surface 的实施约束
系统 SHALL 要求 `phase12-08` 在落实 must-change surface 时同步完成 follow-regression 验证，但不得把 no-change surface 偷渡成扩 scope 改造。

当前阶段至少必须显式验证：

- Dashboard / Review 跟随页在新的四实体表达落地后不出现语义回退；
- `DecisionDetailSummaryCard` 的状态推进区、`DashboardPrimaryActionPanel`、`ReviewActionFooter`、`OnboardingCtaButton` 继续 no-change；
- 不引入新的 query、transport、adapter、mutation 或后端合同变动。

#### Scenario: 实现者完成表达层收口
- **WHEN** `phase12-08` 实现完成
- **THEN** must-change surface 已落地、follow-regression surface 已验证
- **AND** no-change surface 没有被无关改写

## MODIFIED Requirements

### Requirement: phase12-08 与 phase12-04 ~ 07 的关系表达
`phase12-08` 的职责 SHALL 在 `phase12` 三件套与相关 spec 中继续对齐为同一口径：

- `phase12-04` 冻结“哪些 surface 必须改、哪些跟随回归、哪些不改”；
- `phase12-05` 冻结共享 owner 与三类页面承接矩阵；
- `phase12-06` 冻结共享语义来源与前端读路径 / reread 边界；
- `phase12-07` 冻结后端真实字段、L3 单向映射与“不做”的后端候选；
- `phase12-08` 只把这些冻结结论落到前端表达层与只读呈现层。

不得再出现：

- 一份文档说 `phase12-08` 只做表达层，另一份文档又要求它重构 query owner；
- 一份文档要求共享语义来源单值化，另一份文档又允许页面继续硬编码另一套实体解释；
- 一份文档要求四个 detail page 是 primary owner，另一份文档又把它们降回只改 summary card 的被动容器。

#### Scenario: 读者对齐 phase12-08 的实现边界
- **WHEN** 读者同时查看 `phase12-04 ~ 07` 与 `phase12-08`
- **THEN** 能得到同一套“表达层收口、非结构重构”的职责定义
- **AND** 后续实现不需要再发明第二套边界解释

## REMOVED Requirements

### Requirement: 允许页面和组件继续各自硬编码四实体解释
**Reason**: 这会让 Product / Repository / Module / Decision 在不同页面继续被解释成互相冲突的登记对象、普通模块或孤立文本，违背 `phase12` 语义收口目标。
**Migration**: 必须将稳定复用的四实体语义短语收敛到 `frontend/src/features/project-context/data/shared-semantic-constants.ts`，页面与组件只保留本地渲染结构，不再保留第二套冲突文案来源。
