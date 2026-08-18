# Tasks

- [x] Task 1: 建立 phase14-02 关系边界冻结 spec 工件
  - [x] SubTask 1.1: 创建 `phase14_02_freeze_standard_entity_relation_boundaries/` 三件套，冻结多态绑定关系矩阵（4 target_type × 2 role 八格组合合法性 + 存在性校验 + 唯一约束 + 扩展规则）
    - 执行记录：八格矩阵冻结（template_source 仅 repository 合法 + adopts 全开），附组合判定与扩展两个 Scenario；存在性校验 invalid_argument、唯一约束四元组、brief 反查装配规则一并冻结。
  - [x] SubTask 1.2: 冻结画像退役六触点执行层确认（T1 proto 包 / T2 存储表 / T3 后端模块收缩 / T4 RPC / T5 前端切片 / T6 brief 装配；每触点文件级现状定位 + 动作 + 断言）
    - 执行记录：六触点全部基于本轮实际探查（Read/ls/find/grep）落定文件级现状：T1 proto 190 行结构清单、T2 三表 L21/L47/L69 + 主表保留、T3 模块 5 条目、T4 路由清理、T5 前端 10 文件 + 挂载 L27/L305、T6 三处跨包引用 L39/L165/L190/L191。
  - [x] SubTask 1.3: 留档 GetProjectBrief 字段演进前后对照表（7→8 字段；global_assets BREAKING 移除 + reserved 3；governance_profile 类型替换；standards 新增；零丢失映射规则）
    - 执行记录：对照表逐字段冻结（1/4/5-7 不变、2 类型替换、3 BREAKING+reserved、8 新增），附第 9 顶层块禁止约束与零丢失映射 Scenario。
  - [x] SubTask 1.4: 冻结 BriefGovernanceProfile 内联轻量消息字段清单（3 字段保留 + 5 类字段显式不迁移）与 current_phase 三字段保留口径（数据/读取/写入/演进四层冻结）
    - 执行记录：内联消息 proto 草案冻结（3 字段 + 2 内联枚举值域一致）；不迁移字段清单 6 条（含复核后补入的 markdown_resolvable 显式退役说明）；current_phase 四层口径冻结。

- [x] Task 2: 一致性核对
  - [x] SubTask 2.1: 核对 spec 现状定位与仓库实际零漂移
    - 执行记录：wc -l 证实 proto 恰 190 行；sed 证实 0010 迁移 L21/L47/L69 三表行号精确；ls 证实模块 5 条目；find 证实前端切片精确 10 文件（复核代理未数清处由主代理实证）；grep 证实挂载 L27/L305 与 project_context.proto L39/L165/L190/L191；GetProjectBriefResponse 唯一定义源在 proto L188（Go 为生成物）。
  - [x] SubTask 2.2: 核对 spec 与 phase14 三件套及 phase14-01 单值一致
    - 执行记录：template_source/adopts 语义与 baseline §3.4（L103）一致；六触点与 arch §4.7 一致（phase14-01 复核已建立）；brief 演进与 baseline §3.6（standards[] 新增/governance_profile 收缩/两清单迁移）一致；裁决④ phase15 口径（spec 内 2 处）与裁决⑧合并口径与 phase14-01 二值化表逐字一致；12 字段划分闭环（3 保留 + 9 退役）。
  - [x] SubTask 2.3: 确认零代码改动、零三件套正文改动、零根级改动（git status 验证）
    - 执行记录：git status 过滤后无其他变更（grep exit=1），本子任务仅新增 phase14_02 目录。

- [x] Task 3: 独立复核与收口
  - [x] SubTask 3.1: 子代理独立复核（现状定位真实性抽查、矩阵完备性、对照表零丢失映射闭环、与上游零漂移）
    - 执行记录（2026-08-18 独立复核代理）：A 项抽查 7 点中 5 点直接确认（proto 190 行全结构 / SQL 三表行号 / 模块包结构 / 挂载行号 / project_context 引用行号），2 点标记未确认——①前端文件数：主代理 find 实证恰 10 文件（复核代理引用的文件列表本身即这 10 个），闭环；②GetProjectBriefResponse Go 定义位置：属 proto 生成物，唯一定义源在 project_context.proto L188，spec 定位正确，闭环。复核发现 1 个真缺口：C 项 markdown_resolvable 去向未说明——已修复（spec 显式不迁移清单补入第 4 条：服务端计算状态、语义被裁决③结构性吸收、显式退役不进零丢失验收）。B~F"待核查"项由主代理在 Task 2.2 逐项完成（见执行记录）。最终结论：PASS（2 项 FAIL 均为复核代理未完成验证而非 spec 缺陷；1 项真缺口已修复）。
  - [x] SubTask 3.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交
    - 执行记录：checklist.md 全部勾选；变更未提交（phase14_02 目录 untracked），待用户确认后手动提交。

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
