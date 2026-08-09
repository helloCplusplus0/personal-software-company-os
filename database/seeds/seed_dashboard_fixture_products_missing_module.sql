-- seed_dashboard_fixture_products_missing_module.sql
-- phase05-12 验收 fixture 5：products-missing-module（产品缺模块绑定，CTA 8）
--
-- 用途：
--   建立 Dashboard CTA 8 验收状态：product missing module binding → Product Detail。
--   FeedbackSignalRead 产出 PRODUCT_MISSING_MODULE_BINDING 信号。
--
-- 数据状态（phase05-09 spec §"fixture 5 — products-missing-module"）：
--   - 1 条目标 product：已绑定 repository 但未绑定 module
--   - 1 条对照 product：已完整绑定（module + repository）
--   - 不存在 pending_decision_signals
--   - DashboardOverviewRead 的 product_with_module_count < product_count
--   - FeedbackSignalRead 产出 PRODUCT_MISSING_MODULE_BINDING 信号
--
-- 依赖：reset_dashboard_acceptance.sh --fixture products-missing-module 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING，可重复执行。
-- 使用方式：
--   ./database/scripts/reset_dashboard_acceptance.sh --fixture products-missing-module

BEGIN;

-- ============================================================================
-- 模块（1 条）
-- ============================================================================
INSERT INTO modules (name, description, status) VALUES
  ('auth-service', 'Dashboard fixture：认证服务模块', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（2 条）：Product A（目标，缺模块）/ Product B（对照，完整绑定）
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'Dashboard fixture：目标产品，已绑定仓库但缺模块', 'active'),
  ('Product B', 'Dashboard fixture：对照产品，已完整绑定', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 仓库（1 条）
-- ============================================================================
INSERT INTO repositories (name, url, provider, status) VALUES
  ('main-repo', 'https://github.com/psco/main-repo', 'github', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- product_repositories（2 条）：两个产品均绑定 main-repo
-- ============================================================================
INSERT INTO product_repositories (product_id, repository_id)
SELECT p.id, r.id
FROM products p
CROSS JOIN repositories r
WHERE p.name IN ('Product A', 'Product B')
  AND r.name = 'main-repo'
ON CONFLICT (product_id, repository_id) DO NOTHING;

-- ============================================================================
-- product_modules（1 条）：仅对照产品 Product B 绑定 auth-service
-- 目标产品 Product A 不绑定模块，产生 PRODUCT_MISSING_MODULE_BINDING 缺口
-- ============================================================================
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'auth-service'
  AND p.name = 'Product B'
ON CONFLICT (module_id, product_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'products', COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'product_repositories', COUNT(*) FROM product_repositories;

-- 目标产品缺口验证
SELECT p.name,
       EXISTS(SELECT 1 FROM product_modules pm WHERE pm.product_id = p.id) AS has_module,
       EXISTS(SELECT 1 FROM product_repositories pr WHERE pr.product_id = p.id) AS has_repository
FROM products p
ORDER BY p.name;
