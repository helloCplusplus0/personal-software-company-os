// Package service — Onboarding 首轮状态读编排层。
//
// 单一 QueryService 承接 ReadFirstRunState 读组，
// 对齐 phase06-14 spec §"GetFirstRunState 必须由 canonical 数据读时派生"。
//
// 错误语义：
//   - ReadFirstRunState：任一计数 reader 失败 → 整页失败（返回 error）
//
// 跨模块读取全部通过 candidate/ 子包隔离，service 层不直接写跨模块 SQL。
//
// 文件落点：backend/internal/onboarding/service/query_service.go
package service

import (
	"context"

	"github.com/psco/backend/internal/onboarding"
	"github.com/psco/backend/internal/onboarding/candidate"
)

// QueryService 承接 Onboarding 首轮状态读编排。
//
// 依赖通过 platform 装配点注入（phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"）：
//   - firstRunReaders：四个 canonical 模块的计数 reader
//
// service 层不自行构造 candidate reader，也不直接跨模块写 SQL。
type QueryService struct {
	firstRunReaders *candidate.FirstRunReaders
}

// NewQueryService 构造 QueryService。
//
// firstRunReaders 由 platform 装配点构造并注入。
func NewQueryService(firstRunReaders *candidate.FirstRunReaders) *QueryService {
	return &QueryService{firstRunReaders: firstRunReaders}
}

// ReadFirstRunState 编排四个 canonical 模块的计数 reader，推导并返回 FirstRunState。
//
// 状态推导口径（phase06-14 spec §"first_run_state 状态推导"）：
//   - 四类都为 0 → not_started
//   - 至少 1 类、但未满四类 → in_progress
//   - 四类都至少 1 条 → completed
//
// current_step 推导（phase06-14 spec §"current_step 与 completion_progress 推导"）：
//   - 按 Product -> Repository -> Module -> Decision 找到第一个尚未完成的步骤
//   - completed 时 current_step 必须返回 complete
//
// completion_progress 冻结为 0 / 25 / 50 / 75 / 100。
// is_first_entry 在 not_started 时为 true，其他状态为 false。
func (s *QueryService) ReadFirstRunState(ctx context.Context) (*onboarding.FirstRunState, error) {
	counts, err := s.firstRunReaders.ReadCanonicalCounts(ctx)
	if err != nil {
		return nil, onboarding.ErrFirstRunStateReadFailed
	}

	// 统计四类对象中已持久化的类别数
	completedCategories := 0
	if counts.ProductCount > 0 {
		completedCategories++
	}
	if counts.RepositoryCount > 0 {
		completedCategories++
	}
	if counts.ModuleCount > 0 {
		completedCategories++
	}
	if counts.DecisionCount > 0 {
		completedCategories++
	}

	state := &onboarding.FirstRunState{}

	// 状态推导
	switch completedCategories {
	case 0:
		state.Status = onboarding.FirstRunStatusNotStarted
		state.IsFirstEntry = true
		state.CurrentStep = onboarding.OnboardingStepWelcome
		state.CompletionProgress = 0
	case 4:
		state.Status = onboarding.FirstRunStatusCompleted
		state.IsFirstEntry = false
		state.CurrentStep = onboarding.OnboardingStepComplete
		state.CompletionProgress = 100
	default:
		// 1~3 类已完成 → in_progress
		state.Status = onboarding.FirstRunStatusInProgress
		state.IsFirstEntry = false
		state.CurrentStep = deriveCurrentStep(counts)
		state.CompletionProgress = completedCategories * 25
	}

	return state, nil
}

// deriveCurrentStep 按 Product -> Repository -> Module -> Decision 找到第一个尚未完成的步骤。
//
// 推导顺序对齐 phase06-12 冻结的推荐执行顺序。
func deriveCurrentStep(counts *candidate.CanonicalCounts) onboarding.OnboardingStep {
	if counts.ProductCount == 0 {
		return onboarding.OnboardingStepProduct
	}
	if counts.RepositoryCount == 0 {
		return onboarding.OnboardingStepRepository
	}
	if counts.ModuleCount == 0 {
		return onboarding.OnboardingStepModule
	}
	// Decision 尚未完成（completedCategories < 4 时唯一剩余可能）
	return onboarding.OnboardingStepDecision
}
