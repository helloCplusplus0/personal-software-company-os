# Tasks

- [x] Task 1: proto 画像包删除 + project_context.proto 内联切换 + buf 豁免
  - [x] SubTask 1.1: 删除 `proto/psco/governance_profile/v1/` 包；`project_context.proto` 按 phase14-02 草案切换（BriefGovernanceProfile 内联 / global_assets reserved 3 / BriefPhaseStatus / 删 import）
  - [x] SubTask 1.2: `buf.yaml` breaking.ignore 单文件豁免 + 注释留痕；`make gen` 重生成（三端画像产物消失）+ `make lint && make build && make breaking` 全绿

> Task 1 执行记录：画像包（单文件）已删除；project_context.proto 已切换（内联三消息 + reserved 3 + reserved "global_assets" + status 改 BriefPhaseStatus + 删 import）；buf.yaml 落地 ignore + ignore_only（执行期缺口修正：FIELD_WIRE_JSON_COMPATIBLE_TYPE 规则级豁免，见 spec 修正留痕）；make clean && make gen 后 find 断言零输出；make lint/build/breaking 全绿（2026-08-18 实证）。
- [x] Task 2: 后端 governanceprofile 模块收缩
  - [x] SubTask 2.1: 删 4 文件（connect/server.go / connect/server_test.go / service/command_service.go / candidate/repository_reader.go）
  - [x] SubTask 2.2: 收缩 4 文件（errors.go 删 3 哨兵 / types.go 增 GovernanceProfileCoreReadResult / profile_store.go 删写与两表读取 / query_service.go 增 ReadProfileCore）+ integration test 只读化
  - [x] SubTask 2.3: 联动收缩（platform/server.go 单值接收 + 删 mount 调用；router.go 删 mount 函数 + buildGovernanceProfile 单返回；connecterrors 删 3 行；projectcontext candidate 接口签名收缩）
- [x] Task 3: projectcontext brief 装配切换
  - [x] SubTask 3.1: service/query_service.go 画像装配改 ReadProfileCore + 删 GlobalAssetBinding 装配；types.go domain 收缩；connect/server.go 响应组装切换
  - [x] SubTask 3.2: connect/server_integration_test.go 切换（内联类型断言 + BriefGovernanceProfile 三字段 + standards[] 保持）+ `go build / vet / test ./...` 全绿
- [x] Task 4: 前端切片移除
  - [x] SubTask 4.1: 删 `frontend/src/features/governance-profile/` 整目录；验证 gen 产物消失；grep 零命中 + `tsc --noEmit` 零错误

> Task 2/3 执行记录：删 4 文件 + 2 空目录；模块收缩为 types/errors + query_service(ReadProfileCore) + profile_store(只读主表) + 只读集成测试；联动 3 文件闭合；projectcontext 5 文件切换（接口 ReadProfileCore / 内联组装 BriefGovernanceProfile + BriefPhaseStatus 枚举映射 / 删 GlobalAssetBinding）。清单外联动：无（仅 2 处注释措辞对齐 grep 门禁）。门禁：go build/vet/test 全绿（集成测试真实连库，projectcontext 8 子场景 PASS）；三条 grep 断言零命中。
> Task 4 执行记录：切片 10 文件整目录删除；三条残留断言全过；`npx tsc -b` 零错误（正式入口为 project references 模式，裸 tsc --noEmit 会假通过）；oxlint 无画像相关失败。
- [x] Task 5: 0011 迁移补全与执行
  - [x] SubTask 5.1: 补第二段数据迁移 DO 块（两段式算法 + 幂等守卫 + 产物固定值）与第三段 drop；迁移前导出两表全量基线（零丢失对照依据）
  - [x] SubTask 5.2: 验收 DB 手工 psql -f 执行 + 重放一次验证幂等；迁移专属断言（一条合并 Standard / revision 留痕 / 零丢失对照）
- [x] Task 6: 六触点断言矩阵全量验证
  - [x] SubTask 6.1: T1-T6 机械检查命令逐条执行留痕；画像 RPC 404 实测；DoD 三项（六触点全绿 / brief 对照 / 裁决⑧单条）

> Task 5 执行记录：验收库基线=源1 0 行/源2 0 行/主表 1 行——命中冻结"无画像数据场景"（不产生产物 + 两表 drop，设计内分支）；另建一次性验证库（样本 5+8 行覆盖全算法分支含同名合并/三态拆树/幂等重放）端到端证明算法与裁决⑧产物，验证后已删除；证据落档 `migration_evidence.md`（基线导出/断言/重放/零丢失对照）。0011 追加第二段 DO 块（+233 行）与第三段 drop，第一段未动。
> Task 6 执行记录：T1-T6 机械断言全绿；画像 RPC 404 实测通过（对照 Standard ListStandards=200 / healthz=200，排除整体路由故障；服务进程已清理）；DoD 三项成立（brief 对照与裁决⑧以一次性验证库证据为准，验收库无数据场景为设计内分支）。T6 注记：`governance_profile` 在 project_context.proto 的命中为注释 + 内联字段声明本身（phase14-02 设计产物，无旧包 import/类型引用）。
- [x] Task 7: 独立复核与收口
  - [x] SubTask 7.1: 子代理独立复核（0011 与冻结算法逐字一致性 / 收缩清单完整性 / brief 内联消息逐字 / 六触点断言证据 / 范围外改动）
  - [x] SubTask 7.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交

> Task 7 执行记录：独立复核结论 PASS（R1-R5 全过：0011 与冻结算法逐字一致 / 收缩清单完整闭合含接口冻结代码逐字 / brief 内联消息逐字且 ignore_only 认定为最小必要手段 / 六触点证据亲自重跑全真实 / 31 文件改动+spec 四件套无范围外改动）。0 BLOCKER；3 OBS 留档不修：OBS-A（0011 多前导斜杠防御性规范化，对合法输入行为一致）、OBS-B（主表零行守卫与源1-directory 同 path 合并两条低风险分支无样本实证，实现取向符合冻结总则）、OBS-C（phase14-06 T6 断言措辞与 phase14-02 保留字段名的固有重叠，上游措辞瑕疵）。

# Task Dependencies

- Task 2, Task 3 depend on Task 1（生成产物消失后收缩才可编译闭合）
- Task 4 depends on Task 1（gen 产物先清）
- Task 5 depends on Task 2（代码不再读写两表后方可 drop）
- Task 6 depends on Task 1-5
- Task 7 depends on Task 6
