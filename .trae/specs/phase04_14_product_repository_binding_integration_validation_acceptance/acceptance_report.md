# phase04-14 联调、验证与验收报告

## 1. 验收环境

### 1.1 环境重建顺序（按 phase04-09 冻结链路真实执行）

| 步骤 | 入口 | 实际结果 |
|------|------|---------|
| 1 | `database/scripts/init_db.sh` | 数据库 `psco_development` 已存在，脚本幂等跳过创建 |
| 2 | 后端启动 `RUN_SEEDS_ON_BOOT=false ./backend/bin/psco-server` | 自动应用 migration，日志确认无新迁移，`0006_product_repository_binding_mainline.sql` 已在主线中 |
| 3 | `database/scripts/run_seeds.sh` | 真实执行 `seed_readonly_prereqs.sql`，补齐只读前提数据 |
| 4 | `database/scripts/reset_module_mainline.sh` | 真实执行模块主线“清空 + 恢复”，恢复 2 个模块基线 |
| 5 | `database/scripts/reset_product_repository_mainline.sh --clean-only` | 真实清空 `products / repositories / product_modules / product_repositories / module_repositories`，得到空状态前提 |
| 6 | 前端 `npm run dev -- --host 127.0.0.1` | 真实运行 Vite dev server，`http://127.0.0.1:5173`，`/api` 代理到 `:8081` |

### 1.2 运行时前置条件核对

- 后端健康检查：`GET /healthz -> 200 {"status":"ok"}` ✅
- 后端配置：启动日志确认 `run_seeds_on_boot=false`，验收不再依赖“自动 seed 等效” ✅
- 前端真实 API：`frontend/.env` 中 `VITE_USE_REAL_API=true`，开发期通过 Vite proxy 访问真实 `/api` ✅
- `.proto` 合同源：`product_registry.proto`、`repository_binding.proto` 已在仓库主线 ✅
- 当前浏览器联调入口：使用真实前端页面，不使用 mock 数据、不使用临时 SQL 页面替代 ✅

## 2. 最小主线端到端验收

### 2.1 Product Registry 最小闭环

#### 空状态

- 真实访问 `http://127.0.0.1:5173/products?statusFilter=all`
- 页面展示空状态文案：`系统中尚无任何产品，先完成首个产品登记`
- 空状态 CTA `完成首个产品登记` 可点击进入 `Product Create` ✅

#### 创建与 reread

- 通过前端表单创建产品：
  - `name`: `验收产品-A`
  - `description`: `phase04-14 联调验收创建的产品 A`
  - `status`: `active`
- 创建后真实导航到：
  - `/products/6a0ef492-4329-4bf3-ac82-ac239b68479b?fromList=true&statusFilter=all`
- `GET /api/products/6a0ef492-4329-4bf3-ac82-ac239b68479b` 返回新产品详情 ✅

#### BindModuleToProduct

- 在 `ProductDetailPage` 打开“绑定模块”面板
- 真实选择候选模块 `integration-test-module`
- 点击“确认绑定”成功后，页面停留在当前 `ProductDetailPage`
- `GET /api/products/6a0ef492-4329-4bf3-ac82-ac239b68479b` reread 结果：
  - `bound_modules` 包含 `integration-test-module`
- 页面正文真实显示：
  - `已绑定模块`
  - `integration-test-module`

**结论**：`CreateProduct -> ProductDetail -> BindModuleToProduct -> reread` 已在真实 UI 与真实 API 上闭环 ✅

### 2.2 Repository Binding 最小闭环

#### 空状态

- 从 `ProductDetailPage` 点击“进入仓库绑定”，真实进入：
  - `/repositories?fromProductDetail=true&productId=6a0ef492-...&productName=验收产品-A&productFromList=true&productStatusFilter=all&statusFilter=all`
- 页面展示空状态文案：`系统中尚无任何仓库，先完成首个仓库登记`
- 空状态 CTA `完成首个仓库登记` 可点击进入 `Repository Create` ✅

#### 创建与 BindRepositoryToProduct

- 通过前端表单创建仓库：
  - `name`: `验收仓库-A`
  - `url`: `https://github.com/psco/acceptance-repo-a`
  - `provider`: `github`
  - `status`: `active`
- 创建后真实导航到：
  - `/repositories/90142231-36d2-4447-ab35-87ae5715a89d?fromProductDetail=true&productId=6a0ef492-...`
