// Package connect — governanceprofile Connect transport 实现。
//
// 本文件是 phase13 治理画像后端主线的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/governanceprofile/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/psco/backend/internal/connecterrors"
	pbc "github.com/psco/backend/internal/gen/connect/psco/governance_profile/v1/governance_profilev1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/governance_profile/v1"
	"github.com/psco/backend/internal/governanceprofile"
	"github.com/psco/backend/internal/governanceprofile/service"
)

// Server 实现 GovernanceProfileServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.GovernanceProfileServiceHandler = (*Server)(nil)

// NewServer 构造 GovernanceProfile Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{
		querySvc:   querySvc,
		commandSvc: commandSvc,
	}
}

// GetGovernanceProfile 承接治理画像结构化读取。
func (s *Server) GetGovernanceProfile(ctx context.Context, req *pb.GetGovernanceProfileRequest) (*pb.GetGovernanceProfileResponse, error) {
	result, err := s.querySvc.GetGovernanceProfile(ctx, req.RepositoryId)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetGovernanceProfileResponse{
		Profile: domainResultToProto(result),
	}, nil
}

// UpdateGovernanceProfile 承接治理画像保存（手工维护优先，单一事务边界）。
//
// 解包约束：请求不携带 read-only 字段（track_type / current_phase_*）与
// project_profile_version（由服务端固定写入），字段排除在合同层即成立。
func (s *Server) UpdateGovernanceProfile(ctx context.Context, req *pb.UpdateGovernanceProfileRequest) (*pb.UpdateGovernanceProfileResponse, error) {
	input := governanceprofile.UpdateGovernanceProfileInput{
		RepositoryID:        req.RepositoryId,
		TemplateSource:      req.TemplateSource,
		DocsWorkflowLayout:  req.DocsWorkflowLayout,
		CanonicalRootFiles:  protoRootFilesToDomain(req.CanonicalRootFiles),
		GlobalAssetBindings: protoAssetBindingsToDomain(req.GlobalAssetBindings),
	}

	result, err := s.commandSvc.UpdateGovernanceProfile(ctx, input)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.UpdateGovernanceProfileResponse{
		Profile: domainResultToProto(result),
	}, nil
}

// --- 领域结果 → proto 组装 ---

func domainResultToProto(result *governanceprofile.GovernanceProfileReadResult) *pb.GovernanceProfile {
	rootFiles := make([]*pb.CanonicalRootFileBinding, 0, len(result.CanonicalRootFiles))
	for _, f := range result.CanonicalRootFiles {
		rootFiles = append(rootFiles, &pb.CanonicalRootFileBinding{
			FileName: f.FileName,
			Role:     f.Role,
			Required: f.Required,
		})
	}

	assetBindings := make([]*pb.GlobalAssetBinding, 0, len(result.GlobalAssetBindings))
	for _, b := range result.GlobalAssetBindings {
		assetBindings = append(assetBindings, &pb.GlobalAssetBinding{
			Name:               b.Name,
			Kind:               b.Kind,
			EntryRef:           b.EntryRef,
			Role:               b.Role,
			StructuredSummary:  b.StructuredSummary,
			MarkdownResolvable: boolPtr(b.MarkdownResolvable),
		})
	}

	return &pb.GovernanceProfile{
		RepositoryId:          result.Record.RepositoryID,
		ProjectProfileVersion: result.Record.ProjectProfileVersion,
		TrackType:             domainTrackTypeToProto(result.Record.TrackType),
		TemplateSource:        result.Record.TemplateSource,
		DocsWorkflowLayout:    result.Record.DocsWorkflowLayout,
		CurrentPhaseName:      result.Record.CurrentPhaseName,
		CurrentPhaseRef:       result.Record.CurrentPhaseRef,
		CurrentPhaseStatus:    domainPhaseStatusToProto(result.Record.CurrentPhaseStatus),
		CanonicalRootFiles:    rootFiles,
		GlobalAssetBindings:   assetBindings,
		CreatedAt:             timestamppb.New(result.Record.CreatedAt),
		UpdatedAt:             timestamppb.New(result.Record.UpdatedAt),
	}
}

// --- proto 请求 → 领域输入解包 ---

func protoRootFilesToDomain(files []*pb.CanonicalRootFileBinding) []governanceprofile.CanonicalRootFileBinding {
	result := make([]governanceprofile.CanonicalRootFileBinding, 0, len(files))
	for _, f := range files {
		result = append(result, governanceprofile.CanonicalRootFileBinding{
			FileName: f.GetFileName(),
			Role:     f.GetRole(),
			Required: f.GetRequired(),
		})
	}
	return result
}

func protoAssetBindingsToDomain(bindings []*pb.GlobalAssetBinding) []governanceprofile.GlobalAssetBinding {
	result := make([]governanceprofile.GlobalAssetBinding, 0, len(bindings))
	for _, b := range bindings {
		result = append(result, governanceprofile.GlobalAssetBinding{
			Name:              b.GetName(),
			Kind:              b.GetKind(),
			EntryRef:          b.GetEntryRef(),
			Role:              b.GetRole(),
			StructuredSummary: b.StructuredSummary,
		})
	}
	return result
}

func boolPtr(v bool) *bool {
	return &v
}

// --- 枚举转换 ---

func domainTrackTypeToProto(t governanceprofile.TrackType) pb.TrackType {
	switch t {
	case governanceprofile.TrackTypeProduct:
		return pb.TrackType_TRACK_TYPE_PRODUCT
	case governanceprofile.TrackTypeDurableSystem:
		return pb.TrackType_TRACK_TYPE_DURABLE_SYSTEM
	default:
		return pb.TrackType_TRACK_TYPE_UNSPECIFIED
	}
}

func domainPhaseStatusToProto(s governanceprofile.PhaseStatus) pb.PhaseStatus {
	switch s {
	case governanceprofile.PhaseStatusPlanned:
		return pb.PhaseStatus_PHASE_STATUS_PLANNED
	case governanceprofile.PhaseStatusInProgress:
		return pb.PhaseStatus_PHASE_STATUS_IN_PROGRESS
	case governanceprofile.PhaseStatusCompleted:
		return pb.PhaseStatus_PHASE_STATUS_COMPLETED
	case governanceprofile.PhaseStatusBlocked:
		return pb.PhaseStatus_PHASE_STATUS_BLOCKED
	default:
		return pb.PhaseStatus_PHASE_STATUS_UNSPECIFIED
	}
}
