// Package candidate — Onboarding 跨模块 reader 接口定义与实现（由 Onboarding 拥有）。
//
// phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"：
//   - reader 接口的定义与实现均由 Onboarding 模块 candidate/ 子包自己拥有
//   - 沿用 phase05 已验证的 Dashboard OverviewReaders 模式
//   - canonical 模块不需要为 Onboarding 新增 candidate 实现
//   - Onboarding candidate/ 实现可以直接读取 canonical 模块的表，但必须在 candidate/ 子包内隔离
//   - Onboarding service/ 层不得直接跨模块写 SQL
//
// 本文件承接 GetFirstRunState 所需的四类 canonical 对象计数 reader。
// 文件落点：backend/internal/onboarding/candidate/first_run_readers.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// FirstRunReaders 承接 GetFirstRunState 所需的全部跨模块计数 reader。
//
// 由 platform 装配点构造并注入到 Onboarding QueryService。
// 每个 reader 直接读取对应 canonical 模块的表，但在本 candidate/ 子包内隔离。
type FirstRunReaders struct {
	pool *pgxpool.Pool
}

// NewFirstRunReaders 构造 FirstRunReaders。
func NewFirstRunReaders(pool *pgxpool.Pool) *FirstRunReaders {
	return &FirstRunReaders{pool: pool}
}

// CanonicalCounts 四类 canonical 对象的当前持久化数量。
// 用于推导 first_run_state（phase06-14 spec §"first_run_state 状态推导"）。
type CanonicalCounts struct {
	ProductCount    int
	RepositoryCount int
	ModuleCount     int
	DecisionCount   int
}

// ReadCanonicalCounts 读取四类 canonical 对象的当前持久化数量。
//
// 状态推导口径（phase06-01 / phase06-12 / phase06-14 冻结）：
//   - 四类都为 0 → not_started
//   - 至少 1 类、但未满四类 → in_progress
//   - 四类都至少 1 条 → completed
//
// 整页失败语义：任一计数 reader 失败时返回 error，
// 由 service 层归一化为 onboarding.ErrFirstRunStateReadFailed。
func (r *FirstRunReaders) ReadCanonicalCounts(ctx context.Context) (*CanonicalCounts, error) {
	counts := &CanonicalCounts{}

	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM products`).Scan(&counts.ProductCount); err != nil {
		return nil, fmt.Errorf("count products: %w", err)
	}

	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM repositories`).Scan(&counts.RepositoryCount); err != nil {
		return nil, fmt.Errorf("count repositories: %w", err)
	}

	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM modules`).Scan(&counts.ModuleCount); err != nil {
		return nil, fmt.Errorf("count modules: %w", err)
	}

	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM decisions`).Scan(&counts.DecisionCount); err != nil {
		return nil, fmt.Errorf("count decisions: %w", err)
	}

	return counts, nil
}
