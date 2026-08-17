# Tasks

- [x] Task 1: 冻结项目治理画像后端实现的唯一正式主线
  - [x] SubTask 1.1: 明确 `repository_id` 驱动的 repository-scoped governance profile 是唯一后端实现主线
  - [x] SubTask 1.2: 明确四实体既有 service / repository 主线保持不变
  - [x] SubTask 1.3: 复核不会长出第二条治理画像后端实现主线

- [x] Task 2: 冻结第一版 `.proto` 与 ConnectRPC 合同落点
  - [x] SubTask 2.1: 明确继续使用现有 `proto/` workspace 与 `.proto -> ConnectRPC` 主线
  - [x] SubTask 2.2: 明确治理画像读写 RPC 以 `repository_id` 为唯一正式结构化输入锚点
  - [x] SubTask 2.3: 复核不会新增并列 buf workspace 或临时 JSON canonical API

- [x] Task 3: 冻结 Go 分层与模块落点
  - [x] SubTask 3.1: 明确 `connect / service / repository / domain types` 分层职责
  - [x] SubTask 3.2: 明确 router 继续沿用现有平台装配模式接入
  - [x] SubTask 3.3: 复核 candidate 只读层不会被误当作治理画像正式写入层

- [x] Task 4: 冻结 PostgreSQL 持久化结构与数据库迁移要求
  - [x] SubTask 4.1: 明确治理画像主记录与两组 bindings 的正式表结构
  - [x] SubTask 4.2: 明确 `repository_id` 唯一锚点与关联约束
  - [x] SubTask 4.3: 复核 markdown 正文与顶层目录矩阵不会被误落入可写持久化字段

- [x] Task 5: 冻结第一版写路径主线
  - [x] SubTask 5.1: 明确治理画像存在单一正式保存入口
  - [x] SubTask 5.2: 明确 `track_type / current_phase_*` 排除在可写集合之外
  - [x] SubTask 5.3: 明确主记录与 bindings 在同一事务边界内保存
  - [x] SubTask 5.4: 复核不会偷渡目录扫描、模板同步或正文入库

- [x] Task 6: 冻结第一版读取主线
  - [x] SubTask 6.1: 明确治理画像存在单一正式结构化读取入口
  - [x] SubTask 6.2: 明确读取结果区分结构化字段、结构化摘要与正文回源能力
  - [x] SubTask 6.3: 明确 `phase13-08` 只承接治理画像读取，不直接扩写成完整 brief 主线

- [x] Task 7: 冻结治理画像后端主线与现有 `ProjectContextService` 的演进边界
  - [x] SubTask 7.1: 明确 `ProjectContextService` 继续承接 phase12 既有最小只读聚合职责
  - [x] SubTask 7.2: 明确治理画像后端主线承接结构化写读与资产承接结果持久化
  - [x] SubTask 7.3: 复核不会长出第二套项目事实源或第二套并列项目上下文读取协议

- [x] Task 8: 冻结第一版非目标与禁止事项
  - [x] SubTask 8.1: 明确自动扫描、自动同步、全文入库不进入 `phase13-08`
  - [x] SubTask 8.2: 明确不重写四实体既有后端主线
  - [x] SubTask 8.3: 明确不在 `phase13-08` 直接实现完整 agent brief 主线

