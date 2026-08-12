// Package service — Template Reuse 业务编排层。
//
// QueryService 承接模板候选读取、模板预填、派生提示与模板来源复读四类读取能力。
// 所有数据从 product_modules 已持久化事实读时派生，不新增快照表。
//
// 文件落点：backend/internal/templatereuse/service/query_service.go
package service

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/psco/backend/internal/reusesummary"
	"github.com/psco/backend/internal/templatereuse"
	"github.com/psco/backend/internal/templatereuse/candidate"
)

// QueryService 承接 Template Reuse 的只读编排。
//
// 依赖通过 platform 装配点注入：
//   - candidateReaders：跨模块 candidate reader
//
// QueryService 只做 request 校验、candidate reader 调用、只读组合与轻量派生，
// 不得绕过 candidate reader 直接访问下游 repository。
type QueryService struct {
	candidateReaders candidateBindingsReader
	reuseSummaryRead reuseSummaryReader
}

const (
	consumerSurfaceWeeklyReview  = "weekly-review"
	consumerSurfaceProductCreate = "product-create"
	consumerSurfaceProductDetail = "product-detail"
)

type candidateBindingsReader interface {
	ReadAllProductModuleBindings(ctx context.Context) ([]candidate.ProductModuleRow, error)
}

type reuseSummaryReader interface {
	ReadReuseSummary(ctx context.Context, scope reusesummary.ReuseSummaryScope, moduleID, productID string) (*reusesummary.ReuseSummaryReadResult, error)
}

// NewQueryService 构造 QueryService。
func NewQueryService(candidateReaders candidateBindingsReader, reuseSummaryRead reuseSummaryReader) *QueryService {
	return &QueryService{
		candidateReaders: candidateReaders,
		reuseSummaryRead: reuseSummaryRead,
	}
}

// ============================================================================
// 模板候选列表
// ============================================================================

// ListTemplateCandidatesResult 模板候选列表读取结果。
type ListTemplateCandidatesResult struct {
	Candidates               []candidate.TemplateCandidateData
	DefaultActiveCandidateID string
}

// ListTemplateCandidates 读取模板候选列表。
//
// 派生规则（phase09-02 冻结）：
//   - 从 product_modules 读取所有活跃绑定关系
//   - 按 product_id 分组形成 product module sets
//   - 按 module_key（去重并升序排序后的 module_id 集合）分组形成候选
//   - 排序：source_product_count DESC → total_reuse_product_count DESC → latest_source_product_updated_at DESC
func (s *QueryService) ListTemplateCandidates(ctx context.Context) (*ListTemplateCandidatesResult, error) {
	// 1. 读取所有活跃 product-module 绑定
	rows, err := s.candidateReaders.ReadAllProductModuleBindings(ctx)
	if err != nil {
		return nil, fmt.Errorf("template reuse: list candidates: %w", err)
	}

	// 2. 按 product_id 分组
	productModuleSets := make(map[string]*productModuleSet)
	for _, row := range rows {
		set, ok := productModuleSets[row.ProductID]
		if !ok {
			set = &productModuleSet{
				ProductID:          row.ProductID,
				ProductName:        row.ProductName,
				ProductDescription: row.ProductDescription,
				ProductCreatedAt:   row.ProductCreatedAt,
			}
			productModuleSets[row.ProductID] = set
		}
		set.ModuleIDs = append(set.ModuleIDs, row.ModuleID)
		set.ModuleNames = append(set.ModuleNames, row.ModuleName)
		set.ModuleRows = append(set.ModuleRows, row)
	}

	// 3. 按 module_key 分组形成候选
	candidateGroups := make(map[string]*candidateGroup)
	for _, set := range productModuleSets {
		moduleKey := candidate.ComputeTemplateCandidateID(set.ModuleIDs)
		group, ok := candidateGroups[moduleKey]
		if !ok {
			group = &candidateGroup{
				ModuleKey: moduleKey,
				Modules:   buildModuleRefs(set.ModuleRows),
			}
			candidateGroups[moduleKey] = group
		}
		group.SourceProducts = append(group.SourceProducts, set)
		if set.ProductCreatedAt.After(group.LatestSourceProductUpdatedAt) {
			group.LatestSourceProductUpdatedAt = set.ProductCreatedAt
		}
	}

	// 4. 为每个候选计算 total_reuse_product_count
	//    total_reuse_product_count = 使用候选中任一模块的 Product 总数
	for _, group := range candidateGroups {
		moduleIDSet := make(map[string]bool)
		for _, m := range group.Modules {
			moduleIDSet[m.ModuleID] = true
		}
		productSet := make(map[string]bool)
		for _, row := range rows {
			if moduleIDSet[row.ModuleID] {
				productSet[row.ProductID] = true
			}
		}
		group.TotalReuseProductCount = len(productSet)
	}

	// 5. 构建候选列表
	candidates := make([]candidate.TemplateCandidateData, 0, len(candidateGroups))
	for _, group := range candidateGroups {
		// 模板标题：取第一个源产品名称
		title := group.SourceProducts[0].ProductName
		description := fmt.Sprintf("由 %d 个模块组成的模板，源自 %d 个产品",
			len(group.Modules), len(group.SourceProducts))

		moduleRefs := make([]candidate.TemplateModuleRefData, len(group.Modules))
		for i, m := range group.Modules {
			moduleRefs[i] = m
		}

		candidates = append(candidates, candidate.TemplateCandidateData{
			TemplateCandidateID:          group.ModuleKey,
			TemplateTitle:                title,
			TemplateDescription:          description,
			Modules:                      moduleRefs,
			SourceProductCount:           len(group.SourceProducts),
			TotalReuseProductCount:       group.TotalReuseProductCount,
			LatestSourceProductUpdatedAt: group.LatestSourceProductUpdatedAt,
		})
	}

	// 6. 排序
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].SourceProductCount != candidates[j].SourceProductCount {
			return candidates[i].SourceProductCount > candidates[j].SourceProductCount
		}
		if candidates[i].TotalReuseProductCount != candidates[j].TotalReuseProductCount {
			return candidates[i].TotalReuseProductCount > candidates[j].TotalReuseProductCount
		}
		return candidates[i].LatestSourceProductUpdatedAt.After(candidates[j].LatestSourceProductUpdatedAt)
	})

	// 7. 确定 default_active_candidate_id
	defaultID := ""
	if len(candidates) > 0 {
		defaultID = candidates[0].TemplateCandidateID
	}

	return &ListTemplateCandidatesResult{
		Candidates:               candidates,
		DefaultActiveCandidateID: defaultID,
	}, nil
}