- `RepositoryBindingDetailPage` 自动带出来源产品 `验收产品-A`
- 点击“确认绑定”后：
  - 页面停留在当前 `RepositoryBindingDetailPage`
  - `GET /api/repositories/90142231-36d2-4447-ab35-87ae5715a89d` reread 结果中 `bound_products` 包含 `验收产品-A`
  - 页面正文显示 `已绑定产品 -> 验收产品-A`

#### MapModuleToRepository

- 在同一 `RepositoryBindingDetailPage` 打开“映射模块”面板
- 真实选择候选模块 `integration-test-module`
- 点击“确认映射”后：
  - 页面停留在当前 `RepositoryBindingDetailPage`
  - `GET /api/repositories/90142231-36d2-4447-ab35-87ae5715a89d` reread 结果中 `mapped_modules` 包含 `integration-test-module`
  - 页面正文显示 `已映射模块 -> integration-test-module`

**结论**：`CreateRepository -> RepositoryDetail -> BindRepositoryToProduct -> MapModuleToRepository -> reread` 已在真实 UI 与真实 API 上闭环 ✅

## 3. Module Detail 兼容入口验收

### 3.1 Module Detail -> Product Registry

- 真实进入 `Module Detail`：`/modules/a8f72d7d-3d17-4dae-a744-d616af379a8f`（`auth-service`）
- 页面存在 `进入产品绑定` 兼容入口，不存在旧的候选读取 + 直接写入工作台 ✅
- 点击后真实跳转到：
  - `/products?fromModuleDetail=true&moduleId=a8f72d7d-...&moduleName=auth-service&statusFilter=all`
- 选择 `验收产品-A` 进入详情后，页面自动预填 `auth-service`
- 点击“确认绑定”后，页面停留在 `ProductDetailPage`
- 点击“返回模块详情”后真实回到：
  - `/modules/a8f72d7d-3d17-4dae-a744-d616af379a8f`
- 回流后的 `Module Detail` 页面出现已绑定产品链接 `验收产品-A` ✅

### 3.2 Module Detail -> Repository Binding

- 在同一 `Module Detail` 点击 `进入仓库映射`
- 真实跳转到：
  - `/repositories?fromModuleDetail=true&moduleId=a8f72d7d-...&moduleName=auth-service&statusFilter=all`
- 选择 `验收仓库-A` 进入详情后，页面自动预填 `auth-service`
- 点击“确认映射”后，页面停留在 `RepositoryBindingDetailPage`
- 点击“返回模块详情”后真实回到：
  - `/modules/a8f72d7d-3d17-4dae-a744-d616af379a8f`
- 回流后的 `Module Detail` 页面出现已映射仓库链接 `验收仓库-A` ✅

**结论**：`Module Detail` 只保留兼容入口/轻量跳转，正式写入全部发生在 canonical owner 页面 ✅

## 4. 多入口返回路径验证

### 4.1 fromProductDetail

- `Product Detail -> Repository Detail`
- `Repository Detail` 点击“返回产品详情”
- 真实回到：
  - `/products/6a0ef492-4329-4bf3-ac82-ac239b68479b?fromList=true&statusFilter=all`
- 说明 `Product Detail` 自身来源上下文已被透传并恢复 ✅

### 4.2 fromList + queryText

- 在 `Product List` 输入 `queryText=验收`
- 通过列表项进入详情页，真实 URL 为：
  - `/products/6a0ef492-4329-4bf3-ac82-ac239b68479b?fromList=true&queryText=验收&statusFilter=all`
- 在详情页点击“返回列表”
- 真实回到：
  - `/products?queryText=验收&statusFilter=all`
- 页面输入框仍保留 `验收` ✅

### 4.3 fromModuleDetail

- 从 `Module Detail` 进入 `Product Detail` / `Repository Detail`
- 点击“返回模块详情”
- 两条链都真实回到对应 `Module Detail` 页面 ✅

### 4.4 direct-entry

- 直接访问：
  - `/products/6a0ef492-4329-4bf3-ac82-ac239b68479b`
- 点击“返回列表”
- 真实回到：
  - `/products?statusFilter=all`
- 未错误恢复历史 `queryText` 或来源上下文 ✅

