# Tasks

- [x] Task 1: 实现 `decisioncenter` 后端模块支撑文件（types.go / errors.go / validate.go / handler/response.go）
  - [x] SubTask 1.1: 实现 `types.go`，定义 `DecisionStatus` / `DecisionLinkTargetType` 枚举、`Decision` / `DecisionListItem` / `LinkedModule` / `SourceContext` / `DecisionDetail` / `DecisionModuleCandidate` 结构体、`CreateDecisionRequest` / `CreateDecisionResponse` / `LinkDecisionToTargetRequest` 请求体、`ListQuery` 查询参数，从 `.proto` 单向承接语义
  - [x] SubTask 1.2: 实现 `errors.go`，定义 `ErrDecisionNotFound` / `ErrModuleNotFound` / `ErrDuplicateLink` / `ErrInvalidInput` / `ErrInvalidStatus` / `ErrInvalidTargetType` / `ErrInvalidAlternatives` 哨兵错误
  - [x] SubTask 1.3: 实现 `validate.go`，定义 `ValidateDecisionID` / `ValidateModuleID` UUID 校验辅助
  - [x] SubTask 1.4: 实现 `handler/response.go`，定义 `writeJSON` / `writeError` / `decodeJSON`，与 `moduleregistry/handler/response.go` 同构，覆盖 `Decision Center` 错误到 HTTP 状态码映射（404 / 409 / 400 / 500）

- [x] Task 2: 实现 `repository` 数据访问层（decision_store.go / link_store.go）
  - [x] SubTask 2.1: 实现 `repository/decision_store.go`，承接 `decisions` 表的 `Create` / `GetByID` / `List`（含 `link_count` 与 `linked_module_summary` 计算的 SQL 查询）/ `Exists`
  - [x] SubTask 2.2: 实现 `repository/link_store.go`，承接 `decision_links` 表的 `Create` / `ListByDecisionID` / `ExistsByDecisionAndModule` / `ListLinkedModuleIDs`
  - [x] SubTask 2.3: 确认 `decision_store.go` 的 `List` 查询实现 `link_count`（COUNT）与 `linked_module_summary`（按 `module_name` 升序取前 3 + `+N`）计算口径，无关联时返回 `0` 与 `''`

- [x] Task 3: 实现 `candidate` 跨模块候选读取（module_candidate_read.go）
  - [x] SubTask 3.1: 实现 `candidate/module_candidate_read.go`，定义 `ModuleCandidateRead` struct 与 `NewModuleCandidateRead` 构造函数，并新增 `ModuleExists` 跨模块只读校验方法（承接 `LinkDecisionToTarget` 的 Module 存在性校验，service 层不直接写跨模块 SQL）
  - [x] SubTask 3.2: 实现 `List(ctx, decisionID)` 方法，查询 `modules` 表，排除已关联目标，排序 `status(active 优先) -> module_name 升序`，`status` 复用 `moduleregistry.ModuleStatus`
  - [x] SubTask 3.3: 确认空候选返回空列表 `[]`，不返回 `null`

- [x] Task 4: 实现 `service` 业务编排层（query_service.go / command_service.go）
  - [x] SubTask 4.1: 实现 `service/query_service.go`，承接 `ListDecisions` / `GetDecisionDetail` / `ListModuleCandidates`，`GetDecisionDetail` 返回 `DecisionDetail`（含 `linked_modules` 与 `source_context`）
  - [x] SubTask 4.2: 实现 `service/command_service.go` 的 `CreateDecision`，执行必填校验 → `status` 合法性校验 → `alternatives` 条目校验 → `source_module_id` 可选来源校验 → 写入 → 返回 `decision_id`
  - [x] SubTask 4.3: 实现 `service/command_service.go` 的 `LinkDecisionToTarget`，执行 `target_type` 校验 → `Decision` 存在性 → `Module` 存在性（通过 candidate.ModuleExists 跨模块只读校验） → 重复关联检测 → 写入 → 返回空响应
  - [x] SubTask 4.4: 确认 `source_context` 的组装逻辑（从 `decisions.source_module_id` `LEFT JOIN modules` 读取来源 `Module` 标识，支持 `phase03-10 §5.11`"持续到完成或放弃"的跨刷新承接；无来源时 `source_module_id` / `source_module_name` 返回空字符串）

- [x] Task 5: 实现 `handler` HTTP 入口层（query_handler.go / command_handler.go）
  - [x] SubTask 5.1: 实现 `handler/query_handler.go`，承接 `ListDecisions`（GET `/api/decisions`）/ `GetDecisionDetail`（GET `/api/decisions/{decisionId}`）/ `ListDecisionModuleCandidates`（GET `/api/decisions/{decisionId}/candidates/modules`）
  - [x] SubTask 5.2: 实现 `handler/command_handler.go`，承接 `CreateDecision`（POST `/api/decisions`，返回 201 + `decision_id`）/ `LinkDecisionToTarget`（POST `/api/decisions/{decisionId}/links`，返回 204）
  - [x] SubTask 5.3: 确认 URL 路径参数（`decisionId`）在 handler 层组装为 service 调用参数，不放在 JSON 请求体

