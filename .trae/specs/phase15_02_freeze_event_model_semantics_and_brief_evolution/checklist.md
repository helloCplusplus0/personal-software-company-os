# phase15-02 Checklist

- [x] 三轨 × event_kind 合法矩阵单值冻结（3 workflow × 4 kind 全 12 格判定逐格与 shared_baseline §3.3 / architecture_plan §4.5 一致：phase 轨 4 格全合法含 task_key 要求、audit/fix 轨 phase_started/phase_completed 禁止）
- [x] task_key 格式规则 K-1 ~ K-5 冻结（4 组必填正则 + note 可空不强制格式），与 shared_baseline §3.3 规则 3-6/8 单值一致
- [x] 9 条校验规则完成清单级前置确认（逐条引用不改写；两条显式不做约束〔task_key 唯一性不做 / occurred_at 未来时间不校验〕留档）
- [x] 派生规则 4 项冻结（当前 phase key / label / 最新完成任务项 / recent_events N=10），三键链全序 `(occurred_at DESC, created_at DESC, id DESC)` + tiebreak 两类碰撞场景 + 派生执行位（后端 service 层统一计算、web 与 brief 共用、不落库不缓存）
- [x] current_phase_key 空值双情形同型语义冻结（从未开始 / 全部完结均为空字符串零值；不引入第二套状态字段或枚举；web 展示区分归 phase15-05 且受 DP-1 后端统一派生约束）
- [x] brief 演进前后对照表留档：前侧与 project_context.proto L191-204 逐字段一致（repository=1 / reserved 2,3,4 含字段名 / products=5 / modules=6 / decisions=7 / standards=8）；后侧 progress=9 新增、槽位 2/3/4 保持 reserved、1-8 零改动
- [x] BriefProgress 字段清单语义冻结（current_phase_key / current_phase_label / latest_task_completed 可选 / recent_events[]；同型复用 psco.progress.v1.ProgressEvent 不建第二套摘要消息）；空态恒构造（非 nil、零值、空数组）；跨包 import 关系声明；字段号分配显式留给 phase15-04
- [x] 三重边界各有显式断言语句（plan.md 3 条 / PhaseEntry 3 条 / Decision 3 条，共 9 条均可机械判定，可进 phase15-08 裁决⑪门禁验证）
- [x] 与 phase15-01 边界上游一致（四组成部分映射声明 + 裁决③④⑤⑩⑪承接 + 口径辨析引用 phase15-01 留档不重复展开）
- [x] 无 phase15-03/04/05 设计细节偷渡（无错误码 / 无 DDL / 无 RPC envelope 签名 / 无 BriefProgress 字段号分配 / 无 ProgressReader 接口签名 / 无前端组件树与交互规格）
- [x] 独立复核通过（0 阻断；复核维度：合法矩阵与校验规则单值性 / 派生语义无歧义 / brief 对照表前侧实证 / 三重边界可判定 / 无偷渡 / 勾选真实性）
- [x] tasks.md / checklist.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交

---

## 勾选依据（收口留档）

- 检查点 1-3：Task 2 SubTask 2.1 逐项比对 + 独立复核维度 A 独立复验（12 格逐格、K 正则逐字、9 条零增减、清单级/形式化边界清晰）
- 检查点 4-5：Task 2 SubTask 2.1 ④ + 独立复核维度 B（空值双情形为 §3.4 自然延拓补齐、不构成 BriefPhaseStatus 变相复活、DP-1 归属与 phase15-01 单值一致）
- 检查点 6-7：Task 2 SubTask 2.1 ⑤ + SubTask 2.2 ② + 独立复核维度 C（前侧 8 行经复核方独立实读 proto 行号级吻合；字段号无偷渡）
- 检查点 8：独立复核维度 D（9 条断言语义从属 §4.6 与非目标 10、可机械判定、无遗漏）
- 检查点 9：Task 2 SubTask 2.2 ① + 独立复核维度 F2
- 检查点 10：独立复核维度 E（六类后续设计细节零承载）
- 检查点 11：独立复核最终结论 PASS（0 blocker，4 observations；Obs-1 移交 phase15-05 DP-1 裁决必处理事项、Obs-4 移交 phase15-08 取证口径，Obs-2/Obs-3 留档无动作）
- 检查点 12-13：Task 3 SubTask 3.3 收口执行记录；git 工作区最终态仅本 spec 目录新增（3 文件）
