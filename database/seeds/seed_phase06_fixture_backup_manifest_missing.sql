-- seed_phase06_fixture_backup_manifest_missing.sql
-- phase06-14 验收 fixture 7：backup-manifest-missing
--
-- 用途：
--   建立 phase06 备份校验失败路径 1：备份产物存在，但 manifest 缺失或不可读取。
--   GetBackupSnapshot read/verify 子路径必须返回 verify_failed + manifest_missing。
--
-- 数据状态（phase06-11 spec §"backup-manifest-missing"）：
--   - 备份产物存在（instance_backups 有记录）
--   - manifest 缺失或不可解析（manifest_json 为空 / null / 缺少必填字段）
--   - backup verified 不得成立
--   - 失败语义必须单值回落到 manifest_missing
--
-- 依赖：reset_phase06_acceptance.sh --fixture backup-manifest-missing 已执行清空操作
-- 幂等：使用 WHERE NOT EXISTS 守卫，可重复执行。

BEGIN;

-- ============================================================================
-- 预置一条 manifest 缺失的 instance_backups 记录
-- ============================================================================
-- manifest_json 设为空 JSON 对象（不可反序列化为有效 ManifestSummary：
--   manifest_version 为空 + total_asset_entries = 0，verifyManifest 返回 false）
-- asset_coverage_json / schema_version / instance_version 均合法，
-- 但因 manifest 校验在第 1 步即失败，后续步骤不会执行。
-- verified_status 初始为 unverified，由 GetBackupSnapshot 重新校验为 verify_failed + manifest_missing。

INSERT INTO instance_backups (
    manifest_json,
    asset_coverage_json,
    schema_version,
    instance_version,
    verified_status,
    backup_payload_json
)
SELECT
    '{}'::jsonb,
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

-- manifest 缺失验证
SELECT 'manifest_missing_check' AS check_name,
       ib.manifest_json::text AS manifest_json,
       CASE
         WHEN ib.manifest_json::text = '{}' OR ib.manifest_json IS NULL
              OR ib.manifest_json::text = 'null'
         THEN 'manifest_missing'
         ELSE 'manifest_present'
       END AS expected_failure_code
FROM instance_backups ib
ORDER BY ib.created_at DESC
LIMIT 1;
