// 创建事件输入校验（V1-V9 + envelope 前置，phase15-03 冻结规则）。
//
// 本文件承接写路径统一执行的校验（执行序 6 步报第一个错误）：
//   1. envelope 前置：INVALID_REPOSITORY_ID（UUID 格式）→ INVALID_OCCURRED_AT（已设置）
//   2. V1a INVALID_WORKFLOW_TYPE    workflow_type 枚举（phase/audit/fix）
//   3. V1b INVALID_EVENT_KIND       event_kind 枚举（phase_started/phase_completed/task_completed/note）
//   4. V7  EVENT_KIND_NOT_ALLOWED   组合矩阵（audit/fix 轨 × phase_started/phase_completed 禁止）
//   5. V2-V6 task_key 矩阵分支     先必填（TASK_KEY_REQUIRED）后格式（TASK_KEY_FORMAT_INVALID）
//   6. V9  文本字段顺序             INVALID_TITLE → INVALID_DETAIL → INVALID_EVIDENCE_REF → INVALID_SOURCE
//
// V8（note 轨 task_key 可空）无独立错误码：合法路径不产生错误，
// 其语义由矩阵分支覆盖（note 格不走必填与格式分支）。
//
// TrimSpace 规范化边界（phase15-03 冻结）：
//   - task_key（结构化标识符）：trim 后判定并持久化 trim 后值——校验通过时
//     本函数将 input.TaskKey 就地归一为 trim 后值（判定与持久化同源承接）
//   - title / detail / evidence_ref（自由文本/引用）：trim 仅用于必填判定（title），
//     原值持久化（沿 standard name 处理模式）
//   - 长度计量：Unicode rune 计数（中文按字符而非字节计）
//
// 所有校验错误统一包装 ErrInvalidInput（InvalidArgument），错误信息格式沿
// standard 模式 `%w: [CODE] message`，格式类错误携带期望格式说明。
//
// repository 存在性校验（phase15-03 执行序第 6 步存储层语义）不在本函数范围：
// 承接位为 candidate.RepositoryReader + service 层包装（phase15-04 DP-2 裁决）。
//
// 文件落点：backend/internal/progress/validate.go
package progress

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

// 校验稳定错误码（逐字冻结自 phase15-03，前端按码定位）。
const (
	errCodeInvalidWorkflowType  = "INVALID_WORKFLOW_TYPE"
	errCodeInvalidEventKind     = "INVALID_EVENT_KIND"
	errCodeEventKindNotAllowed  = "EVENT_KIND_NOT_ALLOWED"
	errCodeTaskKeyRequired      = "TASK_KEY_REQUIRED"
	errCodeTaskKeyFormatInvalid = "TASK_KEY_FORMAT_INVALID"
	errCodeInvalidTitle         = "INVALID_TITLE"
	errCodeInvalidDetail        = "INVALID_DETAIL"
	errCodeInvalidEvidenceRef   = "INVALID_EVIDENCE_REF"
	errCodeInvalidSource        = "INVALID_SOURCE"
	errCodeInvalidRepositoryID  = "INVALID_REPOSITORY_ID"
	errCodeInvalidOccurredAt    = "INVALID_OCCURRED_AT"
)

// 校验冻结阈值（phase15-03）。
const (
	maxTitleRunes  = 200  // title 字符上限（rune 计数）
	maxDetailRunes = 2000 // detail 字符上限（rune 计数）
)

// task_key 格式正则（K-1~K-4，逐字冻结自 phase15-02/03）。
var (
	// taskKeyPhasePattern K-1：phase 轨 phase_started / phase_completed（phaseNN）。
	taskKeyPhasePattern = regexp.MustCompile(`^phase[0-9]{2,}$`)

	// taskKeyPhaseTaskPattern K-2：phase 轨 task_completed（phaseNN-MM）。
	taskKeyPhaseTaskPattern = regexp.MustCompile(`^phase[0-9]{2,}-[0-9]{2,}$`)

	// taskKeyAuditPattern K-3：audit 轨 task_completed（audit_NNN）。
	taskKeyAuditPattern = regexp.MustCompile(`^audit_[0-9]{3,}$`)

	// taskKeyFixPattern K-4：fix 轨 task_completed（fix_NNN）。
	taskKeyFixPattern = regexp.MustCompile(`^fix_[0-9]{3,}$`)
)

