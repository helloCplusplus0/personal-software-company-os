// Package candidate — ReuseSummary 跨模块 reader 接口定义与实现（由 ReuseSummary 拥有）。
//
// phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"：
//   - reader 接口的定义与实现均由 ReuseSummary 模块 candidate/ 子包自己拥有
//   - ReuseSummary candidate/ 实现可以直接读取 canonical 模块的表
//   - ReuseSummary service/ 层不得直接跨模块写 SQL
//
// 本文件承接三种作用域的复用感知读取。
// 文件落点：backend/internal/reusesummary/candidate/reuse_readers.go
package candidate

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ReuseReaders 承接复用感知所需的三种作用域 reader。
//
// 由 platform 装配点构造并注入到 ReuseSummary QueryService。
type ReuseReaders struct {
	pool *pgxpool.Pool
}

// NewReuseReaders 构造 ReuseReaders。
func NewReuseReaders(pool *pgxpool.Pool) *ReuseReaders {
	return &ReuseReaders{pool: pool}
}

// ModuleReuseData 模块复用原始数据。
type ModuleReuseData struct {
	ModuleID          string
	ModuleName        string
	CapabilityKey     *string
	ReuseProductCount int
	LatestReuseAt     *time.Time
	ModuleCreatedAt   time.Time
}

// CapabilityAggregateData 能力聚合原始数据。
type CapabilityAggregateData struct {
	CapabilityKey         string
	SupportingModuleCount int
	LatestUpdateAt        *time.Time
}

// ReadDashboardReuse 读取 Dashboard 作用域的全局复用快照。
//
// 聚合口径（phase06-14 spec §"module_reuse_summary 聚合口径"）：
//   - reuse_product_count 只表示"当前被多少 Product 直接复用"
//   - 只允许基于 product_modules 与 modules 当前已提交数据聚合
//
// 返回所有模块的复用数据（含 capability_key），由 service 层排序与裁剪。
func (r *ReuseReaders) ReadDashboardReuse(ctx context.Context) ([]ModuleReuseData, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			m.id,
			m.name,
			m.capability_key,
			COALESCE(pm.reuse_count, 0) AS reuse_product_count,
			pm.latest_reuse_at,
			m.created_at
		FROM modules m
		LEFT JOIN (
			SELECT module_id, COUNT(DISTINCT product_id) AS reuse_count, MAX(created_at) AS latest_reuse_at
			FROM product_modules
			GROUP BY module_id
		) pm ON pm.module_id = m.id
		ORDER BY reuse_count DESC, COALESCE(pm.latest_reuse_at, m.created_at) DESC`)
	if err != nil {
		return nil, fmt.Errorf("read dashboard reuse: %w", err)
	}
	defer rows.Close()

	return scanModuleReuseRows(rows)
}

// ReadModuleDetailReuse 读取 Module Detail 作用域的单一模块复用快照。
//
// 围绕单一 module_id 返回该模块的直接复用反馈。
// 若 module_id 不存在，返回 nil + nil error（由 service 层判断为空态）。
func (r *ReuseReaders) ReadModuleDetailReuse(ctx context.Context, moduleID string) (*ModuleReuseData, error) {
	data := &ModuleReuseData{}
	var latestReuseAt *time.Time
	var capabilityKey *string

	err := r.pool.QueryRow(ctx, `
		SELECT
			m.id,
			m.name,
			m.capability_key,
			COALESCE(pm.reuse_count, 0) AS reuse_product_count,
			pm.latest_reuse_at,
			m.created_at
		FROM modules m
		LEFT JOIN (
			SELECT module_id, COUNT(DISTINCT product_id) AS reuse_count, MAX(created_at) AS latest_reuse_at
			FROM product_modules
			GROUP BY module_id
		) pm ON pm.module_id = m.id
		WHERE m.id = $1`, moduleID).Scan(
		&data.ModuleID, &data.ModuleName, &capabilityKey,
		&data.ReuseProductCount, &latestReuseAt, &data.ModuleCreatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, fmt.Errorf("read module detail reuse: %w", err)
	}

	data.CapabilityKey = capabilityKey
	data.LatestReuseAt = latestReuseAt
	return data, nil
}

// ReadProductDetailReuse 读取 Product Detail 作用域的复用快照。
//
// 先限定在当前 Product 已绑定模块范围内，再返回全量复用 / capability 摘要。
func (r *ReuseReaders) ReadProductDetailReuse(ctx context.Context, productID string) ([]ModuleReuseData, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			m.id,
			m.name,
			m.capability_key,
			COALESCE(pm.reuse_count, 0) AS reuse_product_count,
			pm.latest_reuse_at,
			m.created_at
		FROM modules m
		INNER JOIN product_modules ptm ON ptm.module_id = m.id AND ptm.product_id = $1
		LEFT JOIN (
			SELECT module_id, COUNT(DISTINCT product_id) AS reuse_count, MAX(created_at) AS latest_reuse_at
			FROM product_modules
			GROUP BY module_id
		) pm ON pm.module_id = m.id
		ORDER BY reuse_count DESC, COALESCE(pm.latest_reuse_at, m.created_at) DESC`,
		productID)
	if err != nil {
		return nil, fmt.Errorf("read product detail reuse: %w", err)
	}
	defer rows.Close()

	return scanModuleReuseRows(rows)
}

