// Package service — 读组业务编排层。
//
// 统一承接 ListRepositories / GetRepositoryDetail /
// ListRepositoryProductCandidates / ListRepositoryModuleCandidates 编排。
//
// 文件落点：backend/internal/repositorybinding/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/repositorybinding"
	"github.com/psco/backend/internal/repositorybinding/candidate"
	"github.com/psco/backend/internal/repositorybinding/repository"
)

// QueryService 承接读组业务编排。
//
// 依赖注入：
//   - repos: repositories 表读写
//   - bindings: product_repositories / module_repositories 摘要读取
//   - productCandidate: Product 候选读取（跨模块，通过 candidate 子包隔离）
//   - moduleCandidate: Module 候选读取（跨模块，通过 candidate 子包隔离）
type QueryService struct {
	repos            *repository.RepositoryStore
	bindings         *repository.BindingStore
	productCandidate *candidate.ProductCandidateRead
	moduleCandidate  *candidate.ModuleCandidateRead
}

// NewQueryService 构造 QueryService。
func NewQueryService(repos *repository.RepositoryStore, bindings *repository.BindingStore, productCandidate *candidate.ProductCandidateRead, moduleCandidate *candidate.ModuleCandidateRead) *QueryService {
	return &QueryService{repos: repos, bindings: bindings, productCandidate: productCandidate, moduleCandidate: moduleCandidate}
}

// ListRepositories 承接 RepositoryListRead。
//
// 列表读取至少承接（phase04-04）：
// name / url / provider / status / created_at / product_bind_count / module_bind_count
//
// 筛选只允许 queryText / statusFilter。
// 排序按 created_at DESC（phase04-04 冻结）。
func (s *QueryService) ListRepositories(ctx context.Context, q repositorybinding.ListRepositoriesQuery) ([]repositorybinding.RepositoryListItem, error) {
	statusFilter := ""
	if q.StatusFilter != "" && q.StatusFilter != "all" {
		statusFilter = string(q.StatusFilter)
	}

	repos, err := s.repos.List(ctx, q.QueryText, statusFilter)
	if err != nil {
		return nil, err
	}

	items := make([]repositorybinding.RepositoryListItem, 0, len(repos))
	for _, r := range repos {
		productCount, err := s.repos.CountProductBindings(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("count product bindings for repository %s: %w", r.ID, err)
		}
		moduleCount, err := s.repos.CountModuleBindings(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("count module bindings for repository %s: %w", r.ID, err)
		}

		items = append(items, repositorybinding.RepositoryListItem{
			ID:               r.ID,
			Name:             r.Name,
			URL:              r.URL,
			Provider:         r.Provider,
			Status:           r.Status,
			CreatedAt:        r.CreatedAt,
			ProductBindCount: productCount,
			ModuleBindCount:  moduleCount,
		})
	}
	return items, nil
}

// GetRepositoryDetail 承接 RepositoryDetailRead。
//
// 详情读取统一承接（phase04-04）：
//   - Repository 核心字段
//   - 已绑定 Product 列表（通过 RepositoryProductSummaryRead 注入）
//   - 已映射 Module 列表（通过 RepositoryModuleSummaryRead 注入）
//
// service 层不直接写跨模块 SQL（对齐 phase04-07）。
// 候选读取结果不内嵌于此。
func (s *QueryService) GetRepositoryDetail(ctx context.Context, repositoryID string) (*repositorybinding.RepositoryDetail, error) {
	// 校验 ID 格式
	if err := repositorybinding.ValidateRepositoryID(repositoryID); err != nil {
		return nil, err
	}

	// 读取仓库核心字段
	r, err := s.repos.GetByID(ctx, repositoryID)
	if err != nil {
		return nil, err
	}

	// 读取已绑定 Product 摘要（通过注入的 BindingStore，不直接写跨模块 SQL）
	boundProducts, err := s.bindings.ListBoundProductsByRepository(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("list bound products: %w", err)
	}

	// 读取已映射 Module 摘要（通过注入的 BindingStore）
	mappedModules, err := s.bindings.ListMappedModulesByRepository(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("list mapped modules: %w", err)
	}

	// 确保切片非 nil
	if boundProducts == nil {
		boundProducts = []repositorybinding.BoundProductSummary{}
	}
	if mappedModules == nil {
		mappedModules = []repositorybinding.MappedModuleSummary{}
	}

	return &repositorybinding.RepositoryDetail{
		Repository:    *r,
		BoundProducts: boundProducts,
		MappedModules: mappedModules,
	}, nil
}

// ListRepositoryProductCandidates 承接 ProductBindingCandidateRead。
//
// 读取可绑定到指定仓库的 Product 候选，排除已绑定的 Product。
// 排序按 products.created_at DESC（phase04-04 冻结）。
// 无可关联候选时返回空列表，不返回错误。
func (s *QueryService) ListRepositoryProductCandidates(ctx context.Context, repositoryID string) ([]repositorybinding.RepositoryProductCandidate, error) {
	// 校验 ID 格式
	if err := repositorybinding.ValidateRepositoryID(repositoryID); err != nil {
		return nil, err
	}

	// 校验仓库存在性（候选读取前提）
	if _, err := s.repos.GetByID(ctx, repositoryID); err != nil {
		return nil, err
	}

	candidates, err := s.productCandidate.List(ctx, repositoryID)
	if err != nil {
		return nil, err
	}
	if candidates == nil {
		candidates = []repositorybinding.RepositoryProductCandidate{}
	}
	return candidates, nil
}

// ListRepositoryModuleCandidates 承接 RepositoryModuleCandidateRead。
//
// 读取可映射到指定仓库的 Module 候选，排除已映射的 Module。
// 排序按 modules.created_at DESC（phase04-04 冻结）。
// 无可关联候选时返回空列表，不返回错误。
func (s *QueryService) ListRepositoryModuleCandidates(ctx context.Context, repositoryID string) ([]repositorybinding.RepositoryModuleCandidate, error) {
	// 校验 ID 格式
	if err := repositorybinding.ValidateRepositoryID(repositoryID); err != nil {
		return nil, err
	}

	// 校验仓库存在性（候选读取前提）
	if _, err := s.repos.GetByID(ctx, repositoryID); err != nil {
		return nil, err
	}

	candidates, err := s.moduleCandidate.List(ctx, repositoryID)
	if err != nil {
		return nil, err
	}
	if candidates == nil {
		candidates = []repositorybinding.RepositoryModuleCandidate{}
	}
	return candidates, nil
}
