# Personal Software Company OS

# MVP0.2 Next-Stage Direction Plan

**Author:** DeepSeek-V4-Flash
**Date:** 2026-08-09
**Role:** 基于原始方案（`PSCO_0.md ~ PSCO_4.md`）与 MVP0.1（`phase01 ~ phase05`）已验收现实，系统输出 PSCO 下一阶段推进方向的标准计划文档，供后续正式 `/plan` 参考评价
**Document Type:** `review`
**Status:** 供后续 `/plan` 参考，不直接构成正式 phase 命名或执行指令

---

## 1. 文档定位与结论先行

本文不重复总结 `phase01 ~ phase05` 的交付清单，也不替代下一阶段的正式 `/plan`。

本文要完成两件事：

1. 重新确认 PSCO 的长期核心价值与最终呈现形态（对应第一项任务）；
2. 基于 MVP0.1 的真实工程现实，给出下一阶段最符合最佳实践的推进方向（对应第二、三项任务）。

**我的结论先写在前面：**

> **MVP0.2 的第一原则，不是继续扩展对象宽度，也不应把长期理论模型中的 `Feature / Opportunity / Venture / Experiment` 一次性放量推进，而是把 PSCO 从"可登记、可查看的资产系统"推进为"可持续使用、可回流、可复利的经营运行系统"。**

更具体地，我建议下一阶段收敛为一个单主题 phase：

> **Operating Loop Foundation（运行闭环地基）**

并拆成三条顺序主线：

1. **Cold Start / Onboarding**（让已有现实资产低摩擦进入系统）
2. **Review Loop**（把 Dashboard 变成 daily / weekly 经营起点并回流动作）
3. **Derived Asset Intelligence**（让能力增长成为可观测事实，而非叙事）

同时，我作为 DeepSeek-V4-Flash，会额外强调一个容易被"概念讨论"掩盖的工程现实：

> **在规划新方向之前，必须先还清 MVP0.1 遗留的规格合规负债（导出 / 备份），否则新的 operating layer 会叠加在未闭合的底座之上。**

---

## 2. 第一项任务：PSCO 核心价值、关键功能与最终形态（从原始方案再确认）

### 2.1 一句话定位

`PSCO_0.md` 给出了 PSCO 的根本定位：

> 帮助一个个人开发者具备一家软件公司的组织能力，使其能够持续发现机会、验证需求、构建产品、积累软件能力，并形成长期复利资产。

它的核心判断非常关键：

> 软件开发的核心竞争力，正在从"编码能力"转向"组织能力 + 判断能力 + 资产积累能力"。

也就是说，PSCO 不是"更快写代码"的工具，而是"个人软件公司的运行系统"。

### 2.2 三大核心理念（Build / Accumulate / Compound）

`PSCO_0.md` 用三个理念定义了 PSCO 的价值内核：

| 理念 | 含义 | 对应动作 |
| --- | --- | --- |
| **Build（持续创造）** | 个人开发者不是一次性项目开发者，而是持续寻找机会、构建产品 | `Opportunity -> Product -> Software` |
| **Accumulate（持续积累）** | 每次开发都应留下资产，从"项目归档"升级为"提取能力 -> 形成模块 -> 进入资产库" | `项目 -> 模块 -> 资产库` |
| **Compound（持续复利）** | 长期价值来自能力复用，模块随版本演进形成能力，被未来产品反复复用 | `Module v1 -> v2 -> Capability -> 复用` |

这三点直接决定了 PSCO 的差异点：它管理的是"我形成了多少能力"，而不是"我完成了多少任务"。

### 2.3 关键功能与长期对象链

`PSCO_1.md`（Document 02/03）给出了完整的长期运行模型与领域模型：

```
Opportunity -> Venture -> Product -> Feature -> Module -> Release -> Capability
辅助实体：Decision / Experiment / Repository
```

`PSCO_2.md`（Document 04/05）进一步把 PSCO 的核心差异化定位在 **Module System**：

> Module = Capability Boundary + Implementation + Interface + Knowledge + History

它强调 Module 不是代码包（npm package），而是"软件公司资产"，并按能力（Foundation / Application / Domain / AI / Data）而非语言分类，经历 `Prototype -> Candidate -> Internal -> Stable -> Commercial` 生命周期。

