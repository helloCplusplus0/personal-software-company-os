// Package candidate — projectcontext 跨模块 reader 接口定义与实现（由 projectcontext 拥有）。
//
// 本文件承接 GetProjectContext 所需的全部跨模块读取：
//   - repository 身份读取
//   - 关联 product 摘要读取（通过 product_repositories）
//   - 关联 module 摘要读取（通过 module_repositories）
//   - 关联 decision 摘要读取（两类 module-link 派生命中）
//
// 文件落点：backend/internal/projectcontext/candidate/context_readers.go
package candidate

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/projectcontext"
)

// ContextReaders 承接 GetProjectContext 所需的全部跨模块 reader。
//
// 由 platform 装配点构造并注入到 projectcontext QueryService。
// 每个 reader 直接读取对应 canonical 模块的表，但在本 candidate/ 子包内隔离。
type ContextReaders struct {
	pool *pgxpool.Pool
}

// NewContextReaders 构造 ContextReaders。
func NewContextReaders(pool *pgxpool.Pool) *ContextReaders {
	return &ContextReaders{pool: pool}
}

// ReadRepository 读取仓库身份摘要。
// 若仓库不存在，返回 ErrRepositoryNotFound。
func (r *ContextReaders) ReadRepository(ctx context.Context, repositoryID string) (*projectcontext.RepositorySummary, error) {
	var summary projectcontext.RepositorySummary
	err := r.pool.QueryRow(ctx, `
		SELECT id, name, COALESCE(provider, ''), COALESCE(url, ''), ''::text AS description
		FROM repositories
		WHERE id = $1
	`, repositoryID).Scan(
		&summary.ID, &summary.Name, &summary.Provider, &summary.URL, &summary.Description,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: %w", projectcontext.ErrRepositoryNotFound, err)
		}
		return nil, fmt.Errorf("read repository: %w", err)
	}
	return &summary, nil
}

// ValidateRepositoryBinding 校验当前仓库是否已完成 phase11 所要求的 Repository Binding。
//
// 当前阶段将“绑定完成”明确解释为：
//   - 至少存在一条 product_repositories 绑定
//   - 至少存在一条 module_repositories 映射
//
// 若仓库绑定不完整，返回 ErrRepositoryBindingIncomplete。
func (r *ContextReaders) ValidateRepositoryBinding(ctx context.Context, repositoryID string) error {
	var hasProductBinding bool
	if err := r.pool.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM product_repositories WHERE repository_id = $1)
	`, repositoryID).Scan(&hasProductBinding); err != nil {
		return fmt.Errorf("check product bindings: %w", err)
	}

	var hasModuleBinding bool
	if err := r.pool.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM module_repositories WHERE repository_id = $1)
	`, repositoryID).Scan(&hasModuleBinding); err != nil {
		return fmt.Errorf("check module bindings: %w", err)
	}

	if !hasProductBinding || !hasModuleBinding {
		return projectcontext.ErrRepositoryBindingIncomplete
	}
	return nil
}

// ReadProduct 通过 product_repositories 读取关联 product 摘要。
// 若无关联 product，返回 nil（非错误）。
func (r *ContextReaders) ReadProduct(ctx context.Context, repositoryID string) (*projectcontext.ProductSummary, error) {
	var summary projectcontext.ProductSummary
	err := r.pool.QueryRow(ctx, `
		SELECT p.id, p.name, COALESCE(p.description, ''), COALESCE(p.status, 'active')
		FROM products p
		INNER JOIN product_repositories pr ON pr.product_id = p.id
		WHERE pr.repository_id = $1
		LIMIT 1
	`, repositoryID).Scan(
		&summary.ID, &summary.Name, &summary.Description, &summary.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("read product: %w", err)
	}
	return &summary, nil
}

