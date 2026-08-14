# phase10_asset_action_closure_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase10_asset_action_closure_foundation` 的架构规划文档。

`phase08` 已完成 `Operating Review Loop` 的正式交付，`phase09` 已完成 `Template Reuse + Derived Intelligence Deepening` 的正式交付，`mvp0.3` `Real-Project Dry-Run` 第一轮阻断项也已通过 `fix_001 ~ fix_003` 的修复、独立复核与聚焦 rerun 正式收口。

根据 [PSCO-mvp04-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp04-summarize-feedback.md#L380-L387) 已冻结的仲裁结论，下一步不能直接跳到 `Agent Consumption Layer`，也不能重新扩长期对象宽度；而应先把 PSCO 从“实体已经登记、review 已经存在、提示已经给出”推进到“资产与动作真正闭环”。

因此，`phase10` 的职责不是增加更多新概念，而是交付 **`Asset-Action Closure`** 这一单一主交付能力：让 `Onboarding / Decision / Dashboard / Review / Detail pages` 共同承接真实下一步动作，而不是继续要求用户在多个并列实体页之间手动补链、手动猜测“系统到底要我下一步做什么”。

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `docs/README.md`
7. `PSCO-mvp04-summarize-feedback.md`
8. `docs/fix/fix_001_003_mvp03_dry_run_rerun_closure.md`
9. `docs/fix/fix_001_onboarding_cold_start_state_analysis.md`
10. `docs/fix/fix_002_decision_pending_signal_semantics_analysis.md`
11. `docs/fix/fix_003_decision_detail_status_advance_analysis.md`
12. `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
13. `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
14. `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
15. `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
16. `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`
17. `.trae/specs/phase09_11_validate_template_reuse_derived_hint_integration_browser_regression_acceptance/acceptance_report.md`
18. `docs/phase/phase08_operating_review_loop_foundation_*`
19. `docs/phase/phase09_template_reuse_derived_intelligence_foundation_*`

补充说明：

- `PSCO-mvp04-summarize-feedback.md` 冻结的是 `mvp0.4` 的方向排序、边界与推进顺序，不直接替代本阶段 `/plan`
- 三份 `fix analysis` 与 `fix_001_003` 收口记录共同证明：当前阶段不再是“先修 P0 阻断”，而是正式进入“闭动作链、降使用摩擦”的业务 phase
- `phase06 / phase08 / phase09` 提供的是本阶段必须直接复用的既有 `Onboarding / Review / Reuse / Detail` 能力，不允许推翻重写为第二套结构

## 3. 本阶段目标

`phase10` 的目标是：

> 在不新增长期核心实体、不把系统演化为任务管理器或 AI 工作台的前提下，让 PSCO 已有资产主线、review 主线与详情页主线正式承接“下一步动作”，使用户从冷启动到日常 reread 都能理解系统要求推进什么、如何推进、推进后如何回流。

本阶段需要回答的核心问题：

1. `Onboarding` 如何从“六段式登记流程”升级为“首轮建链引导”，减少后续跨详情页手动补链摩擦
2. `Decision` 如何在既有四态内形成最小但真实的生命周期，而不是只停留在“留痕”与“待处理”之间摇摆
3. `Dashboard / Daily Review / Decision Detail / Product Detail / Module Detail / Repository Detail` 如何形成统一的“下一步动作”承接矩阵
4. `Current Focus` 与 pending signals 如何继续锚定真实 canonical facts，而不是继续提示“存在过记录”
5. 哪些动作应该直接在 canonical owner 中完成，哪些动作应以清晰 handoff 的方式回到既有实体主线
6. `Onboarding`、`Review`、`Detail pages` 与 `Dashboard` 的读写 owner、route caller 与返回链应如何收敛，避免散落在多个页面中各自生长一套业务编排
7. 本阶段完成时，如何证明 PSCO 已经从“登记系统 + review 系统”推进到“具备真实动作闭环的经营系统”

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase10` 必须直接承接：

- `PSCO-mvp04-summarize-feedback.md`
- `docs/fix/fix_001_003_mvp03_dry_run_rerun_closure.md`
- `docs/fix/fix_001_onboarding_cold_start_state_analysis.md`
- `docs/fix/fix_002_decision_pending_signal_semantics_analysis.md`
- `docs/fix/fix_003_decision_detail_status_advance_analysis.md`
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
- `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
- `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
- `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`
- `.trae/specs/phase09_11_validate_template_reuse_derived_hint_integration_browser_regression_acceptance/acceptance_report.md`

不允许在本阶段重新解释：

- `.proto` 作为唯一长期合同源
- `ConnectRPC` 作为业务接口正式传输主线
- `Product Create` 作为正式 canonical 创建承接位
- `Decision` 在 MVP 中的中心地位
- `phase08 / phase09` 已完成的正式交付与验收结论
- `mvp0.4` 的方向排序：当前先做 `Asset-Action Closure`，不是 `Agent Consumption Layer`

### 4.2 当前阶段单一主交付能力

`phase10` 的单一主交付能力冻结为：

> **Asset-Action Closure**

它由四条互相配合的动作链组成：

1. `Onboarding` 从首轮登记走向首轮建链引导
2. `Decision` 从“留痕对象”走向“可正式推进的经营对象”
3. `Dashboard / Review / Detail pages` 共同承接“下一步动作”
4. `Current Focus / pending signals` 回到真实经营语义

当前阶段不把以下内容作为主交付对象：

- `Agent Consumption Layer`
- `Cross-Project Convention Asset`
- 新的长期核心业务实体
- 第五个 `DecisionStatus`
- AI 工作台 / 对话式主入口
- 真实连接的重型 GitHub 集成

### 4.3 Onboarding 的正式边界冻结

`Onboarding` 在 `phase10` 中冻结为：

- 继续保留既有六段式最小主线：`welcome / product / repository / module / decision / complete`
- 但正式职责从“逐段登记”升级为“首轮建链引导”
- 目标不是强制用户走固定工作流引擎，而是在创建过程中主动降低“登记后再去多个详情页补链”的使用摩擦
- `Product` 在首轮建链中继续作为唯一主上下文锚点；`Repository / Module / Decision` 只作为 step 级结果与辅助恢复线索，不替代主上下文
- 中途中断后再进入时，只允许依据 canonical facts 与单值 onboarding 恢复读模型回到最近未完成 step，不允许按全局最新实体猜测当前主线

本阶段明确不做：

- 可配置 `Onboarding` 工作流引擎
- 第二套独立草稿系统
- 脱离 canonical `Product / Repository / Module / Decision` 的并列写路径

### 4.4 Decision 生命周期的正式边界冻结

`Decision` 在 `phase10` 中继续冻结为既有四态：

- `proposed`
- `active`
- `superseded`
- `archived`

当前阶段正式结论：

- 不新增 `acknowledged / resolved` 第五态
- `pending decision` 的判定继续锚定 canonical `Decision.status`
- `Decision Detail` 是正式状态推进承接页
- `Dashboard / Daily Review / Current Focus` 的语义必须复用同一套 canonical 状态解释

### 4.5 当前阶段正式消费面

为避免 `Asset-Action Closure` 长成散装全站改造，本阶段正式消费面冻结如下：

- `Onboarding`
- `Dashboard Home`
- `Daily Review`
- `Decision Detail`
- `Product Detail`
- `Module Detail`
- `Repository Detail`

补充约束：

- `Weekly Review` 当前不是本阶段中心承接位，只允许被动享受“动作语义更清晰”的结果，不承接新的复杂动作编排
- `Dashboard` 继续作为经营入口，但不承接第二套局部写路径
- `Detail pages` 必须承接各自实体的“下一步动作”，而不是都回退到列表页或 review 页猜下一步

### 4.6 当前阶段前端交付策略

前端继续统一遵守：

- 单一 `React Web`
- `TanStack Router + TanStack Query`
- 业务写路径唯一 `application` 入口
- `query` 层纯只读
- mutation 固定承接位

当前阶段重点是：

- 把 `Onboarding`、`Decision Detail` 与各实体 detail 页中的动作承接位收敛到固定 owner
- 把 `Dashboard / Review` 中“下一步动作”的组装规则收敛到单值读模型与单值 application owner
- 不在页面组件中散落复制“判断下一步 / 失效刷新 / 成功回流”逻辑
- 当多个 detail CTA 同时成立时，页面必须存在单值主 CTA；主 CTA 优先用于闭当前 canonical 结构缺口，其次才是 `Decision` 推进与 reread 返回

补充冻结：

- `Onboarding` 主线由 `phase10-08` 承接，不允许在 `phase10-09 / 10` 中回头改写其业务编排
- `Decision pending` 主线由 `phase10-09` 承接，不允许在 `phase10-10` 中再定义第二套 detail page pending 解释
- `Product / Module / Repository Detail` 的 CTA inventory 由 `phase10-10` 承接，但必须复用 `phase10-09` 已冻结的 `Decision` canonical 解释

### 4.7 当前阶段后端、合同与数据承接策略

当前阶段继续统一遵守：

- `.proto` 是唯一长期合同源
- `ConnectRPC` 是业务接口正式传输层
- `chi` 继续只承接 router shell、middleware 与非业务端点
- 既有 canonical facts 不复制为第二套动作语义来源

当前阶段重点是：

- 为 `Onboarding` 的首轮建链引导补齐最小正式读写承接
- 在既有四态内补齐 `Decision` 生命周期推进所需的最小合同、服务与 reread 路径
- 为 `Current Focus / pending signals` 与 detail CTA 的统一语义补齐最小读模型与承接矩阵
- 优先复用既有 `Product / Module / Repository / Decision / Review` canonical facts，不新增影子状态表
- 若 `Onboarding` 恢复语义需要额外读取，只允许新增单值恢复辅助读模型，不允许引入第二套草稿真相源

补充冻结：

- 能由既有 canonical facts 组合表达的动作语义，优先复用既有 `.proto / Connect / service / store`
- 只有当页面层会重复编排、且现有合同无法稳定表达“下一步动作”时，才允许新增最小读模型或最小合同
- 任何新增承接位都必须显式回答“它回收了哪段页面层重复逻辑”，否则不得进入实现

### 4.8 当前阶段业务边界原则

为避免 `phase10` 变成“大规模体验重做 + agent-first + 真实连接 + 新实体回归”大杂烩，当前阶段先冻结以下边界：

- 本阶段只交付 `Asset-Action Closure`
- `Agent Consumption Layer` 只允许作为后续 phase 的依赖表达，不进入本阶段实现承诺
- `Cross-Project Convention Asset` 只允许作为候选探索，不进入本阶段实现承诺
- `Current Focus` 与 pending signals 只能回到真实 canonical facts，不允许长出局部“已处理”假象
- `Onboarding` 逻辑化不能演化为工作流引擎
- `Decision` 生命周期不能演化为 `Decision Intelligence`

## 5. 当前阶段完成条件

`phase10` 完成时，至少必须满足：

1. `Onboarding` 已从“逐段登记”升级为“首轮建链引导”，且显著减少后续补链摩擦
2. `Decision` 在既有四态内形成最小但真实的生命周期，并可被正式 reread 消费
3. `Dashboard / Daily Review / Detail pages` 已形成统一的“下一步动作”承接矩阵
4. `Current Focus / pending signals` 已回到真实经营语义，而不是泛化的存在性提示
5. 前端与后端的动作承接位、读写 owner、返回链与 reread 规则已收敛到单值结构
6. 本阶段未把 `Agent Consumption Layer / Cross-Project Convention Asset / 新实体主线` 偷渡为并列主交付
7. 根级入口已回写当前正式 phase 为 `phase10_asset_action_closure_foundation`
