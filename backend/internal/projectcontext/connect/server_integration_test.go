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
	projectcontextcandidate "github.com/psco/backend/internal/projectcontext/candidate"
	projectcontextservice "github.com/psco/backend/internal/projectcontext/service"
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
	querySvc := projectcontextservice.NewQueryService(readers)

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
