# Phase07-05 前端生成客户端、切片承接位与 query/application 迁移设计 Spec

> **执行产出**：`design.md` — 包含 6 个设计区域：generated client 组合与物理落点（共享 transport + 9 切片 client 表）、query 层迁移策略（page / component / route caller 盘点 + read owner 落点 + 代码模板）、mutation owner inventory（11 项写动作逐项分类：4 保持 + 5 回收 + 2 过渡）、旧 adapter 回收顺序（13 项 adapter 盘点 + 三层回收顺序 + 各切片回收明细表 + 5 步 compat switch 收口规则）、调用组织保持（分层不变表 + 3 禁止项）、一致性声明（7 项上游对齐）。
> **执行日期**：2026-08-11
> **参考**：Context7 `/connectrpc/connect-es` + `/tanstack/query` 最新文档
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

`phase07-01` 已冻结 34 条业务 RPC 与前端 mutation owner inventory，`phase07-02` 已冻结前端正式客户端组合必须收敛为 `bufbuild/es + @connectrpc/connect + @connectrpc/connect-web + /api`，`phase07-03` 已冻结 compat adapter 的退场窗口，`phase07-04` 已冻结后端 Connect handler 的正式挂载方式。当前前端真实状态仍是“8 份 hand-written `api-adapter.ts` + 少量 `application` owner + 多处 page/component 级 `useMutation`，以及 page/component/route 级直接读取 caller 并存”，如果不先把 generated client 承接位、query/application 切换策略、mutation owner 收口与旧 adapter 回收顺序冻结清楚，`phase07-07/10` 很容易长出第二套调用组织或留下未声明的临时过渡位。

## What Changes

- 冻结前端 Connect generated client 的共享 transport 与切片承接位
- 冻结 `query` 层从 hand-written `fetch + JSON` adapter 迁移到 generated client 的切换策略
- 冻结 route 级 first-run 读取与 component 级 candidate 读取的正式 read owner 承接位
- 冻结 `application` 层正式写动作的 mutation owner inventory、固定落点与允许过渡位
- 冻结旧 `api-adapter.ts` / `*-adapter.ts` 的分批回收顺序
- 冻结 `module-registry` 兼容 adapter switch 的退场规则
- **BREAKING**：前端不得同时长期保留“generated client 主线 + hand-written `api-adapter` 主线”并列正式状态

## Impact

- Affected specs:
  - `phase07-01` 前端页面/动作 owner 清单
  - `phase07-02` 前端客户端正式组合
  - `phase07-03` compat 入口退场窗口
  - `phase07-04` 后端 Connect path 与 `/api` 挂载方式
  - `phase07-07` 前端 Connect client 实现
  - `phase07-10` 前端 adapter 回收与 mutation 收口
- Affected code:
  - `frontend/package.json`
  - `frontend/vite.config.ts`
  - `frontend/src/shared/rpc/`
  - `frontend/src/features/*/data/`
  - `frontend/src/features/*/application/`
  - `frontend/src/features/*/pages/`
  - `frontend/src/features/*/components/`
  - `frontend/src/routes/`
  - `frontend/src/gen/proto/`

## ADDED Requirements

### Requirement: 前端必须通过单一共享 Connect transport + 切片内 generated client 承接业务调用

系统 SHALL 将前端正式 RPC 运行时冻结为：

- 共享 transport：`@connectrpc/connect-web` 的 `createConnectTransport({ baseUrl: '/api' })`
- 共享 client 工具：`@connectrpc/connect` 的 `createClient()`
- 单一浏览器业务前缀：`/api`
- service descriptor 来源：`frontend/src/gen/proto/**` 的 `bufbuild/es` 产物

正式落点 SHALL 满足：

- 共享 transport 仅允许存在一个稳定 cross-slice 承接位，例如 `frontend/src/shared/rpc/connect-transport.ts`
- 各业务切片各自持有 slice-local generated client 承接位，例如 `features/<slice>/data/connect-client.ts`
- 页面、组件、表单不得直接在 UI 文件里自行 `createConnectTransport()` 或 `createClient()`
- 不得再新增第二套 hand-written fetch 封装作为长期正式 transport

#### Scenario: Slice query owner consumes generated client

- **WHEN** 某个业务切片迁移读取逻辑到 Connect
- **THEN** 该切片必须通过自己的 `data/connect-client.ts` 或等价 slice-local client 承接位调用 generated service descriptor
- **AND** transport 只能复用共享 `/api` baseUrl 配置

#### Scenario: UI does not create a second transport stack

- **WHEN** 页面或组件需要发起 RPC
- **THEN** 它只能通过切片 `query` owner 或 `application` owner 间接调用 generated client
- **AND** 不得在 UI 文件中直接初始化 transport / client

### Requirement: query 层必须保持纯只读，并以 slice-local read owner 承接 Connect 读取

