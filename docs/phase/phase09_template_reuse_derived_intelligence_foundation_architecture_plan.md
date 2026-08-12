# phase09_template_reuse_derived_intelligence_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase09_template_reuse_derived_intelligence_foundation` 的架构规划文档。

`phase08` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并正式交付了 `Operating Review Loop` 最小经营回路。根据 `PSCO-mvp03-summarize-feedback.md` 已冻结的仲裁结论，下一步不应继续扩长期对象宽度，也不应提前切到 `dry-run` 或 `AI Context Enhancement`，而应把 **`Template Reuse`** 与 **`Derived Intelligence Deepening`** 作为承接 `phase08` 的支撑能力正式建立。

因此，`phase09` 的职责不是再造第二条业务主线，而是交付一套面向“下一次创造”的最小加速支撑层：让用户能从既有复用事实中低摩擦开始新建动作，并在 review / create 过程中看到能够支撑下一步判断的最小派生提示。

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `docs/README.md`
7. `PSCO-mvp03-summarize-feedback.md`
8. `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`
9. `docs/phase/phase08_operating_review_loop_foundation_architecture_plan.md`
10. `docs/phase/phase08_operating_review_loop_foundation_dev_plan.md`
11. `docs/phase/phase08_operating_review_loop_foundation_shared_baseline.md`
12. `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
13. `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_shared_baseline.md`
14. `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
15. `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
16. `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`

补充说明：

- `PSCO-mvp03-summarize-feedback.md` 冻结的是 `mvp0.3` 的方向排序与范围边界，不直接替代当前 `/plan`
- `phase08-11` 是本阶段唯一直接业务上游验收结论，说明 `Operating Review Loop` 已稳定可消费
- `phase06` 提供 `module_reuse_summary / capability_summary` 等已落地基础读模型，是本阶段支撑能力设计的直接资产来源
- `phase03 / phase04 / phase05` 的正式规格继续提供 `Decision / Product / Dashboard` 等既有 canonical 动作与页面承接位

## 3. 本阶段目标

`phase09` 的目标是：

> 在不引入新一级核心实体、不把系统演化成模板平台或 AI 工作台的前提下，交付 `Template Reuse + Derived Intelligence Deepening` 的最小支撑能力，使用户能够从 review 或既有复用事实中更快进入下一次 `Product` 创建，并获得足以支撑动作的最小提示。

本阶段需要回答的核心问题：

1. `Template Reuse` 的正式资产应如何定义，才能服务“下一次创造”而不是长成独立平台
2. 模板级复用应从哪些既有事实源派生，才能不复制第二套业务真相源
3. 哪些页面承担正式消费位，才能让支撑能力直接服务 `phase08` 已交付的经营回路
4. `Derived Intelligence Deepening` 最小需要承接哪些提示，才能帮助用户判断“下一步应该复用什么、补齐什么、推进什么”
5. 提示语义如何保持解释性与动作导向，而不滑向“更复杂的统计展示”或“AI 工作台”
6. `Product Create` 如何继续保持唯一 canonical 写入承接位，而不因为模板预填再长出第二套创建路径
7. 支撑能力相关读模型、合同、前后端 owner 应如何收敛，避免编排散落到页面、query 与展示组件
8. `phase09` 完成时，如何证明“复用从可见走向可用、派生智能从展示走向动作支撑”

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase09` 必须直接承接：

