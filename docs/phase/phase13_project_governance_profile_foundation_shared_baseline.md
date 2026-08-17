# phase13_project_governance_profile_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase13` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级入口中重复发散。

`phase13` 当前处于 `/plan` 阶段。本文档只承接当前阶段的单值基线、信息分层矩阵、项目范式 v1、全局规范资产矩阵与 agent 项目简报前提，不替代后续 `/spec`、实现或根级状态正文。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase13_project_governance_profile_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase12_semantic_alignment_and_readonly_consumption_foundation` 已完成正式收口，`phase13` 当前作为下一阶段正式 `/plan` 入口建立
- 当前阶段规划上游统一以 `PSCO-mvp05-summarize-feedback.md`、当前根级真相源文档与 `phase12` 正式收口结论为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp05-summarize-feedback.md`
  - `AGENTS.md`
  - `plan.md`
  - `project_rules.md`
  - `TECH_STACK_BASELINE.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `docs/phase/README.md`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_*`
  - `.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md`
- 当前阶段只承接：
  - 项目级治理画像
  - 全局规范资产
  - agent 项目简报输入
- 当前阶段不反向重写 `phase02 ~ phase12` 已完成的正式交付

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web`
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http + ConnectRPC`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- 运行方式：`Single Server First`

### 2.4 当前阶段特别约束

- `.proto` 仍是唯一长期合同源
- `ConnectRPC` 仍是业务接口正式传输层
- 四实体业务主线继续保持：
  - `Product`
  - `Repository`
  - `Module`
  - `Decision`
- 当前阶段新增承接层不得演化为第五个业务主实体
- 当前阶段不允许：
  - 把 IDE 目录即时上下文默认上升为 PSCO 正式事实源
  - 引入 Git 推进跟踪主线
  - 引入模板仓库自动接入
  - 引入目录全文扫描入库
  - 引入 agent 自动写回、MCP / CLI 或审批流
  - 让 Web 与 agent 各自长出第二套项目治理解释

### 2.5 当前阶段交付模式

- `phase13` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可维护、可读取、可被 agent 消费的“项目治理画像 + 全局规范资产 + agent 项目简报”主线代码

## 3. 当前阶段能力矩阵

### 3.1 `Project Governance Profile Foundation` 单值定义

`Project Governance Profile Foundation` 在当前阶段冻结为：

- 项目级治理画像
- 全局规范资产
- agent 项目简报输入

当前阶段不把以下内容解释为该主交付：

- Git 推进跟踪
- 模板仓库 bootstrap
- 自动同步
- IDE 插件化
- agent 写回
- 第五个业务主实体

### 3.2 三层信息分层矩阵

#### A. `PSCO-native facts`

- 四实体信息与关系
- 全局规范资产
- 项目级治理画像
- 当前阶段入口与状态

#### B. `IDE-accessible context`

- 当前工作区文件全文
- 目录局部漂移
- 即时实现细节
- agent 在 IDE 现场自行读取的项目文件

#### C. `Controlled synced projection`

- Git 推进摘要
- 模板来源与模板版本
- 自动存在性校验
- 自动状态建议

补充冻结：

- 当前阶段只正式承接 A 层
- B 层默认继续由 IDE / agent 现场读取
- C 层只保留为后续进入条件，不在本阶段起点直接实现

### 3.3 项目级治理画像最小字段矩阵

当前阶段冻结的最小字段集合如下：

- `project_profile_version`
- `track_type`
- `template_source`
- `canonical_root_files[]`
  - `file_name`
  - `role`
  - `required`
- `global_constraint_refs[]`
  - `name`
  - `kind`
  - `entry_ref`
- `current_phase_name`
- `current_phase_ref`
- `current_phase_status`

补充冻结：

- 第一版优先手工维护，不默认自动推断
- 第一版优先管理“合同”，不是先管理“扫描结果”
- 存在性校验、自动建议与同步增强属于后续受控进入项

### 3.4 当前项目范式 v1 矩阵

当前项目范式 v1 继续冻结为以下目录与根级文件集合：

#### 目录结构

- `backend/`
- `database/`
- `frontend/`
- `docs/`
  - `audit/`
  - `fix/`
  - `phase/`
  - `review/`
- `proto/`

#### 根级文件

- `.env`
- `AGENTS.md`
- `architecture_map.md`
- `plan.md`
- `global_skills.md`
- `project_skills.md`
- `project_rules.md`
- `README.md`
- `TECH_STACK_BASELINE.md`

补充冻结：

