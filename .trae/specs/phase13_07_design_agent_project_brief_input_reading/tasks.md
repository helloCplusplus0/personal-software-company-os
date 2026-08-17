# Tasks

- [x] Task 1: 冻结 `project brief for agent` 的第一版最小 schema
  - [x] SubTask 1.1: 明确 7 个顶层字段集合
  - [x] SubTask 1.2: 明确各顶层字段与上游 `phase13-04 / 13-05` 的对应关系
  - [x] SubTask 1.3: 复核第一版不会长出第 8 个临时顶层块

- [x] Task 2: 冻结 brief 的组合语义与承接边界
  - [x] SubTask 2.1: 明确 `repository / governance_profile / global_assets / current_phase` 的承接职责
  - [x] SubTask 2.2: 明确 `products[] / modules[] / decisions[]` 的数组摘要语义
  - [x] SubTask 2.3: 冻结各顶层块的最小子字段矩阵
  - [x] SubTask 2.4: 复核 brief 不会退化为 prompt 片段拼装或目录扫描结果

- [x] Task 3: 冻结 brief 的唯一锚点与解析协议
  - [x] SubTask 3.1: 明确 `repository_id` 是唯一正式锚点
  - [x] SubTask 3.2: 明确多 `Module / Decision` 必须返回数组摘要
  - [x] SubTask 3.3: 复核不会伪造单一“当前 module / decision”

- [x] Task 4: 冻结 brief 与 IDE 目录读取能力的协作边界
  - [x] SubTask 4.1: 明确 brief 只承接结构化只读事实
  - [x] SubTask 4.2: 明确目录全文、局部源码与即时工作区上下文继续由 IDE / agent 现场能力读取
  - [x] SubTask 4.3: 复核第一版不会长出第二套目录扫描机制

- [x] Task 5: 冻结后端、前端与 agent 消费侧的共享 schema 规则
  - [x] SubTask 5.1: 明确后端是第一版 brief 的唯一正式生成来源
  - [x] SubTask 5.2: 明确前端不得派生“前端理解版 brief”
  - [x] SubTask 5.3: 明确 agent 消费侧不得跳过正式 schema 自行拼装第二版协议

- [x] Task 6: 冻结 brief 与现有 `ProjectContextService` 的正式关系
  - [x] SubTask 6.1: 明确 brief 是现有 `ProjectContextService` 主线内的受控演进，而不是并列第二 service
  - [x] SubTask 6.2: 明确 `GetProjectContext / ExportProjectContext` 的兼容层或退役策略必须在 formal spec 中声明
  - [x] SubTask 6.3: 复核 `product -> products[]` 与 `rules / phases / boundaries -> governance_profile / global_assets / current_phase` 的映射不再留给实现阶段临场决定

- [x] Task 7: 冻结第一版 brief 的非目标与禁止事项
  - [x] SubTask 7.1: 明确全文、Git 推进、模板自动同步状态不进入第一版 brief
  - [x] SubTask 7.2: 明确前端不承担 brief 主消费职责
  - [x] SubTask 7.3: 复核 brief 不是未结构化 prompt 文本

- [x] Task 8: 完成 spec 包与上游冻结文档的一致性校验
  - [x] SubTask 8.1: 校验本 spec 包与 `phase13-03` 的 `repository_id` 锚点口径保持单值一致
  - [x] SubTask 8.2: 校验本 spec 包与 `phase13-04 / 13-05` 的字段与后端边界保持单值一致
  - [x] SubTask 8.3: 校验本 spec 包与 `phase13-06` 的前端非主消费边界保持单值一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, and `Task 7`

# 执行记录

