# Personal Software Company OS

# MVP0.2 Next-Stage Direction Plan

**Author:** DeepSeek-V4-Pro
**Date:** 2026-08-09
**Role:** 基于原始方案（`PSCO_0.md ~ PSCO_4.md`）与 MVP0.1（`phase01 ~ phase05`）已验收现实，系统输出 PSCO 下一阶段推进方向的标准计划文档，供后续正式 `/plan` 参考评价
**Document Type:** `review`
**Status:** 供后续 `/plan` 参考，不直接构成正式 phase 命名或执行指令

---

## 0. 说明：本文与其他评审文档的关系

在撰写本文之前，已有 GPT54、DeepSeek-V4-Flash、GLM-5.2 三份独立的 MVP0.2 方向评审文档输出到 `docs/review/`。本文与它们的关系如下：

1. **本文独立完成**，基于同一组原始方案与已验收交付物形成判断，但会标注与其他评审的共识与差异；
2. **本文与 GPT54 / DeepSeek-V4-Flash 的评审在三线结构（Onboarding / Review Loop / Derived Intelligence）上高度收敛**，这不是"复制"，而是三份独立评审基于同一组事实得出的相似结论，本身就构成对这三条主线优先级的强验证；
3. **本文的独特价值**在于：作为 DeepSeek-V4-Pro，我曾在 2026-08-04 对 PSCO 原始方案做过一次独立第三方方案评审（`PSCO-evaluation-deepseek-v4-pro.md`），本文会系统性地回看那次评审中的 9 条建议在 MVP0.1 中被如何处理、哪些仍然悬空、哪些已经成为新的推进约束；
4. 本文不重复 GPT54 或 DeepSeek-V4-Flash 已详细展开的内容，而是聚焦于**补充视角、强化论证、给出工程级落点建议**。

---

## 1. 文档定位与结论先行

本文不重复总结 `phase01 ~ phase05` 的交付清单，也不替代下一阶段的正式 `/plan`。

本文要完成三件事：

1. 重新确认 PSCO 的长期核心价值与最终呈现形态；
2. 基于 MVP0.1 的真实工程现实，结合我 2026-08-04 方案评审的 9 条建议回看，给出下一阶段最符合最佳实践的推进方向；
3. 输出一份可供后续 `/plan` 参考的标准方向文档，并明确与前序评审的共识与差异。

**我的结论先写在前面：**

> **MVP0.2 的第一原则，不是继续扩展对象宽度，而是把 PSCO 从"可登记、可查看的资产系统"推进为"可持续使用、可回流、可复利的经营运行系统"。**

更具体地，我建议下一阶段收敛为一个单主题 phase，拆成三条顺序主线：

1. **Cold Start / Onboarding Foundation**（让已有现实资产低摩擦进入系统，并闭合数据主权底座）
2. **Review / Operating Loop**（把 Dashboard 变成 daily / weekly 经营起点，让 review 结论回流到实体）
3. **Derived Asset Intelligence**（让能力增长成为可观测事实，而非叙事）

同时，作为 DeepSeek-V4-Pro，我会额外强调一个在前序评审中已被 DeepSeek-V4-Flash 指出、但值得进一步展开的工程现实：

> **在规划新方向之前，必须先正视并闭合 MVP0.1 遗留的规格合规负债（导出 / 备份），否则新的 operating layer 会叠加在未闭合的数据主权底座之上。**

---

## 2. 第一项任务：PSCO 核心价值、关键功能与最终形态（从原始方案系统再确认）

### 2.1 一句话定位

`PSCO_0.md`（Document 00）给出了 PSCO 的根本定位：

> 让一个个人开发者具备一家软件公司的组织能力，使其能够持续发现机会、验证需求、构建产品、积累软件能力，并形成长期复利资产。

它的核心判断非常关键：

> 软件开发的核心竞争力，正在从"编码能力"转向"组织能力 + 判断能力 + 资产积累能力"。

