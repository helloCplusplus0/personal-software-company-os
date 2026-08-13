# Personal Software Company OS

# MVP0.4 Direction Review — Based on Real-Project Dry-Run Feedback

**Author:** GLM-5.2
**Date:** 2026-08-13
**Role:** 作为 `PSCO-mvp03-GLM52.md` 「Operating Review Loop 为唯一主线」路线的提出者，以及 `phase01 ~ phase09` 实现链路的同行评审者，基于用户真实 `dry-run` 反馈文档对 PSCO 下一阶段推进方向给出第四轮评审性建议
**Document Type:** `review`
**Status:** 供后续正式 `/plan` 与交叉汇总参考，不直接构成正式 phase 命名、spec 路径或执行指令；不预设 `mvp0.4+` 既成事实

---

## 1. 目的说明

本文不重复书写 `phase01 ~ phase09` 的实现清单，也不替代下一阶段正式 `/plan`。

本文要回答的核心问题是：

> 在 `phase09` 已完成 `Template Reuse + Derived Intelligence Deepening` 正式收口、`Real-Project Dry-Run` 已产出 4 轮真实使用反馈之后，作为这套系统设计的同行评审者，我认为这些反馈真正揭示了什么、应该如何分层处理、以及 PSCO 下一阶段最符合长期复利价值的推进方向究竟是什么？

我的结论先写在前面：

> **`dry-run` 反馈不是失败信号，而是 `dry-run` 设计意图的真正达成 — 它揭示了一类此前评审无法发现的问题：PSCO 当前在「实体登记完备」与「经营动作闭环」之间存在一层未闭合的语义断层。下一阶段不应继续扩长期对象宽度，而应优先闭合这层断层，让 PSCO 从「可登记、可复用、可看见」推进到「可被真实经营动作持续消费」。**

更具体地说，我建议下一阶段的核心命题收敛为：

> **Asset-Action Closure（中心主线） + Agent Consumption Layer（支撑能力，受控版本） + Cross-Project Convention Asset（候选探索，严格最小版）**

与之前几轮评审不同，本文不主张把用户反思中的「生产经营 + 工程化」「agent 协同」「跨项目规范复用」直接升格为并列主线，理由会在第 6、7 节展开。

---

## 2. 评审依据

### 2.1 真实使用反馈（本文核心证据）

- `docs/review/PSCO-real-project-dry-run-user-manual-GPT54.md` — dry-run 标准使用手册
- `docs/review/PSCO-real-project-dry-run-user-manual-GPT54 feedback.md` — 4 轮真实使用反馈 + 8 点深度反思

### 2.2 当前正式推进与阶段现实

- `plan.md` — `phase01 ~ phase09` 全链路状态与后续 `dry-run` 进入条件
- `PSCO-mvp03-summarize-feedback.md` — MVP0.3 最终仲裁（Operating Review Loop + Template Reuse + Derived Intelligence + Dry-Run）
- `phase08-11` 与 `phase09-11` acceptance_report — 两个最近完成 phase 的正式验收结论
- `AGENTS.md` / `project_rules.md` — 根级真相源与协作门禁

### 2.3 原始长期方案

- `PSCO_0.md` — 战略愿景（Build / Accumulate / Compound）
- `PSCO_1.md` — 运行模型与领域实体
- `PSCO_2.md` — 模块系统与技术架构
- `PSCO_3.md` — 工作流引擎与 AI 策略
- `PSCO_4.md` — 产品 UX 规格

### 2.4 本文采取的判断原则

1. **真实使用证据优先于理论推演。** 4 轮反馈是 dry-run 设计意图的真正产出，其揭示的问题比任何专家评审都更接近系统真实状态。
2. **分层处理优先于整体扩张。** 反馈中的问题横跨 bug、设计缺陷、产品定位、长期方向四个层次，不能混为一谈。
3. **闭合优先于扩张。** 在当前已交付能力尚未形成真正闭环前，继续扩对象宽度会加重「中间地带」困境。
4. **长期方向优先于短期补丁。** 用户反思中的「agent 协同」「跨项目规范」即使方向正确，也必须受控进入，不能因诉求强烈就提前升格。
5. **不预设 phase 命名。** 遵守 `plan.md §4` 与 `project_rules.md` 的切换条件，本文只给方向，不冻结 `phase10` 名称或 `mvp0.4+` 范围。

