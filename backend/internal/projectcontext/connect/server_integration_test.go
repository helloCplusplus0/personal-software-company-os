package connect

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	connectrpc "connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	pbc "github.com/psco/backend/internal/gen/connect/psco/project_context/v1/project_contextv1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/project_context/v1"
	standardpb "github.com/psco/backend/internal/gen/proto/psco/standard/v1"
	governanceprofilerepo "github.com/psco/backend/internal/governanceprofile/repository"
	governanceprofileservice "github.com/psco/backend/internal/governanceprofile/service"
	projectcontextcandidate "github.com/psco/backend/internal/projectcontext/candidate"
	projectcontextservice "github.com/psco/backend/internal/projectcontext/service"
	"github.com/psco/backend/internal/standard"
	standardcandidate "github.com/psco/backend/internal/standard/candidate"
	standardrepo "github.com/psco/backend/internal/standard/repository"
	standardservice "github.com/psco/backend/internal/standard/service"
)

func TestProjectContextAcceptanceScenarios(t *testing.T) {
	// 本组测试通过真实 fixture 重置共享数据库状态，因此必须串行执行。
	h := newProjectContextIntegrationHarness(t)

	t.Run("repository not found returns not found", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		_, err := h.client.GetProjectContext(t.Context(), &pb.GetProjectContextRequest{
			RepositoryId: "00000000-0000-0000-0000-000000000001",
		})
		if err == nil {
			t.Fatal("expected not found error")
		}
		if got := connectrpc.CodeOf(err); got != connectrpc.CodeNotFound {
			t.Fatalf("expected CodeNotFound, got %v", got)
		}
	})

	t.Run("binding incomplete returns failed precondition", func(t *testing.T) {
		h.resetFixture(t, "completed-unbound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
		_, err := h.client.GetProjectContext(t.Context(), &pb.GetProjectContextRequest{
			RepositoryId: repositoryID,
		})
		if err == nil {
			t.Fatal("expected failed precondition error")
		}
		if got := connectrpc.CodeOf(err); got != connectrpc.CodeFailedPrecondition {
			t.Fatalf("expected CodeFailedPrecondition, got %v", got)
		}
	})

	t.Run("completed binding returns aggregated context", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
		resp, err := h.client.GetProjectContext(t.Context(), &pb.GetProjectContextRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		if resp.GetRepository() == nil || resp.GetRepository().GetName() != "main-repo" {
			t.Fatalf("expected repository main-repo, got %+v", resp.GetRepository())
		}
		if resp.GetProduct() == nil || resp.GetProduct().GetName() != "Product A" {
			t.Fatalf("expected product Product A, got %+v", resp.GetProduct())
		}
		if len(resp.GetModules()) < 2 {
			t.Fatalf("expected at least 2 modules, got %d", len(resp.GetModules()))
		}
		if len(resp.GetDecisions()) == 0 {
			t.Fatal("expected aggregated decisions")
		}
		if len(resp.GetRules()) == 0 {
			t.Fatal("expected rule entries")
		}
		if len(resp.GetPhases()) == 0 {
			t.Fatal("expected phase entries")
		}
		if len(resp.GetBoundaries()) == 0 {
			t.Fatal("expected boundary entries")
		}

		decision := resp.GetDecisions()[0]
		if decision.GetTitle() != "phase06 completed-bound 验收决策" {
			t.Fatalf("unexpected decision title: %s", decision.GetTitle())
		}
		if !containsAll(decision.GetHitSources(), "repository_module_mapping", "bound_product_module") {
			t.Fatalf("expected both hit sources, got %v", decision.GetHitSources())
		}

		for _, rule := range resp.GetRules() {
			if rule.GetEntryRef() == "" || rule.GetEntryKind() == "" {
				t.Fatalf("expected populated rule entry locator, got %+v", rule)
			}
		}
		requiredRuleSummaries := []string{
			"Product = 经营目标与交付容器",
			"Repository = 代码仓库身份对象与项目锚点",
		}
		for _, summary := range requiredRuleSummaries {
			if !ruleSummariesContain(resp.GetRules(), summary) {
				t.Fatalf("expected rules to contain summary %q, got %+v", summary, resp.GetRules())
			}
		}
		requiredRuleRefs := []string{"plan.md", "architecture_map.md", "docs/README.md"}
		for _, entryRef := range requiredRuleRefs {
			if !ruleRefsContain(resp.GetRules(), entryRef) {
				t.Fatalf("expected rules to contain entry ref %q, got %+v", entryRef, resp.GetRules())
			}
		}
		for _, phase := range resp.GetPhases() {
			if phase.GetEntryRef() == "" || phase.GetEntryKind() == "" {
				t.Fatalf("expected populated phase entry locator, got %+v", phase)
			}
		}
		if !phaseEntriesContain(resp.GetPhases(), "phase12", "docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md") {
			t.Fatalf("expected phases to contain phase12 entry, got %+v", resp.GetPhases())
		}
	})

	t.Run("export project context returns markdown derived from structured result", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
		resp, err := h.client.ExportProjectContext(t.Context(), &pb.ExportProjectContextRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		markdown := resp.GetMarkdown()
		requiredContent := []string{
			"# Project Context",
			"Product = 经营目标与交付容器",
			"Repository = 代码仓库身份对象与项目锚点",
			"## Rules & Constraints",
			"plan.md",
			"architecture_map.md",
			"docs/README.md",
			"project_rules.md",
			"## Current Phase",
			"docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md",
			"docs/phase/phase11_project_context_foundation_dev_plan.md",
			"## Boundaries (What This Project Does NOT Do)",
			"不形成第二套事实源",
		}
		for _, item := range requiredContent {
			if !strings.Contains(markdown, item) {
				t.Fatalf("expected markdown to contain %q, got:\n%s", item, markdown)
			}
		}
	})

	t.Run("project brief repository not found returns not found", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		_, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: "00000000-0000-0000-0000-000000000001",
		})
		if err == nil {
			t.Fatal("expected not found error")
		}
		if got := connectrpc.CodeOf(err); got != connectrpc.CodeNotFound {
			t.Fatalf("expected CodeNotFound, got %v", got)
		}
	})

	t.Run("project brief profile not found returns not found", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
		_, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err == nil {
			t.Fatal("expected not found error")
		}
		if got := connectrpc.CodeOf(err); got != connectrpc.CodeNotFound {
			t.Fatalf("expected CodeNotFound, got %v", got)
		}
	})

	t.Run("project brief returns inline brief blocks when governance profile exists", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
		h.mustSeedGovernanceProfileRow(t, repositoryID)

		resp, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		if resp.GetRepository() == nil || resp.GetRepository().GetName() != "main-repo" {
			t.Fatalf("expected repository main-repo, got %+v", resp.GetRepository())
		}
		if resp.GetGovernanceProfile() == nil {
			t.Fatal("expected governance profile block")
		}
		if resp.GetCurrentPhase() == nil {
			t.Fatal("expected current phase block")
		}
		if resp.GetCurrentPhase().GetStatus() != pb.BriefPhaseStatus_BRIEF_PHASE_STATUS_IN_PROGRESS {
			t.Fatalf("expected current phase in progress, got %v", resp.GetCurrentPhase().GetStatus())
		}
		if len(resp.GetProducts()) != 1 {
			t.Fatalf("expected products array length 1, got %d", len(resp.GetProducts()))
		}
		if len(resp.GetModules()) == 0 {
			t.Fatal("expected modules array")
		}
		if len(resp.GetDecisions()) == 0 {
			t.Fatal("expected decisions array")
		}

		// phase14-09 内联切换断言：BriefGovernanceProfile 三字段 round-trip
		// （repository_id / track_type / template_source）。
		profile := resp.GetGovernanceProfile()
		if profile.GetRepositoryId() != repositoryID {
			t.Fatalf("expected repository_id round-trip, got %q want %q", profile.GetRepositoryId(), repositoryID)
		}
		if profile.GetTemplateSource() != "manual://brief-test" {
			t.Fatalf("expected template_source round-trip, got %q", profile.GetTemplateSource())
		}
		if profile.GetTrackType() != pb.BriefTrackType_BRIEF_TRACK_TYPE_DURABLE_SYSTEM {
			t.Fatalf("expected durable system track, got %v", profile.GetTrackType())
		}

		// current_phase 从治理画像主记录 read-only 字段单向派生
		// （fixture 固定投影：phase13 阶段名 / plan.md 入口 / in_progress）。
		if resp.GetCurrentPhase().GetName() != "phase13_project_governance_profile_foundation" ||
			resp.GetCurrentPhase().GetEntryRef() != "plan.md#phase13_project_governance_profile_foundation" ||
			resp.GetCurrentPhase().GetStatus() != pb.BriefPhaseStatus_BRIEF_PHASE_STATUS_IN_PROGRESS {
			t.Fatalf("expected current phase derived from governance profile record, got %+v", resp.GetCurrentPhase())
		}
	})

	t.Run("project brief returns standards with full directory tree when standard is bound to repository", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
		h.mustSeedGovernanceProfileRow(t, repositoryID)
		h.cleanupStandardFixtures(t)
		created := h.mustCreateAndBindStandard(t, repositoryID)

		resp, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		standards := resp.GetStandards()
		if len(standards) != 1 {
			t.Fatalf("expected standards array length 1, got %d", len(standards))
		}

		// 主记录字段 round-trip（与 standard service 写入口径一致）
		got := standards[0]
		if got.GetId() != created.ID {
			t.Fatalf("expected standard id round-trip, got %q want %q", got.GetId(), created.ID)
		}
		if got.GetName() != "brief-standards-roundtrip" {
			t.Fatalf("expected standard name round-trip, got %q", got.GetName())
		}
		if got.GetDescription() != "phase14-07 brief standards round-trip fixture" {
			t.Fatalf("expected standard description round-trip, got %q", got.GetDescription())
		}
		if got.GetStatus() != standardpb.StandardStatus_STANDARD_STATUS_ACTIVE {
			t.Fatalf("expected active status round-trip, got %v", got.GetStatus())
		}
		if got.GetCreatedAt() == nil || got.GetUpdatedAt() == nil {
			t.Fatal("expected standard timestamps round-trip")
		}

		// directory_tree 全树节点字段逐项断言（根 → docs → README.md）
		root := got.GetDirectoryTree()
		if root == nil {
			t.Fatal("expected directory tree round-trip")
		}
		if root.GetName() != "." || root.GetNodeType() != standardpb.NodeType_NODE_TYPE_DIRECTORY ||
			root.GetRole() != "" || root.GetSummary() != "" || root.GetRef() != "" {
			t.Fatalf("unexpected root node fields: %+v", root)
		}
		if len(root.GetChildren()) != 1 {
			t.Fatalf("expected root children length 1, got %d", len(root.GetChildren()))
		}

		docs := root.GetChildren()[0]
		if docs.GetName() != "docs" || docs.GetNodeType() != standardpb.NodeType_NODE_TYPE_DIRECTORY ||
			docs.GetRole() != "" || docs.GetSummary() != "规范正文目录" || docs.GetRef() != "" {
			t.Fatalf("unexpected docs node fields: %+v", docs)
		}
		if len(docs.GetChildren()) != 1 {
			t.Fatalf("expected docs children length 1, got %d", len(docs.GetChildren()))
		}

		readme := docs.GetChildren()[0]
		if readme.GetName() != "README.md" || readme.GetNodeType() != standardpb.NodeType_NODE_TYPE_FILE ||
			readme.GetRole() != "standard-entry" || readme.GetSummary() != "规范入口说明" ||
			readme.GetRef() != "/docs/README.md" {
			t.Fatalf("unexpected README.md node fields: %+v", readme)
		}
		if len(readme.GetChildren()) != 0 {
			t.Fatalf("expected file node without children, got %d", len(readme.GetChildren()))
		}

		// phase14-09：旧 global_assets 顶层块已移除，两组 bindings 信息唯一来自 standards[]；
		// 内联画像块与 current_phase 保持装配不回退。
		if resp.GetGovernanceProfile() == nil || resp.GetCurrentPhase() == nil {
			t.Fatalf("expected governance profile blocks to remain populated, got %+v", resp)
		}
	})

	t.Run("project brief allows empty arrays when governance profile exists but bindings are incomplete", func(t *testing.T) {
		h.resetFixture(t, "completed-unbound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
		h.mustSeedGovernanceProfileRow(t, repositoryID)
		h.cleanupStandardFixtures(t)

		resp, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success with empty arrays, got %v", err)
		}

		if len(resp.GetProducts()) != 0 {
			t.Fatalf("expected empty products array, got %d", len(resp.GetProducts()))
		}
		if len(resp.GetModules()) != 0 {
			t.Fatalf("expected empty modules array, got %d", len(resp.GetModules()))
		}
		if len(resp.GetDecisions()) != 0 {
			t.Fatalf("expected empty decisions array, got %d", len(resp.GetDecisions()))
		}
		// 未绑定任何 Standard 时 standards 为空数组（phase14-07 冻结语义）
		if len(resp.GetStandards()) != 0 {
			t.Fatalf("expected empty standards array, got %d", len(resp.GetStandards()))
		}
		if resp.GetGovernanceProfile() == nil || resp.GetCurrentPhase() == nil {
			t.Fatalf("expected governance profile blocks to still be present, got %+v", resp)
		}
	})
}