AI 正在快速降低编码成本、学习成本、原型成本和技术实现成本，但它不会自动产生商业判断、产品洞察、长期经验和可复用软件能力。因此，PSCO 不是"更快写代码"的工具，而是"个人软件公司的运行系统"。

### 2.2 三大核心理念（Build / Accumulate / Compound）

`PSCO_0.md`（Document 00/01）用三个理念定义了 PSCO 的价值内核：

| 理念 | 含义 | 关键动作 |
| --- | --- | --- |
| **Build（持续创造）** | 个人开发者不是一次性项目开发者，而是像公司一样持续寻找机会、构建产品 | `Opportunity → Product → Software` |
| **Accumulate（持续积累）** | 每次开发都应留下资产，从"项目归档"升级为"提取能力 → 形成模块 → 进入资产库" | `项目 → 模块 → 资产库` |
| **Compound（持续复利）** | 长期价值来自能力复用，模块随版本演进形成能力，被未来产品反复复用 | `Module v1 → v2 → Capability → 复用` |

这三点直接决定了 PSCO 的差异化：它管理的是"我形成了多少能力"，而不是"我完成了多少任务"。

### 2.3 关键功能与长期对象链

`PSCO_1.md`（Document 02/03）给出了完整的长期运行模型与领域模型：

```
Opportunity → Venture → Product → Feature → Module → Release → Capability
辅助实体：Decision / Experiment / Repository
```

`PSCO_2.md`（Document 04/05）进一步把 PSCO 的核心差异化定位在 **Module System**：

> Module = Capability Boundary + Implementation + Interface + Knowledge + History

它强调 Module 不是代码包（如 npm package），而是"软件公司资产"，应按能力（Foundation / Application / Domain / AI / Data）而非语言分类，经历 `Prototype → Candidate → Internal → Stable → Commercial` 生命周期。

`PSCO_3.md`（Document 06/07）定义了两件事：

- **工作流引擎**：`Think → Validate → Build → Extract → Reuse → Compound`，让 PSCO 成为"驱动个人软件公司持续运行的工作流系统"而非"记录工具"；
- **AI 边界**：AI 永远是增强层，不是决策层；绝对不做自动代码扫描、自动知识图谱、AI 自动判断方案。

`PSCO_4.md`（Document 08）定义了最终 UX 形态：

> PSCO 不是 Jira / Notion / GitHub / AI Chat，而是 **Personal Software Company Command Center**（个人软件公司控制台/驾驶舱）。

它的体验目标不是"让用户管理更多东西"，而是"让用户更高效地经营自己的软件公司"。

### 2.4 最终呈现形态：四重视角的 Command Center

综合 `PSCO_0 ~ PSCO_4`，PSCO 的最终形态是一个具备四重视角的个人软件公司运行系统：

1. **经营视角**：知道自己在做什么、为什么做、当前最值得推进什么（Dashboard / Current Focus / Venture）；
2. **产品视角**：把产品、价值主张、用户、模块、决策、发布串成一条可追溯的价值链（Product Registry）；
3. **工程资产视角**：看见哪些模块已沉淀、是否稳定、是否被复用、是否值得进一步抽象（Module Library / Capability）；
4. **上下文增强视角**：AI 在 `Product / Module / Decision / Review` 上下文中做总结、提醒、候选与文档辅助，但不抢走判断权。

一句话定义最终形态：

> **让一个独立开发者把经营、产品、工程与资产积累放进同一套可长期复利的运行系统里。**

### 2.5 我 2026-08-04 初次评审时对长期愿景的再确认

在 `PSCO-evaluation-deepseek-v4-pro.md` 中，我对原始方案给出的总体评价是 7.3/10，核心判断是：

> "PSCO OS 是一份具有原创性洞察和清晰愿景的优秀方案设计。……如果能在 v0.1 中解决冷启动、数据可移植性和 GitHub 集成这三个关键问题，它将从一个'好的设计文档'变成'一个真正可用的产品'。"

现在回头看，这个判断中有相当一部分已经在 MVP0.1 中得到验证或修正，但仍有部分核心问题悬空。详见第 4 节。

---

## 3. 第二项任务：MVP0.1 的真实推进基础（从仓库现实再确认）

