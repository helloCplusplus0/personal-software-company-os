// Package progress 承载项目推进时间轴（三轨 append-only 事件流）结构化写读能力。
//
// 分层语义（对齐 phase15-04 已冻结结论）：
//   - connect/     入口层：Connect handler，proto request 解包 → service 调用 → proto response 组装
//   - service/     业务编排层：写读编排、校验触发与 repository 存在性包装、派生摘要计算
//   - repository/  持久化层：PostgreSQL 读写（progress_events 单表，三键链排序唯一执行位）
//   - candidate/   外部连接层：跨模块 repository 存在性校验 reader（progress 自拥有）
//
// 根包另含 derive.go：派生摘要纯函数（progress 特有；零 I/O、零时间函数，可纯单元测试）。
//
// 本文件定义跨层共享的领域模型与受控枚举。
// 约束：消息结构从 proto/psco/progress/v1/progress.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在本文件 import 生成 pb 包（proto ↔ domain 转换收敛在 connect 层）。
package progress

import "time"

// ============================================================================
// 受控枚举
// ============================================================================

// WorkflowType 三轨 workflow（受控枚举，对齐 DDL workflow_type CHECK）。
type WorkflowType string

const (
	WorkflowTypePhase WorkflowType = "phase"
	WorkflowTypeAudit WorkflowType = "audit"
	WorkflowTypeFix   WorkflowType = "fix"
)

// EventKind 事件类型（受控枚举，对齐 DDL event_kind CHECK）。
type EventKind string

const (
	EventKindPhaseStarted   EventKind = "phase_started"
	EventKindPhaseCompleted EventKind = "phase_completed"
	EventKindTaskCompleted  EventKind = "task_completed"
	EventKindNote           EventKind = "note"
)

// ProgressSource 事件来源（受控枚举；本阶段创建入口仅 manual）。
type ProgressSource string

const (
	ProgressSourceManual ProgressSource = "manual"
	ProgressSourceGit    ProgressSource = "git"
	ProgressSourceAgent  ProgressSource = "agent"
)

// ============================================================================
// 读模型
// ============================================================================

// ProgressEventReadResult 推进事件读取结果（progress_events 表投影）。
type ProgressEventReadResult struct {
	ID           string         `json:"id"`
	RepositoryID string         `json:"repository_id"`
	WorkflowType WorkflowType   `json:"workflow_type"`
	EventKind    EventKind      `json:"event_kind"`
	TaskKey      string         `json:"task_key"`
	Title        string         `json:"title"`
	Detail       string         `json:"detail"`
	EvidenceRef  string         `json:"evidence_ref"`
	Source       ProgressSource `json:"source"`
	OccurredAt   time.Time      `json:"occurred_at"`
	CreatedAt    time.Time      `json:"created_at"`
}

// ProgressSummary 派生摘要（与 BriefProgress 四字段 1:1；空态恒构造零值）。
type ProgressSummary struct {
	CurrentPhaseKey     string                    `json:"current_phase_key"`
	CurrentPhaseLabel   string                    `json:"current_phase_label"`
	LatestTaskCompleted *ProgressEventReadResult  `json:"latest_task_completed"` // 可空：无任务完成时 nil
	RecentEvents        []ProgressEventReadResult `json:"recent_events"`        // min(10, len)；恒非 nil（空态为空切片）
}

// ============================================================================
// 写模型
// ============================================================================

// CreateProgressEventInput 创建事件输入（source 已由 service 归一 manual）。
type CreateProgressEventInput struct {
	RepositoryID string         `json:"repository_id"`
	WorkflowType WorkflowType   `json:"workflow_type"`
	EventKind    EventKind      `json:"event_kind"`
	TaskKey      string         `json:"task_key"`
	Title        string         `json:"title"`
	Detail       string         `json:"detail"`
	EvidenceRef  string         `json:"evidence_ref"`
	Source       ProgressSource `json:"source"`
	OccurredAt   time.Time      `json:"occurred_at"`
}
