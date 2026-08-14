// Package service — Decision Center 写组业务编排层测试。
//
// 测试范围：
//   - isValidDecisionStatus：四态枚举校验
//   - validStatusTransitions：状态推进矩阵完整性
//   - isValidTransition：状态迁移合法性校验
//
// 文件落点：backend/internal/decisioncenter/service/command_service_test.go
package service

import (
	"testing"

	"github.com/psco/backend/internal/decisioncenter"
)

func TestIsValidDecisionStatus(t *testing.T) {
	tests := []struct {
		name   string
		status decisioncenter.DecisionStatus
		want   bool
	}{
		{"proposed is valid", decisioncenter.DecisionStatusProposed, true},
		{"active is valid", decisioncenter.DecisionStatusActive, true},
		{"superseded is valid", decisioncenter.DecisionStatusSuperseded, true},
		{"archived is valid", decisioncenter.DecisionStatusArchived, true},
		{"empty string is invalid", "", false},
		{"unknown status is invalid", decisioncenter.DecisionStatus("unknown"), false},
		{"dismissed is invalid", decisioncenter.DecisionStatus("dismissed"), false},
		{"resolved is invalid", decisioncenter.DecisionStatus("resolved"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidDecisionStatus(tt.status)
			if got != tt.want {
				t.Errorf("isValidDecisionStatus(%q) = %v, want %v", tt.status, got, tt.want)
			}
		})
	}
}

func TestValidStatusTransitionsMatrix(t *testing.T) {
	// proposed → active / superseded / archived
	proposedTargets := validStatusTransitions[decisioncenter.DecisionStatusProposed]
	if len(proposedTargets) != 3 {
		t.Fatalf("proposed should have 3 allowed targets, got %d", len(proposedTargets))
	}
	expectedProposed := map[decisioncenter.DecisionStatus]bool{
		decisioncenter.DecisionStatusActive:     true,
		decisioncenter.DecisionStatusSuperseded: true,
		decisioncenter.DecisionStatusArchived:   true,
	}
	for _, target := range proposedTargets {
		if !expectedProposed[target] {
			t.Errorf("proposed → %s is unexpected", target)
		}
		delete(expectedProposed, target)
	}
	if len(expectedProposed) > 0 {
		t.Errorf("proposed missing targets: %v", expectedProposed)
	}

	// active → superseded / archived
	activeTargets := validStatusTransitions[decisioncenter.DecisionStatusActive]
	if len(activeTargets) != 2 {
		t.Fatalf("active should have 2 allowed targets, got %d", len(activeTargets))
	}
	expectedActive := map[decisioncenter.DecisionStatus]bool{
		decisioncenter.DecisionStatusSuperseded: true,
		decisioncenter.DecisionStatusArchived:   true,
	}
	for _, target := range activeTargets {
		if !expectedActive[target] {
			t.Errorf("active → %s is unexpected", target)
		}
		delete(expectedActive, target)
	}
	if len(expectedActive) > 0 {
		t.Errorf("active missing targets: %v", expectedActive)
	}

	// superseded → 终态（不在 map 中）
	if _, exists := validStatusTransitions[decisioncenter.DecisionStatusSuperseded]; exists {
		t.Error("superseded should be terminal (not in transitions map)")
	}

	// archived → 终态（不在 map 中）
	if _, exists := validStatusTransitions[decisioncenter.DecisionStatusArchived]; exists {
		t.Error("archived should be terminal (not in transitions map)")
	}
}

// isValidTransition 校验从当前状态推进到目标状态是否合法。
// 终态（superseded / archived）不可继续推进。
func isValidTransition(current, target decisioncenter.DecisionStatus) bool {
	allowed := validStatusTransitions[current]
	if allowed == nil {
		return false
	}
	for _, a := range allowed {
		if a == target {
			return true
		}
	}
	return false
}

func TestIsValidTransition(t *testing.T) {
	tests := []struct {
		name    string
		current decisioncenter.DecisionStatus
		target  decisioncenter.DecisionStatus
		want    bool
	}{
		// proposed 允许推进
		{"proposed → active", decisioncenter.DecisionStatusProposed, decisioncenter.DecisionStatusActive, true},
		{"proposed → superseded", decisioncenter.DecisionStatusProposed, decisioncenter.DecisionStatusSuperseded, true},
		{"proposed → archived", decisioncenter.DecisionStatusProposed, decisioncenter.DecisionStatusArchived, true},
		// proposed 禁止推进
		{"proposed → proposed (self)", decisioncenter.DecisionStatusProposed, decisioncenter.DecisionStatusProposed, false},
		// active 允许推进
		{"active → superseded", decisioncenter.DecisionStatusActive, decisioncenter.DecisionStatusSuperseded, true},
		{"active → archived", decisioncenter.DecisionStatusActive, decisioncenter.DecisionStatusArchived, true},
		// active 禁止推进
		{"active → proposed (backward)", decisioncenter.DecisionStatusActive, decisioncenter.DecisionStatusProposed, false},
		{"active → active (self)", decisioncenter.DecisionStatusActive, decisioncenter.DecisionStatusActive, false},
		// 终态禁止推进
		{"superseded → active", decisioncenter.DecisionStatusSuperseded, decisioncenter.DecisionStatusActive, false},
		{"superseded → archived", decisioncenter.DecisionStatusSuperseded, decisioncenter.DecisionStatusArchived, false},
		{"superseded → proposed", decisioncenter.DecisionStatusSuperseded, decisioncenter.DecisionStatusProposed, false},
		{"archived → active", decisioncenter.DecisionStatusArchived, decisioncenter.DecisionStatusActive, false},
		{"archived → superseded", decisioncenter.DecisionStatusArchived, decisioncenter.DecisionStatusSuperseded, false},
		{"archived → proposed", decisioncenter.DecisionStatusArchived, decisioncenter.DecisionStatusProposed, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidTransition(tt.current, tt.target)
			if got != tt.want {
				t.Errorf("isValidTransition(%q, %q) = %v, want %v", tt.current, tt.target, got, tt.want)
			}
		})
	}
}

// TestValidStatusTransitionsExhaustive 验证所有合法迁移 + 禁止越界 + 终态不可推进。
func TestValidStatusTransitionsExhaustive(t *testing.T) {
	allStatuses := []decisioncenter.DecisionStatus{
		decisioncenter.DecisionStatusProposed,
		decisioncenter.DecisionStatusActive,
		decisioncenter.DecisionStatusSuperseded,
		decisioncenter.DecisionStatusArchived,
	}

	allowedCount := 0
	for _, current := range allStatuses {
		for _, target := range allStatuses {
			allowed := isValidTransition(current, target)
			if allowed {
				allowedCount++
			}
			// 终态不可推进
			if current == decisioncenter.DecisionStatusSuperseded || current == decisioncenter.DecisionStatusArchived {
				if allowed {
					t.Errorf("terminal state %q should not allow transition to %q", current, target)
				}
			}
			// 不允许回退到 proposed
			if current != decisioncenter.DecisionStatusProposed && target == decisioncenter.DecisionStatusProposed {
				if allowed {
					t.Errorf("%q should not allow backward transition to proposed", current)
				}
			}
		}
	}

	// 总共应有 5 条合法迁移：proposed→3 + active→2
	if allowedCount != 5 {
		t.Errorf("expected exactly 5 valid transitions, got %d", allowedCount)
	}
}