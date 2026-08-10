// Package platform — HTTP 服务器装配。
//
// 使用 chi v5 装配根路由器，挂载 Module Registry / Decision Center /
// Product Registry / Repository Binding / Dashboard 模块路由，
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
//  3. 构造 Product Registry / Repository Binding 的 service 层（供 Module Registry 兼容委派复用）
//  4. 通过 mountModuleRegistry 把 Module Registry 路由挂到 /api 下（注入兼容委派目标）
//  5. 通过 mountDecisionCenter 把 Decision Center 路由挂到 /api 下（phase03-12）
//  6. 通过 mountProductRegistry 把 Product Registry 路由挂到 /api 下（phase04-12）
//  7. 通过 mountRepositoryBinding 把 Repository Binding 路由挂到 /api 下（phase04-12）
//  8. 构造 http.Server
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

	// 装配业务模块路由到 /api
	r.Route("/api", func(r chi.Router) {
		// phase04-12: 先构造 Product Registry / Repository Binding 的 service 层，
		// 供 Module Registry 的旧绑定入口做兼容委派注入。
		//
		// 装配顺序约束（phase04-07 L162-181）：
		//   Repository Binding 必须先构造，将其 BindingStore 作为 BoundRepositoryReader
		//   注入到 Product Registry 的 QueryService（ProductRepositorySummaryRead owner=Repository Binding）。
		repoQuerySvc, repoCommandSvc, repoBindingStore := buildRepositoryBinding(pool)
		productQuerySvc, productCommandSvc := buildProductRegistry(pool, repoBindingStore)

		mountModuleRegistry(r, pool, productQuerySvc, productCommandSvc, repoQuerySvc, repoCommandSvc)
		mountDecisionCenter(r, pool)
		mountProductRegistry(r, productQuerySvc, productCommandSvc)
		mountRepositoryBinding(r, repoQuerySvc, repoCommandSvc)

		// phase05-12: Dashboard 必须在既有四个 canonical 模块装配之后装配
		// （Dashboard 跨模块读取依赖 canonical 模块的表已建表）。
		// Dashboard 只承接只读聚合，跨模块读依赖在 platform 装配点注入（phase05-07）。
		dashboardQuerySvc := buildDashboard(pool)
		mountDashboard(r, dashboardQuerySvc)

		// phase06-14: Onboarding / Export / Backup / ReuseSummary 必须在既有
		// canonical 模块与 Dashboard 之后装配（phase06 模块依赖 canonical 模块的表已建表）。
		// phase06 模块只承接只读聚合或独立写入，跨模块读依赖在 platform 装配点注入。

		// Onboarding：首轮状态读取（只读，读时派生 first_run_state）
		onboardingQuerySvc := buildOnboarding(pool)
		mountOnboarding(r, onboardingQuerySvc)

		// Export：导出快照读取 + 导出执行
		exportQuerySvc, exportCommandSvc := buildExport(pool)
		mountExport(r, exportQuerySvc, exportCommandSvc)

		// Backup：备份快照读取（read / verify 子路径）+ 备份执行
		backupQuerySvc, backupCommandSvc := buildBackup(pool)
		mountBackup(r, backupQuerySvc, backupCommandSvc)

		// ReuseSummary：复用感知派生读（只读，读时聚合）
		reuseSummaryQuerySvc := buildReuseSummary(pool)
		mountReuseSummary(r, reuseSummaryQuerySvc)
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
//
// phase03-14 修复：新增 Access-Control-Max-Age 头。
// 未设置该头时 Chrome 默认仅缓存 preflight 5 秒，过期后 POST 请求会被中止
// (net::ERR_ABORTED) 再重新 preflight，在控制台产生 error 日志。
// 设置 600 秒缓存可避免开发期间的频繁 re-preflight 与中止告警。
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 缓存 preflight 结果 10 分钟，避免 preflight 过期后 POST 被中止重发
		w.Header().Set("Access-Control-Max-Age", "600")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
