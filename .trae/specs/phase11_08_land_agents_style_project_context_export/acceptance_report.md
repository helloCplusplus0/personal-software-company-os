# phase11-08 acceptance report

## 1. 验收范围

本报告用于补齐 `phase11-08` 的正式验收留档，覆盖以下两类证据：

1. `ExportProjectContext` 已具备正式只读承接位，并与 `GetProjectContext` 保持单向派生；
2. AGENTS 风格导出已能服务固定 `<= 3` 入口的新路径 dogfooding，且可回答预设的 `5` 个恢复问题。

## 2. 实现证据

### 2.1 正式承接位

- 结构化读取：
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/service/query_service.go`
  - `backend/internal/projectcontext/connect/server.go`
- Markdown 导出：
  - `proto/psco/project_context/v1/project_context.proto#ExportProjectContext`
  - `backend/internal/projectcontext/service/query_service.go#ExportProjectContext`
  - `backend/internal/projectcontext/renderer/markdown.go`

### 2.2 单向派生证据

- `ExportProjectContext` 先调用 `GetProjectContext` 获取同一 `repository_id` 的结构化结果，再交给 `renderer.RenderMarkdown` 渲染；
- `renderer` 不读取数据库、不扫描目录；
- 当前阶段边界摘要已纳入结构化只读结果的 `boundaries` 字段，由结构化结果统一提供，不再由 Markdown 层自行硬编码事实。

## 3. 自动化验证证据

已执行：

```bash
cd /home/dell/Projects/personal-software-company-os/proto && buf generate
cd /home/dell/Projects/personal-software-company-os/backend && go test ./internal/projectcontext/... ./internal/connecterrors/...
cd /home/dell/Projects/personal-software-company-os/backend && go build ./...
```

关键覆盖点：

- `renderer/markdown_test.go`
  - 验证 Rules / Current Phase 会输出受控引用；
  - 验证 Boundaries 来自结构化结果；
  - 验证空集合与 `nil Product` 路径。
- `connect/server_integration_test.go`
  - 验证 `GetProjectContext` 成功路径仍返回完整结构化上下文；
  - 验证 `ExportProjectContext` 成功路径返回 Markdown；
  - 验证导出结果包含：
    - 规则入口引用；
    - phase 文档入口引用；
    - 当前阶段边界摘要。

## 4. Dogfooding 入口协议

### 4.1 旧路径基线入口

1. `AGENTS.md`
2. `plan.md`
3. `project_rules.md`
4. `architecture_map.md`
5. `docs/README.md`
6. `PSCO-mvp05-summarize-feedback.md`

### 4.2 新路径目标入口

1. `AGENTS.md`
2. `PSCO-mvp05-summarize-feedback.md`
3. 基于同一 `repository_id` 生成的 AGENTS 风格项目上下文导出

结论：

- 新路径保持在 `<= 3` 固定入口约束内；
- AGENTS 风格导出作为第三入口，负责承接仓库摘要、phase 摘要、实体摘要、规则入口、phase 入口与明确不做边界；
- 新路径不需要临时增加第 `4` 个入口来补齐当前项目的结构化上下文摘要。

## 5. 五个恢复问题对照

### 5.1 当前 phase

- 回答入口：
  - `AGENTS.md`
  - AGENTS 风格项目上下文导出的 `Current Phase` 节

### 5.2 直接上游

- 回答入口：
  - `PSCO-mvp05-summarize-feedback.md`
  - AGENTS 风格项目上下文导出的 `Current Phase` / `Rules & Constraints` 节中的受控入口引用

### 5.3 单一主交付

- 回答入口：
  - `AGENTS.md`
  - `PSCO-mvp05-summarize-feedback.md`
  - AGENTS 风格项目上下文导出的 `Current Phase` 节

### 5.4 明确不做

- 回答入口：
  - AGENTS 风格项目上下文导出的 `Boundaries (What This Project Does NOT Do)` 节

### 5.5 当前项目关联的 Repository / Product / Module / Decision 摘要入口

- 回答入口：
  - AGENTS 风格项目上下文导出的：
    - `Repository`
    - `Product`
    - `Modules`
    - `Decisions`
    - `Rules & Constraints`

## 6. 边界复核

本次验收再次确认：

- 不主动写入 `AGENTS.md`；
- 不主动写入外部仓库；
- 不引入 CLI / MCP / agent 自动写回；
- 不把消费侧项目目录结构上升为输入合同；
- Markdown 导出不再形成独立于结构化读取之外的第二套事实源。

## 7. 验收结论

`phase11-08` 当前实现满足以下条件：

1. 已存在正式只读 Markdown 导出承接位；
2. 已存在“结构化读取 -> Markdown 导出”的单向派生实现；
3. 已补齐 `ExportProjectContext` 成功路径自动化验证；
4. 已补齐可复核的 dogfooding 入口协议与 `5` 问对照留档；
5. 已回收 renderer 层的第二事实源风险。

当前结论：`phase11-08` 通过本轮独立验收，可作为后续 `phase11-09` 联调与 dogfooding 的正式输入。
