# phase11-09 acceptance report

## 1. 文档定位

本报告用于完成 `phase11-09` 的 Task 1-5 冻结与验收留档，只承接：

1. 最小工具链验证矩阵与实际执行结果；
2. 旧路径基线 / 新路径目标 dogfooding 双路径入口清单；
3. 固定 `5` 问的回答结果与是否达标；
4. 根级真相源治理的反回归复核结果；
5. 不做 `MCP / CLI / agent 写回 / 对话入口` 的边界证据；
6. “PSCO 自身治理样本 != 跨项目模板”的分离说明。

本报告不扩大到新代码、新协议层、新 seed 或新写回通道。

## 2. 上游输入与范围冻结

### 2.1 直接上游与 DoD

本轮冻结直接承接以下上游：

- [phase11 dev plan:L241-L278](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase11_project_context_foundation_dev_plan.md#L241-L278)
- [phase11 shared baseline:L13-L247](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase11_project_context_foundation_shared_baseline.md#L13-L247)
- [phase11-08 acceptance report:L5-L132](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase11_08_land_agents_style_project_context_export/acceptance_report.md#L5-L132)

据此冻结 `phase11-09` 范围：

- 只承接联调、dogfooding 与反回归验证；
- 不新增 `MCP / CLI / agent 写回 / 前端对话入口`；
- 验收必须留档双路径入口、`5` 问回答、失败点、是否达标；
- DoD 继续以“固定入口 <= 3、可客观复验、不是主观体感”为准。

### 2.2 `phase11-06 ~ 08` 正式输入与当前实现状态

本轮只复用既有正式输入，不重新发明主线：

| 子任务 | 正式输入 | 当前实现/状态 |
| --- | --- | --- |
| `phase11-06` | [spec](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase11_06_land_root_context_truth_source_governance/spec.md) | 根级真相源治理已落到 `README.md / AGENTS.md / plan.md / architecture_map.md / docs/README.md / docs/phase/README.md / project_rules.md` 等入口；本轮只做反回归复核。 |
| `phase11-07` | [spec](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase11_07_land_minimal_readonly_project_context_read_capability/spec.md) | `repository_id` 唯一输入锚点、只读聚合、失败语义、`.proto + ConnectRPC` 主线已落到 `proto/psco/project_context/v1/project_context.proto`、`backend/internal/projectcontext/service/query_service.go`、`backend/internal/projectcontext/connect/server.go`。 |
| `phase11-08` | [spec](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase11_08_land_agents_style_project_context_export/spec.md)、[acceptance report](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase11_08_land_agents_style_project_context_export/acceptance_report.md) | `ExportProjectContext` 已从结构化结果单向派生 Markdown 导出，并已冻结 `<= 3` 入口 dogfooding 协议；本轮只做验收收口，不改导出实现。 |

### 2.3 固定样本冻结与范围发现

固定样本冻结如下：

- 第一 dogfooding 样本：`PSCO` 自身仓库；
- 该样本身份：当前 phase 的治理对象与验收对象；
- 该样本不是：未来所有项目的目录模板、固定 seed 模板或通用 contract。

范围发现（已在本轮补闭环）：

1. 当前仓库内可复用的自动化 fixture 仍以 `database/seeds/seed_phase06_fixture_*` 的 `main-repo` / `mirror-repo` 为主，最初确实不存在“`PSCO` 自身仓库已完成 repository binding”的现成 seed 证据；
2. 当前集成验证也仍通过 `backend/internal/projectcontext/connect/server_integration_test.go` 中的 `main-repo` fixture 完成；
3. 当前 renderer 单测虽然含有 `personal-software-company-os / PSCO` 文案样本，但它是渲染样本，不等于已登记的 `repository_id` 绑定样本；
4. 针对上述缺口，本轮已在开发库中补齐 `PSCO` 自身仓库样本，并以正式读取路径完成终验。

本轮处理口径：

- 保留“此前缺少现成 fixture”的事实留档；
- 允许在 `psco_development` 中补齐最小 dogfooding 样本数据；
- 不新增 `MCP / CLI / agent 写回 / 前端对话入口`；
- 不把“补齐样本证据”扩写成新的产品主线或 schema 变更。

## 3. 最小工具链验证矩阵

### 3.1 Context7 口径确认

编码前已通过 Context7 确认本轮命令语义：

- `Buf`：选择 `/bufbuild/buf`，确认 `buf generate` 默认读取当前目录下的 `buf.gen.yaml`；
- `Go`：选择 `/golang/go`，确认 `go test [packages]`、`go build [packages]` 为正式命令形态。

因此本轮正式工具链命令冻结为：

1. `cd /home/dell/Projects/personal-software-company-os/proto && buf generate`
2. `cd /home/dell/Projects/personal-software-company-os/backend && go test ./internal/projectcontext/... ./internal/connecterrors/...`
3. `cd /home/dell/Projects/personal-software-company-os/backend && go build ./...`

### 3.2 执行矩阵与实际结果

| 步骤 | 命令 | 通过标准 | 实际结果 |
| --- | --- | --- | --- |
| 1 | `cd /home/dell/Projects/personal-software-company-os/proto && buf generate` | 退出码 `0`；按 `proto/buf.gen.yaml` 完成生成；不得出现插件或配置错误 | 通过。退出码 `0`，无 stderr。 |
| 2 | `cd /home/dell/Projects/personal-software-company-os/backend && go test ./internal/projectcontext/... ./internal/connecterrors/...` | 退出码 `0`；`projectcontext` 与 `connecterrors` 相关包测试全部通过 | 通过。输出：`projectcontext/connect`、`projectcontext/renderer`、`connecterrors` 为 `ok`，其余相关包为 `[no test files]`。 |
| 3 | `cd /home/dell/Projects/personal-software-company-os/backend && go build ./...` | 退出码 `0`；后端全部 package 可构建 | 通过。退出码 `0`，无 stderr。 |

### 3.3 失败、环境异常与局部通过归类口径

本轮冻结以下归类口径，后续不允许临场解释：

- `命令失败`：命令正常启动但退出码非 `0`；
- `环境异常`：命令无法启动或缺少前置环境，例如缺少 `buf`、缺少 `go`、缺少测试数据库；
- `局部通过`：部分步骤成功、部分步骤失败。

统一判定：

- 只要出现 `命令失败` 或 `环境异常`，本轮工具链矩阵整体不得判定通过；
- `局部通过` 必须逐步记录“哪一步通过、哪一步失败”，但总结果仍为不通过；
- 不允许用“`go test` 通过所以可忽略 `go build` 失败”这类口径替代全链路结论。

## 4. 双路径 dogfooding 剧本与固定五问

### 4.1 固定入口集合

旧路径基线固定 `6` 个入口，不允许额外补读：

1. `AGENTS.md`
2. `plan.md`
3. `project_rules.md`
4. `architecture_map.md`
5. `docs/README.md`
6. `PSCO-mvp05-summarize-feedback.md`

新路径目标固定 `3` 个入口，不允许增加第 `4` 个入口：

1. `AGENTS.md`
2. `PSCO-mvp05-summarize-feedback.md`
3. 基于同一 `repository_id` 生成的 AGENTS 风格项目上下文导出

### 4.2 固定五问、回答格式与达标口径

固定 `5` 问：

1. 当前 phase 是什么？
2. 直接上游是什么？
3. 单一主交付是什么？
4. 当前明确不做什么？
5. 当前项目关联的 `Repository / Product / Module / Decision` 摘要入口是什么？

回答格式冻结为：

`问题 -> 直接回答 -> 证据入口 -> 是否达标 -> 失败点`

达标口径冻结为：

- `直接回答` 必须能从固定入口追溯；
- `证据入口` 不得超出对应路径允许的入口集合；
- 新路径必须 `5/5` 全部达标；
- 旧路径作为基线对照，可出现失败项，但必须写明失败点。

### 4.3 共同记录模板

双路径共同使用以下留档模板：

| 路径 | 问题 | 直接回答 | 证据入口 | 是否达标 | 失败点 |
| --- | --- | --- | --- | --- | --- |

### 4.4 实际双路径结果

#### 4.4.1 旧路径基线

| 路径 | 问题 | 直接回答 | 证据入口 | 是否达标 | 失败点 |
| --- | --- | --- | --- | --- | --- |
| 旧路径 | 当前 phase 是什么？ | `phase11_project_context_foundation` | `AGENTS.md`、`plan.md` | 达标 | 无 |
| 旧路径 | 直接上游是什么？ | `PSCO-mvp05-summarize-feedback.md` 是唯一共识上游；`phase10` 已完成交付作为进入条件 | `AGENTS.md`、`plan.md`、`PSCO-mvp05-summarize-feedback.md` | 达标 | 无 |
| 旧路径 | 单一主交付是什么？ | 根级上下文真相源治理 + 最小只读项目上下文导出 | `AGENTS.md`、`plan.md`、`PSCO-mvp05-summarize-feedback.md` | 达标 | 无 |
| 旧路径 | 当前明确不做什么？ | 不做 `MCP / CLI / 前端对话入口 / 自动扫描 / 知识图谱 / agent 自动写回` | `plan.md`、`project_rules.md`、`PSCO-mvp05-summarize-feedback.md` | 达标 | 无 |
| 旧路径 | 当前项目关联的 `Repository / Product / Module / Decision` 摘要入口是什么？ | 固定 `6` 入口内没有单值、结构化、受控的四实体摘要入口 | 固定 `6` 入口交叉复核 | 不达标 | 仍需补读 phase11 导出或实现细节，无法在旧路径入口集内完成回答 |

旧路径结论：

- 入口数量：`6`
- `5` 问结果：`4/5`
- 主要失败点：四实体摘要入口仍然缺少固定单值承接位
- 对照结论：旧路径可恢复总体 phase 上下文，但不能稳定给出项目实体摘要入口

#### 4.4.2 新路径目标

| 路径 | 问题 | 直接回答 | 证据入口 | 是否达标 | 失败点 |
| --- | --- | --- | --- | --- | --- |
| 新路径 | 当前 phase 是什么？ | `phase11_project_context_foundation` | `AGENTS.md`、导出中的 `Current Phase` 节 | 达标 | 无 |
| 新路径 | 直接上游是什么？ | `PSCO-mvp05-summarize-feedback.md` 为唯一共识上游；导出中的 `Rules & Constraints / Current Phase` 负责给出受控回看入口 | `AGENTS.md`、`PSCO-mvp05-summarize-feedback.md`、导出中的 `Current Phase / Rules & Constraints` | 达标 | 无 |
| 新路径 | 单一主交付是什么？ | 根级上下文真相源治理 + 最小只读项目上下文导出 | `AGENTS.md`、`PSCO-mvp05-summarize-feedback.md`、导出中的 `Current Phase` | 达标 | 无 |
| 新路径 | 当前明确不做什么？ | 不做 `MCP / CLI / agent 写回 / 前端对话入口 / 第二套事实源 / 消费侧目录模板合同` | `PSCO-mvp05-summarize-feedback.md`、导出中的 `Boundaries` | 达标 | 无 |
| 新路径 | 当前项目关联的 `Repository / Product / Module / Decision` 摘要入口是什么？ | 导出中的 `Repository / Product / Modules / Decisions / Rules & Constraints` 节 | 导出中的对应节；`phase11-08` 验收已冻结这些入口职责 | 达标 | 无 |

新路径结论：

- 入口数量：`3`
- `5` 问结果：`5/5`
- 新路径达标；
- 相比旧路径，减少了对散落根级文档补读的依赖。

## 5. 根级真相源治理反回归复核

### 5.1 复核范围与判定口径

复核范围冻结为：

- `README.md`
- `AGENTS.md`
- `plan.md`
- `project_rules.md`
- `architecture_map.md`
- `docs/README.md`
- `docs/phase/README.md`

判定口径冻结为：

- `重复 phase 状态`：非 `plan.md` 文件把 phase 状态写成新的 canonical 正文，而不是摘要式入口；
- `重复目录落点`：非 `architecture_map.md` 文件把目录/文档落点写成新的 canonical 正文；
- `重复技术栈正文`：非 `TECH_STACK_BASELINE.md` 文件重写完整技术栈正文，而不是摘要或规则引用；
- `悬空引用回流`：目标文件重新出现 `PSCO-summarize-feedback.md` 等失效入口。

补充口径：

- 摘要式入口允许存在；
- “允许摘要”不等于“允许第二个 canonical writer”。

### 5.2 实际复核结果

先执行了：

```bash
rg -n 'PSCO-summarize-feedback\.md' \
  README.md AGENTS.md plan.md project_rules.md architecture_map.md docs/README.md docs/phase/README.md
```

结果：`No matches found`。

逐项复核结果：

| 文件 | 正式职责 | 重复 phase 状态 | 重复目录落点 | 重复技术栈正文 | 悬空引用回流 | 结论 |
| --- | --- | --- | --- | --- | --- | --- |
| `README.md` | 项目总览入口 | 否，仅入口跳转 | 否 | 否 | 否 | 通过 |
| `AGENTS.md` | 入口摘要 | 否，保留摘要，不替代 `plan.md` | 否 | 否，仅技术路线摘要 | 否 | 通过 |
| `plan.md` | phase 状态唯一正式承接位 | 否，本文件即 canonical writer | 否 | 否，仅路线级摘要 | 否 | 通过 |
| `project_rules.md` | 规则与门禁 | 否 | 否 | 否，承接的是规则与选择约束，不是完整 baseline 正文 | 否 | 通过 |
| `architecture_map.md` | 目录结构、文档分类、迁移落点 | 否，未替代 `plan.md` 成为 phase 状态 writer | 否，本文件即 canonical writer | 否，仅引用 `TECH_STACK_BASELINE.md` | 否 | 通过 |
| `docs/README.md` | docs workflow 总入口 | 否，显式回指 `plan.md` | 否 | 否 | 否 | 通过 |
| `docs/phase/README.md` | phase 索引入口 | 否，显式回指 `plan.md` | 否 | 否 | 否 | 通过 |

反回归复核结论：

- 目标范围内未发现 `PSCO-summarize-feedback.md` 悬空引用回流；
- 未发现新的第二个 phase 状态 canonical writer；
- 未发现新的第二个目录落点 canonical writer；
- 未发现新的技术栈正文回流。

### 5.3 留档格式冻结

后续验收复核统一采用以下格式，不再口头判定：

| 文件 | 正式职责 | 4 类检查项结果 | 关键证据 | 结论 |
| --- | --- | --- | --- | --- |

本报告第 `5.2` 节即为正式样例。

## 6. 边界证据与样本/合同分离

### 6.1 不做 `MCP / CLI / agent 写回 / 对话入口` 的证据

本轮边界证据来自三层：

1. [phase11 shared baseline:L48-L85](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase11_project_context_foundation_shared_baseline.md#L48-L85)
2. [phase11-08 acceptance report:L112-L132](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase11_08_land_agents_style_project_context_export/acceptance_report.md#L112-L132)
3. [context_readers.go:L269-L298](file:///home/dell/Projects/personal-software-company-os/backend/internal/projectcontext/candidate/context_readers.go#L269-L298)

冻结结论：

- 当前阶段不做 `MCP` 协议层；
- 当前阶段不做 `CLI` 工具；
- 当前阶段不做 `agent` 自动写回、`Draft`、审批流；
- 当前阶段不做 web 前端对话式 agent 入口；
- 当前阶段不形成第二套事实源；
- 当前阶段不把消费侧目录结构上升为输入合同。

### 6.2 `PSCO` 样本与跨项目合同分离说明

以下结论已冻结：

- `PSCO` 自身仓库是当前第一 dogfooding 样本；
- 该样本只证明“PSCO 自身治理结果可被当前导出能力消费”；
- 它不等于未来所有项目都必须复制：
  - 同名根级文件；
  - 同样的目录结构；
  - 同样的 seed 方式；
  - 同样的 dogfooding 入口清单。

证据：

- [shared baseline:L72-L85](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase11_project_context_foundation_shared_baseline.md#L72-L85)
- [shared baseline:L170-L185](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase11_project_context_foundation_shared_baseline.md#L170-L185)
- [PSCO-mvp05:L119-L124](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L119-L124)
- [PSCO-mvp05:L192-L200](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L192-L200)

### 6.3 正式结论表达冻结

本阶段正式表述冻结为：

> 本轮 dogfooding 只证明 `PSCO` 自身治理样本在当前 phase 边界内可被固定 `<= 3` 入口消费；
> 该结论不自动上升为跨项目模板、固定目录合同或未来所有项目必须复制的 agent 上下文入口方案。

## 7. Task 1-5 对应关系

| Task | 本报告承接位置 | 结果 |
| --- | --- | --- |
| Task 1 | 第 `2` 节 | 已冻结 |
| Task 2 | 第 `3` 节 | 已冻结并补齐实际执行结果 |
| Task 3 | 第 `4` 节 | 已冻结并补齐双路径回答结果 |
| Task 4 | 第 `5` 节 | 已冻结并补齐反回归复核结果 |
| Task 5 | 第 `6` 节 | 已冻结并补齐边界/样本分离证据 |

## 8. `PSCO` 自身仓库样本补闭环证据

### 8.1 正式样本与 `repository_id`

本轮在开发库 `psco_development` 中补齐了 `PSCO` 自身 dogfooding 样本，正式样本为：

- `Repository`：`personal-software-company-os`
- `repository_id`：`ca261521-8daf-4248-8f12-43525326e759`
- `Product`：`PSCO`
- `Module`：`project-context-foundation`
- `Decision`：`phase11 Project Context Foundation dogfooding 验收决策`

正式绑定证据查询结果：

| repository_id | repository | product | module | decision |
| --- | --- | --- | --- | --- |
| `ca261521-8daf-4248-8f12-43525326e759` | `personal-software-company-os` | `PSCO` | `project-context-foundation` | `phase11 Project Context Foundation dogfooding 验收决策` |

### 8.2 实际补链命令

本轮使用开发库最小补链方式建立正式样本：

1. 向 `repositories` 写入 `personal-software-company-os`
2. 向 `products` 写入 `PSCO`
3. 向 `modules` 写入 `project-context-foundation`
4. 建立：
   - `product_repositories`
   - `module_repositories`
   - `product_modules`
   - `decision_links`
5. 写入 `Decision`：`phase11 Project Context Foundation dogfooding 验收决策`

补充说明：

- 本次补链只发生在开发库 `psco_development`；
- 它的身份是 `phase11-09` 验收样本补闭环，不是新业务能力；
- 它不改变 `phase11` 的只读导出边界，也不改变未来跨项目合同。

### 8.3 基于同一 `repository_id` 的正式读路径终验

先启动后端：

```bash
cd /home/dell/Projects/personal-software-company-os/backend && go run ./cmd/server
```

服务实际监听：

- `healthz`：`http://127.0.0.1:8081/healthz`
- Connect 前缀：`http://127.0.0.1:8081/api`

实际终验命令：

```bash
curl --header "Content-Type: application/json" \
  --data '{"repositoryId":"ca261521-8daf-4248-8f12-43525326e759"}' \
  http://127.0.0.1:8081/api/psco.project_context.v1.ProjectContextService/GetProjectContext

curl --header "Content-Type: application/json" \
  --data '{"repositoryId":"ca261521-8daf-4248-8f12-43525326e759"}' \
  http://127.0.0.1:8081/api/psco.project_context.v1.ProjectContextService/ExportProjectContext
```

终验结果：

- `GetProjectContext`：成功返回 `Repository / Product / Modules / Decisions / Rules / Phases / Boundaries`
- `ExportProjectContext`：成功返回 AGENTS 风格 Markdown 导出
- 命中的 `Decision` 同时包含：
  - `bound_product_module`
  - `repository_module_mapping`

这说明当前导出能力已经基于 `PSCO` 自身样本，而不是仅基于 `main-repo` fixture 完成验证。

## 9. 当前验收结论

本轮此前的阻塞已经补齐，当前正式判定更新为：

- 工具链验证：通过；
- 根级反回归：通过；
- 导出能力协议验证：通过；
- `PSCO` 自身仓库 dogfooding：通过；
- `phase11-09` 总体验收结论：**通过**。

通过依据如下：

1. `PSCO` 自身仓库已具备正式 `repository_id` 与完整 binding 证据；
2. 已基于同一 `repository_id` 执行新路径固定 `3` 入口 dogfooding；
3. 新路径对固定 `5` 问达到 `5/5`；
4. 旧路径基线与新路径目标的差异仍被客观保留；
5. 边界证据与样本/合同分离说明仍然成立。

## 10. 对 phase11 最终复核预期的当前判断

在 `phase11-09` 通过之后，`phase11` 当前已经满足“最终成果复核预期”的验收侧前提：

- 根级真相源治理已被真实复核；
- 最小只读项目上下文读取能力已被真实复核；
- AGENTS 风格导出已被真实复核；
- `PSCO` 自身仓库样本 dogfooding 已闭环；
- 双路径对照、边界证据与样本/合同分离说明均已留档。

因此，当前剩余工作不再是“补验收闭环”，而是进入 `phase11-10` 的根级同步与文档收口。
