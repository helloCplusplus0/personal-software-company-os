# Personal Software Company OS

# MVP0.2 推进方向 — 五方评审汇总与仲裁建议

**Author:** DeepSeek-V4-Pro
**Date:** 2026-08-09
**Role:** 基于 GPT54、DeepSeek-V4-Flash、DeepSeek-V4-Pro、GLM-5.2、Qwen3.7-Pro 五份独立评审，进行交叉汇总、共识提取、分歧仲裁与方向建议
**Document Type:** `review`
**Status:** 供后续正式 `/plan` 参考，不直接构成正式 phase 命名或执行指令

---

## 0. 文档定位

本文是对五份独立 MVP0.2 方向评审的系统性汇总与仲裁。五份评审均已输出到 `docs/review/`：

| 评审文档 | 作者 | 核心主张 |
| --- | --- | --- |
| `PSCO-mvp02-GPT54.md` | GPT54 | **从 Asset Registry 到 Operating System**：Onboarding → Review Loop → Derived Intelligence |
| `PSCO-mvp02-deepseekv4flash.md` | DeepSeek-V4-Flash | **从 Asset Registry 到 Operating System**：Onboarding → Review Loop → Derived Intelligence，并强调导出/备份负债 |
| `PSCO-mvp02-deepseekv4pro.md` | DeepSeek-V4-Pro（本文作者） | **从 Asset Registry 到 Operating System**：Onboarding+Export/Backup → Review Loop → Derived Intelligence，回看 8/4 方案评审的 9 条建议 |
| `PSCO-mvp02-GLM52.md` | GLM-5.2 | **不扩实体，先做复利**：复用感知深化（方向 A）→ 模板复用+导出（方向 B 部分） |
| `PSCO-mvp02-qwen37pro.md` | Qwen3.7-Pro | **不扩实体，先做复利**：复用感知+能力派生（方向 A）→ 模板复用+度量+导出（方向 B+C） |

---

## 1. 强共识（五方一致，无需再讨论）

以下结论在五份评审中**完全一致**，可直接作为下一阶段 `/plan` 的不可谈判边界：

### 1.1 不扩展实体范围

| 共识项 | 五方立场 |
| --- | --- |
| 不引入 Opportunity / Feature / Experiment 流程化 | ✅ 五方一致 |
| 不强制 Venture 实现（继续可选） | ✅ 五方一致 |
| Capability 继续作为派生层，不建重实体 | ✅ 五方一致 |
| 不新增核心实体 CRUD | ✅ 五方一致（GLM/Qwen 明确写入"不新增核心实体"） |

### 1.2 不做的事

| 共识项 | 五方立场 |
| --- | --- |
| 不做自动代码扫描 / 知识图谱 | ✅ 五方一致 |
| 不做 Rust Intelligence Layer | ✅ 五方一致 |
| 不做 AI Assistant 一级主导航 | ✅ 五方一致 |
| 不把 PSCO 做成通用项目管理工具 | ✅ 五方一致 |
| 不做重量级 GitHub OAuth / 全自动导入 | ✅ 五方一致 |
| 不引入第二套路由 / 合同 / ORM / UI 框架 | ✅ 五方一致 |

### 1.3 必须做的事

| 共识项 | 五方立场 |
| --- | --- |
| 数据导出 / 备份必须在 MVP0.2 闭合 | ✅ 五方一致（GPT54 在正文提及，其余四方显式强调） |
| module_reuse_summary 派生视图必须落地 | ✅ 五方一致（GPT54 归入 Derived Intelligence，GLM/Qwen 归入复用感知） |
| capability_summary 派生视图必须落地 | ✅ 五方一致 |
| 技术栈继续冻结，不引入新基础设施 | ✅ 五方一致 |

> **结论：以上 3 组共 15 项已形成不可谈判的强共识，下一阶段 `/plan` 不得突破。**

---

## 2. 核心分歧：两条路线的差异

五份评审虽然大方向一致，但在 **"MVP0.2 的第一优先级是什么"** 上存在一条清晰的分界线：

### 路线一："经营闭环优先"（GPT54 / DS-V4-Flash / DS-V4-Pro）

**核心主张：先让系统变得可日常使用，再让复利可见。**

