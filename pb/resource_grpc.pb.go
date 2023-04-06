// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: resource.proto

package pb

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
	ResourceService_InsertResource_FullMethodName     = "/pb.ResourceService/InsertResource"
	ResourceService_GetResourceById_FullMethodName    = "/pb.ResourceService/GetResourceById"
	ResourceService_GetResourceByEmail_FullMethodName = "/pb.ResourceService/GetResourceByEmail"
	ResourceService_GetAllResources_FullMethodName    = "/pb.ResourceService/GetAllResources"
	ResourceService_UpdateResource_FullMethodName     = "/pb.ResourceService/UpdateResource"
)

// ResourceServiceClient is the client API for ResourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceServiceClient interface {
	InsertResource(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetResourceById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetResourceByEmail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetAllResources(ctx context.Context, in *RequestWithFilters, opts ...grpc.CallOption) (*ArrayResponse, error)
	UpdateResource(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type resourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceServiceClient(cc grpc.ClientConnInterface) ResourceServiceClient {
	return &resourceServiceClient{cc}
}

func (c *resourceServiceClient) InsertResource(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ResourceService_InsertResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetResourceById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ResourceService_GetResourceById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetResourceByEmail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ResourceService_GetResourceByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetAllResources(ctx context.Context, in *RequestWithFilters, opts ...grpc.CallOption) (*ArrayResponse, error) {
	out := new(ArrayResponse)
	err := c.cc.Invoke(ctx, ResourceService_GetAllResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) UpdateResource(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ResourceService_UpdateResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServiceServer is the server API for ResourceService service.
// All implementations must embed UnimplementedResourceServiceServer
// for forward compatibility
type ResourceServiceServer interface {
	InsertResource(context.Context, *Request) (*Response, error)
	GetResourceById(context.Context, *Request) (*Response, error)
	GetResourceByEmail(context.Context, *Request) (*Response, error)
	GetAllResources(context.Context, *RequestWithFilters) (*ArrayResponse, error)
	UpdateResource(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedResourceServiceServer()
}

// UnimplementedResourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServiceServer struct {
}

func (UnimplementedResourceServiceServer) InsertResource(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertResource not implemented")
}
func (UnimplementedResourceServiceServer) GetResourceById(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceById not implemented")
}
func (UnimplementedResourceServiceServer) GetResourceByEmail(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceByEmail not implemented")
}
func (UnimplementedResourceServiceServer) GetAllResources(context.Context, *RequestWithFilters) (*ArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllResources not implemented")
}
func (UnimplementedResourceServiceServer) UpdateResource(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (UnimplementedResourceServiceServer) mustEmbedUnimplementedResourceServiceServer() {}

// UnsafeResourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServiceServer will
// result in compilation errors.
type UnsafeResourceServiceServer interface {
	mustEmbedUnimplementedResourceServiceServer()
}

func RegisterResourceServiceServer(s grpc.ServiceRegistrar, srv ResourceServiceServer) {
	s.RegisterService(&ResourceService_ServiceDesc, srv)
}

func _ResourceService_InsertResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).InsertResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_InsertResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).InsertResource(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetResourceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetResourceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_GetResourceById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetResourceById(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetResourceByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetResourceByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_GetResourceByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetResourceByEmail(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetAllResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWithFilters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetAllResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_GetAllResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetAllResources(ctx, req.(*RequestWithFilters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_UpdateResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).UpdateResource(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceService_ServiceDesc is the grpc.ServiceDesc for ResourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ResourceService",
	HandlerType: (*ResourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertResource",
			Handler:    _ResourceService_InsertResource_Handler,
		},
		{
			MethodName: "GetResourceById",
			Handler:    _ResourceService_GetResourceById_Handler,
		},
		{
			MethodName: "GetResourceByEmail",
			Handler:    _ResourceService_GetResourceByEmail_Handler,
		},
		{
			MethodName: "GetAllResources",
			Handler:    _ResourceService_GetAllResources_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _ResourceService_UpdateResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}
