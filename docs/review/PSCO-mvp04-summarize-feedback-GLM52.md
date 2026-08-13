# Personal Software Company OS

# MVP0.4 第一轮专家讨论标准化汇总反馈

**Author:** GLM-5.2
**Date:** 2026-08-13
**Role:** 基于 `PSCO-mvp04-DPv4flash.md` / `PSCO-mvp04-DPv4pro.md` / `PSCO-mvp04-gemini31pro.md` / `PSCO-mvp04-GLM52.md` / `PSCO-mvp04-GPT54.md` / `PSCO-mvp04-qwen37pro.md` 六份专家文档,以独立评审者视角进行交叉分析,产出供后续正式 `/plan` 与仲裁参考的标准化汇总
**Document Type:** `review`
**Status:** 供后续正式 `/plan` 与最终共识仲裁参考,不直接冻结正式 phase 命名、spec 路径或执行指令

---

## 0. 文档定位与前置说明

### 0.1 本文的特殊性

本文与六份专家文档的关系:

- 六份专家文档由六位不同模型作者基于同一份真实使用反馈(`PSCO-real-project-dry-run-user-manual-GPT54 feedback.md`)独立产出
- 本文不替代任何一份专家文档,而是对六份文档进行交叉分析,识别共识与分歧,产出综合判断
- 本文不替代后续正式 `/plan`,只给方向性输入

本文要回答的核心问题是:

> 在六位专家基于真实 dry-run 反馈独立给出 mvp0.4 方向建议后,这些观点在哪些维度上形成共识、在哪些维度上存在分歧、以及基于这些共识与分歧,PSCO 下一阶段最符合长期复利价值的推进方向究竟是什么?

### 0.2 本文不替代正式 `/plan`

本文归入 `docs/review/`,作为第一轮专家讨论的标准化汇总。正式推进仍需按 `project_rules.md §4.1` 走 `/plan -> /spec -> 实现 -> 验收 -> 收口`。

### 0.3 本文与既有共识基线的关系

本文完全遵守:

- `PSCO-summarize-feedback.md` — 最终共识仲裁
- `PSCO-mvp03-summarize-feedback.md` — MVP0.3 最终仲裁
- `project_rules.md` — 协作门禁与 workflow 规则
- `TECH_STACK_BASELINE.md` — 技术栈基线
- `AGENTS.md` — 根级真相源入口

本文不引入任何已明确后移的方向作为主线。

---

## 1. 六份专家文档核心观点概览

### 1.1 DeepSeek-V4-Flash(`PSCO-mvp04-DPv4flash.md`)

**核心立场**:**先闭合资产到动作的链路,agent 连接与工程化加速才有可消费的资产底座**。

**主线结构**:
- 中心主线:Asset-Action Closure(资产-动作闭环)
- 支撑能力:Agent Consumption Layer(受控最小版)
- 候选探索:Cross-Project Convention Asset(严格最小版或后移)

**关键贡献**:
- 提出分层处理框架:第一层阻断项(fix 范畴)、第二层实体关系与心智模型错位(phase 级,mvp0.4 主线)、第三层长期方向(mvp0.4+ 候选)
- 把反馈中的三个阻断项落到代码根因,给出具体 fix 编号:`fix_001_onboarding_cold_start_state` / `fix_002_decision_status_semantics` / `fix_003_decision_detail_status_advance_cta`
- 明确反对把三条主线升格为并列,理由是会把 mvp0.4 拉回"概念扩张先于使用闭环"的旧问题

**明确不做**:内置 AI 工作台、自动代码生成、agent 双向同步、规范版本治理、自动同步、模板生成、依赖关系图谱。

### 1.2 DeepSeek-V4-Pro(`PSCO-mvp04-DPv4pro.md`)

**核心立场**:**MVP0.4 的方向从"向上叠加智能"调整为"向下扎根连接"**。

**主线结构**:
- 主线 A:闭环修复(阻断项修复)
- 主线 B:真实连接(从"文本描述"到"代码触及")
- 主线 C:Agent-Native 基础(为未来奠基)

**关键贡献**:
- 对 `mvp03-DPv4pro` 自身原建议路线(Venture → Decision Intelligence → AI Context Enhancement)进行公开重新校准,把 Venture 后移到 mvp0.5+
- 提出具体的真实连接能力集:仓库元数据回显(GitHub API 读时获取)、模块源码结构预览、版本与 Git Tag 关联、仓库活动时间线
- 给出具体的 phase 推进建议:`phase10`(闭环修复) → `phase11`(真实连接) → `phase12`(Agent-Native 基础)
- 对用户八点反思逐条回应,给出明确的诊断与处理建议

**明确不做**:完整的 Decision 生命周期工作流引擎、代码质量分析或静态分析、本地文件系统监控、多平台仓库适配、Agent 运行时或编排引擎、"探索型"场景的完整实现。

### 1.3 Gemini-3.1-Pro(`PSCO-mvp04-gemini31pro.md`)

**核心立场**:**纯人工维护的资产登记系统,在单人+AI 的研发模式下是没有未来的**;MVP0.4 应是 PSCO 真正蜕变为"OS"的起点。

**主线结构**:
- 主线一:Agent-Driven Workflow(读写倒置,Web 侧重读与宏观决策,Agent 侧重写与微观执行)
- 主线二:跨项目全局上下文资产化(Context as an Asset,自动向目标仓库注入 `.trae/rules` 或 `.cursorrules`)
- 主线三:融入精益创业模型(Lean Business Loop,解锁 Venture/Experiment/Opportunity 实体)

**关键贡献**:
- 提出"读写倒置"范式:Web 退后一步成为驾驶舱,Agent 通过 MCP Server 自动完成 Module 登记与绑定
- 强调"下发机制":PSCO 不仅记录资产,更向目标仓库下发规范资产,成为多项目 Agent 开发的"母舰"
- 提出精益创业模型的最小数据结构:假设 → 实验设计 → 数据指标 → 验证结果

