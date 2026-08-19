// Package service — progress 集成测试（phase15-06 Task 7）。
//
// 模式沿袭 standard 集成测试冻结模式（phase14-07）：
//   - 真实 PostgreSQL（进程环境变量 DATABASE_URL 优先，回落 backend/.env）
//   - 每个测试独立 fixture（uuid 后缀 repositories 行）+ t.Cleanup 显式清理
//     （progress_events 行先于 repositories 行删除，规避 FK ON DELETE RESTRICT），
//     测试之间互不依赖执行顺序
//   - DATABASE_URL 无法获得时 t.Skipf 跳过（SKIP 非 FAIL）；
//     配置了 DATABASE_URL 但连接失败仍 Fatal（环境故障显式暴露）
//
// 断言范围（phase15-06 spec 冻结）：
//   1. Create→List round-trip：三键链倒序（occurred_at DESC, created_at DESC,
//      id DESC）——含补录历史事件排在正确位置与同 occurred_at 的 created_at tiebreak
//   2. 三轨过滤：workflow_type 各轨（phase/audit/fix）+ 不过滤（nil）返回全量
//   3. Delete 后 List 不含该事件；Delete 不存在 id → ErrProgressEventNotFound
//   4. Create 错误分支：repository 不存在 → InvalidArgument 含 [REPOSITORY_NOT_FOUND]；
//      校验失败码抽查（TASK_KEY_FORMAT_INVALID / EVENT_KIND_NOT_ALLOWED）
//   5. List 读锚点：repository 不存在 → ErrRepositoryNotFound（NotFound 语义）
//
// 文件落点：backend/internal/progress/service/service_integration_test.go
package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/psco/backend/internal/progress"
	"github.com/psco/backend/internal/progress/candidate"
	"github.com/psco/backend/internal/progress/repository"
)

// ============================================================================
// harness 与 fixture helper
// ============================================================================

// progressIntegrationHarness 聚合被测 service 与底层连接池。
type progressIntegrationHarness struct {
	pool  *pgxpool.Pool
	cmd   *CommandService
	query *QueryService
}

// newProgressIntegrationHarness 建立真实 DB 连接并装配被测双 service。
func newProgressIntegrationHarness(t *testing.T) *progressIntegrationHarness {
	t.Helper()

	backendRoot := mustBackendRoot(t)
	databaseURL := mustDatabaseURL(t, backendRoot)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pool, err := newTestDBPool(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open db pool: %v", err)
	}
	t.Cleanup(pool.Close)

	store := repository.NewProgressEventStore(pool)
	repositoryReader := candidate.NewRepositoryReader(pool)
	return &progressIntegrationHarness{
		pool:  pool,
		cmd:   NewCommandService(repositoryReader, store),
		query: NewQueryService(store, repositoryReader),
	}
}

// insertProgressFixtureRepository 建立 repositories 表 fixture 行（uuid 后缀），
// 返回 repository id。注册 cleanup：先删 progress_events 行再删 repositories 行
// （progress_events.repository_id FK ON DELETE RESTRICT，逆序清理）。
func insertProgressFixtureRepository(t *testing.T, pool *pgxpool.Pool) string {
	t.Helper()

	repositoryID := uuid.NewString()
	repositoryName := "progress-it-" + repositoryID[:8]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pool.Exec(ctx, `
		INSERT INTO repositories (id, name, url, provider, status)
		VALUES ($1, $2, $3, $4, $5)
	`, repositoryID, repositoryName, "https://example.com/"+repositoryName, "github", "active")
	if err != nil {
		t.Fatalf("insert test repository: %v", err)
	}

	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		if _, err := pool.Exec(cleanupCtx, `DELETE FROM progress_events WHERE repository_id = $1`, repositoryID); err != nil {
			t.Fatalf("cleanup progress events: %v", err)
		}
		if _, err := pool.Exec(cleanupCtx, `DELETE FROM repositories WHERE id = $1`, repositoryID); err != nil {
			t.Fatalf("cleanup repository: %v", err)
		}
	})
	return repositoryID
}

