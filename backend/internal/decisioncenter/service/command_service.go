// Package service — 写组业务编排层。
//
// 统一承接 DecisionWrite (RecordDecision) 与 DecisionLinkWrite (LinkDecisionToTarget) 编排
// （phase03-10 §10.4 单文件编排，不拆两个独立 service 文件）。
//
// 文件落点：backend/internal/decisioncenter/service/command_service.go
package service

import (
	"context"
	"strings"

	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/decisioncenter/candidate"
	"github.com/psco/backend/internal/decisioncenter/repository"
)

// CommandService 承接写组业务编排。
type CommandService struct {
	decisions  *repository.DecisionStore
	links      *repository.LinkStore
	candidates *candidate.ModuleCandidateRead
}

// NewCommandService 构造 CommandService。
//
// candidates 由应用装配点注入（phase03-10 §10.5），
// service 层不自行构造 ModuleCandidateRead，也不直接写跨模块 SQL。
// Module 存在性校验通过 candidate.ModuleExists 承接。
func NewCommandService(decisions *repository.DecisionStore, links *repository.LinkStore, candidates *candidate.ModuleCandidateRead) *CommandService {
	return &CommandService{decisions: decisions, links: links, candidates: candidates}
}

// CreateDecision 承接 DecisionWrite (RecordDecision)。
//
// 校验顺序（phase03-04 / phase03-10 §5.5 冻结）：
//  1. 必填字段非空校验（title / context / problem / choice / reason / status）
//     去首尾空白后不得为空字符串
//  2. status 取值合法性校验（proposed / active / superseded / archived）
//  3. alternatives 条目校验（每个条目去首尾空白后不得为空字符串）
//  4. source_module_id 可选来源校验（§5.11）：非空时校验 Module 存在性
//
// 成功返回新建 Decision 标识（decision_id），支持前端回流到 DecisionDetailPage
// （phase03-10 §6.4 不返回完整 Decision 对象，避免脱离 DecisionDetailRead 的第二套回流路径）。
func (s *CommandService) CreateDecision(ctx context.Context, req decisioncenter.CreateDecisionRequest) (*decisioncenter.CreateDecisionResponse, error) {
	// 1. 必填字段非空校验
	title := strings.TrimSpace(req.Title)
	context := strings.TrimSpace(req.Context)
	problem := strings.TrimSpace(req.Problem)
	choice := strings.TrimSpace(req.Choice)
	reason := strings.TrimSpace(req.Reason)
	if title == "" || context == "" || problem == "" || choice == "" || reason == "" || req.Status == "" {
		return nil, decisioncenter.ErrInvalidInput
	}

	// 2. status 取值合法性校验
	if !isValidDecisionStatus(req.Status) {
		return nil, decisioncenter.ErrInvalidStatus
	}

	// 3. alternatives 条目校验（去首尾空白后不得为空字符串）
	trimmedAlts := make([]string, 0, len(req.Alternatives))
	for _, a := range req.Alternatives {
		trimmed := strings.TrimSpace(a)
		if trimmed == "" {
			return nil, decisioncenter.ErrInvalidAlternatives
		}
		trimmedAlts = append(trimmedAlts, trimmed)
	}

	// 4. source_module_id 可选来源校验（§5.11）
	//    非空时校验 Module 存在性（跨模块只读校验由 candidate 子包承接）
	//    空字符串表示无来源上下文，跳过校验
	sourceModuleID := strings.TrimSpace(req.SourceModuleID)
	if sourceModuleID != "" {
		if err := decisioncenter.ValidateModuleID(sourceModuleID); err != nil {
			return nil, err
		}
		moduleExists, err := s.candidates.ModuleExists(ctx, sourceModuleID)
		if err != nil {
			return nil, err
		}
		if !moduleExists {
			return nil, decisioncenter.ErrModuleNotFound
		}
	}

	// 写入（已 trim 的字段覆盖原值，保证存储一致）
	req.Title = title
	req.Context = context
	req.Problem = problem
	req.Choice = choice
	req.Reason = reason
	req.Alternatives = trimmedAlts
	req.SourceModuleID = sourceModuleID

	d, err := s.decisions.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &decisioncenter.CreateDecisionResponse{DecisionID: d.ID}, nil
}

// LinkDecisionToTarget 承接 DecisionLinkWrite (LinkDecisionToTarget)。
//
// 校验顺序（phase03-04 / phase03-10 §6.5 冻结）：
//  1. target_type 取值合法性（当前阶段只允许 module），越界返回 ErrInvalidTargetType
//  2. decision_id 格式校验 → Decision 存在性校验，不存在返回 ErrDecisionNotFound
//  3. module_id 格式校验 → Module 存在性校验，不存在返回 ErrModuleNotFound
//  4. 重复关联检测，已存在返回 ErrDuplicateLink
//
// 成功返回空响应（无返回体），前端通过 DecisionDetailRead 重新读取
// （phase03-10 §6.4 不返回 link 结果或 detail reread 标识）。
func (s *CommandService) LinkDecisionToTarget(ctx context.Context, decisionID string, req decisioncenter.LinkDecisionToTargetRequest) error {
	// 1. target_type 取值合法性校验
	if req.TargetType != decisioncenter.DecisionLinkTargetTypeModule {
		return decisioncenter.ErrInvalidTargetType
	}

	// 2. decision_id 格式校验 → Decision 存在性校验
	if err := decisioncenter.ValidateDecisionID(decisionID); err != nil {
		return err
	}
	exists, err := s.decisions.Exists(ctx, decisionID)
	if err != nil {
		return err
	}
	if !exists {
		return decisioncenter.ErrDecisionNotFound
	}

	// 3. module_id 格式校验 → Module 存在性校验（跨模块只读校验由 candidate 子包承接）
	if err := decisioncenter.ValidateModuleID(req.ModuleID); err != nil {
		return err
	}
	moduleExists, err := s.candidates.ModuleExists(ctx, req.ModuleID)
	if err != nil {
		return err
	}
	if !moduleExists {
		return decisioncenter.ErrModuleNotFound
	}

	// 4. 重复关联检测 → 写入
	return s.links.Create(ctx, decisionID, req.ModuleID)
}

// isValidDecisionStatus 校验 status 是否在冻结枚举内
// （phase03-10 §5.6：proposed / active / superseded / archived）。
func isValidDecisionStatus(s decisioncenter.DecisionStatus) bool {
	switch s {
	case decisioncenter.DecisionStatusProposed,
		decisioncenter.DecisionStatusActive,
		decisioncenter.DecisionStatusSuperseded,
		decisioncenter.DecisionStatusArchived:
		return true
	}
	return false
}
