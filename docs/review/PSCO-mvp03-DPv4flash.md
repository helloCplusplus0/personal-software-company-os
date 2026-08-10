# Personal Software Company OS

# MVP0.3 Next-Stage Direction & Outlook

**Author:** DeepSeek-V4-Flash
**Date:** 2026-08-09
**Role:** 基于原始方案（`PSCO_0.md ~ PSCO_4.md`）与 `phase01 ~ phase06` 已验收现实（含 `PSCO-mvp01-summarize-feedback.md`、`PSCO-mvp02-summarize-feedback.md` 共识），以 DeepSeek-V4-Flash 专业能力陈述 PSCO 下一步推进方向与展望
**Document Type:** `review`
**Status:** 供后续正式 `/plan` 参考，不预设正式 phase 名称

---

## 1. 文档定位与结论先行

本文要完成三件事：

1. 从原始方案再确认 PSCO 的"最佳呈现形态"（Command Center 的经营视角），并据此判断 mvp03 应触达的形态增量；
2. 基于 `phase01 ~ phase06` 真实交付（含 mvp01 / mvp02 共识落地核对）给出当前真实推进基础；
3. 以我的专业能力陈述 mvp03 的推进方向与展望。

**我的结论先写在前面：**

> **`phase01 ~ phase06` 已经把 PSCO 从"概念"推进到"能登记、能绑定、能决策、能导出备份、能看见复利的可运行系统"。但直到现在，系统仍然没有证明"用户会周期性回来使用它"。**
>
> **因此 mvp03 的第一原则是：把"一次性录入"推进为"周期性经营"。**
>
> 具体方向收敛为一个主题——**Operating Review Loop + Template-Driven Compound（经营复盘闭环 + 模板驱动的复利加速）**，并以此完成 mvp02 共识中"阶段二"的收尾：日常经营 + 复利可行动 + 真实项目验证。

这与我之前在 `PSCO-mvp02-summarize-feedback-deepseekv4flash.md` 中的立场完全一致：**先做使用与回流闭环，再让复利从"空指标"变成"真事实"**。phase06 已经兑现了"复利感知基础"，现在正是把 use 变成 habit 的时刻。

---

## 2. 第一项任务：从原始方案再确认"最佳呈现形态"与 mvp03 的形态增量

### 2.1 PSCO 的长期价值再确认

`PSCO_0.md ~ PSCO_4.md` 反复确认了 PSCO 的根本定位：

> PSCO 是个人软件公司的经营与资产系统，不是代码管理工具、不是 AI Chat 产品、不是自动扫描系统。

其差异化核心稳定为四件事：

```
Module（能力单元）+ Decision（决策留痕）+ Binding（资产锚定）+ Feedback（复利反馈）
```

三条理念：`Build / Accumulate / Compound`。
长期飞轮：`更多机会 → 更多产品 → 更多模块 → 更强能力 → 更快产品 → 更多机会`。

### 2.2 最佳呈现形态：四重视角的 Command Center

`PSCO_4.md`（Document 08）定义的最终形态是 **Personal Software Company Command Center**。它具备四重视角：

| 视角 | 核心问题 | 对应载体 |
| --- | --- | --- |
| 经营视角 | 我在做什么、为什么、当前最值得推进什么 | Dashboard / Review Loop / Venture |
| 产品视角 | 产品、价值、用户、模块、决策、发布的价值链 | Product Registry |
| 工程资产视角 | 哪些模块已沉淀、是否稳定、是否复用、值得抽象 | Module Library / Reuse / Capability |
| 上下文增强视角 | AI 在 Product/Module/Decision/Review 上下文中做总结与辅助 | AI enhancement（增强层） |

### 2.3 mvp03 应触达的"形态增量"——经营视角的落地

对照四重视角看 `phase01 ~ phase06` 的现状：

- **产品视角**：`phase04` 已交付 ✅
- **工程资产视角**：`phase02`（Module）+ `phase06`（复用感知）已交付 ✅
- **上下文增强视角**：长期后移，属 AI 增强层，非本次主线
- **经营视角**：**尚未真正落地** ❌

`phase05` 交付了 Dashboard 的"总览 + 反馈信号 + 最近活动"，`phase06` 交付了"首轮录入 + 数据主权 + 复用快照"。但 Dashboard 目前仍是"进来看看"的总览页，**还不是"今天该做什么、本周复盘结论怎么回流"的经营动作起点**。