### 3.1 已验收并完结的推进事实

`plan.md` 明确 `phase01 ~ phase05` 均已走完 `/plan → /spec → 实现 → 验收 → 收口`：

| Phase | 交付主题 | 唯一规格收敛入口 | 状态 |
| --- | --- | --- | --- |
| phase01 | MVP 规格收敛 | `mvp_spec_v0.1.md` | ✅ completed |
| phase02 | Module Registry | `module_registry_spec_v0.1.md` | ✅ completed |
| phase03 | Decision Center | `decision_center_spec_v0.1.md` | ✅ completed |
| phase04 | Product Registry + Repository Binding | `product_repository_binding_spec_v0.1.md` | ✅ completed |
| phase05 | Dashboard + Feedback | `dashboard_feedback_spec_v0.1.md` | ✅ completed |

### 3.2 真实工程落地的骨架

我实际核对了仓库代码结构，确认以下事实：

- **前端（`frontend/src`）**：基于 `React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind + shadcn/ui`，已形成 `features/dashboard`、`features/module-registry`、`features/product-registry`、`features/decision-center`、`features/repository-binding` 五个 feature 模块，以及对应的 `routes/*` 与 `routeTree.gen.ts`。

- **后端（`backend`）**：基于 Go 模块化单体，`internal/` 下已形成 `moduleregistry`、`decisioncenter`、`productregistry`、`repositorybinding`、`dashboard`、`platform` 六个领域包，采用 `handler / service / repository / candidate / errors / validate / types / response` 的分层职责。

- **合同（`proto`）**：已建立 `module_registry`、`decision_center`、`product_registry`、`repository_binding`、`dashboard`、`common` 六个 `.proto` 包，并对应生成 Go 与 TS 代码，`proto/Makefile` 含 `build / gen / lint / breaking / clean` 五个受控 target。

- **数据（`database`）**：已有 6 张 migration、多组基线 seed 与 reset seed，以及 `init_db.sh / run_seeds.sh / reset_*_acceptance.sh` 等受控脚本，支持"先清空再加载指定 fixture"的可重复验收语义。

- **验收实证**：`phase05-14` 已证明 Dashboard 的 `overview / feedback-signals / recent-activities` 三类聚合读在真实前后端与数据库上跑通，空状态、有数据、局部错误、跳转返回路径均形成可复验矩阵。

### 3.3 MVP0.1 已经被证明成立的事实

1. **最小资产主线真实成立**：Module / Release / Decision / Product / Repository / Dashboard 不再是概念图，而是可运行、可复验、可交付的现实。
2. **单一合同源与单一事实链路成立**：`.proto → HTTP envelope → 前端 adapter → 真实联调验收` 已经形成稳定一致的工程方式，不存在第二套 JSON 合同或第二套路由事实源。
3. **Dashboard 已具备"驾驶舱雏形"**：聚合读、局部错误隔离、跳转/返回上下文保留、Feedback 作为最小反馈层均已落地。
4. **工程纪律得到验证**：`spec → 实现 → 子代理复核 → 验收 → 收口` 的流程门禁在五个 phase 中持续生效。

### 3.4 MVP0.1 的真实边界

尽管 MVP0.1 成立，但以下四点现状决定了下一阶段的起点：

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

#### 3.4.3 存在规格合规负债（导出 / 备份未落地）

`mvp_spec_v0.1.md` §7.3 / §7.4 明确要求：
- 导出：面向用户带走核心资产数据，范围至少覆盖 `Product / Module / Release / Repository / Decision` 及基础绑定关系；
- 备份：至少提供一种面向当前实例的基础备份路径，不依赖 GitHub 或第三方平台作为唯一前提。

但我在后端代码中**未检索到任何 export / backup / dump 对应的实现或接口**。这意味着：
- 这是一条"规格已要求、实现未落地"的合规欠账；
- 它对"数据所有权优先（Local First）"这一核心原则是直接违背的；
- 若在未闭合此负债的情况下直接铺开 operating layer，会让新功能叠加在未闭合的数据主权底座上。

