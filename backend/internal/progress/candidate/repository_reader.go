// Package candidate — progress 跨模块只读 reader。
//
// 跨模块读取全部通过本子包隔离，service 层不直接写跨模块 SQL
// （对齐工程约定：候选/前提读取接口由消费方模块的 candidate 子包定义和拥有；
// 本文件为 phase15-04 DP-2 裁决的承接位，沿 standard/candidate/TargetReader 模式）。
//
// 文件落点：backend/internal/progress/candidate/repository_reader.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// RepositoryReader 承接 repository_id 存在性事实查询（DP-2 承接位）。
// 纯存在性查询（返回 bool），不承载业务错误语义；错误语义包装归 service 层
// （Create → InvalidArgument [REPOSITORY_NOT_FOUND]；List → NotFound 读锚点）。
type RepositoryReader struct {
	pool *pgxpool.Pool
}

// NewRepositoryReader 构造 RepositoryReader。
func NewRepositoryReader(pool *pgxpool.Pool) *RepositoryReader {
	return &RepositoryReader{pool: pool}
}

// RepositoryExists 查询仓库是否已存在。
// 查询失败包装原始错误（Internal，由 connecterrors 兜底），不承载业务判断。
func (r *RepositoryReader) RepositoryExists(ctx context.Context, repositoryID string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM repositories WHERE id = $1)
	`, repositoryID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("progress: check repository exists: %w", err)
	}
	return exists, nil
}
