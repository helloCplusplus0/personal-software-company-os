# Tasks

- [x] Task 1: 对齐上游边界并冻结后端 owner 矩阵
  - [x] SubTask 1.1: 对齐 `phase04-03` 的 canonical owner、候选读取迁移边界与 `phase02` 兼容前提
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接接口必须迁移为 canonical owner 下的兼容适配」L257-281 + 「phase02 历史绑定数据兼容前提冻结」L283-302
  - [x] SubTask 1.2: 对齐 `phase04-04` 的最小数据读写范围、错误语义与“本阶段不冻结实现工具”的边界
    - 证据：`spec.md` §ADDED Requirements「后端接口必须按读组与写组拆分」L56-78 + 「当前阶段不冻结的实现工具」L377-382
  - [x] SubTask 1.3: 将 `CreateProduct / CreateRepository / BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository` 的后端 canonical owner 与 reread owner 单值化
    - 证据：`spec.md` §ADDED Requirements「Product Registry 后端边界必须冻结为单一业务模块」L24-38 + 「Repository Binding 后端边界必须冻结为单一业务模块」L40-54 + 「Product Detail 与 Repository Detail 必须分别作为 reread owner」L110-133

- [x] Task 2: 冻结 Product Registry 与 Repository Binding 的后端接口分组
  - [x] SubTask 2.1: 冻结 `Product Registry` 的读组、写组与候选读取子组
    - 证据：`spec.md` §ADDED Requirements「后端接口必须按读组与写组拆分」Scenario: Product Registry 最小接口分组 L60-68
  - [x] SubTask 2.2: 冻结 `Repository Binding` 的读组、写组与候选读取子组
    - 证据：`spec.md` §ADDED Requirements「后端接口必须按读组与写组拆分」Scenario: Repository Binding 最小接口分组 L70-78
  - [x] SubTask 2.3: 明确详情读组与候选读取、创建写入与绑定写入之间的边界
    - 证据：`spec.md` §ADDED Requirements「后端接口必须按读组与写组拆分」L60-78（`ProductDetailRead` 只承接详情本体与已绑定列表、候选读取独立子包、创建与绑定写入不混成单接口）
  - [x] SubTask 2.4: 冻结 `Product Registry` 与 `Repository Binding` 的方向级 API 矩阵，承接 `phase02-09` §6.2 与 `phase03-10` §6.2 模式
    - 证据：`spec.md` §ADDED Requirements「方向级 API 矩阵冻结」Scenario: Product Registry 方向级 API 矩阵 L84-94 + Scenario: Repository Binding 方向级 API 矩阵 L96-108

- [x] Task 3: 冻结跨模块服务侧连接边界与兼容迁移策略
  - [x] SubTask 3.1: 明确两模块对 `Module Registry` 的最小只读连接边界
    - 证据：`spec.md` §ADDED Requirements「Module Registry 连接边界必须冻结为只读候选与存在性校验」L135-151
  - [x] SubTask 3.2: 明确 `Decision Center` 当前阶段保持后移，不进入正式绑定主线
    - 证据：`spec.md` §ADDED Requirements「Decision Center 连接边界必须保持后移」L246-254
  - [x] SubTask 3.3: 冻结 `phase02` 旧候选读取与旧绑定写入接口的兼容适配策略，不保留第二业务 owner
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接接口必须迁移为 canonical owner 下的兼容适配」L257-281 + §MODIFIED Requirements「phase02 临时绑定 owner 的解释」L397-405
  - [x] SubTask 3.4: 冻结 `phase02` 历史绑定数据（`product_modules` / `module_repositories`）兼容前提，承接 `phase04-03`
    - 证据：`spec.md` §ADDED Requirements「phase02 历史绑定数据兼容前提冻结」Scenario: 历史绑定数据可读性 L287-294
  - [x] SubTask 3.5: 冻结 `Module Detail` 旧入口后端 endpoint 兼容策略（不为 `Module Detail` 提供新绑定写入 endpoint，旧 endpoint 仅作兼容适配层委派）
    - 证据：`spec.md` §ADDED Requirements「phase02 历史绑定数据兼容前提冻结」Scenario: Module Detail 旧入口后端 endpoint 兼容策略 L296-302
  - [x] SubTask 3.6: 冻结四条跨模块已绑定关系摘要读取边界（`ProductRepositorySummaryRead` / `ProductModuleSummaryRead` / `RepositoryProductSummaryRead` / `RepositoryModuleSummaryRead`），承接 `phase04-04` 详情读取数据范围与 `phase04-03` 关系表 owner 冻结，每条链路单值化到 provider owner、接口名、字段范围、文件落点与注入方式
    - 证据：`spec.md` §ADDED Requirements「跨模块已绑定关系摘要读取边界冻结」L153-244（含 `ProductRepositorySummaryRead` L159-190 与三个新增 Scenario：`ProductModuleSummaryRead` L192-205、`RepositoryProductSummaryRead` L207-220、`RepositoryModuleSummaryRead` L222-235，以及整体边界隔离 L237-244）

