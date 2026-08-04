# phase01_mvp_spec_convergence_architecture_plan

## 1. 文档定位

本文档是 `phase01_mvp_spec_convergence` 的架构规划文档。

目标不是直接进入实现，而是冻结本阶段的架构边界、技术路线、范围边界与输出物，为后续 `dev_plan`、`shared_baseline` 与 `/spec` 提供唯一上游。

> 状态说明：`phase01` 已完成收口；本文档保留为规划侧上游。当前执行层唯一规格入口为 `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`，根级状态以 `AGENTS.md` 与 `plan.md` 为准。

## 2. 上游输入

本阶段唯一上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `PSCO-summarize-feedback.md`
7. 用户明确决策：**PSCO 走 Durable System Track**

## 3. 本阶段目标

`phase01` 的目标是：

> 在 `Durable System Track` 前提下，完成 PSCO v0.1 MVP 规格收敛，形成后续实现可直接遵循的单值规格基线。

本阶段需要回答的核心问题：

1. PSCO 在 `Durable System Track` 下的正式架构是什么
2. `v0.1` 保留哪些对象、动作、页面与边界
3. 哪些技术标准在当前项目中立即生效，哪些只冻结为后续路线
4. 冷启动、导入路径与导出要求如何在 `v0.1` 中落成最小闭环
5. 如何在不扩大范围的前提下，把 MVP 规格压缩为可执行版本

## 4. 架构冻结结论

### 4.1 项目技术路线

PSCO 当前项目正式冻结为：

`Durable System Track`

这意味着当前项目的正式运行主线为：

`React Web + Go Backend + PostgreSQL`

并附带以下约束：

- `Rust` 只作为高性能计算引擎的保留扩展位，不进入 `phase01` 或 `v0.1` 首轮实现范围
- 不允许再为本项目重新解释为 `Product Track`
- 不允许把 `Hono` 再写成当前项目主运行时

### 4.2 前端基线

尽管项目走 `Durable System Track`，前端基线仍统一遵守全局技术方案：

- `React`
- `Vite`
- `TypeScript`
- `TanStack Router`
- `TanStack Query`
- `Zustand`
- `Tailwind CSS`
- `shadcn/ui`

### 4.3 后端基线

后端主线冻结为：

- `Go`
- `PostgreSQL`
- 模块化单体优先
- 单进程、单主运行面优先

当前阶段只冻结方向，不在本阶段额外引入超出基线的 Go 技术分叉。

### 4.4 合同与跨语言边界

本项目遵守 `Contract First`。

在 `Durable System Track` 下：

- TS / Go / Rust 跨语言合同的长期标准冻结为 `Protocol Buffers`
- 当前 `phase01` 不要求立刻实现完整 `proto` 生成链
- 但后续规格与实现不得采用与该方向冲突的第二套跨语言合同路线

### 4.5 部署与运行约束

当前默认部署与运行方式冻结为：

- `Single Server First`
- `Caddy + systemd`
- `PostgreSQL`

当前阶段不引入：

- Kubernetes
- 微服务
- Docker 全流程
- Kafka
- Redis 缓存层
- Elasticsearch

## 5. MVP 范围冻结

`phase01` 中，MVP 继续遵守既有最终共识：

### 5.1 核心实体

- `Product`
- `Module`
- `Release`
- `Decision`
- `Repository`
- `Venture`（可选）

### 5.2 派生层 / 后移层

- `Capability`：派生层，不作为重实体
- `Feature / Opportunity / Experiment`：后移，不进入 `v0.1` 主执行范围

### 5.3 页面级主范围

- `Dashboard`
- `Module Registry`
- `Product Registry`
- `Decision Center`
- `Repository Binding`

### 5.4 最小落地约束

- 必须定义首轮冷启动路径
- 必须定义导入路径与手动录入边界
- 必须定义最小导出 / 备份要求
- 必须把 `Repository Binding` 细化到可执行的绑定动作，而不是停留在泛化命名

## 6. 本阶段输出物

本阶段必须产出：

1. `phase01_mvp_spec_convergence_architecture_plan.md`
2. `phase01_mvp_spec_convergence_dev_plan.md`
3. `phase01_mvp_spec_convergence_shared_baseline.md`

本阶段已完成审核，并已进入 `phase01` 对应 `/spec` 收口。

## 7. 本阶段不做

本阶段明确不做：

- 代码实现
- 前后端脚手架落地
- Rust 引擎接入
- 完整 proto 工具链实施
- GitHub OAuth 自动导入
- AI Assistant 一级主导航
- 对 `v0.1` 范围做二次扩大

## 8. 通过标准

当以下条件满足时，`phase01` 的架构规划才算通过：

1. `Durable System Track` 已写成单值结论
2. 技术栈来源已完全指向 `TECH_STACK_BASELINE.md`
3. MVP 范围、后移范围、非目标范围已无冲突
4. 后续 `dev_plan` 与 `shared_baseline` 可以在此基础上继续展开
