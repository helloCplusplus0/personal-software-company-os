// Package repository — modules 表数据访问层。
//
// 只负责持久化与读取，不承接业务校验或跨模块编排（§9.3 分层语义）。
// 文件落点对齐 phase02-08 spec §"数据访问层文件落点"。
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
)

// ModuleStore 承接 modules 表的读写。
type ModuleStore struct {
	pool *pgxpool.Pool
}

// NewModuleStore 构造 ModuleStore。
func NewModuleStore(pool *pgxpool.Pool) *ModuleStore {
	return &ModuleStore{pool: pool}
}

// Create 插入一条模块记录，返回带 id / created_at 的完整对象。
func (s *ModuleStore) Create(ctx context.Context, name, description string, status moduleregistry.ModuleStatus) (*moduleregistry.Module, error) {
	m := &moduleregistry.Module{}
	err := s.pool.QueryRow(ctx, `
INSERT INTO modules (name, description, status)
VALUES ($1, $2, $3)
RETURNING id, name, description, status, created_at`,
		name, description, string(status),
	).Scan(&m.ID, &m.Name, &m.Description, &m.Status, &m.CreatedAt)
	if err != nil {
		// UNIQUE(name) 冲突映射为业务错误（兜底 service 层的 pre-check，防 TOCTOU）
		if isUniqueViolation(err) {
			return nil, moduleregistry.ErrDuplicateModuleName
		}
		return nil, fmt.Errorf("insert module: %w", err)
	}
	return m, nil
}

// GetByID 按 id 读取单个模块。未找到时返回 moduleregistry.ErrModuleNotFound。
func (s *ModuleStore) GetByID(ctx context.Context, id string) (*moduleregistry.Module, error) {
	m := &moduleregistry.Module{}
	err := s.pool.QueryRow(ctx, `
SELECT id, name, description, status, created_at
FROM modules
WHERE id = $1`,
		id,
	).Scan(&m.ID, &m.Name, &m.Description, &m.Status, &m.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, moduleregistry.ErrModuleNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get module by id: %w", err)
	}
	return m, nil
}

// GetByName 按 name 读取单个模块。用于名称唯一性校验。
func (s *ModuleStore) GetByName(ctx context.Context, name string) (*moduleregistry.Module, error) {
	m := &moduleregistry.Module{}
	err := s.pool.QueryRow(ctx, `
SELECT id, name, description, status, created_at
FROM modules
WHERE name = $1`,
		name,
	).Scan(&m.ID, &m.Name, &m.Description, &m.Status, &m.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil // 名称可用
	}
	if err != nil {
		return nil, fmt.Errorf("get module by name: %w", err)
	}
	return m, nil
}

// List 按筛选条件读取模块列表，返回核心字段。
//
// statusFilter 为空字符串时不过滤状态。queryText 走 ILIKE 模糊匹配 name / description。
// 结果按 created_at DESC 排序。
func (s *ModuleStore) List(ctx context.Context, queryText, statusFilter string) ([]*moduleregistry.Module, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, name, description, status, created_at
FROM modules
WHERE ($1 = '' OR status = $1)
  AND ($2 = '' OR name ILIKE '%' || $2 || '%' OR description ILIKE '%' || $2 || '%')
ORDER BY created_at DESC`,
		statusFilter, queryText,
	)
	if err != nil {
		return nil, fmt.Errorf("list modules: %w", err)
	}
	defer rows.Close()

	var items []*moduleregistry.Module
	for rows.Next() {
		m := &moduleregistry.Module{}
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.Status, &m.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan module row: %w", err)
		}
		items = append(items, m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter module rows: %w", err)
	}
	return items, nil
}

// CountProductBindings 统计指定模块的产品绑定数。
func (s *ModuleStore) CountProductBindings(ctx context.Context, moduleID string) (int, error) {
	var count int
	err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM product_modules WHERE module_id = $1`,
		moduleID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count product bindings: %w", err)
	}
	return count, nil
}

// CountRepositoryBindings 统计指定模块的仓库映射数。
func (s *ModuleStore) CountRepositoryBindings(ctx context.Context, moduleID string) (int, error) {
	var count int
	err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM module_repositories WHERE module_id = $1`,
		moduleID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count repository bindings: %w", err)
	}
	return count, nil
}
