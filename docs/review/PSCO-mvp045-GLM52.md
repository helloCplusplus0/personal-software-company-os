# Personal Software Company OS

# MVP0.45 Next-Stage Direction Review (GLM-5.2)

**Author:** GLM-5.2
**Date:** 2026-08-14
**Role:** 作为 `PSCO-mvp04-GLM52.md` 上一轮评审的延续者,基于对 PSCO 仓库实际源码、`phase01 ~ phase10` 完整交付链与三份 `mvp0.45` 评审稿(GPT54 / DPv4pro / gemini31pro)的交叉阅读,以独立专家身份给出对 PSCO 下一阶段推进方向的评审意见
**Document Type:** `review`
**Status:** 供后续多位专家交叉评价、汇总仲裁与正式 `/plan` 参考;不直接冻结正式 `phase` 名称、`.trae/specs/phaseXX_*` 路径、接口名或实现细节

---

## 1. 文档定位

本文是对 GPT54 [PSCO-mvp045-GPT54.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp045-GPT54.md) 的独立回应,同时与 [PSCO-mvp045-DPv4pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp045-DPv4pro.md) 和 [PSCO-mvp045-gemini31pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp045-gemini31pro.md) 形成交叉。

本文不复述三份已存在评审中已充分论证的内容,而是基于我对仓库实际源码的全面梳理,做四件事:

1. 对 GPT54 的四个核心主张给出独立判断(同意 / 部分同意 / 不同意 + 理由),并对 DPv4pro / gemini31pro 的关键分歧点表明立场;
2. 补充三份评审未充分展开的第六个重点问题:**静态规则文件体系本身的双源风险治理**;
3. 提出一个更具体的、与当前 12 个 ConnectRPC 服务现状直接对齐的最小可验证交付物定义;
4. 为后续多位专家仲裁提供一份基于工程现实的输入端。

本文的判断标准只有三个:

1. 是否符合 PSCO 当前已冻结的共识、边界与规则(`project_rules.md` / `TECH_STACK_BASELINE.md`);
2. 是否真正回应了仓库当前的真实状态(12 个 ConnectRPC 服务、9 个前端切片、`phase10` 已收口),而不是基于理论推演;
3. 是否有助于 PSCO 从"高质量结构化登记系统"继续推进到"可被 agent 稳定消费的项目上下文系统"。

---

## 2. 评审依据

### 2.1 三份 `mvp0.45` 评审稿

- `docs/review/PSCO-mvp045-GPT54.md`(刹车发起人,提出四主张)
- `docs/review/PSCO-mvp045-DPv4pro.md`(逐项评价 GPT54,补充第五问)
- `docs/review/PSCO-mvp045-gemini31pro.md`(全盘接收 GPT54,补充审批门禁与主动注入)

### 2.2 当前最终共识与规则基线

- `PSCO-mvp04-summarize-feedback.md`(当前最终仲裁,候选阶段三 `Agent Consumption Layer` 原定方向)
- `AGENTS.md` / `plan.md` / `project_rules.md` / `TECH_STACK_BASELINE.md` / `architecture_map.md`

### 2.3 仓库实际源码与交付状态

- `phase10-11` 正式验收已完成,`phase10` 三件套保留为最近完成正式业务 phase 的规划与冻结记录
- 后端 12 个 ConnectRPC 服务模块(`backend/internal/`)
- 前端 9 个 feature 切片(`frontend/src/features/`)
- 12 个 `.proto` v1 合同(`proto/psco/*/v1/`)
- 9 个数据库 migration 与完整 fixture 体系(`database/`)

### 2.4 本文的仲裁原则

1. **工程现实优先于概念推演。** 已稳定运行 10 个 phase 的实现不应被"重定义"轻率推翻。
2. **agent 消费场景优先于 agent 协议设计。** 先理解 agent 真实消费什么,再决定是否需要新协议。
3. **静态文件治理优先于新通道建设。** agent 当前已在消费静态规则文件,如果这些文件本身有双源风险,建新通道只会放大错误。
4. **受控最小版优先于全面铺开。** 单人长期维护项目不支持多主线并发。
5. **不越权冻结正式 phase。** 正式阶段命名、spec 路径与具体落点留给后续 `/plan` 与 `/spec` 决定。

---

## 3. 对 GPT54 四个主张的逐项评价

### 3.1 主张一:重新定义四实体关系

**GPT54 的核心主张:**
- `Product` = 经营目标与交付容器,不是代码结构本身
- `Repository` = 代码仓库身份对象,不是所有本地目录的长期登记容器
- `Module` = 可复用能力资产,后置提炼而非前置规划
- `Decision` = 对实体生效的规则/约束/选择索引,上下文层而非并列业务对象