> **mvp03 的核心形态增量，就是让"经营视角"从 Dashboard 上真正站起来——把 Command Center 的第一块（经营驾驶舱）做实。**

---

## 3. 第二项任务：phase01 ~ phase06 的真实推进基础

### 3.1 已验收并完结的推进事实

`plan.md` 明确 `phase01 ~ phase06` 均已走完 `/plan -> /spec -> 实现 -> 验收 -> 收口`：

| Phase | 交付主题 | 状态 |
| --- | --- | --- |
| phase01 | MVP 规格收敛 | completed |
| phase02 | Module Registry | completed |
| phase03 | Decision Center | completed |
| phase04 | Product Registry + Repository Binding | completed |
| phase05 | Dashboard + Feedback | completed |
| phase06 | Onboarding + Data Sovereignty + Reuse Awareness | completed |

### 3.2 两轮共识的落地核对

**mvp01 共识（`PSCO-mvp01-summarize-feedback.md`）锚定了 v0.1 目标**：软件资产登记、决策留痕、基础复用反馈；核心实体 `Product / Module / Release / Decision / Repository`（`Venture` 可选）；`Capability` 派生层不建表；`Decision` 必须进 MVP；Rust 不进 MVP；GitHub 自动化后移；先做冷启动/导入/度量/导出。

**mvp02 共识（`PSCO-mvp02-summarize-feedback.md`）把下一阶段方向冻结为**：
> 从 Asset Registry 走向 Operating System，主轴 = Onboarding + Operating Review Loop + Derived Asset Intelligence。

并按"阶段一 / 阶段二"给出候选顺序。

### 3.3 `phase06` 实际交付了什么（我核对了验收报告）

`phase06` 兑现了 mvp02 共识的**阶段一**：

1. **Onboarding Foundation**：首次引导（cold-start / 根级入口）、`in_progress` 回访继续、`fromOnboarding` 回流链、首轮成功会话真实走通（`GET /api/onboarding/state`）。
2. **Data Sovereignty（数据主权闭合）**：`Export`（`GET/POST /api/dashboard/export`，覆盖 9 类资产含绑定关系）+ `Backup`（`GET/POST /api/dashboard/backup`，含 `verified` 与 `manifest_missing / coverage_incomplete / schema_mismatch` 三类失败语义）。
3. **Reuse Awareness（复用感知基础）**：`GET /api/reuse-summary`（`module_reuse_summary` + `capability_summary`），Dashboard 与 Module/Product Detail 均展示复用数与能力分布。

`phase05` Dashboard 主线在 `phase06` 下保持兼容，无破坏性回归。

### 3.4 `mvp02` 共识"阶段二"尚未落地的内容

对照 mvp02 共识 `§5.1 必做范围`，`phase06` **未覆盖**第 4、5 组：

| 组 | 内容 | phase06 状态 |
| --- | --- | --- |
| 4 | **Operating Review Loop**（daily/weekly review、`Feedback -> Decision -> Update`、把 Dashboard 变成经营动作起点） | ❌ 未落地 |
| 5 | **模板级复用**（Module 组合快照 + 新建 Product 预填） + **真实项目 dry-run** | ❌ 未落地 |

这正是 mvp02 共识 `§6.1 阶段二` 的内容，也就是 **mvp03 的天然起点**。

### 3.5 当前真实边界与未验证假设

经过 `phase06`，系统已经能回答：

- 我的资产是什么、如何绑定、关键决策是什么？
- 首轮如何低摩擦进入、如何导出备份、复用与能力分布如何？

但系统**仍然不能自然回答**：

- 我今天应该优先推进什么？
- 我上周做了什么、结论是什么、如何回流到实体？
- 我从一个已存在的 Product 的 Module 组合出发，能否快速创建下一个 Product？
- 这套闭环在真实项目上是否真的产生复利？

**一句话：`phase01 ~ phase06` 证明了"能登记、能看见复利"，但**尚未证明"会被周期性使用、且使用能产生经营与复利收益"**。这是 mvp03 必须验证的核心假设。**

---

## 4. mvp03 的判断标准

基于以上，我给出 mvp03 的判断标准：

> **mvp03 的第一原则是"验证周期性经营闭环"，第二原则是"让复利从可见走向可行动、可加速"。**

具体三条：

