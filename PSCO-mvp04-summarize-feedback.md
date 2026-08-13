# Personal Software Company OS

# MVP0.4 Final Consensus & Planning Baseline

**Author:** GPT54
**Date:** 2026-08-13
**Purpose:** 基于两轮 `mvp0.4` 专家评审、交叉汇总文档、真实 `dry-run` 反馈与当前仓库现实，形成 PSCO 在 `phase09` 收口之后、进入下一阶段正式 `/plan` 之前的最终仲裁结论、范围基线、推进顺序与明确非目标。

---

## 1. 文档定位

本文不是第十一份专家评审意见，也不是直接替代后续正式 `phase` `/plan` 的执行文档。

本文要解决的是六件事：

1. 明确哪些判断已经形成稳定共识；
2. 对仍有分歧的方向给出最终仲裁；
3. 冻结 `mvp0.4` 的正式方向、范围与非目标；
4. 明确 `mvp0.3` 应如何通过 `fix + rerun dry-run` 正式收口；
5. 约束下一阶段正式 `/plan` 的输入边界与推进顺序；
6. 为后续 `/plan -> /spec -> 实现 -> 验收 -> 收口` 提供唯一上游判断基线。

本文的职责，是为下一阶段正式 phase 入口建立提供上游依据，而不是在这里越权冻结正式 `phase` 名称、`.trae/specs/phaseXX_*` 路径、接口名或实现细节。

## 1.1 本文与 `PSCO-mvp03-summarize-feedback.md` 的关系

本文与 `PSCO-mvp03-summarize-feedback.md` 的关系，明确冻结为：

1. `PSCO-mvp03-summarize-feedback.md` 已经完成了从 `mvp0.2` 阶段二走向 `mvp0.3` 正式主题的总方向仲裁，并冻结了：
   - `Operating Review Loop` 为中心主线；
   - `Template Reuse` 与 `Derived Intelligence Deepening` 为支撑；
   - `Real-Project Dry-Run` 为独立验收闸。
2. `phase08` 与 `phase09` 已经实际完成了上述主线的大部分正式交付；
3. `Real-Project Dry-Run` 已经产出真实使用反馈，第一次把 fixture 验收无法暴露的问题拉到了台面上；
4. 本文要解决的，不是再次仲裁 `phase08 / phase09` 做得对不对，而是基于这轮真实反馈，正式收敛 `mvp0.4` 的执行主题、范围基线与完整路线表达；
5. 因此，本文既承接 `mvp0.3` 的正式共识，也新增冻结了下一步必须优先处理的内容：`dry-run` 阻断项修复、`Asset-Action Closure`、受控 `Agent Consumption Layer`，以及对候选长期方向的时机仲裁。

换句话说：

> **`mvp0.3` 解决的是“PSCO 能否形成 review / reuse / derived hint / dry-run 的可运行经营骨架”；`mvp0.4` 要解决的是“这些骨架能力能否真正闭合成可持续消费的动作链，并开始具备 agent 可消费性”。**

---

## 2. 证据来源

## 2.1 第一轮方向评审

- `docs/review/PSCO-mvp04-DPv4flash.md`
- `docs/review/PSCO-mvp04-DPv4pro.md`
- `docs/review/PSCO-mvp04-gemini31pro.md`
- `docs/review/PSCO-mvp04-GLM52.md`
- `docs/review/PSCO-mvp04-GPT54.md`
- `docs/review/PSCO-mvp04-qwen37pro.md`

## 2.2 第二轮交叉汇总

- `docs/review/PSCO-mvp04-summarize-feedback-DPv4pro.md`
- `docs/review/PSCO-mvp04-summarize-feedback-gemini31pro.md`
- `docs/review/PSCO-mvp04-summarize-feedback-GLM52.md`
- `docs/review/PSCO-mvp04-summarize-feedback-GPT54.md`

## 2.3 本文额外采用的直接证据

- `docs/review/PSCO-real-project-dry-run-user-manual-GPT54 feedback.md`
- 当前根级真相源与项目规则：
  - `AGENTS.md`
  - `plan.md`
  - `project_rules.md`
  - `TECH_STACK_BASELINE.md`
