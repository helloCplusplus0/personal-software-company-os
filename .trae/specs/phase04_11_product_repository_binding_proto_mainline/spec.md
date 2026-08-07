# Phase04-11 Product / Repository / Binding 最小 Protocol Buffers 合同主线 Spec

> 状态说明：`phase04` 已完成收口，项目当前根级阶段已切换到 `phase05_dashboard_feedback_foundation`。本文件保留为 `phase04-11` 的合同主线规格说明，不单独承担根级当前阶段状态表达。

## Why

`phase04-08` 已经冻结了 `Product / Repository / Binding` 的最小 `.proto` 合同设计，`phase04-10` 又把这些结论收口进当前阶段正式规格正文。但截至现在，仓库主线里仍只有 `Module Registry` 与 `Decision Center` 已进入现有 `proto/` workspace，`Product Registry` 与 `Repository Binding` 还没有真正落地到仓库内的单一合同源、生成入口与 `buf` 校验链中。若继续围绕手写 DTO / JSON HTTP 结构推进实现，后续很容易再次长出并列合同源。

因此，`phase04-11` 必须把 `phase04-08` 已冻结的设计推进成仓库内实际存在、可生成、可校验、可被后续后端与前端直接消费的 `.proto` 主线，并把 DTO/HTTP 过渡层与 `.proto` 的单向映射边界写成单值结论，为 `phase04-12 / 13` 与后续验收提供可执行的合同基线。

## What Changes

- 将 `phase04-08` 已冻结的 `Product / Repository / Binding` 最小 `.proto` 合同正式落地到现有 `proto/` workspace
- 冻结 `proto/psco/common/v1/common.proto`、`proto/psco/product_registry/v1/product_registry.proto`、`proto/psco/repository_binding/v1/repository_binding.proto` 的单一文件落点、包名、跨包 import 与生成产物归属
- 冻结 `buf build / lint / generate / breaking` 在当前仓库中的最小可运行入口，并要求继续复用现有 `proto/buf.yaml`、`proto/buf.gen.yaml`、`proto/Makefile`
- 冻结 `Product Registry` 与 `Repository Binding` 的 HTTP DTO / adapter / handler 对 `.proto` 的单向语义映射边界，阻断第二套合同源
- 明确当前阶段只落地合同源、生成入口、校验链与映射边界，不提前完成完整 gRPC / Connect 传输层迁移
- **BREAKING**：后续 `Product / Repository / Binding` 的实现与验收不得再把手写 JSON 结构、页面自定义字段或 handler DTO 视为并列合同源，`.proto` 成为仓库内唯一合同定义入口

## Impact

- Affected specs:
  - `phase04_08_product_repository_binding_proto_contract`
  - `phase04_10_product_repository_binding_formal_spec`
  - 后续 `phase04-12` 后端与数据主线
  - 后续 `phase04-13` 前端主线
  - 后续 `phase04` 验收与收口
  - `phase02_11a_module_registry_proto_contract`
  - `phase03_11_decision_center_proto_mainline`
- Affected code:
  - `proto/psco/common/v1/common.proto`
  - `proto/psco/product_registry/v1/product_registry.proto`
  - `proto/psco/repository_binding/v1/repository_binding.proto`
  - `proto/buf.yaml`
  - `proto/buf.gen.yaml`
  - `proto/Makefile`
  - `proto/README.md`
  - `backend/internal/gen/proto/psco/common/v1/`
  - `backend/internal/gen/proto/psco/product_registry/v1/`
  - `backend/internal/gen/proto/psco/repository_binding/v1/`
  - `frontend/src/gen/proto/psco/common/v1/`
  - `frontend/src/gen/proto/psco/product_registry/v1/`
  - `frontend/src/gen/proto/psco/repository_binding/v1/`
  - `backend/internal/productregistry/types.go`
  - `backend/internal/repositorybinding/types.go`
  - `frontend/src/features/product-registry/`
  - `frontend/src/features/repository-binding/`

## ADDED Requirements

### Requirement: Product / Repository / Binding 合同必须落地到现有 proto workspace

系统 SHALL 将 `Product / Repository / Binding` 的最小 `.proto` 合同落地到现有 `proto/` workspace，而不是为 `phase04` 新建第二个 proto 根目录、第二套 `buf.yaml` 或第二套生成入口。

#### Scenario: 合同源文件落点

- **WHEN** 执行 `phase04-11`
- **THEN** 仓库中必须存在：
  - `proto/psco/common/v1/common.proto`
  - `proto/psco/product_registry/v1/product_registry.proto`
  - `proto/psco/repository_binding/v1/repository_binding.proto`
- **AND** 上述文件必须作为 `Product / Repository / Binding` 当前阶段唯一合同定义入口
- **AND** 必须继续复用现有 `proto/buf.yaml` 作为同一个 buf workspace 的根配置
- **AND** 不得在 `backend/`、`frontend/` 或仓库根新增并列 proto 工作区

#### Scenario: 包名与版本语义落地

