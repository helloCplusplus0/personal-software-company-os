# phase10-03 冻结 Decision 生命周期、detail CTA 与 pending 语义的统一矩阵 Spec

## Why

`phase10-01` 已冻结 `Asset-Action Closure` 的范围边界，`fix_002 / fix_003` 也已经证明 `Decision` 的 pending 语义与状态推进必须走同一条 canonical 主线。但如果不在 `phase10-03` 继续把四态生命周期、`Decision Detail` CTA、`Dashboard / Daily Review / Current Focus` 的 pending 解释与 reread 规则冻结成单值矩阵，后续 `/spec -> 实现 -> 验收` 仍然会在详情页、review 信号与 dashboard 信号之间长出三套解释口径。

## What Changes

- 冻结 `Decision` 在 `phase10` 中继续只使用既有四态：`proposed / active / superseded / archived`
- 冻结 `Decision` 最小生命周期矩阵与终态解释
- 冻结 `Decision Detail` 在各状态下允许展示的 CTA 矩阵
- 冻结 `Dashboard / Daily Review / Current Focus` 与 `Decision.status` 的统一 pending 语义
- 冻结“退出 pending”的正式条件与 reread 规则
- 冻结 `Decision Detail` 作为唯一正式状态推进承接位
- 冻结当前阶段不引入第五态、不引入第二事实源、不让 `decision_links / review_records` 代理 pending 退出

## Impact

- Affected specs:
  - `phase10_01_freeze_asset_action_closure_scope_success_non_goals`
  - `fix_002_003_close_decision_pending_status_loop`
  - `phase03_02_decision_template_status_read_model`
  - `phase03_10_decision_center_formal_spec`
  - `phase08_03_freeze_feedback_decision_update_action_owner`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
  - 后续 `phase10-05 / 06 / 07 / 09 / 10 / 11` 的 `/spec` 与实现规格
- Affected code:
  - 当前无直接源码改动
  - 后续会影响 `frontend/src/features/decision-center/`
  - 后续会影响 `frontend/src/features/dashboard/`
  - 后续会影响 `frontend/src/features/review/`
  - 后续会影响 `backend/internal/decisioncenter/`
  - 后续会影响 `backend/internal/dashboard/`
  - 后续会影响 `backend/internal/review/`

## ADDED Requirements

### Requirement: `Decision` 在 `phase10` 中必须继续保持四态单值生命周期

系统 SHALL 将 `Decision` 在 `phase10` 中继续冻结为既有四态：`proposed / active / superseded / archived`，并要求后续 `/spec`、实现与验收不得新增第五态或局部替代状态。

#### Scenario: 判断 `Decision.status` 范围

- **WHEN** 后续 `/spec`、实现或验收描述 `Decision.status`
- **THEN** 当前阶段只允许使用 `proposed / active / superseded / archived`
- **AND** 不得新增 `acknowledged / resolved / done / dismissed` 等第五态
- **AND** 不得在页面层、读模型层或临时筛选层引入与这四态并列的局部状态真相源

### Requirement: `Decision` 最小生命周期矩阵必须单值冻结

系统 SHALL 将 `Decision` 的最小生命周期矩阵冻结为当前阶段唯一允许的状态推进矩阵。

#### Scenario: `proposed` 的允许迁移

- **WHEN** `Decision.status = proposed`
- **THEN** 当前阶段只允许正式推进到：
  - `active`
  - `superseded`
  - `archived`

#### Scenario: `active` 的允许迁移

- **WHEN** `Decision.status = active`
- **THEN** 当前阶段只允许正式推进到：
  - `superseded`
  - `archived`

#### Scenario: 终态禁止继续推进

- **WHEN** `Decision.status = superseded` 或 `archived`
- **THEN** 当前阶段必须将其视为终态
- **AND** 不得继续展示新的状态推进 CTA
- **AND** 不得在页面局部制造“重新激活”或等价的旁路状态推进语义

### Requirement: `Decision Detail` 必须是唯一正式状态推进承接位

