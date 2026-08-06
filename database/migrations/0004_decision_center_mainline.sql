-- 0004_decision_center_mainline.sql
-- Decision Center 主线：decisions 表从 phase02 只读前提原位升级为结构化主线
--
-- 上游规格：
--   - phase03-09 spec §"decisions 表原位升级 migration 必须冻结"
--   - phase03-10 decision_center_spec_v0.1.md §5.7 decisions 表演进兼容基线
--
-- 设计原则：
--   - 通过 ALTER TABLE decisions ADD COLUMN 原位升级，不新建替代表，不临时双写
--   - 不删除原有 title / created_at 字段，不破坏既有 decision_links 外键引用
--   - alternatives 使用 TEXT[]（PostgreSQL 原生数组），不用 JSONB
--   - 现有示例 Decision 数据通过兼容回填完成升级，不依赖手工 SQL 修补
--
-- 幂等性说明：
--   - 本 migration 由 schema_migrations 跟踪机制保证只执行一次
--   - 回填 UPDATE 通过 WHERE <新字段> IS NULL 保证可重复执行（即便 migration 重跑也安全）

-- ============================================================================
-- 1. 无 DEFAULT 必填字段：按"ADD COLUMN 允许 NULL → 回填 → SET NOT NULL"三步流程
--    避免 NOT NULL 无默认值在已有数据行上失败
-- ============================================================================

-- context（决策上下文，必填）
ALTER TABLE decisions ADD COLUMN IF NOT EXISTS context TEXT;
ALTER TABLE decisions ADD COLUMN IF NOT EXISTS problem TEXT;
ALTER TABLE decisions ADD COLUMN IF NOT EXISTS choice TEXT;
ALTER TABLE decisions ADD COLUMN IF NOT EXISTS reason TEXT;

-- ============================================================================
-- 2. 有 DEFAULT 字段：直接 ADD COLUMN ... NOT NULL DEFAULT ...
--    已有行由 DEFAULT 自动填充，回填 UPDATE 对它们为 no-op（但保留以保证幂等）
-- ============================================================================

ALTER TABLE decisions ADD COLUMN IF NOT EXISTS alternatives TEXT[] NOT NULL DEFAULT '{}';
ALTER TABLE decisions ADD COLUMN IF NOT EXISTS impact TEXT NOT NULL DEFAULT '';
ALTER TABLE decisions ADD COLUMN IF NOT EXISTS status TEXT NOT NULL DEFAULT 'proposed';

-- ============================================================================
-- 3. 现有示例数据兼容回填
--    对无 DEFAULT 的必填字段（context / problem / choice / reason）回填占位文本
--    对有 DEFAULT 的字段（alternatives / impact / status）回填为 no-op，但保留以保证幂等
--    WHERE 条件保证可重复执行
-- ============================================================================
UPDATE decisions
SET context     = '（历史决策，phase03 升级前无结构化上下文）',
    problem     = '（历史决策，phase03 升级前无结构化上下文）',
    choice      = '（历史决策，phase03 升级前无结构化上下文）',
    reason      = '（历史决策，phase03 升级前无结构化上下文）',
    alternatives = COALESCE(alternatives, '{}'),
    impact      = COALESCE(impact, ''),
    status      = COALESCE(NULLIF(status, ''), 'proposed')
WHERE context IS NULL
   OR problem IS NULL
   OR choice IS NULL
   OR reason IS NULL;

-- ============================================================================
-- 4. 对无 DEFAULT 必填字段补 SET NOT NULL（回填完成后可安全加约束）
-- ============================================================================
ALTER TABLE decisions ALTER COLUMN context SET NOT NULL;
ALTER TABLE decisions ALTER COLUMN problem SET NOT NULL;
ALTER TABLE decisions ALTER COLUMN choice SET NOT NULL;
ALTER TABLE decisions ALTER COLUMN reason SET NOT NULL;

-- ============================================================================
-- 5. status CHECK 约束（phase03-10 §5.6 冻结枚举）
-- ============================================================================
ALTER TABLE decisions DROP CONSTRAINT IF EXISTS decisions_status_check;
ALTER TABLE decisions ADD CONSTRAINT decisions_status_check
    CHECK (status IN ('proposed', 'active', 'superseded', 'archived'));

-- ============================================================================
-- 6. 列表读取性能索引：按状态过滤 + 创建时间倒序（phase03-10 §5.7）
-- ============================================================================
CREATE INDEX IF NOT EXISTS idx_decisions_status_created_at
    ON decisions (status, created_at DESC);
