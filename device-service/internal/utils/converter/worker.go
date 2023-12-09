package converter

import (
	"github.com/alserov/device-shop/device-service/internal/broker/worker/models"
	repo "github.com/alserov/device-shop/device-service/internal/db/models"
)

type workerConverter struct {
}

type WorkerConverter interface {
	OrderDevicesToRepo(req []*models.OrderDevice) []*repo.OrderDevice
}

func NewWorkerConverter() WorkerConverter {
	return &workerConverter{}
}

func (*workerConverter) OrderDevicesToRepo(req []*models.OrderDevice) []*repo.OrderDevice {
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