**明确约束**:Local-First 延续、渐进式精益、隔离修复与重构。

### 1.4 GLM-5.2(`PSCO-mvp04-GLM52.md`)

**核心立场**:**先闭合"实体登记"与"经营动作"之间的语义断层,让 PSCO 从"可登记、可复用、可看见"推进到"可被真实经营动作持续消费"**。

**主线结构**:
- 中心主线:Asset-Action Closure(资产-动作闭环)
- 支撑能力:Agent Consumption Layer(受控最小版)
- 候选探索:Cross-Project Convention Asset(严格最小版或后移)

**关键贡献**:
- 延续 `PSCO-mvp03-GLM52.md` 的"Operating Review Loop 为唯一主线"判断思路,把 Asset-Action Closure 定位为 phase08 的真正完成
- 提出明确的 Decision 状态机细化方案:`draft → proposed → acknowledged → active → superseded → archived`
- 提出 Cross-Entity CTA 矩阵:把 Dashboard 的 CTA 矩阵下沉到每个实体详情页,让每页都成为"经营动作的局部起点"
- 提出 5 条守界纪律:防止 agent 协同被偷渡成 AI 主线、跨项目规范被偷渡成知识图谱、onboarding 逻辑化被偷渡成工作流引擎、Decision 状态机被偷渡成 Decision Intelligence、review loop 被偷渡成通用任务管理器

**明确不做**:Venture 主线化、Opportunity/Feature/Experiment 流程化、Capability 重实体 CRUD、GitHub OAuth、AI 一级工作台、自动扫描、Rust Intelligence Layer、完整模板平台、Decision Intelligence、PSCO 自身 IDE 集成、onboarding 可配置工作流引擎。

### 1.5 GPT-5.4(`PSCO-mvp04-GPT54.md`)

**核心立场**:**PSCO 当前已经证明"骨架正确";mvp0.4 之后更重要的任务,是证明它不仅能被登记,而且能真正驱动一个个人软件公司的持续经营与持续生产**。

**主线结构**(三层而非三主线):
- 第一层:必须立即闭合的现实阻断项(fix 范畴)
- 第二层:下一阶段最值得正式规划的产品主题
  - 主题 A:从首轮录入走向首轮建链
  - 主题 B:补齐 Decision 的经营生命周期
  - 主题 C:从人手工维护过渡到 agent-first 协作
- 第三层:应进入长期路线但不应立刻压进当前版本的方向
  - 经营侧方法论承接(Opportunity/Hypothesis/Experiment/Validation)
  - 生产侧真实代码协同
  - 跨项目全局规范资产统一维护

**关键贡献**:
- 明确肯定 PSCO 当前骨架正确、产品质感好、实体链路设计成功,问题不是"做错了"而是"还没从骨架走到闭环"
- 提出"首轮建链"概念,把 onboarding 的职责从"首轮录入"升级为"首轮建链引导"
- 提出 Decision 经营生命周期的最小语义:待决策 → 已形成正式结论 → 已回流到相关实体
- 明确指出 PSCO 长期定位应同时承接"经营系统"与"工程协同系统",因为它是 Personal Software Company OS

**建议其他专家重点回答的问题**:Onboarding 修复范围、Decision 生命周期设计、review 承接位、下一阶段主线优先级、agent-first 是否正式写入产品方向、跨项目规范资产是否纳入长期对象体系、PSCO 经营侧描述是否需要重新显性化。

### 1.6 Qwen3.7-Pro(`PSCO-mvp04-qwen37pro.md`)

**核心立场**:**PSCO 当前面临的不是"功能不够用"的问题,而是"战略定位与实际交付存在根本性错位"的问题**;下一阶段应正式重新校准 PSCO 战略定位。

**主线结构**:
- 方向 A:闭环修复(短期 P0)
- 方向 B:协调方式升级(中期 P1,从 web-first 升级为 agent-first)
- 方向 C:科学实验 workflow 最小承接(长期 P2,引入 Opportunity/Experiment 实体)

**关键贡献**:
- 提出"战略定位偏差"诊断:PSCO 当前是"以数据模型为中心",而非"以用户经营动作为中心"
- 把"协调方式升级"作为独立主线,而非依附于其他主线
- 给出科学实验 workflow 的具体实施路径:引入 Opportunity 实体(承接发现需求)、Experiment 实体(承接验证需求)、强化 Feedback 的"验证"语义
- 明确提出"是否重新定义 PSCO 战略定位"作为关键决策点

**明确建议**:产出 `PSCO-mvp04-summarize-feedback.md` 更新最终共识、正式讨论并仲裁 PSCO 战略定位、基于仲裁结论推进后续 phase。

---

## 2. 六份文档的共识点分析

经过交叉对比,六位专家在以下七个维度上形成了明确共识。这些共识是后续 `/plan` 必须承接的硬性输入。

### 2.1 共识一:Dry-Run 达成了设计目的,不是失败信号

**共识强度**:6/6 完全一致

六位专家均明确认同:`PSCO-mvp03-summarize-feedback.md §4.8` 定义的 dry-run 三项验证目标(真实使用摩擦、review loop 是否改变动作、模板复用是否缩短创造路径)都得到了真实证据,反馈揭示的问题恰恰是 dry-run 设计意图的达成。

**关键引用**:
- DPv4flash:"本轮 dry-run 完全达成设计目的"
- DPv4pro:"这是 PSCO 项目历史上第一次获得'系统在真实使用中的表现'这类数据"
- GLM52:"dry-run 不是证明 PSCO 完美,而是证明 PSCO 在真实使用中暴露了哪些 fixture 验收无法发现的问题"
- GPT54:"这轮反馈最值得珍惜的地方在于:它暴露的问题并不是'产品没有结构',而是'结构已经很好,但闭环没有彻底合上'"

### 2.2 共识二:PSCO 骨架正确,问题是"闭环未合上"而非"做错了"

**共识强度**:6/6 完全一致

六位专家均明确肯定 PSCO 当前骨架方向正确、产品质感克制、实体链路设计成功。问题不是"推翻重来",而是"补全闭环"。

