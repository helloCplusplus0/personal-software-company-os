# phase11-09 Project Context Foundation 联调、dogfooding 与反回归验证 Spec

## Why
`phase11-06` 到 `phase11-08` 已分别完成根级真相源治理、最小结构化只读读取与 AGENTS 风格导出，但当前还缺少一轮正式的联调、dogfooding 与反回归收口。`phase11-09` 需要把“可实现”推进为“可客观复验”，避免 `Project Context Foundation` 只停留在抽象设计或局部实现通过。

## What Changes
- 冻结 `phase11-09` 的最小工具链验证矩阵与通过标准
- 冻结“旧路径基线 / 新路径目标”双路径 dogfooding 剧本、固定入口集合与固定提问集合
- 冻结根级真相源治理的反回归检查项，验证重复承载与悬空引用未回流
- 冻结本阶段边界证据的留档方式，明确不做 MCP / CLI / agent 写回 / 对话入口
- 冻结“PSCO 自身治理样本”与“跨项目通用能力合同”分离的验收口径

## Impact
- Affected specs:
  - `phase11_project_context_foundation`
  - `phase11_06_land_root_context_truth_source_governance`
  - `phase11_07_land_minimal_readonly_project_context_read_capability`
  - `phase11_08_land_agents_style_project_context_export`
