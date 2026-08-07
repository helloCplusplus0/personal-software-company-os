// Package repository — repositories 表数据访问层。
//
// 只负责持久化与读取，不承接业务校验或跨模块编排（对齐 phase04-07 分层语义）。
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/repositorybinding"
)

// RepositoryStore 承接 repositories 表的读写。
type RepositoryStore struct {
	pool *pgxpool.Pool
}

// NewRepositoryStore 构造 RepositoryStore。
func NewRepositoryStore(pool *pgxpool.Pool) *RepositoryStore {
	return &RepositoryStore{pool: pool}
}

// Create 插入一条仓库记录，返回带 id / created_at 的完整对象。
func (s *RepositoryStore) Create(ctx context.Context, name, url, provider string, status repositorybinding.RepositoryStatus) (*repositorybinding.Repository, error) {
	r := &repositorybinding.Repository{}
	err := s.pool.QueryRow(ctx, `
INSERT INTO repositories (name, url, provider, status)
VALUES ($1, $2, $3, $4)
RETURNING id, name, url, provider, status, created_at`,
		name, url, provider, string(status),
	).Scan(&r.ID, &r.Name, &r.URL, &r.Provider, &r.Status, &r.CreatedAt)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("%w: repository name already exists", repositorybinding.ErrInvalidInput)
		}
		return nil, fmt.Errorf("insert repository: %w", err)
	}
	return r, nil
}

// GetByID 按 id 读取单个仓库。未找到时返回 repositorybinding.ErrRepositoryNotFound。
func (s *RepositoryStore) GetByID(ctx context.Context, id string) (*repositorybinding.Repository, error) {
	r := &repositorybinding.Repository{}
	err := s.pool.QueryRow(ctx, `
SELECT id, name, url, provider, status, created_at
FROM repositories
WHERE id = $1`,
		id,
	).Scan(&r.ID, &r.Name, &r.URL, &r.Provider, &r.Status, &r.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repositorybinding.ErrRepositoryNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get repository by id: %w", err)
	}
	return r, nil
}

// List 按筛选条件读取仓库列表，返回核心字段。
//
// statusFilter 为空字符串时不过滤状态。queryText 只对 name 做 ILIKE 模糊匹配。
// 结果按 created_at DESC 排序（phase04-04 冻结）。
func (s *RepositoryStore) List(ctx context.Context, queryText, statusFilter string) ([]*repositorybinding.Repository, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, name, url, provider, status, created_at
FROM repositories
WHERE ($1 = '' OR status = $1)
  AND ($2 = '' OR name ILIKE '%' || $2 || '%')
ORDER BY created_at DESC`,
		statusFilter, queryText,
	)
	if err != nil {
		return nil, fmt.Errorf("list repositories: %w", err)
	}
	defer rows.Close()

	var items []*repositorybinding.Repository
	for rows.Next() {
		r := &repositorybinding.Repository{}
		if err := rows.Scan(&r.ID, &r.Name, &r.URL, &r.Provider, &r.Status, &r.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan repository row: %w", err)
		}
		items = append(items, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter repository rows: %w", err)
	}
	return items, nil
}

// CountProductBindings 统计指定仓库的产品绑定数。
func (s *RepositoryStore) CountProductBindings(ctx context.Context, repositoryID string) (int, error) {
	var count int
	err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM product_repositories WHERE repository_id = $1`,
		repositoryID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count product bindings: %w", err)
	}
	return count, nil
}

// CountModuleBindings 统计指定仓库的模块映射数。
func (s *RepositoryStore) CountModuleBindings(ctx context.Context, repositoryID string) (int, error) {
	var count int
	err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM module_repositories WHERE repository_id = $1`,
		repositoryID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count module bindings: %w", err)
	}
	return count, nil
}
