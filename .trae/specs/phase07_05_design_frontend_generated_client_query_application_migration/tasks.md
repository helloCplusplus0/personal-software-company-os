# Tasks

- [x] Task 1: 冻结前端 generated client 的正式组合方式与物理落点。把 `@connectrpc/connect + @connectrpc/connect-web + bufbuild/es` 与当前 `/api` 代理链、切片结构对齐，避免后续前端实现临时生长第二套客户端组织。
  - [x] SubTask 1.1: 明确共享 Connect transport 的唯一落点与职责边界 → `design.md` §1.2（`shared/rpc/connect-transport.ts`，唯一 cross-slice transport）
  - [x] SubTask 1.2: 明确各业务切片 `data/connect-client.ts` 或等价 slice-local generated client 承接位 → `design.md` §1.3（9 个切片 client 落点表）
  - [x] SubTask 1.3: 明确 `package.json`、`frontend/src/gen/proto/`、`vite.config.ts` 与 `/api` baseUrl 的协同关系 → `design.md` §1.1 架构图 + §1.5 禁止项

- [x] Task 2: 冻结 query 层迁移策略。把当前 page / component / route 级读取 caller 与少量既有 read owner 收敛到统一的 slice-local read owner 模式，确保 `query` 继续纯只读。
  - [x] SubTask 2.1: 盘点当前真实 read owner、page-level `useQuery`、component 级 candidate reads 与 route 级 first-run caller → `design.md` §2.1（按真实 caller 盘点表）
  - [x] SubTask 2.2: 明确 generated client 接入后每个切片的 read owner、candidate read owner 与 route 可复用只读 helper 承接位 → `design.md` §2.2（read owner 落点树）+ §2.3（Onboarding route / CTA 模板）
  - [x] SubTask 2.3: 明确 queryKey、enabled、响应解包、route helper 复用与页面/组件消费边界，禁止 query 层混入写动作 → `design.md` §2.3 read owner 模板 + §2.4 query 约束表

- [x] Task 3: 冻结 mutation owner inventory 与收口策略。把当前 7 处页面/组件级 mutation 逐项分类为 application owner、固定承接位或允许的短时过渡位。
  - [x] SubTask 3.1: 对齐 `phase07-01` 已冻结的 11 个 canonical 写动作清单 → `design.md` §3.1（11 项完整清单，含当前落点、迁移后 owner、owner 类型）
  - [x] SubTask 3.2: 明确 `CreateRelease / Bind / Map / Link / Export / Backup` 各自的最终 owner 与最晚退场时点 → `design.md` §3.2（5 个新增 application owner 文件）+ §3.3（2 个短时过渡位规则）
  - [x] SubTask 3.3: 明确成功回流、query 失效与错误归一化由谁承接 → `design.md` §3.4 application owner 模板（含 onSuccess 失效刷新 + 跨切片失效）

- [x] Task 4: 冻结旧 adapter 回收顺序与 compat 收口规则。把 `api-adapter.ts`、`*-adapter.ts` 和 `module-registry` runtime switch 的退场顺序写成可直接执行的迁移链。
  - [x] SubTask 4.1: 盘点 8 份 hand-written `api-adapter.ts` 与 4 份 adapter 壳的当前角色 → `design.md` §4.1（13 项 adapter 盘点表，含当前角色与回收优先级）
  - [x] SubTask 4.2: 明确各切片 page / component / route caller 切换完成前后的删除顺序 → `design.md` §4.2 三层回收顺序 + §4.3 各切片回收明细表
  - [x] SubTask 4.3: 明确 `module-registry/data/module-registry-adapter.ts` 与 compat dormant export 的最终收口规则 → `design.md` §4.4（5 步收口规则）

- [x] Task 5: 完成与上游 phase07 文档、phase06 mutation owner 基线和 Context7 官方口径的一致性校验。确保本次设计能直接指导 `phase07-07` 与 `phase07-10`。
  - [x] SubTask 5.1: 校验与 `phase07-02` 的 `generated client + /api + 单一 transport` 结论一致 → `design.md` §6 一致性声明
  - [x] SubTask 5.2: 校验与 `phase07-03` 的 compat 退场窗口、`phase07-04` 的 Connect procedure path 结论一致 → `design.md` §4.4 + §6 一致性声明
  - [x] SubTask 5.3: 校验与 `phase06-07` 的 application owner 基线、Context7 的 `connect-es / connect-web / TanStack Query` 最新口径一致 → `design.md` §2.3 + §3.5 + §6 一致性声明

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 1` and `Task 2` ✅
- `Task 4` depends on `Task 2` and `Task 3` ✅
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4` ✅