`PSCO_3.md`（Document 06/07）定义了两件事：

- **工作流引擎**：`Think -> Validate -> Build -> Extract -> Reuse -> Compound`，让 PSCO 成为"驱动个人软件公司持续运行的工作流系统"而非"记录工具"；
- **AI 边界**：AI 永远是增强层，不是决策层；绝对不做自动代码扫描、自动知识图谱、AI 自动判断方案。

`PSCO_4.md`（Document 08）定义了最终 UX 形态：

> PSCO 不是 Jira / Notion / GitHub / AI Chat，而是 **Personal Software Company Command Center**（个人软件公司控制台/驾驶舱）。

它的体验目标不是"让用户管理更多东西"，而是"让用户更高效地经营自己的软件公司"。

### 2.4 最终呈现形态（长期北向）

综合 `PSCO_0 ~ PSCO_4`，PSCO 的最终形态是一个具备四重视角的个人软件公司运行系统：

1. **经营视角**：知道自己在做什么、为什么做、当前最值得推进什么（Dashboard / Current Focus / Venture）；
2. **产品视角**：把产品、价值主张、用户、模块、决策、发布串成一条可追溯的价值链（Product Registry）；
3. **工程资产视角**：看见哪些模块已沉淀、是否稳定、是否被复用、是否值得进一步抽象（Module Library / Capability）；
4. **上下文增强视角**：AI 在 `Product / Module / Decision / Review` 上下文中做总结、提醒、候选与文档辅助，但不抢走判断权。

一句话定义最终形态：

> **让一个独立开发者把经营、产品、工程与资产积累放进同一套可长期复利的运行系统里。**

---

## 3. 第二项任务：MVP0.1 的真实推进基础（从仓库现实再确认）

### 3.1 已验收并完结的推进事实

`plan.md` 明确 `phase01 ~ phase05` 均已走完 `/plan -> /spec -> 实现 -> 验收 -> 收口`：

| Phase | 交付主题 | 唯一规格收敛入口 | 状态 |
| --- | --- | --- | --- |
| phase01 | MVP 规格收敛 | `mvp_spec_v0.1.md` | completed |
| phase02 | Module Registry | `module_registry_spec_v0.1.md` | completed |
| phase03 | Decision Center | `decision_center_spec_v0.1.md` | completed |
| phase04 | Product Registry + Repository Binding | `product_repository_binding_spec_v0.1.md` | completed |
| phase05 | Dashboard + Feedback | `dashboard_feedback_spec_v0.1.md` | completed |

### 3.2 真实工程落地的骨架（我实际核对了仓库代码结构）

- **前端（`frontend/src`）**：基于 `React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind + shadcn/ui`，已形成 `features/dashboard`、`features/module-registry`、`features/product-registry`、`features/decision-center`、`features/repository-binding` 五个 feature 模块，以及对应的 `routes/*` 与 `routeTree.gen.ts`。
- **后端（`backend`）**：基于 Go 模块化单体，`internal/` 下已形成 `moduleregistry`、`decisioncenter`、`productregistry`、`repositorybinding`、`dashboard`、`platform` 六个领域包，采用 `handler / service / repository / candidate / errors / validate / types / response` 的分层职责。
- **合同（`proto`）**：已建立 `module_registry`、`decision_center`、`product_registry`、`repository_binding`、`dashboard`、`common` 六个 `.proto` 包，并对应生成 Go 与 TS 代码，`proto/Makefile` 含 `build / gen / lint / breaking / clean` 五个受控 target。
- **数据（`database`）**：已有 6 张 migration、多组基线 seed 与 reset seed，以及 `init_db.sh / run_seeds.sh / reset_*_acceptance.sh` 等受控脚本，支持"先清空再加载指定 fixture"的可重复验收语义。
- **验收实证**：`phase05-14` 已证明 Dashboard 的 `overview / feedback-signals / recent-activities` 三类聚合读在真实前后端与数据库上跑通，空状态、有数据、局部错误、跳转返回路径均形成可复验矩阵。

### 3.3 MVP0.1 已经被证明成立的事实

