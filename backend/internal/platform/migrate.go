// Package platform — 数据库迁移运行器。
//
// 从文件系统读取 SQL 迁移文件并按文件名顺序执行，使用 schema_migrations 表记录已应用版本。
// 迁移在事务内执行，失败时回滚，保证幂等可重复运行。
//
// 设计权衡：未使用 go:embed 嵌入迁移文件，而是从文件系统读取，
// 以保持 database/migrations/ 作为单一真相源（与 architecture_map.md 目录落点一致），
// 并允许在不重新编译 Go 的情况下调整迁移。
package platform

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// RunMigrations 执行指定目录下所有 .sql 迁移文件，按文件名升序逐个应用。
//
// 流程：
//  1. 确保 schema_migrations 表存在
//  2. 读取目录下所有 *.sql 文件并按文件名排序
//  3. 对每个文件：若已记录则跳过；否则在事务内执行并记录版本
//
// 返回本次新应用的迁移版本列表。
func RunMigrations(ctx context.Context, pool *pgxpool.Pool, migrationsDir string) ([]string, error) {
	// 1. 确保追踪表存在
	if _, err := pool.Exec(ctx, `
CREATE TABLE IF NOT EXISTS schema_migrations (
    version    TEXT PRIMARY KEY,
    applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)`); err != nil {
		return nil, fmt.Errorf("ensure schema_migrations table: %w", err)
	}

	// 2. 读取迁移文件
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, fmt.Errorf("read migrations dir %s: %w", migrationsDir, err)
	}

	var files []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		files = append(files, e.Name())
	}
	sort.Strings(files)

	if len(files) == 0 {
		slog.Warn("no migration files found", "dir", migrationsDir)
		return nil, nil
	}

	// 3. 逐个应用
	var applied []string
	for _, name := range files {
		version := strings.TrimSuffix(name, ".sql")
		ok, err := isMigrationApplied(ctx, pool, version)
		if err != nil {
			return applied, fmt.Errorf("check migration %s: %w", version, err)
		}
		if ok {
			slog.Debug("migration already applied, skip", "version", version)
			continue
		}

		content, err := os.ReadFile(filepath.Join(migrationsDir, name))
		if err != nil {
			return applied, fmt.Errorf("read migration file %s: %w", name, err)
		}

		if err := applyMigration(ctx, pool, version, string(content)); err != nil {
			return applied, fmt.Errorf("apply migration %s: %w", version, err)
		}
		applied = append(applied, version)
		slog.Info("migration applied", "version", version)
	}

	return applied, nil
}

// isMigrationApplied 查询某版本是否已记录在 schema_migrations。
func isMigrationApplied(ctx context.Context, pool *pgxpool.Pool, version string) (bool, error) {
	var exists bool
	err := pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)`,
		version,
	).Scan(&exists)
	return exists, err
}

// applyMigration 在单个事务内执行迁移 SQL 并记录版本。
//
// 注意：PostgreSQL 事务内不能执行 VACUUM、CREATE DATABASE 等命令，
// 当前 phase02 迁移只含 CREATE TABLE / INDEX / CONSTRAINT，均在事务内安全。
func applyMigration(ctx context.Context, pool *pgxpool.Pool, version, sql string) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx) // commit 后 rollback 为 no-op

	if _, err := tx.Exec(ctx, sql); err != nil {
		return fmt.Errorf("exec sql: %w", err)
	}
	if _, err := tx.Exec(ctx,
		`INSERT INTO schema_migrations (version) VALUES ($1)`,
		version,
	); err != nil {
		return fmt.Errorf("record version: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

// RunSeed 执行单个种子数据 SQL 文件。幂等由 SQL 文件自身保证（ON CONFLICT 等）。
//
// 用于 phase02-12 联调前初始化最小候选数据。
func RunSeed(ctx context.Context, pool *pgxpool.Pool, seedsDir, filename string) error {
	path := filepath.Join(seedsDir, filename)
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read seed file %s: %w", path, err)
	}
	if _, err := pool.Exec(ctx, string(content)); err != nil {
		return fmt.Errorf("exec seed %s: %w", filename, err)
	}
	slog.Info("seed applied", "file", filename)
	return nil
}

// RunSeeds 执行指定目录下所有 .sql 种子文件，按文件名升序逐个应用。
//
// 与 RunMigrations 不同，种子执行不记录版本追踪表：
//   - 幂等性由种子 SQL 自身保证（ON CONFLICT DO NOTHING / 守卫块）
//   - 每次调用都会重新执行所有种子文件，依赖 SQL 幂等避免重复数据
//
// 执行策略说明（与 database/scripts/run_seeds.sh 的差异）：
//   - 后端 RunSeeds（本函数）：执行 seedsDir 下全部 .sql 文件，包括 fixture。
//     fixture 自带守卫块（无模块时跳过），因此全量执行是安全的。
//     适用于 RUN_SEEDS_ON_BOOT=true 本地开发场景的"一键就绪"语义。
//   - 脚本 run_seeds.sh：默认只执行 seed_readonly_prereqs.sql，
//     fixture 需通过 RUN_DECISION_LINK_FIXTURE=1 显式启用。
//     适用于需要精细控制的显式操作场景。
//   两个入口的语义差异是有意设计，非缺陷：后端启动追求"零干预就绪"，
//   脚本追求"显式可控"。fixture 的守卫块保证两者在任何场景下都不会报错。
//
// 返回本次执行的种子文件名列表。
//
// 上游规格：phase02-11 spec §"迁移与最小种子数据必须可支撑 phase02-12 验收"
func RunSeeds(ctx context.Context, pool *pgxpool.Pool, seedsDir string) ([]string, error) {
	entries, err := os.ReadDir(seedsDir)
	if err != nil {
		return nil, fmt.Errorf("read seeds dir %s: %w", seedsDir, err)
	}

	var files []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		files = append(files, e.Name())
	}
	sort.Strings(files)

	if len(files) == 0 {
		slog.Warn("no seed files found", "dir", seedsDir)
		return nil, nil
	}

	var applied []string
	for _, name := range files {
		if err := RunSeed(ctx, pool, seedsDir, name); err != nil {
			return applied, fmt.Errorf("run seed %s: %w", name, err)
		}
		applied = append(applied, name)
	}
	return applied, nil
}
