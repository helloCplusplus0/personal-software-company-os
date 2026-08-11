// Package connect — Dashboard Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Dashboard 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/dashboard/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/psco/backend/internal/gen/proto/psco/dashboard/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/dashboard/v1/dashboardv1connect"
	"github.com/psco/backend/internal/dashboard"
	"github.com/psco/backend/internal/dashboard/service"
	"github.com/psco/backend/internal/connecterrors"
)

// Server 实现 DashboardServiceHandler 接口。
type Server struct {
	querySvc *service.QueryService
}

var _ pbc.DashboardServiceHandler = (*Server)(nil)

// NewServer 构造 Dashboard Connect handler。
func NewServer(querySvc *service.QueryService) *Server {
	return &Server{querySvc: querySvc}
}

// GetDashboardOverview 承接 DashboardOverviewRead。
func (s *Server) GetDashboardOverview(ctx context.Context, req *pb.GetDashboardOverviewRequest) (*pb.GetDashboardOverviewResponse, error) {
	overview, err := s.querySvc.ReadOverview(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetDashboardOverviewResponse{
		Overview: &pb.DashboardOverview{
			ModuleCount:                int32(overview.ModuleCount),
			ProductCount:               int32(overview.ProductCount),
			RepositoryCount:            int32(overview.RepositoryCount),
			DecisionCount:              int32(overview.DecisionCount),
			ProductWithRepositoryCount: int32(overview.ProductWithRepositoryCount),
			ProductWithModuleCount:     int32(overview.ProductWithModuleCount),
		},
	}, nil
}

// GetFeedbackSignals 承接 FeedbackSignalRead。
func (s *Server) GetFeedbackSignals(ctx context.Context, req *pb.GetFeedbackSignalsRequest) (*pb.GetFeedbackSignalsResponse, error) {
	result, err := s.querySvc.ReadFeedbackSignal(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetFeedbackSignalsResponse{
		CurrentFocusSignals:  domainFeedbackSignalsToProto(result.CurrentFocusSignals),
		AssetFeedbackSummary: domainAssetCoverageSummaryToProto(&result.AssetFeedbackSummary),
	}, nil
}

// GetRecentActivities 承接 RecentActivityRead。
func (s *Server) GetRecentActivities(ctx context.Context, req *pb.GetRecentActivitiesRequest) (*pb.GetRecentActivitiesResponse, error) {
	items, err := s.querySvc.ReadRecentActivity(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetRecentActivitiesResponse{
		Activities: domainRecentActivityItemsToProto(items),
	}, nil
}

// --- 类型转换函数 ---

func domainFeedbackSignalsToProto(signals []dashboard.FeedbackSignal) []*pb.FeedbackSignal {
	result := make([]*pb.FeedbackSignal, 0, len(signals))
	for _, sig := range signals {
		result = append(result, &pb.FeedbackSignal{
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

func domainAssetCoverageSummaryToProto(s *dashboard.ProductAssetCoverageSummary) *pb.ProductAssetCoverageSummary {
	if s == nil {
		return nil
	}
	return &pb.ProductAssetCoverageSummary{
		FullyBoundProductCount:        int32(s.FullyBoundProductCount),
		MissingBothBindingsCount:      int32(s.MissingBothBindingsCount),
		MissingRepositoryBindingCount: int32(s.MissingRepositoryBindingCount),
		MissingModuleBindingCount:     int32(s.MissingModuleBindingCount),
		RepresentativeSignals:         domainFeedbackSignalsToProto(s.RepresentativeSignals),
	}
}

func domainRecentActivityItemsToProto(items []dashboard.RecentActivityItem) []*pb.RecentActivityItem {
	result := make([]*pb.RecentActivityItem, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.RecentActivityItem{
			ActivityType: domainRecentActivityTypeToProto(item.ActivityType),
			ActivityAt:   timestamppb.New(item.ActivityAt),
			TargetType:   domainDashboardTargetTypeToProto(item.TargetType),
			TargetId:     item.TargetID,
			TargetLabel:  item.TargetLabel,
		})
	}
	return result
}

func domainFeedbackSignalFamilyToProto(f dashboard.FeedbackSignalFamily) pb.FeedbackSignalFamily {
	switch f {
	case dashboard.FeedbackSignalFamilyPendingDecision:
		return pb.FeedbackSignalFamily_FEEDBACK_SIGNAL_FAMILY_PENDING_DECISION
	case dashboard.FeedbackSignalFamilyProductAssetCoverage:
		return pb.FeedbackSignalFamily_FEEDBACK_SIGNAL_FAMILY_PRODUCT_ASSET_COVERAGE
	default:
		return pb.FeedbackSignalFamily_FEEDBACK_SIGNAL_FAMILY_UNSPECIFIED
	}
}

func domainFeedbackSignalCodeToProto(c dashboard.FeedbackSignalCode) pb.FeedbackSignalCode {
	switch c {
	case dashboard.FeedbackSignalCodePendingDecision:
		return pb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PENDING_DECISION
	case dashboard.FeedbackSignalCodeProductMissingBothBindings:
		return pb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_BOTH_BINDINGS
	case dashboard.FeedbackSignalCodeProductMissingRepositoryBinding:
		return pb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_REPOSITORY_BINDING
	case dashboard.FeedbackSignalCodeProductMissingModuleBinding:
		return pb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_PRODUCT_MISSING_MODULE_BINDING
	default:
		return pb.FeedbackSignalCode_FEEDBACK_SIGNAL_CODE_UNSPECIFIED
	}
}

func domainFeedbackSignalPriorityToProto(p dashboard.FeedbackSignalPriority) pb.FeedbackSignalPriority {
	switch p {
	case dashboard.FeedbackSignalPriorityP1PendingDecision:
		return pb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P1_PENDING_DECISION
	case dashboard.FeedbackSignalPriorityP2ProductMissingBothBindings:
		return pb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P2_PRODUCT_MISSING_BOTH_BINDINGS
	case dashboard.FeedbackSignalPriorityP3ProductMissingRepository:
		return pb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P3_PRODUCT_MISSING_REPOSITORY_BINDING
	case dashboard.FeedbackSignalPriorityP4ProductMissingModule:
		return pb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_P4_PRODUCT_MISSING_MODULE_BINDING
	default:
		return pb.FeedbackSignalPriority_FEEDBACK_SIGNAL_PRIORITY_UNSPECIFIED
	}
}

func domainDashboardTargetTypeToProto(t dashboard.DashboardTargetType) pb.DashboardTargetType {
	switch t {
	case dashboard.DashboardTargetTypeDecisionDetail:
		return pb.DashboardTargetType_DASHBOARD_TARGET_TYPE_DECISION_DETAIL
	case dashboard.DashboardTargetTypeDecisionList:
		return pb.DashboardTargetType_DASHBOARD_TARGET_TYPE_DECISION_LIST
	case dashboard.DashboardTargetTypeProductDetail:
		return pb.DashboardTargetType_DASHBOARD_TARGET_TYPE_PRODUCT_DETAIL
	case dashboard.DashboardTargetTypeModuleDetail:
		return pb.DashboardTargetType_DASHBOARD_TARGET_TYPE_MODULE_DETAIL
	case dashboard.DashboardTargetTypeRepositoryDetail:
		return pb.DashboardTargetType_DASHBOARD_TARGET_TYPE_REPOSITORY_DETAIL
	default:
		return pb.DashboardTargetType_DASHBOARD_TARGET_TYPE_UNSPECIFIED
	}
}

func domainRecentActivityTypeToProto(t dashboard.RecentActivityType) pb.RecentActivityType {
	switch t {
	case dashboard.RecentActivityTypeModule:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_MODULE
	case dashboard.RecentActivityTypeRelease:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_RELEASE
	case dashboard.RecentActivityTypeProduct:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_PRODUCT
	case dashboard.RecentActivityTypeRepository:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_REPOSITORY
	case dashboard.RecentActivityTypeDecision:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_DECISION
	case dashboard.RecentActivityTypeProductModuleBinding:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_PRODUCT_MODULE_BINDING
	case dashboard.RecentActivityTypeProductRepositoryBinding:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_PRODUCT_REPOSITORY_BINDING
	case dashboard.RecentActivityTypeModuleRepositoryBinding:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_MODULE_REPOSITORY_BINDING
	default:
		return pb.RecentActivityType_RECENT_ACTIVITY_TYPE_UNSPECIFIED
	}
}