**关键引用**:
- DPv4flash:"phase01 ~ phase09 证明了 PSCO『能登记、能复用、能看见复利』"
- DPv4pro:"PSCO 当前已经证明自己不是一个粗糙的 MVP 玩具"
- gemini31pro:"MVP0.3 成功证明了 PSCO 在『资产关联、来源追踪、留痕管理』上的工程扎实度"
- GPT54:"当前 PSCO 的结构方向是对的;当前 PSCO 的产品质感是好的"
- GLM52:"phase01 ~ phase09 已经证明 PSCO 的最小资产主线、经营回路与支撑能力都成立"
- qwen37pro:"PSCO phase01-09 已完成正式交付,证明了『资产登记 + 决策留痕 + 基础复用反馈』的最小主线成立"

### 2.3 共识三:三个阻断项必须立即通过 fix 闭合

**共识强度**:6/6 完全一致

六位专家均明确认同以下三个阻断项必须在进入 mvp0.4 主线之前通过 `fix*` workflow 闭合:

1. **Onboarding"开始首轮录入"按钮首次点击无响应**(feeback_01/02)
2. **Decision 状态闭环缺失,Dashboard/Review 持续误报"待处理"**(feeback_03/04)
3. **Decision Detail 缺少状态推进 CTA**(feeback_04)

**关键引用**:
- DPv4flash:给出具体 fix 编号 `fix_001/002/003` 并落到代码根因
- DPv4pro:"这些不是『功能增强』,而是『功能修复』"
- gemini31pro:"Dry-Run 中暴露的工程断点必须通过 fix 流程迅速收口"
- GLM52:"这 4 项必须修复,dry-run 才能判为『通过(含已知摩擦点)』并完成 mvp0.3 正式收口"
- GPT54:"这是一个必须立即修的 fix-now 问题"
- qwen37pro:"通过 fix* workflow 修复 P0 阻断项"

### 2.4 共识四:Dry-Run 收口不应无限延期

**共识强度**:5/6 明确一致(qwen37pro 未明确表态)

五位专家明确反对把 8 点反思全部解决后再收口 dry-run,主张 fix 闭合阻断项后即收口,已知摩擦点记入收口报告作为 mvp0.4 输入。

**关键引用**:
- DPv4flash:"fix 必须保持轻量(1~2 周内闭合),不能让阻断项演化为 phase 级工作"
- GLM52:"如果试图把 8 点反思全部解决后再收口 dry-run,会导致 dry-run 验收闸无限延期"
- GPT54:"如果连首轮入口都不稳定,那么后续关于经营系统、真实使用频率、agent-first 协作的讨论都会被削弱"

### 2.5 共识五:Onboarding 必须从"首轮录入"升级为"首轮建链"

**共识强度**:6/6 完全一致

六位专家均认同用户反思第二点(onboarding 六步是"假逻辑链")精准命中核心问题,下一阶段必须让实体之间的逻辑关系在创建时就被主动建立。

**关键引用**:
- DPv4flash:"onboarding 逻辑化(创建时主动建立实体关系,而非事后手动绑定)"
- DPv4pro:"在 Onboarding 流程中建立步骤间的数据传递:产品 → 仓库绑定 → 模块 → 决策"
- gemini31pro:"在向导中创建的 Product、Module、Repository,必须在向导结束时自动建立绑定关系"
- GLM52:"创建 Product 时主动询问『这个产品由哪些已有 Module 支撑?』"
- GPT54:"当前 Onboarding 的职责应从『首轮录入』进一步升级为『首轮建链引导』"
- qwen37pro:"提供『一站式创建产品及其关联』的聚合页面"

### 2.6 共识六:Decision 必须有正式完成出口与状态生命周期

**共识强度**:6/6 完全一致

六位专家均认同 Decision 当前"只能被创建和查看,不能被推进和完成"是结构性缺陷,必须补齐最小生命周期。

**关键引用**:
- DPv4flash:提出 `proposed → acknowledged → active → superseded → archived` 状态机
- DPv4pro:"为 Decision 引入最小状态机(proposed / accepted / rejected / superseded)"
- gemini31pro:"为 Decision 增加明确的生命周期状态...允许标记该决策为『已处理』"
- GLM52:提出 `draft → proposed → acknowledged → active → superseded → archived` 六态
- GPT54:"待决策 → 已形成正式结论 → 已回流到相关实体 之间的最小推进语义"
- qwen37pro:"提供『收口』动作或自动更新机制"

**状态机细节差异**:DPv4flash/GLM52 主张六态(含 `draft` 与 `acknowledged`),DPv4pro 主张四态(不含 `draft`),gemini31pro/GPT54/qwen37pro 未明确状态数。这是可仲裁的细节差异。

### 2.7 共识七:Agent-First 是重要战略方向,但当前应受控进入

**共识强度**:6/6 完全一致

六位专家均认同用户反思第七点(agent 维护为主、Web 展示为辅)是重要战略方向,但对当前 mvp0.4 是否立即全面进入存在差异。

**关键引用**:
- DPv4flash:"Agent Consumption Layer 不能演化为 AI 工作台(只读消费,不做内置对话/自动决策/代码生成)"
- DPv4pro:"Agent-Native 基础最后做...它是对前三层(登记 + 经营 + 连接)的程序化访问层"
- gemini31pro:"Agent 协同机制必须基于本地网络,不得强依赖公网网关"
- GLM52:"当前 PSCO 自身的资产-动作链路未闭合,先让 agent 消费未闭合的资产会放大混乱"
- GPT54:"下一阶段不一定马上实现完整 agent 接入层,但必须把这个方向正式上桌"
- qwen37pro:"将 PSCO 的协调方式从『web 端手动切换』升级为『agent 维护为主 + web 展示为辅』"

**差异**:GPT54 与 qwen37pro 倾向于把 agent-first 正式写入 mvp0.4 方向;DPv4flash/GLM52 主张作为支撑能力受控进入;DPv4pro 主张作为 mvp0.4 收口阶段(主线 C);gemini31pro 主张立即进入 MCP Server 落地。

