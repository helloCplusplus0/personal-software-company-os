# phase09-08 落实支撑能力相关合同、后端承接与前端 read owner Spec

## Why

`phase09-04 ~ 07` 已经把模板复用与派生提示的合同、后端服务、前端 read owner 和返回链边界冻结成单值口径，但仓库里目前还不存在真正可执行的 `template_reuse.proto`、`backend/internal/templatereuse/` 与 `frontend/src/features/template-reuse/data/*`。因此，`phase09-08` 的职责不再是补一轮设计冻结，而是把这些已拍板的承接位真实落到代码树中，并用最小工具链与 smoke 证明它们已经可被后续 `phase09-09 / 10` 消费。

## What Changes

- 实现 `psco.template_reuse.v1` proto 合同并接入现有 `buf` 生成主线
- 实现 `backend/internal/templatereuse/` 后端模块、Connect handler、最小读时派生数据承接与 `/api` 挂载
- 实现 `frontend/src/features/template-reuse/data/` 前端只读切片，包括 slice-local client、query options 与四个正式 read owner
- 明确本阶段源码实现边界：只落合同、后端 query owner、前端 read owner 与最小 smoke，不提前实现模板 handoff、页面展示和 create 回流
- 明确本轮交付必须包含的验证：`buf`、`go build`、`frontend build` 与 template reuse API/read owner smoke

## Impact

- Affected specs:
  - `phase09_04_freeze_contract_read_model_owner_candidate_source_boundary`
  - `phase09_05_design_template_reuse_hint_page_flow_interaction_return_chain`
  - `phase09_06_design_frontend_read_write_owner_state_flow`
  - `phase09_07_design_backend_service_contract_minimal_data_handoff`
  - `phase08_08_land_review_contract_backend_frontend_owner_enablement`
  - `phase06_13_land_minimal_proto_contract_mainline`