- 当前正式仓库现实与此前已核实的实现主线：
  - `Onboarding`
  - `Review`
  - `Decision Center`
  - `Dashboard`
  - `Template Reuse`

## 2.4 本文采用的仲裁原则

1. **真实使用证据优先于理论推演。** `dry-run` 暴露的问题比单纯概念讨论更接近系统真实状态。
2. **高共识优先于单点激进扩张。** 四方以上独立收敛的判断，优先视为稳定输入。
3. **闭合优先于扩张。** 在已有资产与 review 能力尚未形成真实动作闭环前，不应继续扩对象宽度。
4. **时机优先于愿景。** 某个长期方向即使正确，只要当前前提未成立，就不能提前升格为下一阶段主线。
5. **消费优先于集成。** 先证明现有资产能被持续消费，再决定是否追加更深的工程连接与更重的智能层。
6. **agent 能力优先做只读消费层。** 在当前阶段，程序化消费是对的，程序化写入不是主线。
7. **评审不越权冻结正式 phase。** 正式阶段命名、spec 路径与具体落点留给后续 `/plan` 与 `/spec` 决定。

---

## 3. 最终共识

经过两轮评审与交叉汇总，以下内容已经形成稳定共识，可直接视为 `mvp0.4` 的正式前提。

## 3.1 关于这轮 `dry-run` 性质的共识

1. `phase01 ~ phase09` 已经证明 PSCO 的最小资产主线、review 主线、模板复用与派生提示主线成立；
2. 这轮 `dry-run` 是成功的，它达成了“暴露真实使用问题”的设计目的；
3. `dry-run` 揭示的不是“功能完全缺失”，而是“实体登记完备”与“经营动作闭环”之间存在语义断层；
4. 这个断层是此前 fixture 验收很难发现、只有真实使用才能暴露的问题，因此它不是失败信号，而是高价值证据。

## 3.2 关于当前项目状态的共识

1. 当前系统的骨架方向是对的，产品质感也是成立的；
2. 当前系统已经不是单纯的 CRUD 原型，而是具备了经营系统雏形；
3. 当前真正缺的不是更多实体，而是“从创建到下一步动作”的单值承接；
4. 因此下一步的关键问题，不是“还能加什么概念”，而是“已有能力能不能被真实经营动作持续消费”。

## 3.3 关于必须立即修复的阻断项的共识

1. `Onboarding` 欢迎页“开始首轮录入”首次点击无响应；
2. `Decision` 状态语义错位，`proposed` 同时承担“已留痕”与“待处理”双重语义；
3. `Decision Detail` 缺少状态推进 CTA，导致 `Review -> Decision -> 实体更新` 业务回流断裂；
4. `Dashboard / Daily Review / Current Focus` 对 pending signal 的判定必须与真实消费状态重新对齐。

## 3.4 关于下一步总方向的共识

1. 下一步不应继续扩长期对象宽度；
2. 下一步应优先闭合资产与动作之间的断层；
3. `Agent` 方向应正式进入，但进入方式必须克制；
4. 跨项目规范资产值得保留为候选探索，但不应越权成为当前第一主线；
5. `Venture / Opportunity / Experiment / Decision Intelligence / AI Context Enhancement` 不应直接进入当前正式主线。

## 3.5 关于范围边界的共识

1. `.proto` 仍是唯一长期合同源；
2. `ConnectRPC` 仍是业务传输主线；
3. Go 后端继续承载 PSCO 的核心能力与正式业务能力暴露面；
4. React 前端继续作为更友好、更可视化的消费与操作渠道，而不是系统能力的唯一正式入口；
5. 面向 agent 的正式接入通道，应优先落在后端暴露的 `MCP / CLI / API` 能力面，而不是前端页面交互层；
6. 前端 `query` 层继续保持纯只读；
7. 前端写路径继续收敛到切片固定承接位；
8. 不引入第二套路由、第二套状态管理、第二套 UI 事实源；
9. 不做 AI 一级工作台、不做自动扫描 / 知识图谱、不做 GitHub OAuth / 自动导入、不做 Rust Intelligence Layer。

---

## 4. 最终仲裁结论

## 4.1 `mvp0.4` 的版本语义

**最终结论：**

