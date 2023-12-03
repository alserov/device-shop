package client

import (
	"github.com/alserov/device-shop/proto/gen/admin"
	"github.com/alserov/device-shop/proto/gen/auth"
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/alserov/device-shop/proto/gen/order"
	"github.com/alserov/device-shop/proto/gen/user"
	"google.golang.org/grpc"
)

func dial(addr string) (*grpc.ClientConn, error) {
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func DialUser(addr string) (user.UsersClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := user.NewUsersClient(cc)

	return client, cc, nil
}

func DialAdmin(addr string) (admin.AdminClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := admin.NewAdminClient(cc)

	return client, cc, err
}

func DialCollection(addr string) (collection.CollectionsClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := collection.NewCollectionsClient(cc)

	return client, cc, nil
}

func DialDevice(addr string) (device.DevicesClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := device.NewDevicesClient(cc)

	return client, cc, nil
}

func DialOrder(addr string) (order.OrdersClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := order.NewOrdersClient(cc)

	return client, cc, nil
}

func DialAuth(addr string) (auth.AuthClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := auth.NewAuthClient(cc)

	return client, cc, err
}