1. **最小资产主线真实成立**：Module / Release / Decision / Product / Repository / Dashboard 不再是概念图，而是可运行、可复验、可交付的现实。
2. **单一合同源与单一事实链路成立**：`.proto -> HTTP envelope -> 前端 adapter -> 真实联调验收` 已经形成稳定一致的工程方式，不存在第二套 JSON 合同或第二套路由事实源。
3. **Dashboard 已具备"驾驶舱雏形"**：聚合读、局部错误隔离、跳转/返回上下文保留、Feedback 作为最小反馈层均已落地。
4. **工程纪律得到验证**：`spec -> 实现 -> 子代理复核 -> 验收 -> 收口` 的流程门禁在五个 phase 中持续生效。

### 3.4 MVP0.1 的真实边界（这决定下一阶段的起点）

尽管 MVP0.1 成立，但作为深挖工程的视角，我必须如实指出以下四点现状：

#### 3.4.1 仍偏"登记与查看"，尚未进入"日常经营操作"

今天的 PSCO 能回答：

- 我有哪些 Product / Module / Repository / Decision？
- 它们之间是什么关系？
- Dashboard 当前看到哪些反馈信号？

但还不能自然回答：

- 我今天应该优先处理什么？
- 我如何把已有现实资产快速导入并开始使用？
- 我做一次周复盘并回流结果到实体？
- 我如何看见"能力真的在增长"而不是"数据更多了"？

这一条是 MVP0.1 与"个人软件公司运行系统"之间最主要的距离。

#### 3.4.2 冷启动 / 导入摩擦仍是真实缺口

原始方案反复强调"PSCO 不应该让用户感觉自己在填系统"。但从当前代码看，录入路径仍以"手工逐项创建 + 理解较多对象 + 已知绑定关系"为主。对"已理解 PSCO 的人"可用，对"首次把已有现实搬进系统的人"仍偏重。

#### 3.4.3 存在规格合规负债（导出 / 备份未落地）——这是必须正视的工程事实

`mvp_spec_v0.1.md` §7.3 / §7.4 明确要求：

- 导出：面向用户带走核心资产数据，范围至少覆盖 `Product / Module / Release / Repository / Decision` 及基础绑定关系，**"不得把导出留成后续再说"**；
- 备份：至少提供一种面向当前实例的基础备份路径，**不依赖 GitHub 或第三方平台作为唯一前提**。

但我在后端代码中**未检索到任何 export / backup / dump 对应的实现或接口**。这意味着：

- 这是一条"规格已要求、实现未落地"的合规欠账；
- 它对"数据所有权优先（Local First）"这一核心原则是直接违背的；
- 若在未闭合此负债的情况下直接铺开 operating layer，会让新功能叠加在未闭合的数据主权底座上。

我的判断：**下一阶段 `/plan` 在规划 operating loop 之前或之内，必须显式处理这条负债，而不是继续把它当作"后续再说"。**

#### 3.4.4 workflow 仍停留在"项目治理层"，未转化为"产品经营层"

`PSCO_3.md` 把"工作流"当成系统的一等问题。但在 MVP0.1 里，workflow 更多体现为规格推进、开发验收、文档治理这一类"项目级工作流"，尚未转化为最终用户在产品内执行的 daily / weekly operating loop。

---

## 4. 下一阶段的判断标准

基于以上，我给出与项目长期愿景一致的判断标准：

> **下一阶段的第一原则是"先提高使用频率与回流闭环，再扩展对象宽度"，而不是反过来。**

具体三条：

1. **先让事实可进入、可回顾、可转动作，再做更抽象的分析层。**
   → 顺序应为：导入 -> review -> derived intelligence -> 再考虑长期 domain 宽度。

2. **先让"复利"可被观测，再决定是否升级为正式模型。**
   → `Capability` 继续作为派生层，先做 derived signals，再决定何时升级为更正式的能力模型。

3. **继续坚持 AI 是增强层。**
   → 下一阶段可以增加 AI 辅助（周报总结、决策草稿、导入字段建议、模块复用提示、反馈解释），但绝不变成"AI 决定产品方向 / 自动扫描全仓库 / 独立一级导航 AI Assistant"。

---

## 5. 我不建议下一阶段优先做的方向（反向清单）

为了给后续 `/plan` 划清边界，我明确给出反向建议：

