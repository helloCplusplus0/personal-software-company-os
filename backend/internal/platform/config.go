// Package platform 提供 PSCO 后端的横切基础设施：配置、数据库连接、迁移运行器与 HTTP 服务器装配。
//
// 本包不承载任何业务语义，仅服务 internal/moduleregistry 等业务模块。
// 上游规格：phase02-11 spec §"本地开发环境必须复用共享 PostgreSQL 容器"
package platform

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

// Config 承载后端运行所需的环境配置。
//
// 字段语义对齐 phase02-11 spec：
//   - DatabaseURL: 显式最终值，不依赖嵌套环境变量拼接
//   - MigrationsDir: 迁移文件所在目录（默认指向项目根 database/migrations）
//   - SeedsDir: 种子数据文件所在目录（默认指向项目根 database/seeds）
//   - RunSeedsOnBoot: 启动时是否自动执行种子数据（默认 false，仅本地开发启用）
//   - HTTPPort: HTTP 服务监听端口
type Config struct {
	DatabaseURL    string
	MigrationsDir  string
	SeedsDir       string
	RunSeedsOnBoot bool
	HTTPPort       int
}

// LoadConfig 从环境变量加载配置。
//
// DATABASE_URL 必须为显式最终值；若密码含 /、=、@ 等特殊字符，调用方需在写入 env 前完成 URL 编码。
// 敏感凭据不写入仓库源码，应通过本地 .env 文件或运行环境提供。
func LoadConfig() (Config, error) {
	cfg := Config{
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		MigrationsDir:  envOrDefault("MIGRATIONS_DIR", "database/migrations"),
		SeedsDir:       envOrDefault("SEEDS_DIR", "database/seeds"),
		RunSeedsOnBoot: envOrDefault("RUN_SEEDS_ON_BOOT", "false") == "true",
	}

	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required (must be an explicit final value, see phase02-11 spec)")
	}

	port, err := strconv.Atoi(envOrDefault("HTTP_PORT", "8080"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid HTTP_PORT: %w", err)
	}
	cfg.HTTPPort = port

	slog.Info("config loaded",
		"migrations_dir", cfg.MigrationsDir,
		"seeds_dir", cfg.SeedsDir,
		"run_seeds_on_boot", cfg.RunSeedsOnBoot,
		"http_port", cfg.HTTPPort,
	)
	return cfg, nil
}

// envOrDefault 读取环境变量，缺省时返回默认值。
func envOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
