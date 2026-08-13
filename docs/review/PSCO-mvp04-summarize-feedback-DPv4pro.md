# Personal Software Company OS

# MVP0.4 最终共识与规划基线

**Author:** DeepSeek-V4-Pro
**Date:** 2026-08-13
**Purpose:** 基于六份 `mvp0.4` 专家评审文档，形成 PSCO 在 `phase09` 收口与 `Real-Project Dry-Run` 反馈产出之后、进入下一阶段正式 `/plan` 之前的最终仲裁结论、范围基线、推进顺序与明确非目标。

---

## 1. 文档定位

本文不是第七份评审意见，也不是直接替代下一阶段正式 `phase` `/plan` 的执行文档。

本文要解决的是五件事：

1. 明确哪些判断已经形成稳定共识；
2. 对存在分歧的方向给出最终仲裁；
3. 冻结 `mvp0.4` 的正式方向、范围与非目标；
4. 约束下一阶段正式 `/plan` 的输入边界与推进顺序；
5. 为后续 `/plan -> /spec -> 实现 -> 验收 -> 收口` 提供唯一上游判断基线。

本文的职责，是为下一阶段正式 phase 入口建立提供上游依据，而不是在这里越权冻结正式 `phase` 名称、`.trae/specs/phaseXX_*` 路径、接口名或实现细节。

### 1.1 本文与 `PSCO-mvp03-summarize-feedback.md` 的关系

本文不是把 `mvp0.3` 重新写一遍，也不是把 `mvp0.3` 已经完成的 phase08/phase09 机械复制成另一份文档。

本文与 `PSCO-mvp03-summarize-feedback.md` 的关系，明确冻结为：

1. `PSCO-mvp03-summarize-feedback.md` 已经完成了从 `mvp0.2` 阶段二到 `mvp0.3` 正式主题的总方向仲裁，并给出了"Operating Review Loop 为中心主线 + Template Reuse / Derived Intelligence 为支撑 + Real-Project Dry-Run 为验收闸"的结构；
2. `phase08` 和 `phase09` 已经实际完成了其中的 Operating Review Loop、Template Reuse 与 Derived Intelligence Deepening；
3. `Real-Project Dry-Run` 已经产出 4 轮真实使用反馈与 8 点深度反思；
4. 本文要解决的，不是再次仲裁 `phase08/09` 做得对不对，而是基于 dry-run 真实反馈中暴露的问题，正式收敛 `mvp0.4` 的执行主题、范围基线与完整路线表达；
5. 因此，本文既有承接 `mvp0.3` 的部分，也有基于新证据新增冻结的部分；新增冻结的重点，不在于发明更多长期对象，而在于把"闭环修复 + 动作链路闭合 + Agent-Native 基础"从候选方向升级为下一步必须完成的正式主线。

换句话说：

> **`mvp0.3` 解决的是"从 Reuse-Aware Registry 走向 Operating Review System"；`mvp0.4` 解决的是"从 Operating Review System 走向被真实持续消费的 Connected Operating System"。**

---

## 2. 证据来源

### 2.1 第一轮方向评审（六份专家文档）

- `docs/review/PSCO-mvp04-DPv4flash.md` — DeepSeek-V4-Flash
- `docs/review/PSCO-mvp04-DPv4pro.md` — DeepSeek-V4-Pro
- `docs/review/PSCO-mvp04-gemini31pro.md` — Gemini-3.1-Pro
- `docs/review/PSCO-mvp04-GLM52.md` — GLM-5.2
- `docs/review/PSCO-mvp04-GPT54.md` — GPT-5.4
- `docs/review/PSCO-mvp04-qwen37pro.md` — Qwen-3.7-Pro

### 2.2 本轮唯一核心证据

- `docs/review/PSCO-real-project-dry-run-user-manual-GPT54 feedback.md` — 4 轮真实使用反馈 + 8 点深度反思

### 2.3 本文采用的仲裁原则

