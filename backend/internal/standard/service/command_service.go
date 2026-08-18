// Package service — standard 写入编排层。
//
// 单一 CommandService 承接 Standard 写路径编排。
// 对齐 phase14-04 / phase14-07 已冻结的写路径主线：
//   - 创建：name 非空 + status 归一 + 树校验 R1-R8（含 R2 时机）
//   - 更新：整树原子替换 + revision 追加，单一事务边界（事务语义在 repository 层承接）
//   - 删除：active 状态防误删拦截（先经 Update 置 retired）
//   - 绑定：五步校验链顺序固定（standard 存在 → 枚举合法 → 八格矩阵
//     → target 存在 → 唯一约束），不得调换
//
// 文件落点：backend/internal/standard/service/command_service.go
package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/psco/backend/internal/standard"
	"github.com/psco/backend/internal/standard/candidate"
	"github.com/psco/backend/internal/standard/repository"
)

// CommandService 承接 Standard 写路径编排。
//
// 依赖通过 platform 装配点注入：
//   - targetReader：跨模块绑定目标存在性校验（candidate 子包隔离）
//   - store：三表持久化承接位
type CommandService struct {
	targetReader *candidate.TargetReader
	store        *repository.StandardStore
}

// NewCommandService 构造 CommandService。
func NewCommandService(targetReader *candidate.TargetReader, store *repository.StandardStore) *CommandService {
	return &CommandService{
		targetReader: targetReader,
		store:        store,
	}
}

