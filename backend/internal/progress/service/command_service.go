// Package service — progress 写入编排层。
//
// 单一 CommandService 承接 Progress 写路径编排（phase15-04 冻结）。
// 对齐 phase15-03 冻结校验执行序 6 步（报第一个错误）：
//   Create：envelope 前置（INVALID_REPOSITORY_ID → INVALID_OCCURRED_AT）
//         → V1a INVALID_WORKFLOW_TYPE → V1b INVALID_EVENT_KIND
//         → V7 EVENT_KIND_NOT_ALLOWED → task_key 矩阵分支
//         （TASK_KEY_REQUIRED → TASK_KEY_FORMAT_INVALID）
//         → V9 文本顺序（INVALID_TITLE → INVALID_DETAIL
//            → INVALID_EVIDENCE_REF → INVALID_SOURCE）
//         → repository 存在性（[REPOSITORY_NOT_FOUND] InvalidArgument）
//   Delete：id UUID 格式层（INVALID_PROGRESS_EVENT_ID）→ 定位删除（NotFound）
//
// 全部校验统一经根包 ValidateCreateProgressEventInput 执行（含 task_key
// trim 就地归一）；repository 存在性经 candidate.RepositoryReader 承接
// （DP-2 裁决），错误语义包装在本层。
//
// 文件落点：backend/internal/progress/service/command_service.go
package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/psco/backend/internal/progress"
	"github.com/psco/backend/internal/progress/candidate"
	"github.com/psco/backend/internal/progress/repository"
)

// 稳定错误码（phase15-04 新增 ID 格式码；Delete 以事件 id 定位，无 repository_id 输入）。
const errCodeInvalidProgressEventID = "INVALID_PROGRESS_EVENT_ID"

// CommandService 承接 Progress 写路径编排。
//
// 依赖通过 platform 装配点注入：
//   - repositoryReader：跨模块 repository 存在性事实查询（candidate 子包隔离）
//   - store：progress_events 单表持久化承接位
type CommandService struct {
	repositoryReader *candidate.RepositoryReader
	store            *repository.ProgressEventStore
}

// NewCommandService 构造 CommandService。
func NewCommandService(repositoryReader *candidate.RepositoryReader, store *repository.ProgressEventStore) *CommandService {
	return &CommandService{
		repositoryReader: repositoryReader,
		store:            store,
	}
}

// CreateProgressEvent 编排推进事件创建。
//
// 编排顺序（phase15-03 执行序 6 步冻结，报第一个错误）：
//  1. 根包统一校验（envelope 前置 → V1-V9；通过时 input.TaskKey 已就地
//     trim 归一，标识符 trim 后持久化）
//  2. repository 存在性（写入目标外键引用语义：不存在 → InvalidArgument
//     [REPOSITORY_NOT_FOUND]，沿 standard "target 不存在归 InvalidArgument" 模式）
//  3. 持久化（DB FK RESTRICT 为存储层兜底，非校验承接位）
func (s *CommandService) CreateProgressEvent(ctx context.Context, input *progress.CreateProgressEventInput) (*progress.ProgressEventReadResult, error) {
	if err := progress.ValidateCreateProgressEventInput(input); err != nil {
		return nil, err
	}

	exists, err := s.repositoryReader.RepositoryExists(ctx, input.RepositoryID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("%w: [REPOSITORY_NOT_FOUND] repository %s does not exist",
			progress.ErrInvalidInput, input.RepositoryID)
	}

	result, err := s.store.Insert(ctx, *input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteProgressEvent 编排推进事件整条删除（误录修正；无软删除、无 Update）。
// id 非法 UUID → ErrInvalidInput [INVALID_PROGRESS_EVENT_ID]；
// 不存在 → ErrProgressEventNotFound（透传 store 定位结果）。
func (s *CommandService) DeleteProgressEvent(ctx context.Context, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("%w: [%s] progress event id %q is not a valid UUID",
			progress.ErrInvalidInput, errCodeInvalidProgressEventID, id)
	}
	return s.store.DeleteByID(ctx, id)
}
