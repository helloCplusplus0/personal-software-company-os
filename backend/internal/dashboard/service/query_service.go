// Package service — Dashboard 三类聚合读编排层。
//
// 单一 QueryService 承接 ReadOverview / ReadFeedbackSignal / ReadRecentActivity 三个读组，
// 对齐 phase05-07 已冻结的 query service owner 结论与 phase05-12 spec 的编排语义。
//
// 错误语义（对齐 phase05-04）：
//   - ReadOverview：任一计数 reader 失败 → 整页失败（返回 error）
//   - ReadFeedbackSignal：任一信号 reader 失败 → 局部失败（返回 error，不拖垮 ReadOverview）
//   - ReadRecentActivity：任一活动 reader 失败 → 局部失败（返回 error，不拖垮 ReadOverview）
//
// 跨模块读取全部通过 candidate/ 子包隔离，service 层不直接写跨模块 SQL（phase05-07）。
//
// 文件落点：backend/internal/dashboard/service/query_service.go
package service

import (
	"context"
	"fmt"
	"sort"

	"github.com/psco/backend/internal/dashboard"
	"github.com/psco/backend/internal/dashboard/candidate"
)

// 当前版本各读组的展示上限（对齐 phase05-02 / phase05-04 冻结结论，
// 上限由 service 层承接，不进入 .proto 合同本体）。
const (
	maxCurrentFocusSignals   = 5  // current_focus_signals 最多展示 5 条
	maxRepresentativeSignals = 3  // representative_signals 最多展示 3 条代表性缺口项
	maxRecentActivities      = 10 // activities 最多返回 10 条
)

// QueryService 承接 Dashboard 三类聚合读编排。
//
// 依赖通过 platform 装配点注入（phase05-07 §"platform 装配层必须接线 Dashboard 模块"）：
//   - overviewReaders：四个 canonical 模块的计数 reader
//   - feedbackReaders：pending decision 与 product asset coverage reader
//   - activityReaders：四个 canonical 模块的活动 reader
//
// service 层不自行构造 candidate reader，也不直接跨模块写 SQL。
type QueryService struct {
	overviewReaders *candidate.OverviewReaders
	feedbackReaders *candidate.FeedbackReaders
	activityReaders *candidate.ActivityReaders
}

// NewQueryService 构造 QueryService。
//
// 三个 candidate reader 集合由 platform 装配点构造并注入。
func NewQueryService(
	overviewReaders *candidate.OverviewReaders,
	feedbackReaders *candidate.FeedbackReaders,
	activityReaders *candidate.ActivityReaders,
) *QueryService {
	return &QueryService{
		overviewReaders: overviewReaders,
		feedbackReaders: feedbackReaders,
		activityReaders: activityReaders,
	}
}

// ============================================================================
// ReadOverview — 概览聚合读取
// ============================================================================

// ReadOverview 编排四个 canonical 模块的计数 reader，返回 DashboardOverview。
//
// 整页失败语义（phase05-04）：任一 reader 失败时返回整页失败（error），
// 不吞掉局部失败后返回部分计数。
func (s *QueryService) ReadOverview(ctx context.Context) (*dashboard.DashboardOverview, error) {
	moduleCount, err := s.overviewReaders.CountModules(ctx)
	if err != nil {
		return nil, dashboard.ErrOverviewReadFailed
	}

	productCount, productWithModule, productWithRepository, err := s.overviewReaders.CountProducts(ctx)
	if err != nil {
		return nil, dashboard.ErrOverviewReadFailed
	}

	repositoryCount, err := s.overviewReaders.CountRepositories(ctx)
	if err != nil {
		return nil, dashboard.ErrOverviewReadFailed
	}

	decisionCount, err := s.overviewReaders.CountDecisions(ctx)
	if err != nil {
		return nil, dashboard.ErrOverviewReadFailed
	}

	return &dashboard.DashboardOverview{
		ModuleCount:                moduleCount,
		ProductCount:               productCount,
		RepositoryCount:            repositoryCount,
		DecisionCount:              decisionCount,
		ProductWithRepositoryCount: productWithRepository,
		ProductWithModuleCount:     productWithModule,
	}, nil
}

// ============================================================================
// ReadFeedbackSignal — 反馈信号读取（主队列 + 资产缺口摘要）
// ============================================================================

