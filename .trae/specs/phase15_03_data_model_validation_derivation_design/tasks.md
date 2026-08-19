# Tasks

- [x] Task 1: 完成本 spec 三件套并冻结设计正文
  - [x] SubTask 1.1: 撰写 spec.md：progress_events DDL 级设计（注释版 SQL 草案 + 5 项设计决策冻结〔FK RESTRICT / DB 只做单列 CHECK / 索引不含 id / 无 updated_at / 无迁移段〕）；9 条校验规则形式化（错误码总表 9 业务码 + 2 envelope 码 + V8 无码结构性说明 + 执行序 6 步 + TrimSpace 规范化边界 + rune 计量）；派生算法实现序（SQL ORDER BY 唯一排序位 + repository 单查询形态 + service 纯函数 + 三派生项精确算法含 DESC 索引语义 + 空值双情形 + 个人规模边界）；与 phase15-02 单值一致与不偷渡声明
    - 执行记录：spec.md 已按 4 项 ADDED Requirements 完成（DDL 含注释版 SQL 草案可逐字转写 / 校验形式化 9+2 错误码 / 派生实现序与三算法 / 一致性与不偷渡声明）；DDL 草案基于实读 0011（幂等模式）+ 0006（FK RESTRICT 惯例）+ 0002（repositories 表实证）转写；错误码沿实读 standard errors.go（哨兵 + Connect 映射）与 validate.go（errCode 常量 + %w: [CODE] msg 格式）模式；派生算法实现序为方案"SQL 排序唯一执行位 + service 纯函数"（List / brief / web 当前卡三消费方共用单查询 + 单派生函数）；V8 无码 / envelope 2 码显式区分防规则清单 1:1 破坏；DP-2 承接位未裁决仅声明语义边界
  - [x] SubTask 1.2: 撰写 tasks.md / checklist.md（本文件与 checklist）
    - 执行记录：tasks.md 含 3 任务 / 7 子任务；checklist.md 含 14 检查点；依赖关系与 dev_plan §5 一致（phase15-06 depends on 本 spec + phase15-04；phase15-04 消费本 spec 错误码总表）
