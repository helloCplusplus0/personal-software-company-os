# Phase07-05 设计产出：前端生成客户端、切片承接位与 query/application 迁移设计

> 本文档是 phase07-05 spec 的执行产出，定义前端 Connect generated client 承接位、query/application 迁移策略、mutation owner inventory 收口与旧 adapter 回收顺序。
> 产出日期：2026-08-11
> 上游：`phase07-01 frozen_scope.md`（mutation owner inventory）、`phase07-02 frozen_scope.md`（前端客户端正式组合）、`phase07-03 frozen_scope.md`（compat 退场窗口）、`phase07-04 design.md`（Connect procedure path）
> 参考：Context7 `/connectrpc/connect-es` + `/tanstack/query` 最新文档

---

## 1. 设计：前端 Generated Client 正式组合方式与物理落点

### 1.1 整体架构

```
┌──────────────────────────────────────────────────────────────────┐
│  前端调用组织（迁移后正式形态）                                      │
│                                                                    │
│  shared/rpc/                                                       │
│  ├── connect-transport.ts        ← 唯一 cross-slice transport      │
│  │   createConnectTransport({ baseUrl: '/api' })                   │
│  └── (禁止新增第二套 transport 封装)                                │
│                                                                    │
│  features/<slice>/data/                                            │
│  ├── connect-client.ts           ← slice-local generated client  │
│  │   createClient(ServiceDescriptor, transport)                   │
│  ├── use-*-read.ts               ← slice-local read owner        │
│  │   useQuery({ queryKey, queryFn: client.listXxx })             │
│  └── (读层禁止混入 create/bind/link/export/backup)                │
│                                                                    │
│  features/<slice>/application/                                     │
│  └── use-*-draft.ts              ← slice-local mutation owner    │
│      useMutation({ mutationFn: client.createXxx })               │
│      (负责失效刷新、成功回流、错误归一化)                            │
│                                                                    │
│  pages / components                                               │
│  └── 消费 useXxxRead() / useXxxDraft()                            │
│      (不得直接 import transport / client / api-adapter)           │
└──────────────────────────────────────────────────────────────────┘
```

### 1.2 共享 Transport 唯一落点

```typescript
// frontend/src/shared/rpc/connect-transport.ts

import { createConnectTransport } from '@connectrpc/connect-web';

/**
 * 共享 Connect transport — 前端唯一跨切片 transport 承接位。
 *
 * phase07-02 §4.3 冻结：baseUrl 为 '/api'，与 Vite dev proxy、Caddy 统一。
 * 页面、组件、切片 query/application owner 均通过此 transport 发起 RPC。
 */
export const transport = createConnectTransport({
  baseUrl: '/api',
});
```

### 1.3 各切片 Generated Client 承接位

| 切片 | Slice-local client 文件 | 生成的 Service Descriptor 来源 |
|------|------------------------|-------------------------------|
| Module Registry | `features/module-registry/data/connect-client.ts` | `@/gen/proto/psco/module_registry/v1/module_registry_pb` |
| Decision Center | `features/decision-center/data/connect-client.ts` | `@/gen/proto/psco/decision_center/v1/decision_center_pb` |
| Product Registry | `features/product-registry/data/connect-client.ts` | `@/gen/proto/psco/product_registry/v1/product_registry_pb` |
| Repository Binding | `features/repository-binding/data/connect-client.ts` | `@/gen/proto/psco/repository_binding/v1/repository_binding_pb` |
| Dashboard | `features/dashboard/data/connect-client.ts` | `@/gen/proto/psco/dashboard/v1/dashboard_pb` |
| Onboarding | `features/onboarding/data/connect-client.ts` | `@/gen/proto/psco/onboarding/v1/onboarding_pb` |
| Export | `features/dashboard/data/connect-client.ts` | `@/gen/proto/psco/export/v1/export_pb` |
| Backup | `features/dashboard/data/connect-client.ts` | `@/gen/proto/psco/backup/v1/backup_pb` |
| Reuse Summary | `features/reuse-summary/data/connect-client.ts` | `@/gen/proto/psco/reuse_summary/v1/reuse_summary_pb` |

