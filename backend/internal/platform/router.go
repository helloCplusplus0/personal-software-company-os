// Package platform — Module Registry + Decision Center + Product Registry + Repository Binding 路由装配。
//
// 本文件是各业务模块的"组合根"：把 handler/service/repository/candidate
// 四层装配到 chi 子路由上。放在 platform 包而非各业务模块根包，
// 是为了避免业务模块根包（持有 types/errors）与子包之间的导入循环。
//
// 路由设计对齐：
//   - phase02-09 spec §6 接口分组（Module Registry）
//   - phase03-10 §7.7 RPC→HTTP 映射矩阵（Decision Center）
//   - phase04-10 §RPC→HTTP 映射矩阵 + proto/README.md（Product Registry / Repository Binding）
//
// Module Registry 路由：
//
//	读组：
//	  GET    /api/modules                                  ModuleListRead
//	  GET    /api/modules/{moduleId}                       ModuleDetailRead
//	写组：
//	  POST   /api/modules                                  ModuleCreateWrite
//	  POST   /api/modules/{moduleId}/releases              ModuleReleaseWrite
//	  POST   /api/modules/{moduleId}/bindings/products     ModuleBindingWrite (兼容委派到 Product Registry)
//	  POST   /api/modules/{moduleId}/bindings/repositories ModuleBindingWrite (兼容委派到 Repository Binding)
//
//	phase04-12 起，canonical 候选读取由 Product Registry / Repository Binding 各自承接；
//	/api/candidates/products 与 /api/candidates/repositories 仅保留给 Module Detail 历史入口做兼容适配。
//
// Decision Center 路由（phase03-10 §7.7）：
//
//	读组：
//	  GET    /api/decisions                                DecisionListRead
//	  GET    /api/decisions/{decisionId}                   DecisionDetailRead
//	  GET    /api/decisions/{decisionId}/candidates/modules DecisionModuleCandidateRead
//	写组：
//	  POST   /api/decisions                                DecisionWrite (RecordDecision)
//	  POST   /api/decisions/{decisionId}/links             DecisionLinkWrite (LinkDecisionToTarget)
//
// Product Registry 路由（phase04-10 §RPC→HTTP 映射矩阵）：
//
//	读组：
//	  GET    /api/products                                 ProductListRead
//	  GET    /api/products/{productId}                     ProductDetailRead
//	  GET    /api/products/{productId}/candidates/modules  ProductModuleCandidateRead
//	写组：
//	  POST   /api/products                                 ProductCreateWrite
//	  POST   /api/products/{productId}/bindings/modules    ProductModuleBindingWrite
//
// Repository Binding 路由（phase04-10 §RPC→HTTP 映射矩阵）：
//
//	读组：
//	  GET    /api/repositories                             RepositoryListRead
//	  GET    /api/repositories/{repositoryId}              RepositoryDetailRead
//	  GET    /api/repositories/{repositoryId}/candidates/products  RepositoryProductCandidateRead
//	  GET    /api/repositories/{repositoryId}/candidates/modules   RepositoryModuleCandidateRead
//	写组：
//	  POST   /api/repositories                             RepositoryCreateWrite
//	  POST   /api/repositories/{repositoryId}/bindings/products    RepositoryProductBindingWrite
//	  POST   /api/repositories/{repositoryId}/bindings/modules     RepositoryModuleMappingWrite
package platform

import (
	"context"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	dccandidate "github.com/psco/backend/internal/decisioncenter/candidate"
	dchandler "github.com/psco/backend/internal/decisioncenter/handler"
	dcrepository "github.com/psco/backend/internal/decisioncenter/repository"
	dcservice "github.com/psco/backend/internal/decisioncenter/service"
	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/handler"
	"github.com/psco/backend/internal/moduleregistry/repository"
	"github.com/psco/backend/internal/moduleregistry/service"
	"github.com/psco/backend/internal/productregistry"
	productcandidate "github.com/psco/backend/internal/productregistry/candidate"
	producthandler "github.com/psco/backend/internal/productregistry/handler"
	productrepo "github.com/psco/backend/internal/productregistry/repository"
	productservice "github.com/psco/backend/internal/productregistry/service"
	"github.com/psco/backend/internal/repositorybinding"
	repocandidate "github.com/psco/backend/internal/repositorybinding/candidate"
	repohandler "github.com/psco/backend/internal/repositorybinding/handler"
	reporepo "github.com/psco/backend/internal/repositorybinding/repository"
	reposervice "github.com/psco/backend/internal/repositorybinding/service"
)

