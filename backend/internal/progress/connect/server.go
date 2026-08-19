// Package connect — progress Connect transport 实现。
//
// 本文件是 phase15 项目推进时间轴后端主线的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
// 本层不写 SQL、不做业务校验——全部校验（含 envelope 前置与 V1-V9 枚举判定）
// 统一经根包 ValidateCreateProgressEventInput 按冻结执行序执行，保证"报第一个
// 错误"的顺序不被本层提前拦截破坏。
//
// 解包归一语义（phase15-04 合同设计冻结）：
//   - workflow_type / event_kind：UNSPECIFIED / 越界值落空串 domain 值，
//     由 V1a / V1b 校验承接（沿 standard "未知值落空串由下游校验承接" 模式）
//   - source（optional）：未设置 → 归一 manual（与 DDL DEFAULT 对齐）；
//     显式设置非 MANUAL → V9d [INVALID_SOURCE] 承接
//   - occurred_at：nil → 零值 time.Time，envelope 前置 [INVALID_OCCURRED_AT] 承接
//   - List 过滤参数 workflow_type：UNSPECIFIED / 越界 → nil 不过滤（三轨全量）
//
// 文件落点：backend/internal/progress/connect/server.go
package connect

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/psco/backend/internal/connecterrors"
	pbc "github.com/psco/backend/internal/gen/connect/psco/progress/v1/progressv1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/progress/v1"
	"github.com/psco/backend/internal/progress"
	"github.com/psco/backend/internal/progress/service"
)

// Server 实现 ProgressServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.ProgressServiceHandler = (*Server)(nil)

// NewServer 构造 Progress Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{
		querySvc:   querySvc,
		commandSvc: commandSvc,
	}
}

// ListProgressEvents 承接完整事件流读取（三键链倒序，不分页）。
// workflow_type 过滤解包：合法三值 → 单轨过滤；UNSPECIFIED / 越界 → nil
// 不过滤（三轨全量；过滤参数无错误分支，phase15-04 RPC 1 冻结）。
func (s *Server) ListProgressEvents(ctx context.Context, req *pb.ListProgressEventsRequest) (*pb.ListProgressEventsResponse, error) {
	var workflowFilter *progress.WorkflowType
	switch wt := protoWorkflowTypeToDomain(req.WorkflowType); wt {
	case progress.WorkflowTypePhase, progress.WorkflowTypeAudit, progress.WorkflowTypeFix:
		workflowFilter = &wt
	}

	events, err := s.querySvc.ListProgressEvents(ctx, req.RepositoryId, workflowFilter)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	protoEvents := make([]*pb.ProgressEvent, 0, len(events))
	for _, e := range events {
		protoEvents = append(protoEvents, DomainProgressEventToProto(e))
	}
	return &pb.ListProgressEventsResponse{Events: protoEvents}, nil
}

// CreateProgressEvent 承接推进事件创建。
// source / occurred_at 的 optional 与 nil 语义在解包层归一（见包注释），
// 其余校验统一由根包 validate 按执行序承接。
func (s *Server) CreateProgressEvent(ctx context.Context, req *pb.CreateProgressEventRequest) (*pb.CreateProgressEventResponse, error) {
	input := &progress.CreateProgressEventInput{
		RepositoryID: req.RepositoryId,
		WorkflowType: protoWorkflowTypeToDomain(req.WorkflowType),
		EventKind:    protoEventKindToDomain(req.EventKind),
		TaskKey:      req.TaskKey,
		Title:        req.Title,
		Detail:       req.Detail,
		EvidenceRef:  req.EvidenceRef,
		// optional source 未设置 → 归一 manual；显式设置经映射后由 V9d 承接
		Source:     progress.ProgressSourceManual,
		OccurredAt: time.Time{}, // nil → 零值，envelope 前置 [INVALID_OCCURRED_AT] 承接
	}
	if req.Source != nil {
		input.Source = protoProgressSourceToDomain(*req.Source)
	}
	if req.OccurredAt != nil {
		input.OccurredAt = req.OccurredAt.AsTime()
	}

	result, err := s.commandSvc.CreateProgressEvent(ctx, input)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateProgressEventResponse{
		Event: DomainProgressEventToProto(*result),
	}, nil
}

// DeleteProgressEvent 承接推进事件整条删除（误录修正；无软删除、无 Update）。
func (s *Server) DeleteProgressEvent(ctx context.Context, req *pb.DeleteProgressEventRequest) (*pb.DeleteProgressEventResponse, error) {
	if err := s.commandSvc.DeleteProgressEvent(ctx, req.Id); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	return &pb.DeleteProgressEventResponse{}, nil
}

// --- 领域结果 → proto 组装 ---
//
// DomainProgressEventToProto 为导出函数：phase15-06 起 projectcontext 的
// brief progress 块（latest_task_completed / recent_events 元素）复用同一份
// domain → proto 映射，避免在消费方重写第二套事件字段映射
// （沿 standard/connect.DomainStandardToProto 导出模式）。

