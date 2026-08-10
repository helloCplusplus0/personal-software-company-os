-- seed_phase06_fixture_reuse_latest_after_binding.sql
-- phase06-14 验收 fixture 11：reuse-latest-after-binding
--
-- 用途：
--   建立 phase06 复用感知读取最新已提交状态：在 reuse-latest 基础上额外体现一次已提交绑定变化。
--   再次读取 ReuseSummaryRead 时，必须返回更新后的 reuse_product_count、latest_reuse_at 或 capability_summary。
--   该 fixture 作为"读取时反映最新已提交状态"的正式验收证据。
--
-- 数据状态（phase06-11 spec §"reuse-latest-after-binding"）：
--   - 在 reuse-latest 基础上额外体现一次已提交绑定变化
--   - 再次读取 ReuseSummaryRead 时，必须返回更新后的 reuse_product_count / latest_reuse_at / capability_summary
--   - 该 fixture 作为"读取时反映最新已提交状态"的正式验收证据
--
-- 与 reuse-latest 的差异：
--   - reuse-latest：auth-service 被 Product A/B/C 绑定（reuse_product_count = 3）
--   - reuse-latest-after-binding：integration-test-module 额外被 Product B/C 绑定
--     （reuse_product_count 从 1 增长到 3）
--   - 这体现了"绑定补全后读取到最新状态"
--
-- 依赖：reset_phase06_acceptance.sh --fixture reuse-latest-after-binding 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING + WHERE NOT EXISTS，可重复执行。

BEGIN;

-- ============================================================================
-- 模块（2 条）：与 reuse-latest 相同
-- ============================================================================
INSERT INTO modules (name, description, status, capability_key) VALUES
  ('auth-service', 'phase06 fixture：复用感知绑定后模块（认证服务）', 'active', 'auth'),
  ('integration-test-module', 'phase06 fixture：复用感知绑定后模块（集成测试，新增绑定）', 'active', 'web_frontend')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（3 条）：Product A / Product B / Product C
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'phase06 fixture：复用感知绑定后产品 A', 'active'),
  ('Product B', 'phase06 fixture：复用感知绑定后产品 B', 'active'),
  ('Product C', 'phase06 fixture：复用感知绑定后产品 C', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 仓库（1 条）：main-repo
-- ============================================================================
INSERT INTO repositories (name, url, provider, status) VALUES
  ('main-repo', 'https://github.com/psco/main-repo', 'github', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- product_modules：体现"绑定补全后读取到最新状态"
-- ============================================================================
-- auth-service 被 Product A/B/C 绑定（reuse_product_count = 3，与 reuse-latest 一致）
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'auth-service'
  AND p.name IN ('Product A', 'Product B', 'Product C')
ON CONFLICT (module_id, product_id) DO NOTHING;

-- integration-test-module 被 Product A/B/C 绑定（reuse_product_count = 3）
-- 与 reuse-latest 的差异：reuse-latest 中 integration-test-module 仅被 Product A 绑定（count=1）
-- 本 fixture 额外绑定 Product B/C，体现"绑定补全后读取到最新状态"
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'integration-test-module'
  AND p.name IN ('Product A', 'Product B', 'Product C')
ON CONFLICT (module_id, product_id) DO NOTHING;

-- ============================================================================
-- 其他核心绑定关系（满足 completed-bound 前提）
-- ============================================================================
INSERT INTO product_repositories (product_id, repository_id)
SELECT p.id, r.id
FROM products p CROSS JOIN repositories r
WHERE p.name IN ('Product A', 'Product B') AND r.name = 'main-repo'
ON CONFLICT (product_id, repository_id) DO NOTHING;

INSERT INTO module_repositories (module_id, repository_id)
SELECT m.id, r.id
FROM modules m CROSS JOIN repositories r
WHERE m.name IN ('auth-service', 'integration-test-module') AND r.name = 'main-repo'
ON CONFLICT (module_id, repository_id) DO NOTHING;

-- ============================================================================
-- Decision + decision_links（满足四类 canonical 完整性）
-- ============================================================================
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'phase06 reuse-latest-after-binding 验收决策',
       'phase06 fixture：复用感知绑定后决策上下文',
       'phase06 fixture：复用感知绑定后决策问题',
       '{}',
       'phase06 fixture 选择',
       'phase06 fixture 理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'phase06 reuse-latest-after-binding 验收决策');

INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d CROSS JOIN modules m
WHERE d.title = 'phase06 reuse-latest-after-binding 验收决策' AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验：复用感知聚合验证（绑定补全后）
-- ============================================================================

-- module_reuse_summary：两个 Module 都被 3 个 Product 复用
-- 与 reuse-latest 的差异：integration-test-module 从 count=1 增长到 count=3
SELECT 'module_reuse_summary_after_binding' AS check_name,
       m.name AS module_name,
       COUNT(DISTINCT pm.product_id) AS reuse_product_count
FROM modules m
LEFT JOIN product_modules pm ON pm.module_id = m.id
GROUP BY m.name
ORDER BY reuse_product_count DESC;

-- capability_summary：按 capability_key 聚合
SELECT 'capability_summary_after_binding' AS check_name,
       m.capability_key,
       COUNT(*) AS module_count
FROM modules m
WHERE m.capability_key IS NOT NULL
GROUP BY m.capability_key
ORDER BY module_count DESC;
