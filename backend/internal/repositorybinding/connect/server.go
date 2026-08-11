// Package connect — Repository Binding Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Repository Binding 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/repositorybinding/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	commonv1 "github.com/psco/backend/internal/gen/proto/psco/common/v1"
	pb "github.com/psco/backend/internal/gen/proto/psco/repository_binding/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/repository_binding/v1/repository_bindingv1connect"
	module_registryv1 "github.com/psco/backend/internal/gen/proto/psco/module_registry/v1"
	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/connecterrors"
	"github.com/psco/backend/internal/repositorybinding"
	"github.com/psco/backend/internal/repositorybinding/service"
)

// Server 实现 RepositoryBindingServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.RepositoryBindingServiceHandler = (*Server)(nil)

// NewServer 构造 Repository Binding Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{querySvc: querySvc, commandSvc: commandSvc}
}

// ListRepositories 承接 RepositoryListRead。
func (s *Server) ListRepositories(ctx context.Context, req *pb.ListRepositoriesRequest) (*pb.ListRepositoriesResponse, error) {
	items, err := s.querySvc.ListRepositories(ctx, repositorybinding.ListRepositoriesQuery{
		QueryText:    req.GetQueryText(),
		StatusFilter: protoActiveArchivedStatusToDomain(req.GetStatusFilter()),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbItems := make([]*pb.RepositoryListItem, 0, len(items))
	for _, item := range items {
		pbItems = append(pbItems, &pb.RepositoryListItem{
			Id:               item.ID,
			Name:             item.Name,
			Url:              item.URL,
			Provider:         item.Provider,
			Status:           domainRepositoryStatusToProto(item.Status),
			CreatedAt:        timestamppb.New(item.CreatedAt),
			ProductBindCount: int32(item.ProductBindCount),
			ModuleBindCount:  int32(item.ModuleBindCount),
		})
	}

	return &pb.ListRepositoriesResponse{Repositories: pbItems}, nil
}

// GetRepositoryDetail 承接 RepositoryDetailRead。
func (s *Server) GetRepositoryDetail(ctx context.Context, req *pb.GetRepositoryDetailRequest) (*pb.GetRepositoryDetailResponse, error) {
	detail, err := s.querySvc.GetRepositoryDetail(ctx, req.GetRepositoryId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	boundProducts := make([]*pb.BoundProductSummary, 0, len(detail.BoundProducts))
	for _, bp := range detail.BoundProducts {
		boundProducts = append(boundProducts, &pb.BoundProductSummary{
			ProductId:     bp.ProductID,
			ProductName:   bp.ProductName,
			ProductStatus: domainRepositoryStatusToProto(bp.ProductStatus),
		})
	}

	mappedModules := make([]*pb.MappedModuleSummary, 0, len(detail.MappedModules))
	for _, mm := range detail.MappedModules {
		mappedModules = append(mappedModules, &pb.MappedModuleSummary{
			ModuleId:     mm.ModuleID,
			ModuleName:   mm.ModuleName,
			ModuleStatus: domainModuleStatusToProtoCommon(mm.ModuleStatus),
		})
	}

	return &pb.GetRepositoryDetailResponse{
		RepositoryDetail: &pb.RepositoryDetail{
			Repository: &pb.Repository{
				Id:        detail.Repository.ID,
				Name:      detail.Repository.Name,
				Url:       detail.Repository.URL,
				Provider:  detail.Repository.Provider,
				Status:    domainRepositoryStatusToProto(detail.Repository.Status),
				CreatedAt: timestamppb.New(detail.Repository.CreatedAt),
			},
			BoundProducts:  boundProducts,
			MappedModules:  mappedModules,
		},
	}, nil
}

// ListRepositoryProductCandidates 承接 ProductBindingCandidateRead。
func (s *Server) ListRepositoryProductCandidates(ctx context.Context, req *pb.ListRepositoryProductCandidatesRequest) (*pb.ListRepositoryProductCandidatesResponse, error) {
	candidates, err := s.querySvc.ListRepositoryProductCandidates(ctx, req.GetRepositoryId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbCandidates := make([]*pb.RepositoryProductCandidate, 0, len(candidates))
	for _, c := range candidates {
		pbCandidates = append(pbCandidates, &pb.RepositoryProductCandidate{
			ProductId:     c.ProductID,
			ProductName:   c.ProductName,
			ProductStatus: domainRepositoryStatusToProto(c.ProductStatus),
		})
	}

	return &pb.ListRepositoryProductCandidatesResponse{Candidates: pbCandidates}, nil
}

// ListRepositoryModuleCandidates 承接 RepositoryModuleCandidateRead。
func (s *Server) ListRepositoryModuleCandidates(ctx context.Context, req *pb.ListRepositoryModuleCandidatesRequest) (*pb.ListRepositoryModuleCandidatesResponse, error) {
	candidates, err := s.querySvc.ListRepositoryModuleCandidates(ctx, req.GetRepositoryId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbCandidates := make([]*pb.RepositoryModuleCandidate, 0, len(candidates))
	for _, c := range candidates {
		pbCandidates = append(pbCandidates, &pb.RepositoryModuleCandidate{
			ModuleId:     c.ModuleID,
			ModuleName:   c.ModuleName,
			ModuleStatus: domainModuleStatusToProtoCommon(c.ModuleStatus),
		})
	}

	return &pb.ListRepositoryModuleCandidatesResponse{Candidates: pbCandidates}, nil
}

// CreateRepository 承接 RepositoryCreateWrite。
func (s *Server) CreateRepository(ctx context.Context, req *pb.CreateRepositoryRequest) (*pb.CreateRepositoryResponse, error) {
	result, err := s.commandSvc.CreateRepository(ctx, repositorybinding.CreateRepositoryRequest{
		Name:     req.GetName(),
		URL:      req.GetUrl(),
		Provider: req.GetProvider(),
		Status:   protoActiveArchivedStatusToDomain(req.GetStatus()),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateRepositoryResponse{RepositoryId: result.RepositoryID}, nil
}

// BindRepositoryToProduct 承接 RepositoryProductBindingWrite。
func (s *Server) BindRepositoryToProduct(ctx context.Context, req *pb.BindRepositoryToProductRequest) (*pb.BindRepositoryToProductResponse, error) {
	err := s.commandSvc.BindRepositoryToProduct(ctx, req.GetRepositoryId(), req.GetProductId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.BindRepositoryToProductResponse{}, nil
}

// MapModuleToRepository 承接 RepositoryModuleMappingWrite。
func (s *Server) MapModuleToRepository(ctx context.Context, req *pb.MapModuleToRepositoryRequest) (*pb.MapModuleToRepositoryResponse, error) {
	err := s.commandSvc.MapModuleToRepository(ctx, req.GetRepositoryId(), req.GetModuleId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.MapModuleToRepositoryResponse{}, nil
}

// --- 类型转换函数 ---

// protoActiveArchivedStatusToDomain 将 proto ActiveArchivedStatus 转为 domain RepositoryStatus。
func protoActiveArchivedStatusToDomain(s commonv1.ActiveArchivedStatus) repositorybinding.RepositoryStatus {
	switch s {
	case commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ACTIVE:
		return repositorybinding.RepositoryStatusActive
	case commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ARCHIVED:
		return repositorybinding.RepositoryStatusArchived
	default:
		return ""
	}
}

// domainRepositoryStatusToProto 将 domain RepositoryStatus 转为 proto ActiveArchivedStatus。
func domainRepositoryStatusToProto(s repositorybinding.RepositoryStatus) commonv1.ActiveArchivedStatus {
	switch s {
	case repositorybinding.RepositoryStatusActive:
		return commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ACTIVE
	case repositorybinding.RepositoryStatusArchived:
		return commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ARCHIVED
	default:
		return commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_UNSPECIFIED
	}
}

// domainModuleStatusToProtoCommon 将 moduleregistry.ModuleStatus 转为 proto ModuleStatus。
func domainModuleStatusToProtoCommon(s moduleregistry.ModuleStatus) module_registryv1.ModuleStatus {
	switch s {
	case moduleregistry.ModuleStatusActive:
		return module_registryv1.ModuleStatus_MODULE_STATUS_ACTIVE
	case moduleregistry.ModuleStatusArchived:
		return module_registryv1.ModuleStatus_MODULE_STATUS_ARCHIVED
	default:
		return module_registryv1.ModuleStatus_MODULE_STATUS_UNSPECIFIED
	}
}