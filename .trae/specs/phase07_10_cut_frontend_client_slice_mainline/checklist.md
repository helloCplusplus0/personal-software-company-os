- [x] 已明确 `phase07-10` 只承接前端 generated client、切片 owner、adapter 退场与 compat facade 核销，不回卷后端与根级收口
- [x] 已明确前端共享 Connect transport 的唯一落点与单一 `/api` 基址
- [x] 已明确各业务切片 `connect-client.ts` 或等价承接位的物理落点
- [x] 已明确页面、组件不得直接调用 `createConnectTransport()` / `createClient()`，也不得新增第二套 transport 宿主
- [x] 已明确 Module Registry / Decision Center / Product Registry / Repository Binding 的列表、详情与候选读取回收到切片 `data/` 只读 owner
- [x] 已明确 Dashboard / Onboarding / Reuse Summary / Export / Backup 的读取承接位迁移方案
- [x] 已明确 `src/routes/index.tsx` 与 `dashboard/components/onboarding-cta-button.tsx` 必须复用 Onboarding 切片同一只读 helper
- [x] 已明确 `CreateModule / CreateDecision / CreateProduct / CreateRepository` 保留既有 application owner，仅替换 transport
- [x] 已明确 `CreateRelease / BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository / LinkDecisionToTarget` 的正式 owner 落点
- [x] 已明确 `ExportCoreAssets / CreateInstanceBackup` 在 `SovereigntyPanel` 中的允许过渡位规则与退场条件
- [x] 已明确 8 个旧 adapter 文件及其 adapter 壳的删除顺序
- [x] 已明确 `module-registry` mock/real switch、compat 导出与 L3/L4 旧路径前端调用的删除证据
- [x] 已明确不得保留第二套长期 fetch / JSON 主线
- [x] 已明确前端 type-check、构建、关键页面行为等价与静态核销要求
- [x] 已验证本规格对齐 `phase07-05`、`phase07-06`、`phase07-07`、`phase07-09` 与当前源码现状

## 实现验收

- [x] 前端共享 Connect transport 已建立：`frontend/src/shared/rpc/connect-transport.ts`（baseUrl `/api`）
- [x] 7 个切片 connect-client.ts 已建立：dashboard / decision-center / module-registry / onboarding / product-registry / repository-binding / reuse-summary
- [x] 19 个 read owner 已建立并回收到切片 `data/` 下
- [x] 9 个 mutation owner 已建立（含 4 个既有 draft owner 替换 transport + 5 个新增 owner）
- [x] 8 个旧 adapter 文件已全部删除
- [x] `module-registry` compat facade 已核销，无 L3/L4 旧路径引用
- [x] SovereigntyPanel 内 Export/Backup 已切换为 generated client（过渡位保留）
- [x] `(cd frontend && tsc -b --noEmit)` 通过
- [x] `(cd frontend && npm run build)` 通过
- [x] 无旧 `api-adapter.ts` import 引用
- [x] 无 `bindings/products` / `bindings/repositories` 旧路径引用

## 2026-08-11 最终收口验收

- [x] `src/routes/index.tsx` 已不再直连 `connect-client`，而是与 `dashboard/components/onboarding-cta-button.tsx` 共同复用 `onboarding/data/use-onboarding-read.ts` 中的 `fetchOnboardingRead / useOnboardingRead`
- [x] 精确静态核销：`rg "^import .*api-adapter|^import .*module-registry-adapter|^import .*decision-center-adapter|^import .*product-registry-adapter|^import .*repository-binding-adapter|^import .*mock-adapter|^import .*mock-data" frontend/src` 返回空结果
- [x] 精确静态核销：`rg "bindings/products|bindings/repositories" frontend/src` 返回空结果
- [x] 精确静态核销：`rg "createClient\\(|createConnectTransport\\(" frontend/src` 仅命中 `shared/rpc/connect-transport.ts` 与各切片 `data/connect-client.ts`
- [x] 精确静态核销：`rg "useMutation\\(" frontend/src/features` 仅剩 `dashboard/components/sovereignty-panel.tsx` 中的 2 个显式允许过渡位
- [x] 运行验证：`(cd frontend && npx tsc -b --noEmit)` 通过
- [x] 运行验证：`(cd frontend && npm run build)` 通过
