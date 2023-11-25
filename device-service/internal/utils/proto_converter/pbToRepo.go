package proto_converter

import (
	"github.com/alserov/device-shop/device-service/internal/db"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"strings"
)

func CreateDeviceToRepoStruct(s *pb.CreateDeviceReq) db.Device {
	return db.Device{
		UUID:         uuid.New().String(),
		Title:        strings.ToLower(s.Title),
		Description:  strings.ToLower(s.Description),
		Price:        s.Price,
		Manufacturer: strings.ToLower(s.Manufacturer),
		Amount:       s.Amount,
	}
}

func UpdateDeviceToRepoStruct(s *pb.UpdateDeviceReq) db.UpdateDevice {
	return db.UpdateDevice{
		UUID:        s.UUID,
		Title:       strings.ToLower(s.Title),
		Description: strings.ToLower(s.Description),
		Price:       s.Price,
	}
}
