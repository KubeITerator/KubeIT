// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	CreateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Token, error)
	GetToken(ctx context.Context, in *ClientTokenIDRequest, opts ...grpc.CallOption) (*TokenStatus, error)
	GetTokenStatusList(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*TokenStatusList, error)
	DeleteToken(ctx context.Context, in *ClientTokenIDRequest, opts ...grpc.CallOption) (*StatusReport, error)
	CreateOrUpdateClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*ClientID, error)
	GetClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*Client, error)
	GetClientInfo(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientInfo, error)
	DeleteClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*StatusReport, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) CreateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetToken(ctx context.Context, in *ClientTokenIDRequest, opts ...grpc.CallOption) (*TokenStatus, error) {
	out := new(TokenStatus)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetTokenStatusList(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*TokenStatusList, error) {
	out := new(TokenStatusList)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/GetTokenStatusList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteToken(ctx context.Context, in *ClientTokenIDRequest, opts ...grpc.CallOption) (*StatusReport, error) {
	out := new(StatusReport)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CreateOrUpdateClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*ClientID, error) {
	out := new(ClientID)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/CreateOrUpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/GetClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientInfo(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientInfo, error) {
	out := new(ClientInfo)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/GetClientInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*StatusReport, error) {
	out := new(StatusReport)
	err := c.cc.Invoke(ctx, "/kubeit.v1beta1.ClientService/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations should embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	CreateToken(context.Context, *TokenRequest) (*Token, error)
	GetToken(context.Context, *ClientTokenIDRequest) (*TokenStatus, error)
	GetTokenStatusList(context.Context, *ClientID) (*TokenStatusList, error)
	DeleteToken(context.Context, *ClientTokenIDRequest) (*StatusReport, error)
	CreateOrUpdateClient(context.Context, *Client) (*ClientID, error)
	GetClient(context.Context, *ClientID) (*Client, error)
	GetClientInfo(context.Context, *ClientID) (*ClientInfo, error)
	DeleteClient(context.Context, *ClientID) (*StatusReport, error)
}

// UnimplementedClientServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) CreateToken(context.Context, *TokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedClientServiceServer) GetToken(context.Context, *ClientTokenIDRequest) (*TokenStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedClientServiceServer) GetTokenStatusList(context.Context, *ClientID) (*TokenStatusList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenStatusList not implemented")
}
func (UnimplementedClientServiceServer) DeleteToken(context.Context, *ClientTokenIDRequest) (*StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedClientServiceServer) CreateOrUpdateClient(context.Context, *Client) (*ClientID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateClient not implemented")
}
func (UnimplementedClientServiceServer) GetClient(context.Context, *ClientID) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedClientServiceServer) GetClientInfo(context.Context, *ClientID) (*ClientInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientInfo not implemented")
}
func (UnimplementedClientServiceServer) DeleteClient(context.Context, *ClientID) (*StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientTokenIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetToken(ctx, req.(*ClientTokenIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetTokenStatusList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetTokenStatusList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/GetTokenStatusList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetTokenStatusList(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientTokenIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteToken(ctx, req.(*ClientTokenIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CreateOrUpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateOrUpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/CreateOrUpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateOrUpdateClient(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/GetClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClient(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/GetClientInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientInfo(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeit.v1beta1.ClientService/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteClient(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeit.v1beta1.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _ClientService_CreateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _ClientService_GetToken_Handler,
		},
		{
			MethodName: "GetTokenStatusList",
			Handler:    _ClientService_GetTokenStatusList_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _ClientService_DeleteToken_Handler,
		},
		{
			MethodName: "CreateOrUpdateClient",
			Handler:    _ClientService_CreateOrUpdateClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _ClientService_GetClient_Handler,
		},
		{
			MethodName: "GetClientInfo",
			Handler:    _ClientService_GetClientInfo_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _ClientService_DeleteClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/user.proto",
}