> `mvp0.4` 的实质任务，不是开启一轮新的概念扩张，也不是把 PSCO 立即重写成完整 agent-first 平台；
> 而是把 `dry-run` 已经真实暴露的语义断层闭合掉，让 PSCO 从“能登记、能复用、能看见”正式跨到“能闭环、能回流、能被 agent 稳定消费”。

这意味着我正式采纳以下判断：

1. `mvp0.3` 必须先通过 `fix + rerun dry-run` 正式收口；
2. `mvp0.4` 的中心问题是动作链闭合，不是长期对象回归；
3. agent 能力应作为支撑层上桌，但当前不升格为唯一中心主线；
4. “科学实验 workflow”“完整工程连接”“战略定位重写”都是真实长期议题，但当前时机未成熟。

进一步冻结为：

> **`mvp0.4` 不是重新定义 PSCO，而是让 PSCO 先把已经存在的经营骨架真正跑通，再建立最小的 agent 可消费底座。**

## 4.2 `mvp0.4` 的总主题

**最终结论：**

> `mvp0.4` 的总主题确定为：
> **从可运行经营骨架走向可闭环、可消费的 Operating System。**

它的正式主轴收敛为：

1. **dry-run fix 与收口前置**（前提，不是可选项）
2. **Asset-Action Closure**（中心主线）
3. **Agent Consumption Layer**（支撑能力，受控最小版）
4. **Cross-Project Convention Asset**（候选探索，严格最小版或后移）

## 4.3 对“多条并列大主线”的仲裁

**最终结论：**

> 我不采纳把 `mvp0.4` 同时定义成“科学实验 workflow + 真实连接 + 完整 agent-first 重构”的并列结构；
> 我采纳“一个中心主线 + 一个受控支撑层 + 一个候选探索”的收敛结构。

原因如下：

1. 单人长期维护与当前项目约束，不支持多条重主线同时并发推进；
2. 本轮 `dry-run` 暴露出来的第一问题仍然是“动作链未闭合”，而不是“缺少更多新对象”；
3. agent 方向一旦做重，极易滑向 AI 工作台或自动化中枢；
4. 跨项目规范一旦做重，极易滑向知识图谱或文档平台；
5. 当前最需要的是把已有资产变成可持续消费的经营动作底座，而不是重新铺一层更大的叙事。

## 4.4 对“真实连接”路线的正式仲裁

`DPv4pro` 与部分汇总文档提出的“真实连接 / 代码触及 / GitHub API 读时连接”方向，**我部分采纳，但不将其冻结为独立主线。**

我的正式结论是：

1. **真实连接是方法，不是当前版本的中心命题。**
2. 只要某种连接能力直接服务于 `Asset-Action Closure` 或 `Agent Consumption Layer`，就可以进入候选实现范围；
3. 但它不能演化为“为连接而连接”的单独主题，更不能越界变成 GitHub 集成优先、扫描优先或工程工具化优先。

因此，本文件只允许以下最小版连接思路：

1. 读时补充已有 `Repository / Module / Product` 的最小上下文；
2. 为 agent 上下文导出提供更真实的只读素材；
3. 作为既有实体详情页的辅助信息，而不是第二事实源。

同时明确不冻结以下内容：

1. GitHub API 是否成为当前阶段必须项；
2. 是否增加任何新的连接型数据库表；
3. 是否做本地扫描、CI/CD 联动、自动同步、多平台仓库适配。

## 4.5 `Asset-Action Closure` 的正式定义

**最终结论：**

> `Asset-Action Closure` 是 `mvp0.4` 唯一中心主线。

它的正式内涵冻结为五件事：

1. **Onboarding 逻辑化**：从“首轮登记”升级为“首轮建链引导”，在创建过程中主动建立合理关系，而不是事后再让用户手动补链；
2. **Decision 状态机细化**：让 `proposed` 退出“双重语义”，至少能区分“已留痕 / 已确认 / 已生效”；
3. **Cross-Entity CTA 矩阵**：每个关键实体详情页都明确承接“下一步应该做什么”；
4. **Review Loop 业务闭环**：`review -> decision -> entity update -> review completed` 必须形成真实业务回流；
5. **Current Focus / pending signal 重做**：只指向真正待处理的动作，而不是泛化为“存在过决策记录”。