```
Block 1: Cold Start / Onboarding（降低进入门槛）
    ↓
Block 2: Review / Operating Loop（把 Dashboard 变成日常经营起点）
    ↓
Block 3: Derived Asset Intelligence（让复利可观测）
```

**逻辑链：**
- 如果用户不能低摩擦地把数据放进系统，就没有数据可做复用感知
- 如果用户不能把 Dashboard 当作日常经营起点，复用感知只是另一个看板 widget
- 因此：先解决"进得来 + 用得上"，再解决"看得见复利"

**额外强调：** DS-V4-Flash 和 DS-V4-Pro 都将导出/备份作为 Block 1 的贯穿性工作项，认为这是 Local First 底座，不应延迟到后续阶段。

### 路线二："复利感知优先"（GLM-5.2 / Qwen3.7-Pro）

**核心主张：先让复利变得可见，再围绕复利构建经营闭环。**

```
Phase06: 复用感知与能力派生（module_reuse_summary + capability_summary + Dashboard 增强）
    ↓
Phase07: 模板级复用 + 度量收口 + 数据导出 + 真实项目 dry-run
```

**逻辑链：**
- MVP0.1 已验证"资产能登记"，但未验证"资产能复利"
- 没有复用感知，PSCO 的飞轮转不起来，差异化护城河不存在
- 复用感知是后续一切（AI Composition、Decision 复用）的上下文前提
- 因此：先让"能力增长"变成可观测事实，再围绕它构建操作循环

**额外强调：** Qwen 将模板级复用（SaveModuleCompositionAsTemplate）作为独立方向 B，认为这是 "v0.2 第一优先级"（引用 summarize-feedback §4.9）。

---

## 3. 我的仲裁：两条路线不是对立，而是互补

### 3.1 分歧的本质

两条路线的分歧，本质上是**对同一个问题的不同回答**：

> "MVP0.2 的第一块砖应该放在哪里？"

- 路线一认为：第一块砖是**降低进入门槛**（Onboarding），让更多数据进入系统
- 路线二认为：第一块砖是**让复利可见**（Reuse Awareness），让已有数据产生价值反馈

但这两块砖**并不互斥**。Onboarding 解决"数据怎么进来"，Reuse Awareness 解决"进来的数据怎么产生价值"——它们是一个完整的价值闭环的两端。

### 3.2 我的仲裁结论

**我支持两条路线的融合，而非在二者之间二选一。** 具体而言：

> **MVP0.2 应以"Onboarding + 导出/备份 + 复用感知基础"作为第一阶段的融合入口，以"Review Loop + 模板复用 + 派生智能深化"作为第二阶段的经营闭环。**

理由如下：

**理由 1：Onboarding 和 Reuse Awareness 是互补的，不是竞争的。**

- 没有 Onboarding，用户的数据进不来，Reuse Awareness 展示的是空数据
- 没有 Reuse Awareness，用户进来了也看不到"为什么这个系统值得用"
- 二者应该**同时推进**，而不是先后推进

**理由 2：导出/备份不应被延迟到 Phase07。**

- 这是数据主权底座，是 Local First 原则的硬约束
- MVP0.1 已积累真实资产数据，延迟导出意味着用户数据所有权承诺持续未兑现
- DS-V4-Flash 和 DS-V4-Pro 的独立论证均指向这一点，我认可这个判断

**理由 3：模板复用的价值依赖复用感知的先期存在。**

- 用户需要先看到"哪些模块被复用了"，才能理解"为什么要把 Module 组合保存为模板"
- 因此模板复用（GLM/Qwen 的 Phase07 核心功能）放在第二阶段是正确的，但第一阶段必须先建立复用感知的认知基础

**理由 4：Review Loop 是让"可观测"变成"可行动"的桥梁。**

- 路线二（GLM/Qwen）正确指出了"复用感知"的核心价值，但未充分回答"用户看到复用数据后应该做什么"
- 路线一（GPT54/DS）的 Review Loop 恰好补上了这个缺口：从"我看见复用了"到"我决定沉淀更多模块 / 创建更多绑定 / 记录关键决策"
- 因此 Review Loop 应作为第二阶段的核心，承接第一阶段已建立的复用感知

---

## 4. 推荐的融合方案

基于以上仲裁，我建议 MVP0.2 采用以下融合方案：

