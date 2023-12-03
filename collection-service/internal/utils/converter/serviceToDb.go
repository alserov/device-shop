package converter

import (
	"github.com/alserov/device-shop/collection-service/internal/db"
	"github.com/alserov/device-shop/collection-service/internal/service"
)

func ServiceDeviceToRepoStruct(req service.Device) db.Device {
	return db.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}

func RepoDeviceToServiceStruct(req db.Device) service.Device {
	return service.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}
