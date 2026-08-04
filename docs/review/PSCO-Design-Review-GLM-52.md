# Personal Software Company OS

# 方案设计专业评价文档

**评审范围**：PSCO_0.md ~ PSCO_4.md（Document 00–08 全部）
**评审版本**：v1.0 Baseline
**评审模型**：GLM-5.2
**评审日期**：2026-08-04
**评审定位**：建设性同行评审（Peer Review），目标是识别设计风险、澄清模糊点、给出可落地的改进建议

---

# 0. 评审摘要（Executive Summary）

## 0.1 总体判断

PSCO OS v1.0 是一份**愿景清晰、逻辑自洽、范围克制**的体系设计。它完成了从"AI 代码资产记忆系统"到"个人软件公司运营基础设施"的范式跃迁，这一转型方向是正确的——它把工具问题上升为系统问题，把"记忆代码"上升为"积累能力"。

但作为一份将要进入工程阶段的 Baseline 设计，它在**领域模型的精度、技术架构的工程闭环、MVP 的可落地性**三个层面仍存在需要修正的盲区。

## 0.2 评分总览

| 评审维度 | 评分（10 分制） | 一句话结论 |
|---|---|---|
| 战略愿景清晰度 | 9 | 定位准确，飞轮模型有说服力 |
| 领域模型完备性 | 6 | 实体齐全但关系粗糙，关键边界模糊 |
| 模块系统设计 | 7 | 理念正确，但工程契约缺失 |
| 技术架构合理性 | 7 | 选型务实，但闭环未打通 |
| 工作流可执行性 | 6 | 理想化倾向，个人开发者流程过重风险 |
| AI 策略定位 | 8 | 边界清晰，但实现路径空缺 |
| UX 设计 | 6 | 理念对但原型粗糙，信息架构有断层 |
| MVP 范围控制 | 7 | 禁止清单优秀，但正向路径不清 |
| 文档体系一致性 | 7 | 概念串联好，但存在跨文档概念断层 |
| **综合** | **7.0** | **优秀的方向性设计，需补强工程化细节** |

---

# 1. 战略定位与价值主张

## 1.1 优点

### 1.1.1 范式判断准确

Document 00 对"AI 时代软件开发核心竞争力迁移"的判断——从编码能力迁移到**组织能力 + 判断能力 + 资产积累能力**——是本方案最坚实的地基。这一判断与 2025–2026 年行业实际演进方向高度一致：当代码生成成本趋近于零，**决策质量与资产复利**确实成为新的稀缺资源。

### 1.1.2 "不是什么"的边界划得好

Document 01 §5 与 Document 08 §17 通过"不是什么"反向定义产品边界（不是 Jira / 不是 Notion / 不是 GitHub / 不是 Cursor），这比正向描述更有效。特别是明确禁止"自动扫描一切""自动生成知识图谱""AI 自动判断最佳方案"，避免了 AI 时代产品最常见的**过度自动化陷阱**。

### 1.1.3 飞轮模型逻辑自洽

Build → Accumulate → Compound 三层递进，以及"更多模块 → 更强能力 → 更快开发 → 更多机会"的飞轮，在逻辑上是闭环的。

## 1.2 风险与问题

### 1.2.1 【高风险】飞轮的冷启动问题未讨论

飞轮模型有一个被文档回避的核心前提：**飞轮需要足够的"产品开发密度"才能自启动**。

个人开发者的实际产出节奏通常是：**一年 1–3 个项目**。在这种密度下：

- Module 积累速度慢（一年可能只产生 2–5 个真正可复用模块）；
- "复用 → 加速"的反馈周期长（可能 6–12 个月才能感受到第一次复用收益）；
- 飞轮在启动前就可能因"投入产出比不明确"而被放弃。

**建议**：在 Document 06 增补一节《飞轮冷启动策略》，明确：
- 前 3 个月的可量化里程碑（如：从 Rento-miniX 提取 N 个 Module）；
- "最小可感知复用"的定义（如：第二个项目复用第一个项目的 Auth Module）；
- 冷启动期的工作量预算（每周投入 PSCO 维护的时间上限，建议 ≤ 2 小时）。

### 1.2.2 【中风险】"个人软件公司"隐喻的双刃剑

"个人软件公司"是一个强力的心智锚点，但存在**过度拟物化**风险：

