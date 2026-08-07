# Phase04-14 Product Registry 与 Repository Binding 联调、验证与验收 Spec

## Why

`phase04-11`、`phase04-12` 与 `phase04-13` 已分别完成最小 `.proto` 合同主线、后端与数据主线、前端主线，但当前仍缺少一份独立的联调验收规格，把“各侧已经实现”收口为“最小主线已经在同一环境中被验证可运行”。如果缺少这一层，`Product Registry + Repository Binding` 仍会停留在局部自证完成，无法形成可重复复核、可明确证明当前阶段已形成运行交付物的证据。

同时，`phase04` 的联调验收还必须承担阶段收口职责：`CreateProduct / CreateRepository / BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 的最小验收路径、空状态、错误态、候选为空、目标不存在、重复绑定、旧入口兼容跳转、三类绑定成功后的 reread 与多入口返回路径都要进入同一轮验证；联调中发现的实现偏差、兼容断链或正式规格漂移，必须在当前阶段明确收口，不能把隐性阻断带到后续 review 或 root sync。

## What Changes

- 为 `Product Registry + Repository Binding` 新增一份联调、验证与验收 spec
- 冻结联调环境重建顺序、真实前后端联调入口与验收前置条件
- 冻结 `CreateProduct / CreateRepository / BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 的最小端到端验收路径
- 冻结 `Product Detail`、`Repository Binding Detail / Workspace` 与 `Module Detail` 兼容入口的 canonical owner 页面、reread 落点与返回路径验收规则
- 冻结空状态、错误态、候选为空、目标不存在、重复绑定与返回路径验证要求
- 冻结旧 `/api/candidates/*` 兼容读取、`Module Detail` 旧绑定入口兼容跳转与多入口回流验收要求
- 冻结 `.proto`、HTTP 过渡层、前端适配层与正式规格正文的一致性核对要求
- 明确联调中发现的问题必须在当前阶段收口，不得遗留隐性阻断或规格漂移

## Impact

- Affected specs:
  - `phase04_09_integration_acceptance_reset_baseline_compat_migration`
  - `phase04_10_product_repository_binding_formal_spec`
  - `phase04_11_product_repository_binding_proto_mainline`
  - `phase04_12_product_repository_binding_backend_data_mainline`
  - `phase04_13_product_repository_binding_frontend_mainline`
- Affected code:
  - `frontend/src/features/product-registry/**`
  - `frontend/src/features/repository-binding/**`
  - `frontend/src/features/module-registry/**`
  - `frontend/src/routes/products/**`
  - `frontend/src/routes/repositories/**`
  - `backend/internal/productregistry/**`
  - `backend/internal/repositorybinding/**`
  - `backend/internal/moduleregistry/**`
  - `backend/internal/platform/**`
  - `database/migrations/0006_product_repository_binding_mainline.sql`
  - `database/scripts/reset_product_repository_mainline.sh`
  - `database/seeds/seed_product_repository_mainline_baseline.sql`
  - `proto/psco/product_registry/v1/product_registry.proto`
  - `proto/psco/repository_binding/v1/repository_binding.proto`

## ADDED Requirements

### Requirement: 联调环境必须可重复建立并以真实主线运行

系统 SHALL 为 `phase04-14` 提供可重复建立的联调环境前提，使验收在真实前端、真实后端、真实数据库与真实基线主线上执行，而不是依赖 mock、手工 SQL 或一次性会话状态。

#### Scenario: 环境初始化与启动顺序

- **WHEN** 执行 `phase04-14` 联调与验收
- **THEN** 必须复用 `phase04-09` 已冻结的环境顺序：`init_db.sh` -> 后端启动（自动 migration，含 `0006_product_repository_binding_mainline.sql`）-> `run_seeds.sh` -> `reset_module_mainline.sh` -> `reset_product_repository_mainline.sh` -> 前端启动
- **AND** 前端必须直接连接 `phase04-12 / 13` 已落地的真实后端 API，不得切回 mock 数据主线
- **AND** 不得在验收过程中新增临时 SQL、临时 seed 或临时脚本弥补环境缺口

