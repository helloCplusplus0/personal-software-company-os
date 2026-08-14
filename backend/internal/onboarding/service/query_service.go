// Package service — Onboarding 首轮状态读编排层。
//
// 单一 QueryService 承接 ReadFirstRunState 读组与 ReadOnboardingChainState 读组，
// 对齐 phase06-14 spec §"GetFirstRunState 必须由 canonical 数据读时派生"，
// 与 phase10-08 spec §"后端必须落地 GetOnboardingChainState 与最小恢复锚点"。
//
// 错误语义：
//   - ReadFirstRunState：任一计数 reader 失败 → 整页失败（返回 error）
//   - ReadOnboardingChainState：链状态 reader 或 recovery store 失败 → 整页失败（返回 error）
//
// 跨模块读取全部通过 candidate/ 子包隔离，service 层不直接写跨模块 SQL。
//
// 文件落点：backend/internal/onboarding/service/query_service.go
package service

import (
	"context"

	"github.com/psco/backend/internal/onboarding"
	"github.com/psco/backend/internal/onboarding/candidate"
	onboardingrepo "github.com/psco/backend/internal/onboarding/repository"
)

// QueryService 承接 Onboarding 首轮状态读编排与建链状态读编排。
//
// 依赖通过 platform 装配点注入（phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"）：
//   - firstRunReaders：四个 canonical 模块的计数 reader
//   - chainStateReaders：建链状态跨模块 reader（phase10-08 新增）
//   - recoveryStore：建链恢复锚点持久化（phase10-08 新增）
//
// service 层不自行构造 candidate reader，也不直接跨模块写 SQL。
type QueryService struct {
	firstRunReaders   *candidate.FirstRunReaders
	chainStateReaders *candidate.ChainStateReaders
	recoveryStore     *onboardingrepo.RecoveryStore
}

