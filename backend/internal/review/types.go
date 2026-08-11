// Package review 承载 Review 后端模块的全部业务实现。
//
// 分层语义（对齐 phase08-07 已冻结结论）：
//   - service/    业务编排层：QueryService 组合既有 canonical service，CommandService 承接最小 review result sink
//   - repository/ 数据访问层：review_records 单表轻量过程记录
//   - connect/    Connect transport 实现层
//
// 本文件定义跨层共享的 domain types。
//
// 文件落点：backend/internal/review/types.go
package review

import "time"

// ReviewKind review 会话类型，对齐 .proto ReviewKind 枚举。
type ReviewKind string

const (
	ReviewKindDaily  ReviewKind = "daily"
	ReviewKindWeekly ReviewKind = "weekly"
)

// ReviewResultKind review 结果类型，对齐 .proto ReviewResultKind 枚举。
type ReviewResultKind string

const (
	ReviewResultKindDecisionHandoff ReviewResultKind = "decision_handoff"
	ReviewResultKindEntityHandoff   ReviewResultKind = "entity_handoff"
	ReviewResultKindNextStepResult  ReviewResultKind = "next_step_result"
)

// ReviewRecord review 记录的最小 domain 模型。
// 对齐 phase08-04 / 08-07 冻结的最小字段集合。
type ReviewRecord struct {
	ID           string           `json:"id"`
	ReviewKind   ReviewKind       `json:"review_kind"`
	ResultKind   ReviewResultKind `json:"result_kind"`
	DecisionID   string           `json:"decision_id,omitempty"`
	TargetType   string           `json:"target_type,omitempty"`
	TargetID     string           `json:"target_id,omitempty"`
	SummaryText  string           `json:"summary_text"`
	StartedAt    time.Time        `json:"started_at"`
	CompletedAt  time.Time        `json:"completed_at"`
	CreatedAt    time.Time        `json:"created_at"`
}

// SubmitReviewResultInput SubmitReviewResult 的最小输入。
// 对齐 phase08-07 冻结的 SubmitReviewResultRequest 字段。
type SubmitReviewResultInput struct {
	ReviewKind  ReviewKind       `json:"review_kind"`
	ResultKind  ReviewResultKind `json:"result_kind"`
	DecisionID  string           `json:"decision_id,omitempty"`
	TargetType  string           `json:"target_type,omitempty"`
	TargetID    string           `json:"target_id,omitempty"`
	SummaryText string           `json:"summary_text"`
	StartedAt   time.Time        `json:"started_at"`
	CompletedAt time.Time        `json:"completed_at"`
}

// SubmitReviewResultOutput SubmitReviewResult 的最小输出。
// 对齐 phase08-07 冻结的 SubmitReviewResultResponse 字段。
type SubmitReviewResultOutput struct {
	ReviewRecordID string           `json:"review_record_id,omitempty"`
	ResultKind     ReviewResultKind `json:"result_kind"`
	DecisionID     string           `json:"decision_id,omitempty"`
	TargetType     string           `json:"target_type,omitempty"`
	TargetID       string           `json:"target_id,omitempty"`
}