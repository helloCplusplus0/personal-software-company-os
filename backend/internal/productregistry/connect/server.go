// Package connect — Product Registry Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Product Registry 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/productregistry/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	commonv1 "github.com/psco/backend/internal/gen/proto/psco/common/v1"
	pb "github.com/psco/backend/internal/gen/proto/psco/product_registry/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/product_registry/v1/product_registryv1connect"
	module_registryv1 "github.com/psco/backend/internal/gen/proto/psco/module_registry/v1"
	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/connecterrors"
	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/productregistry/service"
)

// Server 实现 ProductRegistryServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.ProductRegistryServiceHandler = (*Server)(nil)

// NewServer 构造 Product Registry Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{querySvc: querySvc, commandSvc: commandSvc}
}

// ListProducts 承接 ProductListRead。
func (s *Server) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	items, err := s.querySvc.ListProducts(ctx, productregistry.ListProductsQuery{
		QueryText:    req.GetQueryText(),
		StatusFilter: protoProductStatusToDomain(req.GetStatusFilter()),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbItems := make([]*pb.ProductListItem, 0, len(items))
	for _, item := range items {
		pbItems = append(pbItems, &pb.ProductListItem{
			Id:                  item.ID,
			Name:                item.Name,
			Description:         item.Description,
			Status:              domainProductStatusToProto(item.Status),
			CreatedAt:           timestamppb.New(item.CreatedAt),
			ModuleBindCount:     int32(item.ModuleBindCount),
			RepositoryBindCount: int32(item.RepositoryBindCount),
		})
	}

	return &pb.ListProductsResponse{Products: pbItems}, nil
}

// GetProductDetail 承接 ProductDetailRead。
func (s *Server) GetProductDetail(ctx context.Context, req *pb.GetProductDetailRequest) (*pb.GetProductDetailResponse, error) {
	detail, err := s.querySvc.GetProductDetail(ctx, req.GetProductId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	boundModules := make([]*pb.BoundModuleSummary, 0, len(detail.BoundModules))
	for _, bm := range detail.BoundModules {
		boundModules = append(boundModules, &pb.BoundModuleSummary{
			ModuleId:     bm.ModuleID,
			ModuleName:   bm.ModuleName,
			ModuleStatus: domainModuleStatusToProtoCommon(bm.ModuleStatus),
		})
	}

	boundRepos := make([]*pb.BoundRepositorySummary, 0, len(detail.BoundRepositories))
	for _, br := range detail.BoundRepositories {
		boundRepos = append(boundRepos, &pb.BoundRepositorySummary{
			RepositoryId:     br.RepositoryID,
			RepositoryName:   br.RepositoryName,
			Provider:         br.Provider,
			RepositoryStatus: domainProductStatusToProto(br.RepositoryStatus),
		})
	}

	return &pb.GetProductDetailResponse{
		ProductDetail: &pb.ProductDetail{
			Product: &pb.Product{
				Id:          detail.Product.ID,
				Name:        detail.Product.Name,
				Description: detail.Product.Description,
				Status:      domainProductStatusToProto(detail.Product.Status),
				CreatedAt:   timestamppb.New(detail.Product.CreatedAt),
			},
			BoundModules:      boundModules,
			BoundRepositories: boundRepos,
		},
	}, nil
}

// ListProductModuleCandidates 承接 ProductModuleCandidateRead。
func (s *Server) ListProductModuleCandidates(ctx context.Context, req *pb.ListProductModuleCandidatesRequest) (*pb.ListProductModuleCandidatesResponse, error) {
	candidates, err := s.querySvc.ListProductModuleCandidates(ctx, req.GetProductId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbCandidates := make([]*pb.ProductModuleCandidate, 0, len(candidates))
	for _, c := range candidates {
		pbCandidates = append(pbCandidates, &pb.ProductModuleCandidate{
			ModuleId:     c.ModuleID,
			ModuleName:   c.ModuleName,
			ModuleStatus: domainModuleStatusToProtoCommon(c.ModuleStatus),
		})
	}

	return &pb.ListProductModuleCandidatesResponse{Candidates: pbCandidates}, nil
}

// CreateProduct 承接 ProductCreateWrite。
func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	result, err := s.commandSvc.CreateProduct(ctx, productregistry.CreateProductRequest{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Status:      protoProductStatusToDomain(req.GetStatus()),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateProductResponse{ProductId: result.ProductID}, nil
}

// BindModuleToProduct 承接 ProductModuleBindingWrite。
func (s *Server) BindModuleToProduct(ctx context.Context, req *pb.BindModuleToProductRequest) (*pb.BindModuleToProductResponse, error) {
	err := s.commandSvc.BindModuleToProduct(ctx, req.GetProductId(), req.GetModuleId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.BindModuleToProductResponse{}, nil
}

// --- 类型转换函数 ---

func protoProductStatusToDomain(s commonv1.ActiveArchivedStatus) productregistry.ProductStatus {
	switch s {
	case commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ACTIVE:
		return productregistry.ProductStatusActive
	case commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ARCHIVED:
		return productregistry.ProductStatusArchived
	default:
		return ""
	}
}

func domainProductStatusToProto(s productregistry.ProductStatus) commonv1.ActiveArchivedStatus {
	switch s {
	case productregistry.ProductStatusActive:
		return commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ACTIVE
	case productregistry.ProductStatusArchived:
		return commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_ARCHIVED
	default:
		return commonv1.ActiveArchivedStatus_ACTIVE_ARCHIVED_STATUS_UNSPECIFIED
	}
}

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