> 我的判断：**下一阶段 `/plan` 在规划 operating loop 之前或之内，必须显式处理这条负债，而不是继续把它当作"后续再说"。** 这一点与 DeepSeek-V4-Flash 的评审结论完全一致。

#### 3.4.4 workflow 仍停留在"项目治理层"，未转化为"产品经营层"

`PSCO_3.md` 把"工作流"当成系统的一等问题。但在 MVP0.1 里，workflow 更多体现为规格推进、开发验收、文档治理这一类"项目级工作流"，尚未转化为最终用户在产品内执行的 daily / weekly operating loop。

---

## 4. 回看：我 2026-08-04 方案评审的 9 条建议在 MVP0.1 后的状态

这是本文的独特视角。我在 `PSCO-evaluation-deepseek-v4-pro.md` 中提出了 9 条具体改进建议。现在逐一回看它们在 MVP0.1 后的状态：

| # | 建议 | 优先级 | MVP0.1 后状态 | 对 MVP0.2 的影响 |
| --- | --- | --- | --- | --- |
| 1 | MVP 领域模型从 9 个实体缩减为 4-5 个 | 🔴 高 | ✅ **已采纳**。MVP0.1 实际执行了 `Product / Module / Release / Decision / Repository`，Venture 为可选，Feature/Opportunity/Experiment 不进入 | 无需再讨论，已收敛 |
| 2 | 冷启动体验设计（Onboarding Flow） | 🔴 高 | ⚠️ **未落地**。当前系统仍依赖理解多对象后手工录入 | **这是 MVP0.2 最优先需补的缺口** |
| 3 | GitHub 集成策略（OAuth + 仓库导入） | 🔴 高 | ⚠️ **未落地**。原始方案已明确"GitHub OAuth / 自动导入不作为当前阶段阻断项"，这在本阶段是正确决策，但长期看仍是采纳障碍 | 建议在 MVP0.2 中评估"辅助导入"而非"全自动导入"，先验证低摩擦录入的价值 |
| 4 | 数据可移植性设计（导出 / 备份） | 🔴 高 | ❌ **未落地**。规格已要求但实现缺失，构成真实合规负债 | **MVP0.2 必须闭合** |
| 5 | Capability 实体澄清（v0.1 作为属性字段，v0.2+ 渐进引入） | 🟡 中 | ✅ **已采纳**。`Capability` 在 MVP0.1 中作为派生结果层 | 建议在 MVP0.2 的 Derived Intelligence 阶段继续坚持派生优先 |
| 6 | 决策记录输入简化（Title-first 模式） | 🟡 中 | ⚠️ **部分采纳**。Decision 已有最小模板，但尚未实现 title-first 渐进补全 | 可在 MVP0.2 的 Review Loop 中优化 |
| 7 | 成功度量标准定义 | 🟡 中 | ⚠️ **未落地**。Dashboard 已有反馈信号，但缺乏面向用户的"能力增长"可解释指标 | 在 MVP0.2 的 Derived Intelligence 中自然承接 |
| 8 | 架构图修正（Rust Engine 标注为 Phase 2） | 🟢 低 | ✅ **已采纳**。`PSCO_2.md` 已明确标注为远期方向，`TECH_STACK_BASELINE.md` 已冻结技术路线 | 无需再处理 |
| 9 | 文档结构优化（Glossary 作为术语唯一来源） | 🟢 低 | ⚠️ **未落地**。但文档职责已通过 `project_rules.md` 和 `AGENTS.md` 去重收敛 | 非当前阻断项 |

**关键发现：** 我 2026-08-04 提出的 4 条 🔴 高优先级建议中，有 2 条（冷启动、数据可移植性）在 MVP0.1 结束时仍未落地，1 条（GitHub 集成）是正确的延迟决策，1 条（领域模型简化）已被采纳。这意味着 MVP0.2 面临着与我 8 月 4 日评审时**高度一致**的未闭合问题，只是现在这些问题出现在一个更成熟的工程基础之上。

---

## 5. 下一阶段的判断标准

