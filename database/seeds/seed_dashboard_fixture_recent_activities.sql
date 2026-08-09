-- seed_dashboard_fixture_recent_activities.sql
-- phase05-12 验收 fixture 7：recent-activities（系统已就绪中性状态，CTA 9）
--
-- 用途：
--   建立 Dashboard CTA 9 验收状态：无缺口且有活动数据 → 系统已就绪中性状态。
--   所有 product 已完成双绑定，不命中任何缺口 CTA。
--   不存在 pending_decision_signals。
--   RecentActivityRead 返回覆盖 8 类活动类型的近期数据，按 activity_at 倒序，最多 10 条。
--
-- 数据状态（phase05-09 spec §"fixture 7 — recent-activities"）：
--   - 覆盖 8 类活动类型：module / release / product / repository / decision /
--     product_module_binding / product_repository_binding / module_repository_binding
--   - 所有 product 已完成模块与仓库双绑定
--   - 不存在 pending_decision_signals（所有 decision 为非 proposed 状态）
--   - 所有活动项带显式时间字段，时间分布在可排序的近期范围内
--   - 最近活动期望排序（DESC）：
--       1. module_repository_binding  2026-08-08T15:00:00Z
--       2. product_repository_binding 2026-08-08T14:00:00Z
--       3. product_module_binding     2026-08-08T13:00:00Z
--       4. decision                   2026-08-08T12:00:00Z
--       5. repository                 2026-08-08T11:00:00Z
--       6. product                    2026-08-08T10:00:00Z
--       7. release                    2026-08-08T09:00:00Z
--       8. module                     2026-08-08T08:00:00Z
--
-- 依赖：reset_dashboard_acceptance.sh --fixture recent-activities 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING + WHERE NOT EXISTS，可重复执行。
-- 使用方式：
--   ./database/scripts/reset_dashboard_acceptance.sh --fixture recent-activities

BEGIN;

-- ============================================================================
-- 模块（1 条）
-- ============================================================================
INSERT INTO modules (name, description, status, created_at) VALUES
  ('auth-service', 'Dashboard fixture：认证服务模块', 'active', '2026-08-08T08:00:00+00:00')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 版本（1 条）：release 活动类型
-- ============================================================================
INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '1.0.0', 'active', '2026-08-08T09:00:00+00:00'
FROM modules m WHERE m.name = 'auth-service'
ON CONFLICT (module_id, version) DO NOTHING;

-- ============================================================================
-- 产品（1 条）：所有产品已完成双绑定
-- ============================================================================
INSERT INTO products (name, description, status, created_at) VALUES
  ('Product A', 'Dashboard fixture：已完整绑定的产品', 'active', '2026-08-08T10:00:00+00:00')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 仓库（1 条）
-- ============================================================================
INSERT INTO repositories (name, url, provider, status, created_at) VALUES
  ('main-repo', 'https://github.com/psco/main-repo', 'github', 'active', '2026-08-08T11:00:00+00:00')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- Decision（1 条）：非 proposed 状态，不产生 pending_decision_signals
-- ============================================================================
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status, created_at)
SELECT 'Dashboard fixture：日志采集方案决策',
       'Dashboard fixture 验证上下文',
       'Dashboard fixture 验证问题',
       '{}',
       'Dashboard fixture 验证选择',
       'Dashboard fixture 验证理由',
       '',
       'active',
       '2026-08-08T12:00:00+00:00'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'Dashboard fixture：日志采集方案决策');

-- ============================================================================
-- product_modules（1 条）：product_module_binding 活动类型
-- 产品完成模块绑定
-- ============================================================================
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'auth-service'
  AND p.name = 'Product A'
ON CONFLICT (module_id, product_id) DO NOTHING;

UPDATE product_modules pm
SET created_at = '2026-08-08T13:00:00+00:00'
FROM modules m, products p
WHERE pm.module_id = m.id
  AND pm.product_id = p.id
  AND m.name = 'auth-service'
  AND p.name = 'Product A';

-- ============================================================================
-- product_repositories（1 条）：product_repository_binding 活动类型
-- 产品完成仓库绑定（双绑定完成）
-- ============================================================================
INSERT INTO product_repositories (product_id, repository_id)
SELECT p.id, r.id
FROM products p
CROSS JOIN repositories r
WHERE p.name = 'Product A'
  AND r.name = 'main-repo'
ON CONFLICT (product_id, repository_id) DO NOTHING;

UPDATE product_repositories pr
SET created_at = '2026-08-08T14:00:00+00:00'
FROM products p, repositories r
WHERE pr.product_id = p.id
  AND pr.repository_id = r.id
  AND p.name = 'Product A'
  AND r.name = 'main-repo';

-- ============================================================================
-- module_repositories（1 条）：module_repository_binding 活动类型
-- ============================================================================
INSERT INTO module_repositories (module_id, repository_id)
SELECT m.id, r.id
FROM modules m
CROSS JOIN repositories r
WHERE m.name = 'auth-service'
  AND r.name = 'main-repo'
ON CONFLICT (module_id, repository_id) DO NOTHING;

UPDATE module_repositories mr
SET created_at = '2026-08-08T15:00:00+00:00'
FROM modules m, repositories r
WHERE mr.module_id = m.id
  AND mr.repository_id = r.id
  AND m.name = 'auth-service'
  AND r.name = 'main-repo';

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'module_releases', COUNT(*) FROM module_releases
UNION ALL SELECT 'products', COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'product_repositories', COUNT(*) FROM product_repositories
UNION ALL SELECT 'module_repositories', COUNT(*) FROM module_repositories;

-- 验证：无 pending decisions
SELECT 'pending_decisions' AS check_name, COUNT(*) AS cnt
FROM decisions WHERE status = 'proposed';

-- 验证：所有产品已完整绑定
SELECT p.name,
       EXISTS(SELECT 1 FROM product_modules pm WHERE pm.product_id = p.id) AS has_module,
       EXISTS(SELECT 1 FROM product_repositories pr WHERE pr.product_id = p.id) AS has_repository
FROM products p
ORDER BY p.name;
