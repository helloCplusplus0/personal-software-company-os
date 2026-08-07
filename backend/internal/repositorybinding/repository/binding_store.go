// Package repository — product_repositories + module_repositories 关系写入与摘要读取层。
//
// 承接：
//   - BindRepositoryToProduct 写入（product_repositories）
//   - MapModuleToRepository 写入（module_repositories）
//   - RepositoryProductSummaryRead（product_repositories JOIN products，按 repository_id）
//   - RepositoryModuleSummaryRead（module_repositories JOIN modules，按 repository_id）
//   - ProductRepositorySummaryRead（product_repositories JOIN repositories，按 product_id，phase04-07 L162-181 冻结 owner=Repository Binding）
//
// owner = 关系表 owner（product_repositories / module_repositories 由 Repository Binding 拥有）。
// service 层通过注入这两个 Read 接口拼装详情，不直接写跨模块 SQL（对齐 phase04-07）。
//
// ProductRepositorySummaryRead 虽服务 ProductDetailRead，但 owner 仍是 Repository Binding
// （关系表 owner 决定接口 owner）。Product Registry 通过 productregistry.BoundRepositoryReader
// 接口注入消费，本方法是唯一实现，不在 productregistry/ 内复制第二套（phase04-07 L174）。
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/repositorybinding"
)

// BindingStore 承接 product_repositories 与 module_repositories 的写入与摘要读取。
type BindingStore struct {
	pool *pgxpool.Pool
}

// NewBindingStore 构造 BindingStore。
func NewBindingStore(pool *pgxpool.Pool) *BindingStore {
	return &BindingStore{pool: pool}
}

// --- 仓库-产品绑定 (product_repositories) ---

// BindProduct 插入仓库-产品绑定关系。重复绑定返回 ErrDuplicateBinding。
func (s *BindingStore) BindProduct(ctx context.Context, repositoryID, productID string) error {
	tag, err := s.pool.Exec(ctx, `
INSERT INTO product_repositories (repository_id, product_id)
VALUES ($1, $2)
ON CONFLICT (repository_id, product_id) DO NOTHING`,
		repositoryID, productID,
	)
	if err != nil {
		return fmt.Errorf("insert repository-product binding: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return repositorybinding.ErrDuplicateBinding
	}
	return nil
}

// ListBoundProductsByRepository 读取指定仓库的已绑定 Product 摘要（RepositoryProductSummaryRead）。
//
// owner = product_repositories（由 Repository Binding 拥有）。
// 读取 products 表是为了获取 product_name / product_status。
// 排序按 products.name 升序。
func (s *BindingStore) ListBoundProductsByRepository(ctx context.Context, repositoryID string) ([]repositorybinding.BoundProductSummary, error) {
	rows, err := s.pool.Query(ctx, `
SELECT pr.product_id, p.name, p.status
FROM product_repositories pr
JOIN products p ON p.id = pr.product_id
WHERE pr.repository_id = $1
ORDER BY p.name`,
		repositoryID,
	)
	if err != nil {
		return nil, fmt.Errorf("list bound products by repository: %w", err)
	}
	defer rows.Close()

	var items []repositorybinding.BoundProductSummary
	for rows.Next() {
		var b repositorybinding.BoundProductSummary
		if err := rows.Scan(&b.ProductID, &b.ProductName, &b.ProductStatus); err != nil {
			return nil, fmt.Errorf("scan bound product row: %w", err)
		}
		items = append(items, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter bound product rows: %w", err)
	}
	return items, nil
}

// ListBoundRepositoriesByProduct 读取指定产品的已绑定 Repository 摘要（ProductRepositorySummaryRead）。
//
// owner = product_repositories（由 Repository Binding 拥有，phase04-07 L162-181 冻结）。
// 本方法是 ProductRepositorySummaryRead 的唯一实现，Product Registry 通过
// productregistry.BoundRepositoryReader 接口在 platform 装配点注入消费（phase04-07 L180）。
// 排序按 repositories.name 升序（与 RepositoryProductSummaryRead 一致，已绑定摘要按 name 升序）。
func (s *BindingStore) ListBoundRepositoriesByProduct(ctx context.Context, productID string) ([]productregistry.BoundRepositorySummary, error) {
	rows, err := s.pool.Query(ctx, `
SELECT pr.repository_id, r.name, r.provider, r.status
FROM product_repositories pr
JOIN repositories r ON r.id = pr.repository_id
WHERE pr.product_id = $1
ORDER BY r.name`,
		productID,
	)
	if err != nil {
		return nil, fmt.Errorf("list bound repositories by product: %w", err)
	}
	defer rows.Close()

	var items []productregistry.BoundRepositorySummary
	for rows.Next() {
		var b productregistry.BoundRepositorySummary
		if err := rows.Scan(&b.RepositoryID, &b.RepositoryName, &b.Provider, &b.RepositoryStatus); err != nil {
			return nil, fmt.Errorf("scan bound repository row: %w", err)
		}
		items = append(items, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter bound repository rows: %w", err)
	}
	return items, nil
}

// --- 仓库-模块映射 (module_repositories) ---

// MapModule 插入仓库-模块映射关系。重复映射返回 ErrDuplicateMapping。
func (s *BindingStore) MapModule(ctx context.Context, repositoryID, moduleID string) error {
	tag, err := s.pool.Exec(ctx, `
INSERT INTO module_repositories (repository_id, module_id)
VALUES ($1, $2)
ON CONFLICT (repository_id, module_id) DO NOTHING`,
		repositoryID, moduleID,
	)
	if err != nil {
		return fmt.Errorf("insert repository-module mapping: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return repositorybinding.ErrDuplicateMapping
	}
	return nil
}

// ListMappedModulesByRepository 读取指定仓库的已映射 Module 摘要（RepositoryModuleSummaryRead）。
//
// owner = module_repositories（由 Repository Binding 拥有）。
// 读取 modules 表是为了获取 module_name / module_status（跨包复用 moduleregistry.ModuleStatus）。
// 排序按 modules.name 升序。
func (s *BindingStore) ListMappedModulesByRepository(ctx context.Context, repositoryID string) ([]repositorybinding.MappedModuleSummary, error) {
	rows, err := s.pool.Query(ctx, `
SELECT mr.module_id, m.name, m.status
FROM module_repositories mr
JOIN modules m ON m.id = mr.module_id
WHERE mr.repository_id = $1
ORDER BY m.name`,
		repositoryID,
	)
	if err != nil {
		return nil, fmt.Errorf("list mapped modules by repository: %w", err)
	}
	defer rows.Close()

	var items []repositorybinding.MappedModuleSummary
	for rows.Next() {
		var m repositorybinding.MappedModuleSummary
		if err := rows.Scan(&m.ModuleID, &m.ModuleName, &m.ModuleStatus); err != nil {
			return nil, fmt.Errorf("scan mapped module row: %w", err)
		}
		items = append(items, m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter mapped module rows: %w", err)
	}
	return items, nil
}

// 确保编译期 moduleregistry 包被引用（module_status 类型复用 moduleregistry.ModuleStatus）。
var _ moduleregistry.ModuleStatus
