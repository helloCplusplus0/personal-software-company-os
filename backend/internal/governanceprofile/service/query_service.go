// Package service — governanceprofile 只读编排层。
//
// 单一 QueryService 承接 GetGovernanceProfile 结构化读取编排。
// 对齐 phase13-08 已冻结的读取主线：
//   - 读取结果区分结构化字段、结构化摘要与正文回源能力
//   - 本服务不把 markdown 全文作为 canonical stored field 返回
//   - 本服务不是完整 agent brief 入口（phase13-10 另行承接）
//
// 文件落点：backend/internal/governanceprofile/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/governanceprofile"
	"github.com/psco/backend/internal/governanceprofile/candidate"
	"github.com/psco/backend/internal/governanceprofile/repository"
)

// QueryService 承接治理画像结构化读取编排。
//
// 依赖通过 platform 装配点注入：
//   - repositoryReader：跨模块仓库存在性前提校验（candidate 子包隔离）
//   - profileStore：治理画像聚合读取
type QueryService struct {
	repositoryReader *candidate.RepositoryReader
	profileStore     *repository.ProfileStore
}

// NewQueryService 构造 QueryService。
func NewQueryService(repositoryReader *candidate.RepositoryReader, profileStore *repository.ProfileStore) *QueryService {
	return &QueryService{
		repositoryReader: repositoryReader,
		profileStore:     profileStore,
	}
}

// GetGovernanceProfile 编排治理画像结构化读取。
//
// 编排顺序：
//  1. 校验目标仓库已在 PSCO 登记（不存在则返回 ErrRepositoryNotFound）
//  2. 聚合读取主记录与两组 bindings（画像未创建则返回 ErrGovernanceProfileNotFound）
//
// 失败语义：
//   - repository 不存在 → ErrRepositoryNotFound（NotFound）
//   - 画像未创建       → ErrGovernanceProfileNotFound（NotFound）
//   - 其他读取失败     → ErrGovernanceProfileReadFailed（Internal）
func (s *QueryService) GetGovernanceProfile(ctx context.Context, repositoryID string) (*governanceprofile.GovernanceProfileReadResult, error) {
	// 1. 仓库存在性前提
	if err := s.repositoryReader.EnsureRepositoryExists(ctx, repositoryID); err != nil {
		return nil, err
	}

	// 2. 聚合读取治理画像
	result, err := s.profileStore.ReadProfile(ctx, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("governanceprofile: get governance profile: %w", err)
	}
	return result, nil
}
