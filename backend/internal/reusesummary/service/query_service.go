// Package service — ReuseSummary 读编排层。
//
// QueryService 承接 GetReuseSummary 读组，
// 对齐 phase06-14 spec §"Reuse Summary 必须通过读时聚合返回最新已提交状态"。
//
// 读取语义：
//   - dashboard 作用域：返回全局复用快照，module_reuse_summary 与 capability_summary
//     都按"数量优先、时间次级"排序并最多返回前 5 条
//   - module_detail 作用域：围绕单一 module_id 返回该模块的直接复用反馈，
//     并在存在 capability_key 时返回对应 capability 摘要
//   - product_detail 作用域：围绕单一 product_id 先限定在已绑定模块范围内，
//     再返回全量复用 / capability 摘要
//
// 跨模块读取全部通过 candidate/ 子包隔离，service 层不直接写跨模块 SQL。
//
// 文件落点：backend/internal/reusesummary/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/reusesummary"
	"github.com/psco/backend/internal/reusesummary/candidate"
)

// dashboard 作用域下每类摘要的最大返回条数。
const maxDashboardSummaryEntries = 5

// capabilityLabelMap 后端内置的单一 capability_key -> capability_label 映射。
//
// phase06-14 spec §"capability_summary 聚合与映射"：
//   - capability_label 必须来自后端内置的单一 capability_key -> capability_label 映射
//   - 当前阶段不得让后端、前端与 fixture 各自维护三套不同映射表
//
// 未在映射表中的 capability_key 以 capability_key 本身作为 label 回退。
var capabilityLabelMap = map[string]string{
	"web_frontend":     "Web Frontend",
	"backend_api":      "Backend API",
	"database":         "Database",
	"auth":             "Authentication",
	"cli_tool":         "CLI Tool",
	"build_system":     "Build System",
	"observability":    "Observability",
	"deployment":       "Deployment",
	"documentation":    "Documentation",
	"state_management": "State Management",
}

// QueryService 承接 ReuseSummary 复用感知读编排。
//
// 依赖通过 platform 装配点注入：
//   - reuseReaders：三种作用域的复用 reader
type QueryService struct {
	reuseReaders *candidate.ReuseReaders
}

// NewQueryService 构造 QueryService。
func NewQueryService(reuseReaders *candidate.ReuseReaders) *QueryService {
	return &QueryService{reuseReaders: reuseReaders}
}

// ReadReuseSummary 按作用域读取复用感知。
//
// 作用域参数使用关系（phase06-10 proto 冻结）：
//   - dashboard：module_id 与 product_id 均不使用
//   - module_detail：使用 module_id，不使用 product_id
//   - product_detail：使用 product_id，不使用 module_id
func (s *QueryService) ReadReuseSummary(ctx context.Context, scope reusesummary.ReuseSummaryScope, moduleID, productID string) (*reusesummary.ReuseSummaryReadResult, error) {
	switch scope {
	case reusesummary.ReuseSummaryScopeDashboard:
		return s.readDashboard(ctx)
	case reusesummary.ReuseSummaryScopeModuleDetail:
		return s.readModuleDetail(ctx, moduleID)
	case reusesummary.ReuseSummaryScopeProductDetail:
		return s.readProductDetail(ctx, productID)
	default:
		return nil, reusesummary.ErrInvalidScope
	}
}

// readDashboard 读取 Dashboard 作用域的全局复用快照。
func (s *QueryService) readDashboard(ctx context.Context) (*reusesummary.ReuseSummaryReadResult, error) {
	moduleData, err := s.reuseReaders.ReadDashboardReuse(ctx)
	if err != nil {
		return nil, reusesummary.ErrReuseSummaryReadFailed
	}

	capData, err := s.reuseReaders.ReadCapabilityAggregates(ctx, nil)
	if err != nil {
		return nil, reusesummary.ErrReuseSummaryReadFailed
	}

	moduleSummary := buildModuleReuseSummaryList(moduleData, maxDashboardSummaryEntries)
	capSummary := buildCapabilitySummaryList(capData, maxDashboardSummaryEntries)

	return &reusesummary.ReuseSummaryReadResult{
		ModuleReuseSummary: moduleSummary,
		CapabilitySummary:  capSummary,
	}, nil
}

