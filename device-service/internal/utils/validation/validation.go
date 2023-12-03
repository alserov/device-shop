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

func ValidateCreateDeviceReq(req *device.CreateDeviceReq) error {
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

func ValidateDeleteDeviceReq(req *device.DeleteDeviceReq) error {
	if req.GetUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	return nil
}

func ValidateUpdateDeviceReq(req *device.UpdateDeviceReq) error {
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

func ValidateGetDeviceByTitleReq(req *device.GetDeviceByTitleReq) error {
	if req.GetTitle() == "" {
		return status.Error(codes.Internal, emptyTitle)
	}
	return nil
}

func ValidateGetDevicesByManufacturerReq(req *device.GetByManufacturer) error {
	if req.GetManufacturer() == "" {
		return status.Error(codes.InvalidArgument, emptyManu)
	}
	return nil
}

func ValidateGetDevicesByPrice(req *device.GetByPrice) error {
	if req.GetMin() < 0 {
		return status.Error(codes.InvalidArgument, invalidPrice)
	}

	if req.Max <= req.Min {
		return status.Error(codes.InvalidArgument, invalidMaxPrice)
	}
	return nil
}

func ValidateGetDeviceByUUID(req *device.GetDeviceByUUIDReq) error {
	if req.GetUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}
	return nil
}
