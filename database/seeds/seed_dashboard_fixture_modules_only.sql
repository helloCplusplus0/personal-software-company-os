-- seed_dashboard_fixture_modules_only.sql
-- phase05-12 验收 fixture 2：modules-only（仅有模块，无产品 / 仓库 / 决策）
--
-- 用途：
--   建立 Dashboard CTA 3 验收状态：module_count > 0 && product_count = 0。
--   主 CTA 指向 Product Registry / Create。
--
-- 数据状态（phase05-09 spec §"fixture 2 — modules-only"）：
--   - 至少 2 条 modules + 2 条 module_releases
--   - products / repositories / decisions / decision_links /
--     product_modules / product_repositories / module_repositories 全部为空
--   - DashboardOverviewRead 返回 module_count >= 2 / product_count=0 / repository_count=0 / decision_count=0
--
-- 依赖：reset_dashboard_acceptance.sh --fixture modules-only 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING，可重复执行。
-- 使用方式：
--   ./database/scripts/reset_dashboard_acceptance.sh --fixture modules-only

BEGIN;

-- ============================================================================
-- 模块（2 条）
-- ============================================================================
INSERT INTO modules (name, description, status) VALUES
  ('auth-service', 'Dashboard fixture：认证服务模块', 'active'),
  ('integration-test-module', 'Dashboard fixture：联调验证用模块', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 版本（3 条）：auth-service 1.0.0 / 1.1.0，integration-test-module 0.1.0
-- ============================================================================
INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '1.0.0', 'active', '2026-07-01T00:00:00+00:00'
FROM modules m WHERE m.name = 'auth-service'
ON CONFLICT (module_id, version) DO NOTHING;

INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '1.1.0', 'active', '2026-07-10T00:00:00+00:00'
FROM modules m WHERE m.name = 'auth-service'
ON CONFLICT (module_id, version) DO NOTHING;

INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '0.1.0', 'active', '2026-08-05T15:00:00+00:00'
FROM modules m WHERE m.name = 'integration-test-module'
ON CONFLICT (module_id, version) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'module_releases', COUNT(*) FROM module_releases
UNION ALL SELECT 'products', COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions;
