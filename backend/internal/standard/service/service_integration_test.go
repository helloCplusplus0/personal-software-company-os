// Package service — standard 集成测试（phase14-07 Task 5）。
//
// 模式沿袭 backend/internal/governanceprofile/repository/profile_store_integration_test.go：
//   - 真实 PostgreSQL（进程环境变量 DATABASE_URL 优先，回落 backend/.env）
//   - 每个测试独立 fixture（uuid 后缀 name / 独立 repositories 行）+ t.Cleanup 显式清理，
//     测试之间互不依赖执行顺序
//   - repositories 表 fixture 直接 INSERT（对齐既有集成测试准备方式）
//
// 与既有模式的唯一偏差：DATABASE_URL 无法获得时 t.Skipf 跳过（SKIP 非 FAIL），
// 系 phase14-07 spec 对本测试的冻结要求；配置了 DATABASE_URL 但连接失败仍 Fatal
// （环境故障应显式暴露，与既有模式一致）。
//
// 断言范围（phase14-07 spec 冻结）：
//   1. UpdateStandard 事务 round-trip（整树替换 + revision 追加原子性）
//   2. QueryService.ListStandardsByRepository 反查（brief 反查主线）
//   3. 写路径错误语义抽查（哨兵错误 + 关键错误信息）
//
// 文件落点：backend/internal/standard/service/service_integration_test.go
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

	"github.com/psco/backend/internal/standard"
	"github.com/psco/backend/internal/standard/candidate"
	"github.com/psco/backend/internal/standard/repository"
)

// ============================================================================
// harness 与 fixture helper
// ============================================================================

// standardIntegrationHarness 聚合被测 service 与底层连接池。
type standardIntegrationHarness struct {
	pool  *pgxpool.Pool
	cmd   *CommandService
	query *QueryService
}

// newStandardIntegrationHarness 建立真实 DB 连接并装配被测 service。
func newStandardIntegrationHarness(t *testing.T) *standardIntegrationHarness {
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

	store := repository.NewStandardStore(pool)
	return &standardIntegrationHarness{
		pool:  pool,
		cmd:   NewCommandService(candidate.NewTargetReader(pool), store),
		query: NewQueryService(store),
	}
}

// createTestStandard 建立规范 fixture 并注册 CASCADE 清理
// （standard_revisions / standard_bindings 经 ON DELETE CASCADE 连带删除）。
func createTestStandard(
	t *testing.T,
	h *standardIntegrationHarness,
	name string,
	tree *standard.DirectoryTreeNode,
	status standard.StandardStatus,
) *standard.StandardReadResult {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := h.cmd.CreateStandard(ctx, standard.CreateStandardInput{
		Name:          name,
		Description:   "phase14-07 standard integration test fixture",
		DirectoryTree: tree,
		Status:        status,
	})
	if err != nil {
		t.Fatalf("create test standard %q: %v", name, err)
	}

	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		if _, err := h.pool.Exec(cleanupCtx, `DELETE FROM standards WHERE id = $1`, result.ID); err != nil {
			t.Fatalf("cleanup standard %s: %v", result.ID, err)
		}
	})
	return result
}

// insertTestRepository 建立 repositories 表 fixture 行（对齐既有集成测试准备方式）。
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

// cleanupTestRepository 删除 repositories fixture 行。
// standard_bindings.target_id 无 DB 外键（多态），standards 已先行 CASCADE 清理。
func cleanupTestRepository(t *testing.T, pool *pgxpool.Pool, repositoryID string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := pool.Exec(ctx, `DELETE FROM repositories WHERE id = $1`, repositoryID); err != nil {
		t.Fatalf("cleanup repository: %v", err)
	}
}

