# Personal Software Company OS

# MVP0.2 Next-Stage Direction — Summarize & Feedback

**Author:** DeepSeek-V4-Flash
**Date:** 2026-08-09
**Role:** 对`docs/review/`下五份 MVP0.2 推进方向评审文档（GPT54 / GLM-5.2 / DeepSeek-V4-Pro / DeepSeek-V4-Flash / Qwen3.7-Pro）做交叉总结，明确本人支持与反对的观点及理由，供后续正式 `/plan` 参考引用
**Document Type:** `review`
**Status:** 交叉共识参考，不构成正式 phase 命名或执行指令

---

## 1. 文档定位与总结方法

本文不做"五份文档观点罗列"，而是以 DeepSeek-V4-Flash 的专业判断，对这五份独立评审做一次**交叉仲裁性质的总结**：

1. 先给出五份文档的立场格局（收敛区与分歧区）；
2. 明确我**支持**的共识与个体观点；
3. 明确我**反对**的观点（主要是方向优先级上的分歧）；
4. 每条都给出基于工程现实与 PSCO 长期愿景的理由；
5. 最后给出一个可执行的合成方向建议。

**透明声明**：我本人（DeepSeek-V4-Flash）是 Camp 1（Operating Loop）的作者之一。因此我在第 3、4 节会刻意用"工程事实"而非"维护自己文档"来论证，避免既当运动员又当裁判的偏差。

---

## 2. 五份文档的立场格局

### 2.1 共识区（五份文档高度一致，可视为强收敛）

| 判断 | GPT54 | DS-Flash | DS-Pro | GLM-5.2 | Qwen3.7-Pro |
| --- | :-: | :-: | :-: | :-: | :-: |
| 下一阶段不扩展实体宽度（不批量引入 Opportunity/Venture/Feature/Experiment） | ✅ | ✅ | ✅ | ✅ | ✅ |
| `Capability` 继续作为派生层，不建重表 | ✅ | ✅ | ✅ | ✅ | ✅ |
| 不做 GitHub OAuth / 重量级自动导入（至少非主线） | ✅ | ✅ | ✅ | ✅ | ✅ |
| 不做 AI 一级工作台 / 自动扫描 / 知识图谱 / Rust 智能层 | ✅ | ✅ | ✅ | ✅ | ✅ |
| 技术栈继续冻结，单一 `.proto` 合同源，不引第二套基础设施 | ✅ | ✅ | ✅ | ✅ | ✅ |
| **导出 / 备份（Local First 数据主权）必须在新阶段闭合** | ✅ | ✅ | ✅ | ✅ | ✅ |

**关键洞察**：五份独立文档在"不做什么"上几乎完全一致，在"导出/备份必须闭合"上也完全一致。这说明这些不是某个模型的偏好，而是基于同一组 MVP0.1 现实与原始方案的**客观事实约束**。

### 2.2 分歧区（唯一真正的分歧：主线优先级）

分歧不在"要不要做复用感知 / 要不要做 review"，而在**先做什么**：

- **Camp 1（Operating Loop Foundation）**：GPT54 / DeepSeek-V4-Flash / DeepSeek-V4-Pro
  - 顺序：**Onboarding → Review Loop → Derived Intelligence**
  - 主张：先解决"数据怎么进入 + 用户凭什么每天用"，复用感知是派生产物，放在最后。
- **Camp 2（Reuse / Capability 复利）**：GLM-5.2 / Qwen3.7-Pro
  - 顺序：**复用感知 + 模板级复用 → 导出收尾**
  - 主张：`module_reuse_summary` / `capability_summary` 是让"复利可见"的关键，应作为主体；锚点是 `PSCO-mvp0.1-summarize-feedback.md §4.9`"模板级复用是 v0.2 第一优先级"。

下面分别给出我的支持与反对。

---

## 3. 我支持的共识与个体观点

### 3.1 支持"不扩实体、不引重集成、不让 AI 抢判断权"的全局共识

理由不需展开——这是五份独立文档基于同一组事实的收敛，且完全符合 `PSCO_0 ~ PSCO_4` 的长期边界（AI 是增强层、绝不做自动扫描/知识图谱）。**Camp 2 与 Camp 1 在这一层没有任何冲突，我全盘支持。**

### 3.2 支持"导出 / 备份必须在下一阶段闭合"（此点所有专家一致，且我首先提出）