// DomainProgressEventToProto 将领域推进事件读取结果转换为 proto ProgressEvent。
// List / Create 响应与 brief progress 块同源，共用本函数。
func DomainProgressEventToProto(e progress.ProgressEventReadResult) *pb.ProgressEvent {
	return &pb.ProgressEvent{
		Id:           e.ID,
		RepositoryId: e.RepositoryID,
		WorkflowType: domainWorkflowTypeToProto(e.WorkflowType),
		EventKind:    domainEventKindToProto(e.EventKind),
		TaskKey:      e.TaskKey,
		Title:        e.Title,
		Detail:       e.Detail,
		EvidenceRef:  e.EvidenceRef,
		Source:       domainProgressSourceToProto(e.Source),
		OccurredAt:   timestamppb.New(e.OccurredAt),
		CreatedAt:    timestamppb.New(e.CreatedAt),
	}
}

// --- 枚举转换（domain string ↔ proto enum） ---

// domain → proto：未知 domain 值落 UNSPECIFIED（沿 standard domainStatusToProto 模式）。

func domainWorkflowTypeToProto(t progress.WorkflowType) pb.WorkflowType {
	switch t {
	case progress.WorkflowTypePhase:
		return pb.WorkflowType_WORKFLOW_TYPE_PHASE
	case progress.WorkflowTypeAudit:
		return pb.WorkflowType_WORKFLOW_TYPE_AUDIT
	case progress.WorkflowTypeFix:
		return pb.WorkflowType_WORKFLOW_TYPE_FIX
	default:
		return pb.WorkflowType_WORKFLOW_TYPE_UNSPECIFIED
	}
}

func domainEventKindToProto(k progress.EventKind) pb.EventKind {
	switch k {
	case progress.EventKindPhaseStarted:
		return pb.EventKind_EVENT_KIND_PHASE_STARTED
	case progress.EventKindPhaseCompleted:
		return pb.EventKind_EVENT_KIND_PHASE_COMPLETED
	case progress.EventKindTaskCompleted:
		return pb.EventKind_EVENT_KIND_TASK_COMPLETED
	case progress.EventKindNote:
		return pb.EventKind_EVENT_KIND_NOTE
	default:
		return pb.EventKind_EVENT_KIND_UNSPECIFIED
	}
}

func domainProgressSourceToProto(s progress.ProgressSource) pb.ProgressSource {
	switch s {
	case progress.ProgressSourceManual:
		return pb.ProgressSource_PROGRESS_SOURCE_MANUAL
	case progress.ProgressSourceGit:
		return pb.ProgressSource_PROGRESS_SOURCE_GIT
	case progress.ProgressSourceAgent:
		return pb.ProgressSource_PROGRESS_SOURCE_AGENT
	default:
		return pb.ProgressSource_PROGRESS_SOURCE_UNSPECIFIED
	}
}

// proto → domain：UNSPECIFIED / 越界值落空串，由根包校验承接
// （Create 路径 V1a / V1b / V9d；List 路径过滤参数归 nil 不过滤）。

func protoWorkflowTypeToDomain(v pb.WorkflowType) progress.WorkflowType {
	switch v {
	case pb.WorkflowType_WORKFLOW_TYPE_PHASE:
		return progress.WorkflowTypePhase
	case pb.WorkflowType_WORKFLOW_TYPE_AUDIT:
		return progress.WorkflowTypeAudit
	case pb.WorkflowType_WORKFLOW_TYPE_FIX:
		return progress.WorkflowTypeFix
	default:
		return "" // UNSPECIFIED / 越界 → 空串，V1a 校验承接
	}
}

func protoEventKindToDomain(v pb.EventKind) progress.EventKind {
	switch v {
	case pb.EventKind_EVENT_KIND_PHASE_STARTED:
		return progress.EventKindPhaseStarted
	case pb.EventKind_EVENT_KIND_PHASE_COMPLETED:
		return progress.EventKindPhaseCompleted
	case pb.EventKind_EVENT_KIND_TASK_COMPLETED:
		return progress.EventKindTaskCompleted
	case pb.EventKind_EVENT_KIND_NOTE:
		return progress.EventKindNote
	default:
		return "" // UNSPECIFIED / 越界 → 空串，V1b 校验承接
	}
}

func protoProgressSourceToDomain(v pb.ProgressSource) progress.ProgressSource {
	switch v {
	case pb.ProgressSource_PROGRESS_SOURCE_MANUAL:
		return progress.ProgressSourceManual
	case pb.ProgressSource_PROGRESS_SOURCE_GIT:
		return progress.ProgressSourceGit
	case pb.ProgressSource_PROGRESS_SOURCE_AGENT:
		return progress.ProgressSourceAgent
	default:
		return "" // UNSPECIFIED / 越界 → 空串，V9d 校验承接
	}
}