// fileNode 构造 file 叶节点（R7：无 children、role 非空）。
func fileNode(name, role, summary, ref string) *standard.DirectoryTreeNode {
	return &standard.DirectoryTreeNode{
		Name:     name,
		NodeType: string(standard.NodeTypeFile),
		Role:     role,
		Summary:  summary,
		Ref:      ref,
		Children: []*standard.DirectoryTreeNode{},
	}
}

// dirNode 构造 directory 节点。
func dirNode(name string, children ...*standard.DirectoryTreeNode) *standard.DirectoryTreeNode {
	return &standard.DirectoryTreeNode{
		Name:     name,
		NodeType: string(standard.NodeTypeDirectory),
		Children: children,
	}
}

// assertTreeEqual 递归逐项核对目录树（name / node_type / role / summary / ref / children 顺序）。
func assertTreeEqual(t *testing.T, got, want *standard.DirectoryTreeNode) {
	t.Helper()

	if got == nil {
		t.Fatal("expected non-nil directory tree")
	}
	if got.Name != want.Name {
		t.Fatalf("node name mismatch: got %q, want %q", got.Name, want.Name)
	}
	if got.NodeType != want.NodeType {
		t.Fatalf("node %q node_type mismatch: got %q, want %q", want.Name, got.NodeType, want.NodeType)
	}
	if got.Role != want.Role {
		t.Fatalf("node %q role mismatch: got %q, want %q", want.Name, got.Role, want.Role)
	}
	if got.Summary != want.Summary {
		t.Fatalf("node %q summary mismatch: got %q, want %q", want.Name, got.Summary, want.Summary)
	}
	if got.Ref != want.Ref {
		t.Fatalf("node %q ref mismatch: got %q, want %q", want.Name, got.Ref, want.Ref)
	}
	if len(got.Children) != len(want.Children) {
		t.Fatalf("node %q children count mismatch: got %d, want %d", want.Name, len(got.Children), len(want.Children))
	}
	for i := range want.Children {
		assertTreeEqual(t, got.Children[i], want.Children[i])
	}
}

// assertInvalidInput 断言错误为 ErrInvalidInput 哨兵且信息包含关键字。
func assertInvalidInput(t *testing.T, err error, msgContains string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q, got nil", msgContains)
	}
	if !errors.Is(err, standard.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
	if !strings.Contains(err.Error(), msgContains) {
		t.Fatalf("expected error message containing %q, got %q", msgContains, err.Error())
	}
}

// ============================================================================
// 断言组 1：UpdateStandard 事务 round-trip
// ============================================================================

