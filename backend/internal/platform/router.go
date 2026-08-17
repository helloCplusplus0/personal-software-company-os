// Package platform — Module Registry + Decision Center + Product Registry + Repository Binding 路由装配。
//
// 本文件是各业务模块的"组合根"：把 handler/service/repository/candidate
// 四层装配到 chi 子路由上。放在 platform 包而非各业务模块根包，
// 是为了避免业务模块根包（持有 types/errors）与子包之间的导入循环。
//
// phase07-11 正式传输主线（compat 已全部退场）：
//   - canonical 业务接口统一通过 Connect handler（.proto + ConnectRPC）运行
//   - chi 只保留 /api shell、middleware 与非业务端点
//   - L1/L2 候选 compat 入口（GET /api/candidates/products, GET /api/candidates/repositories）已退场
//   - L3/L4 绑定 compat 入口（POST /api/modules/{moduleId}/bindings/*）已退场
//
// 路由设计对齐：
//   - phase07-07 formal spec §4 canonical Connect transport 装配规则
//   - phase07-03 compat 退场时点冻结
//   - proto/README.md RPC→HTTP 映射矩阵
package platform

import (
	"context"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	backupcandidate "github.com/psco/backend/internal/backup/candidate"
	backupconnect "github.com/psco/backend/internal/backup/connect"
	backuprepo "github.com/psco/backend/internal/backup/repository"
	backupservice "github.com/psco/backend/internal/backup/service"
	dashboardcandidate "github.com/psco/backend/internal/dashboard/candidate"
	dashboardconnect "github.com/psco/backend/internal/dashboard/connect"
	dashboardservice "github.com/psco/backend/internal/dashboard/service"
	dccandidate "github.com/psco/backend/internal/decisioncenter/candidate"
	dcconnect "github.com/psco/backend/internal/decisioncenter/connect"
	dcrepository "github.com/psco/backend/internal/decisioncenter/repository"
	dcservice "github.com/psco/backend/internal/decisioncenter/service"
	exportcandidate "github.com/psco/backend/internal/export/candidate"
	exportconnect "github.com/psco/backend/internal/export/connect"
	exportrepo "github.com/psco/backend/internal/export/repository"
	exportservice "github.com/psco/backend/internal/export/service"
	"github.com/psco/backend/internal/moduleregistry"
	mrconnect "github.com/psco/backend/internal/moduleregistry/connect"
	"github.com/psco/backend/internal/moduleregistry/repository"
	"github.com/psco/backend/internal/moduleregistry/service"
	onboardingcandidate "github.com/psco/backend/internal/onboarding/candidate"
	onboardingconnect "github.com/psco/backend/internal/onboarding/connect"
	onboardingrepo "github.com/psco/backend/internal/onboarding/repository"
	onboardingservice "github.com/psco/backend/internal/onboarding/service"
	"github.com/psco/backend/internal/productregistry"
	productcandidate "github.com/psco/backend/internal/productregistry/candidate"
	prconnect "github.com/psco/backend/internal/productregistry/connect"
	productrepo "github.com/psco/backend/internal/productregistry/repository"
	productservice "github.com/psco/backend/internal/productregistry/service"
	"github.com/psco/backend/internal/repositorybinding"
	repocandidate "github.com/psco/backend/internal/repositorybinding/candidate"
	rbconnect "github.com/psco/backend/internal/repositorybinding/connect"
	reporepo "github.com/psco/backend/internal/repositorybinding/repository"
	reposervice "github.com/psco/backend/internal/repositorybinding/service"
	reusesummarycandidate "github.com/psco/backend/internal/reusesummary/candidate"
	reusesummaryconnect "github.com/psco/backend/internal/reusesummary/connect"
	reusesummaryservice "github.com/psco/backend/internal/reusesummary/service"
	reviewconnect "github.com/psco/backend/internal/review/connect"
	reviewrepo "github.com/psco/backend/internal/review/repository"
	reviewservice "github.com/psco/backend/internal/review/service"
	projectcontextcandidate "github.com/psco/backend/internal/projectcontext/candidate"
	projectcontextconnect "github.com/psco/backend/internal/projectcontext/connect"
	projectcontextservice "github.com/psco/backend/internal/projectcontext/service"
	governanceprofilecandidate "github.com/psco/backend/internal/governanceprofile/candidate"
	governanceprofileconnect "github.com/psco/backend/internal/governanceprofile/connect"
	governanceprofilerepo "github.com/psco/backend/internal/governanceprofile/repository"
	governanceprofileservice "github.com/psco/backend/internal/governanceprofile/service"
	templatereusecandidate "github.com/psco/backend/internal/templatereuse/candidate"
	templatereuseconnect "github.com/psco/backend/internal/templatereuse/connect"
	templatereuseservice "github.com/psco/backend/internal/templatereuse/service"

	// generated Connect handler constructors
	backupv1connect "github.com/psco/backend/internal/gen/connect/psco/backup/v1/backupv1connect"
	dashboardv1connect "github.com/psco/backend/internal/gen/connect/psco/dashboard/v1/dashboardv1connect"
	decisioncenterv1connect "github.com/psco/backend/internal/gen/connect/psco/decision_center/v1/decision_centerv1connect"
	exportv1connect "github.com/psco/backend/internal/gen/connect/psco/export/v1/exportv1connect"
	moduleregistryv1connect "github.com/psco/backend/internal/gen/connect/psco/module_registry/v1/module_registryv1connect"
	onboardingv1connect "github.com/psco/backend/internal/gen/connect/psco/onboarding/v1/onboardingv1connect"
	productregistryv1connect "github.com/psco/backend/internal/gen/connect/psco/product_registry/v1/product_registryv1connect"
	repositorybindingv1connect "github.com/psco/backend/internal/gen/connect/psco/repository_binding/v1/repository_bindingv1connect"
	reusesummaryv1connect "github.com/psco/backend/internal/gen/connect/psco/reuse_summary/v1/reuse_summaryv1connect"
	reviewv1connect "github.com/psco/backend/internal/gen/connect/psco/review/v1/reviewv1connect"
	templatereusev1connect "github.com/psco/backend/internal/gen/connect/psco/template_reuse/v1/template_reusev1connect"
	projectcontextv1connect "github.com/psco/backend/internal/gen/connect/psco/project_context/v1/project_contextv1connect"
	governanceprofilev1connect "github.com/psco/backend/internal/gen/connect/psco/governance_profile/v1/governance_profilev1connect"
)

