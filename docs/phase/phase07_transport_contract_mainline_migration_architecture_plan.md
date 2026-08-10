# phase07_transport_contract_mainline_migration_architecture_plan

## 1. 文档定位

本文档是 `phase07_transport_contract_mainline_migration` 的架构规划文档。

`phase06` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，但 `audit_001` 已明确给出 `escalate-phase` 结论：在正式推进 `mvp0.3` 的 `Operating Review Loop / Template Reuse / Derived Intelligence / Dry-Run` 业务主线之前，必须先完成 Go 业务传输主线从“`.proto canonical + chi JSON 过渡映射`”到“`.proto + ConnectRPC`”的正式切换。

因此，`phase07` 不是偏离 `mvp0.3`，而是 `mvp0.3` 业务阶段的前置基础 phase。它的任务不是新增业务对象，而是完成 `phase01 ~ phase06` 已交付 canonical 业务接口的正式传输主线切换，并把旧的手写 `chi + JSON HTTP` 业务主线彻底退场。

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `docs/README.md`
7. `PSCO-mvp03-summarize-feedback.md`
8. `docs/audit/audit_001_transport_contract_mainline_issue.md`
9. `docs/audit/audit_001_transport_contract_mainline_analysis.md`
10. `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
11. `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/`
12. `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md`
13. `proto/README.md`
14. `proto/buf.gen.yaml`
15. `proto/Makefile`

补充说明：

- `PSCO-mvp03-summarize-feedback.md` 继续冻结 `mvp0.3` 的业务方向，但不直接替代本阶段 `/plan`
- `audit_001` 提供了本阶段必须存在的原因、最佳实践与去向结论
- `phase06` 的正式规格、合同主线与验收结论继续作为当前仓库 canonical 业务能力的直接交付上游

## 3. 本阶段目标

`phase07` 的目标是：

> 在不推翻当前 `chi` 装配层、不引入重型 gRPC 基础设施的前提下，将 `phase01 ~ phase06` 已落地的 canonical 业务接口一次性迁移到 `.proto + ConnectRPC` 主线，并完成前后端生成链、调用链与验收口径的同步切换。

本阶段需要回答的核心问题：

1. `chi` 在正式架构中的唯一职责是什么，哪些端点应明确保留在 `chi + net/http`
2. `buf.gen.yaml` 应如何补齐 Go Connect handler/client 与 TypeScript Connect 客户端生成
3. `phase01 ~ phase06` 的哪些业务接口属于必须在本阶段一次性切换的 canonical 主线
4. 前端应如何从手写 JSON adapter 切到生成客户端而不破坏既有 `query / application` 分层
5. 浏览器侧、Vite dev proxy、Caddy 与本地启动链应如何在不长出第二套访问路径的前提下承接 Connect procedure path
6. middleware、interceptor 与错误语义映射应如何在传输切换后保持单值
7. phase 收口时，如何证明旧手写 JSON 业务主线已经退场，而非仅仅增加了一套新实现

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase07` 必须直接承接：

- `PSCO-mvp03-summarize-feedback.md`
- `docs/audit/audit_001_transport_contract_mainline_issue.md`
- `docs/audit/audit_001_transport_contract_mainline_analysis.md`
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
- `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/`
- `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md`

不允许在本阶段重新解释：

- `.proto` 作为唯一长期合同源的地位
- `Decision` 作为业务主线的中心地位
- `mvp0.3` 的业务方向排序
- `chi` 仍可保留为装配层与非业务端点承载层这一冻结结论

### 4.2 当前阶段主交付对象

`phase07` 的主交付对象是：

- `ConnectRPC Transport Mainline`
- `Buf Generation Chain Upgrade`
- `Canonical Business API Migration`
- `Phase01 ~ Phase06 Regression Validation`

其最小主线必须优先承接：

- `Module Registry`
- `Decision Center`
- `Product Registry`
- `Repository Binding`
- `Dashboard + Feedback`
- `Onboarding`
- `Export`
- `Backup`
- `Reuse Summary`

当前阶段不把以下内容作为主交付对象：

- `Operating Review Loop`
- `Template Reuse`
- `Derived Intelligence Deepening`
- `Real-Project Dry-Run`
- 新的长期业务实体

### 4.3 当前阶段后端交付策略

后端继续统一遵守：

- `.proto` 是唯一长期合同源
- `chi` 作为顶层路由、middleware 与 mount shell
- `ConnectRPC` 作为业务接口正式传输层
- `healthz / readyz / metrics / debug` 继续留在 `chi + net/http`

当前阶段重点是：

