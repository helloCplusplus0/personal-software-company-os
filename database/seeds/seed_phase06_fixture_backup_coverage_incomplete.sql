-- seed_phase06_fixture_backup_coverage_incomplete.sql
-- phase06-14 验收 fixture 8：backup-coverage-incomplete
--
-- 用途：
--   建立 phase06 备份校验失败路径 2：manifest 可读，但核心覆盖矩阵不完整。
--   GetBackupSnapshot read/verify 子路径必须返回 verify_failed + coverage_incomplete。
--
-- 数据状态（phase06-11 spec §"backup-coverage-incomplete"）：
--   - 备份产物存在（instance_backups 有记录）
--   - manifest 可读（manifest_json 为有效 ManifestSummary）
--   - 核心覆盖矩阵不完整（asset_coverage_json 缺少部分核心资产或 covered=false）
--   - backup verified 不得成立
--   - 失败语义必须单值回落到 coverage_incomplete
--
-- 依赖：reset_phase06_acceptance.sh --fixture backup-coverage-incomplete 已执行清空操作
-- 幂等：使用 WHERE NOT EXISTS 守卫，可重复执行。

BEGIN;

-- ============================================================================
-- 预置一条覆盖矩阵不完整的 instance_backups 记录
-- ============================================================================
-- manifest_json：有效 ManifestSummary（manifest_version 非空 + total_asset_entries > 0）
--   → verifyManifest 通过（第 1 步）
-- asset_coverage_json：只包含 3 类资产（缺少 6 类），不满足 9 类全覆盖
--   → verifyCoverage 失败（第 2 步）
-- schema_version / instance_version：合法（但因覆盖矩阵校验先失败，第 3 步不执行）
-- verified_status 初始为 unverified，由 GetBackupSnapshot 重新校验为 verify_failed + coverage_incomplete。

INSERT INTO instance_backups (
    manifest_json,
    asset_coverage_json,
    schema_version,
    instance_version,
    verified_status,
    backup_payload_json
)
SELECT
    '{"manifest_version":"1.0","total_asset_entries":9,"covered_asset_entries":3}'::jsonb,
    '[
      {"asset_scope":"products","covered":true},
      {"asset_scope":"modules","covered":true},
      {"asset_scope":"repositories","covered":true}
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

-- 覆盖矩阵不完整验证
SELECT 'coverage_incomplete_check' AS check_name,
       jsonb_array_length(ib.asset_coverage_json) AS coverage_entry_count,
       CASE
         WHEN jsonb_array_length(ib.asset_coverage_json) < 9
         THEN 'coverage_incomplete'
         ELSE 'coverage_complete'
       END AS expected_failure_code
FROM instance_backups ib
ORDER BY ib.created_at DESC
LIMIT 1;
