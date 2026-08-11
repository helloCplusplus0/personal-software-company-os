# Tasks

- [x] Task 1: 对齐 `phase07-09` 的直接上游与阶段边界，明确本阶段负责"Go 后端业务传输主线切换 + L1/L2 compat 退场"，不把前端 phase07-10 工作混入当前实现。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase07_transport_contract_mainline_migration_dev_plan.md#L182-199` 的范围与 DoD。
  - [x] SubTask 1.2: 对齐 `phase07-07 formal spec` 中后端正式传输主线、compat 退场时点、错误映射与 `/api` shell 规则。
  - [x] SubTask 1.3: 对齐 `phase07-03` 对 L1/L2 候选 compat 路由"必须在 phase07-09 退场"的冻结约束。

- [x] Task 2: 设计并冻结后端 canonical Connect transport 装配方式。
  - [x] SubTask 2.1: 明确 `server.go` 继续以 `r.Route("/api", ...)` 作为唯一业务前缀壳。
  - [x] SubTask 2.2: 明确 generated `path + handler` 的消费方式（`r.Handle(path+"*", http.StripPrefix("/api", handler))`），不得写回第二套路由字符串。
  - [x] SubTask 2.3: 明确 compat 路由集中到单一 `mountCompatRoutes` 或等价承接位，不得散落在各 canonical mount 中。

- [x] Task 3: 冻结 Connect implementation 与现有 service 分层的接线边界。
  - [x] SubTask 3.1: 明确 9 个业务模块的 Connect implementation 只做 transport 解包、service 调用、response 组装与错误映射。
  - [x] SubTask 3.2: 明确继续复用既有 `QueryService / CommandService` 与 candidate/repository owner，不得在 transport 层重写业务逻辑。
  - [x] SubTask 3.3: 明确 simple 模式签名与 generated Connect handler 的正式承接位。

- [x] Task 4: 冻结 Connect interceptor 与错误语义映射的单值方案。
  - [x] SubTask 4.1: 明确 `chi` middleware 与 Connect interceptor 的职责边界。
  - [x] SubTask 4.2: 明确 domain error → Connect code 的单值映射承接位（`connecterrors.MapToConnectError`）。
  - [x] SubTask 4.3: 明确各模块不得各自维护独立错误映射表。

- [x] Task 5: 冻结 `router.go` 与 hand-written JSON handler 主线的退场策略。
  - [x] SubTask 5.1: 明确 `router.go` 从长期 JSON canonical 路由组织重构为 canonical Connect mount + compat 过渡组。
  - [x] SubTask 5.2: 明确 hand-written JSON canonical handler 必须退出主线，仅允许 phase07-10 前仍需存在的绑定 compat 薄壳暂留。
  - [x] SubTask 5.3: 明确 `chi` 只保留 shell 与非业务端点职责。

- [x] Task 6: 冻结 L1/L2 候选 compat 入口的退场实施与证据要求。
  - [x] SubTask 6.1: 明确删除 `/api/candidates/products` 与 `/api/candidates/repositories` 的路由注册。
  - [x] SubTask 6.2: 明确删除 `ListProductCandidates` / `ListRepositoryCandidates` 及对应 reader 字段。
  - [x] SubTask 6.3: 明确替代 Connect path 回归证据：`ProductRegistryService.ListProducts`、`RepositoryBindingService.ListRepositories` 返回 200。
  - [x] SubTask 6.4: 明确旧路径必须返回 404，且不得推迟到 phase07-10/11。

- [x] Task 7: 冻结 L3/L4 绑定 compat 入口在本阶段的过渡边界。
  - [x] SubTask 7.1: 明确两条绑定 compat 路由在 `phase07-09` 可作为后端薄壳保留。
  - [x] SubTask 7.2: 明确两条路由必须只存在于 compat 过渡组，并标注最晚 `phase07-10` 退场。
  - [x] SubTask 7.3: 明确本阶段不能把前端 adapter / mutation owner 回收错误并入当前 DoD。

- [x] Task 8: 冻结后端回归与验收口径。
  - [x] SubTask 8.1: 明确后端回归以 Connect procedure path 为主线，而不是继续以旧 JSON 业务路径为默认成功路径。
  - [x] SubTask 8.2: 明确 `(cd backend && go build ./...)` 通过、Connect path `curl` 验证通过、legacy 路径 404 验证通过、`/api` 单一基址验证通过。
  - [x] SubTask 8.3: 明确本阶段完成后，`phase07-11` 证据包中哪些条目应由 `phase07-09` 提前提供事实基础。

- [x] Task 9: 完成 `phase07-09` 规格一致性校验。
  - [x] SubTask 9.1: 验证本 spec 已继承 `phase07-03/04/06/07/08` 的关键冻结口径。
  - [x] SubTask 9.2: 验证本 spec 与当前 `server.go / router.go` 现状对齐，且明确指出需要切换的旧 JSON 组织方式。
  - [x] SubTask 9.3: 验证本 spec 未把前端 phase07-10 范围、adapter 删除或 L3/L4 删除提前并入本阶段。

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 1`, `Task 2` ✅
- `Task 4` depends on `Task 2`, `Task 3` ✅
- `Task 5` depends on `Task 2`, `Task 3`, `Task 4` ✅
- `Task 6` depends on `Task 1`, `Task 5` ✅
- `Task 7` depends on `Task 1`, `Task 5`, `Task 6` ✅
- `Task 8` depends on `Task 4`, `Task 5`, `Task 6`, `Task 7` ✅
- `Task 9` depends on `Task 1` through `Task 8` ✅