- 公司运营有"部门协同"的现实约束，个人开发者没有——强行套用 Opportunity/Venture/Product/Feature 四层结构，可能让一个本可 1 天做完的决策变成 1 周的"流程"；
- Document 02 §3 的 Innovation Loop（Idea → Opportunity → Hypothesis → Experiment → Decision）对个人开发者而言**环节过多**，实际工作中很多 Idea 应该可以直接跳到 Prototype 验证，而非走完整 5 步。

**建议**：明确区分**"重型路径"**与**"轻型路径"**：
- 重型路径：用于高投入决策（如是否启动一个新产品）；
- 轻型路径：用于低风险探索（如一个周末原型），允许从 Idea 直接到 Prototype，事后回补 Opportunity 记录。

### 1.2.3 【中风险】商业化路径与自用路径未区分

方案未明确 PSCO OS 是**纯粹的自用工具**还是**未来的商业产品**。这影响设计优先级：

- 若自用：可接受更高的个性化、更低的易用性门槛；
- 若商业化：需要考虑多租户、数据隔离、onboarding 流程等，这些当前完全缺失。

**建议**：在 Document 00 增补一节《产品形态定位》，明确 v1.0–v2.0 为自用阶段，v3.0+ 评估商业化。这会影响后续所有架构决策。

---

# 2. 领域模型设计（Document 03）

## 2.1 优点

实体识别完整，覆盖了从机会到能力的全链路。Venture 作为"长期战略容器"的抽象有价值——它把"一次性项目"和"长期方向"区分开，这是多数个人开发者缺失的视角。

## 2.2 问题

### 2.2.1 【阻断性】Capability 与 Module 的边界模糊

Document 03 §9 定义 Capability 为"多个模块和经验形成的长期个人能力"，但：

- 两者关系只在 ER 图中画了 `Module → Capability` 单向箭头，**没有说明聚合规则**；
- `related_modules` 是数组，但没有定义"一个 Module 可以属于多个 Capability 吗"；
- Capability 的 `experience_level` 取值未定义；
- Document 08 §5.2 的 Capability Growth 用了 Stable/Candidate/Experimental 三态，但 Document 03 没有定义这三个状态。

这是**整个领域模型最严重的概念断层**：Capability 是 PSCO 的最终价值产出（Document 00 §3.3 把 Compound 的落点放在 Capability），但它的定义却最薄弱。

**建议**：重写 Document 03 §9，至少明确：
- Capability 的形成条件（如：≥2 个 Module + ≥3 次实际使用 + 1 次跨项目验证）；
- Capability 与 Module 的多对多关系；
- Capability 的状态机（Dormant / Emerging / Established / Declining）；
- Capability 如何反哺新 Opportunity（闭合飞轮）。

### 2.2.2 【阻断性】Repository 实体定义过于简陋

Document 03 §12 的 Repository 只有 4 个字段（id/url/technology/owner），但实际工程中：

- 一个 Module 可能有**多个实现仓库**（如 Auth Module 有 Go 版和 TypeScript 版）；
- 一个 Product 仓库内部可能包含**多个 Module 的实现**；
- Repository 与 Module、Product 的关系是多对多，但文档未定义；
- 缺少分支策略、版本标签、提交钩子等与"资产沉淀"直接相关的属性。

**建议**：将 Repository 拆分为**逻辑实体**与**物理实体**：
- `Repository`（物理）：url、technology、branch_strategy；
- `ModuleImplementation`（逻辑）：module_id、repository_id、path、language、entry_point。
这样能支持"一个 Module 多语言实现"和"一个仓库多 Module"两种现实情况。

### 2.2.3 【中风险】Decision 实体缺少结构与关联

Document 03 §10 的 Decision 是扁平结构（context/problem/choice/reason/impact），但：

- **缺少关联对象**：一个 Decision 通常关联到具体的 Module / Product / Architecture 选择，但模型中没有 `related_entities` 字段；
- **缺少时间维度**：Decision 会被后续 Decision 覆盖（如 v1 选 PostgreSQL，v3 可能迁移），但模型没有 `supersedes` / `superseded_by` 字段；
- **缺少类型分类**：架构决策、技术选型、产品取舍、商业判断应区分，因为它们的复用价值不同。

**建议**：扩展 Decision 模型，至少增加 `type`（architecture/technology/product/business）、`related_entities[]`、`supersedes_id`、`status`（active/superseded/deprecated）。