// ReadFeedbackSignal 编排 PendingDecisionSignalReader 与 ProductAssetCoverageReader，
// 在 Dashboard 模块内归一化为统一 FeedbackSignal 列表与 ProductAssetCoverageSummary。
//
// 局部失败语义（phase05-04）：任一 reader 失败时返回局部失败（error），
// 不拖垮 ReadOverview 的整页成功语义。
//
// 归一化规则（phase05-12 spec）：
//   - pending decision 已绑定 decision_link → 单项信号（DECISION_DETAIL）
//   - pending decision 未绑定 decision_link → 聚合信号（DECISION_LIST）
//   - product 同时缺 module 与 repository → P2 信号
//   - product 缺 repository → P3 信号
//   - product 缺 module → P4 信号
//   - 主队列按 P1 > P2 > P3 > P4 排序，同优先级按 created_at DESC
//   - current_focus_signals 最多 5 条，representative_signals 最多 3 条
//
// 空态语义：无 pending decisions 且无资产缺口时返回成功语义，
// current_focus_signals 为空列表，ProductAssetCoverageSummary 返回完整结构与零计数。
func (s *QueryService) ReadFeedbackSignal(ctx context.Context) (*dashboard.FeedbackSignalReadResult, error) {
	pendingDecisions, err := s.feedbackReaders.ReadPendingDecisions(ctx)
	if err != nil {
		return nil, dashboard.ErrFeedbackSignalReadFailed
	}

	coverage, err := s.feedbackReaders.ReadProductAssetCoverage(ctx)
	if err != nil {
		return nil, dashboard.ErrFeedbackSignalReadFailed
	}

	// 归一化为统一 FeedbackSignal 列表
	signals := s.normalizeFeedbackSignals(pendingDecisions, coverage)

	// 排序：优先级升序（P1 > P2 > P3 > P4），同优先级按 created_at DESC
	sortFeedbackSignals(signals)

	// current_focus_signals 最多 5 条
	currentFocus := truncateSignals(signals, maxCurrentFocusSignals)

	// representative_signals：从 product asset coverage 缺口中取最多 3 条
	representative := s.buildRepresentativeSignals(coverage)

	return &dashboard.FeedbackSignalReadResult{
		CurrentFocusSignals: currentFocus,
		AssetFeedbackSummary: dashboard.ProductAssetCoverageSummary{
			FullyBoundProductCount:        coverage.FullyBoundCount,
			MissingBothBindingsCount:      coverage.MissingBothCount,
			MissingRepositoryBindingCount: coverage.MissingRepositoryCount,
			MissingModuleBindingCount:     coverage.MissingModuleCount,
			RepresentativeSignals:         representative,
		},
	}, nil
}