**我的评价:方向正确,但应"语义澄清 + 局部增强",不应"结构重定义"。我完全同意 DPv4pro 的"语义叠加"立场,并基于实际源码给出更具体的工程判断。**

理由如下:

#### 3.1.1 `Product` 的定位已被现有实现满足

GPT54 说"Product 不是代码结构本身"——这在当前实现中已经成立。

观察 `backend/internal/productregistry/types.go` 与 `proto/psco/product_registry/v1/product_registry.proto`:`Product` 承接的是 `name / description / value_proposition / stage / status` 等经营语义字段,通过 `product_repositories / product_modules` 表与 `Repository / Module` 建立关系,不承接任何代码结构信息。

GPT54 的语义澄清在概念层面有洞察,但在工程层面已被现有实现满足。不需要"重新定义",只需要在 `AGENTS.md` 或后续 `phase` 规格中做一次显式的语义说明,确认"`Product` 是经营目标容器"这一口径。

#### 3.1.2 `Repository` 的定位已被现有实现满足

GPT54 说"Repository 首先是代码仓库身份对象"——当前实现与此一致。

观察 `proto/psco/repository_binding/v1/repository_binding.proto`:`Repository` 承接 `name / provider / url / binding_scope`,通过显式 `CreateRepository` 产生,不存在隐式登记。`Repository` 与 `Product / Module` 的关系通过 `repository_binding` 表显式建立。

GPT54 强调"不是所有本地目录的长期登记容器"——这一点当前设计已经天然满足。不需要任何结构调整。

#### 3.1.3 `Module` 的定位是 GPT54 主张中与当前实现张力最大的部分,但"前置"与"后置"不互斥

GPT54 说"Module 更适合被定义为后置提炼资产"——这意味着 Module 应该从 Product 开发中"被识别"出来,而非前置创建。

但观察当前实现:
- `Module` 已有 `status` 字段,支持 `prototype / candidate / stable` 等生命周期状态
- `Onboarding` 流程中确实前置创建 Module,但这与"后置提炼"不冲突——前置创建的 Module 可以是 `prototype` 状态,随着跨项目复用逐步升级为 `stable`
- `phase09` 已完成模板候选与派生提示,本质就是"从已持久化的 `product_modules` 绑定事实中读时派生",这本身就是"后置提炼"的工程实现

**我的判断:GPT54 的"后置提炼"洞察应该被吸收为对 Module 生命周期的语义澄清,而不是触发实体关系重定义。** 当前实现已经支持"前置创建 prototype + 后置升级 stable"的演进路径,只需要在规格层显式说明这一口径即可。

#### 3.1.4 `Decision` 的"上下文层"定位已被 `decision_links` 表满足

GPT54 说"Decision 应逐步演进为规则/约束/选择的索引对象,更像上下文层而非并列业务对象"。

观察当前实现:
- `Decision` 已有四态生命周期:`proposed / active / superseded / archived`
- `decision_links` 表已建立 `Decision` 与 `Product / Module / Repository` 的多对多关系
- `phase10` 已完成 `Decision Detail` 状态推进 CTA 与 `Current Focus / pending signals` 反回归验证
- `Decision` 通过 `decision_links` 对实体施加约束语义,这本身就是"上下文层"的工程实现

**我的判断:GPT54 的"上下文层"定位可以理解为对现有 Decision 职责的语义升级,而不是结构重写。** 不需要把 `Decision` 从"并列业务对象"降级为"上下文层"——它当前既是业务对象(有 CRUD),又是上下文层(通过 `decision_links` 对实体施加约束)。这两种角色不互斥。

#### 3.1.5 结论

我同意 GPT54 对四实体语义的澄清方向,但主张采用"语义澄清 + 局部增强"而非"结构重定义"的策略。具体落地方式:

1. 在下一阶段 `/plan` 的 `shared_baseline` 中,显式写明四实体的正式语义口径(吸收 GPT54 的洞察)
2. 不调整任何现有 `.proto` 合同、数据库 schema 或前后端代码结构
3. 若需增强,只在现有结构上做局部字段补充(如 `Module` 的 `status` 字段语义细化),不触发实体关系重写

这个立场与 DPv4pro 完全一致,但我基于实际源码给出了更具体的工程判断依据。

### 3.2 主张二:PSCO 不应成为开发流程控制器

**GPT54 的核心主张:**
- PSCO 不负责规定 agent 下一步必须怎么开发
- IDE / agent 负责项目内的微观推进
- PSCO 负责提供上下文、关系、约束、决策、回写目标与全局视图
- PSCO 更适合作为"上下文系统"而非"流程编排器"

