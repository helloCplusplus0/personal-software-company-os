# Tasks

- [x] Task 1: 对齐 `phase07-08` 的直接上游与阶段边界，明确本次任务是"落实生成链与正式合同产物主线"，不是并行展开 handler / frontend owner 实现。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase07_transport_contract_mainline_migration_dev_plan.md#L168-180` 的范围与 DoD。
  - [x] SubTask 1.2: 对齐 `phase07-07 formal spec` 中关于 3 插件矩阵、生成产物目录、前端依赖与本地等价验证命令的正式口径。
  - [x] SubTask 1.3: 明确本阶段完成条件是"生成链稳定产出正式合同产物"，而不是"9 模块业务实现完成"。

- [x] Task 2: 落实 `proto/buf.gen.yaml` 的正式生成矩阵与产物目录。
  - [x] SubTask 2.1: 将当前 2 插件矩阵升级为 `protocolbuffers/go + connectrpc/gosimple + bufbuild/es`（含 `simple` 选项）。
  - [x] SubTask 2.2: 冻结 Go Connect 产物落点为 `backend/internal/gen/connect/**`，并保持 `simple` 模式。
  - [x] SubTask 2.3: 冻结 TS 合同产物继续单值输出到 `frontend/src/gen/proto/**`，不得新增第二生成根。

- [x] Task 3: 更新 `proto/Makefile` 与生成入口的一致性。
  - [x] SubTask 3.1: 明确 `build / gen / lint / breaking / clean` 继续复用现有 `proto/Makefile` 入口。
  - [x] SubTask 3.2: 验证 `make gen` 与 `buf generate` 在 `proto/` 目录下产出同一套正式合同产物。
  - [x] SubTask 3.3: 更新 `make clean` 覆盖 `backend/internal/gen/proto`、`backend/internal/gen/connect` 与 `frontend/src/gen/proto`。

- [x] Task 4: 补齐前端 Connect runtime 依赖与 lockfile。
  - [x] SubTask 4.1: 安装 `@connectrpc/connect@^2.1.2` 与 `@connectrpc/connect-web@^2.1.2`。
  - [x] SubTask 4.2: 确认 `@bufbuild/protobuf` 继续保留。
  - [x] SubTask 4.3: 确认 lockfile 与正式依赖口径同步更新。

- [x] Task 5: 执行 `buf generate` 生成并验证产物完整性。
  - [x] SubTask 5.1: 执行 `make clean && make gen`，验证 10 个 Go protobuf + 9 个 Go Connect + 10 个 TS 产物完整。
  - [x] SubTask 5.2: 验证 Go Connect simple 模式签名正确（`*pb.ListModulesRequest` 直接使用，无 `connect.Request` 包装）。
  - [x] SubTask 5.3: 验证 `go build ./...` 与 `npx tsc -b --noEmit` 通过。
  - [x] SubTask 5.4: 验证 `make build` 与 `make lint` 通过。

- [x] Task 6: 更新 `proto/README.md` 与文档承接。
  - [x] SubTask 6.1: 更新 `proto/README.md` 的生成产物落点表，从 10 模块手写表升级为 3 插件矩阵表。
  - [x] SubTask 6.2: 更新 `proto/README.md` 的 §6 "当前阶段不做"，反映 Connect 生成链已落地但 handler 实现待后续子任务。
  - [x] SubTask 6.3: 明确后续实现不得再以"缺少 Connect 产物 / 缺少前端 Connect 依赖"为默认前提。

- [x] Task 7: 完成 `phase07-08` 规格一致性校验。
  - [x] SubTask 7.1: 验证本 spec 已继承 `phase07-02` 与 `phase07-07` 的 3 插件矩阵、产物目录与工具链口径。
  - [x] SubTask 7.2: 验证本 spec 已对齐 `phase07-06` 的工具链迁移矩阵与 CI 缺口处理。
  - [x] SubTask 7.3: 验证本 spec 未把 handler / frontend owner 业务实现误纳入当前阶段 DoD。
  - [x] SubTask 7.4: 验证本 spec 与当前仓库事实一致：`buf.gen.yaml` 已切换为 3 插件、前端已安装 Connect runtime、Go 后端已安装 `connectrpc.com/connect`。

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 1`, `Task 2` ✅
- `Task 4` depends on `Task 1`, `Task 2` ✅
- `Task 5` depends on `Task 2`, `Task 3`, `Task 4` ✅
- `Task 6` depends on `Task 3`, `Task 4`, `Task 5` ✅
- `Task 7` depends on `Task 1` through `Task 6` ✅
