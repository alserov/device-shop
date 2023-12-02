package converter

import (
	"github.com/alserov/admin-service/internal/service"
	"github.com/alserov/device-shop/proto/gen/admin"
	"strings"
)

func CreateDeviceToServiceStruct(r *admin.CreateDeviceReq) service.CreateDeviceReq {
	return service.CreateDeviceReq{
		Title:        strings.ToLower(r.Title),
		Description:  strings.ToLower(r.Description),
		Price:        r.Price,
		Manufacturer: strings.ToLower(r.Manufacturer),
		Amount:       r.Amount,
	}
}

func UpdateDeviceToServiceStruct(r *admin.UpdateDeviceReq) service.UpdateDeviceReq {
	return service.UpdateDeviceReq{
		UUID:        r.UUID,
		Title:       strings.ToLower(r.Title),
		Description: strings.ToLower(r.Description),
		Price:       r.Price,
	}
}
