package connect

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	connectrpc "connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/encoding/protojson"

	pbc "github.com/psco/backend/internal/gen/connect/psco/project_context/v1/project_contextv1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/project_context/v1"
	standardpb "github.com/psco/backend/internal/gen/proto/psco/standard/v1"
	projectcontextcandidate "github.com/psco/backend/internal/projectcontext/candidate"
	projectcontextservice "github.com/psco/backend/internal/projectcontext/service"
	"github.com/psco/backend/internal/progress"
	progresscandidate "github.com/psco/backend/internal/progress/candidate"
	progressrepo "github.com/psco/backend/internal/progress/repository"
	progressservice "github.com/psco/backend/internal/progress/service"
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

	// 2026-08-18 phase14-10 T7 裁决：原「画像未创建 → not_found」用例已删除
	// （画像主表随 0012 迁移 drop，语义失效；brief 不再依赖画像主记录行）。
	t.Run("project brief returns six top-level blocks without profile remnants", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")

		resp, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		if resp.GetRepository() == nil || resp.GetRepository().GetName() != "main-repo" {
			t.Fatalf("expected repository main-repo, got %+v", resp.GetRepository())
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

		// 5 顶层块口径（2026-08-18 T7 裁决，槽位 2/3/4 reserved）：响应 JSON 序列化后不得再出现
		// governanceProfile / currentPhase / globalAssets 任何画像残余字段。
		assertNoProfileRemnants(t, resp)
	})

	t.Run("project brief returns standards with full directory tree when standard is bound to repository", func(t *testing.T) {
		h.resetFixture(t, "completed-bound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
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
		// 2026-08-18 phase14-10 T7 裁决：governance_profile / current_phase 画像残余块
		// 同步移除，brief 无画像前提不再 404，Standards 摘要经 brief 正常装配。
		assertNoProfileRemnants(t, resp)
	})

	t.Run("project brief allows empty arrays when bindings are incomplete", func(t *testing.T) {
		h.resetFixture(t, "completed-unbound")

		repositoryID := h.mustRepositoryIDByName(t, "main-repo")
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
		// 2026-08-18 phase14-10 T7 裁决：画像主表已 drop，brief 不再依赖画像行，
		// 空 bindings 场景同样不得出现画像残余字段。
		assertNoProfileRemnants(t, resp)
	})

	// ========================================================================
	// phase15-06：brief progress 块（字段 9）三场景
	// （独立 repository fixture + 事件经 progress CommandService 写入，
	//   不经 resetFixture，不破坏既有共享 fixture 场景）
	// ========================================================================

	t.Run("project brief progress round-trip derives current phase from events", func(t *testing.T) {
		repositoryID := h.insertProgressFixtureRepository(t)
		now := time.Now().UTC()

		startedTitle := "phase15 项目推进时间轴基座"
		h.mustCreateProgressEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindPhaseStarted,
			"phase15", startedTitle, now.Add(-2*time.Hour))
		h.mustCreateProgressEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindTaskCompleted,
			"phase15-06", "后端主线落地", now.Add(-1*time.Hour))

		resp, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		progressBlock := resp.GetProgress()
		if progressBlock == nil {
			t.Fatal("expected progress block always constructed")
		}
		if progressBlock.GetCurrentPhaseKey() != "phase15" {
			t.Fatalf("expected current_phase_key %q, got %q", "phase15", progressBlock.GetCurrentPhaseKey())
		}
		if progressBlock.GetCurrentPhaseLabel() != startedTitle {
			t.Fatalf("expected current_phase_label %q, got %q", startedTitle, progressBlock.GetCurrentPhaseLabel())
		}

		// recent_events：三键链倒序含全部录入事件（最新 task_completed 在前）。
		recent := progressBlock.GetRecentEvents()
		if len(recent) != 2 {
			t.Fatalf("expected 2 recent events, got %d", len(recent))
		}
		if recent[0].GetTaskKey() != "phase15-06" || recent[0].GetTitle() != "后端主线落地" {
			t.Fatalf("unexpected latest recent event: %+v", recent[0])
		}
		if recent[1].GetTaskKey() != "phase15" || recent[1].GetTitle() != startedTitle {
			t.Fatalf("unexpected second recent event: %+v", recent[1])
		}

		// latest_task_completed：最新的 task_completed 事件。
		latest := progressBlock.GetLatestTaskCompleted()
		if latest == nil {
			t.Fatal("expected latest_task_completed set")
		}
		if latest.GetTaskKey() != "phase15-06" || latest.GetTitle() != "后端主线落地" {
			t.Fatalf("unexpected latest_task_completed: %+v", latest)
		}
	})

	t.Run("project brief progress current phase empty after phase_completed", func(t *testing.T) {
		repositoryID := h.insertProgressFixtureRepository(t)
		now := time.Now().UTC()

		// round-trip 基础状态：phase_started + task_completed（同场景①）。
		h.mustCreateProgressEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindPhaseStarted,
			"phase15", "phase15 项目推进时间轴基座", now.Add(-3*time.Hour))
		h.mustCreateProgressEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindTaskCompleted,
			"phase15-06", "后端主线落地", now.Add(-2*time.Hour))
		// 再录入 phase_completed（同 task_key）→ 当前 phase 派生为空（DoD 冻结断言）。
		h.mustCreateProgressEvent(t, repositoryID, progress.WorkflowTypePhase, progress.EventKindPhaseCompleted,
			"phase15", "phase15 阶段收口", now.Add(-1*time.Hour))

		resp, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		progressBlock := resp.GetProgress()
		if progressBlock == nil {
			t.Fatal("expected progress block always constructed")
		}
		// phase_completed 后当前 phase 为空（三态判定：全部完结 → 空值，与从未开始同型零值）。
		if progressBlock.GetCurrentPhaseKey() != "" {
			t.Fatalf("expected empty current_phase_key after phase_completed, got %q", progressBlock.GetCurrentPhaseKey())
		}
		if progressBlock.GetCurrentPhaseLabel() != "" {
			t.Fatalf("expected empty current_phase_label after phase_completed, got %q", progressBlock.GetCurrentPhaseLabel())
		}

		// recent_events 仍含全部 3 条录入事件（phase_completed 同样入流）。
		if len(progressBlock.GetRecentEvents()) != 3 {
			t.Fatalf("expected 3 recent events, got %d", len(progressBlock.GetRecentEvents()))
		}

		// latest_task_completed 不受 phase_completed 影响（仍为最新 task_completed）。
		latest := progressBlock.GetLatestTaskCompleted()
		if latest == nil || latest.GetTaskKey() != "phase15-06" {
			t.Fatalf("expected latest_task_completed phase15-06, got %+v", latest)
		}
	})

	t.Run("project brief progress block always constructed for repository without events", func(t *testing.T) {
		// 空态恒构造：0 事件仓库的 progress 块非 nil + 空数组 + 零值字段。
		repositoryID := h.insertProgressFixtureRepository(t)

		resp, err := h.client.GetProjectBrief(t.Context(), &pb.GetProjectBriefRequest{
			RepositoryId: repositoryID,
		})
		if err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		progressBlock := resp.GetProgress()
		if progressBlock == nil {
			t.Fatal("expected progress block non-nil for repository without events")
		}
		if progressBlock.GetCurrentPhaseKey() != "" {
			t.Fatalf("expected empty current_phase_key, got %q", progressBlock.GetCurrentPhaseKey())
		}
		if progressBlock.GetCurrentPhaseLabel() != "" {
			t.Fatalf("expected empty current_phase_label, got %q", progressBlock.GetCurrentPhaseLabel())
		}
		if progressBlock.GetLatestTaskCompleted() != nil {
			t.Fatalf("expected latest_task_completed unset, got %+v", progressBlock.GetLatestTaskCompleted())
		}
		// recent_events 空数组：proto wire 语义下空 repeated 不编码，客户端零值
		// len==0 即空数组（服务端组装侧恒构造非 nil 空切片，见 connect/server.go）。
		if recent := progressBlock.GetRecentEvents(); len(recent) != 0 {
			t.Fatalf("expected empty recent_events array, got %d events", len(recent))
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
	// 2026-08-18 phase14-10 T7 裁决：画像残余彻底退役，不再注入画像 reader
	// phase14-07：注入 standard 读取主线作为 GetProjectBrief.standards[] 的 candidate 依赖
	standardReader := standardservice.NewQueryService(standardrepo.NewStandardStore(pool))
	// phase15-06：注入 progress 派生摘要主线作为 GetProjectBrief.progress 的 candidate 依赖
	// （签名演进编译最小适配；progress 集成断言归 phase15-06 Task 7）
	progressReader := progressservice.NewQueryService(
		progressrepo.NewProgressEventStore(pool),
		progresscandidate.NewRepositoryReader(pool),
	)
	querySvc := projectcontextservice.NewQueryService(readers, standardReader, progressReader)

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

// assertNoProfileRemnants 断言 brief 响应序列化后不含任何画像残余字段
// （2026-08-18 phase14-10 T7 用户裁决：governance_profile / current_phase
// 已从 proto 移除并 reserved，global_assets 已于 phase14-09 移除；
// protojson 默认输出 camelCase key，因此按 camelCase 形态逐一排除）。
func assertNoProfileRemnants(t *testing.T, resp *pb.GetProjectBriefResponse) {
	t.Helper()

	raw, err := protojson.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal brief response: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal brief response json: %v", err)
	}

	for _, key := range []string{"governanceProfile", "currentPhase", "globalAssets"} {
		if _, ok := decoded[key]; ok {
			t.Fatalf("expected brief response without profile remnant %q, got keys: %v", key, decodedKeys(decoded))
		}
	}
}

// decodedKeys 返回 map 中全部顶层 key（排序后），供失败信息定位。
func decodedKeys(decoded map[string]any) []string {
	keys := make([]string, 0, len(decoded))
	for k := range decoded {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
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

// insertProgressFixtureRepository 建立独立 repository fixture 行（uuid 后缀，
// phase15-06 brief progress 场景专用；不经 resetFixture，不触碰共享 fixture），
// 返回 repository id。注册 cleanup：先删 progress_events 行再删 repositories 行
// （progress_events.repository_id FK ON DELETE RESTRICT，逆序清理）。
func (h *projectContextIntegrationHarness) insertProgressFixtureRepository(t *testing.T) string {
	t.Helper()

	repositoryID := uuid.NewString()
	repositoryName := "brief-progress-it-" + repositoryID[:8]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := h.pool.Exec(ctx, `
		INSERT INTO repositories (id, name, url, provider, status)
		VALUES ($1, $2, $3, $4, $5)
	`, repositoryID, repositoryName, "https://example.com/"+repositoryName, "github", "active"); err != nil {
		t.Fatalf("insert progress fixture repository: %v", err)
	}

	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		if _, err := h.pool.Exec(cleanupCtx, `DELETE FROM progress_events WHERE repository_id = $1`, repositoryID); err != nil {
			t.Fatalf("cleanup progress events: %v", err)
		}
		if _, err := h.pool.Exec(cleanupCtx, `DELETE FROM repositories WHERE id = $1`, repositoryID); err != nil {
			t.Fatalf("cleanup repository: %v", err)
		}
	})
	return repositoryID
}

// mustCreateProgressEvent 经 progress CommandService 创建推进事件 fixture
// （brief progress 场景的事件插入路径：直连 harness 同一 pool 的 service 写入，
// 与既有 mustCreateAndBindStandard 的 service 写入模式一致）。
func (h *projectContextIntegrationHarness) mustCreateProgressEvent(
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

	commandSvc := progressservice.NewCommandService(
		progresscandidate.NewRepositoryReader(h.pool),
		progressrepo.NewProgressEventStore(h.pool),
	)
	created, err := commandSvc.CreateProgressEvent(ctx, &progress.CreateProgressEventInput{
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