---

## 3. 对 dry-run 反馈的整体判断

### 3.1 dry-run 达成了它的设计目的

`PSCO-mvp03-summarize-feedback.md §4.8` 明确写道：

> 真实项目 `dry-run` 不是可有可无的附录，而应作为 `mvp0.3` 后段验收中的独立交付要求。它不能替代 fixture 验收，但必须补上三件事：真实使用摩擦是否可接受；review loop 是否真的改变下一步动作；模板复用是否真的缩短了下一次创造路径。

按这个标尺回看 4 轮反馈，结论非常清楚：

- **真实使用摩擦**：发现 onboarding 冷启动断裂、Decision Detail 缺处理入口、Current Focus 状态语义错位 — 摩擦点被真实暴露
- **review loop 是否真的改变下一步动作**：用户明确反馈「Daily Review 引导进入 Decision Detail 后没有处理入口，未形成闭环」— review loop 的断裂点被精确定位
- **模板复用是否真的缩短创造路径**：用户未走到 B 线深度验证，但已确认 onboarding 阶段的实体并列隔离问题影响了对「复用」的心智建立

换句话说：

> **dry-run 不是证明 PSCO 完美，而是证明 PSCO 在真实使用中暴露了哪些 fixture 验收无法发现的问题。从这个角度看，本轮 dry-run 完全达成了设计目的。**

### 3.2 反馈揭示的不是「功能缺失」，而是「语义断层」

最关键的洞察在用户反思第二点：

> 「onboarding 阶段看似六个流程，其实就是分别新增实体，不存在任何逻辑关系。」

这一句话精准命中了 PSCO 当前设计的核心问题：

- **实体登记能力完备**：Product / Module / Decision / Repository 都能创建、编辑、关联
- **双向关联链路精良**：来源标识、绑定 workflow、candidate 读取都设计到位
- **但实体之间缺乏「动作因果链」**：创建 Product 不强制询问「由哪些 Module 支撑」；创建 Module 不主动询问「属于哪个 Product」；创建 Decision 不联动触发「需要更新哪个实体状态」

这种断层在 fixture 验收中是不可见的，因为 fixture 是按已成立的关系预置的；只有在真实使用中，用户从零开始按 onboarding 顺序创建时，才会感受到「我做了很多动作，但没有形成经营闭环」。

### 3.3 用户反思的真正价值

用户 8 点反思中，最有价值的是第 5、6、7 点：

- **第 5 点**：用户回溯初心 — PSCO 应该是「Personal Software Company」，覆盖生产经营+工程化
- **第 6 点**：PSCO 当前处于「生产经营」与「工程化」中间地带，两边均有功能但不深入
- **第 7 点**：期待 agent 可访问 PSCO、PSCO 前端展示状态、用户做必要操作

这三点共同指向一个判断：

> **PSCO 当前定位的「经营与资产系统」是正确的，但「经营」与「资产」之间的「动作链路」没有真正打通。下一阶段的核心任务不是扩对象，而是把动作链路闭合到可被持续消费的程度。**

---

## 4. 反馈问题的分层分析

我把 4 轮反馈与 8 点反思中的所有问题，按处理层次分为三层。这是本文最重要的分析框架。

### 4.1 第一层：阻断项与高摩擦点（fix 范畴，应在 dry-run 收口前闭合）

这些问题是具体的 bug 或局部设计缺陷，不需要新 phase，应该走 `fix*` workflow。

