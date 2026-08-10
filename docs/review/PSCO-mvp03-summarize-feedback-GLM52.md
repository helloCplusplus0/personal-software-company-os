# Personal Software Company OS

# MVP0.3 五份方向评审交叉汇总与评估

**Author:** GLM-5.2  
**Date:** 2026-08-10  
**Role:** 作为五份 MVP0.3 方向评审的作者之一（`PSCO-mvp03-GLM52.md`），对五份评审进行交叉汇总、立场归档与支持 / 反对评估，作为后续正式 `/plan` 的参考仲裁入口  
**Document Type:** `review`  
**Status:** 供后续正式 `/plan` 参考，不直接构成正式 phase 命名、spec 路径或执行指令

---

## 1. 文档定位与汇总方法

### 1.1 本文要完成的事

本文对以下五份 MVP0.3 方向评审进行交叉汇总：

| 文档 | 作者 | 核心主张 |
| --- | --- | --- |
| `PSCO-mvp03-GPT54.md` | GPT54 | Operating Review Loop + Template Reuse + Real-Project Dry-Run（三主线并列） |
| `PSCO-mvp03-GLM52.md` | GLM-5.2 | Operating Review Loop（唯一主线）+ Derived Intelligence（左翼）+ Template Reuse（右翼）+ Dry-Run（验收闸） |
| `PSCO-mvp03-qwen37pro.md` | Qwen3.7-Pro | Operating Review Loop 为主体，模板 / 度量 / dry-run 为可选 phase08 |
| `PSCO-mvp03-DPv4flash.md` | DeepSeek-V4-Flash | Operating Review Loop + Template-Driven Compound（三顺序主线） |
| `PSCO-mvp03-DPv4pro.md` | DeepSeek-V4-Pro | Venture + Decision Intelligence + AI Context Enhancement（假定 phase07 先完成 MVP0.2 阶段二） |

### 1.2 汇总原则

1. **诚实归档立场。** 每份文档的主张如实记录，不扭曲原意。
2. **区分强共识与弱分歧。** 五份一致的结论可直接作为后续 `/plan` 输入；存在分歧的结论需给出仲裁建议。
3. **以 MVP0.2 共识为仲裁基准。** 凡与 `PSCO-mvp02-summarize-feedback.md` 冲突之处，以共识文档为准。
4. **作者立场透明声明。** 本文作者即 `PSCO-mvp03-GLM52.md` 作者，在涉及自身主张时会显式标注，不回避自我评估。
5. **不冻结 phase 名称。** 遵守 `plan.md §4` 与 `project_rules.md` 切换条件。

### 1.3 关键背景澄清：MVP0.3 的版本语义

五份评审中存在一个隐含的版本语义分歧，必须先澄清：

- **`PSCO-mvp02-summarize-feedback.md §6.1`** 将 MVP0.2 定义为两个阶段：
  - 阶段一（已完成 = phase06）：Onboarding + 数据主权 + 复用感知基础
  - 阶段二（尚未完成）：Operating Review Loop + 模板级复用 + 派生智能深化 + 真实项目 dry-run

- **四份评审（GPT54 / GLM52 / Qwen37pro / DPv4flash）** 将「MVP0.3」理解为**完成 MVP0.2 阶段二**，即把 MVP0.2 的未完成部分作为下一阶段主线。

- **DPv4pro** 将「MVP0.3」理解为**MVP0.2 阶段二完成之后的下一版本**，即假定 phase07 先收口 MVP0.2，然后 MVP0.3 从 Venture / Decision Intelligence / AI 开始。

这个语义分歧是五份评审最根本的结构性差异，后续所有分歧都源于此。本文采用**四份评审的版本语义**（MVP0.3 = 完成 MVP0.2 阶段二），理由见第 5.1 节。

---

## 2. 五份评审立场详细概览

### 2.1 GPT54 — 三主线并列

**核心主张：** `Operating Review Loop + Template Reuse + Real-Project Dry-Run` 三主线并列。

