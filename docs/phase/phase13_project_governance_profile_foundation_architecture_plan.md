# phase13_project_governance_profile_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase13_project_governance_profile_foundation` 的架构规划文档。

`phase12_semantic_alignment_and_readonly_consumption_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并把四实体在 Web / agent 共享只读消费面上的语义表达重新收口到同一条主线。与此同时，本阶段后的新共识也已明确：

- PSCO 不应把 IDE 已打开项目目录中的即时文件上下文，默认上升为自己的正式事实源；
- agent 在 IDE 场景中一次只服务一个项目，因此它首先需要的是“PSCO 管理的当前项目信息 + 全局长期规范资产”，而不是第二套目录扫描机制；
- 当前更值得优先沉淀的，是四实体信息、全局规范与约束、项目级治理画像，以及可被 agent 直接消费的项目简报输入；
- Git 推进跟踪、模板仓库接入、自动同步与更重受控维护能力，属于后续受控进入项，而不是 `phase13` 起点。

因此，`phase13` 的职责不是：

- 直接建设 Git 推进跟踪平台；
- 直接建设模板仓库自动化或项目脚手架生成器；
- 把 PSCO 做成 IDE 插件、目录扫描器或开发流程控制器；
- 继续放大 `phase12` 的共享只读 UI 表达；
- 直接进入 MCP / CLI / agent 写回。

`phase13` 的职责是交付单一主交付能力：

> **Project Governance Profile Foundation**

也就是：

1. 正式冻结 PSCO 应管理的项目级治理信息边界；
2. 正式冻结当前项目范式 v1、根级 canonical 文件集合与全局规范资产承接方式；
3. 正式定义供 agent 在 IDE 场景读取的最小项目简报输入；
4. 为后续项目推进跟踪、模板接入与自动同步建立明确进入条件。

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `project_rules.md`
4. `TECH_STACK_BASELINE.md`
5. `architecture_map.md`
6. `docs/README.md`
7. `docs/phase/README.md`
8. `PSCO-mvp05-summarize-feedback.md`
9. `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
10. `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
11. `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`
12. `.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md`

补充说明：

- `PSCO-mvp05-summarize-feedback.md` 继续作为当前 `mvp0.5` 的唯一直接共识上游；
- `phase12` 提供的是“前端四实体语义一致性收口 + Web / agent 共享只读消费深化”的已完成交付，不再承担下一阶段事实源扩张职责；
- 当前阶段直接承接用户已明确的新共识：PSCO 优先管理四实体信息、全局规范资产与项目级治理画像，而不是替 IDE 复刻目录理解机制。

## 3. 本阶段目标

`phase13` 的目标是：

> 在不侵入 IDE 现场机制、不直接进入自动同步与推进跟踪的前提下，让 PSCO 正式承接“项目级治理画像 + 全局规范资产 + agent 项目简报输入”的最小主线。

本阶段需要回答的核心问题：

1. 哪些信息属于 `PSCO-native facts`，哪些只是 `IDE-accessible context`，哪些属于后续才允许进入的受控同步投影；
2. 项目级治理层如何承接目录结构、根级 canonical 文件、全局规范与当前阶段状态，而不污染四实体主线；
3. 当前项目范式 v1 应冻结哪些目录、哪些根级文件、哪些规范入口；
4. agent 在 IDE 场景下应从 PSCO 读取哪一份最小项目简报；
5. Git 推进跟踪、模板仓库接入、自动同步要满足什么进入条件，才允许进入下一阶段讨论。

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase13` 必须直接承接：

- `PSCO-mvp05-summarize-feedback.md`
- 当前根级真相源文档
- `phase12` 三件套
- `phase12-11` 正式验收与收口入口

不允许在本阶段重新解释：

- `phase12` 已完成四实体语义与只读消费收口的事实；
- `phase11` 与 `phase12` 已明确的“先消费、后维护”顺序；
- `.proto + ConnectRPC` 为 Go 业务接口正式主线的事实；
- 当前项目冻结为 `Durable System Track` 的事实。

### 4.2 当前阶段单一主交付能力

`phase13` 的单一主交付能力冻结为：

> **让 PSCO 正式管理当前项目的治理画像、全局规范资产与 agent 可读项目简报。**

它由三部分组成：

1. **项目级治理画像**
2. **全局规范资产**
3. **agent 项目简报输入**

补充冻结：

- 三者必须在同一交付型 phase 中协同定义，但仍共享同一条四实体主线；
- 本阶段优先解决“PSCO 该管理什么”，而不是“如何自动同步一切”；
- 本阶段允许建立新的项目级承接层，但不允许因此长出第五个业务主实体。

### 4.3 信息分层正式边界

当前阶段冻结三层信息模型如下：

1. **`PSCO-native facts`**
   - 四实体信息与关系：`Product / Repository / Module / Decision`
   - 全局规范资产：长期有效的 workflow、技术基线、协作约束、目录与文档职责
   - 项目级治理画像：项目范式版本、canonical 根级文件集合、docs workflow 布局、当前阶段入口与状态
2. **`IDE-accessible context`**
   - 当前工作区文件全文
   - 临时目录漂移
   - 局部实现细节
   - agent 现场读取的即时文件内容
   - 这些信息默认继续由 IDE / agent 自行读取，不默认上升为 PSCO 正式事实
