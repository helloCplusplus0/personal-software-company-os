// Package governanceprofile 承载项目级治理画像结构化写读能力。
//
// 分层语义（对齐 phase13-08 已冻结结论）：
//   - connect/     入口层：Connect handler，proto request 解包 → service 调用 → proto response 组装
//   - service/     业务编排层：读写编排、字段分类约束、8 项资产矩阵校验与事务边界
//   - repository/  持久化层：PostgreSQL 读写（主记录 + 两组 bindings 单一事务）
//   - candidate/   外部连接层：跨模块 reader 接口定义与实现（governanceprofile 自拥有）
//
// 本文件定义跨层共享的领域结构与治理冻结常量。
// 约束：消息结构从 proto/psco/governance_profile/v1/governance_profile.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在 types.go 中新增 .proto 中不存在的业务字段语义。
package governanceprofile

import "time"

// ============================================================================
// 受控枚举
// ============================================================================

// TrackType 当前项目正式技术路线（read-only 受控枚举）。
// 对齐 TECH_STACK_BASELINE.md 两条受控路线。
type TrackType string

const (
	TrackTypeProduct       TrackType = "product"
	TrackTypeDurableSystem TrackType = "durable_system"
)

// PhaseStatus 当前阶段状态（read-only 受控枚举）。
type PhaseStatus string

const (
	PhaseStatusPlanned    PhaseStatus = "planned"
	PhaseStatusInProgress PhaseStatus = "in_progress"
	PhaseStatusCompleted  PhaseStatus = "completed"
	PhaseStatusBlocked    PhaseStatus = "blocked"
)

// ============================================================================
// 根级正式上游冻结结论投影（read-only 字段唯一受控来源）
// ============================================================================

// read-only 字段（track_type / current_phase_*）只允许来自根级正式上游的冻结结果回读，
// 不允许由治理画像维护写路径改写（phase13-05 写边界）。
// 本组常量是根级冻结结论在服务端的最小受控投影：
//   - 技术路线来自 TECH_STACK_BASELINE.md 冻结结论
//   - 当前阶段来自 plan.md 冻结结论
//
// 根级上游演进时，通过受控方式同步本投影（手工维护优先）。
const (
	// CurrentProfileVersion 治理画像正式版本（第一版唯一取值，由服务端写入）。
	CurrentProfileVersion = "project_governance_profile_v1"

	// RootFrozenTrackType PSCO 当前冻结技术路线。
	RootFrozenTrackType = TrackTypeDurableSystem

	// RootFrozenCurrentPhaseName PSCO 当前正式阶段名。
	RootFrozenCurrentPhaseName = "phase13_project_governance_profile_foundation"

	// RootFrozenCurrentPhaseRef 当前阶段正式入口引用（plan.md 是阶段路线唯一真相源）。
	RootFrozenCurrentPhaseRef = "plan.md#phase13_project_governance_profile_foundation"

	// RootFrozenCurrentPhaseStatus 当前阶段状态。
	RootFrozenCurrentPhaseStatus = PhaseStatusInProgress
)

// ============================================================================
// 8 项全局规范资产冻结矩阵（phase13-05 逐项承接策略）
// ============================================================================

// globalAssetMatrixEntry 单项资产的承接策略。
type globalAssetMatrixEntry struct {
	// SummaryRequired 该资产是否必须携带 structured_summary。
	// 前 5 份摘要型资产必填；README / global_skills / project_skills 第一版允许为空。
	SummaryRequired bool
}

// globalAssetMatrix 8 项全局规范资产逐项承接矩阵。
// name 集合与 summary 必填性冻结自 phase13-05，不得在实现阶段重新猜测。
var globalAssetMatrix = map[string]globalAssetMatrixEntry{
	"project_rules.md":       {SummaryRequired: true},
	"TECH_STACK_BASELINE.md": {SummaryRequired: true},
	"AGENTS.md":              {SummaryRequired: true},
	"architecture_map.md":    {SummaryRequired: true},
	"plan.md":                {SummaryRequired: true},
	"README.md":              {SummaryRequired: false},
	"global_skills.md":       {SummaryRequired: false},
	"project_skills.md":      {SummaryRequired: false},
}