- [x] Task 9: 完成 spec 包与上游冻结文档的一致性校验
  - [x] SubTask 9.1: 校验本 spec 包与 `phase13-05` 的后端边界保持单值一致
  - [x] SubTask 9.2: 校验本 spec 包与 `phase13-07` 的 brief 演进边界保持单值一致
  - [x] SubTask 9.3: 校验本 spec 包与现有 `project_context.proto` 的职责边界保持单值一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 2`, `Task 5`, and `Task 6`
- `Task 8` depends on `Task 5`, `Task 6`, and `Task 7`
- `Task 9` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, `Task 7`, and `Task 8`

# 执行记录

- SubTask 1.1 结论：项目治理画像第一版后端实现主线已冻结为同一 `repository_id` 驱动的 repository-scoped governance profile 持久化与读取主线，和 `phase13-05` 的唯一后端承接位保持单值一致
- SubTask 1.2 结论：四实体既有 `Product / Repository / Module / Decision` 的 service / repository / proto 主线已明确保持不变，治理画像实现不会反向改造四实体主线
- SubTask 1.3 结论：spec 已显式禁止长出第二条治理画像后端实现主线，避免多个后端入口并列维护治理画像写读语义
- SubTask 2.1 结论：治理画像后端合同落点已冻结为继续使用现有 `proto/` workspace 与 `.proto -> ConnectRPC` 主线，不新增并列 buf workspace 或第二套生成链
- SubTask 2.2 结论：治理画像读写 RPC 已明确以 `repository_id` 作为唯一正式结构化输入锚点，与 `phase13-05 / phase13-07 / project_context.proto` 的仓库锚点口径一致
- SubTask 2.3 结论：spec 已显式排除临时 JSON canonical API 与手写 `chi + JSON HTTP` 业务接口，保持 `.proto` 为唯一长期合同源
- SubTask 3.1 结论：Go 代码分层已冻结为 `connect / service / repository / domain types`，治理画像写入约束、字段分类与事务语义统一收敛到 `service/`
- SubTask 3.2 结论：router 装配已明确继续沿用 `backend/internal/platform/router.go` 模式接入，不新增第二套装配机制
- SubTask 3.3 结论：spec 已显式禁止把治理画像正式写入塞进 `projectcontext/candidate` 这类只读 candidate 层，避免职责漂移
- SubTask 4.1 结论：数据库持久化结构已冻结为治理画像主记录 + `canonical_root_file_bindings` + `global_asset_bindings` 三层，直接承接 `phase13-04 / 13-05` 的字段与绑定矩阵
- SubTask 4.2 结论：`repository_id` 已明确成为治理画像主记录唯一业务锚点，并要求在 bindings 层形成明确外键或等价关联约束
- SubTask 4.3 结论：markdown 正文与顶层目录矩阵 `backend / database / frontend / proto` 已明确排除在 repository-scoped 可写持久化字段之外
- SubTask 5.1 结论：治理画像后端写路径已冻结为单一正式保存入口，不允许多个 RPC / handler / service 各自维护一套写语义
- SubTask 5.2 结论：`track_type / current_phase_name / current_phase_ref / current_phase_status` 已被再次显式排除在可写集合之外，和 `phase13-04 / 13-05` 的 `read-only` 口径重新对齐
- SubTask 5.3 结论：治理画像主记录、canonical 根级文件 bindings 与全局规范资产 bindings 已要求在同一事务边界内保存，避免半套状态写入
- SubTask 5.4 结论：自动扫描、模板同步、正文入库与自动建议能力都已显式排除在 `phase13-08` 写路径之外，保持“手工维护优先”
- SubTask 6.1 结论：治理画像后端读取已冻结为单一正式结构化读取入口，供 `phase13-09` 前端回看与维护入口消费
- SubTask 6.2 结论：读取结果已明确区分结构化字段、结构化摘要与 markdown 正文回源能力，不把正文误当 stored field
- SubTask 6.3 结论：spec 已显式限定 `phase13-08` 只承接治理画像读取主线，不在本阶段直接扩写成完整 agent brief 主线
- SubTask 7.1 结论：现有 `ProjectContextService` 已被冻结为继续承接 phase12 最小只读聚合读取与 Markdown 导出职责，不被误改为治理画像正式写读主线
- SubTask 7.2 结论：治理画像后端主线承接结构化写读与资产承接结果持久化，和 `ProjectContextService` 形成职责分层而不是事实并列
- SubTask 7.3 结论：spec 已显式禁止两者长期并列为两套项目正式事实源或两套并列项目上下文读取协议；后续 `phase13-10` 若要把 brief 收口到 `ProjectContextService` 主线，必须按 `phase13-07` 已冻结的受控演进规则处理
- SubTask 8.1 结论：自动扫描、自动同步、存在性校验与全文入库已明确列为 `phase13-08` 非目标
- SubTask 8.2 结论：四实体既有后端主线不重写已显式进入 spec requirement，与 `dev_plan` 范围保持一致
- SubTask 8.3 结论：完整 agent brief 主线明确留给后续 `phase13-10`，`phase13-08` 不偷渡该实现
- SubTask 9.1 结论：本 spec 包与 `phase13-05` 的后端承接位、存储分层、读写边界和全文禁止入库口径保持单值一致
- SubTask 9.2 结论：本 spec 包与 `phase13-07` 的 brief 演进边界保持单值一致：治理画像读取是 brief 的上游结构化数据源，但 `phase13-08` 不直接承担完整 brief 组装
- SubTask 9.3 结论：本 spec 包与现有 `proto/psco/project_context/v1/project_context.proto` 的职责边界保持单值一致：`ProjectContextService` 继续承接 phase12 最小只读聚合，治理画像后端主线承接 phase13 结构化写读

# 实现执行记录（phase13-08 后端主线落地）

实现交付物（`/spec` 冻结后按本 spec 执行）：

1. `.proto` 合同源：`proto/psco/governance_profile/v1/governance_profile.proto`
   - `GovernanceProfileService`（`GetGovernanceProfile / UpdateGovernanceProfile`），`repository_id` 为唯一结构化输入锚点
   - `UpdateGovernanceProfileRequest` 不携带 `track_type / current_phase_*` 与 `project_profile_version`，read-only 排除在合同层即成立
   - `optional string`（proto3 field presence）承接 `template_source / structured_summary` 可空语义
   - 生成产物经 `make lint && make build && make gen` 落地（Go / Go Connect / TS 三链）
2. 数据库迁移：`database/migrations/0010_phase13_governance_profile.sql`
   - `governance_profiles` 主记录：`repository_id UUID NOT NULL UNIQUE REFERENCES repositories(id)` 唯一业务锚点，承接 9 类核心字段列
   - `governance_canonical_root_file_bindings`：`governance_profile_id` 外键 + `UNIQUE(governance_profile_id, file_name)`
   - `governance_global_asset_bindings`：`governance_profile_id` 外键 + `UNIQUE(governance_profile_id, name)`，`structured_summary` 为可空列（必填性由 service 按矩阵校验）
   - markdown 正文无任何 canonical 存储列；顶层目录矩阵不进入迁移
   - 已由后端 `RunMigrations` 正式应用并记录版本（`0010_phase13_governance_profile`）
3. Go 模块 `backend/internal/governanceprofile/`（`connect / service / repository / candidate / types / errors` 分层）
   - `types.go`：领域结构 + 根级上游冻结结论投影（read-only 字段唯一受控来源）+ 8 项资产冻结矩阵
   - `errors.go`：`ErrRepositoryNotFound / ErrGovernanceProfileNotFound / ErrInvalidInput / ErrGovernanceProfileReadFailed / ErrGovernanceProfileSaveFailed`
   - `candidate/repository_reader.go`：仓库存在性前提校验（跨模块读取经 candidate 子包隔离）
   - `repository/profile_store.go`：`SaveProfile` 单一事务边界（UPSERT 主记录 + bindings 全量替换，任一步失败整体回滚）；UPSERT 的 `ON CONFLICT DO UPDATE` 分支只 SET 可写字段，不改写 read-only 列；`ReadProfile` 聚合读取
   - `service/query_service.go`：读取编排（repo 不存在 → NotFound；画像未创建 → NotFound）
   - `service/command_service.go`：保存编排 + 字段分类校验（required 非空、file_name/name 去重、name 限定 8 项矩阵、前 5 项 `structured_summary` 必填）
   - `connect/server.go`：Connect handler（解包 → service → 组装 → `connecterrors.MapToConnectError`）
4. 平台装配：`backend/internal/platform/router.go`（`buildGovernanceProfile` + `mountGovernanceProfileConnect`）、`backend/internal/platform/server.go`（phase13 装配段）、`backend/internal/connecterrors/connect_errors.go`（错误映射注册：NotFound × 2、InvalidArgument × 1、Internal × 2）

验证记录：

- `go build ./... && go vet ./...` 通过；`go test ./...` 全量通过（无回归）
- 前端 `tsc --noEmit` 通过（TS 生成产物为纯新增）
- Connect 冒烟（本地 dev 库，repo `main-repo`）：
  1. repo 不存在 → `not_found: repository not found in PSCO`
  2. repo 存在但画像未创建 → `not_found: governance profile not created`
  3. 创建画像 → read-only 字段由受控投影初始化（`TRACK_TYPE_DURABLE_SYSTEM` / `phase13_project_governance_profile_foundation` / `PHASE_STATUS_IN_PROGRESS`），`project_profile_version` 服务端固定 `project_governance_profile_v1`
  4. 资产名越出 8 项矩阵 → `invalid_argument`
  5. 前 5 项缺 `structured_summary` → `invalid_argument`
  6. 二次保存（`templateSource: null` + bindings 精简）→ read-only 字段保持不变、`createdAt` 不变、`updatedAt` 更新、bindings 全量替换（9→2、8→2）
  7. 冒烟后画像已恢复为完整 9 root files + 8 asset bindings，供 phase13-09 前端承接开发使用

边界遵守确认：

- 未引入自动扫描、自动同步、存在性校验或全文入库
- 未重写四实体既有 service / repository / proto 主线
- 未实现完整 agent brief 主线（留给 phase13-10）
- 未形成第二套项目事实源或第二套并列读取协议

# 独立复核记录

- 复核结论：**通过，无阻断性问题**
- spec 符合性：8 项 ADDED Requirement 逐项确认一致（唯一主线 / .proto+ConnectRPC 落点 / Go 分层 / PostgreSQL 结构 / 写路径 / 读取主线 / ProjectContextService 演进边界 / 非目标）
- 重点核验：SQL 列与 Scan 目标数量顺序全部匹配；UPSERT `ON CONFLICT DO UPDATE` 确实不改写 read-only 列；8 项资产矩阵与 phase13-05 逐项一致；错误映射与枚举转换无遗漏分支；`go build / vet` 复核通过
- 已修复（复核建议 3）：清理 `candidate/repository_reader.go` 中 `SELECT EXISTS` 场景下不可达的 `pgx.ErrNoRows` 死代码分支，修复后重新构建验证通过
- 非阻断建议留档（供后续 phase 显式进入，不在 phase13-08 处理）：
  1. read-only 字段快照的受控同步机制未落地：根级上游演进时已存在行的 read-only 列需后续明确刷新方式
  2. `SaveProfile` 提交后回读失败时错误语义可进一步区分"已提交但回读失败"
  3. "markdown 正文可回源能力状态"第一版未显式返回：phase13-09/10 若引入回源需补能力状态与独立失败语义
  4. `docs_workflow_layout` 仅非空校验，与"受控 string"措辞有落差：后续明确是否需要值级约束
  5. 后 3 项资产允许携带非空 summary（宽松读法）：后续口径统一时确认