---

## 3. 六份文档的分歧点辨析

六位专家在以下四个维度上存在明显分歧。这些分歧是后续 `/plan` 必须仲裁的关键决策点。

### 3.1 分歧一:MVP0.4 主线收敛策略

**分歧维度**:mvp0.4 应该是单一中心主线、三主线并列,还是三层推进?

| 专家 | 主张 | 理由 |
|---|---|---|
| DPv4flash | 单一中心主线(Asset-Action Closure)+ 支撑能力 + 候选探索 | 三主线并列会拉回"概念扩张先于使用闭环"的旧问题 |
| GLM52 | 单一中心主线(Asset-Action Closure)+ 支撑能力 + 候选探索 | 延续 mvp03 的"一主两翼一闸"结构 |
| DPv4pro | 三主线并列(闭环修复 + 真实连接 + Agent-Native 基础) | 三者解决不同层次问题,可以并行推进 |
| qwen37pro | 三方向(闭环修复 + 协调方式升级 + 科学实验 workflow) | 战略定位偏差需要多维度同时修正 |
| GPT54 | 三层推进(第一层 fix + 第二层主题 + 第三层长期路线) | 分层而非并列,避免一次性压太多 |
| gemini31pro | 三主线(Agent-Driven + 跨项目资产 + 精益创业) | 范式转移需要多维度同步 |

**分析**:

- DPv4flash 与 GLM52 立场高度一致,主张收敛,强调"先闭合再扩展"
- DPv4pro 与 qwen37pro 立场相近,主张多主线并行,但 qwen37pro 更激进(引入新实体)
- GPT54 主张三层推进,本质是"分层处理"而非"并列主线",与 DPv4flash/GLM52 的分层框架有相通之处
- gemini31pro 最激进,主张范式转移

**仲裁建议**:

**采纳 DPv4flash/GLM52 的收敛立场**,理由:

1. 单人维护 + AI 协作的工程现实不允许同时推进三条并列主线
2. mvp0.3 的 dry-run 反馈恰恰证明"概念扩张先于使用闭环"会导致语义断层
3. GPT54 的"三层推进"本质与 DPv4flash/GLM52 的"分层处理"相通,可以调和

但**吸纳 DPv4pro 的"真实连接"作为 Asset-Action Closure 的组成部分**(而非独立主线),以及**吸纳 GPT54 的"agent-first 正式上桌"作为 Agent Consumption Layer 的方向锚定**。

### 3.2 分歧二:是否引入新实体(Opportunity/Experiment/Venture)

**分歧维度**:mvp0.4 是否应该引入 Opportunity、Experiment、Venture 等新实体?

| 专家 | 主张 | 理由 |
|---|---|---|
| DPv4flash | 不引入 | Venture/Opportunity/Feature/Experiment 继续后移 |
| GLM52 | 不引入 | 同 DPv4flash,完全遵守 mvp03 后移项清单 |
| GPT54 | 不引入(但承认长期路线) | 第三层"应进入长期路线但不应立刻压进当前版本" |
| DPv4pro | 不引入(Venture 后移到 mvp0.5+) | 用户核心痛点不在此 |
| qwen37pro | 引入(Opportunity/Experiment,P2 长期) | 科学实验 workflow 完全缺失是深层不满来源 |
| gemini31pro | 引入(Venture/Experiment/Opportunity) | Personal Software Company 包含"公司"二字,不仅是写代码 |

**分析**:

- 4/6 主张不引入(DPv4flash/GLM52/GPT54/DPv4pro)
- 2/6 主张引入(qwen37pro 作为 P2 长期,gemini31pro 作为主线)
- gemini31pro 最激进,主张立即解锁并升格

**仲裁建议**:

**采纳多数意见,mvp0.4 不引入新实体**,理由:

1. 4/6 专家明确反对立即引入
2. 用户反思第五点(科学实验 workflow)是"初心回溯",不是"当前阻断项"
3. 引入新实体需要完整的 CRUD + 关联 + review + 复用链路,工程量巨大
4. mvp0.4 的核心任务是闭合已有断层,不是扩张对象宽度

但**承认 qwen37pro 与 gemini31pro 指出的"科学实验 workflow 完全缺失"是真实问题**,作为 mvp0.5+ 的核心方向正式留档。

### 3.3 分歧三:真实连接(工程化协同)的深度与方式

**分歧维度**:PSCO 应该如何"触及真实开发过程"?

| 专家 | 主张 | 理由 |
|---|---|---|
| DPv4pro | GitHub API 读时获取仓库元数据(星标、语言、commit) | 让 PSCO 从"文本描述者"升级为"连接真实开发过程的系统" |
| DPv4flash | 不深入工程化,通过 Agent Consumption Layer 桥接 | PSCO 不应该自己变成工程工具 |
| GLM52 | 同 DPv4flash,PSCO 通过 agent layer 让外部工具消费上下文 | PSCO 不应该执行代码、不应该是 IDE |
| GPT54 | 长期方向,但当前不压进 mvp0.4 | 文本化事实层在 MVP 阶段是正确的 |
| qwen37pro | 最小联动(如 Git 仓库状态读取),再逐步深化 | 评估投入产出比,分阶段推进 |
| gemini31pro | Agent 通过 MCP 自动调用 PSCO 接口完成登记与绑定 | Agent 侧重写与微观执行 |

**分析**:

- DPv4pro 主张最具体的工程化连接(GitHub API 元数据回显)
- DPv4flash/GLM52/GPT54 主张通过 agent layer 间接连接,而非 PSCO 自己做工程化
- qwen37pro 主张最小联动
- gemini31pro 主张 agent 写入侧的自动化

**仲裁建议**:

**采纳 DPv4flash/GLM52 的"通过 Agent Consumption Layer 桥接"立场**,但**吸纳 DPv4pro 的"读时派生、不持久化"原则作为 Agent Consumption Layer 的实现约束**,理由:

