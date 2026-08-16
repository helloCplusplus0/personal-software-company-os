# phase12-10 落实关键页面与共享读路径的联动回流 Spec

## Why
`phase12-08` 已收口四实体语义表达，`phase12-09` 已落地共享只读入口，但关键页面之间的共享摘要接入、返回链、reread 与旧文案回收还没有形成同一套可执行闭环。`phase12-10` 需要把这些衍生消费页与 detail 页真正接到同一条共享读路径上，避免页面各自继续放大语义漂移。

## What Changes
- 冻结 `phase12-10` 为“关键页面与共享读路径联动回流”的实现型子任务，而不是新增结构化主锚点
- 落实 `dashboard / onboarding / reviews/daily / reviews/weekly / detail pages` 对共享语义摘要、固定入口或受控派生摘要的接入矩阵
- 落实共享读路径更新后的 reread、精确失效与返回链协同，避免页面成功回流后继续显示旧解释
- 回收关键页面中仍残留的四实体重复解释、漂移文案或不稳定入口
- **BREAKING** 禁止 `dashboard / onboarding / reviews/daily / reviews/weekly` 再各自生长新的结构化 project-context 主锚点

## Impact
- Affected specs:
  - `phase12_semantic_alignment_and_readonly_consumption_foundation`
  - `phase12_05_design_readonly_consumption_shared_entry`
  - `phase12_06_design_frontend_read_path_owner_shared_summary_reread`
  - `phase12_08_land_frontend_four_entity_semantic_alignment`
  - `phase12_09_land_readonly_consumption_shared_entry`
- Affected code:
  - `frontend/src/features/project-context/*`
  - `frontend/src/features/dashboard/*`
  - `frontend/src/features/onboarding/*`
  - `frontend/src/features/review/*`
  - `frontend/src/features/product-registry/*`
  - `frontend/src/features/repository-binding/*`
  - `frontend/src/features/module-registry/*`
  - `frontend/src/features/decision-center/*`

## ADDED Requirements

### Requirement: 冻结 phase12-10 的联动回流职责与非目标
系统 SHALL 将 `phase12-10` 冻结为“关键页面与共享读路径联动回流”的实现型子任务，只承接 `phase12-05 / 06 / 08 / 09` 已冻结结论，不新增第二套锚点、第二套解释链或第二套写路径。

当前子任务至少必须完成：

- 关键页面通过同一套共享语义摘要、固定入口或受控派生摘要回看当前四实体角色；
- 页面返回链、成功回流与 reread 不再放大旧解释或漂移文案；
- `dashboard / onboarding / reviews/daily / reviews/weekly` 继续保持“衍生消费页”身份，不升格为新的结构化入口；
- 仍残留在页面层的重复解释与入口不稳定实现被识别并回收。

补充冻结：

- `phase12-10` 不新增业务写路径、第二 query owner、第二 adapter 主线或第二 project-context 服务；
- `phase12-10` 不重开 `repository_id` resolver、L1/L3 分层或后端合同讨论；
- `phase12-10` 若需要新增或调整前端承接位，只能沿 `frontend/src/features/project-context/` 与既有页面主线继续实现。

#### Scenario: 执行者开始承接 phase12-10
- **WHEN** 执行者开始实现 `phase12-10`
- **THEN** 能明确这是联动回流与衍生消费页接入任务
- **AND** 不会把范围偷渡成新的结构重构或新协议扩张

### Requirement: 关键页面必须通过共享语义摘要或固定入口解释四实体角色
系统 SHALL 让 `dashboard / onboarding / reviews/daily / reviews/weekly` 与相关 detail pages 使用同一套共享语义来源、固定入口或受控派生摘要来解释 `Product / Repository / Module / Decision` 的当前角色。

当前阶段至少必须落实：

- detail pages 在具备既有锚点条件时优先消费共享只读摘要与固定入口；
- `dashboard / onboarding / reviews/daily / reviews/weekly` 只能消费共享语义常量、固定入口或受控派生摘要；
- 关键页面不得继续依赖切片内旧文案单独解释四实体角色；
- 同一实体在关键页面上的解释句式不得出现新的并列版本。

