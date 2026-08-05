// Package repository — module_releases 表数据访问层。
//
// 文件落点对齐 phase02-08 spec §"数据访问层文件落点"。
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
)

// isUniqueViolation 判断是否为 PostgreSQL 唯一约束冲突（SQLSTATE 23505）。
func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

// ReleaseStore 承接 module_releases 表的读写。
type ReleaseStore struct {
	pool *pgxpool.Pool
}

// NewReleaseStore 构造 ReleaseStore。
func NewReleaseStore(pool *pgxpool.Pool) *ReleaseStore {
	return &ReleaseStore{pool: pool}
}

// Create 在指定模块下插入一条版本记录，返回完整对象。
//
// 调用方应先校验 module 存在性；FK 约束提供兜底保护。
func (s *ReleaseStore) Create(ctx context.Context, moduleID, version string, status moduleregistry.ReleaseStatus, releasedAt string) (*moduleregistry.Release, error) {
	r := &moduleregistry.Release{}
	err := s.pool.QueryRow(ctx, `
INSERT INTO module_releases (module_id, version, status, released_at)
VALUES ($1, $2, $3, $4)
RETURNING id, module_id, version, status, released_at`,
		moduleID, version, string(status), releasedAt,
	).Scan(&r.ID, &r.ModuleID, &r.Version, &r.Status, &r.ReleasedAt)
	if err != nil {
		// UNIQUE(module_id, version) 冲突映射为业务错误
		if isUniqueViolation(err) {
			return nil, moduleregistry.ErrDuplicateReleaseVersion
		}
		return nil, fmt.Errorf("insert release: %w", err)
	}
	return r, nil
}

// ListByModule 按 module_id 读取版本列表，按 released_at DESC 排序。
func (s *ReleaseStore) ListByModule(ctx context.Context, moduleID string) ([]moduleregistry.Release, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, module_id, version, status, released_at
FROM module_releases
WHERE module_id = $1
ORDER BY released_at DESC`,
		moduleID,
	)
	if err != nil {
		return nil, fmt.Errorf("list releases by module: %w", err)
	}
	defer rows.Close()

	var items []moduleregistry.Release
	for rows.Next() {
		var r moduleregistry.Release
		if err := rows.Scan(&r.ID, &r.ModuleID, &r.Version, &r.Status, &r.ReleasedAt); err != nil {
			return nil, fmt.Errorf("scan release row: %w", err)
		}
		items = append(items, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter release rows: %w", err)
	}
	return items, nil
}

// GetLatestVersionByModule 取指定模块的最新版本号。无版本时返回 nil。
func (s *ReleaseStore) GetLatestVersionByModule(ctx context.Context, moduleID string) (*string, error) {
	var version *string
	err := s.pool.QueryRow(ctx, `
SELECT version
FROM module_releases
WHERE module_id = $1
ORDER BY released_at DESC
LIMIT 1`,
		moduleID,
	).Scan(&version)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil // 无版本
	}
	if err != nil {
		return nil, fmt.Errorf("get latest release version: %w", err)
	}
	return version, nil
}
