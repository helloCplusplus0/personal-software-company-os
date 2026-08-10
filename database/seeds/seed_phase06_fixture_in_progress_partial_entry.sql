-- seed_phase06_fixture_in_progress_partial_entry.sql
-- phase06-14 验收 fixture 2：in-progress-partial-entry
--
-- 用途：
--   建立 phase06 部分录入状态：四类首轮对象中存在 1 条但未满四类。
--   first_run_state 必须为 in_progress，根级默认进入路径必须为 /dashboard。
--   Dashboard 必须可重复验证 Continue Onboarding -> /onboarding。
--
-- 数据状态（phase06-11 spec §"in-progress-partial-entry"）：
--   - 四类首轮对象中至少存在 1 条、但未满四类对象最小持久化集合
--   - first_run_state 必须为 in_progress
--   - Onboarding 必须自动定位到第一个未完成步骤
--
-- 依赖：reset_phase06_acceptance.sh --fixture in-progress-partial-entry 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING，可重复执行。

BEGIN;

-- ============================================================================
-- 模块（1 条）：已完成 Product 步骤、Module 步骤，但 Repository / Decision 未完成
-- ============================================================================
INSERT INTO modules (name, description, status, capability_key) VALUES
  ('auth-service', 'phase06 fixture：认证服务模块', 'active', 'auth')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（1 条）：Product 步骤已完成
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'phase06 fixture：部分录入产品', 'active')
ON CONFLICT (name) DO NOTHING;

-- 不插入 repositories / decisions → first_run_state = in_progress
-- current_step 应推导为 repository（Product 已完成，Repository 未完成）

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'products' AS tbl, COUNT(*) AS cnt FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'modules', COUNT(*) FROM modules
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions;
