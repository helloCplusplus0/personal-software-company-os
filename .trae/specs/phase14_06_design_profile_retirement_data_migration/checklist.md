# phase14-06 Checklist

## 0011 迁移设计冻结

- [x] 三段结构冻结（建表 DDL / 数据迁移 / drop 两表；先迁数据后 drop 顺序；单事务原子——migrate.go applyMigration 整文件单事务实证）
- [x] 两段式算法冻结（TEMP TABLE 节点行物化 + priority 1/2 聚合同名合并（源 2 非空优先）+ WITH RECURSIVE 自顶向下组树 + 幂等产物写入）
- [x] 字段映射引用 discipline（逐字引用 phase14-03，本 spec 不重复定义映射规则）
- [x] 迁移产物固定值冻结（name=默认项目范式（迁移自治理画像）/ status=active / description 固定 / 首条 revision 动态计数 / adopts 绑定决策留档（不自动 template_source 理由）/ 源 repository 取主表最新 updated_at）
- [x] 无画像数据场景决策冻结（零行不生成空树 Standard，不产生 revision/binding）
- [x] 幂等守卫矩阵冻结（CREATE IF NOT EXISTS / INSERT WHERE NOT EXISTS / ON CONFLICT DO NOTHING / DROP IF EXISTS / 数据段 to_regclass 守卫——两表已 drop 的重放不报错）

## 后端模块收缩冻结

- [x] 文件级清单冻结（删 4：connect/server.go + server_test.go + command_service.go + candidate/repository_reader.go；收缩 4：query_service / profile_store / types / errors；保留测试收缩）
- [x] candidate 删除条件确认（本轮实证其仅服务画像写入前提校验——T3 条款适用）
- [x] GovernanceProfileReader 接口签名收缩冻结（ReadProfileCore + GovernanceProfileCoreReadResult 三字段组；失败语义保留）
- [x] ReadProfile 收缩必要性论证留档（两表 drop 后 JOIN 必然失败——收缩是前置必要非可选）
- [x] router.go 行级变更清单冻结（L69/L70/L90 import 删 / L305-313 mount 删 / L407-409 双 Reader 注入 / L412-422 build 收缩）

## 前端切片移除冻结

- [x] 10 文件清单 + 挂载点（08 已让位、09 删目录清扫）+ TS 生成产物清理（make gen 重生成自动消失）
- [x] 时机与三断言（目录不存在 / grep 零命中 / tsc 零错误）

## brief 切换与 buf 豁免冻结

- [x] 装配切换执行序四步（proto 四处变更引用 02 / 后端双 Reader / 前端重生成 / 字段面断言）
- [x] buf.yaml breaking.ignore 单文件路径方案 + 注释留痕（决策可追溯）+ 不扩大化约束（禁止目录前缀豁免）
- [x] make breaking 豁免后全绿要求

## 六触点断言矩阵执行化

- [x] T1-T6 每触点机械检查命令冻结（find/grep/psql/tsc/buf 可直接执行）
- [x] 迁移专属断言三条（固定名仅一条 / 首条 revision 含计数与源 id / 零丢失对照引用 02 Scenario）

## 一致性核对

- [x] 与 phase14-03 单值一致（映射引用无重复定义；算法忠实落地；产物天然满足 R1-R8 论证）
- [x] 与 phase14-02 单值一致（六触点/对照表/内联消息引用无漂移；candidate 删除条件确认）
- [x] 与 phase14-04/05 衔接一致（模块并存形态 / StandardReader 双注入 / 让位时序）
- [x] 与仓库机制实证一致（migrate.go 单事务 / buf.yaml v2 / router.go 行号 / 模块文件行数）
- [x] dev_plan L44-47 范围与 DoD 逐条满足
- [x] 零代码改动、零三件套正文改动、零根级改动（git status 验证）

## 复核与收口

- [x] 子代理独立复核通过（发现项已修复回填 spec）
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
