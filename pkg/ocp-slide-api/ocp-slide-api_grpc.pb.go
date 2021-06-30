// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_slide_api

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

// SlideAPIClient is the client API for SlideAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlideAPIClient interface {
	// Creates a new slide
	CreateSlideV1(ctx context.Context, in *CreateSlideV1Request, opts ...grpc.CallOption) (*CreateSlideV1Response, error)
	// Creates new slides
	MultiCreateSlidesV1(ctx context.Context, in *MultiCreateSlidesV1Request, opts ...grpc.CallOption) (*MultiCreateSlidesV1Response, error)
	// Updates a slide
	UpdateSlideV1(ctx context.Context, in *UpdateSlideV1Request, opts ...grpc.CallOption) (*UpdateSlideV1Response, error)
	// Returns a slide by id
	DescribeSlideV1(ctx context.Context, in *DescribeSlideV1Request, opts ...grpc.CallOption) (*DescribeSlideV1Response, error)
	// Returns a list of slides
	ListSlidesV1(ctx context.Context, in *ListSlidesV1Request, opts ...grpc.CallOption) (*ListSlidesV1Response, error)
	// Removes a slide by id
	RemoveSlideV1(ctx context.Context, in *RemoveSlideV1Request, opts ...grpc.CallOption) (*RemoveSlideV1Response, error)
}

type slideAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSlideAPIClient(cc grpc.ClientConnInterface) SlideAPIClient {
	return &slideAPIClient{cc}
}

