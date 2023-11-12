// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: proto/device.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevicesClient interface {
	CreateDevice(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDevice(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateDevice(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllDevices(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*DevicesRes, error)
	GetDevicesByTitle(ctx context.Context, in *GetByTitleReq, opts ...grpc.CallOption) (*DevicesRes, error)
	GetDevicesByManufacturer(ctx context.Context, in *GetByManufacturer, opts ...grpc.CallOption) (*DevicesRes, error)
	GetDevicesByPrice(ctx context.Context, in *GetByPrice, opts ...grpc.CallOption) (*DevicesRes, error)
	GetDeviceByUUID(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*Device, error)
}

type devicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) CreateDevice(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/device.Devices/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) DeleteDevice(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/device.Devices/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) UpdateDevice(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/device.Devices/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetAllDevices(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*DevicesRes, error) {
	out := new(DevicesRes)
	err := c.cc.Invoke(ctx, "/device.Devices/GetAllDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDevicesByTitle(ctx context.Context, in *GetByTitleReq, opts ...grpc.CallOption) (*DevicesRes, error) {
	out := new(DevicesRes)
	err := c.cc.Invoke(ctx, "/device.Devices/GetDevicesByTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDevicesByManufacturer(ctx context.Context, in *GetByManufacturer, opts ...grpc.CallOption) (*DevicesRes, error) {
	out := new(DevicesRes)
	err := c.cc.Invoke(ctx, "/device.Devices/GetDevicesByManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDevicesByPrice(ctx context.Context, in *GetByPrice, opts ...grpc.CallOption) (*DevicesRes, error) {
	out := new(DevicesRes)
	err := c.cc.Invoke(ctx, "/device.Devices/GetDevicesByPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDeviceByUUID(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/device.Devices/GetDeviceByUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
// All implementations must embed UnimplementedDevicesServer
// for forward compatibility
type DevicesServer interface {
	CreateDevice(context.Context, *CreateReq) (*emptypb.Empty, error)
	DeleteDevice(context.Context, *DeleteReq) (*emptypb.Empty, error)
	UpdateDevice(context.Context, *UpdateReq) (*emptypb.Empty, error)
	GetAllDevices(context.Context, *GetAllReq) (*DevicesRes, error)
	GetDevicesByTitle(context.Context, *GetByTitleReq) (*DevicesRes, error)
	GetDevicesByManufacturer(context.Context, *GetByManufacturer) (*DevicesRes, error)
	GetDevicesByPrice(context.Context, *GetByPrice) (*DevicesRes, error)
	GetDeviceByUUID(context.Context, *UUIDReq) (*Device, error)
}

// UnimplementedDevicesServer must be embedded to have forward compatible implementations.
type UnimplementedDevicesServer struct {
}

func (UnimplementedDevicesServer) CreateDevice(context.Context, *CreateReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedDevicesServer) DeleteDevice(context.Context, *DeleteReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedDevicesServer) UpdateDevice(context.Context, *UpdateReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedDevicesServer) GetAllDevices(context.Context, *GetAllReq) (*DevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDevices not implemented")
}
func (UnimplementedDevicesServer) GetDevicesByTitle(context.Context, *GetByTitleReq) (*DevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevicesByTitle not implemented")
}
func (UnimplementedDevicesServer) GetDevicesByManufacturer(context.Context, *GetByManufacturer) (*DevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevicesByManufacturer not implemented")
}
func (UnimplementedDevicesServer) GetDevicesByPrice(context.Context, *GetByPrice) (*DevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevicesByPrice not implemented")
}
func (UnimplementedDevicesServer) GetDeviceByUUID(context.Context, *UUIDReq) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceByUUID not implemented")
}
func (UnimplementedDevicesServer) mustEmbedUnimplementedDevicesServer() {}

// UnsafeDevicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevicesServer will
// result in compilation errors.
type UnsafeDevicesServer interface {
	mustEmbedUnimplementedDevicesServer()
}

func RegisterDevicesServer(s grpc.ServiceRegistrar, srv DevicesServer) {
	s.RegisterService(&Devices_ServiceDesc, srv)
}

func _Devices_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).CreateDevice(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).DeleteDevice(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).UpdateDevice(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetAllDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetAllDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/GetAllDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetAllDevices(ctx, req.(*GetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDevicesByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTitleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDevicesByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/GetDevicesByTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDevicesByTitle(ctx, req.(*GetByTitleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDevicesByManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByManufacturer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDevicesByManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/GetDevicesByManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDevicesByManufacturer(ctx, req.(*GetByManufacturer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDevicesByPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDevicesByPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/GetDevicesByPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDevicesByPrice(ctx, req.(*GetByPrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDeviceByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDeviceByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.Devices/GetDeviceByUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDeviceByUUID(ctx, req.(*UUIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Devices_ServiceDesc is the grpc.ServiceDesc for Devices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Devices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDevice",
			Handler:    _Devices_CreateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _Devices_DeleteDevice_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _Devices_UpdateDevice_Handler,
		},
		{
			MethodName: "GetAllDevices",
			Handler:    _Devices_GetAllDevices_Handler,
		},
		{
			MethodName: "GetDevicesByTitle",
			Handler:    _Devices_GetDevicesByTitle_Handler,
		},
		{
			MethodName: "GetDevicesByManufacturer",
			Handler:    _Devices_GetDevicesByManufacturer_Handler,
		},
		{
			MethodName: "GetDevicesByPrice",
			Handler:    _Devices_GetDevicesByPrice_Handler,
		},
		{
			MethodName: "GetDeviceByUUID",
			Handler:    _Devices_GetDeviceByUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/device.proto",
}
