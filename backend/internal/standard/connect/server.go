// Package connect — standard Connect transport 实现。
//
// 本文件是 phase14 Standard 全局规范实体后端主线的 Connect handler 实现。
// 职责仅限于：proto request 解包 → service 调用 → proto response 组装 → 错误映射。
// 本层不写 SQL；树与枚举的 proto ↔ domain 转换收敛在本文件。
//
// 文件落点：backend/internal/standard/connect/server.go
package connect

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/psco/backend/internal/connecterrors"
	pbc "github.com/psco/backend/internal/gen/connect/psco/standard/v1/standardv1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/standard/v1"
	"github.com/psco/backend/internal/standard"
	"github.com/psco/backend/internal/standard/service"
)

// Server 实现 StandardServiceHandler 接口。
type Server struct {
	querySvc   *service.QueryService
	commandSvc *service.CommandService
}

var _ pbc.StandardServiceHandler = (*Server)(nil)

// NewServer 构造 Standard Connect handler。
func NewServer(querySvc *service.QueryService, commandSvc *service.CommandService) *Server {
	return &Server{
		querySvc:   querySvc,
		commandSvc: commandSvc,
	}
}

// CreateStandard 承接规范创建。
// status 解包语义：未设置（nil）→ draft；显式 UNSPECIFIED → InvalidArgument。
func (s *Server) CreateStandard(ctx context.Context, req *pb.CreateStandardRequest) (*pb.CreateStandardResponse, error) {
	status := standard.StandardStatusDraft
	if req.Status != nil {
		if *req.Status == pb.StandardStatus_STANDARD_STATUS_UNSPECIFIED {
			return nil, connecterrors.MapToConnectError(fmt.Errorf("%w: status must be specified", standard.ErrInvalidInput))
		}
		mapped, err := protoStatusToDomain(*req.Status)
		if err != nil {
			return nil, connecterrors.MapToConnectError(err)
		}
		status = mapped
	}

	result, err := s.commandSvc.CreateStandard(ctx, standard.CreateStandardInput{
		Name:          req.Name,
		Description:   req.Description,
		DirectoryTree: protoTreeToDomain(req.DirectoryTree),
		Status:        status,
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.CreateStandardResponse{
		Standard: DomainStandardToProto(*result),
	}, nil
}

// ListStandards 承接全量规范列表读取（不分页）。
func (s *Server) ListStandards(ctx context.Context, req *pb.ListStandardsRequest) (*pb.ListStandardsResponse, error) {
	results, err := s.querySvc.ListStandards(ctx)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	standards := make([]*pb.Standard, 0, len(results))
	for _, r := range results {
		standards = append(standards, DomainStandardToProto(r))
	}
	return &pb.ListStandardsResponse{Standards: standards}, nil
}

// GetStandard 承接规范详情读取（主记录 + 绑定集合）。
func (s *Server) GetStandard(ctx context.Context, req *pb.GetStandardRequest) (*pb.GetStandardResponse, error) {
	result, bindings, err := s.querySvc.GetStandard(ctx, req.StandardId)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	protoBindings := make([]*pb.StandardBinding, 0, len(bindings))
	for _, b := range bindings {
		protoBindings = append(protoBindings, domainBindingToProto(b))
	}
	return &pb.GetStandardResponse{
		Standard: DomainStandardToProto(*result),
		Bindings: protoBindings,
	}, nil
}

// UpdateStandard 承接规范整树替换更新。
// optional name / description 直接映射 domain *string；optional status 映射 *StandardStatus。
func (s *Server) UpdateStandard(ctx context.Context, req *pb.UpdateStandardRequest) (*pb.UpdateStandardResponse, error) {
	var status *standard.StandardStatus
	if req.Status != nil {
		mapped, err := protoStatusToDomain(*req.Status)
		if err != nil {
			return nil, connecterrors.MapToConnectError(err)
		}
		status = &mapped
	}

	result, err := s.commandSvc.UpdateStandard(ctx, standard.UpdateStandardInput{
		StandardID:    req.StandardId,
		Name:          req.Name,
		Description:   req.Description,
		Status:        status,
		DirectoryTree: protoTreeToDomain(req.DirectoryTree),
		ChangeSummary: req.ChangeSummary,
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.UpdateStandardResponse{
		Standard: DomainStandardToProto(*result),
	}, nil
}

// DeleteStandard 承接规范删除（active 防误删拦截在 service 层）。
func (s *Server) DeleteStandard(ctx context.Context, req *pb.DeleteStandardRequest) (*pb.DeleteStandardResponse, error) {
	if err := s.commandSvc.DeleteStandard(ctx, req.StandardId); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	return &pb.DeleteStandardResponse{}, nil
}

// BindStandard 承接绑定建立。
// 枚举解包前置校验：UNSPECIFIED 或越界值直接 InvalidArgument，
// 剩余校验链（八格矩阵 / target 存在 / 唯一约束）按冻结顺序在 service 层承接。
func (s *Server) BindStandard(ctx context.Context, req *pb.BindStandardRequest) (*pb.BindStandardResponse, error) {
	targetType, err := protoTargetTypeToDomain(req.TargetType)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	role, err := protoRoleToDomain(req.Role)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	var note *string
	if req.Note != "" {
		note = &req.Note
	}

	result, err := s.commandSvc.BindStandard(ctx, standard.BindStandardInput{
		StandardID: req.StandardId,
		TargetType: targetType,
		TargetID:   req.TargetId,
		Role:       role,
		Note:       note,
	})
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	return &pb.BindStandardResponse{
		Binding: domainBindingToProto(*result),
	}, nil
}

// UnbindStandard 承接绑定解除（四元组定位）。
func (s *Server) UnbindStandard(ctx context.Context, req *pb.UnbindStandardRequest) (*pb.UnbindStandardResponse, error) {
	targetType, err := protoTargetTypeToDomain(req.TargetType)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	role, err := protoRoleToDomain(req.Role)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	if err := s.commandSvc.UnbindStandard(ctx, standard.UnbindStandardInput{
		StandardID: req.StandardId,
		TargetType: targetType,
		TargetID:   req.TargetId,
		Role:       role,
	}); err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}
	return &pb.UnbindStandardResponse{}, nil
}

// ListStandardRevisions 承接 revision 回看（不分页）。
func (s *Server) ListStandardRevisions(ctx context.Context, req *pb.ListStandardRevisionsRequest) (*pb.ListStandardRevisionsResponse, error) {
	revisions, err := s.querySvc.ListStandardRevisions(ctx, req.StandardId)
	if err != nil {
		return nil, connecterrors.MapToConnectError(err)
	}

	protoRevisions := make([]*pb.StandardRevision, 0, len(revisions))
	for _, r := range revisions {
		protoRevisions = append(protoRevisions, domainRevisionToProto(r))
	}
	return &pb.ListStandardRevisionsResponse{Revisions: protoRevisions}, nil
}

// --- 领域结果 → proto 组装 ---
//
// DomainTreeToProto / DomainStandardToProto 为导出函数：phase14-07 起
// projectcontext 的 GetProjectBrief.standards[] 复用同一份 domain → proto
// 映射，避免在消费方重写第二套规范字段映射。

// DomainTreeToProto 将领域目录树节点递归转换为 proto DirectoryTreeNode。
// children 为 nil 时输出空切片（proto repeated 语义）。
func DomainTreeToProto(node *standard.DirectoryTreeNode) *pb.DirectoryTreeNode {
	if node == nil {
		return nil
	}
	children := make([]*pb.DirectoryTreeNode, 0, len(node.Children))
	for _, child := range node.Children {
		children = append(children, DomainTreeToProto(child))
	}
	return &pb.DirectoryTreeNode{
		Name:     node.Name,
		NodeType: domainNodeTypeToProto(node.NodeType),
		Role:     node.Role,
		Summary:  node.Summary,
		Ref:      node.Ref,
		Children: children,
	}
}

// DomainStandardToProto 将领域规范读取结果转换为 proto Standard。
// brief 的 standards[] 顶层块与规范详情同源，共用本函数。
func DomainStandardToProto(s standard.StandardReadResult) *pb.Standard {
	return &pb.Standard{
		Id:            s.ID,
		Name:          s.Name,
		Description:   s.Description,
		Status:        domainStatusToProto(s.Status),
		DirectoryTree: DomainTreeToProto(s.DirectoryTree),
		CreatedAt:     timestamppb.New(s.CreatedAt),
		UpdatedAt:     timestamppb.New(s.UpdatedAt),
	}
}

func domainBindingToProto(b standard.StandardBindingReadResult) *pb.StandardBinding {
	return &pb.StandardBinding{
		Id:         b.ID,
		StandardId: b.StandardID,
		TargetType: domainTargetTypeToProto(b.TargetType),
		TargetId:   b.TargetID,
		Role:       domainRoleToProto(b.Role),
		Note:       b.Note,
		CreatedAt:  timestamppb.New(b.CreatedAt),
	}
}

func domainRevisionToProto(r standard.StandardRevisionReadResult) *pb.StandardRevision {
	return &pb.StandardRevision{
		Id:            r.ID,
		StandardId:    r.StandardID,
		ChangeSummary: r.ChangeSummary,
		CreatedAt:     timestamppb.New(r.CreatedAt),
	}
}

// --- proto 请求 → 领域输入解包 ---

// protoTreeToDomain 将 proto DirectoryTreeNode 递归转换为领域目录树节点。
// node_type 的 UNSPECIFIED / 越界值落空串，由树校验（R1 根类型）承接。
func protoTreeToDomain(node *pb.DirectoryTreeNode) *standard.DirectoryTreeNode {
	if node == nil {
		return nil
	}
	children := make([]*standard.DirectoryTreeNode, 0, len(node.GetChildren()))
	for _, child := range node.GetChildren() {
		children = append(children, protoTreeToDomain(child))
	}
	return &standard.DirectoryTreeNode{
		Name:     node.GetName(),
		NodeType: protoNodeTypeToDomain(node.GetNodeType()),
		Role:     node.GetRole(),
		Summary:  node.GetSummary(),
		Ref:      node.GetRef(),
		Children: children,
	}
}

// --- 枚举转换（string ↔ pb enum；未知值落 UNSPECIFIED） ---

func domainStatusToProto(s standard.StandardStatus) pb.StandardStatus {
	switch s {
	case standard.StandardStatusDraft:
		return pb.StandardStatus_STANDARD_STATUS_DRAFT
	case standard.StandardStatusActive:
		return pb.StandardStatus_STANDARD_STATUS_ACTIVE
	case standard.StandardStatusRetired:
		return pb.StandardStatus_STANDARD_STATUS_RETIRED
	default:
		return pb.StandardStatus_STANDARD_STATUS_UNSPECIFIED
	}
}

func domainNodeTypeToProto(t string) pb.NodeType {
	switch standard.NodeType(t) {
	case standard.NodeTypeDirectory:
		return pb.NodeType_NODE_TYPE_DIRECTORY
	case standard.NodeTypeFile:
		return pb.NodeType_NODE_TYPE_FILE
	default:
		return pb.NodeType_NODE_TYPE_UNSPECIFIED
	}
}

func domainTargetTypeToProto(t standard.BindingTargetType) pb.BindingTargetType {
	switch t {
	case standard.BindingTargetRepository:
		return pb.BindingTargetType_BINDING_TARGET_TYPE_REPOSITORY
	case standard.BindingTargetProduct:
		return pb.BindingTargetType_BINDING_TARGET_TYPE_PRODUCT
	case standard.BindingTargetDecision:
		return pb.BindingTargetType_BINDING_TARGET_TYPE_DECISION
	case standard.BindingTargetModule:
		return pb.BindingTargetType_BINDING_TARGET_TYPE_MODULE
	default:
		return pb.BindingTargetType_BINDING_TARGET_TYPE_UNSPECIFIED
	}
}

func domainRoleToProto(r standard.BindingRole) pb.BindingRole {
	switch r {
	case standard.BindingRoleTemplateSource:
		return pb.BindingRole_BINDING_ROLE_TEMPLATE_SOURCE
	case standard.BindingRoleAdopts:
		return pb.BindingRole_BINDING_ROLE_ADOPTS
	default:
		return pb.BindingRole_BINDING_ROLE_UNSPECIFIED
	}
}

// protoStatusToDomain 将 proto StandardStatus 转换为领域受控枚举。
// UNSPECIFIED / 越界值返回 ErrInvalidInput（InvalidArgument）。
func protoStatusToDomain(v pb.StandardStatus) (standard.StandardStatus, error) {
	switch v {
	case pb.StandardStatus_STANDARD_STATUS_DRAFT:
		return standard.StandardStatusDraft, nil
	case pb.StandardStatus_STANDARD_STATUS_ACTIVE:
		return standard.StandardStatusActive, nil
	case pb.StandardStatus_STANDARD_STATUS_RETIRED:
		return standard.StandardStatusRetired, nil
	default:
		return "", fmt.Errorf("%w: invalid status enum value %d", standard.ErrInvalidInput, int32(v))
	}
}

func protoNodeTypeToDomain(v pb.NodeType) string {
	switch v {
	case pb.NodeType_NODE_TYPE_DIRECTORY:
		return string(standard.NodeTypeDirectory)
	case pb.NodeType_NODE_TYPE_FILE:
		return string(standard.NodeTypeFile)
	default:
		return ""
	}
}

// protoTargetTypeToDomain 将 proto BindingTargetType 转换为领域受控枚举。
// UNSPECIFIED / 越界值返回 ErrInvalidInput（InvalidArgument）。
func protoTargetTypeToDomain(v pb.BindingTargetType) (standard.BindingTargetType, error) {
	switch v {
	case pb.BindingTargetType_BINDING_TARGET_TYPE_REPOSITORY:
		return standard.BindingTargetRepository, nil
	case pb.BindingTargetType_BINDING_TARGET_TYPE_PRODUCT:
		return standard.BindingTargetProduct, nil
	case pb.BindingTargetType_BINDING_TARGET_TYPE_DECISION:
		return standard.BindingTargetDecision, nil
	case pb.BindingTargetType_BINDING_TARGET_TYPE_MODULE:
		return standard.BindingTargetModule, nil
	default:
		return "", fmt.Errorf("%w: invalid target_type enum value %d", standard.ErrInvalidInput, int32(v))
	}
}

// protoRoleToDomain 将 proto BindingRole 转换为领域受控枚举。
// UNSPECIFIED / 越界值返回 ErrInvalidInput（InvalidArgument）。
func protoRoleToDomain(v pb.BindingRole) (standard.BindingRole, error) {
	switch v {
	case pb.BindingRole_BINDING_ROLE_TEMPLATE_SOURCE:
		return standard.BindingRoleTemplateSource, nil
	case pb.BindingRole_BINDING_ROLE_ADOPTS:
		return standard.BindingRoleAdopts, nil
	default:
		return "", fmt.Errorf("%w: invalid role enum value %d", standard.ErrInvalidInput, int32(v))
	}
}
