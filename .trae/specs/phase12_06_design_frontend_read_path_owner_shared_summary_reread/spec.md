# phase12-06 产出前端读路径 owner、共享摘要与回流设计 Spec

## Why
`phase12-05` 已冻结共享 owner / 入口矩阵，`phase12-07` 已冻结哪些后端真实字段可直接复用、哪些只允许停留在 L3 单向映射、哪些当前不做。`phase12-06` 的职责是把这些结论收成前端可执行的读路径 owner、共享摘要承接位、缓存 / reread / 成功回流关系设计，避免四实体解释逻辑继续散落在各切片 data 层。

## What Changes
- 冻结 `phase12-06` 在 `phase12` 内的正式职责、范围与非目标
- 冻结 8 个现有前端 read owner 与潜在 `frontend/src/features/project-context/data/*` 的审计清单、分类与边界
- 冻结 `query` 层纯只读、L3 共享只读 owner、页面层渲染 / 回流之间的正式分工
- 冻结四实体共享摘要、入口定位视图与切片内 detail / dashboard / onboarding / review 数据层之间的承接关系
- 冻结页面读取、缓存键、成功回流、局部 reread 与跨页 reread 的关系设计
- 冻结需要回收的散装页面解释逻辑与不允许继续分散承接的场景

## Impact
- Affected specs:
  - `phase12_semantic_alignment_and_readonly_consumption_foundation`
  - `phase12_05_design_readonly_consumption_shared_entry`
  - `phase12_07_design_backend_contract_export_shared_read_views`
  - `phase12_04_design_frontend_four_entity_semantic_alignment`