- 使用 Connect 生成物提供 `net/http` handler，并通过 `chi` 挂载
- 不再新增手写 JSON 业务 handler / DTO 作为 canonical owner
- 把旧业务 handler 的职责回收到 Connect service implementation
- 让错误语义、response envelope 与 request contract 直接服从 `.proto`
- 保持既有 `chi` middleware 作为外层 HTTP shell，同时把 RPC 级校验、错误归一化与必要元数据处理收敛到 Connect interceptor 链

### 4.4 当前阶段前端交付策略

前端继续统一遵守：

- 单一 `React Web`
- `TanStack Router + TanStack Query`
- 业务写路径唯一 `application` 入口
- `query` 层纯只读
- mutation 固定承接位

当前阶段重点是：

- 从 `buf` 生成的 TypeScript / Connect 客户端消费业务接口
- 不再扩写手写 `fetch + JSON DTO` 作为业务主线
- 把调用方式切换限定在传输层与切片承接位，不破坏现有页面组织
- 保证 `phase01 ~ phase06` 页面与动作在传输切换后行为保持等价
- 保持浏览器侧单一 `/api` 访问前缀与既有开发环境代理入口，不允许因为 Connect 迁移长出第二套前端基址约定

### 4.5 当前阶段数据、合同与传输承接原则

当前阶段关于合同与传输的冻结如下：

- `.proto` 仍是唯一长期合同源
- `buf build / lint / generate / breaking` 继续作为合同工具链基线
- `buf.gen.yaml` 必须补齐：
  - `protocolbuffers/go`
  - `connectrpc/go`
  - `bufbuild/es`
- Go 侧业务接口必须由 Connect 生成 handler / procedure 常量与 service interface 承接
- TypeScript 侧业务调用必须由生成客户端或等价 Connect transport adapter 承接
- 浏览器侧对业务接口的访问前缀继续统一冻结为 `/api`
- Vite dev proxy、Caddy 与本地启动链继续只暴露单一 `/api` 基址；Connect procedure path 必须在该前缀下完成挂载，而不是新增并列基址
- 旧手写 `chi + JSON HTTP` 业务接口只允许在迁移过程短时存在，phase 收口前必须退场
- 非业务基础设施端点不纳入 `.proto`，继续保留原生 `chi + net/http`
- 现有 `chi` middleware 链必须继续作为统一 HTTP 外壳；Connect interceptor 只承接 RPC 级横切逻辑，不得复制第二套请求治理体系
- domain error -> proto error code -> Connect error 的映射必须单值化；不得同时保留两套长期错误语义

当前阶段不在架构规划中提前冻结：

- 后续 `mvp0.3` 业务 phase 的正式命名
- 新业务阶段的具体接口名
- Connect 拦截器链的最终完整实现细节

### 4.6 当前阶段迁移边界原则

为了避免 `phase07` 变成“加一套 Connect 试点，同时旧主线继续长期保留”的半切换状态，当前阶段先冻结以下迁移边界：

- 本阶段迁移的是 `phase01 ~ phase06` 已交付的 canonical 业务接口，而不是只做新接口试点
- 本阶段迁移清单必须下钻到 `service / RPC / 当前外部访问路径 / 页面或动作 owner` 级别，不得只停留在模块名层级
- 本阶段不要求迁移所有 HTTP 端点；仅要求迁移所有业务主线接口
- phase 收口后，不允许仓库仍保留 hand-written JSON 业务主线作为正式默认模式
- 若迁移过程中需要短时并存 JSON adapter，只允许作为临时排障手段，不得写入阶段完成态

### 4.7 当前阶段源码设计层输出要求

`phase07` 虽然当前处于 `/plan`，但为了保证后续 `/spec` 可直接进入实现，本阶段必须把以下源码设计层结果纳入任务规划：

- ConnectRPC handler 挂载与 `chi` route shell 组合方式
- `buf.gen.yaml` 升级后的生成链与产物落点
- Go service implementation 与 generated interface 的对接方式
- 前端生成客户端承接位与迁移边界
- `phase01 ~ phase06` canonical 业务接口的 `service / RPC / 当前入口路径 / 页面动作 owner` 迁移清单
- Vite dev proxy、Caddy、本地启动脚本与验收脚本对 `/api` 单一访问前缀的承接设计
- `chi` middleware、Connect interceptor 与错误映射的组合设计
- 回归矩阵、工具链迁移矩阵与最终验收证据清单

## 5. 当前阶段完成条件

`phase07` 完成时，至少必须满足：

1. 新的 `buf` 生成链已稳定生成 Go Connect 产物与前端客户端产物
2. `phase01 ~ phase06` canonical 业务接口已完成正式传输主线切换
3. `chi` 只保留装配层与非业务端点承载职责
4. 旧手写 JSON 业务主线已退场
5. 覆盖 `phase01 ~ phase06` 的联调回归验收已通过
6. 根级入口已回写“`phase07` 为 `mvp0.3` 业务阶段前置基础 phase”的状态结论
