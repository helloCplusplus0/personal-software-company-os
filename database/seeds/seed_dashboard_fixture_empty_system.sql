-- seed_dashboard_fixture_empty_system.sql
-- phase05-12 验收 fixture 1：empty-system（冷启动空系统）
--
-- 用途：
--   建立 Dashboard 冷启动空系统验收状态。所有 Dashboard 相关表为空。
--   对应 CTA 1：冷启动空系统 → Module Registry / Create。
--
-- 数据状态（phase05-09 spec §"fixture 1 — empty-system"）：
--   - modules / module_releases / products / repositories / decisions /
--     decision_links / product_modules / product_repositories / module_repositories 全部为空
--   - DashboardOverviewRead 返回 module_count=0 / product_count=0 / repository_count=0 / decision_count=0
--
-- 依赖：reset_dashboard_acceptance.sh --fixture empty-system 已执行清空操作
-- 幂等：本文件不插入数据，只包含校验查询，天然幂等。
-- 使用方式：
--   ./database/scripts/reset_dashboard_acceptance.sh --fixture empty-system

-- ============================================================================
-- 校验：所有 Dashboard 相关表必须为空
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'module_releases', COUNT(*) FROM module_releases
UNION ALL SELECT 'products', COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions
UNION ALL SELECT 'decision_links', COUNT(*) FROM decision_links
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'product_repositories', COUNT(*) FROM product_repositories
UNION ALL SELECT 'module_repositories', COUNT(*) FROM module_repositories;
