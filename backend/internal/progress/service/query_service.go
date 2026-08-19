// Package service — progress 只读编排层。
//
// 单一 QueryService 承接 Progress 读路径编排（phase15-04 冻结）：
//   - ListProgressEvents：UUID 校验 → repository 存在性（读锚点 NotFound 语义）
//     → store 单一查询（三键链倒序，Go 侧不重排）
//   - GetProgressSummary：store 单查询无过滤 → DeriveProgressSummary 纯函数；
//     本方法是 projectcontext candidate.ProgressReader 的实现位
//     （brief progress 块与 List RPC 同源同派生）
//
// 读路径 store 错误已由 repository 层统一 wrap ErrProgressReadFailed，
// service 直接透传。
//
// 文件落点：backend/internal/progress/service/query_service.go
package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/psco/backend/internal/progress"
	"github.com/psco/backend/internal/progress/candidate"
	"github.com/psco/backend/internal/progress/repository"
)

// 稳定错误码（与根包 validate.go envelope 前置码同值；根包常量私有，
// service 层承接 List 路径的格式层校验）。
const errCodeInvalidRepositoryID = "INVALID_REPOSITORY_ID"

// QueryService 承接 Progress 只读编排。
//
// 依赖通过 platform 装配点注入：
//   - store：progress_events 单表读取
//   - repositoryReader：repository 存在性事实查询（DP-2 承接位）
type QueryService struct {
	store            *repository.ProgressEventStore
	repositoryReader *candidate.RepositoryReader
}

// NewQueryService 构造 QueryService。
func NewQueryService(store *repository.ProgressEventStore, repositoryReader *candidate.RepositoryReader) *QueryService {
	return &QueryService{
		store:            store,
		repositoryReader: repositoryReader,
	}
}

// ListProgressEvents 编排完整事件流读取（三键链倒序，不分页）。
//
// 失败语义（phase15-04 RPC 1 三要素冻结）：
//   - repository_id 非法 UUID → ErrInvalidInput [INVALID_REPOSITORY_ID]
//   - 仓库不存在 → progress.ErrRepositoryNotFound（读锚点 NotFound 语义，沿 GetProjectBrief）
//   - 无事件 / 过滤后空 → 空列表（非错误）
// workflowType 非 nil 时单轨过滤（UNSPECIFIED → nil 不过滤的归一在 connect 解包层承接）。
func (s *QueryService) ListProgressEvents(ctx context.Context, repositoryID string, workflowType *progress.WorkflowType) ([]progress.ProgressEventReadResult, error) {
	if err := validateRepositoryUUID(repositoryID); err != nil {
		return nil, err
	}

	exists, err := s.repositoryReader.RepositoryExists(ctx, repositoryID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, progress.ErrRepositoryNotFound
	}

	return s.store.ListByRepository(ctx, repositoryID, workflowType)
}

// GetProgressSummary 编排进度派生摘要计算。
//
// 本方法是 projectcontext candidate.ProgressReader 的实现位：
// store 单一查询（无 workflow_type 过滤）取全量事件 → 根包纯函数派生四字段，
// 与 ListProgressEvents RPC 同源同派生，无第二套派生路径。
//
// 失败语义（phase15-04 ProgressReader 接口冻结）：
//   - 不做 repository 存在性校验（brief 主流程 ReadRepository 已先行承接
//     仓库存在性；EXISTS 语义由查询天然承接：无行 → 空集 → 零值摘要）
//   - 读取失败 → ErrProgressReadFailed（store 层 wrap 透传）
//   - 仓库无事件 → 零值摘要 + 空 RecentEvents（非错误，derive 纯函数空态恒构造）
func (s *QueryService) GetProgressSummary(ctx context.Context, repositoryID string) (progress.ProgressSummary, error) {
	events, err := s.store.ListByRepository(ctx, repositoryID, nil)
	if err != nil {
		return progress.ProgressSummary{}, err
	}
	return progress.DeriveProgressSummary(events), nil
}

// validateRepositoryUUID 校验 repository_id 为合法 UUID（List 读锚点格式层，
// 错误码与根包 envelope 前置同值；存在性归 candidate 承接位）。
func validateRepositoryUUID(repositoryID string) error {
	if _, err := uuid.Parse(repositoryID); err != nil {
		return fmt.Errorf("%w: [%s] repository_id %q is not a valid UUID",
			progress.ErrInvalidInput, errCodeInvalidRepositoryID, repositoryID)
	}
	return nil
}
