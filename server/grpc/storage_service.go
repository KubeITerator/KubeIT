package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/common"
	"kubeIT/pkg/grpc/storage"
	"kubeIT/server/helpers"
)

type StorageService struct {
	storage.UnimplementedStorageManagementServiceServer
	database   *db.Database
	authorizer *helpers.Authorizer
}

func NewStorageService(db *db.Database, authorizer *helpers.Authorizer) *StorageService {
	return &StorageService{database: db, authorizer: authorizer}
}

func (sts *StorageService) InitUpload(context.Context, *storage.InitRequest) (*storage.IDmessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitUpload not implemented")
}
func (sts *StorageService) GetUploadUrl(context.Context, *storage.IDmessage) (*storage.UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadUrl not implemented")
}
func (sts *StorageService) FinishUpload(context.Context, *storage.IDmessage) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishUpload not implemented")
}
func (sts *StorageService) DeleteUpload(context.Context, *storage.IDmessage) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUpload not implemented")
}
func (sts *StorageService) GetDownload(context.Context, *storage.IDmessage) (*storage.UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownload not implemented")
}
