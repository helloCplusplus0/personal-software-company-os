// Package connect — ReuseSummary Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，ReuseSummary 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/reusesummary/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/psco/backend/internal/gen/proto/psco/reuse_summary/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/reuse_summary/v1/reuse_summaryv1connect"
	"github.com/psco/backend/internal/connecterrors"
	"github.com/psco/backend/internal/reusesummary"
	"github.com/psco/backend/internal/reusesummary/service"
)

// Server 实现 ReuseSummaryServiceHandler 接口。
type Server struct {
	querySvc *service.QueryService
}

var _ pbc.ReuseSummaryServiceHandler = (*Server)(nil)

// NewServer 构造 ReuseSummary Connect handler。
func NewServer(querySvc *service.QueryService) *Server {
	return &Server{querySvc: querySvc}
}

// GetReuseSummary 承接 ReuseSummaryRead。
func (s *Server) GetReuseSummary(ctx context.Context, req *pb.GetReuseSummaryRequest) (*pb.GetReuseSummaryResponse, error) {
	scope := domainReuseSummaryScope(req.GetScope())
	moduleID := req.GetModuleId()
	productID := req.GetProductId()

	result, err := s.querySvc.ReadReuseSummary(ctx, scope, moduleID, productID)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetReuseSummaryResponse{
		ModuleReuseSummary: domainModuleReuseSummaryListToProto(result.ModuleReuseSummary),
		CapabilitySummary:  domainCapabilitySummaryListToProto(result.CapabilitySummary),
	}, nil
}

// domainReuseSummaryScope 将 proto ReuseSummaryScope 映射为 domain 类型。
func domainReuseSummaryScope(s pb.ReuseSummaryScope) reusesummary.ReuseSummaryScope {
	switch s {
	case pb.ReuseSummaryScope_REUSE_SUMMARY_SCOPE_DASHBOARD:
		return reusesummary.ReuseSummaryScopeDashboard
	case pb.ReuseSummaryScope_REUSE_SUMMARY_SCOPE_MODULE_DETAIL:
		return reusesummary.ReuseSummaryScopeModuleDetail
	case pb.ReuseSummaryScope_REUSE_SUMMARY_SCOPE_PRODUCT_DETAIL:
		return reusesummary.ReuseSummaryScopeProductDetail
	default:
		return reusesummary.ReuseSummaryScopeUnspecified
	}
}

func domainModuleReuseSummaryListToProto(items []reusesummary.ModuleReuseSummary) []*pb.ModuleReuseSummary {
	result := make([]*pb.ModuleReuseSummary, 0, len(items))
	for _, item := range items {
		mrs := &pb.ModuleReuseSummary{
			ModuleId:          item.ModuleID,
			ReuseProductCount: int32(item.ReuseProductCount),
			ExplanationText:   item.ExplanationText,
		}
		if item.LatestReuseAt != nil {
			mrs.LatestReuseAt = timestamppb.New(*item.LatestReuseAt)
		}
		result = append(result, mrs)
	}
	return result
}

func domainCapabilitySummaryListToProto(items []reusesummary.CapabilitySummary) []*pb.CapabilitySummary {
	result := make([]*pb.CapabilitySummary, 0, len(items))
	for _, item := range items {
		cs := &pb.CapabilitySummary{
			CapabilityKey:         item.CapabilityKey,
			CapabilityLabel:       item.CapabilityLabel,
			SupportingModuleCount: int32(item.SupportingModuleCount),
			EmptyStateText:        item.EmptyStateText,
		}
		if item.LatestCapabilityUpdateAt != nil {
			cs.LatestCapabilityUpdateAt = timestamppb.New(*item.LatestCapabilityUpdateAt)
		}
		result = append(result, cs)
	}
	return result
}