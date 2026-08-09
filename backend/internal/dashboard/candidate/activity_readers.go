// Package candidate — Recent Activity reader（由 Dashboard 拥有）。
//
// 承接 RecentActivityRead 所需的跨模块 reader：
//   - ModuleActivityReader：返回 Module 与 Release 的最近活动
//   - ProductActivityReader：返回 Product 与 product_module_binding 的最近活动
//   - RepositoryActivityReader：返回 Repository 与 product_repository_binding / module_repository_binding 的最近活动
//   - DecisionActivityReader：返回 Decision 的最近活动
//
// 各 reader 返回带显式 activity_at 的原始活动项，由 service 层统一映射为 RecentActivityItem。
// 文件落点：backend/internal/dashboard/candidate/activity_readers.go
package candidate

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/dashboard"
)

// RawActivityItem 原始活动项。
// 由各 activity reader 返回，由 service 层映射为 dashboard.RecentActivityItem。
type RawActivityItem struct {
	// Source 标识活动来源，用于 service 层映射 activity_type 与 target_type
	// 取值：module / release / product / repository / decision / product_module_binding / product_repository_binding / module_repository_binding
	Source       string
	ActivityAt   time.Time
	TargetID     string
	TargetLabel  string
}

// ActivityReaders 承接 RecentActivityRead 所需的全部跨模块活动 reader。
type ActivityReaders struct {
	pool *pgxpool.Pool
}

// NewActivityReaders 构造 ActivityReaders。
func NewActivityReaders(pool *pgxpool.Pool) *ActivityReaders {
	return &ActivityReaders{pool: pool}
}

