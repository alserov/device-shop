package validation

import (
	"github.com/alserov/device-shop/proto/gen/collection"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyUserUUID   = "user uuid can not be empty"
	emptyDeviceUUID = "device uuid can not be empty"
)

func ValidateChangeCollectionReq(req *collection.ChangeCollectionReq) error {
	if req.GetDeviceUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyDeviceUUID)
	}

	if req.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUserUUID)
	}

	return nil
}