#### Scenario: 联调前置条件核对

- **WHEN** 开始执行验收
- **THEN** 必须先明确当前 `.proto` 合同、后端路由、数据库 migration、重置脚本与前端页面主线都已处于可运行状态
- **AND** 必须先确认 `Product Registry` 与 `Repository Binding` 导航可达、后端 `/api/products` 与 `/api/repositories` 已挂载、`reset_product_repository_mainline.sh` 可执行
- **AND** 若任何前置条件不满足，不得跳过并继续手工联调

### Requirement: 最小主线必须端到端完整走通

系统 SHALL 证明 `Product Registry + Repository Binding` 当前阶段最小主线已经在同一环境中形成真实可运行交付物，而不是停留在单侧 build、单个接口自测或局部页面演示。

#### Scenario: Product Registry 最小闭环

- **WHEN** 执行 `Product Registry` 最小主线验收
- **THEN** 必须至少走通：
- **AND** `reset_product_repository_mainline.sh --clean-only`
- **AND** 进入 `Product Registry / List` 验证空状态入口
- **AND** 从空状态进入 `Product Create`
- **AND** 提交 `CreateProduct`
- **AND** 默认回流到 `ProductDetailPage`
- **AND** 在详情页读取候选列表并执行 `BindModuleToProduct`
- **AND** 停留当前 `ProductDetailPage` 并重新读取出最新已绑定模块结果
- **AND** 全路径不得依赖手工 SQL

#### Scenario: Repository Binding 最小闭环

- **WHEN** 执行 `Repository Binding` 最小主线验收
- **THEN** 必须至少走通：
- **AND** 进入 `Repository Binding / List` 验证空状态或基线列表入口
- **AND** 从列表或空状态进入 `Repository Create`
- **AND** 提交 `CreateRepository`
- **AND** 默认回流到 `RepositoryBindingDetailPage`
- **AND** 在同一详情页先执行 `BindRepositoryToProduct`
- **AND** 再执行 `MapModuleToRepository`
- **AND** 两次写入后都必须停留当前 `RepositoryBindingDetailPage` 并重新读取出最新已绑定产品与已映射模块结果
- **AND** 不得用第二个并列工作台替代 `Repository Binding Detail / Workspace`

#### Scenario: Module Detail 兼容入口闭环

- **WHEN** 用户从 `Module Detail` 发起“绑定到 Product”或“映射到 Repository”
- **THEN** 必须通过 `Module Detail` 兼容入口进入正式主线页面
- **AND** 正式写入必须发生在 canonical owner 页面，而不是继续在 `Module Detail` 内直接提交
- **AND** 写入成功后的 reread 必须落到对应 canonical owner 页面：
  - `BindModuleToProduct` -> `ProductDetailPage`
  - `BindRepositoryToProduct` -> `RepositoryBindingDetailPage`
  - `MapModuleToRepository` -> `RepositoryBindingDetailPage`

### Requirement: 空状态、错误态与返回路径必须被统一验证

系统 SHALL 将关键 UI 状态纳入本轮验收，而不是只验证 happy path。

#### Scenario: 空状态与空候选状态