// normalizeFeedbackSignals 将 pending decisions 与 product asset coverage 归一化为统一 FeedbackSignal 列表。
func (s *QueryService) normalizeFeedbackSignals(
	pendingDecisions []candidate.PendingDecisionData,
	coverage *candidate.ProductAssetCoverageData,
) []dashboard.FeedbackSignal {
	signals := make([]dashboard.FeedbackSignal, 0)

	// —— pending decision 信号 ——
	// phase10-09 起统一 handoff 到唯一正式动作出口 Decision Detail。
	// 不再因为是否存在 decision_link 而退回到 Decision List 聚合入口。
	for _, d := range pendingDecisions {
		signals = append(signals, dashboard.FeedbackSignal{
			SignalFamily:  dashboard.FeedbackSignalFamilyPendingDecision,
			SignalCode:    dashboard.FeedbackSignalCodePendingDecision,
			Priority:      dashboard.FeedbackSignalPriorityP1PendingDecision,
			PriorityLabel: dashboard.FeedbackSignalPriorityP1PendingDecision.PriorityString(),
			Title:         fmt.Sprintf("待决策：%s", d.Title),
			Summary:       "有一条待决策记录需要处理",
			ActionLabel:   "查看决策",
			TargetType:    dashboard.DashboardTargetTypeDecisionDetail,
			TargetID:      d.DecisionID,
			TargetLabel:   d.Title,
			CreatedAt:     d.CreatedAt,
		})
	}

	// —— product asset coverage 缺口信号 ——
	for _, gap := range coverage.Gaps {
		switch gap.GapType {
		case "both":
			signals = append(signals, dashboard.FeedbackSignal{
				SignalFamily:  dashboard.FeedbackSignalFamilyProductAssetCoverage,
				SignalCode:    dashboard.FeedbackSignalCodeProductMissingBothBindings,
				Priority:      dashboard.FeedbackSignalPriorityP2ProductMissingBothBindings,
				PriorityLabel: dashboard.FeedbackSignalPriorityP2ProductMissingBothBindings.PriorityString(),
				Title:         "产品缺少模块与仓库绑定",
				Summary:       fmt.Sprintf("产品「%s」尚未绑定任何模块与仓库", gap.ProductName),
				ActionLabel:   "查看产品",
				TargetType:    dashboard.DashboardTargetTypeProductDetail,
				TargetID:      gap.ProductID,
				TargetLabel:   gap.ProductName,
				CreatedAt:     gap.CreatedAt,
			})
		case "repository":
			signals = append(signals, dashboard.FeedbackSignal{
				SignalFamily:  dashboard.FeedbackSignalFamilyProductAssetCoverage,
				SignalCode:    dashboard.FeedbackSignalCodeProductMissingRepositoryBinding,
				Priority:      dashboard.FeedbackSignalPriorityP3ProductMissingRepository,
				PriorityLabel: dashboard.FeedbackSignalPriorityP3ProductMissingRepository.PriorityString(),
				Title:         "产品缺少仓库绑定",
				Summary:       fmt.Sprintf("产品「%s」已绑定模块但尚未绑定仓库", gap.ProductName),
				ActionLabel:   "查看产品",
				TargetType:    dashboard.DashboardTargetTypeProductDetail,
				TargetID:      gap.ProductID,
				TargetLabel:   gap.ProductName,
				CreatedAt:     gap.CreatedAt,
			})
		case "module":
			signals = append(signals, dashboard.FeedbackSignal{
				SignalFamily:  dashboard.FeedbackSignalFamilyProductAssetCoverage,
				SignalCode:    dashboard.FeedbackSignalCodeProductMissingModuleBinding,
				Priority:      dashboard.FeedbackSignalPriorityP4ProductMissingModule,
				PriorityLabel: dashboard.FeedbackSignalPriorityP4ProductMissingModule.PriorityString(),
				Title:         "产品缺少模块绑定",
				Summary:       fmt.Sprintf("产品「%s」已绑定仓库但尚未绑定模块", gap.ProductName),
				ActionLabel:   "查看产品",
				TargetType:    dashboard.DashboardTargetTypeProductDetail,
				TargetID:      gap.ProductID,
				TargetLabel:   gap.ProductName,
				CreatedAt:     gap.CreatedAt,
			})
		}
	}

	return signals
}

// buildRepresentativeSignals 从 product asset coverage 缺口构建代表性缺口信号（最多 3 条）。
//
// representative_signals 只承接 product asset coverage 缺口，不混入 pending decision 信号。
// 排序规则与主队列一致：P2 > P3 > P4，同优先级按 created_at DESC。
func (s *QueryService) buildRepresentativeSignals(coverage *candidate.ProductAssetCoverageData) []dashboard.FeedbackSignal {
	// 复用 normalizeFeedbackSignals 中 product gap 部分的归一化逻辑，
	// 通过传入空 pending decisions 只取 product gap 信号。
	gapSignals := s.normalizeFeedbackSignals(nil, coverage)
	sortFeedbackSignals(gapSignals)
	return truncateSignals(gapSignals, maxRepresentativeSignals)
}

// sortFeedbackSignals 按优先级升序（P1 > P2 > P3 > P4）排序，
// 同优先级按 created_at 倒序（DESC）回退排序。
func sortFeedbackSignals(signals []dashboard.FeedbackSignal) {
	sort.SliceStable(signals, func(i, j int) bool {
		if signals[i].Priority != signals[j].Priority {
			return signals[i].Priority < signals[j].Priority
		}
		return signals[i].CreatedAt.After(signals[j].CreatedAt)
	})
}

// truncateSignals 截取前 n 条信号，保证返回非 nil 切片。
func truncateSignals(signals []dashboard.FeedbackSignal, n int) []dashboard.FeedbackSignal {
	result := make([]dashboard.FeedbackSignal, 0, n)
	for i := 0; i < len(signals) && i < n; i++ {
		result = append(result, signals[i])
	}
	return result
}

// ============================================================================
// ReadRecentActivity — 最近活动读取
// ============================================================================