func (c *slideAPIClient) CreateSlideV1(ctx context.Context, in *CreateSlideV1Request, opts ...grpc.CallOption) (*CreateSlideV1Response, error) {
	out := new(CreateSlideV1Response)
	err := c.cc.Invoke(ctx, "/ocp.slide.api.SlideAPI/CreateSlideV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideAPIClient) MultiCreateSlidesV1(ctx context.Context, in *MultiCreateSlidesV1Request, opts ...grpc.CallOption) (*MultiCreateSlidesV1Response, error) {
	out := new(MultiCreateSlidesV1Response)
	err := c.cc.Invoke(ctx, "/ocp.slide.api.SlideAPI/MultiCreateSlidesV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideAPIClient) UpdateSlideV1(ctx context.Context, in *UpdateSlideV1Request, opts ...grpc.CallOption) (*UpdateSlideV1Response, error) {
	out := new(UpdateSlideV1Response)
	err := c.cc.Invoke(ctx, "/ocp.slide.api.SlideAPI/UpdateSlideV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideAPIClient) DescribeSlideV1(ctx context.Context, in *DescribeSlideV1Request, opts ...grpc.CallOption) (*DescribeSlideV1Response, error) {
	out := new(DescribeSlideV1Response)
	err := c.cc.Invoke(ctx, "/ocp.slide.api.SlideAPI/DescribeSlideV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideAPIClient) ListSlidesV1(ctx context.Context, in *ListSlidesV1Request, opts ...grpc.CallOption) (*ListSlidesV1Response, error) {
	out := new(ListSlidesV1Response)
	err := c.cc.Invoke(ctx, "/ocp.slide.api.SlideAPI/ListSlidesV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideAPIClient) RemoveSlideV1(ctx context.Context, in *RemoveSlideV1Request, opts ...grpc.CallOption) (*RemoveSlideV1Response, error) {
	out := new(RemoveSlideV1Response)
	err := c.cc.Invoke(ctx, "/ocp.slide.api.SlideAPI/RemoveSlideV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlideAPIServer is the server API for SlideAPI service.
// All implementations must embed UnimplementedSlideAPIServer
// for forward compatibility
type SlideAPIServer interface {
	// Creates a new slide
	CreateSlideV1(context.Context, *CreateSlideV1Request) (*CreateSlideV1Response, error)
	// Creates new slides
	MultiCreateSlidesV1(context.Context, *MultiCreateSlidesV1Request) (*MultiCreateSlidesV1Response, error)
	// Updates a slide
	UpdateSlideV1(context.Context, *UpdateSlideV1Request) (*UpdateSlideV1Response, error)
	// Returns a slide by id
	DescribeSlideV1(context.Context, *DescribeSlideV1Request) (*DescribeSlideV1Response, error)
	// Returns a list of slides
	ListSlidesV1(context.Context, *ListSlidesV1Request) (*ListSlidesV1Response, error)
	// Removes a slide by id
	RemoveSlideV1(context.Context, *RemoveSlideV1Request) (*RemoveSlideV1Response, error)
	mustEmbedUnimplementedSlideAPIServer()
}

// UnimplementedSlideAPIServer must be embedded to have forward compatible implementations.
type UnimplementedSlideAPIServer struct {
}

func (UnimplementedSlideAPIServer) CreateSlideV1(context.Context, *CreateSlideV1Request) (*CreateSlideV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlideV1 not implemented")
}
func (UnimplementedSlideAPIServer) MultiCreateSlidesV1(context.Context, *MultiCreateSlidesV1Request) (*MultiCreateSlidesV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateSlidesV1 not implemented")
}
func (UnimplementedSlideAPIServer) UpdateSlideV1(context.Context, *UpdateSlideV1Request) (*UpdateSlideV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlideV1 not implemented")
}
func (UnimplementedSlideAPIServer) DescribeSlideV1(context.Context, *DescribeSlideV1Request) (*DescribeSlideV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSlideV1 not implemented")
}
func (UnimplementedSlideAPIServer) ListSlidesV1(context.Context, *ListSlidesV1Request) (*ListSlidesV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSlidesV1 not implemented")
}
func (UnimplementedSlideAPIServer) RemoveSlideV1(context.Context, *RemoveSlideV1Request) (*RemoveSlideV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSlideV1 not implemented")
}
func (UnimplementedSlideAPIServer) mustEmbedUnimplementedSlideAPIServer() {}

// UnsafeSlideAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlideAPIServer will
// result in compilation errors.
type UnsafeSlideAPIServer interface {
	mustEmbedUnimplementedSlideAPIServer()
}

func RegisterSlideAPIServer(s grpc.ServiceRegistrar, srv SlideAPIServer) {
	s.RegisterService(&SlideAPI_ServiceDesc, srv)
}

func _SlideAPI_CreateSlideV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSlideV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlideAPIServer).CreateSlideV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.slide.api.SlideAPI/CreateSlideV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlideAPIServer).CreateSlideV1(ctx, req.(*CreateSlideV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlideAPI_MultiCreateSlidesV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateSlidesV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlideAPIServer).MultiCreateSlidesV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.slide.api.SlideAPI/MultiCreateSlidesV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlideAPIServer).MultiCreateSlidesV1(ctx, req.(*MultiCreateSlidesV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlideAPI_UpdateSlideV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSlideV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlideAPIServer).UpdateSlideV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.slide.api.SlideAPI/UpdateSlideV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlideAPIServer).UpdateSlideV1(ctx, req.(*UpdateSlideV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlideAPI_DescribeSlideV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSlideV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlideAPIServer).DescribeSlideV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.slide.api.SlideAPI/DescribeSlideV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlideAPIServer).DescribeSlideV1(ctx, req.(*DescribeSlideV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlideAPI_ListSlidesV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSlidesV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlideAPIServer).ListSlidesV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.slide.api.SlideAPI/ListSlidesV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlideAPIServer).ListSlidesV1(ctx, req.(*ListSlidesV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlideAPI_RemoveSlideV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSlideV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlideAPIServer).RemoveSlideV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.slide.api.SlideAPI/RemoveSlideV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlideAPIServer).RemoveSlideV1(ctx, req.(*RemoveSlideV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SlideAPI_ServiceDesc is the grpc.ServiceDesc for SlideAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SlideAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.slide.api.SlideAPI",
	HandlerType: (*SlideAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSlideV1",
			Handler:    _SlideAPI_CreateSlideV1_Handler,
		},
		{
			MethodName: "MultiCreateSlidesV1",
			Handler:    _SlideAPI_MultiCreateSlidesV1_Handler,
		},
		{
			MethodName: "UpdateSlideV1",
			Handler:    _SlideAPI_UpdateSlideV1_Handler,
		},
		{
			MethodName: "DescribeSlideV1",
			Handler:    _SlideAPI_DescribeSlideV1_Handler,
		},
		{
			MethodName: "ListSlidesV1",
			Handler:    _SlideAPI_ListSlidesV1_Handler,
		},
		{
			MethodName: "RemoveSlideV1",
			Handler:    _SlideAPI_RemoveSlideV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-slide-api/ocp-slide-api.proto",
}
