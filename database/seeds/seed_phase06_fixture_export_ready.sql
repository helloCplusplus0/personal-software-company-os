-- seed_phase06_fixture_export_ready.sql
-- phase06-14 验收 fixture 5：export-ready
--
-- 用途：
--   建立 phase06 数据主权导出闭合状态：至少满足 completed-bound，
--   且具备 Export 当前阶段最小覆盖矩阵所需的 9 类 canonical 数据集。
--   验收时可直接验证 GET /api/dashboard/export 读取 export_snapshot，
--   POST /api/dashboard/export 生成正式导出结果。
--
-- 数据状态（phase06-11 spec §"export-ready"）：
--   - 至少满足 completed-bound（四类 canonical 对象 + 核心绑定关系）
--   - 具备 9 类 canonical 数据集：
--     1. products
--     2. modules
--     3. releases (module_releases)
--     4. repositories
--     5. decisions
--     6. decision_links
--     7. product_modules
--     8. product_repositories
--     9. module_repositories
--   - 不得缺失任何核心绑定关系后仍被判定为"数据主权闭合"
--
-- 依赖：reset_phase06_acceptance.sh --fixture export-ready 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING + WHERE NOT EXISTS，可重复执行。

BEGIN;

-- ============================================================================
-- 模块（2 条）：auth-service / integration-test-module
-- ============================================================================
INSERT INTO modules (name, description, status, capability_key) VALUES
  ('auth-service', 'phase06 fixture：导出就绪模块（认证服务）', 'active', 'auth'),
  ('integration-test-module', 'phase06 fixture：导出就绪模块（集成测试）', 'active', 'web_frontend')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（2 条）：Product A / Product B
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'phase06 fixture：导出就绪产品 A', 'active'),
  ('Product B', 'phase06 fixture：导出就绪产品 B', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 仓库（1 条）：main-repo
-- ============================================================================
INSERT INTO repositories (name, url, provider, status) VALUES
  ('main-repo', 'https://github.com/psco/main-repo', 'github', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- module_releases（2 条）：auth-service 1.0.0 / integration-test-module 0.1.0
-- ============================================================================
-- 第 3 类核心资产：releases
INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '1.0.0', 'active', '2026-07-01T00:00:00+00:00'
FROM modules m WHERE m.name = 'auth-service'
ON CONFLICT (module_id, version) DO NOTHING;

INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '0.1.0', 'active', '2026-08-05T00:00:00+00:00'
FROM modules m WHERE m.name = 'integration-test-module'
ON CONFLICT (module_id, version) DO NOTHING;

-- ============================================================================
-- Decision（2 条）：proposed / active 状态
-- ============================================================================
-- 第 5 类核心资产：decisions
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'phase06 export-ready 验收决策 1',
       'phase06 fixture：导出就绪决策上下文',
       'phase06 fixture：导出就绪决策问题',
       '{}',
       'phase06 fixture 选择',
       'phase06 fixture 理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'phase06 export-ready 验收决策 1');

INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'phase06 export-ready 验收决策 2',
       'phase06 fixture：导出就绪第二决策上下文',
       'phase06 fixture：导出就绪第二决策问题',
       '{"选项 A", "选项 B"}',
       '选项 A',
       'phase06 fixture 理由',
       'phase06 fixture 影响',
       'active'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'phase06 export-ready 验收决策 2');

-- ============================================================================
-- 核心绑定关系补齐（第 6-9 类核心资产）
-- ============================================================================

-- 第 7 类：product_modules
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND p.name = 'Product A'
ON CONFLICT (module_id, product_id) DO NOTHING;

-- 第 8 类：product_repositories
INSERT INTO product_repositories (product_id, repository_id)
SELECT p.id, r.id
FROM products p
CROSS JOIN repositories r
WHERE p.name = 'Product A'
  AND r.name = 'main-repo'
ON CONFLICT (product_id, repository_id) DO NOTHING;

-- 第 9 类：module_repositories
INSERT INTO module_repositories (module_id, repository_id)
SELECT m.id, r.id
FROM modules m
CROSS JOIN repositories r
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND r.name = 'main-repo'
ON CONFLICT (module_id, repository_id) DO NOTHING;

-- 第 6 类：decision_links
INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d
CROSS JOIN modules m
WHERE d.title = 'phase06 export-ready 验收决策 1'
  AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验：9 类核心资产覆盖矩阵验证
-- ============================================================================
SELECT 'export_coverage_matrix' AS check_name,
       (SELECT COUNT(*) FROM products) AS products,
       (SELECT COUNT(*) FROM modules) AS modules,
       (SELECT COUNT(*) FROM module_releases) AS releases,
       (SELECT COUNT(*) FROM repositories) AS repositories,
       (SELECT COUNT(*) FROM decisions) AS decisions,
       (SELECT COUNT(*) FROM decision_links) AS decision_links,
       (SELECT COUNT(*) FROM product_modules) AS product_modules,
       (SELECT COUNT(*) FROM product_repositories) AS product_repositories,
       (SELECT COUNT(*) FROM module_repositories) AS module_repositories;

-- 9 类全部 > 0 才满足 Export 最小覆盖矩阵
SELECT 'export_ready' AS check_name,
       CASE
         WHEN (SELECT COUNT(*) FROM products) > 0
          AND (SELECT COUNT(*) FROM modules) > 0
          AND (SELECT COUNT(*) FROM module_releases) > 0
          AND (SELECT COUNT(*) FROM repositories) > 0
          AND (SELECT COUNT(*) FROM decisions) > 0
          AND (SELECT COUNT(*) FROM decision_links) > 0
          AND (SELECT COUNT(*) FROM product_modules) > 0
          AND (SELECT COUNT(*) FROM product_repositories) > 0
          AND (SELECT COUNT(*) FROM module_repositories) > 0
         THEN 'ready'
         ELSE 'not_ready'
       END AS export_coverage_status;
