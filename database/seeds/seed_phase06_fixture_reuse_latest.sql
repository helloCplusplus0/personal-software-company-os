-- seed_phase06_fixture_reuse_latest.sql
-- phase06-14 验收 fixture 10：reuse-latest
--
-- 用途：
--   建立 phase06 复用感知可见状态：至少 1 个 Module 被多个 Product 直接绑定，
--   且至少 1 个 Module 填写了 capability_key。
--   验收时可在 Dashboard、Module Detail 或 Product Detail 中看到非空
--   module_reuse_summary / capability_summary。
--
-- 数据状态（phase06-11 spec §"reuse-latest"）：
--   - 至少 1 个被多个 Product 直接绑定的 Module（reuse_product_count >= 2）
--   - 至少 1 个填写了 capability_key 的 Module
--   - module_reuse_summary / capability_summary 必须非空可读
--   - 不得依赖异步统计表或离线任务才能展示复用反馈
--
-- 依赖：reset_phase06_acceptance.sh --fixture reuse-latest 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING + WHERE NOT EXISTS，可重复执行。

BEGIN;

-- ============================================================================
-- 模块（2 条）：auth-service（被多 Product 复用）/ integration-test-module
-- ============================================================================
INSERT INTO modules (name, description, status, capability_key) VALUES
  ('auth-service', 'phase06 fixture：复用感知模块（认证服务，被多 Product 复用）', 'active', 'auth'),
  ('integration-test-module', 'phase06 fixture：复用感知模块（集成测试）', 'active', 'web_frontend')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 产品（3 条）：Product A / Product B / Product C
-- ============================================================================
INSERT INTO products (name, description, status) VALUES
  ('Product A', 'phase06 fixture：复用感知产品 A', 'active'),
  ('Product B', 'phase06 fixture：复用感知产品 B', 'active'),
  ('Product C', 'phase06 fixture：复用感知产品 C', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- 仓库（1 条）：main-repo
-- ============================================================================
INSERT INTO repositories (name, url, provider, status) VALUES
  ('main-repo', 'https://github.com/psco/main-repo', 'github', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- product_modules：auth-service 被 Product A / B / C 同时绑定（reuse_product_count = 3）
-- ============================================================================
-- 这是 module_reuse_summary 的核心聚合来源
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'auth-service'
  AND p.name IN ('Product A', 'Product B', 'Product C')
ON CONFLICT (module_id, product_id) DO NOTHING;

-- integration-test-module 仅被 Product A 绑定（reuse_product_count = 1）
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'integration-test-module'
  AND p.name = 'Product A'
ON CONFLICT (module_id, product_id) DO NOTHING;

-- ============================================================================
-- 其他核心绑定关系（满足 completed-bound 前提）
-- ============================================================================
INSERT INTO product_repositories (product_id, repository_id)
SELECT p.id, r.id
FROM products p CROSS JOIN repositories r
WHERE p.name = 'Product A' AND r.name = 'main-repo'
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
SELECT 'phase06 reuse-latest 验收决策',
       'phase06 fixture：复用感知决策上下文',
       'phase06 fixture：复用感知决策问题',
       '{}',
       'phase06 fixture 选择',
       'phase06 fixture 理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'phase06 reuse-latest 验收决策');

INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d CROSS JOIN modules m
WHERE d.title = 'phase06 reuse-latest 验收决策' AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验：复用感知聚合验证
-- ============================================================================

-- module_reuse_summary：auth-service 被 3 个 Product 复用
SELECT 'module_reuse_summary' AS check_name,
       m.name AS module_name,
       COUNT(DISTINCT pm.product_id) AS reuse_product_count
FROM modules m
LEFT JOIN product_modules pm ON pm.module_id = m.id
GROUP BY m.name
ORDER BY reuse_product_count DESC;

-- capability_summary：按 capability_key 聚合
SELECT 'capability_summary' AS check_name,
       m.capability_key,
       COUNT(*) AS module_count
FROM modules m
WHERE m.capability_key IS NOT NULL
GROUP BY m.capability_key
ORDER BY module_count DESC;
