# Phase06-07 前端写路径收敛与 mutation 承接位设计 Spec

## Why

`phase06-02` 已冻结 draft-first 写模型与唯一 `application` 承接位要求，`phase06-05` 已冻结 `query` 纯只读与 mutation 固定承接位约束，`phase06-06` 已定义 Onboarding 步骤组件对 `useCreateDraftProduct / useCreateDraftRepository / useCreateDraftModule / useCreateDraftDecision` 的消费方式。但如果不在进入实现前把四类核心对象的 application owner 切片落点、统一错误归一化、成功回流与 query 失效策略、旧模式回收清单与迁移顺序设计到可直接落地的程度，后续 `phase06-15` 前端实现仍会在 mutation 拆分粒度、失效范围与回流职责边界上各自猜测，且既有 4 个 create 页面中的 page-level `useMutation` 会继续作为第二套写语义与 Onboarding 并存。

## What Changes

- 产出四类核心对象（Product / Repository / Module / Decision）create draft application owner 的切片落点、职责与接口设计
- 产出四类核心对象 update 写路径的预留落点设计
- 产出统一错误归一化策略与 `NormalizedError` 形状
- 产出统一成功回流与 query 失效策略（失效矩阵 + 回流职责拆分）
- 识别前端当前所有 page-level / panel-level `useMutation` 漂移点
- 列出 phase06 必须回收的 4 个 create 页面清单与对应 application owner
- 产出回收迁移顺序与 phase06 收口标准

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-02` draft-first 写路径
  - `phase06-05` 前端四条约束执行口径
  - `phase06-06` Onboarding 前端页面、路由与交互流设计
- Affected code:
  - 新增 `frontend/src/features/product-registry/application/` 切片
  - 新增 `frontend/src/features/repository-binding/application/` 切片
  - 新增 `frontend/src/features/module-registry/application/` 切片
  - 新增 `frontend/src/features/decision-center/application/` 切片
  - 回收 `frontend/src/features/product-registry/pages/product-create-page.tsx` 中的 page-level `useMutation`
  - 回收 `frontend/src/features/repository-binding/pages/repository-create-page.tsx` 中的 page-level `useMutation`
  - 回收 `frontend/src/features/module-registry/pages/module-create-page.tsx` 中的 page-level `useMutation`
  - 回收 `frontend/src/features/decision-center/pages/decision-create-page.tsx` 中的 page-level `useMutation`
  - 识别但允许过渡保留：4 个 binding/link panel 级 `useMutation` + 1 个 release create page 级 `useMutation`

---

## 1. 当前漂移点全景

### 1.1 Page-level useMutation（必须回收）

| 文件 | 行号 | mutationFn | 失效 queryKey | 成功回流 | 错误处理 |
| --- | --- | --- | --- | --- | --- |
| `product-registry/pages/product-create-page.tsx` | L79 | `createProduct` | `['product-list']` | navigate `/products/$productId` + 来源上下文 | `toast.error` |
| `module-registry/pages/module-create-page.tsx` | L38 | `createModule` | `['module-list']` | navigate `/modules/$moduleId` + dashboard 上下文 | `toast.error` |
| `repository-binding/pages/repository-create-page.tsx` | L103 | `createRepository` | `['repository-list']` | navigate `/repositories/$repositoryId` + 来源上下文 + product transit | `toast.error` |
| `decision-center/pages/decision-create-page.tsx` | L49 | `createDecision` | `['decision-list']` | navigate `/decisions/$decisionId` + fromList | `toast.error` |

当前 4 个 create 页面的共同模式：

```
useMutation({
  mutationFn: createXxx,
  onSuccess: (response) => {
    queryClient.invalidateQueries({ queryKey: ['xxx-list'] })
    toast.success('xxx创建成功')
    navigate({ to: '/xxx/$xxxId', params: { xxxId: response.xxx_id }, search: detailSearch })
  },
  onError: (error) => {
    toast.error('创建失败：' + error.message)
  },
})
```

### 1.2 Panel-level useMutation（识别，允许过渡保留）

| 文件 | 行号 | mutationFn | 失效 queryKey | 成功行为 |
| --- | --- | --- | --- | --- |
| `product-registry/components/product-module-binding-panel.tsx` | L79 | `bindModuleToProduct` | 无显式失效（通过 `onBindingSuccess` 回调） | `toast` + 关闭面板 + 回调 |
| `repository-binding/components/repository-product-binding-panel.tsx` | L83 | `bindProductToRepository` | （类似模式） | （类似模式） |
| `repository-binding/components/repository-module-mapping-panel.tsx` | L83 | `mapModuleToRepository` | （类似模式） | （类似模式） |
| `decision-center/components/decision-module-candidate-panel.tsx` | L52 | `linkDecisionToTarget` | `['decision-detail', id]` + `['decision-module-candidates', id]` + `['decision-list']` | `toast.success` |

### 1.3 Release create page useMutation（识别，允许过渡保留）

| 文件 | 行号 | mutationFn |
| --- | --- | --- |
| `module-registry/pages/release-create-page.tsx` | L41 | `createRelease` |

---

## 2. Application Owner 切片落点设计

### 2.1 目录结构

每个 feature slice 新增 `application/` 目录，承接该切片所有正式 mutation owner：

```
product-registry/
  application/
    index.ts                           # barrel export
    normalize-error.ts                 # 切片级错误归一化工具
    use-create-draft-product.ts        # useCreateDraftProduct
    use-update-product.ts              # useUpdateProduct（预留落点）

