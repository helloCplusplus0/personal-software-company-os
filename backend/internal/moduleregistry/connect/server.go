// Package connect — Module Registry Connect transport 实现。
//
// 本文件是 phase07-09 正式传输主线切换后，Module Registry 模块的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
//
// 文件落点：backend/internal/moduleregistry/connect/server.go
package connect

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/psco/backend/internal/gen/proto/psco/module_registry/v1"
	pbc "github.com/psco/backend/internal/gen/connect/psco/module_registry/v1/module_registryv1connect"
	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/service"
	"github.com/psco/backend/internal/connecterrors"
)

// ProductCandidateReader 承接旧 ProductBindingCandidateRead 的兼容委派。
// phase04-12 起，canonical 候选读取已迁移到 Product Registry。
// 该函数类型在主调方（router 层）通过闭包委派到 ProductRegistry.QueryService.ListProducts。
type ProductCandidateReader func(ctx context.Context) ([]moduleregistry.ProductCandidate, error)

// RepositoryCandidateReader 承接旧 RepositoryBindingCandidateRead 的兼容委派。
// phase04-12 起，canonical 候选读取已迁移到 Repository Binding。
// 该函数类型在主调方（router 层）通过闭包委派到 RepositoryBinding.QueryService.ListRepositories。
type RepositoryCandidateReader func(ctx context.Context) ([]moduleregistry.RepositoryCandidate, error)

// BindModuleToProductDelegate 承接旧 ModuleBindingWrite (product) 的兼容委派。
// phase04-12 起，canonical 绑定写入已迁移到 Product Registry。
// 该函数类型在主调方（router 层）通过闭包委派到 ProductRegistry.CommandService.BindModuleToProduct。
type BindModuleToProductDelegate func(ctx context.Context, moduleID, productID string) error

// MapModuleToRepositoryDelegate 承接旧 ModuleBindingWrite (repository) 的兼容委派。
// phase04-12 起，canonical 映射写入已迁移到 Repository Binding。
// 该函数类型在主调方（router 层）通过闭包委派到 RepositoryBinding.CommandService.MapModuleToRepository。
type MapModuleToRepositoryDelegate func(ctx context.Context, moduleID, repositoryID string) error

// Server 实现 ModuleRegistryServiceHandler 接口。
//
// 4 个已迁移 RPC（BindModuleToProduct / MapModuleToRepository /
// ListProductCandidates / ListRepositoryCandidates）通过可选 delegate 承接。
// 若 delegate 为 nil，对应 RPC 返回 CodeInternal 错误。
type Server struct {
	querySvc               *service.QueryService
	commandSvc             *service.CommandService
	productCandidates      ProductCandidateReader
	repositoryCandidates   RepositoryCandidateReader
	bindModuleToProduct    BindModuleToProductDelegate
	mapModuleToRepository  MapModuleToRepositoryDelegate
}

var _ pbc.ModuleRegistryServiceHandler = (*Server)(nil)

// NewServer 构造 Module Registry Connect handler。
//
// 四个 delegate 参数可选（nil 表示对应 RPC 不可用，返回内部错误）。
// phase07-09 阶段，router 层应为 L3/L4 兼容入口提供有效 delegate；
// L1/L2 候选读取 compat 在本阶段退场后不再提供 delegate。
func NewServer(
	querySvc *service.QueryService,
	commandSvc *service.CommandService,
	productCandidates ProductCandidateReader,
	repositoryCandidates RepositoryCandidateReader,
	bindModuleToProduct BindModuleToProductDelegate,
	mapModuleToRepository MapModuleToRepositoryDelegate,
) *Server {
	return &Server{
		querySvc:              querySvc,
		commandSvc:            commandSvc,
		productCandidates:     productCandidates,
		repositoryCandidates:  repositoryCandidates,
		bindModuleToProduct:   bindModuleToProduct,
		mapModuleToRepository: mapModuleToRepository,
	}
}

