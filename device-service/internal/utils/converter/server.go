package converter

import (
	"github.com/alserov/device-shop/device-service/internal/service/models"
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/alserov/device-shop/proto/gen/device"
	"strings"
)

type ServerConverter struct {
	Admin  ServerAdmin
	Device ServerDevice
}

func NewServerConverter() *ServerConverter {
	return &ServerConverter{
		Admin:  &serverAdmin{},
		Device: &serverDevice{},
	}
}

type serverAdmin struct{}
type ServerAdmin interface {
	CreateDeviceToService(req *device.CreateDeviceReq) models.CreateDeviceReq
	UpdateDeviceToService(req *device.UpdateDeviceReq) models.UpdateDeviceReq
	IncreaseDeviceAmountToService(req *device.IncreaseDeviceAmountByUUIDReq) models.IncreaseDeviceAmountReq
	RemoveDeviceFromCollectionsReqToPb(req *device.DeleteDeviceReq) *collection.RemoveDeletedDeviceReq
}

type serverDevice struct{}
type ServerDevice interface {
	DeviceToPb(d models.Device) *device.Device
	GetDevicesByPriceToService(req *device.GetByPrice) models.GetByPrice
	GetAllDevicesReqToService(req *device.GetAllDevicesReq) models.GetAllDevicesReq
	GetAllDevicesResToPb(res []*models.Device) *device.DevicesRes
}

func (*serverAdmin) CreateDeviceToService(req *device.CreateDeviceReq) models.CreateDeviceReq {
	return models.CreateDeviceReq{
		Title:        strings.ToLower(req.Title),
		Description:  strings.ToLower(req.Description),
		Price:        req.Price,
		Manufacturer: strings.ToLower(req.Manufacturer),
		Amount:       req.Amount,
	}
}

func (*serverAdmin) IncreaseDeviceAmountToService(req *device.IncreaseDeviceAmountByUUIDReq) models.IncreaseDeviceAmountReq {
	return models.IncreaseDeviceAmountReq{
		Amount:     req.Amount,
		DeviceUUID: req.DeviceUUID,
	}
}

func (*serverAdmin) UpdateDeviceToService(req *device.UpdateDeviceReq) models.UpdateDeviceReq {
	return models.UpdateDeviceReq{
		UUID:        req.UUID,
		Title:       strings.ToLower(req.Title),
		Description: strings.ToLower(req.Description),
		Price:       req.Price,
	}
}

func (s *serverAdmin) RemoveDeviceFromCollectionsReqToPb(req *device.DeleteDeviceReq) *collection.RemoveDeletedDeviceReq {
	return &collection.RemoveDeletedDeviceReq{
		DeviceUUID: req.UUID,
	}
}

func (*serverDevice) DeviceToPb(d models.Device) *device.Device {
	return &device.Device{
		UUID:         d.UUID,
		Title:        d.Title,
		Description:  d.Description,
		Price:        d.Price,
		Manufacturer: d.Manufacturer,
		Amount:       d.Amount,
	}
}

func (s *serverDevice) GetAllDevicesReqToService(req *device.GetAllDevicesReq) models.GetAllDevicesReq {
	return models.GetAllDevicesReq{
		Amount: req.GetAmount(),
		Index:  req.Index,
	}
}

func (*serverDevice) GetDevicesByPriceToService(req *device.GetByPrice) models.GetByPrice {
	return models.GetByPrice{
		Max: req.Max,
		Min: req.Min,
	}
}

func (*serverDevice) GetAllDevicesResToPb(res []*models.Device) *device.DevicesRes {
	var devices device.DevicesRes
	for _, d := range res {
		device := &device.Device{
			UUID:         d.UUID,
			Title:        d.Title,
			Description:  d.Description,
			Price:        d.Price,
			Manufacturer: d.Manufacturer,
			Amount:       d.Amount,
		}
		devices.Devices = append(devices.Devices, device)
	}

	return &devices
}