系统 SHALL 把 `query` 层正式形态冻结为“slice-local read owner + TanStack Query 缓存键 + Connect client 读调用”。

正式要求：

- `query` 层只承接读取、缓存键、请求参数解包与响应解包
- `query` owner 继续优先放在 `features/<slice>/data/use-*-read.ts` 或等价 slice-local read 文件
- route 级非 React caller 若无法直接消费 hook，必须复用切片 `data/` 中同一正式 read owner 导出的只读 helper，不得回退为直接 import `api-adapter.ts`
- 页面可以消费 `useXxxRead()`，但不得重新拼装同一读接口的 queryKey / queryFn 第二份实现
- 组件级候选读取（如 `*Candidates`）仍属于正式 query owner 范围，不得因为与 mutation 同屏出现而长期内联在组件中
- hand-written `fetchXxx` 可以在迁移阶段短时存在，但只允许作为 generated client 切换过程的过渡内层，不得继续成为页面直接 import 的正式 owner

#### Scenario: Existing page-level useQuery is recovered to read owner

- **WHEN** 当前页面直接写有 `useQuery({ queryKey, queryFn: fetchXxx })`
- **THEN** phase07 迁移后应优先回收到该切片的 `useXxxRead()` 或等价 read owner
- **AND** queryKey、enabled、响应解包不得在多个页面重复维护

#### Scenario: Route-level read caller reuses slice read contract

- **WHEN** 某个 route `beforeLoad`、loader 或其他非 React caller 需要读取业务数据
- **THEN** 它必须复用对应切片 `data/` 中由 read owner 导出的只读 helper / query contract
- **AND** 不得直接 import `api-adapter.ts` 形成第二套正式读取主线

#### Scenario: Candidate reads are recovered out of components

- **WHEN** 某个组件同时承接 candidate 读取与 mutation 提交
- **THEN** candidate 读取也必须回收到切片 `data/` 下的 `use-*-candidates-read.ts` 或等价 read owner
- **AND** 组件只保留展开状态、选中态与提交交互，不继续持有 hand-written candidate fetch 主线

#### Scenario: Query layer remains write-free

- **WHEN** 某个切片完成 transport 迁移
- **THEN** 其 `data/` 读层不得混入 `create / bind / link / export / backup` 等写动作
- **AND** 写动作只能进入 `application` owner 或明确允许的短时过渡位

### Requirement: application 层必须成为正式写动作的唯一前端 owner，过渡位必须显式标记

系统 SHALL 将前端正式写动作分成两类：

- canonical application owner：必须进入 `features/<slice>/application/`
- 短时过渡位：允许在切片内固定 page/component 位置存在，但必须在 spec 中显式声明最晚退场时点

本阶段 canonical 写动作至少覆盖：

- `CreateProduct`
- `CreateRepository`
- `CreateModule`
- `CreateDecision`
- `CreateRelease`
- `BindModuleToProduct`
- `BindRepositoryToProduct`
- `MapModuleToRepository`
- `LinkDecisionToTarget`
- `ExportCoreAssets`
- `CreateInstanceBackup`

当前 mutation owner inventory SHALL 冻结为：

- 已在 `application` owner：`useCreateDraftProduct`、`useCreateDraftRepository`、`useCreateDraftModule`、`useCreateDraftDecision`
- 必须回收到切片内固定承接位：`ReleaseCreatePage`、`ProductModuleBindingPanel`、`RepositoryProductBindingPanel`、`RepositoryModuleMappingPanel`、`DecisionModuleCandidatePanel`
- 允许短时过渡：`SovereigntyPanel` 内的 `ExportCoreAssets` 与 `CreateInstanceBackup`，但必须标记为 `phase07-10` 前显式核销的过渡位

#### Scenario: Component-level mutation is recovered to slice owner

- **WHEN** 某个页面或组件当前内联 `useMutation`
- **THEN** 若它是高频或 canonical 业务写动作，必须回收到对应切片 `application/` 内的固定 owner
- **AND** 页面/组件只保留 UI 状态、参数采集、成功回流与展示逻辑

#### Scenario: Low-frequency sovereignty action stays transitional

- **WHEN** `SovereigntyPanel` 仍在 phase07 中直接承接 Export / Backup mutation
- **THEN** 该位置必须被显式标记为允许的短时过渡位
- **AND** 不能被误判为长期正式 mutation 主线

### Requirement: 旧 adapter 回收必须按三层顺序进行，禁止倒序删除

系统 SHALL 将旧前端 adapter 回收顺序冻结为：

1. 先建立 generated client 与 slice-local query/application owner
2. 再切走真实页面 / 组件 caller
3. 最后删除旧 `api-adapter.ts` / `*-adapter.ts` / compat switch 分支

细化要求：