- Affected code:
  - `proto/psco/template_reuse/v1/template_reuse.proto`
  - `proto/Makefile`
  - `backend/internal/templatereuse/`
  - `backend/internal/platform/router.go`
  - `frontend/src/gen/proto/psco/template_reuse/v1/*`
  - `frontend/src/features/template-reuse/data/connect-client.ts`
  - `frontend/src/features/template-reuse/data/template-reuse-query-options.ts`
  - `frontend/src/features/template-reuse/data/use-template-candidates-read.ts`
  - `frontend/src/features/template-reuse/data/use-template-prefill-read.ts`
  - `frontend/src/features/template-reuse/data/use-template-source-read.ts`
  - `frontend/src/features/template-reuse/data/use-derived-insight-hints-read.ts`
  - `frontend/src/features/review/data/use-weekly-review-read.ts` 或后续等价组合位
  - `frontend/src/features/product-registry/application/use-product-create-template-handoff.ts` 的未来消费位
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx` 的未来消费位

## ADDED Requirements

### Requirement: template reuse 合同必须沿用现有单一 buf 主线正式落地

系统 SHALL 在现有 `proto/` workspace 内新增 `proto/psco/template_reuse/v1/template_reuse.proto`，并通过既有 `buf build / gen / lint / breaking` 主线同时生成后端 Go proto、Go Connect 与前端 TypeScript 合同产物。

#### Scenario: template_reuse proto 的物理落点与生成产物

- **WHEN** 实现 `phase09-08`
- **THEN** 仓库中必须存在 `proto/psco/template_reuse/v1/template_reuse.proto`
- **AND** 生成产物必须继续只落到：
  - `backend/internal/gen/proto/psco/template_reuse/v1/*`
  - `backend/internal/gen/connect/psco/template_reuse/v1/*`
  - `frontend/src/gen/proto/psco/template_reuse/v1/*`
- **AND** 不得为 template reuse 单独新增第二份 `buf.yaml`、第二份 `buf.gen.yaml`、第二套 TS 生成脚本或第二套前端 RPC 客户端生成链

#### Scenario: template_reuse proto 只承接 phase09-07 已冻结的最小 RPC

- **WHEN** 编写 `template_reuse.proto`
- **THEN** `TemplateReuseService` 只允许实现：
  - `ListTemplateCandidates`
  - `GetTemplateCandidatePrefill`
  - `GetDerivedInsightHints`
  - `GetTemplateSourceSummary`
- **AND** 当前阶段不得新增 `CreateProductFromTemplate`、`ApplyTemplateToProduct`、`PersistActiveTemplateCandidate` 或等价模板写 RPC
- **AND** request / response 字段、枚举与 `UNAVAILABLE` 成功态必须继续对齐 `phase09-07` 已冻结合同

### Requirement: template reuse 后端正式承接位必须落到单一 `backend/internal/templatereuse/`

系统 SHALL 将 template reuse 后端正式实现收敛到单一 `backend/internal/templatereuse/` 模块，并沿用现有 phase07 Connect transport 主线，不得把模板读逻辑散落到 `review / productregistry / reusesummary / platform`。

#### Scenario: template reuse 后端最小模块结构

- **WHEN** 实现 template reuse 后端
- **THEN** 至少必须新增以下落点：
  - `backend/internal/templatereuse/service/query_service.go`
  - `backend/internal/templatereuse/candidate/template_candidate_readers.go`
  - `backend/internal/templatereuse/connect/server.go`
- **AND** 若需要 domain types / errors，只允许继续收敛在 `backend/internal/templatereuse/` 模块内部
- **AND** 当前阶段不得新增 `backend/internal/templatereuse/service/command_service.go`
- **AND** 不得把模板候选 derivation、`templateCandidateId` 解析或提示派生直接塞进 `backend/internal/platform/router.go`

#### Scenario: templatereuse QueryService 只承接 phase09 已冻结的读时派生

- **WHEN** 实现 `templatereuse.QueryService`
- **THEN** 它只允许承接：
  - 基于 `product_modules` 已持久化事实派生模板候选
  - 基于 `templateCandidateId` 解析预填详情与模板来源复读
  - 基于 `templateCandidateId + review_scope_key` 计算 `Weekly Review` 场景下的提示
  - 返回 `RESOLVED / UNAVAILABLE` 两种解析状态
- **AND** 不得绕过既有 `ReuseSummaryService` 复制第二套 `capability_summary`
- **AND** 不得把 `ReviewService` 改写成模板候选的 canonical query owner

#### Scenario: candidate reader 的数据承接必须保持最小且单值

- **WHEN** 实现 `template_candidate_readers.go`
- **THEN** 读时派生只能直接依赖：
  - `products`
  - `modules`
  - `product_modules`
- **AND** 可以在 reader 内做去重键、排序与 `templateCandidateId` 派生
- **AND** 当前阶段不得新增模板快照表、候选缓存表、预填持久化记录或服务器侧会话表
- **AND** service 层不得直接写跨模块 SQL

### Requirement: template reuse Connect transport 必须正式挂到现有 `/api` 业务树

系统 SHALL 将 template reuse 的对外传输正式接入现有 `/api` Connect 主线，保持与其他业务模块一致的 handler 构造与路由挂载方式。

#### Scenario: template reuse Connect handler 的挂载方式

- **WHEN** 落实 template reuse Connect transport
- **THEN** `backend/internal/platform/router.go` 必须新增 `mountTemplateReuseConnect(...)` 或等价单值装配位
- **AND** handler 必须继续通过 generated `(path, handler)` 挂到单一 `/api` 业务树
- **AND** template reuse 新增错误必须继续统一收敛到既有 Connect 错误映射入口
- **AND** 对于 `templateCandidateId` 漂移一类可恢复场景，必须继续返回成功 envelope 中的 `UNAVAILABLE`，而不是 transport error
- **AND** 不得新增 `/template-api`、`/rpc`、`/internal-template-reuse` 等第二套路由根

### Requirement: template reuse 前端正式承接位必须以单一 `frontend/src/features/template-reuse/data/` 落地

系统 SHALL 将 `phase09` 的模板候选、预填、来源复读与提示读取正式实现为单一 `template-reuse` 只读切片，而不是让 `Weekly Review / Product Create / Product Detail` 页面继续直接拼 generated client 或底层 `useQuery`。

#### Scenario: template reuse 前端最小切片结构

- **WHEN** 实现 template reuse 前端只读能力
- **THEN** 至少必须新增以下落点：
  - `frontend/src/features/template-reuse/data/connect-client.ts`
  - `frontend/src/features/template-reuse/data/template-reuse-query-options.ts`
  - `frontend/src/features/template-reuse/data/use-template-candidates-read.ts`
  - `frontend/src/features/template-reuse/data/use-template-prefill-read.ts`
  - `frontend/src/features/template-reuse/data/use-template-source-read.ts`
  - `frontend/src/features/template-reuse/data/use-derived-insight-hints-read.ts`
- **AND** 当前阶段不得在 `template-reuse/data/` 混入 handoff 编排、导航、toast、query invalidation 或 mutation
- **AND** 当前阶段不得新增 `frontend/src/features/template-reuse/application/` 作为页面级编排兜底层

#### Scenario: template reuse slice-local connect client 的正式 transport 路径

- **WHEN** 实现 `template-reuse/data/connect-client.ts`
- **THEN** 它必须以 generated `TemplateReuseService` client 为唯一 transport 入口
- **AND** 继续复用 `@/shared/rpc/connect-transport`
- **AND** 不得在 `WeeklyReviewPage / ProductCreatePage / ProductDetailPage` 内直接 `createClient(TemplateReuseService, ...)`

#### Scenario: template reuse query options 与 read owner 的正式职责

- **WHEN** 实现四个 template reuse read owner
- **THEN** 它们必须只承接：
  - queryKey
  - 只读请求
  - 响应解包
  - `empty / unavailable / error` 的只读派生
- **AND** 至少形成以下一对一映射：
  - `use-template-candidates-read` -> `ListTemplateCandidates`
  - `use-template-prefill-read` -> `GetTemplateCandidatePrefill`
  - `use-derived-insight-hints-read` -> `GetDerivedInsightHints`
  - `use-template-source-read` -> `GetTemplateSourceSummary`
- **AND** 当前阶段不得在 read owner 内提前拼 `Product Create` handoff、create 成功导航或页面级 active candidate 状态机

### Requirement: phase09-08 只落 enablement，不提前实现页面消费主链

系统 SHALL 将 `phase09-08` 的实现范围落实为“合同 + 后端 query owner + 前端 read owner + 最小 smoke”，而不是提前把 `phase09-09 / 10` 的页面与交互逻辑偷跑到这一轮。

#### Scenario: 不提前实现模板 handoff 与 create 回流

- **WHEN** 实现 `phase09-08`
- **THEN** 当前阶段只允许交付模板预填读取的正式 read owner 与后端接口
- **AND** 当前阶段不得提前实现：
  - `Product Create` 的模板 handoff application owner
  - 创建成功后的模板来源导航回流
  - `Product Detail` 的模板来源摘要展示
- **AND** 这些能力必须继续留给 `phase09-09`

#### Scenario: 不提前实现提示展示主链

- **WHEN** 实现 `phase09-08`
- **THEN** 当前阶段只允许交付 `GetDerivedInsightHints` 与 `use-derived-insight-hints-read`
- **AND** 当前阶段不得提前把 `reuse_opportunity_hint / capability_gap_hint` 渲染进 `Weekly Review` 或 `Product Create` 页面
- **AND** 提示展示、CTA handoff 与解释性回流必须继续留给 `phase09-10`

### Requirement: 本阶段不得把页面级临时拼装当作长期稳态

系统 SHALL 要求 `phase09-08` 的 enablement 即使暂时尚未被页面消费，也必须以正式切片 owner 形式真实落地，而不是把“先在页面里临时拼一下”当成过渡稳态。

#### Scenario: 禁止页面直接接 generated client 或底层 query

- **WHEN** 后续为了 smoke 或临时调试接入 template reuse 能力
- **THEN** `WeeklyReviewPage`、`ProductCreatePage`、`ProductDetailPage` 不得直接：
  - import `TemplateReuseService`
  - `createClient(...)`
  - 直接 `useQuery(...)` 调用 template reuse RPC
- **AND** 正式长期消费位必须继续只认 `template-reuse/data/*` 导出的 read owner

#### Scenario: review / product-registry 不得成为 template reuse 的过渡宿主

- **WHEN** 实现者需要给后续 `phase09-09 / 10` 预留消费口
- **THEN** 可以在 `use-weekly-review-read`、`use-product-create-template-handoff` 或 `ProductDetailPage` 的后续消费位注明将使用 template reuse read owner
- **AND** 不得把模板 read owner 直接落进 `review/data/` 或 `product-registry/data/`
- **AND** 不得把当前 enablement 退化为 `page-local helper + TODO 后续再抽离`

### Requirement: 验收口径必须同时覆盖合同、后端、前端与最小 API smoke

系统 SHALL 将 `phase09-08` 的实现验收落实为“合同工具链 + 后端构建 + 前端构建 + template reuse API/read owner smoke”四层。

#### Scenario: 最小验证命令

- **WHEN** 验证 `phase09-08` 是否达成 DoD
- **THEN** 至少必须通过：
  - `(cd proto && make build && make gen && make lint)`
  - `(cd proto && make breaking)` 作为有 `main` 基准时的 breaking gate
  - `(cd backend && go build ./...)`
  - `(cd frontend && npm run build)`
- **AND** 不得只验证后端构建而跳过前端 read owner 收敛的类型构建

#### Scenario: template reuse 最小 smoke

- **WHEN** 执行 `phase09-08` 的最小 smoke
- **THEN** 至少必须覆盖：
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/ListTemplateCandidates`
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/GetTemplateCandidatePrefill`
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/GetDerivedInsightHints`
  - `POST /api/psco.template_reuse.v1.TemplateReuseService/GetTemplateSourceSummary`
- **AND** smoke 中必须至少显式覆盖一种 `UNAVAILABLE` 成功态
- **AND** 前端侧必须至少验证四个 read owner 均能成功编译、导出稳定返回类型，并形成后续页面消费的唯一正式入口

## MODIFIED Requirements

### Requirement: `phase09-07` 中 template reuse enablement 的解释口径

`phase09-07` 已冻结 `template_reuse.proto`、`TemplateReuseService`、`templatereuse.QueryService` 与前端 read owner 的逻辑边界。

自 `phase09-08` 起，系统必须把该 requirement 进一步修改为：

- 上述合同与 owner 不再只是设计结论，而必须成为仓库中真实存在的 proto 文件、后端模块、Connect handler 与前端只读切片
- 当前阶段 enablement 必须做到“后续页面可以直接消费正式 owner”，而不是继续依赖页面级临时组合
- 但 `phase09-08` 仍不提前承担 handoff、页面展示与 create 回流逻辑

#### Scenario: 判断 phase09-08 是否仍保持 enablement 边界

- **WHEN** 审查 `phase09-08` 实现结果
- **THEN** 必须同时满足：
  - 合同与 owner 已真实落地
  - 页面级临时拼装没有被当成长期稳态
  - `phase09-09 / 10` 的页面交互逻辑没有被提前偷跑

## REMOVED Requirements

### Requirement: 允许以页面级临时拼 query / connect client / helper 作为 template reuse 正式 enablement

**Reason**: 这会直接破坏 `phase09-06` 已冻结的前端 owner 收敛原则，并让 `phase09-08` 交付物退化成“先能跑，后面再抽离”的散装过渡态。

**Migration**: template reuse 正式 enablement 统一回收到 `template_reuse.proto + TemplateReuseService + templatereuse.QueryService + frontend/src/features/template-reuse/data/*`；`Weekly Review / Product Create / Product Detail` 仅在后续 phase 中消费这些正式 owner，不再并排发明页面级临时接线。 
