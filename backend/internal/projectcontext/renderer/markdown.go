// Package renderer — projectcontext Markdown 导出渲染器。
//
// 本文件承接 phase11-08 AGENTS 风格 Markdown 导出。
// 渲染器只消费 phase11-07 的结构化只读结果，不读取数据库、不扫描目录、不引入第二套事实源。
//
// 文件落点：backend/internal/projectcontext/renderer/markdown.go
package renderer

import (
	"fmt"
	"strings"

	"github.com/psco/backend/internal/projectcontext"
)

// RenderMarkdown 将 ProjectContextReadResult 单向渲染为 AGENTS 风格 Markdown。
//
// 渲染规则：
//   - 标题层级从 H1 开始，按节组织
//   - 空值节不渲染（Product 为 nil 时不输出 Product 节）
//   - 空列表节输出 "(none)" 占位
//   - 不离开结构化结果补充新事实
func RenderMarkdown(result *projectcontext.ProjectContextReadResult) string {
	var b strings.Builder

	// 标题
	b.WriteString("# Project Context\n\n")

	// Repository
	renderRepository(&b, result.Repository)

	// Product
	renderProduct(&b, result.Product)

	// Modules
	renderModules(&b, result.Modules)

	// Decisions
	renderDecisions(&b, result.Decisions)

	// Rules & Constraints
	renderRules(&b, result.Rules)

	// Current Phase
	renderPhases(&b, result.Phases)

        // Boundaries (what this project does NOT do)
        renderBoundaries(&b, result.Boundaries)

	return b.String()
}

func renderRepository(b *strings.Builder, repo *projectcontext.RepositorySummary) {
	if repo == nil {
		return
	}
	b.WriteString("## Repository\n\n")
	b.WriteString(fmt.Sprintf("- **Name**: %s\n", repo.Name))
	b.WriteString(fmt.Sprintf("- **ID**: `%s`\n", repo.ID))
	if repo.Provider != "" {
		b.WriteString(fmt.Sprintf("- **Provider**: %s\n", repo.Provider))
	}
	if repo.URL != "" {
		b.WriteString(fmt.Sprintf("- **URL**: %s\n", repo.URL))
	}
	if repo.Description != "" {
		b.WriteString(fmt.Sprintf("- **Description**: %s\n", repo.Description))
	}
	b.WriteString("\n")
}

func renderProduct(b *strings.Builder, product *projectcontext.ProductSummary) {
	if product == nil {
		return
	}
	b.WriteString("## Product\n\n")
	b.WriteString(fmt.Sprintf("- **Name**: %s\n", product.Name))
	b.WriteString(fmt.Sprintf("- **ID**: `%s`\n", product.ID))
	b.WriteString(fmt.Sprintf("- **Status**: %s\n", product.Status))
	if product.Description != "" {
		b.WriteString(fmt.Sprintf("- **Description**: %s\n", product.Description))
	}
	b.WriteString("\n")
}

func renderModules(b *strings.Builder, modules []projectcontext.ModuleSummary) {
	b.WriteString("## Modules\n\n")
	if len(modules) == 0 {
		b.WriteString("(none)\n\n")
		return
	}
	for _, m := range modules {
		b.WriteString(fmt.Sprintf("- **%s** (`%s`) — %s [%s]\n", m.Name, m.ID, m.Description, m.Status))
	}
	b.WriteString("\n")
}

func renderDecisions(b *strings.Builder, decisions []projectcontext.DecisionSummary) {
	b.WriteString("## Decisions\n\n")
	if len(decisions) == 0 {
		b.WriteString("(none)\n\n")
		return
	}
	for _, d := range decisions {
		hitSources := strings.Join(d.HitSources, ", ")
		b.WriteString(fmt.Sprintf("### %s\n\n", d.Title))
		b.WriteString(fmt.Sprintf("- **Status**: %s\n", d.Status))
		if d.Context != "" {
			b.WriteString(fmt.Sprintf("- **Context**: %s\n", d.Context))
		}
		b.WriteString(fmt.Sprintf("- **Hit Sources**: %s\n", hitSources))
		b.WriteString(fmt.Sprintf("- **Created**: %s\n", d.CreatedAt))
		b.WriteString("\n")
	}
}

func renderRules(b *strings.Builder, rules []projectcontext.RuleEntry) {
	b.WriteString("## Rules & Constraints\n\n")
	if len(rules) == 0 {
		b.WriteString("(none)\n\n")
		return
	}
	for _, r := range rules {
                b.WriteString(fmt.Sprintf("- **%s**: %s%s\n", r.Label, r.Summary, formatControlledRef(r.EntryKind, r.EntryRef)))
	}
	b.WriteString("\n")
}

func renderPhases(b *strings.Builder, phases []projectcontext.PhaseEntry) {
	b.WriteString("## Current Phase\n\n")
	if len(phases) == 0 {
		b.WriteString("(none)\n\n")
		return
	}
	for _, p := range phases {
                b.WriteString(fmt.Sprintf("- **%s** (%s): %s%s\n", p.Label, p.Phase, p.StatusSummary, formatControlledRef(p.EntryKind, p.EntryRef)))
	}
	b.WriteString("\n")
}

// renderBoundaries 输出当前阶段明确不做或不承接的边界摘要。
// 这些边界必须来自结构化读取结果，不能在渲染层额外补充第二套事实。
func renderBoundaries(b *strings.Builder, boundaries []projectcontext.BoundaryEntry) {
	b.WriteString("## Boundaries (What This Project Does NOT Do)\n\n")
        if len(boundaries) == 0 {
                b.WriteString("(none)\n\n")
                return
        }
        for _, boundary := range boundaries {
                b.WriteString(fmt.Sprintf("- **%s**: %s\n", boundary.Label, boundary.Summary))
        }
	b.WriteString("\n")
}

func formatControlledRef(entryKind, entryRef string) string {
        if entryKind == "" || entryRef == "" {
                return ""
        }
        return fmt.Sprintf(" (`%s`: `%s`)", entryKind, entryRef)
}
