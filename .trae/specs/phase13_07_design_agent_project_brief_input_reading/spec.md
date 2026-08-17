# Phase13-07 产出 agent 项目简报输入与读取设计 Spec

## Why

`phase13-03` 已冻结项目级治理层与四实体主线关系，`phase13-04` 已冻结治理画像字段模型，`phase13-05` 已冻结后端承接位与读写边界，`phase13-06` 已冻结前端承接边界；但如果不继续把 agent 在 IDE 场景中到底读取什么、如何组合、如何解析、以及与 IDE 自带目录读取能力如何分工压成单值协议，后续执行者仍会重新猜“给 agent 发什么”“brief 算不算第二套项目上下文”。

本次 `/spec` 的目标，是把 `project brief for agent` 的最小 schema、解析协议、来源边界与消费边界冻结成唯一正式规格，让后端、前端与 agent 消费侧共享同一份 brief schema，而不是各做一版“我理解的项目简报”。

## What Changes

- 冻结 `project brief for agent` 的第一版最小 schema 与分层组合方式
- 冻结 brief 的解析协议、唯一锚点与多实体数组摘要规则
- 冻结 brief 与 IDE 目录读取能力的协作边界
- 冻结 brief 在后端、前端与 agent 三侧的共享口径与禁止事项

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - 后续 `phase13-08` 后端实现、`phase13-09` 前端实现与后续 agent 消费接入
- Affected code:
  - 无直接源代码改动
  - 直接影响后续实现边界：
    - `backend/`
    - `proto/`
    - `frontend/`

## ADDED Requirements

### Requirement: 冻结 agent 项目简报的第一版最小 schema

系统 SHALL 将 `project brief for agent` 的第一版最小 schema 冻结为以下 7 个顶层字段：

1. `repository`
2. `governance_profile`
3. `global_assets`
4. `current_phase`
5. `products[]`
6. `modules[]`
7. `decisions[]`

补充冻结：

- 第一版不得新增与上述 7 个顶层字段并列的第 8 个“临时先塞进去”的正式顶层块
- 第一版所有 agent 项目简报都必须以同一顶层结构返回，不得按调用方各自裁剪出不同 schema
- `governance_profile / global_assets / current_phase` 的语义必须分别对应 `phase13-04 / phase13-05` 已冻结的治理画像字段、全局规范资产承接结果与当前阶段入口状态

#### Scenario: 执行者判断 brief 顶层结构应该是什么

- **WHEN** 执行者设计 `project brief for agent` 的顶层字段
- **THEN** 必须以上述 7 个顶层字段作为第一版正式 schema
- **AND** 不得为了某个 agent 场景额外拼出并列顶层块

### Requirement: 冻结 brief 各顶层块的组合语义

系统 SHALL 将 brief 的各顶层块组合语义冻结如下：

1. `repository`
   - 承接当前项目锚点仓库的结构化摘要
   - 必须显式包含 `repository_id` 作为唯一正式锚点
2. `governance_profile`
   - 承接 `phase13-04` 已冻结的治理画像字段
   - 不得偷换为目录扫描结果、agent 推断结果或前端展示拼装结果
3. `global_assets`
   - 承接 `phase13-05` 已冻结的全局规范资产逐项承接结果
   - 只允许表达 `entry_ref / role / structured_summary / markdown 回源能力` 这类结构化结果
4. `current_phase`
   - 承接当前阶段名称、入口引用与状态
   - 不得退化为“最近对话阶段印象”或自由文本笔记
5. `products[] / modules[] / decisions[]`
   - 承接同一 `repository_id` 驱动的四实体结构化关系摘要
   - 只表达与当前项目锚点相关的结构化数组结果

补充冻结：

- brief 是“PSCO 结构化事实组合结果”，不是 prompt 模板片段集合
- brief 不得混入“给 agent 的额外自然语言指导词”作为并列 canonical 字段
- 若未来需要追加临时提示，只能作为 brief 外围消费协议处理，不得污染第一版 schema

#### Scenario: 执行者判断某项信息应落入 brief 哪一块

- **WHEN** 执行者判断某项 agent 读取信息应放到哪一块
- **THEN** 必须先按 `repository / governance_profile / global_assets / current_phase / products[] / modules[] / decisions[]` 归类
- **AND** 不得把 IDE 目录全文读取结果塞进 `governance_profile` 或 `global_assets`

### Requirement: 冻结 brief 的最小子字段矩阵

系统 SHALL 将 `project brief for agent` 第一版中各顶层块的最小子字段矩阵冻结如下：

1. `repository`
   - `repository_id`
   - `name`
   - `provider`
   - `url`
