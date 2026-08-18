# Tasks

- [x] Task 1: 建立 phase14-06 画像退役与数据迁移设计 spec 工件
  - [x] SubTask 1.1: 冻结 0011 迁移三段结构（建表 DDL / 数据迁移 / drop 两表先迁后删；单事务原子依据 migrate.go L103-124 实证）
    - 执行记录：三段结构冻结（spec §ADDED-1）；原子性依据 = applyMigration L110 整文件单 tx.Exec + L113 版本记录同事务；0011 文件由 phase14-07 落地（本轮实证 database/migrations/ 现止于 0010）。
  - [x] SubTask 1.2: 冻结数据迁移两段式算法（DO 块内节点行物化 TEMP TABLE + priority 聚合同名合并 + 自底向上固定轮次组树（独立复核修正：弃 WITH RECURSIVE 形态）+ 幂等产物写入；字段映射逐字引用 phase14-03 不重复定义）
    - 执行记录：算法四步冻结（spec §ADDED-1 第二段）：DO 块内 TEMP TABLE ON COMMIT DROP 物化（priority 1/2 区分两源）→ GROUP BY path + COALESCE(MAX CASE) 集合化同名合并（源 2 非空优先，03 规则的 SQL-native 落地）→ 自底向上固定 6 轮逐层聚合物化组树（subtree 列 + 每轮非递归 UPDATE，相关子查询作用于普通临时表无递归 CTE 限制；R5 深度 ≤6 覆盖，重复轮次幂等）→ 三条幂等产物写入；映射规则全部引用 03 无重复定义。独立复核 FAIL-1 修正：原 WITH RECURSIVE 自顶向下每层 jsonb_agg 形态违反 PG 递归项双限制（递归项内禁聚合、子查询内禁递归引用），已改形并留痕。
  - [x] SubTask 1.3: 冻结迁移产物固定值（name 固定名幂等键 / status=active / description / 首条 revision 动态计数 change_summary / adopts 绑定决策留档 + 源 repository 取主表最新 updated_at）与无画像数据场景决策（不生成空树 Standard）
    - 执行记录：固定值表冻结（spec §ADDED-1）；adopts 决策留档（template_source 为强语义主张不由迁移自动断言）；零行场景不生成空树 Standard（避免无意义 draft 行）。
  - [x] SubTask 1.4: 冻结幂等与可重放设计（CREATE IF NOT EXISTS / INSERT WHERE NOT EXISTS / ON CONFLICT DO NOTHING / DROP IF EXISTS / 数据迁移段收敛为单 DO 块 + to_regclass 真跳过守卫——手工 psql 重放安全）
    - 执行记录：幂等守卫矩阵冻结（spec §ADDED-1）；DO 块守卫覆盖"两表已 drop 后重放报错"场景——schema_migrations 单次执行（migrate.go L65-68）与手工 psql -f 重放（phase14-10 runbook）双场景均安全；附幂等判定 Scenario。独立复核 FAIL-3 修正：原"WHERE 条件包裹 + 多条独立语句"在 psql autocommit 下 TEMP TABLE 语句级提交即删、后续语句必报 relation does not exist；收敛为单 DO 块（块首 IF to_regclass 真跳过）后两条执行路径自洽，已改形并留痕。
  - [x] SubTask 1.5: 冻结后端模块收缩文件级清单（删 4 文件含 candidate 确认 / 收缩 4 文件行级定位 / 保留测试收缩）+ GovernanceProfileReader 接口签名收缩设计（ReadProfileCore + GovernanceProfileCoreReadResult）+ router.go 行级变更清单（L69/70/90/305-315/407-409/412-422）+ 收缩联动 2 文件（server.go / connect_errors.go，独立复核 FAIL-2 补齐）
    - 执行记录：9 文件动作表冻结（spec §ADDED-2，现状行数全部本轮实证；独立复核 OBS-1 修正计数 10→9）；candidate 删除条件确认（本轮读文件实证其仅被 Query/Command 构造注入服务画像写入前提校验）；接口收缩 Go 草案冻结（ReadProfileCore 失败语义保留）；ReadProfile 收缩必要性论证留档（两表 drop 后 JOIN 必然失败）；router.go 六处行级变更定位（本轮 grep 实证行号；OBS-2 补注 phase14-09 以符号定位为准、L407-409 standardReader 注入由 phase14-07 先行落地）。独立复核 FAIL-2 修正：收缩清单补齐清单外硬编译断点——server.go L92/93/97 三处调用点 + connect_errors.go L67/102/138 哨兵映射行删除，ErrRepositoryNotFound 去向单值化（随 candidate 删除、projectcontext 同名哨兵独立不受影响），已回填 spec 联动清单。
  - [x] SubTask 1.6: 冻结前端切片移除清单（10 文件 + 挂载点已由 08 让位 + TS 生成产物清理 + 时机与三断言）
    - 执行记录：10 文件清单冻结（spec §ADDED-3）；挂载点让位时序引用 phase14-05（08 swap / 09 删目录）；gen 产物随 make gen 自动消失；三断言（目录不存在 / grep 零命中 / tsc 零错误）。
  - [x] SubTask 1.7: 冻结 brief 装配切换执行序（proto 四处变更引用 02 / 后端双 Reader 装配 / 前端重生成）与 buf breaking 豁免配置（buf.yaml v2 breaking.ignore 单文件路径 + 注释留痕 + 不扩大化约束）
    - 执行记录：切换执行序四步冻结（spec §ADDED-4，proto 变更逐字引用 02 内联消息草案）；豁免配置 YAML 草案冻结（单文件路径 psco/governance_profile/v1/governance_profile.proto + 两行注释留痕 + 长期保留说明 + 禁止目录前缀扩大化约束）。
  - [x] SubTask 1.8: 冻结六触点验收断言矩阵执行化（每触点机械检查命令 + 迁移专属断言三条）
    - 执行记录：T1-T6 断言矩阵冻结（spec §ADDED-5，每触点 find/grep/psql/tsc/buf 可直接执行的机械命令）；迁移专属断言三条（固定名仅一条 / 首条 revision 计数与源 id / 零丢失对照引用 02）。