- Affected code:
  - `frontend/src/features/product-registry/data/use-product-detail-read.ts`
  - `frontend/src/features/repository-binding/data/use-repository-detail-read.ts`
  - `frontend/src/features/module-registry/data/use-module-detail-read.ts`
  - `frontend/src/features/decision-center/data/use-decision-detail-read.ts`
  - `frontend/src/features/dashboard/data/use-dashboard-overview-read.ts`
  - `frontend/src/features/onboarding/data/use-onboarding-read.ts`
  - `frontend/src/features/review/data/use-daily-review-read.ts`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/project-context/data/*`
  - `frontend/src/features/project-context/*`
  - `frontend/src/features/*/pages/*`
  - `frontend/src/features/*/components/*`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结 phase12-06 的前端读路径设计职责
系统规格 SHALL 将 `phase12-06` 冻结为“产出前端读路径 owner、共享摘要与回流设计”的正式子任务，而不是重新打开 `phase12-05` 的共享 owner 冻结结论，也不是反向要求 `phase12-07` 新增第二服务或影子合同。

当前子任务至少必须回答：

- 8 个既有 `use-*-read.ts` 的正式 owner 身份分别是什么；
- 哪些共享摘要、语义来源与入口定位视图应收敛到 `frontend/src/features/project-context/data/*`；
- 哪些内容继续保留在各切片 data owner / 页面 / 组件内；
- 页面读取、缓存、成功回流与 reread 之间如何衔接；
- 需要回收哪些散装页面解释逻辑，避免各页继续各讲一套四实体语义。

补充冻结：

- `phase12-06` 只产出前端设计结论，不直接改写后端合同、proto 或 `phase12-05 / 07` 已冻结的候选去留；
- `phase12-06` 必须继续遵守“`query` 层纯只读、业务写路径唯一 application 入口、共享只读不得变成第二 canonical facts”的项目规则；
- `phase12-06` 的完成标准是“足以直接指导后续实现收口”，而不是只写一个目录建议。

#### Scenario: 执行者开始承接 phase12-06
- **WHEN** 执行者开始编写 `phase12-06` 设计结果
- **THEN** 能明确本任务只负责冻结前端读路径 owner、共享摘要与 reread / 回流设计
- **AND** 不会反向改写 `phase12-05 / 07` 的 owner、字段或服务结论

### Requirement: 冻结前端 read owner 审计清单与分类矩阵
系统规格 SHALL 为 `phase12-06` 冻结正式的前端 read owner 审计清单，并逐项标注 `must-change / follow-regression / no-change`，避免实现阶段再临场决定“哪个 hook 该保留、哪个要进入共享只读”。

当前阶段至少必须显式覆盖：

1. `frontend/src/features/product-registry/data/use-product-detail-read.ts`
2. `frontend/src/features/repository-binding/data/use-repository-detail-read.ts`
3. `frontend/src/features/module-registry/data/use-module-detail-read.ts`
4. `frontend/src/features/decision-center/data/use-decision-detail-read.ts`
5. `frontend/src/features/dashboard/data/use-dashboard-overview-read.ts`
6. `frontend/src/features/onboarding/data/use-onboarding-read.ts`
7. `frontend/src/features/review/data/use-daily-review-read.ts`
8. `frontend/src/features/review/data/use-weekly-review-read.ts`
9. 若存在跨切片共享只读，新增 `frontend/src/features/project-context/data/*`

补充冻结：

- detail 页 read owner 默认继续保留切片内查询承接身份，但必须说明是否消费 L3 共享摘要；
- `dashboard / onboarding / review` 的 read owner 必须显式说明哪些继续是页面级 read model，哪些高频语义来源应回收到跨切片共享只读；
- 若当前 `frontend/src/features/project-context/` 目录尚不存在，`phase12-06` 必须显式回答“本阶段是否需要新增、若新增只承接什么”；
- 不允许把 audit 清单外的临时 helper、页面内联 `useQuery` 或组件私有 fetch 当成正式读路径 owner。

#### Scenario: 执行者盘点前端 read owner
- **WHEN** 执行者审阅前端只读路径
- **THEN** 能拿到一份完整的 read owner 清单与分类矩阵
- **AND** 不需要再猜某个页面数据逻辑是否应该迁入共享只读

### Requirement: 冻结 query 层、L3 共享只读 owner 与页面层的分工
系统规格 SHALL 冻结 `query` 层、`frontend/src/features/project-context/data/*` 与页面 / 组件渲染层之间的正式分工，使共享摘要不会继续散落在 `dashboard / onboarding / review / detail page` 各自的数据层。

当前阶段至少必须明确：

- 切片内 `use-*-read.ts` 继续承担原始读取、缓存键、响应解包与页面级 read model；
- `frontend/src/features/project-context/data/*` 若新增，只允许承接基于 `repository_id` 的共享 query options、共享摘要 adapter、入口定位视图与四实体共享语义来源；
- 页面与组件只消费切片 read owner 或 L3 共享只读结果，不得再拼装一套新的共享摘要合同；
- `phase12-07` 已判定“仅停留在 L3 / renderer 映射”的字段，必须在前端继续作为单向 view model / adapter 处理，不得倒逼后端补伪统一字段。

补充冻结：

- `frontend/src/features/project-context/data/*` 不得承接写路径、页面私有状态、局部空态拼接细节或第二套 query 事实源；
- `dashboard / onboarding / review` 若只需要共享语义 label、入口定位或裁剪后的只读摘要，不得继续在各自 data hook 中复制转换逻辑；
- detail 页 read owner 若已具备页面专属字段（如 Product 绑定仓库、Module 仓库映射、Decision linked modules），这些字段继续留在切片内，不得误提升为 L3 共享合同。

#### Scenario: 执行者决定某段只读逻辑落点
- **WHEN** 执行者判断某段四实体解释、入口定位或共享摘要逻辑应落在哪里
- **THEN** 能机械判断它属于切片 read owner、L3 共享只读 data owner，还是页面 / 组件渲染层
- **AND** 不会让多个切片各自复制同一套解释逻辑

### Requirement: 冻结页面读取、缓存、成功回流与 reread 关系
系统规格 SHALL 为 `phase12-06` 冻结页面读取、缓存、成功回流与 reread 的关系设计，使后续实现知道哪些查询应被失效、哪些共享摘要应被重读、哪些页面只需局部回流。

当前阶段至少必须明确：

- detail 页接入共享摘要后，页面私有 detail read 与基于 `repository_id` 的共享只读 read owner 之间的先后关系；
- `dashboard / onboarding / review` 消费共享只读时，哪些只需读取共享语义来源，哪些需要在成功写回后触发 reread；
- 写路径成功后，哪些失效动作仍归切片内 mutation owner 决定，哪些共享只读 query 需要被显式重读；
- 初次加载、局部重试、整页重试与成功回流后的 reread 范围如何区分。

补充冻结：

- `phase12-06` 只负责定义读路径与 reread 关系，不接管 mutation owner；
- 不允许让页面直接根据“写操作看起来可能影响语义”就各自随意失效一批 query；
- 若某页面并不稳定持有 `repository_id`，则不得把共享只读 reread 写成默认强依赖流程。

#### Scenario: 某写路径成功后页面需要刷新共享摘要
- **WHEN** 某个切片写路径成功，并且其结果会影响共享只读摘要或四实体解释
- **THEN** 执行者能明确由哪个 mutation owner 触发哪些 read owner 的失效或 reread
- **AND** 不会让 Dashboard、Review、Onboarding、Detail pages 各自复制一套刷新策略

### Requirement: 冻结散装页面解释逻辑的回收清单
系统规格 SHALL 要求 `phase12-06` 显式识别并冻结需要回收的散装页面解释逻辑，避免四实体语义继续以页面私有文案或 data transform 的形式扩散。

当前阶段至少必须审计：

- detail 页将实体字段临时拼成“共享摘要”的逻辑；
- `dashboard / onboarding / review` 为了展示四实体解释、入口链接或 next-action 提示而各自做的解释性 transform；
- 对 `entry_ref / entry_kind / label / summary / status_summary` 的重复裁剪与命名转换；
- 对 `phase12-07` 已判定为 L3 单向映射字段的重复生成逻辑。

补充冻结：

- 若一段解释逻辑仅服务单页布局，不必强行提升到共享只读；
- 只有在多个页面稳定复用、且不改变 canonical facts 时，才允许进入 `frontend/src/features/project-context/data/*`；
- 回收清单必须与 `phase12-04` 的 surface / 语义承接矩阵一致，不得重新改写 primary owner 页面职责。

#### Scenario: 执行者处理重复解释逻辑
- **WHEN** 执行者发现多个页面都在做相同的四实体解释或入口定位转换
- **THEN** 能判断它是否应进入共享只读 data owner
- **AND** 若不进入，也能说明它继续留在切片内渲染的原因

### Requirement: 冻结 phase12-06 的最小设计产物模板
系统规格 SHALL 要求 `phase12-06` 的设计结果继续满足 `phase12-04 ~ 07` 的统一最小模板，并补充适用于前端读路径设计的审计面。

当前子任务至少必须显式产出：

1. 影响对象清单
2. 结论矩阵
3. 承接位矩阵
4. 共享语义来源 vs 切片内渲染矩阵
5. Before / After 样例
6. 明确不做清单

同时至少补充以下前端读路径审计面：

- 每个 read owner 的缓存键 / 输入锚点 / 输出 shape / 是否消费共享只读；
- detail 页、dashboard、onboarding、daily review、weekly review 的 reread 关系；
- 需要新增的 `frontend/src/features/project-context/data/*` 文件边界；
- 明确不迁入共享只读的页面专属字段与转换。

#### Scenario: 执行者准备把 phase12-06 交给后续 /spec 或实现
- **WHEN** `phase12-06` 设计结果完成
- **THEN** 后续执行者能直接拿到前端 read owner 分类、共享摘要落点与 reread 关系
- **AND** 不需要再发明第二套 query contract 或刷新策略

## MODIFIED Requirements

### Requirement: phase12 三件套中的 phase12-06 顺序与职责表达
`phase12` 三件套中的 `phase12-06` 表达 SHALL 在本 spec 中继续对齐为同一口径：

- `phase12-05` 先冻结共享 owner、入口矩阵与 resolver 输入规则；
- `phase12-07` 再冻结后端真实字段、L3 单向映射边界与“不做”候选；
- `phase12-06` 最后消费 `phase12-05 / 07` 的结论，设计前端 read owner、共享摘要与 reread / 回流关系。

不得再出现：

- 一个文档要求 `phase12-06` 只消费既有结果，另一个文档又让它反向改写 `phase12-05 / 07`；
- 一个文档要求 `query` 层纯只读，另一个文档又允许页面 data owner 私自拼出跨切片共享摘要；
- 一个文档要求 `frontend/src/features/project-context/` 是唯一跨切片共享只读 owner，另一个文档又允许 `dashboard / onboarding / review` 各自保留一套共享解释数据层。

#### Scenario: 读者对齐 phase12-06 的职责
- **WHEN** 读者分别查看 `dev_plan`、`architecture_plan`、`shared_baseline`、`phase12-05` 与 `phase12-07`
- **THEN** 能获得同一套 `phase12-06` 职责、顺序、读路径 owner 边界与 reread 关系
- **AND** 后续 `/spec` 与实现不需要补造第二套前端只读解释

## REMOVED Requirements

### Requirement: 允许页面 data owner 各自保留一套跨切片共享解释逻辑
**Reason**: 这会让四实体语义、入口定位与共享摘要继续散落在 `dashboard / onboarding / review / detail page` 各自的数据层，破坏 `phase12` 的共享只读收敛目标。
**Migration**: 后续设计与实现必须将稳定复用的共享语义来源、入口定位视图与共享摘要 adapter 收敛到 `frontend/src/features/project-context/data/*` 或明确留在切片 read owner / 渲染层，不再允许页面 data owner 私自维护一套跨切片共享解释逻辑。
