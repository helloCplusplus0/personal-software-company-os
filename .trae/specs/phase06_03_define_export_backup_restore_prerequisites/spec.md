# Phase06-03 导出、基础备份与恢复前提语义冻结 Spec

## Why

`phase06` 已经把 first-run onboarding 边界与 draft-first 写路径冻结成单值结论，下一步必须把 `Export / Backup / restore prerequisites / backup verified` 的正式语义收住。只有先明确“什么叫带走数据、什么叫保留当前实例、什么叫恢复前提已验证”，后续 `.proto`、后端 owner、Dashboard 动作入口与验收 fixture 才不会把“文件写出成功”误判为数据主权闭合。

## What Changes

- 冻结 `Export` 与 `Backup` 的职责边界，不再允许把二者当作同义词
- 冻结 `Export` 与 `Backup` 的最小覆盖矩阵
- 冻结 `Export` 与 `Backup` 的正式用户入口位与执行路由
- 冻结当前阶段“恢复前提”的最小含义，只承接 restore prerequisites read / verify，不承接真正 restore 写回
- 冻结 `backup verified` 的最小成立动作：产物生成、manifest 可读、覆盖矩阵可核对、schema / 版本前提可校验
- 冻结当前阶段的基础错误语义与失败归类，避免把“生成文件成功”与“可验证恢复前提成立”混为一谈
- 明确当前阶段不引入自动同步、连续备份、多端同步、复杂灾备或依赖第三方平台的前提

## Impact

- Affected specs: `phase06_onboarding_sovereignty_reuse_foundation`
- Affected code:
  - 后续 `frontend/src/routes/dashboard.tsx` 下新增 `/dashboard/export` 与 `/dashboard/backup` 子路由或等价路由文件
  - 后续 `frontend/src/features/dashboard/` 中的动作区入口与回流语义
  - 后续新增 `frontend/src/features/export/`、`frontend/src/features/backup/` 或等价切片
  - 后续 `ExportRead / ExportWrite / BackupWrite` owner
  - 后续 `export_snapshot / backup_snapshot` 读写模型
  - 后续 `.proto` 合同、HTTP DTO、manifest 读取与验收 fixture

## ADDED Requirements

### Requirement: Export 与 Backup 职责边界冻结

系统 SHALL 将 `Export` 与 `Backup` 冻结为不同职责的能力，不允许在当前阶段把二者视为同义词。

#### Scenario: Export 语义判定

- **WHEN** 接手者判断 `Export` 的正式语义
- **THEN** 必须得到“面向用户带走核心资产数据”的单值结论
- **AND** `Export` 的主要目标是支持用户独立保存、理解与迁移当前核心资产
- **AND** 当前阶段不得把 `Export` 解释为“当前实例完整恢复包”

#### Scenario: Backup 语义判定

- **WHEN** 接手者判断 `Backup` 的正式语义
- **THEN** 必须得到“面向当前实例保留与恢复前提校验”的单值结论
- **AND** `Backup` 的主要目标是支持当前实例保留、迁移与后续恢复准备
- **AND** 当前阶段不得把 `Backup` 收窄为“只是另一种导出文件”

### Requirement: Export 最小覆盖矩阵冻结

系统 SHALL 将 `Export` 的最小覆盖范围冻结为核心资产主表与核心绑定关系，且不得缺失关系数据。

#### Scenario: Export 覆盖范围判定

- **WHEN** 接手者定义当前阶段 `Export` 的覆盖内容
- **THEN** `Export` 最小必须覆盖以下数据集：
  - `products`
  - `modules`
  - `releases`
  - `repositories`
  - `decisions`
  - `decision_links`
  - `product_modules`
  - `product_repositories`
  - `module_repositories`
- **AND** 当前阶段不得把“只导出主实体，不导出绑定 / 关联关系”解释为完成 `Export`

### Requirement: Backup 最小覆盖矩阵冻结

系统 SHALL 将 `Backup` 的最小覆盖范围冻结为“不小于 Export”，并额外带出恢复前提所需的基础元信息。

#### Scenario: Backup 覆盖范围判定

- **WHEN** 接手者定义当前阶段 `Backup` 的覆盖内容
- **THEN** `Backup` 的最小覆盖范围不得小于 `Export`
- **AND** `Backup` 除核心资产外，至少还必须带出：
  - 当前备份清单或 `manifest`
  - 备份创建时间
  - 当前实例恢复所需的 `schema / version` 前提
- **AND** 当前阶段不得把仅包含主实体快照、但缺少 `manifest` 与恢复前提的产物判定为合格 `Backup`

### Requirement: Export / Backup 正式入口与执行路由冻结

系统 SHALL 将 `Export / Backup` 的正式用户入口冻结为 `Dashboard` 动作区中的独立入口，并冻结正式执行路由。

#### Scenario: Export 入口位判定

- **WHEN** 用户从当前阶段正式 UI 中进入 `Export`
- **THEN** 正式用户入口必须位于 `Dashboard` 动作区
- **AND** 正式执行路由必须单值冻结为 `/dashboard/export`
- **AND** 当前阶段不得在 `Dashboard Home` 主内容区内联完成整个导出流程

#### Scenario: Backup 入口位判定