1. **先证明"会回来用"**：一次完整的 daily/weekly review 是否真实成立、是否回流到既有实体，比"再增加一个总览图表"更重要。
2. **先证明"决策在经营里仍居中心"**：`Feedback -> Decision -> Update` 闭环不是新对象，而是让既有 `Decision` 在经营循环里被持续调用。
3. **先证明"复利可加速创造"**：模板级复用是 PSCO 第一个"缩短从想法到新产品路径"的能力，必须先最小化落地并验证。

同时继续坚持：**AI 永远只是增强层**，mvp03 可做 Decision capture 的最小辅助，但不做"AI 决定产品方向 / 自动扫描 / 独立 AI 工作台"。

---

## 5. 我不建议 mvp03 优先做的方向（反向清单）

1. **不建议一次性引入 `Opportunity / Venture / Feature / Experiment` 流程化。** 长期对象仍未进入时机，先验证经营闭环与复利加速。
2. **不建议把 Review Loop 做成通用任务管理系统。** 这是 mvp03 最大的工程风险——review 动作必须锚定既有实体（`Decision / Product / Module`），绝不演化为独立 Kanban / Sprint / 任务面板，否则稀释 `Decision + Asset + Reuse` 差异化。
3. **不建议做重量 AI（Decision 复用推荐引擎 / AI 决策上下文增强）。** 属于 mvp03+，先让用户手动完成闭环，再谈智能辅助。
4. **不建议做 GitHub OAuth / 全自动导入。** `phase06` 已证明手动 + 辅助录入能完成闭环，集成复杂度仍留给后续。
5. **不建议把模板级复用过早工程化为"完整模板系统"。** 边界冻结为"Module 组合快照 + 预填"，不做模板版本管理、参数化、模板 CRUD 列表。

---

## 6. mvp03 推荐方向：Operating Review Loop + Template-Driven Compound

我建议把 mvp03 的推进主题收敛为：

> **经营复盘闭环 + 模板驱动的复利加速**

收敛成三条顺序主线（与 mvp02 共识"阶段二"严格对齐，并吸收复盘复用感知基础）：

### 主线 A：Operating Review Loop（经营复盘闭环）

**目标：让 Dashboard 从"总览页"成为"经营动作起点"。**

建议能力：

1. **Daily / Weekly Review**：进入系统先看 `Current Focus`（本周最值得推进对象，基于已有数据聚合，不新增实体）；Weekly 汇总本周新增 `Decision / Module / Release`、缺口 / 停滞 / 复用进展。
2. **Review 结论回流**：允许记录 review 结论（最小文本块）并**锚定到既有实体**，而非再发明第二套对象。
3. **`Feedback -> Decision -> Update` 闭环**：从 Dashboard 的 signal / gap 出发，快速进入 `Decision`，决策后回流更新相关 `Product / Module / Release`。
4. **Action handoff 但不当任务管理器**：允许把 review 结论转成最小"后续动作"，但动作必须锚定既有实体（如"Product X 需要补充 Decision"，锚定到 Product X）。

**Why 这是第一主线**：它直接验证"用户会周期性回来使用"这一 mvp01/02 从未验证的假设，是 PSCO 从"登记系统"走向"经营系统"的真正分水岭。

### 主线 B：Template-Driven Compound（模板驱动的复利加速）

**目标：让"快速创造新产品"成为现实，兑现 mvp02 共识 §4.9 的"模板级复用"，并承接 phase06 的复用感知。**

建议能力：

1. **Module 组合快照**：从已有 `Product` 的 Module 组合导出为能力模板（`SaveModuleCompositionAsTemplate`）。
2. **新建 Product 预填**：新建 Product 时基于模板预填 Module 绑定（`ApplyModuleCompositionTemplate`），预填结果可编辑。
3. **边界冻结**：模板只是"Module 组合快照 + 预填辅助"，不是新实体，不拥有 Product，不做完整模板系统。

**Why 重要**：它是 PSCO 第一个让 `Compound` 理念"可行动"的能力——不复用经验也能快速起新盘。它把 phase06 的"看见复利"推进到"用复利加速创造"。

### 主线 C：Real-Project Dry-Run 与派生智能深化

**目标：用真实项目证明这套闭环成立，别再只停在 fixture 验收。**

建议能力：

1. **真实项目 dry-run**：优先以 `Rento-miniX` 或等价真实项目走完整流程（创建 Product → 绑定 Repository → 注册 Module → 录入 Decision → 走 review → 观察复用与能力反馈 → 用模板重建），兑现 mvp01 共识 §6.4 Phase 3 承诺。
2. **派生智能深化**：在 phase06 的 `module_reuse_summary / capability_summary` 基础上，把 mvp01 §7.1 的"模块复用率 / 决策复用率 / 资产导入耗时 / 录入摩擦"度量落地为可解释指标，用于 review 与 Dashboard。
3. **dry-run 作为独立交付物**：不混入功能实现，确保不被挤压。