- **WHEN** `Product / Repository / Binding` 合同正式进入仓库
- **THEN** 包名与版本语义必须冻结为：
  - `psco.common.v1`
  - `psco.product_registry.v1`
  - `psco.repository_binding.v1`
- **AND** 后续字段演进必须继续在这些 `v1` 版本语义下进行
- **AND** 不得在落地阶段临时改写为第二套包名、目录层级或版本策略

### Requirement: 跨包依赖必须直接复用现有 Proto 主线

系统 SHALL 要求 `Product / Repository / Binding` 在合同落地时直接复用现有 proto 主线中的共享枚举与跨模块状态枚举，保证跨模块状态语义继续保持单值一致。

#### Scenario: 公共 active / archived 状态枚举落地

- **WHEN** 落地 `Product`、`Repository`、`RepositoryProductCandidate`、`BoundRepositorySummary`、`BoundProductSummary` 等 `status`
- **THEN** `common.proto` 必须定义并承接单一枚举 `psco.common.v1.ActiveArchivedStatus`
- **AND** 不得在 `product_registry.proto` 或 `repository_binding.proto` 中重定义本地等价 `active / archived` 枚举
- **AND** `ACTIVE_ARCHIVED_STATUS_UNSPECIFIED` 在 `status_filter` 中只表示“不过滤”，不得作为写组 request 的合法写入值

#### Scenario: ModuleStatus 依赖落地

- **WHEN** 落地 `BoundModuleSummary.module_status`、`MappedModuleSummary.module_status`、`ProductModuleCandidate.module_status`、`RepositoryModuleCandidate.module_status`
- **THEN** `product_registry.proto` 与 `repository_binding.proto` 必须通过 `import "psco/module_registry/v1/module_registry.proto"` 直接复用 `psco.module_registry.v1.ModuleStatus`
- **AND** 不得在 `psco.product_registry.v1` 或 `psco.repository_binding.v1` 中重定义本地等价枚举
- **AND** 该 import 必须在同一个 `proto/` workspace 下通过 `buf build` 成功解析

### Requirement: buf 工具链入口必须可运行且复用现有入口

系统 SHALL 为 `Product / Repository / Binding` 合同落地继续复用并扩展当前仓库已有的 `buf` 工具链入口，使 `build / lint / generate / breaking` 可以在同一 proto workspace 中执行，而不是为单个模块发明第二套脚本体系。

#### Scenario: build lint generate 入口

- **WHEN** 后续实现者在 `proto/` 目录校验 `Product / Repository / Binding` 合同
- **THEN** 必须能够通过既有入口运行 `buf build`、`buf lint` 与 `buf generate`
- **AND** 这些入口必须同时覆盖 `Module Registry`、`Decision Center`、`Product Registry` 与 `Repository Binding`
- **AND** 不得要求实现者手工拼接单文件命令绕过 `proto/buf.yaml` 或 `proto/buf.gen.yaml`

#### Scenario: breaking 基准路径

- **WHEN** 后续实现者或 CI 对 `Product / Repository / Binding` 合同执行破坏性变更校验
- **THEN** 必须通过既有 `proto/Makefile` 或等价受控入口运行 `buf breaking`
- **AND** `buf breaking` 必须直接对照仓库主线 Git 基准，路径口径与 `proto/` 子目录保持一致
- **AND** `buf breaking` 的基准路径必须继续冻结为 `../.git#branch=main,subdir=proto`
- **AND** 失败时必须保留非零退出码，不得吞掉错误
- **AND** 不得改为对临时导出文件、临时副本目录或手工拼接镜像做 breaking 基准

### Requirement: 生成产物落点必须与现有 proto 主线同构

系统 SHALL 要求 `Product / Repository / Binding` 的代码生成产物继续落在现有 Go / TypeScript 生成目录主线上，使后续后端与前端实现可以复用既有生成模式。

#### Scenario: Go 与 TypeScript 生成产物

- **WHEN** 对 `Product / Repository / Binding` 合同执行 `buf generate`
- **THEN** Go 生成产物必须落在：
  - `backend/internal/gen/proto/psco/common/v1/`
  - `backend/internal/gen/proto/psco/product_registry/v1/`
  - `backend/internal/gen/proto/psco/repository_binding/v1/`
- **AND** TypeScript 生成产物必须落在：
  - `frontend/src/gen/proto/psco/common/v1/`
  - `frontend/src/gen/proto/psco/product_registry/v1/`
  - `frontend/src/gen/proto/psco/repository_binding/v1/`
- **AND** 当前阶段只要求生成消息类型与 service 定义对应的最小合同产物
- **AND** 不得因为 `phase04-11` 额外引入完整 gRPC 服务端、Connect 网关或第二套客户端生成主线

### Requirement: Product Registry 过渡传输层必须从 proto 单向承接

系统 SHALL 冻结 `Product Registry` 的 DTO / HTTP adapter / handler 与 `.proto` 的关系为“`.proto` 单向定义合同，过渡传输层显式映射承接”，不允许继续并列扩张第二套字段语义。

