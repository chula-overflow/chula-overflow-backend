// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: course.proto

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

// CourseClient is the client API for Course service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseClient interface {
	GetAllCourseSummary(ctx context.Context, in *GetAllCourseSummaryRequest, opts ...grpc.CallOption) (*GetAllCourseSummaryResponse, error)
	GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseResponse, error)
}

type courseClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseClient(cc grpc.ClientConnInterface) CourseClient {
	return &courseClient{cc}
}

func (c *courseClient) GetAllCourseSummary(ctx context.Context, in *GetAllCourseSummaryRequest, opts ...grpc.CallOption) (*GetAllCourseSummaryResponse, error) {
	out := new(GetAllCourseSummaryResponse)
	err := c.cc.Invoke(ctx, "/course.Course/GetAllCourseSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseResponse, error) {
	out := new(GetCourseResponse)
	err := c.cc.Invoke(ctx, "/course.Course/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServer is the server API for Course service.
// All implementations should embed UnimplementedCourseServer
// for forward compatibility
type CourseServer interface {
	GetAllCourseSummary(context.Context, *GetAllCourseSummaryRequest) (*GetAllCourseSummaryResponse, error)
	GetCourse(context.Context, *GetCourseRequest) (*GetCourseResponse, error)
}

// UnimplementedCourseServer should be embedded to have forward compatible implementations.
type UnimplementedCourseServer struct {
}

func (UnimplementedCourseServer) GetAllCourseSummary(context.Context, *GetAllCourseSummaryRequest) (*GetAllCourseSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCourseSummary not implemented")
}
func (UnimplementedCourseServer) GetCourse(context.Context, *GetCourseRequest) (*GetCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}

// UnsafeCourseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServer will
// result in compilation errors.
type UnsafeCourseServer interface {
	mustEmbedUnimplementedCourseServer()
}

func RegisterCourseServer(s grpc.ServiceRegistrar, srv CourseServer) {
	s.RegisterService(&Course_ServiceDesc, srv)
}

func _Course_GetAllCourseSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCourseSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).GetAllCourseSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.Course/GetAllCourseSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).GetAllCourseSummary(ctx, req.(*GetAllCourseSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.Course/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).GetCourse(ctx, req.(*GetCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Course_ServiceDesc is the grpc.ServiceDesc for Course service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Course_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "course.Course",
	HandlerType: (*CourseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCourseSummary",
			Handler:    _Course_GetAllCourseSummary_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _Course_GetCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "course.proto",
}