2. `governance_profile`
   - 直接承接 `phase13-04` 已冻结的 9 类治理画像字段
   - 不得删改 `project_profile_version / track_type / template_source / docs_workflow_layout / canonical_root_files[] / global_constraint_refs[] / current_phase_*` 的正式字段语义
3. `global_assets[]`
   - `name`
   - `kind`
   - `entry_ref`
   - `role`
   - `structured_summary`
   - `markdown_resolvable`
4. `current_phase`
   - `name`
   - `entry_ref`
   - `status`
5. `products[]`
   - `id`
   - `name`
   - `description`
   - `status`
6. `modules[]`
   - `id`
   - `name`
   - `description`
   - `status`
7. `decisions[]`
   - `id`
   - `title`
   - `status`
   - `context`
   - `hit_sources[]`

补充冻结：

- `global_assets[]` 的最小子字段矩阵必须显式接住 `phase13-05` 已冻结的 `name / kind / entry_ref / role`，不得退化为只靠数组顺序或外部约定识别资产身份
- `structured_summary` 在 `global_assets[]` 中属于可空摘要字段：前 5 份需摘要资产必须提供，`README.md / global_skills.md / project_skills.md` 第一版允许为空
- `markdown_resolvable` 只表达“该资产正文是否允许回源”的只读能力状态，不得被误写成全文正文或 markdown 副本字段
- `products[] / modules[] / decisions[]` 的 item shape 在第一版不得继续留给后端、前端或 agent 消费侧临场补充
- 若后续需要新增子字段，只能通过 formal schema 变更进入，不得在单侧消费端私自扩写

#### Scenario: 执行者设计 brief 的嵌套对象或数组项 shape

- **WHEN** 执行者设计 brief 中某个顶层块的嵌套对象或数组项字段
- **THEN** 必须先满足上述最小子字段矩阵
- **AND** 不得删除 `global_assets[]` 中识别资产身份所需的 `name / kind`
- **AND** 不得让四实体数组摘要继续停留在“结构化摘要”但无 item shape 的模糊状态

### Requirement: 冻结 brief 的唯一锚点与解析协议

系统 SHALL 将 `project brief for agent` 的第一版解析协议冻结为：

1. `repository_id` 是唯一正式锚点
2. `products / modules / decisions` 只能从同一 `repository_id` 驱动的 PSCO 结构化关系中解析
3. 若存在多个 `Module / Decision`，必须返回数组摘要，不得伪造单一“当前 module / decision”
4. agent 若需目录全文，继续由 IDE / agent 现场能力读取，不通过 PSCO brief 补做第二套目录扫描

补充冻结：

- 第一版不得引入第二个锚点，如 `module_id`、`decision_id` 或文件路径，来替代 `repository_id` 充当 brief 主定位键
- `products[] / modules[] / decisions[]` 的缺省语义应是“该仓库相关结构化摘要数组”，而不是“当前唯一对象”
- 即使数组长度为 1，也不得把协议改写为 singular-only 的另一套 schema

#### Scenario: 执行者设计 brief 解析逻辑

- **WHEN** 执行者设计 `project brief for agent` 的解析逻辑
- **THEN** 必须以同一 `repository_id` 为唯一正式锚点
- **AND** 必须返回数组摘要而不是伪造单一当前实体
- **AND** 不得让 PSCO 代替 IDE 扫描目录全文

### Requirement: 冻结 brief 与 IDE 目录读取能力的协作边界

系统 SHALL 将 brief 与 IDE / agent 现场目录读取能力的关系冻结为协作关系，而不是替代关系。

其正式边界如下：

1. PSCO brief 负责提供结构化只读事实输入
2. IDE / agent 现场能力继续负责目录全文、局部源码、临时文件漂移与工作区即时上下文读取
3. PSCO brief 不负责复刻 IDE 的项目目录扫描能力
4. 第一版不得通过 brief 长出第二套“目录快照 / 文件全文 / 文件列表真相源”

补充冻结：

- brief 可以引用 `entry_ref` 或 markdown 回源能力，但这不等于 PSCO 接管文件全文读取
- brief 允许帮助 agent 更快定位应读哪些正式治理资产，但不负责替 IDE 完成全文消费
- 若 agent 需要更细粒度实现上下文，应继续回到 IDE 工作区读取

#### Scenario: 执行者判断某项目录信息是否应进入 brief

- **WHEN** 执行者评估某项目录相关信息是否应被放进 brief
- **THEN** 只有结构化治理事实、入口关系与资产摘要可以进入 brief
- **AND** 目录全文、临时文件状态与实现细节必须继续留给 IDE / agent 现场能力

### Requirement: 冻结后端、前端与 agent 消费侧共享同一份 brief schema

