# Phase07-06 业务接口迁移矩阵与回归验收设计 Spec

> **执行产出**：`design.md` — 包含 8 个设计区域：输入边界基线（5 项上游 + 6 项工具链事实）、34 条 RPC 迁移矩阵（4 波次 × 10 列表格，逐条 rpc/service/path/owner/wave/regression/evidence）、跨模块回归清单（17 项模块内 + 9 项跨模块联动 CR1-CR9 + 6 条 route 级）、fixture/联调/退场/收口证据矩阵（5 个脚本映射 + 4 条 legacy endpoint 退场证据 + 8 份 phase07-11 证据包 + 8 条阻断条件）、前端 mutation owner 验收映射（11 项逐条 + 4 组 candidate+mutation 联合验收）、工具链迁移清单（Vite/Proto/脚本/CI 4 表 16 项）、执行顺序总览图（phase07-07→11 流程图）、一致性声明（8 项上游对齐）。
> **执行日期**：2026-08-11
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

`phase07-01` 已冻结 34 条 canonical 业务 RPC 与 4 条 legacy / compat 入口，`phase07-03` 已冻结退场标准，`phase07-05` 已冻结前端 read / mutation owner 收口边界，但当前仍缺一份把“迁移顺序、回归覆盖、脚本承接、工具链迁移、最终证据”收成同一闭环的执行规格。若没有这份矩阵，后续 `phase07-07 ~ phase07-11` 很容易出现“Connect 主线已实现，但 regression 没覆盖完整、脚本链没对齐、phase 收口证据不成套”的漂移。

## What Changes

- 冻结 `phase01 ~ phase06` 的 canonical 业务接口迁移矩阵，要求下钻到 `service / RPC / 当前入口路径 / 页面动作 owner / 迁移 owner / 回归项`
- 冻结跨模块回归清单，覆盖页面读链、正式写链、route 级 first-run、Dashboard / Reuse / Sovereignty 等跨切片路径
- 冻结 fixture、联调、验收与 legacy endpoint retirement evidence 的统一矩阵
- 冻结 legacy endpoint retirement inventory 与 frontend mutation owner inventory 的验收映射关系
- 冻结 Vite、本地启动脚本、reset / acceptance 脚本、`proto/Makefile`、`buf.gen.yaml` 与当前 CI 缺口的迁移清单
- 冻结 `phase07-11` 收口所需的最终证据包结构与阻断条件
- **BREAKING**：`phase07` 不得再以“Connect handler 已接通”或“局部 smoke test 通过”替代完整迁移矩阵与回归验收闭环

## Impact

- Affected specs:
  - `phase07_01_freeze_transport_migration_scope_canonical_boundary`
  - `phase07_02_freeze_chi_connectrpc_buf_formal_composition`
  - `phase07_03_freeze_compat_migration_exit_criteria`
  - `phase07_04_design_go_connect_handler_service_chi_mount`
  - `phase07_05_design_frontend_generated_client_query_application_migration`
  - `phase06_16_integration_validation_acceptance`
- Affected code:
  - `proto/psco/**`
  - `proto/Makefile`
  - `proto/buf.gen.yaml`
  - `frontend/package.json`
  - `frontend/vite.config.ts`
  - `frontend/src/routes/`
  - `frontend/src/features/*/data/`
  - `frontend/src/features/*/application/`
  - `frontend/src/features/*/pages/`
  - `frontend/src/features/*/components/`
  - `backend/internal/platform/router.go`
  - `database/scripts/reset_module_mainline.sh`
  - `database/scripts/reset_decision_mainline.sh`
  - `database/scripts/reset_product_repository_mainline.sh`
  - `database/scripts/reset_dashboard_acceptance.sh`
  - `database/scripts/reset_phase06_acceptance.sh`
  - `.github/workflows/*`（当前仓库无现成 workflow，需作为迁移缺口显式建账）

## ADDED Requirements

### Requirement: Canonical 迁移矩阵必须覆盖 34 条业务 RPC 与其正式迁移 owner

