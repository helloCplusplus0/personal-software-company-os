# phase03_decision_center_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase03` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase03_decision_center_foundation`
- 当前技术路线：`Durable System Track`

### 2.2 当前阶段唯一执行层上游

- 直接执行层上游：
  - `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
  - `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
  - `.trae/specs/phase02_11a_module_registry_proto_contract/`
  - `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`
- 当前阶段只承接 `v0.1` 已冻结边界
- 当前阶段不反向重写 `phase02` 已冻结的 `Module Registry` 边界

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- Client State：`Zustand`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- Contract Tooling：`buf build / lint / generate / breaking`
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段不得重新引入 `Feature / Opportunity / Experiment`
- 当前阶段不得引入独立 `AI Assistant` 一级导航
- 当前阶段不得引入独立 `React Native` 客户端
- 当前阶段不得把完整 `PWA` 能力写成前置范围
- 当前阶段不得把 `Module Detail` 扩写为第二个 `Decision` 工作台
- 当前阶段不得把 `Product / Repository` 扩写为并列主交付对象

### 2.5 当前阶段交付模式

- `phase03` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 `Decision Center` 最小主线代码

## 3. 当前阶段动作矩阵

`phase03` 最少需要直接承接：

- `RecordDecision`
- `LinkDecisionToTarget`

当前阶段必须打通的最小目标关系：

- `Decision -> Module`

允许以最小连接位承接但不扩写为独立主线：

- `Decision -> Product`
- `Decision -> Repository`

## 4. 当前阶段页面矩阵

- `Decision Center / List`
- `Decision Create`
- `Decision Detail`

允许存在最小入口或上下文跳转：

- `Module Detail`

### 4.1 当前阶段交互归属矩阵

- `Decision Center / List`：承接决策列表读取、筛选入口、创建入口与进入详情入口
- `Decision Create`：承接 `RecordDecision`
- `Decision Detail`：承接详情读取、目标关联展示、`Decision -> Module` 候选读取与 `LinkDecisionToTarget`
- `Module Detail`：当前阶段只承接跳转入口或带上下文的预填入口，不扩写为独立决策工作台

## 5. 当前阶段数据矩阵

直接承接：

- `decisions`
- `decision_links`

当前阶段必须直接读取或校验：

- `modules`

允许保留轻量候选读取前提，但不要求当前阶段写入主线：

- `products`
- `repositories`

### 5.1 最小读写模型

- 列表读取至少承接：`title / status / created_at / link_count / linked_module_summary`
- 详情读取至少承接：核心对象字段、结构化模板字段、已关联目标列表与最小来源上下文
- `Decision Detail` 的目标候选读取作为附属读取承接，当前阶段只面向 `Module`
- 创建写入承接：`RecordDecision`
- 关联写入承接：`LinkDecisionToTarget`

当前阶段列表读模型的计算口径冻结如下：

- `link_count` 当前阶段仅统计 `decision_links` 中已建立的 `Decision -> Module` 有效关联数
- `linked_module_summary` 仅基于已关联 `Module` 生成，不混入 `Product / Repository`
- `linked_module_summary` 按 `module_name` 升序取前 `3` 个名称；若超出 `3` 个，则在摘要末尾附加 `+N`
- 当 `Decision` 当前没有任何已关联 `Module` 时，`link_count` 返回 `0`，`linked_module_summary` 返回空字符串，不返回 `null`

当前阶段目标候选读取基线冻结如下：

- 候选来源为当前已存在的 `modules`
- 候选范围同时覆盖 `active` 与 `archived` 的 `Module`，避免历史决策无法关联历史模块
- 候选排序采用 `status(active 优先) -> module_name 升序`
- 已建立 `Decision -> Module` 关联的目标不得再次出现在可关联候选中
- 无可关联候选时，页面必须返回明确空状态，而不是把空结果误报为接口错误

当前阶段 `Decision` 结构化模板最小字段冻结为：

- `title`
- `context`
- `problem`
- `alternatives`
- `choice`
- `reason`
- `impact`
- `status`

字段级冻结如下：

- 创建必填：`title / context / problem / choice / reason / status`
- 创建可选：`alternatives / impact`
- `alternatives` 冻结为按顺序保留的文本条目集合；当前阶段不引入嵌套对象结构
- 空字符串不得视为合法必填值；写入前必须完成去首尾空白后的最小非空校验

当前阶段 `status` 最小枚举冻结为：

- `proposed`
- `active`
- `superseded`
- `archived`

当前阶段 `decisions` 表演进兼容基线冻结如下：