系统 SHALL 将 `Decision Detail` 冻结为当前阶段唯一承接 `Decision.status` 正式推进的 canonical 页面。

#### Scenario: 判断状态推进承接页归属

- **WHEN** 用户需要让某条 `Decision` 正式退出 `pending` 或推进到下一个 canonical 状态
- **THEN** 必须进入既有 `Decision Detail`
- **AND** 正式状态推进不得在 `Dashboard / Daily Review / Current Focus` 内联为第二套写路径
- **AND** `Review` 页面也不得成为新的并列状态推进承接位

### Requirement: `Decision Detail` CTA 矩阵必须与四态生命周期严格一致

系统 SHALL 冻结 `Decision Detail` 在各状态下允许展示的 CTA，要求页面 CTA 与四态生命周期矩阵严格对齐，并同时覆盖状态推进 CTA 与非状态型 CTA 的完整 inventory。

#### Scenario: `proposed` 状态下的 CTA

- **WHEN** 用户进入一条 `status = proposed` 的 `Decision Detail`
- **THEN** 页面必须允许展示以下正式状态推进 CTA：
  - `Mark Active`
  - `Mark Superseded`
  - `Archive`
- **AND** 不得展示超出这三类推进语义的旁路状态按钮
- **AND** 页面仍允许展示与当前决策相关的非状态型 CTA，但只限于：
  - 进入或查看既有 canonical 关联目标
  - 进入既有 canonical handoff
  - 返回 `Dashboard / Daily Review / Current Focus` reread
- **AND** 不得把这些非状态型 CTA 解释为退出 pending 的代理动作

#### Scenario: `active` 状态下的 CTA

- **WHEN** 用户进入一条 `status = active` 的 `Decision Detail`
- **THEN** 页面必须允许展示以下正式状态推进 CTA：
  - `Mark Superseded`
  - `Archive`
- **AND** 不得继续展示 `Mark Active`
- **AND** 页面允许继续展示结果消费或 canonical 导航类 CTA，但这些 CTA 只能承接：
  - 查看当前已生效决策的关联目标
  - 进入既有 canonical owner 消费当前生效结果
  - 返回 `Dashboard / Daily Review / Current Focus` reread
- **AND** 不得把 `active` 重新解释为 pending 决策的动作入口

#### Scenario: 终态下的 CTA

- **WHEN** 用户进入一条 `status = superseded` 或 `archived` 的 `Decision Detail`
- **THEN** 页面不得继续展示状态推进 CTA
- **AND** 页面只允许保留以下非状态型 CTA：
  - 结果消费
  - 查看既有 canonical 关联目标
  - 详情 reread
  - 返回来源入口
- **AND** 不得继续展示任何让用户误以为“仍需处理当前决策状态”的 CTA

#### Scenario: `Decision Detail` 非状态型 CTA inventory 的统一边界

- **WHEN** 后续 `/spec`、实现或验收定义 `Decision Detail` 中除状态推进外的 CTA
- **THEN** 当前阶段只允许存在以下三类非状态型 CTA：
  - canonical 关联目标消费或导航 CTA
  - canonical handoff CTA
  - 返回 `Dashboard / Daily Review / Current Focus` 的 reread CTA
- **AND** 不得在 `Decision Detail` 长出第二套 review-local action hub
- **AND** 不得让非状态型 CTA 反向承担状态推进语义

### Requirement: `pending decision` 的判定必须完全锚定 canonical `Decision.status`

系统 SHALL 冻结 `Dashboard / Daily Review / Current Focus` 的 `pending decision` 判定完全锚定 canonical `Decision.status`，不得由 `decision_links`、`review_records`、局部已读状态或页面隐藏逻辑代理。

#### Scenario: `proposed` 继续视为 pending

- **WHEN** `Dashboard / Daily Review / Current Focus` 读取待处理决策
- **THEN** `Decision.status = proposed` 必须继续被解释为 pending
- **AND** 不得因为该 `Decision` 已存在 `decision_links` 就自动退出 pending
- **AND** 不得因为存在 `review_records` 或用户曾看过详情页就自动退出 pending