**结构特征：**
- 三主线地位平等，无明确主从
- Derived Intelligence 未单独列出（隐含并入 Template Reuse）
- 明确反对扩长期实体、反对重智能层、反对模板独立化
- 不冻结 phase 名称

**关键判断：**
> phase06 之后最该验证的不是「能不能扩张长期模型」，而是「能不能把已沉淀资产转化为持续经营动作」。

### 2.2 GLM52 — 一主两翼一闸

**核心主张：** `Operating Review Loop（唯一主线）+ Derived Intelligence Deepening（左翼）+ Template Reuse（右翼）+ Real-Project Dry-Run（验收闸）`。

**结构特征：**
- Operating Review Loop 为唯一主线，其余为支撑
- Derived Intelligence 独立成左翼，不并入 Template Reuse
- 基于一手实现经验，指出三处「接口存在、语义未闭合」隐患
- 强调「先消费再扩展」
- 不冻结 phase 名称

**关键判断：**
> Review Loop 是消费已有能力的唯一枢纽；没有它，Template Reuse 缺触发场景，Derived Intelligence 缺消费方，Dry-Run 无法验证核心命题。

### 2.3 Qwen3.7-Pro — 主体 + 可选 phase08

**核心主张：** 以 Operating Review Loop 为方向 A 主体，方向 B（Decision Intelligence）与方向 C（AI Enhancement）为可选深化；拆分为 phase07（Review Loop）+ phase08（度量 + 模板 + dry-run，可选）。

**结构特征：**
- 最保守的执行收缩——phase08 标注为可选
- 唯一明确给出 phase 编号建议（phase07 / phase08）
- 给出具体动作清单（StartDailyReview / RecordReviewConclusion / CreateActionFromReview）
- 给出度量指标表（资产导入耗时必须落地，其余可观测不强制）
- 明确 Decision 高级复用与 AI 增强后移理由

**关键判断：**
> 不扩实体，先做流程；不引集成，先做闭环；不碰 AI，先用真实项目验证经营收益。

### 2.4 DeepSeek-V4-Flash — 经营复盘 + 模板驱动复利

**核心主张：** `Operating Review Loop + Template-Driven Compound`，三顺序主线（Review Loop → Template → Dry-Run + Derived Intelligence）。

**结构特征：**
- 第一原则：把「一次性录入」推进为「周期性经营」
- 强调「经营视角」是 Command Center 四重视角中唯一尚未落地的
- 明确反对 Review Loop 演化为任务管理器
- 给出技术与实现约束（TanStack Router / Query-first / .proto 单一合同源）
- 不冻结 phase 名称

**关键判断：**
> phase01~06 证明了「能登记、能看见复利」；mvp03 要证明「会被周期性经营，且经营能加速复利」。

### 2.5 DeepSeek-V4-Pro — 智能增强增长系统

**核心主张：** MVP0.3 = `Venture & Strategic Context + Decision Intelligence + AI Context Enhancement`，假定 phase07 先完成 MVP0.2 阶段二。

**结构特征：**
- 唯一主张引入新实体（Venture，可选）
- 唯一主张 MVP0.3 引入 AI 增强层
- 唯一主张 Decision Intelligence（检索 / 引用 / 相似匹配）进入 MVP0.3
- 明确给出 phase07~10 四阶段拆分
- 视角最长期、最雄心勃勃

**关键判断：**
> MVP0.3 的任务，是在经营闭环已经建立的基础上，让 PSCO 从「记录和经营的系统」升级为「加速决策和增长的系统」。

---

## 3. 强共识区域（五份一致，可直接作为 /plan 输入）

以下七项结论五份评审完全一致，且与 `PSCO-mvp02-summarize-feedback.md` 共识对齐，可直接作为后续正式 `/plan` 的稳定输入。

### 3.1 Operating Review Loop 是下一阶段第一优先级

**五份一致支持。** 无论结构如何划分，Review Loop 都是下一阶段的核心或主体。这直接承接 `PSCO-mvp02-summarize-feedback.md §6.1` 阶段二。