// TestUpdateStandardTransactionRoundTrip 验证整树替换 + revision 追加的事务 round-trip：
// 创建后 revision 为空（CreateStandard 不记 revision）；更新后回读返回新树
// （旧树节点消失、新树节点齐全），updated_at 推进，revision 恰多一条且
// change_summary 精确匹配。
func TestUpdateStandardTransactionRoundTrip(t *testing.T) {
	h := newStandardIntegrationHarness(t)

	initialTree := dirNode(".", dirNode("docs",
		fileNode("README.md", "docs", "入口文档", "/docs/README.md"),
	))
	created := createTestStandard(t, h, "standard-it-update-"+uuid.NewString()[:8], initialTree, standard.StandardStatusActive)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 创建后基线：revision 为空，记录更新前读数
	beforeRevisions, err := h.query.ListStandardRevisions(ctx, created.ID)
	if err != nil {
		t.Fatalf("list revisions before update: %v", err)
	}
	if len(beforeRevisions) != 0 {
		t.Fatalf("expected no revision after create, got %d", len(beforeRevisions))
	}
	before, _, err := h.query.GetStandard(ctx, created.ID)
	if err != nil {
		t.Fatalf("get standard before update: %v", err)
	}

	// 保证 updated_at 推进可观测（timestamptz 微秒精度）
	time.Sleep(20 * time.Millisecond)

	newTree := dirNode(".",
		fileNode("ARCHITECTURE.md", "architecture", "结构总览", "/ARCHITECTURE.md"),
		dirNode("guides",
			fileNode("onboarding.md", "guide", "上手指引", "/guides/onboarding.md"),
		),
	)
	const changeSummary = "phase14-07 集成测试：整树替换为 v2 结构"
	updated, err := h.cmd.UpdateStandard(ctx, standard.UpdateStandardInput{
		StandardID:    created.ID,
		DirectoryTree: newTree,
		ChangeSummary: changeSummary,
	})
	if err != nil {
		t.Fatalf("update standard: %v", err)
	}
	if updated.ID != created.ID {
		t.Fatalf("updated standard id mismatch: got %q, want %q", updated.ID, created.ID)
	}

	// 回读：新树齐全、旧树消失（整树逐项核对天然覆盖两者）
	readBack, _, err := h.query.GetStandard(ctx, created.ID)
	if err != nil {
		t.Fatalf("get standard after update: %v", err)
	}
	assertTreeEqual(t, readBack.DirectoryTree, newTree)

	// 时间语义：updated_at 晚于 created_at，且晚于更新前读数
	if !readBack.UpdatedAt.After(readBack.CreatedAt) {
		t.Fatalf("expected updated_at %v after created_at %v", readBack.UpdatedAt, readBack.CreatedAt)
	}
	if !readBack.UpdatedAt.After(before.UpdatedAt) {
		t.Fatalf("expected updated_at %v after pre-update %v", readBack.UpdatedAt, before.UpdatedAt)
	}

	// revision：数量从 0 → 1 且 change_summary 精确匹配
	revisions, err := h.query.ListStandardRevisions(ctx, created.ID)
	if err != nil {
		t.Fatalf("list revisions after update: %v", err)
	}
	if len(revisions) != 1 {
		t.Fatalf("expected exactly 1 revision after update, got %d", len(revisions))
	}
	if revisions[0].ChangeSummary != changeSummary {
		t.Fatalf("revision change_summary mismatch: got %q, want %q", revisions[0].ChangeSummary, changeSummary)
	}
	if revisions[0].StandardID != created.ID {
		t.Fatalf("revision standard_id mismatch: got %q, want %q", revisions[0].StandardID, created.ID)
	}
}

// ============================================================================
// 断言组 2：StandardReader 反查（QueryService.ListStandardsByRepository）
// ============================================================================

