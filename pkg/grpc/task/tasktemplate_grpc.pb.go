// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package task

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "kubeIT/pkg/grpc/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TaskManagementServiceClient is the client API for TaskManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskManagementServiceClient interface {
	CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*common.StatusReport, error)
	GetTask(ctx context.Context, in *common.IDRequest, opts ...grpc.CallOption) (*common.StatusReport, error)
	DeleteGroup(ctx context.Context, in *common.IDRequest, opts ...grpc.CallOption) (*common.StatusReport, error)
}

type taskManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskManagementServiceClient(cc grpc.ClientConnInterface) TaskManagementServiceClient {
	return &taskManagementServiceClient{cc}
}

func (c *taskManagementServiceClient) CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*common.StatusReport, error) {
	out := new(common.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.task.TaskManagementService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementServiceClient) GetTask(ctx context.Context, in *common.IDRequest, opts ...grpc.CallOption) (*common.StatusReport, error) {
	out := new(common.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.task.TaskManagementService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementServiceClient) DeleteGroup(ctx context.Context, in *common.IDRequest, opts ...grpc.CallOption) (*common.StatusReport, error) {
	out := new(common.StatusReport)
	err := c.cc.Invoke(ctx, "/v1alpha2.task.TaskManagementService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskManagementServiceServer is the server API for TaskManagementService service.
// All implementations must embed UnimplementedTaskManagementServiceServer
// for forward compatibility
type TaskManagementServiceServer interface {
	CreateTask(context.Context, *Task) (*common.StatusReport, error)
	GetTask(context.Context, *common.IDRequest) (*common.StatusReport, error)
	DeleteGroup(context.Context, *common.IDRequest) (*common.StatusReport, error)
	mustEmbedUnimplementedTaskManagementServiceServer()
}

// UnimplementedTaskManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskManagementServiceServer struct {
}

func (UnimplementedTaskManagementServiceServer) CreateTask(context.Context, *Task) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskManagementServiceServer) GetTask(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskManagementServiceServer) DeleteGroup(context.Context, *common.IDRequest) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedTaskManagementServiceServer) mustEmbedUnimplementedTaskManagementServiceServer() {}

// UnsafeTaskManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskManagementServiceServer will
// result in compilation errors.
type UnsafeTaskManagementServiceServer interface {
	mustEmbedUnimplementedTaskManagementServiceServer()
}

func RegisterTaskManagementServiceServer(s grpc.ServiceRegistrar, srv TaskManagementServiceServer) {
	s.RegisterService(&TaskManagementService_ServiceDesc, srv)
}

func _TaskManagementService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.task.TaskManagementService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServiceServer).CreateTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagementService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.task.TaskManagementService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServiceServer).GetTask(ctx, req.(*common.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagementService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.task.TaskManagementService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServiceServer).DeleteGroup(ctx, req.(*common.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskManagementService_ServiceDesc is the grpc.ServiceDesc for TaskManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha2.task.TaskManagementService",
	HandlerType: (*TaskManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManagementService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskManagementService_GetTask_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _TaskManagementService_DeleteGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/task/tasktemplate.proto",
}