### 2.2.4 【中风险】缺少"反馈闭环"实体

飞轮的关键是反馈，但领域模型中**没有承载用户反馈、运行指标、复盘结论的实体**。Experiment 只验证"假设"，不收集"上线后的真实数据"。

**建议**：新增 `Feedback` 实体（来自用户/市场）和 `Metric` 实体（来自产品运行），与 Product 关联。这是 Capability 从"主观判断"走向"经验证判断"的必要数据。

### 2.2.5 【低风险】Venture 与 Product 的层级在个人场景可能冗余

Document 02 §4 强调"一个 Venture 可产生多个 Product"，但对多数个人开发者，**Venture 与 Product 常常是 1:1 的**。强制区分可能增加无意义的分类负担。

**建议**：允许 Product 直接挂载到顶层（无 Venture），Venture 作为可选的"上层分组"。这与 Local First 原则一致——结构应服务于内容，而非内容填结构。

---

# 3. 模块系统设计（Document 04）

## 3.1 优点

### 3.1.1 Module 生命周期设计是亮点

Prototype → Candidate → Internal → Stable → Commercial 的五阶段模型，以及"不提前设计模块，从真实需求抽象"的原则，是本方案最成熟的部分。它正确地把 Module 的产生放在**实践验证之后**，避免了过度抽象。

### 3.1.2 Module 与 Package 的区分有价值

"Package 关注代码复用，Module 关注能力复用"的区分，为后续 Module Registry 的设计提供了清晰的定位。

## 3.2 问题

### 3.2.1 【阻断性】Module 接口契约未规范化

Document 04 §4.1 说"接口是资产边界"，但**没有定义接口如何描述**：

- 是 IDL（如 gRPC proto）？
- 是 OpenAPI Spec？
- 是自然语言 + 函数签名？
- 不同语言（Go / TypeScript / Rust）的接口如何统一表达？

这是 Module 能否真正"复用"的工程前提。没有接口契约规范，Module Registry 就只是一个带标签的代码链接列表，无法支撑 Document 08 §14 的"AI 推荐已有能力组合"场景。

**建议**：定义 PSCO Module Manifest 规范（PMM），至少包含：
```yaml
name: auth-module
version: 3.2.0
category: foundation
capability: identity-management
interface:
  spec_format: openapi  # 或 grpc-proto / graphql / custom
  spec_path: ./interface/auth.v3.yaml
  stability: stable
implementations:
  - language: go
    repository: github.com/xxx/auth-go
    entry: github.com/xxx/auth-go/v3
  - language: typescript
    repository: github.com/xxx/auth-ts
    entry: @xxx/auth
dependencies:
  - module: storage-module
    version: ">=2.0.0"
status: stable
used_by: [rento-miniX, ai-platform]
```

### 3.2.2 【阻断性】Module 依赖关系缺失

文档完全没有讨论 Module 之间的依赖。现实中：

- Auth Module 依赖 Storage Module（存 token）；
- Notification Module 依赖 Auth Module（鉴权）；
- 依赖关系是"能力组合"的基础，也是 Document 08 §14 "AI 推荐组合"的必要输入。

没有依赖图，"组合 Module 创建新产品"就是空中楼阁。

**建议**：在 Module 模型中增加 `dependencies[]`，并在 Module Registry 中维护依赖图，支持依赖冲突检测。

### 3.2.3 【中风险】Module 分类体系存在交叉

Foundation / Application / Domain / AI / Data 五类存在交叉：

- Storage 既是 Foundation（基础能力）也是 Data（数据能力）；
- User Management 既是 Application 也是 Domain（取决于上下文）；
- LLM Gateway 是 AI 还是 Foundation？

分类交叉本身不是大问题，但**没有定义分类规则**会导致注册时主观随意。

**建议**：采用**主分类 + 标签**模式：每个 Module 一个主分类（强制），多个标签（可选）。例如 Storage 主分类 Foundation，标签 [data, persistence]。

### 3.2.4 【中风险】Module 的废弃与迁移策略缺失

生命周期止于 Commercial，但**没有 Declining / Deprecated / Retired 状态**。Module 会过时（如从 REST 迁移到 gRPC），如何标记废弃、如何提示依赖项目迁移、如何处理历史 Release，文档未涉及。