1. PSCO 的核心定位是"经营与资产系统",不是"工程工具"
2. GitHub API 元数据回显是"工程化连接"的一种,但不是唯一方式
3. Agent Consumption Layer 可以让外部 agent(cursor/trae/codex)消费 PSCO 上下文后,由 agent 自己完成工程化加速
4. 这样 PSCO 保持轻量,工程化加速由专业工具承接

但**承认 DPv4pro 的"真实连接"诊断是深刻的**,如果 Asset-Action Closure 完成后用户仍感觉"文本描述者"痛点未缓解,可在 mvp0.4 尾部或 mvp0.5+ 引入 DPv4pro 的 GitHub API 读时派生方案作为补充。

### 3.4 分歧四:Agent-First 在 mvp0.4 的落地深度

**分歧维度**:agent-first 在 mvp0.4 应该做到什么程度?

| 专家 | 主张 | 范围 |
|---|---|---|
| DPv4flash | 受控最小版 | 只读 agent 友好接口 + AGENTS.md 风格上下文导出 + 上下文快照 |
| GLM52 | 受控最小版 | 同 DPv4flash |
| DPv4pro | Agent-Native 基础(主线 C) | Agent Context Query API + 跨项目知识基线 + Agent 可写 API 完善 |
| GPT54 | 正式上桌,但不马上实现 | 把方向正式写入产品方向,讨论哪些事实适合 agent 写入 |
| qwen37pro | 协调方式升级(方向 B) | agent 友好 API + web 端必要操作入口优化 + 跨项目规范统一维护 |
| gemini31pro | Agent-Driven Workflow(主线一) | MCP Server + Agent 自动调用 PSCO 接口完成登记与绑定 |

**分析**:

- DPv4flash/GLM52 主张最克制(只读消费)
- DPv4pro 主张中等(Agent Context Query API + 可写 API 完善)
- GPT54 主张方向上桌但实现延后
- qwen37pro 主张协调方式升级(含 web 端改造)
- gemini31pro 主张最激进(MCP Server + 自动写入)

**仲裁建议**:

**采纳 DPv4flash/GLM52 的"受控最小版"立场**,理由:

1. 当前 PSCO 自身的资产-动作链路未闭合,先让 agent 消费未闭合的资产会放大混乱
2. Agent Consumption Layer 一旦做重,容易演化为 AI 工作台,违背 `PSCO_3.md` 的 AI 增强层定位
3. `.proto + ConnectRPC` 主线已经具备机器可消费能力,真正缺的是"agent 友好的上下文输出层"

但**吸纳 GPT54 的"正式上桌"主张**,在 mvp0.4 正式 `/plan` 中明确把 agent-first 作为方向锚定,即使当前只做受控最小版。同时**吸纳 DPv4pro 的 Agent Context Query API 设计**作为受控最小版的具体实现方向。

---

## 4. 综合判断与主线收敛

基于六份专家文档的共识与分歧分析,我给出如下综合判断:

### 4.1 核心判断

> **六份专家文档在"骨架正确、闭环未合上"这一核心诊断上完全一致,在"如何收敛 mvp0.4 主线"上存在分歧。分歧的本质不是"方向之争",而是"节奏之争"——是先闭合再扩展,还是多维度同步推进。**
>
> **基于单人维护 + AI 协作的工程现实,以及 mvp0.3 dry-run 反馈恰恰证明"概念扩张先于使用闭环"会导致语义断层这一历史教训,我采纳"先闭合再扩展"的收敛立场,以 Asset-Action Closure 为中心主线,Agent Consumption Layer 为支撑能力,Cross-Project Convention Asset 为候选探索。**

### 4.2 收敛后的主线结构

```
Asset-Action Closure(中心主线)
   +
Agent Consumption Layer(支撑能力,受控最小版)
   +
Cross-Project Convention Asset(候选探索,严格最小版或后移)
```

这个结构与 DPv4flash/GLM52 的主张完全一致,同时吸纳了:

- DPv4pro 的"真实连接"诊断(作为 Asset-Action Closure 的组成部分,而非独立主线)
- DPv4pro 的 Agent Context Query API 设计(作为 Agent Consumption Layer 的实现方向)
- GPT54 的"首轮建链"概念(作为 Asset-Action Closure 的核心子任务)
- GPT54 的"agent-first 正式上桌"主张(作为方向锚定,即使当前只做受控最小版)
- qwen37pro 的"战略定位偏差"诊断(作为 mvp0.4 必须回应的战略议题)
- gemini31pro 的"Local-First 延续"约束(作为 Agent Consumption Layer 的硬性约束)

### 4.3 与分歧点的对应关系

| 分歧点 | 仲裁结论 | 理由 |
|---|---|---|
| 主线收敛策略 | 单一中心主线 + 支撑 + 候选 | 单人维护现实 + 历史教训 |
| 是否引入新实体 | 不引入 | 4/6 专家反对 + 工程量巨大 + 非当前阻断项 |
| 真实连接深度 | 通过 Agent Consumption Layer 桥接,不直接做工程化 | PSCO 核心定位是经营与资产系统 |
| Agent-First 落地深度 | 受控最小版 + 方向正式上桌 | 资产链路未闭合前不宜让 agent 深度消费 |

---

## 5. 推进顺序建议

### 5.1 推荐的推进路径

