package converter

import (
	repo "github.com/alserov/device-shop/order-service/internal/db/models"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/status"
	"time"
)

type serviceConverter struct{}

type ServiceConverter interface {
	CreateOrderReqToRepo(req models.CreateOrderReq, orderUUID string) repo.CreateOrderReq
	CreateOrderResToService(orderUUID string) models.CreateOrderRes
	CheckOrderToService(res repo.CheckOrderRes) models.CheckOrderRes
}

func NewServiceConverter() ServiceConverter {
	return &serviceConverter{}
}

func (*serviceConverter) CreateOrderReqToRepo(req models.CreateOrderReq, orderUUID string) repo.CreateOrderReq {
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

func repoOrderDevicesToService(res []*repo.OrderDevice) []*models.OrderDevice {
	var devices []*models.OrderDevice

	for _, od := range res {
		d := &models.OrderDevice{
			DeviceUUID: od.DeviceUUID,
			Amount:     od.Amount,
		}
		devices = append(devices, d)
	}

	return devices
}

func (*serviceConverter) CreateOrderResToService(orderUUID string) models.CreateOrderRes {
	return models.CreateOrderRes{
		OrderUUID: orderUUID,
	}
}

func (*serviceConverter) CheckOrderToService(res repo.CheckOrderRes) models.CheckOrderRes {
	return models.CheckOrderRes{
		Status:       status.StatusCodeToString(res.Status),
		CreatedAt:    res.CreatedAt,
		OrderPrice:   res.OrderPrice,
		OrderDevices: repoOrderDevicesToService(res.OrderDevices),
	}
}