### 4.1 整体主题

> **从 Asset Registry 到 Operating System：先让数据进得来、复利看得见，再让经营转得动。**

### 4.2 两个 Phase 的融合拆分

#### Phase06：Onboarding + 数据主权 + 复用感知基础

**融合了路线一的 Block 1（Onboarding + 导出/备份）和路线二的 Phase06（复用感知与能力派生）。**

| 能力 | 来源 | 说明 |
| --- | --- | --- |
| First-run onboarding | 路线一（GPT54/DS） | 受控流程引导用户建立首个 Product → Repository → Module → Decision |
| Draft-first / partial-entry | 路线一（GPT54/DS） | 对象允许最小字段创建，逐步补全 |
| Controlled import helper | 路线一（GPT54/DS） | 从仓库 URL 预填字段，辅助录入 |
| **ExportAssets + Backup** | 路线一（DS-V4-Flash/DS-V4-Pro） + 路线二（GLM/Qwen 的导出） | JSON 导出端点 + pg_dump 备份脚本，**必须在 Phase06 内闭合** |
| `module_reuse_summary` 派生视图 | 路线二（GLM/Qwen） | 每个 Module 被哪些 Product 使用、复用次数 |
| `capability_summary` 派生视图 | 路线二（GLM/Qwen） | 按能力分类聚合模块数、稳定模块数、复用模块数 |
| Module Detail "Used By" 区块 | 路线二（GLM/Qwen） | 展示复用关系 |
| Module List 复用度排序 | 路线二（GLM/Qwen） | 按复用次数排序/筛选 |
| Low-friction decision capture | 路线一（GPT54/DS） | ADR-lite，title-first 渐进补全 |

**Phase06 的 DoD（融合版）：**
1. 新用户可在一次会话内完成首个 Product + Repository + Module + Decision 录入
2. 对象允许最小字段创建并逐步补全
3. 导出/备份基础路径已闭合（JSON 导出 + pg_dump 脚本）
4. module_reuse_summary 与 capability_summary 派生视图已落地且可复验
5. Module Detail / List 对应增强已通过联调
6. 模块复用率度量指标可在 Dashboard 观测

#### Phase07：Review Loop + 模板复用 + 派生智能深化

**融合了路线一的 Block 2+3（Review Loop + Derived Intelligence）和路线二的 Phase07（模板复用 + 度量收口 + dry-run）。**

| 能力 | 来源 | 说明 |
| --- | --- | --- |
| Daily Review | 路线一（GPT54/DS） | 进入系统先看 Current Focus，一键跳转 |
| Weekly Review | 路线一（GPT54/DS） | 汇总新增 Decision / Module / Release / 缺口 / 复用进展，回流结论 |
| Action handoff | 路线一（GPT54/DS） | Review 结论转为最小后续动作，锚定到既有实体 |
| Feedback → Decision → Update 闭环 | 路线一（GPT54/DS） | 从 Dashboard 的 gap/signal 到 Decision 到实体更新 |
| **模板级复用最小版** | 路线二（GLM/Qwen） | SaveModuleCompositionAsTemplate + ApplyModuleCompositionTemplate |
| Product Create 模板预填 | 路线二（GLM/Qwen） | 新建 Product 时基于模板预填 |
| Derived Intelligence 深化 | 路线一（GPT54/DS） + 路线二（GLM/Qwen） | 候选能力提示、跨产品复用标记、release_freshness、decision_link_density |
| 资产导入耗时度量 | 路线二（GLM/Qwen） | 基于已有 timestamp 计算 |
| 真实项目 dry-run | 路线二（Qwen 强调） | 以 Rento-miniX 或等价真实项目走完整流程 |

**Phase07 的 DoD（融合版）：**
1. 用户可完成一次完整 weekly review，结论能回流到既有实体
2. Dashboard 能承接"看见问题 → 做出动作 → 记录决策"主线
3. 能力模板最小版可完成"保存组合 → 新建 Product 时预填"闭环
4. 至少有一组派生指标可解释"能力增长 / 复用增长"
5. 资产导入耗时度量可观测
6. 至少一个真实项目 dry-run 走通
7. 未引入超出 TECH_STACK_BASELINE 的技术选择

---

## 5. 我支持的观点（附理由）

