package grpc

import (
	"context"
	"google.golang.org/grpc/metadata"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/common"
	"kubeIT/pkg/grpc/user"
	"kubeIT/server/helpers"
)

type UserManagerServer struct {
	user.UnimplementedUserManagerServer
	database   *db.Database
	authorizer *helpers.Authorizer
}

func NewUserManagerServer(db *db.Database, authorizer *helpers.Authorizer) *UserManagerServer {
	return &UserManagerServer{database: db, authorizer: authorizer}
}

func (a *UserManagerServer) AddUserToGroup(ctx context.Context, request *user.UserGroupRequest) (*common.StatusReport, error) {
	_, _ = metadata.FromIncomingContext(ctx)

	return nil, nil

}

func (a *UserManagerServer) GetUserInfo(ctx context.Context, request *common.Empty) (*user.User, error) {

	return GetUserPermissions(ctx, a.database, a.authorizer)
}

func (a *UserManagerServer) GetUser(ctx context.Context, request *user.UserIDRequest) (*user.User, error) {
	return nil, nil
}
func (a *UserManagerServer) GetUserPermissions(ctx context.Context, request *user.UserIDRequest) (*user.UserPermissionResponse, error) {
	return nil, nil
}
func (a *UserManagerServer) ChangeUserPermission(ctx context.Context, request *user.ChangePermissionRequest) (*common.StatusReport, error) {
	return nil, nil
}
func (a *UserManagerServer) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*common.StatusReport, error) {
	return nil, nil
}
func (a *UserManagerServer) RemoveUserFromGroup(ctx context.Context, request *user.UserGroupRequest) (*common.StatusReport, error) {
	return nil, nil
}
