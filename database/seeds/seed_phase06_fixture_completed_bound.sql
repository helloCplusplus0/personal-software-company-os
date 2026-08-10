-- seed_phase06_fixture_completed_bound.sql
-- phase06-14 验收 fixture 4：completed-bound
--
-- 用途：
--   建立 phase06 首轮完成且核心绑定补全状态：延续 completed-unbound 的四类最小持久化对象，
--   同时补齐当前阶段冻结的核心绑定关系。
--   该 fixture 作为数据主权闭合、导出覆盖矩阵与复用感知验证的正式前置状态。
--
-- 数据状态（phase06-11 spec §"completed-bound"）：
--   - 延续 completed-unbound 的四类最小持久化对象
--   - 同时补齐核心绑定关系：
--     * product_modules
--     * product_repositories
--     * module_repositories
--     * decision_links（需要时）
--   - first_run_state 必须为 completed
--   - 不得要求验收人员在 completed-unbound 基础上手工补数据
--
-- 依赖：reset_phase06_acceptance.sh --fixture completed-bound 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING + WHERE NOT EXISTS，可重复执行。

BEGIN;

-- ============================================================================
-- 模块（2 条）：auth-service / integration-test-module
-- ============================================================================
INSERT INTO modules (name, description, status, capability_key) VALUES
  ('auth-service', 'phase06 fixture：首轮完成已绑定模块（认证服务）', 'active', 'auth'),
  ('integration-test-module', 'phase06 fixture：首轮完成已绑定模块（集成测试）', 'active', 'web_frontend')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（2 条）：Product A / Product B
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'phase06 fixture：首轮完成已绑定产品 A', 'active'),
  ('Product B', 'phase06 fixture：首轮完成已绑定产品 B', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 仓库（1 条）：main-repo
-- ============================================================================
INSERT INTO repositories (name, url, provider, status) VALUES
  ('main-repo', 'https://github.com/psco/main-repo', 'github', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- Decision（1 条）：proposed 状态
-- ============================================================================
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'phase06 completed-bound 验收决策',
       'phase06 fixture：首轮完成已绑定的决策上下文',
       'phase06 fixture：首轮完成已绑定的决策问题',
       '{}',
       'phase06 fixture 选择',
       'phase06 fixture 理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'phase06 completed-bound 验收决策');

-- ============================================================================
-- 核心绑定关系补齐
-- ============================================================================

-- product_modules：auth-service / integration-test-module -> Product A
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND p.name = 'Product A'
ON CONFLICT (module_id, product_id) DO NOTHING;

-- product_repositories：Product A -> main-repo
INSERT INTO product_repositories (product_id, repository_id)
SELECT p.id, r.id
FROM products p
CROSS JOIN repositories r
WHERE p.name = 'Product A'
  AND r.name = 'main-repo'
ON CONFLICT (product_id, repository_id) DO NOTHING;

-- module_repositories：auth-service / integration-test-module -> main-repo
INSERT INTO module_repositories (module_id, repository_id)
SELECT m.id, r.id
FROM modules m
CROSS JOIN repositories r
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND r.name = 'main-repo'
ON CONFLICT (module_id, repository_id) DO NOTHING;

-- decision_links：phase06 completed-bound 验收决策 -> auth-service
INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d
CROSS JOIN modules m
WHERE d.title = 'phase06 completed-bound 验收决策'
  AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'products' AS tbl, COUNT(*) AS cnt FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'modules', COUNT(*) FROM modules
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'product_repositories', COUNT(*) FROM product_repositories
UNION ALL SELECT 'module_repositories', COUNT(*) FROM module_repositories
UNION ALL SELECT 'decision_links', COUNT(*) FROM decision_links;

-- 绑定完整性验证：所有核心绑定关系必须存在
SELECT 'binding_completeness' AS check_name,
       (SELECT COUNT(*) FROM product_modules) AS pm_count,
       (SELECT COUNT(*) FROM product_repositories) AS pr_count,
       (SELECT COUNT(*) FROM module_repositories) AS mr_count,
       (SELECT COUNT(*) FROM decision_links) AS dl_count;
