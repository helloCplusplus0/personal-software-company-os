# Tasks

- [x] Task 1: 冻结联调环境与启动前提。
  - [x] SubTask 1.1: 明确前端、后端、数据库与种子数据的启动顺序
  - [x] SubTask 1.2: 明确前端必须切换到真实后端模式，禁止以 mock 模式代替联调
  - [x] SubTask 1.3: 明确 `.proto`、HTTP 过渡层与当前数据库状态是本次验收的共同基线

- [x] Task 2: 验证最小主线正向路径。
  - [x] SubTask 2.1: 验证空状态进入 `CreateModule` 并成功回流到详情页
  - [x] SubTask 2.2: 验证 `CreateRelease` 成功后回流并更新版本列表
  - [x] SubTask 2.3: 验证 `BindModuleToProduct` 与 `MapModuleToRepository` 成功后停留详情页并刷新结果

- [x] Task 3: 验证关键状态路径。
  - [x] SubTask 3.1: 验证空状态文案与主动作符合 `phase02-09`
  - [x] SubTask 3.2: 验证错误停留在当前页面或面板上下文，不跳转独立错误页
  - [x] SubTask 3.3: 验证从创建页、详情页与版本登记页的返回路径规则
  - [x] SubTask 3.4: 验证返回列表时恢复 `queryText` 与 `statusFilter`

- [x] Task 4: 验证关键异常路径与边界语义。
  - [x] SubTask 4.1: 验证无效 `moduleId`、非法状态值、重复名称/重复版本/重复绑定的异常路径
  - [x] SubTask 4.2: 验证候选读取与 `Decision` 入口未扩写为新的写入主线
  - [x] SubTask 4.3: 验证系统未暴露当前阶段未冻结的新对象解释或第二套数据语义

- [x] Task 5: 核对合同与验收证据。
  - [x] SubTask 5.1: 以 `.proto` 为基准核对关键请求与响应语义
  - [x] SubTask 5.2: 明确联调中发现的问题、修复结论与剩余风险
  - [x] SubTask 5.3: 输出可复核的验收结果，证明 `Module Registry` 最小主线已形成可运行交付物

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`

# 验收证据索引

- 前端真实后端模式启动与访问记录：`VITE_USE_REAL_API=true`，前端 `:5173`，后端 `:8081`
- 后端服务、数据库迁移、种子数据与健康检查记录：`rento-preview-postgres` 容器 + `migrate.go` + `seed_readonly_prereqs.sql`
- 联调环境可重复建立入口：`init_db.sh` → `run_seeds.sh` → `reset_module_mainline.sh`（新增 `database/scripts/reset_module_mainline.sh` + `database/seeds/seed_module_mainline_baseline.sql`，提供清空+恢复基线可重复入口）
- 正向路径验证记录：`CreateModule / CreateRelease / BindModuleToProduct / MapModuleToRepository` 全部通过（201/201/204/204）
- 状态验证记录：空状态文案与入口（通过 `reset_module_mainline.sh --clean-only` 触发）、错误态停留与草稿保留、返回路径筛选恢复
- 异常路径验证记录：16 条异常路径覆盖 404 / 400 / 409（含产品绑定与仓库映射两组 404/409），无 500 级未收口错误
- `.proto` 与 HTTP 过渡层关键消息对齐记录：见 `acceptance_report.md` §5
- Decision 关联前提说明：`run_seeds.sh` 默认不建立 `decision_links`，phase02-12 验收基线通过 `reset_module_mainline.sh` 重建
- 问题收口清单与最终验收结论：见 `acceptance_report.md` §6 / §7（P-01~P-07 共 7 个问题全部收口）
