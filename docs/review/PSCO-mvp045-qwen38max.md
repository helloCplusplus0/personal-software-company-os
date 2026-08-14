# Personal Software Company OS

# MVP0.45 Next-Stage Direction Review (Qwen3.8-Max)

**Author:** Qwen3.8-Max  
**Date:** 2026-08-14  
**Role:** 基于本轮对仓库源码的逐项核实、`phase01 ~ phase10` 完整交付链与四份已存在的 `mvp0.45` 评审稿（GPT54 / DPv4pro / gemini31pro / GLM52）的交叉阅读，以独立专家身份对 PSCO 下一阶段推进方向给出评审意见与可仲裁主张  
**Document Type:** `review`  
**Status:** 供后续多位专家交叉评价、汇总仲裁与正式 `/plan` 参考；不直接冻结正式 `phase` 名称、`.trae/specs/phaseXX_*` 路径、接口名或实现细节

---

## 1. 文档定位与我的评审依据

本文是对 GPT54 [PSCO-mvp045-GPT54.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp045-GPT54.md) 的独立回应，同时与以下三份评审稿形成交叉：

- [PSCO-mvp045-DPv4pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp045-DPv4pro.md)（补充第五问：最小可验证交付物）
- [PSCO-mvp045-gemini31pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp045-gemini31pro.md)（补充审批门禁与主动注入）
- [PSCO-mvp045-GLM52.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp045-GLM52.md)（补充第六问：静态文件双源风险治理）

本文与四份已有评审的最大差异在于：**我不再以理论推演为主要依据，而是以本轮对仓库的逐项核实为唯一判断基础**，并把评审目标从"继续讨论方向"推进到"产出可仲裁、可执行的最小结论"。

### 1.1 本轮实际核实到的工程事实

以下每一条都是我在本轮会话中直接读取源码或目录结构确认过的，不是转述：

