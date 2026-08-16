# phase12-04 产出前端四实体语义一致性设计 Spec

## Why
`phase12-01 ~ 03` 已分别冻结了本阶段边界、Web 端语义承接矩阵与只读消费深化边界，但如果不继续把前端四实体语义一致性设计正式落成可执行设计结果，后续实现仍会临场猜测“哪些页面必须改文案、哪些摘要应共享呈现、哪些对象允许 no-change”。

## What Changes
- 冻结 `Product / Repository / Module / Decision` 相关前端页面、路由、组件的语义一致性设计输出要求
- 冻结摘要卡片、空态、说明文案、下一步动作说明的承接矩阵与分类规则
- 冻结 route / page / component / surface 的 `must-change / follow-regression / no-change` 逐项留档要求
- 冻结“哪些语义收敛为共享呈现、哪些保留在切片内”的设计边界
- 冻结 `phase12-04` 设计结果的最小产物模板、before/after 样例与明确不做清单
- 对齐 `phase12-02`、`phase12-03` 与 `phase12` 三件套中的前端语义设计口径

## Impact
- Affected specs: `phase12_semantic_alignment_and_readonly_consumption_foundation`
- Affected code:
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`
  - `frontend/src/routes/dashboard.tsx`
  - `frontend/src/routes/onboarding.tsx`
  - `frontend/src/routes/reviews/daily.tsx`
  - `frontend/src/routes/reviews/weekly.tsx`
  - `frontend/src/routes/products/$productId.tsx`
  - `frontend/src/routes/repositories/$repositoryId.tsx`
  - `frontend/src/routes/modules/$moduleId.tsx`
  - `frontend/src/routes/decisions/$decisionId.tsx`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/product-registry/components/product-summary-card.tsx`
  - `frontend/src/features/repository-binding/components/repository-summary-card.tsx`
  - `frontend/src/features/module-registry/components/module-summary-card.tsx`
  - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
    - `frontend/src/features/dashboard/components/current-focus-section.tsx`
    - `frontend/src/features/dashboard/components/asset-feedback-section.tsx`
    - `frontend/src/features/dashboard/components/recent-activity-section.tsx`
    - `frontend/src/features/review/components/review-page-shell.tsx`
    - `frontend/src/features/module-registry/components/module-next-action-bar.tsx`
    - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
    - `frontend/src/features/dashboard/components/onboarding-cta-button.tsx`
    - `frontend/src/features/review/components/review-action-footer.tsx`
    - `frontend/src/features/product-registry/pages/product-list-page.tsx`
    - `frontend/src/features/repository-binding/pages/repository-binding-list-page.tsx`
    - `frontend/src/features/module-registry/pages/module-list-page.tsx`
    - `frontend/src/features/decision-center/pages/decision-list-page.tsx`

## ADDED Requirements

### Requirement: 冻结 `phase12-04` 的最小设计产物模板
系统规格 SHALL 要求 `phase12-04` 产出的前端四实体语义一致性设计至少包含可直接进入实现的最小设计产物，而不是停留在方向描述。

当前阶段至少必须显式产出：

1. **影响对象清单**
   - 逐项列出 route / page / component / surface
   - 每项必须标注 `must-change`、`follow-regression` 或 `no-change`
2. **结论矩阵**
   - 当前承接什么
   - 需要改成什么
   - 为什么要改
   - 若为 `no-change`，为什么仍满足当前阶段冻结口径
3. **承接矩阵**
   - 摘要卡片、空态、说明文案、下一步动作说明分别落在哪个页面、组件或共享呈现承接位
4. **共享语义来源 vs 切片内渲染矩阵**
   - 哪些高频语义短语或共享解释应收敛为唯一共享语义来源
   - 哪些页面布局、组件结构、空态插入位或导语插入位继续保留在切片内渲染
5. **Before / After 样例**
   - 至少提供一组页面级和一组 surface 级 before / after
6. **明确不做清单**
   - 本子任务没有扩入的读路径 owner、后端合同、写回通道、协议扩张或页面重组事项

#### Scenario: 执行者读取 phase12-04 设计结果
- **WHEN** 后续执行者准备进入实现
- **THEN** 能直接从设计结果获得影响对象、结论矩阵、承接矩阵与 before / after 样例
- **AND** 不需要再临场补“这里到底要怎么改”