## 4.6 `Agent Consumption Layer` 的正式定义

在“Agent-Native Foundation”与“Agent Consumption Layer”两种说法中，本文**正式采用 `Agent Consumption Layer`**。

原因是：

1. 这个名称更准确强调“当前阶段先做可消费”，而不是暗示“要做完整 agent 平台”；
2. 它更符合项目规则中“AI 是增强层，不是主线控制层”的边界；
3. 它更容易把范围限制在只读上下文输出，而不是滑向第二套写入主线。

其正式范围冻结为：

1. **后端优先的 agent 能力暴露面**：agent 的正式接入应优先通过 Go 后端提供的 `MCP / CLI / API` 一类能力面进入，而不是把 React 前端当成 agent 的主工作台；
2. **标准化只读上下文接口**：让外部 agent 能按 `Product / Repository / Module / Decision` 获取一份可消费上下文；
3. **AGENTS 风格上下文导出**：把当前项目相关资产、决策与约束导出为稳定文本或结构化结果；
4. **最小上下文快照能力**：为 handoff 与外部消费提供受控快照，但不演化为模板平台或知识库。

其合同级边界进一步冻结为：

1. **后端 canonical contracts 不变**：无论是 Web、CLI 还是 MCP，PSCO 的正式领域语义与核心写入规则都仍以后端正式合同为准；
2. **agent 通道先消费、后扩写**：当前阶段只冻结 agent 的稳定只读消费，不冻结 agent 写入；
3. **不新增 agent 专属领域模型**：prompt、session、chat transcript 这类对象不作为当前阶段 PSCO 的一级业务对象；
4. **适配器不能长出第二套语义**：MCP / CLI 只是消费或适配通道，不得自行长出一套与后端合同并列的字段语义或状态机；
5. **若未来允许 agent 写入**：也只能回到同一套后端 canonical contracts，而不是新增隐藏写路径。

其明确非目标为：

1. AI 工作台；
2. agent 编排器；
3. 在 React 前端中新增对话框式主入口来让用户“通过聊天维护 PSCO”；
4. agent 自动写入 PSCO；
5. 自动代码执行、自动决策、自动建实体；
6. 第二套写入主线。

## 4.7 `Cross-Project Convention Asset` 的正式定位

**最终结论：**

> 这是值得正式留档的长期方向，但在 `mvp0.4` 只保留为候选探索，不升格为中心主线。

它的正确进入方式只能是：

1. 只登记工程可消费的规范、约束、技术栈基线；
2. 只承接“跨项目复用中真正稳定”的资产；
3. 只服务于 agent 上下文消费或新项目启动基线；
4. 不演化为知识图谱、笔记系统、语义搜索平台或跨项目大而全资产中心。

## 4.8 对长期经营对象回归的正式仲裁

`Gemini` 与 `Qwen` 对“科学实验 workflow”“Opportunity / Experiment / Lean Business Loop”的提醒，**方向上我认可，时机上我不采纳。**

正式结论如下：

1. 用户反思中表达的“探索 - 验证 - 迭代”诉求是真实的；
2. 原始方案文档中的长期经营语义并没有失效；
3. 但在当前轮次，把这些对象重新升格为正式主线，会再次回到“概念扩张先于使用闭环”的旧问题；
4. 因此它们应被明确记录为 `mvp0.5+` 的重点候选方向，而不是压入 `mvp0.4`。

## 4.9 当前明确后移的方向

以下方向在本文件中明确**后移到 `mvp0.5+` 候选范围**：

1. `Venture` 正式主线化；
2. `Opportunity / Feature / Experiment` 流程化；
3. `Decision Intelligence`（相似匹配、引用推荐、AI 上下文增强）；
4. `AI Context Enhancement`；
5. `Capability` 重实体 CRUD；
6. GitHub OAuth / 自动导入；
7. 自动扫描 / 知识图谱 / 语义搜索；
8. AI 一级工作台 / agent 编排器 / 自动代码执行；
9. 完整模板平台 / 模板版本管理；
10. Agent 写入 PSCO / Agent 双向同步；
11. 战略定位重新校准作为正式主线；
12. 完整科学实验 workflow。