// ReadModules 通过 module_repositories 读取关联 module 摘要。
// 若无关联 module，返回空列表（非错误）。
func (r *ContextReaders) ReadModules(ctx context.Context, repositoryID string) ([]projectcontext.ModuleSummary, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT m.id, m.name, COALESCE(m.description, ''), COALESCE(m.status, 'active')
		FROM modules m
		INNER JOIN module_repositories mr ON mr.module_id = m.id
		WHERE mr.repository_id = $1
		ORDER BY m.name
	`, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("read modules: %w", err)
	}
	defer rows.Close()

	var summaries []projectcontext.ModuleSummary
	for rows.Next() {
		var s projectcontext.ModuleSummary
		if err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Status); err != nil {
			return nil, fmt.Errorf("scan module: %w", err)
		}
		summaries = append(summaries, s)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter module rows: %w", err)
	}
	return summaries, nil
}

// decisionRaw 从数据库读取的原始 decision 行。
type decisionRaw struct {
	ID        string
	Title     string
	Status    string
	Context   string
	CreatedAt string
	HitSource string
}

// ReadDecisions 按两类 module-link 派生命中范围读取关联 decision 摘要。
//
// 两类命中：
//  1. 命中当前 Repository 已映射 Module 的 Decision（通过 module_repositories → decision_links）
//  2. 命中当前 Repository 已绑定 Product 所属 Module 的 Decision（通过 product_repositories → product_modules → decision_links）
//
// 去重规则：同一 decision_id 命中多类关系时，以 decision_id 去重并合并 hit_sources。
// 过滤规则：只承接 status != 'archived' 的 Decision。
func (r *ContextReaders) ReadDecisions(ctx context.Context, repositoryID string) ([]projectcontext.DecisionSummary, error) {
	// 两类命中统一查询，通过 UNION 合并并在 Go 层去重
	query := `
		SELECT d.id, d.title, COALESCE(d.status, 'proposed'), COALESCE(d.context, ''), 
		       d.created_at::text, $2::text AS hit_source
		FROM decisions d
		INNER JOIN decision_links dl ON dl.decision_id = d.id
		INNER JOIN modules m ON m.id = dl.module_id
		INNER JOIN module_repositories mr ON mr.module_id = m.id
		WHERE mr.repository_id = $1 AND COALESCE(d.status, 'proposed') != 'archived'
		UNION
		SELECT d.id, d.title, COALESCE(d.status, 'proposed'), COALESCE(d.context, ''),
		       d.created_at::text, $3::text AS hit_source
		FROM decisions d
		INNER JOIN decision_links dl ON dl.decision_id = d.id
		INNER JOIN modules m ON m.id = dl.module_id
		INNER JOIN product_modules pm ON pm.module_id = m.id
		INNER JOIN product_repositories pr ON pr.product_id = pm.product_id
		WHERE pr.repository_id = $1 AND COALESCE(d.status, 'proposed') != 'archived'
		ORDER BY created_at DESC
	`

	hitSource1 := "repository_module_mapping"
	hitSource2 := "bound_product_module"
	rows, err := r.pool.Query(ctx, query, repositoryID, hitSource1, hitSource2)
	if err != nil {
		return nil, fmt.Errorf("read decisions: %w", err)
	}
	defer rows.Close()

	// 使用 map 去重并合并 hit_sources
	seen := make(map[string]*projectcontext.DecisionSummary)
	var order []string
	for rows.Next() {
		var raw decisionRaw
		if err := rows.Scan(&raw.ID, &raw.Title, &raw.Status, &raw.Context, &raw.CreatedAt, &raw.HitSource); err != nil {
			return nil, fmt.Errorf("scan decision: %w", err)
		}
		if existing, ok := seen[raw.ID]; ok {
			existing.HitSources = appendUnique(existing.HitSources, raw.HitSource)
		} else {
			seen[raw.ID] = &projectcontext.DecisionSummary{
				ID:         raw.ID,
				Title:      raw.Title,
				Status:     raw.Status,
				Context:    raw.Context,
				HitSources: []string{raw.HitSource},
				CreatedAt:  raw.CreatedAt,
			}
			order = append(order, raw.ID)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter decision rows: %w", err)
	}

	// 保持 ORDER BY 顺序
	result := make([]projectcontext.DecisionSummary, 0, len(order))
	for _, id := range order {
		result = append(result, *seen[id])
	}
	return result, nil
}

func appendUnique(items []string, value string) []string {
	for _, item := range items {
		if item == value {
			return items
		}
	}
	return append(items, value)
}

// ReadRules 读取规则与约束入口摘要。
// 当前阶段从 project_rules.md 与 TECH_STACK_BASELINE.md 的摘要投影，并返回可定位入口。
func (r *ContextReaders) ReadRules(ctx context.Context) []projectcontext.RuleEntry {
	// 当前阶段返回固定规则入口摘要，不依赖数据库
	return []projectcontext.RuleEntry{
		{
			Key:       "project_rules",
			Label:     "项目协作规则",
			Summary:   "PSCO 项目的工作流规范、单一真相源约束、技术栈选择规则与协作门禁",
			EntryRef:  "project_rules.md",
			EntryKind: "repo_relative_path",
		},
		{
			Key:       "tech_stack_baseline",
			Label:     "技术栈基线",
			Summary:   "Durable System Track: React + Go + PostgreSQL + .proto + ConnectRPC",
			EntryRef:  "TECH_STACK_BASELINE.md",
			EntryKind: "repo_relative_path",
		},
	}
}

// ReadPhases 读取当前 phase 相关的文档入口。
// 当前阶段返回固定 phase 入口摘要与可定位入口。
func (r *ContextReaders) ReadPhases(ctx context.Context) []projectcontext.PhaseEntry {
	return []projectcontext.PhaseEntry{
		{
			Phase:         "phase11",
			Label:         "Project Context Foundation",
			StatusSummary: "已完成正式 /plan 并建立三件套，当前冻结的单一主交付能力为根级上下文真相源治理 + 最小只读项目上下文导出",
			EntryRef:      "docs/phase/phase11_project_context_foundation_dev_plan.md",
			EntryKind:     "repo_relative_path",
		},
		{
			Phase:         "phase10",
			Label:         "Asset-Action Closure",
			StatusSummary: "已完成 /plan -> /spec -> 实现 -> 验收 -> 收口，最近完成正式业务 phase",
			EntryRef:      "docs/phase/phase10_asset_action_closure_foundation_dev_plan.md",
			EntryKind:     "repo_relative_path",
		},
	}
}

// ReadBoundaries 读取当前阶段明确不做或不承接的边界摘要。
// 当前阶段返回 phase11 已冻结的最小边界集合，供结构化读取与 Markdown 导出共用。
func (r *ContextReaders) ReadBoundaries(ctx context.Context) []projectcontext.BoundaryEntry {
        return []projectcontext.BoundaryEntry{
                {
                        Key:     "no_mcp_or_cli",
                        Label:   "不做 MCP / CLI / agent 自动写回",
                        Summary: "当前阶段不引入 MCP 协议层、CLI 工具、agent 自动写回、Draft 接口或审批流。",
                },
                {
                        Key:     "not_process_controller",
                        Label:   "不把 PSCO 做成开发流程控制器",
                        Summary: "PSCO 当前定位是上下文系统，而不是 IDE 现场流程编排器或开发过程控制台。",
                },
                {
                        Key:     "web_not_chat_workspace",
                        Label:   "不把 web 退化为对话式 agent 工作台",
                        Summary: "web 继续承接全局查看、关系校对、回顾、人工修正与最终确认。",
                },
                {
                        Key:     "no_consumer_layout_contract",
                        Label:   "不要求消费侧项目复制 PSCO 目录结构",
                        Summary: "当前阶段不把消费侧项目目录结构或固定文件名上升为必要输入合同。",
                },
                {
                        Key:     "no_second_fact_source",
                        Label:   "不形成第二套事实源",
                        Summary: "Markdown 导出只能从结构化只读结果单向派生，不形成独立于 canonical 数据的第二套事实判断。",
                },
        }
}
