// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: application/interview/interfaces/rpc/interview.proto

package interview

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Interview_Ping_FullMethodName = "/interview.Interview/Ping"
)

// InterviewClient is the client API for Interview service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterviewClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type interviewClient struct {
	cc grpc.ClientConnInterface
}

func NewInterviewClient(cc grpc.ClientConnInterface) InterviewClient {
	return &interviewClient{cc}
}

func (c *interviewClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Interview_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterviewServer is the server API for Interview service.
// All implementations must embed UnimplementedInterviewServer
// for forward compatibility.
type InterviewServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedInterviewServer()
}

// UnimplementedInterviewServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInterviewServer struct{}

func (UnimplementedInterviewServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedInterviewServer) mustEmbedUnimplementedInterviewServer() {}
func (UnimplementedInterviewServer) testEmbeddedByValue()                   {}

// UnsafeInterviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterviewServer will
// result in compilation errors.
type UnsafeInterviewServer interface {
	mustEmbedUnimplementedInterviewServer()
}

func RegisterInterviewServer(s grpc.ServiceRegistrar, srv InterviewServer) {
	// If the following call pancis, it indicates UnimplementedInterviewServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Interview_ServiceDesc, srv)
}

func _Interview_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterviewServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Interview_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterviewServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Interview_ServiceDesc is the grpc.ServiceDesc for Interview service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interview_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interview.Interview",
	HandlerType: (*InterviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Interview_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/interview/interfaces/rpc/interview.proto",
}