```
第一步:fix* 闭合 3 个阻断项
  ├─ fix_001_onboarding_cold_start_state
  ├─ fix_002_decision_status_semantics
  └─ fix_003_decision_detail_status_advance_cta
        ↓
第二步:重跑一轮 dry-run(只验证 fix 是否解决阻断项)
        ↓
第三步:dry-run 判"通过(含已知摩擦点)",正式收口 mvp0.3
  └─ 已知摩擦点记入收口报告,作为 mvp0.4 正式输入
        ↓
第四步:建立 mvp0.4 正式 /plan 入口
  ├─ 中心主线:Asset-Action Closure
  │   ├─ onboarding 逻辑化(首轮建链)
  │   ├─ Decision 状态机细化(draft/proposed/acknowledged/active/superseded/archived)
  │   ├─ Cross-Entity CTA 矩阵(详情页主动提示下一步)
  │   ├─ review loop 业务闭环(review → Decision → 实体更新 → review 已完成)
  │   └─ Current Focus 信号派生规则重做
  ├─ 支撑能力:Agent Consumption Layer(受控最小版)
  │   ├─ 只读 agent 友好接口(基于既有 ConnectRPC)
  │   ├─ AGENTS.md 风格上下文导出
  │   └─ 上下文快照机制
  └─ 候选探索:Cross-Project Convention Asset(严格最小版或后移)
        ↓
第五步:mvp0.4 正式验收
  └─ 必须回答:用户能否一次性形成完整实体关系网?Decision 能否推进并回流?每个详情页是否提示下一步?Dashboard 是否只指向真正待处理?外部 agent 能否读取 Product 完整上下文?
```

### 5.2 推进顺序的依赖关系论证

1. **没有第一步,dry-run 无法收口,mvp0.3 不能正式完成**:阻断项会持续干扰后续判断
2. **没有第二步,fix 效果无法验证**:只修不验等于没修
3. **没有第三步,mvp0.4 正式入口无法建立**:没有收口就没有新 phase
4. **没有第四步中心主线,PSCO 仍停留在"实体登记完备但动作链路断裂"**:dry-run 反馈的核心问题未解决
5. **没有第四步支撑能力,用户反思中的"工程化配合"无法落地**:agent-first 方向无法落地
6. **候选探索是可选项,不阻断 mvp0.4 主线收口**

### 5.3 核心纪律

- **fix 必须保持轻量(1~2 周内闭合)**:如果 fix 工作量超过 1-2 周,说明对问题根因的判断有误
- **不预设 phase 命名**:本文只给方向,不冻结 `phase10` 名称或 `mvp0.4+` 范围
- **不把 8 点反思全部塞进 dry-run 收口**:已知摩擦点记入收口报告即可

---

## 6. 风险与守界纪律

### 6.1 五大工程风险与缓解

| 风险 | 级别 | 缓解措施 |
|---|---|---|
| Agent Consumption Layer 被偷渡成 AI 工作台 | 🔴 高 | 只读消费,不做内置对话/自动决策/代码生成/agent 双向同步 |
| Onboarding 逻辑化被偷渡成工作流引擎 | 🔴 高 | 仍是 6 步引导,关系建立是主动询问而非强制依赖,不做 onboarding 模板或版本管理 |
| Cross-Project Convention Asset 被偷渡成知识图谱 | 🟡 中 | 只登记工程可消费的规范资产(技术栈、目录结构、agent 规则、协作约束),不做书籍笔记/语义搜索/依赖图谱 |
| Decision 状态机细化被偷渡成 Decision Intelligence | 🟡 中 | 只细化到能区分"已留痕/已确认/已生效",不做相似匹配/引用推荐/AI 增强 |
| Review loop 被偷渡成通用任务管理器 | 🟡 中 | review 动作必须锚定既有实体,不做通用任务分配/调度 |

### 6.2 五条守界纪律(对后续 `/spec` 与实现的硬约束)

1. `.proto` 仍是唯一长期合同源
2. query 层继续保持纯只读
3. 前端写路径继续收敛到切片固定承接位
4. 不引入第二套路由、第二套状态管理、第二套 UI 事实源
5. **agent-first 方向正式上桌,但 mvp0.4 只做受控最小版**——方向锚定不等于立即全面实现

### 6.3 对战略定位议题的处理

qwen37pro 提出的"战略定位与实际交付存在根本性错位"诊断是深刻的,六位专家中 3/6(DPv4pro/qwen37pro/gemini31pro)明确认同这一诊断。

但**这个议题不应在 mvp0.4 阶段直接仲裁**,理由:

1. 战略定位重新校准需要完整的原始方案文档(`PSCO_0.md ~ PSCO_4.md`)重新评审,工程量巨大
2. mvp0.4 的核心任务是闭合已有断层,不是重新定义战略定位
3. 如果 Asset-Action Closure 完成后用户仍感觉"战略定位偏差"未缓解,可在 mvp0.5+ 正式启动战略定位重新校准

**建议**:把 qwen37pro 的"战略定位偏差"诊断作为 mvp0.4 正式 `/plan` 的上游输入之一,但不作为 mvp0.4 的直接主线。在 mvp0.4 收口时,基于 Asset-Action Closure 的实际效果评估是否需要启动战略定位重新校准。

---

## 7. 对后续正式 `/plan` 的输入

### 7.1 推荐的候选阶段结构(不冻结正式 phase 命名)

```
候选阶段一:dry-run 阻断项 fix 与收口
  ├─ fix_001_onboarding_cold_start_state
  ├─ fix_002_decision_status_semantics
  ├─ fix_003_decision_detail_status_advance_cta
  └─ 重跑 dry-run → mvp0.3 正式收口
        ↓
候选阶段二:Asset-Action Closure 主线
  ├─ onboarding 逻辑化(首轮建链)
  ├─ Decision 状态机细化
  ├─ Cross-Entity CTA 矩阵
  ├─ review loop 业务闭环
  └─ Current Focus 信号派生规则重做
        ↓
候选阶段三:Agent Consumption Layer 支撑能力
  ├─ agent 友好只读接口(基于既有 ConnectRPC)
  ├─ AGENTS.md 风格上下文导出
  └─ 上下文快照机制
        ↓
候选阶段四(可选):Cross-Project Convention Asset 候选探索
  └─ 严格最小版,视前置阶段完成情况决定是否进入
```

### 7.2 推荐的先后关系

1. 没有候选阶段一,dry-run 无法收口,mvp0.3 不能正式完成
2. 没有候选阶段二,PSCO 仍停留在"实体登记完备但动作链路断裂"状态
3. 没有候选阶段三,PSCO 资产无法被外部 agent 消费,用户反思中的"工程化配合"无法落地
4. 候选阶段四是可选项,不阻断 mvp0.4 主线收口

