# phase09_template_reuse_derived_intelligence_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase09` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

`phase09` 当前处于 `/plan` 阶段。本文档只承接当前阶段的单值基线、范围矩阵、动作矩阵与验收前提，不替代后续 `/spec`、实现或根级状态正文。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase09_template_reuse_derived_intelligence_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase08` 已完成正式收口，`phase09` 已作为后续支撑能力 phase 进入 `/plan`
- 当前阶段规划上游统一以 `PSCO-mvp03-summarize-feedback.md`、`phase08-11` 验收结论与 `phase06` 已交付复用摘要主线为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp03-summarize-feedback.md`
  - `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`
  - `docs/phase/phase08_operating_review_loop_foundation_*`
  - `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
  - `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_shared_baseline.md`
  - `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
  - `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
  - `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
- 当前阶段只承接 `phase08` 已冻结并验收的 review loop、`phase06` 已冻结并落地的 `reuse_summary`、以及 `phase04` 已冻结的 `Product Create` canonical 路径
- 当前阶段不反向重做 `phase08` 的 review 会话边界，不把 `dry-run` 混入当前实现承诺

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- Contract Tooling：`buf build / lint / generate / breaking`
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段新增前端写路径必须继续遵守：
  - 业务写路径唯一 `application` 入口
  - `query` 层纯只读
  - mutation 固定承接位
  - 切片优先与 `shared` 延迟晋升
- 当前阶段新增后端接口必须继续遵守：
  - `.proto` 是唯一长期合同源
  - `ConnectRPC` 是新增业务接口默认正式传输层
  - HTTP DTO 只能从 `.proto` 单向派生或显式映射
- 当前阶段不得重新引入 `Feature / Opportunity / Experiment`
- 当前阶段不得把 `Capability` 扩写为独立重实体
- 当前阶段不得把模板级复用扩写为完整模板平台
- 当前阶段不得把派生智能扩写为独立 AI 工作台

### 2.5 当前阶段交付模式

- `phase09` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 `Template Reuse + Derived Intelligence` 最小支撑能力代码

## 3. 当前阶段能力矩阵

### 3.1 Template Reuse 单值定义

`Template Reuse` 在当前阶段冻结为：

- `Module` 组合快照
- 面向 `Product Create` 的预填辅助
- 预填后继续编辑并完成创建

当前阶段不把以下内容解释为 `Template Reuse`：

- 独立模板实体主线
- 模板平台
- 模板版本管理
- 参数化模板系统

### 3.2 Derived Intelligence Deepening 单值定义

`Derived Intelligence Deepening` 在当前阶段冻结为：

- `capability gap` 最小提示
- `reuse opportunity` 最小提示
- 与 `review / create` 直接相连的解释性指标与行动文案

当前阶段不把以下内容解释为 `Derived Intelligence Deepening`：

- AI 主线
- 独立智能中心
- `Decision Intelligence` 独立主线
- 自动生成长期战略或任务系统

### 3.3 当前阶段必须直接承接的最小闭环

当前阶段最少需要直接承接：

- `Weekly Review -> 查看模板候选 / 派生提示`
- `选择模板候选 -> Product Create 预填`
- `Product Create 继续编辑 -> 创建成功 -> Product Detail`
- `结果回流 -> Review / Dashboard / ReuseSummary reread`

允许以最小连接位承接但不扩写为独立主线：

- `Decision` 对“下一步做什么”的语义锚点
- Dashboard 对支撑能力的轻量摘要显示
- 后续 `dry-run` 的前置进入条件

## 4. 当前阶段页面矩阵

- `Weekly Review`
- `Product Create`
- `Product Detail`
- `Dashboard Home`

### 4.1 当前阶段交互归属矩阵

- `Weekly Review`：承接模板候选与派生提示的主要诊断入口
- `Product Create`：承接模板预填、用户确认与正式创建
- `Product Detail`：承接创建成功回流与结果核验
- `Dashboard Home`：承接轻量摘要与结果 reread，不扩写为第二套模板创建宿主

补充冻结：

- `Daily Review` 当前不承担复杂模板编排主入口
- 模板选择若来自 `Review`，正式落点必须仍然是 `Product Create`
- 页面不得各自拼装第二套 `Product Create` mutation、失效刷新与回流语义

### 4.2 `Product Create` 模板来源参数矩阵

当前阶段 `/products/new` 的正式来源参数矩阵冻结如下：

- 既有来源参数：
  - `fromList`
  - `fromModuleDetail`
  - `fromDashboard / dashboardSection / dashboardReturnTo`
- 新增模板来源参数：
  - `fromTemplateReuse`
  - `templateCandidateId`
  - `templateSource`