- [x] Task 2: 一致性核对
  - [x] SubTask 2.1: 与 phase14-03 单值一致（字段映射/拆树三态/同名合并/零丢失范围全部引用无重复定义；两段式算法是对 03 映射的忠实集合化落地；产物天然满足 R1-R8 论证成立——根 name='.' / 深度 / 字符集由源数据性质保证）
    - 执行记录：算法步骤的每个映射动作（源 1 展开 / 源 2 三态 / kind 前缀格式 / priority 合并）逐字对应 03 迁移映射节表格；产物满足 R1-R8 论证——R1 根由组树起点固定、R2 迁移行全为 file、R3/R4 同层唯一与字符集由 GROUP BY path 与源数据（既有合法文件名）保证、R5 源路径深度 ≤6（phase13 资产全为根级或一层路径）、R6 单用户两表十几行远小于 64KB、R7 file 无 children 由物化结构保证、R8 ref 由规范化规则自洽生成；phase14-09 落地后仍有 GetStandard 只读断言兜底。
  - [x] SubTask 2.2: 与 phase14-02 单值一致（六触点动作与断言逐项对应；brief 对照表"后"列引用无漂移；BriefGovernanceProfile 内联消息逐字引用；T2 主表保留 / T3 candidate 删除条件确认适用）
    - 执行记录：断言矩阵 T1-T6 与 02 六触点断言逐项对应且执行化（机械命令化不改变语义）；brief 切换 proto 四处变更（内联消息 / reserved 3 / BriefPhaseStatus / import 移除）逐字引用 02 Requirement；T2 主表保留（列不增不减）在断言中体现；T3 candidate 条款经本轮读文件实证确认适用（仅服务写入前提校验）。
  - [x] SubTask 2.3: 与 phase14-04/05 衔接一致（standard 模块与收缩后画像模块并存形态；StandardReader 双注入对齐 04 装配设计；前端让位时序 08 swap → 09 删目录对齐 05）
    - 执行记录：收缩后画像模块（纯读 Reader 实现）与 standard 模块（新主线）并存形态清晰——两者互不依赖、各自经 candidate 接口服务 projectcontext；router.go L407-409 双 Reader 注入与 phase14-04 StandardReader 装配设计（注入位 router.go）一致；前端让位时序（08 swap L27/L305-306 / 09 删目录）与 phase14-05 §ADDED-6 逐字一致。
  - [x] SubTask 2.4: 与仓库机制实证一致（migrate.go 单事务整文件 / schema_migrations 单次执行 vs 手工重放双场景 / buf.yaml v2 语法 / router.go 行号定位 / 模块 10 文件行数实证）
    - 执行记录：migrate.go L103-124（单事务整文件）、L65-68（单次执行跳过）本轮读取实证；buf.yaml v2 语法（version/modules/breaking.use）本轮 cat 实证——ignore 为 v2 合法字段；router.go L69-90/L305-313/L396-422 行号本轮 grep 实证；模块 10 文件与行数（query 61/command 134/store 274/types 167/errors 31/candidate 42）本轮 find+wc 实证。
  - [x] SubTask 2.5: dev_plan L44-47 范围与 DoD 逐条满足（迁移可机械执行且幂等 / 合并策略单值 / 六触点逐项验收断言 / brief 前后对照表留档）
    - 执行记录：范围逐项覆盖——迁移映射与拆树规则引用 03（避免重复冻结）✓、迁移脚本入口与幂等语义（0011 三段 + 守卫矩阵，入口=migrate.go 既有机制不新建脚本）✓、drop 顺序（先迁后删）✓、前端切片移除清单 ✓、后端收缩文件级清单（删 connect 与写路径、保留 Reader）✓、brief 内联消息（引用 02 + 接口收缩设计）✓、buf breaking 豁免方案 ✓。DoD 四条：可机械执行且幂等（两 Scenario）✓、合并策略单值（固定值表 + priority 聚合）✓、六触点逐项断言（矩阵执行化）✓、brief 前后对照表留档（引用 02 对照表 + 切换断言）✓。
  - [x] SubTask 2.6: 零代码改动、零三件套正文改动、零根级改动（git status 验证）
    - 执行记录：git status 验证仅新增 phase14_06 目录（?? .trae/specs/phase14_06_design_profile_retirement_data_migration/），无其他变更；0011 文件确认尚不存在（database/migrations/ 现止于 0010，属 phase14-07 交付）。

