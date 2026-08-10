-- seed_phase06_fixture_completed_unbound.sql
-- phase06-14 验收 fixture 3：completed-unbound
--
-- 用途：
--   建立 phase06 首轮完成但绑定缺失状态：四类 canonical 对象都已至少持久化 1 条，
--   但允许缺少 Product->Module / Product->Repository / Module->Repository / Decision->target 绑定。
--   first_run_state 必须为 completed，证明"缺少绑定关系仍完成首轮成功会话"。
--
-- 数据状态（phase06-11 spec §"completed-unbound"）：
--   - Product / Repository / Module / Decision 四类对象必须都已至少持久化 1 条
--   - first_run_state 必须为 completed
--   - 明确允许缺少以下绑定关系：
--     * Product -> Module (product_modules)
--     * Product -> Repository (product_repositories)
--     * Module -> Repository (module_repositories)
--     * Decision -> target (decision_links)
--   - 该 fixture 必须证明"缺少绑定关系仍完成首轮成功会话"
--
-- 依赖：reset_phase06_acceptance.sh --fixture completed-unbound 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING，可重复执行。

BEGIN;

-- ============================================================================
-- 模块（1 条）：auth-service，带 capability_key
-- ============================================================================
INSERT INTO modules (name, description, status, capability_key) VALUES
  ('auth-service', 'phase06 fixture：首轮完成未绑定模块', 'active', 'auth')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（1 条）：Product A
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'phase06 fixture：首轮完成未绑定产品', 'active')
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
-- decisions.title 无 UNIQUE 约束，使用 WHERE NOT EXISTS 守卫保证幂等
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'phase06 completed-unbound 验收决策',
       'phase06 fixture：首轮完成未绑定的决策上下文',
       'phase06 fixture：首轮完成未绑定的决策问题',
       '{}',
       'phase06 fixture 选择',
       'phase06 fixture 理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'phase06 completed-unbound 验收决策');

-- ============================================================================
-- 不插入任何绑定关系：
--   - product_modules（空）
--   - product_repositories（空）
--   - module_repositories（空）
--   - decision_links（空）
-- 证明"缺少绑定关系仍完成首轮成功会话"
-- ============================================================================

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

-- first_run_state 应为 completed（四类 canonical 对象各 >= 1）
SELECT 'first_run_state_check' AS check_name,
       (SELECT COUNT(*) FROM products) AS product_count,
       (SELECT COUNT(*) FROM repositories) AS repository_count,
       (SELECT COUNT(*) FROM modules) AS module_count,
       (SELECT COUNT(*) FROM decisions) AS decision_count,
       CASE
         WHEN (SELECT COUNT(*) FROM products) > 0
          AND (SELECT COUNT(*) FROM repositories) > 0
          AND (SELECT COUNT(*) FROM modules) > 0
          AND (SELECT COUNT(*) FROM decisions) > 0
         THEN 'completed'
         ELSE 'not_completed'
       END AS expected_first_run_state;