// TestListStandardsByRepositoryReverseLookup 验证 brief 反查主线：
// 同一 Standard 对同一 repository 绑定两 role（adopts + template_source）后，
// 反查返回该 Standard（含 directory_tree 全树，节点字段逐项核对）；
// 无绑定的另一 repository 返回空列表非错误非 nil。
func TestListStandardsByRepositoryReverseLookup(t *testing.T) {
	h := newStandardIntegrationHarness(t)

	boundRepoID := uuid.NewString()
	unboundRepoID := uuid.NewString()
	insertTestRepository(t, h.pool, boundRepoID, "standard-it-bound-"+boundRepoID[:8])
	t.Cleanup(func() { cleanupTestRepository(t, h.pool, boundRepoID) })
	insertTestRepository(t, h.pool, unboundRepoID, "standard-it-unbound-"+unboundRepoID[:8])
	t.Cleanup(func() { cleanupTestRepository(t, h.pool, unboundRepoID) })

	tree := dirNode(".", dirNode("docs",
		fileNode("README.md", "docs", "入口文档", "/docs/README.md"),
	))
	created := createTestStandard(t, h, "standard-it-lookup-"+uuid.NewString()[:8], tree, standard.StandardStatusActive)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 绑定两 role：adopts + template_source（target_type=repository）
	if _, err := h.cmd.BindStandard(ctx, standard.BindStandardInput{
		StandardID: created.ID,
		TargetType: standard.BindingTargetRepository,
		TargetID:   boundRepoID,
		Role:       standard.BindingRoleAdopts,
	}); err != nil {
		t.Fatalf("bind adopts: %v", err)
	}
	templateNote := "phase14-07 集成测试 template_source 绑定"
	if _, err := h.cmd.BindStandard(ctx, standard.BindStandardInput{
		StandardID: created.ID,
		TargetType: standard.BindingTargetRepository,
		TargetID:   boundRepoID,
		Role:       standard.BindingRoleTemplateSource,
		Note:       &templateNote,
	}); err != nil {
		t.Fatalf("bind template_source: %v", err)
	}

	// 两 role 真实生效（绑定集合各一条）
	_, bindings, err := h.query.GetStandard(ctx, created.ID)
	if err != nil {
		t.Fatalf("get standard bindings: %v", err)
	}
	if len(bindings) != 2 {
		t.Fatalf("expected 2 bindings after dual-role bind, got %d", len(bindings))
	}

	// 反查：同一 (standard, repository) 两 role 去重后恰返回一条，含整树
	results, err := h.query.ListStandardsByRepository(ctx, boundRepoID)
	if err != nil {
		t.Fatalf("list standards by repository: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected exactly 1 standard bound to repository, got %d", len(results))
	}
	if results[0].ID != created.ID {
		t.Fatalf("standard id mismatch: got %q, want %q", results[0].ID, created.ID)
	}
	assertTreeEqual(t, results[0].DirectoryTree, tree)

	// 无绑定 repository：空列表非错误非 nil
	unbound, err := h.query.ListStandardsByRepository(ctx, unboundRepoID)
	if err != nil {
		t.Fatalf("list standards by unbound repository: %v", err)
	}
	if unbound == nil {
		t.Fatal("expected non-nil empty slice for unbound repository, got nil")
	}
	if len(unbound) != 0 {
		t.Fatalf("expected empty result for unbound repository, got %d", len(unbound))
	}
}

// ============================================================================
// 断言组 3：写路径错误语义抽查
// ============================================================================

// TestCreateStandardErrorSemantics 抽查创建写路径错误语义：
// name 重复 / status=retired 拒绝创建 / active + 单根空树（R2）。
func TestCreateStandardErrorSemantics(t *testing.T) {
	h := newStandardIntegrationHarness(t)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	suffix := uuid.NewString()[:8]
	validTree := func() *standard.DirectoryTreeNode {
		return dirNode(".", fileNode("README.md", "docs", "", "/README.md"))
	}
	existing := createTestStandard(t, h, "standard-it-dup-"+suffix, validTree(), standard.StandardStatusDraft)

	// name 重复 → ErrInvalidInput 且信息含 "already exists"（UNIQUE 23505 映射）
	_, err := h.cmd.CreateStandard(ctx, standard.CreateStandardInput{
		Name:          existing.Name,
		DirectoryTree: validTree(),
		Status:        standard.StandardStatusDraft,
	})
	assertInvalidInput(t, err, "already exists")

	// status=retired → ErrInvalidInput（"cannot create ... retired"）
	_, err = h.cmd.CreateStandard(ctx, standard.CreateStandardInput{
		Name:          "standard-it-retired-" + suffix,
		DirectoryTree: validTree(),
		Status:        standard.StandardStatusRetired,
	})
	assertInvalidInput(t, err, "cannot create a standard in retired status")

	// status=active + 单根空树（根 "." 无 file 子节点）→ ErrInvalidInput 含 "EMPTY_TREE_NOT_ALLOWED"
	_, err = h.cmd.CreateStandard(ctx, standard.CreateStandardInput{
		Name:          "standard-it-empty-tree-" + suffix,
		DirectoryTree: dirNode("."),
		Status:        standard.StandardStatusActive,
	})
	assertInvalidInput(t, err, "EMPTY_TREE_NOT_ALLOWED")
}