**建议**：补充生命周期后段：Commercial → Declining → Deprecated → Retired，并定义迁移记录实体 `Migration`。

### 3.2.5 【低风险】"不按语言分类"与工程现实的张力

Document 04 §6 主张"不按语言分类，因为技术会变化"。理念正确，但工程上：

- 不同语言的包管理机制不同（go mod / npm / cargo）；
- 跨语言复用一个 Module 实际上意味着**重新实现**，而非复用；
- "能力复用"在跨语言场景下更多是**设计复用**而非**代码复用**。

**建议**：明确区分"能力复用"（设计 + 接口）与"实现复用"（代码）。Module Registry 同时管理两者，但标注复用层级。

---

# 4. 技术架构（Document 05）

## 4.1 优点

技术选型务实：React + Go + PostgreSQL 是个人开发者维护成本最低的成熟组合。MVP 范围明确排除 Rust Intelligence Layer 和 AI 自动化，范围控制得当。

## 4.2 问题

### 4.2.1 【阻断性】Rust Intelligence Layer 的引入缺乏判断标准

Document 05 §7 把 Rust 定位为"高性能能力"层（代码索引、语义搜索、依赖分析），但：

- 没有定义**何时引入**的判断标准（性能阈值？数据量阈值？）；
- 没有说明 Rust 与 Go 的**通信协议**（gRPC？HTTP？FFI？）；
- 没有评估**维护成本**：个人开发者维护双语言栈的负担是否可接受？

更关键的是：Document 06 §11 明确反对"自动扫描一切"，而 Rust 层的三个能力（代码索引、语义搜索、依赖分析）恰恰需要扫描代码。**这里存在文档间的逻辑矛盾**。

**建议**：在 v1.0 阶段**完全移除 Rust 层**，将相关需求推迟到 v2.0+ 并重新论证。当前架构应简化为 React + Go + PostgreSQL 三层，与 MVP 范围一致。

### 4.2.2 【阻断性】多语言 Module 的统一管理方案缺失

Module 涉及 Go / TypeScript / 可能的 Rust，但文档未讨论：

- Go Module 走 go mod，TS 走 npm，如何统一注册到 PSCO Registry？
- 是自建私有 registry，还是复用语言原生 registry + PSCO 元数据层？
- 项目如何"安装"一个 PSCO Module？（是 go get + 元数据记录，还是私有源？）

**建议**：采用**元数据层方案**——PSCO Registry 只管理 Module 的元信息（能力、接口、版本、使用记录），代码分发复用语言原生机制（go get / npm install）。PSCO 提供 CLI 工具协调两者。

### 4.2.3 【中风险】Local First 与数据同步的冲突未讨论

Document 05 §2 原则 3 是 Local First，但：

- 个人开发者有多台设备（开发机、笔记本、服务器）；
- Local First 如果意味着数据只在本地，如何多设备协同？
- 如果需要同步，是自建同步层还是用 Git / 云盘？

**建议**：明确 Local First 的含义是"数据所有权在本地"而非"数据只在本地"。推荐方案：**数据以文件形式存储（SQLite + JSON），通过 Git 仓库同步**，既满足 Local First 又解决多设备问题。

### 4.2.4 【中风险】缺少数据模型演进策略

PostgreSQL schema 会随版本演进，但文档未讨论：

- 是否使用 migration 工具（如 golang-migrate）？
- 破坏性变更如何处理？
- Module 的版本与 schema 版本如何关联？

**建议**：在 Document 05 增补《数据演进策略》，明确 migration 工具选型与版本兼容策略。

### 4.2.5 【低风险】后端目录结构过于扁平

Document 05 §5 的 `internal/venture/product/module/...` 是按实体分目录，但缺少：

- `domain/`（领域层，DDD 分层）；
- `service/`（应用层）；
- `repository/`（持久层）；
- `api/`（接口层）。

按实体分目录在小规模可行，但扩展性差。

**建议**：采用 DDD 分层 + 实体子包：
```
internal/
  domain/        # 领域模型与领域服务
    venture/
    product/
    module/
  application/   # 应用服务（用例编排）
  infrastructure/# 持久化、外部集成
  interfaces/    # API、CLI
```

---

# 5. 工作流设计（Document 06）

## 5.1 优点

Daily Workflow 与 Weekly Review 的节奏设计合理，"每次工作产生未来价值"的导向明确。

## 5.2 问题