// IsKnownGlobalAsset 判断资产名是否属于 8 项冻结矩阵。
func IsKnownGlobalAsset(name string) bool {
	_, ok := globalAssetMatrix[name]
	return ok
}

// GlobalAssetSummaryRequired 判断矩阵内资产是否必须携带结构化摘要。
// 调用方需先用 IsKnownGlobalAsset 确认成员资格。
func GlobalAssetSummaryRequired(name string) bool {
	return globalAssetMatrix[name].SummaryRequired
}

// MarkdownResolvableForGlobalAsset 判断资产正文是否允许按 entry_ref 回源。
// 第一版冻结矩阵中的 8 项全局规范资产都允许 markdown 正文回源，但正文不入库。
func MarkdownResolvableForGlobalAsset(name string) bool {
	return IsKnownGlobalAsset(name)
}

// ============================================================================
// 核心领域结构
// ============================================================================

// CanonicalRootFileBinding 根级 canonical 文件绑定。
// 对齐 proto CanonicalRootFileBinding（file_name / role / required）。
type CanonicalRootFileBinding struct {
	FileName string `json:"file_name"`
	Role     string `json:"role"`
	Required bool   `json:"required"`
}

// GlobalAssetBinding 全局规范资产绑定。
// 对齐 proto GlobalAssetBinding（name / kind / entry_ref / role / structured_summary / markdown_resolvable）。
type GlobalAssetBinding struct {
	Name              string  `json:"name"`
	Kind              string  `json:"kind"`
	EntryRef          string  `json:"entry_ref"`
	Role              string  `json:"role"`
	StructuredSummary *string `json:"structured_summary"`
	// MarkdownResolvable 是只读能力状态，由服务端按资产矩阵计算返回，不参与持久化写入。
	MarkdownResolvable bool `json:"markdown_resolvable"`
}

// GovernanceProfileRecord 治理画像主记录（承接 phase13-04 冻结的 9 类核心字段）。
type GovernanceProfileRecord struct {
	RepositoryID          string      `json:"repository_id"`
	ProjectProfileVersion string      `json:"project_profile_version"`
	TrackType             TrackType   `json:"track_type"`
	TemplateSource        *string     `json:"template_source"`
	DocsWorkflowLayout    string      `json:"docs_workflow_layout"`
	CurrentPhaseName      string      `json:"current_phase_name"`
	CurrentPhaseRef       string      `json:"current_phase_ref"`
	CurrentPhaseStatus    PhaseStatus `json:"current_phase_status"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
}

// GovernanceProfileReadResult 治理画像聚合读取结果（主记录 + 两组 bindings）。
// 对齐 proto GetGovernanceProfileResponse。
type GovernanceProfileReadResult struct {
	Record              GovernanceProfileRecord    `json:"record"`
	CanonicalRootFiles  []CanonicalRootFileBinding `json:"canonical_root_files"`
	GlobalAssetBindings []GlobalAssetBinding       `json:"global_asset_bindings"`
}

// UpdateGovernanceProfileInput 治理画像保存输入。
// 可写集合冻结（phase13-05 写边界）：template_source / docs_workflow_layout /
// canonical_root_files[] / global_asset_bindings[]。
// 显式排除 read-only 字段（track_type / current_phase_*）与 project_profile_version。
type UpdateGovernanceProfileInput struct {
	RepositoryID        string                     `json:"repository_id"`
	TemplateSource      *string                    `json:"template_source"`
	DocsWorkflowLayout  string                     `json:"docs_workflow_layout"`
	CanonicalRootFiles  []CanonicalRootFileBinding `json:"canonical_root_files"`
	GlobalAssetBindings []GlobalAssetBinding       `json:"global_asset_bindings"`
}
