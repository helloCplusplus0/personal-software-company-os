// Package service — 写组业务编排层。
//
// 统一承接 CreateRepository / BindRepositoryToProduct / MapModuleToRepository 编排。
//
// 文件落点：backend/internal/repositorybinding/service/command_service.go
package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/psco/backend/internal/repositorybinding"
	"github.com/psco/backend/internal/repositorybinding/candidate"
	"github.com/psco/backend/internal/repositorybinding/repository"
)

// CommandService 承接写组业务编排。
//
// 依赖注入：
//   - repos: repositories 表读写
//   - bindings: product_repositories / module_repositories 写入
//   - productCandidate: Product 存在性与 active 状态校验（跨模块）
//   - moduleCandidate: Module 存在性与 active 状态校验（跨模块）
type CommandService struct {
	repos            *repository.RepositoryStore
	bindings         *repository.BindingStore
	productCandidate *candidate.ProductCandidateRead
	moduleCandidate  *candidate.ModuleCandidateRead
}

// NewCommandService 构造 CommandService。
func NewCommandService(repos *repository.RepositoryStore, bindings *repository.BindingStore, productCandidate *candidate.ProductCandidateRead, moduleCandidate *candidate.ModuleCandidateRead) *CommandService {
	return &CommandService{repos: repos, bindings: bindings, productCandidate: productCandidate, moduleCandidate: moduleCandidate}
}

// CreateRepository 承接 RepositoryCreateWrite。
//
// 校验顺序（phase04-04 / phase04-12 spec）：
//  1. 必填字段去首尾空白后非空（name / url / provider）
//  2. status 只允许 active / archived
//  3. Repository.name / url / provider 的最小格式合法性
//
// 成功返回新建仓库 id，支撑前端回流到 Repository Binding Detail / Workspace。
func (s *CommandService) CreateRepository(ctx context.Context, req repositorybinding.CreateRepositoryRequest) (*repositorybinding.CreateRepositoryResponse, error) {
	// 输入校验
	name := strings.TrimSpace(req.Name)
	url := strings.TrimSpace(req.URL)
	provider := strings.TrimSpace(req.Provider)
	if name == "" || url == "" || provider == "" {
		return nil, repositorybinding.ErrInvalidInput
	}
	if req.Status != repositorybinding.RepositoryStatusActive && req.Status != repositorybinding.RepositoryStatusArchived {
		return nil, repositorybinding.ErrInvalidStatus
	}

	// 写入
	r, err := s.repos.Create(ctx, name, url, provider, req.Status)
	if err != nil {
		return nil, err
	}
	return &repositorybinding.CreateRepositoryResponse{RepositoryID: r.ID}, nil
}

// BindRepositoryToProduct 承接 RepositoryProductBindingWrite。
//
// 校验顺序（phase04-04 / phase04-12 spec）：
//  1. repository_id / product_id 格式合法
//  2. Repository 存在且状态合法
//  3. Product 存在且状态为 active
//  4. 重复绑定检测
//
// 成功后默认 reread owner 必须是 RepositoryDetailRead（phase04-12 spec）。
// ProductDetailRead 不得成为 BindRepositoryToProduct 的第二 reread owner。
func (s *CommandService) BindRepositoryToProduct(ctx context.Context, repositoryID, productID string) error {
	// 1. ID 格式校验
	if err := repositorybinding.ValidateRepositoryID(repositoryID); err != nil {
		return err
	}
	if err := repositorybinding.ValidateProductID(productID); err != nil {
		return err
	}

	// 2. Repository 存在性校验
	if _, err := s.repos.GetByID(ctx, repositoryID); err != nil {
		return err
	}

	// 3. Product 存在且 active 校验（通过 candidate 子包，不直接写跨模块 SQL）
	//
	// 三态分流（phase04-04 错误语义前提）：
	//   - 不存在 → ErrProductNotFound（404）
	//   - 存在但非 active → ErrProductNotActive（400）
	productExists, productActive, err := s.productCandidate.CheckProductExistsActive(ctx, productID)
	if err != nil {
		return fmt.Errorf("check product exists and active: %w", err)
	}
	if !productExists {
		return repositorybinding.ErrProductNotFound
	}
	if !productActive {
		return repositorybinding.ErrProductNotActive
	}

	// 4. 重复绑定检测（由 DB UNIQUE 约束兜底）
	return s.bindings.BindProduct(ctx, repositoryID, productID)
}

// MapModuleToRepository 承接 RepositoryModuleMappingWrite。
//
// 校验顺序（phase04-04 / phase04-12 spec）：
//  1. repository_id / module_id 格式合法
//  2. Repository 存在且状态合法
//  3. Module 存在且状态为 active
//  4. 重复映射检测
//
// 成功后默认 reread owner 必须是 RepositoryDetailRead（phase04-12 spec）。
func (s *CommandService) MapModuleToRepository(ctx context.Context, repositoryID, moduleID string) error {
	// 1. ID 格式校验
	if err := repositorybinding.ValidateRepositoryID(repositoryID); err != nil {
		return err
	}
	if err := repositorybinding.ValidateModuleID(moduleID); err != nil {
		return err
	}

	// 2. Repository 存在性校验
	if _, err := s.repos.GetByID(ctx, repositoryID); err != nil {
		return err
	}

	// 3. Module 存在且 active 校验（通过 candidate 子包，不直接写跨模块 SQL）
	//
	// 三态分流（phase04-04 错误语义前提）：
	//   - 不存在 → ErrModuleNotFound（404）
	//   - 存在但非 active → ErrModuleNotActive（400）
	moduleExists, moduleActive, err := s.moduleCandidate.CheckModuleExistsActive(ctx, moduleID)
	if err != nil {
		return fmt.Errorf("check module exists and active: %w", err)
	}
	if !moduleExists {
		return repositorybinding.ErrModuleNotFound
	}
	if !moduleActive {
		return repositorybinding.ErrModuleNotActive
	}

	// 4. 重复映射检测（由 DB UNIQUE 约束兜底）
	return s.bindings.MapModule(ctx, repositoryID, moduleID)
}
