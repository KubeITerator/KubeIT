package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/common"
	"kubeIT/pkg/grpc/task"
	"kubeIT/server/helpers"
)

type TaskManagementService struct {
	task.UnimplementedTaskManagementServiceServer
	database   *db.Database
	authorizer *helpers.Authorizer
}

func NewTaskManagementService(db *db.Database, authorizer *helpers.Authorizer) *TaskManagementService {
	return &TaskManagementService{database: db, authorizer: authorizer}
}

func (tms *TaskManagementService) CreateTask(context.Context, *task.Task) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (tms *TaskManagementService) GetTask(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (tms *TaskManagementService) DeleteGroup(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
