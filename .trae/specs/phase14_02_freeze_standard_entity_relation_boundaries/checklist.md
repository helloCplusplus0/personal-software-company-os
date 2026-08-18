# phase14-02 Checklist

## 关系边界冻结工件

- [x] 绑定关系矩阵单值冻结（4×2 八格：template_source 仅 repository 合法、adopts 全开；存在性校验 invalid_argument；唯一约束四元组；扩展只扩 enum + 登记矩阵）
- [x] 六触点执行层确认留档（T1 proto 包 190 行 / T2 三表含主表保留 / T3 模块收缩纯读 / T4 RPC 路由清理 / T5 前端 10 文件 + 挂载 L27/L305 / T6 三处跨包引用内联化；每触点现状定位 + 动作 + 断言）
- [x] GetProjectBrief 前后对照表留档（7→8 字段；global_assets=3 BREAKING 移除并 reserved；governance_profile=2 类型替换 BriefGovernanceProfile；current_phase=4 status 内联化；standards=8 新增；第 9 顶层块禁止）
- [x] 零丢失映射规则冻结（name→节点名 / entry_ref→ref / role→role / structured_summary→summary / kind 并入摘要语义）
- [x] BriefGovernanceProfile 内联消息字段清单冻结（repository_id / track_type / template_source 三字段 + BriefTrackType / BriefPhaseStatus 内联枚举值域一致）
- [x] 六条字段显式不迁移冻结（project_profile_version / docs_workflow_layout / 两清单 / markdown_resolvable 显式退役 / current_phase 三字段重复 / created_at+updated_at）
- [x] current_phase 四层保留口径冻结（数据层主表不动 / 读取层仅 brief 顶层块 / 写入层无 / 演进层 phase15 进入条件）

## 一致性核对

- [x] 现状定位与仓库实际零漂移（六触点文件/行号逐一实证：proto 190 行、SQL L21/L47/L69、模块 5 条目、前端精确 10 文件、挂载 L27/L305、proto 引用 L39/L165/L188/L190/L191）
- [x] 与 shared_baseline §3.4/§3.5/§3.6 单值一致
- [x] 与 architecture_plan §4.4/§4.7/§4.8 单值一致
- [x] 与 phase14-01 二值化表（裁决④⑧）口径一致
- [x] template_source 仅 repository 的约束理由与 CON-02 不设限的边界说明无矛盾
- [x] 零代码改动、零三件套正文改动、零根级改动（git status 验证）

## 复核与收口

- [x] 子代理独立复核通过（A 项 7 点中 5 点直接确认、2 点由主代理实证闭环；C 项发现 markdown_resolvable 去向缺口已修复；B~F 由主代理 Task 2.2 逐项完成）
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
