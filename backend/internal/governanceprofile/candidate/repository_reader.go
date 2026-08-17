// Package candidate — governanceprofile 跨模块只读 reader。
//
// 跨模块读取全部通过本子包隔离，service 层不直接写跨模块 SQL
// （对齐工程约定：候选/前提读取接口由消费方模块的 candidate 子包定义和拥有）。
//
// 文件落点：backend/internal/governanceprofile/candidate/repository_reader.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/governanceprofile"
)

// RepositoryReader 承接 repositories 表的存在性前提校验（owner=Repository Binding）。
type RepositoryReader struct {
	pool *pgxpool.Pool
}

// NewRepositoryReader 构造 RepositoryReader。
func NewRepositoryReader(pool *pgxpool.Pool) *RepositoryReader {
	return &RepositoryReader{pool: pool}
}

// EnsureRepositoryExists 校验目标仓库已在 PSCO 中登记。
// 仓库不存在时返回 ErrRepositoryNotFound。
func (r *RepositoryReader) EnsureRepositoryExists(ctx context.Context, repositoryID string) error {
	var exists bool
	err := r.pool.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM repositories WHERE id = $1)
	`, repositoryID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("governanceprofile: check repository exists: %w", err)
	}
	if !exists {
		return governanceprofile.ErrRepositoryNotFound
	}
	return nil
}
