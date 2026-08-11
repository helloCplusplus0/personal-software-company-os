// Package connect — Export Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Export 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/export/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/psco/backend/internal/gen/proto/psco/export/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/export/v1/exportv1connect"
	"github.com/psco/backend/internal/export"
	"github.com/psco/backend/internal/export/service"
	"github.com/psco/backend/internal/connecterrors"
)

// Server 实现 ExportServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.ExportServiceHandler = (*Server)(nil)

// NewServer 构造 Export Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{querySvc: querySvc, commandSvc: commandSvc}
}

// GetExportSnapshot 承接 ExportSnapshotRead。
func (s *Server) GetExportSnapshot(ctx context.Context, req *pb.GetExportSnapshotRequest) (*pb.GetExportSnapshotResponse, error) {
	snapshot, err := s.querySvc.ReadExportSnapshot(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetExportSnapshotResponse{
		Snapshot: domainExportSnapshotToProto(snapshot),
	}, nil
}

// ExportCoreAssets 承接 ExportCoreAssets 写动作。
func (s *Server) ExportCoreAssets(ctx context.Context, req *pb.ExportCoreAssetsRequest) (*pb.ExportCoreAssetsResponse, error) {
	snapshot, err := s.commandSvc.ExportCoreAssets(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.ExportCoreAssetsResponse{
		Snapshot: domainExportSnapshotToProto(snapshot),
	}, nil
}

// --- 类型转换函数 ---

func domainExportAssetScopesToProto(scopes []export.ExportAssetScope) []pb.ExportAssetScope {
	result := make([]pb.ExportAssetScope, 0, len(scopes))
	for _, s := range scopes {
		result = append(result, domainExportAssetScopeToProto(s))
	}
	return result
}

func domainExportAssetScopeToProto(s export.ExportAssetScope) pb.ExportAssetScope {
	switch s {
	case export.ExportAssetScopeProducts:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_PRODUCTS
	case export.ExportAssetScopeModules:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_MODULES
	case export.ExportAssetScopeReleases:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_RELEASES
	case export.ExportAssetScopeRepositories:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_REPOSITORIES
	case export.ExportAssetScopeDecisions:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_DECISIONS
	case export.ExportAssetScopeDecisionLinks:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_DECISION_LINKS
	case export.ExportAssetScopeProductModules:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_PRODUCT_MODULES
	case export.ExportAssetScopeProductRepositories:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_PRODUCT_REPOSITORIES
	case export.ExportAssetScopeModuleRepositories:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_MODULE_REPOSITORIES
	default:
		return pb.ExportAssetScope_EXPORT_ASSET_SCOPE_UNSPECIFIED
	}
}

func domainExportResultStatusToProto(s export.ExportResultStatus) pb.ExportResultStatus {
	switch s {
	case export.ExportResultStatusSuccess:
		return pb.ExportResultStatus_EXPORT_RESULT_STATUS_SUCCESS
	case export.ExportResultStatusInProgress:
		return pb.ExportResultStatus_EXPORT_RESULT_STATUS_IN_PROGRESS
	case export.ExportResultStatusFailed:
		return pb.ExportResultStatus_EXPORT_RESULT_STATUS_FAILED
	default:
		return pb.ExportResultStatus_EXPORT_RESULT_STATUS_UNSPECIFIED
	}
}

func domainExportSnapshotToProto(s *export.ExportSnapshot) *pb.ExportSnapshot {
	if s == nil {
		return nil
	}
	return &pb.ExportSnapshot{
		AssetScope:    domainExportAssetScopesToProto(s.AssetScope),
		CreatedAt:     timestamppb.New(s.CreatedAt),
		ResultStatus:  domainExportResultStatusToProto(s.ResultStatus),
		ResultSummary: s.ResultSummary,
	}
}