### 3.2 Template Reuse 必须是「最小版」，不是完整模板平台

**五份一致支持。** 边界冻结为「Module 组合快照 + 新建预填辅助」，不做模板版本管理、参数化、模板 CRUD 列表。这承接 `PSCO-mvp02-summarize-feedback.md §5.1` 第 5 组。

### 3.3 Real-Project Dry-Run 必须成为独立验收证据

**五份一致支持。** 真实项目 dry-run 不替代 fixture 验收，但必须独立留存，与 fixture 验收并列。优先以 Rento-miniX 或等价真实项目走完整流程。

### 3.4 不扩长期实体（Opportunity / Feature / Experiment）

**五份一致支持。** 这些在长期模型成立，但当前优先级不如 operating loop。承接 `PSCO-mvp02-summarize-feedback.md §5.2`。

### 3.5 不做重智能层（AI 一级工作台 / 自动扫描 / 知识图谱 / Rust）

**五份一致支持。** 现在更稀缺的是高质量 operating context，不是更重的自动化。承接 `PSCO-mvp02-summarize-feedback.md §5.2`。

### 3.6 不做 GitHub OAuth / 全自动导入

**五份一致支持。** phase06 已证明手动 + 辅助录入能完成闭环，集成复杂度留给后续。

### 3.7 守住 TECH_STACK_BASELINE 与已冻结工程约束

**五份一致支持。** `.proto` 唯一合同源、query 纯只读、写入收敛到 application owner、不引入第二套框架。`Capability` 继续派生化，不建重表。

---

## 4. 主要分歧点与各方立场

### 4.1 分歧一：MVP0.3 的版本语义

| 立场 | 支持方 | 含义 |
| --- | --- | --- |
| MVP0.3 = 完成 MVP0.2 阶段二 | GPT54 / GLM52 / Qwen37pro / DPv4flash | Review Loop + Template + Derived Intelligence + Dry-Run 就是 MVP0.3 全部 |
| MVP0.3 = MVP0.2 阶段二完成后的下一版本 | DPv4pro | phase07 先收口 MVP0.2，MVP0.3 从 Venture / Decision Intel / AI 开始 |

**这是最根本的分歧。** DPv4pro 的框架更长期，但会拉伸时间线（phase07~10 四阶段）；四份评审的框架更紧凑，聚焦未完成的核心闭环。

### 4.2 分歧二：主线结构（并列 vs. 主从 vs. 顺序）

| 结构 | 支持方 | 特征 |
| --- | --- | --- |
| 三主线并列 | GPT54 | Review / Template / Dry-Run 地位平等 |
| 一主两翼一闸 | GLM52 | Review Loop 唯一主线，其余支撑 |
| 主体 + 可选 phase | Qwen37pro | Review Loop 主体，phase08 可选 |
| 三顺序主线 | DPv4flash | Review → Template → Dry-Run 有先后 |
| 三新主线（假定 phase07 先完成） | DPv4pro | Venture / Decision Intel / AI |

**GLM52 与 DPv4flash 都强调 Review Loop 优先，但 GLM52 更强调主从关系，DPv4flash 更强调顺序关系。** Qwen37pro 最保守，把第二阶段标为可选。

### 4.3 分歧三：Derived Intelligence Deepening 的地位

| 立场 | 支持方 |
| --- | --- |
| 独立成左翼，必须进入下一阶段 | GLM52 |
| 并入主线 C（与 Dry-Run 同组） | DPv4flash |
| 隐含并入 Template Reuse | GPT54 |
| 并入 phase08 可选 | Qwen37pro |
| 并入 phase07（假定 MVP0.2 收口） | DPv4pro |

**只有 GLM52 明确主张 Derived Intelligence 应独立成线。** 其余四份把它并入其他组。这直接影响 `capability_summary` 是否需要从计数进化为缺口提示 / 复用推荐。

### 4.4 分歧四：Venture 是否引入

