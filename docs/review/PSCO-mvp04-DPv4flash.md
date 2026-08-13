# Personal Software Company OS

# MVP0.4 Real-Project Dry-Run 反馈评审与方向展望

**Author:** DeepSeek-V4-Flash
**Date:** 2026-08-13
**Role:** 基于 `phase01 ~ phase09` 已收口现实、`PSCO-real-project-dry-run-user-manual-GPT54 feedback.md` 四轮真实使用反馈与 8 点深度反思，以 DeepSeek-V4-Flash 专业能力陈述 PSCO 下一步推进方向与展望
**Document Type:** `review`
**Status:** 供后续专家讨论、交叉评审与正式 `/plan` 参考，不直接冻结正式 phase 名称、spec 路径或实现指令

---

## 1. 文档定位与结论先行

本轮评审与前三轮（`mvp01 / mvp02 / mvp03`）有本质不同：这是 PSCO 项目历史上**第一次基于真实使用证据**（而非 fixture 验收与理论推演）进行的方向评审。

我要回答的核心问题是：

> 在 `phase09` 已正式收口、首次 `Real-Project Dry-Run` 已产出 4 轮反馈与 8 点反思之后，这些证据揭示了什么、应如何分层处理、以及 PSCO 下一阶段最符合长期复利价值的方向是什么？

**我的结论先写在前面：**

> **`dry-run` 反馈不是失败信号，而是 `dry-run` 设计意图的真正达成——它真实暴露了 fixture 验收无法发现的「实体登记完备」与「经营动作闭环」之间的语义断层。下一阶段不应继续扩长期对象宽度，也不应提前切到 Venture / Decision Intelligence / AI Context Enhancement；而应优先闭合这层断层，让 PSCO 从「可登记、可复用、可看见」推进到「可被真实经营动作持续消费」。**

这与我此前在 `PSCO-mvp03-DPv4flash.md` 中的立场一脉相承：

- mvp03 我主张「先做使用与回流闭环，复利才会从空指标变成真事实」；
- mvp04 我进一步主张「先闭合资产到动作的链路，agent 连接与工程化加速才有可消费的资产底座」。

一句话收口 mvp04 的展望：

> **`phase01 ~ phase09` 证明了 PSCO「能登记、能复用、能看见复利」；dry-run 证明了 PSCO「在真实使用中暴露了语义断层」；mvp04 必须闭合这层断层，让 PSCO 从「资产登记系统」真正推进为「可被持续消费的个人软件公司 operating system」。**

---

## 2. 评审依据

- 使用手册：`docs/review/PSCO-real-project-dry-run-user-manual-GPT54.md`
- **真实反馈（核心证据）**：`docs/review/PSCO-real-project-dry-run-user-manual-GPT54 feedback.md`（4 轮 feeback + 8 点深度反思）
- 当前推进状态：`plan.md`、`AGENTS.md`、`project_rules.md`、`architecture_map.md`
- mvp03 仲裁基线：`PSCO-mvp03-summarize-feedback.md`
- 两个最近完成 phase 的验收：`phase08-11` 与 `phase09-11` acceptance_report
- 原始长期方案：`PSCO_0.md ~ PSCO_4.md`
- 我核对了当前代码以实现（前端 `features/*`、后端 `internal/*`、`proto/psco/*`），把反馈中的三个阻断项落到代码根因

---

## 3. 对 dry-run 反馈的整体判断

### 3.1 dry-run 达成了它的设计目的

`PSCO-mvp03-summarize-feedback.md §4.8` 明确把 `Real-Project Dry-Run` 定位为 mvp0.3 后段验收的**独立交付要求**，要补三件事：

1. 真实使用摩擦是否可接受；
2. review loop 是否真的改变下一步动作；
3. 模板复用是否真的缩短下一次创造路径。

按这个标尺回看 4 轮反馈，结论非常清楚——**三项都得到了真实证据**：

