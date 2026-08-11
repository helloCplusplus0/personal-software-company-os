# Phase07-03 冻结产出：迁移过程兼容策略与 Phase 收口退场标准

> 本文档是 phase07-03 spec 的执行产出，冻结 compat 策略、legacy inventory 退场门槛、最晚时点与 phase 收口判定标准。
> 产出日期：2026-08-11
> 上游：`phase07-01 frozen_scope.md`（inventory 与替代关系）、`phase07-02 frozen_scope.md`（正式传输主线组合）

---

## 1. 冻结：Compat 资产只允许短时存在的正式策略

### 1.1 合法身份定义

在 `phase07` 迁移期间，compat 资产的**唯一合法身份**是：

| 允许的类型 | 示例 | 身份 |
|-----------|------|------|
| 临时 JSON adapter | `backend/internal/moduleregistry/handler/query_handler.go:ListProductCandidates` | 迁移过渡承接位 |
| 兼容委派 handler | `backend/internal/moduleregistry/handler/command_handler.go:BindModuleToProduct` | 旧 caller 的短时兼容层 |
| 前端 adapter 切换入口 | `frontend/src/features/module-registry/data/module-registry-adapter.ts` | 旧 adapter 切换到新 client 前的过渡 switch |

### 1.2 允许并存的前提（4 项必须同时满足）

| # | 前提 | 检查方式 |
|---|------|---------|
| 1 | canonical `.proto + Connect` 替代路径已明确 | 对照 phase07-01 frozen_scope.md §2 迁移总表 |
| 2 | 当前 caller 和承接位可定位 | 对照本文 §2 每条入口的"当前调用方"列 |
| 3 | 最晚并存时点已写入 phase07 子任务链 | 对照本文 §4 退场时点映射 |
| 4 | 删除证据与回归证据已预先定义 | 对照本文 §2 每条入口的"删除证据"和"回归证据"列 |

### 1.3 禁止事项（4 条）

| # | 禁止事项 | 违反后果 |
|---|---------|---------|
| 1 | 不得保留"新 Connect + 旧 JSON"双主线长期并列 | phase07 验收直接阻塞 |
| 2 | 不得新增第二批未列入 inventory 的 compat 业务入口 | 先补入 inventory 或直接删除 |
| 3 | 不得把前端临时 adapter 留在 route、page 或展示组件中长期使用 | 必须回收至切片内固定承接位 |
| 4 | 不得把"当前无 active caller，但导出仍在"解释为可无限期保留 | 仍属于待退场资产，必须在最晚时点前删除 |

### 1.4 特殊规则：Dormant Asset

> 某 compat adapter 当前没有 active UI caller，但其导出仍存在于代码库中 → **仍属于待退场 legacy 资产**，必须在对应最晚时点前删除。

当前适用此规则的前端 adapter：
- `module-registry/data/api-adapter.ts` 中的 `bindModuleToProduct` / `mapModuleToRepository` / `fetchProductCandidates` / `fetchRepositoryCandidates` 导出
- `module-registry/data/module-registry-adapter.ts` 中对应的 runtime switch 分支
- `module-registry/data/mock-adapter.ts` 中对应的历史演示导出；若保留，必须显式标注为非正式业务主线，不计作 `phase07` 收口时的正式 caller

---

## 2. 冻结：Legacy / Compat 业务入口 Inventory（退场版本）

> 每条入口均已从 phase07-01 §4 的静态清单升级为带退场门槛的正式兼容策略。

### 2.1 `GET /api/candidates/products`

| 属性 | 值 |
|------|-----|
| **当前调用方** | 当前前端无 active caller；仅余 `module-registry/data/api-adapter.ts`、`module-registry/data/module-registry-adapter.ts` 与 `module-registry/data/mock-adapter.ts` 中的残留导出 / switch 分支 |
| **存在原因** | phase02 Module Registry 曾临时承接 Product 候选读取；phase04 起 canonical owner 已迁移到 Product Registry |
| **替代 RPC / Connect path** | `ProductRegistryService.ListProducts`（status_filter=ACTIVE），Connect path: `/psco.product_registry.v1.ProductRegistryService/ListProducts` |
| **最晚退场时点** | `phase07-09`（后端 compat 路由退场） |
| **删除证据** | 1) `router.go` 中移除 `r.Get("/api/candidates/products", ...)` 路由注册；2) `moduleregistry/handler/query_handler.go` 中删除 `ListProductCandidates` 方法；3) `moduleregistry/handler/query_handler.go` 中删除 `ProductCandidateReader` 字段与构造函数参数 |
| **回归证据** | 1) 若前端存在候选读取正式承接位，则通过 `ProductRegistryService.ListProducts` Connect path 可正常获取候选列表；2) `GET /api/candidates/products` 返回 404；3) 仓库中不再存在真实页面 / 组件对 `module-registry` compat 候选读取导出的正式调用 |