| 反馈点 | 根因（基于代码定位） | 建议处理 |
|---|---|---|
| feeback_01：onboarding "开始首轮录入"按钮首次点击无响应 | `onboarding-page.tsx` WelcomeStep 的 onStart 触发后，依赖 `first_run_state.current_step` 推进；但首次访问时服务端 `first_run_state` 未初始化，导致状态机断裂 | `fix_001_onboarding_cold_start_state` |
| feeback_02：绕到 products/new 创建产品后再回 onboarding 才正常 | 与 feeback_01 同根因 — 创建产品副作用写入了 `first_run_state`，绕开了 WelcomeStep 的状态推进 | 同 fix_001 |
| feeback_03：Current Focus "待决策"标识与实际状态不符 | `decisioncenter/types.go` 中 `proposed` 同时承担「已留痕」与「待处理」双重语义；`use-feedback-signals-read.ts` 把 `proposed` 直接映射为 `pending_decision` 信号 | `fix_002_decision_status_semantics` — 引入 `acknowledged` 中间态或细化 signal 派生规则 |
| feeback_04：Daily Review → Decision Detail 无处理入口 | `decision-detail-page.tsx` 只展示与关联模块，缺少 `proposed → active` 的状态推进 CTA | `fix_003_decision_detail_status_advance_cta` |

**判断**：这 4 项必须修复，dry-run 才能判为「通过（含已知摩擦点）」并完成 mvp0.3 正式收口。不修复就强行收口，会让 mvp0.4 建立在未闭合的底座上。

### 4.2 第二层：实体关系与心智模型错位（phase 级，候选 mvp0.4 主线）

这些问题不能通过 fix 解决，因为它们涉及设计语义调整，但也不应该升格为 mvp0.4+ 长期方向。它们应该成为 mvp0.4 的中心主线。

#### 4.2.1 onboarding 逻辑化

用户反馈：

> 「onboarding 阶段看似六个流程，其实就是分别新增实体，不存在任何逻辑关系。」

当前 onboarding 6 步是「分别创建 Product / Repository / Module / Decision」的并行引导，而不是「先有产品 → 产品需要仓库 → 产品由模块支撑 → 模块需要决策」的因果引导。

但解法**不是**让 onboarding 变重，而是要让实体之间的逻辑关系在创建时就被主动建立：

- 创建 Product 时主动询问「这个产品由哪些已有 Module 支撑？或需要新建哪些 Module？」
- 创建 Module 时主动询问「这个 Module 属于哪个 Product？或作为独立资产登记？」
- 创建 Decision 时主动询问「这个 Decision 影响哪个 Product / Module？」

这是「Asset-Action Closure」的核心 — 让每一次创建动作都自然形成实体间关系，而不是事后到详情页手动绑定。

#### 4.2.2 Decision 状态机细化

当前 Decision 状态只有 `proposed / active / superseded / archived`，且 `proposed` 承担双重语义。

建议细化为：

```
draft → proposed → acknowledged → active → superseded → archived
```

其中：

- `draft`：草稿，未留痕
- `proposed`：已记录但未审阅
- `acknowledged`：已确认但未执行（用户反馈 feeback_03 的真实诉求）
- `active`：已生效
- `superseded`：被新决策取代
- `archived`：归档

这能让 Current Focus 的「待决策」信号只指向真正待处理的 `proposed` 状态，而不是所有未到 `active` 的决策。

#### 4.2.3 Cross-Entity CTA 矩阵

用户反馈：

> 「不停的在各个实体详情页中切换，点击执行关联，我不能说这样的设计错误，但是我感觉使用摩擦有点大。」

当前每个实体详情页只关心自己的展示与编辑，缺少「下一步应该做什么」的主动 CTA。例如：

- Product Detail（缺仓库时）：主动 CTA「绑定仓库」
- Product Detail（缺模块时）：主动 CTA「添加支撑模块」
- Module Detail（未被任何 Product 使用时）：主动 CTA「加入某个 Product」
- Decision Detail（proposed 状态）：主动 CTA「确认并推进到 acknowledged」
- Decision Detail（影响 Product 但未关联）：主动 CTA「关联到目标 Product」

这本质上是把 Dashboard 的 CTA 矩阵下沉到每个实体详情页，让每页都成为「经营动作的局部起点」。

#### 4.2.4 Dashboard 与 review loop 的回流闭合

