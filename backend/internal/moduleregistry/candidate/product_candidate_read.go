// Package candidate — Product 候选读取（phase02 临时承接）。
//
// §9.6 跨模块候选读取临时承接：
//   - ProductBindingCandidateRead 在 phase02 阶段由 Module Registry 后端模块临时承接
//   - 通过独立 Read 接口定义与独立代码落点隔离，不在 service 层直接写跨模块 SQL
//   - phase03 实现 Product 模块后，本文件可整体迁移到 Product 模块，但接口契约保持不变
//
// 文件落点：backend/internal/moduleregistry/candidate/product_candidate_read.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
)

// ProductCandidateRead 承接 ProductBindingCandidateRead 接口。
//
// 只服务 BindModuleToProduct 的候选 Product 选择，不承接产品主线写入。
type ProductCandidateRead struct {
	pool *pgxpool.Pool
}

// NewProductCandidateRead 构造 ProductCandidateRead。
func NewProductCandidateRead(pool *pgxpool.Pool) *ProductCandidateRead {
	return &ProductCandidateRead{pool: pool}
}

// List 读取全部 Product 候选，按 name 升序。
//
// phase02 阶段不实现分页与筛选，只提供最小候选列表。
func (r *ProductCandidateRead) List(ctx context.Context) ([]moduleregistry.ProductCandidate, error) {
	rows, err := r.pool.Query(ctx, `
SELECT id, name
FROM products
ORDER BY name ASC`)
	if err != nil {
		return nil, fmt.Errorf("list product candidates: %w", err)
	}
	defer rows.Close()

	var items []moduleregistry.ProductCandidate
	for rows.Next() {
		var p moduleregistry.ProductCandidate
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, fmt.Errorf("scan product candidate row: %w", err)
		}
		items = append(items, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter product candidate rows: %w", err)
	}
	return items, nil
}
