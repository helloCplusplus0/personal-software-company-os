# Phase13-06 产出前端信息架构与手工维护设计 Spec

## Why

`phase13-03` 已冻结项目级治理层以 `repository_id` 为唯一项目锚点、以 `Repository detail` 为第一版前端正式承接位，`phase13-04` 已冻结字段模型，`phase13-05` 已冻结后端读写边界。但如果不继续把前端信息架构、入口层级、展示边界与手工维护范围压成单值设计，后续执行者仍可能重新长出独立页面、第二入口，或把 agent / 验收层信息重新做成常驻大块 UI。

本次 `/spec` 的目标，是把项目治理画像在 Web 端的正式承接位、信息层级、展示区块、人工维护边界与 agent 消费边界冻结为可直接进入实现设计的规格。

## What Changes

- 冻结项目治理画像第一版前端正式承接位、入口层级与导航边界
- 冻结 `Repository detail` 内的信息架构与展示分层
- 冻结哪些内容供人类维护、哪些只读、哪些只保留摘要回看
- 冻结目录真实路径、`entry_ref` 与 agent-only 信息在前端的展示约束

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - 后续 `phase13-07` agent brief 设计与 `phase13-09` 前端实现
- Affected code:
  - 无直接源代码改动
  - 直接影响后续实现边界：
    - `frontend/src/features/repository-binding/`
    - `frontend/src/features/project-context/`

## ADDED Requirements

### Requirement: 冻结第一版前端正式承接位与入口层级

系统 SHALL 将项目治理画像第一版前端正式承接位冻结为 `Repository detail`，并以同一 `repository_id` 作为唯一正式前端读取锚点。

补充冻结：

1. 第一版不得新增独立“项目治理画像”一级页面
2. 第一版不得新增与 `Repository detail` 并列的第二入口
3. 第一版不得在 `Dashboard / Daily Review / Weekly Review / Product detail / Module detail / Decision detail` 中再长出并列主承接区
4. 若未来需要提升为独立入口，只能在 `phase13` 正式收口后作为下一阶段进入条件讨论

#### Scenario: 执行者设计治理画像入口

- **WHEN** 执行者设计项目治理画像在 Web 端的正式入口
- **THEN** 必须将其放在 `Repository detail`
- **AND** 必须以同一 `repository_id` 作为唯一正式读取锚点
- **AND** 不得额外建立一级页面或第二入口

### Requirement: 冻结 Repository detail 内的信息架构分层

系统 SHALL 将 `Repository detail` 中的项目治理画像信息架构冻结为以下三层：

1. `治理画像概览层`
   - 用于呈现治理画像版本、技术路线、docs workflow 布局等高层概览
   - 应保持轻量，不得成为大块解释型常驻 UI
2. `结构化维护层`
   - 用于承接第一版允许人类手工维护的结构化字段与全局规范资产承接结果
   - 应以表单 / 面板 / 列表等结构化方式呈现
3. `摘要回看层`
   - 用于回看全局规范资产的角色、摘要、当前阶段入口与状态
   - 应突出“结构化摘要”，而不是正文全文

补充冻结：

- 三层都属于 `Repository detail` 的同一正式承接区，不得拆成多个并列主内容区
- 第一版不得重演 `phase12` 中把解释性或验收性信息做成大块常驻 UI 的问题
- 四实体原有业务详情内容仍保持各自主语义，治理画像区只能作为项目级治理层承接区存在
- `phase12` 遗留的“项目上下文”概念、命名与独立区块应系统性退出前端，不得在 `Repository detail` 或其他详情页中被换名延续
- 当前 `Repository detail` 中 `phase12-09` 遗留的“项目上下文”只读区必须被移除，而不是继续与治理画像承接区并列存在或并区保留
- `rules / phases / boundaries` 这类 `phase12 project-context` 叙事内容，不得作为前端治理画像的默认摘要回看承接对象继续留存；若后续确有必要保留，必须在后续 phase 以 `phase13` 新基线重新定义其人类侧价值，而不是沿用旧概念迁移

#### Scenario: 执行者拆分 Repository detail 页面结构

- **WHEN** 执行者设计 `Repository detail` 的页面结构
- **THEN** 必须围绕“概览层 / 结构化维护层 / 摘要回看层”组织治理画像内容
- **AND** 不得把治理画像拆成多个并列主卡片散落到不同页面
- **AND** 必须移除现有 `phase12`“项目上下文”区，而不是仅把它与治理画像区并列收编

### Requirement: 冻结第一版人类可维护范围

