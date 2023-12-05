package validation

import (
	orderStatus "github.com/alserov/device-shop/order-service/internal/utils/status"
	"github.com/alserov/device-shop/proto/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

const (
	emptyUserUUID  = "user uuid can not be empty"
	emptyOrderUUID = "order uuid can not be empty"
	emptyDevices   = "devices list length can not be less than 1"
	invalidStatus  = "invalid order status"
)

func (*Validator) ValidateCreateOrderReq(req *order.CreateOrderReq) error {
	if req.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUserUUID)
	}

	if len(req.GetDevices()) < 1 {
		return status.Error(codes.InvalidArgument, emptyDevices)
	}

	return nil
}

func (*Validator) ValidateCheckOrderReq(req *order.CheckOrderReq) error {
	if req.GetOrderUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyOrderUUID)
	}

	return nil
}

func (*Validator) ValidateUpdateOrderReq(req *order.UpdateOrderReq) error {
	if req.GetOrderUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyOrderUUID)
	}

	if -1 == orderStatus.StatusToCode(req.GetStatus()) {
		return status.Error(codes.InvalidArgument, invalidStatus)
	}

	return nil
}
