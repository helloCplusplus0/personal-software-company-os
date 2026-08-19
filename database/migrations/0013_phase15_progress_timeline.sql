-- 0013_phase15_progress_timeline.sql — Project Progress Timeline 事件流建表迁移（phase15-06）
--
-- 上游规格：phase15-03 数据模型与校验派生设计 Spec（本 DDL 逐字源）
-- 幂等：CREATE TABLE / INDEX IF NOT EXISTS，整文件可安全重放（沿 0011 第一段模式）
-- 登记：落入 database/migrations/ 即被 RunMigrations 按文件名升序自动登记执行
--       （phase14-07 OBS-01 修复后机制，无需手工登记）

-- progress_events：repository 锚定的三轨 append-only 推进事件流（裁决②③）
-- 无 Update 语义（裁决⑨）；"当前 phase"等为读取侧派生值，不落库（裁决③）
CREATE TABLE IF NOT EXISTS progress_events (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 锚点：进度事实唯一归属仓库；RESTRICT 保护 append-only 历史
    -- （仓库被事件引用时不可物理删除，沿 0006 product_repositories FK 惯例）
    repository_id   UUID NOT NULL REFERENCES repositories(id) ON DELETE RESTRICT,
    -- 三轨 workflow（对齐 docs/ 三目录与三推进链，裁决④）
    workflow_type   TEXT NOT NULL CHECK (workflow_type IN ('phase', 'audit', 'fix')),
    -- 事件类型（裁决⑤：audit/fix 轨禁止 phase 边界标记——应用层规则 7 承接，DB 不建组合约束）
    event_kind      TEXT NOT NULL CHECK (event_kind IN ('phase_started', 'phase_completed', 'task_completed', 'note')),
    -- 任务项标识（可空；格式随 workflow×kind 矩阵变化，应用层 K-1~K-5 承接）
    task_key        TEXT NULL,
    -- 一句话标题（非空上限 200 字符，应用层承接）
    title           TEXT NOT NULL,
    -- 展开说明（可空上限 2000 字符，应用层承接）
    detail          TEXT NULL,
    -- 证据导航引用（/ 或 https:// 前缀，应用层承接；正文零托管，裁决⑦）
    evidence_ref    TEXT NULL,
    -- 来源（预留 manual/git/agent 三值；本阶段创建入口仅 manual，裁决⑧）
    source          TEXT NOT NULL CHECK (source IN ('manual', 'git', 'agent')) DEFAULT 'manual',
    -- 用户声明发生时间（允许补录历史，与 created_at 分离）
    occurred_at     TIMESTAMPTZ NOT NULL,
    -- 系统录入时间
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 时间轴回看与派生计算统一读取序索引（三键链前两键 + repository 过滤；
-- 最终 tiebreak id DESC 由 ORDER BY 补齐，不重复入索引——沿 shared_baseline §3.2 冻结形态）
CREATE INDEX IF NOT EXISTS idx_progress_events_repository_sort
    ON progress_events (repository_id, occurred_at DESC, created_at DESC);
