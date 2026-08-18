# phase14-07 Checklist

## proto 合同

- [x] `standard.proto` 4 枚举值序 / 4 核心消息字段号 / 8 RPC envelope 与 phase14-04 §ADDED-1/2 逐字一致；文件头四段注释齐全
- [x] `project_context.proto` 追加 import + `standards = 8`；`global_assets = 3` / `governance_profile = 2` 未动（过渡态）；brief 顶层注释与字段面一致
- [x] `make gen` 三端产物齐备（Go proto / Go connect / TS）；`make lint && make build && make breaking` 全绿

## 0011 迁移（第一段建表）

- [x] 三表 DDL（列名/类型/约束/索引）与 phase14-03 §ADDED-1 逐字一致；IF NOT EXISTS 幂等形态
- [x] 仅建表：无数据迁移 DO 块、无 drop 画像两表（phase14-09 范围）；头注释含追加段预留声明
- [x] 迁移执行后三表建成、schema_migrations 记录 0011；画像表原样保留
  - 留痕：验收 DB 曾被手动 apply 三表而缺 version 记录（复核 OBS-01）；已通过启动后端走 RunMigrations 幂等补记，日志 `migration applied version=0011_phase14_standard_entity`。

## standard Go 模块

- [x] 8 文件清单与 phase14-04 §ADDED-3 一一对应（connect/server / service/command / service/query / repository/standard_store / candidate/target_reader / errors / types / validate），无多余文件无职责漂移
- [x] types.go DirectoryTreeNode 结构/json tag/omitempty 与 phase14-03 §ADDED-2 逐字一致；受控枚举转换单值
- [x] errors.go 5 哨兵齐备且注释对齐既有风格
- [x] validate.go R1-R8 判定逻辑/稳定错误码/节点路径定位与 phase14-03 §ADDED-3 一致；R2 三时机（创建/整树替换/状态变更）；R6 ≤65536 字节
- [x] UpdateStandard 单事务边界（树替换 + updated_at + revision 追加同事务）
- [x] BindStandard 校验链顺序固定：standard 存在 → 枚举 → 八格矩阵 → target 存在 → 唯一约束
- [x] CreateStandard 不记 revision；ListStandards 按 updated_at DESC；ListRevisions 按 created_at DESC 不分页
- [x] connect 层无 SQL / repository 层无业务判断 / service 层无跨模块直写 SQL

## 测试覆盖

- [x] validate_test.go R1-R8 逐条非法用例（8 错误码全覆盖 + 节点路径断言）+ 合法边界 + nil/[] 等价 + omitempty 反序列化 + 合法完整树
- [x] 集成测试：UpdateStandard 事务 round-trip；StandardReader 反查（两 role 含全树 / 无绑定空列表）
- [x] 集成测试错误语义抽查：name 重复 / RETIRED 拒绝 / active 空树 / Delete active 拒绝 / template_source 携非 repository / target 不存在
- [x] 集成测试沿袭既有模式（真实 DB + 环境检测跳过，不可用时 skip 不 fail）

## projectcontext 接入

- [x] StandardReader 接口签名逐字 phase14-04 §ADDED-5；实现位 query_service.go；projectcontext 无 standard 表直接 SQL
- [x] GetProjectBrief 追加 standards[] 装配；既有 7 步编排与失败语义不回退
- [x] ProjectBriefReadResult 追加 Standards 字段；connect 响应组装含递归树转换
- [x] brief 集成测试 standards[] round-trip 断言（绑定后含全树零丢失 / 未绑定空数组）；global_assets 行为与 phase13 现状一致（过渡态并存）

## 装配与错误映射

- [x] connecterrors 注册 standard 5 哨兵（NotFound 2 / InvalidArgument 1 / Internal 显式留痕 2）
- [x] router.go buildStandard / mountStandardConnect / buildProjectContext 双 Reader 签名扩展落地；server.go 调用点闭合
- [x] 后端启动后 8 个 RPC 经 `/api/psco.standard.v1/StandardService/...` 可达
  - 留痕：实弹 smoke 验证——ListStandards 200；GetStandard 非法 UUID → 400 invalid_argument；CreateStandard → BindStandard(adopts) → GetProjectBrief 返回 standards count=1 含全树（根 "." + AGENTS.md，ref round-trip）→ Unbind → Delete 全链 200，数据已清理。

## DoD 门禁

- [x] `make lint / build / breaking` 通过
- [x] `go build ./... && go vet ./... && go test ./...` 通过
- [x] brief 集成测试含 standards[] round-trip 断言（验收 DB 执行留痕）
  - 留痕：projectcontext/connect 集成 21.2s 真实执行（127.0.0.1:55432/psco_development）；standard/service 集成 5 用例真实执行全 PASS。
- [x] 树校验单测含非法树用例全错误码覆盖

## 复核与收口

- [x] 子代理独立复核通过（发现项已修复回填）
  - 留痕：独立复核 11 项全 PASS、0 阻断、5 观察项；OBS-01（schema_migrations 缺 0011 记录）已处置见上；OBS-02（skip 语义偏差，spec 允许）/ OBS-03（connect 层枚举前置拦截，非跳步）/ OBS-04（Marshal 失败防御性归类，不可达分支）/ OBS-05（长度按 rune 判定，实现与测试自洽）不阻断，已留痕。
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
