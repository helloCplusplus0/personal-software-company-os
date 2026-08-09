# Personal Software Company OS

# MVP0.2 Next-Stage Direction Review

**Author:** GPT54
**Date:** 2026-08-09
**Role:** 基于原始方案与 MVP0.1 已验收现实，对下一阶段推进方向给出评审性建议
**Document Type:** `review`
**Status:** 供后续 `/plan` 参考，不直接构成正式 phase 命名或执行指令

---

## 1. 目的说明

本文的任务不是重复总结 `phase01 ~ phase05`，也不是直接替代下一阶段的正式 `/plan`。

本文要回答的是一个更关键的问题：

> 在 `MVP0.1` 已经完成“资产登记 + 决策留痕 + 基础复用反馈”之后，PSCO 下一阶段最符合最佳实践的推进方向，究竟应该是什么？

我的结论先写在前面：

> **下一阶段不应优先扩展更多重实体，而应优先把 PSCO 从“可登记的资产系统”推进为“可持续使用的经营操作台”。**

更具体地说，建议下一阶段的核心主题收敛为：

> **Onboarding + Operating Loop + Derived Intelligence**

也就是：

1. 让已有现实资产更低摩擦进入系统；
2. 让 Dashboard/Decision/Product/Module 真正形成日常使用闭环；
3. 让 `Capability / Reuse / Focus` 从事实中被派生出来，而不是要求用户手工维护抽象对象。

---

## 2. 评审依据

本文直接依据以下材料形成判断：

### 2.1 原始长期方案

- `PSCO_0.md`
- `PSCO_1.md`
- `PSCO_2.md`
- `PSCO_3.md`
- `PSCO_4.md`

### 2.2 当前正式推进与已验收现实