repository-binding/
  application/
    index.ts
    normalize-error.ts
    use-create-draft-repository.ts     # useCreateDraftRepository
    use-update-repository.ts           # useUpdateRepository（预留落点）

module-registry/
  application/
    index.ts
    normalize-error.ts
    use-create-draft-module.ts         # useCreateDraftModule
    use-update-module.ts               # useUpdateModule（预留落点）

decision-center/
  application/
    index.ts
    normalize-error.ts
    use-create-draft-decision.ts       # useCreateDraftDecision
    use-update-decision.ts             # useUpdateDecision（预留落点）
```

### 2.2 落点约束

- 每个 feature slice 的 `application/` 目录是该切片唯一正式 mutation 承接位
- `application/` 目录不得承接读取逻辑（读取由 `data/` 层 query owner 承接）
- `normalize-error.ts` 是切片级共享工具，只在切片内部复用；当跨切片稳定复用后才允许提升到 `shared/lib/`
- `use-update-*.ts` 文件在 phase06 阶段只设计预留落点与接口签名，不要求实现；后续阶段实现 update 能力时直接填充
- `application/index.ts` 作为 barrel export，只导出该切片的正式 application owner hooks

### 2.3 既有 `data/` 层关系

- `data/api-adapter.ts` 继续承接 HTTP 调用与字段转换，不变
- `data/` 层的 `createXxx` 函数被 application owner 内部调用，不再被页面组件直接调用
- `data/` 层的 `fetchXxx` / `useXxxQuery` 等 query owner 保持纯只读

---

## 3. Create Draft Application Owner 设计

### 3.1 统一接口契约

每个 `useCreateDraft*` application owner SHALL 返回以下接口：

```ts
interface ApplicationOwnerResult<TInput, TResponse> {
  /** fire-and-forget（Onboarding 步骤消费） */
  mutate: (input: TInput) => void
  /** awaitable（create 页面消费，需要响应数据做回流）
   *  失败时抛出 NormalizedError（非原始 ApiError / Error） */
  mutateAsync: (input: TInput) => Promise<TResponse>
  /** 提交中状态 */
  isPending: boolean
  /** 是否出错 */
  isError: boolean
  /** 归一化后的错误（从 error state 派生，非 mutateAsync 抛出值） */
  error: NormalizedError | null
  /** 创建成功响应 */
  data: TResponse | undefined
}
```

> 约束：`mutateAsync` 是 application owner 内部包装版本，在 catch 中调用 `normalizeApiError(e)` 后 re-throw `NormalizedError`。消费方使用 `mutateAsync` 时 catch 到的已是 `NormalizedError`，无需自行归一化。`mutate`（fire-and-forget）不抛出，消费方通过 `error` state 获取归一化错误。

### 3.2 useCreateDraftProduct

**文件**：`frontend/src/features/product-registry/application/use-create-draft-product.ts`

**输入类型**：

```ts
interface CreateDraftProductInput {
  name: string                    // 最小必填（phase06-02）
  description?: string            // 可选，默认 ''（phase06-02 系统填充）
  status?: ProductStatus          // 可选，默认 'active'（phase06-02 系统填充）
}
```

**内部实现伪代码**：

```
useCreateDraftProduct():
  queryClient = useQueryClient()
  mutation = useMutation({
    mutationFn: (input) => createProduct({
      name: input.name,
      description: input.description ?? '',
      status: input.status ?? 'active',
    }),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ['product-list'] })
      queryClient.invalidateQueries({ queryKey: ['onboarding', 'read'] })
      return data
    },
  })

  // 包装 mutateAsync：失败时归一化后 re-throw NormalizedError
  const mutateAsync = async (input) => {
    try {
      return await mutation.mutateAsync(input)
    } catch (e) {
      throw normalizeApiError(e)
    }
  }

  return {
    mutate: mutation.mutate,
    mutateAsync,                      // 包装版本
    isPending: mutation.isPending,
    isError: mutation.isError,
    error: mutation.error ? normalizeApiError(mutation.error) : null,
    data: mutation.data,
  }