1. **高共识优先于单点激进扩张。** 四方以上独立收敛的判断，优先视为稳定输入。
2. **真实使用证据优先于理论推演。** Dry-run 反馈揭示的问题比任何专家评审都更接近系统真实状态。
3. **闭合优先于扩张。** 在当前已交付能力尚未形成真正闭环前，继续扩对象宽度会加重"中间地带"困境。
4. **分层处理优先于整体扩张。** 反馈中的问题横跨 bug、设计缺陷、产品定位、长期方向四个层次，不能混为一谈。
5. **时机优先于愿景。** 某个长期方向即使正确，只要当前前提未成立，就不能提前升格为下一阶段主线。
6. **评审不越权冻结正式 phase。** 正式阶段命名、spec 路径与具体落点留给后续 `/plan` 与 `/spec` 决定。

---

## 3. 最终共识

经过六份专家文档的独立评审，以下内容已经形成稳定共识，可直接视为 `mvp0.4` 的正式前提。

### 3.1 关于 dry-run 反馈价值的共识

1. `phase01 ~ phase09` 已经证明 PSCO 的最小资产主线、经营回路与支撑能力都成立，且系统结构克制、链路清晰、具有真实产品潜力。
2. `Real-Project Dry-Run` 完全达成了它的设计目的——不是证明 PSCO 完美，而是真实暴露了 fixture 验收无法发现的"实体登记完备"与"经营动作闭环"之间的语义断层。
3. 反馈揭示的不是"功能缺失"，而是"语义断层"：实体登记能力完备，双向关联链路精良，但实体之间缺乏"动作因果链"。
4. Dry-run 反馈不是失败信号，而是 dry-run 设计意图的真正达成。

**六份文档全部同意以上判断。** 这是本轮最稳固的共识基础。

### 3.2 关于下一步总方向的共识

1. 下一步不应继续扩长期对象宽度（不引入 `Venture`、`Opportunity`、`Experiment` 等新实体作为主线）。
2. 下一步不应提前切到 `Decision Intelligence` 或 `AI Context Enhancement`。
3. 下一步应优先闭合"实体登记"与"经营动作"之间的语义断层。
4. 下一步的核心任务是从"可登记、可复用、可看见"推进到"可被真实经营动作持续消费"。

**五份文档同意（DPv4pro、DPv4flash、GLM52、GPT54、Qwen37pro），Gemini 持部分保留（认为应同时引入 Venture/Experiment）。**

### 3.3 关于必须立即修复的阻断项的共识

1. Onboarding "开始首轮录入"按钮空态无响应（`fix_001`）。
2. Decision 状态语义错位——`proposed` 同时承担"已留痕"与"待处理"双重语义，导致 Dashboard / Daily Review 误标（`fix_002`）。
3. Decision Detail 缺少状态推进 CTA，review loop 在 Decision 处断裂（`fix_003`）。

**六份文档全部同意这三项必须立即修复，且应走 `fix*` workflow 而非 phase 级工作。**

### 3.4 关于范围边界的共识

1. 不新增长期核心实体主线（`Venture`、`Opportunity`、`Experiment` 继续后移）。
2. `Capability` 继续作为派生层，不进入重实体 CRUD。
3. `Decision` 必须继续保持经营中心地位，不能被弱化。
4. 不做 AI 一级工作台、不做 AI Chat 界面、不做自动扫描 / 知识图谱。
5. 不做 GitHub OAuth / 自动导入、不做 Rust Intelligence Layer。
6. 技术栈、`.proto` 单一合同源、前端 query / application 边界等既有工程约束继续冻结。

**五份文档同意以上边界（DPv4pro、DPv4flash、GLM52、GPT54、Qwen37pro），Gemini 持不同意见（认为应引入 Venture/Experiment 实体）。**

### 3.5 关于 Agent 协同方向的共识

1. PSCO 未来应支持 agent 可编程访问，但当前阶段不应演化为 AI 工作台。
2. Agent 消费层应以**只读上下文消费**为主，不做内置 AI 对话、自动决策或代码生成。
3. PSCO 的 `.proto + ConnectRPC` 主线已经具备机器可消费能力，缺的是 agent 友好的上下文输出层。

