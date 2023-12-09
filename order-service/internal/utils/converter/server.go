package converter

import (
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/alserov/device-shop/proto/gen/order"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type serverConverter struct{}

type ServerConverter interface {
	CreateOrderResToPb(res models.CreateOrderRes) *order.CreateOrderRes
	CreateOrderReqToService(req *order.CreateOrderReq, orderPrice float32) models.CreateOrderReq
	CheckOrderReqToService(req *order.CheckOrderReq) models.CheckOrderReq
	CheckOrderResToPb(res models.CheckOrderRes, devices []*device.Device) *order.CheckOrderRes
	UpdateOrderReqToService(req *order.UpdateOrderReq) models.UpdateOrderReq
	UpdateOrderResToPb(status string) *order.UpdateOrderRes
}

func NewServerConverter() ServerConverter {
	return &serverConverter{}
}

func (*serverConverter) CreateOrderReqToService(req *order.CreateOrderReq, orderPrice float32) models.CreateOrderReq {
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

func (*serverConverter) CreateOrderResToPb(res models.CreateOrderRes) *order.CreateOrderRes {
	return &order.CreateOrderRes{
		OrderUUID: res.OrderUUID,
	}
}

func (*serverConverter) CheckOrderReqToService(req *order.CheckOrderReq) models.CheckOrderReq {
	return models.CheckOrderReq{
		OrderUUID: req.OrderUUID,
	}
}

func (*serverConverter) CheckOrderResToPb(res models.CheckOrderRes, devices []*device.Device) *order.CheckOrderRes {
	return &order.CheckOrderRes{
		Price:     res.OrderPrice,
		Status:    res.Status,
		CreatedAt: toTimepb(res.CreatedAt),
		Devices:   devices,
	}
}

func (*serverConverter) UpdateOrderReqToService(req *order.UpdateOrderReq) models.UpdateOrderReq {
	return models.UpdateOrderReq{
		OrderUUID: req.OrderUUID,
		Status:    req.Status,
	}
}

func (*serverConverter) UpdateOrderResToPb(status string) *order.UpdateOrderRes {
	return &order.UpdateOrderRes{
		Status: status,
	}
}