系统 SHALL 将第一版供人类手工维护的前端字段范围冻结为：

1. `template_source`
2. `canonical_root_files[]`
   - `file_name`
   - `role`
   - `required`
3. `global_asset_bindings[]`
   - `entry_ref`
   - `role`
   - `structured_summary`（仅适用于前 5 份需要摘要的资产）

补充冻结：

- 前端手工维护范围必须与 `phase13-05` 写边界一致
- 第一版不得允许前端直接编辑 markdown 正文
- 第一版不得允许前端编辑 `backend / database / frontend / proto` 顶层目录矩阵
- 第一版不得允许前端编辑 `track_type / current_phase_name / current_phase_ref / current_phase_status`

#### Scenario: 执行者设计治理画像维护表单

- **WHEN** 执行者设计项目治理画像的人类维护表单
- **THEN** 只能覆盖上述允许手工维护的结构化字段
- **AND** 不得把正文全文编辑、只读字段编辑或顶层目录矩阵编辑塞进同一表单

### Requirement: 冻结第一版只读展示范围

系统 SHALL 将以下内容冻结为前端只读展示范围：

1. `project_profile_version`
2. `track_type`
3. `docs_workflow_layout`
4. `current_phase_name`
5. `current_phase_ref`
6. `current_phase_status`
7. `backend / database / frontend / proto` 顶层目录矩阵

补充冻结：

- 这些内容只能展示、复制、回看，不得进入前端可编辑表单
- `current_phase_ref` 可以作为定位入口展示，但其定位能力不得被放大为普通用户主内容
- 顶层目录矩阵只用于表达“当前项目范式 v1 的基线输入”，不应被呈现为需要用户频繁操作的 UI
- `backend / database / frontend / proto` 顶层目录矩阵的第一版前端来源冻结为：当前项目范式 v1 的前端只读基线表达
- 该只读基线表达不得反推为新增后端治理字段、不得要求新增 repository-scoped 读模型字段、也不得来自运行时目录扫描

#### Scenario: 执行者判断某项内容应只读还是可编辑

- **WHEN** 执行者判断治理画像中的某项内容是否可编辑
- **THEN** 上述 7 类内容必须保持只读
- **AND** 不得因其出现在治理画像字段中就自动进入维护表单

### Requirement: 冻结顶层目录矩阵的前端承接方式

系统 SHALL 将 `backend / database / frontend / proto` 顶层目录矩阵的第一版前端承接方式冻结为：

1. 仅作为当前项目范式 v1 的只读基线表达存在于前端
2. 不要求后端为其新增 repository-scoped 持久化字段或专属读模型字段
3. 不允许通过运行时目录扫描、IDE 现场读取或 agent 输入协议临时回填到普通产品 UI
4. 若未来需要更动态的目录矩阵来源，只能作为后续 phase 的显式进入项讨论

补充冻结：

- 该矩阵的前端承接方式必须与 `phase13-05` “只作为只读基线输入保留、不新增专属持久化字段”的后端边界保持单值一致
- 第一版前端若展示该矩阵，应将其视为“当前项目范式 v1 的只读说明性基线”，而不是可写治理字段
- 第一版不得因为普通产品 UI 需要展示该矩阵，就反向推动后端补第 10 个治理画像核心字段

#### Scenario: 执行者设计顶层目录矩阵的前端来源

- **WHEN** 执行者设计 `backend / database / frontend / proto` 顶层目录矩阵在前端的展示来源
- **THEN** 必须将其实现为当前项目范式 v1 的前端只读基线表达
- **AND** 不得通过新增后端字段、运行时目录扫描或 agent brief 解析结果来充当前端正式数据源
- **AND** 不得把该矩阵误设计成可编辑治理字段

### Requirement: 冻结人类维护、agent 消费与摘要回看的边界

系统 SHALL 将前端消费边界冻结为：

1. `供人类维护`
   - 仅限结构化字段维护与资产承接结果维护
2. `供 agent 消费`
   - 完整的 `project brief for agent`、数组摘要解析协议、目录全文读取能力边界等，属于 agent 消费主线
   - 不需要在前端做成并列主内容区
3. `只保留摘要回看`
   - 全局规范资产的 `role / structured_summary / entry_ref`
   - 当前阶段入口与状态的结构化回看

补充冻结：