- `docs/README.md` 继续作为 docs workflow 总入口
- 当前项目范式 v1 先作为 `PSCO` 当前默认项目范式被结构化承接
- 当前阶段不要求立刻把它扩写成模板仓库或自动脚手架

### 3.5 全局规范资产矩阵

当前阶段至少应显式承接以下全局规范资产：

- `project_rules.md`
- `TECH_STACK_BASELINE.md`
- `AGENTS.md`
- `architecture_map.md`
- `plan.md`
- `README.md`
- `global_skills.md`
- `project_skills.md`

补充冻结：

- 当前阶段优先管理这些资产的结构化摘要、职责与入口关系
- markdown 正文继续保留在仓库内作为正文载体
- 不把这些文件的真实路径直接做成面向普通用户的主内容
- 第一版逐项承接策略冻结如下：
  - `project_rules.md`：`entry_ref + role + structured summary`
  - `TECH_STACK_BASELINE.md`：`entry_ref + role + structured summary`
  - `AGENTS.md`：`entry_ref + role + structured summary`
  - `architecture_map.md`：`entry_ref + role + structured summary`
  - `plan.md`：`entry_ref + role + structured summary`
  - `README.md`：`entry_ref + role`
  - `global_skills.md`：`entry_ref + role`
  - `project_skills.md`：`entry_ref + role`
- 第一版统一允许 markdown 正文回源，但不把上述文件全文入库

### 3.6 agent 项目简报矩阵

当前阶段的 agent 项目简报至少应覆盖：

- `repository`
- 当前项目治理画像
- 当前项目继承的全局规范资产
- 当前阶段入口与状态
- `products[]`
- `modules[]`
- `decisions[]`

补充冻结：

- agent 简报是结构化只读输入，不是目录扫描结果回放
- agent 简报与 IDE 自带目录读取能力是协作关系，不是替代关系
- agent 若需读取目录全文，继续由 IDE / agent 现场能力自行完成
- `repository_id` 是第一版唯一正式锚点
- `products / modules / decisions` 只能从同一 `repository_id` 驱动的 PSCO 结构化关系中解析
- 若存在多个 `Module / Decision`，必须返回数组摘要，不得伪造单一“当前 module / decision”

### 3.7 前端承接矩阵

当前阶段前端正式承接策略冻结如下：

- 项目治理画像优先作为项目级治理层承接
- 它不应在四实体详情页内重复长成第二主内容区
- 第一版前端正式承接位冻结为 `Repository detail`
- `phase13` 收口前不新增独立项目治理画像页面或并列第二入口
- 目录真实路径、全文规则与阶段快照不应误做成面向普通用户的大块说明 UI

## 4. 当前阶段验收前提

当前阶段验收必须至少回答：

1. PSCO 是否已能正式管理四实体之外的项目级治理画像，而不新增第五个业务主实体
2. 当前项目范式 v1 是否已被结构化承接，而不是只停留在脑内共识
3. 全局规范资产是否已被 PSCO 正式管理为结构化摘要与入口关系
4. agent 是否已能读取“全局信息 + 当前项目 PSCO 管理信息”的最小项目简报
5. 当前阶段是否仍严格没有引入 Git 推进跟踪、模板仓库自动接入、自动同步与 agent 写回
6. Web 与 agent 是否仍共享同一套 `PSCO-native facts`

补充冻结的第一版正式验收协议：

- 固定样本：
  - `Repository`：`personal-software-company-os`
  - `repository_id`：`ca261521-8daf-4248-8f12-43525326e759`
  - `Product`：`PSCO`
- 固定 Web 页面：
  - `/repositories/$repositoryId`
- 固定 agent 入口：
  - 基于同一 `repository_id` 返回的 `project brief for agent`
  - `AGENTS.md`
  - `plan.md`
  - 治理画像承接的全局规范资产结构化结果
- 固定问题：
  1. 当前项目治理画像版本与技术路线是什么，在哪个固定入口可确认
  2. 当前 canonical 根级文件集合是否已被正式承接，在哪个固定入口可确认
  3. 当前全局规范资产是否以结构化摘要 + 入口关系被正式承接，在哪个固定入口可确认
  4. 当前 agent 项目简报是否由同一 `repository_id` 驱动，且没有伪造第二套目录扫描结果
  5. 当前第一版前端正式承接位是否仍是 `Repository detail`，而没有长出并列第二入口
  6. 当前阶段是否仍严格没有进入 Git 推进跟踪、模板仓库接入、自动同步与 agent 写回
- 固定 rerun 记录：
  - 使用了哪个固定样本与 `repository_id`
  - 使用了哪个固定 Web 页面与 agent 入口
  - 每个问题的回答结果
  - 失败点与是否达标