- [x] Task 4: 冻结实现级包结构与关键文件落点
  - [x] SubTask 4.1: 为 `backend/internal/productregistry/` 冻结 `handler / service / repository / candidate` 结构与关键文件落点，含 `service/query_service.go` 通过注入的 `ProductModuleSummaryRead` 与 `ProductRepositorySummaryRead` 分别承接 `ProductDetailRead` 已绑定模块列表与已绑定仓库列表
    - 证据：`spec.md` §ADDED Requirements「后端模块文件落点必须冻结到实现结构层」Scenario: Product Registry 文件落点 L308-319
  - [x] SubTask 4.2: 为 `backend/internal/repositorybinding/` 冻结 `handler / service / repository / candidate` 结构与关键文件落点，含 `repository/binding_store.go` 承接 `ProductRepositorySummaryRead` / `RepositoryProductSummaryRead` / `RepositoryModuleSummaryRead` 三个跨模块读接口实现，`service/query_service.go` 通过注入的 `RepositoryProductSummaryRead` 与 `RepositoryModuleSummaryRead` 分别承接 `RepositoryDetailRead` 已绑定产品列表与已映射模块列表
    - 证据：`spec.md` §ADDED Requirements「后端模块文件落点必须冻结到实现结构层」Scenario: Repository Binding 文件落点 L321-333
  - [x] SubTask 4.3: 明确 `backend/internal/platform/router.go` 只作为装配点，不反向决定业务 owner
    - 证据：`spec.md` §ADDED Requirements「Chi 路由组织方式只作为实现兼容前提」L384-392
  - [x] SubTask 4.4: 冻结分层语义（入口层 / 业务编排层 / 数据访问层）与读组写组单文件编排原则，承接 `phase02-09` §9.4 与 `phase03-10` §10.4
    - 证据：`spec.md` §ADDED Requirements「后端模块文件落点必须冻结到实现结构层」Scenario: 分层语义与单文件编排原则 L335-344
  - [x] SubTask 4.5: 冻结支撑文件落点（`errors.go / types.go / validate.go / handler/response.go`），与 `moduleregistry` / `decisioncenter` 同构
    - 证据：`spec.md` §ADDED Requirements「后端模块文件落点必须冻结到实现结构层」Scenario: 支撑文件落点 L346-356
  - [x] SubTask 4.6: 冻结候选读取接口拥有者与接线原则（`candidate/` 子包拥有、`service/` 不得直接写跨模块 SQL、装配点注入），承接 `phase03-10` §10.5 与 `phase04-03`，且该接线原则同样适用于 `ProductRepositorySummaryRead` / `ProductModuleSummaryRead` / `RepositoryProductSummaryRead` / `RepositoryModuleSummaryRead` 四个跨模块关系摘要读取接口
    - 证据：`spec.md` §ADDED Requirements「后端模块文件落点必须冻结到实现结构层」Scenario: 候选读取接口拥有者与接线原则 L358-367
  - [x] SubTask 4.7: 明确 `product_repositories` 表为 `phase04` 新增表，`product_modules` / `module_repositories` 表已存在于 `phase02` migration 中迁移后表结构不变
    - 证据：`spec.md` §ADDED Requirements「后端模块文件落点必须冻结到实现结构层」Scenario: product_repositories 表为 phase04 新增 L369-375

- [x] Task 5: 复核规格边界与非目标冻结
  - [x] SubTask 5.1: 复核本 spec 未提前冻结 Go 数据访问层工具、HTTP/RPC 框架与最终 `.proto` 命名
    - 证据：`spec.md` §ADDED Requirements「后端模块文件落点必须冻结到实现结构层」Scenario: 当前阶段不冻结的实现工具 L377-382
  - [x] SubTask 5.2: 复核未把 `Decision Center`、`Module Registry` 扩写为并列绑定主线
    - 证据：`spec.md` §ADDED Requirements「Decision Center 连接边界必须保持后移」L246-254 + §REMOVED Requirements「Module Registry 继续作为绑定写入长期 owner」L409-413 + 「Decision Center 作为 Product / Repository 绑定主线参与者」L415-419
  - [x] SubTask 5.3: 复核 spec 结论足以直接进入后续实现与 `phase04-08` 合同设计
    - 证据：`spec.md` 全文模块边界、接口分组、方向级 API 矩阵、四条跨模块关系摘要读取边界（含每条链路 provider owner、接口名、字段范围、文件落点与注入方式单值化）、文件落点、分层语义、支撑文件、候选读取接线、兼容迁移策略均单值化，运行时实现细节显式不冻结

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
