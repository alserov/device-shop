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
	CancelOrder
	CreateOrder
	CheckOrder
	UpdateOrder
}

type CancelOrder interface {
	DeviceToIncreaseDeviceAmountPb(req *device.Device) *device.IncreaseDeviceAmountByUUIDReq
}

type UpdateOrder interface {
	UpdateOrderReqToService(req *order.UpdateOrderReq) models.UpdateOrderReq
	UpdateOrderResToPb(status string) *order.UpdateOrderRes
}

type CheckOrder interface {
	CheckOrderResToPb(res models.CheckOrderRes, devices []*device.Device) *order.CheckOrderRes
}

type CreateOrder interface {
	CreateOrderResToPb(res string) *order.CreateOrderRes
	CreateOrderReqToService(req *order.CreateOrderReq, orderPrice float32) models.CreateOrderReq
}

func NewServerConverter() ServerConverter {
	return &serverConverter{}
}

func (s *serverConverter) DeviceToIncreaseDeviceAmountPb(req *device.Device) *device.IncreaseDeviceAmountByUUIDReq {
	return &device.IncreaseDeviceAmountByUUIDReq{
		DeviceUUID: req.UUID,
		Amount:     req.Amount,
	}
}

func (s *serverConverter) CreateOrderReqToService(req *order.CreateOrderReq, orderPrice float32) models.CreateOrderReq {
	return models.CreateOrderReq{
		OrderDevices: s.pbOrderDevicesToService(req.OrderDevices),
		UserUUID:     req.UserUUID,
		OrderPrice:   orderPrice,
	}
}

func (*serverConverter) pbOrderDevicesToService(devices []*order.OrderDevice) []models.OrderDevice {
	var orderDevices []models.OrderDevice

	for _, device := range devices {
		d := models.OrderDevice{
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

func (*serverConverter) CreateOrderResToPb(uuid string) *order.CreateOrderRes {
	return &order.CreateOrderRes{
		OrderUUID: uuid,
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
