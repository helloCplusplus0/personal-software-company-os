// Package service — Backup 读编排层（read / verify 子路径）。
//
// QueryService 承接 GetBackupSnapshot 读组，
// 对齐 phase06-14 spec §"Backup 主线必须实现 read / verify 子路径与三类失败语义"。
//
// 读取侧独立 owner 约束（phase06-05 / phase06-13）：
//   - GetBackupSnapshot 必须显式承担当前阶段 read / verify 子路径语义
//   - 不得被 CreateInstanceBackup 写入响应附带的临时结果代替
//   - 由独立读取 owner（BackupRead / QueryService）承接
//
// 校验链（phase06-14 spec §"GetBackupSnapshot 校验链"）：
//  1. manifest_json 是否存在且可承接摘要
//  2. asset_coverage_json 是否覆盖 9 类核心资产
//  3. schema_version 是否与当前 schema_migrations 最新版本对齐，且 instance_version 可读取
//  - 三步全部通过时才允许返回 verified
//  - 当前阶段不得以"写出成功"直接代替 read / verify
//
// 三类失败语义（phase06-14 spec §"三类失败语义保持单值化"）：
//   - manifest_missing
//   - coverage_incomplete
//   - schema_mismatch
// 不得被折叠为泛化的 backup failed。
//
// 文件落点：backend/internal/backup/service/query_service.go
package service

import (
	"context"
	"encoding/json"

	"github.com/psco/backend/internal/backup"
	"github.com/psco/backend/internal/backup/candidate"
	"github.com/psco/backend/internal/backup/repository"
)

// QueryService 承接 Backup 快照读取与恢复前提校验编排。
//
// 依赖通过 platform 装配点注入：
//   - store：instance_backups 表读写
//   - assetReader：schema 版本读取（用于 read / verify 校验链第 3 步）
type QueryService struct {
	store       *repository.BackupStore
	assetReader *candidate.AssetReader
}

// NewQueryService 构造 QueryService。
func NewQueryService(store *repository.BackupStore, assetReader *candidate.AssetReader) *QueryService {
	return &QueryService{store: store, assetReader: assetReader}
}

// ReadBackupSnapshot 读取最新备份快照并执行 read / verify 校验链。
//
// 流程（phase06-14 spec §"GetBackupSnapshot 校验链"）：
//  1. 读取最新一条 instance_backups 记录
//  2. 依次校验：manifest / coverage / schema_version
//  3. 三步全部通过 → verified_status = verified
//  4. 任一步失败 → verified_status = verify_failed + 对应 failure_code
//  5. 回写校验结果到 instance_backups
//  6. 返回校验后的 BackupSnapshot
//
// 若无历史备份记录，返回 nil + nil error（由 handler 返回空态或 404）。
func (s *QueryService) ReadBackupSnapshot(ctx context.Context) (*backup.BackupSnapshot, error) {
	rec, err := s.store.ReadLatest(ctx)
	if err != nil {
		return nil, backup.ErrBackupSnapshotReadFailed
	}

	// 无历史备份记录
	if rec == nil {
		return nil, nil
	}

	// 执行 read / verify 校验链
	verifiedStatus, failureCode := s.verifyBackup(ctx, rec)

	// 回写校验结果到 instance_backups
	if err := s.store.UpdateVerifiedStatus(ctx, rec.ID, verifiedStatus, failureCode); err != nil {
		return nil, backup.ErrBackupPersistFailed
	}

	// 构建 BackupSnapshot
	return s.buildSnapshotFromRecord(ctx, rec, verifiedStatus, failureCode)
}

// verifyBackup 执行 read / verify 三步校验链。
//
// 校验顺序（phase06-14 spec §"GetBackupSnapshot 校验链"）：
//  1. manifest_json 是否存在且可承接摘要 → manifest_missing
//  2. asset_coverage_json 是否覆盖 9 类核心资产 → coverage_incomplete
//  3. schema_version 是否与当前 schema_migrations 最新版本对齐 → schema_mismatch
//
// 返回校验状态与失败原因代码（失败时非 nil）。
func (s *QueryService) verifyBackup(ctx context.Context, rec *repository.BackupRecord) (backup.BackupVerifiedStatus, *backup.VerifyFailureCode) {
	// 第 1 步：manifest 校验
	if !s.verifyManifest(rec) {
		fc := backup.VerifyFailureCodeManifestMissing
		return backup.BackupVerifiedStatusVerifyFailed, &fc
	}

	// 第 2 步：覆盖矩阵校验
	if !s.verifyCoverage(rec) {
		fc := backup.VerifyFailureCodeCoverageIncomplete
		return backup.BackupVerifiedStatusVerifyFailed, &fc
	}

	// 第 3 步：schema / version 校验
	if !s.verifySchemaVersion(ctx, rec) {
		fc := backup.VerifyFailureCodeSchemaMismatch
		return backup.BackupVerifiedStatusVerifyFailed, &fc
	}

	return backup.BackupVerifiedStatusVerified, nil
}

