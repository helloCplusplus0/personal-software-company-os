// Package onboarding 承载 Onboarding 首轮状态读取后端模块的全部业务实现。
//
// 分层语义（对齐 phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"）：
//   - handler/     入口层：只负责承接 HTTP 请求与返回结果
//   - service/     业务编排层：首轮状态推导与 current_step / completion_progress 计算
//   - candidate/   外部连接层：跨模块计数 reader（Onboarding 自拥有）
//
// 当前阶段 Onboarding 只承接 GetFirstRunState 读组，不新增独立写组。
// 四类 draft-first 写入继续复用既有 canonical create 模块（phase06-10 / phase06-12 冻结）。
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/onboarding/v1/onboarding.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在 types.go 或 handler DTO 中新增 .proto 中不存在的业务字段语义。
package onboarding

// ============================================================================
// 枚举类型
// ============================================================================
//
// 所有枚举使用 Go string 类型，常量值使用 snake_case，
// 对齐现有模块模式与 .proto 枚举的 JSON 小写形式。

// FirstRunStatus 首轮 onboarding 运行的状态机。
// 对齐 proto FirstRunStatus：UNSPECIFIED / NOT_STARTED / IN_PROGRESS / COMPLETED。
//
// 状态跃迁口径（phase06-01 / phase06-12 冻结）：
//   - not_started：尚未开始任何首轮对象写入
//   - in_progress：已至少创建 1 条首轮对象记录但四类对象未全部持久化
//   - completed：四类对象均已持久化并满足首轮成功会话条件
type FirstRunStatus string

const (
	FirstRunStatusUnspecified FirstRunStatus = ""
	FirstRunStatusNotStarted  FirstRunStatus = "not_started"
	FirstRunStatusInProgress  FirstRunStatus = "in_progress"
	FirstRunStatusCompleted   FirstRunStatus = "completed"
)

// OnboardingStep Onboarding 首轮录入主线的步骤分段。
// 对齐 proto OnboardingStep。
// 推荐执行顺序冻结为 Product -> Repository -> Module -> Decision。
// COMPLETE 表示首轮成功会话已成立，Onboarding 主线不再作为主入口。
type OnboardingStep string

const (
	OnboardingStepUnspecified OnboardingStep = ""
	OnboardingStepWelcome     OnboardingStep = "welcome"
	OnboardingStepProduct     OnboardingStep = "product"
	OnboardingStepRepository  OnboardingStep = "repository"
	OnboardingStepModule      OnboardingStep = "module"
	OnboardingStepDecision    OnboardingStep = "decision"
	OnboardingStepComplete    OnboardingStep = "complete"
)

// ============================================================================
// 核心消息 DTO
// ============================================================================

// FirstRunState 首轮 onboarding 运行的最小读模型。
// 对齐 proto FirstRunState。
// 该消息只承接读取，不承接任何 draft-first 写入语义。
// current_step 必须显式承接当前引导步骤，而不是要求前端自行推导。
// completion_progress 表达首轮四类对象的完成进度，取值范围冻结为 0 / 25 / 50 / 75 / 100。
type FirstRunState struct {
	Status             FirstRunStatus `json:"status"`
	IsFirstEntry       bool           `json:"is_first_entry"`
	CurrentStep        OnboardingStep `json:"current_step"`
	CompletionProgress int            `json:"completion_progress"`
}

// ============================================================================
// 响应 DTO
// ============================================================================

// FirstRunStateReadResult GetFirstRunState 的响应结构。
// 对齐 proto GetFirstRunStateResponse：单一 first_run_state 字段包装主读模型。
// handler 必须返回此包络结构，不得直接返回裸 FirstRunState，
// 以保证 HTTP JSON 形状与 onboarding.proto 冻结的唯一合同源一致。
type FirstRunStateReadResult struct {
	FirstRunState *FirstRunState `json:"first_run_state"`
}