- `PSCO-mvp03-summarize-feedback.md`
- `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`
- `docs/phase/phase08_operating_review_loop_foundation_*`
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
- `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_shared_baseline.md`
- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
- `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
- `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`

不允许在本阶段重新解释：

- `Decision` 在 MVP 中的中心地位
- `.proto` 作为唯一长期合同源
- `phase08` 已完成的 `Operating Review Loop` 收口结论
- `mvp0.3` 的结构：`Operating Review Loop` 为中心主线，`Template Reuse + Derived Intelligence` 为支撑能力，`dry-run` 为独立验收闸

### 4.2 当前阶段单一主交付能力

`phase09` 的单一主交付能力冻结为：

> **Next-Creation Acceleration Support**

它由两类互相配合的支撑能力组成：

- `Template Reuse`：把既有复用事实转化成可直接用于新建动作的低摩擦起点
- `Derived Intelligence Deepening`：把既有反馈、复用与能力摘要转化成可支撑下一步动作的最小解释性提示

当前阶段不把以下内容作为主交付对象：

- 新的长期核心业务实体
- 独立模板平台
- 独立智能中心 / AI 工作台
- `Real-Project Dry-Run`
- 对 `Operating Review Loop` 主线的大范围回写重做

### 4.3 Template Reuse 的边界冻结

`Template Reuse` 在 `phase09` 中继续冻结为：

- `Module` 组合快照
- 新建 `Product` 时的预填辅助
- 预填后继续编辑并完成创建

其正式职责是：

1. 把当前系统内已经存在的模块组合与复用事实，转化成下一次创造的起点
2. 让用户不必从空白 `Product Create` 开始手工重建已有组合
3. 服务于 review 后的下一步动作，而不是成为独立管理对象

因此，本阶段明确不做：

- 完整模板平台
- 参数化模板体系
- 模板版本管理
- 独立模板 CRUD 主线

### 4.4 Derived Intelligence Deepening 的边界冻结

`Derived Intelligence Deepening` 在 `phase09` 中继续冻结为：

- `capability gap` 最小提示
- `reuse opportunity` 最小提示
- 与 `review / create` 直接相连的解释性指标与行动文案

其正式职责是：

1. 帮助用户判断“当前缺什么能力、下一步可以复用什么”
2. 让 `phase06` 的 `module_reuse_summary / capability_summary` 从“能看见”升级到“能指导动作”
3. 服务 `review -> create / update` 的既有正式动作链

因此，本阶段明确不做：

- AI 主线
- 独立智能工作台
- 自动生成长期策略
- `Decision Intelligence` 独立主线化

### 4.5 当前阶段正式消费面

为避免范围失控，当前阶段正式消费面冻结如下：

- `Weekly Review`：支撑能力的主要诊断入口
- `Product Create`：模板预填与提示消费的主要动作入口
- `Product Detail / Dashboard / Review reread`：结果回流与支撑能力复读验证入口

补充约束：

- `Daily Review` 不承担复杂模板编排或重型智能提示主入口
- `Product Create` 仍是唯一正式创建承接位；模板选择或推荐只允许向该 canonical 写路径输送预填上下文
- Dashboard 若继续展示摘要，只承担导航与轻量状态提示，不承担第二套正式模板创建主线

### 4.5.1 `Product Create` 模板 handoff 合同冻结

`phase09` 对 `/products/new` 的模板来源承接方式冻结如下：

- 正式 handoff 继续使用 route search 参数承接，而不是在页面层临时塞入全量快照 payload
- 当前阶段新增的正式来源参数只允许是：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=weekly-review | dashboard | product-detail`
- `fromDashboard / dashboardSection / dashboardReturnTo` 允许与模板来源参数共存，用于保留返回链
- `fromList`、`fromModuleDetail` 与 `fromTemplateReuse` 不允许并列成立；若 `fromTemplateReuse=true`，则模板来源是本次 create 的唯一业务来源语义
- `Product Create` 页面只读取 `templateCandidateId` 对应的预填读模型；不得把完整 `Module` 组合快照直接塞进 search 参数

补充冻结：

- `templateCandidateId` 对前端是 opaque string；其编码与解码由后端 owner 负责，前端不得自行拼装业务语义
- 模板预填只改变 create 页的初始表单与候选模块上下文，不改变 `Product Create` 作为唯一 canonical write path 的事实

### 4.6 当前阶段前端交付策略

前端继续统一遵守：

- 单一 `React Web`
- `TanStack Router + TanStack Query`
- 业务写路径唯一 `application` 入口
- `query` 层纯只读
- mutation 固定承接位

当前阶段重点是：