### 7.3 对正式 `/spec` 的硬约束

后续 `/spec` 与实现必须继续遵守:

1. `.proto` 仍是唯一长期合同源
2. query 层继续保持纯只读
3. 前端写路径继续收敛到切片固定承接位
4. 不引入第二套路由、第二套状态管理、第二套 UI 事实源
5. onboarding 逻辑化不能演化为可配置工作流引擎
6. Agent Consumption Layer 不能演化为 AI 工作台
7. Cross-Project Convention Asset 不能演化为知识图谱
8. Decision 状态机细化不能演化为 Decision Intelligence

### 7.4 对后续验收的硬要求

下一阶段正式验收至少必须回答以下问题:

1. 用户从 onboarding 进入,能否一次性形成完整实体关系网(产品 + 仓库 + 模块 + 决策)?
2. Decision Detail 能否完成 `proposed → acknowledged → active` 的状态推进,并回流到 review?
3. 每个实体详情页是否都主动提示"下一步应该做什么"?
4. Dashboard Current Focus 是否只指向真正待处理的 Decision?
5. 外部 agent 能否通过标准化接口读取 Product 的完整上下文(含关联 Module / Decision / Repository)?
6. 真实项目中,用户是否愿意围绕 PSCO 持续回来使用,且使用摩擦显著低于 mvp0.3 dry-run?

### 7.5 明确不做范围(延续 mvp03 后移项清单)

- `Venture` 正式主线化
- `Opportunity / Feature / Experiment` 流程化
- `Capability` 重实体 CRUD
- GitHub OAuth / 自动导入
- AI 一级工作台或内置 AI 对话
- 自动扫描 / 知识图谱 / 语义搜索
- Rust Intelligence Layer
- 完整模板平台 / 模板版本管理
- `Decision Intelligence`(相似匹配、引用推荐、AI 上下文增强)
- PSCO 自身的 IDE 集成或代码执行
- onboarding 可配置工作流引擎
- "探索型"场景(假设管理、实验设计、结果追踪)的完整实现

---

## 8. 六份专家文档的评价与定位

### 8.1 各份文档的独特价值

| 文档 | 独特价值 | 适用场景 |
|---|---|---|
| DPv4flash | 分层处理框架 + 代码根因定位 + 守界纪律 | fix 阶段直接落地 |
| DPv4pro | 真实连接具体方案 + Agent Context Query API 设计 + 对自身路线的公开校准 | Agent Consumption Layer 实现参考 |
| gemini31pro | 读写倒置范式 + 跨项目资产下发机制 + 精益创业模型数据结构 | mvp0.5+ 战略参考 |
| GLM52 | Decision 状态机六态 + Cross-Entity CTA 矩阵 + 5 条守界纪律 | Asset-Action Closure 实现参考 |
| GPT54 | "骨架正确"定调 + "首轮建链"概念 + agent-first 正式上桌主张 | mvp0.4 战略锚定 |
| qwen37pro | "战略定位偏差"诊断 + 协调方式升级独立主线 + 科学实验 workflow 实施路径 | mvp0.5+ 战略参考 |

### 8.2 各份文档的局限性

| 文档 | 局限性 |
|---|---|
| DPv4flash | 对真实连接(工程化协同)的处理过于保守,未充分回应反思第四点 |
| DPv4pro | 三主线并列在单人维护现实下工程量过大,且 Agent-Native 基础依赖前两条主线 |
| gemini31pro | 主张过于激进,立即引入 Venture/Experiment/Opportunity 会重蹈"概念扩张先于使用闭环"覆辙 |
| GLM52 | 对战略定位议题的处理过于保守,未充分回应 qwen37pro 的"战略定位偏差"诊断 |
| GPT54 | 三层推进结构在操作层面不够具体,需要其他文档补充实现细节 |
| qwen37pro | 引入新实体的工程量评估不足,且"协调方式升级"与 Asset-Action Closure 存在范围重叠 |

### 8.3 文档间的关系定位

- **DPv4flash 与 GLM52 立场高度一致**:作为收敛立场的共同主张者
- **DPv4pro 与 qwen37pro 立场相近**:作为多主线并行主张者,但 qwen37pro 更激进
- **GPT54 居中调停**:三层推进本质与 DPv4flash/GLM52 的分层框架相通,但吸纳了 agent-first 正式上桌
- **gemini31pro 独立激进**:主张范式转移,作为 mvp0.5+ 战略参考价值最大

---

## 9. 最终结论

### 9.1 一句话总结

> **六份专家文档在"骨架正确、闭环未合上"这一核心诊断上完全一致;分歧在于"如何收敛 mvp0.4 主线"。基于单人维护现实与历史教训,我采纳"先闭合再扩展"的收敛立场——以 Asset-Action Closure 为中心主线,Agent Consumption Layer 为支撑能力,Cross-Project Convention Asset 为候选探索;同时把"战略定位偏差"诊断与"科学实验 workflow 缺失"作为 mvp0.5+ 的核心方向正式留档。**

### 9.2 核心仲裁结论

1. **第一优先级**:走 `fix*` workflow 闭合 3 个阻断项,重跑 dry-run,完成 mvp0.3 正式收口
2. **第二优先级**:以 Asset-Action Closure 为中心主线推进 mvp0.4(onboarding 逻辑化 + Decision 状态机细化 + Cross-Entity CTA 矩阵 + review loop 业务闭环 + Current Focus 信号派生重做)
3. **第三优先级**:以 Agent Consumption Layer 为支撑能力(受控最小版:只读 agent 友好接口 + AGENTS.md 风格上下文导出 + 上下文快照机制)
4. **第四优先级(可选)**:Cross-Project Convention Asset 候选探索(严格最小版或后移到 mvp0.5+)
5. **后移到 mvp0.5+**:`Venture` / `Opportunity` / `Experiment` / `Decision Intelligence` / `AI Context Enhancement` / 完整模板平台 / GitHub OAuth / 自动扫描 / 战略定位重新校准 / 科学实验 workflow 完整实现