- Affected code:
  - `README.md`
  - `AGENTS.md`
  - `plan.md`
  - `project_rules.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `docs/phase/README.md`
  - `proto/psco/project_context/v1/project_context.proto`
  - `backend/internal/projectcontext/`
  - `.trae/specs/phase11_08_land_agents_style_project_context_export/`

## ADDED Requirements
### Requirement: 冻结 phase11-09 的最小工具链验证矩阵
系统 SHALL 为 `phase11-09` 提供一组单值的最小工具链验证矩阵，用于验证 `Project Context Foundation` 的合同生成、后端实现与回归状态，而不是让执行者临场挑选命令。

当前阶段至少冻结：

- `proto/` 下的正式生成命令：`cd /home/dell/Projects/personal-software-company-os/proto && buf generate`；
- `backend/` 下的正式测试命令：`cd /home/dell/Projects/personal-software-company-os/backend && go test ./internal/projectcontext/... ./internal/connecterrors/...`；
- `backend/` 下的正式构建命令：`cd /home/dell/Projects/personal-software-company-os/backend && go build ./...`；
- 固定执行顺序：先生成、再测试、最后构建；
- 固定通过标准：
  - `buf generate` 必须退出码 `0`，并按 `proto/buf.gen.yaml` 完成生成，不得出现插件或配置错误；
  - `go test ./internal/projectcontext/... ./internal/connecterrors/...` 必须退出码 `0`，且 `projectcontext` 与 `connecterrors` 相关包测试全部通过；
  - `go build ./...` 必须退出码 `0`，且后端全部 package 可构建；
- 固定失败归类口径：
  - `命令失败`：命令正常启动但退出码非 `0`；
  - `环境异常`：命令无法启动或缺少前置环境；
  - `局部通过`：部分步骤成功、部分步骤失败；
  - 只要出现 `命令失败` 或 `环境异常`，工具链矩阵整体不得判定通过。

#### Scenario: 执行最小工具链验证
- **WHEN** 执行者准备开始 `phase11-09` 的联调与反回归验证
- **THEN** 必须先按冻结顺序执行最小工具链验证
- **AND** 不得以“局部文件通过”替代全链路通过结论

### Requirement: 冻结旧路径基线与新路径目标的 dogfooding 剧本
系统 SHALL 冻结一份可重复执行的 dogfooding 剧本，并要求同一验收者对“旧路径基线”与“新路径目标”各执行一次，以比较上下文恢复成本是否真实下降。

旧路径基线入口集合固定为：

1. `AGENTS.md`
2. `plan.md`
3. `project_rules.md`
4. `architecture_map.md`
5. `docs/README.md`
6. `PSCO-mvp05-summarize-feedback.md`

新路径目标入口集合固定为：

1. `AGENTS.md`
2. `PSCO-mvp05-summarize-feedback.md`
3. 基于同一 `repository_id` 生成的 AGENTS 风格项目上下文导出

补充冻结：

- 旧路径执行时不得额外读取 `phase11` 新增导出结果；
- 新路径执行时不得临时增加第 `4` 个入口补答案；
- 两条路径必须使用同一份问题清单与同一套达标口径。

#### Scenario: 执行双路径 dogfooding
- **WHEN** 执行者按固定入口集合分别执行旧路径与新路径
- **THEN** 必须能够记录两条路径各自的入口清单、回答结果、失败点与是否达标
- **AND** 新路径结果必须可与旧路径基线直接对照

### Requirement: 冻结五个恢复问题与回答口径
系统 SHALL 将 `phase11-09` 的 dogfooding 提问固定为 `5` 问，并要求新路径必须完整回答，不再依赖散落的根级多文档补读。

固定提问如下：

1. 当前 phase 是什么；
2. 直接上游是什么；
3. 单一主交付是什么；
4. 当前明确不做什么；
5. 当前项目关联的 `Repository / Product / Module / Decision` 摘要入口是什么。

回答格式冻结为：

`问题 -> 直接回答 -> 证据入口 -> 是否达标 -> 失败点`

达标口径冻结为：

- `直接回答` 必须能从固定入口追溯；
- `证据入口` 不得超出对应路径允许的入口集合；
- 新路径必须 `5/5` 全部达标；
- 旧路径作为基线对照，可出现失败项，但必须写明失败点。

#### Scenario: 评估新路径是否达标
- **WHEN** 验收者使用新路径固定入口集合回答 `5` 个恢复问题
- **THEN** 必须能够完整回答全部问题
- **AND** 回答结果必须可追溯到固定入口，而不是依赖额外补读或主观记忆

### Requirement: 冻结根级真相源治理的反回归检查项
系统 SHALL 为 `phase11-09` 冻结根级真相源治理的反回归检查项，验证 `phase11-06` 收口后的根级入口仍保持单值一致，不再重复承载主结论，也不重新引入已清理的悬空引用。

检查范围至少包括：

- `README.md`
- `AGENTS.md`
- `plan.md`
- `project_rules.md`
- `architecture_map.md`
- `docs/README.md`
- `docs/phase/README.md`

判定口径冻结为：

- `重复 phase 状态`：非 `plan.md` 文件把 phase 状态写成新的 canonical 正文，而不是摘要式入口；
- `重复目录落点`：非 `architecture_map.md` 文件把目录或文档落点写成新的 canonical 正文；
- `重复技术栈正文`：非 `TECH_STACK_BASELINE.md` 文件重写完整技术栈正文，而不是摘要或规则引用；
- `悬空引用回流`：目标文件重新出现 `PSCO-summarize-feedback.md` 等失效入口；
- 摘要式入口允许存在，但不得上升为第二个 canonical writer。

#### Scenario: 执行根级反回归复核
- **WHEN** 执行者复核根级入口文档
- **THEN** 必须逐项记录是否出现重复 phase 状态、重复目录落点、重复技术栈正文或悬空引用回流
- **AND** 不得仅凭“主观看起来差不多”判定通过

### Requirement: 冻结边界证据与样本/合同分离口径
系统 SHALL 为 `phase11-09` 留档两类边界证据，避免 dogfooding 结果被误读为更大范围的既成事实。

必须明确记录：

- 当前阶段不做 MCP / CLI / agent 写回 / 前端对话入口；
- 当前 dogfooding 只证明 `PSCO` 自身仓库样本可消费；
- `PSCO` 当前仓库文件清单只用于自身治理与 dogfooding，不等于未来所有项目都必须复制同样目录结构。

#### Scenario: 复核边界证据
- **WHEN** 后续读者查看 `phase11-09` 验收结果
- **THEN** 必须能直接看到当前阶段明确不做的边界与“样本 != 模板”的说明
- **AND** 不会将本次 dogfooding 误读为对未来项目模板的冻结

## MODIFIED Requirements
### Requirement: phase11 完成条件中的 dogfooding 验收
`phase11` 的完成条件 SHALL 从“已存在 dogfooding 方向与入口约束”推进为“已完成正式双路径 dogfooding、反回归复核与边界留档，并形成可独立复验的验收记录”。

补充要求：

- 必须同时留存旧路径基线与新路径目标两次执行记录；
- 必须留存 `5` 个恢复问题的回答结果；
- 必须留存根级反回归复核结果；
- 必须留存边界证据与样本/合同分离说明。

#### Scenario: 判断 phase11 是否可进入收口
- **WHEN** 执行者准备声明 `phase11-09` 完成
- **THEN** 必须能指出正式验收记录、双路径对照、反回归复核结果与边界证据
- **AND** 不得仅以“已有导出能力且个人感觉更顺手”作为通过依据

## REMOVED Requirements
### Requirement: 以主观体感替代 dogfooding 验收结果
**Reason**: `phase11-09` 的目标是把 `Project Context Foundation` 的效果收敛成可复验的证据，而不是停留在执行者个人感受。
**Migration**: 后续验收统一改为留存固定入口清单、固定问题回答、失败点、达标结论与边界说明，不再接受“感觉恢复成本下降”作为唯一结论。 