```

**职责**：

1. 接收最小 draft input，对未提供的可选字段填充系统默认值（`description: ''`, `status: 'active'`）
2. 调用 `createProduct`（来自 `data/api-adapter.ts`）
3. 成功后失效 `['product-list']` + `['onboarding', 'read']`
4. 失败时归一化错误
5. 返回 `{ mutate, mutateAsync, isPending, isError, error, data }`

**禁止**：

- 不得处理 `toast`（由消费方决定展示）
- 不得处理 `navigate` / 回流（由消费方决定回流目标）
- 不得承接读取逻辑

### 3.3 useCreateDraftRepository

**文件**：`frontend/src/features/repository-binding/application/use-create-draft-repository.ts`

**输入类型**：

```ts
interface CreateDraftRepositoryInput {
  name: string                    // 最小必填（phase06-02）
  url: string                     // 最小必填（phase06-02）
  provider?: string               // 可选，默认 'manual'（phase06-02 系统填充）
  status?: RepositoryStatus       // 可选，默认 'active'（phase06-02 系统填充）
}
```

**系统填充**：`provider: 'manual'`, `status: 'active'`

**失效目标**：`['repository-list']` + `['onboarding', 'read']`

### 3.4 useCreateDraftModule

**文件**：`frontend/src/features/module-registry/application/use-create-draft-module.ts`

**输入类型**：

```ts
interface CreateDraftModuleInput {
  name: string                    // 最小必填（phase06-02）
  description?: string            // 可选，默认 ''（phase06-02 系统填充）
  status?: ModuleStatus           // 可选，默认 'active'（phase06-02 系统填充）
  capability_key?: string | null  // 可选，默认 null（phase06-02 系统填充）
}
```

**系统填充**：`description: ''`, `status: 'active'`, `capability_key: null`

**失效目标**：`['module-list']` + `['onboarding', 'read']`

### 3.5 useCreateDraftDecision

**文件**：`frontend/src/features/decision-center/application/use-create-draft-decision.ts`

**输入类型**：

```ts
interface CreateDraftDecisionInput {
  title: string                   // 最小必填（phase06-02）
  choice: string                  // 最小必填（phase06-02）
  reason: string                  // 最小必填（phase06-02）
  context?: string                // 可选，默认 ''（phase06-02 系统填充）
  problem?: string                // 可选，默认 ''（phase06-02 系统填充）
  alternatives?: string[]         // 可选，默认 []（phase06-02 系统填充）
  impact?: string                 // 可选，默认 ''（phase06-02 系统填充）
  status?: DecisionStatus         // 可选，默认 'proposed'（phase06-02 系统填充）
  source_module_id?: string       // 可选，来源上下文自动带入（phase06-02）
}
```

**系统填充**：`context: ''`, `problem: ''`, `alternatives: []`, `impact: ''`, `status: 'proposed'`

**失效目标**：`['decision-list']` + `['onboarding', 'read']`

### 3.6 Draft Input 与既有 Input 的关系

| 既有类型（`types.ts`） | Draft 类型（`application/`） | 关系 |
| --- | --- | --- |
| `CreateProductInput`（全字段必填） | `CreateDraftProductInput`（最小必填 + 可选） | Draft 类型的可选字段填默认值后产生 `CreateProductInput` |
| `CreateRepositoryInput` | `CreateDraftRepositoryInput` | 同上 |
| `CreateModuleInput` | `CreateDraftModuleInput` | 同上 |
| `CreateDecisionInput` | `CreateDraftDecisionInput` | 同上 |

> 约束：Draft 类型定义在 `application/` 目录中，不修改 `types.ts` 中的既有 `CreateXxxInput`。application owner 负责从 Draft 类型组装出既有 `CreateXxxInput` 传给 `data/api-adapter.ts`。

---

## 4. Update 写路径预留落点

### 4.1 预留设计

每个切片的 `application/use-update-*.ts` 预留以下接口签名（phase06 不要求实现）：

```ts
// use-update-product.ts
export function useUpdateProduct(): {
  mutate: (input: UpdateProductInput) => void
  mutateAsync: (input: UpdateProductInput) => Promise<UpdateProductResponse>
  isPending: boolean
  isError: boolean
  error: NormalizedError | null
}

// use-update-repository.ts — 同结构
// use-update-module.ts — 同结构
// use-update-decision.ts — 同结构
```

### 4.2 预留约束

- phase06 阶段只设计接口签名与文件落点，不要求实现 `mutationFn`
- `UpdateXxxInput` 类型在 phase06 阶段可不定义具体字段，只在文件中预留占位注释
- 后续阶段实现 update 时，失效目标至少包含：
  - 对应 detail query（如 `['product-detail', id]`）
  - 对应 list query（如 `['product-list']`）
  - `['onboarding', 'read']`（若 update 影响 first-run 草稿状态）
- update owner 不得内联在 detail 页面或编辑表单中

---

## 5. 统一错误归一化策略

### 5.1 NormalizedError 形状

```ts
interface NormalizedError {
  message: string              // 用户可展示的错误信息
  status?: number              // HTTP 状态码（若可用）
  code?: string                // 业务错误码（若后端提供，当前阶段预留）
}
```

### 5.2 归一化函数

每个切片的 `application/normalize-error.ts` 提供切片级 `normalizeApiError` 函数：

```ts
import { ApiError } from '../data/api-adapter'

export interface NormalizedError {
  message: string
  status?: number
  code?: string
}