### 9.3 核心原则

1. **先修后建**:阻断项修复优先于任何新能力建设
2. **闭合优先于扩张**:在增加新实体类型之前,先让已有实体形成动作闭环
3. **Agent-Native 优先于 AI-Enhanced**:在让 AI 辅助人类操作之前,先让 PSCO 具备 agent 可编程访问的基础
4. **读时派生优先于持久化**:Agent Consumption Layer 的上下文导出不写入 PSCO 数据库,保持系统轻量
5. **方向上桌优先于立即实现**:agent-first 方向正式写入产品方向,但 mvp0.4 只做受控最小版
6. **继续坚持技术栈冻结**:不引入新基础设施、不引入第二套事实源

### 9.4 对后续专家讨论的建议

如果需要第二轮专家讨论,建议重点回答以下问题(吸纳 GPT54 的建议):

1. Onboarding 逻辑化的最小实现方案是什么?(是否需要步骤间数据传递,还是只需要结束时统一建立关系?)
2. Decision 状态机应该是四态还是六态?(`acknowledged` 中间态是否必要?`draft` 草稿态是否必要?)
3. Cross-Entity CTA 矩阵的具体范围是什么?(每个详情页应该有哪些 CTA?如何避免 CTA 过载?)
4. Agent Consumption Layer 的上下文导出格式应该是什么?(Markdown?JSON?Protocol Buffers?)
5. Cross-Project Convention Asset 是否应该在 mvp0.4 进入?如果进入,最小范围是什么?
6. 是否需要在 mvp0.4 正式启动"战略定位重新校准"议题?还是延后到 mvp0.5+?

### 9.5 对用户的直接建议

基于六份专家文档的综合判断,我建议您:

1. **立即启动 fix* workflow 闭合 3 个阻断项**(1~2 周内闭合)
2. **重跑一轮 dry-run 验证 fix 效果**(只验证阻断项是否解决)
3. **正式收口 mvp0.3**,把已知摩擦点记入收口报告
4. **建立 mvp0.4 正式 /plan 入口**,以 Asset-Action Closure 为中心主线
5. **把 agent-first 方向正式写入 mvp0.4 产品方向**,但只做受控最小版
6. **把战略定位重新校准与科学实验 workflow 作为 mvp0.5+ 核心方向正式留档**,不在 mvp0.4 启动

如果用一句话收口这份汇总:

> **六位专家在"骨架正确"上完全一致,在"如何收敛"上存在节奏之争;基于单人维护现实与历史教训,采纳"先闭合再扩展"立场——mvp0.4 以 Asset-Action Closure 为中心主线,Agent Consumption Layer 为支撑能力,战略定位与科学实验 workflow 后移到 mvp0.5+。**

---

## 附录 A:六份专家文档主线结构对照表

| 专家 | 中心主线 | 支撑/并行 | 候选/长期 | 是否引入新实体 |
|---|---|---|---|---|
| DPv4flash | Asset-Action Closure | Agent Consumption Layer(受控) | Cross-Project Convention Asset(候选) | 否 |
| GLM52 | Asset-Action Closure | Agent Consumption Layer(受控) | Cross-Project Convention Asset(候选) | 否 |
| DPv4pro | 闭环修复(主线 A) | 真实连接(主线 B)+ Agent-Native(主线 C) | — | 否(Venture 后移) |
| GPT54 | (第一层 fix) | 首轮建链 + Decision 生命周期 + agent-first(第二层主题) | 经营方法论 + 真实代码协同 + 跨项目规范(第三层) | 否(长期路线) |
| qwen37pro | 闭环修复(方向 A) | 协调方式升级(方向 B) | 科学实验 workflow(方向 C) | 是(Opportunity/Experiment,P2) |
| gemini31pro | Agent-Driven Workflow(主线一) | 跨项目资产化(主线二) | 精益创业模型(主线三) | 是(Venture/Experiment/Opportunity) |

## 附录 B:六份专家文档对用户八点反思的处理对照表

| 用户反思 | DPv4flash | DPv4pro | gemini31pro | GLM52 | GPT54 | qwen37pro |
|---|---|---|---|---|---|---|
| 1. 精良但未闭环 | Asset-Action Closure | 闭环修复 | Fix & Improvement | Asset-Action Closure | 第一层 fix | 闭环修复 |
| 2. onboarding 假逻辑链 | onboarding 逻辑化 | 步骤间数据传递 | 自动建立绑定关系 | onboarding 逻辑化 | 首轮建链 | 一站式聚合页面 |
| 3. 详情页频繁切换 | Cross-Entity CTA 矩阵 | 引导式关联 | — | Cross-Entity CTA 矩阵 | — | web 端必要操作入口优化 |
| 4. 仅文本描述 | Agent Consumption Layer | 真实连接(GitHub API) | Agent 自动调用 | Agent Consumption Layer | 长期方向 | 最小联动 |
| 5. 用户初心回溯 | 不扩张对象 | mvp0.5+ 方向 | 精益创业模型 | 不扩张对象 | 第三层长期路线 | 科学实验 workflow |
| 6. 中间地带困境 | 双层缓解 | 真实连接深化 | 范式转移 | 双层缓解 | 同时承接经营与工程 | 战略定位偏差 |
| 7. agent 协同预想 | Agent Consumption Layer(受控) | Agent-Native 基础 | Agent-Driven Workflow | Agent Consumption Layer(受控) | agent-first 正式上桌 | 协调方式升级 |
| 8. 跨项目规范复用 | Cross-Project Convention(候选) | 跨项目知识基线 | 跨项目资产化 | Cross-Project Convention(候选) | 第三层长期路线 | 跨项目规范统一维护 |

---

**Final Record**
**Signed by: GLM-5.2**
**Date: 2026-08-13**
**Document Type:** `review`
**Status:** 供后续正式 `/plan` 与最终共识仲裁参考