// buildProductRegistry 构造 Product Registry 的 service 层并返回。
//
// 返回 query / command service 供 mountProductRegistry 注册路由，
// 以及供 mountModuleRegistry 的旧绑定入口做兼容委派注入。
//
// 装配依赖（phase04-07 L162-181 冻结）：
//   - boundRepoReader 必须由 Repository Binding 模块的 BindingStore 注入（owner=Repository Binding）
//   - 调用方必须先构造 Repository Binding 的 BindingStore，再传入本函数
func buildProductRegistry(pool *pgxpool.Pool, boundRepoReader productregistry.BoundRepositoryReader) (*productservice.QueryService, *productservice.CommandService) {
	// 1. repository 层
	productStore := productrepo.NewProductStore(pool)
	bindingStore := productrepo.NewBindingStore(pool)

	// 2. candidate 层（跨模块 Module 候选读取，由 Product Registry 拥有）
	moduleCandidateRead := productcandidate.NewModuleCandidateRead(pool)

	// 3. service 层（注入 boundRepoReader：ProductRepositorySummaryRead 的消费侧接口）
	querySvc := productservice.NewQueryService(productStore, bindingStore, boundRepoReader, moduleCandidateRead)
	commandSvc := productservice.NewCommandService(productStore, moduleCandidateRead)

	return querySvc, commandSvc
}

// buildRepositoryBinding 构造 Repository Binding 的 service 层并返回。
//
// 返回 query / command service 供 mountRepositoryBinding 注册路由，
// 以及供 mountModuleRegistry 的旧绑定入口做兼容委派注入。
//
// 同时返回 bindingStore（*reporepo.BindingStore），用于注入到 Product Registry 的
// QueryService 作为 BoundRepositoryReader 实现（phase04-07 L180 装配点接线要求）。
func buildRepositoryBinding(pool *pgxpool.Pool) (*reposervice.QueryService, *reposervice.CommandService, *reporepo.BindingStore) {
	// 1. repository 层
	repositoryStore := reporepo.NewRepositoryStore(pool)
	bindingStore := reporepo.NewBindingStore(pool)

	// 2. candidate 层（跨模块 Product / Module 候选读取，由 Repository Binding 拥有）
	productCandidateRead := repocandidate.NewProductCandidateRead(pool)
	moduleCandidateRead := repocandidate.NewModuleCandidateRead(pool)

	// 3. service 层
	querySvc := reposervice.NewQueryService(repositoryStore, bindingStore, productCandidateRead, moduleCandidateRead)
	commandSvc := reposervice.NewCommandService(repositoryStore, bindingStore, productCandidateRead, moduleCandidateRead)

	return querySvc, commandSvc, bindingStore
}

