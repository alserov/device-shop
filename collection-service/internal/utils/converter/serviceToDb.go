package converter

import (
	repo "github.com/alserov/device-shop/collection-service/internal/db/models"
	"github.com/alserov/device-shop/collection-service/internal/service/models"
)

type ServiceConverter struct {
	Device     ServiceDevice
	Collection ServiceCollection
}

func NewServiceConverter() *ServiceConverter {
	return &ServiceConverter{
		Device:     &serviceDevice{},
		Collection: &serviceCollection{},
	}
}

type serviceDevice struct{}
type ServiceDevice interface {
	DeviceToRepo(req models.Device) repo.Device
	DeviceToService(req repo.Device) models.Device
}

type serviceCollection struct{}
type ServiceCollection interface {
	CollectionToServer(coll []*repo.Device) []*models.Device
}

func (*serviceDevice) DeviceToRepo(req models.Device) repo.Device {
	return repo.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}

func (*serviceDevice) DeviceToService(req repo.Device) models.Device {
	return models.Device{
		UUID:         req.UUID,
		Title:        req.Title,
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Amount:       req.Amount,
	}
}

func (*serviceCollection) CollectionToServer(coll []*repo.Device) []*models.Device {
	var devices []*models.Device
	for _, d := range coll {
		device := models.Device{
			UUID:         d.UUID,
			Title:        d.Title,
			Description:  d.Description,
			Price:        d.Price,
			Manufacturer: d.Manufacturer,
			Amount:       d.Amount,
		}
		devices = append(devices, &device)
	}

	return devices
}
