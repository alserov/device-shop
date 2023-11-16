package utils

import (
	"encoding/json"
	device "github.com/alserov/device-shop/device-service/pkg/entity"
	"github.com/alserov/device-shop/gateway/pkg/models"
	order "github.com/alserov/device-shop/order-service/pkg/entity"
	"github.com/alserov/device-shop/proto/gen"
	user "github.com/alserov/device-shop/user-service/pkg/entity"
	"net/http"
)

func RequestToPBMessage[T any, B any](r *http.Request, fn func(str *T) *B) (*B, error) {
	var req T
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if err := models.Validate(&req); err != nil {
		return nil, err
	}

	msg := fn(&req)

	return msg, nil
}

func DeviceToPB(str *device.Device) *pb.Device {
	return &pb.Device{
		UUID:         str.UUID,
		Title:        str.Title,
		Description:  str.Manufacturer,
		Price:        str.Price,
		Manufacturer: str.Manufacturer,
		Amount:       str.Amount,
	}
}

func CreateOrderToPB(str *order.CreateOrderReq) *pb.CreateOrderReq {
	return &pb.CreateOrderReq{
		UserUUID: str.UserUUID,
	}
}

func UpdateOrderToPB(str *order.UpdateOrderReq) *pb.UpdateOrderReq {
	return &pb.UpdateOrderReq{
		Status:    str.Status,
		OrderUUID: str.OrderUUID,
	}
}

func CreateDeviceToPB(str *device.Device) *pb.CreateReq {
	return &pb.CreateReq{
		Title:        str.Title,
		Description:  str.Description,
		Price:        str.Price,
		Manufacturer: str.Manufacturer,
		Amount:       str.Amount,
	}
}

func UpdateDeviceToPB(str *device.UpdateDeviceReq) *pb.UpdateReq {
	return &pb.UpdateReq{
		Title:       str.Title,
		Description: str.Description,
		Price:       str.Price,
		UUID:        str.UUID,
	}
}

func SignupReqToPB(str *user.SignupReq) *pb.SignupReq {
	return &pb.SignupReq{
		Username: str.Username,
		Password: str.Password,
		Email:    str.Email,
	}
}

func LoginReqToPB(str *user.LoginReq) *pb.LoginReq {
	return &pb.LoginReq{
		Username: str.Username,
		Password: str.Password,
	}
}

func AddReqToPB(str *user.AddToCollectionReq) *pb.AddReq {
	return &pb.AddReq{
		DeviceUUID: str.DeviceUUID,
		UserUUID:   str.UserUUID,
	}
}

func RemoveReqToPB(str *device.RemoveDeviceReq) *pb.RemoveReq {
	return &pb.RemoveReq{
		DeviceUUID: str.DeviceUUID,
		UserUUID:   str.UserUUID,
	}
}

func GetAllReqToPB(str *device.GetAllDevicesReq) *pb.GetAllReq {
	return &pb.GetAllReq{
		Index:  uint32(str.Index),
		Amount: uint32(str.Amount),
	}
}
