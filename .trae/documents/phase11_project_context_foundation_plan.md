# `phase11_project_context_foundation` /plan 执行计划

## Summary

基于 [PSCO-mvp05-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md) 已冻结的 `mvp0.5` 共识，下一步将执行一次正式 `/plan`，建立 `phase11_project_context_foundation` 的三件套：

- `docs/phase/phase11_project_context_foundation_architecture_plan.md`
- `docs/phase/phase11_project_context_foundation_dev_plan.md`
- `docs/phase/phase11_project_context_foundation_shared_baseline.md`

本次 `/plan` 的目标不是进入实现细节，而是为下一阶段单一主交付能力建立正式规划入口，并同步完成开启新 phase 所需的根级入口更新。

本次 `/plan` 已锁定的关键决策：

1. 下一阶段正式 phase 名称冻结为 `phase11_project_context_foundation`
2. 根级“最终共识”入口在本轮治理中统一改为直接指向 `PSCO-mvp05-summarize-feedback.md`
3. 本阶段单一主交付能力冻结为：**根级上下文真相源治理 + 最小只读项目上下文导出**
4. 本阶段不拆成“纯方向收敛 phase + 实现 phase”两段式；而是在一个交付型 phase 内承接边界收敛、治理、设计、实现、验收与根级同步

## Current State Analysis

### 1. 当前 phase 体系与命名模式

经仓库探索，`docs/phase/` 当前已存在 `phase01 ~ phase10` 的完整三件套，命名模式固定为：

- `phaseXX_*_architecture_plan.md`
- `phaseXX_*_dev_plan.md`
- `phaseXX_*_shared_baseline.md`

最近完成正式业务 phase 为：

- [phase10_asset_action_closure_foundation_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md)
- [phase10_asset_action_closure_foundation_dev_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase10_asset_action_closure_foundation_dev_plan.md)
- [phase10_asset_action_closure_foundation_shared_baseline.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md)

因此，下一阶段采用 `phase11_project_context_foundation` 作为正式 phase 名称，符合当前仓库的既有模式。

### 2. 当前根级状态仍停留在 `phase10`

以下根级/入口文档目前仍把“当前阶段入口”停留在 `phase10` 或“`phase10` 收口后进入 Agent Consumption Layer”的口径上：

