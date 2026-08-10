-- seed_phase06_fixture_cold_start_empty.sql
-- phase06-14 验收 fixture 1：cold-start-empty
--
-- 用途：
--   建立 phase06 冷启动空系统状态：四类 canonical 对象全部不存在。
--   first_run_state 必须为 not_started，根级默认进入路径必须回落到 /onboarding。
--
-- 数据状态（phase06-11 spec §"cold-start-empty"）：
--   - Product / Repository / Module / Decision 四类首轮对象必须全部不存在
--   - first_run_state 必须为 not_started
--   - 不得伪造任何已完成首轮对象记录
--
-- 依赖：reset_phase06_acceptance.sh --fixture cold-start-empty 已执行清空操作
-- 幂等：本文件不插入数据，天然幂等。

-- cold-start-empty：清空后不加载任何 canonical 数据
-- first_run_state 由 GetFirstRunState 读时派生（四类计数均为 0 → not_started）

-- ============================================================================
-- 校验：四类 canonical 对象全部为 0
-- ============================================================================
SELECT 'products' AS tbl, COUNT(*) AS cnt FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'modules', COUNT(*) FROM modules
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions
UNION ALL SELECT 'instance_exports', COUNT(*) FROM instance_exports
UNION ALL SELECT 'instance_backups', COUNT(*) FROM instance_backups;