> **注**：Export / Backup 的 client 与 Dashboard 共享 `features/dashboard/data/connect-client.ts`，因为它们都服务于 Dashboard 页面（SovereigntyPanel）。

### 1.4 Slice-local Client 模板

```typescript
// features/module-registry/data/connect-client.ts

import { createClient } from '@connectrpc/connect';
import { ModuleRegistryService } from '@/gen/proto/psco/module_registry/v1/module_registry_pb';
import { transport } from '@/shared/rpc/connect-transport';

/**
 * Module Registry 切片内唯一的 Connect client 承接位。
 * 页面、read owner、application owner 均通过此 client 发起 RPC。
 */
export const moduleRegistryClient = createClient(ModuleRegistryService, transport);
```

### 1.5 禁止

- 页面/组件直接 `import { createConnectTransport } from '@connectrpc/connect-web'`
- 页面/组件直接 `import { createClient } from '@connectrpc/connect'`
- 新增第二个共享 transport 文件（如 `shared/rpc/grpc-transport.ts`）
- 在 `shared/` 下新建 `api/`、`sdk/`、`clients/` 等根级目录

---

## 2. 设计：Query 层迁移策略

### 2.1 当前状态盘点（按真实 caller）

| 切片 | 当前 Read 模式 | 文件 | 迁移动作 |
|------|---------------|------|---------|
| Module Registry | `useQuery({ queryKey, queryFn: fetchModuleList })` 在 `module-list-page.tsx` | `module-list-page.tsx` | 回收至 `data/use-module-list-read.ts` |
| Module Registry | `useQuery({ queryKey, queryFn: fetchModuleDetail })` 在 page | `module-detail-page.tsx` | 回收至 `data/use-module-detail-read.ts` |
| Decision Center | `useQuery` 在 `decision-list-page.tsx` / `decision-detail-page.tsx` | 两个 page 文件 | 回收至 `data/use-decision-list-read.ts` + `data/use-decision-detail-read.ts` |
| Decision Center | `useQuery({ queryKey, queryFn: fetchDecisionModuleCandidates })` 在组件内 | `components/decision-module-candidate-panel.tsx` | 回收至 `data/use-decision-module-candidates-read.ts` |
| Product Registry | `useQuery` 在 `product-list-page.tsx` / `product-detail-page.tsx` | 两个 page 文件 | 回收至 `data/use-product-list-read.ts` + `data/use-product-detail-read.ts` |
| Product Registry | `useQuery({ queryKey, queryFn: fetchProductModuleCandidates })` 在组件内 | `components/product-module-binding-panel.tsx` | 回收至 `data/use-product-module-candidates-read.ts` |
| Repository Binding | `useQuery` 在 `repository-binding-list-page.tsx` / `repository-binding-detail-page.tsx` | 两个 page 文件 | 回收至 `data/use-repository-list-read.ts` + `data/use-repository-detail-read.ts` |
| Repository Binding | `useQuery({ queryKey, queryFn: fetchRepositoryProductCandidates })` 在组件内 | `components/repository-product-binding-panel.tsx` | 回收至 `data/use-repository-product-candidates-read.ts` |
| Repository Binding | `useQuery({ queryKey, queryFn: fetchRepositoryModuleCandidates })` 在组件内 | `components/repository-module-mapping-panel.tsx` | 回收至 `data/use-repository-module-candidates-read.ts` |
| Dashboard | `useQuery` 在 `dashboard-home-page.tsx` | 页面内 | 回收至 `data/use-dashboard-overview-read.ts` + `data/use-feedback-signals-read.ts` + `data/use-recent-activities-read.ts` |
| Onboarding | ✅ 已有 `data/use-onboarding-read.ts` | `data/use-onboarding-read.ts` | 保留 `use-onboarding-read.ts`，并补 route / CTA 可复用的只读 helper |
| Onboarding | `useQuery({ queryKey, queryFn: fetchFirstRunState })` 在 Dashboard CTA 内 | `dashboard/components/onboarding-cta-button.tsx` | 改为消费 `useOnboardingRead()` 或同一 read owner 导出的只读 helper |
| Onboarding | 根路由 `/` 在 `beforeLoad` 直接调用 `fetchFirstRunState()` | `src/routes/index.tsx` | 改为复用 Onboarding slice 的只读 helper，不再直连 `api-adapter.ts` |
| Reuse Summary | ✅ 已有 `data/use-reuse-summary-read.ts` | `data/use-reuse-summary-read.ts` | 仅替换内部 transport 为 generated client |