| 立场 | 支持方 |
| --- | --- |
| 不引入（继续可选不存在） | GPT54 / GLM52 / DPv4flash / Qwen37pro |
| 引入为可选实体（最小版） | DPv4pro |

**DPv4pro 是唯一主张引入 Venture 的。** 理由是「经营闭环建立后，为什么做成为自然需求」。四份评审认为 Venture 优先级不如 operating loop。

### 4.5 分歧五：AI Enhancement 时机

| 立场 | 支持方 |
| --- | --- |
| 后移到 MVP0.4+ | GPT54 / GLM52 / Qwen37pro / DPv4flash |
| MVP0.3 首次引入（phase10） | DPv4pro |

**DPv4pro 是唯一主张 MVP0.3 引入 AI 的。** 理由是「Personal Context Layer 前提已开始成立」。四份评审认为 review 事实与 decision 回流尚未稳定，AI 增强为时过早。

### 4.6 分歧六：Decision Intelligence 时机

| 立场 | 支持方 |
| --- | --- |
| 后移（Decision 高级复用留 MVP0.3+） | GPT54 / GLM52 / Qwen37pro / DPv4flash |
| MVP0.3 主线（phase09） | DPv4pro |

**DPv4pro 是唯一主张 MVP0.3 推进 Decision Intelligence 的。** 理由是「决策数量增长后历史决策价值需释放」。四份评审认为应先让 Decision 在 review 中被消费，再谈智能复用。

### 4.7 分歧七：phase 命名是否冻结

| 立场 | 支持方 |
| --- | --- |
| 不冻结，只给方向 | GPT54 / GLM52 / DPv4flash |
| 给出编号建议但不强制 | Qwen37pro（phase07/08） |
| 明确给出 phase07~10 | DPv4pro |

---

## 5. 我的评估：支持与反对

> **透明声明：** 本文作者即 `PSCO-mvp03-GLM52.md` 作者。以下评估中涉及自身主张时会显式标注，但我会对自身主张保持同样的批判标准。

### 5.1 支持：MVP0.3 = 完成 MVP0.2 阶段二（四份评审的版本语义）

**支持方：** GPT54 / GLM52 / Qwen37pro / DPv4flash  
**反对方：** DPv4pro

**我的立场：支持四份评审。**

理由：

1. **MVP0.2 阶段二尚未完成，不能跳过。** `PSCO-mvp02-summarize-feedback.md §6.1` 明确将 Review Loop + 模板复用 + 派生智能深化 + dry-run 列为阶段二必做。phase06 只完成了阶段一。跳过阶段二直接做 Venture / AI，等于在未验证 operating loop 的前提下扩张，违反「先验证再扩张」原则。

2. **DPv4pro 的框架假设 phase07 先收口 MVP0.2，再启动 MVP0.3。** 这个假设本身合理，但它把 MVP0.3 拉长到 phase08~10 四阶段，时间线过长。更紧凑的做法是把 MVP0.2 阶段二直接命名为 MVP0.3 的核心内容，在一个版本内完成。

3. **版本编号应反映实质演进，不是机械递增。** 从「资产登记」到「经营系统」已经是实质演进，足以构成 MVP0.3。不需要等到 Venture / AI 才算新版本。

**对 DPv4pro 的承认：** DPv4pro 正确识别了 Venture / Decision Intelligence / AI 是 PSCO 的长期方向，这些方向最终需要落地。本文只是反对它们的「时机」，不反对「方向」。DPv4pro 的文档更适合作为 MVP0.4+ 的路线图输入，而非 MVP0.3 的执行计划。

### 5.2 支持：Operating Review Loop 为第一优先级

**五份一致支持。**

**我的立场：支持，且支持 GLM52 的「唯一主线」结构。**

理由已在 `PSCO-mvp03-GLM52.md §6.1` 详述：Review Loop 是消费 phase05 Feedback、phase06 ReuseSummaryRead、phase03 Decision 的唯一枢纽。没有它，这些能力会变成孤儿。