- SubTask 1.1 结论：brief 顶层字段已冻结为 `repository / governance_profile / global_assets / current_phase / products[] / modules[] / decisions[]` 7 项，与 `dev_plan` phase13-07 范围、`architecture_plan` §4.7 两处逐字一致；与 `shared_baseline` §3.6 为语义单值对应（§3.6 中治理画像 / 全局规范资产 / 当前阶段入口与状态三项为中文描述，映射关系清晰无歧义，非逐字逐句）
- SubTask 1.2 结论：`governance_profile` 对应 `phase13-04` 已冻结治理画像字段，`global_assets` 对应 `phase13-05` 已冻结的全局规范资产承接结果，`current_phase` 对应当前阶段入口与状态；brief 各块语义来源单值化
- SubTask 1.3 结论：spec 已显式禁止新增并列第 8 个临时顶层块，避免后续为单个 agent 场景长出专用 schema
- SubTask 2.1 结论：`repository / governance_profile / global_assets / current_phase` 四块的职责已正式冻结，且都回到 PSCO 结构化事实，而不是目录扫描结果、prompt 片段或前端拼装结果
- SubTask 2.2 结论：`products[] / modules[] / decisions[]` 已冻结为同一 `repository_id` 驱动的结构化数组摘要，不再允许被误解释为 singular 当前对象
- SubTask 2.3 结论：brief 的最小子字段矩阵已正式冻结：`repository` 至少包含 `repository_id / name / provider / url`，`global_assets[]` 至少包含 `name / kind / entry_ref / role / structured_summary / markdown_resolvable`，`current_phase` 至少包含 `name / entry_ref / status`，四实体数组也都已冻结最小 summary shape；后续 `.proto`、前端回看与 agent 消费不再需要各自补 item shape
- SubTask 2.4 结论：spec 已显式禁止把 brief 退化为 prompt 模板片段集合、目录扫描结果回放或并列自然语言指导词字段
- SubTask 3.1 结论：`repository_id` 已被冻结为第一版唯一正式锚点，与 `phase13-03` 的项目锚点和 `architecture_plan` §4.7 的解析协议逐条一致
- SubTask 3.2 结论：多 `Module / Decision` 必须返回数组摘要的规则已进入正式 requirement 与 scenario，后续执行者不再需要临场决定“选哪一个当 current”
- SubTask 3.3 结论：即使数组长度为 1，也不得把协议改写为 singular-only 的另一套 schema；单元素数组与多元素数组共享同一协议
- SubTask 4.1 结论：brief 已明确定位为结构化只读事实输入，不承担 IDE 工作区上下文、目录全文或局部源码读取职责
- SubTask 4.2 结论：目录全文、局部源码、临时文件漂移与即时工作区上下文继续由 IDE / agent 现场能力读取，brief 仅提供更快的治理事实定位与摘要输入
- SubTask 4.3 结论：spec 已显式禁止通过 brief 长出第二套“目录快照 / 文件全文 / 文件列表真相源”，与 `phase13-02` 三层边界中 `IDE-accessible context` 的分工保持一致
- SubTask 5.1 结论：后端已被冻结为第一版 brief 的唯一正式生成来源，后续实现时不得由前端或 agent 自行成为 canonical brief 生产者
- SubTask 5.2 结论：前端只允许回看 brief 对应的同源治理事实，不得派生一份“前端理解版 brief schema”，与 `phase13-06`“前端不承担 brief 主消费职责”单值一致
- SubTask 5.3 结论：agent 消费侧必须按同一 formal schema 消费 brief，不得跳过正式 schema 自行拼装第二版协议，从而避免后端 / 前端 / agent 三侧各讲各话
- SubTask 6.1 结论：brief 与现有 `ProjectContextService` 的关系已在 formal spec 中冻结为“同一 `.proto + ConnectRPC` 主线内的受控演进”，而不是并列第二个 repository-scoped 只读聚合 service
- SubTask 6.2 结论：`GetProjectContext / ExportProjectContext` 若继续保留，现已被正式约束为兼容层或单向派生层，并要求在 brief 合同落地时同步声明兼容窗口与退役策略；这不再只是 tasks 备注
- SubTask 6.3 结论：`product -> products[]` 与 `rules / phases / boundaries -> governance_profile / global_assets / current_phase` 的合同映射，现已被 formal spec 要求显式处理，不再留给 `phase13-08 / 10` 实现阶段临场决定
- SubTask 7.1 结论：全文、Git 推进跟踪、模板自动同步状态与自动建议结果均已被冻结为第一版 brief 非目标，与 `phase13-01 / 13-02` 的非目标集合一致
- SubTask 7.2 结论：前端不承担 brief 主消费职责已显式进入 spec requirement，与 `phase13-06` 前端边界对齐
- SubTask 7.3 结论：brief 已明确为结构化 schema，而不是未结构化 prompt 文本；若未来需要附加提示，只能作为 brief 外围消费协议处理
- SubTask 8.1 结论：本 spec 与 `phase13-03` 的 `repository_id` 唯一锚点、Repository 作为项目锚点但不偷换为治理层本体的口径保持单值一致
- SubTask 8.2 结论：本 spec 对 `governance_profile / global_assets / current_phase` 的承接方式与 `phase13-04 / 13-05` 的字段、后端边界和全文禁止入库口径一致；其中 `global_assets[]` 已显式接住 `name / kind / entry_ref / role` 最小身份字段
- SubTask 8.3 结论：本 spec 与 `phase13-06` 的“前端不得承担 brief 主消费职责、agent-only 协议不得做成前端主内容区”保持单值一致；同时与现有 `ProjectContextService` 的收口路径也已进入 formal spec，不再靠执行记录兜底
