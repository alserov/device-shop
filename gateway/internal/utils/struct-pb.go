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
	return &pb.Device{}
}

func CreateOrderToPB(str *order.CreateOrderReq) *pb.CreateOrderReq {
	return &pb.CreateOrderReq{}
}

func UpdateOrderToPB(str *order.UpdateOrderReq) *pb.UpdateOrderReq {
	return &pb.UpdateOrderReq{}
}

func CheckOrderToPB(str *order.CheckOrderReq) *pb.CheckOrderReq {
	return &pb.CheckOrderReq{}
}

func CreateDeviceToPB(str *device.Device) *pb.CreateReq {
	return &pb.CreateReq{
		Title:        str.Title,
		Description:  str.Description,
		Price:        str.Price,
		Manufacturer: str.Manufacturer,
	}
}

func UpdateDeviceToPB(str *device.UpdateDeviceReq) *pb.UpdateReq {
	return &pb.UpdateReq{}
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
	return &pb.AddReq{}
}

func RemoveReqToPB(str *device.RemoveDeviceReq) *pb.RemoveReq {
	return &pb.RemoveReq{}
}

func GetAllReqToPB(str *device.GetAllDevicesReq) *pb.GetAllReq {
	return &pb.GetAllReq{}
}
