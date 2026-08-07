-- seed_product_repository_mainline_baseline.sql
-- phase04-12 验收辅助：清空 Product / Repository 主线 + 恢复基线数据
--
-- 用途：
--   为 phase04-12 空状态验收与基线复验提供可重复执行的"清空 -> 恢复基线"入口。
--   清空仅影响 Product / Repository 主线（products / repositories 及其关系表
--   product_repositories / product_modules / module_repositories），
--   不影响模块主线（modules / module_releases）与 Decision 主线（decisions / decision_links）。
--
-- 幂等：可重复执行。
--   - 清空阶段：按 FK 依赖顺序 DELETE 关系表 -> 实体表
--     （product_repositories / product_modules / module_repositories 先于
--      products / repositories，避免 ON DELETE RESTRICT 冲突）
--   - 恢复阶段：
--     * Product A / B / C 与 main-repo / mirror-repo 使用
--       "先 UPDATE 收口 readonly placeholder -> 再 INSERT 补缺"，
--       确保 --restore-only 模式下也能把 readonly placeholder 收口为正式基线内容
--     * product_modules / module_repositories / product_repositories
--       使用 ON CONFLICT DO NOTHING
--
-- 上游规格：phase04-09 spec §"phase04 基线 seed 与 reset 脚本必须可重复执行"
--           phase04-10 product_repository_binding_spec_v0.1.md §基线种子数据
--           phase04-12 spec §"phase04 基线 seed 与 reset 脚本必须可重复执行"
-- 依赖：reset_module_mainline.sh 已执行（提供 modules 基线，
--       product_modules / module_repositories 依赖 modules）
--
-- 覆盖维度（phase04-12 spec）：
--   - 至少 3 条 Product（Product A / B / C）
--   - 至少 2 条 Repository（main-repo / mirror-repo）
--   - 至少 2 条 product_modules（auth-service / integration-test-module -> Product A）
--   - 至少 1 条 product_repositories（Product A -> main-repo）
--   - 至少 2 条 module_repositories（auth-service / integration-test-module -> main-repo）
--   - 至少 1 条无已绑定仓库的 Product（Product B / Product C）
--   - 至少 1 条无已绑定产品的 Repository（mirror-repo）
--
-- 使用方式：
--   # 通过封装脚本（推荐）
--   ./database/scripts/reset_product_repository_mainline.sh              # 清空 + 恢复
--   ./database/scripts/reset_product_repository_mainline.sh --clean-only # 仅清空（验证空状态）
--   ./database/scripts/reset_product_repository_mainline.sh --restore-only # 仅恢复
--
--   # 直接执行 SQL（需已运行 reset_module_mainline.sh）
--   podman exec -i rento-preview-postgres psql -U rento -d psco_development \
--     -v ON_ERROR_STOP=1 \
--     < database/seeds/seed_product_repository_mainline_baseline.sql

BEGIN;

-- ============================================================================
-- 清空阶段：按 FK 依赖顺序删除 Product / Repository 主线所有数据
-- ============================================================================
-- 清空顺序：先关系表，后实体表。
-- - product_repositories: REFERENCES products / repositories (ON DELETE RESTRICT)
-- - product_modules: REFERENCES products (ON DELETE RESTRICT) / modules (ON DELETE CASCADE)
-- - module_repositories: REFERENCES repositories (ON DELETE RESTRICT) / modules (ON DELETE CASCADE)
-- - products / repositories: 实体表
--
-- 不清空 modules / decisions / module_releases / decision_links（由各自 reset 脚本负责）。
DELETE FROM product_repositories;
DELETE FROM product_modules;
DELETE FROM module_repositories;
DELETE FROM products;
DELETE FROM repositories;

-- ============================================================================
-- 恢复阶段：重建 phase04-12 验收基线数据
-- ============================================================================
-- 设计原则：通过 name 查找 product_id / repository_id / module_id，避免硬编码 UUID。

-- ----------------------------------------------------------------------------
-- Products（3 条）：Product A / B / C
-- ----------------------------------------------------------------------------
-- 先 UPDATE 收口可能来自 seed_readonly_prereqs.sql 的 placeholder 数据，
-- 确保 --restore-only 模式下也能把 readonly placeholder 收口为正式基线内容。
UPDATE products
  SET description = 'phase04 正式基线产品 A，已绑定模块与仓库',
      status = 'active'
  WHERE name = 'Product A';

