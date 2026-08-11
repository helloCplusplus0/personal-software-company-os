# Tasks

- [x] Task 1: 冻结 `phase07` 必须迁移的 canonical 业务模块范围。把 `phase01 ~ phase06` 已交付业务模块统一压成 `phase07` 的正式迁移范围，不允许留下"以后再迁"的灰区。
  - [x] SubTask 1.1: 根据 `phase07` shared baseline 冻结 9 个必须迁移的 canonical 业务模块清单 → `frozen_scope.md` §1
  - [x] SubTask 1.2: 说明这些模块对应的页面 / 业务主线，以及为什么属于 `phase07` 一次性正式切换范围 → `frozen_scope.md` §1 表格
  - [x] SubTask 1.3: 明确哪些"只迁新接口 / 只迁单模块试点 / 旧主线长期兼容"解释属于禁止项 → `frozen_scope.md` §1 禁止项

- [x] Task 2: 冻结 canonical 业务接口迁移总表。把迁移清单下钻到 `service / RPC / 当前入口路径 / 页面或动作 owner` 级别，形成后续 `/spec`、实现和验收都能直接复用的总表要求。
  - [x] SubTask 2.1: 为每个 canonical 业务模块定义 service / RPC 级别的迁移表结构 → `frozen_scope.md` §2.1-2.9
  - [x] SubTask 2.2: 明确每条记录必须至少包含：当前入口路径、目标 Connect procedure path、当前 transport owner、迁移后正式 owner → `frozen_scope.md` §2 各子表均含全部 7 列
  - [x] SubTask 2.3: 明确这份总表将作为 `phase07-06` 回归矩阵与 `phase07-11` 验收核销的直接上游 → `frozen_scope.md` §2.10 统计

- [x] Task 3: 冻结非业务基础设施端点边界与 keep list。明确哪些端点继续留在 `chi + net/http`，哪些端点不得被误判为可长期保留的业务接口。
  - [x] SubTask 3.1: 冻结 `healthz / readyz / metrics / debug / pprof` 为唯一允许保留的非业务基础设施端点 → `frozen_scope.md` §3 表格
  - [x] SubTask 3.2: 明确除上述 keep list 外，其余业务主线接口不得以"暂时先不迁"名义留在正式完成态 → `frozen_scope.md` §3 约束
  - [x] SubTask 3.3: 明确 keep list 的职责只限基础设施，不承担 `.proto` 对外合同职责 → `frozen_scope.md` §3 约束

- [x] Task 4: 冻结当前真实 `legacy / compat` 业务入口 inventory。把已经存在的兼容入口点名建账，为 `phase07-03 / 09 / 11` 的退场、实现和验收提供 endpoint 级依据。
  - [x] SubTask 4.1: 至少点名 `/api/candidates/products`、`/api/candidates/repositories`、`/api/modules/{moduleId}/bindings/products`、`/api/modules/{moduleId}/bindings/repositories` → `frozen_scope.md` §4 表格
  - [x] SubTask 4.2: 为每个入口明确当前调用方、存在原因、替代 RPC / Connect path、允许并存的最晚时点与删除证据 → `frozen_scope.md` §4 表格含全部 7 列
  - [x] SubTask 4.3: 明确这些入口不得作为 `phase07` 收口后的长期兼容层继续保留 → `frozen_scope.md` §4 约束

- [x] Task 5: 冻结 `phase07` 收口时"业务主线已切换"的判定标准。把完成条件写成明确门禁，避免实现结束后再以主观判断决定是否通过。
  - [x] SubTask 5.1: 明确 canonical 业务接口全部切到 ConnectRPC 才能算主线切换完成 → `frozen_scope.md` §5 条件 1
  - [x] SubTask 5.2: 明确单一 `/api` 基址在开发、验收、部署链路中继续成立 → `frozen_scope.md` §5 条件 3
  - [x] SubTask 5.3: 明确 legacy inventory 必须逐项核销，且不允许"新 Connect 主线 + 旧 JSON 主线"并列正式存在 → `frozen_scope.md` §5 条件 4-5

- [x] Task 6: 冻结前端页面或动作 owner 进入迁移范围表的要求。确保后续不仅迁 transport，还明确正式读写动作的承接位。
  - [x] SubTask 6.1: 明确每个 RPC 需要映射到页面、面板、query owner、application owner 或其他正式动作承接位 → `frozen_scope.md` §6.1-6.3
  - [x] SubTask 6.2: 明确当前仍位于页面 / 组件中的正式 mutation 必须被标记为"回收至 application owner"或"短时过渡位" → `frozen_scope.md` §6.3 表格
  - [x] SubTask 6.3: 明确"保持 query / application 边界"必须落为 owner 清单，而不是抽象口号 → `frozen_scope.md` §6.2 含 4 个 concrete application owner 文件

- [x] Task 7: 完成规格一致性校验。确保本次 `phase07-01` 规格与 `phase07` 三件套、`audit_001`、`phase06` 正式规格与验收结论完全对齐。
  - [x] SubTask 7.1: 校验本次规格没有越权扩写 `mvp0.3` 业务能力 → 未在 frozen_scope.md 中出现 mvp0.3 内容
  - [x] SubTask 7.2: 校验本次规格与 `phase07` shared baseline 中的 9 个业务模块、keep list 与验收基线一致 → `frozen_scope.md` §7 一致性声明
  - [x] SubTask 7.3: 校验本次规格与 `architecture_plan` 中"一次性正式切换"的定位一致 → `frozen_scope.md` §7 一致性声明
  - [x] SubTask 7.4: 校验本次规格足以直接支撑后续 `phase07-02 ~ phase07-07` 子任务 → 34 条 proto-defined business RPC 总表 + 4 个 legacy/compat HTTP 入口 inventory 可直接复用

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 1` ✅
- `Task 4` depends on `Task 2` ✅
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4` ✅
- `Task 6` depends on `Task 2` ✅
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6` ✅