// mountModuleRegistry 把 Module Registry 模块的全部路由挂到 /api 下。
//
// 装配顺序：
//  1. 构造 repository 层（module / release / binding store）
//  2. 构造 service 层（query / command），注入 repository
//  3. 构造 handler 层（query / command），注入 service + 跨模块兼容委派目标
//  4. 在 /api 子路由器上注册路径
//
// phase04-12 起：
//   - 旧候选读取入口（/api/candidates/products, /api/candidates/repositories）仅保留 transport 兼容层，
//     实际数据委派给 Product Registry / Repository Binding canonical 读组
//   - 旧绑定入口（/api/modules/{moduleId}/bindings/*）保留为兼容委派，
//     productBindingSvc / repositoryMappingSvc 由调用方注入
func mountModuleRegistry(
	r chi.Router,
	pool *pgxpool.Pool,
	productQuerySvc *productservice.QueryService,
	productBindingSvc *productservice.CommandService,
	repositoryQuerySvc *reposervice.QueryService,
	repositoryMappingSvc *reposervice.CommandService,
) {
	// 1. repository 层
	moduleStore := repository.NewModuleStore(pool)
	releaseStore := repository.NewReleaseStore(pool)
	bindingStore := repository.NewBindingStore(pool)

	// 2. service 层（phase04-12 起 command service 不再接收 bindingStore）
	querySvc := service.NewQueryService(moduleStore, releaseStore, bindingStore)
	commandSvc := service.NewCommandService(moduleStore, releaseStore)

	// 3. handler 层（注入跨模块兼容委派目标）
	legacyProductCandidates := func(ctx context.Context) ([]moduleregistry.ProductCandidate, error) {
		items, err := productQuerySvc.ListProducts(ctx, productregistry.ListProductsQuery{
			StatusFilter: productregistry.ProductStatusActive,
		})
		if err != nil {
			return nil, err
		}

		candidates := make([]moduleregistry.ProductCandidate, 0, len(items))
		for _, item := range items {
			candidates = append(candidates, moduleregistry.ProductCandidate{
				ID:   item.ID,
				Name: item.Name,
			})
		}
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].Name < candidates[j].Name
		})
		return candidates, nil
	}
	legacyRepositoryCandidates := func(ctx context.Context) ([]moduleregistry.RepositoryCandidate, error) {
		items, err := repositoryQuerySvc.ListRepositories(ctx, repositorybinding.ListRepositoriesQuery{
			StatusFilter: repositorybinding.RepositoryStatusActive,
		})
		if err != nil {
			return nil, err
		}

		candidates := make([]moduleregistry.RepositoryCandidate, 0, len(items))
		for _, item := range items {
			candidates = append(candidates, moduleregistry.RepositoryCandidate{
				ID:   item.ID,
				Name: item.Name,
			})
		}
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].Name < candidates[j].Name
		})
		return candidates, nil
	}
	queryH := handler.NewQueryHandler(querySvc, legacyProductCandidates, legacyRepositoryCandidates)
	commandH := handler.NewCommandHandler(commandSvc, productBindingSvc, repositoryMappingSvc)

	// 4. 路由注册
	// --- 读组 ---
	r.Get("/modules", queryH.ListModules)
	r.Get("/modules/{moduleId}", queryH.GetModuleDetail)
	r.Get("/candidates/products", queryH.ListProductCandidates)
	r.Get("/candidates/repositories", queryH.ListRepositoryCandidates)

	// --- 写组 ---
	r.Post("/modules", commandH.CreateModule)
	r.Post("/modules/{moduleId}/releases", commandH.CreateRelease)
	// 兼容委派：旧模块中心绑定入口，委派到 Product Registry / Repository Binding
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

// mountProductRegistry 把 Product Registry 模块的全部路由挂到 /api 下。
//
// 装配顺序（phase04-10 §RPC→HTTP 映射矩阵 + phase04-07 分层语义）：
//  1. 由调用方注入已构造的 query / command service（避免重复构造，支持跨模块委派复用）
//  2. 构造 handler 层（query / command），注入 service
//  3. 在 /api 子路由器上注册路径
func mountProductRegistry(r chi.Router, querySvc *productservice.QueryService, commandSvc *productservice.CommandService) {
	// handler 层
	queryH := producthandler.NewQueryHandler(querySvc)
	commandH := producthandler.NewCommandHandler(commandSvc)

	// 路由注册（phase04-10 §RPC→HTTP 映射矩阵）
	// --- 读组 ---
	r.Get("/products", queryH.ListProducts)
	r.Get("/products/{productId}", queryH.GetProductDetail)
	r.Get("/products/{productId}/candidates/modules", queryH.ListProductModuleCandidates)

	// --- 写组 ---
	r.Post("/products", commandH.CreateProduct)
	r.Post("/products/{productId}/bindings/modules", commandH.BindModuleToProduct)
}

// mountRepositoryBinding 把 Repository Binding 模块的全部路由挂到 /api 下。
//
// 装配顺序（phase04-10 §RPC→HTTP 映射矩阵 + phase04-07 分层语义）：
//  1. 由调用方注入已构造的 query / command service（避免重复构造，支持跨模块委派复用）
//  2. 构造 handler 层（query / command），注入 service
//  3. 在 /api 子路由器上注册路径
func mountRepositoryBinding(r chi.Router, querySvc *reposervice.QueryService, commandSvc *reposervice.CommandService) {
	// handler 层
	queryH := repohandler.NewQueryHandler(querySvc)
	commandH := repohandler.NewCommandHandler(commandSvc)

	// 路由注册（phase04-10 §RPC→HTTP 映射矩阵）
	// --- 读组 ---
	r.Get("/repositories", queryH.ListRepositories)
	r.Get("/repositories/{repositoryId}", queryH.GetRepositoryDetail)
	r.Get("/repositories/{repositoryId}/candidates/products", queryH.ListRepositoryProductCandidates)
	r.Get("/repositories/{repositoryId}/candidates/modules", queryH.ListRepositoryModuleCandidates)

	// --- 写组 ---
	r.Post("/repositories", commandH.CreateRepository)
	r.Post("/repositories/{repositoryId}/bindings/products", commandH.BindRepositoryToProduct)
	r.Post("/repositories/{repositoryId}/bindings/modules", commandH.MapModuleToRepository)
}
