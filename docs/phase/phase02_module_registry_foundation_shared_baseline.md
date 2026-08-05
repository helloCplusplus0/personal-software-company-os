# phase02_module_registry_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase02` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase02_module_registry_foundation`
- 当前技术路线：`Durable System Track`

### 2.2 当前阶段唯一执行层上游

- 唯一执行层上游：`.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- 当前阶段只承接 `v0.1` 已冻结边界

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- Client State：`Zustand`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`（长期方向）
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段不得重新引入 `Feature / Opportunity / Experiment`
- 当前阶段不得引入独立 `AI Assistant` 一级导航
- 当前阶段不得引入独立 `React Native` 客户端
- 当前阶段不得把完整 `PWA` 能力写成前置范围
- 当前阶段不得重新解释 `Module Registry` 之外的独立实现主线

### 2.5 当前阶段交付模式

- `phase02` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 `Module Registry` 最小主线代码

## 3. 当前阶段动作矩阵

`phase02` 最少需要直接承接：

- `CreateModule`
- `CreateRelease`
- `BindModuleToProduct`
- `MapModuleToRepository`

允许以最小入口承接但不扩写为独立主线：

- `LinkDecisionToTarget`

## 4. 当前阶段页面矩阵

- `Module Registry / List`
- `Module Create`
- `Module Detail`
- `Release Create`

允许存在最小跳转或关联入口：

- `Product Registry`
- `Repository Binding`
- `Decision Center`

## 4.1 当前阶段交互归属矩阵

- `Module Registry / List`：承接列表读取、筛选入口、创建入口与进入详情入口
- `Module Create`：承接 `CreateModule`
- `Module Detail`：承接详情读取、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository`
- `Decision Center`：当前阶段只承接跳转或只读关联入口，不在 `phase02` 中扩写为独立写入主线

## 5. 当前阶段数据矩阵

直接承接：

- `modules`
- `module_releases`
- `product_modules`
- `module_repositories`

最小读取或关联前提：

- `decisions`
- `decision_links`

候选读取前提（只读，不要求写入主线，由 `phase02-08` 收口）：

- `products`
- `repositories`

### 5.1 最小读写模型

- 列表读取至少承接：`name / description / status / latest_release / product_bind_count / repository_bind_count`
- 详情读取至少承接：核心对象字段、版本列表、产品绑定、仓库映射与相关 `Decision` 入口
- 创建写入承接：`CreateModule`
- 版本写入承接：`CreateRelease`
- 关联写入承接：`BindModuleToProduct`、`MapModuleToRepository`

### 5.2 最小接口归属前提

- `BindModuleToProduct` 与 `MapModuleToRepository` 在 `phase02` 中按 `Module Detail` 的直接写入动作处理
- `LinkDecisionToTarget` 在 `phase02` 中只作为只读展示或跳转入口，不扩写为当前阶段独立写入主线

## 5.3 当前阶段源码设计层基线

- 前端必须明确页面分层、最小路由结构与组件职责
- 前端必须明确列表、详情、创建、版本登记的状态模型
- 后端必须明确 `Module Registry` 模块边界与读写接口分组
- 当前阶段不提前冻结 Go 数据访问层具体工具

## 6. 当前阶段冷启动矩阵

- 首轮必须允许用户从空状态进入 `CreateModule`
- 首轮必须允许用户完成首个 `Release` 登记
- 首轮允许模块先登记后补充 `Product / Repository` 关联
- 当前阶段不依赖导入、自动扫描或 AI 推荐

## 7. 非目标矩阵

- Product 全量主线
- Decision Center 全量主线
- Repository Binding 全量主线
- Dashboard 聚合反馈
- 自动扫描代码
- 自动知识图谱
- 独立 `React Native` 客户端
- 完整 `PWA`

## 8. 本阶段校验清单

进入 `phase02` 后续 `/spec` 前，必须再次确认：

1. 当前阶段唯一执行层上游是否仍是 `mvp_spec_v0.1.md`
2. `Module Registry` 是否仍然是当前阶段唯一主交付对象
3. 是否仍未重新引入后移对象
4. 是否仍采用单一 `React Web` 前端交付策略
5. 是否已明确模块登记、版本登记与最小关联入口
6. 是否已明确绑定动作与 `Decision` 入口的交互归属
7. 是否已明确前端/后端源码设计层最小输出要求
8. 是否已明确 `phase02` 最终以代码交付而不是文档冻结作为完成条件

## 9. 上游引用

- `AGENTS.md`
- `plan.md`
- `TECH_STACK_BASELINE.md`
- `project_rules.md`
- `architecture_map.md`
- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
