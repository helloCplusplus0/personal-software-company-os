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
	"strings"
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
// phase07-11 装配顺序（Connect handler 主线，compat 已退场）：
//  1. 应用基础中间件（RequestID / Logger / Recoverer / CORS）
//  2. 注册健康检查端点
//  3. 构造 Product Registry / Repository Binding 的 service 层
//  4. 通过 mount*Connect 把各业务模块的 canonical Connect handler 挂到 /api 下
//  5. 构造 http.Server
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
		// phase07-09: 先构造 Product Registry / Repository Binding 的 service 层，
		// 供 Module Registry Connect handler 的 L3/L4 兼容委派注入。
		//
		// 装配顺序约束（phase04-07 L162-181）：
		//   Repository Binding 必须先构造，将其 BindingStore 作为 BoundRepositoryReader
		//   注入到 Product Registry 的 QueryService（ProductRepositorySummaryRead owner=Repository Binding）。
		repoQuerySvc, repoCommandSvc, repoBindingStore := buildRepositoryBinding(pool)
		productQuerySvc, productCommandSvc := buildProductRegistry(pool, repoBindingStore)

		// --- canonical Connect handler 挂载（phase07-09 主线） ---
		mountModuleRegistryConnect(r, pool, productQuerySvc, productCommandSvc, repoQuerySvc, repoCommandSvc)
		mountDecisionCenterConnect(r, pool)
		mountProductRegistryConnect(r, productQuerySvc, productCommandSvc)
		mountRepositoryBindingConnect(r, repoQuerySvc, repoCommandSvc)

		// phase06 模块：Dashboard / Onboarding / Export / Backup / ReuseSummary
		dashboardQuerySvc := buildDashboard(pool)
		mountDashboardConnect(r, dashboardQuerySvc)

		onboardingQuerySvc := buildOnboarding(pool)
		mountOnboardingConnect(r, onboardingQuerySvc)

		exportQuerySvc, exportCommandSvc := buildExport(pool)
		mountExportConnect(r, exportQuerySvc, exportCommandSvc)

		backupQuerySvc, backupCommandSvc := buildBackup(pool)
		mountBackupConnect(r, backupQuerySvc, backupCommandSvc)

		reuseSummaryQuerySvc := buildReuseSummary(pool)
		mountReuseSummaryConnect(r, reuseSummaryQuerySvc)

		// phase08 review 模块：消费 Dashboard / Decision Center / Reuse Summary 既有 service
		dcQuerySvc, _ := buildDecisionCenter(pool)
		reviewQuerySvc, reviewCommandSvc := buildReview(pool, dashboardQuerySvc, dcQuerySvc, reuseSummaryQuerySvc)
		mountReviewConnect(r, reviewQuerySvc, reviewCommandSvc)

		// phase09 template reuse 模块：模板候选读取、模板预填、派生提示与模板来源复读
		templateReuseQuerySvc := buildTemplateReuse(pool, reuseSummaryQuerySvc)
		mountTemplateReuseConnect(r, templateReuseQuerySvc)

		// phase11 project context 模块：最小只读项目上下文聚合读取
		projectContextQuerySvc := buildProjectContext(pool)
		mountProjectContextConnect(r, projectContextQuerySvc)
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
		requestedHeaders := r.Header.Get("Access-Control-Request-Headers")
		if strings.TrimSpace(requestedHeaders) != "" {
			// 预检时优先回显浏览器实际请求的 header 列表，确保 Connect-Web
			// 等跨域调用不会因为额外协议头被浏览器拦截。
			w.Header().Set("Access-Control-Allow-Headers", requestedHeaders)
		} else {
			w.Header().Set(
				"Access-Control-Allow-Headers",
				"Content-Type, Authorization, Connect-Protocol-Version, Connect-Timeout-Ms, X-User-Agent, X-Grpc-Web, Grpc-Timeout",
			)
		}
		w.Header().Set(
			"Access-Control-Expose-Headers",
			"Grpc-Status, Grpc-Message, Grpc-Status-Details-Bin, Connect-Content-Encoding, Connect-Accept-Encoding",
		)
		// 缓存 preflight 结果 10 分钟，避免 preflight 过期后 POST 被中止重发
		w.Header().Set("Access-Control-Max-Age", "600")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
