// Package connect — projectcontext Connect transport 实现。
//
// 本文件是 phase11-07 最小只读项目上下文能力与 phase13-10 agent 项目简报
// 读取主线的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/projectcontext/connect/server.go
package connect

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/psco/backend/internal/connecterrors"
	pbc "github.com/psco/backend/internal/gen/connect/psco/project_context/v1/project_contextv1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/project_context/v1"
	standardpb "github.com/psco/backend/internal/gen/proto/psco/standard/v1"
	"github.com/psco/backend/internal/governanceprofile"
	"github.com/psco/backend/internal/projectcontext"
	"github.com/psco/backend/internal/projectcontext/service"
	standardconnect "github.com/psco/backend/internal/standard/connect"
)

// Server 实现 ProjectContextServiceHandler 接口。
type Server struct {
	querySvc *service.QueryService
}

var _ pbc.ProjectContextServiceHandler = (*Server)(nil)

// NewServer 构造 projectcontext Connect handler。
func NewServer(querySvc *service.QueryService) *Server {
	return &Server{querySvc: querySvc}
}

// GetProjectContext 承接最小只读项目上下文聚合读取。
func (s *Server) GetProjectContext(ctx context.Context, req *pb.GetProjectContextRequest) (*pb.GetProjectContextResponse, error) {
	result, err := s.querySvc.GetProjectContext(ctx, req.GetRepositoryId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.GetProjectContextResponse{
		Repository: domainRepositorySummaryToProto(result.Repository),
		Product:    domainProductSummaryToProto(result.Product),
		Modules:    domainModuleSummariesToProto(result.Modules),
		Decisions:  domainDecisionSummariesToProto(result.Decisions),
		Rules:      domainRuleEntriesToProto(result.Rules),
		Phases:     domainPhaseEntriesToProto(result.Phases),
                Boundaries: domainBoundaryEntriesToProto(result.Boundaries),
	}, nil
}

// ExportProjectContext 承接 AGENTS 风格 Markdown 项目上下文导出。
//
// Deprecated: 兼容导出层（phase13-10 冻结兼容窗口）。单向派生自
// GetProjectContext 结构化结果，不绕过结构化读取主线。
func (s *Server) ExportProjectContext(ctx context.Context, req *pb.ExportProjectContextRequest) (*pb.ExportProjectContextResponse, error) {
	markdown, err := s.querySvc.ExportProjectContext(ctx, req.GetRepositoryId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.ExportProjectContextResponse{
		Markdown: markdown,
	}, nil
}

// GetProjectBrief 承接 agent 项目简报结构化读取（phase13-10 正式主线）。
//
// 组装约束（phase14-09 切换后的 8 顶层块字段面；旧 global_assets 已移除）：
//   - governance_profile 组装为 project_context.proto 内联 BriefGovernanceProfile
//     （repository_id / track_type / template_source 三字段，domain→gen 枚举映射）
//   - current_phase 从治理画像主记录 read-only 字段单向派生，
//     status 映射内联 BriefPhaseStatus
//   - products[] / modules[] / decisions[] 数组语义，空数组合法
//   - standards[] 复用 standard/connect 导出的 DomainStandardToProto
//     （含递归树转换），与 StandardService 读取同源，不重写第二套树映射
//   - 不混入硬编码投影、目录扫描或自然语言指导词
func (s *Server) GetProjectBrief(ctx context.Context, req *pb.GetProjectBriefRequest) (*pb.GetProjectBriefResponse, error) {
	result, err := s.querySvc.GetProjectBrief(ctx, req.GetRepositoryId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	standards := make([]*standardpb.Standard, 0, len(result.Standards))
	for _, item := range result.Standards {
		standards = append(standards, standardconnect.DomainStandardToProto(item))
	}

	return &pb.GetProjectBriefResponse{
		Repository: domainRepositorySummaryToProto(result.Repository),
		GovernanceProfile: &pb.BriefGovernanceProfile{
			RepositoryId:   result.Repository.ID,
			TrackType:      domainTrackTypeToBriefProto(result.GovernanceProfile.TrackType),
			TemplateSource: result.GovernanceProfile.TemplateSource,
		},
		CurrentPhase: &pb.BriefCurrentPhase{
			Name:     result.CurrentPhase.Name,
			EntryRef: result.CurrentPhase.EntryRef,
			Status:   domainPhaseStatusToBriefProto(result.GovernanceProfile.CurrentPhaseStatus),
		},
		Products:  domainProductSummariesToProto(result.Products),
		Modules:   domainModuleSummariesToProto(result.Modules),
		Decisions: domainDecisionSummariesToProto(result.Decisions),
		Standards: standards,
	}, nil
}

// --- 治理画像 domain → gen 枚举映射（phase14-09 内联 brief 切换） ---

// domainTrackTypeToBriefProto 将治理画像技术路线受控枚举转换为
// project_context.proto 内联 BriefTrackType。
func domainTrackTypeToBriefProto(t governanceprofile.TrackType) pb.BriefTrackType {
	switch t {
	case governanceprofile.TrackTypeProduct:
		return pb.BriefTrackType_BRIEF_TRACK_TYPE_PRODUCT
	case governanceprofile.TrackTypeDurableSystem:
		return pb.BriefTrackType_BRIEF_TRACK_TYPE_DURABLE_SYSTEM
	default:
		return pb.BriefTrackType_BRIEF_TRACK_TYPE_UNSPECIFIED
	}
}

// domainPhaseStatusToBriefProto 将治理画像阶段状态受控枚举转换为
// project_context.proto 内联 BriefPhaseStatus（值域与画像 PhaseStatus 一致）。
func domainPhaseStatusToBriefProto(s governanceprofile.PhaseStatus) pb.BriefPhaseStatus {
	switch s {
	case governanceprofile.PhaseStatusPlanned:
		return pb.BriefPhaseStatus_BRIEF_PHASE_STATUS_PLANNED
	case governanceprofile.PhaseStatusInProgress:
		return pb.BriefPhaseStatus_BRIEF_PHASE_STATUS_IN_PROGRESS
	case governanceprofile.PhaseStatusCompleted:
		return pb.BriefPhaseStatus_BRIEF_PHASE_STATUS_COMPLETED
	case governanceprofile.PhaseStatusBlocked:
		return pb.BriefPhaseStatus_BRIEF_PHASE_STATUS_BLOCKED
	default:
		return pb.BriefPhaseStatus_BRIEF_PHASE_STATUS_UNSPECIFIED
	}
}

// domainProductSummariesToProto 将产品摘要数组转换为 proto（brief 数组语义专用）。
func domainProductSummariesToProto(items []projectcontext.ProductSummary) []*pb.ProductSummary {
	result := make([]*pb.ProductSummary, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.ProductSummary{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Status:      item.Status,
		})
	}
	return result
}

// --- 类型转换函数 ---

func domainRepositorySummaryToProto(s *projectcontext.RepositorySummary) *pb.RepositorySummary {
	if s == nil {
		return nil
	}
	return &pb.RepositorySummary{
		Id:          s.ID,
		Name:        s.Name,
		Provider:    s.Provider,
		Url:         s.URL,
		Description: s.Description,
	}
}

func domainProductSummaryToProto(s *projectcontext.ProductSummary) *pb.ProductSummary {
	if s == nil {
		return nil
	}
	return &pb.ProductSummary{
		Id:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Status:      s.Status,
	}
}

func domainModuleSummariesToProto(items []projectcontext.ModuleSummary) []*pb.ModuleSummary {
	result := make([]*pb.ModuleSummary, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.ModuleSummary{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Status:      item.Status,
		})
	}
	return result
}

func domainDecisionSummariesToProto(items []projectcontext.DecisionSummary) []*pb.DecisionSummary {
	result := make([]*pb.DecisionSummary, 0, len(items))
	for _, item := range items {
		createdAt, _ := parseTimestamp(item.CreatedAt)
		result = append(result, &pb.DecisionSummary{
			Id:         item.ID,
			Title:      item.Title,
			Status:     item.Status,
			Context:    item.Context,
			HitSources: item.HitSources,
			CreatedAt:  createdAt,
		})
	}
	return result
}

func domainRuleEntriesToProto(items []projectcontext.RuleEntry) []*pb.RuleEntry {
	result := make([]*pb.RuleEntry, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.RuleEntry{
			Key:       item.Key,
			Label:     item.Label,
			Summary:   item.Summary,
			EntryRef:  item.EntryRef,
			EntryKind: item.EntryKind,
		})
	}
	return result
}

func domainPhaseEntriesToProto(items []projectcontext.PhaseEntry) []*pb.PhaseEntry {
	result := make([]*pb.PhaseEntry, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.PhaseEntry{
			Phase:         item.Phase,
			Label:         item.Label,
			StatusSummary: item.StatusSummary,
			EntryRef:      item.EntryRef,
			EntryKind:     item.EntryKind,
		})
	}
	return result
}

func domainBoundaryEntriesToProto(items []projectcontext.BoundaryEntry) []*pb.BoundaryEntry {
        result := make([]*pb.BoundaryEntry, 0, len(items))
        for _, item := range items {
                result = append(result, &pb.BoundaryEntry{
                        Key:     item.Key,
                        Label:   item.Label,
                        Summary: item.Summary,
                })
        }
        return result
}

// parseTimestamp 将数据库返回的 timestamp 文本解析为 protobuf Timestamp。
func parseTimestamp(s string) (*timestamppb.Timestamp, error) {
	if s == "" {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		// 尝试 PostgreSQL 默认格式
		t, err = time.Parse("2006-01-02 15:04:05-07:00", s)
		if err != nil {
			return nil, err
		}
	}
	return timestamppb.New(t), nil
}