type projectContextIntegrationHarness struct {
	repoRoot    string
	databaseURL string
	pool        *pgxpool.Pool
	client      pbc.ProjectContextServiceClient
}

func newProjectContextIntegrationHarness(t *testing.T) *projectContextIntegrationHarness {
	t.Helper()

	backendRoot := mustBackendRoot(t)
	repoRoot := filepath.Dir(backendRoot)
	databaseURL := mustDatabaseURL(t, backendRoot)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := newTestDBPool(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open db pool: %v", err)
	}
	t.Cleanup(pool.Close)

	readers := projectcontextcandidate.NewContextReaders(pool)
	// phase14-09：注入收缩后的画像主记录轻量读取主线（ReadProfileCore）
	governanceReader := governanceprofileservice.NewQueryService(
		governanceprofilerepo.NewProfileStore(pool),
	)
	// phase14-07：注入 standard 读取主线作为 GetProjectBrief.standards[] 的 candidate 依赖
	standardReader := standardservice.NewQueryService(standardrepo.NewStandardStore(pool))
	querySvc := projectcontextservice.NewQueryService(readers, governanceReader, standardReader)

	mux := http.NewServeMux()
	path, handler := pbc.NewProjectContextServiceHandler(NewServer(querySvc))
	mux.Handle(path, handler)

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	return &projectContextIntegrationHarness{
		repoRoot:    repoRoot,
		databaseURL: databaseURL,
		pool:        pool,
		client:      pbc.NewProjectContextServiceClient(server.Client(), server.URL),
	}
}

