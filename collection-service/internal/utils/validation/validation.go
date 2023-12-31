package validation

import (
	coll "github.com/alserov/device-shop/proto/gen/collection"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyUserUUID   = "user uuid can not be empty"
	emptyDeviceUUID = "device uuid can not be empty"
)

type Validator struct {
	Collection Collection
}

func NewValidator() *Validator {
	return &Validator{
		Collection: &collection{},
	}
}

type collection struct{}
type Collection interface {
	ValidateChangeCollectionReq(req *coll.ChangeCollectionReq) error
}

func (*collection) ValidateChangeCollectionReq(req *coll.ChangeCollectionReq) error {
	if req.GetDeviceUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyDeviceUUID)
	}

	if req.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUserUUID)
	}

	return nil
}
