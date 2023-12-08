package converter

import (
	repo "github.com/alserov/device-shop/device-service/internal/db/models"
	"github.com/alserov/device-shop/device-service/internal/service/models"
)

type ServiceConverter struct {
	Admin  ServiceAdminConverter
	Device ServiceDeviceConverter
}

func NewServiceConverter() *ServiceConverter {
	return &ServiceConverter{
		Admin:  &serviceAdminConverter{},
		Device: &serviceDeviceConverter{},
	}
}

type serviceAdminConverter struct{}
type ServiceAdminConverter interface {
	DeviceToRepo(req models.CreateDeviceReq) repo.Device
	UpdateDeviceReqToRepo(req models.UpdateDeviceReq) repo.UpdateDevice
}

type serviceDeviceConverter struct{}
type ServiceDeviceConverter interface {
	DevicesToService(devices []*repo.Device) []*models.Device
	DeviceToService(device repo.Device) models.Device
}

func (*serviceAdminConverter) DeviceToRepo(req models.CreateDeviceReq) repo.Device {
	return repo.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Description:  req.Description,
		Amount:       req.Amount,
	}
}

func (*serviceAdminConverter) UpdateDeviceReqToRepo(req models.UpdateDeviceReq) repo.UpdateDevice {
	return repo.UpdateDevice{
		UUID:        req.UUID,
		Description: req.Description,
		Title:       req.Title,
		Price:       req.Price,
	}
}

func (*serviceDeviceConverter) DevicesToService(devices []*repo.Device) []*models.Device {
	res := make([]*models.Device, 0, len(devices))
	for _, device := range devices {
		pbDevice := &models.Device{
			UUID:         device.UUID,
			Title:        device.Title,
			Description:  device.Description,
			Manufacturer: device.Manufacturer,
			Price:        device.Price,
			Amount:       device.Amount,
		}
		res = append(res, pbDevice)
	}

	return res
}

func (*serviceDeviceConverter) DeviceToService(device repo.Device) models.Device {
	return models.Device{
		UUID:         device.UUID,
		Title:        device.Title,
		Description:  device.Description,
		Manufacturer: device.Manufacturer,
		Price:        device.Price,
		Amount:       device.Amount,
	}
}
