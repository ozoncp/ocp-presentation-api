// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_presentation_api

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

// PresentationAPIClient is the client API for PresentationAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PresentationAPIClient interface {
	// Creates a new presentation
	CreatePresentationV1(ctx context.Context, in *CreatePresentationV1Request, opts ...grpc.CallOption) (*CreatePresentationV1Response, error)
	// Returns a presentation by id
	DescribePresentationV1(ctx context.Context, in *DescribePresentationV1Request, opts ...grpc.CallOption) (*DescribePresentationV1Response, error)
	// Returns a list of presentations
	ListPresentationsV1(ctx context.Context, in *ListPresentationsV1Request, opts ...grpc.CallOption) (*ListPresentationsV1Response, error)
	// Removes a presentation by id
	RemovePresentationV1(ctx context.Context, in *RemovePresentationV1Request, opts ...grpc.CallOption) (*RemovePresentationV1Response, error)
}

type presentationAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPresentationAPIClient(cc grpc.ClientConnInterface) PresentationAPIClient {
	return &presentationAPIClient{cc}
}

func (c *presentationAPIClient) CreatePresentationV1(ctx context.Context, in *CreatePresentationV1Request, opts ...grpc.CallOption) (*CreatePresentationV1Response, error) {
	out := new(CreatePresentationV1Response)
	err := c.cc.Invoke(ctx, "/ocp.presentation.api.PresentationAPI/CreatePresentationV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationAPIClient) DescribePresentationV1(ctx context.Context, in *DescribePresentationV1Request, opts ...grpc.CallOption) (*DescribePresentationV1Response, error) {
	out := new(DescribePresentationV1Response)
	err := c.cc.Invoke(ctx, "/ocp.presentation.api.PresentationAPI/DescribePresentationV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationAPIClient) ListPresentationsV1(ctx context.Context, in *ListPresentationsV1Request, opts ...grpc.CallOption) (*ListPresentationsV1Response, error) {
	out := new(ListPresentationsV1Response)
	err := c.cc.Invoke(ctx, "/ocp.presentation.api.PresentationAPI/ListPresentationsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationAPIClient) RemovePresentationV1(ctx context.Context, in *RemovePresentationV1Request, opts ...grpc.CallOption) (*RemovePresentationV1Response, error) {
	out := new(RemovePresentationV1Response)
	err := c.cc.Invoke(ctx, "/ocp.presentation.api.PresentationAPI/RemovePresentationV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresentationAPIServer is the server API for PresentationAPI service.
// All implementations must embed UnimplementedPresentationAPIServer
// for forward compatibility
type PresentationAPIServer interface {
	// Creates a new presentation
	CreatePresentationV1(context.Context, *CreatePresentationV1Request) (*CreatePresentationV1Response, error)
	// Returns a presentation by id
	DescribePresentationV1(context.Context, *DescribePresentationV1Request) (*DescribePresentationV1Response, error)
	// Returns a list of presentations
	ListPresentationsV1(context.Context, *ListPresentationsV1Request) (*ListPresentationsV1Response, error)
	// Removes a presentation by id
	RemovePresentationV1(context.Context, *RemovePresentationV1Request) (*RemovePresentationV1Response, error)
	mustEmbedUnimplementedPresentationAPIServer()
}

// UnimplementedPresentationAPIServer must be embedded to have forward compatible implementations.
type UnimplementedPresentationAPIServer struct {
}

func (UnimplementedPresentationAPIServer) CreatePresentationV1(context.Context, *CreatePresentationV1Request) (*CreatePresentationV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePresentationV1 not implemented")
}
func (UnimplementedPresentationAPIServer) DescribePresentationV1(context.Context, *DescribePresentationV1Request) (*DescribePresentationV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePresentationV1 not implemented")
}
func (UnimplementedPresentationAPIServer) ListPresentationsV1(context.Context, *ListPresentationsV1Request) (*ListPresentationsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPresentationsV1 not implemented")
}
func (UnimplementedPresentationAPIServer) RemovePresentationV1(context.Context, *RemovePresentationV1Request) (*RemovePresentationV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePresentationV1 not implemented")
}
func (UnimplementedPresentationAPIServer) mustEmbedUnimplementedPresentationAPIServer() {}

// UnsafePresentationAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresentationAPIServer will
// result in compilation errors.
type UnsafePresentationAPIServer interface {
	mustEmbedUnimplementedPresentationAPIServer()
}

func RegisterPresentationAPIServer(s grpc.ServiceRegistrar, srv PresentationAPIServer) {
	s.RegisterService(&PresentationAPI_ServiceDesc, srv)
}

func _PresentationAPI_CreatePresentationV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePresentationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresentationAPIServer).CreatePresentationV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.presentation.api.PresentationAPI/CreatePresentationV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresentationAPIServer).CreatePresentationV1(ctx, req.(*CreatePresentationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresentationAPI_DescribePresentationV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePresentationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresentationAPIServer).DescribePresentationV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.presentation.api.PresentationAPI/DescribePresentationV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresentationAPIServer).DescribePresentationV1(ctx, req.(*DescribePresentationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresentationAPI_ListPresentationsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPresentationsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresentationAPIServer).ListPresentationsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.presentation.api.PresentationAPI/ListPresentationsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresentationAPIServer).ListPresentationsV1(ctx, req.(*ListPresentationsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresentationAPI_RemovePresentationV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePresentationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresentationAPIServer).RemovePresentationV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.presentation.api.PresentationAPI/RemovePresentationV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresentationAPIServer).RemovePresentationV1(ctx, req.(*RemovePresentationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PresentationAPI_ServiceDesc is the grpc.ServiceDesc for PresentationAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PresentationAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.presentation.api.PresentationAPI",
	HandlerType: (*PresentationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePresentationV1",
			Handler:    _PresentationAPI_CreatePresentationV1_Handler,
		},
		{
			MethodName: "DescribePresentationV1",
			Handler:    _PresentationAPI_DescribePresentationV1_Handler,
		},
		{
			MethodName: "ListPresentationsV1",
			Handler:    _PresentationAPI_ListPresentationsV1_Handler,
		},
		{
			MethodName: "RemovePresentationV1",
			Handler:    _PresentationAPI_RemovePresentationV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-presentation-api/ocp-presentation-api.proto",
}
