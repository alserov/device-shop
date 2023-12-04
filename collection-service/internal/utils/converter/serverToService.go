package converter

import (
	"github.com/alserov/device-shop/collection-service/internal/service/models"
	coll "github.com/alserov/device-shop/proto/gen/collection"
	"github.com/alserov/device-shop/proto/gen/device"
)

type ServerConverter struct {
	Device     ServerDevice
	Collection Collection
}

func NewServerConverter() *ServerConverter {
	return &ServerConverter{
		Device:     &serverDevice{},
		Collection: &serverCollection{},
	}
}

type serverDevice struct{}
type ServerDevice interface {
	PbDeviceToService(req *device.Device) models.Device
	DeviceToPb(req models.Device) *device.Device
	GetDeviceByUUIDReq(uuid string) *device.GetDeviceByUUIDReq
}

type serverCollection struct{}
type Collection interface {
	ChangeCollectionReqToService(req *coll.ChangeCollectionReq) models.ChangeCollectionReq
	GetCollectionResToPb(res []*models.Device) *coll.GetCollectionRes
}

func (*serverDevice) PbDeviceToService(req *device.Device) models.Device {
	return models.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}

func (*serverDevice) DeviceToPb(req models.Device) *device.Device {
	return &device.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}

func (*serverCollection) ChangeCollectionReqToService(req *coll.ChangeCollectionReq) models.ChangeCollectionReq {
	return models.ChangeCollectionReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	}
}

func (*serverCollection) GetCollectionResToPb(res []*models.Device) *coll.GetCollectionRes {
	var devices []*device.Device
	for _, d := range res {
		device := &device.Device{
			UUID:         d.UUID,
			Title:        d.Title,
			Description:  d.Description,
			Price:        d.Price,
			Manufacturer: d.Manufacturer,
			Amount:       d.Amount,
		}
		devices = append(devices, device)
	}

	return &coll.GetCollectionRes{
		Devices: devices,
	}
}

func (*serverDevice) GetDeviceByUUIDReq(uuid string) *device.GetDeviceByUUIDReq {
	return &device.GetDeviceByUUIDReq{
		UUID: uuid,
	}
}
