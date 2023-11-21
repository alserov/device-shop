// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: proto/user.proto

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

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRes, error)
	AddToFavourite(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveFromFavourite(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFavourite(ctx context.Context, in *GetCollectionReq, opts ...grpc.CallOption) (*GetCollectionRes, error)
	AddToCart(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveFromCart(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCart(ctx context.Context, in *GetCollectionReq, opts ...grpc.CallOption) (*GetCollectionRes, error)
	TopUpBalance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*BalanceRes, error)
	DebitBalance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*BalanceRes, error)
	RemoveDeviceFromCollections(ctx context.Context, in *RemoveDeletedDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRes, error) {
	out := new(GetUserInfoRes)
	err := c.cc.Invoke(ctx, "/user.Users/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AddToFavourite(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.Users/AddToFavourite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RemoveFromFavourite(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.Users/RemoveFromFavourite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetFavourite(ctx context.Context, in *GetCollectionReq, opts ...grpc.CallOption) (*GetCollectionRes, error) {
	out := new(GetCollectionRes)
	err := c.cc.Invoke(ctx, "/user.Users/GetFavourite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AddToCart(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.Users/AddToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RemoveFromCart(ctx context.Context, in *ChangeCollectionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.Users/RemoveFromCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetCart(ctx context.Context, in *GetCollectionReq, opts ...grpc.CallOption) (*GetCollectionRes, error) {
	out := new(GetCollectionRes)
	err := c.cc.Invoke(ctx, "/user.Users/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) TopUpBalance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*BalanceRes, error) {
	out := new(BalanceRes)
	err := c.cc.Invoke(ctx, "/user.Users/TopUpBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DebitBalance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*BalanceRes, error) {
	out := new(BalanceRes)
	err := c.cc.Invoke(ctx, "/user.Users/DebitBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RemoveDeviceFromCollections(ctx context.Context, in *RemoveDeletedDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.Users/RemoveDeviceFromCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRes, error)
	AddToFavourite(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error)
	RemoveFromFavourite(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error)
	GetFavourite(context.Context, *GetCollectionReq) (*GetCollectionRes, error)
	AddToCart(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error)
	RemoveFromCart(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error)
	GetCart(context.Context, *GetCollectionReq) (*GetCollectionRes, error)
	TopUpBalance(context.Context, *BalanceReq) (*BalanceRes, error)
	DebitBalance(context.Context, *BalanceReq) (*BalanceRes, error)
	RemoveDeviceFromCollections(context.Context, *RemoveDeletedDeviceReq) (*emptypb.Empty, error)
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUsersServer) AddToFavourite(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToFavourite not implemented")
}
func (UnimplementedUsersServer) RemoveFromFavourite(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromFavourite not implemented")
}
func (UnimplementedUsersServer) GetFavourite(context.Context, *GetCollectionReq) (*GetCollectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavourite not implemented")
}
func (UnimplementedUsersServer) AddToCart(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedUsersServer) RemoveFromCart(context.Context, *ChangeCollectionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCart not implemented")
}
func (UnimplementedUsersServer) GetCart(context.Context, *GetCollectionReq) (*GetCollectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedUsersServer) TopUpBalance(context.Context, *BalanceReq) (*BalanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopUpBalance not implemented")
}
func (UnimplementedUsersServer) DebitBalance(context.Context, *BalanceReq) (*BalanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebitBalance not implemented")
}
func (UnimplementedUsersServer) RemoveDeviceFromCollections(context.Context, *RemoveDeletedDeviceReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDeviceFromCollections not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AddToFavourite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AddToFavourite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/AddToFavourite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AddToFavourite(ctx, req.(*ChangeCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RemoveFromFavourite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RemoveFromFavourite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/RemoveFromFavourite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RemoveFromFavourite(ctx, req.(*ChangeCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetFavourite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetFavourite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetFavourite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetFavourite(ctx, req.(*GetCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/AddToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AddToCart(ctx, req.(*ChangeCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RemoveFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RemoveFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/RemoveFromCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RemoveFromCart(ctx, req.(*ChangeCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetCart(ctx, req.(*GetCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_TopUpBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).TopUpBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/TopUpBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).TopUpBalance(ctx, req.(*BalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DebitBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DebitBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/DebitBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DebitBalance(ctx, req.(*BalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RemoveDeviceFromCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeletedDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RemoveDeviceFromCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/RemoveDeviceFromCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RemoveDeviceFromCollections(ctx, req.(*RemoveDeletedDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Users_GetUserInfo_Handler,
		},
		{
			MethodName: "AddToFavourite",
			Handler:    _Users_AddToFavourite_Handler,
		},
		{
			MethodName: "RemoveFromFavourite",
			Handler:    _Users_RemoveFromFavourite_Handler,
		},
		{
			MethodName: "GetFavourite",
			Handler:    _Users_GetFavourite_Handler,
		},
		{
			MethodName: "AddToCart",
			Handler:    _Users_AddToCart_Handler,
		},
		{
			MethodName: "RemoveFromCart",
			Handler:    _Users_RemoveFromCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _Users_GetCart_Handler,
		},
		{
			MethodName: "TopUpBalance",
			Handler:    _Users_TopUpBalance_Handler,
		},
		{
			MethodName: "DebitBalance",
			Handler:    _Users_DebitBalance_Handler,
		},
		{
			MethodName: "RemoveDeviceFromCollections",
			Handler:    _Users_RemoveDeviceFromCollections_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