3. **`Controlled synced projection`**
   - Git 推进摘要
   - 模板仓库来源
   - 自动存在性校验
   - 自动状态建议
   - 这些属于后续受控进入项，不在 `phase13` 起点直接实现

### 4.4 项目级治理层的正式定位

当前阶段继续冻结：

- 项目级治理层不是第五个业务主实体；
- 它是四实体之外的**项目级承接层 / 治理层 / 合同层**；
- 它服务的对象是“当前项目如何被 PSCO 理解与约束”，而不是替代 `Product / Repository / Module / Decision` 的业务含义；
- 第一版前端正式承接位冻结为 `Repository detail`，因为 `Repository` 是代码仓库身份对象与项目锚点，但不得把该治理层偷换成 `Repository` 的业务字段扩张；
- 在 `phase13` 正式收口前，不新增并列的独立项目治理画像页面或第二入口。

### 4.5 当前项目范式 v1

当前阶段冻结当前项目范式 v1 的目录结构如下：

```text
backend/
database/
frontend/
docs/
  audit/
  fix/
  phase/
  review/
proto/
.env
AGENTS.md
architecture_map.md
plan.md
global_skills.md
project_skills.md
project_rules.md
README.md
TECH_STACK_BASELINE.md
```

补充冻结：

- `docs/README.md` 继续作为 docs workflow 总入口；
- 当前目录结构与根级文件集合作为项目治理画像的直接输入，但不等于“要求所有消费侧项目立刻复制同一目录”；
- 当前阶段先把它冻结为 `PSCO` 当前默认项目范式 v1，不提前扩写为完整模板平台。

### 4.6 全局规范资产的正式边界

当前阶段继续冻结：

- `project_rules.md`
- `TECH_STACK_BASELINE.md`
- `AGENTS.md`
- `architecture_map.md`
- `plan.md`
- `README.md`
- `global_skills.md`
- `project_skills.md`

这些文件中承载的是**长期有效、跨项目不应只停留在脑中的规范与约束**。

当前阶段正式结论：

- PSCO 优先管理这些全局规范资产的**结构化摘要、职责、入口关系与当前采用状态**；
- 当前阶段不强制把这些 markdown 全文都拆成数据库字段；
- 当前阶段允许保留“结构化摘要 + markdown 正文回源”的双层模式；
- 目录中的真实文件路径不应在普通产品 UI 中被误当成面向最终用户的主内容。
- 第一版必须输出逐项承接矩阵，至少明确：
  - 哪些文件只存 `entry_ref + role`
  - 哪些文件存 `entry_ref + role + structured summary`
  - 哪些文件允许 markdown 正文回源但不入库
  - 哪些文件第一版禁止全文入库

### 4.7 agent 项目简报输入的正式边界

当前阶段继续冻结：

- agent 在 IDE 场景中一次只服务一个项目；
- 它优先需要从 PSCO 读取：
  - 当前项目的四实体信息与关系
  - 当前项目的治理画像
  - 当前项目继承的全局规范资产
  - 当前阶段入口与状态
- agent 不要求 PSCO 替它复刻整个目录读取能力；
- agent 不应通过 PSCO 默认读取“目录里所有文件全文”。

当前阶段正式结论：

- `phase13` 必须输出一份正式的 **project brief for agent**；
- 该 brief 是受控的结构化输入，不是第二套目录扫描结果；
- 后续若进入 Git 推进跟踪或模板自动接入，应在此 brief 基础上增量扩展，而不是另起第二套 agent 输入协议；
- brief 的第一版最小 schema 必须至少覆盖：
  - `repository`
  - `governance_profile`
  - `global_assets`
  - `current_phase`
  - `products[]`
  - `modules[]`
  - `decisions[]`
- brief 的第一版解析协议必须冻结：
  1. `repository_id` 是唯一正式锚点
  2. `products / modules / decisions` 只能从同一 `repository_id` 驱动的 PSCO 结构化关系中解析
  3. 多个 `Module / Decision` 必须以数组摘要返回，不得伪造单一“当前 module / decision”
  4. 目录全文读取继续留给 IDE / agent 现场能力，不通过 PSCO brief 补第二套目录扫描

### 4.8 当前阶段验收协议前提

当前阶段必须冻结可机械 rerun 的正式验收协议，至少包括：

- 固定样本：继续使用 `phase11 / phase12` 的 `PSCO` dogfooding 样本与固定 `repository_id`
- 固定 Web 页面：第一版冻结为 `/repositories/$repositoryId`
- 固定 agent 入口：固定为同一 `repository_id` 驱动的 `project brief for agent`、`AGENTS.md`、`plan.md` 与治理画像承接的全局规范资产结果
- 固定问题：至少覆盖治理画像版本与技术路线、canonical 根级文件集合、全局规范资产承接、brief 锚点一致性、第一版前端唯一承接位与当前阶段非目标
- 固定 rerun 记录格式：必须显式记录样本、入口、问题结果、失败点与是否达标

### 4.9 当前阶段明确不做

本阶段明确不做：

1. Git 推进跟踪主线
2. 模板仓库 bootstrap / clone / pull 自动化
3. 目录或文件全文自动扫描入库
4. MCP / CLI / agent 写回 / Draft / 审批流
5. IDE 插件化、会话劫持或开发流程控制台
6. 第五个业务主实体
7. 第二套与四实体并列的事实源