用户反馈 feeback_04 最关键的判断是：

> 「Daily Review->Current Focus 与待处理决策均提示要处理，但是又没有出现处理入口，我感觉并没有形成闭环。」

这指向 review loop 的回流链路在 Decision 处断裂。修复方式不只是 fix_003（Decision Detail 加状态推进 CTA），还要让 review 主动引导用户进入 Decision 的处理动作，并在 Decision 处理完成后回流到 review 的「已完成」状态。

这是 `phase08` Operating Review Loop 的真正完成 — 之前 phase08 完成的是 review 会话闭环，不是业务动作闭环。

### 4.3 第三层：长期方向（mvp0.4+ 候选，受控探索）

用户反思中的第 6、7、8 点属于长期方向，应该作为 mvp0.4+ 的候选输入，但**不能直接升格为 mvp0.4 主线**。理由见第 6 节。

#### 4.3.1 Agent Consumption Layer（agent 可消费的资产接口）

用户期待：

> 「agent 应该可以知道并使用 PSCO，绝大多数的 PSCO 应该以 agent 维护为主，PSCO web 前端主要展示这个 Personal Software Company 实际状态。」

这个方向是正确的，但当前 PSCO 的 `.proto + ConnectRPC` 主线已经具备机器可消费能力。真正缺的不是「agent 可访问」，而是「agent 友好的上下文输出层」— 类似 AGENTS.md 风格的标准化上下文导出。

这可以作为 mvp0.4 的支撑能力（受控最小版），但不应该作为主线。理由：

- 当前 PSCO 自身的资产-动作链路未闭合，先让 agent 消费未闭合的资产会放大混乱
- agent 接口层一旦做重，容易演化为 AI 工作台，违背 `PSCO_3.md` 的 AI 增强层定位

#### 4.3.2 Cross-Project Convention Asset（跨项目规范资产）

用户期待：

> 「在每一个长期推进的项目中探索得出了一些有利于 agent 执行保持一致性的：规范，约束，机制，全局技术栈等非常有价值且可以全局使用的资料，但是现在面临着新开新项目时目录结构，规范，约束都需要从零开始重新建。」

这个诉求非常真实，且与 PSCO 的「资产积累」哲学一致。但它的解法需要慎重：

- **正确解法**：把 `Convention / Standard / TechStack` 作为新一类资产对象，可被跨项目引用
- **错误解法**：把它做成通用知识图谱或文档管理系统

这可以作为 mvp0.4+ 的候选探索方向，但严格最小版 — 只做「资产登记 + 跨项目引用」，不做「自动同步 / 版本治理 / 模板生成」。

#### 4.3.3 PSCO 与真实工程化开发的配合

用户期待：

> 「应该考虑 PSCO 如何才能真正管理并加速实际源代码开发工作，而不能仅仅孤立的做一些文本记录。」

这是最容易被误解的方向。我的判断是：

- PSCO **不应该**自己执行代码、不应该是 IDE、不应该做代码托管
- PSCO **应该**通过 agent consumption layer 让真实开发工具（cursor / trae / codex）可消费 PSCO 的资产上下文
- 这样「工程化加速」是通过「外部 agent 拿到 PSCO 上下文后加速」实现，而不是 PSCO 自己变成工程工具

这个区分非常关键，否则 PSCO 会失去「经营与资产系统」的核心定位。

---

## 5. 对 dry-run 收口策略的建议

### 5.1 不应把所有反馈塞进 dry-run 收口

如果试图把 8 点反思全部解决后再收口 dry-run，会导致：

- dry-run 验收闸无限延期
- mvp0.3 长期处于「进行中」状态
- 阻断 mvp0.4 正式入口建立

### 5.2 推荐的收口路径

**第一步**：走 `fix*` workflow 闭合第一层 4 个阻断项

- `fix_001_onboarding_cold_start_state`
- `fix_002_decision_status_semantics`
- `fix_003_decision_detail_status_advance_cta`
- （feeback_03 的 Current Focus 标识问题在 fix_002 中一并解决）

