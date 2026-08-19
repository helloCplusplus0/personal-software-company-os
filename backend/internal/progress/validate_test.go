// 创建事件输入校验单元测试（V1-V9 + envelope 前置 + 执行序首错 + trim 归一）。
//
// 覆盖冻结要求（phase15-06 Task 3 / phase15-03 Scenario）：
//   - 合法矩阵正例（phase 轨 4 格 + audit/fix 轨各 2 格；audit/fix 轨的
//     phase_started/phase_completed 为 V7 禁止格，作为反例覆盖——12 格全判定）
//   - 11 个稳定错误码逐码覆盖（每码至少 1 个触发用例）
//   - 执行序首错断言（构造多错输入断言只报第一个错误码）
//   - TrimSpace 边界（task_key trim 后持久化；title trim 判定原值入库）
//   - 断言按错误码（errors.Is / strings.Contains "[CODE]"）非错误信息全文
//
// 文件落点：backend/internal/progress/validate_test.go
package progress

import (
	"errors"
	"strings"
	"testing"
	"time"
)

// validInput 构造合法基线输入（phase × phase_started，K-1 合法 task_key）。
func validInput() CreateProgressEventInput {
	return CreateProgressEventInput{
		RepositoryID: "11111111-1111-1111-1111-111111111111",
		WorkflowType: WorkflowTypePhase,
		EventKind:    EventKindPhaseStarted,
		TaskKey:      "phase15",
		Title:        "phase15 开始",
		Source:       ProgressSourceManual,
		OccurredAt:   time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC),
	}
}

