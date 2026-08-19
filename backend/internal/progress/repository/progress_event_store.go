// Package repository — progress 数据访问层。
//
// 本文件承接 progress_events 单表的 PostgreSQL 持久化（phase15-03 冻结）：
//   - ListByRepository：单一读取语句（三键链排序唯一执行位，Go 侧不重排）
//   - Insert：append-only 写入（RETURNING 完整行）
//   - DeleteByID：整条删除（append-only 唯一修正路径，无 Update 语义）
//
// 编解码语义（对齐 0013 迁移列集）：
//   - workflow_type / event_kind / source：受控小写字符串直存
//   - task_key / detail / evidence_ref 可空 TEXT：写侧空串 → NULL，
//     读侧 NULL → 空串（NULL 与空串读取等价，转换收敛在本层）
//
// 本层只做 SQL 与编解码，无业务判断；校验收敛在根包 validate + service 层，
// repository 存在性校验承接位在 candidate 子包（phase15-04 DP-2 裁决）。
//
// 文件落点：backend/internal/progress/repository/progress_event_store.go
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/progress"
)

// 单一读取语句形态（phase15-03 冻结逐字，仅可选过滤子句按调用方拼接）：
// 排序唯一执行位在 SQL ORDER BY；List RPC / brief 派生共用同一查询。
const listProgressEventsSQL = `
	SELECT id::text, repository_id::text, workflow_type, event_kind, task_key, title,
	       detail, evidence_ref, source, occurred_at, created_at
	  FROM progress_events
	 WHERE repository_id = $1
	 ORDER BY occurred_at DESC, created_at DESC, id DESC
`

const listProgressEventsByWorkflowSQL = `
	SELECT id::text, repository_id::text, workflow_type, event_kind, task_key, title,
	       detail, evidence_ref, source, occurred_at, created_at
	  FROM progress_events
	 WHERE repository_id = $1
	   AND workflow_type = $2
	 ORDER BY occurred_at DESC, created_at DESC, id DESC
`

// ProgressEventStore 承接 progress_events 表的持久化操作。
type ProgressEventStore struct {
	pool *pgxpool.Pool
}

// NewProgressEventStore 构造 ProgressEventStore。
func NewProgressEventStore(pool *pgxpool.Pool) *ProgressEventStore {
	return &ProgressEventStore{pool: pool}
}

// ListByRepository 读取仓库完整事件流（三键链倒序：occurred_at DESC,
// created_at DESC, id DESC；排序唯一执行位在 SQL，Go 侧不重排）。
// workflowType 非 nil 时追加单轨过滤；nil = 三轨全量。
// 空结果返回空切片非 nil；读失败 wrap ErrProgressReadFailed。
func (s *ProgressEventStore) ListByRepository(ctx context.Context, repositoryID string, workflowType *progress.WorkflowType) ([]progress.ProgressEventReadResult, error) {
	// 可选过滤子句按调用方拼接（phase15-03 冻结）：nil = 三轨全量单条件语句
	query, args := listProgressEventsSQL, []any{repositoryID}
	if workflowType != nil {
		query = listProgressEventsByWorkflowSQL
		args = append(args, string(*workflowType))
	}

	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: list progress events: %w", progress.ErrProgressReadFailed, err)
	}
	defer rows.Close()

	events := []progress.ProgressEventReadResult{}
	for rows.Next() {
		event, err := scanProgressEvent(rows)
		if err != nil {
			return nil, fmt.Errorf("%w: scan progress event: %w", progress.ErrProgressReadFailed, err)
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate progress events: %w", progress.ErrProgressReadFailed, err)
	}
	return events, nil
}

// Insert 插入推进事件并返回完整读取结果（RETURNING 含服务端生成的
// id / created_at）。可空文本三字段空串写为 NULL；写失败 wrap
// ErrProgressWriteFailed（repository 存在性已由 service 层先行校验，
// FK RESTRICT 为存储层兜底，非校验承接位）。
func (s *ProgressEventStore) Insert(ctx context.Context, input progress.CreateProgressEventInput) (progress.ProgressEventReadResult, error) {
	// 可空文本三字段：空串 → NULL（未填语义；读侧 NULL 归空串，双向等价）
	var taskKey, detail, evidenceRef *string
	if input.TaskKey != "" {
		taskKey = &input.TaskKey
	}
	if input.Detail != "" {
		detail = &input.Detail
	}
	if input.EvidenceRef != "" {
		evidenceRef = &input.EvidenceRef
	}

	event, err := scanProgressEvent(s.pool.QueryRow(ctx, `
		INSERT INTO progress_events (repository_id, workflow_type, event_kind, task_key, title, detail, evidence_ref, source, occurred_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id::text, repository_id::text, workflow_type, event_kind, task_key, title, detail, evidence_ref, source, occurred_at, created_at
	`, input.RepositoryID, string(input.WorkflowType), string(input.EventKind), taskKey, input.Title, detail, evidenceRef, string(input.Source), input.OccurredAt))
	if err != nil {
		return progress.ProgressEventReadResult{}, fmt.Errorf("%w: insert progress event: %w", progress.ErrProgressWriteFailed, err)
	}
	return event, nil
}

// DeleteByID 按 id 整条删除推进事件（append-only 唯一修正路径）。
// 未找到时返回 progress.ErrProgressEventNotFound；删除失败 wrap
// ErrProgressWriteFailed。
func (s *ProgressEventStore) DeleteByID(ctx context.Context, id string) error {
	tag, err := s.pool.Exec(ctx, `
		DELETE FROM progress_events WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("%w: delete progress event: %w", progress.ErrProgressWriteFailed, err)
	}
	if tag.RowsAffected() == 0 {
		return progress.ErrProgressEventNotFound
	}
	return nil
}

// rowScanner 统一 pgx.Row / pgx.Rows 的 Scan 承接位。
type rowScanner interface {
	Scan(dest ...any) error
}

// scanProgressEvent 将 progress_events 表单行解码为 ProgressEventReadResult
// （可空 TEXT 三列 NULL → 空串；受控枚举列直存 string）。
func scanProgressEvent(row rowScanner) (progress.ProgressEventReadResult, error) {
	var event progress.ProgressEventReadResult
	var taskKey, detail, evidenceRef *string
	if err := row.Scan(
		&event.ID, &event.RepositoryID, &event.WorkflowType, &event.EventKind,
		&taskKey, &event.Title, &detail, &evidenceRef, &event.Source,
		&event.OccurredAt, &event.CreatedAt,
	); err != nil {
		return progress.ProgressEventReadResult{}, err
	}
	if taskKey != nil {
		event.TaskKey = *taskKey
	}
	if detail != nil {
		event.Detail = *detail
	}
	if evidenceRef != nil {
		event.EvidenceRef = *evidenceRef
	}
	return event, nil
}