// ReadRecentActivity 编排四个 canonical 模块的活动 reader，
// 在 Dashboard 模块内映射为统一 RecentActivityItem 列表，
// 按 activity_at 倒序排序，最多返回 10 条。
//
// 局部失败语义（phase05-04）：任一 reader 失败时返回局部失败（error），
// 不拖垮 ReadOverview 的整页成功语义。
//
// 活动类型映射对齐 phase05-12 spec §"活动类型映射"与 phase05-03 导航解释。
func (s *QueryService) ReadRecentActivity(ctx context.Context) ([]dashboard.RecentActivityItem, error) {
	moduleItems, err := s.activityReaders.ReadModuleActivities(ctx)
	if err != nil {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	productItems, err := s.activityReaders.ReadProductActivities(ctx)
	if err != nil {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	repositoryItems, err := s.activityReaders.ReadRepositoryActivities(ctx)
	if err != nil {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	decisionItems, err := s.activityReaders.ReadDecisionActivities(ctx)
	if err != nil {
		return nil, dashboard.ErrRecentActivityReadFailed
	}

	// 合并所有原始活动项
	allRaw := make([]candidate.RawActivityItem, 0,
		len(moduleItems)+len(productItems)+len(repositoryItems)+len(decisionItems))
	allRaw = append(allRaw, moduleItems...)
	allRaw = append(allRaw, productItems...)
	allRaw = append(allRaw, repositoryItems...)
	allRaw = append(allRaw, decisionItems...)

	// 映射为统一 RecentActivityItem
	items := make([]dashboard.RecentActivityItem, 0, len(allRaw))
	for _, raw := range allRaw {
		item, ok := mapRawActivity(raw)
		if !ok {
			// 未知 source 跳过，不阻断已知活动项的返回
			continue
		}
		items = append(items, item)
	}

	// 按 activity_at 倒序排序
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].ActivityAt.After(items[j].ActivityAt)
	})

	// 最多返回 10 条
	if len(items) > maxRecentActivities {
		items = items[:maxRecentActivities]
	}

	return items, nil
}

// mapRawActivity 将原始活动项映射为统一 RecentActivityItem。
//
// 映射规则对齐 phase05-12 spec §"活动项跳转目标映射"：
//   - module / release → MODULE_DETAIL
//   - product / product_module_binding → PRODUCT_DETAIL
//   - repository / product_repository_binding / module_repository_binding → REPOSITORY_DETAIL
//   - decision → DECISION_DETAIL
//
// 返回 ok=false 表示未知 source，调用方应跳过该条目。
func mapRawActivity(raw candidate.RawActivityItem) (dashboard.RecentActivityItem, bool) {
	var activityType dashboard.RecentActivityType
	var targetType dashboard.DashboardTargetType

	switch raw.Source {
	case "module":
		activityType = dashboard.RecentActivityTypeModule
		targetType = dashboard.DashboardTargetTypeModuleDetail
	case "release":
		activityType = dashboard.RecentActivityTypeRelease
		targetType = dashboard.DashboardTargetTypeModuleDetail
	case "product":
		activityType = dashboard.RecentActivityTypeProduct
		targetType = dashboard.DashboardTargetTypeProductDetail
	case "repository":
		activityType = dashboard.RecentActivityTypeRepository
		targetType = dashboard.DashboardTargetTypeRepositoryDetail
	case "decision":
		activityType = dashboard.RecentActivityTypeDecision
		targetType = dashboard.DashboardTargetTypeDecisionDetail
	case "product_module_binding":
		activityType = dashboard.RecentActivityTypeProductModuleBinding
		targetType = dashboard.DashboardTargetTypeProductDetail
	case "product_repository_binding":
		activityType = dashboard.RecentActivityTypeProductRepositoryBinding
		targetType = dashboard.DashboardTargetTypeRepositoryDetail
	case "module_repository_binding":
		activityType = dashboard.RecentActivityTypeModuleRepositoryBinding
		targetType = dashboard.DashboardTargetTypeRepositoryDetail
	default:
		return dashboard.RecentActivityItem{}, false
	}

	return dashboard.RecentActivityItem{
		ActivityType: activityType,
		ActivityAt:   raw.ActivityAt,
		TargetType:   targetType,
		TargetID:     raw.TargetID,
		TargetLabel:  raw.TargetLabel,
	}, true
}