- 把模板预填、推荐选择与派生提示读取收敛到固定切片 owner
- 让 `Weekly Review` 与 `Product Create` 只消费已收敛的 read / application 承接位
- 保持 `Product Create` 的正式写入 owner 不变，不在页面、弹层或 review footer 中复制第二套创建语义
- 优先复用 `phase06` 与 `phase08` 已交付的读模型、路由上下文与返回链，而不是新建并列宿主

### 4.7 当前阶段后端、合同与数据承接策略

当前阶段继续统一遵守：

- `.proto` 是唯一长期合同源
- `ConnectRPC` 是业务接口正式传输层
- `chi` 继续只承接 router shell、middleware 与非业务端点
- 既有领域事实源不复制为第二套并列 canonical 数据主线

当前阶段重点是：

- 复用 `phase06` 已交付的 `reuse_summary` 基础读模型
- 为“模板候选 / 组合快照 / 派生提示 / create 预填”补齐最小正式合同
- 优先从既有 `Product / Module / Binding` canonical 事实派生模板与提示；`Review` 只提供消费作用域与返回链上下文，不直接生成模板候选
- 若后续 `/spec` 证明需要轻量快照记录，其身份也只能是支撑能力资产，不得升级为新的长期核心实体

### 4.7.1 模板候选的 canonical 来源、去重键与排序规则

当前阶段模板候选的生成规则冻结如下：

- canonical 来源：当前已持久化的 `product_modules` 绑定事实
- 候选生成口径：以“一个 Product 当前绑定的全部 Module 集合”作为原始候选输入；空集合不参与模板候选
- 去重键：对候选中的 `module_id` 去重、升序排序后形成 normalized module-set key；相同 key 的多个 Product 合并为同一模板候选
- `templateCandidateId` 必须由该 normalized key 单向派生，但编码方式由后端 owner 负责冻结
- 当前阶段模板候选默认按以下顺序排序：
  1. `source_product_count DESC`
  2. `total_reuse_product_count DESC`
  3. `latest_source_product_updated_at DESC`

补充冻结：

- 当前阶段候选不从未持久化草稿、页面本地状态或临时 review 表单输入派生
- 若某个 Product 当前没有任何绑定 Module，则不得把它提升为模板候选
- 当前阶段允许在后续 `/spec` 中细化字段名，但不得改写上述 canonical 来源、去重键与排序口径

### 4.7.1.1 Weekly Review 模板候选选择语义冻结

为避免 `capability_gap_hint` 与模板 CTA 在前后端各自实现一套“当前候选”判断逻辑，当前阶段继续冻结：

- `Weekly Review` 中的模板候选为 **单选模型**
- 若当前存在候选，则按既定排序结果取第一个候选作为默认 active candidate
- 用户可以在当前 review 会话中切换 active candidate，但同一时刻只允许存在一个 active candidate
- `Weekly Review` 内与模板直接相关的派生提示、CTA 与解释文案，都只允许基于当前 active candidate 计算
- 若当前没有任何模板候选，则 `reuse_opportunity_hint` 与依赖模板上下文的 `capability_gap_hint` 都必须返回成功空态；当前阶段不再引入“未选模板时退回 generic review focus”这一并列口径

### 4.7.2 轻量快照记录的默认决策

为避免 `/spec` 与实现阶段出现“读时派生”和“新建持久化快照表”两套并列方案，当前阶段先冻结默认决策：

- `phase09` 默认采用 **读时派生模板候选 + 按 `templateCandidateId` 读取预填详情** 的方案
- 当前阶段默认 **不** 新增独立模板快照持久化表，不把 `Module` 组合快照落成新的长期事实源
- 只有在 `phase09-04` 边界冻结中拿到明确证据，证明既有 canonical 事实无法稳定支撑：
  - 候选读取
  - `templateCandidateId` 定位
  - `Product Create` 预填
  才允许升级为受控的轻量快照记录方案

若触发升级，仍必须同时满足：

- 快照记录只承担支撑能力资产身份，不得升级为新的核心实体
- 原始事实源仍然是 `Product / Module / Binding` canonical 数据
- 前后端不得同时保留“纯读时派生”和“持久化快照”两套长期稳态实现