- `plan.md`
- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`
- `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`
- `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md`
- `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`

### 2.3 当前技术与实现最佳实践基线

本轮额外核对了与当前仓库主线直接相关的最新文档基线：

- `TanStack Router v1.114.3`：typed search params、`retainSearchParams`、`useNavigate({ from })`
- `TanStack Query v5.90.3`：query key 分层、前缀失效、`placeholderData`
- `Buf v1.42.0`：单一 `.proto` 合同源、`lint` / `breaking` 检查

这些最新基线支持了一个很重要的判断：

> **PSCO 下一阶段完全可以继续在现有技术方案内演进，不需要引入第二套路由事实源、第二套 contract 语义或额外复杂基础设施。**

---

## 3. 我对 PSCO 长期目标的再确认

回看 `PSCO_0.md ~ PSCO_4.md`，PSCO 的长期目标始终不是做一个“模块清单工具”，也不是做一个“AI 聊天入口”。

它真正想成立的是三件事：

1. **把个人开发者的工作对象从“项目”提升为“产品、决策、模块、能力、经营动作”。**
2. **把零散的软件实现沉淀为可复用、可解释、可持续复利的软件资产。**
3. **让 AI 成为基于上下文的增强层，而不是替代人的判断层。**

这意味着 PSCO 的最终形态，应该同时具备四种能力：

### 3.1 经营视角

用户能知道自己在做什么、为什么做、当前最值得推进什么。

### 3.2 产品视角

用户能把产品、反馈、价值主张、关键判断与真实实现连接起来。

### 3.3 工程资产视角

用户能看见哪些模块已经存在、是否稳定、是否复用、是否值得进一步沉淀。

### 3.4 上下文增强视角

AI 能在 `Product / Module / Decision / Review` 上下文中给出总结、提醒、候选方案与文档辅助，但不夺走方向判断权。

所以，PSCO 的长期胜负手并不是“实体够不够多”，而是：

> **是否真的让一个独立开发者把经营、产品、工程和资产积累放到同一套运行系统里。**

---

## 4. MVP0.1 已经真实证明了什么

`phase01 ~ phase05` 结束后，我认为 `MVP0.1` 已经证明了以下事实：

### 4.1 资产登记主线已经成立

仓库已经完成并验收：

- `Module Registry`
- `Release` 最小主线
- `Decision Center`
- `Product Registry`
- `Repository Binding`
- `Dashboard + Feedback`

这说明 PSCO 最小闭环已经不是概念图，而是可运行现实。

### 4.2 单一合同源与单一事实链路已经成立

`.proto` 主线、HTTP envelope、前端 adapter、真实联调验收已经形成稳定一致的工程方式。

这很重要，因为下一阶段不需要回头再讨论“是不是还要第二套 JSON 合同”“是不是要换一套路由状态管理”。

### 4.3 Dashboard 已经初步具备“驾驶舱”雏形

`phase05` 证明了：

- 聚合读是可行的；
- 局部错误与整体成功可以共存；
- 跳转与返回上下文可被正确保留；
- Feedback 可以作为资产视角的最小反馈层存在。

这意味着 PSCO 已经有了“看板入口”，但还没有完全形成“经营动作入口”。

### 4.4 当前 MVP 仍然偏“登记与查看”，还没有完全进入“日常经营操作”

这是我认为下一阶段必须正视的核心现实。

今天的 PSCO 已经能回答：

- 我有哪些 Product / Module / Repository / Decision？
- 它们之间是什么关系？
- Dashboard 当前看到哪些反馈信号？

但它还没有完整回答：

- 我今天应该优先处理什么？
- 我如何把已有现实资产快速导入并开始使用？
- 我怎样做一次周复盘，并把结果回流到 Product / Module / Decision？
- 我如何看见“能力真的在增长”而不是只是数据更多了？

---

## 5. 当前最关键的未验证点

如果直接从 `MVP0.1` 往后推进，我认为真正还没有被验证的，不是“还能不能继续加实体”，而是以下四个问题。

## 5.1 冷启动与导入摩擦仍未被真正解决

原始方案一直强调：PSCO 不应该让用户感觉自己在“填系统”。

但从当前交付现实看，系统虽然已经可用，却仍然比较依赖：

- 手工逐项创建；
- 理解较多对象；
- 已知 Product / Module / Repository 关系；
- 愿意投入足够精力维护结构。

这对于“已经理解 PSCO 的人”是可以接受的，但对于“刚把已有现实搬进系统的人”仍然偏重。

如果下一阶段不先补这个缺口，PSCO 很容易停在：

> 逻辑正确，但进入成本偏高；理念成立，但使用频率不足。

## 5.2 Dashboard 还没有真正接管 operating loop

当前 Dashboard 已经能展示全局情况，但还不够像“经营起点”。

它更像：

- 很好的总览页；

但还不完全像：

- 每天打开就能决定下一步动作的工作起点。

下一阶段必须把 Dashboard、Decision、Product、Module 之间的关系推进到：

> **从“我看见了什么”进入“我接下来做什么、记录什么、为什么这样做”。**

## 5.3 `Capability` 的长期价值被确认，但其派生机制还不够强

原始方案与阶段共识都已明确：

- `Capability` 不应在当前阶段作为重实体强维护；
- 它应该从事实里派生。

问题是：当前派生证据还偏薄。

现在的系统更像是在告诉用户：

- 你有多少模块；
- 哪些地方有缺口；

但还没有足够强地告诉用户：

- 哪些模块已经形成真实能力；
- 哪些能力正在跨产品复用；
- 哪些决策确实提升了资产复利；
- 哪些地方值得进入“模块提取 / 能力沉淀”动作。

## 5.4 Workflow 还停留在文档理念层，没有完全变成产品能力层

`PSCO_3.md` 最重要的价值之一，是把“工作流”当成系统的一等问题。

但在 `MVP0.1` 现实里，workflow 还主要体现为：

- 规格推进 workflow；
- 开发与验收 workflow；
- 文档与阶段治理 workflow。

它还没有充分变成“最终用户在产品内执行 daily/weekly operating loop”的产品能力。

---

## 6. 下一阶段的标准判断

基于以上分析，我给出一个非常明确的标准：

> **下一阶段不应该优先做“更多对象”，而应该优先做“更高频、低摩擦、可复利的运行闭环”。**

这意味着，下一阶段的第一原则应是：

### 6.1 先提高使用频率，再扩展对象宽度

如果用户不能稳定每周回到 PSCO，继续扩展 `Opportunity / Venture / Feature / Experiment` 只会让系统更重。

### 6.2 先让事实可进入、可回顾、可转动作，再做更抽象的分析层

也就是说：

- 先解决导入；
- 再解决 review；
- 再解决 capability insight；
- 最后再扩大长期 domain 宽度。

### 6.3 继续坚持 AI 是增强层

下一阶段可以增加 AI 辅助，但必须建立在真实上下文之上，例如：

- 周报总结；
- 决策草稿辅助；
- 导入字段建议；
- 模块复用提示；
- Dashboard 反馈解释。

而不应变成：

- AI 自动决定产品方向；
- 自动扫描全仓库并声称理解一切；
- 单独做成一级导航式 `AI Assistant`。

---

## 7. 推荐的下一阶段主题

我建议把 `MVP0.2` 的推进主题定义为：

## **从 Asset Registry 走向 Operating System**

具体收敛成三条主线。

### 7.1 主线 A：Onboarding Foundation

目标不是“自动扫描所有代码”，而是：

> **让一个已经有现实项目的用户，可以在 30-60 分钟内把第一批核心资产放进 PSCO，并获得初始价值。**

建议能力包括：

1. **First-run onboarding**
   - 以一条受控流程引导用户先建立首个 `Product`
   - 再补首批 `Repository / Module`
   - 再补关键 `Decision`

2. **Draft-first / partial-entry**
   - 允许对象以最少字段先创建
   - 后续逐步补全，而不是首次录入就要求完整结构

3. **Controlled import helper**
   - 支持从已有仓库、README、发布记录、人工表单中导入基础事实
   - 重点是“辅助录入”，不是“自动理解一切”

4. **Low-friction decision capture**
   - `Decision` 采用 ADR-lite
   - 支持 title-first、context-later 的渐进补全

### 7.2 主线 B：Operating Review Loop

这是我认为最关键的一条。

PSCO 下一阶段应验证：

> **用户能否把 Dashboard 作为 daily / weekly operating 起点来使用。**

建议能力包括：

1. **Daily Review**
   - 进入系统后先看 `Current Focus`
   - 明确本周或今天最值得推进的对象
   - 一键跳到相关 `Product / Module / Decision`

2. **Weekly Review**
   - 汇总新增 `Decision`
   - 汇总新增或变化的 `Module / Release`
   - 汇总缺口、停滞、复用进展
   - 允许记录 review 结论并回流到现有实体

3. **Action handoff without task-manager drift**
   - 不把 PSCO 做成完整任务管理工具
   - 但允许把 review 结论转成最小“后续动作”记录
   - 这些动作必须锚定到既有实体，而不是再发明第二套项目对象

4. **Feedback -> Decision -> Update 闭环**
   - 从 Dashboard 的 gap 或 signal 出发
   - 能快速进入 `Decision`
   - 决策后回流更新相关 `Product / Module / Release`

### 7.3 主线 C：Derived Asset Intelligence

下一阶段应开始验证“复利”不是叙事，而是系统可观测事实。

建议优先做的不是完整 `Capability` 实体，而是以下派生指标：

1. `module_reuse_count`
2. `cross_product_usage_count`
3. `release_freshness`
4. `decision_link_density`
5. `orphan_repository / orphan_module / orphan_product` 数量
6. “候选能力”提示，而不是正式 `Capability` 写入流程

这里的原则是：

> **先做 derived signals，再决定何时把其中一部分升级为更正式的能力模型。**

---

## 8. 我不建议下一阶段优先做的方向

为了避免 `/plan` 时范围再次发散，我也明确给出反向建议。

### 8.1 不建议直接全面引入 `Opportunity / Venture / Feature / Experiment`

这些对象在长期模型中都成立，但现在一起进入，会把系统重新拖回“概念完整优先”。

更稳妥的做法是：

- 先在 review / analysis 层讨论其进入条件；
- 等 operating loop 被验证后，再挑一个最有必要的对象进入。

### 8.2 不建议把下一阶段做成通用项目管理系统

PSCO 可以承接 operating action，但不应该演化为：

- Kanban 主工具；
- Sprint 主工具；
- 通用任务平台；
- 第二个 Jira / Linear。

否则会稀释 `Decision + Asset + Reuse` 的差异化价值。

### 8.3 不建议把自动代码扫描、知识图谱、Rust 智能层作为下一阶段主线

原始方案已经明确这些都不是当前主线。

在数据密度、导入质量、日常使用频率尚未建立之前，直接做智能层只会提高工程成本，降低验证效率。

### 8.4 不建议先做重量级 GitHub OAuth / 全自动导入

下一阶段更需要验证的是：

- 辅助导入有没有真实价值；
- 用户愿不愿意持续维护事实；
- review 是否会被真实使用。

而不是先解决最大自动化问题。

---

## 9. 对下一阶段 `/plan` 的建议形状

以下内容不是正式 phase 命名，只是我建议后续 `/plan` 可以参考的结构。

## 9.1 建议先开一个“单主题 phase”，不要一次并行开多个大方向

推荐优先主题：

> **Operating Loop Foundation**

注意：这只是评审建议名，不是正式 phase 名称。

它应直接承接：

- `phase05-10` 正式规格正文
- `phase05-11` 合同主线
- `phase05-14` 验收结论

## 9.2 建议拆成三个顺序子块，而不是一个大而全的 phase

### Block 1：Cold Start / Onboarding

DoD 建议：

- 新用户可在一次会话内完成首个 `Product + Repository + Module + Decision` 录入
- 首次进入可在不查文档的前提下理解系统基本路径
- 对象允许最小字段创建并逐步补全

### Block 2：Review Loop

DoD 建议：

- 用户可完成一次完整 weekly review
- review 结果能回流到既有实体
- Dashboard 能承接“看见问题 -> 做出动作”的主线

### Block 3：Derived Intelligence

DoD 建议：

- Dashboard 或 detail 页面可展示最小复利指标
- 至少有一组指标可解释“能力增长”或“复用增长”
- 不引入手工维护的重 `Capability` 写入流程

---

## 10. 技术与实现层最佳实践建议

下一阶段在实现上，我建议继续坚持以下约束。

### 10.1 路由状态继续留在 TanStack Router

基于当前仓库经验与最新 Router 文档，下一阶段涉及 onboarding、review、multi-step return path 时，应继续：

- 使用 typed search params
- 保持 `useNavigate({ from })`
- 在需要跨层保留来源上下文时优先使用 `retainSearchParams`

不要再引入第二套全局导航事实源。

### 10.2 Dashboard / Review 读模型继续采用 Query-first 思路

基于当前 Query 文档，建议继续：

- 统一 query key 设计
- 使用前缀失效刷新相关聚合读
- 对 review 或分页列表使用 `placeholderData` 保持高密度 UI 稳定

这样能保持页面密度与切换流畅性，而不是在每次刷新时回退到大面积 loading。

### 10.3 合同层继续只认 `.proto`

基于当前 Buf 实践，下一阶段新增 review / onboarding / derived metrics 接口时，应继续：

- 以 `.proto` 为单一合同源
- 维持 `lint` 与 `breaking` 检查
- 不在页面层偷偷长出第二套 contract 语义

### 10.4 后端继续坚持单服务、清晰 query / application 边界

下一阶段更像是在已有系统上补 operating layer，而不是重构成新架构。

因此应继续避免：

- 微服务拆分
- 新的复杂基础设施
- 为未来智能层提前引入不必要依赖

---

## 11. 我对 `MVP0.2` 的单句定义

如果要用一句话定义我建议的下一阶段，我会写成：

> **MVP0.2 的任务，不是让 PSCO 拥有更多概念，而是让它真正成为个人软件公司的每周 operating console。**

换一种更产品化的表达：

> **从“我记录了什么资产”推进到“我知道接下来该做什么，以及这些动作如何继续沉淀为资产”。**

---

## 12. 最终结论

我的最终判断如下：

1. `MVP0.1` 已经成功证明 PSCO 的最小资产主线是成立的。
2. 下一阶段最值得做的，不是继续铺开完整长期 domain，而是优先补足 **导入、review、derived intelligence**。
3. 若下一阶段优先把 PSCO 做成高频 operating loop，长期的 `Feature / Opportunity / Experiment / AI enhancement` 才会有稳固落点。
4. 若下一阶段反而直接扩实体、扩自动化、扩智能层，PSCO 很容易重新变成“理念强、使用弱”的系统。

因此，我支持后续正式 `/plan` 以以下方向为主轴：

> **先把 PSCO 做成可持续使用的 operating system，再把它扩展成更完整的 company OS。**

这会是比“继续增加对象覆盖面”更稳、更符合最佳实践、也更接近原始愿景的一步。
