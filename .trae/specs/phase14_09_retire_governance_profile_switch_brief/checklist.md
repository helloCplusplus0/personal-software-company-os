# phase14-09 Checklist

## 0011 迁移（补全与执行）

- [x] 第二段数据迁移 DO 块与 phase14-06 §ADDED-1 算法逐字一致（块首 to_regclass 守卫 / TEMP TABLE ON COMMIT DROP / 源 1 源 2 展开 / GROUP BY path 合并源 2 优先 / 固定 6 轮组树非递归 / 产物幂等写入）——独立复核 R1 PASS
- [x] 迁移产物固定值逐字（name / status=active / description / revision 含 N/M 与源 repository / binding adopts）——独立复核 R1 逐字核对（含全角括号）
- [x] 第三段 drop 两表；主表保留未触碰——git diff 实证第一段建表零改动
- [x] 验收 DB 执行：一条合并 Standard + 树含源节点 + revision 留痕 + binding 正确——**注记：验收库基线为两源表 0 行（无画像数据冻结场景，设计内分支，不产生产物）**；上述产物断言在一次性验证库（样本 5+8 行）全量实证（migration_evidence.md §3）
- [x] 重放一次无报错无重复产物（幂等实证）——验证库与验收库双实证（migration_evidence.md §3.4 / §4.2）
- [x] 迁移前基线导出 + 迁移后零丢失对照（5 字段映射集逐项回溯）——migration_evidence.md §1 / §5

## proto 退役（T1）

- [x] `proto/psco/governance_profile/v1/` 包已删除
- [x] `buf.yaml` breaking.ignore 单文件路径 + 三行注释留痕；无目录前缀扩大化——另含执行期缺口修正 `ignore_only: FIELD_WIRE_JSON_COMPATIBLE_TYPE`（规则级单文件，spec 修正留痕；独立复核 R3 认定最小必要手段）
- [x] `make gen` 后三端画像产物消失（make clean && make gen；find 断言零输出）
- [x] `make lint && make build && make breaking` 全绿（独立复核亲测复验）

## 后端收缩（T3/T4）

- [x] 删 4 文件（connect/server.go / connect/server_test.go / command_service.go / candidate/repository_reader.go）
- [x] errors.go 删 3 哨兵保留 2；types.go 增 GovernanceProfileCoreReadResult；profile_store 删 SaveProfile 与两表读取、ReadProfile 只读主表；query_service 增 ReadProfileCore
- [x] integration test 收缩为只读主表断言（3 用例真实连库 PASS）
- [x] 联动闭合：server.go 单值接收 + 删 mount 调用；router.go 删 mount 函数 + build 收缩；connecterrors 删 3 行；candidate 接口签名 ReadProfileCore（与 phase14-06 冻结 Go 代码逐字一致含注释——独立复核 R2）
- [x] `grep -r "CommandService\|SaveProfile\|connect/" backend/internal/governanceprofile` 零命中（独立复核亲测 exit=1）
- [x] `grep -rn "GovernanceProfileService\|governance_profile.v1" backend/internal/platform` 零命中
- [x] `go build / vet / test ./...` 全绿（集成测试真实连库，projectcontext 8 子场景 PASS；独立复核复验）

## brief 切换（T6）

- [x] project_context.proto：BriefGovernanceProfile 内联逐字 phase14-02 草案；global_assets reserved 3（另补 reserved "global_assets"——对冻结口径纯增强，buf FIELD_NO_DELETE_UNLESS_NAME_RESERVED 要求）；BriefPhaseStatus；无 governance_profile import
- [x] 显式不迁移字段零复活（7 项清单）——GlobalAssetBinding 等旧类型零残留（仅 4 处"已移除"说明注释）
- [x] projectcontext service/types/connect 装配切换 ReadProfileCore + 删 GlobalAssetBinding；standards[] 装配不回退
- [x] brief 集成测试：BriefGovernanceProfile 三字段 + standards[] round-trip 断言通过
- [x] brief 字段面 = phase14-02 对照表"后"列（8 顶层块）——独立复核逐块核对

## 前端移除（T5）

- [x] `frontend/src/features/governance-profile/` 目录不存在
- [x] `grep -r "governance-profile" frontend/src` 零命中
- [x] `frontend/src/gen/proto/psco/governance_profile/` 不存在
- [x] `tsc --noEmit` 零错误——**实际正式入口为 `npx tsc -b`（project references 模式，裸 tsc --noEmit 会假通过）**，exit 0

## 存储（T2）

- [x] 两张 bindings 表 to_regclass 均 NULL（独立复核亲查验收库）
- [x] governance_profiles 主表行数与迁移前一致（=1，repository_id/updated_at 逐字一致）

## 画像 RPC（T4 运行时）

- [x] `/api/psco.governance_profile.v1/...` Connect 请求返回 404 实测（对照 Standard ListStandards=200 / healthz=200；服务进程已清理）

## DoD 门禁

- [x] 六触点断言全绿（T1-T6 逐条留痕；T6 注记：governance_profile 命中为注释 + 内联字段声明本身，属 phase14-02 设计产物）
- [x] brief 对照验收：旧清单信息经 standards[] 零丢失；ref 承接 entry_ref / external_url——migration_evidence.md §5 零丢失对照表（一次性验证库证据；验收库无数据为设计内分支）
- [x] 全库仅一条合并全局 Standard 且首条 revision 记录合并来源（裁决⑧）——一次性验证库实证（固定名恰 1 条 / revision 含 N/M 与源 repository / adopts 绑定）

## 复核与收口

- [x] 子代理独立复核通过（发现项已修复回填）——结论 PASS、0 BLOCKER、3 OBS 留档（OBS-A 多前导斜杠防御性规范化 / OBS-B 两条低风险分支无样本实证 / OBS-C 上游 T6 断言措辞瑕疵，均不修）
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