基于以上分析，我给出与项目长期愿景一致的判断标准：

> **下一阶段的第一原则是"先提高使用频率与回流闭环，再扩展对象宽度"，而不是反过来。**

具体三条：

1. **先让事实可进入、可回顾、可转动作，再做更抽象的分析层。**
   → 顺序应为：导入 → review → derived intelligence → 再考虑长期 domain 宽度。

2. **先让"复利"可被观测，再决定是否升级为正式模型。**
   → `Capability` 继续作为派生层，先做 derived signals，再决定何时升级为更正式的能力模型。

3. **继续坚持 AI 是增强层。**
   → 下一阶段可以增加 AI 辅助（周报总结、决策草稿、导入字段建议、模块复用提示、反馈解释），但绝不变成"AI 决定产品方向 / 自动扫描全仓库 / 独立一级导航 AI Assistant"。

---

## 6. 我不建议下一阶段优先做的方向（反向清单）

这里与 GPT54 和 DeepSeek-V4-Flash 的判断完全一致，我用自己的语言重新确认：

1. **不建议一次性全面引入 `Opportunity / Venture / Feature / Experiment`。**
   这些对象在长期模型中都成立，但现在一起进入会把系统拖回"概念完整优先"。更稳的做法是：先在 review / analysis 层讨论进入条件，等 operating loop 验证后再挑最必要的一个进入。

2. **不建议把 PSCO 做成通用项目管理系统。**
   PSCO 可以承接 operating action，但不应演化为 Kanban / Sprint / 通用任务平台，否则会稀释 `Decision + Asset + Reuse` 的差异化价值。

3. **不建议把自动代码扫描、知识图谱、Rust 智能层作为下一阶段主线。**
   原始方案已明确这些不是当前主线；在数据密度、导入质量、日常使用频率未建立之前，直接做智能层只会提高工程成本、降低验证效率。

4. **不建议先做重量级 GitHub OAuth / 全自动导入。**
   下一阶段更需要验证的是"辅助导入有没有价值、用户愿不愿意持续维护事实、review 是否被真实使用"，而不是先解决最大自动化问题。

---

## 7. 推荐的下一阶段主题与主线

我建议把 MVP0.2 的推进主题定义为：

> **从 Asset Registry 走向 Operating System**

收敛成三条顺序主线（与 GPT54 和 DeepSeek-V4-Flash 的评审在三线结构上高度一致，但我会从工程落地角度给出更具体的分解与交叉依赖）：

### 主线 A：Cold Start / Onboarding Foundation

**目标不是"自动扫描所有代码"，而是：**

> 让一个已有现实项目的用户，能在 30-60 分钟内把第一批核心资产放进 PSCO，并获得初始价值。

**建议能力：**

1. **First-run onboarding**
   - 以一条受控流程引导用户先建立首个 `Product`
   - 再补首批 `Repository / Module`
   - 再补关键 `Decision`
   - 关键是"引导式、最低字段、能跳过"三步

2. **Draft-first / partial-entry**
   - 允许对象以最少字段先创建（如 Product 只需 name + description）
   - 后续逐步补全，而不是首次录入就要求完整结构
   - 在 Dashboard / List 中以"待补全"标签提示

3. **Controlled import helper**
   - 支持从已有仓库 URL、README 摘要、人工表单中导入基础事实
   - 重点是"辅助录入"（如预填字段建议），不是"自动理解一切"
   - 先做最小版本：输入仓库 URL → 拉取 README 摘要 → 预填 Product 名称、描述

4. **Low-friction decision capture**
   - `Decision` 采用 ADR-lite
   - 支持 title-first、context-later 的渐进补全
   - 在 Dashboard 和 Review 中提示"待补全的决策"

**DoD 建议：**
- 新用户可在一次会话内完成首个 `Product + Repository + Module + Decision` 录入
- 首次进入可在不查文档的前提下理解系统基本路径
- 对象允许最小字段创建并逐步补全

### 贯穿主线：Export / Backup 合规闭合

**这是一个贯穿 MVP0.2 的独立工作项，不应被归入任何单一主线：**