// ============================================================================
// 模板预填详情
// ============================================================================

// TemplateCandidatePrefillResult 模板预填详情读取结果。
type TemplateCandidatePrefillResult struct {
	ResolutionStatus            string // "RESOLVED" or "UNAVAILABLE"
	UnavailableReasonText       string
	TemplateCandidateID         string
	TemplateTitle               string
	TemplateDescription         string
	SuggestedProductName        string
	SuggestedProductDescription string
	Modules                     []candidate.TemplateModuleRefData
	CapabilityGapHints          []DerivedHintData
}

// DerivedHintData 派生提示数据。
type DerivedHintData struct {
	HintType            string // "REUSE_OPPORTUNITY" or "CAPABILITY_GAP"
	Title               string
	ExplanationText     string
	CTAKind             string // "CREATE_PRODUCT_FROM_TEMPLATE" etc.
	TemplateCandidateID string
	CapabilityKey       string
	ModuleID            string
}

// GetTemplateCandidatePrefill 读取模板预填详情。
//
// 解析 template_candidate_id，返回 Product Create 表单的预填建议。
// 当 templateCandidateId 因底层事实变化而无法重新解析时，返回 UNAVAILABLE 成功态。
func (s *QueryService) GetTemplateCandidatePrefill(ctx context.Context, templateCandidateID string) (*TemplateCandidatePrefillResult, error) {
	// 先读取候选列表，从中找到匹配的候选
	candidatesResult, err := s.ListTemplateCandidates(ctx)
	if err != nil {
		return nil, fmt.Errorf("template reuse: get prefill: %w", err)
	}

	// 查找匹配的候选
	var matched *candidate.TemplateCandidateData
	for i := range candidatesResult.Candidates {
		if candidatesResult.Candidates[i].TemplateCandidateID == templateCandidateID {
			matched = &candidatesResult.Candidates[i]
			break
		}
	}

	// 候选未找到 → UNAVAILABLE
	if matched == nil {
		return &TemplateCandidatePrefillResult{
			ResolutionStatus:      "UNAVAILABLE",
			UnavailableReasonText: "模板候选已失效：底层 product_modules 事实已变化，当前候选无法再被解析",
			TemplateCandidateID:   templateCandidateID,
		}, nil
	}

	// 构建预填结果
	suggestedName := fmt.Sprintf("%s (基于模板)", matched.TemplateTitle)
	suggestedDescription := fmt.Sprintf("基于模板「%s」创建，包含 %d 个模块：%s",
		matched.TemplateTitle,
		len(matched.Modules),
		formatModuleNames(matched.Modules))

	// Product Create 场景下的 capability_gap_hint 继续由预填读取内联返回。
	// 当前 phase09-08 仅返回稳定的模块能力上下文提示，不复制第二次独立提示读取。
	capGapHints := deriveCreateCapabilityGapHints(matched.Modules, templateCandidateID)

	return &TemplateCandidatePrefillResult{
		ResolutionStatus:            "RESOLVED",
		TemplateCandidateID:         templateCandidateID,
		TemplateTitle:               matched.TemplateTitle,
		TemplateDescription:         matched.TemplateDescription,
		SuggestedProductName:        suggestedName,
		SuggestedProductDescription: suggestedDescription,
		Modules:                     matched.Modules,
		CapabilityGapHints:          capGapHints,
	}, nil
}

