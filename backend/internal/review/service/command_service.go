// Package service — Review 业务编排层。
//
// CommandService 只承接 SubmitReviewResult 这一类最小 review result sink。
// 实体写入继续直走既有 canonical command service，不得在 review 模块里复制第二套写路径。
//
// 文件落点：backend/internal/review/service/command_service.go
package service

import (
	"context"
	"fmt"

	"github.com/psco/backend/internal/review"
	"github.com/psco/backend/internal/review/repository"
)

// CommandService 承接 Review 的最小写路径。
//
// 只承接 next-step result 和可选 review 过程留痕。
// decision handoff / entity handoff 路径允许完全不调用此服务。
type CommandService struct {
	reviewRecordStore *repository.ReviewRecordStore
}

// NewCommandService 构造 CommandService。
func NewCommandService(reviewRecordStore *repository.ReviewRecordStore) *CommandService {
	return &CommandService{reviewRecordStore: reviewRecordStore}
}

// SubmitReviewResult 提交 review 结果。
//
// 只允许承接 review 流程语义字段，不得承接完整 DecisionDraftInput、
// ProductBindingInput、RepositoryMappingInput 等实体写入。
func (s *CommandService) SubmitReviewResult(ctx context.Context, input review.SubmitReviewResultInput) (*review.SubmitReviewResultOutput, error) {
	record, err := s.reviewRecordStore.CreateReviewRecord(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("review: submit review result: %w", err)
	}

	return &review.SubmitReviewResultOutput{
		ReviewRecordID: record.ID,
		ResultKind:     record.ResultKind,
		DecisionID:     record.DecisionID,
		TargetType:     record.TargetType,
		TargetID:       record.TargetID,
	}, nil
}