**Why 关键**：`phase01 ~ phase06` 全部基于 fixture 验收；只有真实项目 dry-run 才能补上"真实使用摩擦 + 资产复利收益是否成立"的证据链。

---

## 7. 候选推进结构（不冻结正式 phase 名称）

按 mvp02 共识"阶段二"与项目交付型 phase 纪律，我给出候选推进结构（**非正式命名，留给后续 `/plan`**）：

```
候选阶段一：Operating Review Loop
  └─ daily/weekly review、review 结论回流、Feedback→Decision→Update、action handoff
        ↓
候选阶段二：Template-Driven Compound + 派生智能深化 + 真实项目 dry-run
  └─ Module 组合快照 + 新建预填、度量指标落地、Rento-miniX 真实 dry-run（独立交付物）
        ↓
mvp03 收口 → mvp04 候选：GitHub 集成 / AI 决策上下文增强 / Opportunity 等长期对象进入评估
```

**依赖关系**：
- 阶段二依赖阶段一的 review 数据（有经营数据才能验证模板与度量价值）；
- 阶段二依赖 phase06 的复用感知基础（`module_reuse_summary` 是模板与度量的输入）；
- dry-run 依赖阶段一/二都可运行（真实项目要能走完整闭环）。

---

## 8. 技术与实现对 mvp03 的约束

下一阶段在实现上，我建议继续坚持以下硬约束，避免引入第二套事实源：

1. **路由状态留在 TanStack Router**：review、onboarding、模板创建涉及 multi-step 与来源上下文保留，继续用 typed search params 与 `useNavigate({ from })`，不引入第二套导航事实源。
2. **读模型 Query-first**：review / 复用快照 / Dashboard 统一 query key 设计，前缀失效刷新，高密度 UI 用 `placeholderData` 保稳。
3. **合同层只认 `.proto`**：review / 模板 / metric 新增接口一律以 `.proto` 为单一合同源，维持 `lint` 与 `breaking` 检查。
4. **后端单服务、清晰 query / application 边界**：在既有系统上补 review 与模板操作层，不做微服务、不引新基础设施。
5. **`Capability` 继续派生化**：不建重表，复用感知与模板都基于既有 `product_modules` 等聚合，保持 Local First 可导出性。

---

## 9. 度量与验收口径建议

mvp03 正式验收至少应回答：

1. 用户能否完成一次完整 weekly review，并把结论回流到既有实体？
2. `Feedback -> Decision -> Update` 是否形成真实闭环，`Decision` 是否在经营循环中保持中心地位？
3. 模板级复用是否完成"保存组合 → 新建预填 → 继续编辑 → 完成创建"闭环？
4. 模块复用率 / 资产导入耗时等度量是否在 Dashboard 可解释观测？
5. 至少一个真实项目 dry-run 是否走通，并形成独立验收记录？
6. 是否未引入超出 `TECH_STACK_BASELINE` 的技术选择、未回退重做 phase06 已冻结交付物？

---

## 10. 最终结论

1. `phase01 ~ phase06` 已把 PSCO 从"理念"推进到可运行的资产系统：能登记、能绑定、能决策、能导出备份、能看见复用复利。
2. 但直到现在，系统**尚未证明会被周期性使用**——"经营视角"（Command Center 的第一块）仍未真正落地。
3. 因此 mvp03 的第一原则是**把"一次性录入"推进为"周期性经营"**，方向收敛为 **Operating Review Loop + Template-Driven Compound**，并以此完成 mvp02 共识"阶段二"的收尾。
4. 我坚持此前立场：**先做使用与回流闭环，复利才会从"空指标"变成"真事实"**；mvp03 是先补"经营闭环"，再让"复利可行动、可加速"。
5. 最大的工程风险是 Review Loop 失控成任务管理器——必须让 review 动作锚定既有实体，绝不另起第二套对象。

用一句话收口 mvp03 的展望：

> **phase01~06 证明了"能登记、能看见复利"；mvp03 要证明"会被周期性经营，且经营能加速复利"。**

这是 PSCO 从"资产登记系统"走向"个人软件公司 operating system"的最后一段关键闭环。

---

*End of PSCO-mvp03-DPv4flash.md*