// Package candidate — Onboarding 建链状态 reader（由 Onboarding 拥有）。
//
// phase10-08 新增：提供 ReadOnboardingChainState 所需的跨模块读取能力。
// 每个 reader 直接读取对应 canonical 模块的表，但在本 candidate/ 子包内隔离。
//
// 文件落点：backend/internal/onboarding/candidate/chain_state_readers.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ChainStateReaders 承接 ReadOnboardingChainState 所需的全部跨模块读取。
//
// 由 platform 装配点构造并注入到 Onboarding QueryService。
type ChainStateReaders struct {
	pool *pgxpool.Pool
}

// NewChainStateReaders 构造 ChainStateReaders。
func NewChainStateReaders(pool *pgxpool.Pool) *ChainStateReaders {
	return &ChainStateReaders{pool: pool}
}

// ChainStateFacts 建链状态相关的跨模块事实。
type ChainStateFacts struct {
	// HasProduct 是否存在至少一个 Product。
	HasProduct bool
	// HasRepository 是否存在至少一个 Repository。
	HasRepository bool
	// HasRepositoryBound 是否已有 Repository 绑定到当前 Product。
	HasRepositoryBound bool
	// HasModule 是否存在至少一个 Module。
	HasModule bool
	// HasModuleBound 是否已有 Module 绑定到当前 Product。
	HasModuleBound bool
	// HasDecision 是否存在至少一个 Decision。
	HasDecision bool
}

// ReadChainStateFacts 读取建链状态所需的跨模块事实。
//
// 以 currentProductID 为锚点，读取各 canonical 模块的持久化状态。
// currentProductID 为空时，只检查是否存在记录，不检查绑定关系。
func (r *ChainStateReaders) ReadChainStateFacts(ctx context.Context, currentProductID string) (*ChainStateFacts, error) {
	facts := &ChainStateFacts{}

	// 检查 Product 是否存在
	var productCount int
	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM products`).Scan(&productCount); err != nil {
		return nil, fmt.Errorf("count products: %w", err)
	}
	facts.HasProduct = productCount > 0

	// 检查 Repository 是否存在
	var repoCount int
	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM repositories`).Scan(&repoCount); err != nil {
		return nil, fmt.Errorf("count repositories: %w", err)
	}
	facts.HasRepository = repoCount > 0

	// 检查 Module 是否存在
	var moduleCount int
	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM modules`).Scan(&moduleCount); err != nil {
		return nil, fmt.Errorf("count modules: %w", err)
	}
	facts.HasModule = moduleCount > 0

	// 检查 Decision 是否存在
	var decisionCount int
	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM decisions`).Scan(&decisionCount); err != nil {
		return nil, fmt.Errorf("count decisions: %w", err)
	}
	facts.HasDecision = decisionCount > 0

	// 如果有 currentProductID，检查绑定关系
	if currentProductID != "" {
		// 检查 Repository 绑定到 Product
		var boundRepoCount int
		if err := r.pool.QueryRow(ctx,
			`SELECT COUNT(*) FROM product_repositories WHERE product_id = $1`,
			currentProductID,
		).Scan(&boundRepoCount); err != nil {
			return nil, fmt.Errorf("count bound repositories: %w", err)
		}
		facts.HasRepositoryBound = boundRepoCount > 0

		// 检查 Module 绑定到 Product
		var boundModuleCount int
		if err := r.pool.QueryRow(ctx,
			`SELECT COUNT(*) FROM product_modules WHERE product_id = $1`,
			currentProductID,
		).Scan(&boundModuleCount); err != nil {
			return nil, fmt.Errorf("count bound modules: %w", err)
		}
		facts.HasModuleBound = boundModuleCount > 0
	}

	return facts, nil
}

// ReadFirstProductID 读取第一个 Product 的 ID（用于自动冻结锚点）。
//
// 返回空字符串表示尚无 Product。
func (r *ChainStateReaders) ReadFirstProductID(ctx context.Context) (string, error) {
	var id string
	err := r.pool.QueryRow(ctx, `SELECT id FROM products ORDER BY created_at ASC LIMIT 1`).Scan(&id)
	if err != nil {
		// 没有行不是错误
		if err.Error() == "no rows in result set" {
			return "", nil
		}
		return "", fmt.Errorf("read first product id: %w", err)
	}
	return id, nil
}