func (h *projectContextIntegrationHarness) resetFixture(t *testing.T, fixture string) {
	t.Helper()

	scriptPath := filepath.Join(h.repoRoot, "database", "scripts", "reset_phase06_acceptance.sh")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, scriptPath, "--fixture", fixture)
	cmd.Dir = h.repoRoot
	cmd.Env = fixtureCommandEnv(t, h.databaseURL)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("reset fixture %q failed: %v\n%s", fixture, err, string(output))
	}
}

func (h *projectContextIntegrationHarness) mustRepositoryIDByName(t *testing.T, name string) string {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var repositoryID string
	if err := h.pool.QueryRow(ctx, `SELECT id FROM repositories WHERE name = $1`, name).Scan(&repositoryID); err != nil {
		t.Fatalf("lookup repository %q: %v", name, err)
	}
	return repositoryID
}

// mustSeedGovernanceProfileRow 直接经 SQL 写入画像主表 fixture
// （phase14-09：画像写路径已退役，测试不再经 SaveProfile 写入；
// 也不引用两张 bindings 表——集成测试与两表存废状态解耦）。
func (h *projectContextIntegrationHarness) mustSeedGovernanceProfileRow(t *testing.T, repositoryID string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := h.pool.Exec(ctx, `
		INSERT INTO governance_profiles (
			repository_id, project_profile_version, track_type, template_source,
			docs_workflow_layout, current_phase_name, current_phase_ref, current_phase_status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (repository_id) DO UPDATE SET
			template_source      = EXCLUDED.template_source,
			current_phase_name   = EXCLUDED.current_phase_name,
			current_phase_ref    = EXCLUDED.current_phase_ref,
			current_phase_status = EXCLUDED.current_phase_status,
			updated_at           = NOW()
	`, repositoryID,
		"project_governance_profile_v1",
		"durable_system",
		"manual://brief-test",
		"phase/fix/audit/review",
		"phase13_project_governance_profile_foundation",
		"plan.md#phase13_project_governance_profile_foundation",
		"in_progress",
	)
	if err != nil {
		t.Fatalf("seed governance profile row: %v", err)
	}
}

