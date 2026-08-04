# phase01_mvp_spec_convergence_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase01` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

> 状态说明：`phase01` 已完成收口；本文档保留为规划侧共享基线。当前执行层唯一规格入口为 `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase01_mvp_spec_convergence`
- 当前技术路线：`Durable System Track`

### 2.2 正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：`v0.1` 只交付单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI；不引入独立 `React Native` 客户端，`PWA` 仅作可兼容增强方向
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- Client State：`Zustand`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go`
- Database：`PostgreSQL`
- Optional Compute：`Rust`（仅计算瓶颈时启用，不进入 `v0.1` 首轮）
- Contract：`Protocol Buffers`（跨语言长期标准）
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.3 当前项目的特别约束

- 当前项目不得再解释为 `Product Track`
- 当前项目不得把 `Hono` 写成主运行时
- 当前项目不得把 `Drizzle` 写成当前 Go 主链的既定实现依赖
- 当前项目不得把 `Rust` 写成 `v0.1` 的当前必需项

## 3. MVP 边界矩阵

### 3.1 进入 `v0.1` 主执行范围

- `Product`
- `Module`
- `Release`
- `Decision`
- `Repository`
- `Venture`（可选）

### 3.2 派生层

- `Capability`

### 3.3 后移对象

- `Feature`
- `Opportunity`
- `Experiment`

## 4. MVP 页面矩阵

- `Dashboard`
- `Module Registry`
- `Product Registry`
- `Decision Center`
- `Repository Binding`

明确不进入 `v0.1` 主范围：

- `AI Assistant` 一级导航
- `Feature` 页面
- `Opportunity` 页面
- `Experiment` 页面

## 5. MVP 动作矩阵

后续 `/spec` 至少需要承接以下动作：

- `CreateProduct`
- `CreateModule`
- `CreateRelease`
- `CreateRepository`
- `RecordDecision`
- `BindRepositoryToProduct`
- `BindModuleToProduct`
- `MapModuleToRepository`
- `LinkDecisionToTarget`

## 6. 冷启动 / 导入 / 导出矩阵

### 6.1 冷启动

- 首轮必须允许用户从零手动创建 `Product`
- 首轮必须允许用户从零手动创建 `Module`
- 首轮必须允许用户从零手动创建 `Repository`
- 首轮必须允许用户记录 `Decision`
- 首轮必须允许用户手动完成基础绑定关系

### 6.2 导入路径

- `GitHub OAuth / 自动导入` 不进入 `v0.1` 首轮
- 首轮允许的导入路径以低摩擦手动录入为主
- `/spec` 必须明确哪些资产支持导入、哪些资产仅支持手动录入

### 6.3 导出要求

- `/spec` 必须明确最小导出 / 备份要求
- 导出要求必须与 `Local First = 数据所有权优先` 的当前解释一致
- 不允许把“后续再说”作为 `v0.1` 的导出策略

## 7. 非目标矩阵

- 微服务
- Kubernetes
- Docker 全流程
- GraphQL
- Kafka
- Redis 缓存层
- Elasticsearch
- GitHub OAuth 自动导入
- Rust 计算引擎接入
- 自动扫描代码
- 自动知识图谱

## 8. 本阶段校验清单

进入 `phase01` 后续 `/spec` 前，必须再次确认：

1. `Decision` 是否仍在 MVP 中
2. `Capability` 是否仍为派生层
3. `Venture` 是否仍为可选实体
4. `Rust` 是否仍未进入 `v0.1` 首轮
5. 技术栈是否仍只以 `TECH_STACK_BASELINE.md` 为准
6. 是否已明确冷启动、导入路径与导出要求

## 9. 上游引用

- `AGENTS.md`
- `plan.md`
- `TECH_STACK_BASELINE.md`
- `project_rules.md`
- `PSCO-summarize-feedback.md`
