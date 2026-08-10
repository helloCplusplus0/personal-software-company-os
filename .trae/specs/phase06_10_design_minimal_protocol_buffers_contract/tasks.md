# Tasks

- [x] Task 1: 对齐 `phase06-10` 的直接上游与边界，明确这次任务是“最小 Proto 合同设计冻结”，不是合同主线落地。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md` 中 `phase06-10` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase06-05` 关于 `.proto` 唯一合同源与 `chi + HTTP JSON` 传输适配职责的冻结口径
  - [x] SubTask 1.3: 对齐 `phase06-01 / 02 / 03 / 04 / 08 / 09` 中 `first_run_state`、draft-first、`export_snapshot`、`backup_snapshot`、`module_reuse_summary / capability_summary` 的既有语义
  - [x] SubTask 1.4: 对齐 `phase03-11 / 04-11 / 05-11` 已建立的 Proto 设计与主线落地分层模式

- [x] Task 2: 冻结 `phase06` 四组最小 Proto 合同的文件分组、包名与版本语义。
  - [x] SubTask 2.1: 冻结 `onboarding / export / backup / reuse_summary` 四个逻辑 Proto 文件组
  - [x] SubTask 2.2: 冻结 `psco.onboarding.v1 / psco.export.v1 / psco.backup.v1 / psco.reuse_summary.v1` 四个 package
  - [x] SubTask 2.3: 冻结当前阶段仅做设计落点，不要求真实文件已写入仓库
  - [x] SubTask 2.4: 冻结 breaking 变更以后续主版本演进承接

- [x] Task 3: 冻结 Onboarding 的最小 Proto 合同设计。
  - [x] SubTask 3.1: 冻结单一 `OnboardingService`
  - [x] SubTask 3.2: 冻结 `OnboardingService` 当前阶段只承接 `GetFirstRunState`
  - [x] SubTask 3.3: 冻结 `first_run_state` 的最小消息边界与状态枚举
  - [x] SubTask 3.4: 冻结 draft-first 对既有 `CreateProduct / CreateRepository / CreateModule / CreateDecision` request / response 的最小影响边界

- [x] Task 4: 冻结 Export 与 Backup 的最小 Proto 合同设计。
  - [x] SubTask 4.1: 冻结单一 `ExportService` 与 `GetExportSnapshot / ExportCoreAssets`
  - [x] SubTask 4.2: 冻结 `ExportSnapshot` 的最小字段边界，不提前冻结文件格式与介质
  - [x] SubTask 4.3: 冻结单一 `BackupService` 与 `GetBackupSnapshot / CreateInstanceBackup`
  - [x] SubTask 4.4: 冻结 `BackupSnapshot` 的 manifest / coverage / schema/version / verified 状态边界

- [x] Task 5: 冻结 Reuse Summary 的最小 Proto 合同设计。
  - [x] SubTask 5.1: 冻结单一 `ReuseSummaryService`
  - [x] SubTask 5.2: 冻结 `GetReuseSummaryRequest` 的 `scope / module_id / product_id` 作用域设计
  - [x] SubTask 5.3: 冻结 `GetReuseSummaryResponse` 必须同时承接两类 summary
  - [x] SubTask 5.4: 冻结 `capability_summary` 的成功空态承接方式

- [x] Task 6: 冻结 `.proto -> HTTP DTO / JSON` 的最小映射策略与错误语义。
  - [x] SubTask 6.1: 冻结 `GetFirstRunState` 与四类既有 canonical create HTTP 入口的 Proto request / response 映射关系
  - [x] SubTask 6.2: 冻结 Export / Backup / Reuse Summary 的最小 HTTP 入口分组
  - [x] SubTask 6.3: 冻结 DTO / handler / adapter 只允许从 `.proto` 单向派生或显式对齐
  - [x] SubTask 6.4: 冻结四组能力的最小错误语义，不混入成功 response 字段

- [x] Task 7: 冻结枚举零值、保留字段与 breaking 演进规则。
  - [x] SubTask 7.1: 冻结枚举首值必须为 `*_UNSPECIFIED = 0`
  - [x] SubTask 7.2: 冻结删除字段或枚举值后必须 `reserved`
  - [x] SubTask 7.3: 冻结不得复用 tag、不得随意改字段类型或 repeated/scalar 语义
  - [x] SubTask 7.4: 冻结不得新增 `required` 字段

- [x] Task 8: 冻结 `phase06-10` 与 `phase06-13` 的边界并完成规格一致性校验。
  - [x] SubTask 8.1: 明确当前 `/spec` 轮次不要求真实 proto 文件、生成产物与源码实现同轮全部完成
  - [x] SubTask 8.2: 明确 `phase06-13` 承接 buf 工具链与生成产物收口，但真实 `.proto` 进入主线时点不得晚于 `phase06-12`
  - [x] SubTask 8.3: 验证本 spec 与 `phase06-05 / 08 / 09` 保持单值一致
  - [x] SubTask 8.4: 验证本 spec 足以支撑 `phase06-12` 正式规格正文与 `phase06-13` 合同主线落地

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`
- `Task 6` depends on `Task 3`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`
- `Task 8` depends on `Task 1` through `Task 7`
