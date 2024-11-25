// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: whrmi.proto

package whrmi

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
	LocationService_ShowLocation_FullMethodName = "/api.whrmi.v1.LocationService/ShowLocation"
)

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	ShowLocation(ctx context.Context, in *ShowLocationRequest, opts ...grpc.CallOption) (*ShowLocationResponse, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) ShowLocation(ctx context.Context, in *ShowLocationRequest, opts ...grpc.CallOption) (*ShowLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShowLocationResponse)
	err := c.cc.Invoke(ctx, LocationService_ShowLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationServiceServer is the server API for LocationService service.
// All implementations must embed UnimplementedLocationServiceServer
// for forward compatibility.
type LocationServiceServer interface {
	ShowLocation(context.Context, *ShowLocationRequest) (*ShowLocationResponse, error)
	mustEmbedUnimplementedLocationServiceServer()
}

// UnimplementedLocationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLocationServiceServer struct{}

func (UnimplementedLocationServiceServer) ShowLocation(context.Context, *ShowLocationRequest) (*ShowLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowLocation not implemented")
}
func (UnimplementedLocationServiceServer) mustEmbedUnimplementedLocationServiceServer() {}
func (UnimplementedLocationServiceServer) testEmbeddedByValue()                         {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	// If the following call pancis, it indicates UnimplementedLocationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_ShowLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).ShowLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationService_ShowLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).ShowLocation(ctx, req.(*ShowLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.whrmi.v1.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowLocation",
			Handler:    _LocationService_ShowLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "whrmi.proto",
}

const (
	LocationKeeperService_Init_FullMethodName              = "/api.whrmi.v1.LocationKeeperService/Init"
	LocationKeeperService_AddVpnInterface_FullMethodName   = "/api.whrmi.v1.LocationKeeperService/AddVpnInterface"
	LocationKeeperService_ListVpnInterfaces_FullMethodName = "/api.whrmi.v1.LocationKeeperService/ListVpnInterfaces"
	LocationKeeperService_ExportLocations_FullMethodName   = "/api.whrmi.v1.LocationKeeperService/ExportLocations"
	LocationKeeperService_ImportLocations_FullMethodName   = "/api.whrmi.v1.LocationKeeperService/ImportLocations"
)

// LocationKeeperServiceClient is the client API for LocationKeeperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationKeeperServiceClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	AddVpnInterface(ctx context.Context, in *AddVpnInterfaceRequest, opts ...grpc.CallOption) (*AddVpnInterfaceResponse, error)
	ListVpnInterfaces(ctx context.Context, in *ListVpnInterfacesRequest, opts ...grpc.CallOption) (*ListVpnInterfacesResponse, error)
	ExportLocations(ctx context.Context, in *ExportLocationsRequest, opts ...grpc.CallOption) (*ExportLocationsResponse, error)
	ImportLocations(ctx context.Context, in *ImportLocationsRequest, opts ...grpc.CallOption) (*ImportLocationsResponse, error)
}

type locationKeeperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationKeeperServiceClient(cc grpc.ClientConnInterface) LocationKeeperServiceClient {
	return &locationKeeperServiceClient{cc}
}

