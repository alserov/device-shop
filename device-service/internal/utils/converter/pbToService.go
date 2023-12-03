package converter

import (
	"github.com/alserov/device-shop/device-service/internal/service/models"
	"github.com/alserov/device-shop/proto/gen/device"
	"strings"
)

func CreateDeviceToService(req *device.CreateDeviceReq) models.CreateDeviceReq {
	return models.CreateDeviceReq{
		Title:        strings.ToLower(req.Title),
		Description:  strings.ToLower(req.Description),
		Price:        req.Price,
		Manufacturer: strings.ToLower(req.Manufacturer),
		Amount:       req.Amount,
	}
}

func UpdateDeviceToService(req *device.UpdateDeviceReq) models.UpdateDeviceReq {
	return models.UpdateDeviceReq{
		UUID:        req.UUID,
		Title:       strings.ToLower(req.Title),
		Description: strings.ToLower(req.Description),
		Price:       req.Price,
	}
}

func DeviceToPb(d models.Device) *device.Device {
	return &device.Device{
		UUID:         d.UUID,
		Title:        d.Title,
		Description:  d.Description,
		Price:        d.Price,
		Manufacturer: d.Manufacturer,
		Amount:       d.Amount,
	}
}

func GetByPriceToService(req *device.GetByPrice) models.GetByPrice {
	return models.GetByPrice{
		Max: req.Max,
		Min: req.Min,
	}
}