**结论**：`fromList / fromProductDetail / fromModuleDetail / direct-entry` 四类返回路径均完成真实页面验证 ✅

## 5. HTTP 异常路径与兼容读取矩阵

### 5.1 关键异常路径

| 路径 | 实际状态码 | 结果 |
|------|-----------|------|
| `CreateProduct` 缺少 `name` | `400` | ✅ |
| `CreateProduct` 非法 `status` | `400` | ✅ |
| `CreateRepository` 缺少 `name` | `400` | ✅ |
| `CreateRepository` 非法 `status` | `400` | ✅ |
| `ProductDetailRead` 不存在 | `404` | ✅ |
| `RepositoryDetailRead` 不存在 | `404` | ✅ |
| `BindModuleToProduct` 缺少 `product_id` / `module_id` | `404` | ✅ |
| `BindRepositoryToProduct` 缺少 `repository_id` / `product_id` | `404` | ✅ |
| `MapModuleToRepository` 缺少 `repository_id` / `module_id` | `404` | ✅ |
| `BindModuleToProduct` 重复绑定 | `409` | ✅ |
| `BindRepositoryToProduct` 重复绑定 | `409` | ✅ |
| `MapModuleToRepository` 重复映射 | `409` | ✅ |

### 5.2 候选为空与兼容读取

- `GET /api/products/6a0ef492-.../candidates/modules -> 200 []` ✅
- `GET /api/repositories/90142231-.../candidates/products -> 200 []` ✅
- `GET /api/candidates/products -> 200`，返回兼容产品列表 ✅
- `GET /api/candidates/repositories -> 200`，返回兼容仓库列表 ✅

### 5.3 500 级未收口错误检查

- 畸形 JSON -> `400` ✅
- 非法 UUID 路径参数 -> `404` ✅
- 未发现 500 级错误替代业务错误 ✅

## 6. 合同与正式规格一致性核对

- `.proto` 合同、HTTP DTO、前端 `types.ts` 仍保持单值一致 ✅
- `phase04-10` 正式规格正文与当前实现边界一致 ✅
- 旧 `/api/candidates/*` 仍为兼容委派入口，没有让 `moduleregistry` 重新成为候选读取 owner ✅

## 7. 联调中发现的问题与修复

| # | 问题 | 级别 | 收口结果 |
|---|------|------|---------|
| 1 | `phase04-13` 历史问题：`Product Detail` 来源上下文未完整透传到 `Repository Detail` | P1 | 之前已修复，本轮真实回流验证通过 |
| 2 | `phase04-13` 历史问题：列表进入详情时未带 `queryText / statusFilter` | P1 | 之前已修复，本轮真实 `fromList + queryText` 回流验证通过 |
| 3 | 本轮联调新发现：带 `prefillProductId / prefillModuleId` 的面板在成功提交后会因 `useEffect` 自动再次展开，破坏仓库详情页的互斥面板切换 | P1 | 已修复：只在首次带来源上下文进入时自动展开，成功提交或用户关闭后不再被同一来源参数反复重开；修复文件：`product-module-binding-panel.tsx`、`repository-product-binding-panel.tsx`、`repository-module-mapping-panel.tsx`；修复后真实 UI 重新验证通过 |

## 8. 阶段收口结论

- `Product Registry + Repository Binding` 最小主线已在真实前端、真实后端、真实数据库上完整走通 ✅
- 三类绑定动作的 canonical owner 页面、成功写入后的 reread 页面与多入口返回路径都可以重复复核 ✅
- `Module Detail` 旧绑定入口已收敛为兼容入口/轻量跳转，不再形成第二主工作台 ✅
- 兼容读取、异常路径、空状态、候选为空与 direct-entry 返回规则都已纳入本轮验收 ✅
- 本轮发现的阻断问题已在当前阶段修复并复测通过，不遗留隐性阻断 ✅

## 9. DoD 达成情况

| DoD 项 | 达成情况 |
|--------|---------|
| `Product Registry + Repository Binding` 最小主线可完整走通 | ✅ |
| 三类绑定动作的 canonical owner 页面、成功写入后的 reread 页面与返回路径可被重复复核 | ✅ |
| 验收结果可重复复核，并可明确证明当前阶段已形成可运行交付物 | ✅ |
| 发现的问题已收口到当前阶段，不遗留隐性阻断 | ✅ |