| 要验证的事 | 反馈证据 | 结论 |
| --- | --- | --- |
| 真实使用摩擦 | feeback_01/02 冷启动断裂、feeback_03 手动关联摩擦、feeback_04 闭环断裂 | 摩擦被真实暴露 |
| review loop 是否改变动作 | feeback_04「进入 Decision Detail 无处理入口，未形成闭环」 | 断裂点被精确定位 |
| 模板复用是否缩短路径 | 用户尚未深度走通 B 线；但已确认 onboarding 实体并列隔离影响「复用」心智建立 | 未能验证，需修复后重验 |

换句话说：**dry-run 不是证明 PSCO 完美，而是证明 PSCO 在真实使用中暴露了 fixture 验收无法发现的问题。从这个角度看，本轮 dry-run 完全达成设计目的。**

### 3.2 反馈揭示的不是「功能缺失」，而是「语义断层」

最关键的洞察在用户反思第二点：

> 「onboarding 阶段看似六个流程，其实就是分别新增实体，不存在任何逻辑关系。」

这句话精准命中 PSCO 当前设计的核心问题：

- **实体登记能力完备**：Product / Module / Decision / Repository 都能创建、编辑、关联；
- **双向关联链路精良**：来源标识、绑定 workflow、candidate 读取都设计到位；
- **但实体之间缺乏「动作因果链」**：创建 Product 不主动询问「由哪些 Module 支撑」；创建 Module 不主动询问「属于哪个 Product」；创建 Decision 不联动触发「需要更新哪个实体状态」。

这种断层在 fixture 验收中不可见——因为 fixture 按已成立的关系预置；只有真实从零创建时，用户才感受到「我做了很多动作，但没有形成经营闭环」。

---

## 4. 三个阻断项：我从代码层面核实的根因

反馈中 2 轮阻断（f01/f04）、2 轮高摩擦（f02/f03）。这三个问题不需要新 phase，应走 `fix*` workflow。我逐一把它们落到代码：

### 4.1 feeback_01/02：onboarding「开始首轮录入」按钮首次点击无响应

