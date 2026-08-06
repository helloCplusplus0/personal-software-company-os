- [x] 已明确 `decisioncenter` 后端模块文件落点与 `phase03-10` §10.2 完全一致
- [x] 已明确 `types.go` 从 `.proto` 单向承接语义，不形成第二套合同源
- [x] 已明确 `errors.go` 覆盖 `phase03-04` 全部错误语义（7 个哨兵错误）
- [x] 已明确 `handler/response.go` 错误到 HTTP 状态码映射（404 / 409 / 400 / 500）
- [x] 已明确 `RecordDecision` 校验顺序（必填 → status → alternatives → source_module_id 可选来源校验）
- [x] 已明确 `LinkDecisionToTarget` 校验顺序（target_type → Decision 存在性 → Module 存在性 → 重复关联）
- [x] 已明确 `link_count` 计算口径（仅统计 `decision_links` 有效关联数，无关联返回 0）
- [x] 已明确 `linked_module_summary` 计算口径（按 `module_name` 升序取前 3 + `+N`，无关联返回空字符串）
- [x] 已明确候选读取排序（`status(active 优先) -> module_name 升序`）与排除规则（已关联不得再次出现）
- [x] 已明确 `ModuleCandidateRead` 接口由 `candidate/` 子包拥有，在应用装配点接线
- [x] 已明确 `DecisionModuleCandidate.status` 复用 `moduleregistry.ModuleStatus`，不重定义本地枚举
- [x] 已明确 `0004` migration 通过 `ALTER TABLE` 原位升级，不新建替代表
- [x] 已明确无 `DEFAULT` 必填字段按"ADD COLUMN 允许 NULL → 回填 → SET NOT NULL"三步流程
- [x] 已明确 `alternatives` 使用 `TEXT[]`，不使用 `JSONB`
- [x] 已明确现有示例数据兼容回填策略（占位文本 + `'{}'` + `''` + `'proposed'`）
- [x] 已明确基线 seed 覆盖维度（3 条 Decision + 2 条 decision_links + 1 条无关联 Decision）
- [x] 已明确 `seed_readonly_prereqs.sql` 中 `decisions` seed 从 `title-only` 升级为结构化字段
- [x] 已明确 `reset_decision_mainline.sh` 与 `reset_module_mainline.sh` 同构
- [x] 已明确清空范围为 `DELETE FROM decisions`（级联清空 `decision_links`），不清空 `modules`
- [x] 已明确应用装配在 `platform/router.go` + `server.go` 中完成
- [x] 已明确路由注册对齐 `phase03-10` §7.7 RPC → HTTP 映射矩阵
- [x] 已明确后端可编译、migration 可执行、重置脚本可运行
- [x] 已明确 `source_context` 入口上下文来源通过 `decisions.source_module_id` 持久化承接（`phase03-10 §5.11`"持续到完成或放弃"跨刷新语义）
- [x] 已明确 `CreateDecision` 第 4 步校验 `source_module_id` 来源存在性（非空时校验，空字符串跳过，无效返回 `ErrModuleNotFound`）
- [x] 已明确 `DecisionDetailRead` 通过 `source_context` 持久化返回来源 `Module` 标识（`LEFT JOIN modules`，无来源返回空字符串）
- [x] 已明确 `0005_decision_source_context.sql` 独立 migration 添加 `source_module_id` 字段（`ON DELETE SET NULL`），幂等执行
- [x] 已明确 `.proto` `CreateDecisionRequest` 新增 `source_module_id` 字段（编号 `9`，非 breaking change，`buf lint` / `buf breaking` 通过）
- [x] 已明确基线 seed `Decision 1` 带 `source_module_id`（关联 `auth-service`），幂等语义为 `UPDATE` 收口 + `INSERT` 补缺 + `WHERE NOT EXISTS` + `ON CONFLICT`
- [x] 已明确本 spec 可直接作为 `phase03-13 / 14` 的后端 API 上游

# 验收证据补充