### 2.2 迁移后 Read Owner 落点

```
features/<slice>/data/
├── connect-client.ts                        # slice-local client
├── use-<entity>-list-read.ts               # 列表读取 owner
├── use-<entity>-detail-read.ts             # 详情读取 owner
├── use-<entity>-candidates-read.ts         # 候选读取 owner（若需要）
├── read-<entity>.ts / <entity>-query.ts    # route / 非 React caller 复用的只读 helper（若需要）
└── (旧 api-adapter.ts 在 caller 切走后删除)
```

Onboarding 特例：

```
features/onboarding/data/
├── connect-client.ts
├── read-first-run-state.ts      # 根路由 / Dashboard CTA 复用的只读 helper
└── use-onboarding-read.ts       # React 查询 owner，内部复用同一 helper + queryKey
```

### 2.3 Read Owner 模板

```typescript
// features/module-registry/data/use-module-list-read.ts

import { useQuery } from '@tanstack/react-query';
import { moduleRegistryClient } from './connect-client';

/**
 * ModuleListRead — slice-local read owner。
 *
 * 职责：
 *   - 唯一 queryKey 与 queryFn 定义
 *   - 请求参数解包（queryText, statusFilter）
 *   - 响应解包（proto → domain DTO）
 *   - 零计数空态处理
 *
 * 禁止：
 *   - 混入 create / bind / link 等写动作
 *   - 在页面中重新定义 queryKey / queryFn
 */
export function useModuleListRead(queryText: string, statusFilter: string) {
  return useQuery({
    queryKey: ['modules', 'list', queryText, statusFilter],
    queryFn: async () => {
      const res = await moduleRegistryClient.listModules({
        queryText,
        statusFilter: statusFilter as any, // proto enum
      });
      return res.modules ?? [];
    },
    enabled: true,
  });
}
```

Onboarding route / CTA 复用模板：

```typescript
// features/onboarding/data/read-first-run-state.ts

import { onboardingClient } from './connect-client';
import { useQuery } from '@tanstack/react-query';

export const ONBOARDING_STATE_QUERY_KEY = ['onboarding-state'] as const;

/**
 * readFirstRunState — Onboarding slice 的唯一 first-run 读取 helper。
 *
 * 供 route beforeLoad、Dashboard CTA 与 useOnboardingRead() 共同复用，
 * 避免 `fetchFirstRunState()` 在 route / component 中再次长出第二条正式主线。
 */
export async function readFirstRunState() {
  return onboardingClient.getFirstRunState({});
}

export function useOnboardingRead() {
  return useQuery({
    queryKey: ONBOARDING_STATE_QUERY_KEY,
    queryFn: readFirstRunState,
  });
}
```

### 2.4 Query 层约束

| 规则 | 说明 |
|------|------|
| 纯只读 | query 层只承接 `useQuery`，不混入 `useMutation` |
| 单一 owner | 同一读接口的 queryKey / queryFn 只在 slice-local read owner 中定义一次 |
| 页面消费 | 页面通过 `useXxxRead()` 消费，不重新拼装 queryKey/queryFn |
| 组件候选读取 | candidate reads 也必须进入 `use-*-candidates-read.ts`，不得长期留在组件内 |
| route 复用 | route / beforeLoad 等非 React caller 必须复用切片 `data/` 内同一只读 helper，不得直连 `api-adapter.ts` |
| 缓存键一致 | queryKey 模式与 TanStack Query 缓存失效策略对齐 |