**第二步**：重新走一轮 dry-run 验证 fix 效果

- 只验证 fix 是否真的解决了阻断项
- 不要求验证第二层、第三层问题

**第三步**：dry-run 判为「通过（含已知摩擦点）」，正式收口 mvp0.3

- 已知摩擦点（onboarding 逻辑化、Cross-Entity CTA、review loop 业务闭环）记入收口报告
- 8 点反思作为 mvp0.4 仲裁上游输入正式留档

**第四步**：建立 mvp0.4 正式 `/plan` 入口

- 直接承接 dry-run 反馈、fix 结论、8 点反思
- 不预设 phase 名称，先做方向仲裁

### 5.3 dry-run 验收报告应明确记录的「已知摩擦点」

- onboarding 实体并列隔离（待 mvp0.4 Asset-Action Closure 解决）
- Decision 状态机过粗（待 mvp0.4 细化）
- 实体详情页缺 Cross-Entity CTA（待 mvp0.4 下沉）
- review loop 业务闭环未完成（待 mvp0.4 闭合）
- agent 协同接口未建立（待 mvp0.4+ 候选）
- 跨项目规范资产未支持（待 mvp0.4+ 候选）

这些已知摩擦点不应该阻断 dry-run 收口，但必须在收口报告中明确留档，作为 mvp0.4 的正式输入。

---

## 6. 对 mvp0.4 主线方向的判断

### 6.1 我不建议的主线结构

用户反思中最容易引导 mvp0.4 走偏的方向，是把「生产经营 + 工程化」「agent 协同」「跨项目规范」直接升格为三条并列主线。这种结构的风险：

1. **三条主线都需要新建实体类型**，会把 mvp0.4 拉回「概念扩张先于使用闭环」的旧问题
2. **agent 协同一旦做重，容易演化为 AI 工作台**，违背 `PSCO_3.md` 的 AI 增强层定位
3. **跨项目规范一旦做重，容易演化为知识图谱**，违背 `PSCO-mvp01-summarize-feedback.md` 的非目标

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

- onboarding 逻辑化（创建时主动建立实体关系）
- Decision 状态机细化（draft / proposed / acknowledged / active / superseded / archived）
- Cross-Entity CTA 矩阵（每个详情页都成为局部经营起点）
- review loop 业务闭环（review → Decision → 实体更新 → review 已完成状态）
- Current Focus 信号派生规则重做（基于新状态机）

**核心验收问题**：

- 用户从 onboarding 进入，能否一次性形成「产品 + 仓库 + 模块 + 决策」的完整关系网，而不是分别创建后手动绑定？
- Decision Detail 是否能完成 `proposed → acknowledged → active` 的状态推进，并回流到 review？
- 每个实体详情页是否都主动提示「下一步应该做什么」？
- Dashboard Current Focus 是否只指向真正待处理的 Decision？

#### 6.2.2 支撑能力：Agent Consumption Layer（受控最小版）

**目标**：让 PSCO 资产可被外部 agent（cursor / trae / codex）标准化消费，但不让 PSCO 自己变成 agent 工具。

**必做范围（最小版）**：

- 一套只读的 agent 友好接口（基于既有 ConnectRPC，不新增第二套合同）
- AGENTS.md 风格的上下文导出能力（把 Product / Module / Decision / Repository 组合导出为 agent 可读上下文）
- 一个简单的「上下文快照」机制（不演化为完整模板平台）

**明确不做**：

- 不做 PSCO 内置 AI 工作台
- 不做自动代码生成或修改
- 不做 agent 身份认证或权限治理（当前阶段单用户）
- 不做 agent 双向同步（只读消费即可）

**核心验收问题**：

- 外部 agent 能否通过标准化接口读取某个 Product 的完整上下文（含关联 Module / Decision / Repository）？
- 这个上下文是否足够支撑 agent 在新项目中复用既有规范？

#### 6.2.3 候选探索：Cross-Project Convention Asset

**目标**：把跨项目可复用的规范、约束、技术栈作为新一类资产登记，解决用户新开项目时从零开始的痛点。

