// Package platform — Module Registry + Decision Center 路由装配。
//
// 本文件是各业务模块的"组合根"：把 handler/service/repository/candidate
// 四层装配到 chi 子路由上。放在 platform 包而非各业务模块根包，
// 是为了避免业务模块根包（持有 types/errors）与子包之间的导入循环。
//
// 路由设计对齐：
//   - phase02-09 spec §6 接口分组（Module Registry）
//   - phase03-10 §7.7 RPC→HTTP 映射矩阵（Decision Center）
//
// Module Registry 路由：
//   读组：
//     GET    /api/modules                                  ModuleListRead
//     GET    /api/modules/{moduleId}                       ModuleDetailRead
//   候选读取（phase02 临时承接）：
//     GET    /api/candidates/products                      ProductBindingCandidateRead
//     GET    /api/candidates/repositories                  RepositoryBindingCandidateRead
//   写组：
//     POST   /api/modules                                  ModuleCreateWrite
//     POST   /api/modules/{moduleId}/releases              ModuleReleaseWrite
//     POST   /api/modules/{moduleId}/bindings/products     ModuleBindingWrite (产品绑定)
//     POST   /api/modules/{moduleId}/bindings/repositories ModuleBindingWrite (仓库映射)
//
// Decision Center 路由（phase03-10 §7.7）：
//   读组：
//     GET    /api/decisions                                DecisionListRead
//     GET    /api/decisions/{decisionId}                   DecisionDetailRead
//     GET    /api/decisions/{decisionId}/candidates/modules DecisionModuleCandidateRead
//   写组：
//     POST   /api/decisions                                DecisionWrite (RecordDecision)
//     POST   /api/decisions/{decisionId}/links             DecisionLinkWrite (LinkDecisionToTarget)
package platform

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	dccandidate "github.com/psco/backend/internal/decisioncenter/candidate"
	dchandler "github.com/psco/backend/internal/decisioncenter/handler"
	dcrepository "github.com/psco/backend/internal/decisioncenter/repository"
	dcservice "github.com/psco/backend/internal/decisioncenter/service"
	"github.com/psco/backend/internal/moduleregistry/candidate"
	"github.com/psco/backend/internal/moduleregistry/handler"
	"github.com/psco/backend/internal/moduleregistry/repository"
	"github.com/psco/backend/internal/moduleregistry/service"
)

// mountModuleRegistry 把 Module Registry 模块的全部路由挂到 /api 下。
//
// 装配顺序：
//  1. 构造 repository 层（module / release / binding store）
//  2. 构造 candidate 层（product / repository candidate read）
//  3. 构造 service 层（query / command），注入 repository
//  4. 构造 handler 层（query / command），注入 service + candidate
//  5. 在 /api 子路由器上注册路径
func mountModuleRegistry(r chi.Router, pool *pgxpool.Pool) {
	// 1. repository 层
	moduleStore := repository.NewModuleStore(pool)
	releaseStore := repository.NewReleaseStore(pool)
	bindingStore := repository.NewBindingStore(pool)

	// 2. candidate 层
	productReader := candidate.NewProductCandidateRead(pool)
	repoReader := candidate.NewRepositoryCandidateRead(pool)

	// 3. service 层
	querySvc := service.NewQueryService(moduleStore, releaseStore, bindingStore)
	commandSvc := service.NewCommandService(moduleStore, releaseStore, bindingStore)

	// 4. handler 层
	queryH := handler.NewQueryHandler(querySvc, productReader, repoReader)
	commandH := handler.NewCommandHandler(commandSvc)

	// 5. 路由注册
	// --- 读组 + 候选读取 ---
	r.Get("/modules", queryH.ListModules)
	r.Get("/modules/{moduleId}", queryH.GetModuleDetail)
	r.Get("/candidates/products", queryH.ListProductCandidates)
	r.Get("/candidates/repositories", queryH.ListRepositoryCandidates)

	// --- 写组 ---
	r.Post("/modules", commandH.CreateModule)
	r.Post("/modules/{moduleId}/releases", commandH.CreateRelease)
	r.Post("/modules/{moduleId}/bindings/products", commandH.BindModuleToProduct)
	r.Post("/modules/{moduleId}/bindings/repositories", commandH.MapModuleToRepository)
}

// healthz 简单健康检查端点。
func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

// mountDecisionCenter 把 Decision Center 模块的全部路由挂到 /api 下。
//
// 装配顺序（phase03-10 §10.5 接线约束）：
//  1. 构造 repository 层（decision / link store）
//  2. 构造 candidate 层（module candidate read，由 Decision Center 拥有）
//  3. 构造 service 层（query / command），注入 repository + candidate
//  4. 构造 handler 层（query / command），注入 service
//  5. 在 /api 子路由器上注册路径（对齐 phase03-10 §7.7 RPC→HTTP 映射矩阵）
//
// candidate.ModuleCandidateRead 由本装配点构造并注入到 query/command service，
// service 层不自行构造，也不直接写跨模块 SQL（§10.5）。
func mountDecisionCenter(r chi.Router, pool *pgxpool.Pool) {
	// 1. repository 层
	decisionStore := dcrepository.NewDecisionStore(pool)
	linkStore := dcrepository.NewLinkStore(pool)

	// 2. candidate 层（跨模块 Module 候选读取，由 Decision Center 拥有）
	moduleCandidateRead := dccandidate.NewModuleCandidateRead(pool)

	// 3. service 层
	querySvc := dcservice.NewQueryService(decisionStore, linkStore, moduleCandidateRead)
	commandSvc := dcservice.NewCommandService(decisionStore, linkStore, moduleCandidateRead)

	// 4. handler 层
	queryH := dchandler.NewQueryHandler(querySvc)
	commandH := dchandler.NewCommandHandler(commandSvc)

	// 5. 路由注册（phase03-10 §7.7 RPC→HTTP 映射矩阵）
	// --- 读组 ---
	r.Get("/decisions", queryH.ListDecisions)
	r.Get("/decisions/{decisionId}", queryH.GetDecisionDetail)
	r.Get("/decisions/{decisionId}/candidates/modules", queryH.ListDecisionModuleCandidates)

	// --- 写组 ---
	r.Post("/decisions", commandH.CreateDecision)
	r.Post("/decisions/{decisionId}/links", commandH.LinkDecisionToTarget)
}