- `module-registry/data/module-registry-adapter.ts` 是唯一仍带 runtime switch 的特殊 compat 资产，必须最后收口
- `product-registry-adapter.ts`、`repository-binding-adapter.ts`、`decision-center-adapter.ts` 当前仅 re-export 真 API，迁移后应优先被删除或降级为 generated client re-export 壳
- `dashboard/data/api-adapter.ts`、`dashboard/data/sovereignty-api-adapter.ts`、`onboarding/data/api-adapter.ts`、`reuse-summary/data/api-adapter.ts` 应在对应 read/application owner 切换完成后删除 hand-written fetch 实现
- `onboarding/data/api-adapter.ts` 的删除前提不仅包括 `OnboardingPage`，还包括 Dashboard 内 `OnboardingCtaButton` 与根路由 `/` 的 first-run caller 已切离直接 adapter import
- `product / repository / decision` 三个切片的 adapter 删除前提必须包含 candidate read caller 已切到切片 read owner，而不只是 mutation owner 已建立
- 未切走 caller 前不得先删 adapter，避免页面临时长出第三套调用方式

#### Scenario: Adapter retirement follows caller migration

- **WHEN** 某个 adapter 对应的页面 / 组件 caller 尚未切到 generated client owner
- **THEN** 该 adapter 不得先行删除
- **AND** 迁移必须先完成 caller 切换，再删除旧实现

#### Scenario: Module registry compat switch retires last

- **WHEN** `module-registry` 切片仍存在 compat candidate / bind / map 的 dormant export
- **THEN** `module-registry-adapter.ts` 中的 runtime switch 不能提前误删为“已无风险”
- **AND** 必须等 `phase07-03` 定义的 4 条 compat 资产一并核销后再收口

### Requirement: 前端调用组织不得因 transport 迁移长出第二套分层

系统 SHALL 保持前端调用组织为：

- `shared/rpc/`：仅承接稳定 cross-slice Connect transport / 通用 client 工具
- `features/<slice>/data/`：slice-local read owner、generated client 承接位、只读参数与缓存键
- `features/<slice>/application/`：正式写动作 owner、失效策略、错误归一化、成功回流协助
- `pages/components`：仅消费 owner，不定义第二套正式 transport 语义

系统 SHALL NOT：

- 新增并列的 `services/`、`sdk/`、`clients/` 根级体系与现有 `features/*/data|application` 长期并列
- 在 `query` 与 `application` 之外再生长一个长期 “api facade” 层
- 把 generated client 直接散落到各个页面和组件

#### Scenario: Shared promotion stays minimal

- **WHEN** 某个能力只在单一切片内使用
- **THEN** 它应保留在切片内，不得过早提升到 `shared`
- **AND** 只有 Connect transport 这类稳定 cross-slice 能力允许进入共享层

## MODIFIED Requirements

### Requirement: phase06-07 的 application owner 基线扩展到 transport mainline 迁移

`phase06-07` 已冻结四类 create draft 必须收敛到 `application` owner。

自 `phase07-05` 起，系统必须把这一 requirement 修改为：

- 不只 create draft，所有 canonical 写动作都要给出 `application owner / 固定过渡位 / 最晚退场时点`
- `application` owner 内部 transport 必须从 hand-written fetch 迁移到 generated Connect client
- 页面 / 组件级 `useMutation` 只有在被 spec 明确声明为过渡位时才允许短时存在

#### Scenario: Existing draft application owners become Connect-ready

- **WHEN** 团队进入 `phase07-07`
- **THEN** 既有 `useCreateDraft*` hooks 必须继续保留为正式 owner
- **AND** 仅替换其内部 transport 为 generated client，不得重新把 create 写回页面层

### Requirement: phase07-01 的 mutation owner inventory 细化为前端迁移执行表

`phase07-01` 已列出当前 mutation owner inventory。

自 `phase07-05` 起，系统必须把这一 requirement 修改为：

- 每个 mutation 都要明确“回收到 application owner / 保留在切片内固定承接位 / 允许短时过渡”
- 每个旧 adapter 都要明确对应 caller、切换前提与删除顺序
- 前端 read/write 切换必须能直接指导 `phase07-07/10` 实现与核销

#### Scenario: Inventory becomes implementation-ready

- **WHEN** 团队开始编写前端 Connect client
- **THEN** 可以直接从本 spec 读出每个 slice 的 client 落点、query owner、application owner 和 adapter 回收顺序

## REMOVED Requirements

### Requirement: 页面可长期直接 import `api-adapter.ts` 作为正式业务调用主线

**Reason**: 这会让 generated client 只是新增 transport 旁路，而不是正式主线，也会让 query/application 边界在迁移后继续模糊。

**Migration**:

- 读调用回收到 slice-local read owner
- 写调用回收到 `application` owner 或明确允许的短时过渡位
- 旧 `api-adapter.ts` 在 caller 切走后按顺序删除