// TestDeleteStandardActiveGuard 验证 active 规范防误删拦截。
func TestDeleteStandardActiveGuard(t *testing.T) {
	h := newStandardIntegrationHarness(t)

	created := createTestStandard(t, h, "standard-it-delete-"+uuid.NewString()[:8],
		dirNode(".", fileNode("README.md", "docs", "", "/README.md")),
		standard.StandardStatusActive,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.cmd.DeleteStandard(ctx, created.ID)
	assertInvalidInput(t, err, "cannot delete an active standard")
}

// TestBindingErrorSemantics 抽查绑定写路径错误语义：
// 八格矩阵非法 / target 不存在 / 重复四元组 / Unbind 不存在 / GetStandard 不存在。
func TestBindingErrorSemantics(t *testing.T) {
	h := newStandardIntegrationHarness(t)

	repoID := uuid.NewString()
	insertTestRepository(t, h.pool, repoID, "standard-it-binding-"+repoID[:8])
	t.Cleanup(func() { cleanupTestRepository(t, h.pool, repoID) })

	created := createTestStandard(t, h, "standard-it-binding-"+uuid.NewString()[:8],
		dirNode(".", fileNode("README.md", "docs", "", "/README.md")),
		standard.StandardStatusDraft,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 八格矩阵非法：template_source 仅允许 repository target
	_, err := h.cmd.BindStandard(ctx, standard.BindStandardInput{
		StandardID: created.ID,
		TargetType: standard.BindingTargetProduct,
		TargetID:   uuid.NewString(),
		Role:       standard.BindingRoleTemplateSource,
	})
	assertInvalidInput(t, err, "only allowed for repository targets")

	// target 不存在（合法 UUID 但表中无行）
	_, err = h.cmd.BindStandard(ctx, standard.BindStandardInput{
		StandardID: created.ID,
		TargetType: standard.BindingTargetRepository,
		TargetID:   uuid.NewString(),
		Role:       standard.BindingRoleAdopts,
	})
	assertInvalidInput(t, err, "not found")

	// 重复四元组 → ErrInvalidInput 且信息含 "already bound"（UNIQUE 23505 映射）
	bindInput := standard.BindStandardInput{
		StandardID: created.ID,
		TargetType: standard.BindingTargetRepository,
		TargetID:   repoID,
		Role:       standard.BindingRoleAdopts,
	}
	if _, err := h.cmd.BindStandard(ctx, bindInput); err != nil {
		t.Fatalf("first bind: %v", err)
	}
	_, err = h.cmd.BindStandard(ctx, bindInput)
	assertInvalidInput(t, err, "already bound")

	// UnbindStandard 不存在的四元组 → ErrBindingNotFound
	err = h.cmd.UnbindStandard(ctx, standard.UnbindStandardInput{
		StandardID: created.ID,
		TargetType: standard.BindingTargetRepository,
		TargetID:   uuid.NewString(),
		Role:       standard.BindingRoleAdopts,
	})
	if err == nil || !errors.Is(err, standard.ErrBindingNotFound) {
		t.Fatalf("expected ErrBindingNotFound, got %v", err)
	}

	// GetStandardByID 不存在 id → ErrStandardNotFound
	_, _, err = h.query.GetStandard(ctx, uuid.NewString())
	if err == nil || !errors.Is(err, standard.ErrStandardNotFound) {
		t.Fatalf("expected ErrStandardNotFound, got %v", err)
	}
}

// ============================================================================
// DB 连接 helper（对齐 governanceprofile 既有集成测试模式）
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
// 无法获得时跳过（SKIP 非 FAIL，phase14-07 spec 冻结要求）。
func mustDatabaseURL(t *testing.T, backendRoot string) string {
	t.Helper()

	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		return databaseURL
	}

	envPath := filepath.Join(backendRoot, ".env")
	if err := godotenv.Load(envPath); err != nil {
		t.Skipf("skip standard integration test: load %s: %v", envPath, err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skipf("skip standard integration test: DATABASE_URL is empty after loading %s", envPath)
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