- [x] Task 2: 一致性校验
  - [x] SubTask 2.1: 与语义上游单值一致校验：DDL 11 列 vs shared_baseline §3.2 字段矩阵 + 枚举三列值域 vs §3.2/§3.3 + 索引 vs §3.2 冻结形态；V1-V9 判定逻辑 vs phase15-02 spec 合法矩阵 12 格 + K-1~K-5 正则 + 9 条规则清单；三派生算法 vs phase15-02 派生冻结（4 项 + 三键链 + 空值双情形同型）；执行序与报第一个错误策略 vs standard validate.go 模式
    - 执行记录：逐项比对通过——①DDL 11 列与 shared_baseline §3.2 字段矩阵逐列一致（id uuid PK / repository_id uuid FK NOT NULL / workflow_type 三值 / event_kind 四值 / task_key text 可空 / title NOT NULL 200 / detail 可空 2000 / evidence_ref 可空 / source 三值默认 manual / occurred_at NOT NULL / created_at NOT NULL）；枚举三列 CHECK 值域与 §3.2 + §3.3 规则 1 一致；索引 (repository_id, occurred_at DESC, created_at DESC) 为 §3.2 冻结形态，id DESC tiebreak 由 ORDER BY 补齐与 §3.2"读取全序为三键链"表述一致；②V1a/V1b 对应规则 1 两半；V7 判定 `NOT (audit/fix × phase_started/phase_completed)` 与 phase15-02 矩阵 4 禁止格一致；TASK_KEY_REQUIRED/FORMAT_INVALID 覆盖规则 2-6 必填与格式（K-1~K-4 正则经 phase15-02 K 表引用）；V9a-d 与规则 9 四要素一致（title 非空+200 / detail 2000 / evidence_ref 前缀 / source 仅 manual）；③三派生算法与 phase15-02 派生冻结逐项一致（recent N=10 / latest task_completed 三轨同序 / current phase 双步含"j < latestStartedIdx 即序更晚"的 DESC 索引语义 = phase15-02"序更晚的 phase_completed"精确化 / 空值双情形同型零值）；④报第一个错误策略与 standard validate.go L54"DFS 报第一个错误"模式一致；⑤9 条规则 1:1 对应零增减（V8 允许规则无码为结构性处理）
  - [x] SubTask 2.2: 与既有实现模式一致校验：DDL 草案 vs 0011 第一段幂等模式（IF NOT EXISTS / 注释头格式 / gen_random_uuid / TIMESTAMPTZ）；FK RESTRICT vs 0006 product_repositories 惯例；错误码命名与 `%w: [CODE] msg` 格式 vs standard errors.go + validate.go；rune 计量 vs standard maxSummaryRunes；迁移登记机制声明 vs phase14-07 OBS-01 修复后机制；git 工作区确认零代码 / 零迁移文件 / 零 proto / 零根级改动
    - 执行记录：逐项比对通过——①DDL 草案幂等模式（CREATE TABLE / INDEX IF NOT EXISTS）、注释头格式（文件名 + 定位 + 上游规格 + 机制说明）、gen_random_uuid() PK、TIMESTAMPTZ NOT NULL DEFAULT NOW() 均沿 0011 第一段实读模式；②FK `REFERENCES repositories(id) ON DELETE RESTRICT` 沿 0006 product_repositories L71-72 实读惯例（repositories 表经 0002 L16-20 实证：id UUID PK）；③错误码 SCREAMING_SNAKE_CASE 命名 + errCode 常量模式 + `%w: [CODE] message` 包装格式沿 standard validate.go L30-40 + L58 实读模式；ErrInvalidInput 哨兵 + CodeInvalidArgument 映射沿 standard errors.go L20-22；④rune 计量沿 standard validate.go L46 maxSummaryRunes 模式；⑤迁移登记机制（落入 database/migrations/ 自动被 RunMigrations 文件名升序登记，无需手工登记）沿 phase14-07 OBS-01 修复后机制（dev_plan phase15-06 范围原文引用）；⑥git status --porcelain 复验——工作区仅 `.trae/specs/phase15_03_data_model_validation_derivation_design/` 目录 untracked 新增（phase15-02 已被用户提交不在工作区），零 backend / frontend / database / proto 改动，零根级回写
- [x] Task 3: 独立复核与收口
  - [x] SubTask 3.1: 子代理独立复核（DDL 可逐字转写性 / 校验规则可测试性〔每码可触发、执行序可断言〕/ 派生算法精确性〔排序键、DESC 索引语义、空值双情形、tiebreak〕/ 与 phase15-02 1:1 转译零语义漂移 / 无 phase15-04/05/06 偷渡 / 与既有模式一致性 / tasks-checklist 质量与勾选真实性）
    - 执行记录：已指定独立复核子代理对 spec.md / tasks.md / checklist.md 三件套执行七维度复核（DDL 可逐字转写性 / 校验规则可测试性 / 派生算法精确性 / 与 phase15-02 1:1 转译零语义漂移 / 无 phase15-04/05/06 偷渡 / 与既有实现模式一致性 / tasks-checklist 质量与勾选真实性）；复核报告结论为 **PASS（0 阻断性问题，0 需修复项）**
  - [x] SubTask 3.2: 修复独立复核发现的阻断性问题（如有）并复验
    - 执行记录：无阻断性问题需修复（复核 0 blocker / 0 需修复项）；无需复验轮次
  - [x] SubTask 3.3: tasks.md / checklist.md 全部勾选附执行记录；变更保持未提交，待用户最终确认后手动提交
    - 执行记录：tasks.md 3 任务 / 7 子任务全部勾选并附执行记录；checklist.md 14 检查点全部勾选；git status --porcelain 复验——工作区仅本 spec 目录 `.trae/specs/phase15_03_data_model_validation_derivation_design/` untracked，零代码 / 零迁移文件 / 零 proto / 零根级改动；全部变更保持未提交，待用户最终确认后手动提交

# Task Dependencies

- Task 1（三件套产出）为本 spec 获用户批准后的首个执行项
- Task 2 depends on Task 1
- Task 3 depends on Task 2
- 后续：phase15-06 depends on 本 spec + phase15-04（dev_plan §5）；phase15-04 错误语义设计消费本 spec 错误码总表