// CreateStandard 编排规范创建。
//
// 编排顺序：
//  1. name TrimSpace 非空
//  2. status 归一：空串 → draft；draft / active 合法；retired 拒绝创建；其他值非法
//  3. directory_tree 必带（允许单根空树走 R2 判定）
//  4. 树校验 R1-R8（结构规则 + 按 status 的 R2 时机）
//  5. 持久化（name 重复由 23505 承接）
func (s *CommandService) CreateStandard(ctx context.Context, input standard.CreateStandardInput) (*standard.StandardReadResult, error) {
	if strings.TrimSpace(input.Name) == "" {
		return nil, fmt.Errorf("%w: name is required", standard.ErrInvalidInput)
	}

	status, err := normalizeCreateStatus(input.Status)
	if err != nil {
		return nil, err
	}

	if input.DirectoryTree == nil {
		return nil, fmt.Errorf("%w: directory_tree is required", standard.ErrInvalidInput)
	}
	if err := standard.ValidateTreeStructure(input.DirectoryTree); err != nil {
		return nil, err
	}
	if err := standard.ValidateTreeForStatus(input.DirectoryTree, status); err != nil {
		return nil, err
	}

	result, err := s.store.CreateStandard(ctx, standard.CreateStandardInput{
		Name:          input.Name,
		Description:   input.Description,
		DirectoryTree: input.DirectoryTree,
		Status:        status,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateStandard 编排规范整树替换更新。
//
// 编排顺序：
//  1. standard_id 合法 UUID
//  2. change_summary 非空
//  3. directory_tree 必带（整树替换语义）
//  4. name 非 nil 时 TrimSpace 非空
//  5. status 非 nil 时值合法（draft / active / retired 均可，retired 合法）
//  6. 读当前记录（NotFound 透传）
//  7. 按生效 status 执行树校验 R1-R8（含 R2 时机）
//  8. 持久化（name 重复由 23505 承接）
func (s *CommandService) UpdateStandard(ctx context.Context, input standard.UpdateStandardInput) (*standard.StandardReadResult, error) {
	if err := validateUUID("standard", input.StandardID); err != nil {
		return nil, err
	}
	if strings.TrimSpace(input.ChangeSummary) == "" {
		return nil, fmt.Errorf("%w: change_summary is required", standard.ErrInvalidInput)
	}
	if input.DirectoryTree == nil {
		return nil, fmt.Errorf("%w: directory_tree is required for full replacement", standard.ErrInvalidInput)
	}
	if input.Name != nil && strings.TrimSpace(*input.Name) == "" {
		return nil, fmt.Errorf("%w: name must not be empty", standard.ErrInvalidInput)
	}
	if input.Status != nil && !isValidStatus(*input.Status) {
		return nil, fmt.Errorf("%w: invalid status %q", standard.ErrInvalidInput, *input.Status)
	}

	current, err := s.store.GetStandardByID(ctx, input.StandardID)
	if err != nil {
		return nil, err
	}

	effectiveStatus := current.Status
	if input.Status != nil {
		effectiveStatus = *input.Status
	}
	if err := standard.ValidateTreeStructure(input.DirectoryTree); err != nil {
		return nil, err
	}
	if err := standard.ValidateTreeForStatus(input.DirectoryTree, effectiveStatus); err != nil {
		return nil, err
	}

	result, err := s.store.UpdateStandard(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteStandard 编排规范删除。
// active 状态防误删拦截：先经 UpdateStandard 置 retired 再删除。
func (s *CommandService) DeleteStandard(ctx context.Context, standardID string) error {
	if err := validateUUID("standard", standardID); err != nil {
		return err
	}

	current, err := s.store.GetStandardByID(ctx, standardID)
	if err != nil {
		return err
	}
	if current.Status == standard.StandardStatusActive {
		return fmt.Errorf("%w: cannot delete an active standard; retire it first via UpdateStandard", standard.ErrInvalidInput)
	}

	return s.store.DeleteStandard(ctx, standardID)
}

// BindStandard 编排绑定建立。
//
// 校验链顺序固定（phase14-04 Scenario 冻结，不得调换）：
//  1. standard 存在（uuid 校验 + 读主记录 → ErrStandardNotFound）
//  2. 枚举合法（targetType ∈ 四类、role ∈ 两类）
//  3. 八格矩阵：template_source 仅 repository 合法（adopts 对四类全开）
//  4. target 存在（uuid 校验 + EnsureTargetExists）
//  5. 持久化（四元组唯一约束 23505 → "already bound"）
func (s *CommandService) BindStandard(ctx context.Context, input standard.BindStandardInput) (*standard.StandardBindingReadResult, error) {
	// 1. standard 存在
	if err := validateUUID("standard", input.StandardID); err != nil {
		return nil, err
	}
	if _, err := s.store.GetStandardByID(ctx, input.StandardID); err != nil {
		return nil, err
	}

	// 2. 枚举合法
	if err := validateBindingEnums(input.TargetType, input.Role); err != nil {
		return nil, err
	}

	// 3. 八格矩阵（phase14-02：template_source 仅 repository）
	if input.Role == standard.BindingRoleTemplateSource && input.TargetType != standard.BindingTargetRepository {
		return nil, fmt.Errorf("%w: binding role %q is only allowed for repository targets", standard.ErrInvalidInput, standard.BindingRoleTemplateSource)
	}

	// 4. target 存在
	if err := validateUUID("target", input.TargetID); err != nil {
		return nil, err
	}
	if err := s.targetReader.EnsureTargetExists(ctx, input.TargetType, input.TargetID); err != nil {
		return nil, err
	}

	// 5. 持久化
	result, err := s.store.InsertBinding(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindStandard 编排绑定解除（四元组定位）。
// standardID / targetID uuid 校验 + 枚举校验 → 删除（0 行 → ErrBindingNotFound）。
func (s *CommandService) UnbindStandard(ctx context.Context, input standard.UnbindStandardInput) error {
	if err := validateUUID("standard", input.StandardID); err != nil {
		return err
	}
	if err := validateUUID("target", input.TargetID); err != nil {
		return err
	}
	if err := validateBindingEnums(input.TargetType, input.Role); err != nil {
		return err
	}
	return s.store.DeleteBinding(ctx, input)
}

// validateUUID 校验 id 为合法 UUID（service 层统一小 helper）。
func validateUUID(field, value string) error {
	if _, err := uuid.Parse(value); err != nil {
		return fmt.Errorf("%w: invalid %s id", standard.ErrInvalidInput, field)
	}
	return nil
}

// normalizeCreateStatus 归一创建 status：空串 → draft；draft / active 合法；
// retired 拒绝创建；其他值非法。
func normalizeCreateStatus(status standard.StandardStatus) (standard.StandardStatus, error) {
	switch status {
	case "":
		return standard.StandardStatusDraft, nil
	case standard.StandardStatusDraft, standard.StandardStatusActive:
		return status, nil
	case standard.StandardStatusRetired:
		return "", fmt.Errorf("%w: cannot create a standard in retired status", standard.ErrInvalidInput)
	default:
		return "", fmt.Errorf("%w: invalid status %q", standard.ErrInvalidInput, status)
	}
}

// isValidStatus 判断 status 是否属于受控枚举（draft / active / retired）。
func isValidStatus(status standard.StandardStatus) bool {
	switch status {
	case standard.StandardStatusDraft, standard.StandardStatusActive, standard.StandardStatusRetired:
		return true
	default:
		return false
	}
}

// validateBindingEnums 校验绑定枚举合法（targetType ∈ 四类、role ∈ 两类）。
// service 层兜底校验（connect 层枚举转换已先行拦截，双保险供程序化调用方）。
func validateBindingEnums(targetType standard.BindingTargetType, role standard.BindingRole) error {
	switch targetType {
	case standard.BindingTargetRepository, standard.BindingTargetProduct,
		standard.BindingTargetDecision, standard.BindingTargetModule:
	default:
		return fmt.Errorf("%w: invalid target_type %q", standard.ErrInvalidInput, targetType)
	}
	switch role {
	case standard.BindingRoleTemplateSource, standard.BindingRoleAdopts:
	default:
		return fmt.Errorf("%w: invalid role %q", standard.ErrInvalidInput, role)
	}
	return nil
}
