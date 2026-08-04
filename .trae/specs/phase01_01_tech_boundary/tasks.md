# Tasks

- [ ] Task 1: 冻结当前项目技术路线。将 `PSCO` 当前项目的正式路线写成单值结论，明确为 `Durable System Track`，并消除任何 `Product Track` 混写空间。
  - [ ] SubTask 1.1: 明确当前项目正式运行主线为 `React Web + Go Backend + PostgreSQL`
  - [ ] SubTask 1.2: 明确当前项目不得重新解释为 `Product Track`
  - [ ] SubTask 1.3: 明确当前项目不得继续沿用 `AGENTS-OLD.md` 作为技术栈来源

- [ ] Task 2: 冻结前端、后端、部署与跨语言合同边界。将当前项目后续实现必须遵守的系统边界写成可执行约束。
  - [ ] SubTask 2.1: 冻结前端统一基线为 `React + Vite + TypeScript + TanStack Router + TanStack Query + Zustand + Tailwind CSS + shadcn/ui`
  - [ ] SubTask 2.2: 冻结后端主线为 `Go + PostgreSQL + 模块化单体 + 单进程单主运行面`
  - [ ] SubTask 2.3: 冻结部署与运行标准为 `Single Server First + Caddy + systemd`
  - [ ] SubTask 2.4: 冻结跨语言长期合同标准为 `Protocol Buffers`

- [ ] Task 3: 冻结排除边界与未来扩展位。明确当前阶段哪些技术不进入 `v0.1` 首轮实现，防止 `/spec` 与后续实现扩范围。
  - [ ] SubTask 3.1: 明确 `Rust` 只保留未来计算扩展位，不进入 `v0.1` 首轮实现
  - [ ] SubTask 3.2: 明确当前阶段不默认引入微服务、Kubernetes、Docker 全流程、Kafka、Redis 缓存层与 Elasticsearch
  - [ ] SubTask 3.3: 验证上述排除边界与 `TECH_STACK_BASELINE.md` 和 `phase01` 规划文档保持一致

- [ ] Task 4: 完成规格校验。检查本次 `phase01-01` 规格是否具备进入后续子任务的条件。
  - [ ] SubTask 4.1: 验证技术路线已经单值化
  - [ ] SubTask 4.2: 验证无 `Product Track` / `Durable System Track` 混写
  - [ ] SubTask 4.3: 验证无 `AGENTS-OLD.md` 旧技术栈残留解释

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