### Requirement: 冻结前端影响对象全量清单
系统规格 SHALL 冻结 `phase12-04` 必须覆盖的前端 route / page / component 影响对象清单，允许 `no-change`，但不允许整体跳过。

当前阶段至少必须覆盖以下对象：

1. **Route**
   - `frontend/src/routes/dashboard.tsx`
   - `frontend/src/routes/onboarding.tsx`
   - `frontend/src/routes/reviews/daily.tsx`
   - `frontend/src/routes/reviews/weekly.tsx`
   - `frontend/src/routes/products/$productId.tsx`
   - `frontend/src/routes/repositories/$repositoryId.tsx`
   - `frontend/src/routes/modules/$moduleId.tsx`
   - `frontend/src/routes/decisions/$decisionId.tsx`
2. **Page**
   - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
   - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
   - `frontend/src/features/review/pages/daily-review-page.tsx`
   - `frontend/src/features/review/pages/weekly-review-page.tsx`
   - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
   - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
   - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
   - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
3. **Component**
   - `frontend/src/features/product-registry/components/product-summary-card.tsx`
   - `frontend/src/features/repository-binding/components/repository-summary-card.tsx`
   - `frontend/src/features/module-registry/components/module-summary-card.tsx`
   - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
    - `frontend/src/features/dashboard/components/current-focus-section.tsx`
    - `frontend/src/features/dashboard/components/asset-feedback-section.tsx`
    - `frontend/src/features/dashboard/components/recent-activity-section.tsx`
    - `frontend/src/features/review/components/review-page-shell.tsx`
    - `frontend/src/features/module-registry/components/module-next-action-bar.tsx`
    - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
    - `frontend/src/features/dashboard/components/onboarding-cta-button.tsx`
    - `frontend/src/features/review/components/review-action-footer.tsx`
  4. **Surface owner files**
    - `frontend/src/features/product-registry/pages/product-list-page.tsx`
    - `frontend/src/features/repository-binding/pages/repository-binding-list-page.tsx`
    - `frontend/src/features/module-registry/pages/module-list-page.tsx`
    - `frontend/src/features/decision-center/pages/decision-list-page.tsx`
    - 这些对象虽然不是四实体 primary owner 页面，但本阶段已冻结其空态、来源提示或动作说明属于正式 surface 审计面

补充冻结：

- Route 不作为语义事实源，但必须显式记录其承接职责与是否 `follow-regression`；
- Page / component 不得因为“当前看起来没问题”而被默认跳过，必须逐项给出分类结果；
- 已被 `phase12-02` 冻结为 primary owner 的四个 detail page，必须继续保持 primary owner 身份，不得在 `phase12-04` 中被降级为只做被动跟随的容器；
- `List pages`、toolbar 与搜索 store 可以继续作为默认非 primary owner，但若在设计中被引用，也必须记录其跟随关系或 `no-change` 理由。

#### Scenario: 执行者盘点前端影响面
- **WHEN** 执行者开始编写 `phase12-04` 设计
- **THEN** 必须逐项覆盖上述 route / page / component 清单
- **AND** 不允许只提交抽象总结而省略对象级结论

### Requirement: 冻结 surface 承接矩阵
系统规格 SHALL 要求 `phase12-04` 对摘要卡片、空态、说明文案与下一步动作说明产出逐项承接矩阵，而不是只回答页面级 owner。

当前阶段至少必须对以下 surface 类型逐项审计：

- 摘要卡片（Summary Cards）
- 空态（Empty States）
- 说明文案（Hint / Description Copy）
- 下一步动作说明（Next-Action Descriptions）

每类 surface 都必须回答：

- 当前承接什么语义
- 是否与 `Product / Repository / Module / Decision` 冻结语义一致
- 应修改为共享呈现还是保留在切片内
- 最终分类为 `must-change`、`follow-regression` 还是 `no-change`

补充冻结：

- 不允许用“页面已经在 owner 矩阵里”代替对 surface 的独立审计；
- `Module` 相关 surface 不得回落为“模块登记”语义；
- `Decision` 相关 surface 不得回落为“孤立文本卡片/提交决策”语义。

