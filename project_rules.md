# project_rules.md

# Personal Software Company OS Project Rules

## 1. 单一真相源规则

- 最终共识只以 `PSCO-summarize-feedback.md` 为准
- 入口摘要只由 `AGENTS.md` 承担
- 阶段路线只由 `plan.md` 承担
- 目录结构、文档分类、迁移落点只由 `architecture_map.md` 承担
- 文档总览与 workflow 入口只由 `docs/README.md` 承担
- 不允许让同一条主结论被多个根级文档重复承载

## 2. 核心基线

- 当前业务执行主链：`Product -> Module -> Release -> Decision -> Repository Binding -> Feedback`
- 当前实现目录主线：`frontend/ + backend/ + database/ + scripts/`
- 当前技术栈唯一基线文档：`TECH_STACK_BASELINE.md`
- 禁止继续沿用 `AGENTS-OLD.md` 作为技术栈来源
- `Rust Intelligence Layer` 是否进入当前项目，必须由 `TECH_STACK_BASELINE.md` 的 track 选择规则决定，而不是自由发挥
- `Local First` 的当前解释是数据所有权优先，不等于切换到 `SQLite`
- `GitHub OAuth / 自动导入` 不作为当前阶段阻断项

> 技术基线已切换：从现在开始，本项目不再沿用 `AGENTS-OLD.md` 的旧技术口径，而是统一遵守 [TECH_STACK_BASELINE.md](file:///home/dell/Projects/personal-software-company-os/TECH_STACK_BASELINE.md)。

### 2.1 全局技术栈方案

`TECH_STACK_BASELINE.md` 定义的不是“每个项目一个技术栈”，而是一套统一的长期技术方案，按场景分成两条受控路线：

- `Product Track`：`React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui + Hono + PostgreSQL + Drizzle`
- `Durable System Track`：`React + Go + PostgreSQL + Rust`

补充统一标准：

- 移动端：`React Native + Expo`
- Schema：`Zod`
- RPC / Contract：`Hono RPC / Protocol Buffers`
- 分析：`Parquet + DuckDB`
- 部署：`Caddy + systemd`
- 运行方式：`Single Server First`

### 2.2 选择规则

- 业务产品、MVP、管理系统、SaaS、AI Web 应用：默认走 `Product Track`
- 长期 7x24 运行、资源敏感、系统型工程：走 `Durable System Track`
- 只有存在明确计算瓶颈时，才在 `Durable System Track` 中增加 `Rust Engine`
- 新项目默认 `Drizzle ORM`
- 已有项目若已经稳定使用 `Prisma`，允许继续使用，不为了架构纯洁重写

### 2.3 当前项目冻结选择

- `PSCO` 当前项目已明确冻结为 `Durable System Track`
- 当前项目正式运行主线为：`React Web + Go Backend + PostgreSQL`
- 当前项目中 `Rust` 只保留为未来计算扩展位，不进入 `v0.1` 首轮实现
- 当前项目不得重新解释为 `Product Track`

### 2.4 禁止自由发挥

除非在对应 `phase / fix / audit` 文档中给出明确理由并完成审查，否则禁止擅自引入以下超出基线的技术选择：

- Kubernetes
- 微服务
- Docker 全流程
- GraphQL
- Kafka
- Redis 缓存层
- Elasticsearch
- 第二套路由、第二套状态管理、第二套 ORM、第二套 UI 框架

## 3. MVP 边界规则

- `v0.1` 的正式目标是：软件资产登记、决策留痕与基础复用反馈
- `Decision` 必须进入 MVP
- `Capability` 在 `v0.1` 中只作为派生层
- `Venture` 保留，但作为可选实体，不强制创建
- `Feature / Opportunity / Experiment` 保留在长期理论模型中，但不进入 `v0.1` 主执行范围

## 4. 默认工作流

### 4.1 phase 推进链

- 项目的功能推进、规格推进、结构性建设必须通过 `phase*` workflow 实现
- 第一步：先在 `plan.md` 做整体规划，明确当前阶段目标、范围、完成条件与下一阶段
- 第二步：当进入新的结构化阶段任务时，先执行该阶段的 `/plan`
- 第三步：`/plan` 阶段至少产出三份阶段文档，统一放在 `docs/phase/`：
  - `phase*_architecture_plan.md`
  - `phase*_dev_plan.md`
  - `phase*_shared_baseline.md`
- 第四步：阶段文档经复核通过后，再按 `dev_plan` 子任务顺序执行 `/spec`、实现、验收与收口

### 4.2 fix 推进链

- Bug 修复与小型局部问题必须通过 `fix*` workflow 推进
- 先记录 issue，再做 analysis，再进入 `/spec`、实现与验收
- `fix` 文档统一放在 `docs/fix/`

### 4.3 audit 推进链

- 跨模块复核、路线仲裁、结构性审计必须通过 `audit*` workflow 推进
- 先产出 issue 文档，再产出 analysis 文档
- `audit` 文档统一放在 `docs/audit/`
- analysis 只允许进入四类结论之一：`keep-as-is / enter-fix / enter-improvement / escalate-phase`
- 若结论为 `enter-improvement`，再进入对应 `/spec`、实现与验收

## 5. 协作规则

- 开启新 `phase` 前，先同步 `AGENTS.md`、`plan.md`、`project_rules.md`、`architecture_map.md`、`docs/README.md`
- 没有新 `phase` 时，不允许把功能推进写成散装文档或直接起一个孤立 `spec` 文件
- Bug 修复走 `fix`，跨模块复核走 `audit`，不得混写
- 任何目录迁移、文档迁移、文档新增，都必须在 `architecture_map.md` 与 `docs/README.md` 中留下明确入口
- 非归档文档禁止成为孤岛；每个活动文档都必须能从根级入口或 `docs/README.md` 找到
- 修改原始方案文档时，优先做“对齐式更新”，不要推翻重写整套叙事
- 未来路线不得写成当前版本既成事实

## 6. 非 GPT-5.4 模型协作约束

- 非 GPT-5.4 模型默认不具备 PSCO 的隐性上下文，每次接手必须先完成定向探索
- 任何非平凡改动前，必须至少读取：目标文件、直接上游文档、直接下游文档、当前阶段文档
- 遇到文档中未明确解释的设计决策，必须标注为“待确认”，不得猜测后直接修改
- 涉及 `3+` 文件的跨域改动，必须先列出影响文件清单，再逐文件执行
- 优先复用既有模式，最小修改，不得引入第二套命名体系或第二套结构模式
- 每个结构化子任务结束后，必须回到对应 baseline 或共识文档做一致性校验

## 7. 当前阶段禁止事项

- 自动扫描全部代码
- 自动生成知识图谱
- AI 自动判断最佳方案
- 将 `AI Assistant` 做成一级主导航
- 把完整 PMM / PCP 正式规范作为当前阻断项
- 在没有新的 `phase*` 入口前直接开始正式 MVP 编码
- 直接创建孤立的 `docs/specs/PSCO_MVP_Spec_v0.1.md`

## 8. 当前阶段验收规则

在进入下一阶段前，至少满足：

1. 原始方案文档已完成共识回正
2. 根级真相源文档已建立并职责去重
3. `docs/` 已按 `phase / fix / audit / archive` 收口
4. phase / fix / audit / 非 GPT-5.4 协作机制已继承并落到目录
5. 当前阶段的非目标与冻结结论已写清
