package utils

import (
	"encoding/json"
	"github.com/alserov/device-shop/gateway/pkg/models"
	pb "github.com/alserov/device-shop/proto/gen"
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

func DeviceToPB(str *models.Device) *pb.Device {
	return &pb.Device{}
}

func CreateOrderToPB(str *models.CreateOrderReq) *pb.CreateOrderReq {
	return &pb.CreateOrderReq{}
}

func UpdateOrderToPB(str *models.UpdateOrderReq) *pb.UpdateOrderReq {
	return &pb.UpdateOrderReq{}
}

func CheckOrderToPB(str *models.CheckOrderReq) *pb.CheckOrderReq {
	return &pb.CheckOrderReq{}
}

func CreateDeviceToPB(str *models.Device) *pb.CreateReq {
	return &pb.CreateReq{}
}
