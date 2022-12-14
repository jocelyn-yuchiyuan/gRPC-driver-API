// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: protos/driversapp.proto

package protos

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

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverClient interface {
	GetDrivers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Driver_GetDriversClient, error)
	GetDriver(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DriverInfo, error)
	CreateDriver(ctx context.Context, in *DriverInfo, opts ...grpc.CallOption) (*Id, error)
	UpdateDriver(ctx context.Context, in *DriverInfo, opts ...grpc.CallOption) (*Status, error)
	DeleteDriver(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error)
}

type driverClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverClient(cc grpc.ClientConnInterface) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) GetDrivers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Driver_GetDriversClient, error) {
	stream, err := c.cc.NewStream(ctx, &Driver_ServiceDesc.Streams[0], "/driversapp.Driver/GetDrivers", opts...)
	if err != nil {
		return nil, err
	}
	x := &driverGetDriversClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Driver_GetDriversClient interface {
	Recv() (*DriverInfo, error)
	grpc.ClientStream
}

type driverGetDriversClient struct {
	grpc.ClientStream
}

func (x *driverGetDriversClient) Recv() (*DriverInfo, error) {
	m := new(DriverInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *driverClient) GetDriver(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DriverInfo, error) {
	out := new(DriverInfo)
	err := c.cc.Invoke(ctx, "/driversapp.Driver/GetDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) CreateDriver(ctx context.Context, in *DriverInfo, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/driversapp.Driver/CreateDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) UpdateDriver(ctx context.Context, in *DriverInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/driversapp.Driver/UpdateDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) DeleteDriver(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/driversapp.Driver/DeleteDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
// All implementations must embed UnimplementedDriverServer
// for forward compatibility
type DriverServer interface {
	GetDrivers(*Empty, Driver_GetDriversServer) error
	GetDriver(context.Context, *Id) (*DriverInfo, error)
	CreateDriver(context.Context, *DriverInfo) (*Id, error)
	UpdateDriver(context.Context, *DriverInfo) (*Status, error)
	DeleteDriver(context.Context, *Id) (*Status, error)
	mustEmbedUnimplementedDriverServer()
}

// UnimplementedDriverServer must be embedded to have forward compatible implementations.
type UnimplementedDriverServer struct {
}

func (UnimplementedDriverServer) GetDrivers(*Empty, Driver_GetDriversServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDrivers not implemented")
}
func (UnimplementedDriverServer) GetDriver(context.Context, *Id) (*DriverInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriver not implemented")
}
func (UnimplementedDriverServer) CreateDriver(context.Context, *DriverInfo) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDriver not implemented")
}
func (UnimplementedDriverServer) UpdateDriver(context.Context, *DriverInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDriver not implemented")
}
func (UnimplementedDriverServer) DeleteDriver(context.Context, *Id) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDriver not implemented")
}
func (UnimplementedDriverServer) mustEmbedUnimplementedDriverServer() {}

// UnsafeDriverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServer will
// result in compilation errors.
type UnsafeDriverServer interface {
	mustEmbedUnimplementedDriverServer()
}

func RegisterDriverServer(s grpc.ServiceRegistrar, srv DriverServer) {
	s.RegisterService(&Driver_ServiceDesc, srv)
}

func _Driver_GetDrivers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DriverServer).GetDrivers(m, &driverGetDriversServer{stream})
}

type Driver_GetDriversServer interface {
	Send(*DriverInfo) error
	grpc.ServerStream
}

type driverGetDriversServer struct {
	grpc.ServerStream
}

func (x *driverGetDriversServer) Send(m *DriverInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Driver_GetDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GetDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driversapp.Driver/GetDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GetDriver(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_CreateDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).CreateDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driversapp.Driver/CreateDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).CreateDriver(ctx, req.(*DriverInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_UpdateDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).UpdateDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driversapp.Driver/UpdateDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).UpdateDriver(ctx, req.(*DriverInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_DeleteDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).DeleteDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driversapp.Driver/DeleteDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).DeleteDriver(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Driver_ServiceDesc is the grpc.ServiceDesc for Driver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Driver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "driversapp.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDriver",
			Handler:    _Driver_GetDriver_Handler,
		},
		{
			MethodName: "CreateDriver",
			Handler:    _Driver_CreateDriver_Handler,
		},
		{
			MethodName: "UpdateDriver",
			Handler:    _Driver_UpdateDriver_Handler,
		},
		{
			MethodName: "DeleteDriver",
			Handler:    _Driver_DeleteDriver_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDrivers",
			Handler:       _Driver_GetDrivers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/driversapp.proto",
}
