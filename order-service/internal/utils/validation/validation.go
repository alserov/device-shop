package validation

import (
	orderStatus "github.com/alserov/device-shop/order-service/internal/utils/status"
	"github.com/alserov/device-shop/proto/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validator struct{}

func NewValidator() Validator {
	return &validator{}
}

type Validator interface {
	ValidateCreateOrderReq(req *order.CreateOrderReq) error
	ValidateCheckOrderReq(req *order.CheckOrderReq) error
	ValidateUpdateOrderReq(req *order.UpdateOrderReq) error
	ValidateCancelOrderReq(req *order.CancelOrderReq) error
}

const (
	emptyUserUUID  = "user uuid can not be empty"
	emptyOrderUUID = "order uuid can not be empty"
	emptyDevices   = "devices list length can not be less than 1"
	invalidStatus  = "invalid order status"
)

func (*validator) ValidateCreateOrderReq(req *order.CreateOrderReq) error {
	if req.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUserUUID)
	}

	if len(req.GetOrderDevices()) < 1 {
		return status.Error(codes.InvalidArgument, emptyDevices)
	}

	return nil
}

func (*validator) ValidateCheckOrderReq(req *order.CheckOrderReq) error {
	if req.GetOrderUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyOrderUUID)
	}

	return nil
}

func (*validator) ValidateUpdateOrderReq(req *order.UpdateOrderReq) error {
	if req.GetOrderUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyOrderUUID)
	}

	if -1 == orderStatus.StatusToCode(req.GetStatus()) {
		return status.Error(codes.InvalidArgument, invalidStatus)
	}

	return nil
}

func (v *validator) ValidateCancelOrderReq(req *order.CancelOrderReq) error {
	if req.OrderUUID == "" {
		return status.Error(codes.InvalidArgument, emptyOrderUUID)
	}

	return nil
}
