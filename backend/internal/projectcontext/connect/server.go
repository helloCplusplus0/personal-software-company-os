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
	progresspb "github.com/psco/backend/internal/gen/proto/psco/progress/v1"
	standardpb "github.com/psco/backend/internal/gen/proto/psco/standard/v1"
	"github.com/psco/backend/internal/projectcontext"
	"github.com/psco/backend/internal/projectcontext/service"
	progressconnect "github.com/psco/backend/internal/progress/connect"
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
// 组装约束（2026-08-18 phase14-10 T7 裁决后的 5 顶层块字段面（槽位 2/3/4 reserved）
// + phase15 progress 摘要块（字段 9）；旧 global_assets 与画像残余
// governance_profile / current_phase 均已移除）：
//   - products[] / modules[] / decisions[] 数组语义，空数组合法
//   - standards[] 复用 standard/connect 导出的 DomainStandardToProto
//     （含递归树转换），与 StandardService 读取同源，不重写第二套树映射；
//     template_source 语义经 role=template_source 绑定随 standards[] 消费
//   - progress（phase15-06 新增）复用 progress/connect 导出的
//     DomainProgressEventToProto 组装 latest_task_completed / recent_events
//     元素，与 ProgressService 读取同源，不重写第二套事件映射；
//     BriefProgress 恒构造（空态零值字段 + 空数组非 nil）
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

	// progress 摘要块（phase15-06）：恒构造非 nil；RecentEvents 空态组装为
	// 空数组（非 nil 切片）；LatestTaskCompleted 为 nil 时不设置（零值不设置语义）
	recentEvents := make([]*progresspb.ProgressEvent, 0, len(result.Progress.RecentEvents))
	for _, item := range result.Progress.RecentEvents {
		recentEvents = append(recentEvents, progressconnect.DomainProgressEventToProto(item))
	}
	briefProgress := &pb.BriefProgress{
		CurrentPhaseKey:   result.Progress.CurrentPhaseKey,
		CurrentPhaseLabel: result.Progress.CurrentPhaseLabel,
		RecentEvents:      recentEvents,
	}
	if result.Progress.LatestTaskCompleted != nil {
		briefProgress.LatestTaskCompleted = progressconnect.DomainProgressEventToProto(*result.Progress.LatestTaskCompleted)
	}

	return &pb.GetProjectBriefResponse{
		Repository: domainRepositorySummaryToProto(result.Repository),
		Products:   domainProductSummariesToProto(result.Products),
		Modules:    domainModuleSummariesToProto(result.Modules),
		Decisions:  domainDecisionSummariesToProto(result.Decisions),
		Standards:  standards,
		Progress:   briefProgress,
	}, nil
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
