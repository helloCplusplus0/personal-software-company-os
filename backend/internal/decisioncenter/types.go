// Package decisioncenter 承载 Decision Center 后端模块的全部业务实现。
//
// 分层语义（对齐 phase03-10 §10.2 / §10.3 已冻结结论）：
//   - handler/    入口层：只负责承接 HTTP 请求与返回结果
//   - service/    业务编排层：动作语义、校验顺序与跨连接口编排
//   - repository/ 数据访问层：decisions / decision_links
//   - candidate/  外部连接层：Module 候选读取（由 Decision Center 拥有）
//
// 本文件定义跨层共享的 API 消息结构。
// 约束（phase03-10 §6.6 合同与存储解耦）：
//   - 消息结构独立于数据库表结构，不直接暴露存储模型
//   - DTO 必须从 phase03-11 已落地的 .proto 合同源派生或显式对齐，
//     不得形成第二套合同源
//
// 文件落点：backend/internal/decisioncenter/types.go
package decisioncenter

import (
	"time"

	"github.com/psco/backend/internal/moduleregistry"
)

// DecisionStatus 冻结为 proposed / active / superseded / archived
// （对齐 .proto DecisionStatus 枚举，phase03-10 §5.6）。
type DecisionStatus string

const (
	DecisionStatusProposed  DecisionStatus = "proposed"
	DecisionStatusActive    DecisionStatus = "active"
	DecisionStatusSuperseded DecisionStatus = "superseded"
	DecisionStatusArchived  DecisionStatus = "archived"
)

// DecisionLinkTargetType 当前阶段唯一允许的正式目标类型为 module
// （对齐 .proto DecisionLinkTargetType 枚举，phase03-10 §7.4）。
type DecisionLinkTargetType string

const (
	DecisionLinkTargetTypeModule DecisionLinkTargetType = "module"
)

// Decision 核心对象（对齐 .proto Decision 消息，phase03-10 §5.4 / §5.5）。
// alternatives 建模为 []string，对齐 .proto repeated string，按输入顺序保留。
type Decision struct {
	ID           string          `json:"id"`
	Title        string          `json:"title"`
	Context      string          `json:"context"`
	Problem      string          `json:"problem"`
	Alternatives []string        `json:"alternatives"`
	Choice       string          `json:"choice"`
	Reason       string          `json:"reason"`
	Impact       string          `json:"impact"`
	Status       DecisionStatus  `json:"status"`
	CreatedAt    time.Time       `json:"created_at"`
}

// DecisionListItem 列表读取模型（对齐 .proto DecisionListItem，phase03-10 §5.8）。
// linked_module_summary 承接 +N 摘要语义（§5.9）。
type DecisionListItem struct {
	ID                   string         `json:"id"`
	Title                string         `json:"title"`
	Status               DecisionStatus `json:"status"`
	CreatedAt            time.Time      `json:"created_at"`
	LinkCount            int            `json:"link_count"`
	LinkedModuleSummary  string         `json:"linked_module_summary"`
}

// LinkedModule 已关联 Module 的最小展示结构
// （对齐 .proto LinkedModule，phase03-10 §5.9 计算口径的生成基础）。
type LinkedModule struct {
	ModuleID   string `json:"module_id"`
	ModuleName string `json:"module_name"`
}

// SourceContext 入口上下文最小来源标识
// （对齐 .proto SourceContext，phase03-10 §5.11）。
// 无来源时 source_module_id 与 source_module_name 均为空字符串，
// 不得内嵌完整 Module 对象。
type SourceContext struct {
	SourceModuleID   string `json:"source_module_id"`
	SourceModuleName string `json:"source_module_name"`
}

// DecisionDetail 详情读取模型（对齐 .proto DecisionDetail，phase03-10 §5.8）。
// 候选读取结果不内嵌于此，必须通过 DecisionModuleCandidateRead 独立承接。
type DecisionDetail struct {
	Decision      Decision       `json:"decision"`
	LinkedModules []LinkedModule `json:"linked_modules"`
	SourceContext SourceContext  `json:"source_context"`
}

// DecisionModuleCandidate 候选读取返回项（对齐 .proto DecisionModuleCandidate，phase03-10 §5.10）。
// status 复用 moduleregistry.ModuleStatus，不重定义本地等价枚举
// （phase03-10 §7.6 跨包依赖策略）。
type DecisionModuleCandidate struct {
	ModuleID   string                    `json:"module_id"`
	ModuleName string                    `json:"module_name"`
	Status     moduleregistry.ModuleStatus `json:"status"`
}

// --- 写组请求体 ---

// CreateDecisionRequest RecordDecision 写入最小字段
// （对齐 .proto CreateDecisionRequest，phase03-10 §5.5 / §5.8 / §5.11）。
//
// SourceModuleID 为可选入口上下文来源 Module 标识（§5.11）：
//   - 从 Module Detail 带上下文进入 Decision Create 时传入
//   - 后端持久化后通过 DecisionDetailRead.source_context 返回
//   - 无来源时为空字符串
type CreateDecisionRequest struct {
	Title          string          `json:"title"`
	Context        string          `json:"context"`
	Problem        string          `json:"problem"`
	Alternatives   []string        `json:"alternatives"`
	Choice         string          `json:"choice"`
	Reason         string          `json:"reason"`
	Impact         string          `json:"impact"`
	Status         DecisionStatus  `json:"status"`
	SourceModuleID string          `json:"source_module_id"`
}

// CreateDecisionResponse 创建响应，只返回新建 Decision 标识
// （对齐 .proto CreateDecisionResponse，phase03-10 §6.4）。
// 不返回完整 Decision 对象，避免形成脱离 DecisionDetailRead 的第二套回流读取路径。
type CreateDecisionResponse struct {
	DecisionID string `json:"decision_id"`
}

// LinkDecisionToTargetRequest LinkDecisionToTarget 写入最小字段
// （对齐 .proto LinkDecisionToTargetRequest，phase03-10 §5.8）。
// decision_id 由 URL 路径参数承接，不放在请求体。
type LinkDecisionToTargetRequest struct {
	TargetType DecisionLinkTargetType `json:"target_type"`
	ModuleID   string                 `json:"module_id"`
}

// --- 列表查询参数 ---

// ListQuery 列表读取筛选参数（对齐 moduleregistry.ListQuery，phase03-10 §9.5）。
// QueryText 模糊匹配 title；StatusFilter 空 / "all" 表示不过滤。
type ListQuery struct {
	QueryText    string
	StatusFilter DecisionStatus
}