1. **不建议一次性全面引入 `Opportunity / Venture / Feature / Experiment`。**
   这些对象在长期模型中都成立，但现在一起进入会把系统拖回"概念完整优先"。更稳的做法是：先在 review / analysis 层讨论进入条件，等 operating loop 验证后再挑最必要的一个进入。

2. **不建议把 PSCO 做成通用项目管理系统。**
   PSCO 可以承接 operating action，但不应演化为 Kanban / Sprint / 通用任务平台，否则会稀释 `Decision + Asset + Reuse` 的差异化价值。

3. **不建议把自动代码扫描、知识图谱、Rust 智能层作为下一阶段主线。**
   原始方案已明确这些不是当前主线；在数据密度、导入质量、日常使用频率未建立之前，直接做智能层只会提高工程成本、降低验证效率。

4. **不建议先做重量级 GitHub OAuth / 全自动导入。**
   下一阶段更需要验证的是"辅助导入有没有价值、用户愿不愿意持续维护事实、review 是否被真实使用"，而不是先解决最大自动化问题。

---

## 6. 推荐的下一阶段主题与主线

我建议把 MVP0.2 的推进主题定义为：

> **从 Asset Registry 走向 Operating System**

收敛成三条顺序主线（与 GPT54 评审在三线结构上一致，但我会从工程落地角度给出更具体的 DoD 与顺序依赖）：

### 主线 A：Cold Start / Onboarding Foundation

目标不是"自动扫描所有代码"，而是：

> **让一个已有现实项目的用户，能在 30-60 分钟内把第一批核心资产放进 PSCO，并获得初始价值。**

建议能力：

1. **First-run onboarding**：以受控流程引导用户先建首个 `Product` -> 补首批 `Repository / Module` -> 补关键 `Decision`；
2. **Draft-first / partial-entry**：允许对象以最少字段先创建、后续逐步补全，降低首次录入负担；
3. **Controlled import helper**：从已有仓库、README、发布记录、人工表单导入基础事实，重点是"辅助录入"，不是"自动理解一切"；
4. **Low-friction decision capture**：`Decision` 采用 ADR-lite，支持 title-first、context-later 的渐进补全。

### 主线 B：Review / Operating Loop

这是我认为最关键的一条：

> **验证用户能否把 Dashboard 作为 daily / weekly operating 起点来使用。**

建议能力：

1. **Daily Review**：进入系统先看 `Current Focus`，明确今天最值得推进对象，一键跳到相关 `Product / Module / Decision`；
2. **Weekly Review**：汇总新增 `Decision`、新增/变化的 `Module / Release`、缺口/停滞/复用进展，允许记录结论并回流到现有实体；
3. **Action handoff without task-manager drift**：不把 PSCO 做成完整任务工具，但允许把 review 结论转成最小"后续动作"，且动作必须锚定到既有实体；
4. **Feedback -> Decision -> Update 闭环**：从 Dashboard 的 gap / signal 出发进入 `Decision`，决策后回流更新相关 `Product / Module / Release`。

### 主线 C：Derived Asset Intelligence

> **开始验证"复利"不是叙事，而是系统可观测事实。**

建议优先做派生指标，而非完整 `Capability` 实体：

- `module_reuse_count`
- `cross_product_usage_count`
- `release_freshness`
- `decision_link_density`
- `orphan_repository / orphan_module / orphan_product`
- "候选能力"提示（而非正式 `Capability` 写入流程）

原则：**先做 derived signals，再决定何时把其中一部分升级为更正式的能力模型。**

### 贯穿主线：Export / Backup 合规闭合（我的差异化补充）

在三线之上，我建议 MVP0.2 必须包一条贯穿性工作项：

> **补全 `mvp_spec §7.3 / §7.4` 已要求但未落地的导出 / 备份能力。**

理由：

- 这是"数据所有权优先（Local First）"的硬约束，是 PSCO 差异化价值的底座；
- 现有代码无对应实现，属于真实规格负债；
- 导出能力同时是 Onboarding（把已有数据带进系统）与 Review（把经营结果带走）的天然对称能力，放到 MVP0.2 闭合最自然。

---

## 7. 对下一阶段 `/plan` 的建议形状