当前阶段优先级与互斥规则冻结如下：

1. `fromTemplateReuse=true` 时，模板来源是本次 create 的唯一业务来源语义
2. `fromTemplateReuse` 与 `fromList / fromModuleDetail` 互斥，不允许并列成立
3. `fromDashboard / dashboardSection / dashboardReturnTo` 只承担返回链与来源链元数据，可与 `fromTemplateReuse` 共存
4. `Product Create` 页面对模板预填的唯一读取入口是 `templateCandidateId` 对应的正式 read owner；不得在 search 中直接携带完整模板 payload

## 5. 当前阶段数据矩阵

直接承接：

- `module_reuse_summary`
- `capability_summary`
- `Weekly Review` 上下文
- `Product Create` 既有写入主线
- `Product / Module / Binding / Decision` 既有 canonical 事实

当前阶段必须正式消费或新增的支撑读取：

- `reuse template candidates`
- `module composition snapshot`
- `capability gap hints`
- `reuse opportunity hints`
- `explanatory insight items`

### 5.1 模板候选单值规则

- canonical 来源：`product_modules` 已持久化绑定事实；`Review` 只承担消费作用域与返回链上下文，不直接生成模板候选
- 原始候选输入：单个 Product 当前绑定的全部 Module 集合
- 去重键：去重并升序排序后的 `module_id` 集合
- 候选 ID：由后端根据去重键单向派生，前端按 opaque string 消费
- 排序规则：
  1. `source_product_count DESC`
  2. `total_reuse_product_count DESC`
  3. `latest_source_product_updated_at DESC`
- 空态规则：
  - 没有候选时返回成功空态
  - 不得把“无候选”解释成页面错误
- 低质量候选规则：
  - 空 Module 集合不产生候选
  - 只出现一次且没有任何复用事实支撑的组合，不进入默认推荐列表，但可在后续 `/spec` 中评估是否保留为非推荐候选

当前阶段关于 `Weekly Review` 模板候选的选择语义继续冻结如下：

- 候选列表采用 **单选模型**
- 若存在候选，则按既定排序结果取第一名作为默认 active candidate
- 用户可在当前 review 会话中切换 active candidate，但同一时刻只允许一个 active candidate 生效
- `Weekly Review` 中依赖模板上下文的提示、解释文案与 CTA，一律只基于当前 active candidate 计算
- 若没有任何候选，则与模板直接相关的提示返回成功空态；当前阶段不再保留“退回 generic review focus”这一并列触发口径

### 5.2 最小读写模型

- `reuse template candidate` 至少应承接：
  - 候选来源标识
  - 快照名称或解释文案
  - 关联 `Module` 组合摘要
  - 推荐原因或适用语义
- `module composition snapshot` 至少应承接：
  - 候选 `Module` 列表
  - 与当前目标场景相关的最小说明
  - 进入 `Product Create` 所需的正式预填上下文
- `capability gap hint` 至少应承接：
  - 能力键或等价能力语义
  - 缺口解释
  - 建议下一步动作
- `reuse opportunity hint` 至少应承接：
  - 可复用对象或组合
  - 推荐原因
  - 关联正式动作入口

当前阶段关于模板预填成功会话的单值定义如下：

- 推荐执行顺序：`Weekly Review -> 模板候选 -> Product Create -> Product Detail`
- 成功会话成立条件：
  - 用户能从正式模板候选进入 `Product Create`
  - 创建页能看到可继续编辑的预填内容
  - 预填内容至少覆盖：模板名称/说明、候选 `Module` 组合摘要、待创建 Product 的初始模块上下文
  - 用户可在预填基础上修改并成功创建
  - 创建成功后进入 `Product Detail` 时，仍能通过模板来源上下文读取到模板名称/说明与候选 `Module` 组合摘要
  - `Product Detail` 必须提供进入既有 canonical `Product <-> Module Binding` 路径的下一步动作

补充冻结：

- 当前阶段 `Product Create` 提交成功时，不在同一 mutation 内自动写入 `product_modules`
- 模板组合在 `phase09` 中通过“来源摘要 + 下一步 binding CTA”继续承接，而不是通过新增长期事实表或第二套写路径承接

### 5.3 派生提示矩阵

当前阶段正式提示只允许包含：

1. `reuse_opportunity_hint`
   - trigger：存在当前推荐模板候选
   - explanation：说明该组合为何值得复用
   - CTA：进入模板预填创建路径
   - target owner：`TemplateReuseRead` + `ProductCreate` canonical path

2. `capability_gap_hint`
   - trigger：当前存在 active template candidate，且当前 review 作用域的高频 capability 未被该模板候选覆盖
   - explanation：说明缺口来自哪里、为何会阻碍下一次创造
   - CTA：进入既有 `Module Registry` / `Module Detail` canonical 路径继续补齐
   - target owner：`DerivedInsightRead` + 既有 `Module Registry` canonical path