**对 GPT54 三主线并列的保留：** 三主线并列的风险是焦点分散。如果 Review Loop、Template、Dry-Run 同期推进，Review Loop 容易被后两者的工程压力挤压。GLM52 的主从结构更利于聚焦。

**对 Qwen37pro phase08 可选的保留：** 把模板与 dry-run 标为可选，风险是它们可能被无限后移，导致 MVP0.2 阶段二永远无法真正收口。模板与 dry-run 应是 MVP0.3 的必要组成，不是可选。

### 5.3 支持：Derived Intelligence Deepening 应独立成线（GLM52 主张）

**支持方：** GLM52  
**部分支持：** DPv4flash（并入主线 C）  
**隐含反对：** GPT54（并入 Template）、Qwen37pro（并入 phase08）、DPv4pro（并入 phase07）

**我的立场：支持 GLM52，但接受妥协。**

理由：

1. **`capability_summary` 当前语义太薄。** 只有 `supporting_module_count` 计数，无法支撑 review 的「能力缺口提示」与「复用机会推荐」。这是工程事实，不是理论判断。

2. **Derived Intelligence 与 Template Reuse 是两条工程线。** 前者改的是派生读接口的语义（`reusesummary` 切片内深化），后者新增的是组合快照保存与预填（新写路径）。把它们合并会导致切片边界混乱。

3. **Review Loop 需要它作为动作依据。** review 中「下一步做什么」需要能力缺口与复用机会作为输入，否则 review 只能基于 Feedback 信号，无法基于复利事实。

**妥协方案：** 如果工程资源紧张，Derived Intelligence 可作为 Review Loop 的子任务而非独立 phase，但必须在 MVP0.3 内完成最小深化，不能后移。

### 5.4 支持：Template Reuse 最小版边界

**五份一致支持。**

**我的立场：支持，且支持 DPv4flash 的「Template-Driven Compound」命名。**

DPv4flash 把模板复用的价值定位为「让 Compound 可行动」，这比单纯「新建预填」更准确。模板复用的真正价值是缩短从 review 决策到新产品创建的路径，不是多一个资产栏目。

### 5.5 支持：Real-Project Dry-Run 独立验收

**五份一致支持。**

**我的立场：支持，且支持 GLM52 的「验收闸」定位。**

Dry-Run 不是顺手验证，而是 MVP0.3 能否成立的关键证据。它应与 fixture 验收并列，形成独立验收记录。

### 5.6 反对：MVP0.3 引入 Venture（DPv4pro 主张）

**反对方：** GPT54 / GLM52 / Qwen37pro / DPv4flash  
**支持方：** DPv4pro

**我的立场：反对 MVP0.3 引入，支持作为 MVP0.4+ 候选。**

理由：

1. **Venture 没有 operating loop 就没有消费场景。** Venture 的价值是回答「为什么做」，但如果用户还没有 daily / weekly review 习惯，Venture 只是多一个可选字段，不会被真正使用。

2. **DPv4pro 的理由「经营闭环已建立」是假定，不是事实。** phase06 完成的是阶段一，operating loop 尚未验证。在 operating loop 验证前引入 Venture，等于在未验证的地基上加盖。

3. **Venture 会分散 MVP0.3 焦点。** MVP0.3 已有 Review Loop + Template + Derived Intelligence + Dry-Run 四项，再加入 Venture 会超出单版本承载。

**对 DPv4pro 的承认：** DPv4pro 正确指出「当用户看到多个 Product 时，自然会问它们之间什么关系」。这个需求真实存在，但应在 operating loop 验证后再承接，作为 MVP0.4 的自然延伸。

### 5.7 反对：MVP0.3 引入 AI Context Enhancement（DPv4pro 主张）

**反对方：** GPT54 / GLM52 / Qwen37pro / DPv4flash  
**支持方：** DPv4pro

**我的立场：反对 MVP0.3 引入，支持作为 MVP0.4+ 候选。**

理由：