系统 SHALL 在 `phase07-06` 中产出一份单值迁移矩阵，直接复用 `phase07-01` 的 34 条 canonical 业务 RPC 总表，并将每条记录升级为可执行的迁移与回归条目。

每条记录至少必须包含：

- 所属 service
- RPC 名称
- 当前外部访问路径
- 目标 Connect procedure path
- 当前 transport owner
- 当前页面 / route / 面板 / application owner
- 实施迁移 owner（后端 / 前端 / 脚本 / 验收责任位）
- 迁移先后顺序或所在波次
- 最小回归项
- 最终收口证据

#### Scenario: 每条 canonical RPC 都有对应迁移 owner 与回归项

- **WHEN** 接手者检查 `phase07-06` 的迁移矩阵
- **THEN** 34 条 canonical 业务 RPC 必须逐条可查
- **AND** 每条 RPC 都必须有明确迁移 owner、回归项与最终证据
- **AND** 不得只保留模块级摘要而省略 RPC 级责任与验收映射

#### Scenario: Module-centered compat facade 不能被误记为长期 canonical owner

- **WHEN** 矩阵涉及 `ModuleRegistryService` 下仍带 compat 语义的 4 条 RPC
- **THEN** 必须显式区分“`.proto` 中仍存在的 transport inventory”与“正式业务 owner”
- **AND** 不得把旧 JSON compat 路径重新写回长期 canonical 主线

### Requirement: 跨模块回归清单必须覆盖 phase01 ~ phase06 的正式业务链路

系统 SHALL 将回归验收设计冻结为“模块内回归 + 跨模块回归 + 脚本链回归”的统一清单，而不是只做单接口 smoke test。

最小覆盖至少包括：

- `Module Registry` 列表 / 详情 / 创建 / Release
- `Decision Center` 列表 / 详情 / 创建 / Link
- `Product Registry` 列表 / 详情 / 创建 / Bind Module
- `Repository Binding` 列表 / 详情 / 创建 / Bind Product / Map Module
- `Dashboard` Overview / Feedback / Recent Activities
- `Onboarding` first-run route / CTA / 继续主线
- `Export / Backup` 读取与触发
- `Reuse Summary` 在 Dashboard / Module Detail / Product Detail 的正式挂接位
- 从页面读链到 mutation 成功回流再到 read 刷新的跨切片联动

#### Scenario: 回归清单覆盖页面、组件与 route caller

- **WHEN** 接手者定义跨模块回归项
- **THEN** 必须同时覆盖 page 级 read、component 级 candidate read / mutation、以及 route 级 first-run caller
- **AND** 不得只验证页面首屏渲染而忽略 route guard、CTA 分流或 mutation 成功回流

#### Scenario: Dashboard 与 phase05 / phase06 既有验收边界继续保留

- **WHEN** 接手者设计 `Dashboard / Onboarding / Reuse / Sovereignty` 的回归矩阵
- **THEN** 必须复用 `phase05-14` 与 `phase06-16` 已冻结的正式验收边界
- **AND** 不得因为 transport 迁移而降低到“接口可返回 200 即通过”

### Requirement: Fixture、联调、验收与退场证据必须落为同一证据矩阵

系统 SHALL 将 fixture、联调步骤、验收项、legacy endpoint 删除证据与最终 phase 收口证据收敛为同一套矩阵，确保每个回归项都能追溯到环境入口与核销结果。

最小矩阵必须显式承接当前真实入口：

- `database/scripts/reset_module_mainline.sh`
- `database/scripts/reset_decision_mainline.sh`
- `database/scripts/reset_product_repository_mainline.sh`
- `database/scripts/reset_dashboard_acceptance.sh`
- `database/scripts/reset_phase06_acceptance.sh`

并明确：

- 默认恢复与 fixture 恢复入口
- 对应业务模块 / RPC / 页面动作 owner
- 联调步骤
- 期望结果
- 删除证据
- 回归证据
- 最终记录位置

