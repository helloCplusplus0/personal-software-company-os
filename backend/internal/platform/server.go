// Package platform — HTTP 服务器装配。
//
// 使用 chi v5 装配根路由器，挂载 Module Registry 模块路由，
// 应用基础中间件（RequestID / Logger / Recoverer / CORS），并启动 HTTP 服务。
package platform

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Server 承载 HTTP 服务及其依赖。
type Server struct {
	httpServer *http.Server
	pool       *pgxpool.Pool
}

// NewServer 装配并返回 Server。
//
// 装配顺序：
//  1. 应用基础中间件（RequestID / Logger / Recoverer / CORS）
//  2. 注册健康检查端点
//  3. 通过 mountModuleRegistry 把 Module Registry 路由挂到 /api 下
//  4. 构造 http.Server
func NewServer(cfg Config, pool *pgxpool.Pool) *Server {
	r := chi.NewRouter()

	// 基础中间件
	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.Timeout(60 * time.Second))
	r.Use(corsMiddleware)

	// 健康检查
	r.Get("/healthz", healthz)

	// 装配 Module Registry 模块路由到 /api
	r.Route("/api", func(r chi.Router) {
		mountModuleRegistry(r, pool)
	})

	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.HTTPPort),
			Handler:      r,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		pool: pool,
	}
}

// Start 启动 HTTP 服务（阻塞调用）。
func (s *Server) Start() error {
	slog.Info("http server starting", "addr", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http server: %w", err)
	}
	return nil
}

// Shutdown 优雅关闭 HTTP 服务并释放数据库连接池。
func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("http server shutting down")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("http shutdown: %w", err)
	}
	s.pool.Close()
	return nil
}

// corsMiddleware 为本地联调提供最小 CORS 支持。
//
// phase02 本地开发：前端 Vite dev server (默认 5173) 与后端 (8080) 跨端口，
// 需要允许前端Origin 携带凭据访问。生产部署由 Caddy 反代统一处理 CORS。
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
