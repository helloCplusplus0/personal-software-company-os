# Phase07-10 落实前端客户端与业务切片调用切换 Spec

## Why

`phase07-08` 已完成 `buf + ConnectRPC` 正式合同产物主线，`phase07-09` 已完成 Go 后端业务传输主线切换。当前剩余的阻断点在前端：真实业务调用仍分散在 hand-written `api-adapter.ts`、页面内 `useQuery/useMutation`、以及少量 module-centered compat 调用上，尚未形成与 `.proto + ConnectRPC` 对齐的单一正式主线。

## What Changes

- 建立前端唯一共享 Connect transport，并保持浏览器侧单一 `/api` 基址
- 为各业务切片建立 slice-local generated client 承接位，统一消费 `frontend/src/gen/proto/**`
- 将页面/组件内分散的 `useQuery` 回收到切片 `data/` 下的 read owner，与 route/CTA 共享只读 helper 对齐
- 将 canonical 写动作收口到切片 `application/` 下的单一 owner；仅允许 Export / Backup 在 `SovereigntyPanel` 内保留显式过渡位
- 回收 hand-written fetch / JSON 主线，删除 8 个旧 adapter 文件及不再需要的 adapter 壳与 compat facade
- 核销 L3/L4 module-centered compat 的前端调用与导出，确保前端正式调用只走 canonical owner 的 Connect client
- **BREAKING**：前端正式业务主线不再通过 `api-adapter.ts`、`module-registry-adapter.ts` 或其他 hand-written fetch/JSON 封装发起业务请求

## Impact

- Affected specs:
  - `phase07_05_design_frontend_generated_client_query_application_migration`
  - `phase07_06_design_business_api_migration_matrix_regression_acceptance`
  - `phase07_07_formal_transport_mainline_cutover_spec`
  - `phase07_09_cut_go_backend_transport_mainline`
- Affected code:
  - `frontend/src/gen/proto/**`
  - `frontend/src/features/*/data/*.ts`
  - `frontend/src/features/*/application/*.ts`
  - `frontend/src/features/*/pages/*.tsx`
  - `frontend/src/features/*/components/*.tsx`
  - `frontend/src/routes/index.tsx`
  - `frontend/vite.config.ts`

## ADDED Requirements

### Requirement: 前端必须建立单一 Connect transport 与 slice-local client 主线

系统 SHALL 在前端建立唯一共享 Connect transport，并要求各业务切片通过 slice-local generated client 发起 RPC，不得在页面、组件或多个共享层级中并列维护第二套 transport / client 入口。

#### Scenario: 共享 transport 与切片 client 收敛

- **WHEN** 团队执行 `phase07-10`
- **THEN** 前端必须存在唯一共享 Connect transport 承接位，且 `baseUrl` 保持 `/api`
- **AND** Module Registry / Decision Center / Product Registry / Repository Binding / Dashboard / Onboarding / Reuse Summary 必须各自拥有 slice-local `connect-client.ts` 或等价承接位
- **AND** 页面与组件不得直接调用 `createConnectTransport()` 或 `createClient()`
- **AND** 不得新增第二套 API 基址、第二套 transport 文件或第二套跨切片 client 宿主

### Requirement: Query 层必须回收到切片只读 owner

系统 SHALL 将当前分散在页面、组件与 route 中的 canonical 读取动作回收到切片 `data/` 下的只读 owner / helper，保持 `query` 层纯只读，不让 transport 切换演化为新的页面级拼装模式。

#### Scenario: 页面与组件读取回收

- **WHEN** 团队切换前端读取主线
- **THEN** `useQuery` 的正式 `queryKey / queryFn` 必须收敛到切片 `data/use-*-read.ts` 或等价只读 owner
- **AND** 页面与组件只能消费只读 owner，不得继续内联长期正式 `queryFn`
- **AND** 候选读取也必须回收到切片只读 owner，不得长期滞留在组件内
- **AND** `src/routes/index.tsx` 与 `dashboard/components/onboarding-cta-button.tsx` 必须复用 Onboarding 切片同一只读 helper，不得继续直连 `onboarding/data/api-adapter.ts`

### Requirement: Canonical 写动作必须收敛到单一正式 owner

系统 SHALL 为 canonical 写动作建立单一正式 owner，优先落在切片 `application/` 下；只有 `ExportCoreAssets` 与 `CreateInstanceBackup` 可在 `phase07-10` 作为显式过渡位暂留在 `SovereigntyPanel`，但其 transport 也必须切到 generated client。

#### Scenario: Mutation owner 收口

