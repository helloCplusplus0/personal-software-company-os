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

补充冻结：

- `PSCO` 自身仓库治理与“面向未来不同项目的通用上下文能力”是同一 phase 内的两层交付，但不是同一层合同；
- `PSCO` 当前仓库中存在的根级文件清单，只用于自身治理与第一轮 dogfooding，不外推为所有未来项目的固定目录模板。

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

- `Module` 当前代表可复用能力资产，允许在后续真实复用沉淀中继续提炼，当前阶段不要求重写其 schema、层级或注册主线
- `Decision` 当前代表规则、约束、选择与依据的索引对象，用于支撑项目上下文恢复与只读导出，不在本阶段扩写为审批流、流程引擎或结构重构入口
- 当前阶段只确认语义，不重写结构
- 后续 `/spec` 与实现不得把语义确认偷渡为实体重构

### 3.5 根级真相源治理矩阵

当前阶段根级治理先冻结以下单一写者规则：

- `plan.md`：阶段状态与推进路线
- `architecture_map.md`：目录结构、文档分类、迁移落点
- `TECH_STACK_BASELINE.md`：技术栈正文
- `README.md`：项目总览入口与受控跳转
- `AGENTS.md`：入口摘要
- `project_rules.md`：项目级协作规则与单一真相源约束
- `global_skills.md`：项目内通用方法映射说明
- `docs/README.md`：workflow 总入口
- `docs/phase/README.md`：phase 文档入口索引
- `PSCO-mvp05-summarize-feedback.md`：当前阶段 `PSCO` 自身仓库的有效最终共识文档

补充冻结：

- 当前阶段根级治理的正式设计产物至少应包括：
  - 根级入口治理矩阵
  - 重复承载清单与目标落点清单
  - 悬空引用清理清单
  - 收口后的单一写者规则表
- 其他根级入口文档不得重复承载以上主结论正文
- 当前阶段必须清理治理矩阵目标文件范围内指向不存在文件 `PSCO-summarize-feedback.md` 的引用
- 当前阶段不做静态文件全量 backend 派生
- 当前阶段的治理路线冻结为：一次性校准，而不是静态文件全量派生
- 当前阶段有效最终共识入口的具体文件名只代表 `PSCO` 当前时点的有效文档，不上升为未来版本推进或其他项目的固定文件合同

### 3.6 项目上下文导出矩阵

当前阶段的最小项目上下文至少应覆盖：

- 唯一结构化输入锚点：`repository_id`
- 当前 `Repository` 身份
- 关联 `Product` 摘要
- 关联 `Module` 摘要与状态
- 关联 `Decision` 摘要与状态
- 与当前项目直接相关的规则与约束入口
- 与当前 phase 直接相关的文档入口

补充冻结：

- 当前阶段不以本地路径、Git remote URL、`product_id` 或工作区扫描作为并列主锚点
- 当前阶段只承接“已完成 Repository Binding”的仓库上下文读取；当前阶段将“绑定完成”明确解释为：目标 `Repository` 至少已有一条 `product_repositories` 绑定，且至少已有一条 `module_repositories` 映射
- 仓库不存在与仓库绑定不完整都必须返回明确失败态，不允许执行者自行补猜
- `Decision` 聚合口径冻结为：以当前 `Repository` 为根，只合并基于既有 `Decision -> Module` canonical link 可直接投影出的两类命中：
  - 命中“当前 `Repository` 已映射 `Module`”的 `Decision`
  - 命中“当前 `Repository` 已绑定 `Product` 所属 `Module`”的 `Decision`
