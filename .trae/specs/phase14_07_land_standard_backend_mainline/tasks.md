# Tasks

- [x] Task 1: proto 合同落地与三端生成
  - [x] SubTask 1.1: 新增 `proto/psco/standard/v1/standard.proto`（4 枚举 + 4 核心消息 + 8 RPC envelope + StandardService，逐字 phase14-04 §ADDED-1/2；文件头四段注释对齐 governance_profile.proto）
  - [x] SubTask 1.2: 修改 `proto/psco/project_context/v1/project_context.proto`（import standard.proto + `standards = 8` + brief 顶层注释随演进更新；不动 global_assets=3 / governance_profile=2）
  - [x] SubTask 1.3: `make gen` 三端生成 + `make lint && make build && make breaking` 三门禁验证
- [x] Task 2: 0011 迁移第一段建表落地
  - [x] SubTask 2.1: 新增 `database/migrations/0011_phase14_standard_entity.sql`（三表 + 唯一约束 + 索引，CREATE TABLE/INDEX IF NOT EXISTS 幂等形态；头注释声明 phase14-09 追加段预留；不写数据迁移、不 drop 画像两表）
    - 执行记录：三门禁全绿（gen/lint/build/breaking exit 0）；Go pb / Go connect（standardv1connect 子目录，与其余 13 服务一致）/ TS 三端产物齐备；project_context 生成产物含 Standards 字段（Go L950 / TS L544）。
- [x] Task 3: standard Go 模块 8 文件实现
  - [x] SubTask 3.1: 支撑文件 `errors.go`（5 哨兵）+ `types.go`（DirectoryTreeNode 逐字 phase14-03 §ADDED-2 + 读写模型 + 受控枚举转换）
  - [x] SubTask 3.2: `validate.go`（R1-R8 逐条 phase14-03 §ADDED-3：稳定错误码 + 节点路径定位 + R2 三时机 + R6 65536 字节）
  - [x] SubTask 3.3: `repository/standard_store.go`（三表 SQL + jsonb 编解码 + UpdateStandard 单事务边界）
  - [x] SubTask 3.4: `candidate/target_reader.go`（4 实体表 EXISTS 校验）
  - [x] SubTask 3.5: `service/command_service.go`（Create/Update/Delete/Bind/Unbind；BindStandard 五步校验链固定顺序；CreateStandard 不记 revision）
  - [x] SubTask 3.6: `service/query_service.go`（List/Get/ListRevisions + StandardReader 实现 ListStandardsByRepository）
  - [x] SubTask 3.7: `connect/server.go`（8 RPC handler + connecterrors 统一错误映射）
    - 执行记录：8 文件 + validate_test.go 落地；`go build / vet / test ./internal/standard/...` 全绿（34 子用例）。
- [x] Task 4: validate 单元测试
  - [x] SubTask 4.1: `validate_test.go` R1-R8 逐条非法用例（8 错误码全覆盖 + 节点路径断言）+ 合法边界用例 + 序列化等价规则用例 + 合法完整树样例
- [x] Task 5: standard 模块集成测试
  - [x] SubTask 5.1: UpdateStandard 事务 round-trip（整树替换 + updated_at + revision 追加）
  - [x] SubTask 5.2: StandardReader 反查（adopts/template_source 绑定后含全树；无绑定空列表）+ 写路径错误语义抽查（name 重复 / RETIRED 拒绝 / active 空树 / Delete active 拒绝 / 八格矩阵非法 / target 不存在）
    - 执行记录：`service/service_integration_test.go` 5 个集成测试在验收 DB（127.0.0.1:55432/psco_development）真实执行全 PASS，非 SKIP；fixture 清理后 4 表残留为 0。测试环境缺三表时已用既有幂等迁移 0011 apply（非代码变更）。
- [x] Task 6: projectcontext 接入与 brief 装配
  - [x] SubTask 6.1: `candidate/context_readers.go` 追加 StandardReader 接口（签名逐字 phase14-04 §ADDED-5）
  - [x] SubTask 6.2: `service/query_service.go` 注入 standardReader + GetProjectBrief 追加 standards[] 装配步骤；`types.go` ProjectBriefReadResult 追加 Standards 字段
  - [x] SubTask 6.3: `connect/server.go` brief 响应组装 standards（domain → proto 递归树转换，复用 standard/connect.DomainStandardToProto）
  - [x] SubTask 6.4: brief 集成测试追加 standards[] round-trip 断言（绑定后含全树 / 未绑定空数组；global_assets 过渡态行为不回退）
    - 执行记录：projectcontext/connect 集成测试真实执行 21.2s 全 PASS，含 adopts role 绑定后三层树 round-trip 断言与既有 empty-arrays 用例扩展。
- [x] Task 7: 装配与错误映射闭合
  - [x] SubTask 7.1: `connecterrors/connect_errors.go` 注册 standard 5 哨兵（NotFound 2 / InvalidArgument 1 / Internal 显式留痕 2）
  - [x] SubTask 7.2: `platform/router.go` buildStandard + mountStandardConnect + buildProjectContext 双 Reader 签名扩展；`platform/server.go` 装配调用点
- [x] Task 8: DoD 全量验证
  - [x] SubTask 8.1: proto 三门禁 + `go build / vet / test ./...` 全绿；集成测试在验收 DB 可用时执行留痕
    - 执行记录：`proto/` make lint + build + breaking 全绿（breaking 对 .git#branch=main 通过）；`backend/` go build ./... + go vet ./... + go test ./... -count=1 全绿（standard 0.011s + standard/service 0.192s + projectcontext/connect 21.239s 集成真实执行）。
- [x] Task 9: 独立复核与收口
  - [x] SubTask 9.1: 子代理独立复核（合同逐字一致性 / DDL 逐字一致性 / 分层约束 / R1-R8 覆盖 / brief 过渡态不回退 / 装配闭合 / DoD 证据）
    - 执行记录：独立复核 11 项全 PASS、0 阻断、5 观察项（OBS-01~05）；OBS-01（验收 DB schema_migrations 缺 0011 记录）已通过启动后端走 RunMigrations 幂等补记处置；OBS-02~05 不阻断已留痕于 checklist。另完成实弹 smoke：8 RPC 可达 + Create→Bind→brief standards[] 全树→Unbind→Delete 全链 200。
  - [x] SubTask 9.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交

# Task Dependencies

- Task 2, Task 3 depend on Task 1（模块代码依赖生成产物）
- Task 4 depends on Task 3（validate 单测依赖 types/validate 实现）
- Task 5 depends on Task 3（集成测试依赖模块与 0011 建表）
- Task 6 depends on Task 1, Task 3（StandardReader 依赖 standard 模块与 proto 生成）
- Task 7 depends on Task 3, Task 6（装配依赖模块与 projectcontext 接入）
- Task 8 depends on Task 1-7
- Task 9 depends on Task 8
