// Package platform — 数据库连接管理。
package platform

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDBPool 创建 PostgreSQL 连接池。
//
// 使用 pgxpool 而非单连接，以支撑后续并发读写。
// 连接配置：合理 MaxConns、连接超时与空闲超时，避免本地开发时长连接堆积。
//
// 调用方负责在关闭时调用 pool.Close()。
func NewDBPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database url: %w", err)
	}

	// 本地开发单服务器场景的保守池配置
	config.MaxConns = 10
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 15 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	// 主动 ping 一次，确保连接可用
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	slog.Info("database connection pool ready", "max_conns", config.MaxConns)
	return pool, nil
}