// TestValidateCreateProgressEventInput 表驱动覆盖合法矩阵与 11 错误码。
func TestValidateCreateProgressEventInput(t *testing.T) {
	tests := []struct {
		name        string
		mutate      func(*CreateProgressEventInput)
		wantErrCode string // 空串表示合法用例
	}{
		// --- 合法矩阵正例（8 合法格；audit/fix 轨 phase 边界格见 V7 反例）---
		{
			name: "矩阵合法：phase × phase_started（K-1 phaseNN）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypePhase
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "phase15"
			},
		},
		{
			name: "矩阵合法：phase × phase_completed（K-1 phaseNN）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypePhase
				in.EventKind = EventKindPhaseCompleted
				in.TaskKey = "phase15"
			},
		},
		{
			name: "矩阵合法：phase × task_completed（K-2 phaseNN-MM）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypePhase
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "phase15-06"
			},
		},
		{
			name: "矩阵合法：phase × note（task_key 可空）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypePhase
				in.EventKind = EventKindNote
				in.TaskKey = ""
			},
		},
		{
			name: "矩阵合法：audit × task_completed（K-3 audit_NNN）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "audit_001"
			},
		},
		{
			name: "矩阵合法：audit × note（task_key 可空）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindNote
				in.TaskKey = ""
			},
		},
		{
			name: "矩阵合法：fix × task_completed（K-4 fix_NNN）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeFix
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "fix_001"
			},
		},
		{
			name: "矩阵合法：fix × note（task_key 可空）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeFix
				in.EventKind = EventKindNote
				in.TaskKey = ""
			},
		},

		// --- K-5 边界：note 格填写自由标注 task_key 不强制格式 ---
		{
			name: "K-5 边界：phase × note 填写自由标注 task_key（不强制格式）",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindNote
				in.TaskKey = "自由标注-任意形态"
			},
		},
		{
			name: "K-5 边界：audit × note 填写自由标注 task_key",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindNote
				in.TaskKey = "audit 自由标注"
			},
		},

		// --- 合法边界：文本字段（rune 计数 / trim 判定 / evidence 前缀）---
		{
			name: "合法边界：title 恰好 200 个中文字符（rune 计数，非字节）",
			mutate: func(in *CreateProgressEventInput) {
				in.Title = strings.Repeat("标", 200)
			},
		},
		{
			name: "合法边界：title 首尾空白但中间非空（trim 判定、原值入库）",
			mutate: func(in *CreateProgressEventInput) {
				in.Title = "  有内容的标题  "
			},
		},
		{
			name: "合法边界：detail 恰好 2000 个中文字符",
			mutate: func(in *CreateProgressEventInput) {
				in.Detail = strings.Repeat("详", 2000)
			},
		},
		{
			name: "合法边界：evidence_ref 为 / 开头仓库内相对路径",
			mutate: func(in *CreateProgressEventInput) {
				in.EvidenceRef = "/docs/plan.md"
			},
		},
		{
			name: "合法边界：evidence_ref 为 https:// 开头 URL",
			mutate: func(in *CreateProgressEventInput) {
				in.EvidenceRef = "https://github.com/psco/repo"
			},
		},
		{
			name: "合法边界：task_key 首尾空白（trim 后合法并归一持久化）",
			mutate: func(in *CreateProgressEventInput) {
				in.TaskKey = "  phase15  "
			},
		},

		// --- envelope 前置：INVALID_REPOSITORY_ID ---
		{
			name: "envelope 非法：repository_id 非法 UUID 字符串",
			mutate: func(in *CreateProgressEventInput) {
				in.RepositoryID = "not-a-uuid"
			},
			wantErrCode: "INVALID_REPOSITORY_ID",
		},
		{
			name: "envelope 非法：repository_id 为空串",
			mutate: func(in *CreateProgressEventInput) {
				in.RepositoryID = ""
			},
			wantErrCode: "INVALID_REPOSITORY_ID",
		},

		// --- envelope 前置：INVALID_OCCURRED_AT ---
		{
			name: "envelope 非法：occurred_at 零值（用户声明时间无合法零值）",
			mutate: func(in *CreateProgressEventInput) {
				in.OccurredAt = time.Time{}
			},
			wantErrCode: "INVALID_OCCURRED_AT",
		},

		// --- V1a INVALID_WORKFLOW_TYPE ---
		{
			name: "V1a 非法：workflow_type 越界字符串",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowType("sprint")
			},
			wantErrCode: "INVALID_WORKFLOW_TYPE",
		},
		{
			name: "V1a 非法：workflow_type 空串（proto UNSPECIFIED 解包形态）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowType("")
			},
			wantErrCode: "INVALID_WORKFLOW_TYPE",
		},

		// --- V1b INVALID_EVENT_KIND ---
		{
			name: "V1b 非法：event_kind 越界字符串",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKind("milestone")
			},
			wantErrCode: "INVALID_EVENT_KIND",
		},
		{
			name: "V1b 非法：event_kind 空串（proto UNSPECIFIED 解包形态）",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKind("")
			},
			wantErrCode: "INVALID_EVENT_KIND",
		},

		// --- V7 EVENT_KIND_NOT_ALLOWED（4 禁止格全覆盖）---
		{
			name: "V7 非法：audit × phase_started（矩阵禁止格）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "phase15"
			},
			wantErrCode: "EVENT_KIND_NOT_ALLOWED",
		},
		{
			name: "V7 非法：audit × phase_completed（矩阵禁止格）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindPhaseCompleted
				in.TaskKey = "phase15"
			},
			wantErrCode: "EVENT_KIND_NOT_ALLOWED",
		},
		{
			name: "V7 非法：fix × phase_started（矩阵禁止格）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeFix
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "phase15"
			},
			wantErrCode: "EVENT_KIND_NOT_ALLOWED",
		},
		{
			name: "V7 非法：fix × phase_completed（矩阵禁止格）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeFix
				in.EventKind = EventKindPhaseCompleted
				in.TaskKey = "phase15"
			},
			wantErrCode: "EVENT_KIND_NOT_ALLOWED",
		},

		// --- TASK_KEY_REQUIRED（必填格空串与纯空白）---
		{
			name: "必填非法：phase × phase_started task_key 空串",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = ""
			},
			wantErrCode: "TASK_KEY_REQUIRED",
		},
		{
			name: "必填非法：phase × phase_started task_key 纯空白",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "   "
			},
			wantErrCode: "TASK_KEY_REQUIRED",
		},
		{
			name: "必填非法：phase × phase_completed task_key 空串",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindPhaseCompleted
				in.TaskKey = ""
			},
			wantErrCode: "TASK_KEY_REQUIRED",
		},
		{
			name: "必填非法：phase × task_completed task_key 纯空白",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "   "
			},
			wantErrCode: "TASK_KEY_REQUIRED",
		},
		{
			name: "必填非法：audit × task_completed task_key 纯空白",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "   "
			},
			wantErrCode: "TASK_KEY_REQUIRED",
		},
		{
			name: "必填非法：fix × task_completed task_key 空串",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeFix
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = ""
			},
			wantErrCode: "TASK_KEY_REQUIRED",
		},

		// --- TASK_KEY_FORMAT_INVALID（K-1~K-4 每正则至少 1 个不符值）---
		{
			name: "K-1 格式非法：phase × phase_started 用 phaseNN-MM 形态（含 - 段）",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "phase15-06"
			},
			wantErrCode: "TASK_KEY_FORMAT_INVALID",
		},
		{
			name: "K-1 格式非法：phase × phase_started 数字不足 2 位",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "phase1"
			},
			wantErrCode: "TASK_KEY_FORMAT_INVALID",
		},
		{
			name: "K-2 格式非法：phase × task_completed 缺 -NN 任务段",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "phase15"
			},
			wantErrCode: "TASK_KEY_FORMAT_INVALID",
		},
		{
			name: "K-3 格式非法：audit × task_completed 数字不足 3 位",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "audit_01"
			},
			wantErrCode: "TASK_KEY_FORMAT_INVALID",
		},
		{
			name: "K-4 格式非法：fix × task_completed 尾部混入字母",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeFix
				in.EventKind = EventKindTaskCompleted
				in.TaskKey = "fix_001x"
			},
			wantErrCode: "TASK_KEY_FORMAT_INVALID",
		},

		// --- INVALID_TITLE（空/纯空白 与 超 200 rune 两形态）---
		{
			name: "V9a 非法：title 空串",
			mutate: func(in *CreateProgressEventInput) {
				in.Title = ""
			},
			wantErrCode: "INVALID_TITLE",
		},
		{
			name: "V9a 非法：title 纯空白",
			mutate: func(in *CreateProgressEventInput) {
				in.Title = "   "
			},
			wantErrCode: "INVALID_TITLE",
		},
		{
			name: "V9a 非法：title 超 200 个中文字符（201 rune）",
			mutate: func(in *CreateProgressEventInput) {
				in.Title = strings.Repeat("标", 201)
			},
			wantErrCode: "INVALID_TITLE",
		},

		// --- INVALID_DETAIL ---
		{
			name: "V9b 非法：detail 超 2000 个中文字符（2001 rune）",
			mutate: func(in *CreateProgressEventInput) {
				in.Detail = strings.Repeat("详", 2001)
			},
			wantErrCode: "INVALID_DETAIL",
		},

		// --- INVALID_EVIDENCE_REF ---
		{
			name: "V9c 非法：evidence_ref 以 http:// 开头（非 https）",
			mutate: func(in *CreateProgressEventInput) {
				in.EvidenceRef = "http://example.com/doc"
			},
			wantErrCode: "INVALID_EVIDENCE_REF",
		},
		{
			name: "V9c 非法：evidence_ref 为相对路径（无前缀）",
			mutate: func(in *CreateProgressEventInput) {
				in.EvidenceRef = "docs/plan.md"
			},
			wantErrCode: "INVALID_EVIDENCE_REF",
		},

		// --- INVALID_SOURCE ---
		{
			name: "V9d 非法：source 为 git（预留枚举位，本阶段拒绝）",
			mutate: func(in *CreateProgressEventInput) {
				in.Source = ProgressSourceGit
			},
			wantErrCode: "INVALID_SOURCE",
		},
		{
			name: "V9d 非法：source 为 agent（预留枚举位，本阶段拒绝）",
			mutate: func(in *CreateProgressEventInput) {
				in.Source = ProgressSourceAgent
			},
			wantErrCode: "INVALID_SOURCE",
		},
		{
			name: "V9d 非法：source 为空串（proto 显式 UNSPECIFIED 解包形态）",
			mutate: func(in *CreateProgressEventInput) {
				in.Source = ProgressSource("")
			},
			wantErrCode: "INVALID_SOURCE",
		},

		// --- 执行序首错断言（多错输入只报第一个错误码）---
		{
			name: "执行序：envelope 前置最先（repository_id 非法 + 其余全错）",
			mutate: func(in *CreateProgressEventInput) {
				in.RepositoryID = "not-a-uuid"
				in.OccurredAt = time.Time{}
				in.WorkflowType = WorkflowType("sprint")
				in.Title = ""
			},
			wantErrCode: "INVALID_REPOSITORY_ID",
		},
		{
			name: "执行序：occurred_at 先于枚举校验",
			mutate: func(in *CreateProgressEventInput) {
				in.OccurredAt = time.Time{}
				in.WorkflowType = WorkflowType("sprint")
				in.Title = ""
			},
			wantErrCode: "INVALID_OCCURRED_AT",
		},
		{
			name: "执行序：V1a 先于 V1b / V9（非法枚举 + 空 title）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowType("sprint")
				in.EventKind = EventKind("milestone")
				in.Title = ""
			},
			wantErrCode: "INVALID_WORKFLOW_TYPE",
		},
		{
			name: "执行序：V1b 先于 V7 / V9（event_kind 非法 + 空 title）",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKind("milestone")
				in.TaskKey = ""
				in.Title = ""
			},
			wantErrCode: "INVALID_EVENT_KIND",
		},
		{
			name: "执行序：V7 先于 task_key 与 V9（audit×phase_started + 空 task_key + 空 title）",
			mutate: func(in *CreateProgressEventInput) {
				in.WorkflowType = WorkflowTypeAudit
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = ""
				in.Title = ""
			},
			wantErrCode: "EVENT_KIND_NOT_ALLOWED",
		},
		{
			name: "执行序：TASK_KEY_REQUIRED 先于格式与 V9（纯空白 task_key + 空 title）",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "   "
				in.Title = ""
			},
			wantErrCode: "TASK_KEY_REQUIRED",
		},
		{
			name: "执行序：TASK_KEY_FORMAT_INVALID 先于 V9（格式错 + 超 200 title）",
			mutate: func(in *CreateProgressEventInput) {
				in.EventKind = EventKindPhaseStarted
				in.TaskKey = "phase1"
				in.Title = strings.Repeat("标", 201)
			},
			wantErrCode: "TASK_KEY_FORMAT_INVALID",
		},
		{
			name: "执行序：V9 内部顺序 title 先于 detail / evidence_ref / source",
			mutate: func(in *CreateProgressEventInput) {
				in.Title = ""
				in.Detail = strings.Repeat("详", 2001)
				in.EvidenceRef = "docs/plan.md"
				in.Source = ProgressSourceGit
			},
			wantErrCode: "INVALID_TITLE",
		},
		{
			name: "执行序：V9 内部顺序 detail 先于 evidence_ref / source",
			mutate: func(in *CreateProgressEventInput) {
				in.Detail = strings.Repeat("详", 2001)
				in.EvidenceRef = "docs/plan.md"
				in.Source = ProgressSourceGit
			},
			wantErrCode: "INVALID_DETAIL",
		},
		{
			name: "执行序：V9 内部顺序 evidence_ref 先于 source",
			mutate: func(in *CreateProgressEventInput) {
				in.EvidenceRef = "docs/plan.md"
				in.Source = ProgressSourceGit
			},
			wantErrCode: "INVALID_EVIDENCE_REF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := validInput()
			if tt.mutate != nil {
				tt.mutate(&input)
			}
			err := ValidateCreateProgressEventInput(&input)
			if tt.wantErrCode == "" {
				if err != nil {
					t.Fatalf("期望合法，实际报错: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("期望错误码 %s，实际无错误", tt.wantErrCode)
			}
			if !strings.Contains(err.Error(), "["+tt.wantErrCode+"]") {
				t.Errorf("错误信息缺少稳定错误码 %s: %v", tt.wantErrCode, err)
			}
			if !errors.Is(err, ErrInvalidInput) {
				t.Errorf("错误未包装 ErrInvalidInput: %v", err)
			}
		})
	}
}

// TestValidateTaskKeyTrimNormalize 覆盖 TrimSpace 边界：
// task_key trim 后归一持久化；title trim 仅判定、原值保留。
func TestValidateTaskKeyTrimNormalize(t *testing.T) {
	input := validInput()
	input.TaskKey = "  phase15  "
	if err := ValidateCreateProgressEventInput(&input); err != nil {
		t.Fatalf("期望合法，实际报错: %v", err)
	}
	if input.TaskKey != "phase15" {
		t.Errorf("task_key 未归一为 trim 后值: %q, 期望 %q", input.TaskKey, "phase15")
	}

	titleInput := validInput()
	titleInput.Title = "  有内容的标题  "
	if err := ValidateCreateProgressEventInput(&titleInput); err != nil {
		t.Fatalf("期望合法，实际报错: %v", err)
	}
	if titleInput.Title != "  有内容的标题  " {
		t.Errorf("title 应原值保留: %q", titleInput.Title)
	}
}