- **WHEN** `products` 或 `repositories` 为空，或候选读取为空
- **THEN** `Product Registry / List` 与 `Repository Binding / List` 必须展示围绕“先登记首条记录”的空状态入口
- **AND** 空状态主动作必须分别直接进入 `Product Create` 与 `Repository Create`
- **AND** `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 的候选读取为空时必须表现为可解释的空列表状态，不得误报资源不存在或接口错误

#### Scenario: 错误态与表单/面板保留

- **WHEN** 创建、详情读取或绑定写入失败
- **THEN** 错误必须停留在当前页面或当前面板上下文
- **AND** 不得跳转独立错误页
- **AND** `CreateProduct / CreateRepository` 失败时草稿与来源上下文必须继续保留
- **AND** 三类绑定失败时当前详情、当前活动面板与当前已选候选目标必须继续保留

#### Scenario: 多入口返回路径验证

- **WHEN** 用户从 `ProductCreatePage`、`RepositoryCreatePage`、`ProductDetailPage` 或 `RepositoryBindingDetailPage` 主动返回
- **THEN** 必须符合 `phase04-06 / 10 / 13` 已冻结的返回路径规则
- **AND** 从 `fromList` 进入时必须恢复原有 `queryText / statusFilter`
- **AND** 从 `Module Detail` 或 `Product Detail` 上下文入口进入时必须回到真实来源
- **AND** 外部直达或刷新后无来源列表上下文时必须落到默认列表参数（`statusFilter: 'all'`），不得恢复历史筛选

### Requirement: 关键异常路径必须覆盖当前阶段边界

系统 SHALL 验证 `phase04` 当前阶段最容易掩盖阻断的异常路径，并以当前冻结的业务错误语义作为唯一验收标准。

#### Scenario: 最小异常路径覆盖

- **WHEN** 执行 `phase04-14` 异常路径验收
- **THEN** 必须至少覆盖以下 `11` 类异常/边界路径：
- **AND** `CreateProduct` 必填字段缺失或非法 `status` -> 400 校验失败
- **AND** `CreateRepository` 必填字段缺失或非法 `status` -> 400 校验失败
- **AND** `ProductDetailRead` 不存在的 `product_id` -> 404 资源不存在
- **AND** `RepositoryDetailRead` 不存在的 `repository_id` -> 404 资源不存在
- **AND** `BindModuleToProduct` 的 `product_id` 或 `module_id` 不存在 -> 404 资源不存在
- **AND** `BindRepositoryToProduct` 的 `repository_id` 或 `product_id` 不存在 -> 404 资源不存在
- **AND** `MapModuleToRepository` 的 `repository_id` 或 `module_id` 不存在 -> 404 资源不存在
- **AND** `BindModuleToProduct` 重复绑定 -> 409 重复冲突
- **AND** `BindRepositoryToProduct` 重复绑定 -> 409 重复冲突
- **AND** `MapModuleToRepository` 重复映射 -> 409 重复冲突
- **AND** 三条候选读取为空 -> 空列表语义（非错误）
- **AND** 不得出现 500 级未收口错误替代业务错误

#### Scenario: 异常前提建立方式

- **WHEN** 异常路径需要前提数据
- **THEN** 必须优先通过基线 seed、重置脚本与受控 API 操作建立
- **AND** 不得为某个异常路径新建独立 fixture SQL 文件
- **AND** 不得通过手工改库制造异常前提

### Requirement: canonical owner 页面、reread 页面与旧入口兼容必须被同时验证

系统 SHALL 在联调与验收中同时验证三类绑定动作的 canonical owner 页面、成功写入后的 reread 页面，以及旧页面/旧 transport 兼容链路，确保当前阶段不存在“看似可用、实则断链”的隐性阻断。

#### Scenario: 三类绑定动作的 canonical owner 与 reread 落点

- **WHEN** 任一绑定动作提交成功
- **THEN** 必须明确其 canonical owner 页面与 reread 页面是同一页：
  - `BindModuleToProduct` -> `ProductDetailPage`
  - `BindRepositoryToProduct` -> `RepositoryBindingDetailPage`
  - `MapModuleToRepository` -> `RepositoryBindingDetailPage`
- **AND** 成功结果必须由详情 reread 驱动展示
- **AND** 不得仅靠 toast、局部假状态或返回列表宣告成功

#### Scenario: 旧 `/api/candidates/*` 兼容读取验证

- **WHEN** 执行 `phase04-14` 兼容链路验收
- **THEN** 仍保留的旧 `/api/candidates/products` 与 `/api/candidates/repositories` 必须能够返回兼容结果
- **AND** 兼容结果必须来自 `productregistry` / `repositorybinding` 的 canonical query service 委派
- **AND** 不得让 `moduleregistry` 重新持有候选读取业务 owner

#### Scenario: Module Detail 旧入口兼容跳转验证

- **WHEN** 用户在 `Module Detail` 中使用旧绑定入口
- **THEN** 页面只能跳入 `Product Registry` 或 `Repository Binding` 正式主线
- **AND** 必须携带 `fromModuleDetail` 与必要上下文参数
- **AND** 不得继续展示候选读取、选择器、提交按钮组成的第二主工作台

### Requirement: 合同、传输层、前端适配层与正式规格必须一致

系统 SHALL 在联调与验收中同时核对 `.proto` 合同源、HTTP 过渡传输层、前端适配层和正式规格正文，确保当前阶段不存在第二套合同语义或未收口的规格漂移。

#### Scenario: 合同与 HTTP 语义核对

- **WHEN** 执行接口联调验收
- **THEN** 必须以 `.proto` 作为当前阶段唯一合同源
- **AND** HTTP 请求与响应语义必须与 `.proto` 对齐
- **AND** 前端 `types / api-adapter / product-registry-adapter / repository-binding-adapter` 的对象语义必须与 HTTP / `.proto` 对齐
- **AND** 不得在联调过程中发现第二套并列 JSON 合同语义

#### Scenario: 正式规格正文一致性核对

- **WHEN** 联调中发现 `phase04-10` 正式规格、`phase04-12 / 13` 实现 spec 与仓库现实不一致
- **THEN** 必须在当前阶段给出单值结论并完成收口
- **AND** 不得把“formal spec 未同步”“实现已改但上游未回写”视为可忽略瑕疵
- **AND** 当前阶段通过的前提是：`phase04-10` 正式规格正文与已验收实现边界一致，不再保留误导后续阶段的旧承诺

### Requirement: 验收结果必须形成可重复复核证据

系统 SHALL 为 `phase04-14` 输出可追踪、可重复复核的验收结果，而不是停留在口头“已通过”。

#### Scenario: 证据收口

- **WHEN** 完成联调与验证
- **THEN** 必须明确每条最小验收路径的结果
- **AND** 必须明确环境重建入口、执行顺序、关键请求/响应结果与页面行为结果
- **AND** 必须记录发现的问题、修复结论与剩余风险
- **AND** 未解决问题不得以“后续再看”形式遗留到阶段收口之后

## MODIFIED Requirements

### Requirement: phase04 当前阶段完成条件

系统 SHALL 将 `phase04` 当前阶段的完成条件从“合同、后端、前端都已实现”推进为“合同、后端、前端已经在同一环境中被验证可运行，且本阶段发现的问题已全部收口”。

#### Scenario: phase04-14 进入统一验收链

- **WHEN** `phase04-14` 开始执行
- **THEN** `phase04-11 / 12 / 13` 的成果必须进入统一联调
- **AND** 当前阶段 DoD 必须以可重复运行证据而不是实现声明为准
- **AND** `phase04-14` 不仅验证功能可跑通，还必须验证当前阶段不存在隐性阻断、兼容断链或正式规格漂移

## REMOVED Requirements

### Requirement: 仅以单侧 build 通过、局部接口成功或页面可见视为 phase04 验收完成

**Reason**: 单侧 build、静态页面可见或局部接口成功，并不能证明 `Product Registry + Repository Binding` 当前阶段最小主线已经形成完整可运行交付物。

**Migration**: 改为以前后端联调、冷启动路径、关键异常路径、canonical owner reread、兼容入口与合同一致性共同作为 `phase04-14` 的验收标准。
