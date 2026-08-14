// Package candidate — Feedback 信号 reader（由 Dashboard 拥有）。
//
// 承接 FeedbackSignalRead 所需的跨模块 reader：
//   - PendingDecisionSignalReader：读取 proposed 状态的 decisions
//   - ProductAssetCoverageReader：读取 product 绑定覆盖数据
//
// 文件落点：backend/internal/dashboard/candidate/feedback_readers.go
package candidate

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/dashboard"
)

// PendingDecisionData 待决策信号的原始数据。
// 由 PendingDecisionSignalReader 返回，由 service 层归一化为 FeedbackSignal。
//
// phase10-09 起，pending decision 统一 handoff 到 Decision Detail，
// 因此不再区分“已建立 decision_link”和“未建立 decision_link”的两套跳转出口。
type PendingDecisionData struct {
	DecisionID string
	Title      string
	CreatedAt  time.Time
}

// FeedbackReaders 承接 FeedbackSignalRead 所需的跨模块 reader。
type FeedbackReaders struct {
	pool *pgxpool.Pool
}

// NewFeedbackReaders 构造 FeedbackReaders。
func NewFeedbackReaders(pool *pgxpool.Pool) *FeedbackReaders {
	return &FeedbackReaders{pool: pool}
}

// ReadPendingDecisions 读取所有 proposed 状态的 decisions。
//
// pending decision 的状态判定沿用 phase03 Decision Center 已冻结的 proposed status 语义。
// 按 created_at DESC 排序（service 层归一化时用于同优先级回退排序）。
//
// 无 proposed decisions 时返回空列表，不返回错误。
func (r *FeedbackReaders) ReadPendingDecisions(ctx context.Context) ([]PendingDecisionData, error) {
	if isSimulateError("DASHBOARD_SIMULATE_FEEDBACK_ERROR") {
		return nil, dashboard.ErrFeedbackSignalReadFailed
	}

	rows, err := r.pool.Query(ctx, `
SELECT d.id,
       d.title,
       d.created_at
FROM decisions d
WHERE d.status = 'proposed'
ORDER BY d.created_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("read pending decisions: %w", err)
	}
	defer rows.Close()

	items := make([]PendingDecisionData, 0)
	for rows.Next() {
		var d PendingDecisionData
		if err := rows.Scan(&d.DecisionID, &d.Title, &d.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan pending decision row: %w", err)
		}
		items = append(items, d)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter pending decision rows: %w", err)
	}
	return items, nil
}

// ProductGapData 产品资产缺口的原始数据。
type ProductGapData struct {
	ProductID   string
	ProductName string
	GapType     string // "both" / "repository" / "module"
	CreatedAt   time.Time
}

// ProductAssetCoverageData 产品资产覆盖的原始数据。
type ProductAssetCoverageData struct {
	FullyBoundCount        int
	MissingBothCount       int
	MissingRepositoryCount int
	MissingModuleCount     int
	Gaps                   []ProductGapData
}

// ReadProductAssetCoverage 读取 product 绑定覆盖数据。
//
// 缺口判定逻辑：
//   - missing both：product 在 product_modules 与 product_repositories 中均无记录
//   - missing repository：product 在 product_modules 中有记录，但 product_repositories 中无记录
//   - missing module：product 在 product_repositories 中有记录，但 product_modules 中无记录
//   - fully bound：product 在两个表中均有记录
//
// missing_both_bindings 作为独立分类，不回退为两个单缺口的隐式组合。
// 无 products 时返回零计数与空 Gaps，不返回错误。
func (r *FeedbackReaders) ReadProductAssetCoverage(ctx context.Context) (*ProductAssetCoverageData, error) {
	if isSimulateError("DASHBOARD_SIMULATE_FEEDBACK_ERROR") {
		return nil, dashboard.ErrFeedbackSignalReadFailed
	}

	data := &ProductAssetCoverageData{
		Gaps: []ProductGapData{},
	}

	// 读取所有 product 的绑定状态
	rows, err := r.pool.Query(ctx, `
SELECT
  p.id,
  p.name,
  p.created_at,
  EXISTS(SELECT 1 FROM product_modules pm WHERE pm.product_id = p.id) AS has_module,
  EXISTS(SELECT 1 FROM product_repositories pr WHERE pr.product_id = p.id) AS has_repository
FROM products p
ORDER BY p.created_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("read product asset coverage: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			productID     string
			productName   string
			createdAt     time.Time
			hasModule     bool
			hasRepository bool
		)
		if err := rows.Scan(&productID, &productName, &createdAt, &hasModule, &hasRepository); err != nil {
			return nil, fmt.Errorf("scan product asset coverage row: %w", err)
		}

		switch {
		case hasModule && hasRepository:
			data.FullyBoundCount++
		case !hasModule && !hasRepository:
			data.MissingBothCount++
			data.Gaps = append(data.Gaps, ProductGapData{
				ProductID:   productID,
				ProductName: productName,
				GapType:     "both",
				CreatedAt:   createdAt,
			})
		case hasModule && !hasRepository:
			data.MissingRepositoryCount++
			data.Gaps = append(data.Gaps, ProductGapData{
				ProductID:   productID,
				ProductName: productName,
				GapType:     "repository",
				CreatedAt:   createdAt,
			})
		case !hasModule && hasRepository:
			data.MissingModuleCount++
			data.Gaps = append(data.Gaps, ProductGapData{
				ProductID:   productID,
				ProductName: productName,
				GapType:     "module",
				CreatedAt:   createdAt,
			})
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter product asset coverage rows: %w", err)
	}

	return data, nil
}
