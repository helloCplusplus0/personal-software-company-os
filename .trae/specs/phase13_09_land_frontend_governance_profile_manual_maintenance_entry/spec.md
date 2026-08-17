# Phase13-09 落实前端治理画像承接与手工维护入口 Spec

## Why

`phase13-06` 已冻结前端信息架构与展示边界，`phase13-08` 已落治理画像后端写读主线，但如果不继续把前端实现入口、读写承接位、表单结构与旧 UI 退出规则压成实现规格，后续执行仍可能重新长出第二入口、把 `phase12`“项目上下文”换名保留，或把目录真实路径做成面向普通用户的大块主内容。

本次 `/spec` 的目标，是把 Web 端治理画像正式承接位、人类手工维护入口、只读回看区与旧设计退出规则冻结为可直接实施的前端实现规格。

## What Changes

- 冻结 `Repository detail` 内治理画像区的正式落点、页面层级与读写承接关系
- 冻结治理画像前端读路径、写路径、失效刷新与成功回流主线
- 冻结当前项目范式 v1、canonical 根级文件、全局规范资产与当前阶段信息的前端维护/回看方式
- 冻结 `phase12` 遗留“项目上下文 / 共享项目上下文”前端残留的移除规则

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - `phase13-06`
  - `phase13-08`
  - 后续 `phase13-10` brief 前端回看
