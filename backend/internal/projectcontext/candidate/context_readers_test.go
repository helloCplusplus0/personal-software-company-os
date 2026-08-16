package candidate

import (
	"context"
	"testing"

	"github.com/psco/backend/internal/projectcontext"
)

func TestContextReaders_ReadRules_ExposesSemanticAndRootEntries(t *testing.T) {
	readers := &ContextReaders{}

	rules := readers.ReadRules(context.Background())
	if len(rules) == 0 {
		t.Fatal("expected non-empty rules")
	}

	expectedByKey := map[string]struct {
		summary  string
		entryRef string
	}{
		"product_semantic_positioning": {
			summary:  "Product = 经营目标与交付容器",
			entryRef: "docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md",
		},
		"repository_semantic_positioning": {
			summary:  "Repository = 代码仓库身份对象与项目锚点",
			entryRef: "docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md",
		},
		"plan_entry": {
			summary:  "plan.md 是当前阶段状态与推进路线的唯一正式承接位。",
			entryRef: "plan.md",
		},
		"architecture_map_entry": {
			summary:  "architecture_map.md 负责目录结构、文档分类与迁移落点。",
			entryRef: "architecture_map.md",
		},
		"docs_readme_entry": {
			summary:  "docs/README.md 负责文档总览与 workflow 入口。",
			entryRef: "docs/README.md",
		},
	}

	for key, expected := range expectedByKey {
		t.Run(key, func(t *testing.T) {
			rule, ok := findRuleByKey(rules, key)
			if !ok {
				t.Fatalf("expected rule %q to exist", key)
			}
			if rule.Summary != expected.summary {
				t.Fatalf("expected rule %q summary %q, got %q", key, expected.summary, rule.Summary)
			}
			if rule.EntryRef != expected.entryRef {
				t.Fatalf("expected rule %q entry ref %q, got %q", key, expected.entryRef, rule.EntryRef)
			}
			if rule.EntryKind != "repo_relative_path" {
				t.Fatalf("expected rule %q entry kind repo_relative_path, got %q", key, rule.EntryKind)
			}
		})
	}
}

func TestContextReaders_ReadPhases_ExposesPhase12Entry(t *testing.T) {
	readers := &ContextReaders{}

	phases := readers.ReadPhases(context.Background())
	if len(phases) == 0 {
		t.Fatal("expected non-empty phases")
	}

	phase12, ok := findPhaseByName(phases, "phase12")
	if !ok {
		t.Fatal("expected phase12 entry to exist")
	}
	if phase12.EntryRef != "docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md" {
		t.Fatalf("unexpected phase12 entry ref: %q", phase12.EntryRef)
	}
	if phase12.EntryKind != "repo_relative_path" {
		t.Fatalf("unexpected phase12 entry kind: %q", phase12.EntryKind)
	}
	if phase12.StatusSummary == "" {
		t.Fatal("expected phase12 status summary to be populated")
	}
}

func findRuleByKey(items []projectcontext.RuleEntry, key string) (projectcontext.RuleEntry, bool) {
	for _, item := range items {
		if item.Key == key {
			return item, true
		}
	}
	return projectcontext.RuleEntry{}, false
}

func findPhaseByName(items []projectcontext.PhaseEntry, phase string) (projectcontext.PhaseEntry, bool) {
	for _, item := range items {
		if item.Phase == phase {
			return item, true
		}
	}
	return projectcontext.PhaseEntry{}, false
}