// ============================================================================
// 模块构造器（build* 函数保留不变，只构造 service 层）
// ============================================================================

// buildProductRegistry 构造 Product Registry 的 service 层并返回。
//
// 返回 query / command service 供 mountProductRegistryConnect 注册路由，
// 以及供 mountModuleRegistryConnect 的旧绑定入口做兼容委派注入。
//
// 装配依赖（phase04-07 L162-181 冻结）：
//   - boundRepoReader 必须由 Repository Binding 模块的 BindingStore 注入（owner=Repository Binding）
//   - 调用方必须先构造 Repository Binding 的 BindingStore，再传入本函数
func buildProductRegistry(pool *pgxpool.Pool, boundRepoReader productregistry.BoundRepositoryReader) (*productservice.QueryService, *productservice.CommandService) {
	productStore := productrepo.NewProductStore(pool)
	bindingStore := productrepo.NewBindingStore(pool)
	moduleCandidateRead := productcandidate.NewModuleCandidateRead(pool)
	querySvc := productservice.NewQueryService(productStore, bindingStore, boundRepoReader, moduleCandidateRead)
	commandSvc := productservice.NewCommandService(productStore, moduleCandidateRead)
	return querySvc, commandSvc
}

// buildRepositoryBinding 构造 Repository Binding 的 service 层并返回。
//
// 返回 query / command service 供 mountRepositoryBindingConnect 注册路由，
// 以及供 mountModuleRegistryConnect 的旧绑定入口做兼容委派注入。
//
// 同时返回 bindingStore（*reporepo.BindingStore），用于注入到 Product Registry 的
// QueryService 作为 BoundRepositoryReader 实现（phase04-07 L180 装配点接线要求）。
func buildRepositoryBinding(pool *pgxpool.Pool) (*reposervice.QueryService, *reposervice.CommandService, *reporepo.BindingStore) {
	repositoryStore := reporepo.NewRepositoryStore(pool)
	bindingStore := reporepo.NewBindingStore(pool)
	productCandidateRead := repocandidate.NewProductCandidateRead(pool)
	moduleCandidateRead := repocandidate.NewModuleCandidateRead(pool)
	querySvc := reposervice.NewQueryService(repositoryStore, bindingStore, productCandidateRead, moduleCandidateRead)
	commandSvc := reposervice.NewCommandService(repositoryStore, bindingStore, productCandidateRead, moduleCandidateRead)
	return querySvc, commandSvc, bindingStore
}

