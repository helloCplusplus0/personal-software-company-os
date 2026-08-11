# phase07-03 冻结迁移过程兼容策略与 Phase 收口退场标准 Spec

> **执行产出**：`frozen_scope.md` — 包含 compat 策略（4 前提 + 4 禁止）、4 条 legacy inventory 退场版本（含调用方、替代 RPC、最晚时点、删除/回归证据各 3 项）、退场证据模型（后端 4 项 + 前端 5 项）、子任务链时点映射（phase07-09/10/11）、双门槛收口判定标准。
> **执行日期**：2026-08-11
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

`phase07-01` 已冻结迁移范围与 legacy inventory，`phase07-02` 已冻结正式传输主线组合，但当前仓库仍存在手写 JSON business handler、module-centered compat 入口和前端遗留 adapter 导出。如果不先把“哪些兼容允许短时存在、何时必须退场、退场如何举证”冻结成单值规则，后续实现很容易把临时适配层写成长期正式接口。

## What Changes

- 冻结 `phase07` 迁移过程中是否允许短时并存 JSON adapter / hand-written JSON business handler
- 冻结 compat 资产允许存在的前提、禁止事项与责任边界
- 冻结当前真实 legacy / compat 业务入口 inventory 的正式退场清单
- 冻结每个 legacy / compat 入口的替代 RPC / Connect path、最晚并存时点、删除证据与回归证据
- 冻结 `phase07` 收口时“旧 JSON 业务主线已退场”的验收门槛
- **BREAKING**：`phase07` 收口后不允许保留任何未声明的 hand-written JSON canonical 业务主线

## Impact

- Affected specs:
  - `phase07-04` Go Connect handler / compat adapter 承接设计
  - `phase07-05` 前端 query / application 迁移设计
  - `phase07-06` transport 实现与双栈过渡策略
  - `phase07-09` 后端 compat 路由退场
  - `phase07-10` 前端 adapter 回收与 mutation 收口
  - `phase07-11` 验收与收口核销
- Affected code:
  - `backend/internal/platform/router.go`
  - `backend/internal/moduleregistry/handler/query_handler.go`
  - `backend/internal/moduleregistry/handler/command_handler.go`
  - `frontend/src/features/module-registry/data/api-adapter.ts`
  - `frontend/src/features/module-registry/data/module-registry-adapter.ts`
  - `frontend/src/features/product-registry/data/api-adapter.ts`
  - `frontend/src/features/repository-binding/data/api-adapter.ts`
  - `frontend/src/features/*/application/*`
  - 联调验收脚本与 phase07 验收记录

## ADDED Requirements

### Requirement: Compat 资产只允许在迁移过程中短时存在

系统 SHALL 允许 compat 资产在 `phase07` 迁移过程中短时并存，但其身份只能是：

- 迁移中的临时 adapter
- legacy caller 的过渡承接位
- canonical Connect 主线切换前的短时兼容层

系统 SHALL NOT 把 compat 资产写成长期正式接口、长期正式 query owner、长期正式 mutation owner 或长期正式 router 主线。

#### Scenario: Temporary coexistence is allowed during active migration

- **WHEN** 某条 canonical 业务接口尚未完成前后端同步切换
- **THEN** 对应的 legacy JSON adapter / handler 可以短时保留
- **AND** 该保留必须带有明确替代 RPC、最晚退场时点和删除证据要求

#### Scenario: Compat layer cannot become a second canonical mainline

- **WHEN** 某个 compat 入口仍能被调用
- **THEN** 它只能被视为迁移过程的过渡资产
- **AND** 不得被解释为 `phase07` 收口后的长期正式 API

### Requirement: 兼容并存必须满足明确前提与禁止事项

系统 SHALL 仅在以下前提同时满足时允许 compat 并存：

- canonical `.proto + Connect` 替代路径已明确
- 当前 caller 和承接位已可定位
- 最晚并存时点已写入阶段任务链
- 删除证据与回归证据已预先定义

系统 SHALL 明确以下禁止事项：

- 不得为图省事保留“新 Connect + 旧 JSON”双主线长期并列
- 不得新增第二批未列入 inventory 的 compat 业务入口
- 不得把前端临时 adapter 留在 route、page 或展示组件中长期使用
- 不得把“当前无 active caller，但导出仍在”解释为可无限期保留

#### Scenario: Untracked compat entry is rejected

- **WHEN** 实现过程中新增了未进入 inventory 的 JSON compat 入口
- **THEN** 该入口不应被接受为 phase07 合法过渡资产
- **AND** 必须先补入正式 inventory 或直接删除

#### Scenario: Dormant export is still subject to retirement

- **WHEN** 某个 compat adapter 当前没有 active UI caller，但仍保留导出
- **THEN** 它仍属于待退场 legacy 资产
- **AND** 必须在对应最晚时点前删除，而不是以“当前没人调用”为理由长期保留

### Requirement: Legacy / Compat 业务入口 inventory 必须单值冻结

系统 SHALL 将当前真实 legacy / compat 业务入口 inventory 至少冻结为以下 4 项：