**判断**：这个方向价值真实，但风险也最大。建议两种处理方式之一：

- **方式 A（推荐）**：作为 mvp0.4 的尾部候选，先完成 Asset-Action Closure 与 Agent Consumption Layer，再视精力决定是否进入
- **方式 B**：后移到 mvp0.5+，让 mvp0.4 专注闭合

**如果进入，必做范围（严格最小版）**：

- 新增 `Convention` 实体（或复用 Module 的扩展字段，不新增核心表）
- 支持跨项目引用（Product → Convention 的轻量绑定）
- 支持从既有项目导出规范为 Convention 资产

**明确不做**：

- 不做规范版本治理
- 不做自动同步或模板生成
- 不做规范间依赖关系图谱

### 6.3 与用户反思的对应关系

| 用户反思 | 我的处理建议 | 层次 |
|---|---|---|
| 1. PSCO 精良但细节未闭环 | Asset-Action Closure 中心主线 | mvp0.4 主线 |
| 2. onboarding 实体并列隔离 | onboarding 逻辑化 | mvp0.4 主线 |
| 3. 详情页频繁切换 | Cross-Entity CTA 矩阵 | mvp0.4 主线 |
| 4. 仅文本描述未触及源码 | Agent Consumption Layer（让外部工具消费 PSCO 上下文） | mvp0.4 支撑 |
| 5. 用户初心回溯 | 不扩张对象，闭合动作链路 | 方向锚定 |
| 6. 中间地带困境 | Asset-Action Closure 闭合「经营」侧；Agent Consumption Layer 桥接「工程」侧 | mvp0.4 双层 |
| 7. agent 协同预想 | Agent Consumption Layer 受控最小版 | mvp0.4 支撑 |
| 8. 跨项目规范复用 | Cross-Project Convention Asset 候选探索 | mvp0.4 尾部或后移 |

---

## 7. 对几个关键风险的提醒

### 7.1 警惕「agent 协同」被偷渡成 AI 主线

用户反思第 7 点的诉求很强，容易引导 mvp0.4 把 agent consumption layer 做成 AI 工作台。必须守住底线：

- PSCO 只提供**只读**上下文消费接口
- 不做内置 AI 对话
- 不做自动决策或自动代码生成
- agent 是外部消费者，不是 PSCO 的内置模块

### 7.2 警惕「跨项目规范」被偷渡成知识图谱

用户反思第 8 点提到「曾收集过精益创业方法论书籍」，这种背景容易引导 PSCO 做成「个人知识管理系统」。必须守住底线：

- Convention 只登记**可被工程消费的规范资产**（技术栈、目录结构、agent 规则、协作约束）
- 不做书籍笔记、不做知识图谱、不做语义搜索
- Convention 必须是「工程可执行」的，不是「理论可参考」的

### 7.3 警惕「onboarding 逻辑化」被偷渡成完整工作流引擎

onboarding 逻辑化的正确解法是「创建时主动建立关系」，不是「把 onboarding 做成完整的工作流引擎」。必须守住底线：

- onboarding 仍然是 6 步引导，不演化为可配置工作流
- 实体间关系建立是「主动询问」而非「强制依赖」
- 不做 onboarding 模板或 onboarding 版本管理

### 7.4 警惕「Decision 状态机细化」被偷渡成 Decision Intelligence

Decision 状态机细化是为了解决 feeback_03 的语义错位，不是为了引入决策智能。必须守住底线：

- 状态机只细化到能区分「已留痕」「已确认」「已生效」
- 不做相似决策匹配、不做历史决策引用推荐、不做 AI 决策上下文增强
- 这些仍然后移到 mvp0.5+ 或更后

### 7.5 警惕 dry-run 收口被无限延期

第一层 4 个阻断项必须用 fix 快速闭合，不能让它们演化为 phase 级工作。如果 fix 工作量超过 1-2 周，说明对问题根因的判断有误，应该重新审视。

---

## 8. 对 mvp0.4 范围基线的建议

### 8.1 必做范围

