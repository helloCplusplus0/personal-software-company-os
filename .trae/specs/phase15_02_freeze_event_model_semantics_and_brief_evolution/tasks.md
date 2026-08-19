# Tasks

- [x] Task 1: 完成本 spec 三件套并冻结语义正文
  - [x] SubTask 1.1: 撰写 spec.md：三轨 × event_kind 合法矩阵（3×4 全 12 格）+ task_key 格式规则（K-1 ~ K-5 正则）+ 两条显式不做约束；9 条校验规则清单级前置确认；派生规则语义冻结（4 项派生 + 三键链全序 + 派生执行位 + current_phase_key 空值双情形同型语义）；brief 演进前后对照表（前侧 proto 现状 / 后侧 progress = 9 / BriefProgress 字段清单 / 装配约束）；三重边界分离显式断言（每重 3 条共 9 条）；与 phase15-01 边界上游一致性声明
    - 执行记录：spec.md 已按 6 项 ADDED Requirements 完成（合法矩阵与 task_key 格式 / 9 条校验规则清单级确认 / 派生规则含空值边界 / brief 前后对照表 / 三重边界 9 条断言 / 与 phase15-01 一致性声明）；矩阵 12 格与 K-1~K-5 正则逐格逐条对齐 shared_baseline §3.3；brief 前侧对照表基于 project_context.proto L191-204 实读内容（本对话 Read 取证）非凭记忆转写；错误码 / DDL / RPC 签名 / 字段号分配 / 组件树均未承载（防偷渡）
  - [x] SubTask 1.2: 撰写 tasks.md / checklist.md（本文件与 checklist）
    - 执行记录：tasks.md 含 3 任务 / 7 子任务；checklist.md 含 13 检查点；依赖关系与 dev_plan §5 一致（Task 2 depends on Task 1；Task 3 depends on Task 2；phase15-03/04/05 depends on 本 spec）
- [x] Task 2: 一致性校验
  - [x] SubTask 2.1: 与 phase15 三件套单值一致校验：合法矩阵 + task_key 格式 + 9 条校验规则 vs shared_baseline §3.3 + architecture_plan §4.5；派生规则 vs shared_baseline §3.4 + architecture_plan §4.5；brief 对照表 vs shared_baseline §3.5 + architecture_plan §4.7；三重边界 vs architecture_plan §4.6 + shared_baseline §2.2 裁决⑪
    - 执行记录：逐项比对通过——①合法矩阵 12 格逐格与 architecture_plan §4.5 矩阵一致（phase 轨 4 格全合法含 task_key 要求、audit/fix 轨边界事件禁止、task_completed/note 两格含格式要求）；②K-1~K-5 正则与 shared_baseline §3.3 规则 3/4/5/6/8 的正则逐字一致；两条显式不做约束（task_key 唯一性不做、occurred_at 未来时间不校验）沿 §3.3 注；③9 条校验规则逐条引用 §3.3 原文未改写（第 9 条 evidence_ref `/`|`https://` 前缀 + title≤200 + detail≤2000 + source 仅 manual 逐项一致）；④派生规则 4 项与 §3.4 矩阵逐项一致，三键链全序与 tiebreak 两类碰撞场景沿 §3.2 末段 + §3.4，派生执行位（后端 service 层 / web 与 brief 共用 / 不落库不缓存）逐字一致；空值双情形中"全部完结态（空值）"为 §3.4 原文，"从未开始"为 §3.4"当前 phase = 最新 phase_started 的 task_key"在无 phase_started 时的自然空值补齐，不冲突不扩界；⑤brief 对照表与 §3.5 演进矩阵逐项一致（progress=9 / 槽位 2-3-4 reserved / BriefProgress 4 字段 / 同型复用 ProgressEvent / 空态恒构造 / ProgressReader 装配 / 5→5+progress 顶层口径）；⑥三重边界 9 条断言为 architecture_plan §4.6 三段结论的机械化展开（plan.md 零托管 / PhaseEntry 不改动不映射不裁决 / Decision 不合并不互替不互链数据零交叉），无超出 §4.6 语义的新增约束
  - [x] SubTask 2.2: 与 phase15-01 spec 边界上游一致校验（四组成部分映射 / 裁决③④⑤⑩⑪承接 / 口径辨析引用不重复展开）；与 project_context.proto 实际内容逐字段比对（brief 前侧 L191-204：repository=1 + reserved 2/3/4 + products/modules/decisions=5/6/7 + standards=8；PhaseEntry L101-107 五字段；既有 import standard.proto 模式 L42）；git 工作区确认零 proto / 代码 / 根级文档改动
    - 执行记录：①与 phase15-01 spec 一致——四组成部分映射声明成立（矩阵/校验规则→组成部分 1、派生→组成部分 4、brief→组成部分 3、三重边界→裁决⑪展开），裁决③④⑤⑩⑪承接逐项对应 phase15-01 裁决矩阵行，口径辨析引用 phase15-01 留档未重复展开；②与 project_context.proto 实读内容（本对话 Read 取证）逐字段比对——brief 前侧对照表 8 行与 L191-204 一致（repository=1〔L192〕、reserved 2,3,4 含 "governance_profile","global_assets","current_phase" 三字段名〔L198-199〕、products=5〔L200〕、modules=6〔L201〕、decisions=7〔L202〕、standards=8〔L203〕）；PhaseEntry 五字段 phase/label/status_summary/entry_ref/entry_kind〔L101-107〕与边界二断言依据一致；既有 import "psco/standard/v1/standard.proto"〔L42〕与跨包导入模式声明一致；③git status --porcelain 复验——工作区仅 `.trae/specs/phase15_02_freeze_event_model_semantics_and_brief_evolution/` 目录 untracked 新增（phase15-01 及根级改动已被用户提交，不在工作区），零 proto / backend / frontend / database 改动，零根级回写
