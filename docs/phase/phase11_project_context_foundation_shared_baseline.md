# phase11_project_context_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase11` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级入口中重复发散。

`phase11` 当前处于 `/plan` 阶段。本文档只承接当前阶段的单值基线、能力矩阵、治理矩阵与验收前提，不替代后续 `/spec`、实现或根级状态正文。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase11_project_context_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase10_asset_action_closure_foundation` 已完成正式收口，`phase11` 当前作为下一阶段正式 `/plan` 入口建立
- 当前阶段规划上游统一以 `PSCO-mvp05-summarize-feedback.md` 为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp05-summarize-feedback.md`
  - `project_rules.md`
  - `TECH_STACK_BASELINE.md`
  - `AGENTS.md`
  - `plan.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `docs/phase/README.md`
  - `docs/phase/phase10_asset_action_closure_foundation_*`
  - `.trae/specs/phase10_11_complete_asset_action_closure_integration_browser_regression_validation/spec.md`
- 当前阶段只承接：
  - 根级上下文真相源治理
  - 最小只读项目上下文导出
- 当前阶段不反向重写 `phase10` 已完成的正式交付

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http + ConnectRPC`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- 运行方式：`Single Server First`

### 2.4 当前阶段特别约束

- `.proto` 是唯一长期合同源
- `ConnectRPC` 是业务接口正式传输层
- `chi` 只承担 router shell、middleware 与非业务端点承接职责
- 当前阶段新增上下文导出能力必须保持只读
- 当前阶段不允许：
  - 引入 agent 专属一级业务对象
  - 引入第二套 canonical API
  - 把 PSCO 做成 IDE 现场流程控制器
  - 把 web 做成对话式 agent 工作台
  - 把 MCP / CLI / agent 写回偷渡为当前阶段主交付

## 3. 当前阶段能力矩阵

### 3.1 Project Context Foundation 单值定义

`Project Context Foundation` 在当前阶段冻结为：

- 根级上下文真相源治理
- 项目上下文聚合只读读取
- AGENTS 风格 Markdown 导出
- PSCO 仓库自身 dogfooding 验证

当前阶段不把以下内容解释为 `Project Context Foundation`：

- MCP 协议层
- CLI 工具
- agent 自动写回
- Draft / 审批流
- 主动向外部仓库注入文件
- 知识图谱或自动扫描

### 3.2 PSCO 单值定位

当前阶段继续冻结：

- PSCO 是**上下文系统**
- PSCO 不是**开发流程控制器**
- IDE / agent 负责微观执行推进
- PSCO 负责提供上下文、关系、约束、决策依据与回看入口

### 3.3 Web / Agent 分工矩阵

- web：
  - 全局查看
  - 关系校对
  - review 与回顾
  - 历史查阅
  - 人工修正
  - 最终确认

- agent：
  - 在项目现场消费当前项目上下文
  - 读取与当前项目直接相关的规则、决策与文档入口
  - 降低人工重复解释成本

补充冻结：

- web 不退化
- agent 不对称并行进入
- 二者共享同一套 Go backend canonical core

### 3.4 四实体语义矩阵

- `Product`：经营目标与交付容器
- `Repository`：代码仓库身份对象与项目锚点
- `Module`：可复用能力资产，允许后置提炼
- `Decision`：规则、约束、选择与依据的索引对象

补充冻结：

- 当前阶段只确认语义，不重写结构
- 后续 `/spec` 与实现不得把语义确认偷渡为实体重构

### 3.5 根级真相源治理矩阵

当前阶段根级治理先冻结以下单一写者规则：

- `plan.md`：阶段状态与推进路线
- `architecture_map.md`：目录结构、文档分类、迁移落点
- `TECH_STACK_BASELINE.md`：技术栈正文
- `AGENTS.md`：入口摘要
- `docs/README.md`：workflow 总入口
- `PSCO-mvp05-summarize-feedback.md`：当前最终共识

补充冻结：

- 其他根级入口文档不得重复承载以上主结论正文
- 当前阶段必须清理指向不存在文件 `PSCO-summarize-feedback.md` 的引用
- 当前阶段不做静态文件全量 backend 派生

### 3.6 项目上下文导出矩阵

当前阶段的最小项目上下文至少应覆盖：

- 当前 `Repository` 身份
- 关联 `Product` 摘要
- 关联 `Module` 摘要与状态
- 关联 `Decision` 摘要与状态
- 与当前项目直接相关的规则与约束入口
- 与当前 phase 直接相关的文档入口

导出形式至少冻结为两层：

1. 结构化只读聚合读取
2. AGENTS 风格或等价 Markdown 风格导出

## 4. 当前阶段验收前提

当前阶段验收必须至少回答：

1. 新接手 agent 是否能通过少量固定入口恢复当前项目核心上下文
2. 根级入口文档是否已消除重复承载与悬空引用
3. 最小只读导出是否完全只读
4. 导出能力是否服务于 PSCO 仓库自身真实协作场景
5. 当前阶段是否严格没有引入 agent 写入路径、第二套语义或第二套事实源

## 5. 当前阶段完成条件

`phase11` 完成时，至少必须满足：

1. `PSCO-mvp05-summarize-feedback.md` 已成为根级最终共识单值入口
2. 根级入口文档之间不再重复承载 phase 状态、目录落点与技术栈正文
3. 不存在文件 `PSCO-summarize-feedback.md` 的引用已清零
4. 已存在最小只读项目上下文能力的正式承接结果
5. 已存在 AGENTS 风格导出的正式承接结果
6. 已完成 PSCO 仓库自身 dogfooding 验证
7. 本阶段未把 MCP / CLI / agent 写回 / 对话入口偷渡为并列主交付
