// Package service — 读组业务编排层。
//
// 统一承接 ListProducts / GetProductDetail / ListProductModuleCandidates 编排。
//
// 文件落点：backend/internal/productregistry/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/productregistry/candidate"
	"github.com/psco/backend/internal/productregistry/repository"
)

// QueryService 承接读组业务编排。
//
// 依赖注入：
//   - products: products 表读写
//   - bindings: product_modules 摘要读取（owner=Product Registry）
//   - boundRepoReader: product_repositories 摘要读取（owner=Repository Binding，phase04-07 L162-181 冻结）
//   - moduleCandidate: Module 候选读取（跨模块，通过 candidate 子包隔离）
//
// boundRepoReader 通过 productregistry.BoundRepositoryReader 接口注入，
// 实现由 repositorybinding/repository.BindingStore 提供，在 platform 装配点接线。
// service 层不直接读 product_repositories 表（phase04-07 L181）。
type QueryService struct {
	products         *repository.ProductStore
	bindings         *repository.BindingStore
	boundRepoReader  productregistry.BoundRepositoryReader
	moduleCandidate  *candidate.ModuleCandidateRead
}

// NewQueryService 构造 QueryService。
//
// boundRepoReader 参数由 platform 装配点注入（实现为 *repositorybinding/repository.BindingStore，
// 隐式满足 productregistry.BoundRepositoryReader 接口）。
func NewQueryService(products *repository.ProductStore, bindings *repository.BindingStore, boundRepoReader productregistry.BoundRepositoryReader, moduleCandidate *candidate.ModuleCandidateRead) *QueryService {
	return &QueryService{products: products, bindings: bindings, boundRepoReader: boundRepoReader, moduleCandidate: moduleCandidate}
}

// ListProducts 承接 ProductListRead。
//
// 列表读取至少承接（phase04-04）：
// name / description / status / created_at / module_bind_count / repository_bind_count
//
// 筛选只允许 queryText / statusFilter。
// 排序按 created_at DESC（phase04-04 冻结）。
func (s *QueryService) ListProducts(ctx context.Context, q productregistry.ListProductsQuery) ([]productregistry.ProductListItem, error) {
	statusFilter := ""
	if q.StatusFilter != "" && q.StatusFilter != "all" {
		statusFilter = string(q.StatusFilter)
	}

	products, err := s.products.List(ctx, q.QueryText, statusFilter)
	if err != nil {
		return nil, err
	}

	items := make([]productregistry.ProductListItem, 0, len(products))
	for _, p := range products {
		moduleCount, err := s.products.CountModuleBindings(ctx, p.ID)
		if err != nil {
			return nil, fmt.Errorf("count module bindings for product %s: %w", p.ID, err)
		}
		repoCount, err := s.products.CountRepositoryBindings(ctx, p.ID)
		if err != nil {
			return nil, fmt.Errorf("count repository bindings for product %s: %w", p.ID, err)
		}

		items = append(items, productregistry.ProductListItem{
			ID:                  p.ID,
			Name:                p.Name,
			Description:         p.Description,
			Status:              p.Status,
			CreatedAt:           p.CreatedAt,
			ModuleBindCount:     moduleCount,
			RepositoryBindCount: repoCount,
		})
	}
	return items, nil
}

// GetProductDetail 承接 ProductDetailRead。
//
// 详情读取统一承接（phase04-04）：
//   - Product 核心字段
//   - 已绑定 Module 列表（通过 ProductModuleSummaryRead 注入，owner=Product Registry）
//   - 已绑定 Repository 列表（通过 ProductRepositorySummaryRead 注入，owner=Repository Binding）
//
// service 层不直接写跨模块 SQL（对齐 phase04-07 L181）。
// 候选读取结果不内嵌于此。
func (s *QueryService) GetProductDetail(ctx context.Context, productID string) (*productregistry.ProductDetail, error) {
	// 校验 ID 格式
	if err := productregistry.ValidateProductID(productID); err != nil {
		return nil, err
	}

	// 读取产品核心字段
	p, err := s.products.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// 读取已绑定 Module 摘要（通过注入的本模块 BindingStore，owner=Product Registry）
	boundModules, err := s.bindings.ListBoundModulesByProduct(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("list bound modules: %w", err)
	}

	// 读取已绑定 Repository 摘要（通过注入的 BoundRepositoryReader 接口，owner=Repository Binding）
	boundRepos, err := s.boundRepoReader.ListBoundRepositoriesByProduct(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("list bound repositories: %w", err)
	}

	// 确保切片非 nil
	if boundModules == nil {
		boundModules = []productregistry.BoundModuleSummary{}
	}
	if boundRepos == nil {
		boundRepos = []productregistry.BoundRepositorySummary{}
	}

	return &productregistry.ProductDetail{
		Product:           *p,
		BoundModules:      boundModules,
		BoundRepositories: boundRepos,
	}, nil
}

// ListProductModuleCandidates 承接 ProductModuleCandidateRead。
//
// 读取可绑定到指定产品的 Module 候选，排除已绑定的 Module。
// 排序按 modules.created_at DESC（phase04-04 冻结）。
// 无可关联候选时返回空列表，不返回错误。
func (s *QueryService) ListProductModuleCandidates(ctx context.Context, productID string) ([]productregistry.ProductModuleCandidate, error) {
	// 校验 ID 格式
	if err := productregistry.ValidateProductID(productID); err != nil {
		return nil, err
	}

	// 校验产品存在性（候选读取前提）
	if _, err := s.products.GetByID(ctx, productID); err != nil {
		return nil, err
	}

	candidates, err := s.moduleCandidate.List(ctx, productID)
	if err != nil {
		return nil, err
	}
	if candidates == nil {
		candidates = []productregistry.ProductModuleCandidate{}
	}
	return candidates, nil
}