- [plan.md](file:///home/dell/Projects/personal-software-company-os/plan.md)
- [AGENTS.md](file:///home/dell/Projects/personal-software-company-os/AGENTS.md)
- [architecture_map.md](file:///home/dell/Projects/personal-software-company-os/architecture_map.md)
- [docs/README.md](file:///home/dell/Projects/personal-software-company-os/docs/README.md)
- [docs/phase/README.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/README.md)

根据 `project_rules.md` §4.1 与 §5，开启新 phase 前必须同步这些入口，避免出现“新 phase 三件套已建立，但根级入口仍停在旧 phase”的状态。

### 3. 最新共识文档已切换为 `mvp0.5`

根目录中已存在新的共识文档：

- [PSCO-mvp05-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md)

且其正式冻结了：

- PSCO 的定位为“上下文系统”，不是“开发流程控制器”
- agent 路线先消费、后维护
- web 不退化，但 agent 只以受控最小能力面进入
- 当前阶段的单一主交付方向为：
  - 根级上下文真相源治理
  - 最小只读项目上下文导出

因此，`phase11` 的 `/plan` 必须直接承接 `PSCO-mvp05-summarize-feedback.md`，而不再以 `PSCO-mvp04-summarize-feedback.md` 作为最新直接上游。

### 4. 当前存在已确认的根级入口漂移问题

探索结果已确认：

- `project_rules.md` §1 仍写有“最终共识只以 `PSCO-summarize-feedback.md` 为准”
- [docs/README.md](file:///home/dell/Projects/personal-software-company-os/docs/README.md) 当前仍有对不存在文件 `PSCO-summarize-feedback.md` 的引用
- 用户已明确选择：本轮 phase 中将根级“最终共识”入口**统一改为直接指向** `PSCO-mvp05-summarize-feedback.md`，而不是补建薄指针文件

这意味着 `phase11` 的 `/plan` 不仅要规划“以后怎么治理”，还要把“本轮怎么统一根级共识入口”写入边界与子任务。

### 5. 当前对下一阶段的实现方向已足够收敛

结合 `PSCO-mvp05-summarize-feedback.md` 与五份 `mvp0.45` 评审，可视为已稳定的共识包括：

- 不做四实体结构重构，只做语义澄清
- 不先做 MCP / CLI / 前端对话框 / 写回接口
- 先做项目上下文聚合只读导出
- 先做根级上下文真相源治理
- 以 PSCO 仓库自身作为 dogfooding 场景验证上下文恢复成本下降

因此，当前已不存在必须阻止 `/plan` 继续推进的高影响未知项。

## Assumptions & Decisions

### 已锁定决策

1. **Phase 名称**：`phase11_project_context_foundation`
2. **最新共识上游**：`PSCO-mvp05-summarize-feedback.md`
3. **共识入口统一方式**：将根级引用直接改为指向 `PSCO-mvp05-summarize-feedback.md`
4. **阶段组织方式**：单一交付型 phase，不拆纯讨论 phase
5. **阶段单一主交付能力**：让 agent 通过单一稳定入口读取当前项目核心上下文
6. **本阶段前置治理**：根级上下文真相源治理必须作为 phase 内正式子任务进入，而不是停留在评审文档口号层

### 执行假设

1. 本次 `/plan` 的实现范围包括：
   - 产出 `docs/phase/` 下的三件套
   - 同步开启新 phase 所需的根级入口文档
2. 本次 `/plan` **不**产出 `.trae/specs/` 下的执行规格，不进入实现代码
3. `project_rules.md` 作为规则基线需复读校验，但未必必须在本轮修改正文；若其“`PSCO-summarize-feedback.md` 为唯一最终共识”与用户最新明确决策冲突，则应在三件套中显式记为根级治理对象，并在根级同步时对齐
4. `phase11` 仍需沿用既有 `phase02+` 交付型 phase 模式：architecture / dev / shared baseline 三份文档各司其职，不互相复制正文

## Proposed Changes

### A. 新增三件套

#### 1. `docs/phase/phase11_project_context_foundation_architecture_plan.md`

**What**

建立 `phase11` 的架构规划文档。

**Why**

需要明确本阶段的：

- 单一主交付能力
- 直接上游输入
- 架构冻结结论
- 根级真相源治理与项目上下文导出的关系
- 不做范围

**How**

文档应直接承接：

- [PSCO-mvp05-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md)
- [project_rules.md](file:///home/dell/Projects/personal-software-company-os/project_rules.md)
- [plan.md](file:///home/dell/Projects/personal-software-company-os/plan.md)
- [docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md)

并冻结以下核心结论：

- `phase11` 不再处理 `Asset-Action Closure`
- `phase11` 的主交付不是 MCP/CLI/写回
- `phase11` 只做“根级上下文真相源治理 + 最小只读项目上下文导出”
- `Project Context` 是聚合投影，不是新实体主线
- 根级治理遵循“单一写者执行”，而不是全量 backend 派生

#### 2. `docs/phase/phase11_project_context_foundation_dev_plan.md`

**What**

建立 `phase11` 的开发执行拆分文档。

**Why**

需要把本阶段拆成符合 `project_rules.md` 要求的五类子任务：

- 边界收敛类
- 实现设计类
- 源代码实现类
- 验证验收类
- 根级同步类

**How**

Dev plan 应按单一主交付能力展开，至少包含以下子任务组：

1. 边界收敛：
   - 冻结上下文系统定位
   - 冻结 web / agent 分工
   - 冻结四实体语义确认口径
   - 冻结不做清单
2. 根级治理设计：
   - 明确 phase 状态、目录落点、技术栈正文的唯一承接位
   - 明确重复承载的收敛方式
   - 明确 `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识入口的统一改写策略
3. 项目上下文导出设计：
   - 聚合只读读取能力的输入锚点
   - 聚合输出内容边界
   - AGENTS 风格 Markdown 导出职责
4. 源代码实现：
   - 最小只读聚合读取能力
   - 最小导出能力
5. 验证验收：
   - 以 PSCO 仓库自身做 dogfooding
   - 验证上下文恢复成本下降
   - 验证根级文档重复与悬空引用收敛
6. 根级同步：
   - 更新 `AGENTS.md / plan.md / architecture_map.md / docs/README.md / docs/phase/README.md`

#### 3. `docs/phase/phase11_project_context_foundation_shared_baseline.md`

**What**

建立 `phase11` 的共享基线文档。

**Why**

需要集中冻结当前阶段的单值基线，避免架构规划、dev plan 与后续 `/spec` 出现三套说法。

**How**

文档应至少冻结：

- 当前项目路线与技术主线
- 当前 phase 单一主交付定义
- `Project Context` 的正式边界
- 根级真相源治理的正式边界
- 不做范围
- 当前阶段验收前提

需要明确写入的共享基线包括：

- PSCO 是上下文系统，不是开发流程控制器
- web 不退化，agent 受控最小进入
- agent 先消费、后维护
- 四实体只做语义澄清，不做结构重构
- 新增导出能力保持只读
- 不得引入 agent 专属一级业务对象
- 不得新增第二套 canonical API

### B. 根级入口同步

#### 4. `plan.md`

**What**

把“当前阶段”从 `phase10_asset_action_closure_foundation` 切换到 `phase11_project_context_foundation`，并新增 phase11 路线预览。

**Why**

`plan.md` 是阶段路线唯一正式承接位。若不更新，三件套会成为孤岛。

**How**

- 更新“当前状态 / 当前目标 / 当前下一阶段入口”
- 将 `PSCO-mvp05-summarize-feedback.md` 写为直接上游
- 新增 `phase11_project_context_foundation` 的路线预览
- 说明 `phase10` 已完成，`phase11` 为当前正式业务/支撑 phase（以文档最终仲裁口径为准）

#### 5. `AGENTS.md`

**What**

更新根级入口摘要，指向 `phase11` 和 `PSCO-mvp05-summarize-feedback.md`。

**Why**

`AGENTS.md` 是入口摘要。若仍停留在 `phase10` 口径，会直接误导后续接手模型。

**How**

- 更新当前阶段与当前目标摘要
- 将“当前下一阶段入口”改写为已正式进入 `phase11_project_context_foundation`
- 把最终共识引用统一到 `PSCO-mvp05-summarize-feedback.md`

#### 6. `architecture_map.md`

**What**

为 `phase11` 三件套留出明确入口，并修正最终共识引用。

**Why**

`architecture_map.md` 是目录结构与迁移落点唯一正式承接位。

**How**

- 新增 `phase11_*` 三件套入口
- 将根级最终共识引用统一到 `PSCO-mvp05-summarize-feedback.md`
- 保证新文档不是孤岛

#### 7. `docs/README.md`

**What**

把常用入口从 `phase10` 更新到 `phase11`，并修复悬空共识引用。

**Why**

`docs/README.md` 是 docs workflow 总入口，也是当前已确认存在引用漂移的文件之一。

**How**

- 更新“当前最常用入口 / 当前阶段状态”
- 增加 `phase11` 三件套入口
- 将 `[最终共识]` 从不存在的 `PSCO-summarize-feedback.md` 改为 `PSCO-mvp05-summarize-feedback.md`
- 减少与 `architecture_map.md` 的重复承载

#### 8. `docs/phase/README.md`

**What**

把 `phase11` 纳入 phase 索引，并更新“当前状态”。

**Why**

这是 phase workflow 的直接入口，必须显示新 phase 已建立 `/plan` 入口。

**How**

- 更新当前状态说明
- 新增 `phase11` 三件套入口
- 明确 `phase10` 退回“最近完成正式业务 phase 规划记录”角色

### C. 评估但默认不修改的文件

#### 9. `project_rules.md`

**What**

本轮只校验，不默认改动。

**Why**

它是项目规则基线，不应在 `/plan` 阶段为配合新 phase 轻率改写；但如果其“最终共识只以 `PSCO-summarize-feedback.md` 为准”与用户已锁定的 `PSCO-mvp05-summarize-feedback.md` 发生硬冲突，则需要在执行时慎重判断是否同步修改，或在三件套中明确记为后续根级治理对象。

**How**

- 执行时先复读
- 若修改，必须仅做与当前用户决策直接一致的最小对齐
- 若不修改，需在三件套内明确该差异如何被阶段治理承接

## Verification Steps

### 文档结构验证

1. `docs/phase/` 下新增且仅新增：
   - `phase11_project_context_foundation_architecture_plan.md`
   - `phase11_project_context_foundation_dev_plan.md`
   - `phase11_project_context_foundation_shared_baseline.md`
2. 三份文档命名、层次与 `phase10` 模式一致
3. `docs/phase/README.md`、`docs/README.md`、`architecture_map.md` 均能导航到三件套

### 内容一致性验证

1. 三件套均直接承接 `PSCO-mvp05-summarize-feedback.md`
2. 三件套都明确：
   - 上下文系统定位
   - 先消费后维护
   - web / agent 分工
   - 四实体语义澄清而非结构重构
3. `dev_plan` 至少覆盖五类子任务
4. `architecture_plan`、`dev_plan`、`shared_baseline` 之间无互相冲突的主结论

### 根级同步验证

1. `plan.md` 当前阶段已切换到 `phase11_project_context_foundation`
2. `AGENTS.md / docs/README.md / architecture_map.md / docs/phase/README.md` 都已出现 `phase11` 入口
3. 根级“最终共识”引用统一指向 `PSCO-mvp05-summarize-feedback.md`
4. `docs/README.md` 中不再保留对不存在文件 `PSCO-summarize-feedback.md` 的引用

### 边界验证

1. 三件套中未把 MCP / CLI / 写回 / 前端对话框偷渡为当前阶段主交付
2. 三件套中未把四实体结构重构写成既成事实
3. 三件套中明确以 PSCO 仓库自身为第一 dogfooding 场景
4. 三件套中明确“根级上下文真相源治理 + 最小只读项目上下文导出”为单一主交付方向