- [x] Task 6: 实现 `decisions` 表原位升级 migration（0004_decision_center_mainline.sql）
  - [x] SubTask 6.1: 创建 `database/migrations/0004_decision_center_mainline.sql`，通过 `ALTER TABLE decisions ADD COLUMN` 添加 `context / problem / alternatives TEXT[] / choice / reason / impact / status` 字段
  - [x] SubTask 6.2: 对无 `DEFAULT` 的必填字段（`context / problem / choice / reason`）按"先 `ADD COLUMN` 允许 NULL → 回填占位文本 → `SET NOT NULL`"三步流程添加
  - [x] SubTask 6.3: 对有 `DEFAULT` 的字段（`alternatives TEXT[] NOT NULL DEFAULT '{}'` / `impact TEXT NOT NULL DEFAULT ''` / `status TEXT NOT NULL DEFAULT 'proposed'`）直接 `ADD COLUMN ... NOT NULL DEFAULT ...`
  - [x] SubTask 6.4: 添加 `status` 的 `CHECK` 约束与 `idx_decisions_status_created_at` 索引
  - [x] SubTask 6.5: 实现现有示例数据兼容回填（`WHERE <新字段> IS NULL`），保留原有 `title / created_at`

- [x] Task 7: 实现基线 seed 与 seed_readonly_prereqs 更新
  - [x] SubTask 7.1: 创建 `database/seeds/seed_decision_mainline_baseline.sql`，开头 `DELETE FROM decisions`，`decision_links` 与 `decisions` 的 `INSERT` 使用 `ON CONFLICT DO NOTHING` / `WHERE NOT EXISTS` 守卫，`BEGIN / COMMIT` 事务包裹
  - [x] SubTask 7.2: 插入至少 `3` 条结构化 `Decision`（`1 proposed + 1 active + 1 archived/superseded`），`1` 条保留 `phase02` 原有 `title`，覆盖 `alternatives` 数组/空数组与 `impact` 空字符串维度
  - [x] SubTask 7.3: 插入至少 `2` 条 `decision_links`（通过 `module_name` 与 `decision title` 查找 `ID`），`1` 条复用 `phase02` 关联，至少 `1` 条 `Decision` 无 `decision_links`
  - [x] SubTask 7.4: 更新 `database/seeds/seed_readonly_prereqs.sql` 中 `decisions` seed 从 `title-only` 升级为结构化字段插入

- [x] Task 8: 实现重置脚本（reset_decision_mainline.sh）
  - [x] SubTask 8.1: 创建 `database/scripts/reset_decision_mainline.sh`，与 `reset_module_mainline.sh` 同构，支持 `--clean-only` / `--restore-only` / 默认三种模式
  - [x] SubTask 8.2: 实现 `resolve_psql` 自动检测（宿主机 psql → docker exec → podman exec）与环境变量覆盖
  - [x] SubTask 8.3: 实现前置校验（PSCO 数据库存在 + `modules` 基线数据存在）与清空范围（`DELETE FROM decisions` 级联清空 `decision_links`）

- [x] Task 9: 实现应用装配（router.go + server.go）
  - [x] SubTask 9.1: 在 `backend/internal/platform/router.go` 中新增 `mountDecisionCenter(r chi.Router, pool *pgxpool.Pool)` 函数，装配 repository → candidate → service → handler → 路由注册
  - [x] SubTask 9.2: 在 `backend/internal/platform/server.go` 的 `r.Route("/api", ...)` 中调用 `mountDecisionCenter(r, pool)`
  - [x] SubTask 9.3: 确认路由注册对齐 `phase03-10` §7.7 RPC → HTTP 映射矩阵

- [x] Task 10: 验证后端可编译、migration 可执行、重置脚本可运行
  - [x] SubTask 10.1: 执行 `go build ./...` 确认后端可编译（通过，无错误）
  - [x] SubTask 10.2: 执行 migration 确认 `0004` 可在已有 `0001-0003` 基础上运行（通过，schema_migrations 记录 0004，decisions 表字段升级完成）
  - [x] SubTask 10.3: 执行 `reset_decision_mainline.sh` 确认清空 + 恢复可重复执行（通过，三种模式均幂等：clean-only / default / restore-only×2 结果一致）
  - [x] SubTask 10.4: 通过 `curl` 验证 5 个 API 端点可访问（列表 / 详情 / 创建 / 关联 / 候选读取），并验证 7 类错误路径返回正确 HTTP 状态码（400 / 404 / 409 / 204 / 201）