#### Scenario: 执行者审计页面内具体 surface
- **WHEN** 执行者为某个页面写设计结论
- **THEN** 必须同时给出该页所含摘要卡片、空态、说明文案与下一步动作说明的承接结论
- **AND** 不会只记录页面层结论而遗漏具体 surface

### Requirement: 冻结共享语义来源与切片内渲染的设计边界
系统规格 SHALL 要求 `phase12-04` 明确哪些语义应收敛为共享语义来源、哪些必须继续保留在切片内渲染，但不得越权改写 `phase12-03` 已冻结的共享只读 owner 边界。

当前阶段至少固定如下边界：

- `phase12-04` 可以设计“共享语义来源”的表达层收敛，但不得把它直接写成新的跨切片只读数据 owner；
- 若某段语义只是文案、提示或页面解释的稳定复用，可以标记为“共享语义来源候选”；
- 若某段语义依赖 repository-scoped 结构化摘要、规则入口或文档入口，则只允许引用 `phase12-03` 已冻结的共享只读 owner 前提，不得在本任务内重新定义；
- 切片内渲染的对象必须说明为什么继续留在切片内仍能满足当前阶段冻结口径。

#### Scenario: 执行者判断某段语义是否需要共享语义来源
- **WHEN** 执行者发现多个页面复用同一段解释性语言
- **THEN** 必须判断它是共享语义来源候选还是继续保留在切片内渲染
- **AND** 不会在 `phase12-04` 内越权新增第二套共享只读 owner

### Requirement: 冻结 `must-change / follow-regression / no-change` 记录规则
系统规格 SHALL 要求 `phase12-04` 对所有进入影响面的对象逐项记录 `must-change / follow-regression / no-change`，并显式写明理由。

当前阶段至少固定如下规则：

- `must-change`：当前对象存在语义漂移，必须修改文案、摘要、说明或动作表达；
- `follow-regression`：当前对象不独立定义四实体语义，但需要跟随 primary owner 或共享语义来源结果完成回归验证；
- `no-change`：当前对象已满足冻结口径，但必须记录理由与依据；
- 若任一对象未被分类，则 `phase12-04` 不得视为完成。

#### Scenario: 设计结果包含无需改动对象
- **WHEN** 执行者判断某对象当前无需改动
- **THEN** 必须显式写出其 `no-change` 理由
- **AND** 不允许用“本轮先不看”替代正式冻结结果

## MODIFIED Requirements

### Requirement: phase12 三件套中的前端语义设计表达
`phase12` 三件套与 `phase12-02 / 03` 上游 spec 在前端语义设计上的表达 SHALL 对齐为同一口径：

- `phase12-04` 负责产出可直接进入实现的前端语义一致性设计
- 影响对象必须覆盖 route / page / component / surface 四层
- 设计结果必须显式给出共享语义来源 vs 切片内渲染判断
- `phase12-04` 不得越权改写 `phase12-03` 的共享只读 owner、读路径 owner 或更重通道边界
- `no-change` 允许存在，但必须逐项留档

不得再出现：

- 一个文档要求逐项审计，另一个文档又允许只给页面级概述；
- 一个文档允许共享语义来源候选，另一个文档又把它直接写成新的数据 owner；
- 一个文档允许 `no-change` 留档，另一个文档又把未列出的对象默认为无需处理。

#### Scenario: 读者对齐 phase12-04 的职责
- **WHEN** 读者分别查看 `dev_plan`、`architecture_plan`、`shared_baseline` 与 `phase12-02 / 03` 上游 spec
- **THEN** 能获得同一套 `phase12-04` 设计职责、影响面与不做边界
- **AND** 后续实现不需要再次猜测本任务是否已经完成设计冻结

## REMOVED Requirements

### Requirement: 将前端语义一致性设计留到实现阶段边做边定
**Reason**: `phase12-04` 的目标就是先把前端四实体语义一致性设计写成可执行设计结果，避免后续实现继续靠临场判断决定页面、组件与 surface 如何承接冻结语义。
**Migration**: 后续实现必须直接承接 `phase12-04` 设计产物中的影响对象清单、承接矩阵、分类结果与 before / after 样例，不再自行补造第二套设计结论。