1. **Asset-Action Closure**
   - onboarding 逻辑化（创建时主动建立实体关系）
   - Decision 状态机细化（draft / proposed / acknowledged / active / superseded / archived）
   - Cross-Entity CTA 矩阵（每个详情页主动提示下一步）
   - review loop 业务闭环（review → Decision → 实体更新 → review 已完成）
   - Current Focus 信号派生规则重做

2. **Agent Consumption Layer（受控最小版）**
   - 只读 agent 友好接口（基于既有 ConnectRPC）
   - AGENTS.md 风格的上下文导出
   - 上下文快照机制（不演化为完整模板平台）

3. **dry-run 收口前置 fix**
   - fix_001_onboarding_cold_start_state
   - fix_002_decision_status_semantics
   - fix_003_decision_detail_status_advance_cta

### 8.2 候选范围（视精力决定）

- **Cross-Project Convention Asset（严格最小版）**
  - 新增 Convention 实体或扩展字段
  - 跨项目引用
  - 从既有项目导出规范

### 8.3 明确不做范围

- `Venture` 正式主线化
- `Opportunity / Feature / Experiment` 流程化
- `Capability` 重实体 CRUD
- GitHub OAuth / 自动导入
- AI 一级工作台或内置 AI 对话
- 自动扫描 / 知识图谱 / 语义搜索
- Rust Intelligence Layer
- 完整模板平台 / 模板版本管理
- `Decision Intelligence`（相似匹配、引用推荐、AI 上下文增强）
- PSCO 自身的 IDE 集成或代码执行
- onboarding 可配置工作流引擎

### 8.4 与既有约束的一致性

本文建议的 mvp0.4 范围完全遵守 `PSCO-mvp03-summarize-feedback.md §4.9` 的后移项清单，不引入任何已明确后移的方向作为主线。`Venture / Decision Intelligence / AI Context Enhancement / Opportunity / Feature / Experiment / 完整模板平台 / GitHub OAuth / 自动扫描` 全部继续后移。

唯一新增的候选方向是 `Cross-Project Convention Asset`，且严格受控为最小版或后移。

---

## 9. 对后续正式 `/plan` 的输入

### 9.1 推荐的候选阶段结构

本文不冻结正式 `phase` 命名，但冻结如下候选推进顺序：

#### 候选阶段一：dry-run 阻断项 fix 与收口

- 走 `fix*` workflow 闭合 4 个阻断项
- 重新走一轮 dry-run 验证
- 完成 mvp0.3 正式收口

#### 候选阶段二：Asset-Action Closure 主线

- onboarding 逻辑化
- Decision 状态机细化
- Cross-Entity CTA 矩阵
- review loop 业务闭环
- Current Focus 信号派生规则重做

#### 候选阶段三：Agent Consumption Layer 支撑能力

- agent 友好只读接口
- AGENTS.md 风格上下文导出
- 上下文快照机制

#### 候选阶段四（可选）：Cross-Project Convention Asset 候选探索

- 严格最小版
- 视前置阶段完成情况决定是否进入

### 9.2 推荐的先后关系

1. 没有候选阶段一，dry-run 无法收口，mvp0.3 不能正式完成
2. 没有候选阶段二，PSCO 仍停留在「实体登记完备但动作链路断裂」状态
3. 没有候选阶段三，PSCO 资产无法被外部 agent 消费，用户反思中的「工程化配合」无法落地
4. 候选阶段四是可选项，不阻断 mvp0.4 主线收口

### 9.3 对正式 `/spec` 的硬约束

后续 `/spec` 与实现必须继续遵守：

1. `.proto` 仍是唯一长期合同源
2. query 层继续保持纯只读
3. 前端写路径继续收敛到切片固定承接位
4. 不引入第二套路由、第二套状态管理、第二套 UI 事实源
5. onboarding 逻辑化不能演化为可配置工作流引擎
6. Agent Consumption Layer 不能演化为 AI 工作台
7. Cross-Project Convention Asset 不能演化为知识图谱
8. Decision 状态机细化不能演化为 Decision Intelligence