## 编译验证
- `go build ./...` 通过
- `go vet ./...` 通过

## Migration 验证
- `0004_decision_center_mainline` 成功应用到已有 `0001-0003` 基础上
- `decisions` 表升级后全部新字段 `NOT NULL`，`CHECK` 约束与索引就位
- `0005_decision_source_context` 成功应用：新增 `source_module_id UUID REFERENCES modules(id) ON DELETE SET NULL` 列与索引

## 重置脚本幂等性验证
- `--clean-only` / 默认 / `--restore-only`×2 三种模式均幂等
- `--restore-only` 两次执行结果一致（`decisions=3 / decision_links=2`）
- [Issue 1 修复] `--restore-only` 在已有 readonly placeholder 时，UPDATE 收口 placeholder 为基线内容（不保留占位文本），计数与内容均恢复到正式基线

## API 端点与错误路径验证
- 5 个 API 端点全部可访问且返回正确语义
- 8 类错误路径全部返回正确 HTTP 状态码（400 / 404 / 409）与错误消息
- 无 500 级未收口错误替代业务错误
- [Issue 2 修复] source_context 有值/无值/创建带 source/创建带无效 source 全部验证通过

## 修复记录
- 修复 `seed_decision_mainline_baseline.sql` 中 `decisions` INSERT 的幂等性问题：因 `decisions.title` 无 UNIQUE 约束，`ON CONFLICT DO NOTHING` 无法触发冲突，改用 `WHERE NOT EXISTS` 守卫（与 `seed_readonly_prereqs.sql` 中 decisions seed 既有模式一致），`decision_links` 仍使用 `ON CONFLICT (decision_id, module_id) DO NOTHING`
- 在 `candidate/module_candidate_read.go` 中新增 `ModuleExists` 方法承接 `LinkDecisionToTarget` 的 Module 存在性跨模块只读校验，确保 service 层不直接写跨模块 SQL（phase03-10 §10.5 约束）
- [独立复核 P2-1 修复] 清理 `decision_store.go::List` 中 `linked_module_summary` SQL 的死代码（`WHEN COUNT(*) = 0` 不可达分支），简化为直接 `COALESCE(string_agg(...) || CASE ..., '')`，语义不变
- [独立复核 P2-3 修复] 强化 `candidate/module_candidate_read.go` 编译期断言为 `var _ moduleregistry.ModuleStatus = decisioncenter.DecisionModuleCandidate{}.Status`，真正约束字段类型
- [GPT-5.4 复核 Issue 1 修复] `seed_decision_mainline_baseline.sql` 中 Decision 1（复用 phase02 标题）增加 UPDATE 语句收口已有 readonly placeholder 为正式基线内容，再 INSERT 补缺，确保 `--restore-only` 模式也能恢复到同一份正式基线
- [GPT-5.4 复核 Issue 2 修复] 补齐后端 `source_context` 承接（phase03-10 §5.11"持续到完成或放弃"要求跨刷新持久化）：
  - 新增 `database/migrations/0005_decision_source_context.sql`：decisions 表添加 `source_module_id UUID REFERENCES modules(id) ON DELETE SET NULL`
  - `.proto` CreateDecisionRequest 添加 `source_module_id` 字段（编号 9，非 breaking change，buf lint/breaking 通过）
  - `types.go` CreateDecisionRequest 添加 `SourceModuleID` 字段
  - `repository/decision_store.go` 新增 `DecisionWithSource` 结构体，Create 写入 source_module_id，GetByID 通过 LEFT JOIN modules 读取 source_module_name
  - `service/command_service.go` CreateDecision 增加第 4 步校验：source_module_id 非空时校验 Module 存在性（跨模块只读校验由 candidate 子包承接），无效返回 ErrModuleNotFound
  - `service/query_service.go` GetDecisionDetail 组装 SourceContext（从 DecisionWithSource 提取）
  - 基线 seed Decision 1 添加 source_module_id（关联 auth-service 模块）
