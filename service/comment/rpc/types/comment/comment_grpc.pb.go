// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: comment/rpc/comment.proto

package comment

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// -----------------------Comment-----------------------
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
	UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error)
	DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error)
	GetCommentById(ctx context.Context, in *GetCommentByIdReq, opts ...grpc.CallOption) (*GetCommentByIdResp, error)
	SearchComment(ctx context.Context, in *SearchCommentReq, opts ...grpc.CallOption) (*SearchCommentResp, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, "/comment.service/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error) {
	out := new(UpdateCommentResp)
	err := c.cc.Invoke(ctx, "/comment.service/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error) {
	out := new(DelCommentResp)
	err := c.cc.Invoke(ctx, "/comment.service/DelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetCommentById(ctx context.Context, in *GetCommentByIdReq, opts ...grpc.CallOption) (*GetCommentByIdResp, error) {
	out := new(GetCommentByIdResp)
	err := c.cc.Invoke(ctx, "/comment.service/GetCommentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SearchComment(ctx context.Context, in *SearchCommentReq, opts ...grpc.CallOption) (*SearchCommentResp, error) {
	out := new(SearchCommentResp)
	err := c.cc.Invoke(ctx, "/comment.service/SearchComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// -----------------------Comment-----------------------
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error)
	DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error)
	GetCommentById(context.Context, *GetCommentByIdReq) (*GetCommentByIdResp, error)
	SearchComment(context.Context, *SearchCommentReq) (*SearchCommentResp, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedServiceServer) UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedServiceServer) DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelComment not implemented")
}
func (UnimplementedServiceServer) GetCommentById(context.Context, *GetCommentByIdReq) (*GetCommentByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedServiceServer) SearchComment(context.Context, *SearchCommentReq) (*SearchCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchComment not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.service/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.service/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateComment(ctx, req.(*UpdateCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.service/DelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DelComment(ctx, req.(*DelCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.service/GetCommentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetCommentById(ctx, req.(*GetCommentByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_SearchComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SearchComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.service/SearchComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SearchComment(ctx, req.(*SearchCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _Service_AddComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Service_UpdateComment_Handler,
		},
		{
			MethodName: "DelComment",
			Handler:    _Service_DelComment_Handler,
		},
		{
			MethodName: "GetCommentById",
			Handler:    _Service_GetCommentById_Handler,
		},
		{
			MethodName: "SearchComment",
			Handler:    _Service_SearchComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment/rpc/comment.proto",
}
