// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: runme/runner/v2alpha1/runner.proto

package runnerv1

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

const (
	RunnerService_CreateSession_FullMethodName = "/runme.runner.v2alpha1.RunnerService/CreateSession"
	RunnerService_GetSession_FullMethodName    = "/runme.runner.v2alpha1.RunnerService/GetSession"
	RunnerService_ListSessions_FullMethodName  = "/runme.runner.v2alpha1.RunnerService/ListSessions"
	RunnerService_UpdateSession_FullMethodName = "/runme.runner.v2alpha1.RunnerService/UpdateSession"
	RunnerService_DeleteSession_FullMethodName = "/runme.runner.v2alpha1.RunnerService/DeleteSession"
	RunnerService_Execute_FullMethodName       = "/runme.runner.v2alpha1.RunnerService/Execute"
	RunnerService_ResolveEnv_FullMethodName    = "/runme.runner.v2alpha1.RunnerService/ResolveEnv"
)

// RunnerServiceClient is the client API for RunnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunnerServiceClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error)
	ListSessions(ctx context.Context, in *ListSessionsRequest, opts ...grpc.CallOption) (*ListSessionsResponse, error)
	UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error)
	// Execute executes a program. Examine "ExecuteRequest" to explore
	// configuration options.
	//
	// It's a bidirectional stream RPC method. It expects the first
	// "ExecuteRequest" to contain details of a program to execute.
	// Subsequent "ExecuteRequest" should only contain "input_data" as
	// other fields will be ignored.
	Execute(ctx context.Context, opts ...grpc.CallOption) (RunnerService_ExecuteClient, error)
	// ResolveEnv resolves environment variables from a script or a list of commands
	// using the provided sources, which can be a list of environment variables,
	// a session, or a project.
	// The result contains all found environment variables. If the env is in any source,
	// it is considered resolved. Otherwise, it is makred as unresolved.
	ResolveEnv(ctx context.Context, in *ResolveEnvRequest, opts ...grpc.CallOption) (*ResolveEnvResponse, error)
}

type runnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerServiceClient(cc grpc.ClientConnInterface) RunnerServiceClient {
	return &runnerServiceClient{cc}
}

func (c *runnerServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, RunnerService_CreateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error) {
	out := new(GetSessionResponse)
	err := c.cc.Invoke(ctx, RunnerService_GetSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerServiceClient) ListSessions(ctx context.Context, in *ListSessionsRequest, opts ...grpc.CallOption) (*ListSessionsResponse, error) {
	out := new(ListSessionsResponse)
	err := c.cc.Invoke(ctx, RunnerService_ListSessions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerServiceClient) UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error) {
	out := new(UpdateSessionResponse)
	err := c.cc.Invoke(ctx, RunnerService_UpdateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error) {
	out := new(DeleteSessionResponse)
	err := c.cc.Invoke(ctx, RunnerService_DeleteSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerServiceClient) Execute(ctx context.Context, opts ...grpc.CallOption) (RunnerService_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RunnerService_ServiceDesc.Streams[0], RunnerService_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &runnerServiceExecuteClient{stream}
	return x, nil
}

type RunnerService_ExecuteClient interface {
	Send(*ExecuteRequest) error
	Recv() (*ExecuteResponse, error)
	grpc.ClientStream
}

type runnerServiceExecuteClient struct {
	grpc.ClientStream
}

func (x *runnerServiceExecuteClient) Send(m *ExecuteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runnerServiceExecuteClient) Recv() (*ExecuteResponse, error) {
	m := new(ExecuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runnerServiceClient) ResolveEnv(ctx context.Context, in *ResolveEnvRequest, opts ...grpc.CallOption) (*ResolveEnvResponse, error) {
	out := new(ResolveEnvResponse)
	err := c.cc.Invoke(ctx, RunnerService_ResolveEnv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerServiceServer is the server API for RunnerService service.
// All implementations must embed UnimplementedRunnerServiceServer
// for forward compatibility
type RunnerServiceServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error)
	ListSessions(context.Context, *ListSessionsRequest) (*ListSessionsResponse, error)
	UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error)
	// Execute executes a program. Examine "ExecuteRequest" to explore
	// configuration options.
	//
	// It's a bidirectional stream RPC method. It expects the first
	// "ExecuteRequest" to contain details of a program to execute.
	// Subsequent "ExecuteRequest" should only contain "input_data" as
	// other fields will be ignored.
	Execute(RunnerService_ExecuteServer) error
	// ResolveEnv resolves environment variables from a script or a list of commands
	// using the provided sources, which can be a list of environment variables,
	// a session, or a project.
	// The result contains all found environment variables. If the env is in any source,
	// it is considered resolved. Otherwise, it is makred as unresolved.
	ResolveEnv(context.Context, *ResolveEnvRequest) (*ResolveEnvResponse, error)
	mustEmbedUnimplementedRunnerServiceServer()
}

// UnimplementedRunnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRunnerServiceServer struct {
}

func (UnimplementedRunnerServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedRunnerServiceServer) GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedRunnerServiceServer) ListSessions(context.Context, *ListSessionsRequest) (*ListSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSessions not implemented")
}
func (UnimplementedRunnerServiceServer) UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSession not implemented")
}
func (UnimplementedRunnerServiceServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedRunnerServiceServer) Execute(RunnerService_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedRunnerServiceServer) ResolveEnv(context.Context, *ResolveEnvRequest) (*ResolveEnvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveEnv not implemented")
}
func (UnimplementedRunnerServiceServer) mustEmbedUnimplementedRunnerServiceServer() {}

// UnsafeRunnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerServiceServer will
// result in compilation errors.
type UnsafeRunnerServiceServer interface {
	mustEmbedUnimplementedRunnerServiceServer()
}

func RegisterRunnerServiceServer(s grpc.ServiceRegistrar, srv RunnerServiceServer) {
	s.RegisterService(&RunnerService_ServiceDesc, srv)
}

func _RunnerService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerService_GetSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerService_ListSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServiceServer).ListSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerService_ListSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServiceServer).ListSessions(ctx, req.(*ListSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerService_UpdateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServiceServer).UpdateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerService_UpdateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServiceServer).UpdateSession(ctx, req.(*UpdateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerService_DeleteSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerService_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RunnerServiceServer).Execute(&runnerServiceExecuteServer{stream})
}

type RunnerService_ExecuteServer interface {
	Send(*ExecuteResponse) error
	Recv() (*ExecuteRequest, error)
	grpc.ServerStream
}

type runnerServiceExecuteServer struct {
	grpc.ServerStream
}

func (x *runnerServiceExecuteServer) Send(m *ExecuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runnerServiceExecuteServer) Recv() (*ExecuteRequest, error) {
	m := new(ExecuteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RunnerService_ResolveEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServiceServer).ResolveEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerService_ResolveEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServiceServer).ResolveEnv(ctx, req.(*ResolveEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RunnerService_ServiceDesc is the grpc.ServiceDesc for RunnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RunnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runme.runner.v2alpha1.RunnerService",
	HandlerType: (*RunnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _RunnerService_CreateSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _RunnerService_GetSession_Handler,
		},
		{
			MethodName: "ListSessions",
			Handler:    _RunnerService_ListSessions_Handler,
		},
		{
			MethodName: "UpdateSession",
			Handler:    _RunnerService_UpdateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _RunnerService_DeleteSession_Handler,
		},
		{
			MethodName: "ResolveEnv",
			Handler:    _RunnerService_ResolveEnv_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _RunnerService_Execute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "runme/runner/v2alpha1/runner.proto",
}