#### Scenario: 非 `proposed` 不再视为 pending

- **WHEN** 某条 `Decision.status` 已从 `proposed` 正式推进为 `active / superseded / archived`
- **THEN** `Dashboard / Daily Review / Current Focus` 必须停止将其显示为 pending
- **AND** 不得继续误报为“仍待处理”

### Requirement: 退出 pending 的正式条件必须单值冻结

系统 SHALL 冻结“退出 pending”的正式条件：只有 canonical `Decision.status` 正式离开 `proposed`，该决策才算退出 pending。

#### Scenario: 正式退出 pending

- **WHEN** 某条 `Decision` 在 `Decision Detail` 中被正式推进到 `active / superseded / archived`
- **THEN** 该决策必须被判定为已退出 pending
- **AND** 后续 reread 必须与这一结果保持一致

#### Scenario: 不足以退出 pending 的动作

- **WHEN** 用户只完成以下任一动作：
  - 建立 `decision_links`
  - 提交 `review_records`
  - 浏览 `Decision Detail`
  - 页面局部隐藏 pending 卡片
- **THEN** 这些动作都不得单独被解释为退出 pending 的正式条件

### Requirement: `Dashboard / Daily Review / Current Focus` 必须复用同一套 canonical 状态解释

系统 SHALL 冻结 `Dashboard / Daily Review / Current Focus` 对 `Decision.status` 的解释必须完全复用同一套 canonical 口径，不得各自生长不同的 pending 过滤规则，并必须为四态提供完整消费语义。

#### Scenario: 跨消费面的一致性

- **WHEN** 同一条 `Decision` 同时被 `Dashboard / Daily Review / Current Focus` 消费
- **THEN** 三者必须基于同一套 canonical `Decision.status` 解释它是否仍为 pending
- **AND** 不得出现一个页面认为已处理、另一个页面仍认为待处理的分裂语义

#### Scenario: `proposed` 在各消费面上的统一语义

- **WHEN** `Decision.status = proposed`
- **THEN** `Dashboard / Daily Review / Current Focus` 必须统一将其解释为仍待正式推进的决策
- **AND** 若页面展示该决策的下一步动作入口，正式入口必须继续指向 `Decision Detail`

#### Scenario: `active` 在各消费面上的统一语义

- **WHEN** `Decision.status = active`
- **THEN** `Dashboard / Daily Review / Current Focus` 必须统一将其解释为已生效、但不再属于 pending 的当前结论
- **AND** 若页面继续展示该决策，只允许作为当前生效结果消费或进入 `Decision Detail` 的导航入口
- **AND** 不得继续将其归入 pending bucket、pending count 或等价的“仍待处理”语义

#### Scenario: `superseded` 在各消费面上的统一语义

- **WHEN** `Decision.status = superseded`
- **THEN** `Dashboard / Daily Review / Current Focus` 必须统一将其解释为已被后续结论替代的历史结果
- **AND** 若页面继续展示该决策，只允许作为历史结果消费或进入 `Decision Detail` 的导航入口
- **AND** 不得继续将其解释为当前待推进动作

#### Scenario: `archived` 在各消费面上的统一语义

- **WHEN** `Decision.status = archived`
- **THEN** `Dashboard / Daily Review / Current Focus` 必须统一将其解释为已归档保留的历史结果
- **AND** 若页面继续展示该决策，只允许作为归档结果消费或进入 `Decision Detail` 的导航入口
- **AND** 不得继续将其解释为当前待推进动作

### Requirement: 状态推进后的 reread 规则必须单值冻结

系统 SHALL 冻结状态推进成功后的 reread 规则，要求 `Decision Detail / Dashboard / Daily Review / Current Focus` 在 reread 中共同回答同一组问题，而不是只靠 toast 或页面局部隐藏维持一致性假象。

#### Scenario: reread 必须回答的核心问题

