# Tasks

- [x] Task 1: 产出当前后端模块结构概览。把既有模块装配模式、模块清单与跨模块读取模式写成单值结论。
  - [x] SubTask 1.1: 明确既有四层结构（handler / service / candidate / repository）与 platform 组合根模式
  - [x] SubTask 1.2: 明确既有 5 个模块清单与装配函数
  - [x] SubTask 1.3: 明确"消费方模块拥有 candidate reader"跨模块读取模式

- [x] Task 2: 产出 Export 模块边界设计。把模块定位、结构、接口分组、`export_snapshot` 读模型与职责边界冻结到可直接落地的程度。
  - [x] SubTask 2.1: 明确 Export 模块目录结构（handler / service / candidate / types.go / errors.go / response.go）
  - [x] SubTask 2.2: 明确 ExportRead 接口 owner（query_handler → query_service，GET /api/dashboard/export）
  - [x] SubTask 2.3: 明确 ExportWrite 接口 owner（command_handler → command_service，POST /api/dashboard/export）
  - [x] SubTask 2.4: 明确 ExportRead 正式承接 `export_snapshot`（资产范围 + 创建时间 + 创建结果摘要 / 预览态）
  - [x] SubTask 2.5: 明确 Export 模块只读 canonical 表、不承接备份职责，且当前阶段不提前冻结 export 元数据介质

- [x] Task 3: 产出 Backup 模块边界设计。把模块定位、结构、接口分组与职责边界冻结到可直接落地的程度。
  - [x] SubTask 3.1: 明确 Backup 模块目录结构（handler / service / candidate / repository / types.go / errors.go / response.go）
  - [x] SubTask 3.2: 明确 BackupWrite command 子路径（command_handler → command_service，POST /api/dashboard/backup）
  - [x] SubTask 3.3: 明确 BackupWrite 的 read/verify 子路径（query_handler → query_service，GET /api/dashboard/backup）
  - [x] SubTask 3.4: 明确 `backup_snapshot` 由 BackupWrite owner 内的正式 read/verify 子路径承接，BackupWrite 响应只返回摘要
  - [x] SubTask 3.5: 明确 Backup 模块只读 canonical 表、不承接导出职责、不执行 restore 写回

- [x] Task 4: 产出数据装配责任边界与覆盖矩阵。
  - [x] SubTask 4.1: 明确 Export 覆盖矩阵（9 类 canonical 表）与装配 owner（export.candidate.AssetReader）
  - [x] SubTask 4.2: 明确 Backup 覆盖矩阵（9 类 + manifest + 创建时间 + schema/version）与装配 owner
  - [x] SubTask 4.3: 明确两个 candidate reader 各自独立拥有、分别读取相同覆盖范围
  - [x] SubTask 4.4: 明确装配约束（candidate reader 直接使用 pool、service 不直接写跨模块 SQL、覆盖矩阵必须完整）

- [x] Task 5: 产出 `backup verified` 最小后端校验链与 owner 归属。
  - [x] SubTask 5.1: 明确 5 步校验链定义（产物生成 → snapshot 可读 → manifest 可解析 → 覆盖矩阵可核对 → schema/version 可校验）
  - [x] SubTask 5.2: 明确步骤 1 由 BackupWrite.command 承接、步骤 2-5 由 BackupWrite.read_verify 承接
  - [x] SubTask 5.3: 明确 BackupWrite 响应不包含 backup verified 判定结果
  - [x] SubTask 5.4: 明确校验失败必须分类（manifest 缺失 / 覆盖不完整 / schema 缺失），不得笼统收敛

- [x] Task 6: 产出与既有服务、脚本、数据库的边界关系。
  - [x] SubTask 6.1: 明确与 canonical 模块的边界（只读、不依赖 service、canonical 不感知）
  - [x] SubTask 6.2: 明确与 platform 装配点的边界（4 个装配函数、装配顺序约束）
  - [x] SubTask 6.3: 明确与数据库脚本的边界（不修改 migrations / seeds、验收 fixture 需覆盖）
  - [x] SubTask 6.4: 明确不冻结的细节（脚本名、存储介质、目录细节、文件格式）

- [x] Task 7: 产出 `.proto` 合同分组与接口 owner。
  - [x] SubTask 7.1: 明确 export.proto 文件路径与消息类型（含 `ExportSnapshot` 或等价正式读模型）
  - [x] SubTask 7.2: 明确 backup.proto 文件路径与消息类型（含 BackupManifest / BackupCoverage / SchemaVersion）
  - [x] SubTask 7.3: 明确合同约束（复用 buf 入口、types.go 单向派生、ExportRead 承接 export_snapshot、BackupWrite.read_verify 响应包含完整 snapshot）
  - [x] SubTask 7.4: 明确 backup_snapshot 合同约束（BackupWrite 响应只含摘要、BackupWrite.read_verify 响应含完整 snapshot）

- [x] Task 8: 产出模块间依赖关系图。
  - [x] SubTask 8.1: 明确 Export / Backup 通过 candidate reader 只读 canonical 表
  - [x] SubTask 8.2: 明确 Export / Backup 之间无直接依赖
  - [x] SubTask 8.3: 明确 canonical 模块不感知 Export / Backup

- [x] Task 9: 完成规格一致性校验。验证本次设计与 phase06-03/04/05/shared_baseline 已冻结语义保持一致。
  - [x] SubTask 9.1: 验证 Export 覆盖矩阵（9 类）与 phase06-03 §"Export 最小覆盖矩阵冻结"一致
  - [x] SubTask 9.2: 验证 ExportRead 已正式承接 shared baseline 中的 `export_snapshot` 最小读模型
  - [x] SubTask 9.3: 验证 Backup 覆盖矩阵（不小于 Export + manifest + 创建时间 + schema/version）与 phase06-03 一致
  - [x] SubTask 9.4: 验证 backup verified 校验链与 phase06-03 §"Backup Verified 最小验证链冻结"一致
  - [x] SubTask 9.5: 验证 backup_snapshot 读取侧承接位与 phase06-05 §"backup_snapshot 读取侧一致性"一致，不新长第二套 canonical owner
  - [x] SubTask 9.6: 验证 .proto 合同分组与 phase06-05 §"phase06 新增接口的唯一合同源冻结"一致
  - [x] SubTask 9.7: 验证 Export/Backup 路由（/dashboard/export, /dashboard/backup）与 shared_baseline §4.1 一致
  - [x] SubTask 9.8: 验证跨模块读取模式与既有 dashboard candidate 模式一致
  - [x] SubTask 9.9: 验证错误分类与 phase06-03 §"当前阶段错误语义冻结"一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 3`
- `Task 6` depends on `Task 1` through `Task 5`
- `Task 7` depends on `Task 2` and `Task 3`
- `Task 8` depends on `Task 2` and `Task 3`
- `Task 9` depends on `Task 1` through `Task 8`
