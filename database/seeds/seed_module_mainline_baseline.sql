-- seed_module_mainline_baseline.sql
-- phase02-12 验收辅助：清空模块主线 + 恢复基线数据
--
-- 用途：
--   为 phase02-12 空状态验收与基线复验提供可重复执行的"清空 -> 恢复基线"入口。
--   清空仅影响模块主线（modules 及其级联表），不影响只读前提数据
--   （products / repositories / decisions）。
--
-- 幂等：可重复执行。
--   - 清空阶段：DELETE FROM modules 依赖 ON DELETE CASCADE 级联删除
--     module_releases / product_modules / module_repositories / decision_links
--   - 恢复阶段：所有 INSERT 使用 ON CONFLICT DO NOTHING
--
-- 上游规格：phase02-12 spec §"联调环境必须可重复建立"
-- 依赖：seed_readonly_prereqs.sql 已执行（提供 products / repositories / decisions）
--
-- 使用方式：
--   # 通过封装脚本（推荐）
--   ./database/scripts/reset_module_mainline.sh              # 清空 + 恢复
--   ./database/scripts/reset_module_mainline.sh --clean-only # 仅清空（验证空状态）
--   ./database/scripts/reset_module_mainline.sh --restore-only # 仅恢复
--
--   # 直接执行 SQL（需已运行 seed_readonly_prereqs.sql）
--   podman exec -i rento-preview-postgres psql -U rento -d psco_development \
--     -v ON_ERROR_STOP=1 \
--     < database/seeds/seed_module_mainline_baseline.sql

BEGIN;

-- ============================================================================
-- 清空阶段：删除模块主线所有数据
-- ============================================================================
-- 依赖 schema 定义的 ON DELETE CASCADE：
--   - module_releases.module_id REFERENCES modules(id) ON DELETE CASCADE
--   - product_modules.module_id REFERENCES modules(id) ON DELETE CASCADE
--   - module_repositories.module_id REFERENCES modules(id) ON DELETE CASCADE
--   - decision_links.module_id REFERENCES modules(id) ON DELETE CASCADE
-- 因此只需 DELETE FROM modules 即可级联清空所有关联表。
DELETE FROM modules;

-- ============================================================================
-- 恢复阶段：重建 phase02-12 验收基线数据
-- ============================================================================
-- 设计原则：通过 name / title 查找 product_id / repository_id / decision_id，
-- 避免硬编码 UUID，保证在不同环境（不同 UUID 种子）下均可复现。

-- 模块（2 个）
INSERT INTO modules (name, description, status) VALUES
  ('auth-service', '认证服务模块，提供 OAuth2 与 JWT 支持', 'active'),
  ('integration-test-module', 'phase02-12 联调验证用模块', 'active')
ON CONFLICT (name) DO NOTHING;

-- 版本（3 个）：auth-service 1.0.0 / 1.1.0，integration-test-module 0.1.0
INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '1.0.0', 'active', '2026-07-01T00:00:00+00:00'
FROM modules m WHERE m.name = 'auth-service'
ON CONFLICT (module_id, version) DO NOTHING;

INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '1.1.0', 'active', '2026-07-10T00:00:00+00:00'
FROM modules m WHERE m.name = 'auth-service'
ON CONFLICT (module_id, version) DO NOTHING;

INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '0.1.0', 'active', '2026-08-05T15:00:00+00:00'
FROM modules m WHERE m.name = 'integration-test-module'
ON CONFLICT (module_id, version) DO NOTHING;

-- 产品绑定（2 个）：两个模块均绑定到 Product A
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND p.name = 'Product A'
ON CONFLICT (module_id, product_id) DO NOTHING;

-- 仓库映射（2 个）：两个模块均映射到 main-repo
INSERT INTO module_repositories (module_id, repository_id)
SELECT m.id, r.id
FROM modules m
CROSS JOIN repositories r
WHERE m.name IN ('auth-service', 'integration-test-module')
  AND r.name = 'main-repo'
ON CONFLICT (module_id, repository_id) DO NOTHING;

-- Decision 关联（1 个）：auth-service 关联到"关于 auth-service 技术选型的决策"
-- 依赖 seed_readonly_prereqs.sql 已插入 decisions 记录
INSERT INTO decision_links (module_id, decision_id)
SELECT m.id, d.id
FROM modules m
CROSS JOIN decisions d
WHERE m.name = 'auth-service'
  AND d.title = '关于 auth-service 技术选型的决策'
ON CONFLICT (module_id, decision_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 验证恢复结果（信息性输出，不阻断）
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) FROM modules
UNION ALL SELECT 'module_releases', COUNT(*) FROM module_releases
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'module_repositories', COUNT(*) FROM module_repositories
UNION ALL SELECT 'decision_links', COUNT(*) FROM decision_links;
