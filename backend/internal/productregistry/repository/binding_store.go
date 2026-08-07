// Package repository — product_modules 摘要读取层。
//
// 承接 ProductDetailRead 所需的本模块关系摘要读取：
//   - ProductModuleSummaryRead：读取已绑定 Module 摘要（product_modules JOIN modules）
//
// owner = 关系表 owner（product_modules 由 Product Registry 拥有，phase04-03 / phase04-07 冻结）。
// service 层通过注入本 Read 接口拼装详情，不直接写跨模块 SQL。
//
// 注意：ProductRepositorySummaryRead（product_repositories JOIN repositories，按 product_id）
// 的 owner 是 Repository Binding，不在本文件实现。Product Registry 的 service 层通过
// productregistry.BoundRepositoryReader 接口注入消费（phase04-07 L162-181 冻结）。
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/productregistry"
)

// BindingStore 承接 product_modules 的摘要读取。
//
// 仅承接本模块关系表（product_modules）的摘要读取，不读 product_repositories
// （该关系表摘要读取 owner 是 Repository Binding）。
type BindingStore struct {
	pool *pgxpool.Pool
}

// NewBindingStore 构造 BindingStore。
func NewBindingStore(pool *pgxpool.Pool) *BindingStore {
	return &BindingStore{pool: pool}
}

// ListBoundModulesByProduct 读取指定产品的已绑定 Module 摘要（ProductModuleSummaryRead）。
//
// owner = product_modules（由 Product Registry 拥有，phase04-07 冻结）。
// 读取 modules 表是为了获取 module_name / module_status（跨包复用 moduleregistry.ModuleStatus）。
// 排序按 modules.name 升序（已绑定摘要按 name 升序，与现有 moduleregistry 模式一致）。
func (s *BindingStore) ListBoundModulesByProduct(ctx context.Context, productID string) ([]productregistry.BoundModuleSummary, error) {
	rows, err := s.pool.Query(ctx, `
SELECT pm.module_id, m.name, m.status
FROM product_modules pm
JOIN modules m ON m.id = pm.module_id
WHERE pm.product_id = $1
ORDER BY m.name`,
		productID,
	)
	if err != nil {
		return nil, fmt.Errorf("list bound modules by product: %w", err)
	}
	defer rows.Close()

	var items []productregistry.BoundModuleSummary
	for rows.Next() {
		var b productregistry.BoundModuleSummary
		if err := rows.Scan(&b.ModuleID, &b.ModuleName, &b.ModuleStatus); err != nil {
			return nil, fmt.Errorf("scan bound module row: %w", err)
		}
		items = append(items, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter bound module rows: %w", err)
	}
	return items, nil
}

// 确保编译期 moduleregistry 包被引用（module_status 类型复用 moduleregistry.ModuleStatus）。
var _ moduleregistry.ModuleStatus