1. **后端**：`backend/internal/` 共 11 个业务模块（`backup / dashboard / decisioncenter / export / moduleregistry / onboarding / productregistry / repositorybinding / reusesummary / review / templatereuse`）加 `platform` 装配层；[router.go](file:///home/dell/Projects/personal-software-company-os/backend/internal/platform/router.go) 将 11 个 ConnectRPC handler 挂载到单一 `/api` 基址，`chi` 只保留 shell、middleware 与 `healthz` 等非业务端点，**无任何 legacy JSON HTTP 业务入口残留**。
2. **Proto 合同**：`proto/psco/*/v1/` 共 12 个 v1 合同（含 `common`）。**当前不存在任何"项目上下文聚合"类合同**——这正是各评审稿共同指向的新增点。
3. **前端**：10 个 feature 切片、19 个路由文件；技术栈（React 19 + TanStack Router/Query + Zustand + Tailwind 4 + shadcn + ConnectRPC web client）与 `TECH_STACK_BASELINE.md` 单值一致。`phase10` 交付物已实际落地：三类 detail page 均有 `*-next-action-bar.tsx` 与 `*-decision-entry-panel.tsx`。
4. **数据库**：9 个 migration（最新 `0009_phase10_onboarding_recovery_store.sql`）+ 完整 fixture/seed 体系 + `init_db.sh / run_seeds.sh / reset_*.sh` 受控入口。
5. **口径漂移实锤**：`PSCO-summarize-feedback.md` 在根目录**实际不存在**，但被 [AGENTS.md](file:///home/dell/Projects/personal-software-company-os/AGENTS.md)、[architecture_map.md](file:///home/dell/Projects/personal-software-company-os/architecture_map.md)、[docs/README.md](file:///home/dell/Projects/personal-software-company-os/docs/README.md) 多处引用为"最终共识"。当前实际最终共识由 [PSCO-mvp04-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp04-summarize-feedback.md) 承担。
6. **根级文档双源重复实锤**：`AGENTS.md §4` 与 `plan.md §2` 的 phase 状态清单几乎逐条重复；`architecture_map.md §4.1` 与 `docs/README.md §3` 的入口清单高度重复。
7. **规则早已存在但未被执行**：[project_rules.md](file:///home/dell/Projects/personal-software-company-os/project_rules.md) §1 早已写明单一真相源规则（phase 路线只由 `plan.md` 承担、目录结构只由 `architecture_map.md` 承担、不允许同一主结论被多个根级文档重复承载）。上述重复不是规则缺失，而是规则执行漂移。
8. **我自身的直接经验**：作为本轮接手的非 GPT-5.4 模型，我为恢复正确上下文实际读取了约 15 份文档与多个目录结构，并在过程中真实撞见了第 5 条的悬空引用。这不是思想实验，是刚发生的成本。

### 1.2 我的评审原则

1. **方向讨论的边际收益已经递减**：四份评审稿之后，原则层共识已经形成；继续增加方向讨论轮次，其本身正在成为新的漂移源。本文的职责是推动"从讨论走向仲裁"。
2. **只以核实过的工程事实为判断依据**，不重复任何未经源码确认的断言。
3. **不越权冻结**：正式 phase 命名、spec 路径、接口名留给 `/plan` 与 `/spec`。

---

## 2. 对 GPT54 "战略刹车"的总体判断：支持刹车，但风险焦点已经转移

我完全支持 GPT54 发起的战略刹车。在 `mvp0.4` 已把 `Agent Consumption Layer` 列为候选阶段三之后，如果直接跳入 MCP / CLI / API 的接口设计，确实会把未收敛的实体语义与错误的系统站位固化进合同层。这个刹车是必要且及时的。

但我必须指出一个四份评审都未明说的判断：

> **刹车踩到现在，"方向"实际上已经收敛了。** 四份评审稿在以下四点上没有实质分歧：PSCO 是上下文系统而非流程控制器；agent 先消费后维护；web 与 agent 共享 Go backend canonical core；不先建 MCP/CLI 协议层。

因此当前最大的风险已经不再是"方向跑偏"，而是**"讨论漂移"**——即在原则已无分歧的情况下，继续以"还需收敛"为由推迟一切可执行结论，让评审轮次本身变成目的。GPT54 文档中"不宜过早冻结"的清单如果被绝对化执行，其逻辑终点是任何实现都不能开始，因为任何实现都会冻结某些细节。

所以本文的核心任务不是继续论证方向，而是回答：

> **在原则共识已经形成的前提下，下一阶段可以被立即仲裁并执行的最小结论到底是什么？**

---

## 3. 对 GPT54 四个重点命题的逐项评价（基于核实证据）

### 3.1 命题一：四实体关系如何澄清 —— 降级为文档确认，不占用 phase 工程容量

GPT54 主张重述四实体语义。这个主张在概念层有洞察力，但我的核实结论是：**四实体的目标语义已经被现有实现满足，不存在需要下一阶段解决的结构性缺口。**

逐条证据：

- **`Product` = 经营目标与交付容器**：`product_registry.proto` 与后端 `productregistry` 承接的是 `name / description / value_proposition / stage / status` 等经营语义字段，不含任何代码结构信息。GPT54 的口径就是现状。
- **`Repository` = 代码仓库身份对象**：`Repository` 由显式 `CreateRepository` 产生，不存在隐式登记（这也是项目记忆中的硬约束），通过 `repository_binding` 与 Product/Module 建立关系。"不是所有本地目录的长期登记容器"这一条当前设计天然满足。
- **`Module` = 可复用能力资产（后置提炼）**：`Module` 已有 `prototype / candidate / stable` 生命周期状态；`phase09` 的模板候选正是"从已持久化的 `product_modules` 绑定事实中读时派生"——这本身就是"后置提炼"的工程实现。前置创建 prototype 与后置升级 stable 不互斥。
- **`Decision` = 规则/约束/选择的上下文层**：`Decision` 已有生命周期状态机，`decision_links` 表已建立与 Product/Module/Repository 的多对多约束关系，`phase10` 已完成状态推进闭环。"上下文层"语义已在结构中生效。

**我的结论**：命题一应从"下一阶段必须回答的重点工程问题"**降级为一次写入 `shared_baseline` 的语义确认文本工作**——用一页篇幅把 GPT54 的语义口径显式写清，确认"现状即目标语义"，不调整任何 `.proto`、schema 或代码结构。它值得做（因为能终止未来的反复讨论），但它不构成 phase 级工程交付，不应占用实现容量。

这一点我与 DPv4pro、GLM52 的"语义叠加而非结构重定义"立场一致，但我更进一步：**连"局部增强"都不需要立项，现状已满足，确认即收口。**

### 3.2 命题二：PSCO 是上下文系统而非流程控制器 —— 完全同意，且实现已部分落地

完全同意。这是四份评审的第一共识，也是下一阶段所有设计的最高原则。

我补充核实证据：`phase10` 落地的下一步动作承接矩阵（`*-next-action-bar.tsx` / `*-decision-entry-panel.tsx`）在工程上已经是"经营建议"语义——CTA 只指向 canonical path，不规定 IDE 现场开发顺序。也就是说这个定位不需要"重新确立"，只需要**显式声明并写进 `shared_baseline`**，作为 agent 渠道设计的约束条款。

### 3.3 命题三：agent 先消费后维护 —— 同意，并主张第一个消费者就是 PSCO 仓库自己

同意原则。GLM52 已论证消费的第一优先级是聚合端点而非新协议层，我认同。在此基础上我补充一个四份评审都没有显式提出的主张：

> **聚合上下文能力的第一个正式消费者，应该是 PSCO 仓库自己的开发工作流（dogfooding），而不是任何外部项目。**

理由有三：

1. **证据就在现场**：PSCO 本身就是由 agent 协作开发的。每一轮接手模型（包括本轮的我）都在支付"读取 15 份文档重建上下文"的成本。这个成本是真实的、可度量的、正在反复发生的。
2. **反馈回路最短**：以自己仓库为第一消费者，验证"聚合导出是否真的降低上下文成本"不需要任何外部仓库、外部 agent 工具链或真实业务数据，验收可以在本仓库内闭环完成。
3. **天然防漂移**：如果聚合导出连 PSCO 自己的上下文都不能准确表达，它就不可能正确服务外部项目。先对内兑现，再对外承诺。

这也为验收提供了一个极其具体的标准（见 §6）。

### 3.4 命题四：web 与 agent 非对称并行 —— 完全同意，且装配层零新增工作

完全同意。核实证据：`router.go` 的单一 `/api` 挂载方式意味着 agent 渠道**不需要任何新的后端装配工作**——ConnectRPC 天然支持 HTTP/JSON 传输，agent 现在就可以消费全部既有 read RPC。缺的只是"按项目聚合"的读取视角，不是传输通道。

与 DPv4pro、GLM52 一致的关键约束：agent 消费不得绕过 `.proto` 合同，字段语义必须单向派生。

---

## 4. 对现有三处分歧的仲裁立场

### 4.1 分歧一：静态文件治理 —— 既不是"全量派生"，也不是"一次性校准"，而是"单一写者规则执行"

- DPv4pro 主张静态规则文件从 backend canonical 数据派生；
- GLM52 反驳全量派生成本过高且引入新双源风险，主张 MVP0.5 只做去重与口径校准。

**我的仲裁：两者的框架都不够准确。正确的问题是"执行早已存在的单一真相源规则"。**

核实证据（§1.1 第 6、7 条）表明：`project_rules.md` §1 已经规定了每个主结论的唯一承载文档，当前的重复与漂移是**规则执行失败**，不是规则缺失。因此：

1. **反对 DPv4pro 的全量派生**：GLM52 的反驳成立——叙述性内容无法派生，且"backend 成为文档事实源"会引入循环依赖。但更重要的是，派生机制是在用新机制解决一个执行问题，属于过度工程。
2. **不满足于 GLM52 的一次性校准**：只做一轮校准，下一轮 phase 收口时同样的复制粘贴会再次漂移（过去 10 个 phase 已经反复证明了这一点）。
3. **我的方案——单一写者执行**：把根级文档中的"复制"改为"引用"：
   - `AGENTS.md §4` 只保留"当前阶段 + 当前主目标"两行总结，完整 phase 列表只存在于 `plan.md`；
   - `docs/README.md §3` 只保留 top 入口，完整落点只存在于 `architecture_map.md`；
   - `project_rules.md §2` 只保留冻结选择与禁止事项，技术栈正文只存在于 `TECH_STACK_BASELINE.md`。
4. **修复 `PSCO-summarize-feedback.md` 悬空引用**：我建议**补建该文件为薄指针文件**（内容只声明"当前最终共识 = 最新一轮 mvpXX-summarize-feedback"并给出链接），而不是修改三处引用。理由：`project_rules.md` 与 `AGENTS.md` 都把该文件名规定为最终共识的唯一入口；补建薄指针后，未来每一轮共识更替只需更新这一个指针文件，三处入口永久稳定。此口径涉及根级真相源结构，**需用户确认后执行**。

这项工作零代码、纯文档、可逆，是 agent 稳定消费的前置条件——如果 agent 消费的是互相重复且已漂移的静态文件，新建任何通道都只会放大错误。

### 4.2 分歧二：gemini31pro 的"主动注入" —— 拒绝主动注入，复用既有 Export 模式做受控导出

gemini31pro 主张 PSCO 在绑定/初始化时主动把全局规范注入目标仓库。

**我的仲裁：拒绝"主动注入"，采纳其动机，落为"受控导出"。**

1. 主动注入意味着 PSCO 写入外部仓库，这本身是写入行为，会带来文件冲突、版本控制与回滚复杂度；
2. 它与"PSCO 是上下文系统而非控制器"的最高原则存在语义张力——主动写入外部仓库容易被理解为"PSCO 控制目标仓库"；
3. 但 gemini31pro 的动机（让 agent 进场即有对齐基线）是真实的。核实证据：PSCO 已有 `ExportService / BackupService` 的成熟"只读快照 + 受控导出"模式。上下文导出应复用这一模式——**PSCO 提供可随时取用的 Markdown 上下文，是否复制、复制到哪、如何版本化，由用户或 agent 自主决定。**

一句话：提供弹药，不替 agent 扣扳机。

### 4.3 分歧三：gemini31pro 的"Draft & Approval 门禁" —— 支持 GLM52，MVP0.5 禁止设计任何 agent 写入路径

**我的仲裁：完全支持 GLM52。下一阶段禁止任何形式的 agent 写入路径设计，包括 Draft 提交接口。**

补充一条合同级依据：[PSCO-mvp04-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp04-summarize-feedback.md) §6.3 硬约束第 7 条明确"`Agent Consumption Layer` 不得引入 agent 专属一级业务对象作为当前阶段主线"。Draft 表、Draft 状态机、Draft 审批流正是 agent 专属领域模型的变体——一旦设计 Draft 接口，它们就会自然长出来。

若远期确需 agent 辅助写入，正确路径是复用 canonical 写合同（`CreateProduct / CreateModule / CreateDecision` 等）+ 人工在 web 端确认，而不是新建 agent 专属 Draft 系统。这是远期候选，不是当前范围。

---

## 5. 我的核心主张：下一阶段应是一个 phase，而不是两个阶段

DPv4pro 与 GLM52 都建议"方向收敛 phase + 最小交付 phase"两段式结构。**我主张合并为一个交付型 phase。** 这是本文与四份已有评审的最主要分歧。

### 5.1 为什么不应拆成两个 phase

1. **"方向收敛"的交付物已经小到不构成独立 phase**：经过本文 §3 的降级处理后，收敛侧产物只剩——四实体语义确认文本、上下文系统定位声明、先消费后维护原则冻结、web/agent 分工说明、静态文件单一写者执行。全部是文档级工作，零代码。
2. **这些文档工作恰好就是交付型 phase 内部的标准子任务类型**：`project_rules.md` §4.1 要求每个 phase 的 dev_plan 至少覆盖五类子任务，其中"边界收敛类子任务（回答做什么不做什么）"与"根级同步类子任务（回写状态与入口）"正是上述文档工作的正式承接位。把它们拆出去单独立 phase，反而违背了 phase 应"对应单一主交付能力"的定义。
3. **phase workflow 本身有重量**：本项目每个 phase 都要支付三件套 + `/spec` + 实现 + 验收 + 收口 + 根级同步的完整流程成本。对单人长期维护项目，为一个纯文档阶段支付全套流程开销是不划算的。两段式意味着支付两次。

### 5.2 该 phase 的主交付与子任务结构（供 `/plan` 参考，不冻结命名）

**主交付能力（单一）**：agent 可通过单一入口获取当前项目的完整核心上下文。

子任务按五类覆盖：

1. **边界收敛类**：四实体语义确认、上下文系统定位声明、先消费后维护原则、web/agent 分工边界、明确不做清单（见 §5.3）；
2. **实现设计类**：项目上下文聚合只读 RPC 的合同设计与读取聚合设计（按 `repository_id` / `product_id` 锚点，聚合 Product / Module / Repository / Decision / ReuseSummary，过滤 archived Decision，完全只读）；AGENTS 风格 Markdown 渲染设计；
3. **源代码实现类**：新增聚合只读合同（单一 proto 包）+ 后端读取 owner + Markdown 渲染端点；前端不新增长期页面（如需入口，复用 Dashboard 既有 Export 区域的紧凑条带模式）；
4. **验证验收类**：以 PSCO 自身仓库为第一消费者完成一次真实上下文重建 dogfood（§3.3），记录重建摩擦的前后对比；buf build / breaking / go build / frontend build 四层口径；
5. **根级同步类**：静态文件单一写者执行（§4.1）、`PSCO-summarize-feedback.md` 薄指针修复（需用户确认）、根级入口回写。

### 5.3 明确不做清单

1. MCP 协议层；
2. CLI 工具；
3. 任何 agent 写入路径（含 Draft 接口）；
4. 前端对话式 agent 入口；
5. 四实体结构重定义与 schema 扩张；
6. 主动注入外部仓库；
7. 静态文件全量 backend 派生（保留为远期候选）；
8. 知识图谱、自动扫描、重型 GitHub 集成。

### 5.4 我与 GPT54 的一个细微但重要的分歧：聚合端点的合同不属于"过早冻结"

GPT54 的不建议方向第 7 条是"在未收敛实体关系前就冻结具体 API / MCP / CLI 细节"。我支持对 MCP / CLI / 写入合同执行这条禁令，但**反对把"聚合只读端点的 proto"也划入此范围**。

理由：

1. 聚合端点是**既有 12 个稳定合同的只读投影**，其字段语义完全单向派生自现有 proto，不重新定义任何实体语义；
2. 实体语义在工程层已经收敛（§3.1 的核实结论），不存在"等待收敛"的前置条件；
3. 如果"实体关系未完全澄清"可以无限期推迟任何一个只读端点的设计，那么同样的逻辑可以推迟一切实现——这正是本文 §2 警告的讨论漂移。

真正危险的过早冻结是：在语义未稳时固化 agent 写入合同或 MCP 协议细节。只读聚合投影不在此列。**区分"投影型合同"与"语义型合同"，是避免刹车变成永久停车的关键。**

---

## 6. 验收标准：如何判断下一阶段成功

我建议以下五条作为验收门禁，全部可客观检查：

1. **上下文重建成本可度量下降**：以 PSCO 自身仓库为对象，新接手模型通过"聚合端点输出 + 一份 AGENTS 风格 Markdown + 根级入口"重建当前项目核心上下文，所需必读材料从本轮的约 15 份降至 3 份以内，且不丢失当前阶段、四实体关系、关键约束三类信息；
2. **根级文档无重复状态清单**：`AGENTS.md / docs/README.md / project_rules.md` 中不再存在与 `plan.md / architecture_map.md / TECH_STACK_BASELINE.md` 重复承载的主结论；
3. **悬空引用清零**：`PSCO-summarize-feedback.md` 引用不再指向不存在的文件；
4. **合同一致性**：新增聚合合同通过 buf build / breaking 检查，字段语义与既有合同单值一致，且为纯只读（无任何写 RPC）；
5. **零写入路径、零第二套语义**：验收时能明确列举"本阶段没有新增任何 agent 写入接口、没有新增任何与 `/api` 并列的第二 canonical 通道"。

---

## 7. 建议后续仲裁聚焦的五个问题

为避免讨论继续发散，我建议汇总仲裁只回答以下五个问题（前四个在四份评审中已接近收敛，第五个是本文新增）：

1. 是否同意四实体语义**降级为 `shared_baseline` 确认文本**，不作为 phase 级工程工作？
2. 是否同意静态文件治理采用**"单一写者规则执行 + 薄指针修复"**，而非全量派生或一次性校准？
3. 是否同意下一阶段采用**单一交付型 phase**（收敛子任务 + 聚合只读端点 + AGENTS 风格导出），而非两段式？
4. 是否同意**第一个正式消费者是 PSCO 仓库自身开发工作流**，并以此作为验收基准？
5. 是否同意区分"投影型合同"与"语义型合同"——**聚合只读端点可以立即进入设计，MCP / CLI / 写入合同继续冻结**？

---

## 8. 结论

我对 GPT54 四个主张与三份交叉评审的判断汇总如下：

| 议题 | 我的判断 | 关键依据 |
|------|------|------|
| GPT54 战略刹车 | 支持，但方向共识已形成，风险焦点应转向防止讨论漂移 | 四份评审在四点原则上无实质分歧 |
| 四实体关系重述 | 降级为文档确认，现状实现已满足目标语义 | 逐项源码核实（§3.1） |
| 上下文系统定位 | 完全同意，phase10 实现已部分落地 | CTA 组件已是经营建议语义 |
| 先消费后维护 | 同意，第一个消费者应是 PSCO 自身 | 本轮接手实际支付约 15 份文档的重建成本 |
| web/agent 并行 | 完全同意，装配层零新增工作 | router.go 单一 /api 挂载 |
| 静态文件治理 | 单一写者执行 + 薄指针修复 | 规则早已存在，问题是执行漂移 |
| 主动注入 / Draft 门禁 | 均拒绝进入当前范围 | 写入行为越界 + mvp04 §6.3 硬约束 |
| 阶段结构 | 单一交付型 phase，反对两段式 | 收敛产物过小，phase 流程有重量 |

如果用一句话总结我的主张：

> **方向已经收敛，现在需要的是仲裁而不是更多讨论：把四实体语义降级为一页确认文本，把静态文件治理收敛为单一写者执行，然后用一个交付型 phase 落地"聚合只读端点 + AGENTS 风格导出"，并以 PSCO 仓库自身的上下文重建成本下降作为验收标准——不建 MCP，不做 CLI，不开写入，让刹车在正确的地点松开。**