我在五份里最先以"规格合规负债"的形式指出：`mvp_spec_v0.1.md §7.3/§7.4` 明确要求导出/备份，但后端代码中**无对应实现**。GLM / Qwen 将其列为"方向 C 收尾"，V4-Pro 给出了具体技术路径（JSON 导出端点 + pg_dump 备份脚本）。

**我支持把它作为贯穿性硬约束，而非可选收尾**，理由：
- 它直接违背"数据所有权优先（Local First）"这一 PSCO 差异化底座；
- `phase01~05` 已积累真实资产数据，越晚闭合，用户越难信任系统；
- 实现成本低（JSON 序列化 + pg_dump），没有理由继续延后。

### 3.3 支持 V4-Pro 的"2026-08-04 评审 9 条建议回看"方法，及其结论

V4-Pro 回看发现：4 条 🔴 高优先级建议中，**冷启动 与 数据可移植性 两条在 MVP0.1 后仍未闭合**。这与我独立发现的"登记与查看≠日常经营 + 导出缺失"高度互证，构成对 MVP0.2 起点判断的强验证。**我支持把"冷启动"作为最优先补的缺口之一。**

### 3.4 支持 Qwen3.7-Pro 的两个补充

1. **Decision 复用后移**：`Decision` 复用在 mvp0.1 已能完成基础闭环，其上下文匹配逻辑依赖复用感知成熟，应后移 mvp0.3。**我支持。**
2. **真实项目 dry-run 作为独立交付物**：把 dry-run 从"phase07 的一部分"提升为独立子阶段，避免被功能实现挤压。**我支持**——这符合项目"他人可稳定复验"的验收纪律。

---

## 4. 我反对的观点（集中在主线优先级）

### 4.1 反对 Camp 2 把"复用感知（derived views）"作为 MVP0.2 的主体/第一优先级

这是本文最核心的反对意见。理由如下：

**理由一：派生指标在数据密度不足时没有可信价值。**
`module_reuse_summary` / `capability_summary` 是纯 SQL 聚合。MVP0.1 是个人级数据量（零星 Product / Module）。在数据稀疏、且用户尚未形成使用习惯时：
- 复用率大概率趋近 0，"复利可见"呈现为"空指标"，提供不了洞察；
- 反而可能让 Dashboard 变成一个"数据很少的漂亮图表"，这正是 summarize-feedback 警告的"台账工具化"预先温床。

**理由二：复用感知是"输出"，持续使用是"输入"；先做输出、不做输入，飞轮依然转不起来。**
复利的前提是先有"持续的数据进入 + 持续的日常使用"。Camp 2 的方案里，`Onboarding`（让现实资产进入）、`Review Loop`（让用户每周回来）是没有被当作主体的。**做一堆只读聚合视图，无法提高使用频率，也无法填补冷启动缺口。** 这恰恰是 MVP0.1 与"运行系统"之间最真实的距离。

**理由三：Camp 2 的"方向 A"本质上是我的"Derived Intelligence"这一条线的内容，只是被放到了最前面。**
我并非反对 derived views 本身——我支持在 Derived Intelligence 阶段完整落地 `module_reuse_summary` / `capability_summary` / 复用率指标。我反对的是**排序**：把派生输出放在"数据进入 + 使用习惯"之前。

> 一句话：**Camp 2 的内容我全收，Camp 2 的顺序我反对。**

### 4.2 对 Camp 2 引用的"§4.9 模板级复用是 v0.2 第一优先级"的重新校准

GLM / Qwen 以 `PSCO-mvp0.1-summarize-feedback.md §4.9`"模板级复用是 v0.2 第一优先级"作为龙头锚点。我对此做两点校准：

1. **模板级复用本质属于 Onboarding / 快速创造，而非孤立主线。** "从已有 Product 的 Module 组合快速创建新产品"解决的是"新建太慢"的录入摩擦，应归入 Cold Start / Onboarding 的"快速创造"能力，而不是单独作为复用感知主体。
2. **"第一优先级"不等于"第一阶段"。** 它在共识文档里是 v0.2 的优先级排序，但把它作为 MVP0.2 的第一块落地，会再次把"录入/创建"放在"日常使用"之前——仍是优先级次序问题，而非是否要做的问题。

