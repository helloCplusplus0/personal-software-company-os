# Tasks

- [x] Task 1: 对齐上游冻结边界并确定 phase04 Proto 合同源布局
  - [x] SubTask 1.1: 对齐 `phase04-04` 的最小数据读写范围、详情/候选边界与错误语义前提
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小消息结构必须覆盖当前动作矩阵」L89-128 + 「Repository Binding 最小消息结构必须覆盖当前动作矩阵」L130-173 + 「错误语义必须在合同设计中保留显式映射前提」L333-349
  - [x] SubTask 1.2: 对齐 `phase04-07` 的后端模块边界、方向级 API 矩阵、canonical owner 与 reread owner，含四条跨模块关系摘要读取链路
    - 证据：`spec.md` §ADDED Requirements「服务接口必须对齐 phase04-07 已冻结的接口分组」L264-301 + 「详情读取与候选读取必须保持消息边界分离」Scenario: 详情内已绑定列表摘要消息与 phase04-07 跨模块关系摘要读取接口的对应关系 L321-331
  - [x] SubTask 1.3: 冻结 `common / product_registry / repository_binding` 三段式 Proto 布局、包名、版本语义与文件落点
    - 证据：`spec.md` §ADDED Requirements「phase04 Proto 合同源必须进入现有单一 proto workspace」L33-43 + 「Proto 包名、目录与版本语义必须冻结」L45-68

- [x] Task 2: 冻结 Product Registry 最小 Proto 消息结构
  - [x] SubTask 2.1: 冻结 `Product / ProductListItem / ProductDetail / BoundModuleSummary / BoundRepositorySummary`
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小消息结构必须覆盖当前动作矩阵」Scenario: Product Registry 核心对象与读组消息 L93-112
  - [x] SubTask 2.2: 冻结 `ListProducts / GetProductDetail / CreateProduct / BindModuleToProduct / ListProductModuleCandidates` 的 request / response
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小消息结构必须覆盖当前动作矩阵」Scenario: Product Registry 写组与候选读取消息 L114-128
  - [x] SubTask 2.3: 冻结 `Product Registry` 字段编号方案、写组返回语义与页面动作映射
    - 证据：`spec.md` §ADDED Requirements「字段编号方案必须在当前阶段冻结」Scenario: Product Registry 枚举与核心对象字段编号冻结 L210-218 + Scenario: Product Registry request / response 字段编号冻结 L220-233 + 「字段语义与页面动作必须单值映射」Scenario: Product Registry 字段与页面动作映射 L179-187

- [x] Task 3: 冻结 Repository Binding 最小 Proto 消息结构
  - [x] SubTask 3.1: 冻结 `Repository / RepositoryListItem / RepositoryDetail / BoundProductSummary / MappedModuleSummary`
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小消息结构必须覆盖当前动作矩阵」Scenario: Repository Binding 核心对象与读组消息 L134-153
  - [x] SubTask 3.2: 冻结 `ListRepositories / GetRepositoryDetail / CreateRepository / BindRepositoryToProduct / MapModuleToRepository / ListRepositoryProductCandidates / ListRepositoryModuleCandidates` 的 request / response
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小消息结构必须覆盖当前动作矩阵」Scenario: Repository Binding 写组与候选读取消息 L155-173
  - [x] SubTask 3.3: 冻结 `Repository Binding` 字段编号方案、写组返回语义与页面动作映射
    - 证据：`spec.md` §ADDED Requirements「字段编号方案必须在当前阶段冻结」Scenario: Repository Binding 核心对象字段编号冻结 L235-243 + Scenario: Repository Binding request / response 字段编号冻结 L245-262 + 「字段语义与页面动作必须单值映射」Scenario: Repository Binding 字段与页面动作映射 L189-197

