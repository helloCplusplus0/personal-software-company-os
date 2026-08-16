package renderer

import (
	"strings"
	"testing"

	"github.com/psco/backend/internal/projectcontext"
)

func TestRenderMarkdown_FullResult(t *testing.T) {
	result := &projectcontext.ProjectContextReadResult{
		Repository: &projectcontext.RepositorySummary{
			ID:          "repo-1",
			Name:        "personal-software-company-os",
			Provider:    "github",
			URL:         "https://github.com/user/psco",
			Description: "Personal Software Company OS",
		},
		Product: &projectcontext.ProductSummary{
			ID:          "prod-1",
			Name:        "PSCO",
			Description: "经营与资产系统",
			Status:      "active",
		},
		Modules: []projectcontext.ModuleSummary{
			{ID: "mod-1", Name: "module-registry", Description: "Module Registry", Status: "active"},
			{ID: "mod-2", Name: "decision-center", Description: "Decision Center", Status: "active"},
		},
		Decisions: []projectcontext.DecisionSummary{
			{
				ID:         "dec-1",
				Title:      "Use .proto + ConnectRPC as canonical transport",
				Status:     "accepted",
				Context:    "Need unified contract for all business APIs",
				HitSources: []string{"repository", "product"},
				CreatedAt:  "2026-01-15T00:00:00Z",
			},
		},
		Rules: []projectcontext.RuleEntry{
			{Key: "product_semantic_positioning", Label: "Product 语义定位", Summary: "Product = 经营目标与交付容器", EntryRef: "docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md", EntryKind: "repo_relative_path"},
			{Key: "repository_semantic_positioning", Label: "Repository 语义定位", Summary: "Repository = 代码仓库身份对象与项目锚点", EntryRef: "docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md", EntryKind: "repo_relative_path"},
			{Key: "project_rules", Label: "项目协作规则", Summary: "PSCO 项目的工作流规范", EntryRef: "project_rules.md", EntryKind: "repo_relative_path"},
			{Key: "plan_entry", Label: "阶段路线入口", Summary: "plan.md 是当前阶段状态与推进路线的唯一正式承接位。", EntryRef: "plan.md", EntryKind: "repo_relative_path"},
			{Key: "architecture_map_entry", Label: "目录与迁移落点入口", Summary: "architecture_map.md 负责目录结构、文档分类与迁移落点。", EntryRef: "architecture_map.md", EntryKind: "repo_relative_path"},
			{Key: "docs_readme_entry", Label: "docs 工作流入口", Summary: "docs/README.md 负责文档总览与 workflow 入口。", EntryRef: "docs/README.md", EntryKind: "repo_relative_path"},
			{Key: "tech_stack_baseline", Label: "技术栈基线", Summary: "Durable System Track", EntryRef: "TECH_STACK_BASELINE.md", EntryKind: "repo_relative_path"},
		},
		Phases: []projectcontext.PhaseEntry{
			{Phase: "phase12", Label: "Semantic Alignment & Read-Only Consumption Foundation", StatusSummary: "已建立 /plan 入口", EntryRef: "docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md", EntryKind: "repo_relative_path"},
			{Phase: "phase11", Label: "Project Context Foundation", StatusSummary: "已完成 /plan，当前阶段", EntryRef: "docs/phase/phase11_project_context_foundation_dev_plan.md", EntryKind: "repo_relative_path"},
		},
		Boundaries: []projectcontext.BoundaryEntry{
			{Key: "no_second_fact_source", Label: "不形成第二套事实源", Summary: "Markdown 导出只能从结构化只读结果单向派生。"},
		},
	}

	output := RenderMarkdown(result)

	// 必须包含的节
	requiredSections := []string{
		"# Project Context",
		"## Repository",
		"## Product",
		"## Modules",
		"## Decisions",
		"## Rules & Constraints",
		"## Current Phase",
		"## Boundaries",
	}

	for _, section := range requiredSections {
		if !strings.Contains(output, section) {
			t.Errorf("expected section %q not found in output", section)
		}
	}

	// 必须包含的关键信息
	requiredContent := []string{
		"personal-software-company-os",
		"PSCO",
		"module-registry",
		"decision-center",
		"Use .proto + ConnectRPC as canonical transport",
		"Product = 经营目标与交付容器",
		"Repository = 代码仓库身份对象与项目锚点",
		"项目协作规则",
		"plan.md",
		"architecture_map.md",
		"docs/README.md",
		"Semantic Alignment & Read-Only Consumption Foundation",
		"docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md",
		"Project Context Foundation",
		"project_rules.md",
		"docs/phase/phase11_project_context_foundation_dev_plan.md",
		"不形成第二套事实源",
	}

	for _, content := range requiredContent {
		if !strings.Contains(output, content) {
			t.Errorf("expected content %q not found in output", content)
		}
	}
}

func TestRenderMarkdown_EmptyCollections(t *testing.T) {
	result := &projectcontext.ProjectContextReadResult{
		Repository: &projectcontext.RepositorySummary{
			ID:   "repo-1",
			Name: "test-repo",
		},
		Modules:    []projectcontext.ModuleSummary{},
		Decisions:  []projectcontext.DecisionSummary{},
		Rules:      []projectcontext.RuleEntry{},
		Phases:     []projectcontext.PhaseEntry{},
		Boundaries: []projectcontext.BoundaryEntry{},
	}

	output := RenderMarkdown(result)

	// 空集合应显示 (none)
	if !strings.Contains(output, "(none)") {
		t.Error("expected (none) placeholder for empty collections")
	}

	// nil Product 不应渲染 Product 节
	if strings.Contains(output, "## Product") {
		t.Error("Product section should not appear when Product is nil")
	}
}

func TestRenderMarkdown_NilProduct(t *testing.T) {
	result := &projectcontext.ProjectContextReadResult{
		Repository: &projectcontext.RepositorySummary{
			ID:   "repo-1",
			Name: "test-repo",
		},
		// Product is nil
		Modules:    []projectcontext.ModuleSummary{},
		Decisions:  []projectcontext.DecisionSummary{},
		Rules:      []projectcontext.RuleEntry{},
		Boundaries: []projectcontext.BoundaryEntry{},
	}

	output := RenderMarkdown(result)

	if strings.Contains(output, "## Product") {
		t.Error("Product section should not appear when Product is nil")
	}
}
