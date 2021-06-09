package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/common"
	"kubeIT/pkg/grpc/schema"
	"kubeIT/server/helpers"
)

type SchemaManagementService struct {
	schema.UnimplementedSchemaManagementServiceServer
	database   *db.Database
	authorizer *helpers.Authorizer
}

func NewSchemaManagementService(db *db.Database, authorizer *helpers.Authorizer) *SchemaManagementService {
	return &SchemaManagementService{database: db, authorizer: authorizer}
}

func (sms *SchemaManagementService) CreateSchema(context.Context, *schema.Schema) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchema not implemented")
}
func (sms *SchemaManagementService) GetSchema(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (sms *SchemaManagementService) DeleteSchema(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchema not implemented")
}