// readModuleDetail 读取 Module Detail 作用域的单一模块复用快照。
func (s *QueryService) readModuleDetail(ctx context.Context, moduleID string) (*reusesummary.ReuseSummaryReadResult, error) {
	moduleData, err := s.reuseReaders.ReadModuleDetailReuse(ctx, moduleID)
	if err != nil {
		return nil, reusesummary.ErrReuseSummaryReadFailed
	}

	// 模块不存在 → 返回空态
	if moduleData == nil {
		return reusesummary.NewEmptyReuseSummaryReadResult(), nil
	}

	moduleSummary := []reusesummary.ModuleReuseSummary{
		buildModuleReuseSummary(*moduleData),
	}

	// 若模块有 capability_key，返回对应 capability 摘要
	var capSummary []reusesummary.CapabilitySummary
	if moduleData.CapabilityKey != nil && *moduleData.CapabilityKey != "" {
		capData, err := s.reuseReaders.ReadCapabilityAggregates(ctx, []string{moduleID})
		if err != nil {
			return nil, reusesummary.ErrReuseSummaryReadFailed
		}
		capSummary = buildCapabilitySummaryList(capData, 0) // 不裁剪
	} else {
		capSummary = []reusesummary.CapabilitySummary{
			{
				CapabilityKey:  "",
				EmptyStateText: "当前模块未填写 capability_key，暂无能力聚合摘要",
			},
		}
	}

	return &reusesummary.ReuseSummaryReadResult{
		ModuleReuseSummary: moduleSummary,
		CapabilitySummary:  capSummary,
	}, nil
}

// readProductDetail 读取 Product Detail 作用域的复用快照。
func (s *QueryService) readProductDetail(ctx context.Context, productID string) (*reusesummary.ReuseSummaryReadResult, error) {
	moduleData, err := s.reuseReaders.ReadProductDetailReuse(ctx, productID)
	if err != nil {
		return nil, reusesummary.ErrReuseSummaryReadFailed
	}

	// 提取已绑定模块的 ID 列表，用于 capability 聚合过滤
	moduleIDs := make([]string, 0, len(moduleData))
	for _, m := range moduleData {
		moduleIDs = append(moduleIDs, m.ModuleID)
	}

	var capData []candidate.CapabilityAggregateData
	if len(moduleIDs) > 0 {
		capData, err = s.reuseReaders.ReadCapabilityAggregates(ctx, moduleIDs)
		if err != nil {
			return nil, reusesummary.ErrReuseSummaryReadFailed
		}
	}

	// Product Detail 不裁剪，全量返回
	moduleSummary := buildModuleReuseSummaryList(moduleData, 0)
	capSummary := buildCapabilitySummaryList(capData, 0)

	return &reusesummary.ReuseSummaryReadResult{
		ModuleReuseSummary: moduleSummary,
		CapabilitySummary:  capSummary,
	}, nil
}

// buildModuleReuseSummaryList 从原始数据构建 ModuleReuseSummary 列表。
// maxEntries > 0 时裁剪到指定数量；maxEntries <= 0 时不裁剪。
func buildModuleReuseSummaryList(data []candidate.ModuleReuseData, maxEntries int) []reusesummary.ModuleReuseSummary {
	result := make([]reusesummary.ModuleReuseSummary, 0, len(data))
	for _, m := range data {
		result = append(result, buildModuleReuseSummary(m))
	}

	if maxEntries > 0 && len(result) > maxEntries {
		result = result[:maxEntries]
	}

	if result == nil {
		result = []reusesummary.ModuleReuseSummary{}
	}
	return result
}

// buildModuleReuseSummary 从单条原始数据构建 ModuleReuseSummary。
func buildModuleReuseSummary(m candidate.ModuleReuseData) reusesummary.ModuleReuseSummary {
	var explanation string
	switch {
	case m.ReuseProductCount == 0:
		explanation = fmt.Sprintf("模块「%s」当前尚未被任何 Product 复用", m.ModuleName)
	case m.ReuseProductCount == 1:
		explanation = fmt.Sprintf("模块「%s」当前被 1 个 Product 复用，尚未形成跨 Product 复用", m.ModuleName)
	default:
		explanation = fmt.Sprintf("模块「%s」当前被 %d 个 Product 复用", m.ModuleName, m.ReuseProductCount)
	}

	return reusesummary.ModuleReuseSummary{
		ModuleID:          m.ModuleID,
		ReuseProductCount: m.ReuseProductCount,
		LatestReuseAt:     m.LatestReuseAt,
		ExplanationText:   explanation,
	}
}

// buildCapabilitySummaryList 从原始数据构建 CapabilitySummary 列表。
// maxEntries > 0 时裁剪到指定数量；maxEntries <= 0 时不裁剪。
func buildCapabilitySummaryList(data []candidate.CapabilityAggregateData, maxEntries int) []reusesummary.CapabilitySummary {
	result := make([]reusesummary.CapabilitySummary, 0, len(data))
	for _, c := range data {
		label := capabilityLabelMap[c.CapabilityKey]
		if label == "" {
			label = c.CapabilityKey
		}
		result = append(result, reusesummary.CapabilitySummary{
			CapabilityKey:            c.CapabilityKey,
			CapabilityLabel:          label,
			SupportingModuleCount:    c.SupportingModuleCount,
			LatestCapabilityUpdateAt: c.LatestUpdateAt,
		})
	}

	if maxEntries > 0 && len(result) > maxEntries {
		result = result[:maxEntries]
	}

	if result == nil {
		result = []reusesummary.CapabilitySummary{}
	}
	return result
}