- **WHEN** 用户从当前阶段正式 UI 中进入 `Backup`
- **THEN** 正式用户入口必须位于 `Dashboard` 动作区
- **AND** 正式执行路由必须单值冻结为 `/dashboard/backup`
- **AND** 当前阶段不得在 `Dashboard Home` 主内容区内联完成整个备份流程

### Requirement: 恢复前提语义冻结

系统 SHALL 将当前阶段“恢复前提”冻结为 restore prerequisites read / verify，而不是要求当前阶段完成真正 restore 写回。

#### Scenario: 恢复前提判定

- **WHEN** 接手者定义当前阶段的恢复相关能力
- **THEN** 当前阶段必须至少能读取并校验 `Backup` 产物中的 `manifest`、覆盖矩阵与 `schema / version` 前提
- **AND** 当前阶段可以判断“该备份是否具备后续 restore prerequisites”
- **AND** 当前阶段不要求真正执行数据写回式 restore

### Requirement: Backup Verified 最小验证链冻结

系统 SHALL 将 `backup verified` 的最小成立条件冻结为一组可执行动作，而不是单一的文件生成结果。

#### Scenario: Backup Verified 成立

- **WHEN** 用户触发一次基础备份，且系统需要判定当前阶段 `backup verified` 是否成立
- **THEN** 至少必须同时满足以下条件：
  - 已生成可读取的备份产物
  - 可重新读取并解析备份 `manifest`
  - `manifest` 中可见并可核对核心资产覆盖矩阵
  - `manifest` 中可见并可校验 `schema / version` 恢复前提
- **AND** 只有在以上条件同时成立时，当前阶段才允许判定 `backup verified`

#### Scenario: 仅文件写出成功不成立

- **WHEN** 系统只是写出了某个备份文件，但无法重新读取 `manifest`、无法核对覆盖矩阵，或无法校验 `schema / version` 前提
- **THEN** 当前阶段不得把该结果判定为 `backup verified`

### Requirement: Export 成功语义冻结

系统 SHALL 将 `Export` 成功语义冻结为“核心资产与绑定关系已被带出，并形成用户可独立保存的导出结果”。

#### Scenario: Export 成立

- **WHEN** 用户执行 `Export`
- **THEN** 系统必须生成一个可独立保存与带走的导出结果
- **AND** 该结果必须包含当前阶段冻结的核心资产覆盖矩阵
- **AND** 用户必须能够确认导出完成，而不是只看到后台写出动作

### Requirement: 当前阶段错误语义冻结

系统 SHALL 将 `Export / Backup / restore prerequisites verify` 的错误语义冻结为用户可区分的最小失败类型。

#### Scenario: Export 失败归类

- **WHEN** `Export` 执行失败
- **THEN** 至少必须能区分以下失败类型：
  - 核心资产装配失败
  - 导出产物生成失败
  - 导出结果不可读取或不可交付
- **AND** 不得把上述所有错误统一收敛为笼统“导出失败”

#### Scenario: Backup 失败归类

- **WHEN** `Backup` 执行或验证失败
- **THEN** 至少必须能区分以下失败类型：
  - 备份产物生成失败
  - `manifest` 缺失或不可解析
  - 覆盖矩阵缺失或不完整
  - `schema / version` 恢复前提缺失或不可校验
- **AND** 不得把 `backup file written` 与 `backup verified` 混写为同一个成功态

### Requirement: 第三方依赖边界冻结

系统 SHALL 明确当前阶段 `Export / Backup` 不得依赖第三方平台作为唯一完成前提。

#### Scenario: 第三方依赖判定

- **WHEN** 接手者设计当前阶段 `Export / Backup`
- **THEN** 不得要求用户必须依赖 GitHub、云盘、对象存储平台或其他第三方平台才能完成导出或备份
- **AND** 当前系统边界内必须存在独立完成 `Export / Backup` 的路径

### Requirement: 当前阶段非目标冻结

系统 SHALL 明确当前阶段的 `Export / Backup` 不扩写为连续备份、复杂灾备、多端同步或完整 restore 系统。

#### Scenario: 非目标判定

- **WHEN** 接手者扩展当前阶段 `Export / Backup`
- **THEN** 不得把以下能力写成 `phase06-03` 的当前目标：
  - 自动同步
  - 连续备份
  - 多端同步
  - 完整 restore 写回
  - 复杂灾备编排

## MODIFIED Requirements

### Requirement: Local First 的导出 / 备份承诺

`Local First = 数据所有权优先` 在 `phase06-03` 中 SHALL 被操作化为：用户可以独立导出核心资产，且当前实例可以生成具备 restore prerequisites verify 基线的基础备份。

#### Scenario: Local First 语义回落

- **WHEN** 接手者判断当前阶段 `Local First` 在导出 / 备份上的实际承诺
- **THEN** 必须明确导出承诺是“带走核心资产数据”
- **AND** 必须明确备份承诺是“形成可验证恢复前提的当前实例备份”
- **AND** 不得把 `Local First` 继续停留在抽象口号层

## REMOVED Requirements

### Requirement: 只要文件成功写出就可视为恢复前提已验证

**Reason**: 这种表述会把“产物生成成功”和“恢复前提可验证”混为一谈，无法支撑当前阶段对 `manifest`、覆盖矩阵与 `schema / version` 前提的校验要求。

**Migration**: phase06 后续实现统一改为：只有在备份产物可读、`manifest` 可解析、覆盖矩阵可核对、`schema / version` 前提可校验时，才允许把结果判定为 `backup verified`。