### 5.2.1 【高风险】工作流的强制性可能成为负担

个人开发者最大的敌人是**流程摩擦**。Document 06 的 Daily Workflow 要求：

- Morning Review（查看 4 类信息）；
- During Development（记录决策）；
- End Session（更新 3 类状态）。

如果每项都是强制手动操作，实际坚持率会极低。文档没有讨论**自动化采集**与**手动记录的边界**。

**建议**：明确"自动采集 vs 手动记录"清单：
- 自动采集（通过 Git hook / IDE 插件）：commit 记录、文件变更、分支创建；
- 半自动（提示确认）：Module 提取候选、Release 发布；
- 强制手动：Decision、Opportunity、Experiment 结论。

### 5.2.2 【中风险】缺少"轻量入口"

Daily Workflow 假设用户主动打开 PSCO，但个人开发者最自然的动作是**在 IDE 里写代码**或**在浏览器里查资料**。PSCO 缺少"随手记录"的轻量入口：

- 在 IDE 里遇到一个决策点，如何 30 秒内记录到 PSCO？
- 在浏览器看到一个竞品，如何一键收藏为 Opportunity 候选？

**建议**：MVP 阶段提供至少一个轻量入口（CLI 命令或浏览器扩展），降低"进入 PSCO"的摩擦。

### 5.2.3 【中风险】Weekly Review 缺少量化指标

Weekly Review 的 4 个问题（新机会？用户价值？新能力？沉淀代码？）都是定性问题，没有可量化的指标。个人开发者很容易在"复盘"中自我感觉良好，但实际无进展。

**建议**：为每周Review定义最小指标集：
- 本周新增 Module 数 / 复用 Module 数；
- 本周记录 Decision 数；
- 本周 Feature 交付数；
- Capability 增长（经验值变化）。

### 5.2.4 【低风险】Phase 1–7 的命名与 Document 02 的 Loop 不一致

Document 02 用 Innovation/Product/Engineering/Capability Loop，Document 06 用 Phase 1–7（Opportunity/Validation/Venture/Product/Engineering/Module/Capability）。两套术语指代重叠但不一致，增加理解负担。

**建议**：统一术语体系，建议以 Document 06 的 Phase 为主，Document 02 的 Loop 作为 Phase 的分组。

---

# 6. AI 策略（Document 07）

## 6.1 优点

### 6.1.1 "AI 是增强层不是决策层"是正确判断

Document 07 §2 的核心原则与当前 AI 工程最佳实践一致。特别是在"AI 自动判断最佳方案"被明确禁止（Document 08 §17），避免了责任边界混乱。

### 6.1.2 反对自动代码扫描的论证有力

Document 07 §11 对"Scan Everything → LLM Extract → Knowledge Graph → Auto Recommendation"路径的批判，以及四点问题（噪声大、判断难、不理解价值、维护成本高），论证扎实。

## 6.2 问题

### 6.2.1 【高风险】Personal Context Layer 缺乏实现路径

Document 07 §10 提出 PSCO 提供 Product/Module/Decision/Experience Context 形成 Personal Context Layer，但：

- 这些 Context 如何**结构化输出**给外部 AI？（是 API？是文件？是 prompt 模板？）
- 如何与现有的 AI Coding 工具（Cursor、Claude Code、Copilot）**协同**？
- Context 的**体积控制**：个人开发积累几年后，Context 可能达到 GB 级，如何做相关性检索？

这是 Document 07 最大的空缺——提出了正确的目标，但没有给出实现路径。

**建议**：定义 PSCO Context Protocol（PCP），至少包含：
- Context 导出格式（Markdown + 结构化 YAML frontmatter）；
- Context 检索 API（按 Module / Product / Decision 维度查询）；
- 与 AI 工具的集成方式（如生成 `.cursorrules` / `CLAUDE.md` / `AGENTS.md`）。

### 6.2.2 【中风险】6 种 AI 能力未区分优先级与成本

Document 07 §4–9 列出 6 种 AI 能力（Research/Product/Engineering/Documentation/Module/Composition Assistant），但：

- 实现成本差异巨大（Documentation Assistant 是 LLM 调用，Module Assistant 需要代码理解能力）；
- 未区分 MVP 与未来版本；
- 未说明哪些是 PSCO 内建、哪些是依赖外部 AI 工具。

