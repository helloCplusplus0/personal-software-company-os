// Package connect — Decision Center Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Decision Center 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/decisioncenter/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/psco/backend/internal/gen/proto/psco/decision_center/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/decision_center/v1/decision_centerv1connect"
	module_registryv1 "github.com/psco/backend/internal/gen/proto/psco/module_registry/v1"
	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/decisioncenter/service"
	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/connecterrors"
)

// Server 实现 DecisionCenterServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.DecisionCenterServiceHandler = (*Server)(nil)

// NewServer 构造 Decision Center Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{querySvc: querySvc, commandSvc: commandSvc}
}

// ListDecisions 承接 DecisionListRead。
func (s *Server) ListDecisions(ctx context.Context, req *pb.ListDecisionsRequest) (*pb.ListDecisionsResponse, error) {
	items, err := s.querySvc.ListDecisions(ctx, decisioncenter.ListQuery{
		QueryText:    req.GetQueryText(),
		StatusFilter: protoDecisionStatusToDomain(req.GetStatusFilter()),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbItems := make([]*pb.DecisionListItem, 0, len(items))
	for _, item := range items {
		pbItems = append(pbItems, &pb.DecisionListItem{
			Id:                  item.ID,
			Title:               item.Title,
			Status:              domainDecisionStatusToProto(item.Status),
			CreatedAt:           timestamppb.New(item.CreatedAt),
			LinkCount:           int32(item.LinkCount),
			LinkedModuleSummary: item.LinkedModuleSummary,
		})
	}

	return &pb.ListDecisionsResponse{Decisions: pbItems}, nil
}

// GetDecisionDetail 承接 DecisionDetailRead。
func (s *Server) GetDecisionDetail(ctx context.Context, req *pb.GetDecisionDetailRequest) (*pb.GetDecisionDetailResponse, error) {
	detail, err := s.querySvc.GetDecisionDetail(ctx, req.GetDecisionId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	linkedModules := make([]*pb.LinkedModule, 0, len(detail.LinkedModules))
	for _, lm := range detail.LinkedModules {
		linkedModules = append(linkedModules, &pb.LinkedModule{
			ModuleId:   lm.ModuleID,
			ModuleName: lm.ModuleName,
		})
	}

	return &pb.GetDecisionDetailResponse{
		DecisionDetail: &pb.DecisionDetail{
			Decision: &pb.Decision{
				Id:           detail.Decision.ID,
				Title:        detail.Decision.Title,
				Context:      detail.Decision.Context,
				Problem:      detail.Decision.Problem,
				Alternatives: detail.Decision.Alternatives,
				Choice:       detail.Decision.Choice,
				Reason:       detail.Decision.Reason,
				Impact:       detail.Decision.Impact,
				Status:       domainDecisionStatusToProto(detail.Decision.Status),
				CreatedAt:    timestamppb.New(detail.Decision.CreatedAt),
			},
			LinkedModules: linkedModules,
			SourceContext: &pb.SourceContext{
				SourceModuleId:   detail.SourceContext.SourceModuleID,
				SourceModuleName: detail.SourceContext.SourceModuleName,
			},
		},
	}, nil
}

// ListDecisionModuleCandidates 承接 DecisionModuleCandidateRead。
func (s *Server) ListDecisionModuleCandidates(ctx context.Context, req *pb.ListDecisionModuleCandidatesRequest) (*pb.ListDecisionModuleCandidatesResponse, error) {
	candidates, err := s.querySvc.ListModuleCandidates(ctx, req.GetDecisionId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbCandidates := make([]*pb.DecisionModuleCandidate, 0, len(candidates))
	for _, c := range candidates {
		pbCandidates = append(pbCandidates, &pb.DecisionModuleCandidate{
			ModuleId:   c.ModuleID,
			ModuleName: c.ModuleName,
			Status:     domainModuleStatusToProto(c.Status),
		})
	}

	return &pb.ListDecisionModuleCandidatesResponse{Candidates: pbCandidates}, nil
}

// CreateDecision 承接 DecisionWrite (RecordDecision)。
func (s *Server) CreateDecision(ctx context.Context, req *pb.CreateDecisionRequest) (*pb.CreateDecisionResponse, error) {
	result, err := s.commandSvc.CreateDecision(ctx, decisioncenter.CreateDecisionRequest{
		Title:          req.GetTitle(),
		Context:        req.GetContext(),
		Problem:        req.GetProblem(),
		Alternatives:   req.GetAlternatives(),
		Choice:         req.GetChoice(),
		Reason:         req.GetReason(),
		Impact:         req.GetImpact(),
		Status:         protoDecisionStatusToDomain(req.GetStatus()),
		SourceModuleID: req.GetSourceModuleId(),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateDecisionResponse{DecisionId: result.DecisionID}, nil
}

// LinkDecisionToTarget 承接 DecisionLinkWrite (LinkDecisionToTarget)。
func (s *Server) LinkDecisionToTarget(ctx context.Context, req *pb.LinkDecisionToTargetRequest) (*pb.LinkDecisionToTargetResponse, error) {
	err := s.commandSvc.LinkDecisionToTarget(ctx, req.GetDecisionId(), decisioncenter.LinkDecisionToTargetRequest{
		TargetType: decisioncenter.DecisionLinkTargetTypeModule,
		ModuleID:   req.GetModuleId(),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.LinkDecisionToTargetResponse{}, nil
}

// --- 类型转换函数 ---

func protoDecisionStatusToDomain(s pb.DecisionStatus) decisioncenter.DecisionStatus {
	switch s {
	case pb.DecisionStatus_DECISION_STATUS_PROPOSED:
		return decisioncenter.DecisionStatusProposed
	case pb.DecisionStatus_DECISION_STATUS_ACTIVE:
		return decisioncenter.DecisionStatusActive
	case pb.DecisionStatus_DECISION_STATUS_SUPERSEDED:
		return decisioncenter.DecisionStatusSuperseded
	case pb.DecisionStatus_DECISION_STATUS_ARCHIVED:
		return decisioncenter.DecisionStatusArchived
	default:
		return ""
	}
}

func domainDecisionStatusToProto(s decisioncenter.DecisionStatus) pb.DecisionStatus {
	switch s {
	case decisioncenter.DecisionStatusProposed:
		return pb.DecisionStatus_DECISION_STATUS_PROPOSED
	case decisioncenter.DecisionStatusActive:
		return pb.DecisionStatus_DECISION_STATUS_ACTIVE
	case decisioncenter.DecisionStatusSuperseded:
		return pb.DecisionStatus_DECISION_STATUS_SUPERSEDED
	case decisioncenter.DecisionStatusArchived:
		return pb.DecisionStatus_DECISION_STATUS_ARCHIVED
	default:
		return pb.DecisionStatus_DECISION_STATUS_UNSPECIFIED
	}
}

func domainModuleStatusToProto(s moduleregistry.ModuleStatus) module_registryv1.ModuleStatus {
	switch s {
	case moduleregistry.ModuleStatusActive:
		return module_registryv1.ModuleStatus_MODULE_STATUS_ACTIVE
	case moduleregistry.ModuleStatusArchived:
		return module_registryv1.ModuleStatus_MODULE_STATUS_ARCHIVED
	default:
		return module_registryv1.ModuleStatus_MODULE_STATUS_UNSPECIFIED
	}
}