// ============================================================================
// 派生提示
// ============================================================================

// DerivedInsightHintsResult 派生提示读取结果。
type DerivedInsightHintsResult struct {
	ResolutionStatus      string
	UnavailableReasonText string
	Hints                 []DerivedHintData
}

// GetDerivedInsightHints 读取派生提示。
//
// 围绕传入的 template_candidate_id 计算当前消费面的正式提示。
// 在 Weekly Review 场景返回 reuse_opportunity_hint 与可选 capability_gap_hint。
// 当 templateCandidateId 漂移时返回 UNAVAILABLE 成功态。
func (s *QueryService) GetDerivedInsightHints(ctx context.Context, templateCandidateID, consumerSurface, reviewScopeKey string) (*DerivedInsightHintsResult, error) {
	if err := validateDerivedHintsScope(consumerSurface, reviewScopeKey); err != nil {
		return nil, err
	}

	// 先读取候选列表确认候选可解析
	candidatesResult, err := s.ListTemplateCandidates(ctx)
	if err != nil {
		return nil, fmt.Errorf("template reuse: get hints: %w", err)
	}

	var matched *candidate.TemplateCandidateData
	for i := range candidatesResult.Candidates {
		if candidatesResult.Candidates[i].TemplateCandidateID == templateCandidateID {
			matched = &candidatesResult.Candidates[i]
			break
		}
	}

	// 候选未找到 → UNAVAILABLE
	if matched == nil {
		return &DerivedInsightHintsResult{
			ResolutionStatus:      "UNAVAILABLE",
			UnavailableReasonText: "模板候选已失效：底层 product_modules 事实已变化，当前提示无法再被计算",
		}, nil
	}

	if consumerSurface == consumerSurfaceProductCreate {
		return &DerivedInsightHintsResult{
			ResolutionStatus: "RESOLVED",
			Hints:            []DerivedHintData{},
		}, nil
	}

	// 生成提示
	hints := make([]DerivedHintData, 0)

	// 1. reuse_opportunity_hint：当候选有多个源产品时，说明复用模式已形成
	if matched.SourceProductCount > 1 {
		hints = append(hints, DerivedHintData{
			HintType:            "REUSE_OPPORTUNITY",
			Title:               fmt.Sprintf("复用机会：%d 个产品共享相同模块组合", matched.SourceProductCount),
			ExplanationText:     fmt.Sprintf("当前有 %d 个产品使用了相同的 %d 个模块组合。基于此模板创建新产品可以快速复用已验证的模块组合。", matched.SourceProductCount, len(matched.Modules)),
			CTAKind:             "CREATE_PRODUCT_FROM_TEMPLATE",
			TemplateCandidateID: templateCandidateID,
		})
	} else if matched.SourceProductCount == 1 && matched.TotalReuseProductCount > 1 {
		hints = append(hints, DerivedHintData{
			HintType:            "REUSE_OPPORTUNITY",
			Title:               "复用机会：模块组合已被多个产品使用",
			ExplanationText:     fmt.Sprintf("当前模块组合内的模块共被 %d 个产品使用，具有复用潜力。", matched.TotalReuseProductCount),
			CTAKind:             "CREATE_PRODUCT_FROM_TEMPLATE",
			TemplateCandidateID: templateCandidateID,
		})
	}

	// 2. capability_gap_hints：基于 review 作用域 capability_summary 与当前模板未覆盖能力派生。
	capabilitySummary, err := s.readCapabilitySummaryForReviewScope(ctx, reviewScopeKey)
	if err != nil {
		return nil, fmt.Errorf("template reuse: get hints: %w", err)
	}
	capGapHints := deriveReviewCapabilityGapHints(capabilitySummary, matched.Modules, templateCandidateID)
	for _, h := range capGapHints {
		hints = append(hints, h)
	}

	return &DerivedInsightHintsResult{
		ResolutionStatus: "RESOLVED",
		Hints:            hints,
	}, nil
}