#### Scenario: 现有 reset 脚本必须被纳入 phase07 证据链

- **WHEN** 接手者设计 fixture 与验收矩阵
- **THEN** 必须基于当前仓库真实存在的 reset / acceptance 脚本编排
- **AND** 不得凭空假设另一套尚不存在的测试入口

#### Scenario: 每条 legacy endpoint 都有 endpoint 级删除证据与等价回归证据

- **WHEN** 接手者将 `phase07-03` 的 4 条 legacy / compat 入口纳入证据矩阵
- **THEN** 每条入口都必须同时映射：
  - 路由删除证据
  - handler / adapter 删除证据
  - 替代 Connect 路径回归证据
- **AND** 不得只写“已改走 Connect”而缺少旧入口真正退场的 endpoint 级核销

### Requirement: Frontend mutation owner inventory 必须升级为验收映射表

系统 SHALL 将 `phase07-05` 已冻结的 11 项前端正式写动作 owner 清单，升级为可执行的验收映射表。

验收映射至少必须回答：

- 当前 owner 位置
- 目标 owner 位置
- 允许保留的短时过渡位
- 对应触发页面 / 组件
- 触发前置 fixture
- 成功回流与 query 失效检查项
- 最晚核销时点

#### Scenario: Canonical mutation 全部进入验收映射

- **WHEN** 接手者审查写路径验收设计
- **THEN** `CreateDraft*`、`CreateRelease`、`Bind*`、`Map*`、`Link*`、`ExportCoreAssets`、`CreateInstanceBackup` 必须全部在映射表中
- **AND** 不得遗漏 `SovereigntyPanel` 的两个低频过渡位

#### Scenario: Candidate read 与 mutation owner 必须一起验收

- **WHEN** 某个组件同时承接 candidate read 与 mutation
- **THEN** 验收设计必须同时验证 candidate read 已回收到 read owner、mutation 已回收到 application owner 或声明过渡位
- **AND** 不得只验 mutation 成功而忽略旧 adapter 读取主线是否仍残留

### Requirement: `/api` 单一前缀与 proto 生成链承接位必须冻结到工具链迁移清单

系统 SHALL 将以下工具链与运行链路纳入同一份迁移清单：

- `frontend/vite.config.ts` 的 `/api` dev proxy
- 本地启动脚本或当前缺失情况
- `database/scripts/reset_*.sh` 与验收入口
- `proto/Makefile` 的 `build / gen / lint / breaking`
- `proto/buf.gen.yaml` 的生成插件矩阵
- 当前 CI workflow 的存在状态或缺口

迁移清单必须明确：

- 当前源码事实
- 目标 phase07 状态
- 谁负责修改
- 验证命令 / 验证方式
- 阻断条件

#### Scenario: Vite dev proxy 保持单一 `/api` 基址

- **WHEN** 接手者梳理前端 dev 与验收链路
- **THEN** 必须明确 `frontend/vite.config.ts` 继续通过单一 `/api` 前缀转发到后端
- **AND** 不得因为 Connect procedure path 引入并列浏览器访问基址

#### Scenario: 当前仓库没有现成 CI workflow 也必须显式建账

- **WHEN** 接手者梳理 proto 生成链在 CI 中的承接位
- **THEN** 若仓库当前不存在 `.github/workflows/*` 等正式 CI 入口，必须把“当前缺失”记录为显式迁移缺口
- **AND** 不得把不存在的 CI 入口假装成已存在配置
- **AND** phase 收口前必须明确该缺口由何处补齐或如何提供等价证据

### Requirement: Phase07 收口最终证据必须以单一证据包定义

系统 SHALL 将 `phase07-11` 需要核销的最终证据提前冻结成单一证据包结构，至少覆盖：

