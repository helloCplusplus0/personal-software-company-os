// Package connect — Template Reuse Connect transport 实现。
//
// 本文件是 phase09-08 正式传输主线落地后，Template Reuse 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/templatereuse/connect/server.go
package connect

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/psco/backend/internal/connecterrors"
	pbc "github.com/psco/backend/internal/gen/connect/psco/template_reuse/v1/template_reusev1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/template_reuse/v1"
	"github.com/psco/backend/internal/templatereuse"
	"github.com/psco/backend/internal/templatereuse/candidate"
	"github.com/psco/backend/internal/templatereuse/service"
)

// Server 实现 TemplateReuseServiceHandler 接口。
type Server struct {
	querySvc queryServicer
}

var _ pbc.TemplateReuseServiceHandler = (*Server)(nil)

type queryServicer interface {
	ListTemplateCandidates(ctx context.Context) (*service.ListTemplateCandidatesResult, error)
	GetTemplateCandidatePrefill(ctx context.Context, templateCandidateID string) (*service.TemplateCandidatePrefillResult, error)
	GetDerivedInsightHints(ctx context.Context, templateCandidateID, consumerSurface, reviewScopeKey string) (*service.DerivedInsightHintsResult, error)
	GetTemplateSourceSummary(ctx context.Context, templateCandidateID, templateSource string) (*service.TemplateSourceSummaryResult, error)
}

// NewServer 构造 TemplateReuse Connect handler。
func NewServer(querySvc queryServicer) *Server {
	return &Server{querySvc: querySvc}
}

