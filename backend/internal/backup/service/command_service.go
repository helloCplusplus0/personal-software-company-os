// Package service — Backup 写编排层。
//
// CommandService 承接 CreateInstanceBackup 写组，
// 对齐 phase06-14 spec §"Backup 主线必须实现 read / verify 子路径与三类失败语义"。
//
// 写入语义（phase06-14 spec §"CreateInstanceBackup 持久化语义"）：
//   - 装配与 Export 相同的 9 类核心资产
//   - 同时生成 manifest_json、asset_coverage_json、schema_version、instance_version
//   - 初始写入时 verified_status 必须为 unverified
//   - 写入结果必须持久化到 instance_backups
//
// 文件落点：backend/internal/backup/service/command_service.go
package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/psco/backend/internal/backup"
	"github.com/psco/backend/internal/backup/candidate"
	"github.com/psco/backend/internal/backup/repository"
)

// 当前阶段的实例版本标识。
// 用于 schema_version_prerequisite.instance_version 字段。
// 后续可从配置或构建信息注入，当前阶段冻结为静态值。
const currentInstanceVersion = "phase06-v0.1"

// manifest 版本号，对齐 ManifestSummary.manifest_version。
const currentManifestVersion = "1.0"

// CommandService 承接 Backup 备份执行编排。
//
// 依赖通过 platform 装配点注入：
//   - store：instance_backups 表读写
//   - assetReader：9 类核心资产装配 + schema 版本读取
type CommandService struct {
	store       *repository.BackupStore
	assetReader *candidate.AssetReader
}

// NewCommandService 构造 CommandService。
func NewCommandService(store *repository.BackupStore, assetReader *candidate.AssetReader) *CommandService {
	return &CommandService{store: store, assetReader: assetReader}
}

// CreateInstanceBackup 装配 9 类核心资产并持久化备份快照。
//
// 流程（phase06-14 spec §"CreateInstanceBackup 持久化语义"）：
//  1. 通过 candidate.AssetReader 装配 9 类 canonical 数据
//  2. 读取当前 schema_migrations 最新版本作为 schema_version
//  3. 组装 manifest_json（ManifestSummary）
//  4. 组装 asset_coverage_json（9 类资产的 AssetCoverageEntry，全部 covered=true）
//  5. 组装 backup_payload_json（9 类资产完整数据载荷）
//  6. 持久化到 instance_backups，verified_status = unverified
//  7. 返回写入后的 BackupSnapshot（verified_status = unverified）
//
// 注意：写入响应不得替代 GetBackupSnapshot 的正式读取合同（phase06-05 / phase06-13）。
func (s *CommandService) CreateInstanceBackup(ctx context.Context) (*backup.BackupSnapshot, error) {
	// 1. 装配 9 类核心资产
	payload, err := s.assetReader.ReadCoreAssets(ctx)
	if err != nil {
		return nil, backup.ErrAssetReadFailed
	}

	// 2. 读取当前 schema_migrations 最新版本
	schemaVersion, err := s.assetReader.ReadLatestSchemaVersion(ctx)
	if err != nil {
		return nil, backup.ErrSchemaVersionReadFailed
	}

	// 3. 组装 manifest_json（ManifestSummary）
	allScopes := backup.AllBackupAssetScopes()
	manifestSummary := backup.ManifestSummary{
		ManifestVersion:     currentManifestVersion,
		TotalAssetEntries:   len(allScopes),
		CoveredAssetEntries: len(allScopes),
	}
	manifestJSON, err := json.Marshal(manifestSummary)
	if err != nil {
		return nil, fmt.Errorf("marshal manifest: %w", err)
	}

	// 4. 组装 asset_coverage_json（9 类资产全部 covered=true）
	coverageEntries := make([]backup.AssetCoverageEntry, 0, len(allScopes))
	for _, scope := range allScopes {
		coverageEntries = append(coverageEntries, backup.AssetCoverageEntry{
			AssetScope: scope,
			Covered:    true,
		})
	}
	assetCoverageJSON, err := json.Marshal(coverageEntries)
	if err != nil {
		return nil, fmt.Errorf("marshal asset_coverage: %w", err)
	}

	// 5. 组装 backup_payload_json（9 类资产完整数据载荷）
	backupPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal backup_payload: %w", err)
	}

	// 6. 持久化到 instance_backups（verified_status = unverified）
	rec, err := s.store.CreateLatest(ctx, manifestJSON, assetCoverageJSON, schemaVersion, currentInstanceVersion, backupPayload)
	if err != nil {
		return nil, backup.ErrBackupPersistFailed
	}

	// 7. 返回写入后的 BackupSnapshot（verified_status = unverified）
	return &backup.BackupSnapshot{
		CreatedAt:                 rec.CreatedAt,
		ManifestSummary:          &manifestSummary,
		AssetCoverage:            coverageEntries,
		SchemaVersionPrerequisite: &backup.SchemaVersionPrerequisite{
			SchemaVersion:         schemaVersion,
			InstanceVersion:       currentInstanceVersion,
			PrerequisiteCheckable: true,
		},
		VerifiedStatus: backup.BackupVerifiedStatusUnverified,
	}, nil
}