### 5.1 我完全支持的观点

| 观点 | 出处 | 支持理由 |
| --- | --- | --- |
| 不扩实体，先做复利 | GLM/Qwen 核心主张 | 与 MVP0.1 收敛原则一致，长期实体应在 operating loop 验证后再引入 |
| 导出/备份必须在早期闭合 | DS-V4-Flash/DS-V4-Pro | Local First 是差异化底座，延迟闭合会侵蚀用户信任 |
| Dashboard 应成为经营起点而非仅看板 | GPT54/DS-V4-Flash/DS-V4-Pro | 如果没有"从看到做"的闭环，复用感知只是另一个 widget |
| 模板复用应限定为"Module 组合快照" | Qwen 的语义边界强调 | 避免过度工程化为完整模板系统，保持最小可用 |
| 真实项目 dry-run 是独立的验收交付物 | Qwen 的独立交付物建议 | 不被其他功能实现挤压，确保飞轮验证不流于形式 |
| Decision 复用应后移到 MVP0.3 | Qwen 的明确后移理由 | 需要复用感知作为上下文前提，MVP0.2 阶段推进会增加不必要复杂度 |
| 派生视图用 SQL 查询/视图，不引入物化视图 | GLM/Qwen 的实现方式 | 个人级数据量不需要物化视图，保持 Local First 可导出性 |
| 继续坚持 TanStack Router / Query / .proto 单一事实源 | GPT54/DS-V4-Flash/DS-V4-Pro 的技术约束 | 已被 phase01-05 验证有效，不引入第二套事实源 |

### 5.2 我部分支持但需调整的观点

| 观点 | 出处 | 我的调整建议 |
| --- | --- | --- |
| 先做复用感知，再做 Onboarding | GLM/Qwen 的排序 | 我认为二者应并行推进（见 Phase06 融合方案），因为无数据的复用感知是空壳，无复利感知的 Onboarding 缺乏"aha moment" |
| 导出放在 Phase07 | GLM/Qwen 的排序 | 我认为导出/备份应放在 Phase06（见 §3.2 理由 2），因为它是最基础的数据主权承诺 |
| 三个 Block 而非两个 Phase | GPT54/DS 的建议 | 我认可 GLM/Qwen 的 2-phase 结构更清晰，但 Phase06 应融合 Block 1 的 Onboarding + Block 3 的复用感知基础 |

---

## 6. 我反对或有保留的观点（附理由）

### 6.1 我明确反对的观点

| 观点 | 出处 | 反对理由 |
| --- | --- | --- |
| （无） | — | 五份评审在核心方向上高度一致，没有出现需要明确反对的重大分歧 |

### 6.2 我有保留的观点

| 观点 | 出处 | 保留理由 |
| --- | --- | --- |
| 将"复用感知"作为 MVP0.2 的唯一主线，Onboarding 不做或大幅延后 | GLM/Qwen 的隐含取向（未明确说"不做 Onboarding"，但 Phase06 未包含 Onboarding 相关内容） | 如 §3.2 理由 1 所述，Onboarding 和 Reuse Awareness 是互补的。如果 Phase06 只做复用感知不做 Onboarding，系统会变成"对已有数据的人有价值，对新人仍然门槛高"的状态 |
| 将 Review Loop 完全延后到 MVP0.3 | GLM/Qwen 的路线中未包含 Review Loop | 如果 MVP0.2 只产出"复用感知"而不产出"基于复用感知的日常操作"，那么用户看到复用数据后不知道下一步该做什么，复用感知会退化为"好看但不可操作"的看板 |
| 能力模板的 SaveModuleCompositionAsTemplate 语义 | GLM/Qwen 的命名 | 我建议在正式 spec 中进一步明确：模板只是"Module 组合快照"，不是独立实体，不拥有 Product，不引入模板 CRUD 列表，避免过度工程 |

---

## 7. 各评审的独特贡献（值得保留的差异化视角）

