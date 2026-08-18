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

// phase14-09 收缩：画像保存写路径与两组 bindings 读取已退役，
// 本文件只保留主表读路径覆盖（ReadProfile 三组字段 + 未创建语义）。
// fixture 通过直接 SQL 写入 governance_profiles 主表（不依赖已删除的写方法，
// 也不引用两张 bindings 表——集成测试与两表存废状态解耦）。

func TestReadProfileReturnsCoreFieldsOnly(t *testing.T) {
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

	templateSource := "manual://core-read"
	insertTestGovernanceProfile(t, pool, repositoryID, templateSource)

	store := NewProfileStore(pool)
	result, err := store.ReadProfile(ctx, repositoryID)
	if err != nil {
		t.Fatalf("read profile: %v", err)
	}

	if result.TrackType != governanceprofile.TrackTypeDurableSystem {
		t.Fatalf("expected track_type %q, got %q", governanceprofile.TrackTypeDurableSystem, result.TrackType)
	}
	if result.TemplateSource == nil || *result.TemplateSource != templateSource {
		t.Fatalf("expected template_source %q, got %+v", templateSource, result.TemplateSource)
	}
	if result.CurrentPhaseName != "phase13_project_governance_profile_foundation" {
		t.Fatalf("unexpected current_phase_name: %q", result.CurrentPhaseName)
	}
	if result.CurrentPhaseRef != "plan.md#phase13_project_governance_profile_foundation" {
		t.Fatalf("unexpected current_phase_ref: %q", result.CurrentPhaseRef)
	}
	if result.CurrentPhaseStatus != governanceprofile.PhaseStatusInProgress {
		t.Fatalf("expected current_phase_status %q, got %q", governanceprofile.PhaseStatusInProgress, result.CurrentPhaseStatus)
	}
}

func TestReadProfileAllowsNullTemplateSource(t *testing.T) {
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

	insertTestGovernanceProfile(t, pool, repositoryID, "")

	store := NewProfileStore(pool)
	result, err := store.ReadProfile(ctx, repositoryID)
	if err != nil {
		t.Fatalf("read profile: %v", err)
	}
	if result.TemplateSource != nil {
		t.Fatalf("expected nil template_source, got %+v", result.TemplateSource)
	}
}

func TestReadProfileNotCreatedReturnsNotFound(t *testing.T) {
	backendRoot := mustBackendRoot(t)
	databaseURL := mustDatabaseURL(t, backendRoot)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pool, err := newTestDBPool(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open db pool: %v", err)
	}
	t.Cleanup(pool.Close)

	store := NewProfileStore(pool)
	_, err = store.ReadProfile(ctx, uuid.NewString())
	if err == nil {
		t.Fatal("expected not found error")
	}
	if err != governanceprofile.ErrGovernanceProfileNotFound {
		t.Fatalf("expected ErrGovernanceProfileNotFound, got %v", err)
	}
}

// insertTestGovernanceProfile 直接经 SQL 写入主表 fixture
// （templateSource 为空串时写入 NULL，覆盖 optional 语义）。
func insertTestGovernanceProfile(t *testing.T, pool *pgxpool.Pool, repositoryID, templateSource string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pool.Exec(ctx, `
		INSERT INTO governance_profiles (
			repository_id, project_profile_version, track_type, template_source,
			docs_workflow_layout, current_phase_name, current_phase_ref, current_phase_status
		) VALUES ($1, $2, $3, NULLIF($4, ''), $5, $6, $7, $8)
	`, repositoryID,
		"project_governance_profile_v1",
		string(governanceprofile.TrackTypeDurableSystem),
		templateSource,
		"phase/fix/audit/review",
		"phase13_project_governance_profile_foundation",
		"plan.md#phase13_project_governance_profile_foundation",
		string(governanceprofile.PhaseStatusInProgress),
	)
	if err != nil {
		t.Fatalf("insert test governance profile: %v", err)
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