- `phase02` 中仅承接只读入口的 `decisions` 表在 `phase03` 中必须原位升级为结构化主线，不得并行新建替代表后再临时双写
- 现有示例 `Decision` 数据必须通过迁移脚本完成兼容回填，满足当前阶段最小结构化模板的非空约束
- 兼容回填必须保留原有 `title / created_at`，并保证既有 `decision_links` 在迁移后仍可正常读取
- 当前阶段不得依赖手工 SQL 临时修补历史样例，迁移与 seed 必须可重复执行

### 5.2 最小接口归属前提

- `RecordDecision` 在 `phase03` 中按 `Decision Create` 的直接写入动作处理
- `Decision -> Module` 的候选读取在 `phase03` 中按 `Decision Detail` 的附属读取处理，不扩写为独立主交付动作
- `LinkDecisionToTarget` 在 `phase03` 中按 `Decision Detail` 的直接写入动作处理
- `Decision -> Module` 为当前阶段必交付目标关联
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层，但不得形成与 `.proto` 并列的第二套合同源

最小错误语义前提：

- `RecordDecision` 缺少必填字段时，返回明确校验失败，不得降级为模糊通用错误
- `Decision -> Module` 候选读取返回空结果时，返回空列表语义，不得错误映射为资源不存在
- `LinkDecisionToTarget` 目标不存在时，返回资源不存在语义
- `LinkDecisionToTarget` 重复关联时，返回重复冲突语义

### 5.3 当前阶段源码设计层基线

- 前端必须明确页面分层、最小路由结构与组件职责
- 前端必须明确列表、详情、创建、目标关联的状态模型
- 前端必须明确 `Decision List` 的筛选维度、路由搜索参数承接方式与返回列表恢复规则
- 后端必须明确 `Decision Center` 模块边界与读写接口分组
- 当前阶段必须为 `Decision Center` 落地最小 `.proto` 合同源
- 当前阶段必须提前定义联调重置脚本、基线种子与验收 fixture
- 当前阶段不提前冻结 Go 数据访问层具体工具
- 当前阶段必须明确旧 `Decision` 示例数据如何通过迁移与 seed 保持兼容

当前阶段 `Decision List` 最小上下文恢复基线：

- 查询条件冻结到路由搜索参数层：`queryText`、`statusFilter`
- 从 `DecisionCreatePage` 或 `DecisionDetailPage` 返回时，必须按原有 `queryText` 与 `statusFilter` 恢复列表上下文
- 刷新页面后，若路由搜索参数仍在，列表必须按该参数恢复读取状态

## 6. 当前阶段合同与演进基线

- `.proto` 是 `Decision Center` 的唯一合同源
- `buf` 校验链至少覆盖：`build`、`lint`、`generate`、`breaking`
- `.proto` 字段语义必须与正式规格正文、HTTP DTO 与前端消费模型保持单值一致，允许通过显式映射承接，不允许出现并列定义
- 合同演进必须遵守兼容性约束；删除字段后必须保留 `reserved` 字段号，必要时同时保留字段名，避免未来复用
- `breaking` 校验必须直接对照仓库主线基准，不允许吞掉失败退出码

## 7. 当前阶段冷启动与验收基线

- 首轮必须允许用户从空状态进入 `Decision Create`
- 首轮必须允许用户记录首条 `Decision`
- 首轮必须允许用户将该 `Decision` 关联到已存在 `Module`
- 当前阶段验收不得依赖手工补 SQL 才能建立最小联调环境
- 当前阶段必须提供可重复执行的重置脚本、基线种子与异常路径验证前提

## 8. 非目标矩阵

- Product 全量主线
- Repository 全量主线
- Dashboard 聚合反馈
- `pending_decision_signals` 的完整消费链
- 自动扫描代码
- 自动知识图谱
- AI 自动建议或自动裁决
- 独立 `React Native` 客户端
- 完整 `PWA`

## 9. 本阶段校验清单

进入 `phase03` 后续 `/spec` 前，必须再次确认：

1. 当前阶段直接执行层上游是否仍为 `phase01 + phase02` 已冻结规格与验收结果
2. `Decision Center` 是否仍然是当前阶段唯一主交付对象
3. 是否仍未重新引入后移对象
4. 是否仍采用单一 `React Web` 前端交付策略
5. 是否已明确 `RecordDecision`、`LinkDecisionToTarget` 与 `Decision -> Module` 的交互归属
6. 是否已明确前端/后端源码设计层最小输出要求
7. 是否已明确 `.proto` 合同与 `buf` 校验链从阶段一开始进入主线
8. 是否已明确联调环境、重置脚本与基线数据前提
9. 是否已明确 `phase03` 最终以代码交付而不是文档冻结作为完成条件

## 10. 上游引用

- `AGENTS.md`
- `plan.md`
- `TECH_STACK_BASELINE.md`
- `project_rules.md`
- `architecture_map.md`
- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
- `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`