// ReadModuleActivities 读取 Module 与 Release 的最近活动。
//
// module 活动：来自 modules 表的 created_at，target 为 module 自身。
// release 活动：来自 module_releases 表的 released_at，target 回落到所属 Module（phase05-03 冻结）。
// 最多返回 10 条（由 service 层在归并后统一裁剪，reader 层返回更多以便归并排序）。
func (r *ActivityReaders) ReadModuleActivities(ctx context.Context) ([]RawActivityItem, error) {
	if isSimulateError("DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR") {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	// module 活动
	rows, err := r.pool.Query(ctx, `
SELECT id, name, created_at
FROM modules
ORDER BY created_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read module activities: %w", err)
	}
	defer rows.Close()

	items := make([]RawActivityItem, 0)
	for rows.Next() {
		var id, name string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &createdAt); err != nil {
			return nil, fmt.Errorf("scan module activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "module",
			ActivityAt:  createdAt,
			TargetID:    id,
			TargetLabel: name,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter module activity rows: %w", err)
	}

	// release 活动（target 回落到所属 Module）
	releaseRows, err := r.pool.Query(ctx, `
SELECT mr.module_id, m.name, mr.released_at
FROM module_releases mr
JOIN modules m ON m.id = mr.module_id
ORDER BY mr.released_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read release activities: %w", err)
	}
	defer releaseRows.Close()

	for releaseRows.Next() {
		var moduleID, moduleName string
		var releasedAt time.Time
		if err := releaseRows.Scan(&moduleID, &moduleName, &releasedAt); err != nil {
			return nil, fmt.Errorf("scan release activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "release",
			ActivityAt:  releasedAt,
			TargetID:    moduleID,
			TargetLabel: moduleName,
		})
	}
	if err := releaseRows.Err(); err != nil {
		return nil, fmt.Errorf("iter release activity rows: %w", err)
	}

	return items, nil
}

// ReadProductActivities 读取 Product 与 product_module_binding 的最近活动。
func (r *ActivityReaders) ReadProductActivities(ctx context.Context) ([]RawActivityItem, error) {
	if isSimulateError("DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR") {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	items := make([]RawActivityItem, 0)

	// product 活动
	rows, err := r.pool.Query(ctx, `
SELECT id, name, created_at
FROM products
ORDER BY created_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read product activities: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &createdAt); err != nil {
			return nil, fmt.Errorf("scan product activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "product",
			ActivityAt:  createdAt,
			TargetID:    id,
			TargetLabel: name,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter product activity rows: %w", err)
	}

	// product_module_binding 活动（target 落到 Product Detail，phase05-03 冻结）
	bindingRows, err := r.pool.Query(ctx, `
SELECT pm.product_id, p.name, pm.created_at
FROM product_modules pm
JOIN products p ON p.id = pm.product_id
ORDER BY pm.created_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read product_module_binding activities: %w", err)
	}
	defer bindingRows.Close()

	for bindingRows.Next() {
		var productID, productName string
		var createdAt time.Time
		if err := bindingRows.Scan(&productID, &productName, &createdAt); err != nil {
			return nil, fmt.Errorf("scan product_module_binding activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "product_module_binding",
			ActivityAt:  createdAt,
			TargetID:    productID,
			TargetLabel: productName,
		})
	}
	if err := bindingRows.Err(); err != nil {
		return nil, fmt.Errorf("iter product_module_binding activity rows: %w", err)
	}

	return items, nil
}

// ReadRepositoryActivities 读取 Repository、product_repository_binding 与 module_repository_binding 的最近活动。
//
// repository 活动 → target = REPOSITORY_DETAIL
// product_repository_binding 活动 → target = REPOSITORY_DETAIL（phase05-03 冻结）
// module_repository_binding 活动 → target = REPOSITORY_DETAIL（phase05-03 冻结）
func (r *ActivityReaders) ReadRepositoryActivities(ctx context.Context) ([]RawActivityItem, error) {
	if isSimulateError("DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR") {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	items := make([]RawActivityItem, 0)

	// repository 活动
	rows, err := r.pool.Query(ctx, `
SELECT id, name, created_at
FROM repositories
ORDER BY created_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read repository activities: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &createdAt); err != nil {
			return nil, fmt.Errorf("scan repository activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "repository",
			ActivityAt:  createdAt,
			TargetID:    id,
			TargetLabel: name,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter repository activity rows: %w", err)
	}

	// product_repository_binding 活动（target 落到 Repository Detail）
	prRows, err := r.pool.Query(ctx, `
SELECT pr.repository_id, r.name, pr.created_at
FROM product_repositories pr
JOIN repositories r ON r.id = pr.repository_id
ORDER BY pr.created_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read product_repository_binding activities: %w", err)
	}
	defer prRows.Close()

	for prRows.Next() {
		var repoID, repoName string
		var createdAt time.Time
		if err := prRows.Scan(&repoID, &repoName, &createdAt); err != nil {
			return nil, fmt.Errorf("scan product_repository_binding activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "product_repository_binding",
			ActivityAt:  createdAt,
			TargetID:    repoID,
			TargetLabel: repoName,
		})
	}
	if err := prRows.Err(); err != nil {
		return nil, fmt.Errorf("iter product_repository_binding activity rows: %w", err)
	}

	// module_repository_binding 活动（target 落到 Repository Detail）
	mrRows, err := r.pool.Query(ctx, `
SELECT mr.repository_id, r.name, mr.created_at
FROM module_repositories mr
JOIN repositories r ON r.id = mr.repository_id
ORDER BY mr.created_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read module_repository_binding activities: %w", err)
	}
	defer mrRows.Close()

	for mrRows.Next() {
		var repoID, repoName string
		var createdAt time.Time
		if err := mrRows.Scan(&repoID, &repoName, &createdAt); err != nil {
			return nil, fmt.Errorf("scan module_repository_binding activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "module_repository_binding",
			ActivityAt:  createdAt,
			TargetID:    repoID,
			TargetLabel: repoName,
		})
	}
	if err := mrRows.Err(); err != nil {
		return nil, fmt.Errorf("iter module_repository_binding activity rows: %w", err)
	}

	return items, nil
}

// ReadDecisionActivities 读取 Decision 的最近活动。
func (r *ActivityReaders) ReadDecisionActivities(ctx context.Context) ([]RawActivityItem, error) {
	if isSimulateError("DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR") {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	rows, err := r.pool.Query(ctx, `
SELECT id, title, created_at
FROM decisions
ORDER BY created_at DESC
LIMIT 10`)
	if err != nil {
		return nil, fmt.Errorf("read decision activities: %w", err)
	}
	defer rows.Close()

	items := make([]RawActivityItem, 0)
	for rows.Next() {
		var id, title string
		var createdAt time.Time
		if err := rows.Scan(&id, &title, &createdAt); err != nil {
			return nil, fmt.Errorf("scan decision activity row: %w", err)
		}
		items = append(items, RawActivityItem{
			Source:      "decision",
			ActivityAt:  createdAt,
			TargetID:    id,
			TargetLabel: title,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter decision activity rows: %w", err)
	}

	return items, nil
}