---

## 3. 设计：Mutation Owner Inventory 与收口策略

### 3.1 Canonical 写动作完整清单（11 项）

| # | Canonical 写动作 | 当前落点 | 当前 Transport | 迁移后 Owner | Owner 类型 |
|---|-----------------|---------|---------------|-------------|-----------|
| 1 | **CreateProduct** | `application/use-create-draft-product.ts` | hand-written fetch | 同文件，替换 transport | ✅ Application Owner |
| 2 | **CreateRepository** | `application/use-create-draft-repository.ts` | hand-written fetch | 同文件，替换 transport | ✅ Application Owner |
| 3 | **CreateModule** | `application/use-create-draft-module.ts` | hand-written fetch | 同文件，替换 transport | ✅ Application Owner |
| 4 | **CreateDecision** | `application/use-create-draft-decision.ts` | hand-written fetch | 同文件，替换 transport | ✅ Application Owner |
| 5 | **CreateRelease** | `pages/release-create-page.tsx`（页面内 useMutation） | hand-written fetch | `application/use-create-release.ts`（**新增**） | 🔄 回收到 Application Owner |
| 6 | **BindModuleToProduct** | `components/product-module-binding-panel.tsx`（组件内 useMutation） | hand-written fetch | `application/use-bind-module-to-product.ts`（**新增**） | 🔄 回收到 Application Owner |
| 7 | **BindRepositoryToProduct** | `components/repository-product-binding-panel.tsx`（组件内 useMutation） | hand-written fetch | `application/use-bind-repository-to-product.ts`（**新增**） | 🔄 回收到 Application Owner |
| 8 | **MapModuleToRepository** | `components/repository-module-mapping-panel.tsx`（组件内 useMutation） | hand-written fetch | `application/use-map-module-to-repository.ts`（**新增**） | 🔄 回收到 Application Owner |
| 9 | **LinkDecisionToTarget** | `components/decision-module-candidate-panel.tsx`（组件内 useMutation） | hand-written fetch | `application/use-link-decision-to-target.ts`（**新增**） | 🔄 回收到 Application Owner |
| 10 | **ExportCoreAssets** | `components/sovereignty-panel.tsx`（组件内 useMutation） | hand-written fetch | 同组件内，mark 为过渡位 | ⏸️ 短时过渡位 |
| 11 | **CreateInstanceBackup** | `components/sovereignty-panel.tsx`（组件内 useMutation） | hand-written fetch | 同组件内，mark 为过渡位 | ⏸️ 短时过渡位 |

### 3.2 新增 Application Owner 文件清单

| 新增文件 | 承接 Mutation | 所属切片 | 最晚交付时点 |
|---------|-------------|---------|------------|
| `module-registry/application/use-create-release.ts` | `CreateRelease` | Module Registry | phase07-07 |
| `product-registry/application/use-bind-module-to-product.ts` | `BindModuleToProduct` | Product Registry | phase07-07 |
| `repository-binding/application/use-bind-repository-to-product.ts` | `BindRepositoryToProduct` | Repository Binding | phase07-07 |
| `repository-binding/application/use-map-module-to-repository.ts` | `MapModuleToRepository` | Repository Binding | phase07-07 |
| `decision-center/application/use-link-decision-to-target.ts` | `LinkDecisionToTarget` | Decision Center | phase07-07 |

### 3.3 短时过渡位规则