## 4.10 当前不允许提前冻结的内容

以下内容在本文件中明确**不冻结**，留给后续正式 `/plan` 与 `/spec`：

1. 下一阶段正式 `phase` 名称；
2. `.trae/specs/phaseXX_*` 路径；
3. `Decision` 状态机的最终状态数量与最终命名；
4. `Agent Consumption Layer` 的最终接口名、输出格式与物理落点；
5. “真实连接”是否在某个子任务中被最小引入；
6. `Cross-Project Convention Asset` 是否进入 `mvp0.4`，还是整体后移到 `mvp0.5+`。

---

## 5. `mvp0.4` 正式范围基线

## 5.1 前置必做范围

在正式进入 `mvp0.4` 业务主线之前，必须先完成以下工作：

1. 通过 `fix*` workflow 闭合三个阻断项；
2. 重跑一轮聚焦型 `dry-run`，只验证阻断项是否真实消失；
3. 形成 `mvp0.3` 的正式收口记录，把已知非阻断摩擦点留档；
4. 在此基础上再开启下一阶段正式 `/plan`。

其最小收口门槛进一步冻结为：

1. **`fix_001` 完成标准**：用户从冷启动欢迎页进入时，“开始首轮录入”首次点击必须稳定进入下一步，不再出现首击无响应；
2. **`fix_002` 完成标准**：已被消费或已被正式确认的 `Decision`，不得继续在 `Dashboard / Daily Review / Current Focus` 中被误报为待处理；
3. **`fix_003` 完成标准**：`Decision Detail` 必须提供正式状态推进入口，且该推进结果能回流到 review 语义与详情展示；
4. **聚焦型 rerun 标准**：重跑时至少验证冷启动 onboarding、Decision 状态推进、Current Focus 语义、Review 回流这四类场景；
5. **`mvp0.3` 收口标准**：只有在上述阻断项关闭且 rerun 不再出现同类阻断后，才允许把 `mvp0.3` 标记为正式收口。

## 5.2 必做范围

下一阶段正式 `/plan` 必须覆盖以下两组正式范围：

1. **Asset-Action Closure**
   - Onboarding 逻辑化
   - `Decision` 最小生命周期闭合
   - Cross-Entity CTA 矩阵
   - Review 到实体动作的业务回流
   - Current Focus 与 pending signals 的语义重做

2. **Agent Consumption Layer（受控最小版）**
   - 后端优先的 `MCP / CLI / API` 暴露面
   - 标准化只读上下文接口
   - AGENTS 风格上下文导出
   - 最小上下文快照能力

## 5.3 候选范围

以下内容可作为 `mvp0.4` 候选范围，但不构成主线阻断项：

1. `Cross-Project Convention Asset` 严格最小版；
2. 直接服务于详情页上下文或 agent 导出的最小“真实连接”读增强；
3. 对 `Dashboard / Review / Product` 的轻量回流优化，只要它们不引入第二事实源。

## 5.4 明确不做范围

1. 新增长期核心实体主线；
2. 把 Onboarding 做成通用工作流引擎；
3. 把 `Decision` 细化做成完整 Decision Intelligence；
4. 把 agent 能力做成 AI 工作台、自动化控制台或第二套操作面；
5. 把跨项目规范做成知识图谱或通用文档平台；
6. 把“真实连接”做成重型 GitHub 集成、扫描系统或工程平台。

---

## 6. 对后续 `/plan` 的正式输入

## 6.1 推荐的完整候选阶段结构

我建议后续正式 `/plan` 至少按以下结构评估，但**不在本文冻结正式 phase 命名**。

### 候选阶段一：dry-run 阻断项 fix 与收口

目标：

1. 闭合 `Onboarding` 冷启动断裂；
2. 闭合 `Decision` 状态语义错位；
3. 闭合 `Decision Detail` 状态推进与 review 回流断裂；
4. 聚焦重跑 `dry-run` 并完成 `mvp0.3` 正式收口。

### 候选阶段二：Asset-Action Closure 主线

目标：

1. 让 `Onboarding` 变成真正的首轮建链引导；
2. 让 `Decision` 形成最小但真实的生命周期；
3. 让 Dashboard / Review / Detail pages 共同承接“下一步动作”；
4. 让 `Current Focus` 与 pending signals 回到真实经营语义。

