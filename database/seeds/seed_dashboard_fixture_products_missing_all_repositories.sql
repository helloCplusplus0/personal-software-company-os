-- seed_dashboard_fixture_products_missing_all_repositories.sql
-- phase05-12 验收 fixture 8：products-missing-all-repositories（CTA 4 扩展）
--
-- 用途：
--   建立 Dashboard CTA 4 验收状态：module_count > 0 && product_count > 0 && repository_count = 0。
--   主 CTA 指向 Repository Binding / Create。
--
-- 数据状态（phase05-09 spec §"fixture 8 — products-missing-all-repositories"）：
--   - 至少 1 条 module 与至少 2 条 products
--   - 目标 products 在 product_modules 中有记录
--   - repositories 表为空，product_repositories 与 module_repositories 为空
--   - DashboardOverviewRead 返回 module_count > 0 / product_count > 0 / repository_count = 0
--
-- 依赖：reset_dashboard_acceptance.sh --fixture products-missing-all-repositories 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING，可重复执行。
-- 使用方式：
--   ./database/scripts/reset_dashboard_acceptance.sh --fixture products-missing-all-repositories

BEGIN;

-- ============================================================================
-- 模块（1 条）
-- ============================================================================
INSERT INTO modules (name, description, status) VALUES
  ('auth-service', 'Dashboard fixture：认证服务模块', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（2 条）：均已绑定模块但系统无仓库
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'Dashboard fixture：已绑定模块但系统无仓库', 'active'),
  ('Product B', 'Dashboard fixture：已绑定模块但系统无仓库', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- product_modules（2 条）：两个产品均绑定 auth-service
-- ============================================================================
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'auth-service'
  AND p.name IN ('Product A', 'Product B')
ON CONFLICT (module_id, product_id) DO NOTHING;

-- 注意：不插入 repositories / product_repositories / module_repositories
-- repositories 表保持为空，repository_count = 0

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'products', COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'product_repositories', COUNT(*) FROM product_repositories
UNION ALL SELECT 'module_repositories', COUNT(*) FROM module_repositories;
