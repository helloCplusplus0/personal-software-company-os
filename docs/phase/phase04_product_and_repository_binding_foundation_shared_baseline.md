# phase04_product_and_repository_binding_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase04` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase04_product_and_repository_binding_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase04` 已完成收口，项目当前入口已切换为 `phase05_dashboard_feedback_foundation`
- 本文档保留为 `phase04` 的共享冻结基线；文中“当前阶段”均指 `phase04` 当时上下文，不覆盖根级当前阶段状态

### 2.2 当前阶段唯一执行层上游

- 直接执行层上游：
  - `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
  - `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
  - `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
  - `.trae/specs/phase03_11_decision_center_proto_mainline/`
  - `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`
- 当前阶段只承接 `v0.1` 已冻结边界
- 当前阶段不反向重写 `phase02` 已冻结的 `Module Registry` 边界
- 当前阶段不反向重写 `phase03` 已冻结的 `Decision Center` 边界

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
- 当前阶段不得把 `Module Detail` 长期保留为第二个绑定工作台
- 当前阶段不得把 `Decision Center` 重新扩写为 `Product / Repository` 绑定主线
- 当前阶段不得把 GitHub OAuth / 自动导入写成当前阶段阻断项

### 2.5 当前阶段交付模式

- `phase04` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 `Product Registry + Repository Binding` 最小主线代码

## 3. 当前阶段动作矩阵

`phase04` 最少需要直接承接：

- `CreateProduct`
- `CreateRepository`
- `BindRepositoryToProduct`
- `BindModuleToProduct`
- `MapModuleToRepository`

当前阶段必须打通的最小关系闭环：

- `Product -> Repository`
- `Product -> Module`
- `Module -> Repository`

允许以最小连接位承接但不扩写为独立主线：

- `Product -> Decision`
- `Repository -> Decision`
- `Module Detail` 中的绑定摘要与跳转入口

## 4. 当前阶段页面矩阵

- `Product Registry / List`
- `Product Create`
- `Product Detail`
- `Repository Binding / List`
- `Repository Create`
- `Repository Binding Detail / Workspace`

允许存在最小入口或上下文跳转：

- `Module Detail`

### 4.1 当前阶段交互归属矩阵

- `Product Registry / List`：承接产品列表读取、筛选入口、创建入口与进入详情入口
- `Product Create`：承接 `CreateProduct`
- `Product Detail`：承接产品详情读取、当前已绑定模块/仓库读取与进入绑定流程的上下文入口
- `Repository Binding / List`：承接仓库列表读取、筛选入口、创建入口与进入绑定工作台入口
- `Repository Create`：承接 `CreateRepository`
- `Repository Binding Detail / Workspace`：承接仓库详情读取、候选读取、`BindRepositoryToProduct`、`MapModuleToRepository`
- `Module Detail`：当前阶段只承接绑定摘要与跳转入口，不扩写为独立绑定工作台

## 5. 当前阶段数据矩阵

直接承接：

- `products`
- `repositories`
- `product_modules`
- `product_repositories`
- `module_repositories`

当前阶段必须直接读取或校验：

- `modules`

允许保留轻量连接前提，但不要求当前阶段写入主线：

- `decisions`

### 5.1 最小读写模型

- 产品列表读取至少承接：`name / description / status / created_at / module_bind_count / repository_bind_count`
- 产品详情读取至少承接：核心对象字段、已绑定模块列表、已绑定仓库列表与最小上下文入口
- 仓库列表读取至少承接：`name / url / provider / status / created_at / product_bind_count / module_bind_count`
- 仓库详情读取至少承接：核心对象字段、已绑定产品列表、已映射模块列表与绑定工作台上下文
- 创建写入承接：`CreateProduct`、`CreateRepository`
- 绑定写入承接：
  - `BindModuleToProduct` → `Product Detail`
  - `BindRepositoryToProduct` → `Repository Binding Detail / Workspace`
  - `MapModuleToRepository` → `Repository Binding Detail / Workspace`

当前阶段候选读取基线冻结如下：

- `BindRepositoryToProduct` 的候选为当前已存在的 `products`
- `BindModuleToProduct` 的候选为当前已存在的 `modules`
- `MapModuleToRepository` 的候选为当前已存在的 `modules`
- 候选排序、状态过滤、已绑定排除规则必须在后续 `/spec` 中单值化
- 无可绑定候选时，页面必须返回明确空状态，而不是把空结果误报为接口错误

当前阶段 `Product / Repository` 最小字段冻结为：

- `Product`：`name / description / status`
- `Repository`：`name / url / provider / status`

当前阶段 `Product / Repository` 状态语义冻结如下：

- 当前阶段最小持久化枚举统一为：
  - `active`
  - `archived`
- `status` 是 `Product` 与 `Repository` 持久化记录中的必有字段
- `CreateProduct` 与 `CreateRepository` 的默认创建状态统一为 `active`
- 当前阶段创建写入继续按“显式提交 `status`”处理：默认 `active` 的含义是创建页、HTTP DTO 与 `.proto` 写模型在未发生用户改动时都应预填并显式提交 `active`，而不是依赖服务端或合同层隐式补默认值
- 列表筛选中的 `statusFilter` 仅作为 UI/路由层枚举，当前阶段统一为：
  - `all`
  - `active`
  - `archived`
- `all` 不得写入数据库、合同持久化字段或后端领域模型

字段级冻结如下：

- 创建必填：
  - `Product`：`name / description / status`
  - `Repository`：`name / url / provider / status`
- 上述“创建必填”是指当前阶段的 Create 页面输入模型、HTTP DTO 与 `.proto` 写请求都必须携带该字段；不允许把“持久化必有字段”误解释为“创建请求可省略，由服务端默默补值”
- 空字符串不得视为合法必填值；写入前必须完成去首尾空白后的最小非空校验
- 当前阶段不引入自动导入来源字段、扫描状态字段或复杂同步字段

### 5.2 最小接口归属前提

- `CreateProduct` 在 `phase04` 中按 `Product Create` 的直接写入动作处理
- `CreateRepository` 在 `phase04` 中按 `Repository Create` 的直接写入动作处理
- `BindModuleToProduct` 在 `phase04` 中按 `Product Detail` 的直接写入动作处理
- `BindRepositoryToProduct` 与 `MapModuleToRepository` 在 `phase04` 中按 `Repository Binding Detail / Workspace` 的直接写入动作处理
- `Product Detail` 可以提供进入仓库绑定工作台的上下文入口，但自身不承接第二套仓库绑定写入流程
- `Module Detail` 只允许保留进入正式主入口的上下文跳转，不扩写为第二套主写入流程
- 当前阶段允许保留 `chi + JSON HTTP` 作为过渡传输层，但不得形成与 `.proto` 并列的第二套合同源

最小错误语义前提：

- `CreateProduct` 与 `CreateRepository` 缺少必填字段时，返回明确校验失败
- 任一绑定动作的目标不存在时，返回资源不存在语义
- 任一绑定动作重复建立时，返回重复冲突语义
- 候选读取返回空结果时，返回空列表语义，不得错误映射为资源不存在

### 5.3 当前阶段源码设计层基线

- 前端必须明确 `Product / Repository / Binding` 页面分层、最小路由结构与组件职责
- 前端必须明确列表、详情、创建、绑定工作台的状态模型
- 前端必须明确 `Product List` 与 `Repository List` 的筛选维度、路由搜索参数承接方式与返回列表恢复规则
- 后端必须明确 `Product Registry` 与 `Repository Binding` 模块边界与读写接口分组
- 当前阶段必须为 `Product / Repository / Binding` 落地最小 `.proto` 合同源
- 当前阶段必须提前定义联调重置脚本、基线种子、兼容迁移与验收 fixture
- 当前阶段不提前冻结 Go 数据访问层具体工具
- 当前阶段必须明确 `phase02` 临时绑定承接点如何迁移并保持兼容
- 当前阶段必须明确迁移完成后的 canonical owner、旧入口保留级别、兼容跳转参数与成功写入后的 reread 承接页面

## 6. 当前阶段合同与演进基线

- `.proto` 是 `Product / Repository / Binding` 的唯一合同源
- `buf` 校验链至少覆盖：`build`、`lint`、`generate`、`breaking`
- `.proto` 字段语义必须与正式规格正文、HTTP DTO 与前端消费模型保持单值一致
- 合同演进必须遵守兼容性约束；删除字段后必须保留 `reserved` 字段号，必要时同时保留字段名
- `breaking` 校验必须直接对照仓库主线基准，不允许吞掉失败退出码

## 7. 当前阶段冷启动与验收基线

- 首轮必须允许用户从空状态进入 `Product Create`
- 首轮必须允许用户从空状态进入 `Repository Create`
- 首轮必须允许用户完成首个 `Product -> Repository` 绑定
- 首轮必须允许用户将已存在 `Module` 绑定到 `Product`
- 首轮必须允许用户将已存在 `Module` 映射到 `Repository`
- 首轮必须允许用户从 `Module Detail` 的兼容入口跳转进入正式绑定主入口，并完成至少一条绑定写入
- 三类绑定成功后必须回到对应 canonical owner 页面完成 reread，不得只靠 toast 作为成功依据
- 当前阶段验收不得依赖手工补 SQL 才能建立最小联调环境
- 当前阶段必须提供可重复执行的重置脚本、基线种子与异常路径验证前提

## 8. 非目标矩阵

- Dashboard 聚合反馈
- `Decision -> Product / Repository` 正式关联写入
- `Feature / Opportunity / Experiment`
- GitHub OAuth / 自动导入
- 自动扫描代码
- 自动知识图谱
- AI 自动建议或自动绑定
- 独立 `React Native` 客户端
- 完整 `PWA`

## 9. 本阶段校验清单

进入 `phase04` 后续 `/spec` 前，必须再次确认：

1. 当前阶段直接执行层上游是否仍为 `phase01 + phase02 + phase03` 已冻结规格与验收结果
2. `Product Registry` 与 `Repository Binding` 是否仍然是当前阶段唯一主交付对象
3. 是否仍未重新引入后移对象
4. 是否仍采用单一 `React Web` 前端交付策略
5. 是否已明确五个核心动作与三类绑定关系的交互归属
6. 是否已明确前端/后端源码设计层最小输出要求
7. 是否已明确 `.proto` 合同与 `buf` 校验链从阶段一开始进入主线
8. 是否已明确联调环境、重置脚本、兼容迁移与基线数据前提
9. 是否已明确 `phase04` 最终以代码交付而不是文档冻结作为完成条件

## 10. 上游引用

- `AGENTS.md`
- `plan.md`
- `TECH_STACK_BASELINE.md`
- `project_rules.md`
- `architecture_map.md`
- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
- `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
- `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`
