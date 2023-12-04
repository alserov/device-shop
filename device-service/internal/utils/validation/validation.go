package validation

import (
	"github.com/alserov/device-shop/proto/gen/device"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	invalidPrice    = "price can not be less or equal to 0"
	invalidMaxPrice = "max price can not be less or equal to min price"
	invalidAmount   = "amount can not be less than 1"
	emptyTitle      = "title can not be empty"
	emptyDesc       = "description can not be empty"
	emptyManu       = "manufacturer can not be empty"
	emptyUUID       = "uuid can not be empty"
)

type Validator struct {
	Admin  AdminValidator
	Device DeviceValidator
}

func NewValidator() *Validator {
	return &Validator{
		Admin:  &adminValidator{},
		Device: &deviceValidator{},
	}
}

type adminValidator struct{}
type AdminValidator interface {
	ValidateCreateDeviceReq(req *device.CreateDeviceReq) error
	ValidateDeleteDeviceReq(req *device.DeleteDeviceReq) error
	ValidateUpdateDeviceReq(req *device.UpdateDeviceReq) error
}

type deviceValidator struct{}
type DeviceValidator interface {
	ValidateGetDeviceByTitleReq(req *device.GetDeviceByTitleReq) error
	ValidateGetDevicesByManufacturerReq(req *device.GetByManufacturer) error
	ValidateGetDevicesByPrice(req *device.GetByPrice) error
	ValidateGetDeviceByUUID(req *device.GetDeviceByUUIDReq) error
}

func (*adminValidator) ValidateCreateDeviceReq(req *device.CreateDeviceReq) error {
	if req.GetPrice() <= 0 {
		return status.Error(codes.InvalidArgument, invalidPrice)
	}

	if req.GetAmount() < 1 {
		return status.Error(codes.InvalidArgument, invalidAmount)
	}

	if req.GetTitle() == "" {
		return status.Error(codes.InvalidArgument, emptyTitle)
	}

	if req.GetDescription() == "" {
		return status.Error(codes.InvalidArgument, emptyDesc)
	}

	if req.GetManufacturer() == "" {
		return status.Error(codes.InvalidArgument, emptyManu)
	}

	return nil
}

func (*adminValidator) ValidateDeleteDeviceReq(req *device.DeleteDeviceReq) error {
	if req.GetUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	return nil
}

func (*adminValidator) ValidateUpdateDeviceReq(req *device.UpdateDeviceReq) error {
	if req.GetPrice() <= 0 {
		return status.Error(codes.InvalidArgument, invalidPrice)
	}

	if req.GetTitle() == "" {
		return status.Error(codes.InvalidArgument, emptyTitle)
	}

	if req.GetUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	if req.GetDescription() == "" {
		return status.Error(codes.InvalidArgument, emptyDesc)
	}

	return nil
}

func (*deviceValidator) ValidateGetDeviceByTitleReq(req *device.GetDeviceByTitleReq) error {
	if req.GetTitle() == "" {
		return status.Error(codes.Internal, emptyTitle)
	}
	return nil
}

func (*deviceValidator) ValidateGetDevicesByManufacturerReq(req *device.GetByManufacturer) error {
	if req.GetManufacturer() == "" {
		return status.Error(codes.InvalidArgument, emptyManu)
	}
	return nil
}

func (*deviceValidator) ValidateGetDevicesByPrice(req *device.GetByPrice) error {
	if req.GetMin() < 0 {
		return status.Error(codes.InvalidArgument, invalidPrice)
	}

	if req.Max <= req.Min {
		return status.Error(codes.InvalidArgument, invalidMaxPrice)
	}
	return nil
}

func (*deviceValidator) ValidateGetDeviceByUUID(req *device.GetDeviceByUUIDReq) error {
	if req.GetUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}
	return nil
}