补充约束：

- 每类提示都必须同时具有 `trigger / explanation / CTA / target owner`
- 不能满足该矩阵的提示，不进入 `phase09` 正式范围
- `Product Create` 中允许展示解释性提示，但不得在 create 页内联第二套写路径
- 若当前没有 active template candidate，则 `capability_gap_hint` 返回成功空态，而不是切换到未冻结的 generic review focus 口径

### 5.4 最小接口归属前提

- `TemplateReuseRead` 由当前阶段新增或扩展 query owner 承接
- `DerivedInsightRead` 由当前阶段新增或扩展 query owner 承接
- `Product Create` 仍由既有 `ProductCreate` mutation owner 承接
- `Review` 页面只消费正式 read / application owner，不得直接拼装底层 client 调用
- 当前阶段默认采用读时派生模板候选与预填详情读取；不新增持久化模板快照表
- 若 `phase09-04` 证明必须引入轻量快照记录，也不得将其解释为新的 canonical 事实源

## 6. 当前阶段合同与演进基线

- `.proto` 是当前阶段新增接口的唯一合同源
- `buf` 校验链至少覆盖：`build`、`lint`、`generate`、`breaking`
- `.proto` 字段语义必须与前端消费模型保持单值一致
- 合同演进必须遵守兼容性约束
- 当前阶段默认不新增持久化模板快照表；只有 `phase09-04` 拿到明确不足证据后，才允许升级为受控轻量快照记录，且其身份只能是支撑能力资产

当前阶段关于模板与提示的最小覆盖矩阵如下：

- 模板候选最小必须覆盖：
  - 候选来源
  - `Module` 组合摘要
  - 进入 `Product Create` 的正式 handoff 上下文
- 派生提示最小必须覆盖：
  - `capability gap`
  - `reuse opportunity`
  - 解释性文案
  - 指向正式动作的 CTA 或等价 handoff

## 7. 当前阶段验收基线

- 用户必须能从 `Weekly Review` 或等价正式消费位看到模板候选与派生提示
- 用户必须能从模板候选进入 `Product Create` 并完成可编辑预填创建
- 当前阶段必须允许创建成功后回流到 `Product Detail` 并验证 reread
- 当前阶段必须允许模板与提示在无数据时以成功空态返回，而不是视为页面错误
- 当前阶段验收不得依赖新增长期核心实体才能建立联调环境

补充约束：

- 验收时不得把“只是跳转到创建页”算作模板预填闭环成立
- 验收时不得把“只是多展示几条统计摘要”算作派生智能已成立
- 验收时必须至少覆盖以下浏览器闭环：
  1. `/dashboard -> /reviews/weekly`
  2. `Weekly Review` 显示模板候选或成功空态、显示派生提示或成功空态
  3. 从模板候选 CTA 进入 `/products/new?fromTemplateReuse=true&templateCandidateId=...`
  4. `Product Create` 显示可编辑预填内容并成功提交
  5. 创建成功进入 `Product Detail`，确认模板来源摘要、候选 `Module` 组合摘要与 canonical binding CTA 可见，随后完成 `Dashboard / Review / ReuseSummary` 最小 reread 验证
- API smoke 至少必须覆盖：
  - 模板候选读取
  - 模板预填详情读取
  - 派生提示读取
- 验收时必须分别验证：
  - 模板预填真实降低了从空白创建开始的摩擦
  - 提示真实支撑了下一步动作判断
  - `Product Create` 的 canonical write owner 没有被模板逻辑替换或分叉
  - `Product Detail` 已保留模板来源语义，并把后续动作继续导向既有 canonical `Product <-> Module Binding` 路径

当前阶段对 reread 结果的最小观察断言继续冻结如下：

- `Product Detail` 必须成功读取到新建 Product 本身，并显示与 `templateCandidateId` 对应的模板来源摘要 / 候选 `Module` 组合摘要
- `Product Detail` 必须出现单值的“继续绑定模板中的 Module”正式 CTA；若模板读取失败，也必须回退到可解释失败态，而不是丢失整个详情页
- `Dashboard / Review / ReuseSummary` 的 reread 当前只要求成功返回并保持来源链可继续消费；若本次创建尚未完成模块绑定，不要求这些聚合页立即出现新的复用统计变化

## 8. 非目标矩阵

- `Real-Project Dry-Run`
- `Venture / Decision Intelligence / AI Context Enhancement`
- 完整模板平台
- 参数化模板版本体系
- 独立智能工作台
- 自动扫描 / 知识图谱
