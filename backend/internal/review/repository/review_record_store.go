// Package repository — Review 数据访问层。
//
// 本文件承接 review_records 单表的轻量持久化操作。
// 只服务于 next-step result 或可选 review 过程留痕，不复制完整实体快照。
//
// 文件落点：backend/internal/review/repository/review_record_store.go
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/review"
)

// ReviewRecordStore 承接 review_records 的单表持久化操作。
type ReviewRecordStore struct {
	pool *pgxpool.Pool
}

// NewReviewRecordStore 构造 ReviewRecordStore。
func NewReviewRecordStore(pool *pgxpool.Pool) *ReviewRecordStore {
	return &ReviewRecordStore{pool: pool}
}

// CreateReviewRecord 创建一条 review 记录。
func (s *ReviewRecordStore) CreateReviewRecord(ctx context.Context, input review.SubmitReviewResultInput) (*review.ReviewRecord, error) {
	query := `
		INSERT INTO review_records (
			review_kind, result_kind, decision_id, target_type, target_id,
			summary_text, started_at, completed_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, review_kind, result_kind, decision_id, target_type, target_id,
		          summary_text, started_at, completed_at, created_at
	`

	record := &review.ReviewRecord{}
	// decision_id / target_type / target_id 在数据库中为 NULL 可空列，
	// 但 review.ReviewRecord 的对应字段为 string（非指针），
	// pgx 无法将 NULL 直接扫描到 string 字段地址。
	// 使用中间 *string 变量承接，扫描后再解引用赋值。
	var decisionID, targetType, targetID *string
	err := s.pool.QueryRow(ctx, query,
		string(input.ReviewKind),
		string(input.ResultKind),
		nullIfEmpty(input.DecisionID),
		nullIfEmpty(input.TargetType),
		nullIfEmpty(input.TargetID),
		input.SummaryText,
		input.StartedAt,
		input.CompletedAt,
	).Scan(
		&record.ID,
		&record.ReviewKind,
		&record.ResultKind,
		&decisionID,
		&targetType,
		&targetID,
		&record.SummaryText,
		&record.StartedAt,
		&record.CompletedAt,
		&record.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("review: create review record: %w", err)
	}

	if decisionID != nil {
		record.DecisionID = *decisionID
	}
	if targetType != nil {
		record.TargetType = *targetType
	}
	if targetID != nil {
		record.TargetID = *targetID
	}

	return record, nil
}

// nullIfEmpty 返回 nil 指针当字符串为空，否则返回字符串指针。
func nullIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}