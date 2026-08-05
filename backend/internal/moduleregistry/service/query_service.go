// Package service — 读组业务编排层。
//
// 统一承接 ModuleListRead 与 ModuleDetailRead 编排（§9.4 单文件编排，不拆 list/detail）。
// Decision 附属读取作为详情读取的内嵌数据组装逻辑在此承接（§9.5 不设独立 service 文件）。
//
// 文件落点：backend/internal/moduleregistry/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/repository"
)

// QueryService 承接读组业务编排。
type QueryService struct {
	modules   *repository.ModuleStore
	releases  *repository.ReleaseStore
	bindings  *repository.BindingStore
}

// NewQueryService 构造 QueryService。
func NewQueryService(modules *repository.ModuleStore, releases *repository.ReleaseStore, bindings *repository.BindingStore) *QueryService {
	return &QueryService{modules: modules, releases: releases, bindings: bindings}
}

// ListModules 承接 ModuleListRead。
//
// 列表读取至少承接（§5.7）：
// name / description / status / latest_release / product_bind_count / repository_bind_count
//
// 编排流程：
//  1. 从 modules 表按筛选条件读取核心字段
//  2. 对每个模块，分别读取 latest_release / product_bind_count / repository_bind_count
//
// 注：当前实现按模块逐条聚合统计。phase02 数据规模小，无需引入聚合 SQL 优化。
// 后续若性能需要，可改为单条 SQL with LEFT JOIN + GROUP BY 聚合，但接口契约不变。
func (s *QueryService) ListModules(ctx context.Context, q moduleregistry.ListQuery) ([]moduleregistry.ModuleListItem, error) {
	statusFilter := ""
	if q.StatusFilter != "" && q.StatusFilter != "all" {
		statusFilter = string(q.StatusFilter)
	}

	modules, err := s.modules.List(ctx, q.QueryText, statusFilter)
	if err != nil {
		return nil, err
	}

	items := make([]moduleregistry.ModuleListItem, 0, len(modules))
	for _, m := range modules {
		latest, err := s.releases.GetLatestVersionByModule(ctx, m.ID)
		if err != nil {
			return nil, fmt.Errorf("read latest release for module %s: %w", m.ID, err)
		}
		productCount, err := s.modules.CountProductBindings(ctx, m.ID)
		if err != nil {
			return nil, fmt.Errorf("count product bindings for module %s: %w", m.ID, err)
		}
		repoCount, err := s.modules.CountRepositoryBindings(ctx, m.ID)
		if err != nil {
			return nil, fmt.Errorf("count repository bindings for module %s: %w", m.ID, err)
		}

		items = append(items, moduleregistry.ModuleListItem{
			ID:                  m.ID,
			Name:                m.Name,
			Description:         m.Description,
			Status:              m.Status,
			LatestRelease:       latest,
			ProductBindCount:    productCount,
			RepositoryBindCount: repoCount,
		})
	}
	return items, nil
}

// GetModuleDetail 承接 ModuleDetailRead。
//
// 详情读取统一承接（§5.7）：模块核心字段、版本列表、产品绑定、仓库映射与相关 Decision 入口。
//
// 编排流程：
//  1. 读取模块核心字段，不存在返回 ErrModuleNotFound
//  2. 读取版本列表（按 released_at DESC）
//  3. 读取产品绑定列表（附带 product name）
//  4. 读取仓库映射列表（附带 repository name）
//  5. 内嵌读取 Decision 关联（§6.3 不设独立读接口，直接在详情编排中组装）
//
// Decision 读取的 SQL 物理落在 repository/binding_store.go，但业务编排归属本方法，
// 不构成独立 Read 接口组或独立 handler/service 文件。
func (s *QueryService) GetModuleDetail(ctx context.Context, moduleID string) (*moduleregistry.ModuleDetail, error) {
	// 校验 ID 格式，防止无效 ID 打到数据库
	if err := moduleregistry.ValidateModuleID(moduleID); err != nil {
		return nil, err
	}

	m, err := s.modules.GetByID(ctx, moduleID)
	if err != nil {
		return nil, err
	}

	releases, err := s.releases.ListByModule(ctx, moduleID)
	if err != nil {
		return nil, fmt.Errorf("list releases: %w", err)
	}

	productBindings, err := s.bindings.ListProductBindingsByModule(ctx, moduleID)
	if err != nil {
		return nil, fmt.Errorf("list product bindings: %w", err)
	}

	repoMappings, err := s.bindings.ListRepositoryMappingsByModule(ctx, moduleID)
	if err != nil {
		return nil, fmt.Errorf("list repository mappings: %w", err)
	}

	// Decision 附属读取 — 内嵌于此，不设独立读接口（§6.3）
	decisionLinks, err := s.bindings.ListDecisionLinksByModule(ctx, moduleID)
	if err != nil {
		return nil, fmt.Errorf("list decision links: %w", err)
	}

	// 确保切片非 nil，便于前端处理（nil 切片序列化为 null，空切片序列化为 []）
	if releases == nil {
		releases = []moduleregistry.Release{}
	}
	if productBindings == nil {
		productBindings = []moduleregistry.ProductBinding{}
	}
	if repoMappings == nil {
		repoMappings = []moduleregistry.RepositoryMapping{}
	}
	if decisionLinks == nil {
		decisionLinks = []moduleregistry.DecisionLink{}
	}

	return &moduleregistry.ModuleDetail{
		Module:             *m,
		Releases:           releases,
		ProductBindings:    productBindings,
		RepositoryMappings: repoMappings,
		DecisionLinks:      decisionLinks,
	}, nil
}