**建议**：制作 AI 能力优先级矩阵：
| 能力 | MVP | v1.0 | v2.0+ | 实现方式 |
|---|---|---|---|---|
| Documentation | ✓ | | | 内建 LLM 调用 |
| Engineering | | ✓ | | 集成外部工具 |
| Module | | | ✓ | 需要 PCP |
| Composition | | | ✓ | 需要 PCP + 依赖图 |

### 6.2.3 【低风险】"模型趋同"判断需要更新

Document 07 §12 说"模型趋同，竞争在 Personal Context"。这一判断在 2026 年需要修正——模型能力仍在快速分化（推理能力、长上下文、Agent 能力差异显著）。竞争不再是"纯 Context"，而是**Context + 判断力 + 工具协同**。

**建议**：弱化"模型趋同"论述，强调"Context 是个人不可复制的资产，模型是公共可获取的能力"。

---

# 7. UX 设计（Document 08）

## 7.1 优点

Command Center 定位准确，明确反对"Notion 页面集合 / Jira 任务列表 / GitHub 仓库浏览器 / AI Chat 界面"四种错误形态。MVP UX Scope 的 5 项功能（Dashboard / Module Registry / Product Registry / Decision Records / Project Binding）范围克制。

## 7.2 问题

### 7.2.1 【阻断性】信息架构中存在未定义实体

Document 08 §4 的导航包含 `Knowledge`，但 Document 03 领域模型中**没有 Knowledge 实体**。这是跨文档的概念断层——UX 引用了领域模型未定义的对象。

**建议**：二选一：
- 移除 `Knowledge` 导航项（Decision + Experiment 已覆盖大部分知识）；
- 或在 Document 03 补充 Knowledge 实体定义。

建议移除，避免引入新的模糊概念。

### 7.2.2 【中风险】Dashboard 原型过于简陋，缺少信息层级

Document 08 §3 的 Dashboard ASCII 原型是平铺式列表，没有：

- 信息优先级（什么最重要？）；
- 时间维度（今天 vs 本周 vs 长期）；
- 交互引导（点击进入哪里？）；
- 空状态设计（新用户看到什么？）。

**建议**：重新设计 Dashboard 信息层级：
- 第一屏：Current Focus（1 个 Venture + 1 个 Goal + 1 个 Pending Decision）；
- 第二屏：Recent Activity（最近 7 天的 Module/Decision/Release 变更）；
- 第三屏：Capability Growth（长期视角）。

### 7.2.3 【中风险】New Product Creation 被推迟可能错失早期验证

Document 08 §14 称 New Product Creation 是"未来最有价值场景"，但 MVP 不包含。问题在于：

- 这正是 PSCO 与"普通模块管理工具"的核心差异点；
- 不做这个，MVP 容易被感知为"另一个模块列表"；
- 早期用户反馈最需要的正是这个场景。

**建议**：将 New Product Creation 的**最小版本**（选择已有 Module → 生成 starter 模板）纳入 v0.2，作为 MVP 之后的第一优先级。

### 7.2.4 【中风险】Decision Center 与 Experiment Center 的输入成本高

这两个页面价值高，但输入成本也高（每个 Decision 要填 6 个字段，每个 Experiment 要填 5 个字段）。如果没有降低输入摩擦的设计，它们会变成"好看的空架子"。

**建议**：
- 提供 Decision / Experiment 的**快速模板**（常见决策类型预填字段）；
- 支持**渐进式填写**（先填标题，其他字段后补）；
- 允许**语音 / 自然语言输入**后由 AI 结构化（这是 AI 增强的合理用例）。

### 7.2.5 【低风险】缺少移动端 / 轻量访问方案

个人开发者的"记录时机"往往不在电脑前（如通勤时想到一个 Idea）。方案完全没讨论移动端，可能错失 Opportunity 捕获的最佳时机。

**建议**：MVP 不做移动端，但提供**移动友好的轻量入口**（如 PWA 或 Telegram Bot），至少支持 Opportunity 的快速记录。

---

# 8. MVP 范围与落地路径

## 8.1 优点

### 8.1.1 "禁止清单"是优秀的设计纪律

Document 08 §17 与 Document 05 §11 的禁止清单（不自动扫描、不自动生成知识图谱、不 AI 自动判断、不做 Notion 编辑器）是本方案设计纪律的体现。

## 8.2 问题

### 8.2.1 【阻断性】MVP 5 项功能的依赖关系未梳理

