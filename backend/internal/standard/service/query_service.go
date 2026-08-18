// Package service — standard 只读编排层。
//
// 单一 QueryService 承接 Standard 读路径编排。
// 对齐 phase14-04 / phase14-07 已冻结的读取主线：
//   - 全量列表与 revision 回看不分页（单用户系统数量级）
//   - 读路径 store 错误已由 repository 层统一 wrap ErrStandardReadFailed，
//     service 直接透传
//   - ListStandardsByRepository 是 projectcontext candidate.StandardReader
//     的实现位（phase14-04 冻结，brief 反查主线）
//
// 文件落点：backend/internal/standard/service/query_service.go
package service

import (
	"context"

	"github.com/psco/backend/internal/standard"
	"github.com/psco/backend/internal/standard/repository"
)

// QueryService 承接 Standard 只读编排。
//
// 依赖通过 platform 装配点注入：
//   - store：三表聚合读取
type QueryService struct {
	store *repository.StandardStore
}

// NewQueryService 构造 QueryService（只依赖 store）。
func NewQueryService(store *repository.StandardStore) *QueryService {
	return &QueryService{store: store}
}

// ListStandards 编排全量规范列表读取（按 updated_at DESC，不分页）。
func (s *QueryService) ListStandards(ctx context.Context) ([]standard.StandardReadResult, error) {
	return s.store.ListStandards(ctx)
}

// GetStandard 编排规范详情读取（主记录 + 绑定集合）。
// 不存在时返回 ErrStandardNotFound。
func (s *QueryService) GetStandard(ctx context.Context, standardID string) (*standard.StandardReadResult, []standard.StandardBindingReadResult, error) {
	if err := validateUUID("standard", standardID); err != nil {
		return nil, nil, err
	}

	result, err := s.store.GetStandardByID(ctx, standardID)
	if err != nil {
		return nil, nil, err
	}

	bindings, err := s.store.ListBindingsByStandardID(ctx, standardID)
	if err != nil {
		return nil, nil, err
	}
	return result, bindings, nil
}

// ListStandardRevisions 编排 revision 回看（按 created_at DESC，不分页）。
// 先读主记录承接 NotFound 语义，再读留痕列表。
func (s *QueryService) ListStandardRevisions(ctx context.Context, standardID string) ([]standard.StandardRevisionReadResult, error) {
	if err := validateUUID("standard", standardID); err != nil {
		return nil, err
	}

	if _, err := s.store.GetStandardByID(ctx, standardID); err != nil {
		return nil, err
	}

	return s.store.ListRevisions(ctx, standardID)
}

// ListStandardsByRepository 按 repository 绑定关系反查关联规范。
//
// 本方法是 projectcontext candidate.StandardReader 的实现位
// （phase14-04 冻结：GetProjectBrief.standards[] 经本反查主线取数）。
// repositoryID uuid 校验；空列表非错误。
func (s *QueryService) ListStandardsByRepository(ctx context.Context, repositoryID string) ([]standard.StandardReadResult, error) {
	if err := validateUUID("repository", repositoryID); err != nil {
		return nil, err
	}
	return s.store.ListStandardsByRepository(ctx, repositoryID)
}