// ============================================================================
// Connect handler 装配函数（phase07-09 新增）
// ============================================================================

// mountModuleRegistryConnect 把 Module Registry 的 canonical Connect handler 挂到 /api 下。
//
// phase07-09 切换：
//   - canonical 读/写路由（ListModules / GetModuleDetail / CreateModule / CreateRelease）
//     统一切换到 Connect handler
//   - BindModuleToProduct / MapModuleToRepository 通过 delegate 委派到 canonical owner
//   - ListProductCandidates / ListRepositoryCandidates（L1/L2 compat）通过 delegate 委派
//   - 旧 hand-written JSON handler 不再作为 canonical 入口
func mountModuleRegistryConnect(
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

	// 2. service 层
	querySvc := service.NewQueryService(moduleStore, releaseStore, bindingStore)
	commandSvc := service.NewCommandService(moduleStore, releaseStore)

	// 3. L1/L2 compat delegate（候选读取委派到 canonical owner）
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

	// L3/L4 compat delegate（绑定/映射委派到 canonical owner）
	bindModuleToProduct := func(ctx context.Context, moduleID, productID string) error {
		return productBindingSvc.BindModuleToProduct(ctx, productID, moduleID)
	}
	mapModuleToRepository := func(ctx context.Context, moduleID, repositoryID string) error {
		return repositoryMappingSvc.MapModuleToRepository(ctx, repositoryID, moduleID)
	}

	// 4. Connect handler 构造
	connectSvc := mrconnect.NewServer(
		querySvc, commandSvc,
		legacyProductCandidates, legacyRepositoryCandidates,
		bindModuleToProduct, mapModuleToRepository,
	)
	path, handler := moduleregistryv1connect.NewModuleRegistryServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountDecisionCenterConnect 把 Decision Center 的 canonical Connect handler 挂到 /api 下。
func mountDecisionCenterConnect(r chi.Router, pool *pgxpool.Pool) {
	querySvc, commandSvc := buildDecisionCenter(pool)
	connectSvc := dcconnect.NewServer(querySvc, commandSvc)
	path, handler := decisioncenterv1connect.NewDecisionCenterServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountProductRegistryConnect 把 Product Registry 的 canonical Connect handler 挂到 /api 下。
func mountProductRegistryConnect(r chi.Router, querySvc *productservice.QueryService, commandSvc *productservice.CommandService) {
	connectSvc := prconnect.NewServer(querySvc, commandSvc)
	path, handler := productregistryv1connect.NewProductRegistryServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountRepositoryBindingConnect 把 Repository Binding 的 canonical Connect handler 挂到 /api 下。
func mountRepositoryBindingConnect(r chi.Router, querySvc *reposervice.QueryService, commandSvc *reposervice.CommandService) {
	connectSvc := rbconnect.NewServer(querySvc, commandSvc)
	path, handler := repositorybindingv1connect.NewRepositoryBindingServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountDashboardConnect 把 Dashboard 的 canonical Connect handler 挂到 /api 下。
func mountDashboardConnect(r chi.Router, querySvc *dashboardservice.QueryService) {
	connectSvc := dashboardconnect.NewServer(querySvc)
	path, handler := dashboardv1connect.NewDashboardServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountOnboardingConnect 把 Onboarding 的 canonical Connect handler 挂到 /api 下。
func mountOnboardingConnect(r chi.Router, querySvc *onboardingservice.QueryService) {
	connectSvc := onboardingconnect.NewServer(querySvc)
	path, handler := onboardingv1connect.NewOnboardingServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountExportConnect 把 Export 的 canonical Connect handler 挂到 /api 下。
func mountExportConnect(r chi.Router, querySvc *exportservice.QueryService, commandSvc *exportservice.CommandService) {
	connectSvc := exportconnect.NewServer(querySvc, commandSvc)
	path, handler := exportv1connect.NewExportServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountBackupConnect 把 Backup 的 canonical Connect handler 挂到 /api 下。
func mountBackupConnect(r chi.Router, querySvc *backupservice.QueryService, commandSvc *backupservice.CommandService) {
	connectSvc := backupconnect.NewServer(querySvc, commandSvc)
	path, handler := backupv1connect.NewBackupServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountReuseSummaryConnect 把 ReuseSummary 的 canonical Connect handler 挂到 /api 下。
func mountReuseSummaryConnect(r chi.Router, querySvc *reusesummaryservice.QueryService) {
	connectSvc := reusesummaryconnect.NewServer(querySvc)
	path, handler := reusesummaryv1connect.NewReuseSummaryServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountReviewConnect 把 Review 的 canonical Connect handler 挂到 /api 下。
func mountReviewConnect(r chi.Router, querySvc *reviewservice.QueryService, commandSvc *reviewservice.CommandService) {
	connectSvc := reviewconnect.NewServer(querySvc, commandSvc)
	path, handler := reviewv1connect.NewReviewServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountTemplateReuseConnect 把 Template Reuse 的 canonical Connect handler 挂到 /api 下。
//
// phase09-08 新增：
//   - 模板候选读取、模板预填、派生提示与模板来源复读四类只读能力通过本 handler 承接
//   - TemplateReuseService 是模板读能力的唯一 canonical transport owner
func mountTemplateReuseConnect(r chi.Router, querySvc *templatereuseservice.QueryService) {
	connectSvc := templatereuseconnect.NewServer(querySvc)
	path, handler := templatereusev1connect.NewTemplateReuseServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountProjectContextConnect 把 Project Context 的 canonical Connect handler 挂到 /api 下。
//
// phase11-07 新增：
//   - 最小只读项目上下文聚合读取能力
//   - 以 repository_id 为唯一结构化输入锚点
//   - ProjectContextService 是项目上下文读取的唯一 canonical transport owner
func mountProjectContextConnect(r chi.Router, querySvc *projectcontextservice.QueryService) {
	connectSvc := projectcontextconnect.NewServer(querySvc)
	path, handler := projectcontextv1connect.NewProjectContextServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// mountGovernanceProfileConnect 把 Governance Profile 的 canonical Connect handler 挂到 /api 下。
//
// phase13-08 新增：
//   - 项目治理画像结构化写读能力（Get / Update）
//   - 以 repository_id 为唯一结构化输入锚点
//   - GovernanceProfileService 是治理画像写读的唯一 canonical transport owner
func mountGovernanceProfileConnect(r chi.Router, querySvc *governanceprofileservice.QueryService, commandSvc *governanceprofileservice.CommandService) {
	connectSvc := governanceprofileconnect.NewServer(querySvc, commandSvc)
	path, handler := governanceprofilev1connect.NewGovernanceProfileServiceHandler(connectSvc)
	r.Handle(path+"*", http.StripPrefix("/api", handler))
}

// ============================================================================
// phase06 模块构造器
// ============================================================================

// buildDecisionCenter 构造 Decision Center 的 service 层并返回。
// 供 mountDecisionCenterConnect 与 buildReview 共用，避免重复构造。
func buildDecisionCenter(pool *pgxpool.Pool) (*dcservice.QueryService, *dcservice.CommandService) {
	decisionStore := dcrepository.NewDecisionStore(pool)
	linkStore := dcrepository.NewLinkStore(pool)
	moduleCandidateRead := dccandidate.NewModuleCandidateRead(pool)
	querySvc := dcservice.NewQueryService(decisionStore, linkStore, moduleCandidateRead)
	commandSvc := dcservice.NewCommandService(decisionStore, linkStore, moduleCandidateRead)
	return querySvc, commandSvc
}

// buildDashboard 构造 Dashboard 的 QueryService 并返回。
func buildDashboard(pool *pgxpool.Pool) *dashboardservice.QueryService {
	overviewReaders := dashboardcandidate.NewOverviewReaders(pool)
	feedbackReaders := dashboardcandidate.NewFeedbackReaders(pool)
	activityReaders := dashboardcandidate.NewActivityReaders(pool)
	return dashboardservice.NewQueryService(overviewReaders, feedbackReaders, activityReaders)
}

// buildOnboarding 构造 Onboarding 的 QueryService 并返回。
//
// phase10-08 新增 chainStateReaders 与 recoveryStore 注入。
func buildOnboarding(pool *pgxpool.Pool) *onboardingservice.QueryService {
	firstRunReaders := onboardingcandidate.NewFirstRunReaders(pool)
	chainStateReaders := onboardingcandidate.NewChainStateReaders(pool)
	recoveryStore := onboardingrepo.NewRecoveryStore(pool)
	if err := recoveryStore.EnsureSchema(context.Background()); err != nil {
		panic(err)
	}
	return onboardingservice.NewQueryService(firstRunReaders, chainStateReaders, recoveryStore)
}

// buildExport 构造 Export 的 QueryService / CommandService 并返回。
func buildExport(pool *pgxpool.Pool) (*exportservice.QueryService, *exportservice.CommandService) {
	exportStore := exportrepo.NewExportStore(pool)
	assetReader := exportcandidate.NewAssetReader(pool)
	querySvc := exportservice.NewQueryService(exportStore, assetReader)
	commandSvc := exportservice.NewCommandService(exportStore, assetReader)
	return querySvc, commandSvc
}

// buildBackup 构造 Backup 的 QueryService / CommandService 并返回。
func buildBackup(pool *pgxpool.Pool) (*backupservice.QueryService, *backupservice.CommandService) {
	backupStore := backuprepo.NewBackupStore(pool)
	assetReader := backupcandidate.NewAssetReader(pool)
	querySvc := backupservice.NewQueryService(backupStore, assetReader)
	commandSvc := backupservice.NewCommandService(backupStore, assetReader)
	return querySvc, commandSvc
}

// buildReuseSummary 构造 ReuseSummary 的 QueryService 并返回。
func buildReuseSummary(pool *pgxpool.Pool) *reusesummaryservice.QueryService {
	reuseReaders := reusesummarycandidate.NewReuseReaders(pool)
	return reusesummaryservice.NewQueryService(reuseReaders)
}

// buildReview 构造 Review 的 QueryService 和 CommandService 并返回。
// 依赖 dashboard、decisioncenter、reusesummary 的既有 QueryService。
func buildReview(pool *pgxpool.Pool, dashboardQuerySvc *dashboardservice.QueryService, decisionCenterQuerySvc *dcservice.QueryService, reuseSummaryQuerySvc *reusesummaryservice.QueryService) (*reviewservice.QueryService, *reviewservice.CommandService) {
	reviewRecordStore := reviewrepo.NewReviewRecordStore(pool)
	querySvc := reviewservice.NewQueryService(dashboardQuerySvc, decisionCenterQuerySvc, reuseSummaryQuerySvc)
	commandSvc := reviewservice.NewCommandService(reviewRecordStore)
	return querySvc, commandSvc
}

// buildTemplateReuse 构造 Template Reuse 的 QueryService 并返回。
//
// phase09-08 新增：
//   - 模板候选、模板预填、派生提示与模板来源复读四类只读能力
//   - 数据从 product_modules 已持久化事实读时派生，不新增快照表
func buildTemplateReuse(pool *pgxpool.Pool, reuseSummaryQuerySvc *reusesummaryservice.QueryService) *templatereuseservice.QueryService {
	candidateReaders := templatereusecandidate.NewTemplateCandidateReaders(pool)
	return templatereuseservice.NewQueryService(candidateReaders, reuseSummaryQuerySvc)
}

// buildProjectContext 构造 Project Context 的 QueryService 并返回。
//
// phase11-07 新增：
//   - 最小只读项目上下文聚合读取能力
//   - 以 repository_id 为唯一结构化输入锚点
//   - 不依赖消费侧目录结构或固定文件名
func buildProjectContext(pool *pgxpool.Pool) *projectcontextservice.QueryService {
	contextReaders := projectcontextcandidate.NewContextReaders(pool)
	return projectcontextservice.NewQueryService(contextReaders)
}

// buildGovernanceProfile 构造 Governance Profile 的 QueryService / CommandService 并返回。
//
// phase13-08 新增：
//   - 项目治理画像结构化写读主线（repository_id 唯一锚点）
//   - repository 存在性前提经 candidate 子包隔离（service 不直接写跨模块 SQL）
//   - 主记录与两组 bindings 的持久化收敛在 ProfileStore 单一事务边界
func buildGovernanceProfile(pool *pgxpool.Pool) (*governanceprofileservice.QueryService, *governanceprofileservice.CommandService) {
	repositoryReader := governanceprofilecandidate.NewRepositoryReader(pool)
	profileStore := governanceprofilerepo.NewProfileStore(pool)
	querySvc := governanceprofileservice.NewQueryService(repositoryReader, profileStore)
	commandSvc := governanceprofileservice.NewCommandService(repositoryReader, profileStore)
	return querySvc, commandSvc
}

// ============================================================================
// 非业务端点
// ============================================================================

// healthz 简单健康检查端点。
func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
