package grpc

import (
	v1alpha2 "KubeIT-gRPC/model/go"
	"context"
	"google.golang.org/grpc/metadata"
)

type UserManagerServer struct {
	v1alpha2.UnimplementedUserManagerServer
}

func NewUserManagerServer() *UserManagerServer {
	return &UserManagerServer{}
}

func (a *UserManagerServer) AddUserToGroup(ctx context.Context, request *v1alpha2.UserGroupRequest) (*v1alpha2.StatusReport, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	return nil, nil

}
func (a *UserManagerServer) GetUser(ctx context.Context, request *v1alpha2.UserIDRequest) (*v1alpha2.User, error) {
	return nil, nil
}
func (a *UserManagerServer) GetUserPermissions(ctx context.Context, request *v1alpha2.UserIDRequest) (*v1alpha2.UserPermissionResponse, error) {
	return nil, nil
}
func (a *UserManagerServer) ChangeUserPermission(ctx context.Context, request *v1alpha2.ChangePermissionRequest) (*v1alpha2.StatusReport, error) {
	return nil, nil
}
func (a *UserManagerServer) DeleteUser(ctx context.Context, request *v1alpha2.DeleteUserRequest) (*v1alpha2.StatusReport, error) {
	return nil, nil
}
func (a *UserManagerServer) RemoveUserFromGroup(ctx context.Context, request *v1alpha2.UserGroupRequest) (*v1alpha2.StatusReport, error) {
	return nil, nil
}