// ListModules 承接 ModuleListRead。
func (s *Server) ListModules(ctx context.Context, req *pb.ListModulesRequest) (*pb.ListModulesResponse, error) {
	items, err := s.querySvc.ListModules(ctx, moduleregistry.ListQuery{
		QueryText:    req.GetQueryText(),
		StatusFilter: protoModuleStatusToDomain(req.GetStatusFilter()),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbItems := make([]*pb.ModuleListItem, 0, len(items))
	for _, item := range items {
		pbItem := &pb.ModuleListItem{
			Id:                  item.ID,
			Name:                item.Name,
			Description:         item.Description,
			Status:              domainModuleStatusToProto(item.Status),
			ProductBindCount:    int32(item.ProductBindCount),
			RepositoryBindCount: int32(item.RepositoryBindCount),
		}
		if item.LatestRelease != nil {
			pbItem.LatestRelease = item.LatestRelease
		}
		pbItems = append(pbItems, pbItem)
	}

	return &pb.ListModulesResponse{Modules: pbItems}, nil
}

// GetModuleDetail 承接 ModuleDetailRead（含 Decision 内嵌附属读取）。
func (s *Server) GetModuleDetail(ctx context.Context, req *pb.GetModuleDetailRequest) (*pb.GetModuleDetailResponse, error) {
	detail, err := s.querySvc.GetModuleDetail(ctx, req.GetModuleId())
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	releases := make([]*pb.Release, 0, len(detail.Releases))
	for _, r := range detail.Releases {
		releases = append(releases, &pb.Release{
			Id:         r.ID,
			ModuleId:   r.ModuleID,
			Version:    r.Version,
			Status:     domainReleaseStatusToProto(r.Status),
			ReleasedAt: timestamppb.New(r.ReleasedAt),
		})
	}

	productBindings := make([]*pb.ProductBinding, 0, len(detail.ProductBindings))
	for _, pbItem := range detail.ProductBindings {
		productBindings = append(productBindings, &pb.ProductBinding{
			ProductId:   pbItem.ProductID,
			ProductName: pbItem.ProductName,
		})
	}

	repositoryMappings := make([]*pb.RepositoryMapping, 0, len(detail.RepositoryMappings))
	for _, rm := range detail.RepositoryMappings {
		repositoryMappings = append(repositoryMappings, &pb.RepositoryMapping{
			RepositoryId:   rm.RepositoryID,
			RepositoryName: rm.RepositoryName,
		})
	}

	decisionLinks := make([]*pb.DecisionLink, 0, len(detail.DecisionLinks))
	for _, dl := range detail.DecisionLinks {
		decisionLinks = append(decisionLinks, &pb.DecisionLink{
			DecisionId:    dl.DecisionID,
			DecisionTitle: dl.DecisionTitle,
		})
	}

	return &pb.GetModuleDetailResponse{
		ModuleDetail: &pb.ModuleDetail{
			Module: &pb.Module{
				Id:          detail.Module.ID,
				Name:        detail.Module.Name,
				Description: detail.Module.Description,
				Status:      domainModuleStatusToProto(detail.Module.Status),
				CreatedAt:   timestamppb.New(detail.Module.CreatedAt),
			},
			Releases:            releases,
			ProductBindings:     productBindings,
			RepositoryMappings:  repositoryMappings,
			DecisionLinks:       decisionLinks,
		},
	}, nil
}

// CreateModule 承接 ModuleCreateWrite。
func (s *Server) CreateModule(ctx context.Context, req *pb.CreateModuleRequest) (*pb.CreateModuleResponse, error) {
	m, err := s.commandSvc.CreateModule(ctx, moduleregistry.CreateModuleRequest{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Status:      protoModuleStatusToDomain(req.GetStatus()),
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateModuleResponse{
		Module: &pb.Module{
			Id:          m.ID,
			Name:        m.Name,
			Description: m.Description,
			Status:      domainModuleStatusToProto(m.Status),
			CreatedAt:   timestamppb.New(m.CreatedAt),
		},
	}, nil
}

// CreateRelease 承接 ModuleReleaseWrite。
func (s *Server) CreateRelease(ctx context.Context, req *pb.CreateReleaseRequest) (*pb.CreateReleaseResponse, error) {
	// released_at 转换为 RFC3339 字符串供 service 层使用
	releasedAt := req.GetReleasedAt().AsTime().Format("2006-01-02T15:04:05Z07:00")

	r, err := s.commandSvc.CreateRelease(ctx, req.GetModuleId(), moduleregistry.CreateReleaseRequest{
		Version:    req.GetVersion(),
		Status:     protoReleaseStatusToDomain(req.GetStatus()),
		ReleasedAt: releasedAt,
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateReleaseResponse{
		Release: &pb.Release{
			Id:         r.ID,
			ModuleId:   r.ModuleID,
			Version:    r.Version,
			Status:     domainReleaseStatusToProto(r.Status),
			ReleasedAt: timestamppb.New(r.ReleasedAt),
		},
	}, nil
}

// BindModuleToProduct 承接 ModuleBindingWrite (product)。
//
// phase04-12 起，canonical owner 已迁移到 Product Registry。
// 本 handler 通过 BindModuleToProductDelegate 委派到 ProductRegistry.CommandService.BindModuleToProduct。
// 若 delegate 为 nil（router 未正确装配），返回 CodeInternal 错误。
func (s *Server) BindModuleToProduct(ctx context.Context, req *pb.BindModuleToProductRequest) (*pb.BindModuleToProductResponse, error) {
	if s.bindModuleToProduct == nil {
		return nil, connecterrors.MapToConnectError(moduleregistry.ErrInvalidInput)
	}
	if err := s.bindModuleToProduct(ctx, req.GetModuleId(), req.GetProductId()); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	return &pb.BindModuleToProductResponse{}, nil
}

// MapModuleToRepository 承接 ModuleBindingWrite (repository)。
//
// phase04-12 起，canonical owner 已迁移到 Repository Binding。
// 本 handler 通过 MapModuleToRepositoryDelegate 委派到 RepositoryBinding.CommandService.MapModuleToRepository。
// 若 delegate 为 nil（router 未正确装配），返回 CodeInternal 错误。
func (s *Server) MapModuleToRepository(ctx context.Context, req *pb.MapModuleToRepositoryRequest) (*pb.MapModuleToRepositoryResponse, error) {
	if s.mapModuleToRepository == nil {
		return nil, connecterrors.MapToConnectError(moduleregistry.ErrInvalidInput)
	}
	if err := s.mapModuleToRepository(ctx, req.GetModuleId(), req.GetRepositoryId()); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	return &pb.MapModuleToRepositoryResponse{}, nil
}

// ListProductCandidates 承接 ProductBindingCandidateRead。
//
// phase04-12 起，canonical 候选读取已迁移到 Product Registry。
// 本 handler 通过 ProductCandidateReader 委派到 ProductRegistry.QueryService.ListProducts。
// 若 delegate 为 nil（L1 compat 已退场），返回 CodeInternal 错误。
func (s *Server) ListProductCandidates(ctx context.Context, req *pb.ListProductCandidatesRequest) (*pb.ListProductCandidatesResponse, error) {
	if s.productCandidates == nil {
		return nil, connecterrors.MapToConnectError(moduleregistry.ErrInvalidInput)
	}
	items, err := s.productCandidates(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbCandidates := make([]*pb.ProductCandidate, 0, len(items))
	for _, item := range items {
		pbCandidates = append(pbCandidates, &pb.ProductCandidate{
			Id:   item.ID,
			Name: item.Name,
		})
	}
	return &pb.ListProductCandidatesResponse{Products: pbCandidates}, nil
}

// ListRepositoryCandidates 承接 RepositoryBindingCandidateRead。
//
// phase04-12 起，canonical 候选读取已迁移到 Repository Binding。
// 本 handler 通过 RepositoryCandidateReader 委派到 RepositoryBinding.QueryService.ListRepositories。
// 若 delegate 为 nil（L2 compat 已退场），返回 CodeInternal 错误。
func (s *Server) ListRepositoryCandidates(ctx context.Context, req *pb.ListRepositoryCandidatesRequest) (*pb.ListRepositoryCandidatesResponse, error) {
	if s.repositoryCandidates == nil {
		return nil, connecterrors.MapToConnectError(moduleregistry.ErrInvalidInput)
	}
	items, err := s.repositoryCandidates(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	pbCandidates := make([]*pb.RepositoryCandidate, 0, len(items))
	for _, item := range items {
		pbCandidates = append(pbCandidates, &pb.RepositoryCandidate{
			Id:   item.ID,
			Name: item.Name,
		})
	}
	return &pb.ListRepositoryCandidatesResponse{Repositories: pbCandidates}, nil
}

// --- 类型转换函数 ---

func protoModuleStatusToDomain(s pb.ModuleStatus) moduleregistry.ModuleStatus {
	switch s {
	case pb.ModuleStatus_MODULE_STATUS_ACTIVE:
		return moduleregistry.ModuleStatusActive
	case pb.ModuleStatus_MODULE_STATUS_ARCHIVED:
		return moduleregistry.ModuleStatusArchived
	default:
		return ""
	}
}

func domainModuleStatusToProto(s moduleregistry.ModuleStatus) pb.ModuleStatus {
	switch s {
	case moduleregistry.ModuleStatusActive:
		return pb.ModuleStatus_MODULE_STATUS_ACTIVE
	case moduleregistry.ModuleStatusArchived:
		return pb.ModuleStatus_MODULE_STATUS_ARCHIVED
	default:
		return pb.ModuleStatus_MODULE_STATUS_UNSPECIFIED
	}
}

func protoReleaseStatusToDomain(s pb.ReleaseStatus) moduleregistry.ReleaseStatus {
	switch s {
	case pb.ReleaseStatus_RELEASE_STATUS_ACTIVE:
		return moduleregistry.ReleaseStatusActive
	case pb.ReleaseStatus_RELEASE_STATUS_ARCHIVED:
		return moduleregistry.ReleaseStatusArchived
	default:
		return ""
	}
}

func domainReleaseStatusToProto(s moduleregistry.ReleaseStatus) pb.ReleaseStatus {
	switch s {
	case moduleregistry.ReleaseStatusActive:
		return pb.ReleaseStatus_RELEASE_STATUS_ACTIVE
	case moduleregistry.ReleaseStatusArchived:
		return pb.ReleaseStatus_RELEASE_STATUS_ARCHIVED
	default:
		return pb.ReleaseStatus_RELEASE_STATUS_UNSPECIFIED
	}
}