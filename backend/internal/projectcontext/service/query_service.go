// Package service — projectcontext 只读聚合编排层。
//
// 单一 QueryService 承接 GetProjectContext / GetProjectBrief 只读编排。
// 对齐 phase11-07 与 phase13-10 已冻结的 query service owner 结论。
//
// 跨模块读取全部通过 candidate/ 子包隔离，service 层不直接写跨模块 SQL；
// 全局规范读取通过 candidate.StandardReader 接口注入（phase14-07）。
//
// 2026-08-18 phase14-10 T7 用户裁决：画像残余彻底退役，原 candidate.
// GovernanceProfileReader 依赖已随画像后端模块（governance profile internal
// 包）整体删除，brief 编排不再包含画像主记录读取与 current_phase 派生步骤。
//
// 文件落点：backend/internal/projectcontext/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/projectcontext"
	"github.com/psco/backend/internal/projectcontext/candidate"
	"github.com/psco/backend/internal/projectcontext/renderer"
	"github.com/psco/backend/internal/standard"
)

// QueryService 承接 projectcontext 只读聚合编排。
//
// 依赖通过 platform 装配点注入：
//   - contextReaders：跨模块 reader 集合
//   - standardReader：全局规范读取接口（phase14-04，由
//     standard/service.QueryService 实现）
type QueryService struct {
	contextReaders *candidate.ContextReaders
	standardReader candidate.StandardReader
}

// NewQueryService 构造 QueryService。
func NewQueryService(contextReaders *candidate.ContextReaders, standardReader candidate.StandardReader) *QueryService {
	return &QueryService{
		contextReaders: contextReaders,
		standardReader: standardReader,
	}
}

// GetProjectContext 编排跨模块读取，返回最小只读项目上下文。
//
// 编排顺序：
//  1. 读取 Repository 身份（不存在则返回 ErrRepositoryNotFound）
//  2. 校验 Repository Binding 是否完成（不完整则返回 ErrRepositoryBindingIncomplete）
//  3. 读取关联 Product 摘要
//  4. 读取关联 Module 摘要
//  5. 读取关联 Decision 摘要（两类 module-link 派生命中 + 去重 + archived 过滤）
//  6. 读取规则入口
//  7. 读取 phase 入口
//  8. 读取当前阶段边界摘要
//
// 失败语义：
//   - Repository 不存在 → 返回 ErrRepositoryNotFound
//   - Repository Binding 不完整 → 返回 ErrRepositoryBindingIncomplete
//   - 其他读取失败 → 返回 ErrProjectContextReadFailed
//   - Product/Module/Decision 为空 → 返回空列表/空值，不视为错误
func (s *QueryService) GetProjectContext(ctx context.Context, repositoryID string) (*projectcontext.ProjectContextReadResult, error) {
	// 1. Repository 身份
	repo, err := s.contextReaders.ReadRepository(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}

	// 2. Repository Binding 完成态
	if err := s.contextReaders.ValidateRepositoryBinding(ctx, repositoryID); err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}

	// 3. Product 摘要
	product, err := s.contextReaders.ReadProduct(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}

	// 4. Module 摘要
	modules, err := s.contextReaders.ReadModules(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}
	if modules == nil {
		modules = []projectcontext.ModuleSummary{}
	}

	// 5. Decision 摘要
	decisions, err := s.contextReaders.ReadDecisions(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}
	if decisions == nil {
		decisions = []projectcontext.DecisionSummary{}
	}

	// 6. 规则入口
	rules := s.contextReaders.ReadRules(ctx)

	// 7. Phase 入口
	phases := s.contextReaders.ReadPhases(ctx)

        // 8. 当前阶段边界摘要
        boundaries := s.contextReaders.ReadBoundaries(ctx)

	return &projectcontext.ProjectContextReadResult{
		Repository: repo,
		Product:    product,
		Modules:    modules,
		Decisions:  decisions,
		Rules:      rules,
		Phases:     phases,
                Boundaries: boundaries,
	}, nil
}

// ExportProjectContext 先调用 GetProjectContext 获取结构化只读结果，再单向渲染为 Markdown。
//
// 导出语义：
//   - 严格从 GetProjectContext 结构化结果单向派生
//   - 不绕过结构化读取主线直接扫描目录或拼接内容
//   - 失败语义与 GetProjectContext 一致
func (s *QueryService) ExportProjectContext(ctx context.Context, repositoryID string) (string, error) {
	result, err := s.GetProjectContext(ctx, repositoryID)
	if err != nil {
		return "", err
	}
	return renderer.RenderMarkdown(result), nil
}

// GetProjectBrief 编排跨模块读取，返回 agent 项目简报（phase13-10 正式主线）。
//
// 编排顺序（以 repository_id 为唯一锚点）：
//  1. 读取 Repository 身份（不存在则返回 ErrRepositoryNotFound）
//  2. 读取 products[]（数组语义，通过 product_repositories）
//  3. 读取 modules[]（通过 module_repositories）
//  4. 读取 decisions[]（两类 module-link 派生命中 + 去重 + archived 过滤）
//  5. 读取 standards[]（通过 candidate.StandardReader 经 standard_bindings
//     任意 role 反查，含 directory_tree 全树；phase14-07 新增）——
//     旧 global_assets 顶层块已移除，两组 bindings 信息唯一来自 standards[]
//
// 2026-08-18 phase14-10 T7 用户裁决：画像残余彻底退役，原步骤"读取治理画像
// 主记录核心字段"与"current_phase 从画像派生"已删除；template_source 语义由
// standard_bindings(role=template_source) 承接，track_type/时间轴不保留。
//
// 失败语义（phase13-10 冻结，2026-08-18 T7 裁决后画像 not_found 分支移除）：
//   - Repository 不存在 → ErrRepositoryNotFound（CodeNotFound）
//   - 数组为空是合法状态，不做 Repository Binding 完整性强制校验
//   - 其他读取失败 → ErrProjectContextReadFailed（CodeInternal）
//   - standards[] 读取失败 → standard.ErrStandardReadFailed（CodeInternal，
//     phase14-04 冻结 reader 失败语义，由 connecterrors 统一映射）
func (s *QueryService) GetProjectBrief(ctx context.Context, repositoryID string) (*projectcontext.ProjectBriefReadResult, error) {
	// 1. Repository 身份
	repo, err := s.contextReaders.ReadRepository(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}

	// 2. products[]（数组语义）
	products, err := s.contextReaders.ReadProducts(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}

	// 3. modules[]
	modules, err := s.contextReaders.ReadModules(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}
	if modules == nil {
		modules = []projectcontext.ModuleSummary{}
	}

	// 4. decisions[]
	decisions, err := s.contextReaders.ReadDecisions(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, err)
	}
	if decisions == nil {
		decisions = []projectcontext.DecisionSummary{}
	}

	// 5. standards[]（candidate 接口承接，不复制 standard 表 SQL；
	//    失败按 reader 冻结语义透传 wrapped ErrStandardReadFailed）
	standards, err := s.standardReader.ListStandardsByRepository(ctx, repositoryID)
	if err != nil {
		return nil, err
	}
	if standards == nil {
		standards = []standard.StandardReadResult{}
	}

	return &projectcontext.ProjectBriefReadResult{
		Repository: repo,
		Products:   products,
		Modules:    modules,
		Decisions:  decisions,
		Standards:  standards,
	}, nil
}