1. **AI 增强的前提是稳定的 operating context。** DPv4pro 列举的 AI 能力（决策草稿、模块复用提示、周报总结）都依赖 review 事实与 decision 回流。这些在 MVP0.3 才开始建立，尚未稳定。

2. **DPv4pro 自身也承认依赖关系。** 其文档 §7.2 明确 phase10（AI）依赖 phase07（review 数据）、phase08（Venture 战略上下文）、phase09（Decision Intelligence 结构化数据）。这意味着 AI 至少在三个 phase 之后，不可能在 MVP0.3 内完成。

3. **AI 引入会破坏 MVP0.3 的执行收缩。** MVP0.3 的核心命题是验证 operating loop。引入 AI 会增加模型配置、失败降级、草稿编辑等大量工程工作，挤压核心闭环验证。

**对 DPv4pro 的承认：** DPv4pro 提出的 AI 架构约束（显式触发、永远可编辑、不访问代码仓库、失败不阻塞、模型可配置）是正确的长期原则。这些原则应在未来引入 AI 时严格遵守。

### 5.8 反对：MVP0.3 推进 Decision Intelligence（DPv4pro 主张）

**反对方：** GPT54 / GLM52 / Qwen37pro / DPv4flash  
**支持方：** DPv4pro

**我的立场：反对 MVP0.3 推进 Decision Intelligence，支持先让 Decision 在 review 中被消费。**

理由：

1. **Decision 的当务之急是「被消费」而非「被智能检索」。** 当前 Decision 仍是「被记录后被查看」，尚未进入 review 动作流。先让它在 review 中成为「待处理 → 升级 → 回流」的中心，再谈相似匹配与引用推荐。

2. **Decision Intelligence 需要足够决策量才有价值。** DPv4pro 自己说「当决策数量增长到几十上百条后」。当前系统决策量远未达到这个规模，提前建检索与相似匹配会是空跑。

3. **Decision Intelligence 是重工程。** 全文检索、相似度匹配、引用关系管理都是独立工程线，会显著增加 MVP0.3 复杂度。

### 5.9 支持：不冻结 phase 命名

**支持方：** GPT54 / GLM52 / DPv4flash  
**部分支持：** Qwen37pro（给建议但不强制）  
**反对方：** DPv4pro（明确 phase07~10）

**我的立场：支持不冻结。**

理由：`plan.md §4` 与 `project_rules.md` 明确要求下一阶段 phase 入口建立前不预设名称。Qwen37pro 的 phase07/08 建议可作为讨论参考，但不应在评审文档中冻结。DPv4pro 的 phase07~10 四阶段拆分过于超前，违反「不预设 phase 名称」约束。

---

## 6. 综合建议

### 6.1 我建议的 MVP0.3 推进结构

基于以上评估，我给出综合建议：

> **MVP0.3 = 完成 MVP0.2 阶段二，以 Operating Review Loop 为唯一主线，Derived Intelligence Deepening 与 Template Reuse 为左右两翼，Real-Project Dry-Run 为验收闸。**

这采用 GLM52 的主从结构，但吸收其他四份的合理要素：

- 吸收 **DPv4flash** 的「Template-Driven Compound」命名与「第一原则：周期性经营」
- 吸收 **Qwen37pro** 的具体动作清单（StartDailyReview / RecordReviewConclusion / CreateActionFromReview）与度量指标表
- 吸收 **GPT54** 的「先消费再扩展」原则与非目标清单
- 保留 **GLM52** 的主从结构与 Derived Intelligence 独立性

### 6.2 建议的推进顺序

```
MVP0.3 候选推进顺序（不冻结 phase 名称）：

阶段一：Operating Review Loop 基础
  └─ daily / weekly review 入口
  └─ Feedback -> Decision -> Update 闭环
  └─ review 结论回流到既有实体
  └─ Action Handoff（锚定既有实体，不做任务管理器）
        ↓
阶段二：Derived Intelligence Deepening + Template Reuse + Dry-Run
  └─ capability_summary 从计数进化到缺口提示 / 复用推荐
  └─ Module 组合快照 + 新建 Product 预填
  └─ 真实项目 dry-run（独立验收记录）
```