// cleanupStandardFixtures 清理 standard 三表 fixture（按 fixture 名定位，
// 先删 bindings / revisions 再删主记录，规避外键与 UNIQUE(name) 残留）。
func (h *projectContextIntegrationHarness) cleanupStandardFixtures(t *testing.T) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// pgx 带参 Exec 走 prepared statement，不支持多命令，逐条删除
	statements := []string{
		`DELETE FROM standard_bindings WHERE standard_id IN (SELECT id FROM standards WHERE name = $1)`,
		`DELETE FROM standard_revisions WHERE standard_id IN (SELECT id FROM standards WHERE name = $1)`,
		`DELETE FROM standards WHERE name = $1`,
	}
	for _, stmt := range statements {
		if _, err := h.pool.Exec(ctx, stmt, "brief-standards-roundtrip"); err != nil {
			t.Fatalf("cleanup standard fixtures (%q): %v", stmt, err)
		}
	}
}

// mustCreateAndBindStandard 经 standard service 写入 fixture：
// 创建 active 规范（含 docs/README.md 两层树）并按 adopts role 绑定到仓库，
// 返回创建结果供 round-trip 断言对齐写入口径（phase14-07）。
func (h *projectContextIntegrationHarness) mustCreateAndBindStandard(t *testing.T, repositoryID string) *standard.StandardReadResult {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	commandSvc := standardservice.NewCommandService(
		standardcandidate.NewTargetReader(h.pool),
		standardrepo.NewStandardStore(h.pool),
	)

	tree := &standard.DirectoryTreeNode{
		Name:     ".",
		NodeType: string(standard.NodeTypeDirectory),
		Children: []*standard.DirectoryTreeNode{
			{
				Name:     "docs",
				NodeType: string(standard.NodeTypeDirectory),
				Summary:  "规范正文目录",
				Children: []*standard.DirectoryTreeNode{
					{
						Name:     "README.md",
						NodeType: string(standard.NodeTypeFile),
						Role:     "standard-entry",
						Summary:  "规范入口说明",
						Ref:      "/docs/README.md",
					},
				},
			},
		},
	}

	created, err := commandSvc.CreateStandard(ctx, standard.CreateStandardInput{
		Name:          "brief-standards-roundtrip",
		Description:   "phase14-07 brief standards round-trip fixture",
		DirectoryTree: tree,
		Status:        standard.StandardStatusActive,
	})
	if err != nil {
		t.Fatalf("create standard fixture: %v", err)
	}

	note := "bound via phase14-07 brief round-trip"
	if _, err := commandSvc.BindStandard(ctx, standard.BindStandardInput{
		StandardID: created.ID,
		TargetType: standard.BindingTargetRepository,
		TargetID:   repositoryID,
		Role:       standard.BindingRoleAdopts,
		Note:       &note,
	}); err != nil {
		t.Fatalf("bind standard fixture: %v", err)
	}
	return created
}

