package converter

import (
	"github.com/alserov/device-shop/device-service/internal/broker/worker/models"
	repo "github.com/alserov/device-shop/device-service/internal/db/models"
)

type WorkerConverter struct {
}

func NewWorkerConverter() *WorkerConverter {
	return &WorkerConverter{}
}

func (*WorkerConverter) OrderDevicesToRepo(req []*models.OrderDevice) []*repo.OrderDevice {
	var devices []*repo.OrderDevice

	for _, d := range req {
		device := &repo.OrderDevice{
			DeviceUUID: d.DeviceUUID,
			Amount:     d.Amount,
		}
		devices = append(devices, device)
	}

	return devices
}
