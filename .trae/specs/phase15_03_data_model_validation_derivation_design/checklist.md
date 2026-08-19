# phase15-03 Checklist

- [x] DDL 级设计冻结：progress_events 11 列逐字段（id / repository_id / workflow_type / event_kind / task_key / title / detail / evidence_ref / source / occurred_at / created_at）与 shared_baseline §3.2 字段矩阵一致；TEXT + CHECK(IN ...) 枚举三列值域一致；索引 (repository_id, occurred_at DESC, created_at DESC) 为 §3.2 冻结形态；可逐字转写为 0013 迁移
- [x] DDL 5 项设计决策冻结且理由留档：FK ON DELETE RESTRICT（沿 0006 惯例 + 保护 append-only 历史）/ DB 只承接单列枚举完整性（组合与格式归应用层，沿 standard 模式）/ 索引不含 id DESC（ORDER BY 补齐）/ 无 updated_at 列（无更新语义）/ 无数据迁移段（纯新建）
- [x] 幂等 DDL 模式确认（CREATE TABLE / INDEX IF NOT EXISTS 沿 0011 第一段）；迁移落点 database/migrations/0013_phase15_progress_timeline.sql；自动登记机制声明（RunMigrations 文件名升序，无需手工登记）
- [x] 9 条校验规则形式化：9 个业务错误码总表（INVALID_WORKFLOW_TYPE / INVALID_EVENT_KIND / EVENT_KIND_NOT_ALLOWED / TASK_KEY_REQUIRED / TASK_KEY_FORMAT_INVALID / INVALID_TITLE / INVALID_DETAIL / INVALID_EVIDENCE_REF / INVALID_SOURCE）每码含判定逻辑伪码，与 phase15-02 合法矩阵 12 格 + K-1~K-5 正则 1:1 对应
- [x] V8 无独立错误码的结构性说明留档（允许规则不产生错误，语义由 V2-V6 矩阵分支覆盖）；2 个 envelope 前置码（INVALID_REPOSITORY_ID 格式层 / INVALID_OCCURRED_AT 已设置）显式区分于 9 条业务规则，repository 存在性承接位归 phase15-04 DP-2 未偷渡
- [x] 执行序 6 步冻结（envelope → V1a → V1b → V7 → task_key 矩阵分支〔先必填后格式〕→ V9 文本顺序 → repository 存在性）；报第一个错误策略沿 standard 模式；错误信息格式 %w: [CODE] message
- [x] TrimSpace 规范化边界冻结（task_key trim 后持久化 / title·detail·evidence_ref 原值持久化 / 长度 rune 计量沿 maxSummaryRunes 模式）
- [x] 派生实现序冻结：SQL ORDER BY 为唯一排序执行位（Go 不重排）；repository 单一查询形态（repository_id 必选 + workflow_type 可选过滤）供 List / brief / 派生共用；service 纯函数 DeriveProgressSummary 零 I/O 可纯单元测试
- [x] 三派生项精确算法冻结：recent_events min(10, len) 截断 / latest_task_completed 首个 task_completed / current_phase 双步（latestStartedIdx + j < latestStartedIdx 同 key phase_completed 完结判定）；DESC 切片"索引越小越晚"语义明确；空值双情形同型（零值）
- [x] tiebreak 确定性说明（PostgreSQL UUID 字节序与 Go uuid.UUID 字节比较一致；两类碰撞场景）；个人规模边界与万级优化属未来进入条件声明
- [x] 与 phase15-02 语义上游 1:1 转译零语义漂移；无 phase15-04/05/06 偷渡（无 RPC envelope / 无字段号分配 / 无 ProgressReader 签名 / 无 DP-2 裁决 / 无前端细节 / 无代码与迁移文件实体）
- [x] 独立复核通过（0 阻断；复核维度：DDL 可逐字转写性 / 校验可测试性 / 派生精确性 / 1:1 转译 / 无偷渡 / 既有模式一致性 / 勾选真实性）
- [x] tasks.md / checklist.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