// mustCreateEvent 经 CommandService 创建事件并断言成功，返回创建结果。
func (h *progressIntegrationHarness) mustCreateEvent(
	t *testing.T,
	repositoryID string,
	workflowType progress.WorkflowType,
	eventKind progress.EventKind,
	taskKey, title string,
	occurredAt time.Time,
) *progress.ProgressEventReadResult {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	created, err := h.cmd.CreateProgressEvent(ctx, &progress.CreateProgressEventInput{
		RepositoryID: repositoryID,
		WorkflowType: workflowType,
		EventKind:    eventKind,
		TaskKey:      taskKey,
		Title:        title,
		Source:       progress.ProgressSourceManual,
		OccurredAt:   occurredAt,
	})
	if err != nil {
		t.Fatalf("create progress event %q: %v", title, err)
	}
	return created
}

// assertInvalidInput 断言错误为 ErrInvalidInput 哨兵且信息包含关键字（错误码）。
func assertInvalidInput(t *testing.T, err error, msgContains string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q, got nil", msgContains)
	}
	if !errors.Is(err, progress.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
	if !strings.Contains(err.Error(), msgContains) {
		t.Fatalf("expected error message containing %q, got %q", msgContains, err.Error())
	}
}

// eventTitles 返回事件切片的 title 序（失败信息定位用）。
func eventTitles(events []progress.ProgressEventReadResult) []string {
	titles := make([]string, 0, len(events))
	for _, event := range events {
		titles = append(titles, event.Title)
	}
	return titles
}

// ============================================================================
// 断言组 1：Create→List round-trip（三键链倒序）
// ============================================================================

