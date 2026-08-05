// Package service — 写组业务编排层。
//
// 统一承接 ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite 编排（§9.4 单文件编排，不拆三个 service）。
//
// 文件落点：backend/internal/moduleregistry/service/command_service.go
package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/repository"
)

// CommandService 承接写组业务编排。
type CommandService struct {
	modules  *repository.ModuleStore
	releases *repository.ReleaseStore
	bindings *repository.BindingStore
}

// NewCommandService 构造 CommandService。
func NewCommandService(modules *repository.ModuleStore, releases *repository.ReleaseStore, bindings *repository.BindingStore) *CommandService {
	return &CommandService{modules: modules, releases: releases, bindings: bindings}
}

// CreateModule 承接 ModuleCreateWrite。
//
// 校验顺序（§5.6 模块准入规则）：
//  1. 输入字段非空（name / description / status）
//  2. status 取值合法（active / archived）
//  3. name 唯一性
//
// 成功返回带 id / created_at 的完整模块对象，支持前端回流到 ModuleDetailPage。
func (s *CommandService) CreateModule(ctx context.Context, req moduleregistry.CreateModuleRequest) (*moduleregistry.Module, error) {
	// 输入校验
	name := strings.TrimSpace(req.Name)
	description := strings.TrimSpace(req.Description)
	if name == "" || description == "" {
		return nil, moduleregistry.ErrInvalidInput
	}
	if req.Status != moduleregistry.ModuleStatusActive && req.Status != moduleregistry.ModuleStatusArchived {
		return nil, moduleregistry.ErrInvalidStatus
	}

	// 名称唯一性校验（§5.6）
	existing, err := s.modules.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("check name uniqueness: %w", err)
	}
	if existing != nil {
		return nil, moduleregistry.ErrDuplicateModuleName
	}

	// 写入
	m, err := s.modules.Create(ctx, name, description, req.Status)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// CreateRelease 承接 ModuleReleaseWrite。
//
// §5.7 版本写入承接 CreateRelease（最小字段 version / status / released_at，
// module_id 由上下文隐式承接）。
//
// 校验顺序：
//  1. 模块存在性（Release Create 必须依附有效当前模块上下文，不得创建孤儿 release）
//  2. 输入字段非空（version / status / released_at）
//  3. status 取值合法
//  4. released_at 可解析为 RFC3339 时间
func (s *CommandService) CreateRelease(ctx context.Context, moduleID string, req moduleregistry.CreateReleaseRequest) (*moduleregistry.Release, error) {
	// 校验 ID 格式
	if err := moduleregistry.ValidateModuleID(moduleID); err != nil {
		return nil, err
	}

	// 模块存在性校验（防止孤儿 release）
	if _, err := s.modules.GetByID(ctx, moduleID); err != nil {
		return nil, err
	}

	// 输入校验
	version := strings.TrimSpace(req.Version)
	if version == "" || req.ReleasedAt == "" {
		return nil, moduleregistry.ErrInvalidInput
	}
	if req.Status != moduleregistry.ReleaseStatusActive && req.Status != moduleregistry.ReleaseStatusArchived {
		return nil, moduleregistry.ErrInvalidReleaseStatus
	}

	// 时间格式校验（RFC3339）
	parsed, err := time.Parse(time.RFC3339, req.ReleasedAt)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid released_at (expect RFC3339): %v", moduleregistry.ErrInvalidInput, err)
	}
	// 统一以 RFC3339 字符串传给 DB，避免时区漂移
	releasedAt := parsed.UTC().Format(time.RFC3339)

	// 写入
	r, err := s.releases.Create(ctx, moduleID, version, req.Status, releasedAt)
	if err != nil {
		// UNIQUE(module_id, version) 冲突由 DB 兜底，调用方可据此提示用户
		return nil, err
	}
	return r, nil
}

// BindModuleToProduct 承接 ModuleBindingWrite 的产品绑定子动作。
//
// 校验顺序：
//  1. 模块存在性
//  2. Product 候选存在性（候选读取前提校验，§"Product 连接边界"）
//  3. 重复绑定检测
func (s *CommandService) BindModuleToProduct(ctx context.Context, moduleID, productID string) error {
	if err := moduleregistry.ValidateModuleID(moduleID); err != nil {
		return err
	}
	if err := moduleregistry.ValidateProductID(productID); err != nil {
		return moduleregistry.ErrProductNotFound
	}
	if _, err := s.modules.GetByID(ctx, moduleID); err != nil {
		return err
	}
	exists, err := s.bindings.ProductExists(ctx, productID)
	if err != nil {
		return err
	}
	if !exists {
		return moduleregistry.ErrProductNotFound
	}
	return s.bindings.BindProduct(ctx, moduleID, productID)
}

// MapModuleToRepository 承接 ModuleBindingWrite 的仓库映射子动作。
//
// 校验顺序：
//  1. 模块存在性
//  2. Repository 候选存在性（候选读取前提校验，§"Repository 连接边界"）
//  3. 重复映射检测
func (s *CommandService) MapModuleToRepository(ctx context.Context, moduleID, repositoryID string) error {
	if err := moduleregistry.ValidateModuleID(moduleID); err != nil {
		return err
	}
	if err := moduleregistry.ValidateRepositoryID(repositoryID); err != nil {
		return moduleregistry.ErrRepositoryNotFound
	}
	if _, err := s.modules.GetByID(ctx, moduleID); err != nil {
		return err
	}
	exists, err := s.bindings.RepositoryExists(ctx, repositoryID)
	if err != nil {
		return err
	}
	if !exists {
		return moduleregistry.ErrRepositoryNotFound
	}
	return s.bindings.MapRepository(ctx, moduleID, repositoryID)
}
