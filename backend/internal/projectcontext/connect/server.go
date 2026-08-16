// Package connect — projectcontext Connect transport 实现。
//
// 本文件是 phase11-07 最小只读项目上下文能力的 Connect handler 实现。
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
	"github.com/psco/backend/internal/projectcontext"
	"github.com/psco/backend/internal/projectcontext/service"
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
	}, nil
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