// ReadCapabilityAggregates 读取能力聚合数据。
//
// 聚合口径（phase06-14 spec §"capability_summary 聚合与映射"）：
//   - 以 modules.capability_key 作为唯一聚合主键来源
//   - supporting_module_count 表示当前参与该 capability_key 聚合的 Module 数量
//   - latest_capability_update_at 取该 capability_key 下最新的模块 created_at
//
// moduleIDsFilter 非 nil 时，只聚合指定 moduleID 集合内的模块。
func (r *ReuseReaders) ReadCapabilityAggregates(ctx context.Context, moduleIDsFilter []string) ([]CapabilityAggregateData, error) {
	query := `
		SELECT
			m.capability_key,
			COUNT(*) AS supporting_module_count,
			MAX(m.created_at) AS latest_update_at
		FROM modules m
		WHERE m.capability_key IS NOT NULL AND m.capability_key != ''`

	args := []any{}
	if len(moduleIDsFilter) > 0 {
		placeholders := ""
		for i, id := range moduleIDsFilter {
			if i > 0 {
				placeholders += ","
			}
			placeholders += fmt.Sprintf("$%d", i+1)
			args = append(args, id)
		}
		query += fmt.Sprintf(" AND m.id IN (%s)", placeholders)
	}

	query += `
		GROUP BY m.capability_key
		ORDER BY supporting_module_count DESC, latest_update_at DESC`

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("read capability aggregates: %w", err)
	}
	defer rows.Close()

	var results []CapabilityAggregateData
	for rows.Next() {
		var data CapabilityAggregateData
		if err := rows.Scan(&data.CapabilityKey, &data.SupportingModuleCount, &data.LatestUpdateAt); err != nil {
			return nil, fmt.Errorf("scan capability aggregate: %w", err)
		}
		results = append(results, data)
	}

	if results == nil {
		results = []CapabilityAggregateData{}
	}
	return results, nil
}

// scanModuleReuseRows 扫描模块复用数据行。
func scanModuleReuseRows(rows pgxRows) ([]ModuleReuseData, error) {
	var results []ModuleReuseData
	for rows.Next() {
		var data ModuleReuseData
		var capabilityKey *string
		var latestReuseAt *time.Time
		if err := rows.Scan(&data.ModuleID, &data.ModuleName, &capabilityKey,
			&data.ReuseProductCount, &latestReuseAt, &data.ModuleCreatedAt); err != nil {
			return nil, fmt.Errorf("scan module reuse: %w", err)
		}
		data.CapabilityKey = capabilityKey
		data.LatestReuseAt = latestReuseAt
		results = append(results, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate module reuse rows: %w", err)
	}
	if results == nil {
		results = []ModuleReuseData{}
	}
	return results, nil
}

// pgxRows 抽象 pgx.Rows 接口，便于测试。
type pgxRows interface {
	Next() bool
	Scan(dest ...any) error
	Err() error
	Close()
}