**我的评价:完全同意。这是 GPT54 文档中最关键、最正确的判断,也是三份评审(GPT54 / DPv4pro / gemini31pro)的第一共识。**

我补充一个基于实际源码的观察:

当前 `phase10` 已落地的"下一步动作承接矩阵"(`NextActionBar` / `DecisionEntryPanel` 等组件,分布在 `product-detail-page.tsx` / `module-detail-page.tsx` / `repository-binding-detail-page.tsx` 中)在工程实现上已经采用了"经营建议"而非"开发指令"的语义:

- CTA 只指向 canonical path(`Repository / Module / Decision` owner 或其正式 handoff)
- CTA 不规定 IDE 现场的开发顺序
- CTA 的优先级规则是"先闭当前 canonical 结构缺口,其次推进 Decision 状态,最后返回 Dashboard reread"

也就是说,GPT54 主张的"PSCO 是上下文系统"在 `phase10` 实现中已经部分落地。下一阶段需要做的不是"重新定位 PSCO",而是"把这个定位显式写进 `shared_baseline`,并在 agent 渠道设计中严格遵守"。

**结论:完全同意。PSCO 是上下文系统,不是流程控制器。这是下一阶段所有设计的最高原则。**

### 3.3 主张三:agent 渠道应先消费、后维护

**GPT54 的核心主张:**
- 先让 agent 稳定消费 PSCO 上下文
- 再让 agent 辅助维护 PSCO
- 自动写回必须后置,且必须受控

**我的评价:同意原则,但需补充一个三份评审都未充分展开的工程现实观察。**

#### 3.3.1 当前 12 个 ConnectRPC 服务已经是"agent 可消费的只读接口"

这是一个三份评审都没有显式提到的工程事实:

PSCO 当前的 12 个 ConnectRPC 服务中,有大量 read RPC 已经是 agent 可直接消费的只读接口:

| 服务 | agent 可消费的 read RPC |
|---|---|
| `ProductRegistryService` | `ListProducts / GetProductDetail / ListProductModuleCandidates` |
| `ModuleRegistryService` | `ListModules / GetModuleDetail / ListProductCandidates / ListRepositoryCandidates` |
| `RepositoryBindingService` | `ListRepositories / GetRepositoryDetail / ListRepositoryProductCandidates / ListRepositoryModuleCandidates` |
| `DecisionCenterService` | `ListDecisions / GetDecisionDetail / ListDecisionModuleCandidates` |
| `DashboardService` | `GetDashboardOverview / GetFeedbackSignals / GetRecentActivities` |
| `ReviewService` | `GetDailyReviewContext / GetWeeklyReviewContext` |
| `OnboardingService` | `GetFirstRunState / GetOnboardingChainState` |
| `TemplateReuseService` | `ListTemplateCandidates / GetTemplateCandidatePrefill / GetDerivedInsightHints / GetTemplateSourceSummary` |
| `ReuseSummaryService` | `GetReuseSummary` |
| `ExportService` | `GetExportSnapshot` |
| `BackupService` | `GetBackupSnapshot` |

也就是说,**"agent 先消费"的基础设施已经存在**。agent 完全可以通过 ConnectRPC 协议(或其 HTTP/JSON 传输形式)直接消费上述所有 read RPC。

#### 3.3.2 但当前 RPC 是面向前端切片设计的,不是面向 agent 消费场景设计的

问题在于:这些 RPC 的字段语义偏向 UI 展示,而不是 agent 在 IDE 现场的消费场景。

具体差异:

- **前端消费场景**:按页面切片查询,分页加载,字段面向 UI 组件渲染
- **agent 消费场景**:按当前项目(物理仓库)聚合查询,一次性获取完整上下文,字段面向决策与约束理解

例如,agent 在 IDE 现场进入一个仓库后,它需要的是:"这个仓库对应哪个 Product?这个 Product 有哪些 Module?这些 Module 有哪些 active Decision 约束?当前项目的关键技术约束是什么?"

这个问题当前需要 agent 依次调用 `ListRepositories → GetRepositoryDetail → ListProducts(filter by repository) → ListModules(filter by product) → ListDecisions(filter by product/module)` 五六个 RPC 才能拼出答案。这对 agent 来说成本过高。

#### 3.3.3 因此"先消费"的第一优先级是"聚合导出端点",不是"建新协议层"

这与 DPv4pro 的判断完全一致,但我可以基于上述工程观察给出更具体的理由:

