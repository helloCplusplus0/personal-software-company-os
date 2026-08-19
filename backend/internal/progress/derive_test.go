// 派生摘要纯函数单元测试（phase15-03 三派生算法边界全覆盖，纯内存无需 DB）。
//
// 覆盖冻结要求（phase15-06 Task 3 / phase15-03 Scenario）：
//   - 空集（nil 与空切片）/ 仅 note / 从未开始 / 进行中 / 全部完结 /
//     完结不同 key / 补录同刻 tiebreak（同 occurred_at+created_at 下
//     id DESC 序判定，构造输入时按 DESC 序给定）/ recent 截断（>10 条取前 10）
//   - latest_task_completed 扫描全量事件（不受 recent N=10 截断影响）
//
// 输入序约定：本测试构造的 events 切片即 SQL ORDER BY
// (occurred_at DESC, created_at DESC, id DESC) 的返回序——索引 0 = 最新事件；
// DeriveProgressSummary 不重排，直接消费输入序。
//
// 文件落点：backend/internal/progress/derive_test.go
package progress

import (
	"reflect"
	"testing"
	"time"
)

// ev 构造派生测试事件（时间同刻：tiebreak 场景由输入序即 id DESC 序承载）。
func ev(id string, workflowType WorkflowType, eventKind EventKind, taskKey, title string) ProgressEventReadResult {
	return ProgressEventReadResult{
		ID:           id,
		RepositoryID: "11111111-1111-1111-1111-111111111111",
		WorkflowType: workflowType,
		EventKind:    eventKind,
		TaskKey:      taskKey,
		Title:        title,
		Source:       ProgressSourceManual,
		OccurredAt:   time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC),
		CreatedAt:    time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC),
	}
}