**代码根因**（[onboarding-page.tsx](file:///home/dell/Projects/personal-software-company-os/frontend/src/features/onboarding/pages/onboarding-page.tsx#L82-L83)）：

```ts
const currentStep: OnboardingStep =
  focusedStep ?? serverStep ?? startStep ?? 'welcome'
```

「开始首轮录入」按钮只执行 `setStartStep('product')`（[L251-257](file:///home/dell/Projects/personal-software-company-os/frontend/src/features/onboarding/pages/onboarding-page.tsx#L251-L257)）。但冷启动时服务端 `first_run_state.current_step` 已被初始化为 `welcome`，于是 `serverStep` 优先级压过 `startStep`，`currentStep` 仍停留在 `welcome`，步骤不推进。

这解释了 feeback_02 的现象：用户绕到 Product Registry 创建产品后，服务端 `first_run_state` 被回写为 `product`，再回 onboarding 时 `serverStep` 变为 `product`，「流程才正常化」。

**建议**：`fix_001_onboarding_cold_start_state`——让本地 `startStep` 在服务端仍处于 `welcome` 时能覆盖推进（例如点击开始时先推进/初始化服务端状态，或调整优先级让「用户主动开始」优先于「服务端欢迎占位」）。

### 4.2 feeback_03：Current Focus 将已消费决策误标「待决策」

**代码根因**：`proposed` 状态同时承担「已留痕」与「待处理」双重语义。后端 current_focus_signals 将 `proposed` 直接派生为 pending 信号，前端 [current-focus-section.tsx](file:///home/dell/Projects/personal-software-company-os/frontend/src/features/dashboard/components/current-focus-section.tsx#L84) 仅透传渲染。用户已关联模块的决策仍被反复提示「待决策」。

**建议**：`fix_002_decision_status_semantics`——引入 `acknowledged` 中间态（`proposed → acknowledged → active → superseded → archived`），并让 Current Focus / Daily Review 只指向真正待处理的 `proposed`。

### 4.3 feeback_04：Daily Review → Decision Detail 无处理入口

**代码根因**：[decision-detail-page.tsx](file:///home/dell/Projects/personal-software-company-os/frontend/src/features/decision-center/pages/decision-detail-page.tsx#L124-L172) 仅展示概要、已关联目标、待关联目标与候选关联面板，**缺少 `proposed → acknowledged → active` 的状态推进 CTA**。review loop 在 Decision 处断裂——这正是 feeback_04「提示要处理，但没有处理入口」的直接来源。

**建议**：`fix_003_decision_detail_status_advance_cta`——在 Decision Detail 增加状态推进 CTA，并让 review 主动引导进入处理、处理完成后回流 review「已完成」状态。

---

## 5. 用户 8 点深度反思的分层处理

我把 8 点反思按处理层次分为三层。这是本轮评审最重要的分析框架。

### 5.1 第一层：阻断项与高摩擦点（fix 范畴，dry-run 收口前闭合）

即第 4 节三个阻断项。**这 4 项必须修复，dry-run 才能判「通过（含已知摩擦点）」并完成 mvp0.3 正式收口。**

### 5.2 第二层：实体关系与心智模型错位（phase 级，mvp0.4 主线）

- **反思 2 / 3**：onboarding 实体并列隔离、详情页频繁手动切换关联 → onboarding 逻辑化 + Cross-Entity CTA 矩阵
- **反思 4（部分）**：仅文本描述不触及源码 → Agent Consumption Layer 桥接「经营」与「工程」
- **feeback_03/04**：Decision 状态机过粗、review 业务闭环未完成 → Decision 状态机细化 + review 业务闭环

### 5.3 第三层：长期方向（mvp0.4+ 候选，受控探索）

- **反思 5**：用户初心回溯（探索-验证-迭代式经营）→ 锚定方向，不作为当轮主线
- **反思 6**：生产经营与工程化中间地带 → 通过「闭合动作链 + agent 消费」双层缓解，不扩对象
- **反思 7**：agent 为主要维护方、Web 展示状态 → Agent Consumption Layer（受控最小版）
- **反思 8**：跨项目规范/约束/技术栈统一维护 → Cross-Project Convention Asset（严格最小版或后移）

---

## 6. 对 mvp0.4 主线方向的判断

### 6.1 我不建议的主线结构

用户反思中最容易引导 mvp0.4 走偏的，是把「生产经营 + 工程化」「agent 协同」「跨项目规范」直接升格为三条并列主线。风险：

1. 三条主线都需要新建实体类型，把 mvp0.4 拉回「概念扩张先于使用闭环」的旧问题；
2. agent 协同一旦做重，容易演化为 AI 工作台，违背 `PSCO_3.md` 的 AI 增强层定位；
3. 跨项目规范一旦做重，容易演化为知识图谱，违背 mvp01 非目标。

### 6.2 我建议的主线结构

```
Asset-Action Closure（中心主线）
   +
Agent Consumption Layer（支撑能力，受控最小版）
   +
Cross-Project Convention Asset（候选探索，严格最小版或后移）
```

#### 6.2.1 中心主线：Asset-Action Closure

**目标**：把 PSCO 从「实体登记完备」推进到「动作链路完备」，让每一次创建、每一次 review、每一次决策都自然形成实体间因果链。

**必做范围**：

- onboarding 逻辑化（创建时主动建立实体关系，而非事后手动绑定）
- Decision 状态机细化（`draft / proposed / acknowledged / active / superseded / archived`）
- Cross-Entity CTA 矩阵（每个详情页都主动提示「下一步应该做什么」）
- review loop 业务闭环（review → Decision → 实体更新 → review 已完成）
- Current Focus 信号派生规则重做（基于新状态机）

**核心验收问题**：

- 用户从 onboarding 进入，能否一次性形成「产品 + 仓库 + 模块 + 决策」的完整关系网？
- Decision Detail 能否完成 `proposed → acknowledged → active` 状态推进并回流到 review？
- 每个实体详情页是否都主动提示「下一步」？
- Dashboard Current Focus 是否只指向真正待处理的 Decision？

#### 6.2.2 支撑能力：Agent Consumption Layer（受控最小版）

**目标**：让 PSCO 资产可被外部 agent（cursor / trae / codex）标准化消费，但不让 PSCO 自己变成 agent 工具。

**必做范围（最小版）**：

- 只读 agent 友好接口（基于既有 ConnectRPC，不新增第二套合同）
- AGENTS.md 风格的上下文导出（把 Product / Module / Decision / Repository 组合导出为 agent 可读上下文）
- 上下文快照机制（不演化为完整模板平台）

**明确不做**：内置 AI 工作台、自动代码生成/修改、agent 双向同步、agent 身份认证与权限治理（当前单用户）。

#### 6.2.3 候选探索：Cross-Project Convention Asset

**目标**：把跨项目可复用的规范、约束、技术栈作为新一类资产登记，解决用户新开项目从零开始的痛点。

**判断**：价值真实但风险最大。建议作为 mvp0.4 尾部候选或后移到 mvp0.5+。若进入，严格最小版——只做「资产登记 + 跨项目引用」，不做规范版本治理、自动同步、模板生成、依赖关系图谱。

### 6.3 与用户反思的对应关系

| 用户反思 | 我的处理建议 | 层次 |
| --- | --- | --- |
| 1. 精良但细节未闭环 | Asset-Action Closure 中心主线 | mvp0.4 主线 |
| 2. onboarding 实体并列隔离 | onboarding 逻辑化 | mvp0.4 主线 |
| 3. 详情页频繁切换 | Cross-Entity CTA 矩阵 | mvp0.4 主线 |
| 4. 仅文本描述未触及源码 | Agent Consumption Layer（外部工具消费 PSCO 上下文） | mvp0.4 支撑 |
| 5. 用户初心回溯 | 不扩张对象，闭合动作链路 | 方向锚定 |
| 6. 中间地带困境 | Asset-Action Closure 闭合「经营」侧 + Agent Consumption Layer 桥接「工程」侧 | mvp0.4 双层 |
| 7. agent 协同预想 | Agent Consumption Layer 受控最小版 | mvp0.4 支撑 |
| 8. 跨项目规范复用 | Cross-Project Convention Asset 候选探索 | mvp0.4 尾部或后移 |

---

## 7. 对 dry-run 收口与推进顺序的建议

### 7.1 不应把所有反馈塞进 dry-run 收口

如果试图把 8 点反思全部解决后再收口 dry-run，会导致 dry-run 验收闸无限延期、阻断 mvp0.4 正式入口建立。

### 7.2 推荐的推进路径

```
第一步：fix* 闭合 3 个阻断项
  ├─ fix_001_onboarding_cold_start_state
  ├─ fix_002_decision_status_semantics
  └─ fix_003_decision_detail_status_advance_cta
        ↓
第二步：重跑一轮 dry-run（只验证 fix 是否解决阻断项）
        ↓
第三步：dry-run 判「通过（含已知摩擦点）」，正式收口 mvp0.3
  └─ 已知摩擦点（onboarding 逻辑化、Cross-Entity CTA、review 业务闭环、agent 接口、跨项目规范）记入收口报告
        ↓
第四步：建立 mvp0.4 正式 /plan 入口（直接承接反馈、fix 结论、8 点反思）
```

**核心纪律**：fix 必须保持轻量（1~2 周内闭合），不能让阻断项演化为 phase 级工作；否则说明对根因判断有误。

---

## 8. 守界纪律提醒（对后续 `/spec` 与实现的硬约束）

1. `.proto` 仍是唯一长期合同源；
2. query 层继续保持纯只读；
3. 前端写路径继续收敛到切片固定承接位；
4. 不引入第二套路由、第二套状态管理、第二套 UI 事实源；
5. **onboarding 逻辑化不能演化为可配置工作流引擎**（仍是 6 步引导，关系建立是主动询问而非强制依赖）；
6. **Agent Consumption Layer 不能演化为 AI 工作台**（只读消费，不做内置对话/自动决策/代码生成）；
7. **Cross-Project Convention Asset 不能演化为知识图谱**（只登记工程可消费的规范，不做语义搜索/笔记/依赖图谱）；
8. **Decision 状态机细化不能演化为 Decision Intelligence**（只细化到区分已留痕/已确认/已生效，不做相似匹配/引用推荐/AI 增强）；
9. **review loop 不能演化为通用任务管理器**（review 动作必须锚定既有实体）。

---

## 9. 对后续正式 `/plan` 的输入

### 9.1 推荐的候选阶段结构（不冻结正式 phase 命名）

```
候选阶段一：dry-run 阻断项 fix 与收口
  └─ 闭合 3 个阻断项 → 重跑 dry-run → mvp0.3 正式收口
        ↓
候选阶段二：Asset-Action Closure 主线
  └─ onboarding 逻辑化 + Decision 状态机细化 + Cross-Entity CTA 矩阵
     + review 业务闭环 + Current Focus 信号派生重做
        ↓
候选阶段三：Agent Consumption Layer 支撑能力
  └─ agent 友好只读接口 + AGENTS.md 风格上下文导出 + 上下文快照机制
        ↓
候选阶段四（可选）：Cross-Project Convention Asset 候选探索
  └─ 严格最小版，视前置阶段完成情况决定是否进入
```

### 9.2 推荐的先后关系

1. 没有候选阶段一，dry-run 无法收口，mvp0.3 不能正式完成；
2. 没有候选阶段二，PSCO 仍停留在「实体登记完备但动作链路断裂」；
3. 没有候选阶段三，PSCO 资产无法被外部 agent 消费，「工程化配合」无法落地；
4. 候选阶段四是可选项，不阻断 mvp0.4 主线收口。

---

## 10. 对后续验收的硬要求

下一阶段正式验收至少必须回答：

1. 用户从 onboarding 进入，能否一次性形成完整实体关系网？
2. Decision Detail 能否完成状态推进并回流到 review？
3. 每个实体详情页是否都主动提示「下一步应该做什么」？
4. Dashboard Current Focus 是否只指向真正待处理的 Decision？
5. 外部 agent 能否通过标准化接口读取 Product 完整上下文？
6. 真实项目中，用户是否愿意围绕 PSCO 持续回来使用，且使用摩擦显著低于 mvp0.3 dry-run？

---

## 11. 最终结论

1. `phase01 ~ phase09` 已把 PSCO 从「理念」推进到「能登记、能复用、能看见复利」的可运行系统，且 dry-run 证明其结构克制、链路清晰、具有真实产品潜力。
2. **但 dry-run 首次真实暴露了「实体登记完备」与「经营动作闭环」之间的语义断层**——这正是 fixture 验收无法发现、只有真实使用才能暴露的问题。它证明 dry-run 的设计目的完全达成。
3. 因此 mvp0.4 的第一原则是**先闭合动作链**，方向收敛为 **Asset-Action Closure 中心主线 + Agent Consumption Layer 支撑能力 + Cross-Project Convention Asset 候选探索**，并先走 `fix*` 闭合 3 个阻断项完成 mvp0.3 收口。
4. 我坚持此前立场：**先闭合资产到动作的链路，agent 连接与工程化加速才有可消费的资产底座**；mvp0.4 是先补「动作链路」，再让「复利可被持续消费」。
5. 最大工程风险是「agent 协同」被偷渡成 AI 工作台、「onboarding 逻辑化」被偷渡成工作流引擎、「跨项目规范」被偷渡成知识图谱——必须守住只读消费、主动询问、工程可执行的底线。

用一句话收口 mvp0.4 的展望：

> **phase01~09 证明了 PSCO「能登记、能看见复利」；dry-run 证明了 PSCO「在真实使用中暴露了语义断层」；mvp0.4 要闭合这层断层，让 PSCO 从「资产登记系统」真正推进为「可被持续消费的个人软件公司 operating system」。**

---

*End of PSCO-mvp04-DPv4flash.md*