// NewQueryService 构造 QueryService。
//
// firstRunReaders / chainStateReaders / recoveryStore 由 platform 装配点构造并注入。
func NewQueryService(
	firstRunReaders *candidate.FirstRunReaders,
	chainStateReaders *candidate.ChainStateReaders,
	recoveryStore *onboardingrepo.RecoveryStore,
) *QueryService {
	return &QueryService{
		firstRunReaders:   firstRunReaders,
		chainStateReaders: chainStateReaders,
		recoveryStore:     recoveryStore,
	}
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

// ============================================================================
// 建链状态读取（phase10-08 新增）
// ============================================================================

// ReadOnboardingChainState 编排建链状态读取，返回六段式建链引导的最小恢复读模型。
//
// 恢复事实源优先级（phase10-08 spec §"刷新或回流时的恢复事实源优先级"）：
//  1. recovery_store.current_product_id（正式优先事实源）
//  2. GetFirstRunState 的冷启动摘要（仅在该锚点尚不存在时回退）
//
// 步骤推导：
//   - 以 current_product_id 为锚点，检查 product / repository 绑定 / module 绑定 / decision 持久化状态
//   - 按 Product -> Repository -> Module -> Decision 顺序找到当前步骤
//   - 已完成的步骤推导为 completed，下一步推导为 create / bind / handoff / complete
func (s *QueryService) ReadOnboardingChainState(ctx context.Context) (*onboarding.OnboardingChainState, error) {
	// 1. 读取恢复锚点
	currentProductID, err := s.recoveryStore.GetCurrentProductID(ctx)
	if err != nil {
		return nil, onboarding.ErrChainStateReadFailed
	}

	// 2. 读取跨模块事实
	facts, err := s.chainStateReaders.ReadChainStateFacts(ctx, currentProductID)
	if err != nil {
		return nil, onboarding.ErrChainStateReadFailed
	}

	// 3. 推导链状态
	state := deriveChainState(currentProductID, facts)
	return state, nil
}

// FreezeProductAnchor 显式冻结 current_product_id 锚点。
//
// phase10-08 spec §"product 步骤首个 Product 创建成功后必须冻结最小恢复锚点"。
// 供前端在 product 步骤创建成功后调用，确保后续步骤围绕该锚点解释。
func (s *QueryService) FreezeProductAnchor(ctx context.Context, productID string) error {
	if productID == "" {
		return nil
	}
	if err := s.recoveryStore.UpsertCurrentProductID(ctx, productID); err != nil {
		return onboarding.ErrRecoveryStoreOpFailed
	}
	return nil
}

// deriveChainState 根据 currentProductID 和跨模块事实推导链状态。
//
// 推导逻辑：
//   - cold_start：尚无 Product
//   - product 步骤：锚点尚未建立，或 Product 存在但当前主线仍需先建立 Product 锚点
//   - repository 步骤：Repository 存在但未绑定到 Product
//   - module 步骤：Module 存在但未绑定到 Product
//   - decision 步骤：Decision 不存在
//   - completed：四类对象均已持久化
func deriveChainState(currentProductID string, facts *candidate.ChainStateFacts) *onboarding.OnboardingChainState {
	state := &onboarding.OnboardingChainState{
		CurrentProductID: currentProductID,
	}

	// 冷启动：尚无 Product
	if !facts.HasProduct {
		state.CurrentStep = onboarding.OnboardingStepWelcome
		state.ResumeStatus = onboarding.ChainStateResumeStatusColdStart
		state.NextStepKind = onboarding.ChainStateNextStepKindCreate
		state.ReturnHint = "开始创建第一个产品"
		return state
	}

	// 未建立当前主线锚点时，不允许用全局 facts 猜测后续步骤。
	if currentProductID == "" {
		state.CurrentStep = onboarding.OnboardingStepProduct
		state.ResumeStatus = onboarding.ChainStateResumeStatusColdStart
		state.NextStepKind = onboarding.ChainStateNextStepKindCreate
		state.ReturnHint = "请先创建或确认当前主线产品"
		return state
	}

	// 已有主线锚点，检查后续步骤
	state.ResumeStatus = onboarding.ChainStateResumeStatusResuming

	// Product 已创建后，下一步应该进入 Repository 创建，而不是回退到 Product。
	if !facts.HasRepository {
		state.CurrentStep = onboarding.OnboardingStepRepository
		state.NextStepKind = onboarding.ChainStateNextStepKindCreate
		state.ReturnHint = "产品已创建，继续创建仓库"
		return state
	}

	// Repository 存在，检查是否已绑定到 Product
	if currentProductID != "" && !facts.HasRepositoryBound {
		state.CurrentStep = onboarding.OnboardingStepRepository
		state.NextStepKind = onboarding.ChainStateNextStepKindHandoff
		state.CanonicalHandoffTarget = "repository"
		state.ReturnHint = "仓库已创建，请在仓库详情页完成绑定"
		return state
	}

	// Module 存在，检查是否已绑定到 Product
	if !facts.HasModule {
		state.CurrentStep = onboarding.OnboardingStepModule
		state.NextStepKind = onboarding.ChainStateNextStepKindCreate
		state.ReturnHint = "继续创建模块"
		return state
	}

	if currentProductID != "" && !facts.HasModuleBound {
		state.CurrentStep = onboarding.OnboardingStepModule
		state.NextStepKind = onboarding.ChainStateNextStepKindHandoff
		state.CanonicalHandoffTarget = "module"
		state.ReturnHint = "模块已创建，请在模块详情页完成绑定"
		return state
	}

	// Decision 检查
	if !facts.HasDecision {
		state.CurrentStep = onboarding.OnboardingStepDecision
		state.NextStepKind = onboarding.ChainStateNextStepKindCreate
		state.ReturnHint = "继续记录决策"
		return state
	}

	// 四类都已完成
	state.CurrentStep = onboarding.OnboardingStepComplete
	state.ResumeStatus = onboarding.ChainStateResumeStatusCompleted
	state.NextStepKind = onboarding.ChainStateNextStepKindComplete
	state.ReturnHint = "首轮录入已完成"
	return state
}
