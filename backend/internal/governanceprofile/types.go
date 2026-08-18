// Package governanceprofile 承载项目级治理画像主记录轻量读取能力。
//
// phase14-09 收缩（画像系统性退役）：
//   - 写路径（画像保存 RPC 编排与主表/两表事务写入）与两组 bindings
//     （canonical 根级文件 / 全局规范资产）已由 Standard 实体承接
//   - 画像 Connect handler 子包与写入候选服务子包已随写路径删除
//   - 模块收敛为纯读：仅 governance_profiles 主表三组字段
//     （track_type / template_source / current_phase 三字段），
//     服务 brief 内联 BriefGovernanceProfile 装配
//
// 本文件定义跨层共享的领域结构与受控枚举。
// 约束：消息结构从 proto/psco/project_context/v1/project_context.proto 内联
// BriefGovernanceProfile 单向派生或显式对齐，不直接暴露存储模型，
// 不在 types.go 中新增 .proto 中不存在的业务字段语义。
package governanceprofile

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
// 核心领域结构
// ============================================================================

// GovernanceProfileCoreReadResult 治理画像主记录核心字段轻量读取结果。
//
// phase14-09 收缩后模块对外唯一读结果形态，对齐主表三组字段：
//   - track_type（read-only 受控枚举）
//   - template_source（optional，允许为空）
//   - current_phase_name / current_phase_ref / current_phase_status（read-only 三字段）
//
// 对齐 proto/psco/project_context/v1/project_context.proto 内联
// BriefGovernanceProfile（repository_id 为请求锚点，由消费方组装时回填）。
type GovernanceProfileCoreReadResult struct {
	TrackType          TrackType   `json:"track_type"`
	TemplateSource     *string     `json:"template_source"`
	CurrentPhaseName   string      `json:"current_phase_name"`
	CurrentPhaseRef    string      `json:"current_phase_ref"`
	CurrentPhaseStatus PhaseStatus `json:"current_phase_status"`
}