// TestCreateAndListProgressEventsRoundTrip 验证多事件写入后的完整回读：
// 跨三轨 7 事件（含补录历史事件与同 occurred_at 双事件），List 断言
// 三键链倒序（occurred_at DESC 主导 + 同刻 created_at DESC tiebreak +
// 补录历史事件按声明时间排在正确位置），并抽查字段 round-trip。
func TestCreateAndListProgressEventsRoundTrip(t *testing.T) {
	h := newProgressIntegrationHarness(t)
	repositoryID := insertProgressFixtureRepository(t, h.pool)

	now := time.Now().UTC()

	// 补录历史事件：声明发生时间早于录入时间 72h（occurred_at 与 created_at 分离语义）。
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindPhaseStarted,
		"phase14", "补录历史：phase14 开始", now.Add(-72*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypeAudit, progress.EventKindTaskCompleted,
		"audit_001", "审计任务完成", now.Add(-30*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypeFix, progress.EventKindNote,
		"", "修复轨备注（task_key 可空）", now.Add(-20*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindPhaseStarted,
		"phase15", "phase15 开始", now.Add(-10*time.Hour))

	// 同 occurred_at 双事件：先录入 A，sleep 后录入 B（created_at 严格区分，
	// 三键链第二键 created_at DESC 的稳定观测位）。
	tieOccurredAt := now.Add(-5 * time.Hour)
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindTaskCompleted,
		"phase15-05", "同刻任务 A（先录入）", tieOccurredAt)
	time.Sleep(20 * time.Millisecond)
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindTaskCompleted,
		"phase15-06", "同刻任务 B（后录入）", tieOccurredAt)

	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindNote,
		"", "阶段进行中备注", now.Add(-1*time.Hour))

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	events, err := h.query.ListProgressEvents(ctx, repositoryID, nil)
	if err != nil {
		t.Fatalf("list progress events: %v", err)
	}
	if len(events) != 7 {
		t.Fatalf("expected 7 events, got %d", len(events))
	}

	// 三键链倒序断言：occurred_at DESC 主导的期望 title 序。
	wantTitles := []string{
		"阶段进行中备注",           // now-1h（最新）
		"同刻任务 B（后录入）",       // now-5h，created_at 较晚 → 同刻排前
		"同刻任务 A（先录入）",       // now-5h，created_at 较早
		"phase15 开始",            // now-10h
		"修复轨备注（task_key 可空）", // now-20h
		"审计任务完成",              // now-30h
		"补录历史：phase14 开始",     // now-72h（补录历史排在正确位置：最早）
	}
	for i, want := range wantTitles {
		if events[i].Title != want {
			t.Fatalf("event[%d] title mismatch: got %q, want %q (full order: %v)",
				i, events[i].Title, want, eventTitles(events))
		}
	}

	// 同 occurred_at 双事件：后录入（created_at 较晚）排在前（created_at DESC tiebreak）。
	if !events[1].CreatedAt.After(events[2].CreatedAt) {
		t.Fatalf("expected tie pair ordered by created_at DESC: got %v before %v",
			events[1].CreatedAt, events[2].CreatedAt)
	}

	// 字段 round-trip 抽查（最新事件）。
	latest := events[0]
	if latest.RepositoryID != repositoryID {
		t.Fatalf("latest event repository_id mismatch: got %q, want %q", latest.RepositoryID, repositoryID)
	}
	if latest.WorkflowType != progress.WorkflowTypePhase || latest.EventKind != progress.EventKindNote {
		t.Fatalf("latest event enums mismatch: got workflow %q kind %q", latest.WorkflowType, latest.EventKind)
	}
	if latest.Source != progress.ProgressSourceManual {
		t.Fatalf("latest event source mismatch: got %q", latest.Source)
	}
	if _, err := uuid.Parse(latest.ID); err != nil {
		t.Fatalf("latest event id not a valid UUID: %q", latest.ID)
	}
	if latest.CreatedAt.IsZero() {
		t.Fatal("expected latest event created_at set")
	}

	// 补录历史语义：声明发生时间（occurred_at）早于系统录入时间（created_at）。
	historical := events[6]
	if historical.TaskKey != "phase14" {
		t.Fatalf("historical event task_key mismatch: got %q, want %q", historical.TaskKey, "phase14")
	}
	if !historical.OccurredAt.Before(historical.CreatedAt) {
		t.Fatalf("expected backfilled event occurred_at %v before created_at %v",
			historical.OccurredAt, historical.CreatedAt)
	}
}

// ============================================================================
// 断言组 2：三轨过滤
// ============================================================================

