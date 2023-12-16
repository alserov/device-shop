package services

import (
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/alserov/device-shop/proto/gen/order"
	"github.com/alserov/device-shop/proto/gen/user"
	"google.golang.org/grpc"
)

func NewOrderClient(addr string) (order.OrdersClient, *grpc.ClientConn) {
	cl, cc, err := client.DialOrder(addr)
	if err != nil {
		panic("failed to dial order service: " + err.Error())
	}
	return cl, cc
}

func NewDeviceClient(addr string) (device.DevicesClient, *grpc.ClientConn) {
	cl, cc, err := client.DialDevice(addr)
	if err != nil {
		panic("failed to dial order service: " + err.Error())
	}
	return cl, cc
}

func NewUserClient(addr string) (user.UsersClient, *grpc.ClientConn) {
	cl, cc, err := client.DialUser(addr)
	if err != nil {
		panic("failed to dial order service: " + err.Error())
	}
	return cl, cc
}

func NewCollectionClient(addr string) (collection.CollectionsClient, *grpc.ClientConn) {
	cl, cc, err := client.DialCollection(addr)
	if err != nil {
		panic("failed to dial order service: " + err.Error())
	}
	return cl, cc
}
