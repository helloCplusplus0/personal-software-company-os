# Tasks

- [x] Task 1: 冻结 `phase07-11` 的统一验收输入边界与证据来源，确保当前阶段只承接联调、回归、退场核销与正式证据收口。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase07_transport_contract_mainline_migration_dev_plan.md#L234-251` 的范围与 DoD。
  - [x] SubTask 1.2: 对齐 `phase07-07 formal spec` 与 `phase07-06 design.md` 中的 `34` 条 RPC、`4` 条 legacy endpoint、`11` 项 mutation owner 与 `CR1-CR9` 回归矩阵。
  - [x] SubTask 1.3: 继承 `phase02-12`、`phase03-14`、`phase04-14`、`phase05-14`、`phase06-16` 的既有验收入口，但明确它们在本阶段只是上游证据来源，不再是并列直接验收入口。

- [x] Task 2: 冻结 `phase07-11` 的联调环境、启动顺序与工具链验证口径。
  - [x] SubTask 2.1: 明确数据库重置脚本、后端启动、前端启动与 `/api` 路由检查的统一顺序。
  - [x] SubTask 2.2: 明确 `proto / backend / frontend` 的本地命令链，以及它们在当前阶段作为 CI 等价验证的使用方式。
  - [x] SubTask 2.3: 明确开发、验收与部署等价环境都不得引入第二套 API 基址、第二套路由主线或旧 hand-written JSON 业务入口。

- [x] Task 3: 冻结 9 个业务模块的模块内回归矩阵与 route 级回归矩阵。
  - [x] SubTask 3.1: 为 Module / Decision / Product / Repository 的列表、详情、创建、绑定/关联动作定义统一回归项。
  - [x] SubTask 3.2: 为 Dashboard / Onboarding / Export / Backup / Reuse Summary 的读取、分流、快照与动作定义统一回归项。
  - [x] SubTask 3.3: 为 `/`、`/onboarding`、各模块 list/detail/new 路径定义 route 级验证与返回路径验证。

- [x] Task 4: 冻结跨模块联动与 reread 验收矩阵。
  - [x] SubTask 4.1: 继承 `CR1 ~ CR9` 的联动路径，明确前置 fixture、触发步骤与成功回流检查。
  - [x] SubTask 4.2: 明确 mutation 成功后对应 read owner 的 reread / invalidate 证据要求。
  - [x] SubTask 4.3: 明确 Dashboard 计数、Onboarding CTA、Reuse Summary 新鲜度等跨模块派生结果的验收口径。

- [x] Task 5: 冻结 legacy / compat endpoint retirement inventory 的核销步骤。
  - [x] SubTask 5.1: 为 `L1 ~ L4` 分别定义路由删除、handler/adapter 删除与旧路径 `404` 验证要求。
  - [x] SubTask 5.2: 明确替代 Connect path 的成功验证方式，避免只验旧路径 `404` 却不验新主线。
  - [x] SubTask 5.3: 明确发现任何未声明残留时的阻断规则与收口动作。

- [x] Task 6: 冻结 frontend mutation owner inventory 的核销步骤。
  - [x] SubTask 6.1: 逐项列出 `11` 项 mutation owner 的触发页面/组件、前置 fixture 与成功回流检查。
  - [x] SubTask 6.2: 明确页面 / 组件级 `useMutation` 的静态核销规则，区分正式 owner 与允许过渡位。
  - [x] SubTask 6.3: 明确 `SovereigntyPanel` 中 `ExportCoreAssets / CreateInstanceBackup` 作为显式允许过渡位的验收边界。

- [x] Task 7: 冻结 `phase07-11` 的正式验收记录与证据包结构。
  - [x] SubTask 7.1: 明确单一正式验收记录必须包含环境、步骤、结果、问题、复测与 DoD 判定。
  - [x] SubTask 7.2: 明确 `canonical RPC / legacy endpoint / mutation owner / adapter retirement / cross-module regression / toolchain` 六类证据的最小结构。
  - [x] SubTask 7.3: 明确 `phase07-11` 通过后才能进入 `phase07-12` 根级收口，并为 `mvp0.3` 提供可直接承接的正式结论。

- [x] Task 8: 完成 `phase07-11` 规格一致性校验。
  - [x] SubTask 8.1: 验证本规格与 `phase07-06/07/08/09/10` 的冻结口径一致。
  - [x] SubTask 8.2: 验证本规格未把"CI 已落地"写成既成事实，而是正确表述为"CI 等价验证"。
  - [x] SubTask 8.3: 验证本规格未提前混入 `phase07-12` 的根级同步正文，只保留进入下一阶段的条件。

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 1`, `Task 2` ✅
- `Task 4` depends on `Task 1`, `Task 3` ✅
- `Task 5` depends on `Task 1`, `Task 2` ✅
- `Task 6` depends on `Task 1`, `Task 3` ✅
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6` ✅
- `Task 8` depends on `Task 1` through `Task 7` ✅

# 实现总结

## 变更文件清单

### 修改文件
- `backend/internal/platform/server.go` — 删除 `mountCompatRoutes` 调用，更新包注释为 phase07-11
- `backend/internal/platform/router.go` — 删除 `mountCompatRoutes` 函数体，删除 `mrhandler` import，更新包注释
- `backend/internal/*/handler/*.go` — 删除 9 个业务模块遗留的 hand-written JSON handler/response 包，实现源码级正式退场

### 新增文件
- `.trae/specs/phase07_11_*/acceptance_report.md` — 正式验收报告

## 验收证据

- [x] `(cd proto && make build && make gen && make lint)` 全部通过
- [x] `(cd backend && go build ./...)` 通过
- [x] `(cd frontend && npx tsc -b --noEmit)` 通过
- [x] `(cd frontend && npm run build)` 通过
- [x] 数据库 baseline 已通过 `reset_phase06_acceptance.sh` 重建
- [x] 后端已在隔离端口 `:18081` 启动复测，healthz 返回 200
- [x] L1-L4 legacy/compat 端点全部返回 404
- [x] 9 个业务模块遗留 hand-written JSON `handler/` 包已全部删除
- [x] 34 条 canonical RPC 全部回归通过（curl 验证）
- [x] 9 个业务模块全部通过 Connect 主线验证
- [x] CR1-CR9 跨模块联动回归全部通过
- [x] 前端静态核销：0 个旧 adapter、0 个旧 import、0 个 legacy 引用
- [x] 11 项 mutation owner 全部核销
- [x] 正式验收记录已产出