### 4.3 反对"先做复利可见、再谈使用频率"这一隐含排序

Camp 2 的收口话术"不扩实体，先做复利；不引集成，先做导出"在字面上成立，但隐含的假设"复利可见能带来使用"是未被验证的。工程上更稳的因果链是：

> 降低进入门槛（Onboarding）+ 建立使用习惯（Review Loop）→ 数据密度与真实性上升 → 复利指标才有可信度（Derived Intelligence）。

**先做使用，复利才会出现；先做复利图表，只会得到空指标。**

---

## 5. 我对五份文档的个体采纳/反驳汇总

| 专家 | 核心主张 | 我的态度 | 理由 |
| --- | --- | --- | --- |
| GPT54 | Operating Loop Foundation（Onboarding / Review / Derived） | **支持（方向一致）** | 与我的工程判断一致：先解决使用频率与回流闭环 |
| DeepSeek-V4-Pro | 同 Operating Loop + 导出负债 + 9 条建议回看 | **支持** | 回看方法强化了"冷启动 + 数据可移植性"是持久缺口的判断 |
| DeepSeek-V4-Flash（我） | Operating Loop + 导出合规负债 | （自我立场） | 引出"先还负债、再做运营层"的工程纪律 |
| GLM-5.2 | 复用感知为主体 + 模板复用 + 导出收尾 | **反对主体优先级，支持内容** | derived views 是输出不是输入；应先 Onboarding/Review |
| Qwen3.7-Pro | 复用感知为主体 + 模板复用 + 导出收尾 + dry-run 独立 | **反对主体优先级，支持内容与补充** | 同 GLM；特别支持 dry-run 独立交付物与 Decision 复用后移 |

---

## 6. 我的合成方向建议（供后续 `/plan` 引用）

**结论：MVP0.2 主题采用 Camp 1 的 Operating Loop Foundation，但把 Camp 2 的复用/复利内容完整纳入 Derived Intelligence 阶段，并按以下顺序落地。**

```
Block 1: Cold Start / Onboarding + Export/Backup 闭合
  ├─ 首次引导、draft-first/minimum-field、辅助导入、快速创建（模板级复用最小版归入此处）
  ├─ 一键 JSON 导出 + pg_dump 备份（闭合 Local First 负债）
  └─ DoD：新用户一次会话完成 Product + Repository + Module + Decision
        ↓
Block 2: Review / Operating Loop
  ├─ Daily / Weekly Review、Action handoff（锚定既有实体）、Feedback→Decision→Update 闭环
  └─ DoD：用户可完成完整 weekly review 并回流实体
        ↓
Block 3: Derived Asset Intelligence
  ├─ module_reuse_summary / capability_summary、复用率+导入耗时指标、候选能力提示
  ├─ 真实项目 dry-run（独立交付物）
  └─ DoD：复利指标可观测，不引入重 Capability 写入流程
```

**要点：**
1. Camp 2 的 `module_reuse_summary` / `capability_summary` / 模板级复用 / 导出，**全部保留**，只是移动了位置；
2. 复用感知不再是"第一块"，而是"在数据密度与使用习惯建立之后"的派生产出；
3. 这条路径同时满足以 `summarize-feedback` 为真相源的合规要求，不引入任何新实体、新集成、新基础设施。

---

## 7. 最终结论

五份独立评审在"不做什么"与"导出必须闭合"上完全一致，在"用什么技术"上完全一致。**唯一真实分歧是主线先后**。

我的立场明确：

- **支持**：不扩实体、不引重集成、AI 仅增强、`Capability` 派生化、技术栈冻结、导出/备份闭合、dry-run 独立交付、Decision 复用后移。
- **反对**：把复用感知（derived views / 模板复用）作为 MVP0.2 的第一块主体。因为派生指标是"输出"，而"数据进入 + 日常使用"才是"输入"；先做输出只会得到空指标，无法让飞轮转起来。
- **既定方向**：采用 Operating Loop Foundation（Onboarding → Review → Derived），并把 Camp 2 的复用/复利内容完整纳入第三阶段。

用一句话收口我对五份文档的仲裁：

> **Camp 2 的"复利内容"我全要，Camp 2 的"复利优先"我反对；先让数据进来、让人每周回来，复利才会从空指标变成真事实。**

---

*End of PSCO-mvp02-summarize-feedback-deepseekv4flash.md*