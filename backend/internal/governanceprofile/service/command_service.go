// Package service — governanceprofile 写入编排层。
//
// 单一 CommandService 承接 UpdateGovernanceProfile 保存编排。
// 对齐 phase13-08 已冻结的写路径主线：
//   - 单一正式保存入口，可写集合收敛在 service 层校验
//   - 显式排除 read-only 字段（track_type / current_phase_*，不在输入结构中）
//   - 主记录与两组 bindings 在同一事务边界内保存
//   - 不触发目录扫描、模板自动同步、正文拉取入库或自动状态建议（手工维护优先）
//
// 文件落点：backend/internal/governanceprofile/service/command_service.go
package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/psco/backend/internal/governanceprofile"
	"github.com/psco/backend/internal/governanceprofile/candidate"
	"github.com/psco/backend/internal/governanceprofile/repository"
)

// CommandService 承接治理画像保存编排。
//
// 依赖通过 platform 装配点注入：
//   - repositoryReader：跨模块仓库存在性前提校验（candidate 子包隔离）
//   - profileStore：单一事务边界的持久化承接位
type CommandService struct {
	repositoryReader *candidate.RepositoryReader
	profileStore     *repository.ProfileStore
}

// NewCommandService 构造 CommandService。
func NewCommandService(repositoryReader *candidate.RepositoryReader, profileStore *repository.ProfileStore) *CommandService {
	return &CommandService{
		repositoryReader: repositoryReader,
		profileStore:     profileStore,
	}
}

// UpdateGovernanceProfile 编排治理画像保存。
//
// 编排顺序：
//  1. 校验目标仓库已在 PSCO 登记（不存在则返回 ErrRepositoryNotFound）
//  2. 校验可写字段输入（字段分类约束 + 8 项资产矩阵）
//  3. 在单一事务边界内保存主记录与两组 bindings
//
// 失败语义：
//   - repository 不存在 → ErrRepositoryNotFound（NotFound）
//   - 输入违反字段矩阵 → ErrInvalidInput（InvalidArgument）
//   - 保存失败         → ErrGovernanceProfileSaveFailed（Internal，整体失败不写半套状态）
func (s *CommandService) UpdateGovernanceProfile(ctx context.Context, input governanceprofile.UpdateGovernanceProfileInput) (*governanceprofile.GovernanceProfileReadResult, error) {
	// 1. 仓库存在性前提
	if err := s.repositoryReader.EnsureRepositoryExists(ctx, input.RepositoryID); err != nil {
		return nil, err
	}

	// 2. 可写字段校验（read-only 字段不在输入结构中，天然排除）
	if err := validateUpdateInput(input); err != nil {
		return nil, err
	}

	// 3. 单一事务边界保存
	result, err := s.profileStore.SaveProfile(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("governanceprofile: update governance profile: %w", err)
	}
	return result, nil
}

// validateUpdateInput 校验保存输入是否满足 phase13-04 字段分类矩阵与 phase13-05 资产矩阵。
//
// 校验规则：
//   - repository_id 非空
//   - canonical_root_files 至少 1 项（required）；file_name / role 非空；file_name 不重复
//   - global_asset_bindings 至少 1 项（required）；name 必须属于 8 项冻结矩阵；
//     kind / entry_ref / role 非空；name 不重复；前 5 项摘要型资产 structured_summary 必填
//   - template_source 允许为空（optional）
func validateUpdateInput(input governanceprofile.UpdateGovernanceProfileInput) error {
	if strings.TrimSpace(input.RepositoryID) == "" {
		return fmt.Errorf("%w: repository_id is required", governanceprofile.ErrInvalidInput)
	}

	// canonical_root_files：required 集合，至少 1 项
	if len(input.CanonicalRootFiles) == 0 {
		return fmt.Errorf("%w: canonical_root_files must not be empty", governanceprofile.ErrInvalidInput)
	}
	rootFileNames := make(map[string]struct{}, len(input.CanonicalRootFiles))
	for i, f := range input.CanonicalRootFiles {
		if strings.TrimSpace(f.FileName) == "" {
			return fmt.Errorf("%w: canonical_root_files[%d].file_name is required", governanceprofile.ErrInvalidInput, i)
		}
		if strings.TrimSpace(f.Role) == "" {
			return fmt.Errorf("%w: canonical_root_files[%d].role is required", governanceprofile.ErrInvalidInput, i)
		}
		if _, dup := rootFileNames[f.FileName]; dup {
			return fmt.Errorf("%w: canonical_root_files duplicate file_name %q", governanceprofile.ErrInvalidInput, f.FileName)
		}
		rootFileNames[f.FileName] = struct{}{}
	}

	// global_asset_bindings：required 集合，至少 1 项，name 限定 8 项冻结矩阵
	if len(input.GlobalAssetBindings) == 0 {
		return fmt.Errorf("%w: global_asset_bindings must not be empty", governanceprofile.ErrInvalidInput)
	}
	assetNames := make(map[string]struct{}, len(input.GlobalAssetBindings))
	for i, b := range input.GlobalAssetBindings {
		if !governanceprofile.IsKnownGlobalAsset(b.Name) {
			return fmt.Errorf("%w: global_asset_bindings[%d].name %q is not in the frozen 8-asset matrix", governanceprofile.ErrInvalidInput, i, b.Name)
		}
		if strings.TrimSpace(b.Kind) == "" {
			return fmt.Errorf("%w: global_asset_bindings[%d].kind is required", governanceprofile.ErrInvalidInput, i)
		}
		if strings.TrimSpace(b.EntryRef) == "" {
			return fmt.Errorf("%w: global_asset_bindings[%d].entry_ref is required", governanceprofile.ErrInvalidInput, i)
		}
		if strings.TrimSpace(b.Role) == "" {
			return fmt.Errorf("%w: global_asset_bindings[%d].role is required", governanceprofile.ErrInvalidInput, i)
		}
		if _, dup := assetNames[b.Name]; dup {
			return fmt.Errorf("%w: global_asset_bindings duplicate name %q", governanceprofile.ErrInvalidInput, b.Name)
		}
		assetNames[b.Name] = struct{}{}

		// 前 5 项摘要型资产 structured_summary 必填（phase13-05 矩阵）
		if governanceprofile.GlobalAssetSummaryRequired(b.Name) {
			if b.StructuredSummary == nil || strings.TrimSpace(*b.StructuredSummary) == "" {
				return fmt.Errorf("%w: global_asset_bindings[%d].structured_summary is required for %q", governanceprofile.ErrInvalidInput, i, b.Name)
			}
		}
	}

	return nil
}
