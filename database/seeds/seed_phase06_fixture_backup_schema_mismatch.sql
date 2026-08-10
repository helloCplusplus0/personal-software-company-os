-- seed_phase06_fixture_backup_schema_mismatch.sql
-- phase06-14 验收 fixture 9：backup-schema-mismatch
--
-- 用途：
--   建立 phase06 备份校验失败路径 3：manifest 与覆盖矩阵可读，但 schema/version 前提不可校验。
--   GetBackupSnapshot read/verify 子路径必须返回 verify_failed + schema_mismatch。
--
-- 数据状态（phase06-11 spec §"backup-schema-mismatch"）：
--   - 备份产物存在（instance_backups 有记录）
--   - manifest 可读（manifest_json 为有效 ManifestSummary）
--   - 核心覆盖矩阵完整（asset_coverage_json 9 类全部 covered=true）
--   - schema / version 前提不可校验（schema_version 与当前 schema_migrations 最新版本不匹配）
--   - backup verified 不得成立
--   - 失败语义必须单值回落到 schema_mismatch
--
-- 依赖：reset_phase06_acceptance.sh --fixture backup-schema-mismatch 已执行清空操作
-- 幂等：使用 WHERE NOT EXISTS 守卫，可重复执行。

BEGIN;

-- ============================================================================
-- 预置一条 schema_version 不匹配的 instance_backups 记录
-- ============================================================================
-- manifest_json：有效 ManifestSummary → verifyManifest 通过（第 1 步）
-- asset_coverage_json：9 类核心资产全部 covered=true → verifyCoverage 通过（第 2 步）
-- schema_version：故意设为不存在的版本号（0099_nonexistent_migration）
--   → verifySchemaVersion 失败（第 3 步）
-- instance_version：非空（满足 instance_version 可读取前提，但 schema_version 不匹配）
-- verified_status 初始为 unverified，由 GetBackupSnapshot 重新校验为 verify_failed + schema_mismatch。

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
    '0099_nonexistent_migration',
    'phase06-v0.1',
    'unverified',
    '{}'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM instance_backups LIMIT 1);

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'instance_backups' AS tbl, COUNT(*) AS cnt FROM instance_backups;

-- schema 不匹配验证
SELECT 'schema_mismatch_check' AS check_name,
       ib.schema_version AS backup_schema_version,
       (SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1) AS current_schema_version,
       CASE
         WHEN ib.schema_version != (SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1)
         THEN 'schema_mismatch'
         ELSE 'schema_match'
       END AS expected_failure_code
FROM instance_backups ib
ORDER BY ib.created_at DESC
LIMIT 1;