### 2.2 `GET /api/candidates/repositories`

| 属性 | 值 |
|------|-----|
| **当前调用方** | 当前前端无 active caller；仅余 `module-registry/data/api-adapter.ts`、`module-registry/data/module-registry-adapter.ts` 与 `module-registry/data/mock-adapter.ts` 中的残留导出 / switch 分支 |
| **存在原因** | phase02 Module Registry 曾临时承接 Repository 候选读取；phase04 起 canonical owner 已迁移到 Repository Binding |
| **替代 RPC / Connect path** | `RepositoryBindingService.ListRepositories`（status_filter=ACTIVE），Connect path: `/psco.repository_binding.v1.RepositoryBindingService/ListRepositories` |
| **最晚退场时点** | `phase07-09`（后端 compat 路由退场） |
| **删除证据** | 1) `router.go` 中移除 `r.Get("/api/candidates/repositories", ...)` 路由注册；2) `moduleregistry/handler/query_handler.go` 中删除 `ListRepositoryCandidates` 方法；3) `moduleregistry/handler/query_handler.go` 中删除 `RepositoryCandidateReader` 字段与构造函数参数 |
| **回归证据** | 1) 若前端存在候选读取正式承接位，则通过 `RepositoryBindingService.ListRepositories` Connect path 可正常获取候选列表；2) `GET /api/candidates/repositories` 返回 404；3) 仓库中不再存在真实页面 / 组件对 `module-registry` compat 候选读取导出的正式调用 |

### 2.3 `POST /api/modules/{moduleId}/bindings/products`

| 属性 | 值 |
|------|-----|
| **当前调用方** | 当前前端无 active caller；正式写入 owner 位于 `product-registry/components/product-module-binding-panel.tsx`，`module-registry` 侧仅残留 compat 导出 / switch 分支 |
| **存在原因** | phase04 起为兼容委派，实际写动作委派到 Product Registry 的 `BindModuleToProduct` |
| **替代 RPC / Connect path** | `ProductRegistryService.BindModuleToProduct`，Connect path: `/psco.product_registry.v1.ProductRegistryService/BindModuleToProduct` |
| **最晚退场时点** | `phase07-10`（前端 adapter 回收与 mutation 收口） |
| **删除证据** | 1) `router.go` 中移除 `r.Post("/api/modules/{moduleId}/bindings/products", ...)` 路由注册；2) `moduleregistry/handler/command_handler.go` 中删除 `BindModuleToProduct` 兼容委派方法；3) 前端 `module-registry/data/api-adapter.ts` 中删除 `bindModuleToProduct` 导出，且 `module-registry-adapter.ts` 不再保留对应 real-api switch 分支 |
| **回归证据** | 1) `ProductModuleBindingPanel` 通过 `ProductRegistryService.BindModuleToProduct` Connect client 执行成功；2) `POST /api/modules/{moduleId}/bindings/products` 返回 404；3) 仓库中不再存在真实页面 / 组件对 `module-registry` compat 绑定导出的正式调用 |

### 2.4 `POST /api/modules/{moduleId}/bindings/repositories`