### 9.4 对后续验收的硬要求

下一阶段正式验收至少必须回答以下问题：

1. 用户从 onboarding 进入，能否一次性形成完整实体关系网？
2. Decision Detail 能否完成状态推进并回流到 review？
3. 每个实体详情页是否都主动提示「下一步应该做什么」？
4. Dashboard Current Focus 是否只指向真正待处理的 Decision？
5. 外部 agent 能否通过标准化接口读取 Product 完整上下文？
6. 真实项目中，用户是否愿意围绕 PSCO 持续回来使用，且使用摩擦显著低于 mvp0.3 dry-run？

---

## 10. 最终结论

`phase01 ~ phase09` 已经证明 PSCO 的最小资产主线、经营回路与支撑能力都成立。`Real-Project Dry-Run` 也完成了它的设计目的 — 真实暴露了 fixture 验收无法发现的语义断层。

因此，`mvp0.4` 的最终仲裁结论是：

> **下一步不应继续扩长期对象宽度，也不应直接跳到 Venture / Decision Intelligence / AI Context Enhancement；而应优先闭合「实体登记」与「经营动作」之间的语义断层，让 PSCO 从「可登记、可复用、可看见」推进到「可被真实经营动作持续消费」的 Operating System。同时以受控最小版建立 Agent Consumption Layer，让外部开发工具可消费 PSCO 资产上下文，桥接「经营层」与「工程层」。**

正式收敛方向如下：

1. **Asset-Action Closure** 为中心主线；
2. **Agent Consumption Layer**（受控最小版）为支撑能力；
3. **Cross-Project Convention Asset** 为候选探索（严格最小版或后移）；
4. `Venture / Decision Intelligence / AI Context Enhancement / Opportunity / Feature / Experiment / 完整模板平台 / GitHub OAuth / 自动扫描` 继续后移到 `mvp0.5+` 候选范围。

如果用一句话收口：

> **`phase01 ~ phase09` 证明了 PSCO「能登记、能复用、能看见复利」；`dry-run` 证明了 PSCO「在真实使用中暴露了语义断层」；`mvp0.4` 必须闭合这层断层，让 PSCO 从「资产登记系统」真正推进为「可被持续消费的个人软件公司 operating system」。**

---

## 附录：与既有评审的关系

### 与 `PSCO-mvp03-GLM52.md` 的关系

本文延续了 `PSCO-mvp03-GLM52.md` 「Operating Review Loop 为唯一主线」的判断思路 — `mvp0.4` 的 Asset-Action Closure 本质上是 `phase08` Operating Review Loop 的真正完成。`phase08` 完成的是 review 会话闭环，`mvp0.4` 要完成的是业务动作闭环。

### 与 `PSCO-mvp03-summarize-feedback.md` 的关系

本文完全遵守 `PSCO-mvp03-summarize-feedback.md §4.9` 的后移项清单，不引入任何已明确后移的方向作为主线。本文新增的 `Cross-Project Convention Asset` 候选方向，是基于 dry-run 真实反馈新增的仲裁输入，严格受控为最小版或后移。

### 与 `PSCO-mvp03-GPT54.md` 的关系

`PSCO-mvp03-GPT54.md` 的「一主两翼一闸」结构在 mvp0.3 阶段已经完成验证。本文建议的 mvp0.4 结构延续了这个思路 — Asset-Action Closure 为一主，Agent Consumption Layer 与 Cross-Project Convention Asset 为两翼，dry-run 收口为进入闸。

### 与 dry-run 反馈的关系

本文所有判断都基于 `docs/review/PSCO-real-project-dry-run-user-manual-GPT54 feedback.md` 的真实使用证据。反馈中的 4 轮 feeback 与 8 点反思被分层处理：

- 阻断项 → fix 范畴
- 设计语义错位 → mvp0.4 主线
- 长期方向 → mvp0.4+ 候选

这种分层处理是为了避免「把所有问题塞进一个 phase」的常见陷阱。

---

**Final Record**
**Signed by: GLM-5.2**
**Date: 2026-08-13**