export function normalizeApiError(error: unknown): NormalizedError {
  if (error instanceof ApiError) {
    return {
      message: error.message,
      status: error.status,
    }
  }
  if (error instanceof Error) {
    return { message: error.message }
  }
  return { message: '未知错误，请稍后重试' }
}
```

### 5.3 归一化约束

- application owner 不得将原始 `unknown` / `Error` / `ApiError` 直接暴露给消费方
- application owner 返回的 `error` 字段必须是 `NormalizedError | null`
- 消费方（create 页面、Onboarding 步骤）只消费 `NormalizedError.message`，不依赖原始错误类型
- `normalizeApiError` 是切片级工具，当 4 个切片的实现稳定一致后，才允许提升到 `shared/lib/normalize-api-error.ts`
- 当前阶段不在 `normalizeApiError` 中做 HTTP 状态码到中文消息的映射（后端 `ApiError.message` 已提供可展示信息）

### 5.4 消费方错误展示

| 消费方 | 错误展示方式 |
| --- | --- |
| create 页面 | `toast.error('创建失败：' + error.message)` |
| Onboarding 步骤 | 将 `error?.message` 传入 `OnboardingDraftForm` 的 `submitError` prop（由 phase06-06 §11.3 定义） |

> 约束：application owner 不调用 `toast`（展示职责归消费方）。

---

## 6. 统一成功回流与 query 失效策略

### 6.1 query 失效矩阵

| Application Owner | 失效目标 | 说明 |
| --- | --- | --- |
| `useCreateDraftProduct` | `['product-list']` + `['onboarding', 'read']` | canonical 列表 + Onboarding 自动前进 |
| `useCreateDraftRepository` | `['repository-list']` + `['onboarding', 'read']` | 同上 |
| `useCreateDraftModule` | `['module-list']` + `['onboarding', 'read']` | 同上 |
| `useCreateDraftDecision` | `['decision-list']` + `['onboarding', 'read']` | 同上 |

### 6.2 回流职责拆分

| 职责 | 承接方 | 说明 |
| --- | --- | --- |
| 调用 API | application owner | 通过 `mutationFn` |
| 系统填充默认值 | application owner | draft-first 占位值 |
| query 失效 | application owner | `onSuccess` 回调 |
| 错误归一化 | application owner | `normalizeApiError` |
| `toast` 展示 | 消费方（create 页面 / Onboarding 步骤） | 成功与失败 toast |
| `navigate` 回流 | 消费方（create 页面） | 使用 `mutateAsync` 返回的 `data` 导航 |
| 来源上下文携带 | 消费方（create 页面） | `fromList / fromDashboard / fromModuleDetail / fromProductDetail` 等 |
| Onboarding 自动前进 | 消费方（Onboarding 步骤） | 通过 `OnboardingRead` 重新读取后自动定位 |

### 6.3 create 页面回流模式

create 页面使用 `mutateAsync` 消费 application owner。application owner 的 `mutateAsync` 在失败时 SHALL 抛出 `NormalizedError`（而非原始 `ApiError` / `Error`），使消费方 catch 块直接访问 `.message`，无需导入 `normalizeApiError`：

```
// ProductCreatePage 回收后的模式
const { mutateAsync, isPending } = useCreateDraftProduct()

const handleSubmit = async (input) => {
  try {
    const response = await mutateAsync(input)
    // application owner 已完成 query 失效
    toast.success('产品创建成功')
    // create 页面负责回流 + 来源上下文携带
    navigate({
      to: '/products/$productId',
      params: { productId: response.product_id },
      search: detailSearch,  // 来源上下文由 create 页面管理
    })
  } catch (e) {
    // e 已是 NormalizedError（application owner 内部归一化后 re-throw）
    toast.error('创建失败：' + (e as NormalizedError).message)
  }
}
```

> 约束：application owner 内部包装 `mutateAsync`，在 catch 中 `normalizeApiError(e)` 后 re-throw `NormalizedError`。消费方不导入 `normalizeApiError`，只消费 `NormalizedError.message`。

### 6.4 Onboarding 步骤回流模式

Onboarding 步骤使用 `mutate` 消费 application owner（fire-and-forget + query 失效驱动自动前进）：

```
// ProductStep 消费模式（phase06-06 §11）
const { mutate, isPending, error } = useCreateDraftProduct()

// 传入 OnboardingDraftForm 的 onSubmit
<OnboardingDraftForm
  submitting={isPending}
  onSubmit={(input) => mutate(input)}
  submitError={error?.message}