> 当前 12 个 ConnectRPC 服务已经提供了完整的只读接口,但它们是面向前端切片设计的,不是面向 agent 消费场景设计的。agent 在 IDE 现场需要的不是"列表分页查询",而是"按当前项目聚合的完整上下文"。因此,"先消费"的第一优先级是新增一个"项目上下文聚合只读端点",而不是建新的 MCP/CLI 协议层。

#### 3.3.4 对 gemini31pro"审批门禁"的独立判断

gemini31pro 提出 MVP0.6+ 开放写入时必须引入 "Draft & Approval" 机制。我完全同意这个战术,但要补充一个约束:

> **在 MVP0.5 阶段,严禁任何形式的 agent 写入路径设计,包括"Draft 提交"接口。** 原因是:一旦设计了 Draft 接口,就会自然长出 Draft 表、Draft 状态机、Draft 审批流,这本身就是"agent 专属一级业务对象"的变体,违反 `PSCO-mvp04-summarize-feedback.md` §4.6 的合同级边界("不新增 agent 专属领域模型")。

正确做法是:MVP0.5 只做 read-only;如果未来需要 agent 写入,再回到同一套后端 canonical contracts(`CreateProduct / CreateModule / CreateDecision` 等),由 agent 调用这些已有 RPC 提交"待人工确认"的实体——但这需要在 canonical contracts 层增加"draft / pending_review"状态,而不是新建 agent 专属的 Draft 表。

**结论:同意"先消费、后维护"。但主张消费的第一优先级是"项目上下文聚合只读端点",不是建新协议层。同时反对在 MVP0.5 阶段设计任何形式的 agent 写入路径,包括 gemini31pro 的 Draft 提交接口。**

### 3.4 主张四:web 与 agent 可以并行,但不能对称并行

**GPT54 的核心主张:**
- web 不退化,继续作为 PSCO 的正式交互渠道
- agent 渠道以受控最小能力面进入
- 二者共用同一套 Go backend canonical core
- 不允许 web 和 agent 各自长出第二套语义与流程

**我的评价:完全同意。这是保证架构不崩塌的生命线,与当前实现高度吻合。**

我补充一个基于实际源码的观察:

当前 `backend/internal/platform/router.go` 的装配方式已经天然支持"web 与 agent 共享 backend core":

- 所有 12 个 ConnectRPC handler 通过 `http.StripPrefix("/api", handler)` 挂载到单一 `/api` 基址
- 4 个 legacy `chi + JSON HTTP` 入口已全部退场
- `chi` 只保留 `/api` shell、middleware 与 `healthz` 等非业务端点

这意味着:agent 渠道不需要任何新的后端装配工作。agent 可以通过 HTTP/JSON 形式直接消费 `/api` 下的所有 ConnectRPC 端点(ConnectRPC 协议天然支持 HTTP/JSON 传输)。

**我补充一个关键约束(与 DPv4pro 一致):agent 渠道的"只读消费"不能绕过 `.proto` 合同。** 无论 agent 通过 MCP、CLI 还是 HTTP 消费 PSCO 数据,其字段语义、枚举、响应 envelope 必须与 `.proto` 单值一致。这确保了 web 和 agent 共享同一个领域模型事实源。

**结论:完全同意。"web 不退化 + agent 受控最小进入 + 共享 backend core"是正确且可行的架构原则,且当前实现已经天然支持。**

---

## 4. 对 DPv4pro 与 gemini31pro 的独立仲裁

### 4.1 对 DPv4pro 第五问(最小可验证交付物)的判断

DPv4pro 提出三件事作为最小可验证交付物:

1. 静态规则文件 → backend canonical 数据派生
2. 项目上下文聚合导出只读端点(`ProjectContextRead`)
3. AGENTS 风格上下文导出(Markdown 输出)

**我的判断:完全同意第二、第三件事,但对第一件事主张更克制的版本。**

#### 4.1.1 同意第二件事:项目上下文聚合导出端点

这是 MVP0.5 的核心交付物。基于我对当前 12 个 ConnectRPC 服务的观察(§3.3.1 / §3.3.2),这个端点应该:

- 新增一个只读 ConnectRPC 服务(如 `ProjectContextService`),或在现有某个服务下新增 RPC
- 输入:`repository_id` 或 `product_id`(按 agent 在 IDE 现场的锚点查询)
- 输出:聚合的 `ProjectContext` 消息,包含:
  - 关联的 `Product` 信息(name / stage / value_proposition)
  - 关联的 `Module` 列表及状态
  - 关联的 `Repository` 信息(name / provider / url)
  - 关联的 `Decision` 列表及状态(过滤 `archived`)
  - 关联的复用摘要(`ReuseSummary`)