| 属性 | 值 |
|------|-----|
| **当前调用方** | 当前前端无 active caller；正式写入 owner 位于 `repository-binding/components/repository-module-mapping-panel.tsx`，`module-registry` 侧仅残留 compat 导出 / switch 分支 |
| **存在原因** | phase04 起为兼容委派，实际写动作委派到 Repository Binding 的 `MapModuleToRepository` |
| **替代 RPC / Connect path** | `RepositoryBindingService.MapModuleToRepository`，Connect path: `/psco.repository_binding.v1.RepositoryBindingService/MapModuleToRepository` |
| **最晚退场时点** | `phase07-10`（前端 adapter 回收与 mutation 收口） |
| **删除证据** | 1) `router.go` 中移除 `r.Post("/api/modules/{moduleId}/bindings/repositories", ...)` 路由注册；2) `moduleregistry/handler/command_handler.go` 中删除 `MapModuleToRepository` 兼容委派方法；3) 前端 `module-registry/data/api-adapter.ts` 中删除 `mapModuleToRepository` 导出，且 `module-registry-adapter.ts` 不再保留对应 real-api switch 分支 |
| **回归证据** | 1) `RepositoryModuleMappingPanel` 通过 `RepositoryBindingService.MapModuleToRepository` Connect client 执行成功；2) `POST /api/modules/{moduleId}/bindings/repositories` 返回 404；3) 仓库中不再存在真实页面 / 组件对 `module-registry` compat 映射导出的正式调用 |

---

## 3. 冻结：后端与前端退场证据模型

### 3.1 后端退场证据（逐条 compat 入口必须满足）

| 证据项 | 具体内容 | 核实方式 |
|--------|---------|---------|
| 路由删除 | `router.go` 中不再注册该 compat 路径 | `grep` 路径在 `router.go` 中无匹配 |
| handler 删除 | 对应 handler 方法及关联类型（Reader 字段/构造函数参数）已删除 | `grep` 方法名在 handler 文件中无匹配 |
| 编译通过 | `go build ./...` 无错误 | CI / 本地编译 |
| 旧路径不可访 | `curl` 旧 JSON 路径 → 404 | 验收脚本 |

### 3.2 前端退场证据（逐条 compat 入口必须满足）

| 证据项 | 具体内容 | 核实方式 |
|--------|---------|---------|
| 正式 caller 切换 | 所有真实页面 / 组件调用已切到 Connect client 或 canonical slice owner | 代码审查：无页面 / 组件继续通过 `module-registry` compat 导出承接正式动作 |
| real-api 旧导出删除 | `module-registry/data/api-adapter.ts` 中不再导出旧 compat 函数 | 代码审查 + `grep` `api-adapter.ts` |
| adapter switch 收口 | `module-registry-adapter.ts` 中不再保留对应 real-api switch 分支；若 mock 分支保留，必须显式标注为历史演示路径且不计作正式 caller | 代码审查 |
| 编译通过 | `tsc -b --noEmit` 无错误 | CI / 本地编译 |
| 页面功能正常 | 正式页面动作通过 Connect 执行成功，且 ModuleDetailPage 继续保持只读摘要 + 跳转入口 | 联调验收 |

### 3.3 回归证据统一模型

| 证据项 | 覆盖范围 | 验收方式 |
|--------|---------|---------|
| 替代 Connect path 可用 | 4 条 compat 入口的替代 RPC 均可正常调用 | 联调验收脚本逐条调用 Connect RPC |
| 旧 JSON 路径不可再访问 | 4 条 compat 入口的旧路径均返回 404 | 联调验收脚本逐条 curl 旧路径 |
| 前端功能不退化 | ProductDetailPage / RepositoryBindingDetailPage 的正式动作正常，且 ModuleDetailPage 继续只读摘要 + 兼容跳转 | 手工验收 + 截图 |

---

## 4. 冻结：最晚退场时点与 Phase07 子任务链映射

```
phase07 子任务链                    compat 入口退场窗口
──────────────────────────────────────────────────────────
phase07-04  Go Connect handler 设计
phase07-05  前端 query/application 迁移设计
phase07-06  transport 实现与双栈过渡
phase07-07  前端 Connect client 实现
phase07-08  后端 Connect handler 实现
phase07-09  后端 compat 路由退场  ←── GET /api/candidates/products     退场
                                 ←── GET /api/candidates/repositories  退场
phase07-10  前端 adapter 回收    ←── POST /api/modules/{id}/bindings/products  退场
                                 ←── POST /api/modules/{id}/bindings/repositories 退场
phase07-11  验收与收口核销       ←── 全部 4 条核销确认
```

### 4.1 候选读取 compat 入口（#1, #2）

