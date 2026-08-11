# Tasks

- [x] Task 1: 冻结 generated Connect handler 与既有 service implementation 的正式接线方式。把 `build* -> service -> Connect implementation -> generated handler -> chi /api` 的组合链写成单值设计，避免后续实现临时拼装。
  - [x] SubTask 1.1: 对齐 `phase07-01` 的 34 条业务 RPC 范围与 `phase07-02` 的 `/api + generated path` mount 基线 → `design.md` §1.1 接线总览 + §3.2 procedure path 映射
  - [x] SubTask 1.2: 明确 `platform` 装配层、`build*`、Connect service implementation、generated handler 各自职责 → `design.md` §1.2 各层职责表 + §1.3 文件落点树
  - [x] SubTask 1.3: 明确 compat JSON handler 在迁移期的允许位置与 canonical owner 的主次关系 → `design.md` §6 compat 过渡组

- [x] Task 2: 冻结 `chi` middleware、Connect interceptor 与错误链的唯一承接位。把 HTTP 壳层治理与 RPC 级治理拆成清晰两层，避免后续实现复制两套横切逻辑。
  - [x] SubTask 2.1: 明确 request id、real IP、logging、recovery、timeout、CORS 的唯一承接位 → `design.md` §2.1 两层横切架构图 + §2.2 职责边界表
  - [x] SubTask 2.2: 明确 Connect interceptor 承接的 metadata、校验、错误归一化、必要审计扩展 → `design.md` §2.3 Connect interceptor 配置
  - [x] SubTask 2.3: 明确 compat JSON handler 不得再长出第二套长期错误/日志治理 → `design.md` §2.4 compat handler 约束

- [x] Task 3: 冻结 procedure path、route group 与 router 结构调整方案。确保 canonical Connect tree、infra keep list 和 compat 过渡组在 `platform` 装配结构中边界清晰，且不引入第二套路由组织模式。
  - [x] SubTask 3.1: 明确 `server.go` 与 `router.go` 的职责分工 → `design.md` §3.3 职责分工表
  - [x] SubTask 3.2: 明确 `/api` 下 canonical Connect service 的 mount 组织方式 → `design.md` §1.5 装配模板 + §3.1 结构对比
  - [x] SubTask 3.3: 明确 compat JSON 入口的显式分组与 phase07-09/10 退场位置 → `design.md` §6.1 分组代码 + §6.2 分组约束

- [x] Task 4: 冻结 domain error -> proto error code -> Connect error 的单值映射方案。把当前分散在各模块 `response.go` 的 JSON status / error body 语义收敛成可直接实现的长期基线。
  - [x] SubTask 4.1: 盘点当前 domain / service 错误类别与既有 JSON handler 映射模式 → `design.md` §4.1 错误类别盘点表（覆盖 7 个模块共 30+ sentinel error）
  - [x] SubTask 4.2: 为长期错误语义冻结 Connect code 对照表 → `design.md` §4.1 错误类别 → Connect Code 对照表 + §4.2 完整映射函数代码
  - [x] SubTask 4.3: 明确固定错误映射承接位，禁止模块各自维护第二套长期映射 → `design.md` §4.3 错误映射约束表

- [x] Task 5: 冻结 Connect 迁移后的 service 分层保持策略。确保 transport 切换不触发 repository / candidate / service 分层重写，也不长出第二套 service 命名体系。
  - [x] SubTask 5.1: 明确 Connect implementation 只能负责 transport 解包、service 调用、结果装配与错误返回 → `design.md` §5.2 Connect implementation 职责边界
  - [x] SubTask 5.2: 明确不得把跨模块 SQL、candidate 构造搬进 Connect handler → `design.md` §5.2 禁止项
  - [x] SubTask 5.3: 明确既有 `service` 层仍是 canonical 业务实现主线 → `design.md` §5.1 分层不变表

- [x] Task 6: 完成与 phase07 上游冻结文档和官方工具链口径的一致性校验。确保本次设计既承接既有基线，也能直接进入 `phase07-08` 实现。
  - [x] SubTask 6.1: 校验与 `phase07-02` 的 single `/api`、generated path、3 插件生成链一致 → `design.md` §7 一致性声明
  - [x] SubTask 6.2: 校验与 `phase07-03` 的 compat 分组与退场窗口一致 → `design.md` §6 compat 分组 + §7 一致性声明
  - [x] SubTask 6.3: 校验与 Context7 的 `connect-go / chi / buf` 最新口径一致 → `design.md` §1.5 使用 `New*ServiceHandler` + `r.Handle(path, handler)` 模式（对齐 Context7 API）

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 1` and `Task 2` ✅
- `Task 4` depends on `Task 2` ✅
- `Task 5` depends on `Task 1` ✅
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5` ✅