// ============================================================================
// 模板来源复读
// ============================================================================

// TemplateSourceSummaryResult 模板来源复读结果。
type TemplateSourceSummaryResult struct {
	ResolutionStatus      string
	UnavailableReasonText string
	TemplateCandidateID   string
	TemplateTitle         string
	TemplateDescription   string
	Modules               []candidate.TemplateModuleRefData
	TemplateSource        string
}

// GetTemplateSourceSummary 读取模板来源摘要。
//
// 基于成功回流链中的 templateCandidateId + templateSource 为 Product Detail 返回模板来源摘要。
// 当 templateCandidateId 漂移时返回 UNAVAILABLE 成功态。
func (s *QueryService) GetTemplateSourceSummary(ctx context.Context, templateCandidateID, templateSource string) (*TemplateSourceSummaryResult, error) {
	candidatesResult, err := s.ListTemplateCandidates(ctx)
	if err != nil {
		return nil, fmt.Errorf("template reuse: get source summary: %w", err)
	}

	var matched *candidate.TemplateCandidateData
	for i := range candidatesResult.Candidates {
		if candidatesResult.Candidates[i].TemplateCandidateID == templateCandidateID {
			matched = &candidatesResult.Candidates[i]
			break
		}
	}

	if matched == nil {
		return &TemplateSourceSummaryResult{
			ResolutionStatus:      "UNAVAILABLE",
			UnavailableReasonText: "模板来源不可复读：底层 product_modules 事实已变化，来源摘要无法再被解析",
			TemplateCandidateID:   templateCandidateID,
			TemplateSource:        templateSource,
		}, nil
	}

	return &TemplateSourceSummaryResult{
		ResolutionStatus:    "RESOLVED",
		TemplateCandidateID: templateCandidateID,
		TemplateTitle:       matched.TemplateTitle,
		TemplateDescription: matched.TemplateDescription,
		Modules:             matched.Modules,
		TemplateSource:      templateSource,
	}, nil
}

// ============================================================================
// 内部类型与辅助函数
// ============================================================================

// productModuleSet 单个 Product 的模块集合。
type productModuleSet struct {
	ProductID          string
	ProductName        string
	ProductDescription string
	ProductCreatedAt   time.Time
	ModuleIDs          []string
	ModuleNames        []string
	ModuleRows         []candidate.ProductModuleRow
}

// candidateGroup 模板候选分组。
type candidateGroup struct {
	ModuleKey                    string
	Modules                      []candidate.TemplateModuleRefData
	SourceProducts               []*productModuleSet
	TotalReuseProductCount       int
	LatestSourceProductUpdatedAt time.Time
}

// buildModuleRefs 从 ProductModuleRow 列表构建模板模块引用列表。
// 对同一 module_id 去重，保持首次出现顺序。
func buildModuleRefs(rows []candidate.ProductModuleRow) []candidate.TemplateModuleRefData {
	seen := make(map[string]bool)
	var refs []candidate.TemplateModuleRefData
	for _, row := range rows {
		if seen[row.ModuleID] {
			continue
		}
		seen[row.ModuleID] = true
		capKey := ""
		if row.CapabilityKey != nil {
			capKey = *row.CapabilityKey
		}
		refs = append(refs, candidate.TemplateModuleRefData{
			ModuleID:        row.ModuleID,
			ModuleName:      row.ModuleName,
			CapabilityKey:   capKey,
			CapabilityLabel: candidate.CapabilityLabel(capKey),
		})
	}
	return refs
}

