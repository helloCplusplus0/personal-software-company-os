// Package backup 承载备份（Backup）后端模块的全部业务实现。
//
// 分层语义（对齐 phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"）：
//   - handler/     入口层：只负责承接 HTTP 请求与返回结果
//   - service/     业务编排层：备份执行编排与 read / verify 子路径校验
//   - candidate/   外部连接层：跨模块 9 类核心资产 reader（Backup 自拥有）
//   - repository/  持久化层：instance_backups 表读写
//
// 当前阶段 Backup 实现为"创建备份 + 重新读取并校验恢复前提"的双路径主线，
// 并把三类失败语义（manifest_missing / coverage_incomplete / schema_mismatch）
// 稳定保留在后端与数据主线内。
//
// 读取侧独立 owner 约束（phase06-05 / phase06-13）：
//   - GetBackupSnapshot 必须显式承担当前阶段 read / verify 子路径语义
//   - 不得把"CreateInstanceBackup 写入响应里顺带返回一次 manifest"解释为已满足 snapshot 正式读取侧
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/backup/v1/backup.proto 单向派生或显式对齐。
package backup

import "time"

// ============================================================================
// 枚举类型
// ============================================================================

// BackupAssetScope 备份覆盖矩阵的核心资产类型。
// 对齐 proto BackupAssetScope：9 类核心资产（与 ExportAssetScope 语义对齐但独立定义）。
type BackupAssetScope string

const (
	BackupAssetScopeUnspecified       BackupAssetScope = ""
	BackupAssetScopeProducts          BackupAssetScope = "products"
	BackupAssetScopeModules           BackupAssetScope = "modules"
	BackupAssetScopeReleases          BackupAssetScope = "releases"
	BackupAssetScopeRepositories      BackupAssetScope = "repositories"
	BackupAssetScopeDecisions         BackupAssetScope = "decisions"
	BackupAssetScopeDecisionLinks     BackupAssetScope = "decision_links"
	BackupAssetScopeProductModules    BackupAssetScope = "product_modules"
	BackupAssetScopeProductRepositories BackupAssetScope = "product_repositories"
	BackupAssetScopeModuleRepositories BackupAssetScope = "module_repositories"
)

// AllBackupAssetScopes 返回当前阶段冻结的 9 类核心资产完整列表（不含 UNSPECIFIED）。
func AllBackupAssetScopes() []BackupAssetScope {
	return []BackupAssetScope{
		BackupAssetScopeProducts,
		BackupAssetScopeModules,
		BackupAssetScopeReleases,
		BackupAssetScopeRepositories,
		BackupAssetScopeDecisions,
		BackupAssetScopeDecisionLinks,
		BackupAssetScopeProductModules,
		BackupAssetScopeProductRepositories,
		BackupAssetScopeModuleRepositories,
	}
}

// BackupVerifiedStatus 备份恢复前提校验状态。
// 对齐 proto BackupVerifiedStatus。
type BackupVerifiedStatus string

const (
	BackupVerifiedStatusUnspecified   BackupVerifiedStatus = ""
	BackupVerifiedStatusUnverified    BackupVerifiedStatus = "unverified"
	BackupVerifiedStatusVerified      BackupVerifiedStatus = "verified"
	BackupVerifiedStatusVerifyFailed  BackupVerifiedStatus = "verify_failed"
)

// VerifyFailureCode 校验失败原因代码。
// 只允许：manifest_missing / coverage_incomplete / schema_mismatch。
// 对齐 phase06-14 spec §"三类失败语义保持单值化"。
type VerifyFailureCode string

const (
	VerifyFailureCodeUnspecified       VerifyFailureCode = ""
	VerifyFailureCodeManifestMissing   VerifyFailureCode = "manifest_missing"
	VerifyFailureCodeCoverageIncomplete VerifyFailureCode = "coverage_incomplete"
	VerifyFailureCodeSchemaMismatch    VerifyFailureCode = "schema_mismatch"
)

// ============================================================================
// 核心消息 DTO
// ============================================================================

// AssetCoverageEntry 备份覆盖矩阵单值项。
// 对齐 proto AssetCoverageEntry。
type AssetCoverageEntry struct {
	AssetScope BackupAssetScope `json:"asset_scope"`
	Covered    bool             `json:"covered"`
}

// ManifestSummary 备份 manifest 摘要。
// 对齐 proto ManifestSummary。
type ManifestSummary struct {
	ManifestVersion     string `json:"manifest_version"`
	TotalAssetEntries   int    `json:"total_asset_entries"`
	CoveredAssetEntries int    `json:"covered_asset_entries"`
}

// SchemaVersionPrerequisite 备份恢复前提的 schema / version 校验信息。
// 对齐 proto SchemaVersionPrerequisite。
type SchemaVersionPrerequisite struct {
	SchemaVersion       string `json:"schema_version"`
	InstanceVersion     string `json:"instance_version"`
	PrerequisiteCheckable bool `json:"prerequisite_checkable"`
}

// BackupSnapshot 备份快照主读模型。
// 对齐 proto BackupSnapshot。
// 该消息同时承接读取校验（GetBackupSnapshot）与备份执行（CreateInstanceBackup）的结果形状。
type BackupSnapshot struct {
	CreatedAt                time.Time                  `json:"created_at"`
	ManifestSummary          *ManifestSummary           `json:"manifest_summary"`
	AssetCoverage            []AssetCoverageEntry       `json:"asset_coverage"`
	SchemaVersionPrerequisite *SchemaVersionPrerequisite `json:"schema_version_prerequisite"`
	VerifiedStatus           BackupVerifiedStatus       `json:"verified_status"`
	// VerifyFailureCode 仅在 verified_status = verify_failed 时有值（对齐 instance_backups.verify_failure_code）。
	// phase06-14 spec §"三类失败语义保持单值化"：
	//   DTO / HTTP / 前端消费侧不得把这三类失败折叠为泛化的 backup failed。
	//   当前字段已进入 backup.proto 正式合同，用于单值表达 manifest / coverage / schema 三类失败家族。
	VerifyFailureCode VerifyFailureCode `json:"verify_failure_code,omitempty"`
}

// ============================================================================
// 响应 DTO
// ============================================================================

// BackupSnapshotReadResult GetBackupSnapshot 的响应结构。
// 对齐 proto GetBackupSnapshotResponse：单一 snapshot 字段包装主读模型。
// 该响应是当前阶段正式 read / verify 子路径合同出口，
// 由独立读取 owner（BackupRead / QueryService）承接，不与 BackupWrite 写入响应耦合。
type BackupSnapshotReadResult struct {
	Snapshot *BackupSnapshot `json:"snapshot"`
}

// CreateBackupResult CreateInstanceBackup 的响应结构。
// 对齐 proto CreateInstanceBackupResponse：单一 snapshot 字段包装备份结果。
// 该响应不得替代 GetBackupSnapshot 的正式读取合同（phase06-05 / phase06-13 读取侧独立 owner 约束）。
type CreateBackupResult struct {
	Snapshot *BackupSnapshot `json:"snapshot"`
}