- [x] Task 11: 完成 `phase03-12` 规格校验与收口
  - [x] SubTask 11.1: 验证后端模块文件落点与 `phase03-10` §10.2 完全一致
  - [x] SubTask 11.2: 验证错误语义与 `phase03-04` 冻结结论完全一致
  - [x] SubTask 11.3: 验证 migration / seed / 重置脚本与 `phase03-09` 冻结结论完全一致
  - [x] SubTask 11.4: 验证 DTO 从 `.proto` 单向承接，不形成第二套合同源

- [x] Task 12: 修复 GPT-5.4 复核发现的阻断性问题
  - [x] SubTask 12.1: [Issue 1] 修复 `seed_decision_mainline_baseline.sql` 中 `--restore-only` 模式不能稳定恢复到正式基线的问题（对复用 phase02 标题的 Decision 1 增加 UPDATE 收口已有 placeholder，再 INSERT 补缺）
  - [x] SubTask 12.2: [Issue 2] 补齐后端 `source_context` 承接：新增 `0005_decision_source_context.sql` migration（decisions 表添加 source_module_id 列）、.proto CreateDecisionRequest 添加 source_module_id 字段（编号 9）、types.go / repository / service 层支持来源上下文持久化与读取
  - [x] SubTask 12.3: [Issue 2] 验证 source_context API（有值/无值/创建带 source/创建带无效 source 404）
  - [x] SubTask 12.4: [Issue 2] 更新基线 seed 为 Decision 1 添加 source_module_id（关联 auth-service 模块）

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 1` and `Task 4`
- `Task 6` depends on `Task 1`（types 约束）
- `Task 7` depends on `Task 6`
- `Task 8` depends on `Task 7`
- `Task 9` depends on `Task 5`
- `Task 10` depends on `Task 5`, `Task 6`, `Task 8`, and `Task 9`
- `Task 11` depends on `Task 5`, `Task 6`, `Task 7`, `Task 8`, `Task 9`, and `Task 10`
- `Task 12` depends on `Task 11`（GPT-5.4 复核后修复）

# 执行证据

## 编译与静态检查
- `go build ./...` 通过（无输出）
- `go vet ./...` 通过（无输出）

## Migration 验证
- `0004_decision_center_mainline` 成功应用到已有 `0001-0003` 基础上
- `decisions` 表升级后字段：`id / title / created_at / context / problem / choice / reason / alternatives TEXT[] / impact / status`，全部 `NOT NULL`
- `CHECK (status IN ('proposed', 'active', 'superseded', 'archived'))` 约束就位
- `idx_decisions_status_created_at` 索引就位
- `0005_decision_source_context` 成功应用：decisions 表新增 `source_module_id UUID REFERENCES modules(id) ON DELETE SET NULL` 列与索引

## 重置脚本验证
- `--clean-only`：清空后 `decisions=0 / decision_links=0`
- 默认模式（清空+恢复）：`decisions=3 / decision_links=2`
- `--restore-only` ×2（幂等性验证）：两次结果均为 `decisions=3 / decision_links=2`
- [Issue 1 修复验证] `--restore-only` 在已有 readonly placeholder 时，UPDATE 收口 placeholder 为基线内容（context 从"phase02 只读前提占位上下文"变为"团队需要为微服务架构选择认证服务实现方案"），计数与内容均恢复到正式基线

## API 端点验证
- `GET /api/decisions`：返回 3 条 DecisionListItem，`link_count` 与 `linked_module_summary` 计算正确（含空值语义）
- `GET /api/decisions/{id}`：返回 DecisionDetail，含 `linked_modules` 与 `source_context`（Decision 1 有来源 auth-service，Decision 2/3 无来源返回空字符串）
- `GET /api/decisions/{id}/candidates/modules`：返回候选 Module，正确排除已关联目标
- `POST /api/decisions`：返回 201 + `decision_id`（支持可选 `source_module_id` 参数）
- `POST /api/decisions/{id}/links`：返回 204 No Content

## 错误路径验证
- 必填字段缺失 → 400 `invalid input`
- 非法 status → 400 `invalid status`
- alternatives 条目空白 → 400 `invalid alternatives: items must not be blank`
- 重复关联 → 409 `decision link already exists`
- 目标类型越界 → 400 `invalid target type`
- 关联到不存在的 Decision → 404 `decision not found`
- 详情读取不存在的 decision_id → 404 `decision not found`
- [Issue 2] 创建带无效 source_module_id → 404 `module not found`

## source_context 验证（Issue 2 修复）
- buf lint / breaking 通过（.proto 添加 source_module_id 字段编号 9，非 breaking change）
- `go build ./...` / `go vet ./...` 通过
- GET detail for Decision 1（基线有 source）→ source_context 返回 `source_module_id` + `source_module_name: "auth-service"`
- GET detail for Decision 2/3（无 source）→ source_context 返回空字符串
- POST create with valid source_module_id → 持久化并在 detail 返回 source_context
- POST create with invalid source_module_id → 404 `module not found`
