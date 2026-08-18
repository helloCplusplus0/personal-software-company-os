// Package service — governanceprofile 只读编排层。
//
// phase14-09 收缩：画像 RPC 编排（GetGovernanceProfile）与写路径已随画像
// Connect handler 退役；单一 QueryService 收敛为 ReadProfileCore 主表
// 轻量读取编排，服务 projectcontext brief 内联 BriefGovernanceProfile 装配
// （经 candidate.GovernanceProfileReader 接口注入）。
//
// 文件落点：backend/internal/governanceprofile/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/governanceprofile"
	"github.com/psco/backend/internal/governanceprofile/repository"
)

// QueryService 承接治理画像主记录轻量读取编排。
//
// 依赖通过 platform 装配点注入：
//   - profileStore：治理画像主表只读
type QueryService struct {
	profileStore *repository.ProfileStore
}

// NewQueryService 构造 QueryService。
func NewQueryService(profileStore *repository.ProfileStore) *QueryService {
	return &QueryService{
		profileStore: profileStore,
	}
}

// ReadProfileCore 编排治理画像主记录核心字段读取（不含已退役的两组 bindings）。
//
// 失败语义：
//   - 画像未创建       → ErrGovernanceProfileNotFound（NotFound）
//   - 其他读取失败     → ErrGovernanceProfileReadFailed（Internal）
func (s *QueryService) ReadProfileCore(ctx context.Context, repositoryID string) (*governanceprofile.GovernanceProfileCoreReadResult, error) {
	result, err := s.profileStore.ReadProfile(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("governanceprofile: read profile core: %w", err)
	}
	return result, nil
}
