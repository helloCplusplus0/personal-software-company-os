// Package main 是 PSCO 后端的运行入口。
//
// 启动流程：
//  1. 加载 .env（若存在）+ 环境变量，得到 Config
//  2. 创建 PostgreSQL 连接池并 ping
//  3. 运行数据库迁移（幂等）
//  4. 若 RUN_SEEDS_ON_BOOT=true，执行种子数据（幂等）
//  5. 装配并启动 HTTP 服务器
//  6. 监听 SIGINT/SIGTERM 优雅关闭
//
// 上游规格：phase02-11 spec §"实现 Module Registry 后端最小读写主线"
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/psco/backend/internal/platform"
)

func main() {
	// 初始化结构化日志（log/slog，Go 1.21+ 标准库）
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	// 1. 加载 .env（可选，生产环境应直接注入环境变量）
	if err := godotenv.Load(); err != nil {
		slog.Info("no .env file loaded, relying on environment variables", "error", err)
	}

	cfg, err := platform.LoadConfig()
	if err != nil {
		slog.Error("load config failed", "error", err)
		os.Exit(1)
	}

	// 2. 数据库连接池
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := platform.NewDBPool(rootCtx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("database pool init failed", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	// 3. 运行迁移（幂等可重复）
	applied, err := platform.RunMigrations(rootCtx, pool, cfg.MigrationsDir)
	if err != nil {
		slog.Error("run migrations failed", "error", err)
		os.Exit(1)
	}
	if len(applied) > 0 {
		slog.Info("migrations applied", "count", len(applied), "versions", applied)
	} else {
		slog.Info("no new migrations to apply")
	}

	// 4. 按配置执行种子数据（仅本地开发启用，默认关闭以保证生产安全）
	//    种子 SQL 自身幂等（ON CONFLICT / 守卫块），可重复执行。
	//    也可通过 database/scripts/run_seeds.sh 在启动外独立执行。
	if cfg.RunSeedsOnBoot {
		seeds, err := platform.RunSeeds(rootCtx, pool, cfg.SeedsDir)
		if err != nil {
			slog.Error("run seeds failed", "error", err)
			os.Exit(1)
		}
		if len(seeds) > 0 {
			slog.Info("seeds applied", "count", len(seeds), "files", seeds)
		} else {
			slog.Info("no seed files to apply")
		}
	} else {
		slog.Info("run_seeds_on_boot disabled (set RUN_SEEDS_ON_BOOT=true to enable)")
	}

	// 5. 装配并启动 HTTP 服务器
	server := platform.NewServer(cfg, pool)

	// 6. 监听信号，优雅关闭
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	serverErrCh := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil {
			serverErrCh <- err
		}
	}()

	select {
	case err := <-serverErrCh:
		slog.Error("server failed", "error", err)
		os.Exit(1)
	case sig := <-sigCh:
		slog.Info("received signal, shutting down", "signal", sig)
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("graceful shutdown failed", "error", err)
		os.Exit(1)
	}
	slog.Info("server stopped cleanly")
}
