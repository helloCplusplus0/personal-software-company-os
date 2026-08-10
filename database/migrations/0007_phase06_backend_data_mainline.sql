-- 0007_phase06_backend_data_mainline.sql
-- Phase06 Onboarding + Sovereignty + Reuse 后端与数据主线核心表
-- 直接承接表：modules (原位升级), instance_exports (新增), instance_backups (新增)
-- 上游规格：phase06-14 spec §"Phase06 数据主线必须通过 0007 migration 补齐最小承接位"
--           phase06-12 formal spec §"Export / Backup 正式语义与边界冻结"
--           phase06-04 / phase06-09 §"module_reuse_summary / capability_summary 最小读模型"
--
-- 升级策略：
--   - modules 原位新增 capability_key（NULL 允许，不创建独立 capabilities 表）
--   - 新增 instance_exports 承接导出快照元数据
--   - 新增 instance_backups 承接备份快照元数据与恢复前提校验状态
--   - 不引入异步统计表、离线聚合作业或可编辑能力字典
--
-- 兼容性：
--   - 历史 modules 记录 capability_key 为 NULL，不影响既有读取
--   - 未填写 capability_key 的 Module 继续允许存在，不误判为写入失败

-- ============================================================================
-- modules 原位升级：新增 capability_key
-- ============================================================================
-- phase06-04 / phase06-12 §"Reuse Summary 正式读模型与页面挂接位冻结"：
--   capability_summary 事实来源冻结为 Module.capability_key + capability_label 映射
--   不得引入独立 Capability 实体、表或可编辑字典

ALTER TABLE modules
  ADD COLUMN capability_key TEXT NULL;

-- capability_key 索引：支持按能力聚合的读取性能
CREATE INDEX idx_modules_capability_key ON modules (capability_key) WHERE capability_key IS NOT NULL;

-- ============================================================================
-- instance_exports：导出快照元数据（新增）
-- ============================================================================
-- phase06-08 / phase06-12 §"Export / Backup 正式语义与边界冻结"：
--   Export 最小覆盖矩阵必须至少包含 9 类核心资产
--   导出快照必须可通过数据库元数据表形成最新可读取快照
--   不得只返回一次性导出响应

CREATE TABLE instance_exports (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    -- 导出执行结果状态：success / in_progress / failed
    result_status       TEXT NOT NULL CHECK (result_status IN ('success', 'in_progress', 'failed')),
    -- 导出结果摘要文本
    result_summary      TEXT NOT NULL DEFAULT '',
    -- 9 类核心资产覆盖矩阵（JSON 数组形式存储 ExportAssetScope 枚举值）
    asset_scope_json    JSONB NOT NULL DEFAULT '[]'::jsonb,
    -- 导出产物载荷（JSON 形式存储 9 类核心资产的实际数据快照）
    artifact_payload_json JSONB NOT NULL DEFAULT '{}'::jsonb
);

-- 按创建时间倒序读取最新导出快照
CREATE INDEX idx_instance_exports_created_at ON instance_exports (created_at DESC);

-- ============================================================================
-- instance_backups：备份快照元数据与恢复前提校验状态（新增）
-- ============================================================================
-- phase06-05 / phase06-13 §"backup_snapshot 读取侧必须由独立读取 owner 承接"：
--   GetBackupSnapshot 必须显式承担当前阶段 read / verify 子路径语义
--   不得把"CreateInstanceBackup 写入响应里顺带返回一次 manifest"解释为已满足 snapshot 正式读取侧
--
-- phase06-12 §"backup verified 最小成立条件"：
--   1. 已生成可读取的备份产物
--   2. 可重新读取并解析 manifest
--   3. manifest 中可见核心资产覆盖矩阵
--   4. schema / version 恢复前提可校验

CREATE TABLE instance_backups (
    id                      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    -- 备份 manifest 摘要（JSON 形式存储 ManifestSummary）
    manifest_json           JSONB NOT NULL DEFAULT '{}'::jsonb,
    -- 9 类核心资产覆盖矩阵（JSON 数组形式存储 AssetCoverageEntry）
    asset_coverage_json     JSONB NOT NULL DEFAULT '[]'::jsonb,
    -- 备份时的 schema 版本（对应 schema_migrations 最新版本）
    schema_version          TEXT NOT NULL DEFAULT '',
    -- 备份时的实例版本标识
    instance_version        TEXT NOT NULL DEFAULT '',
    -- 恢复前提校验状态：unverified / verified / verify_failed
    verified_status         TEXT NOT NULL DEFAULT 'unverified' CHECK (verified_status IN ('unverified', 'verified', 'verify_failed')),
    -- 校验失败原因代码（仅 verified_status = verify_failed 时有值）
    -- 只允许：manifest_missing / coverage_incomplete / schema_mismatch
    verify_failure_code     TEXT NULL CHECK (verify_failure_code IN ('manifest_missing', 'coverage_incomplete', 'schema_mismatch')),
    -- 备份产物载荷（JSON 形式存储 9 类核心资产的实际数据快照）
    backup_payload_json     JSONB NOT NULL DEFAULT '{}'::jsonb
);

-- 按创建时间倒序读取最新备份快照
CREATE INDEX idx_instance_backups_created_at ON instance_backups (created_at DESC);
-- 按校验状态筛选
CREATE INDEX idx_instance_backups_verified_status ON instance_backups (verified_status);