**与 Qwen37pro 的差异：** Qwen37pro 把阶段二标为可选。我反对「可选」——模板复用与 dry-run 是 MVP0.2 阶段二的必做项，标注可选会导致 MVP0.2 永远无法收口。

**与 DPv4pro 的差异：** DPv4pro 把 Venture / Decision Intelligence / AI 放在 phase08~10。我把它们整体后移到 MVP0.4+，MVP0.3 聚焦完成 operating loop 验证。

### 6.3 建议的 MVP0.4+ 路线图（吸收 DPv4pro 的长期视角）

DPv4pro 的 Venture / Decision Intelligence / AI 方向长期成立，只是时机不对。我建议作为 MVP0.4+ 路线图：

```
MVP0.4 候选（需 MVP0.3 operating loop 验证通过后启动）：
  └─ Venture & Strategic Context（可选实体，承接 review 中的「为什么做」）
  └─ Decision Intelligence（决策量足够后的检索 / 引用 / 相似匹配）
        ↓
MVP0.5+ 候选：
  └─ AI Context Enhancement（基于稳定 operating context 的增强层）
  └─ GitHub 集成（数据主权闭合后的代码连接）
```

这个路线图保留了 DPv4pro 的长期视野，但把时机后移到 operating loop 验证之后。

---

## 7. 给后续正式 `/plan` 的输入

### 7.1 可直接采纳的稳定输入（五份一致）

1. Operating Review Loop 是第一优先级
2. Template Reuse 最小版边界（组合快照 + 预填）
3. Real-Project Dry-Run 独立验收
4. 不扩长期实体 / 不做重智能层 / 不做 GitHub OAuth
5. 守住 TECH_STACK_BASELINE 与已冻结工程约束

### 7.2 需仲裁的分歧输入

| 分歧 | 建议仲裁 |
| --- | --- |
| MVP0.3 版本语义 | 采纳四份评审：MVP0.3 = 完成 MVP0.2 阶段二 |
| 主线结构 | 采纳 GLM52 主从结构：Review Loop 唯一主线 + 两翼 + 验收闸 |
| Derived Intelligence 地位 | 采纳 GLM52：独立成线，MVP0.3 内完成最小深化 |
| Venture 时机 | 采纳四份评审：MVP0.3 不引入，留 MVP0.4+ |
| AI 时机 | 采纳四份评审：MVP0.3 不引入，留 MVP0.5+ |
| Decision Intelligence 时机 | 采纳四份评审：MVP0.3 不推进，留 MVP0.4+ |
| phase 命名 | 采纳不冻结原则，正式 `/plan` 建立时再定 |

### 7.3 正式 `/plan` 必须回答的问题

综合五份评审的建议，正式 `/plan` 至少回答：

1. **经营闭环**：daily / weekly review 的最小入口？review 如何消费 phase05 Feedback 与 phase06 ReuseSummaryRead？review 结果如何回流到 Decision / Product / Module？
2. **派生智能深化**：`capability_summary` 从计数到缺口提示的字段演进？`.proto` 向后兼容？复用机会推荐的最小算法？
3. **模板复用**：组合快照保存走哪个 application owner？预填与 draft-first / canonical create 语义如何对齐？
4. **真实项目验证**：选哪个真实项目？验收记录如何独立留存？哪些指标证明「更愿意回来使用」？
5. **工程约束**：review loop 写动作收敛到哪个切片？是否引入新实体（review 记录 vs 动作日志）？
6. **非目标**：不扩实体 / 不引第二套事实源 / 不做模板平台 / 不升格 AI / 不补 phase06 已收口主线

---

## 8. 对五份评审的总体评价

### 8.1 GPT54 — 战略清晰，结构偏松

GPT54 的战略判断最清晰：「不回头扩实体，不急于推重智能，优先推进为 Operating Review System」。三主线并列的弱点是焦点分散，但非目标清单（§8）非常到位。

