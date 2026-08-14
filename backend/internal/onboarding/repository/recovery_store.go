// Package repository — Onboarding 恢复存储层。
//
// phase10-08 新增：提供最小 current_product_id 恢复锚点的持久化能力。
// 使用单行 key-value 表，不引入新 ORM 或第二套存储抽象。
//
// 文件落点：backend/internal/onboarding/repository/recovery_store.go
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	// recoveryKeyCurrentProductID 是 current_product_id 锚点的存储键。
	recoveryKeyCurrentProductID = "current_product_id"
)

// RecoveryStore 承接 Onboarding 建链恢复锚点的持久化读写。
//
// 当前阶段只存储 current_product_id 一个键，后续可按需扩展。
// 使用与既有代码相同的 pgxpool 模式。
type RecoveryStore struct {
	pool *pgxpool.Pool
}

// NewRecoveryStore 构造 RecoveryStore。
func NewRecoveryStore(pool *pgxpool.Pool) *RecoveryStore {
	return &RecoveryStore{pool: pool}
}

// EnsureSchema 创建 recovery store 所需的表（如果不存在）。
//
// 应在上线/迁移时调用一次，幂等。
func (s *RecoveryStore) EnsureSchema(ctx context.Context) error {
	_, err := s.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS onboarding_recovery_store (
			key TEXT PRIMARY KEY,
			value TEXT NOT NULL,
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("ensure recovery store schema: %w", err)
	}
	return nil
}

// UpsertCurrentProductID 冻结或更新 current_product_id 锚点。
//
// 使用 INSERT ... ON CONFLICT 保证幂等。
func (s *RecoveryStore) UpsertCurrentProductID(ctx context.Context, productID string) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO onboarding_recovery_store (key, value, updated_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT (key) DO UPDATE SET value = $2, updated_at = NOW()
	`, recoveryKeyCurrentProductID, productID)
	if err != nil {
		return fmt.Errorf("upsert current product id: %w", err)
	}
	return nil
}

// GetCurrentProductID 读取 current_product_id 锚点。
//
// 返回空字符串表示锚点尚未设置。
func (s *RecoveryStore) GetCurrentProductID(ctx context.Context) (string, error) {
	var value string
	err := s.pool.QueryRow(ctx, `
		SELECT value FROM onboarding_recovery_store WHERE key = $1
	`, recoveryKeyCurrentProductID).Scan(&value)
	if err != nil {
		// 行不存在不是错误，返回空字符串
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("get current product id: %w", err)
	}
	return value, nil
}
