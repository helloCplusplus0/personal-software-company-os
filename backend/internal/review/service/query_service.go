// Package service — Review 业务编排层。
//
// QueryService 组合既有 Dashboard / Decision Center / Reuse Summary 的 canonical query service，
// 只为 Review 提供轻量只读组合与派生，不重写第二套 SQL 或 candidate reader。
//
// CommandService 只承接 SubmitReviewResult 这一类最小 review result sink，
// 实体写入继续直走既有 canonical command service。
//
// 文件落点：backend/internal/review/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/dashboard"
	dashboardservice "github.com/psco/backend/internal/dashboard/service"
	"github.com/psco/backend/internal/decisioncenter"
	dcservice "github.com/psco/backend/internal/decisioncenter/service"
	"github.com/psco/backend/internal/reusesummary"
	reusesummaryservice "github.com/psco/backend/internal/reusesummary/service"
)

// 当前版本展示上限（对齐 phase08-04 冻结结论，上限由 service 层承接）。
const (
	maxPendingDecisions = 5 // pending decisions 最多展示 5 条 proposed 决策
)

// QueryService 承接 Review 的只读组合编排。
//
// 依赖通过 platform 装配点注入：
//   - dashboardQuerySvc：Dashboard 既有 QueryService
//   - decisionCenterQuerySvc：Decision Center 既有 QueryService
//   - reuseSummaryQuerySvc：Reuse Summary 既有 QueryService
//
// QueryService 只做 request 校验、既有 service 调用、只读组合与轻量派生，
// 不得绕过既有 service 直接访问下游 repository。
type QueryService struct {
	dashboardQuerySvc      *dashboardservice.QueryService
	decisionCenterQuerySvc *dcservice.QueryService
	reuseSummaryQuerySvc   *reusesummaryservice.QueryService
}

// NewQueryService 构造 QueryService。
func NewQueryService(
	dashboardQuerySvc *dashboardservice.QueryService,
	decisionCenterQuerySvc *dcservice.QueryService,
	reuseSummaryQuerySvc *reusesummaryservice.QueryService,
) *QueryService {
	return &QueryService{
		dashboardQuerySvc:      dashboardQuerySvc,
		decisionCenterQuerySvc: decisionCenterQuerySvc,
		reuseSummaryQuerySvc:   reuseSummaryQuerySvc,
	}
}

// DailyReviewContextResult Daily Review 上下文读取结果。
type DailyReviewContextResult struct {
	CurrentFocusSignals    []dashboard.FeedbackSignal
	RepresentativeSignals  []dashboard.FeedbackSignal
	PendingDecisions       []decisioncenter.DecisionListItem
}

// GetDailyReviewContext 获取 Daily Review 上下文。
//
// 消费：
//   - dashboard.QueryService.ReadFeedbackSignal（current_focus_signals + representative_signals）
//   - decisioncenter.QueryService.ListDecisions(status = proposed)（pending_decisions）
func (s *QueryService) GetDailyReviewContext(ctx context.Context) (*DailyReviewContextResult, error) {
	// 读取反馈信号
	feedbackResult, err := s.dashboardQuerySvc.ReadFeedbackSignal(ctx)
	if err != nil {
		return nil, fmt.Errorf("review: get daily review context: read feedback: %w", err)
	}

	// 读取 proposed 决策作为 pending decisions
	decisions, err := s.decisionCenterQuerySvc.ListDecisions(ctx, decisioncenter.ListQuery{
		StatusFilter: decisioncenter.DecisionStatusProposed,
	})
	if err != nil {
		return nil, fmt.Errorf("review: get daily review context: list decisions: %w", err)
	}

	// 截断到 maxPendingDecisions
	if len(decisions) > maxPendingDecisions {
		decisions = decisions[:maxPendingDecisions]
	}

	return &DailyReviewContextResult{
		CurrentFocusSignals:   feedbackResult.CurrentFocusSignals,
		RepresentativeSignals: feedbackResult.AssetFeedbackSummary.RepresentativeSignals,
		PendingDecisions:      decisions,
	}, nil
}

// WeeklyReviewContextResult Weekly Review 上下文读取结果。
type WeeklyReviewContextResult struct {
	Overview               *dashboard.DashboardOverview
	RecentActivities       []dashboard.RecentActivityItem
	RepresentativeSignals  []dashboard.FeedbackSignal
	ModuleReuseSummary     []reusesummary.ModuleReuseSummary
	CapabilitySummary      []reusesummary.CapabilitySummary
}

// GetWeeklyReviewContext 获取 Weekly Review 上下文。
//
// 消费：
//   - dashboard.QueryService.ReadOverview
//   - dashboard.QueryService.ReadRecentActivity
//   - dashboard.QueryService.ReadFeedbackSignal（representative_signals）
//   - reusesummary.QueryService.ReadReuseSummary(scope = dashboard)
func (s *QueryService) GetWeeklyReviewContext(ctx context.Context) (*WeeklyReviewContextResult, error) {
	// 读取概览
	overview, err := s.dashboardQuerySvc.ReadOverview(ctx)
	if err != nil {
		return nil, fmt.Errorf("review: get weekly review context: read overview: %w", err)
	}

	// 读取最近活动
	activities, err := s.dashboardQuerySvc.ReadRecentActivity(ctx)
	if err != nil {
		return nil, fmt.Errorf("review: get weekly review context: read recent activity: %w", err)
	}

	// 读取反馈信号（只取 representative_signals）
	feedbackResult, err := s.dashboardQuerySvc.ReadFeedbackSignal(ctx)
	if err != nil {
		return nil, fmt.Errorf("review: get weekly review context: read feedback: %w", err)
	}

	// 读取复用感知摘要（scope = dashboard）
	reuseResult, err := s.reuseSummaryQuerySvc.ReadReuseSummary(ctx, reusesummary.ReuseSummaryScopeDashboard, "", "")
	if err != nil {
		return nil, fmt.Errorf("review: get weekly review context: read reuse summary: %w", err)
	}

	return &WeeklyReviewContextResult{
		Overview:              overview,
		RecentActivities:      activities,
		RepresentativeSignals: feedbackResult.AssetFeedbackSummary.RepresentativeSignals,
		ModuleReuseSummary:    reuseResult.ModuleReuseSummary,
		CapabilitySummary:     reuseResult.CapabilitySummary,
	}, nil
}