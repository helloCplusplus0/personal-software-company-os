# Phase01-01 技术路线与系统边界冻结 Spec

## Why

`phase01` 的第一步不是实现代码，而是先把 PSCO 当前项目的技术路线与系统边界冻结成单值结论。只有先明确 `Durable System Track`、当前正式运行主线、排除项与后续扩展位，后面的对象、动作、页面、数据与 API 规格才不会反复漂移。

## What Changes

- 将 `PSCO` 当前项目正式冻结为 `Durable System Track`
- 冻结当前正式运行主线为 `React Web + Go Backend + PostgreSQL`
- 冻结前端统一基线为 `React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`
- 冻结后端主线为 `Go + PostgreSQL + 模块化单体 + 单进程单主运行面`
- 冻结跨语言长期合同标准为 `Protocol Buffers`
- 冻结部署与运行标准为 `Single Server First + Caddy + systemd`
- 明确 `Rust` 在 `v0.1` 中不进入首轮实现，仅保留未来计算扩展位
- 明确当前项目不得重新解释为 `Product Track`
- 明确当前项目不得继续沿用 `AGENTS-OLD.md` 作为技术栈来源
- 明确当前项目不得把 `Drizzle` 写成当前 Go 主链的既定实现依赖

## Impact

- Affected specs: `phase01_mvp_spec_convergence`
- Affected code: 当前无代码改动，影响后续 `backend/`、`frontend/`、`database/`、部署与跨语言合同设计

## ADDED Requirements

### Requirement: 当前项目技术路线冻结
系统 SHALL 将 `PSCO` 当前项目冻结为 `Durable System Track`，并以此作为 `phase01` 后续规格与实现的唯一技术路线。

#### Scenario: 当前项目路线确定
- **WHEN** 接手者读取当前项目的正式规格
- **THEN** 必须得到 `PSCO` 当前项目走 `Durable System Track` 的单值结论
- **AND** 不得再将当前项目重新解释为 `Product Track`

### Requirement: 当前正式运行主线冻结
系统 SHALL 将当前项目的正式运行主线定义为 `React Web + Go Backend + PostgreSQL`。

#### Scenario: 运行主线判定
- **WHEN** 接手者查询当前项目的运行主线
- **THEN** 必须得到 `React Web + Go Backend + PostgreSQL`
- **AND** 不得把 `Hono` 写成当前项目主运行时

### Requirement: 前端统一基线冻结
系统 SHALL 规定当前项目前端统一遵守 `React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`。

#### Scenario: 前端技术选择
- **WHEN** 后续规格涉及前端框架、路由、数据请求、状态管理和 UI 基线
- **THEN** 必须只使用上述统一基线
- **AND** 不得引入第二套路由、第二套状态管理或第二套 UI 框架

### Requirement: 后端与部署边界冻结
系统 SHALL 规定当前项目后端主线为 `Go + PostgreSQL + 模块化单体 + 单进程单主运行面`，部署与运行标准为 `Single Server First + Caddy + systemd`。

#### Scenario: 后端与部署约束
- **WHEN** 后续规格涉及后端结构或部署方式
- **THEN** 必须以模块化单体和单进程单主运行面为默认前提
- **AND** 不得默认引入微服务、Kubernetes、Docker 全流程、Kafka、Redis 缓存层或 Elasticsearch

### Requirement: 跨语言合同长期标准冻结
系统 SHALL 将 `Protocol Buffers` 作为 TS / Go / Rust 跨语言合同的长期标准。

#### Scenario: 跨语言合同决策
- **WHEN** 后续规格涉及 TS / Go / Rust 的跨语言边界
- **THEN** 必须与 `Protocol Buffers` 方向保持一致
- **AND** 当前阶段可以暂不实现完整 `proto` 工具链

### Requirement: Rust 扩展位边界冻结
系统 SHALL 明确 `Rust` 只作为未来高性能计算扩展位，在 `v0.1` 首轮实现中不作为必需项。

#### Scenario: Rust 进入条件
- **WHEN** 后续规格或实现讨论 `Rust`
- **THEN** 必须将其视为未来计算扩展位
- **AND** 不得把 `Rust` 写成 `v0.1` 当前必需项

## MODIFIED Requirements

### Requirement: 技术栈来源解释
当前项目的技术栈来源 SHALL 只以 `TECH_STACK_BASELINE.md` 为准；`AGENTS-OLD.md` 不再承担当前项目技术栈来源职责。

#### Scenario: 技术栈来源判定
- **WHEN** 接手者查询当前项目技术栈来源
- **THEN** 只能引用 `TECH_STACK_BASELINE.md`
- **AND** 不得使用 `AGENTS-OLD.md` 作为当前项目的技术栈解释依据

### Requirement: Go 主链数据访问依赖约束
系统 SHALL 明确当前项目走 `Durable System Track` 时，Go 主链不得把 `Drizzle` 作为既定实现依赖；`Drizzle` 仅属于 `Product Track` 的 TS 生态 ORM 基线。

#### Scenario: Go 主链 ORM 选择
- **WHEN** 后续规格或实现讨论当前 Go 主链的数据访问/ORM 选型
- **THEN** 不得把 `Drizzle` 写成当前 Go 主链的既定实现依赖
- **AND** 需要对 Go 主链的数据访问层给出独立选型，而不是直接复用 `Product Track` 的 `Drizzle` 基线
