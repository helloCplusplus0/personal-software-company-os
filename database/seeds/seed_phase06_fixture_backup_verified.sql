-- seed_phase06_fixture_backup_verified.sql
-- phase06-14 验收 fixture 6：backup-verified
--
-- 用途：
--   建立 phase06 备份可验状态：至少满足 completed-bound，
--   且具备可重复触发 CreateInstanceBackup 与 GetBackupSnapshot 的正式前置数据。
--   预置一条通过校验链的 instance_backups 记录，
--   使 GET /api/dashboard/backup 可直接返回 verified 快照。
--
-- 数据状态（phase06-11 spec §"backup-verified"）：
--   - 至少满足 completed-bound（四类 canonical 对象 + 核心绑定关系）
--   - 具备可重复触发 CreateInstanceBackup 与 GetBackupSnapshot 的正式前置数据
--   - 验收时可直接验证：
--     * 备份产物可生成
--     * manifest 可重新读取
--     * 覆盖矩阵完整
--     * schema / version 前提可校验
--   - 只有上述条件同时成立，当前 fixture 才允许作为 backup verified 的正式证据
--
-- 依赖：reset_phase06_acceptance.sh --fixture backup-verified 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING + WHERE NOT EXISTS，可重复执行。

BEGIN;

-- ============================================================================
-- 完整 canonical 数据（满足 completed-bound）
-- ============================================================================

INSERT INTO modules (name, description, status, capability_key) VALUES
  ('auth-service', 'phase06 fixture：备份可验模块（认证服务）', 'active', 'auth'),
  ('integration-test-module', 'phase06 fixture：备份可验模块（集成测试）', 'active', 'web_frontend')
ON CONFLICT (name) DO NOTHING;

INSERT INTO products (name, description, status) VALUES
  ('Product A', 'phase06 fixture：备份可验产品 A', 'active'),
  ('Product B', 'phase06 fixture：备份可验产品 B', 'active')
ON CONFLICT (name) DO NOTHING;

INSERT INTO repositories (name, url, provider, status) VALUES
  ('main-repo', 'https://github.com/psco/main-repo', 'github', 'active')
ON CONFLICT (name) DO NOTHING;

INSERT INTO module_releases (module_id, version, status, released_at)
SELECT m.id, '1.0.0', 'active', '2026-07-01T00:00:00+00:00'
FROM modules m WHERE m.name = 'auth-service'
ON CONFLICT (module_id, version) DO NOTHING;

INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'phase06 backup-verified 验收决策',
       'phase06 fixture：备份可验决策上下文',
       'phase06 fixture：备份可验决策问题',
       '{}',
       'phase06 fixture 选择',
       'phase06 fixture 理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'phase06 backup-verified 验收决策');

-- 核心绑定关系
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m CROSS JOIN products p
WHERE m.name IN ('auth-service', 'integration-test-module') AND p.name = 'Product A'
ON CONFLICT (module_id, product_id) DO NOTHING;

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

INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d CROSS JOIN modules m
WHERE d.title = 'phase06 backup-verified 验收决策' AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

-- ============================================================================
-- 预置一条通过校验链的 instance_backups 记录
-- ============================================================================
-- manifest_json：有效 ManifestSummary（manifest_version 非空 + total_asset_entries > 0）
-- asset_coverage_json：9 类核心资产全部 covered=true
-- schema_version：取 schema_migrations 最新版本（保证校验通过）
-- instance_version：非空
-- verified_status：初始为 unverified，由 GetBackupSnapshot read/verify 子路径重新校验为 verified

INSERT INTO instance_backups (
    manifest_json,
    asset_coverage_json,
    schema_version,
    instance_version,
    verified_status,
    backup_payload_json
)
SELECT
    '{"manifest_version":"1.0","total_asset_entries":9,"covered_asset_entries":9}'::jsonb,
    '[
      {"asset_scope":"products","covered":true},
      {"asset_scope":"modules","covered":true},
      {"asset_scope":"releases","covered":true},
      {"asset_scope":"repositories","covered":true},
      {"asset_scope":"decisions","covered":true},
      {"asset_scope":"decision_links","covered":true},
      {"asset_scope":"product_modules","covered":true},
      {"asset_scope":"product_repositories","covered":true},
      {"asset_scope":"module_repositories","covered":true}
    ]'::jsonb,
    (SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1),
    'phase06-v0.1',
    'unverified',
    '{}'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM instance_backups LIMIT 1);

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'instance_backups' AS tbl, COUNT(*) AS cnt FROM instance_backups;

-- 备份校验前提验证
SELECT 'backup_verified_prerequisite' AS check_name,
       ib.schema_version,
       ib.instance_version,
       (SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1) AS current_schema_version,
       CASE
         WHEN ib.schema_version = (SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1)
          AND ib.instance_version != ''
         THEN 'schema_match'
         ELSE 'schema_mismatch'
       END AS schema_check
FROM instance_backups ib
ORDER BY ib.created_at DESC
LIMIT 1;