// ValidateCreateProgressEventInput 执行创建事件输入的统一校验。
// 执行序沿 phase15-03 冻结 6 步报第一个错误；校验通过时将 input.TaskKey
// 就地归一为 trim 后值（标识符 trim 持久化的承接位），
// title / detail / evidence_ref 保持原值（自由文本原值入库）。
func ValidateCreateProgressEventInput(input *CreateProgressEventInput) error {
	// 第 1 步：envelope 前置。
	// repository_id 合法 UUID 格式（存在性校验归 service 层 DP-2 承接位；
	// 补录历史合法，occurred_at 不做未来时间校验）。
	if _, err := uuid.Parse(input.RepositoryID); err != nil {
		return fmt.Errorf("%w: [%s] repository_id %q is not a valid UUID",
			ErrInvalidInput, errCodeInvalidRepositoryID, input.RepositoryID)
	}
	// occurred_at 已设置（用户声明发生时间，无合法零值）。
	if input.OccurredAt.IsZero() {
		return fmt.Errorf("%w: [%s] occurred_at is required and must be set",
			ErrInvalidInput, errCodeInvalidOccurredAt)
	}

	// 第 2 步：V1a workflow_type 枚举。
	switch input.WorkflowType {
	case WorkflowTypePhase, WorkflowTypeAudit, WorkflowTypeFix:
	default:
		return fmt.Errorf("%w: [%s] workflow_type %q must be one of phase/audit/fix",
			ErrInvalidInput, errCodeInvalidWorkflowType, input.WorkflowType)
	}

	// 第 3 步：V1b event_kind 枚举。
	switch input.EventKind {
	case EventKindPhaseStarted, EventKindPhaseCompleted, EventKindTaskCompleted, EventKindNote:
	default:
		return fmt.Errorf("%w: [%s] event_kind %q must be one of phase_started/phase_completed/task_completed/note",
			ErrInvalidInput, errCodeInvalidEventKind, input.EventKind)
	}

	// 第 4 步：V7 组合矩阵（先于 task_key 判定：组合非法时 task_key 无从校验）。
	if (input.WorkflowType == WorkflowTypeAudit || input.WorkflowType == WorkflowTypeFix) &&
		(input.EventKind == EventKindPhaseStarted || input.EventKind == EventKindPhaseCompleted) {
		return fmt.Errorf("%w: [%s] event_kind %q is not allowed on workflow_type %q (only task_completed/note allowed)",
			ErrInvalidInput, errCodeEventKindNotAllowed, input.EventKind, input.WorkflowType)
	}

	// 第 5 步：task_key 矩阵分支（V2-V6）：先必填后格式。
	// task_key trim 后参与校验与持久化（标识符 trim 无害规范化）。
	taskKey := strings.TrimSpace(input.TaskKey)
	input.TaskKey = taskKey
	if required, pattern := taskKeyRule(input.WorkflowType, input.EventKind); required {
		if taskKey == "" {
			return fmt.Errorf("%w: [%s] task_key is required for workflow_type %q with event_kind %q",
				ErrInvalidInput, errCodeTaskKeyRequired, input.WorkflowType, input.EventKind)
		}
		if !pattern.MatchString(taskKey) {
			return fmt.Errorf("%w: [%s] task_key %q does not match expected format %s",
				ErrInvalidInput, errCodeTaskKeyFormatInvalid, taskKey, pattern.String())
		}
	}
	// note 格（V8/K-5）：task_key 可空；若填写不强制格式——无必填与格式分支。

	// 第 6 步：V9 文本字段顺序：title → detail → evidence_ref → source。
	// title：trim 非空且 rune ≤200（空与超长同码，信息区分；trim 判定原值入库）。
	if strings.TrimSpace(input.Title) == "" {
		return fmt.Errorf("%w: [%s] title must not be empty or whitespace-only",
			ErrInvalidInput, errCodeInvalidTitle)
	}
	if utf8.RuneCountInString(input.Title) > maxTitleRunes {
		return fmt.Errorf("%w: [%s] title exceeds %d characters",
			ErrInvalidInput, errCodeInvalidTitle, maxTitleRunes)
	}
	// detail：可空，仅校验上限（原值入库）。
	if utf8.RuneCountInString(input.Detail) > maxDetailRunes {
		return fmt.Errorf("%w: [%s] detail exceeds %d characters",
			ErrInvalidInput, errCodeInvalidDetail, maxDetailRunes)
	}
	// evidence_ref：可空；非空时以 / 或 https:// 开头（正文零托管，仅导航引用）。
	if input.EvidenceRef != "" &&
		!strings.HasPrefix(input.EvidenceRef, "/") && !strings.HasPrefix(input.EvidenceRef, "https://") {
		return fmt.Errorf("%w: [%s] evidence_ref must start with %q or %q",
			ErrInvalidInput, errCodeInvalidEvidenceRef, "/", "https://")
	}
	// source：现值约束——仅 manual（git/agent 为预留枚举位，本阶段创建入口拒绝）。
	if input.Source != ProgressSourceManual {
		return fmt.Errorf("%w: [%s] source %q is not accepted yet (only manual is allowed)",
			ErrInvalidInput, errCodeInvalidSource, input.Source)
	}

	return nil
}

// taskKeyRule 返回矩阵格的 task_key 规则：必填位与格式正则（K-1~K-4）。
// note 格（K-5）可空不强制格式，返回 required=false；
// audit/fix 轨 phase 边界格已在 V7 拦截，不会进入本函数的该分支。
func taskKeyRule(workflowType WorkflowType, eventKind EventKind) (required bool, pattern *regexp.Regexp) {
	switch {
	case workflowType == WorkflowTypePhase &&
		(eventKind == EventKindPhaseStarted || eventKind == EventKindPhaseCompleted):
		return true, taskKeyPhasePattern // K-1
	case workflowType == WorkflowTypePhase && eventKind == EventKindTaskCompleted:
		return true, taskKeyPhaseTaskPattern // K-2
	case workflowType == WorkflowTypeAudit && eventKind == EventKindTaskCompleted:
		return true, taskKeyAuditPattern // K-3
	case workflowType == WorkflowTypeFix && eventKind == EventKindTaskCompleted:
		return true, taskKeyFixPattern // K-4
	default: // note 格（三轨通用，K-5）：可空、不强制格式
		return false, nil
	}
}