系统 SHALL 要求后端、前端与 agent 消费侧共享同一份 `project brief for agent` schema 与字段语义。

补充冻结：

1. 后端是第一版 brief 的唯一正式生成来源
2. 前端可以展示或回看 brief 对应的治理事实，但不得衍生一份“前端理解版 brief schema”
3. agent 消费侧不得跳过正式 schema 自行拼装第二版 brief 协议
4. 第一版任何字段重命名、字段挪位或语义扩写，都必须回到同一份 formal schema 变更，而不是在单侧偷偷演化

#### Scenario: 执行者在不同消费侧定义 brief

- **WHEN** 执行者尝试在后端、前端或 agent 侧各自定义一版 brief 结构
- **THEN** 应判定为违反本 spec 的单值 schema 约束
- **AND** 必须回收到同一份 formal brief schema

### Requirement: 冻结 brief 与现有 ProjectContextService 的正式关系

系统 SHALL 将 `project brief for agent` 与现有 `ProjectContextService` 的关系冻结为：**在同一 `.proto + ConnectRPC` 主线内的受控演进关系**，而不是独立并列的第二套 repository-scoped 只读聚合协议。

其正式约束如下：

1. 现有 `ProjectContextService` 是第一版 brief 合同演进的正式起点
2. 第一版不得新增与 `ProjectContextService` 长期并列存在的第二个 repository-scoped 只读聚合 service
3. 若为承接 brief 新增或改造 RPC / message，必须在同一 `ProjectContextService` 主线内完成，并显式声明兼容策略
4. `GetProjectContext` 若继续保留，只能作为兼容读取层存在，不得与 brief 长期并列演化为两套 canonical agent 输入协议
5. `ExportProjectContext` 若继续保留，只能作为兼容导出层或单向派生层存在，不得与 brief 并列成为第二套正式 agent 输入主线

补充冻结：

- 若选择在现有 `ProjectContextService` 内新增 brief RPC，必须同步声明 `GetProjectContext / ExportProjectContext` 的兼容窗口与后续退役策略
- 若选择直接演进现有读取 / 导出 RPC 到 brief schema，必须显式处理 `product -> products[]`、`rules / phases / boundaries -> governance_profile / global_assets / current_phase` 的合同映射
- 第一版不得把“如何与现有 `ProjectContextService` 收口”留给 `phase13-08 / phase13-10` 实现阶段临场决定
- 该约束的目标是维持 repository-scoped 只读聚合合同的单 canonical 主线，而不是保留两套长期并列读取协议

#### Scenario: 执行者设计 brief 的后端合同落点

- **WHEN** 执行者设计 `project brief for agent` 的后端 `.proto` / RPC 落点
- **THEN** 必须将其落在现有 `ProjectContextService` 主线的受控演进内
- **AND** 必须同步声明现有 `GetProjectContext / ExportProjectContext` 的兼容或退役策略
- **AND** 不得让 brief 落成长期并列的第二套只读聚合 service

### Requirement: 冻结第一版 brief 的非目标与禁止事项

系统 SHALL 将以下事项冻结为第一版 `project brief for agent` 的明确非目标：

1. 通过 brief 提供目录全文、源码全文或第二套目录扫描结果
2. 通过 brief 提供 Git 推进跟踪、模板自动同步状态或自动建议结果
3. 通过 brief 伪造单一“当前 module / decision”
4. 让前端承担 brief 的主消费职责
5. 让 brief 退化成一段未结构化 prompt 文本

#### Scenario: 执行者尝试扩写第一版 brief

- **WHEN** 执行者试图把全文、Git 推进信息、模板自动同步状态或自由文本 prompt 塞进第一版 brief
- **THEN** 应判定为越出当前阶段边界
- **AND** 必须回到第一版结构化只读输入定位

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的 agent brief 设计前提

`phase13_project_governance_profile_foundation` MUST 先完成 `project brief for agent` 的最小 schema、组合语义、唯一锚点、数组摘要协议与 IDE 协作边界冻结，再进入后续后端实现与 agent 消费接入；若这些 brief 边界仍未冻结，则后续任何 agent 入口、消费协议或验收口径都不得视为稳定。

## REMOVED Requirements

### Requirement: 允许不同消费侧各自维护一版“项目简报”

**Reason**: 这种解释会让后端、前端与 agent 长出三套并列上下文协议，重新制造“到底该信哪一版项目简报”的歧义，也会破坏 `repository_id` 作为唯一正式锚点的单值边界。

**Migration**: 将 `project brief for agent` 统一回收到同一份 formal schema；后端负责生成，前端只做同源事实回看，agent 按同一份 schema 消费。