// formatModuleNames 格式化模块名称列表为可读字符串。
func formatModuleNames(modules []candidate.TemplateModuleRefData) string {
	names := make([]string, len(modules))
	for i, m := range modules {
		names[i] = m.ModuleName
	}
	return strings.Join(names, "、")
}

// deriveCreateCapabilityGapHints 从模板模块派生 Product Create 预填场景的提示上下文。
//
// 当前阶段 Create 侧的 capability_gap_hint 仍由预填读取内联返回，不再额外走第二次提示读取。
// 在 phase09-09/10 真正接入页面消费前，这里只返回稳定的模块能力上下文提示。
func deriveCreateCapabilityGapHints(modules []candidate.TemplateModuleRefData, templateCandidateID string) []DerivedHintData {
	hints := make([]DerivedHintData, 0)
	for _, m := range modules {
		if m.CapabilityKey == "" {
			continue
		}
		hints = append(hints, DerivedHintData{
			HintType:            "CAPABILITY_GAP",
			Title:               fmt.Sprintf("能力模块：%s", m.ModuleName),
			ExplanationText:     fmt.Sprintf("模块「%s」提供 %s 能力，创建新产品后将自动继承此能力。", m.ModuleName, m.CapabilityLabel),
			CTAKind:             "GO_TO_MODULE_DETAIL",
			TemplateCandidateID: templateCandidateID,
			CapabilityKey:       m.CapabilityKey,
			ModuleID:            m.ModuleID,
		})
	}
	return hints
}

func (s *QueryService) readCapabilitySummaryForReviewScope(ctx context.Context, reviewScopeKey string) ([]reusesummary.CapabilitySummary, error) {
	// phase09 当前只有一个正式 Weekly Review 作用域；非空 reviewScopeKey 统一映射到
	// Weekly Review 既有 canonical capability_summary 来源。
	if reviewScopeKey == "" {
		return nil, fmt.Errorf("template reuse: review scope key is required: %w", templatereuse.ErrInvalidInput)
	}
	if s.reuseSummaryRead == nil {
		return nil, fmt.Errorf("template reuse: reuse summary reader is not configured")
	}

	result, err := s.reuseSummaryRead.ReadReuseSummary(ctx, reusesummary.ReuseSummaryScopeDashboard, "", "")
	if err != nil {
		return nil, err
	}
	if result == nil {
		return []reusesummary.CapabilitySummary{}, nil
	}
	return result.CapabilitySummary, nil
}

func validateDerivedHintsScope(consumerSurface, reviewScopeKey string) error {
	switch consumerSurface {
	case consumerSurfaceWeeklyReview:
		if reviewScopeKey == "" {
			return fmt.Errorf("template reuse: weekly review hints require review_scope_key: %w", templatereuse.ErrInvalidInput)
		}
		return nil
	case consumerSurfaceProductCreate:
		if reviewScopeKey != "" {
			return fmt.Errorf("template reuse: product create hints must not carry review_scope_key: %w", templatereuse.ErrInvalidInput)
		}
		return nil
	default:
		return fmt.Errorf("template reuse: unsupported consumer surface for derived hints: %w", templatereuse.ErrInvalidInput)
	}
}

func deriveReviewCapabilityGapHints(capabilitySummary []reusesummary.CapabilitySummary, modules []candidate.TemplateModuleRefData, templateCandidateID string) []DerivedHintData {
	coveredCapabilityKeys := make(map[string]struct{}, len(modules))
	for _, module := range modules {
		if module.CapabilityKey == "" {
			continue
		}
		coveredCapabilityKeys[module.CapabilityKey] = struct{}{}
	}

	hints := make([]DerivedHintData, 0)
	for _, summary := range capabilitySummary {
		if summary.CapabilityKey == "" || summary.EmptyStateText != "" {
			continue
		}
		if _, covered := coveredCapabilityKeys[summary.CapabilityKey]; covered {
			continue
		}

		hints = append(hints, DerivedHintData{
			HintType:            "CAPABILITY_GAP",
			Title:               fmt.Sprintf("能力缺口：补齐 %s", summary.CapabilityLabel),
			ExplanationText:     fmt.Sprintf("当前 Weekly Review 作用域显示「%s」是高频能力，但当前模板尚未覆盖它，可能阻碍下一次创造。", summary.CapabilityLabel),
			CTAKind:             "GO_TO_MODULE_DETAIL",
			TemplateCandidateID: templateCandidateID,
			CapabilityKey:       summary.CapabilityKey,
		})
	}
	return hints
}