- **WHEN** 某条 `Decision` 完成正式状态推进后
- **THEN** `Decision Detail / Dashboard / Daily Review / Current Focus` 的 reread 至少必须共同回答：
  1. 这条 `Decision` 当前是否还需要推进？
  2. 如果不需要，系统是否已经停止误报？
  3. 如果仍需要，系统是否给出了新的正式动作入口？

#### Scenario: reread 禁止项

- **WHEN** 后续实现或验收描述状态推进后的 reread 行为
- **THEN** 不得通过页面局部隐藏、前端临时过滤或 toast 假象掩盖未完成状态
- **AND** 必须基于 canonical facts 重读结果

### Requirement: 当前阶段必须保持 `Decision` pending 主线与 detail CTA 主线合一

系统 SHALL 冻结一个统一原则：`Decision Detail` 的 CTA 主线与 `Dashboard / Daily Review / Current Focus` 的 pending 主线必须是同一条 canonical 业务解释，不允许分别生长。

#### Scenario: CTA 与 pending 必须指向同一业务解释

- **WHEN** 后续 `/spec`、实现或验收同时定义 `Decision Detail` 的状态推进 CTA 与 `Dashboard / Daily Review / Current Focus` 的 pending 语义
- **THEN** 两者必须共同回到同一套 `Decision.status` 生命周期矩阵
- **AND** 不得出现“detail 可以推进，但 dashboard 仍按别的语义催促”或“dashboard 不再催促，但 detail 仍无正式出口”的断裂

## MODIFIED Requirements

### Requirement: `Decision.status` 在 `phase10` 中的解释

`phase10-03` 修改了 `Decision.status` 在当前阶段的消费方式：它不再只承接“详情展示字段”或“中心列表状态标签”，而必须同时承接生命周期推进、pending 判定与 reread 解释。

#### Scenario: 判断 `Decision.status` 的新消费面

- **WHEN** 后续 `/spec`、实现或验收消费 `Decision.status`
- **THEN** 必须同时覆盖：
  - `Decision Detail` 状态推进 CTA
  - `Dashboard / Daily Review / Current Focus` pending 判定
  - 状态推进后的 reread 结果
- **AND** 不得把 `Decision.status` 再拆成多套局部解释

### Requirement: `fix_002 / fix_003` 收口结果在 `phase10` 中的解释

`phase10-03` 修改了对 `fix_002 / fix_003` 收口结果的承接方式：后续不再只要求“已有状态推进写链”和“pending 误报已修复”，而必须继续把两者统一提升为 `phase10` 的单值生命周期与 pending 语义矩阵。

#### Scenario: 判断 `fix_002 / fix_003` 是否被正确继承

- **WHEN** 后续 `/spec`、实现或验收承接 `Decision` 主线
- **THEN** 必须继续保持：
  - `Decision Detail` 是唯一正式状态推进承接位
  - pending 判定继续完全锚定 canonical `Decision.status`
  - `decision_links / review_records` 不是退出 pending 的代理条件
- **AND** 不得把 fix 阶段已收口的问题重新写散

## REMOVED Requirements

### Requirement: 使用 `decision_links`、`review_records` 或页面局部状态代理 pending 退出

**Reason**: 这种解释会再次制造第二事实源，直接破坏 `phase10` 所要求的单值 `Decision` 生命周期与 pending reread 语义，也会把 `fix_002 / fix_003` 已收口的问题重新带回当前阶段。
**Migration**: 后续 `/spec` 与实现必须统一回到 canonical `Decision.status`；只有 `Decision` 正式离开 `proposed`，才允许退出 pending。

### Requirement: 在 `Dashboard / Daily Review / Current Focus` 内联第二套状态推进 CTA

**Reason**: 当前阶段已经冻结 `Decision Detail` 为唯一正式状态推进承接页；如果在其他消费面继续内联第二套状态推进按钮，会让 detail CTA、review signal 与 dashboard signal 再次各讲一套语义。
**Migration**: 其他消费面只允许承接 canonical 导航、进入 `Decision Detail` 或 reread 结果消费；正式状态推进继续统一回到 `Decision Detail`。