### 8.2 GLM52 — 工程最实，主从最清

GLM52 基于一手实现经验，指出三处「接口存在、语义未闭合」隐患（§4.3），这是其他四份没有的独特贡献。主从结构最利于聚焦。Derived Intelligence 独立成线的主张有工程依据。

### 8.3 Qwen3.7-Pro — 执行最细，边界最保守

Qwen37pro 给出最具体的动作清单与度量指标表，执行可操作性最强。但 phase08 标注「可选」过于保守，可能导致 MVP0.2 阶段二无法收口。

### 8.4 DeepSeek-V4-Flash — 命名最准，风险意识最强

DPv4flash 的「Template-Driven Compound」命名最准确，「把一次性录入推进为周期性经营」的第一原则最精炼。对 Review Loop 演化为任务管理器的风险警示（§5 第 2 条）是五份中最强的风险意识。

### 8.5 DeepSeek-V4-Pro — 视野最远，时机最急

DPv4pro 的长期视野最远，正确识别了 Venture / Decision Intelligence / AI 三大长期方向，且给出了最详细的依赖关系论证。但时机判断过急——在 operating loop 未验证前推进这三项，会重演「理论领先、使用滞后」的问题。其文档更适合作为 MVP0.4+ 路线图，而非 MVP0.3 执行计划。

---

## 9. 最终结论

### 9.1 一句话汇总

> **五份评审在「Operating Review Loop 为第一优先级」上强一致；在「MVP0.3 版本语义与结构」上存在分歧；四份评审的「MVP0.3 = 完成 MVP0.2 阶段二」框架更稳健，DPv4pro 的「MVP0.3 = Venture + Decision Intel + AI」框架更适合作为 MVP0.4+ 路线图。**

### 9.2 仲裁结论

1. **采纳四份评审的版本语义**：MVP0.3 = 完成 MVP0.2 阶段二
2. **采纳 GLM52 的主从结构**：Operating Review Loop 唯一主线 + Derived Intelligence 左翼 + Template Reuse 右翼 + Dry-Run 验收闸
3. **采纳 DPv4flash 的命名与第一原则**：Template-Driven Compound、周期性经营
4. **采纳 Qwen37pro 的动作清单与度量指标**：StartDailyReview / RecordReviewConclusion / CreateActionFromReview、资产导入耗时必须落地
5. **采纳 GPT54 的非目标清单**：不扩实体 / 不做重智能 / 不做模板平台 / 不补 phase06 已收口主线
6. **后移 DPv4pro 的三大方向到 MVP0.4+**：Venture / Decision Intelligence / AI Context Enhancement

### 9.3 核心风险警示

五份评审共同警示的最大工程风险是：

> **Review Loop 演化为通用任务管理系统。**

review 动作必须锚定既有实体（Decision / Product / Module），绝不演化为独立 Kanban / Sprint / 任务面板。一旦演化为任务管理器，会稀释 `Module + Decision + Binding + Feedback` 的差异化核心，把 PSCO 拉回「通用项目管理工具」的赛道。

正式 `/plan` 必须在规格层冻结这条边界。

### 9.4 对后续工作的建议

1. **建立正式 phase 入口**：按 `project_rules.md §4.1` 先建立 `/plan` 三件套（architecture_plan / dev_plan / shared_baseline），不预设 phase 名称
2. **根级真相源切换**：phase 完成后同步更新 `AGENTS.md`、`plan.md`、`architecture_map.md`、`docs/README.md`
3. **规格正文互链**：正式规格必须明确直接承接 `phase06-12 / 13 / 16`，不回退重做已冻结交付物
4. **验收环境**：继续通过 `reset_*_acceptance.sh` + fixture 建立可重复复核环境，Review 流程必须有 fixture 覆盖（空状态 / 有数据 / 多轮 review）
5. **真实项目 dry-run**：独立留存验收记录，与 fixture 验收并列

---

*End of PSCO-mvp03-summarize-feedback-GLM52.md*