**五份文档同意以上判断（DPv4pro、DPv4flash、GLM52、GPT54、Qwen37pro），Gemini 持更激进立场（认为 agent 应能写入 PSCO）。**

---

## 4. 最终仲裁结论

### 4.1 `mvp0.4` 的版本语义

**最终结论：**

> `mvp0.4` 的实质任务，不是跳过 `mvp0.3` dry-run 暴露的语义断层去开启一套更大的长期模型；
> 而是基于 dry-run 真实反馈，先闭合阻断项完成 `mvp0.3` 正式收口，再把"动作链路闭合 + Agent-Native 基础"真正做完，并将其作为 PSCO 从"Operating Review System"走向"Connected Operating System"的正式跨越。

这意味着我**不采纳**把 `mvp0.4` 建立在"先默认 dry-run 阻断项已经修复"或"直接跳到 Venture/Opportunity/Experiment"这一前提上的推进方式。

原因：

1. Dry-run 暴露的 3 个阻断项是真实使用中发现的，必须修复后才能谈"下一阶段"；
2. "实体登记完备"与"经营动作闭环"之间的语义断层是 dry-run 最核心的发现，必须在 mvp0.4 中闭合；
3. 在这一步未完成前，直接上跳到 `Venture / Decision Intelligence / AI Context Enhancement / Opportunity / Experiment`，会把系统带回"概念扩张先于使用闭环"的旧问题。

进一步冻结为：

> **`mvp0.4` 不是对 `mvp0.3` 的否定，也不是平行新路线；它是 `mvp0.3` dry-run 验收之后的正式闭环化、连接化与 Agent-Native 基础化。**

### 4.2 `mvp0.4` 的总主题

**最终结论：**

> `mvp0.4` 的总主题确定为：
> **从 Operating Review System 走向 Connected Operating System。**

它的正式主轴不是三条并列新世界，也不是提前切到战略与智能主线，而是：

1. **Asset-Action Closure**（中心主线）
2. **Agent-Native Foundation**（支撑能力，受控最小版）
3. **Cross-Project Convention Asset**（候选探索，严格最小版或后移）

### 4.3 对"三并列主线"与"一主一翼一候选"的仲裁

**最终结论：**

> 我采纳 **"Asset-Action Closure 为中心主线，Agent-Native Foundation 为支撑能力，Cross-Project Convention Asset 为候选探索"** 的结构。

原因如下：

1. Asset-Action Closure 是消费 dry-run 反馈中暴露的语义断层的唯一正确方式——它让每一次创建、每一次 review、每一次决策都自然形成实体间因果链；
2. Agent-Native Foundation 脱离 Asset-Action Closure，容易退化为"为 AI 而 AI"的孤立功能——只有先闭合了资产到动作的链路，agent 消费才有可消费的资产底座；
3. Cross-Project Convention Asset 脱离前两者，容易演化为知识图谱或文档管理系统——必须先闭合核心动作链路，再考虑跨项目资产复用。

因此我**不采纳**把 Agent-Native / Cross-Project Convention 与 Asset-Action Closure 完全并列推进的结构，也**不采纳**把 Venture / Opportunity / Experiment 直接升格为 mvp0.4 主线的提议。

### 4.4 对 Gemini 激进路线的正式仲裁

Gemini-3.1-Pro 提出了三条主线：Agent-Driven Workflow（读写倒置）、跨项目全局上下文资产化、融入精益创业模型（引入 Venture/Opportunity/Experiment 实体）。

**仲裁结论：不采纳作为 mvp0.4 主线。**

理由：

1. **Agent-Driven Workflow（读写倒置）**：方向正确，但时机过早。当前 PSCO 自身的资产-动作链路尚未闭合，先让 agent 写入未闭合的资产会放大混乱。采纳其中的"只读上下文消费"思想，但明确后移"agent 写入"到 mvp0.5+。
2. **跨项目全局上下文资产化**：方向正确，部分采纳。将其降级为 mvp0.4 的候选探索（Cross-Project Convention Asset），严格最小版，不作为主线。
3. **融入精益创业模型（Venture/Opportunity/Experiment）**：方向正确，但时机严重过早。在 PSCO 尚未证明"现有实体能被持续消费"之前引入新实体类型，会加重"概念扩张先于使用闭环"的问题。明确后移到 mvp0.5+。