- **WHEN** 团队执行 `phase07-10`
- **THEN** `CreateModule / CreateDecision / CreateProduct / CreateRepository` 必须继续保留在既有 `application/use-create-draft-*.ts`，仅替换 transport
- **AND** `CreateRelease / BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository / LinkDecisionToTarget` 必须进入切片 `application/` 下的单一正式 owner
- **AND** 页面与组件不得长期保留与上述 canonical 写动作并列的正式 `useMutation`
- **AND** `ExportCoreAssets / CreateInstanceBackup` 若暂留在 `SovereigntyPanel`，必须被显式标记为允许过渡位，并附退场条件

### Requirement: phase07-10 必须完成前端 adapter 与 compat facade 退场

系统 SHALL 在本阶段删除前端 hand-written fetch / JSON 正式主线，并核销 module-centered compat 的前端调用与导出，避免后端已切主线而前端继续长期保留第二套 transport。

#### Scenario: 旧 adapter 与 compat 导出删除

- **WHEN** `phase07-10` 完成
- **THEN** `dashboard/data/api-adapter.ts`、`dashboard/data/sovereignty-api-adapter.ts`、`onboarding/data/api-adapter.ts`、`reuse-summary/data/api-adapter.ts`、`decision-center/data/api-adapter.ts`、`product-registry/data/api-adapter.ts`、`repository-binding/data/api-adapter.ts`、`module-registry/data/api-adapter.ts` 必须完成删除或退场
- **AND** 对应 `*-adapter.ts` 壳文件不得继续保留为长期正式入口
- **AND** `module-registry` 的 mock/real switch 与 L3/L4 compat 导出必须删除
- **AND** 前端正式调用不得再请求 `/api/modules/{moduleId}/bindings/products` 与 `/api/modules/{moduleId}/bindings/repositories`

### Requirement: phase07-10 验收必须证明前端主线已切换

系统 SHALL 以“生成客户端 + 切片 owner + 单一 `/api` 基址”为前端验收主线，而不是只验证页面仍能跑起来。

#### Scenario: 前端回归验证

- **WHEN** 团队验证 `phase07-10` 结果
- **THEN** 必须验证前端 type-check 与构建通过
- **AND** 必须验证页面、切片与 route 不再直接依赖旧 `api-adapter.ts`
- **AND** 必须验证 canonical 写动作 owner 已逐项核销，过渡位已显式列入允许清单
- **AND** 必须验证 L3/L4 module-centered compat 的前端导出与调用已删除，为 `phase07-11` 留下 endpoint 404 验收基础

## MODIFIED Requirements

### Requirement: `phase07` 前端阶段职责划分

`phase07-10` SHALL 负责前端 generated client、read owner、mutation owner、旧 adapter 与 module-centered compat facade 的切换与退场；不得把后端 Connect handler、生成链重建或根级收口重新并入当前阶段。

#### Scenario: 阶段边界

- **WHEN** 团队执行 `phase07-10`
- **THEN** 当前阶段完成条件必须聚焦于前端调用主线切换、owner 收口与 adapter 退场
- **AND** 不得把 `phase07-09` 已完成的后端切换再次作为当前阶段主要交付物
- **AND** 不得把 `phase07-12` 根级同步提前混入当前阶段 DoD

### Requirement: `SovereigntyPanel` 写动作承接方式

`SovereigntyPanel` 的 `ExportCoreAssets` 与 `CreateInstanceBackup` SHALL 在 `phase07-10` 作为显式过渡位保留在组件内，但只允许替换为 generated client 调用，不得继续依赖 hand-written fetch / JSON adapter。

#### Scenario: Dashboard 过渡位处理

- **WHEN** 团队迁移 Dashboard SovereigntyPanel
- **THEN** 组件内的两个 `useMutation` 可以暂留
- **AND** 它们必须改为消费 Connect client 或切片内等价 transport adapter
- **AND** 必须在验收与 checklist 中明确其为允许过渡位，而不是新的长期正式模式

## REMOVED Requirements

### Requirement: hand-written fetch / JSON adapter 继续作为前端正式主线

**Reason**: `phase07-07 formal spec` 已冻结 `.proto + ConnectRPC` 为唯一长期合同与传输主线；前端继续保留 `api-adapter.ts` 正式主线会形成与 generated client 并列的第二套调用体系。

**Migration**: 以共享 Connect transport + slice-local client + `data/` read owner + `application/` mutation owner 替代旧 adapter，并删除对应文件与导出。

### Requirement: module-centered compat facade 可在前端长期保留

**Reason**: `phase07-03` 与 `phase07-07` 已冻结 L3/L4 module-centered compat 入口最晚在 `phase07-10` 退场；前端若继续保留 `module-registry` compat 导出与 switch，会阻断 `phase07-11` 的 legacy endpoint 核销。

**Migration**: 删除 `module-registry` 的 compat 导出、adapter switch 与真实调用分支；将正式调用统一回收到 Product Registry / Repository Binding 的 canonical owner。