Document 08 §16 列出 5 项 MVP 功能，但未说明依赖关系：

- Dashboard 依赖 Module / Product / Decision 数据；
- Module Registry 是基础；
- Product Registry 依赖 Module（Product 由 Module 组成）；
- Decision Records 独立；
- Project Binding 依赖 Module + Product。

没有依赖梳理，开发顺序就不清晰。

**建议**：明确开发顺序：
1. Module Registry（基础）；
2. Decision Records（独立，可并行）；
3. Product Registry + Project Binding（依赖 Module）；
4. Dashboard（聚合展示）。

### 8.2.2 【中风险】冷启动数据策略缺失

MVP 上线后是空系统，没有 Module、没有 Product、没有 Decision。新用户（包括作者自己）打开看到空 Dashboard，体验极差。

**建议**：提供**冷启动脚本**，从 Rento-miniX 仓库自动提取首批 Module 候选（基于目录结构 + go.mod 依赖），人工确认后导入。这与 Document 05 §12 的"从 Rento-miniX 开始验证"一致。

### 8.2.3 【中风险】4–8 周时间盒可能乐观

5 项 MVP 功能 + React/Go/PostgreSQL 全栈 + 数据模型设计 + 冷启动，4–8 周对个人开发者（兼职开发）偏乐观。

**建议**：拆分为两阶段：
- v0.1（4 周）：Module Registry + Decision Records + CLI（无 UI）；
- v0.2（4 周）：Product Registry + Project Binding + Dashboard（完整 UI）。

### 8.2.4 【低风险】缺少验收标准

MVP 的 5 项功能没有定义"做到什么程度算完成"的验收标准。

**建议**：为每项功能定义 Done 标准，例如 Module Registry 的 Done 标准包括：CRUD API、版本管理、搜索、导入导出。

---

# 9. 跨文档一致性问题

## 9.1 术语不一致

| 概念 | Document 02 | Document 06 | 问题 |
|---|---|---|---|
| 循环 | Loop（4 个） | Phase（7 个） | 命名不一致 |
| 产品阶段 | Prototype/MVP/Growth/Stable/Retired | Idea/Prototype/MVP/Growth/Stable/Retirement | 起止状态不一致 |
| Module 状态 | Prototype/Candidate/Internal/Stable/Commercial | （未重述） | Document 08 §5.2 用了 Candidate/Experimental，与 04 不一致 |

**建议**：建立术语表（Glossary），统一所有状态机命名。

## 9.2 概念断层

| UX 引用概念 | 领域模型定义 | 状态 |
|---|---|---|
| Knowledge（导航） | 无 | 缺失 |
| Capability 三态（Stable/Candidate/Experimental） | Document 03 未定义 | 缺失 |
| Asset Evolution（Document 08 §5.3） | 无对应实体 | 缺失 |

**建议**：在进入工程阶段前，完成一次跨文档一致性审查，补齐所有概念断层。

---

# 10. 缺失的关键设计

以下是进入工程阶段前**必须补充**的设计：

## 10.1 数据备份与灾难恢复

个人资产 = 数据。文档完全未讨论备份策略。建议：Git 仓库 + 异地备份（如 rsync 到云存储）。

## 10.2 安全模型

即使是个人使用，也需要：
- 数据库访问认证；
- API 鉴权（若提供 Web UI）；
- 敏感信息（如仓库 token）的加密存储。

## 10.3 测试策略

文档未讨论测试。建议：领域层单元测试 + API 集成测试 + 关键流程 E2E。

## 10.4 部署方案

Local First 意味着自部署，但未说明：
- 部署在本地服务器还是云？
- 用 systemd / Docker / 其他？
- 如何更新版本？

## 10.5 可观测性

个人系统也需要日志、错误追踪、性能监控，否则问题难以定位。

---

# 11. 改进建议优先级排序

按"对落地的阻断程度"排序：

## P0（阻断性，工程化前必须解决）

1. **定义 Module 接口契约规范（PMM）**——没有它，Module Registry 无意义；
2. **补充 Module 依赖关系模型**——没有它，能力组合不可能；
3. **澄清 Capability 与 Module 的边界与关系**——它是飞轮的落点；
4. **扩展 Repository 模型支持多实现**——现实工程需要；
5. **移除 MVP 阶段的 Rust 层**——避免过度工程；
6. **补齐 UX 引用的未定义实体（Knowledge）**——消除概念断层。