> 补全 `mvp_spec_v0.1.md` §7.3 / §7.4 已要求但未落地的导出 / 备份能力。

**理由：**
- 这是"数据所有权优先（Local First）"的硬约束，是 PSCO 差异化价值的底座；
- 现有代码无对应实现，属于真实规格负债；
- 导出能力同时是 Onboarding（把已有数据带进系统）与 Review（把经营结果带走）的天然对称能力；
- 建议在 MVP0.2 的早期阶段（Block 1 内）闭合，不让它成为后续的延迟债务。

**建议最小实现：**
- 导出：提供一键导出全部核心数据为 JSON 的功能（覆盖 `Product / Module / Release / Repository / Decision` 及绑定关系）
- 备份：提供基于 `pg_dump` 的数据库备份脚本，支持本地文件存储
- 不追求"完美迁移方案"，只追求"用户可以带走数据"

### 主线 B：Review / Operating Loop

**这是我认为最关键的一条：**

> 验证用户能否把 Dashboard 作为 daily / weekly operating 起点来使用。

**建议能力：**

1. **Daily Review（每日概览）**
   - 进入系统后先看 `Current Focus`（本周最值得推进的对象）
   - 一键跳到相关 `Product / Module / Decision`
   - 不新增实体，基于已有数据聚合生成

2. **Weekly Review（周复盘）**
   - 汇总本周新增 `Decision`、新增/变化的 `Module / Release`、缺口/停滞/复用进展
   - 允许记录 review 结论（最小文本块），并回流到现有实体
   - 关键：review 结论必须锚定到既有实体，而非再发明第二套对象

3. **Action handoff without task-manager drift**
   - 不把 PSCO 做成完整任务管理工具
   - 但允许把 review 结论转成最小"后续动作"记录
   - 这些动作必须锚定到既有实体（如 "Product X 需要补充 Decision" 锚定到 Product X）

4. **Feedback → Decision → Update 闭环**
   - 从 Dashboard 的 gap 或 signal 出发
   - 能快速进入 `Decision`
   - 决策后回流更新相关 `Product / Module / Release`

**DoD 建议：**
- 用户可完成一次完整 weekly review
- review 结果能回流到既有实体
- Dashboard 能承接"看见问题 → 做出动作 → 记录决策"的主线

### 主线 C：Derived Asset Intelligence

**开始验证"复利"不是叙事，而是系统可观测事实。**

**建议优先做派生指标，而非完整 `Capability` 实体：**

1. `module_reuse_count`（模块被多少个 Product 使用）
2. `cross_product_usage_count`（跨产品模块使用数）
3. `release_freshness`（最近一次 Release 距今时间）
4. `decision_link_density`（每个实体的关联决策密度）
5. `orphan_repository / orphan_module / orphan_product`（孤岛实体数量）
6. "候选能力"提示（如"3 个模块都涉及 Auth，是否考虑形成 Authentication Capability？"）

**原则：**
> 先做 derived signals，再决定何时把其中一部分升级为更正式的能力模型。

**DoD 建议：**
- Dashboard 或 detail 页面可展示最小复利指标
- 至少有一组指标可解释"能力增长"或"复用增长"
- 不引入手工维护的重 `Capability` 写入流程

---

## 8. 对下一阶段 `/plan` 的建议形状

### 8.1 建议先开一个"单主题 phase"，不要一次并行开多个大方向

推荐优先主题（评审建议名，非正式 phase 名）：

> **Operating Loop Foundation**

它应直接承接：
- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
- `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/`
- `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`

### 8.2 建议的三块顺序子阶段与依赖关系

这里我给出比 GPT54 和 DeepSeek-V4-Flash 更细粒度的依赖分析：

