// Package repository — products 表数据访问层。
//
// 只负责持久化与读取，不承接业务校验或跨模块编排（对齐 phase04-07 分层语义）。
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/productregistry"
)

// ProductStore 承接 products 表的读写。
type ProductStore struct {
	pool *pgxpool.Pool
}

// NewProductStore 构造 ProductStore。
func NewProductStore(pool *pgxpool.Pool) *ProductStore {
	return &ProductStore{pool: pool}
}

// Create 插入一条产品记录，返回带 id / created_at 的完整对象。
func (s *ProductStore) Create(ctx context.Context, name, description string, status productregistry.ProductStatus) (*productregistry.Product, error) {
	p := &productregistry.Product{}
	err := s.pool.QueryRow(ctx, `
INSERT INTO products (name, description, status)
VALUES ($1, $2, $3)
RETURNING id, name, description, status, created_at`,
		name, description, string(status),
	).Scan(&p.ID, &p.Name, &p.Description, &p.Status, &p.CreatedAt)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("%w: product name already exists", productregistry.ErrInvalidInput)
		}
		return nil, fmt.Errorf("insert product: %w", err)
	}
	return p, nil
}

// GetByID 按 id 读取单个产品。未找到时返回 productregistry.ErrProductNotFound。
func (s *ProductStore) GetByID(ctx context.Context, id string) (*productregistry.Product, error) {
	p := &productregistry.Product{}
	err := s.pool.QueryRow(ctx, `
SELECT id, name, description, status, created_at
FROM products
WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.Name, &p.Description, &p.Status, &p.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, productregistry.ErrProductNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get product by id: %w", err)
	}
	return p, nil
}

// List 按筛选条件读取产品列表，返回核心字段。
//
// statusFilter 为空字符串时不过滤状态。queryText 只对 name 做 ILIKE 模糊匹配。
// 结果按 created_at DESC 排序（phase04-04 冻结）。
func (s *ProductStore) List(ctx context.Context, queryText, statusFilter string) ([]*productregistry.Product, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, name, description, status, created_at
FROM products
WHERE ($1 = '' OR status = $1)
  AND ($2 = '' OR name ILIKE '%' || $2 || '%')
ORDER BY created_at DESC`,
		statusFilter, queryText,
	)
	if err != nil {
		return nil, fmt.Errorf("list products: %w", err)
	}
	defer rows.Close()

	var items []*productregistry.Product
	for rows.Next() {
		p := &productregistry.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Status, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan product row: %w", err)
		}
		items = append(items, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter product rows: %w", err)
	}
	return items, nil
}

// CountModuleBindings 统计指定产品的模块绑定数。
func (s *ProductStore) CountModuleBindings(ctx context.Context, productID string) (int, error) {
	var count int
	err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM product_modules WHERE product_id = $1`,
		productID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count module bindings: %w", err)
	}
	return count, nil
}

// CountRepositoryBindings 统计指定产品的仓库绑定数。
func (s *ProductStore) CountRepositoryBindings(ctx context.Context, productID string) (int, error) {
	var count int
	err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM product_repositories WHERE product_id = $1`,
		productID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count repository bindings: %w", err)
	}
	return count, nil
}

// BindModule 插入产品-模块绑定关系。重复绑定返回 ErrDuplicateBinding。
func (s *ProductStore) BindModule(ctx context.Context, productID, moduleID string) error {
	tag, err := s.pool.Exec(ctx, `
INSERT INTO product_modules (product_id, module_id)
VALUES ($1, $2)
ON CONFLICT (product_id, module_id) DO NOTHING`,
		productID, moduleID,
	)
	if err != nil {
		return fmt.Errorf("insert product-module binding: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return productregistry.ErrDuplicateBinding
	}
	return nil
}
