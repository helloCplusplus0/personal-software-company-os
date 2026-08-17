package repository

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/psco/backend/internal/governanceprofile"
)

func TestSaveProfileRefreshesRootFrozenReadOnlyProjection(t *testing.T) {
	backendRoot := mustBackendRoot(t)
	databaseURL := mustDatabaseURL(t, backendRoot)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pool, err := newTestDBPool(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open db pool: %v", err)
	}
	t.Cleanup(pool.Close)

	repositoryID := uuid.NewString()
	repositoryName := "governance-profile-test-" + repositoryID[:8]
	insertTestRepository(t, pool, repositoryID, repositoryName)
	t.Cleanup(func() {
		cleanupTestRepository(t, pool, repositoryID)
	})

	store := NewProfileStore(pool)
	summary := "phase13 project rules summary"

	firstResult, err := store.SaveProfile(ctx, governanceprofile.UpdateGovernanceProfileInput{
		RepositoryID:   repositoryID,
		TemplateSource: stringPtr("manual://seed"),
		CanonicalRootFiles: []governanceprofile.CanonicalRootFileBinding{
			{FileName: "plan.md", Role: "plan", Required: true},
		},
		GlobalAssetBindings: []governanceprofile.GlobalAssetBinding{
			{
				Name:              "project_rules.md",
				Kind:              "rules",
				EntryRef:          "project_rules.md",
				Role:              "rules",
				StructuredSummary: &summary,
			},
		},
	})
	if err != nil {
		t.Fatalf("first save profile: %v", err)
	}
	assertRootFrozenProjection(t, firstResult)
	if len(firstResult.GlobalAssetBindings) != 1 || !firstResult.GlobalAssetBindings[0].MarkdownResolvable {
		t.Fatalf("expected markdown_resolvable to be true after readback, got %+v", firstResult.GlobalAssetBindings)
	}

	_, err = pool.Exec(ctx, `
		UPDATE governance_profiles
		SET track_type = 'product',
		    current_phase_name = 'stale_phase',
		    current_phase_ref = 'stale.md',
		    current_phase_status = 'blocked'
		WHERE repository_id = $1
	`, repositoryID)
	if err != nil {
		t.Fatalf("inject stale read-only projection: %v", err)
	}

	updatedSummary := "updated summary"
	secondResult, err := store.SaveProfile(ctx, governanceprofile.UpdateGovernanceProfileInput{
		RepositoryID:   repositoryID,
		TemplateSource: stringPtr("manual://updated"),
		CanonicalRootFiles: []governanceprofile.CanonicalRootFileBinding{
			{FileName: "plan.md", Role: "plan", Required: true},
		},
		GlobalAssetBindings: []governanceprofile.GlobalAssetBinding{
			{
				Name:              "project_rules.md",
				Kind:              "rules",
				EntryRef:          "project_rules.md",
				Role:              "rules",
				StructuredSummary: &updatedSummary,
			},
		},
	})
	if err != nil {
		t.Fatalf("second save profile: %v", err)
	}
	assertRootFrozenProjection(t, secondResult)

	readBack, err := store.ReadProfile(ctx, repositoryID)
	if err != nil {
		t.Fatalf("read profile after update: %v", err)
	}
	assertRootFrozenProjection(t, readBack)
	if len(readBack.GlobalAssetBindings) != 1 || !readBack.GlobalAssetBindings[0].MarkdownResolvable {
		t.Fatalf("expected markdown_resolvable to remain true after persisted read, got %+v", readBack.GlobalAssetBindings)
	}
}

func assertRootFrozenProjection(t *testing.T, result *governanceprofile.GovernanceProfileReadResult) {
	t.Helper()

	if result.Record.TrackType != governanceprofile.RootFrozenTrackType {
		t.Fatalf("expected track_type %q, got %q", governanceprofile.RootFrozenTrackType, result.Record.TrackType)
	}
	if result.Record.CurrentPhaseName != governanceprofile.RootFrozenCurrentPhaseName {
		t.Fatalf("expected current_phase_name %q, got %q", governanceprofile.RootFrozenCurrentPhaseName, result.Record.CurrentPhaseName)
	}
	if result.Record.CurrentPhaseRef != governanceprofile.RootFrozenCurrentPhaseRef {
		t.Fatalf("expected current_phase_ref %q, got %q", governanceprofile.RootFrozenCurrentPhaseRef, result.Record.CurrentPhaseRef)
	}
	if result.Record.CurrentPhaseStatus != governanceprofile.RootFrozenCurrentPhaseStatus {
		t.Fatalf("expected current_phase_status %q, got %q", governanceprofile.RootFrozenCurrentPhaseStatus, result.Record.CurrentPhaseStatus)
	}
	if result.Record.DocsWorkflowLayout != governanceprofile.RootFrozenDocsWorkflowLayout {
		t.Fatalf("expected docs_workflow_layout %q, got %q", governanceprofile.RootFrozenDocsWorkflowLayout, result.Record.DocsWorkflowLayout)
	}
}

func insertTestRepository(t *testing.T, pool *pgxpool.Pool, repositoryID, repositoryName string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pool.Exec(ctx, `
		INSERT INTO repositories (id, name, url, provider, status)
		VALUES ($1, $2, $3, $4, $5)
	`, repositoryID, repositoryName, "https://example.com/"+repositoryName, "github", "active")
	if err != nil {
		t.Fatalf("insert test repository: %v", err)
	}
}

func cleanupTestRepository(t *testing.T, pool *pgxpool.Pool, repositoryID string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := pool.Exec(ctx, `DELETE FROM governance_canonical_root_file_bindings WHERE governance_profile_id IN (SELECT id FROM governance_profiles WHERE repository_id = $1)`, repositoryID); err != nil {
		t.Fatalf("cleanup canonical root file bindings: %v", err)
	}
	if _, err := pool.Exec(ctx, `DELETE FROM governance_global_asset_bindings WHERE governance_profile_id IN (SELECT id FROM governance_profiles WHERE repository_id = $1)`, repositoryID); err != nil {
		t.Fatalf("cleanup global asset bindings: %v", err)
	}
	if _, err := pool.Exec(ctx, `DELETE FROM governance_profiles WHERE repository_id = $1`, repositoryID); err != nil {
		t.Fatalf("cleanup governance profile: %v", err)
	}
	if _, err := pool.Exec(ctx, `DELETE FROM repositories WHERE id = $1`, repositoryID); err != nil {
		t.Fatalf("cleanup repository: %v", err)
	}
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

func stringPtr(value string) *string {
	return &value
}