## P1（高优先级，影响飞轮启动）

7. **补充飞轮冷启动策略**——决定产品能否熬过早期；
8. **明确自动采集 vs 手动记录边界**——决定工作流可坚持性；
9. **定义 PSCO Context Protocol（PCP）**——AI 增强的前提；
10. **梳理 MVP 功能依赖与开发顺序**——进入工程的前提；
11. **提供冷启动数据策略**——避免空系统体验。

## P2（中优先级，v0.2–v0.3 解决）

12. 重写 Dashboard 信息层级；
13. 补充 Module 废弃与迁移策略；
14. 补充 Decision 的关联与版本管理；
15. 明确 Local First 的同步方案；
16. 提供轻量入口（CLI / 浏览器扩展）；
17. 区分重型 / 轻型工作流路径。

## P3（低优先级，v1.0+ 解决）

18. 统一跨文档术语；
19. 补充安全 / 备份 / 部署 / 可观测性设计；
20. 评估移动端轻量入口。

---

# 12. 总体结论

## 12.1 方案的核心价值

PSCO OS v1.0 的核心价值不在于它定义了"个人软件公司"的概念，而在于它**正确地识别了 AI 时代个人开发者的真问题**：

> 不是"如何更快写代码"，而是"如何让每一次创造留下长期资产"。

这一判断如果成立，那么 PSCO 的设计方向就是正确的——它把个人开发从"一次性项目"模式升级为"能力复利"模式。

## 12.2 方案的核心风险

方案的最大风险不是技术实现，而是**飞轮启动前的坚持期**：

- 个人开发者的产出密度天然较低，飞轮转得慢；
- 系统的维护成本如果超过它带来的复用收益，就会被放弃；
- "个人软件公司"的隐喻如果被过度执行，会让轻量工作变重。

 mitigations 的核心是：**降低维护摩擦，提高首次复用感知速度**。

## 12.3 对下一阶段的具体建议

进入 Document 09（MVP Specification）前，建议先完成三件事：

1. **完成一次跨文档一致性审查**，补齐所有概念断层与术语不一致；
2. **定义 PMM（Module Manifest）与 PCP（Context Protocol）两个核心规范**，它们是工程化的地基；
3. **用 Rento-miniX 做一次 dry-run**：手动走一遍"从代码提取 Module → 注册 → 关联 Product → 记录 Decision"的完整流程，验证领域模型是否够用，再写 MVP Spec。

完成这三件事，PSCO OS 就具备了进入工程阶段的条件。

---

# 附录 A：评审检查清单

| 检查项 | 状态 | 说明 |
|---|---|---|
| 愿景清晰 | ✅ | Document 00 定位准确 |
| 边界明确 | ✅ | "不是什么"定义充分 |
| 领域模型完备 | ⚠️ | 实体齐但关系粗，Capability 薄弱 |
| 状态机定义 | ⚠️ | 部分实体缺状态机，部分状态跨文档不一致 |
| 接口契约规范 | ❌ | Module 接口未规范化 |
| 依赖关系模型 | ❌ | Module 依赖完全缺失 |
| 技术选型合理 | ✅ | React+Go+PG 务实 |
| 架构闭环 | ⚠️ | Rust 层与 MVP 矛盾，多语言管理缺失 |
| 工作流可执行 | ⚠️ | 流程偏重，自动/手动边界不清 |
| AI 边界清晰 | ✅ | 增强层定位准确 |
| AI 实现路径 | ❌ | PCP 未定义 |
| UX 信息架构 | ⚠️ | 有未定义实体，Dashboard 粗糙 |
| MVP 范围 | ✅ | 禁止清单优秀 |
| MVP 依赖梳理 | ❌ | 功能依赖未说明 |
| 冷启动策略 | ❌ | 数据初始化缺失 |
| 跨文档一致性 | ⚠️ | 术语与概念有断层 |
| 备份/安全/部署 | ❌ | 完全缺失 |
| 可观测性 | ❌ | 缺失 |

**结论**：18 项中 5 项通过、8 项需改进、5 项缺失。建议完成 P0 级改进后进入工程阶段。

---

**评审完成。**

本评价文档将作为 PSCO OS v1.0 设计基线的同行评审记录留档，供后续 Document 09（MVP Specification）设计参考。