### 4.5 对 Qwen 科学实验 Workflow 的正式仲裁

Qwen-3.7-Pro 提出了"闭环修复 + 协调方式升级 + 科学实验 workflow 最小承接"的主线，其中方向 C 建议引入 Opportunity/Experiment 实体。

**仲裁结论：方向 A 和 B 采纳，方向 C 后移。**

理由：

1. 方向 A（闭环修复）与方向 B（协调方式升级）与共识高度一致，直接采纳；
2. 方向 C（科学实验 workflow 最小承接）涉及新增长期核心实体，在 mvp0.4 阶段前置条件不成熟；
3. 用户反思 5 中表达的"科学实验 workflow"诉求是真实的，但应作为 mvp0.5+ 的核心方向被明确记录，而非提前压入 mvp0.4。

### 4.6 Asset-Action Closure 的正式定义

**最终结论：**

> Asset-Action Closure 的正式范围冻结为以下五项，不得扩展：

1. **Onboarding 逻辑化**：创建时主动建立实体关系（产品→仓库→模块→决策），而非事后手动绑定
2. **Decision 状态机细化**：`draft → proposed → acknowledged → active → superseded → archived`，让 `proposed` 不再承担双重语义
3. **Cross-Entity CTA 矩阵**：每个实体详情页都主动提示"下一步应该做什么"
4. **Review Loop 业务闭环**：review → Decision → 实体更新 → review 已完成状态
5. **Current Focus 信号派生规则重做**：基于新状态机，只指向真正待处理的 Decision

**明确非目标**：Onboarding 逻辑化不能演化为可配置工作流引擎；Decision 状态机细化不能演化为 Decision Intelligence；Cross-Entity CTA 不能演化为推荐系统。

### 4.7 Agent-Native Foundation 的正式定义

**最终结论：**

> Agent-Native Foundation 的正式范围冻结为以下三项，不得扩展：

1. **Agent Context Query API**：基于既有 ConnectRPC，输入仓库 URL 返回相关资产、决策和模块的上下文
2. **AGENTS.md 风格上下文导出**：把 Product / Module / Decision / Repository 组合导出为 agent 可读上下文
3. **上下文快照机制**：不演化为完整模板平台

**明确非目标**：不做 PSCO 内置 AI 工作台；不做自动代码生成或修改；不做 agent 身份认证或权限治理；不做 agent 双向同步（只读消费即可）；不做 agent 写入 PSCO。

### 4.8 Cross-Project Convention Asset 的正式定位

**最终结论：**

> Cross-Project Convention Asset 作为 mvp0.4 的候选探索，严格最小版或后移到 mvp0.5+。

**如果进入，必做范围（严格最小版）**：

- 新增 `Convention` 实体（最小字段：title / content / tags）
- 支持跨项目引用（Product → Convention 的轻量绑定）
- 支持从既有项目导出规范为 Convention 资产

**明确非目标**：不做规范版本治理；不做自动同步或模板生成；不做规范间依赖关系图谱；不做语义搜索或知识图谱。

### 4.9 后移项清单（继续冻结）

以下方向在 `mvp0.4` 中继续明确不做，后移到 `mvp0.5+` 候选范围：

1. `Venture` 正式主线化
2. `Opportunity / Feature / Experiment` 流程化
3. `Decision Intelligence`（相似匹配、引用推荐、AI 上下文增强）
4. `AI Context Enhancement`（决策草稿辅助、周报总结、AI 模块复用建议）
5. `Capability` 重实体 CRUD
6. GitHub OAuth / 自动导入
7. AI 一级工作台 / AI Chat 界面
8. 自动扫描 / 知识图谱 / 语义搜索
9. Rust Intelligence Layer
10. 完整模板平台 / 模板版本管理
11. Agent 写入 PSCO / Agent 双向同步
12. 科学实验 Workflow（Opportunity/Experiment 实体引入）