- `GET /api/candidates/products`
- `GET /api/candidates/repositories`
- `POST /api/modules/{moduleId}/bindings/products`
- `POST /api/modules/{moduleId}/bindings/repositories`

每个入口 SHALL 明确：

- 当前调用方
- 存在原因
- 对应替代 RPC / Connect path
- 允许并存的最晚时点
- 退场时的删除证据
- 退场时的回归证据

#### Scenario: Inventory includes all known module-centered compat entries

- **WHEN** 团队检查当前 module-centered compat 业务入口
- **THEN** 上述 4 条入口都必须在正式 inventory 中
- **AND** 不得遗漏当前仍在 `router.go` 中注册的 compat 路由

#### Scenario: Each entry has explicit retirement mapping

- **WHEN** 团队计划删除某条 compat 入口
- **THEN** 必须能直接查到其替代 RPC / Connect path、最晚时点和删除/回归证据
- **AND** 不允许凭口头共识执行退场

### Requirement: Legacy / Compat 入口的最晚退场时点必须与 phase07 子任务链对齐

系统 SHALL 将 legacy / compat 入口的退场时点冻结到 phase07 子任务链中。

最小冻结要求：

- 候选读取 compat 入口在后端切换完成前退场
- module-centered 绑定 compat 入口在前端 adapter 清理完成前退场
- phase07 收口前不得残留任何旧 JSON business mainline

#### Scenario: Candidate compat exits before phase closure

- **WHEN** `ProductRegistry / RepositoryBinding` 的 Connect 读路径已可承接候选读取
- **THEN** `GET /api/candidates/products` 与 `GET /api/candidates/repositories` 必须按冻结时点退场
- **AND** 不得推迟到 phase07 收口之后

#### Scenario: Binding compat exits after frontend owner switch is complete

- **WHEN** 前端正式写入 owner 已完成切换且 module-centered JSON adapter 不再承担正式写动作
- **THEN** `POST /api/modules/{moduleId}/bindings/products` 与 `POST /api/modules/{moduleId}/bindings/repositories` 必须退场
- **AND** 不得保留为长期兜底入口

### Requirement: Phase07 收口必须满足旧 JSON 业务主线退场标准

系统 SHALL 仅在以下条件全部满足时，才能判定 `phase07` 收口：

- canonical 业务主线以 `.proto + Connect` 承接
- legacy inventory 中的 4 个 compat 业务入口均已核销
- 后端不再保留 hand-written JSON business handler 作为正式业务主线
- 前端不再保留未声明的长期 JSON adapter 主线
- 验收记录包含删除证据和回归证据

#### Scenario: Phase closure is blocked by leftover compat route

- **WHEN** `router.go` 中仍保留 inventory 内的 compat 业务路由
- **THEN** `phase07` 不得被判定为完成
- **AND** 必须继续推进退场任务直到删除完成

#### Scenario: Phase closure is blocked by leftover frontend adapter mainline

- **WHEN** 某个旧 JSON adapter 仍作为正式业务主线承接真实页面动作
- **THEN** `phase07` 不得收口
- **AND** 必须先完成 Connect client 切换、adapter 删除与回归验证

## MODIFIED Requirements

### Requirement: phase07-01 的 legacy inventory 从“列清单”升级为“带退场门槛的正式兼容策略”

`phase07-01` 已冻结 4 个 legacy / compat 业务入口及其替代关系。

自 `phase07-03` 起，系统必须把这一 requirement 修改为：

- inventory 不再只是静态清单，还必须绑定并存前提、最晚退场时点、删除证据与回归证据
- “当前无 active caller”的入口也必须纳入退场计划
- 收口判定必须逐项核销 inventory，而不是只证明 Connect 主线已存在

#### Scenario: Inventory is used as retirement gate instead of documentation only

- **WHEN** 团队执行 phase07 验收
- **THEN** 必须逐项检查 inventory 是否已退场
- **AND** 不得把 inventory 仅当作文档备注

### Requirement: phase07-02 的正式传输主线冻结补充 compat 退场约束

`phase07-02` 已冻结 `chi + ConnectRPC + buf` 的正式组合方式。

自 `phase07-03` 起，系统必须把这一 requirement 修改为：

- 允许 compat 资产只在 Connect 主线迁移完成前短时存在
- Connect 主线落地本身不等于 phase07 完成，必须同时满足 legacy inventory 退场
- `/api` 下的旧 JSON business path 不得在 phase 收口后继续作为长期 fallback

#### Scenario: Connect mainline alone does not complete the phase

- **WHEN** Connect handler、生成链和前端 client 已全部具备
- **THEN** 若 legacy inventory 仍未退场，phase07 仍不能收口
- **AND** 验收必须继续阻塞在 compat 清理项

## REMOVED Requirements

### Requirement: 旧 JSON adapter 可以作为长期 fallback 保留到 phase07 之后

**Reason**: 这会把 compat 资产从迁移工具变成第二条正式主线，直接破坏 `phase07` 的单主线目标。

**Migration**:

- 先用 Connect 主线替代真实 caller
- 再删除 compat 路由、handler 和前端 adapter 导出
- 最后用回归记录证明退场完成，而不是仅保留“理论上可删”