- 完全只读,不引入任何写入语义
- 输出格式与 `.proto` 合同单值一致

#### 4.1.2 同意第三件事:AGENTS 风格上下文导出

这是 `ProjectContextRead` 的 Markdown 渲染层。可以让 agent 通过单一 HTTP 端点(如 `GET /api/project-context/{repository_id}.md`)获取一份 AGENTS 风格的结构化 Markdown 文本,直接复制到目标仓库的 `AGENTS.md` 或 `.psco/context.md`。

#### 4.1.3 对第一件事的主张:更克制的版本

DPv4pro 主张"静态规则文件 → backend canonical 数据派生",意思是 `AGENTS.md` 中的"当前阶段"等信息应该从 backend 数据动态生成。

**我同意这个方向的长期价值,但主张 MVP0.5 阶段只做"静态文件去重与口径校准",不做"全量派生"。**

理由:

1. **全量派生的工程成本过高。** `AGENTS.md` / `plan.md` / `architecture_map.md` / `docs/README.md` 中包含大量叙述性内容(如"为什么这样设计"、"历史决策背景"),这些内容无法从 backend canonical 数据派生。
2. **全量派生会引入新的双源风险。** 如果 `AGENTS.md` 由 backend 派生,那么 backend 就成了"文档事实源",但 backend 的 phase 状态本身也需要被记录——这会形成新的循环依赖。
3. **MVP0.5 的真实痛点是"静态文件之间的重复与口径漂移",不是"静态文件与 backend 的派生关系"。**

因此我主张 MVP0.5 阶段做的是:

- 把"当前阶段状态"这类易漂移信息收敛到单一真相源(`plan.md`)
- `AGENTS.md` / `architecture_map.md` / `docs/README.md` 只引用 `plan.md` 的结论,不复制正文
- 这本身就是为 agent 消费做准备——agent 消费的静态文件应该是"单值一致"的,而不是"互相重复但可能漂移"的

这个主张会在第 5 节展开。

### 4.2 对 gemini31pro"主动注入"战术的判断

gemini31pro 提出:在 `CreateRepository Binding` 或 `Initialize Product` 时,PSCO 应主动将全局规范(`.trae/rules` / `TECH_STACK_BASELINE.md` / `project_skills.md` 模板)一键注入到目标物理仓库。

**我的判断:方向有价值,但 MVP0.5 阶段只做"受控可版本化的上下文镜像",不做"主动注入"。**

理由:

1. **"主动注入"意味着 PSCO 写入外部仓库,这本身就是一种"写入"行为。** 虽然不是写入 PSCO 自己的数据库,但写入外部仓库会引入文件冲突、版本控制、回滚等一系列复杂度。
2. **"主动注入"容易被误解为"PSCO 控制目标仓库",与"PSCO 是上下文系统,不是流程控制器"的最高原则冲突。**
3. **MVP0.5 阶段更稳妥的做法是:提供"上下文镜像导出"能力,由用户或 agent 自主决定是否复制到目标仓库。**

具体落地方式:

- `ProjectContextRead` 端点支持 `.md` 格式输出
- 用户在 web 端或 agent 在 IDE 现场可以通过 HTTP 获取这份 Markdown
- 是否复制到目标仓库、复制到哪个路径、如何版本化,由用户/agent 自主决定
- PSCO 不主动写入外部仓库

这个立场与 gemini31pro 的"武装 Agent"精神一致,但在落地方式上更克制:提供弹药,但不替 Agent 扣扳机。

---

## 5. 我的补充视角:作为非 GPT-5.4 模型的真实协作经验

本节是本文与三份已有评审的最大差异点。我作为 GLM-5.2,正是 `project_rules.md` §6 所定义的"非 GPT-5.4 模型"。我的真实协作经验印证了 GPT54 与 DPv4pro 的判断,但也暴露了一个三份评审都未充分展开的问题。

### 5.1 我的真实协作经验:每次接手 PSCO 都需要大量定向探索

`project_rules.md` §6 明确指出:

> 非 GPT-5.4 模型默认不具备 PSCO 的隐性上下文,每次接手必须先完成定向探索

我作为 GLM-5.2,正是这类模型的代表。在本次评审准备过程中,我实际执行了以下定向探索:

1. 阅读 `PSCO-mvp04-summarize-feedback.md`(当前最终共识,23K)
2. 阅读 `PSCO-mvp045-GPT54.md`(GPT54 刹车评审,11K)
3. 阅读 `PSCO-mvp045-DPv4pro.md`(DPv4pro 评审,13K)
4. 阅读 `PSCO-mvp045-gemini31pro.md`(gemini31pro 评审,4K)
5. 阅读 `plan.md`(全局开发预览,17K)
6. 阅读 `TECH_STACK_BASELINE.md`(技术基线,12K)
7. 阅读 `architecture_map.md`(目录结构,15K)
8. 阅读 `project_rules.md`(项目规则,10K)
9. 阅读 `docs/README.md`(文档总览,5K)
10. 阅读 `phase10-11` 正式验收 spec(7K)
11. 阅读 `phase10` shared_baseline(9K)
12. 阅读 `backend/internal/platform/router.go`(后端装配,9K)
13. 探查 12 个后端模块、9 个前端切片、12 个 `.proto` 文件、9 个 migration 的结构
14. 阅读 `PSCO-mvp04-summarize-feedback-GPT54.md`(上一轮汇总,17K)

这个探索过程本身就在印证 GPT54 与 DPv4pro 的核心判断:

> **当前 PSCO 最大的实际价值,就是降低 agent 的上下文解释成本。** 如果 PSCO 能让 agent 在进入项目现场时,通过单一入口获取完整上下文,而不需要阅读 14 份文档,这就已经创造了巨大的实际价值。

### 5.2 但三份评审都未充分展开的第六个重点问题:静态规则文件体系本身的双源风险

在我的定向探索过程中,我观察到一个三份评审都未充分展开的工程现实:

> **PSCO 当前的静态规则文件之间存在严重的重复与口径漂移风险,而这些文件正是 agent 当前已经在消费的上下文。**

具体观察:

#### 5.2.1 `AGENTS.md` §4 "当前状态" 与 `plan.md` §2 "当前进度概览" 内容高度重复

两份文档都列出了 `phase01 ~ phase10` 的完整状态清单,内容几乎一一对应。任何一次 phase 收口都需要同步更新两份文档,极易产生口径漂移。

#### 5.2.2 `architecture_map.md` §4.1 与 `docs/README.md` §3 内容高度重复

两份文档都列出了所有 phase 三件套、正式规格、验收报告的入口链接。任何一次新 phase 文档创建都需要同步更新两份文档。

#### 5.2.3 `project_rules.md` §2 与 `TECH_STACK_BASELINE.md` 内容高度重复

两份文档都描述了技术栈基线、双路线模型、选择规则、当前项目冻结选择。任何一次技术栈调整都需要同步更新两份文档。

#### 5.2.4 `PSCO-summarize-feedback.md` 不存在但被引用

`architecture_map.md` §1.1 与 `AGENTS.md` 都将 `PSCO-summarize-feedback.md` 列为"当前最终共识文档",但根目录实际无此文件。当前最终共识由 `PSCO-mvp04-summarize-feedback.md` 承担。这是口径漂移的实例——根级入口文档引用了一个不存在的文件。

### 5.3 第六个重点问题的正式提出

基于上述观察,我正式提出第六个重点问题:

> **在"PSCO 是上下文系统、agent 先消费后维护、web 与 agent 共享 backend core"的方向达成共识后,PSCO 当前的静态规则文件体系本身就需要先治"双源风险",再谈 agent 消费。**

这个问题的重要性在于:

1. **agent 当前已经在消费这些静态文件。** `AGENTS.md` / `project_rules.md` / `TECH_STACK_BASELINE.md` 是 agent(包括我)接手 PSCO 时阅读的第一批文档。
2. **如果这些文件本身有双源风险,agent 消费的就是"可能已经漂移的上下文"。** 这比"建新通道"更紧迫——先修好已有通道,再谈新通道。
3. **这与 DPv4pro 的第一优先级方向一致,但主张更克制的版本:** 不需要立即做"静态文件 → backend 派生"(这是远期目标),但需要立即做"静态文件之间的去重与口径校准"。

### 5.4 第六个重点问题的具体落地建议

#### 5.4.1 把"当前阶段状态"收敛到 `plan.md` 作为单一真相源

- `plan.md` §2 "当前进度概览" 是 phase 状态的唯一真相源
- `AGENTS.md` §4 "当前状态" 只保留"当前阶段"与"当前主目标"两行总结,不复制完整 phase 列表,正文指向 `plan.md`
- `docs/README.md` §4 "当前阶段状态" 同样只保留总结,正文指向 `plan.md`

#### 5.4.2 把"目录结构与文档落点"收敛到 `architecture_map.md` 作为单一真相源

- `architecture_map.md` §4 是文档落点的唯一真相源
- `docs/README.md` §3 "当前最常用入口" 只保留 top-5 入口,完整列表指向 `architecture_map.md`