// TestDeriveProgressSummary 表驱动覆盖三派生项全部边界。
func TestDeriveProgressSummary(t *testing.T) {
	tests := []struct {
		name             string
		events           []ProgressEventReadResult
		wantCurrentKey   string   // 两种空值情形（从未开始/全部完结）均为空串
		wantCurrentLabel string   // 同上
		wantLatestTaskID string   // 空串 = 期望 nil
		wantRecentIDs    []string // 期望 recent_events 的 ID 序（截断后）
	}{
		{
			name:             "空集（空切片）：四字段零值，RecentEvents 空切片非 nil",
			events:           []ProgressEventReadResult{},
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{},
		},
		{
			name:             "空集（nil 切片）：同空切片语义",
			events:           nil,
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{},
		},
		{
			name: "仅 note：current_phase 空（从未开始），recent 含该条",
			events: []ProgressEventReadResult{
				ev("e1", WorkflowTypePhase, EventKindNote, "", "记录一条"),
			},
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{"e1"},
		},
		{
			name: "从未开始：有 task_completed 但无 phase_started",
			events: []ProgressEventReadResult{
				ev("e3", WorkflowTypeFix, EventKindTaskCompleted, "fix_002", "修复二"),
				ev("e2", WorkflowTypePhase, EventKindTaskCompleted, "phase15-06", "任务六"),
				ev("e1", WorkflowTypeAudit, EventKindNote, "", "审计备注"),
			},
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "e3",
			wantRecentIDs:    []string{"e3", "e2", "e1"},
		},
		{
			name: "进行中：最新 phase_started 之后无同 key phase_completed",
			events: []ProgressEventReadResult{
				ev("e3", WorkflowTypePhase, EventKindTaskCompleted, "phase15-06", "任务六"),
				ev("e2", WorkflowTypePhase, EventKindPhaseStarted, "phase15", "phase15 开始"),
				ev("e1", WorkflowTypePhase, EventKindNote, "", "备注"),
			},
			wantCurrentKey:   "phase15",
			wantCurrentLabel: "phase15 开始",
			wantLatestTaskID: "e3",
			wantRecentIDs:    []string{"e3", "e2", "e1"},
		},
		{
			name: "全部完结：最新 phase_started 之后存在同 task_key phase_completed → current_phase 空",
			events: []ProgressEventReadResult{
				ev("e3", WorkflowTypePhase, EventKindPhaseCompleted, "phase15", "phase15 完成"),
				ev("e2", WorkflowTypePhase, EventKindPhaseStarted, "phase15", "phase15 开始"),
				ev("e1", WorkflowTypePhase, EventKindNote, "", "备注"),
			},
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{"e3", "e2", "e1"},
		},
		{
			name: "完结不同 key：phase_completed 的 key ≠ 最新 phase_started 的 key → 仍进行中（同 key 匹配）",
			events: []ProgressEventReadResult{
				ev("e3", WorkflowTypePhase, EventKindPhaseCompleted, "phase14", "phase14 完成"),
				ev("e2", WorkflowTypePhase, EventKindPhaseStarted, "phase15", "phase15 开始"),
			},
			wantCurrentKey:   "phase15",
			wantCurrentLabel: "phase15 开始",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{"e3", "e2"},
		},
		{
			name: "完结在开始之前（DESC 序索引更大 = 时间更早）：不算完结 → 进行中",
			events: []ProgressEventReadResult{
				ev("e3", WorkflowTypePhase, EventKindPhaseStarted, "phase15", "phase15 开始"),
				ev("e2", WorkflowTypePhase, EventKindPhaseCompleted, "phase15", "phase15 早期完结（更早）"),
				ev("e1", WorkflowTypePhase, EventKindNote, "", "备注"),
			},
			wantCurrentKey:   "phase15",
			wantCurrentLabel: "phase15 开始",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{"e3", "e2", "e1"},
		},
		{
			name: "补录同刻 tiebreak（一）：id DESC 序中同 key 完结更晚 → 全部完结",
			// 输入按 SQL ORDER BY 结果给定：id …bb > …aa → bb 排前（更晚），
			// bb 为同 key phase_completed → 最新 phase_started(aa) 之后存在完结
			events: []ProgressEventReadResult{
				ev("00000000-0000-0000-0000-0000000000bb", WorkflowTypePhase, EventKindPhaseCompleted, "phase15", "同刻完结（id 更大视为更晚）"),
				ev("00000000-0000-0000-0000-0000000000aa", WorkflowTypePhase, EventKindPhaseStarted, "phase15", "同刻开始"),
			},
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{"00000000-0000-0000-0000-0000000000bb", "00000000-0000-0000-0000-0000000000aa"},
		},
		{
			name: "补录同刻 tiebreak（二）：id DESC 序中 started 更晚 → 进行中",
			events: []ProgressEventReadResult{
				ev("00000000-0000-0000-0000-0000000000bb", WorkflowTypePhase, EventKindPhaseStarted, "phase15", "同刻开始（id 更大视为更晚）"),
				ev("00000000-0000-0000-0000-0000000000aa", WorkflowTypePhase, EventKindPhaseCompleted, "phase15", "同刻完结"),
			},
			wantCurrentKey:   "phase15",
			wantCurrentLabel: "同刻开始（id 更大视为更晚）",
			wantLatestTaskID: "",
			wantRecentIDs:    []string{"00000000-0000-0000-0000-0000000000bb", "00000000-0000-0000-0000-0000000000aa"},
		},
		{
			name: "recent 截断：12 条取前 10 条（索引 0-9，最新优先）",
			events: func() []ProgressEventReadResult {
				events := make([]ProgressEventReadResult, 0, 12)
				for i := 12; i >= 1; i-- { // e12（最新）→ e01
					events = append(events, ev(
						"e"+string(rune('0'+i/10))+string(rune('0'+i%10)),
						WorkflowTypePhase, EventKindNote, "", "记录",
					))
				}
				return events
			}(),
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "",
			wantRecentIDs: []string{
				"e12", "e11", "e10", "e09", "e08", "e07", "e06", "e05", "e04", "e03",
			},
		},
		{
			name: "latest_task_completed 三轨同序取最新：取 i=0 起首个 task_completed（fix 轨）",
			events: []ProgressEventReadResult{
				ev("e5", WorkflowTypeFix, EventKindTaskCompleted, "fix_003", "修复三"),
				ev("e4", WorkflowTypePhase, EventKindTaskCompleted, "phase15-06", "任务六"),
				ev("e3", WorkflowTypeAudit, EventKindTaskCompleted, "audit_001", "审计一"),
				ev("e2", WorkflowTypePhase, EventKindPhaseStarted, "phase15", "phase15 开始"),
				ev("e1", WorkflowTypePhase, EventKindNote, "", "备注"),
			},
			wantCurrentKey:   "phase15",
			wantCurrentLabel: "phase15 开始",
			wantLatestTaskID: "e5",
			wantRecentIDs:    []string{"e5", "e4", "e3", "e2", "e1"},
		},
		{
			name: "latest_task_completed 扫描全量：task_completed 在 recent 截断范围外（索引 10）仍被取到",
			events: func() []ProgressEventReadResult {
				events := make([]ProgressEventReadResult, 0, 12)
				for i := 12; i >= 1; i-- { // e12（最新）→ e01
					kind := EventKindNote
					taskKey := ""
					if i == 2 { // e02：索引 10，位于 recent N=10 截断线之外
						kind = EventKindTaskCompleted
						taskKey = "phase15-01"
					}
					events = append(events, ev(
						"e"+string(rune('0'+i/10))+string(rune('0'+i%10)),
						WorkflowTypePhase, kind, taskKey, "记录",
					))
				}
				return events
			}(),
			wantCurrentKey:   "",
			wantCurrentLabel: "",
			wantLatestTaskID: "e02",
			wantRecentIDs: []string{
				"e12", "e11", "e10", "e09", "e08", "e07", "e06", "e05", "e04", "e03",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summary := DeriveProgressSummary(tt.events)

			if summary.CurrentPhaseKey != tt.wantCurrentKey {
				t.Errorf("CurrentPhaseKey = %q, 期望 %q", summary.CurrentPhaseKey, tt.wantCurrentKey)
			}
			if summary.CurrentPhaseLabel != tt.wantCurrentLabel {
				t.Errorf("CurrentPhaseLabel = %q, 期望 %q", summary.CurrentPhaseLabel, tt.wantCurrentLabel)
			}

			if tt.wantLatestTaskID == "" {
				if summary.LatestTaskCompleted != nil {
					t.Errorf("LatestTaskCompleted 期望 nil, 实际 %+v", summary.LatestTaskCompleted)
				}
			} else {
				if summary.LatestTaskCompleted == nil {
					t.Fatalf("LatestTaskCompleted 期望 %s, 实际 nil", tt.wantLatestTaskID)
				}
				if summary.LatestTaskCompleted.ID != tt.wantLatestTaskID {
					t.Errorf("LatestTaskCompleted.ID = %q, 期望 %q", summary.LatestTaskCompleted.ID, tt.wantLatestTaskID)
				}
			}

			if summary.RecentEvents == nil {
				t.Fatalf("RecentEvents 恒非 nil（空态为空切片），实际 nil")
			}
			gotIDs := make([]string, 0, len(summary.RecentEvents))
			for _, e := range summary.RecentEvents {
				gotIDs = append(gotIDs, e.ID)
			}
			if !reflect.DeepEqual(gotIDs, tt.wantRecentIDs) {
				t.Errorf("RecentEvents IDs = %v, 期望 %v", gotIDs, tt.wantRecentIDs)
			}
		})
	}
}
