-- seed_dashboard_fixture_products_without_modules.sql
-- phase05-12 验收 fixture 3：products-without-modules（有产品无模块，非空缺口）
--
-- 用途：
--   建立 Dashboard CTA 2 验收状态：非空缺口 → Module Registry / Create。
--   module_count=0 && product_count >= 2，主 CTA 指向 Module Registry / Create。
--
-- 数据状态（phase05-09 spec §"fixture 3 — products-without-modules"）：
--   - 至少 2 条 products
--   - modules / module_releases / product_modules / module_repositories /
--     product_repositories / decision_links 必须为空
--   - repositories 与 decisions 允许保留 readonly prereqs 占位数据（本 fixture 保持为空）
--   - DashboardOverviewRead 返回 module_count=0 / product_count >= 2
--
-- 依赖：reset_dashboard_acceptance.sh --fixture products-without-modules 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING，可重复执行。
-- 使用方式：
--   ./database/scripts/reset_dashboard_acceptance.sh --fixture products-without-modules

BEGIN;

-- ============================================================================
-- 产品（3 条）：无模块绑定的产品
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'Dashboard fixture：无模块绑定的产品 A', 'active'),
  ('Product B', 'Dashboard fixture：无模块绑定的产品 B', 'active'),
  ('Product C', 'Dashboard fixture：无模块绑定的产品 C', 'active')
ON CONFLICT (name) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'products', COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules;
