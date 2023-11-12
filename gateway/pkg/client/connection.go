package client

import (
	"github.com/alserov/shop/proto/gen"
	"google.golang.org/grpc"
)

func DialUsers(addr string) (pb.UsersClient, *grpc.ClientConn, error) {
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewUsersClient(cc)

	return client, cc, nil
}

func DialDevices(addr string) (pb.DevicesClient, *grpc.ClientConn, error) {
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewDevicesClient(cc)

	return client, cc, nil
}
