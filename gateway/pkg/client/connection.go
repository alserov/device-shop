package client

import (
	"github.com/alserov/device-shop/proto/gen"
	"google.golang.org/grpc"
)

func dial(addr string) (*grpc.ClientConn, error) {
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func DialUser(addr string) (pb.UsersClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewUsersClient(cc)

	return client, cc, nil
}

func DialDevice(addr string) (pb.DevicesClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewDevicesClient(cc)

	return client, cc, nil
}

func DialOrder(addr string) (pb.OrdersClient, *grpc.ClientConn, error) {
	cc, err := dial(addr)
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewOrdersClient(cc)

	return client, cc, nil
}
