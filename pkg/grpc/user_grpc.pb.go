// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	"KubeIT-gRPC/model/go"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientTokenServiceClient is the client API for ClientTokenService API.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientTokenServiceClient interface {
	CreateToken(ctx context.Context, in *_go.TokenRequest, opts ...grpc.CallOption) (*_go.Token, error)
	GetToken(ctx context.Context, in *_go.UserTokenIDRequest, opts ...grpc.CallOption) (*_go.TokenStatus, error)
	GetTokenStatusList(ctx context.Context, in *_go.UserIDRequest, opts ...grpc.CallOption) (*_go.TokenStatusList, error)
	DeleteToken(ctx context.Context, in *_go.UserTokenIDRequest, opts ...grpc.CallOption) (*_go.StatusReport, error)
}

type clientTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientTokenServiceClient(cc grpc.ClientConnInterface) ClientTokenServiceClient {
	return &clientTokenServiceClient{cc}
}

func (c *clientTokenServiceClient) CreateToken(ctx context.Context, in *_go.TokenRequest, opts ...grpc.CallOption) (*_go.Token, error) {
	out := new(_go.Token)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.ClientTokenService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTokenServiceClient) GetToken(ctx context.Context, in *_go.UserTokenIDRequest, opts ...grpc.CallOption) (*_go.TokenStatus, error) {
	out := new(_go.TokenStatus)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.ClientTokenService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTokenServiceClient) GetTokenStatusList(ctx context.Context, in *_go.UserIDRequest, opts ...grpc.CallOption) (*_go.TokenStatusList, error) {
	out := new(_go.TokenStatusList)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.ClientTokenService/GetTokenStatusList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTokenServiceClient) DeleteToken(ctx context.Context, in *_go.UserTokenIDRequest, opts ...grpc.CallOption) (*_go.StatusReport, error) {
	out := new(_go.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.ClientTokenService/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientTokenServiceServer is the server API for ClientTokenService API.
// All implementations must embed UnimplementedClientTokenServiceServer
// for forward compatibility
type ClientTokenServiceServer interface {
	CreateToken(context.Context, *_go.TokenRequest) (*_go.Token, error)
	GetToken(context.Context, *_go.UserTokenIDRequest) (*_go.TokenStatus, error)
	GetTokenStatusList(context.Context, *_go.UserIDRequest) (*_go.TokenStatusList, error)
	DeleteToken(context.Context, *_go.UserTokenIDRequest) (*_go.StatusReport, error)
	mustEmbedUnimplementedClientTokenServiceServer()
}

// UnimplementedClientTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientTokenServiceServer struct {
}

func (UnimplementedClientTokenServiceServer) CreateToken(context.Context, *_go.TokenRequest) (*_go.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedClientTokenServiceServer) GetToken(context.Context, *_go.UserTokenIDRequest) (*_go.TokenStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedClientTokenServiceServer) GetTokenStatusList(context.Context, *_go.UserIDRequest) (*_go.TokenStatusList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenStatusList not implemented")
}
func (UnimplementedClientTokenServiceServer) DeleteToken(context.Context, *_go.UserTokenIDRequest) (*_go.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedClientTokenServiceServer) mustEmbedUnimplementedClientTokenServiceServer() {}

// UnsafeClientTokenServiceServer may be embedded to opt out of forward compatibility for this API.
// Use of this interface is not recommended, as added methods to ClientTokenServiceServer will
// result in compilation errors.
type UnsafeClientTokenServiceServer interface {
	mustEmbedUnimplementedClientTokenServiceServer()
}

func RegisterClientTokenServiceServer(s grpc.ServiceRegistrar, srv ClientTokenServiceServer) {
	s.RegisterService(&ClientTokenService_ServiceDesc, srv)
}

func _ClientTokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.ClientTokenService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTokenServiceServer).CreateToken(ctx, req.(*_go.TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTokenService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserTokenIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTokenServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.ClientTokenService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTokenServiceServer).GetToken(ctx, req.(*_go.UserTokenIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTokenService_GetTokenStatusList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTokenServiceServer).GetTokenStatusList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.ClientTokenService/GetTokenStatusList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTokenServiceServer).GetTokenStatusList(ctx, req.(*_go.UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTokenService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserTokenIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTokenServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.ClientTokenService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTokenServiceServer).DeleteToken(ctx, req.(*_go.UserTokenIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientTokenService_ServiceDesc is the grpc.ServiceDesc for ClientTokenService API.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha2.user.ClientTokenService",
	HandlerType: (*ClientTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _ClientTokenService_CreateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _ClientTokenService_GetToken_Handler,
		},
		{
			MethodName: "GetTokenStatusList",
			Handler:    _ClientTokenService_GetTokenStatusList_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _ClientTokenService_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/user.proto",
}

// UserManagerClient is the client API for UserManager API.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagerClient interface {
	AddUserToGroup(ctx context.Context, in *_go.UserGroupRequest, opts ...grpc.CallOption) (*_go.StatusReport, error)
	GetUser(ctx context.Context, in *_go.UserIDRequest, opts ...grpc.CallOption) (*_go.User, error)
	GetUserPermissions(ctx context.Context, in *_go.UserIDRequest, opts ...grpc.CallOption) (*_go.UserPermissionResponse, error)
	ChangeUserPermission(ctx context.Context, in *_go.ChangePermissionRequest, opts ...grpc.CallOption) (*_go.StatusReport, error)
	DeleteUser(ctx context.Context, in *_go.DeleteUserRequest, opts ...grpc.CallOption) (*_go.StatusReport, error)
	RemoveUserFromGroup(ctx context.Context, in *_go.UserGroupRequest, opts ...grpc.CallOption) (*_go.StatusReport, error)
}

type userManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagerClient(cc grpc.ClientConnInterface) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) AddUserToGroup(ctx context.Context, in *_go.UserGroupRequest, opts ...grpc.CallOption) (*_go.StatusReport, error) {
	out := new(_go.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.UserManager/AddUserToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) GetUser(ctx context.Context, in *_go.UserIDRequest, opts ...grpc.CallOption) (*_go.User, error) {
	out := new(_go.User)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.UserManager/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) GetUserPermissions(ctx context.Context, in *_go.UserIDRequest, opts ...grpc.CallOption) (*_go.UserPermissionResponse, error) {
	out := new(_go.UserPermissionResponse)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.UserManager/GetUserPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) ChangeUserPermission(ctx context.Context, in *_go.ChangePermissionRequest, opts ...grpc.CallOption) (*_go.StatusReport, error) {
	out := new(_go.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.UserManager/ChangeUserPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) DeleteUser(ctx context.Context, in *_go.DeleteUserRequest, opts ...grpc.CallOption) (*_go.StatusReport, error) {
	out := new(_go.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.UserManager/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) RemoveUserFromGroup(ctx context.Context, in *_go.UserGroupRequest, opts ...grpc.CallOption) (*_go.StatusReport, error) {
	out := new(_go.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.user.UserManager/RemoveUserFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagerServer is the server API for UserManager API.
// All implementations must embed UnimplementedUserManagerServer
// for forward compatibility
type UserManagerServer interface {
	AddUserToGroup(context.Context, *_go.UserGroupRequest) (*_go.StatusReport, error)
	GetUser(context.Context, *_go.UserIDRequest) (*_go.User, error)
	GetUserPermissions(context.Context, *_go.UserIDRequest) (*_go.UserPermissionResponse, error)
	ChangeUserPermission(context.Context, *_go.ChangePermissionRequest) (*_go.StatusReport, error)
	DeleteUser(context.Context, *_go.DeleteUserRequest) (*_go.StatusReport, error)
	RemoveUserFromGroup(context.Context, *_go.UserGroupRequest) (*_go.StatusReport, error)
	mustEmbedUnimplementedUserManagerServer()
}

// UnimplementedUserManagerServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagerServer struct {
}

func (UnimplementedUserManagerServer) AddUserToGroup(context.Context, *_go.UserGroupRequest) (*_go.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToGroup not implemented")
}
func (UnimplementedUserManagerServer) GetUser(context.Context, *_go.UserIDRequest) (*_go.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserManagerServer) GetUserPermissions(context.Context, *_go.UserIDRequest) (*_go.UserPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPermissions not implemented")
}
func (UnimplementedUserManagerServer) ChangeUserPermission(context.Context, *_go.ChangePermissionRequest) (*_go.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPermission not implemented")
}
func (UnimplementedUserManagerServer) DeleteUser(context.Context, *_go.DeleteUserRequest) (*_go.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagerServer) RemoveUserFromGroup(context.Context, *_go.UserGroupRequest) (*_go.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromGroup not implemented")
}
func (UnimplementedUserManagerServer) mustEmbedUnimplementedUserManagerServer() {}

// UnsafeUserManagerServer may be embedded to opt out of forward compatibility for this API.
// Use of this interface is not recommended, as added methods to UserManagerServer will
// result in compilation errors.
type UnsafeUserManagerServer interface {
	mustEmbedUnimplementedUserManagerServer()
}

func RegisterUserManagerServer(s grpc.ServiceRegistrar, srv UserManagerServer) {
	s.RegisterService(&UserManager_ServiceDesc, srv)
}

func _UserManager_AddUserToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).AddUserToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.UserManager/AddUserToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).AddUserToGroup(ctx, req.(*_go.UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.UserManager/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetUser(ctx, req.(*_go.UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_GetUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetUserPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.UserManager/GetUserPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetUserPermissions(ctx, req.(*_go.UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_ChangeUserPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.ChangePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).ChangeUserPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.UserManager/ChangeUserPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).ChangeUserPermission(ctx, req.(*_go.ChangePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.UserManager/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).DeleteUser(ctx, req.(*_go.DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_RemoveUserFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).RemoveUserFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.user.UserManager/RemoveUserFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).RemoveUserFromGroup(ctx, req.(*_go.UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManager_ServiceDesc is the grpc.ServiceDesc for UserManager API.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha2.user.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUserToGroup",
			Handler:    _UserManager_AddUserToGroup_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserManager_GetUser_Handler,
		},
		{
			MethodName: "GetUserPermissions",
			Handler:    _UserManager_GetUserPermissions_Handler,
		},
		{
			MethodName: "ChangeUserPermission",
			Handler:    _UserManager_ChangeUserPermission_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManager_DeleteUser_Handler,
		},
		{
			MethodName: "RemoveUserFromGroup",
			Handler:    _UserManager_RemoveUserFromGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/user.proto",
}