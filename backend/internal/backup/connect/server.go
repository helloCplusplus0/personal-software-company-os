// Package connect — Backup Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Backup 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/backup/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/psco/backend/internal/gen/proto/psco/backup/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/backup/v1/backupv1connect"
	"github.com/psco/backend/internal/backup"
	"github.com/psco/backend/internal/backup/service"
	"github.com/psco/backend/internal/connecterrors"
)

// Server 实现 BackupServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.BackupServiceHandler = (*Server)(nil)

// NewServer 构造 Backup Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{querySvc: querySvc, commandSvc: commandSvc}
}

// GetBackupSnapshot 承接 BackupSnapshotRead（read / verify 子路径）。
func (s *Server) GetBackupSnapshot(ctx context.Context, req *pb.GetBackupSnapshotRequest) (*pb.GetBackupSnapshotResponse, error) {
	snapshot, err := s.querySvc.ReadBackupSnapshot(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	// 无历史备份记录 → 返回空 snapshot
	if snapshot == nil {
		return &pb.GetBackupSnapshotResponse{}, nil
	}

	return &pb.GetBackupSnapshotResponse{
		Snapshot: domainBackupSnapshotToProto(snapshot),
	}, nil
}

// CreateInstanceBackup 承接 CreateInstanceBackup 写动作。
func (s *Server) CreateInstanceBackup(ctx context.Context, req *pb.CreateInstanceBackupRequest) (*pb.CreateInstanceBackupResponse, error) {
	snapshot, err := s.commandSvc.CreateInstanceBackup(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateInstanceBackupResponse{
		Snapshot: domainBackupSnapshotToProto(snapshot),
	}, nil
}

// --- 类型转换函数 ---

func domainBackupAssetScopeToProto(s backup.BackupAssetScope) pb.BackupAssetScope {
	switch s {
	case backup.BackupAssetScopeProducts:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_PRODUCTS
	case backup.BackupAssetScopeModules:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_MODULES
	case backup.BackupAssetScopeReleases:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_RELEASES
	case backup.BackupAssetScopeRepositories:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_REPOSITORIES
	case backup.BackupAssetScopeDecisions:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_DECISIONS
	case backup.BackupAssetScopeDecisionLinks:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_DECISION_LINKS
	case backup.BackupAssetScopeProductModules:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_PRODUCT_MODULES
	case backup.BackupAssetScopeProductRepositories:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_PRODUCT_REPOSITORIES
	case backup.BackupAssetScopeModuleRepositories:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_MODULE_REPOSITORIES
	default:
		return pb.BackupAssetScope_BACKUP_ASSET_SCOPE_UNSPECIFIED
	}
}

func domainBackupVerifiedStatusToProto(s backup.BackupVerifiedStatus) pb.BackupVerifiedStatus {
	switch s {
	case backup.BackupVerifiedStatusUnverified:
		return pb.BackupVerifiedStatus_BACKUP_VERIFIED_STATUS_UNVERIFIED
	case backup.BackupVerifiedStatusVerified:
		return pb.BackupVerifiedStatus_BACKUP_VERIFIED_STATUS_VERIFIED
	case backup.BackupVerifiedStatusVerifyFailed:
		return pb.BackupVerifiedStatus_BACKUP_VERIFIED_STATUS_VERIFY_FAILED
	default:
		return pb.BackupVerifiedStatus_BACKUP_VERIFIED_STATUS_UNSPECIFIED
	}
}

func domainVerifyFailureCodeToProto(c backup.VerifyFailureCode) pb.VerifyFailureCode {
	switch c {
	case backup.VerifyFailureCodeManifestMissing:
		return pb.VerifyFailureCode_VERIFY_FAILURE_CODE_MANIFEST_MISSING
	case backup.VerifyFailureCodeCoverageIncomplete:
		return pb.VerifyFailureCode_VERIFY_FAILURE_CODE_COVERAGE_INCOMPLETE
	case backup.VerifyFailureCodeSchemaMismatch:
		return pb.VerifyFailureCode_VERIFY_FAILURE_CODE_SCHEMA_MISMATCH
	default:
		return pb.VerifyFailureCode_VERIFY_FAILURE_CODE_UNSPECIFIED
	}
}

func domainBackupSnapshotToProto(s *backup.BackupSnapshot) *pb.BackupSnapshot {
	if s == nil {
		return nil
	}
	result := &pb.BackupSnapshot{
		CreatedAt:       timestamppb.New(s.CreatedAt),
		VerifiedStatus:  domainBackupVerifiedStatusToProto(s.VerifiedStatus),
		VerifyFailureCode: domainVerifyFailureCodeToProto(s.VerifyFailureCode),
	}

	if s.ManifestSummary != nil {
		result.ManifestSummary = &pb.ManifestSummary{
			ManifestVersion:     s.ManifestSummary.ManifestVersion,
			TotalAssetEntries:   int32(s.ManifestSummary.TotalAssetEntries),
			CoveredAssetEntries: int32(s.ManifestSummary.CoveredAssetEntries),
		}
	}

	result.AssetCoverage = make([]*pb.AssetCoverageEntry, 0, len(s.AssetCoverage))
	for _, e := range s.AssetCoverage {
		result.AssetCoverage = append(result.AssetCoverage, &pb.AssetCoverageEntry{
			AssetScope: domainBackupAssetScopeToProto(e.AssetScope),
			Covered:    e.Covered,
		})
	}

	if s.SchemaVersionPrerequisite != nil {
		result.SchemaVersionPrerequisite = &pb.SchemaVersionPrerequisite{
			SchemaVersion:        s.SchemaVersionPrerequisite.SchemaVersion,
			InstanceVersion:      s.SchemaVersionPrerequisite.InstanceVersion,
			PrerequisiteCheckable: s.SchemaVersionPrerequisite.PrerequisiteCheckable,
		}
	}

	return result
}