### 7.1 建议先开一个"单主题 phase"，不要一次并行开多个大方向

推荐优先主题（评审建议名，非正式 phase 名）：

> **Operating Loop Foundation**

它应直接承接：

- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
- `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/`
- `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`

### 7.2 建议的三块顺序子阶段与 DoD

#### Block 1：Cold Start / Onboarding

DoD 建议：

- 新用户可在一次会话完成首个 `Product + Repository + Module + Decision` 录入；
- 首次进入可在不查文档前提下理解系统基本路径；
- 对象允许最小字段创建并逐步补全；
- 导出 / 备份基础路径已闭合（与贯穿主线对齐）。

#### Block 2：Review Loop

DoD 建议：

- 用户可完成一次完整 weekly review；
- review 结果能回流到既有实体；
- Dashboard 能承接"看见问题 -> 做出动作"的主线。

#### Block 3：Derived Intelligence

DoD 建议：

- Dashboard 或 detail 页面可展示最小复利指标；
- 至少有一组指标可解释"能力增长 / 复用增长"；
- 不引入手工维护的重 `Capability` 写入流程。

---

## 8. 技术与实现层最佳实践建议

下一阶段在实现上，我建议继续坚持以下约束，避免引入第二套事实源。

### 8.1 路由状态继续留在 TanStack Router

涉及 onboarding、review、multi-step return path 时，应继续使用：

- typed search params；
- `useNavigate({ from })`；
- 跨层保留来源上下文时优先 `retainSearchParams`。

不要引入第二套全局导航事实源。

### 8.2 Dashboard / Review 读模型继续采用 Query-first

- 统一 query key 设计；
- 使用前缀失效刷新相关聚合读；
- 对 review / 分页列表使用 `placeholderData` 保持高密度 UI 稳定。

### 8.3 合同层继续只认 `.proto`

- 以 `.proto` 为单一合同源；
- 维持 `lint` 与 `breaking` 检查；
- 不在页面层长出第二套 contract 语义。

### 8.4 后端继续坚持单服务、清晰 query / application 边界

下一阶段更像在已有系统上补 operating layer，不是重构成新架构。应继续避免微服务拆分、新的复杂基础设施、为未来智能层提前引入不必要依赖。

---

## 9. 对 MVP0.2 的单句定义

> **MVP0.2 的任务，不是让 PSCO 拥有更多概念，而是让它真正成为个人软件公司的每周 operating console，并在此过程中补上数据主权（导出 / 备份）这块必须闭合的底座。**

换一种更产品化的表达：

> **从"我记录了什么资产"推进到"我知道接下来该做什么，以及这些动作如何继续沉淀为资产"。**

---

## 10. 最终结论

综合第一、第二项任务，我的最终判断如下：

1. **PSCO 的长期价值与最终形态已经清晰**：它要成为个人软件公司的经营基础设施，核心是 Build / Accumulate / Compound 三条理念，最终形态是具备经营 / 产品 / 工程资产 / 上下文增强四重视角的 Command Center。

2. **MVP0.1 真实证明了最小资产主线成立**：`phase01 ~ phase05` 已在真实前后端、数据库、proto 合同与验收闭环上交付可运行现实。

3. **但 MVP0.1 仍偏"登记与查看"，且存在一条规格合规负债**：导出 / 备份未落地，直接违背"数据所有权优先"原则；这是下一阶段必须正视的工程事实。

4. **下一阶段最符合最佳实践的方向，不是继续铺开长期 domain，而是补足 Onboarding、Review Loop、Derived Intelligence 三条主线，并贯穿闭合导出 / 备份**。

5. **若下一阶段直接扩实体、扩自动化、扩智能层，PSCO 很容易重新变成"理念强、使用弱"的系统**；只有先做成高频 operating loop，长期的 `Feature / Opportunity / Experiment / AI enhancement` 才会有稳固落点。

因此，我支持后续正式 `/plan` 以以下方向为主轴：

> **先把 PSCO 做成可持续使用的 operating system 并闭合数据主权底座，再把它扩展成更完整的 company OS。**

这会是比"继续增加对象覆盖面"更稳、更符合最佳实践、也更接近原始愿景的一步。

---

*End of PSCO-mvp02-deepseekv4flash.md*