// TestListProgressEventsWorkflowFilter 验证三轨过滤：
// workflow_type 各轨（phase/audit/fix）过滤后仅含该轨事件且保持倒序；
// 不过滤（nil）返回三轨全量。
func TestListProgressEventsWorkflowFilter(t *testing.T) {
	h := newProgressIntegrationHarness(t)
	repositoryID := insertProgressFixtureRepository(t, h.pool)

	now := time.Now().UTC()
	// 跨三轨各 2 条（occurred_at 互异，单轨倒序可稳定断言）。
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindPhaseStarted,
		"phase15", "phase 开始", now.Add(-8*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindTaskCompleted,
		"phase15-01", "phase 任务完成", now.Add(-7*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypeAudit, progress.EventKindTaskCompleted,
		"audit_001", "audit 任务完成", now.Add(-6*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypeAudit, progress.EventKindNote,
		"", "audit 备注", now.Add(-5*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypeFix, progress.EventKindTaskCompleted,
		"fix_001", "fix 任务完成", now.Add(-4*time.Hour))
	h.mustCreateEvent(t, repositoryID, progress.WorkflowTypeFix, progress.EventKindNote,
		"", "fix 备注", now.Add(-3*time.Hour))

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 不过滤（nil）：三轨全量。
	all, err := h.query.ListProgressEvents(ctx, repositoryID, nil)
	if err != nil {
		t.Fatalf("list all progress events: %v", err)
	}
	if len(all) != 6 {
		t.Fatalf("expected 6 events without filter, got %d", len(all))
	}

	// 各轨过滤：仅含该轨事件，且按 occurred_at DESC 倒序（task_key 序断言）。
	cases := []struct {
		workflow progress.WorkflowType
		wantKeys []string
	}{
		{progress.WorkflowTypePhase, []string{"phase15-01", "phase15"}},
		{progress.WorkflowTypeAudit, []string{"", "audit_001"}},
		{progress.WorkflowTypeFix, []string{"", "fix_001"}},
	}
	for _, tc := range cases {
		filtered, err := h.query.ListProgressEvents(ctx, repositoryID, &tc.workflow)
		if err != nil {
			t.Fatalf("list progress events filtered by %q: %v", tc.workflow, err)
		}
		if len(filtered) != len(tc.wantKeys) {
			t.Fatalf("workflow %q: expected %d events, got %d", tc.workflow, len(tc.wantKeys), len(filtered))
		}
		for i, event := range filtered {
			if event.WorkflowType != tc.workflow {
				t.Fatalf("workflow %q: event[%d] workflow_type mismatch: got %q", tc.workflow, i, event.WorkflowType)
			}
			if event.TaskKey != tc.wantKeys[i] {
				t.Fatalf("workflow %q: event[%d] task_key mismatch: got %q, want %q",
					tc.workflow, i, event.TaskKey, tc.wantKeys[i])
			}
		}
	}
}

// ============================================================================
// 断言组 3：Delete 语义
// ============================================================================

// TestDeleteProgressEventSemantics 验证删除主线与 NotFound 分支：
// Delete 不存在 id → ErrProgressEventNotFound；Delete 成功后 List 不含该事件；
// 重复 Delete 已删 id → ErrProgressEventNotFound。
func TestDeleteProgressEventSemantics(t *testing.T) {
	h := newProgressIntegrationHarness(t)
	repositoryID := insertProgressFixtureRepository(t, h.pool)

	now := time.Now().UTC()
	first := h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindPhaseStarted,
		"phase15", "待删除的阶段开始", now.Add(-2*time.Hour))
	second := h.mustCreateEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindTaskCompleted,
		"phase15-06", "保留的任务完成", now.Add(-1*time.Hour))

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Delete 不存在 id（合法 UUID 但表中无行）→ ErrProgressEventNotFound（NotFound）。
	if err := h.cmd.DeleteProgressEvent(ctx, uuid.NewString()); !errors.Is(err, progress.ErrProgressEventNotFound) {
		t.Fatalf("expected ErrProgressEventNotFound, got %v", err)
	}

	// Delete 成功后 List 不含该事件（仅剩未删除事件）。
	if err := h.cmd.DeleteProgressEvent(ctx, first.ID); err != nil {
		t.Fatalf("delete progress event %s: %v", first.ID, err)
	}
	events, err := h.query.ListProgressEvents(ctx, repositoryID, nil)
	if err != nil {
		t.Fatalf("list progress events after delete: %v", err)
	}
	if len(events) != 1 || events[0].ID != second.ID {
		t.Fatalf("expected only event %s after delete, got %d events: %+v", second.ID, len(events), eventTitles(events))
	}

	// 重复 Delete 已删 id → ErrProgressEventNotFound。
	if err := h.cmd.DeleteProgressEvent(ctx, first.ID); !errors.Is(err, progress.ErrProgressEventNotFound) {
		t.Fatalf("expected ErrProgressEventNotFound on repeat delete, got %v", err)
	}
}

// ============================================================================
// 断言组 4：Create 错误分支与 List 读锚点
// ============================================================================

// TestCreateProgressEventErrorSemantics 抽查创建写路径错误语义：
// repository 不存在 → InvalidArgument 含 [REPOSITORY_NOT_FOUND]；
// 校验失败码抽查（TASK_KEY_FORMAT_INVALID / EVENT_KIND_NOT_ALLOWED）。
func TestCreateProgressEventErrorSemantics(t *testing.T) {
	h := newProgressIntegrationHarness(t)
	repositoryID := insertProgressFixtureRepository(t, h.pool)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	now := time.Now().UTC()

	// repository 不存在（合法 UUID 但表中无行）→ ErrInvalidInput 含 [REPOSITORY_NOT_FOUND]
	// （InvalidArgument 语义，沿 standard "target 不存在归 InvalidArgument" 模式）。
	_, err := h.cmd.CreateProgressEvent(ctx, &progress.CreateProgressEventInput{
		RepositoryID: uuid.NewString(),
		WorkflowType: progress.WorkflowTypePhase,
		EventKind:    progress.EventKindPhaseStarted,
		TaskKey:      "phase15",
		Title:        "仓库不存在的录入",
		Source:       progress.ProgressSourceManual,
		OccurredAt:   now,
	})
	assertInvalidInput(t, err, "[REPOSITORY_NOT_FOUND]")

	// 校验失败码抽查 1：TASK_KEY_FORMAT_INVALID——phase 轨 task_completed
	// 必须匹配 K-2（phaseNN-MM），"phase15" 缺任务段。
	_, err = h.cmd.CreateProgressEvent(ctx, &progress.CreateProgressEventInput{
		RepositoryID: repositoryID,
		WorkflowType: progress.WorkflowTypePhase,
		EventKind:    progress.EventKindTaskCompleted,
		TaskKey:      "phase15",
		Title:        "缺任务段的 task_key",
		Source:       progress.ProgressSourceManual,
		OccurredAt:   now,
	})
	assertInvalidInput(t, err, "[TASK_KEY_FORMAT_INVALID]")

	// 校验失败码抽查 2：EVENT_KIND_NOT_ALLOWED——audit 轨禁止 phase 边界标记
	// （组合矩阵非法先于 task_key 判定拦截）。
	_, err = h.cmd.CreateProgressEvent(ctx, &progress.CreateProgressEventInput{
		RepositoryID: repositoryID,
		WorkflowType: progress.WorkflowTypeAudit,
		EventKind:    progress.EventKindPhaseStarted,
		TaskKey:      "audit_001",
		Title:        "audit 轨的非法 phase_started",
		Source:       progress.ProgressSourceManual,
		OccurredAt:   now,
	})
	assertInvalidInput(t, err, "[EVENT_KIND_NOT_ALLOWED]")
}

// TestListProgressEventsRepositoryNotFound 验证 List 读锚点语义：
// 仓库不存在 → ErrRepositoryNotFound（NotFound，沿 GetProjectBrief 读锚点）。
func TestListProgressEventsRepositoryNotFound(t *testing.T) {
	h := newProgressIntegrationHarness(t)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := h.query.ListProgressEvents(ctx, uuid.NewString(), nil)
	if err == nil || !errors.Is(err, progress.ErrRepositoryNotFound) {
		t.Fatalf("expected ErrRepositoryNotFound, got %v", err)
	}
}

// ============================================================================
// DB 连接 helper（沿 standard 集成测试冻结模式）
// ============================================================================

// mustBackendRoot 自当前工作目录向上定位 backend 根（含 go.mod 的目录）。
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

// mustDatabaseURL 读取 DATABASE_URL：进程环境变量优先，回落 backend/.env。
// 无法获得时跳过（SKIP 非 FAIL，沿 standard 集成测试冻结模式）。
func mustDatabaseURL(t *testing.T, backendRoot string) string {
	t.Helper()

	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		return databaseURL
	}

	envPath := filepath.Join(backendRoot, ".env")
	if err := godotenv.Load(envPath); err != nil {
		t.Skipf("skip progress integration test: load %s: %v", envPath, err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skipf("skip progress integration test: DATABASE_URL is empty after loading %s", envPath)
	}
	return databaseURL
}

// newTestDBPool 建立 pgx 连接池并 ping 验证连通。
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