// ListTemplateCandidates 承接模板候选列表读取。
func (s *Server) ListTemplateCandidates(ctx context.Context, req *pb.ListTemplateCandidatesRequest) (*pb.ListTemplateCandidatesResponse, error) {
	if err := validateConsumerSurface(req.GetConsumerSurface(), pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	result, err := s.querySvc.ListTemplateCandidates(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	candidates := make([]*pb.TemplateCandidateSummary, 0, len(result.Candidates))
	for _, c := range result.Candidates {
		candidates = append(candidates, candidateDataToProto(&c))
	}

	return &pb.ListTemplateCandidatesResponse{
		Candidates:               candidates,
		DefaultActiveCandidateId: result.DefaultActiveCandidateID,
	}, nil
}

// GetTemplateCandidatePrefill 承接模板预填详情读取。
func (s *Server) GetTemplateCandidatePrefill(ctx context.Context, req *pb.GetTemplateCandidatePrefillRequest) (*pb.GetTemplateCandidatePrefillResponse, error) {
	if err := validateConsumerSurface(req.GetConsumerSurface(), pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_CREATE); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	result, err := s.querySvc.GetTemplateCandidatePrefill(ctx, req.TemplateCandidateId)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	prefill := &pb.TemplateCandidatePrefill{
		TemplateCandidateId:         result.TemplateCandidateID,
		ResolutionStatus:            resolutionStatusToProto(result.ResolutionStatus),
		UnavailableReasonText:       result.UnavailableReasonText,
		TemplateTitle:               result.TemplateTitle,
		TemplateDescription:         result.TemplateDescription,
		SuggestedProductName:        result.SuggestedProductName,
		SuggestedProductDescription: result.SuggestedProductDescription,
		Modules:                     moduleRefsDataToProto(result.Modules),
		CapabilityGapHints:          derivedHintsToProto(result.CapabilityGapHints),
	}

	return &pb.GetTemplateCandidatePrefillResponse{
		Prefill: prefill,
	}, nil
}

// GetDerivedInsightHints 承接派生提示读取。
func (s *Server) GetDerivedInsightHints(ctx context.Context, req *pb.GetDerivedInsightHintsRequest) (*pb.GetDerivedInsightHintsResponse, error) {
	if err := validateDerivedHintsRequest(req); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	result, err := s.querySvc.GetDerivedInsightHints(
		ctx,
		req.TemplateCandidateId,
		consumerSurfaceToString(req.GetConsumerSurface()),
		req.GetReviewScopeKey(),
	)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetDerivedInsightHintsResponse{
		ResolutionStatus:      resolutionStatusToProto(result.ResolutionStatus),
		UnavailableReasonText: result.UnavailableReasonText,
		Hints:                 derivedHintsToProto(result.Hints),
	}, nil
}

// GetTemplateSourceSummary 承接模板来源复读。
func (s *Server) GetTemplateSourceSummary(ctx context.Context, req *pb.GetTemplateSourceSummaryRequest) (*pb.GetTemplateSourceSummaryResponse, error) {
	if err := validateConsumerSurface(req.GetConsumerSurface(), pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_DETAIL); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	result, err := s.querySvc.GetTemplateSourceSummary(ctx, req.TemplateCandidateId, templateSourceToString(req.TemplateSource))
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	summary := &pb.TemplateSourceSummary{
		TemplateCandidateId:   result.TemplateCandidateID,
		ResolutionStatus:      resolutionStatusToProto(result.ResolutionStatus),
		UnavailableReasonText: result.UnavailableReasonText,
		TemplateTitle:         result.TemplateTitle,
		TemplateDescription:   result.TemplateDescription,
		Modules:               moduleRefsDataToProto(result.Modules),
		TemplateSource:        templateSourceToProto(result.TemplateSource),
	}

	return &pb.GetTemplateSourceSummaryResponse{
		SourceSummary: summary,
	}, nil
}

func validateConsumerSurface(actual, expected pb.TemplateConsumerSurface) error {
	if actual != expected {
		return fmt.Errorf("template reuse: consumer surface %s is invalid for this RPC: %w", actual.String(), templatereuse.ErrInvalidInput)
	}
	return nil
}

func validateDerivedHintsRequest(req *pb.GetDerivedInsightHintsRequest) error {
	switch req.GetConsumerSurface() {
	case pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW:
		if req.GetReviewScopeKey() == "" {
			return fmt.Errorf("template reuse: weekly review hints require review_scope_key: %w", templatereuse.ErrInvalidInput)
		}
	case pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_CREATE:
		if req.GetReviewScopeKey() != "" {
			return fmt.Errorf("template reuse: product create hints must not carry review_scope_key: %w", templatereuse.ErrInvalidInput)
		}
	default:
		return fmt.Errorf("template reuse: consumer surface %s is invalid for derived hints: %w", req.GetConsumerSurface().String(), templatereuse.ErrInvalidInput)
	}
	return nil
}

// ============================================================================
// 类型转换函数
// ============================================================================

// candidateDataToProto 将候选数据转换为 proto 消息。
func candidateDataToProto(c *candidate.TemplateCandidateData) *pb.TemplateCandidateSummary {
	return &pb.TemplateCandidateSummary{
		TemplateCandidateId:          c.TemplateCandidateID,
		TemplateTitle:                c.TemplateTitle,
		TemplateDescription:          c.TemplateDescription,
		Modules:                      moduleRefsDataToProto(c.Modules),
		SourceProductCount:           int32(c.SourceProductCount),
		TotalReuseProductCount:       int32(c.TotalReuseProductCount),
		LatestSourceProductUpdatedAt: timestamppb.New(c.LatestSourceProductUpdatedAt),
	}
}

// moduleRefsDataToProto 将模块引用数据转换为 proto 消息。
func moduleRefsDataToProto(refs []candidate.TemplateModuleRefData) []*pb.TemplateModuleRef {
	result := make([]*pb.TemplateModuleRef, 0, len(refs))
	for _, r := range refs {
		result = append(result, &pb.TemplateModuleRef{
			ModuleId:        r.ModuleID,
			ModuleName:      r.ModuleName,
			CapabilityKey:   r.CapabilityKey,
			CapabilityLabel: r.CapabilityLabel,
		})
	}
	return result
}

// derivedHintsToProto 将派生提示数据转换为 proto 消息。
func derivedHintsToProto(hints []service.DerivedHintData) []*pb.DerivedInsightHint {
	result := make([]*pb.DerivedInsightHint, 0, len(hints))
	for _, h := range hints {
		hint := &pb.DerivedInsightHint{
			HintType:            hintTypeToProto(h.HintType),
			Title:               h.Title,
			ExplanationText:     h.ExplanationText,
			CtaKind:             ctaKindToProto(h.CTAKind),
			TemplateCandidateId: h.TemplateCandidateID,
		}
		if h.CapabilityKey != "" {
			hint.CapabilityKey = &h.CapabilityKey
		}
		if h.ModuleID != "" {
			hint.ModuleId = &h.ModuleID
		}
		result = append(result, hint)
	}
	return result
}

// resolutionStatusToProto 将字符串解析状态转换为 proto 枚举。
func resolutionStatusToProto(status string) pb.TemplateResolutionStatus {
	switch status {
	case "RESOLVED":
		return pb.TemplateResolutionStatus_TEMPLATE_RESOLUTION_STATUS_RESOLVED
	case "UNAVAILABLE":
		return pb.TemplateResolutionStatus_TEMPLATE_RESOLUTION_STATUS_UNAVAILABLE
	default:
		return pb.TemplateResolutionStatus_TEMPLATE_RESOLUTION_STATUS_UNSPECIFIED
	}
}

// templateSourceToString 将 proto TemplateSource 枚举转换为字符串。
func templateSourceToString(s pb.TemplateSource) string {
	switch s {
	case pb.TemplateSource_TEMPLATE_SOURCE_WEEKLY_REVIEW:
		return "weekly-review"
	case pb.TemplateSource_TEMPLATE_SOURCE_DASHBOARD:
		return "dashboard"
	case pb.TemplateSource_TEMPLATE_SOURCE_PRODUCT_DETAIL:
		return "product-detail"
	default:
		return ""
	}
}

// templateSourceToProto 将字符串模板来源转换为 proto 枚举。
func templateSourceToProto(s string) pb.TemplateSource {
	switch s {
	case "weekly-review":
		return pb.TemplateSource_TEMPLATE_SOURCE_WEEKLY_REVIEW
	case "dashboard":
		return pb.TemplateSource_TEMPLATE_SOURCE_DASHBOARD
	case "product-detail":
		return pb.TemplateSource_TEMPLATE_SOURCE_PRODUCT_DETAIL
	default:
		return pb.TemplateSource_TEMPLATE_SOURCE_UNSPECIFIED
	}
}

func consumerSurfaceToString(s pb.TemplateConsumerSurface) string {
	switch s {
	case pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW:
		return "weekly-review"
	case pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_CREATE:
		return "product-create"
	case pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_DETAIL:
		return "product-detail"
	default:
		return ""
	}
}

// hintTypeToProto 将字符串提示类型转换为 proto 枚举。
func hintTypeToProto(t string) pb.DerivedInsightHintType {
	switch t {
	case "REUSE_OPPORTUNITY":
		return pb.DerivedInsightHintType_DERIVED_INSIGHT_HINT_TYPE_REUSE_OPPORTUNITY
	case "CAPABILITY_GAP":
		return pb.DerivedInsightHintType_DERIVED_INSIGHT_HINT_TYPE_CAPABILITY_GAP
	default:
		return pb.DerivedInsightHintType_DERIVED_INSIGHT_HINT_TYPE_UNSPECIFIED
	}
}

// ctaKindToProto 将字符串 CTA 类型转换为 proto 枚举。
func ctaKindToProto(k string) pb.CTAKind {
	switch k {
	case "CREATE_PRODUCT_FROM_TEMPLATE":
		return pb.CTAKind_CTA_KIND_CREATE_PRODUCT_FROM_TEMPLATE
	case "GO_TO_MODULE_DETAIL":
		return pb.CTAKind_CTA_KIND_GO_TO_MODULE_DETAIL
	case "GO_TO_PRODUCT_DETAIL":
		return pb.CTAKind_CTA_KIND_GO_TO_PRODUCT_DETAIL
	default:
		return pb.CTAKind_CTA_KIND_UNSPECIFIED
	}
}
