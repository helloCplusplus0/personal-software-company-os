// Package repository — product_modules + module_repositories 表数据访问层。
//
// 统一承接 §4.1 BindModuleToProduct 与 MapModuleToRepository 的关系写入与读取。
// 文件落点对齐 phase02-08 spec §"数据访问层文件落点"。
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
)

// BindingStore 承接 product_modules 与 module_repositories 两张关系表的读写。
type BindingStore struct {
	pool *pgxpool.Pool
}

// NewBindingStore 构造 BindingStore。
func NewBindingStore(pool *pgxpool.Pool) *BindingStore {
	return &BindingStore{pool: pool}
}

// --- 产品绑定 (product_modules) ---

// BindProduct 插入模块-产品绑定关系。重复绑定返回 ErrDuplicateBinding。
func (s *BindingStore) BindProduct(ctx context.Context, moduleID, productID string) error {
	tag, err := s.pool.Exec(ctx, `
INSERT INTO product_modules (module_id, product_id)
VALUES ($1, $2)
ON CONFLICT (module_id, product_id) DO NOTHING`,
		moduleID, productID,
	)
	if err != nil {
		return fmt.Errorf("insert product binding: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return moduleregistry.ErrDuplicateBinding
	}
	return nil
}

// ListProductBindingsByModule 读取指定模块的产品绑定，附带 product name。
func (s *BindingStore) ListProductBindingsByModule(ctx context.Context, moduleID string) ([]moduleregistry.ProductBinding, error) {
	rows, err := s.pool.Query(ctx, `
SELECT pm.product_id, p.name
FROM product_modules pm
JOIN products p ON p.id = pm.product_id
WHERE pm.module_id = $1
ORDER BY p.name`,
		moduleID,
	)
	if err != nil {
		return nil, fmt.Errorf("list product bindings: %w", err)
	}
	defer rows.Close()

	var items []moduleregistry.ProductBinding
	for rows.Next() {
		var b moduleregistry.ProductBinding
		if err := rows.Scan(&b.ProductID, &b.ProductName); err != nil {
			return nil, fmt.Errorf("scan product binding row: %w", err)
		}
		items = append(items, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter product binding rows: %w", err)
	}
	return items, nil
}

// --- 仓库映射 (module_repositories) ---

// MapRepository 插入模块-仓库映射关系。重复映射返回 ErrDuplicateBinding。
func (s *BindingStore) MapRepository(ctx context.Context, moduleID, repositoryID string) error {
	tag, err := s.pool.Exec(ctx, `
INSERT INTO module_repositories (module_id, repository_id)
VALUES ($1, $2)
ON CONFLICT (module_id, repository_id) DO NOTHING`,
		moduleID, repositoryID,
	)
	if err != nil {
		return fmt.Errorf("insert repository mapping: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return moduleregistry.ErrDuplicateBinding
	}
	return nil
}

// ListRepositoryMappingsByModule 读取指定模块的仓库映射，附带 repository name。
func (s *BindingStore) ListRepositoryMappingsByModule(ctx context.Context, moduleID string) ([]moduleregistry.RepositoryMapping, error) {
	rows, err := s.pool.Query(ctx, `
SELECT mr.repository_id, r.name
FROM module_repositories mr
JOIN repositories r ON r.id = mr.repository_id
WHERE mr.module_id = $1
ORDER BY r.name`,
		moduleID,
	)
	if err != nil {
		return nil, fmt.Errorf("list repository mappings: %w", err)
	}
	defer rows.Close()

	var items []moduleregistry.RepositoryMapping
	for rows.Next() {
		var m moduleregistry.RepositoryMapping
		if err := rows.Scan(&m.RepositoryID, &m.RepositoryName); err != nil {
			return nil, fmt.Errorf("scan repository mapping row: %w", err)
		}
		items = append(items, m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter repository mapping rows: %w", err)
	}
	return items, nil
}

// --- Decision 关联（只读前提） ---
//
// Decision 读取内嵌于 ModuleDetailRead，不设独立读接口组（§6.3）。
// 此处提供按 module_id 读取 decision_links 的能力，供 service/query_service.go 在
// 详情读取编排中调用。文件落点为 binding_store.go，不另建 decision_* 文件（§9.5）。

// ListDecisionLinksByModule 读取指定模块的 Decision 关联，附带 decision title。
//
// 实现说明：Decision 关联读取的 SQL 物理落在 repository 层，
// 但业务编排（何时调用、如何组装）由 service/query_service.go 控制。
// 这不违背 §6.3 —— §6.3 禁止的是独立 handler / service / 读接口，
// 而非禁止在已有 store 文件中提供查询方法。
func (s *BindingStore) ListDecisionLinksByModule(ctx context.Context, moduleID string) ([]moduleregistry.DecisionLink, error) {
	rows, err := s.pool.Query(ctx, `
SELECT dl.decision_id, d.title
FROM decision_links dl
JOIN decisions d ON d.id = dl.decision_id
WHERE dl.module_id = $1
ORDER BY d.created_at DESC`,
		moduleID,
	)
	if err != nil {
		return nil, fmt.Errorf("list decision links: %w", err)
	}
	defer rows.Close()

	var items []moduleregistry.DecisionLink
	for rows.Next() {
		var l moduleregistry.DecisionLink
		if err := rows.Scan(&l.DecisionID, &l.DecisionTitle); err != nil {
			return nil, fmt.Errorf("scan decision link row: %w", err)
		}
		items = append(items, l)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter decision link rows: %w", err)
	}
	return items, nil
}

// ProductExists 校验 product 是否存在（绑定前提校验）。
func (s *BindingStore) ProductExists(ctx context.Context, productID string) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)`,
		productID,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check product exists: %w", err)
	}
	return exists, nil
}

// RepositoryExists 校验 repository 是否存在（映射前提校验）。
func (s *BindingStore) RepositoryExists(ctx context.Context, repositoryID string) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM repositories WHERE id = $1)`,
		repositoryID,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check repository exists: %w", err)
	}
	return exists, nil
}