| 评审 | 独特贡献 | 建议在后续 `/plan` 中采纳 |
| --- | --- | --- |
| **GPT54** | 最先提出"从 Asset Registry 到 Operating System"的核心主题，以及"Operating Review Loop"作为最关键主线的判断 | ✅ 作为 MVP0.2 整体叙事框架 |
| **DeepSeek-V4-Flash** | 最先且最强烈地指出导出/备份是规格合规负债，必须在本阶段闭合 | ✅ 作为 Phase06 的贯穿性工作项 |
| **DeepSeek-V4-Pro**（本文作者） | 回看 2026-08-04 方案评审的 9 条建议，发现 2 条 🔴 高优先级建议在 MVP0.1 后仍未闭合，为 MVP0.2 的优先级排序提供了历史连续性论证 | ✅ 作为 Phase06 优先级的佐证 |
| **GLM-5.2** | 最先系统化地将复用感知、模板复用、数据导出组织为两个 phase 的清晰拆分 | ✅ 作为 Phase06/07 拆分的基础框架 |
| **Qwen3.7-Pro** | 最详细地定义了 Done 标准、度量指标、风险与缓解措施，并强调真实项目 dry-run 的独立交付物地位 | ✅ 作为 Phase06/07 验收标准的具体化输入 |

---

## 8. 最终建议

### 8.1 对后续 `/plan` 的输入

基于以上五方评审的交叉汇总与仲裁，我建议后续正式 `/plan` 以以下方向为主轴：

> **MVP0.2 的主题：从 Asset Registry 到 Operating System**
>
> **Phase06：Onboarding + 数据主权闭合 + 复用感知基础**
> **Phase07：Review Loop + 模板复用 + 派生智能深化 + 真实项目 dry-run**

### 8.2 不可谈判的边界（来自五方强共识）

1. 不新增核心实体（Product / Module / Release / Decision / Repository 保持不变）
2. 不做自动代码扫描 / 知识图谱 / Rust Intelligence Layer
3. 不做 AI Assistant 一级主导航
4. 不做重量级 GitHub OAuth / 全自动导入
5. 不引入第二套技术栈（路由 / 合同 / ORM / UI 框架）
6. 导出 / 备份必须在本阶段闭合
7. module_reuse_summary + capability_summary 必须落地
8. Capability 继续作为派生层

### 8.3 建议在正式 `/plan` 中进一步明确的开放问题

以下问题在五份评审中未被充分讨论或存在不同隐含假设，建议在正式 `/plan` 阶段明确：

1. **Onboarding 的具体范围**：first-run wizard 的步骤数、跳过机制、与现有 Create 页面的关系
2. **Review Loop 的实体建模**：review 结论是否需要一个最小持久化结构（如 `review_notes` 字段附加到实体），还是纯前端聚合展示
3. **能力模板的存储方式**：模板是存为独立的 `module_composition_templates` 表，还是作为 Product 的一个标记字段
4. **派生视图的更新策略**：实时 SQL 查询 vs. 缓存刷新 vs. 物化视图，需在 spec 中明确
5. **真实项目 dry-run 的执行时机**：Phase07 中期（边做边验）还是 Phase07 收尾（全部做完再验）

---

## 9. 总结

五份独立评审在核心方向上高度收敛，这是对 MVP0.2 方向的最强验证。两条路线（"经营闭环优先" vs. "复利感知优先"）的差异不是对错之争，而是**侧重点和先后顺序的差异**。我的仲裁结论是：二者应融合而非二选一。

**融合方案：**

```
Phase06: Onboarding + 导出/备份 + 复用感知基础
    │  （融合：路线一的 Block 1 + 路线二的 Phase06）
    │  （核心：让数据进得来、复利看得见、数据主权闭合）
    │
    ▼
Phase07: Review Loop + 模板复用 + 派生智能深化 + Dry-run
    │  （融合：路线一的 Block 2+3 + 路线二的 Phase07）
    │  （核心：让经营转得动、复利可行动、飞轮可验证）
    │
    ▼
MVP0.2 收口 → MVP0.3 候选：GitHub 集成 / AI context-aware 增强 / Decision 复用
```

**一句话收口：**

> **五方评审共识大于分歧，MVP0.2 应同时推进"降低进入门槛"和"让复利可见"，以 Phase06 融合 Onboarding + 数据主权 + 复用感知，以 Phase07 承接 Review Loop + 模板复用 + 派生智能，最终将 PSCO 从"资产登记系统"升级为"可日常经营的资产复利系统"。**

---

*End of PSCO-mvp02-summarize-feedback-deepseekv4pro.md*