#### Scenario: 用户从关键页面回看当前项目
- **WHEN** 用户在 Dashboard、Review、Onboarding 或 Detail Page 上回看当前项目
- **THEN** 能看到同一套四实体解释来源
- **AND** 不会因为页面切换而读到另一套实体定义

### Requirement: 共享读路径更新后的 reread、失效刷新与返回链必须协同
系统 SHALL 让共享读路径的精确失效、页面成功回流与返回链保持同一口径，避免用户在完成写动作、跳转返回或重新进入页面后看到过期的共享解释。

当前阶段至少必须落实：

- 能拿到唯一 `repositoryId` 的成功回调继续承担 `['project-context', repositoryId]` 精确失效；
- 关键页面返回链回到上游页面后，消费到的共享摘要与固定入口不得回退到旧解释；
- 衍生消费页若只依赖静态共享语义来源，则不得伪造 project-context reread；
- 页面局部错误态、重试与 reread 仍遵守 `phase12-06 / 09` 已冻结的局部降级口径。

#### Scenario: 用户完成一次影响共享摘要的写操作后返回关键页面
- **WHEN** 用户在 detail 链路中完成绑定、映射或相关写操作并返回关键页面
- **THEN** 页面展示的共享摘要与固定入口已与最新状态对齐
- **AND** 不会继续保留旧文案或旧入口解释

### Requirement: 关键页面中的重复解释、漂移文案与不稳定入口必须被回收
系统 SHALL 回收关键页面中仍残留的四实体重复解释、漂移文案与入口不稳定实现，并把正式解释来源收敛到共享语义摘要、固定入口或受控派生摘要。

当前阶段至少必须落实：

- 盘点 `dashboard / onboarding / reviews/daily / reviews/weekly / detail pages` 中仍在单独解释四实体角色的 surface；
- 将需要保留的解释迁回共享语义来源或受控入口；
- 将不能稳定定位、不能稳定 reread 或与四实体冻结语义冲突的旧 surface 标记为必须替换或下线；
- 不允许为了“先过页面验收”保留第二套长期解释链。

#### Scenario: 复核关键页面的解释 surface
- **WHEN** 复核者逐页检查关键页面的实体解释、共享摘要与入口展示
- **THEN** 能机械识别每个 surface 是否来自正式共享承接链
- **AND** 能确认旧文案与不稳定入口已被回收或替换

## MODIFIED Requirements

### Requirement: phase12 的共享只读落地顺序在 phase12-10 上继续收口
`phase12` 内部顺序 SHALL 在 `phase12-10` 上继续保持同一口径：

- `phase12-08` 收口前端四实体语义表达；
- `phase12-09` 落实共享只读入口与基础消费接入；
- `phase12-10` 负责把关键页面、共享摘要、返回链与 reread 行为收成可验证闭环。

不得再出现：

- 一份文档要求关键页面复用共享语义来源，另一份文档又允许保留页面私有解释；
- 一份文档要求衍生消费页只做消费，另一份文档又允许它们自长结构化主锚点；
- 一份文档要求精确失效与返回链协同，另一份文档又允许页面回流后继续显示旧摘要。

#### Scenario: 读者连续查看 phase12-08、09、10
- **WHEN** 读者连续查看 `phase12-08 / 09 / 10`
- **THEN** 能得到从表达收口、共享入口到关键页面回流闭环的单一叙事
- **AND** 后续执行不需要再发明第二套阶段职责说明

## REMOVED Requirements

### Requirement: 允许关键页面长期保留各自的实体解释与入口裁剪
**Reason**: 这会让 `phase12` 的共享只读与语义一致性目标失效，继续放大页面之间的语义漂移与回流不一致。
**Migration**: 后续实现必须将这些解释迁回 `frontend/src/features/project-context/` 或既有固定入口主线，并通过统一的 reread / 返回链规则消费。