| 过渡位 | 允许原因 | 最晚退场时点 | 退场条件 |
|--------|---------|------------|---------|
| `SovereigntyPanel` 内 `ExportCoreAssets` | 低频操作，Dashboard 页面内 SovereigntyPanel 是其唯一触发点 | phase07-10 | 1) 内部 transport 切换为 generated client；2) 显式声明为"允许过渡位"；3) 不建立独立 application owner |
| `SovereigntyPanel` 内 `CreateInstanceBackup` | 同上 | phase07-10 | 同上 |

### 3.4 Application Owner 模板（写动作回收后）

```typescript
// features/product-registry/application/use-bind-module-to-product.ts

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { productRegistryClient } from '../data/connect-client';

/**
 * BindModuleToProduct — Product Registry 切片内固定 mutation 承接位。
 *
 * 职责：
 *   - 唯一 mutationFn 定义
 *   - 失效刷新：product detail + product list + module detail（跨切片失效）
 *   - 成功回流：返回结果供调用方消费
 *   - 错误归一化：Connect error 转为 UI 可消费错误
 */
export function useBindModuleToProduct() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (input: { productId: string; moduleId: string }) =>
      productRegistryClient.bindModuleToProduct({
        productId: input.productId,
        moduleId: input.moduleId,
      }),
    onSuccess: (_data, variables) => {
      // 失效 product detail
      queryClient.invalidateQueries({ queryKey: ['products', 'detail', variables.productId] });
      // 失效 product list
      queryClient.invalidateQueries({ queryKey: ['products', 'list'] });
      // 跨切片失效 module detail
      queryClient.invalidateQueries({ queryKey: ['modules', 'detail', variables.moduleId] });
    },
  });
}
```

### 3.5 迁移后既有 Application Owner 变化

既有 4 个 `use-create-draft-*` hooks 的迁移：
- **文件位置不变**：继续保留在 `features/<slice>/application/`
- **内部 transport 替换**：`fetch()` → `sliceClient.createXxx()`
- **失效策略不变**：保持现有 `onSuccess` 中的 `queryClient.invalidateQueries`
- **成功回流不变**：保持现有返回值结构

---

## 4. 设计：旧 Adapter 回收顺序

### 4.1 当前 Adapter 资产盘点

| # | 文件 | 当前角色 | 回收优先级 |
|---|------|---------|-----------|
| 1 | `module-registry/data/api-adapter.ts` | 8 个 hand-written fetch 函数（read + write + compat candidates） | 🔴 最后（依赖 compat 退场） |
| 2 | `module-registry/data/module-registry-adapter.ts` | runtime switch：`VITE_USE_REAL_API` 切换 mock/real | 🔴 最后（compat switch） |
| 3 | `module-registry/data/mock-adapter.ts` | Mock 数据（仅 phase02 演示，当前无 active caller） | 🟡 可在 adapter 切换完成后删除 |
| 4 | `decision-center/data/api-adapter.ts` | 5 个 hand-written fetch 函数 | 🟠 中等 |
| 5 | `decision-center/data/decision-center-adapter.ts` | re-export 壳（当前仅 re-export `api-adapter`） | 🟠 中等 |
| 6 | `product-registry/data/api-adapter.ts` | 5 个 hand-written fetch 函数 | 🟠 中等 |
| 7 | `product-registry/data/product-registry-adapter.ts` | re-export 壳 | 🟠 中等 |
| 8 | `repository-binding/data/api-adapter.ts` | 7 个 hand-written fetch 函数 | 🟠 中等 |
| 9 | `repository-binding/data/repository-binding-adapter.ts` | re-export 壳 | 🟠 中等 |
| 10 | `dashboard/data/api-adapter.ts` | 3 个 hand-written fetch 函数（read-only） | 🟢 优先 |
| 11 | `dashboard/data/sovereignty-api-adapter.ts` | 4 个 hand-written fetch 函数（Export + Backup） | 🟢 优先 |
| 12 | `onboarding/data/api-adapter.ts` | `useOnboardingRead` + Dashboard CTA + 根路由 `/` 共同依赖的 first-run 读取入口 | 🟢 优先 |
| 13 | `reuse-summary/data/api-adapter.ts` | 1 个 hand-written fetch 函数 | 🟢 优先 |

