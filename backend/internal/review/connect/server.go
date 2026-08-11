// Package connect — Review Connect transport 实现。
//
// 本文件是 phase08-08 正式传输主线落地后，Review 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/review/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/psco/backend/internal/gen/proto/psco/review/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/review/v1/reviewv1connect"
	"github.com/psco/backend/internal/connecterrors"
	"github.com/psco/backend/internal/dashboard"
	dcpb "github.com/psco/backend/internal/gen/proto/psco/decision_center/v1"
	ddpb "github.com/psco/backend/internal/gen/proto/psco/dashboard/v1"
	rspb "github.com/psco/backend/internal/gen/proto/psco/reuse_summary/v1"
	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/reusesummary"
	"github.com/psco/backend/internal/review"
	"github.com/psco/backend/internal/review/service"
)

// Server 实现 ReviewServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.ReviewServiceHandler = (*Server)(nil)

// NewServer 构造 Review Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{
		querySvc:   querySvc,
		commandSvc: commandSvc,
	}
}

// GetDailyReviewContext 承接 DailyReviewContextRead。
func (s *Server) GetDailyReviewContext(ctx context.Context, req *pb.GetDailyReviewContextRequest) (*pb.GetDailyReviewContextResponse, error) {
	result, err := s.querySvc.GetDailyReviewContext(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetDailyReviewContextResponse{
		Context: &pb.DailyReviewContext{
			CurrentFocusSignals:   domainFeedbackSignalsToProto(result.CurrentFocusSignals),
			RepresentativeSignals: domainFeedbackSignalsToProto(result.RepresentativeSignals),
			PendingDecisions:      domainDecisionListItemsToProto(result.PendingDecisions),
		},
	}, nil
}

// GetWeeklyReviewContext 承接 WeeklyReviewContextRead。
func (s *Server) GetWeeklyReviewContext(ctx context.Context, req *pb.GetWeeklyReviewContextRequest) (*pb.GetWeeklyReviewContextResponse, error) {
	result, err := s.querySvc.GetWeeklyReviewContext(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetWeeklyReviewContextResponse{
		Context: &pb.WeeklyReviewContext{
			Overview:              domainOverviewToProto(result.Overview),
			RecentActivities:      domainRecentActivityItemsToProto(result.RecentActivities),
			RepresentativeSignals: domainFeedbackSignalsToProto(result.RepresentativeSignals),
			ModuleReuseSummary:    domainModuleReuseSummariesToProto(result.ModuleReuseSummary),
			CapabilitySummary:     domainCapabilitySummariesToProto(result.CapabilitySummary),
		},
	}, nil
}

// SubmitReviewResult 承接 review 结果提交（仅 next-step result 路径）。
func (s *Server) SubmitReviewResult(ctx context.Context, req *pb.SubmitReviewResultRequest) (*pb.SubmitReviewResultResponse, error) {
	input := review.SubmitReviewResultInput{
		ReviewKind:  review.ReviewKind(domainReviewKindToString(req.ReviewKind)),
		ResultKind:  review.ReviewResultKind(domainReviewResultKindToString(req.ResultKind)),
		DecisionID:  req.DecisionId,
		TargetType:  req.TargetType,
		TargetID:    req.TargetId,
		SummaryText: req.SummaryText,
		StartedAt:   req.StartedAt.AsTime(),
		CompletedAt: req.CompletedAt.AsTime(),
	}

	output, err := s.commandSvc.SubmitReviewResult(ctx, input)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.SubmitReviewResultResponse{
		ReviewRecordId: output.ReviewRecordID,
		ResultKind:     stringToReviewResultKindEnum(string(output.ResultKind)),
		DecisionId:     output.DecisionID,
		TargetType:     output.TargetType,
		TargetId:       output.TargetID,
	}, nil
}

// --- 类型转换函数 ---

func domainReviewKindToString(k pb.ReviewKind) string {
	switch k {
	case pb.ReviewKind_REVIEW_KIND_DAILY:
		return "daily"
	case pb.ReviewKind_REVIEW_KIND_WEEKLY:
		return "weekly"
	default:
		return ""
	}
}

func domainReviewResultKindToString(k pb.ReviewResultKind) string {
	switch k {
	case pb.ReviewResultKind_REVIEW_RESULT_KIND_DECISION_HANDOFF:
		return "decision_handoff"
	case pb.ReviewResultKind_REVIEW_RESULT_KIND_ENTITY_HANDOFF:
		return "entity_handoff"
	case pb.ReviewResultKind_REVIEW_RESULT_KIND_NEXT_STEP_RESULT:
		return "next_step_result"
	default:
		return ""
	}
}

func stringToReviewResultKindEnum(s string) pb.ReviewResultKind {
	switch s {
	case "decision_handoff":
		return pb.ReviewResultKind_REVIEW_RESULT_KIND_DECISION_HANDOFF
	case "entity_handoff":
		return pb.ReviewResultKind_REVIEW_RESULT_KIND_ENTITY_HANDOFF
	case "next_step_result":
		return pb.ReviewResultKind_REVIEW_RESULT_KIND_NEXT_STEP_RESULT
	default:
		return pb.ReviewResultKind_REVIEW_RESULT_KIND_UNSPECIFIED
	}
}

func domainOverviewToProto(o *dashboard.DashboardOverview) *ddpb.DashboardOverview {
	if o == nil {
		return nil
	}
	return &ddpb.DashboardOverview{
		ModuleCount:                int32(o.ModuleCount),
		ProductCount:               int32(o.ProductCount),
		RepositoryCount:            int32(o.RepositoryCount),
		DecisionCount:              int32(o.DecisionCount),
		ProductWithRepositoryCount: int32(o.ProductWithRepositoryCount),
		ProductWithModuleCount:     int32(o.ProductWithModuleCount),
	}
}

func domainFeedbackSignalsToProto(signals []dashboard.FeedbackSignal) []*ddpb.FeedbackSignal {
	result := make([]*ddpb.FeedbackSignal, 0, len(signals))
	for _, sig := range signals {
		result = append(result, &ddpb.FeedbackSignal{
			SignalFamily: domainFeedbackSignalFamilyToProto(sig.SignalFamily),
			SignalCode:   domainFeedbackSignalCodeToProto(sig.SignalCode),
			Priority:     domainFeedbackSignalPriorityToProto(sig.Priority),
			Title:        sig.Title,
			Summary:      sig.Summary,
			ActionLabel:  sig.ActionLabel,
			TargetType:   domainDashboardTargetTypeToProto(sig.TargetType),
			TargetId:     sig.TargetID,
			TargetLabel:  sig.TargetLabel,
		})
	}
	return result
}

func domainRecentActivityItemsToProto(items []dashboard.RecentActivityItem) []*ddpb.RecentActivityItem {
	result := make([]*ddpb.RecentActivityItem, 0, len(items))
	for _, item := range items {
		result = append(result, &ddpb.RecentActivityItem{
			ActivityType: domainRecentActivityTypeToProto(item.ActivityType),
			ActivityAt:   timestamppb.New(item.ActivityAt),
			TargetType:   domainDashboardTargetTypeToProto(item.TargetType),
			TargetId:     item.TargetID,
			TargetLabel:  item.TargetLabel,
		})
	}
	return result
}

func domainDecisionListItemsToProto(items []decisioncenter.DecisionListItem) []*dcpb.DecisionListItem {
	result := make([]*dcpb.DecisionListItem, 0, len(items))
	for _, item := range items {
		result = append(result, &dcpb.DecisionListItem{
			Id:                  item.ID,
			Title:               item.Title,
			Status:              domainDecisionStatusToProto(item.Status),
			CreatedAt:           timestamppb.New(item.CreatedAt),
			LinkCount:           int32(item.LinkCount),
			LinkedModuleSummary: item.LinkedModuleSummary,
		})
	}
	return result
}

func domainModuleReuseSummariesToProto(items []reusesummary.ModuleReuseSummary) []*rspb.ModuleReuseSummary {
	result := make([]*rspb.ModuleReuseSummary, 0, len(items))
	for _, item := range items {
		result = append(result, &rspb.ModuleReuseSummary{
			ModuleId:          item.ModuleID,
			ReuseProductCount: int32(item.ReuseProductCount),
			LatestReuseAt:     timestamppb.New(*item.LatestReuseAt),
			ExplanationText:   item.ExplanationText,
		})
	}
	return result
}

func domainCapabilitySummariesToProto(items []reusesummary.CapabilitySummary) []*rspb.CapabilitySummary {
	result := make([]*rspb.CapabilitySummary, 0, len(items))
	for _, item := range items {
		result = append(result, &rspb.CapabilitySummary{
			CapabilityKey:              item.CapabilityKey,
			CapabilityLabel:            item.CapabilityLabel,
			SupportingModuleCount:      int32(item.SupportingModuleCount),
			LatestCapabilityUpdateAt:   timestamppb.New(*item.LatestCapabilityUpdateAt),
			EmptyStateText:             item.EmptyStateText,
		})
	}
	return result
}

// --- 枚举转换辅助函数 ---

func domainFeedbackSignalFamilyToProto(f dashboard.FeedbackSignalFamily) ddpb.FeedbackSignalFamily {
	switch f {
	case dashboard.FeedbackSignalFamilyPendingDecision:
		return ddpb.FeedbackSignalFamily_FEEDBACK_SIGNAL_FAMILY_PENDING_DECISION
	case dashboard.FeedbackSignalFamilyProductAssetCoverage:
		return ddpb.FeedbackSignalFamily_FEEDBACK_SIGNAL_FAMILY_PRODUCT_ASSET_COVERAGE
	default:
		return ddpb.FeedbackSignalFamily_FEEDBACK_SIGNAL_FAMILY_UNSPECIFIED
	}
}

func domainFeedbackSignalCodeToProto(c dashboard.FeedbackSignalCode) ddpb.FeedbackSignalCode {
	switch c {
	case dashboard.FeedbackSignalCodePendingDecision:
		return ddpb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PENDING_DECISION
	case dashboard.FeedbackSignalCodeProductMissingBothBindings:
		return ddpb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS
	case dashboard.FeedbackSignalCodeProductMissingRepositoryBinding:
		return ddpb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING
	case dashboard.FeedbackSignalCodeProductMissingModuleBinding:
		return ddpb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING
	default:
		return ddpb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_UNSPECIFIED
	}
}

func domainFeedbackSignalPriorityToProto(p dashboard.FeedbackSignalPriority) ddpb.FeedbackSignalPriority {
	switch p {
	case dashboard.FeedbackSignalPriorityP1PendingDecision:
		return ddpb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P1_PENDING_DECISION
	case dashboard.FeedbackSignalPriorityP2ProductMissingBothBindings:
		return ddpb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P2_PRODUCT_MISSING_BOTH_BINDINGS
	case dashboard.FeedbackSignalPriorityP3ProductMissingRepository:
		return ddpb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P3_PRODUCT_MISSING_REPOSITORY_BINDING
	case dashboard.FeedbackSignalPriorityP4ProductMissingModule:
		return ddpb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P4_PRODUCT_MISSING_MODULE_BINDING
	default:
		return ddpb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_UNSPECIFIED
	}
}

func domainDashboardTargetTypeToProto(t dashboard.DashboardTargetType) ddpb.DashboardTargetType {
	switch t {
	case dashboard.DashboardTargetTypeDecisionDetail:
		return ddpb.DashboardTargetType_DASHBOARD_TARGET_TYPE_DECISION_DETAIL
	case dashboard.DashboardTargetTypeDecisionList:
		return ddpb.DashboardTargetType_DASHBOARD_TARGET_TYPE_DECISION_LIST
	case dashboard.DashboardTargetTypeProductDetail:
		return ddpb.DashboardTargetType_DASHBOARD_TARGET_TYPE_PRODUCT_DETAIL
	case dashboard.DashboardTargetTypeModuleDetail:
		return ddpb.DashboardTargetType_DASHBOARD_TARGET_TYPE_MODULE_DETAIL
	case dashboard.DashboardTargetTypeRepositoryDetail:
		return ddpb.DashboardTargetType_DASHBOARD_TARGET_TYPE_REPOSITORY_DETAIL
	default:
		return ddpb.DashboardTargetType_DASHBOARD_TARGET_TYPE_UNSPECIFIED
	}
}

func domainRecentActivityTypeToProto(t dashboard.RecentActivityType) ddpb.RecentActivityType {
	switch t {
	case dashboard.RecentActivityTypeModule:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_MODULE
	case dashboard.RecentActivityTypeRelease:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_RELEASE
	case dashboard.RecentActivityTypeProduct:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_PRODUCT
	case dashboard.RecentActivityTypeRepository:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_REPOSITORY
	case dashboard.RecentActivityTypeDecision:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_DECISION
	case dashboard.RecentActivityTypeProductModuleBinding:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_PRODUCT_MODULE_BINDING
	case dashboard.RecentActivityTypeProductRepositoryBinding:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_PRODUCT_REPOSITORY_BINDING
	case dashboard.RecentActivityTypeModuleRepositoryBinding:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_MODULE_REPOSITORY_BINDING
	default:
		return ddpb.RecentActivityType_RECENT_ACTIVITY_TYPE_UNSPECIFIED
	}
}

func domainDecisionStatusToProto(s decisioncenter.DecisionStatus) dcpb.DecisionStatus {
	switch s {
	case decisioncenter.DecisionStatusProposed:
		return dcpb.DecisionStatus_DECISION_STATUS_PROPOSED
	case decisioncenter.DecisionStatusActive:
		return dcpb.DecisionStatus_DECISION_STATUS_ACTIVE
	case decisioncenter.DecisionStatusSuperseded:
		return dcpb.DecisionStatus_DECISION_STATUS_SUPERSEDED
	case decisioncenter.DecisionStatusArchived:
		return dcpb.DecisionStatus_DECISION_STATUS_ARCHIVED
	default:
		return dcpb.DecisionStatus_DECISION_STATUS_UNSPECIFIED
	}
}