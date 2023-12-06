package converter

import (
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/alserov/device-shop/proto/gen/order"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ServerConverter struct {
}

func NewServerConverter() *ServerConverter {
	return &ServerConverter{}
}

func (*ServerConverter) CreateOrderReqToService(req *order.CreateOrderReq, orderPrice float32) models.CreateOrderReq {
	return models.CreateOrderReq{
		OrderDevices: pbOrderDevicesToService(req.OrderDevices),
		UserUUID:     req.UserUUID,
		OrderPrice:   orderPrice,
	}
}

func pbOrderDevicesToService(devices []*order.OrderDevice) []*models.OrderDevice {
	var orderDevices []*models.OrderDevice

	for _, device := range devices {
		d := &models.OrderDevice{
			DeviceUUID: device.DeviceUUID,
			Amount:     device.Amount,
		}
		orderDevices = append(orderDevices, d)
	}

	return orderDevices
}

func toTimepb(t *time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func (*ServerConverter) CreateOrderResToPb(res models.CreateOrderRes) *order.CreateOrderRes {
	return &order.CreateOrderRes{
		OrderUUID: res.OrderUUID,
	}
}

func (*ServerConverter) CheckOrderReqToService(req *order.CheckOrderReq) models.CheckOrderReq {
	return models.CheckOrderReq{
		OrderUUID: req.OrderUUID,
	}
}

func (*ServerConverter) CheckOrderResToPb(res models.CheckOrderRes, devices []*device.Device) *order.CheckOrderRes {
	return &order.CheckOrderRes{
		Price:     res.OrderPrice,
		Status:    res.Status,
		CreatedAt: toTimepb(res.CreatedAt),
		Devices:   devices,
	}
}

func (*ServerConverter) UpdateOrderReqToService(req *order.UpdateOrderReq) models.UpdateOrderReq {
	return models.UpdateOrderReq{
		OrderUUID: req.OrderUUID,
		Status:    req.Status,
	}
}

func (*ServerConverter) UpdateOrderResToPb(status string) *order.UpdateOrderRes {
	return &order.UpdateOrderRes{
		Status: status,
	}
}
