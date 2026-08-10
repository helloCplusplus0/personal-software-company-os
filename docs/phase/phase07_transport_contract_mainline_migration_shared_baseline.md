# phase07_transport_contract_mainline_migration_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase07` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase07_transport_contract_mainline_migration`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase07` 已建立正式 `/plan` 入口，作为 `mvp0.3` 业务阶段之前的前置基础 phase
- `phase07` 的规划上游统一以 `PSCO-mvp03-summarize-feedback.md` 与 `audit_001` 审计结论为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp03-summarize-feedback.md`
  - `docs/audit/audit_001_transport_contract_mainline_issue.md`
  - `docs/audit/audit_001_transport_contract_mainline_analysis.md`
  - `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
  - `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/`
  - `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md`
- 当前阶段只承接 `phase01 ~ phase06` 已冻结并验收的 canonical 业务主线
- 当前阶段不反向重写 `mvp0.3` 已冻结的业务方向排序

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web`
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http + ConnectRPC`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- Contract Tooling：`buf build / lint / generate / breaking`
- TS Generation：`bufbuild/es`
- Go Generation：`protocolbuffers/go + connectrpc/go`
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段新增或迁移后的业务接口必须遵守：
  - `.proto` 是唯一长期合同源
  - `ConnectRPC` 是业务接口正式传输层
  - `chi` 只承担 router shell、middleware 与非业务端点承载职责
- 当前阶段浏览器侧访问业务接口的外部前缀继续统一冻结为 `/api`
- 当前阶段 Vite dev proxy、Caddy 与本地启动链只允许暴露单一 `/api` 基址，不允许因 Connect 迁移长出并列业务前缀
- 当前阶段允许保留在 `chi + net/http` 的端点只有：
  - `healthz`
  - `readyz`
  - `metrics`
  - `debug / pprof`
- 当前阶段不允许：
  - phase 收口后继续保留 hand-written JSON 业务主线
  - 把迁移过程短时 adapter 解释为长期兼容层
  - 把基础设施端点一并强行纳入 `.proto`
  - 让 `chi` middleware 与 Connect interceptor 长出两套长期并列的横切逻辑体系

### 2.5 当前阶段交付模式

- `phase07` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 ConnectRPC 业务传输主线代码，并完成旧 JSON 业务主线退场

## 3. 当前阶段动作矩阵

`phase07` 最少需要直接承接：

- `GenerateConnectGoArtifacts`
- `GenerateFrontendConnectClients`
- `MountConnectHandlersOnChi`
- `MigrateModuleRegistryTransport`
- `MigrateDecisionCenterTransport`
- `MigrateProductRegistryTransport`
- `MigrateRepositoryBindingTransport`
- `MigrateDashboardFeedbackTransport`
- `MigrateOnboardingExportBackupReuseTransport`
- `RunPhase01To06Regression`
- `RetireLegacyJsonBusinessHandlers`

当前阶段必须打通的最小闭环：

- `.proto -> buf generate -> Go Connect / TS client 产物`
- `chi mount(/api/*) -> Connect handler -> service implementation`
- `frontend generated client -> query / application owner`
- `Vite / 本地启动 / Caddy / 验收脚本 -> 单一 /api 基址`
- `phase01 ~ phase06 业务链路联调回归`
- `旧 JSON 业务主线退场`

允许以最小连接位承接但不扩写为独立主线：

- 迁移过程中的临时 adapter
- 必要的回退脚本或排障开关

## 4. 当前阶段页面与业务矩阵

本阶段必须保持可用并完成传输切换的业务页面 / 模块：

- `Module Registry`
- `Decision Center`
- `Product Registry`
- `Repository Binding`
- `Dashboard Home`
- `Onboarding`
- `Export`
- `Backup`
- `Reuse Summary`

### 4.1 当前阶段交互归属矩阵

- `Module / Decision / Product / Repository`：继续承接 canonical owner 业务写入与读取
- `Dashboard`：继续承接经营入口与反馈可见面
- `Onboarding`：继续承接首轮引导与状态读取
- `Export / Backup`：继续承接数据主权相关业务动作
- `Reuse Summary`：继续承接派生读模型消费

本阶段变化不在于页面职责重写，而在于这些页面背后的业务接口传输主线切换。

## 5. 当前阶段合同与传输矩阵

### 5.1 当前阶段必须完成正式切换的业务模块

- `psco.module_registry.v1`
- `psco.decision_center.v1`
- `psco.product_registry.v1`
- `psco.repository_binding.v1`
- `psco.dashboard.v1`
- `psco.onboarding.v1`
- `psco.export.v1`
- `psco.backup.v1`
- `psco.reuse_summary.v1`

### 5.2 当前阶段最小接口归属前提

- Go 业务接口由 Connect 生成 handler / procedure 常量与 service interface 承接
- 前端业务调用由生成客户端或等价 Connect transport adapter 承接
- canonical 业务接口迁移清单必须下钻到 `service / RPC / 当前入口路径 / 页面动作 owner`
- `chi` 只负责：
  - route mount
  - middleware
  - 非业务端点
- Connect interceptor 只承接 RPC 级校验、错误归一化与必要元数据处理
- domain error、proto error code 与 Connect error 维持单值映射
- phase 收口后不再保留 hand-written JSON business handler 作为正式主线

### 5.3 当前阶段生成链基线

- `proto/buf.gen.yaml` 必须至少覆盖：
  - `buf.build/protocolbuffers/go`
  - `buf.build/connectrpc/go`
  - `buf.build/bufbuild/es`
- `proto/Makefile` 的 `build / gen / lint / breaking` 入口继续保留
- 验收脚本、本地启动脚本与 CI 中对 proto 生成链的入口必须与 `proto/Makefile` 保持一致
- Go 与前端生成产物落点必须继续保持单一入口

## 6. 当前阶段验收基线

- `phase01 ~ phase06` canonical 业务接口均已切到 ConnectRPC
- `phase01 ~ phase06` 的 canonical 迁移总表已覆盖到 `service / RPC / 当前入口路径 / 页面动作 owner`
- `healthz / readyz / metrics / debug` 继续稳定留在 `chi + net/http`
- 根级规则、`audit_001` 与 `phase07` 三件套口径一致
- 前端业务切片不再依赖手写 JSON 主线
- 开发环境、验收环境与部署链路均继续通过单一 `/api` 基址访问业务接口
- phase 收口后不存在“新 Connect 主线 + 旧 JSON 主线”并列正式状态

## 7. 当前阶段非目标

- 不在本阶段直接推进 `mvp0.3` 业务能力
- 不在本阶段扩写 AI、`Venture` 或新实体
- 不在本阶段引入第二套路由、第二套状态管理或第二套 UI 事实源
