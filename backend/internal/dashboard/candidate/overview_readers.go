// Package candidate — Dashboard 跨模块 reader 接口定义与实现（由 Dashboard 拥有）。
//
// phase05-07 §"DashboardOverviewRead query service owner 与跨模块读取边界冻结"：
//   - reader 接口的定义与实现均由 Dashboard 模块 candidate/ 子包自己拥有
//   - 沿用 phase03 已验证的 DecisionModuleCandidateRead 模式
//   - canonical 模块不需要为 Dashboard 新增 candidate 实现
//   - Dashboard candidate/ 实现可以直接读取 canonical 模块的表，但必须在 candidate/ 子包内隔离
//   - Dashboard service/ 层不得直接跨模块写 SQL
//
// 本文件承接 DashboardOverviewRead 所需的四个计数 reader。
// 文件落点：backend/internal/dashboard/candidate/overview_readers.go
package candidate

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/dashboard"
)

// OverviewReaders 承接 DashboardOverviewRead 所需的全部跨模块计数 reader。
//
// 由 platform 装配点构造并注入到 Dashboard QueryService。
// 每个 reader 直接读取对应 canonical 模块的表，但在本 candidate/ 子包内隔离。
type OverviewReaders struct {
	pool *pgxpool.Pool
}

// NewOverviewReaders 构造 OverviewReaders。
func NewOverviewReaders(pool *pgxpool.Pool) *OverviewReaders {
	return &OverviewReaders{pool: pool}
}

// CountModules 返回 modules 表的总数。
// 对齐 DashboardOverview.module_count。
func (r *OverviewReaders) CountModules(ctx context.Context) (int, error) {
	if isSimulateError("DASHBOARD_SIMULATE_OVERVIEW_ERROR") {
		return 0, dashboard.ErrOverviewReadFailed
	}

	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM modules`).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count modules: %w", err)
	}
	return count, nil
}

// CountProducts 返回 products 总数与已绑定 module / repository 的去重 product 数。
// 对齐 DashboardOverview.product_count / product_with_module_count / product_with_repository_count。
func (r *OverviewReaders) CountProducts(ctx context.Context) (productCount, withModule, withRepository int, err error) {
	if isSimulateError("DASHBOARD_SIMULATE_OVERVIEW_ERROR") {
		return 0, 0, 0, dashboard.ErrOverviewReadFailed
	}

	// product_count
	if err = r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM products`).Scan(&productCount); err != nil {
		return 0, 0, 0, fmt.Errorf("count products: %w", err)
	}

	// product_with_module_count：在 product_modules 中存在记录的去重 product 数
	if err = r.pool.QueryRow(ctx, `SELECT COUNT(DISTINCT product_id) FROM product_modules`).Scan(&withModule); err != nil {
		return 0, 0, 0, fmt.Errorf("count products with module: %w", err)
	}

	// product_with_repository_count：在 product_repositories 中存在记录的去重 product 数
	if err = r.pool.QueryRow(ctx, `SELECT COUNT(DISTINCT product_id) FROM product_repositories`).Scan(&withRepository); err != nil {
		return 0, 0, 0, fmt.Errorf("count products with repository: %w", err)
	}

	return productCount, withModule, withRepository, nil
}

// CountRepositories 返回 repositories 表的总数。
// 对齐 DashboardOverview.repository_count。
func (r *OverviewReaders) CountRepositories(ctx context.Context) (int, error) {
	if isSimulateError("DASHBOARD_SIMULATE_OVERVIEW_ERROR") {
		return 0, dashboard.ErrOverviewReadFailed
	}

	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM repositories`).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count repositories: %w", err)
	}
	return count, nil
}

// CountDecisions 返回 decisions 表的总数。
// 对齐 DashboardOverview.decision_count。
func (r *OverviewReaders) CountDecisions(ctx context.Context) (int, error) {
	if isSimulateError("DASHBOARD_SIMULATE_OVERVIEW_ERROR") {
		return 0, dashboard.ErrOverviewReadFailed
	}

	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM decisions`).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count decisions: %w", err)
	}
	return count, nil
}

// isSimulateError 检查环境变量是否设置为模拟错误模式。
// 用于 phase05-09 冻结的局部错误模拟机制。
// 环境变量值为 "true"（不区分大小写）时触发模拟错误。
func isSimulateError(envKey string) bool {
	v := os.Getenv(envKey)
	if v == "" {
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}