- 当前阶段不得把 `Repository` 或 `Product` 伪装成 `Decision` 的直接 link target，也不得继续沿 `Product -> Module -> 其他 Repository` 做递归扩张；超出上述两类命中范围的 `Decision` 不进入当前阶段导出
- 同一 `Decision` 若同时命中多类关系，必须以 `decision_id` 去重，并保留命中来源摘要
- 当前阶段结构化只读主列表只承接非 `archived` 的 `Decision`；`archived` 不进入主导出列表
- 结构化只读读取继续落在 Go backend 的 `.proto + ConnectRPC` 正式主线
- AGENTS 风格 Markdown 导出必须从同一结构化只读结果单向派生
- 当前阶段不把消费侧项目目录中的 `README.md / AGENTS.md / rules` 等固定文件名作为必要输入合同
- 当前阶段的通用能力以 `PSCO` 中已登记的 `Repository / Product / Module / Decision` 关系为准
- 若未来存在最佳实践项目模板，其身份只能是候选 convention/profile，不是当前阶段前置依赖

导出形式至少冻结为两层：

1. 结构化只读聚合读取
2. AGENTS 风格或等价 Markdown 风格导出

补充冻结的字段边界：

- 结构化只读输出字段边界至少承接：
  - `repository_id` 与正式失败语义；
  - 当前 `Repository` 身份；
  - 关联 `Product / Module / Decision` 的最小摘要与状态；
  - `Decision` 命中来源摘要；
  - 规则、约束与文档入口字段（至少包含入口定位值与定位类型）。
- Markdown 导出字段边界至少承接：
  - 当前项目/仓库摘要；
  - 当前 phase 相关 spec / baseline / 根级入口摘要；
  - `Product / Module / Decision` 的最小可读摘要；
  - 规则、约束与文档入口的受控引用。
- Markdown 导出只允许从同一结构化只读结果单向派生，不允许额外长出第二套字段语义或第二套事实源。

## 4. 当前阶段验收前提

当前阶段验收必须至少回答：

1. 新接手 agent 是否能在同一 dogfooding 剧本下，通过不超过 `3` 个固定入口恢复当前项目核心上下文
2. 同一 dogfooding 剧本中，是否能够准确回答预设的 `5` 个恢复问题：当前 phase、直接上游、单一主交付、明确不做、当前项目关联的 Repository / Product / Module / Decision 摘要入口
3. 根级入口文档是否已消除重复承载与悬空引用
4. 最小只读导出是否完全只读，且唯一结构化输入锚点是否保持为 `repository_id`
5. 导出能力是否服务于 PSCO 仓库自身真实协作场景
6. 当前阶段是否严格没有引入 agent 写入路径、第二套语义或第二套事实源
7. 当前阶段是否已经明确：PSCO 当前仓库文件清单只用于自身治理与 dogfooding，而不是未来所有项目的强制模板

补充冻结的 dogfooding 入口集合：

- 旧路径基线入口集合固定为：
  - `AGENTS.md`
  - `plan.md`
  - `project_rules.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `PSCO-mvp05-summarize-feedback.md`
- 旧路径基线执行时，不允许额外读取 `phase11` 新增的结构化导出结果或 AGENTS 风格项目上下文导出
- 新路径固定入口集合冻结为：
  - `AGENTS.md`
  - `PSCO-mvp05-summarize-feedback.md`
  - 基于同一 `repository_id` 生成的 AGENTS 风格项目上下文导出
- 新路径不允许临时增加第 `4` 个入口来补齐答案

## 5. 当前阶段完成条件

`phase11` 完成时，至少必须满足：

1. `PSCO-mvp05-summarize-feedback.md` 已成为根级最终共识单值入口
2. 根级入口文档之间不再重复承载 phase 状态、目录落点与技术栈正文
3. 不存在文件 `PSCO-summarize-feedback.md` 的引用已清零
4. `repository_id` 已成为当前阶段唯一正式结构化输入锚点，未绑定仓库失败态已冻结
5. 已存在最小只读项目上下文能力的正式承接结果
6. 已存在 AGENTS 风格导出的正式承接结果
7. 已按固定 dogfooding 剧本完成 PSCO 仓库自身验证，并满足“固定入口 <= 3、预设 5 问全部可回答”的验收标准
8. 已明确“最佳实践项目模板”当前仅属未来候选增强，不是 phase11 前置条件
9. 本阶段未把 MCP / CLI / agent 写回 / 对话入口偷渡为并列主交付