### 4.7.3 派生提示矩阵冻结

当前阶段只冻结两类正式提示：

1. `reuse_opportunity_hint`
   - 事实来源：模板候选聚合结果 + `module_reuse_summary`
   - 触发条件：当前存在可用模板候选，且候选对应的组合已被至少一个已持久化 Product 证明可复用
   - 主要消费面：`Weekly Review`
   - 正式 CTA：进入 `/products/new` 的模板预填 handoff

2. `capability_gap_hint`
   - 事实来源：`capability_summary` + 当前 active template candidate
   - 触发条件：当前存在 active template candidate，且当前 review 作用域内存在高频 capability 未被该模板组合覆盖
   - 主要消费面：`Weekly Review`、`Product Create`
   - 正式 CTA：进入既有 `Module Registry` / `Module Detail` canonical 路径继续补齐，而不是在提示层内联第二套写动作

补充冻结：

- 当前阶段不再新增第三类“泛化智能提示”
- 每类提示都必须同时给出：触发事实、解释文案与正式下一步动作
- 若某类提示当前只有解释、没有稳定 canonical CTA，则该提示不得进入 `phase09` 正式范围

### 4.7.4 模板创建成功后的承接链冻结

为避免模板复用在 `Product Create` 成功后退化为“只看见预填，没有正式下一步动作”，当前阶段继续冻结：

- `Product Create` 提交成功时，既有 `ProductCreate` canonical mutation owner 只负责创建 Product，不在同一写动作内自动写入 `product_modules`
- 若本次创建来自模板预填，成功跳转到 `Product Detail` 时必须继续携带：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
- 上述参数在 `Product Detail` 中只承担模板来源复读与下一步动作承接语义，不得升级为新的写路径来源
- `Product Detail` 必须能够基于该来源上下文展示模板来源摘要 / 候选 `Module` 组合摘要，并提供进入既有 canonical `Product <-> Module Binding` 路径的正式 CTA
- 当前阶段用“创建成功后保留模板来源语义并显式承接到 canonical binding path”证明模板组合没有丢失；不要求在 `phase09` 中把模板组合自动落成新 Product 的已绑定事实

### 4.8 当前阶段与后续 dry-run 的依赖关系

`phase09` 不是 `mvp0.3` 的收口终点。

后续依赖关系冻结如下：

1. 没有 `phase09`，`phase06` 的复用感知仍停留在“可见”而非“可用”
2. 没有 `phase09`，`phase08` 的 review loop 也缺少“如何更快开始下一次创造”的正式支撑层
3. 只有 `phase09` 把模板复用与派生提示变成可运行支撑能力后，`dry-run` 才有资格验证它们是否真的带来复利加速
4. `dry-run` 的正式入口与命名，留给后续单独 `/plan` 建立时再冻结

## 5. 当前阶段完成条件

`phase09` 完成时，至少必须满足：

1. 用户能够从既有复用事实进入“新建 `Product` 预填”正式路径
2. `Product Create` 仍保持唯一 canonical 创建承接位
3. 用户在 `Weekly Review` 或等价正式消费面中，能够看到最小 `capability gap / reuse opportunity` 提示
4. 提示语义已能够支撑下一步动作，而不仅是多展示一层统计信息
5. 模板候选来源、active candidate 选择语义与 `capability_gap_hint` 触发条件已收敛为单值口径
6. 创建成功后的 `Product Detail` 能继续复读模板来源语义，并显式提供进入 canonical `Product <-> Module Binding` 路径的下一步动作
7. `Template Reuse` 与 `Derived Intelligence Deepening` 的合同、读模型、前后端 owner 已收敛到单值承接位
8. 本阶段未把模板复用扩写为独立平台，也未把派生智能扩写为 AI 主线
9. 根级入口已能明确表达：`phase09` 已作为 `phase08` 后续支撑能力 phase 的正式 `/plan` 入口建立，`dry-run` 继续保留为后续独立验收闸
