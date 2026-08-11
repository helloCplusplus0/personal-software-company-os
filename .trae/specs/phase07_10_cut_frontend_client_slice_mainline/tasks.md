# Tasks

- [x] Task 1: 冻结 `phase07-10` 的前端迁移 inventory 与阶段边界，使实现范围只覆盖 generated client、切片 owner、adapter 退场与 compat facade 核销。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase07_transport_contract_mainline_migration_dev_plan.md#L208-224` 的范围与 DoD。
  - [x] SubTask 1.2: 对齐 `phase07-07 formal spec` 中前端 34 条 RPC owner、8 个旧 adapter 文件与 L3/L4 退场时点。
  - [x] SubTask 1.3: 以当前源码盘点实际遗留项：页面内 `useQuery/useMutation`、route/CTA 直连读取、既有 `application/use-create-draft-*.ts`、`module-registry` mock/real switch。

- [x] Task 2: 建立前端单一 Connect transport 与 slice-local generated client 承接位。
  - [x] SubTask 2.1: 设计共享 Connect transport 唯一落点，保持浏览器侧单一 `/api` 基址。
  - [x] SubTask 2.2: 设计 7 个业务切片与 Dashboard 复用场景下的 `connect-client.ts` 落点，不新增第二套跨切片 transport / client 宿主。
  - [x] SubTask 2.3: 明确 `frontend/vite.config.ts`、运行时依赖与 generated proto 产物的消费方式。

- [x] Task 3: 回收 canonical 读取主线到切片 `data/` 只读 owner。
  - [x] SubTask 3.1: 回收 Module Registry / Decision Center / Product Registry / Repository Binding 的列表、详情与候选读取。
  - [x] SubTask 3.2: 回收 Dashboard / Onboarding / Reuse Summary / Export / Backup 的读取承接位。
  - [x] SubTask 3.3: 明确 `src/routes/index.tsx` 与 `dashboard/components/onboarding-cta-button.tsx` 共享 Onboarding 只读 helper 的方案。

- [x] Task 4: 收口 canonical 写动作到单一正式 owner，并标注允许过渡位。
  - [x] SubTask 4.1: 保留 `use-create-draft-module / decision / product / repository` 的 owner 身份，仅替换 transport。
  - [x] SubTask 4.2: 为 `CreateRelease / BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository / LinkDecisionToTarget` 设计切片 `application/` 下的正式 owner。
  - [x] SubTask 4.3: 明确 `SovereigntyPanel` 内 `ExportCoreAssets / CreateInstanceBackup` 的过渡位规则、退场条件与验收口径。

- [x] Task 5: 规划旧 adapter 与 module-centered compat facade 的退场顺序。
  - [x] SubTask 5.1: 明确 8 个旧 adapter 文件与对应 adapter 壳的删除顺序。
  - [x] SubTask 5.2: 明确 `module-registry` mock/real switch、compat 导出与 L3/L4 旧路径调用的删除证据。
  - [x] SubTask 5.3: 明确不得保留第二套长期 fetch / JSON 主线，也不得让页面/组件继续直接 import 旧 adapter。

- [x] Task 6: 冻结 `phase07-10` 的前端回归与验收证据。
  - [x] SubTask 6.1: 明确 type-check、构建与关键页面行为等价的验证要求。
  - [x] SubTask 6.2: 明确"无旧 adapter 直接 import / 无页面级长期正式 queryFn / 无组件级长期正式 canonical mutation" 的静态核销要求。
  - [x] SubTask 6.3: 明确 L3/L4 compat 前端导出删除完成后，为 `phase07-11` 留下 `/api/modules/{moduleId}/bindings/*` 返回 404 的验收基础。

- [x] Task 7: 完成 `phase07-10` 规格一致性校验。
  - [x] SubTask 7.1: 验证本规格继承 `phase07-05/06/07/09` 的关键冻结口径。
  - [x] SubTask 7.2: 验证本规格与当前前端源码现状对齐，不回退到过时的 adapter / owner 盘点。
  - [x] SubTask 7.3: 验证本规格未把根级收口、后端重构或未建立阶段内容并入当前阶段。

# Task Dependencies

- `Task 2` depends on `Task 1` ✅
- `Task 3` depends on `Task 1`, `Task 2` ✅
- `Task 4` depends on `Task 1`, `Task 2` ✅
- `Task 5` depends on `Task 1`, `Task 3`, `Task 4` ✅
- `Task 6` depends on `Task 3`, `Task 4`, `Task 5` ✅
- `Task 7` depends on `Task 1` through `Task 6` ✅

# 实现总结

## 变更文件清单

### 新增文件
- `frontend/src/shared/rpc/connect-transport.ts` — 共享 Connect transport
- `frontend/src/features/dashboard/data/connect-client.ts` — Dashboard/Export/Backup client
- `frontend/src/features/decision-center/data/connect-client.ts` — Decision Center client
- `frontend/src/features/module-registry/data/connect-client.ts` — Module Registry client
- `frontend/src/features/onboarding/data/connect-client.ts` — Onboarding client
- `frontend/src/features/product-registry/data/connect-client.ts` — Product Registry client
- `frontend/src/features/repository-binding/data/connect-client.ts` — Repository Binding client
- `frontend/src/features/reuse-summary/data/connect-client.ts` — Reuse Summary client

### 新增 read owner（19 个）
- `use-dashboard-overview-read.ts` / `use-feedback-signals-read.ts` / `use-recent-activities-read.ts`
- `use-export-snapshot-read.ts` / `use-backup-snapshot-read.ts`
- `use-module-list-read.ts` / `use-module-detail-read.ts`
- `use-decision-list-read.ts` / `use-decision-detail-read.ts` / `use-decision-module-candidates-read.ts`
- `use-product-list-read.ts` / `use-product-detail-read.ts` / `use-product-module-candidates-read.ts`
- `use-repository-list-read.ts` / `use-repository-detail-read.ts` / `use-repository-module-candidates-read.ts` / `use-repository-product-candidates-read.ts`
- `use-onboarding-read.ts` / `use-reuse-summary-read.ts`

### 新增 mutation owner（5 个）
- `use-create-release.ts` / `use-bind-module-to-product.ts` / `use-bind-repository-to-product.ts` / `use-map-module-to-repository.ts` / `use-link-decision-to-target.ts`

### 修改文件
- 4 个既有 draft owner：替换 transport 为 Connect client
- 17 个页面/组件：消费新的 read/mutation owner
- `tsconfig.app.json`：`erasableSyntaxOnly` 改为 `false`

### 删除文件（8 个旧 adapter）
- `dashboard/data/api-adapter.ts` / `dashboard/data/sovereignty-api-adapter.ts`
- `onboarding/data/api-adapter.ts` / `reuse-summary/data/api-adapter.ts`
- `decision-center/data/api-adapter.ts` / `product-registry/data/api-adapter.ts`
- `repository-binding/data/api-adapter.ts` / `module-registry/data/api-adapter.ts`

## 验收证据

- [x] `(cd frontend && tsc -b --noEmit)` 通过
- [x] `(cd frontend && npm run build)` 通过
- [x] 无旧 `api-adapter.ts` 文件残留
- [x] 无旧 `api-adapter.ts` import 引用
- [x] 无 L3/L4 `bindings/products` / `bindings/repositories` 旧路径引用
- [x] 7 个切片 connect-client 已建立
- [x] 19 个 read owner 已回收
- [x] 9 个 mutation owner 已收口
- [x] SovereigntyPanel Export/Backup 已切换 Connect client（过渡位保留）