### 候选阶段三：Agent Consumption Layer 支撑能力

目标：

1. 为外部 agent 提供稳定只读上下文入口；
2. 优先形成后端侧 `MCP / CLI / API` 暴露面，而不是前端对话式入口；
3. 让 PSCO 已有资产真正可被机器消费；
4. 为未来更深的工程协同与 agent-first 工作方式建立受控底座。

### 候选阶段四（可选）：Cross-Project Convention Asset 候选探索

目标：

1. 只登记稳定且工程可消费的跨项目规范；
2. 验证它是否真的提升新项目启动效率与 agent 消费效率；
3. 如无法保持极小边界，则整体后移到 `mvp0.5+`。

## 6.2 推荐的先后关系

1. 没有候选阶段一，`mvp0.3` 无法正式收口；
2. 没有候选阶段二，PSCO 仍停留在“实体登记完备但动作链断裂”的状态；
3. 没有候选阶段三，用户反思中的 agent-first 方向无法以受控方式落地；
4. 候选阶段四是可选项，不阻断 `mvp0.4` 主线收口。

## 6.3 对正式 `/spec` 的硬约束

1. `.proto` 仍是唯一长期合同源；
2. query 层继续保持纯只读；
3. 前端写路径继续收敛到切片固定承接位；
4. 不引入第二套路由、第二套状态管理、第二套 UI 事实源；
5. Onboarding 逻辑化不能演化为可配置工作流引擎；
6. `Agent Consumption Layer` 必须以后端暴露能力为正式主线，不得误落为 React 前端对话框入口；
7. `Agent Consumption Layer` 不得引入 agent 专属一级业务对象（如 chat/session/transcript）作为当前阶段主线；
8. `Agent Consumption Layer` 不能演化为 AI 工作台；
9. `Cross-Project Convention Asset` 不能演化为知识图谱；
10. `Decision` 状态机细化不能演化为 `Decision Intelligence`；
11. Review loop 不能演化为通用任务管理器；
12. 任何“真实连接”增强都不得形成与 PSCO 数据库并列的第二 canonical 数据源。

## 6.4 对后续验收的硬要求

1. 用户从空态进入，能否顺畅完成首轮建链，而不是分别创建再手动补关系？
2. `Decision Detail` 能否完成最小状态推进，并让 Dashboard / Review 不再误报？
3. 每个关键实体详情页是否都明确提示“下一步应该做什么”？
4. `Review -> Decision -> 实体更新 -> Review completed` 是否形成真实业务回流？
5. 外部 agent 能否通过单一稳定入口读取当前项目的完整核心上下文？
6. agent 能力是否以 `MCP / CLI / API` 这类后端优先通道暴露，而不是以 React 前端对话框形式落地？
7. 重跑真实项目后，用户是否愿意继续围绕 PSCO 持续回来使用，且摩擦显著低于本轮 `dry-run`？

---

## 7. 最终结论

**一句话总结：**

> `mvp0.4` 不是继续扩世界观，而是先把 PSCO 已经做对的骨架闭合成真实动作链，并在这个基础上建立最小的 agent 可消费层。

最终冻结为：

1. `mvp0.3` 先走 `fix + rerun dry-run` 正式收口；
2. `mvp0.4` 的唯一中心主线是 **Asset-Action Closure**；
3. `Agent Consumption Layer` 作为正式支撑能力进入，但只做受控最小版；
4. `Cross-Project Convention Asset` 只保留为候选探索，不直接升格为主线；
5. `真实连接` 可以作为服务于主线的局部方法进入，但不单列成当前版本中心命题；
6. `Venture / Opportunity / Experiment / Decision Intelligence / AI Context Enhancement / 自动扫描 / AI 工作台 / Agent 写入` 继续后移到 `mvp0.5+` 候选范围。

我对本轮两轮专家意见的最终明确是：

> **PSCO 当前不是做错了，而是还没把“从资产到动作”的最后一段真正闭合。**
> **先把这段闭上，再谈更大的 agent-first、跨项目资产化与经营对象回归，时机才是对的。**