---

## 5. 推进顺序与依赖关系

### 5.1 推荐的推进路径

```
第一步：fix* 闭合 3 个阻断项
  ├─ fix_001_onboarding_cold_start_state
  ├─ fix_002_decision_status_semantics
  └─ fix_003_decision_detail_status_advance_cta
        ↓
第二步：重跑一轮 dry-run 验证 fix 效果
  └─ 只验证 fix 是否解决了阻断项，不要求验证第二层、第三层问题
        ↓
第三步：dry-run 判为"通过（含已知摩擦点）"，正式收口 mvp0.3
  └─ 已知摩擦点记入收口报告，8 点反思作为 mvp0.4 仲裁上游留档
        ↓
第四步：建立 mvp0.4 正式 /plan 入口
  └─ 直接承接 dry-run 反馈、fix 结论、8 点反思、本文仲裁结论
```

### 5.2 推荐的候选阶段结构（不冻结正式 phase 命名）

```
候选阶段一：dry-run 阻断项 fix 与 mvp0.3 收口
  └─ 闭合 3 个阻断项 → 重跑 dry-run → mvp0.3 正式收口
        ↓
候选阶段二：Asset-Action Closure 主线
  └─ onboarding 逻辑化 + Decision 状态机细化 + Cross-Entity CTA 矩阵
     + review loop 业务闭环 + Current Focus 信号派生规则重做
        ↓
候选阶段三：Agent-Native Foundation 支撑能力
  └─ Agent Context Query API + AGENTS.md 风格上下文导出 + 上下文快照机制
        ↓
候选阶段四（可选）：Cross-Project Convention Asset 候选探索
  └─ 严格最小版，视前置阶段完成情况决定是否进入
```

### 5.3 推荐的先后关系

1. 没有候选阶段一，dry-run 无法收口，mvp0.3 不能正式完成；
2. 没有候选阶段二，PSCO 仍停留在"实体登记完备但动作链路断裂"状态；
3. 没有候选阶段三，PSCO 资产无法被外部 agent 消费，用户反思中的"工程化配合"无法落地；
4. 候选阶段四是可选项，不阻断 mvp0.4 主线收口。

### 5.4 对正式 `/spec` 的硬约束

后续 `/spec` 与实现必须继续遵守：

1. `.proto` 仍是唯一长期合同源
2. query 层继续保持纯只读
3. 前端写路径继续收敛到切片固定承接位
4. 不引入第二套路由、第二套状态管理、第二套 UI 事实源
5. Onboarding 逻辑化不能演化为可配置工作流引擎（仍是 6 步引导，关系建立是主动询问而非强制依赖）
6. Agent-Native Foundation 不能演化为 AI 工作台（只读消费，不做内置对话/自动决策/代码生成）
7. Cross-Project Convention Asset 不能演化为知识图谱（只登记工程可消费的规范，不做语义搜索/笔记/依赖图谱）
8. Decision 状态机细化不能演化为 Decision Intelligence（只细化到区分已留痕/已确认/已生效，不做相似匹配/引用推荐/AI 增强）
9. Review loop 不能演化为通用任务管理器（review 动作必须锚定既有实体）

### 5.5 对后续验收的硬要求

下一阶段正式验收至少必须回答以下问题：

1. 用户从空态 Dashboard 进入 Onboarding，能否一次性形成"产品 + 仓库 + 模块 + 决策"的完整关系网（而非分别创建后手动绑定）？
2. Decision Detail 能否完成 `proposed → acknowledged → active` 的状态推进，并回流到 review？
3. 每个实体详情页是否都主动提示"下一步应该做什么"？
4. Dashboard Current Focus 是否只指向真正待处理的 Decision？
5. 外部 agent 能否通过标准化接口读取 Product 完整上下文（含关联 Module / Decision / Repository）？
6. 真实项目中，用户是否愿意围绕 PSCO 持续回来使用，且使用摩擦显著低于 mvp0.3 dry-run？