- [x] Task 4: 冻结跨包依赖、共享枚举与服务接口矩阵
  - [x] SubTask 4.1: 冻结 `psco.common.v1.ActiveArchivedStatus` 共享枚举，避免 `Product / Repository` 等价枚举重复定义，明确 UNSPECIFIED = 不过滤语义与 create request 不允许 UNSPECIFIED
    - 证据：`spec.md` §ADDED Requirements「公共 active / archived 状态枚举必须单值化」L70-87（含 L85 UNSPECIFIED = 不过滤语义、L86 create request 不允许 UNSPECIFIED、L87 与 phase02-11A / phase03-11 模式对齐、L82-83 ActiveArchivedStatus 与 ModuleStatus 覆盖范围区分）
  - [x] SubTask 4.2: 冻结 `psco.module_registry.v1.ModuleStatus` 的 import 复用策略
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小消息结构必须覆盖当前动作矩阵」L110（`BoundModuleSummary.module_status` 通过 import 直接复用 `psco.module_registry.v1.ModuleStatus`）+ 「Repository Binding 最小消息结构必须覆盖当前动作矩阵」L152（`MappedModuleSummary.module_status` 通过 import 直接复用），承接 `phase03-11` decision_center.proto 跨包 import 模式
  - [x] SubTask 4.3: 冻结 `ProductRegistryService` 与 `RepositoryBindingService` 的最小 RPC 矩阵
    - 证据：`spec.md` §ADDED Requirements「服务接口必须对齐 phase04-07 已冻结的接口分组」Scenario: Product Registry service 矩阵 L268-277 + Scenario: Repository Binding service 矩阵 L279-290
  - [x] SubTask 4.4: 冻结 `.proto` Summary 消息与 `phase04-07` 四个跨模块关系摘要读取接口的对应关系，明确跨模块读接口不暴露为 RPC
    - 证据：`spec.md` §ADDED Requirements「详情读取与候选读取必须保持消息边界分离」Scenario: 详情内已绑定列表摘要消息与 phase04-07 跨模块关系摘要读取接口的对应关系 L321-331
  - [x] SubTask 4.5: 冻结 phase02 已落地 RPC 与 phase04 新增 RPC 的迁移承接，明确共存期间 canonical owner 与 breaking 兼容策略
    - 证据：`spec.md` §ADDED Requirements「服务接口必须对齐 phase04-07 已冻结的接口分组」Scenario: phase02 已落地 RPC 与 phase04 新增 RPC 的迁移承接 L292-301
  - [x] SubTask 4.6: 冻结排序规则不进入 `.proto` 合同本体的边界，承接现有 `decision_center.proto` 已建立模式
    - 证据：`spec.md` §ADDED Requirements「字段语义与页面动作必须单值映射」Scenario: 排序规则不进入 .proto 合同本体 L199-204

- [x] Task 5: 冻结 Proto 与 chi + JSON HTTP 的显式映射策略
  - [x] SubTask 5.1: 冻结 `productId / repositoryId / moduleId` 的 URL path → Proto request 字段组装规则
    - 证据：`spec.md` §ADDED Requirements「chi + JSON HTTP 必须从 Proto 单向承接」Scenario: 路径参数与消息字段映射 L362-367
  - [x] SubTask 5.2: 冻结 `Product Registry` 的 RPC → HTTP 映射矩阵
    - 证据：`spec.md` §ADDED Requirements「RPC 到 HTTP 的映射矩阵必须明确」Scenario: Product Registry 映射矩阵 L373-381
  - [x] SubTask 5.3: 冻结 `Repository Binding` 的 RPC → HTTP 映射矩阵
    - 证据：`spec.md` §ADDED Requirements「RPC 到 HTTP 的映射矩阵必须明确」Scenario: Repository Binding 映射矩阵 L383-393
  - [x] SubTask 5.4: 明确手写 JSON DTO 只能从 `.proto` 单向承接，不得形成第二套合同源
    - 证据：`spec.md` §ADDED Requirements「chi + JSON HTTP 必须从 Proto 单向承接」Scenario: 过渡层保留 L355-360 + §REMOVED Requirements L457-463

- [x] Task 6: 冻结合同演进与 Buf 校验规则
  - [x] SubTask 6.1: 冻结 `reserved`、递增编号与 `v1 -> v2` breaking 升级边界
    - 证据：`spec.md` §ADDED Requirements「合同演进与 breaking check 规则必须冻结」Scenario: 字段与枚举演进规则 L399-404 + Scenario: 删除字段或枚举值后的 reserved 约束 L406-411 + Scenario: v1 breaking 升级边界 L413-421
  - [x] SubTask 6.2: 冻结 `buf build / lint / generate / breaking` 的最小校验链，复用现有 `proto/Makefile`
    - 证据：`spec.md` §ADDED Requirements「合同演进与 breaking check 规则必须冻结」Scenario: Buf 校验链冻结 L423-430（复用仓库现有 `proto/Makefile`，`buf breaking` 对照 `../.git#branch=main,subdir=proto`）
  - [x] SubTask 6.3: 复核当前 spec 未把完整 gRPC / Connect 迁移、DTO 替换或传输层重写提前纳入当前阶段
    - 证据：`spec.md` §ADDED Requirements「当前阶段合同落地边界必须明确」L432-442（当前阶段可以不完成完整 gRPC / Connect / 网关迁移，可以继续保留 `chi` 作为 HTTP 过渡传输层，不要求立即用生成类型替换全部手写 DTO）

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 4`, and `Task 5`
