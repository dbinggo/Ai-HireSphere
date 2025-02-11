// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: application/user-center/interfaces/rpc/user.proto

package user

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
	User_FindUserById_FullMethodName    = "/user.User/FindUserById"
	User_FindUserByPhone_FullMethodName = "/user.User/FindUserByPhone"
	User_OssUpload_FullMethodName       = "/user.User/OssUpload"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	FindUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfo, error)
	FindUserByPhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*UserInfo, error)
	OssUpload(ctx context.Context, in *OSSUploadReq, opts ...grpc.CallOption) (*OSSUploadResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) FindUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, User_FindUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FindUserByPhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*UserInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, User_FindUserByPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) OssUpload(ctx context.Context, in *OSSUploadReq, opts ...grpc.CallOption) (*OSSUploadResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OSSUploadResp)
	err := c.cc.Invoke(ctx, User_OssUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility.
type UserServer interface {
	FindUserById(context.Context, *Id) (*UserInfo, error)
	FindUserByPhone(context.Context, *Phone) (*UserInfo, error)
	OssUpload(context.Context, *OSSUploadReq) (*OSSUploadResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServer struct{}

func (UnimplementedUserServer) FindUserById(context.Context, *Id) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserById not implemented")
}
func (UnimplementedUserServer) FindUserByPhone(context.Context, *Phone) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByPhone not implemented")
}
func (UnimplementedUserServer) OssUpload(context.Context, *OSSUploadReq) (*OSSUploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssUpload not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}
func (UnimplementedUserServer) testEmbeddedByValue()              {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	// If the following call pancis, it indicates UnimplementedUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_FindUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FindUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindUserById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FindUserByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindUserByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FindUserByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindUserByPhone(ctx, req.(*Phone))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_OssUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OSSUploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).OssUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_OssUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).OssUpload(ctx, req.(*OSSUploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserById",
			Handler:    _User_FindUserById_Handler,
		},
		{
			MethodName: "FindUserByPhone",
			Handler:    _User_FindUserByPhone_Handler,
		},
		{
			MethodName: "OssUpload",
			Handler:    _User_OssUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/user-center/interfaces/rpc/user.proto",
}
