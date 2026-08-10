# Tasks

- [x] Task 1: 对齐 `phase06-13` 的直接上游与现有 proto 主线模式，明确这次任务是“合同主线落地”，不是重新设计一版 `phase06` 合同。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md` 中 `phase06-13` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase06-10` 已冻结的最小 `.proto` 合同设计
  - [x] SubTask 1.3: 对齐 `phase06-12` 正式正文中的合同、HTTP 映射与验收门禁
  - [x] SubTask 1.4: 对齐 `phase03-11 / phase04-11 / phase05-11` 的 proto 主线落地模式与当前 `proto/` workspace 基线

- [x] Task 2: 冻结 `phase06` 合同进入现有 proto workspace 的单一落点与包名语义。
  - [x] SubTask 2.1: 冻结合同源文件落点为 `proto/psco/onboarding/v1/onboarding.proto`、`export/v1/export.proto`、`backup/v1/backup.proto`、`reuse_summary/v1/reuse_summary.proto`
  - [x] SubTask 2.2: 冻结包名与版本语义为 `psco.onboarding.v1`、`psco.export.v1`、`psco.backup.v1`、`psco.reuse_summary.v1`
  - [x] SubTask 2.3: 冻结 `.proto` 作为 `phase06` 当前阶段唯一合同定义入口
  - [x] SubTask 2.4: 冻结 draft-first 继续复用既有 `Product / Repository / Module / Decision` canonical create 合同
  - [x] SubTask 2.5: 冻结四个既有 canonical create request 也必须同步对齐 `phase06` draft-first 最小人工必填与系统默认填充值

- [x] Task 3: 冻结 `phase06` 四个新增 proto 文件必须承接的服务、消息与演进规则。
  - [x] SubTask 3.1: 冻结 `OnboardingService.GetFirstRunState` 与 `first_run_state` 最小消息边界
  - [x] SubTask 3.2: 冻结 `ExportService`、`BackupService`、`ReuseSummaryService` 的最小 RPC 矩阵与消息落地范围
  - [x] SubTask 3.3: 冻结 `GetBackupSnapshot` 的 `read / verify` 读取侧语义与 `GetReuseSummary` 的作用域请求边界
  - [x] SubTask 3.4: 冻结 `backup-manifest-missing / backup-coverage-incomplete / backup-schema-mismatch` 三类失败语义必须继续作为主线合同与映射边界的一部分
  - [x] SubTask 3.5: 冻结字段编号、枚举零值、`reserved` 与 breaking 演进规则必须继续对齐 `phase06-10 / 12`

- [x] Task 4: 冻结 `buf` 工具链、生成产物与 `proto/README.md` 更新要求。
  - [x] SubTask 4.1: 冻结 `buf build / lint / generate / breaking` 必须继续通过既有受控入口运行
  - [x] SubTask 4.2: 冻结 `buf breaking --against '../.git#branch=main,subdir=proto'` 为唯一 breaking 基准路径
  - [x] SubTask 4.3: 冻结 Go 与 TypeScript 生成产物新增落点为四个 `phase06` 包对应目录
  - [x] SubTask 4.4: 冻结 `proto/README.md` 必须把 `phase06` 新增合同纳入目录总览、生成产物表与 RPC → HTTP 映射矩阵

- [x] Task 5: 冻结 `phase06` 过渡传输层与 `.proto` 的单向映射边界。
  - [x] SubTask 5.1: 冻结后端 `types.go`、HTTP handler DTO 与前端 `types.ts / api-adapter.ts` 只允许从 `.proto` 派生或显式对齐
  - [x] SubTask 5.2: 冻结 `GetFirstRunState / GetExportSnapshot / ExportCoreAssets / GetBackupSnapshot / CreateInstanceBackup / GetReuseSummary` 都必须显式组装 Proto request
  - [x] SubTask 5.3: 冻结 `GetReuseSummary` 的 query 参数映射与 `GetBackupSnapshot` 读取侧合同一致性
  - [x] SubTask 5.4: 冻结 DTO / 页面层不得新增 `.proto` 中不存在的业务字段语义

- [x] Task 6: 冻结 `phase06-13` 的最小化边界与阶段推进关系。
  - [x] SubTask 6.1: 冻结当前阶段只承接合同源落地、生成入口、校验链与映射边界
  - [x] SubTask 6.2: 冻结当前阶段继续允许 `chi + JSON HTTP` 作为过渡传输层
  - [x] SubTask 6.3: 冻结当前阶段不要求立即完成 gRPC / Connect 迁移或全量 DTO 替换
  - [x] SubTask 6.4: 冻结后续 `phase06` 实现、联调与验收必须优先引用已落地合同主线

- [x] Task 7: 完成 `phase06-13` 规格一致性校验。
  - [x] SubTask 7.1: 验证本 spec 已把 `phase06` `.proto` 收口为仓库内唯一合同源
  - [x] SubTask 7.2: 验证本 spec 已覆盖 `buf` 工具链入口、breaking 基准路径与生成产物落点
  - [x] SubTask 7.3: 验证本 spec 已明确 `proto/README.md` 的总览更新要求
  - [x] SubTask 7.4: 验证本 spec 已明确 DTO / HTTP 适配层与 `.proto` 的单向映射边界
  - [x] SubTask 7.5: 验证本 spec 已把既有 canonical create 合同同步纳入 `phase06` 主线演进范围
  - [x] SubTask 7.6: 验证本 spec 已把 backup 三类失败语义推进到 proto mainline 与验收承接边界
  - [x] SubTask 7.7: 验证本 spec 与 `phase06-10 / 12` 及既有 `phase03-11 / 04-11 / 05-11` 主线模式一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 3`, `Task 4`
- `Task 6` depends on `Task 3`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 1` through `Task 6`