#### 5.4.3 把"技术栈基线"收敛到 `TECH_STACK_BASELINE.md` 作为单一真相源

- `project_rules.md` §2 只保留"当前项目冻结选择"与"禁止自由发挥"两小节,技术栈正文指向 `TECH_STACK_BASELINE.md`

#### 5.4.4 修复 `PSCO-summarize-feedback.md` 缺失问题

- 当前根级入口文档引用了不存在的 `PSCO-summarize-feedback.md`
- 需要先与用户确认:这是待补建的根级占位文件,还是命名口径漂移?
- 如果是后者,需要更新 `architecture_map.md` 与 `AGENTS.md` 中的引用,统一指向 `PSCO-mvp04-summarize-feedback.md`

#### 5.4.5 这件事应该在 MVP0.5 第一阶段(方向收敛)完成,不需要写代码

这个去重与口径校准工作本身是文档层面的,不需要写代码。它应该在 MVP0.5 第一阶段(方向收敛 phase)完成,作为"为 agent 消费做准备"的最小前置工作。

---

## 6. 我的明确推进路线建议

基于以上分析,我提出以下路线建议:

### 6.1 不直接进入"Agent Consumption Layer"全面实现

原因:

1. GPT54 提出的四个核心问题 + DPv4pro 补充的第五问 + 我补充的第六问,尚未在多专家讨论中完全收敛;
2. 静态规则文件的双源风险治理应先于新通道建设;
3. 当前已有 12 个 ConnectRPC 服务作为 agent 可消费的基础设施,但需要先理解 agent 真实消费场景,再决定是否需要新增聚合端点。

### 6.2 下一阶段建议采用"收敛 + 最小验证"两段式结构

这与 DPv4pro 的建议一致,但我在第一阶段加入"静态文件去重"作为必要前置工作。

#### 6.2.1 第一阶段:方向收敛 + 静态文件去重(1 个轻量 phase,几乎不写代码)

目标:不写代码(或仅写最小代码),只产出共识文档与去重后的根级文件。

交付物:

1. 四实体关系的正式语义澄清文档(在现有实现基础上做"语义澄清",不是"结构重定义")
2. PSCO 作为"上下文系统"的正式定位声明(写入 `shared_baseline`)
3. agent 渠道"先消费、后维护"的正式原则冻结(写入 `shared_baseline`)
4. web 与 agent 的正式分工边界说明(写入 `shared_baseline`)
5. **静态规则文件去重与口径校准**(具体落地见 §5.4)
6. **修复 `PSCO-summarize-feedback.md` 缺失问题**(先与用户确认)

#### 6.2.2 第二阶段:最小可验证交付(1 个交付 phase)

目标:用最小代码量验证"agent 可消费 PSCO 上下文"是否成立。

交付物:

1. **项目上下文聚合导出只读端点**(`ProjectContextRead` 或等价 RPC)
   - 按 `repository_id` 或 `product_id` 聚合输出完整上下文
   - 完全只读,不引入任何写入语义
   - 输出格式与 `.proto` 合同单值一致
   - 新增 `psco/project_context/v1/project_context.proto` 合同
2. **AGENTS 风格上下文导出**(Markdown 输出)
   - 基于 `ProjectContextRead` 的输出,生成 AGENTS 风格的结构化 Markdown 文本
   - 可通过 HTTP 端点获取(如 `GET /api/project-context/{repository_id}.md`)
   - 由用户或 agent 自主决定是否复制到目标仓库(PSCO 不主动写入外部仓库)

明确不做:

- MCP 协议实现
- CLI 工具
- agent 写入路径(包括 gemini31pro 的 Draft 提交接口)
- 前端对话式 agent 入口
- 四实体结构重定义
- 静态规则文件全量派生(远期目标,不是 MVP0.5 范围)
- 主动注入外部仓库(gemini31pro 的"主动注入"战术,改为"受控导出")

### 6.3 对 GPT54 七个"不宜直接主推"方向的确认

我完全同意 GPT54 列出的七个禁区:

1. web 前端对话框式 agent 入口
2. agent 自动写入主线
3. 知识图谱与自动扫描
4. GitHub / Gitea 重型集成
5. Module 默认独立仓库化
6. PSCO 作为开发流程控制器
7. 在未收敛实体关系前冻结具体 API / MCP / CLI 细节

补充两个禁区:

8. **在 MVP0.5 阶段设计任何形式的 agent 写入路径,包括 Draft 提交接口。** 理由见 §4.1.4。
9. **在未验证"agent 能稳定消费"之前,不建任何新的传输协议层(MCP/CLI)。** 这与 DPv4pro 的补充禁区一致——先用现有 ConnectRPC + HTTP 基础设施验证消费链路是否成立,再决定是否需要额外的协议层。

---

## 7. 建议后续专家重点评价的六个问题

为避免后续专家讨论分散,我建议在 GPT54 四问 + DPv4pro 第五问基础上,补充第六问,形成以下六个评价焦点:

1. **四实体关系应如何澄清**(GPT54 原问题):是否同意"语义澄清而非结构重定义"的立场?
2. **PSCO 的定位**(GPT54 原问题):是否同意 PSCO 是"上下文系统"而非"开发流程控制器"?
3. **agent 渠道的优先级**(GPT54 原问题):是否同意"先消费、后维护"?
4. **web 与 agent 的分工**(GPT54 原问题):是否同意"web 偏全局管理、agent 偏现场消费、共享 backend core"?
5. **最小可验证交付物**(DPv4pro 补充):在上述四点达成共识后,下一阶段的最小可验证交付物应是什么?是否同意"项目上下文聚合端点 + AGENTS 风格导出"作为第一优先级?
6. **静态规则文件双源风险治理**(GLM-5.2 补充):是否同意在 MVP0.5 第一阶段先完成"静态文件去重与口径校准",作为 agent 消费的前置工作?是否同意把"静态文件 → backend 全量派生"后移到远期目标?

如果以上六点不能先收敛,后续实现层讨论将很难稳定。

---

## 8. 结论

GPT54 的"战略刹车"是必要且及时的。三份已有评审(GPT54 / DPv4pro / gemini31pro)已经在"PSCO 是上下文系统、agent 先消费后维护、web 与 agent 共享 backend core"上形成稳定共识。我对 GPT54 四主张的判断如下:

| 主张 | 判断 | 理由 |
|---|---|---|
| 重新定义四实体关系 | 方向正确,但应"语义澄清 + 局部增强" | 10 个 phase 的实现已稳定,四实体语义已被现有结构满足,不需要结构重定义 |
| PSCO 不应成为流程控制器 | 完全同意 | 这是 GPT54 文档中最关键的判断,且 `phase10` 实现已部分落地 |
| agent 先消费后维护 | 同意原则,需补充路径 | 当前 12 个 ConnectRPC 服务已是 agent 可消费基础设施,第一优先级是新增聚合端点,不是建新协议层 |
| web 与 agent 并行不称 | 完全同意 | 与当前实现高度吻合,`router.go` 装配方式天然支持双渠道共享 backend core |

我对 DPv4pro 与 gemini31pro 关键分歧点的判断:

| 分歧点 | 判断 | 理由 |
|---|---|---|
| 静态文件 → backend 派生 | 同意方向,但 MVP0.5 只做"去重与口径校准",不做"全量派生" | 全量派生工程成本过高,且会引入新的双源风险 |
| gemini31pro 主动注入 | 方向有价值,但 MVP0.5 只做"受控导出",不做"主动注入" | 主动写入外部仓库与"上下文系统"定位有张力 |
| gemini31pro 审批门禁 | 同意原则,但 MVP0.5 严禁任何形式的 agent 写入路径设计 | Draft 接口本身就是 agent 专属领域模型的变体 |

我补充的核心论点是:

> **当前 PSCO 最紧迫的事不是建新的 agent 通道,而是先把 agent 已经在消费的静态规则文件体系治好"双源风险"。`AGENTS.md` / `plan.md` / `architecture_map.md` / `docs/README.md` 之间存在大量重复,任何一次状态变更都需要同步更新多个文件,极易产生口径漂移。如果 agent 消费的是已经漂移的上下文,建新通道只会放大错误。**

如果用一句话总结我的主张:

> **下一阶段先收敛"PSCO 是上下文系统"的共识,同步完成静态规则文件去重与口径校准,再以"项目上下文聚合端点 + AGENTS 风格导出"作为最小可验证交付物,验证 agent 是否能稳定消费 PSCO 上下文——而不是先建 MCP/CLI/新协议层,也不是先做四实体结构重定义。**

我作为 GLM-5.2 的独特贡献,是基于"非 GPT-5.4 模型真实协作经验"的视角:我每次接手 PSCO 都需要阅读 14+ 份文档才能恢复正确上下文,这个痛点本身就在印证"PSCO 当前最大价值是降低 agent 上下文解释成本"。但如果 PSCO 自己的静态文件体系都有双源风险,这个价值就无法稳定兑现。先治本,再扩张。