func mustBackendRoot(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not locate backend go.mod")
		}
		dir = parent
	}
}

func mustDatabaseURL(t *testing.T, backendRoot string) string {
	t.Helper()

	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		return databaseURL
	}

	envPath := filepath.Join(backendRoot, ".env")
	if err := godotenv.Load(envPath); err != nil {
		t.Fatalf("load %s: %v", envPath, err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Fatalf("DATABASE_URL is empty after loading %s", envPath)
	}
	return databaseURL
}

func fixtureCommandEnv(t *testing.T, databaseURL string) []string {
	t.Helper()

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		t.Fatalf("parse database url: %v", err)
	}

	port := cfg.ConnConfig.Port
	if port == 0 {
		port = 5432
	}

	return append(os.Environ(),
		fmt.Sprintf("PG_HOST=%s", cfg.ConnConfig.Host),
		fmt.Sprintf("PG_PORT=%d", port),
		fmt.Sprintf("PG_USER=%s", cfg.ConnConfig.User),
		fmt.Sprintf("PG_PASSWORD=%s", cfg.ConnConfig.Password),
		fmt.Sprintf("PSCO_DB=%s", cfg.ConnConfig.Database),
	)
}

func newTestDBPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database url: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return pool, nil
}

func containsAll(items []string, wants ...string) bool {
	for _, want := range wants {
		found := false
		for _, item := range items {
			if item == want {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func ruleSummariesContain(items []*pb.RuleEntry, summary string) bool {
	for _, item := range items {
		if item.GetSummary() == summary {
			return true
		}
	}
	return false
}

func ruleRefsContain(items []*pb.RuleEntry, entryRef string) bool {
	for _, item := range items {
		if item.GetEntryRef() == entryRef {
			return true
		}
	}
	return false
}

func phaseEntriesContain(items []*pb.PhaseEntry, phase, entryRef string) bool {
	for _, item := range items {
		if item.GetPhase() == phase && item.GetEntryRef() == entryRef {
			return true
		}
	}
	return false
}
