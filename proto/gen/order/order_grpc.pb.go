// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: order/order.proto

package order

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

// OrdersClient is the client API for Orders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error)
	CheckOrder(ctx context.Context, in *CheckOrderReq, opts ...grpc.CallOption) (*CheckOrderRes, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderRes, error)
}

type ordersClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersClient(cc grpc.ClientConnInterface) OrdersClient {
	return &ordersClient{cc}
}

func (c *ordersClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error) {
	out := new(CreateOrderRes)
	err := c.cc.Invoke(ctx, "/order.Orders/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) CheckOrder(ctx context.Context, in *CheckOrderReq, opts ...grpc.CallOption) (*CheckOrderRes, error) {
	out := new(CheckOrderRes)
	err := c.cc.Invoke(ctx, "/order.Orders/CheckOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderRes, error) {
	out := new(UpdateOrderRes)
	err := c.cc.Invoke(ctx, "/order.Orders/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServer is the server API for Orders service.
// All implementations must embed UnimplementedOrdersServer
// for forward compatibility
type OrdersServer interface {
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRes, error)
	CheckOrder(context.Context, *CheckOrderReq) (*CheckOrderRes, error)
	UpdateOrder(context.Context, *UpdateOrderReq) (*UpdateOrderRes, error)
	mustEmbedUnimplementedOrdersServer()
}

// UnimplementedOrdersServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServer struct {
}

func (UnimplementedOrdersServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrdersServer) CheckOrder(context.Context, *CheckOrderReq) (*CheckOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOrder not implemented")
}
func (UnimplementedOrdersServer) UpdateOrder(context.Context, *UpdateOrderReq) (*UpdateOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrdersServer) mustEmbedUnimplementedOrdersServer() {}

// UnsafeOrdersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServer will
// result in compilation errors.
type UnsafeOrdersServer interface {
	mustEmbedUnimplementedOrdersServer()
}

func RegisterOrdersServer(s grpc.ServiceRegistrar, srv OrdersServer) {
	s.RegisterService(&Orders_ServiceDesc, srv)
}

func _Orders_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Orders/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_CheckOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).CheckOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Orders/CheckOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).CheckOrder(ctx, req.(*CheckOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Orders/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).UpdateOrder(ctx, req.(*UpdateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Orders_ServiceDesc is the grpc.ServiceDesc for Orders service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orders_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Orders",
	HandlerType: (*OrdersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Orders_CreateOrder_Handler,
		},
		{
			MethodName: "CheckOrder",
			Handler:    _Orders_CheckOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Orders_UpdateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/order.proto",
}