#### Scenario: Product Registry DTO 与 adapter 映射边界

- **WHEN** 后续实现者编写 `backend/internal/productregistry/types.go`、HTTP handler DTO 或前端 adapter
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 显式对齐
- **AND** 不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义
- **AND** 路径参数、时间格式、枚举字符串形式等传输差异必须被明确视为适配层差异，而不是第二套合同定义

#### Scenario: Product Registry URL 路径参数与 Proto request 组装

- **WHEN** HTTP 过渡层通过 URL 路径参数承接 `productId`
- **THEN** handler 必须在进入业务层前显式组装为对应的 Proto request 字段
- **AND** 这种组装行为只属于过渡传输层
- **AND** 不得因为 URL 参数存在而在 `.proto` 之外保留另一份请求合同

### Requirement: Repository Binding 过渡传输层必须从 proto 单向承接

系统 SHALL 冻结 `Repository Binding` 的 DTO / HTTP adapter / handler 与 `.proto` 的关系为“`.proto` 单向定义合同，过渡传输层显式映射承接”，不允许继续并列扩张第二套字段语义。

#### Scenario: Repository Binding DTO 与 adapter 映射边界

- **WHEN** 后续实现者编写 `backend/internal/repositorybinding/types.go`、HTTP handler DTO 或前端 adapter
- **THEN** JSON 请求与响应语义必须从 `.proto` 派生或与 `.proto` 显式对齐
- **AND** 不得在 `types.go`、handler DTO 或前端页面层私自新增 `.proto` 中不存在的业务字段语义
- **AND** 路径参数、时间格式、枚举字符串形式等传输差异必须被明确视为适配层差异，而不是第二套合同定义

#### Scenario: Repository Binding URL 路径参数与 Proto request 组装

- **WHEN** HTTP 过渡层通过 URL 路径参数承接 `repositoryId`
- **THEN** handler 必须在进入业务层前显式组装为对应的 Proto request 字段
- **AND** 这种组装行为只属于过渡传输层
- **AND** 不得因为 URL 参数存在而在 `.proto` 之外保留另一份请求合同

### Requirement: 当前阶段合同落地边界必须保持最小化

系统 SHALL 明确 `phase04-11` 的目标是“合同源落地 + 工具链入口 + 映射边界”，而不是把整个 `Product / Repository / Binding` 传输栈一次性改写完成。

#### Scenario: 当前阶段允许保留的实现边界

- **WHEN** 执行 `phase04-11`
- **THEN** 当前阶段可以继续保留 `chi + JSON HTTP` 作为过渡传输层
- **AND** 当前阶段可以只生成消息类型和最小合同产物
- **AND** 当前阶段不要求完成完整 gRPC / Connect 传输层迁移
- **AND** 当前阶段不要求立即用生成类型替换全部现有手写 DTO

## MODIFIED Requirements

### Requirement: Product / Repository / Binding Proto 合同源从“设计冻结”推进为“仓库主线落地”

系统 SHALL 将 `phase04-08` 中已经冻结的 `Product / Repository / Binding` Proto 设计，从“规格层定义”推进为“仓库内实际存在且可被 buf 工具链消费的单一合同源”。

#### Scenario: 合同阶段推进

- **WHEN** `phase04-11` 开始执行
- **THEN** `Product / Repository / Binding` `.proto` 不再只停留在 `phase04-08` 与 `phase04-10` 文档正文中
- **AND** 必须在仓库 `proto/` workspace 内拥有实际文件落点、生成入口与校验入口
- **AND** 后续 `phase04-12 / 13` 与验收必须优先引用该已落地合同源，而不是回到文档层手工解释字段

### Requirement: buf 校验链从“基线要求”推进为“仓库执行入口”

系统 SHALL 将 `phase04` 共享基线与正式规格正文中关于 `buf build / lint / generate / breaking` 的要求，从抽象校验前提推进为仓库中的受控执行入口。

#### Scenario: 工具链收口

- **WHEN** `phase04-11` 完成
- **THEN** `buf` 校验链必须能够在仓库中通过受控入口执行
- **AND** 这些入口必须与现有 `proto/` workspace 保持单一真相源
- **AND** 不得继续保留“工具链要求存在于文档中，但仓库内没有对应执行入口”的状态

## REMOVED Requirements

### Requirement: Product / Repository / Binding 合同只停留在 formal spec 与子规格层

**Reason**: `phase04-11` 的目标就是把 `Product / Repository / Binding` 合同从“已经设计好”推进到“仓库中已经落地、可生成、可校验、可被实现引用”的状态。

**Migration**: 后续 `Product / Repository / Binding` 的实现、联调与验收应统一从 `proto/psco/common/v1/common.proto`、`proto/psco/product_registry/v1/product_registry.proto`、`proto/psco/repository_binding/v1/repository_binding.proto` 与现有 `proto/` 工具链入口进入；`phase04-08` 和 `phase04-10` 继续作为设计与正式规格上游，不再承担仓库内合同主线入口职责。