- 34 条 canonical RPC 已完成迁移的核销结果
- 4 条 legacy / compat 入口的 endpoint 级退场结果
- 前端 mutation owner 收口结果与允许保留过渡项
- `/api` 单一基址在 dev / 验收 / 部署链路中的承接结果
- proto 生成链、本地脚本链与 CI 缺口处理结果
- 跨模块回归通过结果

#### Scenario: Phase closure is blocked by missing final evidence

- **WHEN** 某条 canonical RPC、legacy endpoint、脚本链或回归项没有对应最终证据
- **THEN** `phase07` 不得收口
- **AND** 不得用“实现已完成，证据稍后补”绕过验收门禁

## MODIFIED Requirements

### Requirement: phase07-01 的 canonical 总表升级为迁移与回归执行矩阵

`phase07-01` 已冻结 34 条 canonical 业务 RPC 的基础总表。

自 `phase07-06` 起，系统必须把这一 requirement 修改为：

- 总表不再只是 transport inventory，还必须包含迁移波次、迁移 owner、回归项与最终证据
- 每条 RPC 的页面 / route / 组件 / application owner 必须进入同一执行矩阵
- `phase07-07 ~ phase07-11` 必须直接复用这份矩阵推进实现与验收

#### Scenario: 总表从“建账”升级为“可执行”

- **WHEN** 团队进入实现与验收阶段
- **THEN** 可以直接从 `phase07-06` 读出每条 RPC 的迁移责任与回归方式
- **AND** 不得再二次发明并列迁移清单

### Requirement: phase07-03 的 legacy retirement inventory 升级为 endpoint 级证据矩阵

`phase07-03` 已冻结 legacy / compat 入口的退场窗口与证据模型。

自 `phase07-06` 起，系统必须把这一 requirement 修改为：

- 每条 legacy endpoint 都必须绑定实际回归项、fixture / 环境入口与核销记录位置
- 前端 adapter 删除证据、后端 route 删除证据与替代 Connect 路径回归必须在同一行闭环

#### Scenario: Compat inventory 直接服务 phase07 收口核销

- **WHEN** 团队检查某条 compat 入口是否可退场
- **THEN** 必须能从 `phase07-06` 直接查到删除证据与等价回归项
- **AND** 不得仅靠口头说明或散落日志判断

### Requirement: phase07-05 的 mutation owner inventory 升级为验收映射

`phase07-05` 已冻结前端 read / mutation owner 收口设计。

自 `phase07-06` 起，系统必须把这一 requirement 修改为：

- mutation owner 清单不再停留在设计层，还必须明确 fixture、触发动作、成功回流检查项与最晚核销时点
- route caller、candidate read 与 mutation owner 的联合回归必须显式入表

#### Scenario: Frontend transport migration can be validated end-to-end

- **WHEN** 团队验证某条前端正式写动作是否迁移完成
- **THEN** 必须同时看到 read owner、mutation owner、回流刷新与旧 adapter 退场证据
- **AND** 不得只验证按钮点击成功一次

## REMOVED Requirements

### Requirement: 以局部 smoke test 或单接口连通性证明 phase07 等价迁移成立

**Reason**: 单接口连通性无法证明 34 条 canonical RPC、4 条 legacy / compat 入口、前端 owner 收口、脚本链承接与最终收口证据已经形成闭环。

**Migration**:

- 改为以单一迁移矩阵统筹 canonical RPC、legacy endpoint、frontend owner、脚本链与最终证据
- 改为以跨模块回归与 endpoint 级核销作为 phase07 收口的正式依据

### Requirement: 工具链迁移可以在实现阶段临时决定

**Reason**: 当前仓库已经存在真实 `vite.config.ts`、`proto/Makefile` 与多份 reset / acceptance 脚本，同时 CI workflow 处于缺失状态；若不先显式建账，后续很容易出现本地能跑、验收链不通、CI 口径空缺的漂移。

**Migration**:

- 在 `phase07-06` 中统一建立 Vite、本地脚本、验收脚本、proto 生成链与 CI 缺口清单
- 后续实现与收口统一按该清单核销，而不是各子任务临时发挥