### 4.2 三层回收顺序

```
第一层：建立 connect-client + read/mutation owner（不删任何旧文件）
    ├── shared/rpc/connect-transport.ts（新建）
    ├── 各切片 data/connect-client.ts（新建）
    ├── 各切片 data/use-*-read.ts（新建或修改）
    └── 各切片 application/use-*.ts（新建 5 个 + 修改 4 个）

第二层：切走真实页面/组件 caller（不删旧 adapter）
    ├── 页面/组件替换 import：旧 api-adapter → 新 read owner / application owner
    ├── route caller 替换 import：旧 api-adapter → 切片 `data/` 中同一只读 helper
    ├── 验证：tsc -b --noEmit 通过，npm run build 通过
    └── 验证：页面功能正常（Connect RPC 调用成功）

第三层：删除旧 adapter（按优先级）
    ├── 🟢 优先：dashboard / onboarding / reuse-summary / sovereignty api-adapter.ts
    ├── 🟠 中等：decision-center / product-registry / repository-binding api-adapter.ts + adapter 壳
    ├── 🟡 低：module-registry mock-adapter.ts
    └── 🔴 最后：module-registry api-adapter.ts + module-registry-adapter.ts（compat switch）
```

### 4.3 各切片回收明细

| 切片 | 第一层（建立） | 第二层（切走 caller） | 第三层（删除旧文件） |
|------|-------------|---------------------|-------------------|
| **Dashboard** | `connect-client.ts` + `use-dashboard-overview-read.ts` + `use-feedback-signals-read.ts` + `use-recent-activities-read.ts` | `dashboard-home-page.tsx` 替换 import | 删除 `dashboard/data/api-adapter.ts` |
| **Dashboard (Sovereignty)** | `connect-client.ts`（复用 Dashboard 的） | `sovereignty-panel.tsx` 替换内部 transport | 删除 `dashboard/data/sovereignty-api-adapter.ts` |
| **Onboarding** | `connect-client.ts` + `read-first-run-state.ts`（或等价只读 helper）+ `use-onboarding-read.ts` | `onboarding-page.tsx` 保留 `useOnboardingRead()`；`dashboard/components/onboarding-cta-button.tsx` 与 `src/routes/index.tsx` 停止直接 import `api-adapter.ts` | 删除 `onboarding/data/api-adapter.ts` |
| **Reuse Summary** | `connect-client.ts` | `use-reuse-summary-read.ts` 替换内部 transport | 删除 `reuse-summary/data/api-adapter.ts` |
| **Decision Center** | `connect-client.ts` + `use-decision-list-read.ts` + `use-decision-detail-read.ts` + `use-decision-module-candidates-read.ts` + `use-link-decision-to-target.ts` | `decision-list-page.tsx` + `decision-detail-page.tsx` + `decision-module-candidate-panel.tsx` 替换 import | 删除 `decision-center/data/api-adapter.ts` + `decision-center-adapter.ts` |
| **Product Registry** | `connect-client.ts` + `use-product-list-read.ts` + `use-product-detail-read.ts` + `use-product-module-candidates-read.ts` + `use-bind-module-to-product.ts` | `product-list-page.tsx` + `product-detail-page.tsx` + `product-module-binding-panel.tsx` 替换 import | 删除 `product-registry/data/api-adapter.ts` + `product-registry-adapter.ts` |
| **Repository Binding** | `connect-client.ts` + `use-repository-list-read.ts` + `use-repository-detail-read.ts` + `use-repository-product-candidates-read.ts` + `use-repository-module-candidates-read.ts` + `use-bind-repository-to-product.ts` + `use-map-module-to-repository.ts` | `repository-binding-list-page.tsx` + `repository-binding-detail-page.tsx` + `repository-product-binding-panel.tsx` + `repository-module-mapping-panel.tsx` 替换 import | 删除 `repository-binding/data/api-adapter.ts` + `repository-binding-adapter.ts` |
| **Module Registry** | `connect-client.ts` + `use-module-list-read.ts` + `use-module-detail-read.ts` + `use-create-release.ts` | `module-list-page.tsx` + `module-detail-page.tsx` + `release-create-page.tsx` 替换 import | 最后删除 `module-registry/data/api-adapter.ts` + `module-registry-adapter.ts` + `mock-adapter.ts` |

