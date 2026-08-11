# Tasks

- [x] Task 1: 冻结 compat 资产只允许短时存在的正式策略。把 phase07 中"允许并存"与"禁止长期并列"的边界写成单值规则，避免后续实现把迁移过渡层写成正式接口。
  - [x] SubTask 1.1: 对齐 `phase07-01` legacy inventory 与 `phase07-02` 单主线 transport 结论 → `frozen_scope.md` §1.1 合法身份定义（引用 phase07-01/02）
  - [x] SubTask 1.2: 在规格中明确 compat 资产的合法身份、允许并存前提与禁止事项 → `frozen_scope.md` §1.2（4 项前提）+ §1.3（4 条禁止事项）
  - [x] SubTask 1.3: 在规格中明确"当前无 active caller"的 adapter 导出仍属于待退场资产 → `frozen_scope.md` §1.4 Dormant Asset 规则

- [x] Task 2: 冻结当前真实 legacy / compat 业务入口 inventory。至少覆盖 4 条 module-centered compat 入口，并为每条入口绑定替代路径与退场要求。
  - [x] SubTask 2.1: 核对 `router.go` 中仍注册的 4 条 compat 路由 → `frozen_scope.md` §2.1-2.4（与 phase07-01 一致 + 当前 router.go 验证）
  - [x] SubTask 2.2: 逐条写明当前调用方、存在原因、替代 RPC / Connect path → `frozen_scope.md` §2.1-2.4 表格（每表含 7 个属性列）
  - [x] SubTask 2.3: 逐条写明最晚并存时点、删除证据与回归证据 → `frozen_scope.md` §2.1-2.4 表格（删除证据 3 项、回归证据 3 项）

- [x] Task 3: 冻结 compat 入口与前端残留 adapter 的退场标准。把"路由删除、handler 删除、前端切换、回归验证"收成统一的核销证据模型。
  - [x] SubTask 3.1: 明确后端退场证据必须至少包含路由删除与 handler 删除 → `frozen_scope.md` §3.1 后端退场证据表（4 项）
  - [x] SubTask 3.2: 明确前端退场证据必须至少包含正式 caller 切换与旧 adapter 导出删除 → `frozen_scope.md` §3.2 前端退场证据表（5 项）
  - [x] SubTask 3.3: 明确回归证据至少覆盖替代 Connect path 可用与旧路径不可再访问 → `frozen_scope.md` §3.3 回归证据统一模型（3 项）

- [x] Task 4: 冻结 legacy / compat 入口的最晚退场时点与 phase07 子任务链映射。确保 phase07-09、phase07-10、phase07-11 的职责边界清晰，不出现"到收口时再看"的模糊状态。
  - [x] SubTask 4.1: 为候选读取 compat 入口冻结最晚退场时点 → `frozen_scope.md` §4.1（phase07-09）
  - [x] SubTask 4.2: 为 module-centered 绑定 compat 入口冻结最晚退场时点 → `frozen_scope.md` §4.2（phase07-10）
  - [x] SubTask 4.3: 明确 phase07 收口前不得残留任何旧 JSON business mainline → `frozen_scope.md` §4.3 冻结约束

- [x] Task 5: 冻结 phase07 收口退场标准。把"Connect 主线已存在"与"旧 JSON 主线已退场"拆成同时必须满足的双门槛。
  - [x] SubTask 5.1: 明确 phase07 收口必须逐项核销 legacy inventory → `frozen_scope.md` §5.2 门槛二（4 项条件）
  - [x] SubTask 5.2: 明确后端仍保留 compat 业务路由时不得判定完成 → `frozen_scope.md` §5.3 禁止判定方式
  - [x] SubTask 5.3: 明确前端仍保留未声明的长期 JSON adapter 主线时不得判定完成 → `frozen_scope.md` §5.3 禁止判定方式

- [x] Task 6: 完成与 phase07 上游冻结文档的一致性校验。确保本次兼容策略既承接 `phase07-01` 的 inventory，又不与 `phase07-02` 的正式主线冻结冲突。
  - [x] SubTask 6.1: 校验与 `phase07-01` 的 4 条 compat inventory、替代关系一致 → `frozen_scope.md` §6 一致性声明
  - [x] SubTask 6.2: 校验与 `phase07-02` 的单一 `/api`、单一 Connect 主线、旧 JSON 仅作 compat 资产的结论一致 → `frozen_scope.md` §6 一致性声明
  - [x] SubTask 6.3: 校验本次规格足以直接指导 `phase07-09 / 10 / 11` 的实现与验收 → `frozen_scope.md` §4 子任务链映射 + §5 双门槛

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 2` ✅
- `Task 4` depends on `Task 2` ✅
- `Task 5` depends on `Task 3` and `Task 4` ✅
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5` ✅