| 入口 | 退场子任务 | 最晚时点 | 提前条件 |
|------|-----------|---------|---------|
| `GET /api/candidates/products` | phase07-09 | phase07-09 完成 | `ProductRegistryService.ListProducts` Connect 读路径已可用 |
| `GET /api/candidates/repositories` | phase07-09 | phase07-09 完成 | `RepositoryBindingService.ListRepositories` Connect 读路径已可用 |

### 4.2 Module-centered 绑定 compat 入口（#3, #4）

| 入口 | 退场子任务 | 最晚时点 | 提前条件 |
|------|-----------|---------|---------|
| `POST /api/modules/{moduleId}/bindings/products` | phase07-10 | phase07-10 完成 | ProductDetailPage 的正式 bind action 已切到 Connect client，且 `module-registry` compat caller 已清零 |
| `POST /api/modules/{moduleId}/bindings/repositories` | phase07-10 | phase07-10 完成 | RepositoryBindingDetailPage 的正式 map action 已切到 Connect client，且 `module-registry` compat caller 已清零 |

### 4.3 冻结约束

- **候选读取 compat 入口**：不晚于 phase07-09 退场，不得推迟到 phase07-10 或 phase07-11
- **Module-centered 绑定 compat 入口**：不晚于 phase07-10 退场，不得推迟到 phase07-11
- **phase07-11 收口**：只做核销确认，不做退场动作
- 不得出现"到收口时再看"的模糊状态

---

## 5. 冻结：Phase07 收口退场标准（双门槛）

### 5.1 门槛一：Connect 主线已存在

| 条件 | 验证方式 |
|------|---------|
| 34 条 proto-defined business RPC 均有 Connect handler | phase07-01 迁移总表逐条核销 |
| `buf.gen.yaml` 已包含 3 插件 | `buf generate` 成功 |
| 前端 Connect transport 已建立 | `createConnectTransport` 承接位存在 |
| 所有页面动作通过 Connect client 执行 | 联调验收 |

### 5.2 门槛二：旧 JSON 主线已退场

| 条件 | 验证方式 |
|------|---------|
| 4 条 legacy compat 入口全部核销 | 本文 §2 逐条检查删除证据 + 回归证据 |
| `router.go` 中不再保留 compat 业务路由 | `grep` 验证 |
| `moduleregistry/handler/` 中不再保留 compat handler | `grep` 验证 |
| 前端真实业务主线不再经过 `module-registry` compat adapter | 代码审查：无真实页面 / 组件 caller；real-api 旧导出已删除；mock 若保留则已显式隔离为历史演示路径 |

### 5.3 收口判定规则

```
phase07 收口 = 门槛一（全部满足）AND 门槛二（全部满足）

任何一项不满足 → phase07 不得收口
```

**禁止的判定方式**：
- ❌ "Connect 主线已存在，旧 JSON 可后续再删" → 不通过
- ❌ "4 条 compat 入口已删 3 条，剩 1 条当前无 caller" → 不通过
- ❌ "真实业务 caller 已切走，但仍把 `module-registry` compat real-api 导出留在正式主线中" → 不通过

---

## 6. 与上游文档一致性声明

| 上游文档 | 关键结论 | 本产出对齐 |
|---------|---------|-----------|
| `phase07-01 frozen_scope.md` §4 | 4 条 legacy inventory 与替代关系 | §2 逐条升级为带退场门槛的正式策略 |
| `phase07-01 frozen_scope.md` §5 | phase07 收口判定标准 | §5 双门槛细化，补充逐项核销模型 |
| `phase07-02 frozen_scope.md` §1 | chi 不再承担业务合同定义 | §3.1 后端退场证据要求路由删除 |
| `phase07-02 frozen_scope.md` §2 | 单一 /api 前缀 + Connect | §2 替代 Connect path 全部通过 /api |
| `phase07-02 frozen_scope.md` §4 | 前端 query/application 边界 | §3.2 前端退场证据要求 adapter 切换到固定承接位 |
| `shared_baseline.md` §2.4 | 旧 JSON 只作 compat 资产 | 全文 compat 策略 |
| `architecture_plan.md` §4.6 | 一次性正式切换 | §5.3 双门槛同时满足才能收口 |
| `project_rules.md` §2.6 | `.proto` 唯一合同源 | 所有替代路径均为 Connect procedure path |
