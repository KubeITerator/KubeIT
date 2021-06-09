package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/common"
	"kubeIT/pkg/grpc/workflow"
	"kubeIT/server/helpers"
)

type WorkflowManagementService struct {
	workflow.UnimplementedWorkflowManagementServiceServer
	database   *db.Database
	authorizer *helpers.Authorizer
}

func NewWorkflowManagementService(db *db.Database, authorizer *helpers.Authorizer) *WorkflowManagementService {
	return &WorkflowManagementService{database: db, authorizer: authorizer}
}

func (wfms *WorkflowManagementService) CreateWorkflow(context.Context, *workflow.Workflow) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflow not implemented")
}
func (wfms *WorkflowManagementService) GetWorkflow(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflow not implemented")
}
func (wfms *WorkflowManagementService) DeleteWorkflow(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflow not implemented")
}
func (wfms *WorkflowManagementService) UpdateStatus(context.Context, *workflow.StatusUpdate) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (wfms *WorkflowManagementService) GetStatus(context.Context, *common.IDRequest) (*workflow.WorkflowStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