- Affected code:
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/project-context/`
  - 新增或回收的治理画像前端 slice（应收敛在单一前端承接位）

## ADDED Requirements

### Requirement: 冻结前端正式承接位与页面落点

系统 SHALL 将项目治理画像第一版前端正式承接位冻结为 `Repository detail` 内的单一治理画像承接区。

补充冻结：

1. 治理画像区只能出现在 `Repository detail`
2. 该承接区必须位于仓库业务主内容之后、页面内 secondary 区域中
3. 第一版不得新增独立“治理画像”页面、tab、drawer 入口或并列详情页区块
4. 第一版不得在 `Product / Module / Decision / Dashboard / Review` 页面新增治理画像主内容区

#### Scenario: 执行者决定治理画像 UI 落在哪个页面

- **WHEN** 执行者设计治理画像前端正式落点
- **THEN** 必须放在 `Repository detail`
- **AND** 不得创建第二入口或跨页并列主承接区

### Requirement: 冻结 Repository detail 内的治理画像区结构

系统 SHALL 将 `Repository detail` 中治理画像区的实现结构冻结为：

1. `概览卡片`
   - 只读展示 `project_profile_version / track_type / docs_workflow_layout / current_phase_*`
2. `手工维护表单区`
   - 承接 `template_source / canonical_root_files[] / global_asset_bindings[]` 的维护
3. `摘要回看区`
   - 承接 canonical 根级文件角色、全局规范资产结构化摘要与轻量 `entry_ref`

补充冻结：

- 这三部分必须属于同一治理画像区，而不是分散为多个并列卡片系统
- `概览卡片` 与 `摘要回看区` 都应保持轻量，避免大块解释型文案
- `entry_ref` 只能作为 secondary metadata 或 locator 呈现
- `structured_summary` 必须优先于真实文件路径成为主阅读内容

#### Scenario: 执行者拆分治理画像区内部结构

- **WHEN** 执行者拆分 `Repository detail` 中治理画像区的内部区块
- **THEN** 必须围绕“概览 / 维护 / 摘要回看”三部分组织
- **AND** 不得把真实路径列表做成主视觉内容

### Requirement: 冻结前端读路径与唯一 query owner

系统 SHALL 将治理画像前端读路径冻结为单一只读主线：

1. 读取锚点是 `repository_id`
2. query owner 只能有一个正式承接位
3. 该 query owner 只承接治理画像读取、缓存键、只读解包与错误状态
4. 页面与展示组件不得各自重复拼装治理画像读取逻辑

补充冻结：

- query key 必须以 `repository_id` 为唯一正式参数
- 第一版应直接消费 `phase13-08` 的治理画像读取主线，而不是继续依赖 `phase12 project-context`
- query 层不得混入写动作、提交逻辑或失效刷新策略

#### Scenario: 执行者实现治理画像读取

- **WHEN** 执行者实现前端治理画像读取
- **THEN** 必须收敛到单一 query owner
- **AND** 不得在 page / card / panel 中重复各写一套读取逻辑

### Requirement: 冻结前端写路径与唯一 mutation owner

系统 SHALL 将治理画像前端写路径冻结为单一正式 mutation 主线：

1. mutation owner 只能有一个正式承接位
2. 该 mutation owner 承接保存请求、错误归一化、成功回流与缓存失效
3. 页面组件、表单组件与展示组件不得各自内联 `useMutation`
4. 保存成功后必须精准刷新当前 `repository_id` 对应的治理画像读取结果

补充冻结：

- 写路径只允许提交 `template_source / canonical_root_files[] / global_asset_bindings[]`
- `track_type / current_phase_* / docs_workflow_layout` 在第一版只读，不得进入可编辑提交负载
- 若保存失败，应保持表单可重试与错误可见，不得静默吞错

#### Scenario: 执行者实现治理画像保存

- **WHEN** 执行者实现治理画像保存入口
- **THEN** 必须通过单一 mutation owner 提交
- **AND** 必须精准刷新当前 `repository_id` 的治理画像读取结果
- **AND** 不得把只读字段混入保存负载

### Requirement: 冻结第一版人类手工维护表单范围

系统 SHALL 将第一版前端手工维护表单范围冻结为：

1. `template_source`
2. `canonical_root_files[]`
   - `file_name`
   - `role`
   - `required`
3. `global_asset_bindings[]`
   - `name`
   - `kind`
   - `entry_ref`
   - `role`
   - `structured_summary`

补充冻结：

- `name / kind` 在资产矩阵中属于受控值，应以前端受控展示承接，不得让用户自由新增第 9 项资产
- 前 5 项需要摘要的资产必须有可维护的 `structured_summary` 输入位
- `README.md / global_skills.md / project_skills.md` 第一版允许不填 `structured_summary`
- 第一版不得提供 markdown 正文编辑器

#### Scenario: 执行者设计治理画像维护表单字段

- **WHEN** 执行者设计治理画像第一版维护表单
- **THEN** 只能覆盖上述允许维护的结构化字段
- **AND** 不得让用户新增矩阵外资产或编辑 markdown 正文

### Requirement: 冻结只读展示与摘要回看范围

系统 SHALL 将以下内容冻结为只读展示或摘要回看内容：

1. `project_profile_version`
2. `track_type`
3. `docs_workflow_layout`
4. `current_phase_name / current_phase_ref / current_phase_status`
5. `backend / database / frontend / proto` 顶层目录矩阵
6. `markdown_resolvable`
7. 全局规范资产 `structured_summary + entry_ref`

补充冻结：

- `markdown_resolvable` 只表达“是否允许正文回源”的能力状态，不等于正文内容
- 顶层目录矩阵只能作为当前项目范式 v1 的轻量只读基线表达
- `current_phase_ref` 与资产 `entry_ref` 可以复制或跳转，但不得长成大块路径说明区

#### Scenario: 执行者设计只读信息展示

- **WHEN** 执行者设计治理画像只读展示
- **THEN** 必须让摘要与状态优先成为主阅读内容
- **AND** 不得把真实路径、文件名列表或目录说明放大成主内容

### Requirement: 冻结 phase12 遗留项目上下文 UI 的退出规则

系统 SHALL 将以下旧设计冻结为必须退出前端的遗留内容：

1. `Repository detail` 中现有“项目上下文”区
2. `Decision detail` 中现有“共享项目上下文入口”卡片
3. 任何继续使用“项目上下文 / 共享项目上下文”命名的前端区块

补充冻结：

- 第一版实现不得通过换标题、并区保留或弱化文案的方式延续旧设计
- 治理画像前端实现只能回到 `phase13` 的治理画像术语与字段边界
- `Decision detail` 应只保留自身业务语义，不再引导用户去看“共享项目上下文”

#### Scenario: 执行者处理旧 project-context 残留 UI

- **WHEN** 执行者处理当前前端中遗留的 `project-context` UI
- **THEN** 必须将其整体移除
- **AND** 不得以换名或轻量保留方式继续存在

### Requirement: 冻结 UI 层级与产品定位一致性

系统 SHALL 保持四实体主线与详情页结构稳定，并确保治理画像区只承担“项目级治理信息的人类维护与摘要回看”职责。

补充冻结：

- `Repository detail` 仍以仓库业务主内容为主，治理画像区属于 secondary governance 区
- 治理画像区不得抢占四实体业务摘要、状态操作与关系工作台的主视觉
- 前端不得把 agent-only 协议、IDE 目录能力或治理规则解释文档做成大块产品说明

#### Scenario: 执行者评估治理画像区是否与产品定位一致

- **WHEN** 执行者评估治理画像区的视觉层级与文案重心
- **THEN** 必须保证它服务于“人类维护 + 摘要回看”
- **AND** 不得压过仓库业务主内容或变成验收层说明面板

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的前端实现前提

`phase13_project_governance_profile_foundation` MUST 在 `phase13-09` 中完成治理画像前端正式承接位、单一读写 owner、手工维护表单范围、只读摘要回看与旧 `project-context` UI 退出规则的实现规格冻结；若这些实现规格仍未冻结，则治理画像前端实现不得视为可执行。

## REMOVED Requirements

### Requirement: 允许前端继续沿用 phase12 的 project-context 叙事来承接治理信息

**Reason**: 这会直接违背 `phase13` 已明确的新基线，让治理画像实现继续背着旧概念前进，并把用户重新带回“这是 agent 上下文还是产品治理信息”的混乱语义中。

**Migration**: 前端实现统一改为 `phase13` 治理画像术语、字段与维护边界；`phase12` 遗留 `project-context` UI 在实现阶段整体移除。