/>
// mutate 成功后 application owner 失效 ['onboarding', 'read']
// OnboardingRead 重新读取后自动前进到下一步（phase06-06 §4.3）
```

### 6.5 回流约束

- application owner 不得调用 `navigate`（回流是消费方职责）
- application owner 不得调用 `toast`（展示是消费方职责）
- create 页面的来源上下文（`fromList / fromDashboard / fromModuleDetail / fromProductDetail`）由 create 页面自行管理，不进入 application owner
- Onboarding 步骤不需要显式回流（通过 query 失效 + `OnboardingRead` 重新读取自动前进）
- 回收后 create 页面不得保留 `useMutation` / `useQueryClient` import（这两个由 application owner 内部使用）

---

## 7. 旧模式回收清单

### 7.1 必须回收（phase06 范围内）

| 页面 | 当前模式 | 回收目标 | 对应 application owner |
| --- | --- | --- | --- |
| `ProductCreatePage` | page-level `useMutation` + `invalidateQueries` + `toast` + `navigate` | 消费 `useCreateDraftProduct`，保留来源上下文与回流 | `product-registry/application/use-create-draft-product.ts` |
| `RepositoryCreatePage` | page-level `useMutation` + `invalidateQueries` + `toast` + `navigate` | 消费 `useCreateDraftRepository`，保留来源上下文与 product transit | `repository-binding/application/use-create-draft-repository.ts` |
| `ModuleCreatePage` | page-level `useMutation` + `invalidateQueries` + `toast` + `navigate` | 消费 `useCreateDraftModule`，保留 dashboard 上下文 | `module-registry/application/use-create-draft-module.ts` |
| `DecisionCreatePage` | page-level `useMutation` + `invalidateQueries` + `toast` + `navigate` | 消费 `useCreateDraftDecision`，保留 fromList 与完整来源 Module 链路（`sourceModuleId` + `sourceModuleName`） | `decision-center/application/use-create-draft-decision.ts` |

### 7.2 识别但允许过渡保留

| 组件 / 页面 | 当前模式 | 后续回收建议 | phase06 处理 |
| --- | --- | --- | --- |
| `ProductModuleBindingPanel` | panel-level `useMutation`（`bindModuleToProduct`） | 后续回收到 `product-registry/application/use-bind-module-to-product.ts` | 过渡保留 |
| `RepositoryProductBindingPanel` | panel-level `useMutation`（`bindProductToRepository`） | 后续回收到 `repository-binding/application/use-bind-product-to-repository.ts` | 过渡保留 |
| `RepositoryModuleMappingPanel` | panel-level `useMutation`（`mapModuleToRepository`） | 后续回收到 `repository-binding/application/use-map-module-to-repository.ts` | 过渡保留 |
| `DecisionModuleCandidatePanel` | panel-level `useMutation`（`linkDecisionToTarget`） | 后续回收到 `decision-center/application/use-link-decision-to-target.ts` | 过渡保留 |
| `ReleaseCreatePage` | page-level `useMutation`（`createRelease`） | 后续回收到 `module-registry/application/use-create-release.ts` | 过渡保留 |

### 7.3 回收后 create 页面职责

回收后，create 页面 SHALL 只承接以下职责：

- 来源上下文读取与展示（`fromList / fromDashboard / fromModuleDetail / fromProductDetail`）
- 主动取消返回（按真实来源决定返回路径）
- 调用 `useCreateDraft*` 的 `mutateAsync`
- 成功后 `toast.success` + `navigate` 到 detail（携带来源上下文）
- 失败后 `toast.error`（使用归一化错误信息）
- 表单壳层布局与表单组件渲染

create 页面 SHALL NOT：

- 内联 `useMutation`
- 直接调用 `createXxx` API 函数
- 直接调用 `queryClient.invalidateQueries`
- 自行定义 mutation 的 `onSuccess` / `onError`

### 7.4 表单组件字段级 draft-first 放宽

回收后表单组件（`ProductCreateForm` 等）SHALL 在保持 `onSubmit / submitting / submitError` props 接口不变的前提下，放宽前端必填约束以对齐 phase06-02 draft-first 语义。表单组件职责仍为字段收集 + 提交事件 + 局部 loading / error 展示，但字段级 required 与提交阻断逻辑必须按 §7.5 放宽。

> 约束：如果不放宽表单字段级必填约束，即使 mutation 已回收到 application owner，用户仍无法以 draft-first 最小字段提交（表单前端校验会阻断），phase06-02 的 draft-first 语义在实现后失效。

### 7.5 表单字段级放宽清单

| 表单组件 | 当前阻断提交的字段 | 放宽后必填字段 | 放宽动作 |
| --- | --- | --- | --- |
| `ProductCreateForm` | `name` + `description`（`if (!name \|\| !description) return`） | 仅 `name` | 移除 `description` 的 `required` HTML 属性与提交阻断；移除 `status` Select（系统填充 `active`）；按钮 disabled 条件改为 `!name` |
| `RepositoryCreateForm` | `name` + `url` + `provider`（`if (!name \|\| !url \|\| !provider) return`） | 仅 `name` + `url` | 移除 `provider` 的 `required` HTML 属性与提交阻断；移除 `status` Select（系统填充 `active`）；按钮 disabled 条件改为 `!name \|\| !url` |
| `ModuleCreateForm` | `name` + `description`（`if (!name \|\| !description) return`） | 仅 `name` | 移除 `description` 的 `required` HTML 属性与提交阻断；移除 `status` Select（系统填充 `active`）；按钮 disabled 条件改为 `!name` |
| `DecisionCreateForm` | `title` + `context` + `problem` + `choice` + `reason`（HTML `required`） | 仅 `title` + `choice` + `reason` | 移除 `context` / `problem` 的 `required` HTML 属性与"必填"标签；移除 `status` Select（系统填充 `proposed`）；`alternatives` / `impact` 保持可选不变 |

放宽约束：

- 表单提交时只校验 phase06-02 冻结的最小必填字段
- 未填写的非必填字段不传入 application owner（由 application owner 系统填充默认值）
- `status` 字段不再由用户在表单中选择，改为由 application owner 系统填充（Product / Repository / Module = `active`，Decision = `proposed`）
- 表单的 `onSubmit` 回调签名改为接收 Draft Input 类型（如 `CreateDraftProductInput`），而非既有 `CreateXxxInput`
- `isDirty` 判定条件相应放宽（只看最小必填字段是否非空）

---

## 8. 回收迁移顺序

### 8.1 迁移阶段

| 阶段 | 内容 | 依赖 | 可并行 |
| --- | --- | --- | --- |
| 阶段 1 | 4 个切片各自创建 `application/` 目录 + `normalize-error.ts` + `index.ts` | 无 | 4 切片可并行 |
| 阶段 2 | 实现 `useCreateDraftProduct`（参考实现） | 阶段 1 | — |
| 阶段 3 | 回收 `ProductCreatePage` + 放宽 `ProductCreateForm` 字段级必填消费 `useCreateDraftProduct`（验证模式） | 阶段 2 | — |
| 阶段 4 | 实现 `useCreateDraftModule` / `useCreateDraftRepository` / `useCreateDraftDecision` | 阶段 1 | 3 个可并行 |
| 阶段 5 | 回收 `ModuleCreatePage` / `RepositoryCreatePage` / `DecisionCreatePage` + 放宽对应表单字段级必填 | 阶段 4 | 3 个可并行 |
| 阶段 6 | 实现 4 个 `use-update-*.ts` 预留落点（接口签名 + 文件创建） | 阶段 1 | 4 个可并行 |

### 8.2 迁移约束

- 阶段 2-3 必须先于阶段 4-5 完成，以验证 application owner 模式可工作
- 每个回收完成后，对应 create 页面不得再保留 `useMutation` / `useQueryClient` import
- 每个回收完成后，对应表单组件的必填字段必须已按 §7.5 放宽（不得只迁移 mutation owner 而保留旧的字段级阻断）
- 回收过程中不得改变 create 页面对用户可见的行为（来源上下文、回流目标、toast 文案保持不变）
- 回收过程中不得改变 API 调用形状（`createProduct` 等函数签名不变，application owner 内部调用）
- 阶段 6 的预留落点只创建文件与接口签名，不实现 `mutationFn`

---

## 9. phase06 收口标准

### 9.1 必须满足

1. 4 个 feature slice 各自有 `application/` 目录
2. 4 个 `useCreateDraft*` application owner 已实现
3. 4 个 create 页面已回收，不再内联 `useMutation`
4. 4 个表单组件的必填字段已按 §7.5 放宽，draft-first 最小字段可提交
5. application owner 被 Onboarding 步骤与 create 页面共享
6. 错误归一化通过 `NormalizedError` 统一形状
7. query 失效矩阵已按 §6.1 落地
8. 4 个 `use-update-*.ts` 预留落点已创建（接口签名级别）
9. binding/link panel 与 release create page 的旧模式已识别
10. DecisionCreatePage 回收后完整保留 `sourceModuleId` + `sourceModuleName` 来源链路

### 9.2 禁止

- create 页面继续内联 `useMutation`
- 表单组件保留旧的前端必填阻断（如 `if (!name || !description) return`）或 `status` 用户选择器
- application owner 承接读取逻辑
- application owner 调用 `navigate` 或 `toast`
- 新增写路径复制 page-level `useMutation` 散装模式
- 跨切片直接引用其他切片的 `normalize-error.ts`（应先提升到 `shared/lib/`）
- DecisionCreatePage 回收后丢失 `sourceModuleId` 或 `sourceModuleName` 来源展示

---

## ADDED Requirements

### Requirement: Application Owner 切片落点冻结

系统 SHALL 在 4 个 feature slice 中各自新增 `application/` 目录，作为该切片唯一正式 mutation 承接位。

#### Scenario: 目录落点

- **WHEN** 接手者实现 phase06 写路径收敛
- **THEN** 必须在 `product-registry / repository-binding / module-registry / decision-center` 各自新增 `application/` 目录
- **AND** 每个目录必须包含 `index.ts`、`normalize-error.ts` 与对应的 `use-create-draft-*.ts`
- **AND** 不得在 `application/` 目录之外新增正式 mutation owner

#### Scenario: 与既有 data 层关系

- **WHEN** application owner 执行写动作
- **THEN** 必须通过 `data/api-adapter.ts` 中的 `createXxx` 函数发起 HTTP 调用
- **AND** `data/` 层的 query owner 保持纯只读
- **AND** 页面组件不得再直接调用 `data/api-adapter.ts` 中的 `createXxx` 函数

### Requirement: Create Draft Application Owner 接口契约冻结

系统 SHALL 将 4 个 `useCreateDraft*` application owner 的返回接口冻结为 `{ mutate, mutateAsync, isPending, isError, error, data }`。

#### Scenario: 返回接口

- **WHEN** 接手者实现 `useCreateDraftProduct / useCreateDraftRepository / useCreateDraftModule / useCreateDraftDecision`
- **THEN** 每个 owner 必须返回 `mutate` 与 `mutateAsync` 两个提交函数
- **AND** 必须返回 `isPending / isError / error / data` 状态
- **AND** `error` 必须为 `NormalizedError | null`，不得暴露原始 `unknown` / `Error` / `ApiError`

#### Scenario: 系统填充默认值

- **WHEN** 消费方传入最小 draft input（如 Product 只传 `name`）
- **THEN** application owner 必须对未提供的可选字段填充 phase06-02 冻结的系统默认值
- **AND** Product 默认填充 `description: ''`, `status: 'active'`
- **AND** Repository 默认填充 `provider: 'manual'`, `status: 'active'`
- **AND** Module 默认填充 `description: ''`, `status: 'active'`, `capability_key: null`
- **AND** Decision 默认填充 `context: ''`, `problem: ''`, `alternatives: []`, `impact: ''`, `status: 'proposed'`

### Requirement: query 失效矩阵冻结

系统 SHALL 将 4 个 create draft application owner 的 query 失效目标冻结为对应 canonical 列表 + `['onboarding', 'read']`。

#### Scenario: 失效目标

- **WHEN** 任一 `useCreateDraft*` 成功执行
- **THEN** 必须失效对应的 canonical 列表 query（`['product-list']` / `['repository-list']` / `['module-list']` / `['decision-list']`）
- **AND** 必须失效 `['onboarding', 'read']`
- **AND** 不得在 application owner 中失效其他切片的 query

### Requirement: 回流职责拆分冻结

系统 SHALL 将写路径的回流职责拆分为 application owner（失效 + 归一化）与消费方（toast + navigate）两层。

#### Scenario: application owner 职责边界

- **WHEN** application owner 执行 mutation
- **THEN** 必须在 `onSuccess` 中失效 query
- **AND** 必须返回归一化错误
- **AND** 不得调用 `navigate` 或 `toast`

#### Scenario: create 页面回流

- **WHEN** create 页面消费 `useCreateDraft*`
- **THEN** 必须使用 `mutateAsync` 获取响应数据
- **AND** 必须在 `try` 块中 `toast.success` + `navigate`（携带来源上下文）
- **AND** 必须在 `catch` 块中 `toast.error`（使用归一化错误）

#### Scenario: Onboarding 步骤回流

- **WHEN** Onboarding 步骤组件消费 `useCreateDraft*`
- **THEN** 必须使用 `mutate`（fire-and-forget）
- **AND** 必须将 `isPending` 与 `error?.message` 传入 `OnboardingDraftForm`
- **AND** 自动前进由 `OnboardingRead` 重新读取驱动，不调用 `navigate`

### Requirement: 统一错误归一化冻结

系统 SHALL 将 application owner 的错误归一化冻结为 `NormalizedError` 形状，由切片级 `normalizeApiError` 函数承接。

#### Scenario: 归一化输出

- **WHEN** application owner 的 mutation 失败
- **THEN** 必须通过 `normalizeApiError` 将原始错误转换为 `NormalizedError`
- **AND** `NormalizedError` 必须包含 `message` 字段
- **AND** `NormalizedError` 可选包含 `status` 与 `code` 字段
- **AND** 消费方只消费 `NormalizedError.message`

#### Scenario: 切片级工具延迟晋升

- **WHEN** 接手者抽取错误归一化工具
- **THEN** `normalizeApiError` 必须先落在各切片的 `application/normalize-error.ts` 中
- **AND** 只有跨 4 个切片稳定复用后才允许提升到 `shared/lib/normalize-api-error.ts`
- **AND** 当前阶段不得跨切片直接引用其他切片的 `normalize-error.ts`

### Requirement: 旧模式回收清单冻结

系统 SHALL 将 phase06 必须回收的 create 页面清单冻结为 4 个，并识别允许过渡保留的 panel / release 级旧模式。

#### Scenario: 必须回收

- **WHEN** phase06 实现写路径收敛
- **THEN** 必须回收 `ProductCreatePage / RepositoryCreatePage / ModuleCreatePage / DecisionCreatePage` 中的 page-level `useMutation`
- **AND** 回收后这些页面不得再 import `useMutation` 或 `useQueryClient`
- **AND** 回收后这些页面不得再直接调用 `createXxx` API 函数

#### Scenario: 过渡保留

- **WHEN** 接手者识别 panel / release 级旧模式
- **THEN** `ProductModuleBindingPanel / RepositoryProductBindingPanel / RepositoryModuleMappingPanel / DecisionModuleCandidatePanel / ReleaseCreatePage` 中的 `useMutation` 允许作为过渡现实保留
- **AND** 但 phase06 新增写路径不得复制这种散装模式
- **AND** 后续重构到这些写路径时优先向切片内 application owner 回收

### Requirement: 表单字段级 draft-first 放宽冻结

系统 SHALL 在回收 4 个 create 页面时，同步放宽对应表单组件的前端必填约束，使 phase06-02 冻结的 draft-first 最小字段可在表单层直接提交，而不被旧的前端校验阻断。

#### Scenario: Product 表单放宽

- **WHEN** 接手者回收 `ProductCreateForm`
- **THEN** 必须移除 `description` 的 `required` HTML 属性与 `if (!name || !description) return` 提交阻断
- **AND** 必须移除 `status` Select 用户选择器（系统填充 `active`）
- **AND** 按钮 disabled 条件必须改为仅 `!name`
- **AND** `onSubmit` 回调签名必须改为接收 `CreateDraftProductInput`（`description` / `status` 可选）

#### Scenario: Repository 表单放宽

- **WHEN** 接手者回收 `RepositoryCreateForm`
- **THEN** 必须移除 `provider` 的 `required` HTML 属性与 `if (!name || !url || !provider) return` 提交阻断
- **AND** 必须移除 `status` Select 用户选择器（系统填充 `active`）
- **AND** 按钮 disabled 条件必须改为仅 `!name || !url`
- **AND** `onSubmit` 回调签名必须改为接收 `CreateDraftRepositoryInput`（`provider` / `status` 可选）

#### Scenario: Module 表单放宽

- **WHEN** 接手者回收 `ModuleCreateForm`
- **THEN** 必须移除 `description` 的 `required` HTML 属性与 `if (!name || !description) return` 提交阻断
- **AND** 必须移除 `status` Select 用户选择器（系统填充 `active`）
- **AND** 按钮 disabled 条件必须改为仅 `!name`
- **AND** `onSubmit` 回调签名必须改为接收 `CreateDraftModuleInput`（`description` / `status` / `capability_key` 可选）

#### Scenario: Decision 表单放宽

- **WHEN** 接手者回收 `DecisionCreateForm`
- **THEN** 必须移除 `context` 与 `problem` 的 `required` HTML 属性与"必填"标签
- **AND** 必须移除 `status` Select 用户选择器（系统填充 `proposed`）
- **AND** 表单提交只校验 `title` + `choice` + `reason`
- **AND** `onSubmit` 回调签名必须改为接收 `CreateDraftDecisionInput`（`context` / `problem` / `alternatives` / `impact` / `status` 可选）

#### Scenario: 放宽后 draft-first 可提交

- **WHEN** 用户在放宽后的表单中只填写最小必填字段（如 Product 只填 `name`）
- **THEN** 表单提交不得被前端校验阻断
- **AND** 未填写的非必填字段不传入 application owner
- **AND** application owner 对未传入的字段填充系统默认值

### Requirement: Decision 创建页来源 Module 链路冻结

系统 SHALL 在回收 `DecisionCreatePage` 时完整保留来源 Module 链路：search params 读取 `sourceModuleId` + `sourceModuleName` → `DecisionContextSourcePanel` 展示 → 提交时 `source_module_id` 持久化，回收后不得丢失任一环节。

#### Scenario: 来源参数完整读取

- **WHEN** 接手者回收 `DecisionCreatePage`
- **THEN** 页面必须同时从 search params 读取 `sourceModuleId` 与 `sourceModuleName`
- **AND** 不得只保留 `sourceModuleId` 而丢失 `sourceModuleName`
- **AND** 两个参数同时存在时才渲染 `DecisionContextSourcePanel`

#### Scenario: 来源面板展示不变

- **WHEN** `DecisionContextSourcePanel` 渲染
- **THEN** 必须同时展示 `sourceModuleName`（Badge）与 `sourceModuleId`（mono 文本）
- **AND** 回收后该面板的 props 接口与展示行为不得改变

#### Scenario: 提交持久化不变

- **WHEN** 用户从 Module Detail 带上下文进入 Decision Create 并提交
- **THEN** 提交时 `source_module_id` 必须通过 `CreateDraftDecisionInput.source_module_id` 传入 application owner
- **AND** application owner 必须将其传递给 `createDecision` API 调用
- **AND** 回收后该持久化链路不得断裂

### Requirement: 回收后 create 页面行为不变

系统 SHALL 保证回收 create 页面过程中不改变对用户可见的行为。

#### Scenario: 行为一致性

- **WHEN** 接手者回收任一 create 页面
- **THEN** 回收后的来源上下文处理（`fromList / fromDashboard / fromModuleDetail / fromProductDetail`）必须与回收前一致
- **AND** 回收后的回流目标与搜索参数必须与回收前一致
- **AND** 回收后的 toast 文案必须与回收前一致
- **AND** 回收后的 API 调用形状必须与回收前一致

#### Scenario: Decision 来源 Module 链路一致

- **WHEN** 接手者回收 `DecisionCreatePage`
- **THEN** 回收后 `sourceModuleId` + `sourceModuleName` 的完整链路必须保留（search params 读取 → 面板展示 → 提交持久化）
- **AND** 不得只保留 `sourceModuleId` 而丢失 `sourceModuleName` 的展示
- **AND** 不得因回收导致来源 Module 面板不再渲染

### Requirement: Update 写路径预留落点冻结

系统 SHALL 在 4 个切片中各自预留 `use-update-*.ts` 落点，phase06 阶段只创建文件与接口签名。

#### Scenario: 预留落点

- **WHEN** 接手者创建 update owner 文件
- **THEN** 必须在 `application/use-update-product.ts` / `use-update-repository.ts` / `use-update-module.ts` / `use-update-decision.ts` 中创建
- **AND** 必须预留 `mutate / mutateAsync / isPending / isError / error` 接口签名
- **AND** phase06 阶段不要求实现 `mutationFn`
- **AND** 后续阶段实现时失效目标至少包含 detail query + list query + `['onboarding', 'read']`

## MODIFIED Requirements

### Requirement: 既有 Create 页面写入模式

现有 `Product / Repository / Module / Decision` create 页面在 `phase06-07` 中 SHALL 改为消费各切片的 `useCreateDraft*` application owner，不再保留 page-level `useMutation`、`useQueryClient`、`invalidateQueries` 与直接 `createXxx` 调用。

#### Scenario: create 页面消费模式

- **WHEN** 接手者改造既有 create 页面
- **THEN** 页面必须通过 `useCreateDraft*` hook 获取 `mutateAsync` 与 `isPending / error`
- **AND** 必须通过 `mutateAsync` 提交并在 `try` 块中处理回流
- **AND** 必须在 `catch` 块中直接消费 `NormalizedError.message`（application owner 的 `mutateAsync` 已在内部归一化并 re-throw `NormalizedError`）
- **AND** 页面不得再 import `useMutation` 或 `useQueryClient`

### Requirement: 既有表单组件字段级放宽

`ProductCreateForm / RepositoryCreateForm / ModuleCreateForm / DecisionCreateForm` 在 `phase06-07` 中 SHALL 保持 `onSubmit / submitting / submitError` props 接口结构不变，但必须按 §7.5 放宽前端必填约束与 `status` 用户选择器，使 draft-first 最小字段可提交。

#### Scenario: 表单组件 props 接口不变

- **WHEN** 接手者回收 create 页面
- **THEN** 表单组件的 `onSubmit / submitting / submitError` props 接口结构不变
- **AND** 表单组件内部不引入 `useMutation`
- **AND** 表单组件继续由 create 页面传入 `onSubmit` 回调

#### Scenario: 表单组件字段级必填放宽

- **WHEN** 接手者回收 create 页面
- **THEN** 表单组件的 `onSubmit` 回调签名必须改为接收 Draft Input 类型（可选字段不阻断提交）
- **AND** 必须移除非最小必填字段的 `required` HTML 属性与提交阻断逻辑
- **AND** 必须移除 `status` Select 用户选择器（由 application owner 系统填充）
- **AND** 按钮 disabled 条件必须只校验 phase06-02 冻结的最小必填字段

## REMOVED Requirements

### Requirement: create 页面内联 `useMutation` 直接定义 canonical 写入语义

**Reason**: 这种模式让每个 create 页面各自维护一套 mutation、成功回流与 query 失效策略，直接违反 phase06-02 的"唯一 application 入口"约束与 phase06-05 的"mutation 固定承接位"约束。

**Migration**: phase06 后续实现统一改为：4 类对象的 `CreateDraft*` 写动作收敛到各 feature slice 的 `application/` 目录中的 `useCreateDraft*` owner；create 页面只负责消费 owner、管理来源上下文与回流。