```
Block 1: Cold Start / Onboarding + Export/Backup 闭合
    │
    │ 依赖：Dashboard 已有跳转路径（phase05 已交付）
    │ 产出：Onboarding 流程、partial-entry 支持、辅助导入、导出/备份
    │
    ▼
Block 2: Review / Operating Loop
    │
    │ 依赖：Block 1 的 Onboarding 使系统中有足够数据做 review
    │ 依赖：Block 1 的导出/备份使用户信任系统的数据主权
    │ 产出：Daily Review、Weekly Review、Action handoff、Feedback→Decision 闭环
    │
    ▼
Block 3: Derived Asset Intelligence
    │
    │ 依赖：Block 2 的 review 数据使派生指标有足够的观测基础
    │ 产出：最小复利指标、候选能力提示
```

**关键依赖说明：**
- Block 2 依赖 Block 1 的 Onboarding（因为有数据才能 review）
- Block 3 依赖 Block 2 的 Review（因为 review 数据是派生指标的重要输入）
- Export/Backup 应在 Block 1 内闭合，因为它是数据主权底座，不宜延迟

### 8.3 各 Block 的 DoD 汇总

#### Block 1：Cold Start / Onboarding + Export/Backup

- 新用户可在一次会话内完成首个 `Product + Repository + Module + Decision` 录入
- 首次进入可在不查文档前提下理解系统基本路径
- 对象允许最小字段创建并逐步补全
- 导出 / 备份基础路径已闭合（JSON 导出 + pg_dump 备份脚本）

#### Block 2：Review Loop

- 用户可完成一次完整 weekly review
- review 结果能回流到既有实体
- Dashboard 能承接"看见问题 → 做出动作"的主线

#### Block 3：Derived Intelligence

- Dashboard 或 detail 页面可展示最小复利指标
- 至少有一组指标可解释"能力增长 / 复用增长"
- 不引入手工维护的重 `Capability` 写入流程

---

## 9. 技术与实现层最佳实践建议

下一阶段在实现上，我建议继续坚持以下约束，避免引入第二套事实源。这些与 GPT54 和 DeepSeek-V4-Flash 的评审一致。

### 9.1 路由状态继续留在 TanStack Router

基于当前仓库经验与最新 Router 文档，下一阶段涉及 onboarding、review、multi-step return path 时，应继续：
- 使用 typed search params
- 保持 `useNavigate({ from })`
- 在需要跨层保留来源上下文时优先使用 `retainSearchParams`

不要再引入第二套全局导航事实源。

### 9.2 Dashboard / Review 读模型继续采用 Query-first

基于当前 Query 文档，建议继续：
- 统一 query key 设计
- 使用前缀失效刷新相关聚合读
- 对 review 或分页列表使用 `placeholderData` 保持高密度 UI 稳定

这样能保持页面密度与切换流畅性，而不是在每次刷新时回退到大面积 loading。

### 9.3 合同层继续只认 `.proto`

基于当前 Buf 实践，下一阶段新增 review / onboarding / derived metrics 接口时，应继续：
- 以 `.proto` 为单一合同源
- 维持 `lint` 与 `breaking` 检查
- 不在页面层偷偷长出第二套 contract 语义

### 9.4 后端继续坚持单服务、清晰 query / application 边界

下一阶段更像是在已有系统上补 operating layer，而不是重构成新架构。因此应继续避免：
- 微服务拆分
- 新的复杂基础设施
- 为未来智能层提前引入不必要依赖

### 9.5 导出/备份的技术路径建议

这是 MVP0.2 特有的新增建议：

- **导出**：新增一个后端端点（如 `GET /api/export/data`），返回 JSON，包含所有核心实体的完整数据。前端提供一键下载按钮（放在 Settings 或 Dashboard 中）。
- **备份**：提供 `database/backup.sh` 脚本，封装 `pg_dump`，支持指定输出路径。不追求自动定时备份，先保证"可以备份"。
- **合同**：导出/备份的端点应在 `.proto` 中定义，维持单一合同源。

---

## 10. 与其他评审文档的共识与差异

### 10.1 共识点（三份独立评审的收敛）