UPDATE products
  SET description = 'phase04 正式基线产品 B，无已绑定仓库',
      status = 'active'
  WHERE name = 'Product B';

-- Product C 设为 archived，验证 archived 状态在列表与详情中的读取
UPDATE products
  SET description = 'phase04 正式基线产品 C，已归档',
      status = 'archived'
  WHERE name = 'Product C';

-- INSERT 补缺（仅在记录不存在时插入，保证幂等）
INSERT INTO products (name, description, status)
SELECT 'Product A', 'phase04 正式基线产品 A，已绑定模块与仓库', 'active'
WHERE NOT EXISTS (SELECT 1 FROM products WHERE name = 'Product A');

INSERT INTO products (name, description, status)
SELECT 'Product B', 'phase04 正式基线产品 B，无已绑定仓库', 'active'
WHERE NOT EXISTS (SELECT 1 FROM products WHERE name = 'Product B');

INSERT INTO products (name, description, status)
SELECT 'Product C', 'phase04 正式基线产品 C，已归档', 'archived'
WHERE NOT EXISTS (SELECT 1 FROM products WHERE name = 'Product C');

-- ----------------------------------------------------------------------------
-- Repositories（2 条）：main-repo / mirror-repo
-- ----------------------------------------------------------------------------
-- 先 UPDATE 收口可能来自 seed_readonly_prereqs.sql 的 placeholder 数据。
UPDATE repositories
  SET url = 'https://github.com/psco/main-repo',
      provider = 'github',
      status = 'active'
  WHERE name = 'main-repo';

UPDATE repositories
  SET url = 'https://github.com/psco/mirror-repo',
      provider = 'github',
      status = 'active'
  WHERE name = 'mirror-repo';

-- INSERT 补缺
INSERT INTO repositories (name, url, provider, status)
SELECT 'main-repo', 'https://github.com/psco/main-repo', 'github', 'active'
WHERE NOT EXISTS (SELECT 1 FROM repositories WHERE name = 'main-repo');

INSERT INTO repositories (name, url, provider, status)
SELECT 'mirror-repo', 'https://github.com/psco/mirror-repo', 'github', 'active'
WHERE NOT EXISTS (SELECT 1 FROM repositories WHERE name = 'mirror-repo');

-- ----------------------------------------------------------------------------
-- product_modules（2 条）：auth-service / integration-test-module -> Product A
-- ----------------------------------------------------------------------------
-- 复用 modules 基线（auth-service / integration-test-module），
-- 通过 name 查找 module_id 与 product_id，不硬编码 UUID。
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND p.name = 'Product A'
ON CONFLICT (module_id, product_id) DO NOTHING;

-- ----------------------------------------------------------------------------
-- product_repositories（1 条）：Product A -> main-repo
-- ----------------------------------------------------------------------------
-- Product B / Product C 不绑定仓库（满足"至少 1 条无已绑定仓库的 Product"）。
-- mirror-repo 不绑定产品（满足"至少 1 条无已绑定产品的 Repository"）。
INSERT INTO product_repositories (product_id, repository_id)
SELECT p.id, r.id
FROM products p
CROSS JOIN repositories r
WHERE p.name = 'Product A'
  AND r.name = 'main-repo'
ON CONFLICT (product_id, repository_id) DO NOTHING;

-- ----------------------------------------------------------------------------
-- module_repositories（2 条）：auth-service / integration-test-module -> main-repo
-- ----------------------------------------------------------------------------
-- 复用 modules 基线与 repositories 基线，通过 name 查找 ID。
INSERT INTO module_repositories (module_id, repository_id)
SELECT m.id, r.id
FROM modules m
CROSS JOIN repositories r
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND r.name = 'main-repo'
ON CONFLICT (module_id, repository_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 验证恢复结果（信息性输出，不阻断）
-- ============================================================================
SELECT 'products' AS tbl, COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'product_repositories', COUNT(*) FROM product_repositories
UNION ALL SELECT 'module_repositories', COUNT(*) FROM module_repositories;
