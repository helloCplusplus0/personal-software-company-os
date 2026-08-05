// Package candidate — Repository 候选读取（phase02 临时承接）。
//
// §9.6 跨模块候选读取临时承接：
//   - RepositoryBindingCandidateRead 在 phase02 阶段由 Module Registry 后端模块临时承接
//   - 通过独立 Read 接口定义与独立代码落点隔离，不在 service 层直接写跨模块 SQL
//   - phase04 实现 Repository 模块后，本文件可整体迁移到 Repository 模块，但接口契约保持不变
//
// 文件落点：backend/internal/moduleregistry/candidate/repository_candidate_read.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
)

// RepositoryCandidateRead 承接 RepositoryBindingCandidateRead 接口。
//
// 只服务 MapModuleToRepository 的候选 Repository 选择，不承接仓库主线写入。
type RepositoryCandidateRead struct {
	pool *pgxpool.Pool
}

// NewRepositoryCandidateRead 构造 RepositoryCandidateRead。
func NewRepositoryCandidateRead(pool *pgxpool.Pool) *RepositoryCandidateRead {
	return &RepositoryCandidateRead{pool: pool}
}

// List 读取全部 Repository 候选，按 name 升序。
//
// phase02 阶段不实现分页与筛选，只提供最小候选列表。
func (r *RepositoryCandidateRead) List(ctx context.Context) ([]moduleregistry.RepositoryCandidate, error) {
	rows, err := r.pool.Query(ctx, `
SELECT id, name
FROM repositories
ORDER BY name ASC`)
	if err != nil {
		return nil, fmt.Errorf("list repository candidates: %w", err)
	}
	defer rows.Close()

	var items []moduleregistry.RepositoryCandidate
	for rows.Next() {
		var r moduleregistry.RepositoryCandidate
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, fmt.Errorf("scan repository candidate row: %w", err)
		}
		items = append(items, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter repository candidate rows: %w", err)
	}
	return items, nil
}