// verifyManifest 校验 manifest_json 是否存在且可承接 ManifestSummary 摘要。
//
// 校验条件：
//   - manifest_json 非 nil 且非空
//   - 可反序列化为 ManifestSummary
//   - manifest_version 非空
//   - total_asset_entries > 0
func (s *QueryService) verifyManifest(rec *repository.BackupRecord) bool {
	if len(rec.ManifestJSON) == 0 || string(rec.ManifestJSON) == "null" {
		return false
	}

	var ms backup.ManifestSummary
	if err := json.Unmarshal(rec.ManifestJSON, &ms); err != nil {
		return false
	}
	if ms.ManifestVersion == "" || ms.TotalAssetEntries <= 0 {
		return false
	}
	return true
}

// verifyCoverage 校验 asset_coverage_json 是否覆盖 9 类核心资产。
//
// 校验条件：
//   - asset_coverage_json 可反序列化为 AssetCoverageEntry 列表
//   - 9 类核心资产全部 covered = true
func (s *QueryService) verifyCoverage(rec *repository.BackupRecord) bool {
	if len(rec.AssetCoverageJSON) == 0 || string(rec.AssetCoverageJSON) == "null" {
		return false
	}

	var entries []backup.AssetCoverageEntry
	if err := json.Unmarshal(rec.AssetCoverageJSON, &entries); err != nil {
		return false
	}

	// 构建 covered 状态映射
	coveredSet := make(map[backup.BackupAssetScope]bool, len(entries))
	for _, e := range entries {
		if e.Covered {
			coveredSet[e.AssetScope] = true
		}
	}

	// 检查 9 类核心资产是否全部 covered
	for _, scope := range backup.AllBackupAssetScopes() {
		if !coveredSet[scope] {
			return false
		}
	}
	return true
}

// verifySchemaVersion 校验 schema_version 是否与当前 schema_migrations 最新版本对齐。
//
// 校验条件：
//   - rec.SchemaVersion 非空
//   - rec.InstanceVersion 非空
//   - rec.SchemaVersion == 当前 schema_migrations 最新版本
func (s *QueryService) verifySchemaVersion(ctx context.Context, rec *repository.BackupRecord) bool {
	if rec.SchemaVersion == "" || rec.InstanceVersion == "" {
		return false
	}

	currentVersion, err := s.assetReader.ReadLatestSchemaVersion(ctx)
	if err != nil {
		return false
	}

	return rec.SchemaVersion == currentVersion
}

// buildSnapshotFromRecord 从数据库记录构建校验后的 BackupSnapshot。
func (s *QueryService) buildSnapshotFromRecord(ctx context.Context, rec *repository.BackupRecord, verifiedStatus backup.BackupVerifiedStatus, failureCode *backup.VerifyFailureCode) (*backup.BackupSnapshot, error) {
	snapshot := &backup.BackupSnapshot{
		CreatedAt:      rec.CreatedAt,
		VerifiedStatus: verifiedStatus,
	}
	// failureCode 为 *VerifyFailureCode 指针，解引用后赋值（nil 时保持零值）
	if failureCode != nil {
		snapshot.VerifyFailureCode = *failureCode
	}

	// manifest_summary
	var ms backup.ManifestSummary
	if err := json.Unmarshal(rec.ManifestJSON, &ms); err == nil {
		snapshot.ManifestSummary = &ms
	} else {
		// manifest 不可解析时仍需返回结构，但用空值承接
		snapshot.ManifestSummary = &backup.ManifestSummary{}
	}

	// asset_coverage
	var entries []backup.AssetCoverageEntry
	if err := json.Unmarshal(rec.AssetCoverageJSON, &entries); err == nil {
		snapshot.AssetCoverage = entries
	} else {
		snapshot.AssetCoverage = []backup.AssetCoverageEntry{}
	}

	// schema_version_prerequisite
	currentVersion, _ := s.assetReader.ReadLatestSchemaVersion(ctx)
	snapshot.SchemaVersionPrerequisite = &backup.SchemaVersionPrerequisite{
		SchemaVersion:         rec.SchemaVersion,
		InstanceVersion:       rec.InstanceVersion,
		PrerequisiteCheckable: rec.SchemaVersion != "" && rec.InstanceVersion != "" && rec.SchemaVersion == currentVersion,
	}

	return snapshot, nil
}