- 前端可以提供人类可理解的摘要回看，以帮助校对 agent 读取结果
- 前端不得把 agent-only 协议、解析规则或 IDE 目录读取能力说明做成产品主界面常驻大块说明
- 前端不得承担 agent brief 的主消费职责
- 前端摘要回看层不得继续承接 `phase12`“项目上下文”叙事本身；任何保留内容都必须以 `phase13` 治理画像字段与全局规范资产矩阵为唯一解释来源

#### Scenario: 执行者判断某项信息是否应该在前端详细展开

- **WHEN** 执行者评估某项治理信息是给人维护、给 agent 消费，还是只做摘要回看
- **THEN** 必须先按上述三类边界分类
- **AND** 若该信息主要服务于 agent 协议或 IDE 协作，则不得在前端展开成主内容区

### Requirement: 冻结目录真实路径与 entry_ref 的前端展示约束

系统 SHALL 冻结：

1. 目录真实路径不得被当作面向普通用户的主内容
2. `entry_ref` 的正式角色是定位与回源入口，不是正文主体
3. 前端若展示 `entry_ref`，应以轻量 locator / secondary metadata 方式呈现
4. 第一版不得把大量真实文件路径做成占据主视觉的列表或说明区

补充冻结：

- 第一版允许用户在需要时查看、复制或跳转 `entry_ref`
- 第一版不允许把路径列表做成“项目治理画像”的主价值展示
- `structured_summary` 应优先于真实路径成为主阅读内容

#### Scenario: 执行者设计 entry_ref 的展示方式

- **WHEN** 执行者设计 `entry_ref` 或目录真实路径的前端展示
- **THEN** 必须将其作为轻量定位元数据处理
- **AND** 必须让结构化摘要优先成为主阅读内容
- **AND** 不得让路径本身成为大块主内容

### Requirement: 冻结 phase12 遗留项目上下文设计的退出规则

系统 SHALL 将 `phase12` 遗留“项目上下文 / 共享项目上下文 / project context”设计的前端退出规则冻结为：

1. `Repository detail` 中现有的“项目上下文”区必须移除
2. `Decision detail` 中现有的“共享项目上下文入口”卡片必须移除
3. 第一版不得在任何页面保留以“项目上下文 / 共享项目上下文”为标题、导语或主交互意图的前端区块
4. 第一版前端不得再把“引导用户去看项目上下文”作为 `Decision detail`、`Module detail` 或其他页面的正式交互目标

补充冻结：

- 该退出规则的目标，是让前端彻底看齐 `phase13` 新基线，而不是把 `phase12` 概念换名后继续保留
- 前端正式承接对象只能回到 `phase13` 治理画像的人类维护边界、只读字段与全局规范资产摘要回看
- 若未来确有新的跨页治理入口需求，只能以 `phase13` 新术语、新职责与新价值重新定义，不得复用 `phase12 project-context` 叙事

#### Scenario: 执行者判断旧项目上下文设计应如何处理

- **WHEN** 执行者处理前端中遗留的“项目上下文 / 共享项目上下文”设计
- **THEN** 必须将其视为需要系统性退出的旧设计
- **AND** 不得通过并区、换标题或轻量保留的方式延续旧概念
- **AND** 必须回到 `phase13` 治理画像的正式承接边界重新组织页面

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的前端设计前提

`phase13_project_governance_profile_foundation` MUST 先完成项目治理画像在 `Repository detail` 的唯一正式承接位、信息架构分层、可编辑范围、只读范围与前端展示边界冻结，再进入 `phase13-07` agent brief 设计与后续前端实现；若这些前端边界仍未冻结，则后续任何页面结构、表单实现或 UI 入口都不得视为稳定。

## REMOVED Requirements

### Requirement: 允许第一版在多个页面同时铺开项目治理画像主内容

**Reason**: 这种解释会直接重演 `phase12` 中“把解释层 / 验收层信息铺成多个大块 UI”的问题，也会让执行者再次陷入“到底挂在 Repository detail 还是另起页面”的反复判断。

**Migration**: 第一版将项目治理画像前端正式承接位统一回收到 `Repository detail`；其他页面仅保留各自业务主语义，不承担治理画像主内容区职责。

### Requirement: 允许 phase12“项目上下文”设计以并区或换名方式继续存在

**Reason**: 这种解释会让前端继续背着 `phase12` 的旧概念前进，无法真正看齐 `phase13` 新基线，也会让执行阶段继续模糊“到底是在做治理画像，还是在保留 project-context UI”。

**Migration**: 将 `phase12` 遗留“项目上下文 / 共享项目上下文”设计从前端系统性移除；第一版前端只保留 `phase13` 治理画像正式需要的人类维护、只读字段与全局规范资产摘要回看。
