package converter

import (
	repo "github.com/alserov/device-shop/order-service/internal/db/models"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/status"
	"time"
)

type ServiceConverter struct{}

func NewServiceConverter() *ServiceConverter {
	return &ServiceConverter{}
}

func (*ServiceConverter) CreateOrderReqToRepo(req models.CreateOrderReq, orderUUID string) repo.CreateOrderReq {
	now := time.Now()
	return repo.CreateOrderReq{
		OrderUUID:    orderUUID,
		UserUUID:     req.UserUUID,
		Status:       status.CREATING_CODE,
		OrderPrice:   req.OrderPrice,
		CreatedAt:    &now,
		OrderDevices: serviceOrderDevicesToRepo(req.OrderDevices),
	}
}

func serviceOrderDevicesToRepo(req []*models.OrderDevice) []*repo.OrderDevice {
	var devices []*repo.OrderDevice

	for _, od := range req {
		d := &repo.OrderDevice{
			DeviceUUID: od.DeviceUUID,
			Amount:     od.Amount,
		}
		devices = append(devices, d)
	}

	return devices
}

func (*ServiceConverter) CreateOrderResToService(orderUUID string) models.CreateOrderRes {
	return models.CreateOrderRes{
		OrderUUID: orderUUID,
	}
}

func (*ServiceConverter) CheckOrderToService(res repo.CheckOrderRes) models.CheckOrderRes {
	return models.CheckOrderRes{
		Status:      status.StatusCodeToString(res.Status),
		CreatedAt:   res.CreatedAt,
		OrderPrice:  res.OrderPrice,
		DeviceUUIDs: res.DeviceUUIDs,
	}
}
