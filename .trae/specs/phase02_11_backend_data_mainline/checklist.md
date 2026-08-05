- [x] 已明确 `phase02-11` 的直接上游规格入口是 `phase02-09` 的 `module_registry_spec_v0.1.md`
- [x] 已明确后端实现必须继续遵守 `phase02-08` 冻结的 `handler / service / repository / candidate` 包结构与文件落点
- [x] 已明确 `ModuleListRead / ModuleDetailRead / ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite` 是当前阶段必须实现的最小后端接口组
- [x] 已明确 `Decision` 只作为 `ModuleDetailRead` 的内嵌附属读取承接，不设独立读接口组
- [x] 已明确数据库主线至少覆盖 `modules / module_releases / product_modules / module_repositories`
- [x] 已明确 `decisions / decision_links / products / repositories` 只作为当前阶段读取前提，不扩写为新的写入主线
- [x] 已明确本地开发环境复用现有 Podman PostgreSQL 容器，而不是为 PSCO 新开第二个数据库容器
- [x] 已明确 PSCO 在共享 PostgreSQL 实例中使用独立数据库，而不是直接复用 `rento_production`
- [x] 已明确本地 `DATABASE_URL` 必须为显式最终值，且特殊字符密码需要 URL 编码
- [x] 已明确数据库迁移主线必须可重复执行，并支撑 `phase02-12` 联调验收
- [x] 已明确必须提供 `products / repositories` 的最小候选数据与 `decisions / decision_links` 的最小示例数据或 fixture
- [x] 已明确后端运行入口必须能够接上数据库连接并让接口可运行
- [x] 已明确 `phase02-10` 前端临时适配层后续必须能切换到真实后端，而不引入第二套数据语义
- [x] 已明确 `phase02-11` 的验收以"后端读写接口可运行、数据主线与冻结边界一致、不引入新对象解释或第二套数据主线"为准

# 额外校验项（实现过程中补充）

- [x] 后端 `go build ./...` 与 `go vet ./...` 通过
- [x] 前端 `npm run build` 与 `npm run lint` 在 mock 与 real API 两种模式下均通过（仅 shadcn/ui 预存警告）
- [x] 迁移系统幂等可重复执行（第二次启动显示 "no new migrations to apply"）
- [x] 非 UUID 路径参数返回 404 而非 500（service 层 UUID 格式校验）
- [x] 唯一约束冲突映射为 409 而非 500（DB 层 23505 错误码检测）
- [x] 浏览器端到端验证：列表 → 详情（UUID 路径）→ 版本/绑定/Decision 展示
- [x] `.env` 文件已 gitignore，敏感凭据不写入仓库源码
- [x] 前端切换不改变任何函数签名，页面与组件代码无需修改
- [x] PSCO 独立数据库初始化入口 `database/scripts/init_db.sh` 可重复执行（幂等建库）
- [x] 种子数据统一执行入口 `database/scripts/run_seeds.sh` 可重复执行（幂等，支持 fixture 可选）
- [x] 后端 `RUN_SEEDS_ON_BOOT` 配置项可控制启动时是否自动执行种子（默认 false 保证生产安全）
- [x] 脚本与后端均不硬编码凭据，密码通过环境变量或容器内 peer 认证获取
