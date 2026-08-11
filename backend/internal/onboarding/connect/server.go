// Package connect — Onboarding Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Onboarding 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
// 不包含业务逻辑，不直接访问 repository 或 candidate 层。
//
// 文件落点：backend/internal/onboarding/connect/server.go
package connect

import (
	"context"

	pb "github.com/psco/backend/internal/gen/proto/psco/onboarding/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/onboarding/v1/onboardingv1connect"
	"github.com/psco/backend/internal/onboarding/service"
	"github.com/psco/backend/internal/connecterrors"
)

// Server 实现 OnboardingServiceHandler 接口。
type Server struct {
	querySvc *service.QueryService
}

// 确保 Server 实现生成的 handler 接口。
var _ pbc.OnboardingServiceHandler = (*Server)(nil)

// NewServer 构造 Onboarding Connect handler。
func NewServer(querySvc *service.QueryService) *Server {
	return &Server{querySvc: querySvc}
}

// GetFirstRunState 承接 FirstRunStateRead。
func (s *Server) GetFirstRunState(ctx context.Context, req *pb.GetFirstRunStateRequest) (*pb.GetFirstRunStateResponse, error) {
	state, err := s.querySvc.ReadFirstRunState(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetFirstRunStateResponse{
		FirstRunState: &pb.FirstRunState{
			Status:             domainFirstRunStatusToProto(string(state.Status)),
			IsFirstEntry:       state.IsFirstEntry,
			CurrentStep:        domainOnboardingStepToProto(string(state.CurrentStep)),
			CompletionProgress: int32(state.CompletionProgress),
		},
	}, nil
}

// domainFirstRunStatusToProto 将 domain FirstRunStatus 映射为 proto 枚举。
func domainFirstRunStatusToProto(s string) pb.FirstRunStatus {
	switch s {
	case "not_started":
		return pb.FirstRunStatus_FIRST_RUN_STATUS_NOT_STARTED
	case "in_progress":
		return pb.FirstRunStatus_FIRST_RUN_STATUS_IN_PROGRESS
	case "completed":
		return pb.FirstRunStatus_FIRST_RUN_STATUS_COMPLETED
	default:
		return pb.FirstRunStatus_FIRST_RUN_STATUS_UNSPECIFIED
	}
}

// domainOnboardingStepToProto 将 domain OnboardingStep 映射为 proto 枚举。
func domainOnboardingStepToProto(s string) pb.OnboardingStep {
	switch s {
	case "welcome":
		return pb.OnboardingStep_ONBOARDING_STEP_WELCOME
	case "product":
		return pb.OnboardingStep_ONBOARDING_STEP_PRODUCT
	case "repository":
		return pb.OnboardingStep_ONBOARDING_STEP_REPOSITORY
	case "module":
		return pb.OnboardingStep_ONBOARDING_STEP_MODULE
	case "decision":
		return pb.OnboardingStep_ONBOARDING_STEP_DECISION
	case "complete":
		return pb.OnboardingStep_ONBOARDING_STEP_COMPLETE
	default:
		return pb.OnboardingStep_ONBOARDING_STEP_UNSPECIFIED
	}
}