| 共识点 | GPT54 | DS-V4-Flash | DS-V4-Pro（本文） |
| --- | --- | --- | --- |
| 下一阶段不应扩展对象宽度，应先提高使用频率 | ✅ | ✅ | ✅ |
| 三线结构：Onboarding / Review Loop / Derived Intelligence | ✅ | ✅ | ✅ |
| 不建议全面引入长期实体 | ✅ | ✅ | ✅ |
| 先做 derived signals，不做重 Capability 实体 | ✅ | ✅ | ✅ |
| 继续坚持 AI 是增强层 | ✅ | ✅ | ✅ |
| 技术栈继续冻结，不引入新基础设施 | ✅ | ✅ | ✅ |

### 10.2 差异点（本文的独有补充）

| 差异点 | 本文立场 |
| --- | --- |
| 2026-08-04 方案评审回看 | 本文独有：系统回看了 9 条建议的落地状态，发现 2 条 🔴 高优先级建议（冷启动、数据可移植性）在 MVP0.1 后仍未闭合 |
| 导出/备份负债的工程级论证 | 与 DS-V4-Flash 一致，但本文从"方案评审回看"角度给出了更强的论证链路 |
| 三块子阶段的依赖关系 | 本文给出了更明确的顺序依赖：Block 2 依赖 Block 1（有数据才能 review），Block 3 依赖 Block 2（review 数据是派生指标的输入） |
| 导出/备份的技术路径 | 本文给出了具体的技术实现建议（JSON 导出端点 + pg_dump 备份脚本） |

---

## 11. 对 MVP0.2 的单句定义

> **MVP0.2 的任务，不是让 PSCO 拥有更多概念，而是让它真正成为个人软件公司的每周 operating console，并在此过程中补上数据主权（导出 / 备份）这块必须闭合的底座。**

换一种更产品化的表达：

> **从"我记录了什么资产"推进到"我知道接下来该做什么，以及这些动作如何继续沉淀为资产"。**

---

## 12. 最终结论

综合第一、第二项任务，以及我 2026-08-04 方案评审的回看，我的最终判断如下：

1. **PSCO 的长期价值与最终形态已经清晰**：它要成为个人软件公司的经营基础设施，核心是 Build / Accumulate / Compound 三条理念，最终形态是具备经营 / 产品 / 工程资产 / 上下文增强四重视角的 Command Center。

2. **MVP0.1 真实证明了最小资产主线成立**：`phase01 ~ phase05` 已在真实前后端、数据库、proto 合同与验收闭环上交付可运行现实。这意味着 MVP0.2 不需要"重新证明概念"，而是需要"让概念变成可日常使用的产品"。

3. **MVP0.1 仍偏"登记与查看"，且存在一条规格合规负债**：导出 / 备份未落地，直接违背"数据所有权优先"原则；冷启动和导入摩擦未解决；workflow 尚未从"项目治理层"转化为"产品经营层"。

4. **我 2026-08-04 方案评审的 4 条 🔴 高优先级建议中，有 2 条（冷启动、数据可移植性）在 MVP0.1 后仍未闭合**。这不是偶然，而是说明这些问题是 PSCO 从"设计文档"变成"可日常使用产品"过程中必须跨越的坎。

5. **下一阶段最符合最佳实践的方向，不是继续铺开长期 domain，而是补足 Onboarding、Review Loop、Derived Intelligence 三条主线，并贯穿闭合导出 / 备份**。这个方向与 GPT54 和 DeepSeek-V4-Flash 的独立评审高度收敛，三份独立评审的一致结论本身就是对这三条主线优先级的最强验证。

6. **若下一阶段直接扩实体、扩自动化、扩智能层，PSCO 很容易重新变成"理念强、使用弱"的系统**；只有先做成高频 operating loop 并闭合数据主权底座，长期的 `Feature / Opportunity / Experiment / AI enhancement` 才会有稳固的工程和心理落点。

因此，我支持后续正式 `/plan` 以以下方向为主轴：

> **先把 PSCO 做成可持续使用的 operating system 并闭合数据主权底座，再把它扩展成更完整的 company OS。**

这会是比"继续增加对象覆盖面"更稳、更符合最佳实践、也更接近原始愿景的一步。

---

*End of PSCO-mvp02-deepseekv4pro.md*