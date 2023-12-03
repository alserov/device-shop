package converter

import (
	"github.com/alserov/device-shop/collection-service/internal/service"
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/alserov/device-shop/proto/gen/device"
)

func PbDeviceToServiceStruct(req *device.Device) service.Device {
	return service.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}

func ServiceDeviceToPb(req service.Device) *device.Device {
	return &device.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}

func PbChangeCollectionReqTpServiceStruct(req *collection.ChangeCollectionReq) service.ChangeCollectionReq {
	return service.ChangeCollectionReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	}
}