---

## 6. 风险与缓解

| 风险 | 级别 | 缓解措施 |
| --- | --- | --- |
| Onboarding 逻辑化被过度工程化（演变为工作流引擎） | 🟡 中 | 严格限定为"创建时主动询问建立关系"，不做可配置流程、不做模板 |
| Decision 状态机细化被偷渡为 Decision Intelligence | 🟡 中 | 只做 6 状态区分语义，不做相似匹配、引用推荐、AI 增强 |
| Agent-Native Foundation 被偷渡为 AI 工作台 | 🔴 高 | 严格限定为只读上下文消费，不做内置对话、自动决策、代码生成 |
| Cross-Project Convention Asset 被偷渡为知识图谱 | 🟡 中 | 严格限定为工程可消费的规范资产，不做笔记、语义搜索、依赖图谱 |
| Dry-run 收口被无限延期 | 🟡 中 | fix 必须保持轻量（1-2 周内闭合），不能让阻断项演化为 phase 级工作 |
| Agent 消费层过早优化（在实际使用前做了太多假设） | 🟢 低 | 只做最小 API 和上下文格式定义，不做 agent 运行时、调度或编排 |

---

## 7. 最终结论

`phase01 ~ phase09` 已经证明 PSCO 的最小资产主线、经营回路与支撑能力都成立。`Real-Project Dry-Run` 也完成了它的设计目的——真实暴露了 fixture 验收无法发现的"实体登记完备"与"经营动作闭环"之间的语义断层。

因此，`mvp0.4` 的最终仲裁结论是：

> **下一步不应继续扩长期对象宽度，也不应直接跳到 Venture / Decision Intelligence / AI Context Enhancement / Opportunity / Experiment；而应优先闭合"实体登记"与"经营动作"之间的语义断层，让 PSCO 从"可登记、可复用、可看见"推进到"可被真实经营动作持续消费"的 Connected Operating System。同时以受控最小版建立 Agent-Native Foundation，让外部开发工具可消费 PSCO 资产上下文，桥接"经营层"与"工程层"。**

正式收敛方向如下：

1. **Asset-Action Closure** 为中心主线；
2. **Agent-Native Foundation**（受控最小版，只读消费）为支撑能力；
3. **Cross-Project Convention Asset** 为候选探索（严格最小版或后移）；
4. `Venture / Decision Intelligence / AI Context Enhancement / Opportunity / Feature / Experiment / 完整模板平台 / GitHub OAuth / 自动扫描 / Agent 写入` 继续后移到 `mvp0.5+` 候选范围。

### 对六份专家文档的核心分歧仲裁

| 分歧点 | 最终仲裁 | 理由 |
| --- | --- | --- |
| 是否引入 Venture/Opportunity/Experiment 实体（Gemini、Qwen 主张） | **不采纳**，后移到 mvp0.5+ | 在语义断层未闭合前引入新实体，会加重"概念扩张先于使用闭环" |
| Agent 是否应写入 PSCO（Gemini 主张） | **不采纳**，只读消费即可 | 当前资产-动作链路未闭合，agent 写入会放大混乱 |
| Agent Consumption Layer 范围（各文档略有差异） | **冻结为只读上下文消费**，不做写入 | 五份文档共识，Gemini 的写入立场不采纳 |
| Cross-Project Convention Asset 优先级 | **降级为候选探索**，可后移 | 价值真实但风险大，不阻塞主线收口 |

如果用一句话收口：

> **`phase01 ~ phase09` 证明了 PSCO"能登记、能复用、能看见复利"；`dry-run` 证明了 PSCO"在真实使用中暴露了语义断层"；`mvp0.4` 必须闭合这层断层，让 PSCO 从"Operating Review System"真正推进为"可被持续消费的 Connected Operating System"。**

---

*End of PSCO-mvp04-summarize-feedback-DPv4pro.md*