func (c *locationKeeperServiceClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, LocationKeeperService_Init_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationKeeperServiceClient) AddVpnInterface(ctx context.Context, in *AddVpnInterfaceRequest, opts ...grpc.CallOption) (*AddVpnInterfaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVpnInterfaceResponse)
	err := c.cc.Invoke(ctx, LocationKeeperService_AddVpnInterface_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationKeeperServiceClient) ListVpnInterfaces(ctx context.Context, in *ListVpnInterfacesRequest, opts ...grpc.CallOption) (*ListVpnInterfacesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVpnInterfacesResponse)
	err := c.cc.Invoke(ctx, LocationKeeperService_ListVpnInterfaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationKeeperServiceClient) ExportLocations(ctx context.Context, in *ExportLocationsRequest, opts ...grpc.CallOption) (*ExportLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportLocationsResponse)
	err := c.cc.Invoke(ctx, LocationKeeperService_ExportLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationKeeperServiceClient) ImportLocations(ctx context.Context, in *ImportLocationsRequest, opts ...grpc.CallOption) (*ImportLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportLocationsResponse)
	err := c.cc.Invoke(ctx, LocationKeeperService_ImportLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationKeeperServiceServer is the server API for LocationKeeperService service.
// All implementations must embed UnimplementedLocationKeeperServiceServer
// for forward compatibility.
type LocationKeeperServiceServer interface {
	Init(context.Context, *InitRequest) (*InitResponse, error)
	AddVpnInterface(context.Context, *AddVpnInterfaceRequest) (*AddVpnInterfaceResponse, error)
	ListVpnInterfaces(context.Context, *ListVpnInterfacesRequest) (*ListVpnInterfacesResponse, error)
	ExportLocations(context.Context, *ExportLocationsRequest) (*ExportLocationsResponse, error)
	ImportLocations(context.Context, *ImportLocationsRequest) (*ImportLocationsResponse, error)
	mustEmbedUnimplementedLocationKeeperServiceServer()
}

// UnimplementedLocationKeeperServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLocationKeeperServiceServer struct{}

func (UnimplementedLocationKeeperServiceServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedLocationKeeperServiceServer) AddVpnInterface(context.Context, *AddVpnInterfaceRequest) (*AddVpnInterfaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVpnInterface not implemented")
}
func (UnimplementedLocationKeeperServiceServer) ListVpnInterfaces(context.Context, *ListVpnInterfacesRequest) (*ListVpnInterfacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVpnInterfaces not implemented")
}
func (UnimplementedLocationKeeperServiceServer) ExportLocations(context.Context, *ExportLocationsRequest) (*ExportLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportLocations not implemented")
}
func (UnimplementedLocationKeeperServiceServer) ImportLocations(context.Context, *ImportLocationsRequest) (*ImportLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportLocations not implemented")
}
func (UnimplementedLocationKeeperServiceServer) mustEmbedUnimplementedLocationKeeperServiceServer() {}
func (UnimplementedLocationKeeperServiceServer) testEmbeddedByValue()                               {}

// UnsafeLocationKeeperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationKeeperServiceServer will
// result in compilation errors.
type UnsafeLocationKeeperServiceServer interface {
	mustEmbedUnimplementedLocationKeeperServiceServer()
}

func RegisterLocationKeeperServiceServer(s grpc.ServiceRegistrar, srv LocationKeeperServiceServer) {
	// If the following call pancis, it indicates UnimplementedLocationKeeperServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LocationKeeperService_ServiceDesc, srv)
}

func _LocationKeeperService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationKeeperServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationKeeperService_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationKeeperServiceServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationKeeperService_AddVpnInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVpnInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationKeeperServiceServer).AddVpnInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationKeeperService_AddVpnInterface_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationKeeperServiceServer).AddVpnInterface(ctx, req.(*AddVpnInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationKeeperService_ListVpnInterfaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVpnInterfacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationKeeperServiceServer).ListVpnInterfaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationKeeperService_ListVpnInterfaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationKeeperServiceServer).ListVpnInterfaces(ctx, req.(*ListVpnInterfacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationKeeperService_ExportLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationKeeperServiceServer).ExportLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationKeeperService_ExportLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationKeeperServiceServer).ExportLocations(ctx, req.(*ExportLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationKeeperService_ImportLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationKeeperServiceServer).ImportLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationKeeperService_ImportLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationKeeperServiceServer).ImportLocations(ctx, req.(*ImportLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationKeeperService_ServiceDesc is the grpc.ServiceDesc for LocationKeeperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationKeeperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.whrmi.v1.LocationKeeperService",
	HandlerType: (*LocationKeeperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _LocationKeeperService_Init_Handler,
		},
		{
			MethodName: "AddVpnInterface",
			Handler:    _LocationKeeperService_AddVpnInterface_Handler,
		},
		{
			MethodName: "ListVpnInterfaces",
			Handler:    _LocationKeeperService_ListVpnInterfaces_Handler,
		},
		{
			MethodName: "ExportLocations",
			Handler:    _LocationKeeperService_ExportLocations_Handler,
		},
		{
			MethodName: "ImportLocations",
			Handler:    _LocationKeeperService_ImportLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "whrmi.proto",
}
