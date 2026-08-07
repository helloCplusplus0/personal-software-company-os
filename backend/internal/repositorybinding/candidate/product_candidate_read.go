// Package candidate — Product 候选读取（Repository Binding 拥有）。
//
// 对齐 phase04-07 跨模块候选读取边界：
//   - ProductBindingCandidateRead 由 Repository Binding 的 candidate 子包定义和拥有
//   - 通过独立 Read 接口隔离，service 层不直接写跨模块 SQL
//   - 读取 products 表是为了获取候选 Product 列表（排除已绑定的 Product）
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/repositorybinding"
)

// ProductCandidateRead 承接 ProductBindingCandidateRead 接口。
//
// 只服务 BindRepositoryToProduct 的候选 Product 选择，不承接 Product 主线写入。
type ProductCandidateRead struct {
	pool *pgxpool.Pool
}

// NewProductCandidateRead 构造 ProductCandidateRead。
func NewProductCandidateRead(pool *pgxpool.Pool) *ProductCandidateRead {
	return &ProductCandidateRead{pool: pool}
}

// List 读取指定仓库的可绑定 Product 候选，排除已绑定的 Product。
//
// 排序按 products.created_at DESC（phase04-04 冻结：候选读取均按 created_at 降序）。
// 无可关联候选时返回空列表，不返回错误。
func (r *ProductCandidateRead) List(ctx context.Context, repositoryID string) ([]repositorybinding.RepositoryProductCandidate, error) {
	rows, err := r.pool.Query(ctx, `
SELECT p.id, p.name, p.status
FROM products p
WHERE p.status = 'active'
  AND p.id NOT IN (
    SELECT pr.product_id FROM product_repositories pr WHERE pr.repository_id = $1
  )
ORDER BY p.created_at DESC`,
		repositoryID,
	)
	if err != nil {
		return nil, fmt.Errorf("list repository product candidates: %w", err)
	}
	defer rows.Close()

	var items []repositorybinding.RepositoryProductCandidate
	for rows.Next() {
		var c repositorybinding.RepositoryProductCandidate
		if err := rows.Scan(&c.ProductID, &c.ProductName, &c.ProductStatus); err != nil {
			return nil, fmt.Errorf("scan repository product candidate row: %w", err)
		}
		items = append(items, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter repository product candidate rows: %w", err)
	}
	return items, nil
}

// CheckProductExistsActive 校验 Product 存在性与 active 状态（绑定前提校验）。
//
// 返回三态（phase04-04 错误语义前提：不存在 → 404，存在但非 active → 400）：
//   - exists=false, active=false：Product 不存在（service 层映射 ErrProductNotFound → 404）
//   - exists=true, active=false：Product 存在但状态非 active（service 层映射 ErrProductNotActive → 400）
//   - exists=true, active=true：可用于绑定
//
// 单次 SQL 同时取 EXISTS(id) 与 EXISTS(id AND status='active')，避免两次往返。
func (r *ProductCandidateRead) CheckProductExistsActive(ctx context.Context, productID string) (exists, active bool, err error) {
	err = r.pool.QueryRow(ctx, `
SELECT
  EXISTS(SELECT 1 FROM products WHERE id = $1),
  EXISTS(SELECT 1 FROM products WHERE id = $1 AND status = 'active')`,
		productID,
	).Scan(&exists, &active)
	if err != nil {
		return false, false, fmt.Errorf("check product exists and active: %w", err)
	}
	return exists, active, nil
}
