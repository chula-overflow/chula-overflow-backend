// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: thread.proto

package proto

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

// ThreadClient is the client API for Thread service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThreadClient interface {
	CreateThread(ctx context.Context, in *ThreadRequestCreateBody, opts ...grpc.CallOption) (*ThreadBody, error)
	GetThread(ctx context.Context, in *ThreadPropertyRequestBody, opts ...grpc.CallOption) (*ThreadBodyArr, error)
	GetThreadById(ctx context.Context, in *ThreadIdRequestBody, opts ...grpc.CallOption) (*ThreadBody, error)
	UpvoteThread(ctx context.Context, in *ThreadIdRequestBody, opts ...grpc.CallOption) (*ThreadBody, error)
	DownvoteThread(ctx context.Context, in *ThreadIdRequestBody, opts ...grpc.CallOption) (*ThreadBody, error)
	UpvoteProblem(ctx context.Context, in *ThreadProblemIdRequestBody, opts ...grpc.CallOption) (*ProblemBody, error)
	DownvoteProblem(ctx context.Context, in *ThreadProblemIdRequestBody, opts ...grpc.CallOption) (*ProblemBody, error)
	UpvoteAnswer(ctx context.Context, in *ThreadAnswerIdRequestBody, opts ...grpc.CallOption) (*AnswerBody, error)
	DownvoteAnswer(ctx context.Context, in *ThreadAnswerIdRequestBody, opts ...grpc.CallOption) (*AnswerBody, error)
	AddAnswer(ctx context.Context, in *AnswerRequestCreateBody, opts ...grpc.CallOption) (*ThreadBody, error)
}

type threadClient struct {
	cc grpc.ClientConnInterface
}

func NewThreadClient(cc grpc.ClientConnInterface) ThreadClient {
	return &threadClient{cc}
}

