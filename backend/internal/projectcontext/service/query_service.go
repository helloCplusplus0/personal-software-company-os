// Package service — projectcontext 只读聚合编排层。
//
// 单一 QueryService 承接 GetProjectContext 只读编排。
// 对齐 phase11-07 已冻结的 query service owner 结论。
//
// 跨模块读取全部通过 candidate/ 子包隔离，service 层不直接写跨模块 SQL。
//
// 文件落点：backend/internal/projectcontext/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/projectcontext"
	"github.com/psco/backend/internal/projectcontext/candidate"
	"github.com/psco/backend/internal/projectcontext/renderer"
)

// QueryService 承接 projectcontext 只读聚合编排。
//
// 依赖通过 platform 装配点注入：
//   - contextReaders：跨模块 reader 集合
type QueryService struct {
	contextReaders *candidate.ContextReaders
}

// NewQueryService 构造 QueryService。
func NewQueryService(contextReaders *candidate.ContextReaders) *QueryService {
	return &QueryService{
		contextReaders: contextReaders,
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