- [x] Task 3: 独立复核与收口
  - [x] SubTask 3.1: 子代理独立复核（合法矩阵与校验规则单值性 / 派生语义含空值边界无歧义 / brief 对照表前侧与 proto 实际一致且后侧无偷渡 / 三重边界 9 条断言可机械判定 / 无 phase15-03/04/05 设计细节偷渡 / tasks-checklist 质量与勾选真实性）
    - 执行记录：独立复核六维度全 PASS——A 合法矩阵 12 格逐格比对 + K-1~K-5 正则逐字一致 + 9 条规则零增减 + 清单级/形式化边界清晰（唯一 invalid_argument 出现于排除声明）；B 派生 4 项与三键链逐字一致 + 空值双情形为 §3.4 自然延拓补齐非改写 + 不构成 BriefPhaseStatus 变相复活（CON-08 红线零冲突）+ DP-1 归属与 phase15-01 登记单值一致；C 前侧 8 行经复核方独立实读 proto 逐字段行号级吻合（L192/L198-199/L200-203）+ 后侧与 §3.5 全项一致 + 字段号无偷渡（表格无字段号列且显式声明归 phase15-04）；D 9 条断言语义从属 §4.6 非目标 10、均可经代码检视+运行时取证机械判定、无遗漏；E 六类后续设计细节零承载；F DoD 四要素全覆盖 + Task 1/2 勾选经独立比对全部属实（proto 行号级实读证据 + git 工作区实测吻合 HEAD 为 phase15-01 提交）+ checklist 13 项覆盖 6 项 ADDED Requirements。最终结论 PASS（0 blocker，4 observations）
  - [x] SubTask 3.2: 修复独立复核发现的阻断性问题（如有）并复验
    - 执行记录：0 blocker 无需修复；4 项非阻断 observations 处置——Obs-1（空值同型后 phase15-05 若经 brief 通道则前端无法区分"尚未开始/已完结"文案）已由 spec 条件措辞将裁决留给 phase15-05，作为 DP-1 裁决时必处理事项随本执行记录移交；Obs-2（BriefProgress"自然序"软指引为非约束性预置倾向）phase15-04 冻结时可忽略，不改文；Obs-3（checklist 统一勾选时序属既定收口模式）无动作；Obs-4（边界三-3"数据零交叉"与边界二-1 尾注为具体化展开）移交 phase15-08 按"派生 SQL 检视 + GetProjectContext 行为回归"口径取证
  - [x] SubTask 3.3: tasks.md / checklist.md 全部勾选附执行记录；变更保持未提交，待用户最终确认后手动提交
    - 执行记录：本文件与 checklist.md 已全部勾选附执行记录及勾选依据；git 工作区最终态 = 仅本 spec 目录新增（3 文件）；变更未提交，待用户最终确认后手动提交

# Task Dependencies

- Task 1（三件套产出）为本 spec 获用户批准后的首个执行项
- Task 2 depends on Task 1
- Task 3 depends on Task 2
- 后续：phase15-03 / 04 / 05 depend on 本 spec 收口（dev_plan §5：三者均 depends on phase15-02）