func (c *threadClient) CreateThread(ctx context.Context, in *ThreadRequestCreateBody, opts ...grpc.CallOption) (*ThreadBody, error) {
	out := new(ThreadBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/CreateThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) GetThread(ctx context.Context, in *ThreadPropertyRequestBody, opts ...grpc.CallOption) (*ThreadBodyArr, error) {
	out := new(ThreadBodyArr)
	err := c.cc.Invoke(ctx, "/thread.Thread/GetThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) GetThreadById(ctx context.Context, in *ThreadIdRequestBody, opts ...grpc.CallOption) (*ThreadBody, error) {
	out := new(ThreadBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/GetThreadById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) UpvoteThread(ctx context.Context, in *ThreadIdRequestBody, opts ...grpc.CallOption) (*ThreadBody, error) {
	out := new(ThreadBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/UpvoteThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) DownvoteThread(ctx context.Context, in *ThreadIdRequestBody, opts ...grpc.CallOption) (*ThreadBody, error) {
	out := new(ThreadBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/DownvoteThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) UpvoteProblem(ctx context.Context, in *ThreadProblemIdRequestBody, opts ...grpc.CallOption) (*ProblemBody, error) {
	out := new(ProblemBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/UpvoteProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) DownvoteProblem(ctx context.Context, in *ThreadProblemIdRequestBody, opts ...grpc.CallOption) (*ProblemBody, error) {
	out := new(ProblemBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/DownvoteProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) UpvoteAnswer(ctx context.Context, in *ThreadAnswerIdRequestBody, opts ...grpc.CallOption) (*AnswerBody, error) {
	out := new(AnswerBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/UpvoteAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) DownvoteAnswer(ctx context.Context, in *ThreadAnswerIdRequestBody, opts ...grpc.CallOption) (*AnswerBody, error) {
	out := new(AnswerBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/DownvoteAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) AddAnswer(ctx context.Context, in *AnswerRequestCreateBody, opts ...grpc.CallOption) (*ThreadBody, error) {
	out := new(ThreadBody)
	err := c.cc.Invoke(ctx, "/thread.Thread/AddAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThreadServer is the server API for Thread service.
// All implementations should embed UnimplementedThreadServer
// for forward compatibility
type ThreadServer interface {
	CreateThread(context.Context, *ThreadRequestCreateBody) (*ThreadBody, error)
	GetThread(context.Context, *ThreadPropertyRequestBody) (*ThreadBodyArr, error)
	GetThreadById(context.Context, *ThreadIdRequestBody) (*ThreadBody, error)
	UpvoteThread(context.Context, *ThreadIdRequestBody) (*ThreadBody, error)
	DownvoteThread(context.Context, *ThreadIdRequestBody) (*ThreadBody, error)
	UpvoteProblem(context.Context, *ThreadProblemIdRequestBody) (*ProblemBody, error)
	DownvoteProblem(context.Context, *ThreadProblemIdRequestBody) (*ProblemBody, error)
	UpvoteAnswer(context.Context, *ThreadAnswerIdRequestBody) (*AnswerBody, error)
	DownvoteAnswer(context.Context, *ThreadAnswerIdRequestBody) (*AnswerBody, error)
	AddAnswer(context.Context, *AnswerRequestCreateBody) (*ThreadBody, error)
}

// UnimplementedThreadServer should be embedded to have forward compatible implementations.
type UnimplementedThreadServer struct {
}

func (UnimplementedThreadServer) CreateThread(context.Context, *ThreadRequestCreateBody) (*ThreadBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThread not implemented")
}
func (UnimplementedThreadServer) GetThread(context.Context, *ThreadPropertyRequestBody) (*ThreadBodyArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThread not implemented")
}
func (UnimplementedThreadServer) GetThreadById(context.Context, *ThreadIdRequestBody) (*ThreadBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThreadById not implemented")
}
func (UnimplementedThreadServer) UpvoteThread(context.Context, *ThreadIdRequestBody) (*ThreadBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteThread not implemented")
}
func (UnimplementedThreadServer) DownvoteThread(context.Context, *ThreadIdRequestBody) (*ThreadBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteThread not implemented")
}
func (UnimplementedThreadServer) UpvoteProblem(context.Context, *ThreadProblemIdRequestBody) (*ProblemBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteProblem not implemented")
}
func (UnimplementedThreadServer) DownvoteProblem(context.Context, *ThreadProblemIdRequestBody) (*ProblemBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteProblem not implemented")
}
func (UnimplementedThreadServer) UpvoteAnswer(context.Context, *ThreadAnswerIdRequestBody) (*AnswerBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteAnswer not implemented")
}
func (UnimplementedThreadServer) DownvoteAnswer(context.Context, *ThreadAnswerIdRequestBody) (*AnswerBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteAnswer not implemented")
}
func (UnimplementedThreadServer) AddAnswer(context.Context, *AnswerRequestCreateBody) (*ThreadBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnswer not implemented")
}

// UnsafeThreadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThreadServer will
// result in compilation errors.
type UnsafeThreadServer interface {
	mustEmbedUnimplementedThreadServer()
}

func RegisterThreadServer(s grpc.ServiceRegistrar, srv ThreadServer) {
	s.RegisterService(&Thread_ServiceDesc, srv)
}

func _Thread_CreateThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadRequestCreateBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).CreateThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/CreateThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).CreateThread(ctx, req.(*ThreadRequestCreateBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_GetThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadPropertyRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).GetThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/GetThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).GetThread(ctx, req.(*ThreadPropertyRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_GetThreadById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadIdRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).GetThreadById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/GetThreadById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).GetThreadById(ctx, req.(*ThreadIdRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_UpvoteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadIdRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).UpvoteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/UpvoteThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).UpvoteThread(ctx, req.(*ThreadIdRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_DownvoteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadIdRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).DownvoteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/DownvoteThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).DownvoteThread(ctx, req.(*ThreadIdRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_UpvoteProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadProblemIdRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).UpvoteProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/UpvoteProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).UpvoteProblem(ctx, req.(*ThreadProblemIdRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_DownvoteProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadProblemIdRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).DownvoteProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/DownvoteProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).DownvoteProblem(ctx, req.(*ThreadProblemIdRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_UpvoteAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadAnswerIdRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).UpvoteAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/UpvoteAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).UpvoteAnswer(ctx, req.(*ThreadAnswerIdRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_DownvoteAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadAnswerIdRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).DownvoteAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/DownvoteAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).DownvoteAnswer(ctx, req.(*ThreadAnswerIdRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_AddAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerRequestCreateBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).AddAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/AddAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).AddAnswer(ctx, req.(*AnswerRequestCreateBody))
	}
	return interceptor(ctx, in, info, handler)
}

// Thread_ServiceDesc is the grpc.ServiceDesc for Thread service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Thread_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thread.Thread",
	HandlerType: (*ThreadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateThread",
			Handler:    _Thread_CreateThread_Handler,
		},
		{
			MethodName: "GetThread",
			Handler:    _Thread_GetThread_Handler,
		},
		{
			MethodName: "GetThreadById",
			Handler:    _Thread_GetThreadById_Handler,
		},
		{
			MethodName: "UpvoteThread",
			Handler:    _Thread_UpvoteThread_Handler,
		},
		{
			MethodName: "DownvoteThread",
			Handler:    _Thread_DownvoteThread_Handler,
		},
		{
			MethodName: "UpvoteProblem",
			Handler:    _Thread_UpvoteProblem_Handler,
		},
		{
			MethodName: "DownvoteProblem",
			Handler:    _Thread_DownvoteProblem_Handler,
		},
		{
			MethodName: "UpvoteAnswer",
			Handler:    _Thread_UpvoteAnswer_Handler,
		},
		{
			MethodName: "DownvoteAnswer",
			Handler:    _Thread_DownvoteAnswer_Handler,
		},
		{
			MethodName: "AddAnswer",
			Handler:    _Thread_AddAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thread.proto",
}
