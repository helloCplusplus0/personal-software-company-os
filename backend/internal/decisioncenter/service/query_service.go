// Package service — 读组业务编排层。
//
// 统一承接 DecisionListRead / DecisionDetailRead / DecisionModuleCandidateRead 编排
// （phase03-10 §10.4 单文件编排，不拆 list/detail/candidate 三个 service 文件）。
//
// 文件落点：backend/internal/decisioncenter/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/decisioncenter/candidate"
	"github.com/psco/backend/internal/decisioncenter/repository"
)

// QueryService 承接读组业务编排。
type QueryService struct {
	decisions *repository.DecisionStore
	links     *repository.LinkStore
	candidates *candidate.ModuleCandidateRead
}

// NewQueryService 构造 QueryService。
//
// candidates 由应用装配点注入（phase03-10 §10.5），
// service 层不自行构造 ModuleCandidateRead。
func NewQueryService(decisions *repository.DecisionStore, links *repository.LinkStore, candidates *candidate.ModuleCandidateRead) *QueryService {
	return &QueryService{decisions: decisions, links: links, candidates: candidates}
}

// ListDecisions 承接 DecisionListRead。
//
// 列表读取至少承接（phase03-10 §5.8）：
// id / title / status / created_at / link_count / linked_module_summary
//
// link_count 与 linked_module_summary 的计算口径在 repository 层通过聚合 SQL 完成
// （phase03-10 §5.9），service 层只承接编排与空切片保证。
func (s *QueryService) ListDecisions(ctx context.Context, q decisioncenter.ListQuery) ([]decisioncenter.DecisionListItem, error) {
	statusFilter := ""
	if q.StatusFilter != "" && q.StatusFilter != "all" {
		statusFilter = string(q.StatusFilter)
	}

	items, err := s.decisions.List(ctx, q.QueryText, statusFilter)
	if err != nil {
		return nil, err
	}
	if items == nil {
		items = []decisioncenter.DecisionListItem{}
	}
	return items, nil
}

// GetDecisionDetail 承接 DecisionDetailRead。
//
// 详情读取统一承接（phase03-10 §5.8）：
//   - Decision 核心字段（id / title / context / problem / alternatives / choice / reason / impact / status / created_at）
//   - 结构化模板字段
//   - 已关联 Module 列表（linked_modules）
//   - 最小来源上下文（source_context）
//
// 编排流程：
//  1. 校验 decision_id 格式，防止无效 ID 打到数据库
//  2. 读取 Decision 核心字段 + 来源上下文，不存在返回 ErrDecisionNotFound
//  3. 读取已关联 Module 列表（附带 module_name）
//  4. 组装 source_context（从 decisions.source_module_id LEFT JOIN modules 获取）
//
// source_context 承接语义（phase03-10 §5.11）：
//   - 从 Module Detail 带上下文进入 Decision Create 时，source_module_id 被持久化
//   - DecisionDetailRead 通过 source_context 返回来源 Module 的最小标识
//   - 支持"持续到用户完成正式 LinkDecisionToTarget 或主动放弃关联"的跨刷新承接
//   - 无来源时 source_module_id / source_module_name 均为空字符串
//   - 入口上下文中的预填 Module 在正式 LinkDecisionToTarget 写入前不计入 linked_modules
func (s *QueryService) GetDecisionDetail(ctx context.Context, decisionID string) (*decisioncenter.DecisionDetail, error) {
	// 校验 ID 格式，防止无效 ID 打到数据库
	if err := decisioncenter.ValidateDecisionID(decisionID); err != nil {
		return nil, err
	}

	result, err := s.decisions.GetByID(ctx, decisionID)
	if err != nil {
		return nil, err
	}

	linkedModules, err := s.links.ListByDecisionID(ctx, decisionID)
	if err != nil {
		return nil, fmt.Errorf("list linked modules: %w", err)
	}
	if linkedModules == nil {
		linkedModules = []decisioncenter.LinkedModule{}
	}

	return &decisioncenter.DecisionDetail{
		Decision: result.Decision,
		LinkedModules: linkedModules,
		// 从 decisions.source_module_id LEFT JOIN modules 组装来源上下文（§5.11）
		SourceContext: decisioncenter.SourceContext{
			SourceModuleID:   result.SourceModuleID,
			SourceModuleName: result.SourceModuleName,
		},
	}, nil
}

// ListModuleCandidates 承接 DecisionModuleCandidateRead。
//
// 候选读取排序与排除规则在 candidate 层实现（phase03-10 §5.10）。
// service 层只承接编排与空切片保证。
//
// 无可关联候选时返回空列表（[]），不返回 null，不得将空结果误报为接口错误
// （phase03-10 §5.10 空候选结果语义）。
func (s *QueryService) ListModuleCandidates(ctx context.Context, decisionID string) ([]decisioncenter.DecisionModuleCandidate, error) {
	// 校验 ID 格式；无效 ID 直接返回 ErrDecisionNotFound（对齐 phase03-04 详情读取前置资源不存在）
	if err := decisioncenter.ValidateDecisionID(decisionID); err != nil {
		return nil, err
	}

	// 校验 Decision 存在性（候选读取前置：目标 Decision 必须存在）
	exists, err := s.decisions.Exists(ctx, decisionID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, decisioncenter.ErrDecisionNotFound
	}

	items, err := s.candidates.List(ctx, decisionID)
	if err != nil {
		return nil, err
	}
	if items == nil {
		items = []decisioncenter.DecisionModuleCandidate{}
	}
	return items, nil
}