- [x] Task 3: 独立复核与收口
  - [x] SubTask 3.1: 子代理独立复核（算法可实现性 SQL 级论证 / 幂等矩阵完备性 / 收缩清单与编译闭环 / 豁免配置语法与范围 / 断言矩阵可机械执行 / 与上游零漂移）
    - 执行记录：子代理独立复核完成（八维度 + 四项重点技术论证，全部实证文件/行号）。初判 FAIL：3 项阻断 + 8 项观察。阻断项已全部修复回填 spec：
      - FAIL-1 组树 SQL 形态不可实现（PG 递归项内禁聚合、子查询内禁递归引用）→ 改为 DO 块内自底向上固定 6 轮逐层聚合物化（subtree 列 + 非递归 UPDATE；R5 深度 ≤6 覆盖；重复轮次幂等），修正留痕入 spec。
      - FAIL-2 收缩清单遗漏清单外硬编译断点（server.go L92/93/97 调用点 + connect_errors.go L67/102/138 哨兵映射）→ spec 补"收缩联动文件"表与 errors.go 哨兵单值口径（ErrRepositoryNotFound 随 candidate 删除，projectcontext 同名哨兵独立不受影响）。
      - FAIL-3 重放守卫形态与 psql autocommit 不自洽（TEMP TABLE ON COMMIT DROP 语句级提交即删）→ 数据迁移段收敛为单 DO 块（块首 IF to_regclass 真跳过），migrate.go 单事务与手工 psql -f 双路径自洽，幂等 Scenario 同步改写。
      观察项处置：OBS-1 计数笔误已修（10→9、删 3/收缩 3→删 4/收缩 4）；OBS-2 行号时点注 + 符号定位 + L407-409 时序归属 phase14-07 已补注；OBS-3 并入 FAIL-2 修复；OBS-4 brief 切换补 projectcontext/types.go 与 server_integration_test.go 联动文件；OBS-5 T2 断言去占位表述改机械命令；OBS-6 挂载点补全路径（含 pages/）；OBS-7 源 2 末段 name 口径消歧（按 03 表格主口径）；OBS-8 属 phase14-02 文档行号偏差、06 自身引用无漂移，不在本 spec 修改范围（留档 phase14-09/10 执行时以字段号为准）。
  - [x] SubTask 3.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交
    - 执行记录：checklist 全部勾选收口；git status 复验确认仅 .trae/specs/phase14_06 目录新增、零代码改动、零其他文档改动；变更未提交，待用户最终确认后手动提交。

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