### 4.4 Module Registry Compat Switch 特殊规则

`module-registry-adapter.ts` 是唯一仍带 runtime switch 的特殊资产，收口规则：

1. 所有 caller 已切到 generated client 后，`module-registry-adapter.ts` 不再有 active import
2. 确认 `VITE_USE_REAL_API` 环境变量不再被读取
3. 确认 compat candidate 入口（`GET /api/candidates/*`）已按 phase07-03 §2 退场
4. 确认 compat 绑定入口（`POST /api/modules/{id}/bindings/*`）已按 phase07-03 §2 退场
5. 以上 4 项全部满足后，删除 `module-registry-adapter.ts` + `api-adapter.ts` + `mock-adapter.ts`

---

## 5. 设计：前端调用组织保持

### 5.1 分层不变

| 层 | 迁移前 | 迁移后 | 变化 |
|----|--------|--------|------|
| `shared/rpc/` | 不存在 | `connect-transport.ts`（唯一） | 新增 |
| `features/<slice>/data/` | `api-adapter.ts`（hand-written fetch） | `connect-client.ts` + `use-*-read.ts` + `read-*.ts`（按需） | 替换 |
| `features/<slice>/application/` | 4 个 `use-create-draft-*` | 9 个 `use-*` hooks | 新增 5 个 |
| `routes/` | `beforeLoad` 可直接 import `api-adapter` | 复用切片 `data/` 中只读 helper | 收口 |
| `pages/` | 消费 `api-adapter` 直接 import | 消费 `useXxxRead()` / `useXxxDraft()` | 替换 import |
| `components/` | 内联 `useQuery` / `useMutation` | 消费 `useXxxRead()` / `useXxxDraft()` | 替换 import |

### 5.2 禁止

- 新增 `services/`、`sdk/`、`clients/` 根级目录与 `features/*/data|application` 长期并列
- 在 `query` 与 `application` 之外再生长一个长期 "api facade" 层
- 把 generated client 直接散落到各个页面和组件
- 把只在单一切片内使用的能力过早提升到 `shared/`

---

## 6. 与上游文档一致性声明

| 上游文档 | 关键结论 | 本设计对齐 |
|---------|---------|-----------|
| `phase07-01 frozen_scope.md` §6 | 11 个 canonical 写动作 + 8 个 adapter 文件 | §3.1（11 项完整清单）+ §4.1（13 项 adapter 盘点） |
| `phase07-02 frozen_scope.md` §4 | `@connectrpc/connect` + `@connectrpc/connect-web` + `createClient()` + `/api` | §1.2-1.4 共享 transport + slice-local client |
| `phase07-03 frozen_scope.md` §2-4 | compat 退场窗口 + 三层回收顺序 | §4.2 三层回收顺序 + §4.4 compat switch 特殊规则 |
| `phase07-04 design.md` §3 | Connect procedure path 映射 | §1.3 generated client 通过 `bufbuild/es` service descriptor 调用 |
| `phase06-07` spec | 4 个 create draft 收敛到 application owner | §3.5 既有 application owner 保持不变 |
| `project_rules.md` §2.5 | 切片优先、query 纯只读、mutation 固定承接位 | §2.4 query 约束 + §5.1 分层不变 |
| Context7 connect-es | `createConnectTransport` + `createClient` | §1.2-1.4 代码模板对齐最新 API |
| Context7 tanstack/query | 自定义 read hook / 共享 query contract | §2.3 route / CTA 复用模板 |
